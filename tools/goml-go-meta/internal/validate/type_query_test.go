package validate

import (
	"context"
	"goml.dev/tools/goml-go-meta/internal/protocol"
	"os"
	"path/filepath"
	"testing"
)

func TestTypeQueries(t *testing.T) {
	r := fixture(t)
	source := `package shim
import "time"
type DurationAlias = time.Duration
type IntegerAlias = int64
type Reader interface { Read([]byte) (int, error) }
type Box[T ~int64] struct { Value T }
func (b Box[T]) Get() T { return b.Value }
`
	if err := os.WriteFile(filepath.Join(r.BuildContext.ModuleDir, "shim", "types.go"), []byte(source), 0644); err != nil {
		t.Fatal(err)
	}
	r.TypeQueries = []protocol.TypeQuery{
		{ID: "duration", ImportPath: "time", Name: "Duration"},
		{ID: "same", ImportPath: "time", Name: "Duration"},
		{ID: "alias", ImportPath: "example.com/host/shim", Name: "DurationAlias"},
		{ID: "integer", ImportPath: "example.com/host/shim", Name: "IntegerAlias"},
		{ID: "interface", ImportPath: "example.com/host/shim", Name: "Reader"},
		{ID: "box", ImportPath: "example.com/host/shim", Name: "Box", Arguments: []protocol.Type{{Tag: "int64"}}},
	}
	result := Check(context.Background(), r)
	if len(result.TypeResults) != len(r.TypeQueries) || len(result.Diagnostics) != 0 {
		t.Fatalf("results: %+v", result)
	}
	nodes := map[string]protocol.ReferencedType{}
	for _, node := range result.ReferencedTypes {
		nodes[node.ID] = node
	}
	for _, value := range result.TypeResults {
		if value.Status != "verified" {
			t.Fatalf("query: %+v", value)
		}
	}
	duration := result.TypeResults[0].Type
	if duration != result.TypeResults[1].Type || duration != result.TypeResults[2].Type {
		t.Fatal("canonical identity")
	}
	if nodes[duration].Tag != "named" || nodes[duration].ImportPath != "time" {
		t.Fatal("nominal type")
	}
	if result.TypeResults[2].Kind != "alias" || nodes[result.TypeResults[3].Type].Name != "int64" {
		t.Fatal("alias normalization")
	}
	reader := nodes[result.TypeResults[4].Type]
	if nodes[reader.Underlying].Tag != "interface" {
		t.Fatal("interface kind")
	}
	box := nodes[result.TypeResults[5].Type]
	if len(box.TypeArguments) != 1 || len(box.Methods) != 1 {
		t.Fatalf("generic method set: %+v", box)
	}
	r.Bindings = []protocol.Binding{binding("Scalar", []protocol.Type{{Tag: "int64"}}, []protocol.Type{{Tag: "int64"}})}
	r.TypeQueries = append(r.TypeQueries, protocol.TypeQuery{ID: "bad", ImportPath: "example.com/host/shim", Name: "Box", Arguments: []protocol.Type{{Tag: "string"}}})
	result = Check(context.Background(), r)
	if result.TypeResults[6].Status != "failed" || result.TypeResults[0].Status != "verified" || result.Bindings[0].Status != "verified" {
		t.Fatalf("mixed failure: %+v", result)
	}
}

