# goml

goml is a statically typed programming language inspired by Go and Rust.

The "ml" in goml nods to the [ML (programming language)](https://en.wikipedia.org/wiki/ML_(programming_language)), whose descendants have deeply influenced Rust.

goml aims to empower gophers with a more powerful type system but without leaving the Go ecosystem.

GoML is statically typed and garbage-collected, with Rust-like syntax, monomorphized generics, and no ownership or lifetime system. The compiler, project driver, tests, formatter, and language server are implemented in GoML and compile to Go. Generated Go code and Go interoperability target Go 1.26.

Toolchain releases provide Linux amd64 and macOS arm64 archives. See [installation and platform support](docs/releasing.md#installation) for checksum verification, required installation finalization, and platform-specific library limitations. macOS archives are tested on macOS 15 with Apple Silicon; source bootstrap still starts on Linux amd64.

## Documentation

- [Language guide](docs/goml.md): syntax, semantics, packages, tests, and standard-library APIs
- [Library boundaries](docs/library-boundaries.md): standard-library and ecosystem ownership
- [Formatting](docs/formatting.md): formatter rules and CLI
- [Compiler](gomlc/README.md) and [project driver](goml/README.md): local tools and development commands
- [Compile-time evaluation](docs/comptime.md): CTIR and programmable derive architecture
- [Go bindings](docs/ffi/bind-go.md) and [metadata protocol](docs/ffi/protocol-v1.md): Go interoperability
- [gomlgo](https://github.com/gomlang/gomlgo): independent Go frontend, interpreter, and differential tests
- [Ecosystem libraries](../gomlang/ecosystem/README.md): independent sibling repositories, consumers, examples, and verification
- [VS Code extension](editors/vscode/README.md): editor setup and commands
- [Releasing](docs/releasing.md): release archives, installation, and stage0 advancement

## Development

Run recipes from the repository root. The toolchain build requires Linux amd64, Go 1.26+, `just`, Bash, curl, tar, and sha256sum. Tests require a C compiler for Go's race detector; full CI also uses Node 20+, npm, jq, Python 3, Clang 15+ for C binding smoke tests, Perl's `shasum` for checksum portability tests, and Linux development headers for checking the generated syscall ABI and record layouts. See [.justfile](.justfile) for all commands. The independent [gomlgo frontend and interpreter](https://github.com/gomlang/gomlgo) targets Go 1.26 and requires Go 1.26.x for execution and differential tests.

```sh
just make
just test
just ci
just clean
```

`just make` incrementally builds stage2 directly from the pinned stage0. `just test` builds the tools and runs compiler, driver, and Go metadata tests; `just all` is an alias for it. `just ci` performs a clean stage2 build, fixed-point verification, tests, extension compilation, and release archive smoke checks. The independent gomlgo suite is maintained in [its own repository](https://github.com/gomlang/gomlgo).

Local CI runs its check groups concurrently. Set `GOML_CI_SEQUENTIAL=1` to run them sequentially, `GOML_BUILD_JOBS` to limit bootstrap package workers, and `GOML_TEST_JOBS` to override compiler and driver test concurrency. Compiler fixtures may use Yaegi when it is available on `PATH`; they fall back to native Go compilation when it is unavailable or cannot run a fixture.

GitHub CI distributes compiler tests across four runners using `GOML_TEST_SHARD=1/4` through `4/4`, with four test workers on each runner. Each shard downloads the shared stage2 build and builds its test runner locally. The final `test` check requires every shard and the other CI jobs to pass.

The bootstrap downloads the checksum-pinned stage0 release recorded in [bootstrap/stage0.env](bootstrap/stage0.env). `just bootstrap` rebuilds stage2 from stage0, builds stage3 with stage2, then uses stage3 to rebuild the compiler and driver artifacts and compares them with the first stage3 build. Set `GOML_STAGE0_ARCHIVE` to a previously downloaded pinned archive to avoid downloading stage0.

Use `stage2/bin` for local development. Toolchain prefixes under `stage0`, `stage2`, and `stage3` contain `bin`, library projects under `lib`, and the finalized compiler world under `lib/compiler`. Downloaded archives are cached in `_bootstrap/cache`; compiler and driver bootstrap products live in their module-local `_bootstrap` directories. Generated outputs are ignored by Git.

The root `goml.work` groups `gomlc`, `goml`, and the internal `toolchain/manifest` module. The compiler and driver share manifest parsing, version selection, and dependency graph resolution through explicit path dependencies on that module. Each consumer supplies its own manifest validation and error policy. It is compiled into the tools; installed toolchains do not need the repository sources. Run `stage2/bin/goml fmt --check --workspace` to check workspace formatting, or run `goml test` from `toolchain/manifest` to test the shared parser and resolver. Library modules under `lib/` are formatted and tested separately.

`just install` installs and finalizes the tools under `${GOML_HOME:-$HOME/.goml}`; add its `bin` directory to `PATH`. `just clean` removes the root, compiler/driver, and shared manifest build caches and generated development stages, while retaining the downloaded `stage0` toolchain.

Explore [gomlc/testdata/pipeline](gomlc/testdata/pipeline) for source programs and every compiler-stage golden file. Use `just verify-golden` to check the corpus or `just update-golden` to regenerate it through the self-hosted compiler.

## Disclaimer

This project is a **personal project** and is **NOT** affiliated with, endorsed by, or connected to any organization.

⚠️Do not use this project or any of its derivatives in production environments.

The author assumes no responsibility for any risk or damage resulting from the use of this project.

## Ownership and License

This project does not currently have an open-source license. Until a license is explicitly provided, all rights to the project — including code, documentation, design, and related resources — are reserved by the author. No copying, distribution, modification, or commercial use is permitted without prior authorization.
