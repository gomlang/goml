import argparse
from collections import Counter
import hashlib
import io
import json
import math
from pathlib import Path
import random
import struct
import subprocess
import sys
import tarfile
import urllib.request


ROOT = Path(__file__).resolve().parents[1]
REFERENCE_URL = "https://files.pythonhosted.org/packages/4d/f2/bfb55a6236ed8725a96b0aa3acbd0ec17588e6a2c3b62a93eb513ed8783f/msgpack-1.1.2.tar.gz"
REFERENCE_SHA256 = "3b60763c1373dd60f398488069bcdc703cd08a711477b5d480eecc9f9626f47e"


def reference():
    directory = ROOT / "_artifact/reference"
    directory.mkdir(parents=True, exist_ok=True)
    archive = directory / "msgpack-1.1.2.tar.gz"
    if not archive.exists():
        data = urllib.request.urlopen(REFERENCE_URL, timeout=30).read()
        if hashlib.sha256(data).hexdigest() != REFERENCE_SHA256:
            raise RuntimeError("MessagePack reference checksum mismatch")
        archive.write_bytes(data)
    data = archive.read_bytes()
    if hashlib.sha256(data).hexdigest() != REFERENCE_SHA256:
        raise RuntimeError("cached MessagePack reference checksum mismatch")
    target = directory / "msgpack-1.1.2"
    with tarfile.open(fileobj=io.BytesIO(data), mode="r:gz") as source:
        for name in ("__init__.py", "fallback.py", "exceptions.py", "ext.py"):
            member = source.getmember(f"msgpack-1.1.2/msgpack/{name}")
            if not member.isfile():
                raise RuntimeError("unexpected reference archive member")
            destination = target / "msgpack" / name
            destination.parent.mkdir(parents=True, exist_ok=True)
            destination.write_bytes(source.extractfile(member).read())
    sys.path.insert(0, str(target))
    import msgpack
    if msgpack.__version__ != "1.1.2" or target not in Path(msgpack.__file__).parents:
        raise RuntimeError("incorrect reference imported")
    return msgpack


class Pairs(list):
    pass


