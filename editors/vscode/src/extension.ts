import * as path from 'path';
import * as fs from 'fs';
import {
    commands,
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

let client: LanguageClient | undefined;
let diagnosticsTimer: NodeJS.Timeout | undefined;
let diagnosticsUri: string | undefined;

const serverStatusNotification = new NotificationType<ServerStatusNotification>('goml/status');
const checkDiagnosticsDelay = 750;

function clearDiagnosticsTimer(): void {
    if (diagnosticsTimer) {
        clearTimeout(diagnosticsTimer);
        diagnosticsTimer = undefined;
    }
    diagnosticsUri = undefined;
}

function requestDiagnostics(uri: string): void {
    clearDiagnosticsTimer();
    if (client?.state === State.Running) {
        void client.sendNotification('goml/checkDiagnostics', {
            textDocument: { uri },
        });
    }
}

function scheduleDiagnostics(uri: string): void {
    clearDiagnosticsTimer();
    diagnosticsUri = uri;
    diagnosticsTimer = setTimeout(() => {
        const pendingUri = diagnosticsUri;
        clearDiagnosticsTimer();
        if (pendingUri) {
            requestDiagnostics(pendingUri);
        }
    }, checkDiagnosticsDelay);
}

export function activate(context: ExtensionContext) {
    const status = new LanguageServerStatus();
    context.subscriptions.push(status);

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
                if (document?.isDirty) {
                    await document.save();
                }
                const folder = workspace.getWorkspaceFolder(uri);
                const scope = folder ?? TaskScope.Workspace;
                const cwd = folder?.uri.fsPath ?? path.dirname(uri.fsPath);
                const task = new Task(
                    { type: 'goml', task: 'test' },
                    scope,
                    `test ${testName}`,
                    'goml',
                    new ProcessExecution(
                        'goml',
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
            if (!client) {
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

    const serverOptions: ServerOptions = {
        run: {
            command: serverPath,
            transport: TransportKind.stdio,
        },
        debug: {
            command: serverPath,
            transport: TransportKind.stdio,
        },
    };

    const clientOptions: LanguageClientOptions = {
        documentSelector: [{ scheme: 'file', language: 'goml' }],
        synchronize: {
            fileEvents: workspace.createFileSystemWatcher('**/*.{gom,toml,go,mod,sum}'),
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
    context.subscriptions.push(
        client.onDidChangeState(event => {
            status.setClientState(event.newState);
        })
    );
    context.subscriptions.push(
        client.onNotification(serverStatusNotification, serverStatus => {
            status.handleServerStatus(serverStatus);
        })
    );
    context.subscriptions.push(
        workspace.onDidChangeTextDocument(event => {
            if (event.document.languageId === 'goml') {
                scheduleDiagnostics(event.document.uri.toString());
            }
        })
    );
    context.subscriptions.push(
        workspace.onDidSaveTextDocument(document => {
            if (document.languageId === 'goml') {
                requestDiagnostics(document.uri.toString());
            }
        })
    );
    context.subscriptions.push(
        workspace.onDidCloseTextDocument(document => {
            if (document.uri.toString() === diagnosticsUri) {
                clearDiagnosticsTimer();
            }
        })
    );
    void client.start().catch(error => {
        const message = error instanceof Error ? error.message : String(error);
        status.setUnavailable(`GoML language server failed to start: ${message}`);
    });
}

export function deactivate(): Thenable<void> | undefined {
    clearDiagnosticsTimer();
    if (!client) {
        return undefined;
    }
    return client.stop();
}

function findServerPath(context: ExtensionContext): string | undefined {
    const config = workspace.getConfiguration('goml');
    const configPath = config.get<string>('serverPath');

    if (configPath && configPath.length > 0 && fs.existsSync(configPath)) {
        return configPath;
    }

    const binary = 'gomllsp';
    const bundledPath = path.join(context.extensionPath, 'bin', binary);
    if (fs.existsSync(bundledPath)) {
        return bundledPath;
    }

    const bundledPathExe = path.join(context.extensionPath, 'bin', `${binary}.exe`);
    if (fs.existsSync(bundledPathExe)) {
        return bundledPathExe;
    }

    const envPath = process.env.PATH;
    if (envPath) {
        const pathDirs = envPath.split(path.delimiter);
        for (const dir of pathDirs) {
            const candidate = path.join(dir, binary);
            if (fs.existsSync(candidate)) {
                return candidate;
            }
            const candidateExe = path.join(dir, `${binary}.exe`);
            if (fs.existsSync(candidateExe)) {
                return candidateExe;
            }
        }
    }

    return undefined;
}
