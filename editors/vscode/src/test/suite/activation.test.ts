import * as assert from 'assert';
import * as fs from 'fs';
import * as path from 'path';
import * as vm from 'vm';
import type { ExtensionContext } from 'vscode';
import { DiagnosticsScheduler } from '../../diagnosticsScheduler';

interface Subscription {
  dispose(): void;
}

function activationHarness(initiallyEnabled: boolean) {
  let enabled = initiallyEnabled;
  let lookups = 0;
  let starts = 0;
  let stops = 0;
  let watchers = 0;
  let statusEnabled = true;
  let onConfiguration: ((event: { affectsConfiguration(name: string): boolean }) => void) | undefined;
  const errors: string[] = [];
  const disposable = (): Subscription => ({ dispose() {} });
  const extensionExports: {
    activate(context: ExtensionContext): Promise<void>;
    deactivate(): Promise<void>;
  } = {} as typeof extensionExports;
  const modules: Record<string, unknown> = {
    vscode: {
      commands: { registerCommand: disposable },
      workspace: {
        getConfiguration: () => ({
          get: (name: string, fallback: unknown) => name === 'lsp.enabled' ? enabled : fallback,
        }),
        onDidChangeConfiguration: (callback: typeof onConfiguration) => {
          onConfiguration = callback;
          return disposable();
        },
        createFileSystemWatcher: () => {
          watchers++;
          return { dispose() { watchers--; } };
        },
        onDidChangeTextDocument: disposable,
        onDidSaveTextDocument: disposable,
        onDidCloseTextDocument: disposable,
      },
      window: { showErrorMessage: (message: string) => errors.push(message) },
    },
    'vscode-languageclient/node': {
      LanguageClient: class {
        outputChannel = { error() {} };
        onDidChangeState = disposable;
        onNotification = disposable;
        async start() { starts++; }
        async dispose() { stops++; }
      },
      NotificationType: class {},
      State: { Running: 2, Starting: 1, Stopped: 3 },
      TransportKind: { stdio: 0 },
    },
    './status': {
      LanguageServerStatus: class {
        middleware = {};
        setEnabled(value: boolean) { statusEnabled = value; }
        setClientState() {}
        attachOutputChannel() {}
        setUnavailable(message: string) { errors.push(message); }
        dispose() {}
      },
    },
    './diagnosticsScheduler': { DiagnosticsScheduler },
    './executables': {
      findExecutable: () => {
        lookups++;
        return '/fake/gomllsp';
      },
    },
    path,
  };
  const filename = path.resolve(__dirname, '../../extension.js');
  vm.runInNewContext(fs.readFileSync(filename, 'utf8'), {
    exports: extensionExports,
    require: (name: string) => {
      assert.ok(Object.prototype.hasOwnProperty.call(modules, name), `Unexpected module: ${name}`);
      return modules[name];
    },
    process,
  }, { filename });
  const subscriptions: Subscription[] = [];
  return {
    activate: () => extensionExports.activate({
      extensionPath: '/fake/extension',
      subscriptions,
    } as ExtensionContext),
    async configure(value: boolean) {
      enabled = value;
      onConfiguration?.({ affectsConfiguration: name => name === 'goml.lsp.enabled' });
      await new Promise<void>(resolve => setImmediate(resolve));
    },
    async dispose() {
      await extensionExports.deactivate();
      subscriptions.forEach(subscription => subscription.dispose());
    },
    snapshot: () => ({ lookups, starts, stops, watchers, statusEnabled, errors }),
  };
}

suite('GoML activation', () => {
  test('Highlighting mode activates without resolving or starting a server', async () => {
    const harness = activationHarness(false);
    try {
      await harness.activate();
      assert.deepStrictEqual(harness.snapshot(), {
        lookups: 0, starts: 0, stops: 0, watchers: 0, statusEnabled: false, errors: [],
      });
    } finally {
      await harness.dispose();
    }
  });

  test('Configuration changes stop, dispose and restart server resources', async () => {
    const harness = activationHarness(true);
    try {
      await harness.activate();
      await harness.configure(false);
      assert.deepStrictEqual(harness.snapshot(), {
        lookups: 1, starts: 1, stops: 1, watchers: 0, statusEnabled: false, errors: [],
      });
      await harness.configure(true);
      assert.deepStrictEqual(harness.snapshot(), {
        lookups: 2, starts: 2, stops: 1, watchers: 1, statusEnabled: true, errors: [],
      });
    } finally {
      await harness.dispose();
    }
    assert.strictEqual(harness.snapshot().watchers, 0);
    assert.strictEqual(harness.snapshot().stops, 2);
  });
});
