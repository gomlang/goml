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
Ecosystem owners map to independent modules under ecosystem with their own
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

## Current acceptance status

Of the 74 capability rows, 42 are accepted against their mapped GoML contracts,
none is started, and 32 remain planned-only.
Acceptance here means local implementation and verification, not a published
release or blanket compatibility with every Go API. All 74 rows remain in the
release objective; the migration and publication are not complete.
User review and explicit confirmation are required after full implementation and
before publication; local acceptance of individual rows is not release approval.

### Accepted: net/http/httputil in ecosystem::web::proxy and ecosystem::http

The mapped GoML contract includes a bounded reverse proxy on the existing web
server and request client, plus transport-independent HTTP/1.1 dump views and
hop-by-hop header filtering. The proxy joins an upstream base path, retains the
raw query, forwards request and response bodies under explicit limits, passes
backend redirects through, preserves repeated end-to-end headers, replaces
untrusted forwarding headers, and propagates inbound cancellation. Upgrades,
tunnels, streaming bodies, trailers and 1xx forwarding remain outside this
buffered transport contract. Independent HTTP and web consumers exercise the
public packages. Isolated module/consumer format, build, tests, execution,
cached-build and race checks passed for both modules. This is local acceptance,
not publication or Go API parity.

### Accepted: net/http/httptrace in ecosystem::request::trace

The opt-in bounded trace records request starts, connection attempts and reuse,
HTTP/1.1 writes or HTTP/2 queueing, buffered responses, retries, redirects and
terminal outcomes. Copies share a concurrency-safe collector; request IDs
correlate events, and old events are discarded at capacity. The existing
networking API does not expose separate DNS/TLS timing hooks, so this is not a
claim of Go hook parity or a new tracing backend. The external request consumer
uses the public builder and checks lifecycle events on live HTTP requests.
Request's isolated verification passed library and consumer format, build,
tests, execution, cached build and race checks, including native race checks.
The root request integration remains uncommitted pending its package-level
commit; this local acceptance does not imply publication.

### Accepted: encoding/json in std::json

The mapped contract now includes bounded incremental Read/Write, tokens,
iterative Value traversal and direct typed serde without an intermediate Value
tree. Existing string and serde/value entry points remain compatible; new
single-value reader/writer and string conveniences require explicit limits.
Self-hosted compiler/driver consumers have not migrated to the new APIs.

| Contract | Implementation and evidence |
| --- | --- |
| Incremental grammar and strict scalar validation | Shared cursor/lexer/token state; every chunk size for Unicode typed data, refill boundaries, malformed UTF-8/escapes/numbers, EOF and root separation. |
| Direct typed serde and iterative Value handling | Optional, tuple, sequence, map, struct and enum protocol frames; unknown-field skipping, numeric widths, duplicate/missing fields, mixed completed roots and explicit deserialize_any. |
| Bounded work and diagnostics | Input/output/root bytes, depth, values, scalar lengths, entries, paths, serde frames and event quotas; structured offsets/paths and recoverable failures, including cyclic Value input. |
| Provider and alias safety | Primitive I/O only; short/zero/invalid counts, interrupted/failing calls, confirmed progress, sticky errors, reentrant aliases, expired adapters and panic cleanup. |
| Independent behavior checks | Go encoding/json acceptance/token differential checks on known answers and deterministic mutations; legacy typed encoder comparisons and custom serde protocol fault injection. |
| Public integration | All 18 JSON sources registered, source-only resource closure, external fixture use, completion/navigation tests and documented compatibility boundaries. |

Evidence is in `_artifact/json-only.log`, `_artifact/json-ci.log` and
`_artifact/json-golden.log`. Full CI passed 174 driver tests, fixed-point,
packaging, release smoke, editor and script checks. Of 1,032 compiler tests,
1,031 initially passed; the sole failure was the expected JSON IR snapshot
drift. `just update-golden` regenerated the five affected snapshots and passed
29 formatter, 17 pipeline and 51 integration tests. The original
`pipeline_shard_11` then passed again with UPDATE_EXPECT unset (exit 0).
No production code changed after that CI run.

Limits are logical bounds, not allocator or blocking-provider guarantees;
concurrent alias access is unsupported. Typed recursive user serde still uses
its own stack. Strict UTF-8/surrogates, exact serde field matching, array-shaped
bytes, ordered/duplicate-preserving Value objects and whitespace-separated
stream roots are intentional documented differences from Go's reflection APIs.

### Accepted: fmt in std::text::format

The mapped typed-formatting contract covers bounded text and quoting, Unicode
scalar padding, integer signs/prefixes/digits, separate f32/f64 formatters,
ToString/Debug adaptation and primitive Write with confirmed progress. Shared
aliases cannot reset budgets, snapshots are independent, and failed library
fields do not partially append. User conversion side effects are not rolled back.
The implementation reuses num/text and does not add reflection or f-string syntax.

All four external fixture families pass alongside package-selection, 11 resource
and 95 query-namespace tests. Generated snapshots pass 29 formatter, 17 pipeline
and 51 integration tests. Final CI passes 1023 compiler and 171 driver tests,
bootstrap fixed point, packaging and smoke checks. Evidence is under
_artifact/stdlib-format-*.log. This is the 27th accepted capability; older
under-validation notes below are superseded. Publication still requires the
entire inventory and user review.

### Accepted: testing/fstest in std::testing::fs

The public read-only MemoryFS provides bounded construction, immutable input
snapshots, implicit directories, metadata, link resolution, independent opens,
sequential/positional reads, seeking and paged directory handles. External
composition tests cover helpers, nested SubFs, walk and glob.

The recoverable checker covers primitive file/directory contracts, stable
snapshots, invalid logical paths, expected trees, bounded reports and exactly-once
cleanup. Static Seek, ReadAt and link entry points add capability checks without
runtime interface discovery. Helper/SubFs cross-checks meter underlying calls
and confirmed bytes using the same budgets; compound failures preserve primary
errors, independent cleanup failures and the first stop reason.

All four memory and eight checker fixture families pass. Resource tests (10),
query-namespace tests (95), ecosystem csv/msgpack/rope/archive checks and generated
snapshots pass. Final CI passes 1021 compiler and 171 driver tests, bootstrap
fixed point, packaging and smoke checks. Evidence: _artifact/stdlib-fs-check-ci.log,
_artifact/stdlib-fs-check-golden.log and _artifact/stdlib-fs-check-helpers-*.log.
This is the 26th accepted capability. This acceptance supersedes the historical
outstanding notes below; it is not publication approval or Go ABI compatibility.

### Accepted: io/fs in std::fs

The portable filesystem row covers read-only handles, bounded read/stat helpers,
paged directory collection, namespace composition, iterative traversal, globbing,
link protocols, portable modes, snapshot entries and bounded display helpers.
The whole-row matrix below records the public contracts and external fixtures.
Host metadata remains limited by its existing boundary; this is not a sandbox
or a claim of Go optional-interface ABI compatibility.

Final validation passes 1015 compiler and 171 driver tests, bootstrap fixed-point,
packaging and smoke checks. The external fixture, 97 query tests, 10 resource
tests, ecosystem checks and generated snapshots also pass. Evidence is under
_artifact/stdlib-fs-metadata-*.log. This is the 25th accepted capability.
The separate testing/fstest row has begun with a public memory provider; its
recoverable contract checker and complete validation remain outstanding.

### Accepted: io and bufio in std::io

The mapped rows now cover positional and sequential transfer, composable
adapters, checked seeking, confirmed partial progress, synchronous pipes,
buffered byte/scalar/delimiter operations and bounded standard/custom scanning.
The whole-row family-to-fixture matrix in the synchronous-pipe implementation
record documents their coverage and deliberate GoML contract differences.
Historical notes about missing io/bufio families below are superseded by this
acceptance, not removed from the development history.

All production additions are GoML source over existing allocation, channel and
I/O primitives. Public resources and navigation are integrated without moving
self-hosted compiler/driver consumers onto unreleased APIs. Final validation
passes the complete project093 external fixture, its race-instrumented run,
97 query tests, csv/msgpack/rope/archive library and independent-consumer checks
including configured race runs, and generated snapshots. Full CI passes 1010
compiler and 171 driver tests, bootstrap fixed-point, packaging and smoke checks.
Evidence is recorded under _artifact/stdlib-io-pipe-*.log.

These are the 23rd and 24th accepted capabilities. This does not accept io/fs,
promise every Go interface/optimization ABI, or authorize publication. The full
74-row implementation and subsequent user review/confirmation remain required.

### Accepted: html in ecosystem::html

The independently versioned HTML utility module implements bounded escaping,
explicit quote policies, shared named-entity lookup and bounded text/attribute
character-reference decoding. All 2,125 pinned names and 106 semicolon-optional
legacy names are covered, including multi-scalar replacements, longest matches,
attribute ambiguity, numeric recovery, C1 mappings and output/error limits.
The independent consumer compares normal cases against Go EscapeString and
UnescapeString, with explicit HTML-state expectations for Go's short-decimal,
empty-hex and integer-wraparound quirks. Numeric controls/noncharacters are not
silently deleted. The README links the normative reference states and records
the deliberate differences and whole-string resource limits.

Template and Markdown retain their compatible wrappers and explicit policies;
Markdown keeps its stricter semicolon requirement. Production algorithms are
pure GoML with checked-in data; Go reference calls exist only in the independent
consumer. Web's existing Response::html accepts already-constructed markup;
it has no separate entity engine to migrate, and must not silently escape the
entire markup response. Its callers can use this shared utility explicitly.
Manifests, registry verification and downstream consumers are in place.
Fresh _artifact/stdlib-html-acceptance.log verification passes html, template,
markdown and tui_markdown library/consumer tests, formatting, cached builds and
execution, plus entity-data generation checks and the terminal PTY check.

This accepts the mapped html utility row, not HTML sanitization, an HTML parser
or template contextual analysis. Ecosystem ownership is unchanged. At that
acceptance, counts were 22 accepted, 2 started (io and bufio) and 50 planned. The full A/B
implementation, user review and release gates remain outstanding.

### Accepted: strings in std::text and existing string/bytes/cmp APIs

The mapped text capability covers byte-offset and scalar/predicate searching,
comparison, cuts and trimming, bounded and lazy splitting/fields/lines, joining,
repetition, replacement, scalar mapping/casing/folding, and an explicit tested
legacy word-title recipe. StringBuilder supports checked appends, reservation,
UTF-8 byte validation, shared handles and immutable snapshots; clone_checked
provides bounded detached copying. Compiled Replacer supplies ordered rules,
fallible incremental fragments and opt-in writer output. StringReader belongs
to std::io and supplies sequential/positional reads, byte/scalar rollback, checked
seek/reset and confirmed-progress writer transfer without a text-to-I/O dependency.

Final comparison against local Go 1.26 strings sources found no remaining mapped
functional family. Existing public API compositions are directly tested instead
of adding redundant Go-named aliases. Arbitrary byte search and invalid UTF-8
repair use bytes; text accepts valid UTF-8 only. Empty-pattern policies, reader
rollback/seek differences, shared builder/reset semantics, and replacement scalar
boundaries remain explicit GoML contracts, not blanket Go compatibility. The
legacy title recipe is caller policy, not general linguistic segmentation.
Resource limits are logical bounds, not total-memory or allocator-failure promises.
Failed writes may have unconfirmed side effects and are not automatically replayed.

The standard production implementation passed full CI in
_artifact/stdlib-text-io-safety-ci.log (1007 compiler and 171 driver tests,
bootstrap fixed point, helpers and packaging). Final composition/title tests
passed project093 and golden regeneration (29 formatter, 17 pipeline and 51
integration tests), recorded in _artifact/stdlib-text-composition-focused.log and
_artifact/stdlib-text-composition-golden.log. Related CSV write-policy correction
passed independent consumer, cache and race verification in
_artifact/stdlib-text-composition-csv.log. This accepts strings only: io, bufio and
HTML remain started, and 50 rows remain planned. Counts are 21 accepted, 3 started
and 50 planned. Compiler/driver fallbacks remain until release and stage0
advancement; no publication or user-review gate has been bypassed.

### Accepted: strconv in std::num and std::text

The mapped conversion capability covers boolean and radix/width-aware integer
conversion, f32/f64 parsing and all decimal/binary/hexadecimal formatting modes,
checked scalar casts and rounding, complex component parsing/formatting, three
quotation modes, raw-literal eligibility, complete/prefix/element unquoting, and
checked destination/builder composition. The detailed conversion acceptance
matrix below records the production owners, external tests and contract limits.
Existing Unicode classification is reused; no duplicate data or native algorithm
backend was introduced. Complex arithmetic remains a separate planned capability.

Final full CI passes in `_artifact/stdlib-num-final-ci.log`: all 1006 compiler
tests, 171 driver tests, helper race checks, bootstrap fixed point and packaging
checks. This accepts the mapped GoML contract, not Go API/ABI identity or removal
of stage0-compatible compiler/driver fallbacks. Input/work, allocation, UTF-8 and
intentional Go-reference differences remain explicitly documented. Counts are
20 accepted, 4 started and 50 planned; the full A/B review/release gate remains.

### Accepted: bytes in std::bytes

The mapped byte capability provides reusable linear-time pattern search with
owned search tables; byte/scalar/set/predicate search; cut/prefix/suffix views;
nonoverlapping count, bounded eager and lazy split variants; newline and field
iteration; checked join/repeat/replace; edge trimming, scalar mapping/casing and
simple-fold equality; and bounded malformed UTF-8 run replacement. Unicode rules
reuse the pinned standard tables, while raw pattern operations preserve arbitrary
binary data. Existing Bytes/Builder shared storage, explicit copies and immutable
FrozenBytes snapshots remain compatible.

All production algorithms are pure GoML. The package's sole standard dependency
is Unicode; it does not import io, text, utf8, host or ecosystem code. Memory I/O
uses the existing io::Cursor/BufReader/BufWriter contracts, including checked total
output limits, zero-progress rejection, overlapping write snapshots and buffered
retry. These are composition APIs, not a second Go bytes.Buffer object model.

External tests cover exhaustive binary-alphabet searches, adversarial repetitive
patterns, every one-/two-byte decoder-sensitive input, multibyte mutations,
Go differential split/replace and casing, integer-extreme counts, exact/insufficient
limits, callback ordering, iterator exhaustion/shared state, live views and detached
output. The final audit adds 784 overlapping part/separator construction cases.
Immutable snapshots and original byte-copy contracts are also covered by project089.

Public sources, dependency/link isolation, completion/navigation, documentation
and examples are present. `_artifact/stdlib-bytes-acceptance-focused.log` and
`_artifact/stdlib-bytes-acceptance-compiler.log` pass (1006 compiler tests). The
unchanged production baseline passed complete CI in
`_artifact/stdlib-cursor-limit-ci.log`, including driver tests, helper race,
fixed point and packaging; this final audit changes only tests and records.
Explicit limits bound returned bytes/parts, not all allocations or callback work.
Shared mutable inputs require serialized access; replacement-decoding policies
are intentionally separate from strict UTF-8 validation. This accepts the mapped
capability, not blanket Go API/ABI compatibility or completion of the full A/B goal.

### Accepted: unicode in std::unicode

The dependency-free standard package provides pinned Unicode 15.0.0 scalar
classification, named category/script/property/fold tables, validated immutable
custom ranges, simple upper/lower/title mappings, simple-fold cycles, special
case overrides and locale-independent multi-scalar full folding. Lookup and
conversion are pure GoML; native libraries are only generation/test references.
The shared pinned baseline is reused by text and bytes without dependency cycles.

External tests compare every valid scalar's classifications, simple mappings,
simple fold and Turkish mappings with Go. Every named table's ranges and boundary
membership are checked, including surrogate inspection, maximum stride, immutable
snapshots and invalid custom tables/mappings. Public contracts distinguish scalar
mapping from full folding and code points from valid character values.

Both seven-table generation and full casefold generation now have non-mutating
canonical --check modes in CI, with explicit version rejection and provenance/
license documentation. Full-fold stale-data rejection was additionally exercised
with write operations disabled. No generated data changed in this final audit.

Public source packaging, dependency selection/purity, docs/examples and editor
navigation are covered. Full CI passed in `_artifact/stdlib-unicode-acceptance-ci.log`:
1,006 compiler tests, 171 driver tests, helper race, both generation checks,
bootstrap fixed point and packaging/smoke checks. This accepts the mapped Unicode
tables/casing capability, not grapheme layout, terminal width or normalization.
The full A/B implementation and user-review/publication gates remain active.

### Accepted: unicode/utf8 in std::utf8

The mapped codec provides pure scalar validation/encoding, checked forward and
replacement forward/reverse decoding, completeness/start/scalar predicates,
malformed-byte-aware counting, lossy conversion and strict incremental decoding.
Checked destination writes preserve buffers on range errors; incremental state
retains at most four bytes and guards consumed-offset overflow. Existing string
allocation/scalar primitives remain unchanged and no native algorithm is added.

Public contracts distinguish a real replacement scalar from a one-byte invalid
sequence, incomplete from malformed input, and sticky failures from successful
finish. Shared aliases/reset and caller synchronization are explicit. Existing
tests cover every one/two-byte input, sampled valid scalars throughout Unicode,
random binary data, encoding guard bytes and streaming reset/finish behavior.
The final audit adds 57,600 multibyte boundary combinations, valid-prefix offsets,
reverse walks and independently derived Go DecodeRune/FullRune error positions.
No production defect remained after this audit.

Catalog/dependency selection, source purity, public documentation/examples,
completion/navigation and external consumers are present. Focused execution and
all 1,006 compiler tests passed in
`_artifact/stdlib-utf8-boundaries-{focused,compiler}.log`; unchanged production and
catalog sources were covered by `_artifact/stdlib-url-acceptance-ci.log`.
This accepts UTF-8 encoding/decoding, not grapheme segmentation, Unicode property
tables, normalization or implicit invalid-byte repair. Those capabilities retain
their separately assigned ownership and acceptance state. Release remains gated.

### Accepted: encoding/base32 in std::encoding::base32

The public pure codec covers standard/hex alphabets, padded/raw variants,
validated reusable custom ASCII alphabets/padding, strict canonical decoding,
checked encoded lengths and output budgets. Incremental Encoder/Decoder expose
shared state, cumulative limits, transactional updates, recoverable finish and
reset with independent returned output. The separate stream child composes
Read/Write without adding I/O dependencies to the core codec.

Known-answer and deterministic Go comparisons cover configurations, binary data,
quantum boundaries and chunking. Failure tests cover invalid alphabets, symbols,
padding/trailing bits, integer overflow, exact budgets, retry after failures,
partial destination writes, sticky I/O failures and trailing input validation.
The final audit tests every two-part split of payload lengths zero through sixteen
for all four variants, injecting a late error after potentially produced local
output, then verifying complete rollback, correct absolute offsets, successful
retry/reset and independent buffers. No production defect remained after review.

The guide explains stricter canonical/whitespace rules, update atomicity versus
non-rollbackable I/O, finish/flush/close ownership and cumulative output limits.
Catalog, dependency isolation, driver selection, completion, core update and
stream navigation are covered. Final focused execution and all 1,006 compiler
tests passed in `_artifact/stdlib-base32-transaction-{focused,compiler}.log`;
unchanged production/catalog sources were covered by
`_artifact/stdlib-url-acceptance-ci.log`. This accepts the mapped Base32 capability,
not new runtime encoders or arbitrary memory quotas. Release approval remains gated.

### Accepted: encoding/pem in std::encoding::pem

The mapped framing/header capability provides validated immutable Block values,
independent mutable data/header snapshots, bounded canonical encoding and bounded
first-candidate decoding with consumed offsets for successive blocks. Encoding
sorts headers with Proc-Type first, wraps padded Base64 at 64 columns and preflights
the complete output size. Decode independently limits raw input, decoded bytes
and header count and returns recoverable errors without partial blocks.

Deterministic Go comparisons span binary payloads through multiple line boundaries,
header ordering, LF/CRLF, whitespace, preambles and missing final newlines. Tests
also cover alias isolation, multiple blocks, exact limits, invalid labels/headers,
duplicate keys and strict Base64 trailing bits. Final auditing found and fixed
Unicode trimming that could hide invalid non-ASCII header bytes; regression tests
now enforce ASCII-only space/tab normalization before validation.

The public contract explicitly distinguishes fail-fast malformed candidates from
Go's searching decoder, valid UTF-8 textual input from arbitrary byte preambles,
and framing from cryptography, ASN.1 or trust validation. This row does not promise
a streaming transport adapter. Source catalog, dependency selection, navigation,
public docs and an external consumer are present, with no Go FFI in production.
Incremental toolchain build, focused execution and all 1,006 compiler tests passed
in `_artifact/stdlib-pem-audit-{build,focused,compiler}.log`. No catalog or bootstrap
contract changed in the fix. Full-inventory review and release approval remain open.

### Accepted: errors in std::error

Report[E] adds typed leaf errors, immutable context wrappers and ordered aggregate
causes to the existing Error marker, ErrorKind and Details protocol. It supports
immediate inspection, first-match depth-first search, typed projection and
equality-based membership. Joins and returned cause vectors copy their structure;
leaf values retain ordinary shallow/shared semantics. Empty joins return None,
duplicates remain meaningful and display/search use iterative stacks.

Public examples and external tests cover nested/empty/singleton reports, Unicode,
10,000 context layers, cause-vector isolation, mutable leaves and ordered repeated
subtrees. Domain-enum tests combine io::Error with application validation errors,
preserve kind and OS code through typed projection, propagate through map_err/?,
and use the generic Error protocol without reflective downcasts. Callback work,
tree size and rendered output are caller-controlled, not implicitly budgeted or
sanitized; arbitrary error fields inside a leaf are not automatically traversed.

The error source remains a standard catalog entry without new native or ecosystem
dependencies; completion and context/find_map navigation are verified. Focused
execution and all 1,006 compiler tests passed in
`_artifact/stdlib-errors-acceptance-{focused,compiler}.log`. The unchanged production
and catalog sources were covered by `_artifact/stdlib-url-acceptance-ci.log`.
This accepts the idiomatic typed errors contract, not Go reflection-based Is/As,
new panic machinery or deep freezing of mutable error payloads. The whole-inventory
review and explicit publication approval are still required.

### Accepted: slices in std::collections and existing Vec/Slice methods

The mapped APIs cover view equality/comparison/membership, checked concat/repeat,
overlap-safe vector range edits, stable predicate deletion, checked reserve,
reverse and stable view sorting, sortedness/search/extrema, bounded-view iteration,
zero-copy chunking and generic iterator collection/append/sorted collection.
Existing clone/dedup/value-iterator methods remain canonical; no compiler fallback
is replaced. Go capacity clipping has no GoML view equivalent and independent
length-sized copies use to_vec, as explicitly documented.

Invalid ranges/counts and arithmetic overflow return SequenceError before edit
mutation. Allocation failure recovery is not promised; arbitrary iterator resource
limits and callback synchronization remain caller responsibilities. Views read
shared backing storage, copies are shallow, Vec aliases observe edits and indexed
iterator aliases share cursor state. Ordered helpers require Ord or explicit
comparators rather than silently imposing float/NaN ordering.

Existing deterministic Go differential tests cover equality, comparison, indexing,
repeat, replacement, deletion and compaction across small ranges, plus ordering,
search, iterator exhaustion and maximum-length arithmetic. Final tests add 2,892
self-overlap combinations, live chunk partitions, alias/cursor behavior, callback
counts and stable deletion order. No production defect remained after this audit.

The collections source catalog, dependency closure, public guide, examples,
completion and checked-edit/chunk navigation are covered. Focused execution and
all 1,006 compiler tests passed in
`_artifact/stdlib-slices-acceptance-{focused,compiler}.log`; unchanged production
and catalog sources were covered by `_artifact/stdlib-url-acceptance-ci.log`.
This accepts the mapped slices capability, not runtime storage changes or a
blanket Go slice-header ABI. The full-inventory review/release gate remains active.

### Accepted: maps in std::collections

The mapped helpers provide shallow independent map cloning, overwrite-copy,
membership-aware equal/equal-by, snapshot entries/keys/values, predicate deletion,
generic iterator insertion and collection. They use existing HashMap Hash/Eq
semantics, not a new map runtime or Go algorithm backend. Snapshot allocation,
shared nested values, unspecified order and lack of synchronization are explicit;
arbitrary iterators remain the caller's resource responsibility.

Existing tests include deterministic Go clone/copy/equality comparisons, collision
keys, NaN values, missing versus zero values, custom Iterator associated types,
duplicate overwrite and snapshot timing. The final audit adds comparator-call
counts, cross-entry callback deletion, nested Vec aliasing, iterator-side-effect
ordering, exhaustion and insertion from the destination's own snapshot. No
production defect or missing helper in the mapped scope was found.

The source is included in the collections catalog and standard dependency closure.
Public documentation, examples, completion and API navigation cover these helpers;
compiler-owned HashMap storage remains unchanged. Focused external execution and
all 1,006 compiler tests passed in
`_artifact/stdlib-maps-acceptance-{focused,compiler}.log`. Existing full CI in
`_artifact/stdlib-url-acceptance-ci.log` covers the unchanged production/catalog
sources. This accepts maps independently of the unfinished slices row and does
not authorize publication or claim live Go map-range compatibility.

### Accepted: path in std::path::slash

The public pure lexical operations cover clean, join, split, base, dir, extension
and absolute checks without changing host std::path behavior. The independent
slash package has no standard-package or native algorithm dependency. Its Pattern
type provides reusable whole-name scalar matching, classes, negation, ranges and
escapes with explicit compile/work limits, recoverable malformed-pattern offsets,
pattern-sized state and no recursive backtracking or shared mutable match state.

Existing fixed and seeded tests compare lexical operations and ordinary shell
patterns with Go path, including UTF-8, backslashes, roots and repeated separators.
The acceptance extension checks 11,664 finite-grammar combinations against an
independent anchored regex oracle, malformed suffixes, exact budget boundaries
and reuse after failed calls. Direct Go comparisons now check oracle errors.
Explicit cases document Go's byte-wise star retry and incomplete exploration
around slash-matching classes; GoML follows scalar, whole-name matching semantics.
These are not broad exclusions of differential failures.

Source catalog, driver selection, pure dependency closure and external consumers
are covered, including Pattern completion and method navigation. The guide
documents all operations, budgets, concurrency-by-immutable-state and the absence
of filesystem resolution or sandbox guarantees. Existing full CI evidence in
`_artifact/stdlib-url-acceptance-ci.log` covers unchanged production/catalog
sources; this test/documentation audit additionally passed focused execution and
all 1,006 compiler tests in `_artifact/stdlib-path-acceptance-{focused,compiler}.log`.
No release or self-hosted consumer migration is authorized by this acceptance.

### Accepted: net/url in std::net::url

The pure value layer provides bounded byte-preserving component escaping,
query decoding/encoding, immutable raw Reference and Authority parsing, generic
RFC reference resolution, explicit ASCII serialization and password redaction.
Absent versus explicitly empty components, raw spelling, decoded buffer ownership,
syntax offsets and resource budgets are documented and tested. Port syntax is
separate from checked transport-port conversion; IPv6 zones and IPvFuture are
textual values, not network operations. Invalid input returns typed errors.

Known-answer and boundary tests are supplemented by deterministic Go comparisons
for byte codecs, query ordering, hierarchical resolution, authority fields and
IPv6 mutations. Deliberate deviations have explicit expectations rather than
broad differential exclusions. Additional tests cover output idempotence, raw
round trips, invalid UTF-8, mutable decode isolation and exact error offsets.

