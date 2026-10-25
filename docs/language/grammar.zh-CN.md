<a id="grammar-and-common-mistakes"></a>

# 文法与常见错误

[English](grammar.md) | [简体中文](grammar.zh-CN.md)

[GoML 语言指南](../goml.zh-CN.md)

语法对照和非正式文法。相应语义规则见[语言基础](basics.zh-CN.md)和[类型与 trait](types.zh-CN.md)。

本页内容：[常见写法错误对照](#comparison-of-common-writing-errors)、[非正式语法速查](#informal-grammar-quick-facts)。

<a id="comparison-of-common-writing-errors"></a>

## 常见写法错误对照

库 API 和常见任务见[标准库参考](../stdlib.zh-CN.md#common-tasks)。

类型推断失败时，可以标注局部结果、空容器或闭包参数的类型，或使用 UFCS。方法解析规则见[方法调用与 UFCS](types.zh-CN.md#method-parsing-and-ufcs)。

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
| 跨包共享实现 API，但不向外部模块公开 | 对声明、字段、方法和重导出使用 `pub(module)` |
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
use_decl      = visibility? "use" use_path ("as" ident | "::" "{" use_items "}")? ";"
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
visibility    = "pub" ("(" "module" ")")?
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
