import * as assert from 'assert';
import * as fs from 'fs';
import * as os from 'os';
import * as path from 'path';
import { findExecutable } from '../../executables';

suite('GoML Executable Discovery Tests', () => {
  let directory: string;

  setup(() => {
    directory = fs.mkdtempSync(path.join(os.tmpdir(), 'goml-executable-'));
  });

  teardown(() => {
    fs.rmSync(directory, { recursive: true, force: true });
  });

  test('Skips directories and non-executable candidates', () => {
    fs.mkdirSync(path.join(directory, 'directory'));
    const executable = path.join(directory, 'server');
    fs.writeFileSync(executable, '', { mode: 0o755 });
    assert.strictEqual(findExecutable(undefined, [directory], ['directory', 'server']), executable);
    if (process.platform !== 'win32') {
      fs.writeFileSync(path.join(directory, 'not-executable'), '', { mode: 0o644 });
      assert.strictEqual(findExecutable(undefined, [directory], ['not-executable', 'server']), executable);
    }
  });

  test('An invalid configured server cannot silently select another toolchain', () => {
    fs.writeFileSync(path.join(directory, 'server'), '', { mode: 0o755 });
    assert.strictEqual(findExecutable(path.join(directory, 'missing'), [directory], ['server']), undefined);
  });
});
