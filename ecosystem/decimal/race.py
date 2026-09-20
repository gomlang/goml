import os
from pathlib import Path
import subprocess


ROOT = Path(__file__).resolve().parent


def main():
    generated = ROOT / "_artifact/test/external/goml_generated.go"
    binary = ROOT / "_artifact/decimal-race-tests"
    if not generated.exists():
        raise RuntimeError("Run goml test in ecosystem/decimal before race.py")
    subprocess.run(["go", "build", "-race", "-o", str(binary), str(generated)], cwd=ROOT, check=True, timeout=120)
    environment = os.environ.copy()
    environment["GORACE"] = "halt_on_error=1 atexit_sleep_ms=0"
    subprocess.run([str(binary), "ecosystem::decimal::tests::concurrent_arithmetic_preserves_shared_decimal_values"], cwd=ROOT, env=environment, check=True, timeout=30)
    print("Decimal race detector: shared exact arithmetic, contextual division, and detached coefficient bytes passed")


if __name__ == "__main__":
    main()
