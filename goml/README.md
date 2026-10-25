# goml project driver

[English](README.md) | [简体中文](README.zh-CN.md)

`goml` is the self-hosted project driver. It provides project creation, package discovery, check/build/run/test plans, named examples, dependency resolution, registry cache management, incremental artifact fingerprints, native linking, and parallel test execution.

See [repository development](../docs/development.md) for requirements,
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

GoML compilation and linking are skipped only when their compiler identity, arguments, inputs, and recorded output digests all match. Native builds always invoke Go, whose own cache accounts for the selected Go toolchain, build environment, and native dependencies.

`check`, `build`, and `test` discover the enclosing `goml.toml` and operate on the complete module. The optional argument to `test` is a test-name substring filter. `run [TARGET]` selects an executable package and passes arguments after `--` to the program. `--dry-run` prints the command plan. `goml fmt` and `goml fmt --check` format or verify the module's production and test sources.

See [project tools](../docs/tooling.md) for user commands: [module and workspace configuration](../docs/tooling.md#gomltoml), [package management](../docs/tooling.md#package-registry), [examples](../docs/tooling.md#examples), and [testing](../docs/tooling.md#check-and-run-tests). See also [formatting](../docs/formatting.md), [API documentation generation](../docs/documentation.md), and [Go interoperability](../docs/ffi/bind-go.md).

CLI integration tests live in `cmd/goml/*_test.goml`, with shared helpers in `test_support/`; other packages also contain their own unit tests. Their isolated integration workspaces are written below `goml/_artifact/test-work` relative to the repository root.

The driver owns project configuration and module dependency resolution. `goml __project-info --source <PATH>` is an internal, read-only protocol for editor and compiler query clients. Its JSON response has `protocol: 1`, a nullable `project`, and absolute `watched` paths. A project contains `root`, `module`, `target`, `development`, `modules`, and a nullable dependency-resolution `error`. Module entries contain `name`, `root`, `dependencies`, and `dev_dependencies`; only the root module exposes its development dependencies. Files outside modules return a null project. Invalid root manifests fail the command. Dependency failures preserve root metadata and report the error in the description. The command does not locate a compiler, fetch dependencies, or create artifacts. `--registry-root` provides an explicit cache directory for isolated tests.

Package responsibilities follow the build flow: `project/` discovers and describes
projects, `build_model/` defines commands and plans, `build_plan/` constructs those
plans, and `build/` executes them with caching and scheduling. `go_export/` owns
Go export paths and output assembly. Planning and execution share the command
model without depending on each other; discovery does not import either layer.
Their tests live in the corresponding packages.
