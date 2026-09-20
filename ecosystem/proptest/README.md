# ecosystem::proptest

Deterministic property testing with composable generators and lazy shrink trees.
The runner returns a structured failure; the application decides how to report
it or integrate it with `std::testing`.

```goml
use ecosystem::proptest;

fn check_addition() -> Result[Option[proptest::Failure[isize]], string] {
    let values = proptest::integers(-1000, 1000)?;
    proptest::check(values, proptest::Config::new(42), |value: isize| {
        value + 0 == value
    })
}
```

## Strategies and generators

`Strategy` has an associated `Item` type and a `generate(seed, size)` method.
Downstream packages can implement it directly. `Generator[T]` implements that
trait using a function value. Its combinators are `map`, `filter`, `zip` and
`flat_map`; all retain shrinking. A `Sample[T]` holds a value and a lazy function
returning candidate samples. `Sample::leaf` makes a value with no smaller cases.

| Generator | Contract |
| --- | --- |
| `integers(min, max)` | Half-open machine-integer range, including ranges wider than signed `isize` |
| `unsigneds(min, max)` | Half-open `u64` range with shrinking toward the lower bound |
| `any_u64()`, `any_i64()` | Full-width integer generation, shrinking toward zero |
| `just(value)` | One constant value |
| `booleans()` | Boolean values; true shrinks to false |
| `select(values)` | Snapshots choices; shrinks toward the first |
| `one_of(generators)` / `weighted(entries)` | Alternative generators with positive checked weights |
| `optional(generator)` | Optional values; present values shrink toward absence |
| `vectors(item, min, max)` | Inclusive length bounds, guided by the runner's size |
| `characters()` | Unicode scalar values, excluding surrogates |
| `strings(characters, min, max)` | Character-count bounds with Unicode-preserving shrinking |
| `bytes(min, max)` | Byte vectors |
| `finite_floats(min, max)` | Finite values within inclusive bounds |
| `float64s()` | IEEE bit patterns mixed with zeros, subnormals, extrema, infinities and NaNs |
| `recursive(leaf, branch, depth)` | Recursively composed generators with a construction depth limit |
| `unique_vectors_by(item, key, min, max, attempts)` | Unique keys, bounded generation attempts and invariant-preserving shrinking |
| `sets(item, min, max, attempts)` | `std::collections::HashSet[T]` with distinct-element bounds |
| `maps(keys, values, min, max, attempts)` | `HashMap[K, V]` with distinct-key bounds |
| `stateful(initial, command, transition, min, max, attempts)` | Model-dependent command sequences whose prefixes remain valid during shrinking |

Integer shrinking moves toward zero, or toward the nearest valid bound when zero
is excluded, and probes progressively smaller steps to find failure boundaries.
Vectors try the minimum-length prefix, progressively smaller chunk deletions,
then individual element shrinks. Candidates are produced one at a time, and
large copies are charged before allocation.
Filtering retains only accepted shrink candidates and searches rejected branches
for acceptable descendants, bounded to 1,024 visits per expansion. Dependent
`flat_map` shrinking regenerates inner values after shrinking the outer value,
preserving generation dependencies. Alternative choices shrink within their
selected generator; they do not switch arbitrarily to an unrelated generator.

Invalid ranges and exhausted rejection budgets return errors. Maximum vector
length is 1,000,000 and recursive construction depth is at most 64. Floating
shrinking has a depth limit of 64. These generators are intended for testing;
their modulo-based selections do not promise unbiased statistical sampling.

Collection length bounds are inclusive. `attempts` bounds total draws while
constructing one unique collection or command sequence, including accepted draws.
An impossible uniqueness/precondition requirement returns an error instead of
looping. `stateful` takes `initial: () -> M`, `command: (M) -> Generator[C]` and
`transition: (M, C) -> Result[M, string]`. A rejected transition discards that draw.
Shrinks replay transitions from a fresh initial model and reject invalid command
prefixes. Transitions must be pure. Application execution and reset can use the
`StateMachine` helper described below or be implemented in the property. Keys
and shared values must not mutate during generation or shrinking.

