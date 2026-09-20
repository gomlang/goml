from dataclasses import dataclass
from itertools import zip_longest
from pathlib import Path
import random
import string
import subprocess
import textwrap


ROOT = Path(__file__).resolve().parents[1]
BINARY = ROOT / "consumers/tui_markdown/_artifact/bin/tui_markdown"


@dataclass
class Case:
    name: str
    width: int
    source: str
    lines: list[str]
    links: int = 0
    mode: str = "normal"

    def request(self):
        return f"{self.width}\t{self.mode}\t{self.source.encode().hex()}"


def wrap(value, width):
    return textwrap.wrap(
        value, width, break_on_hyphens=False, replace_whitespace=False
    ) or [""]


def words(rng, width, count=None):
    return " ".join(
        "".join(rng.choices(string.ascii_lowercase, k=rng.randint(1, min(12, width))))
        for _ in range(count if count is not None else rng.randint(1, 24))
    )


def prefixed(value, width, first, rest=None):
    rest = first if rest is None else rest
    lines = wrap(value, width - max(len(first), len(rest)))
    return [(first if index == 0 else rest) + line for index, line in enumerate(lines)]


def table_lines(rows, width, alignments):
    columns = len(rows[0])
    base, extra = divmod(width - columns - 1, columns)
    widths = [base + (column < extra) for column in range(columns)]
    border = "┼" + "┼".join("─" * size for size in widths) + "┼"
    lines = [border]
    for index, row in enumerate(rows):
        wrapped = [wrap(cell, size) for cell, size in zip(row, widths)]
        for cells in zip_longest(*wrapped, fillvalue=""):
            padded = []
            for cell, size, alignment in zip(cells, widths, alignments):
                if alignment == "left":
                    padded.append(cell.ljust(size))
                elif alignment == "right":
                    padded.append(cell.rjust(size))
                else:
                    left = (size - len(cell)) // 2
                    padded.append(" " * left + cell + " " * (size - len(cell) - left))
            lines.append("│" + "│".join(padded) + "│")
        if index == 0:
            lines.append(border)
    return lines + [border]


def table_source(rows, alignments, outer=True):
    delimiters = {"left": ":---", "center": ":---:", "right": "---:"}
    values = [rows[0], [delimiters[value] for value in alignments], *rows[1:]]
    return "\n".join(
        ("| " if outer else "") + " | ".join(row) + (" |" if outer else "")
        for row in values
    )


