# ecosystem::notify

A GoML library extracted from the former `std::fs::notify` package. The public
API and Linux amd64 behavior are preserved; the import and dependency coordinate
now belong to the ecosystem. It uses the published GoML 0.1.50 standard library,
including public `std::os::linux::syscall`, without compiler or private runtime
imports, generated native bindings, or additional host libraries.

## Installation and migration

Declare the dependency in your module-root `goml.toml`:

```toml
[dependencies]
"ecosystem::notify" = "0.1.0"
```

Replace `use std::fs::notify;` with `use ecosystem::notify;`. Existing calls and
public types retain their names and signatures. Source archives in this repository
are resolved by the ecosystem verifier's isolated local registry.

## Filesystem notifications

`ecosystem::notify` implements filesystem notifications in GoML using `std::os::linux::syscall` and Linux inotify. The supported target is Linux amd64. `watch(path: string)` watches one file or a directory and its immediate entries. `watch_recursive(path: string)` requires a directory, watches existing descendants, and adds watches when directories are created or moved into the tree. Both return `Result[Watcher, fs::Error]`. Descendant symbolic links are not traversed, and a symbolic link as the root is rejected. Directory aliases through bind mounts are unsupported.

`watch_with(path, options: Options)` configures a single watcher. `Options::new()` uses `recursive = false` and `mask = CHANGES`; `with_recursive(bool)` and `with_mask(u32)` return adjusted options. Both fields are public. A requested mask must be a nonempty subset of `ALL_EVENTS`. The available request bits are `ACCESS`, `MODIFY`, `ATTRIB`, `CLOSE_WRITE`, `CLOSE_NOWRITE`, `OPEN`, `MOVED_FROM`, `MOVED_TO`, `CREATE`, `DELETE`, `DELETE_SELF`, and `MOVE_SELF`. `CHANGES` includes these except `ACCESS`, `CLOSE_NOWRITE`, and `OPEN`. Recursive directory topology events, root lifecycle events, and recovery notifications are always delivered even when excluded by the requested filter, so filtering cannot disable recursive maintenance.

`Event` has public `path: string`, `mask: u32`, `cookie: u32`, and `rescan: bool` fields. Paths are absolute. `has(mask)` tests whether any requested mask bit is present, and `is_dir()` tests `IS_DIR`. In addition to requested event bits, masks may contain `UNMOUNT`, `Q_OVERFLOW`, `IGNORED`, or `IS_DIR`. Matching nonzero cookies connect `MOVED_FROM` and `MOVED_TO` events within a registration; cookies are not persistent object identities.

Build ignore rules with `Options::new().with_ignored_names(Vec::from_array([".git", "node_modules"]))` or `with_ignore((absolute_path: string, is_directory: bool) -> bool)`. Repeated calls combine rules with OR. Name rules match literal basenames at every depth and snapshot the supplied vector; they are not glob patterns. An ignored directory prunes its entire subtree, avoiding recursive watches and scans there. Rules apply to initial discovery, new directories, renames, and overflow recovery, including registrations in `WatchSet` and subscriptions. Moving a visible directory into an ignored path removes its subtree watches; moving it back installs watches and requests a rescan. The root itself and recovery signals are never ignored. Construct options through `new()`; predicates must be stable, quick, and must not call back into their watcher because they run while its state is locked.

`Watcher.try_read() -> Result[Vec[Event], fs::Error]` reads one available batch without waiting for new kernel events. `Watcher.read(timeout: time::Duration)` has the same return type and waits for a nonempty batch, returning an empty vector on timeout. A zero timeout checks immediately. `read_with(cancel: task::CancelToken, timeout)` returns `Result[task::WaitResult[Vec[Event]], fs::Error]`; cancellation returns `Cancelled` and leaves the watcher available. A read that wins a race with cancellation may return `Completed(events)`. Waits check cancellation and close in intervals of at most 50 ms. Timeouts and cancellation bound waiting for kernel events, not directory scans or time spent waiting for another operation to release shared state.

Single-watcher copies share a synchronized handle. Concurrent reads divide the event stream. `close() -> Result[(), fs::Error]` releases the descriptor and all watches and is idempotent; `is_closed()` reports terminal state. Use explicit close or `defer`; garbage collection does not close watchers. A single-file watcher follows the watched inode until a terminal event. Editors commonly replace a file atomically, which ends that file watch; monitor its parent directory and filter event paths when replacement must be followed.

```goml
use std::fs;
use ecosystem::notify;
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
use ecosystem::notify;
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

## Validation

From the repository root:

```sh
just ecosystem-test notify
```

The independent consumer preserves the complete former compiler regression
fixture and exercises the package through a normal versioned dependency. The module retains all 22 original internal tests and adds 8 public-API
tests for filesystem boundaries and resource lifecycle. The GoML verifier compiles
the generated library and consumer test runners with Go's race detector and runs
the same real-filesystem cases. Tests create unique temporary directories and
remove them after each run; no external service or privileged mount is required.

Coverage includes real inotify queue overflow and rebuild, cross-batch rename
cookies, recursive move-in/out and ignored subtree maintenance, malformed record
decoding, raw non-UTF-8 filenames, cancellable reads, unbuffered subscription
backpressure, concurrent close and descriptor release.
