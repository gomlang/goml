import html.entities
import json
from pathlib import Path


def generate():
    grouped = {}
    for key, value in html.entities.html5.items():
        if key.endswith(";"):
            grouped.setdefault(key[0], []).append((key[:-1], value))
    lines = ["package markdown;", "", "fn named_entity(name: string) -> Option[string] {", '    if name == "" { return Option::None }', "    match name.byte_get(0) {"]
    for prefix in sorted(grouped):
        lines.append(f"        {ord(prefix)} => entity_{ord(prefix)}(name),")
    lines.extend(["        _ => Option::None,", "    }", "}"])
    for prefix, entries in sorted(grouped.items()):
        lines.extend(["", f"fn entity_{ord(prefix)}(name: string) -> Option[string] {{", "    match name {"])
        for key, value in sorted(entries):
            lines.append(f"        {json.dumps(key)} => Option::Some({json.dumps(value, ensure_ascii=False)}),")
        lines.extend(["        _ => Option::None,", "    }", "}"])
    Path(__file__).with_name("entities.gom").write_text("\n".join(lines) + "\n")


if __name__ == "__main__":
    generate()
