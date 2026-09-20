# unicode_text

`ecosystem::unicode_text` implements Unicode 16.0.0 text segmentation and terminal layout in GoML. It contains no runtime FFI, registry dependencies, host Unicode calls, or network access. Its tables are independent of the Unicode version used by `std::unicode`.

```toml
[dependencies]
"ecosystem::unicode_text" = "0.1.0"
```

```gom
use ecosystem::unicode_text;

fn main() -> () {
    let source = "GoML 👩🏽‍💻 中文";
    println(unicode_text::width(source).to_string());
    for cluster in unicode_text::graphemes(source) {
        println(cluster.text);
    }
    for line in unicode_text::wrap(source, 8) {
        println(line);
    }
}
```

## Segmentation

The implementation follows the default extended grapheme and word rules in [UAX #29 revision 45](https://www.unicode.org/reports/tr29/tr29-45.html), and default line breaking in [UAX #14 revision 53](https://www.unicode.org/reports/tr14/tr14-53.html). All offsets are UTF-8 **byte offsets**, never scalar indexes or terminal columns. Inputs use GoML's valid Unicode `string` contract; arbitrary byte decoding belongs at the input boundary.

| API | Result |
| --- | --- |
| `graphemes(text)` | `Vec[Grapheme]`, each with public `text`, `start`, exclusive `end` |
| `grapheme_boundaries(text)` | All byte boundaries, including start and end; empty input returns `[0]` |
| `is_grapheme_boundary(text, offset)` | False for invalid offsets and positions inside a cluster |
| `word_boundaries(text)` | Default word boundaries, including punctuation and whitespace segments |
| `word_segments(text)` | `Vec[Word]` with `text`, `start`, `end`, `is_word` |
| `words(text)` | Word segments containing a Unicode letter or number; punctuation-only and emoji-only segments are excluded |
| `line_breaks(text)` | `Vec[LineBreak]` with `offset` and `mandatory`; includes end-of-input, excludes start except for empty input |

Grapheme segmentation includes Indic conjuncts, Hangul, prepend characters, spacing marks, regional-indicator pairing, and extended pictographic ZWJ sequences. Word segmentation implements ignored marks, Hebrew quotes, numeric punctuation, Katakana, and regional indicators. Line breaking includes combining-mark inheritance, nonbreaking spaces and joiners, East Asian punctuation, contextual quotes, numeric expressions, Brahmic syllables, and explicit Unicode hard separators.

Default word boundaries are not dictionary word segmentation for Chinese, Japanese, Thai, or other languages requiring linguistic context. Line class `SA` receives UAX #14's default category-based fallback. There is no locale tailoring, dictionary hyphenation, normalization, bidirectional reordering, or sentence segmentation.

## Terminal measurement

`WidthOptions::new()` has public `ambiguous_wide: false` and `emoji_wide: true` fields. `width(text)` uses those defaults; `width_with(text, options)` accepts explicit policy. `max_line_width(text, options)` measures each hard-separated line and returns its maximum.

A cluster occupies the maximum width of its visible scalars. East Asian Wide/Fullwidth characters use two columns; other visible scalars use one, with an optional two-column policy for East Asian Ambiguous characters. Combining marks, controls, formatting/default-ignorable characters, line separators, and medial/final Hangul jamo have zero scalar width. Standalone combining marks therefore have width zero.

Emoji-presentation clusters, VS16 emoji, keycaps, flags, and pictographic ZWJ sequences occupy two columns by default. `emoji_wide: false` assigns such clusters one column, even when their constituent scalars have East Asian Wide width. VS15 suppresses the emoji-presentation override; East Asian width still applies. These policies provide deterministic layout; actual glyph width depends on terminal and font behavior. No terminal probing or font shaping is performed.

`width` sums cluster widths across all input. Newlines and tabs contribute zero; they do not move a simulated cursor. ANSI escapes are ordinary input characters: an ESC control itself has zero width, but printable sequence payload has width. Strip or parse ANSI using `ecosystem::ansi` first. `expand_tabs` provides explicit tab-stop expansion when required.

## Truncation, padding and wrapping

| API | Behavior |
| --- | --- |
| `truncate(text, columns)` / `truncate_with(text, columns, width)` | Longest whole-grapheme prefix fitting the column budget; a negative budget returns empty |
| `truncate_checked(text, columns, width, limits)` | Checked input/output limits and nonnegative budget |
| `ellipsize(text, columns, marker, width, limits)` | Whole-grapheme truncation with marker included in the budget |
| `pad(text, columns, Alignment::{Left,Center,Right}, width, limits)` | Pads with ASCII spaces; preserves text already wider than the target |
| `expand_tabs(text, tab_stop, width, limits)` | Expands tabs at measured columns and resets after hard separators |
| `wrap(text, columns)` / `wrap_with(text, columns, width)` | Convenience wrapping with default trim and emergency-break behavior |
| `wrap_configured(text, WrapOptions)` | Checked wrapping returning `Vec[WrappedLine]` with text, source byte range, measured width, and `hard_break` |

Zero-width prefixes fit a zero-column truncation budget. Ellipsis markers are themselves truncated at grapheme boundaries if necessary; if no marker grapheme fits, the result uses the available budget for the original prefix. Padding never truncates; centered padding puts an odd extra space on the right.

`WrapOptions::new(columns)` exposes `columns`, `width`, `trim_spaces: true`, `break_long_words: true`, and `limits`. Wrapping prefers Unicode line-break opportunities, intersects them with grapheme boundaries, removes hard separators, and preserves empty hard-separated lines, including a final empty line. `hard_break` distinguishes a physical separator from a soft wrap or end-of-input.

Trimming removes ASCII spaces at line ends and at the start of a continuation. Set `trim_spaces: false` to preserve whitespace across soft wraps. Nonbreaking and other Unicode spaces are preserved. Long words receive emergency breaks at grapheme boundaries by default; set `break_long_words: false` to permit oversized words. A single cluster wider than the target is emitted intact on an oversized line. Consequently, wrapping can return a line wider than `columns`; consumers requiring strict clipping should call `truncate`. Tabs are not expanded automatically. The convenience wrappers clamp nonpositive columns to one; the checked API rejects them.

`LayoutLimits::new()` allows 1 MiB input, 16 MiB output, and 100,000 lines. Limits are public and checked before output growth; negative byte limits and nonpositive line limits are rejected. Tab stops must be within 1–1,048,576. `Error` distinguishes invalid options, input, output, and line limits, and implements `ToString`. Segmentation and measurement convenience APIs have input-proportional allocations without an arbitrary input cap; callers handling untrusted streams should bound decoded input before invoking them.

## Data and reproducibility

[Unicode source manifest](data/manifest.json) pins the source URLs and SHA-256 hashes for Unicode 16.0.0 properties and all three official segmentation test suites. Only this manifest, the [Unicode license](data/LICENSE.txt), and generated `properties.gom` are checked in. Raw property files, test corpora and compressed archives are not versioned. Every generator or conformance-checker invocation downloads fresh source bytes from the pinned URLs and checks their SHA-256 hashes; downloads are held in memory for that invocation, with no persistent cache or bundled-data fallback. Ordinary GoML builds and module tests use the generated tables directly and do not need these downloads.

```sh
just ecosystem-test unicode_text
```

The separate native GoML module in `tools/` downloads the sources and runs both exact table verification and official conformance as an ordinary `#[test]`. Its `unicode_data` executable also accepts `check` or `generate` followed by the Unicode module directory. With the local registry configured, build it from `ecosystem/unicode_text/tools` using `../../../stage2/bin/goml build`, then run `_artifact/bin/unicode_data generate ..` to rewrite the table.

Every `check`, `generate`, or native conformance-test invocation downloads fresh data again. Download failures, a response above 16 MiB, or checksum mismatches stop verification; they never fall back to old files. Downloads are held in memory and never written to the repository. Injected-fetch native tests also check fresh requests on every invocation, both no-cache headers, checksum and transport failures without reuse, oversized bodies and invalid UTF-8. The conformance step therefore requires network access to unicode.org. The generated table consists of sorted, nonoverlapping intervals encoded in one immutable string and searched with binary search. Segmentation tracks preceding/following significant classes and regional-indicator counts rather than repeatedly scanning long runs. Runtime work is linear in input length with logarithmic property lookup and input-proportional temporary storage.

## Validation

From the repository root:

```sh
GOML_BUILD_JOBS=2 just ecosystem-test unicode_text
```

The independent consumer exercises public imports, byte offsets, widths, and checked layouts. The separate native tooling module imports the versioned library and checks all 1,093 official `GraphemeBreakTest`, 1,826 `WordBreakTest`, and 16,672 `LineBreakTest` cases directly. Eleven terminal-width reference samples also run in the consumer’s ordinary GoML tests. Unit tests also cover emoji/variation policies, CJK and ambiguous widths, long combining and flag runs, source ranges, mandatory separators, tabs, overflowing clusters, trim policies, and explicit allocation limits.
