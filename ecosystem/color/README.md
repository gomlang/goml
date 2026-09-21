# color

`ecosystem::color` is a checked, allocation-free color math core with CSS input,
perceptual interpolation, gradients, alpha compositing and accessibility helpers.
Algorithms, including angle and cube-root calculations, are implemented in GoML.
Elementary functions use `std::math`. There are no direct Go FFI bindings,
third-party Go dependencies, terminal I/O, global configuration or shared mutable
state.

```toml
[dependencies]
"ecosystem::color" = "0.1.0"
```

```gom
use ecosystem::color;
use ecosystem::color::{Srgb, Gradient, MixSpace, HueDirection, Gamut};

fn theme() -> Result[Vec[Srgb], color::Error] {
    let start = color::parse("rebeccapurple")?;
    let end = color::parse("hsl(160 80% 60%)")?;
    let gradient = Gradient::evenly_spaced(
        Vec::from_array([start, end]),
        MixSpace::Oklch(HueDirection::Shorter),
        Gamut::PreserveHue,
    )?;
    gradient.samples(24)
}
```

## Models and conversion

Every model has private fields, a checked `new(...) -> Result[Model, Error]`
constructor, channel accessors, `alpha()`, `PartialEq` and `Debug`. Every
constructor takes alpha last. Alpha, sRGB channels, HSL saturation/lightness and
HSV saturation/value must be finite and in `[0, 1]`. Floating equality is exact;
conversion roundtrip comparisons should use a tolerance.

| Model | Constructor channels | Conversion methods |
| --- | --- | --- |
| `Srgb` | `red, green, blue, alpha` | `to_linear`, `to_hsl`, `to_hsv`, `to_xyz`, `to_lab`, `to_oklab`, `to_oklch` |
| `LinearRgb` | `red, green, blue, alpha` | `to_srgb(gamut)`, `to_xyz`, `to_oklab` |
| `Hsl` | `hue, saturation, lightness, alpha` | `to_srgb` |
| `Hsv` | `hue, saturation, value, alpha` | `to_srgb` |
| `Xyz` | `x, y, z, alpha` | `to_linear`, `to_srgb(gamut)`, `to_lab` |
| `Lab` | `lightness, a, b, alpha` | `to_xyz`, `to_srgb(gamut)`, `to_lch` |
| `Lch` | `lightness, chroma, hue, alpha` | `to_lab`, `to_srgb(gamut)` |
| `Oklab` | `lightness, a, b, alpha` | `to_linear`, `to_srgb(gamut)`, `to_oklch` |
| `Oklch` | `lightness, chroma, hue, alpha` | `to_oklab`, `to_srgb(gamut)` |

All angle accessors are named `hue()` and return degrees in `[0, 360)`.
Achromatic colors have hue zero. Hues supplied to constructors may wrap and are
limited to magnitude `1e9` degrees. Chroma cannot be negative. For predictable
finite intermediate arithmetic, constructor channels for `LinearRgb` and `Xyz`
are limited to magnitude `1e6`; Lab/LCh channels to `1e4`; Oklab/Oklch channels
to `100`. Conversion output can exceed these input limits: intermediate models
retain extended, signed color coordinates and are not clipped.

XYZ is relative **D65**, with white `(0.9504559270516716, 1,
1.0890577507598784)`. Lab/LCh also use this D65 white, so their values differ
from CSS `lab()`/`lch()` and ICC PCS values using D50. There is no chromatic
adaptation or ICC profile support. Lab lightness uses the usual nominal
`0..100` scale; Oklab lightness the nominal `0..1` scale.

`Srgb::from_rgb8(r, g, b)` and `from_rgba8(r, g, b, a)` are infallible byte
constructors. `black()`, `white()` and `transparent()` are convenience
constructors. `with_alpha(alpha)` validates the replacement alpha.
`to_rgb8()` / `to_rgba8()` return tuples, quantizing to nearest integer with
halfway values rounded upward. RGB byte conversion does not flatten alpha.

