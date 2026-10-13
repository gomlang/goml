import * as assert from 'assert';
import { setTimeout as delay } from 'timers/promises';
import { DiagnosticsScheduler } from '../../diagnosticsScheduler';

suite('GoML Diagnostics Scheduler Tests', () => {
  test('Retains pending checks for separate documents', async () => {
    const checked: string[] = [];
    const scheduler = new DiagnosticsScheduler(5, uri => checked.push(uri));
    try {
      scheduler.schedule('first');
      scheduler.schedule('second');
      scheduler.schedule('first');
      await delay(30);
      assert.deepStrictEqual(checked.sort(), ['first', 'second']);
    } finally {
      scheduler.dispose();
    }
  });

  test('Immediate checks and closing affect only their own document', async () => {
    const checked: string[] = [];
    const scheduler = new DiagnosticsScheduler(5, uri => checked.push(uri));
    try {
      scheduler.schedule('saved');
      scheduler.schedule('closed');
      scheduler.schedule('pending');
      scheduler.request('saved');
      scheduler.clear('closed');
      assert.deepStrictEqual(checked, ['saved']);
      await delay(30);
      assert.deepStrictEqual(checked, ['saved', 'pending']);
    } finally {
      scheduler.dispose();
    }
  });

  test('Disposal cancels all pending checks', async () => {
    const checked: string[] = [];
    const scheduler = new DiagnosticsScheduler(5, uri => checked.push(uri));
    scheduler.schedule('first');
    scheduler.schedule('second');
    scheduler.dispose();
    await delay(30);
    assert.deepStrictEqual(checked, []);
  });
});
