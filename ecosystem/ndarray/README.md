# ndarray

`ecosystem::ndarray` is an independent GoML module for generic multidimensional
arrays and numerical computation. It implements shared strided views, checked
shapes, broadcasting, reductions, batched matrix products and dense linear
algebra. The implementation uses ordinary GoML and `std::simd`; it has no NumPy,
BLAS, LAPACK or Go numerical-library runtime dependency.

```toml
[dependencies]
"ecosystem::ndarray" = "0.1.0"
```

```goml
use ecosystem::ndarray::{Array, SolveOptions, Error};

fn solve_example() -> Result[Array[f64], Error] {
    let a = Array::from_vec(
        Vec::from_array([2, 2]),
        Vec::from_array([3.0, 1.0, 1.0, 2.0]),
    )?;
    let b = Array::from_vec(Vec::from_array([2]), Vec::from_array([9.0, 8.0]))?;
    a.solve(b, SolveOptions::standard())
}
```

The solution is `[2.0, 3.0]`. All fallible operations return `Result` with a public
`ErrorKind` and message. Invalid dimensions, indices, axes, incompatible shapes,
read-only writes, integer overflow and unsuitable factorization inputs are
recoverable errors.

## Generic storage and views

`Array[T]` stores any element type, including downstream structs. Shape and
strides use `isize`; strides count elements. Logical traversal is row-major.
Rank zero is a scalar containing one element. Any zero dimension produces an
empty array. Constructors copy both the input vector and shape metadata;
`shape()` and `strides()` return snapshots.

| Operation | Behavior |
| --- | --- |
| `from_vec`, `from_vec_with_limits` | Validate shape, element count and limits, then copy values |
| `full`, `full_with_limits`, `scalar` | Construct filled arrays or a rank-zero scalar |
| `ndim`, `size`, `is_empty`, `limits` | Inspect dimensions, count and configured limits |
| `get`, `set`, `get_flat`, `set_flat` | Checked indexing, including negative indices |
| `transpose`, `reversed_axes`, `swap_axes` | Permute axes while sharing storage |
| `slice(axis, SliceSpec)` | Shared slice with clamped bounds and positive/negative nonzero step |
| `index_axis`, `insert_axis`, `squeeze` | Remove an indexed axis, insert a singleton, remove all singleton axes |
| `diagonal(offset, first, second)` | Shared diagonal, appended after remaining axes |
| `broadcast_to` | Shared view using zero strides for repeated values |
| `take(axis, indices)` | Checked gather into independent storage, retaining index order and repetitions |
| `reshape_view` | Share C-contiguous data; reject a required copy |
| `reshape`, `ravel` | Share when C-contiguous, otherwise copy logical values |
| `copy`, `flatten`, `to_vec` | Independent element buffer; flatten always copies |
| `as_slice`, `as_mut_slice` | Borrow contiguous storage; mutable access also checks writability |
| `shares_storage`, `is_contiguous`, `is_writable` | Inspect storage identity and view properties |
| `fill`, `assign` | Modify existing storage; assignment broadcasts the source |
| `map`, `map_indexed`, `zip_map` | Produce arrays of another element type; zip broadcasts both operands |
| `concatenate`, `stack` | Join arrays along an existing or newly inserted axis |
| `reduce_axes` | Fallible custom reduction over selected axes |

Axes accept negative indices. Transposition rejects repeated/missing axes.
Reshape accepts one inferred `-1`; zero-product inference such as `[0, -1]` is
ambiguous and rejected. Empty views use offset zero and zero strides. Singleton
strides need not be one for an array to be contiguous.

`SliceSpec::all()`, `reverse()` and `between(start, stop)` cover common slices.
Explicit specs have `start: Option[isize]`, `stop: Option[isize]` and `step`.
Bounds follow Python slicing, including the distinction between an omitted stop
and `Some(-1)` with a negative step.

A view shares mutations with its source. Broadcasting an axis to length greater
than one makes the resulting view read-only, and further views preserve that
restriction. `copy()` creates writable storage. Assignment snapshots every
source element before writing, so overlapping transposes, reversals and shifted
slices cannot overwrite unread values. `apply_assign` computes a checked result
before modifying the destination, preserving it when arithmetic fails.