`decode_srgb(channel)` and `encode_srgb(channel)` expose the piecewise transfer
functions, including sign-preserving extended-channel behavior. They reject
nonfinite inputs and magnitudes over `1e6`. No integer color math or SIMD-specific
rounding is used.

## Gamut mapping

Conversions into `Srgb` require an explicit `Gamut`:

- `Clip` clamps encoded RGB channels independently to `[0, 1]`.
- `PreserveHue` leaves colors already in gamut unchanged, otherwise searches
  for the largest in-gamut Oklch chroma while preserving lightness and hue.
  The search uses 40 bounded binary-search iterations. Lightness at or below
  zero maps to black; at or above one maps to white. Alpha is retained.

`LinearRgb.is_in_gamut()` uses exact channel bounds. Matrix roundoff at primary
boundaries can therefore report an infinitesimal excursion. Conversions preserve
normal floating-point error; Oklab inverse roundtrips are generally accurate to
about `2e-6` in encoded RGB with these published matrices.

The hue-preserving map is a deterministic chroma-reduction policy. It does not
implement CSS local-MINDE, EdgeSeeker or an ICC rendering intent, and does not
promise a globally nearest perceptual match.

## Parsing and formatting

`parse(text)` accepts case-insensitive CSS names (all 148 opaque names plus
`transparent`), `#rgb`, `#rgba`, `#rrggbb`, `#rrggbbaa`, and `rgb()`/`rgba()`/
`hsl()`/`hsla()` functions. `named(name)` returns `Option[Srgb]` for a name;
`parse_hex(text)` handles only hash-prefixed hex.

Functional input supports legacy commas, modern whitespace and optional slash
alpha. RGB channels accept numbers on the `0..255` scale or percentages; legacy
RGB requires all three channels to use the same unit. HSL saturation/lightness
require percentages. Hue accepts degrees or `deg`, `rad`, `grad`, and `turn`.
Alpha accepts a number or percentage. Decimal exponents are supported; NaN,
infinity, numeric separators, hexadecimal floats, trailing decimal points and
malformed separator combinations return `Error`.

Like CSS, parsed RGB channels, saturation/lightness and alpha are clamped;
constructors instead reject out-of-range values. Hue is wrapped after validation.
`parse` limits input to 4096 bytes, trims surrounding ASCII whitespace and uses
ASCII CSS whitespace between channels. `currentcolor`, `none`, CSS escapes,
comments, relative colors, `calc()`, `color()`, `lab()`, `lch()`, `oklab()` and
`oklch()` syntax are not supported; use typed constructors for those implemented
spaces. It is a color-value parser, not a CSS tokenizer or stylesheet parser.

- `to_hex()` returns lowercase `#rrggbb` and discards alpha.
- `to_hex_alpha()` returns lowercase `#rrggbbaa`.
- `to_css()` emits modern `rgb(... / ...)` using full-precision numeric text;
  parse/format roundtrips are subject to ordinary float rounding.

## Compositing and contrast

`Srgb.composite_over(background)` performs Porter-Duff source-over using
**linear-light** RGB and unassociated alpha. `LinearRgb.composite_over` exposes
the same operation without output gamut conversion. Fully transparent output
uses zero RGB. These are physical linear-light compositing operations; they do
not emulate applications that blend directly in encoded sRGB.

`relative_luminance()` returns the WCAG sRGB luminance (ignoring alpha).
`contrast_ratio(other)` accepts only opaque colors and returns a ratio in
`[1, 21]`. `contrast_on(background)` composites the receiver over an opaque
background before comparison. A translucent background is rejected because the
result depends on an unspecified surface behind it. `readable_foreground()`
chooses black or white with the greater contrast against an opaque receiver.
These helpers do not account for font weight, display conditions or non-color
accessibility requirements.

## Mixing and gradients

