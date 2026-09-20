# terminal

`ecosystem::terminal` is a Linux amd64 terminal backend written in GoML. It owns terminal modes, serializes input and output across copied session handles, decodes incremental input into typed events, and keeps cancellation and cleanup explicit. It has no Go adapter or third-party dependency.

```toml
[dependencies]
"ecosystem::terminal" = "0.1.0"
```

```goml
use ecosystem::terminal;
use std::context;

fn run() -> Result[(), terminal::Error] {
    terminal::with_session(terminal::Options::interactive(), |session| {
        session.write("Press q to quit.\r\n")?;
        loop {
            if let Some(event) = session.next_event(context::Context::background(), -1)? {
                match event {
                    terminal::Event::Key(key) => {
                        if key.code == terminal::KeyCode::Char('q') {
                            return Result::Ok(())
                        }
                    },
                    terminal::Event::End => return Result::Ok(()),
                    _ => (),
                }
            }
        }
    })
}
```

## Session and lifecycle

- `Session::open(options)` reopens terminal stdin/stdout through `/proc/self/fd/{0,1}` with independent open descriptions, leaving their standard-stream `O_NONBLOCK` flags unchanged. Redirected streams are duplicated to preserve file offsets and append flags; they share status flags temporarily and require the same exclusive-use contract as `from_fds`.
- `Session::from_fds(input, output, options)` duplicates caller descriptors. Caller descriptors remain open. Duplicates share status flags, so the session temporarily enables nonblocking I/O and restores both original flag sets on close. The caller must give the session exclusive use of those open descriptions until close.
- Copied `Session` values share input ownership, an output gate, and close state. Only one `next_event` reads at a time, and each complete `write` holds the output gate, including partial writes. `size` returns a `Size { columns, rows }` using `TIOCGWINSZ`.
- `Options::defaults()` leaves terminal modes unchanged, allowing pipe-based sessions. `Options::interactive()` enables raw input, alternate screen, hidden cursor, and bracketed paste. Mouse and focus reporting are optional. Raw mode disables echo, canonical input, signal-generating input characters, and output postprocessing; applications should use `\r\n` when they need a new line at column zero.
- `close()` wakes active and queued readers/writers, disables enabled reporting modes, restores cursor and alternate screen, restores the saved kernel termios and descriptor flags, and closes retained descriptors. All restoration steps run even if an earlier one fails. Repeated or concurrent close returns the same result. Setup failures also attempt complete restoration.
- `with_session(options, body)` closes the session after the callback returns `Ok` or `Err`. It preserves a callback error and reports cleanup errors when the callback succeeds. Manual sessions can use `defer { let _ = session.close(); };` or explicitly handle the close result.

A terminal should have one application-owned session. Independent sessions controlling the same terminal cannot coordinate saved mode state; nested sessions and external writes or reads are outside the ownership contract. Output modes are disabled on close rather than queried and restored to an unknown prior application state. Kernel termios and flags are restored exactly. Cleanup is guaranteed on ordinary returns and errors, not process exit, an unhandled signal, or a runtime panic. There is no global handler, goroutine, or automatic finalizer.

## Events and cancellation

`next_event(context, timeout_ms)` returns `Result[Option[Event], Error]`. `-1` waits indefinitely, `0` polls without waiting, and a positive timeout includes waiting for input ownership. `None` means timeout or an already-delivered EOF. `Event::End` is delivered exactly once after valid buffered input. Invalid UTF-8 and truncated paste at EOF are errors.

`Context` cancellation yields `io::ErrorKind::Interrupted`; context deadlines yield `TimedOut`. Session close yields `BrokenPipe`. The implementation uses nonblocking descriptors and poll waits of at most 20 ms, so cancellation and close do not depend on input arriving. An input timeout preserves partial UTF-8, escape sequences, and paste. Resize changes are detected by checking the output size during the same bounded polling loop; changes may coalesce and do not require installing a signal handler.

`write(text)` and `write_with(context, text)` write directly, so there is no buffered flush requirement. Writes are performed in chunks of at most 64 KiB. A failed or cancelled write may have written a prefix; callers must not blindly replay the whole value. Setup and cleanup control output have a 200 ms budget, allowing termios restoration even when output is blocked. No arbitrary terminal control bytes are added to ordinary `write` calls.

Supported events:

