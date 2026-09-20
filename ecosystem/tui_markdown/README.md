# tui_markdown

CommonMark terminal layout and an interactive TUI viewer implemented in GoML. It composes the existing `markdown` AST, Unicode line breaking, `ansi` styles and `tui` buffers; it does not launch a browser or a pager subprocess.

```toml
[dependencies]
"ecosystem::tui_markdown" = "0.1.0"
"ecosystem::tui" = "0.1.0"
"ecosystem::ansi" = "0.1.0"
"ecosystem::terminal" = "0.1.0"
```

```gom
use ecosystem::tui_markdown;
use ecosystem::tui;
use ecosystem::ansi;

fn document(source: string) -> Result[string, tui_markdown::Error] {
    let result = tui_markdown::layout(source, tui_markdown::Options::defaults(80))?;
    result.render_ansi(ansi::Profile::TrueColor, false)
}

fn preview(source: string, buffer: tui::Buffer) -> Result[(), tui_markdown::Error] {
    let viewer = tui_markdown::Viewer::new(
        source,
        tui_markdown::Options::defaults(buffer.area().width),
    )?;
    viewer.draw(buffer, buffer.area())
}
```

## Layout and styling

`layout(source, options)` parses CommonMark and returns `Layout { lines, links, width }`. `render(document, options)` accepts an existing or manually constructed `markdown::Document`. `plain_text()` produces a screenshot; `render_ansi(profile, hyperlinks)` produces styled lines separated by LF, without a trailing LF. Explicit hyperlinks are emitted only for a non-plain profile when `hyperlinks` is true. `Layout.lines` contains `ansi::StyledText`, so other TUI widgets and output adapters can consume the result.

Supported terminal presentations include:

- ATX/Setext headings with level markers and hanging continuation indentation.
- Paragraphs, nested emphasis/strong emphasis, code spans, hard/soft breaks and thematic separators.
- Unicode-aware word wrapping that preserves graphemes and styles, including CJK and emoji; narrow columns never emit half a wide glyph.
- Fenced and indented code with language labels, preserved indentation and literal column wrapping. Syntax highlighting is not performed.
- Nested block quotes and tight/loose ordered/unordered lists, preserving list start numbers and continuation indentation.
- Inline/reference/autolinks, images represented by alt-text labels, optional numbered links and optional URL footnotes.
- Raw HTML displayed as sanitized literal text. HTML is never executed or interpreted as terminal markup.

`Theme` supplies text, heading, code, quote, link, selected-link, search, border and muted styles. Inline emphasis combines with its surrounding style. `Options.link_numbers` defaults to true; `link_footnotes` defaults to false. Hyperlink metadata remains on spans even when not emitted, allowing the viewer to select links. TUI cells retain the visible styles and text, while `Layout.render_ansi` can retain OSC 8 hyperlinks for stream output.

`Link.index` is a one-based displayed number. Each link stores its decoded URL, title, plain label and first displayed line. Repeated destinations remain separate links. Safe destinations use the Markdown library's relative/http/https/mailto/ftp policy and additionally reject empty URLs, terminal controls and URLs longer than 4,096 bytes. Unsafe destinations retain their labels but receive neither a selectable link nor OSC 8 metadata. `safe_destination` exposes that policy. The library never opens a link itself.

Line wrapping uses Unicode line-break opportunities, trims paragraph wrapping spaces, and breaks long words when necessary. Code blocks preserve spaces and wrap by columns. Width policies follow the Unicode module defaults. Rendering is left-to-right terminal layout; bidirectional shaping and font-specific width negotiation are not implemented.

## Optional pipe tables

Set `Options { pipe_tables: true, ..Options::defaults(width) }` to enable a deliberately limited table extension in `layout(source, options)`. It recognizes a top-level paragraph whose first row and delimiter row contain pipes. Header/body rows must have equal cell counts; delimiter cells contain at least three hyphens with optional leading/trailing colons for left/center/right alignment.

Tables support optional outer pipes, escaped `\|`, inline Markdown inside cells, Unicode cell wrapping, aligned padding and visible borders. Columns divide available width evenly, with earlier columns receiving leftover columns. Headers use the heading style. A table may contain at most 32 columns, 256 body rows and 2,048 total header/body cells. If the viewport is too narrow for one column per cell plus borders, the text falls back to ordinary CommonMark paragraph layout.

