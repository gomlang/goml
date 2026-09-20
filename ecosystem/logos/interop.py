import itertools
import json
from pathlib import Path
import random
import re
import subprocess


ROOT = Path(__file__).resolve().parents[1]
BINARY = ROOT / "consumers" / "logos" / "_artifact" / "bin" / "logos"


def reference(case):
    rules = case["rules"]
    patterns = []
    for rule in rules:
        pattern = re.escape(rule["pattern"]) if rule.get("literal") else rule["pattern"]
        flags = re.ASCII
        if rule.get("ignore_case"):
            flags |= re.IGNORECASE
        if rule.get("dot_all"):
            flags |= re.DOTALL
        patterns.append(re.compile(pattern, flags))
    source = case["source"]
    offsets = [0]
    for character in source:
        offsets.append(offsets[-1] + len(character.encode()))
    result = []
    start = 0
    while start < len(source):
        best_end, best_priority, candidates = start, -1, []
        for index, pattern in enumerate(patterns):
            for end in range(start + 1, len(source) + 1):
                if pattern.fullmatch(source[start:end]) is None:
                    continue
                priority = rules[index]["priority"]
                if end > best_end or (end == best_end and priority > best_priority):
                    best_end, best_priority, candidates = end, priority, [index]
                elif end == best_end and priority == best_priority:
                    candidates.append(index)
        if not candidates:
            raise AssertionError("oracle cases require a complete fallback rule")
        entry = {"start": offsets[start], "end": offsets[best_end], "slice": source[start:best_end]}
        if len(candidates) > 1:
            entry["error"] = "ambiguous"
            result.append(entry)
        elif not rules[candidates[0]].get("skip"):
            entry["rule"] = candidates[0]
            result.append(entry)
        start = best_end
    return result


def main():
    if not BINARY.is_file():
        raise RuntimeError("build the logos consumer with ecosystem/verify.py first")
    rng = random.Random(20260921)
    fallback = {"pattern": ".", "dot_all": True, "priority": 0}
    patterns = [
        "a|ab", "ab|a", "a(b|c)*", "(ab|a)+", "a{1,4}", "a{2,}",
        "[ab]+", "[^b]+", "(?:ab)?a", "(a?)*b", "(a|)*b", "a{0}b",
        "(a|aa)+b", "[a-z][a-z0-9_]*", "[0-9]+", r"\w+", r"\d+", r"\s+",
        r"\D+", r"\W+", r"\S+", "[-a-c]+", "[A-Z]+", "[^\\n]+", ".+",
        r"\x61+", r"\u0061+", "é|é中", "[中é]+", "(😀|a)+",
    ]
    cases = []
    for pattern in patterns[:13]:
        for length in range(7):
            for letters in itertools.product("ab", repeat=length):
                cases.append({"rules": [{"pattern": pattern, "priority": 3}, fallback], "source": "".join(letters)})
    for _ in range(1500):
        rules = []
        for index in range(rng.randrange(1, 5)):
            rule = {"pattern": rng.choice(patterns), "priority": index + 1}
            if rng.randrange(5) == 0:
                rule["ignore_case"] = True
            if rng.randrange(5) == 0:
                rule["dot_all"] = True
            if rng.randrange(6) == 0:
                rule["skip"] = True
            rules.append(rule)
        rules.append(fallback)
        source = "".join(rng.choice("aaabbcABC019_ !\n\té中😀") for _ in range(rng.randrange(0, 25)))
        cases.append({"rules": rules, "source": source})
    cases.extend([
        {"rules": [{"pattern": "a", "priority": 4}, {"pattern": "[ab]", "priority": 4}, fallback], "source": "abba"},
        {"rules": [{"pattern": "a+b", "literal": True, "priority": 4}, fallback], "source": "a+b aaab"},
        {"rules": [{"pattern": "Ab", "literal": True, "ignore_case": True, "priority": 4}, fallback], "source": "aBAbABab"},
        {"rules": [{"pattern": "ab|a", "priority": 4}, {"pattern": "abc", "literal": True, "priority": 1}, fallback], "source": "abcaba"},
    ])
    expected = [reference(case) for case in cases]
    process = subprocess.run([str(BINARY), "--json"], input=json.dumps(cases, ensure_ascii=False), capture_output=True, text=True, timeout=90, check=True)
    actual = json.loads(process.stdout)
    if len(actual) != len(expected):
        raise AssertionError(f"result length {len(actual)} != {len(expected)}")
    for index, (left, right) in enumerate(zip(actual, expected)):
        if left != right:
            raise AssertionError(f"case {index}: {cases[index]!r}\nGoML: {left!r}\nPython: {right!r}")
    print(f"logos interoperability: {len(cases)} cases passed against exhaustive-prefix Python re matching")


if __name__ == "__main__":
    main()
