import json
import math
from pathlib import Path
import subprocess

ROOT = Path(__file__).resolve().parent
BINARY = ROOT.parent / "consumers/progress/_artifact/bin/progress"


def main():
    result = subprocess.run([str(BINARY), "oracle"], stdin=subprocess.DEVNULL, capture_output=True, text=True, check=True, timeout=30)
    rows = [json.loads(line) for line in result.stdout.splitlines()]
    assert len(rows) == 400, len(rows)
    for index, row in enumerate(rows, 1):
        total = index * 37 % 1000 + 1
        position = index * 17 % (total + 1)
        elapsed = 10007
        assert (row["position"], row["total"], row["elapsed"]) == (position, total, elapsed), row
        if position:
            assert math.isclose(row["rate"], position * 1000 / elapsed, rel_tol=1e-12), row
            expected_eta = math.ceil((total - position) * elapsed / position)
            assert row["eta"] == expected_eta, (row, expected_eta)
        else:
            assert row["rate"] is None and row["eta"] is None, row
        filled = position * 10 // total
        marker = "|/-\\"[(elapsed // 100) % 4]
        expected = f"{marker} work [{'=' * filled}{'-' * (10 - filled)}] {position}/{total}"
        assert row["text"] == expected, (row, expected)
    print("progress reference oracle: 400 independent integer/fraction rate, ETA, spinner and bar cases passed")


if __name__ == "__main__":
    main()
