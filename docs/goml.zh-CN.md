<a id="goml-language-guide"></a>

# GoML 语言指南

[English](goml.md) | [简体中文](goml.zh-CN.md)

本指南介绍本仓库实现的 GoML 语法、类型规则、工具和内置 prelude。标准库 API 见[标准库参考](stdlib.zh-CN.md)。编写、修改和审阅 `.goml` 源码时请以本指南为准。与 Rust、Go 或 OCaml 相似，并不意味着支持这些语言的语法或 API。

GoML 是一种具有垃圾回收的静态类型语言。语法接近 Rust，语义更接近 ML。编译器在生成 Go 代码前，对泛型进行单态化并对 GoML 闭包进行闭包提升。显式 Go FFI 可以保留原生泛型类型和函数值。GoML 没有所有权、借用、生命周期或手动内存管理。

生成的 Go 代码和 Go FFI 以 Go 1.26 为兼容性基线。构建工具链和编译生成的程序需要 Go 1.26 或更高版本。原生 FFI 签名由最后的 Go 编译步骤检查。独立的 [gomlgo 前端和解释器](https://github.com/gomlang/gomlgo) 以 Go 1.26 的语言和标准库行为为目标；其执行和差分测试需要 Go 1.26.x。

快速导航：[包](language/packages.zh-CN.md#modules-packages-and-imports)、[类型](language/types.zh-CN.md#type)、[控制流](language/basics.zh-CN.md#control-flow)、[模式](language/types.zh-CN.md#patterns)、[trait](language/types.zh-CN.md#trait-and-impl)、[编译期求值](language/metaprogramming.zh-CN.md#compile-time-evaluation)、[Go FFI](ffi/bind-go.zh-CN.md#go-ffi)、[测试](tooling.zh-CN.md#tests)、[API 文档](language/basics.zh-CN.md#comments-and-api-documentation)、[prelude](language/prelude.zh-CN.md#built-in-prelude)、[标准库](stdlib.zh-CN.md)和[验证](tooling.zh-CN.md#verify-generated-code)。

构建和安装说明见[开发指南](development.zh-CN.md)和[发布指南](releasing.zh-CN.md#installation)。源码的规范布局见[格式化](formatting.zh-CN.md)，CTIR 内部机制见[编译期求值架构](comptime.zh-CN.md)。

<a id="reading-guide"></a>

## 阅读路线

| 专题 | 内容 |
| --- | --- |
| [语言基础](language/basics.zh-CN.md) | 词法规则、定义、函数、绑定、表达式、控制流和闭包 |
| [类型、模式与 trait](language/types.zh-CN.md) | 类型、结构体、枚举、泛型参数、模式匹配、trait 和动态分派 |
| [模块、包与可见性](language/packages.zh-CN.md) | 源码中的包身份、导入和可见性 |
| [编译期求值与派生](language/metaprogramming.zh-CN.md) | 编译期求值、属性和可编程派生的语言规则 |
| [Prelude 参考](language/prelude.zh-CN.md) | 自动进入作用域的类型、trait 和操作 |
| [文法与常见错误](language/grammar.zh-CN.md) | 语法对照和非正式文法 |
| [项目工具与工作流](tooling.zh-CN.md) | 模块配置、依赖、示例、构建输出、测试、编辑器支持和生成代码检查 |
| [Go FFI](ffi/bind-go.zh-CN.md) | Go 绑定、原始边界类型和 Go 库导出 |

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
12. 当前模块根目录下的包使用 `module::path`。不要生成 `mod`、`crate::`、`self::`、`super::`、根路径 `::x`、Rust 引用或 Go 的 `var` / `:=`。用户的 `extern fn` 只有配合[类型化 Go FFI 属性](ffi/bind-go.zh-CN.md#go-ffi)才有效。
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

<a id="original-section-links"></a>

## 原章节链接

原有章节锚点保留在下表中，链接指向各专题的新位置。

| 原章节 | 当前文档 |
| --- | --- |
| <a id="lexical-rules"></a><a id="词法规则"></a> 词法规则 | [语言基础](language/basics.zh-CN.md#lexical-rules) |
| <a id="comments-and-api-documentation"></a><a id="注释和-api-文档"></a> 注释和 API 文档 | [语言基础](language/basics.zh-CN.md#comments-and-api-documentation) |
| <a id="identifiers-and-keywords"></a><a id="标识符和关键字"></a> 标识符和关键字 | [语言基础](language/basics.zh-CN.md#identifiers-and-keywords) |
| <a id="literals"></a><a id="字面量"></a> 字面量 | [语言基础](language/basics.zh-CN.md#literals) |
| <a id="modules-packages-and-imports"></a><a id="模块包和导入"></a> 模块、包和导入 | [模块、包与可见性](language/packages.zh-CN.md#modules-packages-and-imports) |
| <a id="gomltoml"></a> `goml.toml` | [项目工具与工作流](tooling.zh-CN.md#gomltoml) |
| <a id="package-registry"></a><a id="包注册表"></a> 包注册表 | [项目工具与工作流](tooling.zh-CN.md#package-registry) |
| <a id="examples"></a><a id="示例"></a> 示例 | [项目工具与工作流](tooling.zh-CN.md#examples) |
| <a id="native-dependencies"></a><a id="原生依赖"></a> 原生依赖 | [项目工具与工作流](tooling.zh-CN.md#native-dependencies) |
| <a id="build-output-layout"></a><a id="构建输出布局"></a> 构建输出布局 | [项目工具与工作流](tooling.zh-CN.md#build-output-layout) |
| <a id="directory-packages"></a><a id="目录包"></a> 目录包 | [模块、包与可见性](language/packages.zh-CN.md#directory-packages) |
| <a id="use"></a> `use` | [模块、包与可见性](language/packages.zh-CN.md#use) |
| <a id="visibility"></a><a id="可见性"></a> 可见性 | [模块、包与可见性](language/packages.zh-CN.md#visibility) |
| <a id="type"></a><a id="类型"></a> 类型 | [类型、模式与 trait](language/types.zh-CN.md#type) |
| <a id="basic-types-and-composite-types"></a><a id="基本类型和复合类型"></a> 基本类型和复合类型 | [类型、模式与 trait](language/types.zh-CN.md#basic-types-and-composite-types) |
| <a id="type-aliases"></a><a id="类型别名"></a> 类型别名 | [类型、模式与 trait](language/types.zh-CN.md#type-aliases) |
| <a id="type-syntax-not-currently-available"></a><a id="当前不支持的类型语法"></a> 当前不支持的类型语法 | [类型、模式与 trait](language/types.zh-CN.md#type-syntax-not-currently-available) |
| <a id="top-level-definitions"></a><a id="顶层定义"></a> 顶层定义 | [语言基础](language/basics.zh-CN.md#top-level-definitions) |
| <a id="constants"></a><a id="常量"></a> 常量 | [语言基础](language/basics.zh-CN.md#constants) |
| <a id="statics-and-one-time-initialization"></a><a id="静态变量与一次性初始化"></a> 静态变量与一次性初始化 | [语言基础](language/basics.zh-CN.md#statics-and-one-time-initialization) |
| <a id="compile-time-evaluation"></a><a id="编译期求值"></a> 编译期求值 | [编译期求值与派生](language/metaprogramming.zh-CN.md#compile-time-evaluation) |
| <a id="structs"></a><a id="结构体"></a> 结构体 | [类型、模式与 trait](language/types.zh-CN.md#structs) |
| <a id="enums"></a><a id="枚举"></a> 枚举 | [类型、模式与 trait](language/types.zh-CN.md#enums) |
| <a id="functions-and-generics"></a><a id="函数与泛型"></a> 函数与泛型 | [类型、模式与 trait](language/types.zh-CN.md#functions-and-generics) |
| <a id="top-level-functions"></a><a id="顶层函数"></a> 顶层函数 | [语言基础](language/basics.zh-CN.md#top-level-functions) |
| <a id="opaque-return-types"></a><a id="不透明返回类型"></a> 不透明返回类型 | [类型、模式与 trait](language/types.zh-CN.md#opaque-return-types) |
| <a id="generic-functions"></a><a id="泛型函数"></a> 泛型函数 | [类型、模式与 trait](language/types.zh-CN.md#generic-functions) |
| <a id="blocks-bindings-and-assignments"></a><a id="块绑定和赋值"></a> 块、绑定和赋值 | [语言基础](language/basics.zh-CN.md#blocks-bindings-and-assignments) |
| <a id="const-generic-parameters"></a><a id="常量泛型参数"></a> 常量泛型参数 | [类型、模式与 trait](language/types.zh-CN.md#const-generic-parameters) |
| <a id="block-values-and-semicolons"></a><a id="块值与分号"></a> 块值与分号 | [语言基础](language/basics.zh-CN.md#block-values-and-semicolons) |
| <a id="let-type-annotations-and-shadowing"></a><a id="let类型标注与遮蔽"></a> `let`、类型标注与遮蔽 | [语言基础](language/basics.zh-CN.md#let-type-annotations-and-shadowing) |
| <a id="mut-and-assignment"></a><a id="mut-与赋值"></a> `mut` 与赋值 | [语言基础](language/basics.zh-CN.md#mut-and-assignment) |
| <a id="expressions-and-operators"></a><a id="表达式和运算符"></a> 表达式和运算符 | [语言基础](language/basics.zh-CN.md#expressions-and-operators) |
| <a id="basic-expressions"></a><a id="基本表达式"></a> 基本表达式 | [语言基础](language/basics.zh-CN.md#basic-expressions) |
| <a id="operator-precedence"></a><a id="运算符优先级"></a> 运算符优先级 | [语言基础](language/basics.zh-CN.md#operator-precedence) |
| <a id="operator-type-rules"></a><a id="运算符类型规则"></a> 运算符类型规则 | [语言基础](language/basics.zh-CN.md#operator-type-rules) |
| <a id="range-expressions"></a><a id="范围表达式"></a> 范围表达式 | [语言基础](language/basics.zh-CN.md#range-expressions) |
| <a id="control-flow"></a><a id="控制流"></a> 控制流 | [语言基础](language/basics.zh-CN.md#control-flow) |
| <a id="if"></a> `if` | [语言基础](language/basics.zh-CN.md#if) |
| <a id="match"></a> `match` | [语言基础](language/basics.zh-CN.md#match) |
| <a id="select"></a> `select` | [语言基础](language/basics.zh-CN.md#select) |
| <a id="while"></a> `while` | [语言基础](language/basics.zh-CN.md#while) |
| <a id="loop"></a> `loop` | [语言基础](language/basics.zh-CN.md#loop) |
| <a id="for"></a> `for` | [语言基础](language/basics.zh-CN.md#for) |
| <a id="break-continue-and-return"></a><a id="breakcontinue-和-return"></a> `break`、`continue` 和 `return` | [语言基础](language/basics.zh-CN.md#break-continue-and-return) |
| <a id="try-operator"></a> `?` | [语言基础](language/basics.zh-CN.md#try-operator) |
| <a id="go"></a> `go` | [语言基础](language/basics.zh-CN.md#go) |
| <a id="patterns"></a><a id="模式"></a> 模式 | [类型、模式与 trait](language/types.zh-CN.md#patterns) |
| <a id="refutability-and-binding"></a><a id="可反驳性与绑定"></a> 可反驳性与绑定 | [类型、模式与 trait](language/types.zh-CN.md#refutability-and-binding) |
| <a id="groups-tuples-structures-and-enumerations"></a><a id="分组元组结构体和枚举"></a> 分组、元组、结构体和枚举 | [类型、模式与 trait](language/types.zh-CN.md#groups-tuples-structures-and-enumerations) |
| <a id="fixed-arrays-vec-slice-and-mutslice"></a><a id="定长数组vecslice-和-mutslice"></a> 定长数组、Vec、Slice 和 MutSlice | [类型、模式与 trait](language/types.zh-CN.md#fixed-arrays-vec-slice-and-mutslice) |
| <a id="alias-and-or-pattern"></a><a id="别名与或模式"></a> 别名与或模式 | [类型、模式与 trait](language/types.zh-CN.md#alias-and-or-pattern) |
| <a id="range-patterns"></a><a id="范围模式"></a> 范围模式 | [类型、模式与 trait](language/types.zh-CN.md#range-patterns) |
| <a id="exhaustiveness-and-unreachable-branches"></a><a id="穷尽性与不可达分支"></a> 穷尽性与不可达分支 | [类型、模式与 trait](language/types.zh-CN.md#exhaustiveness-and-unreachable-branches) |
| <a id="closures-and-function-values"></a><a id="闭包与函数值"></a> 闭包与函数值 | [语言基础](language/basics.zh-CN.md#closures-and-function-values) |
| <a id="closure-syntax"></a><a id="闭包语法"></a> 闭包语法 | [语言基础](language/basics.zh-CN.md#closure-syntax) |
| <a id="no-let-polymorphism"></a><a id="没有-let-多态"></a> 没有 let 多态 | [语言基础](language/basics.zh-CN.md#no-let-polymorphism) |
| <a id="trait-and-impl"></a><a id="trait-与-impl"></a> Trait 与 impl | [类型、模式与 trait](language/types.zh-CN.md#trait-and-impl) |
| <a id="define-traits"></a><a id="定义-trait"></a> 定义 trait | [类型、模式与 trait](language/types.zh-CN.md#define-traits) |
| <a id="trait-impl"></a><a id="trait-实现"></a> Trait 实现 | [类型、模式与 trait](language/types.zh-CN.md#trait-impl) |
| <a id="inherent-impl"></a><a id="固有-impl"></a> 固有 impl | [类型、模式与 trait](language/types.zh-CN.md#inherent-impl) |
| <a id="method-parsing-and-ufcs"></a><a id="方法解析与-ufcs"></a> 方法解析与 UFCS | [类型、模式与 trait](language/types.zh-CN.md#method-parsing-and-ufcs) |
| <a id="associated-constants"></a><a id="关联常量"></a> 关联常量 | [类型、模式与 trait](language/types.zh-CN.md#associated-constants) |
| <a id="dyn-trait"></a> `dyn Trait` | [类型、模式与 trait](language/types.zh-CN.md#dyn-trait) |
| <a id="attributes-and-derive"></a><a id="属性与-derive"></a> 属性与 derive | [编译期求值与派生](language/metaprogramming.zh-CN.md#attributes-and-derive) |
| <a id="programmable-derive"></a><a id="可编程-derive"></a> 可编程 derive | [编译期求值与派生](language/metaprogramming.zh-CN.md#programmable-derive) |
| <a id="go-ffi"></a> Go FFI | [Go FFI](ffi/bind-go.zh-CN.md#go-ffi) |
| <a id="go-interface-wrappers"></a><a id="go-接口包装器"></a> Go 接口包装器 | [Go FFI](ffi/bind-go.zh-CN.md#go-interface-wrappers) |
| <a id="raw-go-function-values"></a><a id="原始-go-函数值"></a> 原始 Go 函数值 | [Go FFI](ffi/bind-go.zh-CN.md#raw-go-function-values) |
| <a id="raw-go-strings"></a><a id="原始-go-字符串"></a> 原始 Go 字符串 | [Go FFI](ffi/bind-go.zh-CN.md#raw-go-strings) |
| <a id="raw-go-runes"></a><a id="原始-go-rune"></a> 原始 Go rune | [Go FFI](ffi/bind-go.zh-CN.md#raw-go-runes) |
| <a id="raw-go-slices"></a><a id="原始-go-切片"></a> 原始 Go 切片 | [Go FFI](ffi/bind-go.zh-CN.md#raw-go-slices) |
| <a id="raw-go-maps"></a><a id="原始-go-映射"></a> 原始 Go 映射 | [Go FFI](ffi/bind-go.zh-CN.md#raw-go-maps) |
| <a id="declaring-go-bindings"></a><a id="声明-go-绑定"></a> 声明 Go 绑定 | [Go FFI](ffi/bind-go.zh-CN.md#declaring-go-bindings) |
| <a id="c-compatible-strings"></a><a id="c-兼容字符串"></a> C 兼容字符串 | [Go FFI](ffi/bind-go.zh-CN.md#c-compatible-strings) |
| <a id="exporting-a-go-library"></a><a id="导出-go-库"></a> 导出 Go 库 | [Go FFI](ffi/bind-go.zh-CN.md#exporting-a-go-library) |
| <a id="tests"></a><a id="测试"></a> 测试 | [项目工具与工作流](tooling.zh-CN.md#tests) |
| <a id="test-functions-and-attributes"></a><a id="测试函数与属性"></a> 测试函数与属性 | [项目工具与工作流](tooling.zh-CN.md#test-functions-and-attributes) |
| <a id="white-box-and-black-box-tests"></a><a id="白盒与黑盒测试"></a> 白盒与黑盒测试 | [项目工具与工作流](tooling.zh-CN.md#white-box-and-black-box-tests) |
| <a id="check-and-run-tests"></a><a id="检查与运行测试"></a> 检查与运行测试 | [项目工具与工作流](tooling.zh-CN.md#check-and-run-tests) |
| <a id="lsp-and-editor"></a><a id="lsp-与编辑器"></a> LSP 与编辑器 | [项目工具与工作流](tooling.zh-CN.md#lsp-and-editor) |
| <a id="built-in-prelude"></a><a id="内置-prelude"></a> 内置 prelude | [Prelude 参考](language/prelude.zh-CN.md#built-in-prelude) |
| <a id="option-and-result"></a><a id="option-与-result"></a> `Option` 与 `Result` | [Prelude 参考](language/prelude.zh-CN.md#option-and-result) |
| <a id="default"></a> `Default` | [Prelude 参考](language/prelude.zh-CN.md#default) |
| <a id="output-and-string-conversion"></a><a id="输出与字符串转换"></a> 输出与字符串转换 | [Prelude 参考](language/prelude.zh-CN.md#output-and-string-conversion) |
| <a id="reft"></a> `Ref[T]` | [Prelude 参考](language/prelude.zh-CN.md#reft) |
| <a id="fixed-arrays"></a><a id="定长数组"></a> 定长数组 | [Prelude 参考](language/prelude.zh-CN.md#fixed-arrays) |
| <a id="vect"></a> `Vec[T]` | [Prelude 参考](language/prelude.zh-CN.md#vect) |
| <a id="frozenvect"></a> `FrozenVec[T]` | [Prelude 参考](language/prelude.zh-CN.md#frozenvect) |
| <a id="slicet-and-mutslicet"></a><a id="slicet-与-mutslicet"></a> `Slice[T]` 与 `MutSlice[T]` | [Prelude 参考](language/prelude.zh-CN.md#slicet-and-mutslicet) |
| <a id="hashmapk-v"></a> `HashMap[K, V]` | [Prelude 参考](language/prelude.zh-CN.md#hashmapk-v) |
| <a id="channels"></a><a id="通道"></a> 通道 | [Prelude 参考](language/prelude.zh-CN.md#channels) |
| <a id="iterator"></a><a id="迭代器"></a> 迭代器 | [Prelude 参考](language/prelude.zh-CN.md#iterator) |
| <a id="comparison-of-common-writing-errors"></a><a id="常见写法错误对照"></a> 常见写法错误对照 | [文法与常见错误](language/grammar.zh-CN.md#comparison-of-common-writing-errors) |
| <a id="informal-grammar-quick-facts"></a><a id="非正式语法速查"></a> 非正式语法速查 | [文法与常见错误](language/grammar.zh-CN.md#informal-grammar-quick-facts) |
| <a id="verify-generated-code"></a><a id="验证生成代码"></a> 验证生成代码 | [项目工具与工作流](tooling.zh-CN.md#verify-generated-code) |
