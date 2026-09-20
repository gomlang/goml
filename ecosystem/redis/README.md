# Redis

`ecosystem::redis` is a standalone GoML RESP2/RESP3 client. It includes an
incremental binary codec, typed requests, heterogeneous pipelines, transactions,
optimistic updates, subscriptions, bounded push handling, cancellation and TCP
connection management, bounded pools, DNS/TLS and injectable transports. It uses `std::net`,
`std::net::tls`, `std::context`, `std::task` and `std::time`; it does not
wrap a Go Redis client.

## Connection and typed commands

```gom
use ecosystem::redis::{Connection, ConnectionOptions, Operation, Error, ErrorKind};
use ecosystem::redis::commands;
use std::net;

fn example() -> Result[(), Error] {
    let address = net::SocketAddr::parse("127.0.0.1:6379").map_err(
        |problem| Error::new(ErrorKind::Io, problem.to_string()),
    )?;
    let connection = Connection::connect(address, ConnectionOptions::standard(), Operation::new())?;
    defer { let _ = connection.close(); };
    let stored = connection.execute(
        commands::set("greeting", "hello", commands::SetOptions::standard())?,
        Operation::new(),
    )?;
    let value = connection.execute(commands::get("greeting")?, Operation::new())?;
    if stored {
        if let Some(blob) = value {
            println(blob.text()?);
        }
    }
    Result::Ok(())
}
```

`ConnectionOptions` selects `Protocol::Resp2` or `Resp3` (default), optional
`Credentials`, database, client name, resource limits and read buffer size.
Connection setup issues `HELLO`, optionally authenticates and names the client,
validates the returned protocol version, then selects the database. It therefore
requires Redis 6 or newer even in RESP2 mode. Failed setup closes the socket.
`close()` is idempotent and wakes active/queued operations; `quit()` exchanges
QUIT and then closes. Applications should explicitly close connections.

### Custom transports and dialing

`Connection::connect_host(host, port, options, operation)` resolves DNS and tries
the returned addresses through the standard TCP connector. It enables
TCP_NODELAY just like numeric-address `connect`.
`Connection::connect_tls(host, port, tls_config, options, operation)` uses
`std::net::tls::ClientConfig` for verified server names, system/custom trust roots,
client certificates, ALPN and TLS version policy. There is no insecure mode.
The operation's total deadline includes DNS, TCP/TLS setup and HELLO/SELECT;
the TLS configuration's connect timeout can impose an earlier setup limit.
`Transport::tls(stream)` also wraps an existing standard TLS stream.

`Transport::new(read, write_all, close, is_closed)` adapts an ordered duplex byte
stream. `Connection::from_transport` takes ownership immediately, including on
invalid options or failed setup. `Connection::dial(options, operation, dialer)`
validates options before dialing and gives the dialer the remaining `Operation`.
The same total deadline covers dialing, transport setup and HELLO/SELECT. This
allows applications to supply additional transports, such as Unix sockets,
without reimplementing RESP, pipelines, transactions or subscriptions.

Each read/write callback receives the remaining timeout and cancellation token,
available through `Operation.timeout()` and `cancel_token()`. An attached standard
context is available through `context()`; adapters must honor it as well. Native
socket adapters can use `wait_options()`. A read returns 0 at EOF; counts outside the
provided buffer are rejected. A successful write must transmit all bytes.
Callbacks must honor these limits, avoid retaining buffers after returning, and
return structured `std::io::Error` values. The connection also checks the total
deadline after callbacks return; it cannot preempt a callback that ignores it.

The connection serializes reads and writes. `close` and `is_closed` may run
concurrently with either callback: adapters must synchronize their state and
make close wake blocked I/O. The connection calls its close callback once;
directly shared `Transport` handles still require an idempotent close adapter.
Do not use the underlying stream independently after transferring ownership.
`Transport::tcp` wraps an already connected `std::net::TcpStream`; the normal
`Connection::connect` entry point additionally enables TCP_NODELAY.

