# Library boundaries

`lib/std` contains broadly reusable value types, algorithms and composition
contracts. It depends only on builtin, prelude and other standard packages.
Protocol engines, application policy, large datasets and domain-specific types
live in independently versioned sibling repositories under `../gomlang/`.

Extend existing packages before introducing new owners. Extract a shared module
when a real consumer and its migration test need it. Sharing an implementation
between ecosystem modules does not require moving it into the standard library.

## Ownership

| Standard library | Ecosystem |
| --- | --- |
| `std::bytes`, `std::text`, `std::unicode`, `std::utf8`, `std::num`, `std::collections`, `std::error` | `regexp`, `parser::scanner`, `tabwriter`, `unicode_text`, `template` |
| `std::io`, `std::fs`, `std::testing::fs`, `std::testing::io` | `archive`, `compress`, `object`, `dwarf` |
| `std::encoding`, `std::bytes::endian`, `std::hash`, `std::crypto` | `asn1`, `xml`, `html`, `mime`, `mail`, `x509::pkix` |
| `std::net` address values, `std::net::url`, `std::path::slash` | `http`, `textproto`, `request`, `web` |
| `std::math`, `std::rand` | `bigint`, `bigmath`, `datetime`, `sql`, `tracing`, `image`, `color` |
| Existing `std::json`, `std::toml`, `std::bincode` | Other format engines and `go_doc` |

Standard Unicode scalar classification and casing share one pinned dataset.
Grapheme layout and terminal width belong to `unicode_text`; timezone and
public-suffix datasets remain ecosystem concerns. General URL and IP values
carry no HTTP client policy. Lexical slash paths remain separate from host
filepath operations.

Public APIs and limitations are documented in the [language guide](goml.md) and
individual ecosystem READMEs. The ownership table is not a claim of Go API
compatibility. Go source comment tooling belongs to `go_doc`, independently of
GoML documentation and the independent [gomlgo frontend](https://github.com/gomlang/gomlgo).

## Dependency direction

- `archive` uses `compress` and standard I/O, checksums and endian buffers.
- `image::png` uses `compress::zlib` and standard CRC32; `image` uses `color`.
- `request` and `web` share URL values, MIME, text framing and HTTP helpers.
  `web::proxy` may use `request`; `request` must not depend on `web`.
- `template` and `markdown` share `html`; contextual escaping stays in `template`.
- `regexp` uses Unicode and UTF-8 without adopting lexer-specific matching rules.
- `text` may use `unicode`; Unicode output assembly must not depend on `text`.
- `x509::pkix` uses `asn1`; it does not implement TLS or certificate verification.
- HMAC uses digest algorithms; HKDF and PBKDF2 use HMAC.
- `sqlite` adapts to `sql` contracts; `sql` must not depend on `sqlite`.
- `object` and `dwarf` use byte and random-access I/O interfaces.

## Compatibility

Shared implementations preserve consumer policy. The HTTP-only request URL
contract remains distinct from general URL reference parsing. Template and
Markdown escaping retain their quote and entity rules. Lexical path cleaning is
not archive extraction policy, and HTML entity escaping is not sanitization.

Pure GoML algorithms may use existing allocation, scalar math, I/O and runtime
boundaries. Keep those boundaries explicit; moving an algorithm does not imply a
new runtime, transport or cryptographic security guarantee.

When extracting an implementation, preserve public entry points and test both
existing callers and an independent consumer. Standard-library changes also
need source catalog, dependency, navigation and packaging coverage. Compiler and
driver sources may adopt new APIs only after release and stage0 advancement;
see [repository guidelines](../AGENTS.md#bootstrap-and-language-evolution).

Use [repository checks](../README.md#development) for standard-library changes
and the sibling ecosystem verification repository for its modules. Follow the
[release guide](releasing.md) for publication.
