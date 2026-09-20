import argparse
import json
import os
from pathlib import Path
import random
import re
import select
import subprocess
import time


class Client:
    def __init__(self, executable):
        self.process = subprocess.Popen(
            [str(executable), "--stdio"], stdin=subprocess.PIPE,
            stdout=subprocess.PIPE, stderr=subprocess.PIPE,
        )
        self.next_id = 1
        self.responses = 0

    def send_body(self, body, fragmented=False):
        data = body.encode("utf-8")
        wire = f"Content-Length: {len(data)}\r\n\r\n".encode("ascii") + data
        width = 1 if fragmented else len(wire)
        for offset in range(0, len(wire), width):
            self.process.stdin.write(wire[offset:offset + width])
            self.process.stdin.flush()

    def send(self, value, fragmented=False):
        self.send_body(json.dumps(value, ensure_ascii=False), fragmented)

    def notification(self, method, params=None):
        message = {"jsonrpc": "2.0", "method": method}
        if params is not None:
            message["params"] = params
        self.send(message)

    def read_exact(self, count):
        result = bytearray()
        deadline = time.monotonic() + 10
        while len(result) < count:
            remaining = deadline - time.monotonic()
            ready, _, _ = select.select([self.process.stdout], [], [], max(remaining, 0))
            if not ready:
                raise AssertionError("LSP reply timed out")
            data = os.read(self.process.stdout.fileno(), count - len(result))
            if not data:
                raise AssertionError("LSP output ended before reply")
            result.extend(data)
        return bytes(result)

    def read(self):
        header = bytearray()
        while not header.endswith(b"\r\n\r\n"):
            header.extend(self.read_exact(1))
            assert len(header) <= 8192
        fields = dict(line.split(b":", 1) for line in bytes(header[:-4]).split(b"\r\n"))
        length = int(fields[b"Content-Length"])
        assert 0 <= length <= 16777216
        reply = json.loads(self.read_exact(length).decode("utf-8"))
        assert reply["jsonrpc"] == "2.0"
        assert ("result" in reply) != ("error" in reply)
        self.responses += 1
        return reply

    def request(self, method, params=None, error=None, fragmented=False):
        identity = self.next_id
        self.next_id += 1
        message = {"jsonrpc": "2.0", "id": identity, "method": method}
        if params is not None:
            message["params"] = params
        self.send(message, fragmented)
        reply = self.read()
        assert reply["id"] == identity, reply
        if error is not None:
            assert reply["error"]["code"] == error, reply
            return reply["error"]
        assert "result" in reply, reply
        return reply["result"]

    def close(self):
        if self.process.poll() is None:
            self.process.kill()
        self.process.communicate(timeout=10)


def utf16_length(value):
    return len(value.encode("utf-16-le")) // 2


def offset(value, line, character):
    starts = [0] + [match.end() for match in re.finditer(r"\r\n|\r|\n", value)]
    content = re.split(r"\r\n|\r|\n", value)[line]
    units = content.encode("utf-16-le")[:character * 2]
    prefix = units.decode("utf-16-le")
    return starts[line] + len(prefix)


def edit(value, change):
    if "range" not in change:
        return change["text"]
    span = change["range"]
    start = offset(value, **span["start"])
    end = offset(value, **span["end"])
    return value[:start] + change["text"] + value[end:]


