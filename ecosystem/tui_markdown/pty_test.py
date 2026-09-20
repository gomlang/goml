import codecs
import errno
import fcntl
import os
from pathlib import Path
import pty
import select
import struct
import subprocess
import termios
import time
import unicodedata


BINARY = Path(__file__).resolve().parents[1] / "consumers/tui_markdown/_artifact/bin/tui_markdown"


class Screen:
    def __init__(self, columns, rows):
        self.decoder = codecs.getincrementaldecoder("utf8")()
        self.pending = ""
        self.reverse = False
        self.resize(columns, rows)

    def resize(self, columns, rows):
        self.columns, self.rows = columns, rows
        self.cells = [[(" ", False) for _ in range(columns)] for _ in range(rows)]
        self.x = self.y = 0
        self.head = None
        self.join_next = False

    def sgr(self, parameters):
        codes = [int(item or 0) for item in parameters.split(";")]
        at = 0
        while at < len(codes):
            code = codes[at]
            if code == 0:
                self.reverse = False
            elif code == 7:
                self.reverse = True
            elif code == 27:
                self.reverse = False
            elif code in (38, 48) and at + 1 < len(codes):
                at += 2 if codes[at + 1] == 5 else 4
            at += 1

    def append_cluster(self, character):
        if self.head is not None:
            row, column = self.head
            value, reverse = self.cells[row][column]
            self.cells[row][column] = (value + character, reverse)

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
                        self.head = None
                        self.join_next = False
                    elif final == "m":
                        self.sgr(parameters)
                    elif final == "J" and parameters in ("2", "3"):
                        self.cells = [[(" ", False) for _ in range(self.columns)] for _ in range(self.rows)]
                    at = end + 1
                    continue
                at += 2
                continue
            if character == "\r":
                self.x = 0
                self.head = None
            elif character == "\n":
                self.y = min(self.rows - 1, self.y + 1)
                self.head = None
            elif character == "\u200d":
                self.append_cluster(character)
                self.join_next = True
            elif self.join_next:
                self.append_cluster(character)
                self.join_next = False
            elif unicodedata.combining(character) or character in ("\ufe0e", "\ufe0f"):
                self.append_cluster(character)
            elif ord(character) >= 32 and self.x < self.columns:
                width = 2 if unicodedata.east_asian_width(character) in "WF" else 1
                self.cells[self.y][self.x] = (character, self.reverse)
                self.head = (self.y, self.x)
                if width == 2 and self.x + 1 < self.columns:
                    self.cells[self.y][self.x + 1] = ("", self.reverse)
                self.x += width
            at += 1
        self.pending = self.pending[at:]

    def text(self):
        return "\n".join("".join(value for value, _ in row) for row in self.cells)

    def highlighted(self, word):
        for row in self.cells:
            for start in range(self.columns - len(word) + 1):
                cells = row[start:start + len(word)]
                if "".join(value for value, _ in cells) == word and all(reverse for _, reverse in cells):
                    return True
        return False


def main():
    master, slave = pty.openpty()
    original = termios.tcgetattr(slave)
    flags = fcntl.fcntl(slave, fcntl.F_GETFL)
    fcntl.ioctl(slave, termios.TIOCSWINSZ, struct.pack("HHHH", 8, 48, 0, 0))
    screen = Screen(48, 8)
    output = bytearray()
    child = subprocess.Popen([str(BINARY), "--pty"], stdin=slave, stdout=slave, stderr=subprocess.PIPE, env=dict(os.environ, TERM="xterm-256color", COLORTERM="truecolor", NO_COLOR=""))

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
                errors = child.stderr.read() if child.poll() is not None else b""
                raise AssertionError((screen.text(), child.poll(), errors, bytes(output[-1000:])))
            pump()

    try:
        until(lambda: "Markdown Viewer" in screen.text() and "界👩‍💻." in screen.text() and "second item" in screen.text())
        assert termios.tcgetattr(slave)[3] & (termios.ECHO | termios.ICANON) == 0
        assert screen.cells[2][15][0] == "界", screen.text()
        assert screen.cells[2][16][0] == "", screen.text()
        assert screen.cells[2][17][0] == "👩‍💻", screen.text()
        assert screen.cells[2][18][0] == "", screen.text()
        assert screen.cells[2][19][0] == ".", screen.text()
        assert not screen.highlighted("Guide")
        os.write(master, b"\x1b[B")
        until(lambda: "Markdown Viewer" not in screen.text() and "Hello GoML" in screen.text())
        os.write(master, b"\x1b[A")
        until(lambda: "Markdown Viewer" in screen.text())
        os.write(master, b"\x1b[6~")
        until(lambda: "Name" in screen.text() and "Value" in screen.text() and "┼" in screen.text())
        assert "界" in screen.text(), screen.text()
        os.write(master, b"\t")
        until(lambda: screen.highlighted("Guide"))
        os.write(master, b"\x1b[Z")
        until(lambda: screen.highlighted("Guide"))
        os.write(master, b"\x1b[6~")
        until(lambda: "let ready = true;" in screen.text() and "┌ gom" in screen.text())
        os.write(master, b"\x1b[F")
        until(lambda: "Last paragraph." in screen.text() and "Safe terminal text" in screen.text())
        os.write(master, b"\x1b[5~")
        until(lambda: "Last paragraph." not in screen.text())
        os.write(master, b"\x1b[H")
        until(lambda: "Markdown Viewer" in screen.text())
        screen.resize(32, 10)
        fcntl.ioctl(slave, termios.TIOCSWINSZ, struct.pack("HHHH", 10, 32, 0, 0))
        until(lambda: screen.highlighted("Guide") and "let ready = true;" in screen.text())
        assert b"\x1b[10;32H" in output
        os.write(master, b"\x1b[H")
        until(lambda: "Markdown Viewer" in screen.text() and "界👩‍💻." in screen.text())
        os.write(master, b"\x1b")
        deadline = time.monotonic() + 10
        while child.poll() is None and time.monotonic() < deadline:
            pump()
        assert child.wait(timeout=1) == 0
        pump(0)
        assert child.stderr.read() == b""
        assert termios.tcgetattr(slave) == original
        assert fcntl.fcntl(slave, fcntl.F_GETFL) == flags
        assert b"\x1b[?1049h" in output and b"\x1b[?1049l" in output
        assert b"\x1b[?25l" in output and b"\x1b[?25h" in output
        print("tui_markdown PTY passed: Unicode cell positions, line/page/Home/End scrolling, code/table rendering, Tab/BackTab link highlighting, resize reflow, exit and terminal restoration")
    finally:
        if child.poll() is None:
            child.kill()
            child.wait()
        child.stderr.close()
        os.close(master)
        os.close(slave)


if __name__ == "__main__":
    main()
