# ecosystem::parser

Composable text and binary parsers implemented in GoML. Parsers retain concrete
function values, allowing generic transformations and recursive grammars across
package and module boundaries without a runtime type-erasure convention.

```goml
use ecosystem::parser;

fn numbers(input: string) -> Result[Vec[isize], parser::Error] {
    parser::token(parser::integer())
        .separated(parser::token(parser::literal(",")), 0, false)
        .between(parser::literal("["), parser::literal("]"))
        .parse(input)
}
```

## Text API

`Parser[T]::new` accepts `(string, isize) -> Result[(T, isize), Error]`.
`parse` requires complete consumption; `parse_at` accepts a byte offset and
returns the next offset; `parse_prefix` returns the remaining text. The parser
checks offsets, forward progress and UTF-8 boundaries before exposing a result.

### Resource budgets

Every text and binary entry point starts with `ParseLimits::standard()`: 256
nested parser calls and 1,000,000 work units. `parse_with_limits(input, limits)`
overrides these limits; depths must be 1 through 4096 and work nonnegative.
One `ParseContext` is shared by all built-in combinators, alternatives, repetitions,
dependent parsers and `lazy` calls in that parse. Backtracking never replenishes
work. Each invocation costs one unit; byte scans, literal comparisons and error
expectation concatenation also consume work. Binary `take` returns a slice in
constant time and charges its invocation rather than every returned byte.

Depth/work exhaustion returns a committed error at the stopping byte offset.
It remains sticky through `attempt`, `label`, `optional`, negative lookahead and
user callbacks that swallow the immediate error. Left-recursive grammars
therefore fail recoverably at the depth limit instead of overflowing the stack;
they still need rewriting to parse successfully. The iterative `choice` helper
does not add one call-stack frame per alternative.

For custom combinators, use `Parser::with_context` with a
`(ParseContext, input, offset)` callback and call children with
`child.parse_in(context, input, offset)`. `ParseContext::new`, `charge`,
`work_used`, `peak_depth` and `limit_error` support explicit work accounting and
diagnostics. Context copies share their work budget and failure state; use a
fresh context for an independent parse. `nested` is the invocation-accounting
hook used by parser implementations, and callers ordinarily let `parse_in` use it.

The original two-argument `Parser::new` constructor remains available for leaf
callbacks. Work inside arbitrary callbacks cannot be preempted or measured;
calling a top-level `parse`/`parse_at` from one starts an independent budget.
Custom nested parsers must use the contextual constructor to participate in the
outer limit. These budgets bound library work, not wall time or exact heap size.

| Category | API |
| --- | --- |
| Transformation | `map`, `try_map`, `and_then`, `verify` |
| Sequencing | `then`, `left`, `right`, `between` |
| Alternatives | `or`, `choice`, `optional`, `pure`, `fail`, `lazy` |
| Repetition | `many`, `many1`, `repeat(minimum, maximum)`, `separated(separator, minimum, trailing)` |
| Observation | `peek`, `not`, `position`, `end`, `spanned`, `recognize` |
| Errors | `cut`, `attempt`, `label`, `context`, `Error::render` |
| Operators | `chain_left`, `chain_right` |
| Text primitives | `literal`, `take_while`, `satisfy`, `character`, `any_char`, `one_of`, `none_of` |
| Lexical helpers | `whitespace`, `token`, `digits`, `integer`, `decimal`, `identifier`, `quoted`, `line_ending`, `rest_of_line` |

`or` backtracks unless the failed branch is committed with `cut`. It reports the
furthest error and combines expectations on a tie. `attempt` clears commitment.
`between` commits after its opening delimiter. Repetitions reject successful
zero-width iterations instead of looping forever. Lists explicitly choose
whether a trailing separator is accepted. `peek` retains input; `not` preserves
committed errors. Error offsets and spans use bytes; rendered columns count
Unicode scalar values, not terminal display cells or UTF-16 units.

`identifier` accepts Unicode letters and underscore initially, then Unicode
letters/numbers and underscore. `integer` accepts an optional sign and decimal
digits with machine-range checks. `decimal` additionally accepts a fractional
part and exponent; a decimal point or exponent marker requires following digits.
`quoted` supports the selected quote and `\\`, `\n`, `\r`, `\t` escapes.
`take_while` uses byte predicates and rejects an ending inside a UTF-8 scalar.

`arithmetic::expression()` demonstrates recursive parentheses, multiplication,
addition and subtraction with precedence. `lazy` supports productive recursive
grammars; left recursion must be eliminated. Text parsers consume an available
string, while incremental binary parsing uses the API below.

## Binary API

Import `ecosystem::parser::binary`. Its separate `Parser[T]` takes a
`Slice[byte]` and supports `parse`, `parse_at`, `map`, `try_map`, `then`, `left`,
`right`, `and_then`, bounded `repeat`, `or`, `cut`, `attempt`, `optional`, `peek`,
`not`, `verify` and `between`. `pure`, `lazy` and `end` support recursive layouts.
The contextual constructor and resource-limit APIs match the text parser.

`take`, `literal`, `byte_value`, `uint16/32/64`, `int16/32/64`, `float32/64` and
`length_prefixed` cover common frame layouts. Numeric parsers take
`std::bytes::endian::Endian`. Slices retain the input backing storage, and literal
patterns are snapshotted at construction. Do not concurrently mutate input views.

`Error::Incomplete { offset, needed }` distinguishes a valid truncated prefix
from `Error::Invalid`. Append more input and retry `parse_at` at the frame start;
the parsers keep no hidden partial state. `length_prefixed` validates its maximum
before waiting for or slicing the payload. The caller owns buffering and I/O.
Binary alternatives backtrack only on uncommitted invalid input. `Incomplete`
propagates unchanged through alternatives, optional parsing and lookahead so a
short transport read cannot prematurely select a different branch. `cut` and
`attempt` affect invalid errors, preserving the needed-byte count on truncation.

## Validation

```sh
just ecosystem-test parser
```

The module tests cover precedence, commitment, backtracking, invalid ranges,
zero-width repetitions, Unicode spans, errors, malformed numbers and every prefix
of a binary frame. The separate consumer uses `ecosystem::proptest` to check
integer and list roundtrips through registry-resolved dependencies.
Budget tests cover left recursion, custom nested callbacks, shared sibling work,
repetition, error relabeling, swallowed errors and binary prefix alternatives.
