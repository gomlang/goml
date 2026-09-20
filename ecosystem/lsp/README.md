# ecosystem::lsp

A reusable GoML toolkit for language servers and LSP peers: bounded streaming
frames, validated JSON-RPC messages, typed handler registration, request tracking,
document synchronization, negotiated UTF-8/UTF-16/UTF-32 coordinates and synchronous
or deferred dispatch.
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
| Documents | `PositionEncoding`, `Position`, `Range`, `TextEdit`, `Change`, `Document`, `Documents`, `utf16_length` |
| Requests | `Session`, `RequestToken`, `Completed`, `ExpiredRequest` |
| Dispatch | `RequestContext`, `Router`, `Server`, `Phase`, `Event`, `Dispatch`, `PendingRequest` |
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
and a persistent rope. `Document::new` and `Documents::new` default to UTF-16;
`new_with_encoding` selects `PositionEncoding::Utf8`, `Utf16` or `Utf32` explicitly.
`encoding()` reports that selection. `position(byte_offset)` and `offset(position)`
convert checked UTF-8 boundaries to and from the snapshot's selected coordinate
units: UTF-8 bytes, UTF-16 code units or UTF-32 Unicode scalars. LF, CRLF, lone CR
and trailing empty lines are supported. Columns beyond line content clamp to its
end; negative/out-of-range lines and positions inside a UTF-8 scalar or UTF-16
surrogate pair are rejected. Offsets inside line terminators map to the preceding
line end. `PositionEncoding.length(text)` measures text in the corresponding
units; `utf16_length` remains available for explicitly UTF-16 computations.

`Document.changed(version, changes, max_bytes)` applies full and incremental
changes sequentially: each range addresses the result of the previous change.
Optional deprecated `rangeLength` is checked against the replaced length in the
snapshot encoding, including intervening line terminators.
Versions must increase, but need not be consecutive. Invalid ranges, oversized
results and invalid versions return errors.

`Documents` stores open snapshots by exact URI string and enforces document-count
and per-document byte limits. `open` rejects duplicate opens, `change` commits
only after the entire batch succeeds, and `close` returns the removed snapshot.
Failures leave the stored text/version unchanged. `open` adopts the store encoding
for its stored snapshot without changing the supplied snapshot. `synchronize` decodes and
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
`Session.close` invalidates all incoming and outgoing tokens and prevents new
registrations; it is idempotent. Server exit closes its session automatically.

Outgoing `request` registration is separate from incoming requests, so each peer
can use the same ID independently. `receive` correlates replies and rejects
duplicates/unexpected IDs. `cancel_outgoing` retains the pending request while
creating a cancellation notification; `expire` removes it and creates that
notification. `request_with_timeout(id, method, params, Some(duration))` additionally
registers a monotonic deadline starting when registration succeeds, including any
later transport delay. `None` preserves the untimed `request` behavior. Deadline
storage is bounded by the outgoing pending limit, with one deadline per request.
Invalid messages, duplicate IDs and capacity failures leave registration unchanged.
Cancellation retains the deadline; manual `expire`, receipt and `close` remove it.
Late responses after expiry are recoverable errors.

`Session.next_outgoing_timeout()` returns the shortest remaining duration, or
`None` without timed outgoing work. `poll_outgoing_timeouts()` removes each due
request exactly once and returns `ExpiredRequest { completed, cancellation }`.
The local completion carries the original ID/method and RequestFailed (-32803);
the cancellation is a `$/cancelRequest` notification to send to the peer. If a
response is received after its deadline but before polling, `receive` instead
returns that local timeout completion and consumes the response; no cancellation
is needed for a reply already received. An earlier response removes its deadline
and wins normally, including a successful result after cooperative cancellation.

No request IDs are reserved permanently. Applications must not reuse an expired
outgoing ID while a reply to the old request might still arrive: JSON-RPC responses
carry only IDs, so such replies cannot be distinguished from replies to reused IDs.
Incoming token identity protection remains independent of this wire limitation.

`Router.on_raw_request` registers a dynamic handler. `on_request[P, R]` uses
`std::serde::Deserialize` and `Serialize` for typed parameters/results, including
downstream derives and captured closures. Parameter conversion failures become
-32602; serialization failures become -32603. Handlers receive their token and
the document store. Notifications use a separate registry and never produce
wire responses. Duplicate handler registration is an error; unknown
notifications are ignored and missing request handlers return -32601.

