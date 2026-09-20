import colorsys
from decimal import Decimal, localcontext
import json
import math
from pathlib import Path
import random
import subprocess


ROOT = Path(__file__).resolve().parent
BINARY = ROOT.parent / 'consumers' / 'color' / '_artifact' / 'bin' / 'color'


def decode(value):
    return value / 12.92 if value <= 0.04045 else ((value + 0.055) / 1.055) ** 2.4


def encode(value):
    return value * 12.92 if value <= 0.0031308 else 1.055 * value ** (1 / 2.4) - 0.055


def matrix(rows, vector):
    return [sum(a * b for a, b in zip(row, vector)) for row in rows]


def details(rgba):
    r, g, b, alpha = rgba
    linear = list(map(decode, (r, g, b)))
    h, l, s = colorsys.rgb_to_hls(r, g, b)
    hh, ss, v = colorsys.rgb_to_hsv(r, g, b)
    xyz = matrix([
        [506752 / 1228815, 87881 / 245763, 12673 / 70218],
        [87098 / 409605, 175762 / 245763, 12673 / 175545],
        [7918 / 409605, 87881 / 737289, 1001167 / 1053270],
    ], linear)
    white = [0.3127 / 0.3290, 1, (1 - 0.3127 - 0.3290) / 0.3290]
    f = [x ** (1 / 3) if x > (6 / 29) ** 3 else x / (3 * (6 / 29) ** 2) + 4 / 29 for x in [v / w for v, w in zip(xyz, white)]]
    lab = [116 * f[1] - 16, 500 * (f[0] - f[1]), 200 * (f[1] - f[2])]
    lms = matrix([
        [0.4122214708, 0.5363325363, 0.0514459929],
        [0.2119034982, 0.6806995451, 0.1073969566],
        [0.0883024619, 0.2817188376, 0.6299787005],
    ], linear)
    oklab = matrix([
        [0.2104542553, 0.7936177850, -0.0040720468],
        [1.9779984951, -2.4285922050, 0.4505937099],
        [0.0259040371, 0.7827717662, -0.8086757660],
    ], [x ** (1 / 3) for x in lms])
    luminance = sum(x * y for x, y in zip(linear, [0.2126, 0.7152, 0.0722]))
    return list(rgba) + linear + [h * 360, s, l, hh * 360, ss, v] + xyz + lab + oklab + [luminance]


def quantized(rgba):
    return '#' + ''.join(f'{min(255, max(0, math.floor(x * 255 + 0.5))):02x}' for x in rgba)