The `commands` package provides typed factories for:

| Family | Commands |
| --- | --- |
| Connection | PING, ECHO, CLIENT ID, SELECT |
| Strings | GET, GETDEL, SET, SET GET, MGET, MSET, APPEND, STRLEN, INCRBY, DECRBY |
| Keys | DEL, UNLINK, EXISTS, RENAME, TYPE, EXPIRE, PEXPIRE, TTL, PTTL, PERSIST |
| Lists | LPUSH, RPUSH, LPOP, RPOP, LLEN, LRANGE, BLPOP |
| Hashes | HGET, HMGET, HSET, HDEL, HLEN, HGETALL, HINCRBY, HSCAN |
| Sets | SADD, SREM, SCARD, SISMEMBER, SMEMBERS, SSCAN |
| Sorted sets | ZADD, ZREM, ZCARD, ZSCORE, ZRANGE WITHSCORES, ZINCRBY |
| Iteration and scripts | SCAN, SCRIPT LOAD, EVAL, EVALSHA |
| Publishing | PUBLISH |

Factories return `Result[Request[T], Error]`, validate empty lists and invalid
options, and copy binary arguments. `SetOptions` combines an exclusive condition
with an exclusive expiration policy, including KEEPTTL and absolute expiration.
`set` returns whether the condition allowed storage; `set_get` returns the prior
value according to Redis's SET GET semantics. SCAN methods expose each cursor
and page; callers iterate until the cursor is zero and handle Redis's documented
possible duplicates. BLPOP's server-side zero timeout means indefinite waiting;
the client's default operation timeout still applies.

`Blob` carries arbitrary bytes; `Blob.text()` checks UTF-8. `ToArg` supports
strings, blobs, byte vectors/slices, booleans and numeric scalars, and applications
may implement it. `FromReply` supports blobs, UTF-8 strings, checked integer
widths, floats, booleans, OK unit status, options, vectors, pairs,
`Entries[K, V]`, `ScoredMembers` and raw `Value`. Vectors decode arrays/sets;
use `Blob` for a binary string. Signed integers require integer replies;
unsigned integers also accept decimal strings, needed by SCAN's unsigned cursor.
RESP2 flat pairs and RESP3 maps/nested scored members share typed results.

Use `request[T: FromReply](name, arguments)` for additional Redis commands,
`raw_request(Command)` for dynamic replies, or `Request::new(command, decoder)`
for custom decoding. `Request.map` composes fallible result conversions. A
consumer-defined `FromReply` implementation specializes through the dependency
interface. Protocol-changing commands, replication/monitoring modes and CLIENT
REPLY are rejected by ordinary execution; use the dedicated transaction,
subscription and close APIs. Raw commands that alter server state otherwise have
the semantics of those commands; there is no automatic retry or reconnection.

## Bounded connection pools

`Pool::connect(address, connection_options, pool_options)`, `connect_host` and
`connect_tls` construct a lazy pool. No socket is opened until `acquire` or
`execute`. Each new connection completes the same authenticated HELLO/SELECT
setup as `Connection`. `Pool::new(pool_options, connector)` accepts a custom
`Operation -> Result[Connection, Error]` factory. Each successful factory call
must transfer a fresh, exclusively owned connection; the factory must honor the
remaining operation timeout, cancellation token and context.

```gom
let pool = Pool::connect(address, ConnectionOptions::standard(), PoolOptions::standard())?;
defer { let _ = pool.close(); };
let pong = pool.execute(commands::ping()?, Operation::new())?;
let lease = pool.acquire(Operation::new())?;
defer { let _ = lease.release(); };
let values = lease.pipeline(plan, Operation::new())?;
```

`PoolOptions` has these defaults:

