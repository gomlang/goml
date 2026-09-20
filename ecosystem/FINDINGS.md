# Language findings from ecosystem development

These observations come from compiling and running actual library consumers.
They distinguish supported designs from current compiler or API boundaries.

## Verified capabilities

- Generic parser combinators retain captured callbacks and compose across
  independently resolved modules. Text and binary parsers report checked input
  errors, and parser operators preserve left/right associativity.
- Associated-type strategies support custom downstream implementations. Lazy
  recursive shrink trees, dependent generators and higher-order transformations
  compile and execute across module boundaries. Bounded campaigns aggregate
  counterexamples under shared shrink budgets and preserve typed initial/minimal
  values alongside categorical distributions. A downstream state-machine helper
  resets custom model/system types for every initial and shrink evaluation, with
  explicit invariant and cleanup results.
- Third-party derives export through dependency interfaces. CLI derives generate
  generic struct decoding, unit-enum conversion, definition-site runtime helper
  calls and type predicates. Nested Args and typed subcommand enums compose
  across dependencies, retaining generic payload predicates and positional/global
  option semantics. Import aliases and caller-name collisions are tested.
  Composition errors remain explicit schema diagnostics, including multiple
  subcommand selectors introduced through nested flattening.
- Generic sequence algorithms produce minimal edits over exhaustive short inputs;
  the text layer interoperates with GNU diff and GNU patch in both directions.
- Generic graphs support downstream label types, captured weight callbacks,
  derived handle hashing and iterative algorithms. Shortest paths and spanning
  forests agree with independent matrix and exhaustive-subset oracles.
- A generic LSP router deserializes downstream parameter types and serializes
  downstream result types inside stored closures. Its separate consumer also
  validates public stateful APIs, standard I/O and UTF-16 document coordinates
  against an independent Python client. Negotiated UTF-8/UTF-16/UTF-32 positions
  share rope-based conversion and transactional edit semantics; 4,768 additional
  boundary checks and 48 edit sequences agree with Python. Incoming and outgoing
  deadlines contribute to one explicit event-loop wakeup calculation.
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
  aggregates. All 29 library tests also pass under Go's race detector, including
  concurrent calls, close wakeups and cancellation during partial replies.
  Bounded pools reserve slots before dialing, serialize shared lease aliases,
  retire unhealthy connections and preserve ambiguous-write errors without
  replaying commands. Real Redis checks cover concurrent INCR and replacement
  after CLIENT KILL. DNS/TLS connectors use standard network APIs. Fifteen Python
  TLS-server cases, repeated under the race detector, verify trust, mTLS,
  setup/operation deadlines and both context and legacy cancellation.

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

- Bitflags derives validate integer-backed newtypes and emit trait methods plus
  definition metadata through the ordinary public derive interface. Default
  trait methods, typed iterators, and explicit text/numeric Serde wrappers work
  across a versioned dependency. Exhaustive byte-set algebra and 4,601 comparisons
  with Rust bitflags cover aliases, overlapping flags, unknown bits and parsing.
  The optional `FlagValues` derive now emits public inherent constructors such
  as `Access::flag_read()`, including across package interfaces without trait
  imports. Existing trait APIs and module constants remain compatible;
  associated constants remain unsupported.

- Logos implements a recursive regex AST, bounded Thompson NFA construction,
  generic callbacks with extras/error types and cross-package iterator methods
  without a native regex adapter. Reusable grammars work across independent
  concurrent lexers; Python exhaustive-prefix matching checks 3,155 cases.

- Tempfile combines the public random, Linux descriptor and I/O APIs into
  exclusive creation, descriptor-relative directory cleanup, atomic persistence,
  shared resource lifecycles and in-memory spooling. Real filesystem checks and
  the race detector cover ownership transfer, symlinks, concurrent creation and
  explicit cleanup without assuming destructors or GC finalizers.

