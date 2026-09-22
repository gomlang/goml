# archive

A pure GoML implementation of USTAR/PAX TAR and ZIP archives. Format
parsing, writing, CRC32, DEFLATE/GZIP compression, metadata validation, resource
budgets, and extraction are implemented in GoML using the standard library.
There are no Go adapters or native module dependencies.

## API

```goml
use ecosystem::archive;
use std::bytes;

let entries = Vec::from_array([
    archive::Entry::directory("project/"),
    archive::Entry::file("project/main.gom", bytes::Bytes::from_string("hello")),
]);
let limits = archive::Limits::standard();
let packed = archive::encode_zip(entries, archive::Compression::Deflate, limits)?;
let indexed = archive::ZipArchive::open(packed, limits)?;
let first_file = indexed.entry(1)?;
```

- `encode_tar`, `decode_tar`, `encode_zip`, `decode_zip`, `encode_tar_gz`, and
  `decode_tar_gz` provide bounded in-memory APIs.
- `TarWriter[W: Write]` writes entries immediately. `append_reader` copies an
  explicitly sized payload using an 8 KiB buffer; `finish` writes the two-block
  end marker. It does not retain entry bodies.
- `TarReader[R: Read]::next` reads one entry at a time, materializing at most one
  bounded entry body. The next call releases that body's storage when the caller
  no longer retains it. Global/local PAX headers are applied during iteration.
- `ZipWriter[W: Write]` writes local headers and compressed bodies immediately,
  retaining the central directory and one bounded compression buffer. Both
  `append` and `append_reader` are available; `finish(comment)` emits the central
  directory and archive comment.
- `ZipArchive::open` owns a snapshot of its input and validates the directory,
  local headers, descriptors and entry ranges. `read[R: Read]` reads a bounded
  archive into memory. `records` exposes metadata without decompression;
  `entry`, `entries`, and `copy_entry[W: Write]` validate expanded sizes and CRC32.
  `copy_entry` decompresses into bounded chunks without retaining the expanded
  entry; a late checksum or writer error may leave a partial output.
- `ZipSource[R: ReadAt]::open(reader, size, limits)` indexes bounded directory
  metadata without reading every entry body. `entry` reads one selected entry;
  `copy_entry` streams it through a bounded buffer. `open_zip_source` opens a
  regular file on Linux and returns a source that the caller must `close`.
  The caller must keep any supplied random-access source stable while indexed.
- `open_tar`, `open_tar_gz`, `open_zip`, `read_file`, and `write_file` supply file
  APIs. `write_file` creates a new file exclusively and never overwrites one.
- `crc32` and incremental `crc32_update` implement IEEE CRC32 in GoML.
- `gzip` and bounded `gunzip` implement RFC 1952, including optional headers,
  header checksums, and concatenated members. DEFLATE decoding supports stored,
  fixed Huffman, and dynamic Huffman blocks. Encoding uses LZ77 matching with
  fixed Huffman codes and falls back to stored blocks for incompressible input.

DEFLATE encoding and decoding are shared with `ecosystem::compress::flate`,
using a bounded 32 KiB history window. These archive convenience APIs still
materialize the returned entry/member data; direct incremental encoding and
decoding are available from the compress module. GZIP framing, including
concatenated members and optional headers, is shared there. ZLIB and LZW are
also available from the compress module.

Writers complete short writes through standard traits and propagate write errors,
including interruptions, without automatic replay. Partial
output failures poison writers, and TAR parsing failures poison readers. `finish`
and terminal TAR EOF are idempotent. Callers own and close the underlying stream.
Instances are intended for one sequential owner; independent instances can run
concurrently.

Generic reads also propagate interruptions immediately without retry. A failing
read may already have consumed bytes or modified the destination buffer without
reporting a count; those bytes do not become a successful archive read. TAR entry
input errors poison the writer because headers may already have been emitted.
ZIP `append_reader` reads the entire bounded entry before emitting it: an input
failure leaves that writer usable, but recovery requires a fresh or explicitly
repositioned entry source, not blindly resuming the failed one. TAR reader errors
are terminal; later calls return `Closed` without additional source reads.

## Metadata and limits

`Header` contains path, type, permissions, UID/GID, Unix modification seconds,
user/group names, link target, ZIP comment and unknown PAX key/value extensions.
USTAR prefixes, PAX records, GNU long-name/link records and checked GNU
base-256 numeric fields are decoded; writing uses PAX when strings or IDs do
not fit USTAR. TAR supports files, directories, symbolic links and hard links. ZIP
supports files, directories and Unix symbolic-link entries, Stored and Deflate,
UTF-8 names, extended Unix timestamps, archive/file comments, and both signed and
unsigned classic data descriptors. ZIP directory ranges are checked in sorted
order to reject overlapping entries without quadratic pairwise validation.

`Limits` bounds entries, individual expanded bytes, total expanded bytes,
archive bytes, path bytes, and metadata bytes. Decompression is bounded before
returning a payload, and declared ZIP sizes are checked before decompression.
Default limits are 100,000 entries, 64 MiB per entry, 256 MiB expanded total,
512 MiB archive, 4 KiB paths and 1 MiB metadata. Classic ZIP additionally limits
entry counts and offsets to its non-ZIP64 ranges. In TAR streaming APIs the
archive limit counts consumed/emitted bytes; `decode_tar` additionally verifies
that trailing bytes after the end marker are zero.

