# highlight

Native GoML lexical syntax highlighting with reusable `logos` grammars, immutable
document snapshots, cross-line state, incremental edits, themes and ANSI/HTML
rendering. The design takes inspiration from [syntect](https://github.com/trishume/syntect).

```goml
use ecosystem::highlight;
use ecosystem::ansi;
use std::context;

fn preview(source: string) -> Result[string, highlight::Error] {
    let engine = highlight::Highlighter::new(highlight::Options::new())?;
    let document = engine.document("goml", source, context::Context::background())?;
    document.to_ansi(highlight::Theme::dark(), ansi::Profile::TrueColor, 1048576)
}
```

## Languages and scopes

The built-in registry includes `goml`, `json`, `toml`, `markdown` and `text`.
`detect(filename)` uses registered extensions; language selection also accepts
an extension such as `gom` or `md`. Unknown explicit languages return errors;
unknown Markdown fence languages render as plain text.

- GoML: keywords, primitive/container types, Unicode identifiers, numbers,
  loop labels, punctuation/operators, line comments, nested block comments, escaped strings,
  character/byte literals and arbitrary hash-delimited multiline raw strings.
  Interpolated strings receive string styling as a whole.
- JSON: literals, numbers, escaped strings, punctuation and invalid bare words.
- TOML: keys, numbers/dates, literals, comments, punctuation and basic/literal
  strings including multiline forms. Dates share the number scope.
- Markdown: ATX headings, inline code, basic emphasis/links, HTML comments and
  fenced code with GoML/JSON/TOML/custom embedded languages. Fence closure takes
  precedence over an embedded language's unfinished string or comment.

Highlighting is tolerant lexical classification, not syntax validation. Markdown
inline highlighting is deliberately smaller than the separate CommonMark parser:
nested emphasis/link parsing, setext headings, indented code and block-container
fence indentation are not modeled. TextMate/Sublime grammar compatibility and
semantic symbol classification are not provided.

`Language::compile(name, extensions, rules, regions)` compiles literal/regex
rules through `logos`. `Region` describes opening/closing delimiters, nested
comments, optional escape characters, and multiline policy. Regex syntax and
longest-match/priority semantics are those of `logos`; ambiguities and work
limits become recoverable errors. Regions use priority 200; literal rules use
100, regex rules 10 and the whole-scalar plain fallback 0. Users can choose a
different rule priority. `with_language` returns a new registry, replacing an
existing language of the same name. Original registries remain unchanged.

## Documents, state and incremental edits

`document(language, source, context)` retains a complete source snapshot and
per-line tokens/state. `line(language, text, before, context)` supports incremental
streaming by returning a line and its `after()` state. One call accepts at most
one newline, which must be final. State must come from the same language/registry.
Single-line strings recover at the line end; multiline regions remain open.

`Document::spans()` returns complete, sorted, non-overlapping UTF-8 byte ranges;
`Line::spans()` returns ranges relative to that line. All nonempty source bytes
are covered. Returned vectors are detached. Lines preserve LF/CRLF and a final
empty line after a trailing newline; an empty document has one empty line.

`replace(start, end, replacement, context)` checks UTF-8 boundaries and returns
`Update { document, reparsed_lines, reused_lines }`. It retains unchanged prefix
lines, reparses from the first changed line, and reuses the unchanged suffix as
soon as the incoming lexical state agrees. Old snapshots remain immutable and
can be used concurrently. Line splitting and source reconstruction are O(total
bytes); incremental reuse saves lexical work, not the source-copy cost. A rope
consumer demonstrates keeping document edits and source edits synchronized.

## Rendering and limits

`Theme::dark`, `light`, and persistent `with_style` map scopes to `ansi::Style`.
`to_ansi` supports plain, 16-color, 256-color and truecolor output, emits style
transitions and a final reset, and sanitizes terminal control characters through
`ansi::safe_text`. Consequently raw CR/control bytes are replaced in terminal
output. `to_html` escapes source and emits stable `hl-<scope>` CSS classes inside
`<pre><code>`; the application supplies its stylesheet. Both renderers have a
checked output byte limit.

Default limits: 4 MiB/document, 64 KiB/line, 100,000 lines, 1,000,000 spans and
64 nested region levels, plus `logos` automaton/matching budgets. Definition and
edit failures return `Error { message, line }`; parsing reports one-based line
numbers. Cancellation/deadlines are checked during line splitting and reuse,
between lines/tokens and periodically within regions, regardless of UTF-8 widths.
User grammar compilation and individual bounded lexer
calls are synchronous. Failed edits leave the previous document intact.

## Validation

Run `just ecosystem-test highlight`. Native GoML tests cover built-in/custom
grammars, Unicode-safe ranges, nested/raw/multiline states, embedded fences,
recovery, persistent edits and suffix reuse, output escaping/bounds, budgets,
cancellation, detached snapshots and concurrent registry reuse. Two hundred
deterministic edits compare incremental results with full re-highlighting.
The versioned consumer combines `rope`, Markdown embedded code, and plain/HTML
preview. There is no Python or native adapter.
