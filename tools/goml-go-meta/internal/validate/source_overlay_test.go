package validate

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"goml.dev/tools/goml-go-meta/internal/protocol"
)

func TestSourceOverlayReplacesStaleImportedCodeWithoutWriting(t *testing.T) {
	request := fixture(t)
	file := filepath.Join(request.BuildContext.ModuleDir, "shim", "generated.go")
	original := "package shim\nfunc Generated() int64 { return Missing() }\n"
	if err := os.WriteFile(file, []byte(original), 0644); err != nil {
		t.Fatal(err)
	}
	request.TypeQueries = []protocol.TypeQuery{{Mode: "function", ID: "value", ImportPath: "example.com/host/shim", Name: "Scalar"}}
	if result := Check(context.Background(), request); result.TypeResults[0].Status == "verified" {
		t.Fatal("stale source was ignored")
	}
	request.SourceOverlay = &protocol.SourceOverlay{Path: file, Source: "package shim\n"}
	result := Check(context.Background(), request)
	if result.TypeResults[0].Status != "verified" {
		t.Fatalf("overlay query: %+v", result)
	}
	first := result.GoWorldIdentity
	request.SourceOverlay.Source = "package shim\nfunc Generated(value int64) int64 { return Scalar(value) }\n"
	request.TypeQueries[0].Name = "Generated"
	result = Check(context.Background(), request)
	if result.TypeResults[0].Status != "verified" || result.GoWorldIdentity == first {
		t.Fatalf("candidate validation: %+v", result)
	}
	contents, err := os.ReadFile(file)
	if err != nil || string(contents) != original {
		t.Fatalf("overlay changed source: %s %v", contents, err)
	}
	request.SourceOverlay.Source = "package shim\nfunc Scalar() {}\n"
	if result := Check(context.Background(), request); result.TypeResults[0].Status == "verified" {
		t.Fatal("duplicate symbol accepted")
	}
}

func TestSourceOverlayRejectsUnboundedAndNonModulePaths(t *testing.T) {
	request := fixture(t)
	for _, file := range []string{"relative.go", filepath.Join(filepath.Dir(request.BuildContext.ModuleDir), "escape.go"), filepath.Join(request.BuildContext.ModuleDir, "go.mod"), filepath.Join(request.BuildContext.ModuleDir, "goml_type_query_0.go")} {
		request.SourceOverlay = &protocol.SourceOverlay{Path: file, Source: "package value\n"}
		if err := request.Validate(); err == nil {
			t.Fatalf("accepted overlay: %q", file)
		}
	}
	request.SourceOverlay = &protocol.SourceOverlay{Path: filepath.Join(request.BuildContext.ModuleDir, "candidate.go"), Source: ""}
	if err := request.Validate(); err == nil {
		t.Fatal("accepted empty overlay")
	}
}
