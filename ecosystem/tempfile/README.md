# tempfile

A temporary-file library written entirely in GoML, inspired by [Rust tempfile](https://docs.rs/tempfile/latest/tempfile/). It uses the public standard library's secure random generator, Linux descriptor APIs, and composable I/O traits. No native adapter or external Go dependency is needed.

The implementation targets Linux amd64, matching the current GoML toolchain. It supports named and anonymous temporary files, private temporary directories, atomic persistence, explicit ownership transfer, scoped cleanup, and memory-backed files that spill to disk.

## Example

```goml
use ecosystem::tempfile;
use std::bytes;
use std::io;
use io::{Read, Write};

fn example() -> Result[string, tempfile::ScopeError] {
    tempfile::with_tempfile(
        tempfile::Builder::new().prefix("report-").suffix(".txt"),
        |file| {
            file.write_all(bytes::Bytes::from_string("hello").as_slice())?;
            file.rewind()?;
            file.read_to_string(1024)
        },
    )
}
```

Add `"ecosystem::tempfile" = "0.1.0"` to the module manifest after publishing to a registry. The repository verification runner supplies an isolated local registry snapshot for development.

## API

| API | Behavior |
| --- | --- |
| `Builder::new()` | Default `.goml-` prefix, empty suffix, 16 secure random bytes encoded as 32 hex digits, 128 collision attempts |
| `.parent(path).prefix(value).suffix(value)` | Select the parent directory and filename fragments |
| `.random_bytes(count).attempts(count)` | Configure 8–64 random bytes and 1–1024 attempts |
| `.tempfile()` / `.tempdir()` / `.anonymous()` | Create a named file, directory, or immediately unlinked file |
| `NamedTempFile::new()` / `new_in(parent)` | Named-file convenience constructors |
| `tempdir()` / `tempdir_in(parent)` / `TempDir::new()` | Directory convenience constructors |
| `tempfile()` / `tempfile_in(parent)` | Anonymous-file convenience constructors |
| `temp_dir()` | Nonempty `TMPDIR`, otherwise `/tmp` |
| `NamedTempFile::path()` / `TempDir::path()` | Absolute path recorded at creation |
| `close()` / `is_closed()` | Explicit, shared, idempotent lifecycle |
| `NamedTempFile::persist(destination)` | Atomic rename without replacing any existing destination |
| `NamedTempFile::persist_overwrite(destination)` | Explicit atomic replacement of an existing file or symlink |
| `NamedTempFile::keep()` | Return `KeptFile { file, path }`, disabling temporary cleanup |
| `NamedTempFile::into_file()` | Unlink the name and return an anonymous `File` |
| `NamedTempFile::reopen()` | Independent file offset, checked against the original device/inode |
| `NamedTempFile::duplicate()` / `File::duplicate()` | Independently closed descriptor sharing the same file offset |
| `TempDir::keep()` | Transfer directory cleanup responsibility and return its path |
| `TempDir::persist(destination)` | Atomic directory rename without replacing an existing destination |
| `with_tempfile`, `with_tempdir`, `with_anonymous` | Run a callback and close its resource before returning |
| `SpooledTempFile::new(limit)` / `with_builder(limit, builder)` | Keep up to `limit` bytes in memory, then use an anonymous file |
| `with_spooled(limit, builder, action)` | Scoped spooled file |

`NamedTempFile`, `File`, and `SpooledTempFile` implement `std::io::Read`, `Write`, and `Close`. They work with `io::copy`, `BufReader`, `BufWriter`, and the trait's exact-read and full-write helpers. `TempDir` implements `Close`.

All file forms provide `len`, `seek(SeekFrom::Start/Current/End)`, `rewind`, `set_len`, and `sync_all`. Negative start offsets and lengths are errors. Seeking beyond EOF is supported; subsequent writes and length extension supply zero-filled gaps. Named files and `File` also provide `sync_data`. `flush` checks that the handle remains usable; writes are already unbuffered, and durability requires `sync_all` or `sync_data`.

Spooled files expose `is_rolled`, `rollover`, and `into_file`. Crossing the limit by writing or extending the length spills once; shrinking does not move a file back to memory. Explicit rollover preserves contents and cursor, including a cursor beyond EOF. A failed rollover keeps the in-memory contents and position for retry. A zero limit spills on the first nonempty write; empty writes neither extend nor spill. `into_file` forces rollover and transfers its anonymous descriptor. The builder's parent directory is accessed only when a disk file is needed.

## Ownership and cleanup

GoML values can be copied freely. Copies of one resource share a synchronized lifecycle: successful `keep`, `persist`, or `into_file` disables every old alias. Closing an old alias cannot close the newly returned `File`. Failed persistence leaves the original object open and still responsible for cleanup, so it can be retried or closed. A reopened or duplicated `File` has an independent close lifecycle; it remains usable after the original named file is removed.

There are no GC finalizers or automatic destruction assumptions. Every returned `File`, named file, directory, or spooled file must be closed, explicitly transferred, or managed by a scope helper. A plain `defer { let _ = resource.close(); };` ensures cleanup but discards any cleanup error; call `close` explicitly when reporting cleanup failures matters.

Scope callbacks return `Result[T, io::Error]`. The helper returns `Result[T, ScopeError]`, retaining both `action` and `cleanup` errors when both fail. Creation errors occupy `action`. A callback may deliberately `keep` or `persist` its resource; the helper then has nothing left to clean up. Resource aliases that escape an ordinary scope are already closed. A deferred close also runs during ordinary panic unwinding, but cleanup is not promised on abrupt process termination or signals.

`close` always releases owned descriptors, including when removal fails. It returns that failure and becomes closed; it does not retry removal on a later close. The caller may inspect and recover any remaining path. Cleanup cannot delete a renamed resource under an unknown new name. If the original pathname was removed, closing succeeds; if it was replaced, closing reports `InvalidData` and leaves the replacement untouched.

## Filesystem guarantees and limits

Names come from `std::crypto::rand`, with at least 64 bits of random input and 128 bits by default. File creation uses `O_CREAT | O_EXCL | O_NOFOLLOW` and mode `0600`; directory creation uses atomic `mkdirat` and mode `0700`, subject to a more restrictive process umask. Existing names are never opened or truncated. Prefixes and suffixes reject slash, backslash, and NUL; complete names are limited to 255 bytes. Missing parents, entropy failures, unsupported operations, and exhausted collision attempts return errors.

A parent directory descriptor is retained until cleanup or transfer. Operations continue to address that directory if its pathname is renamed. Directory removal uses descriptor-relative traversal with no-follow opens; symbolic links, including dangling links and links to outside directories, are removed as links. Raw directory-entry bytes support cleanup of non-UTF-8 names. A depth limit of 256 bounds recursive cleanup; deeper trees return an error and may be partly removed.

Persistence uses Linux `renameat2`; `persist` sets `RENAME_NOREPLACE`, including for existing symlinks and empty directories. `persist_overwrite` replaces a symlink itself rather than its target. Cross-filesystem moves and kernels/filesystems without the required rename operation return an error; there is no non-atomic copy fallback. Persistence does not fsync the file or destination directory automatically, so atomic visibility is separate from crash durability.

The recorded `path()` can become stale after parent renames. Device/inode checks detect pathname substitution before reopening, cleanup, and transfer, but they do not provide a transaction against hostile concurrent filesystem changes. Use a trusted parent for long-lived named resources, keep untrusted code from modifying their directories, and prefer anonymous files when a name is unnecessary. Directory cleanup is not a sandbox and assumes no concurrent hostile renames or mount changes.

Lifecycle methods are synchronized for copies of the same object. Named-file `read_exact` and `write_all` hold that object's lock for the entire operation. Separate `seek` and `read` calls are distinct operations, and duplicated descriptors share an offset. Do not concurrently mutate input/output byte buffers supplied to any I/O method. There is no async or cancellation-specific file API.

## Local validation

From the repository root:

```sh
python3 ecosystem/verify.py tempfile
python3 ecosystem/tempfile/interop.py
python3 ecosystem/tempfile/race.py
```

The library has 24 external tests covering permissions, validation, complete file I/O, sparse writes, reopen/duplicate semantics, no-clobber persistence, symlink behavior, non-UTF-8 cleanup, replaced/deleted names, parent renames, scope errors, transfer races, concurrent creation, descriptor counts, spill thresholds, failed rollover recovery, and concurrent spooled writes. An independent consumer tests the public registry boundary and cached artifact reuse. Python cross-checks binary contents, private permissions, cleanup, and 160 concurrent persists across five isolated directories; the race script runs every library test under Go's race detector.
