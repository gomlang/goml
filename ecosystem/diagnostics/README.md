# diagnostics

`ecosystem::diagnostics` is a pure GoML source diagnostic library. It provides reusable source caches, checked byte spans, multi-file and overlapping labels, Unicode-aware snippets, configurable plain/ANSI rendering, and validated multi-file suggestions. It performs no I/O, terminal probing, compiler integration, or automatic source modification.

```toml
[dependencies]
"ecosystem::diagnostics" = "0.1.0"
```

```gom
use ecosystem::diagnostics;

fn example() -> Result[string, diagnostics::Error] {
    let cache = diagnostics::SourceCache::new();
    let file = cache.add("main.gom", "let value = old;\n")?;
    let span = cache.span(file, 12, 15)?;
    let fix = diagnostics::Suggestion::new(
        "use the new name",
        diagnostics::Applicability::MachineApplicable,
    ).with_edit(diagnostics::Edit::new(span, "new"));
    let report = diagnostics::Diagnostic::new(
        diagnostics::Severity::Error,
        "unknown name",
    ).with_code("E001")
        .with_label(diagnostics::Label::primary(span, "not in scope"))
        .with_help("declare the value before using it")
        .with_suggestion(fix);
    diagnostics::render(cache, report, diagnostics::RenderOptions::new())
}
```

The default renderer produces ordinary strings. Pass the result to your application's chosen output writer. ANSI output is explicit: set `RenderOptions.profile` to `ansi::Profile::{Ansi16,Ansi256,TrueColor}` and add `ecosystem::ansi` as a direct dependency when referring to its types. There is no environment-variable or terminal state dependency.

## Sources and positions

`SourceCache::new()` creates an append-only cache. `add(name, text)` returns an opaque `FileId`; duplicate names are allowed and represent independent snapshots. IDs include cache identity and are rejected by other caches even when their numeric registration positions coincide. Source content is immutable after insertion, so previously created spans cannot become stale through replacement. Add a new source for a revision; the application decides how to retain or discard caches. `len()` and `total_bytes()` expose cache size. Cache insertion is not synchronized; coordinate mutations in the application.

`get(file)` returns a reusable `Source` snapshot with `id()`, `name()`, `text()`, `byte_len()`, `line_count()`, and `line(number)`. `SourceLine` exposes zero-based `number`, inclusive byte `start`, exclusive content `end`, byte `next` after the line separator, and separator-free `text`. LF, CRLF, and lone CR are physical separators. Empty input has one line, and a trailing separator creates a final empty line. Other Unicode separators are rendered visibly as replacement characters rather than splitting physical lines.

`cache.span(file, start, end)` and `source.span(start, end)` validate an exclusive UTF-8 byte range and both scalar boundaries. Empty spans and EOF are valid. A span can select part of a grapheme cluster when its boundaries are valid scalar boundaries. `Span::{file,start,end,is_empty}` inspect it, and `cache.text(span)` returns the exact original text.

`source.location(offset, WidthOptions, tab_stop)` returns zero-based `Location::{line,byte_column,display_column}`. Offsets inside CRLF clamp to the preceding line's content end. An offset inside a grapheme maps to that grapheme's starting terminal column; label ends expand to its ending column. Tabs advance to the next configured stop. CJK, combining sequences, emoji ZWJ sequences, and optional ambiguous/emoji width policies use `ecosystem::unicode_text`. Controls and bidi direction marks/overrides/isolates render as `�`. Isolated zero-width clusters gain a dotted circle, so they cannot attach to the snippet gutter. `tab_stop` must be 1–256.

`SourceLimits::new()` defaults to 4,096 files, 16 MiB per file, 64 MiB total text, and 1,000,000 lines per file. `SourceCache::with_limits` checks custom bounds before construction. Hard maxima are 100,000 files, 64 MiB per file, 256 MiB total text, and 1,000,000 lines per file; source names are limited to 64 KiB. Rejected insertion leaves the cache unchanged. The byte budget counts source text; line indexes and names incur additional input-proportional storage.

## Diagnostic model

`Diagnostic::new(severity, message)` supports `Severity::{Error,Warning,Advice,Note}` and composable `with_code`, `with_label`, `with_note`, `with_help`, and `with_suggestion` methods. Labels use `Label::primary(span, message)` or `Label::secondary(span, message)`. The report, label, note, suggestion, and edit records expose their data for serialization or application-specific presentation. No formatting or source lookup occurs in the builders.

Builders copy the collection they extend, so extending a report does not add labels to an earlier report. Public vectors retain normal GoML shared-vector semantics; callers must coordinate direct mutation, and nested vectors are not deeply cloned. A renderer can be reused sequentially for many independent reports.

## Rendering and layout

Call `render(cache, report, options)` or construct `Renderer::new(options)` and call `renderer.render(cache, report)`. Both return `Result[string, Error]`. Invalid cache identities, limits, spans, options, and suggestion conflicts are recoverable errors. A failed render returns no partial output.

Files appear in source registration order. Within a file, labels sort by byte start, primary before secondary at the same start, and original insertion order as the final tie-breaker. The first primary label supplies the file header location; a file with only secondary labels uses its first label. Headers show one-based display columns and line numbers.

