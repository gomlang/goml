<a id="releasing-goml"></a>

# 发布 goml

[English](releasing.md) | [简体中文](releasing.zh-CN.md)

发布使用严格的 `vX.Y.Z` 标签，提供 Linux amd64 和 macOS arm64 二进制文件。

Go 兼容性基线为 Go 1.26。请使用 Go 1.26.x 构建和验证发布；用户需要 Go 1.26 或更高版本来编译生成的程序，包括其中的显式外部调用。归档不捆绑 Go 工具链。

根目录的 [VERSION](../VERSION) 文件是版本的权威来源。`goml`、`gomlc`、`gomlfmt`、`gomldoc`、`gomllsp` 和 VS Code 扩展必须使用相同版本。官方归档包含这五个可执行文件，以及 builtin、prelude 和标准库资源。

<a id="version-policy"></a>

## 版本策略

每次发布必须从最近已发布版本连续推进一个 SemVer 步骤：

- 补丁版本：`0.1.0` 到 `0.1.1`
- 次版本：`0.1.0` 到 `0.2.0`
- 主版本：`0.1.0` 到 `1.0.0`

不允许跳过版本，例如从 `0.1.0` 跳到 `0.1.2` 或 `0.3.0`。

在早期自举阶段，只允许发布连续的 `0.1.x` 补丁版本。

<a id="publish"></a>

## 发布

在仓库根目录运行，准备[开发指南](development.zh-CN.md)中的前置条件，以及用于发布操作且已认证的 GitHub CLI（`gh`）。将 `X.Y.Z` 替换为最近已发布版本之后的下一个连续版本：

```sh
release_version=X.Y.Z
just set-version "$release_version"
just ci
```

`just set-version` 同步 `VERSION`、两个 GoML 版本模块，以及 VS Code 包和锁文件版本。提交并推送变更，等待该确切提交的 main CI 通过后再创建标签。在同一 shell 中保持 `release_version` 已设置：

```sh
(
    set -eu
    release_sha="$(git rev-parse HEAD)"
    git push origin main
    ci_run_id="$(gh run list --repo gomlang/goml --workflow CI --branch main --commit "$release_sha" --event push --limit 1 --json databaseId --jq '.[0].databaseId // empty')"
    test -n "$ci_run_id"
    gh run watch "$ci_run_id" --repo gomlang/goml --exit-status
    git tag -a "v$release_version" -m "goml v$release_version"
    git push origin "v$release_version"
)
```

若推送失败、CI 运行尚未登记或 CI 失败，上述代码块会在打标签前停止。若暂时看不到 CI 运行，请等待其出现后重新运行该代码块。

