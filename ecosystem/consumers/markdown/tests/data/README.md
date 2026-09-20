# Independent reference fixture

652 CommonMark examples and 2,125 HTML5 named entities.

Source: The CommonMark 0.31.2 official spec.json HTML values and the semicolon-ended mappings in Python 3.12 html.entities.html5, escaped independently for HTML text.

Seed: `not randomized`. The input and expected values were extracted once from
`ecosystem/markdown/interop.py` at repository commit
`27f8b1649561504bbe61d5247500e7cadc6eef42`, before invoking the GoML
consumer. No expected value was captured from the implementation under test.
Python is not required to run or update native GoML tests; new reference cases
can be added directly from independent calculations or the cited specification.

Fixture SHA-256: `eb9054a8b7ac902141f84fb05e1ab232321d3ed320250028cce81f2092a54d07`.

The file contains ordinary JSON `input` and `expected` values. Native tests compare every retained expected result.

CommonMark source: https://spec.commonmark.org/0.31.2/spec.json

Original corpus SHA-256: `d431b29d97b6f73e69d547109cf5081578fac931e72afe95639ebe766c1b2a20`. The specification and examples are by John MacFarlane, licensed under [CC BY-SA 4.0](https://creativecommons.org/licenses/by-sa/4.0/). This fixture preserves each example’s Markdown input and expected HTML and combines them with the entity cases; section/example metadata is omitted. HTML entity data retains the [Python source license](../../../../markdown/LICENSE.entities.txt).
