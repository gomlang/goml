# Shared compression codecs

This module owns the E1 compression engines consumed by archive and future
protocol/image adapters. Production algorithms are GoML; Go compression codecs
are used only by differential test oracles.

Raw DEFLATE, GZIP, ZLIB and LZW encoding and decoding are incremental. These
complete the four mapped E1 compression capabilities, not the whole E1
migration.

## Incremental DEFLATE

`flate::Decoder[R: io::Read]::new(reader, compress::Limits)` accepts explicit
nonnegative input-byte, output-byte, block and work budgets.
`with_dictionary(reader, limits, dictionary)` seeds history from an independent
copy of the dictionary's final 32 KiB. A raw stream carries no dictionary ID;
the caller supplies the agreed dictionary.

Stored, fixed-Huffman and dynamic-Huffman blocks share a 32 KiB ring window.
The decoder does not retain complete input or output, and overlapping matches
and matches across block/window boundaries are supported. Dictionaries, trees
and transport scratch use bounded storage independent of stream length.

`read_chunk(MutSlice[byte]) -> Result[isize, compress::Error]` supplies expanded
bytes. Decoder also implements standard Read, mapping detailed errors to
io::Error. Successful partial progress is returned before a pending failure:
if an operation produced bytes before detecting an error, it returns their
count, exposes the failure through `error()`, and the next read returns that
failure without more I/O. Callers must read through EOF to validate the full
stream. A zero-length read does not advance the stream or prove EOF.

`consumed()` counts confirmed compressed bytes; `produced()` counts expanded
bytes handed to callers. `is_finished()` requires a validated final block with
no pending error. Final-block padding bits are discarded without reading a byte
from the enclosing format's trailer. Input is requested one byte at a time;
callers may compose a shared BufReader for transport buffering and must read
subsequent framing through that same buffered reader.

Limits are lifetime totals. Work counts decoder transitions and bit-read
operations, not wall-clock time. An exhausted input budget cannot issue an
extra EOF probe. Parser/provider failures are sticky across aliases.
Reentrancy is rejected, and unwinding poisons the decoder; provider panics still
propagate. Invalid provider counts fail, interruptions are not retried, and
providers are never closed. Concurrent use of aliases is unsupported.
Independent decoder instances can be used concurrently.

`flate::decode(Bytes, limits)` is the explicitly materializing single-stream
convenience. It rejects trailing bytes. Structured errors distinguish malformed
input, truncation, limits, I/O and invalid session state, with compressed byte
offsets and retained provider causes.

`flate::Encoder[W: io::Write]::new(writer, limits, level)` accepts levels -2
(Huffman-only), -1 (default), and 0 through 9. `with_dictionary` copies the
final 32 KiB of a preset dictionary. `write_chunk` accepts bounded input in
32 KiB blocks; `flush` emits a nonfinal synchronization marker; `finish` emits
the final block and is idempotent after success. Levels 1 through 9 vary LZ77
search effort. The encoder uses fixed Huffman codes or stored blocks when they
are smaller; level 0 uses stored blocks. Dynamic Huffman encoding is not
provided. Encoding does not retain the whole stream.

`accepted()` counts accepted source bytes and `written()` counts bytes confirmed
by the underlying writer. A failure after accepting a prefix returns that
prefix once and becomes sticky; a failure after a partial output write retains
the confirmed byte count. Input, output, block and work limits are lifetime
totals. Invalid/zero writer counts, writer errors and reentrant callbacks poison
the encoder; unwinding also poisons it. The underlying writer remains caller
owned and is never closed. `flate::encode(Bytes, limits, level)` is the bounded
materializing convenience.

## Incremental GZIP

`gzip::Writer[W: io::Write]::new(writer, limits, level)` emits an RFC 1952 header,
streams raw DEFLATE blocks and finishes with CRC32 and modulo-2³² input length.
It exposes `write_chunk`, `flush`, `finish`, `accepted`, `written`, `error` and
`is_finished` with the same terminal and partial-progress rules as the DEFLATE
encoder. The output budget includes the header and trailer; 8 bytes are reserved
for the trailer before DEFLATE output is emitted. The writer does not close its
underlying stream.

