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
import termios
import time

BINARY = Path(__file__).resolve().parents[1] / "consumers/tui/_artifact/bin/tui"


def main():
    master, slave = pty.openpty()
    fcntl.ioctl(slave, termios.TIOCSWINSZ, struct.pack("HHHH", 6, 20, 0, 0))
    before = termios.tcgetattr(slave)
    process = subprocess.Popen([str(BINARY), "--pty"], stdin=slave, stdout=slave, stderr=subprocess.PIPE, env=dict(os.environ, TERM="xterm-256color"))
    data = bytearray()
    decoder = codecs.getincrementaldecoder("utf8")()
    visible = ""

    def wait_for(predicate):
        nonlocal visible
        deadline = time.monotonic() + 10
        while not predicate():
            assert time.monotonic() < deadline, (bytes(data[-1000:]), process.poll())
            ready, _, _ = select.select([master], [], [], 0.2)
            if ready:
                try:
                    chunk = os.read(master, 16384)
                except OSError as error:
                    if error.errno != errno.EIO:
                        raise
                    chunk = b""
                assert chunk, (bytes(data[-1000:]), process.poll())
                data.extend(chunk)
                visible += decoder.decode(chunk)

    try:
        wait_for(lambda: b"\x1b[6;20H" in data)
        raw = termios.tcgetattr(slave)
        assert raw[3] & (termios.ECHO | termios.ICANON) == 0
        os.write(master, "ab界".encode())
        wait_for(lambda: "界" in visible)
        os.write(master, b"\x7f")
        wait_for(lambda: b"\x1b[2;3H" in data and b"\x1b[2;4H " in data)
        os.write(master, b"\x1b[200~hello\nworld\x1b[201~")
        wait_for(lambda: b"\x1b[3;5H" in data)
        data.clear()
        fcntl.ioctl(slave, termios.TIOCSWINSZ, struct.pack("HHHH", 9, 32, 0, 0))
        wait_for(lambda: b"\x1b[9;32H" in data)
        os.write(master, b"\x1b")
        process.wait(timeout=10)
        errors = process.stderr.read()
        assert process.returncode == 0 and not errors, (process.returncode, errors)
        assert termios.tcgetattr(slave) == before
        print("tui PTY: interactive drawing, Unicode input/deletion, bracketed multiline paste, resize/full repaint and terminal restoration passed")
    finally:
        if process.poll() is None:
            process.kill()
            process.wait()
        process.stderr.close()
        os.close(master)
        os.close(slave)


if __name__ == "__main__":
    main()
