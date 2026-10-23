# GoML

[English](README.md) | [简体中文](README.zh-CN.md)

GoML 是一门受 Go 和 Rust 启发的静态类型编程语言。

GoML 中的“ML”致意于 [ML 编程语言](https://en.wikipedia.org/wiki/ML_(programming_language))，其后继语言对 Rust 产生了深远影响。

GoML 希望让 Go 开发者在继续使用 Go 生态的同时，获得更强大的类型系统。

GoML 采用静态类型和垃圾回收，具有类似 Rust 的语法和单态化泛型，没有所有权或生命周期系统。编译器、项目驱动程序、测试、格式化工具和语言服务器均由 GoML 实现，并编译为 Go。

<a id="get-started"></a>

## 快速开始

为 Linux amd64 或 macOS arm64 [安装工具链](docs/releasing.zh-CN.md#installation)。先安装 Go 1.26 或更新版本，然后运行：

```sh
curl -fsSL https://gomlang.github.io/install/install.sh | sh
export PATH="${GOML_HOME:-$HOME/.goml}/bin:$PATH"
```

[安装器](https://github.com/gomlang/install)会下载最新发布版本，验证校验和并初始化工具链。将其输出的 PATH 命令加入 shell 配置。

```sh
goml new hello
cd hello
goml run
```

使用 `goml check`、`goml fmt` 和 `goml test` 检查、格式化和测试项目。

<a id="documentation"></a>

## 文档

- [语言指南](docs/goml.zh-CN.md)：语法、类型和包
- [标准库](docs/stdlib.zh-CN.md)：包、API 和常见任务
- [Go 互操作](docs/ffi/bind-go.zh-CN.md)：在 GoML 中使用 Go 包
- [VS Code 扩展](editors/vscode/README.zh-CN.md)：编辑器配置
- [开发指南](docs/development.zh-CN.md)：构建和测试自托管编译器及工具

<a id="status-and-license"></a>

## 项目状态与许可

GoML 是一个实验性的个人项目，不隶属于任何组织。请勿将其或其衍生作品用于生产环境。作者不承担由此产生的风险或损害责任。

目前未授予任何开源许可。保留所有权利。复制、分发、修改或商业使用均须事先获得作者授权。
