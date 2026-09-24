import argparse
import os
import pathlib
import subprocess
import unicodedata


parser = argparse.ArgumentParser()
mode = parser.add_mutually_exclusive_group()
mode.add_argument("--check", action="store_true")
mode.add_argument("--write", action="store_true")
args = parser.parse_args()


def string_literal(value):
    return '"' + value.replace('\\', '\\\\').replace('"', '\\"').replace('\n', '\\n').replace('\r', '\\r').replace('\t', '\\t') + '"'


def char_literal(value):
    return "'" + value.replace('\\', '\\\\').replace("'", "\\'") + "'"


if unicodedata.unidata_version != "15.0.0":
    raise SystemExit("Python Unicode data must be 15.0.0")

folds = []
for codepoint in range(0x110000):
    value = chr(codepoint)
    folded = value.casefold()
    if folded != value:
        folds.append((value, folded))

lines = [
    "package unicode;",
    "",
    "fn case_fold_scalar(value: char) -> string {",
    "    match value {",
]
for source, folded in folds:
    lines.append(f"        {char_literal(source)} => {string_literal(folded)},")
lines.extend([
    "        _ => value.to_string(),",
    "    }",
    "}",
    "",
    "pub fn case_fold(value: string) -> string {",
    "    let result: Vec[byte] = Vec::new();",
    "    for character in value.chars() {",
    "        append_string(result, case_fold_scalar(character))",
    "    }",
    "    finish_string(result)",
    "}",
    "",
])

root = pathlib.Path(__file__).resolve().parents[1]
target = root / "lib/std/unicode/casefold.gom"
formatter = os.environ.get("GOMLFMT", str(root / "stage2/bin/gomlfmt"))
generated = subprocess.run(
    [formatter],
    input="\n".join(lines).encode("utf-8"),
    stdout=subprocess.PIPE,
    check=True,
).stdout
if args.check:
    if target.read_bytes() != generated:
        raise SystemExit("Unicode full casefold source is stale; regenerate with --write")
else:
    target.write_bytes(generated)
