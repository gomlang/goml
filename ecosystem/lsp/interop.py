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

    def read(self, protocol_message=False):
        header = bytearray()
        while not header.endswith(b"\r\n\r\n"):
            header.extend(self.read_exact(1))
            assert len(header) <= 8192
        fields = dict(line.split(b":", 1) for line in bytes(header[:-4]).split(b"\r\n"))
        length = int(fields[b"Content-Length"])
        assert 0 <= length <= 16777216
        reply = json.loads(self.read_exact(length).decode("utf-8"))
        assert reply["jsonrpc"] == "2.0"
        if not protocol_message:
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


def encoding_length(value, encoding):
    if encoding == "utf-8":
        return len(value.encode("utf-8"))
    if encoding == "utf-16":
        return utf16_length(value)
    return len(value)


def check_encodings(executable):
    positions = 0
    edits = 0
    for encoding in ("utf-8", "utf-16", "utf-32"):
        client = Client(executable)
        try:
            client.request("initialize", {"capabilities": {"general": {"positionEncodings": [encoding, 7]}}}, error=-32602)
            result = client.request("initialize", {"capabilities": {"general": {"positionEncodings": ["future-encoding", encoding]}}})
            assert result["capabilities"]["positionEncoding"] == encoding
            client.notification("initialized", {})
            rng = random.Random(2407)
            for sample in range(16):
                text = "".join(rng.choice(["a", "中", "😀", "é", "e\u0301", "\r", "\n", "\r\n"]) for _ in range(24))
                if sample == 0:
                    text = "x" * 1023 + "\r\n😀中é\r\ne\u0301\n"
                uri = f"file:///{encoding}-{sample}"
                client.notification("textDocument/didOpen", {"textDocument": {"uri": uri, "languageId": "text", "version": 1, "text": text}})
                line_start = [0] + [m.end() for m in re.finditer(r"\r\n|\r|\n", text)]
                for line, content in enumerate(re.split(r"\r\n|\r|\n", text)):
                    start_byte = len(text[:line_start[line]].encode("utf-8"))
                    valid = {encoding_length(content[:index], encoding): start_byte + len(content[:index].encode("utf-8")) for index in range(len(content) + 1)}
                    for character in range(encoding_length(content, encoding) + 1):
                        params = {"textDocument": {"uri": uri}, "position": {"line": line, "character": character}}
                        if character in valid:
                            assert client.request("test/offset", params) == valid[character]
                        else:
                            client.request("test/offset", params, error=-32602)
                        positions += 1
                    assert client.request("test/offset", {"textDocument": {"uri": uri}, "position": {"line": line, "character": 2147483647}}) == start_byte + len(content.encode("utf-8"))
                emoji_width = encoding_length("😀", encoding)
                wide_width = encoding_length("中", encoding)
                changes = [
                    {"text": "a😀b\r\n中\n"},
                    {"range": {"start": {"line": 0, "character": 1}, "end": {"line": 0, "character": 1 + emoji_width}}, "rangeLength": emoji_width, "text": "中"},
                    {"range": {"start": {"line": 0, "character": 1}, "end": {"line": 1, "character": wide_width}}, "rangeLength": wide_width * 2 + 3, "text": "done"},
                ]
                client.notification("textDocument/didChange", {"textDocument": {"uri": uri, "version": 2}, "contentChanges": changes})
                assert client.request("test/document", {"uri": uri}) == {"text": "adone\n", "version": 2}
                client.notification("textDocument/didChange", {"textDocument": {"uri": uri, "version": 3}, "contentChanges": [
                    {"text": "a😀"},
                    {"range": {"start": {"line": 0, "character": 1}, "end": {"line": 0, "character": 1 + emoji_width}}, "rangeLength": 99, "text": "invalid"},
                ]})
                assert client.request("test/document", {"uri": uri}) == {"text": "adone\n", "version": 2}
                client.notification("textDocument/didClose", {"textDocument": {"uri": uri}})
                edits += 1
            assert client.request("shutdown") is None
            client.notification("exit")
            _, stderr = client.process.communicate(timeout=10)
            assert client.process.returncode == 0, stderr.decode()
        finally:
            client.close()
    print(f"LSP encoding negotiation: {positions} UTF-8/16/32 positions and invalid boundaries, {edits} transactional change sequences")


def check_outgoing_deadlines(executable):
    client = Client(executable)
    try:
        client.request("initialize", {"capabilities": {}})
        client.notification("initialized", {})
        client.notification("test/outgoing", {"id": "zero", "timeout": 0})
        request = client.read(protocol_message=True)
        assert request["id"] == "zero" and request["method"] == "client/echo", request
        client.notification("test/poll", {})
        cancellation = client.read(protocol_message=True)
        assert cancellation["method"] == "$/cancelRequest" and cancellation["params"]["id"] == "zero", cancellation
        expired = client.read(protocol_message=True)
        assert expired["method"] == "test/completed" and expired["params"]["error"]["code"] == -32803, expired
        client.send({"jsonrpc": "2.0", "id": "zero", "result": "late"})
        assert client.request("test/echo", {}) == {}
        client.notification("test/outgoing", {"id": "early", "timeout": 10000})
        assert client.read(protocol_message=True)["id"] == "early"
        client.send({"jsonrpc": "2.0", "id": "early", "result": {"unicode": "😀"}})
        completed = client.read(protocol_message=True)
        assert completed["method"] == "test/completed" and completed["params"]["result"] == {"unicode": "😀"}, completed
        client.notification("test/outgoing", {"id": "late", "timeout": 20})
        assert client.read(protocol_message=True)["id"] == "late"
        time.sleep(0.05)
        client.send({"jsonrpc": "2.0", "id": "late", "result": "overdue"})
        completed = client.read(protocol_message=True)
        assert completed["method"] == "test/completed" and completed["params"]["error"]["code"] == -32803, completed
        client.notification("test/outgoing", {"id": "poll", "timeout": 20})
        assert client.read(protocol_message=True)["id"] == "poll"
        time.sleep(0.05)
        client.notification("test/poll", {})
        assert client.read(protocol_message=True)["method"] == "$/cancelRequest"
        assert client.read(protocol_message=True)["params"]["error"]["code"] == -32803
        client.notification("test/poll", {})
        assert client.request("test/echo", {}) == {}
        client.notification("test/outgoing", {"id": "exit", "timeout": 10000})
        assert client.read(protocol_message=True)["id"] == "exit"
        assert client.request("shutdown") is None
        client.notification("exit")
        stdout, stderr = client.process.communicate(timeout=10)
        assert client.process.returncode == 0 and not stdout, (stdout, stderr)
        print("LSP outgoing deadlines: explicit polling, cancellation, response/deadline races, duplicate suppression and exit cleanup")
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
    check_encodings(args.consumer)
    check_outgoing_deadlines(args.consumer)
    check_termination(args.consumer)


if __name__ == "__main__":
    main()
