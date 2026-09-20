import argparse
import os
from pathlib import Path
import socket
import ssl
import subprocess
import tempfile
import threading


ROOT = Path(__file__).resolve().parent
CONSUMER = ROOT.parent / 'consumers' / 'redis'


def command(stream):
    header = stream.readline(256)
    if not header:
        return None
    if not header.startswith(b'*') or not header.endswith(b'\r\n'):
        raise AssertionError(f'invalid command header: {header!r}')
    count = int(header[1:-2])
    if not 0 < count < 32:
        raise AssertionError(f'invalid argument count: {count}')
    values = []
    for _ in range(count):
        header = stream.readline(256)
        if not header.startswith(b'$') or not header.endswith(b'\r\n'):
            raise AssertionError(f'invalid argument header: {header!r}')
        length = int(header[1:-2])
        if not 0 <= length < 1024:
            raise AssertionError(f'invalid argument length: {length}')
        value = stream.read(length + 2)
        if len(value) != length + 2 or value[-2:] != b'\r\n':
            raise AssertionError('truncated command')
        values.append(value[:-2])
    return values


def exercise(binary, certificate, key, mode):
    listener = socket.socket()
    listener.bind(('127.0.0.1', 0))
    listener.listen(1)
    listener.settimeout(8)
    port = listener.getsockname()[1]
    stopped = threading.Event()
    stalled = threading.Event()
    failures = []
    observed = []

    def serve():
        try:
            raw, _ = listener.accept()
            with raw:
                raw.settimeout(8)
                if mode == 'handshake_timeout':
                    stopped.wait(8)
                    return
                if mode in ('plain', 'pool_plain'):
                    connection = raw
                else:
                    config = ssl.SSLContext(ssl.PROTOCOL_TLS_SERVER)
                    config.load_cert_chain(certificate, key)
                    if mode in ('mutual', 'pool_mutual'):
                        config.load_verify_locations(certificate)
                        config.verify_mode = ssl.CERT_REQUIRED
                    try:
                        connection = config.wrap_socket(raw, server_side=True)
                    except ssl.SSLError:
                        if mode not in ('untrusted', 'name'):
                            raise
                        return
                with connection, connection.makefile('rb') as stream:
                    while True:
                        request = command(stream)
                        if request is None:
                            return
                        observed.append(request)
                        if request[0] == b'HELLO':
                            if mode == 'hello_timeout':
                                stalled.set()
                                stopped.wait(8)
                                return
                            if request != [b'HELLO', b'3']:
                                raise AssertionError(f'unexpected handshake: {request!r}')
                            connection.sendall(b'%1\r\n+proto\r\n:3\r\n')
                        elif request == [b'PING']:
                            connection.sendall(b'+PONG\r\n')
                        elif request == [b'ECHO', b'stall']:
                            connection.sendall(b'$10\r\nx')
                            stalled.set()
                            stopped.wait(8)
                            return
                        elif request == [b'ECHO', b'secure\x00echo']:
                            data = request[1]
                            connection.sendall(b'$' + str(len(data)).encode() + b'\r\n' + data + b'\r\n')
                        else:
                            raise AssertionError(f'unexpected command: {request!r}')
        except Exception as error:
            failures.append(error)

    worker = threading.Thread(target=serve, daemon=True)
    worker.start()
    process = subprocess.Popen([str(binary), '--tls', 'localhost', str(port), str(certificate), str(key), mode], stdin=subprocess.PIPE, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True, env=dict(os.environ, GORACE='halt_on_error=1 atexit_sleep_ms=0'))
    try:
        if mode in ('cancel', 'legacy') and not stalled.wait(8):
            raise AssertionError(f'{mode}: request did not reach the cancellation barrier')
        stdout, stderr = process.communicate('\n', timeout=12)
        if process.returncode:
            raise AssertionError(f'{mode}: {stdout}{stderr}')
    finally:
        if process.poll() is None:
            process.kill()
            process.communicate()
        stopped.set()
        listener.close()
        worker.join(10)
    if worker.is_alive():
        raise AssertionError(f'{mode}: server did not stop')
    if failures:
        raise AssertionError(f'{mode}: {failures!r}')
    if mode in ('plain', 'tls', 'mutual') and len(observed) != 5:
        raise AssertionError(f'{mode}: incomplete command sequence: {observed!r}')
    if mode in ('pool_plain', 'pool_tls', 'pool_mutual', 'pool_timeout') and len(observed) != 8:
        raise AssertionError(f'{mode}: incomplete pooled command sequence: {observed!r}')
    if mode in ('untrusted', 'name', 'handshake_timeout') and observed:
        raise AssertionError(f'{mode}: commands were sent before TLS validation')


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('--race', action='store_true')
    args = parser.parse_args()
    binary = CONSUMER / '_artifact/bin/redis'
    artifacts = ROOT / '_artifact/network'
    artifacts.mkdir(parents=True, exist_ok=True)
    if args.race:
        binary = artifacts / 'consumer-race'
        source = CONSUMER / '_artifact/build/pkg/consumer/redis/goml_generated.go'
        subprocess.run(['go', 'build', '-race', '-o', str(binary), str(source)], cwd=CONSUMER, check=True, timeout=120)
    with tempfile.TemporaryDirectory(dir=artifacts) as directory:
        certificate = Path(directory) / 'certificate.pem'
        key = Path(directory) / 'key.pem'
        subprocess.run(['openssl', 'req', '-x509', '-newkey', 'rsa:2048', '-nodes', '-days', '2', '-subj', '/CN=localhost', '-addext', 'subjectAltName=DNS:localhost,IP:127.0.0.1', '-keyout', str(key), '-out', str(certificate)], check=True, capture_output=True, timeout=15)
        modes = ('plain', 'tls', 'mutual', 'untrusted', 'name', 'handshake_timeout', 'hello_timeout', 'io_timeout', 'context_timeout', 'cancel', 'legacy', 'pool_plain', 'pool_tls', 'pool_mutual', 'pool_timeout')
        for mode in modes:
            exercise(binary, certificate, key, mode)
    print(f'Redis DNS/TLS/context checks: {len(modes)} passed' + (' under race detector' if args.race else ''))


if __name__ == '__main__':
    main()
