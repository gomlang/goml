# FFI Error alias affects an unrelated I/O error

This intentionally failing module imports both `std::ffi` and `std::io`.
`goml check` resolves the error from `io::read_stdin_to_string()` as `go[].error`
and cannot find its `to_string` method, although the actual function returns
`std::io::Error`. Replacing the error callback with a constant allows checking
but linking reports an ANF type mismatch between those two result types.

```sh
cd ecosystem/repros/ffi_error_alias
../../../stage2/bin/goml check
```

The SQLite consumer keeps standard I/O in a separate `transport` package whose
public function returns `Result[string, string]`. That package does not import
the FFI dependency, and its callers do not need to resolve `std::io::Error` in
the FFI environment. The consumer checks, links and runs with this arrangement.
The compiler limitation remains; this reproducer is excluded from the passing
verification matrix.
