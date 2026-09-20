import os
import re
import subprocess
from pathlib import Path


ROOT = Path(__file__).resolve().parent


def main():
    generated = ROOT / "_artifact/test/external/goml_generated.go"
    binary = ROOT / "_artifact/datetime-race-tests"
    subprocess.run(["go", "build", "-race", "-o", str(binary), str(generated)], cwd=ROOT, check=True, timeout=120)
    names = sorted(name for source in (ROOT / "tests").glob("*.gom") for name in re.findall(r"#\[test\]\s+fn\s+(\w+)\(", source.read_text()))
    environment = dict(os.environ, GORACE="halt_on_error=1 atexit_sleep_ms=0")
    for name in names:
        subprocess.run([str(binary), f"ecosystem::datetime::tests::{name}"], cwd=ROOT, env=environment, check=True, timeout=30)
    print(f"datetime race detector: {len(names)} tests passed, including shared-zone concurrent conversions")


if __name__ == "__main__":
    main()
