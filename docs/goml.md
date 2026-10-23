<a id="goml-language-guide"></a>

# GoML language guide

[English](goml.md) | [简体中文](goml.zh-CN.md)

This guide describes the GoML syntax, type rules, tools, and built-in prelude implemented in this repository. Standard-library APIs are documented in the [standard library reference](stdlib.md). Use it when writing, modifying, and reviewing `.goml` source. Similarities to Rust, Go, or OCaml do not imply that their syntax or APIs are supported.

GoML is a statically typed language with garbage collection. Its syntax is close to Rust, while its semantics are closer to ML. The compiler monomorphizes generics and lambda-lifts GoML closures before emitting Go. Explicit Go FFI can preserve native generic types and function values. GoML has no ownership, borrowing, lifetimes, or manual memory management.

Generated Go code and Go FFI use Go 1.26 as their compatibility baseline. Building the toolchain and compiling generated programs require Go 1.26 or newer. Native FFI signatures are checked by that final Go compilation. The independent [gomlgo frontend and interpreter](https://github.com/gomlang/gomlgo) targets Go 1.26 language and standard-library behavior; its execution and differential tests require Go 1.26.x.

Quick navigation: [packages](language/packages.md#modules-packages-and-imports), [types](language/types.md#type), [control flow](language/basics.md#control-flow), [patterns](language/types.md#patterns), [traits](language/types.md#trait-and-impl), [compile-time evaluation](language/metaprogramming.md#compile-time-evaluation), [Go FFI](ffi/bind-go.md#go-ffi), [tests](tooling.md#tests), [API documentation](language/basics.md#comments-and-api-documentation), [prelude](language/prelude.md#built-in-prelude), [standard library](stdlib.md), and [verification](tooling.md#verify-generated-code).

Build and installation instructions are in the [development guide](development.md) and [release guide](releasing.md#installation). See [formatting](formatting.md) for canonical source layout and [compile-time evaluation architecture](comptime.md) for CTIR internals.

<a id="reading-guide"></a>

## Reading guide

| Topic | Contents |
| --- | --- |
| [Language basics](language/basics.md) | Lexical rules, definitions, functions, bindings, expressions, control flow, and closures |
| [Types, patterns, and traits](language/types.md) | Types, structs, enums, generic parameters, pattern matching, traits, and dynamic dispatch |
| [Modules, packages, and visibility](language/packages.md) | Source-level package identities, imports, and visibility |
| [Compile-time evaluation and derive](language/metaprogramming.md) | Language rules for compile-time evaluation, attributes, and programmable derive |
| [Prelude reference](language/prelude.md) | Automatically scoped types, traits, and operations |
| [Grammar and common mistakes](language/grammar.md) | A syntax comparison and the informal grammar |
| [Project tools and workflows](tooling.md) | Module configuration, dependencies, examples, build output, testing, editor support, and generated-code inspection |
| [Go FFI](ffi/bind-go.md) | Go bindings, raw boundary types, and Go library export |

<a id="essential-rules"></a>

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
12. Use `module::path` for a package below the current module root. Do not generate `mod`, `crate::`, `self::`, `super::`, root paths `::x`, Rust references, or Go `var` / `:=`. A user `extern fn` is valid only with the [typed Go FFI attributes](ffi/bind-go.md#go-ffi).
13. The test function uses `#[test]`, which must have no parameters, no type parameters and return `()`; the white-box test is placed in `*_test.goml` of the same package, and the black-box test is placed in the `tests/` directory of the package under test.

<a id="minimal-program"></a>

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

<a id="original-section-links"></a>

## Original section links

The table preserves the original section anchors and links to their new locations.

| Original section | Current document |
| --- | --- |
| <a id="lexical-rules"></a> Lexical rules | [Language basics](language/basics.md#lexical-rules) |
| <a id="comments-and-api-documentation"></a> Comments and API documentation | [Language basics](language/basics.md#comments-and-api-documentation) |
| <a id="identifiers-and-keywords"></a> Identifiers and keywords | [Language basics](language/basics.md#identifiers-and-keywords) |
| <a id="literals"></a> Literals | [Language basics](language/basics.md#literals) |
| <a id="modules-packages-and-imports"></a> Modules, packages and imports | [Modules, packages, and visibility](language/packages.md#modules-packages-and-imports) |
| <a id="gomltoml"></a> `goml.toml` | [Project tools and workflows](tooling.md#gomltoml) |
| <a id="package-registry"></a> Package registry | [Project tools and workflows](tooling.md#package-registry) |
| <a id="examples"></a> Examples | [Project tools and workflows](tooling.md#examples) |
| <a id="native-dependencies"></a> Native dependencies | [Project tools and workflows](tooling.md#native-dependencies) |
| <a id="build-output-layout"></a> Build output layout | [Project tools and workflows](tooling.md#build-output-layout) |
| <a id="directory-packages"></a> Directory packages | [Modules, packages, and visibility](language/packages.md#directory-packages) |
| <a id="use"></a> `use` | [Modules, packages, and visibility](language/packages.md#use) |
| <a id="visibility"></a> Visibility | [Modules, packages, and visibility](language/packages.md#visibility) |
| <a id="type"></a> Type | [Types, patterns, and traits](language/types.md#type) |
| <a id="basic-types-and-composite-types"></a> Basic types and composite types | [Types, patterns, and traits](language/types.md#basic-types-and-composite-types) |
| <a id="type-aliases"></a> Type aliases | [Types, patterns, and traits](language/types.md#type-aliases) |
| <a id="type-syntax-not-currently-available"></a> Type syntax not currently available | [Types, patterns, and traits](language/types.md#type-syntax-not-currently-available) |
| <a id="top-level-definitions"></a> Top-level definitions | [Language basics](language/basics.md#top-level-definitions) |
| <a id="constants"></a> Constants | [Language basics](language/basics.md#constants) |
| <a id="statics-and-one-time-initialization"></a> Statics and one-time initialization | [Language basics](language/basics.md#statics-and-one-time-initialization) |
| <a id="compile-time-evaluation"></a> Compile-time evaluation | [Compile-time evaluation and derive](language/metaprogramming.md#compile-time-evaluation) |
| <a id="structs"></a> Structs | [Types, patterns, and traits](language/types.md#structs) |
| <a id="enums"></a> Enums | [Types, patterns, and traits](language/types.md#enums) |
| <a id="functions-and-generics"></a> Functions and generics | [Types, patterns, and traits](language/types.md#functions-and-generics) |
| <a id="top-level-functions"></a> Top-level functions | [Language basics](language/basics.md#top-level-functions) |
| <a id="opaque-return-types"></a> Opaque return types | [Types, patterns, and traits](language/types.md#opaque-return-types) |
| <a id="generic-functions"></a> Generic functions | [Types, patterns, and traits](language/types.md#generic-functions) |
| <a id="blocks-bindings-and-assignments"></a> Blocks, bindings and assignments | [Language basics](language/basics.md#blocks-bindings-and-assignments) |
| <a id="const-generic-parameters"></a> Const generic parameters | [Types, patterns, and traits](language/types.md#const-generic-parameters) |
| <a id="block-values-and-semicolons"></a> Block values and semicolons | [Language basics](language/basics.md#block-values-and-semicolons) |
| <a id="let-type-annotations-and-shadowing"></a> `let`, type annotations and shadowing | [Language basics](language/basics.md#let-type-annotations-and-shadowing) |
| <a id="mut-and-assignment"></a> `mut` and assignment | [Language basics](language/basics.md#mut-and-assignment) |
| <a id="expressions-and-operators"></a> Expressions and operators | [Language basics](language/basics.md#expressions-and-operators) |
| <a id="basic-expressions"></a> Basic expressions | [Language basics](language/basics.md#basic-expressions) |
| <a id="operator-precedence"></a> Operator precedence | [Language basics](language/basics.md#operator-precedence) |
| <a id="operator-type-rules"></a> Operator type rules | [Language basics](language/basics.md#operator-type-rules) |
| <a id="range-expressions"></a> Range expressions | [Language basics](language/basics.md#range-expressions) |
| <a id="control-flow"></a> Control flow | [Language basics](language/basics.md#control-flow) |
| <a id="if"></a> `if` | [Language basics](language/basics.md#if) |
| <a id="match"></a> `match` | [Language basics](language/basics.md#match) |
| <a id="select"></a> `select` | [Language basics](language/basics.md#select) |
| <a id="while"></a> `while` | [Language basics](language/basics.md#while) |
| <a id="loop"></a> `loop` | [Language basics](language/basics.md#loop) |
| <a id="for"></a> `for` | [Language basics](language/basics.md#for) |
| <a id="break-continue-and-return"></a> `break`, `continue` and `return` | [Language basics](language/basics.md#break-continue-and-return) |
| <a id="try-operator"></a> `?` | [Language basics](language/basics.md#try-operator) |
| <a id="go"></a> `go` | [Language basics](language/basics.md#go) |
| <a id="patterns"></a> Patterns | [Types, patterns, and traits](language/types.md#patterns) |
| <a id="refutability-and-binding"></a> Refutability and binding | [Types, patterns, and traits](language/types.md#refutability-and-binding) |
| <a id="groups-tuples-structures-and-enumerations"></a> Groups, tuples, structures and enumerations | [Types, patterns, and traits](language/types.md#groups-tuples-structures-and-enumerations) |
| <a id="fixed-arrays-vec-slice-and-mutslice"></a> Fixed arrays, Vec, Slice and MutSlice | [Types, patterns, and traits](language/types.md#fixed-arrays-vec-slice-and-mutslice) |
| <a id="alias-and-or-pattern"></a> Alias and or-pattern | [Types, patterns, and traits](language/types.md#alias-and-or-pattern) |
| <a id="range-patterns"></a> Range patterns | [Types, patterns, and traits](language/types.md#range-patterns) |
| <a id="exhaustiveness-and-unreachable-branches"></a> Exhaustiveness and unreachable branches | [Types, patterns, and traits](language/types.md#exhaustiveness-and-unreachable-branches) |
| <a id="closures-and-function-values"></a> Closures and function values | [Language basics](language/basics.md#closures-and-function-values) |
| <a id="closure-syntax"></a> Closure syntax | [Language basics](language/basics.md#closure-syntax) |
| <a id="no-let-polymorphism"></a> No let-polymorphism | [Language basics](language/basics.md#no-let-polymorphism) |
| <a id="trait-and-impl"></a> Trait and impl | [Types, patterns, and traits](language/types.md#trait-and-impl) |
| <a id="define-traits"></a> Define traits | [Types, patterns, and traits](language/types.md#define-traits) |
| <a id="trait-impl"></a> Trait impl | [Types, patterns, and traits](language/types.md#trait-impl) |
| <a id="inherent-impl"></a> Inherent impl | [Types, patterns, and traits](language/types.md#inherent-impl) |
| <a id="method-parsing-and-ufcs"></a> Method parsing and UFCS | [Types, patterns, and traits](language/types.md#method-parsing-and-ufcs) |
| <a id="associated-constants"></a> Associated constants | [Types, patterns, and traits](language/types.md#associated-constants) |
| <a id="dyn-trait"></a> `dyn Trait` | [Types, patterns, and traits](language/types.md#dyn-trait) |
| <a id="attributes-and-derive"></a> Attributes and derive | [Compile-time evaluation and derive](language/metaprogramming.md#attributes-and-derive) |
| <a id="programmable-derive"></a> Programmable derive | [Compile-time evaluation and derive](language/metaprogramming.md#programmable-derive) |
| <a id="go-ffi"></a> Go FFI | [Go FFI](ffi/bind-go.md#go-ffi) |
| <a id="go-interface-wrappers"></a> Go interface wrappers | [Go FFI](ffi/bind-go.md#go-interface-wrappers) |
| <a id="raw-go-function-values"></a> Raw Go function values | [Go FFI](ffi/bind-go.md#raw-go-function-values) |
| <a id="raw-go-strings"></a> Raw Go strings | [Go FFI](ffi/bind-go.md#raw-go-strings) |
| <a id="raw-go-runes"></a> Raw Go runes | [Go FFI](ffi/bind-go.md#raw-go-runes) |
| <a id="raw-go-slices"></a> Raw Go slices | [Go FFI](ffi/bind-go.md#raw-go-slices) |
| <a id="raw-go-maps"></a> Raw Go maps | [Go FFI](ffi/bind-go.md#raw-go-maps) |
| <a id="declaring-go-bindings"></a> Declaring Go bindings | [Go FFI](ffi/bind-go.md#declaring-go-bindings) |
| <a id="c-compatible-strings"></a> C-compatible strings | [Go FFI](ffi/bind-go.md#c-compatible-strings) |
| <a id="exporting-a-go-library"></a> Exporting a Go library | [Go FFI](ffi/bind-go.md#exporting-a-go-library) |
| <a id="tests"></a> Tests | [Project tools and workflows](tooling.md#tests) |
| <a id="test-functions-and-attributes"></a> Test functions and attributes | [Project tools and workflows](tooling.md#test-functions-and-attributes) |
| <a id="white-box-and-black-box-tests"></a> White-box and black-box tests | [Project tools and workflows](tooling.md#white-box-and-black-box-tests) |
| <a id="check-and-run-tests"></a> Check and run tests | [Project tools and workflows](tooling.md#check-and-run-tests) |
| <a id="lsp-and-editor"></a> LSP and editor | [Project tools and workflows](tooling.md#lsp-and-editor) |
| <a id="built-in-prelude"></a> Built-in prelude | [Prelude reference](language/prelude.md#built-in-prelude) |
| <a id="option-and-result"></a> `Option` and `Result` | [Prelude reference](language/prelude.md#option-and-result) |
| <a id="default"></a> `Default` | [Prelude reference](language/prelude.md#default) |
| <a id="output-and-string-conversion"></a> Output and string conversion | [Prelude reference](language/prelude.md#output-and-string-conversion) |
| <a id="reft"></a> `Ref[T]` | [Prelude reference](language/prelude.md#reft) |
| <a id="fixed-arrays"></a> Fixed arrays | [Prelude reference](language/prelude.md#fixed-arrays) |
| <a id="vect"></a> `Vec[T]` | [Prelude reference](language/prelude.md#vect) |
| <a id="frozenvect"></a> `FrozenVec[T]` | [Prelude reference](language/prelude.md#frozenvect) |
| <a id="slicet-and-mutslicet"></a> `Slice[T]` and `MutSlice[T]` | [Prelude reference](language/prelude.md#slicet-and-mutslicet) |
| <a id="hashmapk-v"></a> `HashMap[K, V]` | [Prelude reference](language/prelude.md#hashmapk-v) |
| <a id="channels"></a> Channels | [Prelude reference](language/prelude.md#channels) |
| <a id="iterator"></a> Iterator | [Prelude reference](language/prelude.md#iterator) |
| <a id="comparison-of-common-writing-errors"></a> Comparison of common writing errors | [Grammar and common mistakes](language/grammar.md#comparison-of-common-writing-errors) |
| <a id="informal-grammar-quick-facts"></a> Informal Grammar Quick Facts | [Grammar and common mistakes](language/grammar.md#informal-grammar-quick-facts) |
| <a id="verify-generated-code"></a> Verify generated code | [Project tools and workflows](tooling.md#verify-generated-code) |