def cases():
    rng = random.Random(20260921)
    result = []
    for index in range(240):
        width = rng.randint(1, 80)
        text = words(rng, width)
        source = text
        if index % 3 == 0:
            source = text.replace(" ", "\n")
        elif index % 3 == 1:
            source = " ".join(
                ("**" + word + "**") if at % 2 else ("*" + word + "*")
                for at, word in enumerate(text.split())
            )
        result.append(Case(f"paragraph {index}", width, source, wrap(text, width)))
    for index in range(60):
        width = rng.randint(1, 50)
        text = "".join(rng.choices(string.ascii_lowercase, k=rng.randint(1, 250)))
        result.append(Case(f"long word {index}", width, text, wrap(text, width)))
    for index in range(150):
        level = rng.randint(1, 6)
        width = rng.randint(level + 3, 80)
        text = words(rng, width - level - 1)
        marker = "#" * level + " "
        result.append(Case(
            f"heading {index}", width, marker + text,
            prefixed(text, width, marker, " " * len(marker)),
        ))
    for index in range(120):
        depth = rng.randint(1, 4)
        width = rng.randint(depth * 2 + 2, 80)
        text = words(rng, width - depth * 2)
        result.append(Case(
            f"quote {index}", width, "> " * depth + text,
            prefixed(text, width, "│ " * depth),
        ))
    for index in range(120):
        ordered = bool(index % 2)
        start = rng.choice([1, 7, 9, 98, 999])
        width = rng.randint(10, 80)
        source, lines = [], []
        for item in range(rng.randint(1, 6)):
            marker = f"{start + item}. " if ordered else "• "
            text = words(rng, width - len(marker))
            source.append((marker if ordered else "- ") + text)
            lines.extend(prefixed(text, width, marker, " " * len(marker)))
        result.append(Case(f"list {index}", width, "\n".join(source), lines))
    for index in range(120):
        width = rng.randint(3, 80)
        info = rng.choice(["", "gom", "text", "sample"])
        body = [
            "".join(rng.choices(string.ascii_letters + string.digits + "   ", k=rng.randint(0, 100)))
            for _ in range(rng.randint(1, 6))
        ]
        lines = [("┌ " + info)[:width]]
        for line in body:
            pieces = [line[start:start + width - 2] for start in range(0, len(line), width - 2)] or [""]
            lines.extend("│ " + piece for piece in pieces)
        lines.append("└")
        result.append(Case(
            f"code {index}", width, "~~~" + info + "\n" + "\n".join(body) + "\n~~~", lines,
        ))
    destinations = [
        ("https://example.org/guide", True),
        ("HTTP://example.org", True),
        ("mailto:hello@example.org", True),
        ("ftp://example.org/a", True),
        ("/guide", True),
        ("../guide", True),
        ("#part", True),
        ("javascript:alert", False),
        ("data:text/html,bad", False),
        ("file:/tmp/demo", False),
        ("vbscript:msgbox", False),
    ]
    for index in range(150):
        sources, rendered = [], []
        count = 0
        for item in range(rng.randint(1, 12)):
            url, allowed = rng.choice(destinations)
            label = "label" + str(item)
            image = rng.randrange(4) == 0
            sources.append(("!" if image else "") + f"[**{label}**]({url})")
            count += allowed
            value = label + (f"[{count}]" if allowed else "")
            rendered.append("[image: " + value + "]" if image else value)
        text = " ".join(rendered)
        result.append(Case(f"links {index}", 1024, " ".join(sources), [text], count))
    for index in range(120):
        columns = rng.randint(2, 5)
        width = rng.randint(columns * 6 + 1, 80)
        available = (width - columns - 1) // columns
        alignment = [rng.choice(["left", "center", "right"]) for _ in range(columns)]
        rows = [
            [words(rng, available, rng.randint(1, 4)) for _ in range(columns)]
            for _ in range(rng.randint(2, 5))
        ]
        result.append(Case(
            f"table {index}", width, table_source(rows, alignment, bool(index % 2)),
            table_lines(rows, width, alignment), mode="table",
        ))
    result.extend([
        Case("empty", 20, "", []),
        Case("blank document", 20, "\n\n \n", []),
        Case("reference definitions", 20, "[guide]: /guide\n", []),
        Case("soft break", 20, "alpha\nbeta", ["alpha beta"]),
        Case("hard break", 20, "alpha  \nbeta", ["alpha", "beta"]),
        Case("escaped hard break", 20, "alpha\\\nbeta", ["alpha", "beta"]),
        Case("paragraph spacing", 20, "alpha\n\nbeta\n\n\ngamma", ["alpha", "", "beta", "", "gamma"]),
        Case("loose list", 20, "- alpha\n\n- beta", ["• alpha", "", "• beta"]),
        Case("line ending CRLF", 20, "# title\r\n\r\nbody\r\n", ["# title", "", "body"]),
        Case("line ending CR", 20, "# title\r\rbody\r", ["# title", "", "body"]),
        Case("literal escapes", 40, r"\*plain\* and \[x](javascript:alert)", ["*plain* and [x](javascript:alert)"]),
        Case("inline code", 40, "a `x  y` and `` a ` b ``", ["a x  y and a ` b"]),
        Case("references", 80, "[first][g] and [g]\n\n[g]: /guide", ["first[1] and g[2]"], 2),
        Case("duplicate links", 80, "[a](/guide) [a](/guide)", ["a[1] a[2]"], 2),
        Case("autolinks", 100, "<https://example.org> <a@example.org>", ["https://example.org[1] a@example.org[2]"], 2),
        Case("entities", 40, "&amp; &#65; &#x42; &lt;", ["& A B <"]),
        Case("literal HTML", 40, "<b>hello</b>", ["<b>hello</b>"]),
        Case("control replacement", 40, "a\x1b[2Jb", ["a�[2Jb"]),
        Case("maximum width", 4096, "x", ["x"]),
    ])
    for width in [1, 2, 8, 41, 80]:
        result.append(Case(f"thematic break {width}", width, "---", ["─" * width]))
    escaped_tables = [
        (r"a | b\|" + "\n--- | ---\nc | d", [["a", "b|"], ["c", "d"]]),
        (r"a | b" + "\n--- | ---\n" + r"\*hi\* | \[x](/guide)", [["a", "b"], ["*hi*", "[x](/guide)"]]),
        (r"a | b" + "\n--- | ---\n" + r"\\*hi* | b\|c", [["a", "b"], ["\\hi", "b|c"]]),
    ]
    for index, (source, rows) in enumerate(escaped_tables):
        result.append(Case(
            f"table escapes {index}", 31, source, table_lines(rows, 31, ["left", "left"]), mode="table",
        ))
    rows = [["a", "b"], *[["x", "y"] for _ in range(256)]]
    result.append(Case(
        "table row boundary", 5, table_source(rows, ["left", "left"]),
        table_lines(rows, 5, ["left", "left"]), mode="table",
    ))
    rows = [["x"] * 32 for _ in range(64)]
    result.append(Case(
        "table cell boundary", 65, table_source(rows, ["left"] * 32),
        table_lines(rows, 65, ["left"] * 32), mode="table",
    ))
    rows = [["a", "b"], ["x" * 65536, "y"]]
    result.append(Case(
        "table byte boundary", 80, table_source(rows, ["left", "left"]),
        table_lines(rows, 80, ["left", "left"]), mode="table",
    ))
    return result