The source catalog includes all four public-source files; driver selection and
link closure retain only standard byte foundations, without socket/DNS/TLS or
ecosystem dependencies. Resource purity, package completion and public method
navigation are covered. Request consumes query codecs and absolute reference
decomposition, retaining its HTTP-specific validation and relative-join policies.
No compiler/driver consumer was migrated to an unreleased standard API.

Final focused execution and full CI passed in
`_artifact/stdlib-url-acceptance-{focused,ci}.log`: 1,006 compiler tests, 171 driver
tests, Go helper race tests, bootstrap fixed point and packaging/smoke checks.
Request integration independently passed 51 library tests, two consumer tests,
cached builds, execution and library/consumer/native race checks in
`_artifact/stdlib-url-request-reference.log`.

This accepts the mapped URI-value capability, not blanket Go URL API compatibility,
IDNA/DNS, universal URL-equivalence normalization, HTTP endpoint authorization or
safe-logging sanitization. Those distinctions are part of the public contract,
not unfinished promises hidden by this acceptance. HTTP/MIME and other ecosystem
inventory rows remain separately tracked. Review and publication gates are unchanged.

### Accepted: sort in std::collections

The existing pure GoML stable merge sort is retained through all six public
default/comparator/Ordering entry points, including those named without stable.
Generic ordering uses Ord; user comparators can define record or descending
orders. All variants preserve equal-key order and mutate shared vector storage.
They use O(n log n) comparisons and O(n) temporary storage. Compiler-owned source
fallbacks remain in place until release/stage0 advancement; no Go backend is used.

The public API includes sortedness predicates, first-match binary search,
lower/upper insertion bounds and equal ranges in default, integer-comparator and
Ordering forms. Abstract indexed search covers both monotone predicates and
comparison-based insertion/found results without materializing a sequence.
Negative lengths fail through Option, empty ranges do not invoke callbacks and
midpoint arithmetic works at maximum isize lengths. The element-to-target sign
convention intentionally differs from Go sort.Find and is explicitly documented.

Tests compare integer sorting/search with Go, and compare all six sorting entry
points against Go stable record ordering on sorted, reversed, all-equal, alternating
signed extremes, duplicate-rich and organ-pipe inputs around merge boundaries.
They verify identity-based stability, comparison counts and alias/view visibility.
Indexed search tests exercise absent/duplicate targets, bounded callback counts,
invalid/empty lengths and maximum-size virtual domains against sort.Find.
Documentation states strict-weak-order, immutability and sorted-input preconditions;
it does not introduce an implicit float/NaN total order, transactional comparator
panic recovery, reflection or automatic callback synchronization.

The existing collections source catalog/dependencies remain unchanged; public
completion and navigation cover the new functions. Build and focused execution
passed, followed by snapshot generation (29 formatter, 17 pipeline and 51
integration tests) and all 1,006 compiler tests. Evidence:
`_artifact/stdlib-sort-audit-build.log`,
`_artifact/stdlib-sort-audit-focused.log`,
`_artifact/stdlib-sort-audit-golden.log` and
`_artifact/stdlib-sort-audit-compiler.log`.
No toolchain-construction or packaging change was made; full CI remains required
before release. This accepts sort, not the separate maps or slices rows.

### Accepted: math/bits in std::math::bits

The public source implements population count, bit length, leading/trailing zeros,
rotation and bit reversal at 8/16/32/64 bits, byte reversal at 16/32/64 bits, and
32/64-bit carry/borrow, double-width multiplication, division and remainder.
Division uses recoverable zero/overflow errors; remainder accepts arbitrary high
words. Zero-divisor errors take precedence over quotient overflow. Boolean carry
inputs rule out invalid multi-bit carry values. Rotation reduces signed counts
modulo the width, including the minimum isize value without negating it.

Auditing both source files found no missing fixed-width family or implementation
defect. Tests cover all 8-bit and 16-bit values for applicable scalar operations,
extreme signed rotation counts, every power-of-two boundary, even/odd divisors,
all high-word overflow thresholds selected by those boundaries, and seeded full
double-width dividends. Go differential checks are combined with independent
quotient/product/remainder reconstruction. Existing FNV-128 is a real standard
consumer of mul64, with its own Go differential tests.

The package has no direct or transitive standard dependency, native algorithm,
runtime backend or compiler intrinsic. Catalog, dependency selection, pure-source
and navigation checks cover it. Explicit widths rather than native-word aliases,
Result rather than division panics, and the absence of constant-time guarantees
are deliberate documented GoML contracts, not unimplemented promises.

Focused consumer execution, snapshot generation (29 formatter, 17 pipeline and
51 integration tests) and all 1,006 compiler tests passed. Evidence:
`_artifact/stdlib-bits-audit-focused.log`,
`_artifact/stdlib-bits-audit-golden.log` and
`_artifact/stdlib-bits-audit-compiler.log`.
The full production tree previously passed CI in
`_artifact/stdlib-endian-blocks-ci.log`; this audit changed only tests and this
ledger, with no production or packaging changes requiring a new bootstrap run.
This accepts math/bits only, not the remaining math, randomness or crypto rows.

### Accepted: encoding/binary in std::bytes::endian and its stream child

The memory package provides checked explicit-offset, cursor and builder operations
for fixed-width unsigned/signed integers, IEEE floats and one-byte booleans, plus
unsigned base-128 and signed ZigZag varints. Endianness is explicit for multibyte
scalars. Failed memory bounds checks preserve output and position; varint cursor
decode errors preserve position. Encoded-length and append helpers are public.

Explicit read_with/write_with callbacks compose fixed-layout structs, arrays and
nested records without reflection or compiler-derived field layout. Readers commit
the declared extent only on success; writers and bounded builders stage a zeroed
block and commit only on success. Padding, callback alias restrictions, scratch
ownership and nested diagnostic coordinates are documented. No automatic serde
mapping, platform-dependent integer layout or generic Go API clone is promised.

The optional stream child covers every scalar, both varint forms and bounded
aggregate callbacks. It distinguishes clean EOF from truncation, retries read
interruption, validates provider counts and reports confirmed read progress.
Writers preserve write_all overrides, perform encoding before writing and never
implicitly flush or close. Stream side effects are not rolled back. Zero-width
records and callback failures have explicit, tested contracts.

External tests cover signed/unsigned boundaries, seeded varint differential cases,
all truncation lengths, tenth-byte overflow and nonminimal encodings, every Boolean
input byte, both byte orders and floating-point payloads. Record tests compare Go
encoding/binary output for mixed fields, fixed arrays and padding, including NaNs;
they also exercise nested failure atomicity, limits, retained scratch isolation,
short/interrupted I/O, invalid counts and no-retry adapters. Production sources
are pure GoML; Go references exist only in test oracles. The deliberate tenth-byte
varint truncation/overflow distinction from Go is retained in the language guide.

Existing catalog, selection and link tests preserve the I/O-free memory package
and separately imported stream child. Pure-source checks cover both packages;
completion/navigation tests cover scalar and aggregate public APIs. Build and
focused execution passed, followed by snapshot generation (29 formatter, 17
pipeline and 51 integration tests) and full CI (1,006 compiler, 171 driver tests,
Go helper race tests, bootstrap fixed point and packaging checks). Evidence:
`_artifact/stdlib-endian-blocks-build.log`,
`_artifact/stdlib-endian-blocks-focused.log`,
`_artifact/stdlib-endian-blocks-golden.log` and
`_artifact/stdlib-endian-blocks-ci.log`.

This accepts the binary encoding row, not the remaining general I/O, buffering or
serialization-format work. User review and publication gates remain unchanged.

### Accepted: hash in std::hash, with std::hash::stream composition

The public Hasher contract provides incremental update, non-consuming independent
sum snapshots, reset, output size and block size. Its bytes-only base package is
independent of I/O and algorithms. Existing checksum implementations and a custom
Hasher defined by the external consumer demonstrate cross-package implementation,
generic dispatch, reset, ownership and independent output snapshots.

The optional stream child adds Read/Write composition without changing the base
contract. It hashes exactly the nonempty successfully reported prefix, validates
counts, preserves provided digest state, forwards flush/errors and supports the
standard interrupted-I/O retry helpers. Malformed counts and unreported partial
progress are explicitly tested; no rollback, implicit close, synchronization or
automatic digest reset/copy is promised. It allocates no transfer buffer.

The source catalog, driver selection, link closure, source-purity checks, namespace
completion and method navigation include the child. The guide documents APIs,
ownership, error behavior and a generic write-and-hash example. External tests
cover real FNV digests, a custom Hasher, short/interrupted transfers, prehashed
prefixes, buffer boundaries, aliases, zero/empty operations and recovery.

Build, focused consumer, snapshot generation (29 formatter, 17 pipeline,
51 integration tests) and full CI passed, including 1006 compiler tests,
171 driver tests, helper race checks, bootstrap fixed-point verification and
packaging/smoke checks. Logs: `_artifact/stdlib-hash-stream-{build,focused,golden,ci}.log`.
This accepts the hash composition contract, not the remaining io/bufio, algorithm
or cryptography rows. Final user review is still required before publication.

### Accepted: hash/adler32, hash/crc32 and hash/crc64

The three standard owners implement the inventory's one-shot, continuation and
stateful checksum contracts entirely in GoML. Adler32 reduces modulo 65521 in
bounded 5552-byte batches. CRC32/CRC64 use private, immutable-by-public-API tables
for arbitrary reflected polynomials, with IEEE/Castagnoli/Koopman and ISO/ECMA
constants respectively. Digest storage is fixed-size; output snapshots own their
bytes, assignment aliases state, and copy creates independent state while safely
sharing CRC tables. No Go FFI, runtime backend or ecosystem dependency was added.

`project093_std_algorithms/checksums.gom` covers fixed answers, deterministic Go
differential inputs from empty through 65,536 bytes, seven-byte chunking, reset,
shared aliases, independent copies and snapshot mutation isolation. The acceptance
extension checks Adler32 continuation across the 5552-byte boundary with maximum
byte input; CRC continuation is compared with Go Update for zero, one, all-ones
and mixed-bit starting states, including zero/one custom polynomials. Splitting
each continuation into two updates must preserve its result. CRC copied states
are also continued independently after the original is reset for every tested
polynomial and input length.

Source/link dependencies, source purity, package completion and actual API
navigation are covered by resource/query tests. The incremental-checksum guide
documents initial state, reflected/complemented CRC convention, byte order,
ownership, non-concurrent access and non-cryptographic limits. Go-compatible state
serialization remains outside these algorithm-row contracts. Shared I/O adapters
are implemented separately in std::hash::stream; no hardware acceleration or
authentication is claimed.

The final fixtures passed snapshot generation (29 formatter, 17 pipeline,
51 integration tests) and all 1006 compiler tests. Logs are
`_artifact/stdlib-checksum-audit-{focused,golden,compiler}.log`; focused execution
preceded the final added CRC-copy checks, which passed in golden/full regression.
Only tests/documentation changed in this audit; subsequent full production CI is
`_artifact/stdlib-hash-stream-ci.log`. These are three accepted capability rows, not
acceptance of the separate hash, io or bufio rows.

### Accepted: hash/fnv in std::hash::fnv

All six FNV-1/FNV-1a variants at 32/64/128 bits implement the public Hasher
contract. `lib/std/hash/fnv/fnv.gom` uses pure integer arithmetic, including
double-width multiplication through std::math::bits; it introduces no Go FFI,
runtime intrinsic or ecosystem dependency. Private width/variant fields are
initialized only through the six supported constructors.

The independent `project093_std_algorithms` consumer compares all six variants
with Go hash/fnv on deterministic inputs from empty through 65,536 bytes, using
both whole-input and seven-byte updates. Fixed empty/abc answers, every split of
short inputs, non-consuming snapshots, snapshot mutation isolation, shared aliases,
independent copies and reset behavior are covered for every variant. The public
source catalog and link closure are tested; editor navigation now checks new128a.

The API/state and non-security contracts are documented in the incremental
checksums section of `docs/goml.md`, with Go provenance retained in
`lib/std/THIRD_PARTY_NOTICES`. Go-compatible binary state serialization is not part
of this row's six-algorithm contract. Shared I/O adapters are implemented separately
in std::hash::stream. No thread-safety or cryptographic claim is made.

Focused execution, snapshot generation (29 formatter, 17 pipeline, 51 integration
tests) and all 1006 compiler tests passed. Evidence is
`_artifact/stdlib-fnv-audit-{focused,golden,compiler}.log`. This audit changed only
tests and documentation; subsequent full production CI is
`_artifact/stdlib-hash-stream-ci.log`, which also included these FNV sources.

### Accepted: net/netip in std::net

The inventory requires extending existing address values with prefixes, zones and
classification without introducing a second address-family model. IpAddr remains
the single V4/V6 representation; ScopedIpAddr composes it with a zone, and IpPrefix
composes it with a checked prefix length. No runtime, DNS, socket or native-backend
capability was added for this row.

| Acceptance requirement | Implementation and evidence |
| --- | --- |
| Prefix construction, masking, membership and overlap | `lib/std/net/address.gom`; `ip_prefix.gom` checks every prefix length, mixed families, unequal prefix widths, invalid input and Go netip differential results. |
| Classification and mapped addresses | `lib/std/net/classify.gom`; `ip_classify.gom` checks all 65,536 IPv6 first words, IPv4/mapped boundaries, literal unspecified addresses and classification policies. |
| Named zone values and explicit numeric socket conversion | `lib/std/net/scoped.gom`; `ip_scoped.gom` checks opaque zones, case, Unicode, percent separators, unmapping, invalid combinations and u32 scope boundaries. |
| Ordering, neighbors and byte representations | `lib/std/net/value.gom`; `ip_value.gom` compares Go netip ordering/bytes/neighbors, carry/borrow boundaries, standard sorting, invalid lengths and copy isolation. |
| HashMap key semantics | Existing Hash derivation on all four address types; `ip_hash.gom` compares map behavior with a linear equality reference, including duplicate spellings, host bits, zones and socket fields. |
| Parser acceptance and numeric correctness | `ip_parse.gom` and native test oracles cover every zero-run compression position, dotted tails, character insertion/deletion, Unicode zone text and invalid prefix lengths; accepted values are compared by raw bytes and exact zone, not just roundtrip strings. |
| Public source, dependency and editor integration | `gomlc/resources/source.gom` registers all value sources and both cmp dependency lists; resource tests, type completion and actual method-navigation tests cover the public APIs. |
| Documentation, provenance and bootstrap | Networking contracts in `docs/goml.md`, Go provenance in `lib/std/THIRD_PARTY_NOTICES`, and full CI including fixed-point, packaging, helper race checks, 1006 compiler tests and 171 driver tests passed after the last production change. |

The fixtures above are in `gomlc/testdata/module/project093_std_algorithms/`.
Latest parser-audit evidence is `_artifact/stdlib-ip-audit-focused.log`,
`_artifact/stdlib-ip-audit-golden.log` (29 formatter, 17 pipeline, 51 integration
tests) and `_artifact/stdlib-ip-audit-compiler.log` (1006 passes). The last full
production-change CI is `_artifact/stdlib-hash-stream-ci.log`.

Intentional contract differences are explicit: Option replaces invalid-address
sentinels and panic-only as4 conversion; nonempty IPv4 zones are rejected rather
than discarded; the existing hexadecimal mapped-IPv6 presentation is retained;
numeric socket scopes do not resolve interface names; scoped values do not enter
prefix matching without explicit zone removal. Hash values are not a stable
serialization format, and global-unicast classification is not a reachability or
security decision. These preserve the planned GoML ownership and value contracts.

## Implementation record

The newest-first records below preserve development history. The acceptance
status above supersedes older outstanding-work notes for accepted rows.

### Accepted: net/http/httptest in ecosystem::web::testing

The child package provides a reusable in-process Recorder with completed-call
history, a loopback TestServer wrapper with validated URLs and bounded shutdown,
and a scripted FaultTransport for reject, body-truncation and post-dispatch
failures. These reuse the web router's dispatch and server machinery. Focused
tests cover snapshots, concurrent recording, fault order, and server lifecycle;
the independent web consumer uses the public package with an HTTP client.
Isolated verification passed format, build, 36 web tests, consumer tests,
execution and library/consumer race checks. The httptest row is locally
accepted; TLS test servers, Go ResponseRecorder identity and packet-level
fault injection are not claimed.

### Accepted: net/http/cookiejar in ecosystem::request::cookies

The child package has a bounded Public Suffix List parser with exact, wildcard
and exception matching, plus a pinned bundled ICANN-and-private rule snapshot
and literal Punycode aliases for Unicode rules. Its bounded jar parses
Set-Cookie, applies host/domain/path/Secure/expiry policy, evicts old entries
at per-domain and total quotas, and is used by the request Client behind its
existing cookie-store option. Independent request consumer coverage exercises
the public jar. Isolated verification passed format, build, 60 request tests,
consumer tests, execution and native/GoML race checks. The cookiejar row is
locally accepted; release and the other E1 rows remain open.

### Accepted: mime/multipart in ecosystem::mime::multipart

The child package provides bounded streaming Reader/Writer APIs with validated
boundaries, textproto header parsing and serialization, folded headers,
incremental binary bodies, cross-chunk delimiter recognition and collision
checks, and per-part, aggregate, header and wire quotas. Focused tests cover
fragmented sources, false delimiters, preambles, closing boundaries at EOF,
malformed framing, quotas, short writes and provider failures. The independent
MIME consumer exercises both directions, and request multipart construction
delegates to the shared writer through its compatible wrapper. Isolated
verification passed format, build, tests, execution and race checks for MIME
and request; the final MIME verification also passed after the last test was
added.

### Accepted: mime/quotedprintable in ecosystem::mime::quotedprintable

The child package has bounded streaming Reader/Writer APIs, strict CRLF and
soft-break decoding, 76-character wire folding, checked trailing whitespace,
text and binary writer modes, and sticky provider failures. Focused chunking,
known-byte, malformed-input, quota, short-write and RFC 2045 soft-break tests
pass. The independent MIME consumer exercises the child package; isolated
verification passed format, build, tests, execution and race checks for the
library and consumer. The separate multipart row is recorded above.

### Accepted: mime in ecosystem::mime

The independent module has bounded media type and parameter parsing and
formatting, including quoted values, UTF-8 RFC 2231 single-part extended
parameters and out-of-order continuations with checked gaps, duplicates and
cross-segment UTF-8. RFC 2047 B/Q encoded words decode UTF-8, US-ASCII and
ISO-8859-1 and encode UTF-8 across 75-byte word boundaries. Malformed-input,
quota, exact-output and round-trip tests pass. The independent consumer uses
the public API, and request multipart metadata now delegates media type
validation to it. Isolated verification passed format, build, tests,
execution and race checks for mime and request. The separate multipart row
is recorded above.

### Accepted: net/textproto in ecosystem::textproto

The independent module provides bounded strict-CRLF physical lines,
byte-preserving header fields with canonical names and optional continuation,
and incremental dot-unstuffing/stuffing. Separate tests cover exact quota
boundaries, malformed framing, chunk splits, source/sink failures and terminal
error behavior. Its independent consumer exercises the public framing API;
request and web use the shared single-header parser while retaining their
existing transport buffering and policy. The isolated verification registry
report passed format, build, test, execution and race checks for textproto,
request and web. Mail policy and the multipart and quoted-printable rows remain
future work.

### Accepted: archive/zip in ecosystem::archive

ZIP64 end-of-directory records, central/local size and offset extra fields,
and 64-bit data descriptors now have checked bounded parsing. Info-ZIP's
streamed-stdin encoding is exercised while preserving the existing rejection
of FIFO entries. `ZipArchive::copy_entry` incrementally decompresses and checks
CRC32 without retaining the expanded body; existing entry APIs remain intact.
ZIP64 writing is automatic at classic-field limits or can be forced for testing;
the forced output passes external `unzip -t`, and the 65,535-entry boundary is
covered. `ZipSource[ReadAt]` indexes metadata without reading all entry bodies,
and `open_zip_source` adapts a regular file with explicit closing. Both memory-
and source-backed `copy_entry` stream decompressed bytes with size and CRC32
checks. Valid and malformed ZIP64 sources, bounded random-access reads and an
independent archive consumer are covered. `verification compress archive`
passed library and consumer formatting, tests, builds, execution and race
checks. `ZipWriter::append_reader` intentionally retains one bounded entry;
streaming writer input is not part of this row's decompression contract. ZIP64
support stays within configured archive, metadata and expanded-data limits.
This locally accepts the ZIP row and completes the archive package; release
and the other E1 rows remain separate.

### Accepted: archive/tar in ecosystem::archive

`TarReader::next_header` and `read_body_chunk` expose bounded incremental entry
bodies; moving to the next header drains an unread body and padding. The existing
`next` and whole-archive APIs remain compatible. The reader accepts USTAR/PAX
and selected GNU extensions: long name/link records and checked base-256 mode,
owner, size and signed timestamp fields. GNU sparse entries, special device/FIFO
types and extraction without an explicit policy remain unsupported. The writer
continues to use PAX for long paths and numeric fields.

GNU `tar --format=gnu` interoperability verifies long paths and links, large
owner/group IDs and negative timestamps. Malformed long records, path limits,
truncation, partial body reads and provider faults are covered. The independent
archive consumer uses streaming entry reads. `verification compress archive`
passed library and consumer format, tests, cached builds, execution and race
checks for both modules. This accepts only the TAR capability locally; archive
remains uncommitted until its ZIP capability is complete.

### Accepted: compress/flate in ecosystem::compress::flate

The raw DEFLATE engine now has bounded incremental decoding and encoding,
preset dictionaries, stored/fixed/dynamic decoding, stored/fixed encoding,
compression levels -2 through 9, synchronization flush and explicit terminal
state. The archive GZIP/ZIP entry points use the shared engine for both
directions; the old archive-local encoder has been removed. The independent
compress consumer exercises the public streaming encoder and decoder.

Go flate interop tests cover all supported encoder levels, dictionaries larger
than 32 KiB, short input/output chunks, flush, incompressible fallback,
malformed/truncated decoder input, quotas, provider failures, reentrancy and
panic cleanup. The `verification compress archive` run passed library and
consumer format, tests, cached build, execution and race checks for both
modules. Dynamic-Huffman encoding is not claimed; RFC
1951 permits fixed and stored blocks. This is local acceptance of the mapped
flate row, not completion of the other compression rows or the E1 objective.

### Accepted: compress/gzip in ecosystem::compress::gzip

The archive GZIP framing has been extracted behind its compatible public
wrappers. Shared `gzip::Writer` and `gzip::Reader` compose with the incremental
flate engine, validate RFC 1952 headers and CRC32/ISIZE trailers, and support
concatenated members without retaining complete streams. Writer limits include
the header and trailer, and reader limits include framing bytes and work. Both
implement standard I/O traits and expose confirmed progress and sticky failures.

Go gzip encoder/decoder interop, representative levels, optional extra/name/
comment/header-CRC fields, concatenated members, corruption, truncation, bounds,
provider faults, reentrancy and panic cleanup pass. The independent compress
consumer uses the public GZIP API. `verification compress archive` passed
library/consumer format, tests, cached build, execution and race checks for both
modules after extraction. This accepts GZIP locally, not the remaining
compression rows or E1 overall.

### Accepted: compress/zlib in ecosystem::compress::zlib

The shared ZLIB reader and writer compose with incremental flate and standard
Adler32. They validate RFC 1950 CMF/FLG/FCHECK, declared window size,
full-dictionary DICTID, and the uncompressed-data trailer. Preset dictionaries
are supported for both directions, including input longer than 32 KiB; only
the final window is retained for DEFLATE history. The streaming reader preserves
following input, while bounded whole-buffer convenience rejects trailing data.

Go zlib bidirectional interop covers representative levels, dictionary/no-
dictionary frames, short output chunks and large payloads. Negative tests cover
missing/wrong dictionaries, malformed headers, advertised-window violations,
Adler32 mismatch, truncation, bounds, invalid/failing providers, reentrancy and
panic cleanup. The independent compress consumer uses the dictionary API, and
`verification compress archive` passed format, library and consumer tests,
cached build, execution and race checks for both modules. This accepts ZLIB
locally, not the other compression rows or E1 overall.

### Accepted: compress/lzw in ecosystem::compress::lzw

Bounded incremental LZW reading and writing now support LSB and MSB bit
packing, literal widths 2 through 8, 12-bit maximum codes, clear/EOF codes,
dictionary growth and reset, and the special just-created-code expansion.
The public contract documents GIF/PDF-style width timing and excludes TIFF's
incompatible variant. Input/output, clear-epoch and work quotas, malformed
codes, short writes, provider faults, reentrancy and panic cleanup are terminal
and recoverable. Streaming readers preserve following framing; whole-buffer
conveniences reject trailing input.

Go `compress/lzw` bidirectional interoperability covers both bit orders,
literal widths 2, 4 and 8, chunk sizes, 30 KiB streams and dictionary resets.
Malformed/truncated input, partial output, quota and provider tests pass. The
independent compress consumer uses the public LZW codec. The compress/archive
verification passed module and consumer format, tests, cached build,
execution and race checks for both modules. All four mapped E1 compression
rows are locally accepted; the remaining E1 rows and publication are not.

### E1 shared compression and archive migration — locally accepted

All mapped E1 capability rows have passed their local acceptance gates. This
does not complete the 74-row migration or authorize publication.

- compress/flate is owned by ecosystem::compress::flate, with std::io/bytes
  dependencies. Archive is the first real consumer; an independent compress
  consumer and isolated verification-registry entry are added.
- Both archive encoding and decoding now use the shared flate and GZIP engines
  with compatible entry points.
- HTTP trace, reverse proxy and protocol helpers are locally accepted alongside
  compress, archive, textproto and MIME. Release and stage0 advancement remain
  separate gates.

### S2 bounded JSON streaming — accepted

- The encoding/json row remains in std::json and retains existing serde/value
  APIs. The new core uses incremental Read/Write, token grammar and
  direct typed serde events, not whole-input buffering or Value intermediates.
  Explicit budgets cover transport, roots, scalar payloads, containers, paths,
  protocol frames and events. Aliases share sticky failure and reentrancy state.
- Initial audit found invalid non-ASCII literal prefixes could panic in both
  parser paths, and as_int could wrap out-of-range values. Bytewise literal
  matching and checked magnitude accumulation now fix these cases; external
  boundary regressions are added. The toolchain build, complete project093
  runtime fixture and eight existing JSON-related compiler/LSP tests pass
  (_artifact/stdlib-json-boundaries-*.log).
- Incremental token and iterative Value decoding/encoding, direct typed serde
  adapters, checked reader/writer and string conveniences are now implemented.
  Explicit protocol frames validate optional, tuple, sequence, map, struct and
  enum events; unknown fields are skipped without building a Value tree.
  Strict scalar parsing, diagnostic paths, sticky failures, reentrant callbacks,
  expired adapters, panic cleanup and confirmed I/O progress share one session.
- External fixtures cover every chunk size for a Unicode typed record,
  typed/value/token root switching, width/range failures, duplicate and missing
  fields, maps/enums, malformed protocol events, all quota families, cyclic
  Value input, reader/writer faults and independent Go token/acceptance
  differential checks including deterministic mutations. The focused JSON suite
  passes in `_artifact/json-only.log`; the earlier full algorithms fixture
  passes in `_artifact/json-focused.log`.
