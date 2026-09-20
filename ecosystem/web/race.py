import os
from pathlib import Path
import re
import subprocess
import sys


ROOT = Path(__file__).resolve().parent
GENERATED = ROOT / "_artifact/test/external/goml_generated.go"
BINARY = ROOT / "_artifact/web-race-tests"


def main():
    if not GENERATED.exists():
        raise RuntimeError("Run goml test in ecosystem/web before race.py")
    subprocess.run(["go", "test", "-race", "./adapter"], cwd=ROOT, check=True, timeout=120)
    subprocess.run(["go", "build", "-race", "-o", str(BINARY), str(GENERATED)], cwd=ROOT, check=True, timeout=180)
    tests = sorted(name for source in (ROOT / "tests").glob("*.gom") for name in re.findall(r"#\[test\]\s+fn\s+(\w+)\(", source.read_text()))
    environment = os.environ.copy()
    environment["GORACE"] = "halt_on_error=1 atexit_sleep_ms=0"
    for name in tests:
        subprocess.run([str(BINARY), f"ecosystem::web::tests::{name}"], cwd=ROOT, env=environment, check=True, timeout=30)
    driver = environment.get("GOML_VERIFY_DRIVER", str(ROOT.parent.parent / "stage2/bin/goml"))
    consumer_environment = environment.copy()
    consumer_environment["GOFLAGS"] = (consumer_environment.get("GOFLAGS", "") + " -race").strip()
    subprocess.run([driver, "test", "--timeout", "60s"], cwd=ROOT.parent / "consumers/web", env=consumer_environment, check=True, timeout=180)
    subprocess.run([driver, "build", "--target-dir", "_artifact/race"], cwd=ROOT.parent / "consumers/web", env=consumer_environment, check=True, timeout=180)
    consumer_binary = ROOT.parent / "consumers/web/_artifact/race/bin/web"
    environment["WEB_BINARY"] = str(consumer_binary)
    subprocess.run([sys.executable, str(ROOT / "interop.py")], cwd=ROOT, env=environment, check=True, timeout=30)
    print(f"web race detector: native adapter, {len(tests)} GoML tests, cross-library consumer tests and concurrent independent consumer passed")


if __name__ == "__main__":
    main()
