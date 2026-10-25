# gomlc

[English](README.md) | [简体中文](README.zh-CN.md)

`gomlc` is the complete self-hosted GoML compiler and language server. It implements:

```text
lexer → parser → CST → AST → HIR → TAST → Core → Mono → Lift → ANF → Go
```

`hir/` lowers AST and resolves names. `checker/` checks HIR and owns inference,
constraint solving, method selection, diagnostics, and query facts. Its
`CheckResult` contains the optional TAST file, compile-time output, diagnostics,
and `TypeckResults` used by `query/`.

`tast/` defines the typed representation and shared signature metadata. It owns
type inspection and substitution, traversal, printing, verification, and builtin
type queries used by later stages. The checker depends on TAST; TAST and
production backend code do not import the checker. Core lowering accepts an
optional `tast::File`, while callers retain the checker diagnostics and query
facts. Builtin and prelude source checking lives in `checker/toolchain.goml`.

Inference variables and deferred callable tags currently remain in the shared
representation. The checker resolves deferred calls before a successful result;
separating inference types from finalized types is a separate concern from the
package boundary.

See [repository development](../docs/development.md) for requirements,
toolchain layout, builds and bootstrap verification.

Run a single source or inspect an IR stage:

```sh
stage2/bin/gomlc run-single file.goml
stage2/bin/gomlc anf file.goml
stage2/bin/gomlc run-single --dump-go file.goml
```

The regression corpus and every generated golden file live in `gomlc/testdata`. Verify or update them with:

```sh
just verify-golden
just update-golden
```

See the [language guide](../docs/goml.md), [formatter rules](../docs/formatting.md), [API documentation](../docs/documentation.md), and [compile-time evaluation architecture](../docs/comptime.md) for the corresponding compiler contracts.

Project-aware queries obtain module descriptions from the matching `goml` driver through `project_info/`; `query::analyze_with_project` also accepts an explicit description from its caller. Manifest, workspace, and registry interpretation remain in the driver. Source import resolution, package analysis, and navigation remain in `query/`. Standalone compiler commands keep their existing inputs and do not need project dependency resolution. Set `GOML_PROJECT_DRIVER` when the driver is not installed beside the compiler or language server and cannot be discovered through `GOML_HOME/bin` or `PATH`.

`compiler/` owns source compilation, package worlds, artifact loading and linking;
`cmd/gomlc/` owns command parsing and command execution. Compiler service tests
live alongside the services, while CLI tests remain with the command.

`diagnostic/` owns shared diagnostic values and formatting. `builtin_source/`
owns builtin source loading, parsing and cached HIR metadata; `tast/` retains
typed builtin conversion and contract validation. Installed worlds inject builtin
metadata through the same loader, and the lazy source fallback remains available.
`names/` supplies shared name matching without depending on typed representations.

`go_ir/` owns the Go representation and printer, `go_types/` owns representation
type lowering and mangling, and `go_runtime/` emits runtime support. `go_backend/`
assembles these pieces, lowers control flow and performs Go dead-code elimination.
Standard-package selection uses the resource catalog and its dependency closure,
rather than a separate boolean field for every package.

`just check-package-boundaries` checks production imports, dependency layers and
cycles for both compiler and driver packages. It also rejects imports of command
packages from reusable packages.
