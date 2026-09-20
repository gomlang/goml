import os
from pathlib import Path
import re
import subprocess
import sys


ROOT = Path(__file__).resolve().parent
GENERATED = ROOT / '_artifact' / 'test' / 'external' / 'goml_generated.go'
BINARY = ROOT / '_artifact' / 'redis-race-tests'


def main():
    if not GENERATED.exists():
        raise RuntimeError('Run goml test in ecosystem/redis before race.py')
    subprocess.run(['go', 'build', '-race', '-o', str(BINARY), str(GENERATED)], cwd=ROOT, check=True, timeout=120)
    tests = sorted(name for source in (ROOT / 'tests').glob('*.gom') for name in re.findall(r'#\[test\]\s+fn\s+(\w+)\(', source.read_text()))
    environment = os.environ.copy()
    environment['GORACE'] = 'halt_on_error=1 atexit_sleep_ms=0'
    for name in tests:
        subprocess.run([str(BINARY), f'ecosystem::redis::tests::{name}'], cwd=ROOT, env=environment, check=True, timeout=30)
    print(f'Redis race detector: {len(tests)} tests passed')
    subprocess.run([sys.executable, str(ROOT / 'network_check.py'), '--race'], check=True, timeout=180)


if __name__ == '__main__':
    main()
