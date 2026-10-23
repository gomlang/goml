<a id="goml-project-driver"></a>

# goml 项目驱动器

[English](README.md) | [简体中文](README.zh-CN.md)

`goml` 是自托管项目驱动器，提供项目创建、包发现、检查/构建/运行/测试计划、命名示例、依赖解析、注册表缓存管理、增量产物指纹、原生链接和并行测试执行。

环境要求、工具链构建和仓库检查见[仓库开发指南](../docs/development.zh-CN.md)。

从仓库根目录进入驱动器模块，直接使用 stage2：

```sh
cd goml
../stage2/bin/goml check \
  --compiler ../stage2/bin/gomlc
../stage2/bin/goml test \
  --compiler ../stage2/bin/gomlc \
  --jobs 4 \
  --timeout 10m
```

`goml test --nocapture` 继承测试输出，`--timeout` 接受以 `ms`、`s` 或 `m` 为单位的正时长。只有编译器标识、参数、输入和记录的输出摘要全部匹配时，才跳过 GoML 编译和链接。原生构建始终调用 Go，由 Go 自身的缓存处理所选 Go 工具链、构建环境和原生依赖。

`check`、`build` 和 `test` 查找所在模块的 `goml.toml` 并操作整个模块。`test` 的可选参数是测试名子串过滤器。`run [TARGET]` 选择可执行包，并把 `--` 后的参数传给程序。`--dry-run` 打印命令计划。`goml fmt` 和 `goml fmt --check` 格式化或验证模块的生产源码和测试源码。

`goml doc` 在 `<target-dir>/doc/` 为模块生成离线 HTML API 文档。使用 `--format json` 获取带版本的 API 模型，使用 `--document-private-items` 生成内部文档。注释语法、链接和独立的 `gomldoc` 命令见[文档说明](../docs/documentation.zh-CN.md)。

Go FFI 信任显式绑定签名；最终 Go 构建检查原生符号和签名。外部 Go 类型是不透明的。边界规则见[语言指南](../docs/ffi/bind-go.zh-CN.md#go-ffi)，显式绑定见 [Go 互操作](../docs/ffi/bind-go.zh-CN.md)，发布生成的 Go 包见 [export-go](../docs/ffi/bind-go.zh-CN.md#exporting-a-go-library)。

`goml clean` 删除当前模块配置的构建目标目录。使用 `goml clean --target-dir <path>` 清理模块内的另一个目标目录。

驱动器依次通过 `--compiler`、`GOMLC`、同目录二进制文件、`GOML_HOME/bin` 和 `PATH` 解析 `gomlc`。

包管理命令如下：

```sh
goml update
goml add owner::module
goml add owner::module --path ../module
goml add owner::test_support --dev
goml remove owner::module
```

默认索引为 [gomlang/registry](https://github.com/gomlang/registry)。依赖使用 `"owner::module" = true`；Git 源码跟随默认分支，不进行版本选择。`add` 获取所需源码，`update` 刷新当前模块的普通依赖和开发依赖。构建和 LSP 查询读取本地缓存。注册表状态存储在 `$GOML_HOME/cache/registry`，默认是 `~/.goml/cache/registry`，源码检出目录为 `sources/<owner>/<module>/`。可通过主目录配置中的 `[registry].default`，或 `update`/`add` 的 `--local-registry <path>` 覆盖索引。迁移和本地测试用例见[注册表文档](../docs/tooling.zh-CN.md#package-registry)。

本地路径依赖和根清单的 `[replace]` 条目支持同时开发库和应用。包含 `[workspace].members` 的 `goml.work` 文件让声明的成员依赖从各自的工作树解析。使用 `--workspace` 对所有成员执行 check/build/test/fmt/doc/clean，或用 `-p owner::module` 选择单个成员。优先级、源码验证和 `GOML_WORKSPACE` 见语言指南的模块章节。

CLI 集成测试位于 `cmd/goml/*_test.goml`，共享辅助工具位于 `test_support/`；其他包也包含各自的单元测试。隔离的集成工作区写入相对于仓库根目录的 `goml/_artifact/test-work` 下。

`[dev-dependencies]` 包含测试和示例依赖，不向生产代码或下游依赖者暴露。可执行示例放在 `examples/<name>/main.goml`，共用根清单。使用 `goml run --example <name>`、`goml build --examples` 或 `goml test --example <name>`；普通的 `goml test` 也会构建并测试全部示例。

驱动器负责项目配置和模块依赖解析。`goml __project-info --source <PATH>` 是供编辑器和编译器查询客户端使用的内部只读协议。其 JSON 响应包含 `protocol: 1`、可为空的 `project` 和绝对 `watched` 路径。项目包含 `root`、`module`、`target`、`development`、`modules` 和可为空的依赖解析 `error`。模块条目包含 `name`、`root`、`dependencies` 和 `dev_dependencies`；只有根模块暴露其开发依赖。模块外的文件返回空项目。无效的根清单使命令失败。依赖失败会保留根元数据，并在描述中报告错误。此命令不查找编译器、不获取依赖，也不创建产物。`--registry-root` 为隔离测试提供显式缓存目录。
