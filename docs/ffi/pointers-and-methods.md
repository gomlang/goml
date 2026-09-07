# B1 pointers and method bindings

This contract implements the B1 requirements of plan.md. B0 identity acceptance is recorded in acceptance-b0.md.

## Representation and public boundary

`use std::ffi;` exposes `ffi::Ptr[T]` as a nullable raw Go pointer. Its Go representation is `*T`, where T uses the existing Go identity rules. It is distinct from GoML `Ref[T]`, integers, handles and a value of T. Copies preserve the same pointer and normal Go GC reachability. There is no object registry, finalizer-driven Close or automatic allocation when a nil pointer is received.

The public surface provides `ffi::null[T]() -> Ptr[T]` and `ffi::is_nil[T](value: Ptr[T]) -> bool`. Non-null pointers enter through checked Go bindings; no integer-to-pointer conversion, arbitrary layout construction or unchecked Ref-to-Ptr reinterpretation is introduced. The exact builtin-to-standard-library wiring must preserve stage0 compatibility. The new capability will not be used in compiler or driver implementation sources before stage0 advances.

`#[go_method("Method")] extern fn ...` declares the first parameter as the receiver. The Go checker determines method selection from that parameter's actual Go type and method set. Value and pointer receiver behavior must match the generated Go operation. The helper validates an explicit Go method expression, so a value receiver cannot silently become a pointer receiver through address-taking. Promoted methods and interface method sets follow Go rules. A nil pointer receiver is passed through to Go; any behavior of the target method, including failure, is retained.

## Protocol and artifacts

The bridge protocol encodes a pointer as `{ "tag": "pointer", "element": <type> }`. Pointees retain canonical named identity and nesting. Pointer nodes count toward existing type depth/node limits. A pointer must have an element; it cannot carry array lengths or channel directions. The legacy direct-value `any` marker cannot be nested inside a pointer.

The Go helper constructs an AST StarExpr for the type and lets the Go compiler validate assignability. Interface binding codecs and persisted instance argument codecs preserve pointer nodes. Go pointer aliases now lower through a distinct GoPointer IR representation to *T. The public Ptr alias and direct nil operations are implemented through builtin language-item normalization and Go AST nil/comparison expressions. Method bindings now lower through explicit GoMethod callable variants and a Go AST MethodExpression.

## Acceptance work

The os.Open and *os.File Read/Close fixture now covers pointer aliases sharing one object, explicit Close and nil receiver semantics. It uses the existing explicit MutSlice view; no implicit Vec conversion was added. Any error and byte-buffer representations needed by those fixtures must follow the subsequent B2/B3 contracts without introducing implicit conversions.

The Go metadata helper now accepts method-mode requests, checks exported receiver method sets and records their actual signatures. Tests cover pointer/value receivers, pointer aliases, interfaces, promoted methods, instantiated generic receivers, duplicate witnesses, invalid signatures and nil receiver execution. GoML bridge requests and interface metadata now preserve the call mode and receiver. Source declarations, callable/Core transport and backend method emission are now connected. Cross-package execution covers method function values, closure capture, shared mutation and nil receivers.
