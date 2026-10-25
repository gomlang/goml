<a id="developing-goml"></a>

# 开发 GoML

[English](development.md) | [简体中文](development.zh-CN.md)

请在仓库根目录运行配方。构建工具链需要 Linux amd64 或 macOS arm64、Go 1.26+、`just`、Bash、curl、tar，以及 sha256sum 或 shasum。测试需要 C 编译器以运行 Go 的竞态检测器；完整 CI 还使用 Node 20+、npm、jq、Python 3、用于校验和可移植性测试的 Perl `shasum`，以及用于检查生成的系统调用 ABI 和记录布局的 Linux 开发头文件。所有命令见 [.justfile](../.justfile)。独立的 [gomlgo 前端和解释器](https://github.com/gomlang/gomlgo) 面向 Go 1.26，其执行和差分测试需要 Go 1.26.x。

```sh
just make
just test
just ci
just clean
```

`just make` 直接使用固定版本的 stage0 增量构建 stage2。`just test` 构建工具并运行编译器、项目驱动器和标准库测试；`just all` 是它的别名。`just ci` 执行干净的 stage2 构建、不动点验证、测试、扩展编译和发布归档冒烟检查。标准库行为测试位于实现旁的 `lib/std/<package>/tests`，使用 `just test-stdlib` 运行。已知答案的来源、数据格式和更新规则见[标准库已知答案](../lib/std/KNOWN_ANSWERS.zh-CN.md)。

本地 CI 并发运行各检查组。设置 `GOML_CI_SEQUENTIAL=1` 可改为顺序运行，`GOML_BUILD_JOBS` 限制自举包构建的工作线程数，`GOML_TEST_JOBS` 覆盖编译器、驱动器和标准库测试的并发度。编译器测试用例可以使用 `PATH` 中的 Yaegi；若不可用或无法运行某个用例，则回退到原生 Go 编译。

GitHub CI 只构建一次测试产物，然后使用 `goml test --no-build` 分片运行。任务依赖见[工作流](../.github/workflows/ci.yml)，分片和产物复用见[测试命令](tooling.zh-CN.md#check-and-run-tests)。

自举从 [bootstrap/stage0.env](../bootstrap/stage0.env) 选择宿主对应的已发布 `linux-amd64` 或 `darwin-arm64` stage0 归档，并验证固定的 SHA-256 校验和。两个平台使用相同的发布版本。将 `GOML_STAGE0_ARCHIVE` 设为宿主平台固定归档的本地副本，可避免下载。不支持的宿主会在下载或清理自举输出前被拒绝。

`just bootstrap` 用 stage0 重建 stage2，用 stage2 构建 stage3，再用 stage3 重建编译器和驱动器产物，并与第一次 stage3 构建比较。

在 macOS arm64 上，执行 `just make` 或 `just bootstrap` 后运行 `just _ci-toolchain-scripts` 和 `just _ci-release-smoke`，检查安装、缓存、打包和已发布工具链工作流。完整的 `just ci`、编译器和标准库测试套件需要 Linux amd64，因为它们会使用 Linux 系统调用和 ABI 头文件。GitHub CI 除 macOS 发布归档冒烟测试外，还运行原生 macOS 自举和工具链脚本检查。

本地开发使用 `stage2/bin`。`stage0`、`stage2` 和 `stage3` 下的工具链前缀包含 `bin`、`lib` 下的库项目，以及 `lib/compiler` 下已完成初始化的编译器 world。下载归档缓存在 `_bootstrap/cache`；编译器和驱动器的自举产物位于各模块自己的 `_bootstrap` 目录。Git 会忽略生成的输出。

根目录的 `goml.work` 将 `gomlc` 和 `goml` 组织在一起。驱动器通过 `goml/manifest` 和 `goml/registry` 负责清单解析、工作区选择、注册表索引和依赖解析。编译器查询引擎使用 `goml` 提供的项目描述，不自行解析清单或注册表依赖。运行 `stage2/bin/goml fmt --check --workspace` 检查工作区格式。`lib/` 下的库模块单独格式化和测试。

`just install` 将工具安装到 `${GOML_HOME:-$HOME/.goml}` 并完成初始化；请把其 `bin` 目录加入 `PATH`。`just clean` 删除根目录和编译器/驱动器构建缓存，以及生成的开发阶段工具链，但保留已下载的 `stage0` 工具链。

[gomlc/testdata/pipeline](../gomlc/testdata/pipeline) 包含源程序和每个编译阶段的 golden 文件。使用 `just verify-golden` 检查整个用例集，或使用 `just update-golden` 通过自托管编译器重新生成。

直接 Go FFI 集成用例位于 [goml/testdata/ffi](../goml/testdata/ffi)，由驱动器回归测试和发布冒烟检查共用。完整的 SQLite 和 LLVM 集成分别维护在 [gomlang/sqlite](https://github.com/gomlang/sqlite) 和 [gomlang/llvm](https://github.com/gomlang/llvm)。

组件专属命令和架构见[编译器](../gomlc/README.zh-CN.md)和[项目驱动器](../goml/README.zh-CN.md)文档。

<a id="bilingual-documentation"></a>

## 双语文档

英文文档保留现有路径，作为规范主版本。仓库中每篇 Markdown 文档都在同目录提供完整的简体中文译本，文件名使用 `.zh-CN.md` 后缀。新增或修改文档、语言语义、公共 API 或开发流程时，在同一次变更中同步更新两个版本。`AGENTS.md` 仍是权威指令文件；中文版供阅读参考。

两种语言均保留代码示例、命令、API 标识符和诊断/输出文本。术语保持一致：package（包）、module（模块）、monomorphization（单态化）、closure lifting（闭包提升）和 compile-time evaluation（编译期求值）；`trait` 和 IR 名称在标识语言或编译器概念时保留原文。完整翻译解释性文字，不用摘要替代原章节。

[语言指南](goml.zh-CN.md)作为概览和阅读导航。语言规则维护在其链接的 `language/` 专题中，项目工作流维护在[项目工具](tooling.zh-CN.md)，Go 边界规则维护在 [Go 互操作](ffi/bind-go.zh-CN.md)。重组内容时，保留概览中的原章节锚点并链接到规范专题位置；其他文档直接链接到这些位置。

每篇文档包含 `English | 简体中文` 切换链接。存在中文对应文档时，中文链接指向中文版本。在中文文档中使用显式锚点保留英文标题片段，使已有交叉引用继续有效。修改标题时同步更新链接和锚点。

在仓库根目录运行 `just check-docs`（或 `python3 tools/check-docs.py`），检查双语覆盖、语言切换、标题结构、代码示例，以及本地链接和锚点。这些检查发现结构偏差；审阅者仍需检查翻译准确性和完整性。纯文档改动只需运行适用的文档检查，不必运行完整编译器 CI 套件。

运行 `just check-package-boundaries` 验证编译器和驱动器生产代码的导入、依赖层次和循环。CI 也会执行此检查。
