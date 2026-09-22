# HTTP protocol utilities

`ecosystem::http` contains pure, transport-independent HTTP helpers shared by
ecosystem clients and servers. It does not open sockets or parse live streams.

`strip_hop_headers` validates a sequence of name/value pairs, removes standard
hop-by-hop fields and every field nominated by a `Connection` header, lowercases
retained names and preserves duplicate value order. An invalid header or
`Connection` token is a recoverable error. Reverse proxies should apply it on
both sides of the exchange and set trusted forwarding headers only afterward.

`dump_request` and `dump_response` construct bounded HTTP/1.1-style byte views
from explicit start-line parts, headers and a copied binary body. They reject
start-line/header injection and return a limit error rather than truncate. They
are diagnostics, not exact round trips: parsed requests may have lost original
header case/order, and HTTP/2 is represented as HTTP/1.1 text. Callers provide
their own headers and framing fields; these functions do not add or reconcile
`Content-Length`, chunked encoding, trailers or connection state.

```gom
use ecosystem::http::{Error, dump_request, strip_hop_headers};
use std::bytes::{Bytes};

fn debug_request() -> Result[Bytes, Error] {
    let headers = strip_hop_headers(
        Vec::from_array([("Connection", "x-private"), ("X-Private", "secret")]),
    )?;
    dump_request("GET", "/", headers, Bytes::new(), 1024)
}
```
