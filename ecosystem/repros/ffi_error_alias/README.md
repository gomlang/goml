# FFI and I/O error type identities

This regression module now checks, builds and runs successfully while importing
both `std::ffi` and `std::io`. Qualified type lookup keeps `io::Error` distinct
from the FFI `Error` alias through checking, interface serialization and linking.

```sh
cd ecosystem/repros/ffi_error_alias
../../../stage2/bin/goml run </dev/null
```

The complete ecosystem verifier includes this reproducer. SQLite's consumer now
uses standard I/O directly; its former transport workaround has been removed.
