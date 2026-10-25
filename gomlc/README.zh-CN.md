<a id="gomlc"></a>

# gomlc

[English](README.md) | [简体中文](README.zh-CN.md)

`gomlc` 是完整的自托管 GoML 编译器和语言服务器，实现以下流水线：

```text
lexer → parser → CST → AST → HIR → TAST → Core → Mono → Lift → ANF → Go
```

`hir/` 负责 AST 降级和名称解析。`checker/` 检查 HIR，负责类型推断、约束求解、方法选择、诊断和查询事实。其 `CheckResult` 包含可选的 TAST 文件、编译期输出、诊断，以及供 `query/` 使用的 `TypeckResults`。

`tast/` 定义带类型的表示和共享签名元数据，负责类型检查与替换、遍历、打印、验证，以及后续阶段使用的内建类型查询。检查器依赖 TAST；TAST 和生产后端代码不导入检查器。Core 降级接受可选的 `tast::File`，调用方保留检查器诊断和查询事实。内建库和 prelude 的源码检查位于 `checker/toolchain.goml`。

推断变量和延迟可调用标记目前仍保留在共享表示中。检查器在成功返回结果前解析延迟调用；将推断类型与最终类型分离，是独立于包边界的另一个问题。

环境要求、工具链布局、构建和自举验证见[仓库开发指南](../docs/development.zh-CN.md)。

运行单个源文件或检查某个 IR 阶段：

```sh
stage2/bin/gomlc run-single file.goml
stage2/bin/gomlc anf file.goml
stage2/bin/gomlc run-single --dump-go file.goml
```

回归用例集和所有生成的 golden 文件位于 `gomlc/testdata`。使用以下命令验证或更新：

```sh
just verify-golden
just update-golden
```

相应的编译器契约见[语言指南](../docs/goml.zh-CN.md)、[格式化规则](../docs/formatting.zh-CN.md)、[API 文档](../docs/documentation.zh-CN.md)和[编译期求值架构](../docs/comptime.zh-CN.md)。

感知项目的查询通过 `project_info/` 从匹配的 `goml` 驱动器获取模块描述；`query::analyze_with_project` 也接受调用方提供的显式描述。清单、工作区和注册表解释仍由驱动器负责。源文件导入解析、包分析和导航仍由 `query/` 负责。独立编译器命令保留现有输入，不需要项目依赖解析。如果驱动器未安装在编译器或语言服务器旁，且无法通过 `GOML_HOME/bin` 或 `PATH` 发现，请设置 `GOML_PROJECT_DRIVER`。

`compiler/` 负责源码编译、包世界、产物加载和链接；`cmd/gomlc/` 负责命令解析和执行。编译服务测试与服务放在一起，CLI 测试保留在命令包中。

`diagnostic/` 负责共享诊断数据和格式化。`builtin_source/` 负责内建源码加载、解析和 HIR 元数据缓存；`tast/` 保留带类型的内建转换和契约验证。已安装的世界通过同一加载器注入内建元数据，并保留按需加载源码的回退路径。`names/` 提供不依赖带类型表示的共享名称匹配。

`go_ir/` 负责 Go 表示和打印，`go_types/` 负责表示类型的降级和名称编码，`go_runtime/` 生成运行时支持。`go_backend/` 组装这些部分，降级控制流并执行 Go 死代码消除。标准包选择使用资源目录及其依赖闭包，不再为每个包维护单独的布尔字段。

`just check-package-boundaries` 检查编译器和驱动器生产代码的导入、依赖层次和循环，也禁止可复用包导入命令包。
