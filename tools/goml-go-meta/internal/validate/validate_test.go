package validate

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"goml.dev/tools/goml-go-meta/internal/protocol"
)

func fixture(t *testing.T) protocol.Request {
	t.Helper()
	dir := filepath.Join(t.TempDir(), "host project")
	write := func(path, contents string) {
		t.Helper()
		if err := os.MkdirAll(filepath.Dir(filepath.Join(dir, path)), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(dir, path), []byte(contents), 0644); err != nil {
			t.Fatal(err)
		}
	}
	write("go.mod", "module example.com/host\n\ngo 1.25.0\n")
	write("shim/shim.go", `package shim

type Alias = int64
type Defined int64
func Scalar(x int64) int64 { return x }
func AliasValue(x Alias) Alias { return x }
func DefinedValue(x Defined) Defined { return x }
func Any(x any) any { return x }
func Pair(x int64) (int64, bool) { return x, true }
func Empty() {}
func Identity[T any](x T) T { return x }
func Variadic(x ...int64) int64 { return x[0] }
func private() {}
var FunctionValue = func() {}
`)
	if err := os.MkdirAll(filepath.Join(dir, "gen"), 0755); err != nil {
		t.Fatal(err)
	}
	goPath, err := exec.LookPath("go")
	if err != nil {
		t.Fatal(err)
	}
	return protocol.Request{ProtocolVersion: 1, BuildContext: protocol.BuildContext{
		GoExecutable: goPath, Toolchain: runtime.Version(), ModuleDir: dir, GOOS: runtime.GOOS, GOARCH: runtime.GOARCH,
		CGOEnabled: "0", GO111MODULE: "on", GOWORK: "off", Dependencies: "readonly", Network: "off",
	}, CallerContext: protocol.CallerContext{LoadMode: "package", ImportPath: "example.com/host/gen", Directory: filepath.Join(dir, "gen"), Package: "gen"}}
}

func binding(symbol string, params, results []protocol.Type) protocol.Binding {
	return protocol.Binding{ID: symbol, ImportPath: "example.com/host/shim", Symbol: symbol, Parameters: params, Results: results, CallMode: "ordinary"}
}

func TestWitnessCalls(t *testing.T) {
	r := fixture(t)
	i64 := []protocol.Type{{Tag: "int64"}}
	anyType := []protocol.Type{{Tag: "any"}}
	r.Bindings = []protocol.Binding{
		binding("Scalar", i64, i64), binding("AliasValue", i64, i64), binding("Any", anyType, anyType),
		binding("Pair", i64, []protocol.Type{{Tag: "int64"}, {Tag: "bool"}}), binding("Empty", nil, nil),
		binding("Identity", i64, i64), binding("Variadic", i64, i64),
	}
	response := Check(context.Background(), r)
	if len(response.Diagnostics) != 0 {
		t.Fatalf("request failed: %+v", response)
	}
	for _, result := range response.Bindings {
		if result.Status != "verified" {
			t.Errorf("%s: %+v", result.ID, result)
		}
	}
	if len(response.Bindings) != len(r.Bindings) {
		t.Fatal("missing binding results")
	}
	files, err := os.ReadDir(r.CallerContext.Directory)
	if err != nil || len(files) != 0 {
		t.Fatalf("witness modified caller directory: %v %v", files, err)
	}
	if _, err := os.Stat(filepath.Join(r.BuildContext.ModuleDir, "go.sum")); !os.IsNotExist(err) {
		t.Fatal("created go.sum")
	}
}

func TestInvalidCalls(t *testing.T) {
	r := fixture(t)
	i64 := []protocol.Type{{Tag: "int64"}}
	r.Bindings = []protocol.Binding{
		binding("DefinedValue", i64, i64), binding("Missing", nil, nil), binding("private", nil, nil),
		binding("Scalar", nil, i64), binding("Pair", i64, i64), binding("FunctionValue", nil, nil),
	}
	response := Check(context.Background(), r)
	if len(response.Bindings) != len(r.Bindings) {
		t.Fatalf("missing results: %+v", response)
	}
	for _, result := range response.Bindings {
		if result.Status != "failed" || len(result.Diagnostics) == 0 {
			t.Errorf("accepted invalid binding: %+v", result)
		}
	}
}