| Option | Default | Behavior |
| --- | --- | --- |
| `max_connections` | 16 | Bounds idle, leased and currently opening connections together; accepted range 1–65,536 |
| `max_idle` | 16 | Closes surplus connections on release; zero disables idle retention; cannot exceed capacity |
| `idle_timeout` | 300 seconds | Retires an idle connection on its next checkout; `None` disables idle expiry |
| `max_lifetime` | `None` | Optional lifetime limit, enforced at checkout and release |
| `health_check` | `HealthCheck::OnCheckout` | PING before reusing an idle connection |
| `health_timeout` | 1 second | Maximum PING budget, capped by the acquisition's remaining deadline |

Expiry durations and the health timeout must be positive. `HealthCheck::Never`
skips PING; `AfterIdle(duration)` checks only after the specified idle duration,
with zero equivalent to `OnCheckout`. Newly established connections rely on
HELLO validation. Expiration never interrupts a borrowed connection, and idle
age starts when it is returned, not when it was originally created. Reaping is
lazy; there are no periodic pool tasks or minimum-idle prefill. `stats()` provides
a synchronized snapshot of `capacity`, `idle`, `leased`, `opening` and `closed`;
`leased` includes an idle connection undergoing checkout validation.

`acquire(operation)` reserves capacity before dialing and waits when all slots
are occupied. One total deadline covers queueing, health checks and connection
setup. Context cancellation, context deadlines and legacy cancel tokens also
interrupt acquisition. Pool closure wakes waiters and cancels pending standard
connectors. Invalid configuration is `ErrorKind::Limit`, exhaustion expires as
`Timeout`, and cancellation/closure returns `Cancelled`/`Closed`. Acquisition
errors never set `may_have_executed`, because no caller command was sent. A failed
connection attempt frees its reservation and returns its error; it does not spin
or retry repeatedly. Waiters have no fairness guarantee or separate queue bound;
applications should bound their own concurrent work.

`Lease` provides `execute`, `pipeline`, `transaction` and `compare_and_set`.
Its connection remains private. Aliases of a lease share one execution gate;
`release()` waits for active use before returning the connection and prevents
later operations through any alias. Repeated release/discard calls are no-ops.
`discard()` closes the connection instead of retaining it. Explicitly release
leases, normally with `defer`; garbage collection does not return pool capacity.
Per-lease operations have their own supplied deadline, including waiting for
another operation on that lease. `pool.execute(request, operation)` acquires,
executes once and releases automatically under one shared total deadline.

A closed, expired or unhealthy idle connection is discarded before checkout;
the pool establishes a replacement on demand. It **never replays a caller
command**, including an ambiguous failed write. Application code receives the
original command error and `may_have_executed` flag. Failed health PINGs may be
replaced within the remaining acquisition budget. Standard TLS interruption
closes the stream, so releasing that lease evicts it. A complete server or typed
decoding error leaves an otherwise synchronized connection reusable.

Pooled requests reject session-changing AUTH, SELECT, READONLY, READWRITE,
ASKING and CLIENT commands other than ID/INFO/GETNAME, in addition to the ordinary
connection's protocol-state restrictions. Pipelines, transactions and WATCH
callbacks apply the same checks. Use a dedicated `Connection` for subscriptions
or mutable session configuration. Applications issuing custom module commands
remain responsible for avoiding module-specific session changes. WATCH callbacks
must not reenter their lease or acquire from an exhausted pool.

`pool.close()` is idempotent, immediately closes all tracked idle/borrowed
connections and wakes queued operations; it does not wait for leases to be
returned. It records and returns the first connection-close error. An in-flight
custom factory that ignores cancellation cannot be preempted; a connection it
returns after closure is immediately closed and never leased. Explicit
`release`/`discard` report a close failure; automatic `execute` cleanup preserves
the command result, so cleanup errors do not masquerade as a failed command.

## Pipelines and transactions

