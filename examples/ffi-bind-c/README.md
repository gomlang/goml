# Calling C from GoML

This module exercises typed C handles, checked scalar and buffer arguments, copied strings, output parameters, explicit release, and compile-time constants. `sample.h` is a self-contained C library. No handwritten Go bridge is needed.

With the repository toolchain built and Go 1.26, Clang and a C compiler available:

```sh
../../stage2/bin/goml bind-c bindings.json
../../stage2/bin/goml test
../../stage2/bin/goml run
```

Output is `42`, `GoML calls C`, and `255`. The driver verifies C inputs before building. `bindings/generated.gom`, `native/generated.go` and `bindings.json.goml-c-bind.json` are generated together. See [C binding configuration](../../docs/ffi/bind-c.md) for the lifetime contract and supported parameter mappings.
