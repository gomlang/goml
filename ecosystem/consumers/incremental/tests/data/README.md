# Independent reference fixture

324 revision histories, including four conditional-cycle configurations.

Source: A from-scratch recursive graph evaluator with no incremental cache or revision algorithm.

Seed: `20260926`. The input and expected values were extracted once from
`ecosystem/incremental/interop.py` at repository commit
`27f8b1649561504bbe61d5247500e7cadc6eef42`, before invoking the GoML
consumer. No expected value was captured from the implementation under test.
Python is not required to run or update native GoML tests; new reference cases
can be added directly from independent calculations or the cited specification.

Fixture SHA-256: `ea90b7b1a83068878c57a8c419e3288a05b63df1a59b78591f847184345e216b`.

The file uses `indexed-json-v2` to share repeated JSON subtrees. Nodes are in dependency order: `[0, value]` is a scalar; `[1, ids]` is an array; `[2, shape_id, ids]` is an object whose field names come from `shapes[shape_id]`. `root` selects the final decoded node. The native helper in `ecosystem/verification/reference` expands these references and compares objects independently of field order.
