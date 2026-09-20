# bench

A native GoML benchmark runner with adaptive sampling, reproducible statistical
analysis, baseline comparison and standalone reports. GoML implements scheduling,
statistics, validation and rendering. The small Go adapter only provides atomic
scalar input/output barriers and `runtime.KeepAlive`.

## Running benchmarks

```goml
use ecosystem::bench;

let case = bench::Case::new("sum", || {
    let length = bench::black_box(1000);
    let mut sum: u64 = 0;
    let mut index: u64 = 0;
    while index < length {
        sum += index;
        index += 1;
    }
    bench::black_box(sum)
}).with_group("arithmetic").with_throughput(bench::Throughput::Elements(1000));
let report = bench::run(case, bench::Options::standard())?;
```

`Case::new` benchmarks a closure returning `u64`; use a checksum or digest of
richer output. `Case::fallible` propagates workload failures. `with_group`,
`with_parameter`, `with_throughput` and `parameterized` build named benchmark
matrices. `run_suite` returns all raw samples and configuration in a `SuiteReport`.
Duplicate group/name/parameter identities are rejected before executing workloads.
A failing case stops the suite; callers needing partial progress can call `run`
individually and retain completed reports.

`Case::with_setup[T]` and `try_with_setup[T]` create fresh input outside the timed
section on every iteration. The work closure alone is measured, followed by
output consumption outside that section. This avoids accumulating a batch of
large setup objects. Normal `Case::new` measures a complete batch, including loop,
fallible-call dispatch, periodic cancellation checks and output consumption.
Reports record `Timing::Batch` versus `Timing::PerIteration`; baseline comparison
rejects different timing models. Both modes include any garbage collection or
runtime scheduling that occurs inside their timed regions. The library does not
force garbage collection or subtract timer/loop overhead.

## Measurement and work bounds

Every run performs calibration even if warmup is zero. Warmup doubles iterations
until the requested accumulated measured duration is reached, subject to 64
rounds and configured work limits. Sampling adapts each subsequent batch toward
`measurement_ns / samples`, bounded by minimum/maximum iterations. Measurement
duration is a target, not an exact sleep or a performance assertion.

`Options` controls warmup, measurement target, sample count, per-batch and total
iterations, wall-clock budget, bootstrap resamples, confidence, noise threshold
and seed. Defaults use 100 ms warmup, 500 ms measurement, 30 samples and 1,000
bootstrap resamples. Each run has a default 30-second wall budget. Samples are
limited to 10,000, bootstrap work to 20 million sample draws, and a suite to 1,000
cases/100,000 samples. The suite uses the options' work/wall budgets per case.

`run_with` and `run_suite_with` accept a `Clock` and `std::context::Context`.
`Clock::monotonic` uses `std::time::Instant`; `Clock::new` accepts an injected
nanosecond clock for deterministic tests. Negative/backward clocks, unresolved
zero durations, invalid configurations and checked accumulation failures return
typed errors. Clocks are stateful and should be used sequentially; use separate
clocks for independent concurrent benchmark runs.

Cancellation/deadlines are cooperative: normal batches check at least every 256
iterations and at batch boundaries; setup mode checks each iteration. A workload
closure cannot be preempted while it is executing. Wall budgets are checked at
measurement boundaries and after statistical analysis; a batch may overshoot the
budget by its own duration. `bootstrap_mean_with`, `analyze_with` and `compare_with`
check the context every 32 resamples; their non-`with` versions use a background
context. Statistics and comparisons have independent explicit work limits.

## Statistics and comparisons

`analyze` retains every raw sample, including outliers, and computes:

- Arithmetic mean of per-iteration sample times and a seeded percentile-bootstrap
  confidence interval.
- Median, minimum/maximum, sample standard deviation and median absolute deviation.
- Mild/severe Tukey outlier counts, using 1.5/3 IQR fences.
- Least-squares timing slope through the origin, exposed as a separate estimate.

`bootstrap_mean` is also available for independent positive finite observations.
The same data, seed and configuration produce identical statistical results.
`throughput_per_second` converts mean iteration time into bytes/second or
elements/second and rejects missing, zero or overflowing values.

`compare(baseline, current)` checks group, name, parameter, throughput and timing
model, then bootstraps the two raw sample distributions independently. It returns
a relative-change interval and `Improved`, `Regressed`, `WithinNoise` or
`Inconclusive`. Classification uses the configured practical noise threshold and
entire interval. It is not a p-value or Criterion-compatible hypothesis test.
The comparison recomputes from raw measurements rather than trusting serialized
summary statistics. Callers remain responsible for matching compiler options,
hardware, operating-system load and dataset semantics across runs.

## Reports and files

`to_json`/`from_json` preserve schema version 1, identity, timing model, throughput,
raw samples, summary statistics, options, iteration totals, wall time and result
checksum. Readers validate schema, shapes, finite statistics and totals and limit
input/output JSON to 16 MiB. Parsing also limits nesting to 64 and structural tokens to one million. Unknown future schema versions are errors.

`to_html` produces a standalone UTF-8 HTML table, confidence intervals, throughput,
outlier counts and an expandable raw JSON section. User-controlled names and
parameters are HTML escaped. `write_json` and `write_html` use atomic file
replacement; `read_json` performs a bounded file read. File reading currently
uses Linux standard file descriptors. Rendering and analysis have no network or
browser dependencies.

## Result consumption and optimization

`black_box(u64)` returns an opaque local atomic Store/Load value, preventing the
Go compiler from propagating the input value through a plain identity return.
Each call owns its atomic cell, so concurrent calls cannot substitute another
caller's input. Workload results are also stored in a process-wide atomic sink;
`observed()` is diagnostic only and may reflect any concurrent runner. Report
checksums are accumulated separately for each run.

Use `black_box` on inputs as well as consuming outputs, as shown above. Consuming
a precomputed constant cannot recreate the computation that produced it. The
barrier and sink have measurable atomic/dispatch overhead, and the scalar API is
not a universal optimizer fence for arbitrary GoML values. There are no unsafe
pointers, no generic memory-reinterpretation tricks and no claims of exact
compatibility with Rust's `std::hint::black_box`.

## Validation and example

```sh
../../stage2/bin/goml fmt --check
../../stage2/bin/goml check
../../stage2/bin/goml test
GOFLAGS=-race ../../stage2/bin/goml test --target-dir _artifact/race
```

Native tests use fixed statistical data and injected clocks. They cover warmup
and batch adaptation, setup exclusion, output retention, cancellation/deadlines,
work limits, backward/zero clocks, bootstrap replay, outlier retention, comparison
categories, identity mismatches, JSON validation, HTML escaping and concurrent
black-box correctness. Tests do not assert absolute machine performance.

The independent `../consumers/bench` module imports version `0.1.0`, measures
insertion sort and standard sorting for 32/128-element inputs, and writes
`_artifact/bench-report.json` and `.html`. Run `just ecosystem-test bench` from the
repository root for the isolated registry consumer workflow.

Design references: [Criterion analysis](https://bheisler.github.io/criterion.rs/book/analysis.html),
[Criterion timing loops](https://docs.rs/criterion/latest/criterion/struct.Bencher.html),
and [Go KeepAlive](https://pkg.go.dev/runtime#KeepAlive).