- All 18 JSON source files are registered; resource count/closure and public
  completion/navigation tests cover the new API. `docs/goml.md` documents
  contracts and intentional Go differences. Self-hosted consumers and stage0
  are unchanged. Whole-row validation is complete as recorded in the acceptance
  section, including generated snapshot refresh and non-update verification.
  Counts are now 28 accepted, 0 started and 46 planned.

### S2 typed formatting — implementation under validation

- The fmt row is owned by std::text::format. Explicit typed field specifications
  compose text/quoting, integer and f32/f64 formatting, scalar-count padding and
  existing ToString/Debug conversions. Existing num and text engines are reused;
  interpolation syntax, runtime reflection and compiler/driver consumers do not
  change. Format-string interpretation and reflection-driven scanning are not
  part of this mapped contract.
- A private shared Formatter enforces total UTF-8 output bytes. Each library
  field validates before committing; conversion callbacks retain their own
  side effects and are not implicitly recovered. write_to uses primitive Write
  with confirmed byte progress, bounded count validation and no error retries.
- Three parallel implementations and four external fixture families are
  integrated, with source registration, package selection and navigation
  coverage. The toolchain build, complete project093 runtime fixture and focused
  dependency-selection regression pass (_artifact/stdlib-format-*.log), along
  with all 11 resource and 95 query-namespace tests. Independent whole-row review
  found no production gap against the typed-formatting contract. Snapshot
  regeneration passes (29 formatter, 17 pipeline, 51 integration tests);
  final CI is running. Counts are 26 accepted,
  1 started (fmt) and 47 planned.

### S2 filesystem helper cross-checks — implementation under validation

- check_helpers compares complete basic-checker node snapshots with direct
  stat/read-file/read-directory helpers and SubFs compositions. Directory results
  must be sorted and preserve duplicate multiplicity; child links are not followed.
- Private observed providers meter underlying calls and confirmed bytes against
  the shared limits. Per-transaction event identity separates provider failures
  from synthetic stop signals and avoids duplicating close-only failures.
  Successful opens immediately register exactly-once cleanup guards, including
  independent catch boundaries for cleanup after errors, budget stops and panics.
- Parallel implementations and external regressions are integrated. Toolchain
  construction, the complete project093 fixture and all 10 resource tests pass
  (_artifact/stdlib-fs-check-helpers-*.log). All 95 query-namespace tests pass.
  The external suite also passes compound read-error/close-panic,
  entry-panic/close-panic, partial-directory-error/entry-panic and
  budget-stop/close-panic regressions, preserving error ordering and the first
  stop reason. Snapshot regeneration and final CI remain pending;
  testing/fstest is still started.
- Independent whole-row review found no remaining production scope gap: memory
  construction/resolution/handles and helper/SubFs/walk/glob composition are
  covered by the four fs_memory fixture families; reports, primitive contracts,
  static capabilities and helper observation are covered by eight fs_check
  families. Read-only snapshots and static capability entry points are deliberate
  GoML adaptations, not promises of Go MapFS mutability or optional-interface ABI.
  Acceptance remains conditional on the final validation gates above.

### S2 static filesystem capability checks — implementation under validation

- check_seek/check_read_at/check_links add statically constrained capability
  checks over the basic tree checker, sharing budgets and cleanup. Seek respects
  provider-specific End support and error side effects. ReadAt validates confirmed
  progress, invalid ranges and cursor preservation with separate fresh probes for
  successful, partial-EOF and empty requests. Links compare raw targets and
  non-following metadata without imposing a following-resolution policy.
- Three additional external fixture families cover valid providers, malformed
  counts and positions, ordinary provider errors, snapshot inconsistencies and
  exact shared-budget boundaries. Further review added nonempty-file empty-read
  cursor probes and retained ordinary errors returned for invalid ReadAt ranges.
- Integration exposed two compiler defects. TAST signature projection candidates
  now include inherited trait applications, deduplicated by receiver, canonical
  trait identity and complete generic arguments. Four new tests preserve true
  ambiguity while covering parent/diamond/generic where constraints; all 197
  TAST tests pass.
- Artifact function exports now use checked function signatures rather than raw
  declarations, preserving projection owners and implied constraints through
  interface/core serialization. A focused cross-package roundtrip and Mono
  regression covers that boundary. Callable/comptime metadata, visibility,
  extern fallback and existing default-method requirements remain unchanged.
  The implementation does not add redundant public bounds or weaken Mono checks.
- The combined build and complete project093 runtime fixture pass, including
  all six checker fixture families. The actual interface/core binary-roundtrip
  and Mono regression also passes. Evidence is under
  _artifact/stdlib-fs-check-capabilities-*.log. Broader resource, navigation,
  snapshot and CI validation is ongoing. All 97 navigation tests and ecosystem
  csv/msgpack/rope/archive checks also pass. Helper/SubFs integration and final
  whole-row verification remain outstanding; testing/fstest is not yet accepted.

### S2 recoverable filesystem contract checker — implementation under validation

- check_fs returns a bounded report instead of invoking testing assertions.
  ExpectedTree distinguishes Contains from Empty, validates and snapshots input,
  and does not infer missing paths after incomplete or failed traversal.
- Proven primitive-contract violations, stable-snapshot inconsistencies, ordinary
  provider errors and catchable provider panics retain distinct structured issue
  kinds. Budgets and report capacity produce explicit incomplete stop reasons.
- Every successful open has an independent exactly-once cleanup boundary.
  Cleanup is exempt from call budgets, runs after stops and panics, and preserves
  its own errors without overwriting earlier findings or the first stop reason.
- The initial implementation covers invalid logical paths, interleaved independent
  reads, checked counts and byte budgets, stat/empty-read cursor preservation,
  directory paging and nonpositive requests, entry lifetime after close, and
  duplicate-preserving multiset comparisons. It permits unsorted raw pages and
  legal NUL/backslash/colon names, and does not follow child symlinks.
- Build and 97 query tests pass. Three external fixture families exercise valid
  providers, budget edges, cursor/count faults, normal errors and panicking reads,
  entry accessors and close operations. The complete external fixture and all 10
  resource tests pass for this initial primitive-checker surface.
- Standard-helper/SubFs cross-checks and static link/Seek/ReadAt extensions are
  still outstanding, along with their tests and final whole-row verification.
  Counts remain 25 accepted, 1 started and 48 planned; no publication is implied.

The remaining helper/SubFs cross-check design uses a private observed-provider
wrapper, not an unmetered call to a convenience helper. Every underlying call
and confirmed read prefix must share the same budgets; helper entry accessors
also count. A per-helper observation transaction distinguishes real provider
events from synthetic budget/panic exits without comparing error text or values.
Successful opens register cleanup guards shared by normal close and outer
fallback cleanup. Cleanup errors retain separate identity so a close-only error
is not duplicated and a close panic cannot erase a prior read/stat error.
Non-Result entry accessors use safe exit values only after marking the transaction
aborted; aborted helper results must never be accepted as confirmed data.
Only complete node baselines are compared with helper outputs. Directory
comparisons require sorted helper names but retain duplicate multiplicity;
SubFs is constructed over the observed provider so prefixed calls remain metered.

### S2 public in-memory filesystem — validated, checker still outstanding

- std::testing::fs now contains MemoryFS, MemoryNode, MemoryContent, MemoryLimits
  and MemoryHandle. Construction snapshots bytes, preserves explicit metadata,
  validates kinds and bounded logical paths, and infers sorted parent directories.
- Links resolve with an explicit component stack, bounded hops and work. Raw
  final targets remain inspectable; dot-dot follows the directory actually
  reached rather than lexically cleaning the target. Independent opens have
  independent cursors, with shared state only between copies of one handle.
- Read, positional read, checked seek, snapshot directory pages and idempotent
  close use existing GoML primitives without new runtime or native bindings.
  Source catalog, navigation, documentation and four external fixture families
  are integrated, including walk/glob/nested-SubFs composition. Build and the
  complete project093 fixture pass, along with 97 query tests, 10 resource
  tests and the focused package-selection regression. Registration includes
  the compiler's explicit package-selection mappings, not only its catalog.
  Snapshot verification passes 29 formatter, 17 pipeline and 51 integration
  tests. Full CI passes 1016 compiler and 171 driver tests, bootstrap fixed-point,
  packaging and smoke checks. Evidence is under _artifact/stdlib-memory-*.log.
  This validates MemoryFS, not the whole testing/fstest row.
- The recoverable checker remains outstanding. Its design separates invalid
  checker input from a bounded report: proven contract violations, ordinary
  provider errors, stable-snapshot inconsistencies and catchable provider panics
  are distinct outcomes. Incomplete budget-limited checks must never pass.
  Expected paths will distinguish Contains from Empty; missing paths can only
  be inferred from successfully completed traversal, not failed enumeration.
- The checker must inspect primitives before comparing standard helpers, retain
  duplicate directory entries as a multiset, accept unsorted raw pages and legal
  NUL/backslash/colon names, and avoid comparing following stat to link entry
  metadata. Separate static checks will cover links, seeking and positional
  reads. Cleanup must close every successful open once, even after exhaustion
  or a catchable panic; cleanup failures must not erase the primary evidence.
- Shared bounds will cover traversal, reads, provider calls and report size.
  They do not interrupt blocking providers or limit allocation inside providers.
  Malicious and valid/erroring provider fixtures are required in addition to
  MemoryFS checks. No self-hosted consumer migration or publication is included.

### S2 portable filesystem metadata — validated, io/fs accepted

- FileMode/NodeKind/ModeFlags preserve advanced portable kinds and distinguish
  unknown flags from known all-false flags without extending the old FileType
  enum. Metadata retains the complete mode while keeping prior getters and
  from_parts behavior; legacy host data remains explicitly incomplete.
- SnapshotEntry is a public metadata-to-entry value adapter. Bounded metadata
  and entry display helpers preserve UTF-8 and never perform implicit metadata
  I/O. The output is deterministic display text, not escaped serialization.
- Source registration, navigation, docs and external value/format/provider tests
  are integrated. Build, the complete project093 external fixture, all 97 query
  tests and csv/msgpack/rope/archive library, independent-consumer, cache and
  configured race checks pass, as do all 10 resource-catalog tests. Snapshot
  updates pass 29 formatter, 17 pipeline and 51 integration tests. Final CI
  passes 1015 compiler and 171 driver tests, fixed-point, packaging and smoke
  checks. No runtime boundary or self-hosted consumer migration is added.
  The io/fs row is accepted; testing/fstest is accounted for separately.

The io/fs whole-row audit maps the following production families to external
project093 fixtures. Final CI has passed for this complete production surface.

| Capability family | Production sources under std/fs | External fixtures |
| --- | --- | --- |
| Read-only filesystem/handle contracts and logical paths | portable.gom | fs_portable.gom |
| Bounded reads, stat and close/error/panic precedence | portable.gom | fs_portable.gom |
| Paged directory results, confirmed prefixes and stable sorting | directory.gom, read_directory.gom | fs_directory.gom, fs_directory_errors.gom |
| Lazy entry metadata and existing host adapter | directory.gom, fs.gom | fs_directory_host.gom |
| Namespace composition | sub.gom | fs_sub.gom plus directory/walk/glob/link composition |
| Iterative traversal, callbacks, skip controls and global budgets | walk.gom | fs_walk.gom, fs_walk_cycle.gom, fs_walk_cleanup.gom |
| Component patterns, literal optimization and sorted partial matches | glob.gom | fs_glob.gom, fs_glob_boundaries.gom |
| Raw link targets and final-link metadata | links.gom | fs_links.gom |
| Advanced portable kinds, permissions and known/unknown flags | mode.gom, fs.gom | fs_mode.gom, fs_directory_host.gom |
| Metadata snapshot-to-entry conversion | mode.gom | fs_mode.gom |
| Bounded deterministic metadata/entry display | format.gom | fs_metadata_format.gom |
| Structured causes and partial-result snapshots | error.gom, directory.gom, glob.gom | directory/glob/cleanup fault fixtures |

This is a portable io/fs contract, not an upgrade of host metadata or Go ABI
compatibility. Advanced values are expressible by portable providers while the
existing host boundary retains its old information limits. Static associated
handle and link bounds replace dynamic optional-interface discovery; pure data
providers adapt through ordinary handles. No sandbox, opaque host Sys value,
new native backend or implicit optimized-dispatch protocol is claimed.
Public MemoryFS and the recoverable provider checker belong to the separate
testing/fstest row and are not counted as implemented by private fixtures.

### S2 portable glob and link capabilities — implementation under validation

- glob_from combines source-only slash Pattern matching with the existing
  named-directory kernel. It uses an explicit stack, validates all components
  before I/O, optimizes literal paths, retains sorted confirmed matches on error
  and bounds pattern size, components, directories, match calls/results and work.
  Conservative matching costs are prepaid with overflow-safe arithmetic.
- LinkFileSystem, read_link_from and symlink_metadata_from require static link
  support and preserve provider results. Nested SubFs prefixes input names but
  never rebases targets or silently falls back to following stat.
- Public docs, source/dependency registration, navigation and external fixtures
  are integrated. Build, the complete external fixture, all 97 query tests and
  csv/msgpack/rope/archive regressions and all 10 resource-catalog tests pass;
  final CI validation remains pending. The whole-row audit identified portable mode values, metadata
  snapshot entries and descriptor formatting as remaining work. These are being
  implemented without extending the runtime; testing/fstest remains a separate
  planned row. Counts stay 24 accepted,
  1 started and 49 planned.

### S2 portable filesystem traversal — implemented and validated

- walk_from adds explicit-stack sorted DFS with typed Continue/SkipDir/SkipAll
  controls and recoverable provider error callbacks. Root metadata and cached
  child descriptors remain usable after closing. Child symlinks are not followed;
  root behavior is provider-defined and no confinement is claimed.
- A shared private directory kernel retains cached names and confirmed prefixes.
  read_dir_from keeps its public behavior and uses no shared work budget.
  WalkLimits bounds depth, per-directory entries, expansion attempts and work;
  Continue/SkipDir cannot bypass a hard limit. Explicit SkipAll may stop cleanly.
- Public docs, source registration, navigation and external fixtures are
  integrated. Build, the external runtime fixture, 97 query tests, 10 resource
  tests and csv/msgpack/rope/archive regression checks pass. The additional
  repeated-directory fixture verifies shared global budgets and finite
  termination, including Continue/SkipDir attempts to ignore the limit.
  Snapshot updates pass 29 formatter, 17 pipeline and 51 integration tests.
  Full CI passes 1015 compiler and 171 driver tests, bootstrap fixed-point,
  packaging and smoke checks. Logs use _artifact/stdlib-fs-walk-*.log.
  Globbing, link protocols and testing/fstest remain outstanding; counts stay
  24 accepted, 1 started and 49 planned.

### S2 portable directory enumeration — implemented and validated

- DirectoryEntry/ReadDirFile protocols preserve lazy metadata, positive-sized
  pages, EOF and confirmed partial errors. ReadDirError snapshots its vector,
  not arbitrary entry objects. Existing host DirEntry gains the portable trait
  without changing its eager listing API.
- read_dir_from uses chained associated types and an explicit handle bound.
  It validates entire pages, caches names once, bounds requests and results,
  stably sorts confirmed prefixes and closes once including panic cleanup.
  Invalid pages outrank provider errors; valid error pages preserve their cause
  even when entries exceed the remaining budget. SubFs composes unchanged.
- Added normal and fault-provider tests, source/dependency registration,
  navigation coverage and documentation. The full project093 runtime fixture,
  including host directory entries and partial-error budget boundaries, passes.
  All 97 query tests and 10 resource-catalog tests pass, as do the
  csv/msgpack/rope/archive module, independent-consumer, cache and configured
  race checks. Snapshot updates pass 29 formatter, 17 pipeline and 51 integration
  tests. Full CI passes 1015 compiler and 171 driver tests, bootstrap fixed-point,
  packaging and smoke checks. Logs use _artifact/stdlib-fs-directory-*.log.
  Traversal, globbing, link protocols and testing/fstest remain outstanding. Counts remain
  24 accepted, 1 started and 49 planned.

### S2 portable filesystem foundations — implemented and validated

- std::fs gains File/ FileSystem associated-handle protocols, logical valid_path,
  stat_from and bounded read_file_from. Existing host functions are unchanged.
  A private forwarding reader prevents overriding read_to_end from bypassing
  the helper's limit. Opened handles close once on normal/error/panic paths;
  primary returned errors take precedence over returned cleanup errors.
- SubFs validates and prefixes logical names, with root identity and nested
  composition. It neither probes directory existence nor confines providers or
  symlinks. Provider errors and handles remain unchanged.
- New external fixtures cover paths, malicious helper overrides, confirmed
  short reads, limit probes, partial side effects, error precedence, cleanup and
  namespace composition. Build, full external project093 runtime, 193 TAST
  tests, 97 query tests and 10 resource-catalog tests pass. csv/msgpack/rope/archive
  module, independent-consumer, cache and configured race checks pass.
  Generated snapshots pass 29 formatter, 17 pipeline and 51 integration tests.
  Full CI passes 1015 compiler and 171 driver tests, bootstrap fixed-point,
  packaging and smoke checks. Logs use _artifact/stdlib-fs-foundations-*.log.
- SubFs exposed an associated-type cycle-checker false positive: F::Handle was
  treated as the current wrapper's Handle. Cycle edges now require the same
  receiver and full trait identity, while still rejecting direct, indirect and
  nested Self cycles. Focused compiler tests cover foreign receivers and trait
  identities; this is a correction to existing associated-type semantics.
- Cross-package validation also exposed an export-qualification problem:
  restored trait-implementation constraints contain bare File instead of its
  canonical std::fs::File identity. Associated-bound expansion now resolves
  the trait name in its declaring trait's package, as supertrait expansion
  already does. A non-main-package test checks generic function and impl
  predicates, and the formerly failing external SubFs consumer passes. The
  speculative Mono normalization change was removed; no artifact format or
  trait-bound changes were needed.
- io/fs is started, not accepted: directory capabilities, sorted/paged results,
  bounded walk/glob, link-specific contracts and complete integration remain.
  testing/fstest's public memory provider and contract checks remain separately
  planned. Counts are 24 accepted, 1 started and 49 planned.

### S2 synchronous pipe — implemented and validated

- PipeReader/PipeWriter use a channel-held state token and broadcast notification
  epochs; no new runtime boundary, permanent worker, or unsynchronized shared
  mutable state. Concurrent writers do not interleave their packets.
- Empty direct operations do not rendezvous. Closing wakes readers and writers;
  first causes, local-close priority and completed packet replies have explicit
  contracts. write_progress preserves only the confirmed consumed prefix.
- Added partial-transfer, concurrent producer/consumer, post-return input
  mutation and close-order regression tests. Build, the full project093 runtime
  fixture and its race-instrumented executable pass. All 97 query tests and the
  csv/msgpack/rope/archive module, consumer, cache and configured race checks
  pass. Snapshot updates pass 29 formatter, 17 pipeline and 51 integration tests.
  Full CI passes 1010 compiler and 171 driver tests, fixed-point verification,
  packaging and smoke checks. Logs use _artifact/stdlib-io-pipe-*.log. The two
  mapped io/bufio rows are accepted: 24 accepted, 0 started, 50 planned.

The io/bufio whole-row audit maps capability families to the following
project093_std_algorithms fixtures. This is functional GoML coverage, not a
promise to duplicate Go's interface names or optimized dispatch protocols.

| Capability family | Fixture files | Deliberate contract differences |
| --- | --- | --- |
| Minimum/exact reads and errors | io_read_minimum, io_read_retry, io_write_retry | No generic Interrupted retry; a failing call can have unknown side effects. |
| Bounded copy and confirmed progress | io_copy, io_copy_progress | Separate confirmed read/write counts; no automatic ReaderFrom/WriterTo dispatch. |
| Positional access, sections and offsets | io_at | TransferError carries partial progress; mutable adapters are not implicitly concurrent-safe. |
| Seek and string input | io_seek, io_seek_at, string_reader | Cursor/Section seeks are bounded; StringReader permits past EOF; OffsetWriter rejects End. |
| Multi-stream, tee and close composition | io_multi, io_tee, io_noop_close | Side effects are not rolled back; NoopCloser leaves the source open. |
| Synchronous concurrent pipe | io_pipe, io_pipe_close | Borrowed input until reply, first close causes, local-close priority, no FIFO guarantee. |
| Buffered capacity, reset and aliases | io_buffer_reset | 8 KiB defaults; reset does no I/O; available is logical capacity. |
| Peek, discard and buffered transfer | io_peek, io_discard, io_buffer_composition | Checked views/progress; buffered transfer composes generic copy. |
| Byte/scalar reads and rollback | io_buffer_bytes, io_buffer_chars | Explicit rollback eligibility; physical scalar storage can exceed configured capacity. |
| Delimiter and line fragments | io_fragments | Explicit end states and consumed counts; errors preserve the consumed fragment. |
| Buffered typed output and input transfer | io_buffered_write | Counts new input acceptance, not physical delivery; no implicit flush. |
| Standard/custom scanner splitting | scanner | Owned token snapshots, explicit limits/final states, sticky failures. |

Trait bounds compose Read/Write/Close without nominal combined interfaces.
ReadString/Scanner.Text-style use composes byte reads with explicit UTF-8
decoding. BufReader-to-writer transfer uses copy_progress; controlled buffered
writes avoid exposing Go's append-capacity AvailableBuffer optimization view.
These are design choices rather than deferred runtime work. Existing resource
cleanup tests also remain in project089_resource_bytes_serde. The audit finds
no remaining production capability family in these two mapped rows; the current
snapshot and full CI results above complete their local acceptance.

### S2 bounded fragments and buffered write composition — implemented and validated

- BufReader::read_fragment and read_line_fragment expose borrowed bounded
  payloads, Delimiter/BufferFull/Eof states and raw consumed-byte counts. Errors
  expose already-consumed confirmed fragments with the original cause; unknown
  failed-call side effects are excluded. Full buffers do not probe for EOF.
  Incremental scanning avoids rescanning prefixes on each short read.
- Line fragments preserve split CRLF, bare CR and precise consumption, including
  configured capacity one using bounded lookahead in existing scalar storage.
  Previously buffered bytes beyond configured capacity remain ordered. Fragment
  reads always clear rollback permissions, and copies retain shared state.
- BufWriter now composes write_byte/write_char/write_string/read_from with the
  existing short-write and no-retry contracts. String scratch is at most 4096
  bytes; write errors report acceptance of new input, not prior buffered output
  or physical sink delivery. Neither EOF nor successful typed writes implicitly
  flush or close. Partial UTF-8 writes are not promised to be atomic.
- External fixtures cover byte delimiters, short I/O, capacity and EOF borders,
  CRLF and error-side-effect paths, aliases, snapshots, typed writes, transfer
  counts and chunk boundaries. Resource packaging, navigation, completion and
  language documentation include these APIs. Production remains pure GoML.
- Formatting, stage2 build, the complete project093 fixture, all 97 query tests
  and CSV/MessagePack/Rope/Archive verification pass, including independent
  consumers, cached builds and configured race tests. Regressions also cover
  empty confirmed error prefixes, CR lookahead after scalar residue and scratch
  boundaries inside UTF-8 scalars. Golden regeneration passes 29 formatter,
  17 pipeline and 51 integration tests. Full CI passes 1010 compiler tests,
  171 driver tests, bootstrap fixed point, helper, packaging and release-smoke
  checks.
  Logs use
  _artifact/stdlib-io-fragments-write-. Counts remain 22 accepted, 2 started and
  50 planned. Pipe and final io/bufio contract audits still precede acceptance;
  full A/B implementation and user review remain required before release.

### S2 buffered byte/scalar reads and rollback — implemented and validated

- BufReader now exposes fallible read_byte/read_char and one-shot
  unread_byte/unread_char. Scalar reads preserve original encoded bytes, replace
  malformed or EOF-truncated sequences one byte at a time, and work with
  configured capacity one. Bulk reads can enable only byte rollback; all
  state-changing attempts invalidate prior permissions as documented.
- Shared offset and rollback permissions live in the existing reader handle.
  Physical storage is max(configured capacity, 4), while capacity() and every
  underlying read remain bounded by the configured capacity. Confirmed scalar
  prefixes survive provider errors without becoming consumed or mixing in bytes
  modified by a failed call. All buffered operations see the same prefix.
- Parallel fixtures cover every byte, UTF-8 widths and malformed sequences,
  tiny capacities, aliases, single rollback, invalidation, buffered composition
  and failed-call side effects. Public documentation and source navigation
  include the new methods. No syntax, native/runtime boundary or ecosystem
  dependency changes are needed.
- Stage2 build, the complete project093 fixture, all 97 query tests and
  CSV/MessagePack/Rope/Archive ecosystem verification pass, including independent
  consumers, cached builds and configured race tests. Invalid provider counts
  after actual side effects and buffered lengths exceeding tiny configured
  capacities have explicit regression coverage. Golden regeneration passes
  29 formatter, 17 pipeline and 51 integration tests. Full CI passes 1010
  compiler tests, 171 driver tests, bootstrap fixed point, helper, packaging
  and release-smoke checks.
  Logs use _artifact/stdlib-io-buffer-runes-.
  Counts remain 22 accepted, 2 started and 50 planned. Delimiter fragments,
  writer transfer/typed composition and synchronous Pipe still prevent io/bufio
  acceptance; full implementation and user review still precede release.

### S2 shared seek contract — implemented and validated

- Added std::io::Seek using the existing SeekFrom enum and checked signed byte
  offsets. Cursor, StringReader, SectionReader and OffsetWriter implement it
  without adding native hooks, buffered seek policy or implicit I/O.
- Concrete boundaries remain explicit: Cursor cannot create holes; StringReader
  permits beyond-EOF positions and clears character rollback on every attempt;
  SectionReader uses section-relative positions and only successful seeks clear
  pending errors; OffsetWriter checks base-plus-position and rejects End as
  Unsupported because WriteAt exposes no length.
- Existing setters and StringReader's inherent seek are retained, with the
  inherent and generic paths sharing implementation. New fixtures exercise
  generic Read/Seek composition, aliases, byte versus scalar boundaries, limits,
  integer overflow and positional-provider no-I/O behavior. Source packaging,
  navigation, completion and language documentation include the trait; no
  compiler/driver implementation starts consuming this unreleased API.
- Stage2 build, the complete project093 fixture and all 97 query tests pass.
  Golden regeneration passes 29 formatter, 17 pipeline and 51 integration
  tests. Full CI passes 1010 compiler tests, 171 driver tests, bootstrap fixed
  point, helper, packaging and release-smoke checks.
  Logs use _artifact/stdlib-io-seek-.
  Counts remain 22 accepted, 2 started and 50 planned. This closes another io
  contract gap, not the full io/bufio rows or the overall review/release gate.

### S2 buffered peek/discard and progress copy — implemented and validated

- Added std::io::copy_progress/copy_buffer_progress and CopyError, reporting
  separate confirmed read/write counts without changing existing copy APIs.
  The new engine handles successful short transfers directly and stops on all
  errors without retry, write_all overrides, flushing or closing. Failed-call
  side effects remain unknown; counts are not automatic recovery offsets.
  Counter exhaustion is checked before another read, including at exact u64
  maximum. This unexecutable-sized boundary is reviewed algebraically, not
  claimed as an executed extreme-size test.
