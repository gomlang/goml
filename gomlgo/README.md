# gomlgo

`gomlgo` is an independent GoML module implementing a Go 1.26 scanner, recursive-descent parser, package type checker, and source importer. The compatibility target is Go 1.26.5 at `/usr/lib/go-1.26`. Production code does not call `go/scanner`, `go/parser`, `go/types`, `go/constant`, or `go/importer`.

The frontend accepts source as `std::bytes::Bytes`, preserving invalid UTF-8 and physical byte offsets. The scanner implements the complete Go token set, literal validation, comments, BOM and NUL handling, Unicode identifiers, and automatic semicolon insertion. Tokens retain kind, literal, start/end offsets, and whether a semicolon is synthetic.

The parser is organized as declarations/files, types, expressions, statements, comments, and build constraints. It uses a comment-skipping token cursor, explicit recovery points, progress guards, a diagnostic cap, and a nesting guard. Its normalized AST mirrors every Go 1.26 `go/ast` expression, statement, declaration, field, comment, and file node used by `parser.ParseFile` with `ParseComments | AllErrors | SkipObjectResolution`. Positions are 0-based physical byte offsets and `-1` represents `token.NoPos`.

The type checker lowers the parser AST into a strongly typed syntax arena with stable file and node IDs. Types, objects, packages, scopes, constants, selections, instances, and initialization dependencies use ID-based arenas and side tables. It supports package and file scopes, declarations, imports, named and alias types, interfaces and method sets, constants, expressions and builtins, generics and inference, statements and control flow, initialization ordering, and Go 1.26 `new(expr)` behavior. Constants use arbitrary-precision integers and rationals plus a 512-bit binary floating representation for values that exceed the exact rational resource boundary.

Syntax paths are opaque, immutable handles. Use `node.path.to_string()` to render them and `syntax::empty_path()` or `syntax::text_path(value)` when constructing syntax nodes.

Type facts retain independent constant snapshots. `FactConstant::SmallInt(i64)` stores signed 64-bit integers directly; larger values retain copied arbitrary-precision storage. Use `type_fact_constant` and `type_fact_constant_kind` to read the exact value and kind across storage variants.

The source importer evaluates `//go:build` and `// +build` constraints, GOOS/GOARCH and release tags, custom tags, cgo selection, platform filename suffixes, import cycles, and source dependencies. It selects ordinary, internal-test, and external-test package sources separately. Its default target is `linux/amd64`, cgo disabled, with `/usr/lib/go-1.26/src` as the source root.

`checker::CheckConfig` exposes `go_version`, target `sizes`, `ignore_func_bodies`, `fake_import_c`, `disable_unused_import_check`, `enable_alias`, and explicit import availability/failure inputs. The default is Go 1.26, gc/amd64 sizes, function bodies enabled, cgo import emulation disabled, unused-import checks enabled, and materialized aliases.

`oracle/` is test-only Go code built with `/usr/lib/go-1.26/bin/go`. It exposes `version`, `scan`, `parse-file`, `parse-expr`, `type-check-package`, and `constant-eval` over stdin/stdout JSON. Type-check requests contain a package path, Go version, target, source files, and checker configuration; responses contain diagnostics, package metadata, normalized scopes and objects, and all `go/types.Info` sections. The test protocol also has `collect_info`; acceptance-only GOROOT runs set it to false so the oracle does not normalize data that will not be compared. GoML differential runners invoke the oracle only as an independent reference process.

## Commands

From the repository root:

```bash
just gomlgo-test
just gomlgo-build
just gomlgo-diff
just gomlgo-type-test
just gomlgo-type-diff
just gomlgo-type-diff-official
just gomlgo-type-diagnostics-official
just gomlgo-type-diff-goroot 10 std
just gomlgo-type-diff-goroot -1 cmd acceptance
just gomlgo-type-diff-goroot -1 std info
just gomlgo-type-diff-goroot 1 std acceptance linux arm64 0 none
just gomlgo-type-diff-goroot-matrix 16
just gomlgo-type-diff-goroot-tests 16 internal
just gomlgo-type-diff-goroot-tests 16 external
just gomlgo-type-mutate-diff 25
just gomlgo-diff-goroot
just gomlgo-parse-diff-goroot
just gomlgo-ast-shape-goroot
just gomlgo-ast-position-goroot
just gomlgo-ast-comments-goroot
just gomlgo-mutate-diff 25
```

