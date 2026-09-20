import json
import subprocess
from pathlib import Path


ROOT = Path(__file__).resolve().parent
BINARY = ROOT.parent / "consumers" / "llvm" / "_artifact" / "bin" / "llvm"
ARTIFACT = ROOT / "_artifact" / "interop"
LLVM_BIN = Path("/usr/lib/llvm-18/bin")
HARNESS = r'''
#include <inttypes.h>
#include <stdint.h>
#include <stdio.h>
extern int64_t goml_mix(int64_t);
extern int64_t goml_absolute_mix(int64_t);
extern int64_t goml_sum(int64_t);
extern double goml_scale(double);
extern int64_t goml_struct(int64_t);
extern double goml_convert(int64_t);
int main(void) {
    for (int64_t x = -200; x <= 200; ++x) {
        int64_t n = x < 0 ? -x : x;
        printf("%" PRId64 " %" PRId64 " %" PRId64 " %.17g %" PRId64 " %.17g\n",
            goml_mix(x), goml_absolute_mix(x), goml_sum(n), goml_scale((double)x), goml_struct(x), goml_convert(x));
    }
    return 0;
}
'''


def run(command):
    return subprocess.run([str(value) for value in command], check=True, capture_output=True, text=True, timeout=90)


def main():
    if not BINARY.is_file():
        raise RuntimeError("Build ecosystem/consumers/llvm before running interop.py")
    ARTIFACT.mkdir(parents=True, exist_ok=True)
    prefix = ARTIFACT / "program"
    run([BINARY, prefix])
    harness = ARTIFACT / "harness.c"
    harness.write_text(HARNESS)
    run([LLVM_BIN / "opt", "-passes=verify", "-disable-output", prefix.with_suffix(".bc")])
    run([LLVM_BIN / "llvm-dis", prefix.with_suffix(".bc"), "-o", ARTIFACT / "decoded.ll"])
    run([LLVM_BIN / "llvm-as", prefix.with_suffix(".ll"), "-o", ARTIFACT / "assembled.bc"])
    run([LLVM_BIN / "llc", "-filetype=obj", "-relocation-model=pic", "-O=2", ARTIFACT / "assembled.bc", "-o", ARTIFACT / "cli.o"])
    count = 0
    for variant, source in (("unoptimized", ARTIFACT / "program-unoptimized.o"), ("optimized", prefix.with_suffix(".o")), ("assembly", prefix.with_suffix(".s")), ("llvm_cli", ARTIFACT / "cli.o")):
        executable = ARTIFACT / variant
        run(["cc", "-O2", harness, source, "-o", executable])
        lines = run([executable]).stdout.splitlines()
        if len(lines) != 401:
            raise AssertionError((variant, len(lines)))
        for x, line in zip(range(-200, 201), lines):
            raw = line.split()
            actual = [int(raw[0]), int(raw[1]), int(raw[2]), float(raw[3]), int(raw[4]), float(raw[5])]
            n = abs(x)
            expected = [3 * x + 7, 3 * n + 7, n * (n - 1) // 2, x * 1.5 + 0.25, x, float(x)]
            if actual != expected:
                raise AssertionError((variant, x, actual, expected))
            count += len(expected)
    print(json.dumps({"llvm": run([LLVM_BIN / "llvm-config", "--version"]).stdout.strip(), "variants": 4, "function_results": count, "bitcode_and_ir_roundtrip": True}))


if __name__ == "__main__":
    main()
