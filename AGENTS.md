# Repository Guidelines

goml is a statically typed, garbage-collected language with Rust-like syntax that compiles to Go. Sources use `.gom`; there is no ownership system or lifetime syntax. The compiler monomorphizes generics and lambda-lifts GoML closures.

## Documentation and Source Map

Read the documentation relevant to the change:

- [Language guide](docs/goml.md): canonical syntax, semantics, packages, builtins, standard-library APIs, examples, and grammar.
- [Formatting](docs/formatting.md): formatter behavior and CLI.
- [Releases](docs/releasing.md): publishing, toolchain installation, and stage0 advancement.
- [Compile-time evaluation](docs/comptime.md), [Go FFI](docs/ffi/bind-go.md), and [FFI protocol](docs/ffi/protocol-v1.md): specialized workflows and the Go metadata boundary.
- [gomlgo](gomlgo/README.md): independent Go frontend, interpreter, compatibility target, and differential tests.
- [VS Code extension](editors/vscode/README.md): editor setup, behavior, and configuration.

| Path | Responsibility |
| --- | --- |
| `gomlc/` | Self-hosted compiler, query engine, LSP, resource loader, and compiler tests |
| `goml/` | Self-hosted project driver, registry client, dependency resolver, and CLI tests |
| `gomlgo/` | Independent Go frontend and interpreter implemented in GoML |
| `lib/builtin/` | Hidden compiler/runtime contract, runtime hooks, and language items |
| `lib/prelude/` | Independent package defining the automatically scoped public API |
| `lib/std/` | Standard-library packages loaded through explicit imports |
| `bootstrap/` | Bootstrap scripts and checksum-pinned released stage0 metadata |
| `tools/` | Build, library packaging, release, and Go metadata tooling |
| `editors/vscode/` | VS Code extension |
| `gomlc/testdata/` | Compiler regression fixtures and generated golden files |

Generated toolchains live under `stage0` through `stage3`; use `stage2/bin` for local development. Each loads resources from its executable-relative `lib/`, including its finalized compiler world under `lib/compiler/`. Build outputs belong under `_bootstrap/`, `_artifact/`, or module-configured target directories.

## Development Workflow

Requirements: Linux amd64, Go 1.25+, Node 20+, npm, `just`, Bash, curl, tar, and sha256sum. See the gomlgo README for its additional Go-version requirements.

Run recipes from the repository root; [.justfile](.justfile) is the command reference.

| Command | Purpose |
| --- | --- |
| `just make` | Incrementally build stage2 from pinned stage0 |
| `just test` / `just all` | Build and run compiler, driver, and Go metadata tests |
| `just ci` | Required full CI, including bootstrap fixed point and packaging |
| `just bootstrap` | Clean bootstrap and fixed-point verification |
| `just verify-golden` / `just update-golden` | Verify / regenerate snapshots through self-hosted tests |
| `just vscode-ext` | Build the LSP and compile the extension |
| `just gomlgo-test` | Run the independent gomlgo test suite |
| `just clean` | Remove local build caches and generated toolchains |

- After editing `.gom` files, run `goml fmt` from every affected module before tests or commits. Modules include `gomlc/`, `goml/`, `gomlgo/`, and the separate library projects under `lib/`.
- Use the repository formatter, for example `cd gomlc && ../stage2/bin/goml fmt`; `fmt --check` verifies formatting.
- Run a focused fixture with `stage2/bin/gomlc run-single <file.gom>`. Add `--dump-ast`, `--dump-hir`, `--dump-tast`, `--dump-core`, `--dump-mono`, `--dump-lift`, `--dump-anf`, or `--dump-go` to inspect lowering.
- `goml check`, `goml build`, and `goml test` discover the enclosing `goml.toml` and operate on the complete module, without package targets. `--dry-run` prints planned commands.
- The driver finds `gomlc` through `--compiler`, `GOMLC`, a sibling binary, `GOML_HOME/bin`, then `PATH`, and verifies the driver protocol.
- Run relevant tests and `just ci` before reporting completion. Changes to `gomlgo/` also need its separate tests; consult its README for focused differential checks.

## Coding and Architecture Rules

- Do not add code comments; write clear, self-explanatory code.
- GoML uses four-space indentation, snake_case functions/packages, CamelCase types, and explicit top-level function signatures. Generics use square brackets; local closures have one concrete type and cannot declare generics or use let-generalization.
- TypeScript uses two-space indentation and PascalCase components; prefer named exports.
- Keep packages small and focused. Prefer method syntax when available, such as `s.len()`, `s.get(i)`, and `x.to_string()`.
- The frontend pipeline is `lexer → parser → CST → AST → HIR → TAST → Core → Mono → Lift → ANF → Go`.
- `gomlc/hir/` owns AST-to-HIR lowering and name resolution. `gomlc/tast/` owns inference, checking, constraints, and type-directed decisions, including contextual enum-pattern resolution.
- Ambiguity, missing lookups, invalid caches, and dependency failures must produce recoverable diagnostics. Environment and lookup code must return failure values instead of terminating the compiler.
- `lib/builtin/` alone owns compiler runtime externs and language items. Keep this contract distinct from ordinary user Go FFI; see the language guide and FFI docs for supported bindings.
- Reuse `gomlc/query/` for LSP features. Preserve file-scoped imports, canonical package identities, dependency navigation, and existing document-analysis caching.
- ANF join points are local continuations reached in tail position through `Jump`; recursive joins represent loops. Preserve lexical scope and jump arguments. Consult [ANF definitions](gomlc/anf/model.gom), [verification](gomlc/anf/verify.gom), and [Go lowering](gomlc/go_backend/lower.gom) when changing control flow.
- Go lowering emits structured control flow without `goto`; labeled loop breaks/continues are supported. Preserve continuation merging, recursive-loop handling, and the post-emission [Go DCE pass](gomlc/go_backend/dce.gom).

