# MIME

Pure GoML MIME value helpers. `parse_media_type` and `format_media_type` handle
type/subtype values, quoted parameters, and UTF-8 RFC 2231 extended parameters
and continuations. Names are ASCII-case-normalized; values remain case-sensitive.
Duplicate names, continuation gaps, malformed percent escapes, unsupported
charsets and decoded control bytes are errors. Formatting emits UTF-8 extended
parameters when values contain non-ASCII characters.

`decode_word` and `decode_header` handle RFC 2047 B/Q encoded words. Decoding
accepts UTF-8, US-ASCII and ISO-8859-1 and produces UTF-8 text. Adjacent
encoded words discard intervening horizontal whitespace. Malformed encoded
words are rejected rather than silently retained. `encode_word` emits UTF-8
B/Q words split at scalar boundaries and limited to 75 bytes per word.

`Limits` bounds input, output, parameter count and value sizes. Inputs and
outputs are GoML strings, so raw invalid UTF-8 octets are outside this API.
The root-package helpers do not parse multipart bodies, transfer encodings, or
email addresses. Applications retain authority over header line folding,
transport limits and accepted media types.

The `quotedprintable` child package provides streaming transfer encoding and
strict decoding as a separate byte-oriented contract. The `multipart` child
package streams bounded part headers and binary bodies using textproto fields.
