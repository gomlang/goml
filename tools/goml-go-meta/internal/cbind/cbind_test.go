package cbind

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func fixture(t *testing.T) (Project, context.Context) {
	t.Helper()
	if _, err := clangProgram(); err != nil {
		t.Skip(err)
	}
	root := t.TempDir()
	write := func(name, value string) {
		t.Helper()
		if err := os.WriteFile(filepath.Join(root, name), []byte(value), 0600); err != nil {
			t.Fatal(err)
		}
	}
	write("goml.toml", "[module]\npath = \"test\"\n")
	write("go.mod", "module example.com/test\n\ngo 1.26.0\n")
	write("api.h", "#include <stdint.h>\n#define MASK UINT64_C(18446744073709551615)\nstatic int answer(void) { return 42; }\n")
	config := Config{Version: 1, Package: "bindings", Output: "bindings/generated.gom", GoPackage: "native", GoOutput: "native/generated.go", Headers: []string{"api.h"}, Functions: []Function{{Name: "answer", Symbol: "answer"}}, Constants: []Constant{{Name: "MASK", Symbol: "MASK"}}}
	data, _ := json.Marshal(config)
	write("bindings.json", string(data))
	p, err := Load(filepath.Join(root, "bindings.json"))
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	t.Cleanup(cancel)
	return p, ctx
}

func execute(t *testing.T, ctx context.Context, p Project, options ...string) string {
	t.Helper()
	var output bytes.Buffer
	if err := Execute(ctx, append([]string{p.File}, options...), "cat", &output); err != nil {
		t.Fatal(err)
	}
	return output.String()
}

func TestGenerationFreshnessAndOwnership(t *testing.T) {
	p, ctx := fixture(t)
	execute(t, ctx, p, "--dry-run")
	if _, err := os.Stat(filepath.Dir(p.GoFile)); !os.IsNotExist(err) {
		t.Fatal("dry run created outputs")
	}
	execute(t, ctx, p)
	before := execute(t, ctx, p, "--check")
	if len(strings.TrimSpace(before)) != 64 {
		t.Fatal(before)
	}
	data, _ := os.ReadFile(p.GomlFile)
	if !bytes.Contains(data, []byte("pub const MASK: u64 = 18446744073709551615;")) {
		t.Fatal(string(data))
	}
	execute(t, ctx, p)
	after := execute(t, ctx, p, "--check")
	if before != after {
		t.Fatal("nondeterministic C fingerprint")
	}
	header := filepath.Join(p.Root, "api.h")
	source, _ := os.ReadFile(header)
	if err := os.WriteFile(header, bytes.ReplaceAll(source, []byte("return 42"), []byte("return 43")), 0600); err != nil {
		t.Fatal(err)
	}
	if before == execute(t, ctx, p, "--check") {
		t.Fatal("C implementation change did not invalidate fingerprint")
	}
	if err := os.WriteFile(header, bytes.ReplaceAll(source, []byte("static int answer"), []byte("static long answer")), 0600); err != nil {
		t.Fatal(err)
	}
	if err := Execute(ctx, []string{p.File, "--check"}, "cat", &bytes.Buffer{}); err == nil || !strings.Contains(err.Error(), "declarations") {
		t.Fatal(err)
	}
	execute(t, ctx, p)
	if err := os.WriteFile(p.GoFile, []byte("package native\n"), 0600); err != nil {
		t.Fatal(err)
	}
	if err := Execute(ctx, []string{p.File}, "cat", &bytes.Buffer{}); err == nil || !strings.Contains(err.Error(), "modified") {
		t.Fatal(err)
	}
}

