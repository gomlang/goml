# ecosystem::cache

A concurrent, bounded, generic in-memory cache written entirely in GoML. It
combines a hash table and a doubly linked LRU order, weighted admission, expiration,
and cancellable singleflight loading. No Go adapter or background worker is used.

```gom
use ecosystem::cache;
use std::context;

fn cached_configuration() -> Result[string, cache::Error] {
    let options = cache::Options {
        max_weight: 1048576,
        expiry: cache::Expiry::ttl(60000),
        ..cache::Options::new(128)
    };
    let values: cache::Cache[string, string] = cache::Cache::new(options)?;
    defer values.close();
    values.get_or_try_insert_with(
        context::Context::background(),
        "configuration",
        |load| {
            load.check()?;
            Result::Ok(cache::Loaded::new("ready", 5))
        },
    )
}
```

Declare `"ecosystem::cache" = "0.1.0"` in the module root's `[dependencies]`.
The versioned consumer under `ecosystem/consumers/cache` exercises independently
loaded package interfaces, generic specialization, and the JSON oracle protocol.

## Operations and bounds

| API | Behavior |
| --- | --- |
| `Cache[K, V]::new(options)` | Shared, synchronized cache for immutable values |
| `Cache::with_policy(options, copy, removed)` | Custom value isolation and removal notifications |
| `get(key)` | Return a value, update LRU and idle time, count a hit or miss |
| `peek(key)` | Return a value without touching LRU, idle time, or hit counters |
| `insert(key, value, weight)` | Replace and return the previous live value |
| `insert_with_expiry(key, value, weight, expiry)` | Override expiration for one insertion |
| `remove(key)` / `invalidate(key)` | Remove a value and invalidate any current load, including absent keys |
| `clear()` | Remove all values and invalidate all current loads; keep the cache open |
| `prune()` | Remove expired values and return the number removed |
| `keys_lru()` | Detached key vector from least to most recently used |
| `stats()` | Expire due entries and return a synchronized statistics snapshot |
| `get_or_try_insert_with(ctx, key, loader)` | Read or participate in one shared load |
| `close()` / `closed()` | Idempotent close and a broadcast receiver for closure |

Keys require `Eq + Hash`. Entry count and total weight are both bounded after
every mutation. Weights are nonnegative `i64`; zero-weight entries still count
against the entry limit. Oversized or negative weights are rejected without
replacing an existing value or invalidating its load. Weight arithmetic cannot
overflow, including a maximum `i64` weight. The caller chooses weights: they do
not automatically measure memory, reachable allocations, or the size of a value.

`Options::new(capacity)` requires a positive entry capacity and defaults to an
`i64::MAX` weight limit, 64 active loaders, 1024 waiting callers, and no expiration.
Weight and loader limits must be positive; zero waiting callers disables joining
pending loads. All configurable fields are public. Options are fixed after
construction. Cache handles copy shared identity; `keys_lru` returns a new vector.

Insertion makes a key most recent, resets both relative expiration clocks, and
evicts least recent entries until both limits permit admission. Expired entries
are removed before a capacity decision, so stale values never displace a live
value. Reads do not expose expired values. `remove` also cancels loading for an
absent or expired key. `clear` and `close` classify their removals as `Cleared` and
`Closed`; they do not run a separate expiration pass first.

## Expiration and clocks

`Expiry` supports a TTL since insertion, a TTI since the last successful `get`
or singleflight cache hit, and an absolute deadline in the configured clock's
millisecond coordinate system. The earliest enabled condition wins. Zero TTL or
TTI disables that condition; negative values are errors. An absolute deadline of
zero expires immediately. Expiration occurs at equality (`now >= deadline`).
`peek`, snapshots, failed reads, and statistics do not extend idle time.

`Expiry::none`, `ttl`, `idle`, and `until` are constructors; public fields allow
combinations. An insertion with an already elapsed absolute deadline replaces
any old value, generates an `Expired` notification, and retains no new entry.
A loader can return `Loaded::new(value, weight).with_expiry(expiry)` for expiry
computed from its result. Without an override, the cache-wide policy applies.
Relative loader expiry begins when its result is committed, excluding load time.
Such a loader still returns its result if its absolute deadline prevents caching.

`Clock::monotonic()` uses elapsed monotonic time. `Clock::new(callback)` supports
application clocks; calls occur outside the cache lock. Samples below zero are
clamped to zero and backward samples are clamped to the last observed value.
`ManualClock` is synchronized, starts at zero, and exposes `clock`, `now`, and a
checked nonnegative `advance`; it permits deterministic expiration tests.
Clock callbacks must be thread safe and finite. A custom clock must use the same
coordinate system as absolute deadlines.

Expiration is lazy. An idle cache retains expired entries until its next read,
insertion, load, prune, key listing, or statistics operation. Applications can
schedule `prune` themselves when prompt release of idle resources matters.

## Singleflight, cancellation, and invalidation

`get_or_try_insert_with` uses a cached hit immediately. On a miss, the first caller
executes its loader synchronously outside the lock. Other callers for that key
wait for the same generation's result; their own loader callbacks are unused.
Different keys load concurrently. Successes and failures are shared with current
waiters. Errors are never cached, so a later caller can retry. `Loaded[V]` carries
the value, its weight, and an optional expiration override.

