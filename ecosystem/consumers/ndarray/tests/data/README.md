# Independent reference vectors

`reference.jsonl` contains 2,929 deterministic cases, one JSON object per line with `input` and `expected` fields. Source: NumPy 2.2.6, CPython 3.12, seeds 20260920 for Python Random and NumPy default_rng. Reference wheel: numpy-2.2.6-cp312-cp312-manylinux_2_17_x86_64.manylinux2014_x86_64.whl; SHA-256 fd83c01228a688733f1ded5201c678f0c53ecc1006ffbc404db9f7a899ac6249. QR/LU rows store the original independent inputs; GoML test code checks orthogonality, triangularity, permutation validity and reconstruction with direct scalar loops, not the library’s matrix operations. Special floats are encoded as strings. Tolerances: rtol 2e-10, atol 2e-11.

These values were exported once from the independent reference implementation used by `ecosystem/ndarray/interop.py` at repository revision `27f8b1649561504bbe61d5247500e7cadc6eef42`, before removing that helper. No expected arithmetic or rendering result was captured from the GoML implementation under test. The reference source remains available in that historical revision for provenance; running these tests needs neither Python nor NumPy nor a downloaded reference runtime.

SHA-256 of this vector file: `f72bba75cfa67253b60b1d06e4d33366801e6257b374c6eb044175f39fb2bae4`.

Run the consumer’s ordinary `goml test` after resolving its local ecosystem dependencies. `tests/reference_test.gom` evaluates each case directly through the GoML consumer API and asserts both the case count and results. New behavior should receive independently calculated expected values or a small native reference model; do not regenerate expectations from the implementation being tested.
