# fuzzy

Unicode-aware ranked subsequence matching in native GoML, inspired by
[nucleo](https://github.com/helix-editor/nucleo). The score is this library's
documented policy, not a claim of numerical compatibility with nucleo or fzf.

```goml
use ecosystem::fuzzy;
use std::context;

fn files(query: string, candidates: Vec[string]) -> Result[fuzzy::Report, fuzzy::Error] {
    let matcher = fuzzy::Matcher::new(fuzzy::Options::new())?;
    matcher.search_parallel(query, candidates, 20, 4, context::Context::background())
}
```

## Matching and offsets

`Matcher::find` performs fuzzy matching. `find_with` also accepts `Substring`,
`Prefix`, `Suffix`, and `Exact` plus a cancellation context. `pattern` and
`prepare` cache pattern normalization and candidate segmentation for repeated
`match_prepared` calls. Prepared values are immutable and reusable concurrently.
They may cross matcher instances when preprocessing policies agree: a candidate
requires the same `path_bonus`, and a pattern requires the same `CaseMode`.
Candidates retain both case representations, so their originating case mode does
not restrict reuse; patterns are independent of path scoring. Incompatible
preparation returns `IncompatiblePreparation`. Budget settings may differ: the
receiving matcher rechecks original UTF-8 byte counts, both candidate unit counts
and the normalized pattern unit count before any empty/impossible-match shortcut.
Reprepare the candidate or pattern when changing its relevant matching policy.

`CaseMode::Sensitive` compares original scalars; `Fold` uses standard Unicode
full case folding; `Smart` folds unless the pattern contains an uppercase scalar.
An expansion such as `ß` to `ss` retains its original grapheme location. Returned
`Match.ranges` are merged half-open UTF-8 byte ranges covering complete extended
graphemes; `graphemes` contains unique zero-based grapheme indices. Combining marks
and emoji sequences are never partially highlighted. Matching itself compares
scalars and can select a scalar inside a grapheme. Canonical normalization,
accent stripping, transliteration, and locale-specific folding are not applied.
Segmentation uses `unicode_text` Unicode 16; case folding/classification use the
standard library's Unicode version.

The dynamic program finds the maximum-score subsequence. Each matched scalar
earns 16 points, with 12 at the input start, 14 after a path separator when enabled,
10 after other non-alphanumeric scalars, or 8 at a lowercase-to-uppercase boundary.
Consecutive matches gain 8. Leading skipped scalars cost 1 each; internal gaps
cost 3 plus their length; trailing skips cost their length divided by four.
Boundary rewards are applied once per original grapheme. Equal scores prefer the
earliest final position, and stable candidate ties preserve input order.

## Search and incremental queries

`search` and `search_parallel` return bounded Top-K hits, total matching candidate
count and scanned count. Limit zero still counts all matches. Parallel workers
produce the same ranking, indices, scores and highlights as serial execution.
Input vectors are copied before parallel work; callers must not concurrently
mutate a vector while the copy is being taken. There are at most 256 workers.

`Session::new(matcher, values)` caches candidate preprocessing. `update(query,
limit, context)` returns `(next_session, report)`; old sessions remain usable.
When a query extends its predecessor with unchanged case policy, only previously
matching candidates are scanned. All eligible candidates are retained regardless
of the previous Top-K limit. Backspacing, replacement and smart-case transitions
rescan the index. Sessions have no mutable shared cache and can branch or be
queried concurrently.

## Bounds and errors

Options validate input byte/unit counts, pattern size, matrix cells and result
count. The defaults are 1 MiB per string, 32,768 normalized candidate scalars,
256 pattern scalars, 2,097,152 DP cells and 10,000 results. Matching uses O(MN)
time and parent storage, plus two O(N) score rows. Top-K uses sorted insertion,
with O(K) insertion cost. Sessions retain O(total candidate size) preprocessing.
Matrix budget exhaustion is an error, never a silently degraded match.

Cancellation/deadlines are checked before matching, each row, every 1,024 columns
and between candidates. Segmentation and case folding of one bounded candidate
are synchronous. `Session::new_with` accepts a context and checks it between
candidates; `new` uses a background context. Searches check the context before
preparing each candidate. A failed update does not modify
the previous session. The scoped parallel search joins every worker on failure.

## Validation

`just ecosystem-test fuzzy` checks library and versioned consumer, repeated builds
and race detection. Native tests cover scores, anchors, Unicode expansions and
graphemes, all budgets, stable Top-K, persistent incremental search, cancellation,
cross-matcher preparation compatibility and budget validation,
parallel equivalence, and 1,270 exhaustive short-input comparisons against an
independent recursive subsequence enumerator. The consumer demonstrates file
search and query refinement. No Python or native adapter is required.