The leader's context governs its loader. Each waiter can cancel independently
without cancelling the leader. Cancellation or deadline failure of the leader is
shared with its waiters and prevents result publication. Context is checked again
at publication. Cancellation racing with an already committed success does not
undo that success. Already-cancelled callers return a context error even on a hit.

The callback receives `LoadContext`, whose `check` and `sleep` also observe cache
closure and invalidation. `context`, `closed`, and `invalidated` expose the original
context and broadcast signals for integration with other blocking operations.
A loader must cooperate with these signals to interrupt its own I/O or computation.
The cache cannot forcibly stop arbitrary user code.

`insert`, `remove`, `invalidate`, `clear`, and `close` terminate affected flights
and immediately wake waiting callers with `Invalidated` or `Closed`. A later miss
can start a new generation while an obsolete loader is still running, subject to
admission limits. Unique checked generation identifiers ensure that an old result
cannot overwrite a replacement, resurrect a removed value, or complete a new
flight. An invalidated leader returns the invalidation error even if its callback
later returns a value. Identifier exhaustion is a recoverable `Exhausted` error.

Active callbacks, including invalidated callbacks that have not returned, count
against `max_inflight`. All joining callers count against a global `max_waiters`.
At either bound, admission returns `Busy` immediately; there is no unbounded
admission queue. A waiter releases its slot on completion or cancellation. Closing
rejects later operations, drops entries, and wakes waiters without joining user
callbacks. `stats().inflight` reports callbacks still completing after close.
The caller can use its own task scope to join them. No goroutine is owned by the
cache, and no explicit close is required to stop a hidden worker.

## Value isolation and callback contracts

`Cache::new` uses identity copies, suitable for values whose reachable data is
immutable. GoML `Vec`, `HashMap`, and `Ref` are shared mutable containers; use
`with_policy` with an appropriate deep copy for mutable values. The copy policy
runs on insertion/load ownership transfer, every returned value, and each removal
notification. A copy must not mutate its input and must be safe to call
concurrently. The caller must not mutate an input concurrently with its initial
copy. The cache protects its own structures; it cannot synchronize user aliases.

Keys and their reachable data must remain immutable while used by the cache.
`Hash` and `Eq` run inside the lock and must be pure, finite, and non-reentrant.
Clock, value copy, removal notification, and loader callbacks all run outside the
lock. They must be thread safe and return normally; panic recovery is not provided.
Callbacks may use ordinary cache operations, including reentrant removal and
insertion. Recursive loading must have an acyclic dependency graph: loading one's
own pending key, or an indirect cycle, waits on itself. GoML exposes no task-local
identity with which this library could distinguish that call from a valid waiter.
A clock or copy callback must likewise avoid recursively invoking itself.

The removal callback receives `(key, isolated_value, RemovalCause)`. Causes are
`Capacity`, `Expired`, `Replaced`, `Removed`, `Cleared`, and `Closed`. It executes
synchronously after the mutation releases the lock. Other tasks may have changed
the cache by then. Notifications for a loading leader are deferred until its
flight is published, enabling a removal callback to read that completed key.
Callbacks from different operations may overlap and their delivery order is not
globally serialized. Bulk clear order is unspecified. Callback failures have no
rollback semantics, and slow callbacks delay the initiating caller.

Counters saturate at `u64::MAX`. `hits` and `misses` count `get` and singleflight
lookups; `peek` does not count. `loads` counts admitted leaders, `coalesced` counts
admitted waiters, and `rejected` counts admission limit failures. Load outcomes
include invalidated and cancelled generations. `removals` counts all causes,
including replacements; `evictions` is capacity-only and `expirations` is
expiration-only. Live entry, weight, loader and waiter gauges do not saturate.

## Complexity and scope

With no expiring entries, hash lookup, LRU touch, removal, and each eviction are
expected O(1). Expiration currently scans the live table in O(n) time and O(n)
temporary storage whenever any entry has an expiration policy. Key listing and
clear are O(n). Insertion can evict O(n) entries. Storage is O(n + active loaders +
waiters), excluding user values and callback working memory. The cache maintains
one short state lock; it provides strict coherent bounds rather than sharded or
approximate admission. There is no periodic maintenance thread, expiry heap,
TinyLFU, refresh-ahead, stale-while-revalidate, persistence, or distributed protocol.

## Validation

Run `just ecosystem-test cache` from the repository root. It creates an
isolated versioned registry, formats/checks the projects, runs 22 black-box library
tests and independent consumer tests, verifies a stable cached build, and runs:

- 160 deterministic histories checked by native consumer tests against retained
  independent `OrderedDict` reference results, with 42,240
  operations and 20,303 query results across LRU, weights, TTL, TTI, Unicode values,
  replacement, removal, and clear.
- The complete library test executable rebuilt with Go's race detector, including
  simultaneous same-key loads, independent keys, cancellation, invalidation and
  replacement generations, closure, bounded admission, callback reentry, copied
  mutable values, and 3,600 concurrent mutation/read/expiry operations.

The regression suite also covers exact expiration boundaries, clock overflow and
backward samples, maximum weights, publication-time cancellation, shared failures,
and expiry notifications that reenter a completed singleflight.
