# GoML language guide

This guide describes the GoML syntax, type rules, tools, and public library APIs implemented in this repository. Use it when writing, modifying, and reviewing `.gom` source. Similarities to Rust, Go, or OCaml do not imply that their syntax or APIs are supported.

GoML is a statically typed language with garbage collection. Its syntax is close to Rust, while its semantics are closer to ML. The compiler monomorphizes generics and lambda-lifts GoML closures before emitting Go. Explicit Go FFI can preserve native generic types and function values. GoML has no ownership, borrowing, lifetimes, or manual memory management.

Generated Go code and Go FFI use Go 1.26 as their compatibility baseline. Building the toolchain, compiling generated programs, and validating Go bindings require Go 1.26 or newer. The independent [gomlgo frontend and interpreter](../gomlgo/README.md) targets Go 1.26 language and standard-library behavior; its execution and differential tests require Go 1.26.x.

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
- The fixed-width vector and mask type spellings listed under [Portable SIMD](#portable-simd), including `u8x16`, `i32x8`, `f32x4`, `f64x4`, and `mask32x4`, are also accepted as structure and type-alias names and in type positions. They are ordinary scoped names, supplied by `std::simd`, rather than globally reserved keywords;
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

### Native dependencies

A module containing a Go adapter can declare its Go module and system prerequisites in the same manifest:

```toml
[native]
go-module = "example.com/goml-ecosystem/llvm"
go-version = "v0.0.0"
cgo = "required"
llvm-major = "18"
```

All values are quoted strings. `go-module` must match the adapter's module-root `go.mod`; `go-version` defaults to `v0.0.0` and can specify a Go module version appropriate for a versioned module path. `cgo` accepts `required` or `optional` (the default). `llvm-major` requires a positive decimal major version. Unknown and duplicate native keys are errors.

The consuming project supplies a minimal module-root `go.mod`. For Go FFI commands, the driver resolves declared adapters from the selected GoML registry sources and generates Go requirements and local replacements under `<target-dir>/native/<hash>/`. It passes the generated module file through `GOFLAGS=-modfile=...` to metadata validation, compilation, tests, and Go export. Source `go.mod` and `go.sum` files remain unchanged. The content hash includes adapter locations and module metadata; dry runs do not run prerequisite probes or create these files. An existing `GOFLAGS -modfile` conflicts with managed dependencies and produces a diagnostic.

The driver checks required `CGO_ENABLED=1` and the requested LLVM major using `LLVM_CONFIG`, `llvm-config-N`, `/usr/lib/llvm-N/bin/llvm-config`, then `llvm-config`. These declarations do not install system packages or configure compiler/linker search paths. Consumers still configure `CGO_CFLAGS`, `CGO_LDFLAGS`, and the shared-library loader when needed. Additional external Go dependencies must satisfy the existing Go FFI read-only module policy; dependency-local replacements are not promoted into the consumer. Registry packages must include their adapter sources and Go module files.

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

An expression of type `never` does not produce a value and can appear where another value type is expected. This does not implicitly convert container or function types: `Vec[never]` is not `Vec[isize]`, and `() -> never` is not `() -> isize`. To adapt a non-returning function to a callback with another return type, use a closure with that expected callback type. Ordinary values and normally returning callbacks cannot satisfy `never` annotations.

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

On Linux amd64, `std::os::linux::syscall` accepts numeric machine words, scoped byte-buffer arguments, and explicit pointer fields between byte buffers for low-level kernel calls. It does not add pointer casts, pointer arithmetic, or native struct layout to the language.

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

`OnceCell.get_or_init(init)` runs one initializer and caches its result. Concurrent first callers wait for that initializer and receive the same value. If the initializer panics, the cell becomes uninitialized again and waiting callers are woken so a later initializer can retry. Recursive initialization of the same cell panics with an error that names the static. A cached `Result` is an ordinary cached value. Statics do not run user-observable destruction at process exit.

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

`#[comptime]` marks a non-generic free function as compile-time-capable. The function remains callable at runtime. A compile-time call may call other `#[comptime]` free functions, the builtin operations listed below, or the `compile_error(string) -> never` intrinsic. The compiler validates the complete body of every marked function, including branches not taken by a particular invocation. Attributes with arguments, duplicate attributes, generic functions, methods, extern functions, and other declarations are rejected.

A top-level constant initializer is an implicit compile-time context, so `const SIX: isize = factorial(3);` and an initializer wrapped in `comptime { ... }` are equivalent. A top-level constant may produce a recursively immutable tuple, fixed array, struct, or enum in addition to scalar values. A `comptime` expression in ordinary code supports the same reifiable value shapes.

Compile-time code may use local bindings and assignment, blocks, `if`, `match`, `while`, `loop`, restricted `for`, `break`, `continue`, `return`, recursion, direct calls, integer conversion methods, integer `to_string()`, and supported operators. A compile-time `for` accepts only a fixed array or the builtin `isize` ranges `start..end` and `start..=end`; its source and range endpoints are evaluated once, and its pattern must be irrefutable. The deterministic string methods `len`, `byte_len`, `get`, `byte_get`, `byte_slice`, `is_char_boundary`, `starts_with`, `ends_with`, and `contains` are also available. String indexes and slices use byte offsets and reject invalid UTF-8 character boundaries. Integer formatting accepts signed and unsigned 8-, 16-, 32-, and 64-bit values, including `byte`, `isize`, and `usize`, and produces decimal text with the full range preserved. This is restricted to the builtin `ToString` implementation; a user trait with the same short name does not become a compile-time intrinsic. For example, `const BITS: u64 = 18446744073709551615; const MASK: string = BITS.to_string();` evaluates without runtime formatting.

Compile-time code cannot capture a surrounding runtime parameter or local. Closures, indirect calls, generic functions, methods other than the integer conversions, integer `to_string`, and string whitelist, trait or dynamic dispatch, general iterators, floating-point computation, `Ref`, `Vec`, `HashMap`, channels, goroutines, extern calls, host I/O, environment access, time, randomness, network access, general type reflection, arbitrary declaration generation, compile-time parameters, value generics, and type-level computation are not supported. The constrained programmable derive interface described below is the only reflection and code-generation facility.

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

A free function may use a type parameter only in its body, with no occurrence in its parameter or result types. Supply such arguments explicitly, for example `schema::[User]()`, including when taking a function value. Those arguments are retained across package interfaces and specialization; an unused generic definition does not generate an unspecialized executable body.

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
- Import `std::num::{ToFloat, TryToInt}` for scalar integer `.to_f32()` / `.to_f64()` and checked float `.try_to_i32()` / `.try_to_u64()` conversions, with analogous methods for every integer width. Float-to-integer methods truncate toward zero and reject nonfinite or out-of-range results.
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

Concrete and nested specializations also support associated constructors, for example `impl Box[f64]` and `impl[T] Box[Vec[T]]`. In `Box::[Vec[string]]::repeated("x")`, the owner argument denotes `Vec[string]`, while the implementation parameter `T` is `string`. Omitting owner arguments creates inference variables; an expected result such as `let value: Box[f64] = Box::empty();` can select a specialization. If several candidates remain, use explicit owner arguments; ambiguous lookup reports a diagnostic.

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

Trait bounds make the corresponding methods available on generic values. Imported short trait names, renamed imports, and public re-exports retain the defining trait identity in generic bounds and `where` predicates through specialization. Supertraits and associated type bounds also participate in method resolution as implied constraints.

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

`serde::Value` remains the explicit dynamic data model. It retains exact signed and unsigned integer widths, floating-point widths, enum declaration indexes, field order, and variant shape. `Binary(Vec[byte])`, `Map(Vec[(Value, Value)])`, and `Extension { tag: i8, data: Vec[byte] }` preserve binary payloads, ordered arbitrary-key map entries (including duplicates), and opaque extension data. `Sequence`, `Tuple`, and `Optional` are distinct, and `Value::Number(string)` represents a textual number whose destination type is not yet known. `serde::to_value` and `serde::from_value` connect typed values to this model through `ValueSerializer` and `ValueDeserializer`; both return `Result`, and using them intentionally constructs or consumes a complete value tree. Unit, booleans, strings, chars, numeric primitives, `Value`, `Vec[T]`, `Option[T]`, and two- or three-element tuples have standard direct implementations.

`Serializer::serialize_extension(tag, data)` and `Deserializer::deserialize_extension()` are optional events with recoverable unsupported defaults. Format implementations opt in explicitly. The value serializer copies binary and extension payloads and retains maps as `Value::Map`; the value deserializer also accepts the legacy sequence-of-key/value-tuples map representation. JSON and TOML reject extension events, and Bincode has no extension representation. TOML preserves its previous value-bridge encoding of binary as integer arrays and maps as arrays of key/value pairs. `json::try_from_serde_value` rejects extensions recursively. The older infallible `json::from_serde_value` remains a lossy projection: binary becomes an integer array, maps become arrays of pairs, and extensions become `{ "tag": ..., "data": [...] }`. Prefer checked conversion for format validation.

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

`Serializer::is_human_readable()` and `Deserializer::is_human_readable()` default to `true`. JSON, TOML, and the generic value adapters use this default; Bincode overrides it with `false`. Custom formats should override both sides consistently. A custom Serde implementation can choose readable names for text formats and an integer representation for binary formats without depending on a particular encoder. This query does not change ordinary derived serialization.

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

The compiler resolves handlers from already compiled dependency interfaces. A handler cannot be defined and applied within the same package compilation. Put reusable handlers and their generated traits in a separate package. The target may be a generic struct or enum; the generated impl inherits its type parameters. Explicit owner arguments also specialize static inherent methods whose parameter and return types do not mention those arguments, for example `Record::[string]::type_name()`. `derive_output_add_predicate` and `derive_output_add_call_site_predicate` add the bounds required by generated methods.

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

The structured output API provides opaque `MetaAttribute`, `MetaType`, `MetaExpr`, `MetaPattern`, `MetaArm`, `MetaBlock`, `MetaParamList`, `MetaGenericList`, `MetaMethod`, and list handles. Constructors use the `meta_type_*`, `meta_expr_*`, `meta_pattern_*`, `meta_arm*`, `meta_block_*`, `meta_param_list_*`, and `meta_generic_list_*` families. `meta_method` creates a concrete method; `meta_method_generic` creates a method with explicit type parameters and bounds. A handler creates one trait impl with `derive_output_new` or `derive_output_new_call_site`, or one inherent impl with `derive_output_inherent`. It then adds predicates and methods and returns that output. `derive_output_add_method` preserves private visibility for inherent methods and normal trait visibility for trait implementations. `derive_output_add_public_method` exports an inherent method and rejects trait output.

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

The result is restricted to one trait or inherent `impl` for the annotated type. It cannot create types, traits, free functions, constants, modules, imports, extern declarations, attributes, associated types, or raw tokens. Associated constants and multiple impl blocks in one output remain unsupported. Generated method type parameters and trait bounds are supported. The generated impl is processed by ordinary name resolution, orphan and coherence checks, type checking, monomorphization, and backend lowering. Duplicate or invalid generated implementations are regular compiler diagnostics.

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

The adapters add no synchronization, global registry, reflection or implicit panic recovery. A panic propagates along the same goroutine, including through both conversion directions, and Go and GoML defers run during stack unwinding. It does not automatically become an error result. Recovery requires an explicit `std::panic::catch` boundary or a caller-owned Go adapter using defer/recover in that goroutine; neither can recover a panic in another goroutine or intercept process exit. `panic::resume` preserves the captured panic's diagnostic information using an internal runtime wrapper; a Go caller recovering after resume observes that wrapper rather than the original payload's dynamic type. Go may retain a converted callback after its registering GoML call returns, invoke it on another goroutine, or reenter GoML recursively. Captured state remains reachable through the Go function value. Concurrent mutation needs ordinary synchronization; conversion does not make a `Ref` thread-safe.

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

### Calling C libraries

`goml bind-c <CONFIG> [--compiler <COMPILER>] [--dry-run | --check]` generates an explicitly allowlisted C binding using Clang declarations and ABI information. The Go backend supports `cgo` (the default) and `dynamic` bindings. The dynamic backend uses GoML's own assembly and pthread runtime with `CGO_ENABLED=0`, without third-party FFI libraries. An existing module-root `go.mod` is required; `[native]` declares `go-module` and `c-bindings = "bindings.json"`, plus `cgo = "required"` for the cgo backend. A dynamic configuration selects `"backend": "dynamic"` and `"libraries": ["libsqlite3.so.0"]`; the driver supplies its bundled runtime module automatically. Project commands verify C inputs before compilation, and changed preprocessed header contents invalidate the native cache. Generated GoML and Go files plus their ownership manifest are checked before replacement. See [the configuration and lifetime contract](ffi/bind-c.md).

Supported mappings include opaque typed C handles with `null`, `is_null` and `same_as`, fixed-arity scalar functions, copied strings and byte buffers, checked buffer lengths, scalar/handle output parameters, and explicitly released output strings. Functions return `Result[T, c::Error]` for adapter failures and preserve C status codes as ordinary return values. Handle copying does not copy or retain C resources. Nullability, aliasing, parent lifetimes, thread safety and destruction remain C API contracts; these low-level bindings do not imply memory safety.

`std::c::CString::new(string)`, `from_bytes(Bytes)` and `from_raw(ffi::String)` reject interior NUL. `bytes()` preserves arbitrary bytes and `text()` returns `Result[string, utf8::Utf8Error]`; `as_raw()` exposes the bridge representation. `c::Error` has `InteriorNul` and `Native(string)` variants and supports `Debug`, `PartialEq`, `Eq` and `ToString`. `c::check_error(ffi::Error)` is the generated adapter's error check.

`const NAME: string = c::literal("example");` validates a C string literal with CTFE. The ordinary `CString` constructor still checks runtime values. Allowlisted integer C constants become typed GoML constants usable in `comptime`; their values come from Clang. CTFE cannot inspect headers or allocate C memory. Examples cover [a self-contained C API](../examples/ffi-bind-c/README.md), [SQLite](../examples/ffi-c-sqlite/README.md), and [LLVM](../examples/ffi-c-llvm/README.md).

Both backends require the host Go target and Clang for generation and verification. The dynamic backend currently requires Linux amd64, glibc 2.34+, Go 1.26.x and `CGO_ENABLED=0`; its application builds need no C compiler. It loads exported shared-library functions on first use and returns loader failures through `c::Error::Native`. Static/inline functions and mixed cgo/dynamic dependencies are unsupported in this backend. Neither backend supports callbacks, variadic functions, struct/union values, pointer dereference, C exports or cross compilation. The bindings use existing declarations and calls, without `extern "C"`, native-layout or unsafe syntax.

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

The built-in `PartialEq`, `Eq`, and `Hash` implementations for `Ref[T]` use reference identity and do not require `T` to implement those traits. Mutating the referenced value therefore does not change equality or hashing. Separate allocations have distinct identities even when `T` is `()`, an empty struct, or another zero-sized type. Copies of one reference keep its identity. `Ref[T]` does not implement `Default`, because implicit allocation and recursive default construction would be surprising.

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

Import `std::collections` for pure GoML map algorithms. `map_clone` makes independent entry storage with shallow keys and values: nested `Ref`, `Vec` and other shared objects remain shared. `map_copy(destination, source)` overwrites source keys while retaining destination-only entries; copying a map into itself is safe. There is no nil-map state in `HashMap`.

`map_equal` compares key membership and `PartialEq` values, distinguishing missing keys from present zero values. `map_equal_by` accepts maps with different value types and an `(A, B) -> bool` comparator. Comparisons can stop early; callbacks must not mutate either compared map. Keys use the existing `Hash + Eq` contract, including consistent hashing and reflexive equality, not Go's raw-map key rules. Non-reflexive values such as floating-point NaNs remain unequal under ordinary `PartialEq`.

`map_all`, `map_keys` and `map_values` return single-pass `FnIterator` snapshots captured when called, with unspecified order. Later entry mutations do not alter the snapshot, but referenced objects remain shared. Iterators retain O(n) snapshot storage; they are not live Go range iterators and do not synchronize concurrent access. `map_insert(destination, iterator)` and `map_collect(iterator)` consume any `Iterator` whose item is `(K, V)`, overwriting duplicate keys in input order. Collection always allocates a new map.

`map_delete_if(map, predicate)` invokes the `(K, V) -> bool` predicate once for each original snapshot entry and removes selected keys from the current map, returning the number actually removed. Keys inserted during callbacks are not visited. If a callback replaces a selected key, the replacement is removed; if it already removes that key, it is not counted again. Predicate order is unspecified. Concurrent mutations still require caller synchronization.

Equal-length maps with disjoint keys do not invoke the value comparator; a comparator that returns false stops comparison immediately. Snapshot deletion still visits original entries whose keys were removed by earlier callbacks. Insertion consumes through the iterator's first None and applies each yielded pair after the iterator callback returns, so that pair overwrites any callback write to the same key. Existing snapshot iterators can safely supply entries back into their originating map. These helpers do not impose an iterator item limit, catch callback panics, roll back earlier writes or provide synchronization; callers must bound untrusted streams and keep hash/equality behavior stable while keys are stored.

```goml
use std::collections;

fn main() -> () {
    let counts = HashMap::from_array([("a", 1), ("b", 2)]);
    let copy = collections::map_clone(counts);
    let keys = collections::map_keys(copy);
    copy.set("c", 3);
    for key in keys {
        println(key);
    }
    let _ = collections::map_delete_if(copy, |_: string, count: isize| count < 2);
    println(collections::map_equal(counts, copy));
}
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
use std::context;
use std::crypto;
use std::encoding::base32;
use std::encoding::base64;
use std::encoding::hex;
use std::error;
use std::env;
use std::ffi;
use std::fs;
use std::hash;
use std::hash::adler32;
use std::hash::crc32;
use std::hash::crc64;
use std::hash::fnv;
use std::io;
use std::iter;
use std::json;
use std::math;
use std::math::bits;
use std::net;
use std::net::tls;
use std::num;
use std::os::linux::syscall;
use std::path;
use std::path::slash;
use std::panic;
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
- `bytes::Bytes`, `bytes::FrozenBytes`, `bytes::Builder`, explicit byte copies, immutable snapshots, checked and zero-copy byte views, plus `bytes::endian::{Builder, Reader, Writer, Endian}` and checked integer and floating-point reads and writes
- `channel::Operation`, `Selection`, `select`, `try_select`, and `try_select_priority` for runtime-sized channel selection
- `cmp::Ordering`, `Ord`, `Reverse`, comparison helpers, and two-value minimum, maximum, and clamping operations. `Ordering` is a builtin type re-exported by `cmp`.
- `collections::Arena`, `BinaryHeap`, `BitSet`, `BTreeMap`, `BTreeSet`, `Deque`, `HashSet`, `IndexMap`, `IndexSet`, `IndexVec`, `Interner`, and `Stack`; hash-backed collections require `Hash + Eq`, while tree collections and heaps use `cmp::Ord`
- `collections::sort`, `stable_sort`, `binary_search`, `min`, `max`, and their comparator variants. The sorting, search, selection, and deduplication methods on `Vec[T]` are the canonical forms.
- `collections::search`, `is_sorted`, `lower_bound`, `upper_bound`, `equal_range` and comparator variants for monotone indexed search, sortedness checks and duplicate ranges
- `collections::slice_*` view algorithms and `vec_replace`, `vec_insert_slice`, `vec_delete_range`, `vec_delete_if`, `vec_grow`, `vec_append_iter` for checked range edits and iterator composition
- `collections::{map_clone, map_copy, map_equal, map_equal_by, map_delete_if, map_all, map_keys, map_values, map_insert, map_collect}` for shallow map copies, membership-aware comparisons, snapshot iteration and collection
- `context::{Context, CancelHandle, Deadline}` and scoped `with_cancel`, `with_timeout`, and `with_deadline`
- `crypto::hash` one-shot SHA-256 and `crypto::rand` operating-system random bytes
- `encoding::hex` lowercase and uppercase hexadecimal encoding plus checked decoding
- `encoding::base32` RFC 4648 standard and extended-hex encoding with padded and unpadded variants and checked canonical decoding
- `encoding::base64` RFC 4648 standard and URL-safe encoding with padded and unpadded variants
- `error::Error`, `ErrorKind`, `Details`, and stable error-kind code conversion
- `env::args`, current-directory and executable queries, and environment-variable reads
- `ffi::String`, `Rune`, `Ptr`, `Error`, `Func`, `RawSlice`, `RawMap`, and explicit Go boundary adapters
- `fs::read_file_structured`, `write_file_structured`, structured byte I/O, directory operations, path inspection, and `sha256_file`
- `hash::Hasher`, incremental Adler32, CRC32, CRC64 and FNV-1/FNV-1a checksum implementations
- `io::{Read, Write, BufRead, Close}`, `Cursor`, `Take`, `BufReader`, `BufWriter`, `copy`, `stdin`, `stdout`, `stderr`, and the existing standard-stream functions
- `iter::empty`, `once`, `from_fn`, iterator adapters, and single-pass consumers
- `json::Value`, `parse`, `encode`, serde `Serialize` and `Deserialize` re-exports, `to_value`, `from_value`, `try_to_string`, `from_string`, `field`, and typed `as_*` accessors
- `math` f32/f64 elementary functions, IEEE 754 classification, and the `E`, `PI`, `TAU`, `SQRT_2`, `LN_2`, and `LN_10` constants
- `math::bits` fixed-width bit counting, reversal, rotation and checked double-width arithmetic
- `net::{IpAddr, ScopedIpAddr, IpPrefix, SocketAddr}` for IP values, zones and prefixes; `net::{TcpListener, TcpStream, UdpSocket, WaitOptions}` for Linux amd64 syscall-backed networking, shared epoll readiness, timeouts, and cancellation
- `net::resolve`, `resolve_with`, `TcpStream::connect_host`, and `net::tls` verified TLS clients
- `num` structured parsing, scalar conversion traits, explicit rounding modes, plus checked and saturating `i64` arithmetic
- `os::linux::syscall` Linux amd64 calls by number, six machine-word arguments, scoped mutable byte-buffer graphs, raw return values, and named errno
- `os::linux::abi` checked Linux amd64 record codecs and `Iovecs`/`Message` nested-buffer builders
- `os::linux::fd` owned Linux descriptors, scalar/vector/positional I/O, metadata, byte paths, directory-relative operations, locks, and anonymous memory files
- `os::linux::process` IDs, signals, pidfds, child creation/waiting, resource usage, limits, and priorities
- `os::linux::memory` anonymous/file mappings, checked copied access, protection, sync/advice, page locking, and residency
- `os::linux::ipc` pipes, Unix socket pairs, descriptor passing, eventfd, timerfd, poll, and epoll
- `path::join`, `clean`, `is_absolute`, component inspection, and `absolute_structured`
- `path::slash::{clean, join, split, base, dir, extension, is_absolute, matches}` and reusable `Pattern` for pure slash-separated logical paths
- `resource::{Scope, ScopeError, scope, with_cleanup, finish}` for explicit cleanup and combined errors
- `process::Command`, structured whole-process execution, `ExitStatus`, `Output`, `exit`, and `look_path_structured`
- `rand::ALGORITHM`, `next_u64`, deterministic byte generation, integer ranges, and shuffle with an explicit seed
- `serde::Value`, `Serializer`, `Deserializer`, `Serialize`, `Deserialize`, `value_serializer`, `value_deserializer`, `to_value`, and `from_value`
- `task::Scope`, `Task[T]`, `CancelToken`, `WaitResult[T]`, `scope`, and `try_scope`
- `testing::fail`, boolean/equality assertions, and `Option`/`Result` shape assertions
- `text::StringBuilder`, `LineIndex`, `LineColumn`, and `PositionEncoding`; the byte-offset search, character iteration and slicing, trimming, splitting, replacement, joining, repetition, and explicit ASCII operations are string methods
- `toml::Value`, `parse`, `encode`, serde `Serialize` and `Deserialize` re-exports, `to_value`, `from_value`, `to_string`, and `from_string`
- `time::Duration`, `Instant`, `SystemTime`, `sleep`, and `sleep_with`
- `utf8::validate`, `decode`, `decode_slice`, `encode`, `encode_into`, `encode_to`, `encode_char`, `encode_char_into`, `encode_char_to`, `encoded_len`, and `Utf8Error`
- `utf8::{decode_first, decode_rune, decode_last_rune, decode_lossy, full_rune, rune_count, rune_start, valid_scalar, Decoder}` for pure scalar decoding and explicit invalid-byte policies
- `utf16::decode`, `decode_bytes`, `encode`, `encode_bytes`, and `Utf16Error`
- `unicode::VERSION`, scalar properties, and Unicode case conversion using Unicode 15.0.0 tables

### Structured error foundation

`std::error` provides the common protocol used by structured standard-library errors. `Error` is a marker trait requiring `Debug` and `ToString`. `ErrorKind` defines stable, message-independent categories including missing files, permission failures, invalid input or data, timeouts, interruption, short I/O, broken pipes, unsupported operations, and `Other`.

`Error` has a blanket implementation for every value implementing `Debug` and `ToString`, so domain errors participate in the common protocol directly.

`Report[E]` preserves typed leaf errors while adding an immutable tree of context and aggregation. `Report::new(value)` creates a leaf; `.context(message)` returns a new parent without changing the original. `Report::join(Slice[Report[E]])` snapshots the supplied sequence, preserves order and duplicates, and returns None for an empty sequence. A singleton remains an aggregate with one cause. Use a domain enum for heterogeneous leaf variants; this API does not erase types or perform reflective downcasts.

`.value()` returns Some only at a leaf, `.message()` only at a context node, and `.causes()` returns a fresh vector of immediate children. Mutating either the join input or the returned child vector does not alter the report tree. Leaf values retain their ordinary copy/shared-handle semantics: reports do not deep-copy mutable application errors.

`.find(predicate)` searches leaves depth-first, left-to-right and stops at the first match. `.find_map(select)` also permits typed projection to another value; `.contains(target)` uses leaf PartialEq rather than comparing display strings. Context nodes and aggregates are not implicit matches, and arbitrary errors embedded inside a leaf are not recursively unwrapped. These explicit operations provide typed cause inspection without changing the existing Error marker trait or adding Go-style reflection-based Is/As behavior.

Reports compose with `Result::map_err` and `?` without losing domain variants or standard I/O error fields. A domain enum can hold `io::Error` alongside application validation failures; `find_map` can project the I/O variant back to its original typed error, including its kind and OS code. Reusing one subtree in multiple aggregate positions visits it once per occurrence, in order; there is no identity-based deduplication. Callback work and output size are not budgeted by Report, and callbacks may observe mutable leaf handles, so callers must bound externally controlled error trees and avoid concurrent payload mutation.

Report implements ToString and Debug when E implements ToString, and therefore participates in the Error protocol. Display uses `message + ": " + cause` for context and newline-separated children for aggregation, including empty messages and duplicate causes. Display is not a serialization format. Search and display use iterative stacks, not recursive calls; work is proportional to visited nodes plus rendered bytes, with allocations for traversal and output. No arbitrary depth limit is imposed. This adds no syntax or runtime error interface.

```goml
use std::error;

fn with_context[T](result: Result[T, error::Details]) -> Result[T, error::Report[error::Details]] {
    result.map_err(|cause: error::Details| error::Report::new(cause).context("load settings"))
}

fn missing_file(report: error::Report[error::Details]) -> bool {
    report.find(|cause: error::Details| {
        error::kind_from_code(cause.kind_code()) == error::ErrorKind::NotFound
    }).is_some()
}
```

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

The scalar decoder rejects overlong encodings, surrogate code points, values above U+10FFFF, stray continuation bytes and invalid lead bytes. `validate` and incremental decoding share this check. An already-invalid prefix is an error even when fewer than its nominal width's bytes are present: for example `E0 80` is invalid (`Some(1)`), while `E1 80` is incomplete (`None`). Invalid-byte recovery consumes one byte; error offsets point to the start of the first failed scalar. This corrects the former classification of some known-invalid short prefixes as incomplete.

`decode_first(Slice[byte])` returns a checked `(char, width)` for the first scalar, ignoring subsequent bytes; empty or incomplete input returns `Utf8Error` with `None`. `decode_rune` and `decode_last_rune` instead implement replacement decoding: empty input returns `(REPLACEMENT, 0)`, invalid or incomplete nonempty input returns `(REPLACEMENT, 1)`, and valid scalars return their encoded width. `REPLACEMENT` is U+FFFD and `MAX_WIDTH` is four. A real encoded U+FFFD has width three, distinguishable from invalid-byte replacement. `full_rune` is true when the prefix is either complete or already known invalid, not a validity check. `rune_count` counts malformed bytes individually; `decode_lossy` replaces each one with U+FFFD, without coalescing adjacent invalid bytes. These APIs accept arbitrary binary input and never normalize Unicode.

`rune_start(byte)` tests whether a byte is not a continuation byte; it does not guarantee a valid lead byte. `valid_scalar(u32)` rejects surrogates and out-of-range values. GoML `char` already guarantees a scalar, so `encode_char(char) -> Vec[byte]` has no invalid-character fallback. `encode_char_into` checks the entire destination range before writing and returns its width or `bytes::BoundsError` without partial writes; `encode_char_to` appends the encoded bytes to a byte builder. The source implementation performs scalar encoding directly and retains the existing string-conversion boundary.

`Decoder::new()` creates a strict incremental decoder retaining at most four bytes. `feed(byte)` returns `Result[Option[char], DecoderError]`: `None` means a valid prefix needs more bytes, and `Some` delivers the completed scalar exactly once. Each byte call preserves already-delivered output even if a later byte fails. `consumed()` counts accepted feed bytes, including the byte proving invalidity. `DecoderError::InvalidUtf8` contains a `Utf8Error` with an absolute stream offset; malformed-input errors are sticky until `reset()`, and later feeds do not consume input. `finish()` reports incomplete final sequences, otherwise closes the decoder; successful repeated finishes are idempotent and later feeds return `Finished`. `OffsetOverflow` rejects a byte that would exceed the Linux amd64 isize offset range without consuming it. `reset()` clears offsets, pending bytes, failures and the finished state. Decoder copies are shared mutable handles, not independent snapshots, and require external synchronization for concurrent use.

```goml
use std::utf8;

fn decode_stream(input: Slice[byte]) -> Result[Vec[char], utf8::DecoderError] {
    let decoder = utf8::Decoder::new();
    let output = Vec::new();
    for value in input {
        if let Option::Some(character) = decoder.feed(value)? {
            output.push(character);
        }
    }
    decoder.finish()?;
    Result::Ok(output)
}
```

Importing `utf8::BytesUtf8` adds `Bytes::to_string_utf8`, which returns `Result[string, Utf8Error]`.

`decode_slice` accepts a read-only byte view. `encode_into` writes into a `MutSlice[byte]` at a checked offset and reports `bytes::BoundsError` without a partial write; `encode_to` appends to the lightweight `bytes::Builder`.

### Byte buffers and endian access

Byte search operates on arbitrary `Slice[byte]` values without UTF-8 decoding.
`bytes::find` and `rfind` return the first/last byte offset as `Option[isize]`;
`contains`, `starts_with` and `ends_with` return booleans. An empty pattern matches
at offset zero for `find` and at the input length for `rfind`. Last-match search
includes overlapping occurrences. `find_byte` and `rfind_byte` search one byte.

`bytes::Finder::new(pattern)` snapshots the pattern and builds a prefix fallback
table. Its `find`, `rfind` and `contains` methods reuse that immutable state across
inputs. Construction takes O(pattern length) time/space; each search takes
O(input length) time and O(1) additional space. Later pattern-buffer mutation
does not change the finder. The free search functions construct a fresh finder.

`bytes::cut(input, separator)` returns `Some((before, after))` around the first
match, excluding the separator, or `None` if absent. `cut_prefix` and `cut_suffix`
return the remaining view only when the corresponding edge matches. Returned
slices share the input's backing storage; mutation through an existing mutable
alias is visible. Copy a returned view with `to_vec()` when isolation is needed.
An empty separator cuts into an empty prefix and the whole input; cutting an
empty prefix/suffix returns the whole input. These APIs add no syntax forms.

`bytes::count(input, pattern) -> Result[isize, TransformError]` counts
nonoverlapping matches. Empty patterns match before the first UTF-8 scalar and
after every scalar; each malformed or truncated byte advances by one. This
matches replacement decoding, not strict UTF-8 validation, and preserves all
original bytes.

`split(input, separator, max_parts)` and `split_after` return checked vectors of
shared input views; the latter retains each matched separator in the preceding
part. `split_n` and `split_after_n` insert a `count` argument before `max_parts`:
negative counts mean all parts, zero means none, and positive counts produce at
most that many parts with an unsplit final remainder. The independent nonnegative
`max_parts` is a resource limit: exceeding it returns an error, not truncation.
Empty separators split at the replacement-decoding boundaries described above,
without leading/trailing empty parts. Empty input and empty separator yield zero
parts; nonempty separators retain empty boundary and adjacent parts.

`split_iter(input, separator)` and `split_after_iter` return single-pass
`FnIterator[Slice[byte]]` values with the same unlimited-split semantics. They
snapshot and compile the separator at construction, then scan only as requested
by `next()`, without allocating a result vector. `lines_iter(input)` yields
newline-delimited views including each terminating `\n`, with a final unterminated
line when present. It emits no extra empty part after a trailing newline and no
parts for empty input; `\r` is preserved, not normalized.

Iterator copies share traversal state; construct another iterator to restart.
After exhaustion, all further calls return `None`. Returned views and pending
input retain shared backing storage, so later input mutation is visible in prior
views and affects future reads; separator mutation does not affect the iterator.
Do not call `next()` concurrently or mutate input concurrently with a call.
Early stopping avoids scanning the remaining input, but the iterator keeps the
input storage reachable until it is released. It uses O(separator length) setup
storage and constant additional traversal state; callers collecting results must
apply their own bound or use the bounded vector APIs. Eager split variants reuse
the same traversal logic, including count-limited final remainders.

`replace(input, old, replacement, max_bytes)` replaces all nonoverlapping matches.
`replace_n` adds a maximum replacement count before `max_bytes`; negative counts
replace all, zero makes an unchanged copy. Empty old patterns insert before the
first scalar and after successive replacement-decoding boundaries. Search,
count, split and replacement remain linear in input/pattern/output size; repeated
searches reuse a compiled pattern. These raw split/replacement algorithms need
no Unicode property tables; the package's Unicode-aware helpers below reuse
the standard Unicode data without introducing a text dependency.

`join(Vec[Slice[byte]], separator, max_bytes)` and
`repeat(input, count, max_bytes)` construct byte buffers. All three construction
families return `Result[Bytes, TransformError]`, check the final byte length
before allocating output storage, and return independent buffers even for a
single part or unchanged replacement. Negative repetition counts are errors;
empty input can be repeated any nonnegative number of times without looping.

`TransformError` distinguishes `InvalidLimit`, `InvalidCount`, `LimitExceeded`
and `LengthOverflow`. Negative limits are rejected even for empty results; exact
limits succeed. Failure never mutates inputs or exposes a partial result.
Limits cover output bytes or part count, not caller-owned inputs, compiled search
tables or allocator overhead. They do not make global allocation failure
recoverable. As with other shared slices, callers must not concurrently mutate
inputs while an operation is reading them. These functions use existing syntax.

The Unicode-aware byte helpers decode valid UTF-8 scalars and treat each malformed
or truncated byte as U+FFFD, without coalescing adjacent invalid bytes. `find_char`
and `rfind_char` search that scalar stream; searching U+FFFD therefore also finds
invalid bytes, unlike a raw byte-pattern search. `find_any`/`rfind_any` accept a
string of scalar-set members. `find_by`/`rfind_by` invoke a predicate once per
visited scalar, forward/backward respectively, stopping at the first match.
All results are original byte offsets, not scalar indexes.

`fields_iter` and `fields_by_iter` yield nonempty shared views separated by
Unicode whitespace or a scalar predicate. Predicates run lazily and once per
visited scalar in input order. `fields`/`fields_by` collect with a nonnegative
`max_fields` result-count limit and return `TransformError` on invalid/exceeded
limits. Aliases of an iterator share progress, and exhaustion is permanent.
`trim_space`, `trim_by`, `trim_start_by`, `trim_end_by`, `trim_chars`,
`trim_start_chars` and `trim_end_chars` return shared edge-trimmed views. Character
sets classify scalars, not byte substrings; trimming U+FFFD can remove invalid
bytes as well as real encoded U+FFFD. These view operations preserve the exact
remaining bytes, including invalid UTF-8.

`map_chars(input, transform, max_bytes)` instead constructs an independent valid
UTF-8 buffer. Its transform returns `Option[char]`: `None` drops the decoded
scalar, and `Some` emits one valid scalar. Each visited scalar invokes the
callback once, and each emitted encoding is checked against the byte limit
before appending. Thus identity mapping replaces invalid input bytes with the
three-byte UTF-8 encoding of U+FFFD. `to_case(input, unicode::Case, max_bytes)`
and `case_with(input, special, kind, max_bytes)` use the same checked mapping
path for normal or special simple casing. Negative limits fail before callbacks;
other failures return no partial buffer but do not roll back callback effects.
Limits exclude source storage and allocation overhead.

`bytes::replace_invalid_utf8(input, replacement, max_bytes)` replaces each maximal
run of malformed UTF-8 bytes once, preserving valid scalars (including encoded
U+FFFD) exactly. Empty replacement deletes such runs. Replacement is arbitrary
bytes and is not decoded again, so valid UTF-8 output requires a valid UTF-8
replacement. Unlike `map_chars`, adjacent malformed bytes share one replacement.
The function preflights the complete output length before allocation, returns
`TransformError` for negative limits, overflow or excessive output, and always
returns independent storage, including when input already contains valid UTF-8.

`equal_fold(left, right)` compares replacement-decoded scalars by Unicode simple
fold cycles, not full multi-scalar folding or normalization. Different malformed
bytes can compare equal because each decodes as U+FFFD. Raw `find`/`rfind`, byte
search, prefix/suffix checks and split/replace remain byte-pattern operations.
Inputs must not be concurrently mutated while a call is reading them; shared
views reflect subsequent mutations. These APIs introduce no new syntax forms.

```goml
use std::bytes;

fn byte_search_example() -> Option[isize] {
    let finder = bytes::Finder::new(bytes::Bytes::from_string("aba").as_slice());
    finder.rfind(bytes::Bytes::from_string("ababa").as_slice())
}

fn invalid_byte_search_example() -> Option[isize] {
    bytes::find_char(b"\xFF".as_slice(), '�')
}

fn repair_byte_runs_example() -> Result[bytes::Bytes, bytes::TransformError] {
    bytes::replace_invalid_utf8(b"a\xFF\x80b".as_slice(), b"?".as_slice(), 3)
}
```

`std::bytes` uses `Slice[byte]` and `MutSlice[byte]` for borrowed views. `Bytes::as_slice` and `as_mut_slice` are zero-copy, while `slice_checked` and `slice_mut_checked` validate a subrange. `bytes::Builder` is a lightweight byte accumulator and returns `Bytes`. `Bytes::from_vec` and `to_vec` share mutable storage. Use `from_vec_copy`, `from_slice_copy`, `to_vec_copy`, or `copy` for independent mutable buffers. `Builder::finish_copy` and `to_vec_copy` detach a snapshot from the builder. These copies do not synchronize concurrent mutations of the source.

`Bytes::freeze` creates a `FrozenBytes` snapshot with private immutable storage. `FrozenBytes::from_vec_copy`, `from_slice_copy`, and `from_string` also create snapshots. `len`, `is_empty`, and `get_checked` inspect them; `slice_checked(start, end)` returns an immutable view sharing that snapshot, including empty ranges. Invalid indexes and ranges return `None`. `to_vec_copy` and `thaw` return independent mutable copies. `FrozenBytes` implements content-based `PartialEq` and `Eq`.

`std::bytes::endian` is an opt-in package for binary formats. Its `Builder` grows while writing typed values; `Reader` advances over a read-only view; `Writer` advances over a fixed mutable view and returns `endian::BoundsError` rather than partially writing past the end. The top-level `read_u16/u32/u64`, `read_i16/i32/i64`, `read_f32/f64` and matching `write_*` functions take an explicit `Endian::Little` or `Endian::Big`. One-byte operations omit endianness. Every operation validates the complete range before reading or writing.

The stateful `Reader`, `Writer`, and `Builder` provide typed methods for `bool`, `u8/u16/u32/u64`, `i8/i16/i32/i64`, and `f32/f64`. Multibyte methods take explicit endianness. Reader and Writer methods advance only after success, with copies sharing the cursor; failed bounds checks leave both position and destination unchanged. Floating-point methods preserve IEEE bit patterns, including signed zero and NaN payloads. Top-level functions take an explicit offset and do not advance any cursor.

Boolean `read_bool` and `write_bool` operations occupy one byte without an endianness argument: zero decodes to false and every nonzero byte to true; writing produces only zero or one. These operations are available as memory functions, cursor/builder methods, and `std::bytes::endian::stream::{read_bool, write_bool}` with the same EOF, error-propagation and partial-progress contracts as other fixed-width stream functions. For example, `builder.write_bool(true)` appends one byte and `reader.read_i16(endian::Endian::Big)` reads a signed two-byte value. These ordinary methods introduce no syntax or implicit struct layout.

`Reader::read_with[T](width, decode)` composes an explicit fixed-width record or array layout. The callback `(Reader) -> Result[T, BoundsError]` receives an independent cursor over exactly that block. The complete outer range is checked before invoking it. Success advances the outer cursor by the declared width, skipping any unread padding; failure leaves the outer cursor unchanged. Nested errors retain offsets and lengths relative to the nested view where they arose, while the initial range error describes the outer view.

`Writer::write_with(width, encode)` accepts `(Writer) -> Result[(), BoundsError]` and first validates the outer range, then invokes the callback on a separate zero-filled buffer of exactly `width` bytes. Success copies the whole block and advances the outer cursor; unwritten bytes are zero padding. Failure discards the scratch buffer without changing the outer position or destination. `Builder::write_with(width, limit, encode)` uses the same staging contract and appends only on success; `limit` bounds the total resulting builder length, including existing bytes. Invalid widths or limits fail before allocation or callback invocation. Zero-width blocks invoke the callback with an empty view and do not advance the cursor. Retained scratch writers never alias the committed destination.

These methods use ordinary callbacks, structs and loops, not reflection, automatic field ordering, a derive macro or native struct layout. Callbacks may nest blocks and capture an explicit endian order. They must not reenter or mutate the outer cursor, builder or its storage through captured aliases; unrelated callback side effects, concurrent mutation and panics are not rolled back. Reading uses a bounded borrowed input view, not an immutable snapshot; writing allocates one scratch buffer per active nested block. Bound untrusted record widths and array counts before constructing output storage; the builder limit controls bytes, not callback execution time.

```goml
use std::bytes::endian;

fn read_header(reader: endian::Reader) -> Result[(u16, bool), endian::BoundsError] {
    reader.read_with(4, |block: endian::Reader| {
        let version = block.read_u16(endian::Endian::Big)?;
        let enabled = block.read_bool()?;
        Result::Ok((version, enabled))
    })
}

fn write_header(writer: endian::Writer, version: u16, enabled: bool) -> Result[(), endian::BoundsError] {
    writer.write_with(4, |block: endian::Writer| {
        block.write_u16(version, endian::Endian::Big)?;
        block.write_bool(enabled)
    })
}
```

`bytes::endian` also provides unsigned base-128 varints and signed ZigZag varints. `uvarint_len` and `varint_len` report encoded sizes, at most `MAX_VARINT_LEN64` (10). `read_uvarint(input, offset)` and `read_varint(input, offset)` return `(value, bytes_consumed)` in a `Result`; the count is relative to the supplied offset. They leave the input unchanged and accept nonminimal encodings that still fit in 64 bits. `VarintError` distinguishes `InvalidOffset`, `Truncated`, and `Overflow`, carrying an absolute byte offset. A tenth byte greater than 1 is immediately an overflow, including a continuation byte; this differs from Go's slice decoder reporting some ten-byte continuation-only inputs as incomplete.

`write_uvarint(output, offset, value)` and `write_varint` validate the complete output range before writing and return the byte count or `BoundsError`; failure does not partially mutate the destination. `append_uvarint(builder, value)` and `append_varint` append to `bytes::Builder`. These free functions do not advance an endian `Reader` or `Writer`.

`endian::Reader::read_uvarint()` and `read_varint()` decode from the current position and advance it only after success, returning the value. Truncated and overflowing input leave the position unchanged; error offsets are absolute within the reader's original view. `endian::Writer::write_uvarint(value)` and `write_varint(value)` return `Result[(), BoundsError]`, advancing by the encoded length only on success. Insufficient space changes neither position nor output bytes. Copies of a reader/writer share its position, consistent with existing fixed-width methods. For example, `writer.write_uvarint(300)` writes two bytes and advances two positions. These ordinary inherent methods introduce no grammar. The endian value package retains its pure memory-only dependency graph; streaming I/O belongs to its separately imported child package.

Import `std::bytes::endian::stream` for `read_uvarint(reader: R) -> Result[Option[u64], stream::ReadError]` and `read_varint(reader: R) -> Result[Option[i64], stream::ReadError]`, with `R: io::Read`. A clean EOF before the first byte returns `None`; after any prefix it returns `ReadError::Decode(VarintError::Truncated(count))`. Overflow and noncanonical-but-valid encodings follow the existing byte decoder. Reads request one byte at a time, return every error including `Interrupted`, validate returned counts and consume at most ten bytes, never the next encoded value. Decode error offsets start at zero for each operation. `ReadError::Io(error, count)` preserves the original I/O error and the number of bytes returned by preceding successful reads; it cannot account for unreported side effects inside a failed read. Unlike memory cursors, streams cannot roll back consumed prefixes after errors.

`stream::write_uvarint(writer, value)` and `write_varint(writer, value)` return `Result[(), io::Error]` for `W: io::Write`. They encode using a bounded ten-byte scratch buffer and call the writer's `write_all`, respecting adapter-specific no-retry behavior. Failure can follow partial writes and does not report a progress count. Neither reading nor writing flushes or closes the underlying stream. For example, `stream::write_uvarint(io::Discard::new(), 300)` emits a two-byte encoding to a discard writer. The stream package depends on endian and I/O, not vice versa; existing generic functions and error enums express the APIs without grammar changes.

The same package exposes `read_u8/i8(reader)`, `read_u16/i16/u32/i32/u64/i64/f32/f64(reader, order)` and corresponding `write_*` functions taking `(writer, value)` or `(writer, value, order)`. Here `order` is `endian::Endian::{Little, Big}`. Each read returns `Result[Option[T], stream::ReadError]` for its scalar type; each write returns `Result[(), io::Error]`. Integer and IEEE floating-point bit encodings reuse the memory codecs, including signed zero and NaN payloads, without native algorithm calls or reflection.

Fixed-width reads fill at most the remaining field width using one buffer of one, two, four or eight bytes, return every error including `Interrupted` and validate each reported count. EOF before any byte returns `None`; EOF after a prefix returns `ReadError::Io(error, count)` with `UnexpectedEof` and the successfully read prefix length. Other I/O errors keep their original details with that same progress convention. A successful read consumes exactly the field width, leaving subsequent fields untouched. Fixed-width writes fully encode before calling `write_all`, retain the destination's retry policy and never flush or close it. For example, `stream::read_u32(reader, endian::Endian::Big)` reads an optional four-byte big-endian field. These are ordinary generic functions, not a new binary-layout language.

For explicit aggregate layouts, `stream::read_with(reader, width, limit, decode)` and `stream::write_with(writer, width, limit, encode)` accept the same bounded endian Reader/Writer callbacks as memory blocks. The nonnegative per-record byte limit is checked before allocation, callback invocation or I/O; invalid parameters return `InvalidInput` (with zero read progress). Each operation stages exactly `width` bytes, with the same zero padding and trailing-field rules as memory blocks. Reads fill the block before decoding, stop on every read error and never consume the following record; clean EOF returns `None`, a partial block returns `UnexpectedEof` with confirmed progress, and a callback bounds error becomes `InvalidData` with progress equal to the whole consumed block. This error includes the callback bounds diagnostic, not a rollback of the stream. A zero-width read invokes the decoder and returns `Some(value)` without probing EOF; do not use it as an EOF-driven loop condition.

Streaming writes invoke the encoder before calling `write_all`; callback bounds errors become `InvalidData` without any underlying writes. Later I/O failure can leave partially written records. Zero-width writes still invoke the encoder and, on success, the writer's `write_all` with an empty slice; custom overrides retain their own semantics. Both functions preserve existing no-flush/no-close behavior and do not synchronize or undo callback side effects. For example, a sixteen-byte record can reuse `read_record` through `stream::read_with(source, 16, 16, |block: endian::Reader| read_record(block))`. Fixed array element counts, field order, padding, mixed byte orders and semantic validation remain explicit caller schema decisions; there is no reflection or automatic serde/native-layout mapping.

```goml
use std::bytes;
use std::bytes::endian;

fn encoded_integer() -> bytes::Bytes {
    let output = bytes::Builder::new();
    endian::append_varint(output, -123);
    output.finish()
}
```

### UTF-16 conversion

`std::utf16` converts between strings and `Slice[u16]`, or between strings and endian-tagged byte slices. Decoding rejects lone surrogates and odd byte lengths with an indexed `Utf16Error`; encoding emits surrogate pairs for non-BMP scalar values.

### Text and source positions

Text search indices are UTF-8 byte offsets, matching the indices accepted by the built-in string APIs. `starts_with_at` returns false for out-of-range and non-character-boundary offsets. `rfind` returns the last matching byte offset. Trimming recognizes ASCII whitespace. Splitting on an empty separator returns the original string as one item, while `split_once` with an empty separator returns `Option::None`.

`text::LineIndex` precomputes line starts and non-ASCII scalar positions for repeated source-position conversion. `offset_to_line_column` and `line_column_to_offset` support UTF-8, UTF-16, and UTF-32 columns, treat CRLF as one line ending, and clamp unchecked out-of-range positions. `line_column_to_offset_checked` rejects positions inside an encoded scalar or outside the source.

`text::find`, `text::rfind`, and `text::find_bytes` return byte offsets; the string methods `find`, `rfind`, and `find_bytes` are the canonical forms. `text::char_indices` yields byte offsets paired with Unicode scalar values, `text::char_count` counts scalar values, and `text::slice_chars` uses scalar-value indexes and returns `None` for an invalid range.

`text::Finder::new(pattern)` compiles an immutable UTF-8 pattern for repeated
`find`, `rfind` and `contains` calls. Prefix-table construction takes O(pattern
byte length) time and space; searching takes O(input byte length) time with
constant additional state, without copying the input into a byte vector. Offsets
are bytes, including for multibyte characters. Last-match searches include
overlap. Empty patterns match at zero for first search and the input byte length
for last search. The free `text::find`, `rfind` and `find_bytes` functions use this
implementation while retaining existing results. Compiler-owned string-method
fallbacks remain in place for bootstrap compatibility.

`find_char`/`rfind_char` search one scalar; `find_any`/`rfind_any` search for a
member of a scalar set, and `contains_any` tests membership. Duplicated set
characters have no special meaning; an empty set never matches. `find_by` and
`rfind_by` accept a scalar predicate, scan forward/backward respectively and
stop at the first matching scalar in that direction. Each visited scalar is
classified once. All searches return byte offsets as `Option[isize]`, never
scalar indexes, and accept only valid Unicode characters. Set-based searches
use a character-keyed hash map, not repeated substring search per input character.

`text::count(value, pattern)` counts nonoverlapping matches and returns
`Result[isize, bytes::TransformError]`. An empty pattern counts scalar boundaries
(scalar count plus one); overflow returns `LengthOverflow`. `cut` returns
`Option[(string, string)]` around the first separator; `cut_prefix` and `cut_suffix`
return an optional remainder. Empty edge patterns match, and `cut(value, "")`
returns `Some(("", value))`. This deliberately differs from the retained
`split_once(value, "")`, which returns `None`.

Existing public operations also compose without additional text-specific aliases:
`cmp::compare(left, right).to_isize()` compares strings lexicographically;
`find_char(value, character).is_some()` and `find_by(value, predicate).is_some()`
test scalar or predicate membership. String `starts_with`/`ends_with` test fixed
prefixes/suffixes. `cut_prefix(value, prefix).unwrap_or(value)` and the analogous
`cut_suffix` composition remove one matching fixed edge, leaving unmatched input
unchanged; they do not trim a character set.

For single-byte search, use
`bytes::find_byte(value.to_bytes().as_slice(), needle)` or `rfind_byte`.
This composition copies the UTF-8 bytes and can find an interior encoding byte:
searching `"é"` for `0xa9` returns byte offset 1, which is not a scalar boundary.
The string-pattern operation `find_bytes` is not a replacement for this arbitrary
byte search. Invalid UTF-8 repair likewise belongs at the bytes boundary:
`bytes::replace_invalid_utf8` followed by checked UTF-8 conversion. GoML strings
are already valid UTF-8; an invalid replacement byte sequence does not guarantee
that the repaired bytes can be converted to a string.

Empty patterns have intentionally distinct contracts across API families:

| Operation with an empty pattern | Result or policy |
| --- | --- |
| `find` / `rfind` | Offset zero / input byte length. |
| `count` | Scalar count plus one, including one for empty input. |
| `cut` / `cut_prefix` / `cut_suffix` | Empty prefix plus whole input / unchanged remainder / unchanged remainder. |
| Retained `split` / `split_iter` / `split_after_iter` | Whole input once, including empty input. |
| Checked splitting | Scalar decomposition, with zero parts for empty input; N variants apply their count limit. |
| Retained `split_once` | `None`. |
| Retained `replace` | Input unchanged. |
| Checked replacement | Insert at scalar boundaries, subject to replacement-count and output limits. |
| Compiled `Replacer` | Ordered-rule matching with its documented scalar-boundary empty-rule policy. |

`text::split_iter` and `split_after_iter` yield substrings without an eager result
vector; the latter retains each separator in the preceding part. They use one
compiled separator and scan only when requested. As with existing GoML string
splitting, an empty separator yields the whole input exactly once, even for empty
input; it does not use Go/byte-split scalar decomposition. Nonempty separators
retain adjacent and edge empty parts. `lines_iter` matches existing `lines`: omit
LF, strip one trailing CR per line (also from an unterminated final line), and
omit the extra empty line after a trailing LF. `lines_inclusive_iter` instead
matches Go's line sequence convention: preserve LF and CR, with no extra trailing
empty line. Both line iterators yield nothing for empty input.

These are single-pass `FnIterator[string]` values: aliases share progress,
exhaustion is permanent, and another constructor starts independent traversal.
They retain their immutable source string while alive. Collecting results may
allocate unbounded output; early stopping does not scan the remaining input.
No new grammar forms are introduced.

`text::fields_iter(value)` lazily yields nonempty fields separated by Unicode
15.0.0 whitespace. `fields_by_iter(value, separator)` instead classifies each
visited scalar with a predicate, exactly once in left-to-right order, and does
not call it until iteration begins. Leading, trailing and repeated separators
produce no empty fields. These iterators share progress when copied and remain
exhausted after returning `None`.

`fields(value, max_fields)` and `fields_by(value, separator, max_fields)` collect
the same fields with an explicit nonnegative result-count limit, returning
`Result[Vec[string], bytes::TransformError]`. Exact limits succeed; negative
limits fail before any predicate calls. An exceeded limit returns no partial
vector, but predicates may already have run and their side effects are not
rolled back. The limit bounds result count, not input scanning or allocator
overhead. An all-separator input succeeds with a zero limit.

`trim_start_by`, `trim_end_by` and `trim_by` remove edge scalars satisfying a
predicate. Left trimming scans forward; right trimming scans backward; the
combined form applies those operations in order and may classify a retained
scalar twice. `trim_chars(value, characters)` removes members of a scalar set,
not a substring pattern; `trim_start_chars` and `trim_end_chars` operate on just
one edge using the same set semantics. `trim_space` uses Unicode whitespace. Existing `trim`,
`trim_start` and `trim_end` keep their ASCII-only contract.

`equal_fold(left, right)` compares Unicode simple-fold equivalence scalar by
scalar, with no locale or normalization rules. It accepts Kelvin-sign/K and
Greek sigma variants, but does not equate `ß` with `ss`; use full `case_fold`
from `std::unicode` when multi-scalar folding is intended.

`join_checked(values, separator, max_bytes)`,
`repeat_checked(value, count, max_bytes)` and
`replace_checked(value, old, replacement, max_bytes)` construct valid UTF-8 text
with a nonnegative byte limit, returning `Result[string, bytes::TransformError]`.
They check the result length before allocating output storage. Exact limits
succeed, arithmetic overflow returns `LengthOverflow`, and negative repetition
counts return `InvalidCount`. Empty text can be repeated any nonnegative number
of times without looping. Existing unbounded `join`, `repeat` and `replace`
remain unchanged.

`replace_n_checked` inserts a maximum replacement count before `max_bytes`:
negative means all nonoverlapping matches, zero means none. Checked replacement
uses Go boundary-insertion semantics for empty old strings (before the first
scalar and after subsequent scalars), unlike the retained unbounded `replace`,
which leaves an empty old string unchanged. Matching uses one compiled Finder
for the sizing and construction passes, without allocating all match offsets.

`map_chars(value, transform, max_bytes)` calls `transform(char) -> Option[char]`
once per visited scalar, in input order. `None` deletes a scalar; `Some` emits
one valid scalar. It performs no callback sizing pass: each emitted scalar's
encoded byte length is checked immediately before appending. On error, partial
text is not returned, but prior callback side effects remain; a negative limit
fails before invoking the callback. Mapping may succeed with a zero limit when
all input scalars are removed.

`split_checked(value, separator, max_parts)` and `split_after_checked` collect
bounded substrings, retaining separators in preceding parts for the latter.
Their `_n_checked` forms insert a `count` before `max_parts`: negative counts
mean all parts, zero means none, and positive counts leave an unsplit final
remainder after at most that many parts. The separate nonnegative result limit
returns `TransformError` on overflow rather than truncating the result; negative
limits fail even with zero count. No partial vector is returned.

Checked splitting follows Go's empty-separator scalar decomposition: no empty
edge parts, zero parts for empty input, and a positive count limits scalar parts
with a final remainder. This is intentionally distinct from retained GoML
`split`/`split_iter`, which yield the entire input for an empty separator.
Nonempty separators retain adjacent and edge empty parts in both API families.
Limits count result parts, not bytes or allocator overhead. These APIs add no
new syntax forms.

`case_checked(value, unicode::Case, max_bytes)` maps simple upper/lower/title
casing. `case_with_checked(value, special, kind, max_bytes)` uses a
`unicode::SpecialCase` override with normal fallback. These are one-scalar maps,
not normalization, full folding or word-title algorithms. Limits count encoded
bytes, not characters, and do not guarantee recovery from allocator exhaustion.
All these APIs use existing function, enum and closure syntax.

```goml
use std::text;

fn find_second_character() -> Option[isize] {
    text::Finder::new("界").find("é界")
}

fn unicode_text_example() -> bool {
    text::equal_fold("K", "k") && text::trim_space("　x　") == "x"
}

fn bounded_text_example() -> bool {
    text::repeat_checked("界", 3, 9).is_ok()
        && text::repeat_checked("界", 3, 8).is_err()
}

fn scalar_split_example() -> bool {
    text::split_n_checked("é界🙂", "", 2, 2).is_ok()
        && text::find_any("é界", "界") == Some(2)
}
```

#### Explicit word-title policy

Simple title mapping converts every scalar; it is not word-title formatting.
A caller can express an explicit word-boundary policy with stateful `map_chars`.
The following bounded compatibility recipe matches Go's deprecated `strings.Title`
on valid UTF-8. It is not a Unicode word-segmentation recommendation: ASCII
punctuation forms boundaries, but most non-ASCII punctuation does not. Linguistic
and locale-aware word segmentation belongs in ecosystem text processing.

```goml
use std::bytes;
use std::text;
use std::unicode;

fn legacy_title_separator(value: char) -> bool {
    if value.to_u32() <= 0x7f {
        !((value >= '0' && value <= '9')
            || (value >= 'a' && value <= 'z')
            || (value >= 'A' && value <= 'Z')
            || value == '_')
    } else {
        !unicode::is_letter(value)
            && !unicode::is_digit(value)
            && unicode::is_whitespace(value)
    }
}

fn legacy_word_title_checked(
    value: string,
    max_bytes: isize,
) -> Result[string, bytes::TransformError] {
    let previous = Ref::new(' ');
    text::map_chars(
        value,
        |character: char| {
            let boundary = legacy_title_separator(previous.get());
            previous.set(character);
            Option::Some(
                if boundary {
                    unicode::to_titlecase(character)
                } else {
                    character
                },
            )
        },
        max_bytes,
    )
}
```

The previous **original** scalar determines whether to map the current one.
Thus `"hELLO world"` becomes `"HELLO World"`, without lowercasing word interiors,
while `"foo—bar"` becomes `"Foo—bar"`. Each call starts fresh state. Output-byte
limits and errors are those of `map_chars`; no partially constructed text is
returned. This recipe uses ordinary closures and references, not a new standard
function or grammar form.

### Typed field formatting

Import `std::text::format` to assemble bounded messages using explicit field
types. This is a pure library API: `f"{expression}"` still uses ToString with no
format specifications. There is no printf format-string parser, reflection,
heterogeneous Any argument list, pointer/type-name formatting or scanf protocol.
Existing numeric parsing and ecosystem scanners remain separate facilities.

`Formatter::new(max_bytes)` rejects a negative byte limit. Its `literal`, `text`,
`integer`, `unsigned`, `float32`, `float64`, `display` and `debug` methods append
one field and return `Result[(), Error]`. `len()` and `remaining()` report bytes.
`finish()` returns an independent string snapshot without closing or consuming
the formatter. Copies of Formatter share its private builder and budget; no
writable builder or slice is exposed. Concurrent use is not supported.

`Padding { width, alignment, fill }` measures the final rendered field in Unicode
scalars, not UTF-8 bytes, grapheme clusters or terminal columns. Alignment is
Left, Right or Center; an odd extra Center padding scalar goes on the right.
Width is a minimum and never truncates. Fill is one char, including a multibyte
scalar. `Padding::plain()` is zero width, right alignment and space fill.

`TextSpec { padding, max_chars, style }` optionally truncates input by Unicode
scalar count before rendering. TextStyle is Plain, Quoted or QuotedAscii; the
last two use the existing text quoting rules. Padding then counts the rendered
scalars, including quotes and escapes. `TextSpec::plain()` does not truncate or
quote. No normalization, grapheme segmentation or display-width table is used.

`IntSpec { padding, radix, uppercase, alternate, sign, min_digits }` accepts bases
2 through 36. Integer fields take i64, unsigned fields u64; narrower values can
be explicitly widened. `alternate` is valid only for bases 2, 8 and 16 and emits
0b, 0o or 0x (uppercase when requested). Sign is NegativeOnly, Always or Space.
The order is sign, prefix, zero-padded digits. `min_digits` is nonnegative; zero
still has one digit when the minimum is zero. Outer padding applies to the
whole field: fill `0` does not imply sign-aware zero padding. For example,
negative hexadecimal 15 with alternate and min_digits 4 is `-0x000f`, whereas
right outer zero padding to width 6 around `-15` produces `000-15`.
`IntSpec::plain()` uses decimal, NegativeOnly and no padding or prefix.

`FloatSpec { padding, sign, style }` reuses num's separate f32 and f64 formatters.
FloatStyle supports Shortest(FloatNotation), Fixed(decimal_places),
Scientific(decimal_places, uppercase), General(significant_digits, uppercase),
Binary and Hex(fractional_digits, uppercase). Fixed, Scientific and General
precisions must be nonnegative; General zero means one significant digit.
Hex accepts -1 for exact trimmed output or a nonnegative fractional precision.
The existing num rules determine rounding, exponent spelling, NaN, infinities
and negative zero. Existing leading signs are preserved: `+Inf` is not given
a second sign. Unsigned NaN can become `+NaN` or ` NaN` under Sign.
`FloatSpec::plain()` uses Shortest(General) and NegativeOnly.

`display[T: ToString](value, padding)` and `debug[T: Debug](value, padding)`
validate padding before invoking the respective conversion exactly once. These
adapters cannot limit allocation or time inside user conversions, and do not
catch their panics. Conversion can reenter the formatter through an alias;
its own successful writes and other side effects remain even if the outer
field fails or panics. The outer field checks the current remaining budget
after conversion, so reentrancy cannot reset or bypass that budget.

All library field methods validate their specifications and full output size
before modifying the destination. A returned error does not partially append
that field or change the builder's capacity; earlier successful fields remain.
Error distinguishes InvalidLimit, InvalidSpec(message), LengthOverflow,
LimitExceeded(total_limit) and ConversionFailure(message). These error values
support ToString, Debug, PartialEq and Eq. LimitExceeded reports the formatter's
total limit, not its current remainder. All output bytes count, including signs,
prefixes, quoting and multibyte fill. Length arithmetic is checked before
width- or precision-dependent buffers are allocated. Temporary fields and the
destination can each use O(max_bytes) storage; this is not a CPU, allocator or
process-memory quota. Input scanning and user conversion have separate costs.

`write_to[W: io::Write](writer, value)` writes an already completed string and
returns `Result[usize, io::TransferError]`, counting confirmed UTF-8 bytes. It
uses primitive write, validates counts, continues short successful writes and
reports WriteZero for no progress. Any error, including Interrupted, ends the
operation without retrying that call; its unconfirmed external side effects
are not replayed. Empty input makes no writer call. It never flushes or closes
the writer and does not invoke an overridden write_all method. Format into a
private Formatter before calling write_to when formatting failure must produce
no external output; writing itself is not atomic.

```goml
use std::text::format;

fn message(value: i64) -> Result[string, format::Error] {
    let output = format::Formatter::new(128)?;
    output.literal("value=")?;
    output.integer(value, format::IntSpec {
        radix: 16, alternate: true, min_digits: 4,
        ..format::IntSpec::plain()
    })?;
    Result::Ok(output.finish())
}
```

This library adds no grammar production. Protocol templating, contextual HTML
escaping and terminal layout remain ecosystem concerns. Compiler and driver
formatting consumers retain released fallbacks until stage0 advancement.

### ASCII and Unicode

`std::ascii` operates on `byte`. Classification and case conversion use only the 7-bit ASCII range, and bytes above `0x7f` remain unchanged. `escape_default` emits short escapes for tabs, carriage returns, newlines, quotes, and backslashes, preserves printable ASCII, and uses lowercase `\\xNN` escapes for other bytes.

`std::unicode` fixes its public data version to Unicode 15.0.0. Character predicates operate on one Unicode scalar value. `lowercase` and `uppercase` apply Unicode case mapping to a complete string. `case_fold` uses the checked-in table generated by `tools/generate_unicode_casefold.py`, supports multi-scalar folds, and is locale independent.

Classification and simple casing use checked-in immutable tables generated by `tools/generate_unicode_tables.go`; production lookup and conversion are pure GoML. Predicates include `is_control`, `is_letter`, `is_number`, `is_digit`, `is_mark`, `is_punctuation`, `is_symbol`, `is_whitespace`, `is_lowercase`, `is_uppercase`, `is_titlecase`, `is_graphic` and `is_printable`. Graphic characters include Unicode space separators; printable characters include only ASCII space in addition to letters, marks, numbers, punctuation and symbols.

`category`, `script`, `property`, `fold_category` and `fold_script` take case-sensitive Unicode/Go table names and return `Option[RangeTable]`; their corresponding `*_names()` functions return fresh, sorted name vectors. Fold tables contain the additional code points needed for simple-fold closure of a category or script. `RangeTable::contains(char)` tests a scalar; `contains_codepoint(u32)` also allows inspecting surrogate categories and returns false above U+10FFFF. `in_any(char, Vec[RangeTable])` tests their union.

`RangeTable::new(Vec[Range])` copies checked ranges into an immutable table. Public `Range` fields are `low`, `high` and `stride`, all `u32`. Bounds are inclusive, membership is `(value - low) % stride == 0`, and the high bound need not itself be a member. Inputs must be sorted, nonoverlapping, within U+0000..U+10FFFF, and have nonzero strides. Failures return indexed `TableError::InvalidRange` or `UnsortedRange`. `ranges()` returns an independent copy; neither subsequent input mutation nor returned-range mutation changes the table.

`to_case(char, Case)` accepts `Upper`, `Lower` or `Title`; `to_uppercase`, `to_lowercase` and `to_titlecase` are shortcuts. `simple_fold(char)` advances the scalar's simple-fold equivalence cycle, unlike the multi-scalar `case_fold(string)`. String `uppercase` and `lowercase` use one-scalar mappings, so uppercase `ß` remains `ß`. They do not perform normalization, contextual casing or grapheme segmentation.

`SpecialCase::new(Vec[CaseMapping])` accepts strictly increasing, distinct `from` characters and explicit `upper`, `lower`, `title` mappings, returning `TableError::UnsortedMapping` otherwise. It snapshots its input and falls back to ordinary Unicode casing for unlisted characters. Its `to_case` method and string `lowercase_with`/`uppercase_with` use those overrides; `turkish_case()` and `azeri_case()` supply standard dotted/dotless-I mappings.

```goml
use std::unicode;

fn unicode_example() -> bool {
    let letters = unicode::category("L");
    unicode::uppercase_with("iı", unicode::turkish_case()) == "İI"
        && match letters {
            Some(table) => table.contains('界'),
            None => false,
        }
}
```

These APIs use existing function, enum and struct syntax; they add no grammar forms.

### Mathematics

`std::math` delegates elementary operations to Go's `math` package and follows its IEEE 754 special-value behavior. The f32 forms calculate through f64 and round the result back to f32. Results therefore use the target Go toolchain's correctly rounded conversions but do not promise bit-for-bit equality across different operating systems or processor implementations for every transcendental function.

### Fixed-width bit operations

`std::math::bits` implements portable integer algorithms in GoML without Go `math/bits` bindings. `ones_count8/16/32/64`, `leading_zeros8/16/32/64`, `trailing_zeros8/16/32/64`, and `len8/16/32/64` return `isize` counts. Zero has a population/bit length of zero and a leading/trailing zero count equal to its width. `reverse8/16/32/64` reverses bits; `reverse_bytes16/32/64` reverses byte order. `rotate_left8/16/32/64` takes an `isize` count reduced modulo the word width; negative counts rotate right.

`add32/64(left, right, carry)` and `sub32/64(left, right, borrow)` accept a boolean carry/borrow and return `(word, bool)`. `mul32/64` returns `(high, low)` for the full double-width product. `div32/64(high, low, divisor)` returns `(quotient, remainder)` or `DivisionError::DivisionByZero`/`Overflow`; a quotient overflows when `high >= divisor`. `rem32/64` accepts any high word and fails only on zero divisors. These are not constant-time cryptographic primitives. There are no native-word-width aliases or new intrinsic/syntax requirements.

### Portable SIMD

`use std::simd;` imports fixed-width value types with portable scalar implementations and GoML-owned amd64 SSE2/AVX2 assembly backends. The implementation does not import Go SIMD packages or third-party SIMD libraries.

| Element | 128-bit vector | 256-bit vector | Matching masks |
| --- | --- | --- | --- |
| `i8`, `u8` | `i8x16`, `u8x16` | `i8x32`, `u8x32` | `mask8x16`, `mask8x32` |
| `i16`, `u16` | `i16x8`, `u16x8` | `i16x16`, `u16x16` | `mask16x8`, `mask16x16` |
| `i32`, `u32` | `i32x4`, `u32x4` | `i32x8`, `u32x8` | `mask32x4`, `mask32x8` |
| `i64`, `u64` | `i64x2`, `u64x2` | `i64x4`, `u64x4` | `mask64x2`, `mask64x4` |
| `f32` | `f32x4` | `f32x8` | `mask32x4`, `mask32x8` |
| `f64` | `f64x2` | `f64x4` | `mask64x2`, `mask64x4` |

All vectors provide `from_array`, `to_array`, `splat`, `from_slice`, and `copy_to_slice`. Array conversions use the corresponding fixed-size scalar array and copy their elements. `from_slice(values, offset)` returns `Option[Vector]`; `copy_to_slice(values, offset)` returns `bool`. Both validate the complete range before accessing memory, require no special alignment, and reject negative offsets or incomplete vectors. An invalid store leaves the destination unchanged.

All vectors provide `add`, `sub`, `mul`, `neg`, `abs`, `min`, `max`, `reduce_sum`, `reduce_product`, `reduce_min`, and `reduce_max`. Integer arithmetic wraps at the element width, including reductions. Negating or taking the absolute value of the minimum signed integer preserves that minimum; unsigned `abs` is the identity. Integer vectors also provide `bitand`, `bitor`, `bitxor`, `bitnot`, `shl(count: u32)`, `shr(count: u32)`, `saturating_add`, `saturating_sub`, `reduce_and`, `reduce_or`, and `reduce_xor`. Shift counts are reduced modulo the element width. Signed right shifts extend the sign bit. Saturating arithmetic clamps to the scalar type's range. Integer vectors do not provide division.

Floating-point vectors additionally provide `div`, `sqrt`, `floor`, `ceil`, `trunc`, `round`, `round_ties_even`, `mul_add`, `is_nan`, `is_infinite`, and `is_finite`. `round` breaks halfway ties away from zero; `round_ties_even` chooses the nearest even integer. Ordinary arithmetic rounds each operation to the element type, so `a.mul(b).add(c)` has two roundings. `a.mul_add(b, c)` computes the fused result with one rounding, including in the scalar fallback. Reductions use a balanced tree with rounding at each operation. For example, four-lane `reduce_sum` computes `(lane0 + lane1) + (lane2 + lane3)`.

Floating-point `min` and `max` return the numeric operand when exactly one operand is NaN, and NaN when both are NaN. For equal zeros, `min` chooses negative zero if present and `max` chooses positive zero if present. Arithmetic NaN payloads are unspecified. Comparisons follow scalar rules, including unordered NaNs and equality of positive and negative zero.

`simd_eq`, `simd_ne`, `simd_lt`, `simd_le`, `simd_gt`, and `simd_ge` return the matching mask; vector `==` returns a single `bool` testing all lanes. Masks provide `from_array`, `to_array`, `splat`, `from_bitmask`, `to_bitmask`, `all`, `any`, `bitand`, `bitor`, `bitxor`, `bitnot`, and `select(if_true, if_false)`. Bit zero corresponds to lane zero; unused high bits are discarded. Bitmasks use `u8` for up to eight lanes, `u16` for sixteen lanes, and `u32` for thirty-two lanes. `select` accepts any matching vector through the public `SimdSelect[M]` trait and copies lane bits without arithmetic.

```goml
use std::simd;

fn select_scaled(a: simd::f64x4, b: simd::f64x4, scale: f64) -> simd::f64x4 {
    a.simd_gt(b).select(a, b).mul(simd::f64x4::splat(scale))
}

fn blend_bytes(a: simd::u8x16, b: simd::u8x16) -> simd::u8x16 {
    a.simd_lt(b).select(a.saturating_add(b), a)
}
```

`swizzle(indices: [u32; N])` selects lanes from one vector. `shuffle(other, indices: [u32; N])` selects from the concatenation of the receiver and `other`. Indices outside `0..N` or `0..2*N`, respectively, produce zero lanes. `reverse`, `interleave_low`, `interleave_high`, `deinterleave_even`, and `deinterleave_odd` provide common fixed rearrangements. Interleaving alternates elements of the two complete vectors and returns the lower or upper half of that sequence; deinterleaving selects even or odd positions from their concatenation. These operations apply across the full vector, including across 128-bit halves of a 256-bit vector.

Numeric conversions use `to_i8`, `to_u8`, `to_i16`, `to_u16`, `to_i32`, `to_u32`, `to_i64`, `to_u64`, `to_f32`, or `to_f64` when the destination vector exists with the same lane count and a different element type. Integer conversions extend according to the source signedness and discard excess high bits when narrowing. Floating-point to integer conversions truncate toward zero, clamp out-of-range values to the destination range, and map NaN to zero. Integer to floating-point and `f64` to `f32` conversions round to nearest, ties to even. Floating-point `to_bits` and `from_bits` convert to and from the matching unsigned integer vector without changing any bits.

```goml
use std::simd;

fn widen_and_scale(values: simd::i16x8) -> simd::f32x8 {
    values.to_i32().to_f32().mul(simd::f32x8::splat(0.5))
}

fn fused(a: simd::f32x4, b: simd::f32x4, c: simd::f32x4) -> simd::f32x4 {
    a.mul_add(b, c)
}
```

On amd64, `goml build/run/test` and `gomlc run-single` extract supported straight-line functions into generated `.s` files. Kernels may combine different supported vector types, matching masks, scalar broadcasts, conversions, and reductions, and may follow pure helper calls. They accept up to eight parameters and keep intermediates in registers. Unsupported expressions, control flow, or excessive register pressure retain the scalar implementation. Slice operations and dynamic shuffle indices remain scalar. Some numeric conversions and 32/64-bit saturating arithmetic also use scalar implementations. Integer reductions may use scalar instructions inside the native kernel. Native kernels have no allocations, callbacks, or stack spills.

Basic 128-bit kernels use SSE2. Kernels using 256-bit vectors or native floating-point rounding require AVX2; eligible 128-bit FMA kernels require FMA independently of AVX2. Dispatch checks CPU features and operating-system preservation of XMM/YMM state with CPUID and XGETBV. Unsupported features retain the scalar path. AVX/FMA kernels execute `VZEROUPPER` before returning. Set `GOML_SIMD=scalar` when building to disable native kernels, or `GOML_SIMD=sse2` to disable generated AVX2 and FMA kernels. Other architectures and exported standalone Go source retain scalar implementations; GoML build commands consume the native metadata embedded in generated Go source.

There are currently no configurable lane counts, gather/scatter, vector arithmetic operators, AVX-512 kernels, or automatic vectorization of ordinary scalar loops. SIMD uses the existing struct, import, trait, and method-call grammar and adds no expression syntax. Assembly snapshots under `gomlc/testdata/simd/` are generated and checked by `just update-golden` and `just verify-golden`.

### Randomness and cryptographic helpers

`std::crypto::hash::sha256` returns the lowercase hexadecimal SHA-256 digest of a byte buffer, and `sha256_file` hashes a complete file before returning. `std::crypto::rand::bytes` reads the requested number of bytes from the operating-system cryptographic random source. These APIs do not expose hasher or random-source handles.

`std::rand` uses the versioned `splitmix64-v1` algorithm. Every operation requires an explicit seed, and identical inputs produce identical outputs. It is intended for tests, simulations, sampling, and shuffling and is not cryptographically secure.

### Incremental checksums

`std::hash::Hasher` is a trait distinct from the prelude `Hash` trait used by collections. Its methods are `update(Slice[byte])`, `reset()`, `sum() -> bytes::Bytes`, `size()`, and `block_size()`. Import `use std::hash; use hash::Hasher;` for method-call syntax. `sum` is non-consuming and returns an independent byte snapshot; more updates remain valid. The checksum implementations below encode sums most-significant byte first and have a block size of one byte.

- `hash::adler32::checksum(input)` and `update(previous, input)` compute or continue an Adler32 checksum. The initial checksum is 1; `previous` is a checksum produced by this algorithm. `Digest::new()` supports `Hasher` and `sum32()`.
- `hash::crc32::Table::new(polynomial)` builds an immutable shared table for a reflected polynomial. Constants are `IEEE`, `CASTAGNOLI`, and `KOOPMAN`. `checksum(input, table)` starts at zero; `update(previous, table, input)` continues a checksum. `Digest::new(table)` supports `Hasher` and `sum32()`.
- `hash::crc64` has the same table/update/digest design, with `ISO`, `ECMA` and `sum64()`. The API follows Go's reflected, complemented CRC convention; the ECMA polynomial name does not imply an unreflected CRC-64/ECMA-182 parameter set.
- `hash::fnv::Digest::new32/new64/new128` constructs FNV-1 and `new32a/new64a/new128a` constructs FNV-1a. All implement `Hasher`; 128-bit state is implemented with ordinary integer arithmetic.

Assigning a digest copies a shared mutable handle. `digest.copy()` explicitly creates independent state; CRC copies may share the immutable table. Updates/reset must not race with other access to the same state. These checksums are not cryptographic hashes or authentication mechanisms. Hardware acceleration and Go-compatible state serialization are not provided.

`std::hash::stream` composes any Hasher with existing I/O:
`Reader::new(reader, hasher)` implements `io::Read`, and
`Writer::new(writer, hasher)` implements `io::Write`. Constructors preserve the
supplied hash state, allowing a prehashed prefix. `inner()` and `hasher()` return
the stored values using their ordinary alias/copy semantics; accessing the inner
stream directly bypasses hashing, and sharing a standard digest shares its state.
Use `Writer::new(io::Discard::new(), digest)` for a hash-only write sink.

Each operation calls the underlying stream once, validates its returned count,
and hashes only the successfully reported nonempty prefix. Short reads/writes
are not filled automatically; the inherited read_exact/write_all methods provide
that behavior and return Interrupted errors without retry. Empty requests are forwarded, but zero
counts and errors never update the hasher. Negative/oversized counts return
InvalidData. Underlying errors propagate unchanged and are not sticky; subsequent
calls are allowed. A stream that changes buffers or consumes bytes before returning
Err cannot report that partial progress through the current Read/Write contract,
so such unreported bytes are not hashed. There is no rollback of underlying I/O.
Writer.flush forwards the underlying flush without hashing or resetting anything.
The adapters neither close streams nor reset/copy digests automatically, allocate
no transfer buffers and add no synchronization. Keep stream and hash access
serialized and do not mutate write input during a call. The base hash package
still depends only on bytes; only the stream child adds an io dependency.
These adapters use existing traits and generics and add no syntax.

```goml
use std::bytes;
use std::hash;
use hash::Hasher;
use std::hash::fnv;
use std::hash::stream;
use std::io;
use io::Write;

fn hash_output[W: Write](writer: W, data: Slice[byte]) -> Result[bytes::Bytes, io::Error] {
    let digest = fnv::Digest::new64a();
    let output = stream::Writer::new(writer, digest);
    output.write_all(data)?;
    output.flush()?;
    Result::Ok(digest.sum())
}
```

```goml
use std::bytes;
use std::hash;
use hash::Hasher;
use std::hash::crc32;

fn checksum_parts() -> u32 {
    let digest = crc32::Digest::new(crc32::Table::new(crc32::IEEE));
    digest.update(bytes::Bytes::from_string("1234").as_slice());
    digest.update(bytes::Bytes::from_string("56789").as_slice());
    digest.sum32()
}
```

### Slice algorithms and checked vector edits

The pure `std::collections` slice helpers operate on explicit `Slice[T]` and `MutSlice[T]` views. They do not introduce a second collection type or change the builtin storage contract. All copies are shallow: entry storage can be independent while referenced objects remain shared.

| Capability | API and contract |
| --- | --- |
| Equality and comparison | `slice_equal`, `slice_equal_by`, `slice_compare`, `slice_compare_by`; lexicographic comparison returns `cmp::Ordering`, and `_by` supports different left/right element types |
| Membership | `slice_index`, `slice_index_by` return the first `Option[isize]`; `slice_contains` and `slice_contains_by` return bool |
| Copy/compact | Existing `Slice::to_vec`, `Vec::copy`, `Vec::dedup` and `dedup_by`; adjacent compaction preserves the first element of each equal run |
| Concatenation/repetition | `slice_concat(Vec[Slice[T]])`, `slice_repeat(Slice[T], count)` return checked independent `Vec` storage |
| Range editing | `vec_replace(Vec[T], start, end, replacement)`, `vec_insert_slice(Vec[T], index, inserted)`, `vec_delete_range(Vec[T], start, end)` return `Result[(), SequenceError]` |
| Predicate deletion | `vec_delete_if(Vec[T], predicate)` removes matching elements, preserving retained order and returning the removed count |
| Capacity reservation | `vec_grow(Vec[T], additional)` checks negative counts and length overflow before reserving; it does not change length or values |
| View mutation | `slice_reverse(MutSlice[T])`, `slice_sort(MutSlice[T])`, `slice_sort_by(MutSlice[T], compare)` affect only the view's bounds; sorting is stable and uses O(n) temporary storage |
| Sortedness/search | `slice_is_sorted`, `slice_is_sorted_by`, `slice_binary_search`, `slice_binary_search_by`; search returns `(insertion_index, found)` for the first equal value |
| Extrema | `slice_min`, `slice_max`, `slice_min_by`, `slice_max_by` return `None` for empty views and the first winner on ties |
| Iteration | `slice_all` and `slice_backward` yield `(relative_index, element)`; use existing `Slice::iter` for values only |
| Chunking | `slice_chunks(view, positive_size)` returns a checked single-pass iterator of bounded read-only subviews; the last chunk may be shorter |
| Collection | `vec_append_iter(iterator, destination)`, `slice_collect(iterator)`, `slice_sorted(iterator)` and stable `slice_sorted_by(iterator, compare)` accept any `Iterator` with the corresponding item type; iterator-first arguments establish the associated item type before checking the destination |

`SequenceError` distinguishes `InvalidRange(start, end, original_length)`, `InvalidCount(count)` and `LengthOverflow`. Negative repetition/reservation counts, nonpositive chunk sizes, and invalid edit ranges are recoverable errors. Checked lengths use the supported Linux amd64 `isize` range. Arithmetic validation does not promise recovery from allocation failure. An empty slice repeated any nonnegative number of times stays empty without looping that many times.

Range edits materialize their complete shallow result before writing, allowing replacement/insert input to overlap any part of the destination, even the whole vector. Index or length errors leave the destination unchanged. Editing mutates the shared `Vec` handle, so other aliases observe its new length and entries; preexisting `Slice` headers retain their original storage/length semantics. Edits use O(result length) temporary storage. Predicate deletion compacts in place; its predicate must not structurally mutate the vector.

Slice iterators and chunks retain the original bounded view and read elements when consumed, not an immutable value snapshot. Writes to that backing storage can therefore be observed, while later vector growth does not expand the view. Iterators are single-pass and do not synchronize concurrent access. Chunk boundaries avoid overflow even with a machine-maximum chunk size. Appending a vector's own ordinary iterator appends its original captured length once. Callers must terminate supplied iterators and avoid mutation during comparison callbacks.

Assigning an indexed iterator to another variable shares its cursor; call `slice_all` or `slice_backward` again for an independent traversal. Each returned chunk is a live view, not a copy. Equality rejects unequal lengths without invoking its comparator and stops on the first unequal pair. Predicate deletion visits original elements in order, preserves the order of retained elements and reports the number removed; callback panics do not promise rollback of earlier compaction writes.

Unlike Go slices, GoML public views cannot be resliced past their length or appended to, and have no public capacity to clip. Use `view.to_vec()` for an independent, length-sized copy rather than a Go `Clip` header operation; it also stops retaining out-of-view elements through that view. Generic ordered helpers require `cmp::Ord`; floating-point callers must choose an explicit comparator/NaN policy. `slice_equal` retains ordinary `PartialEq` behavior, including unequal NaNs. Sorted searches require input already ordered by the supplied comparator and use O(log n) comparisons without an O(n) validation pass.

```goml
use std::collections;

fn edit(values: Vec[isize]) -> Result[(), collections::SequenceError] {
    collections::vec_insert_slice(values, 0, values.as_slice())?;
    collections::slice_sort(values.as_mut_slice());
    let chunks = collections::slice_chunks(values.as_slice(), 2)?;
    for chunk in chunks {
        println(chunk.len());
    }
    Result::Ok(())
}
```

### Sorted collection boundaries

`std::collections::sort` and `stable_sort` retain the existing stable GoML merge-sort contract: equal elements keep their relative order, and sorting mutates the shared vector. The existing builtin-source implementations remain available for stage0 compatibility; no Go sorting backend is introduced. Generic default ordering uses `cmp::Ord`; floats require an explicit comparator and NaN ordering policy rather than an implicit total order.

`is_sorted`, `is_sorted_by` and `is_sorted_by_ordering` check adjacent pairs, accepting empty and singleton vectors. Integer comparators return negative/equal/positive according to their arguments' order; an `Ordering` comparator returns `Less`, `Equal` or `Greater`. They must define a consistent order and must not mutate the input.

`lower_bound(values, expected)` finds the first element not less than the target; `upper_bound` finds the first greater element. Both return an index in `0..=len`, including the insertion point when the target is absent. `equal_range` returns the half-open duplicate interval `(lower, upper)`. Their `_by` forms compare one element to the captured target; all three also have `_by_ordering` variants, including `equal_range_by_ordering`. The existing `binary_search` still returns the first matching index or `None`; it does not change to an insertion-point result. Descending order works by supplying a consistent reversed comparator. Inputs must already be sorted under that comparator; bounds use O(log n) comparisons and O(1) extra storage without rescanning to validate sortedness.

`search(length, predicate)` searches an abstract index range without allocating a vector. The predicate must be false then true as indices increase. The result is `Some(first_true_index)`, or `Some(length)` if all are false. Negative lengths return `None` without invoking the callback; zero returns `Some(0)` without invoking it. Midpoint arithmetic avoids overflow, including machine-maximum lengths. None of these helpers invokes the callback at `length`.

`search_by(length, compare)` returns `Option[(isize, bool)]`: the first index with a nonnegative comparison (or `length`) and whether that index compares equal. Negative lengths return `None`; empty ranges return `Some((0, false))`, both without calling the comparator. `search_by_ordering` accepts `cmp::Ordering` instead of an integer comparison. The callback compares the indexed element to the captured target, matching `lower_bound_by` and reversing Go `sort.Find`'s target-to-element convention. Its signs must form negative, then zero, then positive regions, with any region possibly absent. It may be called again at the candidate index to check equality, never outside `0..length`. These helpers use O(log n) comparisons and constant space, without materializing the sequence; for example, `collections::search_by(100, |index: isize| index * 2 - 37)` returns `Some((19, false))`. Comparator results must remain consistent and callbacks must not mutate the searched sequence.

All sorting variants perform O(n log n) comparisons with O(n) temporary storage, preserving equal-key order even for the names without `stable_`. Comparators must define a strict weak order, and must not mutate the vector, its aliases or ordering keys while sorting. A comparator panic may leave earlier merge passes committed; sorting does not promise transactional rollback or synchronization. Shared vector aliases and existing element views observe the final order. For integer keys near signed limits, compare relationally or use `cmp::compare` instead of subtraction, which can overflow.

```goml
use std::collections;

fn main() -> () {
    let values = Vec::from_array([1, 2, 2, 4]);
    let (start, end) = collections::equal_range(values, 2);
    println(end - start);
    println(collections::lower_bound(values, 3));
    println(collections::is_sorted(values));
}
```

### Logical slash paths

`std::path::slash` is separate from the existing host-filesystem `std::path` API. It treats only `/` as a separator; backslashes and drive-letter prefixes are ordinary text. No operation consults the working directory, resolves symbolic links, performs I/O, or provides a security sandbox. All inputs are valid UTF-8 GoML strings.

`clean` removes empty and `.` components and resolves `..` lexically. It preserves leading relative `..`, prevents traversal above a lexical `/` root, and returns `.` for an empty relative result. `join(Vec[string])` skips empty elements, joins with `/`, and cleans the result; all-empty input returns an empty string. A later absolute-looking element does not discard earlier elements. `split` preserves its input exactly as `(directory_including_last_slash, final_component)`. `base` strips trailing slashes, returning `.` for empty input and `/` for all slashes. `dir` cleans the directory returned by `split`. `extension` returns the suffix starting at the final dot of the final component, including that dot; it does not strip trailing slashes or special-case dotfiles.

`matches(pattern, name)` matches the entire name. `*` matches zero or more non-slash characters, `?` matches one non-slash Unicode scalar, `[...]` specifies scalar ranges, `[^...]` negates them, and `\\` escapes the next scalar. Classes must be nonempty; escape literal `-` or `]` within a class. Descending ranges match nothing. As in Go's `path.Match`, character classes may explicitly or negatively match `/`, unlike `*` and `?`. There is no recursive `**` extension. Malformed patterns return `MatchError::InvalidPattern` with a zero-based Unicode character position, even when the name would not match.

`Pattern::compile` parses once for repeated, independent matching. `compile_with_limit` sets the maximum number of pattern scalars; the default is `DEFAULT_PATTERN_LIMIT` (4096). `matches_with_limit` sets a work budget; `matches` uses `DEFAULT_WORK_LIMIT` (1000000). Negative limits are errors. Matching charges one unit per initial token, input scalar, token transition and character-range comparison. It uses linear pattern-sized state and polynomial dynamic programming, not recursive backtracking; exhausted limits return `PatternLimit` or `WorkLimit`, never a false match result. Caller-selected larger limits trade resource usage for larger inputs. Pattern fields are private and matching does not mutate shared compiled state.

Matching always advances at Unicode scalar boundaries. This intentionally differs from Go 1.26 path.Match's byte-wise retry after a star: Go matches `*[�]` and `*??` against `世` by decoding partial UTF-8 suffixes as replacement characters, whereas GoML returns false for both. An actual replacement scalar, as in `世�`, can match `*[�]`, and two scalars such as `世界` match `*??`. These differences are explicitly tested rather than interpreting an internal byte position as a character. Consecutive stars collapse to one matching token but still count individually against the source-pattern scalar limit. Each range comparison is charged, even for a state that cannot currently match; a failed budgeted call does not affect subsequent calls through the same Pattern or an alias.

The dynamic-programming matcher explores all valid token transitions. In particular, `*[a/]*` matches `a/`: the first star consumes `a`, the class consumes `/` and the final star is empty. Go 1.26 path.Match instead rejects this case after an earlier class match consumes `a` and its remaining star cannot consume `/`. GoML follows the documented whole-name pattern semantics here rather than reproducing that search-order behavior.

```goml
use std::path::slash;

