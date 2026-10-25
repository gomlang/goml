# GoML API documentation

[English](documentation.md) | [简体中文](documentation.zh-CN.md)

`gomldoc` generates documentation from checked GoML source. It shares package analysis, dependency navigation, derive expansion, and type checking with the compiler query engine. Documentation generation does not build a native executable. Compile-time evaluation and Go FFI validation still follow normal analysis rules.

## Commands

From any directory inside a module:

```sh
goml doc
goml doc --format json
goml doc --document-private-items
goml doc --target-dir build --dry-run
```

The driver prepares dependencies and invokes `gomldoc` for the enclosing module. The default output is `<target-dir>/doc`, where the target directory comes from `goml.toml` or `--target-dir`. Production packages are included; tests, `testdata`, hidden directories, the build target directory, and nested modules are excluded.

The documenter is found through `--documenter`, `GOMLDOC`, a sibling of `goml`, `GOML_HOME/bin`, then `PATH`. `--compiler` selects the compiler used by the driver's project discovery.

The standalone command uses locally available dependency sources:

```text
gomldoc [OPTIONS] [MODULE_DIR]

--format html|json        default: html
--output DIR              default: <target-dir>/doc
--target-dir DIR          override the module build target directory
--document-private-items  include private declarations and members
-h, --help
--version
```

`MODULE_DIR` defaults to the current directory; the enclosing manifest determines the module. An explicit relative output path is relative to the working directory. A relative `--target-dir` is relative to the module root and also excludes that build directory from discovery. If dependencies are missing, prepare them with `goml update` or use `goml doc`.

The standalone command exits 0 on success, 1 on analysis or output failure, and 2 on invalid arguments. Link warnings do not make generation fail. Output is written only after successful analysis and rendering. A failed run preserves the previous documentation. Existing output directories must have the `.gomldoc` ownership marker; arbitrary existing directories are not replaced. Generated files belong to the documenter and are replaced on the next successful run.

## Comments

```goml
//! Package overview.
package example;

/// A named value.
///
/// The public field can be read by callers.
pub struct Value {
    /// Display name.
    pub name: string,
}
```

See [comment syntax and attachment](language/basics.md#comments-and-api-documentation) for `///` and `//!` rules. Documentation may appear on functions, types, constants, statics, fields, variants, trait members, implementations, methods, and public re-exports. The [formatter preserves attachment](formatting.md#comments-and-literals).

## Rendering and links

The supported Markdown subset includes paragraphs, ATX headings, flat `-` or `*` bullet lists, single-backtick inline code, triple-backtick fenced code blocks, and links. Raw HTML is escaped. Tables, emphasis, nested lists, images, and full CommonMark semantics are not implemented.

`[Value]`, ``[`Value`]``, and `[a value](goml:models::Value)` are API references. Names are interpreted in the declaring file's package and import context, including explicitly renamed imports and a package's declared default alias. Canonical package paths and `module::` paths are supported. Resolution does not guess by a global name suffix. Unresolved or ambiguous references produce diagnostics. A known target without a generated page, such as an external dependency, is rendered as text.

Ordinary link destinations accept `https:`, `http:`, `mailto:`, and local fragments. Other schemes are rendered as text. All HTML and CSS are local; opening `index.html` works offline.

Public API output excludes private declarations, private fields, private inherent methods, and implementations belonging to private types or traits. Trait members inherit the trait's visibility. Generic parameters, bounds, `where` clauses, aliases, and associated types remain in signatures. Method bodies and constant/static initializers are omitted. Re-exports retain the original target, and generated APIs are marked as generated.

## Output

HTML output contains a module index, a page for each package, dedicated struct/enum/trait pages, and a stylesheet. Symbol identifiers and page names use canonical identities, with implementation signatures distinguishing trait implementations. Relative source filenames and one-based source lines are displayed; generated declarations use line 0.

JSON output is `module.json`:

```json
{"schema_version":1,"module":"example","packages":[]}
```

Each package contains `name`, combined `documentation`, `sections`, and `items`. Each section retains its relative `source`, `documentation`, and resolved `links`, so package descriptions from different files keep their own import contexts. Each item contains `id`, `name`, `qualified_name`, `kind`, `public`, `signature`, `documentation`, `source`, `line`, `generated`, `target`, `links`, and recursively nested `members`. Links contain their reference `text` and canonical `target`. Consumers must check `schema_version`; this schema is separate from compiler artifact formats.

Output ordering is deterministic. Generated contents contain neither timestamps nor absolute source paths. HTML and JSON are projections of the same documentation model.

## Development

`just make` builds `stage2/bin/gomldoc`; `just install` and release archives include it. It uses the compiler version module and executable-relative library resources.

The implementation lives in `gomlc/doc_comment`, `gomlc/query/documentation*`, `gomlc/doc`, and `gomlc/cmd/gomldoc`. Project integration lives in `goml/commands/doc.goml` and `goml/cmd/goml/doc.goml`. Tests cover attachment, visibility, imports, output, errors, and CLI integration. Toolchain and packaging changes require `just ci`, including fixed-point and relocated archive checks.

Dependency documentation sites, search, an HTTP server, automatic browser launch, and doctests are future work.

## Source API snapshots

`goml api snapshot` checks the module and writes a versioned JSON snapshot to standard output. It reuses the normal dependency discovery and interface artifact pipeline; it may create check artifacts and requires the selected compiler. Build progress and diagnostics stay on standard error. Redirect standard output to save a baseline:

```sh
goml api snapshot > baseline.json
goml api snapshot --target-dir build > current.json
goml api diff baseline.json current.json
```

`--compiler PROGRAM` selects the compiler for either operation. `api diff` does not rebuild the module. It writes a JSON report and exits 0 when there are no breaking or potentially breaking changes, 1 when review is required, and 2 when the compiler cannot read or validate the snapshots.

Snapshot schema version 1 records the module, public paths, canonical declarations, structured types, generic parameters and predicates, public fields, enum variants, trait members and implementations, evaluated constant values, and public derive entry points. Re-exports preserve both the consumer path and the target declaration. Object keys and unordered declarations are sorted; private bodies, documentation, source positions, artifact hashes, and function parameter names are excluded. Module-internal declarations are excluded. Inherent method paths include their receiver and specialization arguments; dependency type and trait contracts reachable from public signatures are retained separately in `referenced_contracts`. Dependency interfaces are used to resolve re-export contracts; unavailable targets appear in `unresolved_reexports`, and comparing such snapshots requires review.

The report distinguishes removed public paths or contract members (`breaking`), changed contracts or declaration identities (`potential`), and new public paths (`added`). A preserved facade with a moved target is therefore reported as a declaration change, not silently accepted as identity-preserving. Generic parameter renaming and equivalent but differently expressed constraints may still require review. Adding enum variants or trait members can affect exhaustive matches or downstream implementations and is conservatively reported.

This checks source API structure, not runtime behavior or binary ABI. Constant initializer behavior, macro expansion behavior, platform-dependent foreign ABI compatibility, and all possible downstream inference interactions are not proven compatible. Current interface artifacts do not independently retain whether a trait method first parameter was spelled `self`; that source syntax distinction is outside version 1 coverage. Unrepresentable dynamic compile-time constants make snapshot generation fail instead of producing an incomplete constant contract. Validate important facade migrations with an independently compiled consumer module as well as the snapshot report.
