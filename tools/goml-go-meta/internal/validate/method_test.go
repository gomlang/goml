package validate

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"goml.dev/tools/goml-go-meta/internal/protocol"
)

func methodBinding(id, name string, parameters, results []protocol.Type) protocol.Binding {
	return protocol.Binding{ID: id, Symbol: name, Parameters: parameters, Results: results, CallMode: "method"}
}

func TestMethodBindingsUseActualReceiverMethodSets(t *testing.T) {
	r := fixture(t)
	source := `package shim
type Cell struct { Value int64 }
func (c Cell) Read() int64 { return c.Value }
func (c *Cell) Write(value int64) { c.Value = value }
func (c *Cell) IsNil() bool { return c == nil }
func (c Cell) hidden() int64 { return c.Value }
type Handle = *Cell
type Reader interface { Read() int64 }
type Embedded struct { *Cell }
type Box[T any] struct { Value T }
func (b Box[T]) Get() T { return b.Value }
`
	if err := os.WriteFile(filepath.Join(r.BuildContext.ModuleDir, "shim", "methods.go"), []byte(source), 0644); err != nil {
		t.Fatal(err)
	}
	named := func(name string) protocol.Type {
		return protocol.Type{Tag: "named", ImportPath: "example.com/host/shim", Name: name}
	}
	cell := named("Cell")
	pointer := protocol.Type{Tag: "pointer", Element: &cell}
	integer := protocol.Type{Tag: "int64"}
	box := named("Box")
	box.Arguments = []protocol.Type{integer}
	r.Bindings = []protocol.Binding{
		methodBinding("value", "Read", []protocol.Type{cell}, []protocol.Type{integer}),
		methodBinding("pointer", "Read", []protocol.Type{pointer}, []protocol.Type{integer}),
		methodBinding("write", "Write", []protocol.Type{pointer, integer}, nil),
		methodBinding("nil", "IsNil", []protocol.Type{pointer}, []protocol.Type{{Tag: "bool"}}),
		methodBinding("alias", "Read", []protocol.Type{named("Handle")}, []protocol.Type{integer}),
		methodBinding("interface", "Read", []protocol.Type{named("Reader")}, []protocol.Type{integer}),
		methodBinding("promoted", "Write", []protocol.Type{named("Embedded"), integer}, nil),
		methodBinding("generic", "Get", []protocol.Type{box}, []protocol.Type{integer}),
		methodBinding("duplicate", "Read", []protocol.Type{pointer}, []protocol.Type{integer}),
	}
	result := Check(context.Background(), r)
	if len(result.Diagnostics) != 0 || len(result.Bindings) != len(r.Bindings) {
		t.Fatalf("result: %+v", result)
	}
	for _, b := range result.Bindings {
		if b.Status != "verified" || b.SignatureType == "" || b.ActualSignature == "" {
			t.Fatalf("binding: %+v", b)
		}
	}
	for _, invalid := range []protocol.Binding{
		methodBinding("value-pointer", "Write", []protocol.Type{cell, integer}, nil),
		methodBinding("missing", "Missing", []protocol.Type{pointer}, nil),
		methodBinding("private", "hidden", []protocol.Type{cell}, []protocol.Type{integer}),
		methodBinding("argument", "Write", []protocol.Type{pointer, {Tag: "bool"}}, nil),
		methodBinding("return", "Read", []protocol.Type{pointer}, []protocol.Type{{Tag: "bool"}}),
		methodBinding("depth", "Read", []protocol.Type{{Tag: "pointer", Element: &pointer}}, []protocol.Type{integer}),
		methodBinding("primitive", "Read", []protocol.Type{integer}, []protocol.Type{integer}),
	} {
		t.Run(invalid.ID, func(t *testing.T) {
			r.Bindings = []protocol.Binding{invalid}
			result := Check(context.Background(), r)
			if len(result.Bindings) != 1 || result.Bindings[0].Status == "verified" {
				t.Fatalf("accepted mismatch: %+v", result)
			}
		})
	}
}

func TestMethodWitnessUsesExplicitReceiverWithoutAddressTaking(t *testing.T) {
	element := protocol.Type{Tag: "named", ImportPath: "example.com/host/shim", Name: "Cell"}
	receiver := protocol.Type{Tag: "pointer", Element: &element}
	data, err := witness("caller", methodBinding("nil", "IsNil", []protocol.Type{receiver}, []protocol.Type{{Tag: "bool"}}), 0)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), "(*goml_type_import_0.Cell).IsNil(arg0)") || strings.Contains(string(data), "&arg0") {
		t.Fatalf("witness: %s", data)
	}
}

func TestMethodWitnessPreservesNilReceiverBehavior(t *testing.T) {
	r := fixture(t)
	if err := os.WriteFile(filepath.Join(r.BuildContext.ModuleDir, "shim", "nil.go"), []byte("package shim\ntype Cell struct{}\nfunc (c *Cell) IsNil() bool { return c == nil }\n"), 0644); err != nil {
		t.Fatal(err)
	}
	element := protocol.Type{Tag: "named", ImportPath: "example.com/host/shim", Name: "Cell"}
	receiver := protocol.Type{Tag: "pointer", Element: &element}
	data, err := witness("gen", methodBinding("nil", "IsNil", []protocol.Type{receiver}, []protocol.Type{{Tag: "bool"}}), 0)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(r.CallerContext.Directory, "witness.go"), data, 0644); err != nil {
		t.Fatal(err)
	}
	test := "package gen\nimport \"testing\"\nfunc TestNil(t *testing.T) { if !goml_ffi_witness_0(nil) { t.Fatal(\"nil receiver was replaced\") } }\n"
	if err := os.WriteFile(filepath.Join(r.CallerContext.Directory, "witness_test.go"), []byte(test), 0644); err != nil {
		t.Fatal(err)
	}
	command := exec.Command(r.BuildContext.GoExecutable, "test", "-mod=readonly", "./gen")
	command.Dir = r.BuildContext.ModuleDir
	command.Env = environment(r.BuildContext)
	if output, err := command.CombinedOutput(); err != nil {
		t.Fatalf("nil receiver test: %v\n%s", err, output)
	}
}
