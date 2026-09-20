# tracing

A structured tracing library implemented in GoML, with no Go adapter. It provides typed events, nested span lifecycles, explicit context propagation between tasks, subscriber composition, filtering, deterministic root sampling, and a bounded asynchronous writer. The design borrows the event/span/subscriber model from [Rust tracing](https://docs.rs/tracing/latest/tracing/) and the bounded writer distinction from [tracing-appender](https://docs.rs/tracing-appender/latest/tracing_appender/non_blocking/), adapted to GoML's explicit task scopes and garbage collection.

## Features

- `Level::{Error, Warn, Info, Debug, Trace}` and explicit target/name metadata.
- Typed fields: strings, booleans, signed and unsigned 64-bit integers, floats, and null. `Fields` is an immutable snapshot; duplicate names replace the previous value without changing its position. `from_vec` and `to_vec` copy their mutable outer storage. Field values contain no mutable containers.
- Stable tracer-local trace/span identifiers, strictly ordered record sequence numbers, wall-clock nanosecond timestamps, monotonic span elapsed time, span field updates, and idempotent span closure.
- Explicit `TraceContext` passed to tasks, with foreign-tracer and closed-parent validation. Closing a parent leaves existing children usable; it prevents new children or events through that parent's own context.
- Level and namespace target filters, longest matching namespace precedence, last directive wins on equal targets, and atomic filter reload.
- Deterministic all/none/every-N root sampling. Root spans and standalone events each consume one sampling ordinal, including filtered roots. Child spans and their events inherit the root decision and never consume an ordinal.
- Callback subscribers, fanout, subscriber-local filters, synchronized in-memory collection, and generic `std::io::Write` text/JSON sinks.
- Bounded queue, blocking or drop-newest event publication, dropped/accepted/filtered/live-span metrics, cancellation and deadlines while acquiring admission, queueing, flushing, and waiting for close.
- Ordered flush barriers, coordinated concurrent close, first-error latching, and exactly one sink finish callback. `Span.in_scope` and `Tracer.in_scope` preserve both an action error and cleanup errors through `std::resource::ScopeError`.

## Example

```gom
use ecosystem::tracing;
use std::context;
use std::io;
use std::task;

fn main() -> () {
    task::scope(|scope| {
        let ctx = context::Context::background();
        let sink = tracing::Sink::writer(io::stdout(), tracing::Format::Json);
        let tracer = tracing::Tracer::new(scope, sink, tracing::Options::standard()).unwrap();
        let span = tracer.span(
            ctx,
            Option::None,
            tracing::Metadata::new(tracing::Level::Info, "service::http", "request"),
            tracing::Fields::new().with("path", tracing::Value::Text("/users")),
        ).unwrap();
        let parent = span.context();
        let job = scope.spawn(|_| {
            tracer.event(
                ctx,
                Option::Some(parent),
                tracing::Metadata::new(tracing::Level::Info, "service::db", "query"),
                tracing::Fields::new().with("rows", tracing::Value::Uint(3)),
            ).unwrap()
        });
        let _ = job.join();
        span.close(ctx).unwrap();
        tracer.close(ctx).unwrap();
    });
}
```

The independent `ecosystem/consumers/tracing` module is a versioned consumer that simulates request handlers with concurrently traced lookup tasks. It exercises exported contexts, callbacks, immutable fields, generic scoped results, hidden spans, and parent closure before child completion. `--json` reads an array of scenarios from stdin and emits records plus metrics for each scenario.

## Filtering and ancestry

Sampling and filtering are separate decisions. `TraceContext.sampled()` reports only root sampling. An `Info` threshold can hide a `Debug` span while allowing an `Error` event emitted through it. A hidden span forwards the nearest visible ancestor's ID into child `parent_id` and event `span_id`; zero means no visible ancestor. Its logical ID and lifecycle still exist internally. A visible child belongs to the same trace even when the root itself is hidden.

Existing visible spans always publish their updates and end record after a filter reload, preserving their lifecycle. The current global filter controls new spans and events. Hidden spans are not retroactively published when a filter is relaxed. Events from a still-live hidden child may reference a visible ancestor that has already closed; that link expresses ancestry, not a claim that the ancestor remains open. Events have `parent_id = 0`; their enclosing visible span is `span_id`.

`Sampling::Every(N)` selects root ordinals `1, 1 + N, 1 + 2N, ...`. Root selection is independent of nesting depth. Identifiers and sampling counters are local to a tracer, start at one, and reject exhaustion rather than wrapping. They are not distributed trace IDs. `enabled` is a filter-only advisory query; it does not consume a sampling ordinal, and a subsequent concurrent reload can change admission.

## Queue, ownership, and error contract

There is one worker per tracer, owned by the supplied `std::task::Scope`. The worker serializes write, flush, and finish callbacks. Queue capacity is a record/command count, from 1 to 1,048,576; it does not bound the byte size of an individual field or allocations retained by callers. One additional record may be in the active sink callback.

`Overflow::DropNewest` applies to events and span field updates. Start/end lifecycle records and flush/close commands use reliable queue admission and may wait even in drop mode. A dropped update does not alter the span's stored field snapshot. `Delivery::Accepted` means admission to the queue, not durable output. `metrics.accepted` counts admitted records; barriers are excluded. `filtered` counts explicitly suppressed start/event/update operations. Suppressed end records and idempotent closes do not increase it.

Records and fields may be retained by a sink. A `MemorySink.records()` result can be structurally modified without affecting the collector. `Sink.fanout` freezes its sink list and invokes every sink for the current record even if one fails, returning the first error. Per-sink filters independently select records and can produce a partial lifecycle; use the tracer-level filter when coherent visible ancestry is required.

A sink write/flush/finish failure is latched. Later records already in the queue are drained without further write callbacks, subsequent publication fails with the latched error, and barriers/close report it. Flush and finish callbacks still run. Thus callers should always inspect `flush` or `close` results before treating accepted events as delivered. The first failure is retained; later sink failures are not aggregated.

`flush(ctx)` queues a barrier after previously admitted records and calls the sink's flush callback. Later publications may proceed while the caller waits. `close(ctx)` admits a single close command under the same admission lock, rejects later publications, drains earlier commands, flushes, and calls finish exactly once. Concurrent and subsequent closes share its outcome. A timeout before close admission leaves the tracer open. A timeout after admission leaves it closing; retrying close waits for the same operation. Timeout/cancellation and simultaneous readiness can race at the admission boundary; they cannot retract an already admitted record or barrier.

Span closure is similarly retryable after a queue-admission timeout. Tracer closure does not synthesize ends for live spans; close spans first when a complete lifecycle is required. `metrics.live_spans` counts handles not successfully closed, including hidden/unsampled spans. Closing the tracer with live spans leaves this diagnostic count nonzero and later attempts to close those spans fail.

Sink callbacks must return normally. They must not synchronously call admission, metrics, filter, flush, span, or close APIs on the same tracer: the single worker and admission/queue ordering can otherwise self-wait. Reading `tracer.error()` or using an independent collector is allowed. There is no thread-local current span, automatic function instrumentation, panic recovery, implicit finalizer, or forced interruption of arbitrary blocking callbacks. Context deadlines stop the caller's wait; they cannot interrupt a sink's own I/O. `Sink.writer` flushes but does not close its borrowed writer. Use a custom finish callback when the sink owns a resource that must close.

Always close the tracer before leaving its task scope. Cancellation of the owning scope aborts queue processing, latches a cancellation error, and still calls flush/finish; accepted records may then remain unwritten. A callback that blocks forever can prevent the worker/task scope from finishing. Scoped helpers guarantee cleanup only for normal returns and `Result` propagation, matching `std::resource`; panic cleanup is not guaranteed. Give cleanup an uncancelled or separate bounded context when the action's context may already have expired.

## Validation

From the repository root:

```sh
python3 ecosystem/verify.py tracing
```

This checks formatting, black-box tests, the standalone registry consumer, cached builds, a deterministic independent Python event/ancestry/field oracle, and every library test under the Go race detector. Tests coordinate blocked sinks with channels and cover saturation, cancellation, deadlines, retryable barriers/closure, filter reload, independent sampling ordinals, alias isolation, concurrent task propagation, concurrent span/tracer closure, and sink failures.

The Python oracle generates root/child visibility and sampling decisions independently, compares event multisets, and verifies sequence numbers, start/end pairing, visible ancestry, elapsed-time constraints, typed fields, and metrics. It does not claim binary or API compatibility with Rust tracing or OpenTelemetry.

Future work includes W3C trace propagation, OpenTelemetry exporters, richer sampling policies, subscriber lifecycle aggregation, byte-budget admission, and optional instrumentation syntax once supported by the language.
