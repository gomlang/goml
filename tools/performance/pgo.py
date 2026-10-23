#!/usr/bin/env python3
import argparse
import hashlib
import json
import os
import platform
from pathlib import Path
import shutil
import statistics
import subprocess
import time


def digest(path):
    return hashlib.sha256(Path(path).read_bytes()).hexdigest()


def run(command, directory, environment):
    started = time.perf_counter_ns()
    process = subprocess.run(command, cwd=directory, env=environment, text=True, capture_output=True)
    elapsed = time.perf_counter_ns() - started
    if process.returncode:
        raise RuntimeError(f"command failed: {command!r}\n{process.stdout}\n{process.stderr}")
    return process.stdout, elapsed


def main():
    root = Path(__file__).resolve().parents[2]
    parser = argparse.ArgumentParser(description="Reproduce native GoML PGO builds and balanced samples.")
    parser.add_argument("--goml", type=Path, default=root / "stage2/bin/goml")
    parser.add_argument("--compiler", type=Path, default=root / "stage2/bin/gomlc")
    parser.add_argument("--output", type=Path, default=root / "_artifact/performance-pgo")
    parser.add_argument("--samples", type=int, default=6)
    parser.add_argument("--iterations", type=int, default=2000000)
    parser.add_argument("--train-iterations", type=int, default=40000000)
    parser.add_argument("--wasm", action="store_true", help="also build JS/WASM and validate outputs in Node; does not measure a browser")
    args = parser.parse_args()
    if args.samples < 2 or args.iterations < 1 or args.train_iterations < 1:
        parser.error("samples must be at least two and iteration counts must be positive")
    output = args.output.resolve()
    output.mkdir(parents=True, exist_ok=True)
    for name in ("samples.json", "results.json", "wasm-validation.json"):
        (output / name).unlink(missing_ok=True)
    workspace = output / "source"
    shutil.copytree(root / "tools/performance/example", workspace, dirs_exist_ok=True, ignore=shutil.ignore_patterns("_artifact"))
    environment = dict(os.environ, CGO_ENABLED="0", GOTOOLCHAIN="local")
    profile = output / "training.pgo"
    builds = {}
    binaries = {}

    def build(name, pgo, target_environment):
        target = output / name
        command = [str(args.goml.resolve()), "build", "--compiler", str(args.compiler.resolve()), "--target-dir", str(target), "--pgo", str(pgo)]
        stdout, elapsed = run(command, workspace, target_environment)
        binary = target / "bin/example"
        record = Path(str(binary) + ".build.json")
        evidence = json.loads(record.read_text())
        if not evidence["success"] or evidence["artifact_sha256"] != digest(binary):
            raise RuntimeError(f"invalid build record: {record}")
        builds[name] = {"build_ns": elapsed, "artifact_bytes": binary.stat().st_size, "record": str(record), "record_sha256": digest(record), "artifact_sha256": digest(binary), "stdout": stdout}
        binaries[name] = binary

    build("plain", "off", environment)
    training, training_ns = run([str(binaries["plain"]), "-iterations", str(args.train_iterations), "-profile", str(profile)], workspace, environment)
    build("pgo", profile, environment)
    if args.wasm:
        wasm_environment = dict(environment, GOOS="js", GOARCH="wasm")
        build("wasm-plain", "off", wasm_environment)
        build("wasm-pgo", profile, wasm_environment)
    samples = []
    expected = None

    def sample(name, phase, round_index, position):
        nonlocal expected
        stdout, wall_ns = run([str(binaries[name]), "-iterations", str(args.iterations)], workspace, environment)
        value = json.loads(stdout)
        identity = value["checksum"], value["interpreter_steps"]
        if value["interpreter_steps"] != args.iterations:
            raise RuntimeError("work count changed")
        if expected is None:
            expected = identity
        elif identity != expected:
            raise RuntimeError(f"functional output changed: {name}: {identity} != {expected}")
        samples.append(dict(value, variant=name, phase=phase, round=round_index, position=position, wall_ns=wall_ns))
        (output / "samples.json").write_text(json.dumps(samples, indent=2) + "\n")

    for position, name in enumerate(("plain", "pgo")):
        sample(name, "warmup", -1, position)
    wasm_validation = None
    if args.wasm:
        node = shutil.which("node", path=environment.get("PATH"))
        go = shutil.which("go", path=environment.get("PATH"))
        if node is None or go is None:
            raise RuntimeError("WASM validation requires Node and the selected Go toolchain on PATH")
        node = str(Path(node).resolve())
        go = str(Path(go).resolve())
        go_environment = json.loads(run([go, "env", "-json", "GOROOT", "GOVERSION"], workspace, environment)[0])
        executor = Path(go_environment["GOROOT"]) / "lib/wasm/wasm_exec_node.js"
        runtime = executor.with_name("wasm_exec.js")
        if not executor.is_file() or not runtime.is_file():
            raise RuntimeError(f"selected Go toolchain lacks its Node WASM executor: {executor}")
        wasm_validation = {
            "purpose": "functional-output-validation-only",
            "node": {"path": node, "version": run([node, "--version"], workspace, environment)[0].strip()},
            "go": {"path": go, "version": go_environment["GOVERSION"], "goroot": go_environment["GOROOT"]},
            "executor": {"path": str(executor), "sha256": digest(executor), "runtime_path": str(runtime), "runtime_sha256": digest(runtime)},
            "native_reference": {"checksum": expected[0], "interpreter_steps": expected[1]},
            "runs": [],
            "limits": ["Node validates functional output, not browser performance.", "Work durations in program output are retained as raw output only and excluded from performance comparisons."],
        }
        wasm_environment = {"TMPDIR": environment.get("TMPDIR", "/tmp")}
        wasm_validation["environment"] = wasm_environment
        for name in ("wasm-plain", "wasm-pgo"):
            record = json.loads(Path(builds[name]["record"]).read_text())
            if record["go_environment"]["GOVERSION"] != go_environment["GOVERSION"]:
                raise RuntimeError(f"WASM executor Go version differs from build record: {name}")
            command = [node, str(executor), str(binaries[name]), "-iterations", str(args.iterations)]
            stdout, _ = run(command, workspace, wasm_environment)
            value = json.loads(stdout)
            identity = value["checksum"], value["interpreter_steps"]
            if identity != expected:
                raise RuntimeError(f"WASM functional output changed: {name}: {identity} != {expected}")
            wasm_validation["runs"].append({"variant": name, "command": command, "output": value})
        (output / "wasm-validation.json").write_text(json.dumps(wasm_validation, indent=2) + "\n")
    for round_index in range(args.samples):
        order = ("plain", "pgo") if round_index % 2 == 0 else ("pgo", "plain")
        for position, name in enumerate(order):
            sample(name, "measured", round_index, position)
    medians = {name: statistics.median(s["work_ns"] for s in samples if s["variant"] == name and s["phase"] == "measured") for name in ("plain", "pgo")}
    ratios = []
    for index in range(args.samples):
        paired = {sample["variant"]: sample["work_ns"] for sample in samples if sample["phase"] == "measured" and sample["round"] == index}
        ratios.append(paired["plain"] / paired["pgo"])
    report = {
        "version": 1, "kind": "goml-pgo-example", "machine": platform.uname()._asdict(),
        "cpu_model": next((line.split(":", 1)[1].strip() for line in Path("/proc/cpuinfo").read_text().splitlines() if line.startswith("model name")), platform.processor()) if Path("/proc/cpuinfo").is_file() else platform.processor(),
        "process_affinity": sorted(os.sched_getaffinity(0)) if hasattr(os, "sched_getaffinity") else None,
        "gomaxprocs": os.environ.get("GOMAXPROCS"),
        "go_environment": json.loads(run(["go", "env", "-json", "GOVERSION", "GOOS", "GOARCH", "GOFLAGS", "GOEXPERIMENT"], workspace, environment)[0]),
        "builds": builds, "profile_sha256": digest(profile), "profile_source": {"artifact": str(binaries["plain"]), "output": json.loads(training), "wall_ns": training_ns},
        "wasm_validation": wasm_validation,
        "samples": samples, "work_ns_medians": medians, "paired_speed_ratios": ratios,
        "median_throughput_change_percent": (medians["plain"] / medians["pgo"] - 1) * 100,
        "limits": ["Native process warmup does not model browser JIT warmup.", "Allocation counts cover the measured workload interval and runtime background work.", "No statistical significance or regression threshold is inferred.", "Optional WASM builds and Node runs verify records and functional output only; browser performance is unmeasured.", "Profile training uses one synthetic interpreter workload, not a representative application mix."],
    }
    (output / "results.json").write_text(json.dumps(report, indent=2) + "\n")
    print(json.dumps({"results": str(output / "results.json"), "median_throughput_change_percent": report["median_throughput_change_percent"], "paired_speed_ratios": ratios}, indent=2))


if __name__ == "__main__":
    main()
