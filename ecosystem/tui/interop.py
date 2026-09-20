import copy
import random
import re
import subprocess
from pathlib import Path

BINARY = Path(__file__).resolve().parents[1] / "consumers/tui/_artifact/bin/tui"
GLYPHS = {"a": 1, "b": 1, "x": 1, "界": 2, "中": 2, "🙂": 2, "á": 1, "👩‍💻": 2, " ": 1}


def blank(color=-1, attributes=0):
    return [" ", 1, color, attributes]


class Screen:
    def __init__(self, width, height):
        self.width, self.height = width, height
        self.cells = [[blank() for _ in range(width)] for _ in range(height)]
        self.x = self.y = 0
        self.color = -1
        self.attributes = 0
        self.pending = ""

    def erase(self, x, y):
        cell = self.cells[y][x]
        if cell[1] == 0 and x > 0:
            old = self.cells[y][x - 1]
            self.cells[y][x - 1] = blank(old[2], old[3])
        elif cell[1] == 2 and x + 1 < self.width:
            old = self.cells[y][x + 1]
            self.cells[y][x + 1] = blank(old[2], old[3])
        self.cells[y][x] = blank(cell[2], cell[3])

    def put(self, x, y, glyph, color, attributes):
        width = GLYPHS[glyph]
        if not (0 <= x < self.width and 0 <= y < self.height) or x + width > self.width:
            return
        for column in range(x, x + width):
            self.erase(column, y)
        self.cells[y][x] = [glyph, width, color, attributes]
        if width == 2:
            self.cells[y][x + 1] = ["", 0, color, attributes]

    def feed(self, output):
        text = self.pending + output
        self.pending = ""
        i = 0
        while i < len(text):
            if text[i] == "\x1b":
                match = re.match(r"\x1b\[([0-9;?]*)([A-Za-z])", text[i:])
                if not match:
                    self.pending = text[i:]
                    break
                parameters, final = match.groups()
                i += len(match.group())
                if parameters.startswith("?"):
                    continue
                values = [int(value or 0) for value in parameters.split(";")]
                if final == "H":
                    self.y, self.x = values[0] - 1, values[1] - 1
                elif final == "m":
                    at = 0
                    while at < len(values):
                        value = values[at]
                        if value == 0:
                            self.color, self.attributes = -1, 0
                        elif value == 1:
                            self.attributes |= 1
                        elif value == 2:
                            self.attributes |= 2
                        elif value == 22:
                            self.attributes &= ~3
                        elif 30 <= value <= 37:
                            self.color = value - 30
                        elif 90 <= value <= 97:
                            self.color = value - 90 + 8
                        elif value == 39:
                            self.color = -1
                        elif value == 38 and values[at + 1] == 5:
                            self.color = values[at + 2]
                            at += 2
                        else:
                            raise AssertionError((values, value))
                        at += 1
                else:
                    raise AssertionError(final)
            else:
                glyph = next((value for value in sorted(GLYPHS, key=len, reverse=True) if text.startswith(value, i)), None)
                assert glyph is not None, repr(text[i:i + 20])
                self.put(self.x, self.y, glyph, self.color, self.attributes)
                self.x += GLYPHS[glyph]
                i += len(glyph)


def main():
    randomizer = random.Random(94317)
    commands = []
    expectations = []
    operations = 0
    for width, height in [(1, 1), (2, 2), (7, 3), (16, 5), (31, 8)]:
        commands.append(f"new\t{width}\t{height}")
        screen = Screen(width, height)
        commands.append("frame")
        expectations.append((width, height, copy.deepcopy(screen.cells)))
        for _ in range(500):
            x, y = randomizer.randrange(width), randomizer.randrange(height)
            color, attributes = randomizer.randrange(256), randomizer.randrange(4)
            if randomizer.random() < 0.78:
                glyph = randomizer.choice(list(GLYPHS))
                commands.append(f"put\t{x}\t{y}\t{glyph.encode().hex()}\t{color}\t{attributes}")
                screen.put(x, y, glyph, color, attributes)
            else:
                length = randomizer.randrange(width - x + 1)
                rows = randomizer.randrange(height - y + 1)
                commands.append(f"fill\t{x}\t{y}\t{length}\t{rows}\t{color}")
                for row in range(y, y + rows):
                    for column in range(x, x + length):
                        screen.erase(column, row)
                        screen.cells[row][column] = blank(color)
            commands.append("frame")
            expectations.append((width, height, copy.deepcopy(screen.cells)))
            operations += 1
    result = subprocess.run([str(BINARY), "--oracle"], input="\n".join(commands) + "\n", text=True, capture_output=True, timeout=60, check=True)
    rows = result.stdout.splitlines()
    assert len(rows) == len(expectations), (len(rows), len(expectations), result.stderr)
    actual = None
    for index, (line, (width, height, expected)) in enumerate(zip(rows, expectations)):
        if actual is None or (actual.width, actual.height) != (width, height):
            actual = Screen(width, height)
        encoded, changed = line.split("\t")
        actual.feed(bytes.fromhex(encoded).decode())
        assert actual.cells == expected, (index, line, actual.cells, expected)
        assert (changed == "0") == (encoded == "")
    print(f"tui reference model: {operations} seeded wide-cell/overlap/fill mutations, {len(rows)} ANSI replay frames and five dimensions passed")


if __name__ == "__main__":
    main()
