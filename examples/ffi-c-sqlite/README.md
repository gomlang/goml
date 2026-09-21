# SQLite through generated C bindings

Install SQLite development headers and its linkable library, Go 1.26, Clang, and a C compiler. Then run:

```sh
../../stage2/bin/goml bind-c bindings.json
../../stage2/bin/goml test
../../stage2/bin/goml run
```

The program opens an in-memory database, prepares `SELECT 6 * 7`, reads `42`, and explicitly finalizes the statement and closes the database. C status codes and pointer outputs remain separate values. The SQL tail is copied while the input allocation remains alive.

For dependencies outside system paths, set unquoted `CGO_CPPFLAGS=-I/path/to/include` and `CGO_LDFLAGS=-L/path/to/lib` consistently for generation and builds. Generated sources were validated with SQLite 3.50.4 headers; regenerate against your installed headers. See [C bindings](../../docs/ffi/bind-c.md).
