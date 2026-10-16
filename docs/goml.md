# GoML language guide

This guide describes the GoML syntax, type rules, tools, and built-in prelude implemented in this repository. Standard-library APIs are documented in the [standard library reference](stdlib.md). Use it when writing, modifying, and reviewing `.goml` source. Similarities to Rust, Go, or OCaml do not imply that their syntax or APIs are supported.

GoML is a statically typed language with garbage collection. Its syntax is close to Rust, while its semantics are closer to ML. The compiler monomorphizes generics and lambda-lifts GoML closures before emitting Go. Explicit Go FFI can preserve native generic types and function values. GoML has no ownership, borrowing, lifetimes, or manual memory management.

Generated Go code and Go FFI use Go 1.26 as their compatibility baseline. Building the toolchain and compiling generated programs require Go 1.26 or newer. Native FFI signatures are checked by that final Go compilation. The independent [gomlgo frontend and interpreter](https://github.com/gomlang/gomlgo) targets Go 1.26 language and standard-library behavior; its execution and differential tests require Go 1.26.x.

Quick navigation: [packages](#modules-packages-and-imports), [types](#type), [control flow](#control-flow), [patterns](#patterns), [traits](#trait-and-impl), [compile-time evaluation](#compile-time-evaluation), [Go FFI](#go-ffi), [tests](#tests), [API documentation](#comments-and-api-documentation), [prelude](#built-in-prelude), [standard library](stdlib.md), and [verification](#verify-generated-code).

Build and installation instructions are in the [development guide](development.md) and [release guide](releasing.md#installation). See [formatting](formatting.md) for canonical source layout and [compile-time evaluation architecture](comptime.md) for CTIR internals.

## Essential rules

1. Source files use `.goml`, and internal test files use `*_test.goml`. Legacy `.gom` files must be renamed before compiling or formatting; the project root directory uses `goml.toml` to declare the canonical module path.
2. Each source file in the project first writes `package name;`, then writes the file's own `use`, and finally writes the top-level definition.
3. Parameters of top-level functions must have types; the return type is fixed to `()` when omitted and is not inferred from the function body.
4. Generics use square brackets: `Vec[i32]`, `fn id[T](x: T) -> T`, not `<...>`.
5. Generic calls usually rely on type inference; when explicit type arguments are required, write `id::[i32](1)`, not `id[i32](1)` or Rust's `id::<i32>(1)`.
6. `if`, `match`, and `select` are expressions. `if` without `else` must return `()`; `match` must be exhaustive.
7. The last semicolon-free expression of a block is the block value; adding a semicolon discards the value. Directly discarding a built-in `Result` warns; use `let _ = ...` to explicitly ignore it.
8. `let` and assignment statements must end with a semicolon. Mutable bindings may be introduced with `let mut pattern` or precisely inside a pattern with `mut name`; the semicolon can be omitted for `if`, `match`, `select`, and unlabeled `while`, `loop` and `for` statements. Labeled loop statements require a semicolon before a following statement.
9. Enumeration construction uses full names such as `Option::Some(value)`, or contextual names such as `Some(value)` and `None` when the expected type determines the enum. Patterns also permit contextual variant names.
10. For cross-package calls, write `alias::item`. Top-level items, struct fields, and inherent methods must all be marked with `pub` as required.
11. Before using trait method syntax across packages, import the package and trait with `use alias::Trait;` or a braced import; when in doubt, use UFCS: `Trait::method(value)`.
12. Use `module::path` for a package below the current module root. Do not generate `mod`, `crate::`, `self::`, `super::`, root paths `::x`, Rust references, or Go `var` / `:=`. A user `extern fn` is valid only with the typed Go FFI attribute described below.
13. The test function uses `#[test]`, which must have no parameters, no type parameters and return `()`; the white-box test is placed in `*_test.goml` of the same package, and the black-box test is placed in the `tests/` directory of the package under test.

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

### Comments and API documentation

Line comments start with `//` and end at the next newline. `///` (excluding `////`) at the start of a source line documents the following declaration. Consecutive documentation lines form one block; a blank line or an ordinary comment breaks attachment. Attributes belong to the following declaration and may appear between its documentation and declaration keywords.

`//!` comments before the `package` declaration describe the package. For a package with several source files, `gomldoc` combines these descriptions in relative file-path order. Documentation comments remain ordinary lexer trivia and do not affect program typing or execution.

```goml
//! Small arithmetic helpers.
package arithmetic;

/// Add one to `value`.
pub fn increment(value: isize) -> isize {
    value + 1
}
```

`goml doc` generates offline HTML for all production packages in the current module, defaulting to `<target-dir>/doc/`. `goml doc --format json` writes the versioned API model to `module.json` in that directory. Public declarations and members follow the language visibility rules; `--document-private-items` includes private API. The command checks the module and reports dependency or typing failures before replacing existing output.

Documentation supports paragraphs, headings, flat bullet lists, inline code, fenced code blocks, and links. `[Name]` and `[label](goml:alias::Name)` refer to API declarations using the source file's imports. Ordinary links support HTTP, HTTPS, mailto, and fragment destinations. Raw HTML is escaped. Unresolved or ambiguous API links produce warnings; targets outside the generated module are shown as text. Full Markdown, automatic dependency sites, and executable documentation tests are not supported.

See [documentation](documentation.md) for standalone `gomldoc`, output schemas, and installation.

### Identifiers and keywords

Ordinary identifiers only use ASCII and are of the form:

```text
[A-Za-z][A-Za-z_0-9]* | _[A-Za-z_0-9]+
```

The single `_` is a wildcard character, not an ordinary variable name. The current syntax enforces the first letter case of name categories:

- Package names, package aliases in `use package as alias`, functions, methods, parameters, local bindings and GoML field declarations must start with a lowercase letter or `_`; imported item aliases may follow the naming convention of the imported item;
- Structures, enumerations, traits, enumeration variants, generic parameters, and associated types must start with a capital letter;
- The fixed-width vector and mask type spellings listed under [Portable SIMD](stdlib.md#portable-simd), including `u8x16`, `i32x8`, `f32x4`, `f64x4`, and `mask32x4`, are also accepted as structure and type-alias names and in type positions. They are ordinary scoped names, supplied by `std::simd`, rather than globally reserved keywords;
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
"alice::http" = true
```

Registry dependencies declare presence with `true` (an empty inline table is also accepted). Packages currently follow their repository default branch; there are no package versions, version constraints, minimum-version selection, or `goml.lock`. Existing strict `X.Y.Z` strings and inline `version` fields remain readable during migration, but do not select or constrain a source. New declarations should use `true` or a local `path`.

`[dev-dependencies]` uses the same unversioned and path forms. These dependencies are available to the invoking module's `*_test.goml` files, `tests/` packages and root `examples/` tree. Production sources cannot import them, even during a test build. Dependencies of another module do not inherit its development dependencies. Test and example commands resolve the union of normal and development dependencies; ordinary check/build/run commands resolve only normal requirements. A development helper may depend back on the current library: that edge uses the current source without introducing a second copy. Production dependency cycles remain errors. A helper that imports the package under test must be used from black-box tests or examples; importing it from that same package's white-box tests still creates an unsupported package cycle.

```toml
[dev-dependencies]
"alice::proptest" = true
"alice::test_support" = { path = "../test_support" }
```

Use `goml add alice::proptest --dev`, `goml add alice::test_support --dev --path ../test_support`, and `goml remove alice::proptest --dev` to edit this section. Production dependencies stay in `[dependencies]`.

Local development dependencies can use an inline table. Paths are relative to the declaring module, and the dependency name must match the target's canonical `[module].path`:

```toml
[dependencies]
"alice::http" = { path = "../http" }
"alice::model" = { path = "../model" }

[replace]
"alice::codec" = { path = "../codec" }
```

`goml add alice::http --path ../http` validates the target and writes a module-relative path. Local modules use their working sources, including transitive local dependencies, without requiring a registry cache. Registry modules cannot declare local path dependencies. Only the invoking module's `[replace]` table applies to its whole dependency graph; dependency-local replacements are ignored. Conflicting local sources, mismatched identities and dependency cycles are errors. Artifacts stay under the consuming project's target directory, and local source edits invalidate cached outputs.

A `goml.work` file groups independent modules:

```toml
[workspace]
members = ["app", "http", "model"]

[replace]
"alice::codec" = "../codec"
```

Members retain their own manifests and file-scoped imports. Membership provides a local source for a declared dependency; it does not make undeclared dependencies visible. Selection prefers workspace replacements, invoking-module replacements, workspace members, explicit path dependencies, then registry sources. Member paths and replacement paths are relative to `goml.work`; duplicate canonical member paths or module identities are rejected. An ancestor workspace applies only to its listed members. `GOML_WORKSPACE` selects a workspace directory explicitly, and `GOML_WORKSPACE=off` disables discovery.

`goml check --workspace`, `build --workspace`, `test --workspace`, `fmt --workspace`, `doc --workspace` and `clean --workspace` run for each member in declaration order, stopping on failure. From a workspace directory without an enclosing module, these commands select all members automatically. `-p alice::app` or `--package alice::app` selects one module; `goml run` requires this selection when multiple members exist. Commands invoked inside a member operate on that member unless `--workspace` is supplied. Compiler queries and LSP dependency navigation resolve the same local sources and replacements.

`[build]` can be omitted; `build.target-dir` defaults to `_artifact` under the module root. The manifest value must be a non-empty relative path and cannot contain a `..` segment. `goml check`, `goml build`, `goml run`, and `goml test` can temporarily override it with `--target-dir <path>`; command-line overrides may be relative or absolute. `goml clean` removes the configured target directory. Its optional `--target-dir <path>` override must stay inside the module.

Each module path segment must be non-empty and may contain ASCII letters, digits, `_`, and `-`. Paths rooted at `builtin` or `prelude` are reserved for the toolchain and cannot be used as a module path or dependency. The `[module]` section currently has no `name`, `kind`, `root`, or similar fields.

### Package registry

The default central index is [github.com/gomlang/registry](https://github.com/gomlang/registry). Its `index.toml` maps canonical module names to Git repositories in the `gomlang` organization. A module's canonical name comes from its root manifest, so `ecosystem::ansi` resolves to `github.com/gomlang/ansi` without changing its import path.

```sh
goml add ecosystem::ansi
goml add ecosystem::proptest --dev
goml update
goml remove ecosystem::ansi
```

`goml add` writes an unversioned dependency and fetches required sources. `goml update` refreshes the index and the current module's normal and development dependency closure, honoring local paths and workspace replacements. Outside a module it updates only the index. Dependencies' own development dependencies are not fetched. Builds and compiler/LSP queries use the cached sources without network access; a missing cache is a recoverable error directing the user to `goml update`. Updated source contents invalidate generated artifacts.

The default home is `~/.goml`, overridden by `GOML_HOME`. The index checkout lives in `<home>/cache/registry`; Git package checkouts live in `sources/<owner>/<module>/` below that directory. Each checkout follows its default branch. Updates use fast-forward pulls and do not reset local changes. Changing a cached repository's source produces a diagnostic asking the user to remove that cache directory before updating.

A Git index entry is:

```toml
[modules."ecosystem::ansi"]
git = "https://github.com/gomlang/ansi.git"
```

For local registry fixtures, an entry can instead use `path = "owner/module"`, relative to the index root. Exactly one of `git` and `path` is required. Paths must stay inside the registry tree. Every source must contain a root `goml.toml` whose module path matches its entry. There are no version directories or synthetic snapshot versions.

Set `[registry].default` in `<home>/config.toml` to override the index URL; a missing or empty setting uses the central index. `goml update --local-registry <path>` and `goml add ... --local-registry <path>` use a local Git index. `goml remove` edits the manifest without contacting the index. The previous versioned index format is unsupported; use a fresh cache or remove the old registry cache before switching to this index. Dependency declarations and index entries are configuration and do not change the source-language grammar.

### Examples

Named examples share the module-root manifest and use both normal and development dependencies. Put each executable in `examples/<name>/main.goml`, declaring `package main;` and `fn main() -> ()`. Names contain ASCII letters, digits, `_`, or `-`. An example can contain helper packages, white-box tests and direct `tests/` packages. An example directory containing its own `goml.toml` is an independent module, not a named example.

```text
library/
├── goml.toml
├── library.goml
├── tests/api_test.goml
└── examples/basic/
    ├── main.goml
    └── tests/example_test.goml
```

For a library with module path `alice::library`, the example package is `alice::library::examples::basic`. Production packages and other modules cannot import example packages. Examples use the existing GoML source grammar; no target declarations or extra manifests are needed.

```sh
goml run --example basic -- argument
goml build --example basic
goml check --examples
goml test --example basic
goml test
```

`check`, `build` and `test` accept `--example <name>` or `--examples`; these options select only the named example or all examples. `run` accepts one `--example` and cannot also take a positional target. Ordinary check/build commands exclude the root `examples/` tree. Ordinary `goml test` runs module tests, builds all examples and runs their tests; it does not execute example `main` functions. `goml fmt` includes examples and their tests. The binary for `basic` is `<target-dir>/bin/examples/basic/basic`. Test processes run from the module root, including example tests.

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

The consuming project supplies a minimal module-root `go.mod`. For modules with declared native adapters, the driver resolves those adapters from the selected GoML registry sources and generates Go requirements and local replacements under `<target-dir>/native/<hash>/`. It passes the generated module file through `GOFLAGS=-modfile=...` to compilation, tests, and Go export. This managed module-file path leaves the source `go.mod` and `go.sum` files unchanged. The content hash includes adapter locations and module metadata; dry runs do not run prerequisite probes or create these files. An existing `GOFLAGS -modfile` conflicts with managed dependencies and produces a diagnostic.

The driver checks required `CGO_ENABLED=1` and the requested LLVM major using `LLVM_CONFIG`, `llvm-config-N`, `/usr/lib/llvm-N/bin/llvm-config`, then `llvm-config`. These declarations do not install system packages or configure compiler/linker search paths. Consumers still configure `CGO_CFLAGS`, `CGO_LDFLAGS`, and the shared-library loader when needed. Additional Go dependencies are managed by the consuming Go module and the user's Go environment; dependency-local replacements are not promoted into the consumer. Ordinary project builds use normal Go dependency resolution. Go library export and native SIMD builds retain readonly, offline module handling and require dependencies to be prepared beforehand. Registry packages must include their adapter sources and Go module files.

### Build output layout

The default product layout is as follows; the standard package paths in the path are expanded by directory level, for example, `alice::myapp::utils` corresponds to `alice/myapp/utils`:

```text
_artifact/
├── check/
│   ├── pkg/<canonical-package-path>/<last-path-segment>.interface
│   └── deps/<owner>/<module>/pkg/<canonical-package-path>/<last-path-segment>.interface
├── build/
│   ├── pkg/<canonical-package-path>/<last-path-segment>.interface
│   ├── pkg/<canonical-package-path>/<last-path-segment>.core
│   ├── pkg/<entry-package-path>/goml_generated.go
│   └── deps/<owner>/<module>/pkg/<canonical-package-path>/<last-path-segment>.*
├── test/
│   ├── base/
│   │   ├── pkg/<canonical-package-path>/<last-path-segment>.*
│   │   └── deps/<owner>/<module>/pkg/<canonical-package-path>/<last-path-segment>.*
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

The executable file of the module root entry package is `bin/<module name>`; the nested entry package retains the directory within the module and appends the entry name. For example, the output of `alice::app::cmd::server` is `bin/cmd/server/server`. Production test dependencies are built once under `test/base`; internal and external tests each share one Go entry, manifest, and runner at the corresponding test-kind root. The runner file has the `.exe` suffix on Windows. External dependencies use the same `deps/<owner>/<module>/pkg/...` structure in the root directory of each stage.

The configured production directory will not participate in package discovery and cannot be a target to be inspected, built, or tested. The `.gitignore` generated by `goml new` contains `/_artifact/` by default; after modifying `build.target-dir`, the project's own `.gitignore` should be modified simultaneously.

### Directory packages

Each directory containing a `.goml` file is a package. All source files in the same directory must declare the same package name:

```goml
package utils;
```

The canonical identity of a package is the module path plus the relative directory path. The package declaration name determines the local name when imported without aliases, but not the global identity of the package. The canonical identity of the root package is the module path.

A typical project could be:

```text
goml.toml
main.goml
utils/utils.goml
```

`utils/utils.goml`:

```goml
package utils;

pub fn message() -> string {
    "hello"
}
```

`main.goml`:

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

GoML has no Rust reference or lifetime syntax, pointer arithmetic, slice literals, or union types. Use `Ref[T]` for shared mutable storage, `Option[T]` for optional values, and `Slice[T]` or `MutSlice[T]` for read-only or mutable contiguous views. `std::ffi::Ptr[T]` carries nullable Go pointers and remains distinct from `Ref[T]` and numeric types. `ffi::null()` and `ffi::is_nil()` operate on this explicit pointer type. A type declared with `extern type` remains opaque even when its Go declaration aliases a pointer; it does not gain these pointer operations automatically.

On Linux amd64, `std::os::linux::syscall` accepts numeric machine words, scoped byte-buffer arguments, and explicit pointer fields between byte buffers for low-level kernel calls. It does not add pointer casts, pointer arithmetic, or native struct layout to the language.

`A + B` combines trait bounds in generic constraints, supertrait lists, `dyn A + B` objects, and `impl A + B` returns. It is not a general intersection type for concrete values.

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

`#[comptime]` marks a non-generic free function or concrete implementation method as compile-time-capable. The function remains callable at runtime. A compile-time call may call other `#[comptime]` free functions, the builtin operations listed below, or the `compile_error(string) -> never` intrinsic. The compiler validates the complete body of every marked function, including branches not taken by a particular invocation. Attributes with arguments, duplicate attributes, generic functions or implementations, extern functions, and other declarations are rejected.

A top-level constant initializer is an implicit compile-time context, so `const SIX: isize = factorial(3);` and an initializer wrapped in `comptime { ... }` are equivalent. A top-level constant may produce a recursively immutable tuple, fixed array, struct, or enum in addition to scalar values. A `comptime` expression in ordinary code supports the same reifiable value shapes.

Compile-time code may use local bindings and assignment, blocks, `if`, `match`, `while`, `loop`, restricted `for`, `break`, `continue`, `return`, recursion, direct calls, integer conversion methods, integer `to_string()`, and supported operators. A compile-time `for` accepts only a fixed array or the builtin `isize` ranges `start..end` and `start..=end`; its source and range endpoints are evaluated once, and its pattern must be irrefutable. The deterministic string methods `len`, `byte_len`, `get`, `byte_get`, `byte_slice`, `is_char_boundary`, `starts_with`, `ends_with`, and `contains` are also available. String indexes and slices use byte offsets and reject invalid UTF-8 character boundaries. Integer formatting accepts signed and unsigned 8-, 16-, 32-, and 64-bit values, including `byte`, `isize`, and `usize`, and produces decimal text with the full range preserved. This is restricted to the builtin `ToString` implementation; a user trait with the same short name does not become a compile-time intrinsic. For example, `const BITS: u64 = 18446744073709551615; const MASK: string = BITS.to_string();` evaluates without runtime formatting.

Compile-time code cannot capture a surrounding runtime parameter or local. Closures, function parameters and indirect calls are supported within one evaluation, including shared mutable captures. Concrete `#[comptime]` methods support static calls and dynamic dispatch through local implementations. `Ref` supports `new`, `get`, and `set`; `Vec` supports `new`, `with_capacity`, `from_array`, `get`, `set`, `push`, `len`, `is_empty`, `clear`, and `truncate`. These operations retain ordinary aliasing and bounds checks. Function values, dynamic objects, references, and vectors cannot escape as runtime constants. Generic compile-time functions and implementations, general iterators, floating-point computation, `HashMap`, channels, goroutines, extern calls, host I/O, environment access, time, randomness, network access, general type reflection, arbitrary declaration generation, and general type-level computation are not supported. The constrained programmable derive interface described below is the only reflection and code-generation facility.

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

An explicit generic argument list must contain at least one argument. Omit the whole list to use inference; `identity::[](1)` is rejected.

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

## Blocks, bindings and assignments

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
| Enum tuple variant | `Some(value)` or qualified `Option::Some(value)` |
| Enum struct-like variant | `Message::Named { value }` or contextual `Named { value }` |
| Array, Vec, Slice or MutSlice | `[first, second]`, `[first, .., last]` |
| rest binding | `[first, middle @ .., last]` |
| Alias | `whole @ Some(value)` |
| or-pattern | `Color::Red \| Color::Blue` |
| Range | `0..10`, `'a'..='z'` |
| Nested pattern | `Ok((key, value))` |

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
    whole @ Some(inner) => use_both(whole, inner),
    None => fallback(),
}
```

`@` is more tightly bound than `|`. Therefore `whole @ A | B` resolves to `(whole @ A) | B`, and usually an error will be reported because the variables bound to the two alternatives are different. To make the alias cover the entire or-pattern, you must write parentheses:

```goml
whole @ (Left(value) | Right(value))
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

Trait implementation and method candidate searches do not retain type bindings from rejected candidates. Reordering otherwise equivalent implementation declarations does not select a different inferred type.

Trait bounds make the corresponding methods available on generic values. Imported short trait names, renamed imports, and public re-exports retain the defining trait identity in generic bounds and `where` predicates through specialization. Supertraits and associated type bounds also participate in method resolution as implied constraints.

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

All fields or variant payloads used by a derive must support that trait. A `ToString` variant template only uses the fields referenced by its placeholders. For a generic definition, the generated impl constrains each distinct participating field or payload type that mentions a type parameter. An unused phantom parameter receives no constraint. Compiler-owned runtime, intrinsic, and lang-item attributes remain unavailable to ordinary projects. User projects may use the `go_ffi`, `go_type`, `go_method` and `go_interface` attributes described in the next section.

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

The standard library provides serialization derives through `std::serde`. See [Serde and JSON](stdlib.md#serde-and-json) for their APIs, attributes, and examples.

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

## Go FFI

The runnable [file reading fixture](../goml/testdata/ffi/file-read/README.md) combines raw pointers, methods, shared byte buffers, partial reads, error identity, explicit Close and Go↔GoML result transport.

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

Bind an opaque Go type with `#[go_type("time", "Duration")] pub extern type Duration;`. The GoML name may differ from the Go object name. Declarations cannot have a body or an `= Type` target. Import path, Go name and type arguments identify the type: `Duration` differs from `i64`, while two GoML packages binding `time.Duration` share the same type. Values can pass through ordinary GoML functions and explicitly typed foreign calls; generated Go uses the original Go type. The compiler does not inspect external fields, normalize Go aliases or discover methods and interfaces. External field access and struct construction are unavailable; use explicit function or method bindings, or a Go shim.

```goml
#[go_type("time", "Duration")]
extern type Duration;

#[go_ffi("time", "Sleep")]
extern fn sleep(duration: Duration) -> ();

fn wait(duration: Duration) -> () {
    sleep(duration)
}
```

For a Go `type Point struct { X, Y int64 }`, bind explicit helpers such as `NewPoint` and `PointX` from its Go package:

```goml
#[go_type("example.com/geometry", "Point")]
extern type Point;

#[go_ffi("example.com/geometry", "NewPoint")]
extern fn point(x: i64, y: i64) -> Point;

#[go_ffi("example.com/geometry", "PointX")]
extern fn point_x(value: Point) -> i64;
```

`use std::ffi;` exposes `ffi::Ptr[T]`, a transparent spelling for a raw Go pointer to T. `ffi::null()` constructs a nil pointer with its type inferred from context, and `ffi::is_nil(value)` tests it. Non-null pointers enter through explicit Go bindings. For example, `let empty: ffi::Ptr[Cell] = ffi::null();` declares a nil pointer. These functions are also ordinary function values. Passing a pointer through functions, closures or containers does not copy the pointed-to object. `Option::Some` containing nil remains distinct from `Option::None`. GoML cannot construct a pointer from an integer or a `Ref`. A separately declared Go alias such as `extern type Handle` remains opaque and is not normalized to `ffi::Ptr[Cell]`. It cannot be passed to `ffi::is_nil()` or receive a value from `ffi::null()`; use the explicit pointer representation in binding signatures or bind Go helper functions for the opaque type. There is no `*T` source type syntax.

Bind a Go method with `#[go_method("Method")] extern fn`; its first parameter is the receiver. Declarations must be monomorphic and have exactly one ABI attribute. The method name must be an exported ASCII Go identifier. The final Go build checks the receiver's actual method set: pointer-only methods require an appropriate pointer receiver, while value methods can also be called through pointers. There is no implicit address-taking of a copied value. Interface and promoted methods follow the same Go rules. Bindings can be public, imported across packages, or passed as ordinary function values, and are runtime-only, including when loaded from artifacts.

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

`ffi::Error` is the original nullable Go `error` interface. `ffi::nil_error()` constructs a nil interface; `ffi::error_is_nil(value)` tests interface nilness. An interface holding a typed nil pointer is not nil. Errors retain their original object and chain through functions, closures and artifacts. Raw `(T, error)` returns remain tuples; `Option::Some` containing a nil error remains distinct from `None`.

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

Generic external declarations such as `#[go_type("example.com/shim", "Box")] extern type Box[T];` accept applications such as `Box[i64]` and symbolic `Box[T]` in generic GoML code. GoML checks the declared arity and retains the external identity and arguments across packages and artifacts. It does not query the Go declaration, inspect constraints or expand aliases. Final Go compilation checks reachable concrete instances after specialization, including whether the real type and arguments are legal runtime values. Unused generic functions may produce no Go instantiation. Query and LSP analysis use the same declared opaque types.

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

The declaration must be monomorphic. GoML checks the declared boundary shape and trusts the explicit signature. `goml check` and editor analysis do not inspect the native package. The final Go build validates emitted calls against actual symbols, receiver method sets, assignability and generic constraints. Incorrect declarations fail at that build step; unused bindings may not be emitted and therefore need not be checked. A Go named type such as `time.Duration` remains distinct from a GoML `i64` even when their underlying representations match.

The final executable build uses the selected Go executable, target, flags, module policy and generated caller directory. Go checks import legality, including `internal` restrictions. Ordinary builds disable Go workspace mode and select module mode from the presence of the module-root `go.mod`; other Go dependency and toolchain behavior follows the user's Go environment. They impose no additional FFI-specific readonly or offline policy. Go library export and native SIMD builds retain readonly, offline module handling, so prepare their dependencies separately. A successful Go build does not prove Unicode, mutability, lifetime or concurrency contracts. A unit-returning legacy binding may discard Go results; ordinary calls retain Go's generic inference and variadic-call rules.

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

The first argument is the raw type path, which may name an imported external type. The second is an unqualified wrapper name; its visibility follows the trait. Named string arguments map every declared trait method to exactly one exported Go method. GoML checks complete mapping of its trait, duplicate or unexported names, and the declared raw boundary types. Generated type/helper names and reserved wrapper members cannot collide with user declarations. The final Go build checks the real native interface and method signatures; GoML does not discover its method set.

`ReaderAdapter::from_go(raw)` stores the original interface value in a private field; `wrapper.into_go()` returns it. These preserve nil interfaces, typed-nil receivers and the identity of referenced Go objects. Forwarded calls retain Go receiver behavior, including nil-receiver handling and panics. No receiver is replaced and no automatic close is added. The raw Go interface remains distinct from `dyn Reader`; use `wrapper as dyn Reader` for an explicit GoML trait conversion.

Method signatures use explicit raw boundaries: `ffi::String`/`ffi::Rune` for text and `ffi::RawSlice` for slices. Apply copy/share adapters explicitly at the caller. Every Go result is preserved: `()` means no results, and `(isize, ffi::Error)` retains both a partial count and its original error. Ordinary GoML text, read-only slices, implicit closures and runtime containers are not automatic interface method conversions.

`ReaderAdapter::from_trait(value: dyn Reader)` creates a wrapper backed by a generated Go bridge. Use `ReaderAdapter::from_trait(implementation as dyn Reader).into_go()` to pass a GoML implementation to Go. Typed forwarding closures retain the supplied trait object after the creating function returns. Go calls preserve shared captured state and every return component; panics propagate in the same goroutine. The generated bridge includes a compile-time Go interface-satisfaction assertion. The resulting interface is non-nil, including when the supplied trait object wraps a nil raw interface. `from_go`, `into_go` and `from_trait` are reserved wrapper method names. Subscription removal and resource closing remain explicit caller responsibilities.

The adapter requires a raw external type and a non-generic GoML trait with instance methods, without supertraits, associated types or additional method predicates. Its explicit signatures and mappings supply the information needed for bridge generation. The final Go build checks that the raw type is a usable interface and that the generated bridge satisfies it, including embedded methods and signature compatibility. Constraint-only interfaces cannot be used as Go runtime values.

### Raw Go function values

`ffi::Func[F]` is a nullable raw Go function value, distinct from a GoML closure with signature F. For example, `ffi::Func[(i64) -> (i64, ffi::Error)]` represents Go `func(int64) (int64, error)`, and `ffi::Func[() -> ()]` represents `func()`. F must resolve to a function signature with FFI-safe parameters and results; functions cannot be raw map keys.

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

The runnable [callback subscription fixture](../goml/testdata/ffi/callbacks/README.md) provides explicit unregister, gated asynchronous dispatch, concurrent atomic updates, recursive reentry and registration stress coverage. Its callback reference is removed by unregister; already selected invocations may still run, so callers wait for their batches before disposing of captured resources. GC reachability and registration lifetime are separate. The fixture and its generated Go program are tested with the race detector. GoLibrary function-value exports remain unsupported.

Executable, test and library module links embed a versioned JSON index in the generated Go string constant `_goml_source_origins`, relating surviving functions and callback adapters to enclosing GoML declarations. Shared adapters may have several source entries. External-call wrappers also identify the actual extern declarations used, including `go_method`; unused aliases of the same native target are excluded. The index records canonical symbols, the source package, package-relative source files and declaration-name positions, including explicitly defined impl methods; it is not an expression-level stack-trace line table.

The index has `version: 1`, `scope: "enclosing-declaration"` and an `entries` array. Entries contain `generated_symbol`, `source_symbol`, `source_package`, `source_file`, `start`, `end`, `line` and `column`; callback methods use `ReceiverType.Apply` as their generated symbol. Decode the Go string literal before parsing the JSON.

### Raw Go strings

`std::ffi::String` represents arbitrary Go string bytes and is distinct from the GoML Unicode `string` type. It can cross foreign calls, package interfaces, generic functions and captured closures without changing its bytes. Neither implicit conversion nor direct construction is supported.

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

Legacy foreign bindings declared with GoML `string` retain their existing trusted text contract. Use `ffi::String` when a Go producer can return arbitrary bytes. GoLibrary exports may use `ffi::String` to select the raw-byte policy. Ordinary GoML `string` and `char` exports remain unsupported; use explicit checked adapters at the boundary.

### Raw Go runes

`ffi::Rune` is a transparent alias of `i32`, matching Go's `rune`. Raw values may be negative, surrogate code points or larger than the Unicode scalar range. `ffi::rune_from_char(char) -> ffi::Rune` preserves a valid character. `ffi::rune_char(ffi::Rune) -> Result[char, ffi::BoundaryError]` checks the scalar range and returns `InvalidRune(original_value)` for invalid input. It never substitutes a replacement character. These adapters use ordinary type alias and function syntax.

### Raw Go slices

`ffi::RawSlice[T]` is an alias of `MutSlice[T]` and directly represents a Go `[]T` header. Header assignment copies length, capacity and the backing-storage reference; it does not copy elements. Existing checked indexing and subslicing methods remain available.

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

GoLibrary slice exports remain unavailable.

### Raw Go maps

`ffi::RawMap[K, V]` represents Go `map[K]V` directly and is distinct from GoML `HashMap[K, V]`. It preserves nil and shared storage through assignments, foreign calls, package interfaces, generics and captured closures. Go operations on an alias observe the same entries.

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

The final Go compiler enforces Go's map-key comparability rules. GoML checking rejects known non-comparable keys, including slices, raw maps, arrays or tuples containing them, and local structs with such fields; inferred generic applications are checked after inference. Repeated local generic wrappers are checked at each concrete instantiation: `Key[Key[RawSlice[i64]]]` is rejected when `Key[T]` stores `T` directly, while wrappers that only store a pointer to `T` remain comparable. Recursive struct layouts and structural checks exceeding 64 nested struct instantiations produce recoverable diagnostics. A pointer to an otherwise non-comparable type remains a valid key. Public GoML struct interfaces retain a comparability summary for private fields, so downstream key checking observes their restrictions without exposing field names or types. Summaries may depend on generic arguments; unknown hidden structure does not establish comparability. The artifact decoder also rejects directly non-comparable key shapes. External named keys are opaque; the final Go build checks their actual fields and concrete type arguments for comparability. No native field structure or constraints are imported into GoML analysis. Module checking infers internal comparability requirements from generic function bodies, including requirements reached through function values and nested closures. Package checks propagate them across files until signatures stabilize, with a recoverable limit of 256 passes. Exported function interfaces retain the inferred requirements, so downstream module checks can reject invalid keys before building Go. Private structural key wrappers are reduced to their required type parameters or associated-type projections. These predicates have no user-written bound syntax and do not introduce let-generalization. Source-only dependency analysis checks dependency bodies and retains their inferred requirements before checking callers, including unsaved source overrides. Struct and enum interfaces also retain the key requirements of maps stored in fields or variant payloads, including private storage and aliases. These requirements propagate between nominal declarations with a recoverable 256-pass limit and are substituted at concrete type applications. Recursive type applications are checked through their declaration requirements without repeatedly expanding their layouts. Default trait methods retain their inferred map-key requirements separately from the method declaration bounds. An implementation that uses a default body must satisfy those requirements after substituting its receiver and associated types. An explicit method override does not inherit requirements from the unused default body. GoLibrary map exports remain unsupported. Raw maps add no synchronization or automatic deep copy.

### Declaring Go bindings

Write explicit `go_ffi`, `go_type` and `go_method` declarations, with Go forwarding functions when an API needs adaptation. See [Go interoperability](ffi/bind-go.md) for the declaration model and the [file reading example](../goml/testdata/ffi/file-read/README.md) for explicit pointer and method bindings.

### C-compatible strings

`std::c` provides `CString`, checked C string literals and boundary errors using ordinary GoML types and calls. See [C-compatible strings](stdlib.md#c-compatible-strings) for its API and compile-time validation. These APIs add no pointer arithmetic, native layouts or `extern "C"` syntax.

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

A raw echo returning `ffi::String` can preserve bytes without decoding. A rune adapter can return `(ffi::Rune, bool)` in the same way. Fixed arrays and flat tuple returns preserve these element policies. `Result` itself is not a Go export ABI type; selecting and documenting the failure shape belongs to the adapter author.

Interface and Core artifacts preserve the selected Go export name, stable package/function identity, resolved signature, and relative source origin. Export names and signatures participate in the interface semantic hash; source positions do not. The internal library linker selects only the chosen package’s declared exports as function roots before dead-code elimination and monomorphization; reachable private and cross-package helpers are retained without requiring a `main`. The internal Go library backend emits callable wrappers for scalar and fixed-array parameters and results, maps unit to no Go result, and maps a flat tuple to multiple Go results. Library wrappers use package-level OnceCell initialization and preserve panics. Generated internal package declarations are private Go identifiers; only explicitly selected export names become public API. `pub` alone does not request a Go export.

Use `goml export-go` to generate a library from one package in the current GoML module:

```text
goml export-go [PACKAGE_DIR] --import-path example.com/host/gen/calclib --out ./gen/calclib
```

The default source is the GoML module root. Explicit source and output paths are relative to the current working directory. The source must be a package in the current module and must not declare `package main`. Only its `go_export` declarations become Go API; dependency exports are retained only when reachable. The output directory's basename is the Go package name and must be a non-main ASCII Go identifier.

The GoML module root must already contain `go.mod`. The output must remain inside that Go module, outside nested Go modules, and the supplied import path must equal the Go module path plus the canonical relative output path. Parent traversal is rejected and symlinks are resolved before checking boundaries. The command never creates or repairs `go.mod` or `go.sum` or downloads missing dependencies.

The output contains `goml_generated.go` and `goml_exports.json`. The manifest records public signatures, generation version, Go build prerequisites and ownership digests. Build prerequisites include the Go toolchain, target, cgo setting, flags and build tags. Handwritten Go files are preserved and checked together with the candidate generated source. New output is checked using a Go overlay before publication; malformed Go, duplicate symbols and import cycles fail without publishing the candidate. Modified or unowned generated files are not overwritten. Concurrent generation for the same output is rejected. Publication rolls back reported file replacement errors; two-file publication is not crash-atomic across power loss or forced process termination.

Go compilation checks the candidate generated package and every emitted foreign call before publication. Foreign declarations remain trusted during GoML analysis; unused declarations need not appear in the candidate. `--compiler`, `--target-dir`, `--jobs` and `--dry-run` are supported. The generated pure-Go library builds with the Go toolchain alone. A complete bidirectional example, including a Go executable and Go tests, is available in [the interoperability fixture](../tools/release/testdata/go-export).

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

See [testing assertions](stdlib.md#testing-assertions) for the assertion and payload-extraction helpers in `std::testing`.

### White-box and black-box tests

The white-box test file and the source code under test are located in the same package directory. The file name must end with `_test.goml` and declare the same package name:

```text
math/
├── math.goml
└── math_test.goml
```

`math_test.goml`:

```goml
package math;

use std::testing;

#[test]
fn private_helper_works() -> () {
    testing::assert_eq(private_helper(), 42)
}
```

The test build will merge the production source code and all `*_test.goml` in the same directory into a single package, so white-box testing can access the private top-level items of the package. These files will not participate in compilation when executing normal `goml check`, `goml build` or `goml run` for production targets.

Black-box tests are located in the `tests/` directory of the package under test. This directory as a whole constitutes a package named `tests`. The package under test should be imported explicitly and only its public API can be accessed; nested test suites cannot be created under `tests/`:

```text
math/
├── math.goml
└── tests/
    ├── api_test.goml
    └── smoke_test.goml
```

`math/tests/api_test.goml`:

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

In the compiler repository, standard-library behavior tests live in `lib/std/<package>/tests/` and use native GoML assertions or fixed known-answer data. Run `just test-stdlib` from the repository root. For this source module, `goml test` first builds a private compiler world from the sibling `builtin`, `prelude`, and `std` sources under the selected target directory, then compiles the black-box tests against it. This tests the current library source. Standard-library white-box tests are currently unsupported and produce a diagnostic. Installed toolchains omit these test sources and data.

### Check and run tests

`goml check`, `goml build` and `goml test` discover the module from the current directory upwards. They process complete modules unless an example selector is supplied, and do not accept package or file targets. `goml check` only checks the production source code by default; using `--tests` will also check all white-box and black-box tests after the production package check is successful:

```sh
goml check
goml check --tests
goml build
goml fmt
goml fmt --check
```

Test source code cannot be used to fix a production package that itself fails inspection.

`goml fmt` formats the current module's production files, internal tests, black-box test packages, and examples with their tests using the same package graphs as builds and tests. It can be run from any nested package directory. Package discovery excludes `testdata`, build output, hidden directories, nested modules, and external dependencies. `goml fmt --check` only checks formatting and exits unsuccessfully when any source file would change.

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

Each test is executed in a separate runner process, and failure to exit and timeout do not affect other tests; `--jobs` controls the number of test processes running at the same time. Building test runners with `goml test` requires an available Go toolchain.

`goml test --no-run` builds all selected test groups without executing tests. `goml test --no-build` executes existing test runners and manifests from the target directory without compiling or linking them. It can be combined with filters, `--list`, and sharding. Prebuilt test artifacts must come from the same source revision and use the execution machine’s OS and architecture; `--no-build` deliberately skips freshness checks, so missing artifacts fail instead of being rebuilt. To transfer artifacts, preserve each test group’s `runner` and `tests.json` paths and any executable-relative runtime libraries. `--no-build` conflicts with `--no-run` and `--dry-run`; `--no-run` conflicts with `--list`.

`GOML_TEST_SHARD=INDEX/COUNT` splits matching tests across independent machines. Both numbers are positive, one-based decimal integers with at most nine digits, and `INDEX` must not exceed `COUNT`. After applying the name and ignored-test filters, the driver sorts tests by display name and assigns successive tests to shards in round-robin order. For example, run `GOML_TEST_SHARD=1/4 goml test --jobs 4` through `GOML_TEST_SHARD=4/4 goml test --jobs 4` on four machines. `--list` uses the same shard selection. An unset or empty value runs the full selection; an invalid value produces a diagnostic. The driver clears this variable in test processes so nested `goml test` commands run their complete selection.

`goml test --seed N` supplies one reproducible positive seed to every test process through `testing::seed()`. Failure summaries print the replay seed, and every JSON result event includes it. Captured stdout and stderr remain isolated per test process and are emitted only for failures unless `--nocapture` is selected.

### LSP and editor

Project configuration and module dependency resolution belong to the `goml` driver. The LSP requests a versioned project description from `goml __project-info --source <PATH>` and passes it with open document contents to the compiler query engine. The description supplies canonical module identities, source directories, separate production/development dependency edges, and configuration watch paths. This operation reads local files and cached dependencies without compiling, fetching, or writing build artifacts. Dependency failures remain diagnostics; available root metadata is retained for local package completion.

The project-description client validates protocol version 1 and caches successful descriptions while their watched configuration contents and canonical paths remain unchanged. Ordinary source edits reuse project resolution and still participate in the compiler's source-analysis fingerprints. Workspace creation, dependency manifest changes, replacements, and registry index changes trigger resolution again. Missing dependencies and invalid descriptions are retried after edits rather than cached as successful projects. The compiler retains source import resolution, visibility checks, and type checking.

`GOML_PROJECT_DRIVER` selects an explicit driver. Otherwise clients use a sibling `goml`, `GOML_HOME/bin/goml`, or `PATH`; repository tests may supply `GOML_TEST_GOML`. The driver and language server must support the same project-description protocol. The extension bundles both tools and offers `goml.driverPath` for custom installations. Existing compiler build commands and driver protocol remain unchanged, so pinned stage0 can continue building the tools.

The LSP analyzes production files, internal tests, and black-box tests in their corresponding package contexts, so diagnostics, completion, hover, and go-to-definition follow the same visibility rules as compilation. Test code lenses appear on ordinary tests and parameterized test cases. The VS Code extension saves the file and invokes module-level `goml test` with the selected test display name and kind.

Type inlay hints use short names without package qualifiers and a 25-character type budget, excluding the leading `: `. Nested types and long names are abbreviated with `…`; the hint tooltip retains the complete inferred type and its package paths.

Package analysis caches parsed/lowered sources and successful checked packages within a query context. Source contents, package membership, canonical module identity, and transitive dependency fingerprints determine reuse. Editing a package rechecks it and its dependents while independent packages remain reusable; a function-body edit in a dependency conservatively invalidates its dependents too. Open document overrides participate in the same fingerprints. External Go types remain opaque during cached and fresh analysis; native validation belongs to the final Go build. Watched source and manifest changes invalidate cached document analyses; the VS Code extension also watches `goml.work`, Go sources, `go.mod`, and `go.sum`.

The custom `goml/expandedDerive` request returns the formatted AST after built-in and programmable derives have run for the requested document. It uses the same package aliases, explicit derive imports, ambiguity checks, dependency interfaces, CTIR verifier, and resource limits as `goml check`. The VS Code command `GoML: Show Expanded Derive` opens that result beside the source file.

`gomllsp` supports full-document formatting through `textDocument/formatting`. Formatting uses the latest unsaved document text and the fixed rules described in [formatting.md](formatting.md). Invalid documents are left unchanged.

CLI and LSP diagnostics share a model with optional stable codes, notes, help, and suggested edits. The CLI displays codes as `error[typer/try_error_mismatch]` or `warning[typer/unused_result]`; the LSP exposes the same identifier in `Diagnostic.code` and secondary locations in `relatedInformation`. Notes and help appear in the displayed message and separately in diagnostic `data`, alongside fixes and the document version. Fixes carry a title, one or more ranged edits, and `machine-applicable` or `requires-review` applicability. Code actions reject fixes from an older document version. Codes are currently assigned to the unused-Result warning and the three `?` diagnostics above; other diagnostics may omit them.

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

Prefer `Some`, `None`, `Ok`, and `Err` in construction and patterns when the expected type identifies the enum. Qualified names such as `Option::Some` and `Result::Ok` remain available when type context is insufficient or another declaration shadows the variant.

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
let first = values[0];
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

See [map algorithms](stdlib.md#map-algorithms) for cloning, copying, comparing, and iterating maps with `std::collections`.

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

See [dynamic channel selection](stdlib.md#dynamic-channel-selection) for runtime-sized selection with `std::channel`.

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

See [iterator adapters](stdlib.md#iterator-adapters) for construction, adapters, and consumers in `std::iter`.

Iterators are single pass. Fixed arrays use native indexed `for` lowering. `Vec[T]`, `Slice[T]` and `MutSlice[T]` can be directly used in `for`, and a value implementing `Iterator` is also directly iterable through the identity `IntoIterator`.

## Comparison of common writing errors

For library APIs and common tasks, see the [standard library reference](stdlib.md#common-tasks).

| Avoid | GoML form |
| --- | --- |
| `////` or a blank source line inside an API documentation block | Use `///`; use a bare `///` line to separate paragraphs |
| `Vec<isize>` | `Vec[isize]` |
| `fn id<T>(x: T) -> T` | `fn id[T](x: T) -> T` |
| `fn count[N](values: [i32; N])` | Declare the value parameter as `fn count[const N: usize](values: [i32; N]) -> usize` |
| `Buffer[i32, N + 1]` | `Buffer[i32, {N + 1}]` |
| Using `serde(rename)` to change `.to_string()` | Use `#[to_string("text")]` on the variant with `#[derive(ToString)]` |
| `Type::LIMIT()` for an associated constant | `Type::LIMIT` |
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
| `Message::Value::[i32](42)` | Enum arguments belong to the owner: `Message::[i32]::Value(42)` |
| `identity::[](1)` | Omit the empty turbofish and write `identity(1)` |
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

## Verify generated code

After running `just make`, verify a standalone source file from the repository root:

```sh
stage2/bin/gomlc run-single path/to/main.goml
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

`goml check`, `goml build`, `goml test`, and `goml fmt` operate on complete modules and do not accept package or file targets. In a workspace, use `--workspace` for all members or `-p <module-path>` for one member. `goml run [TARGET]` accepts an optional entry package file or directory when a module has multiple executable packages.

When you need to inspect a compilation phase, add `--dump-ast`, `--dump-expanded-ast`, `--dump-hir`, `--dump-tast`, `--dump-ctir`, `--dump-core`, `--dump-mono`, `--dump-lift`, `--dump-anf`, or `--dump-go` to `gomlc run-single`. `--dump-ast` shows source lowering before derive expansion, while `--dump-expanded-ast` includes every generated implementation.

The code agent should at least run the corresponding `goml check` or `gomlc run-single` before submitting the source code; when modifying the test, it should also run `goml check --tests` and the related `goml test`. When type inference fails, give priority to adding local result types, empty container types, closure parameter types, or using UFCS instead of rewriting to unsupported Rust/Go syntax.
