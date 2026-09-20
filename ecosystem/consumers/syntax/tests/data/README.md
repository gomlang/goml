# Independent reference fixture

320 tree models, 3,840 persistent edits, and 240 configuration rewrites.

Source: An independent nested-dictionary tree walker, UTF-8 offset model, scalar-boundary checks, and lossless text replacement model.

Seed: `2026092003`. The input and expected values were extracted once from
`ecosystem/syntax/interop.py` at repository commit
`27f8b1649561504bbe61d5247500e7cadc6eef42`, before invoking the GoML
consumer. No expected value was captured from the implementation under test.
Python is not required to run or update native GoML tests; new reference cases
can be added directly from independent calculations or the cited specification.

Fixture SHA-256: `af671e3b61a4f9aaa5f663b94be5aad7f6c47e553c6ac442f3b409593f05ffd4`.

The file uses `indexed-json-v2` to share repeated JSON subtrees. Nodes are in dependency order: `[0, value]` is a scalar; `[1, ids]` is an array; `[2, shape_id, ids]` is an object whose field names come from `shapes[shape_id]`. `root` selects the final decoded node. The native helper in `ecosystem/verification/reference` expands these references and compares objects independently of field order.
