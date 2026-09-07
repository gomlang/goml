package validate

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"goml.dev/tools/goml-go-meta/internal/protocol"
)

func TestFunctionQueriesRetainConcreteSignaturesWithoutInitialization(t *testing.T) {
	request := fixture(t)
	source := `package shim
func init() { panic("metadata queries must not execute init") }
func Number[T ~int64](value T) (T, error) { return value, nil }
`
	if err := os.WriteFile(filepath.Join(request.BuildContext.ModuleDir, "shim", "queries.go"), []byte(source), 0644); err != nil {
		t.Fatal(err)
	}
	request.TypeQueries = []protocol.TypeQuery{
		{Mode: "function", ID: "count", ImportPath: "strings", Name: "Count"},
		{Mode: "function", ID: "identity", ImportPath: "example.com/host/shim", Name: "Identity", Arguments: []protocol.Type{{Tag: "int64"}}},
		{Mode: "function", ID: "number", ImportPath: "example.com/host/shim", Name: "Number", Arguments: []protocol.Type{{Tag: "int64"}}},
		{Mode: "function", ID: "variadic", ImportPath: "example.com/host/shim", Name: "Variadic"},
	}
	result := Check(context.Background(), request)
	if len(result.Diagnostics) != 0 || len(result.TypeResults) != 4 {
		t.Fatalf("response: %+v", result)
	}
	nodes := map[string]protocol.ReferencedType{}
	for _, node := range result.ReferencedTypes {
		nodes[node.ID] = node
	}
	for _, item := range result.TypeResults {
		signature := nodes[item.Type]
		if item.Status != "verified" || item.Kind != "function" || signature.Tag != "signature" || len(signature.TypeParameters) != 0 {
			t.Fatalf("function: %+v signature: %+v", item, signature)
		}
	}
	count := nodes[result.TypeResults[0].Type]
	if len(count.Parameters) != 2 || nodes[count.Parameters[0]].Name != "string" || nodes[count.Results[0]].Name != "int" {
		t.Fatalf("count: %+v", count)
	}
	identity := nodes[result.TypeResults[1].Type]
	if nodes[identity.Parameters[0]].Name != "int64" || identity.Parameters[0] != identity.Results[0] {
		t.Fatalf("identity: %+v", identity)
	}
	generic := nodes[result.TypeResults[1].DeclaredType]
	if len(generic.TypeParameters) != 1 {
		t.Fatalf("generic: %+v", generic)
	}
	number := nodes[result.TypeResults[2].Type]
	if len(number.Results) != 2 || nodes[number.Results[1]].Name != "error" {
		t.Fatalf("number: %+v", number)
	}
	if !nodes[result.TypeResults[3].Type].Variadic {
		t.Fatal("lost variadic signature")
	}
	repeated := Check(context.Background(), request)
	if repeated.GoWorldIdentity != result.GoWorldIdentity {
		t.Fatal("nondeterministic query world")
	}
}

func TestFunctionQueriesRejectMissingInstancesAndNonFunctions(t *testing.T) {
	request := fixture(t)
	source := "package shim\nfunc Number[T ~int64](value T) T { return value }\n"
	if err := os.WriteFile(filepath.Join(request.BuildContext.ModuleDir, "shim", "queries.go"), []byte(source), 0644); err != nil {
		t.Fatal(err)
	}
	request.TypeQueries = []protocol.TypeQuery{
		{Mode: "function", ID: "missing", ImportPath: "example.com/host/shim", Name: "Missing"},
		{Mode: "function", ID: "private", ImportPath: "example.com/host/shim", Name: "private"},
		{Mode: "function", ID: "variable", ImportPath: "example.com/host/shim", Name: "FunctionValue"},
		{Mode: "function", ID: "type", ImportPath: "example.com/host/shim", Name: "Defined"},
		{Mode: "function", ID: "generic", ImportPath: "example.com/host/shim", Name: "Identity"},
		{Mode: "function", ID: "constraint", ImportPath: "example.com/host/shim", Name: "Number", Arguments: []protocol.Type{{Tag: "string"}}},
		{Mode: "function", ID: "valid", ImportPath: "example.com/host/shim", Name: "Scalar"},
	}
	result := Check(context.Background(), request)
	if len(result.TypeResults) != len(request.TypeQueries) {
		t.Fatalf("response: %+v", result)
	}
	for index, item := range result.TypeResults {
		if index == len(request.TypeQueries)-1 {
			if item.Status != "verified" {
				t.Fatalf("unrelated function failed: %+v", item)
			}
		} else if item.Status != "failed" || len(item.Diagnostics) == 0 {
			t.Fatalf("invalid query accepted: %+v", item)
		}
	}
}
