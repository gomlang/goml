import bisect
import json
from pathlib import Path
import random
import re
import subprocess


ROOT = Path(__file__).resolve().parents[1]
BINARY = ROOT / "consumers" / "rope" / "_artifact" / "bin" / "rope"


def boundaries(text):
    result = [0]
    for scalar in text:
        result.append(result[-1] + len(scalar.encode("utf-8")))
    return result


def utf16_length(text):
    return len(text.encode("utf-16-le")) // 2


def describe(text, detailed=False):
    starts = [0]
    for match in re.finditer(r"\r\n|\r|\n", text):
        starts.append(match.end())
    ends = starts[1:] + [len(text)]
    lines = [text[start:end] for start, end in zip(starts, ends)]
    contents = [re.sub(r"(?:\r\n|\r|\n)$", "", line) for line in lines]
    offsets = boundaries(text)
    result = {
        "text": text,
        "bytes": offsets[-1],
        "scalars": len(text),
        "utf16": utf16_length(text),
        "lines": lines,
        "contents": contents,
        "starts": [offsets[start] for start in starts],
    }
    if detailed:
        positions = []
        units = [0]
        for scalar in text:
            units.append(units[-1] + utf16_length(scalar))
        for index, offset in enumerate(offsets):
            line = bisect.bisect_right(starts, index) - 1
            position = None
            if index <= starts[line] + len(contents[line]):
                position = [line, units[index] - units[starts[line]]]
            positions.append([offset, index, units[index], line, position])
        offset_set = set(offsets)
        unit_set = set(units)
        result["positions"] = positions
        result["invalid_bytes"] = [index for index in range(offsets[-1] + 1) if index not in offset_set]
        result["invalid_utf16"] = [index for index in range(units[-1] + 1) if index not in unit_set]
    return result


def apply(text, op):
    kind = op["kind"]
    start, end = op.get("start", 0), op.get("end", 0)
    inserted = op.get("text", "")
    encoded = text.encode("utf-8")
    offsets = boundaries(text)
    if kind in ("replace", "slice"):
        if start not in offsets or end not in offsets or start > end:
            return False, text
        if kind == "replace":
            return True, (encoded[:start] + inserted.encode() + encoded[end:]).decode()
        return True, encoded[start:end].decode()
    if kind == "insert":
        return (True, text[:start] + inserted + text[start:]) if 0 <= start <= len(text) else (False, text)
    if kind == "remove":
        return (True, text[:start] + text[end:]) if 0 <= start <= end <= len(text) else (False, text)
    if kind in ("split", "rotate"):
        if start not in offsets:
            return False, text
        return True, text if kind == "split" else (encoded[start:] + encoded[:start]).decode()
    if kind == "concat":
        return True, text + inserted
    if kind == "compact":
        return True, text
    if kind == "build":
        return True, "".join(op["parts"])
    raise AssertionError(kind)


def check_report(actual, text, detailed, context):
    expected = describe(text, detailed)
    for key, value in expected.items():
        if actual.get(key) != value:
            raise AssertionError(f"{context}: {key} differs; expected {value!r}, got {actual.get(key)!r}")
    chunks, height = actual["chunks"], actual["height"]
    if height > 2 * (chunks + 1).bit_length():
        raise AssertionError(f"{context}: tree depth {height} for {chunks} chunks")
    if bool(chunks) != bool(text):
        raise AssertionError(f"{context}: empty/chunk mismatch")


def cases():
    rng = random.Random(20260920)
    result = []
    initial = ["", "\n", "\r", "\r\n", "\r\r\n\n", "中😀é\u2028\u0085\x00", "x" * 1023 + "\r\n", "x" * 1022 + "😀\r\n", "a\r\n" * 1200]
    for text in initial:
        ops = [{"kind": "split", "start": at} for at in boundaries(text)[::max(1, len(text) // 16)]]
        ops.extend([
            {"kind": "replace", "start": -1, "end": 0, "text": "x"},
            {"kind": "slice", "start": 1, "end": 0},
            {"kind": "insert", "start": 9223372036854775807, "text": "x"},
            {"kind": "compact"},
        ])
        result.append({"text": text, "operations": ops})
    for at in range(1018, 1027):
        text = "x" * at + "\r\n😀中\r\n"
        result.append({"text": text, "operations": [
            {"kind": "replace", "start": at + 1, "end": at + 1, "text": "😀"},
            {"kind": "replace", "start": at + 1, "end": at + 5, "text": ""},
            {"kind": "rotate", "start": at + 1},
            {"kind": "compact"},
        ]})
    alphabet = ["a", "b", "é", "中", "😀", "𝄞", "\r", "\n", "\r\n", "\x00", "é", "\u2028"]
    for _ in range(48):
        text = "".join(rng.choice(alphabet) for _ in range(rng.randrange(120, 380)))
        current = text
        ops = []
        for step in range(65):
            offsets = boundaries(current)
            kind = rng.choice(["replace", "replace", "insert", "remove", "split", "rotate", "concat", "compact"])
            start = rng.randrange(len(current) + 1)
            end = min(len(current), start + rng.randrange(8))
            inserted = "".join(rng.choice(alphabet) for _ in range(rng.randrange(7)))
            if kind == "replace":
                op = {"kind": kind, "start": offsets[start], "end": offsets[end], "text": inserted}
            elif kind in ("insert", "remove"):
                op = {"kind": kind, "start": start, "end": end, "text": inserted}
            elif kind in ("split", "rotate"):
                op = {"kind": kind, "start": offsets[start]}
            elif kind == "concat":
                op = {"kind": kind, "text": inserted}
            else:
                op = {"kind": kind}
            if step % 13 == 0:
                op = {"kind": "replace", "start": rng.randrange(-2, len(current.encode()) + 3), "end": rng.randrange(-2, len(current.encode()) + 3), "text": inserted}
            ops.append(op)
            _, current = apply(current, op)
        result.append({"text": text, "operations": ops})
    parts = ["a" * 1023, "😀", "\r", "\n", "中" * 600, "\r", "x", "\n"]
    result.append({"text": "before", "operations": [{"kind": "build", "parts": parts}, {"kind": "slice", "start": 1023, "end": 1029}]})
    return result


def main():
    if not BINARY.is_file():
        raise RuntimeError("build the rope consumer with ecosystem/verify.py first")
    inputs = cases()
    executed = subprocess.run([str(BINARY), "oracle"], input=json.dumps(inputs, ensure_ascii=False), text=True, capture_output=True, check=True, timeout=120)
    outputs = json.loads(executed.stdout)
    if len(outputs) != len(inputs):
        raise AssertionError("case count differs")
    edits = 0
    for index, (case, output) in enumerate(zip(inputs, outputs)):
        current = case["text"]
        if len(output["steps"]) != len(case["operations"]) + 1:
            raise AssertionError(f"case {index}: result count differs")
        check_report(output["steps"][0], current, False, f"case {index} initial")
        for step, (op, actual) in enumerate(zip(case["operations"], output["steps"][1:])):
            ok, current = apply(current, op)
            if actual["ok"] != ok:
                raise AssertionError(f"case {index} step {step}: acceptance mismatch for {op}")
            check_report(actual["value"], current, False, f"case {index} step {step}")
            edits += 1
        check_report(output["final"], current, True, f"case {index} final")
    print(f"rope Python Unicode/flat-string oracle: {len(inputs)} cases, {edits} edits, all final UTF-8/UTF-16/line boundaries passed")


if __name__ == "__main__":
    main()
