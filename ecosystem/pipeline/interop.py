from collections import Counter
import json
from pathlib import Path
import random
import subprocess


ROOT = Path(__file__).resolve().parents[1]
BINARY = ROOT / 'consumers' / 'pipeline' / '_artifact' / 'bin' / 'pipeline'


def main():
    rng = random.Random(428901)
    requests = []
    expected = []
    comparison = []

    def add(request, result, mode='exact'):
        requests.append(request)
        expected.append(result)
        comparison.append(mode)

    for _ in range(600):
        values = [rng.randrange(-50, 51) for _ in range(rng.randrange(70))]
        mode = rng.choice(['map', 'filter', 'expand'])
        factor = rng.randrange(-4, 5)
        offset = rng.randrange(-10, 11)
        divisor = rng.randrange(1, 8)
        ordered = rng.randrange(2)
        skip = rng.randrange(8) if ordered else 0
        take = rng.randrange(40) if ordered else 1000000
        result = [value * factor + offset for value in values]
        if mode == 'filter':
            result = [value for value in result if value % divisor]
        if mode == 'expand':
            result = [item for value in result for item in (value, -value)]
        result = result[skip:skip + take]
        request = {'mode': mode, 'values': values, 'factor': factor, 'offset': offset, 'divisor': divisor, 'workers': rng.randrange(1, 7), 'capacity': rng.randrange(5), 'window': rng.randrange(1, 15), 'ordered': ordered, 'skip': skip, 'take': take}
        add(request, result, 'exact' if ordered else 'multiset')
    for _ in range(100):
        values = [rng.randrange(-20, 21) for _ in range(rng.randrange(30))]
        size = rng.randrange(1, 9)
        step = rng.randrange(1, 12)
        add({'mode': 'chunks', 'values': values, 'size': size}, [values[i:i + size] for i in range(0, len(values), size)])
        add({'mode': 'windows', 'values': values, 'size': size, 'step': step}, [values[i:i + size] for i in range(0, len(values) - size + 1, step)])
        add({'mode': 'fold', 'values': values}, sum(values))
        sums = []
        total = 0
        for value in values:
            total += value
            sums.append(total)
        add({'mode': 'scan', 'values': values}, sums)
        end = rng.randrange(30)
        add({'mode': 'zip', 'values': values, 'end': end, 'capacity': rng.randrange(5)}, [list(pair) for pair in zip(values, range(end))])
        add({'mode': 'merge', 'values': values, 'end': 100 + end, 'capacity': rng.randrange(5)}, values + list(range(100, 100 + end)), 'multiset')
    for index in range(30):
        values = list(range(30))
        add({'mode': 'error', 'values': values, 'fail': index, 'ordered': index % 2, 'capacity': index % 4}, {'error': 'ErrorKind::Stage', 'stage': 'parallel map', 'index': index, 'cause': 422})
    for request in [
        {'workers': 0}, {'workers': 1025}, {'capacity': -1}, {'window': 0}, {'window': 2097153}, {'skip': -1}, {'take': -1},
        {'mode': 'chunks', 'size': 0}, {'mode': 'windows', 'size': 0}, {'mode': 'windows', 'step': 0}, {'mode': 'zip', 'capacity': -1}, {'limit': -1},
    ]:
        add({**request, 'values': [1, 2, 3]}, 'InvalidConfig', 'error_kind')
    for limit in range(10):
        add({'values': list(range(10)), 'limit': limit}, 'Limit', 'error_kind')
    add({'values': [1], 'timeout': 0}, 'Timeout', 'error_kind')
    result = subprocess.run([str(BINARY), '--json'], input=json.dumps(requests), capture_output=True, text=True, timeout=60)
    if result.returncode:
        raise RuntimeError(result.stdout + result.stderr)
    actual = json.loads(result.stdout)
    if len(actual) != len(expected):
        raise AssertionError('response count mismatch')
    for index, (got, want, mode) in enumerate(zip(actual, expected, comparison)):
        if mode == 'multiset':
            valid = isinstance(got, list) and Counter(got) == Counter(want)
        elif mode == 'error_kind':
            valid = isinstance(got, dict) and got.get('error') == 'ErrorKind::' + want
        else:
            valid = got == want
        if not valid:
            raise AssertionError(f'case {index}: {requests[index]!r}: {got!r} != {want!r}')
    print(f'Pipeline Python sequence oracle: {len(expected)} cases passed')


if __name__ == '__main__':
    main()