- Reqwest uses an ordinary Go FFI transport with GoML request/response types,
  redirect policy and scoped cancellation. Native HTTP/HTTPS servers and Python
  interoperability exercise certificate validation, HTTP/2, bounded bodies,
  multipart, sensitive-header isolation and connection reuse. Private body
  storage now uses `FrozenBytes`, with shared immutable accessors and explicit
  mutable copies at the public boundary. `Bytes::to_vec()` is not an isolation
  guarantee for later mutations of the returned `Vec`.

- LLVM exposes distinct GoML handle types over an opaque Go FFI interface and
  a cgo binding to LLVM 18. Context locking, module generations and explicit
  closure preserve native ownership without language lifetimes or destructors.
  Black-box tests cover IR/bitcode round trips, invalid operand combinations,
  cross-context references and stale handles after optimization. Native and
  GoML race checks cover shared lifecycles; four independently linked code
  variants agree across 9,624 function results. This design requires ordinary
  native dependencies but no compiler or builtin changes.

The rope library uses private GC-managed references to implement an immutable
AVL tree with shared subtrees. UTF-8, UTF-16 and CRLF metadata survive edits,
splits and concatenation across separately compiled consumers. LSP documents
now use this storage without changing their clamping or transactional semantics.
Logical sizes can grow through sharing without large allocations, so public
operations check integer overflow before updating cached counts. Leaf copies
prevent small slices from retaining unrelated large source strings.

The incremental library composes heterogeneous input and query types through
typed handles and closures that erase dependency validation only. Revision
tracking, equal-result cutoff and atomic input batches work across versioned
module boundaries. Channel gates serialize root operations and join in-flight
callback reads before publishing memoized values. Thirteen tests pass under
the race detector; 15,847 queries agree with independent from-scratch evaluation.
Mutable values require an explicit copy policy, and callbacks use their scoped
`Evaluation` rather than reentering a blocking database operation.

The web framework passes lifted GoML callbacks through ordinary Go FFI into
concurrent HTTP handlers and streaming producers. Typed form/query decoding
implements the Serde deserializer protocol in GoML; body and response streams
implement public I/O traits. Opaque native interfaces keep FFI metadata bounded.
Cancellation callbacks must finish before connection deadlines are reset, so a
completed request cannot poison a later keep-alive request. Native tests and a
separate race-built consumer exercise network backpressure and SSE disconnects.

Bigint implements signed and unsigned arbitrary-precision arithmetic using
immutable `FrozenVec[u32]` limbs and `u64` intermediates, including normalized
multi-limb division and negative two's-complement bitwise semantics. Public
numeric and Serde traits specialize in independent consumers. Exact methods
provide the arithmetic API without operator overloading or wide integer
primitives; the decimal module can reuse this through an ordinary dependency.

Tracing combines immutable structured fields, explicit task contexts and a
bounded channel worker without a native adapter. Root sampling has a separate
ordinal from span identity allocation, so nesting cannot bias the sample.
Trace sampling remains separate from record filtering: a hidden parent can
carry a visible error event. All 26 tests pass under the race detector, including
queue saturation, concurrent closure and first-error propagation. Sink callbacks
must cooperate with shutdown and cannot synchronously reenter their own tracer.

Datetime implements checked calendar arithmetic, nanosecond normalization,
bounded TZif decoding and POSIX timezone rules entirely in GoML. Shared immutable
zone snapshots support concurrent conversions and explicit gap/fold resolution.
Python and Go reference implementations disagree on some synthetic POSIX edge
cases; the README records the specification, minimal cases and oracle selection.
All 18 library tests pass under the race detector and 8,140 cases agree with
their documented calendar, timezone or specification reference.

Decimal consumes bigint through a versioned dependency and implements exact
coefficient/scale arithmetic, numeric hashing and format-sensitive Serde without
floating-point intermediates. Independent Python comparisons cover both values
and `Rounded`/`Inexact` status: discarding zero positions can count as rounding
without changing the value. Private bounded representations keep intermediate
allocation predictable, while explicit context and quantum policies avoid global
rounding state.

Cache uses generic hash keys and private linked LRU entries behind a channel
gate. Singleflight results are published through completion channels; generations
prevent invalidated loads from restoring stale values, while active loader counts
remain bounded until callbacks return. Clock, copy, loader and removal callbacks
run outside the gate. A final context check at publication handles cancellation
during callbacks. All 22 tests pass under the race detector; 42,240 operations
agree with an independent weighted LRU/TTL/TTI model.

