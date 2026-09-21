# LLVM through generated C bindings

This module builds and verifies a function returning `i32 42`, then copies its printed LLVM IR and releases the native message. It exercises LLVM opaque handles, enum arguments, output strings and explicit parent/child destruction without handwritten Go adapters.

The configuration selects LLVM 18 headers under `/usr/lib/llvm-18` and dynamically loads `libLLVM-18.so`; adjust paths for another installation. It uses GoML's own C ABI runtime on Linux amd64 with glibc 2.34+, Go 1.26.x and Clang. No cgo or third-party FFI library is used:

```sh
CGO_ENABLED=0 ../../stage2/bin/goml bind-c bindings.json
CGO_ENABLED=0 ../../stage2/bin/goml test
CGO_ENABLED=0 ../../stage2/bin/goml run
```

`TypeArray::null()` is used only for a zero-parameter function signature. This example does not provide array storage or lifetime-checked LLVM wrappers. The higher-level ecosystem LLVM package remains responsible for safe editing and shared lifetime state. See [C bindings](../../docs/ffi/bind-c.md).
