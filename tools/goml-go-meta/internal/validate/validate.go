package validate

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/tools/go/packages"
	"goml.dev/tools/goml-go-meta/internal/protocol"
)

func failure(request protocol.Request, code, message string) protocol.Response {
	response := protocol.Response{ProtocolVersion: protocol.Version, Diagnostics: []protocol.Diagnostic{{Code: code, Message: message}}}
	for _, binding := range request.Bindings {
		response.Bindings = append(response.Bindings, protocol.BindingResult{ID: binding.ID, Status: "failed", Diagnostics: []protocol.Diagnostic{{Code: code, Message: message}}})
	}
	for _, query := range request.TypeQueries {
		response.TypeResults = append(response.TypeResults, protocol.TypeResult{ID: query.ID, Status: "failed", Diagnostics: []protocol.Diagnostic{{Code: code, Message: message}}})
	}
	return response
}

func Check(ctx context.Context, request protocol.Request) protocol.Response {
	if err := request.Validate(); err != nil {
		return failure(request, "ffi-protocol", err.Error())
	}
	if err := ctx.Err(); err != nil {
		return failure(request, "ffi-canceled", err.Error())
	}
	if len(request.Bindings) == 0 && len(request.TypeQueries) == 0 {
		return protocol.Response{ProtocolVersion: protocol.Version}
	}
	env := environment(request.BuildContext)
	if err := verifyToolchain(ctx, request.BuildContext, env); err != nil {
		if ctx.Err() != nil {
			return failure(request, "ffi-canceled", ctx.Err().Error())
		}
		return failure(request, "ffi-tool-unavailable", err.Error())
	}
	caller := request.CallerContext
	if !token.IsIdentifier(caller.Package) || caller.Package == "_" {
		return failure(request, "ffi-protocol", "invalid caller package name")
	}
	if !filepath.IsAbs(caller.Directory) || !filepath.IsAbs(request.BuildContext.ModuleDir) {
		return failure(request, "ffi-protocol", "build and caller directories must be absolute")
	}
	overlay := map[string][]byte{}
	if value := request.SourceOverlay; value != nil {
		overlay[value.Path] = []byte(value.Source)
	}
	if caller.GeneratedSource != "" {
		if _, exists := overlay[filepath.Join(caller.Directory, "goml_generated.go")]; exists {
			return failure(request, "ffi-protocol", "conflicting generated source overlays")
		}
		overlay[filepath.Join(caller.Directory, "goml_generated.go")] = []byte(caller.GeneratedSource)
	}
	fileBindings := map[string][]int{}
	keys := map[string]string{}
	var firstFile string
	for i, binding := range request.Bindings {
		if !token.IsIdentifier(binding.Symbol) || binding.Symbol == "_" {
			return failure(request, "ffi-protocol", "invalid Go symbol")
		}
		identity := binding
		identity.ID = ""
		encoded, _ := json.Marshal(identity)
		if file, found := keys[string(encoded)]; found {
			fileBindings[file] = append(fileBindings[file], i)
			continue
		}
		file := filepath.Join(caller.Directory, fmt.Sprintf("goml_ffi_witness_%d.go", i))
		if _, err := os.Lstat(file); !os.IsNotExist(err) {
			return failure(request, "ffi-package-load", "witness file conflicts with an existing file: "+file)
		}
		data, err := witness(caller.Package, binding, i)
		if err != nil {
			return failure(request, "ffi-protocol", err.Error())
		}
		overlay[file] = data
		keys[string(encoded)] = file
		fileBindings[file] = []int{i}
		if firstFile == "" {
			firstFile = file
		}
	}
	fileQueries := map[string]int{}
	for i, query := range request.TypeQueries {
		if !token.IsIdentifier(query.Name) || query.Name == "_" {
			return failure(request, "ffi-protocol", "invalid Go type name")
		}
		file := filepath.Join(caller.Directory, fmt.Sprintf("goml_type_query_%d.go", i))
		if _, err := os.Lstat(file); !os.IsNotExist(err) {
			return failure(request, "ffi-package-load", "type witness conflicts with an existing file: "+file)
		}
		data, err := typeWitness(caller.Package, query, i)
		if err != nil {
			return failure(request, "ffi-protocol", err.Error())
		}
		overlay[file] = data
		fileQueries[file] = i
		if firstFile == "" {
			firstFile = file
		}
	}
	flags := []string{}
	if request.BuildContext.GO111MODULE == "on" {
		flags = append(flags, "-mod=readonly")
	}
	if len(request.BuildContext.BuildTags) > 0 {
		flags = append(flags, "-tags="+strings.Join(request.BuildContext.BuildTags, ","))
	}
	patterns := []string{"file=" + firstFile}
	if caller.LoadMode == "files" {
		patterns = nil
		for file := range overlay {
			patterns = append(patterns, file)
		}
		sort.Strings(patterns)
	}
	loadMode := packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles | packages.NeedImports | packages.NeedDeps | packages.NeedTypes | packages.NeedTypesSizes | packages.NeedModule
	for _, binding := range request.Bindings {
		if binding.CallMode == "method" {
			loadMode |= packages.NeedSyntax | packages.NeedTypesInfo
		}
	}
	loaded, err := packages.Load(&packages.Config{
		Context: ctx, Dir: request.BuildContext.ModuleDir, Env: env, BuildFlags: flags, Overlay: overlay,
		Mode: loadMode,
	}, patterns...)
	if err := ctx.Err(); err != nil {
		return failure(request, "ffi-canceled", err.Error())
	}
	if err != nil {
		return failure(request, "ffi-package-load", err.Error())
	}
	if len(loaded) != 1 {
		return failure(request, "ffi-package-load", "expected exactly one caller package")
	}
	root := loaded[0]
	if request.BuildContext.GO111MODULE == "on" && root.PkgPath != caller.ImportPath {
		return failure(request, "ffi-package-load", fmt.Sprintf("caller import path mismatch: expected %s, found %s", caller.ImportPath, root.PkgPath))
	}
	response := protocol.Response{ProtocolVersion: protocol.Version}
	for _, binding := range request.Bindings {
		response.Bindings = append(response.Bindings, protocol.BindingResult{ID: binding.ID, Status: "verified"})
	}
	for _, query := range request.TypeQueries {
		response.TypeResults = append(response.TypeResults, protocol.TypeResult{ID: query.ID, Status: "verified"})
	}
	addType := func(index int, code, message string) {
		result := &response.TypeResults[index]
		result.Status = "failed"
		for _, d := range result.Diagnostics {
			if d.Code == code && d.Message == message {
				return
			}
		}
		result.Diagnostics = append(result.Diagnostics, protocol.Diagnostic{Code: code, Message: message})
	}
	add := func(index int, code, message string) {
		r := &response.Bindings[index]
		r.Status = "failed"
		for _, d := range r.Diagnostics {
			if d.Code == code && d.Message == message {
				return
			}
		}
		r.Diagnostics = append(r.Diagnostics, protocol.Diagnostic{Code: code, Message: message})
	}
	allPackages := map[string]*packages.Package{}
	packages.Visit(loaded, func(p *packages.Package) bool { allPackages[p.PkgPath] = p; return true }, nil)
	paths := make([]string, 0, len(allPackages))
	for path := range allPackages {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		p := allPackages[path]
		for _, problem := range p.Errors {
			indices := []int{}
			for file, bindings := range fileBindings {
				if strings.HasPrefix(problem.Pos, file+":") {
					indices = append(indices, bindings...)
				}
			}
			queryIndices := []int{}
			for file, index := range fileQueries {
				if strings.HasPrefix(problem.Pos, file+":") {
					queryIndices = append(queryIndices, index)
				}
			}
			if len(indices) == 0 && len(queryIndices) == 0 {
				for i := range response.TypeResults {
					queryIndices = append(queryIndices, i)
				}
				for i := range response.Bindings {
					indices = append(indices, i)
				}
			}
			code := "ffi-package-load"
			if p == root && problem.Kind == packages.TypeError {
				code = "ffi-signature-mismatch"
			}
			for _, index := range indices {
				add(index, code, problem.Msg)
			}
			for _, index := range queryIndices {
				addType(index, code, problem.Msg)
			}
		}
	}
	graph := typeGraph{ids: map[types.Type]string{}}
	methods := selectedWitnessMethods(root, fileBindings)
	for i, binding := range request.Bindings {
		if binding.CallMode == "method" {
			method := methods[i]
			if method == nil || !method.Exported() {
				add(i, "ffi-symbol-not-found", "Go receiver has no exported method: "+binding.Symbol)
				continue
			}
			response.Bindings[i].SignatureType = graph.add(method.Type())
			response.Bindings[i].ActualSignature = types.TypeString(method.Type(), func(p *types.Package) string { return p.Path() })
			continue
		}
		p := allPackages[binding.ImportPath]
		if p == nil || p.Types == nil {
			add(i, "ffi-package-load", "could not load "+binding.ImportPath)
			continue
		}
		if p.Name == "main" {
			add(i, "ffi-package-load", "cannot import package main")
			continue
		}
		object := p.Types.Scope().Lookup(binding.Symbol)
		if object == nil {
			add(i, "ffi-symbol-not-found", "Go symbol not found: "+binding.ImportPath+"."+binding.Symbol)
			continue
		}
		response.Bindings[i].SignatureType = graph.add(object.Type())
		response.Bindings[i].ActualSignature = types.TypeString(object.Type(), func(p *types.Package) string { return p.Path() })
		if !object.Exported() {
			add(i, "ffi-symbol-not-found", "Go symbol is not exported: "+binding.Symbol)
		}
		if _, ok := object.(*types.Func); !ok {
			add(i, "ffi-signature-mismatch", "Go symbol is not a package-level function: "+binding.Symbol)
		}
	}
	for i, query := range request.TypeQueries {
		p := allPackages[query.ImportPath]
		if p == nil || p.Types == nil || p.Name == "main" {
			addType(i, "ffi-package-load", "could not import "+query.ImportPath)
			continue
		}
		object := p.Types.Scope().Lookup(query.Name)
		if object == nil || !object.Exported() {
			addType(i, "ffi-symbol-not-found", "Go type is missing or not exported: "+query.Name)
			continue
		}
		if query.Mode == "function" {
			function, ok := object.(*types.Func)
			if !ok {
				addType(i, "ffi-signature-mismatch", "Go object is not a package-level function: "+query.Name)
				continue
			}
			result := &response.TypeResults[i]
			result.Kind = "function"
			result.DeclaredType = graph.add(function.Type())
			if result.Status != "verified" {
				continue
			}
			resolved := root.Types.Scope().Lookup(fmt.Sprintf("goml_type_query_%d", i))
			if resolved == nil {
				addType(i, "ffi-package-load", "missing resolved function witness")
				continue
			}
			signature, ok := resolved.Type().(*types.Signature)
			if !ok || signature.TypeParams().Len() != 0 {
				addType(i, "ffi-signature-mismatch", "function query requires a concrete signature")
				continue
			}
			result.Type = graph.add(signature)
			continue
		}
		declared, ok := object.(*types.TypeName)
		if !ok {
			addType(i, "ffi-signature-mismatch", "Go object is not a type: "+query.Name)
			continue
		}
		result := &response.TypeResults[i]
		result.Kind = "defined"
		if declared.IsAlias() {
			result.Kind = "alias"
		}
		result.DeclaredType = graph.add(declared.Type())
		if query.Mode == "declaration" {
			continue
		}
		resolved := root.Types.Scope().Lookup(fmt.Sprintf("goml_type_query_%d", i))
		if result.Status != "verified" {
			continue
		}
		if resolved == nil {
			addType(i, "ffi-package-load", "missing resolved type witness")
			continue
		}
		result.Type = graph.add(types.Unalias(resolved.Type()))
	}
	if graph.err != nil {
		return failure(request, "ffi-protocol", graph.err.Error())
	}
	response.ReferencedTypes = graph.nodes
	world := sha256.New()
	worldContext := request.BuildContext
	worldContext.GoExecutable = ""
	worldContext.ModuleDir = ""
	data, _ := json.Marshal(worldContext)
	world.Write(data)
	world.Write([]byte(caller.ImportPath + "\x00" + caller.LoadMode))
	for _, path := range paths {
		p := allPackages[path]
		world.Write([]byte("\x00" + path))
		files := append([]string{}, p.CompiledGoFiles...)
		sort.Strings(files)
		for _, file := range files {
			contents, exists := overlay[file]
			if !exists {
				var readErr error
				contents, readErr = os.ReadFile(file)
				if readErr != nil {
					return failure(request, "ffi-package-load", readErr.Error())
				}
			}
			world.Write([]byte("\x00" + filepath.Base(file) + "\x00"))
			digest := sha256.Sum256(contents)
			world.Write(digest[:])
		}
		if p.Module != nil {
			world.Write([]byte("\x00" + p.Module.Path + "\x00" + p.Module.Version))
		}
	}
	response.GoWorldIdentity = hex.EncodeToString(world.Sum(nil))
	return response
}

func selectedWitnessMethods(root *packages.Package, files map[string][]int) map[int]*types.Func {
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
		for _, index := range files[root.Fset.Position(expression.Pos()).Filename] {
			result[index] = method
		}
	}
	return result
}