`float64s` mixes explicit IEEE edge cases with random 64-bit payloads. Nonfinite
values and negative zero shrink to positive zero. Finite values retain the
bounded numeric shrinker; NaN payload bits are not normalized during generation.

## Running and replaying

`Config::new(seed)` selects 100 cases, size 64 and 10,000 shrink attempts. All
fields are public. `check` returns `Ok(None)` on success and `Ok(Some(failure))`
on a counterexample. Generator errors return `Err`. A failure records the exact
case seed, case index, initial value, final candidate and attempted shrink count.
`replay(strategy, failure.seed, size)` reproduces the initial case. Preserve the
generator definition and size as well as the seed. `Failure::report` is available
when the item implements `ToString`.

`check` now also applies a default 1,000,000-unit shrink-work budget. Its existing
signature, `Config` fields and `Sample::new(value, vector_factory)` remain
compatible. Vector candidate ordering changes because chunk deletion precedes
individual element changes; seeds still reproduce the generated initial value.

## Structured runner and coverage

`run(strategy, RunOptions, property)` accepts a property returning `Case`:

```goml
let options = proptest::RunOptions {
    coverage: Vec::from_array([
        proptest::Coverage { label: "nonnegative", minimum_percent: 25 },
    ]),
    ..proptest::RunOptions::new(42),
};
let report = proptest::run(values, options, |value: isize| {
    proptest::Case::pass().classify(value >= 0, "nonnegative")
})?;
```

`Case::pass`, `Case::fail(message)` and `Case::discard(reason)` distinguish
successful cases, counterexamples and unmet preconditions. `label` and `classify`
attach categories; repeated labels on one case count once. Shrink evaluations do
not contribute coverage, discard totals or generated-case totals. Discarded
shrink candidates cannot replace a failing candidate.

`RunOptions::new(seed)` wraps the usual `Config`, allows 1,000 discarded initial
cases and 1,000,000 units of shrinking work, and starts with empty `regressions`
and `coverage`. `from_config` preserves an existing configuration. `config.cases`
counts accepted fresh cases, so discards cause replacement draws with new seeds.
An additional discard beyond `max_discards` returns `RunStatus::Rejected`.

`RunReport[T]` exposes `status` (`Passed`, `Failed` or `Rejected`), `cases`,
`generated`, `regressions`, `discarded`, label/discard-reason counts, `message`,
optional `failure`/`replay`, `shrink_work` and `shrink_limited`. `cases` counts
accepted initial evaluations, including replayed regressions. `generated` counts
fresh draws including discarded ones; `regressions` counts visited saved entries.
The failure's case index addresses the combined regression/fresh evaluation order.
Generator/configuration errors still return `Err`. Callers must inspect `status`;
an `Ok` report with `Rejected` does not establish the property.

Coverage requirements apply after successful execution, count both regressions
and fresh accepted cases, and round the required count upward. Discarded cases
are excluded. Unmet coverage, including a positive requirement with no accepted
cases, returns `Rejected` with a diagnostic.

## Campaigns, distributions and failure aggregation

`run_campaign(strategy, CampaignOptions, property)` continues after failures and
returns a `CampaignReport[T]`. The property returns an `Observation`, created by
`Observation::new(case)` or `case.collect(name, value)`:

```goml
let campaign = proptest::run_campaign(
    values,
    proptest::CampaignOptions::new(42),
    |value: isize| {
        proptest::Case::pass()
            .classify(value == 0, "zero")
            .collect("sign", if value < 0 { "negative" } else { "nonnegative" })
    },
)?;
let rendered = campaign.render(|value| value.to_string(), 24)?;
```

