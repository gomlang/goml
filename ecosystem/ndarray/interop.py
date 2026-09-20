import argparse
from collections import Counter
import hashlib
import io
import json
import math
from pathlib import Path
import random
import subprocess
import sys
import urllib.request
import warnings
import zipfile


ROOT = Path(__file__).resolve().parents[1]
WHEEL = "numpy-2.2.6-cp312-cp312-manylinux_2_17_x86_64.manylinux2014_x86_64.whl"
URL = "https://files.pythonhosted.org/packages/8c/3d/1e1db36cfd41f895d266b103df00ca5b3cbe965184df824dec5c08c6b803/" + WHEEL
DIGEST = "fd83c01228a688733f1ded5201c678f0c53ecc1006ffbc404db9f7a899ac6249"


def reference():
    if sys.version_info[:2] != (3, 12):
        raise RuntimeError("the pinned NumPy reference wheel requires CPython 3.12 on Linux amd64")
    directory = ROOT / "_artifact/reference"
    directory.mkdir(parents=True, exist_ok=True)
    cache = directory / WHEEL
    if not cache.exists():
        data = urllib.request.urlopen(URL, timeout=60).read()
        if hashlib.sha256(data).hexdigest() != DIGEST:
            raise RuntimeError("NumPy reference download checksum mismatch")
        cache.write_bytes(data)
    data = cache.read_bytes()
    if hashlib.sha256(data).hexdigest() != DIGEST:
        raise RuntimeError("NumPy reference cache checksum mismatch")
    target = directory / "numpy-2.2.6-reference"
    with zipfile.ZipFile(io.BytesIO(data)) as archive:
        for name in archive.namelist():
            relative = Path(name)
            if relative.is_absolute() or ".." in relative.parts:
                raise RuntimeError("invalid reference wheel entry")
            if name.endswith("/"):
                continue
            destination = target / relative
            destination.parent.mkdir(parents=True, exist_ok=True)
            destination.write_bytes(archive.read(name))
    sys.path.insert(0, str(target))
    import numpy as np
    if np.__version__ != "2.2.6" or target not in Path(np.__file__).parents:
        raise RuntimeError("wrong NumPy reference imported")
    np.seterr(all="ignore")
    warnings.filterwarnings("ignore", message="Degrees of freedom <= 0 for slice", category=RuntimeWarning)
    return np


