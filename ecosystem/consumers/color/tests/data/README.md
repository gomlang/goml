# Independent reference vectors

`reference.jsonl` contains 4,659 deterministic cases, one JSON object per line with `input` and `expected` fields. Source: CPython 3.12 `colorsys`, W3C CSS Color 4 rational D65 matrices, published Oklab matrices, 148 CSS named colors, all 34 Sharma/Wu/Dalal CIEDE2000 reference pairs, and 1100-digit Python Decimal arithmetic for 810 alpha-boundary mixtures. Seed: 20260922. Floating tolerances are stored per row where needed; exact alpha is checked separately.

These values were exported once from the independent reference implementation used by `ecosystem/color/interop.py` at repository revision `27f8b1649561504bbe61d5247500e7cadc6eef42`, before removing that helper. No expected arithmetic or rendering result was captured from the GoML implementation under test. The reference source remains available in that historical revision for provenance; running these tests needs neither Python nor NumPy nor a downloaded reference runtime.

SHA-256 of this vector file: `7074c6d4313789b824b79d976036487e242db7d093f6c56adc8c7ad2ba3d2bd7`.

Run the consumer’s ordinary `goml test` after resolving its local ecosystem dependencies. `tests/reference_test.gom` evaluates each case directly through the GoML consumer API and asserts both the case count and results. New behavior should receive independently calculated expected values or a small native reference model; do not regenerate expectations from the implementation being tested.
