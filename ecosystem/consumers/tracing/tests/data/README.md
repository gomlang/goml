# Independent reference fixture

92 scenarios, including concurrent child work and a 100-task request.

Source: An independent request/span/event lifecycle model. Native tests compare record multisets and separately verify dynamic timestamps, sequence order, and live-span rules.

Seed: `781329`. The input and expected values were extracted once from
`ecosystem/tracing/interop.py` at repository commit
`27f8b1649561504bbe61d5247500e7cadc6eef42`, before invoking the GoML
consumer. No expected value was captured from the implementation under test.
Python is not required to run or update native GoML tests; new reference cases
can be added directly from independent calculations or the cited specification.

Fixture SHA-256: `6df3f7d63cc735c3870a560fc38729437de26217e0da4d43985d774177085f84`.

The file uses `indexed-json-v2` to share repeated JSON subtrees. Nodes are in dependency order: `[0, value]` is a scalar; `[1, ids]` is an array; `[2, shape_id, ids]` is an object whose field names come from `shapes[shape_id]`. `root` selects the final decoded node. The native helper in `ecosystem/verification/reference` expands these references and compares objects independently of field order.