func TestInternalImport(t *testing.T) {
	r := fixture(t)
	dir := filepath.Join(r.BuildContext.ModuleDir, "owner", "internal", "secret")
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "secret.go"), []byte("package secret\nfunc Value() int64 { return 1 }\n"), 0644); err != nil {
		t.Fatal(err)
	}
	b := binding("Value", nil, []protocol.Type{{Tag: "int64"}})
	b.ImportPath = "example.com/host/owner/internal/secret"
	r.Bindings = []protocol.Binding{b}
	result := Check(context.Background(), r)
	if result.Bindings[0].Status != "failed" {
		t.Fatalf("accepted forbidden internal import: %+v", result)
	}
	found := false
	for _, d := range result.Bindings[0].Diagnostics {
		found = found || strings.Contains(d.Message, "internal")
	}
	if !found {
		t.Fatalf("missing internal diagnostic: %+v", result)
	}
	r.CallerContext.Directory = filepath.Join(r.BuildContext.ModuleDir, "owner", "gen")
	r.CallerContext.ImportPath = "example.com/host/owner/gen"
	if err := os.MkdirAll(r.CallerContext.Directory, 0755); err != nil {
		t.Fatal(err)
	}
	result = Check(context.Background(), r)
	if result.Bindings[0].Status != "verified" {
		t.Fatalf("rejected permitted internal import: %+v", result)
	}
}

func TestDifferentialGoBuild(t *testing.T) {
	r := fixture(t)
	if err := os.WriteFile(filepath.Join(r.BuildContext.ModuleDir, "main.go"), []byte("package main\nfunc Entry() {}\nfunc main() {}\n"), 0644); err != nil {
		t.Fatal(err)
	}
	missingPackage := binding("Entry", nil, nil)
	missingPackage.ImportPath = "example.com/host/missing"
	mainPackage := binding("Entry", nil, nil)
	mainPackage.ImportPath = "example.com/host"
	i64 := []protocol.Type{{Tag: "int64"}}
	boolean := []protocol.Type{{Tag: "bool"}}
	anyType := []protocol.Type{{Tag: "any"}}
	pair := []protocol.Type{{Tag: "int64"}, {Tag: "bool"}}
	defined := []protocol.Type{{Tag: "named", ImportPath: "example.com/host/shim", Name: "Defined"}}
	for _, test := range []struct {
		name    string
		binding protocol.Binding
	}{
		{"scalar", binding("Scalar", i64, i64)},
		{"alias", binding("AliasValue", i64, i64)},
		{"any", binding("Any", anyType, anyType)},
		{"assignable-argument", binding("Any", i64, anyType)},
		{"nonassignable-result", binding("Any", anyType, i64)},
		{"tuple", binding("Pair", i64, pair)},
		{"defined", binding("DefinedValue", defined, defined)},
		{"defined-parameter-mismatch", binding("DefinedValue", i64, defined)},
		{"defined-result-mismatch", binding("DefinedValue", defined, i64)},
		{"generic-inference", binding("Identity", i64, i64)},
		{"variadic-fixed-argument", binding("Variadic", i64, i64)},
		{"unit", binding("Empty", nil, nil)},
		{"discard-result", binding("Scalar", i64, nil)},
		{"missing-package", missingPackage},
		{"main-package", mainPackage},
		{"missing-symbol", binding("Missing", nil, nil)},
		{"private-symbol", binding("private", nil, nil)},
		{"missing-argument", binding("Scalar", nil, i64)},
		{"extra-argument", binding("Scalar", []protocol.Type{{Tag: "int64"}, {Tag: "int64"}}, i64)},
		{"argument-type", binding("Scalar", boolean, i64)},
		{"result-type", binding("Scalar", i64, boolean)},
		{"missing-result", binding("Pair", i64, i64)},
		{"extra-result", binding("Scalar", i64, pair)},
		{"unit-result-mismatch", binding("Empty", nil, i64)},
	} {
		t.Run(test.name, func(t *testing.T) {
			b := test.binding
			r.Bindings = []protocol.Binding{b}
			response := Check(context.Background(), r)
			data, err := witness("gen", b, 0)
			if err != nil {
				t.Fatal(err)
			}
			file := filepath.Join(r.CallerContext.Directory, "actual.go")
			if err := os.WriteFile(file, data, 0644); err != nil {
				t.Fatal(err)
			}
			command := exec.Command(r.BuildContext.GoExecutable, "build", "-mod=readonly", "./gen")
			command.Dir = r.BuildContext.ModuleDir
			command.Env = environment(r.BuildContext)
			output, buildErr := command.CombinedOutput()
			if err := os.Remove(file); err != nil {
				t.Fatal(err)
			}
			verified := len(response.Bindings) == 1 && response.Bindings[0].Status == "verified"
			if verified != (buildErr == nil) {
				t.Fatalf("helper/compiler disagree: %+v, %v\n%s", response, buildErr, output)
			}
		})
	}
}

