# tui

A GoML terminal UI library with Unicode cells, bounded layout, retained frame comparison, interactive widgets and a grapheme-aware editor. The terminal integration uses `ecosystem::terminal::Session` as the input/output owner; applications retain control of their event loop.

```toml
[dependencies]
"ecosystem::tui" = "0.1.0"
"ecosystem::ansi" = "0.1.0"
"ecosystem::terminal" = "0.1.0"
```

```gom
use ecosystem::tui;
use ecosystem::ansi;

fn render(buffer: tui::Buffer) -> Result[(), tui::Error] {
    let theme = tui::Theme::defaults();
    let inner = tui::block(buffer, buffer.area(), "Status", theme)?;
    tui::paragraph(
        buffer,
        inner,
        ansi::StyledText::plain("GoML 界面 👩‍💻"),
        tui::ParagraphOptions::defaults(),
    )
}
```

## Frames, text and layout

- `Rect::new(x, y, width, height)` checks nonnegative coordinates and dimensions. `intersection`, `inset`, `contains`, `right` and `bottom` provide geometry helpers. Because fields are public, drawing and layout entry points validate literal rectangles too.
- `Buffer::new(width, height)` starts with default-style spaces. `get` exposes an immutable `Cell` value; `put` accepts one grapheme and returns its column width, or zero when outside the buffer. `fill` blanks a rectangle; `draw_text` draws an `ansi::StyledText` with multiline clipping and eight-column tab stops. It does not wrap. `snapshot` copies cells; `lines` returns a textual screenshot; `validate` checks wide-cell invariants.
- A two-column grapheme owns a head cell and a continuation. Overwriting either half erases the old pair. A clipped glyph never leaves half a glyph. Repair may blank the adjacent half just outside a fill rectangle, which is necessary to keep the buffer coherent. Width-zero standalone clusters receive a dotted-circle base. Controls are sanitized; cells cannot inject terminal commands. Styles follow the first span of a grapheme spanning multiple spans. Hyperlink metadata is not stored in cells.
- `paragraph` handles styled wrapping or explicit lines, left/center/right alignment, and horizontal/vertical scrolling. Wrapping preserves graphemes; it is column-based rather than language-specific word wrapping. Its inputs follow ANSI literal text rules: LF and tabs survive; other terminal controls are replaced. The editor additionally normalizes CRLF before sanitizing.
- `layout(area, direction, constraints, spacing)` returns ordered rectangles. `Fixed(n)` and `Ratio(n,d)` receive their requested space first, in declaration order; `Min(n)` receives a preferred minimum and participates in remaining-space allocation; `Max(n)` grows only up to its cap; `Fill(weight)` shares remaining space. Overconstrained preferred sizes are truncated in declaration order. Weighted growth and leftover columns are deterministic, favoring earlier entries for small remainders. Gaps are reserved before content. Fixed-only layouts may leave trailing space.

## Rendering and terminal ownership

`Renderer::new(width, height, profile)` compares complete buffers with its last successfully written snapshot. Each changed head cell emits an absolute cursor move, a style transition and its grapheme. Continuations are skipped. A frame begins and ends with an SGR reset. Unchanged frames produce no writes. `FrameStats` reports changed head cells, output bytes and whether the frame was a full redraw.

`render(buffer, writer)` accepts a fallible output callback. It snapshots before calling user code, commits only after a successful write, invalidates after write failure, preserves invalidation performed inside a callback, and rejects reentrant rendering. `resize` invalidates when dimensions change; `invalidate` requests a repaint after outside terminal output. `MemoryBackend` records writes, exposes `output`, and supports `fail_once` and `clear` for deterministic testing.

`Terminal::new(session)` checks interactive capabilities and selects the ANSI color profile. `draw` queries size, prepares a blank frame, calls the rendering callback, then writes the diff. `draw_with(context, callback)` also makes writes cancellable, including under output backpressure. `next_event(context, timeout_ms)` forwards keyboard, paste, mouse, focus, resize and EOF events; resize invalidates the frame. `close` restores the session. Use a `defer` to close the session even when constructing the TUI or rendering fails:

```gom
use ecosystem::terminal;
use ecosystem::tui;

fn run(session: terminal::Session) -> Result[(), tui::Error] {
    defer {
        let _ = session.close();
    };
    let screen = tui::Terminal::new(session)?;
    screen.draw(|buffer| render(buffer))?;
    screen.close()
}
```

