# Independent reference fixture

2,800 cases: 1,200 SGR sequences, 400 command-stripping inputs, and 1,200 rendered palette/style cases.

Source: An independent ECMA-48 SGR state model and xterm RGB palette distance calculation.

Seed: `20260920`. The input and expected values were extracted once from
`ecosystem/ansi/interop.py` at repository commit
`27f8b1649561504bbe61d5247500e7cadc6eef42`, before invoking the GoML
consumer. No expected value was captured from the implementation under test.
Python is not required to run or update native GoML tests; new reference cases
can be added directly from independent calculations or the cited specification.

Fixture SHA-256: `32ffd9f47b013ff9a7a9f26f4469f284762085207956ff375ff98633edaff023`.

The file contains ordinary JSON `input` and `expected` values. Native tests compare every retained expected result.
