# web

`ecosystem::web` is a pure GoML HTTP server framework with routing, typed extraction, middleware, response construction and server-sent events. HTTP parsing, TCP transport, body streams, deadlines and server shutdown use `std::net` and `std::context`. There are no Go adapters, native module dependencies or cgo requirements.

The design takes inspiration from [Axum](https://docs.rs/axum/latest/axum/). It adapts the handler model to GoML closures and garbage collection, rather than reproducing Rust ownership or async futures. The transport supports HTTP/1.1 keepalive, chunked bodies, pipelined requests and `100 Continue`.

## Using the library

Use this checkout's `stage2/bin` toolchain, built with `just make`. Panic isolation
requires the new `std::panic` API and panic-safe lexical cleanup, which are not yet
available in the pinned stage0 release.

Declare a versioned dependency:

```toml
[dependencies]
"ecosystem::web" = "0.1.0"
```

Consumers do not need a native dependency declaration or handwritten Go `require` or `replace` entries. The independent consumer under `../consumers/web` exercises the versioned dependency boundary.

```goml
use ecosystem::web::{Router, Response};
use std::time::{Duration};

let router = Router::new().get("/hello/{name}", |request| {
    Result::Ok(Response::text(200, "Hello " + request.params().required("name")?))
}).layer(|request, next| {
    Result::Ok(next.run(request)?.header("x-service", "example"))
});
let server = router.serve("127.0.0.1:0")?;
println(server.address()?);
server.shutdown(Duration::from_seconds(3))?;
```

`serve` binds before returning. Port zero requests an ephemeral port. Keep the application alive until its own shutdown signal, then call `shutdown` with a positive bound. The example above deliberately shuts down immediately; the consumer's `--serve` mode waits for standard input to close.

## Routing and middleware

- `get`, `post`, `put`, `patch`, `delete` and `route(method, path, handler)` create new routers; previous router values retain their routes and layers. Registration errors are retained and returned by `validate`, `dispatch` or `serve`.
- Patterns support complete `{name}` path segments and a final `{*path}` wildcard. Wildcards require a nonempty remaining path. Repeated capture names, duplicate methods on the same pattern and different capture names for the same route shape are rejected.
- Matching prefers literals over parameters over wildcards, independently of registration order. Method selection happens after finding the most specific matching path. A less specific route does not receive a request merely because its method matches.
- Matching splits the escaped path before percent decoding. An encoded slash stays within one capture; `+` is literal in paths. Query/form decoding converts `+` to space. Invalid UTF-8 and malformed percent escapes are recoverable errors.
- Paths are case sensitive and preserve repeated or trailing slashes. There is no implicit redirect, slash merging or dot-segment normalization. Route patterns are unescaped absolute paths and reject `%`, `?`, `#` and malformed braces. `Request::path` exposes the escaped path.
- GET supplies HEAD unless an explicit HEAD handler exists. HEAD executes the selected handler and middleware, suppresses response bodies and skips stream producers. Known paths supply OPTIONS with `Allow`; other unsupported methods return 405. Explicit OPTIONS overrides the automatic response.
- `nest` mounts child routes beneath a static prefix without a trailing slash. It wraps their child middleware. The parent's limits and fallback govern the combined router; child limits and fallback are not imported.
- Layers execute in registration order, first outermost, including fallback and method-error responses. A `Next` can be called explicitly; middleware can short circuit, change string attributes, inspect errors, or transform responses. Request-body aliases share one stream cursor.
- Application state is captured by handler closures. Use `Channel`, atomic types or another synchronized state container for concurrent mutation. Capturing an ordinary mutable `Ref` does not make it thread safe.

## Typed requests and responses

`Request` exposes method, escaped path, raw query, headers, path parameters, remote address, host, body, cancellation checks and cancellation-aware waits. `Headers` is an immutable builder with case-insensitive names, duplicate preservation, injection validation and copied entry snapshots.

`json[T: Deserialize]` accepts `application/json` and `application/*+json` with optional parameters. Missing or incompatible content type produces 415; invalid typed data produces 422. JSON output uses `Response::json[T: Serialize]`. Binary bytes, text, HTML, redirects, custom status and repeated response headers are also supported. HTML output accepts already prepared HTML and does not perform template escaping.

`query` and `form` return ordered `Pairs`, preserving repeated and empty fields. `query_as`, `form_as` and `Pairs::decode` use a GoML Serde deserializer. They support flat derived structs with string, character, boolean, signed/unsigned integer, float, present optional, and repeated vector fields. Numeric conversions retain width/range checks; string values such as `0001` stay strings. Scalar duplicates fail instead of silently taking a first/last value. Vector fields accept one or more occurrences. Unknown fields follow the normal Serde derive behavior; missing fields, including missing optional fields, follow the existing GoML derive's required-field policy. Nested objects, bracket-index conventions, enum fields and arbitrary map roots are not supported by this form format.

`Error` records a kind and message. Extractor errors become HTTP responses; internal/configuration/transport errors expose a generic message. Middleware can implement a different error response policy. Errors after response headers have been sent abort that stream and connection; they cannot change its HTTP status.

## Streaming and resource boundaries

`Body::read(maximum)` reads up to 1–65536 bytes and returns an empty buffer at EOF. `collect(limit)` is explicitly bounded. Body implements `std::io::Read` and `Close`; `StreamWriter` implements `std::io::Write`. `io::copy(request.body(), writer)` therefore supports a true streaming echo or transform without buffering the whole body.

`Response::stream(status, content_type, producer)` invokes the producer while the response is active. `StreamWriter::write` buffers through the HTTP transport, `send` writes and flushes, `text` sends UTF-8, and `event` sends an SSE event. Each bespoke write/send/text/event call accepts at most 65536 encoded bytes and returns a configuration error above that bound. The standard `io::Write` implementation uses partial writes of at most 65536 bytes, so `write_all`/`copy` handle larger input.

Streaming sends headers before invoking the producer, including when the producer has not yet emitted an event. Writes execute synchronously against the TCP stream; a slow receiver blocks the producer instead of growing a background output queue. Request bodies remain readable while streaming a response. The input reader uses 16 KiB reads and one queued chunk, and the operating system retains its own bounded buffers. HTTP/1.1 streams use chunked framing; HTTP/1.0 streams finish by closing the connection.

SSE supports event name, ID, retry milliseconds and multiline UTF-8 data. CRLF/CR data is normalized into separate `data:` lines. Event names and IDs reject newline/NUL injection. `Response::sse` supplies the event-stream content type and cache/proxy buffering headers. Reconnect storage, replay, heartbeat scheduling and application event queues are the application's responsibility.

A request remains active until its response producer finishes. Body/request/writer aliases become invalid after dispatch ends, and late operations return `Closed`. Request-body reads and response writes are serialized internally. A client disconnect cancels `Request::wait`, subsequent body operations and writes. Finite streams and cancellation-aware loops should finish their own work before returning; the framework does not join arbitrary goroutines spawned by application code.

## Limits, deadlines and shutdown

`Limits::new` sets a 1 MiB request body limit, 1 MiB extraction limit, 1024 query/form fields, 4 MiB captured response limit, 30-second total request timeout, 5-second header timeout, 32 KiB header limit and 256 concurrent requests. `header_bytes` is an exact bound including the request line, CRLF delimiters and final empty line; trailers have a separate bound of the same size. Both known-length and chunked bodies are checked. Ambiguous Content-Length/Transfer-Encoding framing, duplicate framing headers, invalid header names and control-byte injection are rejected. Concurrent admission rejects excess requests with 503 without building an unbounded wait queue. Idle keepalive connections do not occupy request slots; accepted connections have a separate cap of `concurrent_requests + 256`. Idle connections expire after 60 seconds.

The total request context and connection deadlines bound body reads, writes and cooperative work. `Request::check` and `wait` are suitable for application loops and external cancellation-aware operations. Arbitrary CPU loops, sleeps and foreign calls that ignore cancellation cannot be forcibly preempted. They continue to occupy an admitted slot. Expired connections are closed instead of reusing a cancelled HTTP connection context.

Shutdown first stops accepting connections, closes idle connections and waits for active requests. On timeout it forcibly closes connections and cancels their request contexts, returning `Timeout`. Concurrent shutdown calls share the first result and each caller has its own bound while waiting. Repeated shutdown is safe. Forced shutdown cannot terminate an uncooperative application function; it releases transport resources and relies on that function to observe cancellation.

Handlers, middleware, fallbacks and stream producers run inside `std::panic::catch` boundaries. Callback panics unwind their deferred cleanup and become an `Internal` error without exposing the panic payload. Before response headers are sent, the server returns a generic 500 and closes that connection. After headers are sent, including when a stream producer panics before its first write, the server aborts the stream and closes the connection without changing its status or writing a successful terminal chunk. The request's admission slot and transport resources are released, and other requests continue normally. Ordinary `Result::Err` values and cancellation retain their existing behavior. A panic propagated by a handler-owned `task::scope` is caught when it reaches that handler; detached application goroutines need their own recovery boundaries.

`dispatch(method, target, headers, bytes)` runs the same GoML routing and body/stream machinery against an in-process response recorder. It does not open a socket. Handler-produced error responses remain ordinary captured responses; callback panics, transport body-limit, timeout and capture-limit failures return `Err`. A panic returns a generic `Internal` error, and request/writer aliases expire before `dispatch` returns. Targets must use origin form (`/path?query`). Captured responses copy their bytes on access.

`ecosystem::web::testing` adds a reusable dispatch recorder with exchange
history, a loopback HTTP test-server wrapper, and scripted fault dispatches.
The child package reuses this router and server implementation; see its
`testing/README.md` for the bounded behavior and failure modes.

`ecosystem::web::proxy` adds a bounded reverse-proxy handler using the existing
request client. It strips hop-by-hop and untrusted forwarding headers on both
legs, retains duplicate end-to-end headers, passes backend redirects through,
and propagates the incoming cancellation context. See [proxy/README.md](proxy/README.md)
for body limits, response mapping and unsupported upgrade/streaming cases.

## Scope and verification

The current listener serves plain HTTP/1.1 and HTTP/1.0 over TCP. TLS termination, HTTP/2, HTTP/3, WebSocket upgrades, multipart extraction, static file serving, compression negotiation and CORS are not bundled. They can be implemented as later transport or middleware additions. The library deliberately reserves transport framing headers and does not expose connection hijacking.

Run from the repository root:

```sh
just ecosystem-test web
```

GoML black-box tests, versioned-consumer tests and cached rebuild checks cover
routing, typed payloads, request/body limits and lifecycle.
The consumer sends 120 generated Unicode query/form requests through
`ecosystem::request`, exercises concurrent requests and checks real HTTP status
codes and duplicate headers. GoML `std::process` invokes curl as an independent
HTTP client for a 200 KiB chunked binary upload and a timed SSE disconnect; it then
asserts producer cleanup. Streaming events, shared counters and deadline responses
are tested over actual loopback connections. GoML TCP tests cover early stream
arrival, backpressure, admission, idle keepalive, graceful/forced shutdown,
request framing, panic isolation, deferred cleanup and transport failures. The shared verifier runs generated tests
under the race detector.
