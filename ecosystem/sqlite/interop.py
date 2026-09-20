import argparse
from collections import Counter
import json
import math
from pathlib import Path
import random
import sqlite3
import subprocess
import tempfile


ROOT = Path(__file__).resolve().parents[1]


class TextBytes(bytes):
    pass


def encoded(value):
    if value is None:
        return {"type": "null", "value": None}
    if isinstance(value, int):
        return {"type": "integer", "value": str(value)}
    if isinstance(value, float):
        return {"type": "real", "value": str(value)}
    if isinstance(value, TextBytes):
        try:
            return {"type": "text", "value": value.decode("utf-8")}
        except UnicodeDecodeError:
            return {"type": "textbytes", "value": list(value)}
    if isinstance(value, bytes):
        return {"type": "blob", "value": list(value)}
    return {"type": "text", "value": value}


def decoded(value):
    kind, raw = value["type"], value["value"]
    if kind == "integer": return int(raw)
    if kind == "real": return float(raw)
    if kind == "blob": return bytes(raw)
    if kind == "textbytes": raise ValueError("raw TEXT input is tested through CAST in the independent reference")
    return raw


def reference(case):
    connection = sqlite3.connect(":memory:", isolation_level=None)
    connection.text_factory = TextBytes
    connection.execute("PRAGMA foreign_keys=ON")
    results, statements, transactions = [], [], []
    try:
        for command in case["commands"]:
            try:
                op = command["op"]
                if op == "prepare":
                    statements.append(command["sql"])
                    results.append(len(statements) - 1)
                    continue
                if op == "begin":
                    connection.execute("BEGIN " + command.get("mode", "deferred").upper())
                    transactions.append(None)
                    results.append(None)
                    continue
                if op == "savepoint":
                    name = "reference_savepoint_" + str(len(transactions))
                    connection.execute("SAVEPOINT " + name)
                    transactions.append(name)
                    results.append(None)
                    continue
                if op in ("commit", "rollback"):
                    name = transactions[-1]
                    if name is None:
                        connection.execute("COMMIT" if op == "commit" else "ROLLBACK")
                    elif op == "commit":
                        connection.execute("RELEASE SAVEPOINT " + name)
                    else:
                        connection.execute("ROLLBACK TO SAVEPOINT " + name)
                        connection.execute("RELEASE SAVEPOINT " + name)
                    transactions.pop()
                    results.append(None)
                    continue
                if op == "stmt_close":
                    results.append(None)
                    continue
                values = [decoded(value) for value in command.get("params", [])]
                if "names" in command:
                    values = dict(zip(command["names"], values))
                sql = statements[command["id"]] if op.startswith("stmt_") else command["sql"]
                cursor = connection.execute(sql, values)
                if op in ("query", "stmt_query"):
                    rows = cursor.fetchall()
                    if len(rows) > command.get("limit", 10000):
                        results.append({"error": "Limit", "code": 0})
                    else:
                        results.append({"columns": [item[0] for item in cursor.description], "rows": [[encoded(value) for value in row] for row in rows]})
                else:
                    changes, rowid = connection.execute("SELECT changes(),last_insert_rowid()").fetchone()
                    results.append({"changes": str(changes), "last_id": str(rowid)})
            except sqlite3.Error as error:
                code = getattr(error, "sqlite_errorcode", 0)
                kind = "Constraint" if code & 255 == 19 else "Busy" if code & 255 in (5, 6) else "Sqlite"
                results.append({"error": kind, "code": code})
    finally:
        connection.close()
    return results


def compare(actual, expected):
    if isinstance(expected, dict):
        if not isinstance(actual, dict):
            raise AssertionError(f"expected object, got {actual}")
        if "error" in expected:
            if actual.get("error", "").split("::")[-1] != expected["error"] or actual.get("code") != expected["code"]:
                raise AssertionError(f"error differs: {actual} != {expected}")
            return
        if expected.get("type") == "real":
            a, b = float(actual["value"]), float(expected["value"])
            if actual.get("type") != "real" or not (a == b or math.isnan(a) and math.isnan(b) or math.isclose(a, b, rel_tol=1e-12, abs_tol=1e-12)):
                raise AssertionError(f"real differs: {actual} != {expected}")
            return
        if set(actual) != set(expected):
            raise AssertionError(f"object fields differ: {actual} != {expected}")
        for key in expected:
            compare(actual[key], expected[key])
    elif isinstance(expected, list):
        if not isinstance(actual, list) or len(actual) != len(expected):
            raise AssertionError(f"list lengths differ: {actual} != {expected}")
        for a, b in zip(actual, expected):
            compare(a, b)
    elif actual != expected:
        raise AssertionError(f"{actual!r} != {expected!r}")


