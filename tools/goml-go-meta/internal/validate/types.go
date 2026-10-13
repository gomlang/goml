package validate

import (
	"fmt"
	"go/types"

	"goml.run/tools/goml-go-meta/internal/protocol"
)

type typeGraph struct {
	ids   map[types.Type]string
	nodes []protocol.ReferencedType
	err   error
}

type declaredType interface {
	types.Type
	Obj() *types.TypeName
	TypeParams() *types.TypeParamList
	TypeArgs() *types.TypeList
}

func (g *typeGraph) declaration(node *protocol.ReferencedType, ty declaredType, origin types.Type) {
	object := ty.Obj()
	node.Name = object.Name()
	if object.Pkg() != nil {
		node.ImportPath = object.Pkg().Path()
	}
	node.TypeParameters = g.typeParams(ty.TypeParams())
	node.TypeArguments = g.typeArgs(ty.TypeArgs())
	if origin != ty {
		node.Origin = g.add(origin)
	}
}

func (g *typeGraph) field(object types.Object) protocol.TypeField {
	field := protocol.TypeField{Name: object.Name(), Type: g.add(object.Type())}
	if object.Pkg() != nil {
		field.ImportPath = object.Pkg().Path()
	}
	return field
}

func (g *typeGraph) tuple(tuple *types.Tuple) []string {
	result := make([]string, tuple.Len())
	for i := 0; i < tuple.Len(); i++ {
		result[i] = g.add(tuple.At(i).Type())
	}
	return result
}

func (g *typeGraph) typeParams(params *types.TypeParamList) []string {
	result := make([]string, params.Len())
	for i := 0; i < params.Len(); i++ {
		result[i] = g.add(params.At(i))
	}
	return result
}

func (g *typeGraph) typeArgs(args *types.TypeList) []string {
	result := make([]string, args.Len())
	for i := 0; i < args.Len(); i++ {
		result[i] = g.add(args.At(i))
	}
	return result
}

func (g *typeGraph) add(ty types.Type) string {
	if g.err != nil {
		return ""
	}
	if id, ok := g.ids[ty]; ok {
		return id
	}
	if len(g.nodes) >= protocol.MaxTypeNodes {
		g.err = fmt.Errorf("referenced type node limit exceeded")
		return ""
	}
	index := len(g.nodes)
	node := protocol.ReferencedType{ID: fmt.Sprintf("t%d", index)}
	g.ids[ty] = node.ID
	g.nodes = append(g.nodes, node)
	switch t := ty.(type) {
	case *types.Basic:
		node.Tag = "basic"
		node.Name = t.Name()
	case *types.Pointer:
		node.Tag = "pointer"
		node.Element = g.add(t.Elem())
	case *types.Array:
		node.Tag = "array"
		length := t.Len()
		node.Length = &length
		node.Element = g.add(t.Elem())
	case *types.Slice:
		node.Tag = "slice"
		node.Element = g.add(t.Elem())
	case *types.Chan:
		node.Tag = "channel"
		node.Direction = "both"
		if t.Dir() == types.SendOnly {
			node.Direction = "send"
		}
		if t.Dir() == types.RecvOnly {
			node.Direction = "receive"
		}
		node.Element = g.add(t.Elem())
	case *types.Map:
		node.Tag = "map"
		node.Key = g.add(t.Key())
		node.Element = g.add(t.Elem())
	case *types.Named:
		node.Tag = "named"
		node.Underlying = g.add(t.Underlying())
		g.declaration(&node, t, t.Origin())
		for i := 0; i < t.NumMethods(); i++ {
			node.Methods = append(node.Methods, g.field(t.Method(i)))
		}
	case *types.Alias:
		node.Tag = "alias"
		node.Underlying = g.add(t.Rhs())
		g.declaration(&node, t, t.Origin())
	case *types.Struct:
		node.Tag = "struct"
		for i := 0; i < t.NumFields(); i++ {
			field := g.field(t.Field(i))
			field.Embedded = t.Field(i).Embedded()
			field.Tag = t.Tag(i)
			node.Fields = append(node.Fields, field)
		}
	case *types.Interface:
		node.Tag = "interface"
		t.Complete()
		runtimeInterface := t.IsMethodSet()
		node.RuntimeInterface = &runtimeInterface
		methods := make([]protocol.TypeField, t.NumMethods())
		for i := 0; i < t.NumMethods(); i++ {
			methods[i] = g.field(t.Method(i))
		}
		node.InterfaceMethods = &methods
		for i := 0; i < t.NumExplicitMethods(); i++ {
			node.Methods = append(node.Methods, g.field(t.ExplicitMethod(i)))
		}
		for i := 0; i < t.NumEmbeddeds(); i++ {
			node.Embedded = append(node.Embedded, g.add(t.EmbeddedType(i)))
		}
	case *types.Signature:
		node.Tag = "signature"
		node.Parameters = g.tuple(t.Params())
		node.Results = g.tuple(t.Results())
		node.TypeParameters = g.typeParams(t.TypeParams())
		if t.Recv() != nil {
			node.Receiver = g.add(t.Recv().Type())
		}
		node.Variadic = t.Variadic()
	case *types.Tuple:
		node.Tag = "tuple"
		node.Parameters = g.tuple(t)
	case *types.TypeParam:
		node.Tag = "type_parameter"
		node.Name = t.Obj().Name()
		node.Constraint = g.add(t.Constraint())
	case *types.Union:
		node.Tag = "union"
		for i := 0; i < t.Len(); i++ {
			node.Terms = append(node.Terms, protocol.TypeTerm{Type: g.add(t.Term(i).Type()), Tilde: t.Term(i).Tilde()})
		}
	default:
		g.err = fmt.Errorf("unsupported Go metadata type %T", ty)
	}
	g.nodes[index] = node
	return node.ID
}