fn matches_docs(name: string) -> Result[bool, slash::MatchError] {
    let pattern = slash::Pattern::compile("docs/*.md")?;
    pattern.matches(name)
}

fn main() -> () {
    println(slash::clean("docs/../images//icon.png"));
    println(slash::join(Vec::from_array(["docs", "guide", "../index.md"])));
    match matches_docs("docs/index.md") {
        Result::Ok(matched) => println(matched),
        Result::Err(error) => println(error.to_string()),
    }
}
```

### Base32

`std::encoding::base32::stream` provides `Encoder[W: io::Write]` and `Decoder[R: io::Read]`. Both constructors take the underlying handle, an Encoding and a cumulative output byte limit, returning `Result[..., io::Error]`. This child package depends on I/O; the core Base32 package remains independent of I/O.

The writer implements Write and Close. `finish()` or `close()` emits the tail and padding exactly once, without flushing or closing the underlying writer. `flush()` only forwards flushing and does not finalize pending bits. Writes after finish fail with BrokenPipe. Output-limit failures leave the submitted input unconsumed and permit smaller retries. Underlying write/flush failures are sticky because output may already have escaped: subsequent operations do not replay it. `failure()` retains the original error; terminal Interrupted is surfaced as Other, preserving the historical terminal-error mapping for compatibility. The default `write_all` no longer retries Interrupted. `input_len()` counts state-accepted input, not confirmed destination delivery. `is_finished()` requires successful finalization and no terminal failure.

The reader consumes encoded bytes incrementally, retains at most one decoded quantum, returns source Interrupted without retry and validates raw ASCII before decoding. Source, decode and limit errors become sticky and remain available through `failure()`, with original source error details preserved. Earlier returned bytes cannot be rolled back. Read through EOF to validate the entire input, including trailing bytes after padding; obtaining the expected decoded length alone is insufficient. Empty reads do not consume input, but still report a previously recorded failure. Handle copies share state and require external synchronization. Limits bound cumulative output, not total process memory; writer updates allocate output proportional to the supplied input chunk.

`std::encoding::base32::{encode, decode}` uses the uppercase RFC 4648 alphabet and `=` padding. `encode_with`/`decode_with` accept `Variant::Standard`, `StandardNoPadding`, `Hex`, or `HexNoPadding`; Hex means the extended Base32 alphabet `0123456789ABCDEFGHIJKLMNOPQRSTUV`, not hexadecimal byte encoding. Input bytes may be arbitrary binary data.

Decoding is strict and all-or-error: the standard variants reject lowercase, whitespace/newlines, misplaced or extra padding, invalid lengths, invalid digits, and nonzero unused trailing bits. `DecodeError::index()` reports a byte offset and `kind()` distinguishes `InvalidLength`, `InvalidByte`, `InvalidPadding` and `NonCanonical`. Unlike Go's Base32 decoder, no CR/LF is silently removed and no partial output is returned. These one-shot APIs materialize the complete result; the incremental state APIs emit chunks and the separate `base32::stream` package supplies I/O adapters.

`Encoding::new(alphabet: string, padding: Option[byte]) -> Result[Encoding, ConfigError]` constructs a reusable immutable configuration and decoding table. The alphabet must have exactly 32 distinct ASCII bytes, excluding CR/LF; an optional padding byte must also be ASCII, exclude CR/LF and not appear in the alphabet. This explicit ASCII restriction keeps every encoded result a valid GoML string; it is narrower than Go's arbitrary-byte alphabet. NUL, spaces and tabs may be deliberately configured and are then literal symbols, never ignored whitespace. `ConfigError` distinguishes invalid length, invalid/duplicate symbol byte offsets and invalid padding.

`Encoding::from_variant(variant)` provides the established alphabets and padding settings. `alphabet()` and `padding()` inspect a configuration; `with_padding(option)` returns a newly validated configuration without changing the original. `.encode(bytes)` and `.decode(text)` share the standard algorithm and canonical trailing-bit checks. With no padding, `=` may be an alphabet digit; otherwise a stray `=` remains rejected. For example, `Encoding::new("abcdefghijklmnopqrstuvwxyz234567", Option::Some(b'!'))` creates a lowercase alphabet with exclamation padding. Free functions and Variant retain their previous results and errors. These APIs use ordinary structs, enums, methods and generic containers without syntax changes.

`Encoding::encoded_len(input_length)` computes the exact padded or raw output byte length without allocating output and returns `Result[isize, OutputError]`; negative lengths and arithmetic overflow are recoverable errors. `.encode_checked(input, max_bytes)` rejects negative limits and excessive/overflowing output before encoding. `.decode_checked(input, max_bytes)` checks padding/length shape and the resulting decoded length before allocating the output buffer, then performs the usual digit and canonical-bit validation. Exact limits succeed; failures return no partial result. Free `encode_checked` and `decode_checked` functions use the standard padded configuration. Custom alphabets use the methods on their Encoding value.

`OutputError` distinguishes `InvalidLength(length)`, `InvalidLimit(limit)`, `LengthOverflow`, `LimitExceeded(limit)` and `Decode(original_error)`. Decode error precedence is negative limit, invalid shape, exceeded output limit, then invalid digits/trailing bits; an oversized malformed body can therefore report a limit error before a digit error. Limits bound output storage, not input length or total process memory. Existing unbounded entry points remain compatible, and no new syntax is introduced.

`Encoder::new(encoding, max_bytes) -> Result[Encoder, OutputError]` creates incremental encoding state with a cumulative output limit. `update(input: Slice[byte]) -> Result[string, StreamError]` emits all complete five-bit symbols available from that input and prior pending bits; concatenate returned chunks in call order. It retains only a scalar bit accumulator and counters, not input slices or emitted chunks. Calls can split input anywhere. `finish() -> Result[string, StreamError]` emits the final zero-filled symbol and configured padding, then closes the encoder; repeated finish succeeds with an empty string. Update after finish, even with empty input, returns `StreamError::Closed`.

Each update checks the final encoded length of the entire accepted input, including padding that finish will need, before allocating output or changing state. `StreamError::Output(error)` leaves the whole submitted chunk unconsumed and state unchanged, so callers can retry a smaller chunk. Every successful update therefore reserves enough budget to finish. `input_len()` reports accepted input bytes; `output_len()` reports emitted bytes rather than reserved final size. `is_finished()` reports closure. Handle copies share state; `reset()` resets all shared counters and pending bits, reopens the encoder and retains its configuration and limit. Calls require external synchronization. This pure incremental state API does not depend on I/O.

`Decoder::new(encoding, max_bytes) -> Result[Decoder, OutputError]` bounds cumulative decoded bytes. `update(input: string) -> Result[bytes::Bytes, StreamError]` emits decoded complete eight-symbol groups and retains at most seven pending ASCII bytes. Symbols are checked as they arrive; complete groups use the same strict padding and trailing-bit validation as one-shot decoding. A padded final group prohibits further nonempty input. Raw encodings may finish with a valid shorter group. `finish()` validates/emits that tail and closes the decoder; repeated successful finish returns empty bytes. A failed finish leaves the decoder open so a missing suffix can be supplied. Update after successful finish returns `Closed`, including empty input.

Each update is transactional for the supplied chunk: decode or limit failure returns no bytes from that chunk and leaves counters, pending input and terminal state unchanged. Earlier successful chunks are not rolled back. `StreamError::Output(OutputError::Decode(error))` reports absolute input byte offsets across previously accepted chunks; invalid length at finish may point one byte past the accepted input. Limits include bytes implied by pending data, so a valid pending tail can finish within budget. Error precedence follows incremental discovery rather than necessarily matching one-shot shape-first validation. `input_len()`, `output_len()`, `is_finished()` and shared-handle `reset()` follow the encoder conventions. Returned byte buffers are independent; input strings and emitted chunks are not retained. Decoder calls require synchronization. Reader/Writer adapters live in the separate `base32::stream` package; these APIs add no grammar.

```goml
use std::bytes;
use std::encoding::base32;

