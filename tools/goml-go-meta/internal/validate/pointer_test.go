package validate

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"goml.dev/tools/goml-go-meta/internal/protocol"
)

func TestPointerBindingsPreservePointeeIdentity(t *testing.T) {
	r := fixture(t)
	source := `package shim
type Node struct { Value int64 }
type Other struct { Value int64 }
func Pointer(value *Node) *Node { return value }
func Nested(value **Node) **Node { return value }
func NilPointer() *Node { return nil }
`
	if err := os.WriteFile(filepath.Join(r.BuildContext.ModuleDir, "shim", "pointer.go"), []byte(source), 0644); err != nil {
		t.Fatal(err)
	}
	node := protocol.Type{Tag: "named", ImportPath: "example.com/host/shim", Name: "Node"}
	pointer := protocol.Type{Tag: "pointer", Element: &node}
	nested := protocol.Type{Tag: "pointer", Element: &pointer}
	other := protocol.Type{Tag: "named", ImportPath: "example.com/host/shim", Name: "Other"}
	r.Bindings = []protocol.Binding{
		binding("Pointer", []protocol.Type{pointer}, []protocol.Type{pointer}),
		binding("Nested", []protocol.Type{nested}, []protocol.Type{nested}),
		binding("NilPointer", nil, []protocol.Type{pointer}),
	}
	result := Check(context.Background(), r)
	if len(result.Diagnostics) != 0 || len(result.Bindings) != 3 {
		t.Fatalf("result: %+v", result)
	}
	for _, b := range result.Bindings {
		if b.Status != "verified" {
			t.Fatalf("binding: %+v", b)
		}
	}
	for _, wrong := range []protocol.Type{node, {Tag: "pointer", Element: &other}, nested} {
		r.Bindings = []protocol.Binding{binding("Pointer", []protocol.Type{wrong}, []protocol.Type{pointer})}
		result = Check(context.Background(), r)
		if len(result.Bindings) != 1 || result.Bindings[0].Status == "verified" {
			t.Fatalf("accepted mismatch: %+v", result)
		}
	}
}
