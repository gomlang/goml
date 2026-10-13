import * as assert from 'assert';
import { formatDuration, requestLabel } from '../../statusText';

suite('GoML Language Server Status Tests', () => {
  test('Labels common language server requests', () => {
    assert.strictEqual(requestLabel('textDocument/hover'), 'Hover');
    assert.strictEqual(requestLabel('textDocument/inlayHint'), 'Inlay hints');
    assert.strictEqual(requestLabel('goml/expandedDerive'), 'Expanding derive');
    assert.strictEqual(requestLabel('custom/request'), 'custom/request');
    assert.strictEqual(requestLabel('constructor'), 'constructor');
    assert.strictEqual(requestLabel('__proto__'), '__proto__');
  });

  test('Formats elapsed request time', () => {
    assert.strictEqual(formatDuration(50), '0.1s');
    assert.strictEqual(formatDuration(2450), '2.5s');
    assert.strictEqual(formatDuration(12500), '13s');
  });
});