`CampaignOptions.run` contains the compatible `RunOptions`. `max_failures`
defaults to 10 and must be between 1 and 10,000. The limit counts failing initial
evaluations, including saved regressions; failures are not deduplicated by error
message. Every retained `Counterexample` contains its `Failure`, exact `Replay`,
final message, shrink work and whether shrinking reached a limit. Shrink attempt
and work budgets are shared by the entire campaign. When they run out, further
failing cases retain their initial value without expanding their shrink trees.

`summary` retains the original `RunReport` shape and its first counterexample,
with aggregate case/discard/label counts and shrink work. `completed` means all
configured regressions and accepted fresh cases were evaluated; it does not
mean the property passed. An early failure/discard limit sets `completed = false`
and `stop_reason`. A discovered failure keeps the status `Failed` even if later
discards exhaust the budget. Coverage is checked after a campaign without
failures. Generator, configuration and observation errors return `Err`.
The existing `run` and `check` retain their stop-at-first-failure behavior.

`Observation.collect` attaches string-valued categorical distributions. Numeric
values can use `to_string()` or application-selected bins. Histograms count only
accepted initial evaluations, including failures and regressions. Shrinks and
discards do not contribute. A distribution's denominator is the number of cases
that observed it, allowing optional observations. Repeating the same name/value
on a case counts once; conflicting values for one name return an error.
Distributions and buckets preserve first-observed order.

`max_histogram_buckets` bounds distinct `(name, value)` pairs across the campaign;
it defaults to 1,024 and allows 0–10,000. Zero disables observations. Each case
allows at most 10,000 observations, names must be nonempty, and both names and
values are limited to 1,024 UTF-8 bytes. Exceeding these limits returns an error
instead of silently dropping samples.

`RunReport.render(format_value)`, `CampaignReport.render(format_value, width)`
and `Histogram.render(width)` produce deterministic text with replay details,
counterexamples, classifications, discard reasons, counts, integer percentages
and histogram bars. Widths are 0–120. Labels, messages and rendered values are
JSON-quoted to preserve Unicode and distinguish embedded line breaks. The value
formatter supports types without `ToString`; its own work is caller-controlled.
The percentage and bar calculations avoid integer multiplication overflow.

## Executing state machines

`StateMachine[M, S, C]` connects pure model transitions to a real system under
test. Its constructor takes a command limit and five callbacks:

| Callback | Contract |
| --- | --- |
| `setup: () -> Result[(M, S), string]` | Create a fresh model and system for every evaluation, including shrink candidates |
| `transition: (M, C) -> Result[M, string]` | Compute the next model without changing the system; a rejected precondition discards the case |
| `execute: (S, C) -> Result[(), string]` | Apply one command to the system; errors fail the case |
| `invariant: (M, S) -> Result[(), string]` | Check the initial state and every successfully executed command |
| `cleanup: (S) -> Result[(), string]` | Clean up exactly once after successful setup on every normal return path |

`run(commands)` returns `MachineReport`; `check(commands)` returns its `Case`
for direct use with `run`, `run_campaign` or the `stateful` generator. The report
records the stage and zero-based step of a failure/discard, the count of commands
whose execution returned successfully, and any cleanup error. An invariant
failure includes the command that just executed in that count.

The command limit allows 0–1,000,000. An over-limit sequence is discarded before
setup. The helper snapshots the command vector, checks each transition before
executing the command, and stops at the first failure or rejected precondition.
Cleanup failures promote passing/discarded cases to failures; an existing error
and its stage remain available alongside the cleanup error. Setup failures must
clean up partially created resources within the setup callback. Callbacks must
return normally: panics and nontermination are not caught, and model/command
values with mutable internals still require a caller-defined isolation policy.

## Lazy shrinking and budgets

`Sample::lazy(value, factory)` accepts `() -> Candidates[T]`.
`Candidates::from_fn` takes `(ShrinkBudget) -> Option[Sample[T]]`; each factory
must create fresh iterator state. `sample.candidates()` starts an independent
traversal. `next(budget)` charges a unit before invoking the callback and has
permanent exhaustion, including after a budget failure. Copies share traversal
state and are not for concurrent use.

