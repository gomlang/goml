# Explorer integration example

A read-only, two-pane terminal file browser written in GoML. It combines `walkdir`, `notify`, `terminal`, `tui`, `ansi`, `color`, `progress`, and `tui_markdown` through ordinary versioned dependencies. The left pane shows an expanded directory tree; the right pane previews the selected Markdown file. Scans run on one background task and report bounded batches to the event loop.

This is an application example with module path `example::explorer`, not a published ecosystem library or an entry in the shared library verifier's `MODULES` list.

## Build and run

From the repository root:

```sh
python3 ecosystem/examples/explorer/verify.py

ecosystem/examples/explorer/_artifact/bin/explorer .

ecosystem/examples/explorer/_artifact/bin/explorer --snapshot

ecosystem/examples/explorer/_artifact/bin/explorer --snapshot ecosystem/examples/explorer
```

The verifier builds against a private registry snapshot from `ecosystem/verify.py`; it does not publish packages or alter the user's registry. `--demo` is an alias for the built-in snapshot. Snapshots are 100 columns by 28 rows, plain text without terminal escapes. Interactive mode requires Linux amd64 and a real terminal; `--help` works without a terminal.

## Controls

| Key | Action |
| --- | --- |
| Up / Down, `k` / `j` | Select an entry while the tree has focus |
| Home / End / PageUp / PageDown | Navigate the tree selection |
| Tab / BackTab | Switch between tree and preview |
| Up / Down / PageUp / PageDown / Home / End | Scroll the focused Markdown preview |
| `r` | Refresh, coalescing repeated requests while a scan is active |
| `q`, Escape, Ctrl+C | Cancel the worker, join it, close watches, restore the terminal and exit |

The tree stays expanded, making every scanned entry directly reachable. Only `.md` and `.markdown` regular files are previewed; directories, other formats and symlinks show an explanatory message. Markdown rendering includes headings, lists, fenced code, quotes, tables and link labels. Links are displayed but never launched, fetched, or executed. Pane sizes adapt to terminal resize; tiny terminals show an enlargement message while quit remains available.

## Integration and lifecycle

One `terminal::Session` owns input and output. `tui::Terminal` commits frames and restores the alternate screen, cursor, termios and descriptor flags when closed. Progress is rendered from a `progress::JobSnapshot` into a TUI row; no independent progress manager writes over the screen. `color` supplies the accent color, while `ansi` carries structured styles between progress, Markdown and TUI buffers.

The worker receives scan requests through a capacity-one channel. It walks in sorted preorder, sends batches of at most 32 entries through a capacity-16 channel, and never mutates a batch after publishing it. The event-loop thread owns all widget and viewer state. Event sends and request waits observe the task cancellation token, so shutdown also releases a worker blocked by channel backpressure. The scope joins the worker before the terminal is restored. Frame writes have a two-second timeout; a failed frame unwinds through the same resource cleanup path.

Notifications watch **only the root directory and its immediate entries**. Modifying the root README, creating/deleting an immediate entry, or renaming an immediate directory requests a rescan. Deep file modifications are intentionally outside this watch; press `r` to refresh those. Build/VCS names are excluded from both scans and root notices. Bursts are coalesced into at most one pending refresh while scanning. Scans preserve the selected path when it still exists, and reload its preview after completion. The status row shows scan and notification-batch counts. The counter is not an audit log of individual filesystem operations.

A failed watch is visible in the status row and leaves manual refresh usable. A removed or moved watched root ends automatic watching; this example does not automatically attach to a replacement root. Traversal warnings are reported without discarding readable branches. Preview size, encoding and parsing errors appear in the preview pane. Session/watch cleanup is explicit; normal cleanup failures are returned, and an earlier operation error remains the primary error during unwinding.

## Bounds and filesystem behavior

`Limits::defaults()` uses 1,024 entries including the root, depth 8, and 256 KiB per preview. Public helpers accept checked custom limits up to 4,096 entries, depth 32 and 1 MiB preview size. Scans retain at most 16 error messages and reject tree identifiers above 4,096 bytes. Hidden build directories `.git`, `node_modules`, `_artifact`, `_bootstrap`, `stage0`, `stage2`, `stage3`, and `__pycache__` are pruned. Other dotfiles remain visible.

Preview reads open read-only with `O_NOFOLLOW`, `O_NONBLOCK` and `O_CLOEXEC`, check the opened descriptor is a regular file, and enforce the byte budget while reading, including when a file grows. Invalid UTF-8 is rejected. The application neither executes commands nor writes, renames, deletes or creates files. The verification tests create and remove their own temporary fixtures.

Traversal never follows directory symlinks, and interactive mode rejects a symlink root. This remains a local filesystem browser, not a confinement API: renamed ancestors and path-prefix replacement can change what a later pathname resolves to. `walkdir` collects/sorts sibling names before yielding them, so the displayed entry limit does not impose a hard bound on the underlying work or memory required to enumerate a single enormous directory. Cancellation is checked between iterator steps and batches; it cannot interrupt a filesystem syscall blocked in the kernel. Preview reads are bounded but synchronous on the UI thread, so a slow filesystem can delay interaction. Use ordinary local directories for this example.

## Verification

`verify.py` checks formatting, five application tests, normal and cached builds, repeatable built-in snapshots, real PTY interaction, and the same tests under Go's race detector. It stores logs and a report under `_artifact/verification/`.

The PTY test uses an isolated temporary tree and a small screen model. It verifies background scan completion, ignored directories, root-file modification/create/delete notifications, selection, Markdown scrolling, manual refresh, resize, read-only behavior and restoration of termios, descriptor flags, cursor and alternate-screen state. The application tests also cover entry/depth bounds, symlink pruning, invalid UTF-8, oversized previews, replacement by a symlink, selection retention, missing roots, and cancellation while the worker is blocked on a full result channel.
