import json
from pathlib import Path
import subprocess
import sys

ROOT = Path(__file__).resolve().parent
BINARY = ROOT.parent / 'consumers/unicode_text/_artifact/bin/unicode_text'
sys.path.insert(0, str(ROOT))
import generate


def official(name, kind):
    cases = []
    for number, line in enumerate(generate.source(name).splitlines(), 1):
        body = line.split('#', 1)[0].strip()
        if not body:
            continue
        text = ''
        positions = []
        for token in body.split():
            if token == '÷':
                positions.append(len(text.encode('utf-8')))
            elif token != '×':
                text += chr(int(token, 16))
        cases.append(({'kind': kind, 'text': text}, positions, f'{name}:{number}'))
    return cases


def check(cases):
    actual = json.loads(subprocess.run([str(BINARY), '--json'], input=json.dumps([case for case, _, _ in cases]), capture_output=True, text=True, check=True, timeout=120).stdout)
    failures = []
    for (case, expected, context), result in zip(cases, actual, strict=True):
        if result != expected:
            failures.append((context, case, expected, result))
    if failures:
        for context, case, expected, result in failures[:20]:
            print(context, repr(case['text']), 'expected', expected, 'actual', result)
        raise AssertionError(f'{len(failures)} / {len(cases)} conformance cases failed')
    print(f'{len(cases)} cases passed')


def main():
    if generate.generate() != (ROOT / 'properties.gom').read_text():
        raise AssertionError('Unicode generated tables do not match sources')
    selected = sys.argv[1:]
    for name, kind in [('GraphemeBreakTest.txt', 'grapheme'), ('WordBreakTest.txt', 'word'), ('LineBreakTest.txt', 'line')]:
        if not selected or kind in selected:
            cases = official(name, kind)
            print(name, flush=True)
            check(cases)
    samples = [('hello', 5), ('中文', 4), ('e\u0301', 1), ('\u0301', 0), ('👩🏽\u200d💻', 2), ('🇨🇳', 2), ('1\ufe0f\u20e3', 2), ('\x00\x1b\n\t', 0), ('\U0001fae9', 2), ('\u1100\u1161\u11a8', 2), ('\U00016d63\U00016d67', 1)]
    check([({'kind': 'width', 'text': text}, expected, 'display width') for text, expected in samples])
    print('Unicode 16.0.0 conformance and terminal width checks passed')


if __name__ == '__main__':
    main()