func TestFreshSourcesAndTargets(t *testing.T) {
	r := fixture(t)
	file := filepath.Join(r.BuildContext.ModuleDir, "shim", "target_linux.go")
	if err := os.WriteFile(file, []byte("package shim\nfunc Target(x int64) int64 { return x }\n"), 0644); err != nil {
		t.Fatal(err)
	}
	windowsFile := filepath.Join(r.BuildContext.ModuleDir, "shim", "target_windows.go")
	if err := os.WriteFile(windowsFile, []byte("package shim\nfunc Target(x string) string { return x }\n"), 0644); err != nil {
		t.Fatal(err)
	}
	r.BuildContext.GOOS = "linux"
	r.BuildContext.GOARCH = "amd64"
	r.Bindings = []protocol.Binding{binding("Target", []protocol.Type{{Tag: "int64"}}, []protocol.Type{{Tag: "int64"}})}
	before := Check(context.Background(), r)
	if before.Bindings[0].Status != "verified" {
		t.Fatalf("linux: %+v", before)
	}
	if err := os.WriteFile(file, []byte("package shim\nfunc Target(x int64) int64 { return x+1 }\n"), 0644); err != nil {
		t.Fatal(err)
	}
	after := Check(context.Background(), r)
	if after.Bindings[0].Status != "verified" || before.GoWorldIdentity == after.GoWorldIdentity {
		t.Fatalf("body change not observed: %+v", after)
	}
	r.BuildContext.GOOS = "windows"
	after = Check(context.Background(), r)
	if after.Bindings[0].Status != "failed" {
		t.Fatalf("windows signature not observed: %+v", after)
	}
	r.BuildContext.GOOS = "linux"
	if err := os.WriteFile(file, []byte("package shim\nfunc Target(x string) string { return x }\n"), 0644); err != nil {
		t.Fatal(err)
	}
	after = Check(context.Background(), r)
	if after.Bindings[0].Status != "failed" {
		t.Fatalf("signature change not observed: %+v", after)
	}
}

func TestDuplicateBindingAndMixedFailure(t *testing.T) {
	r := fixture(t)
	b := binding("Scalar", []protocol.Type{{Tag: "int64"}}, []protocol.Type{{Tag: "int64"}})
	duplicate := b
	duplicate.ID = "other::package::scalar"
	r.Bindings = []protocol.Binding{b, duplicate, binding("Missing", nil, nil)}
	response := Check(context.Background(), r)
	if len(response.Bindings) != 3 || response.Bindings[0].Status != "verified" || response.Bindings[1].Status != "verified" || response.Bindings[2].Status != "failed" {
		t.Fatalf("wrong per-binding status: %+v", response)
	}
	if response.Bindings[0].SignatureType != response.Bindings[1].SignatureType {
		t.Fatal("same type got different identities")
	}
}