Copies are shallow with respect to reference-valued elements. Read-only views
prevent replacing array slots; they do not freeze objects stored in those slots.
Arrays and borrowed slices provide no synchronization. Concurrent mutation
requires caller coordination, including mutations through another view.

## Shapes and limits

`Limits::standard()` allows rank 32 and 16,777,216 logical elements. Custom limits
allow ranks through 1,024 and nonnegative element caps. Shape multiplication is
checked before allocating. A zero dimension permits otherwise large dimensions
without multiplying them; later operations still check output and working-group
sizes. Multi-input operations use the smallest limits of their inputs.
`binary_scalar`, `add_scalar` and `scale` preserve the array's configured limits.

Limits constrain logical element/group counts, not total process memory or the
size of individual `T` values. Matrix multiplication packs inputs, factorization
copies inputs, assignment snapshots its source, and custom reduction materializes
one group at a time. Callbacks may allocate their own data. OS allocation failure
is outside the recoverable shape-error contract.

Broadcasting compares trailing axes: dimensions must match, or one must be one.
A length-one axis may broadcast to zero. See the
[NumPy broadcasting rules](https://numpy.org/doc/2.2/user/basics.broadcasting.html)
for the reference convention used by the interoperability tests.

## Arithmetic and statistics

`Scalar` requires `zero`, `one` and fallible `add`, `subtract`, `multiply`,
`divide`. Its default methods provide batch evaluation, sums, products and dot
products. Downstream implementations inherit these across module boundaries;
the independent consumer supplies a complex-number type. Implementations may
override kernels. The array API rejects a custom batch result of incorrect
length before constructing an array.

Built-in implementations cover `f32`, `f64`, `i64` and `u64`. Integer operations
check overflow, unsigned underflow, division by zero and signed MIN/-1. Floating
operations follow IEEE arithmetic, including infinities and NaNs. Operands have
the same element type; conversion uses `map`.

| API | Scope |
| --- | --- |
| `zeros`, `ones`, `eye(rows, columns, offset)` | Constructors for `T: Scalar` |
| `binary_with(other, op, execution)` | Broadcast add/subtract/multiply/divide with selected execution |
| `add`, `subtract`, `multiply`, `divide` | Broadcast arithmetic with SIMD preference |
| `binary_scalar`, `add_scalar`, `scale` | Checked elementwise scalar operations |
| `apply_assign` | Broadcast arithmetic into the existing shape, with transactional failure |
| `sum`, `product`, `sum_axes`, `product_axes` | Generic scalar reductions, with optional retained dimensions |
| `where_array` | Broadcast boolean selection over generic arrays |
| `all`, `any` | Boolean reductions; empty results are true/false respectively |
| `Array::[f64]::linspace(start, end, count, endpoint)` | Specialized associated constructor; `Array::linspace` infers the owner and the module-level `linspace` remains available |
| `mean`, `variance(ddof)`, `stddev(ddof)`, `norm` | `f64` statistics and scaled Euclidean/Frobenius norm |
| `mean_axes`, `variance_axes`, `min_axes`, `max_axes` | Multi-axis `f64` reductions |
| `minimum`, `maximum`, `argmin`, `argmax` | Global `f64` extrema and flat positions |
| `argmin_axis`, `argmax_axis` | Axis-wise positions as `Array[isize]` |
| `sqrt`, `exp`, `log`, `abs`, `negate`, `pow`, `clip` | Elementwise `f64` operations |
| `is_finite`, `is_nan`, `equal`, `less` | Boolean masks; binary comparisons broadcast |
| `all_close(other, relative, absolute, equal_nan)` | Broadcast tolerance comparison |

Empty sums/products return zero/one. An actual empty group has no mean or
extremum and returns `Empty`; variance requires `0 <= ddof < group_size`.
An output with no groups is empty without invoking a reduction. Empty axis lists
reduce each element as a singleton; sum/product simply copy. Axes may appear in
any order but cannot repeat. Custom groups traverse reduced axes in their
original axis order, with the last axis varying fastest.

Floating sums use Neumaier compensation; `f32` sums accumulate in `f64` before
narrowing. Variance uses Welford updates, and norm uses scaling to avoid squaring
large raw values. These improve common numerical cases without guaranteeing
arbitrary-precision results. Extrema choose the first tie or first NaN. NaNs
propagate through statistics; `all_close` can optionally equate NaNs and always
handles equal infinities explicitly. Sum/dot rounding can differ from NumPy or
between execution modes; use appropriate tolerances.

## Matrix computation

`matmul`/`matmul_with` support vectors, matrices, leading batch broadcasting,
noncontiguous inputs and zero-length contractions. Vector/vector produces a
rank-zero dot product; promoted vector axes are removed from other results.
A zero-length contraction produces zeros. This follows
[NumPy matmul shapes](https://numpy.org/doc/2.2/reference/generated/numpy.matmul.html).
The right input is packed by column and each output uses a dot kernel.

The following decompositions operate on finite `f64` matrices:

| API | Algorithm and result |
| --- | --- |
| `lu(options)` | Partial row pivoting, packed LU and permutation |
| `Lu.lower`, `upper`, `packed`, `permutation` | Independent snapshots, with `P*A = L*U` |
| `solve(rhs, options)`, `Lu.solve(rhs)` | Square solve with vector or multiple-column RHS |
| `inverse(options)`, `Lu.inverse()` | Solve against the identity |
| `determinant`, `slogdet` | Determinant or `(sign, log(abs(det)))`; singular gives zero/negative infinity |
| `cholesky(options)` | Symmetry validation and lower Cholesky for positive-definite matrices |
| `Cholesky.lower`, `solve`, `slogdet` | Factor snapshot, solve and log determinant |
| `qr(options)` | Thin Householder QR for a matrix with at least as many rows as columns |
| `Qr.q`, `r`, `solve` | Thin Q, triangular R and full-column-rank least-squares solve |
| `least_squares(rhs, options)` | QR-based least squares without forming normal equations |

Factors own snapshots and can solve repeated right-hand sides after the original
input changes. A 0x0 matrix has determinant one and valid empty factors.
Decompositions reject nonfinite inputs/intermediates and incompatible RHS shapes.

`SolveOptions::standard()` uses absolute tolerance zero and relative tolerance
`1e-12`. The cutoff is `absolute + relative * max(abs(input))`. LU compares pivot
magnitude, QR compares the current column-tail norm, and Cholesky compares
symmetry differences and updated diagonal values against that cutoff. This is
an explicit rejection policy, not a condition-number estimate. Setting both
tolerances to zero permits any nonzero pivot. Determinant/slogdet use zero
cutoffs automatically; determinant multiplication itself can overflow.

Factorizations accept one matrix per call. There is no batched factorization,
rank-deficient/underdetermined least-squares solution, SVD or eigenvalue solver.
The library is a dense CPU implementation and does not claim BLAS performance.

## SIMD and verification

`Execution::Scalar` selects scalar batch/dot loops; `Execution::Simd` processes
four `f32` or `f64` lanes and finishes a scalar tail. Broadcasting or striding may
select elementwise scalar traversal. Generic and checked integer operations use
the trait defaults. Native kernel dispatch belongs to GoML's portable SIMD
backend, including CPU capability checks and scalar fallback. Dot kernels use
separate multiply/add rounding.

From the repository root:

```sh
just ecosystem-test ndarray
```

The verifier runs module and independent versioned-consumer tests, fresh/cached
build checks, 2,929 deterministic NumPy 2.2.6 reference cases and SIMD checks.
The native consumer tests read [frozen NumPy 2.2.6 reference vectors](../consumers/ndarray/tests/data/README.md), with the original seed and SHA-256-pinned reference-wheel provenance recorded. No Python or NumPy installation is required.
Interoperability includes negative cases and independent QR/LU reconstruction,
not just round trips through the library.

The native GoML verifier checks generated assembly metadata and actual linked symbols,
then builds separate `GOML_SIMD=sse2` and `GOML_SIMD=scalar` targets under the
consumer's `_artifact/`. Both repeat all reference cases and consumer smoke
checks. The current consumer links 10 native kernels by default, five with SSE2
only, and none with scalar-only code generation. This establishes emitted and
linked paths; it is not a speed benchmark or proof that AVX2 executes on a CPU
without AVX2 support.