`ShrinkBudget::new(work)`, `consume(work)`, `used` and `exhausted` support custom
shrinkers. All composed shrink streams share the runner's budget. It charges
iterator advances, expansion starts, chunk-search steps and elements copied into
vector candidates. Filtering additionally retains its 1,024-visit expansion cap.
Budget/cap exhaustion returns the best counterexample found, with
`shrink_limited = true`; it does not claim a fully explored local minimum.
Zero work or zero shrink attempts does not invoke the sample's shrink factory.

The runner never calls the compatibility `children()` materializer. Explicit
calls to `children()` collect the candidate stream into a vector and can use
substantial memory. Legacy `Sample::new` callbacks still produce their own vector
on the first requested candidate; adopt `Sample::lazy` to control that allocation.
Budgets cover library-controlled work, not arbitrary allocations or running time
inside application callbacks. Custom callbacks must cooperate with the budget.

## Persistent regression seeds

`Replay { seed, size }` records an initial generation. Entries in
`RunOptions.regressions` run before fresh random cases, even when `config.cases`
is zero. Regression discards consume the same discard budget but are not replaced.

`RegressionStore::new(path, property)` binds a corpus to a property identifier.
`load()` treats a missing file as an empty corpus; malformed content, mismatched
property IDs, unsupported versions and I/O failures return errors. `save(replay)`
deduplicates exact `(seed, size)` entries and atomically replaces the file.
The parent directory must already exist. Coordinate writers externally: atomic
replacement prevents a partial file, but read/merge/write is not a concurrent
append transaction.

`run_persisted(strategy, options, store, property)` loads regressions, runs them
and fresh cases, then saves a failing case's seed and size automatically. A save
failure is reported as `Err`. The standalone `encode_replays` and `decode_replays`
functions use the same versioned text format. Corpora are limited to 10,000 entries
and 1 MiB. Property IDs must be nonempty, trimmed, at most 1,024 bytes and contain
no tab, newline or NUL. Full-width unsigned seeds are preserved.

`run_campaign_persisted` performs the same workflow for campaigns and atomically
saves every collected failure. `RegressionStore.save_all(entries)` merges and
deduplicates an entire batch before replacing the file once. Invalid entries or
an exceeded corpus limit leave the old file intact. It has the same external
writer-coordination requirement as `save`. If generation or observation validation
returns `Err` before a campaign produces its report, no new failures are saved.

Store one property/generator version per file and keep the generator definition
stable for replay. The corpus stores seeds and sizes, not arbitrary serialized
minimal values; changing a generator can change what an old seed produces.

The final candidate is locally minimal under the visited shrink tree and budget;
it is not a proof of global minimality. Properties and custom generators should
be deterministic and should not mutate shared sample values. Custom shrink trees
should make progress. The attempt budget bounds shrink calls to the property but
cannot interrupt a custom generator or property that never returns. Panics are
not converted into property failures.

Closures used with an associated `Strategy::Item` sometimes need explicit
parameter types for numeric operators in the current GoML checker. The examples
and consumer tests exercise this form.

## Validation

```sh
python3 ecosystem/verify.py proptest
```

Tests include exact integer failure boundaries, wide signed ranges, filtering,
dependent invariants during shrinking, recursive data, valid Unicode, finite
floats, invalid configurations and exact seed replay. The expanded suite also
covers large vectors with tiny budgets, lazy composition, rejected shrink work,
unique collections, stateful prefixes, IEEE edge cases, discard/coverage
accounting, persistent replay and corruption preservation. Campaign tests cover
failure limits, shared shrink budgets, distribution cardinality, overflow-safe
formatting and atomic multi-failure persistence. State-machine tests verify
per-step checks, rejected preconditions, reset during shrinking and combined
execution/cleanup errors. A separate module
implements its own associated-type strategy and lazy shrink callbacks and uses
the structured runner, campaigns and a custom model/system through a normally
resolved dependency.