- BufRead::discard skips across fills and returns TransferError with successful
  consume progress. BufReader::peek compacts and refills within fixed capacity,
  does not consume, and exposes the confirmed partial view through PeekError.
  Invalid requests and zero requests avoid I/O. Peek errors/EOF are not sticky;
  view lifetimes, shared state and explicit retry limitations are documented.
- Parallel external fixtures cover short I/O, shared handles, capacity and
  range checks, EOF, failed-call mutations/consumption and exact confirmation
  boundaries, including peek/discard/copy and peek/line-reading composition.
  Resource packaging, navigation and completion include the new
  public surface; compiler/driver implementations do not consume these APIs.
- Stage2 build, the complete project093 fixture and CSV/MessagePack/Rope/Archive
  consumer, cache and configured race verification pass. Query verification
  exposed missing navigation for inherited default trait methods: their synthetic
  function identity now maps through exact trait metadata to the declaration,
  with explicit same-name trait disambiguation coverage. All 97 query tests pass.
  Golden regeneration passes all 29 formatter, 17 pipeline and 51 integration
  tests. Full CI passes 1010 compiler tests, 171 driver tests, bootstrap fixed
  point, helper, packaging and release-smoke checks. Logs use
  _artifact/stdlib-io-peek-copy-. Ownership stays std::io; no runtime/FFI changes
  or ecosystem dependency are introduced. Counts remain 22 accepted, 2 started
  and 50 planned; io/bufio still require their remaining mapped capabilities.
- Outstanding work includes byte/scalar rollback, delimiter fragments,
  writer transfer/typed-write composition, shared seeking contracts and the
  synchronous Pipe with close/error ordering and backpressure. Existing runtime
  primitives may be composed for Pipe; adding a runtime primitive remains out
  of scope. Earlier historical notes about missing peek/discard are superseded
  by this implementation record, not by acceptance of the entire buffer row.

### S2 conservative generic read failures — implemented

- Removed automatic Interrupted retries from Read::read_exact/read_to_end,
  copy_buffer, BufReader, Scanner, endian varint/fixed/block readers, Base32
  decoding and ecosystem CSV/MessagePack/Rope readers. Short successful reads
  still complete normally. Every provider error stops the current operation;
  a failed call may have consumed input or modified its buffer without reporting
  progress, so replay cannot be assumed safe. Specific descriptor syscall EINTR
  handling is unchanged.
- Base32 decoder source failures now retain their original kind/details in both
  returned and sticky errors. The encoder's existing write-error translation and
  TeeReader's mirror-error translation remain unchanged for compatibility; Tee
  source errors still propagate directly. Scanner/CSV/MessagePack preserve their
  terminal-state contracts, including CSV's completed CR row before its saved
  lookahead error.
- Parallel tests inject finite Interrupted failures after actual source
  consumption, both initially and after successful prefixes. They cover generic
  helpers, SectionReader pending errors, buffered/scanning paths, codecs and
  ecosystem streams. Ordinary success fixtures no longer depend on injected
  interruption retry. Documentation records the intentional behavior change,
  confirmed-progress boundaries and unsafe whole-operation replay.
- Formatting, stage2 build and the complete project093 fixture pass. CSV,
  MessagePack, Rope and Archive ecosystem verification passes, including
  independent consumers, cached builds and configured race tests. Archive tests
  now distinguish TAR's terminal input failure from ZIP's pre-output read
  failure and verify consumed-byte interruptions without replay. Golden update
  passes all 29 formatter, 17 pipeline and 51 integration tests. Full compiler
  and driver suites pass 1008 and 171 tests respectively. This batch does not
  rerun the bootstrap fixed point or release packaging; the preceding minimum
  read/reset batch's full CI remains the latest evidence for those checks.
  Logs use _artifact/stdlib-read-interruption-. This is further io/bufio work,
  not acceptance of either full row. Counts remain 22 accepted, 2 started and
  50 planned; full implementation and user review still precede release.

### E1 HTML final acceptance audit — validated

- Read-only comparison with the mapped Go html utility family found no missing
  production function: shared escaping, 2,125 pinned entity names, 106 legacy
  names, text/attribute decoding, numeric recovery and bounded output are already
  implemented. HTML parsing and template context analysis remain separate rows.
- Parallel test hardening adds all pinned names in attribute context, legacy
  letter/digit/equals ambiguity and semicolon disambiguation, longest/unknown
  names, both-context byte-limit matrices and structured error payloads. The
  independent consumer now checks Go EscapeString as well as UnescapeString,
  ASCII/Unicode round trips, hexadecimal/scalar boundaries and long prefixes.
  Go short-numeric and overflow quirks use explicit correct expected results,
  not a misleading claim of exact Go parity. The expanded consumer discovered
  that the short-decimal quirk includes non-semicolon terminators, not only EOF;
  Go also decodes empty hexadecimal references incorrectly. Inspection of the
  local Go source confirmed both, and the tests/README now preserve literal
  empty-hex input and decode single-digit decimal references with explicit
  expected values. Production semantics did not change.
- Production implementations and module ownership are unchanged. Historical
  html/template/markdown/tui_markdown verification passed; a fresh joint run
  passes in _artifact/stdlib-html-acceptance.log. The first attempt stopped at consumer
  formatting; formatting under the private registry is now complete. A later
  consumer run exposed the Go-oracle differences described above; the corrected
  verifier passes all four libraries and consumers, cached builds, entity data
  conformance and terminal PTY checks. The mapped HTML row is accepted above.

### S1 minimum reads, no-op closing and buffer reuse — validated

- Parallel work added io::read_at_least with TransferError confirmed progress,
  preflight minimum validation, zero-minimum no-I/O behavior and conservative
  propagation of every provider error. Successful calls can fill beyond the
  minimum up to the supplied view. Exact reading composes minimum=view length.
  Failure-call side effects remain unconfirmed and are not automatically retried.
- NoopCloser supplies a transparent Read wrapper and always-successful Close
  without forwarding close or disabling later reads. Resource-scope tests cover
  action errors, aliases, raw counts/errors and underlying-close suppression.
- BufReader/BufWriter expose capacity and shared-handle reset; BufWriter exposes
  logical available space. Their stream fields now use shared references so all
  aliases switch together while retaining buffer storage. Reset intentionally
  drops unread/pending bytes without I/O or closing the old stream. Tests cover
  partial failed drains, replacement close, capacity boundaries and snapshots.
- The existing source catalog already includes the three edited io files. New
  external fixtures and query navigation/completion cases exercise the public
  APIs. Formatting, toolchain build and the focused project093 execution pass;
  full ecosystem verification of CSV, msgpack, rope and archive also passes,
  including independent consumers, cached builds and applicable race checks.
  Logs use _artifact/stdlib-io-minimum-reset-. Initial query navigation failures
  came from stale executable-relative test/lib resources, not a production type
  inference defect: a fresh-resource independent runner accepts both original
  and byte-literal cases. The test resources are now refreshed, the original
  navigation case is retained, and a three-case source-analysis regression was
  added. Final query verification passes all 93 tests in
  _artifact/stdlib-io-minimum-reset-query-final.log. Golden regeneration passes
  all 29 formatter, 17 pipeline and 51 integration tests; see
  _artifact/stdlib-io-minimum-reset-golden.log. Full CI passes all 1008 compiler
  tests, 171 driver tests, bootstrap fixed point, helper and packaging checks;
  see _artifact/stdlib-io-minimum-reset-ci.log.
- A separate read-retry audit found that legacy generic read helpers still
  retry Interrupted without knowing whether the failed call consumed data.
  TeeReader already maps sticky mirror interruptions to Other, avoiding that
  particular loop, but arbitrary source failures and SectionReader pending
  interruptions can still be swallowed. Uniform generic read-loop correction
  remains required across io, endian/base32 streams and msgpack/rope/CSV readers;
  concrete syscall EINTR policies are distinct. This batch does not claim those
  existing paths are fixed, nor accept the complete io/bufio rows.
- Remaining buffer work includes cross-fill bounded peek/discard, byte/scalar
  rollback and delimiter fragments, plus writer transfer/typed-write composition.
  A synchronous Pipe can be investigated using existing Channel/select and task
  primitives: channel-backed state gates and close notifications already exist
  in std::resource. It is not excluded as R merely because it is concurrent;
  endpoint close/error ordering, backpressure and race tests remain necessary.

### S1 text composition acceptance and CSV write consistency — validated

- Three parallel tasks delivered string comparison/prefix/suffix/byte-search
  compositions, an explicit bounded legacy word-title recipe, and CSV's remaining
  generic write-error replay. Existing public text/bytes/cmp/Unicode APIs provide
  the composition mechanics; no duplicate standard aliases or native production
  algorithms are needed. The word-title recipe belongs to caller policy, not a
  universal Unicode segmentation contract.
- The language guide collects empty-pattern differences and the bytes-to-text
  repair boundary. CSV retains its ecosystem ownership. Its specific write loop
  now stops on Interrupted like other generic writers; concrete Linux descriptor
  syscall retries are a separate contract and were not mechanically removed.
  Direct, MultiWriter and partially successful OffsetWriter tests verify that
  subsequent writes/flushes retain the original failure without replay. CSV's
  reported position excludes unconfirmed effects inside a failed call.
- Project093 passes the string composition matrix, 13 strings across all 256
  byte needles, membership/predicate short-circuit tests, and legacy title versus
  scalar title differential tests including ASCII/Unicode boundaries and exact,
  insufficient and negative output limits. See
  _artifact/stdlib-text-composition-focused.log.
- Complete CSV ecosystem verification passes, including library/independent
  consumer tests, cached builds and library/consumer races; see
  _artifact/stdlib-text-composition-csv.log. Golden regeneration passed all 29
  formatter, 17 pipeline and 51 integration tests; see
  _artifact/stdlib-text-composition-golden.log. The previous full CI validated
  unchanged standard production code; the new tests and CSV change have the
  additional evidence listed here.

### S1 text capacity/copy and conservative write failures — validated

- Parallel implementation added StringBuilder::capacity/reserve_checked and
  text::clone_checked. Reservation validates negative/overflow/total-byte bounds
  before touching storage; detached copying uses existing byte/string primitives.
  Tests cover aliases, snapshots, capacity reuse, invalid UTF-8 byte appends and
  source-storage detachment using a test-only pointer-range oracle. No production
  native algorithm backend was added.
- Parallel I/O audit identified replay bugs: MultiWriter may have written earlier
  targets before returning Interrupted, while OffsetWriter may discard reported
  partial progress when converting a positional failure to sequential Write.
  Generic retry cannot infer zero progress from this error. Write::write_all,
  BufWriter draining, StringReader::write_to and text::stream::write_replaced now
  return all write errors immediately, including Interrupted. Read retries are
  unchanged. This is an intentional behavior change, not signature compatibility
  disguised as identical retry behavior.
- Confirmed counts exclude side effects in a failed call. BufWriter retains an
  unconfirmed suffix, not necessarily an unsent suffix; retrying it requires
  provider-specific knowledge. Tests exercise MultiWriter/OffsetWriter, forward
  wrappers, nested buffered write_all, direct faults and Interrupted after a
  successful prefix. Existing successful short-write fixtures no longer depend
  on injected automatic retry; independent interruption cases remain.
- Updated msgpack and rope consumer tests and rope/archive write-contract prose.
  The audit also found CSV's separate explicit write-retry loop; that is tracked
  as a remaining ecosystem consistency issue, not silently covered by this fix.
  The io/bufio audit still requires progress-aware generic copy/read-at-least,
  broader buffered operations, generic seek/close adapters and a Pipe feasibility
  check using existing runtime primitives. None is reclassified as R merely to
  reduce scope.
- Unified formatting, toolchain build, project093 and project085 runtime tests
  pass. Module tests pass for msgpack (21) and rope (12). Evidence logs use the
  _artifact/stdlib-text-io-safety- prefix. Full ecosystem verification passed for
  msgpack, rope and archive, including independent consumers, cached builds and
  applicable library/consumer race checks; see
  _artifact/stdlib-text-io-safety-ecosystem.log. Golden regeneration passed all 29
  formatter, 17 pipeline and 51 integration tests; see
  _artifact/stdlib-text-io-safety-golden.log. Full CI passed: 1007 compiler tests,
  171 driver tests, bootstrap fixed point, helpers, packaging and smoke checks;
  see _artifact/stdlib-text-io-safety-ci.log. Counts stay
  at 20 accepted, 4 started and 50 planned.

### S1 compiled multi-rule text replacement — validated

- Added standard text::Replacer with a private compiled byte trie and snapshotted
  ordered rule pairs. First matching rule wins rather than longest match; duplicate
  keys retain the earliest rule, and emitted replacement text is never rescanned.
  Construction checks rule count and combined pattern/replacement bytes before
  allocating compiled state. Existing text/Unicode ownership remains unchanged.
- Replacement has independent output-byte and trie-work budgets, recoverable
  errors without partial output and no mutation of compiled state. Work counts
  root lookups and attempted byte transitions. Empty rules match once at a
  position, then nonempty rules can still match there; forward progress respects
  GoML scalar boundaries instead of Go's byte-wise empty insertion into UTF-8.
- Added Go NewReplacer references for compatible cases and an independent scalar
  scanner oracle for the intentional Unicode empty-pattern policy. Tests cover
  ordered/duplicate/overlapping/empty rules, deletion, exact/short construction
  and output limits, exact work exhaustion, retries and input-rule snapshots.
  Source catalog/counts, public navigation, docs and examples are updated.
  Build and focused runtime tests pass, including long repetitive-prefix work
  exhaustion and successful reuse. Golden verification passed all 51 integration
  tests in _artifact/stdlib-replacer-golden.log. Full CI completed successfully
  in _artifact/stdlib-replacer-ci.log, including 1006 compiler tests and 171
  driver tests. This validates the whole-string replacement implementation;
  incremental writer output remains part of the open audit below.
  This advances the
  existing strings row, not acceptance of that full row or the full migration.

### S1 strings acceptance audit — closed

The local Go 1.26 strings sources were compared with public text APIs and the
project093 fixtures. The follow-up batches closed the previously identified
family gaps and passed their verification gates. The mapped row is accepted
above; this does not accept io/bufio or expand strings into locale segmentation.

| Family | Current evidence | Acceptance disposition |
| --- | --- | --- |
| Search, cuts and splitting | text_search.gom and text_split.gom exercise Finder, cuts, bounded splitting and iterators against Go references, plus byte-only search, comparison, membership and prefix/suffix compositions. | Passed final composition/golden checks; intentional empty-separator policies are documented together. |
| Scalar predicates, casing and folding | text_unicode.gom and text_transform.gom exercise Unicode predicates, fields, trimming, mapping and checked transformations; text_title.gom proves the explicit legacy word-title recipe. | Passed final composition/golden checks; scalar title, word-title policy and invalid UTF-8 repair through bytes are distinct documented contracts. |
| Building and snapshots | text_append.gom, text_capacity.gom and text_clone.gom exercise checked appends, reservation, byte validation, aliases, snapshots and detached copying; the safety batch passed full CI. | No remaining family-specific gap identified; include this evidence in final strings-wide acceptance. |
| Compiled replacement | replacer.gom covers rule order, empty matches, limits, snapshots and retry after failure; replacer_stream.gom covers bounded incremental writer output and failure progress. | Included in the successful full safety CI and subsequent golden verification. |
| String-backed reading | io::StringReader supplies read-only sequential/positional bytes, reset, byte/scalar reads and rollback, checked seek including beyond EOF, and confirmed-progress writer transfer; full CI includes fault-injection and composite-writer tests. | No remaining family-specific gap identified; the writable Cursor remains a separate contract. |

String I/O adapters belong to std::io so text algorithms need not import I/O.

The builder/copy batch above now supplies builder-level capacity/reservation and
bounded detached copying, with negative/overflow/output-bound, alias, snapshot
and byte-write UTF-8 tests. A test-only pointer-range check establishes detachment
from a larger source allocation; substring/value equality alone cannot do so.
Full CI verified this batch. Capacity bounds distinguish
logical requested bytes from allocator overhead; there is no general
allocation-failure recovery or global memory quota.

The replacement engine exposes fallible, incremental text fragments;
whole-string replacement and an opt-in text::stream writer adapter consume
the same engine. This follows the existing hash::stream split: the adapter may
depend on both text and I/O, while text itself remains independent of I/O. Do not
copy the trie algorithm into the adapter or materialize the complete output to
claim streaming support. Each operation has independent cursor and budget
state; exhausted and failed cursors do not silently resume or repeat output.
Sequential write failures expose known
progress without promising rollback of an external writer, and budget failures
are not disguised as transport errors. Tests cover fragmented and failing writers,
EOF and zero-length reads, Unicode byte boundaries, failed unread transitions,
seek overflow, reset and reuse.

### S1 incremental replacement fragments — validated

- Added Replacer::chunks returning a fallible, lazy iterator of nonempty UTF-8
  fragments. Whole-string replace now consumes this same engine instead of owning
  a second matching loop. No new package, native backend or I/O dependency was
  introduced; the planned writer adapter remains separate work.
- Limit validation is eager, matching is lazy, and output/work budgets cover the
  full iteration. A failing fragment is not emitted; one error terminates the
  iterator permanently. Deletions still charge work, aliases share progress,
  and fresh calls have independent state. Fragments may retain input/rule data
  and are not individually size-bounded.
- The existing exhaustive rule/reference comparisons now also concatenate the
  lazy fragments and compare their output. Added exact exhaustion, late work and
  output failures, fused error/EOF behavior, alias progress and invalid limits.
  Build and focused runtime tests passed in
  _artifact/stdlib-replacer-chunks-build.log and
  _artifact/stdlib-replacer-chunks-focused.log. Navigation coverage and public
  documentation were added. Initial broad compiler regression verification found
  two stale pipeline snapshot shards and an invalid unwrap call in the navigation
  fixture (1003 passed, 3 failed). Subsequent golden regeneration passed and the
  corrected navigation fixture passed in the 1006-test compiler recheck recorded
  with the writer adapter below.
  This batch does not inherit the earlier whole-string implementation's full-CI
  result.

### S1 replacement writer adapter — validated

- Added opt-in std::text::stream::write_replaced on the public chunk iterator and
  io::Write. It uses a fixed 4096-byte scratch buffer, completes short writes,
  retries Interrupted and rejects zero/invalid progress without materializing the
  complete output. It neither closes nor flushes the caller's writer.
- ReplaceWriteError preserves replacement versus I/O failures and the confirmed
  written byte count. Existing destination bytes are not rolled back. Unknown
  side effects of a failing or invalid-count writer cannot be included; retrying
  the entire operation is not resumable and can duplicate earlier output.
- Added a fault-injecting writer fixture covering all failure offsets of small
  Unicode replacements, short writes, interrupted retry, zero/negative/oversized
  counts, empty output, budget failures and a replacement exceeding scratch
  capacity. Source catalog, dependency/purity checks, navigation and public docs
  include the new package. Module testing exposed a missing standard-package
  selection flag, now added with its dependency-selection regression case.
  The public bound is explicitly io::Write for cross-package specialization.
- Toolchain build and the complete focused module pass in
  _artifact/stdlib-replacer-stream-build.log and
  _artifact/stdlib-replacer-stream-focused.log. Golden regeneration passed all
  29 formatter, 17 pipeline and 51 integration tests in
  _artifact/stdlib-replacer-stream-golden.log, resolving the two stale snapshot
  shards noted above. Full CI in _artifact/stdlib-replacer-stream-ci.log completed
  with 1005 compiler tests passing and one navigation-fixture failure: its private
  helper triggered an unused-function diagnostic. All 171 driver tests passed.
  Both navigation helpers are now public. Compiler verification subsequently
  passed all 1006 tests in
  _artifact/stdlib-replacer-stream-compiler-recheck.log, including navigation,
  dependency selection and source-only package checks. The implementation was
  unchanged by this fixture correction; the earlier full-CI invocation remains
  recorded as failed, not retrospectively green. The strings row is open and
  counts remain unchanged.

### S1 read-only string reader — implemented

- Added io::StringReader with shared-handle reset, size/remaining-length queries,
  sequential and positional byte reads, byte/scalar reads and rollback. The
  immutable string is retained rather than copied into a writable Cursor.
  Character decoding reuses the existing scalar primitive; seeking into an
  encoded scalar consumes one replacement byte per invalid start.
- SeekFrom provides Start/Current/End inputs for the concrete seek method.
  Checked arithmetic rejects negative/overflowing results without moving the
  position, while nonnegative positions beyond EOF are allowed. Positional reads
  preserve position and rollback state, report partial progress and validate
  ranges before touching the destination. Shared handles require serialized use.
- Empty/EOF sequential reads deliberately invalidate character rollback; this
  state rule is documented rather than claiming every Go Reader quirk. The
  progress-preserving transfer addition is recorded separately below; this is not
  full strings or io acceptance. No new runtime/backend or generic seek trait was added.
- Added external tests over every byte offset of representative UTF-8 strings,
  Go strings.Reader scalar/width references, short and positional reads, untouched
  invalid-range destinations, aliases/reset, EOF, unread transitions and signed
  seek boundaries. Toolchain build and focused tests, including the additional
  transition cases, passed in _artifact/stdlib-string-reader-build.log and
  _artifact/stdlib-string-reader-focused.log. Golden regeneration passed all 29
  formatter, 17 pipeline and 51 integration tests in
  _artifact/stdlib-string-reader-golden.log. Broader verification is pending. Source catalog,
  public navigation, examples and comparison guidance include the reader.

### S1 string reader transfer — implemented, verification pending

- StringReader::write_to completes short writes through a bounded scratch buffer,
  retries Interrupted, rejects zero/invalid counts and advances the reader only
  after confirmed write progress. TransferError retains the count for the current
  call. A second call can transfer the remaining suffix, including from a byte
  inside a scalar, without the read-ahead loss of generic read-then-write copying.
- The operation clears character rollback even at EOF, does not implicitly flush
  or close, and does not call the writer for empty input. Concurrent/reentrant
  mutation through reader aliases is unsupported. Unknown side effects of a
  contract-violating writer cannot be counted or undone. Short-write completion
  follows GoML conventions rather than Go Reader.WriteTo's single-write policy.
- Added fault injection at each byte of small Unicode inputs, variable short
  writes, interrupted retry, zero/negative/oversized counts, suffix resumption and
  a 12KB input failing at byte 5000. Navigation and docs cover the method.
  Toolchain build and focused runtime tests pass in
  _artifact/stdlib-string-reader-transfer-build.log and
  _artifact/stdlib-string-reader-transfer-focused.log. Golden regeneration passed
  all 29 formatter, 17 pipeline and 51 integration tests in
  _artifact/stdlib-string-reader-transfer-golden.log. Full CI for the combined
  reader and transfer batch finished in _artifact/stdlib-string-reader-ci.log
  with 1005 compiler tests passing and one navigation-test helper panic; all 171
  driver tests passed. The helper searched by slicing at arbitrary byte offsets,
  so the new Unicode navigation example exposed an invalid scalar boundary.
  It now uses string::rfind, with a regression covering Unicode, missing and empty
  patterns. Unicode navigation examples are retained unchanged. The first recheck
  exposed a Debug bound in the new assertion; using direct equality assertions
  fixed that test-only issue. All 92 query tests subsequently passed in
  _artifact/stdlib-string-reader-query-recheck.log, including navigation and the
  UTF-8 helper regression. The complete compiler recheck remains outstanding;
  inventory acceptance counts are unchanged.

### S1 conversion capability acceptance audit — validated

The reviewed strconv row maps to conversion in num and quotation in text, not
to a new Go-compatible module. The final audit covers each functional family:

| Requirement | Current implementation and external evidence |
| --- | --- |
| Boolean conversion | num boolean parsing and canonical bool ToString; boolean.gom verifies accepted spellings, malformed strings and canonical formatting. |
| Integer conversion and destinations | num radix/width parsing, formatting and transactional buffer writes; integer_width.gom and integer_format.gom cover every radix/width, boundaries, signs, separators, destination offsets and insufficient capacities. |
| Floating parsing and error classification | Structured f32/f64 parsers over the pure scalar implementation; float_parse.gom covers bits, grammar, syntax/range precedence, signed underflow, exact long mantissas and exponent cancellation. |
| Floating representation | Exact fixed/scientific/general/shortest/binary/hex formatters; float_format.gom, float_shortest.gom and float_power.gom compare Go outputs with precision, output and IEEE boundaries. |
| Numeric casts | ToFloat, TryToInt and explicit rounding; float_conversion.gom compares all integer widths, both float widths, all rounding modes, nonfinite/range errors and double-rounding-sensitive inputs. |
| Complex text conversion | complex.gom returns typed component pairs and bounded formatted text without an arithmetic type dependency; complex_parse.gom and complex_format.gom cover grammar, component bits, all notation families and complete framing limits. |
| Quotation and classification | Printable/ASCII/graphic string and scalar quoting, raw-literal eligibility and existing Unicode predicates; text_quote.gom covers Unicode, escapes, controls, limits and reference comparisons. |
| Complete/prefix/element decoding | Byte-preserving and strict-UTF-8 unquoting, bounded prefix recognition and typed byte/scalar elements; text_quote.gom and quote_element.gom cover grammar, limits, offsets, contexts and untouched tails. |
| Append composition | Checked integer destinations and checked StringBuilder writes; text_append.gom compares Go append sequences and verifies per-call atomicity, existing prefixes, alias behavior, snapshots and total-byte limits. |
| Source/metadata/tooling | num/text source catalog, source-only checks, completion/navigation, external project093, docs/examples and comparison guidance. No standard-to-ecosystem import or new native algorithm backend. |

Complex arithmetic remains the independent math/cmplx row. Aliases such as Atoi,
Itoa and per-conversion Append names compose existing typed APIs rather than
adding duplicate algorithms. Invalid UTF-8 source strings are not representable;
raw escaped bytes use the explicit byte decoding surface. The documented one-bit
integer and long-float differences are verified mathematical corrections, not
unexplained divergence from the reference. Numeric input size/work and retained
strings remain caller-bounded; formatter limits are output bounds, not total
allocation or allocation-failure guarantees. Shared mutable builders require
serialized access. Compiler/driver-owned fallbacks are retained until release
and stage0 advancement.

The latest compiler suite passes all 1006 tests. Final full CI passes in
`_artifact/stdlib-num-final-ci.log`, including 171 driver tests, helper race,
bootstrap fixed point and packaging checks. The mapped row is accepted above;
counts are 20 accepted, 4 started and 50 planned. The complete A/B implementation,
user-review and release gate is unchanged.

### S1 bounded complex component formatting — validated

- Added format_complex_f32/f64 with ComplexFormat selecting shortest decimal,
  fixed/scientific/general decimal, binary or hexadecimal output. Both component
  widths call their existing scalar formatter directly; no duplicated conversion
  algorithm, complex runtime type or math-package dependency was introduced.
- Output always includes both components, parentheses and trailing i. Existing
  imaginary signs are retained, positive infinity is not given a duplicate plus,
  and NaN receives the joining plus. Checked final framing prevents length
  arithmetic overflow and respects the complete output limit; independently
  bounded component temporaries are not a total-allocation guarantee.
- Added Go FormatComplex references for all modes, cross-product special values,
  f32/f64 boundary values, deterministic random pairs, exact/short limits, huge
  precisions and negative-limit/invalid-precision precedence. Public navigation,
  documentation and comparison guidance are updated. Build and focused runtime
  tests pass; golden generation passes all 29 formatter, 17 pipeline and 51
  integration tests. Full compiler verification passes all 1006 tests in
  `_artifact/stdlib-complex-format-compiler.log`. Final conversion-row acceptance
  still requires its complete audit.

