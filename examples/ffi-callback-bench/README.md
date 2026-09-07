# Callback invocation and allocation benchmark

From the repository root:

```sh
just make
cd examples/ffi-callback-bench
../../stage2/bin/goml run --ffi-check required
```

The module uses Go's testing.Benchmark and its standard calibration. It prints the Go version, platform, GOMAXPROCS, time per operation and allocation statistics. Compilation and FFI checking happen before measurement. Run on an otherwise idle machine and repeat measurements when comparing compiler changes; the benchmark is intentionally excluded from timing-based CI assertions.

`native-call` invokes an ordinary Go function through the same function-parameter loop as the adapted cases. `goml-to-go-call` invokes a converted GoML closure with an immutable captured offset. `go-to-goml-to-go-call` invokes a native Go function after conversion to a GoML closure and back. These three cases construct the function before starting the benchmark; they measure invocation rather than construction.

`native-create-escape` assigns a non-capturing Go function to a global sink. `native-capture-escape` creates a Go closure with a runtime offset and assigns it to that sink. `goml-create-escape` calls an already converted factory that constructs a capturing GoML closure and converts it to a Go function each iteration. That case includes the factory callback invocation as well as closure and adapter creation. The global sink forces the resulting function to escape; it is benchmark instrumentation, not part of the adapter implementation. Native baselines retain ordinary Go optimizer opportunities, so differences are end-to-end observations rather than an isolated instruction cost for one conversion step.

One local run on an Intel Core i7-13700K with go1.25.8, linux/amd64 and GOMAXPROCS=22 produced:

| Case | ns/op | B/op | allocs/op |
| --- | ---: | ---: | ---: |
| native-call | 0.8054 | 0 | 0 |
| goml-to-go-call | 1.564 | 0 | 0 |
| go-to-goml-to-go-call | 2.147 | 0 | 0 |
| native-create-escape | 0.3336 | 0 | 0 |
| native-capture-escape | 8.918 | 16 | 1 |
| goml-create-escape | 22.13 | 32 | 2 |

The invocation cases measured zero allocations in this run. Escaping GoML callback creation measured two allocations, so the adapter API makes no general zero-allocation promise. These are single-run values from a shared development machine, not statistically established performance bounds. The module has no third-party Go dependencies.