def main():
    cases = []
    expected = []

    def add(op, rgba, text='', values=None, full=True):
        cases.append({'op': op, 'text': text, 'values': list(rgba) if values is None else values})
        expected.append({'hex': quantized(rgba), 'values': details(rgba) if full else rgba})

    for name, value in json.loads((ROOT / 'named-colors.json').read_text()).items():
        rgba = [int(value[i:i + 2], 16) / 255 for i in (0, 2, 4)] + [1]
        add('parse', rgba, name.upper())
        add('parse', rgba, '#' + value)
    add('parse', [0, 0, 0, 0], 'transparent')
    for text, rgba in [
        ('#1234', [17 / 255, 34 / 255, 51 / 255, 68 / 255]),
        ('rgb(255 0 128 / 25%)', [1, 0, 128 / 255, .25]),
        ('rgba(100%, 0%, 50%, .5)', [1, 0, .5, .5]),
        ('hsl(-120deg 100% 50%)', [0, 0, 1, 1]),
        ('hsl(.5turn 100% 50% / 50%)', [0, 1, 1, .5]),
        ('rgb(-2 260 0 / -1)', [0, 1, 0, 0]),
    ]:
        add('parse', rgba, text)
    rng = random.Random(20260922)
    for _ in range(1500):
        rgba = [rng.random() for _ in range(4)]
        add('convert', rgba)
    for _ in range(500):
        hue = rng.uniform(-720, 720)
        saturation, lightness, alpha = [rng.random() for _ in range(3)]
        rgb = list(colorsys.hls_to_rgb((hue % 360) / 360, lightness, saturation))
        text = f'hsl({hue:.15g}deg {saturation * 100:.15g}% {lightness * 100:.15g}% / {alpha:.15g})'
        add('parse', rgb + [alpha], text)
    for _ in range(500):
        fg = [rng.random() for _ in range(4)]
        bg = [rng.random() for _ in range(4)]
        alpha = fg[3] + bg[3] * (1 - fg[3])
        result = [encode((decode(x) * fg[3] + decode(y) * bg[3] * (1 - fg[3])) / alpha) for x, y in zip(fg[:3], bg[:3])] + [alpha]
        add('composite', result, values=fg + bg, full=False)
    for space in ['srgb', 'linear']:
        for _ in range(500):
            a = [rng.random() for _ in range(4)]
            b = [rng.random() for _ in range(4)]
            t = rng.random()
            alpha = a[3] * (1 - t) + b[3] * t
            aa, bb = a[:3], b[:3]
            if space == 'linear':
                aa, bb = list(map(decode, aa)), list(map(decode, bb))
            result = [(x * a[3] * (1 - t) + y * b[3] * t) / alpha for x, y in zip(aa, bb)]
            if space == 'linear':
                result = list(map(encode, result))
            add('mix', result + [alpha], space, values=a + b + [t], full=False)
    alpha_boundaries = [0.0, math.ulp(0.0), math.ldexp(1.0, -1073), math.ldexp(1.0, -1071), math.ldexp(1.0, -1067), math.ldexp(1.0, -1022), math.ldexp(1.0, -1000), math.ldexp(1.0, -500), 1.0]
    with localcontext() as decimal_context:
        decimal_context.prec = 1100
        for space in ['srgb', 'linear']:
            for first_alpha in alpha_boundaries:
                for second_alpha in alpha_boundaries:
                    for fraction in [0.0, 0.25, 0.5, 0.75, 1.0]:
                        a = [0.125, 0.5, 0.75, first_alpha]
                        b = [0.875, 0.25, 0.5, second_alpha]
                        t = Decimal.from_float(fraction)
                        left = Decimal.from_float(first_alpha) * (1 - t)
                        right = Decimal.from_float(second_alpha) * t
                        alpha = left + right
                        if fraction == 0:
                            rgba = a
                        elif fraction == 1:
                            rgba = b
                        elif float(alpha) == 0:
                            rgba = [0.0, 0.0, 0.0, 0.0]
                        else:
                            aa, bb = a[:3], b[:3]
                            if space == 'linear':
                                aa, bb = list(map(decode, aa)), list(map(decode, bb))
                            channels = [float((Decimal.from_float(x) * left + Decimal.from_float(y) * right) / alpha) for x, y in zip(aa, bb)]
                            if space == 'linear':
                                channels = list(map(encode, channels))
                            rgba = channels + [float(alpha)]
                        add('mix', rgba, space, values=a + b + [fraction], full=False)
                        expected[-1].pop('hex')
                        expected[-1]['exact_alpha'] = rgba[3]
    for line in (ROOT / 'ciede2000-testdata.txt').read_text().splitlines():
        values = list(map(float, line.split()))
        cases.append({'op': 'delta', 'text': '', 'values': values[:6]})
        expected.append({'values': [values[6], values[6], math.dist(values[:3], values[3:6])], 'tolerance': 0.00005})
    for text in ['', '#12', '#ggg', 'currentColor', 'rgb(NaN 0 0)', 'rgb(1, 2%, 3)', 'rgb(1 2 3 // 1)', 'rgb(1e999 0 0)', 'rgb(1_0 0 0)', 'hsl(0 1 1)', 'hsl(1foo 10% 10%)', 'lab(1 2 3)']:
        cases.append({'op': 'parse', 'text': text, 'values': []})
        expected.append(None)
    result = subprocess.run([str(BINARY), '--json'], input=json.dumps(cases), capture_output=True, text=True, check=True, timeout=60)
    actual = json.loads(result.stdout)
    if len(actual) != len(expected):
        raise AssertionError(f'reply count: {len(actual)} != {len(expected)}')
    for i, (got, want) in enumerate(zip(actual, expected)):
        if want is None:
            if got['ok']:
                raise AssertionError(f'accepted invalid input: {cases[i]}')
            continue
        if not got['ok'] or ('hex' in want and got['hex'] != want['hex']):
            raise AssertionError(f'case {i}: {cases[i]}\nactual {got}\nexpected {want}')
        if len(got['values']) != len(want['values']):
            raise AssertionError(f'case {i}: wrong component count')
        if 'exact_alpha' in want and got['values'][3] != want['exact_alpha']:
            raise AssertionError(f'case {i}: alpha {got["values"][3]} != {want["exact_alpha"]}')
        tolerance = want.get('tolerance', 2e-10)
        for component, (a, b) in enumerate(zip(got['values'], want['values'])):
            if not math.isfinite(a) or abs(a - b) > tolerance:
                raise AssertionError(f'case {i}, component {component}: {cases[i]}\nactual {a}, expected {b}, tolerance {tolerance}')
    print(f'color interoperability: {len(cases)} cases pass Python colorsys, independent D65 matrices, linear compositing/premultiplied mixing, all 148 CSS names, all 34 published CIEDE2000 pairs, and 810 high-precision Decimal alpha-boundary mixtures')


if __name__ == '__main__':
    main()