fn token_text() -> string {
    base32::encode_with(bytes::Bytes::from_string("foobar"), base32::Variant::StandardNoPadding)
}
```

### Hexadecimal and base64

`std::encoding::hex` encodes `bytes::Bytes` to lowercase hexadecimal by default, while `encode_upper` emits uppercase digits. Decoding accepts either case and returns `DecodeError` with the byte offset of an odd length or invalid digit.

`std::encoding::base64` uses the padded RFC 4648 standard alphabet by default. `Variant` selects standard or URL-safe alphabets with required or omitted padding. Decoding is strict: it rejects invalid lengths, alphabet mixing, misplaced padding, and nonzero unused trailing bits.

### PEM framing

`std::encoding::pem` implements textual PEM framing in pure GoML over the standard Base64 codec. It does not encrypt/decrypt private keys, parse ASN.1, validate certificates or implement trust policy. `Block::new(label, headers: Vec[(string, string)], data: bytes::Bytes) -> Result[Block, Error]` validates and snapshots its inputs. `label()`, `headers()` and `data()` expose values with independent mutable snapshots. Labels are nonempty printable ASCII without leading/trailing spaces. Header keys are nonempty printable ASCII without spaces or colons; values permit printable ASCII and tabs, with outer ASCII whitespace normalized away. Newlines and duplicate case-sensitive keys are rejected.

`encode(block, max_bytes) -> Result[string, Error]` emits LF-delimited BEGIN/END lines and padded Base64 wrapped at 64 columns. Headers are sorted lexically with `Proc-Type` first and followed by a blank line. Empty payloads have no body line. The complete byte length, including headers and line endings, is checked before output construction; exact limits succeed, negative limits and oversized output fail. No partial string is returned.

`decode(input: string, max_input, max_data, max_headers) -> Result[Option[(Block, isize)], Error]` finds the first line-start BEGIN candidate after optional preamble. Success returns its block and the absolute byte offset immediately following its END line, including a terminating newline if present. Slice the original string at that offset to parse subsequent blocks. `None` means no BEGIN candidate. LF and CRLF, trailing line spaces/tabs and spaces/tabs inside the Base64 body are accepted. The whole input byte limit is checked first; decoded byte and header-count limits are independent. Temporary parsing storage is bounded by the input limit. Header field count includes duplicates before validation.

Header normalization removes only outer ASCII space and tab bytes. Unicode whitespace around keys or values is rejected as invalid header text rather than silently removed before validation. This keeps decoded blocks within the same printable-ASCII contract as explicitly constructed blocks.

This is deliberately stricter than Go's searching decoder: malformed candidates, mismatched/missing END labels, duplicate headers and noncanonical Base64 trailing bits return errors rather than being silently skipped. Input is a valid UTF-8 string, not an arbitrary binary preamble. `Error::Malformed(offset)` uses original input byte offsets; nested Base64 error offsets refer to the whitespace-filtered body. Invalid labels/headers, duplicate names, negative limits and exceeded bounds have distinct error variants. The parser materializes one block, provides no cryptographic authenticity guarantee and is not yet a streaming I/O adapter.

For example, `pem::decode(source, 1048576, 65536, 16)` accepts at most a one-MiB input containing a block with 64-KiB decoded content and 16 headers. Existing generic containers, structs, enums and calls express these APIs; no grammar or compiler intrinsic is added.

### Time

`Duration` stores a non-negative number of nanoseconds and offers constructors and whole-unit accessors for nanoseconds, microseconds, milliseconds, and seconds. Subtraction saturates at zero. `Instant` is monotonic and is suitable for elapsed-time measurement. `SystemTime` exposes Unix nanosecond, millisecond, and second timestamps. `time::sleep` blocks the current goroutine for a `Duration`.

`Duration` provides checked and saturating scaled constructors, addition, subtraction, and multiplication. Checked operations return `None` on overflow, underflow, or a negative input. Saturating operations clamp to zero or the largest signed 64-bit nanosecond value. `Duration`, `Instant`, and `SystemTime` expose `compare`; `Instant::checked_duration_since` returns `None` when the receiver precedes the supplied instant.

### Files and standard streams

`fs::read_file_structured`, `read_bytes_structured`, `write_file_structured`, and `write_bytes_structured` perform whole-file I/O. They return `fs::Error` with a stable `io::ErrorKind`, operation, path, optional raw operating-system code, and display message. Directory creation and removal, canonicalization, directory listing, and file hashing use the same error type.

`fs::create_dir`, `rename`, `copy`, `hard_link`, and `symbolic_link` are eager operations returning `fs::Error`. `copy` reads and writes the complete file and currently creates the destination with portable `0644` permissions; it does not preserve source metadata.

`fs::metadata` and `symlink_metadata` return value-only metadata including file type, length, portable permission bits, and modification time. `Metadata::from_parts(file_type, length, mode, modified_unix_nanoseconds)` constructs the same value without filesystem access; permission bits are masked to `0o777`. `read_dir_structured` eagerly snapshots directory entries. `read_dir_names_structured` returns sorted names without fetching each child's metadata, preserving structured listing errors; a disappearing child does not invalidate the other names. `atomic_write` writes, synchronizes, closes, and atomically renames a same-directory temporary file before returning; temporary cleanup stays inside the runtime call. `replace` exposes the host atomic rename operation under replacement semantics.

`io::read_stdin_structured`, `read_stdin_exact_structured`, `write_stdout_structured`, and `write_stderr_structured` provide structured errors for standard streams. `read_stdin_to_string` validates the complete input as UTF-8 and reports `InvalidData` on failure. A negative exact-read length reports `InvalidInput` before accessing stdin.

### Portable metadata values and formatting

`fs::FileMode` adds portable value semantics without changing the existing
four-case FileType enum. `NodeKind` distinguishes File, Directory, Symlink,
BlockDevice, CharacterDevice, NamedPipe, Socket, Irregular and Unknown. Its
compatibility `file_type()` maps the first three directly and every other kind
to FileType::Other. These are GoML values, not Go or host ABI mode bit numbers.

`ModeFlags` has six public bool fields: set_uid, set_gid, sticky, append_only,
exclusive and temporary. `FileMode::from_parts(kind, permissions, flags)` takes
a NodeKind, u32 permission bits and Option[ModeFlags]. Permissions are masked to
`0o777`; flags=None means these attributes are unknown, which differs from
Some containing six false values. Kind and flags knowledge are independent, so
a known NamedPipe with unknown flags is representable.
`FileMode::new(kind, permissions, flags)` is the known-flags convenience form.
`FileMode::from_file_type(file_type, permissions)` maps the old categories and
sets flags to None, mapping Other to Unknown. Getters expose `kind()`,
`permissions() -> Permissions`, `flags()` and the compatible `file_type()`.
NodeKind, ModeFlags and FileMode implement PartialEq and Eq.

`Metadata::from_mode(mode, length, modified_unix_nanoseconds)` preserves the
complete FileMode, retrievable through `.mode()`. Existing file_type, len,
permissions and time getters retain their meanings. The old from_parts
constructor still masks permissions and now explicitly records unknown flags.
Existing host queries provide only their prior categories and nine permission
bits: they cannot recover special attributes or distinguish the advanced kinds
hidden behind Other. Host metadata therefore also reports unknown flags, never
an invented all-false value. No native interface has been expanded.

`SnapshotEntry::new(name, metadata)` stores a name and a metadata value and
implements DirectoryEntry without filesystem I/O. Its metadata method always
returns that snapshot; `.mode()` exposes the full mode directly. Construction
does not validate the supplied name; normal directory collection still rejects
invalid component names. This is a public value adapter for providers, not a
new host directory backend or a claim that metadata stays current on disk.

`format_metadata(name, metadata, max_bytes)` and
`format_directory_entry(entry, max_bytes)` return `Result[string, fs::Error]`.
They use deterministic, locale-independent display formats, with no newline:

```text
kind=file permissions=rw-r----- flags=unknown size=12 modified_ns=-7 name=example
kind=directory name=config/
```

The metadata format includes the advanced kind in snake_case, nine rwx/dash
permission positions, flags, signed byte length and signed Unix nanoseconds.
Known flags appear as `set_uid:0,set_gid:0,sticky:0,append_only:0,exclusive:0,temporary:0`,
with each value replaced by 1 when set. Unknown flags use the word `unknown`.
Entry formatting calls only name and file_type, once each for a nonnegative
budget, and never fetches metadata; its compatible kind labels are file,
directory, symlink and other. Both formats append `/` for a directory, even if
the supplied name already ends in a slash. Names are copied verbatim, not
escaped; these display strings are not a safe logging or serialization format.

Negative budgets return InvalidInput, before any entry method is called.
Insufficient output budgets return InvalidData without a partial string. Each
field is checked against the remaining UTF-8 byte allowance before copying;
there is no truncation inside a character. Error operations are `"format metadata"`
and `"format directory entry"`; their paths contain the supplied or retrieved
name, except a negative entry budget has no name yet. Arbitrary entry-method
panics follow normal panic propagation, not error conversion.
The byte allowance bounds output length, not allocator capacity or process
memory, and does not make allocation failure recoverable.

```goml
use std::fs;

