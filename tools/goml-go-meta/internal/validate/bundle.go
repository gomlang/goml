package validate

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"sort"
	"strconv"
	"strings"
)

type witnessOrigins struct {
	path  string
	lines map[int][]string
}

type witnessDeclaration struct {
	node   ast.Decl
	owners []string
}

func (o witnessOrigins) files(path string, line int) []string {
	if path == o.path {
		return o.lines[line]
	}
	return []string{path}
}

func (o witnessOrigins) diagnostic(position string) []string {
	end := strings.LastIndex(position, ":")
	if end < 0 {
		return nil
	}
	start := strings.LastIndex(position[:end], ":")
	if start < 0 {
		return nil
	}
	line, err := strconv.Atoi(position[start+1 : end])
	if err != nil {
		return nil
	}
	return o.files(position[:start], line)
}

func bundleWitnesses(packageName, path string, sources map[string][]byte) ([]byte, witnessOrigins, error) {
	paths := make([]string, 0, len(sources))
	for name := range sources {
		paths = append(paths, name)
	}
	sort.Strings(paths)
	aliases := map[string]string{}
	nextAlias := 0
	importOwners := map[string][]string{}
	declarations := []witnessDeclaration{}
	for _, name := range paths {
		file, err := parser.ParseFile(token.NewFileSet(), name, sources[name], 0)
		if err != nil {
			return nil, witnessOrigins{}, err
		}
		renames := map[string]string{}
		for _, spec := range file.Imports {
			importPath, err := strconv.Unquote(spec.Path.Value)
			if err != nil {
				return nil, witnessOrigins{}, err
			}
			importOwners[importPath] = append(importOwners[importPath], name)
			alias, exists := aliases[importPath]
			if !exists || alias == "_" && spec.Name.Name != "_" {
				alias = "_"
				if spec.Name.Name != "_" {
					alias = fmt.Sprintf("goml_bundle_import_%d", nextAlias)
					nextAlias++
				}
				aliases[importPath] = alias
			}
			if spec.Name.Name != "_" {
				renames[spec.Name.Name] = alias
			}
		}
		for _, declaration := range file.Decls {
			if general, ok := declaration.(*ast.GenDecl); ok && general.Tok == token.IMPORT {
				continue
			}
			ast.Inspect(declaration, func(node ast.Node) bool {
				if expression, ok := node.(*ast.SelectorExpr); ok {
					if ident, ok := expression.X.(*ast.Ident); ok {
						if replacement, found := renames[ident.Name]; found {
							ident.Name = replacement
						}
					}
				}
				return true
			})
			declarations = append(declarations, witnessDeclaration{node: declaration, owners: []string{name}})
		}
	}
	file := &ast.File{Name: ast.NewIdent(packageName)}
	allDeclarations := make([]witnessDeclaration, 0, len(aliases)+len(declarations))
	imports := make([]string, 0, len(aliases))
	for name := range aliases {
		imports = append(imports, name)
	}
	sort.Strings(imports)
	for _, name := range imports {
		declaration := &ast.GenDecl{Tok: token.IMPORT, Specs: []ast.Spec{
			&ast.ImportSpec{Name: ast.NewIdent(aliases[name]), Path: &ast.BasicLit{Kind: token.STRING, Value: strconv.Quote(name)}},
		}}
		allDeclarations = append(allDeclarations, witnessDeclaration{node: declaration, owners: importOwners[name]})
	}
	allDeclarations = append(allDeclarations, declarations...)
	for _, declaration := range allDeclarations {
		file.Decls = append(file.Decls, declaration.node)
	}
	var output bytes.Buffer
	if err := format.Node(&output, token.NewFileSet(), file); err != nil {
		return nil, witnessOrigins{}, err
	}
	positions := token.NewFileSet()
	parsed, err := parser.ParseFile(positions, path, output.Bytes(), 0)
	if err != nil {
		return nil, witnessOrigins{}, err
	}
	origins := witnessOrigins{path: path, lines: map[int][]string{}}
	for index, declaration := range parsed.Decls {
		for line := positions.Position(declaration.Pos()).Line; line <= positions.Position(declaration.End()).Line; line++ {
			origins.lines[line] = allDeclarations[index].owners
		}
	}
	return output.Bytes(), origins, nil
}
