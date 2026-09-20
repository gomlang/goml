# bitflags

Typed integer flag sets in GoML, modeled on
[Rust bitflags 2.13.2](https://docs.rs/bitflags/2.13.2/bitflags/). The library is
implemented entirely in GoML and uses an ordinary third-party derive handler.
Each flags type remains a distinct user-defined struct with its original integer
storage; it is not a string-keyed map or a shared mutable handle.

```goml
use ecosystem::bitflags;
use bitflags::Flags;

#[derive(Flags, PartialEq, Eq, Hash, Default)]
#[flags(NONE = "0", READ = "0b1", WRITE = "0b10", READ_WRITE = "READ | WRITE", unnamed = "0x80")]
pub struct Permissions {
    bits: u8,
}

pub const READ: Permissions = Permissions { bits: 1 };

fn example() -> Result[Permissions, bitflags::ParseError] {
    let write: Permissions = bitflags::parse("WRITE")?;
    let combined = READ.union(write);
    let mut current = combined;
    current = current.set(READ, false);
    bitflags::parse(current.format())
}
```

The manifest dependency is `"ecosystem::bitflags" = "0.1.0"`. Repository
verification supplies an isolated local registry; this coordinate is not a claim
that the module has been published to a remote registry.

## Declarations

`#[derive(bitflags::Flags)]` accepts a non-generic struct with exactly one field
of type `u8`, `u16`, `u32`, `u64`, `usize`, `i8`, `i16`, `i32`, `i64`, `isize`, or
`byte`. The field can have any name and stay private. Other derives such as
`PartialEq`, `Eq`, `Hash`, and `Default` compose normally. There is no dependency
on a Rust toolchain at application build or runtime.

Use exactly one `#[flags(...)]` attribute. Empty declarations are valid. Each
named argument is a flag name and a string mask expression. Expressions support
decimal, `0x` hexadecimal, `0b` binary, `0o` octal, `~0` for all storage bits, and
`|` combinations of those values or previously declared flag names. Forward
references, empty components, overflow, duplicate names, and invalid storage
types are compile-time diagnostics. Up to 256 declarations are accepted, subject
to the compiler's existing derive evaluation budget.

The reserved argument `unnamed = "mask"` expands the set of known bits without
creating a name. It may occur more than once. Different names may have identical
masks. Zero and multi-bit flags are supported. `unnamed` replaces the Rust macro's
`const _` spelling; the GoML attribute metadata does not represent `_` as a named
argument. Flag names are case-sensitive.

The derive supplies the `Flags` trait implementation. It cannot generate
associated constants or inherent methods under GoML's current derive contract.
Declare ordinary module constants when needed, or use `Permissions::from_name`.
Import `Flags` in each file using its methods. GoML does not overload bitwise
operators for structs; use the methods below.

## Bit operations

All raw bit APIs use normalized unsigned `u64` values. `bits()` masks sign
extension for signed storage. `from_bits_retain` keeps all bits fitting the
storage width and discards bits above it; `from_bits` rejects unknown bits and
out-of-width input; `from_bits_truncate` keeps only known bits.

| API | Behavior |
| --- | --- |
| `empty`, `all`, `all_named`, `from_name` | Construct zero, all known bits, all named bits, or one named flag |
| `bits`, `storage_mask`, `known_mask`, `named_mask` | Inspect the stored pattern and definition masks |
| `known_bits`, `unknown_bits`, `contains_unknown_bits` | Inspect a value's known and unknown portions |
| `is_empty`, `is_all`, `contains`, `intersects` | Set predicates; `is_all` permits additional unknown bits |
| `union`, `intersection`, `difference`, `symmetric_difference` | Set algebra preserving unknown bits where the operation preserves them |
| `complement` | Invert within the known mask |
| `insert`, `remove`, `toggle`, `set`, `truncate`, `clear` | Return an updated value; assign it back to change a local variable |
| `definitions`, `type_name`, `debug_flags` | Metadata and readable diagnostics |

`difference` removes the other value's actual bits, including unknown bits; it
is deliberately different from intersecting with the other value's truncated
complement. Every value contains a zero flag, but no value intersects one.

## Iteration and text

`iter_names()` returns named flags in declaration order. A name is yielded when
its complete mask is contained in the original value and it covers at least one
remaining bit. This handles overlapping flags while skipping redundant aliases
and zero-bit flags. `IterNames::remaining()` includes unnamed, unknown, or partial
multi-bit masks not consumed by names. `iter()` appends this remainder as one
last value; `bitflags::collect(iterator)` reconstructs the original bit pattern.
Both iterators are fused. Copies share the current cursor and should be consumed
serially; creating another iterator starts a fresh traversal.

`Type::iter_defined_names()` includes every named definition, including aliases
and zero. `iter_equal_names()` includes every name whose mask exactly equals the
value. Both return ordinary `FnIterator` values.

The text form is `READ | WRITE | 0x80`; zero formats as the empty string. Parsing
accepts Unicode whitespace around terms, exact names, and lowercase `0x` prefixes
with case-insensitive hex digits. Repeated terms are allowed. Empty terms,
unknown names, malformed numbers, and storage overflow return a `ParseError`
containing its kind, token, and UTF-8 byte offset.

| Functions | Policy |
| --- | --- |
| `parse`, `format` | Preserve every in-width bit and round-trip |
| `parse_truncate`, `format_truncate` | Discard unknown bits, retaining known unnamed/partial bits |
| `parse_strict`, `format_strict` | Parse only names; format only contained named flags |
| `parse_checked` | Accept names and hex but reject unknown bits |

Hex parsing uses the unsigned storage pattern even for signed fields, preserving
round-trips for sign bits. This intentionally differs from Rust's signed integer
hex parser. There are no integer underscores, negative numbers, arbitrary
expressions, or comments in runtime text. Declaration expressions and runtime
text have different grammars.

## Serde

`Text[F] { value }` implements `std::serde::{Serialize, Deserialize}` using the
lossless text form. `Number[F] { value }` uses the unsigned 8/16/32/64-bit Serde
event matching the storage width. Both preserve unknown in-width bits; numeric
decoding rejects overflow. Signed storage also uses its unsigned pattern.
JSON and Bincode round-trips are tested at every width. The wrappers make the
wire format explicit because GoML's Serde protocol has no human-readable format
query. Ordinary Serde derives on the flags struct still serialize its field as
an ordinary struct.

Manual `Flags` implementations are supported: implement `bits`,
`from_bits_retain`, `storage_mask`, `definitions`, and `type_name`. The default
known/named masks are computed from definitions. Implementations must keep masks
within their storage width and supply unique, parseable names with stable
metadata. The derive enforces these conditions and emits constant mask methods.

## Verification

```sh
python3 ecosystem/verify.py bitflags
```

This runs public API tests, a separate versioned consumer, fresh/cached build
checks, Rust reference comparisons, and compile-time diagnostic checks. Tests
cover all 65,536 pairs of 8-bit operands, iterator reconstruction, aliases,
overlaps, empty and unnamed definitions, signed/full-width storage, Unicode
whitespace, malformed input, and explicit text/numeric Serde representations.

`interop.py` needs `rustc` and downloads the checksum-pinned bitflags 2.13.2 crate
into ignored `_artifact/`. It compares seven definition profiles with the real
Rust implementation across operations, iteration, formatting, and parsing;
no third-party Rust code is vendored or linked into the GoML library.
