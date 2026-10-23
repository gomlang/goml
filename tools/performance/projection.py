#!/usr/bin/env python3
import argparse
import hashlib
import json
import os
from pathlib import Path
import platform
import re
import shutil
import statistics
import subprocess
import time


def digest(path):
    return hashlib.sha256(Path(path).read_bytes()).hexdigest()


def command(args, cwd, env):
    started = time.perf_counter_ns()
    result = subprocess.run(args, cwd=cwd, env=env, text=True, capture_output=True)
    if result.returncode:
        raise RuntimeError(f"command failed: {args!r}\n{result.stdout}\n{result.stderr}")
    return result, time.perf_counter_ns() - started


def main():
    root = Path(__file__).resolve().parents[2]
    parser = argparse.ArgumentParser(description="Compare generated immutable field and array reads.")
    parser.add_argument("--baseline", type=Path, default=root / "stage0/bin/gomlc")
    parser.add_argument("--candidate", type=Path, default=root / "stage2/bin/gomlc")
    parser.add_argument("--goml", type=Path, default=root / "stage2/bin/goml")
    parser.add_argument("--output", type=Path, default=root / "_artifact/performance-projection")
    parser.add_argument("--samples", type=int, default=6)
    args = parser.parse_args()
    if args.samples < 2:
        parser.error("at least two samples are required")
    output = args.output.resolve()
    output.mkdir(parents=True, exist_ok=True)
    for name in ("samples.json", "results.json", "performance.json"):
        (output / name).unlink(missing_ok=True)
    source = output / "source"
    shutil.copytree(root / "tools/performance/projection", source, dirs_exist_ok=True, ignore=shutil.ignore_patterns("_artifact"))
    env = dict(os.environ, GOTOOLCHAIN="local", GOFLAGS="-pgo=off", CGO_ENABLED="0")
    cpu = min(os.sched_getaffinity(0)) if hasattr(os, "sched_getaffinity") and shutil.which("taskset") else None
    harness = '''package main

import "testing"

var sink int

func BenchmarkProjection(b *testing.B) {
    var value Outer
    for i := range value.inner.values { value.inner.values[i] = i + 1 }
    if project(&value, 1) != 528 { b.Fatal("incorrect checksum") }
    b.ReportAllocs()
    b.ResetTimer()
    for i := 0; i < b.N; i++ { sink = project(&value, 1) }
}
'''
    builds = {}
    for name, compiler in (("baseline", args.baseline.resolve()), ("candidate", args.candidate.resolve())):
        directory = output / name
        directory.mkdir(exist_ok=True)
        result, elapsed = command([str(compiler), "run-single", str(source / "main.goml"), "--dump-go"], root, env)
        (directory / "dump.txt").write_text(result.stdout)
        (directory / "diagnostics.txt").write_text(result.stderr)
        generated = result.stdout.split("== Go ==\n", 1)[1]
        generated, actual = generated.rsplit("\n3200\n", 1)
        if actual.strip():
            raise RuntimeError(f"unexpected execution output: {actual}")
        (directory / "generated.go").write_text(generated + "\n")
        (directory / "projection_test.go").write_text(harness)
        (directory / "go.mod").write_text("module example.com/projection\n\ngo 1.26\n")
        build, native_ns = command(["go", "test", "-c", "-o", "benchmark", "-gcflags=-m=2"], directory, env)
        (directory / "escape.txt").write_text(build.stderr)
        assembly, _ = command(["go", "tool", "objdump", "-s", r"^example\.com/projection\.(read|project|BenchmarkProjection)$", "benchmark"], directory, env)
        (directory / "assembly.txt").write_text(assembly.stdout)
        builds[name] = {
            "compiler": str(compiler), "compiler_sha256": digest(compiler),
            "generated_sha256": digest(directory / "generated.go"), "generated_bytes": (directory / "generated.go").stat().st_size,
            "compile_and_run_ns": elapsed, "native_test_build_ns": native_ns,
            "benchmark_bytes": (directory / "benchmark").stat().st_size,
            "benchmark_sha256": digest(directory / "benchmark"),
        }
    explanation, _ = command([str(args.goml.resolve()), "build", "--compiler", str(args.candidate.resolve()), "--target-dir", str(output / "explain"), "--explain-performance", "--format", "json", "--pgo", "off"], source, env)
    json.loads(explanation.stdout)
    (output / "performance.json").write_text(explanation.stdout)
    samples = []

    def sample(name, phase, index, position):
        invocation = [str(output / name / "benchmark"), "-test.run=^$", "-test.bench=^BenchmarkProjection$", "-test.benchtime=250ms", "-test.benchmem"]
        if cpu is not None:
            invocation = ["taskset", "-c", str(cpu)] + invocation
        result, elapsed = command(invocation, root, env)
        match = re.search(r"BenchmarkProjection\S*\s+(\d+)\s+([\d.]+) ns/op\s+(\d+) B/op\s+(\d+) allocs/op", result.stdout)
        if not match:
            raise RuntimeError(result.stdout)
        samples.append({"variant": name, "phase": phase, "round": index, "position": position, "iterations": int(match[1]), "ns_per_op": float(match[2]), "bytes_per_op": int(match[3]), "allocs_per_op": int(match[4]), "wall_ns": elapsed, "stdout": result.stdout})
        (output / "samples.json").write_text(json.dumps(samples, indent=2) + "\n")

    for position, name in enumerate(("baseline", "candidate")):
        sample(name, "warmup", -1, position)
    for index in range(args.samples):
        order = ("baseline", "candidate") if index % 2 == 0 else ("candidate", "baseline")
        for position, name in enumerate(order):
            sample(name, "measured", index, position)
    medians = {name: statistics.median(s["ns_per_op"] for s in samples if s["variant"] == name and s["phase"] == "measured") for name in ("baseline", "candidate")}
    result = {
        "version": 1, "kind": "goml-projection-comparison", "machine": platform.uname()._asdict(),
        "cpu_model": next((line.split(":", 1)[1].strip() for line in Path("/proc/cpuinfo").read_text().splitlines() if line.startswith("model name")), platform.processor()) if Path("/proc/cpuinfo").is_file() else platform.processor(),
        "process_affinity": sorted(os.sched_getaffinity(0)) if hasattr(os, "sched_getaffinity") else None,
        "gomaxprocs": os.environ.get("GOMAXPROCS"), "cpu_affinity": cpu,
        "go_environment": json.loads(command(["go", "env", "-json", "GOVERSION", "GOOS", "GOARCH", "GOFLAGS", "GOEXPERIMENT"], root, env)[0].stdout),
        "source_sha256": digest(source / "main.goml"), "builds": builds, "samples": samples, "medians_ns": medians,
        "median_throughput_change_percent": (medians["baseline"] / medians["candidate"] - 1) * 100,
        "scope": "One project call reading 32 integers; checksum 528; no PGO; all builds finish before samples.",
        "limitations": ["Single-machine native microbenchmark, not browser or simulator throughput.", "Build timings are single observations and affected by existing caches.", "The explanation describes a module build of the same source; benchmark uses run-single symbol names.", "No statistical significance or hard regression threshold is inferred."],
    }
    (output / "results.json").write_text(json.dumps(result, indent=2) + "\n")
    print(json.dumps({"results": str(output / "results.json"), "medians_ns": medians}, indent=2))


if __name__ == "__main__":
    main()
