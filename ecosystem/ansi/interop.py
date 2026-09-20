from pathlib import Path
import random
import re
import subprocess


ROOT = Path(__file__).resolve().parents[1]
BINARY = ROOT / "consumers/ansi/_artifact/bin/ansi"
ESC = "\x1b"
BASE = [(0, 0, 0), (128, 0, 0), (0, 128, 0), (128, 128, 0),
        (0, 0, 128), (128, 0, 128), (0, 128, 128), (192, 192, 192),
        (128, 128, 128), (255, 0, 0), (0, 255, 0), (255, 255, 0),
        (0, 0, 255), (255, 0, 255), (0, 255, 255), (255, 255, 255)]
LEVELS = [0, 95, 135, 175, 215, 255]
PALETTE = BASE + [(r, g, b) for r in LEVELS for g in LEVELS for b in LEVELS]
PALETTE += [(n, n, n) for n in range(8, 239, 10)]
ATTRIBUTES = {1: 1, 2: 2, 3: 4, 4: 8, 5: 16, 7: 32, 8: 64, 9: 128, 53: 256}
RESETS = {22: 3, 23: 4, 24: 8, 25: 16, 27: 32, 28: 64, 29: 128, 55: 256}


def encode(text):
    return text.encode().hex()


def advance(parameters, state):
    values = [int(value or 0) for value in parameters.split(";")]
    bits, foreground, background = state
    at = 0
    while at < len(values):
        code = values[at]
        if code == 0:
            bits, foreground, background = 0, "default", "default"
        elif code in ATTRIBUTES:
            bits |= ATTRIBUTES[code]
        elif code in RESETS:
            bits &= ~RESETS[code]
        elif 30 <= code <= 37 or 90 <= code <= 97:
            foreground = f"i:{code - 30 if code < 90 else code - 82}"
        elif 40 <= code <= 47 or 100 <= code <= 107:
            background = f"i:{code - 40 if code < 100 else code - 92}"
        elif code in (39, 49):
            if code == 39:
                foreground = "default"
            else:
                background = "default"
        elif code in (38, 48):
            if values[at + 1] == 5:
                color = f"i:{values[at + 2]}"
                at += 2
            else:
                color = "r:" + ":".join(map(str, values[at + 2:at + 5]))
                at += 4
            if code == 38:
                foreground = color
            else:
                background = color
        at += 1
    return bits, foreground, background


def screen(text):
    state = (0, "default", "default")
    printed = []
    at = 0
    for match in re.finditer(r"\x1b\[([0-9;]*)m", text):
        for character in text[at:match.start()]:
            printed.append((character, state))
        state = advance(match[1], state)
        at = match.end()
    for character in text[at:]:
        printed.append((character, state))
    return printed, state


def main():
    rng = random.Random(20260920)
    rows, checks = [], []
    choices = list(ATTRIBUTES) + list(RESETS) + [0, 39, 49] + list(range(30, 38)) + list(range(40, 48)) + list(range(90, 98)) + list(range(100, 108))
    for _ in range(1200):
        parameters = []
        for _ in range(rng.randrange(1, 30)):
            if rng.randrange(4):
                parameters.append(str(rng.choice(choices)))
            else:
                prefix = rng.choice([38, 48])
                if rng.randrange(2):
                    parameters.extend(map(str, [prefix, 5, rng.randrange(256)]))
                else:
                    parameters.extend(map(str, [prefix, 2, *[rng.randrange(256) for _ in range(3)]]))
        value = ";".join(parameters)
        rows.append("sgr\t" + encode(value))
        checks.append(("equal", "\t".join(map(str, advance(value, (0, "default", "default"))))))
    for _ in range(400):
        parts = [rng.choice(["hello", "中文", "é", "👩‍💻", "\n", "\t"]) for _ in range(rng.randrange(1, 12))]
        controls = [f"{ESC}[2J", f"{ESC}[31m", f"{ESC}[0m", f"{ESC}]0;title\x07", f"{ESC}]8;;https://example.org{ESC}\\", f"{ESC}]8;;{ESC}\\", f"{ESC}Popaque{ESC}\\", "\x07"]
        source = "".join(rng.choice(controls) + part + rng.choice(controls) for part in parts)
        rows.append("strip\t" + encode(source))
        checks.append(("equal", encode("".join(parts))))
    for profile in ["plain", "16", "256", "true"]:
        for _ in range(300):
            rgb = tuple(rng.randrange(256) for _ in range(3))
            bits = rng.randrange(512)
            payload = "hello世界"
            rows.append("\t".join(map(str, ["render", profile, *rgb, bits, encode(payload)])))
            if profile == "plain":
                expected = (0, "default", "default")
            elif profile in ("16", "256"):
                palette = PALETTE[:int(profile)]
                index = min(range(len(palette)), key=lambda i: sum((a - b) ** 2 for a, b in zip(rgb, palette[i])))
                expected = (bits, f"i:{index}", "default")
            else:
                expected = (bits, "r:" + ":".join(map(str, rgb)), "default")
            checks.append(("render", (payload, expected)))
    result = subprocess.run([str(BINARY), "--oracle"], input="\n".join(rows) + "\n", text=True, capture_output=True, check=True, timeout=90)
    actual = result.stdout.splitlines()
    if len(actual) != len(checks):
        raise AssertionError((len(actual), len(checks), result.stderr))
    for index, (row, (kind, expected)) in enumerate(zip(actual, checks)):
        if kind == "equal":
            if row != expected:
                raise AssertionError((index, rows[index], row, expected))
        else:
            payload, state = expected
            cells, final = screen(bytes.fromhex(row).decode())
            if cells != [(character, state) for character in payload] or final != (0, "default", "default"):
                raise AssertionError((index, cells, final, expected))
    print(f"ansi interoperability: {len(checks)} independent SGR, escape stripping, palette and rendered-state checks passed")


if __name__ == "__main__":
    main()
