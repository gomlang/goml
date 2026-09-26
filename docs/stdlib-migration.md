# Pure GoML standard-library migration

## Scope and placement

Implement the 74 A/B entries from the reviewed Go 1.26 inventory: 37 standard
library capabilities and 37 ecosystem capabilities. These are capability counts,
not counts of new GoML packages. Exclude all R entries, including A+R and B+R,
as well as C entries, unranked experiments and already-covered entries.

The target is idiomatic GoML functionality, not blanket Go API compatibility.
Production algorithms must not delegate to Go standard-library implementations.
Existing allocation/scalar primitives, scalar math, I/O and runtime boundaries
may be used unchanged; no new native backend, runtime feature, TLS engine,
certificate verifier or compiler intrinsic is authorized by this scope.

### Standard library boundary

Use lib/std for broadly reusable, stable value types, small algorithms and
composition contracts. Standard packages must not depend on ecosystem modules.
Extend existing packages rather than mirroring every Go import path.
Logical slash paths stay separate from existing host filepath operations.
Existing std::json, std::toml and std::bincode remain where they are; this work
does not relocate established APIs merely to make the classification uniform.
URL/address values are standard because they are useful without an HTTP stack;
MIME, HTML, XML and ASN.1 engines remain independently versioned ecosystem work.

Checksum algorithms, SHA families and hash-based constructions extend the
existing standard hash/crypto APIs. This does not confer FIPS validation or
compiler-enforced constant-time guarantees. Define contracts for implemented
algorithms, not empty claims of signing or key-exchange support.

### Ecosystem boundary

Use ecosystem for protocol/format engines, application frameworks, large
datasets, domain-specific numerical types and tooling. Their release cadence
and dependencies should not force standard-library/bootstrap churn.
Reuse archive, request, web, template, parser, tracing, bigint, datetime and color.

Introduce shared modules only with a tested consumer. The proposed http module
contains protocol values/helpers only, not another client or server framework.
All paths below specify intended ownership, not already available APIs.

### Placement decisions before implementation

A/B priority determines whether to implement a capability, not whether it belongs
in std. Being part of Go's standard library is not sufficient reason to add it to
GoML's standard library. Use the ownership map below before adding public APIs.

- Standard ownership requires broadly shared semantics, a stable contract and
  dependencies confined to builtin, prelude and other standard packages. A small
  implementation alone does not justify standard ownership.
- Ecosystem ownership is preferred for domain policy, protocol engines, evolving
  formats, application frameworks and independently updated datasets. Pure GoML
  implementation is required in either layer; purity does not imply std placement.
- Unicode scalar classification/casing is a deliberate standard-data exception:
  text and language-facing algorithms need one pinned semantic baseline. Grapheme
  layout and terminal width stay in unicode_text; independently updated timezone
  and public-suffix data do not enter std merely because Unicode tables do.
- Keep existing JSON/TOML/bincode and crypto entry points compatible. Conversely,
  request, web, archive, color and markdown remain ecosystem modules after their
  pure GoML migration; do not promote them as a side effect of removing FFI.
- Split shared mechanics from domain policy only when a real consumer needs the
  split. URL/IP values and byte/I/O contracts belong in std; HTTP framing,
  redirects, cookies, MIME, proxy behavior and extraction policy do not.

Standard owners map to directories under lib/std and ship with the toolchain.
Ecosystem owners map to independent sibling repositories under `../gomlang/` with their own
goml.toml and versioned dependencies; nested owners in the table are packages
inside those modules, not automatically separate registry releases. Each new
module needs a public README, module tests, an independent consumer and ecosystem
verification registration. Shared-module extraction must retain compatible
wrappers and avoid dependency cycles before migrating existing consumers.

Changing an owner requires updating this map, documenting the dependency and
compatibility consequences, and reviewing that decision before implementation.
Do not silently move a capability into std to simplify compiler packaging or
avoid declaring an ecosystem dependency.

## Complete ownership map

### Boundaries for the original five migrations

The five requested pure GoML migrations keep their existing ecosystem ownership.
Extract only the shared foundations listed below; extraction is not permission
to relocate the whole module or expand the inventory into excluded R work.

