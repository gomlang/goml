# Independent reference fixture

67 scenarios, 3,266 edits, and every final byte/scalar/UTF-16/line boundary.

Source: Independent flat-string Unicode scalar and UTF-16 indexing, CR/LF/CRLF splitting, and byte-boundary editing.

Seed: `20260920`. The input and expected values were extracted once from
`ecosystem/rope/interop.py` at repository commit
`27f8b1649561504bbe61d5247500e7cadc6eef42`, before invoking the GoML
consumer. No expected value was captured from the implementation under test.
Python is not required to run or update native GoML tests; new reference cases
can be added directly from independent calculations or the cited specification.

Fixture SHA-256: `da09c88ab3019b2b3d4eb81138b39a4e49c4a923bb16f92a64aa5b04161dede8`.

The file uses `indexed-json-v2` to share repeated JSON subtrees. Nodes are in dependency order: `[0, value]` is a scalar; `[1, ids]` is an array; `[2, shape_id, ids]` is an object whose field names come from `shapes[shape_id]`. `root` selects the final decoded node. The native helper in `ecosystem/verification/reference` expands these references and compares objects independently of field order.
