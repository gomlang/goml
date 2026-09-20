package adapter

import (
	"fmt"
	"path/filepath"
	"strings"
	"sync"
	"testing"
)

func checked(t *testing.T) func(Handle, string) Handle {
	t.Helper()
	return func(value Handle, err string) Handle {
		t.Helper()
		if err != "" {
			t.Fatal(err)
		}
		return value
	}
}

func TestMalformedBitcodeAndRecoverableErrors(t *testing.T) {
	c := NewContext()
	defer Close(c)
	must := checked(t)
	m := must(Parse(c, "define i64 @f(i64 %x) { ret i64 %x }", false))
	bytes, err := Bitcode(m)
	if err != "" {
		t.Fatal(err)
	}
	for _, length := range []int{0, 1, 2, 3, 4, 8, 12, 16, 32, len(bytes) / 2, len(bytes) - 1} {
		if _, err := Parse(c, string(bytes[:length]), true); err == "" {
			t.Fatalf("accepted bitcode prefix %d", length)
		}
	}
	for _, input := range []string{"", "not bitcode", "BC\xc0\xde\x00\x00\x00\x00"} {
		if _, err := Parse(c, input, true); err == "" {
			t.Fatalf("accepted %q", input)
		}
	}
	if err := Verify(m); err != "" {
		t.Fatal(err)
	}
	if _, err := Primitive(c, 1, 0); !strings.HasPrefix(err, "argument:") {
		t.Fatal(err)
	}
	if _, err := NewModule(c, "bad\x00name"); !strings.HasPrefix(err, "argument:") {
		t.Fatal(err)
	}
	if _, err := Print(nil); !strings.HasPrefix(err, "argument:") {
		t.Fatal(err)
	}
	symbols := must(Parse(c, "@base = global i64 0\n@duplicate = alias i64, ptr @base\n@dispatch = ifunc i64 (), ptr @resolver\ndefine ptr @resolver() { ret ptr @target }\ndefine i64 @target() { ret i64 1 }", false))
	if err := Verify(symbols); err != "" {
		t.Fatal(err)
	}
	ty := must(Primitive(c, 1, 64))
	signature := must(Composite(c, 2, []Handle{ty}, 0, false))
	for _, name := range []string{"base", "duplicate", "dispatch", "target"} {
		if _, err := Function(symbols, name, signature, true); !strings.HasPrefix(err, "argument:") {
			t.Fatalf("duplicate %s: %s", name, err)
		}
	}
}

func TestOptimizationResetsBuildersAndInvalidatesFailedPassHandles(t *testing.T) {
	c := NewContext()
	defer Close(c)
	must := checked(t)
	ty := must(Primitive(c, 1, 64))
	signature := must(Composite(c, 2, []Handle{ty}, 0, false))
	m := must(NewModule(c, "optimize"))
	fn := must(Function(m, "before", signature, true))
	block := must(NewBlock(fn, "entry"))
	builder := must(NewBuilder(c))
	if err := Position(builder, block); err != "" {
		t.Fatal(err)
	}
	zero := must(Constant(ty, 0, 0, 0))
	must(Emit(builder, 0, []Handle{zero}, ""))
	before, err := Print(m)
	if err != "" {
		t.Fatal(err)
	}
	for _, assembly := range []bool{false, true} {
		if err := EmitFile(m, filepath.Join(t.TempDir(), "output"), assembly, 2); err != "" {
			t.Fatal(err)
		}
		after, err := Print(m)
		if err != "" || before != after {
			t.Fatal("emission changed module", err)
		}
		if _, err := Print(fn); err != "" {
			t.Fatal("emission invalidated function", err)
		}
	}
	if err := Optimize(m, "not-a-pass"); !strings.HasPrefix(err, "llvm:") {
		t.Fatal(err)
	}
	if _, err := Print(fn); !strings.HasPrefix(err, "stale:") {
		t.Fatal(err)
	}
	if _, err := Emit(builder, 0, []Handle{zero}, ""); !strings.HasPrefix(err, "argument:") {
		t.Fatal(err)
	}
	fresh := must(Function(m, "after", signature, true))
	if err := Position(builder, must(NewBlock(fresh, "entry"))); err != "" {
		t.Fatal(err)
	}
	must(Emit(builder, 0, []Handle{zero}, ""))
	if err := Verify(m); err != "" {
		t.Fatal(err)
	}
	if err := Optimize(m, "default<O2>"); err != "" {
		t.Fatal(err)
	}
	if _, err := Print(fresh); !strings.HasPrefix(err, "stale:") {
		t.Fatal(err)
	}
}

func TestSharedContextConcurrentModules(t *testing.T) {
	c := NewContext()
	defer Close(c)
	var workers sync.WaitGroup
	for worker := 0; worker < 12; worker++ {
		workers.Add(1)
		go func(worker int) {
			defer workers.Done()
			for iteration := 0; iteration < 20; iteration++ {
				m, err := Parse(c, fmt.Sprintf("define i64 @f%d(i64 %%x) { %%v = add i64 %%x, %d\n ret i64 %%v\n }", worker, iteration), false)
				if err != "" {
					t.Error(err)
					return
				}
				if err := Optimize(m, "default<O1>"); err != "" {
					t.Error(err)
					return
				}
				if err := Verify(m); err != "" {
					t.Error(err)
					return
				}
				if err := Close(m); err != "" {
					t.Error(err)
					return
				}
			}
		}(worker)
	}
	workers.Wait()
}

func TestCloseSerializesWithReadersAndOptimizers(t *testing.T) {
	c := NewContext()
	m := checked(t)(Parse(c, "define i64 @f(i64 %x) { ret i64 %x }", false))
	fn := checked(t)(Function(m, "f", nil, false))
	ty := checked(t)(Primitive(c, 1, 64))
	start := make(chan struct{})
	var workers sync.WaitGroup
	for worker := 0; worker < 16; worker++ {
		workers.Add(1)
		go func(worker int) {
			defer workers.Done()
			<-start
			for iteration := 0; iteration < 100; iteration++ {
				var err string
				switch worker % 4 {
				case 0:
					_, err = Print(fn)
				case 1:
					_, err = Constant(ty, 0, uint64(iteration), 0)
				case 2:
					err = Verify(m)
				case 3:
					err = Optimize(m, "default<O0>")
				}
				if err != "" && !strings.HasPrefix(err, "closed:") && !strings.HasPrefix(err, "stale:") {
					t.Errorf("unexpected result: %s", err)
					return
				}
			}
		}(worker)
	}
	workers.Add(1)
	go func() {
		defer workers.Done()
		<-start
		for i := 0; i < 10; i++ {
			Close(c)
			Close(m)
		}
	}()
	close(start)
	workers.Wait()
	if !IsClosed(c) || !IsClosed(m) || !IsClosed(fn) || !IsClosed(ty) {
		t.Fatal("close did not invalidate resources")
	}
}
