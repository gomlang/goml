# C2 callback acceptance

Scope: panic propagation, pruning, declaration provenance, reproducible generated output, packaged-toolchain execution and measured callback costs. This audit does not cover deferred map constraints or interface adapters.

## Panic and reachability

The driver regression `callback_panics_preserve_identity_unwind_and_survive_pruning` checks original Go panic object identity through Go/GoML calls and a native function converted into a closure and back. Explicit recovery belongs to the Go fixture. Deferred work runs during unwinding, unit and multi-result callbacks do not fabricate successful results, and a nil-map write retains its native runtime panic. Separate processes verify unhandled synchronous and asynchronous failure. The asynchronous fixture releases its waiting caller only on normal callback return, avoiding a race between process exit and fatal panic reporting.

Captured state and a private helper from another package remain reachable exclusively through a callback after pruning. An unrelated helper is absent from emitted Go. The adapter uses ordinary Go reachability, with no registration roots, reflection dispatch or implicit recovery. Fresh and cached builds preserve execution and generated source.

## Declaration provenance

`project_generated_callback_origins_match_go_declarations_and_sources` parses executable, internal test, black-box test and exported library output with Go's AST parser. Every indexed generated symbol must exist. Source entries identify their canonical declaration, defining package, package-relative file and declaration-name byte range and Unicode position. The fixture validates generic nested closures, distinct captured functions, three same-named inherent/trait methods, shared adapters and public library wrappers. Library mappings use the actual collision-resolved private names.

Foreign callable identities are carried explicitly through specialization and artifact codecs. Two used bindings sharing a target and signature share a wrapper with both origins; an unused alias does not acquire an origin. Different ordinary/raw string signatures remain distinct, and receiver methods retain their binding identity. The compiler codec test `foreign_callable_origins_accept_legacy_bodies_and_reject_invalid_ids` covers absent legacy IDs and malformed values. Persisted source-origin tests cover missing source files and reject malformed positions and fabricated declaration identities.

The unexported `_goml_source_origins` constant has a versioned JSON payload and `enclosing-declaration` scope. Its positions refer to declaration names, including binding and export declarations. It does not rewrite runtime stack traces or describe individual expressions and inlining history. Synthetic methods without source spans are omitted. The index is embedded after pruning and follows existing generated-source cache and transactional output paths. Shared adapters may legitimately have multiple declaration entries.

## Packaging, determinism and measurement

The origin regression compares fresh/cached executable and test output, repeats library export and invokes an exported callback-using function from a normal Go host. Extracted-release smoke builds and runs the committed subscription example twice with required FFI validation and verifies generated-source digests. The example also runs under the Go race detector in driver tests. Full CI checks stage0 compilation, stage2/stage3 fixed point, golden output and release packaging.

The [benchmark](../../examples/ffi-callback-bench/README.md) separates invocation from escaping construction, records Go version/platform and uses standard Go benchmark calibration. Its recorded local run measured no allocations in invocation and two allocations for escaping GoML callback construction. These observations are not portable limits or a zero-allocation guarantee. Performance is not used as a timing threshold in CI.

## Verification status

The panic, source-provenance, test-link, library-link and method-origin changes each passed full CI; their logs and counts are recorded in [the implementation log](implementation-plan.md). The direct-extern regression passed its focused driver run. Latest-source `just ci` passed with exit status 0, including 850 compiler tests, 129 driver tests, helper race checks, golden verification, stage2/stage3 fixed point and extracted-release smoke (`/tmp/goml-extern-origins-ci.log`, `/tmp/goml-extern-origins-ci.status`). This establishes C2 acceptance for the declaration-level provenance and callback contract described above. B3 deferred constraints, B4 acceptance and batch D remain separately tracked.
