package validate

import (
	"encoding/json"
	"fmt"
	"go/token"
	"os"
	"path/filepath"

	"goml.dev/tools/goml-go-meta/internal/protocol"
)

type witnessTargets struct {
	bindings []int
	queries  []int
}

type witnessSet struct {
	overlay   map[string][]byte
	targets   map[string]witnessTargets
	firstFile string
	origins   witnessOrigins
}

func (w *witnessSet) add(file string, source []byte, targets witnessTargets) {
	w.overlay[file] = source
	w.targets[file] = targets
	if w.firstFile == "" {
		w.firstFile = file
	}
}

func prepareWitnesses(request protocol.Request) (*witnessSet, *protocol.Diagnostic) {
	w := &witnessSet{overlay: map[string][]byte{}, targets: map[string]witnessTargets{}}
	caller := request.CallerContext
	if value := request.SourceOverlay; value != nil {
		w.overlay[value.Path] = []byte(value.Source)
	}
	if caller.GeneratedSource != "" {
		if _, exists := w.overlay[filepath.Join(caller.Directory, "goml_generated.go")]; exists {
			return nil, &protocol.Diagnostic{Code: "ffi-protocol", Message: "conflicting generated source overlays"}
		}
		w.overlay[filepath.Join(caller.Directory, "goml_generated.go")] = []byte(caller.GeneratedSource)
	}
	keys := map[string]string{}
	for i, binding := range request.Bindings {
		if !token.IsIdentifier(binding.Symbol) || binding.Symbol == "_" {
			return nil, &protocol.Diagnostic{Code: "ffi-protocol", Message: "invalid Go symbol"}
		}
		identity := binding
		identity.ID = ""
		encoded, _ := json.Marshal(identity)
		if file, found := keys[string(encoded)]; found {
			target := w.targets[file]
			target.bindings = append(target.bindings, i)
			w.targets[file] = target
			continue
		}
		file := filepath.Join(caller.Directory, fmt.Sprintf("goml_ffi_witness_%d.go", i))
		if _, err := os.Lstat(file); !os.IsNotExist(err) {
			return nil, &protocol.Diagnostic{Code: "ffi-package-load", Message: "witness file conflicts with an existing file: " + file}
		}
		data, err := witness(caller.Package, binding, i)
		if err != nil {
			return nil, &protocol.Diagnostic{Code: "ffi-protocol", Message: err.Error()}
		}
		w.add(file, data, witnessTargets{bindings: []int{i}})
		keys[string(encoded)] = file
	}
	for i, query := range request.TypeQueries {
		if !token.IsIdentifier(query.Name) || query.Name == "_" {
			return nil, &protocol.Diagnostic{Code: "ffi-protocol", Message: "invalid Go type name"}
		}
		file := filepath.Join(caller.Directory, fmt.Sprintf("goml_type_query_%d.go", i))
		if _, err := os.Lstat(file); !os.IsNotExist(err) {
			return nil, &protocol.Diagnostic{Code: "ffi-package-load", Message: "type witness conflicts with an existing file: " + file}
		}
		data, err := typeWitness(caller.Package, query, i)
		if err != nil {
			return nil, &protocol.Diagnostic{Code: "ffi-protocol", Message: err.Error()}
		}
		w.add(file, data, witnessTargets{queries: []int{i}})
	}
	if caller.LoadMode == "files" {
		bundlePath := filepath.Join(caller.Directory, "goml_ffi_witness_bundle.go")
		if _, err := os.Lstat(bundlePath); !os.IsNotExist(err) {
			return nil, &protocol.Diagnostic{Code: "ffi-package-load", Message: "witness file conflicts with an existing file: " + bundlePath}
		}
		sources := map[string][]byte{}
		for path := range w.targets {
			sources[path] = w.overlay[path]
			delete(w.overlay, path)
		}
		data, mapping, err := bundleWitnesses(caller.Package, bundlePath, sources)
		if err != nil {
			return nil, &protocol.Diagnostic{Code: "ffi-protocol", Message: err.Error()}
		}
		w.overlay[bundlePath] = data
		w.origins = mapping
		w.firstFile = bundlePath
	}
	return w, nil
}
