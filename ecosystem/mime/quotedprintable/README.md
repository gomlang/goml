# Quoted-printable

`Reader[R: Read]` decodes bounded quoted-printable data one physical line at a
time and exposes `read_chunk`. It accepts hard CRLF and `=\r\n` soft breaks,
hex escapes with either hex case, and final unterminated lines. It rejects bare
CR/LF, malformed escapes, raw non-ASCII/control bytes, trailing raw horizontal
whitespace and encoded lines longer than 76 characters. Errors are terminal.

`Writer[W: Write]` accepts arbitrary byte chunks with `write_chunk` and emits
ASCII quoted-printable. The default text mode preserves CRLF pairs as hard line
breaks even across chunk boundaries; lone CR or LF bytes are escaped. Binary
mode escapes every CR and LF byte. Raw spaces/tabs are delayed so trailing
whitespace is escaped. Soft breaks keep physical lines within 76 characters.
`finish` flushes the final line but does not close the sink.

`Limits` bounds input bytes, output bytes and physical lines. Read and write
provider failures are retained as causes. This package intentionally uses a
strict RFC 2045-oriented policy rather than Go's permissive broken-mail
recovery behavior.
