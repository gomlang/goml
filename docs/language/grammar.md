<a id="grammar-and-common-mistakes"></a>

# Grammar and common mistakes

[English](grammar.md) | [简体中文](grammar.zh-CN.md)

[GoML language guide](../goml.md)

A syntax comparison and the informal grammar. Read [language basics](basics.md) and [types and traits](types.md) for the corresponding semantic rules.

On this page: [Comparison of common writing errors](#comparison-of-common-writing-errors), [Informal Grammar Quick Facts](#informal-grammar-quick-facts).

<a id="comparison-of-common-writing-errors"></a>

## Comparison of common writing errors

For library APIs and common tasks, see the [standard library reference](../stdlib.md#common-tasks).

| Avoid | GoML form |
| --- | --- |
| `////` or a blank source line inside an API documentation block | Use `///`; use a bare `///` line to separate paragraphs |
| `Vec<isize>` | `Vec[isize]` |
| `fn id<T>(x: T) -> T` | `fn id[T](x: T) -> T` |
| `fn count[N](values: [i32; N])` | Declare the value parameter as `fn count[const N: usize](values: [i32; N]) -> usize` |
| `Buffer[i32, N + 1]` | `Buffer[i32, {N + 1}]` |
| Using `serde(rename)` to change `.to_string()` | Use `#[to_string("text")]` on the variant with `#[derive(ToString)]` |
| `Type::LIMIT()` for an associated constant | `Type::LIMIT` |
| `id::<i32>(1)` | `id::[i32](../1)` |
| Ordinary function `id[i32](../1)` | `id::[i32](../1)`, or rely on parameter/result type inference |
| `1i32`, `1u64`, `1.0f32` | Use the expected type, such as `let value: u64 = 1;` |
| Non-ASCII source text in `b"é"` | Use a `string`, or write its encoded bytes explicitly such as `b"\xC3\xA9"` |
| A bare block as a control-flow header value | Parenthesize it, for example `if ({ prepare(); ready() }) { ... }` |
| `let mut x: &T` | Use value `T` or `Ref[T]` as required |
| Write `if cond { value }` in the value position | `if cond { value } else { other }` |
| `let Option::Some(x) = value;` | `let Some(x) = value else { return };`, `if let`, or `match` |
| `let Point { x } = point;` | `let Point { x, .. } = point;` |
| Non-exhaustive `match` | Cover every possible variant or add a `_` branch |
| Discard a `Result` with `operation();` | Handle it, propagate with a compatible `?`, or explicitly ignore it with `let _ = operation();` |
| `x++`, `x--` | `x += 1;`, `x -= 1;` |
| Assign through an immutable structure binding | Declare the binding with `let mut`, or create a new value with `Point { field: value, ..point }` |
| `var x = 1`, `x := 1` | `let x = 1;` |
| A loop with a condition | `while condition { ... }` |
| `for i := 0; ...` | `while`, or `for i in start..end` |
| `switch` | `match` |
| `null`, `nil` | `None` with an expected `Option` type, or `Option::None`; `ffi::null()` for raw Go pointers; `ffi::nil_error()` for Go errors |
| `throw`, exception | `Result` and `?` for expected errors; `std::panic::catch` for an explicit runtime-panic boundary |
| Implicit conversion to `dyn A + B` | Write `value as dyn A + B`; the value must implement every bound |
| `dyn Consumer` for `trait Consumer[T]` | Supply trait arguments, for example `dyn Consumer[string]` |
| `dyn TraitWithAssociatedType` | Bind every associated type, for example `dyn Iterator[Item = isize]` |
| Different concrete return types under `-> impl Trait` | Return one concrete type from every branch, or explicitly return `dyn Trait` |
| `fn consume(value: impl Trait)` | Use a named generic parameter, for example `fn consume[T: Trait](value: T)` |
| `[initializer(); count]` for a runtime `count` | Repeat lengths are compile-time integers; use a `Vec` and a loop for runtime lengths |
| Expect `[initializer(); N]` to call once per element | The initializer runs once, including for zero length; use a loop for per-element work |
| `if let Some(x) = value || ready` | Use nested control flow, or `&& (first || second)` for a boolean term within a let chain |
| Use `type UserId = u64;` when `UserId` must be distinct | Use `struct UserId(u64);` and construct it explicitly |
| `use pkg::*` | List the required public items explicitly with `use pkg::{A, B};` |
| `mod`, `crate::`, `super::` | Directory packages, `module::path` for the current module, and canonical paths for dependencies |
| Develop several modules together | `goml.work` lists members; dependency `{ path = "../library" }` and root `[replace]` select local sources while preserving declared dependencies and canonical identities; `--workspace` or `-p` selects modules |
| Put test-only helpers in production dependencies | Use root `[dev-dependencies]`; tests and `examples/<name>/` can import them |
| Create a separate consumer manifest for every example | Use `goml run --example <name>` and `goml test`; examples share the module-root manifest |
| `fn helper` inside function | Top-level function or local closure |
| Use a local binding after leaving its block | Return the value from the block and bind its result outside |
| `Message::Value::[i32](../42)` | Enum arguments belong to the owner: `Message::[i32]::Value(42)` |
| `identity::[](../1)` | Omit the empty turbofish and write `identity(1)` |
| `Message::[bool]::Named { value: true }` | Use `let value: Message[bool] = Message::Named { value: true };` |
| Expect a function's const parameter to change a top-level constant initializer | Constants retain the bindings from their declaration site |
| Go external type | `#[go_type("pkg", "Name")] extern type Name[T];` retains an opaque Go identity and declared arguments; use explicit functions and methods, with `std::ffi::Ptr[T]` for raw pointers. No automatic fields, struct construction, alias expansion or constraint discovery; final Go compilation validates native uses |
| Go interface adapter | `#[go_interface(RawType, Wrapper, method = "GoMethod")]` generates adapters from explicit trait signatures and mappings; final Go compilation checks the real interface. `from_trait` retains the supplied dyn object; nil/typed-nil and multiple results are preserved |
| Go-callable export | Annotate a supported public function with `#[go_export("Name")]` and generate a Go package with `goml export-go` |
| Customize enum display text | `#[derive(ToString)]` with `#[to_string("text {field}")]` on variants |
| Generate public inherent methods | `derive_output_inherent` and `derive_output_add_public_method` |
| Generate multiple implementations and helpers | `derive_output_merge`, associated item builders, and private function/type builders |
| Build runtime callbacks in derive | `meta_expr_closure` and `meta_type_function` |
| Reuse compile-time computations | Closures, indirect calls, concrete methods, local dynamic dispatch, `Ref`, and `Vec` operations |
| Unannotated user `extern fn` | Use a normal GoML function or `#[go_ffi("import/path", "ExportedSymbol")] extern fn`; GoML trusts the declared signature and final Go compilation validates emitted calls |
| Call an ordinary function from `comptime` | Mark a supported free function with `#[comptime]` |
| Capture a runtime local in `comptime` | Pass a literal or compile-time value to a `#[comptime]` function |

The documentation spelling `////` is an ordinary comment, not an item documentation block. A blank line between `///` and a declaration leaves that comment unattached; use a bare `///` line to separate paragraphs within a block.

<a id="informal-grammar-quick-facts"></a>

## Informal Grammar Quick Facts

Development dependencies, named examples and downstream verification affect manifests, directory discovery and CLI commands. They add no GoML source grammar productions; examples use ordinary package declarations and imports.

Lexical comment conventions (comments are trivia in the grammar below):

```text
LineComment    ::= "//" { any character except newline }
ItemDocLine    ::= "///" { any character except newline }
PackageDocLine ::= "//!" { any character except newline }
```

`ItemDocLine` excludes lines beginning `////` and requires a standalone comment before a declaration. `PackageDocLine` is recognized only before the first source token.

Base32, varints, checksums, UTF-8 scalar/stream decoding, map/slice helpers, logical slash paths and fixed-width bit operations use ordinary imports, traits, structs, enums, slices and calls. They add no grammar, implicit byte conversions or compiler intrinsics. Map and slice collection use existing `Iterator` associated-type constraints. Shell patterns are string data parsed by `std::path::slash`, not GoML syntax.

`std::c` uses existing structs, constants, `#[comptime]`, imports and calls. It adds no C pointer, C layout or `extern "C"` grammar.

Linux syscall buffers, `Pointer` descriptors, `Errno` values, ABI codecs, descriptor/process/memory/IPC wrappers use the ordinary struct, enum, array, slice, and call forms below. Native kernel layouts are encoded into bytes; there is no pointer-cast, native-layout, or `unsafe` grammar.

Panic boundaries use ordinary function calls, closures, `Result`, and the existing `never` type. There is no `try/catch`, `throw`, or bare Go-style `recover()` syntax; panic cleanup extends the semantics of the existing lexical `defer` statement.

`Report::map`, `ResultExt::into_report`, and `ReportContext::{context, with_context}` use the existing generic-method, trait-import and closure forms. They add no grammar production or implicit error conversion.

I/O traits, cancellation contexts, DNS/TLS clients and listeners, scalar conversion traits, and Serde extension events use the existing imports, generic bounds, enums, methods, and calls. They introduce no new grammar. Explicit generic calls retain the `::[Type]` form even when their type arguments do not appear in the function signature.

The following EBNF only describes the canonical form that should be generated; `?` means optional, `*` means repeated, and the terminator is placed in quotes.

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

The expression-statement grammar also permits a built-in `Result` value, but discarding it triggers `unused_result`. The existing `let` statement with wildcard pattern `_` provides explicit ignoring; this lint adds no syntax.

The parser will do some error recovery for commas and semicolons, but the code agent should always generate the above canonical form: list items separated by commas, `let`, assignments and ordinary non-tail expressions with semicolons, unlabeled control-flow statements without unnecessary semicolons, labeled loop statements with semicolons before following statements, and trait method signatures with semicolons. The sequence pattern contains at most one rest; the `..` in the structure pattern appears at most once and must be at the end.

In a condition chain containing `let`, `boolean_term` and `initializer_term` are expressions without top-level `&&` or `||`; grouping parentheses permit a boolean disjunction inside one term. `opaque_return` is restricted to complete top-level function and inherent method returns, even though trait signatures share the abbreviated `return_type` rule above. An unqualified enum variant uses the ordinary `path`, call, or struct-literal grammar and is resolved using its expected type.

In an enum `expression_path`, generic arguments apply to the owner before the variant name, as in `Message::[i32]::Value`; variants cannot take a second list. The `block` rule introduces a lexical binding scope, while names in constant initializers and generic type parameters follow their declaration scopes rather than their eventual use sites.