### S1 complex component parsing — validated

- Added num::parse_complex_f32 and parse_complex_f64 returning real/imaginary
  component tuples. Precision suffixes name component widths; no complex value
  type, arithmetic, source syntax or num-to-math dependency was introduced.
  Decimal/hexadecimal components reuse existing pure scalar parsers, preserving
  direct f32 rounding, signed zero and nonfinite values.
- Supports real-only, imaginary-only, combined and singly parenthesized forms.
  Distinguishes joining signs from exponent signs and retains plus-negative/NaN
  compatibility. Syntax in either component precedes another component's range
  error; errors return no partial/saturated components or retained input strings.
- Registered the new public source, navigation and source-only coverage. Added
  test-only Go ParseComplex references across both widths, signed zeros, special
  values, grammar combinations, exponent signs, overflow/syntax precedence and
  generated bit-pattern component strings. Build and focused runtime tests pass;
  golden generation passes all 29 formatter, 17 pipeline and 51 integration tests.
  Full CI passes in `_artifact/stdlib-complex-parse-ci.log`, including 1006 compiler
  tests, 171 driver tests, helper race checks, packaging and bootstrap fixed point.
- Formatting remains outstanding for strconv; full complex arithmetic belongs to
  the separately planned math/cmplx row. Overall counts and review gates remain
  unchanged rather than accepting either row from this parser alone.

### S1 checked text-builder append composition — validated

- Added total-byte-bounded StringBuilder string, scalar and line append methods.
  Each call validates before buffer growth or mutation; a line's text and LF
  are one transaction. Reuses existing TransformError and length checks, retaining
  unchecked methods, shared handles and public storage without a new object model.
- Added prefix/capacity matrices, empty/oversized existing contents, Unicode byte
  widths, negative limits, alias visibility, unchanged buffers on failure and
  independent finished string snapshots. Numeric/boolean/graphic-quote sequences
  compare composition with Go's AppendInt/Uint/Float/Bool/QuoteToGraphic, including
  signed zero, infinities and NaN. Earlier successful appends are not rolled back.
- Added public navigation and documented per-call limits, temporary conversion
  budgets and concurrency/public-storage limitations. Build and focused runtime
  tests pass; golden generation passes all 29 formatter, 17 pipeline and 51
  integration tests. Full compiler verification passes all 1006 tests in
  `_artifact/stdlib-text-append-compiler.log`. This advances existing
  strings/strconv scope without adding aliases
  for each Go append entry point or claiming either whole row accepted.

### S1 typed quoted-element decoding — validated

- Added QuoteContext and QuotedElement with unquote_element returning one byte
  or Unicode scalar plus the untouched tail. Quote context controls escaped and
  unescaped delimiters. Unicode escapes retain scalar identity even below 128;
  hexadecimal/octal escapes retain byte identity even when not valid UTF-8.
  Empty/malformed elements report Syntax(0) with no partial value. Literal CR/LF
  remain valid elements, with enclosing-literal validation left to callers.
- Refactored the shared escape parser to produce the typed value directly;
  complete unquoting converts it to bytes while prefix validation discards it.
  The no-delimiter context does not accidentally permit backslash plus NUL.
  Per-call input inspection/output scratch is bounded by a single scalar/escape.
- Added independent Go UnquoteChar references for all contexts, Latin-1 input,
  all three-digit octal values, malformed Unicode/byte escapes, scalar sampling,
  byte/scalar flags, tail preservation, sequential consumption and navigation.
  Build passes after correcting a string emptiness check to byte_len; focused
  runtime tests and golden generation pass (29 formatter, 17 pipeline and 51
  integration tests). Full compiler verification passes all 1006 tests in
  `_artifact/stdlib-quote-element-compiler.log`.

### S1 bounded quoted-prefix recognition — validated

- Added text::quoted_prefix, validating one Go-style literal at byte zero and
  returning its original spelling, including delimiters, escapes and raw CRs.
  The source-byte output limit includes delimiters and stops prefix scanning;
  trailing text is ignored. Reuses escape validation with bounded per-escape
  scratch, without constructing the decoded body. Negative limits precede syntax;
  syntax/output errors retain byte positions and return no partial prefix.
- Generalized the shared OutputLimit diagnostic wording to cover both decoded
  output and preserved quoted-prefix output; error variants/payloads are unchanged.
- Added public navigation, documentation and Go QuotedPrefix differential checks
  across the existing literal grammar matrix, malformed escapes, all insufficient
  prefix limits, exact limits, Unicode boundaries, raw CR/LF and unexamined tails.
  Build and focused runtime tests pass. An initial golden run stopped before
  tests because make attempted to replace a compiler used by the focused run;
  after that process finished, the serial retry passed all 29 formatter, 17
  pipeline and 51 integration tests. Full compiler verification passes all 1006
  tests in `_artifact/stdlib-quote-prefix-compiler.log`.
- This implements the prefix part of the pending conversion surface. Single
  escaped-element decoding, complex conversion and append-composition acceptance
  remain outstanding; the overall scope/counts and review gate are unchanged.

### S1 graphic quotation and raw-literal eligibility — validated

- Added bounded text::quote_graphic and quote_char_graphic, reusing the same
  private quoting engine and pinned Unicode classification. The new mode
  preserves Zs characters without relaxing delimiter/backslash/control escaping;
  printable and ASCII modes retain their existing contracts and output preflight.
- Added allocation-free-output can_backquote over valid strings: reject backticks,
  U+FEFF, DEL and ASCII controls except tab, retaining Go's distinction between
  raw-literal eligibility and graphic/terminal-safe text. No source grammar or
  native production backend changes.
- Added Go reference comparisons, exact/insufficient/negative output limits,
  decoded round trips, Latin-1 and sampled Unicode scalars, separator ranges,
  embedded BOMs and public navigation. Build and focused runtime tests pass.
  The initial compiler suite passed 1004 tests with two stale IR snapshots;
  `just update-golden` regenerated them and passed all 29 formatter, 17 pipeline
  and 51 integration tests. Compiler recheck passes all 1006 tests in
  `_artifact/stdlib-graphic-quote-compiler-recheck.log`. The immediately preceding
  exponent correction passed full CI; no packaging or bootstrap contract changes
  were introduced by this source/API batch. Overall counts remain 19 accepted,
  5 started and 50 planned, with the remaining conversion work listed below.

### S1 remaining conversion surface audit

The scalar numeric and complete-literal implementations do not yet accept the
entire strconv capability. Comparing the installed Go 1.26 public source against
the mapped GoML owners identifies these remaining capabilities:

| Capability | Existing coverage or remaining work |
| --- | --- |
| Boolean, integer/radix and scalar floating conversion | Public num APIs, explicit precision/notation/output bounds and checked integer destinations are implemented and tested. |
| Printable/graphic classification | Reuse existing std::unicode::is_printable/is_graphic; no duplicate tables or strconv aliases. |
| Printable and ASCII string/scalar quotation; complete unquoting | Implemented in std::text with bounded output and explicit byte/UTF-8 decoding policies. |
| Graphic string/scalar quotation | Implemented bounded variants preserving Unicode Zs characters; validation is recorded above. |
| Raw backquote eligibility | Implemented public predicate with explicit CR, control, delimiter and UTF-8 rules; validation is recorded above. |
| Quoted prefix and single escaped-element decoding | Bounded prefix recognition and typed byte/scalar element decoding are implemented; validation is recorded above. |
| Complex parsing/formatting | Component parsing and bounded formatting are implemented in num without a dependency on a complex arithmetic type; validation is recorded above. |
| Append-style conversion | Checked integer destinations and total-byte-bounded text-builder composition are implemented; validation and failure/aliasing contracts are recorded above. |

This records remaining work without changing package ownership, reducing the
74-row scope, or marking strconv accepted.

### S1 input-sized floating exponent saturation — validated

- Reproduced the fixed-cutoff bug before changing it: `0x0.` followed by 250000
  zeros and `1p10000001` returned 0.0625 rather than a range error. Evidence is
  `_artifact/stdlib-float-exponent-before.log`. The old loop stopped accumulating
  after reaching 1000000, dropping the final exponent digit.
- Replaced the fixed threshold with a checked saturating accumulator bounded by
  four times input byte length plus 2048, with a machine-integer ceiling. Each
  mantissa byte contributes at most four binary scale positions; exponents past
  this bound cannot be canceled into the finite IEEE range by that input.
  Overflow checks precede decimal accumulation, and syntax scanning continues.
- Added mathematical regression cases for huge overflow, separators and invalid
  suffixes, plus exact-one cancellation with 2.5-million-zero hexadecimal and
  10-million-zero decimal fractions at both widths. These deliberately use exact
  known answers rather than Go's own bounded exponent/digit-buffer behavior.
  Toolchain build and focused runtime tests pass. The reproducer now reports a
  range error in `_artifact/stdlib-float-exponent-after.log`; all cancellation
  cases pass. Golden generation passes all 29 formatter, 17 pipeline and 51
  integration tests. Full CI passes in `_artifact/stdlib-float-exponent-ci.log`,
  including all 1006 compiler tests, 171 driver tests, helper race checks,
  bootstrap fixed point and packaging/smoke checks.
  No public API or builtin ABI changes.

### S1 rational float exponent preflight — validated

- Added conservative bit-length bounds before rational conversion constructs
  shifted arbitrary-precision integers. A nonzero numerator with bit-length
  difference d has a binary exponent of d or d-1 before the requested shift;
  only definite overflow/underflow bypasses exact arithmetic. Zero, signed zero,
  half-subnormal ties and overflow-rounding boundaries keep their established
  behavior. No builtin ABI, native algorithm backend or language syntax changes.
- Added hexadecimal exponents up to 9999999, signed zero/nonzero significands,
  invalid suffixes, and neighbors of both widths' underflow/overflow thresholds.
  Toolchain build, focused runtime tests and golden generation pass (29 formatter,
  17 pipeline and 51 integration tests). Full CI passes in
  `_artifact/stdlib-float-shift-ci.log`, including all 1006 compiler and 171 driver
  tests, helper race checks, bootstrap fixed point and packaging.
- This addresses exponent-proportional temporary allocation, not mantissa scan
  or storage limits. The fixed exponent parsing threshold remains a separate
  correctness audit item before full strconv acceptance; counts remain unchanged.

### S1 numeric long-input and transactional-output audit — validated

- Extended decimal/hex parsing tests to 64-, 768- and 2048-digit mantissas and
  exponents, including exponent cancellation, signed underflow, overflow and
  malformed suffixes. Found an oracle limitation: Go 1.26 ParseFloat's fixed
  decimal buffer returns zero for a decimal power canceled by e-2048, whereas
  the exact result and GoML result are one. Long finite inputs now use test-only
  math/big.Rat conversion plus explicit one assertions, not the lossy oracle.
  Huge exponents and malformed syntax retain the ordinary strconv reference.
- Added every insufficient/exact/extra destination capacity for selected unsigned
  values in every radix, including subview offsets, untouched prefix/suffix and
  error payloads. Standard num now participates in the source-only/no-native-
  algorithm check alongside text. Production algorithms are unchanged.
- Focused validation and all 1006 compiler tests pass after correcting the
  reference; logs are `_artifact/stdlib-num-acceptance-focused.log` and
  `_artifact/stdlib-num-acceptance-compiler.log`.
  The initial compiler run passed 1005 tests and failed the numeric fixture on
  that same oracle discrepancy, not a production regression.
- The full strconv row remains unaccepted. Source audit also found the existing
  scalar parser's fixed 1000000 exponent-accumulation threshold; it needs review
  against mantissas long enough to cancel that threshold before final acceptance.
  Increasing that threshold alone is insufficient: rational-to-float conversion
  currently applies the binary shift before classifying overflow/underflow, so
  exponent handling must also avoid allocating enormous shifted temporaries.
  No broader input-size/resource guarantee is inferred from the new test sizes.

### S1 scalar numeric conversion differential audit — validated

- Added independent test-only Go references for integer-to-float bit patterns
  and rounded float-to-integer results. The reference rejects nonfinite values
  and checks exact power-of-two bounds before invoking Go integer conversions,
  avoiding reliance on platform-specific overflow behavior.
- Covered every integer source/destination width, all five explicit rounding
  modes, both float widths, signed zero, subnormals, NaN/infinity, signed/unsigned
  endpoints, halfway values and deterministic random samples. Float samples span
  every IEEE exponent encoding with selected fraction/sign combinations, not
  every representable float. Integer samples include powers of two and adjacent
  values, plus u64-to-f32 double-rounding-sensitive halfway neighbors.
- Focused external runtime tests pass. Existing pure GoML conversion algorithms
  needed no changes; native references remain confined to test fixtures. Golden
  generation passes all 29 formatter, 17 pipeline and 51 integration tests.
  The preceding parse-error change passed full CI; this subsequent batch changes
  only tests and this implementation record.
- This closes the scalar conversion coverage gap, not the entire strconv
  acceptance audit. Counts remain 19 accepted, 5 started and 50 planned;
  remaining A/B implementation and user review/release gates stay active.

### S1 floating parse diagnostics and differential audit — validated

- Found and fixed information loss in standard numeric wrappers: the existing
  pure scalar parser reports overflow with a failed status and infinite value,
  but std num previously flattened it to invalid syntax. Structured errors now
  retain `is_range()` and `value out of range`, while preserving existing error
  kind/context/constructor contracts. No builtin/runtime or parser algorithm
  change; explicit infinity and underflow remain successful values.
- Added native test-only Go references comparing validity, ErrRange and exact
  result bits for both widths (NaN classification, not payload identity).
  Cases include decimal/hex boundaries, halfway/subnormal rounding, signed zero,
  extreme exponents, separators, malformed text, generated grammar combinations
  and random bit-pattern strings. Custom error construction/equality and public
  navigation are tested. Toolchain build, focused runtime differential tests and
  golden generation pass. Full CI passes with 1006 compiler tests, 171 driver
  tests, helper race tests, bootstrap fixed-point and packaging checks.
- This improves the existing strconv capability; full conversion/acceptance audit
  remains outstanding and all A/B implementation/review gates remain active.

### S1 shortest decimal floating formatting — validated

- Added f32/f64 shortest conversion with explicit fixed/scientific/general
  notation and case choices. Increasing significant-digit search checks exact
  nearest and neighboring decimal candidates through the existing pure GoML
  scalar parser, retaining bit identity, negative zero and rounding ties. Existing
  compiler-owned formatting/parser fallbacks are not removed or modified.
- General shortest selection uses the Go -4/6 exponent thresholds; documented
  that shortest significant digits do not imply minimum total characters for
  every notation. Output limits are checked before final rendering; unexpected
  candidate exhaustion is a recoverable RoundtripFailure, never silent fallback.
- Added exact Go output comparisons for all five notations, f32/f64 exponent-
  spanning powers and neighbors, random bit patterns, subnormals, finite extremes,
  nonfinite values and notation thresholds, with exact/insufficient bounds.
  An explicit statement separator fixed a loop-followed-by-tuple ambiguity.
  Toolchain rebuild and focused runtime tests pass.
- Golden generation passes 29 formatter, 17 pipeline and 51 integration tests.
  Complete `just ci` passes 1006 compiler tests, 171 driver tests, metadata helper
  race tests, Unicode generator checks, bootstrap fixed point and packaging smoke
  tests. A final test-only expansion adds explicit signed f32 zero/subnormal/normal
  boundary/max-finite/infinity/NaN cases; the focused external consumer passes again
  after that expansion. Production sources are unchanged from the full-CI run.
  Formatting and whitespace checks pass. Evidence:
  `_artifact/stdlib-float-shortest-build-retry.log`,
  `_artifact/stdlib-float-shortest-{focused,golden,ci,f32-edges}.log`.
- Remaining numeric parsing/conversion and overall strconv acceptance still
  require audit. No inventory counts or review/publication gates are changed.

### S1 binary and hexadecimal floating formatting — validated

- Added pure f32/f64 binary-significand and normalized hexadecimal formatting.
  Hex supports exact trimmed fractions or nonnegative fractional precision,
  ties-to-even rounding, carry normalization and uppercase output. Output lengths
  are checked before construction, including arbitrarily large requested padding.
- Binary output preserves IEEE subnormal scale (also for signed zero); hex
  normalizes subnormals and gives zero exponent zero. Documentation distinguishes
  the decimal integer significand in binary mode from literal binary digits.
- Added Go differential tests across every f32/f64 exponent encoding, both signs,
  representative mantissas, multiple precisions/cases, exact/insufficient limits,
  and each hexadecimal rounding cutoff's midpoint/neighbors. Public navigation
  and examples are updated. Build and runtime validation pass, including 276,558
  reference comparisons for power-format output (plus exact/insufficient bounds).
- Golden generation passes 29 formatter, 17 pipeline and 51 integration tests;
  all 1006 compiler tests subsequently pass. Formatting and whitespace checks
  pass. Logs: `_artifact/stdlib-float-power-{build,focused,golden,compiler}.log`.
  No package/dependency or bootstrap mechanism changed; the previous decimal
  format full CI remains the broad baseline, with full CI required before release.
- Explicit shortest decimal formatting and the remaining numeric conversion
  audit are still outstanding; no capability count or release gate is changed.

### S1 scientific and general floating formatting — validated

- Added f32/f64 scientific and general decimal formatting, sharing exact IEEE
  conversion, rounding, nonfinite handling and checked rendering with fixed
  formatting. Scientific output has explicit exponent sign and at least two digits;
  general output trims zeros and selects notation after rounding at the -4 and
  precision thresholds. Both support uppercase exponent markers.
- General precision zero means one significant digit; huge positive precision
  does not pad to that length. Negative precision remains an explicit error;
  shortest/default, hex and binary formatting remain separately tracked work.
- Expanded all prior finite/edge/nonfinite/precision differential cases to both
  letter cases of Go e/E and g/G; added notation-transition/carry cases and maximum
  integer precision checks. Public docs/examples and navigation are updated.
  Toolchain build and focused runtime differential validation pass.
- Golden generation passes 29 formatter, 17 pipeline and 51 integration tests.
  Complete `just ci` passes 1006 compiler tests, 171 driver tests, helper race
  tests, Unicode generation checks, bootstrap fixed point and packaging smoke
  tests. Formatting and whitespace checks pass. Logs:
  `_artifact/stdlib-float-decimal-{build,focused,golden,ci}.log`.
  The full strconv row remains in progress: explicit shortest/default, hex and
  binary format APIs and remaining conversion audit are not complete. Counts
  remain 19 accepted, 5 started and 50 planned-only; release remains review-gated.

### S1 fixed-point floating formatting — validated

- Added pure standard f32/f64 fixed formatting from IEEE bits through exact
  decimal digit arithmetic, with ties-to-even rounding, retained negative zero,
  nonfinite tokens, nonnegative decimal places and preflight output bounds.
  No native formatting/float arithmetic backend or compiler-owned consumer change.
- Registered the public source, documented format/limit/error behavior and added
  navigation plus Go strconv differential tests for rounding ties and neighbors,
  decimal carries, signed zeros, normal/subnormal boundaries, exponent-spanning
  samples, maximum finite values, NaNs/infinities and precision through 1074 places.
- Initial lowering exposed an ambiguous if-followed-by-tuple expression; explicit
  statement separation fixed it. A test sign-mask literal also needed an explicit
  unsigned construction. Toolchain rebuild and focused runtime differential tests
  pass. Golden generation passes 29 formatter, 17 pipeline and 51 integration tests.
- Complete `just ci` passes 1006 compiler tests, 171 driver tests, metadata helper
  race tests, generator checks, bootstrap fixed point and packaging smoke tests.
  Formatting and whitespace checks pass. Evidence:
  `_artifact/stdlib-float-fixed-build-retry.log`,
  `_artifact/stdlib-float-fixed-focused-final.log`,
  `_artifact/stdlib-float-fixed-golden.log` and `_artifact/stdlib-float-fixed-ci.log`.
  Scientific/general/hex/binary formatting and the remaining conversion audit
  are still outstanding; this does not accept the full strconv row. Counts remain
  19 accepted, 5 started and 50 planned-only; full implementation/review gates apply.

### S1 boolean conversion — validated

- Added pure `num::parse_bool_structured` with the exact Go strconv boolean
  spellings, recoverable `ParseBoolError` and no default-false fallback. Error
  input is inspectable explicitly, but diagnostic formatting does not echo it.
  Existing boolean ToString supplies canonical formatting without a duplicate API.
- Added Go differential coverage for all ASCII single characters, every casing
  of true/false, prefixed/suffixed whitespace and punctuation, Unicode lookalikes,
  canonical roundtrips and retained input. Public docs and navigation are updated.
  Toolchain build and focused runtime validation pass. Golden generation passes
  29 formatter, 17 pipeline and 51 integration tests; all 1006 compiler tests
  subsequently pass, including public navigation and the external consumer.
  Formatting and whitespace checks pass. No source catalog/dependency, runtime
  contract or bootstrap mechanism changed; the preceding quoted-literal full CI
  remains the broad baseline, with full CI required again before release.
  Evidence: `_artifact/stdlib-bool-{build,focused,golden,compiler}.log`.
- This continues the existing strconv capability; floating-point conversion
  coverage/API work and final inventory acceptance remain outstanding. Counts
  remain 19 accepted, 5 started and 50 planned-only; no release is authorized.

### S1 quoted literal decoding — validated

- Added pure `text::unquote_bytes` and strict-UTF-8 `unquote` alongside quotation.
  Complete double/single/backquoted forms support simple, byte, octal and Unicode
  escapes; raw CR removal and Go Unquote's empty-single-quote behavior are explicit.
  Malformed escapes/scalars, trailing input and excessive output are recoverable.
- `UnquoteError` distinguishes input syntax/output positions from invalid decoded
  UTF-8 positions, without echoing input. Output is checked before append; errors
  expose no partial bytes. Input scanning is separately caller-bounded.
- Added Go differential parsing cases, 5,184 delimiter/body combinations,
  byte-versus-scalar distinctions, quote roundtrips across the existing Unicode
  sample and exact/insufficient limits. Public docs/examples and definition
  navigation are updated. Build and focused external tests pass.
- `just update-golden` passes 29 formatter, 17 pipeline and 51 integration tests.
  Complete `just ci` passes all 1006 compiler tests and 171 driver tests, metadata
  helper race tests, Unicode generator checks, bootstrap fixed point and packaging
  smoke tests. This also revalidates quotation and its corrected source-count
  assertion from the previous batch. Formatting and whitespace checks pass.
  Logs: `_artifact/stdlib-text-unquote-{build,focused,golden,ci}.log`.
- The full numeric/quoting row still requires its remaining API audit. Ledger
  counts remain 19 accepted, 5 started and 50 planned-only; no release is authorized.

### S1 quoted text/scalar encoding — validated

- Added pure `text::quote`, `quote_ascii`, `quote_char` and `quote_char_ascii`
  using the existing pinned printable-scalar classification and checked output
  sizing. Delimiters count toward the limit; ASCII controls and nonprintable
  scalars use explicit Go-style escapes. Valid scalar input excludes invalid
  UTF-8 byte repair and surrogate handling from this API.
- This is the text-owned portion of the existing strconv row, not a new std
  package or a relocation of domain encoders. Documentation distinguishes the
  literal format from JSON, shell/HTML escaping and GoML source syntax.
- Registered public source and definition navigation; added Go strconv
  differential tests for all Latin-1 scalars, a deterministic full-range sample,
  Unicode whitespace, quote/backslash/control cases, supplementary characters
  and exact/insufficient output limits. Build and focused external tests pass.
- `just update-golden` passes 29 formatter, 17 pipeline and 51 integration tests.
  Full CI found one stale source-count assertion (text now has four files);
  updating that assertion and rerunning the compiler suite passes all 1006 tests.
  The original CI's other branches passed: 171 driver tests, metadata helper race
  tests, generator checks, bootstrap fixed point and packaging smoke tests. The
  complete CI command was not rerun after this test-only correction; a final
  green full CI run remains a release gate. Formatting and whitespace checks pass.
  Evidence: `_artifact/stdlib-text-quote-{build,focused,golden,ci,compiler-recheck}.log`.
- Unquoting and the remaining numeric API audit are still outstanding. This
  does not accept strconv or strings; counts remain 19 accepted, 5 started and
  50 planned-only, with full implementation and review required before release.

### S1 bytes acceptance audit — validated

- Inspected all five public byte-package source files, the Unicode-only package
  dependency, source catalog and source/link isolation tests. Production search,
  transforms and Unicode-aware operations contain no Go FFI or native algorithm
  delegation. Existing byte-vector/scalar/allocation boundaries remain unchanged.
- Added 784 overlapping source/part/separator combinations for join and replace,
  including failure non-mutation, single-repeat copies and independence after
  mutating every source byte. Expanded split/replace differential counts to both
  machine integer extremes. Focused external tests pass, and definition navigation
  now checks reusable reverse search and malformed-run replacement directly.
- The prior residuals are implemented: lazy split/line/field iteration, bounded
  construction, shared Unicode scalar operations and existing-Cursor buffered I/O
  integration. Frozen-byte and builder-copy contracts also have external coverage
  in project089. No duplicate byte-reader/buffer type or io dependency is added to
  bytes; stream contracts remain in io. Full compiler verification passes all
  1006 tests with unchanged execution snapshots. Formatting and whitespace checks
  pass. The mapped bytes row is now accepted; the ledger is 19 accepted, 5 started
  and 50 planned-only. Full implementation, user review and release remain pending.

### S1/S2 bounded byte-buffer I/O integration — validated

- Extended the existing standard `io::Cursor` rather than introducing a second
  byte reader/buffer: `with_limit` validates initial data before copying and
  bounds total storage length; `limit` exposes the fixed bound. Existing `new`
  retains its maximum-isize bound and all established seek/aliasing behavior.
- Sequential and positional writes reject excessive output before snapshotting
  or mutation. Rejections preserve storage and shared position, report zero
  positional progress and do not prevent subsequent within-limit overwrites.
  Limits are not cumulative-write budgets or allocation-failure recovery.
- Added an overlap matrix covering sequential/positional writes at five bounds,
  copied input and shared handles, zero capacity, negative/insufficient initial
  limits, and buffered integration: repaired bytes survive source mutation,
  repeated failed flushes preserve pending output, a seek enables successful
  retry, and a one-byte buffered reader retrieves the exact result.
- Public documentation, example and definition navigation are updated. Toolchain
  rebuild and focused consumer pass. An initial concurrent golden invocation
  stopped at executable replacement with `Text file busy`; after the consumer
  exited, the serial retry passed 29 formatter, 17 pipeline and 51 integration
  tests. No generated snapshots were edited by hand.
- Full `just ci` passes, covering 1006 compiler tests, 171 driver tests, metadata
  helper race tests, Unicode generator checks, bootstrap fixed point and toolchain
  packaging. Evidence is in `_artifact/stdlib-cursor-limit-build.log`,
  `_artifact/stdlib-cursor-limit-focused.log`,
  `_artifact/stdlib-cursor-limit-golden-retry.log` and
  `_artifact/stdlib-cursor-limit-ci.log`. Formatting and whitespace checks pass.
  This advances bytes/io integration without declaring either inventory row
  complete; the ledger remains 18 accepted, 6 started and 50 planned-only.

### S1 malformed UTF-8 byte-run replacement — validated

