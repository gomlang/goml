# SQLite through generated C bindings

Install SQLite development headers and `libsqlite3.so.0`, Go 1.26.x and Clang on Linux amd64 with glibc 2.34+. This example uses GoML's own dynamic C ABI runtime without cgo or a third-party FFI library:

```sh
CGO_ENABLED=0 ../../stage2/bin/goml bind-c bindings.json
CGO_ENABLED=0 ../../stage2/bin/goml test
CGO_ENABLED=0 ../../stage2/bin/goml run
```

The program opens an in-memory database, prepares `SELECT 6 * 7`, reads `42`, and explicitly finalizes the statement and closes the database. C status codes and pointer outputs remain separate values. The SQL tail is copied while the input allocation remains alive.

For dependencies outside system paths, set unquoted `CGO_CPPFLAGS=-I/path/to/include` for generation and checks, and `LD_LIBRARY_PATH=/path/to/lib` when running. An absolute library path can also be configured. Generated sources were validated with SQLite 3.50.4 headers; regenerate against your installed headers. See [C bindings](../../docs/ffi/bind-c.md).
