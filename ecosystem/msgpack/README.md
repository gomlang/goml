# MessagePack for GoML

A native GoML implementation of the [MessagePack wire specification](https://github.com/msgpack/msgpack/blob/master/spec.md).
The dynamic codec supports every wire type. The typed codec implements
`std::serde::Serializer` and `Deserializer` directly, including generic structs,
enums, tuples, options, sequences, binary data and maps with arbitrary key types.
No Go FFI or runtime reflection is used. Python is only a verification dependency.

## Typed API

```gom
use ecosystem::msgpack;
use ecosystem::msgpack::{Serialize, Deserialize, Binary};

#[derive(Serialize, Deserialize)]
struct Record[T] {
    name: string,
    payload: T,
}

fn example() -> Result[Record[Binary], msgpack::Error] {
    let original = Record {
        name: "中文😀",
        payload: Binary { data: Vec::from_array([0, 255]) },
    };
    let encoded = msgpack::to_vec(original)?;
    msgpack::from_slice(encoded.as_slice())
}
```

`to_vec_with_options` and `from_slice_with_options` accept `Options`.
`Encoder::new`, `Encoder::finish`, `Decoder::new`, `Decoder::position` and
`Decoder::finish` expose the Serde event interfaces for custom implementations.
Import the standard `Serializer`/`Deserializer` traits to call their methods.
Handles share mutable state; use one caller at a time and discard a handle after
an error. A successful finish requires exactly one complete root. Decoding also
rejects trailing bytes.

| GoML value | Default wire representation | Alternative |
| --- | --- | --- |
| Integers | Smallest fitting signed/unsigned representation | Positive signed values use unsigned wire tags |
| `f32`, `f64` | `float32`, `float64` | Width and bits retained when decoding to the same type |
| `()`, `bool`, `string`, `char` | nil, boolean, UTF-8 string, one-scalar string | No implicit string/number conversion |
| `Vec[T]`, tuples | array | Tuple length must match exactly |
| `Vec[u8]` | array of integers | `Binary { data }` uses the binary wire family |
| `Pairs[K, V]` | map | Ordered entries, arbitrary serializable keys, duplicate pairs preserved |
| Struct | map keyed by field name | `structs_as_arrays` writes positional arrays |
| Unit enum variant | variant name string | `variants_as_indices` writes declaration index |
| Tuple enum variant | `{ name: [fields...] }` | Index tag; one-field variants still contain a one-element array |
| Named enum variant | `{ name: { field: value, ... } }` | Index tag and/or positional payload array |
| `Option[T]` | nil for `None`, transparent payload for `Some` | `tagged_options`: `[0]` or `[1, payload]` |

Structs and named variants accept either map or positional input regardless of
output options. Enum tags accept names or indices. Named fields can arrive in
any order; unknown fields are skipped without constructing a value tree.
Duplicate names, including unknown names, are rejected. Derives reject missing
fields and unknown variants; `#[serde(rename = "...")]` works on fields and
variants. Tuple payload lengths and custom event order/counts are checked.

The default nil option representation cannot distinguish `Some(())`,
`Some(None)` and `None`. Enable `tagged_options` at both ends when this distinction
matters. This option changes the wire convention; it is not an automatic decoder
heuristic. Enum and struct conventions are likewise library conventions on top
of MessagePack's data model.

Signed/unsigned integer conversions check the destination range. Floats accept
only floating wire types; conversion between float widths uses IEEE rounding
and can overflow to infinity when narrowing. Integers are not silently converted
to floats. Numeric-text Serde events parse checked integers or floating literals.

`Pairs::from_hash_map` snapshots entries in the hash map's iteration order.
`Pairs::to_hash_map` rejects duplicate keys instead of silently overwriting them.
Use ordered `Pairs.entries` when encoded byte order must be reproducible.

## Dynamic API and timestamps

`Value` represents nil, booleans, `Int(i64)`, `Uint(u64)`, both float widths,
strings, raw strings, binary data, arrays, ordered map pairs and
`Extension(i8, Vec[u8])`. Map keys may themselves be arrays, maps or other values.
The API preserves duplicates and does not hash untrusted dynamic keys.

- `pack` / `pack_with_options`: encode one value.
- `unpack` / `unpack_with_options`: decode exactly one value.
- `unpack_prefix(input, options)`: return the first value and consumed byte count.
- `skip_prefix(input, options)`: validate and skip the first value without a tree.
- `Timestamp { seconds: i64, nanoseconds: u32 }.to_value()`: shortest valid
  timestamp extension, using 32-, 64- or 96-bit payloads.
- `Timestamp::from_value`: require extension type `-1`, a valid payload length and
  nanoseconds below one billion. Negative seconds and the full signed 64-bit
  seconds range are supported.

The reader accepts legal nonminimal integer/container encodings; the writer
normalizes integer widths and length headers. It preserves extension payloads,
float widths, NaN payload bits and negative zero. Extension types are opaque to
the ordinary codec, including reserved negative types. Timestamp validation is
explicit so unknown extension values remain transportable.

UTF-8 is validated by default. `allow_invalid_utf8` decodes invalid string bytes
as `RawString` and permits writing them. Valid raw input becomes `String`.
Typed string/char decoding still requires valid UTF-8 even with this option.

`std::serde::Value` has no binary, extension or arbitrary-map-key representation.
Its `deserialize_any` path rejects those cases explicitly. Use `msgpack::Value`
for full dynamic fidelity, `Binary` for typed binary fields and `Pairs` for typed
maps. The current standard Serde event protocol has no extension event, so
extensions/timestamps use the dynamic API rather than a magic struct convention.

## Incremental input and limits

`Unpacker::new(options)`, `push(Slice[u8])`, `next()` and `finish()` handle
concatenated values split at any byte boundary. `next()` returns `None` while a
value is incomplete; repeatedly drain it after each push. `finish()` requires an
empty, fully drained stream. Malformed input or buffer overflow poisons the
unpacker until `reset()`. `buffered()` counts only unconsumed bytes.

The scanner retains its token cursor and outstanding child counts between
chunks. It materializes a value only once complete. Consumed messages advance a
buffer head, avoiding a tail copy on every `next()` call. Pushes compact consumed
storage when needed. This is bounded buffering of complete values, not an API
that streams individual fields to a callback.

`Limits::standard()` sets:

| Limit | Default |
| --- | ---: |
| Whole input / encoded output / unconsumed stream buffer | 16 MiB |
| Individual string, binary or extension payload | 16 MiB |
| Array elements or map entries | 1,000,000 |
| Wire values, counting keys and container headers | 2,000,000 |
| Nesting depth, with root at depth zero | 128 |

All limits are configurable nonnegative integers. Lengths also obey the
32-bit wire ceiling. Prefix APIs apply the byte limit to their entire supplied
slice. Streaming byte limits apply to the currently buffered input, while value
and depth counts restart for each message. Serde also bounds its open event
frames, including transparent options, to prevent unbounded logical nesting.

`Decoder::new` performs an iterative validation pass before typed code receives
container length hints. This detects truncated/oversized/deep input before a
standard `Vec` deserializer preallocates from those hints. Direct decoding then
reads the same bytes without an intermediate dynamic tree. Encoder finish uses
an iterative validation pass to check the resulting wire nesting. Dynamic
materialization is recursive under the depth limit; skipping and stream scanning
are iterative. Raising limits also raises allowed memory and stack consumption.

Errors expose `kind`, byte `offset` and `message`. Wire validation retains its
specific error kind; custom/typed Serde failures use `ErrorKind::Serde`.
Streaming decode offsets refer to the current value, not a lifetime stream byte
counter.

## Verification

From the repository root:

```sh
python3 ecosystem/verify.py msgpack
python3 ecosystem/msgpack/interop.py
```

The library has 14 external tests covering integer boundaries, all length
families (including 65,536-element maps/arrays), floating bits, Unicode and raw
strings, arbitrary/duplicate map keys, timestamps, every truncation of selected
nested values, limits, streaming compaction, all chunk sizes, direct Serde
derives, event protocol errors and tagged options.

The separate consumer resolves a normal versioned dependency from the isolated
registry snapshot. It tests downstream derives and generic specialization;
the verifier also builds, runs and checks unchanged cached build artifacts.

The interoperability harness uses checksum-pinned
[msgpack-python 1.1.2](https://pypi.org/project/msgpack/1.1.2/), extracted only under
`ecosystem/_artifact/reference/`. It performs 2,490 checks, including every valid
first-byte tag, 1,234 truncations, 400 generated recursive values, 480 typed
map/positional cases, timestamp formats, raw strings, and 20,000 concatenated
small messages. GoML output is checked against independent reference bytes and
read by the Python implementation. Reports go under
`ecosystem/_artifact/verification/msgpack/`.