- Added pure `bytes::replace_invalid_utf8` over the shared boundary decoder.
  Each maximal malformed-byte run receives one arbitrary-byte replacement;
  valid encodings, including U+FFFD, remain unchanged. Empty replacement deletes
  malformed runs, and replacement bytes are not decoded recursively.
- Preflight checked output length before allocating, preserve independent output
  storage even for already-valid input, and expose recoverable negative-limit,
  overflow and output-limit failures. No new dependency or native primitive.
- Documented the distinction from scalar mapping, and added public completion
  coverage plus external Go `bytes.ToValidUTF8` differential tests for all one-
  and two-byte inputs with four replacement policies, multibyte boundaries,
  exact/insufficient limits and alias isolation. Toolchain rebuild and the
  external consumer pass. Initial compiler validation found five stale pipeline
  snapshot groups after adding the public source function; `just update-golden`
  passed 29 formatter, 17 pipeline and 51 integration tests. The subsequent full
  compiler run passes all 1006 tests. Formatting and diff whitespace checks pass.
  Logs: `_artifact/stdlib-bytes-repair-{build,focused,golden,compiler-recheck}.log`.
  This change adds no package/dependency or bootstrap mechanism; full CI remains
  required at the final release gate.
- This fills a bytes API gap, not acceptance of the whole bytes capability.
  The ledger remains 18 accepted, 6 started and 50 planned-only; I/O integration
  and the other mapped A/B requirements remain part of the full objective.

### S1 Unicode reproducibility acceptance audit — validated

- Reviewed pinned Unicode 15.0.0 property/range/case sources, pure lookup code,
  custom table validation and exhaustive external scalar comparisons. Standard
  Unicode remains dependency-free; full folding and simple case mappings retain
  their distinct public semantics.
- Closed a reproducibility gap: the full casefold generator now supports a
  non-mutating --check mode with canonical formatting and GOMLFMT selection,
  alongside explicit --write and its compatible historical default write mode.
  The script CI gate now checks full folds as well as the seven generated tables.
- Both local generation checks passed against the existing sources, with no
  regenerated data changes. Python Unicode 15.0.0 remains mandatory and version
  mismatches fail clearly rather than silently changing semantics.
- Full CI passed in `_artifact/stdlib-unicode-acceptance-ci.log`; generation
  evidence is in `_artifact/stdlib-unicode-{tables,casefold}-check.log`.
  The accepted-row record above states the final scope and evidence.

### S1 UTF-8 multibyte boundary audit — validated

- Added 57,600 combinations of boundary lead bytes, every second byte and
  continuation/control/extreme trailing bytes. Tests cover both standalone data
  and a valid multiscalar prefix, strict incremental decoding and complete
  reverse walks against Go DecodeLastRune.
- Extended differential checks to independently derive the first invalid offset
  and incomplete-versus-invalid status using Go DecodeRune/FullRune, rather than
  relying only on agreement between the GoML validator and incremental decoder.
- Added failed-finish stickiness and shared-handle reset checks after an incomplete
  four-byte sequence. Existing exhaustive one/two-byte and sampled scalar tests
  remain active. Production code is unchanged.
- Focused execution and all 1,006 compiler tests passed in
  `_artifact/stdlib-utf8-boundaries-{focused,compiler}.log`. The accepted-row
  record above states the final scope and evidence.

### S1 Base32 transactional decoder audit — validated

- Reviewed incremental decode commit points, pending quantum handling, finish
  recovery, strict padding validation and the stream adapter's sticky failures.
- Added every two-part split of encoded payload lengths zero through sixteen
  across standard/hex padded/raw modes. Each case appends an invalid trailing
  symbol after valid input, checks the absolute error offset and unchanged
  accepted counters, retries the valid tail, finishes, resets and decodes again.
- Mutating prior returned buffers does not affect later output or decoder state;
  successful prefixes remain intact even after a failed update with local output.
  Added definition navigation for the incremental update method.
- Focused execution and all 1,006 compiler tests passed in
  `_artifact/stdlib-base32-transaction-{focused,compiler}.log`. The accepted-row
  record above states the final scope and evidence.

### S1 PEM header validation audit — validated

- Found that decode's Unicode trim could remove non-ASCII header bytes before
  Block validation, contradicting the documented printable-ASCII contract.
  Restricted constructor/decoder header normalization to ASCII space and tab.
- Added regression cases for nonbreaking/fullwidth/em whitespace around names
  and values, accepted ASCII padding, all negative limit positions, malformed
  Base64 padding/canonical bits and the exact empty-envelope byte budget.
- Explicitly compared fail-fast malformed-candidate handling with Go's searching
  decoder when a later valid block exists; the documented difference is retained.
- Incremental toolchain build, focused execution and all 1,006 compiler tests
  passed; evidence is recorded under `_artifact/stdlib-pem-audit-*.log`.
  The accepted-row record above states the final scope and evidence.

### S1 errors interoperability acceptance audit — validated

- Reviewed typed Report wrapping, aggregation, matching/projection, immutable
  tree structure and iterative rendering against the mapped errors capability.
  Existing Error/ErrorKind/Details protocols and runtime codes stay unchanged.
- Added a heterogeneous domain enum containing standard io::Error and validation
  failures, with map_err/? propagation, generic Error-protocol formatting and
  typed projection retaining the original kind/OS code. No reflection is needed.
- Added duplicate shared-subtree visitation, exact depth-first callback order,
  early-stop checks and mutable leaf sharing with isolated cause vectors.
- Documented interoperability and resource/concurrency responsibilities, and
  added context/find_map definition-navigation checks.
- Focused execution passed in `_artifact/stdlib-errors-acceptance-focused.log`.
  All 1,006 compiler tests passed in `_artifact/stdlib-errors-acceptance-compiler.log`.
  The accepted-row record above states the final evidence and contract boundaries.

### S1 slices acceptance audit — validated

- Audited all view comparisons, membership, bounded arithmetic, overlapping
  vector edits, stable deletion, reserve, ordering/search/extrema and iterator
  composition helpers against the mapped collection contract and existing tests.
- Added 2,892 self-overlapping replacement cases spanning all source and target
  ranges for lengths zero through seven, with independent expected sequences and
  shared Vec alias checks. Added complete small chunk partitions, live view
  mutations, shared versus independent iterator cursors, comparator short-circuit
  counts and ordered stable-deletion visitation.
- Expanded public documentation and navigation for checked replacement and chunk
  creation. Production source/catalog remain unchanged and continue to use only
  ordinary GoML plus standard comparison contracts.
- Focused execution and all 1,006 compiler tests passed in
  `_artifact/stdlib-slices-acceptance-{focused,compiler}.log`. The accepted-row
  record above states the final contract and evidence.

### S1 maps acceptance audit — validated

- Reviewed clone/copy, membership-aware equality, snapshot iteration/deletion,
  generic insertion and collection against the mapped HashMap helper contract.
  Production remains pure GoML over existing HashMap storage and hashing.
- Added exact comparator call counts for equality, short-circuit failure and
  disjoint keys; deletion callbacks removing future snapshot keys; nested Vec
  sharing with independent map entries; and iterator side effects followed by
  yielded-pair overwrite, exhaustion and self-snapshot insertion.
- Documented iterator resource responsibility, callback order and mutation
  boundaries without implying synchronization, rollback or deep cloning.
  Expanded completion/navigation checks for map predicates and iterators.
- Focused execution passed in `_artifact/stdlib-maps-acceptance-focused.log`;
  all 1,006 compiler tests passed in
  `_artifact/stdlib-maps-acceptance-compiler.log`. The accepted-row record above
  states the final evidence and contract boundaries.

### S1 logical path acceptance audit — validated

- Reviewed all lexical functions, pattern compilation, immutable matching state,
  budget accounting and documented boundary semantics against the existing scope.
- Added 11,664 compositional pattern/name checks with explicit equivalent regular
  expressions for each finite source atom and an anchored Go regexp test oracle.
  This checks full scalar matching independently of Go path.Match's byte retries
  and incomplete choice exploration around slash-matching classes. Existing
  direct Go path differential tests remain; production has no regex dependency.
- Direct comparison exposed Go matching *?? against one Unicode scalar by
  retrying inside its UTF-8 encoding. Explicit tests assert the Go and GoML
  results for this case and replacement-character classes. A subsequent ASCII
  reduction exposed Go rejecting *[a/]* against a/ despite a valid match. Explicit
  tests retain both outcomes; the guide explains these semantic differences.
- Added malformed suffixes after matching/nonmatching prefixes, exact range and
  collapsed-star work budgets, and reuse through aliases after budget failures.
  The Go oracle now reports validity instead of discarding its error result.
- Focused execution and all 1,006 compiler tests passed in
  `_artifact/stdlib-path-acceptance-{focused,compiler}.log`. Expanded Pattern
  completion and navigation checks passed as part of this compiler run.
  The accepted-row record above states the final scope and evidence.

### S1 URL final acceptance audit — validated

- Audited the mapped value-parsing, component/query codec, authority, reference
  resolution and ASCII serialization contracts against their public source,
  guide, resource catalog, driver selection and dependency tests. Production
  imports are confined to bytes; no HTTP/runtime policy was moved into std.
- Added authority decode ownership checks for usernames, passwords, binary
  hostnames and zones; independent results remain unchanged after buffer mutation.
  Added exact authority-relative syntax offsets, empty-password redaction bounds
  and absent-password negative-limit coverage.
- Existing evidence includes RFC resolution examples, deterministic Go codec,
  query, authority and IPv6 differential tests, explicit intended differences,
  resource/link purity, completion/navigation and independent HTTP consumers.
  Request's absolute decomposition/query integration is validated; its relative
  compatibility policy intentionally remains ecosystem-owned.
- Focused execution and final full CI passed in
  `_artifact/stdlib-url-acceptance-{focused,ci}.log`. The accepted-row record
  above summarizes the complete evidence and remaining intentional boundaries.

### E1 request integration with standard URL references — validated

- Replaced absolute URL scheme/authority/path/query decomposition and duplicate
  path-percent validation with std::net::url::Reference. This also removes the
  old host scan's repeated full byte-vector copies.
- Kept request's endpoint policy, path escape spelling, default/empty ports,
  raw query contents, fragment dropping and relative-join compatibility logic.
  No standard Authority or ASCII-output policy is silently imposed on existing
  HTTP clients. General RFC resolution remains independently available in std.
- Added library cases for uppercase authorities, empty components, zero-padded
  ports, IPv6, Unicode, raw malformed query escapes, ignored fragments, invalid
  paths and relative joins. Registry verification passed: 51 library tests,
  two independent consumer tests, cached build and execution, plus library,
  consumer and native race checks. Evidence is recorded in
  `_artifact/stdlib-url-request-reference.log`. The shared hex-digit helper was
  retained for transport, Cookie and HPACK consumers after the initial build
  caught those remaining uses. Final standard URL acceptance audit remains open.

### E1 request integration with standard URL query codecs — validated

- Replaced request's duplicate form/query decoder, encoder and sorting with
  std::net::url::parse_query_bounded and encode_query. Request retains UTF-8 text
  validation over the standard byte-preserving results, plus its existing HTTP
  scheme, credential, path, authority and redirect policies.
- Internal form encoding now returns Result; Url::query propagates errors and
  RequestBuilder::form uses the existing deferred Builder error channel. Public
  signatures and stable duplicate-value ordering are unchanged. Decode budgets
  derive from raw input size; encode uses the representable string-size bound,
  not a new application-level size policy.
- Added library compatibility tests for byte spelling, ordering, empty fields,
  malformed escapes, invalid UTF-8, NUL and semicolon handling, plus an independent
  HTTP consumer assertion. Registry-based verification passed, including library
  and independent consumer tests, cached build, execution, library/consumer race
  and native race checks, in `_artifact/stdlib-url-request-integration.log`.
- The standard package has no ecosystem dependency. This development integration
  requires the new toolchain APIs; no ecosystem version is published or overwritten.
  General URL decomposition/resolution integration and final acceptance remain open.

### S1 URL ASCII serialization — implemented and validated

- Added Reference::to_ascii with explicit authority validation, uppercase percent
  spelling, UTF-8 byte escaping and a complete encoded-length preflight. Existing
  component markers and reserved escapes are retained; ToString stays lossless.
- Raw query escapes are checked during serialization without applying form or
  HTTP policy. Canonical escaping deliberately does not claim URL equivalence,
  DNS/IDNA normalization, safe logging or endpoint validation.
- Added exact output and budget cases, parse/serialize idempotence, decoded-path
  preservation, printable-ASCII coverage and method navigation checks. The first
  printable-byte fixture accidentally treated // as a path; its prefix was fixed
  while preserving the independent explicit-empty-authority case.
- Incremental toolchain build, focused consumer execution and all 1,006 compiler
  tests passed in `_artifact/stdlib-url-ascii-{build,focused,compiler}.log`.
  No package/resource ownership changed in this batch. HTTP consumer integration
  and the overall URL acceptance audit remain pending.

### S1 URL authority differential and navigation audit — validated

- Added a test-only Go netip oracle for compressed/uncompressed IPv6, dotted
  tails, invalid octets, group widths and inserted delimiters/non-hex characters.
  Production authority parsing still has no Go delegation or network dependency.
- Compared hostname, raw port and password-redacted reference output for 180
  combinations against Go net/url. Separately asserted deliberate differences:
  IPvFuture and percent-encoded ASCII names are accepted here, and multiple raw
  user-info separators are rejected rather than split at the final separator.
- Added completion checks for Authority, UserInfo and HostKind plus definition
  navigation checks for authority parsing, Reference::authority and redacted.
- Focused execution and all 1,006 compiler tests passed, recorded in
  `_artifact/stdlib-url-authority-audit-{focused,compiler}.log`. This test-only
  follow-up does not replace the preceding production change's full CI evidence.
  The overall URL row remains in progress pending canonical wire output and
  HTTP consumer integration.

### S1 URL authorities — implemented, acceptance in progress

- Added an independently validated Authority value with immutable raw user, host,
  port and zone components. IPv4/IPv6/IPvFuture classification is lexical; the
  package still depends only on standard bytes, without networking backends.
- Distinguished decimal port spelling from checked u16 conversion and absent
  fields from explicit empty fields. Decode operations retain arbitrary bytes.
  Reference authority validation is explicit, with authority-relative diagnostics.
- Added bounded password redaction without Debug implementations for credential-
  bearing values. Other components are deliberately not sanitized or rewritten.
- Registered the source in toolchain resources and added external tests for
  literal grammar, invalid separators, zone escapes, limits, port overflow,
  lossless recomposition and reference integration. The first build and focused
  consumer passed in `_artifact/stdlib-url-authority-{build,focused}.log`.
- Full `just ci` passed in `_artifact/stdlib-url-authority-ci.log`, including
  1,006 compiler tests, 171 driver tests, fixed-point and packaging checks.
  Canonical wire formatting, expanded differential/navigation coverage and
  HTTP consumer integration remain pending;
  this does not mark the net/url capability accepted or authorize a release.

### S1 raw URL references and resolution — implemented and validated

- Added parse_query_bounded with an explicit raw-byte budget checked before copying
  or scanning, including separator-only input. Retained the three-argument query
  parser and documented the distinct InputLimit error without changing its contract.
- Added an immutable raw Reference value with checked input length, scheme/control/
  path-fragment escape diagnostics, exact component access and lossless recomposition.
  Absent and explicitly empty authority/query/fragment markers remain distinct;
  byte-decoded paths/fragments preserve invalid UTF-8 through owned buffers.
- Implemented generic RFC 3986 component inheritance, path merging and literal
  dot-segment removal with a checked final output budget. Encoded dots/slashes and
  query/fragment contents are not normalized. A /. guard retains the no-authority
  interpretation when normalized paths would otherwise start with //.
- Tests cover section 5 normal/abnormal reference-resolution examples, Go differential
  hierarchical paths, duplicate separators, escaped delimiters, component round trips,
  all ASCII controls, Unicode byte offsets and input/output limits. Explicit empty-
  authority and fragment semantics use independent expected results, since Go
  collapses/inherits those cases. Two root-level ..// cases preserve repeated
  separators according to the segment rules; their Go outputs and distinct expected
  outputs are both asserted instead of broadly suppressing differential failures.
- Resource registration, public docs and completion/navigation include the reference
  source and bounded query API. Build and focused tests passed, followed by snapshot
  generation (29 formatter, 17 pipeline and 51 integration tests) and full CI
  (1,006 compiler, 171 driver tests, Go helper race tests, bootstrap fixed point and
  packaging checks). Evidence: `_artifact/stdlib-url-reference-build.log`,
  `_artifact/stdlib-url-reference-focused.log`,
  `_artifact/stdlib-url-reference-golden.log` and
  `_artifact/stdlib-url-reference-ci.log`.
- Authority/host/port/user-info validation, canonical wire escaping/redaction and
  HTTP consumer integration remain required before net/url acceptance. Counts remain
  9 accepted, 15 started awaiting acceptance and 50 planned. No release was made.

### S1 URL component/query foundation — implemented and validated

- Started net/url in its planned standard owner, std::net::url. Pure component
  codecs distinguish path segments from form-query components, preserve arbitrary
  decoded bytes (including invalid UTF-8/NUL), accept either hex case and enforce
  encoded/decoded byte limits with typed errors and byte offsets.
- Query parsing retains duplicate and empty name/value fields, enforces field and
  aggregate decoded-byte limits, rejects raw semicolons and returns no partial
  collection on error. Encoding sorts a copied pair sequence by raw name bytes,
  preserves equal-name value order, preflights output size and leaves input intact.
- Added source/resource catalog registration, dependency selection/link closure,
  source purity, namespace completion/navigation and public documentation. The
  package depends on bytes, not its net namespace parent or any socket/TLS/HTTP
  backend. Existing request::Url restrictions and ecosystem manifests are unchanged.
- External Go net/url differential tests cover all 256 input bytes, seeded binary
  inputs, invalid escapes, Unicode, duplicate/binary query fields, exact/short limits,
  offsets, field order and independent decoded storage. Expanded random binary
  query cases passed alongside snapshot generation (29 formatter, 17 pipeline and
  51 integration tests) and full CI (1,006 compiler, 171 driver tests, Go helper race
  tests, bootstrap fixed point and packaging checks). Evidence:
  `_artifact/stdlib-url-codec-build.log`,
  `_artifact/stdlib-url-codec-focused.log`,
  `_artifact/stdlib-url-codec-golden.log` and
  `_artifact/stdlib-url-codec-ci.log`.
- After CI, corrected only the display label from "byte" to "index" because query
  encoding reports a pair index, added its regression, and rebuilt/retested the
  external consumer successfully in `_artifact/stdlib-url-codec-final-build.log`
  and `_artifact/stdlib-url-codec-final-focused.log`. No API/dependency/catalog
  change followed CI. Documentation explicitly requires caller-owned raw input
  byte budgets in addition to field/decoded limits; ignored separators still
  require input scanning and the current parser copies the input bytes.
- General URL value/authority parsing, raw-component preservation, formatting,
  reference resolution and request-compatible consumer integration remain required.
  No placeholder Url type was added and net/url is not accepted. Counts are now
  9 accepted, 15 started awaiting acceptance and 50 planned.

### S1 sorting and indexed comparison search — implemented and validated

- Added search_by/search_by_ordering for allocation-free virtual sequences,
  returning a checked insertion index and found flag, and equal_range_by_ordering
  for parity with existing bounds APIs. Comparator direction matches element-to-
  target collection searches, explicitly opposite Go sort.Find's callback sign.
- Retained all existing pure GoML stable merge-sort entry points and compiler-owned
  fallbacks. Public documentation now states comparator ordering/mutation rules,
  temporary-space/comparison bounds, alias visibility and nontransactional panic
  behavior. No float total-order policy, reflection or Go algorithm backend added.
- External tests exercise all six sort entry points on sorted, reversed, equal,
  alternating signed extremes, seeded duplicates and organ-pipe inputs around
  merge-width boundaries. They compare record identities with Go stable sorting,
  assert comparison bounds, and observe updates through vector and slice aliases.
- Indexed search tests compare Go sort.Find, verify machine-maximum lengths,
  callback bounds/counts, absent targets and duplicate first matches; new public
  APIs have completion/navigation coverage. Build, focused execution, snapshot
  generation and all 1,006 compiler tests passed; see the acceptance record above
  for evidence. Counts are now 9 accepted, 14 started awaiting acceptance and
  51 planned. Release and user-review gates remain unchanged.

### S1 fixed-width bit-operation acceptance audit — implemented and validated

- Audited both production source files against the mapped fixed-width contract
  and Go 1.26 math/bits function families. Counts, lengths, rotations and reversals
  cover 8/16/32/64 bits; carry/borrow and double-width arithmetic cover 32/64 bits.
  The documented API deliberately uses boolean carry/borrow, checked division
  errors and explicit widths, not native-word aliases or constant-time guarantees.
- Added exhaustive 16-bit count/length/reversal checks and rotation differential
  tests, plus minimum/maximum isize shifts for all rotation widths. Added each
  power-of-two neighbor to arithmetic boundary tests, even and odd divisors,
  high-word overflow thresholds, full-high-word remainder cases and explicit
  zero-divisor precedence at both arithmetic widths.
- Seeded arbitrary dividend/divisor checks compare against Go and reconstruct
  each successful 128-bit dividend from quotient, divisor and remainder using
  independent multiplication and carry propagation. Existing FNV-128 tests remain
  a cross-package consumer of mul64. Added navigation coverage for rem64.
- Focused external tests, snapshot generation and all 1,006 compiler tests passed;
  see the acceptance record above for evidence. No production algorithm, resource
  catalog, dependency or runtime boundary changed. Counts are now 8 accepted,
  15 started awaiting acceptance and 51 planned.

### S1/S2 explicit binary record composition — implemented and validated

- Added bounded Reader::read_with and staged Writer::write_with / Builder::write_with
  callbacks for explicit field, array and nested-record schemas. Success consumes
  the declared width, with unread padding skipped and unwritten padding zeroed.
  Callback errors preserve outer cursor/storage; nested diagnostics retain their
  originating view offsets. Builder writes enforce a total byte limit.
- Added stream::read_with / write_with with per-record byte limits, exact block
  consumption, interrupted/short-I/O handling and existing write_all retry policy.
  Bounds failures after decoding report InvalidData with full consumed progress;
  encoding failures happen before underlying writes. Zero-width reads return
  Some without probing EOF; documentation explicitly warns against EOF loops.
- No reflection, new runtime boundary, implicit layout or automatic serde mapping
  was added. Callbacks cannot transactionally roll back unrelated side effects or
  captured aliases; the documented contract forbids mutation of outer state.
- External tests cover mixed fields and nested arrays against Go encoding/binary,
  both byte orders, NaN payloads, padding, every short buffer/block length, nested
  failures, invalid/overflow-sized widths and limits, cursor aliases, retained
  scratch isolation, zero-width records, invalid I/O counts and no-retry writers.
  Build, focused execution, snapshots and complete CI passed; see the current
  encoding/binary acceptance record for evidence. This closes the explicit
  aggregate composition gap and accepts that capability against its mapped
  contract. Counts are now 7 accepted, 16 started awaiting acceptance and 51 planned.

### S1 binary scalar cursor completion — implemented and validated

- Extended the existing endian Reader, Writer and Builder with signed integer
  and IEEE floating-point methods. They reuse the checked memory codecs, preserve
  exact bit patterns, and advance shared cursors only after successful operations.
- Added bool codecs to memory, cursor, builder and stream APIs: one byte, zero
  false/nonzero true, canonical zero/one writes. Stream adapters inherit the
  existing interruption, count validation, clean EOF and partial-progress rules.
- The external consumer checks both byte orders, signed limits, signed zero,
  infinities and NaN payloads against explicit expected bytes. It exercises all
  shorter buffer lengths, unchanged guards/cursors on failure, shared reader
  positions, every Boolean input byte and interrupted/fragmented Boolean streams.
- No new package, dependency, runtime primitive or compiler-owned consumer was
  introduced. Documentation and namespace completion coverage include the new APIs.
  Build and focused execution passed, followed by snapshot generation (29 formatter,
  17 pipeline and 51 integration tests) and all 1,006 compiler tests. Evidence:
  `_artifact/stdlib-endian-scalars-build.log`,
  `_artifact/stdlib-endian-scalars-focused.log`,
  `_artifact/stdlib-endian-scalars-golden.log` and
  `_artifact/stdlib-endian-scalars-compiler.log`.
  Aggregate schemas remain outstanding, so encoding/binary is not accepted yet.
  Counts remain 6 accepted, 17 started awaiting acceptance and 51 planned.

### S2 shared hashing streams — implemented and validated

- Added std::hash::stream::Reader/Writer over arbitrary Read/Write and Hasher
  implementations. They preserve supplied state, hash only the successfully
  reported nonempty prefix, reject invalid counts before hashing and forward
  flush/errors without reset, close, buffering or synchronization.
- The base hash package remains I/O-independent. Registered the child source,
  hash/io dependency and link closure, driver selection flags, namespace completion,
  method navigation and resource purity tests. No runtime/Go backend was added.
- The external consumer implements its own counting Hasher and composes real
  FNV digests with interrupted/short I/O through io::copy_buffer. Tests cover
  prehashed prefixes, shared aliases, buffer boundaries, zero counts, empty calls,
  malformed counts, flush/error propagation and recovery, and underlying I/O
  that consumes bytes before returning Err. Unreported progress is explicitly
  not hashed; no rollback is promised.
- Build, focused execution, snapshot generation and full CI passed; see the hash
  acceptance record above for commands, coverage and evidence. This extends the
  existing hash/io work, not a new inventory row; acceptance counts are now
  6 accepted, 17 started awaiting acceptance and 51 planned.
  Final user review and approval remain mandatory before any release.

### S1 address hash keys — implemented and validated

- Added existing prelude Hash derivation to IpAddr, ScopedIpAddr, IpPrefix and
  SocketAddr. Their hashes use exactly the stored fields already compared by Eq,
  including enum family, zone spelling, prefix host bits/length and socket scope.
  No new hash backend, package dependency or compiler contract is introduced.
- Added a generic independent consumer comparing HashMap behavior with a linear
  equality reference for all four types: duplicate canonical spellings, byte
  reconstruction, IPv4/mapped separation, zone case/decimal spellings, masked
  versus unmasked prefixes, ports/scopes, insertion/replacement/removal and key
  value-copy behavior. Equal keys must hash equally; no assumption of collision
  freedom or a stable numeric hash is made.
- Build and focused consumer passed. Snapshot generation passed 29 formatter,
  17 pipeline and 51 integration tests. Full compiler regression exposed a
  source-fallback gap: standard packages bypassed derive expansion, while the
  installed toolchain path already supported it. The fallback now reuses the
  existing derive environment with builtin, prelude and dependency interfaces
  before expanded HIR lowering. The isolated interface-hash regression passed,
  followed by full `just ci`: 1006 compiler tests, 171 driver tests, helper race
  checks, bootstrap fixed-point verification and packaging/smoke checks. Evidence
  is in `_artifact/stdlib-ip-hash-{build,focused,golden,compiler,interface,ci}.log`;
  the compiler log retains the initial failure and the CI log records the fix.
- The net/netip row still needs its final acceptance audit; the 74-capability
  migration and publication remain unfinished.

### S1 IP ordering, neighbors and byte values — implemented and validated

- Added standard Ord/PartialOrd implementations for IpAddr and ScopedIpAddr:
  family then unsigned address order, with zone byte order as the final scoped
  tie-breaker. Ordering agrees with equality and composes with standard collection
  sorting. Registered std::cmp in both source and link dependencies for std::net.
