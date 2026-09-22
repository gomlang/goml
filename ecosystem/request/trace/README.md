# Request trace

`ecosystem::request::trace` records opt-in lifecycle events for a request
client. It uses a bounded, shared in-memory collector; it does not add a
runtime tracing backend or perform network operations itself.

```gom
use ecosystem::request::Client;
use ecosystem::request::trace;
use std::testing;

fn fetch(url: string) -> () {
    let record = testing::expect_some(trace::Trace::new(128), "trace");
    let client = testing::expect_ok(Client::builder().trace(record).build(), "client");
    defer client.close();
    let _ = client.get(url).send();
    for event in record.drain() {
        println(event.request_id.to_string() + ": " + event.detail);
    }
}
```

`Trace::new(capacity)` accepts 1–65,536 events. When full, the oldest event is
discarded. `snapshot()` returns a copy without clearing the collector;
`drain()` returns the events and clears it. Copies of a `Trace` share the same
collector, so concurrent requests receive distinct `request_id` values. The
collector adds no callbacks to connection-pool critical sections.

`RequestStart` records the original method and URL. Other events leave their
`method` and `url` fields empty and are correlated by `request_id`. The `attempt`
field is the zero-based redirect hop; a retry within that hop emits `Retry` but
does not increment `attempt`. `elapsed` is measured from creation of the trace,
not from the request start. `RequestDone` and `RequestFailed` mark the final
outcome, including redirects. HTTP status errors are successful transport
responses until the caller explicitly invokes `error_for_status()`.

Connection attempts emit `ConnectStart` and `ConnectDone`. `Connection` says
whether an established HTTP/1.1 or HTTP/2 connection was reused. HTTP/1.1
emits `RequestWritten` after its request bytes are sent. HTTP/2 emits
`RequestQueued` when a stream is submitted to the asynchronous writer; it does
not claim the bytes have reached the socket. `ResponseReceived` means the wire
response body has been buffered within its limit, before any automatic gzip
decompression. Proxy TLS tunneling contributes its own connection attempt.

DNS lookup, TLS negotiation and individual response-header timing are hidden
inside the existing networking APIs, so this package does not expose exact
Go `net/http/httptrace` hook parity. Recorded URLs may include query secrets,
and error details may contain endpoint information. Choose the capacity and
handling of snapshots accordingly.
