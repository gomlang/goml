import json
import os
from pathlib import Path
import random
import subprocess
import tempfile


ROOT = Path(__file__).resolve().parents[1]
BINARY = ROOT / "consumers" / "ignore" / "_artifact" / "bin" / "ignore"
FILES = [
    "main.gom", "main.go", "keep.gom", ".hidden.gom", "中文.gom", "é.gom",
    "a.gom", "aa.gom", "a1.gom", "file.tmp", "!literal", "#literal",
    "space ", "space", "star*name", "notes.txt", "src/main.gom",
    "src/keep.gom", "src/a1.gom", "src/file.tmp", "src/generated/main.gom",
    "src/generated/keep.gom", "src/nested/a.gom", "src/nested/keep.gom",
    "build/main.gom", "build/keep.gom", "build/sub/a.gom", "vendor/main.gom",
    "vendor/keep.gom", "vendor/deep/a.gom", "tests/main.gom", "tests/a.go",
    "tests/nested/a.gom", "other/src/main.gom", "other/build/main.gom",
]
PATTERNS = [
    "*.gom", "!keep.gom", "/main.gom", "?.gom", "??.gom", "[a-z].gom",
    "[!a-z].gom", "a[0-9].gom", "**/main.gom", "src/**", "src/*",
    "!src/keep.gom", "src/**/keep.gom", "**/nested/**", "build/",
    "!build/keep.gom", "build/*", "!build/", "vendor/**", "!vendor/keep.gom",
    "*.tmp", "!*.tmp", ".*", "!.*", "**/*.go", "tests/**/a.*",
    "\\!literal", "\\#literal", "space\\ ", "space   ", "star\\*name",
    "# ignored comment", "", "中文.gom", "**/generated/", "!**/generated/",
]


def cases():
    directories = sorted({str(parent) for name in FILES for parent in Path(name).parents if str(parent) != "."})
    paths = [{"path": name, "directory": False} for name in FILES]
    paths += [{"path": name, "directory": True} for name in directories]
    result = []
    for source in [
        "", "build/\n!build/keep.gom\n", "build/*\n!build/keep.gom\n",
        "*\n!src/\nsrc/*\n!src/keep.gom\n", "src/**/main.gom\n",
        "\\!literal\n\\#literal\nspace\\ \nstar\\*name\n",
        "*.gom\n!keep.gom\n", "[a-z].gom\n[!a-z].gom\n",
    ]:
        result.append({"exclude": "", "rules": [{"base": "", "source": source}], "paths": paths})
    rng = random.Random(20260920)
    for _ in range(192):
        rules = [{"base": "", "source": "\n".join(rng.choices(PATTERNS, k=rng.randrange(1, 12))) + "\n"}]
        for base in ["src", "src/nested", "src/generated", "vendor", "tests"]:
            if rng.randrange(3) == 0:
                rules.append({"base": base, "source": "\n".join(rng.choices(PATTERNS, k=rng.randrange(1, 5))) + "\n"})
        excluded = "\n".join(rng.choices(PATTERNS, k=rng.randrange(4))) + "\n"
        result.append({"exclude": excluded, "rules": rules, "paths": paths})
    return result


def reference(case, directory):
    environment = os.environ.copy()
    for name in ("GIT_DIR", "GIT_WORK_TREE", "GIT_INDEX_FILE", "GIT_COMMON_DIR", "GIT_OBJECT_DIRECTORY", "GIT_ALTERNATE_OBJECT_DIRECTORIES"):
        environment.pop(name, None)
    environment.update({"GIT_CONFIG_NOSYSTEM": "1", "GIT_CONFIG_GLOBAL": "/dev/null", "LC_ALL": "C"})
    command = ["git", "-c", "core.excludesFile=/dev/null", "-c", "core.ignoreCase=false", "-C", str(directory)]
    subprocess.run(command + ["init", "--quiet"], env=environment, check=True, capture_output=True, timeout=10)
    (directory / ".git" / "info" / "exclude").write_text(case["exclude"])
    for entry in case["paths"]:
        path = directory / entry["path"]
        if entry["directory"]:
            path.mkdir(parents=True, exist_ok=True)
        else:
            path.parent.mkdir(parents=True, exist_ok=True)
            path.write_text("")
    for rules in case["rules"]:
        parent = directory / rules["base"]
        parent.mkdir(parents=True, exist_ok=True)
        (parent / ".gitignore").write_text(rules["source"])
    paths = [entry["path"] + ("/" if entry["directory"] else "") for entry in case["paths"]]
    checked = subprocess.run(
        command + ["check-ignore", "--no-index", "--stdin", "-z", "--verbose", "--non-matching"],
        input=("\0".join(paths) + "\0").encode(), capture_output=True, env=environment, timeout=10,
    )
    if checked.returncode not in (0, 1):
        raise RuntimeError(checked.stderr.decode())
    fields = checked.stdout.decode().split("\0")
    if fields[-1] != "" or len(fields) != len(paths) * 4 + 1:
        raise AssertionError(f"unexpected git check-ignore output: {checked.stdout!r}")
    result = []
    for index, path in enumerate(paths):
        source, line, pattern, actual = fields[index * 4:index * 4 + 4]
        if actual != path:
            raise AssertionError((path, actual))
        result.append(bool(source) and not pattern.startswith("!"))
    return result


def main():
    inputs = cases()
    expected = []
    with tempfile.TemporaryDirectory(prefix="goml-ignore-reference-") as temporary:
        for index, case in enumerate(inputs):
            directory = Path(temporary) / str(index)
            directory.mkdir()
            expected.append(reference(case, directory))
    result = subprocess.run([str(BINARY), "--json"], input=json.dumps(inputs, ensure_ascii=False), capture_output=True, text=True, check=True, timeout=90)
    actual = json.loads(result.stdout)
    if len(actual) != len(expected):
        raise AssertionError(f"case count {len(actual)} != {len(expected)}")
    for index, (left, right) in enumerate(zip(actual, expected)):
        if left != right:
            raise AssertionError(f"case {index}: {inputs[index]!r}\nGoML: {left!r}\nGit: {right!r}")
    print(f"ignore interoperability: {len(inputs)} rule hierarchies, {sum(len(case['paths']) for case in inputs)} paths agree with Git")


if __name__ == "__main__":
    main()
