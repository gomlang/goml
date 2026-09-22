# HTTP testing helpers

`ecosystem::web::testing` wraps the ordinary web router and transport for
deterministic tests. It does not define another handler pipeline.

`Recorder::new(router)` records in-process dispatch results. Use
`recorder.request(method, target).header(name, value)?.body(bytes).send()` or
`recorder.record(method, target, headers, bytes)`. `exchanges()` returns a
snapshot of completed calls in completion order, including dispatch errors.
Concurrent calls serialize history updates; handler execution remains governed
by the router's normal limits and cancellation behavior.

`TestServer::start(router)` binds `127.0.0.1:0` and exposes `base_url()` and
validated origin-form `url(target)`. Shut it down explicitly with a positive
`Duration`; shutdown is idempotent and follows the underlying server's
graceful/forced semantics. This helper serves plain HTTP only.

`FaultTransport::new(router, actions)` consumes one `Fault` per dispatch call.
`Pass` uses ordinary dispatch; `Reject` fails before a handler runs;
`TruncateBody` dispatches a bounded prefix; `FailAfter` runs dispatch and then
returns an injected error if it succeeded. Calls beyond the script pass.
Scripts provide deterministic application/transport boundary failures, not
packet-level corruption or an adapter for `ecosystem::request::Client`.
