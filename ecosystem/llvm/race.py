import os
from pathlib import Path
import re
import subprocess


ROOT = Path(__file__).resolve().parent
GENERATED = ROOT / "_artifact" / "test" / "external" / "goml_generated.go"
BINARY = ROOT / "_artifact" / "llvm-race-tests"


def main():
    if not GENERATED.is_file():
        raise RuntimeError("Run goml test in ecosystem/llvm before race.py")
    environment = os.environ.copy()
    environment["GORACE"] = "halt_on_error=1 atexit_sleep_ms=0"
    subprocess.run(["go", "test", "-race", "-count=1", "./adapter"], cwd=ROOT, env=environment, check=True, timeout=180)
    subprocess.run(["go", "build", "-race", "-o", str(BINARY), str(GENERATED)], cwd=ROOT, env=environment, check=True, timeout=180)
    tests = sorted(name for source in (ROOT / "tests").glob("*.gom") for name in re.findall(r"#\[test\]\s+fn\s+(\w+)\(", source.read_text()))
    for name in tests:
        subprocess.run([str(BINARY), f"ecosystem::llvm::tests::{name}"], cwd=ROOT, env=environment, check=True, timeout=60)
    print(f"LLVM race detector: 4 native concurrency/lifetime tests and {len(tests)} GoML tests passed")


if __name__ == "__main__":
    main()
