# GoML formatting

`gomlfmt` is the canonical formatter for GoML source code. It parses source into the lossless concrete syntax tree, converts that tree into a document describing possible line breaks, and renders the document at a fixed width.

Formatting never uses the lowered AST. Comments, identifiers, numeric spelling, character and byte escapes, strings, and interpolated-string text are emitted from their original lexer tokens. Multiline strings retain their marker and content text while the layout indentation before each marker is normalized.

## Fixed settings

- Maximum line width: 100 columns
- Indentation: four spaces
- Line ending: LF
- Non-empty file ending: exactly one LF

The formatter has no configuration file or command-line style settings. Tests may use narrower widths to exercise line-breaking behavior.

Columns are counted by Unicode code point. East Asian display width and grapheme clusters are not measured specially. A trailing `//` comment may exceed the maximum line width.

## Files and items

File header comments remain first. A `package` declaration occupies one line and is followed by a blank line. Consecutive `use` declarations occupy consecutive lines and the use block is followed by a blank line. Imports are not sorted, merged, or reordered.

Top-level items are separated by one blank line. Attributes remain attached to the item that follows them, with one attribute per line.

Struct, enum, trait, impl, function, match, and other non-empty bodies use multiple lines:

```goml
fn choose(value: bool) -> string {
    if value {
        "yes"
    } else {
        "no"
    }
}
```

Empty blocks are written as `{}`. `else` remains on the line of the preceding closing brace. An empty `else {}` branch without comments is omitted because an `if` without `else` has the same implicit `()` fallback.

An exact `()` at the end of a block and a standalone `();` are omitted when they contain no comments. Empty statements written as extra semicolons are also removed. These normalizations preserve the block's `()` result and do not remove comments.

Semicolons after unlabeled `while` and `for` expressions are removed. Semicolons after `if`, `match`, `select`, and unlabeled `loop` are removed when another statement follows. A final `if` with an implicit `()` fallback also omits its semicolon. Labeled loop statements retain their semicolons because the parser requires them before a following statement. A final control-flow semicolon is retained when it may discard a value, and any control-flow semicolon is retained when the next token could continue the preceding expression, such as `(`, `[`, `.`, `?`, `as`, a range operator, or a binary operator.

Struct-like enum variants with one field stay on one line when every such variant fits. If any struct-like variant has multiple fields, contains a comment, or exceeds the line width, all non-empty struct-like variants in that enum use the same expanded layout:

```goml
enum Message[T] {
    Empty,
    Value { value: T },
}

enum Pair[T] {
    Left {
        value: T,
    },
    Both {
        left: T,
        right: T,
    },
}
```

## Spaces

- A comma or colon is followed by one space when content follows on the same line.
- Binary operators, `=`, `=>`, `->`, and `as` have one space on both sides.
- `::`, `.`, postfix `?`, calls, and indexing have no surrounding padding.
- Prefix operators have no space before their operand.
- Parentheses and brackets have no inner edge padding.

The formatter preserves the CST's parentheses and associativity. It does not insert or remove grouping parentheses.

## Lists and expressions

A group remains flat when it fits:

```goml
call(first, second, third)
```

When it does not fit, each list element is placed on its own indented line:

```goml
call(
    first,
    second,
    third,
)
```

Broken lists receive a trailing comma where the grammar permits it. In addition to the block normalizations described above, conditional trailing commas are the only non-trivia tokens that formatting may add or remove. Nested groups make their own width decisions.

Long binary chains break before operators:

```goml
first
    + second
    + third
```

## Comments and literals

`//` comment text and order are preserved. A same-line comment stays attached to the preceding syntax and is emitted at the end of that line. A standalone comment remains on its own line. At most one blank line is retained between standalone comments.

Comments inside a list force a line break. A trailing comment is not wrapped or rewritten.

Literal token text is preserved byte for byte, including radix prefixes, numeric separators, escapes, and interpolated-string text. A multiline string starts on a new line, all of its `\\` markers are aligned, and the terminating syntax returns to the enclosing indentation. Only spaces and tabs before a `\\` marker are rewritten; content after the marker is preserved byte for byte.

```goml
let poem =
    \\roses are red
    \\violets are blue
;
```

## Errors

Invalid source is not partially formatted. The library returns the parser diagnostics, stdout mode emits no formatted source, and `-w` validates every input before writing any file.

## Command line

```text
gomlfmt [OPTIONS] [FILE...]

-w          write formatted source back to files
-l          list files whose contents would change
--check     exit 1 when any file needs formatting
-h, --help
--version
```

With no file, or with a single `-`, the formatter reads stdin and writes stdout. `-w` cannot be used with stdin. A single file without a mode is formatted to stdout. Multiple files require `-w`, `-l`, or `--check`.

`--check` exits 0 when every file is formatted, 1 when a change is needed, and 2 for argument, I/O, or parse errors.

From anywhere inside a module, use the project driver to format every `.gom` file:

```text
goml fmt
goml fmt --check
```

`goml fmt` searches upward for `goml.toml` and uses the same package graphs as module builds and tests. It formats the current module's production files, internal `*_test.gom` files, and black-box `tests` packages, while excluding external dependencies. Package discovery excludes `testdata`, the configured build target directory, hidden directories, and nested modules. `--check` does not write files.

The formatter executable is resolved from `--formatter`, `GOMLFMT`, the directory containing `goml`, `GOML_HOME/bin`, then `PATH`.

## Editor formatting

`gomllsp` implements `textDocument/formatting`. It formats the latest in-memory document and returns one full-document edit when the source changes. Already formatted or syntactically invalid documents return no edits.

Editor-provided indentation options are ignored because GoML formatting uses the fixed settings defined above. The full-document edit range uses UTF-16 positions as required by the language server protocol.
