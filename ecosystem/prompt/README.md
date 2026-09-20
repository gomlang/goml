# ecosystem::prompt

Typed terminal prompts implemented in GoML on `tui` and `terminal`. The same
models can be tested by feeding events into `Model.update` and drawing to an
in-memory `tui::Buffer`; interactive execution adds a cancellable event loop.

```gom
use ecosystem::prompt;
use std::context;

fn ask_port() -> Result[isize, prompt::Error] {
    prompt::interact(prompt::integer("Port", 1, 65535)?, context::Context::background())
}
```

## Models

- `TextInput[T]` parses and validates into any concrete result type through a
  `(string) -> Result[T, string]` callback. Invalid input stays editable and its
  validation message is displayed. `text`, `password`, and bounded `integer`
  constructors cover common cases.
- `InputOptions` controls initial text, history, help, maximum bytes, multiline
  editing, password masking and theme. Editing uses extended graphemes and
  supports selection, undo/redo, clipboard-style paste events and horizontal or
  vertical viewport scrolling through `tui::Editor`.
- `with_completion` supplies whole-value candidates; Tab cycles them. Up/Down
  or Ctrl-P/Ctrl-N traverse history and restore the current draft. Single-line
  paste converts line breaks to spaces. Multiline Enter inserts a newline;
  Alt-Enter submits (Ctrl-Enter also works when supplied by an event backend).
  Set a corresponding help string for multiline input.
- `Select[T]` returns a typed choice; `Selection[T]` also supports multiple
  choices, minimum/maximum cardinality, disabled entries, descriptions,
  Unicode case-folded substring search and paginated navigation. Hidden
  selections are retained and submitted in original choice order. Space
  toggles multiple selection; paste can enter a search query containing spaces.
- `Confirm` accepts Y/N, an optional default and cancellation. Enter with no
  default remains pending.

All models yield `Outcome::Pending`, `Submitted`, `Cancelled`, or `EndOfInput`.
Esc and Ctrl-C cancel. Ctrl-D on an empty text field indicates end of input.
Updating or submitting a finished model returns an error. Models are mutable
single-consumer handles; copies share state. Supplied choices, history and
completion vectors are copied, while arbitrary typed payloads retain their
ordinary GoML value/storage semantics. `Theme` customizes five visual roles.

## Execution and cleanup

`run(screen, model, context)` borrows the application's TUI session; the caller
closes it. `interact(model, context)` opens a session and restores terminal state
on success, cancellation, EOF, deadline and ordinary render/input/output errors.
Both action and cleanup failures are retained when they occur together. Render
callbacks and validators execute synchronously, must return cooperatively, and
must not reenter the same model.
No background thread or global terminal state is installed by this library.

The interactive backend currently supports Linux amd64. Redirected input/output
returns a terminal error; applications can implement a separate batch input
policy. Panic, process termination and external signals have the same cleanup
limitations as `terminal`; they are not converted into prompt outcomes.

Password mode masks rendered graphemes and disables completion, history and
undo/redo retention. It is a display policy: its value and parser result remain
ordinary garbage-collected strings, with no secure erasure guarantee. A custom
validator is responsible for keeping secrets out of its error messages.

## Limits and validation

Questions, descriptions, help and validation messages are bounded to 4 KiB;
terminal controls are made literal. Text input defaults to 64 KiB and permits
at most 1 MiB. History is bounded to 1,000 entries/4 MiB, completions to 256
entries/4 MiB, and choices to 10,000 entries/4 MiB. Search accepts up to 4 KiB;
page sizes are 1–100. Invalid options return `InvalidOptions`; typing or paste
beyond a field's bound leaves its prior value intact and displays a message.
User callbacks can allocate independently of these returned-value limits.

Run `python3 ecosystem/verify.py prompt` for model tests, a separate registry
consumer and real PTY tests. PTY coverage includes combining/emoji editing,
validation retry, password output inspection, searchable selection, cancellation,
EOF, malformed UTF-8, deadline, deliberate render failure and termios restoration.
