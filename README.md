# GoML

[English](README.md) | [简体中文](README.zh-CN.md)

GoML is a statically typed programming language inspired by Go and Rust.

The "ML" in GoML nods to the [ML (programming language)](https://en.wikipedia.org/wiki/ML_(programming_language)), whose descendants have deeply influenced Rust.

GoML aims to empower gophers with a more powerful type system but without leaving the Go ecosystem.

GoML is statically typed and garbage-collected, with Rust-like syntax, monomorphized generics, and no ownership or lifetime system. The compiler, project driver, tests, formatter, and language server are implemented in GoML and compile to Go.

## Get started

[Install the toolchain](docs/releasing.md#installation) for Linux amd64 or macOS arm64. Install Go 1.26 or newer first, then run:

```sh
curl -fsSL https://gomlang.github.io/install/install.sh | sh
export PATH="${GOML_HOME:-$HOME/.goml}/bin:$PATH"
```

The [installer](https://github.com/gomlang/install) downloads the latest release, verifies its checksum, and initializes the toolchain. Add the printed PATH command to your shell configuration.

```sh
goml new hello
cd hello
goml run
```

Use `goml check`, `goml fmt`, and `goml test` to check, format, and test your project.

## Documentation

- [Language guide](docs/goml.md): syntax, types, and packages
- [Standard library](docs/stdlib.md): packages, APIs, and common tasks
- [Go interoperability](docs/ffi/bind-go.md): use Go packages from GoML
- [VS Code extension](editors/vscode/README.md): editor setup
- [Development](docs/development.md): build and test the self-hosted compiler and tools

## Status and license

GoML is an experimental personal project, unaffiliated with any organization. Do not use it or its derivatives in production. The author assumes no responsibility for resulting risk or damage.

No open-source license is currently granted. All rights reserved. Copying, distribution, modification, or commercial use requires the author's prior authorization.
