import {
  Disposable,
  LogOutputChannel,
  StatusBarAlignment,
  StatusBarItem,
  ThemeColor,
  window,
} from 'vscode';
import { Middleware, State } from 'vscode-languageclient/node';
import { performance } from 'perf_hooks';
import { formatDuration, requestLabel } from './statusText';

const slowOperationThreshold = 2000;

interface Activity {
  label: string;
  startedAt: number;
}

interface CompletedActivity {
  label: string;
  duration: number;
}

export interface ServerStatusNotification {
  state: 'busy' | 'ready';
  operation: string;
}

export class LanguageServerStatus implements Disposable {
  readonly middleware: Middleware;

  private readonly item: StatusBarItem;
  private readonly editorSubscription: Disposable;
  private readonly requests = new Map<number, Activity>();
  private clientState = State.Stopped;
  private nextRequestId = 0;
  private outputChannel: LogOutputChannel | undefined;
  private serverOperation: Activity | undefined;
  private serverStatusSupported = false;
  private lastActivity: CompletedActivity | undefined;
  private unavailableMessage: string | undefined;
  private timer: NodeJS.Timeout | undefined;
  private disposed = false;

  constructor() {
    this.item = window.createStatusBarItem(
      'goml.languageServerStatus',
      StatusBarAlignment.Left,
      100
    );
    this.item.name = 'GoML Language Server';
    this.item.command = 'goml.showLspOutput';
    this.middleware = {
      sendRequest: async (type, param, token, next) => {
        const method = typeof type === 'string' ? type : type.method;
        const requestId = this.beginRequest(method);
        try {
          return await next(type, param, token);
        } finally {
          this.endRequest(requestId);
        }
      },
    };
    this.editorSubscription = window.onDidChangeActiveTextEditor(() => {
      this.updateVisibility();
    });
    this.render();
    this.updateVisibility();
  }

  attachOutputChannel(outputChannel: LogOutputChannel): void {
    this.outputChannel = outputChannel;
  }

  setClientState(state: State): void {
    if (this.disposed) {
      return;
    }
    this.clientState = state;
    if (state !== State.StartFailed) {
      this.unavailableMessage = undefined;
    }
    if (state !== State.Running) {
      this.serverOperation = undefined;
      this.requests.clear();
      this.serverStatusSupported = false;
      this.lastActivity = undefined;
    }
    this.render();
    this.syncTimer();
  }

  setUnavailable(message: string): void {
    if (this.disposed) {
      return;
    }
    this.unavailableMessage = message;
    this.clientState = State.StartFailed;
    this.serverOperation = undefined;
    this.requests.clear();
    this.render();
    this.syncTimer();
  }

  handleServerStatus(status: ServerStatusNotification): void {
    if (this.disposed || this.clientState !== State.Running) {
      return;
    }
    this.serverStatusSupported = true;
    if (status.state === 'busy') {
      this.serverOperation = {
        label: status.operation,
        startedAt: performance.now(),
      };
    } else {
      if (this.serverOperation) {
        this.completeActivity(this.serverOperation);
      }
      this.serverOperation = undefined;
    }
    this.render();
    this.syncTimer();
  }

  showOutput(): void {
    if (this.outputChannel) {
      this.outputChannel.show(true);
    } else {
      void window.showErrorMessage(
        this.unavailableMessage ?? 'GoML language server output is unavailable.'
      );
    }
  }

  dispose(): void {
    this.disposed = true;
    if (this.timer) {
      clearInterval(this.timer);
      this.timer = undefined;
    }
    this.requests.clear();
    this.serverOperation = undefined;
    this.editorSubscription.dispose();
    this.item.dispose();
  }

  private beginRequest(method: string): number {
    if (this.disposed) {
      return -1;
    }
    const requestId = this.nextRequestId;
    this.nextRequestId += 1;
    this.requests.set(requestId, {
      label: requestLabel(method),
      startedAt: performance.now(),
    });
    this.render();
    this.syncTimer();
    return requestId;
  }

  private endRequest(requestId: number): void {
    if (this.disposed) {
      return;
    }
    const activity = this.requests.get(requestId);
    this.requests.delete(requestId);
    if (activity && !this.serverStatusSupported) {
      this.completeActivity(activity);
    }
    this.render();
    this.syncTimer();
  }