def verify_limits():
    rejected = [
        Case("zero width", 0, "x", []),
        Case("negative width", -1, "x", []),
        Case("oversize width", 4097, "x", []),
        Case("input byte limit", 80, "x" * 1048577, []),
        Case("output line limit", 1, "x" * 65537, []),
        Case("table row limit", 10, table_source([["a", "b"]] * 258, ["left", "left"]), [], mode="table"),
        Case("table cell limit", 65, table_source([["x"] * 32] * 65, ["left"] * 32), [], mode="table"),
        Case("table byte limit", 80, table_source([["a", "b"], ["x" * 65537, "y"]], ["left", "left"]), [], mode="table"),
    ]
    for case in rejected:
        result = subprocess.run(
            [str(BINARY), "--oracle"], input=case.request() + "\n",
            text=True, capture_output=True, timeout=30,
        )
        if result.returncode != 101 or result.stdout or result.stderr != "layout\n":
            raise AssertionError((case.name, result.returncode, result.stdout[:200], result.stderr[:200]))
    return len(rejected)


def main():
    checks = cases()
    result = subprocess.run(
        [str(BINARY), "--oracle"], input="\n".join(case.request() for case in checks) + "\n",
        text=True, capture_output=True, check=True, timeout=90,
    )
    actual = result.stdout.splitlines()
    if len(actual) != len(checks):
        raise AssertionError((len(actual), len(checks), result.stderr))
    for case, row in zip(checks, actual):
        encoded, count, links = row.split("\t")
        observed = (bytes.fromhex(encoded).decode(), int(count), int(links))
        expected = ("\n".join(case.lines), len(case.lines), case.links)
        if observed != expected:
            raise AssertionError((case.name, case.width, case.source[:400], observed, expected))
    rejected = verify_limits()
    print(
        f"tui_markdown reference oracle: {len(checks)} independent ASCII layout, block, link and table cases "
        f"plus {rejected} resource-limit rejection cases passed"
    )


if __name__ == "__main__":
    main()
