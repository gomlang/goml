# Language findings from ecosystem development

These observations come from compiling and running actual library consumers.
They distinguish supported designs from current compiler or API boundaries.

## Verified capabilities

- Generic parser combinators retain captured callbacks and compose across
  independently resolved modules. Text and binary parsers report checked input
  errors, and parser operators preserve left/right associativity.
- Associated-type strategies support custom downstream implementations. Lazy
  recursive shrink trees, dependent generators and higher-order transformations
  compile and execute across module boundaries.
- Third-party derives export through dependency interfaces. CLI derives generate
  generic struct decoding, unit-enum conversion, definition-site runtime helper
  calls and type predicates. Import aliases and caller-name collisions are tested.
- Generic sequence algorithms produce minimal edits over exhaustive short inputs;
  the text layer interoperates with GNU diff and GNU patch in both directions.
- Generic graphs support downstream label types, captured weight callbacks,
  derived handle hashing and iterative algorithms. Shortest paths and spanning
  forests agree with independent matrix and exhaustive-subset oracles.
- A generic LSP router deserializes downstream parameter types and serializes
  downstream result types inside stored closures. Its separate consumer also
  validates public stateful APIs, standard I/O and UTF-16 document coordinates
  against an independent Python client.
- The Markdown library combines mutually recursive block/inline trees, shared
  node arenas, captured visitors, Unicode case folding and ordinary Go FFI
  character classifiers. Its generated entity lookup table compiles as GoML;
  the independent consumer matches every CommonMark 0.31.2 reference example.

- MessagePack implements the entire Serde serializer/deserializer event protocol
  in an ordinary dependency. Downstream generic structs and renamed enum variants
  specialize directly into binary encoding/decoding calls. The library also
  exercises all integer widths, floating bit conversions, checked UTF-8, shared
  state handles and bounded incremental binary parsing. Independent reference
  bytes agree across all wire tags and configurable typed representations.

- The template engine compiles recursive expressions and statement trees into
  reusable private ASTs. Rendering combines lexical environment stacks, captured
  application callbacks, independently derived Serde contexts and inherited block
  dispatch. Checked numeric operators and bit-based integer/float comparisons
  agree with the independent Jinja reference on the shared expression dialect.

- Redis exercises native TCP, cancellation-aware channel gates and whole-operation
  deadlines without a Go client adapter. Heterogeneous pipeline tickets retain
  generic reply decoders and captured fallible transformations; consumer-defined
  `FromReply` implementations specialize across the versioned module boundary.
  Incremental recursive RESP values preserve binary data, attributes and streamed
  aggregates. All 15 library tests also pass under Go's race detector, including
  concurrent calls, close wakeups and cancellation during partial replies.

- Pipeline composes lazy generic push functions across module interfaces, with
  separate item/error type parameters and consumer-defined records. Parallel
  workers use admission credits to bound both channel queues and ordered
  reassembly. A shared typed failure channel cancels the root run so downstream
  callbacks wake on an upstream error; local early-stop cancellation preserves
  concatenation and the caller's scope. Accepted task handles are joined before
  reading final errors, preserving failures during cleanup. Independent sequence
  oracles and all 21 tests under the race detector verify the implementation,
  including concurrent reuse, zero-capacity channels and nested scopes.

- Ndarray implements shared strided `Array[T]` views and consumer-defined numeric
  types through `Scalar` defaults. A downstream complex implementation inherits
  batch, sum and dot methods across the registry dependency boundary. Checked
  shape transformations, overlap-safe assignment and LU/Cholesky/Householder QR
  agree with 2,929 independent NumPy cases. Generated metadata and linked symbols
  confirm ten native SIMD kernels in the consumer, five in an SSE2-only build
  and none in a scalar-only build; every build passes the numerical oracle.

- SQLite validates explicit named Go types and raw FFI calls from an independently
  versioned GoML dependency. Opaque native handles preserve shared resource state
  without exposing database/sql's large implementation graph. Downstream value
  and row traits, nullable generic conversions and transaction closures execute
  against real SQLite. Cursor/statement ownership, nested savepoints, automatic
  engine rollback detection, cancellation and close are covered by GoML tests,
  native lock-contention tests and race-detector builds. A separate Python engine
  agrees across 243 sequences and 2,754 operations, including file interchange.

## Zero-size reference identities

`repros/unit_identity` stores separately allocated unit references in a vector.
On the current development toolchain it prints `true` for equality of the two
unit references, and `false` for separately allocated boolean references. Direct
local allocations can give different results, so the vector storage is part of
the reproducer. This also caused real cross-graph handle tests to fail when
graph identity used `Ref[()]`.