func TestRecursiveReferencedTypes(t *testing.T) {
	r := fixture(t)
	file := filepath.Join(r.BuildContext.ModuleDir, "shim", "node.go")
	if err := os.WriteFile(file, []byte("package shim\ntype Node struct { Next *Node }\nfunc (n *Node) Get() *Node { return n.Next }\nfunc NodeValue() *Node { return nil }\n"), 0644); err != nil {
		t.Fatal(err)
	}
	r.Bindings = []protocol.Binding{binding("NodeValue", nil, []protocol.Type{{Tag: "any"}})}
	response := Check(context.Background(), r)
	if response.Bindings[0].Status != "verified" {
		t.Fatalf("any result rejected: %+v", response)
	}
	nodes := map[string]protocol.ReferencedType{}
	var named protocol.ReferencedType
	for _, node := range response.ReferencedTypes {
		if _, duplicate := nodes[node.ID]; duplicate {
			t.Fatal("duplicate type id")
		}
		nodes[node.ID] = node
		if node.Tag == "named" && node.Name == "Node" {
			named = node
		}
	}
	if named.ID == "" || named.ImportPath != "example.com/host/shim" || len(named.Methods) != 1 {
		t.Fatalf("missing named type: %+v", named)
	}
	structure := nodes[named.Underlying]
	if len(structure.Fields) != 1 {
		t.Fatal("missing struct field")
	}
	pointer := nodes[structure.Fields[0].Type]
	if pointer.Tag != "pointer" || pointer.Element != named.ID {
		t.Fatal("recursive type identity was not preserved")
	}
}

func TestMissingPackageAndMain(t *testing.T) {
	r := fixture(t)
	for _, path := range []string{"example.com/host/missing", "example.com/host/program"} {
		if strings.HasSuffix(path, "/program") {
			dir := filepath.Join(r.BuildContext.ModuleDir, "program")
			if err := os.MkdirAll(dir, 0755); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(filepath.Join(dir, "main.go"), []byte("package main\nfunc main() {}\nfunc Value() {}\n"), 0644); err != nil {
				t.Fatal(err)
			}
		}
		b := binding("Value", nil, nil)
		b.ImportPath = path
		r.Bindings = []protocol.Binding{b}
		response := Check(context.Background(), r)
		if len(response.Bindings) != 1 || response.Bindings[0].Status != "failed" {
			t.Fatalf("invalid package accepted: %+v", response)
		}
	}
}

func TestSleepDefinedType(t *testing.T) {
	r := fixture(t)
	b := binding("Sleep", []protocol.Type{{Tag: "int64"}}, nil)
	b.ImportPath = "time"
	r.Bindings = []protocol.Binding{b}
	response := Check(context.Background(), r)
	if response.Bindings[0].Status != "failed" || !strings.Contains(response.Bindings[0].ActualSignature, "time.Duration") {
		t.Fatalf("missing named-type mismatch: %+v", response)
	}
}

func TestToolAndCancellationFailures(t *testing.T) {
	r := fixture(t)
	r.Bindings = []protocol.Binding{binding("Empty", nil, nil)}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if response := Check(ctx, r); response.Bindings[0].Status != "failed" || response.Diagnostics[0].Code != "ffi-canceled" {
		t.Fatalf("cancellation: %+v", response)
	}
	r.BuildContext.GoExecutable = filepath.Join(t.TempDir(), "missing-go")
	if response := Check(context.Background(), r); response.Bindings[0].Status != "failed" || response.Diagnostics[0].Code != "ffi-tool-unavailable" {
		t.Fatalf("missing tool: %+v", response)
	}
}

