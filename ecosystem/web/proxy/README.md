# Reverse proxy

`ecosystem::web::proxy::Proxy` is a bounded reverse-proxy handler built on the
existing `ecosystem::web` server and `ecosystem::request` client. Create and own
a reusable client, then attach `proxy.handler()` as a router fallback or route.

```gom
use ecosystem::request::Client;
use ecosystem::web::{Router, Error};
use ecosystem::web::proxy::Proxy;

fn proxy_router(client: Client) -> Result[Router, Error] {
    let proxy = Proxy::new("http://127.0.0.1:9000/api", client, 1048576, 1048576)?;
    Result::Ok(Router::new().fallback(proxy.handler()))
}
```

The target must be an HTTP(S) URL without a query or fragment. Its base path is
prefixed to the incoming path, and the raw incoming query is retained. Redirects
from the backend are passed through; the outbound client does not follow them.
Request and response bodies are buffered under separate explicit limits. Large
incoming bodies return 413; upstream transport or response-limit failures
return 502, and upstream timeouts return 504. The incoming request context
cancels the outbound exchange on disconnect or deadline.

Hop-by-hop headers, including names nominated by `Connection`, are removed on
both legs. The proxy does not forward an incoming `Forwarded` or
`X-Forwarded-*` value. It sets `X-Forwarded-For` from the connected peer,
`X-Forwarded-Host` from the inbound Host and `X-Forwarded-Proto` to `http`,
which matches the current plain-HTTP web server. Host and framing headers are
recreated by the respective transports. This API does not support WebSocket
upgrades, CONNECT tunnels, trailers, streaming bodies or informational
responses; those require transport capabilities beyond buffered handlers.
