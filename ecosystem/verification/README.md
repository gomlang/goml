# Native ecosystem verification

This standalone GoML module owns the local ecosystem test runner, atomic private
registry snapshots, race and SIMD checks, and Linux PTY sessions. The
`reference` package supplies test fixture decoding, structural JSON comparisons
and an independent VT screen model. It depends only on the standard library.

From the repository root:

```sh
just ecosystem-test
just ecosystem-test color ndarray goml_stats
just ecosystem-test --no-race terminal explorer
just ecosystem-test --goml /path/to/goml lsp
just ecosystem-test --list
```

No module arguments selects all libraries registered in `modules()` and their separate consumers,
`goml_stats`, Explorer and four compiler regression projects. Each selected
module must exist. The runner checks formatting, builds consumers before tests,
runs library and consumer `#[test]` suites, verifies cached build fingerprints
and executes smoke checks. Native PTY checks cover terminal, tui, prompt,
progress, tui_markdown and Explorer. The ndarray check also compiles SSE2 and
scalar variants, inspects the linked kernel symbols and repeats reference tests.
Unicode tools freshly download checksum-pinned inputs on every conformance run.

Modules with concurrency checks run their suites again using `GOFLAGS=-race`
and separate `_artifact/race/` targets. This includes declared native Go adapter
tests where applicable. `--no-race` explicitly skips this additional pass.
Individual test deadlines are 300 seconds; command deadlines are 600 seconds.
Failures write captured output and return a nonzero status. The final JSON
report records success, selected modules, registry identity, command durations,
exit codes and log paths under `ecosystem/_artifact/verification/`.
`GOML_VERIFY_REPORT` can select a different JSON report path when running disjoint
module batches concurrently.

Reference fixtures retain results from independent implementations, along with
their provenance. Native tests consume those fixed expectations; they do not
require Python, NumPy or Rust at test time. Actual filesystem, process, socket,
TLS, PTY and external LLVM/GNU tool checks continue to execute. See each module's
README for prerequisites.

To run an individual consumer manually, first build this runner and obtain its
private registry path:

```sh
cd ecosystem/verification
../../stage2/bin/goml build
export GOML_HOME="$(_artifact/bin/verification --registry-only)"
cd ../consumers/color
../../../stage2/bin/goml build
../../../stage2/bin/goml test
```

The runner supplies `GOML_VERIFY_ROOT`, `GOML_VERIFY_DRIVER` and
`GOML_VERIFY_BINARY` to test subprocesses. Native test helpers use local build
paths when these variables are absent. Shared consumer dependencies are resolved
from a content-addressed snapshot, whose files are captured once and published
with an atomic directory rename. Concurrent publishers can only observe a
complete index; generated directories and symbolic links are excluded.

The runner's own regression suite covers publication races, captured source
consistency, coordinate validation, binary data, symlink exclusion, command
failure logs, path delimiters, UTF-8 failures and option parsing. The reference
model tests also check fragmented terminal sequences, erased password history
and cleanup when starting a PTY child fails:

```sh
cd ecosystem/verification
../../stage2/bin/goml test
```

These commands are local development tools; no ecosystem CI job is installed.
