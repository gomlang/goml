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
	results := newCheckResults(request)
	results.Diagnostics = []protocol.Diagnostic{{Code: code, Message: message}}
	results.failAll(code, message)
	return results.Response
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
	witnesses, diagnostic := prepareWitnesses(request)
	if diagnostic != nil {
		return failure(request, diagnostic.Code, diagnostic.Message)
	}
	flags, err := splitGoFlags(request.BuildContext.GOFLAGS)
	if err != nil {
		return failure(request, "ffi-protocol", err.Error())
	}
	if request.BuildContext.GO111MODULE == "on" {
		flags = append(flags, "-mod=readonly")
	}
	if len(request.BuildContext.BuildTags) > 0 {
		flags = append(flags, "-tags="+strings.Join(request.BuildContext.BuildTags, ","))
	}
	patterns := []string{"file=" + witnesses.firstFile}
	if caller.LoadMode == "files" {
		patterns = []string{witnesses.firstFile}
	}
	loadMode := packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles | packages.NeedImports | packages.NeedDeps | packages.NeedTypes | packages.NeedTypesSizes | packages.NeedModule
	for _, binding := range request.Bindings {
		if binding.CallMode == "method" {
			loadMode |= packages.NeedSyntax | packages.NeedTypesInfo
		}
	}
	loaded, err := packages.Load(&packages.Config{
		Context: ctx, Dir: request.BuildContext.ModuleDir, Env: append(env[:len(env):len(env)], "GOFLAGS="), BuildFlags: flags, Overlay: witnesses.overlay,
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
	results := newCheckResults(request)
	allPackages := map[string]*packages.Package{}
	packages.Visit(loaded, func(p *packages.Package) bool { allPackages[p.PkgPath] = p; return true }, nil)
	packageList := make([]*packages.Package, 0, len(allPackages))
	for _, p := range allPackages {
		packageList = append(packageList, p)
	}
	sort.Slice(packageList, func(i, j int) bool { return packageList[i].PkgPath < packageList[j].PkgPath })
	results.packageErrors(root, packageList, witnesses)
	graph := typeGraph{ids: map[types.Type]string{}}
	results.resolveBindings(request.Bindings, allPackages, selectedWitnessMethods(root, witnesses), &graph)
	results.resolveQueries(request.TypeQueries, allPackages, root, &graph)
	if graph.err != nil {
		return failure(request, "ffi-protocol", graph.err.Error())
	}
	results.ReferencedTypes = graph.nodes
	world := sha256.New()
	worldContext := request.BuildContext
	worldContext.GoExecutable = ""
	worldContext.ModuleDir = ""
	data, _ := json.Marshal(worldContext)
	world.Write(data)
	world.Write([]byte(caller.ImportPath + "\x00" + caller.LoadMode))
	for _, p := range packageList {
		world.Write([]byte("\x00" + p.PkgPath))
		files := append([]string{}, p.CompiledGoFiles...)
		sort.Strings(files)
		for _, file := range files {
			contents, exists := witnesses.overlay[file]
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
	results.GoWorldIdentity = hex.EncodeToString(world.Sum(nil))
	return results.Response
}
