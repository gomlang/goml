# ecosystem::diff

Generic sequence differences and strict unified text patches, implemented in
GoML. Select Myers' shortest-edit-path algorithm or a linear-space Hirschberg
algorithm, both with deterministic ties and explicit work and workspace budgets.

```goml
use ecosystem::diff;

fn update(old: string, new: string) -> Result[string, string] {
    diff::patch(old, new, "a/file", "b/file", 3)?.to_unified()
}
```

## Sequence and text differences

`sequence(old, new)` accepts slices of any `PartialEq` type. `sequence_by` accepts
a custom comparator and `Options { max_work, max_trace }`. The default budgets
are 10,000,000 search steps and 1,000,000 retained diagonal cells. Exceeding either
returns a recoverable error rather than silently producing a nonminimal result.
Empty-side changes take a direct path. Search costs are proportional to sequence
length and edit distance; this implementation retains quadratic trace space in
the edit distance, with the configured cap.

`sequence_linear` / `sequence_linear_by` compute a shortest edit script using
Hirschberg's LCS divide-and-conquer algorithm. They retain two reusable score
rows of at most `min(old.len(), new.len()) + 1` cells each, an O(log(max(N, M)))
explicit task stack, and the output. Common prefixes/suffixes and one-element
segments use direct paths. There is no recursive call-stack growth or quadratic
retained trace. `max_trace` bounds score-row cells; `max_work` bounds comparisons,
row initialization, split selection and segment processing. Budget exhaustion
returns an error without exposing a partial script.

Worst-case time is O(NM), so this mode is useful when large edit distances would
exhaust the Myers trace. It can choose a different equally minimal script on
ties. `sequence_with`, `line_diff_with`, `character_diff_with` and `patch_with`
accept `Options` and `Algorithm::Myers` or `Algorithm::Hirschberg`. Existing entry
points retain Myers behavior. Algorithm selection is explicit; there is no
automatic retry that silently resets the work budget.

Results are coalesced `Op` records with `Equal`, `Delete` or `Insert` tags and
half-open old/new index ranges. `distance` counts inserted/deleted elements;
`statistics` separates equal, inserted and deleted counts. `invert` swaps the
directions. Comparators should define a stable equivalence relation and should
not mutate their inputs.

`line_diff` preserves line endings, including CRLF and a missing final newline.
`lines` returns those same newline-preserving units. `character_diff` uses
Unicode scalar indices, not bytes or grapheme clusters.

## Unified patches

`patch(old, new, old_name, new_name, context)` builds a `Patch` containing public
`Hunk` and `PatchLine` records. Nearby changes merge when their context overlaps.
Hunk positions are zero-based internally. `Patch::to_unified` writes standard
one-based ranges and missing-newline markers. Identical input produces an empty
patch. `parse_unified` accepts omitted single-line counts, file timestamps and
hunk section descriptions.

`Patch::apply` checks every context/deleted line, range count and destination
position before returning a new string. Conflicts return an error with a source
line; no input file is modified. `Patch::reverse` constructs the inverse patch.
Publicly constructed models are validated too, including overlapping ranges,
inconsistent counts, multi-line entries and misplaced missing-newline markers.

The format API handles one file per patch. It does not perform fuzzy matching,
rename detection, filesystem writes, binary Git patches or three-way merging.
Filenames cannot contain tab or newline characters. Missing-newline markers use
the standard English spelling; generate reference patches with `LC_ALL=C`.

## Validation and example executable

```sh
just ecosystem-test diff
ecosystem/consumers/diff/_artifact/bin/diff produce old.txt new.txt change.patch
ecosystem/consumers/diff/_artifact/bin/diff produce-linear old.txt new.txt change.patch
ecosystem/consumers/diff/_artifact/bin/diff apply old.txt change.patch output.txt
```

The tests exhaustively compare all pairs of binary sequences of lengths zero
through five against an independent dynamic-programming distance oracle (3,969
pairs per algorithm), and validate edit ranges and inverse edits. Workspace tests
cover a large edit distance, exact score-row limits, bounded comparator calls,
trimmed identical inputs and comparator direction when inputs are transposed.
Other tests cover Unicode,
CRLF, no-final-newline files, split hunks, malformed patches and conflict checks.
The separate consumer adds randomized properties through the independently
resolved `ecosystem::proptest` module. The consumer’s ordinary native GoML test checks [48 preserved input pairs](../consumers/diff/tests/data/README.md) with both GoML generation algorithms and GNU `patch`, and verifies application of fresh GNU `diff -u` output in the reverse direction. GNU diffutils and patch must be installed; Python is not required. Temporary files use `ecosystem::tempfile` and are cleaned up after the test.

Algorithm reference: [Eugene W. Myers, An O(ND) Difference Algorithm and Its Variations](https://neil.fraser.name/writing/diff/myers.pdf).