Ignore implements byte-oriented glob dynamic programming and immutable
hierarchical rule sets without a native matcher. Bounded descriptor iteration,
canonical ancestor checks and scoped callback workers handle real filesystem
traversal. Git reference queries distinguish an explicit trailing slash from a
directory entry visited during traversal. Worktree Git metadata can live outside
the scan root and requires `gitdir`/`commondir` resolution that preserves path
spaces. All 21 tests pass under the race detector; 9,400 queries agree with Git.

Syntax combines immutable green trees with parent-aware red views and
consumer-defined typed AST traits. Frozen children, checked cached byte/element
counts and iterative traversal support deep trees and persistent path-copy edits
without ownership syntax. Bounded shared interning works across tasks; all 18
tests pass under the race detector. Independent tree models check 3,840 edits,
and a versioned configuration-language consumer preserves trivia through 240
text rewrites.

The statistics tool consumes ignore through a normal versioned dependency.
Its public scan API composes a custom exclusion closure with hierarchical Git
rules and bounded filesystem traversal, preserving explicit-root behavior and
returning errors instead of partial totals when traversal fails.

Notify and walkdir now compile as independent ecosystem modules using only
public standard filesystem, byte, syscall, task and time APIs. Moving their
implementations out of the compiler's standard-package catalog preserves their
public API and Linux descriptor lifecycles. Their separate consumers retain
the original compiler-module scenarios for recursive notifications, cancellation,
subscriptions, symlink traversal and directory syscall behavior.

## Fixed compiler regressions

Ignore-file BOM handling exposed a Go output bug: a decoded `\uFEFF` string
escape was emitted as a literal BOM inside generated Go, which Go rejects.
The backend now emits `\ufeff`, preserving the UTF-8 value. Printer and real
Go compilation/execution regressions cover embedded and repeated marks.

GoML 0.1.50 passes all four retained reproducers in `repros/`:

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
`TypedCommand[T]` is retained for compatibility, while CLI now exposes
`schema::[T]() -> Command` and exercises erased generic arguments through
forwarding and function values. Ndarray adds `Array::[f64]::linspace` with inferred
owner calls and a compatible module-level forwarder. SQLite's consumer removes
the transport workaround and uses standard I/O directly in its FFI package.

## Type inference and standard-library interfaces

- Numeric operations in closures whose expected parameter is an associated-type
  projection can require an explicit parameter annotation. Consumer tests retain
  this syntax. A typed intermediate `Result[T, string]` also resolves `Self` for
  static trait calls before chaining `map_err`.
- Imported short trait bounds now retain the defining trait identity through
  monomorphization. Regression modules cover aliases, re-exports, forwarding
  helpers, and inherent impl bounds. Generic static inherent methods also retain
  owner arguments when their parameter and return types erase those arguments.
- Integer `to_string` now works in CTFE; bitflags uses it when generating masks.
  Derives can generate public inherent methods. General CTFE collections,
  associated constants, and const generics remain future work.
- Derive attribute metadata does not retain every token form: numeric and raw
  string values can be omitted by attribute lowering. CLI validates the retained
  raw attribute text, rejects invalid non-string values and decodes raw strings
  itself; escaped ordinary strings continue to use decoded compiler metadata.
- `std::resource` combines action and cleanup errors and provides concurrent,
  idempotent LIFO scopes. `Bytes::copy` and `freeze` make buffer isolation explicit.
  Serde formats can distinguish text from binary through `is_human_readable`.
  Tempfile's scope helpers now use `io::with_resource` while retaining their
  existing single-cleanup-error API and deferred idempotent close.
- Native adapter declarations now resolve Go modules from registry source paths
  using generated artifact-local module files. LLVM major and cgo prerequisites
  are checked; system packages and linker search paths still require setup.
- A separate pre-existing namespace limitation remains: an imported user trait
  named `ToString` can interfere with builtin formatting predicates. A minimal
  package-import reproduction also fails with pinned stage0. The new CTFE path
  checks builtin identity exactly and does not evaluate such user trait calls.
