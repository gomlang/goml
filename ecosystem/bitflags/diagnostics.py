import os
from pathlib import Path
import subprocess
import sys
import tempfile


ROOT = Path(__file__).resolve().parents[1]
sys.path.insert(0, str(ROOT))
import verify


def main():
    goml = str(ROOT.parent / "stage2/bin/goml")
    environment = dict(os.environ, GOML_HOME=str(verify.registry_snapshot()))
    cases = [
        ('#[flags(A="1")] enum Bad { Item }', "exactly one integer field"),
        ('#[flags(A="1")] struct Bad { a: u8, b: u8 }', "exactly one integer field"),
        ('#[flags(A="1")] struct Bad[T] { value: T }', "non-generic struct"),
        ('#[flags(A="1")] struct Bad { value: bool }', "one primitive integer field"),
        ('#[flags(A="1")] struct Bad { value: Vec[u8] }', "one primitive integer field"),
        ('struct Bad { value: u8 }', "requires #[flags(...)]"),
        ('#[flags(A="1")] #[flags(B="2")] struct Bad { value: u8 }', "exactly one flags attribute"),
        ('#[flags(A="1", A="2")] struct Bad { value: u8 }', "duplicate flag name: A"),
        ('#[flags(A="256")] struct Bad { value: u8 }', "out-of-range flag mask"),
        ('#[flags(A="0x10000000000000000")] struct Bad { value: u64 }', "out-of-range flag mask"),
        ('#[flags(A="MISSING")] struct Bad { value: u8 }', "out-of-range flag mask"),
        ('#[flags(A="B", B="1")] struct Bad { value: u8 }', "out-of-range flag mask"),
        ('#[flags(A="1|")] struct Bad { value: u8 }', "empty flag mask component"),
        ('#[flags(A="-1")] struct Bad { value: u8 }', "out-of-range flag mask"),
        ('#[flags(A=invalid)] struct Bad { value: u8 }', 'require NAME = "mask"'),
        ('#[flags("1")] struct Bad { value: u8 }', 'require NAME = "mask"'),
    ]
    artifacts = ROOT / "bitflags/_artifact/diagnostics"
    artifacts.mkdir(parents=True, exist_ok=True)
    with tempfile.TemporaryDirectory(dir=artifacts) as temporary:
        for index, (declaration, expected) in enumerate(cases):
            directory = Path(temporary) / str(index)
            directory.mkdir()
            (directory / "goml.toml").write_text('[module]\npath="diagnostics::bitflags"\n[dependencies]\n"ecosystem::bitflags"="0.1.0"\n')
            (directory / "main.gom").write_text('package main;\nuse ecosystem::bitflags;\n#[derive(bitflags::Flags)]\n' + declaration + '\nfn main() -> () {}\n')
            subprocess.run([goml, "fmt"], cwd=directory, env=environment, check=True, capture_output=True, text=True)
            result = subprocess.run([goml, "check"], cwd=directory, env=environment, capture_output=True, text=True, timeout=60)
            output = result.stdout + result.stderr
            (artifacts / f"case-{index}.log").write_text(output)
            if result.returncode == 0 or expected not in output:
                raise AssertionError(f"case {index} expected {expected!r}:\n{output}")
    print(f"bitflags derive diagnostics: {len(cases)} cases passed")


if __name__ == "__main__":
    main()
