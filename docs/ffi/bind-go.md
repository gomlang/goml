# Allowlisted binding generation

The configuration uses versioned JSON and explicitly selects each native function and type. The command is `goml bind-go <CONFIG>`, with optional `--compiler <COMPILER>` and `--dry-run`; `gomlc bind-go <CONFIG>` also works directly. Relative output paths resolve from the configuration file, within the existing GoML and Go module roots. Generation must not create or modify dependency manifests or execute Go package initializers.

```json
{
  "version": 1,
  "package": "bindings",
  "output": "bindings/generated.gom",
  "go_package": "native_bindings",
  "go_output": "native_bindings/generated.go",
  "go_import_path": "example.com/host/native_bindings",
  "functions": [
    {"name": "count", "import_path": "strings", "symbol": "Count"}
  ],
  "types": [
    {"name": "Duration", "import_path": "time", "symbol": "Duration"}
  ]
}
```

`functions` and `types` are explicit allowlists; omitted lists are empty and the combined list must contain at least one entry. Names are unique across both lists. GoML names must be identifiers other than the reserved `ffi` alias and `__goml_bind_` or `GomlBindType_` prefixes. Type names start with an uppercase ASCII letter; function and package names start with a lowercase ASCII letter or underscore, following GoML declaration syntax. Native symbols must be exported ASCII identifiers. Unknown or duplicate fields, unsupported versions and malformed arguments produce errors. Configurations are limited to 1 MiB, 4096 bindings and the existing protocol nesting/type complexity limits.

Each entry may have `type_arguments`, a finite array using the [bridge-type JSON schema](protocol-v1.md). For example, `[{"tag":"int64"}]` requests one concrete Go `int64` argument. The Go checker validates instantiation and constraints. Function queries preserve both the original declaration signature and the concrete signature; GoML substitutes the requested arguments and verifies that the signatures agree.

The Go output is intended for native forwarding functions that retain explicit generic arguments and variadic slice expansion. It is constructed through Go AST nodes. The GoML output exposes the selected bindings with explicit raw boundaries and no implicit error, nullable or record conversion. Such higher-level adapters require explicit mapping configuration before they can be generated.

Publication validates module-relative output paths and Go import identity, retains deterministic output, and refuses to overwrite files whose contents no longer match the generator's recorded output. A second identical generation leaves files unchanged. `--dry-run` validates configuration, module boundaries, and output paths and prints destinations without writing files or querying native symbols. See the [standard-library example](../../examples/ffi-bind-go/README.md) for generation, checking and execution commands.

The publisher records both source hashes and module-relative output names in `<CONFIG>.goml-bind.json`. It accepts existing outputs only when both files and that exact canonical manifest agree. Repeating identical generation preserves file timestamps. Output paths reject parent traversal, symbolic links, nested module boundaries and the reserved `.goml-bind-go-lock` recovery and `.goml-bind-go-query` query directories. A module-level lock serializes generators. Publication stages both sources and the manifest, rechecks the previous contents, and rolls back committed sources if a later rename fails. If rollback itself fails, the lock and backup files remain for recovery; subsequent generation refuses to proceed until the failure is resolved. These checks protect generated files; handwritten files outside the configured output names are not publication targets.

Metadata queries use a temporary child package of the configured native output directory, so Go internal-import permissions follow that output location while stale generated wrappers remain separate from the query package. Existing query directories are preserved and reported as conflicts. The directory is removed after successful or failed queries. Native functions and types from the output package use unqualified references, with no self import. Discovery overlays only the owned generated file with a package placeholder, so old signatures cannot block regeneration. Before publication, the helper overlays the complete candidate and checks its real native package, including handwritten files, function signatures, import cycles and declaration conflicts. Type-only configurations also validate that native package. Source overlays never modify the existing file.
