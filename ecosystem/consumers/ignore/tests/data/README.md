# Independent reference fixture

200 rule hierarchies and 9,400 path queries.

Source: Actual git check-ignore --no-index --stdin -z --verbose --non-matching in isolated repositories, with global/system Git configuration disabled.

Seed: `20260920`. The input and expected values were extracted once from
`ecosystem/ignore/interop.py` at repository commit
`27f8b1649561504bbe61d5247500e7cadc6eef42`, before invoking the GoML
consumer. No expected value was captured from the implementation under test.
Python is not required to run or update native GoML tests; new reference cases
can be added directly from independent calculations or the cited specification.

Fixture SHA-256: `a6b2b6c3c43738b6aa65310f6d70a794eb95d9de7789937a85637d39ce407c57`.

The file uses `indexed-json-v2` to share repeated JSON subtrees. Nodes are in dependency order: `[0, value]` is a scalar; `[1, ids]` is an array; `[2, shape_id, ids]` is an object whose field names come from `shapes[shape_id]`. `root` selects the final decoded node. The native helper in `ecosystem/verification/reference` expands these references and compares objects independently of field order.
