import math
from pathlib import Path
import random
import subprocess


ROOT = Path(__file__).resolve().parents[1]
DIGITS = "0123456789abcdefghijklmnopqrstuvwxyz"


def radix(value, base):
    if value == 0:
        return "0"
    remaining = abs(value)
    result = []
    while remaining:
        remaining, digit = divmod(remaining, base)
        result.append(DIGITS[digit])
    return ("-" if value < 0 else "") + "".join(reversed(result))


def signed_bytes(value):
    length = max(1, (value.bit_length() + 8) // 8)
    result = value.to_bytes(length, "big", signed=True)
    while len(result) > 1 and ((result[0] == 0 and result[1] < 128) or (result[0] == 255 and result[1] >= 128)):
        result = result[1:]
    return result


def expected_arithmetic(a, b, shift, base, exponent, modulus):
    division = ["error"] * 4
    if b:
        q = abs(a) // abs(b) * (-1 if (a < 0) != (b < 0) else 1)
        r = a - q * b
        euclidean_remainder = a % abs(b)
        euclidean_quotient = (a - euclidean_remainder) // b
        division = [str(q), str(r), str(euclidean_quotient), str(euclidean_remainder)]
    values = [str(a + b), str(a - b), str(a * b)] + division
    values += [str(a & b), str(a | b), str(a ^ b), str(~a)]
    values += [str(a << shift), str(a >> shift)] if shift >= 0 else ["error", "error"]
    values += [str(math.gcd(a, b)), str(math.lcm(a, b)), radix(a, base), signed_bytes(a).hex()]
    values += [abs(a).to_bytes(max(1, (abs(a).bit_length() + 7) // 8), "big").hex()]
    values += [str(pow(a, exponent, modulus)) if modulus else "error", str((a > b) - (a < b))]
    return "\t".join(values)


def main():
    rng = random.Random(781963120)
    cases = []
    expected = []
    edges = [0, 1, -1]
    for bit in [7, 8, 15, 16, 31, 32, 33, 63, 64, 65, 95, 96, 127, 128, 255, 256, 511, 1024, 2048]:
        for offset in [-1, 0, 1]:
            edges.extend([(1 << bit) + offset, -((1 << bit) + offset)])
    pairs = [(value, rng.choice(edges)) for value in edges]
    for _ in range(800):
        a = rng.getrandbits(rng.choice([32, 64, 96, 128, 256, 512, 1024, 2048, 4096]))
        b = rng.getrandbits(rng.choice([32, 64, 96, 128, 256, 512, 1024]))
        pairs.append((a * rng.choice([-1, 1]), b * rng.choice([-1, 1])))
    for _ in range(150):
        divisor = rng.getrandbits(rng.randrange(33, 1024)) | (1 << 32)
        quotient = rng.getrandbits(rng.randrange(1, 512))
        remainder = rng.choice([0, 1, divisor - 1, rng.randrange(divisor)])
        pairs.append((divisor * quotient + remainder, divisor))
    for a, b in pairs:
        shift = rng.choice([-1, 0, 1, 7, 31, 32, 33, 63, 64, 65, 127, 255, 4097])
        base = rng.randrange(2, 37)
        exponent = rng.randrange(0, 150)
        modulus = rng.getrandbits(rng.choice([0, 1, 32, 64, 128, 256]))
        cases.append(f"arithmetic\t{a}\t{b}\t{shift}\t{base}\t{exponent}\t{modulus}")
        expected.append(expected_arithmetic(a, b, shift, base, exponent, modulus))
    for base in range(2, 37):
        for _ in range(12):
            value = rng.getrandbits(rng.randrange(0, 4096)) * rng.choice([-1, 1])
            text = radix(value, base)
            text = text.upper() if rng.randrange(2) else text
            if value >= 0 and rng.randrange(2):
                text = "+" + text
            cases.append(f"parse\t{text}\t{base}")
            expected.append(str(value))
    for text, base in [("", 10), ("+", 10), ("-", 10), ("--1", 10), ("1_2", 10), ("0x10", 16), (" 1", 10), ("1 ", 10), ("１２", 10), ("2", 2), ("z", 35), ("1", 1), ("1", 37), ("-0000", 10), ("+0000", 10)]:
        cases.append(f"parse\t{text}\t{base}")
        expected.append("0" if text in ["-0000", "+0000"] else "error")
    for _ in range(180):
        value = rng.randrange(-100000, 100001)
        exponent = rng.randrange(0, 100)
        cases.append(f"pow\t{value}\t{exponent}")
        expected.append(str(pow(value, exponent)))
    for value in [0, 1, 2, -1, -2]:
        cases.append(f"pow\t{value}\t1000001")
        expected.append("error")
    cases.append("pow\t2\t65536")
    expected.append("error")
    for _ in range(250):
        data = rng.randbytes(rng.randrange(0, 512))
        cases.append(f"bytes\t{data.hex()}")
        expected.append(f"{int.from_bytes(data, 'big', signed=True)}\t{int.from_bytes(data, 'big')}")
    binary = ROOT / "consumers/bigint/_artifact/bin/bigint"
    result = subprocess.run([str(binary), "--oracle"], input="\n".join(cases) + "\n", capture_output=True, text=True, timeout=90)
    if result.returncode:
        raise AssertionError(result.stderr)
    actual = result.stdout.splitlines()
    if len(actual) != len(expected):
        raise AssertionError(f"expected {len(expected)} output rows, received {len(actual)}")
    for index, (got, wanted) in enumerate(zip(actual, expected)):
        if got != wanted:
            raise AssertionError(f"case {index}: {cases[index]}\nGoML: {got}\nPython: {wanted}")
    print(f"Bigint: {len(cases)} deterministic Python int oracle cases passed, including {len(pairs)} multi-operation arithmetic rows")


if __name__ == "__main__":
    main()