```sh
cd ecosystem/repros/unit_identity
../../../stage2/bin/goml run
```

Graph identity now uses `Ref[bool]`. Cross-graph equality and rejected foreign
handles pass for library and independently compiled consumer instances. No
compiler or runtime changes are included in this workaround.

## Erased generic function parameters

`repros/erased_generic` is a minimal reproducer. Its generic function uses `T`
only in its body, returning an ordinary `isize`. The module type-checks, but
linking reports `generic parameter T remained after specialization`, even when
that function is unused. This was reproduced while implementing the CLI schema
accessor, then reduced independently of derives and the CLI library.

```sh
cd ecosystem/repros/erased_generic
../../../stage2/bin/goml check
../../../stage2/bin/goml build
```

The library uses `TypedCommand[T]` as the schema accessor's return type. That
retains the type parameter in the public signature and passes compilation and
consumer tests. The reproducer is intentionally excluded from the normal
passing-library matrix; it is evidence for a future compiler fix.

## Type inference and standard-library interfaces

- Numeric operations in closures whose expected parameter is an associated-type
  projection can require an explicit parameter annotation. Consumer tests retain
  this syntax. A typed intermediate `Result[T, string]` also resolves `Self` for
  static trait calls before chaining `map_err`.
- Scalar integer-to-float methods are not part of the current public API. The
  property generator creates uniform-grid fractions through IEEE-754 bit
  construction and `float64_from_bits`.
- `Debug` is not universally implemented for standard generic containers. Tests
  compare container values through `PartialEq`, while scalar assertions retain
  detailed diagnostics.
- Strings and slices do not expose `is_empty`; use `byte_len() == 0` or
  `len() == 0`. String byte slices must end on UTF-8 boundaries. Binary parsing
  uses byte slices, and text diagnostics clamp to scalar boundaries.
- The implemented derive builders are `meta_expr_int(isize)` and
  `meta_expr_char(string)`; the language guide's abbreviated builder list does not
  precisely match those spellings/signatures. Ecosystem derives use the actual
  exported API, also used by the standard library.

## Serde format boundaries

The standard Serde event protocol includes binary and arbitrary-key map events,
so an external format can support typed `Binary` and `Pairs[K, V]` wrappers without
changing the standard library. It has no extension event. Its separate dynamic
`std::serde::Value` representation also lacks binary, extension and arbitrary-key
map variants. MessagePack therefore supplies its own full-fidelity dynamic value
and timestamp helpers, while keeping its ordinary typed path directly on the
standard event protocol. Unsupported `deserialize_any` cases return errors
instead of silently changing the wire type.

## Associated constructors in concrete generic specializations

`repros/specialized_static` reproduces a lookup limitation for an associated
constructor declared in `impl Box[f64]`. `Box::from_float(...)` reports method
not found even when the result is explicitly `Box[f64]`. The explicit spelling
`Box::[f64]::from_float(...)` instead reports that the method expects zero owner
type arguments. The intentionally failing module is excluded from verification.

```sh
cd ecosystem/repros/specialized_static
../../../stage2/bin/goml check
```

Ndarray exposes `linspace` as a module function. Specialized instance methods,
including `Array[f64]` statistics and factorizations, work across the dependency
boundary and pass the independent consumer tests. No compiler change is included.

## FFI Error alias collides with unrelated package error types

`repros/ffi_error_alias` imports `std::ffi` and `std::io` in one package. Checking
then treats `io::read_stdin_to_string()`'s error as `go[].error`, so its normal
`to_string` call fails. Ignoring the error payload allows checking but linking
reports an ANF mismatch between `Result[string, std::io::Error]` and
`Result[string, go[].error]`. This occurs without a custom native adapter.

```sh
cd ecosystem/repros/ffi_error_alias
../../../stage2/bin/goml check
```

The SQLite consumer reads standard input in a separate `transport` package that
imports only standard I/O and returns `Result[string, string]`. This prevents the
I/O error type from being resolved in the FFI-importing package. That arrangement
passes checking, linking and execution; it does not fix the compiler limitation.
SQLite's public error type is named `DbError`, with an opaque retained native
cause rather than a public raw Go error return.

## FFI metadata graph size

The SQLite adapter initially exposed pointer aliases whose fields reached its
session, connection, database/sql objects and context state. Even after exporting
the pointee names, checking reported `ffi-type: type graph comparison limit
exceeded`. Exported interfaces with private marker methods keep the public
boundary small while their dynamic implementations retain the same native
state. This interface-based adapter passes required FFI checking; the compiler's
graph-comparison limit remains unchanged. Go pointer alias pointees must also
be publicly nameable: an exported alias of a private pointee was rejected before
that graph comparison.
