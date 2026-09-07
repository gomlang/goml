package shim

import (
	"fmt"
	"runtime"
	"testing"
)

var result int64
var escaped func(int64) int64

func Native() func(int64) int64 {
	return func(value int64) int64 { return value + 7 }
}

func capturing(offset int64) func(int64) int64 {
	return func(value int64) int64 { return value + offset }
}

func measure(name string, body func(*testing.B)) {
	value := testing.Benchmark(body)
	fmt.Printf("%s\t%s\t%s\n", name, value.String(), value.MemString())
}

func calls(name string, callback func(int64) int64) {
	measure(name, func(b *testing.B) {
		var value int64
		for i := 0; i < b.N; i++ {
			value = callback(value)
		}
		result = value
	})
}

func Run(callback func(int64) int64, roundtrip func(int64) int64, factory func() func(int64) int64) {
	fmt.Printf("%s %s/%s GOMAXPROCS=%d\n", runtime.Version(), runtime.GOOS, runtime.GOARCH, runtime.GOMAXPROCS(0))
	calls("native-call", Native())
	calls("goml-to-go-call", callback)
	calls("go-to-goml-to-go-call", roundtrip)
	measure("native-create-escape", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			escaped = Native()
		}
	})
	offset := callback(0)
	measure("native-capture-escape", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			escaped = capturing(offset)
		}
	})
	measure("goml-create-escape", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			escaped = factory()
		}
	})
	if result == 0 || escaped(1) != 8 {
		panic("invalid benchmark result")
	}
}
