# Reqwest

`ecosystem::reqwest` is a GoML HTTP client inspired by the
[Rust reqwest blocking API](https://docs.rs/reqwest/latest/reqwest/blocking/index.html).
It provides reusable concurrent clients, request builders, typed JSON, forms,
multipart uploads, validated multivalue headers, URL handling, authentication,
redirect policies, cancellation and bounded response bodies.

GoML owns the public types, builder validation, response conversion, redirect
loop, origin checks and integration with `std::context`. A normal Go FFI adapter
uses [Go's net/http transport](https://pkg.go.dev/net/http) for HTTP/1.1, HTTP/2,
connection pooling, DNS, TLS, compression and proxies. It adds no compiler hooks,
runtime externs or external Go dependencies. This is a synchronous API; call it
from `std::task` tasks for concurrent requests.

## Requests and typed responses

```gom
use ecosystem::reqwest::{Client, Error};
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

```gom
use ecosystem::reqwest::{Client, Multipart, Part, Error};
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

`send()` reads the complete response within its limit and closes its native body
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
failures. Native failures are converted through an opaque FFI type rather than
Go's `error` interface graph.

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
Request limits are checked before network I/O; ordinary bodies are already
caller-owned buffers. Zero body limits are valid; zero timeouts, negative limits
and negative redirect counts are errors.

`close_idle_connections()` releases idle pooled connections. `close()` is
idempotent, cancels active operations, closes idle connections and prevents new
requests. Copies of a `Client` share this lifecycle. Use `defer client.close()`;
there is no finalizer. The transport and cancellation handles are synchronized,
and concurrent requests on a shared client are supported. Separate `Bytes`
buffers should be used if callers mutate bytes from different tasks.

## Cancellation, redirects and TLS

`send_with(context)` and `execute_with(request, context)` accept
`std::context::Context`. Its earlier deadline and explicit cancellation apply
through response-body completion, including all redirect hops. A scoped GoML
watcher cancels an ordinary native request control; it is joined before the call
returns. Request timeout, context deadline and client closure also interrupt
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
values, ignoring every cookie with a Domain attribute. Path, Secure, expiry and
replacement behavior come from the standard cookie jar. This avoids relying on
an absent public-suffix database but does not implement browser-wide domain
cookies. Cookie storage is memory-only.

## Scope and verification

This implementation does not claim Rust API or feature parity. Streaming
uploads/responses, async/await APIs, HTTP/3, WebSockets, Brotli/Zstd, a persistent
cookie jar, public-suffix-aware domain cookies, custom DNS resolution, custom
TLS backends, automatic application retries and middleware are not implemented.
Text decoding is strict UTF-8 and does not inspect charset labels. Header values
are UTF-8 strings rather than arbitrary octets. HTTP/2 over TLS uses ALPN; h2c
prior knowledge and protocol forcing are not exposed. Client-side file reading
is separate from multipart byte parts. Proxy handling is delegated to net/http;
proxy tunnel behavior is not exercised by the local interoperability script.

`go.mod` contains only the module path and Go version. The library declares
`native.go-module` in `goml.toml`; the driver generates requirements and
replacements pointing at the selected registry copy. The independent consumer
therefore needs only its own minimal `go.mod`, with no manual replacement. `testserver/` is a local HTTP/HTTPS fixture
used by tests and the consumer; the library's production API imports only
`adapter/`.

From the repository root:

```sh
python3 ecosystem/verify.py reqwest
```

The suite includes ten black-box GoML tests, an independently compiled registry
consumer, Python HTTP/1.1 interoperability and Go race checks for native and
GoML-generated clients. Tests use local ephemeral ports and generated
certificates, with no external services. Native race tests synchronize with
server request arrival before cancellation/closure. GoML tests cover the same
in-flight boundary, TLS/mTLS and HTTP/2, redirect credential stripping, gzip and
chunked limits, typed Serde, multivalue headers, multipart, cookies and pooling.
To run the supplemental checks after building/testing:

```sh
python3 ecosystem/reqwest/interop.py
python3 ecosystem/reqwest/race.py
```

Reference API: [Rust reqwest ClientBuilder](https://docs.rs/reqwest/latest/reqwest/blocking/struct.ClientBuilder.html),
[redirect policy](https://docs.rs/reqwest/latest/reqwest/redirect/struct.Policy.html),
and [Go HTTP client/transport](https://pkg.go.dev/net/http).
