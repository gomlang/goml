# Multipart

`Writer[W: Write]` streams MIME multipart bodies with a caller-supplied
validated boundary. `begin_part(textproto::Headers)` serializes one header
block, `write_body_chunk` emits arbitrary binary data, and `finish` writes the
closing delimiter without closing the sink. Body data that would form a
delimiter is rejected, including matches split across input chunks.

`Reader[R: Read]` accepts a bounded preamble, exposes part headers through
`next_part`, and streams each body through `read_body_chunk`. A body must be
drained until it returns zero before advancing to the next part. Header
continuations join with one space; malformed framing and source failures are
terminal. The closing delimiter may end with CRLF or at EOF. The reader leaves
any epilogue unread.

`Limits` bounds wire bytes, preamble, headers, line and value bytes, part bytes,
aggregate body bytes, and part count. The package preserves duplicate headers
and byte-valued metadata; applications choose their own form-data field and
filename policy. It does not perform filesystem extraction or transfer-codec
decoding.
