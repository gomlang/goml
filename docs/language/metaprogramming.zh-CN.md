<a id="compile-time-evaluation-and-derive"></a>

# 编译期求值与派生

[English](metaprogramming.md) | [简体中文](metaprogramming.zh-CN.md)

[GoML 语言指南](../goml.zh-CN.md)

编译期求值、属性和可编程派生的语言规则。编译器实现细节见[编译期求值架构](../comptime.zh-CN.md)。

本页内容：[编译期求值](#compile-time-evaluation)、[属性与 derive](#attributes-and-derive)。

<a id="compile-time-evaluation"></a>

## 编译期求值

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

标准库通过 `std::serde` 提供序列化 derive。API、属性和示例见 [Serde 与 JSON](../stdlib.zh-CN.md#serde-and-json)。

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