fn describe_pipe() -> Result[string, fs::Error] {
    let mode = fs::FileMode::from_parts(fs::NodeKind::NamedPipe, 0o600, Option::None);
    let metadata = fs::Metadata::from_mode(mode, 0, 0);
    fs::format_metadata("events", metadata, 256)
}
```

### Portable read-only filesystems

`fs::File` extends `io::Read + io::Close` with
`stat(self: Self) -> Result[fs::Metadata, fs::Error]`. `fs::FileSystem` has an
associated `Handle: File` and
`open(self: Self, name: string) -> Result[Self::Handle, fs::Error]`. Implementations
provide their own storage and return a fresh logical open handle on each
successful open. Reading follows the normal `io::Read` count/EOF/error contract;
metadata inspection must not move its read position. Handle copies may share
position and closed state; these traits do not imply concurrent safety.
Providers must reject invalid logical paths on direct `open` calls as well;
the generic helpers also validate before dispatching to a provider.
These portable protocols coexist with the host-path functions above and do not
introduce a native file backend or reinterpret OS files as memory snapshots.

`fs::valid_path(name)` accepts `"."` for the root, or nonempty slash-separated
components other than `"."` and `".."`. It rejects absolute paths, trailing
slashes and empty components without normalizing the input. Backslash, colon
and NUL are ordinary logical-name characters; a concrete provider may reject
names its storage cannot represent. Strings use GoML's UTF-8 representation.
Validation is lexical, not an OS path-safety or sandbox guarantee.

`fs::read_file_from(filesystem, name, max_bytes)` returns `Result[bytes::Bytes,
fs::Error]`; `fs::stat_from(filesystem, name)` returns `Result[fs::Metadata,
fs::Error]`. Invalid names and negative byte limits fail before opening. After
open succeeds, each helper closes the handle exactly once on normal completion,
including failures; a deferred cleanup also attempts close during panic
unwinding. A primary operation error wins over a returned close error; a close
error after a successful operation is returned. Open/stat provider errors remain
unchanged. Read/close errors preserve kind, raw OS code and message, with the
requested logical name and operation `"read file"`/`"close file"`. Provider
panics follow the ordinary panic/defer rules, not error-return conversion.

Reading uses the `read` primitive, not a provider's potentially overridden
`read_to_end` helper. It checks counts, stops immediately on any error including
`Interrupted`, and reads at most one byte beyond the limit to distinguish exact
EOF from excess input. Excess input or invalid counts return `InvalidData`;
partial bytes are not returned on failure. Metadata size is not trusted as a
read bound. The limit bounds returned bytes, not all allocations or provider
side effects; allocation failure is not made recoverable by these helpers.

`fs::SubFs::new(filesystem, root)` validates a logical root without opening it
or checking that it exists or is a directory. Its `FileSystem` implementation
validates each requested name, then prefixes the root without cleaning or
rewriting components. Root `"."` is an identity; opening `"."` selects the
configured root. Nested SubFs values compose. Underlying handles and provider
errors are returned unchanged, so errors may mention the prefixed path. SubFs
is a namespace adapter, not confinement: symlinks or provider policies can
resolve outside that namespace.

```goml
use std::fs;
use fs::FileSystem;
use std::bytes;

fn load_settings[F: FileSystem](filesystem: F) -> Result[bytes::Bytes, fs::Error] {
    let config = fs::SubFs::new(filesystem, "config")?;
    fs::read_file_from(config, "settings.toml", 65536)
}
```

`fs::DirectoryEntry` describes a stable `name()` and `file_type()` plus a possibly
lazy `metadata() -> Result[Metadata, Error]`. Entries must remain usable after
later pages and directory close; they must not borrow storage reused by the
enumerator. Metadata may fail if an entry disappears or permissions change.
The existing host `DirEntry` implements this trait: names/types retain their
listing snapshot, while metadata calls `symlink_metadata` on the stored path.
The existing eager `read_dir_structured` API is unchanged.

`fs::ReadDirFile: File` has `type Entry: DirectoryEntry` and
`read_dir(amount: isize) -> Result[Vec[Self::Entry], ReadDirError[Self::Entry]]`.
Providers must reject nonpositive amounts with `InvalidInput` before advancing
the cursor. A successful page contains at most amount entries; an empty page
means EOF at that call, whereas a short nonempty page does not. Errors may
include confirmed consumed entries. No retry or sticky-error behavior is
implied by the trait. Names must be single valid components, excluding `"."`
and `".."`; NUL, backslash and colon remain ordinary logical characters.

`ReadDirError::new(entries, error)` copies the vector structure, not the entry
objects. `.entries()` exposes a read-only slice and `.error()` the original
cause. Its display/debug representations use the cause without imposing
formatting bounds on the entry type. Unknown side effects of a failed provider
call are not represented by fabricated entries.

`fs::read_dir_from(filesystem, name, max_entries)` returns
`Result[Vec[F::Handle::Entry], ReadDirError[F::Handle::Entry]]` for
`F: FileSystem` where `F::Handle: ReadDirFile`. It checks the path and nonnegative
budget before opening, then requests pages of at most 128 entries. Near the
limit it requests one extra entry to distinguish EOF from excess data; zero
budget still opens and probes one entry. The probe can be consumed but is not
returned. The budget bounds result cardinality, not total allocations, name
lengths or provider side effects. It does not promise a consistent directory
snapshot or make allocation failure recoverable.

Each page is validated before any of it is accepted. Too many entries, or any
invalid name (including in the probe), rejects the whole page with `InvalidData`
and preserves earlier pages. Oversized pages are rejected without calling entry
methods. Otherwise each name is fetched once and cached; metadata and file_type
are never called by this helper. Provider errors stop immediately, including
`Interrupted`, and retain their original cause plus valid entries that fit the
budget. On a valid error page, the provider cause wins over a limit excess;
invalid-page errors take priority over both. Successful excess data instead
produces `InvalidData`. The retained prefix is selected in provider order and
then stably sorted by cached names; duplicate names are retained. Success and
all errors return sorted confirmed results, never more than the budget.

Every successfully opened handle is closed once, including during panic cleanup.
An operation error wins over an ordinary close error. A close-only error retains
the complete list and uses the same error mapping as `read_file_from`. A close
panic follows normal panic propagation and can replace an earlier panic; it is
not converted into an ordinary close error. SubFs
forwards its handle type, so nested directory composition needs no separate
adapter or dynamic interface detection. For example:

```goml
use std::fs;
use fs::{FileSystem, ReadDirFile};

fn list_config[F: FileSystem](filesystem: F) -> Result[
    Vec[F::Handle::Entry], fs::ReadDirError[F::Handle::Entry],
] where F::Handle: ReadDirFile {
    fs::read_dir_from(filesystem, "config", 1024)
}
```

`fs::walk_from(filesystem, root, limits, visit)` performs a bounded, stable,
lexically ordered depth-first traversal without recursion. It returns
`Result[(), fs::Error]` and requires the same FileSystem/ReadDirFile bounds as
`read_dir_from`. The callback takes the logical path,
`Option[fs::WalkEntry[F::Handle::Entry]]` and `Option[fs::Error]`, returning
`Result[fs::WalkControl, fs::Error]`. WalkEntry implements DirectoryEntry: its
name/type are cached, root metadata comes from the initial stat, and child
metadata remains lazy. The walker does not fetch child metadata. Cached names
are shared with directory validation and sorting, so the provider's name method
is not called again when constructing child paths. File types are fetched only
for children actually visited, once each.

`WalkControl::Continue` visits a directory's children; `SkipDir` skips them.
For a non-directory, SkipDir skips the remaining siblings in its parent, not
the remaining entries at ancestor levels. `SkipAll` successfully stops the whole walk.
A callback Err stops immediately and is returned unchanged. The callback runs
before opening a directory, allowing it to skip that I/O. Duplicate names are
retained in stable provider order. Child symlinks are visited but not followed;
the root uses the provider's open/stat behavior, which may follow a link.
Traversal does not promise confinement, cycle detection or an atomic snapshot.

The root is inspected through `stat_from`, including its separate open/close.
If this fails, the callback receives no entry and the original error; any
successful control response ends the walk successfully. A directory read/open
or close failure instead triggers a second callback for that directory, with
its entry and the error. Continue traverses its confirmed, sorted prefix;
SkipDir discards that prefix, and SkipAll stops. Provider errors and ordinary
close precedence match `read_dir_from`. All successfully opened handles close
once, before any callback; callbacks therefore run without a held handle.
Provider and callback panics follow normal panic/defer propagation, not Result
conversion.

`WalkLimits` has four explicit, public isize fields:

- `max_depth`: the root has depth zero. A directory at the maximum depth may be
  visited and skipped, but Continue would exceed the expansion limit and fails.
- `max_directory_entries`: the per-directory result limit, with the same
  limit-plus-one probe as read_dir_from, further capped by remaining work.
- `max_directories`: directory expansion attempts, including failed opens;
  initial root stat and directories skipped before expansion do not count.
- `max_work`: one unit for the initial root stat operation, one per directory
  open attempt, one per directory page call, and one per returned entry with a
  valid count. Entry units include probes and pages rejected for invalid names.
  Oversized counts are rejected before entry methods or entry charging. A page
  call requires at least two remaining units for the call and a possible entry;
  its requested count is capped by the remaining work minus one. Unused entry
  capacity is not charged, but a last EOF probe still needs this reservation.
  Ordinary callbacks and mandatory cleanup do not consume units.

Invalid paths or negative limits return InvalidInput before I/O or callbacks.
Zero work reports a limit error to the callback without I/O. Algorithmic limit
failures use InvalidData with operation `"walk"` and the current logical path.
The error callback may explicitly stop with SkipAll or replace the error with
Err; Continue and SkipDir cannot erase a limit failure or resume traversal.
Thus work and directory budgets bound even cyclic or repeatedly failing
providers. These are algorithmic budgets, not limits on arbitrary provider or
callback execution time, name lengths, total allocation or process memory.

```goml
use std::fs;
use fs::{FileSystem, ReadDirFile};

fn inspect_tree[F: FileSystem](filesystem: F) -> Result[(), fs::Error]
where F::Handle: ReadDirFile {
    let limits = fs::WalkLimits {
        max_depth: 32,
        max_directory_entries: 4096,
        max_directories: 1024,
        max_work: 100000,
    };
    fs::walk_from(filesystem, ".", limits, |path, _, error| {
        if let Some(error) = error {
            Result::Err(error)
        } else {
            println(path);
            Result::Ok(fs::WalkControl::Continue)
        }
    })
}
```

`fs::glob_from(filesystem, pattern, limits)` returns
`Result[Vec[string], fs::GlobError]` with the same FileSystem/ReadDirFile bounds.
It uses `std::path::slash::Pattern` for each component and an explicit stack,
not recursive traversal. Slash always separates components: escaped slashes
and slashes within a character class are rejected when compiling the affected
component. All components compile before any filesystem I/O. Empty, absolute,
trailing-slash and empty-component paths are invalid; literal dot/dot-dot
components, including escaped spellings, are invalid except the exact pattern
`"."` for the root. Paths are not cleaned. Repeated stars have the ordinary
single-component Pattern meaning; `**` is not recursive glob syntax.

Entirely literal patterns, including escaped literal metacharacters, use a
single `stat_from` on the decoded path without enumerating its parents. For a
pattern containing wildcards, the longest initial literal prefix selects the
first directory to enumerate. Subsequent components use the compiled matcher.
Final matches need only names, not type or metadata. Intermediate matched File
and Other entries are skipped; Directory and Symlink entries are opened through
the provider for the next component. Unlike walk_from, glob may therefore
follow intermediate links. Neither helper provides confinement.

Successful glob results are globally, stably sorted by the full logical path,
with duplicates retained; no matches is a successful empty vector. Provider
errors are not silently ignored, including NotFound on a literal path and
Interrupted. A failure on an intermediate directory stops without opening its
confirmed children. On the final component, confirmed directory entries may
still be matched within the remaining budgets, without further filesystem I/O;
then the original provider or close error is returned. Such an already-known
error wins over subsequent matching or result-limit errors. Algorithmic
directory/work-limit failures instead stop immediately without matching that
batch. Earlier confirmed matches survive all failures.

`GlobError::matches()` exposes the confirmed, sorted prefix and `.cause()`
returns `GlobCause::FileSystem(fs::Error)` or
`GlobCause::Pattern(component_index, slash::MatchError)`. Component indices are
zero-based; positions inside MatchError are relative to that component.
Provider errors retain their fields. Glob-generated InvalidInput/InvalidData
errors use operation `"glob"`; their path is the input pattern for preflight
failures or the current logical path during execution. GlobError's display and
debug formats show the cause. `GlobError::new(matches, cause)` copies the vector
structure in the supplied order; the helper, not this constructor, sorts it.

`GlobLimits` requires seven explicit, nonnegative public isize fields:

- `max_pattern_chars` bounds the whole input's Unicode character count,
  including separators and escapes; `max_components` bounds its component
  count. The root pattern `"."` counts as one character and one component.
- `max_directory_entries` and `max_directories` bound each directory's collected
  entries and total expansion attempts. Entirely literal stat calls do not count
  as directory expansions.
- `max_match_calls` bounds matcher invocations across all directories.
- `max_matches` bounds returned matches. Zero still searches for a first match
  to distinguish no result from excess results; a literal pattern still stats
  its path. Retained results are selected in traversal order, then globally
  sorted, rather than selecting the lexically smallest possible subset.
- `max_work` is shared across all directory I/O and matching. Directory work has
  the same open/page/entry accounting and EOF reservation as walk_from; a fully
  literal stat costs one unit. Before matching a candidate, glob precharges
  `P + N * (1 + 2 * P)` units, where P is the component's original character
  count and N is the candidate name's UTF-8 byte length. This conservative bound
  covers token, character and class-range work in Pattern; unused matching
  allowance is not refunded. Division checks avoid overflow before computing
  the cost. Pattern compilation is separately bounded by the first two limits
  and does not consume max_work.

Negative limits return InvalidInput before I/O; exceeded limits return
InvalidData. Sorting, arbitrary provider execution and total allocations are
not a time or memory quota. Successful opens close exactly once, including
panic cleanup; panic and ordinary close precedence are unchanged.

```goml
use std::fs;
use fs::{FileSystem, ReadDirFile};

fn text_files[F: FileSystem](filesystem: F) -> Result[Vec[string], fs::GlobError]
where F::Handle: ReadDirFile {
    let limits = fs::GlobLimits {
        max_pattern_chars: 1024,
        max_components: 32,
        max_directory_entries: 4096,
        max_directories: 1024,
        max_match_calls: 100000,
        max_matches: 4096,
        max_work: 10000000,
    };
    fs::glob_from(filesystem, "docs/*.txt", limits)
}
```

`fs::LinkFileSystem: FileSystem` supplies `read_link(name) -> Result[string,
fs::Error]` and `symlink_metadata(name) -> Result[fs::Metadata, fs::Error]`.
The latter inspects the final entry without following that link; resolution of
earlier components is provider-defined. Providers validate direct-call logical
names. `read_link_from` and `symlink_metadata_from` also validate before
dispatching, then return the provider's result unchanged. The helpers require
LinkFileSystem statically: there is no dynamic capability test or fallback that
opens/follows a link instead of inspecting it.

Link targets are opaque provider strings, not validated logical input paths.
They may be relative, absolute, contain dot-dot or use provider-specific syntax.
`SubFs[F]` implements LinkFileSystem when F does, prefixing the requested name
exactly as for open while leaving targets and errors unchanged. A target is not
cleaned or rebased into the sub-filesystem. This is namespace composition, not
a symlink-safe sandbox. These protocols add no native backend. Chained associated
types, bounds and helpers use existing syntax; no grammar changes are needed.

### In-memory test filesystems

Import `std::testing::fs as memory` for the pure GoML, read-only `MemoryFS`.
It implements `fs::FileSystem` and `fs::LinkFileSystem`; its `MemoryHandle`
implements `fs::File`, `fs::ReadDirFile`, `io::Read`, `io::ReadAt`, `io::Seek`
and `io::Close`. This package is a portable test utility, not an OS mount or
a replacement for the host filesystem.

`MemoryFS::new(nodes: Vec[MemoryNode], limits: MemoryLimits)` returns
`Result[MemoryFS, fs::Error]`. Each node has public `path`, `mode`,
`modified_unix_nanoseconds` and `content` fields. `MemoryContent` is `File(Bytes)`,
`Directory`, `Link(string)` or `Special`; the first three require the matching
NodeKind, and Special requires a different kind. Convenience constructors
`MemoryNode::file`, `directory` and `link` use permissions 0444, 0555 and 0777,
zero modification time and known all-false flags. Construction copies file
bytes and containers; later mutation of the input cannot change the snapshot.

Names use `fs::valid_path`. Duplicate paths fail with AlreadyExists; non-directory
ancestors, inconsistent kinds and a non-directory root fail with InvalidInput.
Missing parents and root are inferred as 0555 directories, preserving explicitly
supplied directory metadata. `node_count()` includes inferred nodes. Metadata
length is the byte length of file data or the raw link target, and zero for
directories and special nodes. Directory entries have unique, sorted names.

MemoryLimits has six nonnegative fields. `max_nodes` includes inferred nodes;
`max_path_bytes` sums all unique stored path byte lengths, including root `.`;
`max_data_bytes` sums file and link-target bytes; `max_depth` counts path
components with root at zero. Exceeding a construction limit returns InvalidData;
negative limits return InvalidInput. These are logical data bounds, not an
allocator or CPU-time quota.

Each resolution separately uses `max_link_hops` and `max_resolve_work`.
Work charges input bytes, one root lookup, one per processed component, one
per child lookup and the bytes of every followed link target. Thus resolving
`.` costs three units. Logical input validation precedes this counter. These
units do not count every byte processed by path concatenation or hashing and
are not a CPU-time quota. Link
targets are resolved relative to their parent, without lexical pre-cleaning:
dot-dot applies to the directory actually reached. Repeated separators and
dot components are supported in targets, but traversal through a file, including
a trailing separator, fails. Absolute targets and attempts to leave the virtual
root fail with InvalidInput; empty targets and missing entries fail with
NotFound; exhausted work/hops fail with InvalidData. Backslash, colon and NUL
remain ordinary name characters. `read_link` returns the unmodified final link
target, and `symlink_metadata` does not follow that final link; intermediate
links are resolved. Errors retain the original requested name.

Every open creates a separate cursor; copies of one handle share its cursor
and closed state. Stat does not advance it. Reads copy bytes into caller storage,
return zero at EOF, and positional reads never move the cursor. Short positional
reads report UnexpectedEof with their confirmed byte count; negative or
overflowing ranges fail before copying. Seek permits positions past EOF but
rejects negative/overflowing results without moving the cursor. Directory pages
require a positive amount, return fresh vectors and cached SnapshotEntry values,
and repeatedly return an empty page at EOF. Entries remain usable after close.
Directories reject byte reads and seeks; special nodes return Unsupported for
data and directory operations. Close is idempotent. Primitive stat/read/read_at/
seek/read_dir operations on a closed handle fail with InvalidInput before other
checks. Inherited I/O helpers retain their own semantics, including an empty
read_exact succeeding without calling read. Concurrent operations on one shared
handle are not supported.

```goml
use std::bytes;
use std::fs;
use std::testing::fs as memory;

fn sample_tree() -> Result[bytes::Bytes, fs::Error] {
    let limits = memory::MemoryLimits {
        max_nodes: 16, max_data_bytes: 1024, max_path_bytes: 1024,
        max_depth: 8, max_link_hops: 16, max_resolve_work: 4096,
    };
    let source = memory::MemoryFS::new(Vec::from_array([
        memory::MemoryNode::file("config/name", bytes::Bytes::from_vec(Vec::from_array([b'o', b'k']))),
        memory::MemoryNode::link("latest", "config/name"),
    ]), limits)?;
    fs::read_file_from(source, "latest", 16)
}
```

### Filesystem contract reports

`memory::check_fs(source, expected, limits)` checks a stable, read-only provider
and returns `Result[CheckReport, fs::Error]`. It requires `FileSystem` with a
`ReadDirFile` handle. Invalid checker inputs return InvalidInput before any
provider call. Findings about a provider are returned in the report rather
than raised as testing assertions. The caller must keep the filesystem stable
for the duration of the check; changed content is reported as snapshot
inconsistency, not proof that a live or mutable filesystem violates its traits.

`ExpectedTree::Contains(paths)` requires those physically enumerated paths and
allows additional entries. Paths are validated and copied before calling the
provider; duplicates are ignored. An empty Contains list imposes no coverage
requirement. `ExpectedTree::Empty` requires no entries below root `.`. Traversal
does not follow child symlinks, so expected paths through such links are not
implicitly expanded. Missing paths are only reported after traversal finishes
without a prior finding or exhausted budget.

The basic checker probes invalid logical names, validates sequential read
counts and empty reads, compares interleaved independent opens with different
read widths, and checks that stat and empty reads do not disturb position.
It compares confirmed content length and metadata under the stable-snapshot
assumption. Directory checks compare different page widths, probe nonpositive
amounts before and during enumeration, preserve error prefixes, and recheck
entry names, types and metadata after later pages and close. Raw pages may be
unsorted and may contain duplicate names; comparisons retain multiplicity.
NUL, backslash and colon remain valid component characters. Child symlink
metadata is not compared with following open/stat. Special nodes are inspected
but are not assumed to support regular-file reads. Generic close behavior is
not tested against MemoryHandle-specific idempotence or closed-state policies.

CheckLimits contains nonnegative `max_nodes`, `max_depth`,
`max_directory_entries`, `max_file_bytes`, `max_total_read_bytes`,
`max_provider_calls`, `max_issues` and `max_path_bytes`. Nodes and full path
bytes are charged at discovery, including root and duplicate entries; depth
starts at zero. Every explicit provider operation, including entry accessors,
consumes a call. Confirmed successful read counts across all rereads consume
the total byte budget. File and directory limits permit one additional probe
to distinguish EOF from excess; a read probe also needs remaining total-byte
allowance. Without enough allowance to prove EOF, checking stops as incomplete.
Cleanup is mandatory and exempt from these budgets. Bounds do not interrupt a
blocking provider, bound its internal allocations, or impose a CPU-time quota;
directory multiset comparison may perform quadratic local comparisons. Input
validation and copying the expected list also precede provider-call accounting.

Report accessors are `issues()`, `complete()`, `passed()`, `stop_reason()`,
`discovered_nodes()`, `path_bytes()`, `read_bytes()`, `provider_calls()` and
`cleanup_calls()`. `issues()` is a read-only snapshot. Every issue has a path,
operation and CheckIssueKind: `ContractViolation(message)`,
`ProviderError(fs::Error)`, `InconsistentSnapshot(message)` or
`ProviderPanic(panic::Panic)`. Ordinary provider errors, including Interrupted,
are recorded without automatic retries and are not labeled contract violations.
Any finding makes complete/passed false. Global stops distinguish
`Budget(path, CheckBudget)`, `IssueLimit` and `ProviderPanic`; CheckBudget names
the exhausted Nodes, Depth, DirectoryEntries, FileBytes, TotalReadBytes,
ProviderCalls or PathBytes limit. The first stop reason is retained.

Every successful open is closed exactly once by the checker, including an
unexpectedly successful invalid-path open. Close is attempted even after a
budget stop or catchable panic. Each close has an independent panic boundary,
so a panicking cleanup does not prevent other open handles being closed.
Cleanup errors are retained alongside earlier issues while report space remains;
report exhaustion never prevents cleanup. Fatal runtime termination and failures
outside the language's catchable panic mechanism are not made recoverable.

`check_seek`, `check_read_at` and `check_links` take the same expected tree and
limits and return the same report. Each runs the basic tree checks and its
additional protocol checks with one shared budget. Capabilities are required
statically: check_seek needs handles implementing ReadDirFile and Seek;
check_read_at needs ReadDirFile and ReadAt; check_links needs LinkFileSystem
and ReadDirFile handles. No dynamic interface discovery is used.

Seek checking compares a bounded sequential snapshot with reads at Start,
Current and End positions, verifies returned positions, and probes negative
or overflowing arithmetic. It does not require seeking beyond EOF or preserving
all stream state after an error. Absolute repositioning is used after invalid
seeks. An unsupported End operation is an ordinary provider error that leaves
the check incomplete, not a contract violation.

ReadAt checking compares confirmed prefixes with the sequential snapshot,
requires full counts on success and bounded progress on errors, rejects
accepted invalid ranges or reported progress for them, and tests EOF, empty
buffers and interleaved sequential position. Partial UnexpectedEof at the actual
snapshot boundary is expected; other errors remain provider errors, even when
they report a full transfer. Confirmed prefix counts on errors consume the shared
read budget. Separate fresh handles probe cursor preservation for successful
and partial-EOF reads, including one-byte files. Empty positional reads on
nonempty files also get fresh handles, so moving a cursor beyond an already
observed EOF cannot hide a fault. For a truly empty file, Read and ReadAt alone
cannot distinguish numerical positions at or past EOF; no Seek capability is
silently assumed. Invalid-range probes expect InvalidInput with zero progress;
other error causes remain provider errors rather than passing unnoticed.
The checker does not impose
Cursor-specific rules about untouched destination tails on all providers.

Link checking repeats raw read-link and non-following metadata calls, comparing
them with retained link entry metadata. It also probes invalid input names for
both methods. It neither cleans nor validates target strings and does not open
or recursively traverse child links. Absolute, empty, escaping, cyclic and
dangling targets are not inherently protocol violations; following-resolution
policy is outside this check.

`check_helpers` has the same signature and static requirements as check_fs.
It runs the basic checks and compares complete node snapshots with `stat_from`,
`read_file_from` and `read_dir_from`, both directly and through `SubFs`.
Files use their parent as the SubFs root; directories use themselves and query
`.`. Child symlinks remain unfollowed. Helper directory results must be sorted,
but duplicate entries retain their multiplicity. Failed or incomplete file-read
or directory-enumeration baselines are not used to infer helper mismatches.

An observed provider charges every underlying helper operation against the
same report budgets, including entry accessors and all confirmed read bytes.
Cached entry values do not cause another underlying call. Helpers do not get
a separate allowance, and cleanup remains exempt. Ordinary errors retain their
provider classification; a primary error and an independent cleanup error are
both retained, while a close-only error is reported once. Budget stops and
catchable panics discard the helper result without reporting internal stop
sentinels as provider errors. These comparisons test stable provider composition,
not host confinement or atomic snapshots.

The public types and functions use existing syntax; there is no new grammar
production.

```goml
use std::fs;
use std::testing::fs as memory;

