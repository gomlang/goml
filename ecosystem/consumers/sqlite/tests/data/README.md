# Independent reference fixtures

Reference: [Python standard-library sqlite3 linked to SQLite 3.45.1](https://www.sqlite.org/releaselog/3_45_1.html).

243 command sequences covering 2,747 operations. The independent SQLite engine produced storage values, prepared statements, transactions, savepoints, aggregates, joins, recursive CTEs, JSON and errors; generation used seed 20260920. REAL results use a 1e-12 absolute/relative tolerance. reference.sqlite is an independently produced SQLite database with one binary/Unicode row; the native GoML test copies it, reads and writes a real file, verifies rollback-on-close, then reopens read-only.

The fixture was captured once during migration of the verification harness. Fixture replay reads it directly with GoML, without a Python interpreter or package download. Inputs and expected values are independent of the GoML implementation. The separate live filesystem test also uses the system SQLite runtime described below.

Fixture SHA-256: `aac24e63d80b89dc392f7e2fa29909846858469ee1e03ee01ae4c8c3ac7942ca`.

`sqlite_reference.c` is a minimal SQL executor for the system SQLite shared library. The native GoML test compiles it with `cc -Wl,-l:libsqlite3.so.0`, checks that this independent engine can read GoML writes and observe rollback-on-close, asks it to insert binary/Unicode data, and then verifies those values through GoML. It uses public SQLite C ABI declarations directly, so development headers are unnecessary. All expected-value assertions remain in GoML.
