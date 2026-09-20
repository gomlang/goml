# Concrete specialization with an associated constructor

This intentionally failing module isolates a method-lookup boundary on the
current development toolchain:

```sh
cd ecosystem/repros/specialized_static
../../../stage2/bin/goml check
```

For `impl Box[f64]`, `Box::from_float(...)` reports that the method is not found,
even with an explicitly typed result. `Box::[f64]::from_float(...)` finds a method
but reports that it expects zero owner type arguments. This reproducer is
excluded from the passing ecosystem verification matrix.

The ndarray library exposes its specialized constructor as the module function
`ndarray::linspace(...)`. Ordinary instance methods in `impl Array[f64]` work
across the versioned dependency boundary and are covered by consumer tests.
