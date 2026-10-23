<a id="repository-guidelines"></a>

# 仓库规范

[English](AGENTS.md) | [简体中文](AGENTS.zh-CN.md)

goml 是一门采用静态类型和垃圾回收、具有类似 Rust 语法并编译为 Go 的语言。源文件使用 `.goml` 扩展名；没有所有权系统或生命周期语法。编译器会对泛型进行单态化，并对 GoML 闭包进行闭包提升。

<a id="documentation-and-source-map"></a>

## 文档与源码导航

[AGENTS.md](AGENTS.md) 是权威的代理执行指令入口。本文件为其中文参考译文。两个版本应保持同步；译文不定义额外的指令作用域。

请阅读与改动相关的文档：

- [语言指南](docs/goml.zh-CN.md)：规范语法、语义、包、内建功能、示例和文法。
- [标准库](docs/stdlib.zh-CN.md)：公开标准库 API、示例和限制。
- [格式化](docs/formatting.zh-CN.md)：格式化工具行为和 CLI。
- [发布](docs/releasing.zh-CN.md)：发布、工具链安装和 stage0 升级。
- [编译期求值](docs/comptime.zh-CN.md)与 [Go FFI](docs/ffi/bind-go.zh-CN.md)：专门工作流和显式外部声明。
- [gomlgo](https://github.com/gomlang/gomlgo)：独立的 Go 前端、解释器、兼容性目标和差分测试。
- [VS Code 扩展](editors/vscode/README.zh-CN.md)：编辑器设置、行为和配置。

| 路径 | 职责 |
| --- | --- |
| `gomlc/` | 自托管编译器、查询引擎、LSP、资源加载器及编译器测试 |
| `goml/` | 自托管项目驱动程序、注册表客户端、依赖解析器及 CLI 测试 |
| `lib/builtin/` | 隐藏的编译器/运行时契约、运行时钩子和语言项 |
| `lib/prelude/` | 定义自动进入作用域的公开 API 的独立包 |
| `lib/std/` | 通过显式导入加载的标准库包 |
| `bootstrap/` | 自举脚本以及通过校验和固定的已发布 stage0 元数据 |
| `tools/` | 构建、库打包、发布和源码数据生成 |
| `editors/vscode/` | VS Code 扩展 |
| `gomlc/testdata/` | 编译器回归测试用例和生成的黄金文件 |

当前自举使用 `stage0`、`stage2` 和 `stage3` 下的工具链前缀；本地开发使用 `stage2/bin`。各工具链从相对于可执行文件的 `lib/` 加载资源，包括 `lib/compiler/` 下已最终化的编译器 world。构建输出应放在 `_bootstrap/`、`_artifact/` 或模块配置的目标目录中。

<a id="development-workflow"></a>

## 开发工作流

环境要求：Linux amd64 或 macOS arm64、Go 1.26+、用于竞态检测测试的 C 编译器、Node 20+、npm、`just`、Bash、curl、tar、sha256sum 或 shasum，以及 jq。完整 `just ci` 和 Linux 系统调用测试要求 Linux amd64；macOS 支持原生自举、不动点验证、工具链脚本测试和发布冒烟测试。生成的 Go 代码和 Go FFI 面向 Go 1.26。gomlgo 解释器和差分测试要求 Go 1.26.x；详情见其 README。

从仓库根目录运行任务；[.justfile](.justfile) 是命令参考。

| 命令 | 用途 |
| --- | --- |
| `just make` | 从固定的 stage0 增量构建 stage2 |
| `just test` / `just all` | 构建并运行编译器、驱动程序和标准库测试 |
| `just test-stdlib` | 构建 stage2 并运行原生标准库包测试 |
| `just ci` | 完整 CI，包括自举不动点和打包检查 |
| `just bootstrap` | 全新自举和不动点验证 |
| `just verify-golden` / `just update-golden` | 通过自托管测试验证或重新生成快照 |
| `just vscode-ext` | 构建 LSP 并编译扩展 |
| `just clean` | 删除根目录和编译器/驱动程序的构建缓存及生成的开发阶段；保留 stage0 |

- 编辑 `.goml` 文件后，在测试或提交之前，从每个受影响的模块运行 `goml fmt`。模块包括 `gomlc/`、`goml/` 以及 `lib/` 下的独立库项目。
- 使用仓库的格式化工具，例如 `cd gomlc && ../stage2/bin/goml fmt`；`fmt --check` 用于验证格式。
- 使用 `stage2/bin/gomlc run-single <file.goml>` 运行单个测试用例。可添加 `--dump-ast`、`--dump-expanded-ast`、`--dump-hir`、`--dump-tast`、`--dump-ctir`、`--dump-core`、`--dump-mono`、`--dump-lift`、`--dump-anf` 或 `--dump-go` 检查降级结果。
- `goml check`、`goml build` 和 `goml test` 查找所属的 `goml.toml`，并操作整个模块，不支持以包为目标。`--dry-run` 输出计划执行的命令。
- 驱动程序依次通过 `--compiler`、`GOMLC`、同目录二进制文件、`GOML_HOME/bin` 和 `PATH` 查找 `gomlc`，并验证驱动协议。
- 运行与改动相关的检查。涉及自举兼容性、工具链构建或打包的改动、发布准备，以及明确要求时，应在本地运行 `just ci`。只读审阅和纯文档改动只需运行适用的检查。
- 英文和简体中文文档应一同维护：保留现有英文路径，每篇 Markdown 文档配对一个 `.zh-CN.md` 译文。在同一改动中更新两种语言，保持代码示例和 API 名称不变，并使用 `python3 tools/check-docs.py` 验证覆盖率和本地链接。详情见[文档维护](docs/development.zh-CN.md)。
- gomlgo 在独立仓库中维护，使用自身的测试工作流。检查通过后，仅在新改动、失败或尚未解决的问题需要时才重跑或扩大检查范围。

<a id="coding-and-architecture-rules"></a>

## 编码与架构规则

- 不添加代码注释；编写清晰且能自我说明的代码。
- GoML 使用四空格缩进、snake_case 函数名和包名、CamelCase 类型名，以及显式的顶层函数签名。泛型使用方括号；局部闭包只有一个具体类型，不能声明泛型，也不能进行 let 泛化。
- TypeScript 使用两空格缩进和 PascalCase 组件名；优先使用具名导出。
- 保持包小而职责集中。存在方法语法时优先使用，例如 `s.len()` 和 `x.to_string()`。
- 当预期类型能够确定枚举时，优先使用 `Ok(value)`、`Some(value)` 和 `None` 等上下文枚举变体。推断或消歧需要时，保留限定变体。
- 访问 `Vec` 元素时优先使用 `values[index]`。对于包括 `FrozenVec` 在内的必须使用 `get` 的类型，以及语义不同的查找 API，保留 `get`。
- 接收者明确时，将具体对象上的操作定义为方法，并通过接收者调用。独立构造函数及没有自然接收者的操作仍使用自由函数。
- 前端流水线为 `lexer → parser → CST → AST → HIR → TAST → Core → Mono → Lift → ANF → Go`。
- `gomlc/hir/` 负责 AST 到 HIR 的降级和名称解析。`gomlc/checker/` 负责推断、检查、约束、诊断、查询事实和类型驱动的决策，包括上下文枚举模式解析。`gomlc/tast/` 负责带类型的表示、共享类型和签名元数据、遍历、打印和验证。检查器依赖 TAST；生产用 TAST 和后端代码不得依赖检查器。
- 歧义、查找缺失、无效缓存和依赖失败必须产生可恢复的诊断。环境和查找代码必须返回失败值，而不能终止编译器。
- 只有 `lib/builtin/` 拥有编译器运行时 extern 和语言项。此契约必须与普通用户 Go FFI 区分；标准库算法使用 GoML 和编译器拥有的运行时钩子。用户 FFI 信任显式签名，将原生验证留给最终 Go 构建；外部 Go 类型是不透明的。受支持绑定见语言指南和 FFI 文档。
- LSP 功能应复用 `gomlc/query/`。保留文件作用域导入、规范包身份、依赖导航和现有文档分析缓存。
- ANF 连接点是通过尾位置的 `Jump` 到达的局部续延；递归连接点表示循环。保留词法作用域和跳转参数。修改控制流时，请查阅 [ANF 定义](gomlc/anf/model.goml)、[验证](gomlc/anf/verify.goml)和 [Go 降级](gomlc/go_backend/lower.goml)。
- Go 降级输出不含 `goto` 的结构化控制流；支持带标签的循环 break/continue。保留续延合并、递归循环处理和代码输出后的 [Go 死代码消除过程](gomlc/go_backend/dce.goml)。

<a id="packages-and-dependencies"></a>

## 包与依赖

- 模块通过根目录 `goml.toml` 的 `[module].path` 拥有规范路径；目录路径决定包身份。项目源文件声明所属包，同一目录中的文件共享包名。独立单文件命令可省略该声明。
- 导入具有文件作用域。声明的包名提供默认别名；使用 `as` 消歧。`module::path` 从模块根解析，`use alias::Trait` 将导入的 trait 带入方法调用作用域。
- 跨包 API 需要 `pub`。结构体字段和固有方法默认私有；trait 实现方法继承 trait 可见性，不得声明 `pub`。
- 接口和依赖环境公开公共 API，同时保留使用它所需的元数据。当前包代码生成必须保留私有辅助函数和完整内部环境。
- 可执行文件来自包含 `fn main()` 的 `package main` 包。直接子目录 `tests/` 是其父包的一个黑盒测试包，声明 `package tests`；不支持嵌套测试套件。
- 依赖只能放在模块根清单中。默认注册表为 `https://github.com/gomlang/registry.git`。其权威 `index.toml` 将模块映射到 Git 源。依赖使用 `true` 或本地路径；包源码跟随默认分支，不进行版本选择，也不使用 `goml.lock`。旧版版本字符串仅作为无版本迁移输入接受。
- 全局状态默认存放在 `~/.goml`，可由 `GOML_HOME` 覆盖。依赖源码位于其中的 `cache/registry`；生成的构建产物放在项目配置的目标目录中，默认为 `_artifact/`。

<a id="tests-and-golden-files"></a>

## 测试与黄金文件

- 优先使用快速、确定性的测试，以及覆盖相关解析、类型和运行时边界的最小用例。
- 需要包 world 的编译器 CLI 测试应使用 `gomlc/test_support/compiler.goml` 的 `compiler_command` 复用 `GOML_TEST_COMPILER_WORLD`，包括改变子进程工作目录时。仅在专门测试源码 world 构建或命令不接受 `--world` 时直接执行命令。
- 标准库行为测试位于相应的 `lib/std/<package>/tests/` 目录，声明 `package tests`。使用原生 GoML 断言和固定的已知答案数据；用户 Go FFI 仅用于专门的互操作测试。`just test-stdlib` 运行这些测试。
- 流水线用例位于 `gomlc/testdata/pipeline/NNN[_description]/main.goml`。添加或编辑源文件后，运行 `just update-golden` 生成 IR 快照和执行输出。
- 多包用例位于 `gomlc/testdata/module/projectNNN[_description]/`。包含根 `goml.toml`、显式包声明和导入、公开的跨包 API，以及 `package main` 入口。通过 `just update-golden` 生成其 `.out` 文件；这些用例不生成 IR 快照。
- 可见性和包诊断用例位于 `gomlc/testdata/module_diagnostics/`。
- 不得手动修改生成的黄金文件，包括 `.cst`、`.ast`、`.hir`、`.tast`、`.core`、`.mono`、`.lift`、`.anf`、`.go` 和 `.out`。只能使用 `just update-golden` 更新快照。
- builtin 契约改动通过 builtin 接口哈希使 prelude 和下游产物失效。prelude 改动影响其自身哈希；应重新生成受影响的产物和快照。
- 避免将测试或示例中的 trait 命名为 `Eq` 或 `Hash`，除非有意与内建 trait 消歧。
- 版本测试必须从版本模块推导预期值，不得硬编码发布版本号。

<a id="bootstrap-and-language-evolution"></a>

## 自举与语言演进

- `bootstrap/stage0.env` 是信任根：只有已发布的 Linux amd64 和 macOS arm64 Release 归档，且分别固定了 SHA-256 校验和，才能成为 stage0。两个平台必须使用相同发布版本。不得使用未发布的工作流产物。
- 当前 stage0 必须能够编译当前编译器和驱动程序源码。每项改动都必须保持这一不变式；根据上面的开发工作流选择本地检查。
- 在编译器或驱动程序源码使用新语法、内建功能、标准库 API、trait 或类型系统能力之前，先实现这些功能。测试和测试用例可以立即使用；自托管使用方必须等到发布并升级 stage0 后才能使用。
- 分两个阶段引入标准库能力：先交付公开源码、资源打包、依赖选择、导航、文档和外部测试，同时保留编译器拥有的回退实现；发布并升级 stage0 后，再迁移使用方并删除回退实现。
- 不兼容语法改动需要先发布同时接受新旧形式的过渡版本。升级 stage0，迁移自托管源码，再在后续版本中移除旧形式。
- 驱动协议、编译器 CLI 契约及其他自举接口至少需要一个发布版本的兼容重叠期，使旧 stage0 驱动程序能够使用新编译器。
- 只有当每个全新自举阶段都消费该阶段内生成的产物时，产物格式才可以破坏兼容性。缓存和外部依赖错误必须保持可恢复。
- 语法、语义或内建功能改动必须在同一改动中更新[语言指南](docs/goml.zh-CN.md)相关的说明、示例、限制、对照表和非正式文法。公开标准库 API 改动必须在同一改动中更新[标准库文档](docs/stdlib.zh-CN.md)相关的说明、示例、限制和常见任务表。

<a id="commits-and-releases"></a>

## 提交与发布

- 优先使用简短、祈使语气的 Conventional Commits（`feat:`、`fix:`、`refactor:`、`chore:`）。PR 应说明改动、链接相关 issue、描述行为变化并报告验证结果。
- 逻辑和生成的快照分为两个提交。如果同时新增测试，则分为三个提交：逻辑、测试源文件、由 `just update-golden` 生成的快照。
- 运行 `just bootstrap` 后，将 stage0 升级放在独立提交中；自托管使用方的迁移放在后续提交。
- 在早期自举阶段，只发布连续的 `0.1.x` 补丁版本，直到明确解除此限制。除此之外，兼容修复使用补丁版本，功能和 1.0 前的破坏性改动使用次版本，1.0 后的破坏性改动使用主版本。标签严格采用 `vX.Y.Z`，每次按 SemVer 连续推进一步。
- 遵循[发布指南](docs/releasing.zh-CN.md)：运行 `just set-version X.Y.Z` 和 `just ci`，推送发布提交，等待 main CI 在该准确提交上成功，然后创建并推送标签。
- `just set-version` 必须同步更新 `VERSION`、两个 GoML 版本模块，以及 VS Code 的 package 和 lockfile 版本。
- 合并或发布前 main CI 必须通过；main CI 负责不动点验证和完整测试套件。Release CI 要求带标签的准确提交已经通过 main CI，重新构建 stage2，并运行归档/LSP 冒烟测试，不重复完整套件。
- 发布后，从 `SHA256SUMS` 获取两个归档的校验和，运行 `just set-bootstrap-stage0 X.Y.Z <linux-sha256> <darwin-sha256>` 和 `just bootstrap`，然后提交 `bootstrap/stage0.env`，之后才可依赖新发布的能力。
