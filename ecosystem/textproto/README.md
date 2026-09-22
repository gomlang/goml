# textproto

Pure GoML framing for line-oriented text protocols. The package keeps protocol
bytes separate from character decoding so callers can apply HTTP or mail policy
without silently replacing non-UTF-8 octets.

`Reader[R: Read]` owns a buffered source. `read_line` returns one physical line
without its CRLF, or `None` only at clean EOF. Bare CR/LF and unterminated lines
fail. `read_headers(allow_folded)` reads a blank-line-terminated block, keeps
duplicate fields in order and canonicalizes ASCII names. Folded continuation
lines are rejected unless explicitly enabled; enabled folds join with one space.
Values preserve non-ASCII bytes, while `Field::value_text` requests UTF-8
validation. `parse_header_line` exposes the same single-field validation to
protocol adapters that own their own transport buffering.

`read_dot_chunk` incrementally removes dot stuffing and recognizes a line
containing only a dot as the terminator. `read_dot_bytes` is a bounded convenience
that materializes the decoded block. `Writer[W: Write]` writes CRLF lines and
header blocks, then `begin_dot`, `write_dot_chunk` and `finish_dot` frame a dot
block. Input chunks can split CRLF. A final unterminated data line gets CRLF
before the terminator. Writer I/O errors and reader parse/source errors are
terminal; invalid writer input is checked before emitting that chunk.

`Limits` independently bounds physical line content, total header wire bytes,
field count, unfolded field value bytes and decoded dot-body bytes. Those limits
are application policy, not SMTP's fixed transport line limit. Callers needing
SMTP's 1000-octet line rule should set `max_line_bytes` accordingly. Header
values reject control bytes other than horizontal tab but can contain arbitrary
high bytes; consumers must choose a character policy explicitly. The package
does not parse SMTP status codes, MIME encoded words, HTTP start lines or
multipart bodies.

Run `just ecosystem-test textproto` from the repository root.
