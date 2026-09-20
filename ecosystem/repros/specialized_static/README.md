# Specialized associated constructors

This regression module now checks, builds and runs successfully. Constructors
in `impl Box[f64]` accept both `Box::from_float(...)` with an expected result and
`Box::[f64]::from_float(...)`. Compiler fixtures additionally cover nested generic
owners and cross-package interfaces.

```sh
cd ecosystem/repros/specialized_static
../../../stage2/bin/goml run
```

The complete ecosystem verifier includes this reproducer. Ndarray now exposes
`Array::[f64]::linspace` and inferred `Array::linspace`, with its existing
module-level `linspace` retained as a forwarding convenience API.
