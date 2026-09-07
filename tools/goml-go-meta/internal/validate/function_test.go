package validate

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"goml.dev/tools/goml-go-meta/internal/protocol"
)

func callbackType(parameters, results []protocol.Type) protocol.Type {
	return protocol.Type{Tag: "function", Signature: &protocol.FunctionSignature{Parameters: parameters, Results: results}}
}

func TestFunctionBindingsValidateExactSignatures(t *testing.T) {
	r := fixture(t)
	source := `package shim
import "time"
func Use(value func(int64) (string, error)) { _, _ = value(1) }
func Return() func(int64) (string, error) { return nil }
func Unit(value func()) func() { return value }
func Nested(value func(func()) func()) func(func()) func() { return value }
func Named(value func(time.Duration) *time.Timer) {}
func VariadicCallback(value func(...int64)) {}
`
	if err := os.WriteFile(filepath.Join(r.BuildContext.ModuleDir, "shim", "function.go"), []byte(source), 0644); err != nil {
		t.Fatal(err)
	}
	integer := protocol.Type{Tag: "int64"}
	text := protocol.Type{Tag: "string"}
	failure := protocol.Type{Tag: "named", Name: "error"}
	empty := callbackType([]protocol.Type{}, []protocol.Type{})
	callback := callbackType([]protocol.Type{integer}, []protocol.Type{text, failure})
	nested := callbackType([]protocol.Type{empty}, []protocol.Type{empty})
	duration := protocol.Type{Tag: "named", ImportPath: "time", Name: "Duration"}
	timer := protocol.Type{Tag: "named", ImportPath: "time", Name: "Timer"}
	named := callbackType([]protocol.Type{duration}, []protocol.Type{{Tag: "pointer", Element: &timer}})
	r.Bindings = []protocol.Binding{
		binding("Use", []protocol.Type{callback}, nil),
		binding("Return", nil, []protocol.Type{callback}),
		binding("Unit", []protocol.Type{empty}, []protocol.Type{empty}),
		binding("Nested", []protocol.Type{nested}, []protocol.Type{nested}),
		binding("Named", []protocol.Type{named}, nil),
	}
	result := Check(context.Background(), r)
	if len(result.Diagnostics) != 0 || len(result.Bindings) != len(r.Bindings) {
		t.Fatalf("result: %+v", result)
	}
	for _, value := range result.Bindings {
		if value.Status != "verified" {
			t.Fatalf("binding: %+v", value)
		}
	}
	for _, wrong := range []protocol.Type{
		empty,
		callbackType([]protocol.Type{text}, []protocol.Type{text, failure}),
		callbackType([]protocol.Type{integer}, []protocol.Type{failure, text}),
		callbackType([]protocol.Type{integer}, []protocol.Type{text}),
	} {
		r.Bindings = []protocol.Binding{binding("Use", []protocol.Type{wrong}, nil)}
		result = Check(context.Background(), r)
		if len(result.Bindings) != 1 || result.Bindings[0].Status == "verified" {
			t.Fatalf("accepted mismatch: %+v", result)
		}
	}
	r.Bindings = []protocol.Binding{binding("VariadicCallback", []protocol.Type{callbackType([]protocol.Type{{Tag: "slice", Element: &integer}}, []protocol.Type{})}, nil)}
	result = Check(context.Background(), r)
	if len(result.Bindings) != 1 || result.Bindings[0].Status == "verified" {
		t.Fatalf("accepted variadic mismatch: %+v", result)
	}
}
