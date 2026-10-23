<a id="language-basics"></a>

# 语言基础

[English](basics.md) | [简体中文](basics.zh-CN.md)

[GoML 语言指南](../goml.zh-CN.md)

词法规则、定义、函数、绑定、表达式、控制流和闭包。数据定义、泛型和模式规则见[类型与 trait](types.zh-CN.md)。

本页内容：[词法规则](#lexical-rules)、[顶层定义](#top-level-definitions)、[顶层函数](#top-level-functions)、[块、绑定和赋值](#blocks-bindings-and-assignments)、[表达式和运算符](#expressions-and-operators)、[控制流](#control-flow)、[闭包与函数值](#closures-and-function-values)。

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

独立 `gomldoc` 命令、输出模式和安装说明见[文档生成](../documentation.zh-CN.md)。

<a id="identifiers-and-keywords"></a>

### 标识符和关键字

普通标识符仅使用 ASCII，形式如下：

```text
[A-Za-z][A-Za-z_0-9]* | _[A-Za-z_0-9]+
```

单独的 `_` 是通配符，不是普通变量名。当前语法对各类名称的首字母大小写有以下要求：

- 包名、`use package as alias` 中的包别名、函数、方法、参数、局部绑定和 GoML 字段声明必须以小写字母或 `_` 开头；导入项的别名可以遵循被导入项的命名规则；
- 结构体、枚举、trait、枚举变体、泛型参数和关联类型必须以大写字母开头；
- [可移植 SIMD](../stdlib.zh-CN.md#portable-simd) 中列出的定宽向量和掩码类型写法，包括 `u8x16`、`i32x8`、`f32x4`、`f64x4` 和 `mask32x4`，也可以用作结构体名、类型别名和类型位置中的名称。它们是由 `std::simd` 提供的普通作用域名称，不是全局保留关键字；
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

<a id="top-level-functions"></a>

## 顶层函数

```goml
fn add(left: i32, right: i32) -> i32 {
    left + right
}

fn log(message: string) {
    println(message)
}
```

参数类型不能省略。省略 `-> ...` 等价于 `-> ()`。同包顶层函数名必须唯一，不支持按参数类型重载。

<a id="blocks-bindings-and-assignments"></a>

## 块、绑定和赋值

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
- 表达式中的 `as` 专用于创建 `dyn Trait` 值。数值转换使用[整数和浮点转换方法](types.zh-CN.md#basic-types-and-composite-types)；其他非 `dyn` 目标会被拒绝。`use package as alias;` 中独立的 `as` 形式仍用于选择导入别名。

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

<a id="if"></a>

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

<a id="match"></a>

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

<a id="select"></a>

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

<a id="while"></a>

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

<a id="loop"></a>

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

<a id="for"></a>

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

<a id="try-operator"></a>

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

<a id="go"></a>

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
