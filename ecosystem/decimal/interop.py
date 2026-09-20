import decimal
import json
from pathlib import Path
import random
import subprocess


ROOT = Path(__file__).resolve().parents[1]
BINARY = ROOT / "consumers" / "decimal" / "_artifact" / "bin" / "decimal"
ROUNDINGS = {
    "half_even": decimal.ROUND_HALF_EVEN,
    "half_up": decimal.ROUND_HALF_UP,
    "half_down": decimal.ROUND_HALF_DOWN,
    "up": decimal.ROUND_UP,
    "down": decimal.ROUND_DOWN,
    "ceiling": decimal.ROUND_CEILING,
    "floor": decimal.ROUND_FLOOR,
}


def reference(case):
    context = decimal.Context(prec=case["precision"], rounding=ROUNDINGS[case["rounding"]])
    left = decimal.Decimal(case["a"])
    right = decimal.Decimal(case["b"])
    try:
        if case["op"] == "add":
            value = context.add(left, right)
        elif case["op"] == "sub":
            value = context.subtract(left, right)
        elif case["op"] == "mul":
            value = context.multiply(left, right)
        elif case["op"] == "div":
            value = context.divide(left, right)
        elif case["op"] == "round":
            value = context.plus(left)
        elif case["op"] == "quantize":
            target = decimal.Decimal((0, (1,), -case["scale"]))
            value = context.quantize(left, target)
        else:
            raise AssertionError(case["op"])
        return {"ok": True, "value": str(value), "rounded": context.flags[decimal.Rounded], "inexact": context.flags[decimal.Inexact]}
    except decimal.DecimalException:
        return {"ok": False, "value": ""}


def cases():
    result = []
    for rounding in ROUNDINGS:
        for value in ["0", "-0", "1.25", "-1.25", "1.35", "-1.35", "9.95", "-9.95", "1.2300", "0.000050", "-0.000050", "999999999999999999999999.5"]:
            for scale in [-2, 0, 1, 3, 7]:
                result.append({"op": "quantize", "a": value, "b": "1", "precision": 28, "scale": scale, "rounding": rounding})
        for value in ["1", "-1", "2", "10", "0", "9999999999999999999999999999"]:
            for divisor in ["0", "3", "6", "-7", "0.0003", "10000000000000000000000000001"]:
                result.append({"op": "div", "a": value, "b": divisor, "precision": 17, "scale": 0, "rounding": rounding})
    rng = random.Random(20260921)
    for _ in range(2400):
        left = str(rng.randrange(-(10 ** rng.randrange(1, 41)), 10 ** rng.randrange(1, 41))) + "e" + str(rng.randrange(-24, 13))
        right = str(rng.randrange(-(10 ** rng.randrange(1, 41)), 10 ** rng.randrange(1, 41))) + "e" + str(rng.randrange(-24, 13))
        result.append({
            "op": rng.choice(["add", "sub", "mul", "div", "round", "quantize"]),
            "a": left, "b": right, "precision": rng.randrange(1, 51),
            "scale": rng.randrange(-15, 31), "rounding": rng.choice(list(ROUNDINGS)),
        })
    return result


def main():
    inputs = cases()
    expected = [reference(case) for case in inputs]
    process = subprocess.run([str(BINARY), "--json"], input=json.dumps(inputs), capture_output=True, text=True, check=True, timeout=120)
    actual = json.loads(process.stdout)
    if len(actual) != len(expected):
        raise AssertionError(f"case count {len(actual)} != {len(expected)}")
    for index, (left, right) in enumerate(zip(actual, expected)):
        equal = left["ok"] == right["ok"]
        if equal and left["ok"]:
            equal = decimal.Decimal(left["value"]) == decimal.Decimal(right["value"])
            equal = equal and left["rounded"] == right["rounded"] and left["inexact"] == right["inexact"]
        if not equal:
            raise AssertionError(f"case {index}: {inputs[index]!r}\nGoML: {left!r}\nPython: {right!r}")
    print(f"decimal interoperability: {len(inputs)} arithmetic/rounding cases agree with Python decimal values, errors, Rounded and Inexact across seven rounding modes")


if __name__ == "__main__":
    main()
