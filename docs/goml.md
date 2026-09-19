# GoML language guide

This guide describes the GoML syntax, type rules, tools, and public library APIs implemented in this repository. Use it when writing, modifying, and reviewing `.gom` source. Similarities to Rust, Go, or OCaml do not imply that their syntax or APIs are supported.

GoML is a statically typed language with garbage collection. Its syntax is close to Rust, while its semantics are closer to ML. The compiler monomorphizes generics and lambda-lifts GoML closures before emitting Go. Explicit Go FFI can preserve native generic types and function values. GoML has no ownership, borrowing, lifetimes, or manual memory management.

Quick navigation: [packages](#modules-packages-and-imports), [types](#type), [control flow](#control-flow), [patterns](#patterns), [traits](#trait-and-impl), [compile-time evaluation](#compile-time-evaluation), [Go FFI](#go-ffi), [tests](#tests), [prelude](#built-in-prelude), [standard library](#standard-library-package), and [verification](#verify-generated-code).

Build and installation instructions are in the [repository README](../README.md) and [release guide](releasing.md). See [formatting](formatting.md) for canonical source layout and [compile-time evaluation architecture](comptime.md) for CTIR internals.

## Essential rules

1. The project source code uses `.gom`; the project root directory uses `goml.toml` to declare the canonical module path.
2. Each source file in the project first writes `package name;`, then writes the file's own `use`, and finally writes the top-level definition.
3. Parameters of top-level functions must have types; the return type is fixed to `()` when omitted and is not inferred from the function body.
4. Generics use square brackets: `Vec[i32]`, `fn id[T](x: T) -> T`, not `<...>`.
5. Generic calls usually rely on type inference; when explicit type arguments are required, write `id::[i32](1)`, not `id[i32](1)` or Rust's `id::<i32>(1)`.
6. `if`, `match`, and `select` are expressions. `if` without `else` must return `()`; `match` must be exhaustive.
7. The last semicolon-free expression of a block is the block value; adding a semicolon discards the value.
8. `let` and assignment statements must end with a semicolon. Mutable bindings may be introduced with `let mut pattern` or precisely inside a pattern with `mut name`; the semicolon can be omitted for `if`, `match`, `select`, and unlabeled `while`, `loop` and `for` statements. Labeled loop statements require a semicolon before a following statement.
9. Enumeration construction uses full names such as `Option::Some(value)`. In patterns, the enum qualifier may be omitted when the matched type determines it, such as `Some(value)` and `None`.
10. For cross-package calls, write `alias::item`. Top-level items, struct fields, and inherent methods must all be marked with `pub` as required.
11. Before using trait method syntax across packages, import the package and trait with `use alias::Trait;` or a braced import; when in doubt, use UFCS: `Trait::method(value)`.
12. Use `module::path` for a package below the current module root. Do not generate `mod`, `crate::`, `self::`, `super::`, root paths `::x`, Rust references, or Go `var` / `:=`. A user `extern fn` is valid only with the typed Go FFI attribute described below.
13. The test function uses `#[test]`, which must have no parameters, no type parameters and return `()`; the white-box test is placed in `*_test.gom` of the same package, and the black-box test is placed in the `tests/` directory of the package under test.

## Minimal program

Single-file compilation allows omission of `package main;`:

```goml
fn main() -> () {
    println("hello, goml")
}
```

The entry file in the project should explicitly declare the package:

```goml
package main;

fn main() -> () {
    println("hello, goml")
}
```

`main` cannot have parameters or type parameters. When generating an executable file, the package selected as the entry must be declared as `package main;` and define `fn main()`. It is recommended to let `main` return `()`.

## Lexical rules

### Identifiers and keywords

Ordinary identifiers only use ASCII and are of the form:

```text
[A-Za-z][A-Za-z_0-9]* | _[A-Za-z_0-9]+
```

The single `_` is a wildcard character, not an ordinary variable name. The current syntax enforces the first letter case of name categories:

- Package names, package aliases in `use package as alias`, functions, methods, parameters, local bindings and fields must start with a lowercase letter or `_`; imported item aliases may follow the naming convention of the imported item;
- Structures, enumerations, traits, enumeration variants, generic parameters, and associated types must start with a capital letter;
- Paths retain the appropriate case for the referenced name.

Enumeration construction should use `Enum::Variant`. Patterns may omit `Enum::` because their expected type determines the variant owner.

Common keywords include:

```text
package use as pub fn struct enum trait impl for type const static where defer
let mut if else match while loop for in break continue return go dyn
true false bool isize i8 i16 i32 i64 usize u8 u16 u32 u64
f32 f64 string char extern
```

`()` is the only spelling of the empty-tuple type.

`self` is a special abbreviation for the receiver parameter, and can also be used as a common receiver variable name; `Self` has special meaning in the type position of trait and impl.

Whitespace separates tokens but does not determine block structure. Only line comments from `//` to the end of the line are supported; block comments are not supported.

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

```goml
let small: u8 = 42;
let ratio: f32 = 0.5;
let values: [i16; 3] = [1, 2, 3];
```

Strings support `\"`, `\\`, `\n`, `\r`, `\t`, `\b`, `\f`, `\/` and four-digit `\uXXXX` escaping; characters are escaped using the same set of control characters, and `\'` is used to represent single quotes. Ordinary strings cannot span lines.

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

## Modules, packages and imports

### `goml.toml`

A project module declares the canonical path in the root directory:

```toml
[module]
path = "alice::myapp"

[build]
target-dir = "_artifact"

[dependencies]
"alice::http" = "1.2.0"
```

Dependency versions must use the strict `X.Y.Z` form. A dependency version is a minimum version requirement resolved using MVS; there is currently no `goml.lock`.

`[build]` can be omitted; `build.target-dir` defaults to `_artifact` under the module root. The manifest value must be a non-empty relative path and cannot contain a `..` segment. `goml check`, `goml build`, `goml run`, and `goml test` can temporarily override it with `--target-dir <path>`; command-line overrides may be relative or absolute. `goml clean` removes the configured target directory. Its optional `--target-dir <path>` override must stay inside the module.

Each module path segment must be non-empty and may contain ASCII letters, digits, `_`, and `-`. Paths rooted at `builtin` or `prelude` are reserved for the toolchain and cannot be used as a module path or dependency. The `[module]` section currently has no `name`, `kind`, `root`, or similar fields.

### Build output layout

The default product layout is as follows; the standard package paths in the path are expanded by directory level, for example, `alice::myapp::utils` corresponds to `alice/myapp/utils`:

```text
_artifact/
├── check/
│   ├── pkg/<canonical-package-path>/<last-path-segment>.interface
│   └── deps/<owner>/<module>/<version>/pkg/<canonical-package-path>/<last-path-segment>.interface
├── build/
│   ├── pkg/<canonical-package-path>/<last-path-segment>.interface
│   ├── pkg/<canonical-package-path>/<last-path-segment>.core
│   ├── pkg/<entry-package-path>/goml_generated.go
│   └── deps/<owner>/<module>/<version>/pkg/<canonical-package-path>/<last-path-segment>.*
├── test/
│   ├── base/
│   │   ├── pkg/<canonical-package-path>/<last-path-segment>.*
│   │   └── deps/<owner>/<module>/<version>/pkg/<canonical-package-path>/<last-path-segment>.*
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

`<last-path-segment>` is the last segment of the canonical package path, not the package name declared in the source code; for example, `alice::app::cmd::server` declared as `package main` still uses `server.core` and `server.interface`.

`check` only requires interface artifacts; `build` and `test` also generate Core. Each entry package declared as `package main` generates a Go link file named `goml_generated.go` in its own artifact directory; ordinary library packages do not generate Go link files. Fixed filenames do not trigger Go's special file rules when a package path ends with `_test`, an operating system name, or an architecture name.

Interface and Core artifacts use the deterministic GAF binary container. GAF
payloads are encoded directly from compiler data without an intermediate JSON
value tree.

`for` loops over builtin `range`, `Vec`, `Slice`, and `MutSlice` values lower directly to indexed loops without allocating iterator closures.

The executable file of the module root entry package is `bin/<module name>`; the nested entry package retains the directory within the module and appends the entry name. For example, the output of `alice::app::cmd::server` is `bin/cmd/server/server`. Production test dependencies are built once under `test/base`; internal and external tests each share one Go entry, manifest, and runner at the corresponding test-kind root. The runner file has the `.exe` suffix on Windows. External dependencies use the same `deps/<owner>/<module>/<version>/pkg/...` structure in the root directory of each stage.

The configured production directory will not participate in package discovery and cannot be a target to be inspected, built, or tested. The `.gitignore` generated by `goml new` contains `/_artifact/` by default; after modifying `build.target-dir`, the project's own `.gitignore` should be modified simultaneously.

### Directory packages

Each directory containing a `.gom` file is a package. All source files in the same directory must declare the same package name:

```goml
package utils;
```

The canonical identity of a package is the module path plus the relative directory path. The package declaration name determines the local name when imported without aliases, but not the global identity of the package. The canonical identity of the root package is the module path.

A typical project could be:

```text
goml.toml
main.gom
utils/utils.gom
```

`utils/utils.gom`:

```goml
package utils;

pub fn message() -> string {
    "hello"
}
```

`main.gom`:

```goml
package main;

use alice::myapp::utils;

fn main() -> () {
    println(utils::message())
}
```

### `use`

`use` is a file-level declaration and must be placed after `package` and before any top-level items. Imports are not shared between different files.

Imports are not transitive: package A imports package B, which does not allow files using A to automatically see B. Each file must directly import each package it references.

```goml
use alice::http::client;
use alice::http::client as http_client;
```

The contextual `module` path head denotes the canonical path declared by the current project's `[module].path`. It may be used from any package in that module, including modules whose canonical path contains multiple segments:

```goml
use module::http::client;
use module::rendering::api::{Canvas, Render};
pub use module::model::Request;
```

For a module declared as `alice::myapp`, these paths resolve to `alice::myapp::http::client`, `alice::myapp::rendering::api`, and `alice::myapp::model::Request`. The marker applies only at the start of a `use` path followed by `::`; dependency paths and aliases named `module` elsewhere retain their ordinary meaning. External dependencies continue to use canonical paths such as `alice::http::client`.

After importing the package, access its public items through local aliases:

```goml
let request = http_client::new_request();
```

Braced imports bring selected public top-level items directly into the current file. Functions, constants, structs, enums, type aliases, and traits can be imported, and each item may be renamed independently:

```goml
use alice::rendering::api::{Canvas, Color as Paint, Render, render_to_string};

fn describe(canvas: Canvas) -> string {
    render_to_string(canvas)
}
```

The path before the braces is always a package path. Nested braced imports, `self`, glob imports with `*`, enum-variant imports, and subpackage aggregation are not supported. Imported names remain file-local, and normal top-level visibility rules apply.

The following special form also adds a trait from an already loaded package to the method calling scope:

```goml
use alice::rendering::api;
use api::Render;
```

Importing `Render` in a braced list also adds it to the method scope. Afterwards, the specific value can be written as `value.render()`. Even if the trait is not added to the method scope, you can still write a qualified call to `api::Render::render(value)`.

`pub use` re-exports a named package or public item from the current package without changing its identity:

```goml
pub use alice::http::client;
pub use client::{Client, Request, Response};
pub use client::Request as HttpRequest;
```

Downstream code can use `facade::HttpRequest`, import it directly, or bring a re-exported trait into method scope. The compiler resolves every such path back to the original declaration, so constructors, trait implementations, methods, constants, and functions behave exactly as they do at the original path. Re-exported package namespaces also remain usable, for example `facade::client::Request`. Re-export names in one package must be unique. Glob re-exports remain unsupported.

The following path models are not supported:

```text
mod child;
crate::x
self::x
super::x
::rooted::x
```

### Visibility

Top-level items are only visible within the package by default. Top-level functions, structures, enumerations and traits are exported using `pub`; structure fields and inherent impl methods are also private by default and need to be marked `pub` separately for cross-package access:

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

All files in the same package can use private top-level items. The trait impl method inherits the visibility of the trait method and cannot write `pub`.

## Type

### Basic types and composite types

| Type | Example | Description |
| --- | --- | --- |
| empty tuple | `()` | The only value is `()` |
| Boolean | `bool` | `true`, `false` |
| platform-sized integer | `isize` | Corresponds to the `int` of the target Go platform and is also the default type of integers. |
| signed integer | `i8`, `i16`, `i32`, `i64` | fixed width |
| unsigned integer | `usize`, `u8`, `u16`, `u32`, `u64` | `usize` corresponds to the target Go platform's `uint`; the others have fixed widths |
| byte | `byte` | Transparent builtin alias of `u8` |
| floating point | `f32`, `f64` | IEEE floating point |
| string | `string` | Go string backend |
| character | `char` | Compile to Go `rune` |
| tuple | `(i32, string)` | Nonempty tuples in type syntax have at least two elements |
| fixed array | `[i32; 4]` | The length is part of the type |
| function | `(i32, string) -> bool` | parameter type list to return type |
| Generic application | `Option[i32]`, `pkg::Box[string]` | Use square brackets |
| channel | `Channel[isize]`, `Sender[isize]`, `Receiver[isize]` | Bidirectional and directional Go channel backends |
| trait object | `dyn Render`, `dyn Iterator[Item = isize]` | A single, non-generic dyn-safe trait; associated types must be bound |
| Associated type projection | `I::Item`, `Self::Output`, `I::IntoIter::Item` | There must be corresponding trait constraints; projections may be chained |

Example of function type:

```goml
let predicate: (i32) -> bool = |value: i32| value > 0;
let combine: (i32, i32) -> i32 = |a, b| a + b;
let action: () -> () = || println("run");
let accepts_empty: (()) -> () = |value: ()| value;
```

`()` is the empty-tuple type. At the left of `->`, `() -> T` is a zero-parameter function type, while `(()) -> T` is a function taking one `()` parameter.

`(value)` in value and pattern is a group, `(value,)` is a single-element tuple. The current type syntax cannot directly annotate single-element tuples: `(T,)` still normalizes to `T`, so the type of such values must be inferred from the local context.

`A -> B -> C` is parsed by right associative analysis. It is recommended to always write function argument lists in parentheses, especially for higher-order functions: `(A) -> (B) -> C`.

The array length must be a non-negative decimal integer in the source code, not a constant expression:

```goml
let pair: [string; 2] = ["left", "right"];
```

The number of array literal elements must match the array type. Empty arrays and empty generic containers usually require type annotations.

### Type aliases

Top-level aliases are transparent and may have type parameters:

```goml
type UserId = u64;

type Pair[T] = (T, T);

pub type Names = Vec[string];
```

Aliases do not create nominally distinct types and recursive alias cycles are rejected. Public aliases can be referenced across packages.

Use a single-field tuple struct when a value needs a distinct nominal type instead of an alias:

```goml
struct UserId(u64);

struct Box[T](T);
```

This form is the newtype pattern. Construct it with `UserId(value)`, access its field with `value.0`, and destructure it with `let UserId(inner) = value;`. The field is private by default. A public newtype that can be constructed and unwrapped across package boundaries marks both the type and field public: `pub struct UserId(pub u64);`.

### Type syntax not currently available

GoML has no Rust reference or lifetime syntax, pointer arithmetic, slice literals, or union types. Use `Ref[T]` for shared mutable storage, `Option[T]` for optional values, and `Slice[T]` or `MutSlice[T]` for read-only or mutable contiguous views. `std::ffi::Ptr[T]` and external Go type aliases can carry nullable Go pointers; they remain distinct from `Ref[T]` and numeric types.

On Linux amd64, `std::os::linux::syscall` accepts numeric machine words and scoped byte-buffer arguments for low-level kernel calls. It does not add pointer casts, pointer arithmetic, or native struct layout to the language.

`A + B` is only usable as a trait bound or supertrait list. The parser reserves `dyn A + B`, but the type checker deliberately rejects multiple dyn bounds in the current object model.

`Self` is only used in the trait signature and the type position of impl; ordinary top-level functions cannot use `Self` as an implicit type parameter.

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

### Constants

Top-level constants require an explicit type and a compile-time expression:

```goml
const BASE: i32 = 0x20;

pub const ANSWER: i32 = BASE + 10;

const NEWLINE: byte = b'\n';
```

Constant names may start with an uppercase letter, a lowercase letter, or `_`; `UPPER_SNAKE_CASE` is the preferred convention. Constant expressions support literals, constructors, references to other constants, unary and binary operators, integer conversion methods, and comptime-capable calls. They may use forward references, but cycles are rejected. Allowed constant types are recursively immutable values: `bool`, numeric types, `string`, `char`, `byte`, tuples, fixed arrays, and structs or enums whose fields satisfy the same rule. `Vec`, `HashMap`, `Ref`, `Channel`, closures, and dynamic values are rejected. Composite constants cannot be used as patterns. Public constants are available through package-qualified paths.

Visible top-level constants can be used as value patterns in `match`, `if let`, `while let`, and `let else`. Both local names and package-qualified names are supported, and constant patterns can be nested or combined with or-patterns:

```goml
const ANSWER: isize = 42;

match value {
    ANSWER | config::FALLBACK => "known",
    _ => "other",
}
```

In these refutable pattern contexts, a visible constant takes precedence over introducing a binding with the same name. Use an explicit alias such as `answer @ _` to bind that name instead. Ordinary `let` bindings and `for` patterns continue to introduce bindings, even when a top-level constant has the same name. Range-pattern endpoints remain integer or character literals.

### Statics and one-time initialization

A top-level `static` has one program-wide storage location. Its binding cannot be reassigned. Its explicit type must be `OnceCell[T]` and its initializer must be exactly `OnceCell::new()`:

```goml
static EMOJI: OnceCell[FrozenVec[FrozenVec[u16]]] = OnceCell::new();

fn emoji() -> FrozenVec[FrozenVec[u16]] {
    EMOJI.get_or_init(|| build_emoji().freeze())
}
```

`OnceCell.get_or_init(init)` runs one initializer and caches its result. Concurrent first callers wait for that initializer and receive the same value. Recursive initialization of the same cell terminates with an error that names the static. A cached `Result` is an ordinary cached value. Statics do not run user-observable destruction at process exit.

Unlike a `const`, a `static` has observable identity. `OnceCell[T]` controls initialization but does not make `T` immutable. Shared caches should therefore expose `FrozenVec` or another immutable value instead of a mutable `Vec`.

### Compile-time evaluation

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

`#[comptime]` marks a non-generic free function as compile-time-capable. The function remains callable at runtime. A compile-time call may call only other `#[comptime]` free functions or the `compile_error(string) -> never` intrinsic. The compiler validates the complete body of every marked function, including branches not taken by a particular invocation. Attributes with arguments, duplicate attributes, generic functions, methods, extern functions, and other declarations are rejected.

A top-level constant initializer is an implicit compile-time context, so `const SIX: isize = factorial(3);` and an initializer wrapped in `comptime { ... }` are equivalent. A top-level constant may produce a recursively immutable tuple, fixed array, struct, or enum in addition to scalar values. A `comptime` expression in ordinary code supports the same reifiable value shapes.

Compile-time code may use local bindings and assignment, blocks, `if`, `match`, `while`, `loop`, restricted `for`, `break`, `continue`, `return`, recursion, direct calls, integer conversion methods, and supported operators. A compile-time `for` accepts only a fixed array or the builtin `isize` ranges `start..end` and `start..=end`; its source and range endpoints are evaluated once, and its pattern must be irrefutable. The deterministic string methods `len`, `byte_len`, `get`, `byte_get`, `byte_slice`, `is_char_boundary`, `starts_with`, `ends_with`, and `contains` are also available. String indexes and slices use byte offsets and reject invalid UTF-8 character boundaries.

Compile-time code cannot capture a surrounding runtime parameter or local. Closures, indirect calls, generic functions, methods other than the integer conversions and string whitelist, trait or dynamic dispatch, general iterators, floating-point computation, `Ref`, `Vec`, `HashMap`, channels, goroutines, extern calls, host I/O, environment access, time, randomness, network access, general type reflection, arbitrary declaration generation, compile-time parameters, value generics, and type-level computation are not supported. The constrained programmable derive interface described below is the only reflection and code-generation facility.

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

At runtime, integer addition, subtraction, and multiplication use the same fixed-width wrapping representation. Integer division or remainder by zero terminates the current process with the generated Go runtime failure. Integer conversion methods use the same narrowing, sign-extension, and signedness rules as compile-time evaluation.

### Structs

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

### Enums

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

Structural variants use field constructs and field patterns:

```goml
match named {
    Message::Named { value } => value,
    _ => 0,
}
```

Patterns may omit the enum qualifier. When the expected pattern type is `Option[T]`, `Some(value)` and `None` resolve to `Option::Some(value)` and `Option::None`. The same rule applies to unit, tuple-like, and struct-like variants of user enums, including enums imported from another package and enums reached through a type alias. Duplicate variant names in unrelated enums are not ambiguous because the expected enum type selects the owner. Nested patterns are resolved recursively, so `Some(Ok(value))` uses the payload type of `Some` to resolve `Ok`.

## Functions and generics

### Top-level functions

```goml
fn add(left: i32, right: i32) -> i32 {
    left + right
}

fn log(message: string) {
    println(message)
}
```

The parameter type cannot be omitted. Omitting `-> ...` is equivalent to `-> ()`. The top-level function name must be unique in the same package, and overloading by parameter type is not supported.

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

Don't write `identity[i32](1)` or Rust's `identity::<i32>(1)`. Generic arguments for an owner type or trait appear before the member name, while arguments owned by a method appear after it:

```goml
let inferred = value.convert(fallback);
let explicit = value.convert::[string](fallback);
let inherent = Box::[i32]::convert::[string](value, fallback);
let trait_call = Convert::[i32]::convert::[string](value, fallback);
```

Top-level functions and methods may introduce their own type parameters. A method's parameters are distinct from the parameters of its enclosing trait or impl and may have their own bounds and `where` predicates. Local named functions do not exist; use closures. Closures do not have generics. Structures, enumerations, traits, and impl blocks can also have type parameters.

GoML monomorphizes generic calls. Recursive generic code must produce a limited number of concrete instances and cannot continually change to a new nested type with each recursive call.

## Blocks, bindings and assignments

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

Unlike Go's `defer`, GoML does not evaluate call arguments when the statement is reached. The complete expression is evaluated at block exit, so reads through `Ref` observe the value at cleanup time. A closure body is a separate control-flow scope. Deferred expressions cannot contain `return`, `break`, `continue`, or `?`, and cleanup during an unrecovered runtime panic is not currently guaranteed. The compiler lowers cleanup to ordinary structured control flow and never emits a Go `defer` statement.

```goml
fn work() -> () {
    let state = Ref::new("open");
    defer println("closing:" + state.get());
    defer println("flush");
    state.set("ready")
}
```

This prints `flush` first, then `closing:ready`.

### `let`, type annotations and shadowing

```goml
let inferred = 42;
let explicit: i64 = 42;
let _ = println("discard explicitly");
let name = "first";
let name = "second";
```

Local variables allow shadowing with the same name. `let _ = expr;` executes explicitly and discards the result.

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

Ordinary local bindings without `mut` cannot be reassigned. Assignment targets include mutable locals, tuple projections, structure fields, and supported index locations. Compound assignment supports `+=`, `-=`, `*=`, `/=`, `%=`, `&=`, `|=`, `^=`, `<<=`, and `>>=`:

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

`Slice[T]` is a read-only view and cannot be assigned by index. `MutSlice[T]` is a mutable bounded view and supports indexed assignment without copying its `Vec[T]` backing storage.

## Expressions and operators

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

### Operator type rules

- The two numeric operands of `+ - * /` must be of the same concrete numeric type; there is no implicit numeric promotion.
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
- Expression `as` is reserved for creating `dyn Trait` values. Numeric conversions use the integer and floating-point methods above; other non-`dyn` targets are rejected. The separate `as` form in `use package as alias;` still selects an import alias.

There are no exponentiation, null coalescing or user-defined operators.

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

## Control flow

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
let number = if let Option::Some(value) = candidate {
    value
} else {
    0
};
```

The `else` of `if let` can also be omitted; in this case the compiler treats the else branch as `()`, so the then branch must produce `()`:

```goml
if let Option::Some(value) = candidate {
    println(value);
}
```

When you need to get values from two branches, you must write `else` explicitly.

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
    Option::Some(inner) => inner,
    Option::None => fallback,
}
```

The rule is shared by `match`, `if let`, `while let`, `let`, `for`, nested patterns, and or-patterns. Existing refutability requirements still apply to `let` and `for`.

A match arm may have a guard of type `bool`. The guard can use bindings introduced by that arm’s pattern:

```goml
match value {
    Option::Some(inner) if inner > 0 => inner,
    Option::Some(_) => 0,
    Option::None => -1,
}
```

A guard is evaluated only after a successful pattern match. Branches with ordinary guards are not counted in exhaustive coverage; only branches without guards and branches with guard literal `true` provide coverage. Branches with guard literal `false` are unreachable. The compiler also warns about branches that are completely covered by earlier patterns.

A branch body can be a single expression or a block. Different branches must be separated by commas, and the comma in the last branch can be omitted:

```goml
match value {
    Option::Some(inner) => {
        let doubled = inner * 2;
        doubled
    },
    Option::None => 0,
}
```

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
        Some(Event::Data { id, value }) => handle(id, value),
        Some(Event::Stop) | None => stop(),
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
while let Option::Some(value) = iterator.next() {
    println(value);
}
```

The loop body must return `()`. Pattern binding is only visible within the loop body.

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

### `?`

The suffix `?` only supports the built-in semantics `Option[T]` and `Result[T, E]`:

```goml
fn plus_one(value: Option[i32]) -> Option[i32] {
    Option::Some(value? + 1)
}

fn read_number(flag: bool) -> Result[i32, string] {
    let value = parse_number(flag)?;
    Result::Ok(value)
}
```

When using `?` with `Option[T]`, the nearest function or closure must return `Option[_]`. Use of `Result[T, E]` must return `Result[_, E]` with the same error type; there is currently no Rust `From`-style error conversion. `?` will evaluate the operand once.

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
| Enum tuple variant | `Option::Some(value)` or contextual `Some(value)` |
| Enum struct-like variant | `Message::Named { value }` or contextual `Named { value }` |
| Array, Vec, Slice or MutSlice | `[first, second]`, `[first, .., last]` |
| rest binding | `[first, middle @ .., last]` |
| Alias | `whole @ Option::Some(value)` |
| or-pattern | `Color::Red \| Color::Blue` |
| Range | `0..10`, `'a'..='z'` |
| Nested pattern | `Result::Ok((key, value))` |

### Refutability and binding

`let` and `for` require that the pattern be irrefutable, that is, the type at that position must succeed. `match`, `if let` and `while let` can use refutable patterns:

```goml
let (left, right) = pair;
if let Option::Some(value) = candidate {
    println(value);
}
```

Irrefutability is judged jointly by type and sub-pattern rather than by surface syntax alone. For example, the destructuring of a single-variant enum is irrefutable, as is the complete destructuring of a fixed-length array of known length; [first, ..] of a Vec[T] may fail with an empty Vec.

Variables with the same name cannot be bound repeatedly in the same pattern branch. Each alternative of the or-pattern must be bound to exactly the same set of variables, and the corresponding variables must be of the same type:

```goml
match value {
    Either::Left(shared) | Either::Right(shared) => shared,
}
```

Pattern binding is only visible in the corresponding `let` subsequent scope, `for` / `while let` loop body, `if let` then branch, or `match` branch guard and branch body.

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

### Alias and or-pattern

`name @ pattern` binds the entire matched value to `name` while continuing to match the inner pattern:

```goml
match value {
    whole @ Option::Some(inner) => use_both(whole, inner),
    Option::None => fallback(),
}
```

`@` is more tightly bound than `|`. Therefore `whole @ A | B` resolves to `(whole @ A) | B`, and usually an error will be reported because the variables bound to the two alternatives are different. To make the alias cover the entire or-pattern, you must write parentheses:

```goml
whole @ (Either::Left(value) | Either::Right(value))
```

Or-pattern alternatives are tried from left to right and can be nested in tuple, struct, enum, and sequence patterns.

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

## Closures and function values

### Closure syntax

```goml
let add = |left: i32, right: i32| left + right;
let increment = |value: i32| {
    value + 1
};
let greet = || println("hello");
```

Closure parameter types can be inferred from the expected function type or the calling location; when inference is unstable, it should be explicitly marked:

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

Top-level functions, closures, and tuple variant constructors can all be passed or returned as function values.
Closures can be nested and return another closure that captures the environment, but each resulting closure expression still has only one concrete function type.

## Trait and impl

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

Method type arguments are normally inferred. Write `box_value.map::[string](convert)` when explicit arguments are needed, or `Box::[i32]::map::[string](box_value, convert)` in associated form.

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

Trait bounds make the corresponding methods available on generic values. Supertraits and associated type bounds also participate in method resolution as implied constraints.

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

Every associated type declared by the trait must be bound exactly once. The bracket grammar reserves positional trait arguments followed by associated bindings, such as `dyn Consumer[string, Error = IoError]`; positional arguments after the first `Name = Type` binding are rejected. Generic trait objects are still rejected, so the positional form is reserved for forward compatibility rather than enabled today.

Current dyn-safe conditions:

- Traits cannot have type parameters;
- Each method must have a first receiver parameter of exactly type `Self`;
- Direct `Self` cannot appear in other parameters or return types, while a bound projection such as `Self::Item` is allowed;
- Methods cannot declare type parameters.

Current limitations:

- Generic trait object is not supported;
- Multiple bounds such as `dyn Read + Close` are parsed for forward compatibility but rejected by the type checker;
- Pattern matching on `dyn Trait` is not supported.

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

All fields or variant payloads used by a derive must support that trait. For a generic definition, the generated impl constrains each distinct participating field or payload type that mentions a type parameter. An unused phantom parameter receives no constraint. Compiler-owned runtime, intrinsic, and lang-item attributes remain unavailable to ordinary projects. User projects may use the `go_ffi`, `go_type`, `go_method` and `go_interface` attributes described in the next section.

The eight prelude derives are supplied by verified handlers in the toolchain's builtin sources. They are not hard-coded code generators in the compiler. The `PartialOrd` and `Ord` traits themselves live in `std::cmp`; import that package when using their derives, and import the traits for unqualified bounds and method calls. Standard-library and third-party derives use the same handler and artifact mechanism, but they are not added to the prelude: import the trait or its package before using the derive name.

Derived `Debug` provides `Debug::debug(value)` and the `.debug()` method. It formats structs with their type and field names and enums with their type and variant names. Primitive field values use their ordinary textual representation, while nested values recursively use `Debug`.

`PartialEq` compares struct fields in declaration order and compares enum variants and payloads structurally. `Eq` generates only the marker implementation and does not implicitly derive `PartialEq`. The usual total-equivalence spelling is therefore `#[derive(PartialEq, Eq)]`.

`PartialOrd` and `Ord` use lexicographic field order. Enum variants are ordered by declaration order, followed by their payload fields. `PartialOrd` immediately propagates `None`; `Ord` returns a definite `Ordering`. Neither derive implicitly generates its supertraits, so the usual total-order spelling is `#[derive(PartialEq, Eq, PartialOrd, Ord)]`.

`Hash` combines fields in declaration order and includes the enum variant discriminant. It can be derived independently, although hash-map keys still require both `Eq` and `Hash`.

`Default` initializes every struct field with `Default::default()`. For an enum it always selects the first declared variant and recursively defaults that variant's tuple or named payload. An empty enum cannot derive `Default`. Reordering enum variants therefore changes both the derived default value and the derived ordering; `#[default]` is not supported.

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

A public `#[comptime_derive(Name)]` handler exports the derive name `Name`, independently of its implementation function name. Export names are unique within one package and cannot contain `::`. The named form requires `pub`; private `#[comptime_derive]` functions remain implementation helpers. The unnamed public form remains available and exports the handler function's short name.

Derive names have their own namespace, but ordinary `use` declarations populate it alongside the type and trait namespaces. For example, after `use std::serde;`, `use serde::Serialize;` makes both the `Serialize` trait and a derive export named `Serialize` available as `Serialize`, so `#[derive(Serialize)]` works without another import form. An import alias applies to both namespaces: `use serde::Serialize as DataSerialize;` permits `#[derive(DataSerialize)]`. A package import permits the qualified spelling `#[derive(serde::Serialize)]`. Public re-exports preserve the derive handler's definition identity, so facade packages can use `pub use` to expose a standard-library or third-party derive without copying its implementation.

`std::serde` owns the format-independent `Serialize`, `Deserialize`, `Serializer`, and `Deserializer` traits and the two derives. Typed formats use `Serialize::serialize` and `Deserialize::deserialize`: the format handle is a method-level generic parameter, so monomorphization produces ordinary static calls without a runtime serializer trait object, Go interface dispatch, or runtime reflection. A derived struct emits fields directly and a derived enum emits its declaration index, wire name, shape, and payload directly. Derived named-field decoding accepts arbitrary field order, skips unknown fields, and rejects duplicate or missing fields. Struct fields, enum variants, and named variant fields support `#[serde(rename = "wire_name")]`.

`serde::Value` remains the explicit dynamic data model. It retains exact signed and unsigned integer widths, floating-point widths, enum declaration indexes, field order, and variant shape. `Sequence`, `Tuple`, and `Optional` are distinct, and `Value::Number(string)` represents a textual number whose destination type is not yet known. `serde::to_value` and `serde::from_value` connect typed values to this model through `ValueSerializer` and `ValueDeserializer`; both return `Result`, and using them intentionally constructs or consumes a complete value tree. Unit, booleans, strings, chars, numeric primitives, `Value`, `Vec[T]`, `Option[T]`, and two- or three-element tuples have standard direct implementations.

`Serialize` and `Deserialize` have no `Value` fallback methods and `Deserialize` does not require a runtime schema. Handwritten implementations must implement the generic direct protocol. `Schema` remains available only for explicit dynamic operations such as `bincode::decode_value`; it is not derived or consulted by typed decoding.

This design removes the mandatory intermediate tree, but it is not an absolute zero-cost or Rust-style zero-copy guarantee. GoML has GC rather than ownership and `Deserialize<'de>` lifetimes, format buffers and decoded strings still allocate when required, and the Go compiler decides which static calls it inlines. The direct path means no data-model tree or dynamic dispatch is inherently required; allocations must still be measured for each format and value shape.

```goml
use std::serde;
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
struct User {
    #[serde(rename = "display_name")]
    name: string,
    active: bool,
}
```

`std::json` supports two deliberately separate modes. The value mode uses `json::Value`, `json::parse`, and `json::encode` for schema-free inspection and editing. JSON numbers remain their exact source text in `Value::Number`. In the typed mode, `json::try_to_string` and `from_string` write and consume JSON directly through the streaming serde traits; neither operation first builds a `json::Value` or `serde::Value` tree. `json::to_value` and `from_value` return `Result` and are the explicit bridge to the dynamic JSON model. Numeric range and destination-width checks happen while deserializing into the requested type.

`json` publicly re-exports the shared `Serialize` and `Deserialize` traits and derive handlers, so either `use serde::Serialize` or `use json::Serialize` selects the same implementation identity. The JSON serializer emits struct fields in source order. Direct maps use JSON objects and therefore require keys whose direct representation is a string or char; `try_to_string` returns a recoverable error for other key types. The deserializer accepts any field order, recursively skips unknown values, rejects duplicate and missing fields, and rejects trailing input. Typed errors retain the byte offset and nested struct, sequence, map, or enum path. JSON uses externally tagged enums: a unit variant is a string, a tuple variant is an object whose value is an array, and a struct-like variant is an object whose value is another object.

```goml
use std::json;
use json::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
struct User {
    name: string,
    active: bool,
}

fn encode_user(value: User) -> Result[string, string] {
    json::try_to_string(value)
}

fn decode_user(input: string) -> Result[User, string] {
    json::from_string(input)
}
```

A qualified custom derive must have the form `package_alias::export_name`, where `package_alias` is introduced by a `use` in the same source file. A bare library derive must be explicitly imported. Unimported canonical package paths and handlers visible only through a transitive dependency are rejected. If two explicit imports or a prelude derive and an explicit import provide the same bare name, the compiler reports ambiguity and requires a qualified name or import alias.

`#[comptime_derive]` and `#[comptime_derive(Name)]` are valid only on non-generic free functions. A public handler must have the exact signature `(DeriveInput) -> DeriveOutput`. Private functions with the unnamed attribute are compile-time-only helpers and are included in the interface when reachable from a public handler. A derive handler may call those helpers and ordinary `#[comptime]` functions, but it cannot be called from runtime code or ordinary value `comptime`. Derive handlers are not exported as runtime functions.

The compiler resolves handlers from already compiled dependency interfaces. A handler cannot be defined and applied within the same package compilation. Put reusable handlers and their generated traits in a separate package. The target may be a generic struct or enum; the generated impl inherits its type parameters. `derive_output_add_predicate` and `derive_output_add_call_site_predicate` add the bounds required by generated methods.

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

The structured attribute handle exposes `meta_attribute_name`, `meta_attribute_text`, `meta_attribute_has_argument_list`, `meta_attribute_argument_count`, `meta_attribute_argument_kind`, `meta_attribute_argument_name`, `meta_attribute_argument_value_kind`, and `meta_attribute_argument_text`. Argument kind is `ident`, `path`, `string`, or `named`. Named arguments use `name = value`, where the value may be an identifier, path, or string. `argument_name` returns the left-hand name, `argument_value_kind` describes the right-hand value, and `argument_text` returns its decoded value.

The structured output API provides opaque `MetaAttribute`, `MetaType`, `MetaExpr`, `MetaPattern`, `MetaArm`, `MetaBlock`, `MetaParamList`, `MetaGenericList`, `MetaMethod`, and list handles. Constructors use the `meta_type_*`, `meta_expr_*`, `meta_pattern_*`, `meta_arm*`, `meta_block_*`, `meta_param_list_*`, and `meta_generic_list_*` families. `meta_method` creates a concrete trait method; `meta_method_generic` creates a method with explicit type parameters and bounds. A handler creates its single result with `derive_output_new` or `derive_output_new_call_site`, adds trait predicates and methods, and returns it.

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
derive_output_add_predicate(output, type, trait_name) -> ()
derive_output_add_call_site_predicate(output, type, trait_name) -> ()
derive_output_add_method(output, method) -> ()
```

`meta_expr_unary` uses operator numbers `0..2` for `-`, `!`, and `~`. `meta_expr_binary` uses operator numbers `0..17` for `+`, `-`, `*`, `/`, `%`, `&`, `|`, `^`, `<<`, `>>`, `&&`, `||`, `<`, `>`, `<=`, `>=`, `==`, and `!=`, respectively. `meta_expr_integer` accepts a normalized integer literal string plus its exact integer type, so builders can represent values outside the host `isize` range. `meta_expr_cast` accepts only a `dyn` target type; generated numeric conversions should use `meta_expr_method_call`. `meta_expr_trait_call` resolves the trait in the handler's defining package and builds a static trait method call. `meta_type_equal` compares structural type identity. `meta_type_kind` returns `primitive`, `named`, `tuple`, `application`, `array`, `function`, or `dyn`; shape-specific accessors reject other kinds. List handles are mutable only through their matching `push` operation and remain local to one derive evaluation.

Unqualified names passed to `derive_output_new`, `derive_output_add_predicate`, `meta_type_named`, `meta_expr_call`, and `meta_generic_list_add_bound` resolve in the handler's defining package. Their `_call_site` variants resolve in the target package. `derive_fresh_name` should be used for generated local bindings that must not collide with user names. The `*_target_*` builders construct or match the annotated item by compiler identity and should be preferred over spelling its name manually.

The result is restricted to one trait `impl` for the annotated type. It cannot create types, traits, functions, constants, modules, imports, inherent impls, extern declarations, attributes, associated types, or raw tokens. Generated method type parameters and trait bounds are supported. The generated impl is processed by ordinary name resolution, orphan and coherence checks, type checking, monomorphization, and backend lowering. Duplicate or invalid generated implementations are regular compiler diagnostics.

Handlers are deterministic and have no host access. Imported derive CTIR is verified as untrusted artifact data. Evaluation uses the ordinary compile-time limits plus a limit of 100,000 metadata and syntax-builder operations. Failures are anchored to the requesting derive attribute and include the compile-time derive call stack.

## Go FFI

The runnable [file reading example](../examples/ffi-file-read/README.md) combines raw pointers, methods, shared byte buffers, partial reads, error identity, explicit Close and Go↔GoML result transport.

The typed Go FFI binds a top-level GoML declaration to an exported package-level Go function:

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

The first attribute argument is the Go import path and the second is an exported ASCII Go identifier. The GoML function name is local and may differ from the Go symbol. Add `pub` before `extern fn` to expose the binding through another GoML package; interface and Core artifacts preserve the Go import path and symbol.

Bind a Go type with `#[go_type("time", "Duration")] pub extern type Duration;`. The GoML name may differ from the Go object name. Declarations cannot have a body or an `= Type` target. Go defined types retain their identity: `Duration` differs from `i64`, while two GoML packages binding `time.Duration` share the same type. Go aliases normalize to their target type. Values can pass through ordinary GoML functions and typed Go foreign calls; generated Go uses the original type and preserves its method set. GoML field access, struct construction and pointer dereferencing are not yet supported for these values.

```goml
#[go_type("time", "Duration")]
extern type Duration;

#[go_ffi("time", "Sleep")]
extern fn sleep(duration: Duration) -> ();

fn wait(duration: Duration) -> () {
    sleep(duration)
}
```

A Go alias such as `type Handle = *Cell` can be bound with `#[go_type("example.com/shim", "Handle")] extern type Handle;`. The generated type is `*Cell`, including its nil value and pointer identity. Passing it through functions, closures or containers does not copy the pointed-to object. `Option::Some` containing a nil pointer remains distinct from `Option::None`. GoML cannot construct such a value from an integer or a `Ref`, or use struct layout syntax to create it. `use std::ffi;` exposes `ffi::Ptr[T]`, a transparent spelling for a raw Go pointer to T. `ffi::null()` constructs a nil pointer with its type inferred from context, and `ffi::is_nil(value)` tests it. Non-null pointers enter through explicit Go function bindings. For example, `let empty: ffi::Ptr[Cell] = ffi::null();` and `ffi::is_nil(empty)` use the same representation as a Go alias of `*Cell`. These functions are also ordinary function values. This uses the existing external type declaration grammar, with no `*T` source type syntax.

Bind a Go method with `#[go_method("Method")] extern fn`; its first parameter is the receiver. Declarations must be monomorphic and have exactly one ABI attribute. The method name must be an exported ASCII Go identifier. Go checks the receiver's actual method set: pointer-only methods require an appropriate pointer receiver, while value methods can also be called through pointers. There is no implicit address-taking of a copied value. Interface and promoted methods follow the same Go rules. Bindings can be public, imported across packages, or passed as ordinary function values, and are runtime-only, including when loaded from artifacts.

```goml
use std::ffi;

#[go_type("example.com/shim", "Cell")]
extern type Cell;

#[go_method("Read")]
extern fn read(cell: ffi::Ptr[Cell]) -> i64;

#[go_method("Write")]
extern fn write(cell: ffi::Ptr[Cell], value: i64) -> ();
```

These bindings emit Go method expressions such as `(*shim.Cell).Read(cell)`. A nil receiver reaches the Go method unchanged, including any panic or special nil behavior the method defines. Methods do not automatically allocate objects, close resources, or convert errors into `Result`. Obtaining non-null values still requires an explicit Go binding.

`ffi::Error` is the original nullable Go `error` interface. `ffi::nil_error()` constructs a nil interface; `ffi::error_is_nil(value)` tests interface nilness. An interface holding a typed nil pointer is not nil. Errors retain their original object and chain through functions, closures and artifacts, and Go aliases of `error` have the same type. Raw `(T, error)` returns remain tuples; `Option::Some` containing a nil error remains distinct from `None`.

`ffi::NonNilError::from_error(error)` returns `None` for a nil interface and `Some` for a non-nil interface, including typed-nil errors. Its private storage cannot be constructed directly; `as_error()` returns the original error. Consequently `Result[T, ffi::NonNilError]` cannot accept `Err(ffi::nil_error())`. This wrapper does not call `Error()` or inspect concrete error payloads.

`ffi::Outcome::new(value, error)` preserves both results in public `value` and `error` fields, including a partial value returned alongside a failure. The explicitly named `into_result_discarding_value_on_error()` adapter returns `Result[T, ffi::NonNilError]`: it keeps the value only on success and discards it on failure. No such conversion happens automatically. For example:

```goml
let (count, error) = read(buffer);
let outcome = ffi::Outcome::new(count, error);
let partial_count = outcome.value;
let result = outcome.into_result_discarding_value_on_error();
```

Reverse conversion is also explicit. `ffi::Outcome::from_raw_result(result, failure_value)` accepts `Result[T, ffi::Error]` and returns `Result[ffi::Outcome[T], ffi::BoundaryError]`. An `Ok(value)` produces that value with a nil Go error. An `Err(error)` uses the caller-supplied failure value and preserves the error. If that error is a true nil interface, conversion fails with `BoundaryError::NilError`, whose message is `Go FFI boundary rejected Err(nil)`. Typed-nil errors remain failures. The failure-value argument is evaluated normally even when the input is `Ok`.

`ffi::Outcome::from_result(result, failure_value)` accepts `Result[T, ffi::NonNilError]` and applies the same boundary check. `outcome.into_tuple()` returns the exact `(value, error)` pair for an explicit Go call; it does not normalize partial results or nil values. None of these adapters invent a default failure value or implicitly turn a contract error into Go success.

`ffi::error_matches(value, target)` uses Go `errors.Is`, including wrapped chains, custom `Is` methods and Go nil matching semantics. `NonNilError::matches(target)` provides the same operation for a checked wrapper. Matching compares the original error objects and chains, never message strings.

`NonNilError::message_bytes()` invokes the original Go `Error()` method and returns a fresh `std::bytes::Bytes` copy of its bytes. `message()` returns `Result[string, std::utf8::Utf8Error]`, accepting valid UTF-8 and rejecting invalid bytes without replacement. Message access does not replace the stored error. Typed-nil interfaces can enter the checked wrapper, so the underlying Go method's nil handling, side effects and panics remain its own. These APIs do not catch Go panics or implicitly convert errors to strings.

Package compilation queries Go metadata for every declaration, including private and unused declarations. Interface and Core artifacts preserve this metadata, and cached inputs are rechecked against the current Go declarations. Changed declarations require rebuilding the GoML package from source. `--ffi-check off` cannot skip external type resolution or artifact type revalidation. Generic declarations such as `#[go_type("example.com/shim", "Box")] extern type Box[T];` accept concrete applications such as `Box[i64]`. The Go type checker validates the original Go constraints, including constraints on alias parameters that disappear when the alias is expanded. Imported applications are checked too, and concrete instance requests are saved for artifact revalidation. Concrete instance validation requires that a variable of the resolved Go type is legal. Constraint-only interfaces can be described in declaration metadata but cannot be used as runtime instances. Applications involving unresolved GoML generic parameters currently report that specialization is required; using `Box[T]` inside a generic GoML function or alias is not yet supported. Aliases to Go pointers are supported when their pointee has a supported representation. Aliases requiring other unsupported raw representations, such as raw Go channels, produce diagnostics. Query and LSP checking also resolve external type declarations and validate concrete instances, including cross-package references and unsaved GoML source changes. The language server uses its adjacent `goml-go-meta` helper and the GoML module root as its Go build context, with the same offline, read-only Go loading policy. Missing tools and metadata failures produce recoverable diagnostics. Native foreign-call signatures, interface adapters, and exports still receive their complete validation during module compilation.

The foreign-call ABI supports values whose generated Go representations are directly assignable. Raw boundary types and adapters are detailed in the following sections:

| GoML type | Go representation |
| --- | --- |
| `bool`, numeric primitives, `string` | Corresponding Go primitive |
| External named type with concrete arguments | Original Go package type, without a new defined type |
| `ffi::Ptr[T]` | Nullable Go pointer `*T` |
| `ffi::Error` | Nullable Go `error` interface |
| `ffi::String` | Go `string` preserving arbitrary bytes |
| `ffi::Rune` | Go `rune` / `int32` without scalar validation |
| `ffi::RawSlice[T]` | Nullable Go slice `[]T` |
| `ffi::RawMap[K, V]` | Nullable Go map `map[K]V` with comparable keys |
| `char` | `rune` / `int32` |
| `byte` | `byte` / `uint8` |
| `[T; N]` | `[N]T` |
| `Slice[T]` | `[]T` |
| `MutSlice[T]` | `[]T` |
| `ffi::Func[F]` | Nullable Go function with signature F; unit/flat tuple results use zero/multiple Go results |
| `Channel[T]` | `chan T` |
| `Sender[T]` | `chan<- T` |
| `Receiver[T]` | `<-chan T` |
| Direct `dyn Marker` parameter or single return | Go `any` through the trait object's `data` field |
| `()` return | Go function with no result |
| `(A, B, ...)` return | Multiple Go results in the same order |

Tuple types are supported only as the complete return type. Tuple elements and array, slice, or channel elements must themselves be FFI-safe. Parameters cannot be `()`.

`dyn Marker` is supported only as a direct parameter or as the complete single return type. `Marker` must be strictly empty: it cannot declare generic parameters, predicates, supertraits, associated types, or methods, and the `dyn` type cannot have arguments, associated-type bindings, or additional bounds. Artifact validation resolves a local unqualified marker name in the declaring function’s package, including private markers. The wrapper passes the object's `data` field to Go as `any` and wraps a returned Go value back into the empty marker object. Empty marker objects are not yet supported inside tuples, arrays, slices, or channels.

The declaration must be monomorphic. Project `goml check`, `build`, `run`, and `test` default to `--ffi-check required`: the adjacent `goml-go-meta` helper uses Go package loading and type checking to validate the actual generated call, including unused private declarations. Missing tools, unavailable packages, and incompatible signatures produce errors before successful package artifact publication. Cached GoML artifacts are checked again against current Go sources; verification success is not reused across commands. Binding diagnostics include the GoML declaration location when its source is readable, along with the Go target and parameter-name mapping. When source is unavailable, the package/binding identity and Go diagnostic remain available. Go remains the authority for symbol existence and assignability. In particular, a Go named type such as `time.Duration` is not interchangeable with a GoML `i64` parameter even when its underlying representation is the same.

Use `goml check --ffi-check off` (also accepted by `build`, `run`, and `test`) to explicitly skip Go signature validation. Commands with foreign declarations emit an `ffi-unverified` warning in this mode. This does not guarantee that the final Go build succeeds. Checks with no foreign declarations do not require Go or the helper. Standalone compiler commands and LSP queries do not yet perform this project validation.

Required validation and the final executable build share the selected Go executable, target, flags, module policy, and generated caller directories. Validation honors `internal` import restrictions, uses readonly module resolution, disables workspace mode and automatic toolchain downloads, and does not fetch packages or modify `go.mod`/`go.sum`. Prepare required Go dependencies separately. A valid signature does not prove Unicode, mutability, lifetime, or concurrency contracts. A unit-returning legacy binding may still discard Go results; ordinary calls also retain Go's generic inference and variadic-call rules.

`Vec`, `Ref`, `HashMap`, `Option`, `Result`, user structs and enums, nonempty trait objects, ordinary GoML function values, nested tuples, generic function declarations, implicit callback conversion, automatic Go object lifetime management, and automatic `error` conversion are not supported by this ABI. Write a small Go shim with an exported function and FFI-safe parameters when adapting such an API:

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

When the GoML module root contains `go.mod`, `goml build`, `goml run`, and test linking invoke Go in module mode, so local shim packages and declared Go module dependencies can be imported. Go workspace mode remains disabled. Without `go.mod`, builds retain the existing module-off behavior. Module-mode Go builds always execute and delegate dependency and source freshness to Go's own build cache, so changes to `go.mod`, `go.sum`, and local `.go` shims are observed.

The package-function form calls Go from GoML; method calls use the explicit `go_method` form described above. Dynamic symbol lookup and C ABI interoperation require separate mechanisms.

### Go interface wrappers

In module compilation, `go_interface` on a top-level trait generates a wrapper holding an explicitly bound native Go interface and forwarding methods:

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

The first argument is the raw type path, which may name an imported external type. The second is an unqualified wrapper name; its visibility follows the trait. Named string arguments map every trait method to exactly one exported Go method. The completed Go interface method set must be covered exactly, including embedded methods. Missing, duplicate, unexported or mismatched methods produce diagnostics at the adapter declaration. Generated type/helper names and reserved wrapper members cannot collide with user declarations.

`ReaderAdapter::from_go(raw)` stores the original interface value in a private field; `wrapper.into_go()` returns it. These preserve nil interfaces, typed-nil receivers and the identity of referenced Go objects. Forwarded calls retain Go receiver behavior, including nil-receiver handling and panics. No receiver is replaced and no automatic close is added. The raw Go interface remains distinct from `dyn Reader`; use `wrapper as dyn Reader` for an explicit GoML trait conversion.

Method signatures use explicit raw boundaries: `ffi::String`/`ffi::Rune` for text and `ffi::RawSlice` for slices. Apply copy/share adapters explicitly at the caller. Every Go result is preserved: `()` means no results, and `(isize, ffi::Error)` retains both a partial count and its original error. Ordinary GoML text, read-only slices, implicit closures and runtime containers are not automatic interface method conversions.

`ReaderAdapter::from_trait(value: dyn Reader)` creates a wrapper backed by a generated Go bridge. Use `ReaderAdapter::from_trait(implementation as dyn Reader).into_go()` to pass a GoML implementation to Go. Typed forwarding closures retain the supplied trait object after the creating function returns. Go calls preserve shared captured state and every return component; panics propagate in the same goroutine. The generated bridge includes a compile-time Go interface-satisfaction assertion. The resulting interface is non-nil, including when the supplied trait object wraps a nil raw interface. `from_go`, `into_go` and `from_trait` are reserved wrapper method names. Subscription removal and resource closing remain explicit caller responsibilities.

The current adapter requires a concrete Go interface and a non-generic GoML trait with instance methods, without supertraits, associated types or additional method predicates. Constraint-only interfaces are rejected as runtime values. Both `--ffi-check required` and `--ffi-check off` require current interface metadata and complete typed method validation before publishing a generated implementation. Cached interfaces retain the external metadata for revalidation. This feature currently follows the module-only external-type path; standalone and source-only LSP analysis do not yet resolve Go metadata. For explicit allowlisted raw bindings and finite generic instances, use `goml bind-go` as described below; it does not infer interface adapter mappings.

### Raw Go function values

`ffi::Func[F]` is a nullable raw Go function value, distinct from a GoML closure with signature F. For example, `ffi::Func[(i64) -> (i64, ffi::Error)]` represents Go `func(int64) (int64, error)`, and `ffi::Func[() -> ()]` represents `func()`. F must resolve to a function signature with FFI-safe parameters and results; functions cannot be raw map keys. A Go alias of a supported non-variadic function type normalizes to this representation.

`ffi::func_nil()` constructs nil with its signature inferred from context, and `ffi::func_is_nil(value)` checks it. Values returned by Go can pass through packages, generic functions, containers and captured GoML closures without losing their Go closure environment. `Option::Some` containing a nil function remains distinct from `Option::None`. Go bindings can receive these raw values and invoke them in Go. Raw functions are not directly callable with GoML call syntax, and neither implicit closure conversion nor struct construction is supported.

```goml
use std::ffi;

fn empty_callback() -> ffi::Func[(i64) -> i64] {
    ffi::func_nil()
}
```

`ffi::func_from_closure(value)` explicitly converts a GoML function or closure into a non-nil raw Go function. `ffi::func_to_closure(value)` returns `None` for nil, or `Some(closure)` for a callable GoML closure. Both directions preserve captured state; converting two closures with the same signature keeps their environments distinct. Captured `Ref` cells retain their ordinary sharing behavior.

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

Conversions generate typed Go method values that retain the original function and its environment. GoML closures still undergo lambda lifting. Unit and flat tuple results are converted explicitly between GoML storage and Go's zero/multiple results. Raw errors and other supported raw values are preserved. Callback text values must use `ffi::String` or `ffi::Rune`; ordinary GoML `string`/`char`, including in arrays, slices, channels or raw maps, are rejected during adapter validation. Perform checked Unicode conversions inside the callback and choose an explicit failure result shape.

The adapters add no synchronization, global registry, reflection or panic recovery. A panic propagates with its original value along the same goroutine, including through both conversion directions, and Go defers run during stack unwinding. It does not become an error result. Recovery requires an explicit caller-owned Go adapter using defer/recover in that goroutine; it cannot recover a panic in another goroutine or intercept process exit. Go may retain a converted callback after its registering GoML call returns, invoke it on another goroutine, or reenter GoML recursively. Captured state remains reachable through the Go function value. Concurrent mutation needs ordinary synchronization; conversion does not make a `Ref` thread-safe.

The runnable [callback subscription example](../examples/ffi-callbacks/README.md) provides explicit unregister, gated asynchronous dispatch, concurrent atomic updates, recursive reentry and registration stress coverage. Its callback reference is removed by unregister; already selected invocations may still run, so callers wait for their batches before disposing of captured resources. GC reachability and registration lifetime are separate. The example and its generated Go program are tested with the race detector. GoLibrary function-value exports remain unsupported; the [callback benchmark](../examples/ffi-callback-bench/README.md) measures invocation and escaping construction separately.

Executable, test and library module links embed a versioned JSON index in the generated Go string constant `_goml_source_origins`, relating surviving functions and callback adapters to enclosing GoML declarations. Shared adapters may have several source entries. External-call wrappers also identify the actual extern declarations used, including `go_method`; unused aliases of the same native target are excluded. The index records canonical symbols, the source package, package-relative source files and declaration-name positions, including explicitly defined impl methods; it is not an expression-level stack-trace line table.

The index has `version: 1`, `scope: "enclosing-declaration"` and an `entries` array. Entries contain `generated_symbol`, `source_symbol`, `source_package`, `source_file`, `start`, `end`, `line` and `column`; callback methods use `ReceiverType.Apply` as their generated symbol. Decode the Go string literal before parsing the JSON.

### Raw Go strings

`std::ffi::String` represents arbitrary Go string bytes and is distinct from the GoML Unicode `string` type. Go aliases of `string` normalize to this raw type. It can cross foreign calls, package interfaces, generic functions and captured closures without changing its bytes. Neither implicit conversion nor direct construction is supported.

| Explicit adapter | Result and policy |
| --- | --- |
| `ffi::string_from_bytes(bytes::Bytes)` | `ffi::String`; copies bytes, including invalid UTF-8 and NUL |
| `ffi::string_from_text(string)` | `ffi::String`; preserves validated text |
| `ffi::string_bytes(ffi::String)` | `bytes::Bytes`; returns an independent mutable byte copy |
| `ffi::string_text(ffi::String)` | `Result[string, utf8::Utf8Error]`; validates UTF-8 without replacing invalid bytes |

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

Legacy foreign bindings declared with GoML `string` retain their existing trusted text contract. Use `ffi::String` when a Go producer can return arbitrary bytes. GoLibrary exports may use `ffi::String` to select the raw-byte policy. Ordinary GoML `string` and `char` exports remain unsupported; use explicit checked adapters at the boundary. Raw Go strings use ordinary qualified type and function syntax; no new grammar is introduced.

### Raw Go runes

`ffi::Rune` is a transparent alias of `i32`, matching Go's `rune`. Raw values may be negative, surrogate code points or larger than the Unicode scalar range. `ffi::rune_from_char(char) -> ffi::Rune` preserves a valid character. `ffi::rune_char(ffi::Rune) -> Result[char, ffi::BoundaryError]` checks the scalar range and returns `InvalidRune(original_value)` for invalid input. It never substitutes a replacement character. These adapters use ordinary type alias and function syntax.

### Raw Go slices

`ffi::RawSlice[T]` is an alias of `MutSlice[T]` and directly represents a Go `[]T` header. Go slice aliases normalize to this type. Header assignment copies length, capacity and the backing-storage reference; it does not copy elements. Existing checked indexing and subslicing methods remain available.

| Adapter or operation | Policy |
| --- | --- |
| `ffi::slice_nil::[T]()` / `ffi::slice_is_nil(value)` | Preserve and inspect nil separately from a non-nil empty slice |
| `ffi::slice_capacity(value)` | Return the original header capacity |
| `ffi::slice_append(value, item)` | Return Go append's new header; retain it to observe the new length |
| `ffi::slice_shared(vector)` | Share mutable element storage with a snapshot of the Vec header |
| `ffi::slice_from_vec_copy(vector)` / `ffi::slice_copy(value)` | Copy elements into independent storage |
| `ffi::slice_to_vec_copy(value)` | Copy elements into a new Vec |
| `ffi::slice_trusted_readonly(value)` | Share storage through a read-only `Slice[T]` view under the binding author's promise |

Copies are shallow: copying a slice of pointers isolates the element slots, while the pointed-to objects remain shared. Copy adapters produce a non-nil empty allocation for empty input and do not preserve nil identity or spare capacity. Appending may reuse backing storage or allocate another backing array; neither Go append nor Vec append updates previously copied headers. Go writes are visible through aliases sharing the same backing storage. No adapter adds synchronization.

`trusted_readonly` is an explicit author promise that the storage is suitable for read-only use, including coordinating any other writers; Go signatures cannot prove this promise. The conversion does not freeze the original slice or recursively freeze referenced objects. Copy first when independent element storage is needed. Vec and HashMap still have no implicit foreign conversion.

```goml
use std::ffi;

fn append_header() -> bool {
    let original: ffi::RawSlice[i64] = ffi::slice_nil();
    let updated = ffi::slice_append(original, 7);
    ffi::slice_is_nil(original) && updated.len() == 1
}
```

These APIs reuse ordinary aliases, generic functions and existing slice operations; no new grammar is introduced. GoLibrary slice exports remain unavailable.

### Raw Go maps

`ffi::RawMap[K, V]` represents Go `map[K]V` directly and is distinct from GoML `HashMap[K, V]`. Go aliases of map types normalize to this representation. It preserves nil and shared storage through assignments, foreign calls, package interfaces, generics and captured closures. Go operations on an alias observe the same entries.

`ffi::map_nil::[K, V]()` produces a nil `ffi::RawMap[K, V]`. `ffi::map_is_nil(value) -> bool` distinguishes nil from an allocated empty map. An `Option::Some` containing nil remains distinct from `None`. The raw type has no public struct constructor and no implicit conversion to or from `HashMap`; their storage, hashing and equality contracts differ. Raw map syntax uses ordinary qualified generic aliases and functions.

| Operation | Semantics |
| --- | --- |
| `ffi::map_new::[K, V]()` | Allocate a non-nil empty raw map |
| `ffi::map_len(value)` | Entry count; nil has length zero |
| `ffi::map_get(value, key)` | `Option[V]`; distinguish absence from a present zero or nil value |
| `ffi::map_set(value, key, item)` | Insert or replace using Go equality; writing nil retains Go's panic |
| `ffi::map_delete(value, key)` | Delete if present; missing keys and nil maps are no-ops |
| `ffi::map_copy(value)` | Independent entry storage with shallow values; preserve nil |

Assignment and argument passing share entry storage. Use `map_copy` for independent entries. Pointer values in a copy still refer to the original objects. Copy traverses key/value pairs directly, so NaN-keyed entries retain their values even though NaN cannot be looked up by equality. None of these operations freezes storage, introduces synchronization or converts HashMap hashing semantics.

```goml
use std::ffi;

#[go_ffi("example.com/myapp/goshim", "Entries")]
extern fn entries() -> ffi::RawMap[i64, i64];

fn available() -> bool {
    !ffi::map_is_nil(entries())
}
```

Foreign signature validation uses Go's key comparability rules. GoML checking rejects known non-comparable keys, including slices, raw maps, arrays or tuples containing them, and local structs with such fields; inferred generic applications are checked after inference. Repeated local generic wrappers are checked at each concrete instantiation: `Key[Key[RawSlice[i64]]]` is rejected when `Key[T]` stores `T` directly, while wrappers that only store a pointer to `T` remain comparable. Recursive struct layouts and structural checks exceeding 64 nested struct instantiations produce recoverable diagnostics. A pointer to an otherwise non-comparable type remains a valid key. Public GoML struct interfaces retain a comparability summary for private fields, so downstream key checking observes their restrictions without exposing field names or types. Summaries may depend on generic arguments; unknown hidden structure does not establish comparability. The artifact decoder also rejects directly non-comparable key shapes. External named keys are checked from Go metadata, including private fields, aliases and imported declarations. Generic external keys retain comparability requirements on their type parameters; concrete arguments are substituted when the type is used as a key, including nested types and aliases that reorder or erase arguments. A parameter used only behind a Go pointer imposes no comparability requirement. Disabling foreign call validation does not disable these checks. Build/link validation also checks keys through materialized local struct fields and external named types after generic specialization. External checks use persisted Go metadata and substitute concrete arguments into comparability rules, including nested types and reordered aliases. Missing or unresolved external key metadata produces a recoverable diagnostic, even with foreign call validation disabled. Module checking infers internal comparability requirements from generic function bodies, including requirements reached through function values and nested closures. Package checks propagate them across files until signatures stabilize, with a recoverable limit of 256 passes. Exported function interfaces retain the inferred requirements, so downstream module checks can reject invalid keys before building Go. Private structural key wrappers are reduced to their required type parameters or associated-type projections. These predicates have no user-written bound syntax and do not introduce let-generalization. Source-only dependency analysis checks dependency bodies and retains their inferred requirements before checking callers, including unsaved source overrides. Struct and enum interfaces also retain the key requirements of maps stored in fields or variant payloads, including private storage and aliases. These requirements propagate between nominal declarations with a recoverable 256-pass limit and are substituted at concrete type applications. Recursive type applications are checked through their declaration requirements without repeatedly expanding their layouts. Default trait methods retain their inferred map-key requirements separately from the method declaration bounds. An implementation that uses a default body must satisfy those requirements after substituting its receiver and associated types. An explicit method override does not inherit requirements from the unused default body. GoLibrary map exports remain unsupported. Raw maps add no synchronization or automatic deep copy.

### Generating Go bindings

Use `goml bind-go <CONFIG> [--compiler <COMPILER>] [--dry-run]` to generate explicitly allowlisted Go bindings from a versioned JSON configuration. `gomlc bind-go <CONFIG> [--dry-run]` is the compiler entry point. The configuration selects each native package/symbol and may list finite `type_arguments`; Go checks their constraints. It generates a GoML raw binding file and native Go forwarding functions without running package initializers or creating dependencies. An existing GoML module and an enclosing Go module are required. Output paths resolve relative to the configuration, remain within the GoML module, and must match the configured Go import path. Parent traversal, symbolic links and nested module boundaries are rejected.

Native APIs may reside in the output Go package; same-package references do not introduce self imports. Metadata queries omit the previously owned output in memory, then Go checks the complete candidate package, including handwritten files, before publication. No source overlay is written to disk.

The ownership manifest `<CONFIG>.goml-bind.json` records both generated files. Identical regeneration preserves timestamps; modifications to either source or the manifest prevent overwriting. Put handwritten conversions and wrappers in separate files. Raw strings and errors retain their explicit `std::ffi` boundaries; the generator does not infer error, nullable or record adapters. Dry runs validate configuration, module and output paths and print destinations without writing files; they do not query the selected symbols. See [the configuration contract](ffi/bind-go.md) and [the complete standard-library example](../examples/ffi-bind-go/README.md). This command uses existing extern/type-alias syntax and introduces no grammar changes.

### Exporting a Go library

The frontend recognizes `#[go_export("Add")] pub fn add(x: i64, y: i64) -> i64 { x + y }`. It checks that the declaration is an ordinary top-level public function without generics, comptime/derive capability, or test attributes. The Go name starts with an ASCII uppercase letter and contains only ASCII letters, digits, or underscores; names must be unique across the selected GoML package. Duplicate attributes and attributes on methods, externs, types, fields, and other non-function items are rejected.

Export parameters currently allow bool, integer and floating-point primitives, `ffi::String`, and fixed arrays of those values. Returns additionally allow unit or a flat tuple of non-unit permitted values. Transparent aliases normalize to their underlying type. Ordinary GoML strings and chars, slices, channels, user-defined types, function values, tuple parameters, and nested tuple returns are outside this export boundary.

The text boundary policy is explicit in the GoML types and adapter body:

| Exported GoML type | Go signature | Boundary policy |
| --- | --- | --- |
| `ffi::String` | `string` | Preserve arbitrary bytes, including invalid UTF-8 and NUL; no implicit Unicode conversion |
| `ffi::Rune` (alias of `i32`) | `int32` / `rune` | Preserve the raw integer; no implicit scalar validation |
| ordinary `string` / `char` | Rejected | Choose raw types and write an explicit checked adapter |

Use `ffi::string_text` or `ffi::rune_char` inside the exported function to validate values before treating them as Unicode. Failure must be expressed by the adapter's supported return signature or a deliberate boundary failure strategy. The following adapter returns the original bytes with `false` for invalid UTF-8, and validated text with `true` otherwise. It never silently substitutes replacement characters:

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

A raw echo returning `ffi::String` can preserve bytes without decoding. A rune adapter can return `(ffi::Rune, bool)` in the same way. Fixed arrays and flat tuple returns preserve these element policies. `Result` itself is not a Go export ABI type; selecting and documenting the failure shape belongs to the adapter author. These exports reuse existing type aliases, function syntax and `go_export` grammar.

Interface and Core artifacts preserve the selected Go export name, stable package/function identity, resolved signature, and relative source origin. Export names and signatures participate in the interface semantic hash; source positions do not. The internal library linker selects only the chosen package’s declared exports as function roots before dead-code elimination and monomorphization; reachable private and cross-package helpers are retained without requiring a `main`. The internal Go library backend emits callable wrappers for scalar and fixed-array parameters and results, maps unit to no Go result, and maps a flat tuple to multiple Go results. Library wrappers use package-level OnceCell initialization and preserve panics. Generated internal package declarations are private Go identifiers; only explicitly selected export names become public API. `pub` alone does not request a Go export.

Use `goml export-go` to generate a library from one package in the current GoML module:

```text
goml export-go [PACKAGE_DIR] --import-path example.com/host/gen/calclib --out ./gen/calclib
```

The default source is the GoML module root. Explicit source and output paths are relative to the current working directory. The source must be a package in the current module and must not declare `package main`. Only its `go_export` declarations become Go API; dependency exports are retained only when reachable. The output directory's basename is the Go package name and must be a non-main ASCII Go identifier.

The GoML module root must already contain `go.mod`. The output must remain inside that Go module, outside nested Go modules, and the supplied import path must equal the Go module path plus the canonical relative output path. Parent traversal is rejected and symlinks are resolved before checking boundaries. The command never creates or repairs `go.mod` or `go.sum` or downloads missing dependencies.

The output contains `goml_generated.go` and `goml_exports.json`. The manifest records public signatures, generation version, build prerequisites, FFI check mode and ownership digests. Handwritten Go files are preserved and checked together with the candidate generated source. New output is checked using a Go overlay before publication; malformed Go, duplicate symbols and import cycles fail without publishing the candidate. Modified or unowned generated files are not overwritten. Concurrent generation for the same output is rejected. Publication rolls back reported file replacement errors; two-file publication is not crash-atomic across power loss or forced process termination.

`--ffi-check required` is the default. All declared foreign bindings, including unused private ones, are checked before publication in the final generated-package context. `--ffi-check off` emits an unverified warning when foreign bindings exist and records `off` in the manifest; Go still compiles the candidate package. `--compiler`, `--target-dir`, `--jobs` and `--dry-run` are supported. The generated pure-Go library builds without the GoML compiler or metadata helper installed. A complete bidirectional example, including a Go executable and Go tests, is available in [the interoperability fixture](../tools/release/testdata/go-export).

## Tests

### Test functions and attributes

Top-level functions are marked as tests using `#[test]`:

```goml
package math;

use std::testing;

#[test]
fn addition_works() -> () {
    testing::assert_eq(1 + 1, 2)
}
```

Test functions must meet the following rules:

- It can only be a top-level function, not an impl method, `extern fn`, structure, enumeration or trait;
- Cannot be named `main`;
- Cannot have parameters or type parameters;
- The return type must be `()`;
- Canonical test IDs generated within the same test package must be unique.

Test functions do not require `pub`. `#[test]` does not accept parameters. When you need to skip a test by default, you can use `#[ignore]` without parameters or `#[ignore("reason")]` with a string reason; `#[ignore]` cannot be used alone without `#[test]`:

```goml
#[test]
#[ignore]
fn unfinished_case() -> () {}

#[test]
#[ignore("requires an external service")]
fn integration_case() -> () {}
```

Parameterized tests use one or more `#[test_case(...)]` attributes on a top-level function. Test cases accept string and boolean arguments, which are checked against the function parameters. Each case receives a content-derived stable ID, while list, filter, text/JSON reporting, artifact manifests, and CodeLens continue to use its readable display name. Ordinary zero-argument `#[test]` functions remain unchanged.

```goml
#[test_case("left", true)]
#[test_case("right", false)]
fn is_left(value: string, expected: bool) -> () {
    testing::assert_eq(value == "left", expected)
}
```

`std::testing` provides the following assertions:

- `testing::fail(message)`: Fail the current test immediately;
- `testing::assert(condition)`: requires the condition to be `true`;
- `testing::assert_eq(actual, expected)`: requires two values of the same type that implement `PartialEq + Debug` to be equal;
- `testing::assert_ne(actual, expected)`: requires two values of the same type that implement `PartialEq + Debug` to be unequal;
- `testing::assert_some`, `assert_none`, `assert_ok`, and `assert_err`: check the variant of an `Option` or `Result`;
- `testing::expect_some`, `expect_ok`, and `expect_err`: return the selected payload or fail with the supplied message.

### White-box and black-box tests

The white-box test file and the source code under test are located in the same package directory. The file name must end with `_test.gom` and declare the same package name:

```text
math/
├── math.gom
└── math_test.gom
```

`math_test.gom`:

```goml
package math;

use std::testing;

#[test]
fn private_helper_works() -> () {
    testing::assert_eq(private_helper(), 42)
}
```

The test build will merge the production source code and all `*_test.gom` in the same directory into a single package, so white-box testing can access the private top-level items of the package. These files will not participate in compilation when executing normal `goml check`, `goml build` or `goml run` for production targets.

Black-box tests are located in the `tests/` directory of the package under test. This directory as a whole constitutes a package named `tests`. The package under test should be imported explicitly and only its public API can be accessed; nested test suites cannot be created under `tests/`:

```text
math/
├── math.gom
└── tests/
    ├── api_test.gom
    └── smoke_test.gom
```

`math/tests/api_test.gom`:

```goml
package tests;

use alice::myapp::math;
use std::testing;

#[test]
fn public_add_works() -> () {
    testing::assert_eq(math::add(20, 22), 42)
}
```

If the identity of the package under test is `alice::myapp::math`, the canonical identity of the test package in the above example is `alice::myapp::math::tests`. Ordinary package discovery will exclude the `tests` directory, and production packages cannot import black-box test packages.

### Check and run tests

`goml check`, `goml build` and `goml test` always process complete modules discovered from the current directory upwards, and do not accept package or file targets. `goml check` only checks the production source code by default; using `--tests` will also check all white-box and black-box tests after the production package check is successful:

```sh
goml check
goml check --tests
goml build
goml fmt
goml fmt --check
```

Test source code cannot be used to fix a production package that itself fails inspection.

`goml fmt` formats the current module's production files, internal tests, and black-box test packages using the same package graphs as builds and tests. It can be run from any nested package directory. Package discovery excludes `testdata`, build output, hidden directories, nested modules, and external dependencies. `goml fmt --check` only checks formatting and exits unsuccessfully when any source file would change.

Run the test using:

```sh
goml test [FILTER]
```

`FILTER` performs substring matching on the complete test display name, such as `goml test addition`. Common options include:

- `--kind internal|external|all`: run only white box, only black box or all tests, the default is `all`;
- `--list`: List matching tests without running them;
- `--ignored`: only run ignored tests;
- `--include-ignored`: Run normal tests and ignored tests at the same time;
- `--jobs N`: Number of parallel workers, default is `1`;
- `--seed N`: Supply a reproducible positive seed to every test;
- `--timeout 500ms|30s|2m`: timeout for a single test, the default is `30s`;
- `--nocapture`: Let the standard output and standard error of the test directly inherit the terminal;
- `--format text|json`: Select text or line-by-line JSON results; JSON format cannot be used with `--nocapture` at the same time;
- `--target-dir`, `--dry-run` and `--compiler`: have the same meaning as project build commands.

Each test is executed in a separate runner process, and failure to exit and timeout do not affect other tests; `--jobs` controls the number of test processes running at the same time. Executing `goml test` requires an available Go toolchain to build the test runner.

`goml test --seed N` supplies one reproducible positive seed to every test process through `testing::seed()`. Failure summaries print the replay seed, and every JSON result event includes it. Captured stdout and stderr remain isolated per test process and are emitted only for failures unless `--nocapture` is selected.

### LSP and editor

The LSP analyzes production files, internal tests, and black-box tests in their corresponding package contexts, so diagnostics, completion, hover, and go-to-definition follow the same visibility rules as compilation. Test code lenses appear on ordinary tests and parameterized test cases. The VS Code extension saves the file and invokes module-level `goml test` with the selected test display name and kind.

Package analysis caches parsed/lowered sources and successful checked packages within a query context. Source contents, package membership, canonical module identity, and transitive dependency fingerprints determine reuse. Editing a package rechecks it and its dependents while independent packages remain reusable; a function-body edit in a dependency conservatively invalidates its dependents too. Open document overrides participate in the same fingerprints. Packages with external Go types and their dependents revalidate native metadata on each analysis. Watched source and manifest changes invalidate cached document analyses; the VS Code extension also watches Go sources, `go.mod`, and `go.sum`.

The custom `goml/expandedDerive` request returns the formatted AST after built-in and programmable derives have run for the requested document. It uses the same package aliases, explicit derive imports, ambiguity checks, dependency interfaces, CTIR verifier, and resource limits as `goml check`. The VS Code command `GoML: Show Expanded Derive` opens that result beside the source file.

`gomllsp` supports full-document formatting through `textDocument/formatting`. Formatting uses the latest unsaved document text and the fixed rules described in [formatting.md](formatting.md). Invalid documents are left unchanged.

## Built-in prelude

The toolchain library has three separate layers. `builtin` is the hidden compiler and runtime contract: it owns primitive and built-in type identities, runtime hooks, language items, built-in implementations, and derive handlers. `prelude` is an independently compiled package that depends on `builtin` and defines the public names automatically placed in ordinary source scope. `std` contains normal standard-library packages and is available only through explicit imports.

User and third-party packages cannot explicitly import `builtin` or `prelude`; attempts produce a diagnostic instead of resolving a registry package. This keeps raw runtime helpers and derive internals hidden while preserving the stable identities of re-exported types and traits. Names such as `Option`, `Result`, `Vec`, `ToString`, `print`, `println`, and `range` remain available without `use` through the prelude.

Compiler-recognized protocols are declared in `builtin` with reserved `#[lang(...)]` metadata. The metadata identifies the owning type or trait together with protocol members such as enum variants, associated types, and methods, so compiler behavior does not depend on their public spelling. `#[lang(...)]` is rejected in user, third-party, prelude, and standard-library sources.

The following names can be used without `use`.

### `Option` and `Result`

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

Construction uses `Option::Some`, `Option::None`, `Result::Ok` and `Result::Err`. Patterns may use either those full names or contextual `Some`, `None`, `Ok` and `Err`.

`Option[T]` provides `is_some`, `is_none`, `unwrap_or`, and `unwrap_or_else`, plus type-changing generic methods:

- `value.map(map_fn) -> Option[U]`
- `value.and_then(next) -> Option[U]`
- `value.ok_or(error) -> Result[T, E]`

`Result[T, E]` provides `is_ok`, `is_err`, `unwrap_or`, and `unwrap_or_else`. Its type-changing generic methods are:

- `value.map(map_fn) -> Result[U, E]`
- `value.map_err(map_fn) -> Result[T, F]`
- `value.and_then(next) -> Result[U, E]`

### `Default`

The prelude `Default` implementations use `()` for the empty tuple, `false` for `bool`, zero for numeric types, and the empty string for `string`. `Vec` and `HashMap` default to empty containers, and the standard collection types `HashSet`, `IndexMap`, `Stack`, `Deque`, `BitSet`, `Arena`, `IndexVec`, and `Interner` likewise default through their empty constructors. `Option[T]` defaults to `None` without requiring `T: Default`, while tuples and fixed arrays default each element. `Result[T, E]` and `Ref[T]` intentionally have no global default implementation.

### Output and string conversion

- `print[T: ToString](value: T) -> ()`
- `println[T: ToString](value: T) -> ()`
- `value.to_string() -> string`, suitable for values that implement `ToString`
- `value.debug() -> string`, suitable for values that implement `Debug`
- `string.len() -> isize` and `string.byte_len() -> isize`, both returning the UTF-8 byte length
- `string.get(index: isize) -> char`, decoding a character at a UTF-8 byte boundary
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
- `string.starts_with_at(start: isize, prefix: string) -> bool`, checking at a byte offset
- `string.find(expected: string) -> Option[isize]`, `string.rfind(expected: string) -> Option[isize]`, and `string.find_bytes(expected: string) -> Option[isize]`, returning byte offsets
- `string.char_count() -> isize`, counting Unicode scalar values
- `string.slice_chars(start: isize, end: isize) -> Option[string]`, using scalar-value indexes
- `string.trim() -> string`, `string.trim_start() -> string`, and `string.trim_end() -> string`, trimming ASCII whitespace
- `string.split(separator: string) -> Vec[string]`, `string.split_once(separator: string) -> Option[(string, string)]`, and `string.lines() -> Vec[string]`
- `string.replace(expected: string, replacement: string) -> string`
- `string.repeat(count: isize) -> string`
- `string.is_ascii() -> bool`
- `string.eq_ignore_ascii_case(other: string) -> bool`
- `string.to_ascii_lowercase() -> string` and `string.to_ascii_uppercase() -> string`

The basic scalar types all implement `ToString` and `Debug`. String concatenation uses `+`.

The key-related traits are the method-bearing `PartialEq::eq(self, other: Self) -> bool`, the marker `Eq: PartialEq`, and `Hash::hash(self) -> u64`. User types can use handwritten impls or derive the corresponding traits.

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
- `ptr_eq(a, b) -> bool`, compare reference identities

The built-in `PartialEq`, `Eq`, and `Hash` implementations for `Ref[T]` use reference identity and do not require `T` to implement those traits. Mutating the referenced value therefore does not change equality or hashing. `Ref[T]` does not implement `Default`, because implicit allocation and recursive default construction would be surprising.

### Fixed arrays

```goml
let mut values: [i32; 3] = [1, 2, 3];
let first = values[0];
values[1] = 20;
```

The index type is `isize`. The underlying `array_get` and `array_set` can also be called; index syntax is preferred for daily code.

### `Vec[T]`

```goml
let values = Vec::from_array([1, 2, 3]);
let first = values.get(0);
```

`Vec::from_array([values...])` creates a vector from a fixed array. Array items are evaluated once from left to right. The vector receives new outer backing storage, so replacing an array element later does not change the vector, while referenced elements remain shared because the copy is shallow. Use `Vec::new()` for an empty vector. Direct nonempty array calls coallocate the vector header and backing array. Passing an array variable retains the same shallow-copy semantics. The associated function is an ordinary first-class value, for example `let make: ([i32; 3]) -> Vec[i32] = Vec::from_array;`.

Commonly used methods:

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

`vec[index]` is equivalent to reading the element, `vec[index] = value;` modifies the element.

`copy()` shallow-copies the outer backing storage. `freeze()` also copies the backing storage and returns a `FrozenVec[T]`, so later structural or element changes through aliases to the original `Vec` cannot affect the frozen value.

`contains(value)` requires the element type to implement `PartialEq` and returns whether any element equals `value`.

Sorting and selection methods take an integer comparator returning negative/zero/positive, or an `Ordering` comparator for the `_ordering` variants. `dedup()` requires `PartialEq`; `position` takes a predicate; `binary_search_*` methods expect the vector to already be ordered and return the first matching index. `join(separator)` requires `ToString` and concatenates the element string forms.

### `FrozenVec[T]`

`FrozenVec[T]` is a dynamically sized read-only vector. It exposes `get`, `len`, `is_empty`, `contains`, `iter`, and `to_vec`, but no mutation methods. `to_vec()` returns a new shallow-copied mutable backing store. For nested containers, freezing is shallow: use `FrozenVec[FrozenVec[T]]` when both levels must be read-only.

### `Slice[T]` and `MutSlice[T]`

`Slice[T]` is a bounded read-only view on the `Vec[T]` storage:

```goml
let view: Slice[i32] = values.slice(1, 3);
let item = view.get(0);
```

Commonly used read-only methods are `get`, `get_checked`, `len`, `contains`, `sub`, `sub_checked`, `to_vec`, and `iter`. `contains(value)` requires the element type to implement `PartialEq`. `to_vec()` creates a shallow copy in a new `Vec[T]`; `view[index]` can be read, but `view[index] = value;` is rejected.

`MutSlice[T]` is a bounded mutable zero-copy view. `Vec::slice_mut(start, end)` and `Vec::as_mut_slice()` create it; `as_slice()` converts it to a read-only view without copying:

```goml
let values = Vec::from_array([1, 2, 3, 4]);
let output = values.slice_mut(1, 4);
output[0] = 20;
let written = output.copy_from(1, Vec::from_array([7, 8]).as_slice());
```

`get_checked`, `set_checked`, `sub_checked`, `Vec::slice_checked`, and `Vec::slice_mut_checked` report invalid ranges without indexing. `copy_from(offset, source)` validates the complete destination range before writing and uses overlap-safe copy semantics; `copy_within` is the corresponding same-view operation. `fill`, `to_vec`, and iteration are also available. `Slice` and `MutSlice` keep their current backing array alive. A later `Vec` growth may move the vector to a new backing array, so an older view continues to refer to the storage captured when the view was created.

### `HashMap[K, V]`

The key type must implement both `Hash` and `Eq`; bucket equality dispatches through the `PartialEq` supertrait:

```goml
let counts = HashMap::from_array([("a", 1), ("b", 2)]);
let value: Option[i32] = counts.get("a");
```

`HashMap::from_array([(key, value), ...])` creates a map from an entries array. It first evaluates the complete array from left to right, then inserts its pairs in array order. Later duplicate keys overwrite earlier entries. Use `HashMap::new()` for an empty map. Direct array-literal calls preallocate their build storage from the entry count. The associated function is an ordinary first-class value when its array length is fixed by a function type.

Commonly used methods: `new`, `from_array`, `get`, `set`, `insert`, `get_or_insert_with`, `update`, `remove`, `remove_value`, `entry`, `len`, `contains`, and `entries`. `insert` returns the previous value, `update` applies a `(V) -> V` closure only when the key exists, and `remove_value` returns the removed value. `entries()` returns a snapshot `Vec[(K, V)]`.

`entry(key)` returns a `HashMapEntry[K, V]` with `get`, `or_insert`, `or_insert_with`, `insert`, `and_modify`, `update`, `remove`, and `remove_entry`. The fused entry operations hash and search the key once; the creation closure passed to `or_insert_with` runs only for a missing key:

```goml
let counts: HashMap[string, isize] = HashMap::new();
let count = counts.entry("jobs").or_insert_with(|| 0);
let _ = counts.entry("jobs").update(|value| value + 1);
let removed = counts.entry("jobs").remove_entry();
```

Index reading returns `Option[V]`, and index assignment writes `V`:

```goml
let value = counts["a"];
counts["b"] = 2;
```

### Channels

`Channel[T]` is a Go channel backend that supports buffered and unbuffered communication:

```goml
let channel = Channel::[string]::new(0);
go || channel.send("ready");
let value: Option[string] = channel.recv();
channel.close();
```

Commonly used methods:

- `Channel::[T]::new(capacity: isize) -> Channel[T]`
- `send(value: T) -> ()`
- `recv() -> Option[T]`, returns `Option::None` when closed and drained
- `close() -> ()`
- `sender() -> Sender[T]`
- `receiver() -> Receiver[T]`
- `split() -> (Sender[T], Receiver[T])`

`Sender[T]` provides `send` and `close`; `Receiver[T]` provides `recv`. Direction is enforced statically and emitted as Go's `chan<- T` and `<-chan T`. Converting a bidirectional channel requires an explicit endpoint method. `Channel[T]` keeps its original send, receive, and close API.

Capacity `0` creates an unbuffered channel. Sending to an unbuffered channel should generally be concurrently received by another `go` closure, and vice versa.

Use `select` when a goroutine must wait on several channels, choose between sending and receiving, observe closure through `Option[T]`, or perform a non-blocking operation with `default`. The channel operands and send values are prepared once before selection.

`std::channel` provides runtime-sized homogeneous selection. `Operation[T]` is either `Receive(Receiver[T])` or `Send(Sender[T], T)`. `select` blocks, `try_select` returns `Ok(None)` when nothing is ready, and `try_select_priority` probes from the lowest input index. Successful results are `Selection::Received { index, value }` or `Selection::Sent { index }`; a closed receive carries `None`. An empty operation slice returns `SelectError::Empty`.

### Iterator

The built-in protocols are:

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

The protocol entry points are `FnIterator::from_fn(next_fn)`, `iterator.next()` or `Iterator::next(iterator)`. `range(start, end)` and `start..end` produce incrementing half-open `FnIterator[isize]` values; `start >= end` is empty.

New code should import `std::iter` for iterator construction, adapters, and consumers:

```goml
use std::iter;

let values = iter::map(Vec::from_array([1, 2, 3]).iter(), |value: isize| value * 2);
let total = iter::fold(values, 0, |sum: isize, value: isize| sum + value);
```

- producers: `empty`, `once`, and `from_fn`
- adapters: `map`, `filter`, `filter_map`, `take`, `take_while`, `skip`, `skip_while`, `enumerate`, `zip`, `chain`, `inspect`, and `map_while`
- consumers: `fold`, `collect`, `find`, `find_map`, `any`, `all`, `count`, `position`, `nth`, `last`, `for_each`, and `reduce`

Iterators are single pass. Fixed arrays use native indexed `for` lowering. `Vec[T]`, `Slice[T]` and `MutSlice[T]` can be directly used in `for`, and a value implementing `Iterator` is also directly iterable through the identity `IntoIterator`.

## Standard library package

The standard library is distinct from both the hidden builtin contract and the automatically scoped prelude. Its packages must be imported explicitly:

```goml
use std::ascii;
use std::bincode;
use std::bytes;
use std::bytes::endian;
use std::channel;
use std::cmp;
use std::collections;
use std::crypto;
use std::encoding::base64;
use std::encoding::hex;
use std::error;
use std::env;
use std::ffi;
use std::fs;
use std::fs::notify;
use std::fs::walkdir;
use std::io;
use std::iter;
use std::json;
use std::math;
use std::num;
use std::os::linux::syscall;
use std::path;
use std::process;
use std::rand;
use std::serde;
use std::task;
use std::testing;
use std::text;
use std::toml;
use std::time;
use std::utf8;
use std::utf16;
use std::unicode;
```

Public APIs include:

- `ascii::is_ascii`, character-class predicates, ASCII case conversion and comparison, and `escape_default`
- `bincode::standard`, `legacy`, configuration builders, serde `Serialize` and `Deserialize` re-exports, `encode_to_vec`, and `decode_from_slice`
- `bytes::Bytes`, `bytes::Builder`, checked and zero-copy byte views, plus `bytes::endian::{Builder, Reader, Writer, Endian}` and checked integer and floating-point reads and writes
- `channel::Operation`, `Selection`, `select`, `try_select`, and `try_select_priority` for runtime-sized channel selection
- `cmp::Ordering`, `Ord`, `Reverse`, comparison helpers, and two-value minimum, maximum, and clamping operations. `Ordering` is a builtin type re-exported by `cmp`.
- `collections::Arena`, `BinaryHeap`, `BitSet`, `BTreeMap`, `BTreeSet`, `Deque`, `HashSet`, `IndexMap`, `IndexSet`, `IndexVec`, `Interner`, and `Stack`; hash-backed collections require `Hash + Eq`, while tree collections and heaps use `cmp::Ord`
- `collections::sort`, `stable_sort`, `binary_search`, `min`, `max`, and their comparator variants. The sorting, search, selection, and deduplication methods on `Vec[T]` are the canonical forms.
- `crypto::hash` one-shot SHA-256 and `crypto::rand` operating-system random bytes
- `encoding::hex` lowercase and uppercase hexadecimal encoding plus checked decoding
- `encoding::base64` RFC 4648 standard and URL-safe encoding with padded and unpadded variants
- `error::Error`, `ErrorKind`, `Details`, and stable error-kind code conversion
- `env::args`, current-directory and executable queries, and environment-variable reads
- `ffi::String`, `Rune`, `Ptr`, `Error`, `Func`, `RawSlice`, `RawMap`, and explicit Go boundary adapters
- `fs::read_file_structured`, `write_file_structured`, structured byte I/O, directory operations, path inspection, and `sha256_file`
- `fs::notify` Linux amd64 file and recursive-directory notifications, event filters, ignore rules with subtree pruning, independent multi-path registrations, cancellable reads, bounded subscriptions, rename cookies, and rescan signals
- `fs::walkdir::{walk, WalkDir, WalkIterator, DirEntry, Error}` for lazy directory traversal, depth bounds, pruning, link following, and contextual errors
- `io::print`, `println`, `eprint`, `eprintln`, and byte-oriented standard stream I/O
- `iter::empty`, `once`, `from_fn`, iterator adapters, and single-pass consumers
- `json::Value`, `parse`, `encode`, serde `Serialize` and `Deserialize` re-exports, `to_value`, `from_value`, `try_to_string`, `from_string`, `field`, and typed `as_*` accessors
- `math` f32/f64 elementary functions, IEEE 754 classification, and the `E`, `PI`, `TAU`, `SQRT_2`, `LN_2`, and `LN_10` constants
- `net::{IpAddr, SocketAddr, TcpListener, TcpStream, UdpSocket, WaitOptions}` for Linux amd64 syscall-backed IPv4/IPv6 networking, shared epoll readiness, timeouts, and cancellation
- `num` structured parsing plus checked and saturating `i64` arithmetic
- `os::linux::syscall` Linux amd64 calls by number, six machine-word arguments, scoped mutable byte buffers, raw return values, and numeric errno
- `path::join`, `clean`, `is_absolute`, component inspection, and `absolute_structured`
- `process::Command`, structured whole-process execution, `ExitStatus`, `Output`, `exit`, and `look_path_structured`
- `rand::ALGORITHM`, `next_u64`, deterministic byte generation, integer ranges, and shuffle with an explicit seed
- `serde::Value`, `Serializer`, `Deserializer`, `Serialize`, `Deserialize`, `value_serializer`, `value_deserializer`, `to_value`, and `from_value`
- `task::Scope`, `Task[T]`, `CancelToken`, `WaitResult[T]`, `scope`, and `try_scope`
- `testing::fail`, boolean/equality assertions, and `Option`/`Result` shape assertions
- `text::StringBuilder`, `LineIndex`, `LineColumn`, and `PositionEncoding`; the byte-offset search, character iteration and slicing, trimming, splitting, replacement, joining, repetition, and explicit ASCII operations are string methods
- `toml::Value`, `parse`, `encode`, serde `Serialize` and `Deserialize` re-exports, `to_value`, `from_value`, `to_string`, and `from_string`
- `time::Duration`, `Instant`, `SystemTime`, `sleep`, and `sleep_with`
- `utf8::validate`, `decode`, `decode_slice`, `encode`, `encode_into`, `encode_to`, `encoded_len`, and `Utf8Error`
- `utf16::decode`, `decode_bytes`, `encode`, `encode_bytes`, and `Utf16Error`
- `unicode::VERSION`, scalar properties, and Unicode case conversion using Unicode 15.0.0 tables

### Structured error foundation

`std::error` provides the common protocol used by structured standard-library errors. `Error` is a marker trait requiring `Debug` and `ToString`. `ErrorKind` defines stable, message-independent categories including missing files, permission failures, invalid input or data, timeouts, interruption, short I/O, broken pipes, unsupported operations, and `Other`.

`Error` has a blanket implementation for every value implementing `Debug` and `ToString`, so domain errors participate in the common protocol directly.

`kind_code` and `kind_from_code` convert between `ErrorKind` and the stable integer representation used by runtime and artifact boundaries. Unknown integer values map to `Other`. `Details` stores the stable kind code, operation, optional context, optional raw operating-system code, and a display-only message. Programs may branch on the kind code or `ErrorKind`, but must not parse the host message.

```goml
use std::error;

fn describe[E: error::Error](value: E) -> string {
    value.to_string()
}

fn example() -> string {
    let value = error::Details::new(
        error::kind_code(error::ErrorKind::NotFound),
        "read",
        Option::Some("missing.txt"),
        Option::None,
        "file does not exist",
    );
    describe(value)
}
```

`io`, `fs`, `path`, `env`, `process`, `num`, and `time` expose domain-specific error types. Callers can branch on stable error kinds and use `to_string()` only when a display message is needed.

### UTF-8 validation and conversion

`std::utf8` validates byte slices and converts complete byte vectors to strings without admitting invalid UTF-8. `Utf8Error::valid_up_to` is the length of the valid prefix. `error_length` is the length of the invalid sequence when known and is `None` for an incomplete sequence at the end of the input. `encode` returns the UTF-8 bytes of a string, and `encoded_len` returns the byte length of one Unicode scalar value.

Importing `utf8::BytesUtf8` adds `Bytes::to_string_utf8`, which returns `Result[string, Utf8Error]`.

`decode_slice` accepts a read-only byte view. `encode_into` writes into a `MutSlice[byte]` at a checked offset and reports `bytes::BoundsError` without a partial write; `encode_to` appends to the lightweight `bytes::Builder`.

### Byte buffers and endian access

`std::bytes` uses `Slice[byte]` and `MutSlice[byte]` for borrowed views. `Bytes::as_slice` and `as_mut_slice` are zero-copy, while `slice_checked` and `slice_mut_checked` validate a subrange. `bytes::Builder` is a lightweight byte accumulator and returns `Bytes`.

`std::bytes::endian` is an opt-in package for binary formats. Its `Builder` grows while writing typed values; `Reader` advances over a read-only view; `Writer` advances over a fixed mutable view and returns `endian::BoundsError` rather than partially writing past the end. The top-level `read_u16/u32/u64`, `read_i16/i32/i64`, `read_f32/f64` and matching `write_*` functions take an explicit `Endian::Little` or `Endian::Big`. One-byte operations omit endianness. Every operation validates the complete range before reading or writing.

The stateful `Reader`, `Writer`, and `Builder` currently provide typed methods for `u8`, `u16`, `u32`, and `u64`. Signed and floating-point access uses the top-level functions with an explicit offset; those functions do not advance a reader or writer.

### UTF-16 conversion

`std::utf16` converts between strings and `Slice[u16]`, or between strings and endian-tagged byte slices. Decoding rejects lone surrogates and odd byte lengths with an indexed `Utf16Error`; encoding emits surrogate pairs for non-BMP scalar values.

### Text and source positions

Text search indices are UTF-8 byte offsets, matching the indices accepted by the built-in string APIs. `starts_with_at` returns false for out-of-range and non-character-boundary offsets. `rfind` returns the last matching byte offset. Trimming recognizes ASCII whitespace. Splitting on an empty separator returns the original string as one item, while `split_once` with an empty separator returns `Option::None`.

`text::LineIndex` precomputes line starts and non-ASCII scalar positions for repeated source-position conversion. `offset_to_line_column` and `line_column_to_offset` support UTF-8, UTF-16, and UTF-32 columns, treat CRLF as one line ending, and clamp unchecked out-of-range positions. `line_column_to_offset_checked` rejects positions inside an encoded scalar or outside the source.

`text::find`, `text::rfind`, and `text::find_bytes` return byte offsets; the string methods `find`, `rfind`, and `find_bytes` are the canonical forms. `text::char_indices` yields byte offsets paired with Unicode scalar values, `text::char_count` counts scalar values, and `text::slice_chars` uses scalar-value indexes and returns `None` for an invalid range.

### ASCII and Unicode

`std::ascii` operates on `byte`. Classification and case conversion use only the 7-bit ASCII range, and bytes above `0x7f` remain unchanged. `escape_default` emits short escapes for tabs, carriage returns, newlines, quotes, and backslashes, preserves printable ASCII, and uses lowercase `\\xNN` escapes for other bytes.

`std::unicode` fixes its public data version to Unicode 15.0.0. Character predicates operate on one Unicode scalar value. `lowercase` and `uppercase` apply Unicode case mapping to a complete string. `case_fold` uses the checked-in table generated by `tools/generate_unicode_casefold.py`, supports multi-scalar folds, and is locale independent.

### Mathematics

`std::math` delegates elementary operations to Go's `math` package and follows its IEEE 754 special-value behavior. The f32 forms calculate through f64 and round the result back to f32. Results therefore use the target Go toolchain's correctly rounded conversions but do not promise bit-for-bit equality across different operating systems or processor implementations for every transcendental function.

### Randomness and cryptographic helpers

`std::crypto::hash::sha256` returns the lowercase hexadecimal SHA-256 digest of a byte buffer, and `sha256_file` hashes a complete file before returning. `std::crypto::rand::bytes` reads the requested number of bytes from the operating-system cryptographic random source. These APIs do not expose hasher or random-source handles.

`std::rand` uses the versioned `splitmix64-v1` algorithm. Every operation requires an explicit seed, and identical inputs produce identical outputs. It is intended for tests, simulations, sampling, and shuffling and is not cryptographically secure.

### Hexadecimal and base64

`std::encoding::hex` encodes `bytes::Bytes` to lowercase hexadecimal by default, while `encode_upper` emits uppercase digits. Decoding accepts either case and returns `DecodeError` with the byte offset of an odd length or invalid digit.

`std::encoding::base64` uses the padded RFC 4648 standard alphabet by default. `Variant` selects standard or URL-safe alphabets with required or omitted padding. Decoding is strict: it rejects invalid lengths, alphabet mixing, misplaced padding, and nonzero unused trailing bits.

### Time

`Duration` stores a non-negative number of nanoseconds and offers constructors and whole-unit accessors for nanoseconds, microseconds, milliseconds, and seconds. Subtraction saturates at zero. `Instant` is monotonic and is suitable for elapsed-time measurement. `SystemTime` exposes Unix nanosecond, millisecond, and second timestamps. `time::sleep` blocks the current goroutine for a `Duration`.

`Duration` provides checked and saturating scaled constructors, addition, subtraction, and multiplication. Checked operations return `None` on overflow, underflow, or a negative input. Saturating operations clamp to zero or the largest signed 64-bit nanosecond value. `Duration`, `Instant`, and `SystemTime` expose `compare`; `Instant::checked_duration_since` returns `None` when the receiver precedes the supplied instant.

### Files and standard streams

`fs::read_file_structured`, `read_bytes_structured`, `write_file_structured`, and `write_bytes_structured` perform whole-file I/O. They return `fs::Error` with a stable `io::ErrorKind`, operation, path, optional raw operating-system code, and display message. Directory creation and removal, canonicalization, directory listing, and file hashing use the same error type.

`fs::create_dir`, `rename`, `copy`, `hard_link`, and `symbolic_link` are eager operations returning `fs::Error`. `copy` reads and writes the complete file and currently creates the destination with portable `0644` permissions; it does not preserve source metadata.

`fs::metadata` and `symlink_metadata` return value-only metadata including file type, length, portable permission bits, and modification time. `Metadata::from_parts(file_type, length, mode, modified_unix_nanoseconds)` constructs the same value without filesystem access; permission bits are masked to `0o777`. `read_dir_structured` eagerly snapshots directory entries. `read_dir_names_structured` returns sorted names without fetching each child's metadata, preserving structured listing errors; a disappearing child does not invalidate the other names. `atomic_write` writes, synchronizes, closes, and atomically renames a same-directory temporary file before returning; temporary cleanup stays inside the runtime call. `replace` exposes the host atomic rename operation under replacement semantics.

`io::read_stdin_structured`, `read_stdin_exact_structured`, `write_stdout_structured`, and `write_stderr_structured` provide structured errors for standard streams. `read_stdin_to_string` validates the complete input as UTF-8 and reports `InvalidData` on failure. A negative exact-read length reports `InvalidInput` before accessing stdin.

### Numeric parsing and checked arithmetic

Numeric parsing returns `Result[_, num::ParseIntError]` or `Result[_, num::ParseFloatError]`. Integer radix parsing accepts radix `0` or `2..36`; radix `0` recognizes `0b`, `0o`, and `0x` prefixes and permits Go-style digit separators. Invalid radices, malformed input, and overflow return `Result::Err`. Floating-point parsing supports decimal and hexadecimal IEEE 754 input, signed exponents, digit separators, `inf`, `infinity`, and `NaN`, and rounds directly to the requested `f32` or `f64` width.

`num::parse_int_structured`, radix and unsigned variants, and the structured float parsers return domain parse errors. The `checked_*_int64` operations return `None` on overflow; the corresponding `saturating_*_int64` operations clamp to the signed 64-bit bounds.

### Environment, paths, and processes

`env::current_dir_structured`, `current_exe_structured`, and `var_structured`, `path::absolute_structured`, and the `process` structured execution methods expose whole-operation errors. Process timeout methods take `time::Duration`, terminate and wait through the command runtime, and return `TimedOut` through `process::Error`.

Paths remain UTF-8 `string` values. `path::separator` reports the host separator, `components` recognizes both slash forms, `relative` uses host path rules, and `windows_prefix` recognizes drive and UNC prefixes independently of the host operating system. Non-UTF-8 operating-system names cannot be represented and therefore cannot appear in these APIs.

### Linux amd64 system calls

`std::os::linux::syscall` is an explicit low-level escape hatch supported only on Linux amd64. `syscall6(number: usize, args: [usize; 6]) -> Result[SyscallResult, Error]` accepts any syscall number and six machine words. Unused argument positions should contain zero. Signed arguments use their machine-word bit pattern, for example `(-100).to_usize()` for Linux `AT_FDCWD`.

`SyscallResult` has public `r1`, `r2`, and `errno` fields, all `usize`, preserving the values reported by Go's `syscall.Syscall6`. `is_ok()` tests `errno == 0`. A kernel failure still returns `Ok(SyscallResult)` with a nonzero errno; callers must inspect it. The outer `Err(Error::UnsupportedTarget)` reports that the executing target is not Linux amd64, before any syscall is issued. Other Go targets are not guaranteed to compile. Syscall numbers, layouts, and constants are specific to the Linux amd64 ABI; this package does not select numbers for another architecture.

```goml
use std::os::linux::syscall;

fn process_id() -> Result[usize, string] {
    let result = syscall::syscall6(syscall::SYS_GETPID, [0, 0, 0, 0, 0, 0]).map_err(
        |error| error.to_string(),
    )?;
    if result.is_ok() {
        Result::Ok(result.r1)
    } else {
        Result::Err("getpid errno " + result.errno.to_string())
    }
}
```

`syscall6_with_buffers(number: usize, args: [Arg; 6])` returns the same result type. `Arg::Word(usize)` passes an integer unchanged. `Arg::Buffer(MutSlice[byte])` passes the address of the view's first byte, or a null pointer for an empty view. The runtime retains and pins each nonempty backing allocation until the call returns; writes made by the kernel are visible through that view. Multiple buffers, subviews, and aliases of the same allocation are supported. A buffer argument does not implicitly add a length argument or a trailing NUL byte.

```goml
use std::os::linux::syscall;
use std::os::linux::syscall::{Arg};

fn write_once(fd: usize, data: Vec[byte]) -> Result[syscall::SyscallResult, syscall::Error] {
    syscall::syscall6_with_buffers(syscall::SYS_WRITE, [
        Arg::Word(fd),
        Arg::Buffer(data.as_mut_slice()),
        Arg::Word(data.len().to_usize()),
        Arg::Word(0),
        Arg::Word(0),
        Arg::Word(0),
    ])
}
```

The caller is responsible for the syscall ABI, valid addresses, buffer lengths, alignment, native structure encoding, resource cleanup, and synchronization. Byte buffers can carry explicitly encoded native records; ordinary GoML structs have no kernel-layout guarantee. `Word` does not retain or pin any Go allocation, so a Go-managed address must not be smuggled through an integer. Numeric addresses returned by operations such as `mmap` remain raw words and require appropriate explicit cleanup such as `munmap`.

Buffer lifetimes cover synchronous calls only. Operations that retain a pointer after returning, nested pointer graphs, and arbitrary Go object memory are outside this buffer API. The runtime uses scheduling-aware `Syscall6`, not `RawSyscall6`. It performs one call without automatic `EINTR` retry or completion of partial reads/writes. Per-thread operations require thread-affinity handling that this API does not provide. Raw changes to threads, process creation, signal handlers, or the Go runtime's address space can violate runtime invariants; accepting a number does not make every kernel operation safe to use from GoML. Calls remain subject to kernel availability, process permissions, and sandbox policy.

The exported syscall constants are `SYS_READ`, `SYS_WRITE`, `SYS_CLOSE`, `SYS_FSTAT`, `SYS_POLL`, `SYS_MMAP`, `SYS_MUNMAP`, `SYS_GETPID`, `SYS_SOCKET`, `SYS_CONNECT`, `SYS_SENDTO`, `SYS_RECVFROM`, `SYS_SHUTDOWN`, `SYS_BIND`, `SYS_LISTEN`, `SYS_GETSOCKNAME`, `SYS_GETPEERNAME`, `SYS_SETSOCKOPT`, `SYS_GETSOCKOPT`, `SYS_GETUID`, `SYS_GETPPID`, `SYS_GETDENTS64`, `SYS_CLOCK_GETTIME`, `SYS_EPOLL_WAIT`, `SYS_EPOLL_CTL`, `SYS_INOTIFY_ADD_WATCH`, `SYS_INOTIFY_RM_WATCH`, `SYS_OPENAT`, `SYS_NEWFSTATAT`, `SYS_ACCEPT4`, `SYS_EPOLL_CREATE1`, `SYS_PIPE2`, and `SYS_INOTIFY_INIT1`, plus `EINTR`, `EBADF`, `EAGAIN`, `EINVAL`, and `ENOSYS`. Other numbers can be passed directly. This package uses ordinary imports, functions, enums, arrays, and mutable slices; it introduces no new grammar or `unsafe` syntax and cannot be used in `comptime`.

### Networking

`std::net` implements IPv4/IPv6 TCP and UDP in GoML using Linux amd64 socket syscalls. It does not delegate networking to Go's `net` package. All sockets are nonblocking and close-on-exec; a lazy shared epoll worker wakes waiting GoML tasks through channels. One epoll descriptor and one worker remain for the lifetime of the process, independent of the number of sockets. The implementation uses one-shot, level-triggered readiness and registration identities that prevent queued events from targeting a reused descriptor. Native layouts and readiness handling follow [socket](https://man7.org/linux/man-pages/man2/socket.2.html), [connect](https://man7.org/linux/man-pages/man2/connect.2.html), [accept](https://man7.org/linux/man-pages/man2/accept.2.html), and [epoll_ctl](https://man7.org/linux/man-pages/man2/epoll_ctl.2.html).

`IpAddr` is `V4([byte; 4])` or `V6([u16; 8])`. `IpAddr::parse(string)` accepts numeric addresses, including compressed IPv6 and an IPv4 tail in IPv6. IPv4 octets use decimal without leading zeroes. `SocketAddr` exposes `ip: IpAddr`, `port: u16`, and `scope_id: u32`; `new(ip, port)` sets scope zero. `SocketAddr::parse` accepts `127.0.0.1:8080`, `[::1]:8080`, or `[fe80::1%3]:8080`. IPv6 scope IDs are numeric; interface names and hostnames are unsupported. Scope must be zero for IPv4. Both address types support equality, debug output, and `to_string()`; IPv6 output uses lowercase hexadecimal with the longest first zero run compressed. Port zero asks the kernel to assign a port when binding; read it with `local_addr()`.

All operations return `Result[T, net::Error]`; `Error` and `ErrorKind` re-export `std::io` types. Errors retain operation names and numeric errno when available. Cancellation is `Interrupted`, timeout is `TimedOut`, use after close is `InvalidInput`, and premature TCP EOF in `read_exact` is `UnexpectedEof`. Connection errors not represented by `io::ErrorKind`, such as connection refused or reset, use `Other` with `raw_os_code()`. Address parsing failures are `InvalidInput`.

| Type and operation | Result value and behavior |
| --- | --- |
| `TcpListener::bind(address)` | `TcpListener`; binds with address reuse enabled and backlog 128 |
| `TcpListener::bind_with_backlog(address, backlog: isize)` | `TcpListener`; backlog must fit a positive signed 32-bit value |
| `TcpListener.accept()` | `(TcpStream, SocketAddr)`; waits for one incoming connection and returns its peer address |
| `TcpStream::connect(address)` | `TcpStream`; waits for connection establishment and checks the socket error after readiness |
| `TcpStream.read(buffer: MutSlice[byte])` | `isize` byte count; may read fewer bytes than requested; zero indicates EOF for a nonempty buffer |
| `TcpStream.write(buffer: Slice[byte])` | `isize` byte count; may write fewer bytes than requested |
| `TcpStream.read_exact(buffer)` / `write_all(buffer)` | `()`; loops until complete, EOF, or error |
| `TcpStream.shutdown(Shutdown::Read / Write / Both)` | `()`; shuts down a direction without releasing the descriptor |
| `TcpStream.set_nodelay(bool)` | `()`; enables or disables TCP_NODELAY |
| `TcpStream.peer_addr()` | `SocketAddr` |
| `UdpSocket::bind(address)` | `UdpSocket` |
| `UdpSocket.send_to(data: Slice[byte], destination: SocketAddr)` | `isize`; sends one datagram, including an empty datagram |
| `UdpSocket.recv_from(buffer: MutSlice[byte])` | `Datagram` with public `source`, copied `len`, full `datagram_len`, and `truncated` fields |
| All three socket types: `local_addr()`, `close()`, `is_closed()` | `SocketAddr`, `()`, and a plain `bool`, respectively |

UDP receives consume exactly one datagram, including when the buffer is empty. Excess bytes are discarded; `datagram_len` and `truncated` report this according to Linux [MSG_TRUNC semantics](https://man7.org/linux/man-pages/man2/recv.2.html). A zero-length UDP receive is a valid datagram, not EOF. TCP zero-length reads and writes complete without data transfer on an open socket. Sending uses `MSG_NOSIGNAL` so a broken connection becomes an error. Send methods copy the immutable input into a syscall buffer; callers must not concurrently mutate input or receive buffers.

`WaitOptions::new()` waits indefinitely. `.with_timeout(time::Duration)` and `.with_cancel(task::CancelToken)` return modified options. `accept_with(options)`, `connect_with(address, options)`, `read_with(buffer, options)`, `read_exact_with(buffer, options)`, `write_with(buffer, options)`, `write_all_with(buffer, options)`, `send_to_with(data, destination, options)`, and `recv_from_with(buffer, options)` apply these settings. A timeout covers the whole operation, including waiting behind another reader or writer and all partial transfers; it does not restart after progress. Zero timeout returns `TimedOut` without attempting I/O. Completed I/O may win a race with cancellation or timeout. After a failed `read_exact` or `write_all`, some bytes may already have transferred; the buffer is not rolled back and the error does not report a partial count. Use individual `read`/`write` calls when tracking progress is required.

Socket values are shared handles. Multiple reads or multiple writes serialize per socket; one reader and one writer can progress concurrently. Concurrent `close()` wakes active and queued operations, releases the socket once, and is idempotent. Failed binds and connects release their descriptors. Always close sockets explicitly or with `defer`; garbage collection does not close them. IPv6 sockets are IPv6-only, so bind separate IPv4 and IPv6 listeners when both are needed. DNS resolution, Unix-domain sockets, TLS, HTTP, and other platforms are outside this initial API. It uses existing imports, enums, methods, channels, and tasks and adds no grammar or compile-time networking.

```goml
use std::net;
use std::time;

fn echo_once() -> Result[(), net::Error] {
    let address = net::SocketAddr::parse("127.0.0.1:8080")?;
    let listener = net::TcpListener::bind(address)?;
    defer { let _ = listener.close(); };
    let wait = net::WaitOptions::new().with_timeout(time::Duration::from_seconds(10));
    let (stream, _) = listener.accept_with(wait)?;
    defer { let _ = stream.close(); };
    let buffer = Vec::from_array([0, 0, 0, 0]);
    let count = stream.read_with(buffer.as_mut_slice(), wait)?;
    stream.write_all_with(buffer.slice(0, count), wait)
}
```

### Directory traversal

`std::fs::walkdir` implements lazy, iterative depth-first traversal in GoML using Linux amd64 syscalls: `openat`, `getdents64`, `newfstatat`, `fstat`, and `close`. Other targets are unsupported; it does not fall back to host directory-listing or metadata APIs. `walk(root: string) -> WalkIterator` uses the defaults from `WalkDir::new(root)`: include the root at depth zero, visit directories before their contents, sort siblings by file name, impose no depth limit, and do not follow symbolic links, including a root link. A file root yields one entry. Construction performs no filesystem I/O; the first iterator step checks the root.

`WalkDir` is a reusable value builder. Each `iter()` or `IntoIterator::into_iter()` creates an independent `WalkIterator`, so it can also be used directly in `for`. Builder methods return a modified value:

| Method | Behavior |
| --- | --- |
| `min_depth(isize)` | Hide entries shallower than the bound while still traversing them. |
| `max_depth(isize)` | Include this depth but never read directories below it. |
| `follow_links(bool)` | Follow file and directory symbolic links, including the root. |
| `contents_first(bool)` | Yield a directory after its descendants instead of before them. |
| `filter_entry((DirEntry) -> bool)` | Reject an entry and prune its entire subtree. Repeated filters are combined with short-circuiting AND. |

Negative bounds or `min_depth > max_depth` produce one `InvalidInput` error before filesystem access, followed by exhaustion. The predicate runs after metadata is obtained, before descending, and even for entries hidden by `min_depth`; it also prunes correctly with `contents_first(true)`. Use `std::iter::filter` when only output should be filtered without pruning.

`WalkIterator` implements `Iterator` with `Item = Result[DirEntry, Error]` and composes with `std::iter`. In preorder, call `skip_current_dir()` immediately after receiving a directory to skip its descendants. The directory is not opened until the next `next()` call. Skipping before iteration, after a file or error, or in contents-first order does nothing. `close()` releases all retained directory descriptors, clears pending work, and permanently exhausts the iterator; it is idempotent. Complete exhaustion also releases all descriptors. When stopping early, including after an error, call `close()` or use `defer`; garbage collection does not close descriptors. Explicit `close()` attempts every descriptor and discards close errors; a close failure during normal iteration is returned as an error. On Linux, close is never retried, including after `EINTR`. Exhaustion is permanent. Copies of an iterator share progress and must be used by one consumer; copies of the builder can start independent walks.

`DirEntry` exposes `path()`, `file_name() -> Option[string]`, `depth()`, `file_type()`, `metadata()`, and `path_is_symlink()`. Metadata is a snapshot taken before yielding the entry. When following a link, `file_type()` and `metadata()` describe the target while `path_is_symlink()` remains true. Paths retain the supplied root spelling and use `path::join` for children; they are not replaced with canonical target paths.

`Error` exposes `path()`, `depth()`, `kind()`, `fs_error()`, and `loop_ancestor() -> Option[string]`, plus `ToString` and `Debug`. The underlying `fs::Error` retains its operation and available raw OS code. A failed entry or directory emits an error and traversal continues with remaining branches. Errors are not suppressed by `min_depth`. A directory-read error follows its entry in preorder and precedes its entry in contents-first order. Device and inode identities from opened directory descriptors detect ancestor loops, including directory links and bind-mount aliases; these return `InvalidData` with the offending path and the ancestor's traversal path, then skip that subtree. Sibling aliases to the same directory are each traversed. Dangling links are ordinary link entries by default and metadata errors when followed; a link chain rejected by the OS reports its filesystem error without a loop ancestor.

```goml
use std::fs::walkdir;
use std::io;

fn main() -> () {
    let tree = walkdir::WalkDir::new(".")
        .max_depth(8)
        .filter_entry(|entry| entry.file_name() != Option::Some(".git"));
    let iterator = tree.iter();
    defer iterator.close();
    for item in iterator {
        match item {
            Ok(entry) => println(entry.path()),
            Err(error) => io::eprintln(error.to_string()),
        }
    }
}
```

Directory entries are read in 32 KiB `getdents64` batches, decoded with record-length and filename validation, and sorted before traversal. Full metadata is still read for each yielded entry to preserve `metadata()` snapshot semantics; `d_type` is not used to omit these queries, so `DT_UNKNOWN` requires no special fallback. Invalid UTF-8 filenames report `InvalidData` for the directory. Paths remain UTF-8 strings, and embedded NUL bytes are rejected before calling the kernel.

One descriptor per active ancestor is retained until its subtree finishes. Descriptors use `O_CLOEXEC`; child lookup uses the parent descriptor, so renaming an opened ancestor does not redirect traversal to a replacement at its old path. Returned paths still reflect the original traversal spelling. Opening a directory checks that its identity matches the earlier metadata; a changed directory produces an error and is skipped. Without link following, `O_NOFOLLOW` also rejects replacement of the final directory component by a symbolic link. Path prefixes and explicit links follow kernel resolution rules; this is not a confinement API or an atomic filesystem snapshot.

Traversal uses an explicit stack rather than recursive calls; memory scales with pending names along the active branch. Very deep trees can reach the process descriptor limit, which produces an OS error for that branch; `max_depth` can bound this usage. Interrupted open, metadata, and directory-read calls retry. Entry failures leave other branches available, so callers must still close an iterator when abandoning it after an error. This API uses ordinary imports, closures, traits, and `for` syntax and adds no grammar or compile-time filesystem access. The native layouts and lifecycle rules follow [getdents64](https://man7.org/linux/man-pages/man2/getdents.2.html), [stat](https://man7.org/linux/man-pages/man2/stat.2.html), [openat](https://man7.org/linux/man-pages/man2/open.2.html), and [close](https://man7.org/linux/man-pages/man2/close.2.html).

### Filesystem notifications

`std::fs::notify` implements filesystem notifications in GoML using `std::os::linux::syscall` and Linux inotify. The supported target is Linux amd64. `watch(path: string)` watches one file or a directory and its immediate entries. `watch_recursive(path: string)` requires a directory, watches existing descendants, and adds watches when directories are created or moved into the tree. Both return `Result[Watcher, fs::Error]`. Descendant symbolic links are not traversed, and a symbolic link as the root is rejected. Directory aliases through bind mounts are unsupported.

`watch_with(path, options: Options)` configures a single watcher. `Options::new()` uses `recursive = false` and `mask = CHANGES`; `with_recursive(bool)` and `with_mask(u32)` return adjusted options. Both fields are public. A requested mask must be a nonempty subset of `ALL_EVENTS`. The available request bits are `ACCESS`, `MODIFY`, `ATTRIB`, `CLOSE_WRITE`, `CLOSE_NOWRITE`, `OPEN`, `MOVED_FROM`, `MOVED_TO`, `CREATE`, `DELETE`, `DELETE_SELF`, and `MOVE_SELF`. `CHANGES` includes these except `ACCESS`, `CLOSE_NOWRITE`, and `OPEN`. Recursive directory topology events, root lifecycle events, and recovery notifications are always delivered even when excluded by the requested filter, so filtering cannot disable recursive maintenance.

`Event` has public `path: string`, `mask: u32`, `cookie: u32`, and `rescan: bool` fields. Paths are absolute. `has(mask)` tests whether any requested mask bit is present, and `is_dir()` tests `IS_DIR`. In addition to requested event bits, masks may contain `UNMOUNT`, `Q_OVERFLOW`, `IGNORED`, or `IS_DIR`. Matching nonzero cookies connect `MOVED_FROM` and `MOVED_TO` events within a registration; cookies are not persistent object identities.

Build ignore rules with `Options::new().with_ignored_names(Vec::from_array([".git", "node_modules"]))` or `with_ignore((absolute_path: string, is_directory: bool) -> bool)`. Repeated calls combine rules with OR. Name rules match literal basenames at every depth and snapshot the supplied vector; they are not glob patterns. An ignored directory prunes its entire subtree, avoiding recursive watches and scans there. Rules apply to initial discovery, new directories, renames, and overflow recovery, including registrations in `WatchSet` and subscriptions. Moving a visible directory into an ignored path removes its subtree watches; moving it back installs watches and requests a rescan. The root itself and recovery signals are never ignored. Construct options through `new()`; predicates must be stable, quick, and must not call back into their watcher because they run while its state is locked.

`Watcher.try_read() -> Result[Vec[Event], fs::Error]` reads one available batch without waiting for new kernel events. `Watcher.read(timeout: time::Duration)` has the same return type and waits for a nonempty batch, returning an empty vector on timeout. A zero timeout checks immediately. `read_with(cancel: task::CancelToken, timeout)` returns `Result[task::WaitResult[Vec[Event]], fs::Error]`; cancellation returns `Cancelled` and leaves the watcher available. A read that wins a race with cancellation may return `Completed(events)`. Waits check cancellation and close in intervals of at most 50 ms. Timeouts and cancellation bound waiting for kernel events, not directory scans or time spent waiting for another operation to release shared state.

Single-watcher copies share a synchronized handle. Concurrent reads divide the event stream. `close() -> Result[(), fs::Error]` releases the descriptor and all watches and is idempotent; `is_closed()` reports terminal state. Use explicit close or `defer`; garbage collection does not close watchers. A single-file watcher follows the watched inode until a terminal event. Editors commonly replace a file atomically, which ends that file watch; monitor its parent directory and filter event paths when replacement must be followed.

```goml
use std::fs;
use std::fs::notify;
use std::time;

fn monitor_once(directory: string) -> Result[(), fs::Error] {
    let watcher = notify::watch_recursive(directory)?;
    defer {
        let _ = watcher.close();
    };
    for event in watcher.read(time::Duration::from_seconds(5))? {
        if event.rescan {
            println("rescan " + event.path);
        }
        if event.has(notify::CREATE | notify::MODIFY | notify::DELETE) {
            println(event.path);
        }
    }
    Result::Ok(())
}
```

`WatchSet::new()` creates an initially empty manager for multiple independent registrations. It does not open a kernel descriptor until a path is added.

| Method | Result and behavior |
| --- | --- |
| `add(path)` | `Result[WatchId, fs::Error]`; add a file or nonrecursive directory |
| `add_recursive(path)` | `Result[WatchId, fs::Error]`; add a directory tree |
| `add_with(path, options)` | `Result[WatchId, fs::Error]`; add with explicit options |
| `remove(id)` | `Result[bool, fs::Error]`; close that registration; `false` means it was absent |
| `watches()` | `Result[Vec[WatchInfo], fs::Error]`; snapshot of public `id`, absolute `path`, and `options` |
| `try_read()` | `Result[Vec[Notice], fs::Error]`; read at most one kernel batch per registration |
| `read(timeout)` | Same result; wait for notices or return an empty vector on timeout |
| `read_with(cancel, timeout)` | `Result[task::WaitResult[Vec[Notice]], fs::Error]`; cancellable waiting |
| `close()` / `is_closed()` | Close every registration idempotently / query manager state |

`WatchId` supports equality, hashing, debug output, and `value() -> u64`. IDs are local to one `WatchSet` and are not reused during its lifetime. Duplicate or overlapping paths receive independent IDs and may produce duplicate events, each tagged with its registration. Removing one ID leaves other registrations intact. Registration failure cleans up the attempted registration and leaves existing registrations unchanged. Each registered root uses a separate inotify instance, descriptor, and event queue; recursive roots additionally consume one kernel watch per directory. Kernel limits such as `max_user_instances`, `max_user_watches`, and the process descriptor limit are reported as structured errors with numeric errno.

`Notice` is `Event(WatchId, Event)`, `Error(WatchId, fs::Error)`, or `Removed(WatchId)`. A terminal root event is followed by `Removed`; a decoding or maintenance failure produces `Error` followed by `Removed`. Such failures affect only that registration. The manager remains available, even after its last registration ends, and accepts new paths. Explicit `remove` reports its result directly and does not enqueue a `Removed` notice. Events are ordered within each registration; the set promises no total order across roots and collects from every root to prevent one busy root from starving another.

Both `Watcher` and `WatchSet` provide `subscribe(scope: task::Scope, capacity: isize)`, returning `Result[Subscription[Event], fs::Error]` and `Result[Subscription[Notice], fs::Error]`, respectively. The subscription starts a task in the supplied scope and takes responsibility for closing the source on exit. Its `events()` returns a `Receiver[Result[Vec[T], fs::Error]]`; `done()` becomes ready after the worker has finished. A single-watcher error is delivered as `Err` before the channel closes. Per-registration errors in a set remain `Notice::Error` values, allowing other registrations to continue.

The channel capacity counts batches, not individual events. Zero is an unbuffered channel; a negative capacity is rejected. A full channel applies backpressure and does not silently drop delivered batches. The kernel queue can still overflow while a consumer is slow, and recovery is reported through `rescan`. `Subscription.close()` signals shutdown, joins the worker, closes the source, and returns its close result; copies can call it repeatedly or concurrently. Scope cancellation and closing the underlying source also unblock a producer waiting to send. Shutdown may discard an undelivered batch, while already buffered batches can be drained. A rejected subscription, including one attempted in an already cancelled scope, leaves its source available. Consume through one subscription, or through explicit reads; mixing them divides events between consumers.

```goml
use std::fs;
use std::fs::notify;
use std::task;
use std::io;

fn monitor_pair(scope: task::Scope, first: string, second: string) -> Result[(), fs::Error] {
    let watches = notify::WatchSet::new();
    defer {
        let _ = watches.close();
    };
    let _ = watches.add_recursive(first)?;
    let _ = watches.add_recursive(second)?;
    let stream = watches.subscribe(scope, 16)?;
    defer {
        let _ = stream.close();
    };
    while let Some(batch) = stream.events().recv() {
        for notice in batch? {
            match notice {
                notify::Notice::Event(id, event) => {
                    println(id.value().to_string() + ": " + event.path);
                    if event.rescan {
                        println("rescan " + event.path);
                    }
                },
                notify::Notice::Error(_, error) => io::eprintln(error.to_string()),
                notify::Notice::Removed(_) => (),
            }
        }
    }
    Result::Ok(())
}
```

Without a subscription, call a reading method regularly: recursive registration and rename maintenance run while consuming kernel events. For an internal directory rename, the watcher matches cookies across read batches and updates descendant paths. While a move is unresolved, events inside that subtree are suppressed to avoid reporting paths outside the root. Unmatched moves expire after 100 ms when the input queue becomes empty. Reusing an old path creates a new subscription to that directory without retaining the moved-out tree. If a directory scan discovers a relocation before its queued rename records have been processed, the library reconciles descriptor paths and requests a root rescan rather than treating that ordinary race as a fatal alias error.

When a single watcher's root moves, is deleted, or is unmounted, the terminal event is returned and the watcher closes; subsequent reads return `InvalidInput`. Errors during decoding or watch maintenance also close that watcher, so failures cannot leave a silently incomplete subscription. Errors retain the operation, path when applicable, and numeric errno; unsupported targets return `Unsupported` during registration. Permission failures and exhausted watch limits are not retried indefinitely.

Inotify does not provide an atomic recursive subscription. Files may change before a newly discovered directory gets its own watch. A directory `CREATE` or `MOVED_TO` event therefore sets `rescan = true`, requesting that the caller refresh that subtree's contents. On `Q_OVERFLOW`, the library recreates the affected inotify descriptor and all its watches, discards the remaining stale batch, and returns a root event with `rescan = true`. A synthetic root event with `mask = 0` and `rescan = true` requests reconciliation after paths were discovered ahead of queued rename records. Applications maintaining a cache must rescan the indicated path, then continue processing queued events. If rebuilding fails, the registration closes and reports an error. Initial registration similarly requires the caller's own scan if an initial snapshot is needed.

Notifications may be coalesced, and paths can change again before events are consumed; they are not an audit log. Non-UTF-8 event names produce `InvalidData` rather than replacement characters, consistent with the standard library's UTF-8 path API. Network-filesystem remote changes, mounts placed over watched paths, and memory-mapped writes have the underlying [inotify limitations](https://man7.org/linux/man-pages/man7/inotify.7.html). This API uses ordinary imports, structs, functions, and methods; no grammar changes are introduced.

### Bincode typed binary data

`std::bincode` is primarily a typed serde format. `encode_to_vec` writes the requested type directly to the output buffer, while `decode_from_slice` lets the requested type pull its fields directly from the input and returns the value together with the number of consumed bytes. The typed path does not construct `serde::Value` or request a runtime schema. The lower-level `encode_value` and `decode_value` functions remain available for explicit dynamic work; `decode_value` needs an explicitly supplied schema because bincode does not carry field types, tuple lengths, or struct layouts on the wire.

The typed encoder and decoder validate every compound begin, element or field, and end transition. A malformed handwritten serde implementation returns an error instead of producing partial data or panicking. Decode errors include both the byte offset and a path through positional fields, collection elements, map entries, and enum variants.

`bincode::standard()` uses little-endian variable integer encoding. `bincode::legacy()` uses little-endian fixed-width integers. `with_little_endian`, `with_big_endian`, `with_variable_int_encoding`, and `with_fixed_int_encoding` return adjusted configurations. The wire representation follows bincode 2 conventions for booleans, ZigZag signed varints, integer markers, IEEE floating-point bits, UTF-8 strings, collection lengths, one-byte option tags, source-order struct fields, and enum declaration indexes.

```goml
use std::bincode;
use bincode::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
struct Message {
    id: u32,
    text: string,
}

fn round_trip(value: Message) -> Result[Message, string] {
    let encoded = bincode::encode_to_vec(value, bincode::standard())?;
    let decoded: (Message, isize) = bincode::decode_from_slice(
        encoded.slice(0, encoded.len()),
        bincode::standard(),
    )?;
    Result::Ok(decoded.0)
}
```

Typed bincode supports the shared serde primitives, byte slices through custom direct implementations, `Vec`, `Option`, two- and three-element tuples, and derived structs and enums. Fixed arrays are not yet supported because GoML does not yet have const-generic serde implementations. Dynamic `serde::Value`, `json::Value`, and `toml::Value` do not describe the concrete binary layout needed by typed bincode decode.

### TOML values and typed documents

`std::toml` follows the same public split as JSON. `toml::Value` has string, signed 64-bit integer, 64-bit float, boolean, datetime text, array, and table variants. `parse` and `encode` operate on that schema-free representation. `to_value`, `from_value`, `to_string`, and `from_string` use the shared serde traits and enforce the destination type. TOML still plans a complete table tree before emission because headers and dotted paths require document-wide organization; it is compatible with the streaming traits but is not currently a fully direct format. Unsigned values above the TOML signed-integer range are rejected, and `()` or `None` cannot be encoded because TOML has no null value.

The TOML parser accepts basic and literal strings, Unicode escapes, booleans, decimal and base-prefixed integers, floats, datetime text, arrays, inline tables, dotted keys, and ordinary table headers. It preserves table and field order for deterministic output. Multiline strings, array-of-table headers, and dotted keys inside inline tables are not yet supported.

```goml
use std::toml;
use toml::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
struct Server {
    host: string,
    port: u16,
}

fn load(input: string) -> Result[Server, string] {
    toml::from_string(input)
}
```

### Comparison and sorting

Sorting mutates a `Vec[T]` in place. `sort` and `stable_sort` use `cmp::Ord`; `sort_by_ordering` and `stable_sort_by_ordering` use `cmp::Ordering`; `sort_by` and `stable_sort_by` accept negative/zero/positive integer comparators. All sorting variants are stable. `binary_search` and its comparator-based variants expect the vector to already be ordered and return the first matching index.

`std::cmp` provides `PartialOrd` and `Ord`. `()`, `bool`, `string`, `char`, and all signed and unsigned integer types implement both. Floating-point values implement only `PartialOrd`, because NaN does not form a total order; comparison with NaN yields `Option::None`. Tuples, fixed arrays, `Vec`, `Slice`, `Option`, and `Result` implement the comparison traits conditionally and use lexicographic order. `cmp::compare` returns `Ordering::Less`, `Equal`, or `Greater`; `Ordering` supports predicates, reversal, lexicographic chaining, and conversion to the negative/zero/positive integer convention. `cmp::Reverse[T]` reverses an existing ordering. `cmp::clamp` returns an error when the minimum exceeds the maximum.

### Structured concurrency

`std::task` creates lexical task scopes on top of goroutines:

```goml
use std::task;

fn load_pair() -> Result[(isize, isize), string] {
    task::try_scope(
        |scope: task::Scope| {
            let left = scope.spawn_try(|cancel| load_left(cancel));
            let right = scope.spawn_try(|cancel| load_right(cancel));
            Result::Ok((left.join()?, right.join()?))
        },
    )
}
```

The contextual `scope` and `spawn` forms provide the same structured lifetime with a hidden scope capability:

```goml
use std::task;

fn load_pair() -> Result[(isize, isize), string] {
    scope {
        let left = spawn |cancel| load_left(cancel);
        let right = spawn |cancel| load_right(cancel);
        Result::Ok((left.join()?, right.join()?))
    }
}
```

`scope { body }` is lowered before HIR checking to `task::scope(|hidden_scope| body)`, and each directly nested `spawn |cancel| body` is lowered to `hidden_scope.spawn(|cancel| body)`. The file must import `std::task`. Directly nested lexical scopes use `scope_with` and inherit their parent's cancellation. The hidden scope value cannot be named or returned. A lexical `spawn` cannot cross a user closure boundary, which prevents a returned closure from capturing the capability; create a nested `scope` inside that closure instead. `scope` and `spawn` remain ordinary identifiers outside these contextual forms.

`scope` stops accepting new work after its body returns, waits for every direct child task, and then returns the body's value. It still waits when a `Task[T]` handle is discarded. `Task::join` may be called repeatedly or concurrently; every call observes the same stored result. `Scope::try_spawn` returns `Some(task)` when it registers the task before closing begins and `None` after the scope starts closing. Ordinary `Scope::spawn` retains the stricter runtime-error behavior for closed scopes.

`try_scope` cancels its scope when the body returns `Result::Err`, waits for all direct children to exit, and then returns the original error. `Scope::spawn_try` also cancels sibling tasks as soon as its child returns `Result::Err`. A nested scope inherits cancellation when it is created with `scope_with(parent_token, body)` or `try_scope_with(parent_token, body)`.

Cancellation is cooperative. `Scope::cancel` changes the state observed by `CancelToken::is_cancelled`; it does not forcibly terminate a goroutine. Blocking work can use:

- `task::recv_with(token, channel) -> WaitResult[Option[T]]`
- `task::send_with(token, channel, value) -> WaitResult[()]`
- `time::sleep_with(token, duration) -> WaitResult[()]`
- `command.output_cancel_structured(token) -> Result[process::Output, process::Error]`
- `command.status_cancel_structured(token) -> Result[process::ExitStatus, process::Error]`

`CancelToken::done() -> Receiver[()]` exposes the scope context's shared completion channel, and `Task::done() -> Receiver[()]` exposes the task's shared ready channel. Neither method starts a bridge goroutine. A task stores its result before closing the ready channel, so `join()` is immediately observable after its completion event.

`std::time::Timer::new(duration)` creates a stoppable one-shot timer. `done()` returns its `Receiver[()]`, `stop()` reports whether it prevented a pending firing, and `time::after(duration)` is the one-shot convenience form. A successful stop leaves the completion channel unready. Timer firing and stopping are synchronized so the channel closes at most once.

Channel and sleep operations return `WaitResult::Cancelled` when cancellation wakes them. Process operations instead return `process::Error` with kind `io::ErrorKind::Interrupted`. Process cancellation uses the host command context, so the scope waits for the process operation to return before it exits. Task scopes never close user channels automatically. `active_scope_count()` exposes the number of live runtime scopes for tests and leak diagnostics.

GoML has no lifetime or linear type system, so a `Scope` value can currently escape its body. Calling `spawn` after the scope begins closing is a runtime error. Panic remains a fatal runtime exception and is not converted into `Result`. A panic in the scope body or a child task cancels sibling tasks, waits for them, removes the runtime scope, and is then re-raised in the scope owner.

`join_all` returns values in input order. `join_all_results` waits for every task and returns errors in input order, independent of goroutine scheduling.

`join_all_indexed_results` preserves each error's input index. `for_each_concurrent(limit, values, body)` runs at most `limit` calls at once, uses one worker when the limit is non-positive, waits for every value, and returns indexed errors in input order. `ConcurrencyLimit::run` can apply the same cooperative limit to custom task layouts.

`race(bodies)` returns the first completed value, cancels the remaining bodies, and still waits for every losing body to exit. It returns `None` for an empty input. A body that does not cooperate with cancellation can therefore delay the return from `race`.

### `collections::IndexMap[K, V]`

`IndexMap` is an insertion-ordered hash map. Its key type must implement `Eq` and `Hash`, while actual key comparison uses `PartialEq`:

```goml
use std::collections;

let headers: collections::IndexMap[string, string] = collections::IndexMap::new();
headers.insert("content-type", "text/plain");
headers.insert("content-length", "12");
headers.insert("content-type", "application/json");
for (name, value) in headers {
    println(name + ": " + value)
}
```

Inserting a new key appends it to the iteration order. Replacing an existing value does not move the key. Removing and inserting the key again appends it to the end.

Common methods are `new`, `with_capacity`, `len`, `is_empty`, `contains`, `get`, `insert`, `remove`, `reserve`, `clear`, `entries`, `keys`, `values`, and `iter`. `insert` and `remove` return the previous value as `Option[V]`. `entries`, `keys`, and `values` return ordered snapshots, while `iter` and `for` traverse `(K, V)` pairs in insertion order.

The implementation uses a sparse open-addressed index table and an insertion-ordered entry array. Deleted entries become tombstones and are compacted during later growth or when deletion density becomes high. Lookup, insertion, and removal are expected O(1); iteration and compaction are O(n). Structural mutation while an iterator is active is unsupported.

`IndexMap` does not currently have literal or indexing syntax. Use `insert` and `get`.

## Comparison of common writing errors

| Avoid | GoML form |
| --- | --- |
| `Vec<isize>` | `Vec[isize]` |
| `fn id<T>(x: T) -> T` | `fn id[T](x: T) -> T` |
| `id::<i32>(1)` | `id::[i32](1)` |
| Ordinary function `id[i32](1)` | `id::[i32](1)`, or rely on parameter/result type inference |
| `1i32`, `1u64`, `1.0f32` | Use the expected type, such as `let value: u64 = 1;` |
| Non-ASCII source text in `b"é"` | Use a `string`, or write its encoded bytes explicitly such as `b"\xC3\xA9"` |
| A bare block as a control-flow header value | Parenthesize it, for example `if ({ prepare(); ready() }) { ... }` |
| `let mut x: &T` | Use value `T` or `Ref[T]` as required |
| Write `if cond { value }` in the value position | `if cond { value } else { other }` |
| `let Option::Some(x) = value;` | `let Some(x) = value else { return };`, `if let`, or `match` |
| `let Point { x } = point;` | `let Point { x, .. } = point;` |
| Non-exhaustive `match` | Cover every possible variant or add a `_` branch |
| `x++`, `x--` | `x += 1;`, `x -= 1;` |
| Assign through an immutable structure binding | Declare the binding with `let mut`, or create a new value with `Point { field: value, ..point }` |
| `var x = 1`, `x := 1` | `let x = 1;` |
| A loop with a condition | `while condition { ... }` |
| `for i := 0; ...` | `while`, or `for i in start..end` |
| `switch` | `match` |
| `null`, `nil` | `Option::None` for optional values; `ffi::null()` for raw Go pointers; `ffi::nil_error()` for Go errors |
| `throw`, exception | `Result` and `?` |
| `float_value.to_i32()` | Floating point to integer conversion is not supported; use dedicated parsing or conversion APIs |
| `dyn A + B` | Use one dyn-safe trait; multiple bounds are reserved syntax but not yet supported |
| `dyn TraitWithAssociatedType` | Bind every associated type, for example `dyn Iterator[Item = isize]` |
| Use `type UserId = u64;` when `UserId` must be distinct | Use `struct UserId(u64);` and construct it explicitly |
| `use pkg::*` | List the required public items explicitly with `use pkg::{A, B};` |
| `mod`, `crate::`, `super::` | Directory packages, `module::path` for the current module, and canonical paths for dependencies |
| `fn helper` inside function | Top-level function or local closure |
| Go external type | `#[go_type("pkg", "Name")] extern type Name[T];` retains Go identity and validates concrete instances across packages and artifacts; `std::ffi::Ptr[T]` and Go pointer aliases preserve nullable pointer values, with explicit `ffi::null()` and `ffi::is_nil`; method bindings use `#[go_method("Method")]`; symbolic instances remain unsupported |
| Go interface adapter | `#[go_interface(RawType, Wrapper, method = "GoMethod")]` generates a checked native-interface wrapper and trait implementation; `from_trait` explicitly creates a typed Go bridge retaining the supplied dyn object; nil/typed-nil and multiple results are preserved |
| Go binding generator | `goml bind-go <CONFIG>` selects explicit package/symbol allowlists and finite Go-checked generic arguments; emits raw bindings with protected deterministic output |
| Go-callable export | Annotate a supported public function with `#[go_export("Name")]` and generate a Go package with `goml export-go` |
| Traverse a directory tree | `std::fs::walkdir` Linux amd64 syscall-backed depth-first iteration with depth bounds, pruning, optional link following, and per-path errors |
| Watch a directory tree for changes | Use `fs::notify::watch_recursive` or `WatchSet` on Linux amd64, prune ignored paths through `Options`, consume timed reads or scoped subscriptions, handle `Event.rescan`, and close the handle |
| TCP and UDP networking | `std::net` Linux amd64 syscall-backed sockets with numeric IPv4/IPv6 addresses, shared epoll readiness, explicit close, timeouts, and task cancellation |
| Treat a GoML integer or struct as a kernel pointer/layout | Use `syscall::Arg::Buffer` for synchronous byte storage on Linux amd64; encode the native ABI explicitly and inspect `SyscallResult.errno` |
| Unannotated user `extern fn` | Use a normal GoML function or `#[go_ffi("import/path", "ExportedSymbol")] extern fn`; project commands validate Go calls by default (`--ffi-check required`) |
| Call an ordinary function from `comptime` | Mark a supported free function with `#[comptime]` |
| Capture a runtime local in `comptime` | Pass a literal or compile-time value to a `#[comptime]` function |

## Informal Grammar Quick Facts

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
extern_type = "extern" "type" upper_ident generic_params? ";"
go_ffi_attribute = "#[" "go_ffi" "(" string_literal "," string_literal ")" "]"
go_method_attribute = "#[" "go_method" "(" string_literal ")" "]"
go_interface_attribute = "#[" "go_interface" "(" path "," ident ("," ident "=" string_literal)* ","? ")" "]"
go_ffi_extern = (go_ffi_attribute | go_method_attribute) visibility? "extern" "fn" lower_ident
                param_list return_type? ";"

function      = "fn" lower_ident generic_params? param_list return_type? where_clause? block
type_alias    = "type" upper_ident type_names? "=" type ";"
constant      = "const" ident ":" type "=" expression ";"
static        = "static" ident ":" type "=" expression ";"
method        = visibility? "fn" lower_ident generic_params? param_list return_type? where_clause? block
generic_params = "[" generic_param ("," generic_param)* "]"
generic_param = upper_ident (":" trait_set)?
param_list    = "(" (parameter ("," parameter)*)? ")"
parameter     = lower_ident ":" type | "self"
return_type   = "->" type

struct_def    = "struct" upper_ident type_names?
                ("{" struct_fields? "}" | "(" newtype_field ","? ")" ";")
struct_fields = struct_field ("," struct_field)* ","?
struct_field  = visibility? lower_ident ":" type
newtype_field = visibility? type
enum_def      = "enum" upper_ident type_names? "{" variants? "}"
variants      = variant ("," variant)* ","?
variant       = upper_ident
              | upper_ident "(" type_list? ")"
              | upper_ident "{" variant_fields? "}"
variant_fields = lower_ident ":" type ("," lower_ident ":" type)* ","?
type_names    = "[" upper_ident ("," upper_ident)* "]"

trait_def     = "trait" upper_ident generic_params? (":" trait_set)? where_clause?
                "{" trait_member* "}"
trait_member  = "type" upper_ident (":" trait_set)? ";"
              | "fn" lower_ident generic_params? param_list return_type? where_clause?
                (";" | block)

impl_def      = "impl" generic_params? trait_ref "for" type where_clause?
                "{" impl_member* "}"
              | "impl" generic_params? type where_clause?
                "{" method* "}"
impl_member   = "type" upper_ident "=" type ";" | method

trait_set     = trait_ref ("+" trait_ref)*
trait_ref     = path type_args?
where_clause  = "where" where_predicate ("," where_predicate)* ","?
where_predicate = type ":" trait_set | type "=" type

type          = primitive_type
              | path type_args?
              | dyn_type
              | "[" type ";" integer_literal "]"
              | "(" type_list ")"
              | type "->" type
primitive_type = "()" | "bool" | "isize" | "i8" | "i16" | "i32" | "i64"
               | "usize" | "u8" | "u16" | "u32" | "u64" | "f32" | "f64"
               | "string" | "char"
type_args     = "[" type_list "]"
type_list     = type ("," type)* ","?
dyn_bound     = path dyn_args?
dyn_args      = "[" (type ",")* dyn_assoc ("," dyn_assoc)* ","? "]"
              | "[" type_list "]"
dyn_assoc     = upper_ident "=" type
dyn_type      = "dyn" dyn_bound ("+" dyn_bound)*

block         = "{" statement* expression? "}"
statement     = "let" "mut"? pattern (":" type)? "=" expression ("else" block)? ";"
              | "defer" expression ";"
              | assign_target assignment_operator expression ";"
              | expression ";"
              | unlabeled_control_expression

assignment_operator = "=" | "+=" | "-=" | "*=" | "/=" | "%="
                    | "&=" | "|=" | "^=" | "<<=" | ">>="

expression    = literal | raw_string | byte_string | raw_byte_string
              | interpolated_string | path | tuple | array | block
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

interpolated_string = "f\"" (string_text | "{{" | "}}" | "{" expression "}")* "\""
byte_string   = "b\"" byte_string_content* "\""
raw_byte_string = "br\"...\"" | "br#\"...\"#" | "br##\"...\"##" | ...
raw_string    = "r" raw_hashes? "\"" raw_text "\"" raw_hashes?
struct_literal = path "{" (struct_literal_field ("," struct_literal_field)*
                 ("," ".." expression)? ","? | ".." expression ","?)? "}"
struct_literal_field = lower_ident (":" expression)?

unlabeled_control_expression = if_expression | match_expression | select_expression
                   | "while" expression block | "while" "let" pattern "=" expression block
                   | "loop" block | "for" pattern "in" expression block
if_expression = "if" expression block ("else" (block | if_expression))?
              | "if" "let" pattern "=" expression block
                ("else" (block | if_expression))?
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
while_expression = loop_label_decl? "while" expression block
              | loop_label_decl? "while" "let" pattern "=" expression block
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

The parser will do some error recovery for commas and semicolons, but the code agent should always generate the above canonical form: list items separated by commas, `let`, assignments and ordinary non-tail expressions with semicolons, unlabeled control-flow statements without unnecessary semicolons, labeled loop statements with semicolons before following statements, and trait method signatures with semicolons. The sequence pattern contains at most one rest; the `..` in the structure pattern appears at most once and must be at the end.

## Verify generated code

After running `just make`, verify a standalone source file from the repository root:

```sh
stage2/bin/gomlc run-single path/to/main.gom
```

For a project, run the installed driver from anywhere inside its module:

```sh
goml fmt
goml check
goml build
goml check --tests
goml test
goml run
```

`goml check`, `goml build`, `goml test`, and `goml fmt` always operate on the complete module and do not accept package or file targets. `goml run [TARGET]` accepts an optional entry package file or directory when a module has multiple executable packages.

When you need to inspect a compilation phase, add `--dump-ast`, `--dump-expanded-ast`, `--dump-hir`, `--dump-tast`, `--dump-ctir`, `--dump-core`, `--dump-mono`, `--dump-lift`, `--dump-anf`, or `--dump-go` to `gomlc run-single`. `--dump-ast` shows source lowering before derive expansion, while `--dump-expanded-ast` includes every generated implementation.

The code agent should at least run the corresponding `goml check` or `gomlc run-single` before submitting the source code; when modifying the test, it should also run `goml check --tests` and the related `goml test`. When type inference fails, give priority to adding local result types, empty container types, closure parameter types, or using UFCS instead of rewriting to unsupported Rust/Go syntax.
