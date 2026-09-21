package adapter

import (
	"strings"
	"sync"
	"testing"
)

func readIR(t *testing.T, handle Handle, query int) []Handle {
	t.Helper()
	values, err := ReadIR(handle, query)
	if err != "" {
		t.Fatal(err)
	}
	return values
}

func TestPhiRebuildPreservesMetadataSelfUsesAndAnchors(t *testing.T) {
	ctx := NewContext()
	defer Close(ctx)
	must := checked(t)
	module := must(Parse(ctx, "define double @loop(i1 %c, double %x) {\nentry: br label %loop\nloop: %p = phi fast double [%x, %entry], [%p, %loop], !custom !0\n br i1 %c, label %loop, label %exit\nexit: ret double %p\n}\n!0 = !{i64 7}", false))
	fn := must(Function(module, "loop", nil, false))
	blocks := readIR(t, fn, 1)
	phi := readIR(t, blocks[1], 2)[0]
	builder := must(NewBuilder(ctx))
	if err := PositionBefore(builder, phi); err != "" {
		t.Fatal(err)
	}
	saved := must(SavePosition(builder))
	x := must(Parameter(fn, 1))
	before, _ := Print(module)
	for iteration := 0; iteration < 20; iteration++ {
		if err := EditIncoming(phi, 1, 0, []Handle{x, phi}, []Handle{blocks[0], blocks[1]}); err != "" {
			t.Fatal(err)
		}
		if err := Verify(module); err != "" {
			t.Fatal(err)
		}
		incoming := readIR(t, phi, 7)
		if equal, err := HandleEqual(incoming[2], phi); err != "" || !equal {
			t.Fatal("self edge lost", err)
		}
		if equal, err := HandleEqual(phi, readIR(t, blocks[1], 2)[0]); err != "" || !equal {
			t.Fatal("phi alias diverged", err)
		}
		if err := RestorePosition(builder, saved); err != "" {
			t.Fatal(err)
		}
	}
	ir, err := Print(module)
	if err != "" || ir != before || !strings.Contains(ir, "!custom !0") || !strings.Contains(ir, "%p = phi") {
		t.Fatal(ir, err)
	}
}

func TestSSAMalformedNativeRequestsAreRecoverableAndAtomic(t *testing.T) {
	ctx := NewContext()
	defer Close(ctx)
	must := checked(t)
	module := must(Parse(ctx, "define i64 @f(i64 %x) { entry: br label %end\nend: %p = phi i64 [%x, %entry]\n ret i64 %p }", false))
	fn := must(Function(module, "f", nil, false))
	blocks := readIR(t, fn, 1)
	phi := readIR(t, blocks[1], 2)[0]
	x := must(Parameter(fn, 0))
	builder := must(NewBuilder(ctx))
	before, _ := Print(module)
	failures := []string{
		EditIncoming(phi, -1, 0, nil, nil),
		EditIncoming(phi, 2, 0, nil, nil),
		EditIncoming(phi, 3, 0, []Handle{x}, []Handle{blocks[0]}),
		EditIncoming(phi, 0, 0, []Handle{x}, nil),
		EditIncoming(phi, 1, 0, []Handle{x}, []Handle{fn}),
		EditIncoming(x, 1, 0, nil, nil),
		PositionBefore(builder, x),
		PositionBefore(fn, phi),
		PositionStart(builder, phi),
		RestorePosition(builder, phi),
		ClearPosition(phi),
		EraseInstruction(phi),
		SetOperand(phi, 0, x),
		SetOperand(x, 0, x),
		ReplaceUses(module, x),
	}
	for _, err := range failures {
		if !strings.HasPrefix(err, "argument:") {
			t.Fatalf("expected argument error, got %q", err)
		}
	}
	for query := -1; query <= 9; query++ {
		if _, err := ReadIR(builder, query); !strings.HasPrefix(err, "argument:") {
			t.Fatalf("query %d: %q", query, err)
		}
	}
	if _, err := EntryAlloca(fn, module, "bad"); !strings.HasPrefix(err, "argument:") {
		t.Fatal(err)
	}
	after, _ := Print(module)
	if before != after {
		t.Fatal("rejected edits changed module")
	}
}

func TestModuleReferencedConstantsCannotEscapeTheirModule(t *testing.T) {
	ctx := NewContext()
	defer Close(ctx)
	must := checked(t)
	module := must(Parse(ctx, "@global = global i64 0\ndefine ptr @address() { ret ptr @global }", false))
	fn := must(Function(module, "address", nil, false))
	block := readIR(t, fn, 1)[0]
	ret := readIR(t, block, 2)[0]
	global := readIR(t, ret, 5)[0]
	builder := must(NewBuilder(ctx))
	if err := PositionBefore(builder, ret); err != "" {
		t.Fatal(err)
	}
	integer := must(Primitive(ctx, 1, 64))
	folded := must(Cast(builder, 9, global, integer, "address"))
	pure := must(Constant(integer, 0, 7, 0))
	pureFold := must(Binary(builder, 0, pure, pure, "pure"))
	if err := Close(module); err != "" {
		t.Fatal(err)
	}
	for _, value := range []Handle{global, folded} {
		if _, err := Print(value); !strings.HasPrefix(err, "closed:") {
			t.Fatal("module constant outlived module", err)
		}
	}
	if _, err := Print(pureFold); err != "" {
		t.Fatal("pure constant lost context lifetime", err)
	}
}

func TestConcurrentSSAReadersSerializeWithEditsErasureAndClose(t *testing.T) {
	ctx := NewContext()
	defer Close(ctx)
	must := checked(t)
	module := must(Parse(ctx, "define i64 @f(i64 %x) { entry: %unused = add i64 %x, 1\n ret i64 %x }", false))
	fn := must(Function(module, "f", nil, false))
	block := readIR(t, fn, 1)[0]
	value := readIR(t, block, 2)[0]
	alias := readIR(t, block, 2)[0]
	if value != alias {
		t.Fatal("instruction aliases do not share invalidation state")
	}
	start := make(chan struct{})
	var workers sync.WaitGroup
	for worker := 0; worker < 12; worker++ {
		workers.Add(1)
		go func(worker int) {
			defer workers.Done()
			<-start
			for iteration := 0; iteration < 50; iteration++ {
				var err string
				switch worker % 4 {
				case 0:
					_, err = Print(alias)
				case 1:
					_, err = ReadIR(value, 5)
				case 2:
					err = EraseInstruction(value)
				case 3:
					err = Close(module)
				}
				if err != "" && !strings.HasPrefix(err, "closed:") {
					t.Errorf("unexpected result: %s", err)
					return
				}
			}
		}(worker)
	}
	close(start)
	workers.Wait()
	if !IsClosed(alias) {
		t.Fatal("erased alias remained open")
	}
}
