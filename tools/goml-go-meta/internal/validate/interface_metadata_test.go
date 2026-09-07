package validate

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"goml.dev/tools/goml-go-meta/internal/protocol"
)

func TestInterfaceMetadataCompletedMethodSets(t *testing.T) {
	request := fixture(t)
	source := `package shim
type Base interface { Ping() }
type Left interface { Base }
type Right interface { Base }
type Diamond interface { Left; Right }
type Getter[T any] interface { Get() T }
type Sealed interface { private() }
type Constraint interface { ~int | ~string }
type Comparable interface { comparable }
type EmptyInterface interface {}
`
	if err := os.WriteFile(filepath.Join(request.BuildContext.ModuleDir, "shim", "interfaces.go"), []byte(source), 0644); err != nil {
		t.Fatal(err)
	}
	request.TypeQueries = []protocol.TypeQuery{
		{ID: "reader", ImportPath: "io", Name: "ReadCloser", Mode: "declaration"},
		{ID: "diamond", ImportPath: "example.com/host/shim", Name: "Diamond", Mode: "declaration"},
		{ID: "getter", ImportPath: "example.com/host/shim", Name: "Getter", Arguments: []protocol.Type{{Tag: "int64"}}},
		{ID: "sealed", ImportPath: "example.com/host/shim", Name: "Sealed", Mode: "declaration"},
		{ID: "constraint", ImportPath: "example.com/host/shim", Name: "Constraint", Mode: "declaration"},
		{ID: "comparable", ImportPath: "example.com/host/shim", Name: "Comparable", Mode: "declaration"},
		{ID: "empty", ImportPath: "example.com/host/shim", Name: "EmptyInterface", Mode: "declaration"},
	}
	response := Check(context.Background(), request)
	if len(response.Diagnostics) != 0 || len(response.TypeResults) != len(request.TypeQueries) {
		t.Fatalf("interface queries: %+v", response)
	}
	nodes := map[string]protocol.ReferencedType{}
	for _, node := range response.ReferencedTypes {
		nodes[node.ID] = node
	}
	interfaces := map[string]protocol.ReferencedType{}
	for _, result := range response.TypeResults {
		if result.Status != "verified" {
			t.Fatalf("interface declaration: %+v", result)
		}
		id := result.Type
		if id == "" {
			id = result.DeclaredType
		}
		node := nodes[id]
		for node.Tag == "named" || node.Tag == "alias" {
			node = nodes[node.Underlying]
		}
		if node.Tag != "interface" || node.RuntimeInterface == nil {
			t.Fatalf("interface classification missing: %+v", node)
		}
		expectedRuntime := result.ID != "constraint" && result.ID != "comparable"
		if *node.RuntimeInterface != expectedRuntime {
			t.Fatalf("runtime classification for %s: %+v", result.ID, node)
		}
		interfaces[result.ID] = node
	}
	complete := func(node protocol.ReferencedType) []protocol.TypeField {
		if node.InterfaceMethods == nil {
			t.Fatalf("missing complete interface method list: %+v", node)
		}
		return *node.InterfaceMethods
	}
	reader := interfaces["reader"]
	if len(reader.Methods) != 0 || len(reader.Embedded) != 2 || len(complete(reader)) != 2 {
		t.Fatalf("embedded io.ReadCloser methods: %+v", reader)
	}
	if complete(reader)[0].Name != "Close" || complete(reader)[1].Name != "Read" {
		t.Fatalf("complete deterministic method order: %+v", complete(reader))
	}
	read := nodes[complete(reader)[1].Type]
	if len(read.Parameters) != 1 || len(read.Results) != 2 || nodes[read.Parameters[0]].Tag != "slice" {
		t.Fatalf("Read signature: %+v", read)
	}
	if nodes[read.Results[0]].Name != "int" || nodes[read.Results[1]].Name != "error" {
		t.Fatalf("Read results: %+v", read.Results)
	}
	diamond := interfaces["diamond"]
	if len(complete(diamond)) != 1 || complete(diamond)[0].Name != "Ping" {
		t.Fatalf("duplicate inherited methods: %+v", diamond)
	}
	getter := interfaces["getter"]
	if len(complete(getter)) != 1 {
		t.Fatalf("generic method set: %+v", getter)
	}
	signature := nodes[complete(getter)[0].Type]
	if len(signature.Results) != 1 || nodes[signature.Results[0]].Name != "int64" {
		t.Fatalf("instantiated method result: %+v", signature)
	}
	sealed := interfaces["sealed"]
	if len(complete(sealed)) != 1 || complete(sealed)[0].Name != "private" || complete(sealed)[0].ImportPath != "example.com/host/shim" {
		t.Fatalf("unexported method identity: %+v", sealed)
	}
	if len(complete(interfaces["empty"])) != 0 {
		t.Fatal("empty runtime interface gained methods")
	}
}

func TestRuntimeTypeQueriesRejectConstraintInterfaces(t *testing.T) {
	request := fixture(t)
	source := `package shim
type Constraint interface { ~int }
type Comparable interface { comparable }
`
	if err := os.WriteFile(filepath.Join(request.BuildContext.ModuleDir, "shim", "constraints.go"), []byte(source), 0644); err != nil {
		t.Fatal(err)
	}
	request.TypeQueries = []protocol.TypeQuery{
		{ID: "declaration", ImportPath: "example.com/host/shim", Name: "Constraint", Mode: "declaration"},
		{ID: "runtime", ImportPath: "example.com/host/shim", Name: "Constraint"},
		{ID: "comparable", ImportPath: "example.com/host/shim", Name: "Comparable"},
		{ID: "reader", ImportPath: "io", Name: "Reader"},
	}
	response := Check(context.Background(), request)
	if len(response.TypeResults) != 4 {
		t.Fatalf("missing type results: %+v", response)
	}
	for i, result := range response.TypeResults {
		expected := "verified"
		if i == 1 || i == 2 {
			expected = "failed"
		}
		if result.Status != expected {
			t.Fatalf("runtime type query %s: %+v", result.ID, result)
		}
		if expected == "failed" && len(result.Diagnostics) == 0 {
			t.Fatalf("constraint rejection omitted diagnostics: %+v", result)
		}
	}
}
