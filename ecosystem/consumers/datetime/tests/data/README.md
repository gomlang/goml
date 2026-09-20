# Independent reference vectors

`reference.jsonl` contains 8,140 deterministic cases, one JSON object per line with `input` and `expected` fields. Source: CPython 3.12 `datetime`, `calendar`, and `zoneinfo.ZoneInfo.from_file` using the SHA-256-pinned timezone fixtures in `../../../../datetime/fixtures`; Go 1.25 `time.LoadLocationFromTZData` for nonnegative-epoch synthetic POSIX rules; RFC 9636 fixed all-year-DST expectations. Seed: 221937. The reference implementations’ known differences are documented in the datetime library README.

These values were exported once from the independent reference implementation used by `ecosystem/datetime/interop.py` at repository revision `27f8b1649561504bbe61d5247500e7cadc6eef42`, before removing that helper. No expected arithmetic or rendering result was captured from the GoML implementation under test. The reference source remains available in that historical revision for provenance; running these tests needs neither Python nor NumPy nor a downloaded reference runtime.

SHA-256 of this vector file: `9536abd361bdb03a3b8ffc87074ca198d8b64f71541e5dca066f7683c4d1a6f3`.

Run the consumer’s ordinary `goml test` after resolving its local ecosystem dependencies. `tests/reference_test.gom` evaluates each case directly through the GoML consumer API and asserts both the case count and results. New behavior should receive independently calculated expected values or a small native reference model; do not regenerate expectations from the implementation being tested.
