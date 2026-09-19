import argparse
import hashlib
import json
import pathlib
import re
import urllib.request


ROOT = pathlib.Path(__file__).resolve().parents[2]
MANIFEST = pathlib.Path(__file__).with_name("linux-amd64.json")
TARGET = ROOT / "lib/std/os/linux/syscall/constants.gom"


def verify_sources(manifest):
    syscalls = []
    errors = []
    for source in manifest["sources"]:
        url = f'https://raw.githubusercontent.com/torvalds/linux/v{manifest["version"]}/{source["path"]}'
        data = urllib.request.urlopen(url, timeout=30).read()
        if hashlib.sha256(data).hexdigest() != source["sha256"]:
            raise SystemExit(f"checksum mismatch: {url}")
        for line in data.decode().splitlines():
            fields = line.split()
            if source["path"].endswith(".tbl"):
                if len(fields) >= 3 and fields[0].isdigit() and fields[1] in ("common", "64"):
                    syscalls.append([fields[2].upper(), int(fields[0])])
            else:
                match = re.match(r"#define\s+(E\w+)\s+(\w+)", line)
                if match:
                    name, value = match.groups()
                    errors.append([name, int(value) if value.isdigit() else value])
    if syscalls != manifest["syscalls"] or errors != manifest["errno"]:
        raise SystemExit("ABI manifest differs from upstream sources")


def generate(manifest):
    lines = ["package syscall;", ""]
    names = set()
    for name, value in [("SYS_" + name, value) for name, value in manifest["syscalls"]] + manifest["errno"]:
        if name in names or (isinstance(value, str) and value not in names):
            raise SystemExit(f"duplicate constant or unresolved alias: {name}")
        names.add(name)
        lines.extend([f"pub const {name}: usize = {value};", ""])
    lines.extend(["pub fn errno_name(code: usize) -> Option[string] {", "    match code {"])
    for name, value in manifest["errno"]:
        if isinstance(value, int):
            lines.append(f'        {value} => Option::Some("{name}"),')
    lines.extend(["        _ => Option::None,", "    }", "}", ""])
    return "\n".join(lines)


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--check", action="store_true")
    parser.add_argument("--verify-upstream", action="store_true")
    args = parser.parse_args()
    manifest = json.loads(MANIFEST.read_text())
    if args.verify_upstream:
        verify_sources(manifest)
    output = generate(manifest)
    if args.check:
        if TARGET.read_text() != output:
            raise SystemExit(f"regenerate {TARGET.relative_to(ROOT)} with tools/syscall/generate.py")
    else:
        TARGET.write_text(output)


if __name__ == "__main__":
    main()
