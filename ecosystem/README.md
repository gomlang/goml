# GoML ecosystem libraries

This directory exercises GoML through reusable libraries, each with its own
`goml.toml`, public API, documentation, and external tests. Implementations are
independent GoML modules. SQLite additionally uses an explicit Go adapter.

All thirteen libraries have implementations, public documentation, independent
consumers and executable verification. The table records their implemented scope;
individual READMEs describe API semantics and limits.

| Module | Functional target | Status |
| --- | --- | --- |
| [parser](parser/README.md) | Combinators, recursive grammars, text and binary primitives, spans, contextual errors, bounded repetition and operator precedence | Implemented; module and consumer tests pass |
| [proptest](proptest/README.md) | Composable generators, associated-type strategies, shrinking, bounded rejection, deterministic replay, collections and recursive data | Implemented; module and consumer tests pass |
| [cli](cli/README.md) | Explicit command schema, options, positional arguments, subcommands, help, validation and third-party `Args` derive | Implemented; module and consumer tests pass |
| [msgpack](msgpack/README.md) | MessagePack wire types, direct Serde integration, typed and dynamic APIs, malformed-input limits and interoperability | Implemented; 14 library tests, consumer checks and 2,490 reference interoperability cases pass |
| [graph](graph/README.md) | Mutable directed/undirected graphs, stable IDs, traversal, components, topological order, shortest paths and spanning trees | Implemented; independent algorithm checks and consumer tests pass |
| [template](template/README.md) | Expressions, lexical scopes, conditions, loops, filters, includes, inheritance, escaping and contextual diagnostics | Implemented; 11 library tests, 2 consumer tests and 1,367 Jinja shared-syntax comparisons pass |
| [redis](redis/README.md) | RESP2/3 codec, typed commands, pipelining, transactions, Pub/Sub, cancellation, timeout and connection lifecycle | Implemented; 15 library tests and race checks, versioned consumer, 2,391 protocol cases and Redis 7.2.5 RESP2/3 interoperability pass |
| [pipeline](pipeline/README.md) | Lazy streams, bounded parallel transforms, filtering, ordering, batching/windows, merge/zip, backpressure and cancellation | Implemented; 21 library tests and race checks, versioned consumer and 1,253 Python oracle cases pass |
| [ndarray](ndarray/README.md) | Generic shared views, slicing, broadcasting, checked arithmetic, reductions, batched multiplication, LU/Cholesky/QR solves and SIMD | Implemented; 17 library tests, versioned consumer, 2,929 NumPy cases and native/SSE2/scalar builds pass |
| [sqlite](sqlite/README.md) | Typed binding/rows, prepared statements, streaming queries, nested savepoints, rollback, cancellation and explicit resource management | Implemented; 13 GoML tests, 4 native tests, versioned consumer, 2,754 SQLite comparisons and race checks pass |
| [lsp](lsp/README.md) | JSON-RPC framing, protocol types, document synchronization, UTF-16 positions, request lifecycle and dispatch | Implemented; module, consumer and subprocess interoperability tests pass |
| [markdown](markdown/README.md) | Block and inline parsing, AST, HTML rendering, escaping, links, code, lists and reference conformance | Implemented; 652/652 CommonMark examples, entity, module and consumer checks pass |
| [diff](diff/README.md) | Sequence and text differences, unified patches, checked application, context and newline preservation | Implemented; tests and GNU interoperability pass |

Validation includes module-local public API tests, separate consuming modules,
deterministic negative cases, reference interoperability where applicable, and
fresh/cached builds. Build products belong in `_artifact/` and are ignored by the
repository. Source code follows the root repository guidelines.

Current development command, from a module directory:

```sh
../../stage2/bin/goml fmt
../../stage2/bin/goml test
```

## Tools

[`goml_stats`](goml_stats/README.md) is a standalone GoML project statistics tool
with no third-party dependencies. It counts files, code/comment/blank lines,
bytes, test files, modules and package directories, with exclusions, detailed
tables and JSON output. Its lexer-aware counting handles raw and multiline
strings without mistaking their contents for comments.

```sh
(cd ecosystem/goml_stats && ../../stage2/bin/goml build)
ecosystem/goml_stats/_artifact/bin/cmd/goml_stats/goml_stats .
python3 ecosystem/goml_stats/verify.py
```

## Library verification

Run the available library and independent consumer checks from the repository root:

```sh
python3 ecosystem/verify.py
python3 ecosystem/verify.py lsp markdown diff
```

With no module arguments, the verifier checks all thirteen libraries and their
consumers. A missing module or failed check is an error. It creates an isolated,
content-addressed registry snapshot under `ecosystem/_artifact/`, leaving the
user's registry untouched. Consumers resolve normal versioned dependencies from
that snapshot. Verification logs and command timings are written under
`ecosystem/_artifact/verification/`.

Each implemented library has a README describing its API, semantics, limits and
tests. [FINDINGS.md](FINDINGS.md) records language capabilities and compiler/API boundaries discovered during this work.

SQLite also requires its declared native Go dependencies to be fetched before
readonly compilation (`cd ecosystem/sqlite && go mod download all`). The NumPy
reference check uses CPython 3.12 on Linux amd64. Reference programs and wheels
are downloaded into ignored `_artifact/` directories as documented by each
library; race checks require the repository's C compiler prerequisite.

Compiler limitations found during implementation remain documented in
[FINDINGS.md](FINDINGS.md), with intentionally failing reproducers excluded from
the passing library matrix. This work does not modify the compiler or standard
library to conceal those boundaries.
