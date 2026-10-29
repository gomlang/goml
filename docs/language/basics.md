<a id="language-basics"></a>

# Language basics

[English](basics.md) | [简体中文](basics.zh-CN.md)

[GoML language guide](../goml.md)

Lexical rules, definitions, functions, bindings, expressions, control flow, and closures. See [types and traits](types.md) for data definitions, generics, and pattern rules.

On this page: [Lexical rules](#lexical-rules), [Top-level definitions](#top-level-definitions), [Top-level functions](#top-level-functions), [Blocks, bindings and assignments](#blocks-bindings-and-assignments), [Expressions and operators](#expressions-and-operators), [Control flow](#control-flow), [Closures and function values](#closures-and-function-values).

<a id="lexical-rules"></a>

## Lexical rules

<a id="comments-and-api-documentation"></a>

### Comments and API documentation

Line comments start with `//` and end at the next newline; block comments are unsupported. Standalone `///` lines (excluding `////`) document the following declaration; same-line trailing comments do not. Consecutive documentation lines form one block; a blank line or an ordinary comment breaks attachment. Attributes belong to the following declaration and may appear between its documentation and declaration keywords.

`//!` comments before the first source token describe the package. For a package with several source files, `gomldoc` combines these descriptions in relative file-path order. Documentation comments remain ordinary lexer trivia and do not affect program typing or execution.

```goml
//! Small arithmetic helpers.
package arithmetic;

/// Add one to `value`.
pub fn increment(value: isize) -> isize {
    value + 1
}
```

`goml doc` generates API documentation from these comments. See [API documentation](../documentation.md) for commands, output formats, visibility, supported Markdown, and API links.

<a id="identifiers-and-keywords"></a>

### Identifiers and keywords

Ordinary identifiers only use ASCII and are of the form:

```text
[A-Za-z][A-Za-z_0-9]* | _[A-Za-z_0-9]+
```

The single `_` is a wildcard character, not an ordinary variable name. The current syntax enforces the first letter case of name categories:

- Package names, package aliases in `use package as alias`, functions, methods, parameters, local bindings and GoML field declarations must start with a lowercase letter or `_`; imported item aliases may follow the naming convention of the imported item;
- Structures, enumerations, traits, enumeration variants, generic parameters, and associated types must start with a capital letter;
- The fixed-width vector and mask type spellings listed under [Portable SIMD](../stdlib.md#portable-simd), including `u8x16`, `i32x8`, `f32x4`, `f64x4`, and `mask32x4`, are also accepted as structure and type-alias names and in type positions. They are ordinary scoped names, supplied by `std::simd`, rather than globally reserved keywords;
- Paths retain the appropriate case for the referenced name.

Prefer unqualified enumeration variants in construction and patterns when their expected type determines the variant owner. Use `Enum::Variant` when there is insufficient type context or ordinary name resolution selects another declaration. Prefer receiver methods such as `value.to_string()` for operations on an existing value, and indexing such as `values[i]` for ordinary vector element access.

Common keywords include:

```text
package use as pub fn struct enum trait impl for type const static where defer
let mut if else match while loop for in break continue return go dyn
true false bool isize i8 i16 i32 i64 usize u8 u16 u32 u64
f32 f64 string char extern
```

`()` is the only spelling of the empty-tuple type.

`self` is a special abbreviation for the receiver parameter, and can also be used as a common receiver variable name; `Self` has special meaning in the type position of trait and impl.

Whitespace separates tokens but does not determine block structure.

<a id="literals"></a>

### Literals

| category | Syntax example | Default or description |
| --- | --- | --- |
| empty tuple | `()` | Type is `()` |
| bool | `true`, `false` | Type is `bool` |
| integer | `0`, `42`, `1_000`, `0b1010`, `0o755`, `0xff` | Determined by context; defaults to `isize` when unconstrained |
| floating point number | `1.25`, `1e3`, `2.5e-2` | Determined by context; defaults to `f64` when unconstrained |
| string | `"text"` | Type is `string` |
| raw string | `r"text"`, `r#"text with \"quotes\""#` | Type is `string`; escapes are not processed |
| byte string | `b"text\\n"` | Type is `Vec[byte]`; ASCII contents and byte escapes are supported |
| raw byte string | `br"text"`, `br#"text with "quotes""#` | Type is `Vec[byte]`; escapes are not processed |
| interpolated string | `f"value={value}"` | Type is `string`; embedded values use `ToString` |
| character | `'a'`, `'\n'`, `'\u0041'` | Type is `char`, representing a Unicode scalar value |
| byte | `b'A'`, `b'\n'`, `b'\xFF'` | Type is `byte`, a transparent alias of `u8` |

Numbers have no type suffix. Integer literals support binary `0b`/`0B`, octal `0o`/`0O`, decimal, and hexadecimal `0x`/`0X` forms. The `_` delimiter can be used between two digits; floating point numbers support `e`/`E` exponent and optional exponent sign. When using a decimal point, there must be digits on both sides of the decimal point. Negative numbers are composed of unary `-` and positive numeric literals.

Floating-point literals are rounded directly from their source text to the context-selected `f32` or `f64` width using round-to-nearest, ties-to-even. Overflowing literals are rejected; underflow may round to zero.

Unsuffixed numbers can get the width from the context:

Integer and floating-point literals can share an inferred numeric type. Without a concrete numeric context, a group containing a floating-point literal defaults to `f64`, so `[1, 2.5]` has type `[f64; 2]`. A concrete integer context still requires every literal to be exactly representable in that integer type.

```goml
let small: u8 = 42;
let ratio: f32 = 0.5;
let values: [i16; 3] = [1, 2, 3];
```

Strings support `\"`, `\\`, `\n`, `\r`, `\t`, `\b`, `\f`, `\/` and four-digit `\uXXXX` escaping; characters are escaped using the same set of control characters, and `\'` is used to represent single quotes. Ordinary strings cannot span lines.

Escaped Unicode characters retain their UTF-8 value through Go code generation. For example, `"\uFEFF"` contains one character and three bytes, including when embedded in another string.

Raw strings use `r"..."` or matching hash delimiters such as `r#"..."#` and `r##"..."##`. Their contents may span lines, and backslashes, quotes, braces, and newlines are retained exactly. Escapes and interpolation are not processed. A quote closes the literal only when it is followed by the same number of `#` characters as the opening delimiter:

```goml
let windows_path = r"C:\users\alice";
let quoted = r#"She said "hello"."#;
let multiline = r##"first line
second line"##;
```

Interpolated strings use `f"..."`. Each `{expression}` is evaluated once from left to right and converted with `ToString`; a value that is already `string` is inserted directly. Write `{{` and `}}` for literal braces. Formatting specifications and interpolated multiline strings are not supported:

```goml
let name = "Ada";
let count = 3;
let message = f"{name} has {count} items {{ready}}";
```

Byte literals contain exactly one ASCII byte and support `\'`, `\\`, `\0`, `\b`, `\f`, `\n`, `\r`, `\t`, and two-digit `\xNN` escapes. Non-ASCII contents are rejected.

Byte strings use `b"..."` and have type `Vec[byte]` without an expected type. In a `[byte; N]` context they produce a fixed array and the length must match. Vector byte strings compile to one backing-slice allocation with batch literal initialization rather than per-byte `push` calls. Ordinary byte strings accept ASCII contents together with `\"`, `\\`, `\0`, `\b`, `\f`, `\n`, `\r`, `\t`, `\/`, and two-digit `\xNN` escapes. Unicode escapes and non-ASCII source characters are rejected:

```goml
let request = b"GET / HTTP/1.0\r\n\r\n";
let marker = b"\x00\xFF";
```

Raw byte strings use `br"..."` or matching hash delimiters such as `br#"..."#`. Their contents may span lines and are copied as ASCII bytes without escapes or interpolation:

```goml
let path = br"C:\tmp\file";
let quoted = br#"say "hello""#;
```

The standard library type `std::bytes::Bytes` remains a separate buffer abstraction. It provides read-only and mutable zero-copy views plus a growable byte builder. `std::bytes::endian` provides checked endian access and matching growable and fixed-buffer readers and writers. Use `bytes::Bytes::from_vec(value)` when an API requires it.

Each line of a multiline string begins with two backslashes; the indentation before the mark is removed, the lines are connected with newlines, and the content after the mark is retained as is. At the end, the next line no longer starts with two backslashes, usually written directly `;` or `}`. The formatter places the multiline string on its own indented lines:

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

## Top-level definitions

The top level of an ordinary source code file should only contain:

- `fn`
- `struct`
- `enum`
- `trait`
- `impl`
- `type`
- `const`
- `static`

`package` and `use` can only appear at the beginning of a file. The top level cannot write local variables or arbitrary execution statements.

<a id="constants"></a>

### Constants

Top-level constants require an explicit type and a compile-time expression:

```goml
const BASE: i32 = 0x20;

pub const ANSWER: i32 = BASE + 10;

const NEWLINE: byte = b'\n';
```

Constant names may start with an uppercase letter, a lowercase letter, or `_`; `UPPER_SNAKE_CASE` is the preferred convention. Constant expressions support literals, constructors, references to other constants, unary and binary operators, integer conversion methods, and comptime-capable calls. They may use forward references, but cycles are rejected. Allowed constant types are recursively immutable values: `bool`, numeric types, `string`, `char`, `byte`, tuples (including `()`), fixed arrays, and structs or enums whose fields satisfy the same rule. `Vec`, `HashMap`, `Ref`, `Channel`, closures, and dynamic values are rejected. Composite constants cannot be used as patterns. Finite nesting of immutable generic types, such as `Wrap[Wrap[i32]]`, is allowed; repeating a generic type name is not itself recursive. Public constants are available through package-qualified paths.

Names in a constant initializer resolve where the constant is declared, using that file's imports and the enclosing declaration's generic parameters. Referencing the constant from another file or generic function does not reinterpret its initializer in the caller's scope. In particular, a top-level constant cannot capture a generic parameter of a function that uses it:

```goml
const SIZE: usize = 2;
const WIDTH: usize = SIZE;

fn first[const SIZE: usize](values: [isize; WIDTH]) -> isize {
    values[0]
}
```

`first::[3]([10, 20])` accepts a two-element array because `WIDTH` remains `2`.

Visible top-level constants can be used as value patterns in `match`, `if let`, `while let`, and `let else`. Both local names and package-qualified names are supported, and constant patterns can be nested or combined with or-patterns:

```goml
const ANSWER: isize = 42;

match value {
    ANSWER | config::FALLBACK => "known",
    _ => "other",
}
```

In these refutable pattern contexts, a visible constant takes precedence over introducing a binding with the same name. Use an explicit alias such as `answer @ _` to bind that name instead. Ordinary `let` bindings and `for` patterns continue to introduce bindings, even when a top-level constant has the same name. Range-pattern endpoints remain integer or character literals.

<a id="statics-and-one-time-initialization"></a>

### Statics and one-time initialization

A top-level `static` has one program-wide storage location. Its binding cannot be reassigned. Its explicit type must be `OnceCell[T]` and its initializer must be exactly `OnceCell::new()`:

```goml
static EMOJI: OnceCell[FrozenVec[FrozenVec[u16]]] = OnceCell::new();

fn emoji() -> FrozenVec[FrozenVec[u16]] {
    EMOJI.get_or_init(|| build_emoji().freeze())
}
```

`OnceCell.get_or_init(init)` runs one initializer and caches its result. Concurrent first callers wait for that initializer and receive the same value. If the initializer panics, the cell becomes uninitialized again and waiting callers are woken so a later initializer can retry. Recursive initialization of the same cell panics with an error that names the static. A cached `Result` is an ordinary cached value. Statics do not run user-observable destruction at process exit.

Unlike a `const`, a `static` has observable identity. `OnceCell[T]` controls initialization but does not make `T` immutable. Shared caches should therefore expose `FrozenVec` or another immutable value instead of a mutable `Vec`.

<a id="top-level-functions"></a>

## Top-level functions

```goml
fn add(left: i32, right: i32) -> i32 {
    left + right
}

fn log(message: string) {
    println(message)
}
```

The parameter type cannot be omitted. Omitting `-> ...` is equivalent to `-> ()`. The top-level function name must be unique in the same package, and overloading by parameter type is not supported.

<a id="blocks-bindings-and-assignments"></a>

## Blocks, bindings and assignments

<a id="block-values-and-semicolons"></a>

### Block values and semicolons

The last semicolon-less expression of the block is the return value:

```goml
fn square(value: i32) -> i32 {
    let result = value * value;
    result
}
```

A block without a tail expression returns `()`. Adding `;` after an expression turns it into an expression statement and discards the result:

```goml
fn run() -> () {
    println("first");
    1 + 2;
}
```

`let`, ordinary assignments, and general non-tail expression statements require semicolons. When used as statements and followed by code, `if`, `match`, `select`, and unlabeled `while`, `loop` and `for` can omit the semicolon; other expression statements still require semicolons. Functions, brace-delimited structures, enumerations, traits, impl and blocks themselves are not declared with a semicolon after them. Newtype declarations such as `struct UserId(u64);` do require the trailing semicolon.

`defer expression;` registers a `()` expression to run when the current lexical block is left. Deferred expressions run in last-in-first-out order on normal completion and when `return`, `?`, `break`, or `continue` crosses their block. A return or break value is evaluated before cleanup begins. Each loop-body block has its own cleanup stack, so a deferred expression registered during one iteration runs before that iteration exits.

Unlike Go's `defer`, GoML does not evaluate call arguments when the statement is reached. The complete expression is evaluated at block exit, so reads through `Ref` and captured mutable locals observe the value at cleanup time. A closure body is a separate control-flow scope. Deferred expressions cannot contain `return`, `break`, `continue`, or `?`. Panic unwinding also runs registered cleanup in last-in-first-out order. A cleanup is removed before invocation, so it runs exactly once even if it panics; remaining cleanups still run and the newest panic propagates. The compiler uses a function-local cleanup stack with explicit pops at lexical exits and a Go `defer` guard for unwinding. Loop iterations release their completed entries instead of accumulating deferred calls until function return. Abrupt process termination and fatal runtime failures do not guarantee cleanup.

```goml
fn work() -> () {
    let state = Ref::new("open");
    defer println("closing:" + state.get());
    defer println("flush");
    state.set("ready")
}
```

This prints `flush` first, then `closing:ready`.

<a id="let-type-annotations-and-shadowing"></a>

### `let`, type annotations and shadowing

```goml
let inferred = 42;
let explicit: i64 = 42;
let _ = println("discard explicitly");
let name = "first";
let name = "second";
```

Local variables allow shadowing with the same name. `let _ = expr;` executes explicitly and discards the result.

Each block has a lexical binding scope, including ordinary expression blocks, `comptime` blocks, and task `scope` blocks. A binding is visible after its initializer through the end of its block, including nested blocks and closures. Leaving a block restores any outer binding with the same name; newly introduced names are unavailable outside that block. Assignment to an existing outer mutable binding still changes that binding:

```goml
let value = 1;
let inner = {
    let value = 2;
    value
};
println(inner);
println(value);
```

This prints `2` and then `1`.

The left side of `let` can be an irrefutable pattern:

```goml
let (left, right) = pair;
let Point { x, y: vertical } = point;
```

An ordinary `let` cannot use an enumeration or literal pattern that may fail. Use `let ... else` when the failure path must leave the surrounding control flow:

```goml
fn require_value(value: Option[isize]) -> isize {
    let Some(item) = value else {
        return -1;
    };
    item
}
```

The initializer is evaluated once. The pattern must be refutable, bindings become visible after the statement, and the `else` block cannot use those bindings. The `else` block must diverge with `return`, `break`, `continue`, an infinite `loop`, or another expression of type `never`. Use `if let` or `match` when both paths continue locally.

<a id="mut-and-assignment"></a>

### `mut` and assignment

```goml
let mut count = 0;
count = count + 1;
```

Destructuring patterns can mark only selected bindings mutable:

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

`mut` inside a pattern is allowed only on a binding and does not introduce borrowing or reference semantics. The existing `let mut pattern = value;` form remains supported and recursively makes every binding in that pattern mutable. Per-binding `mut` composes with tuple, struct, enum, array, `match`, and `for` patterns.

Ordinary local bindings without `mut` cannot be reassigned. Assignment targets include mutable locals, tuple projections, structure fields, and supported index locations. Literals, computed values, function definitions, and method values are not assignable locations. Compound assignment supports `+=`, `-=`, `*=`, `/=`, `%=`, `&=`, `|=`, `^=`, `<<=`, and `>>=`:

```goml
count += 1;
point.x *= 2;
values[index] <<= 1;
```

The target root and each index are evaluated once before the right-hand side. `++` and `--` are not supported. Compound assignment through `HashMap` indexing is not supported because an indexed read returns `Option[V]`; use `get` and `set`, or ordinary indexed assignment.

Array index assignment requires that the root value is a `let mut` local array, or comes from the built-in `Ref.get()`:

```goml
let mut values = [1, 2, 3];
values[1] = 20;
let shared = Ref::new([4, 5]);
shared.get()[0] = 40;
```

`Vec` and `HashMap` are internal mutable containers, so the contents can be modified via methods or indexes even if the binding itself is not `mut`:

```goml
let values: Vec[i32] = Vec::new();
values.push(10);
values[0] = 20;
let counts: HashMap[string, i32] = HashMap::new();
counts["answer"] = 42;
```

`Slice[T]` is a read-only view and cannot be assigned by index, including through fields, tuple projections, or array indexes within an element. Interior mutation through an element's `Ref`, `Vec`, or other mutable container remains available. `MutSlice[T]` is a mutable bounded view and supports indexed assignment without copying its `Vec[T]` backing storage.

<a id="expressions-and-operators"></a>

## Expressions and operators

<a id="basic-expressions"></a>

### Basic expressions

GoML supports:

- Literals, variables and paths
- Tuple, array, and structure literals
- Block expressions `{ ... }`
- Function call `f(a, b)`
- Field access `value.field`
- Tuple projection `pair.0`, `triple.2`
- Method call `value.method(arg)`
- Index `value[index]`
- Unary and binary operations
- Integer conversion method `value.to_u32()`
- Explicit trait-object conversion `value as dyn Trait`
- Half-open and inclusive range expressions `start..end` and `start..=end`
- `if`, `if let`, `match`, `select`, `while`, `while let`, `loop`, `for`
- closure
- `return`, `break`, `continue`, `go` and `?`

A block is an expression in any ordinary expression position. Its last expression without a semicolon supplies the value; a block without a tail expression has type `()`:

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

In an `if`, `if let`, `while`, `while let`, `for`, or `match` header, the top-level `{` starts the control-flow body. Parenthesize a block expression used as the header value:

```goml
if ({
    prepare();
    ready()
}) {
    run()
}
```

<a id="operator-precedence"></a>

### Operator precedence

From low to high:

| Precedence | Operator | Description |
| --- | --- | --- |
| 1 | `..`, `..=` | Half-open or inclusive range; cannot be chained |
| 2 | `\|\|` | Short-circuit logical OR |
| 3 | `&&` | short circuit logical AND |
| 4 | `==`, `!=`, `<`, `>`, `<=`, `>=` | Compare; not chainable |
| 5 | `\|` | Bitwise OR |
| 6 | `^` | Bitwise XOR |
| 7 | `&` | Bitwise AND |
| 8 | `<<`, `>>` | shift |
| 9 | `+`, `-` | Addition, subtraction, string concatenation |
| 10 | `*`, `/`, `%` | Multiplication, division, remainder |
| 11 | `as` | explicit `dyn Trait` conversion |
| 12 | Unary `-`, `!`, `~` | prefix |
| 13 | Call `()`, index `[]` | suffix |
| 14 | `?` | error or absence propagation |
| 15 | Member `.` | field access, tuple projection, or method selection |

The binary operator is left associative, and the function type `->` is right associative. Don’t write comparisons in chains; use combinations of logical operations:

```goml
let inside = lower <= value && value < upper;
```

Calls, indexing, `?`, and member access bind more tightly than unary operators. For example, `-compute()` negates the call result and `!values[0]` negates the indexed value:

```goml
let negative = -compute();
let valid = !values[0];
```

<a id="operator-type-rules"></a>

### Operator type rules

- The two numeric operands of `+ - * /` must be of the same concrete numeric type; there is no implicit numeric promotion. Nominal structs and enums, and unconstrained generic parameters, do not acquire arithmetic operations from their representation or eventual instantiation. These rules also apply to compound assignments.
- `%` only accepts integers of the same type.
- `& | ^ ~` only accepts integers, and both sides of binary bitwise operations must be of the same type. Both sides of `<< >>` must be integers and the result type is the same as the left operand.
- `+` also supports `string + string`.
- `&& || !` only accepts `bool`.
- The unary `-` only accepts signed integers or floating point numbers.
- Both sides of the comparison must be of the same type.
- `< > <= >=` uses `PartialOrd`. Primitive numbers, `string`, `char`, tuples, fixed arrays, `Vec`, `Slice`, `Option`, and `Result` provide the corresponding conditional implementations.
- `== !=` uses `PartialEq`. Tuples, fixed arrays, `Vec`, `Slice`, `Option`, and `Result` compare recursively when their elements implement `PartialEq`.
- Floating-point values implement `PartialEq + PartialOrd`, but not `Eq + Ord + Hash`. In particular, NaN is unequal to itself and its `partial_cmp` result is `None`.
- Every integer type provides `to_isize`, `to_i8`, `to_i16`, `to_i32`, `to_i64`, `to_usize`, `to_u8`, `to_u16`, `to_u32`, and `to_u64` for the other integer types. Identity methods are omitted. `byte` uses the `u8` methods, `char` provides `to_u32`, and `u32` to `char` uses `char_from_u32`, which returns `Option[char]`.
- `f32` provides `to_bits` and `to_f64`; `f64` provides `to_bits` and `to_f32`. The prelude functions `float32_from_bits` and `float64_from_bits` reconstruct values from their IEEE 754 bit patterns. The `f64` to `f32` conversion uses round-to-nearest, ties-to-even; the reverse conversion is exact.
- Import `std::num::{ToFloat, TryToInt}` for scalar integer `.to_f32()` / `.to_f64()` and checked float `.try_to_i32()` / `.try_to_u64()` conversions, with analogous methods for every integer width. Float-to-integer methods truncate toward zero and reject nonfinite or out-of-range results.
- Expression `as` is reserved for creating `dyn Trait` values. Numeric conversions use the [integer and floating-point conversion methods](types.md#basic-types-and-composite-types); other non-`dyn` targets are rejected. The separate `as` form in `use package as alias;` still selects an import alias.

There are no exponentiation, null coalescing or user-defined operators.

<a id="range-expressions"></a>

### Range expressions

`start..end` constructs an incrementing half-open `FnIterator[isize]`, and `start..=end` constructs an inclusive iterator. Both can be used directly in `for`:

```goml
for value in 0..10 {
    println(value)
}

for value in 0..=10 {
    println(value)
}
```

Both ends are `isize`. A half-open range is empty when `start >= end`; an inclusive range is empty when `start > end` and contains one value when both ends are equal. Each endpoint is evaluated once from left to right. Inclusive iteration does not compute `end + 1`, so the maximum `isize` endpoint does not overflow. Range expressions cannot be chained. Open ranges, character ranges, and custom step syntax are not supported. The `..` and `..=` in patterns are a separate range-pattern syntax.

<a id="control-flow"></a>

## Control flow

<a id="if"></a>

### `if`

`if` is always an expression. `else` is required when generating values other than `()`:

```goml
fn absolute(value: i32) -> i32 {
    if value < 0 {
        -value
    } else {
        value
    }
}
```

Both branches must produce compatible types. Conditions with only side effects can omit `else`; in this case the then branch must produce `()`:

```goml
if enabled {
    println("enabled")
}
```

`else if` is another `if` in the `else` branch:

```goml
let label = if score > 90 {
    "high"
} else if score > 60 {
    "middle"
} else {
    "low"
};
```

`if let` executes the then branch when the pattern match is successful and restricts the pattern binding to that branch. The matched expression is evaluated only once:

```goml
let number = if let Some(value) = candidate {
    value
} else {
    0
};
```

The `else` of `if let` can also be omitted; in this case the compiler treats the else branch as `()`, so the then branch must produce `()`:

```goml
if let Some(value) = candidate {
    println(value);
}
```

When you need to get values from two branches, you must write `else` explicitly.

Conditions can chain boolean expressions and `let` matches with `&&`:

```goml
fn combine(left: Option[isize], right: Option[isize]) -> Option[isize] {
    if let Some(a) = left && a > 0 && let Some(b) = right && b > a {
        Some(a + b)
    } else {
        None
    }
}
```

Terms run from left to right and stop at the first false condition or failed pattern. Every reached initializer is evaluated once. A `let` binding is visible to later terms and the successful branch, but not to earlier terms, the `else` branch, or code after the `if`. A binding may shadow an outer name; `else` still resolves that name in the outer scope. Non-pattern terms must have type `bool`. A let chain cannot contain top-level `||`; parenthesize a boolean disjunction within one term, for example `if let Some(x) = value && (x == 1 || x == 2) { ... }`.

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

The matched expression is evaluated only once, and the pattern is tried from top to bottom. All branches must produce compatible types, and the compiler is required to cover all possible values.

An unqualified variant in a pattern is resolved only against the matched value's type. The compiler does not search unrelated enums for a fallback. If that type is not an enum, does not contain the variant, or cannot be inferred, the pattern is rejected. Fully qualified patterns remain available:

```goml
match value {
    Some(inner) => inner,
    None => fallback,
}
```

The rule is shared by `match`, `if let`, `while let`, `let`, `for`, nested patterns, and or-patterns. Existing refutability requirements still apply to `let` and `for`.

A match arm may have a guard of type `bool`. The guard can use bindings introduced by that arm’s pattern:

```goml
match value {
    Some(inner) if inner > 0 => inner,
    Some(_) => 0,
    None => -1,
}
```

A guard is evaluated only after a successful pattern match. Branches with ordinary guards are not counted in exhaustive coverage; only branches without guards and branches with guard literal `true` provide coverage. Branches with guard literal `false` are unreachable. The compiler also warns about branches that are completely covered by earlier patterns.

A branch body can be a single expression or a block. Different branches must be separated by commas, and the comma in the last branch can be omitted:

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

`select` waits for one channel operation that can proceed and evaluates that arm as the value of the expression:

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

A receive arm has the form `recv(channel) as binding => body`. The channel must have type `Channel[T]` or `Receiver[T]`, and the binding has type `Option[T]`; a closed and drained channel selects immediately and supplies `None`. Write `_` when the received value and close state are not needed. A send arm has the form `send(channel, value) => body`, accepts `Channel[T]` or `Sender[T]`, and requires the value to match the channel element type.

Communication arms may have a `when` guard after their operands. Operands and guards are evaluated exactly once in source order. A false guard disables its arm by selecting on a nil directional channel; if every arm is disabled, a select without `default` blocks and one with `default` runs immediately. The receive binding is not in scope in its guard.

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

`recv(channel) match { ... }` exhaustively matches the `Option[T]` result after communication succeeds. It is equivalent to receiving into a fresh binding and applying an ordinary `match`, so contextual variants, nested patterns, or-patterns, and exhaustiveness checking work normally. A refutable pattern cannot directly replace the receive binding because discarding a received value after a pattern failure would lose data invisibly.

`select priority` is a deterministic non-blocking form. It requires one final `default`, probes communication arms from first to last, chooses the first arm that can complete immediately, and otherwise runs `default`. False guards are skipped. Ordinary `select` retains Go's unspecified choice among simultaneously ready operations.

Every `select` contains at least one `recv` or `send` arm. It may have one `default` arm, which must be last. Without `default`, execution blocks until an operation can proceed. With `default`, that arm runs immediately when no communication arm is ready. When several operations are ready, the selected arm is unspecified.

Channel operands, send values, and guards are evaluated exactly once from top to bottom before selection. Arm bodies are evaluated only after their operation is chosen, receive bindings are visible only in their own bodies, and all bodies must produce compatible types. A body may be a single expression or a block. Arms are comma-separated, with an optional final comma. `select`, `priority`, `recv`, `send`, `when`, `match`, and `default` are contextual spellings in these positions rather than globally reserved identifiers.

`select` is a runtime operation and is unavailable in `comptime` evaluation. Cancellation, task completion, and timers participate through `Receiver[()]` values rather than dedicated select-arm syntax.

<a id="while"></a>

### `while`

```goml
let mut index = 0;
while index < limit {
    index = index + 1;
}
```

The condition must be `bool` and the loop result is `()`. A `while` loop accepts only `break` without a value.

`while let` evaluates the right-hand expression once in each round; when the match is successful, it enters the loop body and provides pattern binding, and when the match fails, it exits the loop:

```goml
while let Some(value) = iterator.next() {
    println(value);
}
```

The loop body must return `()`. A single `while let` binding is only visible within the loop body. `while` also supports the same short-circuit let chains as `if`:

```goml
while remaining > 0 && let Some(value) = iterator.next() && value > 0 {
    println(value);
    remaining -= 1;
}
```

Each iteration evaluates the chain again from its first term. A false boolean or failed pattern exits the loop. Pattern bindings are visible to later terms and the body of that iteration, and are recreated for the next iteration.

`break` and `continue` cannot directly exit a `while` condition, including a `while let` initializer or a term in a let chain. A loop nested inside the condition may use its own `break` and `continue` normally.

<a id="loop"></a>

### `loop`

`loop` repeats a block until control leaves it. Unlike `while` and `for`, it is an expression whose result comes from `break`:

```goml
let result = loop {
    let candidate = next();
    if candidate >= 0 {
        break candidate;
    }
};
```

All `break` values targeting the same `loop` must have compatible types. `break;` supplies `()`, so it cannot be mixed with other break value types. A `loop` with no `break` targeting it has type `never` and does not continue to the following expression. Unlabeled `break` and `continue` target the nearest loop.

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

`for pattern in source { ... }` accepts fixed arrays and values that implement `IntoIterator`. Both the source expression and the `into_iter` transformation are executed only once. The pattern must be irrefutable and the loop body must return `()`. `start..end` can be used directly as a native `isize` range.

Tuple destructuring can be used directly in loops:

```goml
for (key, value) in pairs {
    println(key + value.to_string());
}
```

Fixed arrays use a native indexed loop. `Vec[T]`, `Slice[T]`, `MutSlice[T]` and all `Iterator` values have corresponding `IntoIterator` implementations.

<a id="break-continue-and-return"></a>

### `break`, `continue` and `return`

```goml
while true {
    if done {
        break
    } else {
        continue
    };
}
```

`break` and `continue` can only appear in loops. `continue` never has a value. `break value` is allowed only for `loop`; `while` and `for` accept `break;` only. Both are divergent control expressions and can appear in `if`, `match`, or other value positions.

Rust-style labels select an enclosing `while`, `loop`, or `for` without conflicting with `break value`:

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

`continue 'label;` resumes the selected enclosing loop. `break 'label value;` may carry a value only when the selected target is a `loop`. Labels must refer to an active enclosing loop and cannot be duplicated while an outer label of the same name is active. Crossing one or more blocks still runs their `defer` cleanups from inner to outer before control reaches the target.

A labeled loop used as a statement requires a trailing semicolon before the next statement. The formatter preserves this semicolon. A labeled loop used as a block's tail expression does not need one.

`return` can be taken with or without a value and checks the return type of the current function or closure:

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

The `unreachable_code` lint warns when a statement or tail expression follows control flow that cannot continue, including `return`, `break`, `continue`, an infinite `loop`, or an exhaustive `if` or `match` whose branches all exit. The compiler reports one warning for each continuous unreachable region, continues checking that code for other errors, and does not fail the build for the warning alone. A function or method can suppress intentional unreachable code with `#[allow(unreachable_code)]`; closures inherit the setting of their enclosing function.

```goml
#[allow(unreachable_code)]
fn platform_exit() -> () {
    return;
    println("unreachable fallback")
}
```

The `unused_function` lint warns about private top-level functions that are not reachable from a public function, `main`, a test, a language item, a top-level initializer, or any method body. The analysis follows calls transitively, so a recursive function or a group of private functions that only call each other is still unused. Public functions are treated as externally reachable API. Use `#[allow(unused_function)]` on an intentionally unused private function to suppress the warning.

The `unused_result` lint warns when an expression statement directly discards a built-in `Result[T, E]`. The check uses the resolved type, including aliases, generic call results, and cross-package return types. It also checks statements in nested blocks and closures. Warnings do not fail the build.

```goml
fn save() -> Result[(), string] {
    Err("save failed")
}

fn main() -> () {
    save();
    let _ = save();
}
```

The first call warns; the second explicitly ignores the result. Prefer handling the error or propagating it with `?` when the enclosing return type permits it. A free function or an `impl` block can suppress the lint with `#[allow(unused_result)]`; methods within that impl and nested closures inherit the setting. To explicitly ignore a particular result inside a method, use `let _ = ...`.

This warning has diagnostic code `unused_result`. Its quick fix inserts `let _ = ` and, when needed, a terminating semicolon. The fix requires review because ignoring an error is a deliberate choice; it is not a preferred automatic fix. The compiler does not suggest inserting `?` automatically.

This lint only checks values whose outer type is the built-in `Result`. It does not warn for `Option`, other ordinary return values, user-defined types named `Result`, or a `Result` nested in a tuple or container. It does not track unused bindings or specialize an unconstrained generic parameter to inspect its callers. Returned values and values consumed by another expression are not directly discarded. There is no general `#[must_use]` attribute.

<a id="try-operator"></a>

### `?`

The suffix `?` only supports the built-in semantics `Option[T]` and `Result[T, E]`:

```goml
fn plus_one(value: Option[i32]) -> Option[i32] {
    Some(value? + 1)
}

fn read_number(flag: bool) -> Result[i32, string] {
    let value = parse_number(flag)?;
    Ok(value)
}
```

When using `?` with `Option[T]`, the nearest function or closure must return `Option[_]`. Use of `Result[T, E]` must return `Result[_, E]` with the same error type; there is currently no Rust `From`-style error conversion. `?` will evaluate the operand once.

Diagnostics distinguish an unsupported operand (`try_invalid_operand`), an incompatible enclosing return type (`try_invalid_return`), and different `Result` error types (`try_error_mismatch`). They highlight `?`; propagation failures also identify the enclosing return annotation, or the function name or closure expression when no annotation is available. Error-type mismatches show both types and suggest `map_err` before `?` or changing the enclosing error type. These suggestions do not add implicit conversions.

<a id="go"></a>

### `go`

`go` starts a zero-argument closure that returns `()`:

```goml
go || {
    println("background")
};
```

Don't write `go work();`; that evaluates `work()` first, while `go` requires a value of type `() -> ()`. Write `go || work();` instead.

`go` is a detached, unstructured escape hatch. Its lifetime is not tied to the caller, it does not return a task handle, and its failures are not propagated. Use `std::task` when the caller must wait for child work or coordinate cancellation.

The `unstructured_go` lint warns on each `go` expression. A function or method that intentionally owns detached background work can use `#[allow(unstructured_go)]`.

<a id="closures-and-function-values"></a>

## Closures and function values

<a id="closure-syntax"></a>

### Closure syntax

```goml
let add = |left: i32, right: i32| left + right;
let increment = |value: i32| {
    value + 1
};
let greet = || println("hello");
```

Closure parameter types can be inferred from the expected function type or the calling location. Deferred method calls also resolve inside nested expressions, including tuples, arrays, and control-flow branches. When inference is ambiguous, annotate the closure parameter or expected function type:

```goml
let transform: (i32) -> string = |value| value.to_string();
```

Closures can capture external variables or modify captured `let mut`:

```goml
let mut count = 0;
let next = || {
    count = count + 1;
    count
};
```

`return` in a closure exits that closure and supplies its result.

<a id="no-let-polymorphism"></a>

### No let-polymorphism

Each closure and local function value has only one concrete type:

```goml
let identity = |value| value;
let number = identity(1);
```

The `string` cannot be processed with the same `identity` thereafter. Define top-level generic functions when polymorphism is required:

```goml
fn identity[T](value: T) -> T {
    value
}
```

Top-level functions, closures, and tuple variant constructors can all be passed or returned as function values. Trait method references such as `Read::read` are also function values, with the receiver as their first argument. Calls through these values preserve dynamic dispatch when that receiver is a `dyn` object.
Closures can be nested and return another closure that captures the environment, but each resulting closure expression still has only one concrete function type.