Create `Pipeline::new()`, enqueue typed requests with `queue`, retain each
`Ticket[T]`, call `connection.pipeline(plan, operation)`, then decode with
`Responses.get(ticket)`. The plan can mix unrelated result types. All command
replies are drained even when some contain server errors. Each ticket checks
pipeline identity, including independently constructed pipelines; `at` and
`raw` provide indexed access. A server/type error identifies the response index.
A pipeline can be reused; its builder and mutable argument/reply buffers must
not be mutated concurrently with use.

`connection.transaction(plan, operation)` holds the connection for MULTI,
queued commands and EXEC. It verifies MULTI before sending the plan. Queue-time
errors trigger DISCARD and return `TransactionOutcome::Rejected` with indexed
server errors. A successful EXEC returns `Committed(Responses)`; execution-time
errors remain individual results and **do not roll back** successful commands.
A WATCH conflict returns `Aborted`. Unexpected protocol/transport failure after
MULTI closes the connection so queued transaction state cannot escape.

`compare_and_set(keys, reads, build, operation)` performs WATCH, executes the
read pipeline, invokes `build(responses)` to create the transaction and runs
MULTI/EXEC under one connection gate. It clears watch state on application
errors and on completion; failed cleanup closes the socket. The callback must
not reenter this connection, wait on a task using it, or perform unbounded work.
Use another connection when an independent operation is required. Callers own
conflict retry policy. The callback is ordinary synchronous code and cannot be
preempted by the operation timeout.

## Deadlines, concurrency and pushes

`Operation::new()` has a five-second total timeout. `with_timeout`,
`without_timeout`, `with_cancel` and `with_context` configure each operation.
Context cancellation and legacy cancel tokens are both honored when supplied;
the earlier of the context deadline and operation timeout wins. `without_timeout` does
not disable a context deadline. Scoped cancellation bridges for DNS/TLS are joined
before their operation returns. The deadline covers
waiting for the connection, writing requests and reading every reply, including
interleaved push traffic. Continued network progress does not reset it.
Concurrent requests share a channel gate and cannot exchange each other's
replies. No background reader or connection-owned task is created.

Cancellation/timeout before transmission leaves the connection usable. Failed
transmission or interrupted reply collection closes it; `Error.may_have_executed`
marks this transport ambiguity and `partial` retains complete pipeline replies
already received. It does not prove whether the server committed a write, and
it is not a transaction rollback indicator. Server errors and typed decoding
errors after a complete ordinary reply leave framing synchronized.

RESP3 pushes encountered while collecting command replies are retained in a
bounded queue. `drain_pushes` removes currently queued values; `next_push` waits
for a queued or newly arriving push. Push limits count both frames and wire
bytes; overflow during a request closes the connection. `next_push` timeout or
cancellation preserves an incomplete decoder frame for the next call because
no request was sent, provided the transport remains open. Standard TLS closes
the stream when active I/O times out or is cancelled; check `is_closed()` and
establish a new connection in that case. Queued cancellation before TLS I/O
begins leaves it usable. Reading an unexpected ordinary reply this way is a
protocol failure.

`subscribe(channels, SubscriptionKind, operation)` enters dedicated channel or
pattern subscription mode. `Subscription.change` adds/removes explicit nonempty
lists, consumes and validates all acknowledgements, and queues messages arriving
between them. `next` returns `Event::Message(channel, payload, pattern)`,
subscription changes, pong or an unknown event. Binary channel names and payloads
are preserved. A read-only subscription timeout can resume a partial frame on
transports that remain open; standard TLS requires reconnection after active I/O
interruption. An interrupted subscription change closes the connection. Ordinary command
execution is rejected while subscribed, in both RESP versions. Removing the last
subscription restores ordinary mode. `count` reads the server's acknowledged
subscription count; `close` closes the underlying connection. Publishers use a
separate connection. Sharded subscriptions are not implemented.

## Codec and limits

