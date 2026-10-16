# gomlc

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
