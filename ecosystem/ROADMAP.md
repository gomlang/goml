# Functional backlog

The modules are usable within their documented scope. This list tracks the
remaining gaps from the ecosystem audit; an implemented module is not a claim
of feature parity with every mature library in that category.

## First improvement batch

| Module | Added |
| --- | --- |
| proptest | Lazy shrinking with shared work limits, structured pass/fail/discard reports, classification and coverage, persistent regression seeds, unique collections, model-valid command sequences and full-width numeric edge cases |
| redis | Injectable duplex transports and dialing with a shared setup/operation deadline, adapter validation and failure cleanup |
| lsp | Deferred request acceptance/completion, cooperative cancellation, deadline polling and pending-request cleanup on exit |
| cli | Argument/command aliases, inherited global options, cardinality-constrained argument groups and derive support for option aliases/globals |
| diff | Explicit linear-space Hirschberg algorithm for sequences, text and patches, with work/workspace limits |
| parser | Shared text/binary work and depth limits, contextual custom parsers, iterative alternatives and binary backtracking/lookahead primitives |

Each addition has library and independently resolved consumer coverage. The
verification entry point remains `python3 ecosystem/verify.py`; CI integration
is intentionally deferred.

## Remaining work

| Module | Remaining capabilities |
| --- | --- |
| proptest | Distribution histograms and richer report formatting, failure aggregation, application-oriented state-machine execution helpers; custom callbacks remain responsible for their own bounded work |
| redis | Bundled DNS/TLS adapters, connection pools and health/reconnection policies, Cluster/Sentinel, typed Streams commands and sharded subscriptions; no automatic retry of ambiguous writes |
| lsp | Position-encoding negotiation, broader typed feature models (completion, code actions, workspace edits and semantic tokens), outgoing deadline scheduling and more efficient document storage |
| cli | Nested/flattened argument derives, derived subcommands, shell completion, defaults/environment for flags and counters |
| diff | Multi-file and Git metadata support, three-way merge, offset/fuzzy application; the linear-space algorithm is explicitly selected and has O(NM) worst-case time |
| parser | Recovery with multiple diagnostics, token-stream and incremental text parsing, further binary/text combinator parity; grammar left recursion still requires rewriting |
| template | Macros/imports/call blocks, keyword arguments, file-loader invalidation and incremental output writing |
| markdown | GFM extensions and finer inline source spans; existing CommonMark behavior must remain covered |
| sqlite | Row derives, batch helpers, connection/statement caching, custom functions, backup and incremental blob APIs |
| pipeline | Error recovery/retry, time-based operators, parallel flat-map and metrics |
| ndarray | Masked selection/scatter, sorting/quantiles, NPY interchange, SVD/eigen and rank-deficient solve support |
| msgpack | Reader/writer integration and incremental field processing, reduced copying and richer typed extension support |
| graph | Flow/matching algorithms, serialization and configurable cost types |
| goml_stats | Gitignore/glob semantics, manifest-based canonical identities, declaration counts and historical comparisons |
| bitflags | Associated-constant or operator syntax depends on language support; wire representation uses explicit Serde wrappers; arbitrary declaration expressions and generic storage newtypes are not generated |
| logos | Compile-time derive/DFA generation, streaming/byte input, named subpatterns and broader Unicode regex properties; current runtime NFA reports equal-priority ambiguity during matching |
| tempfile | Platforms beyond Linux amd64, cancellation-aware file operations, crash-durable persistence helpers; cleanup assumes no hostile concurrent filesystem changes |
| reqwest | Streaming requests/responses, HTTP/3, WebSocket, full domain-cookie policy, custom DNS, application retries and middleware; the synchronous API buffers within explicit limits |
| llvm | JIT execution, cross-target configuration, debug metadata, atomics, exception handling and broader LLVM instruction/API coverage; the initial binding targets LLVM 18 and native object generation |

Further work should preserve resource bounds, recoverable errors, normal
versioned dependency consumption and the independent reference checks already
present. Bundled networking adapters and concurrent resource pools need actual
I/O and race coverage, beyond a public interface or a mock implementation.