`Value` preserves every RESP2/RESP3 type, including the three null encodings,
attributes, arbitrary map keys, duplicate pairs, big-number decimal text,
verbatim format prefixes, binary errors and push frames. `checked()` turns a
server error into structured `Error.server`; `without_attributes()` exposes the
attributed payload without changing the original value. Floats support NaN,
infinities and negative zero. Big numbers remain decimal text without narrowing.

`encode`/`decode` operate on one complete value; decode rejects trailing bytes.
`decode_prefix` returns the value and its consumed wire length. Their
`*_with_limits` variants use explicit `Limits`.
`Decoder::new(limits)`, `feed`, `next_frame`/`next`, `finish`, and `reset` implement
incremental parsing without reparsing previously consumed payloads or headers.
`next_frame` reports each root frame's wire byte count. `finish` requires drained,
complete input. A malformed/over-limit feed or decode poisons the decoder until
reset. `buffered` counts unconsumed input bytes; `is_partial` also detects bytes
already materialized into an incomplete frame.

Streamed bulk strings, arrays, maps and sets are accepted. The writer emits
fixed-size equivalents; it preserves values rather than original header spelling
or chunk boundaries. Redis 7.2 does not emit streamed aggregates itself, so
independent specification fixtures cover those forms. Simple string/error
payloads preserve binary bytes but exclude CR/LF. Numeric/control headers and
three-character verbatim format labels are checked.

Default limits are 16 MiB per frame, input buffer and blob; 64 KiB per header;
1,000,000 aggregate elements; 2,000,000 values; depth 128; 1,024 commands and
16 MiB encoded bytes per pipeline; 64 MiB received bytes per batch; and 128 pushes
with 16 MiB queued wire bytes. Root depth is zero. Attribute pairs and their
payload consume the value/depth budgets. Incremental input buffers can be smaller
than a frame because consumed payload bytes move into the decoded value. Limits
bound wire data and counts, not an exact process heap size. Cyclic caller-created
values terminate at the encoder's depth limit. Decoder/builders are mutable and
require external serialization if shared; `Connection` provides its own gate.

The built-in transport follows `std::net`: Linux amd64 numeric IPv4/IPv6 TCP
addresses. Custom transports extend that boundary. Cluster redirection and Sentinel discovery are outside this implementation. Command availability depends
on the selected server version; unsupported commands return normal server errors.

## Validation

```sh
just ecosystem-test redis
```

All verification is driven by GoML tests. Library tests cover RESP2/RESP3 codecs,
malformed frames, bounded allocation, incremental decoding, typed commands,
pipelines, transactions, transport failures, cancellation and concurrent pools.
The specification generator constructs binary scalars, streamed arrays/maps,
attributes and blob chunks independently of the codec and compares canonical wire bytes.

Consumer tests download the official Redis 7.2.5 tag archive with pinned SHA-256,
build it under `ecosystem/_artifact/reference`, and run both RESP protocols against
a fresh authenticated loopback server with persistence disabled. `std::process`
and scoped cancellation own the server lifetime; startup and subprocess commands
have deadlines. This requires curl, tar, make, a C compiler and network access on
the first run. Tests include command families, binary values, Lua, pipelines,
transaction errors, WATCH cleanup, Pub/Sub, concurrent pooled increments, and
health replacement after CLIENT KILL.

Fifteen further consumer cases exercise DNS, verified TLS, mTLS, untrusted roots,
wrong hostnames, handshake/HELLO/I/O deadlines, both cancellation APIs and pooled
TLS reuse/eviction. A test-only Go TLS peer generates ephemeral certificates;
GoML tests own scenarios and assertions and synchronize cancellation with request
arrival. The production Redis package remains entirely GoML. The shared verifier
runs native and generated tests with Go's race detector.

Protocol references: [Redis RESP specification](https://redis.io/docs/latest/develop/reference/protocol-spec/),
[RESP3 streamed types](https://github.com/antirez/RESP3/blob/master/spec.md),
and [Redis transactions](https://redis.io/docs/latest/develop/using-commands/transactions/).
