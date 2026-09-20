import hashlib
import struct
from pathlib import Path


ROOT = Path(__file__).resolve().parent / "fixtures"


def block(version, width, types, transitions):
    names = bytearray()
    records = bytearray()
    for seconds, dst, name in types:
        records.extend(struct.pack(">iBB", seconds, dst, len(names)))
        names.extend(name.encode() + b"\0")
    header = b"TZif" + version + bytes(15) + struct.pack(">6I", 0, 0, 0, len(transitions), len(types), len(names))
    times = b"".join(struct.pack(">q" if width == 8 else ">i", time) for time, _ in transitions)
    return header + times + bytes(index for _, index in transitions) + records + names


def build():
    normal = [(0, 0, "STD"), (3600, 1, "DST")]
    fixtures = {
        "AllYear": ([(-10800, 0, "XXX"), (-14400, 1, "EDT")], [], "XXX3EDT4,0/0,J365/23"),
        "Finite": (normal, [(0, 1)], ""),
        "Shifted": (normal, [], "STD0DST,J60/-2,J300/26"),
        "Extreme": (normal, [], "STD0DST,M1.1.0/-167,M12.5.0/167"),
        "Julian": (normal, [], "STD0DST,59,300"),
        "NegativeDst": ([(3600, 0, "IST"), (0, 1, "GMT")], [], "IST-1GMT0,M10.5.0,M3.5.0/1"),
        "MultipleFold": ([(0, 0, "ZERO"), (-3600, 0, "ONE"), (-7200, 0, "TWO")], [(0, 1), (1800, 2)], "TWO2"),
    }
    target = ROOT / "Synthetic"
    target.mkdir(exist_ok=True)
    for name, (types, transitions, tail) in fixtures.items():
        data = block(b"3", 4, types, []) + block(b"3", 8, types, transitions) + b"\n" + tail.encode() + b"\n"
        (target / name).write_bytes(data)
    (target / "VersionOne").write_bytes(block(b"\0", 4, [(0, 0, "UTC")], []))
    (target / "VersionFour").write_bytes(block(b"4", 4, [(0, 0, "UTC")], []) + block(b"4", 8, [(0, 0, "UTC")], []) + b"\nUTC0\n")
    paths = sorted(path for path in ROOT.rglob("*") if path.is_file() and path.name not in {"VERSION", "SHA256SUMS"})
    (ROOT / "SHA256SUMS").write_text("".join(f"{hashlib.sha256(path.read_bytes()).hexdigest()}  {path.relative_to(ROOT)}\n" for path in paths))


if __name__ == "__main__":
    build()
