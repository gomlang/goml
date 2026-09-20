# Independent reference fixture

160 LRU histories, 42,240 operations, and 20,303 queries.

Source: An independent OrderedDict model with separately computed TTL/TTI, entry and weight eviction.

Seed: `20260920`. The input and expected values were extracted once from
`ecosystem/cache/interop.py` at repository commit
`27f8b1649561504bbe61d5247500e7cadc6eef42`, before invoking the GoML
consumer. No expected value was captured from the implementation under test.
Python is not required to run or update native GoML tests; new reference cases
can be added directly from independent calculations or the cited specification.

Fixture SHA-256: `11f88c093a96b9b3880380c837fdfa65dff0e0efb93042c67772d5c0a73ddeb0`.

The file uses `indexed-json-v2` to share repeated JSON subtrees. Nodes are in dependency order: `[0, value]` is a scalar; `[1, ids]` is an array; `[2, shape_id, ids]` is an object whose field names come from `shapes[shape_id]`. `root` selects the final decoded node. The native helper in `ecosystem/verification/reference` expands these references and compares objects independently of field order.
