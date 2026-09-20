# incremental

A pure GoML library for typed, demand-driven incremental computation. Inputs and
query families can have different types in the same database. Query callbacks
compose through `Evaluation`, which records the inputs and queries actually read.
The design takes revision validation and unchanged-output cutoff from the
[Salsa model](https://salsa-rs.github.io/salsa/overview.html); it does not emulate
Rust macros, lifetimes, or Salsa's entire API.

## Example

```goml
use ecosystem::incremental as inc;

fn example() -> Result[string, inc::Error] {
    let db = inc::Database::new();
    let source = db.input("source", "hello");
    let length = db.query("length", 64, |read: inc::Evaluation, _: isize| {
        Result::Ok(read.input(source)?.byte_len())
    })?;
    let label = db.query("label", 64, |read: inc::Evaluation, key: isize| {
        Result::Ok(read.query(length, key)?.to_string() + " bytes")
    })?;
    let _ = db.fetch(label, 0)?;
    let _ = db.set(source, "world")?;
    db.fetch(label, 0)
}
```

The second fetch validates `length`, discovers that its output is still five,
and reuses the cached label without invoking its callback again.

## API and behavior

- `Database::new()` creates a database at revision 1 with a 256-query depth
  limit. `with_depth_limit` validates a custom positive limit.
- `input(name, value)` creates an `Input[T]`; `set` returns whether its value
  changed. Equal assignments preserve the revision. Input reads happen through
  `Evaluation::input`, so they become tracked dependencies.
- `Input::change` prepares a type-erased update while preserving the input's
  typed setter. `Database::apply(Vec[Change])` changes heterogeneous inputs
  atomically under one revision. Foreign handles are checked before any update.
  All equal assignments form a no-op batch. Repeated writes to the same input
  execute in order; a batch that changes a value and restores it still advances
  the revision. Prepared changes can be reused.
- `query(name, capacity, callback)` creates a typed `Query[K, V]`. Keys implement
  `Hash + Eq + ToString`; values implement `PartialEq`. `query_with` accepts a
  custom key label and `ValuePolicy[V]`, requiring only `Hash + Eq` for keys.
- `fetch` evaluates a root query. Callbacks use `read.query(other, key)` to
  compose different key and result types. A query can recursively invoke its own
  family through an installed `Ref[Option[Query[K, V]]]`.
- Cached results are validated lazily after input changes. Dependencies retain
  the revision at which their observable result last changed. Equal recomputed
  values retain that earlier revision, stopping propagation to parent queries.
  Every successful recomputation replaces its dependency list, including when
  the new result compares equal. Unread branches do not invalidate a query.
- Failed evaluations are retried on the next request and never publish partial
  results. A parent may catch a child error and return a fallback; this records
  a conservative revision dependency so later input changes retry that branch.
  Validation errors cause the parent callback to run again, allowing its normal
  error-handling logic to decide the result.
- Cycles return `ErrorKind::Cycle` with a path of labeled query invocations.
  Depth exhaustion, cross-database handles, cancellation, expired evaluations,
  concurrent capability reuse, and user errors also return structured errors.
- Each query family bounds its memo entries with LRU eviction. Active recursive
  frames are pinned; the capacity can temporarily exceed its target by active
  depth and returns to the target as frames finish. Evicted dependencies are
  safely recomputed, potentially sacrificing early cutoff. The bound counts
  entries, not retained bytes or dependencies within an entry.
- `clear(query)` releases its cached entries and returns their count. Handles
  remain usable. `inspect(query)` returns copied labels, change/validation
  revisions and dependency names. `stats(query)` reports executions, direct
  cache hits, validations, unchanged results, evictions and failures. These are
  cumulative counters; clearing entries does not reset them.

## Concurrency and cancellation

Database handles, input handles, and query handles can be shared across tasks.
Root queries, input changes, batches, inspection, and cache clearing serialize
through the database gate. Concurrent requests for the same key therefore share
one successful execution. Distinct root queries also serialize; this release
does not evaluate independent graph branches in parallel or offer snapshots.

`fetch_with`, `set_with`, and `apply_with` accept `std::context::Context`. Waiting
for the gate is cancellable, and evaluation checks cancellation between
dependency operations and before publishing results. Long-running callbacks
should call `read.check()` or use `read.context()` for cancellable I/O. Input
batches complete atomically once admitted. Cancellation does not interrupt an
arbitrary computation or blocking operation inside a user callback.

An `Evaluation` belongs to one callback invocation. Its read methods reject
concurrent or reentrant access with `ConcurrentEvaluation`. Nested query
callbacks receive their own capabilities. When a callback returns, an already
started read is joined before its dependency list or result is published; the
capability then expires. Subsequent reads report `ExpiredEvaluation`.

Callbacks must use their `Evaluation` for database access. Calling blocking
database methods such as `fetch`, `set`, `revision`, `inspect`, or `clear` from
inside a callback on the same database would wait for the gate already held by
that root evaluation. `try_fetch` and `try_set` provide nonblocking alternatives
that return `Busy`; they do not enable nested mutation. The library does not
inspect goroutine identities. Panics are outside the fallible API contract;
return `Error::user` for application failures.

## Value and key contracts

Callbacks must be deterministic functions of the key and tracked inputs. Do not
read mutable external state, the clock, or randomness without representing the
relevant value as an input. Captured query handles and immutable configuration
are fine. Hash/equality behavior of keys must remain stable after insertion.

The default `ValuePolicy::immutable()` compares with `PartialEq` and returns
values directly. This is suitable for scalars, strings and genuinely immutable
application values. It does **not** make a `Vec`, `Ref`, or mutable object graph
immutable. Use `input_with` and `query_with` with
`ValuePolicy::new(equal, copy)` when values contain mutable storage. For example,
`|value: Vec[isize]| value.copy()` isolates a vector of scalars; nested mutable
values need a deep copy. The library copies on input assignment, query result
storage and public reads. Copy/equality callbacks must themselves be pure and
respect the same database-access rule. Preparing changes concurrently also
requires the supplied copy callback to be safe for that use.

There are no tracked struct macros, durability classes, interned values, cycle
fixed-point recovery, persistence, automatic filesystem reads, or incremental
compiler migration in this release. Errors are not memoized as query values;
applications may explicitly use a result-like value type when they want that
behavior.

## Verification

From the repository root:

```sh
python3 ecosystem/verify.py incremental
```

This formats/checks the independent modules, runs library and consumer tests,
builds the versioned consumer twice to verify cache stability, executes its
smoke case, and runs both interoperability and Go race-detector checks.

The Python oracle evaluates every requested result from scratch, without sharing
the implementation's cache or revision algorithm. Histories cover random DAGs,
dynamic branches, equality cutoff, failing child queries, caught errors, cache
eviction, clearing, and cycles that appear and disappear after edits. Focused
GoML tests cover heterogeneous composition, diamond reuse, dependency replacement,
copy isolation, atomic batches, foreign handles, limits, cancellation, concurrent
root deduplication and callback lifetime races.
