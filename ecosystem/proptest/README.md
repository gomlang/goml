# ecosystem::proptest

Deterministic property testing with composable generators and lazy shrink trees.
The runner returns a structured failure; the application decides how to report
it or integrate it with `std::testing`.

```goml
use ecosystem::proptest;

fn check_addition() -> Result[Option[proptest::Failure[isize]], string] {
    let values = proptest::integers(-1000, 1000)?;
    proptest::check(values, proptest::Config::new(42), |value: isize| {
        value + 0 == value
    })
}
```

## Strategies and generators

`Strategy` has an associated `Item` type and a `generate(seed, size)` method.
Downstream packages can implement it directly. `Generator[T]` implements that
trait using a function value. Its combinators are `map`, `filter`, `zip` and
`flat_map`; all retain shrinking. A `Sample[T]` holds a value and a lazy function
returning candidate samples. `Sample::leaf` makes a value with no smaller cases.

| Generator | Contract |
| --- | --- |
| `integers(min, max)` | Half-open machine-integer range, including ranges wider than signed `isize` |
| `just(value)` | One constant value |
| `booleans()` | Boolean values; true shrinks to false |
| `select(values)` | Snapshots choices; shrinks toward the first |
| `one_of(generators)` / `weighted(entries)` | Alternative generators with positive checked weights |
| `optional(generator)` | Optional values; present values shrink toward absence |
| `vectors(item, min, max)` | Inclusive length bounds, guided by the runner's size |
| `characters()` | Unicode scalar values, excluding surrogates |
| `strings(characters, min, max)` | Character-count bounds with Unicode-preserving shrinking |
| `bytes(min, max)` | Byte vectors |
| `finite_floats(min, max)` | Finite values within inclusive bounds |
| `recursive(leaf, branch, depth)` | Recursively composed generators with a construction depth limit |

Integer shrinking moves toward zero, or toward the nearest valid bound when zero
is excluded, and probes progressively smaller steps to find failure boundaries.
Vectors shrink by reducing length and then shrinking individual elements.
Filtering retains only accepted shrink candidates and searches rejected branches
for acceptable descendants, bounded to 1,024 visits per expansion. Dependent
`flat_map` shrinking regenerates inner values after shrinking the outer value,
preserving generation dependencies. Alternative choices shrink within their
selected generator; they do not switch arbitrarily to an unrelated generator.

Invalid ranges and exhausted rejection budgets return errors. Maximum vector
length is 1,000,000 and recursive construction depth is at most 64. Floating
shrinking has a depth limit of 64. These generators are intended for testing;
their modulo-based selections do not promise unbiased statistical sampling.

## Running and replaying

`Config::new(seed)` selects 100 cases, size 64 and 10,000 shrink attempts. All
fields are public. `check` returns `Ok(None)` on success and `Ok(Some(failure))`
on a counterexample. Generator errors return `Err`. A failure records the exact
case seed, case index, initial value, final candidate and attempted shrink count.
`replay(strategy, failure.seed, size)` reproduces the initial case. Preserve the
generator definition and size as well as the seed. `Failure::report` is available
when the item implements `ToString`.

The final candidate is locally minimal under the visited shrink tree and budget;
it is not a proof of global minimality. Properties and custom generators should
be deterministic and should not mutate shared sample values. Custom shrink trees
should make progress. The attempt budget bounds runner calls to the property but
cannot interrupt a custom generator or property that never returns. Panics are
not converted into property failures.

Closures used with an associated `Strategy::Item` sometimes need explicit
parameter types for numeric operators in the current GoML checker. The examples
and consumer tests exercise this form.

## Validation

```sh
python3 ecosystem/verify.py proptest
```

Tests include exact integer failure boundaries, wide signed ranges, filtering,
dependent invariants during shrinking, recursive data, valid Unicode, finite
floats, invalid configurations and exact seed replay. A separate module implements
its own associated-type strategy and captures state in a property closure.