| Existing module | Standard foundations it may reuse | Implementation that stays in ecosystem |
| --- | --- | --- |
| request | URL/IP values, byte/text codecs, I/O contracts and buffering | HTTP framing, client pools, redirects, cookies, multipart policy and transport orchestration; share MIME/textproto engines with ecosystem peers. |
| web | I/O contracts, URL values and existing resource/panic facilities | Routing, middleware, HTTP serving, reverse proxies and HTTP-specific testing; proxy code may use request, never the reverse. |
| archive | Byte/endian operations, checksums, lexical paths and I/O contracts | TAR/ZIP formats, compression engines and extraction policy; share compression through ecosystem::compress, not std. |
| color | Numeric conversion and scalar math | Color spaces, alpha conversion, palettes and image/terminal conventions; image consumes color without a reverse dependency. |
| markdown | Text, Unicode scalar operations, UTF-8 and generic collections | Markdown parsing/rendering and link policy; shared HTML entities belong to ecosystem::html, contextual escaping to template. |

Standard lexical path cleaning is not safe archive extraction, standard URL
parsing is not request authorization, and HTML entity escaping is not sanitizing
untrusted HTML. Keep these policy contracts explicit in their ecosystem owners.
Existing runtime and transport boundaries are retained; changing their
implementation requires a separate scope decision.

### Placement at a glance

| Standard library: shared foundations | Ecosystem: independently versioned consumers |
| --- | --- |
| Bytes, text, Unicode, numeric conversion, collections and errors | Regex engines, scanners, tabular layout and contextual templates |
| I/O contracts, buffering, portable filesystem interfaces and test adapters | Archives, compression, object formats and debug-information readers |
| Base encodings, endian operations, checksums and hash-based cryptography | ASN.1, XML, HTML, MIME and mail processing |
| URL/IP values, lexical paths, complex arithmetic and explicit random generators | HTTP clients/servers/proxies, cookies and transport tracing |
| Existing JSON/TOML/bincode APIs remain standard | SQL policy, arbitrary-precision numbers, timezone data, images and logging |

This summary does not add capabilities to the inventory. In particular, existing
standard JSON is retained for compatibility; it is not a precedent for moving
every serialization format into std. Compression remains ecosystem-owned even
when several ecosystem consumers need it: shared ecosystem use alone does not
make a capability a toolchain foundation.

### Per-capability implementation checklist

Before implementing each row, record its owner, public contract, direct
dependencies, existing consumers and remaining acceptance work. Use these
decisions to keep the implementation within its assigned layer:

1. Put reusable value semantics and composition contracts in the mapped standard
   package; put protocol state machines, format engines and application policy in
   the mapped ecosystem module. Do not introduce a standard facade that imports
   an ecosystem implementation.
2. Prefer extending an existing owner. Create a shared ecosystem module only
   when an actual consumer and its migration test are included; do not create
   empty modules for every row or mirror Go's package tree mechanically.
3. Establish the standard dependency first, then implement and validate the
   ecosystem consumer against the public API. Keep compiler/driver consumers on
   released stage0-compatible APIs until release and stage0 advancement.
4. Preserve old public entry points during extraction. Record changes to module
   dependencies and version requirements, and test both the existing entry point
   and the new independent consumer before removing duplicate implementations.
5. Track planned, in-progress and validated scope separately. A validated subset
   does not complete its inventory row; unfinished streaming, limits, integration
   or compatibility work must remain visible in the implementation record.

All entries are planned unless the implementation record explicitly says
otherwise. A scope statement is not an assertion of completion.

