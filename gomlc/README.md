# gomlc

`gomlc` is the complete self-hosted GoML compiler and language server. It implements:

```text
lexer → parser → CST → AST → HIR → TAST → Core → Mono → Lift → ANF → Go
```

See [repository development](../README.md#development) for requirements,
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
