import copy
import json
from pathlib import Path
import random
import subprocess


ROOT = Path(__file__).resolve().parents[1]
BINARY = ROOT / "consumers" / "syntax" / "_artifact" / "bin" / "syntax"


def leaf(kind, text):
    return {"kind": kind, "text": text}


def branch(kind, children):
    return {"kind": kind, "children": children}


def content(tree):
    return tree["text"] if "text" in tree else "".join(map(content, tree["children"]))


def describe(tree):
    result = []
    tokens = []

    def walk(value, path, start):
        length = len(content(value).encode())
        node = "children" in value
        record = {"path": path, "node": node, "kind": value["kind"], "start": start, "end": start + length, "children": len(value.get("children", []))}
        result.append(record)
        if node:
            for index, child in enumerate(value["children"]):
                walk(child, path + [index], start)
                start += len(content(child).encode())
        else:
            tokens.append(record)

    walk(tree, [], 0)
    return result, tokens


def depth(tree):
    return 1 + max(map(depth, tree["children"]), default=0) if "children" in tree else 1


def expected(tree, offsets, slices):
    source = content(tree)
    encoded = source.encode()
    descriptions, tokens = describe(tree)
    boundaries = {0}
    at = 0
    for scalar in source:
        at += len(scalar.encode())
        boundaries.add(at)
    offset_reports = []
    for offset in offsets:
        valid = offset in boundaries
        selected = []
        if valid:
            left = next((token for token in tokens if token["start"] < offset <= token["end"]), None)
            right = next((token for token in tokens if token["start"] <= offset < token["end"]), None)
            if left is not None:
                selected.append(left)
            if right is not None and right is not left:
                selected.append(right)
        offset_reports.append({"offset": offset, "ok": valid, "kinds": [item["kind"] for item in selected], "starts": [item["start"] for item in selected], "ends": [item["end"] for item in selected]})
    slice_reports = []
    for start, end in slices:
        valid = start in boundaries and end in boundaries and start <= end
        slice_reports.append({"ok": valid, "text": encoded[start:end].decode() if valid else ""})
    return {"text": source, "bytes": len(encoded), "elements": len(descriptions), "depth": depth(tree), "descriptions": descriptions, "offsets": offset_reports, "slices": slice_reports}


def commands(tree):
    if "text" in tree:
        return [{"op": "token", "kind": tree["kind"], "text": tree["text"]}]
    result = [{"op": "start", "kind": tree["kind"], "text": ""}]
    for child in tree["children"]:
        result.extend(commands(child))
    result.append({"op": "finish", "kind": 0, "text": ""})
    return result


def apply(tree, replacement):
    if not replacement["path"]:
        return False
    parent = tree
    for index in replacement["path"][:-1]:
        if "children" not in parent or not 0 <= index < len(parent["children"]):
            return False
        parent = parent["children"][index]
    index = replacement["path"][-1]
    if "children" not in parent or not 0 <= index < len(parent["children"]):
        return False
    parent["children"][index] = leaf(replacement["kind"], replacement["text"])
    return True


