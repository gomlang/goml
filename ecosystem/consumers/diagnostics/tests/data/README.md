# Independent reference vectors

`reference.jsonl` contains 6,236 deterministic cases, one JSON object per line with `input` and `expected` fields. Source: Independent byte/scalar-boundary, physical-line, tab/controlled-grapheme-width and edit-conflict models authored for these tests. Seed: 20260921. 2,976 location, 1,440 span and 1,500 edit results are frozen reference outputs; 320 render rows store inputs/property constraints, checked natively for width, label/caret alignment and plain/ANSI equivalence.

These values were exported once from the independent reference implementation used by `ecosystem/diagnostics/interop.py` at repository revision `27f8b1649561504bbe61d5247500e7cadc6eef42`, before removing that helper. No expected arithmetic or rendering result was captured from the GoML implementation under test. The reference source remains available in that historical revision for provenance; running these tests needs neither Python nor NumPy nor a downloaded reference runtime.

SHA-256 of this vector file: `7695bb1e1ad1d7df15b0ef8b8dedad2de44767006bd6cebe14b121c2b4b3b0c1`.

Run the consumer’s ordinary `goml test` after resolving its local ecosystem dependencies. `tests/reference_test.gom` evaluates each case directly through the GoML consumer API and asserts both the case count and results. New behavior should receive independently calculated expected values or a small native reference model; do not regenerate expectations from the implementation being tested.
