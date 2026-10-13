const requestLabels: Record<string, string> = {
  'textDocument/hover': 'Hover',
  'textDocument/completion': 'Completion',
  'textDocument/signatureHelp': 'Signature help',
  'textDocument/inlayHint': 'Inlay hints',
  'textDocument/definition': 'Go to definition',
  'textDocument/codeLens': 'Code lenses',
  'textDocument/codeAction': 'Code actions',
  'textDocument/formatting': 'Formatting',
  'goml/expandedDerive': 'Expanding derive',
};

export function requestLabel(method: string): string {
  return Object.prototype.hasOwnProperty.call(requestLabels, method) ? requestLabels[method] : method;
}

export function formatDuration(duration: number): string {
  const seconds = duration / 1000;
  return `${seconds.toFixed(seconds < 10 ? 1 : 0)}s`;
}