func TestInvalidTypeQueries(t *testing.T) {
	r := fixture(t)
	r.TypeQueries = []protocol.TypeQuery{
		{ID: "missing", ImportPath: "example.com/host/shim", Name: "Missing"},
		{ID: "private", ImportPath: "example.com/host/shim", Name: "private"},
		{ID: "function", ImportPath: "example.com/host/shim", Name: "Scalar"},
		{ID: "args", ImportPath: "time", Name: "Duration", Arguments: []protocol.Type{{Tag: "int64"}}},
		{ID: "valid", ImportPath: "time", Name: "Duration"},
	}
	result := Check(context.Background(), r)
	if len(result.TypeResults) != len(r.TypeQueries) {
		t.Fatalf("missing results: %+v", result)
	}
	for i, value := range result.TypeResults {
		if i == 4 {
			if value.Status != "verified" {
				t.Fatalf("unrelated type failed: %+v", value)
			}
			continue
		}
		if value.Status != "failed" || len(value.Diagnostics) == 0 {
			t.Fatalf("accepted invalid type: %+v", value)
		}
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	result = Check(ctx, r)
	if len(result.TypeResults) != len(r.TypeQueries) || result.TypeResults[4].Status != "failed" {
		t.Fatal("canceled query results")
	}
	r.CallerContext.LoadMode = "files"
	r.CallerContext.ImportPath = "command-line-arguments"
	r.TypeQueries = r.TypeQueries[4:]
	result = Check(context.Background(), r)
	if len(result.TypeResults) != 1 || result.TypeResults[0].Status != "verified" {
		t.Fatalf("file caller: %+v", result)
	}
}

func TestNamedTypeArgumentsAndBindings(t *testing.T) {
	r := fixture(t)
	if err := os.WriteFile(filepath.Join(r.BuildContext.ModuleDir, "shim", "generic.go"), []byte("package shim\ntype Box[T ~int64] struct { Value T }\ntype Container[T any] struct { Value T }\n"), 0644); err != nil {
		t.Fatal(err)
	}
	duration := protocol.Type{Tag: "named", ImportPath: "time", Name: "Duration"}
	box := protocol.Type{Tag: "named", ImportPath: "example.com/host/shim", Name: "Box", Arguments: []protocol.Type{duration}}
	r.TypeQueries = []protocol.TypeQuery{
		{ID: "box", ImportPath: "example.com/host/shim", Name: "Box", Arguments: []protocol.Type{duration}},
		{ID: "nested", ImportPath: "example.com/host/shim", Name: "Container", Arguments: []protocol.Type{{Tag: "slice", Element: &box}}},
	}
	r.Bindings = []protocol.Binding{{ID: "sleep", ImportPath: "time", Symbol: "Sleep", Parameters: []protocol.Type{duration}, CallMode: "ordinary"}}
	result := Check(context.Background(), r)
	if len(result.TypeResults) != 2 || result.TypeResults[0].Status != "verified" || result.TypeResults[1].Status != "verified" || result.Bindings[0].Status != "verified" {
		t.Fatalf("named arguments: %+v", result)
	}
	nodes := map[string]protocol.ReferencedType{}
	for _, node := range result.ReferencedTypes {
		nodes[node.ID] = node
	}
	instance := nodes[result.TypeResults[0].Type]
	argument := nodes[instance.TypeArguments[0]]
	if argument.ImportPath != "time" || argument.Name != "Duration" {
		t.Fatalf("lost named argument: %+v", argument)
	}
	outer := nodes[result.TypeResults[1].Type]
	slice := nodes[outer.TypeArguments[0]]
	nested := nodes[slice.Element]
	if slice.Tag != "slice" || nested.Name != "Box" || nodes[nested.TypeArguments[0]].Name != "Duration" {
		t.Fatal("lost nested identity")
	}
	r.Bindings[0].Parameters = []protocol.Type{{Tag: "int64"}}
	result = Check(context.Background(), r)
	if result.Bindings[0].Status != "failed" || result.TypeResults[0].Status != "verified" {
		t.Fatalf("erased nominal distinction: %+v", result)
	}
}

func TestGenericDeclarationQuery(t *testing.T) {
	r := fixture(t)
	if err := os.WriteFile(filepath.Join(r.BuildContext.ModuleDir, "shim", "declaration.go"), []byte("package shim\ntype Box[T ~int64] struct { Value T }\n"), 0644); err != nil {
		t.Fatal(err)
	}
	r.TypeQueries = []protocol.TypeQuery{{ID: "declaration", ImportPath: "example.com/host/shim", Name: "Box", Mode: "declaration"}, {ID: "value", ImportPath: "example.com/host/shim", Name: "Box"}}
	result := Check(context.Background(), r)
	if len(result.TypeResults) != 2 || result.TypeResults[0].Status != "verified" || result.TypeResults[1].Status != "failed" {
		t.Fatalf("declaration vs instance: %+v", result)
	}
	if result.TypeResults[0].Type != "" || result.TypeResults[0].DeclaredType == "" {
		t.Fatal("declaration was treated as value type")
	}
	for _, node := range result.ReferencedTypes {
		if node.ID == result.TypeResults[0].DeclaredType {
			if len(node.TypeParameters) != 1 {
				t.Fatal("missing generic parameters")
			}
			return
		}
	}
	t.Fatal("missing declaration node")
}
