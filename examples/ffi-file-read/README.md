# File reading across the Go boundary

From the repository root, build the toolchain and run this example:

```sh
just make
cd examples/ffi-file-read
../../stage2/bin/goml run
```

The example uses only the Go standard library and the local `shim` package. Run it from this directory so `data.txt` resolves to the bundled three-byte file. With an installed toolchain, use `goml run` here.

```text
read: 3
partial data: abc
Go round-trip preserved partial read: true
read error: unexpected EOF
explicit close succeeded: true
original handle is closed: true
```

`os.Open` receives an explicitly converted raw Go string. Its returned `*os.File` passes through Go and back unchanged. `io.ReadFull` requests four bytes from a three-byte file, so it returns three useful bytes together with `io.ErrUnexpectedEOF`. `ffi::Outcome` keeps both results. The Go shim receives those results back and checks the original error identity and count.

Go writes into a shared slice of the GoML buffer. The example copies only the returned byte range before checking UTF-8. Invalid text is reported without replacing bytes; the `Bytes` value still contains the original data. Converting the read outcome into a Result that discards the value on error would lose the partial read.

The explicit `Close` occurs immediately after the read, before interpreting its result. Closing the pointer returned by `EchoFile` also closes the original handle, as the last read demonstrates. The close error is retained separately from the read error. No adapter creates a handle registry or schedules automatic Close. The example does not depend on Go's finalizers running.

This is a bounded read demonstration, not a complete file-loading API. A longer file fills the four-byte buffer successfully; an empty file reports EOF with zero bytes. A missing file reports the open error and returns without attempting to read or close a nil pointer. The `IsPartialRead` shim intentionally recognizes only the bundled three-byte case.

The driver regression runs these committed sources in an isolated workspace, including full-buffer, empty-file, invalid UTF-8 and missing-file cases. Invalid text still preserves the partial count and original read error, and the handle is closed before decoding. Release smoke tests copy only the source files and run the example twice with an extracted toolchain, checking fresh and cached compilation and generated-source digests without third-party network access.
