# sqlite

`ecosystem::sqlite` provides a typed GoML database API backed by an explicit Go
FFI adapter and `modernc.org/sqlite v1.39.1`. It includes parameter binding,
prepared statements, streaming and typed queries, transactions, nested
savepoints, cancellation, deadlines and deterministic resource closure.
The pinned driver embeds SQLite 3.50.4 in the tested Linux amd64 build.

## Native dependency setup

The GoML manifest is independent:

```toml
[dependencies]
"ecosystem::sqlite" = "0.1.0"
```

The Go adapter is a separate Go dependency named
`example.com/goml-ecosystem/sqlite`. It is local example infrastructure, not a
published Go module. The independent consumer's `go.mod` uses a local `replace`
to `../../sqlite`; its GoML dependency still resolves version `0.1.0` through the
isolated registry. Applications need an equivalent Go module mapping for the
adapter, plus its pinned transitive dependencies. GoML does not implicitly
translate GoML dependencies into Go module dependencies.

Fetch native modules explicitly before compiling:

```sh
cd ecosystem/sqlite
go mod download all
../../stage2/bin/goml test
```

`go.mod` and `go.sum` pin the driver and its dependencies, including the matching
`modernc.org/libc`. Compiler FFI checks and GoML builds use readonly module
resolution with network access disabled. The database driver itself does not
require CGo; the verification race detector does require a C compiler.
See the [driver documentation](https://pkg.go.dev/modernc.org/sqlite) and
[GoML FFI guide](../../docs/goml.md#go-ffi).

## Usage

Import `Executor` for the shared database/transaction methods and `FromValue` or
`ToValue` when calling their conversion methods directly.

```goml
use ecosystem::sqlite::{Database, Executor, Params, Value, Row, DbError};

fn example() -> Result[string, DbError] {
    let database = Database::memory()?;
    defer { let _ = database.close(); };
    database.execute("CREATE TABLE users(id INTEGER PRIMARY KEY, name TEXT)", Params::empty())?;
    database.execute(
        "INSERT INTO users VALUES(?, ?)",
        Params::positional(Vec::from_array([Value::Integer(7), Value::Text("Ada")])),
    )?;
    let row: Row = database.query_one("SELECT name FROM users WHERE id=?",
        Params::positional(Vec::from_array([Value::Integer(7)])))?;
    row.get(0)
}
```

`Database::memory()` creates an independent in-memory connection.
`Database::open(dsn)` accepts the driver's filename/URI DSN conventions, including
`file:/path/database.db?mode=ro`. `open_with` accepts `OpenOptions` and a `Control`.
Defaults enable foreign keys and set SQLite's busy timeout to 5,000 milliseconds.
A DSN is configuration, not a bound SQL parameter.

## Values, binding and row mapping

`Value` distinguishes `Null`, `Integer(i64)`, `Real(f64)`, `Text(string)`,
`TextBytes(Bytes)` and `Blob(Bytes)`. `TextBytes` preserves SQLite TEXT containing
invalid UTF-8 without silently changing its storage class. BLOB and text data
may contain NUL bytes. Empty blobs remain BLOBs, distinct from NULL. Native
binding and returned blobs use copies.

`ToValue` and `FromValue` support `Value`, `i64`, `f64`, `bool`, `string`, `Bytes`
and generic `Option[T]`. Boolean decoding accepts only INTEGER zero/one; `None`
maps to NULL. Conversions are strict: an INTEGER does not implicitly become
`f64`, and BLOB or invalid UTF-8 TEXT does not implicitly become `string`.
Use SQL `CAST`, inspect `Value`, or implement `FromValue` for another policy.
SQLite applies its own column affinity before values are read back.

`Params::positional` accepts ordered values. `Params::named` accepts `(name,
value)` pairs without the SQL `:`, `@` or `$` prefix. Names start with an ASCII
letter and continue with ASCII letters, digits or underscores; an empty name
represents an ordinal argument. Duplicate names are rejected. Values always go
through driver binding, including SQL-looking strings. Parameter values cannot
stand for table names, column names or SQL syntax. Missing and excess bindings
follow the driver's statement-binding behavior; do not depend on rejection of
unused extra arguments.

`Row` owns a snapshot, so advancing/closing a cursor does not invalidate earlier
rows. `value(index)` accesses a dynamic value; `get[T](index)` decodes it.
`index(name)` and `get_named[T](name)` use exact names and reject ambiguity.
`columns()` returns names and driver-reported declared types; computed columns
may have an empty declared type. `values()` returns a new vector.

Implement `FromRow` to map records. `Row.decode`, `Rows.next_as`,
`query_one`, `query_optional` and `query_all` use it. The independent consumer
defines both a `UserId` value mapping and a `User` record mapping, with nullable
fields, across the versioned module boundary.

The native driver sometimes returns Go `time.Time` for declared DATE/DATETIME
columns. The adapter exposes those as RFC3339Nano TEXT; original formatting is
not preserved in that case. Store date/time strings in TEXT columns or select a
TEXT cast when exact source formatting is required. SQLite converts bound NaN
to NULL; infinities remain REAL in the tested engine.

## Queries and statements

| API | Behavior |
| --- | --- |
| `Executor.execute` / `execute_with` | Execute one statement; return rows affected and connection-local last insert ID |
| `Executor.query` / `query_with` | Return a streaming `Rows` cursor |
| `Executor.prepare` / `prepare_with` | Create a reusable statement tied to its database/transaction |
| `Statement.execute`, `query` and `_with` variants | Rebind parameters for each execution |
| `query_all[T](..., max_rows)` | Bounded collection; closes the cursor on success/error |
| `query_one[T]` | Require exactly one row, reporting `NoRows` or `MultipleRows` |
| `query_optional[T]` | Accept zero/one rows; reject multiple rows |
| `Rows.next` / `next_as[T]` | Return `Option[Row]` / `Option[T]`, with repeated EOF allowed |
| `Rows.try_for_each` | Fallible callback; false stops early and closes the cursor |
| `Rows.columns` | Inspect result columns, including an empty result |
| `Statement.close`, `Rows.close` | Idempotent explicit closure |

Execution reports SQLite's changes and last-insert-ID values; a statement that
performs no insert can retain the preceding insert ID. Prepared statement
validation may be deferred by the native driver until execution.

Each call accepts one SQL statement, up to 1 MiB and 32,766 supplied parameters.
The scanner handles comments, escaped quotes, trailing semicolons and CREATE
TRIGGER bodies with CASE expressions. Multiple statements and direct
BEGIN/COMMIT/END/ROLLBACK/SAVEPOINT/RELEASE are rejected so SQL cannot bypass
managed transaction state. Run a sequence of statements inside `transaction`
when it must be atomic. This wrapper does not expose a multi-statement script API.

`Rows.collect` bounds retained row count, not bytes per value or total process
memory. Streaming avoids retaining an entire result. A cursor owns its query
context until EOF, failure or explicit close. Callers using `next` directly
should use `defer` to close a partially consumed cursor.

## Transactions and resource lifecycle

`Database.begin` supports `Deferred`, `Immediate` and `Exclusive` modes.
`Transaction` implements `Executor`; `commit`, `rollback` and `close` control
its lifetime. `savepoint()` returns a nested transaction handle. Only the
innermost active handle can execute or commit; rolling back an ancestor also
invalidates descendants and closes their resources.

`Database.transaction(mode, callback)` commits an `Ok` callback result and rolls
back an `Err`. `Transaction.with_savepoint(callback)` applies the same rule to a
nested savepoint. Both install deferred rollback for unwinding. Primary callback
or commit errors are preserved, with a separate `cleanup` message if rollback
also fails. GoML panics still propagate; they are not converted to `DbError`.

A failed commit whose native transaction remains active, such as a deferred
foreign-key constraint failure, permits correction, retry or rollback. If SQLite
has ended that transaction, recovery invalidates its handles instead. Resources owned by the
transaction are closed before the commit attempt. A successful commit or
rollback invalidates all handles belonging to the ended scope.

SQLite can automatically roll back a transaction after certain failures,
including `INSERT OR ROLLBACK`. The adapter maintains private sentinel savepoints
and checks them after execution and commit failures. If the sentinel was lost, it closes
transaction resources and invalidates every transaction handle before allowing
further operations. Ordinary statement-level constraint failures retain the
transaction, allowing a nested savepoint to roll back and its parent to
continue. Regression tests cover both paths. See SQLite's
[transaction error behavior](https://www.sqlite.org/lang_transaction.html) and
[conflict resolution rules](https://www.sqlite.org/lang_conflict.html).

A database owns one physical connection. Calls are serialized; multiple
independent databases can operate concurrently. An active cursor blocks new
execution/preparation on that connection with `Busy`. During a transaction,
root database handles and prepared statements owned by a parent scope similarly
return `Busy`. Close or drain a cursor before another statement. This prevents
pool self-deadlocks and implicit execution outside a managed transaction.

`Database.close()` cancels active work, closes cursors and prepared statements,
rolls back unfinished transactions and releases the connection and pool. It is
idempotent. Closing a prepared statement closes its active cursor. Garbage
collection is not a resource-closing API. `resources()` reports outstanding
statements, cursors and transaction depth for diagnostics.

## Cancellation and errors

`Control::unlimited()` creates a cancellable control. `Control::timeout(ms)`
adds a deadline beginning at construction; zero means no deadline. The control
can cover several operations and `cancel()` may be called from another task.
Use `_with` methods for explicit control, and cancel timeout controls after use
to release their timers. `begin_with` controls transaction creation, not its
entire subsequent lifetime.

Waiting for the connection gate observes cancellation/deadlines. Query control
remains effective during iteration; SQLite execution receives the operation
context, and database close cancels the root context. Timeout is cooperative,
including driver interruption and SQLite lock handling, rather than a hard
real-time guarantee. Set SQLite's busy timeout deliberately when sharing a file
with other writers. Rollback and close perform cleanup independently of an
already-cancelled operation control.

`DbError` exposes `kind`, `message`, extended SQLite `code`, `primary_code()` and
optional cleanup information. `Busy` distinguishes connection occupancy through
code zero from engine lock errors through SQLite codes. Constraint errors retain
the extended code. `native_cause()` retains an opaque native error object and
provides message bytes and code through `Cause`; the Go adapter's `FailureError`
recovers the original error for native `errors.Is`/`errors.As` consumers.

## Verification and compiler boundaries

```sh
python3 ecosystem/verify.py sqlite
```

The matrix includes 13 GoML library tests, a separate consumer, cached builds,
243 independent Python SQLite sequences with 2,754 checked operations, and
bidirectional database-file interoperability. The reference run used Python
SQLite 3.45.1 against the pinned adapter's SQLite 3.50.4. It compares storage
classes and exact integers/blobs/text, with tolerances for REAL values. File
checks also cover read-only opening and rollback on close.

`race.py` runs four native Go tests under the race detector, then builds and
runs all 13 GoML tests with `go build -race`. Native tests cover lock contention
between physical connections, cancellation while waiting for the connection
gate, blob copying, retained error identity and a native-transaction termination
fault injected before commit.

External handle types use small named Go interfaces rather than exposing the
large `database/sql` implementation graph. This avoids the current FFI type-graph
comparison limit while keeping concrete state in the native adapter.

GoML 0.1.50 keeps `std::io::Error` and `std::ffi::Error` distinct. The consumer
uses standard I/O directly in its FFI-importing package; the former transport
wrapper has been removed. The retained
[regression reproducer](../repros/ffi_error_alias/README.md) checks this boundary.

The library does not provide a connection pool, ORM/migrations, asynchronous
query objects, custom SQL function registration, backup or incremental BLOB APIs.
Callers can execute SQLite queries, joins, recursive CTEs, window functions,
JSON functions, ordinary PRAGMAs and triggers through the typed statement API.
Changing SQLite journaling policies remains subject to SQLite's own guarantees.
