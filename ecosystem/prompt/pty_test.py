import errno
import fcntl
import os
from pathlib import Path
import pty
import re
import select
import struct
import subprocess
import termios
import time

BINARY = Path(__file__).resolve().parents[1] / "consumers/prompt/_artifact/bin/prompt"
CSI = re.compile(rb"\x1b\[[0-?]*[ -/]*[@-~]")


class Child:
    def __init__(self, mode):
        self.master, self.slave = pty.openpty()
        fcntl.ioctl(self.slave, termios.TIOCSWINSZ, struct.pack("HHHH", 8, 72, 0, 0))
        self.before = termios.tcgetattr(self.slave)
        self.flags = fcntl.fcntl(self.slave, fcntl.F_GETFL)
        self.process = subprocess.Popen([str(BINARY), mode], stdin=self.slave, stdout=self.slave, stderr=subprocess.PIPE, env=dict(os.environ, TERM="xterm-256color"))
        self.data = bytearray()

    def wait(self, marker):
        deadline = time.monotonic() + 10
        while marker not in CSI.sub(b"", self.data):
            assert time.monotonic() < deadline, (marker, bytes(self.data[-1000:]), self.process.poll())
            ready, _, _ = select.select([self.master], [], [], 0.1)
            if ready:
                try:
                    chunk = os.read(self.master, 16384)
                except OSError as error:
                    if error.errno != errno.EIO:
                        raise
                    chunk = b""
                assert chunk, (marker, bytes(self.data[-1000:]))
                self.data.extend(chunk)

    def send(self, value):
        os.write(self.master, value)

    def finish(self, marker):
        self.wait(marker)
        assert self.process.wait(timeout=10) == 0
        assert self.process.stderr.read() == b""
        assert termios.tcgetattr(self.slave) == self.before
        assert fcntl.fcntl(self.slave, fcntl.F_GETFL) == self.flags
        for sequence in (b"\x1b[?1049h", b"\x1b[?1049l", b"\x1b[?25l", b"\x1b[?25h", b"\x1b[?2004h", b"\x1b[?2004l"):
            assert sequence in self.data, sequence

    def close(self):
        if self.process.poll() is None:
            self.process.kill()
            self.process.wait()
        self.process.stderr.close()
        os.close(self.master)
        os.close(self.slave)


def scenario(mode, body):
    child = Child(mode)
    try:
        body(child)
    finally:
        child.close()


def text(child):
    child.wait(b"? Name")
    child.send("é👩‍💻".encode() + b"\x7f" + "中".encode() + b"\r")
    child.finish("RESULT é中".encode())


def validation(child):
    child.wait(b"? Port")
    child.send(b"70000\r")
    child.wait(b"between 1 and 65535")
    child.send(b"\x7f" * 5 + b"8080\r")
    child.finish(b"RESULT 8080")


def password(child):
    child.wait(b"? Secret")
    secret = b"SENSITIVE_PASSWORD"
    child.send(b"\x1b[200~" + secret + b"\x1b[201~")
    child.wait("•••".encode())
    child.send(b"\r")
    child.finish(b"RESULT secret-bytes=18")
    assert secret not in CSI.sub(b"", child.data)


def selection(child):
    child.wait(b"? Language")
    child.send(b"Rust\r")
    child.finish(b"RESULT rust")


def end(child, value, marker):
    child.wait(b"? Name")
    child.send(value)
    child.finish(marker)


def main():
    scenario("text", text)
    scenario("integer", validation)
    scenario("password", password)
    scenario("select", selection)
    scenario("text", lambda child: end(child, b"\x03", b"ERROR prompt cancelled"))
    scenario("text", lambda child: end(child, b"\x04", b"ERROR end of input"))
    scenario("text", lambda child: end(child, b"\xff", b"ERROR terminal:"))
    scenario("timeout", lambda child: child.finish(b"ERROR prompt deadline exceeded"))
    scenario("broken", lambda child: child.finish(b"deliberate render failure"))
    result = subprocess.run([str(BINARY), "text"], input=b"", capture_output=True, timeout=10, check=True)
    assert b"ERROR terminal:" in result.stdout and b"\x1b" not in result.stdout
    print("prompt: nine PTY sessions passed Unicode editing, validation retry, secret masking, search, cancellation, EOF, malformed input, deadline/render-error cleanup; redirected input rejected cleanly")


if __name__ == "__main__":
    main()
