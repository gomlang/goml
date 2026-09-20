import json
from pathlib import Path
import random
import subprocess


ROOT = Path(__file__).resolve().parents[1]
BINARY = ROOT / "consumers" / "incremental" / "_artifact" / "bin" / "incremental"


class EvaluationError(Exception):
    pass


def reference(case):
    inputs = list(case["inputs"])
    revision = 1
    output = []

    def evaluate(index, active):
        if index in active:
            raise EvaluationError("cycle")
        active = active | {index}
        node = case["nodes"][index]
        kind, a, b, c = (node[key] for key in ("kind", "a", "b", "c"))
        if kind == "input":
            return inputs[a]
        if kind == "constant":
            return a
        if kind == "add":
            return evaluate(a, active) + evaluate(b, active)
        if kind == "subtract":
            return evaluate(a, active) - evaluate(b, active)
        if kind == "parity":
            return evaluate(a, active) % 2
        if kind == "branch":
            return evaluate(b if evaluate(a, active) != 0 else c, active)
        if kind == "nonnegative":
            value = evaluate(a, active)
            if value < 0:
                raise EvaluationError("negative")
            return value
        if kind == "fallback":
            try:
                return evaluate(a, active)
            except EvaluationError:
                return b
        raise AssertionError(kind)

    for action in case["actions"]:
        kind, index, value = (action[key] for key in ("kind", "index", "value"))
        if kind == "set":
            if inputs[index] != value:
                inputs[index] = value
                revision += 1
        elif kind == "query":
            try:
                output.append(evaluate(index, set()))
            except EvaluationError as error:
                output.append(str(error))
    return {"values": output, "revision": revision}


def node(kind, a=0, b=0, c=0):
    return {"kind": kind, "a": a, "b": b, "c": c}


def action(kind, index=0, value=0):
    return {"kind": kind, "index": index, "value": value}


def main():
    rng = random.Random(20260926)
    cases = []
    for _ in range(320):
        inputs = [rng.randrange(-8, 9) for _ in range(5)]
        nodes = [node("input", index) for index in range(len(inputs))]
        for index in range(5, 24):
            kind = rng.choice(["add", "subtract", "parity", "branch", "nonnegative", "fallback"])
            nodes.append(node(kind, rng.randrange(index), rng.randrange(index), rng.randrange(index)))
        actions = []
        for _ in range(100):
            kind = rng.randrange(10)
            if kind < 4:
                actions.append(action("set", rng.randrange(5), rng.randrange(-8, 9)))
            elif kind == 4:
                actions.append(action("clear"))
            else:
                actions.append(action("query", rng.randrange(len(nodes))))
        cases.append({"inputs": inputs, "nodes": nodes, "actions": actions, "capacity": rng.choice([2, 4, 8, 64])})
    for capacity in (1, 2, 8, 64):
        cases.append({
            "inputs": [0, 9],
            "nodes": [node("input", 0), node("input", 1), node("branch", 0, 3, 1), node("add", 2, 1), node("fallback", 2, 100)],
            "actions": [action("query", 2), action("set", 0, 1), action("query", 2), action("query", 4), action("set", 0, 0), action("query", 4), action("query", 2)],
            "capacity": capacity,
        })
    expected = [reference(case) for case in cases]
    process = subprocess.run([str(BINARY), "--json"], input=json.dumps(cases), capture_output=True, text=True, check=True, timeout=90)
    actual = json.loads(process.stdout)
    if len(actual) != len(expected):
        raise AssertionError(f"case count {len(actual)} != {len(expected)}")
    for index, (left, right) in enumerate(zip(actual, expected)):
        if left != right:
            raise AssertionError(f"case {index}: {cases[index]!r}\nGoML: {left!r}\nPython: {right!r}")
    queries = sum(action["kind"] == "query" for case in cases for action in case["actions"])
    print(f"incremental interoperability: {len(cases)} histories, {queries} queries match independent from-scratch evaluation")


if __name__ == "__main__":
    main()
