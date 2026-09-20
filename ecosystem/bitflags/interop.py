import hashlib
from pathlib import Path
import random
import shutil
import subprocess
import tarfile
import urllib.request


ROOT = Path(__file__).resolve().parents[1]
VERSION = "2.13.2"
CHECKSUM = "3ded4057c258ba199e2d26386d3af3780957ecaee6c4ef4041c6b4b8b97c0b06"


def reference():
    rustc = shutil.which("rustc")
    if rustc is None:
        raise RuntimeError("bitflags reference verification requires rustc")
    artifacts = ROOT / "bitflags" / "_artifact" / "reference"
    artifacts.mkdir(parents=True, exist_ok=True)
    archive = artifacts / f"bitflags-{VERSION}.crate"
    if not archive.exists():
        with urllib.request.urlopen(f"https://static.crates.io/crates/bitflags/bitflags-{VERSION}.crate", timeout=60) as response:
            data = response.read()
        if hashlib.sha256(data).hexdigest() != CHECKSUM:
            raise RuntimeError("bitflags reference checksum mismatch")
        archive.write_bytes(data)
    if hashlib.sha256(archive.read_bytes()).hexdigest() != CHECKSUM:
        raise RuntimeError("cached bitflags reference checksum mismatch")
    with tarfile.open(archive) as source:
        source.extractall(artifacts, filter="data")
    library = artifacts / "libbitflags.rlib"
    subprocess.run([rustc, "--crate-name", "bitflags", "--crate-type", "rlib", "--edition=2021", "--cfg", 'feature="std"', str(artifacts / f"bitflags-{VERSION}" / "src/lib.rs"), "-o", str(library)], check=True)
    binary = artifacts / "oracle"
    subprocess.run([rustc, "--edition=2021", str(ROOT / "bitflags/oracle.rs"), "--extern", f"bitflags={library}", "-o", str(binary)], check=True)
    return binary


def main():
    oracle = reference()
    binary = ROOT / "consumers/bitflags/_artifact/bin/bitflags"
    rng = random.Random(771029)
    cases = []
    profiles = {"access": 8, "overlap": 8, "empty": 8, "external": 64, "wide": 64, "medium": 16, "word": 32}
    texts = ["", " ", "\t\n\r", "\u2003\u00a0", "READ", "NONE", "READ | WRITE", "BOTH", "ALIAS", "GROUP", "A | AB", "HIGH", "LOW | HIGH", "ALL", "unknown", "_", "read", "|", "READ|", "|READ", "READ||WRITE", "0x", "0X01", "0x+1", "0x-1", "0x1_0", "0xff", "0x100", "0xffffffffffffffff", "0x10000000000000000", "\u2003READ | 0x80\u00a0"]
    for profile, width in profiles.items():
        mask = (1 << width) - 1
        values = list(range(256)) if width == 8 else [0, 1, mask, mask - 1, 1 << (width - 1), (1 << (width - 1)) - 1] + [rng.getrandbits(width) for _ in range(350)]
        for a in values:
            for text in [rng.choice(texts), f"0x{a:x}"]:
                cases.append((profile, a, rng.getrandbits(width), text))
        for text in texts:
            cases.append((profile, mask, 1, text))
    payload = "".join(f"{profile}\t{a:x}\t{b:x}\t{text.encode().hex()}\n" for profile, a, b, text in cases)
    expected = subprocess.run([str(oracle)], input=payload, capture_output=True, text=True, check=True, timeout=60).stdout.splitlines()
    result = subprocess.run([str(binary), "--oracle"], input=payload, capture_output=True, text=True, timeout=60)
    if result.returncode:
        raise AssertionError(result.stderr)
    actual = result.stdout.splitlines()
    if len(actual) != len(expected):
        raise AssertionError(f"expected {len(expected)} responses, received {len(actual)}")
    for index, (got, wanted) in enumerate(zip(actual, expected)):
        if got != wanted:
            raise AssertionError(f"case {index} {cases[index]!r}\nGoML: {got!r}\nRust: {wanted!r}")
    print(f"bitflags {VERSION}: {len(cases)} Rust reference comparisons passed across {len(profiles)} profiles")


if __name__ == "__main__":
    main()