def normalize(value):
    if isinstance(value, float):
        if math.isnan(value):
            return ("nan",)
        return ("float", struct.pack(">d", value).hex())
    if isinstance(value, Pairs):
        return ("map", tuple((normalize(k), normalize(v)) for k, v in value))
    if isinstance(value, list):
        return tuple(normalize(item) for item in value)
    return value


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--consumer", type=Path, default=ROOT / "consumers/msgpack/_artifact/bin/msgpack")
    args = parser.parse_args()
    msgpack = reference()
    cases = []
    packer = msgpack.Packer(use_bin_type=True)

    def pack(value):
        if isinstance(value, Pairs):
            return packer.pack_map_header(len(value)) + b"".join(pack(k) + pack(v) for k, v in value)
        if isinstance(value, list):
            return packer.pack_array_header(len(value)) + b"".join(pack(item) for item in value)
        return packer.pack(value)

    def unpack(data):
        return msgpack.unpackb(data, strict_map_key=False, object_pairs_hook=Pairs)

    def add(group, wire, expected=None, **options):
        request = {"hex": wire.hex(), **options}
        cases.append((group, request, {"hex": (wire if expected is None else expected).hex()}))

    integers = [-2**63, -2**31-1, -2**31, -32769, -32768, -129, -128, -33, -32, -1,
                0, 127, 128, 255, 256, 65535, 65536, 2**32-1, 2**32, 2**63-1, 2**63, 2**64-1]
    for value in integers + [None, False, True, "", "中文😀", b"\0\xff", [], Pairs()]:
        add("scalar", pack(value))
    for width, bit_patterns in ((32, [0, 0x80000000, 0x7f800000, 0xff800000, 0x7fc00123, 1]),
                                (64, [0, 0x8000000000000000, 0x7ff0000000000000, 0xfff0000000000000, 0x7ff8000000000123, 1])):
        for bits in bit_patterns:
            add("float-bits", bytes([0xca if width == 32 else 0xcb]) + bits.to_bytes(width // 8, "big"))
    for length in [0, 1, 2, 4, 8, 16, 31, 32, 255, 256, 65535, 65536]:
        for value in ["a" * length, b"\xff" * length, msgpack.ExtType(42, b"x" * length)]:
            add("blob-lengths", pack(value))
    for length in [0, 15, 16, 65535, 65536]:
        add("container-lengths", pack([None] * length))
        add("container-lengths", pack(Pairs([(False, None)] * length)))
    add("arbitrary-map", pack(Pairs([([1, 2], "array"), (None, 1), (None, 2), (b"x", False)])))

    for tag in range(256):
        if tag == 0xc1:
            cases.append(("reserved-tag", {"hex": "c1"}, None))
            continue
        if tag < 0x80 or tag >= 0xe0 or tag in [0xc0, 0xc2, 0xc3]:
            body = b""
        elif tag < 0x90:
            body = b"\xc0\xc0" * (tag - 0x80)
        elif tag < 0xa0:
            body = b"\xc0" * (tag - 0x90)
        elif tag < 0xc0:
            body = b"a" * (tag - 0xa0)
        elif tag in [0xc4, 0xc5, 0xc6]:
            body = bytes(1 << (tag - 0xc4))
        elif tag in [0xc7, 0xc8, 0xc9]:
            body = (3).to_bytes(1 << (tag - 0xc7), "big") + b"\x05abc"
        elif tag == 0xca:
            body = struct.pack(">f", 1.5)
        elif tag == 0xcb:
            body = struct.pack(">d", 1.5)
        elif 0xcc <= tag <= 0xd3:
            body = bytes(1 << ((tag - 0xcc) % 4))
        elif 0xd4 <= tag <= 0xd8:
            body = b"\x05" + b"a" * (1 << (tag - 0xd4))
        elif 0xd9 <= tag <= 0xdb:
            body = bytes(1 << (tag - 0xd9))
        else:
            body = bytes(2 if tag in [0xdc, 0xde] else 4)
        wire = bytes([tag]) + body
        expected = wire if tag == 0xca else pack(unpack(wire))
        add("all-tags", wire, expected)
        for end in range(len(wire)):
            cases.append(("truncation", {"hex": wire[:end].hex()}, None))

    random_source = random.Random(20260920)

    def generated(depth):
        choice = random_source.randrange(9 if depth else 6)
        if choice == 0:
            return random_source.choice(integers)
        if choice == 1:
            return random_source.choice([None, True, False])
        if choice == 2:
            return "".join(random_source.choices("ab中😀\n\0", k=random_source.randrange(20)))
        if choice == 3:
            return random_source.randbytes(random_source.randrange(20))
        if choice == 4:
            return random_source.uniform(-1e12, 1e12)
        if choice == 5:
            return msgpack.ExtType(random_source.randrange(128), random_source.randbytes(random_source.randrange(20)))
        if choice == 6:
            return [generated(depth - 1) for _ in range(random_source.randrange(6))]
        return Pairs([(generated(depth - 1), generated(depth - 1)) for _ in range(random_source.randrange(6))])

    random_values = [generated(4) for _ in range(400)]
    for value in random_values:
        add("recursive-reference", pack(value))
    for seconds, nanoseconds in [(0, 0), (2**32-1, 0), (2**32, 0), (0, 1), (2**34-1, 999999999), (2**34, 0), (-1, 500000000), (-2**63, 0), (2**63-1, 999999999)]:
        wire = pack(msgpack.Timestamp(seconds, nanoseconds))
        cases.append(("timestamp", {"hex": wire.hex(), "operation": "timestamp"}, {"hex": wire.hex(), "seconds": str(seconds), "nanoseconds": nanoseconds}))
    for raw in [b"\xa1\xff", b"\xa3\xed\xa0\x80", b"\xa2\xc0\x80"]:
        cases.append(("invalid-utf8", {"hex": raw.hex()}, None))
        add("raw-utf8", raw, raw=True)
    for wire in [b"\x01\x02", b"\xdd\xff\xff\xff\xff", b"\xc6\xff\xff\xff\xff", b"\x91" * 130 + b"\xc0"]:
        cases.append(("invalid-input", {"hex": wire.hex()}, None))

    states = ["Idle", {"Move": [-7, 9]}, {"Named": {"value": 65535, "label": "中文"}}]
    for index in range(80):
        variant = index % 3
        present = index % 2 == 0
        message = {"id": random_source.randrange(2**64), "name": "😀" + str(index), "active": present,
                   "payload": random_source.randbytes(index), "points": [[-index, index + .25], [index, -0.0]],
                   "state": states[variant], "maybe": "present" if present else None}
        for compact in [False, True]:
            for tagged in [False, True]:
                source = dict(message)
                if compact:
                    source["state"] = variant if variant == 0 else {variant: [-7, 9] if variant == 1 else [65535, "中文"]}
                if tagged:
                    source["maybe"] = [1, "present"] if present else [0]
                expected = pack(list(source.values()) if compact else source)
                source["unknown"] = {"nested": msgpack.ExtType(5, b"opaque")}
                shuffled = dict(reversed(list(source.items())))
                add("typed-serde", pack(shuffled), expected, operation="typed", compact=compact, tagged=tagged)
                if compact:
                    add("typed-positional", expected, expected, operation="typed", compact=True, tagged=tagged)
    malformed_message = {"id": -1, "name": "x", "active": True, "payload": b"x", "points": [], "state": "Idle", "maybe": None}
    cases.append(("typed-error", {"hex": pack(malformed_message).hex(), "operation": "typed"}, None))
    cases.append(("typed-error", {"hex": "80", "operation": "typed"}, None))

    stream_values = [pack(value) for value in random_values[:40]]
    stream = b"".join(stream_values)
    for chunk in [1, 2, 3, 7, 31, 256, 1024, len(stream)]:
        cases.append(("stream", {"hex": stream.hex(), "operation": "stream", "chunk": chunk}, {"values": [value.hex() for value in stream_values]}))
    tiny = pack(None) * 20000
    cases.append(("stream-many", {"hex": tiny.hex(), "operation": "stream", "chunk": len(tiny)}, {"values": ["c0"] * 20000}))
    cases.append(("stream-truncated", {"hex": "a36162", "operation": "stream", "chunk": 1}, None))

    result = subprocess.run([str(args.consumer), "--json"], input=json.dumps([request for _, request, _ in cases]), text=True, capture_output=True, timeout=60)
    if result.returncode:
        raise RuntimeError(result.stderr)
    actual = json.loads(result.stdout)
    if len(actual) != len(cases):
        raise RuntimeError("wrong response count")
    failures = []
    for index, ((group, request, expected), response) in enumerate(zip(cases, actual)):
        if expected is None:
            valid = isinstance(response.get("error"), str) and bool(response["error"])
        else:
            valid = response == expected
            if valid and "hex" in response and not request.get("raw"):
                wire = bytes.fromhex(response["hex"])
                expected_wire = bytes.fromhex(expected["hex"])
                valid = normalize(unpack(wire)) == normalize(unpack(expected_wire))
        if not valid:
            failures.append({"index": index, "group": group, "request": request, "expected": expected, "actual": response})
    report = ROOT / "_artifact/verification/msgpack/interoperability.json"
    report.parent.mkdir(parents=True, exist_ok=True)
    counts = dict(Counter(group for group, _, _ in cases))
    report.write_text(json.dumps({"reference": "msgpack-python 1.1.2 (pure Python)", "sha256": REFERENCE_SHA256, "cases": len(cases), "groups": counts, "failures": failures}, indent=2) + "\n")
    print(json.dumps({"cases": len(cases), "groups": counts, "failures": len(failures)}, sort_keys=True))
    if failures:
        raise RuntimeError(f"{len(failures)} interoperability failures; see {report}")


if __name__ == "__main__":
    main()
