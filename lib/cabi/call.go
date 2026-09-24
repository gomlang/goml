package cabi

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"unsafe"
)

type Frame struct {
	Function    uint64
	Integer     [6]uint64
	Float       [8]uint64
	StackCount  uint64
	Stack       [128]uint64
	Result      uint64
	FloatResult uint64
}

func (f *Frame) Call() {
	if f.Function == 0 || f.StackCount > uint64(len(f.Stack)) {
		panic("invalid GoML C ABI frame")
	}
	runtimeCall(unsafe.Pointer(&invokeABI), unsafe.Pointer(f))
	runtime.KeepAlive(f)
}

func direct(function *byte, arguments ...uint64) uint64 {
	f := Frame{Function: uint64(uintptr(unsafe.Pointer(function)))}
	copy(f.Integer[:], arguments)
	f.Call()
	return f.Result
}

func Alloc(size int) (unsafe.Pointer, error) {
	if size < 0 || size > 1<<30 {
		return nil, fmt.Errorf("invalid C allocation size")
	}
	if size == 0 {
		return nil, nil
	}
	p := Pointer(direct(&mallocABI, uint64(size)))
	if p == nil {
		return nil, fmt.Errorf("C allocation failed")
	}
	clear(unsafe.Slice((*byte)(p), size))
	return p, nil
}

func Free(p unsafe.Pointer) {
	direct(&freeABI, uint64(uintptr(p)))
}

func String(p unsafe.Pointer, limit int) (string, error) {
	if p == nil {
		return "", nil
	}
	for n := 0; n <= limit; n++ {
		if *(*byte)(unsafe.Add(p, n)) == 0 {
			return string(unsafe.Slice((*byte)(p), n)), nil
		}
	}
	return "", fmt.Errorf("C string exceeds configured copy limit")
}

func Bytes(p unsafe.Pointer, size int) []byte {
	return append([]byte(nil), unsafe.Slice((*byte)(p), size)...)
}

func Copy(data []byte, terminate bool) (unsafe.Pointer, error) {
	size := len(data)
	if terminate {
		size++
	}
	p, err := Alloc(size)
	if err == nil {
		copy(unsafe.Slice((*byte)(p), size), data)
	}
	return p, err
}

type Library struct {
	Names   []string
	once    sync.Once
	handles []uint64
	err     error
	mu      sync.Mutex
	symbols map[string]uint64
}

func loaderCall(function *byte, first uint64, name string) (uint64, error) {
	if strings.IndexByte(name, 0) >= 0 {
		return 0, fmt.Errorf("C loader name contains NUL")
	}
	p, err := Copy([]byte(name), true)
	if err != nil {
		return 0, err
	}
	defer Free(p)
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	direct(&errorABI)
	var result uint64
	if function == &openABI {
		result = direct(function, uint64(uintptr(p)), 2)
	} else {
		result = direct(function, first, uint64(uintptr(p)))
	}
	if message := Pointer(direct(&errorABI)); message != nil {
		value, copyError := String(message, 1<<20)
		if copyError != nil {
			return 0, copyError
		}
		return 0, fmt.Errorf("C loader: %s", value)
	}
	if result == 0 {
		return 0, fmt.Errorf("C loader: null address for %s", name)
	}
	return result, nil
}

func (l *Library) Symbol(name string) (uint64, error) {
	l.once.Do(func() {
		l.symbols = make(map[string]uint64)
		for _, name := range l.Names {
			handle, err := loaderCall(&openABI, 0, name)
			if err != nil {
				l.err = err
				return
			}
			l.handles = append(l.handles, handle)
		}
	})
	if l.err != nil {
		return 0, l.err
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if address, ok := l.symbols[name]; ok {
		return address, nil
	}
	for _, handle := range l.handles {
		address, err := loaderCall(&symbolABI, handle, name)
		if err == nil {
			l.symbols[name] = address
			return address, nil
		}
	}
	return 0, fmt.Errorf("C symbol %s is unavailable in %v", name, l.Names)
}
