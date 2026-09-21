package adapter

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
)

func TestTargetEmissionRejectsConflictsWithoutTouchingFiles(t *testing.T) {
	c := NewContext()
	defer Close(c)
	must := checked(t)
	m := must(Parse(c, "define i64 @f() { ret i64 42 }", false))
	x86 := must(NewTarget(c, "x86_64-unknown-linux-gnu", "x86-64", "-avx", 2, 2, 3))
	arm := must(NewTarget(c, "aarch64-unknown-linux-gnu", "", "", 2, 2, 0))
	if err := ConfigureTarget(arm, m); err != "" {
		t.Fatal(err)
	}
	path := filepath.Join(t.TempDir(), "output.o")
	if err := os.WriteFile(path, []byte("preserve"), 0600); err != nil {
		t.Fatal(err)
	}
	if err := EmitTargetFile(x86, m, path, false); !strings.HasPrefix(err, "argument:") {
		t.Fatal(err)
	}
	if err := EmitFile(m, path, false, 2); !strings.HasPrefix(err, "argument:") {
		t.Fatal(err)
	}
	content, err := os.ReadFile(path)
	if err != nil || string(content) != "preserve" {
		t.Fatal("rejected emission overwrote output", err)
	}
	if err := EmitTargetFile(arm, m, filepath.Join(path, "missing.o"), false); !strings.HasPrefix(err, "llvm:") {
		t.Fatal(err)
	}
	if err := EmitTargetFile(arm, m, path, false); err != "" {
		t.Fatal(err)
	}
	object, failure := EmitTargetBytes(arm, m, false, 1<<20)
	if failure != "" {
		t.Fatal(failure)
	}
	content, err = os.ReadFile(path)
	if err != nil || !bytes.Equal(content, object) {
		t.Fatal("memory and file emission disagree", err)
	}
}

func TestTargetEntryPointsRejectWrongKindsAndOptions(t *testing.T) {
	c := NewContext()
	defer Close(c)
	must := checked(t)
	m := must(NewModule(c, "m"))
	target := must(NewTarget(c, "x86_64-unknown-linux-gnu", "", "", 2, 2, 0))
	if _, err := NewTarget(m, "x86_64-unknown-linux-gnu", "", "", 2, 2, 0); !strings.HasPrefix(err, "argument:") {
		t.Fatal(err)
	}
	for _, options := range [][3]int{{-1, 2, 0}, {4, 2, 0}, {2, 7, 0}, {2, 2, 1}, {2, 2, 2}, {2, 2, 4}, {2, 2, 5}} {
		if _, err := NewTarget(c, "x86_64-unknown-linux-gnu", "", "", options[0], options[1], options[2]); !strings.HasPrefix(err, "argument:") {
			t.Fatalf("options %v: %s", options, err)
		}
	}
	if _, err := NewTarget(c, "x86_64-unknown-linux-gnu", "", "+sse2,", 2, 2, 0); !strings.HasPrefix(err, "argument:") {
		t.Fatal(err)
	}
	if err := ConfigureTarget(m, target); !strings.HasPrefix(err, "argument:") {
		t.Fatal(err)
	}
	if _, err := TargetText(m, 0); !strings.HasPrefix(err, "argument:") {
		t.Fatal(err)
	}
	if _, err := TargetText(target, 5); !strings.HasPrefix(err, "argument:") {
		t.Fatal(err)
	}
	if _, err := EmitTargetBytes(m, target, false, 1<<20); !strings.HasPrefix(err, "argument:") {
		t.Fatal(err)
	}
}

func TestSharedTargetConcurrentEmissionAndContextClose(t *testing.T) {
	c := NewContext()
	defer Close(c)
	must := checked(t)
	m := must(Parse(c, "define i64 @f(i64 %x) { %v = add i64 %x, 42\n ret i64 %v }", false))
	target := must(NewTarget(c, "aarch64-unknown-linux-gnu", "", "", 2, 2, 0))
	ty := must(Primitive(c, 1, 64))
	var workers sync.WaitGroup
	for worker := 0; worker < 8; worker++ {
		workers.Add(1)
		go func() {
			defer workers.Done()
			for iteration := 0; iteration < 4; iteration++ {
				output, err := EmitTargetBytes(target, m, false, 1<<20)
				if err != "" || len(output) < 20 || output[18] != 183 {
					t.Errorf("bad concurrent emission: %s", err)
					return
				}
			}
		}()
	}
	workers.Wait()
	start := make(chan struct{})
	for worker := 0; worker < 8; worker++ {
		workers.Add(1)
		go func(worker int) {
			defer workers.Done()
			<-start
			for iteration := 0; iteration < 30; iteration++ {
				var err string
				switch worker % 3 {
				case 0:
					_, err = EmitTargetBytes(target, m, false, 1<<20)
				case 1:
					_, err = TargetLayout(target, ty, 0, 0)
				case 2:
					_, err = TargetText(target, 4)
				}
				if err != "" && !strings.HasPrefix(err, "closed:") {
					t.Error(err)
				}
			}
		}(worker)
	}
	workers.Add(1)
	go func() { defer workers.Done(); <-start; Close(c); Close(target) }()
	close(start)
	workers.Wait()
	if !IsClosed(target) || !IsClosed(m) {
		t.Fatal("context close did not invalidate target and module")
	}
}

func TestSupportedRelocationsAndCodeModelsEmitObjects(t *testing.T) {
	c := NewContext()
	defer Close(c)
	must := checked(t)
	m := must(Parse(c, "define i64 @f() { ret i64 42 }", false))
	for _, triple := range []string{"x86_64-unknown-linux-gnu", "i386-unknown-linux-gnu", "aarch64-unknown-linux-gnu"} {
		for _, model := range []int{0, 3, 6} {
			if strings.HasPrefix(triple, "i386") && model == 6 {
				continue
			}
			for relocation := 0; relocation <= 2; relocation++ {
				target := must(NewTarget(c, triple, "", "", 2, relocation, model))
				output, err := EmitTargetBytes(target, m, false, 1<<20)
				if err != "" || len(output) < 20 || output[0] != 127 {
					t.Fatalf("triple=%s model=%d relocation=%d: %s", triple, model, relocation, err)
				}
				if err := Close(target); err != "" {
					t.Fatal(err)
				}
			}
		}
	}
}
