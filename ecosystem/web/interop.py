import concurrent.futures
import http.client
import json
import os
import random
from pathlib import Path
import subprocess
import time
from urllib.parse import urlencode


ROOT = Path(__file__).resolve().parent
BINARY = Path(os.environ.get("WEB_BINARY", str(ROOT.parent / "consumers/web/_artifact/bin/web")))


def main():
    process = subprocess.Popen([str(BINARY), "--serve"], stdin=subprocess.PIPE, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)
    address = process.stdout.readline().strip()
    if not address.startswith("127.0.0.1:"):
        process.kill()
        raise RuntimeError(f"server did not report address: {address!r}")
    host, port = address.rsplit(":", 1)
    port = int(port)

    def request(method, path, body=None, headers=None):
        connection = http.client.HTTPConnection(host, port, timeout=5)
        try:
            connection.request(method, path, body=body, headers=headers or {})
            response = connection.getresponse()
            return response.status, response.getheaders(), response.read()
        finally:
            connection.close()

    try:
        status, headers, body = request("GET", "/api/users/%E4%B8%AD%F0%9F%98%80")
        assert status == 200 and json.loads(body) == {"name": "中😀", "count": 1}
        assert ("X-Web", "goml") in headers
        status, _, body = request("POST", "/api/json", json.dumps({"name": "typed", "count": 41}), {"Content-Type": "application/json"})
        assert status == 201 and json.loads(body)["count"] == 42
        form = urlencode([("label", "0001 +中"), ("count", "-9"), ("enabled", "true"), ("tag", "a"), ("tag", "b")])
        for method, path, body, headers in (("GET", "/api/query?" + form, None, {}), ("POST", "/api/form", form, {"Content-Type": "application/x-www-form-urlencoded"})):
            status, _, output = request(method, path, body, headers)
            assert status == 200 and json.loads(output) == {"label": "0001 +中", "count": -9, "enabled": True, "tag": ["a", "b"]}
        generator = random.Random(82471)
        for _ in range(60):
            label = "".join(generator.choice("abc01 +&=%中😀") for _ in range(generator.randrange(1, 24)))
            count = generator.randrange(-100000, 100001)
            enabled = bool(generator.randrange(2))
            tags = ["".join(generator.choice("a1 +&中") for _ in range(5)) for _ in range(generator.randrange(1, 5))]
            query = urlencode([("label", label), ("count", str(count)), ("enabled", str(enabled).lower())] + [("tag", tag) for tag in tags])
            status, _, payload = request("GET", "/api/query?" + query)
            assert status == 200 and json.loads(payload) == {"label": label, "count": count, "enabled": enabled, "tag": tags}
        assert request("GET", "/files/a/b%20c")[2] == b"a/b c"
        assert request("POST", "/api/users/1")[0] == 405
        assert request("HEAD", "/api/users/1")[2] == b""
        assert request("OPTIONS", "/api/users/1")[0] == 204
        assert request("GET", "/missing")[0] == 404
        assert request("POST", "/api/json", "bad", {"Content-Type": "application/json"})[0] == 422
        assert request("POST", "/api/json", "{}")[0] == 415
        assert request("POST", "/api/json", b"x" * 262145, {"Content-Type": "application/json"})[0] == 413
        status, headers, _ = request("GET", "/headers")
        assert status == 200 and len([v for k, v in headers if k.lower() == "set-cookie"]) == 2
        assert request("GET", "/redirect")[0] == 307
        connection = http.client.HTTPConnection(host, port, timeout=5)
        payload = bytes(range(256)) * 800
        connection.request("POST", "/stream", body=(payload[i:i + 8192] for i in range(0, len(payload), 8192)), encode_chunked=True)
        response = connection.getresponse()
        assert response.status == 200 and response.read() == payload
        connection.close()
        connection = http.client.HTTPConnection(host, port, timeout=5)
        connection.request("GET", "/events")
        response = connection.getresponse()
        assert response.getheader("Content-Type") == "text/event-stream"
        assert response.readline() == b"id: 0\n"
        assert response.readline() == b"data: value 0\n"
        assert b"data: value 2\n\n" in response.read()
        connection.close()
        connection = http.client.HTTPConnection(host, port, timeout=5)
        connection.request("GET", "/forever")
        response = connection.getresponse()
        assert response.readline() == b"data: tick\n"
        response.close()
        connection.close()
        deadline = time.monotonic() + 2
        while json.loads(request("GET", "/active")[2])["count"] != 0:
            if time.monotonic() >= deadline:
                raise AssertionError("disconnected SSE producer remains active")
            time.sleep(0.01)
        with concurrent.futures.ThreadPoolExecutor(max_workers=12) as pool:
            counts = list(pool.map(lambda _: json.loads(request("GET", "/counter")[2])["count"], range(48)))
        assert sorted(counts) == list(range(1, 49))
        assert request("GET", "/slow")[0] == 504
        assert request("GET", "/")[2] == b"web consumer"
    finally:
        process.stdin.close()
        try:
            process.wait(timeout=5)
        except subprocess.TimeoutExpired:
            process.kill()
            process.wait()
        error = process.stderr.read()
        if process.returncode:
            raise RuntimeError(f"consumer exited {process.returncode}: {error}")
    print("web interoperability: Python HTTP client, typed JSON/form/query, routing, Unicode, chunked binary streaming, SSE/disconnect cleanup, concurrency, limits and deadline passed")


if __name__ == "__main__":
    main()
