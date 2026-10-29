<a id="compile-time-evaluation-and-derive"></a>

# Compile-time evaluation and derive

[English](metaprogramming.md) | [简体中文](metaprogramming.zh-CN.md)

[GoML language guide](../goml.md)

Language rules for compile-time evaluation, attributes, and programmable derive. Compiler implementation details are covered in [compile-time evaluation architecture](../comptime.md).

On this page: [Compile-time evaluation](#compile-time-evaluation), [Attributes and derive](#attributes-and-derive).

<a id="compile-time-evaluation"></a>

## Compile-time evaluation

`comptime { expression }` requires the expression to be evaluated while the current package is compiled. Its static type is the type of the inner expression, and the compiler replaces it with an equivalent ordinary value before Core lowering:

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

`#[comptime]` marks a non-generic free function or concrete implementation method as compile-time-capable. The function remains callable at runtime. A concrete implementation may apply a generic trait to fixed type or const arguments; each trait application retains its own method identity and compile-time capability. A compile-time call may call other `#[comptime]` free functions, the builtin operations listed below, or the `compile_error(string) -> never` intrinsic. The compiler validates the complete body of every marked function, including branches not taken by a particular invocation. Attributes with arguments, duplicate attributes, generic functions or implementations, extern functions, and other declarations are rejected.

A top-level constant initializer is an implicit compile-time context, so `const SIX: isize = factorial(3);` and an initializer wrapped in `comptime { ... }` are equivalent. A top-level constant may produce a recursively immutable tuple, fixed array, struct, or enum in addition to scalar values. A `comptime` expression in ordinary code supports the same reifiable value shapes.

Compile-time code may use local bindings and assignment, blocks, `if`, `match`, `while`, `loop`, restricted `for`, `break`, `continue`, `return`, recursion, direct calls, integer conversion methods, integer `to_string()`, and supported operators. A compile-time `for` accepts only a fixed array or the builtin `isize` ranges `start..end` and `start..=end`; its source and range endpoints are evaluated once, and its pattern must be irrefutable. The deterministic string methods `len`, `byte_len`, `get`, `byte_get`, `byte_slice`, `is_char_boundary`, `starts_with`, `ends_with`, and `contains` are also available. String indexes and slices use byte offsets and reject invalid UTF-8 character boundaries. Integer formatting accepts signed and unsigned 8-, 16-, 32-, and 64-bit values, including `byte`, `isize`, and `usize`, and produces decimal text with the full range preserved. This is restricted to the builtin `ToString` implementation; a user trait with the same short name does not become a compile-time intrinsic. For example, `const BITS: u64 = 18446744073709551615; const MASK: string = BITS.to_string();` evaluates without runtime formatting.

Compile-time code cannot capture a surrounding runtime parameter or local. Closures, function parameters and indirect calls are supported within one evaluation, including shared mutable captures. Concrete `#[comptime]` methods support static calls and dynamic dispatch through local implementations. Their method values can be stored, passed, and called inside one evaluation, using the same capability checks and trait selection as direct calls. Dynamic dispatch retains all additional and inherited trait bounds, including their type arguments. Same-name methods remain distinct: child methods take precedence for method syntax, and UFCS selects the named declaring trait. Every component implementation must be local and its methods must be marked `#[comptime]`. `Ref` supports `new`, `get`, and `set`; `Vec` supports `new`, `with_capacity`, `from_array`, `get`, `set`, `push`, `len`, `is_empty`, `clear`, and `truncate`. These operations retain ordinary aliasing and bounds checks. Concrete instances of generic structs and enums substitute their field and payload types recursively, including nested references, vectors, callbacks, and dynamic trait arguments; generic compile-time functions and implementations remain unsupported. Function values, dynamic objects, references, and vectors cannot escape as runtime constants. Generic compile-time functions and implementations, general iterators, floating-point computation, `HashMap`, channels, goroutines, extern calls, host I/O, environment access, time, randomness, network access, general type reflection, arbitrary declaration generation, and general type-level computation are not supported. The constrained programmable derive interface described below is the only reflection and code-generation facility.

Compile-time value types must have finite structural layouts. Recursive type layouts, including recursion through `Ref`, `Vec`, or function signatures, are not supported and produce a compile error. Finite nested instantiations of the same generic type are supported. Type expansion also has deterministic depth and visited-node limits; exceeding either limit produces a diagnostic instead of terminating the compiler.

`compile_error` is accepted only in a `#[comptime]` function, a `comptime` block, or a top-level constant initializer. It terminates compile-time evaluation with its message. If runtime execution of a `#[comptime]` function reaches it, the program traps:

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

Failures include the compile-time call stack and source locations. Evaluation uses deterministic instruction, call-depth, temporary-value-node, temporary-memory, and final-value-size limits. Exceeding a limit is a compile error; no wall-clock timeout participates in the language semantics. Successful direct calls are memoized within one evaluation. Memoization is not observable because compile-time code has no user-visible side effects.

Public `#[comptime]` functions can be called from another package. The defining package's interface contains verified compile-time IR for the public entry and the private compile-time helpers and constants it reaches. Public compile-time constant values are also exported. A compile-time body or value change affects the interface hash; source formatting, local names, and source locations do not. A downstream package needs only the dependency interface for checking and evaluation. A value containing hidden fields from another package cannot currently be reified.

Compile-time integer evaluation uses the same fixed-width, wrapping representation as generated runtime code. Division by zero, a negative shift count, and an out-of-bounds index fail compilation. Signed minimum divided by `-1` yields the signed minimum, and its remainder is zero. Narrowing conversions retain the low bits of the destination width; widening a signed source sign-extends before conversion to the destination signedness. The CTIR target specification is part of its semantic hash. `isize` and `usize` currently use the 64-bit Linux amd64 toolchain target width.

Compile-time indexing with `value[index]` supports fixed arrays and vectors. It evaluates the receiver before the index, exactly once each. Vector reads observe mutations through aliases. Invalid indexes produce compile-time diagnostics.

At runtime, integer addition, subtraction, and multiplication use the same fixed-width wrapping representation. Integer division or remainder by zero terminates the current process with the generated Go runtime failure. Integer conversion methods use the same narrowing, sign-extension, and signedness rules as compile-time evaluation.

<a id="attributes-and-derive"></a>

## Attributes and derive

User source code supports deriving `ToString`, `Debug`, `PartialEq`, `Eq`, `PartialOrd`, `Ord`, `Hash`, and `Default` for structures and enumerations, including generic types. `Serialize` and `Deserialize` are also available after importing them from `std::serde` or a re-exporting format package:

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

All fields or variant payloads used by a derive must support that trait. A `ToString` variant template only uses the fields referenced by its placeholders. For a generic definition, the generated impl constrains each distinct participating field or payload type that mentions a type parameter. An unused phantom parameter receives no constraint. Compiler-owned runtime, intrinsic, and lang-item attributes remain unavailable to ordinary projects. User projects may use the `go_ffi`, `go_type`, `go_method` and `go_interface` attributes described in [Go FFI](../ffi/bind-go.md#go-ffi).

The eight prelude derives are supplied by verified handlers in the toolchain's builtin sources. They are not hard-coded code generators in the compiler. The `PartialOrd` and `Ord` traits themselves live in `std::cmp`; import that package when using their derives, and import the traits for unqualified bounds and method calls. Standard-library and third-party derives use the same handler and artifact mechanism, but they are not added to the prelude: import the trait or its package before using the derive name.

Derived `ToString` formats an enum with its type name, variant name, and payload by default. A variant can replace its entire output with a string template:

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

`Status::Ready.to_string()` returns `"ready"`, `Status::Retry(3).to_string()` returns `"retry after 3s"`, and `Status::Unknown.to_string()` retains `"Status::Unknown"`. Named payloads use `{field}`; tuple payloads use zero-based canonical decimal indexes such as `{0}`. `{{` and `}}` produce literal braces. Ordinary and raw string literals are accepted, including empty strings and Unicode text. A template may omit any or all payload fields. Referenced fields are converted once each, in declaration order, and repeated placeholders reuse the converted string. Only referenced fields require `ToString`; generic predicates also retain the requirements of every uncustomized variant. String payloads are inserted directly.

A `to_string` attribute requires exactly one positional string literal and may occur once per variant. Duplicate attributes, invalid arguments, unmatched braces, unknown fields, and out-of-range tuple indexes are compile-time errors located at the attribute or its template argument. When deriving `ToString`, placing this attribute on an enum, struct, or payload field is an error. Templates do not support implicit `{}` placeholders, expressions, format specifiers, callbacks, or automatic case conversion. They are interpreted by the `ToString` derive, independently of `Debug` and `serde(rename)`, and do not generate string parsing or require unique outputs. More complex formatting uses a handwritten `impl ToString`.

Derived `Debug` provides `Debug::debug(value)` and the `.debug()` method. It formats structs with their type and field names and enums with their type and variant names. Primitive field values use their ordinary textual representation, while nested values recursively use `Debug`.

`PartialEq` compares struct fields in declaration order and compares enum variants and payloads structurally. `Eq` generates only the marker implementation and does not implicitly derive `PartialEq`. The usual total-equivalence spelling is therefore `#[derive(PartialEq, Eq)]`.

`PartialOrd` and `Ord` use lexicographic field order. Enum variants are ordered by declaration order, followed by their payload fields. `PartialOrd` immediately propagates `None`; `Ord` returns a definite `Ordering`. Neither derive implicitly generates its supertraits, so the usual total-order spelling is `#[derive(PartialEq, Eq, PartialOrd, Ord)]`.

`Hash` combines fields in declaration order and includes the enum variant discriminant. It can be derived independently, although hash-map keys still require both `Eq` and `Hash`.

`Default` initializes every struct field with `Default::default()`. For an enum it always selects the first declared variant and recursively defaults that variant's tuple or named payload. An empty enum cannot derive `Default`. Reordering enum variants therefore changes both the derived default value and the derived ordering; `#[default]` is not supported.

<a id="programmable-derive"></a>

### Programmable derive

A dependency package may export a compile-time derive handler:

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

Another package applies the handler through an imported package path:

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

A public `#[comptime_derive(Name)]` handler exports the derive name `Name`, independently of its implementation function name. Export names are unique within one package and cannot contain `::`. Private named handlers can be used in their own package; `pub` is required to import a handler from another package. Private unnamed `#[comptime_derive]` functions can remain implementation helpers. The unnamed public form remains available and exports the handler function's short name.

Derive names have their own namespace, but ordinary `use` declarations populate it alongside the type and trait namespaces. For example, after `use std::serde;`, `use serde::Serialize;` makes both the `Serialize` trait and a derive export named `Serialize` available as `Serialize`, so `#[derive(Serialize)]` works without another import form. An import alias applies to both namespaces: `use serde::Serialize as DataSerialize;` permits `#[derive(DataSerialize)]`. A package import permits the qualified spelling `#[derive(serde::Serialize)]`. Public re-exports preserve the derive handler's definition identity, so facade packages can use `pub use` to expose a standard-library or third-party derive without copying its implementation.

Handlers can also be defined and used in the same package, including across source files. The compiler prepares the package's compile-time functions before expanding derives. The preparation pass cannot use methods or types that depend on that expansion; such cycles produce a diagnostic. Local derive names take precedence over imported derive names.

The generation API includes:

| Capability | Builders |
| --- | --- |
| Closures and function types | `meta_expr_closure(params, body)`, `meta_type_function(params, result)`, `meta_expr_invoke(function, args)` |
| Qualified generic member calls | `meta_expr_type_member(owner_type, member, member_type_args, args)` |
| Multiple implementations | `derive_output_merge(output, additional_output)` |
| Associated items | `derive_output_add_associated_type(output, name, type)`, `derive_output_add_constant(output, name, type, value)` |
| Trait parameters and equalities | `derive_output_set_trait_arguments(output, args)`, `derive_output_add_equality(output, lhs, rhs)` |
| Private helper declarations | `derive_output_add_function(output, method)`, `derive_output_add_struct(output, name, fields)`, `derive_output_add_type_alias(output, name, type)` |
| Structured attribute syntax | `meta_attribute_argument_type(attribute, index)`, `meta_attribute_argument_expr(attribute, index)` |
| Source locations | `derive_span(input)`, `derive_field_span(input, index)`, `derive_variant_field_span(input, variant, field)`, `meta_attribute_span(attribute)`, `meta_attribute_argument_span(attribute, index)` |
| Precise diagnostics | `derive_error_at(attribute, message)`, `derive_error_at_argument(attribute, index, message)`, `derive_error_span(span, message)` |

`meta_expr_closure` takes a `MetaParamList`; `meta_type_function` takes a `MetaTypeList`. Helper structs take a `MetaParamList` of field names and types. Use `derive_fresh_name` for helper names. Associated constants use the same expression restrictions as handwritten associated constants. `derive_set_span(span)` assigns the given location to subsequently generated syntax.

Structured attribute access preserves named and positional arguments, nested delimiters, and raw strings. Argument value kinds include `ident`, `path`, `string`, `bool`, `integer`, `float`, `char`, and `syntax`. Type and expression accessors validate a single type or expression and resolve names at the derive use site. For example, `#[config(priority = 10, extras = Option[State], callback = make_callback(2))]` has three arguments. Definition-site member builders qualify relative type owners such as `Rule::regex`; explicit call-site builders retain the deriving file's scope.

Static inherent methods can be used as function values: `let callback = Parser::parse;`. Their resolved signature supplies the function value's parameters and return type; generic method values still need enough context to infer their type arguments.

The standard library provides serialization derives through `std::serde`. See [Serde and JSON](../stdlib.md#serde-and-json) for their APIs, attributes, and examples.

A qualified custom derive must have the form `package_alias::export_name`, where `package_alias` is introduced by a `use` in the same source file. A bare library derive must be explicitly imported. Unimported canonical package paths and handlers visible only through a transitive dependency are rejected. If two explicit imports or a prelude derive and an explicit import provide the same bare name, the compiler reports ambiguity and requires a qualified name or import alias.

`#[comptime_derive]` and `#[comptime_derive(Name)]` are valid only on non-generic free functions. A public or named handler must have the exact signature `(DeriveInput) -> DeriveOutput`. Private functions with the unnamed attribute are compile-time-only helpers and are included in the interface when reachable from a public handler. A derive handler may call those helpers and ordinary `#[comptime]` functions, but it cannot be called from runtime code or ordinary value `comptime`. Derive handlers are not exported as runtime functions.

The compiler resolves imported handlers from already compiled dependency interfaces and prepares local handlers before expanding the current package. Reusable handlers and their generated traits may live in a separate package. The target may be a generic struct or enum; the generated impl inherits its type parameters. Explicit owner arguments also specialize static inherent methods whose parameter and return types do not mention those arguments, for example `Record::[string]::type_name()`. `derive_output_add_predicate` and `derive_output_add_call_site_predicate` add the bounds required by generated methods.

The input reflection operations are:

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

`derive_item_kind` returns zero for a struct and one for an enum. `derive_variant_kind` returns zero for a unit variant, one for a tuple variant, and two for a struct-like variant. The count operation for a nested attribute takes the owner indexes; its name and text operations take one additional attribute index. Field operations require a struct, while variant operations require an enum. Invalid kinds and indexes are compile-time errors at the `#[derive(...)]` site. Attributes may be attached to struct fields, enum variants, and tuple or named variant fields. `name` returns the attribute name and `text` returns its complete source spelling.

The structured attribute handle exposes `meta_attribute_name`, `meta_attribute_text`, `meta_attribute_has_argument_list`, `meta_attribute_argument_count`, `meta_attribute_argument_kind`, `meta_attribute_argument_name`, `meta_attribute_argument_value_kind`, and `meta_attribute_argument_text`. Argument kind is a literal or syntax kind listed above, or `named`. Named arguments use `name = value`, where the value may also be a number, boolean, nested type, or expression. `argument_name` returns the left-hand name, `argument_value_kind` describes the right-hand value, and `argument_text` returns its decoded value.

The structured output API provides opaque `MetaSpan`, `MetaAttribute`, `MetaType`, `MetaExpr`, `MetaPattern`, `MetaArm`, `MetaBlock`, `MetaParamList`, `MetaGenericList`, `MetaMethod`, and list handles. Constructors use the `meta_type_*`, `meta_expr_*`, `meta_pattern_*`, `meta_arm*`, `meta_block_*`, `meta_param_list_*`, and `meta_generic_list_*` families. `meta_method` creates a concrete method; `meta_method_generic` creates a method with explicit type parameters and bounds. A handler starts with one trait impl with `derive_output_new` or `derive_output_new_call_site`, or one inherent impl with `derive_output_inherent`. It then adds predicates and methods and returns that output. `derive_output_add_method` preserves private visibility for inherent methods and normal trait visibility for trait implementations. `derive_output_add_public_method` exports an inherent method and rejects trait output.

For example, a handler in a separate package can generate a public static method:

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

The builder operations are:

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

`meta_expr_unary` uses operator numbers `0..2` for `-`, `!`, and `~`. `meta_expr_binary` uses operator numbers `0..17` for `+`, `-`, `*`, `/`, `%`, `&`, `|`, `^`, `<<`, `>>`, `&&`, `||`, `<`, `>`, `<=`, `>=`, `==`, and `!=`, respectively. `meta_expr_integer` accepts a normalized integer literal string plus its exact integer type, so builders can represent values outside the host `isize` range. `meta_expr_cast` accepts only a `dyn` target type; generated numeric conversions should use `meta_expr_method_call`. `meta_expr_trait_call` resolves the trait in the handler's defining package and builds a static trait method call. `meta_type_equal` compares structural type identity. `meta_type_kind` returns `primitive`, `named`, `tuple`, `application`, `array`, `function`, or `dyn`; shape-specific accessors reject other kinds. List handles are mutable only through their matching `push` operation and remain local to one derive evaluation.

Unqualified names passed to `derive_output_new`, `derive_output_add_predicate`, `meta_type_named`, `meta_expr_call`, and `meta_generic_list_add_bound` resolve in the handler's defining package. Their `_call_site` variants resolve in the target package. `derive_fresh_name` should be used for generated local bindings that must not collide with user names. The `*_target_*` builders construct or match the annotated item by compiler identity and should be preferred over spelling its name manually.

The result may contain several trait or inherent implementations for the annotated type, associated types and constants, and private helper functions, structs, and aliases. It cannot create traits, modules, imports, extern declarations, attributes, or raw tokens. Generated method type parameters and trait bounds are supported. The generated impl is processed by ordinary name resolution, orphan and coherence checks, type checking, monomorphization, and backend lowering. Duplicate or invalid generated implementations are regular compiler diagnostics.

Handlers are deterministic and have no host access. Imported derive CTIR is verified as untrusted artifact data. Evaluation uses the ordinary compile-time limits plus a limit of 100,000 metadata and syntax-builder operations. Failures default to the requesting derive attribute; span-aware error operations can select a field or attribute argument. Diagnostics include the compile-time derive call stack.
