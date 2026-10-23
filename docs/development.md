# Developing GoML

[English](development.md) | [简体中文](development.zh-CN.md)

Run recipes from the repository root. The toolchain build requires Linux amd64 or macOS arm64, Go 1.26+, `just`, Bash, curl, tar, and sha256sum or shasum. Tests require a C compiler for Go's race detector; full CI also uses Node 20+, npm, jq, Python 3, Perl's `shasum` for checksum portability tests, and Linux development headers for checking the generated syscall ABI and record layouts. See [.justfile](../.justfile) for all commands. The independent [gomlgo frontend and interpreter](https://github.com/gomlang/gomlgo) targets Go 1.26 and requires Go 1.26.x for execution and differential tests.

```sh
just make
just test
just ci
just clean
```

`just make` incrementally builds stage2 directly from the pinned stage0. `just test` builds the tools and runs compiler, driver, and standard-library tests; `just all` is an alias for it. `just ci` performs a clean stage2 build, fixed-point verification, tests, extension compilation, and release archive smoke checks. Standard-library behavior tests live beside their implementations in `lib/std/<package>/tests`; run them with `just test-stdlib`.

Local CI runs its check groups concurrently. Set `GOML_CI_SEQUENTIAL=1` to run them sequentially, `GOML_BUILD_JOBS` to limit bootstrap package workers, and `GOML_TEST_JOBS` to override compiler and driver test concurrency. Compiler fixtures may use Yaegi when it is available on `PATH`; they fall back to native Go compilation when it is unavailable or cannot run a fixture.

GitHub CI builds test artifacts once and runs them in shards with `goml test --no-build`. See [the workflow](../.github/workflows/ci.yml) for job dependencies and [test commands](tooling.md#check-and-run-tests) for sharding and artifact reuse.

The bootstrap downloads the checksum-pinned stage0 release recorded in [bootstrap/stage0.env](../bootstrap/stage0.env). `just bootstrap` rebuilds stage2 from stage0, builds stage3 with stage2, then uses stage3 to rebuild the compiler and driver artifacts and compares them with the first stage3 build. Set `GOML_STAGE0_ARCHIVE` to a previously downloaded pinned archive to avoid downloading stage0.

Source bootstrap selects the published `linux-amd64` or `darwin-arm64` archive for the host and verifies its separately pinned SHA-256 checksum. Both platforms use the same stage0 release version. An offline `GOML_STAGE0_ARCHIVE` must match the host's pinned archive. Unsupported hosts are rejected before downloading stage0 or cleaning existing bootstrap outputs.

On macOS arm64, run `just bootstrap` for a clean native build and fixed-point verification, or `just make` for an incremental build. Run `just _ci-toolchain-scripts` and `just _ci-release-smoke` after building to check installation, caching, packaging, and the published-toolchain workflows. Full `just ci` and the complete compiler and standard-library test suites require Linux amd64 because they exercise Linux syscalls and ABI headers. GitHub CI also runs the native macOS bootstrap and toolchain script checks alongside the macOS release archive smoke tests.

Use `stage2/bin` for local development. Toolchain prefixes under `stage0`, `stage2`, and `stage3` contain `bin`, library projects under `lib`, and the finalized compiler world under `lib/compiler`. Downloaded archives are cached in `_bootstrap/cache`; compiler and driver bootstrap products live in their module-local `_bootstrap` directories. Generated outputs are ignored by Git.

The root `goml.work` groups `gomlc` and `goml`. The driver owns manifest parsing, workspace selection, registry indexing, and dependency resolution under `goml/manifest` and `goml/registry`. The compiler's query engine consumes project descriptions from `goml`; it does not resolve manifests or registry dependencies itself. Run `stage2/bin/goml fmt --check --workspace` to check workspace formatting. Library modules under `lib/` are formatted and tested separately.

`just install` installs and finalizes the tools under `${GOML_HOME:-$HOME/.goml}`; add its `bin` directory to `PATH`. `just clean` removes the root and compiler/driver build caches and generated development stages, while retaining the downloaded `stage0` toolchain.

Explore [gomlc/testdata/pipeline](../gomlc/testdata/pipeline) for source programs and every compiler-stage golden file. Use `just verify-golden` to check the corpus or `just update-golden` to regenerate it through the self-hosted compiler.

Direct Go FFI integration fixtures live under [goml/testdata/ffi](../goml/testdata/ffi) and are shared by driver regression tests and release smoke checks. Full SQLite and LLVM integrations are maintained in [gomlang/sqlite](https://github.com/gomlang/sqlite) and [gomlang/llvm](https://github.com/gomlang/llvm).

See the [compiler](../gomlc/README.md) and [project driver](../goml/README.md) documentation for component-specific commands and architecture.

## Bilingual documentation

English documentation remains canonical at its existing paths. Every repository Markdown document has a complete Simplified Chinese translation beside it, named with the `.zh-CN.md` suffix. Keep both versions up to date in the same change when adding or changing documentation, language semantics, public APIs, or development workflows. `AGENTS.md` remains the authoritative instruction file; its Chinese counterpart is a reading reference.

Preserve code examples, commands, API identifiers, and diagnostic/output text in both languages. Use consistent terminology: package (包), module (模块), monomorphization (单态化), closure lifting (闭包提升), and compile-time evaluation (编译期求值); retain `trait` and IR names where they identify language or compiler concepts. Translate explanatory prose fully rather than replacing sections with summaries.

The [language guide](goml.md) is an overview and reading map. Maintain language rules in the linked `language/` topics, project workflows in [project tools](tooling.md), and Go boundary rules in [Go interoperability](ffi/bind-go.md). Keep the overview's original section anchors as links to the canonical topic locations when reorganizing content; link directly to those locations from other documents.

Each document includes an `English | 简体中文` switch. Chinese links point to Chinese counterparts where available. Preserve English heading fragments in Chinese documents with explicit anchors so existing cross-references remain valid. When changing headings, update links and anchors together.

Run `just check-docs` (or `python3 tools/check-docs.py`) from the repository root to check bilingual coverage, language switches, heading structure, code examples, and local links and anchors. These checks detect structural drift; reviewers still need to check translation accuracy and completeness. Documentation-only changes require the applicable documentation checks, not the full compiler CI suite.
