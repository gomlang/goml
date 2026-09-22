package native

import (
	"runtime"
	"unsafe"
)

func SharesTextStorage(source, candidate string) bool {
	if len(source) == 0 || len(candidate) == 0 {
		return false
	}
	start := uintptr(unsafe.Pointer(unsafe.StringData(source)))
	position := uintptr(unsafe.Pointer(unsafe.StringData(candidate)))
	shared := position >= start && position-start < uintptr(len(source))
	runtime.KeepAlive(source)
	runtime.KeepAlive(candidate)
	return shared
}
