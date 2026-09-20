# CSV inventory consumer

This independent consumer imports `ecosystem::csv = "0.1.0"` and
`ecosystem::tempfile = "0.1.0"` through the isolated ecosystem registry.

The runnable example imports an inventory file with a UTF-8 BOM, quoted and
reordered headers, Unicode names, embedded commas, escaped quotes and a multiline
note. A flat Serde `Inventory` struct preserves text SKUs such as `0042`, decodes
unsigned quantities/prices and booleans explicitly, and maps empty notes to
`Option::None`. Three-byte reader buffers exercise boundary handling.

It summarizes active inventory using checked integer cents, exports a canonical
column order with automatic headers/BOM (including a header-only empty export), rewinds the real temporary output file
and imports it again. Both temporary files and their directory are cleaned up.
The expected inventory is three rows, two active products, six units and 6,960
cents. The example prints its summary and normalized CSV.

```sh
just ecosystem-test csv
```

For direct commands, obtain an isolated registry with
`ecosystem/verification/_artifact/bin/verification --registry-only` from the
repository root, then set the returned path as `GOML_HOME` while running
`../../../stage2/bin/goml check`, `goml test`, or `goml run` in this directory.
`GOFLAGS=-race goml test --target-dir _artifact/race` validates the same native
consumer tests with the race detector.

Tests cover real file round trips and cleanup, cross-module generic `Read`/`Write`
adapters, derived Serde conversion, leading-zero identifiers, optional notes,
quoted multiline fields, empty inventory exports, integer conversion errors and inventory-total overflow.
