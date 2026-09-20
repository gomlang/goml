import hashlib
import importlib.util
import json
import os
from pathlib import Path
import re
import subprocess
import sys
import time


ROOT = Path(__file__).resolve().parent
REPOSITORY = ROOT.parents[2]
SPEC = importlib.util.spec_from_file_location("ecosystem_verifier", REPOSITORY / "ecosystem/verify.py")
ECOSYSTEM = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(ECOSYSTEM)


def main():
    environment = dict(os.environ, GOML_HOME=str(ECOSYSTEM.registry_snapshot()), GOML_BUILD_JOBS="2")
    goml = REPOSITORY / "stage2/bin/goml"
    logs = ROOT / "_artifact/verification"
    logs.mkdir(parents=True, exist_ok=True)
    records = []

    def run(name, command):
        start = time.monotonic()
        result = subprocess.run([str(part) for part in command], cwd=ROOT, env=environment, capture_output=True, text=True, timeout=300)
        (logs / f"{name}.log").write_text(result.stdout + result.stderr)
        records.append({"name": name, "seconds": time.monotonic() - start, "exit_code": result.returncode})
        if result.returncode:
            print(result.stdout + result.stderr)
            raise RuntimeError(f"{name} failed; see {logs}")
        print(f"explorer {name}: passed", flush=True)
        return result.stdout

    try:
        run("format", [goml, "fmt", "--check"])
        run("tests", [goml, "test", "--timeout", "60s"])
        run("build", [goml, "build"])
        generated = ROOT / "_artifact/build/pkg"
        fingerprint = lambda: {str(path.relative_to(generated)): hashlib.sha256(path.read_bytes()).hexdigest() for path in generated.rglob("*") if path.is_file() and not path.name.startswith(".")}
        before = fingerprint()
        run("cached-build", [goml, "build"])
        assert before == fingerprint(), "cached build changed artifacts"
        binary = ROOT / "_artifact/bin/explorer"
        first = run("snapshot", [binary, "--snapshot"])
        second = run("snapshot-repeat", [binary, "--demo"])
        assert first == second and "GoML terminal ecosystem" in first and "6/6" in first and "\x1b" not in first
        run("pty", [sys.executable, ROOT / "pty_test.py"])
        source = ROOT / "_artifact/test/external/goml_generated.go"
        race_binary = ROOT / "_artifact/race-tests"
        run("race-build", ["go", "build", "-race", "-o", race_binary, source])
        environment["GORACE"] = "halt_on_error=1 atexit_sleep_ms=0"
        for path in sorted((ROOT / "tests").glob("*.gom")):
            for name in re.findall(r"#\[test\]\s+fn\s+(\w+)\(", path.read_text()):
                run("race-" + name, [race_binary, "example::explorer::tests::" + name])
    finally:
        (logs / "report.json").write_text(json.dumps(records, indent=2) + "\n")


if __name__ == "__main__":
    main()
