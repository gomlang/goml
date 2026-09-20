# rope

A pure GoML persistent UTF-8 rope, inspired by [Ropey](https://docs.rs/ropey/latest/ropey/struct.Rope.html). It implements a chunked AVL tree, not a flat string with editing helpers. There are no Go adapters or native dependencies.

The package owns `ecosystem::rope`; the separate `consumer::rope` module resolves version `0.1.0` through the ecosystem registry fixture. The LSP library uses this rope for document storage and edits.

## Representation and snapshots

Every leaf owns at most 1024 UTF-8 bytes and ends at a scalar boundary. Branches cache byte, Unicode scalar, UTF-16 code-unit and line-break counts, boundary CR/LF information, and height. Internal node references remain private and are never mutated. Edits rebuild and balance the affected paths, sharing untouched subtrees. `snapshot()` and ordinary assignment are constant-time immutable snapshots; all editing methods return a new `Rope`. Readers and independent edits can operate concurrently on snapshots.

Leaf creation copies its bounded text into independent storage. A small slice therefore retains its selected subtrees, rather than retaining an unrelated large source string through a Go substring. Splitting copies at most the affected boundary leaves. Small adjacent leaves are merged opportunistically. There is no minimum leaf occupancy guarantee after edits; `compact()` repacks text into nearly full leaves.

## API

```goml
use ecosystem::rope;

fn edit() -> Result[rope::Rope, rope::Error] {
    let original = rope::Rope::from_string("hello 😀\r\nworld")?;
    let snapshot = original.snapshot();
    let changed = original.insert(6, "GoML ")?;
    let middle = changed.slice(6, 10)?;
    let at = changed.position_to_byte(
        rope::Position { line: 0, column: 11 },
        rope::Encoding::Utf16,
    )?;
    let _ = (snapshot, middle, at);
    Result::Ok(changed)
}
```

| Area | Operations |
| --- | --- |
| Construction | `new`, checked `from_string`, bounded generic `from_reader`, `Builder` |
| Lengths and access | `len_bytes`, `len_scalars`, `len_utf16`, `len_lines`, `is_empty`, checked `byte` and `scalar`, `is_char_boundary` |
| Immutable edits | Scalar-indexed `insert`, `remove`; byte-indexed `insert_bytes`, `remove_bytes`, `replace_bytes`; `concat` |
| Structural operations | Scalar-indexed `slice`, `split_at`; byte-indexed `slice_bytes`, `split_at_byte`; `compact`, `snapshot` |
| Index conversion | `byte_to_scalar`, `scalar_to_byte`, `byte_to_utf16`, `utf16_to_byte`, `scalar_to_utf16`, `utf16_to_scalar`, `byte_to_line`, `line_to_byte`, `line_to_scalar`, `line_to_utf16` |
| Positions | Checked `byte_to_position` / `position_to_byte` with byte, scalar or UTF-16 columns |
| Lines | `line_bounds`, shared `line` / `line_content` slices, `lines` / `line_contents` / `lines_from` iterators |
| Iteration | Lazy fused `chunks`, `bytes`, `scalars`, `chunks_from_byte`, `scalars_from` |
| Streaming and validation | Generic `write_to`, full AVL/summary/UTF-8 `validate`, `ToString`, `Debug`, content-based `PartialEq` / `Eq` |

`replace_bytes(start, end, replacement)` accepts a `Rope`, so a replacement may share existing text. Ranges are half-open. Slicing returns an independent immutable rope with the same API, sharing its interior subtrees. Scalar indexes count Unicode scalar values, not grapheme clusters.

All indexes use nonnegative `isize`. Boundary conversions accept EOF. Accessors `byte` and `scalar` require an actual element. UTF-8 interior bytes and UTF-16 surrogate interiors are rejected instead of rounded. `line_to_byte(len_lines())` is a one-past-line sentinel returning the byte length; `line(index)` accepts only actual lines.

## Newlines and positions

LF, CRLF and lone CR terminate lines. CRLF remains one terminator even when its bytes occupy different leaves or become adjacent through editing. Every rope has at least one line; a final terminator adds a final empty line. U+0085, U+2028 and U+2029 are ordinary scalars in this package.

`line()` includes its terminator; `line_content()` excludes it. `byte_to_line()` assigns the byte boundary between CR and LF to the preceding line. Positions represent line content: the position at the beginning of a terminator is the final content column, while the interior CRLF boundary is rejected. Positions beyond a line's content are rejected; protocol consumers such as LSP can implement their own clamping policy.

## Streaming and builder

`from_reader[R: std::io::Read](reader, max_bytes)` reads incrementally, preserves incomplete UTF-8 scalars across read boundaries, and reports absolute malformed-byte offsets. It retries interrupted reads, validates returned byte counts, rejects truncated UTF-8 at EOF, and enforces the caller's byte limit. At the limit it may consume one extra byte to distinguish EOF from oversized input. I/O errors retain the original `std::io::Error` in `Error::Io`.

`write_to[W: std::io::Write]` writes chunks through `write_all`, handling partial writes and interruptions. It does not flush or close the caller's stream. Errors may occur after partial input consumption or output delivery.

`Builder` appends validated strings or ropes and keeps a bounded partial leaf plus completed leaves. `finish()` creates a snapshot without consuming or clearing the builder; `clear()` does not alter previously finished ropes. Builder and iterator values contain mutable state: copying one shares that state, and concurrent access requires caller synchronization. Calling `rope.chunks()` or another iterator constructor twice creates independent cursors.

## Costs and limits

For `N` logical text bytes, with fixed 1024-byte maximum leaf size:

- Construction, full traversal, flattening, equality and compaction take `O(N)` time.
- Lengths and snapshot creation take `O(1)` time.
- Point queries, conversions, splitting and shared slicing take `O(log N)` time plus a bounded leaf scan/copy. Concatenation takes `O(|left height - right height| + 1)` plus bounded leaf work; insertion of `M` fresh bytes takes `O(M + log N)`.
- Chunk, byte and scalar iterators use `O(log N)` cursor space and visit text once. Line iterators produce shared slices using indexed line lookup, so traversal costs `O(number of lines × log N)` plus any requested text materialization.
- Builder appends copy at most a leaf-sized pending buffer per append. `finish()` takes time proportional to completed chunks; repeated calls rebuild branch nodes while sharing completed leaves.
- `validate()` walks all logical leaf occurrences and recomputes metadata. It is a diagnostic operation, not a constant-time property check.

The logical byte limit is `isize::MAX - 1` on the repository's Linux amd64 target, keeping the additional final-line count representable. Concatenation, replacement and builder append detect overflow before changing state, including enormous logical ropes made through shared doubling. Such ropes can be represented cheaply, but flattening, full traversal and validation still require resources proportional to their logical contents. Allocation exhaustion is not converted into `Error`.

This is not a drop-in Ropey port: there are no grapheme iterators, reverse iterators, Unicode newline feature flags, editing-history manager, search engine, or memory-mapped backing. A rope stores valid UTF-8 only. Tree nodes use GoML GC-managed private references rather than ownership or reference-counted copy-on-write. The API has no panic-unwind cleanup promise and does not add cancellation to arbitrary caller-provided blocking streams.

## Verification

From the repository root:

```sh
just ecosystem-test rope
```

The verifier checks formatting, 11 library black-box tests, separate versioned consumer tests, initial and cached consumer builds, the runnable consumer, native reference corpus checks, and Go's race detector over all library tests.

Native consumer tests retain every result from an independent flat-string / UTF-8 / UTF-16 / newline model: 67 deterministic scenarios, 3266 edits, and every final byte/scalar/UTF-16/line boundary. It covers malformed ranges, split surrogate and UTF-8 positions, CRLF at leaf seams, rotations, concatenation, compaction, builder fragments and snapshot preservation. Black-box tests additionally exercise AVL balance under repeated left/right joins, retained snapshots through edits, fragmented/invalid streaming I/O, parallel readers and editors, and checked logical-size overflow through shared doubling without huge allocations.
