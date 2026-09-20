# Independent reference fixture

2,500 mutations and 2,505 frames over five screen dimensions.

Source: An independent cell-grid model of wide-cell erasure, overlap, and fill. Native tests replay emitted ANSI into a separately implemented screen.

Seed: `94317`. The input and expected values were extracted once from
`ecosystem/tui/interop.py` at repository commit
`27f8b1649561504bbe61d5247500e7cadc6eef42`, before invoking the GoML
consumer. No expected value was captured from the implementation under test.
Python is not required to run or update native GoML tests; new reference cases
can be added directly from independent calculations or the cited specification.

Fixture SHA-256: `c0d878418b50659927b4fb0ee2434ec388356c9d04a817dab231483000e506ae`.

The file uses `indexed-json-v2` to share repeated JSON subtrees. Nodes are in dependency order: `[0, value]` is a scalar; `[1, ids]` is an array; `[2, shape_id, ids]` is an object whose field names come from `shapes[shape_id]`. `root` selects the final decoded node. The native helper in `ecosystem/verification/reference` expands these references and compares objects independently of field order.