The built CLI supports:

```bash
_artifact/bin/cmd/gomlgo/gomlgo scan FILE
_artifact/bin/cmd/gomlgo/gomlgo parse FILE
_artifact/bin/cmd/gomlgo/gomlgo parse-expr 'a + b*c'
_artifact/bin/cmd/gomlgo/gomlgo check FILE...
_artifact/bin/cmd/gomlgo/gomlgo check-package DIRECTORY
_artifact/bin/cmd/gomlgo/gomlgo run FILE [-- PROGRAM_ARGS...]
```

## Single-file interpreter

`gomlgo run` accepts exactly one explicit `.go` file. The file must declare `package main` and define `main.main`. Only that file is interpreted: other files in the directory, local packages, third-party modules, cgo, and assembly are outside the current execution scope. Imports are limited to Go 1.26 standard-library packages.

The execution loader uses the selected Go 1.26 toolchain for environment and standard-library dependency metadata. User functions are lowered to typed bytecode and executed by the GoML VM; `gomlgo run` is not a wrapper around `go run`. The VM supports package initialization, functions and closures, control flow, arrays, structs, pointers, slices, maps, methods, interfaces, type assertions and switches, method values and expressions, variadic calls, generic function and method instantiation, string, byte-slice, rune-slice, and rune conversions, and `defer`, `panic`, and `recover`. Runtime faults use the same panic unwinding path and therefore execute deferred calls and can be recovered. A cooperative scheduler implements goroutines, buffered and unbuffered channels, close, channel range, seeded select, and deadlock detection. `--seed` selects deterministic select choices and `--max-goroutines` limits scheduler growth.

Generated files named `.gom.go` or `goml_generated.go` use the same standard-library source declarations as ordinary `.go` files. Their filenames do not select alternate type definitions. Native image builds sharing a cache key are serialized with a process lock; failed builds release the lock and remove their staging files.

Standard-library bodies are never interpreted. Used package functions, including concrete generic instances, are linked to stable native call IDs and a generated Go registry. Generic standard-library calls inside interpreted generic functions are specialized with the outer function's concrete type arguments. The execution image places the VM and registry in one process and passes opaque values that directly contain `reflect.Value`; there is no RPC, pipe protocol, serialization, session, or remote object-handle table.

Exported standard-library named structs are registered with their real Go types, so their zero values and addressable pointer receivers retain Go semantics. Used standard-library methods are bound as static method expressions, and function values returned by native calls are invoked directly through the same asynchronous `reflect.Value` bridge. Native slice results preserve both length and capacity. This covers compiler-generated task-scope code using `sync.Mutex`, `sync.WaitGroup`, `context.Context`, and `context.CancelFunc`. Checked-in `gomlc` products for lexical task scope, file I/O, JSON parsing and encoding, and bytes, UTF-8, vector, and collection operations execute with output identical to Go 1.26.

Native calls run on host goroutines and wake the VM through a completion queue. User closures cross the boundary as direct callback tokens and execute on the VM scheduler, preserving captures, panic, and exit behavior. The image generator emits concrete proxies for exported standard-library interfaces and type shells for exported, non-generic user structs with supported fields, including struct tags. Native channels such as values returned by `time.After` participate in blocking receive and receive-only select. `os.Args` contains the image path followed only by arguments after `--`.

The native value boundary currently supports scalar values, recursively supported slices, concrete function callbacks, selected exported interfaces, non-generic user structs whose fields use supported value types, and receive-capable native channels. Unsupported package variables other than `os.Args`, interface methods with generic or variadic signatures, native channel send/mixed native-and-interpreted select, user aggregate pointers requiring persistent writeback identity, `unsafe`, cgo, and third-party packages are diagnosed rather than approximated. Native code has the same operating-system authority as an ordinary Go program, so the interpreter is not a security sandbox.

