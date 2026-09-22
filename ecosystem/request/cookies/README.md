# Cookie domain policy

`SuffixList::parse` accepts bounded Public Suffix List text and resolves exact,
wildcard and exception rules. `default_suffix_list` embeds an ICANN-and-private
snapshot from [publicsuffix.org](https://publicsuffix.org/list/) so cookie
domain policy does not depend on a system file or a network call at runtime.
Unicode rules are also stored under their literal Punycode A-label spelling,
so ASCII-IDNA hosts use the same suffix boundaries.
The snapshot was fetched from
`https://publicsuffix.org/list/public_suffix_list.dat` and pinned at SHA-256
`e81c6f5f11359a79a2479238e732e08bd8521071fd95ee47053471e3426d7b54`.
Regenerate `builtin.gom` with `python3 ecosystem/request/tools/generate_psl.py`
after updating the expected checksum. The list is licensed under MPL 2.0; see
`LICENSE.MPL-2.0`.

`Jar::set` accepts a single Set-Cookie field with an explicit response host,
request path and Unix timestamp. `Jar::header` selects a Cookie field for a
target host, path and scheme, pruning expired entries on access. `Jar::new`
uses the bundled list; `with_suffixes` accepts a caller-provided list.
