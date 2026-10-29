<a id="types-patterns-and-traits"></a>

# 类型、模式与 trait

[English](types.md) | [简体中文](types.zh-CN.md)

[GoML 语言指南](../goml.zh-CN.md)

类型、结构体、枚举、泛型参数、模式匹配、trait 和动态分派。函数声明、表达式和控制流见[语言基础](basics.zh-CN.md)。

本页内容：[类型](#type)、[结构体](#structs)、[枚举](#enums)、[函数与泛型](#functions-and-generics)、[模式](#patterns)、[Trait 与 impl](#trait-and-impl)、[`dyn Trait`](#dyn-trait)。

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

`never` 类型的表达式不产生值，可出现在期望其他值类型的位置。但不会隐式转换容器或函数类型：`Vec[never]` 不是 `Vec[isize]`，`() -> never` 也不是 `() -> isize`。要将不返回的函数适配为其他返回类型的回调，请使用具有该预期回调类型的闭包。普通值和正常返回的回调不能满足 `never` 标注。如果 `if` 或 `while` 的条件、`match` 的被匹配表达式，或 `select` 的通道、发送值、守卫具有 `never` 类型，整个控制表达式也具有 `never` 类型，因为必须先求值该输入才能执行分支。分支体仍接受类型检查，并保留普通预期结果类型提供的上下文。外层 `never` 标注不要求这些不可达分支体本身也发散。`match` 分支守卫是条件执行的，单个守卫发散不会直接使整个匹配发散。

值和模式中的 `(value)` 是分组，`(value,)` 是单元素元组。当前类型语法无法直接标注单元素元组：`(T,)` 仍会规范化为 `T`，因此此类值的类型必须从局部上下文推断。

`A -> B -> C` 按右结合解析。建议始终给函数参数列表加括号，尤其是高阶函数：`(A) -> (B) -> C`。

数组长度是非负整数常量表达式。它接受常见进制的整数字面量、具名整数常量、常量泛型参数、括号、一元 `+`、`-`、`~`，以及 `+`、`-`、`*`、`/`、`%`、`<<`、`>>`、`&`、`|`、`^`。整数溢出、除零、无效移位、负长度和循环常量引用均产生诊断。负的最小值字面量 `-9223372036854775808` 可用于常量表达式，但最终长度必须非负。最小有符号索引值对 `-1` 取余为零，而除以 `-1` 会溢出。长度须在目标的有符号索引范围内。这些表达式不执行函数调用：

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

别名不会创建名义上不同的类型，递归别名循环会被拒绝。公开别名可以跨包引用。通过别名进行的关联方法调用和方法值引用使用底层类型的方法；泛型别名的实参可以显式提供，也可以从调用和预期类型中推断。

枚举构造器调用、构造器值和限定模式也接受枚举类型的别名。特化别名保留其固定实参：

```goml
type Maybe[T] = Option[T];
type Flag = Maybe[bool];
let present = Maybe::[i32]::Some(1);
let flag = Flag::Some(true);
```

结构体字面量、结构体更新、newtype 构造器及其模式也接受别名。别名的固定实参和底层字段的可见性仍然生效。

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

`Self` 在 trait 签名和 impl 中表示实现类型；普通顶层函数不能将 `Self` 用作隐式类型参数。在所有者已知为结构体或枚举的 impl 中，结构体构造和模式可使用 `Self { ... }` 或 `Self(value)`，枚举可使用 `Self::Variant(...)` 和 `Self::Variant { ... }`。这些形式保留所有者的泛型实参。通过相等约束确定为该名义类型的类型参数也可使用相同的构造和模式形式；无约束的类型参数不能选择构造器。

<a id="structs"></a>

## 结构体

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

## 枚举

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

关联类型绑定必须引用该约束或唯一父 trait 声明的类型。每个约束内，同一名称只能绑定一次。若多个父 trait 声明同名关联类型，应在声明它的 trait 上指定绑定，例如 `impl Child + A[Item = i32] + B[Item = string]`。

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

类型等式允许在函数内使用相等类型支持的操作。例如，`where T = Vec[i32]` 允许对 `T` 值使用索引，`where T = (i32, bool)` 允许访问 `.0` 和 `.1`。受 `where F = (i32) -> i32` 约束的参数可以作为函数调用，仍检查参数数量和类型。索引赋值仍遵循底层容器限制：`Slice` 保持只读，数组要求可写根，`HashMap` 索引不支持复合赋值。

相等的 dyn 类型也意味着对应的 trait 实参和关联类型绑定相等。trait 约束和关联类型绑定按名称匹配，因此重排它们不会改变推导出的等式。

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

显式泛型实参列表至少要有一个实参，且数量必须与声明的参数数量一致，包括常量参数。使用推断时省略整个列表；`identity::[](1)` 会被拒绝。

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

GoML 对泛型调用进行单态化。递归泛型代码必须只产生有限个具体实例，不能每次递归都不断变为新的嵌套类型。常量实参也受此要求约束。同一调用链上同一函数的递归实例达到 256 个，或一个函数的实例总数达到 4096 个后，继续特化会报告诊断。复用已有实例的调用不会消耗新的实例名额。

<a id="const-generic-parameters"></a>

### 常量泛型参数

函数、结构体、枚举、别名、trait 和实现都可以在类型参数旁声明 `const N: usize`。常量参数在表达式中是值，在泛型应用中是整数实参；不能用作元素类型。每个具体实参都参与单态化和名义类型身份。 结构体和枚举字段中的数组长度要求也适用于每个具体类型应用，包括未使用的签名、未选择的枚举载荷，以及从其他包导入的私有字段。 trait 方法实现必须保持方法泛型参数的数量、顺序和种类（类型或常量）；参数名可以不同。

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

直接的长度参数可从数组实参、预期类型或 `where [T; N] = [i32; 2]` 这样的数组类型等式推断。对于 `N + 1` 这样的长度表达式，请显式提供 `N`；推断不求解算术方程。泛型实参使用字面量或名称，复合整数实参使用花括号，例如 `Buffer[i32, {2 + 1}]`。常量参数目前类型为 `usize`，没有默认值，取值范围为 `0..=9223372036854775807`。泛型替换后会再次检查 trait 约束，包括实现所需的约束；即使常量实参没有出现在函数签名中，也会拒绝负数实参和无效的常量运算。数组重复接受这些参数，例如 `[value; N]`。不支持数组长度中的类型相关关联常量。

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

当 `N` 是尚未确定值的 const 泛型参数时，数组模式会在各个具体实例中检查所需长度。`[rest @ ..]` 这样的仅含剩余绑定的模式是不可反驳的，并保留 `[T; N]` 类型。要求特定长度或非空前缀、后缀的模式需要后备分支，除非长度已知。具名剩余绑定仍是数组；在匹配的分支中，其长度等于原长度减去显式元素数量。更短的实例会进入后备分支，不会求值剩余绑定。

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

每一步都有其声明 trait 的约束时，关联类型投影可链式连接。例如，`I::IntoIter::Item` 先从 `I` 投影 `IntoIter`，再从迭代器类型投影 `Item`。关联类型约束可以递归，例如 `trait Node { type Next: Node; }`；任意有限的 `Next` 投影链都保留 `Node` 约束。链式投影可出现在签名和 `where` 等式中：

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

默认方法的泛型参数属于其 trait 声明，与实现类型的参数保持独立，即使名称相同也不会混淆。const 参数也遵循此规则。

即使显式方法泛型实参只用于方法体，它们也参与特化。此规则同样适用于方法值、默认实现和直接调用。无法从接收者、实参、结果或约束推断的泛型实参必须显式提供。

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

方法类型实参通常自动推断。需要显式实参时写 `box_value.map::[string](../convert)`，关联形式写 `Box::[i32]::map::[string](box_value, convert)`。

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

多个可见 trait 定义同名方法时，预期结果类型可过滤不兼容的候选。对于 `dyn` 结果，这包括所有 trait 实参、关联类型绑定和附加约束。返回 `never` 的方法与任意预期结果类型兼容。这种表达式转换不会使 `Vec[never]` 等嵌套类型与 `Vec[i32]` 兼容。如果仍有多个候选，`value.render()` 会有歧义。使用 UFCS 显式选择：

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

Trait 约束使相应方法可用于泛型值。通过泛型约束调用时，会使用实参类型和预期结果筛选同名候选，包括通过类型参数进行的关联调用。 约束也允许通过类型参数使用关联调用和方法值，例如 `T::method()` 和 `T::method`；没有接收者的方法仍不能通过值调用。导入的 trait 短名、重命名导入和公开重新导出，在特化过程中保留泛型约束及 `where` 谓词中定义 trait 的身份。父 trait 和关联类型约束也以隐含约束参与方法解析。

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

<a id="dyn-trait"></a>

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

预期 `dyn Trait` 值时，具体值不会隐式装箱。赋值、实参、返回、分支、构造器和集合操作都必须在转换处写 `as dyn Trait`。已有 `dyn Trait` 值仍可在兼容类型上下文间传递，无需再次转换。类型为 `never` 的源表达式会在转换前发散；转换表达式也具有 `never` 类型，不要求该类型实现目标 trait。

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

父 trait 方法仍可用。 只要 trait 实参匹配，dyn 类型也满足其主 trait、附加 trait 和父 trait 对应的泛型约束。泛型代码中的调用仍使用动态分派。不同组成 trait 提供同名方法时，使用 UFCS 选择 trait，例如 `A::render(value)` 或 `B::render(value)`。泛型 trait 使用 `Consumer::[string]::consume(value, "hello")`。

约束顺序不改变对象类型：`dyn A + B` 与 `dyn B + A` 兼容。重复 trait 名会被拒绝，即使泛型实参不同。父 trait 路径对同一 trait 的泛型实参必须一致。子 trait 方法优先于从父 trait 继承的方法；无关同名方法需要 UFCS。

唯一名称的继承关联类型可在子 trait 上绑定，例如 `dyn Child[Item = isize]`，或在声明它的 trait 上绑定，如 `dyn Child + Parent[Item = isize]`。不同父 trait 声明同名关联类型时，分别在声明 trait 上绑定，如 `dyn Child + A[Item = isize] + B[Item = string]`。

当前 dyn 安全条件：

- 每个方法的第一个接收者参数必须恰为 `Self` 类型；
- 其他参数或返回类型不能直接出现 `Self`，包括嵌套 `dyn` 的实参、关联类型绑定和附加约束内部；允许 `Self::Item` 这样的已绑定投影；
- 方法不能声明类型参数；
- Trait 不能声明关联常量；
- 每个组成 trait 和父 trait 都必须满足这些规则。

当前限制：

- dyn 实参列表不支持 `dyn Fixed[4]` 这样的整数常量实参；
- 不支持从一种 dyn 对象类型转换为另一种，包括删除约束；
- 不支持对 `dyn Trait` 进行模式匹配。
