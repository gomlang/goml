package validate

import (
	"fmt"
	"go/types"

	"golang.org/x/tools/go/packages"
	"goml.run/tools/goml-go-meta/internal/protocol"
)

func (r *checkResults) resolveBindings(bindings []protocol.Binding, loaded map[string]*packages.Package, methods map[int]*types.Func, graph *typeGraph) {
	for index, binding := range bindings {
		var object types.Object
		if binding.CallMode == "method" {
			method := methods[index]
			if method == nil || !method.Exported() {
				r.failBinding(index, "ffi-symbol-not-found", "Go receiver has no exported method: "+binding.Symbol)
				continue
			}
			object = method
		} else {
			p := loaded[binding.ImportPath]
			if p == nil || p.Types == nil {
				r.failBinding(index, "ffi-package-load", "could not load "+binding.ImportPath)
				continue
			}
			if p.Name == "main" {
				r.failBinding(index, "ffi-package-load", "cannot import package main")
				continue
			}
			object = p.Types.Scope().Lookup(binding.Symbol)
			if object == nil {
				r.failBinding(index, "ffi-symbol-not-found", "Go symbol not found: "+binding.ImportPath+"."+binding.Symbol)
				continue
			}
		}
		result := &r.Bindings[index]
		result.SignatureType = graph.add(object.Type())
		result.ActualSignature = types.TypeString(object.Type(), func(p *types.Package) string { return p.Path() })
		if !object.Exported() {
			r.failBinding(index, "ffi-symbol-not-found", "Go symbol is not exported: "+binding.Symbol)
		}
		if _, ok := object.(*types.Func); !ok {
			r.failBinding(index, "ffi-signature-mismatch", "Go symbol is not a package-level function: "+binding.Symbol)
		}
	}
}

func (r *checkResults) resolveQueries(queries []protocol.TypeQuery, loaded map[string]*packages.Package, root *packages.Package, graph *typeGraph) {
	for index, query := range queries {
		p := loaded[query.ImportPath]
		if p == nil || p.Types == nil || p.Name == "main" {
			r.failQuery(index, "ffi-package-load", "could not import "+query.ImportPath)
			continue
		}
		object := p.Types.Scope().Lookup(query.Name)
		if object == nil || !object.Exported() {
			r.failQuery(index, "ffi-symbol-not-found", "Go type is missing or not exported: "+query.Name)
			continue
		}
		result := &r.TypeResults[index]
		witnessKind := "type"
		if query.Mode == "function" {
			if _, ok := object.(*types.Func); !ok {
				r.failQuery(index, "ffi-signature-mismatch", "Go object is not a package-level function: "+query.Name)
				continue
			}
			result.Kind = "function"
			witnessKind = "function"
		} else {
			declared, ok := object.(*types.TypeName)
			if !ok {
				r.failQuery(index, "ffi-signature-mismatch", "Go object is not a type: "+query.Name)
				continue
			}
			result.Kind = "defined"
			if declared.IsAlias() {
				result.Kind = "alias"
			}
		}
		result.DeclaredType = graph.add(object.Type())
		if query.Mode == "declaration" || result.Status != "verified" {
			continue
		}
		resolved := root.Types.Scope().Lookup(fmt.Sprintf("goml_type_query_%d", index))
		if resolved == nil {
			r.failQuery(index, "ffi-package-load", "missing resolved "+witnessKind+" witness")
			continue
		}
		if query.Mode == "function" {
			signature, ok := resolved.Type().(*types.Signature)
			if !ok || signature.TypeParams().Len() != 0 {
				r.failQuery(index, "ffi-signature-mismatch", "function query requires a concrete signature")
				continue
			}
			result.Type = graph.add(signature)
		} else {
			result.Type = graph.add(types.Unalias(resolved.Type()))
		}
	}
}

func selectedWitnessMethods(root *packages.Package, witnesses *witnessSet) map[int]*types.Func {
	result := map[int]*types.Func{}
	if root == nil || root.TypesInfo == nil || root.Fset == nil {
		return result
	}
	for expression, selection := range root.TypesInfo.Selections {
		if selection.Kind() != types.MethodExpr {
			continue
		}
		method, ok := selection.Obj().(*types.Func)
		if !ok {
			continue
		}
		position := root.Fset.Position(expression.Pos())
		for _, file := range witnesses.origins.files(position.Filename, position.Line) {
			for _, index := range witnesses.targets[file].bindings {
				result[index] = method
			}
		}
	}
	return result
}