| Go capability | Priority | Layer | GoML owner | Scope | Wave |
| --- | --- | --- | --- | --- | --- |
| `bytes` | A | std | `std::bytes` | Expand checked byte search/split/replace and views; preserve aliasing rules. | S1 |
| `strings` | A | std | `std::text and string methods` | Fill text operations using shared Unicode data; preserve valid UTF-8 and byte offsets. | S1 |
| `strconv` | A | std | `std::num`; `std::text` | Numeric/radix parsing and formatting stay in num; quoting belongs to text. | S1 |
| `fmt` | B | std | `std::text::format` | Typed formatting built on existing interpolation/traits; no runtime reflection or fmt ABI clone. | S2 |
| `errors` | A | std | `std::error` | Structured causes, wrapping, aggregation and matching; expected failures remain Result. | S1 |
| `sort` | B | std | `std::collections` | Reuse collection algorithms; preserve existing stable-sort contract. | S1 |
| `slices` | B | std | `std::collections`; `Vec/Slice methods` | Checked collection algorithms; use public source helpers before any compiler-owned migration. | S1 |
| `maps` | B | std | `std::collections` | Clone/equality/merge/iteration helpers over existing HashMap; no map-runtime rewrite. | S1 |
| `unicode` | A | std | `std::unicode` | Generated property/case tables with one pinned data version; replace existing Go algorithm calls. | S1 |
| `unicode/utf8` | A | std | `std::utf8` | Incremental/forward/reverse scalar decoding, explicit invalid-byte policies. | S1 |
| `regexp/syntax` | A | ecosystem | `ecosystem::regexp::syntax` | Regex AST, parser, simplification and compilation; reuse logos ideas without changing lexer semantics. | E2 |
| `regexp` | A | ecosystem | `ecosystem::regexp` | Search/captures/replacement/split; explicit linear-time matching contract and budgets. | E2 |
| `text/scanner` | B | ecosystem | `ecosystem::parser::scanner` | Reusable text scanner; keep Go-specific tokenization in gomlgo. | E2 |
| `text/tabwriter` | B | ecosystem | `ecosystem::tabwriter` | Streaming tab alignment; optional terminal-width policy reuses unicode_text. | E2 |
| `html` | A | ecosystem | `ecosystem::html` | Entity escaping/unescaping shared by markdown/template/web; no sanitizer claim. | E1 |
| `html/template` | B | ecosystem | `ecosystem::template` | Add contextual HTML/attribute/URL/JS/CSS escaping to existing engine, not a second Go-template engine. | E2 |
| `io` | A | std | `std::io` | ReaderAt/WriterAt and composable reader/writer adapters; precise partial-progress semantics. | S2 |
| `bufio` | A | std | `std::io` | Extend existing BufReader/BufWriter with bounded scanner and split rules. | S2 |
| `io/fs` | A | std | `std::fs` | Portable read-only filesystem traits and helpers alongside existing OS APIs; not a security sandbox. | S2 |
| `path` | A | std | `std::path::slash` | Pure lexical clean/join/split/base/dir/extension/absolute checks and bounded reusable shell patterns; preserve host filepath behavior. | S1 |
| `encoding/asn1` | B | ecosystem | `ecosystem::asn1` | Bounded DER/OID codec and explicit typed schemas; no implicit reflection. | E3 |
| `encoding/base32` | A | std | `std::encoding::base32` | RFC 4648 standard/hex alphabets, padding options and checked canonical decoding. | S1 |
| `encoding/binary` | A | std | `std::bytes::endian` | Extend existing endian buffers with unsigned/ZigZag varints and composable typed reads/writes. | S1 |
| `encoding/json` | A | std | `std::json` | Extend existing serde/value implementation with streaming I/O, bounds and diagnostics. | S2 |
| `encoding/pem` | A | std | `std::encoding::pem` | PEM framing and headers over Base64; no certificate validation. | S1 |
| `encoding/xml` | B | ecosystem | `ecosystem::xml` | Bounded streaming tokens, namespaces/escaping, then serde integration. | E3 |
| `archive/tar` | A | ecosystem | `ecosystem::archive` | Streaming entry bodies and selected GNU extensions with explicit extraction policy. | E1 |
| `archive/zip` | A | ecosystem | `ecosystem::archive` | ZIP64, random access and streaming decompression; retain existing APIs. | E1 |
| `compress/flate` | A | ecosystem | `ecosystem::compress::flate` | Extract archive's pure DEFLATE engine; incremental streams, dictionaries and compression levels. | E1 |
| `compress/gzip` | A | ecosystem | `ecosystem::compress::gzip` | Extract/reuse archive framing over shared flate; stream composition. | E1 |
| `compress/zlib` | A | ecosystem | `ecosystem::compress::zlib` | RFC framing, Adler32 and preset dictionaries over shared flate. | E1 |
| `compress/lzw` | B | ecosystem | `ecosystem::compress::lzw` | Checked incremental LZW with explicit bit order/code-width conventions. | E1 |
| `hash` | A | std | `std::hash` | Incremental non-consuming digest interface, reset and documented state ownership. | S1 |
| `hash/adler32` | A | std | `std::hash::adler32` | One-shot/update/stateful Adler32 with known-answer and chunking tests. | S1 |
| `hash/crc32` | A | std | `std::hash::crc32` | One-shot/update/stateful CRC32 with reusable polynomial tables. | S1 |
| `hash/crc64` | B | std | `std::hash::crc64` | ISO/ECMA and custom reflected polynomials; same incremental contract. | S1 |
| `hash/fnv` | B | std | `std::hash::fnv` | FNV-1/FNV-1a at 32/64/128 bits; not security hashes. | S1 |
| `math/bits` | A | std | `std::math::bits` | Portable bit operations and checked double-width arithmetic; no new intrinsic dependency. | S1 |
| `math/cmplx` | B | std | `std::math::complex` | Complex value arithmetic and functions using existing scalar math boundary unchanged. | S3 |
| `math/big` | B | ecosystem | `ecosystem::bigint`; `ecosystem::bigmath` | Extend integers; rational/binary arbitrary-precision float in bigmath. Decimal is not a substitute. | E3 |
| `math/rand/v2` | B | std | `std::rand` | Explicit stateful PCG/ChaCha8, unbiased sampling/distributions; preserve splitmix64-v1 outputs. | S3 |
| `time/tzdata` | B | ecosystem | `ecosystem::datetime::tzdata` | Versioned embeddable timezone dataset using existing datetime parser. | E3 |
| `net/netip` | A | std | `std::net` | Extend existing address values with prefixes/zones/classification; no duplicate address type. | S1 |
| `net/url` | A | std | `std::net::url` | Pure URL value parsing, escaping and reference resolution; no DNS/TLS/socket calls. | S1 |
| `net/textproto` | A | ecosystem | `ecosystem::textproto` | Bounded line/header/dot framing shared by HTTP and mail. | E1 |
| `mime` | A | ecosystem | `ecosystem::mime` | Media type/parameter/encoded-word processing with explicit data policy. | E1 |
| `mime/multipart` | A | ecosystem | `ecosystem::mime::multipart` | Shared bounded streaming reader/writer; migrate request multipart via compatible wrappers. | E1 |
| `mime/quotedprintable` | B | ecosystem | `ecosystem::mime::quotedprintable` | Streaming transfer encoding with strict error handling. | E1 |
| `net/http/cookiejar` | A | ecosystem | `ecosystem::request::cookies` | Complete existing store's matching/expiry/public-suffix policy; retain request API. | E1 |
| `net/http/httptest` | A | ecosystem | `ecosystem::web::testing` | Reusable recorder, test server and controlled-failure transports on existing dispatch. | E1 |
| `net/http/httptrace` | B | ecosystem | `ecosystem::request::trace` | Opt-in transport lifecycle events; no new runtime tracing backend. | E1 |
| `net/http/httputil` | B | ecosystem | `ecosystem::web::proxy`; `ecosystem::http` | Reverse proxy on existing request/web; shared protocol-only dumps/header helpers in http. | E1 |
| `net/mail` | B | ecosystem | `ecosystem::mail` | Bounded address/header parsing on MIME/textproto; SMTP is excluded. | E3 |
| `crypto` | B | std | `std::crypto` | Keep crypto namespace and reusable algorithm contracts; no new runtime/FFI crypto backend. | S3 |
| `crypto/sha256` | B | std | `std::crypto::hash` | Pure incremental SHA-224/256; preserve sha256 convenience API. | S3 |
| `crypto/sha512` | B | std | `std::crypto::hash` | Pure SHA-384/512 and standardized truncated variants over shared digest interface. | S3 |
| `crypto/sha3` | B | std | `std::crypto::hash` | Pure SHA-3 and SHAKE with explicit absorb/squeeze lifecycle. | S3 |
| `crypto/hmac` | B | std | `std::crypto::hmac` | Generic hash-based MAC with known-answer tests; no claim of compiler-enforced constant time. | S3 |
| `crypto/hkdf` | B | std | `std::crypto::hkdf` | Extract/expand with RFC output bounds and hash/MAC reuse. | S3 |
| `crypto/pbkdf2` | B | std | `std::crypto::pbkdf2` | Checked iteration/output parameters over HMAC; preserve security backend exclusions. | S3 |
| `crypto/x509/pkix` | B | ecosystem | `ecosystem::x509::pkix` | Names/OIDs/extensions and DER data only; no certificate-chain/trust-store/TLS implementation. | E3 |
| `database/sql` | B | ecosystem | `ecosystem::sql` | Typed rows, transactions and pool policy; integrate existing sqlite without replacing its native backend. | E3 |
| `log/slog` | B | ecosystem | `ecosystem::tracing` | Extend existing attributes/handlers/groups/filtering, no second logging framework. | E2 |
| `testing/fstest` | A | std | `std::testing::fs` | In-memory filesystem and interface contract tests on the new FS abstraction. | S2 |
| `testing/iotest` | A | std | `std::testing::io` | Deterministic short/fragmented I/O and delayed-error fixtures. | S2 |
| `testing/slogtest` | B | ecosystem | `ecosystem::tracing::testing` | Reusable handler contract suite aligned with the existing tracing API. | E2 |
| `image` | B | ecosystem | `ecosystem::image` | Checked pixel buffers, rectangles/subimages and codec interfaces. | E3 |
| `image/color` | B | ecosystem | `ecosystem::color` | Extend existing color types with pixel conversion/alpha conventions. | E3 |
| `image/color/palette` | B | ecosystem | `ecosystem::color::palette` | Reusable immutable palette data and nearest-color queries. | E3 |
| `image/draw` | B | ecosystem | `ecosystem::image::draw` | Clipped copy/compositing with explicit premultiplied-alpha semantics. | E3 |
| `image/png` | B | ecosystem | `ecosystem::image::png` | Bounded PNG codec using shared zlib and CRC32. | E3 |
| `go/doc/comment` | B | ecosystem | `ecosystem::go_doc` | Go comment AST/rendering, separate from GoML docs and gomlgo's frontend. | E3 |
| `debug/dwarf` | B | ecosystem | `ecosystem::dwarf` | Bounded DWARF reader on random-access I/O; no runtime debugger backend. | E3 |
| `debug/elf` | B | ecosystem | `ecosystem::object::elf` | Bounded ELF reader; object namespace allows later formats without putting them in std. | E3 |