This extension is not full GFM. It does not recognize tables nested inside lists/quotes, multiline cells, merged cells, pipes inside code spans, or reference definitions inherited from outside a cell. Malformed row shapes fall back to ordinary paragraphs; recognized tables exceeding resource limits return an error. `render(document, options)` cannot recognize this extension because a CommonMark AST does not retain the required original pipe-delimiter source; use `layout` for tables. Task lists, strikethrough, footnotes and math are not added by this module.

## Viewer

`Viewer::new(source, options)` owns parsed/layout state, a viewport and selection/search state. `draw(buffer, area)` reflows on width changes, clears the target area, clips visible lines and highlights selected links and search matches. A width change keeps the selected link visible. Zero-size rectangles draw nothing. Use one viewer across frames to retain its state.

- `update(event)` handles Up/Down, PageUp/PageDown, Home/End, mouse-wheel scrolling, Tab/BackTab link cycling, and `n`/`N` search navigation. Its boolean indicates that the event was handled. The application owns focus and activation actions.
- `scroll_to(line)`, `scroll_offset()` and `line_count()` expose viewport control. Scrolling clamps to the rendered document and the last drawn height.
- `select_link(index)` accepts a **zero-based vector index** and scrolls that link into view. `selected_link()` returns its public `Link`, whose displayed `index` is one-based.
- `search(query)` performs literal, case-sensitive matching within displayed lines, returns the match count and selects the first match. Empty text clears the query. Search queries are at most 4,096 bytes and cannot contain LF. `next_match(reverse)` cycles; `search_matches()` returns byte ranges into displayed line text. Searches do not span wrapped line boundaries.
- `set_source(source)` replaces the document atomically after layout and search validation, then resets viewport and link selection. Failed replacements leave the existing viewer intact.
- `snapshot()` copies the outer layout, each line's span storage and link storage. Editing a snapshot does not change the viewer.

Link selection styling takes precedence over search styling on the same grapheme. Search offsets are byte offsets, while highlighting paints whole overlapping graphemes. After a width change the current query is rerun; the active search-match index resets. The viewer and its shared containers are intended for one event-loop thread and do not synchronize concurrent mutation.

The viewer does not own a terminal session. Draw it inside `tui::Terminal::draw`/`draw_with` and route input from that same terminal's `next_event`; close the session with a `defer`. The independent consumer includes a complete `--pty` event loop using this arrangement.

## Limits and verification

Default limits are width 1–4,096 columns, 1 MiB source/AST text, 4 MiB layout text, 65,536 lines, depth 64, 100,000 traversal nodes and 4,096 links. Options validate upper bounds (4 MiB input, 16 MiB output, 262,144 lines, depth 128, 250,000 nodes, 16,384 links). Search retains at most 16,384 matches. Copied link labels/URLs/titles share a separate metadata budget equal to `max_output_bytes`, preventing manually nested links from multiplying metadata without a limit. ANSI stream rendering also enforces a 16 MiB output limit.

A preflight traversal checks each child vector length before descending, counts scalar text/URL/title bytes and nodes, and rejects cyclic/deep manually constructed ASTs without bulk-copying child vectors onto an unbounded traversal stack. Formatting uses separate output/traversal limits. Errors contain a recoverable message. Callbacks and concurrent mutation of supplied ASTs are outside the traversal contract.

Tests cover block snapshots, Unicode/style boundaries, safe links and literal HTML, optional table escaping and alignment, cycles and resource limits, viewport/search behavior, snapshot isolation and selected-link visibility after resizing. Independent consumers exercise public composition with TUI. `interop.py` checks 1,110 successful cases against a separate Python layout model plus eight resource-limit rejection cases. `pty_test.py` exercises Unicode columns, scrolling, code, tables, quotes, link selection, reflow and terminal restoration through a real Linux PTY.

```sh
python3 ecosystem/verify.py tui_markdown
python3 ecosystem/tui_markdown/interop.py
python3 ecosystem/tui_markdown/pty_test.py
```
