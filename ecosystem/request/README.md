# Request

`ecosystem::request` is a GoML HTTP client inspired by the
[Rust reqwest blocking API](https://docs.rs/reqwest/latest/reqwest/blocking/index.html).
It provides reusable concurrent clients, request builders, typed JSON, forms,
multipart uploads, validated multivalue headers, URL handling, authentication,
redirect policies, cancellation and bounded response bodies.

The implementation is pure GoML: URL and MIME encoding, HTTP/1.1 framing,
HTTP/2 and HPACK, connection pooling, cookies, proxies and redirects are ordinary
GoML code. TCP, DNS, TLS and cancellation use `std::net`, `std::net::tls` and
`std::context`; gzip uses the pure GoML `ecosystem::archive` implementation.
There are no production Go adapters, Go FFI declarations or native module
dependencies. This is a synchronous API; call it from `std::task` tasks for
concurrent requests.

## Requests and typed responses

```gom
use ecosystem::request::{Client, Error};
use std::serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize)]
struct Message {
    text: string,
    count: isize,
}

fn send_message(url: string) -> Result[Message, Error] {
    let client = Client::builder().user_agent("goml-example").build()?;
    defer client.close();
    client.post(url)
        .bearer_auth("token")
        .json(Message { text: "hello", count: 1 })
        .send()?
        .error_for_status()?
        .json()
}
```

Reuse a `Client` across requests to retain its connection pool. `Client::new()`
and `Client::builder().build()` return recoverable errors. Convenience methods
cover GET, HEAD, POST, PUT, PATCH and DELETE; `request(method, url)` accepts other
HTTP token methods except CONNECT and TRACE. The top-level `get(url)` creates
and closes a temporary client.

A `RequestBuilder` is a value that preserves its first validation error until
`build()` or `send()`. It supports:

- `header` replaces a header; `append_header` retains multiple values; `headers`
  replaces the supplied names while preserving their multiple values.
- `query` appends encoded pairs to existing query parameters, retaining duplicate
  names and encoding Unicode. Pair order between distinct names is sorted;
  repeated values preserve insertion order.
- `body(Bytes)`, `text`, generic `json[T: Serialize]` and `form` produce replayable
  buffered bodies. Input bytes are copied, so later caller mutation is isolated.
- `multipart(Multipart)` accepts text fields and byte/file parts. `Part::bytes`,
  `file_name` and `mime_type` configure a part without reading the filesystem.
  Metadata is validated and MIME encoding uses a random boundary. Multipart
  output, including MIME overhead, is capped by the client's request limit.
- `basic_auth` rejects usernames containing `:`; `bearer_auth` supplies an
  Authorization header. Header names and values reject control injection.
- `timeout`, `body_limit` and `redirect` override the request's defaults.

`build()` yields an inspectable `Request` with `method`, `url`, `headers` and
copied `body` accessors. `Client.execute` and `execute_with` send it. Builders,
requests and buffered responses can be reused. `Headers.entries`, response
history and body accessors return snapshots; mutating them does not mutate the
original request or response. Host, Content-Length, Transfer-Encoding,
Connection, Proxy-Connection, Trailer and Upgrade are transport-managed and
cannot be supplied manually.

Request, response and multipart bodies use private `FrozenBytes` snapshots.
`RequestBuilder.frozen_body(snapshot)` and `Part::frozen_bytes(snapshot)` accept
an existing immutable snapshot without copying it; `Request.frozen_body()` and
`Response.frozen_bytes()` share it for repeated inspection or forwarding.
The existing `body(Bytes)` and `Part::bytes(Bytes)` snapshot mutable inputs, and
the `body()` / `bytes()` accessors still return independent mutable copies.
This does not change the buffered request/response size limits.

```gom
use ecosystem::request::{Client, Multipart, Part, Error};
use std::bytes::{Bytes};

fn upload(client: Client, url: string, bytes: Bytes) -> Result[(), Error] {
    let form = Multipart::new()
        .text("description", "upload")
        .part("file", Part::bytes(bytes).file_name("data.bin"));
    let _ = client.post(url).multipart(form).send()?.error_for_status()?;
    Result::Ok(())
}
```

## Responses, limits and lifecycle

`send()` reads the complete response within its limit and releases its transport
before returning, on both success and failure. The resulting `Response` owns no
open connection. It exposes status, headers, final URL, HTTP version, redirect
history, copied bytes, strict UTF-8 text and generic JSON decoding. Its
`content_length()` is the actual buffered byte count, including after automatic
decompression; a server's declared length remains accessible in its headers.

`StatusCode` classifies informational, successful, redirect, client-error and
server-error status families. Receiving an HTTP error status is a successful
transport operation; `error_for_status()` converts 4xx/5xx into an `Error` with
`ErrorKind::Status` and the status number. Other error kinds distinguish builder,
transport, timeout, cancellation, body limit, redirect, closed and decoding
failures. All protocol and socket errors are recoverable GoML values.

| Default | Value | Configuration |
| --- | --- | --- |
| Whole request timeout | 30 seconds | `timeout(Duration)` |
| TCP/TLS setup timeout | 10 seconds | `connect_timeout(Duration)` |
| Pool idle timeout | 90 seconds | `pool_idle_timeout(Duration)` |
| Idle connections total/per host | 16 | `pool_max_idle(isize)` |
| Response header budget | 64 KiB | `max_header_bytes(i64)` |
| Response body budget per hop | 16 MiB | `body_limit(isize)` |
| Request body budget | 16 MiB | `request_body_limit(isize)` |
| Redirect limit | 10 hops | `redirect(RedirectPolicy)` |
| Automatic gzip | enabled | `gzip(bool)` |
| Cookie store | disabled | `cookie_store(bool)` |
| Proxy environment | ignored | `environment_proxy(bool)` |

The body limit is enforced for declared lengths, chunked messages and
uncompressed bytes produced by gzip. HEAD ignores the representation's declared
length because it has no response body. Redirect response bodies have the same
per-hop limit. An oversized body fails instead of returning a truncated success.
Automatic gzip permits a bounded wire buffer of the body budget plus 64 KiB
and 0.1% framing overhead before enforcing the decompressed body budget.
Request limits are checked before network I/O; ordinary bodies are already
caller-owned buffers. Zero body limits are valid; zero timeouts, negative limits
and negative redirect counts are errors.

`close_idle_connections()` releases idle pooled HTTP/1.1 and HTTP/2 connections,
leaving HTTP/2 connections with active or queued streams open. `close()` is
idempotent, cancels active operations, closes idle connections and prevents new
requests. Copies of a `Client` share this lifecycle. Use `defer client.close()`;
there is no finalizer. The transport and cancellation handles are synchronized,
and concurrent requests on a shared client are supported. Separate `Bytes`
buffers should be used if callers mutate bytes from different tasks.
If an idle HTTP/1.1 connection was closed by its peer, an empty-body GET, HEAD
or OPTIONS is retried once on a fresh connection only when no response byte
has been received. Other transport failures are returned directly.

HTTP/2 connections retain their HPACK decoder, stream identifiers and connection
flow-control windows across requests. Concurrent requests to the same target
share a connection, including through HTTP and HTTPS CONNECT proxies. A single
reader dispatches frames to independent streams and a single writer serializes
frames; uploads use per-stream and connection windows with round-robin DATA
scheduling. The peer's concurrent-stream limit queues excess requests without
opening redundant connections. Connection setup is coordinated per target, so
a slow TLS handshake does not delay requests to unrelated targets.

Cancellation, request deadlines and response-body limits reset only the affected
HTTP/2 stream. Client closure closes the connection and joins its reader, writer,
idle timer and any CONNECT relay tasks. Idle expiry prevents subsequent reuse;
idle timers also reclaim unused connections. `pool_max_idle` bounds the combined
idle HTTP/1.1 and HTTP/2 pool, not the number of active HTTP/2 streams.

A graceful GOAWAY drains accepted streams while new requests use another
connection. Buffered requests rejected by GOAWAY, refused before response headers
with REFUSED_STREAM, or still queued when a connection fails are replayed at most
twice. Requests that might already have been processed are not automatically
retried. These protocol retries remain within the original request deadline.

## Cancellation, redirects and TLS

`send_with(context)` and `execute_with(request, context)` accept
`std::context::Context`. Its earlier deadline and explicit cancellation apply
through response-body completion, including all redirect hops. A scoped GoML
watcher connects client closure to the request context and is joined before
the call returns. Request timeout, context deadline and client closure interrupt
DNS, TLS handshakes and body reads.

`RedirectPolicy::None` returns redirect responses unchanged;
`Limited(count)` follows 301/302/303/307/308; `SameOrigin(count)` rejects changes
of scheme, hostname or effective port. Relative Location values are resolved
against the current URL. POST becomes GET on 301/302, and 303 converts non-HEAD
methods to GET; 307/308 retain method and body. A rewritten GET drops its body,
Content-Type and Content-Encoding.

Origin changes remove Authorization, Proxy-Authorization, Cookie, Cookie2 and
WWW-Authenticate headers. This includes a change of port on the same hostname.
These headers stay removed if a later redirect returns to the original origin.
All HTTPS-to-HTTP redirects are rejected. Custom application secret headers are
not classified automatically; use `SameOrigin` or disable redirects for them.
The optional cookie jar supplies cookies for the target URL independently.

TLS verifies system trust roots and hostnames, with TLS 1.2 as its minimum.
`add_root_certificate(pem)` appends PEM trust roots; `identity(certificate, key)`
configures mTLS. There is no insecure-verification switch. `https_only(true)`
rejects plaintext requests. URL credentials are rejected; use authentication
methods instead. Proxy selection is explicit: `proxy(url)`,
`environment_proxy(true)`, or `no_proxy()`.

The optional cookie store intentionally accepts only host-only Set-Cookie
values, ignoring every cookie with a Domain attribute. GoML implements path
matching, Secure, Max-Age, HTTP-date expiry and replacement. This avoids relying on
an absent public-suffix database but does not implement browser-wide domain
cookies. Cookie storage is memory-only.

## Scope and verification

This implementation does not claim Rust API or feature parity. Streaming
uploads/responses, async/await APIs, HTTP/3, WebSockets, Brotli/Zstd, a persistent
cookie jar, public-suffix-aware domain cookies, custom DNS resolution, custom
TLS backends, automatic application retries and middleware are not implemented.
Text decoding is strict UTF-8 and does not inspect charset labels. Header values
are UTF-8 strings rather than arbitrary octets. HTTP/2 over TLS uses ALPN;
h2c prior knowledge and protocol forcing are not exposed. Server push is
disabled, request header encoding uses literals without dynamic indexing, and
priority hints do not alter the round-robin upload scheduler. HTTP/1.1 connections
are reused and idle age is checked when checking a connection out of the pool. Client-side file
reading is separate from multipart byte parts. HTTP and HTTPS forward proxies
and CONNECT tunnels are implemented in GoML. CONNECT uses a lifecycle-managed
loopback relay so the standard TLS client can verify the original target hostname.
Proxy credentials use the Proxy-Authorization header; credentials embedded in
proxy URLs are rejected. Environment proxy bypass supports hosts, suffixes,
host:port, wildcard and loopback addresses, but not CIDR ranges. Certificate
PEM framing is checked during building; TLS validates certificate contents
and identity keys when connecting.

`go.mod` belongs solely to the local `testserver/` interoperability fixture;
the library manifest has no native module declaration. The independent consumer
and its local HTTP peer are pure GoML and need no `go.mod`.
The RFC 7541 HPACK table constants were transcribed from Go's vendored
`x/net/http2/hpack`; its BSD license is retained in `LICENSE.hpack.txt`.

From the repository root:

```sh
just ecosystem-test request
```

GoML library and consumer tests use local ephemeral HTTP/HTTPS servers and
fresh certificates. Consumer tests check HTTP/1.1, Unicode query/form data,
duplicate headers, typed JSON, redirects, response limits and chunked bodies.
Native Go test fixtures provide HTTP/TLS transport; all scenarios and assertions
run from `#[test]`. Cancellation tests synchronize with server request arrival.
HTTP/2 wire tests cover persistent HPACK, interleaved streams, flow-control
isolation, queued cancellation, GOAWAY draining, REFUSED_STREAM and malformed
frames. HTTPS tests verify shared-socket parallelism, reuse, idle closure and
per-stream cancellation through direct and proxied connections.
The shared verifier runs native test-server checks and GoML-generated
clients under Go's race detector. No external service is required.

Reference API: [Rust reqwest ClientBuilder](https://docs.rs/reqwest/latest/reqwest/blocking/struct.ClientBuilder.html),
[redirect policy](https://docs.rs/reqwest/latest/reqwest/redirect/struct.Policy.html),
and [Go HTTP client/transport](https://pkg.go.dev/net/http).
