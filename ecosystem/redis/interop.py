import hashlib
import json
from pathlib import Path
import random
import socket
import subprocess
import sys
import tarfile
import tempfile
import time
import urllib.request


ROOT = Path(__file__).resolve().parents[1]
REFERENCE = ROOT / '_artifact' / 'reference'
BINARY = ROOT / 'consumers' / 'redis' / '_artifact' / 'bin' / 'redis'
VERSION = '7.2.5'
URL = f'https://codeload.github.com/redis/redis/tar.gz/refs/tags/{VERSION}'
SHA256 = '98a8502a2e902d2a9785ef46a69a5f8d5e24cbf9ea3ae4d845afcfc6778aa783'


def reference_server():
    REFERENCE.mkdir(parents=True, exist_ok=True)
    archive = REFERENCE / f'redis-{VERSION}-github.tar.gz'
    if not archive.exists():
        with urllib.request.urlopen(URL, timeout=60) as response:
            data = response.read()
        if hashlib.sha256(data).hexdigest() != SHA256:
            raise RuntimeError('Redis reference archive checksum mismatch')
        archive.write_bytes(data)
    if hashlib.sha256(archive.read_bytes()).hexdigest() != SHA256:
        raise RuntimeError('Redis reference archive checksum mismatch')
    source = REFERENCE / f'redis-{VERSION}'
    if not source.exists():
        with tarfile.open(archive) as contents:
            contents.extractall(REFERENCE, filter='data')
    binary = source / 'src' / 'redis-server'
    if not binary.exists():
        with (REFERENCE / 'redis-build.log').open('w') as log:
            subprocess.run(['make', '-j2', '-C', str(source / 'src'), 'MALLOC=libc', 'REDIS_CFLAGS=', 'REDIS_LDFLAGS=', 'redis-server'], stdout=log, stderr=subprocess.STDOUT, check=True, timeout=240)
    actual = subprocess.check_output([str(binary), '--version'], text=True)
    if f'v={VERSION} ' not in actual:
        raise RuntimeError(f'Unexpected reference version: {actual}')
    return binary


def live_cases():
    binary = reference_server()
    password = 'ecosystem-reference-only'
    with tempfile.TemporaryDirectory(prefix='redis-session-', dir=REFERENCE) as directory:
        with socket.socket() as reservation:
            reservation.bind(('127.0.0.1', 0))
            port = reservation.getsockname()[1]
        with (REFERENCE / 'redis-server.log').open('w') as log:
            server = subprocess.Popen([str(binary), '--bind', '127.0.0.1', '--port', str(port), '--save', '', '--appendonly', 'no', '--protected-mode', 'yes', '--daemonize', 'no', '--dir', directory, '--requirepass', password], stdout=log, stderr=subprocess.STDOUT)
            try:
                deadline = time.monotonic() + 10
                while True:
                    if server.poll() is not None:
                        raise RuntimeError('Reference server exited; see redis-server.log')
                    try:
                        with socket.create_connection(('127.0.0.1', port), timeout=0.2):
                            break
                    except OSError:
                        if time.monotonic() > deadline:
                            raise RuntimeError('Reference Redis startup timed out')
                        time.sleep(0.02)
                for version in ('2', '3'):
                    result = subprocess.run([str(BINARY), '--server', f'127.0.0.1:{port}', version, password], capture_output=True, text=True, timeout=30)
                    if result.returncode:
                        raise RuntimeError(result.stdout + result.stderr)
                    print(result.stdout.strip(), flush=True)
            finally:
                server.terminate()
                try:
                    server.wait(timeout=10)
                except subprocess.TimeoutExpired:
                    server.kill()
                    server.wait()


def line(tag, value=b''):
    return tag + value + b'\r\n'


def bulk(tag, payload):
    return line(tag, str(len(payload)).encode()) + payload + b'\r\n'


def aggregate(tag, values):
    return line(tag, str(len(values)).encode()) + b''.join(values)


