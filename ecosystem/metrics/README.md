# metrics

A GoML-native metrics registry with typed handles, bounded label cardinality,
consistent snapshots and Prometheus text exposition. There are no Go adapters,
Python helpers, global singleton registries or background workers.

```toml
[dependencies]
"ecosystem::metrics" = "0.1.0"
```

The handle design is inspired by [Rust metrics](https://docs.rs/metrics/latest/metrics/).
The exporter implements the classic
[Prometheus text format 0.0.4](https://prometheus.io/docs/instrumenting/exposition_formats/)
and uses the conservative ASCII metric/label identifier syntax described by the
[Prometheus data model](https://prometheus.io/docs/concepts/data_model/).

## Recording measurements

```gom
use ecosystem::metrics as metrics;

fn requests(registry: metrics::Registry) -> Result[metrics::Counter, metrics::Error] {
    let descriptor = metrics::Descriptor::counter("http_requests_total", "Completed requests")
        .with_labels(Vec::from_array(["method", "route"]));
    let labels = metrics::Labels::new(Vec::from_array([("method", "GET"), ("route", "/")]))?;
    let counter = registry.counter(descriptor, labels)?;
    counter.increment()?;
    Result::Ok(counter)
}
```

Create registries with `Registry::new(Options::standard())`. Repeated registration
of the same descriptor and label values returns handles for the same series.
Handles can be copied and used concurrently across tasks.

- `Counter` stores an exact `u64`. `increment`, `add` and `absolute` are monotonic;
  `absolute` keeps the larger of the current and supplied value. Overflow returns
  an error without changing the value.
- `Gauge` stores finite `f64`. `set`, signed `add`, nonnegative `increment` and
  nonnegative `decrement` reject nonfinite inputs and arithmetic overflow.
- `Histogram` stores finite, strictly increasing explicit bucket boundaries,
  cumulative inclusive bucket counts, a checked `u64` count and a finite `f64`
  sum. `observe_many(value, count)` records weighted repeated observations
  atomically. An implicit `+Inf` bucket is always exported with the total count.
  Empty explicit boundaries are allowed. Negative observations are allowed;
  therefore the sum is not necessarily monotonic.

Use `linear_buckets` or `exponential_buckets` to construct checked boundaries.
They reject overflow and boundaries that become equal through floating-point
rounding. Histogram count and sum errors leave all buckets unchanged.

`Histogram::observe_duration` records seconds. `start_timer` creates a monotonic
timer; `stop` records once and returns elapsed time, while `cancel` discards it.
Copies share the stop/cancel state. A second stop returns `Finished`. A failed
recording also consumes the timer. `Histogram::time` measures a synchronous
callback and returns its result after recording. Timers have no implicit drop
behavior; use explicit `stop` or `defer` when appropriate. Use a seconds unit and
seconds bucket boundaries for duration histograms.

## Registration, labels and lifecycle

A descriptor contains name, help, unit, kind, label names and histogram bounds.
`with_labels` and `with_unit` construct variants. `register` reserves metadata
without creating a sample. Help text, unit, kind, normalized label schema and
bucket boundaries must agree on every registration of a name. Units are retained
in snapshots; text 0.0.4 has no `UNIT` directive, and units do not rename metrics.
Counter names are exported verbatim; append `_total` in application schemas.

Label names are normalized into sorted order. Each series must supply exactly
the descriptor's labels. Duplicate, malformed and reserved `__` names are
rejected. Histograms reserve the `le` label. Metric identifiers may contain a
colon; label and unit identifiers may not. Label values support Unicode and
empty strings. Source vectors are copied so later caller edits cannot change
registered schemas or series identities.

The registry reserves histogram base, `_bucket`, `_sum` and `_count` names
against every other metric family, in both registration orders. This prevents
ambiguous duplicate samples in exposition output.

`Options::standard()` bounds the registry to 1,024 families, 10,000 total series,
1,024 series per family, 16 labels per series, 1,024 UTF-8 bytes per label value,
256 bytes per identifier, 4,096 help bytes and 256 explicit histogram boundaries.
Rejected registrations do not leave a partially created family or series.
These limits can be adjusted when constructing the registry.

`remove_series` removes one normalized label set but retains its descriptor.
`unregister` removes a family and releases its reserved sample names. `clear`
removes all families and series. All three operations mark previously returned
handles as detached; future reads or updates through those handles return
`Detached`. Re-registering the same name/labels creates a fresh zero-valued
series. It never reconnects stale handles. Existing snapshots remain valid.
Callers retaining detached handles retain those handle objects, independently
of the active registry cardinality limit.

## Snapshots, collectors and concurrency

All registry mutations and observations use a shared gate. Handle updates go
directly to their registered cell; they do not scan the registry. `snapshot`
copies all descriptors and values while holding the gate, producing one
consistent view across the registry. Snapshots expose isolated copies through
`families()` and `series_count()`. Later mutations or deletion do not change a
snapshot. Exposition orders families by name and series by label values for
reproducible output.

`collect_gauges(context, callback)` adapts external state into existing gauges.
The callback receives the context and returns `Vec[GaugeReading]`, each with
name, labels and value. It runs outside the registry gate, so it may perform I/O
or call registry APIs. The complete result is validated and applied under the
gate as one atomic batch. Unknown targets, duplicate readings, invalid values,
wrong kinds or cancellation reject the entire batch. Register target series
before collecting. Duplicate detection uses normalized, length-prefixed keys
in a hash map. The validation loop checks cancellation before each reading;
target lookup scans the bounded family and series lists. Large batches therefore
cost up to readings times registered families/series while holding the gate.
The callback itself is responsible for checking its context
during long work; arbitrary user code cannot be forcibly interrupted.

`snapshot_with`, `render_with`, counter `add_with`, gauge `set_with` and histogram
`observe_with` honor context cancellation and deadlines while waiting for the
gate, and check before committing changes. Rendering checks again before
returning output. Plain methods wait until the gate is available. The registry
uses no callbacks while holding its gate.

## Prometheus export and HTTP

`Registry::render()` and `Snapshot::prometheus()` produce UTF-8 text with HELP,
TYPE and correctly escaped labels/help. Histogram buckets are ascending and
cumulative; the `+Inf` bucket equals `_count`. Nonempty output ends in a newline.
`content_type()` returns `text/plain; version=0.0.4; charset=utf-8` for an HTTP
response. An empty registry produces an empty body. Metadata-only families emit
HELP/TYPE without samples. Negative zero exports as `0`.

The independent versioned consumer in `../consumers/metrics` records request
counts, active requests, pool state and elapsed time, serves `/metrics` over a
real local TCP HTTP exchange, and checks the scraped output. Applications using
`ecosystem::web` can place `render_with` in their route handler and set the
returned content type; the metrics library itself has no HTTP dependency.

```sh
cd ecosystem/metrics
../../stage2/bin/goml fmt --check
../../stage2/bin/goml test
GOFLAGS=-race ../../stage2/bin/goml test --target-dir _artifact/race --timeout 300s
```

Native tests cover exact exposition vectors, Unicode and escaping, normalized
schemas, suffix collisions, finite values, integer/count/sum overflow,
cardinality and deletion, snapshot isolation, atomic collector failures,
cancellation, timers, concurrent registration, observations and scrapes. The
race pass exercises the same GoML tests.

This initial release provides classic histograms, not native histograms,
summary quantiles, exemplars, OpenMetrics, remote write, OTLP or automatic
process collectors. It favors consistent snapshots over sharded update
throughput; a large scrape temporarily serializes updates. Floating sums use
ordinary IEEE-754 arithmetic and can accumulate rounding error. The registry
retains exact integer counts, but Prometheus ingestion represents sample values
as floating point, so very large counts may lose precision downstream.
