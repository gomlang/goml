package validate

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"sort"
	"strconv"

	"goml.dev/tools/goml-go-meta/internal/protocol"
)

type witnessTypes struct{ aliases map[string]string }

func (w *witnessTypes) named(path, name string, arguments []protocol.Type) ast.Expr {
	if path == "" && name == "error" {
		return ast.NewIdent("error")
	}
	alias, found := w.aliases[path]
	if !found {
		alias = fmt.Sprintf("goml_type_import_%d", len(w.aliases))
		w.aliases[path] = alias
	}
	var result ast.Expr = &ast.SelectorExpr{X: ast.NewIdent(alias), Sel: ast.NewIdent(name)}
	if len(arguments) > 0 {
		indices := []ast.Expr{}
		for _, argument := range arguments {
			indices = append(indices, w.expression(argument))
		}
		result = &ast.IndexListExpr{X: result, Indices: indices}
	}
	return result
}

func (w *witnessTypes) imports() *ast.GenDecl {
	paths := []string{}
	for path := range w.aliases {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	specs := []ast.Spec{}
	for _, path := range paths {
		specs = append(specs, &ast.ImportSpec{Name: ast.NewIdent(w.aliases[path]), Path: &ast.BasicLit{Kind: token.STRING, Value: strconv.Quote(path)}})
	}
	return &ast.GenDecl{Tok: token.IMPORT, Specs: specs}
}

func (w *witnessTypes) expression(t protocol.Type) ast.Expr {
	switch t.Tag {
	case "named":
		return w.named(t.ImportPath, t.Name, t.Arguments)
	case "function":
		params, results := &ast.FieldList{}, &ast.FieldList{}
		for _, parameter := range t.Signature.Parameters {
			params.List = append(params.List, &ast.Field{Type: w.expression(parameter)})
		}
		for _, result := range t.Signature.Results {
			results.List = append(results.List, &ast.Field{Type: w.expression(result)})
		}
		return &ast.FuncType{Params: params, Results: results}
	case "pointer":
		return &ast.StarExpr{X: w.expression(*t.Element)}
	case "array":
		return &ast.ArrayType{Len: &ast.BasicLit{Kind: token.INT, Value: strconv.FormatInt(*t.Length, 10)}, Elt: w.expression(*t.Element)}
	case "map":
		return &ast.MapType{Key: w.expression(*t.Key), Value: w.expression(*t.Element)}
	case "slice":
		return &ast.ArrayType{Elt: w.expression(*t.Element)}
	case "channel":
		dir := ast.SEND | ast.RECV
		if t.Direction == "send" {
			dir = ast.SEND
		}
		if t.Direction == "receive" {
			dir = ast.RECV
		}
		return &ast.ChanType{Dir: dir, Value: w.expression(*t.Element)}
	default:
		return ast.NewIdent(t.Tag)
	}
}

func witness(packageName string, binding protocol.Binding, index int) ([]byte, error) {
	w := witnessTypes{aliases: map[string]string{}}
	if binding.CallMode == "ordinary" {
		w.aliases[binding.ImportPath] = "goml_target"
	}
	params := &ast.FieldList{}
	results := &ast.FieldList{}
	args := []ast.Expr{}
	for i, ty := range binding.Parameters {
		name := fmt.Sprintf("arg%d", i)
		params.List = append(params.List, &ast.Field{Names: []*ast.Ident{ast.NewIdent(name)}, Type: w.expression(ty)})
		args = append(args, ast.NewIdent(name))
	}
	stmts := []ast.Stmt{}
	returns := []ast.Expr{}
	for i, ty := range binding.Results {
		name := fmt.Sprintf("result%d", i)
		results.List = append(results.List, &ast.Field{Type: w.expression(ty)})
		stmts = append(stmts, &ast.DeclStmt{Decl: &ast.GenDecl{Tok: token.VAR, Specs: []ast.Spec{&ast.ValueSpec{Names: []*ast.Ident{ast.NewIdent(name)}, Type: w.expression(ty)}}}})
		returns = append(returns, ast.NewIdent(name))
	}
	var receiver ast.Expr = ast.NewIdent("goml_target")
	if binding.CallMode == "method" {
		if len(binding.Parameters) == 0 {
			return nil, fmt.Errorf("method binding requires a receiver")
		}
		receiver = &ast.ParenExpr{X: w.expression(binding.Parameters[0])}
	}
	call := &ast.CallExpr{Fun: &ast.SelectorExpr{X: receiver, Sel: ast.NewIdent(binding.Symbol)}, Args: args}
	if len(returns) == 0 {
		stmts = append(stmts, &ast.ExprStmt{X: call})
	} else {
		stmts = append(stmts, &ast.AssignStmt{Lhs: returns, Tok: token.ASSIGN, Rhs: []ast.Expr{call}}, &ast.ReturnStmt{Results: returns})
	}
	file := &ast.File{Name: ast.NewIdent(packageName)}
	if len(w.aliases) > 0 {
		file.Decls = append(file.Decls, w.imports())
	}
	file.Decls = append(file.Decls, &ast.FuncDecl{Name: ast.NewIdent(fmt.Sprintf("goml_ffi_witness_%d", index)), Type: &ast.FuncType{Params: params, Results: results}, Body: &ast.BlockStmt{List: stmts}})
	var output bytes.Buffer
	if err := format.Node(&output, token.NewFileSet(), file); err != nil {
		return nil, err
	}
	return output.Bytes(), nil
}

func typeWitness(packageName string, query protocol.TypeQuery, index int) ([]byte, error) {
	w := witnessTypes{aliases: map[string]string{query.ImportPath: "goml_target"}}
	target := w.named(query.ImportPath, query.Name, query.Arguments)
	file := &ast.File{Name: ast.NewIdent(packageName), Decls: []ast.Decl{
		w.imports(),
		&ast.GenDecl{Tok: token.TYPE, Specs: []ast.Spec{&ast.TypeSpec{Name: ast.NewIdent(fmt.Sprintf("goml_type_query_%d", index)), Assign: token.Pos(1), Type: target}}},
		&ast.GenDecl{Tok: token.VAR, Specs: []ast.Spec{&ast.ValueSpec{Names: []*ast.Ident{ast.NewIdent("_")}, Type: ast.NewIdent(fmt.Sprintf("goml_type_query_%d", index))}}},
	}}
	if query.Mode == "function" {
		file.Decls = []ast.Decl{
			w.imports(),
			&ast.GenDecl{Tok: token.VAR, Specs: []ast.Spec{&ast.ValueSpec{
				Names:  []*ast.Ident{ast.NewIdent(fmt.Sprintf("goml_type_query_%d", index))},
				Values: []ast.Expr{target},
			}}},
		}
	}
	if query.Mode == "declaration" {
		w.aliases = map[string]string{query.ImportPath: "_"}
		file.Decls = []ast.Decl{w.imports()}
	}
	var output bytes.Buffer
	if err := format.Node(&output, token.NewFileSet(), file); err != nil {
		return nil, err
	}
	return output.Bytes(), nil
}
