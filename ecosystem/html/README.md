# HTML utilities

Pure GoML, independently versioned HTML utilities. This module owns shared
escaping mechanics, not template context analysis or an HTML sanitizer.

`escape(value, max_bytes)` escapes ampersand, angle brackets and both quotes,
using numeric quote references. `escape_with(value, options, max_bytes)` selects
apostrophe escaping and numeric versus named double-quote references explicitly.
Negative limits fail; exact output length is checked before output allocation.
`escaped_len` computes that byte length with overflow checking.

`escape_text` is the unbounded legacy-compatible text helper: double quotes become
`&quot;`, apostrophes remain unchanged. It is not suitable for single-quoted
attribute interpolation. All helpers preserve other UTF-8 bytes, including NUL,
and escape ampersands in already-escaped text again. Output is not marked safe.
Unquoted attributes, JavaScript, CSS, URL schemes and contextual template safety
need their own policy; escaping alone does not make those contexts safe.

Template and Markdown retain compatible public wrappers backed by this module.
`named_entity(name)` looks up one of the 2,125 canonical HTML5 names, without
an ampersand or trailing semicolon, and returns its one- or two-scalar replacement.
Lookup is case-sensitive and unknown names return None. Markdown's stricter
entity parser uses this same table. This lookup does not parse character references
or apply HTML's semicolon-optional or attribute-context rules by itself.

`unescape(value, max_bytes)` decodes character references in HTML text.
`unescape_with(value, DecodeContext::Attribute, max_bytes)` selects attribute-value
rules: a semicolon-free named reference followed by an ASCII letter, digit or `=`
is not consumed. Both use longest matching names, including the 106 legacy names
that permit missing semicolons. Unknown names and malformed numeric syntax remain
literal. Decoding is single-pass: `&amp;lt;` becomes `&lt;`, not `<`.

Numeric references accept decimal or hexadecimal digits with optional semicolons.
Zero, surrogates and out-of-range values become U+FFFD; C1 values use HTML's
Windows-1252 replacement mapping. Arbitrarily long digit sequences saturate without
integer wraparound. Other control/noncharacter scalars are retained. There are
intentional differences from Go 1.26's numeric-reference parser: a single decimal
digit without a semicolon is decoded both at EOF and before a non-digit, so `&#1!`
becomes U+0001 followed by `!`, and `&#0&amp;` becomes U+FFFD followed by `&`.
Go incorrectly leaves those numeric references literal because its consumed-length
check mistakes them for missing digits. Conversely, `&#x;` and `&#X;` contain no
hexadecimal digits and remain literal here; Go incorrectly converts them to U+FFFD.
Overflowing digit sequences also do not reproduce Go's 32-bit integer wraparound.

Negative limits fail. Output growth is checked before each append and errors return
no partial string. Limits bound logical output bytes, not input length or total
process memory. Decoding does not sanitize markup, check URL schemes or validate
which tokenizer state a caller is in. These are whole-string utilities, not
streaming HTML tokenizers.

Boundary expectations follow the HTML Standard's
[numeric reference states](https://html.spec.whatwg.org/multipage/parsing.html#numeric-character-reference-state)
and [recovery rules](https://html.spec.whatwg.org/multipage/parsing.html#numeric-character-reference-end-state):
missing semicolons do not prevent numeric decoding, absent hexadecimal digits
leave literal text, and control/noncharacter parse errors do not imply deleting
the scalar. The utility returns decoded text rather than tokenizer diagnostics.

The independently sourced data lives in [data/entities.json](data/entities.json),
with provenance in [data/README.md](data/README.md) and its retained
[license](LICENSE.entities.txt). The existing native generator under
`../markdown/tools` generates both this module's lookup and Markdown's compatibility
wrapper; no Python installation or runtime data file access is required by HTML.

Run `just ecosystem-test html template markdown` from the repository root.