## Packages and Dependencies

- A module owns the canonical path in `[module].path` in its root `goml.toml`; directory paths determine package identities. Project sources declare their package, and files in one directory share a package name. Standalone single-file commands may omit the declaration.
- Imports are file-scoped. The declared package name supplies the default alias; `as` disambiguates it. `module::path` resolves from the module root, and `use alias::Trait` brings an imported trait into method-call scope.
- Cross-package APIs require `pub`. Struct fields and inherent methods are private by default; trait implementation methods inherit trait visibility and must not declare `pub`.
- Interfaces and dependency environments expose public API while retaining metadata needed to use it. Current-package codegen must retain private helpers and the full internal environment.
- Executables come from `package main` packages with `fn main()`. A direct `tests/` directory is one black-box test package for its parent and declares `package tests`; nested test suites are unsupported.
- Dependencies belong only in the module-root manifest. Registry resolution reads the authoritative `index.toml`; published versions are immutable. Strict `X.Y.Z` requirements are minimums resolved through MVS, and there is no `goml.lock`.
- Global state defaults to `~/.goml`, overridden by `GOML_HOME`. Dependency sources stay in its `cache/registry`; generated artifacts go under the project's configured target directory, defaulting to `_artifact/`.

## Tests and Golden Files

- Prefer fast, deterministic tests and minimal fixtures covering relevant parsing, typing, and runtime edges.
- Pipeline cases live in `gomlc/testdata/pipeline/NNN[_description]/main.gom`. Add or edit the source, then run `just update-golden` to generate IR snapshots and execution output.
- Multi-package cases live in `gomlc/testdata/module/projectNNN[_description]/`. Include a root `goml.toml`, explicit package declarations/imports, public cross-package APIs, and a `package main` entry. Generate their `.out` files with `just update-golden`; these cases do not produce IR snapshots.
- Visibility and package-diagnostic fixtures belong in `gomlc/testdata/module_diagnostics/`.
- Never hand-edit generated golden files, including `.cst`, `.ast`, `.hir`, `.tast`, `.core`, `.mono`, `.lift`, `.anf`, `.go`, and `.out`. Use only `just update-golden` to update snapshots.
- Builtin contract changes invalidate prelude and downstream artifacts through the builtin interface hash. Prelude changes affect its own hash; regenerate affected artifacts and snapshots.
- Avoid test/example traits named `Eq` or `Hash` unless deliberately disambiguated from the builtins.
- Version tests must derive expectations from the version module rather than hard-coding a release number.

## Bootstrap and Language Evolution

- `bootstrap/stage0.env` is the trust root: only published Linux amd64 Release archives with pinned SHA-256 checksums may become stage0. Never use unreleased workflow artifacts.
- Current stage0 must compile the current compiler and driver sources. Every change must preserve this invariant and pass `just ci`.
- Implement new syntax, builtins, standard-library APIs, traits, or type-system capabilities before using them in compiler or driver sources. Tests and fixtures may use them immediately; self-hosted consumers must wait until release and stage0 advancement.
- Introduce standard-library capabilities in two phases: first ship public source, resource packaging, dependency selection, navigation, docs, and external tests while retaining compiler-owned fallbacks; after release and stage0 advancement, migrate consumers and remove fallbacks.
- Incompatible syntax changes need a transition release accepting both forms. Advance stage0, migrate self-hosted sources, then remove the old form in a later release.
- Driver protocols, compiler CLI contracts, and other bootstrap interfaces need at least one release of compatibility overlap so the old stage0 driver can use the new compiler.
- Artifact formats may break compatibility only when each clean bootstrap stage consumes artifacts produced within that stage. Cache and external-dependency errors must remain recoverable.
- Syntax, semantics, builtin, or public standard-library API changes must update the relevant prose, examples, limitations, comparison table, and informal grammar in [docs/goml.md](docs/goml.md) in the same change.

## Commits and Releases

- Prefer concise, imperative Conventional Commits (`feat:`, `fix:`, `refactor:`, `chore:`). PRs should explain the change, link relevant issues, describe behavior changes, and report validation.
- Separate logic and generated snapshots into two commits. When adding tests too, use three commits: logic, test sources, then snapshots generated by `just update-golden`.
- Keep stage0 advancement in its own commit after running `just bootstrap`; migrate self-hosted consumers in later commits.
- During early bootstrap, publish only continuous `0.1.x` patch releases until explicitly lifted. Otherwise use patch for compatible fixes, minor for features and pre-1.0 breaks, and major for post-1.0 breaks. Tags are strict `vX.Y.Z`, advancing one continuous SemVer step.
- Follow [docs/releasing.md](docs/releasing.md): run `just set-version X.Y.Z` and `just ci`, push the release commit, wait for successful main CI on that exact commit, then tag and push it.
- `just set-version` must synchronize `VERSION`, both GoML version modules, and the VS Code package and lockfile versions.
- Main CI owns fixed-point verification and the complete test suite. Release CI requires successful main CI for the exact tagged commit, rebuilds stage2, and runs archive/LSP smoke tests without repeating the full suite.
- After publication, take the archive checksum from `SHA256SUMS`, run `just set-bootstrap-stage0 X.Y.Z <sha256>` and `just bootstrap`, then commit `bootstrap/stage0.env` before relying on newly released capabilities.