func TestGenerationLinksWithoutExecutingInitializers(t *testing.T) {
	p, ctx := fixture(t)
	if err := os.Mkdir(filepath.Dir(p.GoFile), 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(filepath.Dir(p.GoFile), "init.go"), []byte("package native\nfunc init() { panic(\"must not execute\") }\n"), 0600); err != nil {
		t.Fatal(err)
	}
	execute(t, ctx, p)
	header := filepath.Join(p.Root, "api.h")
	if err := os.WriteFile(header, []byte("#include <stdint.h>\n#define MASK 1\nint answer(void);\n"), 0600); err != nil {
		t.Fatal(err)
	}
	before, _ := os.ReadFile(p.GoFile)
	if err := Execute(ctx, []string{p.File}, "cat", &bytes.Buffer{}); err == nil || !strings.Contains(err.Error(), "native compile/link") {
		t.Fatal(err)
	}
	after, _ := os.ReadFile(p.GoFile)
	if !bytes.Equal(before, after) {
		t.Fatal("failed link changed output")
	}
}

func TestInspectionUsesGoDefaultCFlags(t *testing.T) {
	p, ctx := fixture(t)
	t.Setenv("CGO_CFLAGS", "")
	t.Setenv("GOENV", "off")
	source := "#include <stdint.h>\n#if defined(__OPTIMIZE__) && defined(__STRICT_ANSI__)\n#define MASK 1\n#else\n#define MASK 2\n#endif\nstatic int answer(void) { return MASK; }\n"
	if err := os.WriteFile(filepath.Join(p.Root, "api.h"), []byte(source), 0600); err != nil {
		t.Fatal(err)
	}
	execute(t, ctx, p)
	data, _ := os.ReadFile(p.GomlFile)
	if !bytes.Contains(data, []byte("pub const MASK: i32 = 1;")) {
		t.Fatal(string(data))
	}
	t.Setenv("CGO_CFLAGS", "-O2 -std=gnu11")
	execute(t, ctx, p, "--check")
}

func TestUnsupportedSignaturesAndAbi(t *testing.T) {
	for _, entry := range []struct{ source, message string }{
		{"int answer(int, ...);", "non-variadic"},
		{"int answer();", "non-variadic"},
		{"struct X { int value; }; struct X answer(void);", "unsupported C type"},
		{"void answer(void (*callback)(void));", "unsupported C type"},
		{"long double answer(void);", "unsupported C type"},
		{"int missing(void);", "declared"},
	} {
		t.Run(entry.message+entry.source[:4], func(t *testing.T) {
			p, ctx := fixture(t)
			p.Config.Constants = nil
			if err := os.WriteFile(filepath.Join(p.Root, "api.h"), []byte(entry.source), 0600); err != nil {
				t.Fatal(err)
			}
			w, err := Inspect(ctx, p)
			if err == nil {
				_, err = Generate(w)
			}
			if err == nil || !strings.Contains(err.Error(), entry.message) {
				t.Fatal(err)
			}
		})
	}
}

func TestConfigurationRejectsAmbiguityAndUnsafePaths(t *testing.T) {
	for _, source := range []string{`{"version":1,"version":1}`, `{"unknown":true}`, `{} {}`, `{"headers":["x"],"headers":["y"]}`} {
		var config Config
		if err := decode([]byte(source), &config); err == nil {
			t.Fatal(source)
		}
	}
	p, _ := fixture(t)
	for _, output := range []string{"../outside.go", ".goml-bind-c-lock/recovery.go", "/tmp/outside.go"} {
		if _, err := checkedPath(p.Root, p.Directory, output); err == nil {
			t.Fatal(output)
		}
	}
	if err := os.Symlink(t.TempDir(), filepath.Join(p.Root, "link")); err != nil {
		t.Fatal(err)
	}
	if _, err := checkedPath(p.Root, p.Root, "link/out.go"); err == nil {
		t.Fatal("followed output symlink")
	}
	if err := os.Symlink(filepath.Join(p.Root, "other"), p.Manifest); err != nil {
		t.Fatal(err)
	}
	if _, err := Load(p.File); err == nil {
		t.Fatal("followed manifest symlink")
	}
}