  private completeActivity(activity: Activity): void {
    const duration = performance.now() - activity.startedAt;
    this.lastActivity = { label: activity.label, duration };
    if (duration >= slowOperationThreshold) {
      this.outputChannel?.warn(
        `${activity.label} took ${formatDuration(duration)}`
      );
    }
  }

  private render(): void {
    this.item.backgroundColor = undefined;
    if (this.unavailableMessage) {
      this.item.text = '$(error) GoML: Unavailable';
      this.item.tooltip = `${this.unavailableMessage}\n\nClick to open language server output.`;
      this.item.backgroundColor = new ThemeColor('statusBarItem.errorBackground');
      this.item.accessibilityInformation = { label: 'GoML language server unavailable' };
      return;
    }
    if (this.clientState === State.Starting) {
      this.item.text = '$(sync~spin) GoML: Starting';
      this.item.tooltip = 'GoML language server is starting.';
      this.item.accessibilityInformation = { label: 'GoML language server starting' };
      return;
    }
    if (this.clientState === State.StartFailed) {
      this.item.text = '$(error) GoML: Failed';
      this.item.tooltip = 'GoML language server failed to start. Click to open output.';
      this.item.backgroundColor = new ThemeColor('statusBarItem.errorBackground');
      this.item.accessibilityInformation = { label: 'GoML language server failed' };
      return;
    }
    if (this.clientState === State.Stopped) {
      this.item.text = '$(circle-slash) GoML: Stopped';
      this.item.tooltip = 'GoML language server is stopped.';
      this.item.accessibilityInformation = { label: 'GoML language server stopped' };
      return;
    }
    const activity = this.serverOperation ?? this.requests.values().next().value;
    if (activity) {
      const duration = performance.now() - activity.startedAt;
      const extraRequests = this.serverOperation ? 0 : this.requests.size - 1;
      const suffix = extraRequests > 0 ? ` +${extraRequests}` : '';
      this.item.text = `$(sync~spin) GoML: ${activity.label}${suffix} (${formatDuration(duration)})`;
      this.item.tooltip = this.busyTooltip(activity, this.serverOperation !== undefined);
      if (duration >= slowOperationThreshold) {
        this.item.backgroundColor = new ThemeColor('statusBarItem.warningBackground');
      }
      this.item.accessibilityInformation = {
        label: `GoML language server busy with ${activity.label}`,
      };
      return;
    }
    this.item.text = '$(check) GoML: Ready';
    this.item.tooltip = this.readyTooltip();
    this.item.accessibilityInformation = { label: 'GoML language server ready' };
  }

  private busyTooltip(activity: Activity, isServerOperation: boolean): string {
    const activityKind = isServerOperation
      ? 'Current server operation'
      : 'Waiting for server response';
    const lines = [
      `${activityKind}: ${activity.label} (${formatDuration(performance.now() - activity.startedAt)})`,
    ];
    for (const request of this.requests.values()) {
      lines.push(
        `Client request: ${request.label} (${formatDuration(performance.now() - request.startedAt)})`
      );
    }
    lines.push('Click to open language server output.');
    return lines.join('\n\n');
  }

  private readyTooltip(): string {
    const lines = ['GoML language server is ready.'];
    if (this.lastActivity) {
      lines.push(
        `Last operation: ${this.lastActivity.label} (${formatDuration(this.lastActivity.duration)})`
      );
    }
    lines.push('Click to open language server output.');
    return lines.join('\n\n');
  }

  private syncTimer(): void {
    const needsTimer = this.clientState === State.Running
      && (this.serverOperation !== undefined || this.requests.size > 0);
    if (needsTimer && !this.timer) {
      this.timer = setInterval(() => this.render(), 100);
    } else if (!needsTimer && this.timer) {
      clearInterval(this.timer);
      this.timer = undefined;
    }
  }

  private updateVisibility(): void {
    if (window.activeTextEditor?.document.languageId === 'goml') {
      this.item.show();
    } else {
      this.item.hide();
    }
  }
}
