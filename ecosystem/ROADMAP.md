# Functional backlog

The modules are usable within their documented scope. This list tracks the
remaining gaps from the ecosystem audit; an implemented module is not a claim
of feature parity with every mature library in that category.

## First improvement batch

| Module | Added |
| --- | --- |
| proptest | Lazy shrinking with shared work limits, structured pass/fail/discard reports, classification and coverage, persistent regression seeds, unique collections, model-valid command sequences and full-width numeric edge cases |
| redis | Injectable duplex transports and dialing with a shared setup/operation deadline, adapter validation and failure cleanup |
| lsp | Deferred request acceptance/completion, cooperative cancellation, deadline polling, pending-request cleanup on exit and persistent rope document snapshots |
| cli | Argument/command aliases, inherited global options, cardinality-constrained argument groups and derive support for option aliases/globals |
| diff | Explicit linear-space Hirschberg algorithm for sequences, text and patches, with work/workspace limits |
| parser | Shared text/binary work and depth limits, contextual custom parsers, iterative alternatives and binary backtracking/lookahead primitives |

Each addition has library and independently resolved consumer coverage. The
verification entry point is `just ecosystem-test`, implemented in GoML; CI integration
is intentionally deferred.

## Adoption of GoML 0.1.50

| Module | Improved |
| --- | --- |
| cli | Ordinary command schemas through erased generic functions, including generic forwarding and function values; existing typed wrappers remain compatible |
| ndarray | Specialized `Array::[f64]::linspace` constructor and scalar `ToFloat` count conversion |
| template | Direct standard integer-to-float conversion for numeric coercion |
| sqlite | Direct standard I/O in the FFI consumer; removed the error-alias transport workaround |
| tempfile | Standard resource cleanup/error combination with compatible scope results |
| reqwest | Immutable request/response/multipart byte snapshots and shared snapshot accessors |
| msgpack | Standard reader/writer integration, direct typed frame decoding, bounded concatenated values and partial-I/O handling |
| redis | Bundled DNS/TLS connectors, mTLS and standard Context cancellation/deadlines composed with existing operation controls |
| bitflags | Optional `FlagValues` inherent derive for named flag constructors without trait imports |

## Second improvement batch

| Module | Added |
| --- | --- |
| proptest | Bounded multi-failure campaigns with shared shrink budgets, categorical histograms and text reports, atomic batch replay persistence, and state-machine execution with reset/invariant/cleanup handling |
| redis | Bounded connection pools, shared-deadline checkout, idempotent leases, configurable PING health checks, idle/lifetime expiry and replacement of unusable connections without replaying user commands |
| lsp | UTF-8/UTF-16/UTF-32 position-encoding negotiation, consistent document queries/edits, outgoing request deadlines and combined event-loop wakeup scheduling |
| cli | Nested flattened Args, typed required/optional subcommand derives, generic payloads, multilevel aliases/help/globals and composition diagnostics |
| notify / walkdir | Complete filesystem notification and traversal packages moved from `lib/std/fs` into normal versioned dependencies, with their original behavior suites retained as independent consumers |

## Terminal and color batch

| Module | Added |
| --- | --- |
| color | Checked color spaces, CSS values, alpha compositing, contrast and differences, gamut policies, hue interpolation and gradients |
| unicode_text | Version-pinned Unicode 16 tables, full grapheme/word/line conformance, terminal width policies and bounded text layout |
| ansi | Typed styles and hyperlinks, palette reduction, streaming escape tokenizer, styled Unicode text and standard writer adapters |
| terminal | Linux raw sessions, incremental typed input, mouse/paste/focus/resize, cancellable I/O, synchronized aliases and explicit restoration |
| tui | Cell invariants, constrained layout, incremental frames, common widgets, focus, Unicode editing and bounded history |
| prompt | Generic validators, history/completion, password display, search/select/multiselect, confirmation and cancellable model execution |
| progress | Concurrent job state, pure snapshots, rate/ETA, throttled bars/spinners, coordinated logs, plain output and bounded shutdown |
| diagnostics | Owned source identity, byte spans, multi-file labels, Unicode/tab alignment, themes, clipping and checked multi-file suggestions |
| tui_markdown | CommonMark terminal rendering, optional pipe tables, link handling, themes, scrolling and search |

The Explorer example composes the libraries with `walkdir` and `notify`.
Verification uses independent registry consumers, official/reference data,
pseudo-terminals and race checks where relevant. CI integration remains deferred.

