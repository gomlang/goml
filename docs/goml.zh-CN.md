<a id="goml-language-guide"></a>

# GoML 语言指南

[English](goml.md) | [简体中文](goml.zh-CN.md)

本指南介绍本仓库实现的 GoML 语法、类型规则、工具和内置 prelude。标准库 API 见[标准库参考](stdlib.zh-CN.md)。编写、修改和审阅 `.goml` 源码时请以本指南为准。与 Rust、Go 或 OCaml 相似，并不意味着支持这些语言的语法或 API。

GoML 是一种具有垃圾回收的静态类型语言。语法接近 Rust，语义更接近 ML。编译器在生成 Go 代码前，对泛型进行单态化并对 GoML 闭包进行闭包提升。显式 Go FFI 可以保留原生泛型类型和函数值。GoML 没有所有权、借用、生命周期或手动内存管理。

生成的 Go 代码和 Go FFI 以 Go 1.26 为兼容性基线。构建工具链和编译生成的程序需要 Go 1.26 或更高版本。原生 FFI 签名由最后的 Go 编译步骤检查。独立的 [gomlgo 前端和解释器](https://github.com/gomlang/gomlgo) 以 Go 1.26 的语言和标准库行为为目标；其执行和差分测试需要 Go 1.26.x。

快速导航：[包](#modules-packages-and-imports)、[类型](#type)、[控制流](#control-flow)、[模式](#patterns)、[trait](#trait-and-impl)、[编译期求值](#compile-time-evaluation)、[Go FFI](#go-ffi)、[测试](#tests)、[API 文档](#comments-and-api-documentation)、[prelude](#built-in-prelude)、[标准库](stdlib.zh-CN.md)和[验证](#verify-generated-code)。

构建和安装说明见[开发指南](development.zh-CN.md)和[发布指南](releasing.zh-CN.md#installation)。源码的规范布局见[格式化](formatting.zh-CN.md)，CTIR 内部机制见[编译期求值架构](comptime.zh-CN.md)。

<a id="essential-rules"></a>

## 基本规则

1. 源文件使用 `.goml`，包内测试文件使用 `*_test.goml`。旧的 `.gom` 文件必须在编译或格式化前重命名；项目根目录通过 `goml.toml` 声明规范模块路径。
2. 项目中的每个源文件先写 `package name;`，再写该文件自己的 `use`，最后写顶层定义。
3. 顶层函数的参数必须标注类型；省略返回类型时固定为 `()`，不会从函数体推断。
4. 泛型使用方括号：`Vec[i32]`、`fn id[T](x: T) -> T`，不使用 `<...>`。
5. 泛型调用通常依赖类型推断；需要显式类型实参时，写 `id::[i32](1)`，不写 `id[i32](1)` 或 Rust 的 `id::<i32>(1)`。
6. `if`、`match` 和 `select` 都是表达式。没有 `else` 的 `if` 必须返回 `()`；`match` 必须穷尽所有情况。
7. 块末尾没有分号的表达式是块的值；加上分号会丢弃该值。直接丢弃内置 `Result` 会产生警告；使用 `let _ = ...` 显式忽略。
8. `let` 和赋值语句必须以分号结尾。可变绑定可以通过 `let mut pattern` 引入，也可以在模式中用 `mut name` 精确指定；`if`、`match`、`select` 以及无标签的 `while`、`loop` 和 `for` 语句可以省略分号。有标签的循环语句后若还有语句，必须加分号。
9. 构造枚举可以使用 `Option::Some(value)` 这样的完整名称；预期类型能确定枚举时，也可以使用 `Some(value)` 和 `None` 这样的上下文名称。模式也允许上下文变体名称。
10. 跨包调用写作 `alias::item`。顶层项、结构体字段和固有方法都必须按需要标注 `pub`。
11. 跨包使用 trait 方法语法前，先通过 `use alias::Trait;` 或花括号导入导入包和 trait；不确定时使用 UFCS：`Trait::method(value)`。
12. 当前模块根目录下的包使用 `module::path`。不要生成 `mod`、`crate::`、`self::`、`super::`、根路径 `::x`、Rust 引用或 Go 的 `var` / `:=`。用户的 `extern fn` 只有配合下文所述的类型化 Go FFI 属性才有效。
13. 测试函数使用 `#[test]`，必须没有参数、没有类型参数且返回 `()`；白盒测试放在同包的 `*_test.goml` 中，黑盒测试放在被测包的 `tests/` 目录中。

<a id="minimal-program"></a>

## 最小程序

单文件编译允许省略 `package main;`：

```goml
fn main() -> () {
    println("hello, goml")
}
```

项目中的入口文件应显式声明包：

```goml
package main;

fn main() -> () {
    println("hello, goml")
}
```

`main` 不能有参数或类型参数。生成可执行文件时，选作入口的包必须声明 `package main;` 并定义 `fn main()`。建议让 `main` 返回 `()`。

<a id="lexical-rules"></a>

## 词法规则

<a id="comments-and-api-documentation"></a>

### 注释和 API 文档

行注释以 `//` 开始，到下一个换行结束。源代码行首的 `///`（不包括 `////`）为后续声明提供文档。连续的文档行组成一个块；空行或普通注释会中断附着关系。属性属于后续声明，可以放在文档与声明关键字之间。

`package` 声明前的 `//!` 注释描述整个包。包有多个源文件时，`gomldoc` 按相对文件路径顺序合并这些描述。文档注释仍是词法分析器中的普通附属信息，不影响程序的类型检查或执行。

```goml
//! Small arithmetic helpers.
package arithmetic;

/// Add one to `value`.
pub fn increment(value: isize) -> isize {
    value + 1
}
```

`goml doc` 为当前模块的所有生产包生成离线 HTML，默认输出到 `<target-dir>/doc/`。`goml doc --format json` 将带版本的 API 模型写入该目录的 `module.json`。公开声明和成员遵循语言可见性规则；`--document-private-items` 可包含私有 API。命令在替换现有输出前会检查模块并报告依赖或类型错误。

文档支持段落、标题、单层项目列表、行内代码、围栏代码块和链接。`[Name]` 和 `[label](goml:alias::Name)` 根据源文件的导入引用 API 声明。普通链接支持 HTTP、HTTPS、mailto 和片段目标。原始 HTML 会被转义。无法解析或存在歧义的 API 链接产生警告；生成模块之外的目标显示为文本。不支持完整 Markdown、自动生成依赖文档站点或可执行文档测试。

独立 `gomldoc` 命令、输出模式和安装说明见[文档生成](documentation.zh-CN.md)。

<a id="identifiers-and-keywords"></a>

### 标识符和关键字

普通标识符仅使用 ASCII，形式如下：

```text
[A-Za-z][A-Za-z_0-9]* | _[A-Za-z_0-9]+
```

单独的 `_` 是通配符，不是普通变量名。当前语法对各类名称的首字母大小写有以下要求：

- 包名、`use package as alias` 中的包别名、函数、方法、参数、局部绑定和 GoML 字段声明必须以小写字母或 `_` 开头；导入项的别名可以遵循被导入项的命名规则；
- 结构体、枚举、trait、枚举变体、泛型参数和关联类型必须以大写字母开头；
- [可移植 SIMD](stdlib.zh-CN.md#portable-simd) 中列出的定宽向量和掩码类型写法，包括 `u8x16`、`i32x8`、`f32x4`、`f64x4` 和 `mask32x4`，也可以用作结构体名、类型别名和类型位置中的名称。它们是由 `std::simd` 提供的普通作用域名称，不是全局保留关键字；
- 路径中的名称保留其所引用名称应有的大小写。

构造枚举和编写模式时，如果预期类型能确定变体所属枚举，优先使用不带限定的变体名称。类型上下文不足或普通名称解析选择了其他声明时，使用 `Enum::Variant`。操作已有值时优先使用 `value.to_string()` 这样的接收者方法；访问普通向量元素时优先使用 `values[i]` 这样的索引。

常用关键字包括：

```text
package use as pub fn struct enum trait impl for type const static where defer
let mut if else match while loop for in break continue return go dyn
true false bool isize i8 i16 i32 i64 usize u8 u16 u32 u64
f32 f64 string char extern
```

`()` 是空元组类型唯一的写法。

`self` 是接收者参数的特殊简写，也可以作为普通的接收者变量名；`Self` 在 trait 和 impl 的类型位置中有特殊含义。

空白分隔词法单元，但不决定块结构。仅支持从 `//` 到行尾的行注释，不支持块注释。

<a id="literals"></a>

### 字面量

| 类别 | 语法示例 | 默认类型或说明 |
| --- | --- | --- |
| 空元组 | `()` | 类型为 `()` |
| 布尔值 | `true`, `false` | 类型为 `bool` |
| 整数 | `0`, `42`, `1_000`, `0b1010`, `0o755`, `0xff` | 由上下文决定；无约束时默认为 `isize` |
| 浮点数 | `1.25`, `1e3`, `2.5e-2` | 由上下文决定；无约束时默认为 `f64` |
| 字符串 | `"text"` | 类型为 `string` |
| 原始字符串 | `r"text"`, `r#"text with \"quotes\""#` | 类型为 `string`；不处理转义 |
| 字节字符串 | `b"text\\n"` | 类型为 `Vec[byte]`；支持 ASCII 内容和字节转义 |
| 原始字节字符串 | `br"text"`, `br#"text with "quotes""#` | 类型为 `Vec[byte]`；不处理转义 |
| 插值字符串 | `f"value={value}"` | 类型为 `string`；嵌入的值使用 `ToString` |
| 字符 | `'a'`, `'\n'`, `'\u0041'` | 类型为 `char`，表示 Unicode 标量值 |
| 字节 | `b'A'`, `b'\n'`, `b'\xFF'` | 类型为 `byte`，是 `u8` 的透明别名 |

数字没有类型后缀。整数字面量支持二进制 `0b`/`0B`、八进制 `0o`/`0O`、十进制和十六进制 `0x`/`0X` 形式。两个数字之间可以使用 `_` 分隔；浮点数支持 `e`/`E` 指数和可选的指数符号。使用小数点时，两侧都必须有数字。负数由一元 `-` 和正数字面量组成。

浮点数字面量直接从源码文本舍入到上下文选择的 `f32` 或 `f64` 宽度，采用舍入到最近值、相等时取偶数的规则。溢出的字面量会被拒绝；下溢可能舍入为零。

无后缀数字可以从上下文获得宽度：

```goml
let small: u8 = 42;
let ratio: f32 = 0.5;
let values: [i16; 3] = [1, 2, 3];
```

字符串支持 `\"`、`\\`、`\n`、`\r`、`\t`、`\b`、`\f`、`\/` 和四位 `\uXXXX` 转义；字符使用相同的控制字符转义，并用 `\'` 表示单引号。普通字符串不能跨行。

转义后的 Unicode 字符在生成 Go 代码时保留其 UTF-8 值。例如，`"\uFEFF"` 包含一个字符和三个字节，嵌入其他字符串时也一样。

原始字符串使用 `r"..."` 或匹配的井号分隔符，如 `r#"..."#` 和 `r##"..."##`。内容可以跨行，反斜杠、引号、花括号和换行均原样保留。不处理转义和插值。只有引号后紧跟与起始分隔符数量相同的 `#` 时，才结束字面量：

```goml
let windows_path = r"C:\users\alice";
let quoted = r#"She said "hello"."#;
let multiline = r##"first line
second line"##;
```

插值字符串使用 `f"..."`。每个 `{expression}` 从左到右求值一次，并通过 `ToString` 转换；已经是 `string` 的值直接插入。字面花括号写作 `{{` 和 `}}`。不支持格式说明符和插值多行字符串：

```goml
let name = "Ada";
let count = 3;
let message = f"{name} has {count} items {{ready}}";
```

字节字面量恰好包含一个 ASCII 字节，支持 `\'`、`\\`、`\0`、`\b`、`\f`、`\n`、`\r`、`\t` 和两位 `\xNN` 转义。非 ASCII 内容会被拒绝。

字节字符串使用 `b"..."`，没有预期类型时类型为 `Vec[byte]`。在 `[byte; N]` 上下文中生成定长数组，且长度必须匹配。向量字节字符串编译为一次底层切片分配和批量字面量初始化，而不是逐字节调用 `push`。普通字节字符串接受 ASCII 内容和 `\"`、`\\`、`\0`、`\b`、`\f`、`\n`、`\r`、`\t`、`\/` 及两位 `\xNN` 转义。Unicode 转义和非 ASCII 源字符会被拒绝：

```goml
let request = b"GET / HTTP/1.0\r\n\r\n";
let marker = b"\x00\xFF";
```

原始字节字符串使用 `br"..."` 或匹配的井号分隔符，如 `br#"..."#`。内容可以跨行，并按 ASCII 字节原样复制，不处理转义或插值：

```goml
let path = br"C:\tmp\file";
let quoted = br#"say "hello""#;
```

标准库类型 `std::bytes::Bytes` 仍是独立的缓冲区抽象，提供只读和可变零拷贝视图及可增长字节构建器。`std::bytes::endian` 提供带检查的端序访问，以及配套的可增长和定长缓冲区读写器。API 要求该类型时，使用 `bytes::Bytes::from_vec(value)`。

多行字符串的每一行以两个反斜杠开头；标记前的缩进会移除，各行以换行连接，标记后的内容原样保留。结束后的下一行不再以两个反斜杠开头，通常直接写 `;` 或 `}`。格式化器将多行字符串放在独立的缩进行中：

```goml
fn poem() -> string {
    let text =
        \\roses are red
        \\violets are blue
        \\"quotes" need no escaping here
    ;
    text
}
```

<a id="modules-packages-and-imports"></a>

## 模块、包和导入

### `goml.toml`

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

### 包注册表

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

### 示例

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

### 原生依赖

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

### 构建输出布局

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

<a id="directory-packages"></a>

### 目录包

每个包含 `.goml` 文件的目录都是一个包。同一目录的所有源文件必须声明相同包名：

```goml
package utils;
```

包的规范身份由模块路径和相对目录路径组成。包声明名称决定无别名导入时的本地名称，而不决定包的全局身份。根包的规范身份就是模块路径。

典型项目可以是：

```text
goml.toml
main.goml
utils/utils.goml
```

`utils/utils.goml`:

```goml
package utils;

pub fn message() -> string {
    "hello"
}
```

`main.goml`:

```goml
package main;

use alice::myapp::utils;

fn main() -> () {
    println(utils::message())
}
```

### `use`

`use` 是文件级声明，必须放在 `package` 之后、任何顶层项之前。不同文件不共享导入。

导入不传递：包 A 导入包 B，并不会让使用 A 的文件自动看到 B。每个文件必须直接导入它引用的每个包。

```goml
use alice::http::client;
use alice::http::client as http_client;
```

上下文路径首段 `module` 表示当前项目 `[module].path` 声明的规范路径。模块中的任何包都可以使用，包括规范路径包含多段的模块：

```goml
use module::http::client;
use module::rendering::api::{Canvas, Render};
pub use module::model::Request;
```

对于声明为 `alice::myapp` 的模块，这些路径解析为 `alice::myapp::http::client`、`alice::myapp::rendering::api` 和 `alice::myapp::model::Request`。此标记仅在 `use` 路径开头且后跟 `::` 时生效；其他位置名为 `module` 的依赖路径和别名保留普通含义。外部依赖继续使用 `alice::http::client` 这样的规范路径。

导入包后，通过本地别名访问其公开项：

```goml
let request = http_client::new_request();
```

花括号导入将选定的公开顶层项直接引入当前文件。可以导入函数、常量、结构体、枚举、类型别名和 trait，并可分别重命名：

```goml
use alice::rendering::api::{Canvas, Color as Paint, Render, render_to_string};

fn describe(canvas: Canvas) -> string {
    render_to_string(canvas)
}
```

花括号前的路径始终是包路径。不支持嵌套花括号导入、`self`、`*` 通配导入、枚举变体导入或子包聚合。导入名仍限于当前文件，并遵循普通顶层可见性规则。

以下特殊形式还可以将已加载包中的 trait 加入方法调用作用域：

```goml
use alice::rendering::api;
use api::Render;
```

在花括号列表中导入 `Render` 也会将其加入方法作用域。之后可对具体值写 `value.render()`。即使未将 trait 加入方法作用域，仍可写限定调用 `api::Render::render(value)`。

`pub use` 从当前包重新导出具名包或公开项，且不改变其身份：

```goml
pub use alice::http::client;
pub use client::{Client, Request, Response};
pub use client::Request as HttpRequest;
```

下游代码可以使用 `facade::HttpRequest`、直接导入它，或将重新导出的 trait 引入方法作用域。编译器将这些路径都解析回原始声明，因此构造器、trait 实现、方法、常量和函数的行为与原始路径完全一致。重新导出的包命名空间也仍可使用，例如 `facade::client::Request`。同包中的重新导出名必须唯一。仍不支持通配重新导出。

不支持以下路径模型：

```text
mod child;
crate::x
self::x
super::x
::rooted::x
```

<a id="visibility"></a>

### 可见性

顶层项默认只在包内可见。顶层函数、结构体、枚举和 trait 用 `pub` 导出；结构体字段和固有 impl 方法也默认私有，跨包访问时须单独标注 `pub`：

```goml
pub struct Point {
    pub x: i32,
    pub y: i32,
}

pub fn origin() -> Point {
    Point { x: 0, y: 0 }
}

impl Point {
    pub fn sum(self) -> i32 {
        self.x + self.y
    }
}
```

同包中的所有文件都可以使用私有顶层项。trait impl 方法继承 trait 方法的可见性，不能写 `pub`。

<a id="type"></a>

## 类型

<a id="basic-types-and-composite-types"></a>

### 基本类型和复合类型

| 类型 | 示例 | 说明 |
| --- | --- | --- |
| 空元组 | `()` | 唯一的值是 `()` |
| never | `never` | 没有值；标注 `-> never` 的函数不能正常返回 |
| 布尔值 | `bool` | `true`, `false` |
| 平台宽度整数 | `isize` | 对应目标 Go 平台的 `int`，也是整数的默认类型。 |
| 有符号整数 | `i8`, `i16`, `i32`, `i64` | 固定宽度 |
| 无符号整数 | `usize`, `u8`, `u16`, `u32`, `u64` | `usize` 对应目标 Go 平台的 `uint`；其余为固定宽度 |
| 字节 | `byte` | `u8` 的透明内置别名 |
| 浮点数 | `f32`, `f64` | IEEE 浮点数 |
| SIMD 向量和掩码 | `simd::u8x16`, `simd::i32x8`, `simd::f32x4`, `simd::f64x4` 及匹配的掩码 | 从 `std::simd` 导入的固定 128/256 位整数和浮点向量 |
| 字符串 | `string` | 使用 Go string 后端 |
| 字符 | `char` | 编译为 Go `rune` |
| 元组 | `(i32, string)` | 类型语法中的非空元组至少有两个元素 |
| 定长数组 | `[i32; 4]` | 长度是类型的一部分 |
| 函数 | `(i32, string) -> bool` | 从参数类型列表到返回类型 |
| 泛型应用 | `Option[i32]`, `pkg::Box[string]` | 使用方括号 |
| 通道 | `Channel[isize]`, `Sender[isize]`, `Receiver[isize]` | 双向和单向 Go 通道后端 |
| trait 对象 | `dyn Render`, `dyn Consumer[string] + Close` | 一个或多个 dyn 安全的 trait 约束；关联类型必须绑定 |
| 不透明返回类型 | `impl Iterator[Item = isize]` | 函数唯一的隐藏具体返回类型，通过 trait 约束暴露能力 |
| 关联类型投影 | `I::Item`, `Self::Output`, `I::IntoIter::Item` | 必须有相应 trait 约束；投影可链式连接 |

函数类型示例：

```goml
let predicate: (i32) -> bool = |value: i32| value > 0;
let combine: (i32, i32) -> i32 = |a, b| a + b;
let action: () -> () = || println("run");
let accepts_empty: (()) -> () = |value: ()| value;
```

`()` 是空元组类型。在 `->` 左侧，`() -> T` 表示零参数函数类型，`(()) -> T` 表示接收一个 `()` 参数的函数。

`never` 类型的表达式不产生值，可出现在期望其他值类型的位置。但不会隐式转换容器或函数类型：`Vec[never]` 不是 `Vec[isize]`，`() -> never` 也不是 `() -> isize`。要将不返回的函数适配为其他返回类型的回调，请使用具有该预期回调类型的闭包。普通值和正常返回的回调不能满足 `never` 标注。

值和模式中的 `(value)` 是分组，`(value,)` 是单元素元组。当前类型语法无法直接标注单元素元组：`(T,)` 仍会规范化为 `T`，因此此类值的类型必须从局部上下文推断。

`A -> B -> C` 按右结合解析。建议始终给函数参数列表加括号，尤其是高阶函数：`(A) -> (B) -> C`。

数组长度是非负整数常量表达式。它接受常见进制的整数字面量、具名整数常量、常量泛型参数、括号、一元 `+`、`-`、`~`，以及 `+`、`-`、`*`、`/`、`%`、`<<`、`>>`、`&`、`|`、`^`。整数溢出、除零、无效移位、负长度和循环常量引用均产生诊断。长度须在目标的有符号索引范围内。这些表达式不执行函数调用：

```goml
const WIDTH: usize = 2;
let pair: [string; WIDTH] = ["left", "right"];
let bytes: [byte; WIDTH * 2] = [1, 2, 3, 4];
```

数组字面量的元素数量必须匹配数组类型。空数组和空泛型容器通常需要类型标注。

`[value; N]` 对 `value` 求值一次，再将结果复制到每个元素，构造定长数组。长度遵循与数组类型相同的整数常量表达式规则，也可以是常量泛型参数：

```goml
let zeros: [byte; 16] = [0; 16];
let labels = ["ready"; WIDTH * 2];

fn filled[T, const N: usize](value: T) -> [T; N] {
    [value; N]
}
```

即使 `N` 为零，也会对值求值。重复使用普通值复制，因此重复 `Ref`、`Vec` 或其他共享句柄会保留共享存储，不会克隆存储或为每个元素调用一次初始化器。值受支持时，数组重复也可用于 `comptime` 和常量初始化器。

<a id="type-aliases"></a>

### 类型别名

顶层别名是透明的，可以带类型参数：

```goml
type UserId = u64;

type Pair[T] = (T, T);

pub type Names = Vec[string];
```

别名不会创建名义上不同的类型，递归别名循环会被拒绝。公开别名可以跨包引用。

需要独立的名义类型而非别名时，使用单字段元组结构体：

```goml
struct UserId(u64);

struct Box[T](T);
```

这种形式是 newtype 模式。用 `UserId(value)` 构造，用 `value.0` 访问字段，用 `let UserId(inner) = value;` 解构。字段默认私有。可跨包构造和解包的公开 newtype 必须将类型和字段都标为公开：`pub struct UserId(pub u64);`。

<a id="type-syntax-not-currently-available"></a>

### 当前不支持的类型语法

GoML 没有 Rust 引用或生命周期语法、指针算术、切片字面量或联合类型。共享可变存储使用 `Ref[T]`，可选值使用 `Option[T]`，只读或可变连续视图使用 `Slice[T]` 或 `MutSlice[T]`。`std::ffi::Ptr[T]` 承载可空 Go 指针，与 `Ref[T]` 及数值类型不同。`ffi::null()` 和 `ffi::is_nil()` 操作此显式指针类型。用 `extern type` 声明的类型即使在 Go 中是指针别名，也仍然不透明，不会自动获得这些指针操作。

在 Linux amd64 上，`std::os::linux::syscall` 接受数值机器字、限定作用域的字节缓冲区参数，以及字节缓冲区之间的显式指针字段，用于底层内核调用。它不向语言添加指针转换、指针算术或原生结构体布局。

`A + B` 在泛型约束、父 trait 列表、`dyn A + B` 对象和 `impl A + B` 返回类型中组合 trait 约束。它不是具体值的通用交叉类型。

`Self` 仅用于 trait 签名和 impl 的类型位置；普通顶层函数不能将 `Self` 用作隐式类型参数。

<a id="top-level-definitions"></a>

## 顶层定义

普通源文件的顶层应只包含：

- `fn`
- `struct`
- `enum`
- `trait`
- `impl`
- `type`
- `const`
- `static`

`package` 和 `use` 只能出现在文件开头。顶层不能写局部变量或任意执行语句。

<a id="constants"></a>

### 常量

顶层常量需要显式类型和编译期表达式：

```goml
const BASE: i32 = 0x20;

pub const ANSWER: i32 = BASE + 10;

const NEWLINE: byte = b'\n';
```

常量名可用大写字母、小写字母或 `_` 开头；推荐 `UPPER_SNAKE_CASE`。常量表达式支持字面量、构造器、其他常量引用、一元和二元运算符、整数转换方法及支持编译期求值的调用。允许前向引用，但拒绝循环。允许的常量类型为递归不可变值：`bool`、数值类型、`string`、`char`、`byte`、元组、定长数组，以及字段满足同样规则的结构体或枚举。`Vec`、`HashMap`、`Ref`、`Channel`、闭包和动态值会被拒绝。复合常量不能用作模式。公开常量可通过包限定路径访问。

常量初始化器中的名称在常量声明位置解析，使用该文件的导入及外层声明的泛型参数。从其他文件或泛型函数引用常量，不会在调用方作用域中重新解释初始化器。尤其是，顶层常量不能捕获使用它的函数的泛型参数：

```goml
const SIZE: usize = 2;
const WIDTH: usize = SIZE;

fn first[const SIZE: usize](values: [isize; WIDTH]) -> isize {
    values[0]
}
```

`first::[3]([10, 20])` 接受双元素数组，因为 `WIDTH` 仍为 `2`。

可见的顶层常量可在 `match`、`if let`、`while let` 和 `let else` 中用作值模式。同时支持本地名和包限定名，常量模式可嵌套或与或模式组合：

```goml
const ANSWER: isize = 42;

match value {
    ANSWER | config::FALLBACK => "known",
    _ => "other",
}
```

在这些可反驳模式上下文中，可见常量优先于同名绑定的引入。要绑定该名称，使用 `answer @ _` 这样的显式别名。普通 `let` 绑定和 `for` 模式仍引入绑定，即使存在同名顶层常量。范围模式端点仍必须是整数或字符字面量。

<a id="statics-and-one-time-initialization"></a>

### 静态变量与一次性初始化

顶层 `static` 具有整个程序唯一的存储位置。其绑定不能重新赋值。显式类型必须是 `OnceCell[T]`，初始化器必须恰为 `OnceCell::new()`：

```goml
static EMOJI: OnceCell[FrozenVec[FrozenVec[u16]]] = OnceCell::new();

fn emoji() -> FrozenVec[FrozenVec[u16]] {
    EMOJI.get_or_init(|| build_emoji().freeze())
}
```

`OnceCell.get_or_init(init)` 运行一次初始化器并缓存结果。并发的首次调用者等待该初始化器，并获得同一个值。如果初始化器 panic，单元恢复为未初始化状态，等待者被唤醒，以便后续初始化器重试。递归初始化同一单元会 panic，错误中包含静态变量名称。缓存的 `Result` 只是普通缓存值。进程退出时，静态变量不执行用户可观察的析构。

与 `const` 不同，`static` 有可观察的身份。`OnceCell[T]` 控制初始化，但不会让 `T` 不可变。因此共享缓存应暴露 `FrozenVec` 或其他不可变值，而不是可变 `Vec`。

<a id="compile-time-evaluation"></a>

### 编译期求值

`comptime { expression }` 要求在编译当前包时对表达式求值。其静态类型为内部表达式的类型，编译器在 Core 降级前将它替换为等价的普通值：

```goml
#[comptime]
fn factorial(value: isize) -> isize {
    if value < 2 {
        1
    } else {
        value * factorial(value - 1)
    }
}

const SIX: isize = factorial(3);

fn table() -> [isize; 4] {
    comptime {
        [factorial(1), factorial(2), factorial(3), factorial(4)]
    }
}
```

`#[comptime]` 将非泛型自由函数或具体实现方法标记为可在编译期求值。函数仍可在运行时调用。编译期调用可以调用其他 `#[comptime]` 自由函数、下文列出的内置操作或 `compile_error(string) -> never` 内建函数。编译器验证每个标记函数的完整函数体，包括某次调用未经过的分支。带参数的属性、重复属性、泛型函数或实现、extern 函数和其他声明会被拒绝。

顶层常量初始化器是隐式编译期上下文，因此 `const SIX: isize = factorial(3);` 与将初始化器包在 `comptime { ... }` 中等价。除了标量，顶层常量还可产生递归不可变的元组、定长数组、结构体或枚举。普通代码中的 `comptime` 表达式支持相同的可具体化值形态。

编译期代码可使用局部绑定和赋值、块、`if`、`match`、`while`、`loop`、受限的 `for`、`break`、`continue`、`return`、递归、直接调用、整数转换方法、整数 `to_string()` 和受支持运算符。编译期 `for` 仅接受定长数组或内置 `isize` 范围 `start..end` 和 `start..=end`；源值和范围端点求值一次，模式必须不可反驳。还支持确定性的字符串方法 `len`、`byte_len`、`get`、`byte_get`、`byte_slice`、`is_char_boundary`、`starts_with`、`ends_with` 和 `contains`。字符串索引和切片使用字节偏移，拒绝无效的 UTF-8 字符边界。整数格式化接受有符号和无符号的 8、16、32、64 位值，包括 `byte`、`isize` 和 `usize`，生成完整保留取值范围的十进制文本。此能力仅限内置 `ToString` 实现；同名用户 trait 不会成为编译期内建操作。例如，`const BITS: u64 = 18446744073709551615; const MASK: string = BITS.to_string();` 无需运行时格式化即可求值。

编译期代码不能捕获外部运行时参数或局部变量。一次求值内支持闭包、函数参数和间接调用，包括共享可变捕获。具体的 `#[comptime]` 方法支持静态调用和通过本地实现进行动态分派。`Ref` 支持 `new`、`get`、`set`；`Vec` 支持 `new`、`with_capacity`、`from_array`、`get`、`set`、`push`、`len`、`is_empty`、`clear` 和 `truncate`。这些操作保留普通别名关系和边界检查。函数值、动态对象、引用和向量不能作为运行时常量逃逸。不支持泛型编译期函数和实现、通用迭代器、浮点计算、`HashMap`、通道、goroutine、extern 调用、宿主 I/O、环境访问、时间、随机性、网络访问、通用类型反射、任意声明生成和通用类型级计算。下文的受限可编程 derive 接口是唯一的反射和代码生成机制。

`compile_error` 只能用于 `#[comptime]` 函数、`comptime` 块或顶层常量初始化器。它以给定消息终止编译期求值。如果运行时执行 `#[comptime]` 函数到达它，程序会陷入失败：

```goml
#[comptime]
fn checked_size(value: isize) -> isize {
    if value < 0 {
        compile_error("size must be non-negative")
    } else {
        value
    }
}
```

失败信息包含编译期调用栈和源码位置。求值使用确定性的指令数、调用深度、临时值节点、临时内存和最终值大小限制。超出限制是编译错误；语言语义不使用墙钟超时。一次求值内，成功的直接调用会被记忆化缓存。由于编译期代码没有用户可见副作用，记忆化不可观察。

公开的 `#[comptime]` 函数可以跨包调用。定义包的接口包含公开入口及其可达的私有编译期辅助函数和常量的已验证编译期 IR。公开编译期常量值也会导出。编译期函数体或值的变化影响接口哈希；源码格式、局部名称和源码位置不影响。下游包只需依赖接口即可检查和求值。目前无法具体化包含其他包隐藏字段的值。

编译期整数求值使用与生成的运行时代码相同的定宽回绕表示。除零、负移位次数和越界索引会导致编译失败。有符号最小值除以 `-1` 得到该最小值，余数为零。窄化转换保留目标宽度的低位；有符号源值在扩宽时先做符号扩展，再转换为目标符号类型。CTIR 目标规格属于其语义哈希的一部分。`isize` 和 `usize` 当前使用 64 位 Linux amd64 工具链目标宽度。

运行时整数加、减、乘使用相同的定宽回绕表示。整数除零或取余除零会因生成的 Go 运行时错误而终止当前进程。整数转换方法使用与编译期求值相同的窄化、符号扩展和符号性规则。

<a id="structs"></a>

### 结构体

```goml
struct Point {
    x: i32,
    y: i32,
}

struct Box[T] {
    value: T,
}
```

newtype 是恰有一个未命名字段的名义结构体：

```goml
struct UserId(u64);

fn increment(value: UserId) -> UserId {
    UserId(value.0 + 1)
}
```

与 `type UserId = u64;` 不同，newtype 不能与 `u64` 互换，必须显式构造或解构。newtype 可为泛型，字段默认私有，括号内允许尾逗号。此语法要求恰好一个字段；不支持零字段或多字段元组结构体。

结构体和枚举的类型参数列表不接受约束。约束应放在使用该类型的函数、trait 或 impl 上。

构造时写出字段，允许字段简写：

```goml
fn make_point(x: i32, y: i32) -> Point {
    Point { x, y }
}
```

具名字段字面量不接受 turbofish 实参。字段值不能提供预期泛型类型时，请添加类型标注，例如 `let value: Box[bool] = Box { value: true };`。结构体更新也遵循此规则。

字段访问使用点号：

```goml
let x = point.x;
```

结构体更新从同类型基值复制省略的字段。`..base` 必须放在最后。先从左到右求值显式字段表达式，再求值基值表达式，每个表达式只求值一次：

```goml
let moved = Point { x: point.x + 1, ..point };
```

跨包构造时，如果不可访问的字段阻止构造，结构体更新会被拒绝。

直接递归的结构体会有无限大小，必须通过 `Ref`、`Vec` 等间接层递归：

```goml
struct Node {
    value: i32,
    next: Option[Ref[Node]],
}
```

<a id="enums"></a>

### 枚举

```goml
enum Message[T] {
    Quit,
    Value(T),
    Pair(T, T),
    Named { value: T },
}
```

无载荷变体是值，元组变体使用函数调用语法，结构体式变体使用具名字段：

```goml
let quit: Message[i32] = Message::Quit;
let value: Message[i32] = Message::Value(42);
let pair = Message::Pair("left", "right");
let named = Message::Named { value: 42 };
```

无载荷变体没有可供推断泛型实参的载荷，因此通常需要预期类型：

```goml
let none: Option[i32] = Option::None;
```

元组变体构造器也可以作为一等函数值使用：

```goml
let some: (i32) -> Option[i32] = Option::Some;
```

显式泛型实参属于枚举本身，放在变体名之前。对于直接构造和构造器函数值，它们都会约束载荷和结果类型：

```goml
let value = Message::[i32]::Value(42);
let quit = Message::[string]::Quit;
let make = Message::[bool]::Value;
let ready = make(true);
```

带常量参数的枚举在同一列表中包含常量实参，例如 `Storage::[i32, 2]::Values([1, 2])`。提供的实参数量和种类必须与枚举参数匹配。变体没有独立泛型参数列表，显式实参必须与预期类型及载荷类型一致。

结构体式变体使用预期类型而非 turbofish 实参：写作 `let value: Message[bool] = Message::Named { value: true };`。`Message::[bool]::Named { value: true }` 会被拒绝。

预期枚举类型也允许在表达式中使用未限定构造器：

```goml
let none: Option[i32] = None;
let nested: Option[Result[i32, string]] = Some(Ok(42));
let message: Message[string] = Named { value: "ready" };
let some: (i32) -> Option[i32] = Some;

fn success(value: i32) -> Result[i32, string] {
    Ok(value)
}
```

类型标注、函数参数、返回类型和外层构造器载荷都可以提供此上下文。它适用于无载荷、元组式和结构体式变体，包括导入枚举的公开变体及经类型别名访问的枚举。预期构造器函数类型通过返回类型选择枚举。普通名称解析优先；已有局部值或结构体不会被枚举构造器替代。当前包中唯一的变体即使没有预期类型信息也仍可使用。当普通名称解析和预期类型都无法选择变体时，使用 `Enum::Variant` 限定。

结构体式变体使用字段构造和字段模式：

```goml
match named {
    Message::Named { value } => value,
    _ => 0,
}
```

模式可省略枚举限定。预期模式类型为 `Option[T]` 时，`Some(value)` 和 `None` 解析为 `Option::Some(value)` 和 `Option::None`。相同规则适用于用户枚举的无载荷、元组式和结构体式变体，包括从其他包导入及经类型别名访问的枚举。无关枚举中的同名变体不会产生歧义，因为预期枚举类型会选择所属枚举。嵌套模式递归解析，因此 `Some(Ok(value))` 用 `Some` 的载荷类型解析 `Ok`。

<a id="functions-and-generics"></a>

## 函数与泛型

<a id="top-level-functions"></a>

### 顶层函数

```goml
fn add(left: i32, right: i32) -> i32 {
    left + right
}

fn log(message: string) {
    println(message)
}
```

参数类型不能省略。省略 `-> ...` 等价于 `-> ()`。同包顶层函数名必须唯一，不支持按参数类型重载。

<a id="opaque-return-types"></a>

### 不透明返回类型

函数可以通过 `-> impl Trait` 暴露 trait 约束，同时隐藏具体返回类型。约束可包含泛型实参、关联类型绑定和 `+`：

```goml
trait Source {
    type Item;
    fn get(self) -> Self::Item;
}

struct Stored[T] {
    value: T,
}

impl[T] Source for Stored[T] {
    type Item = T;

    fn get(self) -> T {
        self.value
    }
}

fn source[T](value: T) -> impl Source[Item = T] {
    Stored { value }
}

let number = source(42);
let value: isize = number.get();
```

对于给定泛型实例，同一函数的所有返回路径必须产生相同具体类型，且满足全部声明约束。不同泛型实例可以使用不同具体类型。公开函数可以隐藏私有类型，包括跨包隐藏。调用者可使用声明的 trait 约束和关联类型绑定，但不能访问隐藏类型的字段或额外固有方法。即使两个函数返回同一具体类型，每个函数仍定义自己的不透明类型。

分派仍是静态的，并参与单态化；`impl Trait` 不创建 trait 对象，也不要求约束 dyn 安全。需要 trait 对象时，请显式转换为 `as dyn Trait`。`impl Trait` 只支持作为顶层函数或固有方法的完整返回类型。不能出现在参数、类型别名、局部类型标注、`Option[impl Trait]` 这样的嵌套返回类型或 trait 方法声明中。

不透明约束当前接受位置类型实参和关联类型绑定，不接受 `impl Fixed[4]` 这样的整数常量实参。

<a id="generic-functions"></a>

### 泛型函数

```goml
fn identity[T](value: T) -> T {
    value
}

fn render[T: ToString + Eq](value: T) -> string {
    value.to_string()
}
```

约束也可以放在 `where` 子句中：

```goml
use std::iter;

fn collect_items[T, I: Iterator](iterator: I) -> Vec[T] where I::Item = T {
    iter::collect(iterator)
}
```

当前支持两种 `where` 谓词：

```text
Type: TraitA + TraitB
TypeA = TypeB
```

普通泛型函数调用通常从实参和预期结果类型推断类型实参：

```goml
let number: i32 = identity(1);
let text: string = identity("text");
```

需要显式指定时，使用 GoML turbofish：

```goml
let number = identity::[i32](1);
let text = identity::[string]("text");
```

显式泛型实参列表至少要有一个实参。使用推断时省略整个列表；`identity::[](1)` 会被拒绝。

不要写 `identity[i32](1)` 或 Rust 的 `identity::<i32>(1)`。所属类型或 trait 的泛型实参放在成员名之前，方法自己的实参放在成员名之后：

```goml
let inferred = value.convert(fallback);
let explicit = value.convert::[string](fallback);
let inherent = Box::[i32]::convert::[string](value, fallback);
let trait_call = Convert::[i32]::convert::[string](value, fallback);
```

顶层函数和方法可以引入自己的类型参数。方法参数不同于外层 trait 或 impl 的参数，可有自己的约束和 `where` 谓词。没有局部具名函数；请使用闭包。闭包没有泛型。结构体、枚举、trait 和 impl 块也可有类型参数。

作用域内的泛型类型参数优先于未限定的包类型或导入名。例如，包中同时定义了 `struct T` 时，`fn identity[T](value: T) -> T { value }` 的参数和返回类型引用泛型参数。限定路径仍引用包项。

自由函数可以仅在函数体中使用类型参数，而不在参数或结果类型中出现。此类实参需要显式提供，例如 `schema::[User]()`，取得函数值时也一样。这些实参在跨包接口和特化中保留；未使用的泛型定义不会生成未特化的可执行函数体。

GoML 对泛型调用进行单态化。递归泛型代码必须只产生有限个具体实例，不能每次递归都不断变为新的嵌套类型。

<a id="blocks-bindings-and-assignments"></a>

## 块、绑定和赋值

<a id="const-generic-parameters"></a>

### 常量泛型参数

函数、结构体、枚举、别名、trait 和实现都可以在类型参数旁声明 `const N: usize`。常量参数在表达式中是值，在泛型应用中是整数实参；不能用作元素类型。每个具体实参都参与单态化和名义类型身份。

```goml
struct Buffer[T, const N: usize] {
    values: [T; N],
}

fn first[T, const N: usize](values: [T; N]) -> T {
    values[0]
}

fn count[const N: usize](values: [i32; N]) -> usize {
    N
}

fn use_buffers() -> () {
    let buffer: Buffer[i32, 3] = Buffer { values: [4, 5, 6] };
    let first_value = first(buffer.values);
    let capacity = count::[3](buffer.values);
}
```

直接的长度参数可从数组实参或预期类型推断。对于 `N + 1` 这样的长度表达式，请显式提供 `N`；推断不求解算术方程。泛型实参使用字面量或名称，复合整数实参使用花括号，例如 `Buffer[i32, {2 + 1}]`。常量参数目前类型为 `usize`，没有默认值，取值范围为 `0..=9223372036854775807`。数组重复接受这些参数，例如 `[value; N]`。不支持数组长度中的类型相关关联常量。

<a id="block-values-and-semicolons"></a>

### 块值与分号

块中最后一个无分号表达式是返回值：

```goml
fn square(value: i32) -> i32 {
    let result = value * value;
    result
}
```

没有尾表达式的块返回 `()`。表达式后加 `;` 会使其成为表达式语句并丢弃结果：

```goml
fn run() -> () {
    println("first");
    1 + 2;
}
```

`let`、普通赋值和一般非尾表达式语句需要分号。作为语句且后面还有代码时，`if`、`match`、`select` 及无标签的 `while`、`loop`、`for` 可以省略分号；其他表达式语句仍需要分号。函数、花括号结构体、枚举、trait、impl 和块本身的声明后不加分号。`struct UserId(u64);` 这样的 newtype 声明则需要尾分号。

`defer expression;` 注册一个 `()` 表达式，在离开当前词法块时运行。正常完成以及 `return`、`?`、`break` 或 `continue` 穿过该块时，延迟表达式按后进先出顺序运行。返回值或 break 值先求值，再开始清理。每个循环体块都有自己的清理栈，因此某次迭代注册的延迟表达式会在该迭代退出前运行。

与 Go 的 `defer` 不同，GoML 不在到达语句时求值调用实参。整个表达式在块退出时求值，因此通过 `Ref` 和捕获的可变局部变量读取的是清理时的值。闭包体是独立控制流作用域。延迟表达式不能包含 `return`、`break`、`continue` 或 `?`。panic 展开也按后进先出顺序运行已注册清理。清理在调用前移除，因此即使 panic 也恰好运行一次；其余清理仍会运行，最新的 panic 继续传播。编译器使用函数局部清理栈，在词法退出处显式弹出，并用 Go `defer` 守卫处理展开。循环迭代释放已完成条目，不会累积延迟调用到函数返回。突然终止进程和致命运行时故障不保证清理。

```goml
fn work() -> () {
    let state = Ref::new("open");
    defer println("closing:" + state.get());
    defer println("flush");
    state.set("ready")
}
```

此例先输出 `flush`，再输出 `closing:ready`。

<a id="let-type-annotations-and-shadowing"></a>

### `let`、类型标注与遮蔽

```goml
let inferred = 42;
let explicit: i64 = 42;
let _ = println("discard explicitly");
let name = "first";
let name = "second";
```

局部变量允许同名遮蔽。`let _ = expr;` 显式执行表达式并丢弃结果。

每个块都有词法绑定作用域，包括普通表达式块、`comptime` 块和任务 `scope` 块。绑定在初始化器之后直到块末尾可见，包括嵌套块和闭包。离开块会恢复同名外层绑定；新引入的名称在块外不可用。对已有外层可变绑定赋值仍会修改该绑定：

```goml
let value = 1;
let inner = {
    let value = 2;
    value
};
println(inner);
println(value);
```

此例先输出 `2`，再输出 `1`。

`let` 左侧可以是不可反驳模式：

```goml
let (left, right) = pair;
let Point { x, y: vertical } = point;
```

普通 `let` 不能使用可能失败的枚举或字面量模式。失败路径必须离开周围控制流时，使用 `let ... else`：

```goml
fn require_value(value: Option[isize]) -> isize {
    let Some(item) = value else {
        return -1;
    };
    item
}
```

初始化器求值一次。模式必须可反驳，绑定在语句之后可见，`else` 块不能使用这些绑定。`else` 块必须通过 `return`、`break`、`continue`、无限 `loop` 或其他 `never` 类型表达式发散。两条路径都要在本地继续时，使用 `if let` 或 `match`。

<a id="mut-and-assignment"></a>

### `mut` 与赋值

```goml
let mut count = 0;
count = count + 1;
```

解构模式可以只将选定的绑定标为可变：

```goml
let (mut index, value) = pair;
match state {
    Some(mut count) => {
        count += 1;
    },
    None => (),
}
for mut item in values {
    item += 1;
}
```

模式中的 `mut` 只能用于绑定，不引入借用或引用语义。现有的 `let mut pattern = value;` 形式仍受支持，并递归地将模式内所有绑定设为可变。逐绑定的 `mut` 可与元组、结构体、枚举、数组、`match` 和 `for` 模式组合。

没有 `mut` 的普通局部绑定不能重新赋值。赋值目标包括可变局部变量、元组投影、结构体字段和受支持索引位置。复合赋值支持 `+=`、`-=`、`*=`、`/=`、`%=`、`&=`、`|=`、`^=`、`<<=` 和 `>>=`：

```goml
count += 1;
point.x *= 2;
values[index] <<= 1;
```

目标根和每个索引都在右侧表达式之前求值一次。不支持 `++` 和 `--`。不支持通过 `HashMap` 索引复合赋值，因为索引读取返回 `Option[V]`；请使用 `get` 和 `set`，或普通索引赋值。

数组索引赋值要求根值为 `let mut` 局部数组，或来自内置 `Ref.get()`：

```goml
let mut values = [1, 2, 3];
values[1] = 20;
let shared = Ref::new([4, 5]);
shared.get()[0] = 40;
```

`Vec` 和 `HashMap` 是内部可变容器，因此即使绑定本身没有 `mut`，也可以通过方法或索引修改内容：

```goml
let values: Vec[i32] = Vec::new();
values.push(10);
values[0] = 20;
let counts: HashMap[string, i32] = HashMap::new();
counts["answer"] = 42;
```

`Slice[T]` 是只读视图，不能按索引赋值。`MutSlice[T]` 是有界可变视图，支持索引赋值且不会复制其 `Vec[T]` 底层存储。

<a id="expressions-and-operators"></a>

## 表达式和运算符

<a id="basic-expressions"></a>

### 基本表达式

GoML 支持：

- 字面量、变量和路径
- 元组、数组和结构体字面量
- 块表达式 `{ ... }`
- 函数调用 `f(a, b)`
- 字段访问 `value.field`
- 元组投影 `pair.0`、`triple.2`
- 方法调用 `value.method(arg)`
- 索引 `value[index]`
- 一元和二元运算
- 整数转换方法 `value.to_u32()`
- 显式 trait 对象转换 `value as dyn Trait`
- 半开及闭区间范围表达式 `start..end` 和 `start..=end`
- `if`, `if let`, `match`, `select`, `while`, `while let`, `loop`, `for`
- 闭包
- `return`、`break`、`continue`、`go` 和 `?`

块可出现在任何普通表达式位置。最后一个无分号表达式提供值；没有尾表达式的块类型为 `()`：

```goml
let answer = {
    let base = 40;
    base + 2
};
return {
    cleanup();
    answer
};
```

在 `if`、`if let`、`while`、`while let`、`for` 或 `match` 头部中，顶层 `{` 开始控制流主体。将块表达式用作头部值时，请加括号：

```goml
if ({
    prepare();
    ready()
}) {
    run()
}
```

<a id="operator-precedence"></a>

### 运算符优先级

从低到高：

| 优先级 | 运算符 | 说明 |
| --- | --- | --- |
| 1 | `..`, `..=` | 半开或闭区间范围；不可链式连接 |
| 2 | `\|\|` | 短路逻辑或 |
| 3 | `&&` | 短路逻辑与 |
| 4 | `==`, `!=`, `<`, `>`, `<=`, `>=` | 比较；不可链式连接 |
| 5 | `\|` | 按位或 |
| 6 | `^` | 按位异或 |
| 7 | `&` | 按位与 |
| 8 | `<<`, `>>` | 移位 |
| 9 | `+`, `-` | 加法、减法、字符串拼接 |
| 10 | `*`, `/`, `%` | 乘法、除法、取余 |
| 11 | `as` | 显式 `dyn Trait` 转换 |
| 12 | 一元 `-`, `!`, `~` | 前缀 |
| 13 | 调用 `()`、索引 `[]` | 后缀 |
| 14 | `?` | 错误或缺失值传播 |
| 15 | 成员 `.` | 字段访问、元组投影或方法选择 |

二元运算符左结合，函数类型 `->` 右结合。不要链式比较；请组合逻辑运算：

```goml
let inside = lower <= value && value < upper;
```

调用、索引、`?` 和成员访问比一元运算符结合更紧。例如，`-compute()` 对调用结果取负，`!values[0]` 对索引值取反：

```goml
let negative = -compute();
let valid = !values[0];
```

<a id="operator-type-rules"></a>

### 运算符类型规则

- `+ - * /` 的两个数值操作数必须是相同的具体数值类型；没有隐式数值提升。
- `%` 仅接受相同类型的整数。
- `& | ^ ~` 仅接受整数，二元按位运算两侧必须同类型。`<< >>` 两侧必须是整数，结果类型与左操作数相同。
- `+` 还支持 `string + string`。
- `&& || !` 仅接受 `bool`。
- 一元 `-` 仅接受有符号整数或浮点数。
- 比较两侧必须同类型。
- `< > <= >=` 使用 `PartialOrd`。基本数值、`string`、`char`、元组、定长数组、`Vec`、`Slice`、`Option` 和 `Result` 提供相应的条件实现。
- `== !=` 使用 `PartialEq`。元组、定长数组、`Vec`、`Slice`、`Option` 和 `Result` 在元素实现 `PartialEq` 时递归比较。
- 浮点值实现 `PartialEq + PartialOrd`，但不实现 `Eq + Ord + Hash`。尤其是 NaN 不等于自身，其 `partial_cmp` 结果为 `None`。
- 每种整数类型为其他整数类型提供 `to_isize`、`to_i8`、`to_i16`、`to_i32`、`to_i64`、`to_usize`、`to_u8`、`to_u16`、`to_u32` 和 `to_u64`。省略同类型转换方法。`byte` 使用 `u8` 的方法，`char` 提供 `to_u32`，`u32` 转 `char` 使用返回 `Option[char]` 的 `char_from_u32`。
- `f32` 提供 `to_bits` 和 `to_f64`；`f64` 提供 `to_bits` 和 `to_f32`。prelude 函数 `float32_from_bits` 和 `float64_from_bits` 从 IEEE 754 位模式重建值。`f64` 转 `f32` 使用舍入到最近值、相等时取偶数；反向转换精确。
- 导入 `std::num::{ToFloat, TryToInt}`，即可使用标量整数 `.to_f32()` / `.to_f64()` 和带检查的浮点 `.try_to_i32()` / `.try_to_u64()` 转换，每种整数宽度都有类似方法。浮点转整数向零截断，拒绝非有限值或超范围结果。
- 表达式中的 `as` 专用于创建 `dyn Trait` 值。数值转换使用上述整数和浮点方法；其他非 `dyn` 目标会被拒绝。`use package as alias;` 中独立的 `as` 形式仍用于选择导入别名。

没有幂运算、空值合并或用户自定义运算符。

<a id="range-expressions"></a>

### 范围表达式

`start..end` 构造递增的半开 `FnIterator[isize]`，`start..=end` 构造包含上界的迭代器。二者均可直接用于 `for`：

```goml
for value in 0..10 {
    println(value)
}

for value in 0..=10 {
    println(value)
}
```

两端均为 `isize`。半开范围在 `start >= end` 时为空；闭区间范围在 `start > end` 时为空，两端相等时包含一个值。每个端点从左到右求值一次。闭区间迭代不计算 `end + 1`，因此最大 `isize` 端点不会溢出。范围表达式不能链式连接。不支持无界范围、字符范围或自定义步长语法。模式中的 `..` 和 `..=` 属于独立的范围模式语法。

<a id="control-flow"></a>

## 控制流

### `if`

`if` 始终是表达式。产生 `()` 以外的值时必须有 `else`：

```goml
fn absolute(value: i32) -> i32 {
    if value < 0 {
        -value
    } else {
        value
    }
}
```

两个分支必须产生兼容类型。只有副作用的条件语句可以省略 `else`；此时 then 分支必须产生 `()`：

```goml
if enabled {
    println("enabled")
}
```

`else if` 是 `else` 分支中的另一个 `if`：

```goml
let label = if score > 90 {
    "high"
} else if score > 60 {
    "middle"
} else {
    "low"
};
```

`if let` 在模式匹配成功时执行 then 分支，并将模式绑定限制在该分支内。被匹配表达式只求值一次：

```goml
let number = if let Some(value) = candidate {
    value
} else {
    0
};
```

`if let` 的 `else` 也可省略；此时编译器将 else 分支视为 `()`，所以 then 分支必须产生 `()`：

```goml
if let Some(value) = candidate {
    println(value);
}
```

需要从两个分支获得值时，必须显式写 `else`。

条件可用 `&&` 串联布尔表达式和 `let` 匹配：

```goml
fn combine(left: Option[isize], right: Option[isize]) -> Option[isize] {
    if let Some(a) = left && a > 0 && let Some(b) = right && b > a {
        Some(a + b)
    } else {
        None
    }
}
```

各项从左到右执行，在首个假条件或失败模式处停止。每个实际到达的初始化器求值一次。`let` 绑定对后续项和成功分支可见，对先前项、`else` 分支及 `if` 后的代码不可见。绑定可遮蔽外层名；`else` 中仍在外层作用域解析该名。非模式项必须为 `bool`。let 链不能包含顶层 `||`；单项内的布尔析取须加括号，例如 `if let Some(x) = value && (x == 1 || x == 2) { ... }`。

### `match`

```goml
fn unwrap_or(value: Option[i32], fallback: i32) -> i32 {
    match value {
        Some(inner) => inner,
        None => fallback,
    }
}
```

被匹配表达式只求值一次，模式从上到下尝试。所有分支必须产生兼容类型，编译器要求覆盖所有可能值。

模式中的未限定变体仅根据被匹配值的类型解析。编译器不会搜索无关枚举作为后备。若该类型不是枚举、不包含该变体或无法推断，模式会被拒绝。仍可使用完全限定模式：

```goml
match value {
    Some(inner) => inner,
    None => fallback,
}
```

此规则适用于 `match`、`if let`、`while let`、`let`、`for`、嵌套模式和或模式。`let` 和 `for` 仍遵循现有的可反驳性要求。

match 分支可以带 `bool` 类型的守卫。守卫可使用该分支模式引入的绑定：

```goml
match value {
    Some(inner) if inner > 0 => inner,
    Some(_) => 0,
    None => -1,
}
```

守卫仅在模式匹配成功后求值。带普通守卫的分支不计入穷尽覆盖；只有无守卫或守卫字面量为 `true` 的分支提供覆盖。守卫字面量为 `false` 的分支不可达。编译器还会警告完全被先前模式覆盖的分支。

分支体可以是单表达式或块。不同分支必须用逗号分隔，最后一个分支的逗号可省略：

```goml
match value {
    Some(inner) => {
        let doubled = inner * 2;
        doubled
    },
    None => 0,
}
```

### `select`

`select` 等待一个可以进行的通道操作，并将该分支的求值结果作为表达式值：

```goml
let outcome = select {
    recv(messages) as message => match message {
        Some(value) => "received " + value,
        None => "messages closed",
    },
    send(events, "ready") => "sent",
    default => "not ready",
};
```

接收分支形式为 `recv(channel) as binding => body`。通道必须是 `Channel[T]` 或 `Receiver[T]`，绑定类型为 `Option[T]`；已关闭且排空的通道立即就绪，并提供 `None`。不需要接收值和关闭状态时写 `_`。发送分支形式为 `send(channel, value) => body`，接受 `Channel[T]` 或 `Sender[T]`，要求值匹配通道元素类型。

通信分支的操作数后可带 `when` 守卫。操作数和守卫按源码顺序恰好求值一次。假守卫通过选择 nil 单向通道禁用该分支；若全部分支被禁用，无 `default` 的 select 阻塞，有 `default` 的立即运行。接收绑定在其守卫中不在作用域内。

```goml
select {
    recv(messages) when accepting match {
        Some(Data { id, value }) => handle(id, value),
        Some(Stop) | None => stop(),
    },
    send(events, "ready") when publishing => published(),
    default => idle(),
}
```

`recv(channel) match { ... }` 在通信成功后穷尽匹配 `Option[T]` 结果。它等价于接收到新绑定后应用普通 `match`，因此上下文变体、嵌套模式、或模式和穷尽检查正常工作。可反驳模式不能直接替换接收绑定，否则模式失败后丢弃已接收值会造成不可见的数据丢失。

`select priority` 是确定性的非阻塞形式。它要求最后有一个 `default`，从首到尾探测通信分支，选择第一个能立即完成的分支，否则运行 `default`。假守卫被跳过。普通 `select` 保留 Go 对同时就绪操作的不确定选择。

每个 `select` 至少包含一个 `recv` 或 `send` 分支。可有一个 `default` 分支，且必须在最后。无 `default` 时阻塞到某操作可进行；有 `default` 时，没有通信分支就绪便立即执行它。多个操作就绪时，选择哪个分支未作规定。

选择前，通道操作数、发送值和守卫从上到下恰好求值一次。仅在操作被选中后才求值分支体；接收绑定仅在自身分支体可见，所有分支体必须产生兼容类型。分支体可为单表达式或块。分支用逗号分隔，允许尾逗号。`select`、`priority`、`recv`、`send`、`when`、`match` 和 `default` 在这些位置是上下文写法，不是全局保留标识符。

`select` 是运行时操作，不能用于 `comptime` 求值。取消、任务完成和计时器通过 `Receiver[()]` 值参与，而不使用专门的 select 分支语法。

### `while`

```goml
let mut index = 0;
while index < limit {
    index = index + 1;
}
```

条件必须是 `bool`，循环结果为 `()`。`while` 仅接受不带值的 `break`。

`while let` 每轮对右侧表达式求值一次；匹配成功则进入循环体并提供模式绑定，失败则退出循环：

```goml
while let Some(value) = iterator.next() {
    println(value);
}
```

循环体必须返回 `()`。单个 `while let` 绑定仅在循环体内可见。`while` 也支持与 `if` 相同的短路 let 链：

```goml
while remaining > 0 && let Some(value) = iterator.next() && value > 0 {
    println(value);
    remaining -= 1;
}
```

每次迭代从第一项重新求值链。假布尔值或失败模式退出循环。模式绑定对后续项和本轮循环体可见，下次迭代重新创建。

`break` 和 `continue` 不能直接退出 `while` 条件，包括 `while let` 初始化器或 let 链中的项。嵌套在条件内部的循环可以正常使用自己的 `break` 和 `continue`。

### `loop`

`loop` 重复执行块，直到控制流离开。与 `while` 和 `for` 不同，它的表达式结果来自 `break`：

```goml
let result = loop {
    let candidate = next();
    if candidate >= 0 {
        break candidate;
    }
};
```

指向同一 `loop` 的所有 `break` 值必须类型兼容。`break;` 提供 `()`，不能与其他 break 值类型混用。没有指向自身的 `break` 的 `loop` 类型为 `never`，不会继续到后续表达式。无标签 `break` 和 `continue` 指向最近的循环。

### `for`

```goml
let values: Vec[i32] = Vec::new();
values.push(10);
values.push(20);
for value in values {
    println(value);
}
```

`for pattern in source { ... }` 接受定长数组和实现 `IntoIterator` 的值。源表达式和 `into_iter` 转换都只执行一次。模式必须不可反驳，循环体必须返回 `()`。`start..end` 可直接用作原生 `isize` 范围。

循环中可直接使用元组解构：

```goml
for (key, value) in pairs {
    println(key + value.to_string());
}
```

定长数组使用原生索引循环。`Vec[T]`、`Slice[T]`、`MutSlice[T]` 和所有 `Iterator` 值都有相应的 `IntoIterator` 实现。

<a id="break-continue-and-return"></a>

### `break`、`continue` 和 `return`

```goml
while true {
    if done {
        break
    } else {
        continue
    };
}
```

`break` 和 `continue` 只能出现在循环中。`continue` 始终不带值。`break value` 仅适用于 `loop`；`while` 和 `for` 只接受 `break;`。二者都是发散控制表达式，可出现在 `if`、`match` 或其他值位置。

Rust 风格标签选择外层 `while`、`loop` 或 `for`，不会与 `break value` 冲突：

```goml
'outer: for row in rows {
    for item in row {
        if found(item) {
            break 'outer;
        } else {
            continue;
        }
    }
};
let answer = 'result: loop {
    break 'result 42;
};
```

`continue 'label;` 继续所选外层循环。`break 'label value;` 仅在目标为 `loop` 时可携带值。标签必须引用当前有效的外层循环，外层同名标签有效时不能重复。跨越一个或多个块时，仍会先从内到外运行其 `defer` 清理，再将控制流交给目标。

有标签循环作为语句使用时，下一条语句前需要尾分号。格式化器保留此分号。作为块尾表达式的有标签循环不需要分号。

`return` 可以带值或不带值，并检查当前函数或闭包的返回类型：

```goml
fn first_positive(values: Vec[i32]) -> i32 {
    for value in values {
        if value > 0 {
            return value
        }
    };
    -1
}
```

`unreachable_code` 检查会在无法继续的控制流后还有语句或尾表达式时警告，包括 `return`、`break`、`continue`、无限 `loop`，或所有分支都退出的穷尽 `if` / `match`。编译器对每段连续不可达区域报告一次警告，继续检查其中的其他错误，不会仅因警告而使构建失败。函数或方法可用 `#[allow(unreachable_code)]` 抑制有意的不可达代码警告；闭包继承外层函数的设置。

```goml
#[allow(unreachable_code)]
fn platform_exit() -> () {
    return;
    println("unreachable fallback")
}
```

`unused_function` 检查会警告无法从公开函数、`main`、测试、语言项、顶层初始化器或任何方法体到达的私有顶层函数。分析传递跟踪调用，因此仅互相调用的递归函数或私有函数组仍属于未使用。公开函数视为外部可达 API。对有意未使用的私有函数添加 `#[allow(unused_function)]` 可抑制警告。

表达式语句直接丢弃内置 `Result[T, E]` 时，`unused_result` 检查会警告。检查使用解析后的类型，包括别名、泛型调用结果和跨包返回类型，也检查嵌套块和闭包中的语句。警告不会导致构建失败。

```goml
fn save() -> Result[(), string] {
    Err("save failed")
}

fn main() -> () {
    save();
    let _ = save();
}
```

第一次调用产生警告；第二次显式忽略结果。优先处理错误，或在外层返回类型允许时用 `?` 传播。自由函数或 `impl` 块可用 `#[allow(unused_result)]` 抑制检查；该 impl 中的方法及嵌套闭包继承设置。要在方法中显式忽略某个结果，使用 `let _ = ...`。

此警告的诊断码为 `unused_result`。快速修复会插入 `let _ = `，必要时补上结尾分号。忽略错误是需要有意作出的选择，因此修复需要审阅，不作为首选自动修复。编译器不会建议自动插入 `?`。

此检查只检查最外层类型为内置 `Result` 的值。不会警告 `Option`、其他普通返回值、用户自定义同名 `Result` 类型，或嵌套在元组/容器中的 `Result`。它不跟踪未使用绑定，也不会为检查调用者而特化无约束泛型参数。返回的值和被其他表达式消费的值并非直接丢弃。没有通用 `#[must_use]` 属性。

### `?`

后缀 `?` 仅支持内置 `Option[T]` 和 `Result[T, E]` 语义：

```goml
fn plus_one(value: Option[i32]) -> Option[i32] {
    Some(value? + 1)
}

fn read_number(flag: bool) -> Result[i32, string] {
    let value = parse_number(flag)?;
    Ok(value)
}
```

对 `Option[T]` 使用 `?` 时，最近的函数或闭包必须返回 `Option[_]`。对 `Result[T, E]` 使用时，必须返回具有相同错误类型的 `Result[_, E]`；目前没有 Rust `From` 风格的错误转换。`?` 只求值操作数一次。

诊断区分不支持的操作数（`try_invalid_operand`）、不兼容的外层返回类型（`try_invalid_return`）和不同的 `Result` 错误类型（`try_error_mismatch`）。它们高亮 `?`；传播失败还会指出外层返回标注，没有标注时指出函数名或闭包表达式。错误类型不匹配时显示两种类型，并建议在 `?` 前使用 `map_err` 或修改外层错误类型。这些建议不添加隐式转换。

### `go`

`go` 启动一个返回 `()` 的零参数闭包：

```goml
go || {
    println("background")
};
```

不要写 `go work();`；它会先求值 `work()`，而 `go` 要求 `() -> ()` 类型的值。应写 `go || work();`。

`go` 是分离、非结构化的逃生通道。生命周期不绑定调用者，不返回任务句柄，失败也不传播。调用者必须等待子任务或协调取消时，使用 `std::task`。

`unstructured_go` 检查会对每个 `go` 表达式警告。有意管理分离后台任务的函数或方法可使用 `#[allow(unstructured_go)]`。

<a id="patterns"></a>

## 模式

模式可用于 `let`、`for`、`match`、`if let` 和 `while let`。

| 模式 | 示例 |
| --- | --- |
| 变量绑定 | `value` |
| 通配符 | `_` |
| 空元组 | `()` |
| 布尔值/数字/字符串/字符 | `true`, `42`, `"ok"`, `'x'` |
| 顶层常量 | `ANSWER` 或 `config::ANSWER` |
| 分组 | `(pattern)` |
| 元组 | `(left, right)` |
| 单元素元组 | `(value,)` |
| 精确结构体 | `Point { x, y: other }` |
| 部分结构体 | `Point { x, .. }` |
| newtype | `UserId(inner)` |
| 枚举无载荷变体 | `Color::Red` 或上下文形式 `Red` |
| 枚举元组变体 | `Some(value)` 或限定形式 `Option::Some(value)` |
| 枚举结构体式变体 | `Message::Named { value }` 或上下文形式 `Named { value }` |
| 数组、Vec、Slice 或 MutSlice | `[first, second]`, `[first, .., last]` |
| 剩余部分绑定 | `[first, middle @ .., last]` |
| 别名 | `whole @ Some(value)` |
| 或模式 | `Color::Red \| Color::Blue` |
| 范围 | `0..10`, `'a'..='z'` |
| 嵌套模式 | `Ok((key, value))` |

<a id="refutability-and-binding"></a>

### 可反驳性与绑定

`let` 和 `for` 要求模式不可反驳，即对该位置的类型必须匹配成功。`match`、`if let` 和 `while let` 可以使用可反驳模式：

```goml
let (left, right) = pair;
if let Some(value) = candidate {
    println(value);
}
```

不可反驳性由类型和子模式共同判断，而非只看表面语法。例如，单变体枚举的解构不可反驳，已知长度定长数组的完整解构也不可反驳；Vec[T] 的 [first, ..] 可能因空 Vec 失败。

同一模式分支中不能重复绑定同名变量。或模式的每个候选必须绑定完全相同的变量集合，对应变量类型也必须相同：

```goml
match value {
    Left(shared) | Right(shared) => shared,
}
```

模式绑定仅在对应 `let` 的后续作用域、`for` / `while let` 循环体、`if let` 的 then 分支或 `match` 分支守卫和分支体中可见。在 `if` 或 `while` 的 let 链中，也对后续项可见。

<a id="groups-tuples-structures-and-enumerations"></a>

### 分组、元组、结构体和枚举

`(pattern)` 对模式分组；`(pattern,)` 匹配单元素元组：

```goml
let (only,) = one_tuple;
```

结构体字段模式支持同名简写：

```goml
let Point { x, y } = point;
```

结构体模式默认精确匹配，必须列出所有字段。可用末尾的 `..` 忽略未列出字段：

```goml
let Point { x, .. } = point;
```

结构体模式中的 `..` 最多出现一次，且必须在最后。重复字段、未知字段，以及未写 `..` 时缺少字段，都会产生诊断。

枚举模式的构造器参数数量必须与定义一致。未限定变体由预期枚举类型解析，而非由全局唯一性决定。例如，匹配 `First` 的模式中，`Shared` 表示 `First::Shared`，即使也存在 `Second::Shared`。若 `First` 没有 `Shared`，编译器报告该错误，不会选择 `Second::Shared`。类型别名和导入枚举类型在规范化后参与。嵌套变体的载荷为嵌套模式提供预期类型。

<a id="fixed-arrays-vec-slice-and-mutslice"></a>

### 定长数组、Vec、Slice 和 MutSlice

定长数组 `[T; N]`、`Vec[T]`、`Slice[T]` 和 `MutSlice[T]` 共用方括号模式语法。没有 `..` 时要求精确长度；有 `..` 时，前缀和后缀仅指定最小长度：

```goml
match values {
    [] => "empty",
    [only] => "one",
    [first, .., last] => "many",
}
```

序列模式最多允许一个剩余部分。它可出现在任意位置；`name @ ..` 绑定显式前缀和后缀之间的元素：

```goml
let values: [i32; 4] = [1, 2, 3, 4];
let [first, middle @ .., last] = values;
```

对于定长数组，省略剩余部分时元素数必须恰等于 `N`；包含剩余部分时，显式前缀和后缀的总元素数不能超过 `N`。上例中 `middle` 的类型为 `[i32; 2]`。

对于 `Vec[T]`、`Slice[T]` 和 `MutSlice[T]`，`name @ ..` 的绑定类型为只读 `Slice[T]`。例如 `[head, tail @ ..]` 要求至少一个元素，`tail` 不会被复制成新的 Vec。动态序列可能长度不足，因此此类模式通常放在 `match`、`if let` 或 `while let` 中。

<a id="alias-and-or-pattern"></a>

### 别名与或模式

`name @ pattern` 将整个匹配值绑定到 `name`，同时继续匹配内部模式：

```goml
match value {
    whole @ Some(inner) => use_both(whole, inner),
    None => fallback(),
}
```

`@` 比 `|` 结合更紧。因此 `whole @ A | B` 解析为 `(whole @ A) | B`，通常会因两个候选绑定的变量不同而报错。要让别名覆盖整个或模式，必须加括号：

```goml
whole @ (Left(value) | Right(value))
```

或模式候选从左到右尝试，可嵌套在元组、结构体、枚举和序列模式中。

<a id="range-patterns"></a>

### 范围模式

`start..end` 不包含上界，`start..=end` 包含上界：

```goml
match character {
    'a'..='z' => "lowercase",
    'A'..='Z' => "uppercase",
    _ => "other",
}
```

两端必须是相同具体类型的整数或字符字面量。有符号整数端点可以为负。开上界范围要求下界严格小于上界，闭上界范围要求下界小于等于上界。当前不支持无界范围、浮点范围或字符串范围；序列和结构体中单独的 `..` 表示剩余部分，不是范围。

<a id="exhaustiveness-and-unreachable-branches"></a>

### 穷尽性与不可达分支

编译器检查 `match` 是否穷尽被匹配类型，未穷尽时给出缺失模式示例。分析覆盖 bool、枚举、元组、结构体、定长数组、Vec/Slice/MutSlice、整数和字符范围、别名与或模式，也会警告永远无法匹配的分支。

空 `match` 仅适用于没有可构造值的类型：

```goml
enum Never {}

fn absurd(value: Never) -> i32 {
    match value {}
}
```

穷尽分析忽略载荷类型没有值的枚举变体，也识别没有基础变体的纯递归枚举。例如，`enum MaybeNever { Empty, Filled(Never) }` 只需 `Empty` 分支；`enum Loop { Next(Loop) }` 没有可构造值。

普通守卫可能为假，因此带普通守卫的分支不提供穷尽覆盖；通常需要后续无守卫分支。字符串和浮点字面量无法枚举整个类型，匹配时通常也需要 `_`。浮点模式中，`-0.0` 和 `0.0` 视为相同值。

没有 `ref` 或 `ref mut` 模式。支持对单个绑定使用 `mut name`，不支持对 `dyn Trait` 进行模式匹配。

<a id="closures-and-function-values"></a>

## 闭包与函数值

<a id="closure-syntax"></a>

### 闭包语法

```goml
let add = |left: i32, right: i32| left + right;
let increment = |value: i32| {
    value + 1
};
let greet = || println("hello");
```

闭包参数类型可从预期函数类型或调用位置推断。延迟解析的方法调用也能在嵌套表达式中解析，包括元组、数组和控制流分支。推断有歧义时，标注闭包参数或预期函数类型：

```goml
let transform: (i32) -> string = |value| value.to_string();
```

闭包可以捕获外部变量，或修改捕获的 `let mut`：

```goml
let mut count = 0;
let next = || {
    count = count + 1;
    count
};
```

闭包中的 `return` 退出该闭包并提供结果。

<a id="no-let-polymorphism"></a>

### 没有 let 多态

每个闭包和局部函数值只有一个具体类型：

```goml
let identity = |value| value;
let number = identity(1);
```

之后不能用同一个 `identity` 处理 `string`。需要多态时，定义顶层泛型函数：

```goml
fn identity[T](value: T) -> T {
    value
}
```

顶层函数、闭包和元组变体构造器都可作为函数值传递或返回。
闭包可嵌套并返回另一个捕获环境的闭包，但每个闭包表达式仍只有一个具体函数类型。

<a id="trait-and-impl"></a>

## Trait 与 impl

<a id="define-traits"></a>

### 定义 trait

Trait 方法签名必须包含参数名。第一个接收者可写作 `self`，是 `self: Self` 的简写：

```goml
trait Render {
    fn render(self) -> string;
}

trait Convert[T] {
    fn convert(self, fallback: T) -> T;
}
```

Trait 也可声明无接收者的关联函数。`Self` 从实参或预期返回类型推断：

```goml
trait Decode {
    fn decode(input: string) -> Self;
}

trait Default {
    fn default() -> Self;
}

fn decode_value[T: Decode](input: string) -> T {
    Decode::decode(input)
}

fn default_value[T: Default]() -> T {
    Default::default()
}
```

impl 方法必须保持 trait 声明是否具有 `self: Self` 接收者这一特征。关联函数使用静态分派，不能用方法语法调用，也不参与 `dyn` 分派。若实参和预期结果都不能确定 `Self`，调用因歧义被拒绝。

Trait 可以有父 trait、泛型约束和关联类型：

```goml
trait Provider: Render {
    type Item: ToString;
    fn get(self) -> Self::Item;
}
```

每一步都有其声明 trait 的约束时，关联类型投影可链式连接。例如，`I::IntoIter::Item` 先从 `I` 投影 `IntoIter`，再从迭代器类型投影 `Item`。链式投影可出现在签名和 `where` 等式中：

```goml
trait IntoIterator where Self::Item = Self::IntoIter::Item {
    type Item;
    type IntoIter: Iterator;
}
```

泛型 trait 示例：

```goml
trait Child[T: ToString]: Parent[T] + Render where T: Eq {
    fn child(self) -> string;
}
```

标准比较层次区分偏比较和全等价：

```goml
trait PartialEq {
    fn eq(self: Self, other: Self) -> bool;
}

trait Eq: PartialEq {}

trait PartialOrd: PartialEq {
    fn partial_cmp(self: Self, other: Self) -> Option[Ordering];
}

trait Ord: Eq + PartialOrd {
    fn cmp(self: Self, other: Self) -> Ordering;
}
```

`Eq` 是无方法的标记 trait。相等行为属于 `PartialEq`；全等价类型通过 `fn eq` 实现 `PartialEq`，再添加空的 `impl Eq`。`Eq::eq` 和 `impl Eq` 内的 `eq` 方法均无效。

Trait 方法可以提供默认方法体。impl 可省略该方法并继承默认实现；显式声明的 impl 方法会覆盖它。默认方法体使用 trait 的泛型参数、谓词、关联类型和 `Self` 检查，且跨包仍可用：

```goml
trait Named {
    fn name(self) -> string;
    fn describe(self) -> string {
        "named:" + self.name()
    }
}

impl Named for Point {
    fn name(self) -> string {
        "point"
    }
}
```

Trait 方法也可声明自己的类型参数和约束：

```goml
trait Convert {
    fn convert[U: ToString](self, fallback: U) -> string;
}
```

对应 impl 方法必须具有相同的方法泛型参数数量、签名和约束。泛型 trait 方法使用静态分派，并使该 trait 不能用作 `dyn`。

<a id="trait-impl"></a>

### Trait 实现

```goml
struct Point {
    x: i32,
    y: i32,
}

impl Render for Point {
    fn render(self: Point) -> string {
        self.x.to_string() + "," + self.y.to_string()
    }
}
```

泛型 impl 的类型参数写在 `impl` 之后：

```goml
impl[T: ToString] Render for Box[T] {
    fn render(self: Box[T]) -> string {
        self.value.to_string()
    }
}
```

泛型 trait 的不同应用属于不同 impl：

```goml
impl Convert[i32] for Token {
    fn convert(self: Token, fallback: i32) -> i32 {
        7
    }
}

impl Convert[string] for Token {
    fn convert(self: Token, fallback: string) -> string {
        "seven"
    }
}
```

关联类型在 impl 中绑定：

```goml
impl Iterator for Counter {
    type Item = i32;

    fn next(self: Counter) -> Option[i32] {
        next_value(self)
    }
}
```

实现有父 trait 的 trait 时，还必须为目标类型提供每个父 trait 的 impl。编译器拒绝重叠 impl，并实施孤儿规则：trait 或被实现的名义类型至少一个属于当前包。固有 impl 的目标名义类型必须属于当前包。

<a id="inherent-impl"></a>

### 固有 impl

```goml
impl Point {
    fn new(x: i32, y: i32) -> Self {
        Point { x, y }
    }

    fn sum(self: Self) -> i32 {
        self.x + self.y
    }
}
```

无接收者的关联函数通过 `Point::new(1, 2)` 调用；第一个参数为接收者的方法可通过 `point.sum()` 或 `Point::sum(point)` 调用。

泛型类型的固有 impl：

```goml
impl[T] Box[T] {
    fn get(self: Self) -> T {
        self.value
    }

    fn map[U](self: Self, map_fn: (T) -> U) -> Box[U] {
        Box { value: map_fn(self.value) }
    }
}
```

方法类型实参通常自动推断。需要显式实参时写 `box_value.map::[string](convert)`，关联形式写 `Box::[i32]::map::[string](box_value, convert)`。

具体和嵌套特化也支持关联构造器，例如 `impl Box[f64]` 和 `impl[T] Box[Vec[T]]`。在 `Box::[Vec[string]]::repeated("x")` 中，所属类型实参为 `Vec[string]`，实现参数 `T` 则为 `string`。省略所属类型实参会创建推断变量；`let value: Box[f64] = Box::empty();` 这样的预期结果可选择特化。若仍有多个候选，使用显式所属类型实参；歧义查找会报告诊断。

<a id="method-parsing-and-ufcs"></a>

### 方法解析与 UFCS

对于具体值，使用：

```goml
value.render();
Render::render(value)
```

跨包方法语法要求 trait 位于当前文件的方法作用域：

```goml
use acme::render;
use render::Render;
```

多个可见 trait 定义同名方法时，`value.render()` 会有歧义。使用 UFCS 显式选择：

```goml
let a = A::render(value);
let b = B::render(value);
```

泛型 trait 的 UFCS 在 trait 名后使用 `::[...]` 提供类型参数：

```goml
let number = Convert::[i32]::convert(token, 0);
let text = Convert::[string]::convert(token, "");
```

trait 实现和方法候选搜索不会保留被拒候选的类型绑定。调整其他方面等价的实现声明顺序，不会选择不同的推断类型。

Trait 约束使相应方法可用于泛型值。导入的 trait 短名、重命名导入和公开重新导出，在特化过程中保留泛型约束及 `where` 谓词中定义 trait 的身份。父 trait 和关联类型约束也以隐含约束参与方法解析。

<a id="associated-constants"></a>

### 关联常量

Trait 可要求常量或提供默认值。固有常量遵循方法可见性规则；trait 实现常量继承 trait 可见性。通过 `Type::NAME`、`Type::[Arguments]::NAME` 或 `T::NAME` 这样的受约束参数读取常量。

```goml
trait Capacity {
    const LIMIT: usize;
    const LABEL: string = "buffer";
}

struct FixedBuffer {}

impl FixedBuffer {
    pub const HEADER: usize = 2;
}

impl Capacity for FixedBuffer {
    const LIMIT: usize = 16;
}

fn capacity[T: Capacity]() -> usize {
    T::LIMIT
}
```

初始化器支持字面量、常量路径、一元和二元运算、转换及复合常量构造。不能执行普通函数/方法调用或访问运行时局部变量。泛型实现可在初始化器中使用常量参数。非泛型固有整数常量也可提供数组长度。不支持关联常量模式及长度表达式中的 trait 关联常量。带关联常量的 trait 不能用作 `dyn` trait。

## `dyn Trait`

使用 `as dyn Trait` 将满足 dyn 安全 trait 的具体值转换为 trait 对象：

```goml
trait Display {
    fn show(self) -> string;
}

fn erase[T: Display](value: T) -> dyn Display {
    value as dyn Display
}

fn show_dynamic(value: dyn Display) -> string {
    value.show()
}
```

预期 `dyn Trait` 值时，具体值不会隐式装箱。赋值、实参、返回、分支、构造器和集合操作都必须在转换处写 `as dyn Trait`。已有 `dyn Trait` 值仍可在兼容类型上下文间传递，无需再次转换。

`dyn` 值支持方法语法，也可使用 UFCS；父 trait 方法同样可用：

```goml
value.show();
Display::show(value)
```

关联类型在方括号中固定，可用于 dyn 分派的方法签名：

```goml
trait Source {
    type Item;
    fn get(self: Self) -> Self::Item;
}

fn read(source: dyn Source[Item = isize]) -> isize {
    source.get()
}
```

类型推断拒绝递归类型，包括经 trait 对象关联类型绑定形成的循环。拒绝循环时产生类型诊断，不会创建自引用推断变量。泛型替换和推断解析会遍历嵌套类型构造器，包括 trait 对象实参、关联类型绑定和附加约束。

泛型 trait 在关联类型绑定前接受位置类型实参，例如 `dyn Consumer[string, Error = IoError]`。每个必需泛型实参都须提供，每个关联类型必须且只能绑定一次。首个 `Name = Type` 绑定后的位置实参会被拒绝：

```goml
trait Consumer[T] {
    fn consume(self, value: T) -> string;
}

fn send(value: dyn Consumer[string]) -> string {
    value.consume("hello")
}
```

用 `+` 组合 dyn 安全 trait，通过一个对象暴露其方法。具体值必须以指定实参和关联类型实现每项约束：

```goml
trait Close {
    fn close(self) -> ();
}

fn erase[T: Consumer[string] + Close](value: T) -> dyn Consumer[string] + Close {
    value as dyn Consumer[string] + Close
}

fn finish(value: dyn Consumer[string] + Close) -> string {
    let result = value.consume("done");
    value.close();
    result
}
```

父 trait 方法仍可用。不同组成 trait 提供同名方法时，使用 UFCS 选择 trait，例如 `A::render(value)` 或 `B::render(value)`。泛型 trait 使用 `Consumer::[string]::consume(value, "hello")`。

约束顺序不改变对象类型：`dyn A + B` 与 `dyn B + A` 兼容。重复 trait 名会被拒绝，即使泛型实参不同。父 trait 路径对同一 trait 的泛型实参必须一致。子 trait 方法优先于从父 trait 继承的方法；无关同名方法需要 UFCS。

唯一名称的继承关联类型可在子 trait 上绑定，例如 `dyn Child[Item = isize]`，或在声明它的 trait 上绑定，如 `dyn Child + Parent[Item = isize]`。不同父 trait 声明同名关联类型时，分别在声明 trait 上绑定，如 `dyn Child + A[Item = isize] + B[Item = string]`。

当前 dyn 安全条件：

- 每个方法的第一个接收者参数必须恰为 `Self` 类型；
- 其他参数或返回类型不能直接出现 `Self`，但允许 `Self::Item` 这样的已绑定投影；
- 方法不能声明类型参数；
- Trait 不能声明关联常量；
- 每个组成 trait 和父 trait 都必须满足这些规则。

当前限制：

- dyn 实参列表不支持 `dyn Fixed[4]` 这样的整数常量实参；
- 不支持从一种 dyn 对象类型转换为另一种，包括删除约束；
- 不支持对 `dyn Trait` 进行模式匹配。

<a id="attributes-and-derive"></a>

## 属性与 derive

用户源码支持为结构体和枚举（包括泛型类型）派生 `ToString`、`Debug`、`PartialEq`、`Eq`、`PartialOrd`、`Ord`、`Hash` 和 `Default`。从 `std::serde` 或重新导出它们的格式包导入后，也可使用 `Serialize` 和 `Deserialize`：

```goml
use std::cmp;
use cmp::{Ord, PartialOrd};

#[derive(ToString, Debug, PartialEq, Eq, PartialOrd, Ord, Hash, Default)]
struct Key {
    name: string,
    version: i32,
}

#[derive(ToString, Debug, PartialEq, Eq, PartialOrd, Ord, Hash, Default)]
enum Entry[T, Marker] {
    Empty,
    Value(Vec[T]),
}
```

derive 使用的所有字段或变体载荷必须支持相应 trait。`ToString` 变体模板只使用占位符引用的字段。对泛型定义，生成的 impl 为每个涉及类型参数且实际参与的不同字段/载荷类型添加约束。未使用的 phantom 参数不加约束。编译器拥有的 runtime、intrinsic 和 lang-item 属性仍不可用于普通项目。用户项目可使用下一节的 `go_ffi`、`go_type`、`go_method` 和 `go_interface` 属性。

八个 prelude derive 由工具链内置源码中的已验证处理函数提供，不是编译器硬编码的生成器。`PartialOrd` 和 `Ord` trait 本身位于 `std::cmp`；使用其 derive 时导入该包，使用未限定约束和方法调用时导入 trait。标准库和第三方 derive 使用相同处理函数与产物机制，但不加入 prelude：使用 derive 名前先导入 trait 或其包。

派生的 `ToString` 默认用类型名、变体名和载荷格式化枚举。变体可用字符串模板替换整个输出：

```goml
#[derive(ToString, Debug)]
enum Status {
    #[to_string("ready")]
    Ready,
    #[to_string("retry after {0}s")]
    Retry(isize),
    #[to_string("failed: {message}")]
    Failed { code: isize, message: string },
    Unknown,
}
```

`Status::Ready.to_string()` 返回 `"ready"`，`Status::Retry(3).to_string()` 返回 `"retry after 3s"`，`Status::Unknown.to_string()` 保持 `"Status::Unknown"`。具名载荷使用 `{field}`；元组载荷使用从零开始的规范十进制索引，如 `{0}`。`{{` 和 `}}` 产生字面花括号。接受普通和原始字符串字面量，包括空字符串和 Unicode 文本。模板可省略任意或全部载荷字段。引用字段按声明顺序各转换一次，重复占位符复用已转换字符串。仅引用的字段需要 `ToString`；泛型谓词还保留所有未自定义变体的要求。字符串载荷直接插入。

`to_string` 属性要求恰好一个位置字符串字面量，每个变体最多出现一次。重复属性、无效实参、未匹配花括号、未知字段和越界元组索引都是定位于属性或模板实参的编译期错误。派生 `ToString` 时，将此属性放在枚举、结构体或载荷字段上是错误。模板不支持隐式 `{}` 占位符、表达式、格式说明符、回调或自动大小写转换。模板由 `ToString` derive 解释，独立于 `Debug` 和 `serde(rename)`，不生成字符串解析，也不要求输出唯一。复杂格式化请手写 `impl ToString`。

派生的 `Debug` 提供 `Debug::debug(value)` 和 `.debug()` 方法。结构体显示类型名和字段名，枚举显示类型名和变体名。基本字段值使用普通文本表示，嵌套值递归使用 `Debug`。

`PartialEq` 按声明顺序比较结构体字段，并结构化比较枚举变体及载荷。`Eq` 只生成标记实现，不隐式派生 `PartialEq`。因此通常用 `#[derive(PartialEq, Eq)]` 表示全等价。

`PartialOrd` 和 `Ord` 按字段顺序做字典序比较。枚举先按变体声明顺序排序，再比较载荷字段。`PartialOrd` 立即传播 `None`；`Ord` 返回确定的 `Ordering`。两者都不隐式生成父 trait，因此全序通常写作 `#[derive(PartialEq, Eq, PartialOrd, Ord)]`。

`Hash` 按声明顺序组合字段，并包含枚举变体判别值。可以独立派生，但哈希表键仍同时需要 `Eq` 和 `Hash`。

`Default` 用 `Default::default()` 初始化每个结构体字段。对枚举始终选择第一个声明的变体，并递归默认初始化其元组或具名载荷。空枚举不能派生 `Default`。因此重排枚举变体会同时改变派生默认值和派生顺序；不支持 `#[default]`。

<a id="programmable-derive"></a>

### 可编程 derive

依赖包可导出编译期 derive 处理函数：

```goml
package labels;

pub trait Label {
    fn label(self: Self) -> string;
}

#[comptime_derive(Label)]
pub fn derive_label(input: DeriveInput) -> DeriveOutput {
    let output = derive_output_new(input, "Label");
    let params = meta_param_list_new();
    meta_param_list_push(params, "self", derive_target_type(input));
    let body = meta_expr_string(derive_item_name(input));
    let method = meta_method("label", params, meta_type_call_site("string"), body);
    derive_output_add_method(output, method);
    output
}
```

其他包通过导入的包路径应用处理函数：

```goml
use alice::labels as labels;
use labels::Label;

#[derive(Label)]
struct User {
    name: string,
}

#[derive(labels::Label)]
struct Group {
    name: string,
}
```

公开 `#[comptime_derive(Name)]` 处理函数导出 derive 名 `Name`，独立于实现函数名。导出名在同包唯一，不能包含 `::`。私有具名处理函数可在本包使用；跨包导入需要 `pub`。私有无名 `#[comptime_derive]` 函数可保留为实现辅助函数。仍支持公开无名形式，导出处理函数的短名。

Derive 名有独立命名空间，但普通 `use` 声明会同时填充它与类型/trait 命名空间。例如，`use std::serde;` 后，`use serde::Serialize;` 使 `Serialize` trait 和名为 `Serialize` 的 derive 导出都可通过 `Serialize` 使用，因此 `#[derive(Serialize)]` 无需其他导入形式。导入别名同时作用于两种命名空间：`use serde::Serialize as DataSerialize;` 允许 `#[derive(DataSerialize)]`。包导入允许限定写法 `#[derive(serde::Serialize)]`。公开重新导出保留处理函数的定义身份，因此门面包可用 `pub use` 暴露标准库或第三方 derive，无需复制实现。

处理函数也可在同包定义和使用，包括跨源文件。编译器在展开 derive 前准备该包的编译期函数。准备阶段不能使用依赖该展开的方法或类型；这种循环会产生诊断。本地 derive 名优先于导入的 derive 名。

生成 API 包括：

| 能力 | 构建器 |
| --- | --- |
| 闭包与函数类型 | `meta_expr_closure(params, body)`, `meta_type_function(params, result)`, `meta_expr_invoke(function, args)` |
| 限定泛型成员调用 | `meta_expr_type_member(owner_type, member, member_type_args, args)` |
| 多个实现 | `derive_output_merge(output, additional_output)` |
| 关联项 | `derive_output_add_associated_type(output, name, type)`, `derive_output_add_constant(output, name, type, value)` |
| Trait 参数与等式 | `derive_output_set_trait_arguments(output, args)`, `derive_output_add_equality(output, lhs, rhs)` |
| 私有辅助声明 | `derive_output_add_function(output, method)`, `derive_output_add_struct(output, name, fields)`, `derive_output_add_type_alias(output, name, type)` |
| 结构化属性语法 | `meta_attribute_argument_type(attribute, index)`, `meta_attribute_argument_expr(attribute, index)` |
| 源码位置 | `derive_span(input)`, `derive_field_span(input, index)`, `derive_variant_field_span(input, variant, field)`, `meta_attribute_span(attribute)`, `meta_attribute_argument_span(attribute, index)` |
| 精确诊断 | `derive_error_at(attribute, message)`, `derive_error_at_argument(attribute, index, message)`, `derive_error_span(span, message)` |

`meta_expr_closure` 接受 `MetaParamList`；`meta_type_function` 接受 `MetaTypeList`。辅助结构体接受包含字段名和类型的 `MetaParamList`。辅助名称请使用 `derive_fresh_name`。关联常量遵循与手写关联常量相同的表达式限制。`derive_set_span(span)` 为随后生成的语法设置给定位置。

结构化属性访问保留具名和位置实参、嵌套分隔符和原始字符串。实参值种类包括 `ident`、`path`、`string`、`bool`、`integer`、`float`、`char` 和 `syntax`。类型和表达式访问器验证单个类型或表达式，并在 derive 使用位置解析名称。例如，`#[config(priority = 10, extras = Option[State], callback = make_callback(2))]` 有三个实参。定义位置成员构建器会限定 `Rule::regex` 这样的相对类型所属者；显式调用位置构建器保留派生文件的作用域。

静态固有方法可用作函数值：`let callback = Parser::parse;`。解析后的签名提供函数值的参数和返回类型；泛型方法值仍需足够上下文推断类型实参。

标准库通过 `std::serde` 提供序列化 derive。API、属性和示例见 [Serde 与 JSON](stdlib.zh-CN.md#serde-and-json)。

限定自定义 derive 必须为 `package_alias::export_name`，其中 `package_alias` 由同一源文件中的 `use` 引入。不带限定的库 derive 必须显式导入。未导入的规范包路径及仅通过传递依赖可见的处理函数会被拒绝。若两个显式导入，或 prelude derive 与显式导入提供相同短名，编译器报告歧义，要求限定名或导入别名。

`#[comptime_derive]` 和 `#[comptime_derive(Name)]` 仅适用于非泛型自由函数。公开或具名处理函数必须具有精确签名 `(DeriveInput) -> DeriveOutput`。带无名属性的私有函数是仅编译期辅助函数，从公开处理函数可达时会包含在接口中。derive 处理函数可调用这些辅助函数及普通 `#[comptime]` 函数，但不能从运行时代码或普通值 `comptime` 调用。处理函数不作为运行时函数导出。

编译器从已编译依赖接口解析导入处理函数，并在展开当前包前准备本地处理函数。可复用处理函数及其生成的 trait 可位于独立包。目标可为泛型结构体或枚举；生成的 impl 继承其类型参数。显式所属类型实参也会特化参数和返回类型未提及这些实参的静态固有方法，例如 `Record::[string]::type_name()`。`derive_output_add_predicate` 和 `derive_output_add_call_site_predicate` 添加生成方法所需的约束。

输入反射操作如下：

```text
derive_item_name(input) -> string
derive_item_kind(input) -> isize
derive_generic_count(input) -> isize
derive_generic_name(input, index) -> string
derive_field_count/name/type(...)
derive_variant_count/name(...)
derive_variant_kind(input, variant) -> isize
derive_variant_field_count/name/type(...)
derive_attribute_count/name/text(...)
derive_field_attribute_count/name/text(...)
derive_variant_attribute_count/name/text(...)
derive_variant_field_attribute_count/name/text(...)
derive_attribute(input, index) -> MetaAttribute
derive_field_attribute(input, field, index) -> MetaAttribute
derive_variant_attribute(input, variant, index) -> MetaAttribute
derive_variant_field_attribute(input, variant, field, index) -> MetaAttribute
derive_target_type(input) -> MetaType
derive_fresh_name(input, prefix) -> string
```

`derive_item_kind` 对结构体返回零，对枚举返回一。`derive_variant_kind` 对无载荷变体返回零，对元组变体返回一，对结构体式变体返回二。嵌套属性的计数操作接受所属者索引；名称和文本操作额外接受一个属性索引。字段操作要求结构体，变体操作要求枚举。无效种类和索引会在 `#[derive(...)]` 位置产生编译期错误。属性可附着在结构体字段、枚举变体及元组或具名变体字段上。`name` 返回属性名，`text` 返回完整源码写法。

结构化属性句柄暴露 `meta_attribute_name`、`meta_attribute_text`、`meta_attribute_has_argument_list`、`meta_attribute_argument_count`、`meta_attribute_argument_kind`、`meta_attribute_argument_name`、`meta_attribute_argument_value_kind` 和 `meta_attribute_argument_text`。实参种类为上列字面量/语法种类或 `named`。具名实参使用 `name = value`，值也可为数字、布尔值、嵌套类型或表达式。`argument_name` 返回左侧名称，`argument_value_kind` 描述右侧值，`argument_text` 返回解码后的值。

结构化输出 API 提供不透明的 `MetaSpan`、`MetaAttribute`、`MetaType`、`MetaExpr`、`MetaPattern`、`MetaArm`、`MetaBlock`、`MetaParamList`、`MetaGenericList`、`MetaMethod` 和列表句柄。构造器使用 `meta_type_*`、`meta_expr_*`、`meta_pattern_*`、`meta_arm*`、`meta_block_*`、`meta_param_list_*` 和 `meta_generic_list_*` 系列。`meta_method` 创建具体方法；`meta_method_generic` 创建带显式类型参数和约束的方法。处理函数通过 `derive_output_new` 或 `derive_output_new_call_site` 从一个 trait impl 开始，或通过 `derive_output_inherent` 从一个固有 impl 开始，再添加谓词和方法并返回输出。`derive_output_add_method` 保留固有方法的私有可见性及 trait 实现的普通 trait 可见性。`derive_output_add_public_method` 导出固有方法，并拒绝 trait 输出。

例如，独立包中的处理函数可生成公开静态方法：

```goml
#[comptime_derive(TypeName)]
pub fn type_name(input: DeriveInput) -> DeriveOutput {
    let output = derive_output_inherent(input);
    derive_output_add_public_method(
        output,
        meta_method(
            "type_name",
            meta_param_list_new(),
            meta_type_call_site("string"),
            meta_expr_string(derive_item_name(input)),
        ),
    );
    output
}
```

构建器操作如下：

```text
meta_type_named(input, name) -> MetaType
meta_type_call_site(name) -> MetaType
meta_type_list_new() -> MetaTypeList
meta_type_list_push(list, type) -> ()
meta_type_tuple(list) -> MetaType
meta_type_apply(type, arguments) -> MetaType
meta_type_array(element, length) -> MetaType
meta_type_projection(type, associated_name) -> MetaType
meta_type_kind/name(type) -> string
meta_type_argument_count/argument(type, index...) -> isize | MetaType
meta_type_tuple_count/tuple_item(type, index...) -> isize | MetaType
meta_type_array_length/array_element(type) -> isize | MetaType
meta_type_function_parameter_count/parameter(type, index...) -> isize | MetaType
meta_type_function_return(type) -> MetaType
meta_type_contains_generic(input, type) -> bool
meta_type_equal(left, right) -> bool

meta_expr_var(name) -> MetaExpr
meta_expr_unit/bool/isize/string/char(value...) -> MetaExpr
meta_expr_integer(text, type) -> MetaExpr
meta_expr_field(value, name) -> MetaExpr
meta_expr_index(value, index) -> MetaExpr
meta_expr_unary(operator, value) -> MetaExpr
meta_expr_binary(operator, left, right) -> MetaExpr
meta_expr_call(input, name, arguments) -> MetaExpr
meta_expr_call_site(name, arguments) -> MetaExpr
meta_expr_trait_call(input, trait_name, method_name, arguments) -> MetaExpr
meta_expr_method_call(receiver, name, arguments) -> MetaExpr
meta_expr_constructor(name, arguments) -> MetaExpr
meta_expr_target_constructor(input, variant, arguments) -> MetaExpr
meta_expr_tuple/array(elements) -> MetaExpr
meta_expr_field_list_new() -> MetaExprFieldList
meta_expr_field_list_push(list, name, value) -> ()
meta_expr_struct(input, name, fields) -> MetaExpr
meta_expr_struct_call_site(name, fields) -> MetaExpr
meta_expr_target_struct(input, fields) -> MetaExpr
meta_expr_if(condition, then, else) -> MetaExpr
meta_expr_while(condition, body) -> MetaExpr
meta_expr_match(value, arms) -> MetaExpr
meta_expr_cast(value, type) -> MetaExpr
meta_expr_return(value) -> MetaExpr
meta_expr_try(value) -> MetaExpr
meta_expr_list_new() -> MetaExprList
meta_expr_list_push(list, expression) -> ()

meta_pattern_wild() -> MetaPattern
meta_pattern_bind(name) -> MetaPattern
meta_pattern_unit/bool/isize/string/char(value...) -> MetaPattern
meta_pattern_tuple/array(patterns) -> MetaPattern
meta_pattern_constructor(name, patterns) -> MetaPattern
meta_pattern_target_constructor(input, variant, patterns) -> MetaPattern
meta_pattern_field_list_new() -> MetaPatternFieldList
meta_pattern_field_list_push(list, name, pattern) -> ()
meta_pattern_struct(name, fields, has_rest) -> MetaPattern
meta_pattern_target_struct(input, fields, has_rest) -> MetaPattern
meta_pattern_alias(name, pattern) -> MetaPattern
meta_pattern_or(patterns) -> MetaPattern
meta_pattern_range(start, end, inclusive) -> MetaPattern
meta_pattern_list_new() -> MetaPatternList
meta_pattern_list_push(list, pattern) -> ()
meta_arm(pattern, expression) -> MetaArm
meta_arm_guarded(pattern, guard, expression) -> MetaArm
meta_arm_list_new() -> MetaArmList
meta_arm_list_push(list, arm) -> ()

meta_block_new() -> MetaBlock
meta_block_let(block, name, value) -> ()
meta_block_let_mut(block, name, value) -> ()
meta_block_let_mut_typed(block, name, type, value) -> ()
meta_block_let_typed(block, name, type, value) -> ()
meta_block_let_pattern(block, pattern, value) -> ()
meta_block_assign(block, target, value) -> ()
meta_block_expr(block, expression) -> ()
meta_block_finish(block, tail) -> MetaExpr
meta_block_finish_unit(block) -> MetaExpr
meta_param_list_new() -> MetaParamList
meta_param_list_push(list, name, type) -> ()
meta_generic_list_new() -> MetaGenericList
meta_generic_list_push(list, name) -> ()
meta_generic_list_add_bound(input, list, name, trait_name) -> ()
meta_method(name, parameters, return_type, body) -> MetaMethod
meta_method_generic(name, generics, parameters, return_type, body) -> MetaMethod

derive_output_new(input, trait_name) -> DeriveOutput
derive_output_new_call_site(input, trait_name) -> DeriveOutput
derive_output_inherent(input) -> DeriveOutput
derive_output_add_predicate(output, type, trait_name) -> ()
derive_output_add_call_site_predicate(output, type, trait_name) -> ()
derive_output_add_method(output, method) -> ()
derive_output_add_public_method(output, method) -> ()
```

`meta_expr_unary` 用运算符编号 `0..2` 分别表示 `-`、`!`、`~`。`meta_expr_binary` 用 `0..17` 分别表示 `+`、`-`、`*`、`/`、`%`、`&`、`|`、`^`、`<<`、`>>`、`&&`、`||`、`<`、`>`、`<=`、`>=`、`==` 和 `!=`。`meta_expr_integer` 接受规范化整数字面量字符串及精确整数类型，因此构建器可表示宿主 `isize` 范围外的值。`meta_expr_cast` 仅接受 `dyn` 目标类型；生成数值转换应使用 `meta_expr_method_call`。`meta_expr_trait_call` 在处理函数定义包解析 trait，构建静态 trait 方法调用。`meta_type_equal` 比较结构类型身份。`meta_type_kind` 返回 `primitive`、`named`、`tuple`、`application`、`array`、`function` 或 `dyn`；形态专用访问器拒绝其他种类。列表句柄仅可通过对应 `push` 操作修改，且限于单次 derive 求值。

传给 `derive_output_new`、`derive_output_add_predicate`、`meta_type_named`、`meta_expr_call` 和 `meta_generic_list_add_bound` 的未限定名在处理函数定义包解析。其 `_call_site` 变体在目标包解析。必须避免与用户名称冲突的生成局部绑定应使用 `derive_fresh_name`。`*_target_*` 构建器通过编译器身份构造或匹配被标注项，应优先于手动拼写名称。

结果可包含被标注类型的多个 trait 或固有实现、关联类型和常量，以及私有辅助函数、结构体和别名。不能创建 trait、模块、导入、extern 声明、属性或原始词法单元。支持生成方法类型参数和 trait 约束。生成的 impl 经过普通名称解析、孤儿与一致性检查、类型检查、单态化和后端降级。重复或无效的生成实现作为普通编译器诊断报告。

处理函数具有确定性，不能访问宿主。导入的 derive CTIR 作为不可信产物数据验证。求值使用普通编译期限制，额外限制为 100,000 次元数据和语法构建器操作。失败默认定位到请求 derive 的属性；支持 span 的错误操作可选择字段或属性实参。诊断包含编译期 derive 调用栈。

## Go FFI

可运行的[文件读取夹具](../goml/testdata/ffi/file-read/README.zh-CN.md) 综合展示原始指针、方法、共享字节缓冲区、部分读取、错误身份、显式 Close 和 Go↔GoML 结果传输。

类型化 Go FFI 将顶层 GoML 声明绑定到导出的包级 Go 函数：

```goml
#[go_ffi("strings", "ToUpper")]
extern fn to_upper(value: string) -> string;

#[go_ffi("strings", "Cut")]
extern fn cut(value: string, separator: string) -> (string, string, bool);

fn example() -> string {
    let (before, after, found) = cut("left:right", ":");
    if found {
        to_upper(before + after)
    } else {
        ""
    }
}
```

属性第一个实参是 Go 导入路径，第二个是导出的 ASCII Go 标识符。GoML 函数名是本地名，可不同于 Go 符号。在 `extern fn` 前加 `pub` 可通过其他 GoML 包暴露绑定；接口和 Core 产物保留 Go 导入路径及符号。

使用 `#[go_type("time", "Duration")] pub extern type Duration;` 绑定不透明 Go 类型。GoML 名可不同于 Go 对象名。声明不能有主体或 `= Type` 目标。导入路径、Go 名和类型实参共同标识类型：`Duration` 不同于 `i64`，而两个 GoML 包绑定的 `time.Duration` 共享类型。值可通过普通 GoML 函数和显式类型的外部调用传递；生成的 Go 使用原始 Go 类型。编译器不检查外部字段、不规范化 Go 别名，也不发现方法和接口。不可访问外部字段或构造外部结构体；请使用显式函数/方法绑定或 Go shim。

```goml
#[go_type("time", "Duration")]
extern type Duration;

#[go_ffi("time", "Sleep")]
extern fn sleep(duration: Duration) -> ();

fn wait(duration: Duration) -> () {
    sleep(duration)
}
```

对于 Go 的 `type Point struct { X, Y int64 }`，从其 Go 包绑定 `NewPoint`、`PointX` 等显式辅助函数：

```goml
#[go_type("example.com/geometry", "Point")]
extern type Point;

#[go_ffi("example.com/geometry", "NewPoint")]
extern fn point(x: i64, y: i64) -> Point;

#[go_ffi("example.com/geometry", "PointX")]
extern fn point_x(value: Point) -> i64;
```

`use std::ffi;` 暴露 `ffi::Ptr[T]`，它是指向 T 的原始 Go 指针的透明写法。`ffi::null()` 构造类型从上下文推断的 nil 指针，`ffi::is_nil(value)` 检查它。非空指针通过显式 Go 绑定进入。例如，`let empty: ffi::Ptr[Cell] = ffi::null();` 声明 nil 指针。这些函数也可作为普通函数值。指针在函数、闭包或容器间传递不复制所指对象。包含 nil 的 `Option::Some` 仍不同于 `Option::None`。GoML 不能从整数或 `Ref` 构造指针。独立声明的 Go 别名如 `extern type Handle` 仍不透明，不会规范化为 `ffi::Ptr[Cell]`。它不能传给 `ffi::is_nil()` 或接收 `ffi::null()` 的值；请在绑定签名中使用显式指针表示，或为不透明类型绑定 Go 辅助函数。没有 `*T` 源类型语法。

通过 `#[go_method("Method")] extern fn` 绑定 Go 方法；第一个参数为接收者。声明必须单态，且恰好有一个 ABI 属性。方法名必须为导出的 ASCII Go 标识符。最终 Go 构建检查接收者实际方法集：仅指针方法要求适当的指针接收者，值方法也可通过指针调用。不会对复制的值隐式取地址。接口和提升方法遵循相同 Go 规则。绑定可公开、跨包导入或作为普通函数值传递，且始终仅限运行时，即使从产物加载也一样。

```goml
use std::ffi;

#[go_type("example.com/shim", "Cell")]
extern type Cell;

#[go_method("Read")]
extern fn read(cell: ffi::Ptr[Cell]) -> i64;

#[go_method("Write")]
extern fn write(cell: ffi::Ptr[Cell], value: i64) -> ();
```

这些绑定生成 `(*shim.Cell).Read(cell)` 等 Go 方法表达式。nil 接收者原样传给 Go 方法，包括该方法定义的 panic 或特殊 nil 行为。方法不会自动分配对象、关闭资源或将错误转换为 `Result`。取得非空值仍需显式 Go 绑定。

`ffi::Error` 是原始的可空 Go `error` 接口。`ffi::nil_error()` 构造 nil 接口；`ffi::error_is_nil(value)` 检查接口是否为 nil。持有带类型 nil 指针的接口不是 nil。错误在函数、闭包和产物间保留原对象及错误链。原始 `(T, error)` 返回仍是元组；包含 nil 错误的 `Option::Some` 仍不同于 `None`。

`ffi::NonNilError::from_error(error)` 对 nil 接口返回 `None`，对非 nil 接口（含带类型 nil 错误）返回 `Some`。其私有存储不能直接构造；`as_error()` 返回原错误。因此 `Result[T, ffi::NonNilError]` 不能接受 `Err(ffi::nil_error())`。此包装器不调用 `Error()`，也不检查具体错误载荷。

`ffi::Outcome::new(value, error)` 在公开 `value` 和 `error` 字段中保留两个结果，包括失败时同时返回的部分值。显式命名的 `into_result_discarding_value_on_error()` 适配器返回 `Result[T, ffi::NonNilError]`：仅成功时保留值，失败时丢弃值。此转换不会自动发生。例如：

```goml
let (count, error) = read(buffer);
let outcome = ffi::Outcome::new(count, error);
let partial_count = outcome.value;
let result = outcome.into_result_discarding_value_on_error();
```

反向转换也需显式进行。`ffi::Outcome::from_raw_result(result, failure_value)` 接受 `Result[T, ffi::Error]`，返回 `Result[ffi::Outcome[T], ffi::BoundaryError]`。`Ok(value)` 产生该值和 nil Go 错误。`Err(error)` 使用调用者提供的失败值并保留错误。若错误是真正的 nil 接口，转换以 `BoundaryError::NilError` 失败，消息为 `Go FFI boundary rejected Err(nil)`。带类型 nil 错误仍是失败。即使输入为 `Ok`，失败值实参也正常求值。

`ffi::Outcome::from_result(result, failure_value)` 接受 `Result[T, ffi::NonNilError]` 并应用相同边界检查。`outcome.into_tuple()` 返回用于显式 Go 调用的精确 `(value, error)` 对；不会规范化部分结果或 nil 值。这些适配器都不臆造默认失败值，也不会隐式将契约错误变成 Go 成功。

`ffi::error_matches(value, target)` 使用 Go `errors.Is`，包括包装链、自定义 `Is` 方法和 Go 的 nil 匹配语义。`NonNilError::matches(target)` 为已检查包装器提供相同操作。匹配比较原始错误对象和错误链，从不比较消息字符串。

`NonNilError::message_bytes()` 调用原始 Go `Error()` 方法，返回其字节的新 `std::bytes::Bytes` 副本。`message()` 返回 `Result[string, std::utf8::Utf8Error]`，接受有效 UTF-8，拒绝无效字节而不替换。访问消息不替换存储的错误。带类型 nil 接口可进入已检查包装器，因此底层 Go 方法的 nil 处理、副作用和 panic 仍由其自身决定。这些 API 不捕获 Go panic，也不隐式将错误转为字符串。

泛型外部声明如 `#[go_type("example.com/shim", "Box")] extern type Box[T];` 接受 `Box[i64]`，以及泛型 GoML 代码中的符号形式 `Box[T]`。GoML 检查声明参数数量，并跨包和产物保留外部身份及实参。不查询 Go 声明、不检查约束、不展开别名。最终 Go 编译检查特化后可达的具体实例，包括实际类型和实参能否作为合法运行时值。未使用的泛型函数可能不产生 Go 实例。查询和 LSP 分析使用同样的声明不透明类型。

外部调用 ABI 支持生成的 Go 表示可直接赋值的值。原始边界类型和适配器详见后续各节：

| GoML 类型 | Go 表示 |
| --- | --- |
| `bool`、数值基本类型、`string` | 对应 Go 基本类型 |
| 带具体实参的外部具名类型 | 原始 Go 包类型，不创建新的定义类型 |
| `ffi::Ptr[T]` | 可空 Go 指针 `*T` |
| `ffi::Error` | 可空 Go `error` 接口 |
| `ffi::String` | 保留任意字节的 Go `string` |
| `ffi::Rune` | 不做标量验证的 Go `rune` / `int32` |
| `ffi::RawSlice[T]` | 可空 Go 切片 `[]T` |
| `ffi::RawMap[K, V]` | 键可比较的可空 Go 映射 `map[K]V` |
| `char` | `rune` / `int32` |
| `byte` | `byte` / `uint8` |
| `[T; N]` | `[N]T` |
| `Slice[T]` | `[]T` |
| `MutSlice[T]` | `[]T` |
| `ffi::Func[F]` | 签名为 F 的可空 Go 函数；空元组/扁平元组结果对应零个/多个 Go 结果 |
| `Channel[T]` | `chan T` |
| `Sender[T]` | `chan<- T` |
| `Receiver[T]` | `<-chan T` |
| 直接 `dyn Marker` 参数或单个返回值 | 通过 trait 对象 `data` 字段传递 Go `any` |
| `()` 返回 | 无结果的 Go 函数 |
| `(A, B, ...)` 返回 | 同顺序的多个 Go 结果 |

元组类型仅支持作为完整返回类型。元组元素及数组、切片、通道元素本身也必须满足 FFI 安全。参数不能是 `()`。

`dyn Marker` 仅支持直接参数或完整的单返回类型。`Marker` 必须严格为空：不能声明泛型参数、谓词、父 trait、关联类型或方法，`dyn` 类型不能有实参、关联类型绑定或附加约束。产物验证在声明函数的包内解析本地未限定标记名，包括私有标记。包装器将对象的 `data` 字段作为 Go `any` 传递，再将返回 Go 值包装回空标记对象。空标记对象暂不支持嵌套在元组、数组、切片或通道内。

声明必须单态。GoML 检查声明的边界形态，并信任显式签名。`goml check` 和编辑器分析不检查原生包。最终 Go 构建根据实际符号、接收者方法集、可赋值性和泛型约束验证生成调用。错误声明在该构建步骤失败；未使用绑定可能不生成，因此可能不检查。即使底层表示匹配，`time.Duration` 这样的 Go 具名类型仍不同于 GoML `i64`。

最终可执行文件构建使用选定的 Go 可执行程序、目标、标志、模块策略和生成调用者目录。Go 检查导入合法性，包括 `internal` 限制。普通构建禁用 Go 工作区模式，并根据模块根是否存在 `go.mod` 选择模块模式；其他 Go 依赖和工具链行为遵循用户 Go 环境。不会额外实施 FFI 专用只读或离线策略。Go 库导出和原生 SIMD 构建保留只读、离线模块处理，请单独准备依赖。Go 构建成功不能证明 Unicode、可变性、生命周期或并发契约。返回空元组的旧绑定可能丢弃 Go 结果；普通调用保留 Go 泛型推断和可变参数调用规则。

此 ABI 不支持 `Vec`、`Ref`、`HashMap`、`Option`、`Result`、用户结构体和枚举、非空 trait 对象、普通 GoML 函数值、嵌套元组、泛型函数声明、隐式回调转换、自动 Go 对象生命周期管理和自动 `error` 转换。适配此类 API 时，编写带导出函数和 FFI 安全参数的小型 Go shim：

```go
package goshim

import "os"

func ReadText(name string) (string, bool) {
    data, err := os.ReadFile(name)
    return string(data), err == nil
}
```

```goml
#[go_ffi("example.com/myapp/goshim", "ReadText")]
extern fn read_text(name: string) -> (string, bool);
```

GoML 模块根包含 `go.mod` 时，`goml build`、`goml run` 和测试链接以模块模式调用 Go，从而可导入本地 shim 包及声明的 Go 模块依赖。Go 工作区模式仍禁用。无 `go.mod` 时保留现有关闭模块的行为。模块模式 Go 构建始终执行，并将依赖和源码新鲜度交给 Go 自身构建缓存，因此会观察到 `go.mod`、`go.sum` 和本地 `.go` shim 的变化。

包函数形式从 GoML 调用 Go；方法调用使用上述显式 `go_method` 形式。动态符号查找和 C ABI 互操作需要独立机制。

<a id="go-interface-wrappers"></a>

### Go 接口包装器

模块编译中，顶层 trait 上的 `go_interface` 生成持有显式绑定的原生 Go 接口并转发方法的包装器：

```goml
use std::ffi;

#[go_type("io", "Reader")]
pub extern type GoReader;

#[go_interface(GoReader, ReaderAdapter, read = "Read")]
pub trait Reader {
    fn read(self: Self, buffer: ffi::RawSlice[u8]) -> (isize, ffi::Error);
}

fn wrap(raw: GoReader) -> ReaderAdapter {
    ReaderAdapter::from_go(raw)
}
```

第一个实参是原始类型路径，可指向导入的外部类型。第二个是未限定包装器名，其可见性随 trait。具名字符串实参将每个声明的 trait 方法映射到恰好一个导出 Go 方法。GoML 检查 trait 映射完整性、重复/未导出名称和声明的原始边界类型。生成类型/辅助名称及保留包装器成员不能与用户声明冲突。最终 Go 构建检查真实原生接口和方法签名；GoML 不发现其方法集。

`ReaderAdapter::from_go(raw)` 将原接口值存入私有字段；`wrapper.into_go()` 返回它。二者保留 nil 接口、带类型 nil 接收者及所引用 Go 对象身份。转发调用保留 Go 接收者行为，包括 nil 处理和 panic。不替换接收者，也不自动关闭。原始 Go 接口仍不同于 `dyn Reader`；用 `wrapper as dyn Reader` 显式转换为 GoML trait。

方法签名使用显式原始边界：文本使用 `ffi::String`/`ffi::Rune`，切片使用 `ffi::RawSlice`。在调用方显式应用复制/共享适配器。所有 Go 结果均保留：`()` 表示无结果，`(isize, ffi::Error)` 同时保留部分计数和原错误。普通 GoML 文本、只读切片、隐式闭包和运行时容器不会自动转换成接口方法边界。

`ReaderAdapter::from_trait(value: dyn Reader)` 创建由生成的 Go 桥接支撑的包装器。用 `ReaderAdapter::from_trait(implementation as dyn Reader).into_go()` 将 GoML 实现传给 Go。类型化转发闭包在创建函数返回后仍保留所提供 trait 对象。Go 调用保留共享捕获状态和每个返回分量；panic 在同一 goroutine 传播。生成的桥接包含编译期 Go 接口满足性断言。结果接口非 nil，即使提供的 trait 对象包装 nil 原始接口也如此。`from_go`、`into_go` 和 `from_trait` 是包装器保留方法名。取消订阅和关闭资源仍由调用者显式负责。

适配器要求原始外部类型和带实例方法的非泛型 GoML trait，不能有父 trait、关联类型或额外方法谓词。显式签名和映射提供生成桥接所需信息。最终 Go 构建检查原始类型是否为可用接口，以及生成桥接是否满足它，包括嵌入方法和签名兼容性。仅约束用途的接口不能用作 Go 运行时值。

<a id="raw-go-function-values"></a>

### 原始 Go 函数值

`ffi::Func[F]` 是可空原始 Go 函数值，不同于签名 F 的 GoML 闭包。例如，`ffi::Func[(i64) -> (i64, ffi::Error)]` 表示 Go `func(int64) (int64, error)`，`ffi::Func[() -> ()]` 表示 `func()`。F 必须解析为参数和结果均 FFI 安全的函数签名；函数不能作原始映射键。

`ffi::func_nil()` 构造签名从上下文推断的 nil，`ffi::func_is_nil(value)` 检查它。Go 返回的值可通过包、泛型函数、容器和捕获它的 GoML 闭包传递，不丢失 Go 闭包环境。包含 nil 函数的 `Option::Some` 仍不同于 `Option::None`。Go 绑定可接收这些原始值并在 Go 中调用。不能用 GoML 调用语法直接调用原始函数，不支持隐式闭包转换或结构体构造。

```goml
use std::ffi;

fn empty_callback() -> ffi::Func[(i64) -> i64] {
    ffi::func_nil()
}
```

`ffi::func_from_closure(value)` 显式将 GoML 函数或闭包转为非 nil 原始 Go 函数。`ffi::func_to_closure(value)` 对 nil 返回 `None`，否则返回可调用 GoML 闭包 `Some(closure)`。两个方向都保留捕获状态；转换同签名的两个闭包时仍保留各自环境。捕获的 `Ref` 单元保持普通共享行为。

```goml
use std::ffi;

fn round_trip() -> bool {
    let callback: ffi::Func[(i64) -> i64] = ffi::func_from_closure(|x| x + 1);
    let Some(call) = ffi::func_to_closure(callback) else {
        return false
    };
    call(41) == 42
}
```

转换生成保留原函数及环境的类型化 Go 方法值。GoML 闭包仍经过闭包提升。空元组和扁平元组结果在 GoML 存储与 Go 的零个/多个结果间显式转换。原始错误和其他受支持原始值保持不变。回调文本值必须使用 `ffi::String` 或 `ffi::Rune`；普通 GoML `string`/`char`（包括数组、切片、通道或原始映射中的）会在适配器验证时被拒绝。请在回调内部做带检查的 Unicode 转换，并显式选择失败结果形态。

适配器不添加同步、全局注册表、反射或隐式 panic 恢复。panic 沿同一 goroutine 传播，包括穿过两个转换方向，Go 和 GoML 的 defer 在栈展开时运行。panic 不会自动成为错误结果。恢复需要显式 `std::panic::catch` 边界，或调用方拥有的、在该 goroutine 中用 defer/recover 的 Go 适配器；二者都不能恢复其他 goroutine 的 panic 或拦截进程退出。`panic::resume` 用内部运行时包装器保留捕获 panic 的诊断信息；Go 调用方在 resume 后 recover 看到的是包装器，而非原载荷动态类型。Go 可在注册回调的 GoML 调用返回后继续保留转换后的回调，在另一 goroutine 调用，或递归重入 GoML。捕获状态通过 Go 函数值保持可达。并发修改需要普通同步；转换不会让 `Ref` 线程安全。

可运行的[回调订阅夹具](../goml/testdata/ffi/callbacks/README.zh-CN.md) 覆盖显式注销、受控异步分派、并发原子更新、递归重入和注册压力。注销会移除其回调引用；已选中的调用仍可能执行，因此调用方在释放捕获资源前等待相应批次。GC 可达性和注册生命周期是不同概念。夹具及其生成的 Go 程序使用竞态检测器测试。仍不支持 GoLibrary 函数值导出。

可执行文件、测试和库模块链接在生成的 Go 字符串常量 `_goml_source_origins` 中嵌入带版本 JSON 索引，将保留的函数及回调适配器关联到外层 GoML 声明。共享适配器可有多个源码条目。外部调用包装器还标识实际使用的 extern 声明，包括 `go_method`；同一原生目标的未使用别名被排除。索引记录规范符号、源包、包相对源文件和声明名称位置，包括显式定义的 impl 方法；它不是表达式级栈追踪行号表。

索引具有 `version: 1`、`scope: "enclosing-declaration"` 和 `entries` 数组。条目包含 `generated_symbol`、`source_symbol`、`source_package`、`source_file`、`start`、`end`、`line` 和 `column`；回调方法用 `ReceiverType.Apply` 作为生成符号。解析 JSON 前先解码 Go 字符串字面量。

<a id="raw-go-strings"></a>

### 原始 Go 字符串

`std::ffi::String` 表示任意 Go 字符串字节，不同于 GoML Unicode `string`。跨外部调用、包接口、泛型函数和捕获闭包传递时字节不变。不支持隐式转换或直接构造。

| 显式适配器 | 结果与策略 |
| --- | --- |
| `ffi::string_from_bytes(bytes::Bytes)` | `ffi::String`；复制字节，包括无效 UTF-8 和 NUL |
| `ffi::string_from_text(string)` | `ffi::String`；保留已验证文本 |
| `ffi::string_bytes(ffi::String)` | `bytes::Bytes`；返回独立可变字节副本 |
| `ffi::string_text(ffi::String)` | `Result[string, utf8::Utf8Error]`；验证 UTF-8，不替换无效字节 |

```goml
use std::ffi;
use std::bytes;

fn checked_text() -> bool {
    let raw = ffi::string_from_bytes(bytes::Bytes::from_vec(Vec::from_array([255, 0])));
    match ffi::string_text(raw) {
        Ok(_) => false,
        Err(error) => error.valid_up_to() == 0,
    }
}
```

用 GoML `string` 声明的旧外部绑定保留现有受信文本契约。Go 生产者可能返回任意字节时，使用 `ffi::String`。GoLibrary 导出可用 `ffi::String` 选择原始字节策略。普通 GoML `string` 和 `char` 导出仍不受支持；请在边界使用显式带检查适配器。

<a id="raw-go-runes"></a>

### 原始 Go rune

`ffi::Rune` 是 `i32` 的透明别名，与 Go `rune` 一致。原始值可以为负、代理码点或超出 Unicode 标量范围。`ffi::rune_from_char(char) -> ffi::Rune` 保留有效字符。`ffi::rune_char(ffi::Rune) -> Result[char, ffi::BoundaryError]` 检查标量范围，无效输入返回 `InvalidRune(original_value)`，从不替换成替换字符。这些适配器使用普通类型别名和函数语法。

<a id="raw-go-slices"></a>

### 原始 Go 切片

`ffi::RawSlice[T]` 是 `MutSlice[T]` 的别名，直接表示 Go `[]T` 头。头赋值复制长度、容量和底层存储引用，不复制元素。现有带检查索引和子切片方法仍可用。

| 适配器或操作 | 策略 |
| --- | --- |
| `ffi::slice_nil::[T]()` / `ffi::slice_is_nil(value)` | 保留并区分 nil 与非 nil 空切片 |
| `ffi::slice_capacity(value)` | 返回原始头的容量 |
| `ffi::slice_append(value, item)` | 返回 Go append 的新头；保留它才能观察新长度 |
| `ffi::slice_shared(vector)` | 共享可变元素存储，使用 Vec 头的快照 |
| `ffi::slice_from_vec_copy(vector)` / `ffi::slice_copy(value)` | 将元素复制到独立存储 |
| `ffi::slice_to_vec_copy(value)` | 将元素复制到新 Vec |
| `ffi::slice_trusted_readonly(value)` | 在绑定作者承诺下，通过只读 `Slice[T]` 视图共享存储 |

复制为浅复制：复制指针切片会隔离元素槽，但所指对象仍共享。空输入的复制适配器生成非 nil 空分配，不保留 nil 身份或剩余容量。追加可能复用底层存储，也可能分配新底层数组；Go append 和 Vec append 都不更新先前复制的头。Go 写入通过共享底层存储的别名可见。所有适配器都不添加同步。

`trusted_readonly` 是作者对存储适合只读使用的显式承诺，包括协调其他写入者；Go 签名无法证明此承诺。转换不冻结原切片，也不递归冻结引用对象。需要独立元素存储时先复制。Vec 和 HashMap 仍无隐式外部转换。

```goml
use std::ffi;

fn append_header() -> bool {
    let original: ffi::RawSlice[i64] = ffi::slice_nil();
    let updated = ffi::slice_append(original, 7);
    ffi::slice_is_nil(original) && updated.len() == 1
}
```

仍不支持 GoLibrary 切片导出。

<a id="raw-go-maps"></a>

### 原始 Go 映射

`ffi::RawMap[K, V]` 直接表示 Go `map[K]V`，不同于 GoML `HashMap[K, V]`。在赋值、外部调用、包接口、泛型和捕获闭包间保留 nil 和共享存储。对别名的 Go 操作观察到相同条目。

`ffi::map_nil::[K, V]()` 产生 nil `ffi::RawMap[K, V]`。`ffi::map_is_nil(value) -> bool` 区分 nil 和已分配空映射。包含 nil 的 `Option::Some` 仍不同于 `None`。原始类型没有公开结构体构造器，也不能与 `HashMap` 隐式互转；二者的存储、哈希和相等契约不同。原始映射语法使用普通限定泛型别名和函数。

| 操作 | 语义 |
| --- | --- |
| `ffi::map_new::[K, V]()` | 分配非 nil 空原始映射 |
| `ffi::map_len(value)` | 条目数；nil 长度为零 |
| `ffi::map_get(value, key)` | `Option[V]`；区分缺失与存在的零值或 nil 值 |
| `ffi::map_set(value, key, item)` | 使用 Go 相等规则插入或替换；写入 nil 映射保留 Go 的 panic |
| `ffi::map_delete(value, key)` | 存在则删除；缺失键和 nil 映射不做操作 |
| `ffi::map_copy(value)` | 独立条目存储，值为浅复制；保留 nil |

赋值和传参共享条目存储。需要独立条目时用 `map_copy`。副本中的指针值仍引用原对象。复制直接遍历键值对，因此即使 NaN 不能通过相等查找，以 NaN 为键的条目仍保留值。这些操作均不冻结存储、不引入同步，也不转换 HashMap 哈希语义。

```goml
use std::ffi;

#[go_ffi("example.com/myapp/goshim", "Entries")]
extern fn entries() -> ffi::RawMap[i64, i64];

fn available() -> bool {
    !ffi::map_is_nil(entries())
}
```

最终 Go 编译器实施 Go 映射键可比较规则。GoML 检查拒绝已知不可比较键，包括切片、原始映射、包含它们的数组或元组，以及含此类字段的本地结构体；推断出的泛型应用在推断后检查。重复的本地泛型包装器在每个具体实例检查：若 `Key[T]` 直接存储 `T`，则 `Key[Key[RawSlice[i64]]]` 被拒绝；仅存储指向 `T` 指针的包装器仍可比较。递归结构体布局及超过 64 层嵌套结构体实例的结构检查产生可恢复诊断。指向不可比较类型的指针仍是合法键。公开 GoML 结构体接口保留私有字段的可比较性摘要，下游键检查可遵循限制而不暴露字段名或类型。摘要可依赖泛型实参；未知隐藏结构不能证明可比较。产物解码器也拒绝直接不可比较的键形态。外部具名键是不透明的；最终 Go 构建检查实际字段和具体类型实参的可比较性。GoML 分析不导入原生字段结构或约束。模块检查从泛型函数体推断内部可比较性要求，包括经函数值和嵌套闭包到达的要求。包检查跨文件传播，直到签名稳定，设有可恢复的 256 轮限制。导出函数接口保留推断要求，因此下游模块检查可在 Go 构建前拒绝无效键。私有结构键包装器被归约为所需类型参数或关联类型投影。这些谓词没有用户编写的约束语法，也不引入 let 泛化。仅源码依赖分析会在检查调用者前检查依赖体并保留其推断要求，包括未保存源码覆盖。结构体和枚举接口也保留字段或变体载荷中映射的键要求，包括私有存储和别名。这些要求在名义声明间传播，设有可恢复的 256 轮限制，并在具体类型应用时替换。递归类型应用通过声明要求检查，不反复展开布局。默认 trait 方法单独保留推断映射键要求，与方法声明约束分开。使用默认体的实现必须在替换接收者和关联类型后满足这些要求。显式方法覆盖不继承未使用默认体的要求。仍不支持 GoLibrary 映射导出。原始映射不添加同步或自动深复制。

<a id="declaring-go-bindings"></a>

### 声明 Go 绑定

编写显式 `go_ffi`、`go_type` 和 `go_method` 声明，API 需要适配时使用 Go 转发函数。声明模型见 [Go 互操作](ffi/bind-go.zh-CN.md)，显式指针和方法绑定见[文件读取示例](../goml/testdata/ffi/file-read/README.zh-CN.md)。

<a id="c-compatible-strings"></a>

### C 兼容字符串

`std::c` 通过普通 GoML 类型和调用提供 `CString`、带检查的 C 字符串字面量及边界错误。API 和编译期验证见 [C 兼容字符串](stdlib.zh-CN.md#c-compatible-strings)。这些 API 不添加指针算术、原生布局或 `extern "C"` 语法。

<a id="exporting-a-go-library"></a>

### 导出 Go 库

前端识别 `#[go_export("Add")] pub fn add(x: i64, y: i64) -> i64 { x + y }`。检查声明必须为普通顶层公开函数，不带泛型、comptime/derive 能力或测试属性。Go 名以 ASCII 大写字母开头，仅含 ASCII 字母、数字或下划线；所选 GoML 包内名称必须唯一。重复属性及方法、extern、类型、字段和其他非函数项上的属性会被拒绝。

导出参数目前允许 bool、整数和浮点基本类型、`ffi::String` 及这些值的定长数组。返回额外允许空元组或由允许的非 `()` 值组成的扁平元组。透明别名规范化为底层类型。普通 GoML 字符串和字符、切片、通道、用户类型、函数值、元组参数及嵌套元组返回不在此导出边界内。

文本边界策略在 GoML 类型和适配器体中显式体现：

| 导出的 GoML 类型 | Go 签名 | 边界策略 |
| --- | --- | --- |
| `ffi::String` | `string` | 保留任意字节，包括无效 UTF-8 和 NUL；无隐式 Unicode 转换 |
| `ffi::Rune`（`i32` 的别名） | `int32` / `rune` | 保留原始整数；无隐式标量验证 |
| 普通 `string` / `char` | 拒绝 | 选择原始类型并编写显式带检查适配器 |

在导出函数内用 `ffi::string_text` 或 `ffi::rune_char` 验证值后再视为 Unicode。失败须通过适配器支持的返回签名或有意选择的边界失败策略表达。以下适配器对无效 UTF-8 返回原字节和 `false`，否则返回已验证文本和 `true`，从不静默替换为替换字符：

```goml
use std::ffi;

#[go_export("CheckedText")]
pub fn checked_text(value: ffi::String) -> (ffi::String, bool) {
    match ffi::string_text(value) {
        Ok(text) => (ffi::string_from_text(text), true),
        Err(_) => (value, false),
    }
}
```

返回 `ffi::String` 的原始回显可不经解码保留字节。rune 适配器同样可返回 `(ffi::Rune, bool)`。定长数组和扁平元组返回保留这些元素策略。`Result` 本身不是 Go 导出 ABI 类型；选择并记录失败形态由适配器作者负责。

接口和 Core 产物保留所选 Go 导出名、稳定包/函数身份、解析签名及相对源码来源。导出名和签名参与接口语义哈希，源码位置不参与。内部库链接器在死代码消除和单态化前，仅选取目标包声明的导出作为函数根；可达私有及跨包辅助函数被保留，无需 `main`。内部 Go 库后端为标量和定长数组参数/结果生成可调用包装器，将空元组映射为无 Go 结果，将扁平元组映射为多个 Go 结果。库包装器使用包级 OnceCell 初始化并保留 panic。生成的内部包声明使用私有 Go 标识符；只有显式选定的导出名成为公开 API。单独 `pub` 不请求 Go 导出。

使用 `goml export-go` 从当前 GoML 模块中的一个包生成库：

```text
goml export-go [PACKAGE_DIR] --import-path example.com/host/gen/calclib --out ./gen/calclib
```

默认源为 GoML 模块根。显式源路径和输出路径相对于当前工作目录。源必须是当前模块中的包，且不能声明 `package main`。仅其 `go_export` 声明成为 Go API；依赖导出仅在可达时保留。输出目录基本名是 Go 包名，必须是非 main 的 ASCII Go 标识符。

GoML 模块根必须已有 `go.mod`。输出必须位于该 Go 模块内、嵌套 Go 模块外，提供的导入路径必须等于 Go 模块路径加规范相对输出路径。拒绝父目录穿越，检查边界前解析符号链接。命令从不创建或修复 `go.mod` / `go.sum`，也不下载缺失依赖。

输出包含 `goml_generated.go` 和 `goml_exports.json`。清单记录公开签名、生成版本、Go 构建先决条件和所有权摘要。先决条件包括 Go 工具链、目标、cgo 设置、标志和构建标签。保留手写 Go 文件，并与候选生成源码一起检查。发布前通过 Go overlay 检查新输出；无效 Go、重复符号和导入循环会失败，不发布候选。已修改或不归生成器所有的文件不会被覆盖。同一输出的并发生成会被拒绝。发布会回滚已报告的文件替换错误；断电或强制终止进程时，两文件发布不具有崩溃原子性。

发布前，Go 编译检查候选生成包和每个已生成外部调用。GoML 分析仍信任外部声明；未使用声明不必出现在候选中。支持 `--compiler`、`--target-dir`、`--jobs` 和 `--dry-run`。生成的纯 Go 库仅需 Go 工具链即可构建。完整双向示例（含 Go 可执行程序和 Go 测试）见[互操作夹具](../tools/release/testdata/go-export)。

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

`goml fmt` 使用与构建和测试相同的包图，格式化当前模块的生产文件、内部测试、黑盒测试包、示例及其测试。可从任何嵌套包目录运行。包发现排除 `testdata`、构建输出、隐藏目录、嵌套模块和外部依赖。`goml fmt --check` 仅检查格式，任何源文件需要改变时以失败状态退出。

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

### LSP 与编辑器

项目配置和模块依赖解析由 `goml` 驱动负责。LSP 通过 `goml __project-info --source <PATH>` 请求带版本项目描述，并连同打开文档内容传给编译器查询引擎。描述提供规范模块身份、源码目录、分开的生产/开发依赖边及配置监视路径。此操作读取本地文件和缓存依赖，不编译、获取或写入构建产物。依赖失败保持为诊断；可用根元数据被保留以供本地包补全。

项目描述客户端验证协议版本 1，在所监视配置内容和规范路径不变时缓存成功描述。普通源码编辑复用项目解析，并仍参与编译器源码分析指纹。创建工作区、依赖清单变化、替换和注册表索引变化触发重新解析。缺失依赖和无效描述在编辑后重试，不缓存为成功项目。编译器仍负责源码导入解析、可见性和类型检查。

`GOML_PROJECT_DRIVER` 显式选择驱动。否则客户端使用同级 `goml`、`GOML_HOME/bin/goml` 或 `PATH`；仓库测试可提供 `GOML_TEST_GOML`。驱动和语言服务器必须支持相同项目描述协议。扩展捆绑两种工具，并提供 `goml.driverPath` 用于自定义安装。现有编译器构建命令和驱动协议不变，因此固定 stage0 可继续构建工具。

LSP 在各自包上下文中分析生产文件、内部和黑盒测试，因此诊断、补全、悬停及跳转定义遵循与编译相同的可见性规则。普通和参数化测试用例上显示测试 CodeLens。VS Code 扩展保存文件后，以所选测试显示名和种类调用模块级 `goml test`。

类型内嵌提示使用不带包限定的短名，类型长度预算为 25 字符，不含前导 `: `。嵌套类型和长名用 `…` 缩写；悬浮提示保留完整推断类型和包路径。

包分析在查询上下文内缓存解析/降级源码和成功检查的包。源码内容、包成员、规范模块身份和传递依赖指纹决定复用。编辑包会重查该包及依赖它的包，独立包仍可复用；依赖中的函数体编辑也保守地使依赖方失效。打开文档覆盖参与相同指纹。外部 Go 类型在缓存和新分析中都保持不透明；原生验证属于最终 Go 构建。监视的源码和清单变化使文档分析缓存失效；VS Code 扩展还监视 `goml.work`、Go 源码、`go.mod` 和 `go.sum`。

自定义 `goml/expandedDerive` 请求返回所请求文档完成内置和可编程 derive 后的格式化 AST。它使用与 `goml check` 相同的包别名、显式 derive 导入、歧义检查、依赖接口、CTIR 验证器和资源限制。VS Code 命令 `GoML: Show Expanded Derive` 在源文件旁打开结果。

`gomllsp` 通过 `textDocument/formatting` 支持整文档格式化。格式化使用最新未保存文本及[格式化](formatting.zh-CN.md)中的固定规则。无效文档保持不变。

CLI 和 LSP 诊断共享包含可选稳定代码、备注、帮助和建议编辑的模型。CLI 显示 `error[typer/try_error_mismatch]` 或 `warning[typer/unused_result]`；LSP 在 `Diagnostic.code` 暴露相同标识符，并在 `relatedInformation` 暴露次级位置。备注和帮助出现在显示消息中，也单独放在诊断 `data`，连同修复和文档版本。修复包含标题、一项或多项范围编辑，以及 `machine-applicable` 或 `requires-review` 适用性。代码操作拒绝旧文档版本的修复。目前为未使用 Result 警告及上述三种 `?` 诊断分配代码；其他诊断可能省略。

<a id="built-in-prelude"></a>

## 内置 prelude

工具链库分为三层。`builtin` 是隐藏的编译器和运行时契约：拥有基本及内置类型身份、运行时钩子、语言项、内置实现和 derive 处理函数。`prelude` 是依赖 `builtin` 的独立编译包，定义自动加入普通源码作用域的公开名称。`std` 包含普通标准库包，仅通过显式导入可用。

用户和第三方包不能显式导入 `builtin` 或 `prelude`；尝试会产生诊断，不会解析注册表包。这隐藏原始运行时辅助函数和 derive 内部机制，同时保留重新导出类型/trait 的稳定身份。`Option`、`Result`、`Vec`、`ToString`、`print`、`println` 和 `range` 等名称通过 prelude 无需 `use` 即可使用。

编译器识别的协议在 `builtin` 中用保留的 `#[lang(...)]` 元数据声明。元数据标识所属类型或 trait 及枚举变体、关联类型和方法等协议成员，使编译器行为不依赖公开拼写。用户、第三方、prelude 和标准库源码中的 `#[lang(...)]` 会被拒绝。

由运行时支持的集合类型在重新导出时保留 builtin 身份。其他包的限定类型即使短名相同也仍不同：`use std::collections::imm;` 后，`imm::HashMap[K, V]` 是持久化标准库映射，而 prelude 的 `HashMap[K, V]` 是可变运行时映射。这使用普通包限定类型路径，不增加语法。

以下名称无需 `use` 即可使用。

<a id="option-and-result"></a>

### `Option` 与 `Result`

```goml
enum Option[T] {
    None,
    Some(T),
}

enum Result[T, E] {
    Ok(T),
    Err(E),
}
```

预期类型能确定枚举时，构造和模式中优先使用 `Some`、`None`、`Ok` 和 `Err`。类型上下文不足或其他声明遮蔽变体时，仍可用 `Option::Some`、`Result::Ok` 等限定名。

`Option[T]` 提供 `is_some`、`is_none`、`unwrap_or` 和 `unwrap_or_else`，以及改变类型的泛型方法：

- `value.map(map_fn) -> Option[U]`
- `value.and_then(next) -> Option[U]`
- `value.ok_or(error) -> Result[T, E]`

`Result[T, E]` 提供 `is_ok`、`is_err`、`unwrap_or` 和 `unwrap_or_else`。改变类型的泛型方法为：

- `value.map(map_fn) -> Result[U, E]`
- `value.map_err(map_fn) -> Result[T, F]`
- `value.and_then(next) -> Result[U, E]`

### `Default`

prelude 的 `Default` 实现为空元组使用 `()`，为 `bool` 使用 `false`，为数值类型使用零，为 `string` 使用空字符串。`Vec` 和 `HashMap` 默认为空容器，标准集合 `HashSet`、`IndexMap`、`Stack`、`Deque`、`BitSet`、`Arena`、`IndexVec` 和 `Interner` 也通过空构造器默认初始化。`Option[T]` 默认为 `None`，不要求 `T: Default`；元组和定长数组逐元素默认初始化。`Result[T, E]` 和 `Ref[T]` 有意不提供全局默认实现。

<a id="output-and-string-conversion"></a>

### 输出与字符串转换

- `print[T: ToString](value: T) -> ()`
- `println[T: ToString](value: T) -> ()`
- `value.to_string() -> string`，适用于实现 `ToString` 的值
- `value.debug() -> string`，适用于实现 `Debug` 的值
- `string.len() -> isize` 和 `string.byte_len() -> isize`，都返回 UTF-8 字节长度
- `string.get(index: isize) -> char`，在 UTF-8 字节边界解码字符
- `string.byte_get(index: isize) -> u8`
- `string.byte_slice(start: isize, end: isize) -> string`
- `string.is_char_boundary(index: isize) -> bool`
- `string.decode_at(index: isize) -> Option[(char, isize)]`
- `string.to_bytes() -> Vec[u8]`
- `string.chars() -> FnIterator[char]`
- `string.char_indices() -> FnIterator[(isize, char)]`
- `string.starts_with(prefix: string) -> bool`
- `string.ends_with(suffix: string) -> bool`
- `string.contains(expected: string) -> bool`
- `string.starts_with_at(start: isize, prefix: string) -> bool`，在字节偏移处检查
- `string.find(expected: string) -> Option[isize]`、`string.rfind(expected: string) -> Option[isize]` 和 `string.find_bytes(expected: string) -> Option[isize]`，返回字节偏移
- `string.char_count() -> isize`，计数 Unicode 标量值
- `string.slice_chars(start: isize, end: isize) -> Option[string]`，使用标量值索引
- `string.trim() -> string`、`string.trim_start() -> string` 和 `string.trim_end() -> string`，修剪 ASCII 空白
- `string.split(separator: string) -> Vec[string]`、`string.split_once(separator: string) -> Option[(string, string)]` 和 `string.lines() -> Vec[string]`
- `string.replace(expected: string, replacement: string) -> string`
- `string.repeat(count: isize) -> string`
- `string.is_ascii() -> bool`
- `string.eq_ignore_ascii_case(other: string) -> bool`
- `string.to_ascii_lowercase() -> string` 和 `string.to_ascii_uppercase() -> string`

基本标量类型都实现 `ToString` 和 `Debug`。字符串拼接使用 `+`。

与键相关的 trait 包括带方法的 `PartialEq::eq(self, other: Self) -> bool`、标记 `Eq: PartialEq` 和 `Hash::hash(self) -> u64`。用户类型可手写实现或派生相应 trait。

### `Ref[T]`

```goml
let cell = Ref::new(1);
let before = cell.get();
cell.set(before + 1);
```

API:

- `Ref::new(value) -> Ref[T]`
- `reference.get() -> T`
- `reference.set(value) -> ()`
- `ptr_eq(a, b) -> bool`，比较引用身份

`Ref[T]` 的内置 `PartialEq`、`Eq` 和 `Hash` 实现使用引用身份，不要求 `T` 实现这些 trait。因此修改引用值不会改变相等性或哈希。即使 `T` 为 `()`、空结构体或其他零大小类型，独立分配也有不同身份。同一引用的副本保留身份。`Ref[T]` 不实现 `Default`，因为隐式分配和递归默认构造容易造成意外。

<a id="fixed-arrays"></a>

### 定长数组

```goml
let mut values: [i32; 3] = [1, 2, 3];
let first = values[0];
values[1] = 20;
```

索引类型为 `isize`。也可调用底层 `array_get` 和 `array_set`；日常代码优先使用索引语法。

### `Vec[T]`

```goml
let values = Vec::from_array([1, 2, 3]);
let first = values[0];
```

`Vec::from_array([values...])` 从定长数组创建向量。数组项从左到右求值一次。向量获得新的外层底层存储，因此之后替换数组元素不影响向量；由于浅复制，引用元素仍共享。空向量使用 `Vec::new()`。直接非空数组调用联合分配向量头和底层数组。传入数组变量保留相同浅复制语义。该关联函数是普通一等值，例如 `let make: ([i32; 3]) -> Vec[i32] = Vec::from_array;`。

常用方法：

- `Vec::new() -> Vec[T]`
- `Vec::from_array(values: [T; N]) -> Vec[T]`
- `Vec::with_capacity(capacity) -> Vec[T]`
- `push(value) -> ()`
- `pushed(value) -> Vec[T]`
- `copy() -> Vec[T]`
- `freeze() -> FrozenVec[T]`
- `get(index) -> T`
- `set(index, value) -> ()`
- `len() -> isize`
- `capacity() -> isize`
- `is_empty() -> bool`
- `contains(value) -> bool`
- `position(predicate) -> Option[isize]`
- `sort_by(compare) -> ()`
- `stable_sort_by(compare) -> ()`
- `sort_by_ordering(compare) -> ()`
- `stable_sort_by_ordering(compare) -> ()`
- `binary_search_by(compare) -> Option[isize]`
- `binary_search_by_ordering(compare) -> Option[isize]`
- `min_by(compare) -> Option[T]`
- `max_by(compare) -> Option[T]`
- `min_by_ordering(compare) -> Option[T]`
- `max_by_ordering(compare) -> Option[T]`
- `dedup_by(equal) -> ()`
- `dedup() -> ()`
- `join(separator) -> string`
- `reserve(additional) -> ()`
- `truncate(len) -> ()`
- `clear() -> ()`
- `last() -> Option[T]`
- `pop() -> Option[T]`
- `swap(left, right) -> ()`
- `swap_remove(index) -> T`
- `insert(index, value) -> ()`
- `remove(index) -> T`
- `reverse() -> ()`
- `extend(other) -> ()`
- `slice(start, end) -> Slice[T]`
- `iter() -> FnIterator[T]`

`vec[index]` 读取元素，`vec[index] = value;` 修改元素。

`copy()` 浅复制外层底层存储。`freeze()` 也复制底层存储并返回 `FrozenVec[T]`，因此之后通过原 `Vec` 别名进行结构或元素修改不会影响冻结值。

`contains(value)` 要求元素类型实现 `PartialEq`，返回是否存在等于 `value` 的元素。

排序和选择方法接受返回负数/零/正数的整数比较器，`_ordering` 变体则接受 `Ordering` 比较器。`dedup()` 要求 `PartialEq`；`position` 接受谓词；`binary_search_*` 要求向量已排序，并返回首个匹配索引。`join(separator)` 要求 `ToString`，拼接元素的字符串形式。

### `FrozenVec[T]`

`FrozenVec[T]` 是动态大小的只读向量，暴露 `get`、`len`、`is_empty`、`contains`、`iter` 和 `to_vec`，但没有修改方法。`to_vec()` 返回新的浅复制可变底层存储。嵌套容器的冻结是浅层的：两层都须只读时使用 `FrozenVec[FrozenVec[T]]`。

<a id="slicet-and-mutslicet"></a>

### `Slice[T]` 与 `MutSlice[T]`

`Slice[T]` 是 `Vec[T]` 存储上的有界只读视图：

```goml
let view: Slice[i32] = values.slice(1, 3);
let item = view.get(0);
```

常用只读方法有 `get`、`get_checked`、`len`、`contains`、`sub`、`sub_checked`、`to_vec` 和 `iter`。`contains(value)` 要求元素类型实现 `PartialEq`。`to_vec()` 在新 `Vec[T]` 中创建浅副本；可以读取 `view[index]`，但 `view[index] = value;` 会被拒绝。

`MutSlice[T]` 是有界可变零拷贝视图。`Vec::slice_mut(start, end)` 和 `Vec::as_mut_slice()` 创建它；`as_slice()` 无需复制即可转为只读视图：

```goml
let values = Vec::from_array([1, 2, 3, 4]);
let output = values.slice_mut(1, 4);
output[0] = 20;
let written = output.copy_from(1, Vec::from_array([7, 8]).as_slice());
```

`get_checked`、`set_checked`、`sub_checked`、`Vec::slice_checked` 和 `Vec::slice_mut_checked` 在索引前报告无效范围。`copy_from(offset, source)` 写入前验证完整目标范围，并使用重叠安全复制语义；`copy_within` 是对应的同视图操作。还支持 `fill`、`to_vec` 和迭代。`Slice` 和 `MutSlice` 保持当前底层数组存活。之后 `Vec` 增长可能将向量移到新底层数组，因此旧视图继续引用创建时捕获的存储。

### `HashMap[K, V]`

键类型必须同时实现 `Hash` 和 `Eq`；桶相等比较通过父 trait `PartialEq` 分派：

```goml
let counts = HashMap::from_array([("a", 1), ("b", 2)]);
let value: Option[i32] = counts.get("a");
```

`HashMap::from_array([(key, value), ...])` 从条目数组创建映射。先从左到右求值整个数组，再按数组顺序插入键值对。后面的重复键覆盖先前条目。空映射使用 `HashMap::new()`。直接数组字面量调用根据条目数预分配构建存储。数组长度由函数类型固定时，该关联函数是普通一等值。

常用方法：`new`、`from_array`、`get`、`set`、`insert`、`get_or_insert_with`、`update`、`remove`、`remove_value`、`entry`、`len`、`contains` 和 `entries`。`insert` 返回旧值，`update` 仅键存在时应用 `(V) -> V` 闭包，`remove_value` 返回移除值。`entries()` 返回快照 `Vec[(K, V)]`。

`entry(key)` 返回 `HashMapEntry[K, V]`，提供 `get`、`or_insert`、`or_insert_with`、`insert`、`and_modify`、`update`、`remove` 和 `remove_entry`。融合的 entry 操作只对键哈希和搜索一次；传给 `or_insert_with` 的创建闭包仅在键缺失时运行：

```goml
let counts: HashMap[string, isize] = HashMap::new();
let count = counts.entry("jobs").or_insert_with(|| 0);
let _ = counts.entry("jobs").update(|value| value + 1);
let removed = counts.entry("jobs").remove_entry();
```

索引读取返回 `Option[V]`，索引赋值写入 `V`：

```goml
let value = counts["a"];
counts["b"] = 2;
```

使用 `std::collections` 克隆、复制、比较和迭代映射见[映射算法](stdlib.zh-CN.md#map-algorithms)。

<a id="channels"></a>

### 通道

`Channel[T]` 使用 Go 通道后端，支持有缓冲和无缓冲通信：

```goml
let channel = Channel::[string]::new(0);
go || channel.send("ready");
let value: Option[string] = channel.recv();
channel.close();
```

常用方法：

- `Channel::[T]::new(capacity: isize) -> Channel[T]`
- `send(value: T) -> ()`
- `recv() -> Option[T]`，关闭且排空时返回 `Option::None`
- `close() -> ()`
- `sender() -> Sender[T]`
- `receiver() -> Receiver[T]`
- `split() -> (Sender[T], Receiver[T])`

`Sender[T]` 提供 `send` 和 `close`；`Receiver[T]` 提供 `recv`。方向在静态检查中强制执行，并生成为 Go 的 `chan<- T` 和 `<-chan T`。转换双向通道需要显式端点方法。`Channel[T]` 保留原有发送、接收和关闭 API。

容量 `0` 创建无缓冲通道。发送到无缓冲通道通常应由另一个 `go` 闭包并发接收，反之亦然。

goroutine 需要等待多个通道、在发送和接收间选择、通过 `Option[T]` 观察关闭，或通过 `default` 做非阻塞操作时，使用 `select`。通道操作数和发送值在选择前准备一次。

使用 `std::channel` 进行运行时数量的选择见[动态通道选择](stdlib.zh-CN.md#dynamic-channel-selection)。

<a id="iterator"></a>

### 迭代器

内置协议如下：

```goml
trait Iterator {
    type Item;
    fn next(self) -> Option[Self::Item];
}

trait IntoIterator {
    type Item;
    type IntoIter: Iterator;
    fn into_iter(self) -> Self::IntoIter;
}
```

协议入口为 `FnIterator::from_fn(next_fn)`、`iterator.next()` 或 `Iterator::next(iterator)`。`range(start, end)` 和 `start..end` 产生递增半开 `FnIterator[isize]`；`start >= end` 时为空。

`std::iter` 中的构造、适配器和消费操作见[迭代器适配器](stdlib.zh-CN.md#iterator-adapters)。

迭代器只能单次遍历。定长数组使用原生索引 `for` 降级。`Vec[T]`、`Slice[T]` 和 `MutSlice[T]` 可直接用于 `for`，实现 `Iterator` 的值也通过恒等 `IntoIterator` 直接迭代。

<a id="comparison-of-common-writing-errors"></a>

## 常见写法错误对照

库 API 和常见任务见[标准库参考](stdlib.zh-CN.md#common-tasks)。

| 避免的写法 | GoML 写法 |
| --- | --- |
| API 文档块内使用 `////` 或空源码行 | 使用 `///`；段落之间使用仅含 `///` 的行 |
| `Vec<isize>` | `Vec[isize]` |
| `fn id<T>(x: T) -> T` | `fn id[T](x: T) -> T` |
| `fn count[N](values: [i32; N])` | 将值参数声明为 `fn count[const N: usize](values: [i32; N]) -> usize` |
| `Buffer[i32, N + 1]` | `Buffer[i32, {N + 1}]` |
| 用 `serde(rename)` 修改 `.to_string()` | 在变体上使用 `#[to_string("text")]`，并配合 `#[derive(ToString)]` |
| 用 `Type::LIMIT()` 访问关联常量 | `Type::LIMIT` |
| `id::<i32>(1)` | `id::[i32](1)` |
| 普通函数调用 `id[i32](1)` | `id::[i32](1)`，或依赖参数/结果类型推断 |
| `1i32`, `1u64`, `1.0f32` | 使用预期类型，如 `let value: u64 = 1;` |
| `b"é"` 中的非 ASCII 源文本 | 使用 `string`，或显式写编码字节，如 `b"\xC3\xA9"` |
| 控制流头部值使用裸块 | 加括号，如 `if ({ prepare(); ready() }) { ... }` |
| `let mut x: &T` | 按需使用值 `T` 或 `Ref[T]` |
| 在值位置写 `if cond { value }` | `if cond { value } else { other }` |
| `let Option::Some(x) = value;` | `let Some(x) = value else { return };`、`if let` 或 `match` |
| `let Point { x } = point;` | `let Point { x, .. } = point;` |
| 不穷尽的 `match` | 覆盖所有可能变体或添加 `_` 分支 |
| 用 `operation();` 丢弃 `Result` | 处理结果，用兼容的 `?` 传播，或用 `let _ = operation();` 显式忽略 |
| `x++`, `x--` | `x += 1;`, `x -= 1;` |
| 通过不可变结构体绑定赋值 | 用 `let mut` 声明绑定，或用 `Point { field: value, ..point }` 创建新值 |
| `var x = 1`, `x := 1` | `let x = 1;` |
| 有条件的循环 | `while condition { ... }` |
| `for i := 0; ...` | `while` 或 `for i in start..end` |
| `switch` | `match` |
| `null`, `nil` | 有预期 `Option` 类型的 `None` 或 `Option::None`；原始 Go 指针用 `ffi::null()`；Go 错误用 `ffi::nil_error()` |
| `throw`、异常 | 预期错误使用 `Result` 和 `?`；显式运行时 panic 边界使用 `std::panic::catch` |
| 隐式转换为 `dyn A + B` | 写 `value as dyn A + B`；值必须实现所有约束 |
| 对 `trait Consumer[T]` 使用 `dyn Consumer` | 提供 trait 实参，如 `dyn Consumer[string]` |
| `dyn TraitWithAssociatedType` | 绑定每个关联类型，如 `dyn Iterator[Item = isize]` |
| `-> impl Trait` 下返回不同具体类型 | 每个分支返回同一具体类型，或显式返回 `dyn Trait` |
| `fn consume(value: impl Trait)` | 使用具名泛型参数，如 `fn consume[T: Trait](value: T)` |
| 对运行时 `count` 使用 `[initializer(); count]` | 重复长度是编译期整数；运行时长度使用 `Vec` 和循环 |
| 期望 `[initializer(); N]` 为每个元素调用一次 | 初始化器只运行一次，长度为零也一样；逐元素工作使用循环 |
| `if let Some(x) = value || ready` | 使用嵌套控制流，或在 let 链的布尔项中使用 `&& (first || second)` |
| `UserId` 必须独立却使用 `type UserId = u64;` | 使用 `struct UserId(u64);` 并显式构造 |
| `use pkg::*` | 用 `use pkg::{A, B};` 显式列出所需公开项 |
| `mod`, `crate::`, `super::` | 使用目录包，当前模块用 `module::path`，依赖用规范路径 |
| 同时开发多个模块 | `goml.work` 列出成员；依赖 `{ path = "../library" }` 和根 `[replace]` 选择本地来源，同时保留声明依赖和规范身份；`--workspace` 或 `-p` 选择模块 |
| 将仅测试辅助模块放进生产依赖 | 使用根 `[dev-dependencies]`；测试和 `examples/<name>/` 可导入 |
| 每个示例创建独立消费方清单 | 使用 `goml run --example <name>` 和 `goml test`；示例共享模块根清单 |
| 函数内部的 `fn helper` | 顶层函数或局部闭包 |
| 离开块后使用其局部绑定 | 从块返回值并在外部绑定结果 |
| `Message::Value::[i32](42)` | 枚举实参属于枚举本身：`Message::[i32]::Value(42)` |
| `identity::[](1)` | 省略空 turbofish，写 `identity(1)` |
| `Message::[bool]::Named { value: true }` | 使用 `let value: Message[bool] = Message::Named { value: true };` |
| 期望函数常量参数改变顶层常量初始化器 | 常量保留声明位置的绑定 |
| Go 外部类型 | `#[go_type("pkg", "Name")] extern type Name[T];` 保留不透明 Go 身份和声明实参；使用显式函数和方法，原始指针用 `std::ffi::Ptr[T]`。不自动提供字段、结构体构造、别名展开或约束发现；最终 Go 编译验证原生使用 |
| Go 接口适配器 | `#[go_interface(RawType, Wrapper, method = "GoMethod")]` 从显式 trait 签名和映射生成适配器；最终 Go 编译检查真实接口。`from_trait` 保留提供的 dyn 对象；保留 nil/带类型 nil 和多个结果 |
| 可被 Go 调用的导出 | 用 `#[go_export("Name")]` 标注受支持公开函数，并用 `goml export-go` 生成 Go 包 |
| 自定义枚举显示文本 | `#[derive(ToString)]` 配合变体上的 `#[to_string("text {field}")]` |
| 生成公开固有方法 | `derive_output_inherent` 和 `derive_output_add_public_method` |
| 生成多个实现和辅助项 | `derive_output_merge`、关联项构建器及私有函数/类型构建器 |
| 在 derive 中构建运行时回调 | `meta_expr_closure` 和 `meta_type_function` |
| 复用编译期计算 | 闭包、间接调用、具体方法、本地动态分派、`Ref` 和 `Vec` 操作 |
| 未标注的用户 `extern fn` | 使用普通 GoML 函数或 `#[go_ffi("import/path", "ExportedSymbol")] extern fn`；GoML 信任声明签名，最终 Go 编译验证生成调用 |
| 从 `comptime` 调用普通函数 | 用 `#[comptime]` 标记受支持自由函数 |
| 在 `comptime` 捕获运行时局部变量 | 向 `#[comptime]` 函数传递字面量或编译期值 |

文档写法 `////` 是普通注释，不是项文档块。`///` 与声明之间的空行会让注释失去附着；块内段落之间使用仅含 `///` 的行。

<a id="informal-grammar-quick-facts"></a>

## 非正式语法速查

开发依赖、具名示例和下游验证影响清单、目录发现及 CLI 命令，不增加 GoML 源语法产生式；示例使用普通包声明和导入。

词法注释约定（下文语法中注释为附属信息）：

```text
LineComment    ::= "//" { any character except newline }
ItemDocLine    ::= "///" { any character except newline }
PackageDocLine ::= "//!" { any character except newline }
```

`ItemDocLine` 排除以 `////` 开头的行，要求声明前的独立注释。`PackageDocLine` 仅在首个源码词法单元前识别。

Base32、变长整数、校验和、UTF-8 标量/流解码、映射/切片辅助操作、逻辑斜杠路径和定宽位操作使用普通导入、trait、结构体、枚举、切片和调用。不增加语法、隐式字节转换或编译器内建操作。映射和切片收集使用现有 `Iterator` 关联类型约束。Shell 模式是 `std::path::slash` 解析的字符串数据，不是 GoML 语法。

`std::c` 使用现有结构体、常量、`#[comptime]`、导入和调用。不增加 C 指针、C 布局或 `extern "C"` 语法。

Linux 系统调用缓冲区、`Pointer` 描述符、`Errno` 值、ABI 编解码器及描述符/进程/内存/IPC 包装器使用下文普通结构体、枚举、数组、切片和调用形式。原生内核布局编码为字节；没有指针转换、原生布局或 `unsafe` 语法。

Panic 边界使用普通函数调用、闭包、`Result` 和现有 `never` 类型。没有 `try/catch`、`throw` 或裸 Go 风格 `recover()` 语法；panic 清理扩展现有词法 `defer` 语句的语义。

`Report::map`、`ResultExt::into_report` 和 `ReportContext::{context, with_context}` 使用现有泛型方法、trait 导入和闭包形式。不增加语法产生式或隐式错误转换。

I/O trait、取消上下文、DNS/TLS 客户端和监听器、标量转换 trait 及 Serde 扩展事件使用现有导入、泛型约束、枚举、方法和调用，不引入新语法。即使类型实参未出现在函数签名中，显式泛型调用也保持 `::[Type]` 形式。

以下 EBNF 仅描述应生成的规范形式；`?` 表示可选，`*` 表示重复，终结符放在引号中。

```text
file          = package_decl? use_decl* item*
package_decl  = "package" lower_ident ";"
use_decl      = "pub"? "use" use_path ("as" ident | "::" "{" use_items "}")? ";"
use_items     = use_item ("," use_item)* ","?
use_item      = ident ("as" ident)?
use_path      = path | "module" "::" path
path          = ident ("::" ident)*

item          = attribute* visibility? function
              | attribute* visibility? type_alias
              | attribute* visibility? constant
              | attribute* visibility? static
              | attribute* visibility? struct_def
              | attribute* visibility? enum_def
              | attribute* visibility? trait_def
              | attribute* impl_def
              | go_ffi_extern
              | attribute* visibility? extern_type
visibility    = "pub"
attribute     = "#[" attribute_body "]"
go_export_attribute = "#[" "go_export" "(" string_literal ")" "]"
comptime_attribute = "#[" "comptime" "]"
comptime_derive_attribute = "#[" "comptime_derive" ("(" ident ")")? "]"
derive_attribute = "#[" "derive" "(" path ("," path)* ")" "]"
to_string_attribute = "#[" "to_string" "(" string_literal ")" "]"
extern_type = "extern" "type" upper_ident generic_params? ";"
go_ffi_attribute = "#[" "go_ffi" "(" string_literal "," string_literal ")" "]"
go_method_attribute = "#[" "go_method" "(" string_literal ")" "]"
go_interface_attribute = "#[" "go_interface" "(" path "," ident ("," ident "=" string_literal)* ","? ")" "]"
go_ffi_extern = (go_ffi_attribute | go_method_attribute) visibility? "extern" "fn" lower_ident
                param_list return_type? ";"

function      = "fn" lower_ident generic_params? param_list return_type? where_clause? block
type_alias    = "type" type_ident type_names? "=" type ";"
type_ident    = upper_ident | simd_type_ident
simd_type_ident = ("i8" | "u8" | "mask8") ("x16" | "x32")
                | ("i16" | "u16" | "mask16") ("x8" | "x16")
                | ("i32" | "u32" | "f32" | "mask32") ("x4" | "x8")
                | ("i64" | "u64" | "f64" | "mask64") ("x2" | "x4")
constant      = "const" ident ":" type "=" expression ";"
static        = "static" ident ":" type "=" expression ";"
method        = visibility? "fn" lower_ident generic_params? param_list return_type? where_clause? block
generic_params = "[" generic_param ("," generic_param)* "]"
generic_param = upper_ident (":" trait_set)? | "const" upper_ident ":" "usize"
param_list    = "(" (parameter ("," parameter)*)? ")"
parameter     = lower_ident ":" type | "self"
return_type   = "->" (type | opaque_return)

struct_def    = "struct" type_ident type_names?
                ("{" struct_fields? "}" | "(" newtype_field ","? ")" ";")
struct_fields = struct_field ("," struct_field)* ","?
struct_field  = visibility? lower_ident ":" type
newtype_field = visibility? type
enum_def      = "enum" upper_ident type_names? "{" variants? "}"
variants      = variant ("," variant)* ","?
variant       = attribute* (upper_ident
              | upper_ident "(" type_list? ")"
              | upper_ident "{" variant_fields? "}")
variant_fields = lower_ident ":" type ("," lower_ident ":" type)* ","?
type_names    = "[" type_parameter ("," type_parameter)* "]"
type_parameter = upper_ident | "const" upper_ident ":" "usize"

trait_def     = "trait" upper_ident generic_params? (":" trait_set)? where_clause?
                "{" trait_member* "}"
trait_member  = "type" upper_ident (":" trait_set)? ";"
              | "const" ident ":" type ("=" constant_expression)? ";"
              | "fn" lower_ident generic_params? param_list return_type? where_clause?
                (";" | block)

impl_def      = "impl" generic_params? trait_ref "for" type where_clause?
                "{" impl_member* "}"
              | "impl" generic_params? type where_clause?
                "{" (attribute* method | visibility? constant)* "}"
impl_member   = "type" upper_ident "=" type ";" | attribute* method | constant

trait_set     = trait_ref ("+" trait_ref)*
trait_ref     = path type_args?
where_clause  = "where" where_predicate ("," where_predicate)* ","?
where_predicate = type ":" trait_set | type "=" type

type          = primitive_type
              | path type_args?
              | dyn_type
              | "[" type ";" integer_constant_expression "]"
              | "(" type_list ")"
              | type "->" type
primitive_type = "()" | "never" | "bool" | "isize" | "i8" | "i16" | "i32" | "i64"
               | "usize" | "u8" | "u16" | "u32" | "u64" | "f32" | "f64"
               | "string" | "char"
type_args     = "[" generic_argument ("," generic_argument)* ","? "]"
generic_argument = type | integer_literal | "{" integer_constant_expression "}"
integer_constant_expression = integer_literal | path
              | "(" integer_constant_expression ")"
              | ("+" | "-" | "~") integer_constant_expression
              | integer_constant_expression integer_operator integer_constant_expression
integer_operator = "+" | "-" | "*" | "/" | "%" | "<<" | ">>" | "&" | "|" | "^"
constant_expression = expression
type_list     = type ("," type)* ","?
dyn_bound     = path dyn_args?
dyn_args      = "[" (type ",")* dyn_assoc ("," dyn_assoc)* ","? "]"
              | "[" type_list "]"
dyn_assoc     = upper_ident "=" type
dyn_type      = "dyn" dyn_bound ("+" dyn_bound)*
opaque_return = "impl" dyn_bound ("+" dyn_bound)*

block         = "{" statement* expression? "}"
statement     = "let" "mut"? pattern (":" type)? "=" expression ("else" block)? ";"
              | "defer" expression ";"
              | assign_target assignment_operator expression ";"
              | expression ";"
              | unlabeled_control_expression

assignment_operator = "=" | "+=" | "-=" | "*=" | "/=" | "%="
                    | "&=" | "|=" | "^=" | "<<=" | ">>="

expression    = literal | raw_string | byte_string | raw_byte_string
              | interpolated_string | expression_path | tuple | array | block
              | struct_literal | closure
              | call | field | index | unary | binary | cast | range_expression
              | try_expression
              | comptime_expression
              | if_expression | match_expression | select_expression | while_expression | loop_expression | for_expression
              | scope_expression | spawn_expression
              | "return" expression?
              | "break" loop_label? expression?
              | "continue" loop_label?
              | "go" expression

expression_path = path ("::" type_args)? ("::" ident ("::" type_args)?)?

interpolated_string = "f\"" (string_text | "{{" | "}}" | "{" expression "}")* "\""
byte_string   = "b\"" byte_string_content* "\""
raw_byte_string = "br\"...\"" | "br#\"...\"#" | "br##\"...\"##" | ...
raw_string    = "r" raw_hashes? "\"" raw_text "\"" raw_hashes?
array         = "[" (expression ("," expression)* ","?)? "]"
              | "[" expression ";" integer_constant_expression "]"
struct_literal = path "{" (struct_literal_field ("," struct_literal_field)*
                 ("," ".." expression)? ","? | ".." expression ","?)? "}"
struct_literal_field = ident (":" expression)?

unlabeled_control_expression = if_expression | match_expression | select_expression
                   | "while" condition block
                   | "loop" block | "for" pattern "in" expression block
if_expression = "if" condition block ("else" (block | if_expression))?
condition     = expression | "let" pattern "=" expression
              | condition_term ("&&" condition_term)+
condition_term = boolean_term | "let" pattern "=" initializer_term
match_expression = "match" expression
                   "{" (match_arm ",")* match_arm? "}"
match_arm     = pattern ("if" expression)? "=>" (expression | block)
select_expression = "select" "priority"? "{" select_arm ("," select_arm)* ","? "}"
select_arm    = "recv" "(" expression ")" select_guard? receive_continuation
              | "send" "(" expression "," expression ")" select_guard?
                "=>" (expression | block)
              | "default" "=>" (expression | block)
select_guard  = "when" expression
receive_continuation = "as" (lower_ident | "_") "=>" (expression | block)
              | "match" "{" (match_arm ",")* match_arm? "}"
while_expression = loop_label_decl? "while" condition block
loop_expression = loop_label_decl? "loop" block
for_expression = loop_label_decl? "for" pattern "in" expression block
loop_label_decl = loop_label ":"
loop_label    = "'" lower_ident
scope_expression = "scope" block
spawn_expression = "spawn" closure
comptime_expression = "comptime" block
closure       = "||" (expression | block)
              | "|" closure_params? "|" (expression | block)
cast          = expression "as" dyn_type
range_expression = expression (".." | "..=") expression

pattern       = or_pattern
or_pattern    = alias_pattern ("|" alias_pattern)*
alias_pattern = ident "@" alias_pattern | range_pattern
range_pattern = primary_pattern ((".." | "..=") range_endpoint)?
range_endpoint = "-"? integer_literal | char_literal
primary_pattern = ident | "mut" lower_ident | "_" | literal | "-" numeric_literal | "()"
              | "(" pattern ")"
              | "(" pattern "," (pattern ("," pattern)* ","?)? ")"
              | path
              | path "(" pattern_list? ")"
              | path "{" struct_pattern_fields? "}"
              | array_pattern
pattern_list  = pattern ("," pattern)* ","?
struct_pattern_fields = struct_pattern_field ("," struct_pattern_field)*
                        ("," "..")? ","?
                      | ".." ","?
struct_pattern_field = ident (":" pattern)?
array_pattern = "[" (array_pattern_item ("," array_pattern_item)* ","?)? "]"
array_pattern_item = pattern | ".." | ident "@" ".."
```

表达式语句语法也允许内置 `Result` 值，但丢弃会触发 `unused_result`。现有带通配模式 `_` 的 `let` 语句提供显式忽略；此检查不增加语法。

解析器会对逗号和分号做一定错误恢复，但代码代理应始终生成上述规范形式：列表项以逗号分隔；`let`、赋值和普通非尾表达式带分号；无标签控制流语句不加多余分号；有标签循环后若有语句则加分号；trait 方法签名带分号。序列模式最多一个剩余部分；结构体模式中的 `..` 最多一次，且必须位于末尾。

含 `let` 的条件链中，`boolean_term` 和 `initializer_term` 是不含顶层 `&&` 或 `||` 的表达式；分组括号允许单项内部布尔析取。虽然 trait 签名共享上述简化的 `return_type` 规则，`opaque_return` 仍仅限顶层函数和固有方法的完整返回类型。未限定枚举变体使用普通 `path`、调用或结构体字面量语法，并按预期类型解析。

枚举 `expression_path` 中，泛型实参应用于变体名前的枚举本身，如 `Message::[i32]::Value`；变体不能再带第二个列表。`block` 规则引入词法绑定作用域；常量初始化器中的名称和泛型类型参数则遵循声明作用域，而非最终使用位置。

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

代码代理提交源码前至少应运行对应 `goml check` 或 `gomlc run-single`；修改测试时，还应运行 `goml check --tests` 和相关 `goml test`。类型推断失败时，优先添加局部结果类型、空容器类型、闭包参数类型或使用 UFCS，不要改写为不支持的 Rust/Go 语法。
