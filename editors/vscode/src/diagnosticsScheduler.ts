export class DiagnosticsScheduler {
  private readonly timers = new Map<string, NodeJS.Timeout>();

  constructor(
    private readonly delay: number,
    private readonly check: (uri: string) => void
  ) {}

  schedule(uri: string): void {
    this.clear(uri);
    this.timers.set(uri, setTimeout(() => this.request(uri), this.delay));
  }

  request(uri: string): void {
    this.clear(uri);
    this.check(uri);
  }

  clear(uri: string): void {
    const timer = this.timers.get(uri);
    if (timer !== undefined) {
      clearTimeout(timer);
      this.timers.delete(uri);
    }
  }

  dispose(): void {
    for (const timer of this.timers.values()) {
      clearTimeout(timer);
    }
    this.timers.clear();
  }
}
