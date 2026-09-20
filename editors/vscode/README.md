# GoML VS Code extension

The extension provides syntax highlighting, diagnostics, hover, completion, go-to-definition, signature help, inlay hints, formatting, test code lenses, and quick fixes.

The status bar shows whether the language server is starting, ready, stopped, or busy. While it is busy, it shows the current server operation and elapsed time. Hover over the status item to see queued client requests, or click it to open the language server output. Operations taking at least two seconds are recorded there as warnings.

Diagnostics run after editing pauses, while completion reuses the latest checked types immediately. Saving requests diagnostics without waiting for the edit delay. Package checks reuse unchanged dependencies across edits. External Go types and concrete generic instances use the bundled `goml-go-meta` helper; Go 1.26 or newer must be installed for these checks. Changes to Go sources and module files invalidate analysis alongside GoML changes.

From the repository root, build the self-hosted language server and extension:

```sh
just vscode-ext
```

Open `editors/vscode` as the VS Code workspace and press F5 to launch the Extension Development Host. The checked-in launch task uses `pnpm run compile`, so debugging also requires pnpm; the `just` build and packaging recipes use npm.

Use **GoML: Show Expanded Derive** to inspect the current document's AST after derive expansion. **GoML: Show Language Server Output** opens the server log. Test code lenses save the source and invoke `goml test` with the selected test name and kind; the project driver must be on `PATH`, and the workspace folder must be inside the GoML module containing the test.

Use **Format Document** to format the current GoML buffer. The language server uses GoML's fixed formatting rules and leaves syntactically invalid documents unchanged.

The bundled server is built for the host platform. A custom server needs its complete executable-relative toolchain resources and an adjacent `goml-go-meta` helper for external Go type checks. See the [language guide](../../docs/goml.md#lsp-and-editor) and [formatting rules](../../docs/formatting.md) for details.

Configuration:

- `goml.serverPath` overrides the bundled or `PATH`-resolved `gomllsp`.
- `goml.trace.server` controls language-server tracing.

Package a `.vsix` with:

```sh
just package-vscode-ext
```

The package is written to `editors/vscode/goml-<version>.vsix`. Packaging resolves README links relative to `editors/vscode` on the repository's `main` branch. CI also packages the extension to validate this step.