fn inspect_tree(source: memory::MemoryFS) -> Result[memory::CheckReport, fs::Error] {
    memory::check_helpers(source, memory::ExpectedTree::Contains(Vec::new()), memory::CheckLimits {
        max_nodes: 256, max_depth: 16, max_directory_entries: 128,
        max_file_bytes: 65536, max_total_read_bytes: 1048576,
        max_provider_calls: 100000, max_issues: 64, max_path_bytes: 65536,
    })
}
```

An Ok result means the checker inputs were valid, not that the provider passed;
inspect the returned report's `passed()`, issues and stop reason.

### Resource scopes

Import `std::resource` for explicit cleanup of resources managed by libraries. `with_cleanup(action, cleanup)` executes the action and then cleanup on normal return, including an action returning `Err`. `scope(action)` passes a `Scope[C]` into the action, then closes it. `Scope::register` accepts a `() -> Result[(), C]` cleanup callback. Closing runs all registered callbacks in reverse registration order, continues after cleanup errors, and runs each callback once.

`Scope::new`, `register`, `is_closed`, and `close` support concurrent callers. Registration returns `false` once closing begins; the caller retains responsibility for a rejected callback's resource. Concurrent close calls wait for the first close and receive copies of its error vector; subsequent calls return the same errors without repeating cleanup. Callbacks run outside the scope's state gate and may inspect or attempt registration, but must not synchronously close their own scope.

`ScopeError[E, C]` retains `action: Option[E]` and `cleanup: Vec[C]`; cleanup failures do not hide the action error. `finish(action_result, cleanup_errors)` combines already-completed operations using the same rule. Successful actions produce `Ok` only when cleanup also succeeds. `ScopeError` implements `ToString` when both error types do. `io::with_resource(value, action)` uses the same machinery for any `io::Close` value.

```goml
use std::resource;

fn run_job() -> Result[isize, resource::ScopeError[string, string]] {
    resource::with_cleanup(|| Result::Ok(42), || Result::Ok(()))
}
```

These helpers do not implement destructors or implicit panic recovery. Cleanup runs on normal control flow, `Result` errors, and panic unwinding. A cleanup that panics does not prevent older registered cleanups from running, and the scope still publishes its closed state to waiting closers. The newest panic propagates instead of becoming a `ScopeError`; ordinary action and cleanup errors retain their existing combined-result behavior. Abrupt process termination and fatal runtime failures can bypass cleanup. A scope that escapes its callback remains closed afterward.

### Composable I/O

Import `std::io::{Read, Write, BufRead, Seek, Close}` to compose streams through generic bounds. `Read::read(MutSlice[byte])` and `Write::write(Slice[byte])` return the transferred byte count; short transfers are normal and zero from a nonempty read means EOF. Counts outside the supplied buffer are invalid. `read_exact` and `write_all` complete partial transfers and report `UnexpectedEof` or `WriteZero` when progress stops. `read_exact` and `write_all` return every provider error immediately, including `Interrupted`. Failed calls do not report their progress: retrying reads can lose already-consumed input, while replaying writes can duplicate side effects, including with MultiWriter or OffsetWriter. This intentionally changes the earlier generic interruption-retry behavior without changing signatures; concrete providers that know an interruption made no progress can retry internally. Errors may follow partial progress and do not roll back data.

`Seek::seek(from: SeekFrom) -> Result[isize, Error]` provides checked byte
positioning independently of `Read` or `Write`. `Start(offset)` is relative to
the logical stream start, `Current(delta)` to its current position, and
`End(delta)` to its logical length when that length is available. Success
returns the new nonnegative position. Position arithmetic rejects negative or
overflowing results with `InvalidInput`; offsets and positions are machine-sized
signed integers. `seek(SeekFrom::Current(0))` reports the current position but
still applies the concrete stream's seek side effects. The trait does not
promise arbitrary beyond-EOF positioning or restoration of all state on error.

| Seek implementation | Valid positions and side effects |
| --- | --- |
| `StringReader` | Allows positions through maximum `isize`, including beyond EOF and inside UTF-8 scalars; every attempt clears character rollback, even if it fails. |
| `Cursor` | Restricted to its current byte length, not its configured write limit; never grows storage or creates holes. End observes the current length after writes. |
| `SectionReader[R: ReadAt]` | Restricted to the declared section length, with all positions relative to the section. Success clears a pending read error; failure preserves it. |
| `OffsetWriter[W: WriteAt]` | Start/Current are relative to its base and limited so base plus position remains representable. End returns `Unsupported` because `WriteAt` has no length contract. |

These implementations share positions through existing handle aliases, perform
no reads or writes while seeking, and neither flush nor close their providers.
Invalid positioning leaves the position unchanged. A successful OffsetWriter
seek does not guarantee its provider accepts a subsequent write at that offset.
Existing `set_position` and StringReader's inherent `seek` remain available.
Buffered wrappers do not implement Seek: their prefetch/drain state needs a
separate positioning policy. The new trait composes through existing syntax,
for example `fn rewind[S: io::Seek](stream: S) -> Result[isize, io::Error] {
stream.seek(io::SeekFrom::Start(0)) }`; it adds no grammar or runtime primitive.

`Read::read_to_end(limit)` returns `bytes::Bytes`; `read_to_string(limit)` additionally validates UTF-8. Limits are nonnegative byte counts. Both return every provider error immediately, including `Interrupted`, without retrying or returning partial output. Both reject oversized input, consuming at most one extra byte to detect overflow; use `Take::new(reader, limit)` when bytes beyond a limit must remain unread. `io::copy(reader, writer)` copies until EOF and returns a `u64` count; it does not flush or close either stream.

`io::read_at_least(reader, buffer: MutSlice[byte], minimum: isize)` returns
`Result[isize, TransferError]`, preserving confirmed progress on failure. It rejects
negative minima or minima exceeding the view length with `InvalidInput` before
I/O; zero succeeds without calling the reader. Each read receives the entire
remaining view, so success may exceed the minimum, but never the view length.
Pass `buffer.len()` as the minimum for an exact read with progress reporting.
EOF before the minimum returns `UnexpectedEof`, even with zero confirmed bytes;
an impossible provider count returns `InvalidData`. Untouched view suffixes and
out-of-view storage are not initialized by this function.

Every provider error, including `Interrupted`, is returned immediately with the
sum of preceding valid successful counts. A failed call can mutate the supplied
buffer or consume input without reporting how much; those effects are not counted
or rolled back, and the count is not a universally safe resume offset. Both this
function and `Read::read_exact` return interruptions immediately.
It allocates no scratch buffer, flushes/closes nothing and uses
ordinary generic/function syntax. For example, a one-byte view and minimum one
read one byte from `io::StringReader::new("abc")` without consuming the remainder.

`io::NoopCloser::new(reader)` wraps any `Read` value with a `Close` implementation
that always succeeds without touching the underlying stream. Reading forwards
the original call, count and error unchanged, including empty buffers; the
wrapper does not validate provider counts at this boundary. Repeated close calls
do not disable later reads or close a provider that also implements `Close`.
Use it with `io::with_resource` when the scope borrows responsibility for reading
but must not take responsibility for closing. Provider aliasing and concurrency
rules remain unchanged. The wrapper supplies `Read` and `Close`, not automatic
forwarding of positional, buffered or optimized transfer interfaces.

`io::copy_buffer(reader, writer, buffer: MutSlice[byte]) -> Result[u64, io::Error]` uses the supplied nonempty scratch view instead of allocating its own buffer. An empty buffer returns `InvalidInput` before accessing either stream. Reads validate counts and return every error including `Interrupted`; writes use the destination's `write_all` implementation, including any adapter-specific no-retry policy. The scratch view is overwritten and must not alias live source/destination storage or be used concurrently. No optimized transfer protocol bypasses this buffer. `copy` now uses this same engine with an 8 KiB buffer, preserving its existing behavior.

`io::copy_n(reader, writer, count: isize) -> Result[u64, io::Error]` copies exactly the requested number of bytes through a bounded reader. Negative counts fail before accessing either stream; zero succeeds without reading, writing or allocating a copy buffer. It never reads beyond the requested count. An earlier EOF returns `UnexpectedEof`; already copied bytes remain in the destination. These legacy copy variants stop on other errors, do not roll back side effects and return no progress count on failure. A destination failure may occur after input was consumed. None flushes or closes streams.

`io::copy_progress(reader, writer)` and
`copy_buffer_progress(reader, writer, scratch: MutSlice[byte])` return
`Result[u64, CopyError]`. Success reports the copied byte count. Failure exposes
`CopyError::read_count()`, `written_count()` and `error()`: separate confirmed
input and output counts plus the original provider error. Read-ahead means the
input count can exceed the output count. Failed calls may consume input, modify
scratch or produce output without reporting progress; those effects are not
counted or rolled back. Neither count is a universally safe recovery offset.

These functions call `read` and `write` directly, complete successful short
transfers and never retry errors, including `Interrupted`. They do not invoke a
custom `write_all` override, flush or close either stream. Invalid successful
counts yield `InvalidData`; a zero write yields `WriteZero`. Empty scratch is
rejected with `InvalidInput` and zero counts before I/O. Scratch must not alias
live source/destination storage or be accessed concurrently. `copy_progress`
allocates an 8 KiB scratch buffer. Reads are clipped to the remaining `u64`
counter range; once the maximum count is reached the next iteration reports
`InvalidData` before I/O, even if a further read would have discovered EOF.
These opt-in APIs leave the existing copy signatures unchanged and add no syntax.

`io::Discard::new()` creates a stateless `Write` implementation that accepts the entire supplied buffer, including empty buffers, without retaining data; `flush` succeeds without work. For example, `io::copy_n(reader, io::Discard::new(), 32)` skips exactly 32 bytes or reports premature EOF. These APIs use existing function, trait and generic syntax.

`io::ReadAt::read_at(output: MutSlice[byte], offset: isize)` and `io::WriteAt::write_at(input: Slice[byte], offset: isize)` perform positional transfers without changing the sequential position. Their result is `Result[isize, io::TransferError]`: success reports the complete buffer length; failure can report a transferred prefix together with its cause. `TransferError::new(transferred: usize, error: io::Error)`, `.transferred()` and `.error()` preserve both values without changing existing `Read`/`Write` error signatures. Implementations must reject negative offsets and offset-plus-length overflow, and must report short transfers as errors. They must not report progress beyond the supplied buffer.

`read_exact_at` and `write_all_at` validate the range before calling the implementation once. They validate returned counts and error progress, preserve valid underlying errors (including `Interrupted`), and do not retry. A short successful read is normalized to `UnexpectedEof`; a short successful write or impossible count is `InvalidData`. Impossible progress is reported as zero because the provider's count cannot be trusted; this does not imply no mutation occurred. Errors never roll back already transferred bytes or other provider side effects. Implementations must not retain the supplied buffer after returning. Unlike Go's ReaderAt/WriterAt concurrency contracts, these traits alone do not promise concurrency safety: callers must follow the concrete provider's synchronization requirements. The offset type is GoML's signed machine word, not an independent 64-bit file-offset type.

`Cursor` also implements `ReadAt` and `WriteAt`. A nonempty read extending past the end copies the available prefix and returns `UnexpectedEof` with its count, leaving the remaining destination untouched. An empty read at any nonnegative offset succeeds. Positional writes may overwrite or append at the current length, but cannot create gaps; validation failures leave storage unchanged. Both sequential and positional writes snapshot their input before mutation, including overlapping views obtained from `fill_buf`. Positional operations preserve `position()`, including on errors. For example, with `use std::io::ReadAt`, `cursor.read_exact_at(output.as_mut_slice(), 4)` fills a buffer starting at byte four without seeking the cursor. These APIs use ordinary traits and methods and add no grammar.

`BufRead::fill_buf` exposes the currently buffered bytes and `consume(amount)` checks the available range. `read_until(delimiter, limit)` includes a found delimiter; `read_line(limit)` includes the newline, validates UTF-8, and returns `None` only at EOF with no bytes. A final unterminated line may exactly fill the limit. Treat a buffer view as valid only until the next operation on that reader.

`BufRead::discard(amount: isize) -> Result[isize, TransferError]` skips exactly
the requested number of bytes across buffer fills. A negative amount fails with
`InvalidInput`; zero succeeds without calling `fill_buf` or `consume`. Existing
buffered bytes are consumed first; insufficient input returns `UnexpectedEof`.
The error's transferred count includes only successful `consume` calls. Every
fill or consume error is returned immediately without retry, preserving the
original error; a failed call's unreported side effects are not included.
The method allocates no scratch storage and flushes/closes nothing. A concrete
buffered provider may read ahead, so discarded bytes need not equal source
position advancement. Import `io::BufRead` to use the default method.

`Cursor::new(bytes)` copies its initial data and shares position among handle copies. It implements `Read`, `Write`, and `BufRead`; writes overwrite or append, `set_position` accepts offsets through the current length, and `bytes()` returns a copy. `Cursor::with_limit(bytes, max_bytes)` checks the initial length before copying and bounds all subsequent writes by total buffer length, not cumulative bytes written. Negative limits return `InvalidInput`; excessive initial data or writes return `InvalidData`. Rejected writes preserve data and position and report zero positional progress, including when the input aliases `fill_buf`. Empty writes and overwrites within the limit remain valid. `limit()` returns the configured limit; `new` uses maximum `isize`. The limit excludes snapshots, input copies and allocation overhead and does not recover allocation failure.

`BufReader::new(reader)` and `BufWriter::new(writer)` use 8 KiB buffers; `with_capacity(stream, size)` rejects nonpositive sizes. Buffered readers return every provider error immediately, including `Interrupted`, without retrying the failed read. Buffered writers require explicit `flush`; on failure they retain the suffix not confirmed by successful write counts, which is not necessarily an unsent suffix. Draining does not retry errors, including `Interrupted`. Retrying a failed flush can duplicate unknown side effects of the failing call; recovery requires knowledge of the concrete writer. A buffered wrapper implements `Close` when the underlying stream does; writer close flushes first. These in-memory adapters require callers to serialize shared access and have no finalizers.

`BufReader::peek(amount: isize) -> Result[Slice[byte], PeekError]` returns exactly
that many buffered bytes without consuming them. Requests below zero or above
`capacity()` fail with `InvalidInput` before I/O or buffer compaction; zero
returns an empty view without I/O. A valid request can compact unread bytes and
perform several short reads, using the remaining capacity and possibly reading
ahead. It never grows the buffer. EOF before the requested size returns
`UnexpectedEof`; provider errors, including `Interrupted`, are preserved without
retry, and invalid counts yield `InvalidData`.

`PeekError::available()` exposes the currently buffered confirmed prefix and
`error()` exposes the cause. Failed-call mutations beyond that prefix are not
exposed, but input consumed by the failed call cannot be recovered. Neither EOF
nor errors are cached permanently: a later explicit peek can read again if more
bytes are needed. Success and error views borrow shared mutable buffer storage;
copy them before another operation or an alias modifies the reader. Peeking is
not a snapshot, a consume operation, or a promise that retrying a failure is safe.

`BufReader::read_byte()` returns `Result[Option[byte], Error]` and
`read_char()` returns `Result[Option[(char, isize)], Error]`, where the tuple
contains a Unicode scalar and its original byte width. `None` means EOF without
remaining bytes. A malformed or EOF-truncated UTF-8 sequence consumes one byte
and returns U+FFFD with width one; valid U+FFFD consumes its three-byte encoding.
Character reads work even with configured capacity one. The physical buffer
holds at least four bytes for scalar assembly, but `capacity()` retains the
configured value and each provider read is limited to that value. Thus
`buffered_len()` and `fill_buf()` may expose up to `max(capacity(), 4)` bytes
after a character operation. `peek` still rejects requests above `capacity()`.
This fixed scalar storage does not grow with input or change writer capacity.

Character reads return provider errors immediately, including `Interrupted`,
without consuming any previously confirmed prefix of the current character.
That prefix stays buffered for every subsequent byte, bulk, peek or discard
operation. Bytes modified or consumed by the failed provider call are unknown
and excluded from the buffered prefix; retries cannot reconstruct lost input.
EOF and errors are not sticky. These methods neither flush nor close the source.

`unread_byte()` and `unread_char()` return `Result[(), Error]` and permit one
rollback after a qualifying successful read. Byte rollback is available after
`read_byte`, `read_char`, or a nonempty successful bulk `Read::read`; it restores
only the last byte. Character rollback is available only after `read_char` and
restores the original encoded bytes, including a single malformed byte rather
than re-encoding U+FFFD. Rollback changes only the shared logical buffer offset,
does not seek the provider, and requires no I/O. Repeated rollback without a new
qualifying read, or the wrong rollback kind, returns `InvalidInput`.

Every rollback attempt clears both rollback permissions, including a failed
attempt. `peek`, `fill_buf`, `consume`, and `reset` attempts also clear them,
even for zero or invalid arguments; empty bulk reads, EOF and read errors clear
them too. `capacity()` and `buffered_len()` are observational and preserve them.
Default helpers follow the operations they actually perform: `discard(0)` and
`read_exact` on an empty view do not call the reader and leave permissions
unchanged. Handle aliases share permissions and offset. A copied byte snapshot
survives later operations, but a borrowed buffer view does not. For example,
after `reader.read_char()` returns `Some(('界', 3))`, `reader.unread_char()` makes
the same three original bytes available again. These APIs use existing syntax.

`BufReader::read_fragment(delimiter: byte)` returns
`Result[ReadFragment, FragmentError]`, consuming a bounded piece of the logical
input. `ReadFragment::bytes()` is a borrowed view, `consumed()` reports the
confirmed raw bytes consumed by this call, and `end()` identifies the boundary:

| `FragmentEnd` | Meaning |
| --- | --- |
| `Delimiter` | The requested delimiter was found and is included in the raw fragment. A delimiter at the capacity boundary takes precedence over BufferFull. |
| `BufferFull` | The configured capacity was reached without a delimiter. No extra read probes for EOF; continue with another call. |
| `Eof` | A provider read returned zero before the buffer filled. The fragment can be nonempty or empty. |

Each raw fragment consumes at most `capacity()` bytes, even when a preceding
character operation left more bytes buffered. Unconsumed cached bytes remain
available. Reads append only within the configured capacity and scan new bytes
once; no whole-input allocation or growing fragment buffer is used. EOF is not
cached permanently, so a later explicit call can observe a source that resumes.

`FragmentError::bytes()` exposes a confirmed partial fragment that has already
been consumed, unlike `PeekError::available()`. `consumed()` reports its length
and `error()` preserves the provider cause; invalid successful counts become
`InvalidData`. All errors, including `Interrupted`, stop immediately. Effects of
the failed call are not included, and retrying cannot restore input lost by that
call. Both success and failure views share the reader's storage and are valid
only until another reader operation; copy them for retention. Aliases share
consumption. Fragment reads always clear byte/character rollback permissions,
including on failure, and do not enable rollback after returning.

`read_line_fragment()` has the same result type, using LF as delimiter and
removing that LF and an immediately preceding CR from a `Delimiter` payload.
`consumed()` still counts the removed wire bytes. It does not validate UTF-8.
`BufferFull` means a continued line, not a complete line. A trailing CR at a full
boundary is retained for the next fragment so a split CRLF is recognized without
dropping a bare CR. With capacity one, a leading CR uses the fixed scalar storage
for one-byte lookahead: CRLF yields an empty `Delimiter` payload with two consumed
bytes; CR followed by another byte yields a one-byte `BufferFull` payload and
retains the next byte. CR at EOF is returned as data; a lookahead error returns
the confirmed CR with its cause. Each underlying read remains capacity-bounded.
No `BufferFull` result is empty. An exact-capacity line can therefore be followed
by an empty `Delimiter` or `Eof` result. Unterminated tails retain bare CR bytes.
The existing whole-line `read_line(limit)` contract is unchanged.

Both wrappers expose `capacity()` and `reset(replacement)`; the replacement has
the same concrete stream type. Capacity is the configured logical buffer size,
not allocator capacity. Reset reuses storage, clears buffered state and switches
all existing aliases to the replacement. It performs no read, write, flush or
close on either stream. Reader reset discards unread prefetched bytes; writer
reset deliberately discards pending output, including an unconfirmed suffix after
a failed drain. Flush explicitly before reset if that output must be preserved,
subject to the failure caveat above. Existing reader views become invalid for
further use after reset; copy them first if they must survive.

`BufWriter::available()` reports `capacity() - buffered_len()`, the logical spare
space. After a partial drain failure it is not a promise that the next write
avoids underlying I/O: the wrapper drains its retained suffix first. Reset clears
that drain offset as well as the data. A later close targets the replacement,
not the previous stream. These operations require serialized access, do not
recover allocation failure, and add no grammar or runtime primitive.

`BufWriter::write_byte(value)` returns `Result[(), Error]` after accepting one
byte. `write_char(value)` and `write_string(value)` return
`Result[isize, TransferError]`, completing successful short writes and reporting
the number of new input bytes accepted into the buffered writer. String writes
use scratch storage of at most 4096 bytes rather than copying the whole input;
character writes encode at most four bytes. Empty strings perform no I/O and do
not drain existing output. Errors stop immediately without retry; the transferred
count excludes the failing call and any output buffered before this invocation.
Partial UTF-8 output is possible when an error interrupts a character or string;
these operations are not atomic and never flush or close the provider.

`BufWriter::read_from[R: Read](reader: R) -> Result[u64, CopyError]` composes
`copy_progress(reader, self)`, preserving separate confirmed input and buffered
acceptance counts. EOF does not flush the writer. Acceptance is not confirmed
delivery to the underlying sink: an operation may drain old pending output,
retain new output, or fail after unknown sink side effects. Existing partial
drain, explicit flush, alias and reset contracts still apply. The method adds no
optimized transfer bypass, automatic close, retry or recovery guarantee.

`io::pipe() -> (PipeReader, PipeWriter)` creates a synchronous in-memory byte
stream using existing channels. `PipeReader` implements `Read` and `Close`;
`PipeWriter` implements `Write` and `Close`. Unlike the mutable memory adapters,
copied pipe handles support concurrent calls. Writers publish one packet at a
time without interleaving; readers may consume it in several short reads. A
nonempty write succeeds only after all its bytes have been copied into reader
buffers. Input must remain unchanged until the write returns. There is no FIFO
or fairness guarantee and no background task or additional runtime primitive.
Concurrent operations must not use overlapping mutable buffers; synchronizing
pipe state does not synchronize caller-owned storage outside the operation.

`PipeWriter::write_progress(input)` returns `Result[isize, TransferError]` and
preserves the confirmed prefix if a close interrupts a write. Ordinary `write`
returns the same error without the prefix. Each endpoint offers
`close_error(error)`; repeated closes preserve that endpoint's first cause.
Closing never waits for partner progress and wakes blocked operations. A local
closed reader or writer returns `BrokenPipe`, taking priority over the other
endpoint's cause. Otherwise a reader sees EOF after normal writer close, or the
writer's custom error; a writer sees the reader's custom error, or `BrokenPipe`
after normal reader close. Closing the writer itself aborts its pending write
with `BrokenPipe`, even when its custom cause is supplied to readers. A packet's
completed reply cannot be replaced by a later close.

Open direct empty reads and writes return zero immediately without publishing a
packet or signalling EOF to a waiting reader. Closed-state checks precede these
empty operations. The inherited `flush` is a no-op, and generic
`write_all(empty)` does not call `write`, even after close. Pipe operations have
no cancellation-token parameter: cancellation cleanup must close an endpoint
to wake blocked tasks. Read and write must run concurrently for nonempty input:

```gom
use std::io;
use io::{Read, Write, Close};
use std::task;

task::scope(|scope| {
    let (reader, writer) = io::pipe();
    defer { let _ = reader.close(); let _ = writer.close(); };
    let producer = scope.spawn(|_| {
        defer { let _ = writer.close(); };
        writer.write_all(b"hello".as_slice())
    });
    let received = reader.read_to_end(16);
    let sent = producer.join();
});
```

`Scanner[R: Read]::new(reader, max_buffer)` creates a bounded line scanner and rejects nonpositive limits. `.next() -> Result[Option[bytes::Bytes], Error]` returns an independently owned token, None after successful termination, or a sticky error retained by `.failure()`. Calls after termination do not read again; handle copies share state and require serialized access. The scanner never closes the source. Default `scan_lines` strips LF and an immediately preceding CR, emits empty lines, and emits a nonempty final unterminated line after stripping its trailing CR. A final newline does not create an extra token. It preserves arbitrary byte content; callers bring `std::utf8::BytesUtf8` into scope to explicitly validate with `token.to_string_utf8()`. `scan_bytes` emits individual bytes, including invalid UTF-8.

`Scanner::with_split(reader, max_buffer, split)` accepts a `(Slice[byte], bool) -> Result[ScanStep, Error]` callback. The boolean means confirmed EOF. `ScanStep::More(advance)` discards the specified prefix and retries, or requests more input when advance is zero. `Token(advance, start, end)` returns the indicated byte range and consumes advance bytes: `0 <= start <= end <= advance <= input.len()` and advance must be positive. `Final(Some((start, end)))` emits one final token, possibly empty, and stops; Final(None) stops without a token. Final may stop before EOF and discards any buffered remainder. Invalid callback ranges and non-progressing tokens return InvalidData instead of looping. More(0) with nonempty input at EOF returns UnexpectedEof; with empty input it ends normally. Callback errors remain terminal and are not retried.

The limit bounds pending input bytes, including delimiters, not just returned token size or total process memory. Reads use scratch storage of at most 4096 bytes. When pending input reaches the exact limit and the callback still needs data, one additional byte may be consumed to distinguish EOF from overflow: EOF permits final token delivery, while an additional byte yields terminal InvalidData. Scanning may read ahead, so callers must not assume the source position is the end of the last returned token. All source errors, including Interrupted, are preserved without retry; buffered incomplete data is discarded on failure and prior tokens remain valid. Callbacks and returned token ranges use byte offsets. These APIs introduce no new grammar or native backend.

`scan_runes` waits for complete UTF-8 scalars across reads and returns each scalar as a valid UTF-8 token. Every malformed byte, including each byte of a truncated EOF sequence, yields U+FFFD; a literal U+FFFD is indistinguishable from a replacement token. `scan_words` splits on the pinned Unicode White_Space property, skips leading/trailing whitespace and never emits empty words. It preserves the original bytes of each word, including malformed UTF-8, and waits for incomplete scalars before deciding whether they delimit a word. This is whitespace tokenization, not linguistic word-boundary segmentation.

`ScanStep::Emit(advance, token)` supports transformed tokens such as rune replacements. Advance must be positive and no greater than the pending input length. Scanner copies the supplied token, so later mutation by the callback cannot change a returned token. The pending-input limit does not bound callback-created output; a replacement rune can occupy three bytes while consuming one input byte. Custom callbacks are responsible for their own output/resource limits.

`SectionReader::new(reader, start, length)` returns `Result[SectionReader[R], io::Error]` for any `R: ReadAt`. It rejects negative bounds and end-offset overflow without accessing the reader. It exposes a live logical subrange, not a snapshot or a security sandbox: the declared length need not exist in the underlying input yet. `size()` returns that declared length; `position()` and checked `set_position()` use section-relative offsets from zero through the length. Handle copies share position and pending errors; independently constructed sections do not. No operation moves the underlying reader's sequential position, closes it or reads outside the declared section. Nested sections are supported.

The section implements `ReadAt` using relative offsets, clips the provider's destination to the section boundary and validates provider counts. A request extending beyond the section reports `UnexpectedEof` with the available prefix count. Empty positional reads at nonnegative offsets succeed without calling the provider. Sequential `Read` instead returns zero at the logical section end. Because `Read` cannot return both a count and an error, an underlying failure with valid positive progress advances the section position and returns that count; the next nonempty sequential read returns the saved error once, without invoking the provider. Empty reads preserve this pending error. A successful `set_position` clears it; an invalid seek and positional reads leave it unchanged. Further reads after delivery can retry the provider at the advanced position. All shared section access requires caller synchronization. For example, `io::SectionReader::new(cursor, 16, 32)` exposes at most bytes 16 through 47 using the existing generic/type/method syntax.

`OffsetWriter::new(writer, base)` returns `Result[OffsetWriter[W], io::Error]` for `W: WriteAt`, rejecting a negative base. It implements `WriteAt` by translating relative offsets to `base + offset`, validating both additions and the provider's counts. Positional writes leave its sequential `position()` unchanged. Checked `set_position(relative)` accepts nonnegative positions whose absolute offset is representable; it does not check the underlying destination's length or create a gap. The provider determines whether a later write at that position is supported. Handle copies share position; independently constructed and nested offset writers have separate relative positions.

Its sequential `Write::write` advances position by valid reported progress on both success and failure, then returns the underlying error immediately. Use positional `write_at` when the error's progress count must be returned directly. This adapter's `write_all` makes one validated positional transfer and does not retry `Interrupted`: a failed transfer may already have changed the destination and must not be replayed. An error accompanying a full-length transfer is still an error. Invalid provider progress is not trusted and does not advance position; actual provider side effects cannot be undone. Empty writes are forwarded to the provider and can fail. `flush` is a no-op because `WriteAt` has no flushing contract; explicitly flush the concrete destination when needed. The adapter neither closes the provider nor changes its sequential position, and callers must synchronize shared access. For example, `io::OffsetWriter::new(cursor, 16)` writes starting at byte 16 using existing generic/type/method syntax.

`MultiReader::new(readers: Slice[R])` for `R: Read` concatenates readers in order. It snapshots the handle list, not their data or positions. Empty reads return zero without accessing any reader; a nonempty zero read exhausts the current reader permanently and releases that stored handle. It skips exhausted readers until one returns positive progress, an error, or the list ends. Counts are validated; errors do not advance to the next reader, so a later call retries the same reader. Copies share traversal state, and exhaustion is permanent. An empty list is already exhausted. The adapter does not close readers and callers must synchronize shared use.

`MultiWriter::new(writers: Slice[W])` for `W: Write` snapshots a destination handle list and writes the complete input to each destination in order, including empty inputs. It stops at the first error or short write; zero progress on a nonempty input is `WriteZero`, other short or invalid counts are `InvalidData`. Earlier destinations may already contain the full input and the failing destination may contain a prefix; there is no rollback or aggregate progress count. `write_all` calls this broadcast once and never retries `Interrupted`, since replaying would duplicate earlier writes. `flush` visits destinations in order and stops at the first failure. An empty destination list accepts the entire input. Repeated handles are deliberately visited repeatedly, and the adapter does not close destinations.

Both constructors use one concrete element type per list and can nest; heterogeneous sources can use an application-defined common `Read`/`Write` implementation. Nesting is not dynamically flattened. For example, with `use std::io::Write`, `io::MultiWriter::new(Vec::from_array([left, right]).as_slice()).write_all(data.as_slice())` broadcasts to two writers of the same type. No new syntax or runtime dispatch mechanism is introduced.

`TeeReader::new(reader, writer)` implements `Read` for `R: Read, W: Write`. Each positive source read is synchronously passed to the mirror's `write_all` before returning success. Short writes and interruption are handled by that writer's implementation, including its no-retry overrides. The tee does not allocate a staging buffer, retain the destination view, flush or close either stream. Empty reads and source EOF perform no mirror operation; source errors and invalid counts are returned without calling the mirror or poisoning the tee.

A mirror failure is terminal for that tee and all its copies. The original error is available through `failure() -> Option[io::Error]`; `failed_read_count() -> isize` records how many source bytes the failing call placed in the caller's buffer, not how many reached the mirror. Previously completed reads are not included. Before failure it is zero. The failing call returns an error immediately, while source consumption, buffer changes and partial mirror writes remain. All later reads, including empty reads, return the failure without touching either stream. There is no reset or replay operation. This intentionally differs from Go's count-plus-error TeeReader API: GoML's `Read` cannot return both, so callers can inspect the explicit failure state when partial input matters.

If the mirror's final error is `Interrupted`, reads report `Other` with a terminal-mirror message while `failure()` retains the exact original error. This historical terminal-error mapping is retained for compatibility; generic `read_exact` and `copy` now return interruptions without retrying. Other mirror errors keep their original kind and details. Tee handles require caller synchronization, and source, mirror and caller buffer must not alias in a way that mutates unread source data. For example, `io::TeeReader::new(source, io::Discard::new())` forwards successful reads without retaining a copy, using existing trait and generic syntax.

`stdin()` implements `Read`, `stdout()` and `stderr()` implement `Write`, and `TcpStream`, `net::tls::TlsStream`, and Linux `fd::Fd` implement `Read`, `Write`, and `Close`. Existing inherent socket methods remain available. For cancellation, call the socket's context-aware methods directly; the generic traits carry no context parameter.

```goml
use std::bytes;
use std::io;
use std::io::{BufRead};

fn first_line(data: bytes::Bytes) -> Result[Option[string], io::Error] {
    let reader = io::BufReader::new(io::Cursor::new(data));
    reader.read_line(4096)
}

