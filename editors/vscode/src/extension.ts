import * as path from 'path';
import {
  commands,
  Disposable,
  ExtensionContext,
  ProcessExecution,
  Task,
  TaskPanelKind,
  TaskRevealKind,
  TaskScope,
  tasks,
  Uri,
  ViewColumn,
  window,
  workspace,
} from 'vscode';
import {
  LanguageClient,
  LanguageClientOptions,
  NotificationType,
  ServerOptions,
  State,
  TransportKind,
} from 'vscode-languageclient/node';
import { LanguageServerStatus, ServerStatusNotification } from './status';
import { DiagnosticsScheduler } from './diagnosticsScheduler';
import { findExecutable } from './executables';

let client: LanguageClient | undefined;
let serverUpdate = Promise.resolve();
let deactivating = false;
const serverSubscriptions: Disposable[] = [];

const serverStatusNotification = new NotificationType<ServerStatusNotification>('goml/status');
const checkDiagnosticsDelay = 750;

const diagnostics = new DiagnosticsScheduler(checkDiagnosticsDelay, uri => {
  if (client?.state === State.Running) {
    void client.sendNotification('goml/checkDiagnostics', {
      textDocument: { uri },
    }).catch(error => {
      client?.outputChannel.error(`Failed to request diagnostics: ${String(error)}`);
    });
  }
});

export async function activate(context: ExtensionContext) {
  deactivating = false;
  const status = new LanguageServerStatus();
  context.subscriptions.push(status, diagnostics);

  context.subscriptions.push(
    commands.registerCommand('goml.showLspOutput', () => {
      status.showOutput();
    })
  );

  context.subscriptions.push(
    commands.registerCommand(
      'goml.runTest',
      async (uriText: string, testName: string, kind: string) => {
        const uri = Uri.parse(uriText);
        const document = workspace.textDocuments.find(
          candidate => candidate.uri.toString() === uri.toString()
        );
        if (document?.isDirty && !await document.save()) {
          void window.showErrorMessage('Save the GoML source file before running its test.');
          return;
        }
        const folder = workspace.getWorkspaceFolder(uri);
        const scope = folder ?? TaskScope.Workspace;
        const cwd = path.dirname(uri.fsPath);
        const driver = workspace.getConfiguration('goml', uri).get<string>('driverPath') || 'goml';
        const task = new Task(
          { type: 'goml', task: 'test' },
          scope,
          `test ${testName}`,
          'goml',
          new ProcessExecution(
            driver,
            ['test', testName, '--kind', kind],
            { cwd }
          )
        );
        task.presentationOptions = {
          reveal: TaskRevealKind.Always,
          panel: TaskPanelKind.Dedicated,
          clear: true,
        };
        await tasks.executeTask(task);
      }
    )
  );

  context.subscriptions.push(
    commands.registerCommand('goml.showExpandedDerive', async () => {
      const editor = window.activeTextEditor;
      if (!editor || editor.document.languageId !== 'goml') {
        window.showErrorMessage('Open a GoML source file to show expanded derives.');
        return;
      }
      if (client?.state !== State.Running) {
        window.showErrorMessage('GoML language server is not running.');
        return;
      }
      const text = await client.sendRequest<string>('goml/expandedDerive', {
        textDocument: { uri: editor.document.uri.toString() },
      });
      const document = await workspace.openTextDocument({ language: 'goml', content: text });
      await window.showTextDocument(document, {
        preview: true,
        viewColumn: ViewColumn.Beside,
      });
    })
  );

  context.subscriptions.push(workspace.onDidChangeConfiguration(event => {
    if (event.affectsConfiguration('goml.lsp.enabled')
      || event.affectsConfiguration('goml.serverPath')
      || event.affectsConfiguration('goml.driverPath')) {
      void updateLanguageServer(context, status);
    }
  }));
  await updateLanguageServer(context, status);
}

function updateLanguageServer(context: ExtensionContext, status: LanguageServerStatus): Promise<void> {
  serverUpdate = serverUpdate.then(async () => {
    await stopLanguageServer();
    const enabled = !deactivating && workspace.getConfiguration('goml').get<boolean>('lsp.enabled', true);
    status.setEnabled(enabled);
    if (enabled) {
      await startLanguageServer(context, status);
    }
  }).catch(error => {
    status.setUnavailable(`GoML language server failed: ${String(error)}`);
  });
  return serverUpdate;
}

async function stopLanguageServer(): Promise<void> {
  diagnostics.dispose();
  for (const subscription of serverSubscriptions.splice(0)) {
    subscription.dispose();
  }
  const previous = client;
  client = undefined;
  await previous?.dispose();
}

async function startLanguageServer(context: ExtensionContext, status: LanguageServerStatus): Promise<void> {
  const serverPath = findServerPath(context);

  if (!serverPath) {
    status.setUnavailable(
      'GoML language server not found. Install gomllsp or set goml.serverPath in settings.'
    );
    window.showErrorMessage(
      'GoML language server not found. Please install gomllsp or set goml.serverPath in settings.'
    );
    return;
  }

  const driverPath = workspace.getConfiguration('goml').get<string>('driverPath', '');
  const driverOptions = driverPath ? { env: { ...process.env, GOML_PROJECT_DRIVER: driverPath } } : undefined;
  const serverExecutable = {
    command: serverPath,
    transport: TransportKind.stdio,
    options: driverOptions,
  };
  const serverOptions: ServerOptions = {
    run: serverExecutable,
    debug: serverExecutable,
  };

  const watcher = workspace.createFileSystemWatcher('**/*.{goml,toml,work}');
  serverSubscriptions.push(watcher);
  const clientOptions: LanguageClientOptions = {
    documentSelector: [{ scheme: 'file', language: 'goml' }],
    synchronize: {
      fileEvents: watcher,
    },
    outputChannelName: 'GoML Language Server',
    middleware: status.middleware,
  };

  client = new LanguageClient(
    'goml',
    'GoML Language Server',
    serverOptions,
    clientOptions
  );

  status.attachOutputChannel(client.outputChannel);
  status.setClientState(State.Starting);
  serverSubscriptions.push(
    client.onDidChangeState(event => {
      status.setClientState(event.newState);
    })
  );
  serverSubscriptions.push(
    client.onNotification(serverStatusNotification, serverStatus => {
      status.handleServerStatus(serverStatus);
    })
  );
  serverSubscriptions.push(
    workspace.onDidChangeTextDocument(event => {
      if (event.document.languageId === 'goml') {
        diagnostics.schedule(event.document.uri.toString());
      }
    })
  );
  serverSubscriptions.push(
    workspace.onDidSaveTextDocument(document => {
      if (document.languageId === 'goml') {
        diagnostics.request(document.uri.toString());
      }
    })
  );
  serverSubscriptions.push(
    workspace.onDidCloseTextDocument(document => {
      diagnostics.clear(document.uri.toString());
    })
  );
  await client.start().catch(error => {
    const message = error instanceof Error ? error.message : String(error);
    status.setUnavailable(`GoML language server failed to start: ${message}`);
  });
}

export async function deactivate(): Promise<void> {
  deactivating = true;
  await serverUpdate;
  await stopLanguageServer();
}

function findServerPath(context: ExtensionContext): string | undefined {
  const config = workspace.getConfiguration('goml');
  const configPath = config.get<string>('serverPath');

  const directories = [
    path.join(context.extensionPath, 'bin'),
    ...(process.env.PATH ? process.env.PATH.split(path.delimiter) : []),
  ];
  return findExecutable(configPath, directories, ['gomllsp', 'gomllsp.exe']);
}