`mix(a, b, fraction, space, gamut)` validates a finite fraction in `[0, 1]`.
`MixSpace` supports `Srgb`, `LinearRgb`, `Lab`, `Oklab`,
`Hsl(HueDirection)` and `Oklch(HueDirection)`. Rectangular components and
non-hue cylindrical components are premultiplied by alpha before interpolation,
then unpremultiplied. Alpha weights are normalized after scaling by the larger
endpoint alpha, retaining chromatic components for subnormal positive alpha.
An interpolated alpha that rounds to zero produces zero non-hue components.
Hue interpolates using the original fraction, independently
of alpha. Fractions zero and one preserve the exact supplied endpoint.

`HueDirection::{Shorter, Longer, Increasing, Decreasing}` selects the path around
the hue circle; equal hues with `Longer` make a complete increasing turn. Exactly
180-degree ties retain their signed difference. An achromatic endpoint borrows
the other hue before path selection; the chroma threshold is `1e-7` (HSL uses
`saturation * min(lightness, 1 - lightness)`). Transparent endpoints can still
have chromatic hue; transparency itself does not make hue powerless.

`Gradient::new(stops, space, gamut, extension)` accepts 2 through 65536 sorted
`Stop { position, color }` values, with positions in `[0, 1]`. It snapshots the
stop vector so caller mutations do not change the gradient. Duplicate positions
make hard edges; the rightmost stop wins exactly at the edge. Positions before
the first stop or after the last stop return the respective edge color.

`Extension::{Clamp, Repeat, Mirror}` controls positions outside `[0, 1]`, with
repeat mapping integer positions to zero and mirror preserving alternating
endpoints. `sample(position)` accepts finite magnitudes through `1e9` and uses
binary search. `samples(count)` returns 0 through 65536 evenly spaced samples;
one sample means position zero. `evenly_spaced(colors, space, gamut)` constructs
a clamped gradient. `len()` returns the number of stops.

## Color difference

`Lab.delta_e_76(other)` calculates Euclidean CIELAB distance.
`Lab.delta_e_2000(other)` implements CIEDE2000 with standard unit parametric
factors. `Oklab.delta_e(other)` calculates Euclidean Oklab distance. All three
ignore alpha and compare coordinates under the same white point. Oklab distance
has a different numerical scale from Lab distance.

## Verification and sources

From the repository root, `just ecosystem-test color` formats/checks,
tests the library and independent consumer, verifies cached builds, runs a theme
palette example, and runs 4,659 independent reference vectors in the consumer’s ordinary GoML tests. No Python runtime or network access is needed.

The GoML suite covers constructor errors, transfer thresholds, a 343-color
conversion grid, extended signed channels, gamut mapping, premultiplied alpha,
CSS grammar failures, hue paths, gradient snapshots and hard edges, contrast,
color difference, and subnormal-alpha mixing across all six interpolation spaces. The consumer uses only the published dependency interface.
The reference vectors record comparisons of deterministic random colors with `colorsys`, an
independent rational-matrix implementation, linear compositing and mixing;
it checks all CSS names, all 34 published CIEDE2000 test pairs, and 810 alpha-boundary
mixtures against high-precision Decimal arithmetic, including exact subnormal alpha. [Vector provenance](../consumers/color/tests/data/README.md) identifies the original sources and seed.

Mathematical definitions, reference matrices and data provenance:

- [W3C CSS Color 4](https://www.w3.org/TR/2026/CRD-css-color-4-20260913/):
  sRGB/XYZ conversion, CSS syntax, hue interpolation, and the 148 named-color
  values in `named-colors.json` / `named.gom`.
- [Björn Ottosson's Oklab specification](https://bottosson.github.io/posts/oklab/):
  the 2021 Oklab forward/inverse matrices.
- [Sharma, Wu and Dalal, CIEDE2000 implementation notes and supplementary data](https://hajim.rochester.edu/ece/sites/gsharma/ciede2000/):
  `ciede2000-testdata.txt` contains their 34 published reference pairs and
  expected distances. No reference MATLAB code is included.
- [W3C WCAG relative luminance](https://www.w3.org/WAI/WCAG22/Understanding/contrast-minimum.html):
  the luminance and contrast ratio definitions.