main CI 直接使用固定的 Linux amd64 stage0 构建 stage2。Linux 检查覆盖编译器、驱动器、标准库和脚本测试、扩展编译和 VSIX 打包、归档冒烟测试，以及 stage3 不动点验证。不动点验证比较 stage2 构建的编译器和驱动器产物与使用 stage3 重建的产物。macOS 15 arm64 任务使用 Linux 构建导出的生成 Go 源码构建并测试发布归档，再独立地从固定的 macOS arm64 stage0 自举、验证原生不动点，并运行工具链脚本测试。该任务不接收 Linux 可执行文件或编译器 world。两个平台都必须通过最终的 `test` 任务。独立的 [gomlgo 项目](https://github.com/gomlang/gomlgo) 维护自己的独立测试套件。

Release 工作流要求标签指向的确切提交已通过 main CI，并验证版本、前一个发布版本和 stage0。它在 Linux 上重建 stage2，导出五个生成的 Go 入口，然后在 Ubuntu amd64 和 macOS arm64 上并行原生构建全部五个工具。每个任务验证二进制目标元数据、打包资源，并运行归档位置迁移、初始化、项目命令、格式化、文档、LSP、直接 Go FFI 和 Go 导出冒烟测试。macOS 还检查对 Linux 系统调用的拒绝。Release CI 不重复完整 Linux 测试套件或不动点验证。

每个平台上传归档及其专属 `.sha256` 文件。发布任务等待两个平台完成，检查两个摘要，生成恰好包含两个归档的 `SHA256SUMS`，并在正式发布前将它们上传到草稿发布。只有该任务具有发布写权限。已发布版本不会被覆盖。

在 Linux 上运行 `just make-tools` 后，可复现打包：

```sh
bash tools/release/package.sh "$(cat VERSION)" linux amd64
bash tools/release/smoke.sh "$(cat VERSION)" linux amd64
bash tools/release/sources.sh _artifact/release-sources.tar.gz
```

将源码归档传到 macOS arm64 上相同提交的检出目录，确保 Go 1.26.x、Bash、tar、jq 和 C 编译器可用，然后运行：

```sh
tar -xzf release-sources.tar.gz
bash tools/release/package.sh "$(cat VERSION)" darwin arm64
bash tools/release/smoke.sh "$(cat VERSION)" darwin arm64
```

打包脚本也支持指定目标进行交叉编译；冒烟测试要求在目标宿主运行。`sha256sum` 不可用时，脚本使用 `shasum -a 256`。两个归档及其 `.sha256` 文件都位于 `dist` 后，运行 `bash tools/release/checksums.sh "$(cat VERSION)"` 生成 `dist/SHA256SUMS`。

<a id="compiler-world-loading"></a>

## 编译器 world 加载

工具链初始化写入并完整验证 `lib/compiler/compiler-world.gaf`。world 包含内建初始化元数据，以及带索引的包接口和 Core。编译器将自身可执行文件相对安装路径下的 world 视为不可变工具链资源：正常加载检查其结构、格式、编译器标识和依赖指纹，不重新计算文件校验和。驱动器将已安装 world 的大小、修改时间和编译器指纹一起用于缓存新鲜度判断，不在每条命令中重新哈希 world。从其他路径提供的 world 和初始化验证仍执行完整校验和验证。发布归档和 stage0 下载仍进行 SHA-256 验证。

包产物存储 12 字节的索引节点，在每个包产物内部共享相同子树，包括重复类型。引用仅指向之前已完成的子树；加载时验证边界并限制值展开。只支持当前产物模式。格式变化后，重建项目产物并重新运行工具链初始化。每个自举阶段读取由自身工具链生成的产物。

请通过安装或初始化替换工具链资源，不要原地修改已安装 world。正常编译不保证能发现已安装 world 的损坏。保持其大小和修改时间的更改可能保留过期的驱动器缓存。包检查加载接口、类型和编译期元数据；运行时 Core 在链接时加载。

<a id="installation"></a>

## 安装

安装 Go 1.26 或更新版本，并使其在 `PATH` 中可用。独立 POSIX shell 安装器维护在 [gomlang/install](https://github.com/gomlang/install)，需要 curl、tar，以及 sha256sum 或 shasum：

```sh
curl -fsSL https://gomlang.github.io/install/install.sh | sh
export PATH="${GOML_HOME:-$HOME/.goml}/bin:$PATH"
goml version
```

安装器识别 Linux amd64 或 macOS arm64，下载最新已发布版本及其 `SHA256SUMS`，验证归档并初始化编译器 world。它安装到 `${GOML_HOME:-$HOME/.goml}`，打印可加入 shell 配置的 PATH 命令。它不修改 shell 配置文件，也不要求 `just` 或源码检出目录。

选择版本或安装目录：

```sh
curl -fsSL https://gomlang.github.io/install/install.sh | sh -s -- --version X.Y.Z --prefix "$HOME/tools/goml"
```

再次运行安装器即可升级。它替换五个 GoML 可执行文件和 `builtin`、`prelude`、`std`、`compiler` 资源目录，保留依赖缓存和其他文件。替换或初始化失败时恢复原工具链。拒绝向同一前缀并发安装；强制终止后，请在重试前检查报告的 `.goml-install` 目录和所有备份。安装器拒绝不支持的宿主、低于 1.26 的 Go 版本，以及覆盖 GoML 源码检出目录的安装。

手动安装请按下列步骤操作。

从同一个发布版本下载适合宿主的归档和 `SHA256SUMS`：

| 宿主 | 归档 | 原生发布验证环境 |
| --- | --- | --- |
| Linux amd64 | `goml-X.Y.Z-linux-amd64.tar.gz` | Ubuntu 24.04 |
| macOS arm64（Apple Silicon） | `goml-X.Y.Z-darwin-arm64.tar.gz` | macOS 15 |

例如，在 macOS 上，将 `X.Y.Z` 替换为发布版本，并在解压前验证下载的归档：

```sh
release_version=X.Y.Z
archive="goml-$release_version-darwin-arm64.tar.gz"
awk -v archive="$archive" '$2 == archive' SHA256SUMS > archive.sha256
test "$(wc -l < archive.sha256 | tr -d ' ')" = 1
shasum -a 256 -c archive.sha256
tar -xzf "$archive"
prefix="$PWD/goml-$release_version-darwin-arm64"
"$prefix/bin/goml" __toolchain-finalize --prefix "$prefix"
export PATH="$prefix/bin:$PATH"
goml version
```

Linux 用户选择 `linux-amd64`，可使用 `sha256sum -c archive.sha256`。迁移安装位置时请保持 `bin` 和 `lib` 在一起。构建程序要求 `PATH` 中存在 Go 1.26 或更新版本。用户提供的 cgo 适配器还需要 C 编译器。

macOS 支持编译器、驱动器、格式化器、文档生成器、LSP、Go FFI、用户提供的 cgo 适配器，以及通过 `just bootstrap` 进行原生源码自举。自举使用已发布的 macOS arm64 归档及 `bootstrap/stage0.env` 中固定的校验和。`std::os::linux`、`std::net` 中当前基于 Linux 系统调用的套接字操作及依赖它们的网络 API 仍仅支持 Linux amd64。SIMD 在 arm64 上使用标量实现。

发布归档采用完整的工具链前缀：

```text
goml-X.Y.Z-<os>-<arch>/
├── bin/
│   ├── goml
│   ├── gomlc
│   ├── gomlfmt
│   ├── gomldoc
│   └── gomllsp
└── lib/
    ├── builtin/
    │   ├── goml.toml
    │   └── ...
    ├── prelude/
    │   ├── goml.toml
    │   └── prelude.goml
    └── std/
        ├── goml.toml
        └── ...
```

编译器相对于其可执行文件解析 `lib`。归档必须严格保留此布局。`builtin`、`prelude` 和 `std` 是独立的 GoML 项目；不得打包平铺的 `builtin_*.goml` 文件。

发布归档包含工具链项目源码和清单，但不包含 `lib/compiler/compiler-world.gaf`。上述安装命令使用同一归档中的二进制文件进行一次初始化。

初始化在临时路径构建并验证编译器 world，再将其原子安装到 `lib/compiler/compiler-world.gaf`。普通的 `goml check`、`goml build` 和 `goml test` 命令只读取相对于可执行文件的 world。缺少 world 表示安装不完整，不得隐式触发源码构建。

<a id="advance-stage0"></a>

## 推进 stage0

发布使用前一个 Linux amd64 发布版本作为 stage0 构建；原生 macOS 自举使用同一版本的 macOS arm64 归档。发布后，保持 `release_version` 为已发布版本，从 `SHA256SUMS` 取得两个平台的校验和，然后推进 stage0：

```sh
linux_sha256=LINUX_AMD64_SHA256_FROM_SHA256SUMS
darwin_sha256=DARWIN_ARM64_SHA256_FROM_SHA256SUMS
just set-bootstrap-stage0 "$release_version" "$linux_sha256" "$darwin_sha256"
just bootstrap
```

使用前一个 stage0 无法编译的语言特性之前，先提交更新后的 `bootstrap/stage0.env`。下一个发布工作流要求 stage0 与最近已发布版本一致。

离线自举时，下载宿主对应的固定版本归档，然后运行，例如在 macOS arm64 上：

```sh
GOML_STAGE0_ARCHIVE=/path/to/goml-X.Y.Z-darwin-arm64.tar.gz just bootstrap
```
