# C bindings

`goml bind-c bindings.json` generates GoML declarations and a native bridge from an explicit C allowlist. The Go backend uses cgo internally; binding users write GoML and C configuration without handwritten Go adapters. `gomlc bind-c` exposes the same generator. The packaged `goml-c-bind` helper inspects headers using Clang, formats the sources, and compiles and links a probe without executing it or package initializers.

Generation and project checks require the native host target, `CGO_ENABLED=1`, a C compiler, and Clang 15 or newer. `GOML_CLANG` selects Clang; otherwise the tool searches `clang` and versioned executables. This first version supports Linux amd64 as the toolchain's development target. Cross compilation is rejected. Inspection and native compilation default to C11; a configured `-std=` flag overrides it in both. The C compiler and Clang must agree on ABI and preprocessing flags.

## Module configuration

An existing module-root `goml.toml` and `go.mod` are required. For example:

```toml
[module]
path = "sqlite_example"

[native]
go-module = "example.com/sqlite-example"
cgo = "required"
c-bindings = "bindings.json"
```

```text
module example.com/sqlite-example

go 1.26.0
```

`native.c-bindings` is one module-relative configuration path. Project check, build, run and test verify bindings in the module and its dependencies before foreign compilation, including configurations containing only constants. Normal builds do not rewrite generated sources. An outdated declaration or a modified generated file produces a diagnostic directing the author to regenerate or restore it. Clang is therefore also required when consuming these configured bindings.

The fingerprint covers configuration, Clang identity, Go target, and preprocessed headers, including transitive headers and their implementation bodies. It is passed into native compilation so changed C header contents invalidate Go's cgo cache even when the generated wrapper stays identical. Ordinary changes to a linked shared library remain the system linker's responsibility; replacing a static archive without changing inputs may require cleaning the Go build cache.

## Binding configuration

```json
{
  "version": 1,
  "package": "sqlite",
  "output": "sqlite/generated.gom",
  "go_package": "native",
  "go_output": "native/generated.go",
  "headers": ["sqlite3.h"],
  "ldflags": ["-lsqlite3"],
  "types": [{"name": "Database", "c_type": "sqlite3 *"}],
  "constants": [{"name": "OK", "symbol": "SQLITE_OK"}],
  "functions": [
    {
      "name": "open",
      "symbol": "sqlite3_open",
      "parameters": [
        {"name": "filename", "kind": "cstring"},
        {"name": "database", "kind": "out"}
      ]
    },
    {"name": "close", "symbol": "sqlite3_close"}
  ]
}
```

Output paths resolve relative to the configuration and stay in the GoML module. Native output belongs in a subdirectory of its Go module. Parent traversal, output symlinks and nested modules are rejected. The Go import identity is derived from `go.mod` and the native output directory.

`include_dirs`, `cflags`, `ldflags`, and `pkg_config` are optional lists. Relative include and library directories resolve from the configuration. Supported explicit flags are `-I`, `-D`, `-U`, `-std=`, `-L`, `-l`, and `-pthread` for linking, each as one argument. Whitespace and quoting in explicit flags are rejected; use `include_dirs` for include paths containing spaces. Quoted or escaped environment `CGO_CFLAGS`/`CGO_CPPFLAGS` and pkg-config output are currently rejected during inspection. Host cgo linker configuration also applies to compilation. A header in the configuration directory can be listed directly.

Unknown fields, duplicate JSON fields, unsupported signatures, ambiguous type mappings and missing symbols are errors. Configuration is limited to 1 MiB, 4096 bindings and 128 parameter mappings per function.

## Types and calls

Opaque types name C pointer types, including typedefs such as `LLVMContextRef`. Each becomes a distinct GoML struct with private native storage and `null()`, `is_null()` and `same_as(other)` methods. No pointer-to-integer conversion or dereference is exposed. A copied GoML handle refers to the same C resource. Calling a release function does not invalidate copies, and GC does not release C resources. Callers must obey the C API's nullability, parent lifetime, thread, aliasing and ownership requirements; these low-level bindings do not make arbitrary C APIs memory safe.

Scalar C widths and signedness come from Clang. Supported values are `_Bool`, 8/16/32/64-bit integers, `float`, `double`, and C enums of at most four bytes. Enums use `i64` and checked conversion into the actual C type; this checks representation, not whether a value names an enumerator. `long` and `size_t` follow the host ABI. Generated C assertions verify scalar ABI, function signatures and integer constants against the compiler actually used by cgo.

Every generated function returns `Result[T, c::Error]`. This error reports adapter failures such as a copy limit or a buffer length that does not fit the C parameter. C status codes remain ordinary values: `sqlite::open` returns `Result[(i32, Database), c::Error]`, preserving the database even when the C status is nonzero. There is no implicit `errno` or status-to-error convention. An adapter failure after a C call does not roll back C side effects; other returned resources may require an API-specific wrapper and cleanup policy.

