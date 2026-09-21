package cabi

import (
	"bytes"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

func TestProfilingAndTracingNativeCalls(t *testing.T) {
	lib := probeLibrary(t)
	fn := address(t, lib, "probe_spin")
	var profile, events bytes.Buffer
	if err := pprof.StartCPUProfile(&profile); err != nil {
		t.Fatal(err)
	}
	defer pprof.StopCPUProfile()
	if err := trace.Start(&events); err != nil {
		t.Fatal(err)
	}
	defer trace.Stop()
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			f := Frame{Function: fn, Integer: [6]uint64{10000000}}
			f.Call()
			if f.Result != 49999995000000 {
				t.Errorf("spin: %d", f.Result)
			}
			runtime.GC()
		}()
	}
	wg.Wait()
	trace.Stop()
	pprof.StopCPUProfile()
	if profile.Len() == 0 || events.Len() == 0 {
		t.Fatal("missing runtime observability output")
	}
}

func probeLibrary(t *testing.T) *Library {
	t.Helper()
	library := filepath.Join(t.TempDir(), "probe.so")
	command := exec.Command("cc", "-shared", "-fPIC", "-O0", "-o", library, "testdata/probe.c")
	if output, err := command.CombinedOutput(); err != nil {
		t.Fatalf("C fixture: %s: %v", output, err)
	}
	return &Library{Names: []string{library}}
}

func address(t *testing.T, lib *Library, name string) uint64 {
	t.Helper()
	result, err := lib.Symbol(name)
	if err != nil {
		t.Fatal(err)
	}
	return result
}

func TestRegistersAndStackArguments(t *testing.T) {
	lib := probeLibrary(t)
	f := Frame{Function: address(t, lib, "probe_integers"), Integer: [6]uint64{1, 2, 3, 4, 5, 6}, StackCount: 2, Stack: [128]uint64{7, 8}}
	f.Call()
	if f.Result != 204 {
		t.Fatal(f.Result)
	}
	f = Frame{Function: address(t, lib, "probe_mixed"), Integer: [6]uint64{1, 4, 6, 8, 10, 12}, Float: [8]uint64{math.Float64bits(2), uint64(math.Float32bits(3)), math.Float64bits(5), math.Float64bits(7), math.Float64bits(9), math.Float64bits(11), math.Float64bits(13), math.Float64bits(15)}, StackCount: 4, Stack: [128]uint64{14, 16, math.Float64bits(17), uint64(math.Float32bits(18))}}
	f.Call()
	if value := math.Float64frombits(f.FloatResult); value != 171 {
		t.Fatal(value)
	}
	f = Frame{Function: address(t, lib, "probe_float"), Float: [8]uint64{uint64(math.Float32bits(3))}}
	f.Call()
	if value := math.Float32frombits(uint32(f.FloatResult)); value != -4.5 {
		t.Fatal(value)
	}
	f = Frame{Function: address(t, lib, "probe_stack"), Integer: [6]uint64{32}}
	f.Call()
	if f.Result != 528 {
		t.Fatal(f.Result)
	}
}

func TestThreadLocalStorageAndThreadExit(t *testing.T) {
	lib := probeLibrary(t)
	fn := address(t, lib, "probe_tls")
	var wg sync.WaitGroup
	for i := 1; i <= 32; i++ {
		wg.Add(1)
		go func() {
			runtime.LockOSThread()
			defer wg.Done()
			for j := 0; j < 5; j++ {
				f := Frame{Function: fn, Integer: [6]uint64{uint64(i)}}
				f.Call()
				if f.Result != 1 {
					t.Errorf("TLS/errno changed for %d", i)
				}
				runtime.GC()
			}
		}()
	}
	wg.Wait()
}

func TestBlockingCallReleasesProcessor(t *testing.T) {
	lib := probeLibrary(t)
	fn := address(t, lib, "probe_block")
	old := runtime.GOMAXPROCS(1)
	defer runtime.GOMAXPROCS(old)
	read, write, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer read.Close()
	defer write.Close()
	entered, err := Alloc(4)
	if err != nil {
		t.Fatal(err)
	}
	defer Free(entered)
	done := make(chan uint64, 1)
	go func() {
		f := Frame{Function: fn, Integer: [6]uint64{uint64(uintptr(entered)), uint64(read.Fd())}}
		f.Call()
		done <- f.Result
	}()
	deadline := time.Now().Add(5 * time.Second)
	for atomic.LoadUint32((*uint32)(entered)) == 0 {
		if time.Now().After(deadline) {
			t.Fatal("C call did not start")
		}
		time.Sleep(time.Millisecond)
	}
	runtime.GC()
	if _, err := write.Write([]byte{42}); err != nil {
		t.Fatal(err)
	}
	select {
	case value := <-done:
		if value != 42 {
			t.Fatal(value)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("C call did not finish")
	}
}

func TestLoaderErrors(t *testing.T) {
	for _, lib := range []*Library{{Names: []string{"/goml-missing-library.so"}}, {Names: []string{"libc.so.6"}}} {
		if _, err := lib.Symbol("goml_missing_symbol"); err == nil {
			t.Fatal("missing loader diagnostic")
		}
	}
	lib := &Library{Names: []string{"libc.so.6"}}
	if _, err := lib.Symbol("x\x00y"); err == nil {
		t.Fatal("accepted NUL")
	}
	p, err := Copy([]byte{255, 'a'}, true)
	if err != nil {
		t.Fatal(err)
	}
	defer Free(p)
	if _, err := String(p, 1); err == nil || !strings.Contains(err.Error(), "limit") {
		t.Fatal(err)
	}
	if value, err := String(p, 2); err != nil || value != "\xffa" {
		t.Fatalf("%q %v", value, err)
	}
}

func TestLibc(t *testing.T) {
	lib := &Library{Names: []string{"libc.so.6"}}
	address, err := lib.Symbol("strlen")
	if err != nil {
		t.Fatal(err)
	}
	var wg sync.WaitGroup
	for i := 0; i < 64; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				p, err := Copy([]byte("goml"), true)
				if err != nil {
					t.Error(err)
					return
				}
				f := Frame{Function: address}
				f.Integer[0] = uint64(uintptr(p))
				f.Call()
				Free(p)
				if f.Result != 4 {
					t.Errorf("strlen: %d", f.Result)
					return
				}
				if j%20 == 0 {
					runtime.GC()
				}
			}
		}()
	}
	wg.Wait()
	if err := os.Setenv("GOML_CABI_TEST", "works"); err != nil {
		t.Fatal(err)
	}
	defer os.Unsetenv("GOML_CABI_TEST")
	getenv, err := lib.Symbol("getenv")
	if err != nil {
		t.Fatal(err)
	}
	p, err := Copy([]byte("GOML_CABI_TEST"), true)
	if err != nil {
		t.Fatal(err)
	}
	defer Free(p)
	f := Frame{Function: getenv}
	f.Integer[0] = uint64(uintptr(p))
	f.Call()
	value, err := String(Pointer(f.Result), 100)
	if err != nil || value != "works" {
		t.Fatalf("getenv: %q %v", value, err)
	}
}

func TestFrameLayout(t *testing.T) {
	f := Frame{}
	if unsafe.Offsetof(f.Integer) != 8 || unsafe.Offsetof(f.Float) != 56 || unsafe.Offsetof(f.StackCount) != 120 || unsafe.Offsetof(f.Stack) != 128 || unsafe.Offsetof(f.Result) != 1152 || unsafe.Offsetof(f.FloatResult) != 1160 {
		t.Fatal("assembly frame layout changed")
	}
}
