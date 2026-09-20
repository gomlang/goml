import os
from pathlib import Path
import stat
import subprocess
import tempfile


ROOT = Path(__file__).resolve().parent
BINARY = ROOT.parent / "consumers" / "tempfile" / "_artifact" / "bin" / "tempfile"


def main():
    if not BINARY.exists():
        raise RuntimeError("Build ecosystem/consumers/tempfile before interop.py")
    for iteration in range(5):
        with tempfile.TemporaryDirectory(prefix="goml-tempfile-interop-") as directory:
            parent = Path(directory)
            environment = os.environ.copy()
            environment["TMPDIR"] = directory
            subprocess.run([str(BINARY)], env=environment, text=True, capture_output=True, timeout=30, check=True)
            assert not list(parent.iterdir()), "default TMPDIR scoped cleanup left resources"
            completed = subprocess.run([str(BINARY), "fs-roundtrip", directory], env=environment, text=True, capture_output=True, timeout=30, check=True)
            lines = completed.stdout.splitlines()
            assert lines[-1] == "tempfile consumer: ok", completed.stdout
            kept = Path(lines[0])
            assert kept.parent == parent and kept.read_bytes() == b"kept"
            assert kept.name.startswith("upload-") and kept.name.endswith(".part")
            assert (parent / "payload.bin").read_bytes() == bytes(range(256))
            assert (parent / "tree" / "nested.txt").read_bytes() == b"retained tree"
            assert stat.S_IMODE((parent / "payload.bin").stat().st_mode) & 0o077 == 0
            assert stat.S_IMODE((parent / "tree").stat().st_mode) & 0o077 == 0
            expected = {"payload.bin", "tree", kept.name}
            for index in range(32):
                name = f"concurrent-{index}"
                expected.add(name)
                assert (parent / name).read_bytes() == str(index).encode()
            assert {entry.name for entry in parent.iterdir()} == expected
    print("Tempfile filesystem interoperability: 5 runs, 160 concurrent persists, binary data, permissions, ownership and cleanup passed")


if __name__ == "__main__":
    main()
