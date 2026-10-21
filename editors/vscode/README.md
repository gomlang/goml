# GoML VS Code extension

[English](README.md) | [简体中文](README.zh-CN.md)

The extension provides syntax highlighting, diagnostics, hover, completion, go-to-definition, signature help, inlay hints, formatting, test code lenses, and quick fixes for `.goml` files.

Type inlay hints use short type names and a 25-character budget. Long types use `…` while preserving their outer structure; hover over a hint to see its full type, including package paths.

Unused `Result` and invalid `?` diagnostics expose stable codes and handling suggestions. Propagation errors link to the enclosing return type. The unused-Result quick fix makes ignoring the error explicit with `let _ = ...`; review that choice before applying it. Fixes from an older document version are discarded.

The status bar shows whether the language server is starting, ready, stopped, or busy. While it is busy, it shows the current server operation and elapsed time. Hover over the status item to see queued client requests, or click it to open the language server output. Operations taking at least two seconds are recorded there as warnings.

Diagnostics run after editing pauses, while completion reuses the latest checked types immediately. Saving requests diagnostics without waiting for the edit delay. Package checks reuse unchanged dependencies and prepared source ASTs across edits. Expanded derive text is formatted only when requested. Go FFI analysis uses the signatures declared in GoML and treats external Go types as opaque. Native symbol, method and generic-constraint errors are reported by the final Go build.

Source file notifications invalidate documents that use the affected package while retaining unchanged package checks. Unrelated source files, build outputs, and native Go file changes preserve cached hover results. Changes to tracked module and workspace configuration refresh project resolution.

Enum completions include payload types and editable placeholders for tuple and named fields. In `match` patterns, suggestions follow the matched enum and insert an arm body when needed. Existing payloads and arm bodies are preserved, and unit variants already covered by an unguarded arm are omitted. Contextual constructors such as `Some` also use the expected type.

For syntax highlighting without a language server, set `"goml.lsp.enabled": false`. This works without an installed toolchain. Changing the setting stops or starts the server immediately; disabling it clears server diagnostics and hides the server status item. Syntax highlighting and bracket matching remain available.

From the repository root, build the self-hosted language server and extension:

```sh
just vscode-ext
```

Open `editors/vscode` as the VS Code workspace and press F5 to launch the Extension Development Host. The launch task runs `npm run compile`.

Use **GoML: Show Expanded Derive** to inspect the current document's AST after derive expansion. **GoML: Show Language Server Output** opens the server log. Test code lenses save the source and invoke `goml test` with the selected test name and kind from the source file's directory. Set `goml.driverPath` or put the project driver on `PATH`; module discovery starts at the test file, including when a workspace contains several modules. A failed save cancels the test run.

Use **Format Document** to format the current GoML buffer. The language server uses GoML's fixed formatting rules and leaves syntactically invalid documents unchanged.

The bundled server is built for the host platform. A custom server needs its complete executable-relative toolchain resources, a matching `goml` project driver. See the [language guide](../../docs/tooling.md#lsp-and-editor) and [formatting rules](../../docs/formatting.md) for details.

Configuration:

- `goml.lsp.enabled` enables the language server (default: `true`); disable it for syntax highlighting only.
- `goml.driverPath` selects the matching project driver used for module and dependency resolution.
- `goml.serverPath` overrides the bundled or `PATH`-resolved `gomllsp`.
- `goml.trace.server` controls language-server tracing.

Package a `.vsix` with:

```sh
just package-vscode-ext
```

The package is written to `editors/vscode/goml-<version>.vsix`. Packaging resolves README links relative to `editors/vscode` on the repository's `main` branch. CI also packages the extension to validate this step.
