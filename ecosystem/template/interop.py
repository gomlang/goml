import argparse
from collections import Counter
import hashlib
import io
import json
from pathlib import Path
import random
import subprocess
import sys
import tarfile
import urllib.request
import zipfile


ROOT = Path(__file__).resolve().parents[1]
REFERENCES = (
    ("jinja2-3.1.6-py3-none-any.whl", "https://files.pythonhosted.org/packages/62/a1/3d680cbfd5f4b8f15abc1d571870c5fc3e594bb582bc3b64ea099db13e56/jinja2-3.1.6-py3-none-any.whl", "85ece4451f492d0c13c5dd7c13a64681a86afae63a5f347908daf103ce6d2f67"),
    ("markupsafe-3.0.2.tar.gz", "https://files.pythonhosted.org/packages/b2/97/5d42485e71dfc078108a86d6de8fa46db44a1a9295e89c5d6d4a06e23a62/markupsafe-3.0.2.tar.gz", "ee55d3edf80167e48ea11a923c7386f4669df67d7994554387f84e7d8b0a2bf0"),
)


def reference():
    directory = ROOT / "_artifact/reference"
    target = directory / "jinja-3.1.6-reference"
    target.mkdir(parents=True, exist_ok=True)
    for filename, url, digest in REFERENCES:
        cache = directory / filename
        if not cache.exists():
            data = urllib.request.urlopen(url, timeout=30).read()
            if hashlib.sha256(data).hexdigest() != digest:
                raise RuntimeError(f"reference checksum mismatch: {filename}")
            cache.write_bytes(data)
        data = cache.read_bytes()
        if hashlib.sha256(data).hexdigest() != digest:
            raise RuntimeError(f"cached reference checksum mismatch: {filename}")
        if filename.endswith(".whl"):
            with zipfile.ZipFile(io.BytesIO(data)) as archive:
                for name in archive.namelist():
                    relative = Path(name)
                    if relative.is_absolute() or ".." in relative.parts:
                        raise RuntimeError("invalid reference archive path")
                    if name.startswith("jinja2/") and name.endswith(".py"):
                        destination = target / relative
                        destination.parent.mkdir(parents=True, exist_ok=True)
                        destination.write_bytes(archive.read(name))
        else:
            with tarfile.open(fileobj=io.BytesIO(data), mode="r:gz") as archive:
                for name in ("__init__.py", "_native.py"):
                    member = archive.getmember(f"markupsafe-3.0.2/src/markupsafe/{name}")
                    if not member.isfile():
                        raise RuntimeError("unexpected reference archive entry")
                    destination = target / "markupsafe" / name
                    destination.parent.mkdir(parents=True, exist_ok=True)
                    destination.write_bytes(archive.extractfile(member).read())
    sys.path.insert(0, str(target))
    import jinja2
    import markupsafe
    if jinja2.__version__ != "3.1.6" or target not in Path(jinja2.__file__).parents or target not in Path(markupsafe.__file__).parents:
        raise RuntimeError("reference packages loaded from incorrect location")
    return jinja2


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--consumer", type=Path, default=ROOT / "consumers/template/_artifact/bin/template")
    args = parser.parse_args()
    jinja = reference()
    cases = []
    random_source = random.Random(20260920)

    def add(group, source, context=None, **options):
        cases.append((group, {"source": source, "context": context or {}, **options}))

    def expression(depth):
        if depth == 0 or random_source.randrange(4) == 0:
            return str(random_source.randint(-30, 30))
        operator = random_source.choice(["+", "-", "*", "//", "%"])
        right = str(random_source.choice([-13, -3, -1, 1, 2, 7, 19])) if operator in ["//", "%"] else expression(depth - 1)
        return f"({expression(depth - 1)} {operator} {right})"

    for _ in range(250):
        add("integer-arithmetic", "{{ " + expression(3) + " }}")
    for _ in range(100):
        a, b, c = (random_source.randint(-50, 50) for _ in range(3))
        add("comparisons", "{{ 'yes' if a < b <= c else 'no' }}|{{ 'yes' if a not in [b,c] else 'no' }}", {"a": a, "b": b, "c": c})
    for integer in [-2**63, -2**53-1, -2**53, -1, 0, 1, 2**53, 2**53+1, 2**63-1]:
        for number in [-float(2**63), -float(2**53), -0.5, 0.0, 0.5, float(2**53), float(2**63)]:
            add("mixed-number-comparison", "{{ 'lt' if a < b else 'gt' if a > b else 'eq' }}", {"a": integer, "b": number})
    for text in ["", "abc", "A中😀BC", "é<&'\"", "e\u0301"]:
        for start in [None, -8, -2, 0, 1, 7]:
            for stop in [None, -8, -1, 0, 2, 8]:
                for step in [-3, -1, 1, 2]:
                    source = "{{ text[" + ("" if start is None else str(start)) + ":" + ("" if stop is None else str(stop)) + f":{step}]" + " }}"
                    add("unicode-slicing", source, {"text": text})
    for text in ["<x>&\"'", "中文😀", "", "a\nb", "&amp;<script>"]:
        for autoescape in [False, True]:
            for source in ["{{text}}", "{{text|escape}}", "{{text|safe}}", "{{text|safe|forceescape}}", "{{ '<b>'|safe ~ text }}", "{{ ['<i>'|safe, text]|join('&') }}", "{% set captured %}<b>{{text}}</b>{% endset %}{{captured}}"]:
                add("escaping-capture", source, {"text": text}, autoescape=autoescape)
    for _ in range(100):
        values = random_source.choices(range(-20, 20), k=random_source.randrange(20))
        add("loops-and-filters", "{% set outer='kept' %}{% for x in xs if x is odd %}{{loop.index}}/{{loop.revindex}}/{{loop.length}}={{x}};{% set outer='changed' %}{% else %}empty{% endfor %}|{{outer}}|{{xs|unique|sort|join(',')}}", {"xs": values})
    for text in ["alpha", "éCOLE", "中文 AbC", " ABC ", "a\tb\n"]:
        add("text-filters", "{{x|upper}}|{{x|lower}}|{{x|trim}}|{{x|capitalize}}|{{x|length}}|{{x|replace('a','<b>')}}", {"x": text})
    for source in ["A\u00a0{{- 'x' -}}\u3000B", "{{\u00a0'x'\u3000}}", "{{ '\u00a0é\u3000'|trim }}", "{{ '<B>'|safe|lower }}", "A {# ignored {{ #} B", "a \n {{- 'x' -}} \n b", "A{% raw %}{{ bad + }} {% if ??? %}{% endraw %}B", " A {%- raw -%} \n{{ x }} \n{%- endraw -%} B", "{{ '}}' }}|{{ {'x': '%}'}.x }}", "{% with a=2,b=a %}{{a}}/{{b}}{% endwith %}|{{a}}", "{{ missing|default('fallback') }}", "{% if missing is defined %}bad{% else %}ok{% endif %}", "{{ 'ok' or missing }}|{{ '' and missing }}", "\n{{'a\\n\\u4e2d'}}\n"]:
        add("syntax-and-scope", source, {"a": 1})
    for _ in range(30):
        context = {"value": "<" + str(random_source.randrange(10000)) + ">", "title": "Title"}
        templates = {
            "base": "<h1>{% block title %}{{title}}{% endblock %}</h1>{% block body %}base:{% block inner %}inner{% endblock %}{% endblock %}",
            "middle": "{% extends 'base' %}{% block title %}middle+{{super()}}{% endblock %}{% block body %}middle+{{super()}}{% endblock %}",
            "child": "{% extends 'middle' %}{% block title %}child+{{super()}}{% endblock %}{% block inner %}{{value}}{% endblock %}",
        }
        cases.append(("inheritance", {"name": "child", "templates": templates, "context": context}))
    templates = {"piece": "{{value|default('-')}}{% set value='changed' %}", "main": "{% include ['missing','piece'] %}|{% include 'piece' without context %}|{{value}}|{% include 'absent' ignore missing %}"}
    cases.append(("includes", {"name": "main", "templates": templates, "context": {"value": "<parent>"}}))
    for source in ["{{", "{% raw %}x", "{% if true %}", "{% endif %}", "{{ 1e }}", "{{ 'open }}", "{{ 1 + }}", "{% block a %}{% endblock b %}", "{% unknown %}", "{{ absent }}", "{{ 1 / 0 }}", "{{ range(1,2,0) }}", "{{ 'x'[::0] }}", "{% for a,b in [[1]] %}{% endfor %}"]:
        add("errors", source)

    expected = []
    for group, request in cases:
        environment = jinja.Environment(loader=jinja.DictLoader(request.get("templates", {})), autoescape=request.get("autoescape", True), undefined=jinja.StrictUndefined, keep_trailing_newline=True)
        try:
            compiled = environment.get_template(request["name"]) if "name" in request else environment.from_string(request["source"])
            expected.append({"output": compiled.render(request.get("context", {}))})
        except (jinja.TemplateError, TypeError, ValueError, ZeroDivisionError):
            expected.append(None)
    result = subprocess.run([str(args.consumer), "--json"], input=json.dumps([request for _, request in cases]), text=True, capture_output=True, timeout=60)
    if result.returncode:
        raise RuntimeError(result.stderr)
    actual = json.loads(result.stdout)
    if len(actual) != len(cases):
        raise RuntimeError("incorrect consumer result count")
    failures = []
    for index, ((group, request), wanted, got) in enumerate(zip(cases, expected, actual)):
        matches = isinstance(got.get("error"), str) if wanted is None else wanted == got
        if not matches:
            failures.append({"index": index, "group": group, "request": request, "expected": wanted, "actual": got})
    groups = dict(Counter(group for group, _ in cases))
    report = ROOT / "_artifact/verification/template/interoperability.json"
    report.parent.mkdir(parents=True, exist_ok=True)
    report.write_text(json.dumps({"reference": "Jinja 3.1.6 / MarkupSafe 3.0.2", "cases": len(cases), "groups": groups, "failures": failures}, ensure_ascii=False, indent=2) + "\n")
    print(json.dumps({"cases": len(cases), "groups": groups, "failures": len(failures)}))
    if failures:
        raise RuntimeError(f"{len(failures)} differential failures; see {report}")


if __name__ == "__main__":
    main()
