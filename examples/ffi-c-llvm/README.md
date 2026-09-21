# LLVM through generated C bindings

This module builds and verifies a function returning `i32 42`, then copies its printed LLVM IR and releases the native message. It exercises LLVM opaque handles, enum arguments, output strings and explicit parent/child destruction without handwritten Go adapters.

The configuration selects LLVM 18 headers and `libLLVM-18` under `/usr/lib/llvm-18`; adjust paths for another installation. With LLVM 18 development files, Go 1.26, Clang and a C compiler installed:

```sh
../../stage2/bin/goml bind-c bindings.json
../../stage2/bin/goml test
../../stage2/bin/goml run
```

`TypeArray::null()` is used only for a zero-parameter function signature. This example does not provide array storage or lifetime-checked LLVM wrappers. The higher-level ecosystem LLVM package remains responsible for safe editing and shared lifetime state. See [C bindings](../../docs/ffi/bind-c.md).
