<a id="library-boundaries"></a>

# 库的边界

[English](library-boundaries.md) | [简体中文](library-boundaries.zh-CN.md)

`lib/std` 包含可广泛复用的值类型、算法和组合契约。其算法以 GoML 实现，仅依赖 builtin、prelude 和其他标准包。宿主操作使用编译器拥有的内建运行时钩子；标准包不依赖用户 Go FFI 或元数据发现。协议引擎、应用策略、大型数据集和领域专属类型位于 `../gomlang/` 下独立版本管理的同级仓库。

引入新的归属模块前，优先扩展现有包。当真实消费者及其迁移测试需要时，再提取共享模块。生态模块之间共享实现并不意味着必须将其移入标准库。

<a id="ownership"></a>

## 归属

| 标准库 | 生态 |
| --- | --- |
| `std::bytes`, `std::text`, `std::unicode`, `std::utf8`, `std::num`, `std::collections`, `std::error` | `regexp`, `parser::scanner`, `tabwriter`, `unicode_text`, `template` |
| `std::io`, `std::fs`, `std::testing::fs`, `std::testing::io` | `archive`, `compress`, `object`, `dwarf` |
| `std::encoding`, `std::bytes::endian`, `std::hash`, `std::crypto` | `asn1`, `xml`, `html`, `mime`, `mail`, `x509::pkix` |
| `std::net` 地址值、`std::net::url`、`std::path::slash` | `http`, `textproto`, `request`, `web` |
| `std::math`, `std::rand` | `bigint`, `bigmath`, `datetime`, `sql`, `tracing`, `image`, `color` |
| 已有的 `std::json`, `std::toml`, `std::bincode` | 其他格式引擎和 `go_doc` |

标准 Unicode 标量分类和大小写处理共用一份固定版本的数据集。字素布局和终端宽度属于 `unicode_text`；时区和公共后缀数据集仍由生态维护。通用 URL 和 IP 值不携带 HTTP 客户端策略。词法斜杠路径与宿主文件路径操作保持分离。

公共 API 和限制记录在[标准库参考](stdlib.zh-CN.md)及各生态项目的 README 中。归属表不代表 Go API 兼容性承诺。Go 源码注释工具属于 `go_doc`，独立于 GoML 文档和独立的 [gomlgo 前端](https://github.com/gomlang/gomlgo)。

<a id="dependency-direction"></a>

## 依赖方向

- `archive` 使用 `compress` 和标准 I/O、校验和及字节序缓冲区。
- `image::png` 使用 `compress::zlib` 和标准 CRC32；`image` 使用 `color`。
- `request` 和 `web` 共用 URL 值、MIME、文本分帧和 HTTP 辅助工具。`web::proxy` 可以使用 `request`；`request` 不得依赖 `web`。
- `template` 和 `markdown` 共用 `html`；上下文转义保留在 `template` 中。
- `regexp` 使用 Unicode 和 UTF-8，但不引入词法分析器专用匹配规则。
- `text` 可以使用 `unicode`；Unicode 输出组装不得依赖 `text`。
- `x509::pkix` 使用 `asn1`；它不实现 TLS 或证书验证。
- HMAC 使用摘要算法；HKDF 和 PBKDF2 使用 HMAC。
- `sqlite` 适配 `sql` 契约；`sql` 不得依赖 `sqlite`。
- `object` 和 `dwarf` 使用字节和随机访问 I/O 接口。

<a id="compatibility"></a>

## 兼容性

共享实现必须保留消费者的策略。仅用于 HTTP 的请求 URL 契约与通用 URL 引用解析保持区别。模板和 Markdown 转义保留各自的引号和实体规则。词法路径清理不是归档解压策略，HTML 实体转义也不是内容净化。

纯 GoML 算法可以使用已有的分配、标量数学、I/O 和运行时边界。保持这些边界显式；移动算法不意味着新增运行时、传输或密码学安全保证。

提取实现时，保留公共入口，并测试已有调用方和独立消费者。标准库改动还需要覆盖源码目录、依赖、导航和打包。编译器及驱动器源码只能在发布并推进 stage0 后采用新 API；见[仓库指南](../AGENTS.zh-CN.md#bootstrap-and-language-evolution)。

标准库改动使用[仓库检查](development.zh-CN.md)，生态模块使用同级生态验证仓库。发布遵循[发布指南](releasing.zh-CN.md)。
