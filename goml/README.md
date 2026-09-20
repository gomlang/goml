# goml project driver

`goml` is the self-hosted project driver. It provides project creation, package discovery, check/build/run/test plans, dependency resolution, registry cache management, incremental artifact fingerprints, native linking, and parallel test execution.

Generated programs and Go FFI target Go 1.26. Install Go 1.26 or newer for native builds, test runners, and Go metadata validation.

From the repository root, build the toolchain and run compiler, driver, and Go metadata tests:

```sh
just all
```

`just test` includes the same tests. `just ci` additionally checks the bootstrap fixed point, scripts, extension, and release packaging.

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

Go FFI is validated by default with `--ffi-check required`. See the [language guide](../docs/goml.md#go-ffi) for the boundary rules, [bind-go](../docs/ffi/bind-go.md) for allowlisted binding generation, and [export-go](../docs/goml.md#exporting-a-go-library) for publishing a generated Go package.

`goml clean` removes the current module's configured build target directory. Use `goml clean --target-dir <path>` to clean another target directory inside the module.

The driver resolves `gomlc` from `--compiler`, `GOMLC`, a sibling binary, `GOML_HOME/bin`, then `PATH`.

Package-management commands are:

```sh
goml update
goml add owner::module
goml add owner::module@1.2.3
goml remove owner::module
```

`update`, `add`, and `remove` accept `--local-registry <path>`. Registry state is stored under `$GOML_HOME/cache/registry`, defaulting to `~/.goml/cache/registry`.

CLI integration tests live in `cmd/goml/*_test.gom`, with shared helpers in `test_support/`; other packages also contain their own unit tests. Their isolated integration workspaces are written below `goml/_artifact/test-work` relative to the repository root.