func TestBuildTagsAndRemovedFiles(t *testing.T) {
	r := fixture(t)
	dir := filepath.Join(r.BuildContext.ModuleDir, "shim")
	ordinary := filepath.Join(dir, "tag_ordinary.go")
	special := filepath.Join(dir, "tag_special.go")
	if err := os.WriteFile(ordinary, []byte("//go:build !special\n\npackage shim\nfunc Tagged(x int64) int64 { return x }\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(special, []byte("//go:build special\n\npackage shim\nfunc Tagged(x string) string { return x }\n"), 0644); err != nil {
		t.Fatal(err)
	}
	r.Bindings = []protocol.Binding{binding("Tagged", []protocol.Type{{Tag: "int64"}}, []protocol.Type{{Tag: "int64"}})}
	if result := Check(context.Background(), r); result.Bindings[0].Status != "verified" {
		t.Fatalf("ordinary tag: %+v", result)
	}
	r.BuildContext.BuildTags = []string{"special"}
	if result := Check(context.Background(), r); result.Bindings[0].Status != "failed" {
		t.Fatalf("build tags ignored: %+v", result)
	}
	r.BuildContext.BuildTags = nil
	if err := os.Remove(ordinary); err != nil {
		t.Fatal(err)
	}
	if result := Check(context.Background(), r); result.Bindings[0].Status != "failed" {
		t.Fatalf("removed file ignored: %+v", result)
	}
}

func TestAllBridgeRepresentations(t *testing.T) {
	r := fixture(t)
	for _, tag := range []string{"bool", "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "float32", "float64", "string"} {
		b := binding("Identity", []protocol.Type{{Tag: tag}}, []protocol.Type{{Tag: tag}})
		b.ID = tag
		r.Bindings = append(r.Bindings, b)
	}
	length := int64(2)
	for i, ty := range []protocol.Type{
		{Tag: "array", Length: &length, Element: &protocol.Type{Tag: "int32"}},
		{Tag: "slice", Element: &protocol.Type{Tag: "uint8"}},
		{Tag: "map", Key: &protocol.Type{Tag: "string"}, Element: &protocol.Type{Tag: "int64"}},
		{Tag: "channel", Direction: "both", Element: &protocol.Type{Tag: "string"}},
		{Tag: "channel", Direction: "send", Element: &protocol.Type{Tag: "string"}},
		{Tag: "channel", Direction: "receive", Element: &protocol.Type{Tag: "string"}},
	} {
		b := binding("Identity", []protocol.Type{ty}, []protocol.Type{ty})
		b.ID = string(rune('a' + i))
		r.Bindings = append(r.Bindings, b)
	}
	result := Check(context.Background(), r)
	if len(result.Bindings) != len(r.Bindings) {
		t.Fatalf("missing results: %+v", result)
	}
	for _, b := range result.Bindings {
		if b.Status != "verified" {
			t.Errorf("bridge representation failed: %+v", b)
		}
	}
}

func TestPortableWorldIdentity(t *testing.T) {
	a, b := fixture(t), fixture(t)
	a.Bindings = []protocol.Binding{binding("Scalar", []protocol.Type{{Tag: "int64"}}, []protocol.Type{{Tag: "int64"}})}
	b.Bindings = a.Bindings
	first := Check(context.Background(), a)
	second := Check(context.Background(), b)
	if first.Bindings[0].Status != "verified" || second.Bindings[0].Status != "verified" || first.GoWorldIdentity == "" || first.GoWorldIdentity != second.GoWorldIdentity {
		t.Fatalf("relocated request changed world identity: %+v / %+v", first, second)
	}
}

func TestPersistedGOPATHInModuleOffMode(t *testing.T) {
	r := fixture(t)
	gopath := filepath.Join(t.TempDir(), "saved gopath")
	dependency := filepath.Join(gopath, "src", "example.com", "saved")
	if err := os.MkdirAll(dependency, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dependency, "saved.go"), []byte("package saved\nfunc Value() int64 { return 42 }\n"), 0644); err != nil {
		t.Fatal(err)
	}
	config := filepath.Join(t.TempDir(), "go-env")
	if err := os.WriteFile(config, []byte("GOPATH="+gopath+"\n"), 0644); err != nil {
		t.Fatal(err)
	}
	t.Setenv("GOENV", config)
	t.Setenv("GOPATH", "")
	r.BuildContext.GO111MODULE = "off"
	b := binding("Value", nil, []protocol.Type{{Tag: "int64"}})
	b.ImportPath = "example.com/saved"
	r.Bindings = []protocol.Binding{b}
	result := Check(context.Background(), r)
	if len(result.Bindings) != 1 || result.Bindings[0].Status != "verified" {
		t.Fatalf("saved GOPATH was ignored: %+v", result)
	}
}

