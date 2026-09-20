# progress

Concurrent progress bars, spinners and coordinated logging for GoML terminal applications. State updates are atomic across copied handles; deterministic snapshots and rendering are separate from terminal output.

```toml
[dependencies]
"ecosystem::progress" = "0.1.0"
"ecosystem::terminal" = "0.1.0"
```

```gom
use ecosystem::progress;
use ecosystem::terminal;
use std::context;

fn download() -> Result[(), progress::Error] {
    let session = terminal::Session::open(terminal::Options::defaults())?;
    defer { let _ = session.close(); };
    let manager = progress::Manager::new(session, progress::Options::defaults())?;
    defer { let _ = manager.shutdown(); };
    let job = manager.add("download", Option::Some(100))?;
    for _ in 0..100 {
        job.advance(1)?;
        manager.tick(context::Context::background())?;
    }
    job.finish("saved")?;
    manager.shutdown()
}
```

## Model and lifecycle

`Manager::add(name, total)` returns a shared `Job`. `None` creates a spinner with an unbounded logical total; `Some(0)` is an empty, complete-sized bar. Position and total are integers in `0..MAX_POSITION` (`2^53 - 1`). `advance`, `set_position`, `set_total`, and `set_message` are atomic. Positions cannot decrease or exceed a known total; totals cannot decrease. Invalid updates leave the entry unchanged. A job at its total is still running until explicitly finished.

`finish(message)` marks success and advances to a known total. `fail(message)` and `cancel(message)` preserve the current position. Terminal states reject updates except `reset(total)`, which starts a fresh run with position zero, an empty message, and fresh timing. `remove(context)` frees its capacity and invalidates every alias of that job handle; identifiers are never reused. Completed jobs stay visible until removed. `Job::apply_with(context, Update)` provides cancellation-aware access to every update.

`Manager::snapshot` returns detached `JobSnapshot` values including elapsed time, average units per second, estimated remaining milliseconds, status, and spinner frame. The `Clock` callback supplies monotonic milliseconds; `Clock::monotonic` is the default. Regressions and negative readings are clamped to the last observed timestamp. Values above `MAX_POSITION` saturate. Rate is absent until both elapsed time and progress are positive. ETA is absent without a known total, rate, or representable estimate. Rate includes idle time. Completed elapsed time and rate freeze; reset clears the estimate history. A clock callback runs under the manager lock and must be quick and non-reentrant.

## Drawing and logs

Updates modify state only. Call `tick(context)` from the application's event loop; it coalesces updates and respects `refresh_ms`. `flush(context)` forces pending output. Running interactive spinners animate on ticks; redirected output prints only changed jobs. There is no implicit background worker or stdin reader. A scoped task may call `tick` periodically if the application needs independent refresh.

`Mode::Auto` uses the supplied session's interactive capability. `Lines` always emits plain, newline-delimited progress; `Interactive` enables cursor control explicitly. Automatic output queries terminal size on each draw, reserves the final column to avoid autowrap, limits rows to the screen, and summarizes hidden jobs. `set_width(Some(columns))` overrides width; `None` restores the configured fallback or terminal width. `Style` customizes single-column spinner frames, bar cells/width, elapsed/rate/ETA fields, and four ANSI status styles. Color respects session capabilities and the requested ANSI profile. `render_job` and `render` are pure, bounded, grapheme-aware renderers usable in memory and application-owned displays.

`log(context, message)` erases the currently drawn block, writes the log, and redraws under one output lock. Interactive log text wraps at the drawable width; an indivisible grapheme wider than that width is clipped. Lines mode writes the complete message. Names, messages and logs must be single-line text without C0/C1 or Unicode line/paragraph controls, preventing cursor-control injection. They are never parsed as ANSI.

One manager owns one live progress region. Route concurrent logs through that manager. The supplied terminal session serializes writes, but direct writes, other progress managers, prompts, or full-screen renderers cannot automatically preserve each other's layouts. To combine a TUI with progress, render snapshots inside the TUI and leave progress output undrawn. No input is consumed by this package.

## Cancellation, cleanup and limits

Every operation using a caller context honors cancellation while waiting for the state/output lock and during terminal backpressure. Background-context convenience updates can wait behind a draw. `shutdown` closes the manager, wakes queued operations, interrupts an active write, cancels unfinished jobs, and emits the final state with a 200 ms cleanup deadline. It is shared and idempotent and retains its first result. **It leaves the caller's terminal session open.** The caller closes that session to restore modes and cursor visibility.

A write error, including cancellation during a possible partial write, becomes sticky: subsequent drawing returns that error instead of blindly replaying output. Snapshots and state updates remain available until shutdown. Cleanup after a partial write is best effort; an interrupted escape sequence cannot guarantee a pristine display. Shutdown returns the earlier write error if present, otherwise its cleanup result. Cancellation before acquiring the lock does not poison output.

Defaults: 128 retained jobs, 4 KiB per name/message/log, 20 visible rows, 80 fallback columns, 100 ms refresh/spinner intervals. Configurable hard limits: 1,000 jobs, 64 KiB text, 100 visible rows, 4,096 columns, 64 single-column spinner frames, and 512 bar columns. Each operation handles bounded retained state. No log queue accumulates: output pressure is synchronous and cancellable. Terminal I/O currently uses the Linux amd64 `terminal` backend. There is no byte-stream read/write adapter, recursive job tree, pause/resume accounting, arbitrary template language, automatic drop cleanup, or persistent background daemon.

## Validation

From the repository root:

```sh
GOML_BUILD_JOBS=2 python3 ecosystem/verify.py progress
python3 ecosystem/progress/interop.py
python3 ecosystem/progress/pty_test.py
python3 ecosystem/progress/race.py
```

The test suite covers atomic and concurrent updates, frozen timing and reset, invalid-input rollback, capacity and stale handles, detached snapshots, Unicode/style rendering, deterministic throttling, plain logs, interactive redraw/resize, shared-session input ownership, cancellation before work, queued deadlines, output pressure, sticky errors, shutdown wakeups, and session reuse. The independent consumer exercises the public dependency boundary. The reference oracle checks 400 rate/ETA/spinner/bar cases using Python arithmetic; PTY tests use a real shared terminal session and verify restored termios/flags/cursor modes.
