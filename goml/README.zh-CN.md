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

只有编译器标识、参数、输入和记录的输出摘要全部匹配时，才跳过 GoML 编译和链接。原生构建始终调用 Go，由 Go 自身的缓存处理所选 Go 工具链、构建环境和原生依赖。

`check`、`build` 和 `test` 查找所在模块的 `goml.toml` 并操作整个模块。`test` 的可选参数是测试名子串过滤器。`run [TARGET]` 选择可执行包，并把 `--` 后的参数传给程序。`--dry-run` 打印命令计划。`goml fmt` 和 `goml fmt --check` 格式化或验证模块的生产源码和测试源码。

用户命令见[项目工具](../docs/tooling.zh-CN.md)：[模块与工作区配置](../docs/tooling.zh-CN.md#gomltoml)、[包管理](../docs/tooling.zh-CN.md#package-registry)、[示例](../docs/tooling.zh-CN.md#examples)和[测试](../docs/tooling.zh-CN.md#check-and-run-tests)。另见[格式化](../docs/formatting.zh-CN.md)、[API 文档生成](../docs/documentation.zh-CN.md)和 [Go 互操作](../docs/ffi/bind-go.zh-CN.md)。

CLI 集成测试位于 `cmd/goml/*_test.goml`，共享辅助工具位于 `test_support/`；其他包也包含各自的单元测试。隔离的集成工作区写入相对于仓库根目录的 `goml/_artifact/test-work` 下。

驱动器负责项目配置和模块依赖解析。`goml __project-info --source <PATH>` 是供编辑器和编译器查询客户端使用的内部只读协议。其 JSON 响应包含 `protocol: 1`、可为空的 `project` 和绝对 `watched` 路径。项目包含 `root`、`module`、`target`、`development`、`modules` 和可为空的依赖解析 `error`。模块条目包含 `name`、`root`、`dependencies` 和 `dev_dependencies`；只有根模块暴露其开发依赖。模块外的文件返回空项目。无效的根清单使命令失败。依赖失败会保留根元数据，并在描述中报告错误。此命令不查找编译器、不获取依赖，也不创建产物。`--registry-root` 为隔离测试提供显式缓存目录。

包职责按构建流程划分：`project/` 发现和描述项目，`build_model/` 定义命令和计划，`build_plan/` 构造计划，`build/` 负责执行、缓存和调度。`go_export/` 负责 Go 导出路径和输出组装。规划与执行共享命令模型，但互不依赖；项目发现也不导入这两层。相关测试放在各自包中。
