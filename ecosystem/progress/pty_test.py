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

ROOT = Path(__file__).resolve().parent
BINARY = ROOT.parent / "consumers/progress/_artifact/bin/progress"


def main():
    master, slave = pty.openpty()
    original = termios.tcgetattr(slave)
    flags = fcntl.fcntl(slave, fcntl.F_GETFL)
    fcntl.ioctl(slave, termios.TIOCSWINSZ, struct.pack("HHHH", 12, 80, 0, 0))
    child = subprocess.Popen([str(BINARY), "pty"], stdin=slave, stdout=slave, stderr=subprocess.PIPE, env=dict(os.environ, TERM="xterm-256color"))
    output = bytearray()

    def until(marker):
        deadline = time.monotonic() + 10
        while marker not in output:
            assert time.monotonic() < deadline, (marker, bytes(output[-4000:]), child.poll())
            ready, _, _ = select.select([master], [], [], 0.1)
            if ready:
                try:
                    data = os.read(master, 65536)
                except OSError as error:
                    if error.errno != errno.EIO:
                        raise
                    data = b""
                assert data, (marker, bytes(output[-4000:]))
                output.extend(data)

    try:
        until(b"READY")
        raw = termios.tcgetattr(slave)
        assert not raw[3] & (termios.ECHO | termios.ICANON)
        os.write(master, b"u")
        until(b"UPDATED")
        until(b"10/100")
        os.write(master, b"l")
        until(b"LOG coordinated output")
        assert b"\r\x1b[1A\x1b[2K\x1b[1A\x1b[2KLOG coordinated output\r\n" in output
        before_resize = len(output)
        fcntl.ioctl(slave, termios.TIOCSWINSZ, struct.pack("HHHH", 8, 24, 0, 0))
        until(b"RESIZED 24")
        os.write(master, b"q")
        until(b"CLOSED")
        until(b"\x1b[?25h")
        assert child.wait(timeout=10) == 0
        assert child.stderr.read() == b""
        assert termios.tcgetattr(slave) == original
        assert fcntl.fcntl(slave, fcntl.F_GETFL) == flags
        assert b"\x1b[?25l" in output and b"\x1b[?25h" in output
        resized = re.sub(rb"\x1b\[[0-9;?]*[A-Za-z]", b"", bytes(output[before_resize:])).decode()
        assert "+ download" in resized and "x 扫描" in resized, resized
        assert b"\x1b[?1049" not in output
        redirected = subprocess.run([str(BINARY)], stdin=subprocess.DEVNULL, capture_output=True, check=True, timeout=10)
        assert b"\x1b" not in redirected.stdout
        assert b"located 12 files\n" in redirected.stdout
        assert b"12/12" in redirected.stdout
        print("progress PTY: raw shared input, bars/spinners, coordinated log redraw, Unicode, resize, finish/cancel, cursor/termios/flags restoration and redirected plain output passed")
    finally:
        if child.poll() is None:
            child.kill()
            child.wait()
        child.stderr.close()
        os.close(master)
        os.close(slave)


if __name__ == "__main__":
    main()