`gzip::Reader[R: io::Read]::new(reader, limits)` processes concatenated members,
including optional extra data, name, comment and header-CRC fields. It validates
each member's trailer before reporting final EOF. `read_chunk`, `consumed`,
`produced`, `members`, `error` and `is_finished` expose stream state. The reader
does not retain the entire input or expanded output; header fields are validated
without materializing them. A caller must read through EOF to validate every
trailer and reject trailing garbage. A generic reader needs one allowed input
byte to probe EOF after the final member; `gzip::decode(Bytes, limits)` accepts
an exact input-byte limit by making that probe internally on its known-length
cursor. A zero-length read does not advance the stream.

GZIP uses the same input, output, DEFLATE-block and work budgets as raw DEFLATE.
Framing bytes also consume input and work budget. Provider errors and invalid
counts are terminal, aliases share failure state, and reentrancy or unwinding
poisons the stream. Concurrent alias access is unsupported.

## Incremental ZLIB

`zlib::Writer[W: io::Write]::new(writer, limits, level)` and
`with_dictionary(writer, limits, level, dictionary)` emit RFC 1950 CMF/FLG,
optional full-dictionary Adler32 identifier, raw DEFLATE and an Adler32 trailer
of uncompressed data. `write_chunk`, `flush` and `finish` follow the DEFLATE
writer's bounded, sticky-state and partial-progress rules. The output budget
includes framing, with four trailer bytes reserved before compressed output.

`zlib::Reader[R: io::Read]::new(reader, limits)` and `with_dictionary` validate
compression method, advertised window size, FCHECK, dictionary identifier and
the final Adler32. A dictionary is copied only as its final 32 KiB for history,
but the identifier is computed from the complete dictionary. The reader leaves
following bytes in its source untouched. `zlib::decode` and
`decode_with_dictionary` are strict whole-buffer conveniences that reject
trailing bytes; streaming callers must read through EOF to validate the trailer.
The reader and writer implement standard I/O traits and expose consumed or
accepted byte counts, confirmed output, terminal state and sticky errors.

## Incremental LZW

`lzw::Reader[R: io::Read]::new(reader, limits, order, literal_width)` and
`lzw::Writer[W: io::Write]::new(writer, limits, order, literal_width)` implement
the GIF/PDF-style variable-width code stream used by Go's `compress/lzw`.
`BitOrder::Lsb` packs least-significant bits first (GIF); `BitOrder::Msb` packs
most-significant bits first. Literal width must be 2 through 8 bits, input
symbols must fit that width, and code width grows to at most 12 bits. The first
two nonliteral codes are Clear and EOF. The writer emits an initial Clear and
resets the dictionary on saturation; the reader also accepts valid streams
without an initial Clear. TIFF's incompatible early-change variant is not
provided.

`write_chunk` accepts a confirmed input prefix, `finish` emits the pending code
and EOF without closing the underlying writer, and `read_chunk` yields bounded
expansions without retaining complete input/output. The block limit counts
dictionary epochs begun by a Clear code or an initial data code. Input, output
and work limits, invalid codes and prefix chains, provider counts/errors,
reentrancy and panic cleanup are terminal and sticky. Materializing `encode`
and `decode` conveniences are bounded; `decode` rejects trailing bytes, while
the streaming reader leaves following framing unread.

## Integration and validation

Archive's DEFLATE/GZIP/ZIP convenience APIs are being migrated through separate
compatibility adapters in the archive package. The independent
`../consumers/compress` module checks encoding, GZIP composition, Read
composition and preservation of following framing through the isolated
verification registry.

Run `just ecosystem-test compress archive` from the repository root.
Tests compare Go flate decoding at all supported encoder levels and Go flate
encoding at representative decoder levels, stored/dynamic/fixed blocks,
dictionaries larger than the window, overlapping/window-spanning matches,
input/output chunk sizes, truncated/corrupt input, all quotas, provider faults,
aliases and panic cleanup. GZIP tests cover Go encoder/decoder interoperability,
representative compression levels, concatenated members, optional header
fields and header/payload CRCs, truncated inputs, quotas and provider faults.
ZLIB tests cover Go encoder/decoder interoperability, preset dictionaries larger
than 32 KiB, FCHECK/DICTID/Adler32 failures, smaller declared windows, trailing
data, quotas and provider faults.
LZW tests compare both directions with Go across both bit orders, literal
widths 2, 4 and 8, long streams, dictionary resets, malformed/truncated codes,
short writes, quotas, provider faults, aliases and panic cleanup.

Format reference: [RFC 1951](https://www.rfc-editor.org/rfc/rfc1951).
GZIP reference: [RFC 1952](https://www.rfc-editor.org/rfc/rfc1952).
ZLIB reference: [RFC 1950](https://www.rfc-editor.org/rfc/rfc1950).