func TestFileCallerIgnoresStaleGeneratedPackage(t *testing.T) {
	r := fixture(t)
	r.CallerContext.LoadMode = "files"
	r.CallerContext.ImportPath = "command-line-arguments"
	r.CallerContext.Package = "main"
	stale := filepath.Join(r.CallerContext.Directory, "goml_generated.go")
	contents := []byte("package main\nfunc stale() { missing_function() }\n")
	if err := os.WriteFile(stale, contents, 0644); err != nil {
		t.Fatal(err)
	}
	r.Bindings = []protocol.Binding{binding("Scalar", []protocol.Type{{Tag: "int64"}}, []protocol.Type{{Tag: "int64"}})}
	result := Check(context.Background(), r)
	if len(result.Bindings) != 1 || result.Bindings[0].Status != "verified" {
		t.Fatalf("file validation read stale output: %+v", result)
	}
	after, err := os.ReadFile(stale)
	if err != nil || string(after) != string(contents) {
		t.Fatalf("modified stale output: %s, %v", after, err)
	}
	r.CallerContext.LoadMode = "package"
	r.CallerContext.ImportPath = "example.com/host/gen"
	result = Check(context.Background(), r)
	if result.Bindings[0].Status != "failed" {
		t.Fatalf("package validation ignored package files: %+v", result)
	}
}

func TestCancellationDuringPackageLoad(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("requires a shell executable wrapper")
	}
	for _, deadline := range []bool{false, true} {
		name := "cancel"
		if deadline {
			name = "deadline"
		}
		t.Run(name, func(t *testing.T) {
			r := fixture(t)
			r.Bindings = []protocol.Binding{binding("Empty", nil, nil)}
			dir := t.TempDir()
			marker := filepath.Join(dir, "loading")
			quote := func(value string) string {
				return "'" + strings.ReplaceAll(value, "'", "'\"'\"'") + "'"
			}
			wrapper := filepath.Join(dir, "go")
			script := "#!/bin/sh\nif [ \"$1\" = list ]; then\n: > " + quote(marker) + "\nexec sleep 30\nfi\nexec " + quote(r.BuildContext.GoExecutable) + " \"$@\"\n"
			if err := os.WriteFile(wrapper, []byte(script), 0755); err != nil {
				t.Fatal(err)
			}
			t.Setenv("PATH", dir+string(os.PathListSeparator)+os.Getenv("PATH"))
			r.BuildContext.GoExecutable = wrapper
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			responses := make(chan protocol.Response, 1)
			go func() { responses <- Check(ctx, r) }()
			ticker := time.NewTicker(10 * time.Millisecond)
			defer ticker.Stop()
			loading := false
			for !loading {
				select {
				case response := <-responses:
					t.Fatalf("validation finished before package loading: %+v", response)
				case <-ctx.Done():
					t.Fatal("package loading did not start before deadline")
				case <-ticker.C:
					_, err := os.Stat(marker)
					loading = err == nil
				}
			}
			if !deadline {
				cancel()
			}
			select {
			case response := <-responses:
				if len(response.Diagnostics) != 1 || response.Diagnostics[0].Code != "ffi-canceled" || response.Diagnostics[0].Message != ctx.Err().Error() {
					t.Fatalf("cancellation diagnostic: %+v", response)
				}
				if len(response.Bindings) != 1 || response.Bindings[0].Status != "failed" {
					t.Fatalf("canceled binding: %+v", response)
				}
			case <-time.After(8 * time.Second):
				t.Fatal("package loading did not stop after cancellation")
			}
		})
	}
}

func TestMapKeysUseGoComparability(t *testing.T) {
	for name, key := range map[string]protocol.Type{
		"slice": {Tag: "slice", Element: &protocol.Type{Tag: "uint8"}},
		"map":   {Tag: "map", Key: &protocol.Type{Tag: "string"}, Element: &protocol.Type{Tag: "int64"}},
	} {
		t.Run(name, func(t *testing.T) {
			r := fixture(t)
			ty := protocol.Type{Tag: "map", Key: &key, Element: &protocol.Type{Tag: "int64"}}
			r.Bindings = []protocol.Binding{binding("Identity", []protocol.Type{ty}, []protocol.Type{ty})}
			if err := r.Validate(); err != nil {
				t.Fatal(err)
			}
			result := Check(context.Background(), r)
			if len(result.Bindings) != 1 || result.Bindings[0].Status != "failed" {
				t.Fatalf("non-comparable map key accepted: %+v", result)
			}
		})
	}
}
