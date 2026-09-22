import hashlib
import json
from pathlib import Path
from urllib.request import urlopen


SOURCE = "https://publicsuffix.org/list/public_suffix_list.dat"
SHA256 = "e81c6f5f11359a79a2479238e732e08bd8521071fd95ee47053471e3426d7b54"
TARGET = Path(__file__).resolve().parents[1] / "cookies" / "builtin.gom"


def ascii_rule(rule):
    prefix = "!" if rule.startswith("!") else "*." if rule.startswith("*.") else ""
    domain = rule[len(prefix):]
    labels = [
        label if label.isascii() else "xn--" + label.encode("punycode").decode("ascii")
        for label in domain.split(".")
    ]
    return prefix + ".".join(labels)


def main():
    with urlopen(SOURCE, timeout=30) as response:
        source = response.read()
    if hashlib.sha256(source).hexdigest() != SHA256:
        raise ValueError("public suffix list checksum changed")
    rules = [
        line.strip()
        for line in source.decode("utf-8").splitlines()
        if line.strip() and not line.lstrip().startswith("//")
    ]
    expanded = []
    seen = set()
    for rule in rules:
        for variant in (rule, ascii_rule(rule)):
            if variant not in seen:
                seen.add(variant)
                expanded.append(variant)
    rules = expanded
    chunks = []
    current = ""
    for rule in rules:
        entry = rule + "\n"
        if len(current.encode("utf-8")) + len(entry.encode("utf-8")) > 4096:
            chunks.append(current)
            current = ""
        current += entry
    if current:
        chunks.append(current)
    output = ["package cookies;", "", "fn builtin_rules() -> string {"]
    output.append("    " + json.dumps(chunks[0], ensure_ascii=False))
    output.extend("        + " + json.dumps(chunk, ensure_ascii=False) for chunk in chunks[1:])
    output.extend(
        [
            "}",
            "",
            "pub fn default_suffix_list() -> Option[SuffixList] {",
            f"    SuffixList::parse(builtin_rules(), 200000, {len(rules)})",
            "}",
            "",
        ]
    )
    TARGET.write_text("\n".join(output), encoding="utf-8")


if __name__ == "__main__":
    main()
