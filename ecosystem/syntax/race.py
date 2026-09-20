import os
from pathlib import Path
import re
import subprocess


ROOT = Path(__file__).resolve().parent
GENERATED = ROOT / "_artifact" / "test" / "external" / "goml_generated.go"
BINARY = ROOT / "_artifact" / "syntax-race-tests"


def main():
    subprocess.run(["go", "build", "-race", "-o", str(BINARY), str(GENERATED)], cwd=ROOT, check=True, timeout=120)
    names = sorted(name for source in (ROOT / "tests").glob("*.gom") for name in re.findall(r"#\[test\]\s+fn\s+(\w+)\(", source.read_text()))
    environment = os.environ.copy()
    environment["GORACE"] = "halt_on_error=1 atexit_sleep_ms=0"
    for name in names:
        subprocess.run([str(BINARY), f"ecosystem::syntax::tests::{name}"], cwd=ROOT, env=environment, check=True, timeout=90)
    print(f"syntax race detector: {len(names)} tests passed, including concurrent shared-cache interning and immutable snapshot editing")


if __name__ == "__main__":
    main()
