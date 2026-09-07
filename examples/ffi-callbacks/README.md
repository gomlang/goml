# Retained and concurrent Go callbacks

From the repository root:

```sh
just make
cd examples/ffi-callbacks
../../stage2/bin/goml run
```

The example uses only standard Go packages and its local shim. Its output is:

```text
deferred until release: true
concurrent retained callbacks: true
unregistered callbacks stopped: true
recursive reentry: true
callback can unregister itself: true
registration stress balanced: true
```

`retained` creates a GoML closure and returns a Go subscription that owns the converted function. The local closure is no longer held by its creating function. `Start` queues a batch on a goroutine and returns before any callback runs. `Release` opens a gate; `Wait` joins the batch and reports its number of dispatched callbacks. Every started batch must be released and waited. Wait is single-use; Release is idempotent. The example forces a GC between Start and Release without relying on finalizers.

The first batch invokes one retained callback from 64 goroutines. The captured Go `Counter` uses atomic operations. The captured offset is immutable. Ordinary GoML `Ref` cells are not made thread-safe by FFI conversion; concurrent mutation needs the same synchronization as other shared state.

`Unregister` clears the subscription's callback reference and is idempotent. Workers obtain the callback under the subscription mutex, release that mutex, and then invoke it. This permits a callback to unregister itself and permits Go↔GoML recursive reentry. A callback already selected by a worker may still start or finish after Unregister; wait for active batches before disposing of resources those callbacks use. A queued batch released after Unregister dispatches no callbacks. New batches also dispatch none.

The self-unregistering callback reads a Ref that is initialized before Release. The channel gate supplies the required ordering; the Ref is not concurrently mutated. The example then performs 128 register/dispatch/unregister cycles, verifies every expected counter increment, and checks balanced registration accounting. The shim's global counter stores only a count, not references or handles. The subscription object holds the function directly, and clearing that field releases the subscription's reference. Reachability and explicit registration lifetime remain separate concerns.

To run the generated program with Go's race detector on a supported host:

```sh
go run -race _artifact/build/pkg/ffi_callbacks/goml_generated.go
```

The driver regression copies these committed sources into an isolated workspace, runs fresh and cached builds, checks stable generated source and executes this race command with CGO enabled. It checks output as well as process success and does not depend on wall-clock sleeps or finalizer timing.
