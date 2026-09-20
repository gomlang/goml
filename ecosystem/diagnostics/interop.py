from pathlib import Path
import random
import re
import subprocess


ROOT = Path(__file__).resolve().parents[1]
BINARY = ROOT / "consumers/diagnostics/_artifact/bin/diagnostics"
TOKENS = ["a", "Z", " ", "中", "é", "e\u0301", "👩‍💻", "\t", "\x1b"]
WIDTHS = {"中": 2, "👩‍💻": 2, "e\u0301": 1}


def encoded(value):
    return value.encode().hex()


def boundary_offsets(value):
    result, offset = [0], 0
    for character in value:
        offset += len(character.encode())
        result.append(offset)
    return result


def physical_lines(value):
    source = value.encode()
    result, start, at = [], 0, 0
    while at < len(source):
        if source[at] in (10, 13):
            end = at
            at += 1
            if source[end] == 13 and at < len(source) and source[at] == 10:
                at += 1
            result.append((start, end))
            start = at
        else:
            at += 1
    result.append((start, at))
    return result


def location(value, offset, tab):
    if offset not in boundary_offsets(value) or not 1 <= tab <= 256:
        return "error"
    lines = physical_lines(value)
    number = max(index for index, (start, _) in enumerate(lines) if start <= offset)
    start, end = lines[number]
    selected = min(offset, end)
    source = value.encode()[start:end].decode()
    at, column = start, 0
    while source and at < selected:
        token = next((item for item in ("👩‍💻", "e\u0301") if source.startswith(item)), source[0])
        if at + len(token.encode()) > selected:
            break
        column += tab - column % tab if token == "\t" else WIDTHS.get(token, 1)
        at += len(token.encode())
        source = source[len(token):]
    return f"{number},{selected - start},{column}"


def apply(value, edits, limit):
    boundaries = set(boundary_offsets(value))
    for start, end, _ in edits:
        if start > end or start not in boundaries or end not in boundaries:
            return "error"
    for index, (start, end, _) in enumerate(edits):
        for other_start, other_end, _ in edits[index + 1:]:
            if start == end and other_start == other_end:
                conflict = start == other_start
            elif start == end:
                conflict = other_start < start < other_end
            elif other_start == other_end:
                conflict = start < other_start < end
            else:
                conflict = max(start, other_start) < min(end, other_end)
            if conflict:
                return "error"
    output, at, source = bytearray(), 0, value.encode()
    for start, end, replacement in sorted(edits, key=lambda edit: (edit[0], edit[1])):
        output.extend(source[at:start])
        output.extend(replacement.encode())
        at = end
    output.extend(source[at:])
    if edits and len(output) > limit:
        return "error"
    return output.hex()


def main():
    rng = random.Random(20260921)
    cases, expected = [], []
    counts = {"locations": 0, "spans": 0, "edits": 0, "renders": 0}
    for _ in range(120):
        value = "".join(rng.choice(TOKENS + ["\n", "\r", "\r\n"]) for _ in range(rng.randrange(0, 18)))
        for offset in range(-1, len(value.encode()) + 2):
            tab = rng.choice([1, 2, 4, 8])
            cases.append(f"location\t{encoded(value)}\t{offset}\t{tab}")
            expected.append(location(value, offset, tab))
            counts["locations"] += 1
        for _ in range(12):
            start = rng.randrange(-1, len(value.encode()) + 2)
            end = rng.randrange(-1, len(value.encode()) + 2)
            valid = start <= end and start in boundary_offsets(value) and end in boundary_offsets(value)
            cases.append(f"span\t{encoded(value)}\t{start}\t{end}")
            expected.append(value.encode()[start:end].hex() if valid else "error")
            counts["spans"] += 1
    for _ in range(1500):
        value = "".join(rng.choice(TOKENS) for _ in range(rng.randrange(0, 12)))
        candidates = boundary_offsets(value)
        if rng.randrange(4) == 0:
            candidates = list(range(-1, len(value.encode()) + 2))
        edits = []
        for _ in range(rng.randrange(0, 6)):
            a, b = rng.choice(candidates), rng.choice(candidates)
            if rng.randrange(8):
                a, b = sorted((a, b))
            edits.append((a, b, "".join(rng.choice(TOKENS) for _ in range(rng.randrange(0, 4)))))
        limit = rng.choice([0, 10, 20, 1000])
        cases.append("\t".join(["edit", encoded(value), str(limit)] + [f"{a}:{b}:{encoded(replacement)}" for a, b, replacement in edits]))
        expected.append(apply(value, edits, limit))
        counts["edits"] += 1
    render_start = len(cases)
    for _ in range(160):
        value = "".join(rng.choice(["a", "b", " "]) for _ in range(rng.randrange(0, 100)))
        start, end = sorted((rng.randrange(len(value) + 1), rng.randrange(len(value) + 1)))
        columns = rng.randrange(24, 100)
        for profile in ["plain", "ansi"]:
            cases.append(f"render\t{encoded(value)}\t{columns}\t{start}\t{end}\t{profile}")
            expected.append((columns, profile, value, start, end))
            counts["renders"] += 1
    result = subprocess.run([str(BINARY), "oracle"], input="\n".join(cases) + "\n", capture_output=True, text=True, check=True, timeout=120)
    output = result.stdout.splitlines()
    if len(output) != len(cases):
        raise AssertionError(f"expected {len(cases)} results, got {len(output)}: {result.stderr}")
    previous = None
    for index, (actual, wanted) in enumerate(zip(output, expected)):
        if index < render_start:
            if actual != wanted:
                raise AssertionError(f"case {index}: {cases[index]!r}; expected {wanted!r}, got {actual!r}")
        else:
            columns, profile, value, start, end = wanted
            decoded = bytes.fromhex(actual).decode()
            plain = re.sub(r"\x1b\[[0-9;]*m", "", decoded)
            if any(len(line) > columns for line in plain.splitlines()):
                raise AssertionError(f"render width {columns}: {plain!r}")
            if "^" not in plain or "label" not in plain:
                raise AssertionError(f"missing label: {plain!r}")
            if profile == "plain":
                previous = plain
            elif plain != previous or "\x1b[" not in decoded:
                raise AssertionError("ANSI profile changed rendered text or omitted styling")
            if columns >= len(value) + 20 and end - start <= 5 and start <= 5:
                expected_line = "  |   " + " " * start + "^" * max(1, end - start) + " label\n"
                if expected_line not in plain:
                    raise AssertionError(f"caret alignment: {plain!r}")
    print(f"diagnostics independent model checks passed: {len(cases)} cases; {counts}")


if __name__ == "__main__":
    main()