Each overlapping label has its own marker row, avoiding ambiguous shared caret messages. Primary markers use `^`; secondary markers use `-`. Multiline labels show `[begins]`, `[continues]`, and `[ends]` on visible lines. Empty, EOF, and newline-only spans retain a visible marker. Context windows around the beginning and end of every label are merged and deduplicated. Large multiline interiors are omitted with a count instead of printing unlimited source text.

Horizontal clipping centers a window near the first label on a line. `<` and `>` indicate hidden source text or an offscreen label. A wide grapheme partially intersecting a window is replaced by spaces at that boundary; its bytes are never split. Long marker rows move the label message to a separate row. Messages, names, notes, and replacement previews are clipped with `...` to the configured terminal width; metadata newlines/tabs are shown as `\n`/`\t`, and terminal controls are neutralized. This is a compact snippet renderer rather than a paragraph reflow engine.

The source-line budget is shared across all files. When exceeded, the renderer retains the first half and last half of selected lines in deterministic source order. Gaps and trailing omissions are explicit. Files with no retained excerpt still show their header and omitted-label summaries. A wholly hidden label retains a summary with its line range and message. Label rows and metadata rows are bounded by their separate counts and the final output byte limit.

`Theme` provides independent ANSI styles for errors, warnings, advice, notes, secondary labels, gutters/headings, and help. `Profile::Plain` ignores styles. The supplied colors are downsampled by `ecosystem::ansi` to the requested output profile, and every styled fragment restores default terminal attributes.

| `RenderOptions` field | Default | Accepted range |
| --- | --- | --- |
| `columns` | 100 | 24–4,096 |
| `tab_stop` | 4 | 1–256 |
| `context_lines` | 2 | 0–32 |
| `max_source_lines` | 64 | 1–4,096 |
| `max_labels` | 64 | 0–1,024 |
| `max_notes` | 64 | 0–1,024 |
| `max_suggestions` | 16 | 0–1,024 |
| `max_edits` | 64 total across suggestions | 0–10,000 |
| `max_text_bytes` | 64 KiB per message/code/label/note/replacement | 0–1 MiB |
| `max_line_bytes` | 1 MiB per measured source line | 0–64 MiB |
| `max_output_bytes` | 1 MiB, including ANSI bytes | 0–64 MiB |

`profile`, `theme`, and `width` are also configurable. Cache names use their separate source limit. The line-byte bound applies to excerpt lines and file/suggestion anchor lines before grapheme layout. Plain `Source.location` is bounded by the cache's file limit instead. Repeated layout is linear in measured source bytes; this initial API does not memoize display columns or shape fonts. It inherits the deterministic terminal-width policy from `unicode_text`; it performs no bidi reordering or font-dependent measurement.

## Suggestions and edit application

`Suggestion::new(message, applicability).with_edit(Edit::new(span, replacement))` models one concrete, possibly multi-file change. `Applicability` records `MachineApplicable`, `MaybeIncorrect`, `HasPlaceholders`, or `Unspecified`; it is descriptive, and the application decides whether to offer or apply the suggestion.

`validate(cache, EditLimits)` checks every file, range, replacement budget, edit conflict, and the combined final output size before any changed text is constructed. `apply(cache, limits)` performs the same preflight and returns `Vec[ChangedSource]` in source registration order. Each result includes `file`, `name`, `original`, and `updated`. The cache is never changed, and no filesystem writes occur. Separate suggestions are alternatives and are validated independently; merge their edits into one suggestion if they must apply together.

All edit offsets refer to the original snapshots. Nonempty overlapping ranges are rejected, as are an insertion strictly inside a replacement and two insertions at the same byte offset. Adjacent replacements are allowed. An insertion at a replacement's start occurs before it; one at its end occurs after it. Sorting by `(file, start, end)` gives these boundary edits deterministic behavior. Conflict errors retain the original edit indexes. Replacement strings may contain arbitrary valid Unicode, including newlines; sanitization affects the displayed preview only.

`EditLimits::new()` permits 1,024 edits, 16 MiB combined replacement text, and 64 MiB combined resulting text across changed files. Hard maxima are 10,000 edits, 64 MiB replacement text, and 256 MiB output. Empty suggestions return an empty change list. Limits count final output, so an insertion followed by a deletion can fit a budget smaller than the original file. The renderer also checks suggestion validity and imposes its own total edit and preview bounds.

## Validation

```sh
GOML_BUILD_JOBS=2 just ecosystem-test diagnostics
```

The suite covers 22 library tests and 4 independent consumer tests: cache identity, malformed spans, CRLF/CR/EOF, combining and emoji boundaries, tabs, overlapping and multiline labels, exact plain layouts, ANSI equivalence, custom themes, horizontal/vertical clipping, output limits, control injection, edit conflicts, final-size preflight, and immutable source revisions.

The consumer’s ordinary GoML test replays [independent reference vectors and native render properties](../consumers/diagnostics/tests/data/README.md) for 6,236 deterministic model checks: 2,976 locations, 1,440 byte spans, 1,500 edit batches, and 320 plain/ANSI rendering checks. No Python runtime is required. The source-position model covers a controlled corpus of independently specified graphemes; the underlying Unicode library owns full Unicode conformance tests. The diagnostics library is not an implementation of Ariadne's API or output format.
