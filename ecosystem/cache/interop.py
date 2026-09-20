import json
import random
import subprocess
from collections import OrderedDict
from pathlib import Path


ROOT = Path(__file__).resolve().parent
BINARY = ROOT.parent / "consumers/cache/_artifact/bin/cache"


def reference(case):
    values = OrderedDict()
    now = 0
    output = []

    def expire():
        for key, (_, _, created, used) in list(values.items()):
            if ((case["ttl_ms"] and now - created >= case["ttl_ms"])
                    or (case["tti_ms"] and now - used >= case["tti_ms"])):
                del values[key]

    for operation in case["operations"]:
        kind = operation["op"]
        if kind == "advance":
            now += operation["ms"]
            continue
        expire()
        key = operation.get("key")
        if kind == "put":
            values.pop(key, None)
            values[key] = (operation["value"], operation["weight"], now, now)
            while (len(values) > case["max_entries"]
                   or sum(item[1] for item in values.values()) > case["max_weight"]):
                values.popitem(last=False)
        elif kind in ("get", "peek", "remove"):
            found = values.get(key)
            output.append(None if found is None else found[0])
            if found is not None and kind == "get":
                value, weight, created, _ = found
                values[key] = (value, weight, created, now)
                values.move_to_end(key)
            elif kind == "remove":
                values.pop(key, None)
        elif kind == "clear":
            values.clear()
        else:
            raise AssertionError(kind)
    return output


def cases():
    rng = random.Random(20260920)
    result = []
    for index in range(160):
        max_entries = rng.randrange(1, 17)
        max_weight = rng.randrange(1, 49)
        ttl = rng.choice([0, 1, 3, 10, 50])
        tti = rng.choice([0, 1, 5, 20])
        operations = []
        for step in range(240):
            kind = rng.choices(["put", "get", "peek", "remove", "advance", "clear"],
                               [35, 25, 10, 8, 20, 2])[0]
            operation = {"op": kind}
            if kind in ("put", "get", "peek", "remove"):
                operation["key"] = rng.randrange(24)
            if kind == "put":
                operation.update(value=f"{index}:{step}:" + rng.choice(["", "中文", "😀"]),
                                 weight=rng.randrange(1, min(max_weight, 12) + 1))
            if kind == "advance":
                operation["ms"] = rng.choice([0, 1, 2, 5, 10, 50])
            operations.append(operation)
        operations.extend({"op": "peek", "key": key} for key in range(24))
        result.append(dict(max_entries=max_entries, max_weight=max_weight,
                           ttl_ms=ttl, tti_ms=tti, operations=operations))
    return result


def main():
    inputs = cases()
    expected = [reference(case) for case in inputs]
    result = subprocess.run([str(BINARY), "--json"], input=json.dumps(inputs),
                            text=True, capture_output=True, timeout=120)
    if result.returncode:
        raise RuntimeError(result.stdout + result.stderr)
    actual = json.loads(result.stdout)
    if len(actual) != len(expected):
        raise AssertionError((len(actual), len(expected)))
    for index, (got, wanted) in enumerate(zip(actual, expected)):
        if got != wanted:
            raise AssertionError(f"case {index}: {inputs[index]}\nGoML: {got!r}\nPython: {wanted!r}")
    queries = sum(len(values) for values in expected)
    operations = sum(len(case["operations"]) for case in inputs)
    print(f"cache: {len(inputs)} Python LRU/TTL/TTI histories, {operations} operations and {queries} queries passed")


if __name__ == "__main__":
    main()