fn bounded_memory_output() -> Result[io::Cursor, io::Error] {
    io::Cursor::with_limit(bytes::Bytes::new(), 65536)
}
```

### Read-only string streams

`io::StringReader::new(value)` holds an immutable string without copying it into
a writable byte buffer. `size()` is the original byte length, `position()` the
current byte offset, and `len()`/`is_empty()` describe unread bytes. Handle copies
share position and `reset(value)`; reset replaces the string, returns to byte zero
and clears character rollback state. Shared access requires serialization.

The reader implements `Read` and `ReadAt`. Sequential reads return zero at EOF;
empty sequential reads succeed. Positional reads preserve both position and
character rollback state, follow the checked `ReadAt` range contract, and return
`TransferError` with the available prefix count on short reads. Empty positional
reads succeed at any nonnegative offset. Neither interface grants write access.

`seek(SeekFrom::Start(offset))`, `Current(delta)` and `End(delta)` return the new
absolute byte position. Nonnegative positions beyond EOF are allowed. Negative
or overflowing results return `InvalidInput` without moving the position; every
seek attempt invalidates character rollback. The inherent method and the
`io::Seek` implementation share this behavior; existing callers do not need to
import the trait to keep using the inherent method.

`read_byte()` returns `Option[byte]`; `read_char()` returns
`Option[(char, isize)]`, containing the scalar and consumed byte width. EOF is
`None`. Starting inside a UTF-8 scalar (after byte reads or seeking) yields the
replacement character and consumes one byte, as Go's strings.Reader does.
`unread_byte()` moves back one byte at any positive position, including beyond
EOF, and fails at zero. `unread_char()` rolls back the last successful character
read once. Byte reads, byte rollback, sequential reads (including empty/EOF),
seek attempts, reset and an EOF character read invalidate character rollback.
`write_to`, including an empty transfer, also invalidates rollback.
Metadata queries and positional reads do not. The empty/EOF sequential-read
invalidation rule is deliberate rather than promising all of Go's state-machine
quirks. These operations use existing string primitives and add no grammar.

For example, `io::StringReader::new("é界").read_char()` yields `Some(('é', 2))`;
seeking to byte 1 then reading a character yields `Some(('�', 1))`.

`reader.write_to(writer)` returns `Result[isize, io::TransferError]` and transfers
the unread suffix using a 4096-byte scratch buffer. Unlike generic read-then-write
copying, the reader advances only by successful, valid writer counts. It completes
short writes, returns all writer errors immediately (including `Interrupted`), rejects zero progress with `WriteZero`, and
rejects negative/oversized counts with `InvalidData`. Failure reports the confirmed
count for this call and leaves the unconfirmed suffix available for another call;
the position may be inside a UTF-8 scalar. Empty/EOF transfers do not call the
writer or change position. No implicit flush or close occurs. The reader and its
aliases must not be concurrently or reentrantly mutated by the writer callback.
Unknown side effects of a failing or invalid-count writer cannot be counted or
rolled back. Even valid MultiWriter or OffsetWriter adapters can have side effects
in a failed call that are not represented in this sequential Write result.
Resumption therefore requires provider-specific knowledge of that failed call,
not merely a zero confirmed count or an Interrupted error. Go Reader.WriteTo's
single-call short-write behavior is not promised: this method deliberately
completes short writes using the GoML stream conventions.

### Compiled multi-rule text replacement

`text::Replacer::new(rules: Slice[(string, string)], max_rules, max_rule_bytes)`
returns a compiled replacer or `ReplaceError`. Each pair is a literal old/new
rule, not a regular expression. Rules are snapshotted; later edits to the input
collection do not affect the replacer. Limits bound the rule count and the sum
of old/new UTF-8 byte lengths before trie construction; empty rules still count.
Compiled state is private and replacement calls do not mutate it.

`replacer.replace(input, max_bytes, max_work)` scans left to right. At each
position, the earliest supplied matching rule wins, not the longest rule.
Duplicate old patterns retain the first rule. Replacement text is emitted once
and is never rescanned. An empty rule can match once at a position before trying
nonempty rules at that same position; otherwise scanning advances by one scalar.
Empty patterns therefore preserve UTF-8 boundaries, unlike Go strings.Replacer's
byte-wise empty-pattern insertion. For example, the rule `("", "-")` maps `界🙂`
to `-界-🙂-` rather than inserting bytes inside either scalar's encoding.

The output limit counts UTF-8 bytes. Matching charges one unit per root lookup
and attempted trie byte transition, including failed transitions; reaching the
earliest possible rule stops that lookup. Exhaustion returns `WorkLimit`, not an
incorrect partial match. Output copying is separately bounded by max_bytes, not
charged as trie work. Each call uses independent output/work state and errors
return no partial output; the same compiled value remains reusable after failure.
There is no linear-time claim for adversarial overlapping patterns.

`replacer.chunks(input, max_bytes, max_work)` returns
`Result[FnIterator[Result[string, ReplaceError]], ReplaceError]`. Negative limits
fail during construction; matching and output-budget checks happen on demand.
Successful items are nonempty UTF-8 fragments in output order. Concatenating them
produces the same result as `replace`, which uses this same iterator internally.
Deletion rules produce no fragment but still consume matching work. Fragment
boundaries are not a stable API contract; a replacement can be one large fragment.
No complete output buffer is allocated by the iterator.

An iterator reports a matching/output error once, then remains exhausted. Already
yielded fragments are not rolled back. A fragment that would exceed the total
output budget is not yielded, even partially. Stopping iteration avoids scanning
the remaining input. Iterator aliases share progress and require serialized use;
separate `chunks` calls have independent state. The iterator retains its input
and compiled rules; laziness is not a retained-memory quota. For example, a
no-rule replacer on `a界b` with output budget 3 yields `a`, then an output-limit
error, then no more items. A fresh call may retry with a larger budget.

`ReplaceError` distinguishes negative limits, rule-count/rule-byte limits, output
and work limits, and an invalid input scalar boundary. Limits exclude caller-owned
inputs and general allocation overhead; they are not retained-memory quotas.
Normal inputs are valid GoML strings;
arbitrary byte rewriting belongs to bytes APIs. Constructor limit validation is
rule count then rule bytes, before checking actual rule sizes; replacement checks
output then work limits before scanning. For example, compile `[("ab", "x"),
("a", "y")]` with limits `(2, 5)` and replacing `aba` yields `xy` when output
and work budgets suffice. These APIs add no grammar forms or native backend.

### Streaming replacement output

The opt-in `std::text::stream` package composes replacement with `std::io::Write`;
importing `std::text` alone does not import I/O. Its
`write_replaced(replacer, input, writer, max_bytes, max_work)` returns the written
byte count or `ReplaceWriteError`. It consumes `Replacer::chunks` lazily and copies
fragments through a fixed 4096-byte scratch buffer, never a complete output
buffer. A large replacement fragment is written in bounded pieces. Short writes
are completed, all writer errors including `Interrupted` return immediately, zero progress becomes `WriteZero`, and
negative or oversized writer counts become `InvalidData`. No implicit flush or
close occurs; the caller owns those operations.

`ReplaceWriteError::Replacement(count, error)` preserves a `text::ReplaceError`;
`Write(count, error)` preserves an I/O error. `transferred()` returns the count
confirmed by successful valid write calls before failure, excluding any unknown
side effects of a writer that returns an error or violates its count contract.
Already written bytes are not rolled back and may end inside a UTF-8 scalar.
Replacement budgets retain their text semantics: a fragment exceeding the output
budget is rejected before any of its bytes reach the writer. I/O calls do not
consume matching work; this API is not a timeout or
cancellation policy. Retrying the whole function against the same destination
can duplicate earlier output; it is not a resumable write cursor.

For example, with `use std::text::stream` and `use std::io`,
`stream::write_replaced(replacer, "aba", io::Discard::new(), 64, 1024)` validates
and counts replacement output without retaining it. These are ordinary generic
functions and error variants, with no new syntax or native algorithm backend.

### Bounded text building

`StringBuilder::capacity()` reports its current byte capacity.
`reserve_checked(additional, max_bytes)` ensures space for at least
`len() + additional` bytes without changing content or length. It checks a negative
limit first (`InvalidLimit`), then a negative increment (`InvalidCount`), arithmetic
overflow (`LengthOverflow`) and the logical total-byte bound (`LimitExceeded`).
Validation failures also leave capacity unchanged. Existing larger capacity is
retained; allocator growth may exceed the requested amount. The bound is not a
strict limit on actual allocation, and `clear()` retains the existing capacity.
For example, an empty builder can reserve 16 additional bytes under a limit of 16
and still report length zero. Aliases observe the same reservation.

`write_bytes(bytes)` validates the complete byte value as UTF-8 before appending;
invalid input leaves contents, length and capacity unchanged. It copies accepted
bytes, so later source mutation does not change the builder. `finish()` returns
an immutable snapshot unaffected by subsequent writes or `clear()`. Unlike Go's
byte-oriented string builder, this text API does not accumulate invalid UTF-8
fragments across calls. Use `bytes::Builder` for arbitrary bytes or split encoded
scalars, then perform checked UTF-8 conversion when complete. Direct mutation of
public `values` bypasses these text checks and remains the caller's responsibility.

`text::clone_checked(value, max_bytes)` returns an equal string detached from the
source's backing storage, using a new byte vector and the existing immutable
string conversion boundary. Negative or insufficient output-byte limits fail
before copying. Use it when a small substring should not retain a large source.
Empty/short strings need not have globally unique addresses. The bound excludes
temporary allocations and allocator overhead and does not recover allocation
failure. For example, `text::clone_checked("界", 3)` succeeds and limit 2 fails.

`text::StringBuilder::write_string_checked(value, max_bytes)`,
`write_char_checked(character, max_bytes)` and `write_line_checked(value, max_bytes)`
append under a total output-byte limit, including existing content. All return
`Result[(), bytes::TransformError]`: negative limits produce `InvalidLimit`,
length arithmetic overflow produces `LengthOverflow`, and excessive total length
produces `LimitExceeded`. Validation precedes buffer growth and mutation; line
appends preflight both the supplied text and trailing LF as one operation. UTF-8
characters count by encoded bytes, not scalar count. Empty appends still reject
an already oversized builder.

The limit is supplied per call, not stored on the builder. Existing unchecked
methods, public `values` access and shared mutable aliases remain unchanged;
callers must preserve UTF-8 and serialize access. A failed call leaves the buffer
unchanged but does not roll back earlier successful calls. This is not an
allocation-failure recovery mechanism or a cap on separately formatted inputs.

Numeric and quoting append workflows compose existing conversion functions with
these checked writes, without a second conversion algorithm or Go append ABI.
For example, `builder.write_string_checked(number.to_string(), 4096)` preserves
the existing prefix and checks the final total size before appending. Use bounded
float/quote formatters with an appropriate remaining budget when their temporary
output must also be limited. These are ordinary methods with no grammar changes.

### Bounded quoted literals

`text::quote(value, max_bytes)` returns a double-quoted Go-style string literal;
`quote_char(value, max_bytes)` returns a single-quoted scalar literal. The
`quote_ascii` and `quote_char_ascii` variants escape every non-ASCII scalar.
The `quote_graphic` and `quote_char_graphic` variants additionally preserve Unicode
space-separator (Zs) characters such as nonbreaking and ideographic spaces.
All six functions return `Result[string, bytes::TransformError]`. The output
limit includes both delimiters; even an empty string needs two bytes. Negative
limits return `InvalidLimit`, excessive output returns `LimitExceeded`, and
length arithmetic is checked before allocating the final output buffer.

Quotes matching the delimiter and backslashes are escaped. Bell, backspace, tab,
LF, vertical tab, form feed and CR use `\a`, `\b`, `\t`, `\n`, `\v`, `\f`
and `\r`; remaining ASCII controls use two-digit `\x` escapes. Nonprintable
Unicode scalars use four-digit `\u` or eight-digit `\U` escapes with lowercase
hex digits. Normal quoting leaves Unicode 15.0.0 printable scalars unchanged;
Graphic quoting uses the same pinned Unicode version and leaves graphic scalars
unchanged, but still escapes quotes, backslashes, controls and non-graphic values.
ASCII quoting escapes all non-ASCII scalars, including printable ones. Inputs
are valid GoML strings/chars, so there is no invalid-byte or surrogate input mode.

`text::can_backquote(value)` applies Go's conservative single-line eligibility
policy for enclosing a string in backticks without changing its contents.
It rejects backticks, U+FEFF, DEL and
ASCII controls other than tab (including CR and LF). Other valid non-ASCII scalars
are permitted even when not graphic; this is not a terminal-safety predicate.
It scans without producing an output string. For example, `can_backquote("a\tb")`
is true, while `can_backquote("a\nb")` is false. The result concerns interchange
literals, not the GoML source grammar.

These APIs encode Go-style interchange literals, not JSON, shell commands,
HTML or GoML source syntax. In particular, the full set of Go escapes is not a
new set of accepted GoML source escapes. No parsing, normalization, locale rules
or security sanitization is implied. Output limits exclude existing inputs and
bounded per-scalar temporary storage; allocation failure is not recovered.

`text::unquote_bytes(literal, max_bytes)` decodes exactly one complete Go-style
double-quoted, single-quoted or backquoted literal into independent `bytes::Bytes`.
It rejects trailing text rather than parsing a prefix. Double/single quotes accept
the short escapes above, exactly two hexadecimal digits after `\x`, four after
`\u`, eight after `\U`, or exactly three octal digits. Octal values must fit a
byte; Unicode escapes must denote a valid scalar (no surrogates or values above
U+10FFFF). Escaping a quote is permitted only for the current delimiter.
Unescaped LF is rejected inside double/single quotes. Single quotes accept at
most one literal scalar or escape, including a byte escape; the empty `''` form
is accepted to match Go's Unquote behavior, not Go source character-literal syntax.
Backquoted content is literal except that CR bytes are discarded; LF is retained.

`text::quoted_prefix(input, max_bytes)` validates and returns the first complete
quoted literal at byte zero, including its delimiters and original escapes.
Trailing text is not inspected; advance by the returned string's byte length to
obtain the remainder. No leading whitespace is skipped. The limit counts source
bytes of the returned prefix, not decoded bytes: even `""` requires two bytes,
and raw CR bytes remain present and count toward the limit. Unlike unquoting,
this operation does not construct the decoded body or require byte escapes to
form UTF-8. It uses bounded per-escape scratch storage.

Errors use `UnquoteError`, with negative limits checked first. After validating
the opening delimiter, scanning stops at the budget; `OutputLimit(offset)` points
to the first element or closing delimiter that cannot fit. Escape validation may
inspect a constant-size escape past that limit before reporting its error. Syntax
errors use source byte offsets and no partial prefix is returned. This is a
Go-style interchange-literal parser, not a new GoML token or grammar production.
For example, `quoted_prefix("\"a\"tail", 3)` returns `Ok("\"a\"")`.

`text::unquote(literal, max_bytes)` additionally requires the complete decoded
result to be valid UTF-8 and returns a string. Thus byte escapes can reconstruct
multibyte UTF-8, but a lone `\xff` is available only through `unquote_bytes`.
Both complete-input decoders return `UnquoteError`: `InvalidLimit(value)` for negative limits,
`Syntax(offset)` for malformed syntax, `OutputLimit(offset)` when the next decoded
element would exceed the limit, or (for `unquote`) `InvalidUtf8(offset)`.
Syntax and output-limit offsets are zero-based input byte positions; malformed
escapes point at their backslash and a missing closing quote points at input end.
UTF-8 error offsets instead refer to the decoded byte sequence. Diagnostics do
not echo input contents. Negative limits take precedence; otherwise syntax and
output limits are checked in traversal order, and UTF-8 validation follows
successful complete parsing. Errors expose no partial result.

`text::unquote_element(input, context)` decodes just the first literal scalar or
escape, returning `Result[(QuotedElement, string), UnquoteError]`; the second
value is the unchanged, unexamined tail. `QuoteContext::{Unquoted, Single, Double}`
selects quote rules. Single/Double reject their own unescaped delimiter and allow
only that delimiter's quote escape. Unquoted allows either literal quote but
neither quote escape. Backticks are ordinary bytes here, not a raw-literal mode.

`QuotedElement::Byte(byte)` represents literal ASCII, short escapes and hexadecimal
or octal byte escapes. `Scalar(char)` represents literal non-ASCII characters and
Unicode escapes, including ASCII-valued escapes such as `\u0041`. Thus `\xff`
is a byte, not an instruction to encode U+00FF as UTF-8. Unlike a whole-literal
parser, this primitive accepts literal LF and CR: callers enforce enclosing
literal rules. Empty or malformed input yields `Syntax(0)` without a partial
result. It inspects at most one scalar/escape (ten input bytes) and uses bounded
scratch, so it has no variable-sized decoded-output limit. It introduces no
source grammar forms and does not process invalid-UTF-8 input strings.
For example, `unquote_element("\\xffrest", QuoteContext::Unquoted)` returns
`Ok((QuotedElement::Byte(255), "rest"))`.

The decoded-output limit is enforced before each append, including each complete
scalar encoding. It does not bound input scanning (for example discarded raw CRs)
or allocation overhead; callers processing untrusted input must also limit its
size. These are in-memory decoders, not incremental stream readers or extensions
to the GoML source grammar.

```goml
use std::text;
use std::bytes;

fn quoted_name(value: string) -> Result[string, bytes::TransformError] {
    text::quote_ascii(value, 4096)
}

fn decoded_name(literal: string) -> Result[string, text::UnquoteError] {
    text::unquote(literal, 4096)
}
```

### Numeric parsing and checked arithmetic

`num::format_float32_fixed(value, decimal_places, max_bytes)` and
`format_float64_fixed` return `Result[string, num::FormatFloatError]` with exactly
the requested nonnegative number of digits after the decimal point. Zero places
omits the point. Conversion starts from the IEEE bits, constructs the exact
decimal value with pure integer arithmetic and rounds once to nearest, ties to
even. This formats the represented binary value, not an assumed original decimal
input: for example 2.675 at two places follows the actual f64 value. Negative zero
and negative values rounded to zero retain their minus sign.

Nonfinite values produce `NaN`, `+Inf` or `-Inf`; a valid precision does not pad
these tokens. This spelling is specific to these APIs and does not change the
existing shortest `ToString` output. Limits include sign, decimal point and all
digits. A negative limit returns `InvalidLimit` before precision validation;
negative places return `InvalidPrecision`, even for NaN/infinity. Excessive
output returns `OutputLimit` without returning a partial string. The complete
output length is checked before allocating its buffer; bounded exact-decimal
scratch storage depends on IEEE width, not requested precision. Allocation failure
is not recovered. No native formatting backend, new syntax or floating-point
environment/rounding-mode dependency is introduced.

`num::format_float32_scientific(value, decimal_places, uppercase, max_bytes)` and
`format_float64_scientific` use the same exact conversion and rounding with one
digit before the point and exactly `decimal_places` after it. `uppercase` selects
`E` instead of `e`; the exponent always has a sign and at least two decimal digits.
Rounding may increase the exponent (9.5 with zero places becomes `1e+01`). Zero
uses exponent zero, retaining a negative sign when present.

`num::format_float32_general(value, significant_digits, uppercase, max_bytes)` and
`format_float64_general` round to the requested nonnegative number of significant
digits, treating zero as one. They remove insignificant trailing zeros, then use
scientific notation when the rounded decimal exponent is below -4 or at least
the effective precision; otherwise they use fixed notation. `uppercase` changes
only the scientific exponent marker. Large precision does not force padding or
allocation proportional to precision: once the exact digits fit, it preserves
them subject to the output bound. Negative precision is an error, not a request
for shortest formatting; existing ToString still provides its established shortest
representation. Both families share `FormatFloatError`, validation order,
nonfinite spelling and complete-output preflight with fixed formatting.

`num::format_float32_binary(value, max_bytes)` and `format_float64_binary` emit
an exact decimal integer significand followed by `p` and a signed, unpadded
binary exponent. The significand is the IEEE integer mantissa, not a string of
binary digits. Subnormals retain their unnormalized mantissa. Zero likewise uses
the subnormal scale (`0p-149` for f32, `0p-1074` for f64); negative zero retains
its sign. This representation has no precision parameter.

`num::format_float32_hex(value, fractional_digits, uppercase, max_bytes)` and
`format_float64_hex` emit normalized hexadecimal significands with a binary
exponent: `0x1...p±dd` for nonzero finite values and `0x0...p+00` for zero.
Uppercase selects `0X`, hexadecimal A–F and `P`. Exponents have at least two
decimal digits. Nonnegative precision gives exactly that many fractional hex
digits, with ties-to-even rounding and exponent adjustment on carry. Precision
`-1` instead emits the exact value without insignificant trailing zeros; lower
precisions are invalid. Subnormals are normalized, not rounded away. Large
precisions append zeros only after checking the complete output byte limit.

Both families use bit/integer operations and the same nonfinite spelling,
sign preservation, recoverable errors and allocation limitations as the decimal
formatters. Negative output limits are rejected before precision validation, and
invalid hex precision is rejected even for nonfinite values. Bounded mantissa and
exponent scratch strings are independent of the requested output size. No new
native formatting primitive or source grammar is introduced.

`num::format_float32_shortest(value, notation, max_bytes)` and
`format_float64_shortest` return the shortest significant decimal digits that
parse back to the original f32/f64 bit pattern, rendered using `FloatNotation`:
`Fixed`, `Scientific`, `ScientificUpper`, `General` or `GeneralUpper`. This does
not promise the fewest total characters across different notations. Fixed output
can require hundreds of leading/trailing zeros; scientific output retains signed,
at-least-two-digit exponents. General shortest output selects scientific notation
below exponent -4 or at exponent 6 and above, independently of the number of
significant digits; this differs from explicit-precision general formatting.

Conversion uses exact decimal digits and checks nearest/lower/upper candidates
in increasing significant-digit counts against the existing pure GoML scalar
parser. It checks at most nine counts for f32 and seventeen for f64, preserving
negative zero. Nearest candidates use ties-to-even rounding; adjacent candidates
cover asymmetric rounding intervals. An unexpected inability to find a matching
candidate returns `FormatFloatError::RoundtripFailure`, not a non-shortest
fallback or a panic. Nonfinite spelling, negative-limit rejection and output
preflight match the other formatters. The output limit excludes bounded candidate
scratch and conversion work; existing compiler-owned ToString fallbacks and source
syntax are unchanged.

```goml
use std::num;

fn fixed_measurement(value: f64) -> Result[string, num::FormatFloatError] {
    num::format_float64_fixed(value, 3, 64)
}

fn compact_measurement(value: f64) -> Result[string, num::FormatFloatError] {
    num::format_float64_general(value, 6, false, 64)
}

fn exact_hex_measurement(value: f64) -> Result[string, num::FormatFloatError] {
    num::format_float64_hex(value, -1, false, 64)
}

fn shortest_measurement(value: f64) -> Result[string, num::FormatFloatError] {
    num::format_float64_shortest(value, num::FloatNotation::General, 64)
}
```

`num::parse_bool_structured(value)` returns `Result[bool, num::ParseBoolError]`.
It accepts exactly `1`, `t`, `T`, `true`, `TRUE`, `True` for true and `0`, `f`,
`F`, `false`, `FALSE`, `False` for false. No trimming, general case folding,
numeric coercion, localized spelling or `yes`/`no` convention is applied.
Malformed input returns an error, never a default false value. `error.input()`
retains the original string for explicit inspection; its ToString/Debug message
does not echo that input. Error equality compares the original input. Bound input
size before parsing untrusted data if retaining error strings would be costly.
The existing boolean `ToString` implementation remains the canonical formatter,
producing lowercase `true`/`false`; no separate formatting API is needed.
This uses ordinary functions, structs and Result and adds no grammar forms.

```goml
use std::num;

fn parse_switch(value: string) -> Result[bool, num::ParseBoolError] {
    num::parse_bool_structured(value)
}
```

`num::format_int(value: i64, radix)` and `format_uint(value: u64, radix)` return
`Result[string, FormatIntError]` using pure GoML integer arithmetic. Formatting
accepts radices 2 through 36, emits lowercase digits without prefixes, separators
or a positive sign, and formats zero as `"0"`. Signed negatives use `-` followed
by the magnitude, including the minimum i64; this is not two's-complement display.
Radix zero is invalid for formatting even though it is supported by parsing.

`write_int(output: MutSlice[byte], value: i64, radix)` and `write_uint` return the
number of ASCII bytes written. They validate before modifying the destination:
invalid radix or insufficient capacity leaves all bytes unchanged, and success
leaves the unused suffix unchanged. `FormatIntError` distinguishes
`InvalidRadix(radix)` and `BufferTooSmall(needed, available)`; radix validation takes
precedence. These helpers allocate bounded temporary digit vectors (at most 65
output bytes) and do not promise allocation-free conversion. Existing parsing,
scalar decimal to_string and floating-point behavior are unchanged. No grammar or
compiler/runtime intrinsic is added.

Numeric parsing returns `Result[_, num::ParseIntError]` or `Result[_, num::ParseFloatError]`. Integer radix parsing accepts radix `0` or `2..36`; radix `0` recognizes `0b`, `0o`, and `0x` prefixes and permits Go-style digit separators. Invalid radices, malformed input, and overflow return `Result::Err`. Floating-point parsing supports decimal and hexadecimal IEEE 754 input, signed exponents, digit separators, `inf`, `infinity`, and `NaN`, and rounds directly to the requested `f32` or `f64` width.

`ParseFloatError::is_range()` distinguishes numeric overflow from malformed
syntax for `parse_float32_structured` and `parse_float64_structured`. Overflow
returns `Err` with message `value out of range`; it does not return a successful
infinity or a saturated numeric value. Explicit infinity spellings are successful
values. Underflow rounds to a subnormal or signed zero and succeeds. Malformed
syntax (including a trailing invalid character after a huge exponent) has
`is_range() == false` and message `invalid syntax`. Both error classes retain the
existing `kind() == io::ErrorKind::InvalidInput`, input/context, operation and
raw-OS-code interfaces. The flag is part of error equality. The compatibility
constructor `ParseFloatError::new(input, message)` leaves the flag false and does
not infer a classification from caller-supplied diagnostic text. No syntax or
runtime primitive is added; parsing continues to use the existing pure GoML
scalar conversion implementation. These APIs have no input-length limit, so
callers must bound untrusted numeric strings before parsing or retaining errors.
The rational conversion stage checks conservative binary exponent bounds before
allocating shifted integer temporaries. Definite overflow and underflow therefore
do not allocate storage proportional to a huge binary exponent; boundary cases
still use exact rounding, including the half-smallest-subnormal tie. This does
not bound the work or storage needed to scan and accumulate a long mantissa.

Long mantissas are interpreted by their numeric value, not truncated to a fixed
decimal digit buffer. For example, `1` followed by 2048 zeros and `e-2048` parses
as exactly one at both widths. This intentionally differs from Go 1.26's
fixed-buffer fallback for some long inputs; exact rational test references are
used for these cases instead of treating every Go result as authoritative.
Exponent accumulation saturates only beyond an input-length-derived bound that
exceeds the mantissa's possible compensating scale. Long fractional zero runs
can therefore be canceled by a large exponent without an arbitrary fixed
exponent cutoff changing the value. Remaining exponent digits and separators
are still syntax-checked after saturation; an invalid suffix is not an overflow.

```goml
use std::num;

fn parse_float_input(value: string) -> Result[f64, string] {
    num::parse_float64_structured(value).map_err(|error: num::ParseFloatError| {
        if error.is_range() {
            "outside f64 range"
        } else {
            "invalid numeric syntax"
        }
    })
}
```

`num::parse_int_structured`, radix and unsigned variants, and the structured float parsers return domain parse errors. The `checked_*_int64` operations return `None` on overflow; the corresponding `saturating_*_int64` operations clamp to the signed 64-bit bounds.

`num::parse_int_bits(input, radix, bits)` and `parse_uint_bits` return i64 and u64
respectively, constrained to the requested signed or unsigned range. Widths 1..64
are accepted; zero selects GoML's 64-bit native integer width. Thus signed width 1
accepts only -1 and 0, while unsigned width 1 accepts 0 and 1. The existing radix,
prefix, sign and underscore rules apply; unsigned parsing does not accept a sign.
These functions never truncate or return a saturated value on failure.

`IntWidthError::InvalidWidth(bits)` precedes input/radix validation.
`IntWidthError::Parse(error)` retains the existing ParseIntError for syntax,
invalid radix or full-width overflow. A value successfully parsed at 64 bits but
outside a narrower requested range yields `OutOfRange(effective_bits)`. This
deliberate validation order can differ from Go's error-category precedence;
callers receive Result rather than Go's value-plus-error pair. No existing parser
signature, cast rule, literal inference or compiler/runtime boundary changes.

The signed one-bit range is enforced mathematically: values below -1 are errors.
Go 1.26's ParseInt can report success with -1 for some overflowing one-bit negative
inputs; GoML deliberately does not reproduce that saturation quirk.

`num::parse_complex_f32(input)` and `parse_complex_f64(input)` return
`Result[(f32, f32), ParseComplexError]` and `Result[(f64, f64), ParseComplexError]`.
The tuple is `(real, imaginary)`; suffixes name each component's width, not Go's
total complex width. These are numeric conversion functions, not a second complex
arithmetic type or source-literal syntax. Each component reuses the corresponding
pure floating parser, including direct f32 rounding and signed zero preservation.

Accepted forms are `N`, `Ni`, `N+Ni`, `N-Ni`, optionally surrounded by one pair of
parentheses, with no whitespace trimming. Numeric components use the existing
decimal/hexadecimal, separator, infinity and NaN rules. Imaginary-only forms still
need a numeric coefficient: `i`, `+i` and `-i` are errors. A combining plus may
precede a negative component (`1+-2i`) or NaN (`1+NaNi`); doubled plus and negative
NaN are invalid. An omitted component is positive zero. `ParseComplexError` is
`Syntax` or `OutOfRange`, with `is_range()` for classification; malformed syntax
in either component takes precedence over the other component's overflow.
Errors return no saturated/partial components and their diagnostics do not retain
or echo the input. There is no input-length limit; bound untrusted strings before
parsing. For example, `parse_complex_f64("(1-2i)")` returns `Ok((1.0, -2.0))`.

`num::format_complex_f32(real, imaginary, format, max_bytes)` and
`format_complex_f64` return `Result[string, FormatFloatError]`, always spelling
both components as `(real+imaginaryi)` or `(real-imaginaryi)`. Inputs have the
component width named by the function; no complex runtime type is involved.
`ComplexFormat` selects `Shortest(FloatNotation)`, `Fixed(decimal_places)`,
`Scientific(decimal_places, uppercase)`, `General(significant_digits, uppercase)`,
`Binary`, or `Hex(fractional_digits, uppercase)`. Component precision, rounding,
nonfinite tokens and validation rules are exactly those of the corresponding
scalar formatter. Shortest is explicit; only Hex uses -1 for exact trimmed output.

The limit includes parentheses, separator sign and trailing `i`. Negative zero
and negative imaginary values retain their signs, positive infinity already
supplies its plus, and imaginary NaN is preceded by plus. The first component
error is returned with no partial output; negative output limits precede precision
validation. Complete framing length is checked without overflowing before string
assembly. Each temporary component string is separately bounded by max_bytes;
this is a final-output limit, not a bound on total allocations or conversion work.
For example, `format_complex_f64(1.0, -2.0, ComplexFormat::Fixed(1), 16)` returns
`Ok("(1.0-2.0i)")`. Binary significand notation is a display format and is not
accepted by the complex decimal/hex parser. These APIs introduce no grammar forms.

`std::num::ToFloat` supplies `.to_f32()` and `.to_f64()` for every scalar integer width. Conversions round directly to the destination IEEE 754 width using nearest, ties to even; `u64` to `f32` does not round through `f64`. Large integers can lose precision.

`std::num::TryToInt` supplies `.try_to_isize()`, `.try_to_i8/i16/i32/i64()`, and unsigned counterparts for both `f32` and `f64`. These truncate toward zero before checking the destination range, returning `FloatConversionError::NonFinite` for NaN/infinity or `OutOfRange` for overflow. Consequently `(-0.9).try_to_u64()` succeeds with zero. `float_to_i64(value, rounding)` and `float_to_u64(value, rounding)` accept `Rounding::{TowardZero, Floor, Ceil, NearestAway, NearestEven}` for explicit rounding followed by the same range checks. These methods require the corresponding trait imports and add no cast syntax.

```goml
use std::num::{ToFloat, TryToInt};

let count: u64 = 9007199254740993;
let approximate: f64 = count.to_f64();
let truncated = (12.75).try_to_i32();
```

### Environment, paths, and processes

`env::current_dir_structured`, `current_exe_structured`, and `var_structured`, `path::absolute_structured`, and the `process` structured execution methods expose whole-operation errors. Process timeout methods take `time::Duration`, terminate and wait through the command runtime, and return `TimedOut` through `process::Error`.

Paths remain UTF-8 `string` values. `path::separator` reports the host separator, `components` recognizes both slash forms, `relative` uses host path rules, and `windows_prefix` recognizes drive and UNC prefixes independently of the host operating system. Non-UTF-8 operating-system names cannot be represented and therefore cannot appear in these APIs.

### Linux amd64 system calls

`std::os::linux::syscall` is an explicit low-level escape hatch supported only on Linux amd64. `syscall6(number: usize, args: [usize; 6]) -> Result[SyscallResult, Error]` accepts any syscall number and six machine words. Unused argument positions should contain zero. Signed arguments use their machine-word bit pattern, for example `(-100).to_usize()` for Linux `AT_FDCWD`.

`SyscallResult` has public `r1`, `r2`, and `errno` fields, all `usize`, preserving the values reported by Go's `syscall.Syscall6`. `is_ok()` tests `errno == 0`. A kernel failure still returns `Ok(SyscallResult)` with a nonzero errno; callers must inspect it. The outer `Err(Error::UnsupportedTarget)` reports that the executing target is not Linux amd64, before any syscall is issued. Other Go targets are not guaranteed to compile. Syscall numbers, layouts, and constants are specific to the Linux amd64 ABI; this package does not select numbers for another architecture.

`SyscallResult::into_result()` converts a successful first return word to `Ok(usize)` and a nonzero errno to `Err(Errno)`. The public tuple field in `Errno(pub usize)` preserves the numeric code, including unknown future codes. `Errno::name()` and `errno_name(usize)` return `Option[string]` with the canonical symbolic name; aliases resolve to the original name, such as `EWOULDBLOCK` to `EAGAIN`. `to_string()` returns that name, or `"errno N"` for an unknown code. Zero has no errno name.

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

`syscall6_with_pointers(number: usize, args: [Arg; 6], pointers: Slice[Pointer])` extends the buffer API to synchronous native structures containing pointers, including `iovec` arrays and `msghdr` records. `Pointer` has public `buffer: MutSlice[byte]`, `offset: isize`, and `target: MutSlice[byte]` fields. Each descriptor temporarily writes the target's address into the eight-byte little-endian pointer field at the offset within the buffer view; an empty target writes zero. Root buffers use `Arg::Buffer`, and descriptors may connect intermediate buffers to other buffers at any depth, including aliases and cycles. Every participating allocation is pinned before pointer fields are installed and remains pinned until the call and pointer restoration finish. There is no limit of six nested buffers.

Every pointer slot is bounds-checked before any mutation or syscall. An invalid slot returns `Error::InvalidPointer { index, offset, length }`, identifying the descriptor and its buffer view. Descriptors are installed in list order. All original pointer-field bytes are restored on return, including kernel errors; overlapping slots are restored to their original bytes as well. Kernel writes outside these slots remain visible. Treat pointer slots as temporary input fields: kernel output written into them is discarded during restoration. Buffers and pointer fields must not be read, mutated, or used by another syscall concurrently with the call.

For example, Linux amd64 `iovec` records contain an eight-byte pointer followed by an eight-byte length. This writes two views with one kernel operation, using the native layout documented by [readv/writev](https://man7.org/linux/man-pages/man2/readv.2.html):

```goml
use std::bytes::endian;
use std::os::linux::syscall;
use std::os::linux::syscall::{Arg, Pointer};

