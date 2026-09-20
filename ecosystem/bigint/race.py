import os
from pathlib import Path
import subprocess


ROOT = Path(__file__).resolve().parent


def main():
    generated = ROOT / "_artifact/test/external/goml_generated.go"
    binary = ROOT / "_artifact/bigint-race-tests"
    if not generated.exists():
        raise RuntimeError("Run goml test in ecosystem/bigint before race.py")
    subprocess.run(["go", "build", "-race", "-o", str(binary), str(generated)], cwd=ROOT, check=True, timeout=120)
    environment = os.environ.copy()
    environment["GORACE"] = "halt_on_error=1 atexit_sleep_ms=0"
    subprocess.run([str(binary), "ecosystem::bigint::tests::concurrent_arithmetic_preserves_shared_values"], cwd=ROOT, env=environment, check=True, timeout=30)
    print("Bigint race detector: concurrent shared arithmetic and detached byte mutation passed")


if __name__ == "__main__":
    main()
