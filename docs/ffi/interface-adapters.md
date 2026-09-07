# Interface adapter implementation contract

Status: D0 is implemented and accepted for module compilation; see acceptance-d0.md for focused and full CI evidence. D1 is implemented and accepted with reader, artifact, concurrency and negative-contract coverage. See acceptance-d1.md for the full CI evidence and supported boundary.

```goml
use std::ffi;

#[go_type("io", "Reader")]
pub extern type GoReader;

#[go_interface(GoReader, ReaderAdapter, read = "Read")]
pub trait Reader {
    fn read(self: Self, buffer: ffi::RawSlice[u8]) -> (isize, ffi::Error);
}
```

The first argument names an explicitly declared raw Go interface type; it may be a qualified imported type path. The second is an explicit unqualified wrapper type name. Named string arguments map GoML trait methods to exported Go method names. Every trait method and every method of the completed Go interface must have exactly one mapping. Duplicate mappings and ambiguous method identities are errors. Wrapper/member name collisions must be diagnosed rather than silently shadowing declarations.

The generated wrapper inherits the trait's visibility and holds one private raw interface value. `ReaderAdapter::from_go(value)` creates the wrapper; `wrapper.into_go()` returns the original raw value. Those method names are reserved by the adapter declaration. A nil interface and an interface containing a typed-nil pointer remain distinct after both conversions. Forwarding uses the original receiver and normal Go method behavior, without adding nil guards or replacing receiver objects. Conversion to a GoML `dyn Reader` remains an explicit ordinary GoML trait conversion of the wrapper.

The initial supported mapping is a concrete Go interface and a non-generic GoML trait with instance methods, no supertraits, no associated types and no additional method predicates. Unsupported contracts produce recoverable diagnostics. A finite Go interface instantiation exposed through a concrete Go alias can retain its original identity; unrestricted GoML-to-Go generic API translation is not promised.

Parameter and result representations must match the completed Go signatures exactly. Unit corresponds to no Go results; multiple Go results retain their tuple shape. Result components cannot be discarded. Text requires `ffi::String`/`ffi::Rune`; slices require `ffi::RawSlice` and explicit copy/share adapters at the caller. Native named types, raw errors, pointers, maps and raw function values retain the existing B/C boundary rules. Ordinary GoML closures, trait objects and runtime containers do not become raw Go interface method parameters implicitly.

Project validation must require current Go interface metadata and verify the complete mapping before publishing the generated trait implementation, including in `--ffi-check off` mode. Persisted external metadata and ordinary generated function/trait artifacts must retain the native interface identity, complete method dependencies and source origin. The wrapper declaration and forwarding helpers need deterministic names, formatter support, diagnostics and cross-package/GoLibrary checks.

D1 consumes the same native method metadata to emit a Go bridge struct with a static interface-satisfaction assertion. Unexported methods that the target package cannot implement must be rejected. D2 reuses the shared metadata and boundary conversion rules for allowlisted source generation through bind-go. Its full acceptance, including packaged-toolchain verification, is recorded in acceptance-d2.md.

## D1 implementation contract

D1 extends the same explicit method mapping with `ReaderAdapter::from_trait(value: dyn Reader) -> ReaderAdapter`. Callers explicitly coerce concrete implementations to `dyn Reader`, then call `into_go()` when a raw `io.Reader` is needed. `from_go` continues to preserve an existing raw interface unchanged. A bridge created from a trait object is a new non-nil Go interface value, even if that trait object originally wrapped a nil raw interface; it must not silently unwrap or replace its captured receiver. `from_trait` is a reserved generated member. The implementation is connected to module compilation; reader, GoLibrary runtime and full D1 CI checks pass.

Each bridge stores one typed forwarding closure per mapped method. The closures retain the same supplied trait object and dispatch through its existing GoML method table. Ordinary lambda lifting and Go references retain captured state; the bridge does not add a global registry or FFI lock. Go methods use the verified native names, parameters and result lists. They invoke the stored closures and unpack GoML tuple returns into native multiple results, preserving all value and error components. Unit maps to zero native results. Panic propagates without recovery.

The emitted Go AST must include a static interface-satisfaction assertion for the concrete bridge type. Method declarations, struct fields, factory functions and the assertion must use Go AST nodes, deterministic collision-safe names and portable declaration origins. Unexported interface methods, unsupported associated types, non-object-safe receivers, missing mappings and boundary mismatches remain recoverable errors before publication, including with signature checking disabled.

The compiler representation must carry a structured bridge construction plan containing the native interface identity and each mapped method signature. Do not encode executable Go fragments in attributes or reconstruct conversion policy in the backend. Persist the plan through interface/Core codecs, substitution and specialization, then detach it into representation-level types for Lift/ANF validation and emission. Any compiler-generated construction helper must be unavailable as a user-authored escape hatch. Imported plans must undergo structural validation, runtime-only/comptime restrictions and current Go metadata revalidation just as source plans do.

Source inspection identifies existing mechanisms to reuse: `tast/go_interface.gom` owns the complete native method plan; `hir/go_interface.gom` constructs ordinary wrapper and forwarding syntax; `lift/detach.gom` translates typed callables into representation adapters; `repr/callback.gom` and `go_backend/callback.gom` validate and emit typed closure conversion; `go_backend/source_origins.gom` attaches generated helper origins. Callback tuple/unit conversion is a reusable emission rule, not permission to equate GoML closures or dyn values with raw Go ABI values.

D1 acceptance must execute ordinary Go retaining and invoking a GoML reader after the creating call returns; partial reads with non-nil errors; captured mutable state shared with GoML; explicit registration removal; asynchronous calls, reentry and race-checked concurrent calls with properly synchronized fixture state; and identity-preserving panic propagation in the same goroutine. Repeat across package artifacts and GoLibrary pruning. Exercise complete inherited interfaces, unimplementable private methods, invalid receivers/signatures, unsupported associated types and stale Go method metadata in both required and off modes. These runtime and negative-contract checks are implemented. The full CI passes; acceptance-d1.md maps each requirement to its tests and records the final result separately from earlier checkpoints.
