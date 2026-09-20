import os
from pathlib import Path
import re
import subprocess


ROOT = Path(__file__).resolve().parent


def run_tests(module, coordinate):
    environment = os.environ.copy()
    environment["GORACE"] = "halt_on_error=1 atexit_sleep_ms=0"
    total = 0
    for kind, sources, package in (
        ("internal", module.glob("*_test.gom"), coordinate),
        ("external", (module / "tests").glob("*.gom"), coordinate + "::tests"),
    ):
        tests = sorted(name for source in sources for name in re.findall(r"#\[test\]\s+fn\s+(\w+)\(", source.read_text()))
        if not tests:
            continue
        generated = module / "_artifact/test" / kind / "goml_generated.go"
        if not generated.is_file():
            raise RuntimeError(f"Run goml test in {module} before race.py")
        binary = module / "_artifact" / ("race-" + kind + "-tests")
        subprocess.run(["go", "build", "-race", "-o", str(binary), str(generated)], cwd=module, check=True, timeout=180)
        for name in tests:
            subprocess.run([str(binary), f"{package}::{name}"], cwd=module, env=environment, check=True, timeout=60)
        total += len(tests)
    return total


def main():
    library = run_tests(ROOT, f"ecosystem::{ROOT.name}")
    consumer = run_tests(ROOT.parent / "consumers" / ROOT.name, f"consumer::{ROOT.name}")
    print(f"{ROOT.name} race detector: {library} library tests and {consumer} complete migrated consumer suite passed")


if __name__ == "__main__":
    main()
