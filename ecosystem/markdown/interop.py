import argparse
from collections import Counter
import hashlib
import html.entities
import json
from pathlib import Path
import subprocess
import urllib.request


ROOT = Path(__file__).resolve().parents[1]
REFERENCE_URL = "https://spec.commonmark.org/0.31.2/spec.json"
REFERENCE_SHA256 = "d431b29d97b6f73e69d547109cf5081578fac931e72afe95639ebe766c1b2a20"


def corpus():
    cache = ROOT / "_artifact/reference/commonmark-0.31.2.json"
    if not cache.exists():
        cache.parent.mkdir(parents=True, exist_ok=True)
        data = urllib.request.urlopen(REFERENCE_URL, timeout=30).read()
        if hashlib.sha256(data).hexdigest() != REFERENCE_SHA256:
            raise RuntimeError("CommonMark reference checksum mismatch")
        cache.write_bytes(data)
    data = cache.read_bytes()
    if hashlib.sha256(data).hexdigest() != REFERENCE_SHA256:
        raise RuntimeError("cached CommonMark reference checksum mismatch")
    return json.loads(data)


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--consumer", type=Path, default=ROOT / "consumers/markdown/_artifact/bin/markdown")
    args = parser.parse_args()
    cases = corpus()
    result = subprocess.run(
        [str(args.consumer), "--json"], input=json.dumps([case["markdown"] for case in cases]),
        text=True, capture_output=True, timeout=60,
    )
    if result.returncode:
        raise RuntimeError(result.stderr)
    actual = json.loads(result.stdout)
    if len(actual) != len(cases):
        raise RuntimeError("consumer returned wrong number of results")
    totals = Counter(case["section"] for case in cases)
    failures = []
    for case, rendered in zip(cases, actual):
        if rendered != case["html"]:
            failures.append({"example": case["example"], "section": case["section"], "markdown": case["markdown"], "expected": case["html"], "actual": rendered})
    failed = Counter(case["section"] for case in failures)
    report = ROOT / "_artifact/verification/markdown/commonmark.json"
    report.parent.mkdir(parents=True, exist_ok=True)
    report.write_text(json.dumps({"reference": REFERENCE_URL, "sha256": REFERENCE_SHA256, "total": len(cases), "passed": len(cases) - len(failures), "sections": {section: {"passed": count - failed[section], "total": count} for section, count in totals.items()}, "failures": failures}, ensure_ascii=False, indent=2) + "\n")
    print(f"CommonMark 0.31.2: {len(cases) - len(failures)}/{len(cases)} examples passed")
    for section, count in totals.items():
        if failed[section]:
            print(f"  {section}: {count - failed[section]}/{count}")
    for failure in failures[:8]:
        print(json.dumps(failure, ensure_ascii=False))
    print(f"Full report: {report}")
    if failures:
        raise SystemExit(1)
    entities = sorted((name, value) for name, value in html.entities.html5.items() if name.endswith(";"))
    result = subprocess.run([str(args.consumer), "--json"], input=json.dumps(["&" + name for name, _ in entities]), text=True, capture_output=True, timeout=60)
    if result.returncode:
        raise RuntimeError(result.stderr)
    outputs = json.loads(result.stdout)
    assert len(outputs) == len(entities)
    for (name, value), output in zip(entities, outputs):
        escaped = value.replace("&", "&amp;").replace("<", "&lt;").replace(">", "&gt;").replace('"', "&quot;")
        assert output == "<p>" + escaped + "</p>\n", (name, output, escaped)
    print(f"HTML5 entities: {len(entities)} mappings checked against Python's source table")


if __name__ == "__main__":
    main()