- Scalar conversion is available through `std::num::{ToFloat, TryToInt}`.
  Integer-to-float rounding and checked float-to-integer conversion are verified
  against native Go across all rounding modes, boundary values and random bit
  patterns. Ndarray count conversion and template numeric coercion now use
  `ToFloat`; property generators retain bit construction for reproducible IEEE
  distributions.
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

MessagePack now provides generic `StreamReader[R: Read]` and
`StreamWriter[W: Write]` adapters. Its retained frame scanner supports typed
decoding directly from buffered bytes, preserving binary and extension events.
Tests cover downstream Serde derives, bounded concatenated frames, partial
transfers, invalid stream counts, clean/truncated EOF, type-mismatch retry and
terminal partial-write failure. Values remain buffered individually; incremental
field callbacks are still outside this interface.

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

## Terminal and color libraries

The terminal batch compiles as ordinary versioned ecosystem modules without new
compiler syntax, builtin hooks or standard-library changes. Public Linux
syscalls, descriptor wrappers, contexts, scoped tasks and channels are sufficient
for raw terminal sessions, cancellable poll/read/write and concurrent progress
state. Real PTY and race checks exercise descriptor restoration and copied handle
ownership. Resize is detected by polling; a portable, runtime-coordinated signal
subscription API remains absent. No library installs raw signal handlers in the
Go runtime or claims restoration after process termination or panic.

Unicode segmentation and layout can be implemented entirely in GoML with compact,
version-pinned property tables. The Unicode 16 implementation passes all official
grapheme, word and line-break cases. Applications must distinguish UTF-8 byte
positions, grapheme boundaries and terminal columns: tabs, combining marks and
wide emoji exposed actual editor and diagnostic alignment bugs during review.
Terminal glyph width still depends on fonts and emulator policy.

Color conversion, gamut mapping and interpolation are GoML algorithms; ordinary
FFI supplies only primitive math operations absent from the public math API.
A high-precision reference exposed underflow from multiplying colors directly by
subnormal alpha. Normalizing the blend weights before mixing preserves tiny
nonzero alpha and finite channels, including unequal color endpoints.

Associated output types and generic callbacks support `prompt::Model` consumers
returning strings, integers, booleans and arbitrary selected values. Closures
that only produce an error can require an explicit expected concrete output
type; typed helper functions also avoid unstable inference around `?`. Generic
`Option` and some tuple arities lack `Debug`, so public APIs use explicit debug
implementations and tests compare such values without assuming that bound.

Snapshots need deliberate container copies because `Vec`, `Ref` and model handle
copies share storage. Cross-library tests cover detached buffers, choices,
gradient stops and progress snapshots. Mutable render/editor/prompt state has one
event-loop owner; synchronized terminal/progress handles explicitly support
concurrency. Progress snapshots let a full-screen UI own output without competing
with a second cursor-control writer.

Terminal lifecycle and output limits remain API responsibilities. Ordinary
failure must leave input intact, invalidate partially written frames, wake
blocked operations and restore modes. Cross-review added regressions for wide
cell replacement budgets, tab navigation, replacing selected text with the same
text, bounded undo/redo storage and callback invalidation during rendering.

One remaining compiler boundary is independently reproducible with the current
stage2 toolchain (version 0.1.50):
assigning a public field of a dependency struct that also contains private fields
can typecheck and then fail ANF validation at link time. For example, with a normal
`ecosystem::tui = "0.1.0"` dependency:

```gom
use ecosystem::tui;

fn assign(state: tui::TreeState) -> tui::TreeState {
    let mut value = state;
    value.selection = tui::SelectionState { selected: 2, ..value.selection };
    value
}
```

The linker reports `constructor argument count does not match ecosystem::tui::TreeState`. This is a compiler limitation, not an intended field
visibility rule. The ecosystem batch does not alter compiler lowering;
`TreeState::with_selection` performs the update inside its defining package,
and the Explorer uses that public builder. The API's independent consumer and
Explorer normal/race/PTY tests verify the working path.