fn write_pair(fd: usize, first: MutSlice[byte], second: MutSlice[byte]) -> Result[syscall::SyscallResult, syscall::Error] {
    let vectors: Vec[byte] = Vec::new();
    for _ in 0..32 { vectors.push(0); }
    let _ = endian::write_u64(vectors.as_mut_slice(), 8, first.len().to_u64(), endian::Endian::Little);
    let _ = endian::write_u64(vectors.as_mut_slice(), 24, second.len().to_u64(), endian::Endian::Little);
    let pointers = Vec::from_array([
        Pointer { buffer: vectors.as_mut_slice(), offset: 0, target: first },
        Pointer { buffer: vectors.as_mut_slice(), offset: 16, target: second },
    ]);
    syscall::syscall6_with_pointers(syscall::SYS_WRITEV, [
        Arg::Word(fd), Arg::Buffer(vectors.as_mut_slice()), Arg::Word(2),
        Arg::Word(0), Arg::Word(0), Arg::Word(0),
    ], pointers.as_slice())
}
```

The caller is responsible for the syscall ABI, valid addresses, buffer lengths, alignment, native structure encoding, resource cleanup, and synchronization. Byte buffers can carry explicitly encoded native records; ordinary GoML structs have no kernel-layout guarantee. Pointer-slot bounds checks do not validate syscall-specific structure layouts or lengths stored in those records. `Word` does not retain or pin any Go allocation, so a Go-managed address must not be smuggled through an integer. Numeric addresses returned by operations such as `mmap` remain raw words and require appropriate explicit cleanup such as `munmap`.

Buffer lifetimes cover synchronous calls only, including nested pointer graphs supplied through `Pointer`. Operations that retain a pointer after returning and arbitrary Go object memory are outside this buffer API. The runtime uses scheduling-aware `Syscall6`, not `RawSyscall6`. It performs one call without automatic `EINTR` retry or completion of partial reads/writes. Per-thread operations require thread-affinity handling that this API does not provide. Raw changes to threads, process creation, signal handlers, or the Go runtime's address space can violate runtime invariants; accepting a number does not make every kernel operation safe to use from GoML. Calls remain subject to kernel availability, process permissions, and sandbox policy.

The package exports all 385 native amd64 `SYS_*` numbers in the pinned [Linux 7.2 syscall table](https://github.com/torvalds/linux/blob/v7.2/arch/x86/entry/syscalls/syscall_64.tbl), excluding the separate x32 ABI. This includes legacy entries and newer interfaces such as `SYS_OPENAT2`, `SYS_CLONE3`, `SYS_PIDFD_OPEN`, `SYS_IO_URING_SETUP`, and `SYS_LANDLOCK_CREATE_RULESET`. A constant does not guarantee that the running kernel implements or permits the call. All 136 errno constants and aliases from the same release's `asm-generic/errno-base.h` and `errno.h` are exported. Other numbers can still be passed directly. The numeric manifest and upstream SHA-256 checksums live in `tools/syscall/linux-amd64.json`; `python3 tools/syscall/generate.py` regenerates constants, `--check` verifies them offline, and `--check --verify-upstream` also checks the original kernel files. This package uses ordinary imports, functions, enums, arrays, and mutable slices; it introduces no new grammar or `unsafe` syntax and cannot be used in `comptime`.

### Linux amd64 ABI records

`std::os::linux::abi` provides public-field records with `Type::size() -> isize`, `encode() -> Vec[byte]`, and `Type::decode(Slice[byte]) -> Result[Type, endian::BoundsError]`. Encoding uses the Linux amd64 little-endian layout, including zeroed padding. Decoding requires the whole record and accepts trailing bytes. These codecs validate byte bounds; the kernel validates semantic values such as time ranges and flags. Ordinary GoML structs still have no native layout guarantee.

| Records | Encoded sizes in bytes |
| --- | --- |
| `Timespec { seconds, nanoseconds }`, `Timeval { seconds, microseconds }` | 16 each |
| `Itimerspec { interval, value }` | 32 |
| `Rlimit { current, maximum }` | 16 |
| `PollFd { fd, events, revents }`, `EpollEvent { events, data }` | 8, 12 (packed) |
| `Iovec { base, length }`, `Msghdr`, `Cmsghdr { length, level, kind }` | 16, 56, 16 |
| `Flock { kind, whence, start, length, pid }`, `OpenHow { flags, mode, resolve }` | 32, 24 |
| `Stat`, `Statx`, `StatxTimestamp { seconds, nanoseconds }` | 144, 256, 16 |
| `Rusage` | 144 |

`Stat` exposes device/inode/link identifiers, mode and ownership, size/block information, and `atime`, `mtime`, `ctime` as `Timespec`. `Statx` also exposes `mask`, attributes, birth time, device major/minor identifiers, mount ID, and direct-I/O alignment. Check its returned mask before using optional fields. Fields after the direct-I/O alignment prefix are currently reserved by this codec. `Rusage` carries user/system `Timeval` values and the Linux resource counters. Pointer-valued record fields are machine words; use the nested-buffer builders to reference Go-managed storage.

`Iovecs::new(buffers: Slice[MutSlice[byte]]) -> Result[Iovecs, string]` accepts up to 1024 views and creates the encoded vector array and pointer descriptors. `buffer()`, `pointers()`, and `len()` supply the arguments for `readv`, `writev`, and related calls. `Message::new(name: MutSlice[byte], vectors: Iovecs, control: MutSlice[byte]) -> Result[Message, string]` constructs an `msghdr` graph for synchronous `sendmsg`/`recvmsg`; empty name/control views represent null pointers. Its `buffer()` and `pointers()` are used together with `syscall6_with_pointers`, and `header()` decodes lengths and flags after the call. Pointer slots are restored by that call. Builders retain the views and do not copy their contents. A builder, its views, and its graph must not be accessed concurrently while a call is using it; editing embedded lengths requires following the raw syscall ABI.

The layout manifest is `tools/syscall/linux-amd64-layouts.json`. Run `python3 tools/syscall/layouts.py`, then `goml fmt` in `lib/std`, to regenerate the codecs. `python3 tools/syscall/layouts.py --check --verify-headers` checks canonical output and compiles C assertions for every size, offset, and field width against Linux amd64 system headers. This API uses ordinary structs and methods and adds no grammar.

### Linux amd64 file descriptors

`std::os::linux::fd` provides explicit descriptor ownership and returns `Result[..., io::Error]`; it re-exports `Error` and `ErrorKind`. Errors retain `raw_os_code()` when the kernel supplies errno. Calls that can be interrupted retry `EINTR`, except `close` and the request-dependent raw `ioctl`. Nonblocking operations return `WouldBlock` for `EAGAIN`. These interfaces do not add readiness waits or cancellation.

`open(path: string, flags: usize, mode: u32)`, `open_at(directory: isize, path: string, flags: usize, mode: u32)`, and `open_at2(directory: isize, path: string, how: abi::OpenHow)` return `Fd`. `open_at_bytes` and `open_at2_bytes` accept raw `Slice[byte]` paths. Paths reject embedded NUL and otherwise preserve their bytes; string paths are encoded as UTF-8. `Fd` also has `open_at`, `open_at_bytes`, and `open_at2` methods that keep the directory descriptor alive during the operation. `memfd(name: string, flags: usize)` creates an anonymous memory-backed file. Open, memfd, and duplication operations set close-on-exec atomically. Flags include `O_*`, `AT_*`, `RESOLVE_*`, `MFD_*`, and `F_SEAL_*`; a flag or syscall constant does not guarantee kernel/filesystem support.

Assigning or passing an `Fd` shares its close state. `close()` is explicit and idempotent; there is no finalizer. It rejects new operations, waits for active operations/borrows, and closes the raw descriptor once, without retrying a failed close. `is_closed()` becomes true when closing begins. Blocking I/O must finish before close can return; closing does not cancel it. Operations on the same open handle may run concurrently. Callers must synchronize shared buffers and any compound sequence involving the shared file offset.

`duplicate()` creates an independent descriptor owner while preserving the kernel's shared file description and offset. `Fd::from_raw(raw: isize)` validates and adopts an existing descriptor without changing its flags; the caller transfers sole responsibility for closing it on success. `Fd::from_owned_raw(raw: isize)` only checks the numeric range and adopts an already-owned descriptor without an `fcntl` probe; use it for successful syscall results, and ensure the number is valid and exclusively owned. `into_raw()` waits for active operations and transfers that responsibility back, invalidating all aliases. `with_raw_fd[T](action: (usize) -> Result[T, Error])` keeps the descriptor open for the callback. `fd::with_raw_fds[T](descriptors: Slice[Fd], action: (Slice[usize]) -> Result[T, Error])` borrows a whole list iteratively, including repeated aliases, and releases every successful borrow if a later acquisition or the callback fails. The callback must not close, replace, retain, or re-adopt the raw number, or call `close`/`into_raw` on the same shared handle. Free functions accepting a raw directory number require the caller to keep it valid; use `with_raw_fd` for an owned directory or `AT_FDCWD` for the current directory.

| Operations on `Fd` | Behavior |
| --- | --- |
| `read(MutSlice[byte])`, `write(Slice[byte])` | Return the transferred byte count; reads return zero at EOF |
| `read_at(output, offset: i64)`, `write_at(input, offset: i64)` | Positional I/O without moving the file offset |
| `read_vectored(Slice[MutSlice[byte]])`, `write_vectored(Slice[Slice[byte]])` | Scatter/gather I/O, at most 1024 buffers |
| `read_vectored_at(outputs, offset: i64)`, `write_vectored_at(inputs, offset: i64)` | Full 64-bit positional scatter/gather I/O |
| `read_exact(output)`, `write_all(input)` | Loop over short transfers; report `UnexpectedEof` or `WriteZero` when progress stops |
| `seek(offset: i64, whence: usize)` | Return the new `u64` position; `SEEK_SET/CUR/END/DATA/HOLE` |
| `stat()`, `truncate(length: i64)`, `chmod(mode: u32)`, `chown(uid: u32, gid: u32)` | Metadata and file changes; `0xffffffff` leaves an ownership ID unchanged |
| `sync()`, `sync_data()`, `allocate(mode: usize, offset: i64, length: i64)` | `fsync`, `fdatasync`, and `fallocate` with `FALLOC_FL_*` flags |
| `descriptor_flags()`, `set_descriptor_flags(flags)`, `status_flags()`, `set_status_flags(flags)` | Descriptor and open-file status flags |
| `flock(operation: usize)`, `record_lock(command: usize, lock: abi::Flock)` | Whole-file and POSIX/OFD byte-range locks; the latter returns the resulting `Flock` |
| `fcntl(command: usize, argument: usize)`, `ioctl(request: usize, buffer: MutSlice[byte])` | Raw word-argument fcntl and synchronous buffer ioctl; caller supplies the request ABI |
| `read_dir()` | Read one batch of `Vec[DirEntry]`, empty at EOF, advancing the directory offset |

Writes copy immutable input views into temporary buffers. Neither input copying nor kernel I/O permits concurrent mutation of those views. `read_exact` and `write_all` can have made partial progress when returning an error. Positional writes retain Linux `O_APPEND` behavior, which can append regardless of the supplied offset. Raw `fcntl` duplication commands return a raw descriptor owned by the caller; `duplicate()` is the managed alternative. `ioctl` requires a buffer large enough for the request and cannot register memory for use after return.

Directory-relative free functions are `stat_at`, `statx_at`, `mkdir_at`, `unlink_at`, `rename_at`, `link_at`, `symlink_at`, `readlink_at`, `chmod_at`, `chown_at`, and `utimens_at`, each with a `_bytes` variant. Rename uses `renameat2` and supports `RENAME_NOREPLACE`, `RENAME_EXCHANGE`, and `RENAME_WHITEOUT`; unlink uses `AT_REMOVEDIR` for directories. `utimens_at` takes access and modification `Timespec` values, including `UTIME_NOW`/`UTIME_OMIT` in the nanoseconds field. `readlink_at` returns raw target bytes. `DirEntry` preserves raw `name: Vec[byte]`, `inode: u64`, `next_offset: i64`, and Linux `kind: byte`; `.` and `..` are included and unknown kinds are retained. Directory order and offset cookies are kernel-defined.

```gom
use std::os::linux::fd;

fn append_bytes(path: string, bytes: Slice[byte]) -> Result[(), fd::Error] {
    let file = fd::open(path, fd::O_WRONLY | fd::O_CREAT | fd::O_APPEND, 0x180)?;
    let written = file.write_all(bytes);
    let closed = file.close();
    written?;
    closed
}
```

These interfaces support Linux amd64 only. They use ordinary imports, structs, functions, and methods and add no native pointer or `unsafe` syntax.

### Linux amd64 processes

`std::os::linux::process` provides Linux process operations and re-exports `io::Error`/`ErrorKind`. PID arguments must fit signed 32 bits; signal numbers range from zero through 64. Zero checks existence/permission, and negative or zero PIDs retain the kernel's process-group selection semantics for `kill` and `wait`. Exported `SIG*` constants use native Linux numbers, including 32/33, rather than glibc's adjusted real-time range.

| Functions | Result or behavior |
| --- | --- |
| `pid()`, `parent_pid()`, `thread_id()` | `Result[isize, Error]` |
| `uid()`, `effective_uid()`, `gid()`, `effective_gid()` | `Result[u32, Error]` |
| `user_ids()`, `group_ids()`, `groups()` | Real/effective/saved `ResIds`, or supplementary `Vec[u32]` |
| `process_group(pid)`, `session(pid)` | Group/session ID; zero selects the caller |
| `set_process_group(pid, group)`, `create_session()` | Apply Linux group/session rules |
| `kill(pid, signal)`, `signal_thread(group, thread, signal)` | Process/group signaling and `tgkill` |
| `pidfd_open(pid, flags)`, `pidfd_signal(handle, signal)` | Owned close-on-exec `fd::Fd` and signaling through a stable process handle |
| `wait(pid, options)` | `Result[Option[WaitOutcome], Error]`; `None` when `WNOHANG` finds no ready child |
| `resource_usage(who)` | `abi::Rusage`; `RUSAGE_SELF`, `RUSAGE_CHILDREN`, `RUSAGE_THREAD` |
| `resource_limit(pid, resource)`, `set_resource_limit(pid, resource, limit)` | Read or atomically replace `abi::Rlimit`, returning the previous limit on replacement |
| `priority(which, who)`, `set_priority(which, who, value)` | Normal nice values from -20 to 19 with kernel permission checks |

`WaitOutcome` exposes `pid`, `status: WaitStatus`, and `usage: abi::Rusage`. `WaitStatus { raw: u32 }` has `exited`, `signaled`, `stopped`, `continued`, `core_dumped`, and `success` predicates plus optional `exit_code`, `signal`, and `stop_signal`. Wait options include `WNOHANG`, `WUNTRACED`, and `WCONTINUED`. Resource limits expose all 16 `RLIMIT_*` selectors and `RLIM_INFINITY`; Linux maximum RSS is in KiB. `thread_id`, `RUSAGE_THREAD`, and `PRIO_PROCESS` with `who == 0` refer to the executing OS thread, which can change as GoML tasks are scheduled. They do not identify or configure a stable GoML task. Resource-limit setters update the kernel directly. Go child creation may restore its original `RLIMIT_NOFILE` soft limit when the current soft value equals the initial hard limit minus one, including after an explicit raw update to that value; other current values are inherited normally.

`Command::new(program)` builds an asynchronous child command. Its chaining methods are `arg(string)`, `args(Slice[string])`, `current_dir(string)`, `env(key, value)`, `env_clear()`, `stdin(fd)`, `stdout(fd)`, `stderr(fd)`, `extra_fd(fd)`, `new_session(bool)`, and `new_process_group(bool)`. `spawn()` returns a `Child`. The builtin runtime uses Go's [runtime-coordinated ForkExec](https://go.dev/src/syscall/exec_linux.go); GoML code does not issue raw fork/clone. Program names without a slash use the parent's PATH through Go's `exec.LookPath`; paths containing a slash are executed after applying the child working directory. Command environment changes apply to the child. Environment is inherited at spawn unless cleared, and later overrides win. Keys must be nonempty and contain neither NUL nor `=`; values reject NUL. An empty working directory means inherit.

Standard descriptors 0/1/2 are inherited by default. Explicit descriptors remain borrowed throughout spawn. Extra descriptors occupy consecutive child slots beginning at 3. The child receives separate descriptor-table entries referring to the same file descriptions; the parent's owners remain open. A new session also creates a new process group; enabling both options has the same effect as `new_session(true)`. Command values share their argument, environment, and descriptor lists through ordinary `Vec` semantics; do not mutate a command concurrently with spawning it.

Spawn requests an atomic pidfd using Linux `CLONE_PIDFD`, requiring the corresponding kernel capability. `Child.pid()` returns its numeric PID, while `signal(signal)` and `kill()` use the pidfd to avoid PID reuse. `wait()` reaps once and caches the outcome across aliases; simultaneous waiters receive the cached result. `try_wait()` returns `None` while the child is running or another waiter owns the wait operation. `wait_event(options)` additionally observes stop/continue events with the supported wait flags. Signaling remains possible while another task waits. Terminal reaping closes the internal pidfd; subsequent signals report a closed descriptor. Call `wait()` after a successful spawn, including after `kill()`; dropping a child does not kill or reap it. Its `Child` owns terminal reaping, so other wait APIs and automatic SIGCHLD reaping must not take that responsibility.

```gom
use std::os::linux::process;

fn child_status() -> Result[isize, process::Error] {
    let child = process::Command::new("/bin/sh").arg("-c").arg("exit 7").spawn()?;
    let result = child.wait()?;
    Result::Ok(result.status.exit_code().unwrap_or(-1))
}
```

Credential mutation, signal-handler replacement, and signal-mask replacement are not provided as convenience wrappers because they require coordination with the multithreaded runtime. These process APIs support Linux amd64 only, use ordinary language constructs, and add no syntax.

### Linux amd64 memory mappings

`std::os::linux::memory::anonymous(length: isize, protection: usize, flags: usize)` and `file(descriptor: fd::Fd, length: isize, offset: i64, protection: usize, flags: usize)` return `Result[Mapping, io::Error]`. The package re-exports `Error` and `ErrorKind`. Anonymous mappings add `MAP_ANONYMOUS`; callers choose `MAP_PRIVATE` or `MAP_SHARED`. File mappings borrow the descriptor during `mmap` and remain valid after it closes. `MAP_PRIVATE` provides copy-on-write behavior, while `MAP_SHARED` exposes shared backing-file changes.

Length must be positive, and file offsets must be nonnegative multiples of `PAGE_SIZE` (4096 on Linux amd64). Copied access is limited to the requested length, including when the last mapped page is partial. `Mapping` rejects fixed-address, growing-stack, explicit hugetlb, and uninitialized mapping flags. These specialized operations remain accessible through raw syscalls. Managed mappings expose no raw address, Go slice view, executable function pointer, remap, or partial unmap.

| `Mapping` methods | Behavior |
| --- | --- |
| `len()`, `is_closed()` | Requested length and shared lifecycle state |
| `read_at(output: MutSlice[byte], offset: isize)`, `write_at(input: Slice[byte], offset: isize)` | Checked copied access, returning the actual byte count |
| `read_exact_at(output, offset)`, `write_all_at(input, offset)` | Repeat short transfers until complete or error |
| `protect(protection)` | Change the whole mapping's protection using `PROT_NONE/READ/WRITE/EXEC/SEM` |
| `sync(flags)`, `advise(advice)` | Whole-mapping `msync` and `madvise` |
| `lock(flags)`, `unlock()` | Whole-mapping `mlock`/`mlock2` and `munlock`; `MLOCK_ONFAULT` is supported |
| `residency()` | One boolean per base page from `mincore`, describing a momentary snapshot |
| `close()` | Explicit, idempotent unmap across aliases |

Copying uses [process_vm_readv/process_vm_writev](https://man7.org/linux/man-pages/man2/process_vm_readv.2.html) against the current process, with local Go buffers pinned only during the call. It has syscall/allocation overhead and can return `ENOSYS`/`EPERM` under unsupported kernels or syscall restrictions. Inaccessible or truncated file pages return short copies or `EFAULT` without directly dereferencing mapped storage from Go. Errors retain the kernel errno. Exact operations can modify a prefix before an error; writes copy the immutable input first. Do not concurrently mutate caller buffers during a call.

Aliases share a channel that serializes copying, protection changes, other operations, and unmap. `close()` waits for active access, unmaps once, and preserves the mapping if unmap fails so it can be retried. Other operations after close return `EBADF`; `len()` retains the requested length. There is no finalizer. Reads require explicit `PROT_READ` and writes explicit `PROT_WRITE`, even where amd64 hardware implies additional permissions. Missing declared access returns `EACCES`. Independent mappings and other processes require their own synchronization; mapped bytes are not an atomic IPC protocol.

```gom
use std::os::linux::memory;

fn mapped_bytes() -> Result[Vec[byte], memory::Error] {
    let region = memory::anonymous(4096, memory::PROT_READ | memory::PROT_WRITE, memory::MAP_PRIVATE)?;
    defer { let _ = region.close(); };
    region.write_all_at(b"hello".as_slice(), 0)?;
    let output = b"?????";
    region.read_exact_at(output.as_mut_slice(), 0)?;
    Result::Ok(output)
}
```

The package exposes Linux protection, mapping, synchronization, advice, and locking constants; a constant does not guarantee kernel support. Operations obey kernel permissions and resource limits. This package is Linux amd64 only and introduces no pointer or layout syntax.

### Linux amd64 IPC and readiness

`std::os::linux::ipc` uses owned `fd::Fd` values and structured `io::Error` results. `pipe(flags)` returns `Pipe { reader, writer }`; `socket_pair(kind, flags)` returns `SocketPair { first, second }` using the Unix domain. Socket kinds include `SOCK_STREAM`, `SOCK_DGRAM`, and `SOCK_SEQPACKET`. Flags include nonblocking operation; newly created descriptors always have close-on-exec. `shutdown(socket, direction)` supports `SHUT_RD`, `SHUT_WR`, and `SHUT_RDWR`. Close each returned descriptor explicitly.

`send_message(socket, data: Slice[byte], descriptors: Slice[fd::Fd], flags)` sends one message with optional `SCM_RIGHTS` descriptors, returns the transferred byte count, and adds `MSG_NOSIGNAL`. It copies payload input and borrows all descriptors for the call. At most 253 descriptors can be sent, and descriptor transfer requires at least one payload byte. A partial stream send may already transfer the rights; retrying the entire descriptor-bearing send can duplicate them at the receiver.

`receive_message(socket, output: MutSlice[byte], max_descriptors: isize, flags)` returns `ReceivedMessage { byte_count: usize, copied: isize, flags: usize, descriptors: Vec[fd::Fd] }`. It atomically sets close-on-exec on received descriptors with `MSG_CMSG_CLOEXEC`. The limit is 0–253. `MSG_TRUNC` can make `byte_count` larger than `copied`; zero-byte datagrams are valid. Control truncation or a descriptor-limit error closes the received descriptors and returns an error after the receive; the payload is consumed unless `MSG_PEEK` was requested. On success the caller owns every returned descriptor. Immutable send inputs and mutable receive buffers must not be mutated concurrently with these calls.

`eventfd(initial: u32, flags)` returns a counter descriptor; `EFD_NONBLOCK` and `EFD_SEMAPHORE` control its behavior. `read_counter(fd)` and `write_counter(fd, value: u64)` perform the exact eight-byte counter operation. `timerfd(clock, flags)` creates a timer descriptor; `set_timer(timer, flags, value: abi::Itimerspec)` returns the previous setting and `get_timer(timer)` reads it. A zero value disarms the timer, an interval makes it periodic, and `TFD_TIMER_ABSTIME` selects absolute expiration. Read expirations through `read_counter`; nonblocking empty counters report `WouldBlock`.

`poll(interests: Slice[PollInterest], timeout_ms: i32)` borrows all descriptors and returns `Vec[PollEvent]`. `PollInterest { descriptor, events: i16 }` specifies a `POLL*` mask; `PollEvent { index, events }` identifies the ready input entry. A timeout of -1 waits indefinitely, zero polls immediately, and positive values are milliseconds. Poll returns `EINTR` instead of restarting a finite timeout.

`Epoll::new()` creates a shared reactor handle that must be closed explicitly. `add(descriptor, events: u32, token: u64)` duplicates and retains the watched file description, so the caller can close its original owner without invalidating the watch. Tokens must be unique among current watches and can use any `u64` value. `modify(token, events)` updates/rearms a watch; `remove(token)` unregisters it and closes its retained descriptor. `wait(max_events: isize, timeout_ms: i32)` returns `Vec[abi::EpollEvent]`, with each `data` field holding the supplied token. It supports `EPOLL*` masks, including edge-triggered and one-shot modes; callers must follow the kernel's draining/rearming rules. Interrupted waits return `EINTR`.

Each epoll registration uses an internal identity that is never reused; events for removed registrations are discarded, which can produce an empty result during concurrent removal. Adding/modifying/removing watches can proceed during waits. `close()` wakes every blocked waiter through a private eventfd before releasing descriptors; those waiters return `EBADF`. Closing is idempotent across aliases and closes all retained watches. Empty indefinitely waiting epoll instances can therefore be closed from another task.

Shared-memory IPC uses `fd::memfd`, `memory::file(..., MAP_SHARED)`, and descriptor transfer together. Each process owns its descriptor and mapping separately and supplies synchronization for shared contents. Pipes, sockets, counters, timers, poll, and epoll are Linux amd64 only; these wrappers add no grammar.

### URL component and query codecs

`std::net::url` provides pure byte-preserving component/query codecs, lossless reference decomposition/resolution, explicit authority validation and ASCII serialization. It depends on `std::bytes`, not socket, DNS, TLS or HTTP packages. The raw reference type is not a validated network endpoint. Ecosystem `request::Url` reuses query codecs and absolute reference decomposition while retaining its HTTP endpoint, UTF-8, fragment and relative-join policies.

`escape(input: Slice[byte], mode: EscapeMode, limit: isize) -> Result[string, CodecError]` accepts arbitrary bytes and emits ASCII with uppercase percent escapes. `EscapeMode::PathSegment` escapes slashes and other segment delimiters while retaining the path-segment-safe reserved bytes `$&+:=@`; it is not a whole-path encoder. `QueryComponent` escapes all reserved bytes and encodes spaces as `+`. Both retain ASCII letters, digits and `-._~`. `path_escape(string, limit)` and `query_escape(string, limit)` are UTF-8 string conveniences. The limit bounds encoded output bytes and is checked before constructing the output buffer.

`unescape(input: string, mode, limit) -> Result[bytes::Bytes, CodecError]`, `path_unescape` and `query_unescape` accept upper/lowercase hexadecimal escapes and decode each `%HH` exactly once. Query mode turns `+` into space; path mode retains `+`. Decoded bytes need not be valid UTF-8: NUL and `%FF` are preserved, not rejected or replaced. Callers needing text use the checked `utf8::BytesUtf8` conversion. Each successful decode owns independent mutable storage. Its limit bounds decoded bytes, not encoded input length; processing is linear in the supplied input and the current implementation copies its UTF-8 bytes. Decoding does not validate UTF-8, host syntax or reserved-character policy.

`parse_query(input, max_fields, max_decoded_bytes)` returns `Result[Vec[(bytes::Bytes, bytes::Bytes)], CodecError]`. Supply the query itself, without removing or interpreting delimiters inside this function: a leading `?` or `#` is literal query data, not a URL boundary. Parsing splits nonempty `&`-separated fields at their first `=`, with an empty value when `=` is absent. Empty fields are ignored; empty names/values and duplicate names are retained in input order. Raw semicolons are rejected, while `%3B` is valid data. Names and values use query unescaping into separate owned buffers. The two limits bound nonempty fields and the combined decoded name/value bytes. Any failure returns an error rather than Go ParseQuery's partial map plus error; already parsed fields are not exposed.

`parse_query_bounded(input, max_input_bytes, max_fields, max_decoded_bytes)` also checks the raw input byte length before copying or scanning it. It rejects any negative limit first, then reports `InputLimit` at the first excluded input byte when the raw budget is exceeded. This prevents ignored separators from bypassing resource limits. The older three-argument `parse_query` retains its decoded/field limits but requires the caller to bound raw input separately; neither API replaces a complete HTTP request-size policy.

`encode_query(fields: Slice[(bytes::Bytes, bytes::Bytes)], limit)` returns an ASCII `Result[string, CodecError]`. It sorts a copy of the pair sequence lexicographically by decoded name bytes, retaining relative value order for equal names, then emits every pair as `name=value` joined by `&`. Input order and buffers are unchanged. It performs a size preflight before sorting or constructing the output, using checked remaining capacity rather than overflowing length arithmetic. The limit bounds the total encoded bytes; sort workspace is O(fields) in addition to output. Do not concurrently mutate input views or shared name/value buffers during either encoding pass; these functions add no synchronization. An empty pair sequence encodes to an empty string.

All limits must be nonnegative, including on empty inputs. `CodecError::kind()` returns `CodecErrorKind::{InvalidLimit, InputLimit, OutputLimit, InvalidEscape, FieldLimit, InvalidSeparator}`. `index()` identifies the input byte for component codecs and query parsing; query encoding instead reports the original pair index. Invalid-limit errors use index zero. Malformed percent escapes point to `%`; query decode errors use offsets in the complete query. Error kinds support equality and debug formatting, and errors support debug/display formatting. These are ordinary enums, structs and functions without grammar or implicit conversion changes.

```goml
use std::net::url;

fn canonical_query(input: string) -> Result[string, url::CodecError] {
    let fields = url::parse_query_bounded(input, 16384, 100, 4096)?;
    url::encode_query(fields.as_slice(), 16384)
}
```

For example, `canonical_query("b=a+b&a=%FF&b=%2B")` returns `Ok("a=%FF&b=a+b&b=%2B")`. Percent decoding is not sanitization; application code must validate decoded components for its destination context.

### Raw URL references and relative resolution

`url::Reference::parse(input, max_input_bytes) -> Result[Reference, ParseError]` decomposes a reference into immutable raw components. The input byte limit is checked before its byte copy. The parser rejects literal ASCII controls including DEL, invalid scheme syntax before the first path slash, and malformed percent escapes in the path or fragment. It preserves scheme case, percent spelling, Unicode, empty components and repeated slashes. It does not decode delimiters before splitting. `scheme()`, `raw_authority()`, `raw_query()` and `raw_fragment()` return `Option[string]`; `raw_path()` returns a string, and `is_absolute()` tests scheme presence. `None` differs from `Some("")`: `?`, `#`, and `//` retain their explicit empty component markers. Query escapes remain uninterpreted until query decoding. The authority is a raw substring until explicitly validated with `authority()`; malformed hosts and ports can therefore survive raw reference parsing. Raw spaces, Unicode and backslashes are not automatically transformed into a canonical wire URI, so parse success does not authorize using a reference as an HTTP endpoint or request target.

`decoded_path(limit)` returns independent `bytes::Bytes`; `decoded_fragment(limit)` returns `Option[bytes::Bytes]` in a Result, preserving absence. Both use path-style percent decoding, leave plus signs literal and retain non-UTF-8 decoded bytes. Negative limits fail even when a fragment is absent. `ToString` recomposes raw components losslessly after parsing; it does not normalize schemes/hosts, redact credentials or escape raw characters. Use caution before logging references. `PartialEq`/`Eq` compare exact raw component values, not URI equivalence, DNS identity or decoded bytes.

