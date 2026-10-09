# goml project driver

`goml` is the self-hosted project driver. It provides project creation, package discovery, check/build/run/test plans, named examples, dependency resolution, registry cache management, incremental artifact fingerprints, native linking, and parallel test execution.

See [repository development](../README.md#development) for requirements,
toolchain builds and repository checks.

Starting at the repository root, enter the driver module to use stage2 directly:

```sh
cd goml
../stage2/bin/goml check \
  --compiler ../stage2/bin/gomlc
../stage2/bin/goml test \
  --compiler ../stage2/bin/gomlc \
  --jobs 4 \
  --timeout 10m
```

`goml test --nocapture` inherits test output, while `--timeout` accepts positive `ms`, `s`, or `m` durations. GoML compilation and linking are skipped only when their compiler identity, arguments, inputs, and recorded output digests all match. Native builds always invoke Go, whose own cache accounts for the selected Go toolchain, build environment, and native dependencies.

`check`, `build`, and `test` discover the enclosing `goml.toml` and operate on the complete module. The optional argument to `test` is a test-name substring filter. `run [TARGET]` selects an executable package and passes arguments after `--` to the program. `--dry-run` prints the command plan. `goml fmt` and `goml fmt --check` format or verify the module's production and test sources.

`goml doc` generates offline HTML API documentation for the module in `<target-dir>/doc/`. Use `--format json` for a versioned API model or `--document-private-items` for internal documentation. See [documentation](../docs/documentation.md) for comment syntax, links, and the standalone `gomldoc` command.

Go FFI is validated by default with `--ffi-check required`. See the [language guide](../docs/goml.md#go-ffi) for the boundary rules, [bind-go](../docs/ffi/bind-go.md) for allowlisted binding generation, and [export-go](../docs/goml.md#exporting-a-go-library) for publishing a generated Go package.

`goml clean` removes the current module's configured build target directory. Use `goml clean --target-dir <path>` to clean another target directory inside the module.

The driver resolves `gomlc` from `--compiler`, `GOMLC`, a sibling binary, `GOML_HOME/bin`, then `PATH`.

Package-management commands are:

```sh
goml update
goml add owner::module
goml add owner::module --path ../module
goml add owner::test_support --dev
goml remove owner::module
```

The default index is [gomlang/registry](https://github.com/gomlang/registry). Dependencies use `"owner::module" = true`; Git sources follow their default branch without version selection. `add` fetches required sources, and `update` refreshes the current module's normal and development dependencies. Builds and LSP queries read the local cache. Registry state is stored under `$GOML_HOME/cache/registry`, defaulting to `~/.goml/cache/registry`, with source checkouts under `sources/<owner>/<module>/`. Override the index through `[registry].default` in the home configuration or `--local-registry <path>` on `update`/`add`. See the [registry documentation](../docs/goml.md#package-registry) for migration and local fixtures.

Local path dependencies and root `[replace]` entries support simultaneous library and application development. A `goml.work` file with `[workspace].members` makes declared member dependencies resolve from their working trees. Use `--workspace` for all-member check/build/test/fmt/doc/clean, or `-p owner::module` to select one member. See the language guide's module section for precedence, source validation and `GOML_WORKSPACE`.

CLI integration tests live in `cmd/goml/*_test.goml`, with shared helpers in `test_support/`; other packages also contain their own unit tests. Their isolated integration workspaces are written below `goml/_artifact/test-work` relative to the repository root.

`[dev-dependencies]` contains test and example dependencies without exposing them to production code or downstream dependents. Put executable examples in `examples/<name>/main.goml`, sharing the root manifest. Use `goml run --example <name>`, `goml build --examples`, or `goml test --example <name>`; ordinary `goml test` also builds and tests all examples.

The driver owns project configuration and module dependency resolution. `goml __project-info --source <PATH>` is an internal, read-only protocol for editor and compiler query clients. Its JSON response has `protocol: 1`, a nullable `project`, and absolute `watched` paths. A project contains `root`, `module`, `target`, `development`, `modules`, and a nullable dependency-resolution `error`. Module entries contain `name`, `root`, `dependencies`, and `dev_dependencies`; only the root module exposes its development dependencies. Files outside modules return a null project. Invalid root manifests fail the command. Dependency failures preserve root metadata and report the error in the description. The command does not locate a compiler, fetch dependencies, or create artifacts. `--registry-root` provides an explicit cache directory for isolated tests.
