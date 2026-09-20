# web

`ecosystem::web` is an HTTP server framework with GoML routing, typed extraction, middleware, response construction and server-sent events. A small ordinary Go FFI adapter owns Go `net/http` connections, body streams, deadlines and server shutdown. It uses no third-party Go dependencies and requires no cgo.

The design takes inspiration from [Axum](https://docs.rs/axum/latest/axum/). It adapts the handler model to GoML closures and garbage collection, rather than reproducing Rust ownership or async futures. Transport behavior follows [Go net/http](https://pkg.go.dev/net/http).

## Using the library

Declare a versioned dependency and retain a minimal Go module for the executable:

```toml
[dependencies]
"ecosystem::web" = "0.1.0"
```

The library's `[native]` declaration supplies its native Go module through the GoML driver. Consumers do not need handwritten `require` or `replace` entries. The independent consumer under `../consumers/web` exercises this boundary.

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

Streaming sends and flushes headers before invoking the producer, including when the producer has not yet emitted an event. Writes execute synchronously against `net/http`; a slow receiver blocks the producer instead of growing a background queue. HTTP/1 full-duplex mode preserves request bodies while streaming a response. The transport and operating system retain their own bounded buffers.

SSE supports event name, ID, retry milliseconds and multiline UTF-8 data. CRLF/CR data is normalized into separate `data:` lines. Event names and IDs reject newline/NUL injection. `Response::sse` supplies the event-stream content type and cache/proxy buffering headers. Reconnect storage, replay, heartbeat scheduling and application event queues are the application's responsibility.

A request remains active until its response producer finishes. Body/request/writer aliases become invalid after dispatch ends, and late operations return `Closed`. Request-body reads and response writes are serialized internally. A client disconnect cancels `Request::wait`, subsequent body operations and writes. Finite streams and cancellation-aware loops should finish their own work before returning; the framework does not join arbitrary goroutines spawned by application code.

## Limits, deadlines and shutdown

`Limits::new` sets a 1 MiB request body limit, 1 MiB extraction limit, 1024 query/form fields, 4 MiB captured response limit, 30-second total request timeout, 5-second header timeout, 32 KiB header limit and 256 concurrent requests. `header_bytes` is passed to Go's `MaxHeaderBytes`, which includes Go's documented buffering allowance; it is not an exact byte cutoff. Both known-length and chunked bodies are checked. Concurrent admission rejects excess requests with 503 without building an unbounded wait queue.

The total request context and connection deadlines bound body reads, writes and cooperative work. `Request::check` and `wait` are suitable for application loops and external cancellation-aware operations. Arbitrary CPU loops, sleeps and foreign calls that ignore cancellation cannot be forcibly preempted. They continue to occupy an admitted slot. Expired connections are closed instead of reusing a cancelled HTTP connection context.

Shutdown first stops accepting connections and waits for active requests. On timeout it cancels server contexts and forcibly closes connections, returning `Timeout`. Concurrent shutdown calls share the first result and each caller has its own bound while waiting. Repeated shutdown is safe. Forced shutdown cannot terminate an uncooperative application function; it releases transport resources and relies on that function to observe cancellation. Callback panics on the serving goroutine become a generic 500 before headers or abort an already started response; this does not catch panics in unrelated application goroutines.

`dispatch(method, target, headers, bytes)` runs the same GoML routing and native body/stream machinery against an in-process response recorder. It does not open a socket. Handler-produced error responses remain ordinary captured responses; transport body-limit, timeout, panic and capture-limit failures return `Err`. Targets must use origin form (`/path?query`). Captured responses copy their bytes on access.

## Scope and verification

The current listener serves plain HTTP/1.1 over TCP. TLS termination, HTTP/2, HTTP/3, WebSocket upgrades, multipart extraction, static file serving, compression negotiation and CORS are not bundled. They can be implemented as later adapter or middleware additions. The library deliberately reserves transport framing headers and does not expose connection hijacking.

Run from the repository root:

```sh
just ecosystem-test web
```

GoML black-box tests, versioned-consumer tests, cached rebuild checks and native
adapter tests cover routing, typed payloads, request/body limits and lifecycle.
The consumer sends 120 generated Unicode query/form requests through
`ecosystem::reqwest`, exercises concurrent requests and checks real HTTP status
codes and duplicate headers. GoML `std::process` invokes curl as an independent
HTTP client for a 200 KiB chunked binary upload and a timed SSE disconnect; it then
asserts producer cleanup. Streaming events, shared counters and deadline responses
are tested over actual loopback connections. Native adapter tests cover
backpressure, admission, graceful/forced shutdown and transport failures. The
shared verifier runs applicable native and generated tests under the race detector.
