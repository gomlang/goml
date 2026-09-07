# B4 resource example and release acceptance

Scope: the real file-resource example, boundary documentation, format/diagnostic/artifact coverage and packaged-toolchain compatibility required by plan.md B4. Deferred generic map-key obligations from B3 remain separately outstanding.

## Resource behavior

The committed [file-reading example](../../examples/ffi-file-read/README.md) binds actual os.Open, io.ReadFull and File.Read/Close. Its bundled three-byte input produces a positive count together with io.ErrUnexpectedEOF. Outcome retains both, and the Go shim receives the pair back and compares the original error object. Shared slice storage exposes Go writes; only the returned range is copied before checked UTF-8 decoding.

The pointer passes through a Go identity function before Close. Closing that alias makes a read through the original pointer fail. Close runs immediately after the bounded read, before decoding or reporting its outcome; its error remains separate. Missing files return on the open-error branch without reading or closing the nil handle. No adapter handle registry or finalizer-driven Close is involved.

`file_read_example_preserves_partial_results_and_closes_handles` runs the committed sources in isolation, including fresh/cached partial reads, full buffers, empty input, invalid UTF-8 and missing files. Invalid bytes still produce the original partial count and error and a successfully closed handle. The earlier method-binding regression additionally covers nil receiver behavior and an actual Reader returning positive bytes with EOF.

## Boundary contract and compiler coverage

The canonical language guide documents raw pointer/error/string/rune/slice/map representations and explicit conversion policies. Raw strings preserve arbitrary bytes; Unicode conversion is checked. Raw rune is i32, and scalar conversion is checked. Slice copy/share/trusted-readonly policies and independent headers are explicit. Map conversion does not imply HashMap equivalence or deep copy. Ordinary string/char library exports are rejected; raw string/rune exports require the author to select a supported failure result or deliberate boundary failure behavior.

External type declarations have parser recovery, AST attribute validation, HIR identity, inference, substitution and formatter tests. Pointer bridge persistence, foreign metadata binary round trips and raw map key artifact rejection have dedicated compiler tests. The pointer/error/raw-string/rune/slice/map driver regressions execute through separate package artifacts and cached rebuilds; invalid signatures, conversions and keys have diagnostics tests. Public API navigation is exercised by the corresponding query tests. The boundary aliases reuse ordinary type syntax; a combined formatter fixture covers their nested spelling and receiver attributes.

## Release and bootstrap

Extracted-release smoke copies the file example's sources, checks formatting and required FFI validation, runs it twice and compares generated-source digests. It uses the installed archive's executable-relative libraries without third-party dependencies. Main CI builds from the checksum-pinned released stage0, checks the stage2/stage3 fixed point, runs compiler/driver tests and verifies all golden output. Compiler and driver implementation sources do not adopt the new FFI boundary APIs before advancing stage0; fixture source strings may exercise them immediately.

## Verification status

Latest-source `just ci` passed with explicit exit status 0: 851 compiler tests and 129 driver tests, helper race checks, golden verification, stage3 fixed point and extracted-release smoke (/tmp/goml-b4-final-ci.log, /tmp/goml-b4-final-ci.status). This includes the combined formatter fixture and invalid-text file regression. B4 is accepted against the scope above. B3 deferred map-key requirements and batch D are not accepted by this audit.