## Dependency and delivery order

1. S1: value algorithms, Unicode, encodings/checksums, bit operations and
   logical path/URL/address values. First batch: Base32, varints, digest trait,
   Adler32/CRC32/CRC64/FNV and bit arithmetic.
2. S2: composable I/O, filesystem interfaces, JSON streaming, typed formatting
   and reusable test adapters.
3. E1: extract compression from archive; shared HTTP/MIME/text framing and
   archive improvements. Existing request/web transport backends stay unchanged.
4. E2: regex, scanner/tabular output, contextual template safety and logging.
   Regex can proceed independently once its Unicode prerequisite is available.
5. S3: complex math, explicit random generators and hash-based cryptography.
6. E3: ASN.1/XML/PKIX, mail, SQL policy, big numbers, timezone data,
   image/PNG and binary/debug/documentation tooling.

Key dependencies and cycle constraints:

The inspected request::Url API is an HTTP(S)-only client value: it rejects
credentials, keeps origin/transport fields and applies request query policy.
Do not promote that restriction into the general standard URL value or silently
relax request's contract while sharing parsing. Standard URL work must cover
non-HTTP schemes and relative references independently, then adapt the existing
request entry points with explicit HTTP policy and compatibility tests. Similarly,
extend the existing standard IpAddr instead of adding a second address family.

