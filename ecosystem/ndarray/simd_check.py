import importlib.util
import json
import os
from pathlib import Path
import re
import subprocess
import sys


ROOT = Path(__file__).resolve().parents[1]
CONSUMER = ROOT / "consumers/ndarray"
GOML = ROOT.parent / "stage2/bin/goml"


def main():
    spec = importlib.util.spec_from_file_location("ecosystem_verify", ROOT / "verify.py")
    verifier = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(verifier)
    environment = os.environ.copy()
    environment["GOML_HOME"] = str(verifier.registry_snapshot())
    environment.pop("GOML_SIMD", None)
    source = CONSUMER / "_artifact/build/pkg/consumer/ndarray/goml_generated.go"
    marker = "const _goml_simd_native = "
    line = next(line for line in source.read_text().splitlines() if line.startswith(marker))
    metadata = json.loads(json.loads(line[len(marker):]))
    kernels = metadata["kernels"]
    assembly = "\n".join(kernel["assembly"] for kernel in kernels)
    for instruction in ("ADDPS", "SUBPS", "MULPS", "DIVPS", "VADDPD", "VSUBPD", "VMULPD", "VDIVPD"):
        if instruction not in assembly:
            raise AssertionError(f"missing native SIMD instruction {instruction}")
    symbols = [(re.search(r"TEXT ·([^ (]+)\(SB\)", kernel["assembly"])[1] + ".abi0", kernel["avx2"] or kernel["fma"]) for kernel in kernels]
    report = []
    for mode in ("native", "sse2", "scalar"):
        target = CONSUMER / "_artifact" if mode == "native" else CONSUMER / "_artifact" / f"simd-{mode}"
        if mode != "native":
            environment["GOML_SIMD"] = mode
            result = subprocess.run([str(GOML), "build", "--target-dir", str(target)], cwd=CONSUMER, env=environment, capture_output=True, text=True, timeout=180)
            (CONSUMER / "_artifact" / f"{mode}-build.log").write_text(result.stdout + result.stderr)
            if result.returncode:
                raise RuntimeError(f"{mode} build failed: {result.stdout}{result.stderr}")
        binary = target / "bin/ndarray"
        symbol_table = subprocess.run(["go", "tool", "nm", str(binary)], capture_output=True, text=True, check=True).stdout
        linked = [symbol for symbol, _ in symbols if symbol in symbol_table]
        expected = [symbol for symbol, needs_avx in symbols if mode == "native" or mode == "sse2" and not needs_avx]
        if set(linked) != set(expected):
            raise AssertionError(f"{mode}: unexpected linked assembly kernels: {linked}, expected {expected}")
        subprocess.run([str(binary)], check=True, timeout=30)
        if mode != "native":
            subprocess.run([sys.executable, str(ROOT / "ndarray/interop.py"), "--consumer", str(binary)], check=True, timeout=120)
        report.append({"mode": mode, "linked_native_kernels": len(linked), "symbols": linked})
        print(f"{mode}: {len(linked)} linked assembly kernels, consumer checks passed", flush=True)
    destination = CONSUMER / "_artifact/simd-verification.json"
    destination.write_text(json.dumps(report, indent=2) + "\n")
    print(f"SIMD validation: {destination}")


if __name__ == "__main__":
    main()