| Event | Behavior |
| --- | --- |
| `Key(KeyEvent)` | Unicode character, Enter, Escape, Backspace, Tab/BackTab, arrows, Home/End, Insert/Delete, PageUp/PageDown, F1–F12; Shift/Alt/Control modifiers |
| `Mouse(MouseEvent)` | SGR mouse press/release, drag/move, vertical/horizontal wheels, modifiers; zero-based coordinates |
| `Paste(string)` | Bracketed paste preserved as one text event, including embedded escapes and newlines |
| `Focus(bool)` | Xterm focus-in/focus-out |
| `Resize(Size)` | Changed nonzero row/column dimensions |
| `Unknown(string)` | Unsupported CSI/SS3 and OSC/DCS/APC/PM replies kept as one event; undecodable/incomplete UTF-8 replies use a descriptive placeholder |
| `End` | Input EOF, delivered once |

Alphabetic CSI keys accept no parameters, an explicit first parameter of `1`, or `1;modifier` with modifiers `1..8`. Bare SS3 cursor/Home/End/F1–F4 keys and the legacy CSI `11~` through `14~` F1–F4 forms remain supported; modified function keys use CSI. Non-key parameter shapes such as CSI `999A` and the cursor-position reply CSI `2;3R` stay `Unknown`. CSI `1;modifierR` is intrinsically ambiguous with a cursor-position reply on row one; this decoder interprets it as F3 with modifiers. Applications querying cursor position must account for that legacy ambiguity. This matches the overlapping encodings documented by [Xterm](https://invisible-island.net/xterm/ctlseqs/ctlseqs.html).

The decoder supports standard xterm CSI/SS3 keys and SGR mouse. It does not negotiate Kitty keyboard, modifyOtherKeys, key-release events, legacy X10 mouse, terminal-specific terminfo keys, or IME composition. A terminal may encode multiple physical keys identically; the library does not invent missing distinctions. Mouse reporting enables button-motion mode, not all-motion mode. Ctrl-C is a typed control key while raw mode is enabled; the application decides how to handle it.

## Pure decoding and bounds

`Decoder::new(DecoderOptions::defaults())`, `feed(Slice[byte])`, `pending_escape()`, `flush_escape()`, and `finish()` allow custom transports and deterministic tests. Arbitrary byte fragmentation, including inside a UTF-8 scalar and paste delimiter, is supported. `flush_escape` emits a lone Escape key or an `Unknown` incomplete escape sequence. Session applies it after `Options.escape_timeout_ms` (default 30 ms); this ambiguity timeout also bounds incomplete terminal replies. Pure decoder callers own that clock. Decoder copies share state and require one owner; Session provides synchronization.

Escape/control sequences default to a 256-byte maximum and accept limits of 8–65,536 bytes. Paste defaults to 1 MiB and permits 0–16 MiB, measured in UTF-8 bytes, excluding delimiters. A paste may retain six additional delimiter bytes while incomplete. Invalid UTF-8 in a key or completed paste, or an exceeded bound, permanently fails that decoder, preventing a failed paste from being interpreted as keystrokes. `feed` returns events atomically for each call; an error discards that call's events. The session reads at most 4 KiB per batch. Total events returned by a direct `feed` scale with its caller-provided input.

## Capabilities and control helpers

`detect_capabilities(input_tty, output_tty, term, colorterm, no_color, policy)` is deterministic and independently testable. `Session.capabilities()` supplies the actual TTY checks and environment values. `ColorPolicy::Auto` disables color for redirected output, empty/`dumb` TERM, and a nonempty NO_COLOR; explicit `Always`/`Never` override automatic color policy. Truecolor is inferred from COLORTERM `truecolor`/`24bit` or a `direct` TERM, and 256-color support from `256color`. Hyperlinks are conservative name-based hints for xterm, kitty, and wezterm, not a negotiated guarantee.

`cursor_to(column, row)` validates zero-based coordinates and returns a cursor-position sequence. `clear_screen`, `clear_line`, `save_cursor`, and `restore_cursor` return control strings. Styling, color conversion, and display width belong in the sibling `ansi`, `color`, and `unicode_text` libraries.

Protocol and raw-mode references: [Xterm control sequences](https://invisible-island.net/xterm/ctlseqs/ctlseqs.html), [Linux termios](https://man7.org/linux/man-pages/man3/termios.3.html).

## Validation

Run `just ecosystem-test terminal` from the repository root. The native GoML verifier builds the consumer and exercises it under real Linux PTYs: fragmented keys/UTF-8/paste, ESC ambiguity, mouse/focus, ioctl resize, cancellation, ordinary error cleanup, exact termios/flag restoration, mode enable/disable output, a 4 MiB write with backpressure, and redirected-file offset/append preservation. Pipe tests cover EOF, raw-mode rejection, session aliases, queued/active reader and writer cancellation, serialization under backpressure, idempotent close, and flag restoration. The verifier also rebuilds and runs generated library and consumer tests with Go's race detector. No Python interpreter is required.
