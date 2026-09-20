import json
from http.server import BaseHTTPRequestHandler, ThreadingHTTPServer
from pathlib import Path
import subprocess
import threading
from urllib.parse import parse_qs, urlsplit


ROOT = Path(__file__).resolve().parents[1]
BINARY = ROOT / "consumers/reqwest/_artifact/bin/reqwest"


class Handler(BaseHTTPRequestHandler):
    protocol_version = "HTTP/1.1"

    def log_message(self, format, *args):
        pass

    def do_GET(self):
        self.handle_request()

    def do_POST(self):
        self.handle_request()

    def handle_request(self):
        address = urlsplit(self.path)
        body = self.rfile.read(int(self.headers.get("Content-Length", "0")))
        headers = []
        status = 200
        if address.path == "/echo":
            payload = json.dumps({"method": self.command, "body": body.decode(), "query": address.query,
                                  "authorization": self.headers.get("Authorization", ""),
                                  "custom": "|".join(self.headers.get_all("X-Custom", []))}).encode()
        elif address.path == "/json":
            payload = body
        elif address.path == "/redirect":
            payload = b""
            status = 302
            headers.append(("Location", parse_qs(address.query)["to"][0]))
        elif address.path == "/headers":
            payload = b"headers"
            headers.extend([("X-Many", "one"), ("X-Many", "two")])
        elif address.path == "/large":
            payload = b"x" * 4096
        elif address.path == "/chunked":
            self.send_response(200)
            self.send_header("Transfer-Encoding", "chunked")
            self.end_headers()
            self.wfile.write(b"9\r\nchunk-one\r\n9\r\nchunk-two\r\n0\r\n\r\n")
            self.wfile.flush()
            return
        else:
            payload = b"unknown"
            status = 404
        self.send_response(status)
        for name, value in headers:
            self.send_header(name, value)
        self.send_header("Content-Length", str(len(payload)))
        self.end_headers()
        self.wfile.write(payload)


def main():
    server = ThreadingHTTPServer(("127.0.0.1", 0), Handler)
    worker = threading.Thread(target=server.serve_forever, daemon=True)
    worker.start()
    try:
        subprocess.run([str(BINARY), f"http://127.0.0.1:{server.server_port}"], check=True, timeout=30)
    finally:
        server.shutdown()
        server.server_close()
        worker.join()
    print("reqwest interoperability: Python HTTP/1.1, duplicate headers, Unicode query/form, typed JSON, redirect, limits and chunked transfer passed")


if __name__ == "__main__":
    main()