- archive -> compress -> standard I/O, checksums and endian buffers.
- image::png -> compress::zlib + standard CRC32; image -> color.
- request/web -> shared URL values, mime, textproto and small HTTP helpers.
  web::proxy -> request; request must not depend on web.
- template/markdown -> html; contextual safety analysis remains in template.
- regexp -> standard Unicode/UTF-8; it does not reuse logos's longest-token
  execution semantics as if they were general regex search semantics.
- Unicode-aware text helpers must not introduce text -> unicode -> text.
  Unicode owns private builtin byte-vector output assembly, not the higher-level
  text builder. This retains public unicode entry points and allows text to use
  the same pinned Unicode tables without a dependency cycle or a second dataset.
- x509::pkix -> asn1, not a new TLS or certificate verifier.
- HMAC -> digest algorithms; HKDF/PBKDF2 -> HMAC.
- sql owns pure policies/contracts; existing sqlite adapts to sql, not vice versa.
- object/dwarf consume byte/random-access interfaces, not LLVM.

Retain compatible entry points while extracting shared implementations.
Published ecosystem versions are immutable; never overwrite an existing version.

HTML extraction must preserve the inspected consumer contracts: template escapes
both quotes, uses `&#34;` for double quotes and returns its source-aware bounded
output errors; markdown escapes double quotes as `&quot;`, leaves apostrophes
unchanged and retains its existing string-returning helper. Share mechanics with
explicit policies rather than silently changing either wrapper's output. Markdown's
entity parser requires semicolons and constrains numeric reference length; do not
replace it with permissive HTML text/attribute decoding. Its generated 2,125-name
table and retained CPython-derived data/license provide the extraction baseline,
not evidence of complete semicolon-optional HTML5 decoding. Add missing HTML rules
and fixtures in the shared ecosystem module while retaining Markdown semantics.
The extraction needs tests for both wrappers, an independent HTML consumer and
registry-verification dependency coverage before switching existing consumers.

