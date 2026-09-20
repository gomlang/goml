# gomlc

`gomlc` is the complete self-hosted GoML compiler and language server. It implements:

```text
lexer → parser → CST → AST → HIR → TAST → Core → Mono → Lift → ANF → Go
```

Generated Go code and Go FFI target Go 1.26. Building the compiler and its generated programs requires Go 1.26 or newer.

Run repository recipes from the repository root. On Linux amd64, a fresh checkout downloads the checksum-pinned binary stage0 and uses it to build the stage2 toolchain directly:

```sh
just make
```

Use `just bootstrap` for a clean bootstrap and stage3 artifact fixed-point verification. The development tools include:

```text
stage2/bin/gomlc
stage2/bin/gomlfmt
stage2/bin/gomllsp
stage2/bin/goml
stage2/bin/goml-go-meta
```

Each installed stage is a complete toolchain prefix. The compiler resolves `lib` from its executable's location. Module commands read the finalized compiler world at `lib/compiler/compiler-world-v2.gaf`; the prefix also carries the builtin, prelude, and standard-library projects.

Run a single source or inspect an IR stage:

```sh
stage2/bin/gomlc run-single file.gom
stage2/bin/gomlc anf file.gom
stage2/bin/gomlc run-single --dump-go file.gom
```

The regression corpus and every generated golden file live in `gomlc/testdata`. Verify or update them with:

```sh
just verify-golden
just update-golden
```

Run all self-hosted compiler, pipeline, query, and language-server tests with
`just test`. Run `just ci` for the complete repository checks, including fixed-point and packaging verification.

See the [language guide](../docs/goml.md), [formatter rules](../docs/formatting.md), and [compile-time evaluation architecture](../docs/comptime.md) for the corresponding compiler contracts.
