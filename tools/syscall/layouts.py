import argparse
import json
import pathlib
import subprocess


ROOT = pathlib.Path(__file__).resolve().parents[2]
MANIFEST = pathlib.Path(__file__).with_name("linux-amd64-layouts.json")
TARGET = ROOT / "lib/std/os/linux/abi/records.gom"
WIDTHS = {"i16": 2, "u16": 2, "i32": 4, "u32": 4, "i64": 8, "u64": 8, "usize": 8}


def generate(records):
    sizes = dict(WIDTHS)
    sizes.update((record["name"], record["size"]) for record in records)
    lines = ["package abi;", "", "use module::bytes::endian;", ""]
    for record in records:
        name, size = record["name"], record["size"]
        lines.append(f"pub struct {name} {{")
        for field in record["fields"]:
            lines.append(f'    pub {field["name"]}: {field["type"]},')
        lines.extend(["}", "", f"impl {name} {{", f"    pub fn size() -> isize {{", f"        {size}", "    }", ""])
        lines.extend([f"    pub fn encode(self: {name}) -> Vec[byte] {{", f"        let data = zeroed({size});"])
        occupied = set()
        for field in record["fields"]:
            key, ty, offset = field["name"], field["type"], field["offset"]
            width = sizes[ty]
            positions = set(range(offset, offset + width))
            if offset < 0 or offset + width > size or occupied & positions:
                raise SystemExit(f"invalid or overlapping field: {name}.{key}")
            occupied |= positions
            if ty in WIDTHS:
                codec = "u64" if ty == "usize" else ty
                value = f"self.{key}.to_u64()" if ty == "usize" else f"self.{key}"
                lines.append(f"        let _ = endian::write_{codec}(data.as_mut_slice(), {offset}, {value}, endian::Endian::Little);")
            else:
                lines.append(f"        let _ = data.as_mut_slice().copy_from({offset}, self.{key}.encode().as_slice());")
        lines.extend(["        data", "    }", "", f"    pub fn decode(data: Slice[byte]) -> Result[{name}, endian::BoundsError] {{", f"        require_size(data, {size})?;", f"        Result::Ok({name} {{"])
        for field in record["fields"]:
            key, ty, offset = field["name"], field["type"], field["offset"]
            if ty in WIDTHS:
                codec = "u64" if ty == "usize" else ty
                value = f"endian::read_{codec}(data, {offset}, endian::Endian::Little)?"
                if ty == "usize":
                    value += ".to_usize()"
            else:
                value = f"{ty}::decode(data.sub({offset}, {offset + sizes[ty]}))?"
            lines.append(f"            {key}: {value},")
        lines.extend(["        })", "    }", "}", ""])
    return "\n".join(lines)


def verify_headers(records):
    lines = ["#define _GNU_SOURCE", "#include <stddef.h>", "#include <stdint.h>", "#include <time.h>", "#include <sys/time.h>", "#include <sys/stat.h>", "#include <sys/resource.h>", "#include <sys/socket.h>", "#include <sys/uio.h>", "#include <sys/epoll.h>", "#include <poll.h>", "#include <fcntl.h>", "#include <linux/openat2.h>", "#if !defined(__linux__) || !defined(__x86_64__)", '#error "Linux amd64 required"', "#endif", '_Static_assert(sizeof(void *) == 8, "64-bit pointers required");']
    sizes = dict(WIDTHS)
    sizes.update((record["name"], record["size"]) for record in records)
    for record in records:
        c_type = record["c_type"]
        lines.append(f'_Static_assert(sizeof({c_type}) == {record["size"]}, "{record["name"]} size");')
        for field in record["fields"]:
            label = f'{record["name"]}.{field["name"]}'
            lines.append(f'_Static_assert(offsetof({c_type}, {field["c_name"]}) == {field["offset"]}, "{label} offset");')
            lines.append(f'_Static_assert(sizeof((({c_type} *)0)->{field["c_name"]}) == {sizes[field["type"]]}, "{label} size");')
    lines.append("int main(void) { return 0; }")
    directory = ROOT / "_artifact/syscall-layout-check"
    directory.mkdir(parents=True, exist_ok=True)
    source = directory / "layouts.c"
    source.write_text("\n".join(lines) + "\n")
    subprocess.run(["cc", "-std=c11", "-Wall", "-Werror", str(source), "-o", str(directory / "layouts")], check=True)


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--check", action="store_true")
    parser.add_argument("--verify-headers", action="store_true")
    args = parser.parse_args()
    records = json.loads(MANIFEST.read_text())
    output = generate(records)
    if args.check:
        result = subprocess.run([str(ROOT / "stage2/bin/gomlfmt")], input=output, text=True, capture_output=True, check=True)
        if TARGET.read_text() != result.stdout:
            raise SystemExit("regenerate ABI records with tools/syscall/layouts.py and goml fmt")
    else:
        TARGET.write_text(output)
    if args.verify_headers:
        verify_headers(records)


if __name__ == "__main__":
    main()
