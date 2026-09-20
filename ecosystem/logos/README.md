# logos

A typed lexer toolkit written in GoML, inspired by [Rust Logos](https://logos.maciej.codes/). Rules compile to one Thompson NFA. Matching is anchored at the current cursor, selects the longest accepted prefix, then uses rule priority. The regex engine has no recursive backtracking. There is no native adapter or dependency on a host regex engine.

This package uses a runtime `Grammar::compile` API. It does not implement Rust's derive macro or generate a specialized DFA. Compiled grammars can be reused across input strings and shared by independent lexers.

```goml
use ecosystem::logos;

enum Token {
    Let,
    Name(string),
    Number(string),
    Equals,
}

fn tokenize(source: string) -> Result[Vec[Token], string] {
    let rules: Vec[logos::Rule[Token, (), string]] = Vec::from_array([
        logos::Rule::token("let", Token::Let),
        logos::Rule::regex("[a-zA-Z_][a-zA-Z0-9_]*", |ctx| {
            logos::Action::Emit(Token::Name(ctx.slice()))
        }),
        logos::Rule::regex("[0-9]+", |ctx| {
            logos::Action::Emit(Token::Number(ctx.slice()))
        }),
        logos::Rule::token("=", Token::Equals),
        logos::Rule::skip("\\s+"),
    ]);
    let grammar = logos::Grammar::compile(rules).map_err(|error| error.to_string())?;
    let lexer = grammar.lexer(source, ());
    let tokens: Vec[Token] = Vec::new();
    for result in lexer {
        tokens.push(result.map_err(|_| "invalid token")?);
    }
    Result::Ok(tokens)
}
```

## Rules and disambiguation

`Rule[T, X, E]` carries a token type `T`, extras type `X`, and callback error type `E`.

| API | Behavior |
| --- | --- |
| `Rule::token(text, value)` | Literal text returning a fixed token |
| `Rule::literal(text, callback)` | Literal text with a typed callback |
| `Rule::regex(pattern, callback)` | Regex with a typed callback |
| `Rule::skip(pattern)` | Regex discarding its matches |
| `with_priority(n)` | Override automatic priority with a nonnegative integer |
| `ignore_ascii_case()` | ASCII case folding for literals and ranges |
| `with_dot_all()` | Make `.` also match LF |
| `Grammar::compile(rules)` | Compile with default budgets |
| `Grammar::compile_with(rules, options)` | Compile with explicit budgets |
| `state_count()` / `rule_count()` | Inspect compiled size |

The longest match wins even when a shorter rule has a higher priority. Default priorities count two per literal Unicode scalar and one per character class or dot. Concatenation adds, alternation takes the minimum, and repetition multiplies by its minimum count. Thus a keyword beats an identifier at the same length; `letter` still beats the keyword `let`. Repeating or optional suffixes contribute only their required portion.

Equal length and priority across different rules produce `LexError::Ambiguous(rule_indices)`, with sorted zero-based rule indices. No callback runs for an ambiguous match. This is a recoverable runtime error, unlike Rust Logos' compile-time overlap diagnosis. Rules are never silently ordered by registration. Priority values are deliberately documented here rather than promised to be numerically identical to every Rust Logos version.

## Regex support

- UTF-8 literal characters; concatenation; `a|b`; capturing-style `(a)` and non-capturing `(?:a)` groups, both without captured submatches.
- `.`, positive/negative classes, ranges, and escaped literals: `[a-z_]`, `[^"\\]`, `[\dA-F]`, `[-a-z]`.
- `?`, `*`, `+`, `{n}`, `{n,}`, `{n,m}`. Nullable internal expressions such as `(a?)*b` are supported. An entire rule that accepts the empty string is rejected.
- ASCII `\d`, `\D`, `\w`, `\W`, `\s`, `\S`. Here word is `[A-Za-z0-9_]`; whitespace is space, HT, LF, CR, VT, FF.
- `\n`, `\r`, `\t`, `\f`, `\v`, `\a`, `\0`, `\xHH`, `\uHHHH`, `\u{H...}` with checked Unicode scalar values.
- `\p{L}` / `\p{Letter}`, `\p{N}` / `\p{Number}`, `\p{White_Space}`, `\p{Lowercase}`, `\p{Uppercase}`, `\p{ASCII}`, and their `\P` complements, including inside classes. Unicode classification delegates to `std::unicode`.

Unsupported constructs produce `CompileError { rule, offset, message }`: lookarounds, anchors, word boundaries, backreferences, lazy/stacked quantifiers, inline flags, nested classes, class set operations, and other Unicode properties. Use rule methods for ASCII case folding/dot-all and callbacks for context-dependent text. Unicode case folding, named subpatterns, byte-string input, streaming/partial input, and capture groups are not provided. Pattern offsets and input spans are UTF-8 byte offsets.

## Lexer, callbacks, extras, and modes

`grammar.lexer(source, extras)` creates `Lexer[T, X, E]`. `lexer_at(source, offset, extras)` validates a starting UTF-8 boundary. `next()` returns `Option[Result[T, LexError[E]]]`; the lexer also implements `Iterator`. `spanned()` returns an iterator of `(result, Span)` pairs. End of input is fused and sets the span to the empty range at EOF.

Both the lexer and callback `Context[X]` expose:

| API | Meaning |
| --- | --- |
| `source()` | Complete input string |
| `span()` / `slice()` | Current half-open byte range / matching text |
| `remainder()` | Input after the match |
| `extras()` / `set_extras(value)` | Application state shared with callbacks |
| `bump(bytes)` | Extend the match by nonnegative bytes; reject overflow, out-of-range offsets and non-UTF-8 boundaries |
| `position()` | One-based line and Unicode scalar column at the match start; LF starts a new line |

Callbacks return `Action::Emit(token)`, `Action::Skip`, or `Action::Error(error)`. They may extend a token with `bump`, maintain indentation or nesting in extras, parse token payloads, or reject a match with an application error. Invalid `bump` leaves the cursor unchanged. Returning an error consumes the matched and bumped text, so a caller can continue lexing.

`lexer.fork()` creates an independent cursor and extras cell at the same position. The extras **value** is copied with ordinary GoML semantics: nested `Ref`, `Vec`, and other mutable values remain shared. Assignment of a lexer itself shares its cursor/extras. `lexer.morph(other_grammar)` changes token and error types while sharing the same cursor/extras; this supports lexer modes and nested language parsers. Old handles observe the new cursor position. One lexer and its aliases are not safe for concurrent mutation. Separate lexers can share a compiled grammar; callbacks must synchronize any additional captured mutable state themselves.

When no rule accepts, an error consumes the viable prefix up to the first impossible character, or one entire Unicode scalar when the prefix is empty. The character that terminates a nonempty failed prefix is left for the next token. The caller observes `LexError::Unexpected` and can resume. Matching semantics follow the general recovery approach documented in the [Logos regex guide](https://logos.maciej.codes/common-regex.html).

## Budgets and cost

`Options::new()` limits each pattern to 65,536 bytes, nesting to 64 groups, explicit repeat counts to 1,024, the combined automaton to 16,384 states, inspected token text to 1,048,576 bytes, and matching work per `next()` to 10,000,000 units. Work includes skipped rules within that same call. NFA expansion also has a cap derived from `max_states`, so repetitions of empty internal groups cannot cause unbounded compile work.

Budget exhaustion produces `LexError::BudgetExceeded` and consumes the inspected prefix, or one complete scalar if needed to guarantee progress. A single multibyte scalar can therefore exceed a tiny byte limit while recovering. `max_token_bytes` bounds lookahead as well as the accepted match. Compile options reject invalid values; nesting has an absolute ceiling of 256 and repeat/state counts a ceiling of 1,000,000.

Simulation costs approximately the inspected characters times the active NFA states/class terms, plus per-token state bookkeeping. Token fallback can revisit input; there is no whole-input linear-time guarantee. `position()` scans from the start of the source and should be replaced by extras-based line tracking in hot paths. Callback work, caller-owned source strings, payload allocation and text explicitly consumed by `bump` are outside the engine budgets.

## Validation

From the repository root:

```sh
just ecosystem-test logos
```

The external library tests cover maximal munch, automatic/explicit priorities, ambiguity, nullable loops, bounded repeats, Unicode classes/spans, error recovery, callbacks/bump, extras, modes, snapshots, limits, adversarial alternation, and independent concurrent lexers. A registry consumer verifies the public generic API. Native consumer tests compare all 3,155 retained reference cases from independent `re` full-match evaluation of every candidate prefix; they include exhaustive short binary inputs and generated multi-rule, Unicode, skip, flag and priority cases. [Fixture provenance](../consumers/logos/tests/data/README.md) records the model and seed. Running the tests requires only GoML.

Reference design: [Logos token rules](https://logos.maciej.codes/attributes/token_and_regex.html), [disambiguation](https://logos.maciej.codes/token-disambiguation.html), [callbacks](https://logos.maciej.codes/callbacks.html), and [extras](https://logos.maciej.codes/extras.html).
