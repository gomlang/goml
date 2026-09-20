# syntax consumer

An independent versioned client of `ecosystem::syntax 0.1.0`. `config.gom` implements this small lossless configuration language:

```text
# preserved document comment
server {
  host = "本地😀" # preserved trailing comment
  port = 8080
  tls {
    enabled = true
  }
}
```

Keys and section names are nonnumeric words. Values are one word, a checked machine integer, or a JSON-escaped quoted string. Spaces, tabs, CRLF, comments, malformed lines and missing closing braces are retained as syntax. Parsing recovers at line/brace boundaries and reports byte ranges with domain diagnostics. Nested sections are bounded to 128 levels. This example grammar is not TOML or a general configuration standard.

`Entry` implements the external `syntax::AstNode` trait. `entries` exposes direct typed children; `section` navigates a name path; `string_value` and `integer_value` check scalar types. `rewrite` replaces the first matching direct key's scalar token, preserving every other byte and the old snapshot. Duplicate section names select the first matching section. A replacement must be exactly one valid scalar, without leading/trailing trivia. It does not create missing keys or close missing braces.

The executable runs a checked parse/read/rewrite example by default. `--json` accepts the deterministic tree and configuration requests used by the native reference tests; it returns structural/range reports and edit results. Use `just ecosystem-test syntax` at the repository root to provision its registry snapshot, format-check, build, test and run the consumer.
