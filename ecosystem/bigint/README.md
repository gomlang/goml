# bigint

Pure GoML arbitrary-precision signed and unsigned integers, with no Go adapter or native dependency. The module is `ecosystem::bigint`; `ecosystem/consumers/bigint` resolves its independent `0.1.0` registry dependency and exercises exact coefficients, serialization, and map keys.

The API takes inspiration from [num-bigint](https://docs.rs/num-bigint/latest/num_bigint/struct.BigInt.html). Division explicitly distinguishes truncating quotient/remainder from Euclidean division, as in [Go math/big](https://pkg.go.dev/math/big#Int.QuoRem). These are API references; the arithmetic implementation is GoML code.

## Representation and immutability

`BigUint` stores normalized little-endian base-2^32 limbs in a private `FrozenVec[u32]`. Zero has no limbs. `BigInt` combines a magnitude with a sign; negative zero is normalized to zero. Assignment and `abs()` share immutable storage. Arithmetic returns new values. Byte and limb imports copy input data; exports return detached mutable vectors. Values can be shared across concurrent tasks and used as hash-map keys.

Addition and subtraction propagate carries and borrows. Multiplication uses schoolbook limb multiplication with exact `u64` intermediates. Division has a single-limb fast path and normalized multi-limb long division with quotient estimation, correction, and add-back. There is no conversion to strings or floating point inside arithmetic.

## Core API

```goml
use ecosystem::bigint::{BigInt, BigUint, Error};

fn calculate() -> Result[string, Error] {
    let coefficient = BigInt::parse("-123456789012345678901234567890")?;
    let scaled = coefficient.mul(BigInt::pow10(4)?);
    let (quotient, remainder) = scaled.div_rem_euclid(BigInt::from_u64(97))?;
    Result::Ok(quotient.to_string() + ":" + remainder.to_string())
}
```

Both types provide `zero`, `one`, `from_u64`, `parse`, `parse_radix`, `to_radix`, `to_string`, `is_zero`, `is_one`, `is_odd`, `bits`, `cmp`, `add`, `mul`, `div`, `rem`, `div_rem`, `pow`, `pow10`, `gcd`, `lcm`, `modpow`, `shl`, `shr`, `bit`, `bitand`, `bitor`, and `bitxor`. `PartialEq`, `Eq`, `Hash`, `Debug`, `Default`, `std::cmp::PartialOrd`, and `std::cmp::Ord` are implemented.

`BigInt` also provides `from_i64`, `from_sign_magnitude(negative, BigUint)`, `abs() -> BigUint`, `is_negative`, `signum() -> i8`, `neg`, `sub`, `bitnot`, and `div_rem_euclid`. Its `gcd` and `lcm` return nonnegative `BigUint` values. `BigUint::sub` returns `Result` and reports unsigned underflow. Its additional methods are `count_ones`, `trailing_zeros() -> Option[isize]`, and detached little-endian limb import/export.

Checked `to_i8/i16/i32/i64/isize` and `to_u8/u16/u32/u64/usize` conversions return `Option`, never truncate. Machine-width conversions use the current Linux amd64 target's 64-bit `isize` and `usize`.

### Arithmetic semantics

- `div_rem` truncates signed division toward zero; nonzero remainder has the dividend's sign. `a = q * b + r`, with `abs(r) < abs(b)`.
- `div_rem_euclid` returns `0 <= r < abs(b)`, including negative divisors. `div` and `rem` use truncating semantics.
- Zero divisors return `Error::DivisionByZero`. `gcd(0, 0)` and any `lcm` with a zero argument return zero.
- `pow` takes a `u64` exponent, with `0^0 = 1`. `modpow` takes a `BigUint` exponent and positive `BigUint` modulus; signed bases produce a nonnegative residue. Modulus one produces zero, including a zero exponent.
- Signed bitwise operations act on infinite two's-complement representations. `bitnot(x) = -x - 1`; right shift rounds negative values toward negative infinity. `bits()` measures the magnitude. Negative bit indexes return false. `trailing_zeros(0)` returns `None`.
- Shift counts are signed machine integers; negative counts return `Error::NegativeShift`. Large right shifts allocate no huge intermediate value and return zero or negative one as appropriate.

### Text and byte encodings

Text parsing accepts optional ASCII `+`, signed `-`, and ASCII digits in radix 2 through 36. Letters are case-insensitive. Leading zeros are accepted. Whitespace, prefixes such as `0x`, underscores, Unicode digits, and a bare sign are rejected. `BigUint` rejects even `-0`. Output is canonical lowercase without a prefix or leading zeros.

`BigUint::from_bytes_be/le` imports an unsigned magnitude. `BigInt::from_signed_bytes_be/le` imports two's-complement signed values. Corresponding `to_bytes_be/le` and `to_signed_bytes_be/le` methods return minimal encodings; zero is a single zero byte. Signed positive values receive a zero sign byte when required. Raw import accepts redundant leading sign/zero bytes and empty input as zero, subject to input size limits. Use `from_sign_magnitude` with unsigned byte import when the sign is stored separately.

## Resource limits

`Limits::standard()` permits 65,536 magnitude bits, 20,000 input digit bytes, and exponent 1,000,000. Parsing, left shift, powers, modular powers, limb import, and byte import have `_with_limits` variants. `parse_radix_with_limits` includes sign bytes in error offsets but excludes them from the digit count. Limits must be nonnegative. Zero-bit budgets accept zero values.

`pow_with_limits` checks the numeric exponent and every intermediate power/result; it returns an error before constructing an arbitrarily large requested exponentiation result. `modpow_with_limits` interprets `max_exponent` as the exponent's bit count (the number of binary exponentiation rounds), and bounds base/modulus magnitude bits. Multiplication intermediates before reduction or limit rejection can occupy up to twice the permitted bit length. Byte and limb import also cap the incoming representation length before copying, so a huge redundant encoding is rejected even if its mathematical value is small.

`shl_with_limits` checks the count and result size before allocating; a zero shifted by any nonnegative count stays zero. Right shift needs no growth limit. Custom limits are an explicit caller-selected budget, not a global hard maximum. Ordinary `add`, `sub`, `mul`, bitwise operations, `gcd`, `lcm`, and formatting do not impose this budget; applications that repeatedly grow values must enforce their own overall size/work policy. There is no operation cancellation or constant-time cryptographic guarantee.

## Serde

Both types implement direct `std::serde::Serialize` and `Deserialize`.

| Format kind | BigUint | BigInt |
| --- | --- | --- |
| Human readable | Canonical decimal string, including JSON quotes | Canonical signed decimal string, including JSON quotes |
| Binary | `serialize_bytes` event containing minimal unsigned big-endian bytes | `serialize_bytes` event containing minimal signed two's-complement big-endian bytes |

For example, signed `-129` is JSON `"-129"` or binary payload `ff 7f`. Binary framing/length prefixes belong to the format; Bincode legacy emits its eight-byte little-endian length followed by that payload. Zero is payload `00` for both types. These encodings are library-specific and are not claimed wire-compatible with Rust num-bigint's Serde implementation. Signed and unsigned binary payloads must be decoded using the matching type.

Deserialization uses standard limits and rejects noncanonical representations: empty binary input, redundant leading bytes, `+1`, `01`, and `-0`. Ordinary parsing/import is more permissive as described above. Values built using larger custom limits or unchecked growth may need application-specific serialization to exceed the standard decode budget. Configure the outer format's own input/allocation limit: its string/byte payload can be allocated before this library receives it.

## Complexity and limits

With `n` and `m` limbs, addition/subtraction and bitwise operations are linear, multiplication is O(nm), and normalized division is O((n-m+1)m) for `n >= m`. Powers use repeated squaring. GCD uses Euclidean division. Text parsing uses multiply-add per digit; formatting divides by the largest radix power that fits a limb and emits grouped digits. General-radix conversion is quadratic in digit length. This implementation does not include Karatsuba/FFT multiplication, Montgomery reduction, modular inverse, roots, random-prime generation, rational arithmetic, floating-point conversions, or operator overloading.

## Verification

From the repository root:

```sh
just ecosystem-test bigint
```

Black-box tests cover signed division identities, quotient-correction boundaries, carry chains, fixed-width conversion extremes, all radices, canonical Serde, resource limits, detached storage, and shared concurrent arithmetic. The versioned consumer has its own tests. The consumer’s native GoML tests check 1,938 frozen independent reference vectors covering multi-limb arithmetic, signed bitwise operations and shifts, both division conventions, radices, signed bytes, powers, GCD/LCM, and modular powers. [Vector provenance](../consumers/bigint/tests/data/README.md) records the independent arbitrary-precision reference and seed. The native ecosystem verifier also runs the concurrent immutable arithmetic tests under Go’s race detector. Python is not required.
