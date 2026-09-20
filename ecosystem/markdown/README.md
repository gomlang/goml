# ecosystem::markdown

A CommonMark parser, public AST and configurable HTML renderer implemented in
GoML. The current implementation passes all **652 CommonMark 0.31.2 examples**
with exact HTML comparison. Parsing and rendering are implemented in this module;
two small Go standard-library bindings supply Unicode punctuation/symbol
classification, alongside `std::unicode` whitespace and case folding.

```gom
use ecosystem::markdown;

fn render_page(source: string) -> Result[string, markdown::Error] {
    markdown::to_html(source)
}
```

## Coverage

- Paragraphs, ATX and Setext headings, thematic breaks, fenced and indented code.
- Nested block quotes, lazy paragraph continuation, ordered and unordered lists,
  empty items, tight/loose rendering and tab-aware container indentation.
- All seven CommonMark HTML block forms, inline tags, comments, processing
  instructions, declarations and CDATA.
- Code spans, nested emphasis/strong emphasis with delimiter-run rules,
  escapes, soft/hard breaks and Unicode-aware delimiter classification.
- Inline, full/collapsed/shortcut reference links, images, URI/email autolinks,
  balanced destinations, multiline titles and first-definition-wins references.
- Case-folded, whitespace-normalized reference labels, all 2,125 semicolon-ended
  HTML5 named entities, numeric references and replacement of invalid scalar
  references/NUL input.

Input line endings are normalized from CRLF and lone CR to LF. Code preserves
literal interior tabs; tabs used for container structure obey four-column stops.
Link/image destinations are percent-encoded during rendering, while existing
percent escapes are preserved.

This module implements CommonMark, without GFM tables, task lists, strikethrough,
footnotes, math or syntax highlighting. Passing the complete reference corpus is
evidence of the tested behavior, not proof that every possible input matches
every other Markdown implementation.

## Parsing and AST

`parse(source)` uses `ParseOptions::standard()`. `parse_with_options` accepts
limits for input bytes, nesting, allocated parse nodes and parsing work. Defaults
are 4 MiB input, depth 128, 250,000 parse nodes and 20,000,000 work units.
Malformed Markdown normally remains text, as CommonMark requires; resource-limit
failures return `Error { message, line }`. Error lines are zero based internally
and one based in their string representation.

`Document` exposes `blocks` and its normalized `references` map. `Block` contains
a `BlockKind` and an inclusive, zero-based `LineSpan`; `Inline` describes the
inline tree. Public enum cases preserve heading levels, list start numbers and
tightness, fence information, literal code, link/image destinations and titles.
Applications can inspect or construct these values and pass the AST to `render`.
Source spans describe whole blocks, not byte-accurate inline token locations.

Parsing first builds block structure and collects reference definitions. A
second pass resolves inline trees, including references defined later in the
document or inside containers. Emphasis uses a delimiter stack and linked node
indices; bracket handling prevents nested links while supporting formatted image
descriptions.

`walk(document, visit_block, visit_inline)` visits nodes in source order with
parents before children. It uses an explicit stack and accepts captured
callbacks. `inline_text(inlines)` extracts decoded text, code content and line
breaks without formatting tags or HTML escaping; image/link labels contribute
their text. Walks have limits of 250,000 visited nodes and depth 256, and return
errors for manually constructed cyclic/deep ASTs. Callbacks must not mutate the
tree during traversal.

AST vectors and maps have normal GoML shared-storage semantics. Returning an AST
does not make its container storage immutable, and the module does not provide
concurrent mutation support.

## Rendering

`render(document, options)` returns HTML or a checked limit/AST error.
`to_html(source)` combines parsing with `RenderOptions::standard()`.

The default renderer escapes raw HTML and permits relative URLs plus `http`,
`https`, `mailto` and `ftp` schemes. Other schemes produce empty link/image
destinations. Scheme checking accounts for whitespace/control characters after
Markdown entity decoding. `escape_html` and `safe_url` are also public helpers.

`RenderOptions::commonmark()` preserves raw HTML and arbitrary URI schemes to
match the reference corpus. Applications select that policy explicitly. Render
options also control soft-break conversion and HTML versus XHTML void-element
spelling. Titles, alt text and code language attributes are escaped; formatted
image descriptions become plain alt text.

Default rendering limits are depth 256 and 16 MiB output. They are independent of
parse limits. Manually constructed ASTs can therefore be checked without going
through the parser. Heading levels outside 1–6 and exceeded limits return errors.

## Resource behavior

There are no recursive calls proportional to plain-text length or delimiter-run
length. Container parsing and finalization recurse on bounded nesting; inline
nesting is checked as nodes are grouped. Some unsuccessful delimiter searches
can be quadratic. Work accounting charges delimiter scans and conservative
remaining-input bounds for candidate links/HTML so those paths can terminate
with a recoverable error. This is an operation budget, not a wall-clock timeout.

Parsing allocates block text and inline nodes; it is not an incremental editor
parser or a zero-copy rope. Input limits and parse-node limits also constrain
large valid documents. Applications handling larger documents can set explicit
limits. Output checking limits emitted HTML; intermediate escaped strings still
allocate before they are written.

## Validation and data provenance

Run from the repository root:

```sh
python3 ecosystem/verify.py markdown
```

The command checks formatting, library tests, a separately resolved consumer,
fresh/cached consumer builds, its executable and reference interoperability.
Library tests cover AST construction/inspection, spans, captured visitors,
Unicode/reference normalization, nested lists, literal tabs, escaping, URI
policy, render options and resource/cycle errors. The independent consumer
imports only public APIs and also offers stdin conversion through `--safe`,
`--commonmark` and a JSON-array batch interface through `--json`.

`interop.py` runs every example from the [CommonMark 0.31.2 reference corpus](https://spec.commonmark.org/0.31.2/spec.json)
and fails on any mismatch. The corpus is downloaded into `_artifact/reference`
and checked against a fixed SHA-256 digest before use. Per-section counts and
every failure are recorded in `_artifact/verification/markdown/commonmark.json`.
The CommonMark specification and examples are by John MacFarlane, licensed under
[CC BY-SA 4.0](https://creativecommons.org/licenses/by-sa/4.0/).

The runner also checks all 2,125 named entities against Python's `html.entities`
table. `entities.gom` is generated data derived from that table in Python 3.12.3;
the transformation keeps semicolon-ended names and emits GoML lookup functions.
The source license is retained in [LICENSE.entities.txt](LICENSE.entities.txt).
Regenerate with `python3 ecosystem/markdown/generate_entities.py`, then run the
module's GoML formatter and verification command. Generated data is checked in;
Python is needed for verification/regeneration, not for library execution.