## Application libraries batch

| Module | Added |
| --- | --- |
| fuzzy | Unicode case-folded matching, original grapheme ranges, explicit anchored modes, bounded maximum-score DP, stable Top-K, immutable incremental search sessions and parallel cancellation |
| config | Layered sources, typed Serde, JSON Pointer overrides, array policies, provenance, validators, immutable concurrent snapshots and caller-driven inotify reloads |
| websocket | Shared RFC 6455 protocol core, client/server handshakes, strict frames, fragmented UTF-8, control/close states, queue bounds and concurrent cancellable TCP/TLS I/O |
| highlight | Extensible logos grammars, built-in language scopes, nested/multiline state, embedded Markdown code, immutable incremental documents and bounded ANSI/HTML renderers |
| archive | USTAR/PAX and classic ZIP read/write, CRC32, bounded standard I/O, Deflate/GZIP adapters, metadata, rooted extraction and GNU/Info-ZIP interoperability |
| metrics | Concurrent metric handles, checked counters/gauges/histograms, consistent snapshots, registration/cardinality policies, atomic gauge collection, timers and Prometheus text output |
| csv | Standard reader/writer streaming, configurable dialects and quoting, multiline fields, BOM and header policies, byte/UTF-8 records, positions, resource bounds and typed Serde schemas |
| bench | Adaptive warmup/sampling, setup exclusion, checked monotonic clocks, bootstrap intervals, outlier statistics, baseline comparisons, throughput and standalone JSON/HTML reports |

The batch uses native GoML tests and independent versioned consumers. Local
verification includes race checks; no CI workflow was added.

## Remaining work

