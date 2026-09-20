package adapter

import (
	"runtime"
	"sync/atomic"
)

var observed atomic.Uint64

func Consume(value uint64) {
	observed.Store(value)
	runtime.KeepAlive(value)
}

func Observed() uint64 { return observed.Load() }

func Opaque(value uint64) uint64 {
	var cell atomic.Uint64
	cell.Store(value)
	result := cell.Load()
	runtime.KeepAlive(&cell)
	return result
}
