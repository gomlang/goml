# config

Layered, typed configuration implemented in GoML. JSON and TOML use the standard
format libraries; file notifications use `ecosystem::notify = "0.1.0"`.

```goml
use ecosystem::config;

let builder = config::Builder::new()
    .with_source(config::Source::json("defaults", "{\"port\":8000}"))
    .with_source(config::Source::file("local", "app.toml", config::Format::Toml, false)?)
    .with_source(config::Source::environment("environment", config::EnvOptions::new("APP_")))
    .with_source(config::Source::overrides("CLI", Vec::from_array([
        config::Override::value("/port", 9000)?,
    ])))
    .validate_with(|snapshot| {
        let port: isize = snapshot.decode_at("/port")?;
        if port < 1 || port > 65535 {
            Result::Err(snapshot.invalid("/port", "port must be between 1 and 65535"))
        } else {
            Result::Ok(())
        }
    });
let live = config::Live::new(builder)?;
let saved = live.snapshot();
let port: isize = saved.decode_at("/port")?;
let origin = saved.origin("/port")?;
```

## Sources and precedence

`Builder::with_source` returns a new builder and appends a source. Later sources
win. All document roots must be objects. Objects merge recursively; scalars and
type changes replace the earlier value. Arrays replace by default;
`with_arrays(ArrayMerge::Append)` concatenates them while retaining each
element's origin. An empty object preserves existing object members. JSON null
is a value, never a deletion marker.

| Source | Behavior |
| --- | --- |
| `Source::defaults[T: Serialize](value)` | Captures JSON-compatible typed defaults as an owned serialized value |
| `Source::json(name, text)` / `Source::toml(name, text)` | Parses the supplied document on each build |
| `Source::file(name, path, format, required)` | Captures an absolute path; reopens and reads it on each build; an optional missing file is skipped |
| `Source::environment(name, options)` | Reads the current process environment on each build |
| `Source::environment_values(name, options, entries)` | Captures supplied entries for deterministic tests and caller-managed environments |
| `Source::overrides(name, entries)` | Applies ordered `Override::value`, `Override::text`, and `Override::unset` operations; callers use their CLI parser to construct these entries |

An optional file suppresses only `NotFound`. Malformed data, permission errors,
directory reads, and other I/O failures are errors. Files use bounded standard
`Read` operations over Linux descriptors; a nonblocking open avoids blocking on
named pipes. Filesystem paths are fixed when the source is created.

Environment options specify a literal prefix, nonempty hierarchy separator,
Unicode lowercase mapping, and a value mode. Defaults map `APP_DATABASE__PORT`
to `/database/port`. `Auto` accepts valid JSON values and keeps other input as an
unchanged string. `String` never coerces. `Json` rejects invalid JSON. Empty key
segments, duplicate mapped keys, and parent/child collisions within one
environment source are errors. Selected entries are applied in sorted path
order; operating-system environment order does not affect results.

## Paths, typing and provenance

Paths are JSON Pointers: `""` identifies the root, `/` an empty object key,
`/a~1b/~0key` the keys `a/b` and `~key`. Dot characters are ordinary key data.
Array indices must be canonical nonnegative decimals. A final `-` appends an
override to an existing array. Creating missing paths creates objects, never
guesses an array. Traversing a scalar or an out-of-range index is an error.
Removing an absent object member is a no-op; array removal shifts later
elements and retains their origins. Removing the root is rejected.

`Snapshot::decode[T: Deserialize]` decodes the root; `decode_at[T]` decodes a
subtree. Serde errors retain the requested pointer and source, with the standard
decoder's nested field context in their message. `validate_with` hooks run in
registration order after all sources merge. Hooks can decode whole settings or
individual values and produce `snapshot.invalid(path, message)` diagnostics.

Every node has an `Origin { source, detail }`. File details contain the absolute
path; environment details contain the original variable name. `origin(path)`
returns the winning node's origin, and `origins()` returns all current node
paths, including containers. Merged containers identify the latest contributing
source; untouched children keep earlier origins. Parse diagnostics retain the
source and standard parser's location text; origins are source-level metadata,
not line/column source maps.

`get(path)` returns detached JSON containers. Decoded vectors and `origins()`
results may be mutated without affecting the snapshot. `json()` uses recursively
sorted object keys. Builder copies, captured input vectors, and snapshots do not
share publicly mutable storage. User-defined Serde implementations and validator
callbacks remain responsible for their own external state.

## Reload and notifications

`Live::new(builder)` validates an initial snapshot at revision 1. `reload()`
serializes concurrent reload operations, loads and validates outside the state
lock, then publishes the whole new snapshot. Readers can retain old snapshots.
An unsuccessful reload leaves both the value and revision unchanged and records
`last_error()`. Success clears that error and increments the revision, even if
values are unchanged. `Reload.changed` compares deterministic JSON values and
ignores provenance differences. Numeric spelling is retained, so `1` and `1.0`
are distinct serialized values.

`live.watch()` registers each file's parent directory, then reloads to cover
changes between initial loading and registration. This sees atomic file
replacement and optional-file creation/deletion. The parent directories must
already exist. It returns an explicitly owned `HotReload`:

```goml
use std::time;

let watcher = live.watch()?;
defer { let _ = watcher.close(); };
let result = watcher.poll(time::Duration::from_milliseconds(100))?;
```

`poll` processes
one event batch and reloads once when a watched file changes. Unrelated file
events return `None`; overflow/rescan events trigger a reload. `poll_with` adds
task cancellation. Failed parsing/validation retains the previous snapshot;
applications can report the error and retry explicitly or await another event.
The library does not spawn a background task or impose a debounce interval.
Parent directory invalidation is reported and requires recreating the watcher.
`close()` is idempotent. `Live::last_error` records configuration reload errors;
watch transport errors are returned by the watcher.

## Bounds, scope and validation

`Limits` bounds source count, source/total bytes, cumulative parsed nodes and
merged nodes, and nesting depth. Defaults are 128 sources, 1 MiB per source,
4 MiB total, 100,000 nodes and depth 64; maximum configurable depth is 128.
Text nesting is checked before invoking recursive parsers. TOML also bounds
unquoted dotted-key segments. Byte/node limits include overridden source data,
so many discarded layers cannot bypass them. Override/environment sources are
bounded as a whole. Serialization of user-supplied typed defaults takes place
when that source is constructed, before builder limits apply.

Configuration values use the JSON-compatible data model. TOML support inherits
`std::toml`'s currently documented subset; multiline TOML strings, datetime
literals and nonfinite numbers are not added here. There is no interpolation,
secret provider, YAML/INI parser, automatic CLI flag discovery, or disk writing.
Native file loading and notifications currently target Linux.

The native GoML tests cover type/merge/path errors, all source kinds, independent
provenance, bounds, alias isolation, real files, optional files, invalid reloads,
atomic replacement through inotify, cancellation and concurrent readers/reloads.
The independent `consumer::config` exercises normal versioned resolution.

Run `just ecosystem-test config` from the repository root. No Python helpers or
CI integration are required.