The parser differential binary accepts `parse-acceptance`, `ast-shape`, `ast-position`, `ast-comments`, or `diagnostic-strict`. A mismatch reports the file, normalized AST path, surrounding node offset, expected value, and actual value.

Mutation differential failures are stored under `_artifact/gomlgo-diff/failures/<hash>/` with `input.go`, `gomlgo.json`, `oracle.json`, and `diff.txt`.

Type-check mutation failures are stored under `_artifact/gomlgo-type-diff/failures/<hash>/` with `request.json`, `sources/`, `gomlgo.json`, `oracle.json`, and `diff.txt`.

The GOROOT type differential recipes run their checker in a user systemd scope with `MemoryHigh=6G`, `MemoryMax=8G`, `MemorySwapMax=1G`, `TasksMax=512`, and `CPUQuota=300%`. They accept GOOS, GOARCH, cgo, and `none|internal|external` test-package modes, and reset the loader session every 16 target packages. `acceptance` compares accept/reject status and diagnostics needed to determine it; `info` additionally compares normalized `TypeInfo`.

The official type corpus recipes also run in a resource-limited user systemd scope. They invoke one checker process per source file so retained dependency state is released between cases; the hard limits are 10 GiB memory, 2 GiB swap, 256 tasks, and 300% CPU.

`testdata/invalid/` contains recovery smoke cases. `testdata/regressions/` contains minimized differential failures and is checked by `just gomlgo-diff`.

## Current parity

Against all 7,710 `.go` files under `/usr/lib/go-1.26/src`:

- Scanner: 17,255,942 tokens, zero kind/literal/start/end/synthetic mismatches.
- Parse acceptance: zero mismatches, with 7,666 accepted and 44 rejected files.
- Valid AST shape and important data: zero mismatches.
- Valid AST byte positions: zero mismatches.
- Comments, doc/trailing attachment, `File.Comments`, and `GoVersion`: zero mismatches.
- Mutation sample: 25 valid seeds and 272 generated cases had zero acceptance mismatches; 8 cases differed only in the first diagnostic byte offset.
- Local type-check fixtures: 3 files, with 2 accepted and 1 rejected, have zero acceptance differences; accepted files have complete normalized TypeInfo parity and the rejected fixture has exact diagnostic parity.
- Official Go type-check corpus: 363 files have zero acceptance and exact-diagnostic differences. The result is 83 accepted and 280 rejected files; all 83 accepted files have complete normalized package, type, Defs, Uses, Implicits, Selections, Scopes, Instances, InitOrder, FileVersions, Universe, and unsafe parity, and all 280 rejected files have partial normalized TypeInfo parity.
  - `internal/types/testdata/check`: 79 files, 13 accepted and 66 rejected.
  - `internal/types/testdata/spec`: 12 files, 1 accepted and 11 rejected.
  - `internal/types/testdata/examples`: 8 files, all rejected.
  - `internal/types/testdata/fixedbugs`: 262 files, 68 accepted and 194 rejected.
  - `go/types/testdata/local`: 2 files, 1 accepted and 1 rejected.
- Type-check mutation smoke corpus: 2 accepted seeds and 9 semantic mutations have zero acceptance differences.
- GOROOT `std` on linux/amd64 with cgo disabled: 339 selected packages, 336 accepted and 3 rejected, with zero acceptance differences. Complete normalized TypeInfo matches for 335 accepted source packages; the remaining accepted package is the built-in `unsafe` package.
- GOROOT `cmd/...` on linux/amd64 with cgo disabled: 210 selected packages, 188 accepted and 22 rejected, with zero acceptance differences.
- Targeted extension smoke checks have zero acceptance differences for `internal/goarch` on linux/arm64, `internal/goos` on windows/amd64, `runtime/cgo` with cgo enabled, the internal tests of `bytes`, and the external test package `cmp_test`.

Strict parser diagnostic count and message parity for invalid files is still best-effort. Complete `cmd/...` Info, the full cgo-enabled corpus, broad internal/external test-package corpora, and full GOOS/GOARCH matrices have not been established as zero-diff.