After the full A/B implementation and acceptance gates pass, notify the user for
review with the complete change summary, acceptance evidence and remaining known
limitations. Do not publish, tag, or push release artifacts until the user explicitly
confirms the review is satisfactory. Address review findings and repeat affected
checks before requesting confirmation again when necessary. Only after approval,
publish the next continuous 0.1.x release using docs/releasing.md: synchronize
versions, pass local
CI, push the release commit, require successful main CI on that exact commit,
then tag and verify the published artifacts. Do not release an incomplete batch
as completion of this goal.

## Acceptance gates

Every delivered capability requires a documented API, recoverable malformed-input
errors, pure production algorithms, known-answer/boundary tests, deterministic
differential tests where useful, and explicit parser/decompression/output limits.
Test chunking, short I/O, concurrency and aliasing where applicable.

Standard changes include source catalog, dependency/link closure, navigation,
packaging and external consumer coverage. Ecosystem changes include manifests,
consumer and verification integration. Format every affected module before tests.
Run focused tests and applicable broad checks; packaging/bootstrap changes need CI.

Preserve established UTF-8, sorting, aliasing and seeded RNG behavior.
Retain compiler/driver fallbacks until public release and stage0 advancement.
A package stub, narrow convenience API or round trips alone do not complete an
inventory item; record residual scope rather than marking it finished.

## Acceptance status

All 74 capability rows are locally implemented and accepted against their
mapped GoML contracts. This records implementation and verification, not a
published release or blanket compatibility with every Go API. User review and
explicit confirmation are still required before publication; self-hosted
consumers retain their fallbacks until release and stage0 advancement.

Public contracts and limitations live in the [language guide](goml.md) and the
individual ecosystem module READMEs. Standard-library regression coverage lives
in [module fixtures](../gomlc/testdata/module/), with source catalog, dependency
and navigation checks in [query tests](../gomlc/query/). Use `just verify-golden`
and `just ci` for current standard-library results and the sibling ecosystem
verification repository's `just ecosystem-test` for its modules. Release checks
follow [the release guide](releasing.md).
