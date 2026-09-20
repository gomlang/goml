from collections import Counter
import json
from pathlib import Path
import random
import subprocess


ROOT = Path(__file__).resolve().parents[1]
BINARY = ROOT / 'consumers' / 'tracing' / '_artifact' / 'bin' / 'tracing'


def signature(record):
    return (record['kind'], record['trace_id'], record['span_id'], record['parent_id'],
            record['level'], record['target'], record['name'],
            json.dumps(record['fields'], sort_keys=True, ensure_ascii=False))


def expected_records(scenario):
    output = []
    next_id = 1
    maximum = scenario.get('maximum', 3)
    period = scenario.get('period', 1)
    levels = {1: 'ERROR', 2: 'WARN', 3: 'INFO', 4: 'DEBUG', 5: 'TRACE'}

    def add(kind, trace, span, parent, level, target, name, fields):
        output.append(dict(kind=kind, trace_id=trace, span_id=span, parent_id=parent,
                           level=levels[level], target=target, name=name, fields=fields.copy()))

    for ordinal, request in enumerate(scenario.get('requests', [])):
        root = next_id
        next_id += 1
        sampled = period > 0 and ordinal % period == 0
        root_level = request.get('level', 3)
        visible_root = sampled and root_level <= maximum
        fields = dict(request=ordinal, method='GET', path=f'/items/{ordinal}')
        if visible_root:
            add('span_start', root, root, 0, root_level, 'service::http', 'request', fields)
        jobs = []
        for job in range(request.get('tasks', 3)):
            child = next_id
            next_id += 1
            child_level = 3 if job % 2 == 0 else 4
            visible_child = sampled and child_level <= maximum
            child_fields = dict(request=ordinal, job=job)
            parent = root if visible_root else 0
            if visible_child:
                add('span_start', root, child, parent, child_level, 'service::worker', 'lookup', child_fields)
            jobs.append((job, child, child_level, visible_child, child_fields, parent))
        fields['status'] = 200
        if visible_root:
            add('span_update', root, root, 0, root_level, 'service::http', 'request', fields)
            add('span_end', root, root, 0, root_level, 'service::http', 'request', fields)
        for job, child, child_level, visible_child, child_fields, parent in jobs:
            event_level = 1 if job % 3 == 0 else 3
            if sampled and event_level <= maximum:
                add('event', root, child if visible_child else parent, 0, event_level,
                    'service::db', 'query_complete', dict(request=ordinal, job=job,
                                                        rows=job + 1, cached=job % 2 == 0,
                                                        detail='row\n"雪"'))
            if visible_child:
                add('span_end', root, child, parent, child_level, 'service::worker', 'lookup', child_fields)
    return output


def main():
    rng = random.Random(781329)
    scenarios = [dict(requests=[])]
    for maximum in range(6):
        for period in range(5):
            for capacity in (1, 3, 16):
                requests = [dict(tasks=rng.randrange(9), level=rng.randrange(1, 6),
                                 early_close=bool(rng.randrange(2))) for _ in range(rng.randrange(1, 7))]
                scenarios.append(dict(maximum=maximum, period=period, capacity=capacity, requests=requests))
    scenarios.append(dict(maximum=5, capacity=1, requests=[dict(tasks=100, early_close=True)]))
    result = subprocess.run([str(BINARY), '--json'], input=json.dumps(scenarios),
                            capture_output=True, text=True, timeout=60)
    if result.returncode:
        raise RuntimeError(result.stdout + result.stderr)
    outputs = json.loads(result.stdout)
    assert len(outputs) == len(scenarios)
    total = 0
    for index, (scenario, output) in enumerate(zip(scenarios, outputs)):
        records = output['records']
        expected = expected_records(scenario)
        assert Counter(map(signature, records)) == Counter(map(signature, expected)), (index, scenario, output)
        assert output['accepted'] == len(records)
        assert output['dropped'] == 0 and output['live_spans'] == 0
        starts = {}
        ends = {}
        last_request = -1
        for sequence, record in enumerate(records, 1):
            assert record['sequence'] == sequence
            assert isinstance(record['timestamp_ns'], int) and record['timestamp_ns'] > 0
            assert isinstance(record['elapsed_ns'], int) and record['elapsed_ns'] >= 0
            request = record['fields']['request']
            assert request >= last_request
            last_request = request
            span = record['span_id']
            if record['kind'] == 'span_start':
                assert span not in starts
                if record['parent_id']:
                    assert record['parent_id'] in starts
                starts[span] = record
            elif record['kind'] == 'span_end':
                assert span in starts and span not in ends
                assert record['elapsed_ns'] >= starts[span]['elapsed_ns']
                ends[span] = record
            elif record['kind'] == 'span_update':
                assert span in starts and span not in ends
            elif span:
                assert span in starts
                if starts[span]['target'] == 'service::worker':
                    assert span not in ends
        assert starts.keys() == ends.keys()
        total += len(records)
    print(f'Tracing independent Python oracle: {len(scenarios)} scenarios, {total} records passed')


if __name__ == '__main__':
    main()
