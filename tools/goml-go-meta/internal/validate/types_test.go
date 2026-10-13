package validate

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"

	"goml.run/tools/goml-go-meta/internal/protocol"
)

func TestTypeGraphRecursiveGenericDeclarations(t *testing.T) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "types.go", `package example
type Node[T interface{}] struct {
	Next *Node[T]
	Value T
}
func (n Node[T]) Get() T { return n.Value }
type Link[T interface{}] = *Node[T]
`, 0)
	if err != nil {
		t.Fatal(err)
	}
	pkg, err := new(types.Config).Check("example.com/types", fset, []*ast.File{file}, nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"Node", "Link"} {
		t.Run(name, func(t *testing.T) {
			declaration := pkg.Scope().Lookup(name).Type()
			instance, err := types.Instantiate(nil, declaration, []types.Type{types.Typ[types.Int]}, true)
			if err != nil {
				t.Fatal(err)
			}
			graph := typeGraph{ids: map[types.Type]string{}}
			id := graph.add(instance)
			if graph.err != nil {
				t.Fatal(graph.err)
			}
			nodes := map[string]protocol.ReferencedType{}
			for _, node := range graph.nodes {
				nodes[node.ID] = node
			}
			root := nodes[id]
			if root.Name != name || root.ImportPath != pkg.Path() || root.ID != "t0" || root.Underlying != "t1" {
				t.Fatalf("instantiated declaration identity and traversal: %+v", root)
			}
			if len(root.TypeParameters) != 1 || len(root.TypeArguments) != 1 {
				t.Fatalf("generic metadata: %+v", root)
			}
			if parameter := nodes[root.TypeParameters[0]]; parameter.Tag != "type_parameter" || parameter.Name != "T" {
				t.Fatalf("declaration parameter: %+v", parameter)
			}
			argumentID := root.TypeArguments[0]
			if argument := nodes[argumentID]; argument.Tag != "basic" || argument.Name != "int" {
				t.Fatalf("instance argument: %+v", argument)
			}
			origin := nodes[root.Origin]
			if root.Origin != graph.add(declaration) || origin.Name != name || origin.Origin != "" || origin.TypeArguments == nil || len(origin.TypeArguments) != 0 {
				t.Fatalf("declaration origin: %+v", origin)
			}
			named := root
			if name == "Link" {
				pointer := nodes[root.Underlying]
				if root.Tag != "alias" || pointer.Tag != "pointer" {
					t.Fatalf("alias target: %+v, %+v", root, pointer)
				}
				named = nodes[pointer.Element]
			}
			if named.Tag != "named" || named.Name != "Node" || nodes[named.Origin].Name != "Node" {
				t.Fatalf("recursive named instance: %+v", named)
			}
			structure := nodes[named.Underlying]
			if structure.Tag != "struct" || len(structure.Fields) != 2 {
				t.Fatalf("recursive structure: %+v", structure)
			}
			pointer := nodes[structure.Fields[0].Type]
			if pointer.Tag != "pointer" || pointer.Element != named.ID || structure.Fields[1].Type != argumentID {
				t.Fatalf("recursive and substituted field types: %+v, %+v", structure, pointer)
			}
			if len(named.Methods) != 1 || named.Methods[0].Name != "Get" {
				t.Fatalf("instantiated method set: %+v", named)
			}
			method := nodes[named.Methods[0].Type]
			if method.Receiver != named.ID || len(method.Results) != 1 || method.Results[0] != argumentID {
				t.Fatalf("instantiated method signature: %+v", method)
			}
			count := len(graph.nodes)
			if graph.add(instance) != id || len(graph.nodes) != count || len(nodes) != count {
				t.Fatal("revisiting declarations changed graph identity")
			}
		})
	}
}

func TestTypeGraphNamedPackageIdentity(t *testing.T) {
	graph := typeGraph{ids: map[types.Type]string{}}
	for _, path := range []string{"example.com/first", "example.com/second"} {
		pkg := types.NewPackage(path, "example")
		named := types.NewNamed(types.NewTypeName(token.NoPos, pkg, "Number", nil), types.Typ[types.Int], nil)
		id := graph.add(named)
		var node protocol.ReferencedType
		for _, candidate := range graph.nodes {
			if candidate.ID == id {
				node = candidate
			}
		}
		if node.ImportPath != path || node.Name != "Number" || node.TypeParameters == nil || node.TypeArguments == nil || node.Methods != nil {
			t.Fatalf("package identity and empty metadata: %+v", node)
		}
	}
	if len(graph.nodes) != 3 {
		t.Fatalf("distinct named identities with shared underlying type: %+v", graph.nodes)
	}
	graph.add(types.Universe.Lookup("error").Type())
	if graph.err != nil {
		t.Fatal(graph.err)
	}
	if universe := graph.nodes[3]; universe.Name != "error" || universe.ImportPath != "" {
		t.Fatalf("universe object identity: %+v", universe)
	}
}
