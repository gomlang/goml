# C0 and C1 callback acceptance

Scope: explicit conversion between GoML closures and nullable native Go functions, including retention, delayed invocation, concurrency and reentry. C2 panic/provenance/performance acceptance and interface adapters are outside this audit.

## Representation and conversion

`ffi::Func[F]` uses a distinct GoFunction representation through TAST, Core, Mono and repr. The signature is preserved in artifacts and checked at source, artifact, specialization and backend boundaries. Native unit returns have zero Go results, and flat tuple returns have separate Go results. Raw function values cannot be invoked as GoML closures or used as comparable map keys; Some(nil) remains distinct from None.

`func_from_closure` and `func_to_closure` retain the existing lambda-lifted function representation. Detachment introduces a typed CallbackAdapter callable, validated in Lift and ANF. The backend generates a signature-specific holder and Apply method; its bound method value strongly retains the stored function and captured environment. Reverse conversion checks nil and returns None before constructing an adapter. There is no reflection, global function registry or global FFI lock. Conversion rejects ordinary string/char signatures requiring implicit text policies; callers use raw ffi::String and ffi::Rune.

The driver regression `project_callback_conversions_preserve_captures_results_and_nil` covers captured and uncaptured closures, shared Ref state, distinct same-signature closures, generic conversion in another package, repeated conversion, Go-returned functions, nil, unit, multiple results and original error identity. It also verifies invalid UTF-8 bytes through a raw string callback and rejects implicit text conversion. Independent compiler tests validate bridge signatures, artifact transport, adapter direction and Go result shapes.

## Lifetime and concurrency

The runnable [subscription example](../../examples/ffi-callbacks/README.md) stores a converted closure in a Go object after its creating GoML function returns. A channel gate delays dispatch until after Start returns and an explicit GC completes. No finalizer timing is asserted. Sixty-four workers invoke the same retained callback with an immutable capture and an atomic counter. Recursive Go/GoML reentry reaches depth 32; self-unregistration confirms that the subscription mutex is released before invocation. Another 128 registration cycles verify expected calls and balanced explicit registration accounting.

Unregister clears the owned callback reference and is idempotent. Queued batches released after unregister and subsequent batches dispatch no calls. An invocation already selected by a worker may still start or finish after unregister; Wait must join active batches before their captured resources are disposed. This is the example's explicit cancellation contract. GC reachability does not cancel registrations, and FFI does not make unsynchronized mutable state thread-safe.

The driver regression `callback_example_retains_unregisters_and_reenters_without_races` copies the committed example into an isolated workspace, executes fresh and cached builds, compares generated source and runs Go's race detector. The test uses synchronization rather than sleeps. Extracted-release smoke additionally checks formatting, required FFI validation, repeat execution and generated-source digests using the packaged toolchain with network access disabled for Go dependency resolution.

## Evidence and limits

Focused adapter and lifetime regressions passed in `/tmp/goml-callback-adapter-test2.log` and `/tmp/goml-callback-lifetime-test.log`. The adapter implementation passed full CI with 848 compiler tests and 126 driver tests in `/tmp/goml-callback-adapter-ci.log`. The lifetime regression brings the driver count to 127. Final CI including extracted-release callback smoke is recorded in the implementation log.

These checks establish the supported C0/C1 contract. They do not establish callback performance, source provenance, explicit panic tests or the remainder of C2. B3 deferred map constraints, remaining B4 acceptance and batch D remain separately tracked.
