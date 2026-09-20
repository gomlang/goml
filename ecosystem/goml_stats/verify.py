import argparse
import json
import os
from pathlib import Path
import random
import subprocess
import tempfile


ROOT = Path(__file__).resolve().parent
COUNTERS = ("files", "test_files", "lines", "code", "comments", "blank", "bytes")


def counts(source, code, comments, blank, test=False):
    return dict(zip(COUNTERS, (1, int(test), code + comments + blank,
                               code, comments, blank, len(source.encode()))))


def total(items):
    items = list(items)
    return {key: sum(item[key] for item in items) for key in COUNTERS}


def main():
    parser = argparse.ArgumentParser(description="Build and verify the GoML statistics tool")
    parser.add_argument("--goml", type=Path, default=ROOT.parents[1] / "stage2/bin/goml")
    arguments = parser.parse_args()
    goml = arguments.goml.resolve()
    artifact = ROOT / "_artifact/verification"
    artifact.mkdir(parents=True, exist_ok=True)
    log = (artifact / "commands.log").open("w")
    commands = []

    def run(command, cwd=ROOT, expected=0):
        result = subprocess.run([str(value) for value in command], cwd=cwd,
                                capture_output=True, text=True, timeout=180)
        log.write(f"$ {command!r}\n{result.stdout}{result.stderr}\n")
        log.flush()
        commands.append({"command": [str(value) for value in command],
                         "exit_code": result.returncode, "expected": expected})
        assert result.returncode == expected, (command, result.returncode, result.stdout, result.stderr)
        return result

    run([goml, "fmt", "--check"])
    run([goml, "test", "--timeout", "60s"])
    run([goml, "build"])
    binary = ROOT / "_artifact/bin/cmd/goml_stats/goml_stats"

    with tempfile.TemporaryDirectory(prefix="project-", dir=artifact) as temporary:
        root = Path(temporary)

        def write(name, value):
            target = root / name
            target.parent.mkdir(parents=True, exist_ok=True)
            target.write_bytes(value.encode() if isinstance(value, str) else value)
            return target

        sources = {
            "main.gom": ("package main;\n\n// docs\nfn main() -> () {}\n", 2, 1, 1, False),
            "src/core.gom": ('let url = "https://example.invalid"; // inline\n', 1, 0, 0, False),
            "src/core_test.gom": ("// test\r\n\r\nfn check() -> () {}", 1, 1, 1, True),
            "tests/api.gom": ("fn test() -> () {}\n", 1, 0, 0, True),
            "vendor/nested/lib.gom": ('let s = br#"first\n// literal\n\nlast"#;\n// outside', 4, 1, 0, False),
            "empty.gom": ("", 0, 0, 0, False),
            'unicode/中文 "name".gom': ('let s = "😀";\r// 中文\r', 1, 1, 0, False),
        }
        expected = {}
        for name, (source, code, comments, blank, test) in sources.items():
            write(name, source)
            expected[name] = counts(source, code, comments, blank, test)
        for name in ("goml.toml", "vendor/nested/goml.toml", "empty_module/goml.toml"):
            write(name, '[module]\npath = "fixture"\n')
        defaults = (".git", ".hg", ".svn", ".goml", "_artifact", "_bootstrap",
                    "stage0", "stage1", "stage2", "stage3", "node_modules", "__pycache__")
        for excluded in defaults:
            write(f"{excluded}/ignored.gom", "ignored\n")
        write("src/_artifact/ignored.gom", "ignored\n")
        write("README.md", "ignored\n")
        write("generated.go", "ignored\n")
        (root / "link.gom").symlink_to(root / "main.gom")
        (root / "loop").symlink_to(root, target_is_directory=True)
        (root / "dangling.gom").symlink_to(root / "missing")
        os.mkfifo(root / "pipe.gom")

        def report(*options, target=root, cwd=ROOT):
            result = run([binary, "--json", "--files", *options, target], cwd=cwd)
            assert result.stderr == "", result.stderr
            return json.loads(result.stdout)

        def check(report_value, wanted):
            assert report_value["totals"] == total(wanted.values()), report_value["totals"]
            files = {entry["path"]: entry["counts"] for entry in report_value["files"]}
            assert files == wanted, (files, wanted)
            assert [entry["path"] for entry in report_value["files"]] == sorted(wanted)
            assert report_value["package_count"] == len({str(Path(name).parent) for name in wanted})
            assert total(entry["counts"] for entry in report_value["packages"]) == report_value["totals"]
            assert total([report_value["unassigned"], *(entry["counts"] for entry in report_value["modules"])]) == report_value["totals"]
            for group in report_value["packages"]:
                assert group["counts"] == total(value for name, value in wanted.items()
                                                if str(Path(name).parent) == group["path"])

        original = report()
        check(original, expected)
        assert original["schema_version"] == 1
        assert original["root"] == str(root)
        assert original["module_count"] == 3
        modules = {entry["path"]: entry["counts"] for entry in original["modules"]}
        assert modules["empty_module"] == total([])
        assert modules["vendor/nested"] == expected["vendor/nested/lib.gom"]
        for file in original["files"]:
            assert file["module"] == ("vendor/nested" if file["path"].startswith("vendor/nested/") else ".")
        assert report() == original
        assert report(target=".", cwd=root) == original
        default_path = run([binary, "--json", "--files"], cwd=root)
        assert json.loads(default_path.stdout) == original

        for flags, removed in [
            (("--exclude", "src"), {"src/core.gom", "src/core_test.gom"}),
            (("--exclude=core.gom",), {"src/core.gom"}),
            (("--exclude", "vendor/nested/"), {"vendor/nested/lib.gom"}),
            (("--exclude", "src/core.gom", "--exclude", "tests"), {"src/core.gom", "tests/api.gom"}),
            (("--exclude", "missing"), set()),
        ]:
            check(report(*flags), {name: value for name, value in expected.items() if name not in removed})
        without_nested = report("--exclude", "vendor/nested")
        assert without_nested["module_count"] == 2
        no_manifests = report("--exclude", "goml.toml")
        check(no_manifests, expected)
        assert no_manifests["module_count"] == 0
        assert no_manifests["unassigned"] == original["totals"]
        assert all(file["module"] is None for file in no_manifests["files"])

        all_expected = dict(expected)
        for name in (*defaults, "src/_artifact"):
            all_expected[f"{name}/ignored.gom"] = counts("ignored\n", 1, 0, 0)
        check(report("--no-default-excludes"), all_expected)
        without_stage = {name: value for name, value in all_expected.items() if not name.startswith("stage2/")}
        check(report("--no-default-excludes", "--exclude", "stage2"), without_stage)
        explicit_excluded_root = report(target=root / "_artifact")
        check(explicit_excluded_root, {"ignored.gom": counts("ignored\n", 1, 0, 0)})

        for name in ("main.gom", "empty.gom", "tests/api.gom"):
            single = report(target=root / name)
            check(single, {Path(name).name: expected[name]})
            assert single["module_count"] == 0
        empty = root / "empty_directory"
        empty.mkdir()
        check(report(target=empty), {})
        dash = write("-dash.gom", "x\n")
        check(report("--", target=dash.name, cwd=root), {dash.name: counts("x\n", 1, 0, 0)})
        dash.unlink()

        no_files = json.loads(run([binary, "--json", root]).stdout)
        assert "files" not in no_files
        for flags, expected_text in [((), "Total"), (("--files",), "src/core.gom"),
                                     (("--modules",), "vendor/nested"),
                                     (("--packages",), "Package directory")]:
            output = run([binary, *flags, root]).stdout
            assert expected_text in output and "Modules: 3" in output and "Test files: 2" in output
        assert "Usage: goml_stats" in run([binary, "--help"]).stdout
        assert "Usage: goml_stats" in run([binary, "-h"]).stdout

        for flags in (("--unknown",), ("--exclude",), ("--exclude=",),
                      ("--exclude", "../outside"), ("--exclude", "/absolute"),
                      ("--exclude", "."), (root, root), ("",)):
            failure = run([binary, *flags], expected=2)
            assert failure.stdout == "" and "goml_stats:" in failure.stderr
        for name in ("missing", "README.md", "link.gom", "loop", "pipe.gom"):
            failure = run([binary, root / name], expected=1)
            assert failure.stdout == "" and "goml_stats:" in failure.stderr
        invalid = write("invalid.gom", b"\xff\xfe")
        failure = run([binary, "--json", root], expected=1)
        assert failure.stdout == "" and "invalid.gom" in failure.stderr
        invalid.unlink()

        rng = random.Random(20260920)
        blocks = [(["// comment"], (0, 1, 0)), (["", " \t"], (0, 0, 2)),
                  (['let url = "https://example.invalid"; // inline'], (1, 0, 0)),
                  (['let s = r##"first', '// literal', '', '"#', 'last"##;'], (5, 0, 0)),
                  (['let s =', '  \\\\r#"', '  \\\\// literal', ';'], (4, 0, 0))]
        generated = {}
        for index in range(48):
            lines, expected_lines = [], [0, 0, 0]
            for _ in range(rng.randrange(1, 25)):
                block, numbers = rng.choice(blocks)
                lines.extend(block)
                expected_lines = [left + right for left, right in zip(expected_lines, numbers)]
            ending = rng.choice(["\n", "\r\n", "\r"])
            source = ending.join(lines) + ending
            name = f"case_{index:02}.gom"
            write(f"generated_cases/{name}", source)
            generated[name] = counts(source, *expected_lines)
        check(report(target=root / "generated_cases"), generated)

    log.close()
    (artifact / "report.json").write_text(json.dumps({"passed": True, "commands": commands,
                                                    "generated_cases": 48}, indent=2) + "\n")
    print(f"goml_stats: unit tests and {len(commands) - 3} CLI checks passed, including 48 generated sources")


if __name__ == "__main__":
    main()
