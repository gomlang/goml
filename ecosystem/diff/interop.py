import os
from pathlib import Path
import random
import subprocess
import tempfile


ROOT = Path(__file__).resolve().parents[1]
BINARY = ROOT / "consumers" / "diff" / "_artifact" / "bin" / "diff"


def main():
    if not BINARY.is_file():
        raise RuntimeError("build the diff consumer with ecosystem/verify.py first")
    rng = random.Random(20260920)
    cases = [
        ("", ""), ("", "hello"), ("hello", ""), ("hello", "hello\n"),
        ("a\r\nb\r\n", "a\r\nc\r\n"), ("你好\n世界", "你好\n朋友"),
        ("\n", ""), ("a\nb\nc\nd\ne\nf\ng\nh\ni\n", "A\nb\nc\nd\ne\nf\ng\nh\nI\n"),
    ]
    alphabet = ["a\n", "b\n", "你好\n", "😀\n", "\n", "tab\there\r\n"]
    for _ in range(40):
        old = "".join(rng.choice(alphabet) for _ in range(rng.randrange(30)))
        new = "".join(rng.choice(alphabet) for _ in range(rng.randrange(30)))
        if rng.randrange(2):
            old += "last"
        if rng.randrange(2):
            new += "末尾"
        cases.append((old, new))
    artifacts = ROOT / "_artifact" / "interop"
    artifacts.mkdir(parents=True, exist_ok=True)
    environment = os.environ.copy()
    environment["LC_ALL"] = "C"
    with tempfile.TemporaryDirectory(prefix="diff-", dir=artifacts) as temporary:
        directory = Path(temporary)
        old_path, new_path = directory / "old", directory / "new"
        generated, reference = directory / "generated.patch", directory / "reference.patch"
        applied, native = directory / "applied", directory / "native"
        for index, (old, new) in enumerate(cases):
            old_path.write_bytes(old.encode())
            new_path.write_bytes(new.encode())
            for algorithm in ("produce", "produce-linear"):
                subprocess.run([str(BINARY), algorithm, str(old_path), str(new_path), str(generated)], check=True)
                if old != new:
                    subprocess.run(["patch", "--batch", "--silent", "--output", str(native), str(old_path), str(generated)], env=environment, check=True)
                    if native.read_bytes() != new.encode():
                        raise AssertionError(f"GNU patch mismatch for {algorithm} case {index}")
            result = subprocess.run(["diff", "-u", "--label", "old", "--label", "new", str(old_path), str(new_path)], env=environment, capture_output=True)
            if result.returncode not in (0, 1):
                raise RuntimeError(result.stderr.decode())
            reference.write_bytes(result.stdout)
            subprocess.run([str(BINARY), "apply", str(old_path), str(reference), str(applied)], check=True)
            if applied.read_bytes() != new.encode():
                raise AssertionError(f"GoML apply mismatch for GNU diff case {index}")
    print(f"diff interoperability: {len(cases)} cases passed with both GoML algorithms and GNU diff/patch")


if __name__ == "__main__":
    main()
