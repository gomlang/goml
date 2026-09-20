# Pipeline

`ecosystem::pipeline` provides lazy, typed streaming pipelines implemented in
GoML with structured tasks and bounded channels. Streams compose sequential and
parallel transformations, filtering, flattening, batching, windows, fan-in and
zipping. Terminal operations propagate typed failures, cancellation, deadlines
and backpressure and wait for all library-owned tasks to exit.

## Example

```gom
use ecosystem::pipeline;
use ecosystem::pipeline::{Stream, ParallelOptions, Order, RunOptions};

fn example() -> Result[Vec[isize], pipeline::Error[string]] {
    let source: Stream[isize, string] = pipeline::range(0, 100, 1)?;
    let parallel = source.parallel_map(
        ParallelOptions::new(4, 8, Order::Input),
        |_, value| Result::Ok(value * value),
    )?;
    parallel.filter(|value| value % 2 == 0).take(10)?.collect(RunOptions::standard())
}
```

The element type and application error type are independent. The consumer module
uses its own `Job`, `Output` and `Issue` records, including an error type without
`Debug` or `ToString` implementations. Captured callbacks specialize across the
versioned module boundary. Pipelines do not start tasks or consume sources until
a terminal operation runs them.

## Sources and transformations

| API | Behavior |
| --- | --- |
| `from_values(Vec[T])` | Copies the vector structure and replays its elements for each run |
| `range(start, end, step)` | Exclusive-end signed range, with checked direction/overflow boundaries; zero step is rejected |
| `repeat(value)` | Infinite source, stopped by downstream control, limits or cancellation |
| `generate(factory)` | Calls the factory per run to obtain a cancellation-aware, fallible next-item callback |
| `from_producer(producer)` | Push source receiving a cancellation token and synchronous `emit(T) -> bool` callback |
| `from_receiver(receiver)` | Consumes a caller-owned channel until closure or cancellation; never closes that channel |
| `map`, `try_map` | One-to-one transformation; the fallible form receives a cancellation token |
| `filter`, `filter_map`, `try_filter_map` | Retains, removes or transforms items; a missing result consumes its input position |
| `inspect` | Runs a fallible observation callback and forwards the original item |
| `enumerate` | Attaches a zero-based stage-local input index |
| `take`, `skip`, `take_while`, `skip_while` | Selects stream portions; `take(0)` does not enter the source |
| `concat` | Runs two streams sequentially; a downstream stop prevents entering the second |
| `flat_map` | Runs each returned inner stream sequentially, preserving early-stop and failure behavior |
| `scan(seed, update)` | Creates fresh state per run; the update returns new state and an optional output |
| `chunks(size)` | Emits independent vectors, including the final incomplete chunk |
| `windows(size, step)` | Emits independent full windows, supporting overlap and gaps; drops an incomplete final window |
| `named(label)` | Adds context to errors originating within that stream |

The `generate` factory has type
`() -> (CancelToken) -> Result[Option[T], E]`. It should construct fresh mutable
state for a repeatable run. The `from_producer` callback has type
`(CancelToken, (T) -> bool) -> Result[(), E]`. It may use `defer` for resources it
owns. It must call `emit` serially during the callback, stop when it returns false,
and never retain it or invoke it after returning. Use `merge` or parallel stages
for concurrent production. Sources and callbacks that wait must observe their
cancellation token. A receiver source consumes a shared queue and is not replayable.

Sequential operations execute inline and add no queues or worker tasks. Mutable
state owned by operators is recreated per run. Copies of elements are shallow;
callers must serialize changes to shared references/buffers they place inside
items or capture in callbacks. A stream with immutable sources and safe callbacks
can be run repeatedly; application-owned callback state controls whether results
are repeatable.

## Bounded parallelism

`parallel_map`, `parallel_filter` and `parallel_filter_map` accept
`ParallelOptions` and a cancellation-aware callback returning `Result[..., E]`.
Each stage owns a producer, a fixed worker pool and a coordinator. No complete
input collection is materialized. `buffer(capacity)` inserts a bounded one-worker
identity stage to decouple adjacent stages.

- `workers` limits concurrently executing callbacks (1–1,024).
- `capacity` bounds each job/result channel (0–1,048,576). Zero is a rendezvous
  channel and is supported throughout the library.
- `max_in_flight` bounds admitted inputs across queues, active callbacks, completed
  results and ordered reassembly (1–2,097,152). The constructor defaults it to
  `workers + capacity`. A smaller value than the worker count is supported.
- `Order::Input` emits successful results in input order, advancing past filtered
  items. A slow first item cannot cause an unbounded reorder buffer.
- `Order::Completion` emits results as they arrive at the coordinator, allowing
  later inputs to pass a blocked earlier input. Ordering among simultaneous
  completions depends on scheduling.