## Extraction policy

`extract(entries, existing_destination, ExtractOptions::standard())` validates
all archive paths and parent relationships before creating entries. It rejects
absolute paths, Windows drive/colon paths, backslashes, NUL, empty/`.`/`..`
components, duplicate normalized paths, and file/link parents. Symbolic and hard
links are rejected by default; `LinkPolicy::Skip` explicitly ignores them.
Existing files and symlinks are never overwritten. Directory components must be
real directories. Where available, Linux `openat2` resolves paths relative to
the opened destination with `RESOLVE_BENEATH | RESOLVE_NO_SYMLINKS`. If the kernel
returns `ENOSYS`, the extraction switches to an `openat` directory-descriptor
walk. Each intermediate component is opened with `O_DIRECTORY | O_NOFOLLOW`,
and the final component also uses `O_NOFOLLOW`; file creation uses
`O_CREAT | O_EXCL`. Absolute paths and `..` are rejected before either backend
performs an operation. Permission, symlink and other errors do not trigger a
fallback. All temporary descriptors are closed on success and failure.

Directory creation uses opened parent descriptors and verifies the resulting
root-relative directory. Replacing an ancestor or leaf with a symlink cannot
redirect a pending operation to that symlink's target: an operation either uses
its already opened directory or rejects the replacement. An opened directory
continues to identify the same object if renamed, including when moved outside
the destination. This is the directory-handle behavior of Go's `os.Root`, and
the fallback does not make an entire path walk atomic. The caller chooses and
creates the trusted destination and must prevent untrusted actors from moving
opened directories out of it when pathname containment is required. Neither
backend prohibits traversal of existing mount points. Extraction is not
transactional: a filesystem error can leave earlier completed files behind.

Permissions default to `0644` for files and `0755` for directories, subject to
umask. `preserve_permissions` restores only ordinary file permission bits;
setuid/setgid/sticky bits, ownership and timestamps are never applied. Directory
permissions are kept traversable while extracting.

## Deliberate format boundaries

- ZIP64 end records, entry extra fields and 64-bit descriptors can be read
  within limits. ZIP64 writing is automatic when classic fields overflow and
  may be forced through `ZipWriter::with_zip64`.
- Split/encrypted archives, non-Deflate compression, legacy non-ASCII
  codepage names, TAR devices/FIFOs and GNU sparse extensions produce
  recoverable errors.
- PAX numeric fields support nonnegative values; modification fractions are
  reduced to whole seconds. ZIP writing supports unsigned 32-bit Unix seconds;
  decoding ZIP entries without extended timestamps currently yields zero.
- `TarReader::next_header` and `read_body_chunk` stream entry bodies within
  limits; calling `next_header` again skips any unread body. `TarReader::next`,
  `decode_tar`, `ZipArchive` indexing and GZIP convenience APIs still materialize
  bounded data. `ZipSource` indexes only metadata; `ZipWriter::append_reader`
  buffers one bounded entry.
- File convenience APIs use the Linux standard file-descriptor API. Extraction
  also works on kernels without `openat2`, using `openat`, `mkdirat` and
  no-follow directory descriptors. Format and standard `Read`/`Write` APIs do
  not assume seekable streams.

## Validation

Run from this module:

```sh
../../stage2/bin/goml fmt --check
../../stage2/bin/goml check
../../stage2/bin/goml test
GOFLAGS=-race ../../stage2/bin/goml test --target-dir _artifact/race
```

Native GoML tests cover Unicode/PAX metadata, CRC vectors, GZIP corruption,
truncation boundaries, invalid ZIP directory/local records, data descriptors,
short/interrupted I/O, resource limits, input snapshot isolation, path/link
attacks and real extraction. Internal tests force the `openat` fallback, replace
ancestors and leaves with symlinks after opening the parent, document the
rename-outside boundary, and check descriptor cleanup across repeated failures.
Interoperability tests create archives with GNU
`tar` and Info-ZIP `zip`, read them in GoML, and have `tar`/`unzip` verify GoML
output. The codec tests also exchange fixed, dynamic and stored DEFLATE streams
with `gzip`, exercise concatenated members and optional headers, and reject
malformed trees and truncated or over-budget data. These four programs must be
installed. There are no Python helpers.
The independent `../consumers/archive` module imports version `0.1.0` through the
isolated verification registry.

Format references: [PKWARE ZIP APPNOTE](https://pkware.cachefly.net/webdocs/casestudies/APPNOTE.TXT),
[GNU TAR USTAR description](https://www.gnu.org/software/tar/manual/html_node/Standard.html),
[POSIX pax](https://pubs.opengroup.org/onlinepubs/9699919799/utilities/pax.html),
[DEFLATE RFC 1951](https://www.rfc-editor.org/rfc/rfc1951),
[GZIP RFC 1952](https://www.rfc-editor.org/rfc/rfc1952), and
[Linux openat2](https://man7.org/linux/man-pages/man2/openat2.2.html).
The fallback follows the directory-descriptor approach used by
[Go 1.26 os.Root](https://cs.opensource.google/go/go/+/refs/tags/go1.26.0:src/os/root_openat.go),
with archive symlinks and parent traversal rejected outright.