`Server` combines these layers. Initialize parameters must contain an object
`capabilities`; the server supplies negotiated position encoding and incremental
document-sync capabilities plus caller-supplied capabilities. It selects the first
supported entry in `capabilities.general.positionEncodings`. Unknown strings are
ignored; absent, empty or unknown-only lists fall back to mandatory UTF-16.
Malformed lists or non-string entries fail initialization without changing phase
or document encoding, allowing a valid retry. This follows the
[LSP 3.17 negotiation contract](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#clientCapabilities).
`Server::new` supports all three encodings. `Server::new_with_encodings` restricts
that set and requires UTF-16; it copies the supplied vector. The negotiated value
is returned as `capabilities.positionEncoding` and exposed by
`Server.position_encoding()`. Initialization updates the store and its existing
snapshots atomically; previously captured snapshots keep their own encoding.
Subsequent synchronization, ranges, positions and edits use the same encoding.
Its phases are New,
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

### Deferred work and deadlines

`Server.accept(message)` returns `Dispatch { event, request }`. Built-in
lifecycle methods, document updates, cancellation and client replies produce an
immediate event. A registered ordinary request produces a `PendingRequest`
without invoking its handler. Duplicate IDs, pending limits and lifecycle rules
still apply. The event loop can now process further messages before finishing
that request.

Call `Server.execute(pending)` to invoke its registered raw/typed handler and
produce the final response, or `Server.complete(pending, result)` when the
application has computed the result separately. Both reject foreign, completed
and stale tokens before running user code. `pending.token()` supports cooperative
cancellation and `params()` exposes the request parameters. Document snapshots
should be captured before scheduling work when it needs a fixed source version.
`execute` checks cancellation before invoking a handler; `complete` permits a
successful partial result after client cancellation, as the protocol allows.

`accept_with_timeout(message, Some(duration))` adds a total deadline for deferred
work. `next_timeout()` gives the shortest remaining duration across incoming and
outgoing deadlines for an event-loop timer. When it fires, call `poll_timeouts()`
and send every returned incoming response; also call `poll_outgoing_timeouts()`,
deliver each local completion and send each cancellation notification. Both polls
are non-blocking and require the same owning event loop as normal dispatch.
Recompute `next_timeout()` after registration, response receipt, expiry and either
poll; `None` means no timer is needed. A zero duration is immediately due.
Expired requests finish once with RequestFailed (-32803); late completions are
rejected, including after an ID is reused. `complete` and `execute` also check
deadlines, so an overdue result cannot bypass a delayed timer poll. No timer or
worker goroutine is created. The application must arrange a timer wakeup even when
no input arrives; polling only after reading messages cannot enforce prompt idle
timeouts. The stdio consumer uses an explicit `test/poll` notification for its
protocol deadline fixtures. Built-in immediate methods do not receive deadlines.

Shutdown stops accepting ordinary work while allowing already accepted requests
to finish. Exit invalidates remaining requests, clears deadline tracking and
ends the protocol stream without generating additional responses.

## Execution and scope

All mutable objects use shared storage and are intended for one owning event
loop. `Server.handle` invokes callbacks synchronously; a cancellation message
cannot interrupt a callback running on that same loop. Applications needing
deferred work can use `accept`, yield to their event loop, and call `execute` or
`complete` later. Handler callbacks themselves remain synchronous and cannot be
preempted. All server/session/token access must remain on the owning loop; worker
tasks should receive immutable inputs and return results to that loop. This library does not
provide a worker pool, timer scheduler or synchronization for concurrent mutation.

Documents use `ecosystem::rope` persistent text storage. Edits share unchanged
tree nodes with previous snapshots, and encoding conversions use cached subtree
byte, UTF-16 and scalar counts. `content()` and `apply_edits()` explicitly materialize their string
results; incremental document changes retain the tree representation. Existing
CRLF clamping, surrogate rejection, size limits and transactional updates remain
part of the document API.
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

The library runs 19 tests and its separate consumer runs 3. The independent Python
client uses its own JSON framing and Unicode encoding oracle. The original
UTF-16 suite checks 708 replies, 431 Unicode positions and 30 document-change
sequences. The negotiation suite adds 4,768 UTF-8/16/32 positions and invalid
boundaries, 48 transactional change sequences, CRLF split across rope chunks,
unknown encodings and malformed-initialize retries. Its outgoing request suite
checks actual cancellation/completion messages, early and overdue responses,
explicit timer polling, duplicate suppression and exit cleanup through a GoML
subprocess. This is targeted interoperability coverage, not certification against
the entire LSP specification.