An admission credit is returned only after the corresponding result has been
consumed downstream or filtered out. Therefore a slow sink eventually stops
workers and upstream production. The producer can hold **one additional input**
while waiting for a credit. This is separate from the in-flight bound. Chaining
parallel stages adds their individual bounded queues; it does not establish one
global item/byte budget. Limits count elements, not the heap occupied by values
or arbitrary allocations performed by user callbacks.

`merge(streams, capacity)` runs at most 1,024 source streams concurrently and
interleaves them through one bounded channel. Each producer can retain one value
while blocked on a send. Per-source ordering is preserved; cross-source order is
unspecified. Empty input completes immediately.

`left.zip(right, capacity)` runs both inputs through independent bounded channels
and emits pairs until either input ends. It then cancels and joins the longer
input. An empty side can finish while the other is blocked before its first
value. Either side can read ahead by its channel capacity, a pending producer
value and the coordinator's pending element. Zipping does not guarantee a shared
channel or other external source remains unconsumed beyond the last pair.

## Terminals, errors and cancellation

`collect(options)` returns a vector. `for_each(options, callback)` receives a
cancellation token and item; it returns `Report { emitted, stopped }`. The
callback returns `Result[Control, E]`; `Control::Stop` includes the current item
in the emitted count and requests a graceful early finish. `fold` applies a
fallible accumulator. `count`, `first`, `find`, `any` and `all` provide common
terminal operations; searches stop when their result is known.

`RunOptions::standard()` sets a terminal limit of 1,000,000 items, with no deadline
or external cancellation token. `without_limit`, `with_timeout` and `with_cancel`
configure a run. The public `max_items` field can set an explicit nonnegative
limit. Exactly that many items may be emitted; encountering an additional item
returns `ErrorKind::Limit`. A filtered item does not count toward the terminal
limit. This limit does not bound source work: an infinite source filtered to no
outputs needs a timeout, cancellation or source-specific bound.

The timeout covers the complete run, including source production, waiting,
callbacks and cleanup. Progress does not reset it. Zero timeout and an already
cancelled token prevent entry into the source. External cancellation is inherited
through a child scope and never cancels the caller's scope.

Each asynchronous failure records an error and cancels the whole run. This wakes
cooperating downstream callbacks as well as sibling workers and sources. Early
termination cancels only scopes that must stop producing, allowing compositions
such as `take(...).concat(...)` to continue normally. Every accepted worker is
joined before returning. Errors from cleanup after an early stop are still
reported. Tasks refused during cancellation are not represented by handles that
would be joined indefinitely.

`Error[E]` exposes `kind`, `message`, `stage`, optional input `index` and optional
application `cause: E`. `Stage` preserves user callback failures; other kinds are
`InvalidConfig`, `Cancelled`, `Timeout` and `Limit`. Indices are relative to the
input of the failing stage, after earlier filters/transforms. With concurrent
failures, the first recorded failure wins; it need not have the lowest input
index. Named stages compose paths such as `parse/parallel map`. Application
errors observed during cleanup take precedence over a bare cancellation/deadline.
The library's error formatting does not require formatting traits on `E`; inspect
`cause` for application-specific details.

Failure is fail-fast: previously emitted values and application side effects are
not rolled back, and buffered successful results may be discarded. `collect`
returns an error instead of a partial vector; use `for_each` to retain explicitly
processed outputs. A callback returning an application error in response to its
token remains an application error; the library cannot infer the meaning of `E`.

Cancellation and deadlines are cooperative. They stop channel waits and signal
callbacks, but cannot preempt arbitrary GoML code or non-cancellable external
operations. A callback that ignores cancellation can delay return because the
library still waits for cleanup. Panics are not converted into typed pipeline
errors. Use the error return for expected failures.

## Validation

From the repository root:

```sh
just ecosystem-test pipeline
```

The verifier checks formatting, the library test suite, a separately resolved
versioned consumer, fresh/cached builds, executable behavior, native sequence reference tests and Go's race detector.
No Python interpreter, external service or package is required.

Library tests cover lazy/repeated execution, generator factories, shallow input
snapshots, transforms, flattening, chunk/window boundaries, typed errors,
zero-buffer channels, worker limits, bounded reassembly, sink backpressure,
ordered versus completion output, nested pools, cancellation, absolute deadlines,
cleanup errors, empty/unequal zip inputs and merge failures. Rendezvous channels
hold specific tasks at known points instead of relying on sleeps to infer
concurrency. Scope counts verify cleanup after success, stop, failure and cancel.
A dedicated test verifies upstream failure wakes blocked downstream callbacks.

Native consumer tests compare all 1,253 retained cases against independent list,
slice, accumulation, zip and multiset reference calculations, including malformed
configurations and typed errors. [Fixture provenance](../consumers/pipeline/tests/data/README.md) records the model and seed. The native verifier rebuilds and runs
both library and consumer tests under Go's race detector.