def tree_cases(rng):
    alphabet = ["", "a", "中", "😀", "é", "\r\n", "\t", "\x00", "é", "\u2028", "# hi\n"]

    def generate(level):
        if level > 5 or rng.random() < 0.45:
            return leaf(rng.randrange(16), "".join(rng.choices(alphabet, k=rng.randrange(5))))
        return branch(rng.randrange(16, 24), [generate(level + 1) for _ in range(rng.randrange(5))])

    trees = [branch(0, []), branch(0, [leaf(1, "")]), branch(0, [leaf(1, ""), leaf(2, "中"), branch(3, []), leaf(4, ""), leaf(5, "😀"), leaf(6, "")])]
    trees.extend(branch(0, [generate(0) for _ in range(rng.randrange(1, 7))]) for _ in range(317))
    inputs, models = [], []
    for initial in trees:
        current = copy.deepcopy(initial)
        length = len(content(initial).encode())
        offsets = list(range(-1, length + 2))
        slices = [[rng.randrange(-2, length + 3), rng.randrange(-2, length + 3)] for _ in range(10)] + [[0, length], [length, length]]
        replacements = []
        reports = []
        for step in range(12):
            paths = [item["path"] for item in describe(current)[0]]
            path = rng.choice(paths) if step % 4 else [rng.randrange(-2, 9), 500]
            replacement = {"path": path, "kind": rng.randrange(32), "text": "".join(rng.choices(alphabet, k=rng.randrange(4)))}
            replacements.append(replacement)
            ok = apply(current, replacement)
            reports.append({"ok": ok, "report": expected(current, offsets, slices)})
        inputs.append({"commands": commands(initial), "replacements": replacements, "offsets": offsets, "slices": slices})
        models.append({"original": expected(initial, offsets, slices), "edits": reports, "snapshot": content(initial)})
    return inputs, models


def config_cases(rng):
    inputs, models = [], []
    for case in range(240):
        values = {}
        locations = {}
        pieces = []
        length = 0

        def write(value):
            nonlocal length
            pieces.append(value)
            length += len(value)

        def scalar():
            return rng.choice([str(rng.randrange(-100000, 100000)), "true", "false", json.dumps(rng.choice(["本地😀", "escaped\nvalue", "quote\"", "\\path", "é"]), ensure_ascii=False)])

        def body(path, level):
            values[path] = []
            for index in range(rng.randrange(1, 5)):
                key = f"key_{index}" if index % 2 else f"名称{index}"
                literal = scalar()
                write("  " * level + key + rng.choice([" = ", "\t=\t", "=", " =\t"]))
                locations[path, key] = (length, length + len(literal))
                values[path].append([key, literal])
                write(literal + rng.choice(["\n", " # trailing\r\n", "\t# 注释😀\n"]))
            if level < 2:
                name = f"section_{level}"
                write("  " * level + name + " { # section\n")
                body(path + (name,), level + 1)
                write("  " * level + "}\n")

        write("# document\r\n")
        body((), 0)
        source = "".join(pieces)
        path = rng.choice(list(values))
        key, old = rng.choice(values[path])
        literal = scalar()
        ok = case % 11 != 0
        if not ok:
            literal = rng.choice(["1 2", '"bad\\q"', "", "# comment", "{"])
        start, end = locations[path, key]
        inputs.append({"text": source, "path": list(path), "key": key, "value": literal})
        models.append({"lossless": source, "diagnostics": [], "ok": ok, "rewritten": source[:start] + literal + source[end:] if ok else source, "snapshot": source, "entries": values[path]})
    return inputs, models


def main():
    rng = random.Random(2026092003)
    trees, tree_models = tree_cases(rng)
    configs, config_models = config_cases(rng)
    executed = subprocess.run([str(BINARY), "--json"], input=json.dumps({"trees": trees, "configs": configs}, ensure_ascii=False), text=True, capture_output=True, check=True, timeout=120)
    actual = json.loads(executed.stdout)
    for name, models in [("trees", tree_models), ("configs", config_models)]:
        if len(actual[name]) != len(models):
            raise AssertionError(f"{name}: case count differs")
        for index, (observed, expected_value) in enumerate(zip(actual[name], models)):
            if observed != expected_value:
                target = ROOT / "syntax" / "_artifact" / "oracle-mismatch.json"
                target.write_text(json.dumps({"kind": name, "index": index, "input": (trees if name == "trees" else configs)[index], "actual": observed, "expected": expected_value}, ensure_ascii=False, indent=2))
                raise AssertionError(f"{name} case {index} differs; see {target}")
    print(f"syntax: {len(trees)} independent tree models, {sum(len(case['replacements']) for case in trees)} persistent edits, and {len(configs)} lossless configuration rewrites passed")


if __name__ == "__main__":
    main()