`base.resolve(reference, max_output_bytes)` returns `Result[Reference, ParseError]`. It applies the component inheritance, path merge and literal dot-segment processing of [RFC 3986 section 5](https://www.rfc-editor.org/rfc/rfc3986.html#section-5.2): a reference with its own scheme replaces the base; otherwise it inherits the base scheme, and an explicit authority replaces the base authority. An empty path inherits the base path and, only if no query marker was supplied, its query. Nonempty relative paths merge with the base directory. The resulting fragment is always the reference's fragment, not the base's. A relative reference requires an absolute base; an independently absolute reference may replace a relative base. Rootless scheme paths use generic URI rules rather than Go's scheme-specific opaque-path behavior.

Only complete literal `.`/`..` path segments are removed. Escaped dots and slashes remain escaped, repeated separators are retained, and query/fragment text is not path-normalized. For a result without an authority whose normalized path starts with `//`, the stored path gets a `/.` prefix to keep reparsing from inventing an authority. This preserves the intended path interpretation instead of blindly emitting ambiguous `scheme://...` text. Parsed inputs themselves are never normalized. Empty query and fragment markers survive resolution, unlike APIs that collapse absent and empty values. Unlike Go's treatment of an empty reference/host, a missing reference fragment never inherits the base fragment and an explicit empty authority replaces the base authority instead of inheriting it.

Repeated slashes produced at the root are also retained: resolving `..//` against `https://a` yields `https://a//`, and against `file:///a/b` yields `file:////`. Go 1.26 produces one fewer slash in those two cases; tests record this deliberate difference instead of treating Go output as the sole specification.

Resolution checks the final recomposed byte length without constructing that output string; negative output limits fail first. Temporary path work is linear in the base/reference path sizes, not bounded by the final output limit alone. Bound both parse inputs to bound intermediate work. Returned values retain only immutable strings and do not mutate either input. `ParseError::kind()` returns `ParseErrorKind::{InvalidLimit, InputLimit, OutputLimit, RelativeBase, ControlCharacter, InvalidScheme, InvalidEscape}`; kinds support equality/debug formatting and errors support debug/display. `index()` is a source byte offset for syntax errors, the excluded byte boundary for input/output limits, and zero for invalid limits or relative-base errors. These APIs use ordinary structs, enums, methods and existing string/byte primitives; they introduce no syntax, reflection or runtime backend.

```goml
use std::net::url;

fn resolve_link(base: string, link: string) -> Result[string, url::ParseError] {
    let base = url::Reference::parse(base, 4096)?;
    let reference = url::Reference::parse(link, 4096)?;
    Result::Ok(base.resolve(reference, 8192)?.to_string())
}
```

For example, resolving `../d?` against `https://example.com/a/b?old#base` yields `https://example.com/d?`. This generic reference resolution is independent of ecosystem request's existing HTTP relative-join compatibility rules; consuming the value layer does not require replacing application policy.

### ASCII URL serialization

`Reference::to_ascii(max_output_bytes) -> Result[string, ParseError]` explicitly validates the authority and emits an ASCII URI reference. It retains component boundaries, scheme/host spelling, absent versus empty markers, dot segments, port spelling and credentials. It uppercases hexadecimal digits in existing percent escapes without decoding them; escaped slashes, dots and invalid UTF-8 bytes retain their interpretation. Other disallowed raw bytes, including spaces, backslashes and each byte of Unicode UTF-8, become uppercase `%HH`. Scheme syntax and separators are unchanged. Paths retain unreserved characters, subdelimiters, `:`, `@` and `/`; query and fragment additionally retain `?`. Authority brackets and separators remain structural after authority validation.

The operation checks the complete encoded length before constructing its output buffer, returns OutputLimit at the excluded output-byte boundary, and rejects negative limits first. Temporary scanning storage is bounded by parsed input size, not output size alone. Raw queries are not form-decoded, reordered or validated as name/value pairs, but malformed percent sequences are rejected with InvalidEscape at their byte offset in the complete raw reference. Authority-validation errors retain authority-relative offsets as documented below. `ToString` remains the lossless alternative, including for raw query text that is not a valid wire escape sequence.

For example, `/café?q=a b+#é` serializes as `/caf%C3%A9?q=a%20b+#%C3%A9`; literal plus is not converted to space. Serialization is idempotent after reparsing. This is a canonical escape spelling, not a universal URL-equivalence normalization: it does not remove default ports, lowercase hosts, resolve dot segments, decode unreserved escapes or perform IDNA. Percent-encoded Unicode registered names are generic URI syntax, not a DNS-ready hostname. The result is not automatically an HTTP request target or safe log text, and its credentials are not redacted. Use application policy and explicit redaction separately. This method adds no grammar or implicit conversion.

### URL authorities and credential redaction

`url::Authority::parse(input, max_input_bytes)` validates a standalone authority, without the leading `//`, and returns `Result[Authority, ParseError]`. The input budget is checked before copying or scanning. `Reference::authority()` explicitly validates its raw authority and returns `Result[Option[Authority], ParseError]`; reference parsing itself remains lossless decomposition, not authority validation. Authority error offsets are relative to the authority substring, including user information, even when accessed through a Reference.

`raw_host()` retains brackets around IP literals; `host_kind()` returns `HostKind::{Name, Ipv4, Ipv6, IpvFuture}`. Bracketed IPv6 supports compression and final IPv4 tails; IPv4 components forbid redundant leading zeroes. A non-IP registered name is not validated as a DNS name. IPvFuture uses `v` plus a hexadecimal version, a dot and a nonempty permitted address spelling. `%25` introduces a nonempty IPv6 zone; its raw characters are unreserved ASCII or Unicode, and percent escapes must not decode to ASCII controls. This is a textual zone value, not interface lookup. `raw_zone()` preserves zone spelling without the separator. `hostname(limit)` removes brackets and percent-decodes into fresh Bytes, including a decoded `%zone` suffix; `zone(limit)` returns optional fresh Bytes.

`raw_port()` distinguishes no separator from an explicit empty port. Nonempty ports must contain only ASCII decimal digits. `port_number()` separately returns `Result[Option[u16], ParseError]`, rejecting values above 65535 instead of wrapping; absent and empty ports both return None. Port zero and leading zeroes are accepted. Syntax parsing alone does not impose transport-port range limits.

`userinfo()` returns optional UserInfo. `raw_username()`, `raw_password()` and `has_password()` preserve an absent versus empty password. `username(limit)` and `password(limit)` decode into independent Bytes, retaining literal plus signs and arbitrary percent-encoded bytes. Exactly one raw `@` is allowed; embedded `@` must be percent-encoded, unlike Go's permissive last-`@` splitting. Raw user information is ASCII unreserved/subdelimiter/colon text; Unicode must be escaped. Registered names permit Unicode and generic percent-encoded bytes, unlike Go's restrictions on percent-encoded ASCII hosts. IPvFuture is accepted independently of Go's host policy. None of these operations performs DNS, IDNA, HTTP policy or endpoint validation.

Authority and UserInfo implement raw equality and lossless ToString, but deliberately omit Debug. `Authority::redacted(limit)` and `Reference::redacted(limit)` replace a present password, including an empty one, with `xxxxx`, retaining all other spelling. For example, `https://u:p@host/a?q=secret` becomes `https://u:xxxxx@host/a?q=secret`. Query, path, fragment and username secrets are not removed; this is not a general safe-logging sanitizer. Reference redaction first validates the authority. Negative limits fail, and output byte limits are checked before joining the result. Decode limits bound decoded bytes; bound the original parse input separately.

Additional ParseErrorKind variants are `InvalidHost`, `InvalidPort`, `PortOutOfRange`, `InvalidUserInfo`, `InvalidZone` and `InvalidIpLiteral`; malformed percent sequences use `InvalidEscape`. Diagnostics report kinds and byte offsets without echoing credentials. These APIs use existing struct, enum and method syntax without new grammar or runtime hooks.

### Networking

`std::net` implements IPv4/IPv6 TCP and UDP in GoML using Linux amd64 socket syscalls. It does not delegate networking to Go's `net` package. All sockets are nonblocking and close-on-exec; a lazy shared epoll worker wakes waiting GoML tasks through channels. One epoll descriptor and one worker remain for the lifetime of the process, independent of the number of sockets. The implementation uses one-shot, level-triggered readiness and registration identities that prevent queued events from targeting a reused descriptor. Native layouts and readiness handling follow [socket](https://man7.org/linux/man-pages/man2/socket.2.html), [connect](https://man7.org/linux/man-pages/man2/connect.2.html), [accept](https://man7.org/linux/man-pages/man2/accept.2.html), and [epoll_ctl](https://man7.org/linux/man-pages/man2/epoll_ctl.2.html).

`IpAddr` is `V4([byte; 4])` or `V6([u16; 8])`. `IpAddr::parse(string)` accepts numeric addresses, including compressed IPv6 and an IPv4 tail in IPv6. IPv4 octets use decimal without leading zeroes. `SocketAddr` exposes `ip: IpAddr`, `port: u16`, and `scope_id: u32`; `new(ip, port)` sets scope zero. `SocketAddr::parse` accepts `127.0.0.1:8080`, `[::1]:8080`, or `[fe80::1%3]:8080`. IPv6 scope IDs are numeric; interface names and hostnames are unsupported. Scope must be zero for IPv4. Both address types support equality, debug output, and `to_string()`; IPv6 output uses lowercase hexadecimal with the longest first zero run compressed. Port zero asks the kernel to assign a port when binding; read it with `local_addr()`.

`IpAddr::bit_len()` returns 32 or 128. `is_ipv4_mapped()` recognizes only
`::ffff:0:0/96`; `unmap()` converts those addresses to V4 and preserves all others.
Neither operation changes the original value. Classification methods are
`is_unspecified`, `is_loopback`, `is_private`, `is_multicast`,
`is_link_local_unicast`, `is_link_local_multicast`,
`is_interface_local_multicast`, and `is_global_unicast`, all returning bool.
Private ranges are IPv4 RFC 1918 and IPv6 `fc00::/7`. Classification uses the
embedded IPv4 address for mapped addresses, except `is_unspecified`, which tests
only literal `0.0.0.0` or `::`, and `is_interface_local_multicast`, which is
IPv6-only. Thus `::ffff:0.0.0.0` is not itself unspecified, but its unmapped value
is. Global-unicast classification excludes unspecified, loopback, multicast,
link-local unicast and IPv4 all-ones broadcast; it includes private and
documentation ranges and is not a public-reachability or security-policy test.
These value operations use no DNS, socket or platform calls and add no syntax.

IpAddr implements `std::cmp::{Ord, PartialOrd}`: IPv4 sorts before IPv6, then
unsigned network-order address bytes determine the order. Mapped IPv6 remains
IPv6 for ordering and equality. `next()` and `prev()` return Option[IpAddr],
preserving the family and returning None at its numeric maximum/minimum rather
than wrapping or producing an invalid address. They do not mutate the original.

`as4()` returns Option[[byte; 4]] for IPv4 and IPv4-mapped IPv6, or None for other
IPv6 values. `as16()` returns network-order [byte; 16], mapping IPv4 into
`::ffff:0:0/96`. `IpAddr::from16([byte; 16])` always constructs IPv6, including
mapped IPv6; use `.unmap()` explicitly if IPv4 is wanted. `IpAddr::from_slice`
accepts Slice[byte] of exactly 4 or 16 bytes, returns Result[IpAddr, Error], and
copies the input. `to_bytes()` returns a fresh Vec[byte] of the family's own
width (4 or 16 bytes). Returned arrays/vectors do not alias address storage.
Construct fixed-width IPv4 directly with `IpAddr::V4([byte; 4])`.

IpAddr, ScopedIpAddr, IpPrefix and SocketAddr implement the prelude Hash trait
and can be HashMap keys. Hashing includes every field used by equality: address
family, exact zone text, stored prefix address/length, and socket port/scope.
It does not unmap IPv6, mask prefix host bits or normalize zones. Equal parsed
spellings have equal hashes; different values remain distinct keys even if their
hashes collide. Hash outputs are internal collection details, not cryptographic
digests or a stable wire/storage format. Normalize explicitly with `unmap()` or
`masked()` before insertion when the application wants that key equivalence.

`ScopedIpAddr` composes an existing IpAddr with an opaque, case-sensitive zone
string, without changing the V4/V6 enum or socket layout. `new(address, zone)`
and `IpAddr::with_zone(zone)` return Result: nonempty zones require IPv6,
including IPv4-mapped IPv6. Unlike Go's WithZone, supplying a nonempty zone for
IPv4 is an error rather than silently discarding it. An empty zone means absent.
`ScopedIpAddr::parse(string)` accepts plain IPs or `IPv6%zone`; an explicitly
empty suffix is invalid. The first `%` separates the zone, whose remaining
contents are preserved literally, including further percent signs and Unicode.
No URL decoding, interface-name validation or interface lookup is performed.

`address()` and `without_zone()` return the underlying IpAddr; `zone()` returns
the string. `with_zone(string)` returns a checked replacement without mutating
the original. `unmap()` removes a mapped IPv6 prefix and its zone when converting
to IPv4, preserving other addresses and zones. Equality compares address and
exact zone; Debug/to_string append `%zone` only when present. For example,
`ScopedIpAddr::parse("fe80::1%eth0")?.zone()` is `"eth0"`.
ScopedIpAddr also implements Ord/PartialOrd, comparing IpAddr first and then
zone text in byte order (absent sorts before nonempty). Its `next()`/`prev()`
return Option[ScopedIpAddr] and preserve the zone, including when crossing a
mapped IPv6 range boundary. No operation implicitly changes address family.

`to_socket_addr(port: u16)` accepts absent zones or decimal u32 zones and returns
Result[SocketAddr, Error]; named, signed and overflowing zones fail without
performing network operations. Callers must resolve interface names explicitly.
`ScopedIpAddr::from_socket_addr(socket)` validates the socket's address/scope
combination and converts a nonzero scope to decimal. Scope zero becomes an absent
zone, so numeric conversion deliberately normalizes `"0"` to absent and `"003"`
to `"3"`. Plain scoped-value parsing/formatting preserves those spellings.
Prefixes continue to accept only unscoped IpAddr values; dropping a zone for
prefix matching therefore requires an explicit `address()` or `without_zone()`.

`IpPrefix::new(address: IpAddr, bits: isize)` validates prefix lengths 0..32 for
IPv4 and 0..128 for IPv6, returning Result with the existing network Error type.
`IpPrefix::parse(string)` accepts numeric `address/length` notation; lengths must
be unsigned decimal without leading zeroes (except `0`). Zones, hostnames, missing
lengths and out-of-range lengths are rejected. Both constructors preserve host
bits: `192.0.2.129/24` stays that value until `.masked()` returns `192.0.2.0/24`.
`.address()` and `.bits()` inspect the original value; `.is_single_ip()` recognizes
full-width prefixes. Equality compares address plus prefix length, not just network
coverage. Debug/to_string reuse the existing IpAddr presentation.

`.contains(address)` compares network bits, ignoring the prefix's host bits.
`.overlaps(other)` tests whether two prefixes share any addresses. Both require
the same address family: IPv4-mapped IPv6 stays IPv6 and never implicitly matches
an IPv4 prefix. The zero-length prefix contains every address in its own family.
Prefix fields are private, so public construction cannot produce invalid lengths.
All operations are pure value computations, with no DNS/socket/syscall access or
new syntax. Prefixes carry no zone; numeric socket scopes remain on SocketAddr.

Fallible socket operations return `Result[T, net::Error]`; `Error` and `ErrorKind` re-export `std::io` types. Errors retain operation names and numeric errno when available. Cancellation is `Interrupted`, timeout is `TimedOut`, use after close is `InvalidInput`, and premature TCP EOF in `read_exact` is `UnexpectedEof`. Connection errors not represented by `io::ErrorKind`, such as connection refused or reset, use `Other` with `raw_os_code()`. Address parsing failures are `InvalidInput`.

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

Socket values are shared handles. Multiple reads or multiple writes serialize per socket; one reader and one writer can progress concurrently. Concurrent `close()` wakes active and queued operations, releases the socket once, and is idempotent. Failed binds and connects release their descriptors. Always close sockets explicitly or with `defer`; garbage collection does not close them. IPv6 sockets are IPv6-only, so bind separate IPv4 and IPv6 listeners when both are needed. Unix-domain listeners, HTTP, and other platforms are outside this socket API. DNS and TLS clients are provided by the APIs below. It uses existing imports, enums, methods, channels, and tasks and adds no grammar or compile-time networking.

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

### DNS and TLS clients

`net::resolve(host, port)` and `resolve_with(host, port, context)` return deduplicated numeric `SocketAddr` values using the host resolver; numeric inputs bypass DNS. Empty or NUL-containing names fail with `InvalidInput`. Resolution observes cancellation and deadlines, with a 30-second maximum when the context has no deadline. `TcpStream::connect_host(host, port, context)` resolves once and tries addresses in resolver order under the same context. Supply a deadline to bound the entire connection attempt.

`WaitOptions::with_context(context)` adds parent cancellation and an absolute deadline to TCP/UDP operations. Existing `with_cancel` tokens remain active too; the effective timeout is the earlier of the context deadline and `with_timeout`. Context cancellation reports `Interrupted`, expiration reports `TimedOut`, and a timed-out TCP wait leaves the socket usable.

`std::net::tls` provides `connect(host, port, config)` and `connect_with(host, port, config, context)`. `ClientConfig::new(server_name)` verifies the certificate chain and server name using system roots, requires TLS 1.2 or newer, and limits DNS, dial, and handshake together to 30 seconds. Hosts are bare names or numeric addresses without a port or IPv6 brackets. Configuration builders are `with_ca_pem` (a nonempty PEM bundle replaces system roots; an empty bundle selects system roots), `with_client_certificate(certificate_pem, private_key_pem)`, `with_alpn`, `with_minimum_version(Version::Tls12 / Tls13)`, and `with_connect_timeout`. There is no insecure verification switch. Client certificates and private keys must be supplied together; ALPN identifiers are 1–255 bytes.

`TlsStream::connection_info()` returns the negotiated version and ALPN protocol. `read_with`, `write_with`, `read_exact_with`, and `write_all_with` accept a `context::Context`; exact/all operations hold their direction's gate and keep one deadline across all partial transfers. A reader and writer can run concurrently, while operations in one direction serialize. Cancelling or timing out active TLS I/O closes the connection to interrupt the native operation; create a new connection afterward. Cancellation detected before native I/O begins, including while waiting for a gate, leaves the connection open. Explicit local close reports `BrokenPipe` to interrupted operations; remote EOF remains a successful zero-byte read. `close()` is explicit, shared, idempotent, and wakes blocked operations. Exact reads report `UnexpectedEof` when the peer closes before filling the buffer. Partial data may already have transferred on failure.

```goml
use std::context;
use std::io;
use std::net::tls;
use std::time;

fn connect_example() -> Result[tls::TlsStream, io::Error] {
    context::with_timeout(context::Context::background(), time::Duration::from_seconds(5), |ctx, _| {
        tls::connect_with("example.com", 443, tls::ClientConfig::new("example.com"), ctx)
    })
}
```

### Filesystem ecosystem packages

Recursive directory traversal and filesystem notifications are independent
third-party modules: [`ecosystem::walkdir`](../ecosystem/walkdir/README.md) and
[`ecosystem::notify`](../ecosystem/notify/README.md). They use public standard
filesystem, syscall, task and time APIs. Their sources are no longer bundled
with the toolchain as `std::fs::walkdir` or `std::fs::notify`.

Add dependencies to the module-root manifest and import the new paths:

```toml
[dependencies]
"ecosystem::walkdir" = "0.1.0"
"ecosystem::notify" = "0.1.0"
```

```goml
use ecosystem::walkdir;
use ecosystem::notify;
use std::fs;
use std::time;
```

The public traversal and notification APIs keep their existing names and
behavior. Walkers provide lazy iterative traversal, depth bounds, pruning,
optional link following and contextual errors. Watchers provide recursive
inotify maintenance, event filters, ignored-subtree pruning, multi-path watch
sets, timed/cancellable reads and bounded subscriptions. Both currently target
Linux amd64 and require explicit resource closure when stopping early.
Their complete API examples, lifecycle rules and platform limits are documented
in the module READMEs. Normal package/dependency syntax applies; this move adds
no language grammar. For this checkout, the [ecosystem verifier](../ecosystem/README.md#library-verification)
constructs an isolated registry for these development modules.

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

GoML has no lifetime or linear type system, so a `Scope` value can currently escape its body. Calling `spawn` after the scope begins closing is a runtime error. Panic is not implicitly converted into `Result`. A panic in the scope body or a child task cancels sibling tasks, waits for them, removes the runtime scope, and is then re-raised in the scope owner. An explicit `std::panic::catch` around the scope can recover that re-raised panic after the scope has finished. A catch inside a child task handles its own panic before the task scope observes it. Cancellation remains cooperative, so an uncooperative sibling can delay scope completion and recovery.

`join_all` returns values in input order. `join_all_results` waits for every task and returns errors in input order, independent of goroutine scheduling.

`join_all_indexed_results` preserves each error's input index. `for_each_concurrent(limit, values, body)` runs at most `limit` calls at once, uses one worker when the limit is non-positive, waits for every value, and returns indexed errors in input order. `ConcurrencyLimit::run` can apply the same cooperative limit to custom task layouts.

`race(bodies)` returns the first completed value, cancels the remaining bodies, and still waits for every losing body to exit. It returns `None` for an empty input. A body that does not cooperate with cancellation can therefore delay the return from `race`.

### Panic boundaries

`std::panic` provides explicit recovery from language and host runtime panics:

```goml
use std::panic;

fn validate_internal_state(valid: bool) -> () {
    if !valid {
        panic::raise("invalid internal state")
    }
}

fn guarded() -> Result[(), panic::Panic] {
    panic::catch(|| validate_internal_state(false))
}
```

`raise(message: string) -> never` starts unwinding. `catch[T](body: () -> T) -> Result[T, Panic]` returns `Ok` for a normal callback return, including a callback returning an ordinary `Result::Err`. It returns `Err` only for a panic after all exited blocks have run their registered defers. Bounds errors and other recoverable Go runtime panics, including panics crossing an FFI callback, are caught on the same goroutine. Execution resumes after `catch`, not at the failing expression. This is not a transaction: shared-state mutations and I/O are not rolled back.

`Panic.message()` and `to_string()` return diagnostic text; invalid UTF-8 in host panic diagnostics is replaced with U+FFFD. `stack_trace()` returns the Go stack captured at the first `std::panic::catch` boundary; generated function names may appear. Child task panics are captured before transfer to the scope owner, while a scope-body panic may first be captured after the scope re-raises it. `resume(info: Panic) -> never` rethrows while preserving that diagnostic text and stack in subsequent GoML catches, including across generated Go packages. The original payload remains private and garbage-collected; there is no global panic registry. If a defer panics during cleanup, remaining defers run and the newest panic is delivered to the catch boundary.

An uncaught panic terminates the process. Catch does not intercept another goroutine's panic, `process::exit`, runtime fatal errors, or forced termination. Recovering does not establish that shared application state is still consistent; use boundaries around isolated requests or tasks, and retain `Result` for expected failures. Runtime hooks belong to `lib/builtin`; the public API and consumers are ordinary GoML and need no user Go FFI adapter.

### Cancellation contexts

`std::context` combines cooperative cancellation and monotonic deadlines. `Context::background()` never cancels and has no deadline; `Context::from_token(token)` adapts an existing task token. `with_cancel(parent, callback)`, `with_timeout(parent, duration, callback)`, and `with_deadline(parent, Deadline::after(duration), callback)` invoke `callback(context, cancel_handle)` inside a lexical task scope. A child inherits the earlier deadline, propagates parent cancellation, and cancels when the callback returns. Internal watcher tasks and timers are joined or stopped before the scope exits; an escaped child context is already cancelled.

`CancelHandle::cancel()` is idempotent. The first cancellation cause is retained as `Error::Cancelled` or `Error::DeadlineExceeded`. `Context::check()` returns `Result[(), Error]`; `error()`, `deadline()`, and `remaining()` expose optional state. `done()` is a `Receiver[()]` usable in `select`; `token()` exposes an optional task token for existing APIs. `sleep(duration)` waits cooperatively and reports the cancellation cause. `Deadline` uses `time::Instant`, so wall-clock adjustments do not move it. Cancellation does not forcibly stop arbitrary user code or blocking I/O without context support.

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
| `Simd[T, N]` or vector `a + b` | Import `std::simd`, select a fixed 128/256-bit vector type, and use `a.add(b)` |
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
| `throw`, exception | `Result` and `?` for expected errors; `std::panic::catch` for an explicit runtime-panic boundary |
| `float_value.to_i32()` | Import `std::num::TryToInt` and use `float_value.try_to_i32()`; handle nonfinite and range errors |
| `dyn A + B` | Use one dyn-safe trait; multiple bounds are reserved syntax but not yet supported |
| `dyn TraitWithAssociatedType` | Bind every associated type, for example `dyn Iterator[Item = isize]` |
| Use `type UserId = u64;` when `UserId` must be distinct | Use `struct UserId(u64);` and construct it explicitly |
| `use pkg::*` | List the required public items explicitly with `use pkg::{A, B};` |
| `mod`, `crate::`, `super::` | Directory packages, `module::path` for the current module, and canonical paths for dependencies |
| `fn helper` inside function | Top-level function or local closure |
| Go external type | `#[go_type("pkg", "Name")] extern type Name[T];` retains Go identity and validates concrete instances across packages and artifacts; `std::ffi::Ptr[T]` and Go pointer aliases preserve nullable pointer values, with explicit `ffi::null()` and `ffi::is_nil`; method bindings use `#[go_method("Method")]`; symbolic instances remain unsupported |
| Go interface adapter | `#[go_interface(RawType, Wrapper, method = "GoMethod")]` generates a checked native-interface wrapper and trait implementation; `from_trait` explicitly creates a typed Go bridge retaining the supplied dyn object; nil/typed-nil and multiple results are preserved |
| Go binding generator | `goml bind-go <CONFIG>` selects explicit package/symbol allowlists and finite Go-checked generic arguments; emits raw bindings with protected deterministic output |
| C binding generator | `goml bind-c <CONFIG>` checks C declarations with Clang and emits typed handles, copied strings/buffers, output adapters and compile-time integer constants; supports cgo and a first-party dynamic C ABI backend with `CGO_ENABLED=0` |
| Go-callable export | Annotate a supported public function with `#[go_export("Name")]` and generate a Go package with `goml export-go` |
| Traverse a directory tree | `ecosystem::walkdir` dependency for Linux amd64 syscall-backed depth-first iteration with depth bounds, pruning, optional link following, and per-path errors |
| Manipulate logical slash paths or shell patterns | `std::path::slash` for pure lexical operations and bounded Unicode matching; keep host paths in `std::path` |
| Watch a directory tree for changes | Use the `ecosystem::notify` dependency and its `watch_recursive` or `WatchSet` on Linux amd64, prune ignored paths through `Options`, consume timed reads or scoped subscriptions, handle `Event.rescan`, and close the handle |
| TCP and UDP networking | `std::net` sockets, DNS resolution, shared epoll readiness, explicit close, and context-aware waits; `std::net::tls` for verified TLS clients |
| Escape URL segments or query parameters | `std::net::url` bounded component/query codecs preserve decoded bytes; full URL parsing is separate |
| Resolve a raw URL reference | `url::Reference::parse` and `base.resolve` preserve component markers and remove literal dot segments; authority validation and ASCII serialization are explicit separate operations |
| Inspect a URL host or hide its password | `url::Authority::parse` validates user information, host literals and decimal port syntax; `redacted` hides only passwords, not query or other secrets |
| Keep byte snapshots independent | `Bytes::copy` for mutable copies; `Bytes::freeze` and `FrozenBytes` for immutable snapshots |
| Search binary data repeatedly | `bytes::Finder` snapshots a pattern and reuses linear-time search state; `cut` returns shared views, not copied buffers |
| Apply Unicode rules to binary input | `bytes::fields`, `trim_space` and scalar searches preserve view bytes; `map_chars` emits valid UTF-8; `replace_invalid_utf8` replaces malformed runs once with caller-selected bytes and an explicit output limit |
| Search or traverse text without collecting parts | `text::Finder` uses byte offsets and linear search; `split_iter` preserves GoML empty-separator semantics, and `lines_inclusive_iter` preserves newline bytes |
| Unicode whitespace and case-insensitive text | `text::fields`/`fields_iter` and `trim_space` use pinned Unicode data; `equal_fold` is simple scalar folding, not multi-scalar full folding; legacy `trim` stays ASCII-only |
| Bound generated text | `text::*_checked` construction counts UTF-8 bytes; `map_chars` invokes a stateful callback once per visited scalar and never returns a partial result |
| Encode/decode quoted literals | `text::quote`/`quote_char` emit bounded Go-style literals; `quoted_prefix` recognizes one leading literal; `unquote_element` distinguishes byte/scalar values; `unquote_bytes` preserves bytes and `unquote` validates decoded UTF-8; these are not JSON or GoML syntax extensions |
| Parse a boolean | `num::parse_bool_structured` accepts the documented exact spellings and returns an error for all others; boolean `ToString` emits canonical lowercase text |
| Distinguish floating parse failures | Structured float parsing returns recoverable errors; `ParseFloatError::is_range` identifies overflow, while explicit infinity and rounded underflow remain successful values |
| Format fixed-point floating values | `num::format_float32_fixed`/`format_float64_fixed` use exact decimal conversion, ties-to-even rounding and an explicit output byte limit |
| Parse complex numeric text | `num::parse_complex_f32`/`parse_complex_f64` return typed real/imaginary component tuples with syntax/range errors; no complex source literal or arithmetic type is introduced |
| Format complex components | `num::format_complex_f32`/`format_complex_f64` use typed `ComplexFormat` modes and a limit covering both components, signs and parentheses |
| Reuse multiple literal replacement rules | `text::Replacer` compiles ordered rules with rule/output/work budgets; `chunks` yields fallible fragments lazily; `text::stream::write_replaced` handles short writes and reports partial progress |
| Read from immutable text | `io::StringReader` provides sequential/positional bytes, scalar reads, rollback, reset and checked byte seeking without a writable buffer |
| Position streams through a shared contract | `io::Seek` uses `SeekFrom` and checked byte positions; Cursor/SectionReader remain bounded, StringReader permits beyond EOF, and OffsetWriter does not support End |
| Read and roll back buffered characters | `io::BufReader` reads bytes or UTF-8 scalars, supports one rollback after qualifying reads, and retains confirmed scalar prefixes on I/O errors even with configured capacity one |
| Read bounded stream fragments | `BufReader::read_fragment` and `read_line_fragment` distinguish delimiter/full/EOF boundaries and expose consumed partial data on errors; line payloads omit CRLF while reporting raw consumption |
| Compose buffered output | `BufWriter::write_byte/write_char/write_string/read_from` retain explicit flush and report input acceptance, not guaranteed destination delivery |
| Connect concurrent byte producers and consumers | `io::pipe` uses existing channels; writes wait for reads, close wakes waiters, and `write_progress` retains a confirmed prefix |
| Inject a read-only filesystem | `fs::FileSystem`/`File` use associated handles; bounded read/stat helpers close handles; `SubFs` composes logical names without sandboxing |
| Enumerate portable directories | `ReadDirFile` provides pages with explicit partial errors; `read_dir_from` bounds, validates and stably sorts confirmed entries, then closes the handle |
| Walk portable filesystems | `walk_from` uses an explicit stack, typed skip controls, error callbacks and explicit depth/directory/work budgets |
| Glob portable paths | `glob_from` matches slash-separated components with global budgets, sorted partial results and explicit provider errors |
| Inspect portable symbolic links | `LinkFileSystem` provides static read-link and non-following metadata capabilities; SubFs preserves raw targets |
| Describe portable metadata | `FileMode` preserves advanced kinds and known/unknown flags; SnapshotEntry caches metadata and bounded display helpers avoid implicit I/O |
| Build an in-memory filesystem | `std::testing::fs::MemoryFS` snapshots input bytes, infers directories, resolves links with explicit budgets and opens independent read/seek/directory handles |
| Check a portable filesystem | `std::testing::fs::check_fs` returns bounded, recoverable reports; `check_helpers` adds standard-helper/SubFs comparisons, and static Seek/ReadAt/link checks share the same error and budget model |
| Assemble typed formatted fields | `std::text::format::Formatter` composes bounded text, integer, float and trait-converted fields; f-string syntax remains unchanged |
| Format scientific/significant-digit values | The `num::format_float*_scientific` and `format_float*_general` families reuse exact rounding; precision controls fractional or significant digits respectively |
| Format exact power-of-two representations | `num::format_float*_binary` emits an integer significand with binary exponent; `format_float*_hex` supports exact or rounded hexadecimal fractions |
| Format shortest decimal representations | `num::format_float*_shortest` finds roundtripping significant digits and renders an explicit `FloatNotation`; it does not change default ToString |
| Bound text splitting | `text::split_n_checked` separates count from the result-part limit and decomposes empty separators into scalars; legacy `split` retains the whole string |
| Split or expand untrusted byte data | `bytes::split_n` separates count semantics from a part limit; `replace_n`, `join` and `repeat` check explicit byte limits and return `TransformError` |
| Go `maps` helpers | `std::collections::map_*` over `HashMap`; shallow copies and snapshot iterators, no nil map or implicit Go-map conversion |
| Go `slices` helpers | `std::collections::slice_*` and checked `vec_*` edits plus existing Vec/Slice methods; read-only views have no reslicing capacity or nil distinction |
| Decode arbitrary UTF-8 incrementally | `std::utf8::Decoder` is strict and reports absolute offsets; `decode_rune`, `decode_last_rune` and `decode_lossy` explicitly replace malformed bytes |
| Unicode classification and caseless matching | `std::unicode` supplies pinned category/script/property tables; distinguish one-scalar `simple_fold` cycles from multi-scalar string `case_fold` and locale-specific casing |
| Go `sort.Search` or duplicate boundaries | `collections::search` checks negative lengths; `lower_bound`, `upper_bound` and `equal_range` return insertion points/ranges while `binary_search` retains its first-match `Option` |
| Encode a variable-width integer | `std::bytes::endian::{append_uvarint, append_varint}`; checked reads/writes return `Result` |
| Encode a fixed-layout struct or array | Explicit fields/loops in endian `read_with`/`write_with` callbacks; memory writes stage atomically, stream operations require a byte limit |
| Compute a streaming checksum | Import `std::hash::Hasher` through the package alias and use an Adler32/CRC/FNV digest; checksums are not MACs |
| Hash successful I/O | Wrap a reader or writer with `std::hash::stream`; only successfully reported bytes enter the supplied Hasher |
| Detect a double-width division overflow | Use `std::math::bits::div32/64`, which returns `Result` rather than panicking |
| Preserve action and cleanup errors | `std::resource::{with_cleanup, scope, ScopeError}` or `io::with_resource` |
| Isolate a runtime panic | `std::panic::catch`, with lexical `defer` cleanup; `raise` and `resume` return `never` |
| Generate public inherent methods | `derive_output_inherent` and `derive_output_add_public_method` |
| Create and supervise a Linux child | `std::os::linux::process::Command` with runtime-coordinated spawn, stable pidfd signals, and shared wait results |
| Access mapped memory | `std::os::linux::memory::Mapping` checked copied access with explicit shared unmap state |
| Pass descriptors or wait for Linux readiness | `std::os::linux::ipc` Unix messages, eventfd/timerfd, poll, and epoll |
| Use Linux descriptor operations | `std::os::linux::fd` for explicit shared ownership, scalar/vector I/O, byte paths, metadata, locks, and directory-relative operations |
| Encode a Linux amd64 native record | `std::os::linux::abi` checked byte codecs and `Iovecs`/`Message` builders |
| Treat a GoML integer or struct as a kernel pointer/layout | Use `syscall::Arg::Buffer` and `syscall::Pointer` for synchronous byte-buffer graphs on Linux amd64; encode the native ABI explicitly and inspect `SyscallResult.errno` or `into_result()` |
| Unannotated user `extern fn` | Use a normal GoML function or `#[go_ffi("import/path", "ExportedSymbol")] extern fn`; project commands validate Go calls by default (`--ffi-check required`) |
| Call an ordinary function from `comptime` | Mark a supported free function with `#[comptime]` |
| Capture a runtime local in `comptime` | Pass a literal or compile-time value to a `#[comptime]` function |

## Informal Grammar Quick Facts

Base32, varints, checksums, UTF-8 scalar/stream decoding, map/slice helpers, logical slash paths and fixed-width bit operations use ordinary imports, traits, structs, enums, slices and calls. They add no grammar, implicit byte conversions or compiler intrinsics. Map and slice collection use existing `Iterator` associated-type constraints. Shell patterns are string data parsed by `std::path::slash`, not GoML syntax.

C bindings and `std::c` use existing structs, constants, `#[comptime]`, extern attributes, imports and calls. They add no C pointer, C layout or `extern "C"` grammar.

Linux syscall buffers, `Pointer` descriptors, `Errno` values, ABI codecs, descriptor/process/memory/IPC wrappers use the ordinary struct, enum, array, slice, and call forms below. Native kernel layouts are encoded into bytes; there is no pointer-cast, native-layout, or `unsafe` grammar.

Panic boundaries use ordinary function calls, closures, `Result`, and the existing `never` type. There is no `try/catch`, `throw`, or bare Go-style `recover()` syntax; panic cleanup extends the semantics of the existing lexical `defer` statement.

I/O traits, cancellation contexts, DNS/TLS, scalar conversion traits, and Serde extension events use the existing imports, generic bounds, enums, methods, and calls. They introduce no new grammar. Explicit generic calls retain the `::[Type]` form even when their type arguments do not appear in the function signature.

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
generic_param = upper_ident (":" trait_set)?
param_list    = "(" (parameter ("," parameter)*)? ")"
parameter     = lower_ident ":" type | "self"
return_type   = "->" type

struct_def    = "struct" type_ident type_names?
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
primitive_type = "()" | "never" | "bool" | "isize" | "i8" | "i16" | "i32" | "i64"
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
