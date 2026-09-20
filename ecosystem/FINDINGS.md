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

## Fixed compiler regressions

The development compiler now passes all four retained reproducers in `repros/`:

- `unit_identity`: separately allocated `Ref[()]` values remain distinct when
  stored in containers. Empty structs, nested zero-size values and zero-length
  arrays also have stable reference equality and hashing. Graph handles now use
  `Ref[()]` directly and retain cross-graph rejection tests.
- `erased_generic`: free-function type arguments used only in the body survive
  specialization. Unused generic definitions compile, explicit calls and function
  values execute, and generic forwarding works across package interfaces.
- `specialized_static`: constructors in `impl Box[f64]` accept inferred or explicit
  owner arguments. Nested implementations such as `impl[T] Box[Vec[T]]` retain
  their complete owner type in exported interfaces. Ambiguous calls require an
  explicit owner type or expected result.
- `ffi_error_alias`: importing `std::ffi` no longer substitutes its `Error` alias
  for the distinct canonical `std::io::Error` type.

`python3 ecosystem/verify.py` runs these reproducers after the library matrix.
Compiler pipeline fixtures 286–288 and module fixtures 084–087 cover these fixes,
including serialization through independently compiled interfaces. The existing
`TypedCommand[T]`, ndarray module constructor, and SQLite transport package remain
valid API designs, but their original compiler workarounds are no longer required.

## Type inference and standard-library interfaces

- Numeric operations in closures whose expected parameter is an associated-type
  projection can require an explicit parameter annotation. Consumer tests retain
  this syntax. A typed intermediate `Result[T, string]` also resolves `Self` for
  static trait calls before chaining `map_err`.
- Scalar conversion is available through `std::num::{ToFloat, TryToInt}`.
  Integer-to-float rounding and checked float-to-integer conversion are verified
  against native Go across all rounding modes, boundary values and random bit
  patterns. Existing bit-based generator code remains valid.
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

## Standard-library capabilities

`std::io` now provides `Read`, `Write`, `BufRead`, `Close`, bounded reads, copying,
in-memory cursors, limiting and buffered adapters, and standard-stream handles.
TCP, TLS and Linux descriptors implement the stream traits. Partial transfers,
failed flush retries, exact limits, binary stdin and malformed stream counts have
regression coverage. Buffered adapters require serialized shared access and
explicit flushing/closing.

`std::context` provides scoped cancellation, inherited monotonic deadlines and
sleep. Existing TCP/UDP waits compose contexts with legacy cancel tokens and
retain whole-operation deadlines. `std::net` adds DNS and named-host connection;
`std::net::tls` adds certificate-verified clients, custom roots, mutual TLS, ALPN,
version policy, deadlines, and cancellation. Local TLS tests use an ephemeral CA
and server and run under Go's race detector. Active TLS I/O cancellation closes
the connection; TCP readiness cancellation leaves the socket reusable. TLS
keeps local close and cancellation distinct from remote EOF, including when the
native read returns EOF during a concurrent shutdown.

## Serde format boundaries

The standard dynamic model now preserves `Binary`, ordered arbitrary-key `Map`
entries and `Extension` values. Optional extension events have recoverable
unsupported defaults, so ordinary format implementations need not implement them.
MessagePack implements the new events, adds a typed `Extension` wrapper and
round-trips binary, duplicate/arbitrary map keys and opaque extension payloads
through `std::serde::Value` without changing wire types. Its own dynamic model
still supports invalid UTF-8 raw strings and timestamp helpers.

JSON, TOML and Bincode do not acquire an extension wire format. Their checked
APIs reject extension values. TOML retains the previous integer-array and
key/value-pair-array projections for binary and map events. The legacy infallible JSON projection is explicitly
lossy; use checked conversion when representation matters.

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
