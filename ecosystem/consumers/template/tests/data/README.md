# Independent reference fixture

1,367 shared-syntax cases, including 720 Unicode slicing cases.

Source: Jinja 3.1.6 with MarkupSafe 3.0.2, StrictUndefined, keep_trailing_newline=True, and explicit per-case autoescape.

Seed: `20260920`. The input and expected values were extracted once from
`ecosystem/template/interop.py` at repository commit
`27f8b1649561504bbe61d5247500e7cadc6eef42`, before invoking the GoML
consumer. No expected value was captured from the implementation under test.
Python is not required to run or update native GoML tests; new reference cases
can be added directly from independent calculations or the cited specification.

Fixture SHA-256: `d8f4d71c6b07e15058e9e5aa8745cd39a89778996ac8c4674b0d7af766a37496`.

The file uses `indexed-json-v2` to share repeated JSON subtrees. Nodes are in dependency order: `[0, value]` is a scalar; `[1, ids]` is an array; `[2, shape_id, ids]` is an object whose field names come from `shapes[shape_id]`. `root` selects the final decoded node. The native helper in `ecosystem/verification/reference` expands these references and compares objects independently of field order.

Reference distributions: [Jinja 3.1.6](https://pypi.org/project/Jinja2/3.1.6/) wheel SHA-256 `85ece4451f492d0c13c5dd7c13a64681a86afae63a5f347908daf103ce6d2f67`; [MarkupSafe 3.0.2](https://pypi.org/project/MarkupSafe/3.0.2/) sdist SHA-256 `ee55d3edf80167e48ea11a923c7386f4669df67d7994554387f84e7d8b0a2bf0`. Only independently generated input/output data is retained, not either implementation.
