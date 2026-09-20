# ecosystem::lsp

A reusable GoML toolkit for language servers and LSP peers: bounded streaming
frames, validated JSON-RPC messages, typed handler registration, request tracking,
document synchronization, UTF-16 coordinates and a synchronous server lifecycle.
The compiler's internal LSP implementation is not imported.

The wire profile targets [LSP 3.17](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/)
and [JSON-RPC 2.0](https://www.jsonrpc.org/specification). LSP uses one message per
frame; JSON-RPC batch arrays are rejected. The module does not claim complete
coverage of every optional LSP feature or generated protocol type.

## Layers and APIs

| Layer | Public entry points |
| --- | --- |
| Framing | `FrameLimits`, `FrameDecoder`, `frame_body`, `frame` |
| Messages | `Id`, `Message`, `RpcError`, `decode`, `decode_value`, `encode`, `encode_value` |
| Documents | `Position`, `Range`, `TextEdit`, `Change`, `Document`, `Documents`, `utf16_length` |
| Requests | `Session`, `RequestToken`, `Completed` |
| Dispatch | `RequestContext`, `Router`, `Server`, `Phase`, `Event` |
| Feature helpers | `DocumentPosition`, `Location`, `Diagnostic`, `publish_diagnostics`, `hover`, `progress` |
| Standard I/O | `Stdio::new`, `read_body`, `write` |

`../consumers/lsp` is a runnable server with hover, typed addition, echo and
document-inspection handlers. Its black-box tests also instantiate public APIs
across a normally resolved versioned dependency.

## Framing and messages

Call `FrameDecoder.push(bytes)` with arbitrary byte chunks, then repeatedly call
`next` until it returns `None`. Bodies are returned as checked UTF-8 strings.
Headers can span chunks, a chunk can contain several frames, and a Unicode scalar
can span any number of chunks. Header scanning resumes at the previous cursor;
consumed buffer prefixes are compacted when needed.

`FrameLimits::standard()` permits 8 KiB headers, 16 MiB bodies and 32 MiB plus
16 KiB of buffered input. Limits are configurable and validated. Header parsing
checks ASCII, field names, decimal byte lengths, duplicate length/type/charset
fields and UTF-8 charset declarations, including `utf8` and quoted values.
Header names are case insensitive. Unknown headers and non-charset content-type
parameters are ignored. Media type names themselves are not restricted.

Malformed framing poisons the decoder until `reset`; it never guesses a new
frame boundary. `finish` succeeds only after all complete frames have been drained
and no partial header/body remains. `buffered` exposes queued bytes and `needed`
helps blocking transports request the remaining body bytes. `Stdio` reads header
bytes incrementally, then reads the known body length in one operation. It keeps
protocol output separate from diagnostic stderr.

`decode` distinguishes parse errors (-32700) from invalid message shapes (-32600).
It rejects nesting deeper than 128 before invoking the standard JSON parser.
Top-level duplicate fields, invalid version strings, scalar/null params,
null request IDs and mixed request/response fields are rejected. A response must
contain exactly one of `result` or `error`; a null result is still a success.
Unknown extension fields are accepted. Integer IDs and error codes use checked
decimal integer parsing and the signed 32-bit LSP range. String and numeric IDs
remain distinct, including `"1"` versus `1`.

`encode` validates the message envelope; `encode_value` is an unchecked structural
builder. Raw JSON values supplied by callers must be valid JSON values, including
numeric lexemes, and must not contain reference cycles. Message parameters and
results preserve dynamic JSON so custom protocol methods do not require library
changes. Framing limits apply to received streams, not standalone `decode` calls
or outgoing messages.

## Documents and edits

`Document` is an immutable snapshot with URI, language ID, version, source string
and a standard-library line index. `position(byte_offset)` and `offset(position)`
convert checked UTF-8 boundaries and UTF-16 coordinates. LF, CRLF, lone CR and
trailing empty lines are supported. Columns beyond line content clamp to its end;
negative/out-of-range lines and positions within a surrogate pair are rejected.
Offsets inside line terminators map to the preceding line end.

`Document.changed(version, changes, max_bytes)` applies full and incremental
changes sequentially: each range addresses the result of the previous change.
Optional deprecated `rangeLength` is checked against the replaced UTF-16 length.
Versions must increase, but need not be consecutive. Invalid ranges, oversized
results and invalid versions return errors.

`Documents` stores open snapshots by exact URI string and enforces document-count
and per-document byte limits. `open` rejects duplicate opens, `change` commits
only after the entire batch succeeds, and `close` returns the removed snapshot.
Failures leave the stored text/version unchanged. `synchronize` decodes and
handles didOpen/didChange/didClose parameters; other methods return `false`.
URIs are opaque identifiers; file access and URI normalization belong to callers.

`Document.apply_edits` implements original-document `TextEdit[]` coordinates,
which differ from sequential didChange events. It validates every range, sorts
by original position, rejects overlap, and preserves insertion order for edits
at the same position. Inserts can precede one replacement at that position.
The result is returned as a string without changing the snapshot.

## Requests and dispatch

`Session.begin` returns a session-specific identity token for an incoming
request. Duplicate active IDs and configured pending limits are checked.
`cancel` marks a known token; unknown cancellation IDs are ignored.
`RequestToken.check` reports -32800 after cancellation. `complete` emits one
response and invalidates that token; tokens from other sessions, completed
tokens and stale tokens after ID reuse are rejected. Cancellation does not
automatically discard a successful or partial result: the handler decides its
response, and every accepted request must eventually be completed.

Outgoing `request` registration is separate from incoming requests, so each peer
can use the same ID independently. `receive` correlates replies and rejects
duplicates/unexpected IDs. `cancel_outgoing` retains the pending request while
creating a cancellation notification; `expire` removes it and creates that
notification. Applications schedule their own timeouts and send the returned
messages. Late responses after expiry are recoverable errors.

`Router.on_raw_request` registers a dynamic handler. `on_request[P, R]` uses
`std::serde::Deserialize` and `Serialize` for typed parameters/results, including
downstream derives and captured closures. Parameter conversion failures become
-32602; serialization failures become -32603. Handlers receive their token and
the document store. Notifications use a separate registry and never produce
wire responses. Duplicate handler registration is an error; unknown
notifications are ignored and missing request handlers return -32601.

`Server` combines these layers. Initialize parameters must contain an object
`capabilities`; the server supplies UTF-16 and incremental document-sync
capabilities plus caller-supplied capabilities. Its phases are New,
AwaitInitialized, Running, Shutdown and Exited. The initialized notification
enables normal request dispatch; earlier ordinary requests return -32002.
Shutdown returns null and subsequent requests fail. Exit reports status 0 after
shutdown and 1 otherwise, leaving actual process termination to the caller.
Caller registrations for built-in lifecycle or synchronization methods do not
override these built-in handlers.

`Event` returns a reply, a correlated client response or an exit code.
Notification/transport failures are returned to the application for local
diagnostics; they are not sent as unsolicited responses. The consumer demonstrates
parse-error responses with null IDs and continuing after a malformed JSON body.

## Execution and scope

All mutable objects use shared storage and are intended for one owning event
loop. `Server.handle` invokes callbacks synchronously; a cancellation message
cannot interrupt a callback running on that same loop. Applications needing
deferred work can use `Session` and `Router` directly, poll cancellation tokens,
and complete requests after yielding to their event loop. This library does not
provide a worker pool, timer scheduler or synchronization for concurrent mutation.

Document changes currently rebuild strings and line indexes, so a batch of C
changes on a document of size N can take O(CN) work and allocate intermediate
snapshots. The library provides explicit size limits, but does not use a rope.
Optional features such as completion/semantic-token generation, workspace-edit
execution, URI file loading and capability-specific language semantics belong to
the application. Raw JSON handlers support their protocol messages.

## Validation

Run from the repository root:

```sh
python3 ecosystem/verify.py lsp
```

The matrix formats/checks both modules, runs their tests, builds the consumer,
checks cached artifact stability, executes its smoke path and runs `interop.py`.
The protocol tests cover every two-part split of a Unicode frame, single-byte
chunks, multiple frames, limits, truncated input, malformed headers, reset,
malformed JSON, integer overflow, all message variants and request lifecycle.
Document tests cover transactional rollback, original-coordinate edits, Unicode
boundaries and source snapshot independence. Typed dispatch is tested in the
library and its separate consumer.

The independent Python client uses its own JSON framing and UTF-16 encoding
oracle. It checks 708 replies, 431 Unicode positions, 30 document-change
sequences, hover/typed dispatch, lifecycle ordering and termination behavior
through an actual GoML subprocess. This is targeted interoperability coverage,
not certification against the entire LSP specification.
