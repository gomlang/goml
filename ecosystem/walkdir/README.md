# ecosystem::walkdir

A GoML library extracted from the former `std::fs::walkdir` package. The public
API and Linux amd64 behavior are preserved; the import and dependency coordinate
now belong to the ecosystem. It uses the published GoML 0.1.50 standard library,
including public `std::os::linux::syscall`, without compiler or private runtime
imports, generated native bindings, or additional host libraries.

## Installation and migration

Declare the dependency in your module-root `goml.toml`:

```toml
[dependencies]
"ecosystem::walkdir" = "0.1.0"
```

Replace `use std::fs::walkdir;` with `use ecosystem::walkdir;`. Existing calls and
public types retain their names and signatures. Source archives in this repository
are resolved by the ecosystem verifier's isolated local registry.

## Directory traversal

`ecosystem::walkdir` implements lazy, iterative depth-first traversal in GoML using Linux amd64 syscalls: `openat`, `getdents64`, `newfstatat`, `fstat`, and `close`. Other targets are unsupported; it does not fall back to host directory-listing or metadata APIs. `walk(root: string) -> WalkIterator` uses the defaults from `WalkDir::new(root)`: include the root at depth zero, visit directories before their contents, sort siblings by file name, impose no depth limit, and do not follow symbolic links, including a root link. A file root yields one entry. Construction performs no filesystem I/O; the first iterator step checks the root.

`WalkDir` is a reusable value builder. Each `iter()` or `IntoIterator::into_iter()` creates an independent `WalkIterator`, so it can also be used directly in `for`. Builder methods return a modified value:

| Method | Behavior |
| --- | --- |
| `min_depth(isize)` | Hide entries shallower than the bound while still traversing them. |
| `max_depth(isize)` | Include this depth but never read directories below it. |
| `follow_links(bool)` | Follow file and directory symbolic links, including the root. |
| `contents_first(bool)` | Yield a directory after its descendants instead of before them. |
| `filter_entry((DirEntry) -> bool)` | Reject an entry and prune its entire subtree. Repeated filters are combined with short-circuiting AND. |

Negative bounds or `min_depth > max_depth` produce one `InvalidInput` error before filesystem access, followed by exhaustion. The predicate runs after metadata is obtained, before descending, and even for entries hidden by `min_depth`; it also prunes correctly with `contents_first(true)`. Use `std::iter::filter` when only output should be filtered without pruning.

`WalkIterator` implements `Iterator` with `Item = Result[DirEntry, Error]` and composes with `std::iter`. In preorder, call `skip_current_dir()` immediately after receiving a directory to skip its descendants. The directory is not opened until the next `next()` call. Skipping before iteration, after a file or error, or in contents-first order does nothing. `close()` releases all retained directory descriptors, clears pending work, and permanently exhausts the iterator; it is idempotent. Complete exhaustion also releases all descriptors. When stopping early, including after an error, call `close()` or use `defer`; garbage collection does not close descriptors. Explicit `close()` attempts every descriptor and discards close errors; a close failure during normal iteration is returned as an error. On Linux, close is never retried, including after `EINTR`. Exhaustion is permanent. Copies of an iterator share progress and must be used by one consumer; copies of the builder can start independent walks.

`DirEntry` exposes `path()`, `file_name() -> Option[string]`, `depth()`, `file_type()`, `metadata()`, and `path_is_symlink()`. Metadata is a snapshot taken before yielding the entry. When following a link, `file_type()` and `metadata()` describe the target while `path_is_symlink()` remains true. Paths retain the supplied root spelling and use `path::join` for children; they are not replaced with canonical target paths.

`Error` exposes `path()`, `depth()`, `kind()`, `fs_error()`, and `loop_ancestor() -> Option[string]`, plus `ToString` and `Debug`. The underlying `fs::Error` retains its operation and available raw OS code. A failed entry or directory emits an error and traversal continues with remaining branches. Errors are not suppressed by `min_depth`. A directory-read error follows its entry in preorder and precedes its entry in contents-first order. Device and inode identities from opened directory descriptors detect ancestor loops, including directory links and bind-mount aliases; these return `InvalidData` with the offending path and the ancestor's traversal path, then skip that subtree. Sibling aliases to the same directory are each traversed. Dangling links are ordinary link entries by default and metadata errors when followed; a link chain rejected by the OS reports its filesystem error without a loop ancestor.

```goml
use ecosystem::walkdir;
use std::io;

fn main() -> () {
    let tree = walkdir::WalkDir::new(".")
        .max_depth(8)
        .filter_entry(|entry| entry.file_name() != Option::Some(".git"));
    let iterator = tree.iter();
    defer iterator.close();
    for item in iterator {
        match item {
            Ok(entry) => println(entry.path()),
            Err(error) => io::eprintln(error.to_string()),
        }
    }
}
```

Directory entries are read in 32 KiB `getdents64` batches, decoded with record-length and filename validation, and sorted before traversal. Full metadata is still read for each yielded entry to preserve `metadata()` snapshot semantics; `d_type` is not used to omit these queries, so `DT_UNKNOWN` requires no special fallback. Invalid UTF-8 filenames report `InvalidData` for the directory. Paths remain UTF-8 strings, and embedded NUL bytes are rejected before calling the kernel.

One descriptor per active ancestor is retained until its subtree finishes. Descriptors use `O_CLOEXEC`; child lookup uses the parent descriptor, so renaming an opened ancestor does not redirect traversal to a replacement at its old path. Returned paths still reflect the original traversal spelling. Opening a directory checks that its identity matches the earlier metadata; a changed directory produces an error and is skipped. Without link following, `O_NOFOLLOW` also rejects replacement of the final directory component by a symbolic link. Path prefixes and explicit links follow kernel resolution rules; this is not a confinement API or an atomic filesystem snapshot.

Traversal uses an explicit stack rather than recursive calls; memory scales with pending names along the active branch. Very deep trees can reach the process descriptor limit, which produces an OS error for that branch; `max_depth` can bound this usage. Interrupted open, metadata, and directory-read calls retry. Entry failures leave other branches available, so callers must still close an iterator when abandoning it after an error. This API uses ordinary imports, closures, traits, and `for` syntax and adds no grammar or compile-time filesystem access. The native layouts and lifecycle rules follow [getdents64](https://man7.org/linux/man-pages/man2/getdents.2.html), [stat](https://man7.org/linux/man-pages/man2/stat.2.html), [openat](https://man7.org/linux/man-pages/man2/open.2.html), and [close](https://man7.org/linux/man-pages/man2/close.2.html).


## Validation

From the repository root:

```sh
python3 ecosystem/verify.py walkdir
```

The independent consumer preserves the complete former compiler regression
fixture and exercises the package through a normal versioned dependency. The module retains all 8 original internal tests and adds 6 public-API
tests for filesystem boundaries and resource lifecycle. `race.py` compiles
the generated library and consumer test runners with Go's race detector and runs
the same real-filesystem cases. Tests create unique temporary directories and
remove them after each run; no external service or privileged mount is required.

Coverage includes multi-batch directory reads, native stat decoding, Unicode and
invalid-byte filenames, symlink loops and aliases, descriptor-relative traversal
after ancestor renames, replaced-directory rejection, depth bounds and pruning,
independent parallel iterators, and descriptor release on errors or early close.
