package cbind

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDynamicDeclarationsAndArgumentPlacement(t *testing.T) {
	p, ctx := fixture(t)
	t.Setenv("CGO_ENABLED", "0")
	t.Setenv("CC", "/bin/false")
	p.Config.Backend = "dynamic"
	p.Config.Libraries = []string{"libc.so.6"}
	p.Config.Constants = nil
	declaration := "enum Unused; typedef enum Unused Unused; typedef enum { Negative = -1, Positive = 3 } Mode; double answer(Mode, double, float, int, double, int, double, int, double, int, double, int, double, int, double, int, double, float);"
	if err := os.WriteFile(filepath.Join(p.Root, "api.h"), []byte(declaration), 0600); err != nil {
		t.Fatal(err)
	}
	w, err := Inspect(ctx, p)
	if err != nil {
		t.Fatal(err)
	}
	if w.EnumTypes["enum Mode"] != "int32" {
		t.Fatal(w.EnumTypes)
	}
	g, err := Generate(w)
	if err != nil {
		t.Fatal(err)
	}
	for _, expected := range []string{"StackCount: 4", "f.Float[1] = uint64(math.Float32bits(p2))", "f.Stack[0] = uint64(p13)", "f.Stack[1] = uint64(p15)", "f.Stack[2] = math.Float64bits(p16)", "f.Stack[3] = uint64(math.Float32bits(p17))", "return math.Float64frombits(f.FloatResult)"} {
		if !strings.Contains(string(g.Go), expected) {
			t.Fatalf("missing %s\n%s", expected, g.Go)
		}
	}
	if strings.Contains(string(g.Go), "import \"C\"") || strings.Contains(string(g.Go), "purego") {
		t.Fatal(string(g.Go))
	}
}

func TestDynamicUnsupportedDeclarations(t *testing.T) {
	for _, declaration := range []string{
		"static int answer(void) { return 1; }",
		"static inline int answer(void) { return 1; }",
		"__attribute__((ms_abi)) int answer(int);",
		"int answer(int, ...);",
		"long double answer(void);",
		"struct Pair { int a, b; }; struct Pair answer(void);",
		"void answer(void (*callback)(void));",
		"enum Huge { Maximum = 0xffffffffffffffffULL }; enum Huge answer(void);",
	} {
		t.Run(declaration, func(t *testing.T) {
			p, ctx := fixture(t)
			t.Setenv("CGO_ENABLED", "0")
			p.Config.Backend = "dynamic"
			p.Config.Libraries = []string{"libc.so.6"}
			p.Config.Constants = nil
			if err := os.WriteFile(filepath.Join(p.Root, "api.h"), []byte(declaration), 0600); err != nil {
				t.Fatal(err)
			}
			w, err := Inspect(ctx, p)
			if err == nil {
				_, err = Generate(w)
			}
			if err == nil {
				t.Fatal("accepted unsupported declaration")
			}
		})
	}
}

func TestDynamicConfigurationAndTarget(t *testing.T) {
	p, ctx := fixture(t)
	p.Config.Backend = "dynamic"
	if err := validate(p.Config); err == nil {
		t.Fatal("missing libraries")
	}
	p.Config.Libraries = []string{"libc.so.6"}
	if err := validate(p.Config); err != nil {
		t.Fatal(err)
	}
	t.Setenv("CGO_ENABLED", "1")
	if _, err := Inspect(ctx, p); err == nil || !strings.Contains(err.Error(), "CGO_ENABLED=0") {
		t.Fatal(err)
	}
	for _, library := range []string{"", "./local.so", "lib\x00c.so"} {
		p.Config.Libraries = []string{library}
		if err := validate(p.Config); err == nil {
			t.Fatal(library)
		}
	}
	p.Config.Libraries = []string{"libc.so.6"}
	p.Config.LDFlags = []string{"-lc"}
	if err := validate(p.Config); err == nil {
		t.Fatal("accepted linker flags")
	}
}
