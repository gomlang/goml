<a id="types-patterns-and-traits"></a>

# Types, patterns, and traits

[English](types.md) | [简体中文](types.zh-CN.md)

[GoML language guide](../goml.md)

Types, structs, enums, generic parameters, pattern matching, traits, and dynamic dispatch. See [language basics](basics.md) for function declarations, expressions, and control flow.

On this page: [Type](#type), [Structs](#structs), [Enums](#enums), [Functions and generics](#functions-and-generics), [Patterns](#patterns), [Trait and impl](#trait-and-impl), [`dyn Trait`](#dyn-trait).

<a id="type"></a>

## Type

<a id="basic-types-and-composite-types"></a>

### Basic types and composite types

| Type | Example | Description |
| --- | --- | --- |
| empty tuple | `()` | The only value is `()` |
| never | `never` | No values; a function annotated `-> never` cannot return normally |
| Boolean | `bool` | `true`, `false` |
| platform-sized integer | `isize` | Corresponds to the `int` of the target Go platform and is also the default type of integers. |
| signed integer | `i8`, `i16`, `i32`, `i64` | fixed width |
| unsigned integer | `usize`, `u8`, `u16`, `u32`, `u64` | `usize` corresponds to the target Go platform's `uint`; the others have fixed widths |
| byte | `byte` | Transparent builtin alias of `u8` |
| floating point | `f32`, `f64` | IEEE floating point |
| SIMD vectors and masks | `simd::u8x16`, `simd::i32x8`, `simd::f32x4`, `simd::f64x4`, matching masks | Fixed 128/256-bit integer and floating-point vectors, imported from `std::simd` |
| string | `string` | Go string backend |
| character | `char` | Compile to Go `rune` |
| tuple | `(i32, string)` | Nonempty tuples in type syntax have at least two elements |
| fixed array | `[i32; 4]` | The length is part of the type |
| function | `(i32, string) -> bool` | parameter type list to return type |
| Generic application | `Option[i32]`, `pkg::Box[string]` | Use square brackets |
| channel | `Channel[isize]`, `Sender[isize]`, `Receiver[isize]` | Bidirectional and directional Go channel backends |
| trait object | `dyn Render`, `dyn Consumer[string] + Close` | One or more dyn-safe trait bounds; associated types must be bound |
| opaque return | `impl Iterator[Item = isize]` | A function's single hidden concrete return type, exposed through its trait bounds |
| Associated type projection | `I::Item`, `Self::Output`, `I::IntoIter::Item` | There must be corresponding trait constraints; projections may be chained |

Example of function type:

```goml
let predicate: (i32) -> bool = |value: i32| value > 0;
let combine: (i32, i32) -> i32 = |a, b| a + b;
let action: () -> () = || println("run");
let accepts_empty: (()) -> () = |value: ()| value;
```

`()` is the empty-tuple type. At the left of `->`, `() -> T` is a zero-parameter function type, while `(()) -> T` is a function taking one `()` parameter.

An expression of type `never` does not produce a value and can appear where another value type is expected. This does not implicitly convert container or function types: `Vec[never]` is not `Vec[isize]`, and `() -> never` is not `() -> isize`. To adapt a non-returning function to a callback with another return type, use a closure with that expected callback type. Ordinary values and normally returning callbacks cannot satisfy `never` annotations.

`(value)` in value and pattern is a group, `(value,)` is a single-element tuple. The current type syntax cannot directly annotate single-element tuples: `(T,)` still normalizes to `T`, so the type of such values must be inferred from the local context.

`A -> B -> C` is parsed by right associative analysis. It is recommended to always write function argument lists in parentheses, especially for higher-order functions: `(A) -> (B) -> C`.

Array lengths are nonnegative integer constant expressions. They accept integer literals in the usual bases, named integer constants, const generic parameters, parentheses, unary `+`, `-`, `~`, and `+`, `-`, `*`, `/`, `%`, `<<`, `>>`, `&`, `|`, `^`. Integer overflow, division by zero, invalid shifts, negative lengths, and cyclic constant references are diagnostics. Lengths fit the target signed indexing range. These expressions do not execute function calls:

```goml
const WIDTH: usize = 2;
let pair: [string; WIDTH] = ["left", "right"];
let bytes: [byte; WIDTH * 2] = [1, 2, 3, 4];
```

The number of array literal elements must match the array type. Empty arrays and empty generic containers usually require type annotations.

`[value; N]` constructs a fixed array by evaluating `value` once and copying that result into every element. Its length follows the same integer constant-expression rules as array types and can be a const generic parameter:

```goml
let zeros: [byte; 16] = [0; 16];
let labels = ["ready"; WIDTH * 2];

fn filled[T, const N: usize](value: T) -> [T; N] {
    [value; N]
}
```

The value is evaluated even when `N` is zero. Repetition uses ordinary value copying, so repeating a `Ref`, `Vec`, or another shared handle preserves shared storage; it does not clone that storage or call the initializer once per element. Array repetition is also available in `comptime` and constant initializers when its value is supported there.

<a id="type-aliases"></a>

### Type aliases

Top-level aliases are transparent and may have type parameters:

```goml
type UserId = u64;

type Pair[T] = (T, T);

pub type Names = Vec[string];
```

Aliases do not create nominally distinct types and recursive alias cycles are rejected. Public aliases can be referenced across packages. Associated method calls and method values through an alias use the underlying type's methods; generic alias arguments may be explicit or inferred from the call and expected type.

Use a single-field tuple struct when a value needs a distinct nominal type instead of an alias:

```goml
struct UserId(u64);

struct Box[T](T);
```

This form is the newtype pattern. Construct it with `UserId(value)`, access its field with `value.0`, and destructure it with `let UserId(inner) = value;`. The field is private by default. A public newtype that can be constructed and unwrapped across package boundaries marks both the type and field public: `pub struct UserId(pub u64);`.

<a id="type-syntax-not-currently-available"></a>

### Type syntax not currently available

GoML has no Rust reference or lifetime syntax, pointer arithmetic, slice literals, or union types. Use `Ref[T]` for shared mutable storage, `Option[T]` for optional values, and `Slice[T]` or `MutSlice[T]` for read-only or mutable contiguous views. `std::ffi::Ptr[T]` carries nullable Go pointers and remains distinct from `Ref[T]` and numeric types. `ffi::null()` and `ffi::is_nil()` operate on this explicit pointer type. A type declared with `extern type` remains opaque even when its Go declaration aliases a pointer; it does not gain these pointer operations automatically.

On Linux amd64, `std::os::linux::syscall` accepts numeric machine words, scoped byte-buffer arguments, and explicit pointer fields between byte buffers for low-level kernel calls. It does not add pointer casts, pointer arithmetic, or native struct layout to the language.

`A + B` combines trait bounds in generic constraints, supertrait lists, `dyn A + B` objects, and `impl A + B` returns. It is not a general intersection type for concrete values.

`Self` is only used in the trait signature and the type position of impl; ordinary top-level functions cannot use `Self` as an implicit type parameter.

<a id="structs"></a>

## Structs

```goml
struct Point {
    x: i32,
    y: i32,
}

struct Box[T] {
    value: T,
}
```

A newtype is a nominal structure with exactly one unnamed field:

```goml
struct UserId(u64);

fn increment(value: UserId) -> UserId {
    UserId(value.0 + 1)
}
```

Unlike `type UserId = u64;`, the newtype is not interchangeable with `u64`. It must be explicitly constructed or destructured. Newtypes may be generic, their field is private by default, and a trailing comma inside the parentheses is accepted. This syntax requires exactly one field; tuple structs with zero or multiple fields are not supported.

Struct and enum type parameter lists do not accept bounds. Put constraints on the functions, traits, or impls that use the type.

Write out the fields during construction, allowing field abbreviations:

```goml
fn make_point(x: i32, y: i32) -> Point {
    Point { x, y }
}
```

Named-field literals do not accept turbofish arguments. Supply a type annotation when field values do not provide the intended generic type, for example `let value: Box[bool] = Box { value: true };`. This also applies to structure updates.

Field access uses dot notation:

```goml
let x = point.x;
```

Structure update copies omitted fields from a base value of the same structure type. The `..base` item must be last. Explicit field expressions are evaluated from left to right, followed by the base expression, and every expression is evaluated once:

```goml
let moved = Point { x: point.x + 1, ..point };
```

Structure update is rejected when inaccessible fields prevent construction across a package boundary.

Directly recursive structures will have infinite size and must be recursed through indirect layers such as `Ref` and `Vec`:

```goml
struct Node {
    value: i32,
    next: Option[Ref[Node]],
}
```

<a id="enums"></a>

## Enums

```goml
enum Message[T] {
    Quit,
    Value(T),
    Pair(T, T),
    Named { value: T },
}
```

Unit variants are values, tuple variants use function-call syntax, and struct-like variants use named fields:

```goml
let quit: Message[i32] = Message::Quit;
let value: Message[i32] = Message::Value(42);
let pair = Message::Pair("left", "right");
let named = Message::Named { value: 42 };
```

A unit variant provides no payload from which to infer generic arguments, so it usually needs an expected type:

```goml
let none: Option[i32] = Option::None;
```

Tuple variant constructors are also available as first-class function values:

```goml
let some: (i32) -> Option[i32] = Option::Some;
```

Explicit generic arguments belong to the enum owner and appear before the variant name. They constrain the payload and result type for both direct construction and constructor function values:

```goml
let value = Message::[i32]::Value(42);
let quit = Message::[string]::Quit;
let make = Message::[bool]::Value;
let ready = make(true);
```

For an enum with const parameters, the same list includes its const arguments, such as `Storage::[i32, 2]::Values([1, 2])`. Supplied arguments must match the enum's parameter count and kinds. A variant has no independent generic parameter list, and explicit arguments must agree with any expected type and payload types.

Struct-like variants use an expected type instead of turbofish arguments: write `let value: Message[bool] = Message::Named { value: true };`. `Message::[bool]::Named { value: true }` is rejected.

An expected enum type also permits an unqualified constructor in an expression:

```goml
let none: Option[i32] = None;
let nested: Option[Result[i32, string]] = Some(Ok(42));
let message: Message[string] = Named { value: "ready" };
let some: (i32) -> Option[i32] = Some;

fn success(value: i32) -> Result[i32, string] {
    Ok(value)
}
```

Annotations, function parameters, return types, and enclosing constructor payloads can provide this context. It works for unit, tuple-like, and struct-like variants, including public variants from imported enums and enums reached through a type alias. An expected constructor function type uses its return type to select the enum. Ordinary name resolution takes precedence; an existing local value or struct is not replaced with an enum constructor. A unique variant declared in the current package remains usable without expected type information. When neither ordinary name resolution nor the expected type selects a variant, qualify it with `Enum::Variant`.

Structural variants use field constructs and field patterns:

```goml
match named {
    Message::Named { value } => value,
    _ => 0,
}
```

Patterns may omit the enum qualifier. When the expected pattern type is `Option[T]`, `Some(value)` and `None` resolve to `Option::Some(value)` and `Option::None`. The same rule applies to unit, tuple-like, and struct-like variants of user enums, including enums imported from another package and enums reached through a type alias. Duplicate variant names in unrelated enums are not ambiguous because the expected enum type selects the owner. Nested patterns are resolved recursively, so `Some(Ok(value))` uses the payload type of `Some` to resolve `Ok`.

<a id="functions-and-generics"></a>

## Functions and generics

<a id="opaque-return-types"></a>

### Opaque return types

A function can expose trait bounds while hiding the concrete return type with `-> impl Trait`. Bounds may contain generic arguments, associated type bindings, and `+`:

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

Every return path of one function must produce the same concrete type for a given generic instantiation, and that type must satisfy all declared bounds. Different generic instantiations may use different concrete types. A public function may hide a private type, including across package boundaries. Callers can use the declared trait bounds and associated type bindings, but cannot access the hidden type's fields or additional inherent methods. Each function defines its own opaque type even if two functions return the same concrete type.

Dispatch remains static and participates in monomorphization; `impl Trait` does not create a trait object or require dyn-safe bounds. Use an explicit `as dyn Trait` conversion when a trait object is needed. `impl Trait` is supported only as the complete return type of a top-level function or inherent method. It cannot appear in parameters, type aliases, local type annotations, nested return types such as `Option[impl Trait]`, or trait method declarations.

Opaque bounds currently accept positional type arguments and associated type bindings, but not integer const arguments such as `impl Fixed[4]`.

<a id="generic-functions"></a>

### Generic functions

```goml
fn identity[T](value: T) -> T {
    value
}

fn render[T: ToString + Eq](value: T) -> string {
    value.to_string()
}
```

Constraints can also be placed in a `where` clause:

```goml
use std::iter;

fn collect_items[T, I: Iterator](iterator: I) -> Vec[T] where I::Item = T {
    iter::collect(iterator)
}
```

There are currently two types of `where` predicates:

```text
Type: TraitA + TraitB
TypeA = TypeB
```

Ordinary generic function calls usually infer the type arguments from the argument and expected result types:

```goml
let number: i32 = identity(1);
let text: string = identity("text");
```

Use GoML turbofish when explicit specification is required:

```goml
let number = identity::[i32](1);
let text = identity::[string]("text");
```

An explicit generic argument list must contain at least one argument and match the declared parameter count, including const parameters. Omit the whole list to use inference; `identity::[](1)` is rejected.

Don't write `identity[i32](1)` or Rust's `identity::<i32>(1)`. Generic arguments for an owner type or trait appear before the member name, while arguments owned by a method appear after it:

```goml
let inferred = value.convert(fallback);
let explicit = value.convert::[string](fallback);
let inherent = Box::[i32]::convert::[string](value, fallback);
let trait_call = Convert::[i32]::convert::[string](value, fallback);
```

Top-level functions and methods may introduce their own type parameters. A method's parameters are distinct from the parameters of its enclosing trait or impl and may have their own bounds and `where` predicates. Local named functions do not exist; use closures. Closures do not have generics. Structures, enumerations, traits, and impl blocks can also have type parameters.

An in-scope generic type parameter takes precedence over an unqualified package type or imported name. For example, in a package that also defines `struct T`, the parameter and return types of `fn identity[T](value: T) -> T { value }` refer to the generic parameter. A qualified path continues to refer to the package item.

A free function may use a type parameter only in its body, with no occurrence in its parameter or result types. Supply such arguments explicitly, for example `schema::[User]()`, including when taking a function value. Those arguments are retained across package interfaces and specialization; an unused generic definition does not generate an unspecialized executable body.

GoML monomorphizes generic calls. Recursive generic code must produce a limited number of concrete instances and cannot continually change to a new nested type with each recursive call.

<a id="const-generic-parameters"></a>

### Const generic parameters

Functions, structs, enums, aliases, traits and implementations can declare `const N: usize` alongside type parameters. A const parameter is a value in an expression and an integer argument in a generic application; it cannot be used as an element type. Each concrete argument participates in monomorphization and nominal type identity.

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

A direct length parameter can be inferred from an array argument or expected type. For a length expression such as `N + 1`, provide `N` explicitly; inference does not solve arithmetic equations. Generic arguments use literals or names, and compound integer arguments use braces, such as `Buffer[i32, {2 + 1}]`. Const parameters currently have type `usize`, have no defaults, and use values in `0..=9223372036854775807`. Array repetition accepts these parameters, for example `[value; N]`. Type-dependent associated constants in array lengths are not supported.

<a id="patterns"></a>

## Patterns

Patterns can be used with `let`, `for`, `match`, `if let` and `while let`.

| Pattern | Example |
| --- | --- |
| variable binding | `value` |
| Wildcard | `_` |
| empty tuple | `()` |
| bool/number/string/character | `true`, `42`, `"ok"`, `'x'` |
| top-level constant | `ANSWER` or `config::ANSWER` |
| Group | `(pattern)` |
| tuple | `(left, right)` |
| Single element tuple | `(value,)` |
| exact structure | `Point { x, y: other }` |
| partial structure | `Point { x, .. }` |
| newtype | `UserId(inner)` |
| Enum unit variant | `Color::Red` or contextual `Red` |
| Enum tuple variant | `Some(value)` or qualified `Option::Some(value)` |
| Enum struct-like variant | `Message::Named { value }` or contextual `Named { value }` |
| Array, Vec, Slice or MutSlice | `[first, second]`, `[first, .., last]` |
| rest binding | `[first, middle @ .., last]` |
| Alias | `whole @ Some(value)` |
| or-pattern | `Color::Red \| Color::Blue` |
| Range | `0..10`, `'a'..='z'` |
| Nested pattern | `Ok((key, value))` |

<a id="refutability-and-binding"></a>

### Refutability and binding

`let` and `for` require that the pattern be irrefutable, that is, the type at that position must succeed. `match`, `if let` and `while let` can use refutable patterns:

```goml
let (left, right) = pair;
if let Some(value) = candidate {
    println(value);
}
```

Irrefutability is judged jointly by type and sub-pattern rather than by surface syntax alone. For example, the destructuring of a single-variant enum is irrefutable, as is the complete destructuring of a fixed-length array of known length; [first, ..] of a Vec[T] may fail with an empty Vec.

Variables with the same name cannot be bound repeatedly in the same pattern branch. Each alternative of the or-pattern must be bound to exactly the same set of variables, and the corresponding variables must be of the same type:

```goml
match value {
    Left(shared) | Right(shared) => shared,
}
```

Pattern binding is only visible in the corresponding `let` subsequent scope, `for` / `while let` loop body, `if let` then branch, or `match` branch guard and branch body. In an `if` or `while` let chain, it is also visible to subsequent terms in that chain.

<a id="groups-tuples-structures-and-enumerations"></a>

### Groups, tuples, structures and enumerations

`(pattern)` groups a pattern; `(pattern,)` matches a single-element tuple:

```goml
let (only,) = one_tuple;
```

Structure field patterns support abbreviations with the same name:

```goml
let Point { x, y } = point;
```

Struct patterns are exact by default and must list all fields. Unlisted fields can be ignored using the trailing `..`:

```goml
let Point { x, .. } = point;
```

`..` appears at most once in the structure pattern and must be the last item. Duplicate fields, unknown fields, and missing fields without writing `..` will all generate diagnoses.

The number of constructor parameters of the enumeration pattern must be consistent with the definition. An unqualified variant is resolved from the expected enum type rather than from global uniqueness. For example, `Shared` in a pattern for `First` means `First::Shared` even when `Second::Shared` also exists. If `First` has no `Shared` variant, the compiler reports that error instead of selecting `Second::Shared`. Type aliases and imported enum types participate after normalization. Nested variant payloads supply the expected type for nested patterns.

<a id="fixed-arrays-vec-slice-and-mutslice"></a>

### Fixed arrays, Vec, Slice and MutSlice

Fixed arrays `[T; N]`, `Vec[T]`, `Slice[T]` and `MutSlice[T]` share square-bracket pattern syntax. Without `..`, the pattern requires the exact length; with `..`, the prefix and suffix only specify the minimum length:

```goml
match values {
    [] => "empty",
    [only] => "one",
    [first, .., last] => "many",
}
```

A sequence pattern allows at most one rest. It can appear at any position; `name @ ..` binds the elements between the explicit prefix and suffix:

```goml
let values: [i32; 4] = [1, 2, 3, 4];
let [first, middle @ .., last] = values;
```

For fixed arrays, the number of elements must be exactly equal to `N` when rest is omitted; when rest is included, the total number of elements of explicit prefixes and suffixes cannot exceed `N`. In the above example, the type of `middle` is `[i32; 2]`.

For `Vec[T]`, `Slice[T]` and `MutSlice[T]`, the binding type of `name @ ..` is read-only `Slice[T]`. For example `[head, tail @ ..]` requires at least one element, `tail` will not be copied into a new Vec. Since dynamic sequences may not be long enough, such patterns are usually placed inside a `match`, `if let` or `while let`.

<a id="alias-and-or-pattern"></a>

### Alias and or-pattern

`name @ pattern` binds the entire matched value to `name` while continuing to match the inner pattern:

```goml
match value {
    whole @ Some(inner) => use_both(whole, inner),
    None => fallback(),
}
```

`@` is more tightly bound than `|`. Therefore `whole @ A | B` resolves to `(whole @ A) | B`, and usually an error will be reported because the variables bound to the two alternatives are different. To make the alias cover the entire or-pattern, you must write parentheses:

```goml
whole @ (Left(value) | Right(value))
```

Or-pattern alternatives are tried from left to right and can be nested in tuple, struct, enum, and sequence patterns.

<a id="range-patterns"></a>

### Range patterns

`start..end` does not contain an upper bound, `start..=end` contains an upper bound:

```goml
match character {
    'a'..='z' => "lowercase",
    'A'..='Z' => "uppercase",
    _ => "other",
}
```

Both endpoints must be integer or character literals of the same concrete type. Signed integer endpoints can be negative. The exclusive range requires the lower bound to be strictly less than the upper bound, and the inclusive range requires the lower bound to be less than or equal to the upper bound. Open ranges, floating point ranges, or string ranges are not currently supported; separate `..` in sequences and structures is a rest, not a range.

<a id="exhaustiveness-and-unreachable-branches"></a>

### Exhaustiveness and unreachable branches

The compiler checks whether `match` is exhaustive for the matched type and gives examples of missing patterns when it is not. The analysis covers bool, enumerations, tuples, structures, fixed arrays, Vec/Slice/MutSlice, integer and character ranges, aliases and or-patterns, and will also warn about branches that are never matched.

An empty `match` is only valid for types that have no constructible value:

```goml
enum Never {}

fn absurd(value: Never) -> i32 {
    match value {}
}
```

Exhaustive analysis ignores enumeration variants with no value for the payload type and also identifies purely recursive enumerations with no base variant. For example, `enum MaybeNever { Empty, Filled(Never) }` requires only an `Empty` arm; `enum Loop { Next(Loop) }` has no constructible value.

Branches with normal guards do not provide exhaustive coverage because guard may be false; subsequent unguarded branches are usually required. String and floating-point literals cannot enumerate the entire type, and `_` is also usually required when matching these types. In floating-point patterns, `-0.0` and `0.0` are regarded as the same value.

There are no `ref` or `ref mut` patterns. `mut name` is supported for individual bindings, while pattern matching on `dyn Trait` is not supported.

<a id="trait-and-impl"></a>

## Trait and impl

<a id="define-traits"></a>

### Define traits

Trait method signatures must include parameter names. The first receiver can be written as `self`, which is short for `self: Self`:

```goml
trait Render {
    fn render(self) -> string;
}

trait Convert[T] {
    fn convert(self, fallback: T) -> T;
}
```

A trait may also declare an associated function without a receiver. `Self` is inferred from its arguments or the expected return type:

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

An impl method must preserve whether the trait declaration has a `self: Self` receiver. Associated functions use static dispatch, cannot be called with method syntax, and do not participate in `dyn` dispatch. If neither arguments nor the expected result determine `Self`, the call is rejected as ambiguous.

Traits can have supertraits, generic constraints, and associated types:

```goml
trait Provider: Render {
    type Item: ToString;
    fn get(self) -> Self::Item;
}
```

Associated type projections may be chained when every step is constrained by its declaring trait. For example, `I::IntoIter::Item` projects `IntoIter` from `I`, then projects `Item` from that iterator type. Chained projections can appear in signatures and `where` equalities:

```goml
trait IntoIterator where Self::Item = Self::IntoIter::Item {
    type Item;
    type IntoIter: Iterator;
}
```

Generic trait example:

```goml
trait Child[T: ToString]: Parent[T] + Render where T: Eq {
    fn child(self) -> string;
}
```

The standard comparison hierarchy follows the distinction between partial comparison and total equivalence:

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

`Eq` is a marker trait and has no methods. Equality behavior belongs in `PartialEq`; a total-equivalence type implements `PartialEq` with `fn eq` and adds an empty `impl Eq`. `Eq::eq` and an `eq` method inside `impl Eq` are invalid.

Trait methods may provide a default body. An impl may omit such a method and will inherit the default; an explicitly declared impl method overrides it. Default bodies are checked using the trait's generic parameters, predicates, associated types, and `Self`, and remain available across package boundaries:

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

Trait methods may also declare their own type parameters and constraints:

```goml
trait Convert {
    fn convert[U: ToString](self, fallback: U) -> string;
}
```

The corresponding impl method must have the same method-generic arity, signature, and constraints. Generic trait methods use static dispatch and make the trait unavailable as `dyn`.

<a id="trait-impl"></a>

### Trait impl

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

The type parameters of generic impl are written after `impl`:

```goml
impl[T: ToString] Render for Box[T] {
    fn render(self: Box[T]) -> string {
        self.value.to_string()
    }
}
```

Different applications of generic traits are different impl:

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

Associated types are bound in impl:

```goml
impl Iterator for Counter {
    type Item = i32;

    fn next(self: Counter) -> Option[i32] {
        next_value(self)
    }
}
```

When implementing a trait with supertraits, you also need to provide the impl of the target type for each supertrait. The compiler rejects overlapping impls and enforces the orphan rule: at least one of the trait or the nominal type being implemented must belong to the current package. The target nominal type of inherent impl must belong to the current package.

<a id="inherent-impl"></a>

### Inherent impl

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

The associated function without a receiver is called with `Point::new(1, 2)`; the method whose first parameter is the receiver can be called with `point.sum()` or `Point::sum(point)`.

Inherent impl of generic types:

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

Method type arguments are normally inferred. Write `box_value.map::[string](../convert)` when explicit arguments are needed, or `Box::[i32]::map::[string](box_value, convert)` in associated form.

Concrete and nested specializations also support associated constructors, for example `impl Box[f64]` and `impl[T] Box[Vec[T]]`. In `Box::[Vec[string]]::repeated("x")`, the owner argument denotes `Vec[string]`, while the implementation parameter `T` is `string`. Omitting owner arguments creates inference variables; an expected result such as `let value: Box[f64] = Box::empty();` can select a specialization. If several candidates remain, use explicit owner arguments; ambiguous lookup reports a diagnostic.

<a id="method-parsing-and-ufcs"></a>

### Method parsing and UFCS

For specific values, use:

```goml
value.render();
Render::render(value)
```

Cross-package method syntax requires that the trait be in the method scope of the current file:

```goml
use acme::render;
use render::Render;
```

If multiple visible traits define methods with the same name, `value.render()` will be ambiguous. Use UFCS to explicitly select:

```goml
let a = A::render(value);
let b = B::render(value);
```

Generic trait UFCS uses `::[...]` after the trait name to give the type parameters:

```goml
let number = Convert::[i32]::convert(token, 0);
let text = Convert::[string]::convert(token, "");
```

Trait implementation and method candidate searches do not retain type bindings from rejected candidates. Reordering otherwise equivalent implementation declarations does not select a different inferred type.

Trait bounds make the corresponding methods available on generic values. Imported short trait names, renamed imports, and public re-exports retain the defining trait identity in generic bounds and `where` predicates through specialization. Supertraits and associated type bounds also participate in method resolution as implied constraints.

<a id="associated-constants"></a>

### Associated constants

Traits can require constants or provide defaults. Inherent constants follow method visibility rules; trait implementation constants inherit trait visibility. Read a constant through `Type::NAME`, `Type::[Arguments]::NAME`, or a bound parameter such as `T::NAME`.

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

Initializers support literals, constant paths, unary and binary operations, casts, and composite constant construction. They cannot perform ordinary function or method calls or access runtime locals. A generic implementation can use its const parameters in the initializer. Non-generic inherent integer constants can also supply array lengths. Associated constant patterns and trait-associated constants in length expressions are not supported. Traits with associated constants cannot be used as `dyn` traits.

<a id="dyn-trait"></a>

## `dyn Trait`

Use `as dyn Trait` to convert a concrete value that satisfies a dyn-safe trait into a trait object:

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

Concrete values are never boxed implicitly when a `dyn Trait` value is expected. Assignments, arguments, returns, branches, constructors, and collection operations must place `as dyn Trait` at the conversion site. Existing `dyn Trait` values can still flow through compatible typed contexts without another conversion.

The `dyn` value supports method syntax, and UFCS can also be used; it can be seen that supertrait methods are also available:

```goml
value.show();
Display::show(value)
```

Associated types are fixed in square brackets and are available to dyn-dispatched method signatures:

```goml
trait Source {
    type Item;
    fn get(self: Self) -> Self::Item;
}

fn read(source: dyn Source[Item = isize]) -> isize {
    source.get()
}
```

Type inference rejects recursive types, including cycles through a trait object’s associated type bindings. A rejected cycle produces a type diagnostic instead of creating a self-referential inference variable. Generic substitution and inference resolution traverse nested type constructors, including trait-object arguments, associated type bindings, and additional bounds.

Generic traits accept positional type arguments before associated type bindings, such as `dyn Consumer[string, Error = IoError]`. Every required generic argument must be supplied, and every associated type must be bound exactly once. Positional arguments after the first `Name = Type` binding are rejected:

```goml
trait Consumer[T] {
    fn consume(self, value: T) -> string;
}

fn send(value: dyn Consumer[string]) -> string {
    value.consume("hello")
}
```

Combine dyn-safe traits with `+` to expose their methods through one object. The concrete value must implement every bound with the specified arguments and associated types:

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

Supertrait methods remain available. If different component traits provide the same method name, use UFCS to choose the trait, for example `A::render(value)` or `B::render(value)`. For a generic trait, use `Consumer::[string]::consume(value, "hello")`.

The order of bounds does not change the object type: `dyn A + B` and `dyn B + A` are compatible. Repeating a trait name is rejected, even with different generic arguments. Supertrait paths must agree on a trait's generic arguments. A child trait's method takes precedence over a method inherited from its parent; unrelated methods with the same name require UFCS.

An inherited associated type with a unique name may be bound on the child, for example `dyn Child[Item = isize]`, or on its declaring trait with `dyn Child + Parent[Item = isize]`. When different parent traits declare the same associated type name, bind each on its declaring trait, such as `dyn Child + A[Item = isize] + B[Item = string]`.

Current dyn-safe conditions:

- Each method must have a first receiver parameter of exactly type `Self`;
- Direct `Self` cannot appear in other parameters or return types, while a bound projection such as `Self::Item` is allowed;
- Methods cannot declare type parameters;
- Traits cannot declare associated constants;
- Every component trait and supertrait must satisfy these rules.

Current limitations:

- Integer const arguments such as `dyn Fixed[4]` are not supported in dyn argument lists;
- Conversion from one dyn object type to another, including dropping a bound, is not supported;
- Pattern matching on `dyn Trait` is not supported.