- Added next/prev with Option overflow/underflow, preserving address family and
  zone even across mapped-IPv6 boundaries. No invalid-address sentinel, wraparound
  or implicit unmapping is introduced.
- Added optional as4, mapped network-order as16, IPv6-only from16, checked
  four/sixteen-byte from_slice and independently owned to_bytes. Input copying and
  output isolation are covered; non-IP lengths produce recoverable errors.
- Build and the independent Go netip differential consumer passed, including
  carry/borrow boundaries at every IPv6 byte, scoped neighbors, pairwise ordering,
  collection sorting, byte roundtrips and invalid lengths. Added source navigation
  and explicit dependency-catalog regression coverage. Snapshot generation passed
  29 formatter, 17 pipeline and 51 integration tests. Full `just ci` passed all
  1006 compiler tests, 171 driver tests, helper race checks, bootstrap fixed-point
  verification and packaging/smoke checks. Evidence is recorded in
  `_artifact/stdlib-ip-value-{build,focused,golden,ci}.log`.
- This continues the net/netip row; complete capability acceptance still needs
  review, including hash-key integration for address values. The full migration
  remains 23 started and 51 planned-only capabilities, with publication pending.

### S1 scoped IP address values — implemented and validated

- Added ScopedIpAddr as a checked composition of the existing IpAddr and opaque
  zone string, preserving the public V4/V6 enum and SocketAddr layout. Supports
  parsing, checked zone replacement, explicit zone removal, equality, formatting
  and unmapping with zone removal only when a mapped address becomes IPv4.
- Nonempty zones require IPv6; unlike Go WithZone, an IPv4 construction error is
  reported instead of silently discarding the zone. Parsing uses the first percent
  separator and preserves subsequent contents, including Unicode and percent signs.
- Numeric socket conversion accepts absent/u32 zones and explicitly normalizes
  scope zero/leading zeroes. Interface names are not resolved implicitly; named,
  signed and overflowing socket scopes produce recoverable errors. Converting
  existing SocketAddr values also checks the IPv4/nonzero-scope invariant.
- Added public completion coverage and an independent consumer with Go netip
  parsing/unmapping comparisons, zone identity/roundtrips, invalid input and
  numeric scope boundaries. Build and focused tests passed; snapshot generation
  passed 29 formatter, 17 pipeline and 51 integration tests. Initial complete
  compiler regression exposed a missing explicit source-catalog registration for
  classify.gom/scoped.gom: runtime loading succeeded but editor completion failed.
  Both sources are now registered, with actual method-navigation coverage for
  classification and scoped parsing as well as ScopedIpAddr completion. Full
  `just ci` then passed: 1006 compiler tests, 171 driver tests, helper race checks,
  bootstrap fixed-point verification and packaging/smoke checks. Logs are
  `_artifact/stdlib-ip-scoped-{build,focused,golden,compiler,ci}.log`; the compiler
  log retains the initial failure, while the CI log records the successful fix.
- This extends the existing standard net owner without adding a new package or
  runtime boundary. The net/netip row still needs value ordering, neighboring
  address operations and byte conversion contracts plus their acceptance tests.
  Inventory counts remain 23 started and 51 planned-only; no release is complete.

### S1 IP address classification — implemented and validated

- Extended the existing IpAddr with bit_len, is_ipv4_mapped and non-mutating
  unmap, plus unspecified, loopback, private, multicast, link/interface-local
  and global-unicast classification. These are pure GoML value operations with
  no additional package, backend or runtime dependency.
- Classification follows the existing Go netip value contracts: mapped addresses
  use embedded IPv4 semantics except literal unspecified/interface-local tests.
  Global-unicast includes private/documentation ranges and is explicitly not a
  reachability or security-policy decision. Prefix membership remains family-strict.
- Added a Go netip differential consumer covering all 65,536 IPv6 first segments,
  5,632 IPv4/mapped boundary combinations and explicit zero, broadcast, loopback,
  multicast and mapped-zero cases. Focused execution passed; snapshot generation
  passed 29 formatter, 17 pipeline and 51 integration tests. Complete compiler
  regression passed all 1006 tests. Evidence is in
  `_artifact/stdlib-ip-classify-{build,focused,golden,compiler}.log`.
- This extends the existing net/netip row rather than completing it. Zone/value
  APIs and remaining acceptance work are still outstanding; the inventory remains
  23 started and 51 planned-only capabilities, with no release yet.

### FFI validation prerequisite — long-path witness loading

- File-mode metadata validation now combines generated function/type witnesses
  into one virtual Go file with canonical import aliases, preventing the Go
  package driver's argument partitioning from dropping later witnesses. Package
  mode remains separate; neither mode writes witnesses to the caller directory.
- Declaration origins retain per-binding and per-query error isolation and actual
  method metadata, including deduplicated method bindings. Existing-file conflicts
  remain recoverable errors without overwriting user files.
- Added a regression exceeding the 16,383-byte argument threshold with 160 distinct
  functions, method bindings and type queries. Tests also cover invalid-binding
  isolation, blank-import promotion, alias reuse and source paths containing a
  colon. The helper suite and original failing long-path consumer now pass.
- Full `just ci` passed, including 1006 compiler tests, 171 driver tests, helper
  race checks, bootstrap fixed-point verification and packaging/smoke checks.
  Evidence is recorded in `_artifact/ffi-bundle-ci.log`; the original long-path
  consumer passed in `_artifact/ffi-bundle-long-target.log`. This prerequisite
  adds no migration inventory item and does not expand runtime scope.

### S1 IP prefix values — implementation and focused validation

- Extended existing standard net with IpPrefix over IpAddr, validated construction
  and parsing, explicit masking, membership, overlap and single-address detection.
  Host bits are retained until masking; equality compares stored values. IPv4 and
  IPv6 remain separate families, including IPv4-mapped IPv6; zones are rejected.
- Prefix fields are private, and operations use pure GoML value arithmetic with
  no new socket, DNS, backend or package dependency. Existing IpAddr/SocketAddr APIs
  and their presentation stay compatible.
- Build and initial Go netip differential tests passed across every prefix length,
  address families, extreme addresses and parse failures. Expanded unequal-length
  overlap tests passed as well. Snapshot generation passed (29 formatter, 17 pipeline,
  51 integration tests). Initial full compiler regression had 1005 passes and one
  reproducible FFI metadata failure in project093 under its longer copied path.
- Diagnosed the regression to x/tools/go/packages pattern chunking: it splits
  explicit file arguments above 16,383 bytes, then merges command-line-arguments
  packages by ID and drops later witness files. The original long target loses
  encoding/binary witnesses 84/85; a fresh shorter target passes. Single-job retry
  and temporarily removing the old generated Go file reproduce the failure, so
  neither concurrency nor that cached file explains it. The file was restored
  and temporary diagnostic instrumentation removed. The witness-loading fix and
  long-path/many-witness regression described above now pass full CI, closing
  this compiler-validation gap without shortening the failing caller path.
- This starts net/netip: 23 capabilities started, 51 planned-only. Further address
  classification, zone/value APIs and acceptance work remain unfinished, as do
  the complete migration and publication.

### S1 width-constrained integer parsing — implemented and validated

- Added parse_int_bits/parse_uint_bits returning i64/u64 constrained to 1..64
  bits, with zero selecting native 64-bit width. IntWidthError separates invalid
  widths, retained ParseIntError values and successful full-width values that
  exceed a narrower range. No truncation or saturated error value is returned.
- Differential tests cover every width/base, signed/unsigned boundary values,
  base-zero prefixes, separators, malformed input and full-width overflow.
  They exposed an existing base-zero decimal prefix bug accepting `_1`; fixed
  decimal prefix classification without changing explicit prefix behavior.
- Build and focused differential validation passed after that fix. Go 1.26's
  signed one-bit parser can silently saturate overflowing negatives; tests use
  full-width reference parsing plus the mathematical [-1, 0] bound for that width
  instead of cloning the quirk. This intentional difference is documented.
  Snapshot generation passed (29 formatter, 17 pipeline, 51 integration tests),
  followed by all 1006 compiler tests; execution outputs are unchanged. Full CI
  was not rerun for this source-only batch. This continues strconv: 22 started,
  52 planned-only; all remaining capabilities and release work stay in scope.

### S1 integer radix formatting — implemented and validated

- Extended standard num with pure format_int/format_uint and transactional
  write_int/write_uint for i64/u64 in bases 2 through 36. Output is lowercase,
  prefix-free ASCII; signed minimum uses safe magnitude arithmetic. Invalid radix
  and insufficient destination capacity are recoverable FormatIntError variants.
- No new package, dependency, native algorithm or compiler intrinsic. Temporary
  digit storage is bounded by scalar width, not claimed allocation-free. Existing
  numeric parsing and float behavior remain unchanged.
- Build and external consumer passed: 9,590 Go strconv comparisons across all
  supported bases, signed/unsigned extremes, deterministic bit-pattern values,
  parse/format round trips and destination atomicity/suffix tests. Completion and
  documentation were updated. Snapshot generation passed (29 formatter, 17 pipeline,
  51 integration tests), followed by all 1006 compiler tests; execution outputs are
  unchanged. This source-only batch did not rerun full CI; the Unicode scanner
  batch remains the latest full-CI baseline.
- This starts strconv: 22 capabilities started, 52 planned-only. Floating-point
  conversion algorithms, quoting/unquoting and remaining numeric API acceptance
  remain unfinished. Full migration and publication remain outstanding.

### E1 HTML character-reference decoding — implemented and validated

- Added bounded one-shot unescape/unescape_with APIs with explicit Text/Attribute
  context. They reuse all 2,125 names, choose longest matches and recognize the
  106 pinned legacy semicolon-optional names. Attribute ambiguity preserves the
  original text; Markdown retains its independent strict parsing rules.
- Numeric decimal/hex references support optional semicolons, C1 replacements and
  invalid-scalar replacement without overflow. Unknown/malformed references remain
  literal; output is checked before append and decoded content is not rescanned.
- Retained the Go-derived data license and checksum-pinned legacy names. Extended
  native generation/verification, all-name decode tests, legacy-context tests,
  numeric edge/limit tests and the independent Go differential consumer.
- Focused tests and expanded html/template/markdown/tui_markdown verification
  passed, including generated-data checks, cached builds, smoke and PTY checks.
  Differential coverage includes 1,024 numeric references and named-prefix cases;
  all 2,125 canonical and 106 legacy names are independently checked. These are
  whole-string helpers, not an HTML tokenizer or sanitizer.
  No change to the overall counts: 21 started, 53 planned-only. Full A/B migration
  and publication remain outstanding.

### E1 shared HTML named-reference data — implemented and validated

- Moved the canonical 2,125-name JSON dataset into ecosystem/html/data without
  changing its pinned checksum. HTML retains the PSF license and owns the generated
  public `named_entity` lookup. Markdown's generated file now delegates to HTML;
  its semicolon/numeric parsing rules remain untouched.
- Extended the existing native Markdown generator to produce and verify both
  files from the shared data. Kept its CLI and verification registration compatible.
  HTML's own tests compare every public lookup against the independent data; the
  independent consumer exercises cross-module lookup as well as escaping.
- HTML/template/markdown tests and generator conformance passed; expanded
  tui_markdown verification also passed, including cache, smoke and real PTY checks.
  This is data/mechanics extraction, not yet
  HTML text/attribute reference parsing. Numeric references, legacy semicolon-free
  names and context-sensitive decoding remain unfinished. Counts stay 21 started
  and 53 planned-only; full migration and publication remain outstanding.

### E1 shared HTML escaping — implemented and validated

- Added independently versioned ecosystem::html with checked escaped lengths,
  bounded escaping, explicit quote options and the legacy unbounded text helper.
  The implementation uses pure GoML byte/string operations and no native backend.
- Template now delegates to the shared bounded engine while preserving numeric
  quote references, its source-aware Limit errors and public signatures. Markdown
  delegates to the text helper while preserving named quotes and apostrophes.
  Both declare the HTML module dependency; no standard package imports ecosystem.
- Added a module README, library tests, an independent versioned consumer and
  ecosystem verification registration. Private registry verification passed for
  html/template/markdown, including reference suites, cache stability, smoke tests
  and the existing Markdown entity-data conformance check. Expanded verification
  with explicit wrapper assertions and tui_markdown also passed, including the
  latter's real PTY checks. The registry/verification module passed all 12 tests.
- This starts HTML: 21 capabilities started, 53 planned-only. Entity decoding,
  shared table extraction and full HTML-specific conformance remain unfinished.
  No public registry version or toolchain release was published; existing immutable
  published versions are unchanged. Full migration/release remains outstanding.

### S2 Unicode scanner splitters — implemented and validated

- Added scan_runes and scan_words using the existing UTF-8 scalar decoder and
  pinned Unicode White_Space table. Rune tokens replace each malformed byte;
  whitespace-delimited words preserve original bytes and never emit empty tokens.
- Added ScanStep::Emit for transformed, independently copied tokens. Input limits
  do not bound custom callback output; this is explicit in the language guide.
  Standard io now declares Unicode directly in both source and link dependencies.
- Build and external Go bufio differential tests passed for line/rune/word
  splitters over nine source fragment sizes, all byte values, Unicode whitespace,
  malformed and truncated UTF-8. The earlier scanner's 1006 compiler tests passed.
- Snapshot generation passed (29 formatter, 17 pipeline, 51 integration tests).
  Full just ci exited successfully: 1006 compiler and 171 driver tests, bootstrap
  fixed point and packaging checks passed for the new dependency. Execution golden
  outputs remain unchanged. Logs use the stdlib-scanner-unicode prefix in _artifact.
- This continues bufio rather than starting another capability: 20 started and
  54 planned-only; full migration and release remain outstanding.

### S2 bounded scanner — implemented and validated

- Extended standard io with `Scanner[R]`, explicit positive pending-input limits,
  owned byte tokens, sticky errors and configurable `ScanStep` callbacks. Built-in
  splitters cover lines (LF/CRLF and final unterminated lines) and single bytes.
- Validates split ranges, progress and source byte counts. Exact-limit EOF works;
  overflow detection can consume one probe byte. Interrupted reads retry, while
  other failures stop without replay. No source ownership/close or syntax changes.
- Build and external tests passed: fragmented input, invalid UTF-8 preservation,
  empty tokens, exact/overflow limits, invalid callbacks/readers, skip progress,
  repeated EOF/error and token-storage independence. Snapshot generation passed:
  29 formatter, 17 pipeline and 51 integration tests, with unchanged execution
  outputs. All 1006 compiler tests passed; the subsequent Unicode scanner batch
  also passed full CI including this implementation.
- This starts bufio: 20 capabilities started, 54 planned-only. Unicode word/rune
  splitters and differential coverage were added by the following batch documented
  above. Broader buffered-interface acceptance remains unfinished; the complete
  migration and release are still outstanding.

### S1 typed error reports — implemented and validated

- Extended the existing standard-owned error package with `Report[E]`: typed
  leaf errors, immutable context wrappers and ordered aggregates. Empty joins
  return None; joins and immediate-cause inspection copy container structure.
- `find`, `find_map` and equality-based `contains` traverse leaves depth-first
  without recursion. Domain enums preserve heterogeneous causes without runtime
  reflection. Formatting also uses an iterative stack and a byte output buffer.
- Existing Error, ErrorKind, Details and runtime codes remain compatible. Reports
  do not recursively inspect arbitrary application error payloads, freeze mutable
  leaf handles or add reflection-based Go Is/As compatibility.
- Toolchain build and the external consumer passed, including 10,000 nested
  contexts, nested groups, matching/projection order, Unicode and alias isolation.
  Snapshot generation passed (29 formatter, 17 pipeline, 51 integration tests),
  as did all 1006 compiler tests including completion coverage. Execution outputs
  remain unchanged. This source-only batch did not rerun full CI; the preceding
  Base32 stream batch has the latest full-CI result.
- This starts the errors capability: 19 started and 55 planned-only. Older count
  statements below describe their respective batches, not the current total.
  The language guide includes Result::map_err and typed-kind lookup examples.
  Broader error interoperability remains acceptance work; the complete A/B
  migration and release are still outstanding.

### S2 positional I/O contracts — in progress

- Added standard ReadAt/WriteAt contracts and TransferError carrying both a
  transferred byte count and the existing structured I/O error. Exact/all
  helpers validate ranges and provider counts, preserve valid error progress,
  and make one call without retrying Interrupted or discarding its progress.
- Cursor implements positional reads and overwrite/append writes without
  changing its sequential position. Sparse writes are rejected. Sequential and
  positional writes snapshot overlapping input views before mutation.
- The focused external consumer passes all small offset/length read ranges,
  exhaustive source/destination overlap combinations, EOF progress, untouched
  output tails, offset overflow, malformed provider counts and single-call
  failure behavior. Public contracts and completion coverage are updated.
- Toolchain rebuild and just update-golden pass: 29 formatter, 17 pipeline and
  51 integration tests, with unchanged execution outputs. The full compiler
  suite passes 1006 tests, including public API completion coverage. Formatting
  and diff whitespace checks pass. This batch changes no package catalog or
  dependency graph; full CI was not rerun and remains required before release.
  This starts the io inventory item: 17 capabilities started, 57 planned-only.
  Section/offset/multi-stream adapters, further concrete providers and consumer
  integration remain; the full inventory and release are not complete.

### S2 bounded section readers — implemented and validated

- Added pure SectionReader over any ReadAt implementation, with checked base
  and length, relative positional access, sequential reads, shared seek state
  and nesting. Provider buffers are clipped to the declared interval, and
  invalid provider counts cannot advance the section position.
- Sequential reads preserve positive progress before delivering the underlying
  error on the next nonempty read. Empty reads and invalid seeks preserve that
  pending error; successful seeking clears it. Exact/all helpers and positional
  error semantics remain unchanged. A section is a live bounded view, not a
  snapshot or a guarantee that the underlying source has the declared length.
- Focused external tests pass for all small base/length/offset/buffer ranges,
  nested sections, shared handles, underlying position isolation, partial
  backing storage, untouched destination tails, deferred Interrupted/EOF,
  empty reads, invalid seeks and malformed provider counts. Public contracts
  and API completion coverage are updated.
- Toolchain rebuild and just update-golden pass: 29 formatter, 17 pipeline and
  51 integration tests. Execution outputs are unchanged. The complete compiler
  suite passes 1006 tests, including SectionReader completion coverage.
  Formatting and diff whitespace checks pass. No package catalog or dependency
  changes are needed; full CI was not rerun for this batch and remains a release
  gate. This extends the existing io item; 17 capabilities are started and 57
  remain planned-only.
  Offset writers, multi-stream adapters and integration remain outstanding.

### S2 relative offset writers — implemented and validated

- Added pure OffsetWriter over WriteAt with validated base/relative/end offset
  arithmetic, shared sequential position, independent positional writes and
  nesting. It does not seek, close or implicitly flush the underlying provider.
- Sequential writes advance by validated progress even on failure and return
  the cause immediately. The write_all override makes one call without retrying
  Interrupted, preventing partial writes from being replayed. Full-count errors
  are preserved; malformed provider counts do not advance position.
- The focused external consumer passes every small base/relative-offset pair,
  nested writes, shared handles, underlying position isolation, invalid seeks,
  arithmetic overflow, partial/full-progress errors, malformed provider counts
  and maximum-offset empty writes. Public contracts and completion coverage are
  updated. Toolchain rebuild and just update-golden pass: 29 formatter,
  17 pipeline and 51 integration tests; execution outputs are unchanged.
  The full compiler suite passes 1006 tests. Formatting and diff whitespace
  checks pass. The catalog and dependency graph are unchanged; full CI was not
  rerun for this batch and remains required before release.
- This extends the io item; 17 capabilities are started and 57 remain
  planned-only. Multi-stream adapters and consumer integration remain, along
  with all other recorded inventory work and release gates.

### S2 multi-stream adapters — implemented and validated

- Added pure MultiReader and MultiWriter over snapshotted homogeneous handle
  lists. Reader copies share traversal, release exhausted handles and preserve
  the current source on errors; empty reads never consume a source.
- Broadcast writes and flushes visit destinations in order and stop at the
  first failure. Short/invalid counts are diagnosed; write_all never replays
  earlier destinations after Interrupted. Repeated destinations remain repeated
  operations, and an empty list accepts writes or immediately exhausts reads.
- Focused tests pass for fragmented concatenation, empty inputs/lists, shared
  exhaustion, list snapshot isolation, nested broadcast, repeated destinations,
  malformed counts, short writes, failure ordering and recovery on the current
  reader. Public contracts and completion coverage are updated.
- Toolchain rebuild and just update-golden pass: 29 formatter, 17 pipeline and
  51 integration tests; execution outputs are unchanged. The full compiler
  suite passes 1006 tests. Formatting and diff whitespace checks pass. This
  batch changes no catalog or dependencies; full CI was not rerun and remains
  a release gate. The io item remains in progress: stream-copy helpers, tee/discard adapters and
  consumer integration still need work. Inventory counts stay at 17 started
  capabilities and 57 planned-only, with release gated on the full scope.

### S2 copy helpers and discard — implemented and validated

- Added copy_buffer over caller-provided nonempty scratch storage and reused
  it from the existing copy entry point. Added copy_n using the bounded Take
  adapter, with no extra input consumption, negative-count rejection and a
  zero-count fast path that performs no stream operations.
- Added a stateless Discard writer. Copy failures retain existing Result error
  semantics and side effects; no flush, close or rollback is implied. Provider
  counts remain checked, read interruption is retried, and destination-specific
  write_all overrides retain control over replay-sensitive failures.
- Focused external tests pass across input lengths, scratch sizes and requested
  copy counts, including fragmented writes, Interrupted, empty streams/buffers,
  early EOF, malformed counts, scratch-view boundaries and input consumption on
  destination failure. Public contracts and completion coverage are updated.
- Toolchain rebuild and just update-golden pass: 29 formatter, 17 pipeline and
  51 integration tests, with unchanged execution outputs. The full compiler
  suite passes 1006 tests. Formatting and diff whitespace checks pass. The
  catalog and dependency graph are unchanged; full CI was not rerun for this
  batch and remains a release gate. The io capability remains in progress,
  including tee and consumer integration;
  inventory counts remain 17 started and 57 planned-only. Full migration and
  release acceptance are still outstanding.

### S2 synchronous tee reader — implemented and validated

- Added pure TeeReader using Read and Write, with synchronous complete mirror
  writes and no retained staging buffer. Empty reads and source EOF avoid the
  mirror; source failures remain retryable according to their existing kind.
- Mirror failures terminate the tee and shared copies. The original cause and
  failed source-read count remain inspectable; consumed source bytes and partial
  mirror writes are not replayed. A terminal Interrupted cause is surfaced as
  Other to generic Read consumers so their retry loops cannot spin or replay;
  the original structured error remains available through failure().
- Focused external tests pass for chunked reads, fragmented/interrupted mirror
  writes, partial mirror failures, original error retention, shared terminal
  state, EOF/empty-read suppression, malformed source counts and non-retry of
  terminal interruption through read_exact. Public contracts and completion
  coverage are updated. Toolchain rebuild and just update-golden pass:
  29 formatter, 17 pipeline and 51 integration tests, with unchanged execution
  outputs. The full compiler suite passes 1006 tests. Formatting and diff
  whitespace checks pass. This batch changes no catalog or dependency graph;
  full CI was not rerun and remains a release gate.
- The io item still needs consumer integration and an inventory-wide remaining
  API audit; this does not complete its full scope or the migration/release.
  Counts remain 17 started capabilities and 57 planned-only.

### S1 varint cursor integration — implemented and validated

- Added signed/unsigned varint methods to the existing endian Reader and Writer.
  Successful operations advance shared cursor state by the encoded length;
  truncated/overflowing reads and insufficient-capacity writes leave the
  position unchanged. Failed writes leave destination bytes unchanged.
- Reused the existing checked algorithms and absolute error offsets without
  introducing an I/O dependency into the pure endian value package. Streaming
  adapters will live in a separate endian child package with explicit I/O
  dependencies; they are not implemented by this cursor batch.
- Toolchain rebuild and focused boundary tests pass for signed/unsigned limits,
  shared cursor state, every insufficient output capacity and truncated or
  overflowing input after a successfully consumed prefix. just update-golden
  passes 29 formatter, 17 pipeline and 51 integration tests; execution outputs
  are unchanged. The full compiler suite passes 1006 tests. Formatting and diff
  whitespace checks pass. This source-only batch changes no catalog or package
  dependencies; full CI was not rerun and remains a release gate.
  This addresses a foundation residual, not a new inventory capability or
  completion of the binary encoding item. Counts remain 17 started and 57
  planned-only, with full migration and release gates still outstanding.

### S2 streaming varint integration — implemented and validated

- Added std::bytes::endian::stream with bounded signed/unsigned varint reads
  and writes over the public I/O traits. Clean EOF is distinct from truncated
  prefixes, decode overflow and underlying I/O failures. Reads consume one byte
  at a time and at most ten bytes; writers reuse the checked memory encoder and
  respect destination write_all policies.
- Kept the endian value package independent of I/O. Registered the child source
  package, compiler dependency selection, source closure, completion/navigation
  and external consumer coverage. Source/decoded offsets and failed-read
  progress limitations are documented along with non-rollback behavior.
- External tests pass against Go-encoded signed/unsigned boundary values,
  noncanonical valid input, interruption, malformed reader counts, early EOF,
  all prefix truncations, tenth-byte overflow, tail preservation, original I/O
  errors and write_all no-retry policy. Toolchain rebuild and just update-golden
  pass: 29 formatter, 17 pipeline and 51 integration tests, with unchanged
  execution outputs. Full just ci passes with Go 1.26.8: 1006 compiler tests,
  171 driver tests, bootstrap fixed-point verification and packaging/release
  checks. Formatting checks were rerun successfully after CI rebuilt stage2;
  diff whitespace checks pass. This closes the separate varint streaming gap,
  not the entire encoding/binary capability or migration.
  Counts remain 17 started and 57 planned-only; fixed-width stream composition
  and remaining inventory work still require implementation before release.

### S2 fixed-width scalar streams — implemented and validated

- Added all signed/unsigned 8/16/32/64-bit and IEEE f32/f64 stream readers and
  writers. Shared bounded helpers compose the existing memory codecs with
  public I/O traits; no new source package, dependency or native algorithm is
  introduced. Multibyte APIs accept explicit endian order.
- Reads handle short transfers and interruption, distinguish clean EOF from
  truncated fields, preserve original errors with confirmed prefix progress
  and consume exactly the field width. Writes encode before I/O and preserve
  the destination write_all retry policy. Float bit patterns are retained.
- Boundary tests pass for all scalar types and both byte orders, chunk sizes,
  all truncation lengths, invalid provider counts, original error progress,
  no-retry writes, signed zero and NaN payloads. Toolchain rebuild and
  just update-golden pass: 29 formatter, 17 pipeline and 51 integration tests,
  with unchanged execution outputs. The full compiler suite passes 1006 tests.
  Formatting and diff whitespace checks pass. The source catalog and dependency
  graph are unchanged; full CI was not rerun for this batch and remains required
  before release. This extends encoding/binary, not a new inventory item. Explicit
  aggregate schemas and the remaining inventory audit are still outstanding;
  counts remain 17 started and 57 planned-only, with release gates unchanged.

### S1 PEM framing — implemented and validated

