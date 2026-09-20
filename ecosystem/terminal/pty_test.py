import errno
import fcntl
import os
from pathlib import Path
import pty
import select
import struct
import subprocess
import termios
import tempfile
import time

ROOT = Path(__file__).resolve().parent
BINARY = ROOT.parent / "consumers/terminal/_artifact/bin/terminal"


class TerminalProcess:
    def __init__(self, mode):
        self.master, self.slave = pty.openpty()
        fcntl.ioctl(self.slave, termios.TIOCSWINSZ, struct.pack("HHHH", 24, 80, 0, 0))
        self.before = termios.tcgetattr(self.slave)
        self.flags = fcntl.fcntl(self.slave, fcntl.F_GETFL)
        self.process = subprocess.Popen([str(BINARY), mode], stdin=self.slave, stdout=self.slave, stderr=subprocess.PIPE, env=dict(os.environ, TERM="xterm-256color"))
        self.output = bytearray()
        self.read_until(b"READY 80 24\n")
        raw = termios.tcgetattr(self.slave)
        assert raw[3] & (termios.ECHO | termios.ICANON | termios.ISIG) == 0, raw
        assert raw[1] & termios.OPOST == 0
        assert fcntl.fcntl(self.slave, fcntl.F_GETFL) == self.flags

    def read_until(self, marker, timeout=10):
        deadline = time.monotonic() + timeout
        while marker not in self.output:
            remaining = deadline - time.monotonic()
            assert remaining > 0, (marker, bytes(self.output[-2000:]), self.process.poll())
            ready, _, _ = select.select([self.master], [], [], min(remaining, 0.2))
            if ready:
                try:
                    data = os.read(self.master, 4096)
                except OSError as error:
                    if error.errno != errno.EIO:
                        raise
                    data = b""
                if not data:
                    raise AssertionError((marker, bytes(self.output[-2000:])))
                self.output.extend(data)
        return bytes(self.output)

    def send(self, data):
        os.write(self.master, data)

    def finish(self, marker=b"CLOSED"):
        self.read_until(marker)
        assert self.process.wait(timeout=10) == 0
        errors = self.process.stderr.read()
        assert errors == b"", errors
        assert termios.tcgetattr(self.slave) == self.before
        assert fcntl.fcntl(self.slave, fcntl.F_GETFL) == self.flags
        for marker in (b"\x1b[?1049h", b"\x1b[?25l", b"\x1b[?1002h", b"\x1b[?1006h", b"\x1b[?1004h", b"\x1b[?2004h", b"\x1b[?1049l", b"\x1b[?25h", b"\x1b[?1002l", b"\x1b[?1006l", b"\x1b[?1004l", b"\x1b[?2004l"):
            assert marker in self.output, marker
        os.close(self.master)
        os.close(self.slave)
        self.process.stderr.close()


def interaction():
    child = TerminalProcess("events")
    child.send(b"\x1b[")
    time.sleep(0.005)
    child.send(b"1;5A")
    child.read_until(b"control: true")
    encoded = "中".encode()
    for byte in encoded:
        child.send(bytes([byte]))
        time.sleep(0.005)
    child.read_until(encoded)
    child.send(b"\x1b")
    child.read_until(b"Escape")
    child.send(b"\x1b[200~hello\n")
    child.send("世界".encode() + b"\x1b[201~")
    child.read_until("PASTE hello\n世界\n".encode())
    child.send(b"\x1b[I\x1b[O\x1b[<0;12;8M\x1b[<0;12;8m")
    child.read_until(b"FOCUS false")
    child.read_until(b"Release")
    fcntl.ioctl(child.slave, termios.TIOCSWINSZ, struct.pack("HHHH", 41, 101, 0, 0))
    child.read_until(b"RESIZE 101 41\n")
    child.send(b"q")
    child.finish()


def output_pressure():
    child = TerminalProcess("write")
    time.sleep(0.08)
    child.finish()
    payload = bytes(child.output).split(b"READY 80 24\n", 1)[1].split(b"\nWRITE DONE\n", 1)[0]
    assert payload == b"0123456789abcdef" * 262144, len(payload)


def main():
    interaction()
    child = TerminalProcess("cancel")
    child.read_until(b"CANCEL TimedOut\n")
    child.finish()
    child = TerminalProcess("events")
    child.send(b"\xff")
    child.finish(b"ERROR InvalidData")
    output_pressure()
    for append in (False, True):
        with tempfile.TemporaryFile() as output:
            output.write(b"prefix")
            output.flush()
            if append:
                fcntl.fcntl(output, fcntl.F_SETFL, fcntl.fcntl(output, fcntl.F_GETFL) | os.O_APPEND)
                output.seek(0)
            flags = fcntl.fcntl(output, fcntl.F_GETFL)
            subprocess.run([str(BINARY), "redirect"], stdin=subprocess.DEVNULL, stdout=output, check=True, timeout=10)
            assert fcntl.fcntl(output, fcntl.F_GETFL) == flags
            output.seek(0)
            assert output.read() == b"prefixsessionDONE\n"
    print("terminal PTY: fragmented UTF-8/CSI/paste, ESC timeout, mouse, focus, resize, cancellation, error cleanup, raw/flag restoration 4 MiB partial writes and redirected file offsets/append flags passed")


if __name__ == "__main__":
    main()
