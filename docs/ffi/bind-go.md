# Go interoperability

[English](bind-go.md) | [简体中文](bind-go.zh-CN.md)

[GoML language guide](../goml.md)

GoML calls Go through explicit foreign declarations. The compiler checks the declared GoML types and trusts the binding author's signature. The final Go build checks native symbol existence, assignability, method sets and generic constraints. `goml check` and editor analysis do not load Go package metadata.

```goml
use std::ffi;

#[go_type("os", "File")]
extern type File;

#[go_ffi("os", "Open")]
extern fn open_file(name: ffi::String) -> (ffi::Ptr[File], ffi::Error);

#[go_method("Close")]
extern fn close_file(file: ffi::Ptr[File]) -> ffi::Error;
```

An external type is opaque. Its import path, Go name and declared type arguments identify it; GoML does not inspect its fields, normalize Go aliases, discover methods or derive interface structure. Bind functions and methods explicitly, or provide a small Go shim when a native API needs conversion. `#[go_interface]` remains available with an explicit raw type, trait signatures and method mappings; final Go compilation checks the generated adapter against the real interface.

Use a module-root `go.mod` for local Go shims and Go dependencies. Ordinary project builds disable Go workspace mode, select module mode from the presence of `go.mod`, and otherwise use the user's Go environment for dependency and toolchain behavior. Go library export and native SIMD builds retain readonly, offline module handling and require dependencies to be prepared beforehand. Unicode validity, ownership, mutability and concurrency guarantees remain explicit adapter contracts; successful Go compilation cannot prove them.

Unused foreign declarations may not appear in generated Go and therefore need not be checked by the final build. Changes to native Go sources are checked when the generated program or library is built; `goml check` continues to use the declared GoML signatures.

See [Go FFI](#go-ffi) for the complete boundary rules, the [file reading example](../../goml/testdata/ffi/file-read/README.md) for explicit pointer and method bindings, and [Go library export](#exporting-a-go-library) for the reverse direction.

<a id="go-ffi"></a>

## Go FFI

The runnable [file reading fixture](../../goml/testdata/ffi/file-read/README.md) combines raw pointers, methods, shared byte buffers, partial reads, error identity, explicit Close and Go↔GoML result transport.

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

<a id="go-interface-wrappers"></a>

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

<a id="raw-go-function-values"></a>

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

The runnable [callback subscription fixture](../../goml/testdata/ffi/callbacks/README.md) provides explicit unregister, gated asynchronous dispatch, concurrent atomic updates, recursive reentry and registration stress coverage. Its callback reference is removed by unregister; already selected invocations may still run, so callers wait for their batches before disposing of captured resources. GC reachability and registration lifetime are separate. The fixture and its generated Go program are tested with the race detector. GoLibrary function-value exports remain unsupported.

Executable, test and library module links embed a versioned JSON index in the generated Go string constant `_goml_source_origins`, relating surviving functions and callback adapters to enclosing GoML declarations. Shared adapters may have several source entries. External-call wrappers also identify the actual extern declarations used, including `go_method`; unused aliases of the same native target are excluded. The index records canonical symbols, the source package, package-relative source files and declaration-name positions, including explicitly defined impl methods; it is not an expression-level stack-trace line table.

The index has `version: 1`, `scope: "enclosing-declaration"` and an `entries` array. Entries contain `generated_symbol`, `source_symbol`, `source_package`, `source_file`, `start`, `end`, `line` and `column`; callback methods use `ReceiverType.Apply` as their generated symbol. Decode the Go string literal before parsing the JSON.

<a id="raw-go-strings"></a>

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

<a id="raw-go-runes"></a>

### Raw Go runes

`ffi::Rune` is a transparent alias of `i32`, matching Go's `rune`. Raw values may be negative, surrogate code points or larger than the Unicode scalar range. `ffi::rune_from_char(char) -> ffi::Rune` preserves a valid character. `ffi::rune_char(ffi::Rune) -> Result[char, ffi::BoundaryError]` checks the scalar range and returns `InvalidRune(original_value)` for invalid input. It never substitutes a replacement character. These adapters use ordinary type alias and function syntax.

<a id="raw-go-slices"></a>

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

<a id="raw-go-maps"></a>

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

<a id="declaring-go-bindings"></a>

### Declaring Go bindings

Write explicit `go_ffi`, `go_type` and `go_method` declarations, with Go forwarding functions when an API needs adaptation. See [Go interoperability](bind-go.md) for the declaration model and the [file reading example](../../goml/testdata/ffi/file-read/README.md) for explicit pointer and method bindings.

<a id="c-compatible-strings"></a>

### C-compatible strings

`std::c` provides `CString`, checked C string literals and boundary errors using ordinary GoML types and calls. See [C-compatible strings](../stdlib.md#c-compatible-strings) for its API and compile-time validation. These APIs add no pointer arithmetic, native layouts or `extern "C"` syntax.

<a id="exporting-a-go-library"></a>

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

Go compilation checks the candidate generated package and every emitted foreign call before publication. Foreign declarations remain trusted during GoML analysis; unused declarations need not appear in the candidate. `--compiler`, `--target-dir`, `--jobs` and `--dry-run` are supported. The generated pure-Go library builds with the Go toolchain alone. A complete bidirectional example, including a Go executable and Go tests, is available in [the interoperability fixture](../../tools/release/testdata/go-export).
