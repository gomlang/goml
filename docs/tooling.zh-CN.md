<a id="project-tools-and-workflows"></a>

# 项目工具与工作流

[English](tooling.md) | [简体中文](tooling.zh-CN.md)

[GoML 语言指南](goml.zh-CN.md)

模块配置、依赖、示例、构建输出、测试、编辑器支持和生成代码检查。源码级导入和可见性见[模块与包](language/packages.zh-CN.md)。

本页内容：[`goml.toml`](#gomltoml)、[包注册表](#package-registry)、[示例](#examples)、[原生依赖](#native-dependencies)、[构建输出布局](#build-output-layout)、[测试](#tests)、[LSP 与编辑器](#lsp-and-editor)、[验证生成代码](#verify-generated-code)。

<a id="gomltoml"></a>

## `goml.toml`

项目模块在根目录中声明规范路径：

```toml
[module]
path = "alice::myapp"

[build]
target-dir = "_artifact"

[dependencies]
"alice::http" = true
```

注册表依赖使用 `true` 声明存在（也接受空内联表）。包目前跟随其仓库的默认分支；没有包版本、版本约束、最低版本选择或 `goml.lock`。迁移期间仍可读取严格的 `X.Y.Z` 字符串和内联 `version` 字段，但它们不选择或约束来源。新声明应使用 `true` 或本地 `path`。

`[dev-dependencies]` 使用相同的无版本和路径形式。这些依赖可用于调用模块的 `*_test.goml` 文件、`tests/` 包和根目录 `examples/` 树。生产源码即使在测试构建期间也不能导入它们。其他模块的依赖不会继承其开发依赖。测试和示例命令解析普通依赖与开发依赖的并集；普通 check/build/run 命令只解析普通依赖。开发辅助模块可以反向依赖当前库：该依赖边使用当前源码，不引入第二份副本。生产依赖循环仍是错误。导入被测包的辅助模块必须用于黑盒测试或示例；从同包白盒测试导入它仍会形成不受支持的包循环。

```toml
[dev-dependencies]
"alice::proptest" = true
"alice::test_support" = { path = "../test_support" }
```

使用 `goml add alice::proptest --dev`、`goml add alice::test_support --dev --path ../test_support` 和 `goml remove alice::proptest --dev` 编辑此节。生产依赖仍放在 `[dependencies]`。

本地开发依赖可使用内联表。路径相对于声明模块，依赖名必须匹配目标的规范 `[module].path`：

```toml
[dependencies]
"alice::http" = { path = "../http" }
"alice::model" = { path = "../model" }

[replace]
"alice::codec" = { path = "../codec" }
```

`goml add alice::http --path ../http` 验证目标并写入相对于模块的路径。本地模块使用工作源码，包括传递性的本地依赖，无需注册表缓存。注册表模块不能声明本地路径依赖。只有调用模块的 `[replace]` 表作用于整个依赖图；依赖模块自己的替换会被忽略。本地来源冲突、身份不匹配和依赖循环均为错误。产物保存在消费项目的目标目录中，本地源码修改会使缓存输出失效。

`goml.work` 文件将独立模块组织到一起：

```toml
[workspace]
members = ["app", "http", "model"]

[replace]
"alice::codec" = "../codec"
```

成员保留自己的清单和文件作用域导入。成员资格为已声明依赖提供本地来源，不会让未声明依赖变得可见。选择顺序为工作区替换、调用模块替换、工作区成员、显式路径依赖，最后是注册表来源。成员路径和替换路径相对于 `goml.work`；重复的成员规范路径或模块身份会被拒绝。祖先目录中的工作区仅作用于列出的成员。`GOML_WORKSPACE` 显式选择工作区目录，`GOML_WORKSPACE=off` 禁用发现。

`goml check --workspace`、`build --workspace`、`test --workspace`、`fmt --workspace`、`doc --workspace` 和 `clean --workspace` 按声明顺序对每个成员执行，失败即停止。从不属于任何模块的工作区目录执行时，这些命令自动选择所有成员。`-p alice::app` 或 `--package alice::app` 选择一个模块；有多个成员时，`goml run` 要求此选择。在成员内部执行的命令只操作该成员，除非提供 `--workspace`。编译器查询和 LSP 依赖导航解析相同的本地来源与替换。

可以省略 `[build]`；`build.target-dir` 默认为模块根目录下的 `_artifact`。清单值必须是非空相对路径，且不能包含 `..` 段。`goml check`、`goml build`、`goml run` 和 `goml test` 可以用 `--target-dir <path>` 临时覆盖；命令行覆盖可为相对或绝对路径。`goml clean` 删除配置的目标目录，其可选的 `--target-dir <path>` 覆盖必须仍位于模块内部。

模块路径每段必须非空，可包含 ASCII 字母、数字、`_` 和 `-`。以 `builtin` 或 `prelude` 为根的路径保留给工具链，不能用作模块路径或依赖。`[module]` 节目前没有 `name`、`kind`、`root` 或类似字段。

<a id="package-registry"></a>

## 包注册表

默认中央索引为 [github.com/gomlang/registry](https://github.com/gomlang/registry)。其 `index.toml` 将规范模块名映射到 `gomlang` 组织下的 Git 仓库。模块的规范名来自根清单，因此 `ecosystem::ansi` 解析为 `github.com/gomlang/ansi`，而导入路径不变。

```sh
goml add ecosystem::ansi
goml add ecosystem::proptest --dev
goml update
goml remove ecosystem::ansi
```

`goml add` 写入无版本依赖并获取所需源码。`goml update` 刷新索引以及当前模块的普通和开发依赖闭包，并遵循本地路径和工作区替换。在模块外执行时仅更新索引。不获取依赖模块自身的开发依赖。构建和编译器/LSP 查询使用缓存源码，不访问网络；缺失缓存是可恢复错误，会提示用户运行 `goml update`。源码内容更新会使生成产物失效。

默认主目录为 `~/.goml`，可用 `GOML_HOME` 覆盖。索引检出位于 `<home>/cache/registry`；Git 包检出位于该目录下的 `sources/<owner>/<module>/`。每个检出跟随默认分支。更新使用快进拉取，不重置本地修改。更改缓存仓库的来源会产生诊断，要求用户先删除该缓存目录再更新。

Git 索引条目如下：

```toml
[modules."ecosystem::ansi"]
git = "https://github.com/gomlang/ansi.git"
```

本地注册表测试夹具可以改用相对于索引根目录的 `path = "owner/module"`。`git` 和 `path` 必须且只能指定一个。路径必须位于注册表树内。每个来源都必须包含根 `goml.toml`，且模块路径与条目匹配。不存在版本目录或合成快照版本。

在 `<home>/config.toml` 中设置 `[registry].default` 可覆盖索引 URL；缺失或为空时使用中央索引。`goml update --local-registry <path>` 和 `goml add ... --local-registry <path>` 使用本地 Git 索引。`goml remove` 只编辑清单，不联系索引。不支持以前的带版本索引格式；切换到此索引前请使用新缓存或删除旧注册表缓存。依赖声明和索引条目属于配置，不改变源语言语法。

<a id="examples"></a>

## 示例

具名示例共享模块根清单，可使用普通和开发依赖。将每个可执行示例放在 `examples/<name>/main.goml`，声明 `package main;` 和 `fn main() -> ()`。名称包含 ASCII 字母、数字、`_` 或 `-`。示例可包含辅助包、白盒测试和直接的 `tests/` 包。包含自己 `goml.toml` 的示例目录是独立模块，不是具名示例。

```text
library/
├── goml.toml
├── library.goml
├── tests/api_test.goml
└── examples/basic/
    ├── main.goml
    └── tests/example_test.goml
```

对于模块路径为 `alice::library` 的库，示例包为 `alice::library::examples::basic`。生产包和其他模块不能导入示例包。示例使用现有 GoML 源语法，不需要目标声明或额外清单。

```sh
goml run --example basic -- argument
goml build --example basic
goml check --examples
goml test --example basic
goml test
```

`check`、`build` 和 `test` 接受 `--example <name>` 或 `--examples`，分别仅选择具名示例或所有示例。`run` 接受一个 `--example`，且不能同时指定位置目标。普通 check/build 命令排除根 `examples/` 树。普通 `goml test` 运行模块测试、构建所有示例并运行其测试，但不执行示例的 `main` 函数。`goml fmt` 包含示例及其测试。`basic` 的二进制文件为 `<target-dir>/bin/examples/basic/basic`。测试进程均从模块根目录运行，包括示例测试。

<a id="native-dependencies"></a>

## 原生依赖

包含 Go 适配器的模块可以在同一清单中声明 Go 模块和系统先决条件：

```toml
[native]
go-module = "example.com/goml-ecosystem/llvm"
go-version = "v0.0.0"
cgo = "required"
llvm-major = "18"
```

所有值均为带引号字符串。`go-module` 必须匹配适配器模块根目录的 `go.mod`；`go-version` 默认为 `v0.0.0`，可指定适合带版本模块路径的 Go 模块版本。`cgo` 接受 `required` 或 `optional`（默认值）。`llvm-major` 要求正的十进制主版本号。未知或重复的 native 键均为错误。

消费项目提供最小的模块根 `go.mod`。对声明了原生适配器的模块，驱动程序从选定的 GoML 注册表来源解析适配器，并在 `<target-dir>/native/<hash>/` 下生成 Go 依赖要求和本地替换。通过 `GOFLAGS=-modfile=...` 将生成的模块文件传给编译、测试和 Go 导出。此托管模块文件路径不会修改源 `go.mod` 和 `go.sum`。内容哈希包括适配器位置和模块元数据；试运行不会探测先决条件或创建这些文件。已有的 `GOFLAGS -modfile` 与托管依赖冲突，会产生诊断。

驱动程序检查所需的 `CGO_ENABLED=1`，并依次通过 `LLVM_CONFIG`、`llvm-config-N`、`/usr/lib/llvm-N/bin/llvm-config`、`llvm-config` 检查请求的 LLVM 主版本。这些声明不会安装系统包或配置编译器/链接器搜索路径。需要时消费方仍须配置 `CGO_CFLAGS`、`CGO_LDFLAGS` 和共享库加载器。额外的 Go 依赖由消费方 Go 模块和用户的 Go 环境管理；依赖模块自己的替换不会提升到消费方。普通项目构建使用正常的 Go 依赖解析。Go 库导出和原生 SIMD 构建保留只读、离线的模块处理，要求事先准备依赖。注册表包必须包含适配器源码和 Go 模块文件。

<a id="build-output-layout"></a>

## 构建输出布局

默认产物布局如下；其中规范包路径按目录层级展开，例如 `alice::myapp::utils` 对应 `alice/myapp/utils`：

```text
_artifact/
├── check/
│   ├── pkg/<canonical-package-path>/<last-path-segment>.interface
│   └── deps/<owner>/<module>/pkg/<canonical-package-path>/<last-path-segment>.interface
├── build/
│   ├── pkg/<canonical-package-path>/<last-path-segment>.interface
│   ├── pkg/<canonical-package-path>/<last-path-segment>.core
│   ├── pkg/<entry-package-path>/goml_generated.go
│   └── deps/<owner>/<module>/pkg/<canonical-package-path>/<last-path-segment>.*
├── test/
│   ├── base/
│   │   ├── pkg/<canonical-package-path>/<last-path-segment>.*
│   │   └── deps/<owner>/<module>/pkg/<canonical-package-path>/<last-path-segment>.*
│   ├── internal/
│   │   ├── pkg/<canonical-package-path>/<last-path-segment>.*
│   │   ├── goml_generated.go
│   │   ├── tests.json
│   │   └── runner
│   └── external/
│       ├── pkg/<black-box-test-package-path>/<last-path-segment>.*
│       ├── goml_generated.go
│       ├── tests.json
│       └── runner
└── bin/
    ├── <module-name>
    └── <module-relative-entry-path>/<entry-name>
```

`<last-path-segment>` 是规范包路径的最后一段，不是源码中声明的包名；例如，声明为 `package main` 的 `alice::app::cmd::server` 仍使用 `server.core` 和 `server.interface`。

`check` 只需要接口产物；`build` 和 `test` 还生成 Core。每个声明为 `package main` 的入口包在自己的产物目录生成名为 `goml_generated.go` 的 Go 链接文件；普通库包不生成 Go 链接文件。包路径以 `_test`、操作系统名或架构名结尾时，固定文件名不会触发 Go 的特殊文件规则。

接口和 Core 产物使用确定性的 GAF 二进制容器。GAF
载荷直接从编译器数据编码，不经过中间 JSON
值树。

遍历内置 `range`、`Vec`、`Slice` 和 `MutSlice` 值的 `for` 循环直接降级为索引循环，不分配迭代器闭包。

模块根入口包的可执行文件为 `bin/<module name>`；嵌套入口包保留模块内目录并附加入口名。例如，`alice::app::cmd::server` 的输出为 `bin/cmd/server/server`。生产测试依赖在 `test/base` 下只构建一次；内部和外部测试分别在对应测试种类根目录共享一个 Go 入口、清单和运行器。运行器文件在 Windows 上带 `.exe` 后缀。外部依赖在各阶段根目录使用相同的 `deps/<owner>/<module>/pkg/...` 结构。

配置的产物目录不参与包发现，也不能作为检查、构建或测试目标。`goml new` 生成的 `.gitignore` 默认包含 `/_artifact/`；修改 `build.target-dir` 后应同时修改项目自身的 `.gitignore`。

<a id="tests"></a>

## 测试

<a id="test-functions-and-attributes"></a>

### 测试函数与属性

使用 `#[test]` 将顶层函数标为测试：

```goml
package math;

use std::testing;

#[test]
fn addition_works() -> () {
    testing::assert_eq(1 + 1, 2)
}
```

测试函数必须满足：

- 只能是顶层函数，不能是 impl 方法、`extern fn`、结构体、枚举或 trait；
- 不能命名为 `main`；
- 不能有参数或类型参数；
- 返回类型必须为 `()`；
- 同一测试包内生成的规范测试 ID 必须唯一。

测试函数不要求 `pub`。`#[test]` 不接受参数。需要默认跳过测试时，可用无参数 `#[ignore]` 或带字符串原因的 `#[ignore("reason")]`；`#[ignore]` 不能脱离 `#[test]` 单独使用：

```goml
#[test]
#[ignore]
fn unfinished_case() -> () {}

#[test]
#[ignore("requires an external service")]
fn integration_case() -> () {}
```

参数化测试在顶层函数上使用一个或多个 `#[test_case(...)]` 属性。测试用例接受字符串和布尔实参，并对照函数参数检查。每个用例获得内容派生的稳定 ID，而列表、过滤、文本/JSON 报告、产物清单和 CodeLens 仍使用可读显示名。普通零参数 `#[test]` 函数不变。

```goml
#[test_case("left", true)]
#[test_case("right", false)]
fn is_left(value: string, expected: bool) -> () {
    testing::assert_eq(value == "left", expected)
}
```

`std::testing` 中的断言和载荷提取辅助函数见[测试断言](stdlib.zh-CN.md#testing-assertions)。

<a id="white-box-and-black-box-tests"></a>

### 白盒与黑盒测试

白盒测试文件与被测源码位于同包目录。文件名必须以 `_test.goml` 结尾，并声明相同包名：

```text
math/
├── math.goml
└── math_test.goml
```

`math_test.goml`:

```goml
package math;

use std::testing;

#[test]
fn private_helper_works() -> () {
    testing::assert_eq(private_helper(), 42)
}
```

测试构建将生产源码和同目录所有 `*_test.goml` 合并为一个包，因此白盒测试可访问包的私有顶层项。对生产目标执行普通 `goml check`、`goml build` 或 `goml run` 时，这些文件不参与编译。

黑盒测试位于被测包的 `tests/` 目录。整个目录构成名为 `tests` 的包。必须显式导入被测包，且只能访问其公开 API；不能在 `tests/` 下创建嵌套测试套件：

```text
math/
├── math.goml
└── tests/
    ├── api_test.goml
    └── smoke_test.goml
```

`math/tests/api_test.goml`:

```goml
package tests;

use alice::myapp::math;
use std::testing;

#[test]
fn public_add_works() -> () {
    testing::assert_eq(math::add(20, 22), 42)
}
```

若被测包身份为 `alice::myapp::math`，上例测试包的规范身份为 `alice::myapp::math::tests`。普通包发现排除 `tests` 目录，生产包不能导入黑盒测试包。

编译器仓库中，标准库行为测试位于 `lib/std/<package>/tests/`，使用原生 GoML 断言或固定已知答案数据。从仓库根运行 `just test-stdlib`。对这个源码模块，`goml test` 先从同级 `builtin`、`prelude` 和 `std` 源码在选定目标目录下构建私有编译器世界，再针对它编译黑盒测试，从而测试当前库源码。目前不支持标准库白盒测试，会产生诊断。安装的工具链不包含这些测试源码和数据。

<a id="check-and-run-tests"></a>

### 检查与运行测试

`goml check`、`goml build` 和 `goml test` 从当前目录向上发现模块。除非提供示例选择器，否则处理完整模块，不接受包或文件目标。`goml check` 默认只检查生产源码；使用 `--tests` 会在生产包检查成功后检查所有白盒和黑盒测试：

```sh
goml check
goml check --tests
goml build
goml fmt
goml fmt --check
```

测试源码不能用于修复自身检查失败的生产包。

`goml fmt` 格式化模块源码；`goml fmt --check` 报告格式差异而不写入文件。源码发现规则和选项见[格式化命令](formatting.zh-CN.md#command-line)。

运行测试：

```sh
goml test [FILTER]
```

`FILTER` 对完整测试显示名做子串匹配，例如 `goml test addition`。常用选项：

- `--kind internal|external|all`：只运行白盒、只运行黑盒或所有测试，默认为 `all`；
- `--list`：列出匹配测试，不运行；
- `--ignored`：只运行被忽略测试；
- `--include-ignored`：同时运行普通测试和被忽略测试；
- `--jobs N`：并行工作进程数，默认为 `1`；
- `--seed N`：为每个测试提供可复现的正数种子；
- `--timeout 500ms|30s|2m`：单个测试超时，默认为 `30s`；
- `--nocapture`：让测试的标准输出和标准错误直接继承终端；
- `--format text|json`：选择文本或逐行 JSON 结果；JSON 不能与 `--nocapture` 同用；
- `--target-dir`、`--dry-run` 和 `--compiler`：含义与项目构建命令相同。

每个测试在独立运行器进程执行，异常退出和超时不影响其他测试；`--jobs` 控制同时运行的测试进程数。使用 `goml test` 构建运行器需要可用 Go 工具链。

`goml test --no-run` 构建所有选定测试组但不执行。`goml test --no-build` 执行目标目录现有的测试运行器和清单，不编译或链接，可与过滤、`--list` 和分片结合。预构建测试产物必须来自同一源码修订，且使用执行机器的操作系统和架构；`--no-build` 有意跳过新鲜度检查，因此缺失产物会失败而非重建。传输产物时，保留各测试组的 `runner`、`tests.json` 路径及所有相对于可执行文件的运行时库。`--no-build` 与 `--no-run`、`--dry-run` 冲突；`--no-run` 与 `--list` 冲突。

`GOML_TEST_SHARD=INDEX/COUNT` 将匹配测试分配给独立机器。两个数均为最多九位、从 1 开始的正十进制整数，`INDEX` 不能超过 `COUNT`。应用名称和忽略测试过滤后，驱动按显示名排序并以轮询顺序分配到分片。例如，在四台机器分别运行 `GOML_TEST_SHARD=1/4 goml test --jobs 4` 到 `GOML_TEST_SHARD=4/4 goml test --jobs 4`。`--list` 使用相同分片选择。未设置或空值运行全部选择；无效值产生诊断。驱动在测试进程中清除此变量，让嵌套 `goml test` 运行完整选择。

`goml test --seed N` 通过 `testing::seed()` 向每个测试进程提供一个可复现的正数种子。失败摘要打印重放种子，每个 JSON 结果事件都包含它。捕获的 stdout 和 stderr 按测试进程隔离，除非选择 `--nocapture`，否则仅失败时输出。

<a id="lsp-and-editor"></a>

## LSP 与编辑器

项目配置和模块依赖解析由 `goml` 驱动负责。LSP 通过 `goml __project-info --source <PATH>` 请求带版本项目描述，并连同打开文档内容传给编译器查询引擎。描述提供规范模块身份、源码目录、分开的生产/开发依赖边及配置监视路径。此操作读取本地文件和缓存依赖，不编译、获取或写入构建产物。依赖失败保持为诊断；可用根元数据被保留以供本地包补全。

项目描述客户端验证协议版本 1，在所监视配置内容和规范路径不变时缓存成功描述。普通源码编辑复用项目解析，并仍参与编译器源码分析指纹。创建工作区、依赖清单变化、替换和注册表索引变化触发重新解析。缺失依赖和无效描述在编辑后重试，不缓存为成功项目。编译器仍负责源码导入解析、可见性和类型检查。

`GOML_PROJECT_DRIVER` 显式选择驱动。否则客户端使用同级 `goml`、`GOML_HOME/bin/goml` 或 `PATH`；仓库测试可提供 `GOML_TEST_GOML`。驱动和语言服务器必须支持相同项目描述协议。扩展捆绑两种工具，并提供 `goml.driverPath` 用于自定义安装。现有编译器构建命令和驱动协议不变，因此固定 stage0 可继续构建工具。

LSP 在各自包上下文中分析生产文件、内部和黑盒测试，因此诊断、补全、悬停及跳转定义遵循与编译相同的可见性规则。普通和参数化测试用例上显示测试 CodeLens。VS Code 扩展保存文件后，以所选测试显示名和种类调用模块级 `goml test`。

枚举补全根据期望类型提供上下文构造器，并根据被匹配表达式的类型提供 `match` 模式建议。变体建议显示载荷类型，并为元组或具名字段提供占位符；新匹配模式还提供分支体占位符。现有载荷和分支体会保留。无守卫的单元变体分支会排除对应变体建议。编辑文档期间，补全继续复用已检查的类型。

在 VS Code 中设置 `"goml.lsp.enabled": false`，即可使用语法高亮和括号匹配而不启动语言服务器，也无需安装工具链。此设置立即生效；重新启用会再次启动服务器。

类型内嵌提示使用不带包限定的短名，类型长度预算为 25 字符，不含前导 `: `。嵌套类型和长名用 `…` 缩写；悬浮提示保留完整推断类型和包路径。

包分析在查询上下文内缓存解析/降级源码和成功检查的包。源码内容、包成员、规范模块身份和传递依赖指纹决定复用。编辑包会重查该包及依赖它的包，独立包仍可复用；依赖中的函数体编辑也保守地使依赖方失效。打开文档覆盖参与相同指纹。外部 Go 类型在缓存和新分析中都保持不透明；原生验证属于最终 Go 构建。监视的源码和清单变化使文档分析缓存失效；VS Code 扩展还监视 `goml.work`、Go 源码、`go.mod` 和 `go.sum`。

自定义 `goml/expandedDerive` 请求返回所请求文档完成内置和可编程 derive 后的格式化 AST。它使用与 `goml check` 相同的包别名、显式 derive 导入、歧义检查、依赖接口、CTIR 验证器和资源限制。VS Code 命令 `GoML: Show Expanded Derive` 在源文件旁打开结果。

`textDocument/formatting` 的行为见[编辑器格式化](formatting.zh-CN.md#editor-formatting)。

CLI 和 LSP 诊断共享包含可选稳定代码、备注、帮助和建议编辑的模型。CLI 显示 `error[typer/try_error_mismatch]` 或 `warning[typer/unused_result]`；LSP 在 `Diagnostic.code` 暴露相同标识符，并在 `relatedInformation` 暴露次级位置。备注和帮助出现在显示消息中，也单独放在诊断 `data`，连同修复和文档版本。修复包含标题、一项或多项范围编辑，以及 `machine-applicable` 或 `requires-review` 适用性。代码操作拒绝旧文档版本的修复。目前为未使用 Result 警告及三种 [`?` 诊断](language/basics.zh-CN.md#try-operator)分配代码；其他诊断可能省略。

<a id="verify-generated-code"></a>

## 验证生成代码

运行 `just make` 后，从仓库根验证独立源文件：

```sh
stage2/bin/gomlc run-single path/to/main.goml
```

对于项目，在模块内任意位置运行已安装的驱动：

```sh
goml fmt
goml check
goml build
goml check --tests
goml test
goml run
```

`goml check`、`goml build`、`goml test` 和 `goml fmt` 操作完整模块，不接受包或文件目标。工作区中，用 `--workspace` 选择所有成员，或 `-p <module-path>` 选择一个成员。模块有多个可执行包时，`goml run [TARGET]` 接受可选入口包文件或目录。

需要检查编译阶段时，在 `gomlc run-single` 后加 `--dump-ast`、`--dump-expanded-ast`、`--dump-hir`、`--dump-tast`、`--dump-ctir`、`--dump-core`、`--dump-mono`、`--dump-lift`、`--dump-anf` 或 `--dump-go`。`--dump-ast` 展示 derive 展开前的源码降级，`--dump-expanded-ast` 则包含所有生成实现。

修改仓库时，按[开发指南](development.zh-CN.md)运行相关检查。
