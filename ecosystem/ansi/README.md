# ecosystem::ansi

Structured terminal styles, bounded streaming escape parsing and Unicode-aware
styled text, implemented in GoML. This package depends on `ecosystem::color` and
`ecosystem::unicode_text`; it does not read environment variables or terminal
input. The terminal session chooses the output profile and hyperlink policy.

## Styles and color

Declare `"ecosystem::ansi" = "0.1.0"` in the module-root dependencies.

```goml
use ecosystem::ansi;

fn greeting() -> string {
    let style = ansi::Style::new()
        .foreground(ansi::Color::Rgb(120, 180, 255))
        .with_attribute(ansi::BOLD, true);
    style.paint("Hello 世界", ansi::Profile::TrueColor)
}
```

`Color` is `Default`, `Indexed(u8)` or `Rgb(u8,u8,u8)`. `Color::from_srgb`
uses the color library's rounded RGB channels; terminal colors have no alpha.
Profiles are `Plain`, `Ansi16`, `Ansi256` and `TrueColor`. Quantization minimizes
squared distance in encoded RGB against the conventional 16-color palette and
xterm 6×6×6 cube/grayscale palette, with the lowest index winning ties. A user's
custom terminal palette can differ. `palette_rgb` exposes the assumed table.

`Style` exposes `fg`, `bg`, `attributes`, builders and `patch(StylePatch)`.
Patches distinguish inherited colors from an explicit terminal default and
provide separate attribute additions/removals; removal wins. Attributes include
bold, dim, italic, underline, blink, reverse, hidden, strike and overline.
Unknown bits are masked during rendering. `transition(previous, profile)`
emits state changes, including the shared bold/dim reset. `paint` and styled
text rendering return foreground/background/attributes to terminal defaults.
They assume the previous state starts at defaults; callers embedding output in
another renderer should use explicit transitions.

## Parsing and text

`Parser::new()` accepts UTF-8 byte fragments through `feed(Slice[byte])` and
returns `Token` values: text, C0/DEL controls, ESC sequences, CSI parameter and
intermediate fields, OSC bodies and opaque DCS/SOS/PM/APC strings. OSC accepts
BEL or ST; other strings require ST. UTF-8 and escape sequences may span feeds.
`finish()` rejects incomplete input. Failure poisons the parser until `reset()`;
copies share state and require one consumer. Tokens from a failed feed are not
published. `pending_bytes` exposes retained incomplete input.

`Limits` defaults to an 8 KiB escape sequence, 1 MiB input chunk and 128 CSI
parameter components. Its checked constructor bounds all limits. Plain text is
emitted per feed; callers comparing different chunkings should coalesce adjacent
text tokens. Parsing is strict about malformed sequences and invalid UTF-8;
single-byte C1 mode is unsupported. Opaque string payloads must also be UTF-8.
This tokenizer does not emulate terminal screen operations.

`parse_tokens` parses a complete string. `strip` removes escape sequences and
controls while retaining text, LF and tab. `parse` produces `StyledText`, applying
SGR and OSC 8 hyperlinks and discarding other commands. SGR supports semicolon
and colon extended colors; unsupported numeric attributes are ignored, while
malformed supported color forms return an error. Underline variants normalize
to one underline attribute; underline color, font selection and palette mutation
are outside the style model.

`Span::new(text, style)` and `StyledText::push` sanitize terminal controls in
ordinary text to U+FFFD, preserving LF/tab. Hyperlinks have checked URL/id fields,
reject controls and bound URL/id lengths; the library does not open their URLs.
`render_with_links` can disable OSC 8 independently of color. `Plain` always
disables both styles and links. `write`/`write_line` use generic `std::io::Write`
and preserve partial-write errors without implicit flushes.

`StyledText` provides `spans`, `snapshot`, `append`, `plain_text`, `render`,
`graphemes`, `width`, `clip`, `truncate` and `wrap`. Builder copies share their
span vector; `snapshot` and returned span vectors are independent. Graphemes are
computed across span boundaries; the style/link at the first scalar governs a
whole grapheme during layout. Default width uses Unicode 16 data, ambiguous
width one and emoji width two. Width is the maximum line width; tabs advance to
eight-column stops.

Clipping selects a column range on the first line. A partial wide-grapheme
intersection becomes spaces of the intersecting width. Wrapping is a hard
grapheme wrap preserving LF, styles and links; it expands tabs and replaces a
grapheme wider than the complete target line with U+FFFD. It always emits at least
one line and preserves a trailing empty line. Word/line-break-aware plain text
wrapping is available from `unicode_text`. Column arguments are checked and
limited to 16,777,216; storage for complete strings/span builders scales with
caller-provided input. There is no background state or implicit I/O.

Protocol semantics follow [XTerm control sequences](https://invisible-island.net/xterm/ctlseqs/ctlseqs.html).

## Verification

```sh
python3 ecosystem/verify.py ansi
```

Tests cover every attribute combination through all four profiles, fragmented
UTF-8/CSI/OSC at every byte boundary, malformed/over-limit input, wide-cell clips,
cross-span combining/ZWJ sequences, hyperlinks and partial consumer writes.
`interop.py` independently checks 2,800 generated SGR, command stripping, palette
quantization and rendered terminal-state cases with a deterministic Python model.