def main():
    arguments = argparse.ArgumentParser()
    arguments.add_argument("--consumer", type=Path, default=ROOT / "consumers/sqlite/_artifact/bin/sqlite")
    args = arguments.parse_args()
    rng = random.Random(20260920)
    cases, groups = [], []

    def add(group, commands):
        cases.append({"commands": commands})
        groups.append(group)

    def command(op, sql, values=(), **extra):
        return {"op": op, "sql": sql, "params": [encoded(value) for value in values], **extra}

    special = [None, -2**63, 2**63-1, -2**53-1, 2**53+1, 0, 1, -1, 0.0, -0.0, 1.25, -1e200, 1e-250, math.inf, -math.inf, math.nan, "", "中😀é", "NUL\0tail", "x'); DROP TABLE t; --", b"", bytes(range(256))]
    for value in special:
        add("storage-values", [command("query", "SELECT ? AS value, typeof(?) AS kind", [value, value])])
    for _ in range(80):
        commands = [command("exec", "CREATE TABLE records(id INTEGER PRIMARY KEY, category INTEGER, score REAL, name TEXT, payload BLOB, optional INTEGER)"), {"op": "prepare", "sql": "INSERT INTO records VALUES(:id,:category,:score,:name,:payload,:optional)"}]
        count = rng.randrange(1, 20)
        for index in range(count):
            values = [index+1, rng.randrange(4), rng.randrange(-1000, 1000) / 16., rng.choice(["", "中文😀", "quote'\"", "line\nend", "nul\0end"])+str(index), rng.randbytes(rng.randrange(12)), None if rng.randrange(2) else rng.randrange(-100, 100)]
            names = ["id", "category", "score", "name", "payload", "optional"]
            pairs = list(zip(names, values)); rng.shuffle(pairs)
            commands.append({"op": "stmt_exec", "id": 0, "names": [p[0] for p in pairs], "params": [encoded(p[1]) for p in pairs]})
        commands += [
            command("query", "SELECT * FROM records ORDER BY id"),
            command("query", "SELECT category,count(*) AS count,sum(score) AS total,avg(score) AS average,min(score) AS minimum,max(score) AS maximum FROM records GROUP BY category ORDER BY category"),
            command("query", "SELECT id, row_number() OVER(PARTITION BY category ORDER BY score,id) AS ordinal,lag(id) OVER(ORDER BY id) AS previous FROM records ORDER BY id"),
            command("query", "SELECT a.id,b.id,a.score-b.score FROM records a JOIN records b ON a.category=b.category WHERE a.id<b.id ORDER BY a.id,b.id"),
            command("query", "SELECT id,length(payload),hex(payload),optional IS NULL,coalesce(optional,-1) FROM records WHERE id>? ORDER BY id", [count//2]),
            {"op": "begin", "mode": rng.choice(["deferred", "immediate", "exclusive"])},
            command("exec", "UPDATE records SET score=score+10 WHERE category=?", [rng.randrange(4)]),
            {"op": "savepoint"}, command("exec", "DELETE FROM records WHERE id%2=0"),
            {"op": rng.choice(["commit", "rollback"])},
            command("query", "SELECT id,score FROM records ORDER BY id"),
            {"op": rng.choice(["commit", "rollback"])},
            command("query", "SELECT * FROM records ORDER BY id"),
            {"op": "prepare", "sql": "SELECT id,name FROM records WHERE id>=? ORDER BY id DESC"},
            {"op": "stmt_query", "id": 1, "params": [encoded(rng.randrange(count+1))]},
            {"op": "stmt_query", "id": 1, "params": [encoded(count+10)]},
            {"op": "stmt_close", "id": 0}, {"op": "stmt_close", "id": 1},
        ]
        add("queries-transactions-prepared", commands)
    for _ in range(100):
        n = rng.randrange(1, 100)
        add("recursive-and-json", [command("query", "WITH RECURSIVE n(x) AS (VALUES(1) UNION ALL SELECT x+1 FROM n WHERE x<?) SELECT count(*),sum(x),sum(x*x) FROM n", [n]), command("query", "SELECT json_extract(?, '$.value'),json_array_length(?, '$.items'),json_valid(?)", [json.dumps({"value": n}), json.dumps({"items": list(range(n%7))}), "{bad"])])
    add("constraints-and-errors", [
        command("exec", "CREATE TABLE parent(id INTEGER PRIMARY KEY, name TEXT NOT NULL UNIQUE CHECK(length(name)>0))"),
        command("exec", "CREATE TABLE child(id INTEGER REFERENCES parent(id) DEFERRABLE INITIALLY DEFERRED)"),
        command("exec", "INSERT INTO parent VALUES(?,?)", [1, "first"]),
        command("exec", "INSERT INTO parent VALUES(?,?)", [1, "second"]),
        command("exec", "INSERT INTO parent VALUES(?,?)", [2, "first"]),
        command("exec", "INSERT INTO parent VALUES(?,?)", [3, None]),
        command("exec", "INSERT INTO parent VALUES(?,?)", [4, ""]),
        {"op": "begin"}, command("exec", "INSERT INTO child VALUES(99)"), {"op": "commit"}, {"op": "rollback"},
        command("query", "SELECT * FROM child"), command("query", "SELECT * FROM parent"),
        command("query", "SELECT * FROM missing_table"), command("query", "SELECT 1/0, CAST(x'fffe' AS TEXT), typeof(x''), length(x'')"),
        command("query", "SELECT 1 UNION ALL SELECT 2", limit=1), command("query", "SELECT 1 WHERE 0", limit=0),
    ])
    for _ in range(40):
        data = rng.randbytes(rng.randrange(1, 50))
        add("invalid-text", [command("query", "SELECT CAST(? AS TEXT),hex(CAST(? AS TEXT)),typeof(CAST(? AS TEXT))", [data, data, data])])

    def run(requests):
        result = subprocess.run([str(args.consumer.resolve()), "oracle"], input=json.dumps(requests, allow_nan=False), text=True, capture_output=True, timeout=90)
        if result.returncode:
            raise RuntimeError(f"consumer failed: {result.stderr[-5000:]}")
        values = json.loads(result.stdout)
        if len(values) != len(requests):
            raise AssertionError("case count differs")
        return values

    actual = run(cases)
    counts = Counter()
    for index, (case, group, outputs) in enumerate(zip(cases, groups, actual)):
        expected = reference(case)
        if len(outputs) != len(expected):
            raise AssertionError(f"case {index}: reply count differs")
        for command_index, (result, reference_value) in enumerate(zip(outputs, expected)):
            try:
                compare(result, reference_value)
            except Exception as error:
                raise AssertionError(f"case {index}, command {command_index}, {group}: {json.dumps(case['commands'][command_index])}\n{error}") from error
        counts[group] += len(expected)

    with tempfile.TemporaryDirectory(prefix="goml-sqlite-reference-") as directory:
        path = Path(directory) / "shared.db"
        created = run([{"dsn": str(path), "commands": [command("exec", "CREATE TABLE shared(id INTEGER PRIMARY KEY, value BLOB, name TEXT)"), command("exec", "INSERT INTO shared VALUES(?,?,?)", [1, bytes([0, 255, 128]), "GoML 中😀"]), {"op": "begin"}, command("exec", "INSERT INTO shared VALUES(2,x'01','rollback on close')"), {"op": "close"}]}])[0]
        if any(isinstance(value, dict) and "error" in value for value in created):
            raise AssertionError(created)
        with sqlite3.connect(path) as connection:
            if connection.execute("SELECT * FROM shared").fetchall() != [(1, bytes([0, 255, 128]), "GoML 中😀")]:
                raise AssertionError("Python could not read GoML data or close did not roll back")
            connection.execute("INSERT INTO shared VALUES(?,?,?)", (3, b"Python\0bytes", "Python é"))
        reopened = run([{"dsn": "file:" + str(path) + "?mode=ro", "commands": [command("query", "SELECT * FROM shared ORDER BY id"), command("exec", "DELETE FROM shared")]}])[0]
        compare(reopened[0], {"columns": ["id", "value", "name"], "rows": [[encoded(1), encoded(bytes([0,255,128])), encoded("GoML 中😀")], [encoded(3), encoded(b"Python\0bytes"), encoded("Python é")]]})
        compare(reopened[1], {"error": "Sqlite", "code": 8})
        counts["file-interoperability"] += 7
    version = run([{"commands": [command("query", "SELECT sqlite_version()")]}])[0][0]["rows"][0][0]["value"]
    print(f"SQLite reference: Python {sqlite3.sqlite_version}, GoML adapter {version}; {len(cases)} sequences, {sum(counts.values())} checked operations passed")
    print(json.dumps(dict(sorted(counts.items())), sort_keys=True))


if __name__ == "__main__":
    main()
