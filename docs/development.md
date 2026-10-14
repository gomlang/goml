# Developing GoML

Run recipes from the repository root. The toolchain build requires Linux amd64, Go 1.26+, `just`, Bash, curl, tar, and sha256sum. Tests require a C compiler for Go's race detector; full CI also uses Node 20+, npm, jq, Python 3, Clang 15+ for C binding smoke tests, Perl's `shasum` for checksum portability tests, and Linux development headers for checking the generated syscall ABI and record layouts. See [.justfile](../.justfile) for all commands. The independent [gomlgo frontend and interpreter](https://github.com/gomlang/gomlgo) targets Go 1.26 and requires Go 1.26.x for execution and differential tests.

```sh
just make
just test
just ci
just clean
```

`just make` incrementally builds stage2 directly from the pinned stage0. `just test` builds the tools and runs compiler, driver, standard-library, and Go metadata tests; `just all` is an alias for it. `just ci` performs a clean stage2 build, fixed-point verification, tests, extension compilation, and release archive smoke checks. Standard-library behavior tests live beside their implementations in `lib/std/<package>/tests`; run them with `just test-stdlib`.

Local CI runs its check groups concurrently. Set `GOML_CI_SEQUENTIAL=1` to run them sequentially, `GOML_BUILD_JOBS` to limit bootstrap package workers, and `GOML_TEST_JOBS` to override compiler and driver test concurrency. Compiler fixtures may use Yaegi when it is available on `PATH`; they fall back to native Go compilation when it is unavailable or cannot run a fixture.

GitHub CI builds test artifacts once and runs them in shards with `goml test --no-build`. See [the workflow](../.github/workflows/ci.yml) for job dependencies and [test commands](goml.md#check-and-run-tests) for sharding and artifact reuse.

The bootstrap downloads the checksum-pinned stage0 release recorded in [bootstrap/stage0.env](../bootstrap/stage0.env). `just bootstrap` rebuilds stage2 from stage0, builds stage3 with stage2, then uses stage3 to rebuild the compiler and driver artifacts and compares them with the first stage3 build. Set `GOML_STAGE0_ARCHIVE` to a previously downloaded pinned archive to avoid downloading stage0.

Use `stage2/bin` for local development. Toolchain prefixes under `stage0`, `stage2`, and `stage3` contain `bin`, library projects under `lib`, and the finalized compiler world under `lib/compiler`. Downloaded archives are cached in `_bootstrap/cache`; compiler and driver bootstrap products live in their module-local `_bootstrap` directories. Generated outputs are ignored by Git.

The root `goml.work` groups `gomlc` and `goml`. The driver owns manifest parsing, workspace selection, registry indexing, and dependency resolution under `goml/manifest` and `goml/registry`. The compiler's query engine consumes project descriptions from `goml`; it does not resolve manifests or registry dependencies itself. Run `stage2/bin/goml fmt --check --workspace` to check workspace formatting. Library modules under `lib/` are formatted and tested separately.

`just install` installs and finalizes the tools under `${GOML_HOME:-$HOME/.goml}`; add its `bin` directory to `PATH`. `just clean` removes the root and compiler/driver build caches and generated development stages, while retaining the downloaded `stage0` toolchain.

Explore [gomlc/testdata/pipeline](../gomlc/testdata/pipeline) for source programs and every compiler-stage golden file. Use `just verify-golden` to check the corpus or `just update-golden` to regenerate it through the self-hosted compiler.

Basic FFI integration fixtures live under [goml/testdata/ffi](../goml/testdata/ffi) and are shared by driver regression tests and release smoke checks. Full SQLite and LLVM integrations are maintained in [gomlang/sqlite](https://github.com/gomlang/sqlite) and [gomlang/llvm](https://github.com/gomlang/llvm).

See the [compiler](../gomlc/README.md) and [project driver](../goml/README.md) documentation for component-specific commands and architecture.