def check_server(executable):
    client = Client(executable)
    try:
        client.send_body("{")
        assert client.read()["error"]["code"] == -32700
        client.send_body('[{"jsonrpc":"2.0","id":1,"method":"x"}]')
        assert client.read()["error"]["code"] == -32600
        client.request("test/echo", {}, error=-32002)
        result = client.request("initialize", {"capabilities": {}}, fragmented=True)
        assert result["capabilities"]["positionEncoding"] == "utf-16"
        assert result["capabilities"]["textDocumentSync"] == 2
        client.notification("initialized", {})
        assert client.request("test/add", {"left": 20, "right": 22}) == {"value": 42}
        client.request("test/add", {"left": "bad", "right": 22}, error=-32602)
        client.request("missing", {}, error=-32601)
        client.notification("$/cancelRequest", {"id": 9999})
        assert client.request("test/echo", {"text": "你好😀"}, fragmented=True) == {"text": "你好😀"}
        rng = random.Random(7331)
        position_cases = 0
        for sample in range(30):
            text = "".join(rng.choice(["a", "中", "😀", "é", "e\u0301", "\r", "\n", "\r\n"]) for _ in range(12))
            uri = f"file:///sample-{sample}.gom"
            client.notification("textDocument/didOpen", {"textDocument": {
                "uri": uri, "languageId": "goml", "version": 1, "text": text,
            }})
            lines = re.split(r"\r\n|\r|\n", text)
            for line, content in enumerate(lines):
                for index in range(len(content) + 1):
                    character = utf16_length(content[:index])
                    expected = len(text[:offset(text, line, character)].encode("utf-8"))
                    actual = client.request("test/offset", {
                        "textDocument": {"uri": uri}, "position": {"line": line, "character": character},
                    })
                    assert actual == expected, (text, line, character, actual, expected)
                    position_cases += 1
                actual = client.request("test/offset", {
                    "textDocument": {"uri": uri}, "position": {"line": line, "character": 100000},
                })
                expected = len(text[:offset(text, line, utf16_length(content))].encode("utf-8"))
                assert actual == expected
            changes = [
                {"text": "a😀b\r\n中\n"},
                {"range": {"start": {"line": 0, "character": 1}, "end": {"line": 0, "character": 3}}, "rangeLength": 2, "text": "XY"},
                {"range": {"start": {"line": 0, "character": 3}, "end": {"line": 1, "character": 1}}, "text": "done"},
            ]
            for change in changes:
                text = edit(text, change)
            client.notification("textDocument/didChange", {"textDocument": {"uri": uri, "version": 4}, "contentChanges": changes})
            assert client.request("test/document", {"uri": uri}) == {"text": text, "version": 4}
            client.notification("textDocument/didChange", {"textDocument": {"uri": uri, "version": 5}, "contentChanges": [
                {"text": "partial"}, {"range": {"start": {"line": 99, "character": 0}, "end": {"line": 99, "character": 0}}, "text": "invalid"},
            ]})
            assert client.request("test/document", {"uri": uri}) == {"text": text, "version": 4}
            hover = client.request("textDocument/hover", {"textDocument": {"uri": uri}, "position": {"line": 0, "character": 0}})
            assert hover["contents"]["kind"] == "markdown"
            client.notification("textDocument/didClose", {"textDocument": {"uri": uri}})
            client.request("test/document", {"uri": uri}, error=-32602)
        assert client.request("shutdown") is None
        client.request("test/echo", {}, error=-32600)
        client.notification("exit")
        _, stderr = client.process.communicate(timeout=10)
        assert client.process.returncode == 0, stderr.decode()
        print(f"LSP interoperability: {client.responses} replies, {position_cases} independently checked Unicode positions, 30 transactional change sequences")
    finally:
        client.close()


def check_termination(executable):
    cases = [
        (b"Content-Length: 10\r\n\r\nx", 1),
        (b"Content-Length: -1\r\n\r\n", 1),
        (b"", 1),
    ]
    body = json.dumps({"jsonrpc": "2.0", "method": "exit"}).encode()
    cases.append((f"Content-Length: {len(body)}\r\n\r\n".encode() + body, 1))
    for wire, expected in cases:
        result = subprocess.run([str(executable), "--stdio"], input=wire, capture_output=True, timeout=10)
        assert result.returncode == expected, result
        assert not result.stdout, result.stdout
    print("LSP termination: premature exit, clean EOF, truncated body and invalid framing checked")


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--consumer", type=Path, default=Path(__file__).resolve().parents[1] / "consumers/lsp/_artifact/bin/lsp")
    args = parser.parse_args()
    if not args.consumer.is_file():
        raise RuntimeError("build the consumer with ecosystem/verify.py lsp first")
    check_server(args.consumer)
    check_termination(args.consumer)


if __name__ == "__main__":
    main()
