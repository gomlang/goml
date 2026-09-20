# decimal

An immutable, finite, base-10 decimal library implemented in GoML, using the
versioned `ecosystem::bigint = "0.1.0"` dependency for integer arithmetic. There
is no floating-point intermediate or Go decimal adapter.

`Decimal` stores a signed `BigInt` coefficient and an `isize` scale:
`value = coefficient × 10^(-scale)`. Private fields and immutable bigint
storage allow safe sharing between tasks. Scale is preserved even when it
does not affect numeric equality: `1.2300` retains coefficient `12300` and
scale `4`.

## Using the library

```toml
[dependencies]
"ecosystem::decimal" = "0.1.0"
```

```gom
use ecosystem::decimal::{Decimal, Context, Rounding};

fn invoice() -> Result[string, ecosystem::decimal::Error] {
    let price = Decimal::parse("19.995")?;
    let exact_total = price.mul(Decimal::from_u64(7))?;
    exact_total.format_scale(2, Rounding::HalfEven)
}

fn ratio() -> Result[Decimal, ecosystem::decimal::Error] {
    let context = Context::new(12, Rounding::HalfEven)?;
    context.div(Decimal::one(), Decimal::from_u64(7))
}
```

The first result is `139.96`; the second is exactly the stored decimal
`0.142857142857`.

## Operations

| API | Behavior |
| --- | --- |
| `parse`, `new`, `from_i64`, `from_u64` | Exact construction; parsing accepts optional sign, decimal point, and signed `e`/`E` exponent. |
| `coefficient`, `scale`, `precision`, `adjusted_exponent` | Representation inspection; zero has one coefficient digit. |
| `add`, `sub`, `mul` | Exact arithmetic, returning a recoverable error if the result representation exceeds the limits. Addition uses the larger input scale; multiplication adds scales. |
| `div_exact` | Reduce common factors, then factor the denominator into powers of 2 and 5. Reject nonterminating quotients. Results normalize trailing zeros. |
| `rescale`, `rescale_exact`, `rescale_status` | Round to a decimal scale, reject any loss of nonzero digits, or return explicit rounding information. |
| `normalize` | Remove coefficient trailing zeros down to the minimum supported scale; normalize all zeros to scale zero. |
| `neg`, `abs`, `signum`, `is_zero`, `is_negative` | Exact sign operations. |
| `trunc`, `floor`, `ceil` | Integral decimals with scale zero and explicit directional rounding. |
| `to_i64`, `to_u64`, `to_bigint` | Exact conversions; fractions and out-of-range results are errors. |
| `cmp`, `Eq`, `Ord`, `Hash` | Numeric semantics independent of scale, suitable for sorted collections and hash keys. |
| `same_quantum`, `same_representation` | Compare scales only, or compare coefficient and scale together. |
| `to_fixed`, `to_string`, `to_scientific`, `format_scale` | Fixed output, exact scientific output, or rounded fixed-scale output. |

`Context::new(precision, rounding)` applies a significant-digit precision;
`Context::standard()` uses 28 digits and half-even rounding. Context provides
`round`, `add`, `sub`, `mul`, `div`, and `quantize`. Addition/subtraction and
multiplication compute their full bounded intermediate before rounding, so
cancellation and carry do not cause double rounding. Division calculates the
decimal exponent and rounds one exact quotient/remainder pair.

Context `quantize(value, scale)` fixes the scale and then checks precision.
For precision 2, rounding `9.95` can produce `10`, while quantizing it to scale
1 fails because `10.0` needs three coefficient digits. Plain `rescale` has no
context precision restriction and can return `10.0`.

Seven rounding modes are available: `HalfEven`, `HalfUp`, `HalfDown`, `Up`
(away from zero), `Down` (toward zero), `Ceiling`, and `Floor`. The three half
modes differ only on an exact midpoint. Directional modes account for the
sign even when the truncated quotient is zero.