| Module | Remaining capabilities |
| --- | --- |
| redis | Cluster/Sentinel, typed Streams commands and sharded subscriptions; no automatic retry of ambiguous writes; active standard TLS I/O interruption requires reconnection |
| lsp | Broader typed feature models (completion, code actions, workspace edits and semantic tokens); the application event loop drives deadline polling |
| cli | Shell completion and defaults/environment for flags and counters |
| diff | Multi-file and Git metadata support, three-way merge, offset/fuzzy application; the linear-space algorithm is explicitly selected and has O(NM) worst-case time |
| parser | Recovery with multiple diagnostics, token-stream and incremental text parsing, further binary/text combinator parity; grammar left recursion still requires rewriting |
| template | Macros/imports/call blocks, keyword arguments, file-loader invalidation and incremental output writing |
| markdown | GFM extensions and finer inline source spans; existing CommonMark behavior must remain covered |
| sqlite | Row derives, batch helpers, connection/statement caching, custom functions, backup and incremental blob APIs |
| pipeline | Error recovery/retry, time-based operators, parallel flat-map and metrics |
| ndarray | Masked selection/scatter, sorting/quantiles, NPY interchange, SVD/eigen and rank-deficient solve support |
| msgpack | Incremental field processing, reduced materialization and richer typed extension support; standard I/O adapters buffer one bounded value at a time |
| graph | Flow/matching algorithms, serialization and configurable cost types |
| goml_stats | Manifest-based canonical identities, declaration counts and historical comparisons; hierarchical Git ignore rules are provided by the ignore dependency |
| bitflags | Associated-constant or operator syntax depends on language support; Serde wrappers support explicit or format-sensitive representations; arbitrary declaration expressions and generic storage newtypes are not generated |
| logos | Compile-time derive/DFA generation, streaming/byte input, named subpatterns and broader Unicode regex properties; current runtime NFA reports equal-priority ambiguity during matching |
| tempfile | Platforms beyond Linux amd64, cancellation-aware file operations, crash-durable persistence helpers; cleanup assumes no hostile concurrent filesystem changes |
| reqwest | Streaming requests/responses, HTTP/3, an adapter to the separate websocket library, full domain-cookie policy, custom DNS, application retries and middleware; the synchronous API buffers within explicit limits |
| llvm | JIT execution, cross-target configuration, debug metadata, atomics, exception handling and broader LLVM instruction/API coverage; the initial binding targets LLVM 18 and native object generation |
| incremental | Parallel branch evaluation, immutable database snapshots, persistent caches, durability classes and cycle fixed-point recovery; current root operations serialize and callbacks use scoped Evaluation handles |
| rope | Grapheme and reverse iterators, search, editing history, optional Unicode newline policies and memory-mapped backing; current storage is persistent UTF-8 with LF/CRLF/CR line semantics |
| web | TLS listeners, HTTP/2 and HTTP/3, an upgrade adapter to the separate websocket library, multipart extraction, static files, compression and bundled CORS middleware; current adapter serves HTTP/1.1 with streaming and SSE |
| bigint | Faster multiplication/division for very large operands, primality and modular inverses, roots and rational arithmetic; ordinary arithmetic allocates proportionally to results while input/shift/power operations have explicit budgets |
| tracing | Distributed trace propagation, OpenTelemetry exporters, richer sampling, byte-budget admission and instrumentation syntax; contexts are explicit and arbitrary sink callbacks must cooperate with shutdown |
| datetime | Arbitrary-pattern parsing, localization, recurrence scheduling and automatic timezone-data updates; dates cover Gregorian years 1–9999 and timestamps use POSIX seconds without leap records |
| decimal | Roots and transcendental functions, locale formatting, binary-float conversion and special-value/trap models; current finite arithmetic has explicit 4,096-digit coefficient/precision and scale bounds |
| cache | Frequency-based admission, sharding and indexed/background expiry; current exact LRU uses O(n) expiry scans when timed entries exist, and synchronous loaders cooperate with cancellation |
| ignore | Combined multi-pattern automata, tracked-file/index-aware selection, configurable file-type groups and additional platforms; current matching is byte-oriented with explicit work budgets and traversal targets Linux amd64 |
| syntax | Incremental parsing/reparse orchestration, syntax pointers stable across revisions, multi-edit transactions and weak-reference interning; current library provides immutable lossless trees and checked persistent edits |
| color | CSS Color 4's full grammar, additional RGB profiles, chromatic adaptation, HDR and ICC; Lab currently uses D65 and the parser documents its subset |
| unicode_text | Unicode version upgrades, locale/dictionary tailoring, normalization, bidi shaping and sentence segmentation; current tables are pinned to Unicode 16 |
| ansi | Screen emulation, single-byte C1 mode, extended underline styles/colors and terminal-specific palette discovery |
| terminal | Additional operating systems, portable signal subscriptions and suspend/resume, Kitty keyboard/modifyOtherKeys, broader terminfo negotiation; current resize detection uses bounded polling |
| tui | Configurable font/terminal width policies, soft-wrapped persistent editing, system clipboard, widget mouse-hit routing and graphics protocols |
| prompt | Ranked completion-menu integration with the separate fuzzy library, date/file pickers and batch-mode policies; current callbacks are synchronous and password values are ordinary GC strings |
| progress | Byte-stream adapters, recursive job trees, pause/resume accounting and arbitrary format templates; applications explicitly drive ticks |
| diagnostics | Bidi/font shaping, richer graphical label routing and persistent source revisions; edit application returns checked new text without writing files |
| tui_markdown | Full GFM extensions and integration with the separate highlight library; terminal rendering and optional pipe tables do not change the CommonMark parser's scope |
| fuzzy | Canonical normalization, accent/transliteration policies, richer query expressions and faster very-large Top-K; scores are library-specific and DP matrix limits are explicit |
| config | Interpolation, additional format/secret providers, debounce/background scheduling and line/column provenance; file reload targets Linux and TOML inherits the standard parser's scope |
| websocket | Compression extensions, HTTP/2 CONNECT, proxy/redirect integrations and Autobahn certification; the core currently implements RFC 6455 without extensions |
| highlight | Semantic scopes, full Markdown inline/block semantics, interpolated-expression scopes and TextMate/Sublime grammar compatibility; source copying remains linear despite lexical suffix reuse |
| archive | ZIP64, additional compression/encryption, GNU sparse/base-256/longname extensions and fully streaming decompression; ZIP indexing and GZIP convenience APIs are bounded in-memory operations |
| metrics | OpenMetrics/native histograms, distributed exporters, automatic runtime instrumentation and sharded update paths; registry operations currently serialize for consistent snapshots |
| csv | Asynchronous I/O, indexing/seeking, additional escape dialects and nested Serde structures; typed flat rows use an explicit scalar schema and stream operations are synchronous |
| bench | Allocation/CPU counters, plotting dashboards, process isolation and automated baseline selection; timings include Go runtime effects, cancellation is cooperative and input barriers support scalar values |

Further work should preserve resource bounds, recoverable errors, normal
versioned dependency consumption and the independent reference checks already
present. Bundled networking adapters and concurrent resource pools need actual
I/O and race coverage, beyond a public interface or a mock implementation.
