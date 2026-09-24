package validate

import (
	"context"
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"goml.dev/tools/goml-go-meta/internal/protocol"
)

func TestWitnessBundleAliasesAndOrigins(t *testing.T) {
	sources := map[string][]byte{
		"a.go": []byte("package caller\nimport _ \"bytes\"\n"),
		"b.go": []byte("package caller\nimport target \"bytes\"\ntype B = target.Buffer\n"),
		"c.go": []byte("package caller\nimport target \"strings\"\ntype C = target.Reader\n"),
		"d.go": []byte("package caller\nimport alternate \"bytes\"\ntype D = alternate.Buffer\n"),
	}
	path := "/path:with-colon/bundle.go"
	data, origins, err := bundleWitnesses("caller", path, sources)
	if err != nil {
		t.Fatal(err)
	}
	positions := token.NewFileSet()
	file, err := parser.ParseFile(positions, path, data, 0)
	if err != nil {
		t.Fatal(err)
	}
	config := types.Config{Importer: importer.Default()}
	if _, err := config.Check("caller", positions, []*ast.File{file}, nil); err != nil {
		t.Fatalf("invalid merged imports: %v\n%s", err, data)
	}
	for _, declaration := range file.Decls {
		general := declaration.(*ast.GenDecl)
		if general.Tok != token.TYPE {
			continue
		}
		name := general.Specs[0].(*ast.TypeSpec).Name.Name
		want := strings.ToLower(name) + ".go"
		position := positions.Position(declaration.Pos()).String()
		if files := origins.diagnostic(position); len(files) != 1 || files[0] != want {
			t.Errorf("%s: want %s, got %v", position, want, files)
		}
	}
	if files := origins.diagnostic("/other.go:4:2"); len(files) != 1 || files[0] != "/other.go" {
		t.Fatalf("lost unbundled origin: %v", files)
	}
}

func TestFileWitnessBundlePreservesExistingFile(t *testing.T) {
	r := fixture(t)
	r.CallerContext.LoadMode = "files"
	r.CallerContext.ImportPath = "command-line-arguments"
	r.Bindings = []protocol.Binding{binding("Empty", nil, nil)}
	path := filepath.Join(r.CallerContext.Directory, "goml_ffi_witness_bundle.go")
	contents := "package gen\nvar Existing = true\n"
	if err := os.WriteFile(path, []byte(contents), 0644); err != nil {
		t.Fatal(err)
	}
	response := Check(context.Background(), r)
	if len(response.Diagnostics) == 0 || !strings.Contains(response.Diagnostics[0].Message, "conflicts") {
		t.Fatalf("missing collision diagnostic: %+v", response)
	}
	data, err := os.ReadFile(path)
	if err != nil || string(data) != contents {
		t.Fatalf("existing file changed: %q, %v", data, err)
	}
}

func TestFileWitnessesBeyondDriverArgumentLimit(t *testing.T) {
	r := fixture(t)
	r.CallerContext.LoadMode = "files"
	r.CallerContext.ImportPath = "command-line-arguments"
	r.CallerContext.Directory = filepath.Join(r.CallerContext.Directory, strings.Repeat("long-path-", 20))
	if err := os.MkdirAll(r.CallerContext.Directory, 0755); err != nil {
		t.Fatal(err)
	}
	var source strings.Builder
	source.WriteString("package shim\n")
	argumentBytes := 0
	for i := range 160 {
		name := fmt.Sprintf("Value%d", i)
		fmt.Fprintf(&source, "func %s() int64 { return %d }\n", name, i)
		r.Bindings = append(r.Bindings, binding(name, nil, []protocol.Type{{Tag: "int64"}}))
		argumentBytes += len(filepath.Join(r.CallerContext.Directory, fmt.Sprintf("goml_ffi_witness_%d.go", i))) + 1
	}
	if argumentBytes <= 16383 {
		t.Fatal("fixture does not exceed the package driver's argument chunk size")
	}
	if err := os.WriteFile(filepath.Join(r.BuildContext.ModuleDir, "shim", "many.go"), []byte(source.String()), 0644); err != nil {
		t.Fatal(err)
	}
	buffer := protocol.Type{Tag: "named", ImportPath: "bytes", Name: "Buffer"}
	receiver := protocol.Type{Tag: "pointer", Element: &buffer}
	for _, id := range []string{"method", "duplicate-method"} {
		r.Bindings = append(r.Bindings, methodBinding(id, "Len", []protocol.Type{receiver}, []protocol.Type{{Tag: "int"}}))
	}
	r.TypeQueries = []protocol.TypeQuery{
		{ID: "declaration", Mode: "declaration", ImportPath: "bytes", Name: "Buffer"},
		{ID: "instance", ImportPath: "bytes", Name: "Buffer"},
		{ID: "other-import", ImportPath: "strings", Name: "Reader"},
	}
	for _, invalid := range []bool{false, true} {
		t.Run(fmt.Sprint(invalid), func(t *testing.T) {
			request := r
			request.Bindings = append([]protocol.Binding(nil), r.Bindings...)
			request.TypeQueries = append([]protocol.TypeQuery(nil), r.TypeQueries...)
			if invalid {
				request.Bindings = append(request.Bindings, binding("Scalar", []protocol.Type{{Tag: "string"}}, []protocol.Type{{Tag: "int64"}}))
				request.TypeQueries = append(request.TypeQueries, protocol.TypeQuery{ID: "invalid", ImportPath: "bytes", Name: "Missing"})
			}
			response := Check(context.Background(), request)
			if len(response.Diagnostics) != 0 || len(response.Bindings) != len(request.Bindings) {
				t.Fatalf("request failed: %+v", response)
			}
			for _, result := range response.Bindings {
				want := "verified"
				if result.ID == "Scalar" {
					want = "failed"
				}
				if result.Status != want {
					t.Errorf("binding %s: want %s, got %+v", result.ID, want, result)
				}
				if strings.HasSuffix(result.ID, "method") && result.ActualSignature == "" {
					t.Errorf("method signature missing: %+v", result)
				}
			}
			if len(response.TypeResults) != len(request.TypeQueries) {
				t.Fatalf("missing type results: %+v", response)
			}
			for _, result := range response.TypeResults {
				want := "verified"
				if result.ID == "invalid" {
					want = "failed"
				}
				if result.Status != want {
					t.Errorf("type %s: want %s, got %+v", result.ID, want, result)
				}
			}
		})
	}
	files, err := os.ReadDir(r.CallerContext.Directory)
	if err != nil || len(files) != 0 {
		t.Fatalf("caller directory changed: %v %v", files, err)
	}
}