Omitting `parameters` infers ordinary scalar/opaque parameters named `arg0`, `arg1`, etc. Otherwise list every physical C parameter in order, with unique GoML names:

| Kind | C parameter | Generated API |
| --- | --- | --- |
| `value` or omitted | Scalar or configured opaque pointer | Ordinary argument; `type` disambiguates opaque mappings |
| `cstring` | `char *` or `const char *` | `c::CString`, copied into C allocation for this call |
| `bytes` | Byte or void pointer | `Bytes`, copied into C allocation for this call |
| `inout_bytes` | Writable byte or void pointer | Input `Bytes` plus an independent copied `Bytes` result |
| `length` | Integer | Hidden parameter filled with the length of the buffer named by `of`, checked for overflow |
| `out` | Pointer to scalar or opaque pointer | Hidden, zero-initialized storage; becomes a result; optional `type` disambiguates opaque types |
| `out_string` | `char **` or `const char **` | Hidden output pointer, copied into `Option[c::CString]` |

The result lists a non-void C return first, then output parameters in declaration order. An empty result is `()`, one result is unwrapped, and multiple results form a tuple.

Input strings and buffers are limited to 64 MiB. C cannot retain these pointers after the call or write beyond the allocation. Empty buffers pass NULL and length zero. Inputs are independently copied even when the same GoML buffer is passed twice; use a handwritten C adapter for APIs requiring aliasing or retained allocations. In/out buffers preserve the supplied capacity, not a C-written output length.

Set `"return": {"kind": "cstring"}` for a C `char *` return. Null becomes `None`; an empty C string becomes `Some` with zero bytes. Copying preserves non-UTF-8 bytes. An optional `release` names a `void(char *)` or `void(void *)` function called after copying, including copy-limit failure. This also applies to `out_string`. Without it the pointer is borrowed and only copied. A non-null result must permit reading up to `max_bytes + 1` bytes if no earlier NUL is present; a limit cannot make an invalid pointer safe. `max_bytes` defaults to 1 MiB and may be at most 1 GiB. An explicit release function must tolerate the output values its C producer can return, including NULL.

## Compile-time integration

Integer object macros and enumerators in `constants` become typed GoML `pub const` declarations, with values evaluated by Clang. They can participate in GoML `comptime` arithmetic, masks and validation without C calls at runtime. Floating constants, string macros and function-like macros are not supported in this version.

`std::c` provides `CString::new(string)`, `from_bytes(Bytes)` and `from_raw(ffi::String)`, each rejecting interior NUL with `Result[CString, c::Error]`. `bytes()` preserves bytes, `text()` validates UTF-8, and `as_raw()` is the bridge representation. `Error` has `InteriorNul` and `Native(string)` variants and implements `Debug`, `PartialEq`, `Eq` and `ToString`.

`c::literal(string) -> string` is a `#[comptime]` function that rejects interior NUL through `compile_error`. Use a constant initializer or `comptime` block to require compile-time validation:

```goml
use std::c;

const DATABASE: string = c::literal(":memory:");

fn database_name() -> Result[c::CString, c::Error] {
    c::CString::new(DATABASE)
}
```

The ordinary `CString` constructor still enforces its invariant at runtime. GoML CTFE does not perform host I/O, invoke Clang, allocate C memory, or determine C layouts. A normal runtime call to `c::literal` is not forced into CTFE and traps if invalid. `c::check_error(ffi::Error)` is the checked adapter helper used by generated functions.

## Generation and recovery

`--dry-run` validates configuration and ownership paths and prints outputs without querying headers or writing. `--check` compares current generated sources and prints a 64-character C input fingerprint. It does not publish sources or execute C code. The driver also accepts `--compiler`.

The generator owns two source files and `<CONFIG>.goml-c-bind.json`, which records content hashes. Commit all three together with the configuration. Put handwritten wrappers in separate files. Files without matching ownership, edited outputs and missing outputs prevent replacement. Restore the previous generated files before regeneration; to move outputs, remove the old owned files and manifest deliberately first.

A module-local lock excludes concurrent generators. Sources are staged, native code is compiled and linked, and output renames publish the validated result. Publication errors retain the staging directory with `previous.json` mapping previous outputs to backup files. No multi-file filesystem transaction is claimed; a process or machine crash during publication can leave a partial set that ownership checks reject. A stale `.goml-bind-c-lock` must be inspected and removed once no generator is active. The generator does not download C libraries or execute their initialization code during validation.

Callbacks, variadic calls, struct/union values, field access, exported C functions, cross-target ABI descriptions and a cgo-free calling backend remain outside this version. An explicit C wrapper can reduce many such APIs to supported fixed signatures. See the [self-contained C example](../../examples/ffi-bind-c/README.md), [SQLite example](../../examples/ffi-c-sqlite/README.md), and [LLVM example](../../examples/ffi-c-llvm/README.md).
