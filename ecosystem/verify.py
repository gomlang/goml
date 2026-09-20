import argparse
import hashlib
import json
import os
from pathlib import Path
import shutil
import subprocess
import sys
import time
import tomllib


ROOT = Path(__file__).resolve().parent
MODULES = (
    "parser", "proptest", "cli", "msgpack", "graph", "template", "redis",
    "pipeline", "ndarray", "sqlite", "lsp", "markdown", "diff",
)
IGNORED = {"_artifact", "_bootstrap", ".git", "__pycache__"}


def source_files(directory):
    return sorted(
        path for path in directory.rglob("*")
        if path.is_file() and not any(part in IGNORED for part in path.relative_to(directory).parts)
    )


def registry_snapshot():
    available = [name for name in MODULES if (ROOT / name / "goml.toml").is_file()]
    digest = hashlib.sha256()
    for name in available:
        for path in source_files(ROOT / name):
            digest.update(str(path.relative_to(ROOT)).encode())
            digest.update(b"\0")
            digest.update(path.read_bytes())
            digest.update(b"\0")
    home = ROOT / "_artifact" / "registry-snapshots" / digest.hexdigest()
    registry = home / "cache" / "registry"
    if not (registry / "index.toml").exists():
        registry.mkdir(parents=True, exist_ok=True)
        entries = []
        for name in available:
            source = ROOT / name
            manifest = tomllib.loads((source / "goml.toml").read_text())
            coordinate = manifest["module"]["path"]
            if coordinate != f"ecosystem::{name}":
                raise RuntimeError(f"unexpected module coordinate: {coordinate}")
            target = registry / "ecosystem" / name / "0.1.0"
            for path in source_files(source):
                destination = target / path.relative_to(source)
                destination.parent.mkdir(parents=True, exist_ok=True)
                destination.write_bytes(path.read_bytes())
            entries.append(f'[modules."{coordinate}"]\nlatest = "0.1.0"\nversions = ["0.1.0"]\n')
        (registry / "index.toml").write_text("\n".join(entries))
    return home


def run(command, cwd, environment, log, records, input_text=None):
    started = time.monotonic()
    result = subprocess.run(command, cwd=cwd, env=environment, capture_output=True, text=True, input=input_text)
    elapsed = time.monotonic() - started
    log.parent.mkdir(parents=True, exist_ok=True)
    log.write_text(result.stdout + result.stderr)
    records.append({"command": command, "cwd": str(cwd), "seconds": elapsed, "exit_code": result.returncode, "log": str(log)})
    if result.returncode:
        print(result.stdout + result.stderr, flush=True)
        raise RuntimeError(f"command failed: {command}; see {log}")
    print(f"  {log.stem}: ok ({elapsed:.2f}s)", flush=True)


def main():
    arguments = argparse.ArgumentParser()
    arguments.add_argument("modules", nargs="*", choices=MODULES)
    arguments.add_argument("--goml", type=Path, default=ROOT.parent / "stage2" / "bin" / "goml")
    args = arguments.parse_args()
    selected = args.modules or list(MODULES)
    for name in selected:
        for directory in (ROOT / name, ROOT / "consumers" / name):
            if not (directory / "goml.toml").is_file():
                raise RuntimeError(f"required module is not implemented: {directory}")
    goml = str(args.goml.resolve())
    environment = os.environ.copy()
    environment["GOML_HOME"] = str(registry_snapshot())
    records = []
    report = ROOT / "_artifact" / "verification" / "report.json"
    report.parent.mkdir(parents=True, exist_ok=True)
    try:
        for name in selected:
            print(f"Verifying {name}", flush=True)
            library = ROOT / name
            consumer = ROOT / "consumers" / name
            logs = report.parent / name
            for directory, prefix in ((library, "library"), (consumer, "consumer")):
                run([goml, "fmt", "--check"], directory, environment, logs / f"{prefix}-format.log", records)
                run([goml, "test", "--timeout", "60s"], directory, environment, logs / f"{prefix}-tests.log", records)
            run([goml, "build"], consumer, environment, logs / "consumer-build.log", records)
            generated = consumer / "_artifact" / "build" / "pkg"
            before = {str(path.relative_to(generated)): hashlib.sha256(path.read_bytes()).hexdigest() for path in generated.rglob("*") if path.is_file() and not path.name.startswith(".")}
            run([goml, "build"], consumer, environment, logs / "consumer-cached-build.log", records)
            after = {str(path.relative_to(generated)): hashlib.sha256(path.read_bytes()).hexdigest() for path in generated.rglob("*") if path.is_file() and not path.name.startswith(".")}
            if before != after:
                raise RuntimeError(f"cached build changed public artifacts for {name}")
            binary = consumer / "_artifact" / "bin" / name
            run([str(binary)], consumer, environment, logs / "consumer-run.log", records)
            if name in ("diff", "lsp", "markdown", "msgpack", "template", "redis", "pipeline", "ndarray", "sqlite"):
                run([sys.executable, str(library / "interop.py")], ROOT.parent, environment, logs / "interoperability.log", records)
            if name == "ndarray":
                run([sys.executable, str(library / "simd_check.py")], ROOT.parent, environment, logs / "simd.log", records)
            if name in ("redis", "pipeline", "sqlite"):
                run([sys.executable, str(library / "race.py")], ROOT.parent, environment, logs / "race-detector.log", records)
            if name == "cli":
                run([sys.executable, str(library / "diagnostics.py")], ROOT.parent, environment, logs / "derive-diagnostics.log", records)
        if set(selected) == set(MODULES):
            for name in ("unit_identity", "erased_generic", "specialized_static", "ffi_error_alias"):
                print(f"Verifying compiler regression {name}", flush=True)
                directory = ROOT / "repros" / name
                logs = report.parent / "repros" / name
                for command in ("check", "build", "run"):
                    run([goml, command], directory, environment, logs / f"{command}.log", records, input_text="")
    finally:
        report.write_text(json.dumps({"modules": selected, "registry_home": environment["GOML_HOME"], "commands": records}, indent=2) + "\n")
    print(f"Verification report: {report}", flush=True)


if __name__ == "__main__":
    main()