- Added pure PEM blocks, bounded canonical encoding and bounded one-block
  decoding over Base64. Input/header validation prevents newline injection;
  duplicate keys and malformed candidates produce explicit errors. Block
  construction and mutable accessors use snapshots. Encoding sorts headers with
  Proc-Type first and wraps body lines at 64 columns.
- Registered the standard source package, compiler dependency selection,
  navigation/completion and source-only coverage. Documented strict differences
  from Go's searching decoder, byte-offset conventions, resource limits and the
  exclusion of certificate/encryption policy. External tests pass against Go
  encoding/pem for 193 payload lengths with and without headers, LF/CRLF,
  trailing whitespace, missing final newline, preambles and consumed offsets.
  Exact/insufficient bounds, malformed input, duplicate headers, multiple blocks
  and snapshot isolation also pass. Toolchain rebuild and just update-golden
  pass: 29 formatter, 17 pipeline and 51 integration tests, with unchanged
  execution outputs. Full just ci passes with Go 1.26.8: 1006 compiler tests,
  171 driver tests, bootstrap fixed point and packaging/release checks.
  Formatting and diff whitespace checks pass. Streaming adapters remain
  separate residual work.
- This starts the PEM inventory capability: 18 started, 56 planned-only. It does
  not complete the full A/B migration or authorize an incomplete release.

### S1 reusable Base32 configurations — implemented and validated

- Added validated immutable Encoding values with custom ASCII alphabets,
  optional padding, variant construction and configuration accessors. A private
  lookup table replaces per-digit alphabet scans; existing free functions use
  the same engine and preserve their strict default behavior.
- Configurations reject duplicate/non-ASCII/CR/LF symbols and conflicting
  padding; explicit space/tab/NUL symbols are data, not ignored whitespace.
  Raw alphabets may contain equals as an ordinary digit. Constructors return
  recoverable errors rather than panicking. Public restrictions and examples
  are documented, with completion coverage and Go differential tests added.
- Go differential tests pass for 32 alphabet rotations, 41 input lengths,
  padded/raw variants, every allowed ASCII padding byte, literal equals and
  whitespace symbols, invalid configurations and canonical decode failures.
  Toolchain rebuild and just update-golden pass: 29 formatter, 17 pipeline and
  51 integration tests; execution outputs are unchanged. The full compiler
  suite passes 1006 tests. Formatting and diff whitespace checks pass. The
  package graph and source catalog are unchanged; full CI was not rerun for
  this batch and remains a release gate. Streaming Base32 adapters and bounded output
  interfaces remain residual work; inventory counts stay at 18 started and 56
  planned-only, with the full migration/release goal unchanged.

### S1 bounded Base32 output — implemented and validated

- Added allocation-free checked encoded lengths, Encoding encode_checked and
  decode_checked methods, and standard-configuration convenience functions.
  Overflow-safe group/tail arithmetic handles both padded and raw encodings.
  Shared decode shape validation checks output bounds before allocation while
  retaining original strict digit/trailing-bit errors inside OutputError.
- Documented error precedence, exact-limit behavior and the distinction between
  output bounds and input/process-memory bounds. Tests pass for four variants,
  257 input lengths, exact/insufficient limits, maximum representable lengths,
  overflow without allocation, custom equals digits, error precedence and
  retained DecodeError details. Toolchain rebuild and just update-golden pass:
  29 formatter, 17 pipeline and 51 integration tests, with unchanged execution
  outputs. The complete compiler suite passes 1006 tests. Formatting and diff
  whitespace checks pass. No catalog/dependency changes are required; full CI
  was not rerun for this batch and remains a release gate. Streaming adapters remain
  outstanding, and this does not complete the full Base32 or migration scope.
  Counts remain 18 started capabilities and 56 planned-only.

### S1 incremental Base32 encoder — implemented and validated

- Added bounded Encoder state with arbitrary chunk boundaries, a scalar bit
  accumulator, shared counters, explicit/idempotent finish and reset. Input
  slices and output chunks are not retained. Updates reserve cumulative final
  size including tail/padding, so limit failures leave state unchanged and any
  accepted prefix can always be finalized within budget.
- Added external Go differential tests across standard/custom alphabets,
  chunk boundaries, padding policies, handle aliases, reset/finish lifecycle
  and atomic limit failures. Tests pass for six configurations, 81 input lengths
  and twelve chunk sizes, including custom equals digits and NUL padding.
  Contracts and completion coverage are updated. Toolchain rebuild and
  just update-golden pass: 29 formatter, 17 pipeline and 51 integration tests,
  with unchanged execution outputs. The full compiler suite passes 1006 tests.
  Formatting and diff whitespace checks pass. No catalog/dependency changes
  are needed; full CI was not rerun and remains a release gate. Incremental
  decoding and I/O wrappers remain.
  Counts stay at 18 started and 56 planned-only; release remains gated on the
  complete inventory, not this subset of Base32.

### S1 Base32 I/O adapters — implemented and validated

- Added standard-owned `std::encoding::base32::stream`, with generic Write/Close
  encoder and Read decoder over the existing incremental state APIs. Core Base32
  remains independent of I/O; catalog and driver selection include the child.
- Explicit finish emits padding without closing the destination. Flush does not
  finalize. Destination failures are sticky, preserve the original cause and do
  not replay partially emitted output. Reader validation continues through EOF.
- Toolchain build and external consumer passed. The consumer compares four built-in
  and two custom encodings against Go's test-only oracle across 34 input lengths
  and nine chunk sizes. It covers fragmented/interrupted I/O, repeated finish/EOF,
  malformed input, limit rejection, partial-output failure without replay, sticky
  flush/source failures, invalid reader counts and trailing data after padding.
- Full just ci passed: 1006 compiler tests, 171 driver tests, fixed-point and
  packaging checks. After adding failure injection and completion/navigation cases,
  the focused consumer and all 1006 compiler tests passed again. Execution golden
  outputs are unchanged. Core Base32 remains free of an I/O dependency.
- This advances the existing Base32 capability, not the completed inventory count.
  The full migration and release remain outstanding; counts stay 18 started and
  56 planned-only.

### S1 incremental Base32 decoder — implemented and validated

- Added bounded Decoder state with at most seven retained symbols, complete
  quantum decoding, absolute error offsets, final-padding enforcement and
  explicit finish/reset. Raw final groups retain strict canonical-bit checks.
- Updates commit state/output only after validating the complete submitted
  chunk. Limit/decode failures allow a corrected or smaller retry without
  replaying earlier output. Incomplete finish leaves state open; successful
  finish is idempotent and closes updates. Pending decoded bytes reserve budget.
- Added standard/custom alphabet, split-padding, lifecycle, aliasing, absolute
  error, limit and transactional-failure tests. Public contracts and completion
  coverage are updated. Tests pass for six configurations, 65 input lengths
  and eleven chunk sizes against Go-encoded input. Toolchain rebuild and
  just update-golden pass: 29 formatter, 17 pipeline and 51 integration tests,
  with unchanged execution outputs. The complete compiler suite passes 1006
  tests. Formatting and diff whitespace checks pass. The catalog and dependency
  graph are unchanged; full CI was not rerun and remains a release gate.
  I/O adapters remain, so this
  does not complete Base32 or the full migration/release goal. Counts remain
  18 started and 56 planned-only.

### S1 foundation batch — in progress

- Implemented initial public surfaces for eight capabilities: Base32, binary
  varints, hash interfaces, Adler32, CRC32, CRC64, FNV and fixed-width bit operations.
- Added the standard source catalog, compiler dependency selection, source/link
  closure tests, completion/navigation tests, public API examples and limitations.
- External consumer project093_std_algorithms compares against Go standard-library
  test-only FFI oracles: Base32 across 257 input lengths, signed/unsigned varints,
  checksums at block/chunk boundaries and wide integer arithmetic. It also covers
  malformed input, overflow, no-partial-write errors, reset and snapshot isolation.
- Targeted consumer and source-world interface reconstruction tests pass.
  just update-golden passes: 29 formatter, 17 pipeline and 51 compiler integration
  tests. Full just ci passes with Go 1.26.8: 1006 compiler tests, 171 driver tests,
  bootstrap fixed-point verification and packaging/release smoke checks.
- Error types use explicit trait implementations so both finalized-world and
  retained source-world construction paths can consume the new packages.
- Base32 custom alphabets and streaming adapters, varint stream/stateful cursor
  adapters, and checksum I/O adapters/state serialization remain outside this
  initial surface. Fixed-width bit APIs have no native-word-width aliases.
- At the end of this batch, 66 inventory items remained planned. It was not completion of
  the full A/B migration and does not claim drop-in Go API compatibility.

### S1 logical slash paths — implemented and validated

- Added std::path::slash without host/runtime algorithm dependencies: clean,
  join, split, base, dir, extension and is_absolute. Existing std::path semantics
  are unchanged; lexical normalization is explicitly not a security sandbox.
- Added reusable immutable Pattern compilation and full-name Unicode matching
  with stars, single-character wildcards, classes, negation, ranges and escapes.
  Pattern length and deterministic work budgets return recoverable errors;
  matching uses pattern-sized dynamic-programming state without backtracking.
- Added source catalog, dependency selection and isolation, navigation/completion,
  public API documentation and external consumer tests. The consumer compares
  lexical operations and valid patterns to Go path, including Unicode and
  backslash cases, invalid patterns, repeated calls and resource-limit boundaries.
- The focused consumer and expanded full just ci pass with Go 1.26.8: 1006
  compiler tests, 171 driver tests, bootstrap fixed-point verification and
  packaging/release smoke checks. Formatting and diff whitespace checks pass.
  There are now nine inventory capabilities with implementations and 65 still planned; the
  initial eight retain the residual scope recorded above. Release remains gated
  on the entire A/B inventory, not this batch.

### S1 maps and sorting — implemented and validated

- Added pure std::collections map helpers corresponding to Go maps clone,
  equality/custom equality, copy, conditional deletion, entry/key/value iteration,
  iterator insertion and collection. Existing HashMap storage and hashing are
  unchanged. Snapshot timing, shallow aliasing, duplicate overwrite, callback
  mutations and unspecified order are explicit public contracts.
- Retained the existing pure GoML stable sorting implementations and added
  monotone indexed search, sortedness checks, lower/upper insertion boundaries
  and equal ranges, with generic and comparator forms. No compiler-owned fallback
  is removed, no new sort backend or float total order is introduced.
- Expanded external differential coverage against Go maps and sort, plus custom
  hashing collisions, independent entry storage/shared references, generic
  iterators, NaNs, callback mutation, comparator direction, duplicate stability,
  machine-maximum search lengths and empty-input behavior.
- Combined external differential tests and formatting checks pass with Go 1.26.8.
  The CI run passed 171 driver tests, bootstrap fixed-point verification and
  packaging/release smoke checks; its only compiler failure was the old expected
  collections source-file count (13 instead of 14). After updating that test,
  the complete compiler CI suite was rerun: 1006 passed, zero failed. No production
  code changed between the CI run and this successful compiler recheck.
- Eleven inventory capabilities now have implementations; 63 remain planned,
  in addition to the recorded initial-batch residual interfaces. These two items
  extend the existing collections package rather than introducing new modules.

### S1 slices — implemented and validated

- Added pure std::collections view equality/comparison/membership, checked
  concatenation/repetition, overlap-safe vector range edits, stable predicate
  deletion, checked capacity reservation, view reversal/sorting, sortedness,
  insertion-point search and optional extrema.
- Added forward/backward indexed view iterators, checked zero-copy chunking,
  generic iterator collection, stable sorted collection and iterator append.
  Iterator append uses iterator-first arguments so the associated item type is
  inferred before the destination. Existing clone, compact and value-iterator
  methods remain canonical. Go capacity clipping has no equivalent public view
  operation; independent length-sized copies use Slice::to_vec.
- Documented live bounded views versus shallow independent copies, shared Vec
  edits, no-partial-mutation validation errors, negative/overflow checks,
  allocation limits, callback restrictions and explicit float ordering policy.
- Added catalog and navigation coverage. External differential tests pass
  against Go slices for equality, comparison, indexing, repetition, deletion,
  replacement and compaction, with all small edit ranges, overlapping/self
  insertion, view boundaries, extrema ties, iterator exhaustion and huge counts.
- Full just ci passes with Go 1.26.8: 1006 compiler tests, 171 driver tests,
  bootstrap fixed-point verification and packaging/release smoke checks.
  Formatting and diff whitespace checks pass. Twelve inventory capabilities now
  have implementations; 62 remain planned, in addition to the initial foundation
  batch's residual scope. This is not completion of the full migration/release goal.

### S1 UTF-8 scalar and incremental decoding — implemented and validated

- Added pure shared prefix inspection, checked first-scalar decoding,
  replacement forward/reverse decoding, full-prefix checks, scalar validity,
  rune-start checks, malformed-byte-aware counting and lossy text conversion.
- Added direct scalar encoding and checked destination/builder output. Whole
  validation reuses the same decoder and now correctly distinguishes already
  invalid short prefixes from genuinely incomplete ones. Existing valid-string
  conversion boundaries remain unchanged.
- Added a strict incremental Decoder with bounded pending storage, per-byte
  scalar delivery, absolute offsets, sticky UTF-8 failures, idempotent successful
  finish, closed-state rejection, offset-overflow checks and shared-handle reset.
- Added source/catalog, navigation/completion and source-only checks, API
  examples and explicit strict-versus-replacement contracts. Focused external
  tests pass against Go unicode/utf8 for every single-byte and two-byte input,
  sampled scalar encoding, boundary prefixes and deterministic malformed inputs;
  streaming failures/reset, shared handles and no-partial-write encoding pass.
- CI passed 171 driver tests, fixed-point verification and packaging/release
  smoke checks. Five pipeline shards initially reported outdated IR snapshots
  for the new constants, decoder types and shared validator. just update-golden
  regenerated them and passed 29 formatter, 17 pipeline and 51 integration tests;
  execution outputs were unchanged. The complete compiler CI suite then passed
  on recheck: 1006 tests, zero failures. Formatting and diff checks pass.
- Thirteen inventory capabilities now have implementations; 61 remain planned,
  in addition to the initial foundation batch's residual scope. Release is still
  gated on the full inventory, not this validated batch.

### S1 Unicode tables and casing — implemented and validated

- Replaced the remaining Go classification and casing calls in std::unicode
  with pure GoML lookup/conversion. Retained Unicode 15.0.0 and the existing
  multi-scalar full case-fold API; no normalization or grapheme policy is added.
- Added all named category, script, property, fold-category and fold-script
  tables, immutable checked custom range tables, all scalar predicates, simple
  fold cycles, title casing and immutable custom/Turkish/Azeri case overrides.
  Raw code-point membership supports surrogate categories without admitting
  invalid values into char. Constructors and exported snapshots do not alias.
- Added reproducible Go-based table generation with a Unicode-version guard,
  canonical formatting and CI drift checks. Production code has no Go FFI;
  test-only Go oracles stay in the external consumer. Documented data provenance,
  range encoding, API contracts, limitations, examples and syntax compatibility.
- The focused consumer passes comparisons over every valid Unicode scalar for
  all 13 predicates, upper/lower/title/simple-fold and Turkish casing. Every named
  table matches Go's exact ranges and boundary membership. Invalid constructors,
  maximum-width strides, unknown names and snapshot isolation are covered.
- Source-only dependency checks and navigation/completion coverage are included.
  Full just ci passes with Go 1.26.8: 1006 compiler tests, 171 driver tests,
  bootstrap fixed-point verification, data regeneration checks and packaging/
  release smoke checks. Formatting and diff whitespace checks pass. This batch
  requires no additional pipeline snapshot changes.
- Fourteen inventory capabilities now have implementations; 60 remain planned,
  in addition to the foundation batch's residual scope. This is not completion
  of the full A/B migration or authorization to release an incomplete batch.

### S1 byte search and shared cuts — in progress

- Added pure arbitrary-byte first/last search, membership, prefix/suffix checks,
  single-byte searches and shared-view cut/cut-prefix/cut-suffix operations.
  Empty patterns, overlapping last matches and absent matches have explicit
  contracts. No UTF-8 interpretation or new dependency cycle is introduced.
- Added reusable immutable Finder state with a copied pattern and KMP fallback
  table: linear preprocessing and linear search, constant search workspace.
  Mutating the original pattern cannot invalidate the compiled search state;
  cut results intentionally remain views of the original input storage.
- Added source catalog, source-only checks, completion coverage and public
  documentation. Toolchain rebuild and focused external differential tests pass:
  all 65,025 pairs of binary-alphabet sequences through length seven, all byte
  values, overlapping matches, snapshots, aliasing and long repetitive inputs.
- just update-golden passes 29 formatter, 17 pipeline and 51 integration tests;
  generated IR includes Finder while execution outputs remain unchanged. Full
  just ci passes with Go 1.26.8: 1006 compiler tests, 171 driver tests, bootstrap
  fixed-point verification and packaging/release smoke checks. Diff whitespace
  checks pass. This advances the bytes item but does not mark it complete:
  split/replace, bounded
  output construction and remaining byte/text operations still need work. The
  previously recorded inventory count is unchanged for this partial batch.

### S1 bounded byte transformations — implemented and validated

- Added pure count, split/split-after with optional maximum part counts,
  replacement with optional maximum replacement counts, join and repeat.
  Split count semantics are separate from the checked resource limit; negative
  counts mean all splits/replacements, but negative repetition counts fail.
- Empty separators follow Go replacement-decoding boundaries: valid UTF-8
  scalars stay intact and invalid/truncated bytes each advance one byte. A small
  private width checker avoids a bytes -> utf8 -> bytes cycle; no Unicode
  property tables or production native calls are added to the byte foundation.
- Split results are shared views. Constructed output buffers are independent;
  exact byte limits succeed, overflow and limits fail before output allocation,
  empty repetition is constant-time, and errors expose no partial output.
- Toolchain rebuild and external differential tests pass: short binary sequence
  pairs, all two-byte sequences, multibyte truncations and byte mutations, bounded
  counts, empty separators, aliasing, precise limits and error variants.
  just update-golden passes 29 formatter, 17 pipeline and 51 integration tests;
  execution outputs are unchanged. Full just ci passes with Go 1.26.8: 1006
  compiler tests, 171 driver tests, bootstrap fixed-point verification and
  packaging/release smoke checks. Formatting and diff whitespace checks pass.
- This completes the previously recorded split/replace and bounded construction
  gap, not the entire bytes migration. Unicode-oriented byte helpers, lazy
  traversal and reader/buffer integration remain to be addressed alongside text
  and I/O. Counting the partial bytes work, 15 inventory capabilities now have
  implementations and 59 remain planned. This is a started-capability count,
  not a completed-capability count; all recorded residual scope still applies.
  Release remains gated on the full A/B scope.

### S1 lazy byte traversal — implemented and validated

- Added split_iter and split_after_iter with snapshotted separators, live shared
  input views, constant traversal state and no eager result-vector allocation.
  Iterator aliases share consumption; independent construction restarts it, and
  exhaustion is permanent. Eager/count-limited splitting now uses the same engine.
- Added lines_iter preserving terminating LF and any CR bytes, with no extra
  empty line after a trailing LF. Arbitrary malformed binary input stays intact.
- Toolchain rebuild and external consumer tests pass against Go split and Lines,
  including existing malformed-UTF-8/count/limit cases, repeated exhaustion,
  separator snapshots, aliasing and changing an unconsumed delimiter between
  calls to prove lazy scanning. Completion coverage and API contracts are updated.
- just update-golden passes 29 formatter, 17 pipeline and 51 integration tests;
  execution outputs are unchanged. The full compiler suite passes 1006 tests.
  Formatting and diff whitespace checks pass. The preceding transformation
  batch passed full CI; this source/API-only follow-up does not
  change the catalog, dependency graph, driver or release packaging mechanism.
  Full CI is not rerun for this follow-up and remains a gate before release.
- Lazy split/line traversal is implemented; Unicode-oriented byte helpers and
  reader/buffer integration remain. This does not complete the bytes inventory
  item or the overall migration/release goal.

### S1 text search, cuts and lazy traversal — implemented and validated

- Added a reusable pure text Finder with linear prefix preprocessing and linear
  byte-offset search, without copying the source string into a byte vector.
  Existing public text find/rfind/find_bytes keep their results while using it;
  compiler-owned string-method implementations remain unchanged.
- Added checked nonoverlapping/scalar-boundary count, cut and edge cuts, lazy
  split/split-after, legacy-compatible lines and inclusive line traversal.
  Existing empty-separator and CR handling stay explicit: GoML string split
  retains the whole input for an empty separator; cut accepts the empty match,
  while existing split_once still rejects it. Inclusive lines preserve CR/LF.
- Toolchain rebuild and external consumer tests pass against Go strings for
  search/count/cuts and nonempty splits, and against existing GoML methods for
  retained semantics. Unicode, NUL, overlapping matches, long repeated prefixes,
  iterator aliasing, exhaustion and line endings are covered. Source-only and
  completion checks plus API documentation are updated.
- just update-golden passes 29 formatter, 17 pipeline and 51 integration tests;
  execution outputs are unchanged. The full compiler suite passes 1006 tests.
  Formatting and diff whitespace checks pass. No catalog, dependency, driver or
  packaging changes are needed for this follow-up. Full CI is not rerun in this
  batch and remains a gate before release.
- This starts the strings inventory item, bringing the started count to 16 and
  planned-only count to 58. Unicode-aware field/trim/map operations, checked
  construction and remaining strings APIs are still required; this is not full
  strings completion or completion of the A/B migration/release goal.

### S1 Unicode-aware text and dependency layering — implemented and validated

- Removed Unicode's dependency on text::StringBuilder through private builtin
  byte-vector assembly. All Unicode public entry points, data versions and table
  encodings remain unchanged; the full-fold generator now emits that same private
  assembly path. Unicode has no standard-package dependencies and text depends
  on Unicode, with source-closure tests guarding the acyclic direction.
- Added Unicode whitespace and predicate field iterators, bounded collection,
  left/right/both predicate trimming, scalar-set trimming, trim_space and
  simple-fold equality. Existing ASCII trim APIs and multi-scalar full folding
  remain distinct and compatible. Predicate visit order, exhaustion, result
  limits and callback side effects are explicit contracts.
- Focused consumer tests pass, including the existing full-scalar Unicode suite,
  every whitespace-table member, non-whitespace lookalikes, all changed BMP
  uppercase mappings, special fold classes, predicate order/laziness, exact
  field limits and existing ASCII behavior. Go predicate oracles use test-only
  fixed-predicate adapters because char callbacks are not accepted by Go FFI.
- Table regeneration checks pass. just update-golden passes 29 formatter,
  17 pipeline and 51 integration tests; execution outputs are unchanged. CI
  passed 171 driver tests, fixed-point verification and packaging/release smoke
  checks. Its only compiler failure was the old JSON source/link closure
  expectation omitting Unicode. After updating both expected closures, the full
  compiler suite passed on recheck: 1006 tests, zero failures. The initial just ci
  invocation exited nonzero; this record does not describe it as a first-pass
  success. Formatting and diff whitespace checks pass.
- Checked text construction, text/byte mapping, byte Unicode helpers and I/O
  integration still need work. The started-capability count stays 16, with 58
  planned-only entries and all previously recorded residual scope still active.

### S1 bounded text construction and scalar mapping — implemented and validated

- Added checked join, repeat, all/count-limited replacement, one-pass optional
  scalar mapping and simple/special Unicode casing. UTF-8 byte limits are
  explicit, exact bounds succeed, and overflow/invalid counts are recoverable.
- Join/repeat/replace check final lengths before allocating output buffers.
  Replacement reuses one compiled pattern in two constant-workspace scans;
  checked empty-pattern replacement inserts at scalar boundaries while legacy
  unbounded replacement remains unchanged. Empty repetition avoids count-sized
  work even at the maximum representable count.
- Mapping calls a stateful transform once per visited scalar, supports deletion,
  and checks each encoded result before appending. Failure returns no partial
  text but retains prior callback side effects; negative limits invoke no callbacks.
- Toolchain rebuild and external differential tests pass against Go strings
  replacement, repetition, mapping and normal/Turkish casing, plus exact and
  insufficient limits, overflow, stateful callbacks and legacy compatibility.
  Public contracts, examples and completion coverage are updated.
- just update-golden passes 29 formatter, 17 pipeline and 51 integration tests;
  execution outputs are unchanged. The full compiler suite passes 1006 tests.
  Formatting and diff whitespace checks pass. This follow-up does not change
  catalog/dependency/packaging structure; full CI is not rerun in this batch and
  remains a release gate. Remaining text APIs,
  byte Unicode helpers and I/O integration are still outstanding; started count
  remains 16, with 58 planned-only entries and the full release gate unchanged.

### S1 bounded text splitting and scalar search — implemented and validated

- Added checked split/split-after and count-limited forms, with independent
  result-part limits, scalar decomposition for empty separators and unsplit
  remainders for positive counts. Existing split and split iterators retain
  their whole-input empty-separator behavior; count and resource bounds are not
  conflated, and no partial vector escapes on errors.
- Added scalar, scalar-set and predicate first/last search, set membership and
  one-sided scalar-set trimming. Search results are UTF-8 byte offsets; predicates
  visit each inspected scalar once in the chosen direction and stop at a match.
- Toolchain rebuild and external differential tests pass against Go strings for
  all split/count variants, scalar/set search and edge trimming. Coverage includes
  machine-minimum/maximum counts, exact and insufficient limits, Unicode/NUL,
  callback visit order and retained legacy behavior. Docs and completion tests
  are updated.
- just update-golden passes 29 formatter, 17 pipeline and 51 integration tests;
  execution outputs are unchanged. The full compiler suite passes 1006 tests.
  Formatting and diff whitespace checks pass. The package graph and packaging
  mechanism are unchanged; full CI is not rerun in this batch and remains a
  release gate. This is further work on the
  existing strings item, not completion of the remaining A/B scope or release.

### S1 Unicode-aware byte operations — implemented and validated

- Added replacement-decoded scalar/set/predicate searches, forward/reverse
  traversal, Unicode/predicate fields, bounded collection, edge trimming,
  scalar-set trimming, one-pass scalar mapping, normal/special casing and
  simple-fold equality. The existing validated boundary-width helper supplies
  forward decoding; reverse decoding handles invalid suffixes one byte at a time.
- Views preserve exact original bytes and alias the source. Mapping/casing
  return independent valid UTF-8 and apply explicit byte limits before appending.
  U+FFFD matching includes malformed bytes; fold equality is scalar/simple,
  not raw byte equality or full multi-scalar folding. Callback effects are not
  rolled back after an output-limit error.
- Bytes now depends on the dependency-free Unicode package, not text or utf8;
  the graph stays acyclic and shares one pinned Unicode dataset. Source catalog,
  source/link closure expectations, isolation checks and completion coverage are
  updated, with API contracts and examples.
- Toolchain rebuild and external tests pass against Go bytes: every single and
  two-byte input for decoder-sensitive operations, multibyte truncation/mutation,
  normal/Turkish casing, fields/trimming, malformed-byte fold equivalence, exact
  limits, aliasing, independent output and callback behavior.
- just update-golden passes 29 formatter, 17 pipeline and 51 integration tests;
  execution outputs are unchanged. Full just ci passes with Go 1.26.8: 1006
  compiler tests, 171 driver tests, bootstrap fixed-point verification and
  packaging/release smoke checks. Formatting and diff whitespace checks pass.
  I/O integration and remaining inventory
  items are not complete; the started count remains 16 and planned-only count 58.