`round_status`, `add_status`, `sub_status`, `mul_status`, `div_status`, `quantize_status`,
and `rescale_status` return `Outcome { value, rounded, inexact }`.
`rounded` records discarded coefficient positions, including zeros;
`inexact` records a change in numeric value. Division also reports rounding
when precision forces an exponent above the preferred input scale
difference. These are per-operation outcomes, with no global flags or traps.
Context division can pad an exact result with zeros to the requested
precision; use `normalize()` when that quantum is unnecessary. Division of
zero returns the canonical zero.

## Serialization

Human-readable Serde formats encode an exact scientific string, retaining
both coefficient and scale. For example `1.2300` becomes `"1.2300e+0"` and
`1e2` becomes `"1e+2"`. Deserialization also accepts valid fixed decimal
strings. JSON numbers are rejected to prevent accidental floating-point
conversion elsewhere in a data pipeline.

Binary Serde formats encode a two-element tuple: signed 32-bit scale followed
by the bigint canonical signed big-endian two's-complement byte encoding.
Tuple framing, scale endianness, and byte-sequence framing follow the selected
serializer. Deserialization checks scale and coefficient limits, and bigint
rejects redundant sign bytes and empty integer encodings. Both formats
preserve representation, except that a negative zero sign is never stored.

## Resource bounds and exclusions

- Coefficients contain at most **4096 decimal digits**, and context precision
  is `1..4096`. These are decimal digit limits, not bit precision. An early
  13607-bit check avoids converting an obviously oversized bigint to text.
- Stored scale is `-4096..4096`. Scale and precision arguments are validated
  before subtraction, exponentiation, or output allocation, including machine
  minimum/maximum integer inputs.
- Parsing accepts at most 4096 mantissa digits, including leading zeros, and
  8208 input bytes. A parsed exponent magnitude above 8192 is rejected while
  scanning, before machine-integer overflow. The final scale still must be in
  range.
- Powers of ten are limited to exponent 12288; arithmetic's largest aligned
  coefficient is at most 12289 digits. Multiplication intermediates have at
  most 8192 digits. These bounds remain below bigint's standard 65536-bit
  capacity. Exact operations check their resulting representation, rather
  than automatically dropping trailing zeros to make it fit.
- Division results must also fit the stored scale. Near a scale boundary,
  context division's precision padding can exceed that limit even for a
  numerically representable exact quotient; `div_exact` is appropriate when
  a finite quotient is required. Exact division factors at most 8192 powers
  of either 2 or 5 before returning a resource error.
- Fixed formatting allocates at most 8193 ASCII bytes including the sign;
  no locale, grouping, binary-float conversion, NaN, infinity, negative zero,
  subnormal/trap model, square root, or transcendental functions are provided.
  Inputs such as `-0.00` become a positive zero with scale 2.

All invalid input, division-by-zero, inexact exact-conversion, precision,
scale, coefficient, and resource failures return `Result` errors. No global
rounding mode, mutable coefficient, or hidden thread-local state is used.

## Algorithms and validation

For `n` coefficient limbs, bigint uses linear addition, schoolbook quadratic
multiplication and long division. Decimal scale alignment multiplies by a
bounded power of ten; comparisons inspect exponents and padded decimal digit
strings without constructing a giant aligned bigint. Normalization scans
trailing digits once, then reparses the remaining coefficient. Conversion
between limbs and decimal text and integer parsing have quadratic worst-case
cost. Exact division uses a gcd plus bounded repeated factor removal.

Black-box tests cover seven rounding modes, signed midpoints, carries,
precision versus fixed-scale semantics, cancellation, recurring and finite
division, representation-preserving serialization, checked conversions,
machine-integer and resource boundaries, numeric hash identity, immutable
aliasing, and concurrent shared arithmetic. The independent consumer computes
an invoice and recurring ratio through a versioned dependency. Its native GoML tests compare 3,072 frozen independent `decimal` reference vectors across all seven rounding modes, including error-versus-success agreement for quantize and division and agreement on the `Rounded` and `Inexact` flags of every successful result. [Vector provenance](../consumers/decimal/tests/data/README.md) records the reference and seed. The native ecosystem verifier exercises concurrent arithmetic and detached coefficient bytes under Go’s race detector. Python is not required.

Run from the repository root:

```sh
just ecosystem-test decimal
```
