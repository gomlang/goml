<a id="project-tools-and-workflows"></a>

# Project tools and workflows

[English](tooling.md) | [简体中文](tooling.zh-CN.md)

[GoML language guide](goml.md)

Module configuration, dependencies, examples, build output, testing, editor support, and generated-code inspection. Source-level imports and visibility are covered in [modules and packages](language/packages.md).

On this page: [`goml.toml`](#gomltoml), [Package registry](#package-registry), [Examples](#examples), [Native dependencies](#native-dependencies), [Build output layout](#build-output-layout), [Tests](#tests), [LSP and editor](#lsp-and-editor), [Verify generated code](#verify-generated-code).

<a id="gomltoml"></a>

## `goml.toml`

A project module declares the canonical path in the root directory:

```toml
[module]
path = "alice::myapp"

[build]
target-dir = "_artifact"

[dependencies]
"alice::http" = true
```

Registry dependencies declare presence with `true` (an empty inline table is also accepted). Unpinned packages use the cached repository default branch. `goml.lock` can pin exact Git sources; package versions, version constraints, and minimum-version selection are not implemented. Existing strict `X.Y.Z` strings and inline `version` fields remain readable during migration, but do not select or constrain a source. New declarations should use `true` or a local `path`.

`[dev-dependencies]` uses the same unversioned and path forms. These dependencies are available to the invoking module's `*_test.goml` files, `tests/` packages and root `examples/` tree. Production sources cannot import them, even during a test build. Dependencies of another module do not inherit its development dependencies. Test and example commands resolve the union of normal and development dependencies; ordinary check/build/run commands resolve only normal requirements. A development helper may depend back on the current library: that edge uses the current source without introducing a second copy. Production dependency cycles remain errors. A helper that imports the package under test must be used from black-box tests or examples; importing it from that same package's white-box tests still creates an unsupported package cycle.

```toml
[dev-dependencies]
"alice::proptest" = true
"alice::test_support" = { path = "../test_support" }
```

Use `goml add alice::proptest --dev`, `goml add alice::test_support --dev --path ../test_support`, and `goml remove alice::proptest --dev` to edit this section. Production dependencies stay in `[dependencies]`.

Local development dependencies can use an inline table. Paths are relative to the declaring module, and the dependency name must match the target's canonical `[module].path`:

```toml
[dependencies]
"alice::http" = { path = "../http" }
"alice::model" = { path = "../model" }

[replace]
"alice::codec" = { path = "../codec" }
```

`goml add alice::http --path ../http` validates the target and writes a module-relative path. Local modules use their working sources, including transitive local dependencies, without requiring a registry cache. Registry modules cannot declare local path dependencies. Only the invoking module's `[replace]` table applies to its whole dependency graph; dependency-local replacements are ignored. Conflicting local sources, mismatched identities and dependency cycles are errors. Artifacts stay under the consuming project's target directory, and local source edits invalidate cached outputs.

A `goml.work` file groups independent modules:

```toml
[workspace]
members = ["app", "http", "model"]

[replace]
"alice::codec" = "../codec"
```

Members retain their own manifests and file-scoped imports. Membership provides a local source for a declared dependency; it does not make undeclared dependencies visible. Selection prefers workspace replacements, invoking-module replacements, workspace members, explicit path dependencies, then registry sources. Member paths and replacement paths are relative to `goml.work`; duplicate canonical member paths or module identities are rejected. An ancestor workspace applies only to its listed members. `GOML_WORKSPACE` selects a workspace directory explicitly, and `GOML_WORKSPACE=off` disables discovery.

`goml check --workspace`, `build --workspace`, `test --workspace`, `fmt --workspace`, `doc --workspace` and `clean --workspace` run for each member in declaration order, stopping on failure. From a workspace directory without an enclosing module, these commands select all members automatically. `-p alice::app` or `--package alice::app` selects one module; `goml run` requires this selection when multiple members exist. Commands invoked inside a member operate on that member unless `--workspace` is supplied. Compiler queries and LSP dependency navigation resolve the same local sources and replacements.

`[build]` can be omitted; `build.target-dir` defaults to `_artifact` under the module root. The manifest value must be a non-empty relative path and cannot contain a `..` segment. `goml check`, `goml build`, `goml run`, and `goml test` can temporarily override it with `--target-dir <path>`; command-line overrides may be relative or absolute. `goml clean` removes the configured target directory. Its optional `--target-dir <path>` override must stay inside the module.

`goml build` and `goml run` accept `--pgo off|auto|<path>` and `--build-record`. An explicit PGO option writes a versioned `<executable>.build.json` record containing the toolchain, target, build arguments, generated source/profile/artifact fingerprints, and build outcome. Paths resolve relative to the module root; `auto` selects its `default.pgo` or disables PGO if absent. Without the option, existing Go environment behavior is preserved. Native and WASM records use build-time evidence rather than executable metadata inspection. `goml build --explain-performance --format text|json` reports generated size, specialization counts, selected static costs, and GoML inline decisions with enclosing source declarations. Default builds collect neither explanation reports nor build records. See [performance explanations and PGO](performance.md) for exact semantics, machine-readable formats, examples, and limitations.

Each module path segment must be non-empty and may contain ASCII letters, digits, `_`, and `-`. Paths rooted at `builtin` or `prelude` are reserved for the toolchain and cannot be used as a module path or dependency. The `[module]` section currently has no `name`, `kind`, `root`, or similar fields.

<a id="package-registry"></a>

## Package registry

The default central index is [github.com/gomlang/registry](https://github.com/gomlang/registry). Its `index.toml` maps canonical module names to Git repositories in the `gomlang` organization. A module's canonical name comes from its root manifest, so `ecosystem::ansi` resolves to `github.com/gomlang/ansi` without changing its import path.

```sh
goml add ecosystem::ansi
goml add ecosystem::proptest --dev
goml update
goml remove ecosystem::ansi
```

`goml add` writes an unversioned dependency and fetches required sources. `goml update` refreshes the index and the current module's normal and development dependency closure, honoring local paths and workspace replacements, then writes `goml.lock`. Outside a module it updates only the index. Dependencies' own development dependencies are not fetched. Without a lock, builds use cached sources; a missing cache is a recoverable error directing the user to `goml update`. With a lock, ordinary builds may fetch an absent exact revision. Compiler/LSP metadata queries remain read-only and report missing pinned sources without fetching. Updated source contents invalidate generated artifacts.

The default home is `~/.goml`, overridden by `GOML_HOME`. The index checkout lives in `<home>/cache/registry`; Git package checkouts live in `sources/<owner>/<module>/` below that directory. Each checkout follows its default branch. Updates use fast-forward pulls and do not reset local changes. Changing a cached repository's source produces a diagnostic asking the user to remove that cache directory before updating.

`goml lock` creates a version 1 JSON `goml.lock` from the selected normal and root development dependency closure. It retains an existing matching lock; use `goml update` to advance Git revisions. Commit the lock with the project. Each entry records canonical module identity, source kind, source URL or local path, exact Git commit, SHA-256 of its sorted Git file records (path, executable mode and SHA-256 of raw file bytes), and direct dependencies. The root fingerprint covers dependency declarations and effective workspace/replacement sources, excluding build settings and comments. Dependency development requirements are not inherited.

```sh
goml lock
goml build --locked
goml test --locked --offline
goml lock --offline
```

`--locked` requires a matching lock and never updates it. `--offline` prohibits dependency network access; registry Git dependencies require a lock and its exact cached revision, rather than accepting a different mutable checkout. The flags are independent: `--locked` may download a missing pinned commit. `goml lock --offline` can create a lock from already cached sources, but cannot refresh the index. `goml update --offline` is rejected. Manual manifest, workspace selection, or local dependency graph changes require an explicit `goml lock` or `goml update`. `goml add` and `goml remove` refresh an existing lock while retaining existing Git pins.

Pinned Git checkouts coexist under `revisions/<source-sha256>/<commit>/` in the registry cache. Resolution verifies the commit, raw working files against committed Git blobs, executable modes and the SHA-256 file-record checksum, including extra files; corruption never falls back to a mutable checkout. Git submodules, symlinks and checkout filters that change raw file bytes are not supported in locked sources. Workspace members, replacements and path dependencies retain the existing source precedence and remain mutable: their source selection and dependency edges are recorded, but their contents are not claimed to be immutable. Local registry path entries are likewise mutable and record absolute paths. Such locks require the same local layout; they are not portable source archives. Root-relative local dependencies and workspace layouts can move together.

Locks fix dependency sources, not compiler versions or binary ABI. Toolchain and target evidence belongs in build records; identical locks alone do not guarantee identical binaries across toolchains.

A Git index entry is:

```toml
[modules."ecosystem::ansi"]
git = "https://github.com/gomlang/ansi.git"
```

For local registry fixtures, an entry can instead use `path = "owner/module"`, relative to the index root. Exactly one of `git` and `path` is required. Paths must stay inside the registry tree. Every source must contain a root `goml.toml` whose module path matches its entry. There are no version directories or synthetic snapshot versions.

Set `[registry].default` in `<home>/config.toml` to override the index URL; a missing or empty setting uses the central index. `goml update --local-registry <path>` and `goml add ... --local-registry <path>` use a local Git index. `goml remove` edits the manifest without contacting the index. The previous versioned index format is unsupported; use a fresh cache or remove the old registry cache before switching to this index. Dependency declarations and index entries are configuration and do not change the source-language grammar.

<a id="examples"></a>

## Examples

Named examples share the module-root manifest and use both normal and development dependencies. Put each executable in `examples/<name>/main.goml`, declaring `package main;` and `fn main() -> ()`. Names contain ASCII letters, digits, `_`, or `-`. An example can contain helper packages, white-box tests and direct `tests/` packages. An example directory containing its own `goml.toml` is an independent module, not a named example.

```text
library/
├── goml.toml
├── library.goml
├── tests/api_test.goml
└── examples/basic/
    ├── main.goml
    └── tests/example_test.goml
```

For a library with module path `alice::library`, the example package is `alice::library::examples::basic`. Production packages and other modules cannot import example packages. Examples use the existing GoML source grammar; no target declarations or extra manifests are needed.

```sh
goml run --example basic -- argument
goml build --example basic
goml check --examples
goml test --example basic
goml test
```

`check`, `build` and `test` accept `--example <name>` or `--examples`; these options select only the named example or all examples. `run` accepts one `--example` and cannot also take a positional target. Ordinary check/build commands exclude the root `examples/` tree. Ordinary `goml test` runs module tests, builds all examples and runs their tests; it does not execute example `main` functions. `goml fmt` includes examples and their tests. The binary for `basic` is `<target-dir>/bin/examples/basic/basic`. Test processes run from the module root, including example tests.

<a id="native-dependencies"></a>

## Native dependencies

A module containing a Go adapter can declare its Go module and system prerequisites in the same manifest:

```toml
[native]
go-module = "example.com/goml-ecosystem/llvm"
go-version = "v0.0.0"
cgo = "required"
llvm-major = "18"
```

All values are quoted strings. `go-module` must match the adapter's module-root `go.mod`; `go-version` defaults to `v0.0.0` and can specify a Go module version appropriate for a versioned module path. `cgo` accepts `required` or `optional` (the default). `llvm-major` requires a positive decimal major version. Unknown and duplicate native keys are errors.

The consuming project supplies a minimal module-root `go.mod`. For modules with declared native adapters, the driver resolves those adapters from the selected GoML registry sources and generates Go requirements and local replacements under `<target-dir>/native/<hash>/`. It passes the generated module file through `GOFLAGS=-modfile=...` to compilation, tests, and Go export. This managed module-file path leaves the source `go.mod` and `go.sum` files unchanged. The content hash includes adapter locations and module metadata; dry runs do not run prerequisite probes or create these files. An existing `GOFLAGS -modfile` conflicts with managed dependencies and produces a diagnostic.

The driver checks required `CGO_ENABLED=1` and the requested LLVM major using `LLVM_CONFIG`, `llvm-config-N`, `/usr/lib/llvm-N/bin/llvm-config`, then `llvm-config`. These declarations do not install system packages or configure compiler/linker search paths. Consumers still configure `CGO_CFLAGS`, `CGO_LDFLAGS`, and the shared-library loader when needed. Additional Go dependencies are managed by the consuming Go module and the user's Go environment; dependency-local replacements are not promoted into the consumer. Ordinary project builds use normal Go dependency resolution. Go library export and native SIMD builds retain readonly, offline module handling and require dependencies to be prepared beforehand. Registry packages must include their adapter sources and Go module files.

<a id="build-output-layout"></a>

## Build output layout

The default product layout is as follows; the standard package paths in the path are expanded by directory level, for example, `alice::myapp::utils` corresponds to `alice/myapp/utils`:

```text
_artifact/
├── check/
│   ├── pkg/<canonical-package-path>/<last-path-segment>.interface
│   └── deps/<owner>/<module>/pkg/<canonical-package-path>/<last-path-segment>.interface
├── build/
│   ├── pkg/<canonical-package-path>/<last-path-segment>.interface
│   ├── pkg/<canonical-package-path>/<last-path-segment>.core
│   ├── pkg/<entry-package-path>/goml_generated.go
│   └── deps/<owner>/<module>/pkg/<canonical-package-path>/<last-path-segment>.*
├── test/
│   ├── base/
│   │   ├── pkg/<canonical-package-path>/<last-path-segment>.*
│   │   └── deps/<owner>/<module>/pkg/<canonical-package-path>/<last-path-segment>.*
│   ├── internal/
│   │   ├── pkg/<canonical-package-path>/<last-path-segment>.*
│   │   ├── goml_generated.go
│   │   ├── tests.json
│   │   └── runner
│   └── external/
│       ├── pkg/<black-box-test-package-path>/<last-path-segment>.*
│       ├── goml_generated.go
│       ├── tests.json
│       └── runner
└── bin/
    ├── <module-name>
    └── <module-relative-entry-path>/<entry-name>
```

`<last-path-segment>` is the last segment of the canonical package path, not the package name declared in the source code; for example, `alice::app::cmd::server` declared as `package main` still uses `server.core` and `server.interface`.

`check` only requires interface artifacts; `build` and `test` also generate Core. Each entry package declared as `package main` generates a Go link file named `goml_generated.go` in its own artifact directory; ordinary library packages do not generate Go link files. Fixed filenames do not trigger Go's special file rules when a package path ends with `_test`, an operating system name, or an architecture name.

Interface and Core artifacts use the deterministic GAF binary container. GAF
payloads are encoded directly from compiler data without an intermediate JSON
value tree.

`for` loops over builtin `range`, `Vec`, `Slice`, and `MutSlice` values lower directly to indexed loops without allocating iterator closures.

The executable file of the module root entry package is `bin/<module name>`; the nested entry package retains the directory within the module and appends the entry name. For example, the output of `alice::app::cmd::server` is `bin/cmd/server/server`. Production test dependencies are built once under `test/base`; internal and external tests each share one Go entry, manifest, and runner at the corresponding test-kind root. The runner file has the `.exe` suffix on Windows. External dependencies use the same `deps/<owner>/<module>/pkg/...` structure in the root directory of each stage.

The configured production directory will not participate in package discovery and cannot be a target to be inspected, built, or tested. The `.gitignore` generated by `goml new` contains `/_artifact/` by default; after modifying `build.target-dir`, the project's own `.gitignore` should be modified simultaneously.

<a id="tests"></a>

## Tests

<a id="test-functions-and-attributes"></a>

### Test functions and attributes

Top-level functions are marked as tests using `#[test]`:

```goml
package math;

use std::testing;

#[test]
fn addition_works() -> () {
    testing::assert_eq(1 + 1, 2)
}
```

Test functions must meet the following rules:

- It can only be a top-level function, not an impl method, `extern fn`, structure, enumeration or trait;
- Cannot be named `main`;
- Cannot have parameters or type parameters;
- The return type must be `()`;
- Canonical test IDs generated within the same test package must be unique.

Test functions do not require `pub`. `#[test]` does not accept parameters. When you need to skip a test by default, you can use `#[ignore]` without parameters or `#[ignore("reason")]` with a string reason; `#[ignore]` cannot be used alone without `#[test]`:

```goml
#[test]
#[ignore]
fn unfinished_case() -> () {}

#[test]
#[ignore("requires an external service")]
fn integration_case() -> () {}
```

Parameterized tests use one or more `#[test_case(...)]` attributes on a top-level function. Test cases accept string and boolean arguments, which are checked against the function parameters. Each case receives a content-derived stable ID, while list, filter, text/JSON reporting, artifact manifests, and CodeLens continue to use its readable display name. Ordinary zero-argument `#[test]` functions remain unchanged.

```goml
#[test_case("left", true)]
#[test_case("right", false)]
fn is_left(value: string, expected: bool) -> () {
    testing::assert_eq(value == "left", expected)
}
```

See [testing assertions](stdlib.md#testing-assertions) for the assertion and payload-extraction helpers in `std::testing`.

<a id="white-box-and-black-box-tests"></a>

### White-box and black-box tests

The white-box test file and the source code under test are located in the same package directory. The file name must end with `_test.goml` and declare the same package name:

```text
math/
├── math.goml
└── math_test.goml
```

`math_test.goml`:

```goml
package math;

use std::testing;

#[test]
fn private_helper_works() -> () {
    testing::assert_eq(private_helper(), 42)
}
```

The test build will merge the production source code and all `*_test.goml` in the same directory into a single package, so white-box testing can access the private top-level items of the package. These files will not participate in compilation when executing normal `goml check`, `goml build` or `goml run` for production targets.

Black-box tests are located in the `tests/` directory of the package under test. This directory as a whole constitutes a package named `tests`. The package under test should be imported explicitly and only its public API can be accessed; nested test suites cannot be created under `tests/`:

```text
math/
├── math.goml
└── tests/
    ├── api_test.goml
    └── smoke_test.goml
```

`math/tests/api_test.goml`:

```goml
package tests;

use alice::myapp::math;
use std::testing;

#[test]
fn public_add_works() -> () {
    testing::assert_eq(math::add(20, 22), 42)
}
```

If the identity of the package under test is `alice::myapp::math`, the canonical identity of the test package in the above example is `alice::myapp::math::tests`. Ordinary package discovery will exclude the `tests` directory, and production packages cannot import black-box test packages.

In the compiler repository, standard-library behavior tests live in `lib/std/<package>/tests/` and use native GoML assertions or fixed known-answer data. Run `just test-stdlib` from the repository root. For this source module, `goml test` first builds a private compiler world from the sibling `builtin`, `prelude`, and `std` sources under the selected target directory, then compiles the black-box tests against it. This tests the current library source. Standard-library white-box tests are currently unsupported and produce a diagnostic. Installed toolchains omit these test sources and data.

<a id="check-and-run-tests"></a>

### Check and run tests

`goml check`, `goml build` and `goml test` discover the module from the current directory upwards. They process complete modules unless an example selector is supplied, and do not accept package or file targets. `goml check` only checks the production source code by default; using `--tests` will also check all white-box and black-box tests after the production package check is successful:

```sh
goml check
goml check --tests
goml build
goml fmt
goml fmt --check
```

Test source code cannot be used to fix a production package that itself fails inspection.

`goml fmt` formats module sources; `goml fmt --check` reports formatting differences without writing files. See [formatter commands](formatting.md#command-line) for source discovery and options.

Run the test using:

```sh
goml test [FILTER]
```

`FILTER` performs substring matching on the complete test display name, such as `goml test addition`. Common options include:

- `--kind internal|external|all`: run only white box, only black box or all tests, the default is `all`;
- `--list`: List matching tests without running them;
- `--ignored`: only run ignored tests;
- `--include-ignored`: Run normal tests and ignored tests at the same time;
- `--jobs N`: Number of parallel workers, default is `1`;
- `--seed N`: Supply a reproducible positive seed to every test;
- `--timeout 500ms|30s|2m`: timeout for a single test, the default is `30s`;
- `--nocapture`: Let the standard output and standard error of the test directly inherit the terminal;
- `--format text|json`: Select text or line-by-line JSON results; JSON format cannot be used with `--nocapture` at the same time;
- `--target-dir`, `--dry-run` and `--compiler`: have the same meaning as project build commands.

Each test is executed in a separate runner process, and failure to exit and timeout do not affect other tests; `--jobs` controls the number of test processes running at the same time. Building test runners with `goml test` requires an available Go toolchain.

`goml test --no-run` builds all selected test groups without executing tests. `goml test --no-build` executes existing test runners and manifests from the target directory without compiling or linking them. It can be combined with filters, `--list`, and sharding. Prebuilt test artifacts must come from the same source revision and use the execution machine’s OS and architecture; `--no-build` deliberately skips freshness checks, so missing artifacts fail instead of being rebuilt. To transfer artifacts, preserve each test group’s `runner` and `tests.json` paths and any executable-relative runtime libraries. `--no-build` conflicts with `--no-run` and `--dry-run`; `--no-run` conflicts with `--list`.

`GOML_TEST_SHARD=INDEX/COUNT` splits matching tests across independent machines. Both numbers are positive, one-based decimal integers with at most nine digits, and `INDEX` must not exceed `COUNT`. After applying the name and ignored-test filters, the driver sorts tests by display name and assigns successive tests to shards in round-robin order. For example, run `GOML_TEST_SHARD=1/4 goml test --jobs 4` through `GOML_TEST_SHARD=4/4 goml test --jobs 4` on four machines. `--list` uses the same shard selection. An unset or empty value runs the full selection; an invalid value produces a diagnostic. The driver clears this variable in test processes so nested `goml test` commands run their complete selection.

`goml test --seed N` supplies one reproducible positive seed to every test process through `testing::seed()`. Failure summaries print the replay seed, and every JSON result event includes it. Captured stdout and stderr remain isolated per test process and are emitted only for failures unless `--nocapture` is selected.

<a id="lsp-and-editor"></a>

## LSP and editor

Project configuration and module dependency resolution belong to the `goml` driver. The LSP requests a versioned project description from `goml __project-info --source <PATH>` and passes it with open document contents to the compiler query engine. The description supplies canonical module identities, source directories, separate production/development dependency edges, and configuration watch paths. This operation reads local files and cached dependencies without compiling, fetching, or writing build artifacts. Dependency failures remain diagnostics; available root metadata is retained for local package completion.

The project-description client validates protocol version 1 and caches successful descriptions while their watched configuration contents and canonical paths remain unchanged. Ordinary source edits reuse project resolution and still participate in the compiler's source-analysis fingerprints. Workspace creation, dependency manifest changes, replacements, and registry index changes trigger resolution again. Missing dependencies and invalid descriptions are retried after edits rather than cached as successful projects. The compiler retains source import resolution, visibility checks, and type checking.

`GOML_PROJECT_DRIVER` selects an explicit driver. Otherwise clients use a sibling `goml`, `GOML_HOME/bin/goml`, or `PATH`; repository tests may supply `GOML_TEST_GOML`. The driver and language server must support the same project-description protocol. The extension bundles both tools and offers `goml.driverPath` for custom installations. Existing compiler build commands and driver protocol remain unchanged, so pinned stage0 can continue building the tools.

The LSP analyzes production files, internal tests, and black-box tests in their corresponding package contexts, so diagnostics, completion, hover, and go-to-definition follow the same visibility rules as compilation. Test code lenses appear on ordinary tests and parameterized test cases. The VS Code extension saves the file and invokes module-level `goml test` with the selected test display name and kind.

Enum completion uses expected types for contextual constructors and the scrutinee type for `match` patterns. Variant suggestions show payload types and provide tuple or named-field placeholders; new match patterns also provide an arm body placeholder. Existing payloads and arm bodies are preserved. Unguarded unit arms suppress the corresponding variant suggestion. Completion continues to reuse checked types while the document is being edited.

Set `"goml.lsp.enabled": false` in VS Code to use syntax highlighting and bracket matching without starting the language server or requiring a toolchain. The setting takes effect immediately; enabling it starts the server again.

Type inlay hints use short names without package qualifiers and a 25-character type budget, excluding the leading `: `. Nested types and long names are abbreviated with `…`; the hint tooltip retains the complete inferred type and its package paths.

Package analysis caches prepared ASTs per source file and successful checked packages within a query context. Exact source contents and module identity determine AST reuse; package membership and transitive dependency fingerprints determine checked-package reuse. Editing a package rechecks it and its dependents while independent packages remain reusable; a function-body edit in a dependency conservatively invalidates its dependents too. Open document overrides participate in the same fingerprints. External Go types remain opaque during cached and fresh analysis; native validation belongs to the final Go build. Watched source changes invalidate documents that use the affected package while retaining reusable package checks. Unrelated source files, build outputs, and native Go file changes preserve cached document analyses. Changes to tracked configuration files, including module manifests, `goml.work`, and registry indexes, refresh project resolution.

The custom `goml/expandedDerive` request returns the formatted AST after built-in and programmable derives have run for the requested document. The expanded text is formatted on demand and cached within that document analysis. It uses the same package aliases, explicit derive imports, ambiguity checks, dependency interfaces, CTIR verifier, and resource limits as `goml check`. The VS Code command `GoML: Show Expanded Derive` opens that result beside the source file.

See [editor formatting](formatting.md#editor-formatting) for `textDocument/formatting` behavior.

CLI and LSP diagnostics share a model with optional stable codes, notes, help, and suggested edits. The CLI displays codes as `error[typer/try_error_mismatch]` or `warning[typer/unused_result]`; the LSP exposes the same identifier in `Diagnostic.code` and secondary locations in `relatedInformation`. Notes and help appear in the displayed message and separately in diagnostic `data`, alongside fixes and the document version. Fixes carry a title, one or more ranged edits, and `machine-applicable` or `requires-review` applicability. Code actions reject fixes from an older document version. Codes are currently assigned to the unused-Result warning and the three [`?` diagnostics](language/basics.md#try-operator); other diagnostics may omit them.

<a id="verify-generated-code"></a>

## Verify generated code

After running `just make`, verify a standalone source file from the repository root:

```sh
stage2/bin/gomlc run-single path/to/main.goml
```

For a project, run the installed driver from anywhere inside its module:

```sh
goml fmt
goml check
goml build
goml check --tests
goml test
goml run
```

`goml check`, `goml build`, `goml test`, and `goml fmt` operate on complete modules and do not accept package or file targets. In a workspace, use `--workspace` for all members or `-p <module-path>` for one member. `goml run [TARGET]` accepts an optional entry package file or directory when a module has multiple executable packages.

When you need to inspect a compilation phase, add `--dump-ast`, `--dump-expanded-ast`, `--dump-hir`, `--dump-tast`, `--dump-ctir`, `--dump-core`, `--dump-mono`, `--dump-lift`, `--dump-anf`, or `--dump-go` to `gomlc run-single`. `--dump-ast` shows source lowering before derive expansion, while `--dump-expanded-ast` includes every generated implementation.

For repository changes, follow the checks in the [development guide](development.md).

## Architecture dependency rules

`goml architecture` checks direct package dependencies against the module-root `goml.architecture.json`. It uses the parsed graph from `goml metadata`, including alias imports and reexports, and does not compile packages or fetch dependencies. `--rules PATH` selects another rules file, relative to the current directory; `--compiler PROGRAM` selects the compiler used to parse imports.

```sh
goml architecture
goml architecture --rules goml.architecture.json
```

A version 1 rules file declares named prohibitions:

```json
{
  "schema_version": 1,
  "rules": [
    {
      "name": "reusable-no-app",
      "from": [{"package": "demo::device", "subpackages": true}],
      "forbid": [{"package": "demo::web", "subpackages": true}],
      "scope": "production"
    }
  ]
}
```

Selectors use canonical package names. A selector matches only that package unless `subpackages` is `true`, which includes descendants separated by `::`; `demo::web` never matches `demo::website`. Lists match any selector. There are no implicit wildcards. `scope` defaults to `production`; `tests` checks white-box and black-box test imports, while `all` also includes examples. Reexports count as dependencies in the source file's category.

Violations name the rule, the source and target packages, and the import's source location. Results are sorted independently of discovery order. An incomplete graph, unresolved dependency, unsupported metadata, or invalid rules file fails the check rather than reporting success. Unknown fields and duplicate rule names are errors. Version 1 checks direct edges only; it does not enforce transitive prohibitions, allowances, or exceptions. This command checks project architecture, not symbol visibility or source API compatibility.

## Package metadata

`goml metadata --format json` emits a versioned, deterministic package graph for external tools. `--mode files` only discovers files and needs no compiler; the default `--mode graph` calls the matching compiler's package-info parser to resolve package dependencies. `--compiler PATH` selects that compiler. Both modes are read-only: they do not fetch dependencies, compile packages, or write build artifacts. Existing locks are respected using locally available sources.

Schema version 1 contains `schema_version`, `toolchain_version`, `root_module`, `mode`, `complete`, `modules`, `packages`, `edges`, and `diagnostics`. Module and package IDs are canonical. Package directories and source files are module-relative; files distinguish production, internal tests, external tests, and examples. Module source kinds distinguish the root, paths, workspace members, replacements, and registry dependencies. A `locked` object records the selected locked source, revision, checksum, and dependency names when available; otherwise it is null. Edges identify the source and target packages, import spelling, file category, reexport status, and source location. Edge source paths are absolute; lines and UTF-8 byte columns start at one.

Missing dependencies, unresolved package imports, or unsupported compiler metadata produce `complete: false`, diagnostic messages, and a nonzero exit status. JSON remains on stdout, while diagnostic text goes to stderr. Graph mode resolves owning packages and import aliases; it does not type-check imported members or validate access permissions. Standard-library endpoints resolve through executable-relative library sources; their full transitive graph is not expanded.

File metadata can replace manually maintained package lists in test snapshots and source archives. Publishers must still choose additional resources and license files explicitly.