def main():
    arguments = argparse.ArgumentParser()
    arguments.add_argument("--consumer", type=Path, default=ROOT / "consumers/ndarray/_artifact/bin/ndarray")
    args = arguments.parse_args()
    np = reference()
    rng = random.Random(20260920)
    generator = np.random.default_rng(20260920)
    cases = []

    def number(value):
        value = float(value)
        return "nan" if math.isnan(value) else "inf" if value == math.inf else "-inf" if value == -math.inf else value

    def encode(value, views=None):
        value = np.asarray(value, dtype=np.float64)
        return {"shape": list(value.shape), "data": [number(x) for x in value.flat], **({"views": views} if views else {})}

    def add(group, request, expected):
        cases.append((group, request, expected))

    def arr(shape):
        return generator.normal(size=shape)

    def view_case(a, steps):
        value = a
        for step in steps:
            op = step["op"]
            if op == "slice":
                index = [slice(None)] * value.ndim
                index[step.get("axis", 0)] = slice(step.get("start"), step.get("stop"), step.get("step", 1))
                value = value[tuple(index)]
            elif op in ("reshape", "reshape_view"):
                value = value.reshape(step["shape"])
            elif op == "transpose":
                value = value.transpose(step["axes"])
            elif op == "broadcast":
                value = np.broadcast_to(value, step["shape"])
            elif op == "index":
                index = [slice(None)] * value.ndim
                index[step.get("axis", 0)] = step["index"]
                value = value[tuple(index)]
            elif op == "insert":
                value = np.expand_dims(value, step["axis"])
            elif op == "squeeze":
                value = np.squeeze(value)
            elif op == "diagonal":
                value = np.diagonal(value, step.get("offset", 0), step["first"], step["second"])
            elif op == "take":
                value = np.take(value, step["indices"], axis=step["axis"])
        add("views", {"op": "view", "a": encode(a, steps)}, value)
        return value

    for shape in [(), (0,), (1,), (7,), (2, 0, 3), (3, 4), (2, 3, 4), (1, 2, 1, 3)]:
        a = arr(shape)
        view_case(a, [])
        if a.ndim:
            for _ in range(70):
                axis = rng.randrange(-a.ndim, a.ndim)
                step = rng.choice([-9, -3, -2, -1, 1, 2, 5, 2**63-1, -2**63])
                start, stop = (rng.choice([None, -20, -4, -1, 0, 1, 2, 20]) for _ in range(2))
                view_case(a, [{"op": "slice", "axis": axis, "start": start, "stop": stop, "step": step}])
            permutation = list(range(a.ndim)); rng.shuffle(permutation)
            view_case(a, [{"op": "transpose", "axes": permutation}, {"op": "reshape", "shape": [-1]}])
            for axis in range(a.ndim):
                if a.shape[axis]:
                    view_case(a, [{"op": "index", "axis": axis, "index": -1}])
                    view_case(a, [{"op": "take", "axis": axis, "indices": [-1, 0, -1]}])
        view_case(a, [{"op": "insert", "axis": -1}, {"op": "squeeze"}])
        if a.ndim >= 2:
            for offset in [-6, -1, 0, 2, 9]:
                view_case(a, [{"op": "diagonal", "offset": offset, "first": -1, "second": 0}])
    view_case(np.array(3.5), [{"op": "broadcast", "shape": [2, 0, 3]}])
    view_case(arr((3,)), [{"op": "broadcast", "shape": [2, 1, 3]}])

    for _ in range(220):
        shape = tuple(rng.randrange(1, 5) for _ in range(rng.randrange(5)))
        left = tuple(d if rng.randrange(2) else 1 for d in shape)
        right = tuple(d if rng.randrange(2) else 1 for d in shape)
        left = left[rng.randrange(len(left) + 1):]
        right = right[rng.randrange(len(right) + 1):]
        a, b = arr(left), arr(right) + 0.25
        for op, function in [("add", np.add), ("subtract", np.subtract), ("multiply", np.multiply), ("divide", np.divide)]:
            add("broadcast-arithmetic", {"op": op, "a": encode(a), "b": encode(b)}, function(a, b))
    for count in range(34):
        a, b = arr((count,)), arr((count,))
        for scalar in [False, True]:
            add("simd-tails", {"op": "divide", "a": encode(a), "b": encode(b), "scalar": scalar}, a / b)
    special = np.array([0.0, -0.0, np.inf, -np.inf, np.nan, 1.0, -2.0, 1e-300, 1e300])
    for op, function in [("add", np.add), ("multiply", np.multiply), ("divide", np.divide)]:
        for scalar in [False, True]:
            add("ieee", {"op": op, "a": encode(special), "b": encode(special[::-1]), "scalar": scalar}, function(special, special[::-1]))

    for shape in [(), (0,), (2, 0, 3), (7,), (2, 3), (2, 3, 4), (1, 2, 1, 3)]:
        a = arr(shape)
        for _ in range(18):
            axes = rng.sample(range(a.ndim), rng.randrange(a.ndim + 1))
            count = math.prod(a.shape[x] for x in axes)
            keep = bool(rng.randrange(2))
            for op, function in [("sum", np.sum), ("product", np.prod), ("mean", np.mean), ("variance", np.var), ("min", np.min), ("max", np.max)]:
                ddof = rng.randrange(2) if op == "variance" else 0
                output_size = math.prod(a.shape[i] for i in range(a.ndim) if i not in axes)
                request = {"op": op, "a": encode(a), "axes": axes, "keep": keep, "ddof": ddof}
                if output_size and op not in ("sum", "product") and count <= ddof:
                    add("empty-reductions", request, "Empty")
                elif not output_size and op in ("min", "max") and not count:
                    add("empty-reductions", request, np.empty(tuple(1 if keep and i in axes else d for i, d in enumerate(shape) if keep or i not in axes)))
                else:
                    kwargs = {"ddof": ddof} if op == "variance" else {}
                    add("reductions", request, function(a, axis=tuple(axes), keepdims=keep, **kwargs))
    for _ in range(120):
        m, k, n = (rng.randrange(0, 7) for _ in range(3))
        av, bv = bool(rng.randrange(2)), bool(rng.randrange(2))
        batch = rng.choice([(), (2,), (2, 3), (0,), (2, 1)])
        ashape = (k,) if av else batch + (m, k)
        bshape = (k,) if bv else tuple(1 if rng.randrange(2) else d for d in batch) + (k, n)
        a, b = arr(ashape), arr(bshape)
        for scalar in [False, True]:
            add("batched-matmul", {"op": "matmul", "a": encode(a), "b": encode(b), "scalar": scalar}, a @ b)
    for _ in range(30):
        a, b = arr((4, 5)), arr((4, 3))
        add("strided-matmul", {"op": "matmul", "a": encode(a, [{"op": "transpose", "axes": [1, 0]}]), "b": encode(b, [{"op": "slice", "axis": 0, "step": -1}])}, a.T @ b[::-1])

    for _ in range(35):
        n = rng.randrange(1, 8)
        a = arr((n, n)) + (n + 1) * np.eye(n)
        rhs = arr((n, rng.randrange(1, 4))) if rng.randrange(2) else arr((n,))
        for op, expected in [("solve", np.linalg.solve(a, rhs)), ("inverse", np.linalg.inv(a)), ("det", np.linalg.det(a)), ("slogdet", np.linalg.slogdet(a))]:
            add("lu-solvers", {"op": op, "a": encode(a), "b": encode(rhs)}, np.array(expected))
        add("lu-reconstruction", {"op": "lu", "a": encode(a)}, ("lu", a))
        spd = a.T @ a + np.eye(n)
        add("cholesky", {"op": "cholesky", "a": encode(spd)}, np.linalg.cholesky(spd))
        add("cholesky", {"op": "cholesky_solve", "a": encode(spd), "b": encode(rhs)}, np.linalg.solve(spd, rhs))
        tall = arr((n + rng.randrange(1, 6), n))
        rhs = arr((len(tall), rng.randrange(1, 4))) if rng.randrange(2) else arr((len(tall),))
        add("qr-reconstruction", {"op": "qr", "a": encode(tall)}, ("qr", tall))
        add("least-squares", {"op": "least_squares", "a": encode(tall), "b": encode(rhs)}, np.linalg.lstsq(tall, rhs, rcond=None)[0])
    for scale in [1e-150, 1e150]:
        a = np.array([[4.0, 2.0], [1.0, 3.0], [2.0, 1.0]]) * scale
        add("scaled-qr", {"op": "qr", "a": encode(a)}, ("qr", a))
        add("scaled-qr", {"op": "least_squares", "a": encode(a), "b": encode(np.array([1.0, 2.0, 3.0]) * scale)}, np.linalg.lstsq(a, np.array([1.0, 2.0, 3.0]) * scale, rcond=None)[0])
    for length in [0, 1, 2, 3, 17]:
        for endpoint in [False, True]:
            add("linspace", {"op": "linspace", "start": -3.0, "stop": 9.0, "count": length, "endpoint": endpoint}, np.linspace(-3.0, 9.0, length, endpoint=endpoint))
    for shape in [(0,), (3,), (2, 3)]:
        a = arr(shape)
        for op, function in [("sqrt", np.sqrt), ("exp", np.exp), ("log", np.log), ("abs", np.abs)]:
            add("unary", {"op": op, "a": encode(a)}, function(a))
        add("norm", {"op": "norm", "a": encode(a)}, np.linalg.norm(a))
        add("clip", {"op": "clip", "a": encode(a), "lower": -0.5, "upper": 0.8}, np.clip(a, -0.5, 0.8))
        for axis in range(a.ndim):
            add("combine", {"op": "concat", "a": encode(a), "b": encode(a), "axis": axis}, np.concatenate([a, a], axis=axis))
        for axis in range(-a.ndim-1, a.ndim+1):
            add("combine", {"op": "stack", "a": encode(a), "b": encode(a), "axis": axis}, np.stack([a, a], axis=axis))
    a = np.arange(12.).reshape(3, 4)
    destination = [{"op": "slice", "axis": 1, "start": 1}]
    source = [{"op": "slice", "axis": 1, "stop": -1}]
    expected = a.copy(); expected[:, 1:] = a[:, :-1].copy()
    add("overlap", {"op": "assign", "a": encode(a), "destination": destination, "source": source}, expected)
    add("overlap", {"op": "assign", "a": encode(a), "destination": [], "source": [{"op": "slice", "axis": 1, "step": -1}]}, a[:, ::-1])
    add("invalid", {"op": "view", "a": {"shape": [2, 3], "data": [1, 2]}}, "Shape")
    add("invalid", {"op": "view", "a": {"shape": [-1], "data": []}}, "Shape")
    add("invalid", {"op": "view", "a": {"shape": [2**62, 4], "data": []}}, "Limit")
    for step, kind in [({"op": "slice", "step": 0}, "Index"), ({"op": "transpose", "axes": [0, 0]}, "Axis"), ({"op": "index", "axis": 2, "index": 0}, "Axis"), ({"op": "index", "axis": 0, "index": 3}, "Index"), ({"op": "reshape", "shape": [-1, -1]}, "Shape"), ({"op": "broadcast", "shape": [2, 4]}, "Shape")]:
        add("invalid", {"op": "view", "a": encode(a, [step])}, kind)
    add("invalid", {"op": "view", "a": encode(a, [{"op": "transpose", "axes": [1, 0]}, {"op": "reshape_view", "shape": [12]}])}, "NonContiguous")
    add("invalid", {"op": "assign", "a": encode(np.array(3.0)), "destination": [{"op": "broadcast", "shape": [3]}], "source": []}, "ReadOnly")
    add("invalid", {"op": "sum", "a": encode(a), "axes": [0, 0]}, "Axis")
    add("invalid", {"op": "matmul", "a": encode(a), "b": encode(a)}, "Shape")
    add("invalid", {"op": "matmul", "a": encode(np.array(1.0)), "b": encode(a)}, "Shape")
    singular = np.array([[1.0, 2.0], [2.0, 4.0]])
    for op in ["lu", "inverse", "qr"]:
        add("invalid", {"op": op, "a": encode(singular)}, "Singular")
    add("singular-det", {"op": "det", "a": encode(singular)}, np.array(0.0))
    add("singular-det", {"op": "slogdet", "a": encode(singular)}, np.array([0.0, -np.inf]))
    add("invalid", {"op": "lu", "a": encode(np.array([[np.nan]]))}, "NonFinite")
    add("invalid", {"op": "cholesky", "a": encode(np.array([[1., 2.], [2., 1.]]))}, "NotPositiveDefinite")

    requests = [request for _, request, _ in cases]
    process = subprocess.run([str(args.consumer.resolve()), "oracle"], input=json.dumps(requests, allow_nan=False), text=True, capture_output=True, timeout=90)
    if process.returncode:
        raise RuntimeError(f"consumer failed ({process.returncode}): {process.stderr[-5000:]}")
    output = json.loads(process.stdout)
    if len(output) != len(cases):
        raise AssertionError(f"reply count differs: {len(output)} != {len(cases)}")

    def decode(value):
        if "error" in value:
            raise AssertionError(value)
        return np.array([float(x) for x in value["data"]], dtype=np.float64).reshape(value["shape"])

    def compare(actual, expected):
        expected = np.asarray(expected)
        if actual.shape != expected.shape:
            raise AssertionError(f"shape {actual.shape} != {expected.shape}")
        np.testing.assert_allclose(actual, expected, rtol=2e-10, atol=2e-11, equal_nan=True)

    counts = Counter()
    for index, ((group, request, expected), result) in enumerate(zip(cases, output)):
        try:
            if isinstance(expected, str):
                if result.get("error", "").split("::")[-1] != expected:
                    raise AssertionError(f"expected {expected}, got {result}")
            elif isinstance(expected, tuple):
                kind, a = expected
                if kind == "qr":
                    q, r = decode(result["q"]), decode(result["r"])
                    compare(q.T @ q, np.eye(a.shape[1]))
                    scale = np.max(np.abs(a))
                    compare((q @ r) / scale, a / scale)
                    compare(np.tril(r, -1), np.zeros_like(r))
                else:
                    lower, upper = decode(result["lower"]), decode(result["upper"])
                    compare(lower @ upper, a[result["permutation"]])
                    compare(np.diag(lower), np.ones(a.shape[0]))
                    compare(np.triu(lower, 1), np.zeros_like(lower))
                    compare(np.tril(upper, -1), np.zeros_like(upper))
            else:
                compare(decode(result), expected)
        except Exception as error:
            raise AssertionError(f"case {index} ({group}) failed: {json.dumps(request)}\n{error}") from error
        counts[group] += 1
    print(f"NumPy {np.__version__}: {len(cases)} ndarray reference cases passed")
    print(json.dumps(dict(sorted(counts.items())), sort_keys=True))


if __name__ == "__main__":
    main()