def protocol_cases():
    rng = random.Random(270496)
    requests = []
    expected = []

    def add(source, canonical=None, chunk=0, error=False):
        requests.append({'hex': source.hex(), 'chunk': chunk})
        expected.append(None if error else [(source if canonical is None else canonical).hex()])

    scalars = [b'+OK\r\n', b'-WRONGTYPE wrong\r\n', b'$-1\r\n', b'*-1\r\n', b'_\r\n', b'#t\r\n', b'#f\r\n', b',inf\r\n', b',-inf\r\n', b',nan\r\n', b',1.25\r\n', b'(1234567890123456789012345678901234567890\r\n']
    scalars.extend(line(b':', str(value).encode()) for value in [-2**63, -1, 0, 1, 2**63 - 1])
    for length in [0, 1, 2, 31, 255, 256, 4096, 65536]:
        payload = bytes(rng.randrange(256) for _ in range(length))
        scalars.extend([bulk(b'$', payload), bulk(b'!', b'ERR ' + payload), bulk(b'=', b'txt:' + payload)])
    for source in scalars:
        add(source)
        for chunk in (1, 2, 7, 257):
            add(source, chunk=chunk)
    for _ in range(400):
        values = rng.choices(scalars[:25], k=rng.randrange(7))
        tag = rng.choice([b'*', b'~'])
        canonical = aggregate(tag, values)
        add(canonical)
        streamed = line(tag, b'?') + b''.join(values) + b'.\r\n'
        add(streamed, canonical, rng.choice([1, 3, 17]))
        pairs = [(rng.choice(scalars[:20]), rng.choice(scalars[:20])) for _ in range(rng.randrange(5))]
        fixed_map = line(b'%', str(len(pairs)).encode()) + b''.join(k + v for k, v in pairs)
        streamed_map = b'%?\r\n' + b''.join(k + v for k, v in pairs) + b'.\r\n'
        add(streamed_map, fixed_map, rng.choice([1, 8, 97]))
        value = rng.choice(scalars)
        attributed = line(b'|', str(len(pairs)).encode()) + b''.join(k + v for k, v in pairs) + value
        add(attributed, chunk=rng.choice([0, 1, 4096]))
        chunks = [bytes(rng.randrange(256) for _ in range(rng.randrange(1, 50))) for _ in range(rng.randrange(8))]
        source = b'$?\r\n' + b''.join(bulk(b';', chunk) for chunk in chunks) + b';0\r\n'
        add(source, bulk(b'$', b''.join(chunks)), rng.choice([0, 1, 5, 31]))
    for source in [aggregate(b'*', scalars[:12]), b'|1\r\n+k\r\n:1\r\n>2\r\n+invalidate\r\n*1\r\n+k\r\n', b'$?\r\n;3\r\nabc\r\n;0\r\n']:
        for end in range(len(source)):
            add(source[:end], error=True)
    for source in [b'+x\n', b':9223372036854775808\r\n', b'$-2\r\n', b'#true\r\n', b',1e9999\r\n', b'%?\r\n+k\r\n.\r\n', b'>0\r\n', b'>1\r\n:1\r\n', b'*1\r\n.\r\n', b'=4\r\ntxtx\r\n', b'$?\r\n;-1\r\n', b'$99999999999999999999999\r\n', b'*1000001\r\n']:
        add(source, error=True)
    requests.append({'hex': b''.join(scalars[:17]).hex(), 'chunk': 3})
    expected.append([source.hex() for source in scalars[:17]])
    result = subprocess.run([str(BINARY), '--codec-json'], input=json.dumps(requests), capture_output=True, text=True, timeout=60)
    if result.returncode:
        raise RuntimeError(result.stdout + result.stderr)
    actual = json.loads(result.stdout)
    if len(actual) != len(expected):
        raise AssertionError('response count differs')
    for index, (value, want) in enumerate(zip(actual, expected)):
        if want is None:
            if not isinstance(value, dict) or 'error' not in value:
                raise AssertionError(f'case {index}: malformed input accepted')
        elif value != want:
            raise AssertionError(f'case {index}: {value!r} != {want!r}')
    print(f'RESP specification fixtures: {len(expected)} cases passed', flush=True)


def main():
    if not BINARY.exists():
        raise RuntimeError('Build the Redis consumer with ecosystem/verify.py first')
    protocol_cases()
    live_cases()
    subprocess.run([sys.executable, str(ROOT / 'redis' / 'network_check.py')], check=True, timeout=120)


if __name__ == '__main__':
    main()