The caller drives redraw timing and owns session lifetime. Buffers, editor state, renderer state and the memory backend are shared mutable GoML values intended for a single event-loop thread; they are not synchronized collections. Session operations retain the terminal library's synchronization. Do not write through another output owner while a frame is being committed. `Profile::Plain` disables colors but a TUI still emits cursor and reset controls. Interactive rendering requires a terminal; textual snapshots provide a noninteractive representation.

## Widgets and interaction

| Widget/state | Behavior |
| --- | --- |
| `Theme`, `block` | Base/border/title/selected/muted/accent styles, Unicode border and clipped title, inner content rectangle |
| `paragraph` | Styled text, column wrapping, alignment, vertical and horizontal scroll |
| `SelectionState`, `list` | Selection, focus, viewport following, arrow/Home/End/Page navigation and mouse wheel |
| `table` | Styled header/cells, constrained column widths, selected rows and vertical scrolling |
| `TreeNode`, `TreeState`, `tree` | Validated flat preorder hierarchy, unique IDs, expansion, depth indentation, selection and scrolling; Enter/Space/Left/Right expansion |
| `tabs` | Selected tab styling and horizontal scrolling to keep selection visible |
| `scrollbar` | Horizontal/vertical proportional thumb with clamped offset |
| `gauge` | Checked fraction, fill and centered label |
| `sparkline` | Eight-level Unicode bars, most recent values clipped to width |
| `Series`, `chart` | Integer-scaled axes, multiple styled series, point clipping and connected segments with a raster work budget |
| `Focus` | Bounded focus ring, Tab/BackTab cycling |
| `Editor` | Grapheme-aware text editing, selection, history, password view and cursor-following viewport |

List/table/tree/tabs rendering returns updated state; retain that value for subsequent events. `TreeState::with_selection` updates its public selection while preserving the private expansion state. Widgets clear the areas they own where appropriate. Chart and sparkline can be overlaid; the caller can fill their area before rendering. Widget operations may have drawn part of their area before returning an input/budget error, so render into a new frame and propagate errors instead of committing a failed frame. Mouse click hit-testing and application actions belong to the application; selection state directly supports wheel navigation.

## Editor

`Editor::new(text, max_bytes)` creates shared editor state. `text`, `cursor` and `selection` expose text, a UTF-8 byte offset, and an ordered selected range. `set_cursor(offset, extend_selection)` rejects offsets inside a grapheme. `insert` replaces the selection; `set_text` replaces the entire document. Both are atomic on byte-limit errors. `update(terminal::Event)` handles printable keys, paste, grapheme Left/Right/Delete/Backspace, line/document Home/End, Up/Down, Ctrl+Left/Right word movement, Ctrl+A selection, Ctrl+Z/Ctrl+Shift+Z/Ctrl+Y history. Its boolean indicates whether it handled the event.

`render(buffer, area, EditorOptions)` follows the cursor with horizontal/vertical scrolling. Options select multiline rendering, a masked password view, focus, and base/selection/cursor styles. Masking emits one bullet per grapheme. Single-line rendering replaces line breaks visually; applications should intercept Enter for submission and choose their own paste policy. `undo`, `redo`, and `clear_history` manage at most 64 snapshots within a 4 MiB history budget. Password masking is a display feature: text and undo history are ordinary garbage-collected strings and are not securely erased. Call `clear_history` when retaining password undo records is undesirable.

## Bounds and validation

Coordinates and each dimension are at most 4,096; buffers have at most 1,048,576 cells. A cell accepts at most 4,096 input bytes, total cell text is capped at 4 MiB, and a drawing text input at 1 MiB. Renderer output is capped at 16 MiB per frame. The memory backend retains at most 1,024 writes and 16 MiB total. Layouts allow 4,096 constraints; tables allow 256 columns and 65,536 rows; lists and trees allow 65,536 items/nodes. Trees reject depth jumps, duplicate IDs, depth above 256 and more than 8 MiB total ID/label text. Charts allow 256 series, 65,536 points and 1,048,576 segment raster steps. Editor capacity is configurable up to 1 MiB.

The implementation uses the Unicode library's default width policy (narrow ambiguous characters, wide emoji). Terminal/font-specific alternative width policies, bidirectional shaping, persistent editor storage, soft-wrapped editing, system clipboard, widget click routing and GPU/image protocols are outside the current API.

Validation includes library snapshots and edge cases, independent consumer tests, 392 exhaustive wide-cell overwrite/fill cases, 1,650 small-layout cases, 2,500 seeded model mutations replayed across 2,505 emitted ANSI frames, and a real PTY test for input, paste, resize and terminal restoration:

```sh
just ecosystem-test tui
```
