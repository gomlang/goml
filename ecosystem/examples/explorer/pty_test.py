import codecs
import errno
import fcntl
import os
from pathlib import Path
import pty
import re
import select
import struct
import subprocess
import tempfile
import termios
import time
import unicodedata


ROOT = Path(__file__).resolve().parent
BINARY = ROOT / "_artifact/bin/explorer"


class Screen:
    def __init__(self, columns, rows):
        self.decoder = codecs.getincrementaldecoder("utf8")()
        self.pending = ""
        self.resize(columns, rows)

    def resize(self, columns, rows):
        self.columns, self.rows = columns, rows
        self.cells = [[" " for _ in range(columns)] for _ in range(rows)]
        self.x = self.y = 0

    def feed(self, data):
        self.pending += self.decoder.decode(data)
        at = 0
        while at < len(self.pending):
            character = self.pending[at]
            if character == "\x1b":
                if at + 1 == len(self.pending):
                    break
                if self.pending[at + 1] == "[":
                    end = at + 2
                    while end < len(self.pending) and not "@" <= self.pending[end] <= "~":
                        end += 1
                    if end == len(self.pending):
                        break
                    parameters = self.pending[at + 2:end]
                    final = self.pending[end]
                    if final in "Hf":
                        values = [int(item or "1") for item in parameters.split(";")]
                        self.y = max(0, min(self.rows - 1, values[0] - 1))
                        self.x = max(0, min(self.columns - 1, (values[1] if len(values) > 1 else 1) - 1))
                    elif final == "J" and parameters in ("2", "3"):
                        self.cells = [[" " for _ in range(self.columns)] for _ in range(self.rows)]
                    at = end + 1
                    continue
                at += 2
                continue
            if character == "\r":
                self.x = 0
            elif character == "\n":
                self.y = min(self.rows - 1, self.y + 1)
            elif ord(character) >= 32:
                if unicodedata.combining(character):
                    if self.x:
                        self.cells[self.y][self.x - 1] += character
                elif self.x < self.columns:
                    width = 2 if unicodedata.east_asian_width(character) in "WF" else 1
                    self.cells[self.y][self.x] = character
                    if width == 2 and self.x + 1 < self.columns:
                        self.cells[self.y][self.x + 1] = ""
                    self.x += width
            at += 1
        self.pending = self.pending[at:]

    def text(self):
        return "\n".join("".join(row) for row in self.cells)


def main():
    with tempfile.TemporaryDirectory(prefix="goml-explorer-") as directory:
        root = Path(directory)
        readme = root / "README.md"
        readme.write_text("# Initial preview\n\nA read-only file browser.\n", encoding="utf8")
        (root / "docs").mkdir()
        (root / "docs/guide.md").write_text("# Guide heading\n\n" + "\n\n".join(f"Guide paragraph {index:03d}" for index in range(100)), encoding="utf8")
        (root / "_artifact").mkdir()
        (root / "_artifact/ignored.md").write_text("# Hidden")
        master, slave = pty.openpty()
        before = termios.tcgetattr(slave)
        flags = fcntl.fcntl(slave, fcntl.F_GETFL)
        fcntl.ioctl(slave, termios.TIOCSWINSZ, struct.pack("HHHH", 28, 110, 0, 0))
        screen = Screen(110, 28)
        output = bytearray()
        child = subprocess.Popen([str(BINARY), str(root)], stdin=slave, stdout=slave, stderr=subprocess.PIPE, env=dict(os.environ, TERM="xterm-256color"))

        def pump(timeout=0.1):
            ready, _, _ = select.select([master], [], [], timeout)
            if ready:
                try:
                    chunk = os.read(master, 65536)
                except OSError as error:
                    if error.errno != errno.EIO:
                        raise
                    chunk = b""
                if chunk:
                    output.extend(chunk)
                    screen.feed(chunk)

        def until(predicate):
            deadline = time.monotonic() + 15
            while not predicate():
                if time.monotonic() > deadline or child.poll() is not None:
                    raise AssertionError((screen.text(), child.poll(), bytes(output[-1000:])))
                pump()

        try:
            until(lambda: "Initial preview" in screen.text() and "Scan complete" in screen.text())
            assert "README.md" in screen.text() and "ignored.md" not in screen.text()
            assert "4/4" in screen.text()
            assert termios.tcgetattr(slave)[3] & (termios.ECHO | termios.ICANON) == 0
            readme.write_text("# Updated preview\n\nRoot notifications refreshed this file.\n", encoding="utf8")
            until(lambda: "Updated preview" in screen.text() and re.search(r"changes [1-9]", screen.text()))
            fresh = root / "fresh.md"
            fresh.write_text("# Fresh file\n")
            until(lambda: "fresh.md" in screen.text() and "Scan complete" in screen.text())
            fresh.unlink()
            until(lambda: "fresh.md" not in screen.text() and "Scan complete" in screen.text())
            os.write(master, b"\x1b[B\x1b[B")
            until(lambda: "Guide heading" in screen.text())
            os.write(master, b"\t")
            until(lambda: "Markdown [focus]" in screen.text())
            os.write(master, b"\x1b[6~")
            until(lambda: "Guide heading" not in screen.text() and "Guide paragraph" in screen.text())
            prior = re.search(r"scans (\d+)", screen.text())
            prior = int(prior[1])
            os.write(master, b"r")
            until(lambda: any(int(value) > prior for value in re.findall(r"scans (\d+)", screen.text())) and "Scan complete" in screen.text())
            screen.resize(84, 20)
            fcntl.ioctl(slave, termios.TIOCSWINSZ, struct.pack("HHHH", 20, 84, 0, 0))
            until(lambda: "GoML Explorer" in screen.text() and "q quit" in screen.text())
            os.write(master, b"q")
            deadline = time.monotonic() + 10
            while child.poll() is None and time.monotonic() < deadline:
                pump()
            assert child.wait(timeout=1) == 0
            pump(0)
            assert child.stderr.read() == b""
            assert termios.tcgetattr(slave) == before
            assert fcntl.fcntl(slave, fcntl.F_GETFL) == flags
            assert b"\x1b[?1049h" in output and b"\x1b[?1049l" in output
            assert b"\x1b[?25l" in output and b"\x1b[?25h" in output
            assert readme.read_text() == "# Updated preview\n\nRoot notifications refreshed this file.\n"
            print("explorer PTY passed: background scan, tree navigation, Markdown scrolling, root watch modify/create/delete, refresh, resize, clean exit and terminal restoration")
        finally:
            if child.poll() is None:
                child.kill()
                child.wait()
            child.stderr.close()
            os.close(master)
            os.close(slave)


if __name__ == "__main__":
    main()
