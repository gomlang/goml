# GoML API documentation

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

Consecutive standalone `///` lines attach to the next declaration, including its attributes. A blank source line or ordinary comment interrupts attachment. `////` and same-line trailing comments are ordinary comments. Documentation may appear on functions, types, constants, statics, fields, variants, trait members, implementations, methods, and public re-exports.

`//!` is package documentation when it appears before the first source token. Descriptions from several package files are combined in relative path order. The formatter preserves documentation attachment. The lexer continues to treat these spellings as ordinary comments, so previous toolchains can compile sources containing them.

The repository's restriction on implementation comments remains applicable. Extraction tests keep documented source examples in string fixtures; this change does not add documentation comments throughout the compiler or libraries.

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

The implementation lives in `gomlc/doc_comment`, `gomlc/query/documentation*`, `gomlc/doc`, and `gomlc/cmd/gomldoc`. Project integration lives in `goml/commands/doc.gom` and `goml/cmd/goml/doc.gom`. Tests cover attachment, visibility, imports, output, errors, and CLI integration. Toolchain and packaging changes require `just ci`, including fixed-point and relocated archive checks.

Dependency documentation sites, search, an HTTP server, automatic browser launch, and doctests are future work.
