# Standard-library known answers

[English](KNOWN_ANSWERS.md) | [简体中文](KNOWN_ANSWERS.zh-CN.md)

These native numeric and cryptographic fixtures use known answers captured from independent Go 1.26.0 reference functions before removing their Go FFI declarations. The original references are in commit `2e7afb9f102c78d17c280e445fff39c7c77af476`; numeric references are under `gomlc/testdata/module/project093_std_algorithms`.

| Data file | Distinct records |
| --- | ---: |
| [num/tests/known_answers.tsv](num/tests/known_answers.tsv) | 34,244 |
| [num/tests/conversion_known_answers.tsv](num/tests/conversion_known_answers.tsv) | 13,300 |
| [crypto/hash/tests/sha2_known_answers.tsv](crypto/hash/tests/sha2_known_answers.tsv) | 100 |
| [crypto/hash/tests/sha3_known_answers.tsv](crypto/hash/tests/sha3_known_answers.tsv) | 111 |
| [crypto/hmac/tests/known_answers.tsv](crypto/hmac/tests/known_answers.tsv) | 133 |
| [crypto/hkdf/tests/known_answers.tsv](crypto/hkdf/tests/known_answers.tsv) | 58 |
| [crypto/pbkdf2/tests/known_answers.tsv](crypto/pbkdf2/tests/known_answers.tsv) | 146 |

Each tab-separated record contains an input key and its expected return values. Keys contain the original reference name and pipe-separated arguments. Integer and Boolean fields use decimal and lowercase Boolean text; floating-point inputs use their exact IEEE 754 bits in decimal. String and byte-slice fields use lowercase hexadecimal bytes. String return values use the same hexadecimal representation. Records are sorted by key, and duplicate inputs were required to have identical results. The native loader rejects duplicate or missing keys.

The test still runs its original deterministic inputs, edge cases, and streaming or stateful checks. Updating a test input requires an independently verified expected answer; do not regenerate expected answers from the GoML implementation under test. Running this fixture requires no Go reference implementation or FFI declarations.

## Numeric aggregate answers

The four largest reference streams are represented by fixed FNV-1a 64-bit digests and invocation counts in [aggregate_answers.goml](num/tests/aggregate_answers.goml) instead of millions of bytes of individual vectors. Every original invocation is retained: 283,786 float formatting calls, 358,720 float-to-integer conversion calls, and 29,640 calls to each signed and unsigned integer parser. The native checked helpers feed actual results into those digests; the expected digests came exclusively from the original Go reference outputs, and `numeric_algorithms` verifies all four counts and digests. The existing exact-length, short-buffer, round-trip, width, and error checks remain in place.

The hashed record is the input key, a tab, tab-separated expected values, and a newline, encoded as UTF-8, with the same field encoding as the vector file. Invalid integer parser values are normalized to zero because the original assertions checked only validity on errors. The FNV offset is 14695981039346656037 and the multiplier is 1099511628211, with unsigned 64-bit wraparound. The expected digests fix both the input order and outputs. A mismatch identifies the affected reference stream; individual vector lookups continue to identify their exact input.
