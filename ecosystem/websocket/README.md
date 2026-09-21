# websocket

An RFC 6455 WebSocket protocol implementation in GoML. The protocol, HTTP
upgrade validation, SHA-1 accept calculation, framing, masking, message assembly,
close state machine and bounded queues are implemented in `.gom` files. The
standard library supplies cryptographic randomness, Base64, UTF-8, contexts and
TCP/TLS streams; this module contains no Go adapter or Python helper.

```toml
[dependencies]
"ecosystem::websocket" = "0.1.0"
```

The reference protocol is [RFC 6455](https://www.rfc-editor.org/rfc/rfc6455),
particularly sections 4–8. [tungstenite](https://docs.rs/tungstenite/latest/tungstenite/)
informs the separation of transport-independent protocol state and stream I/O.

## Opening a connection

```gom
use ecosystem::websocket as ws;
use std::context;
use std::net;

fn connect(stream: net::TcpStream, ctx: context::Context) -> Result[ws::Connection, ws::Error] {
    let request = ws::ClientRequest::new("localhost:8080", "/events", Vec::from_array(["events.v1"]))?;
    let (socket, _) = ws::client(ws::Transport::tcp(stream), request, ws::Limits::standard(), ctx)?;
    socket.send(ws::Message::Text("hello"), ctx)?;
    Result::Ok(socket)
}
```

`client` performs the HTTP/1.1 handshake and returns the selected subprotocol.
`ClientRequest::new` generates a cryptographically random 16-byte nonce;
`with_key` supports explicit nonce vectors. `with_origin` adds an Origin header.
The client verifies upgrade tokens, the accept digest and subprotocol selection,
and rejects extensions it did not negotiate. Requests and responses reject
ambiguous critical headers, body framing, line folding and injected control bytes.

`server(transport, limits, context, select_protocol)` parses the request, invokes
an application callback with target, host, origin and offered protocols, and
sends the response. The callback can enforce route/authentication/origin policy
and return `Err(Error { kind: ErrorKind::Handshake, message: ... })` to reject.
Returning `Some(protocol)` selects one offered protocol; `None` selects none.
Rejected stream handshakes close their transport. Applications that need custom
HTTP error responses can use `ServerRequest::parse` and `response` directly.

`Connection::upgraded` attaches a stream upgraded by another HTTP implementation
and accepts bytes already read after the headers. `HandshakeDecoder::push`
returns the exact consumed byte count, allowing callers to preserve coalesced
frame bytes. `ClientRequest::validate_response` and `ServerRequest::response`
are also available without network I/O.

## Messages and transport

`Message` has `Text`, `Binary`, `Ping`, `Pong` and `Close` variants. `receive`
returns control messages as well as data; incoming pings automatically queue
and flush matching pongs. `close` sends a closing frame. Continue receiving
until the peer's close arrives; the close response is flushed before returning,
and the connection then closes its transport. `abort` immediately closes and
wakes pending operations; it is idempotent. A raw EOF before the closing
handshake is an error. The caller owns the lifetime and should defer `abort`.

`Transport::tcp` and `Transport::tls` integrate standard TCP/TLS streams and
actively support context cancellation and deadlines. Dial and configure TLS
through `std::net` / `std::net::tls` before calling `client`. Incoming TLS can be
supplied as a custom transport; this package does not implement a TLS server.
`Transport::new` accepts context-aware read/write callbacks and a close callback.
`Transport::from_io` accepts any `std::io::{Read, Write, Close}` stream, including
short-read/short-write adapters. For generic streams, cancellation is checked
between calls; interruption of an already-blocked call requires a context-aware
transport implementation. As documented by std, cancelling active TLS I/O may
close the underlying TLS stream.

`Connection` supports concurrent callers with separate reader and writer gates.
A pending receive does not prevent another task from sending. Concurrent sends
serialize complete output frames; concurrent receives consume messages one at
a time. Gate waits honor context cancellation. A write failure aborts the
transport because an unknown prefix may already be on the wire. A TCP read
timeout retains a partial frame for a later receive. Applications should treat
an error after enqueueing a send as potentially transmitted; do not blindly
repeat application operations.

## Protocol engine and limits

`Decoder` incrementally parses frames, validates the client/server mask
orientation, opcodes, RSV bits, minimal unsigned length representation, control
frame limits and configured payload limits before allocating payload storage.
`encode_frame` accepts an optional explicit four-byte mask for codecs/tests.
`Session` automatically generates a fresh cryptographic client mask per frame.

`Session::feed`, `next`, `send` and `send_fragmented` expose the protocol engine
without I/O. Text and close reasons require valid UTF-8; text scalars may span
fragments. Continuations, interrupted messages, close status codes and aggregate
message sizes are checked. A ping/pong may occur between data fragments. Only
`Connection` synchronizes access: an individual `Session`, `Decoder` or
`HandshakeDecoder` requires one task or external synchronization.

`Limits::standard()` permits 1 MiB frames, 8 MiB assembled messages, 2 MiB input
and output queues, 256 queued output frames and 16 KiB HTTP handshakes. Configured
frame limits range from 125 bytes to 1 GiB. Limits are validated on construction.
Outgoing data is automatically fragmented at `frame_bytes`; `send_fragmented`
can choose a smaller size. A message must fit the available output queue and
frame-count budget atomically. Large-message callers can increase queue limits;
this version does not expose an unbounded streaming message writer.

`Backpressure` leaves rejected input/output unconsumed. Drain `next` before
retrying input, or `output`/`consume_output` before retrying a send. Incoming
control frames awaiting queue capacity are retained for retry. `output` returns
a copy of the next frame's unsent bytes; `consume_output` supports partial
writes. Payloads are snapshotted when enqueued. Protocol errors poison the
session; network connections abort on these errors. `finish` distinguishes a
complete closing handshake from truncated/abnormal EOF.

## Validation

```sh
cd ecosystem/websocket
../../stage2/bin/goml fmt --check
../../stage2/bin/goml test
GOFLAGS=-race ../../stage2/bin/goml test --target-dir _artifact/race --timeout 300s
```

Native GoML tests cover the RFC accept digest and exact Hello wire vectors,
every split of a handshake/frame, 16/64-bit lengths, invalid headers and
opcodes, masking direction, UTF-8 split across fragments, interleaved controls,
close states, bounded queue retry, payload isolation, short I/O, coalesced
upgrade bytes, real local TCP echo/close, concurrent send/receive, cancellation,
read-deadline resume and wakeup on abort. The independent versioned consumer in
`../consumers/websocket` performs a real TCP handshake, Unicode echo and closing
handshake using only exported APIs.

This version deliberately has no compression extensions, HTTP/2 extended
CONNECT, HTTP/3, redirects, proxy negotiation, URL parser or automatic heartbeat
scheduler. HTTP upgrade integration with `web`/`request` can use the public
handshake helpers and `Connection::upgraded`. Origin policy belongs to the
server application. Cross-implementation Autobahn certification is not claimed.
