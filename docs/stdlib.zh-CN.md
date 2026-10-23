<a id="goml-standard-library"></a>

# GoML 标准库

[English](stdlib.md) | [简体中文](stdlib.zh-CN.md)

本参考文档介绍本仓库已实现的标准库 API。语法和类型规则见[语言指南](goml.zh-CN.md)，自动进入作用域的 API 见其中的 [prelude 参考](language/prelude.zh-CN.md#built-in-prelude)，标准库与生态包的划分见[库的边界](library-boundaries.zh-CN.md)。

快速导航：[包](#packages)、[文本](#text-and-source-positions)、[集合](#slice-algorithms-and-checked-vector-edits)、[I/O](#composable-io)、[文件系统](#portable-read-only-filesystems)、[网络](#networking)、[Serde 和 JSON](#serde-and-json)、[并发](#structured-concurrency)以及[常见任务](#common-tasks)。

<a id="packages"></a>

## 包

标准库独立于隐藏的 builtin 契约和自动进入作用域的 prelude。算法由 GoML 实现；宿主操作使用编译器拥有的 builtin 运行时钩子，不使用用户 Go FFI 声明，也不进行元数据发现。标准库的包必须显式导入：

```goml
use std::ascii;
use std::bincode;
use std::bytes;
use std::bytes::endian;
use std::c;
use std::channel;
use std::cmp;
use std::collections;
use std::collections::imm;
use std::context;
use std::crypto;
use std::encoding::base32;
use std::encoding::base64;
use std::encoding::hex;
use std::error;
use std::env;
use std::ffi;
use std::fs;
use std::hash;
use std::hash::adler32;
use std::hash::crc32;
use std::hash::crc64;
use std::hash::fnv;
use std::io;
use std::iter;
use std::json;
use std::math;
use std::math::bits;
use std::math::complex;
use std::net;
use std::net::tls;
use std::num;
use std::os::linux::syscall;
use std::path;
use std::path::slash;
use std::panic;
use std::process;
use std::rand;
use std::serde;
use std::task;
use std::testing;
use std::testing::io as iotest;
use std::text;
use std::toml;
use std::time;
use std::utf8;
use std::utf16;
use std::unicode;
```

公共 API 包括：

- `ascii::is_ascii`、字符类别谓词、ASCII 大小写转换和比较，以及 `escape_default`
- `bincode::standard`、`legacy`、配置构建器、serde `Serialize` 和 `Deserialize` 再导出、`encode_to_vec` 和 `decode_from_slice`
- `bytes::Bytes`、`bytes::FrozenBytes`、`bytes::Builder`、显式字节复制、不可变快照、带检查的零拷贝字节视图，以及 `bytes::endian::{Builder, Reader, Writer, Endian}` 和带检查的整数、浮点数读写
- `c::CString`、`literal` 和 `check_error`，用于显式的 C 兼容字符串和错误边界
- `channel::Operation`、`Selection`、`select`、`try_select` 和 `try_select_priority`，用于数量在运行时确定的通道选择
- `cmp::Ordering`、`Ord`、`Reverse`、比较辅助函数，以及两个值的最小值、最大值和钳位操作。`Ordering` 是由 `cmp` 再导出的内建类型。
- `collections::Arena`、`BinaryHeap`、`BitSet`、`BTreeMap`、`BTreeSet`、`Deque`、`HashSet`、`IndexMap`、`IndexSet`、`IndexVec`、`Interner` 和 `Stack`；基于哈希的集合要求 `Hash + Eq`，树集合和堆使用 `cmp::Ord`
- `collections::imm::{HashMap, HashSet, Vector, List, BTreeMap, BTreeSet}`，提供共享内部结构、更新时返回新版本的不可变集合
- `collections::sort`、`stable_sort`、`binary_search`、`min`、`max` 及其接受比较器的变体。`Vec[T]` 上的排序、搜索、选择和去重方法是推荐形式。
- `collections::search`、`is_sorted`、`lower_bound`、`upper_bound`、`equal_range` 及其接受比较器的变体，用于基于索引的单调搜索、有序性检查和重复元素区间查询
- `collections::slice_*` 视图算法及 `vec_replace`、`vec_insert_slice`、`vec_delete_range`、`vec_delete_if`、`vec_grow`、`vec_append_iter`，用于带检查的区间编辑和迭代器组合
- `collections::{map_clone, map_copy, map_equal, map_equal_by, map_delete_if, map_all, map_keys, map_values, map_insert, map_collect}`，用于映射的浅拷贝、考虑键是否存在的比较、快照迭代和收集
- `context::{Context, CancelHandle, Deadline}` 以及限定作用域的 `with_cancel`、`with_timeout` 和 `with_deadline`
- `crypto::{digest_equal, verify_digest}` 全长度摘要比较、`crypto::hash` 纯增量 SHA-2、SHA-3 和 SHAKE、`crypto::hmac` 带密钥哈希、`crypto::hkdf` 和 `crypto::pbkdf2` 密钥派生，以及 `crypto::rand` 操作系统随机字节
- `encoding::hex` 小写和大写十六进制编码及带检查的解码
- `encoding::base32` RFC 4648 标准与扩展十六进制编码，带填充和无填充变体，以及带检查的规范解码
- `encoding::base64` RFC 4648 标准和 URL 安全编码，带填充和无填充变体
- `error::Error`、`ErrorKind`、`Details`、保留类型的 `Report[E]`、`ResultExt`、`ReportContext` 以及稳定的错误类别代码转换
- `env::args`、当前目录和可执行文件查询，以及环境变量读取
- `ffi::String`、`Rune`、`Ptr`、`Error`、`Func`、`RawSlice`、`RawMap` 和显式 Go 边界适配器
- `fs::read_file_structured`、`write_file_structured`、结构化字节 I/O、目录操作、路径检查和 `sha256_file`
- `hash::Hasher`，增量 Adler32、CRC32、CRC64 和 FNV-1/FNV-1a 校验和实现
- `io::{Read, Write, BufRead, Close}`、`Cursor`、`Take`、`BufReader`、`BufWriter`、`copy`、`stdin`、`stdout`、`stderr` 和现有标准流函数
- `iter::empty`、`once`、`from_fn`、迭代器适配器和单次遍历消费操作
- `json::Value`、`parse`、`encode`、serde `Serialize` 和 `Deserialize` 再导出、`to_value`、`from_value`、`try_to_string`、`from_string`、`field` 和带类型的 `as_*` 访问器
- `math` 的 f32/f64 初等函数、IEEE 754 分类，以及 `E`、`PI`、`TAU`、`SQRT_2`、`LN_2` 和 `LN_10` 常量
- `math::bits` 定宽位计数、反转、旋转和带检查的双倍宽度算术
- `net::{IpAddr, ScopedIpAddr, IpPrefix, SocketAddr}`，用于 IP 值、区域和前缀；`net::{TcpListener, TcpStream, UdpSocket, WaitOptions}`，用于基于 Linux amd64 系统调用的网络、共享 epoll 就绪检测、超时和取消
- `net::resolve`、`resolve_with`、`TcpStream::connect_host` 和 `net::tls` 验证 TLS 客户端
- `num` 结构化解析、标量转换 trait、显式舍入模式，以及带检查和饱和的 `i64` 算术
- `os::linux::syscall`，通过编号和六个机器字参数发起 Linux amd64 调用，支持限定作用域的可变字节缓冲区图、原始返回值和命名 errno
- `os::linux::abi` 带检查的 Linux amd64 记录编解码器和 `Iovecs`/`Message` 嵌套缓冲区构建器
- `os::linux::fd` 自主管理的 Linux 描述符、标量/向量/定位 I/O、元数据、字节路径、相对于目录的操作、锁和匿名内存文件
- `os::linux::process` ID、信号、pidfd、子进程创建与等待、资源用量、限制和优先级
- `os::linux::memory` 匿名/文件映射、带检查的复制式访问、保护、同步/建议、页面锁定和驻留情况
- `os::linux::ipc` 管道、Unix socket 对、描述符传递、eventfd、timerfd、poll 和 epoll
- `path::join`、`clean`、`is_absolute`、路径组件检查和 `absolute_structured`
- `path::slash::{clean, join, split, base, dir, extension, is_absolute, matches}` 和可复用的 `Pattern`，用于纯斜杠分隔的逻辑路径
- `resource::{Scope, ScopeError, scope, with_cleanup, finish}`，用于显式清理和合并错误
- `process::Command`、结构化的完整进程执行、`ExitStatus`、`Output`、`exit` 和 `look_path_structured`
- `rand::ALGORITHM`、`next_u64`、确定性字节生成、整数区间，以及使用显式种子的洗牌
- `serde::Value`、`Serializer`、`Deserializer`、`Serialize`、`Deserialize`、`value_serializer`、`value_deserializer`、`to_value` 和 `from_value`
- `task::Scope`、`Task[T]`、`CancelToken`、`WaitResult[T]`、`scope` 和 `try_scope`
- `testing::fail`、布尔/相等断言，以及 `Option`/`Result` 形态断言
- `text::StringBuilder`、`LineIndex`、`LineColumn` 和 `PositionEncoding`；字节偏移搜索、字符迭代与切片、修剪、拆分、替换、连接、重复和显式 ASCII 操作都是字符串方法
- `toml::Value`、`parse`、`encode`、serde `Serialize` 和 `Deserialize` 再导出、`to_value`、`from_value`、`to_string` 和 `from_string`
- `time::Duration`、`Instant`、`SystemTime`、`sleep` 和 `sleep_with`
- `utf8::validate`、`decode`、`decode_slice`、`encode`、`encode_into`、`encode_to`、`encode_char`、`encode_char_into`、`encode_char_to`、`encoded_len` 和 `Utf8Error`
- `utf8::{decode_first, decode_rune, decode_last_rune, decode_lossy, full_rune, rune_count, rune_start, valid_scalar, Decoder}`，用于纯标量解码和显式非法字节策略
- `utf16::decode`、`decode_bytes`、`encode`、`encode_bytes` 和 `Utf16Error`
- `unicode::VERSION`、标量属性，以及使用 Unicode 15.0.0 数据表的 Unicode 大小写转换

<a id="structured-error-foundation"></a>

## 结构化错误基础

`std::error` 提供标准库结构化错误使用的通用协议。`Error` 是要求 `Debug` 和 `ToString` 的标记 trait。`ErrorKind` 定义稳定且不依赖消息的类别，包括文件缺失、权限失败、非法输入或数据、超时、中断、短 I/O、管道断开、不支持的操作和 `Other`。

`Error` 为每个实现 `Debug` 和 `ToString` 的值提供通用实现，因此领域错误可直接参与通用协议。领域枚举可以派生这两个 trait，并通过各变体的 `#[to_string("...")]` 模板自定义面向用户的消息；见[属性与派生](language/metaprogramming.zh-CN.md#attributes-and-derive)。这些模板只影响 `ToString`：`Debug` 和序列化保留各自的表示。

`Report[E]` 保留叶子错误的类型，同时增加由上下文和聚合组成的不可变树。`Report::new(value)` 创建叶子；`.context(message)` 返回新的父节点，不改变原值。`Report::join(Slice[Report[E]])` 对给定序列建立快照，保留顺序和重复项，空序列返回 None。单元素序列仍是只有一个原因的聚合。对于不同类型的叶子变体，请使用领域枚举；此 API 不擦除类型，也不执行反射式向下转型。

`.value()` 仅在叶子处返回 Some，`.message()` 仅在上下文节点处返回 Some，`.causes()` 返回包含直接子节点的新向量。修改 join 的输入或返回的子节点向量都不会改变报告树。叶子值保留普通的复制/共享句柄语义：报告不会深拷贝可变的应用错误。

`.find(predicate)` 按深度优先、从左到右的顺序搜索叶子，在首个匹配处停止。`.find_map(select)` 还允许以保留类型的方式投影为另一种值；`.contains(target)` 使用叶子的 PartialEq，不比较显示字符串。上下文节点和聚合不会隐式匹配，叶子内部嵌入的任意错误也不会被递归解包。这些显式操作支持带类型的原因检查，无需修改现有 Error 标记 trait，也不添加 Go 风格的基于反射的 Is/As 行为。

报告可与 `Result::map_err` 和 `?` 组合，不会丢失领域变体或标准 I/O 错误字段。领域枚举可以同时容纳 `io::Error` 和应用验证失败；`find_map` 可将 I/O 变体投影回原来的带类型错误，包括类别和操作系统代码。在聚合的多个位置复用同一子树时，每次出现都会按顺序访问一次，不按对象身份去重。Report 不限制回调工作量和输出大小，回调也可能观察可变叶子句柄，因此调用方必须限制由外部控制的错误树大小，并避免并发修改载荷。

`.map(convert)` 将 `Report[E]` 转为 `Report[F]`，不要求任一叶子类型实现 Error 或 ToString。它按深度优先、从左到右的顺序，对每次出现的叶子调用一次 `convert`，保留每条上下文消息、分组边界、单元素分组、重复项和子节点顺序。原树不变；返回的叶子值遵循通常的复制/共享句柄语义。映射使用迭代栈，包括深层嵌套上下文；时间和临时存储与所访问的树成正比。共享子树在每次出现时单独映射。

导入 `error::ResultExt` 后，可在 `Result[T, E]` 上调用 `.into_report()`。它保持 Ok 不变，将 Err 转为报告叶子。导入 `error::ReportContext` 后，可在 `Result[T, Report[E]]` 上调用 `.context(message)` 或 `.with_context(make_message)`。后者仅在 Err 时将无参数闭包求值一次，Ok 时不求值。两个方法都为现有报告增加上下文父节点，因此重复调用 context 保持相同的 `Report[E]` 错误类型，并保留聚合。立即求值的 `.context` 参数遵循普通参数求值规则，Ok 时也如此。

转换是显式的：普通错误使用 `result.into_report().with_context(...)`，错误已是报告时使用 `result.with_context(...)`。`into_report` 总将错误包装为叶子；再次调用会有意创建 `Report[Report[E]]`。不存在自动展平或重叠的通用特化。这些辅助方法不会为 `?` 转换叶子类型；使用 `map_err(|report| report.map(convert))` 可将报告适配为领域错误。

当 E 实现 ToString 时，Report 实现 ToString 和 Debug，因此参与 Error 协议。显示时，上下文使用 `message + ": " + cause`，聚合使用换行分隔子节点，包括空消息和重复原因。显示不是序列化格式。搜索和显示使用迭代栈而非递归调用；工作量与访问节点数加渲染字节数成正比，遍历和输出都需要分配。不施加任意的深度上限。映射和上下文辅助方法使用普通泛型方法与扩展 trait，不增加语法、语法产生式或运行时错误接口。

```goml
use std::error;
use error::ResultExt;
use error::ReportContext;

fn with_context[T](result: Result[T, error::Details]) -> Result[T, error::Report[error::Details]] {
    result.into_report().with_context(|| "load settings")
}

fn adapt_report[E, F](report: error::Report[E], convert: (E) -> F) -> error::Report[F] {
    report.map(convert)
}

fn missing_file(report: error::Report[error::Details]) -> bool {
    report.find(|cause: error::Details| {
        error::kind_from_code(cause.kind_code()) == error::ErrorKind::NotFound
    }).is_some()
}
```

`kind_code` 和 `kind_from_code` 在 `ErrorKind` 与运行时和产物边界使用的稳定整数表示之间转换。未知整数映射为 `Other`。`Details` 保存稳定类别代码、操作、可选上下文、可选原始操作系统代码，以及仅供显示的消息。程序可以按类别代码或 `ErrorKind` 分支，但不得解析宿主消息。

```goml
use std::error;

fn describe[E: error::Error](value: E) -> string {
    value.to_string()
}

fn example() -> string {
    let value = error::Details::new(
        error::kind_code(error::ErrorKind::NotFound),
        "read",
        Some("missing.txt"),
        None,
        "file does not exist",
    );
    describe(value)
}
```

`io`、`fs`、`path`、`env`、`process`、`num` 和 `time` 暴露各领域专用的错误类型。调用方可按稳定错误类别分支，仅在需要显示消息时使用 `to_string()`。

<a id="utf-8-validation-and-conversion"></a>

## UTF-8 验证与转换

`std::utf8` 验证字节切片，并将完整字节向量转为字符串，不允许非法 UTF-8。`Utf8Error::valid_up_to` 是有效前缀的长度。`error_length` 是已知时的非法序列长度；输入末尾序列不完整时为 `None`。`encode` 返回字符串的 UTF-8 字节，`encoded_len` 返回单个 Unicode 标量值的字节长度。

标量解码器拒绝过长编码、代理码点、大于 U+10FFFF 的值、孤立的续字节和非法首字节。`validate` 与增量解码共用此检查。即使现有字节数小于标称宽度，已确定非法的前缀仍是错误：例如 `E0 80` 非法（`Some(1)`），而 `E1 80` 不完整（`None`）。非法字节恢复消耗一个字节；错误偏移指向首个失败标量的起始位置。这修正了此前将某些已知非法的短前缀归类为不完整序列的行为。

`decode_first(Slice[byte])` 返回首个标量经检查的 `(char, width)`，忽略后续字节；空或不完整输入返回带 `None` 的 `Utf8Error`。`decode_rune` 和 `decode_last_rune` 则实现替换解码：空输入返回 `(REPLACEMENT, 0)`，非法或不完整的非空输入返回 `(REPLACEMENT, 1)`，合法标量返回其编码宽度。`REPLACEMENT` 是 U+FFFD，`MAX_WIDTH` 是四。真正编码的 U+FFFD 宽度为三，可与非法字节替换区分。当前缀完整或已确定非法时，`full_rune` 为 true；它不是有效性检查。`rune_count` 逐个统计畸形字节；`decode_lossy` 将每个这样的字节替换为 U+FFFD，不合并相邻非法字节。这些 API 接受任意二进制输入，从不进行 Unicode 规范化。

`rune_start(byte)` 检查一个字节是否不是续字节，不保证它是合法首字节。`valid_scalar(u32)` 拒绝代理码点和越界值。GoML `char` 已保证是标量，因此 `encode_char(char) -> Vec[byte]` 没有非法字符的回退处理。`encode_char_into` 在写入前检查整个目标区间，返回宽度或 `bytes::BoundsError`，不做部分写入；`encode_char_to` 将编码字节追加到字节构建器。源码实现直接执行标量编码，并保留现有字符串转换边界。

`Decoder::new()` 创建严格增量解码器，最多保留四个字节。`feed(byte)` 返回 `Result[Option[char], DecoderError]`：`None` 表示合法前缀还需要更多字节，`Some` 恰好交付一次完整标量。每次字节调用都保留已交付输出，即使后续字节失败。`consumed()` 统计已接受的 feed 字节，包括证明非法的那个字节。`DecoderError::InvalidUtf8` 包含带绝对流偏移的 `Utf8Error`；畸形输入错误会持续到 `reset()`，后续 feed 不消耗输入。`finish()` 报告末尾不完整序列，否则关闭解码器；成功后的重复 finish 是幂等的，后续 feed 返回 `Finished`。`OffsetOverflow` 拒绝会超出 Linux amd64 isize 偏移范围的字节，不消耗它。`reset()` 清除偏移、待处理字节、失败和完成状态。解码器副本是共享可变句柄，不是独立快照，并发使用需要外部同步。

```goml
use std::utf8;

fn decode_stream(input: Slice[byte]) -> Result[Vec[char], utf8::DecoderError] {
    let decoder = utf8::Decoder::new();
    let output = Vec::new();
    for value in input {
        if let Some(character) = decoder.feed(value)? {
            output.push(character);
        }
    }
    decoder.finish()?;
    Ok(output)
}
```

导入 `utf8::BytesUtf8` 会添加 `Bytes::to_string_utf8`，返回 `Result[string, Utf8Error]`。

`decode_slice` 接受只读字节视图。`encode_into` 在经检查的偏移处写入 `MutSlice[byte]`，失败报告 `bytes::BoundsError`，不做部分写入；`encode_to` 追加到轻量的 `bytes::Builder`。

<a id="byte-buffers-and-endian-access"></a>

## 字节缓冲区与字节序访问

字节搜索可用于任意 `Slice[byte]`，不执行 UTF-8 解码。`bytes::find` 和 `rfind` 以 `Option[isize]` 返回首个/最后一个字节偏移；`contains`、`starts_with` 和 `ends_with` 返回布尔值。空模式在 `find` 中匹配偏移零，在 `rfind` 中匹配输入长度处。末次匹配搜索包含重叠出现。`find_byte` 和 `rfind_byte` 搜索单个字节。

`bytes::Finder::new(pattern)` 为模式建立快照，并构建前缀回退表。其 `find`、`rfind` 和 `contains` 方法在不同输入之间复用该不可变状态。构建花费 O(模式长度) 的时间/空间；每次搜索花费 O(输入长度) 时间和 O(1) 额外空间。后续修改模式缓冲区不会改变 finder。自由搜索函数每次构建新的 finder。

`bytes::cut(input, separator)` 返回首个匹配两侧的 `Some((before, after))`，不含分隔符；不存在时返回 `None`。`cut_prefix` 和 `cut_suffix` 仅在对应边缘匹配时返回剩余视图。返回的切片共享输入的底层存储；通过已有可变别名进行的修改可见。需要隔离时，用 `to_vec()` 复制返回的视图。空分隔符切分为空前缀和整个输入；切掉空前缀/后缀返回整个输入。

`bytes::count(input, pattern) -> Result[isize, TransformError]` 统计非重叠匹配。空模式匹配首个 UTF-8 标量之前和每个标量之后；每个畸形或截断字节前进一字节。这与替换解码一致，而非严格 UTF-8 验证，并保留所有原始字节。

`split(input, separator, max_parts)` 和 `split_after` 返回经过检查的共享输入视图向量；后者将每个匹配分隔符保留在前一部分。`split_n` 和 `split_after_n` 在 `max_parts` 前增加 `count` 参数：负数表示所有部分，零表示无部分，正数表示最多生成相应数量的部分，最后一部分是未拆分的余下内容。独立且非负的 `max_parts` 是资源限制：超出时返回错误，不截断。空分隔符在上述替换解码边界拆分，不产生开头/结尾的空部分。空输入加空分隔符产生零部分；非空分隔符保留边界和相邻分隔符之间的空部分。

`split_iter(input, separator)` 和 `split_after_iter` 返回单次遍历的 `FnIterator[Slice[byte]]` 值，具有相同的无限制拆分语义。它们在构建时为分隔符建立快照并编译，仅在 `next()` 请求时扫描，不分配结果向量。`lines_iter(input)` 产生以换行分隔的视图，包含每个结尾 `\n`，若存在未以换行结束的末行也会产生它。末尾换行后不额外产生空部分，空输入不产生部分；保留 `\r`，不做规范化。

迭代器副本共享遍历状态；需重新开始时应构建另一个迭代器。耗尽后，后续调用均返回 `None`。返回的视图和待处理输入保留共享底层存储，因此后续输入修改会体现在之前的视图中，并影响未来读取；分隔符修改不会影响迭代器。不要并发调用 `next()`，也不要在调用期间并发修改输入。提前停止可避免扫描余下输入，但迭代器会保持输入存储可达，直至自身释放。初始化存储为 O(分隔符长度)，额外遍历状态为常量；收集结果的调用方必须自行施加上限，或使用有界向量 API。立即拆分的变体复用相同遍历逻辑，包括限制数量时的最终剩余部分。

`replace(input, old, replacement, max_bytes)` 替换所有非重叠匹配。`replace_n` 在 `max_bytes` 前增加最大替换次数；负数替换全部，零生成不变的副本。空 old 模式在首个标量前以及后续替换解码边界后插入。搜索、计数、拆分和替换保持与输入/模式/输出大小成线性的复杂度；重复搜索复用已编译模式。这些原始拆分/替换算法不需要 Unicode 属性表；下述支持 Unicode 的辅助函数复用标准 Unicode 数据，不引入 text 依赖。

`join(Vec[Slice[byte]], separator, max_bytes)` 和 `repeat(input, count, max_bytes)` 构建字节缓冲区。这三类构建操作都返回 `Result[Bytes, TransformError]`，在分配输出存储前检查最终字节长度，即使只有一部分或替换没有改变内容，也返回独立缓冲区。负的重复次数是错误；空输入可以重复任意非负次数，无需循环。

`TransformError` 区分 `InvalidLimit`、`InvalidCount`、`LimitExceeded` 和 `LengthOverflow`。即使结果为空也拒绝负限制；恰好达到限制可成功。失败从不修改输入，也不暴露部分结果。限制涵盖输出字节数或部分数量，不涵盖调用方拥有的输入、已编译搜索表或分配器开销；它们不会使全局分配失败变为可恢复错误。与其他共享切片一样，操作读取输入期间，调用方不得并发修改输入。

支持 Unicode 的字节辅助函数解码合法 UTF-8 标量，并将每个畸形或截断字节视为 U+FFFD，不合并相邻非法字节。`find_char` 和 `rfind_char` 搜索该标量流；因此搜索 U+FFFD 也会找到非法字节，这与原始字节模式搜索不同。`find_any`/`rfind_any` 接受由标量集合成员构成的字符串。`find_by`/`rfind_by` 分别正向/反向遍历，每个访问到的标量调用一次谓词，在首个匹配处停止。所有结果都是原始字节偏移，不是标量索引。

`fields_iter` 和 `fields_by_iter` 产生由 Unicode 空白或标量谓词分隔的非空共享视图。谓词按输入顺序惰性执行，每个访问到的标量一次。`fields`/`fields_by` 按非负 `max_fields` 结果数量上限收集，限制非法或超出时返回 `TransformError`。迭代器别名共享进度，耗尽后不会恢复。`trim_space`、`trim_by`、`trim_start_by`、`trim_end_by`、`trim_chars`、`trim_start_chars` 和 `trim_end_chars` 返回修剪边缘后的共享视图。字符集合对标量分类，不匹配字节子串；修剪 U+FFFD 可同时移除非法字节和真正编码的 U+FFFD。这些视图操作原样保留剩余字节，包括非法 UTF-8。

`map_chars(input, transform, max_bytes)` 则构建独立的合法 UTF-8 缓冲区。其转换函数返回 `Option[char]`：`None` 丢弃解码标量，`Some` 输出一个合法标量。每个访问到的标量调用一次回调，每个输出编码在追加前检查字节限制。因此恒等映射会将非法输入字节替换为 U+FFFD 的三字节 UTF-8 编码。`to_case(input, unicode::Case, max_bytes)` 和 `case_with(input, special, kind, max_bytes)` 使用相同的带检查映射流程，执行普通或特殊的简单大小写转换。负限制在回调前失败；其他失败不返回部分缓冲区，但不会回滚回调效果。限制不包括源存储和分配开销。

`bytes::replace_invalid_utf8(input, replacement, max_bytes)` 对每段最长连续畸形 UTF-8 字节执行一次替换，原样保留合法标量（包括编码的 U+FFFD）。空替换删除这些连续段。替换内容是任意字节，不会再次解码，因此要得到合法 UTF-8 输出，替换内容必须是合法 UTF-8。与 `map_chars` 不同，相邻畸形字节共用一次替换。函数在分配前预检查完整输出长度，负限制、溢出或输出过大时返回 `TransformError`，始终返回独立存储，即使输入本已是合法 UTF-8。

`equal_fold(left, right)` 按 Unicode 简单折叠循环比较替换解码后的标量，不做完整的多标量折叠或规范化。不同畸形字节可能相等，因为每个都解码为 U+FFFD。原始 `find`/`rfind`、字节搜索、前缀/后缀检查和拆分/替换仍是字节模式操作。调用读取输入时不得并发修改输入；共享视图会反映后续修改。

```goml
use std::bytes;

fn byte_search_example() -> Option[isize] {
    let finder = bytes::Finder::new(bytes::Bytes::from_string("aba").as_slice());
    finder.rfind(bytes::Bytes::from_string("ababa").as_slice())
}

fn invalid_byte_search_example() -> Option[isize] {
    bytes::find_char(b"\xFF".as_slice(), '�')
}

fn repair_byte_runs_example() -> Result[bytes::Bytes, bytes::TransformError] {
    bytes::replace_invalid_utf8(b"a\xFF\x80b".as_slice(), b"?".as_slice(), 3)
}
```

`std::bytes` 使用 `Slice[byte]` 和 `MutSlice[byte]` 作为借用视图。`Bytes::as_slice` 和 `as_mut_slice` 是零拷贝的，`slice_checked` 和 `slice_mut_checked` 则验证子区间。`bytes::Builder` 是轻量字节累加器，返回 `Bytes`。`Bytes::from_vec` 和 `to_vec` 共享可变存储。需要独立可变缓冲区时，使用 `from_vec_copy`、`from_slice_copy`、`to_vec_copy` 或 `copy`。`Builder::finish_copy` 和 `to_vec_copy` 从构建器分离出快照。这些复制操作不对源的并发修改进行同步。

`Bytes::freeze` 创建具有私有不可变存储的 `FrozenBytes` 快照。`FrozenBytes::from_vec_copy`、`from_slice_copy` 和 `from_string` 也创建快照。使用 `len`、`is_empty` 和 `get_checked` 检查；`slice_checked(start, end)` 返回共享该快照的不可变视图，包括空区间。非法索引和区间返回 `None`。`to_vec_copy` 和 `thaw` 返回独立可变副本。`FrozenBytes` 实现按内容比较的 `PartialEq` 和 `Eq`。

`std::bytes::endian` 是需显式选择的二进制格式包。其 `Builder` 在写入带类型值时增长；`Reader` 在只读视图上前进；`Writer` 在固定可变视图上前进，越界时返回 `endian::BoundsError`，不会先写入部分内容。顶层 `read_u16/u32/u64`、`read_i16/i32/i64`、`read_f32/f64` 和对应 `write_*` 函数接受显式 `Endian::Little` 或 `Endian::Big`。单字节操作省略字节序。每个操作在读写前验证完整区间。

有状态的 `Reader`、`Writer` 和 `Builder` 为 `bool`、`u8/u16/u32/u64`、`i8/i16/i32/i64` 和 `f32/f64` 提供带类型方法。多字节方法接受显式字节序。Reader 和 Writer 方法仅在成功后前进，副本共享游标；边界检查失败保持位置和目标不变。浮点方法保留 IEEE 位模式，包括带符号零和 NaN 载荷。顶层函数接受显式偏移，不推进任何游标。

布尔 `read_bool` 和 `write_bool` 操作占用一个字节，不接受字节序参数：零解码为 false，所有非零字节解码为 true；写入仅产生零或一。这些操作提供内存函数、游标/构建器方法以及 `std::bytes::endian::stream::{read_bool, write_bool}`，具有与其他定宽流函数相同的 EOF、错误传播和部分进度契约。例如，`builder.write_bool(true)` 追加一个字节，`reader.read_i16(Big)` 读取一个带符号双字节值。这些普通方法不引入语法或隐式结构体布局。

`Reader::read_with[T](width, decode)` 用于组合显式定宽记录或数组布局。回调 `(Reader) -> Result[T, BoundsError]` 接收恰好覆盖该块的独立游标。调用前检查完整外层区间。成功时外层游标按声明宽度前进，跳过未读填充；失败时外层游标不变。嵌套错误保留相对于发生错误的嵌套视图的偏移和长度，初始区间错误则描述外层视图。

`Writer::write_with(width, encode)` 接受 `(Writer) -> Result[(), BoundsError]`，先验证外层区间，再在恰好 `width` 字节的独立零填充缓冲区上调用回调。成功时复制整个块并推进外层游标；未写字节是零填充。失败时丢弃临时缓冲区，不改变外层位置或目标。`Builder::write_with(width, limit, encode)` 使用相同的暂存契约，仅在成功时追加；`limit` 限制结果构建器的总长度，包括已有字节。非法宽度或限制在分配或调用回调前失败。零宽块使用空视图调用回调，不推进游标。保留下来的临时 writer 从不与提交后的目标形成别名。

这些方法使用普通回调、结构体和循环，不使用反射、自动字段排序、派生宏或原生结构体布局。回调可以嵌套块并捕获显式字节序，不得通过捕获的别名重入或修改外层游标、构建器或其存储；无关回调副作用、并发修改和 panic 不会回滚。读取使用有界借用输入视图，不是不可变快照；写入为每个活动嵌套块分配一个临时缓冲区。构建输出存储前应限制不可信记录宽度和数组数量；构建器限制控制字节数，不控制回调执行时间。

```goml
use std::bytes::endian;

fn read_header(reader: endian::Reader) -> Result[(u16, bool), endian::BoundsError] {
    reader.read_with(4, |block: endian::Reader| {
        let version = block.read_u16(Big)?;
        let enabled = block.read_bool()?;
        Ok((version, enabled))
    })
}

fn write_header(writer: endian::Writer, version: u16, enabled: bool) -> Result[(), endian::BoundsError] {
    writer.write_with(4, |block: endian::Writer| {
        block.write_u16(version, Big)?;
        block.write_bool(enabled)
    })
}
```

`bytes::endian` 还提供无符号 base-128 varint 和有符号 ZigZag varint。`uvarint_len` 和 `varint_len` 报告编码大小，最多 `MAX_VARINT_LEN64`（10）。`read_uvarint(input, offset)` 和 `read_varint(input, offset)` 在 `Result` 中返回 `(value, bytes_consumed)`；数量相对于给定偏移。它们保持输入不变，接受仍可放入 64 位的非最短编码。`VarintError` 区分 `InvalidOffset`、`Truncated` 和 `Overflow`，携带绝对字节偏移。第十字节大于 1 时立即视为溢出，包括续字节；这不同于 Go 切片解码器将某些十字节、全为续字节的输入报告为不完整的行为。

`write_uvarint(output, offset, value)` 和 `write_varint` 在写入前验证完整输出区间，返回字节数或 `BoundsError`；失败不会部分修改目标。`append_uvarint(builder, value)` 和 `append_varint` 向 `bytes::Builder` 追加。这些自由函数不推进 endian `Reader` 或 `Writer`。

`endian::Reader::read_uvarint()` 和 `read_varint()` 从当前位置解码，仅成功后推进位置并返回值。截断和溢出输入保持位置不变；错误偏移是 reader 原始视图内的绝对偏移。`endian::Writer::write_uvarint(value)` 和 `write_varint(value)` 返回 `Result[(), BoundsError]`，仅成功时按编码长度前进。空间不足既不改变位置也不改变输出字节。reader/writer 副本共享位置，与现有定宽方法一致。例如，`writer.write_uvarint(300)` 写入两个字节并前进两个位置。endian 值包保留纯内存依赖图；流式 I/O 属于单独导入的子包。

导入 `std::bytes::endian::stream` 可使用 `read_uvarint(reader: R) -> Result[Option[u64], stream::ReadError]` 和 `read_varint(reader: R) -> Result[Option[i64], stream::ReadError]`，要求 `R: io::Read`。首字节前的正常 EOF 返回 `None`；读取任何前缀后 EOF 返回 `ReadError::Decode(VarintError::Truncated(count))`。溢出和非规范但有效的编码遵循现有字节解码器。读取每次请求一个字节，返回包括 `Interrupted` 在内的所有错误，验证返回数量，最多消耗十字节，绝不消耗下一个编码值。每次操作的解码错误偏移从零开始。`ReadError::Io(error, count)` 保留原始 I/O 错误及此前成功读取返回的字节数，无法计入失败读取内部未报告的副作用。与内存游标不同，流无法在错误后回滚已消耗前缀。

对于 `W: io::Write`，`stream::write_uvarint(writer, value)` 和 `write_varint(writer, value)` 返回 `Result[(), io::Error]`。它们使用有界的十字节临时缓冲区编码，并调用 writer 的 `write_all`，遵循适配器特有的不重试行为。失败可能发生在部分写入后，不报告进度数量。读写都不刷新或关闭底层流。例如，`stream::write_uvarint(io::Discard::new(), 300)` 向丢弃型 writer 输出双字节编码。stream 包依赖 endian 和 I/O，反向不依赖；这些 API 由现有泛型函数和错误枚举表达，无语法变更。

同一包还提供 `read_u8/i8(reader)`、`read_u16/i16/u32/i32/u64/i64/f32/f64(reader, order)`，以及接受 `(writer, value)` 或 `(writer, value, order)` 的对应 `write_*` 函数。这里 `order` 是 `endian::Endian::{Little, Big}`。每次读取按标量类型返回 `Result[Option[T], stream::ReadError]`；每次写入返回 `Result[(), io::Error]`。整数和 IEEE 浮点位编码复用内存编解码器，包括带符号零和 NaN 载荷，不调用原生算法，也不使用反射。

定宽读取使用一个一、二、四或八字节缓冲区，最多填充字段剩余宽度，返回包括 `Interrupted` 在内的所有错误，并验证每次报告的数量。任何字节前的 EOF 返回 `None`；前缀后的 EOF 返回带 `UnexpectedEof` 和成功读取前缀长度的 `ReadError::Io(error, count)`。其他 I/O 错误按同一进度约定保留原始详情。成功读取恰好消耗字段宽度，不碰后续字段。定宽写入在调用 `write_all` 前完成编码，保留目标重试策略，从不刷新或关闭目标。例如，`stream::read_u32(reader, Big)` 读取一个可选的四字节大端字段。这些是普通泛型函数，不是新的二进制布局语言。

对于显式聚合布局，`stream::read_with(reader, width, limit, decode)` 和 `stream::write_with(writer, width, limit, encode)` 接受与内存块相同的有界 endian Reader/Writer 回调。在分配、调用回调或 I/O 前检查每条记录的非负字节限制；非法参数返回 `InvalidInput`（读取进度为零）。每次操作暂存恰好 `width` 字节，零填充和尾部字段规则与内存块相同。读取在解码前填满块，遇到任何读取错误均停止，绝不消耗下一记录；正常 EOF 返回 `None`，部分块返回带已确认进度的 `UnexpectedEof`，回调边界错误变为 `InvalidData`，进度等于已消耗的整个块。该错误包含回调边界诊断，不表示流已回滚。零宽读取调用解码器并返回 `Some(value)`，不探测 EOF；不要将其用作依赖 EOF 终止的循环条件。

流式写入先调用编码器，再调用 `write_all`；回调边界错误变为 `InvalidData`，不进行任何底层写入。后续 I/O 失败可能留下部分写入的记录。零宽写入仍调用编码器，成功后以空切片调用 writer 的 `write_all`；自定义覆盖保留自身语义。两个函数均保留既有的不刷新、不关闭行为，也不对回调副作用同步或撤销。例如，十六字节记录可通过 `stream::read_with(source, 16, 16, |block: endian::Reader| read_record(block))` 复用 `read_record`。固定数组元素数、字段顺序、填充、混合字节序和语义验证仍由调用方显式决定；没有反射或自动 serde/原生布局映射。

```goml
use std::bytes;
use std::bytes::endian;

fn encoded_integer() -> bytes::Bytes {
    let output = bytes::Builder::new();
    endian::append_varint(output, -123);
    output.finish()
}
```

<a id="utf-16-conversion"></a>

## UTF-16 转换

`std::utf16` 在字符串与 `Slice[u16]` 之间，或字符串与标明字节序的字节切片之间转换。解码以带索引的 `Utf16Error` 拒绝孤立代理码和奇数字节长度；编码为非 BMP 标量值输出代理对。

<a id="text-and-source-positions"></a>

## 文本与源码位置

文本搜索索引是 UTF-8 字节偏移，与内建字符串 API 接受的索引一致。`starts_with_at` 对越界和非字符边界偏移返回 false。`rfind` 返回最后匹配的字节偏移。修剪识别 ASCII 空白。按空分隔符拆分时，将原字符串作为一个元素返回；`split_once` 使用空分隔符时返回 `None`。

`text::LineIndex` 预计算行起始位置和非 ASCII 标量位置，用于重复的源码位置转换。`offset_to_line_column` 和 `line_column_to_offset` 支持 UTF-8、UTF-16 和 UTF-32 列，将 CRLF 视为一个行尾，并将未检查的越界位置钳位。`line_column_to_offset_checked` 拒绝编码标量内部或源码之外的位置。

`text::find`、`text::rfind` 和 `text::find_bytes` 返回字节偏移；字符串方法 `find`、`rfind` 和 `find_bytes` 是推荐形式。`text::char_indices` 产生字节偏移与 Unicode 标量值的配对，`text::char_count` 统计标量值，`text::slice_chars` 使用标量值索引，非法区间返回 `None`。

`text::Finder::new(pattern)` 编译不可变 UTF-8 模式，供重复调用 `find`、`rfind` 和 `contains`。前缀表构建需要 O(模式字节长度) 时间和空间；搜索需要 O(输入字节长度) 时间和常量额外状态，不将输入复制到字节向量。偏移以字节计，包括多字节字符。末次匹配搜索包含重叠。空模式的首次搜索匹配零偏移，末次搜索匹配输入字节长度处。自由函数 `text::find`、`rfind` 和 `find_bytes` 使用此实现，同时保留已有结果。为保持自举兼容性，仍保留编译器拥有的字符串方法回退实现。

`find_char`/`rfind_char` 搜索一个标量；`find_any`/`rfind_any` 搜索标量集合的成员，`contains_any` 检查成员关系。集合中重复字符无特殊含义；空集合从不匹配。`find_by` 和 `rfind_by` 接受标量谓词，分别正向/反向扫描，在该方向首个匹配标量处停止。每个访问到的标量只分类一次。所有搜索以 `Option[isize]` 返回字节偏移，绝不是标量索引，并且仅接受有效 Unicode 字符。集合搜索使用以字符为键的哈希映射，不对每个输入字符反复搜索子串。

`text::count(value, pattern)` 统计非重叠匹配，返回 `Result[isize, bytes::TransformError]`。空模式统计标量边界（标量数加一）；溢出返回 `LengthOverflow`。`cut` 返回首个分隔符两侧的 `Option[(string, string)]`；`cut_prefix` 和 `cut_suffix` 返回可选剩余部分。空边缘模式会匹配，`cut(value, "")` 返回 `Some(("", value))`。这有意区别于保留的 `split_once(value, "")`，后者返回 `None`。

现有公共操作也能组合使用，无需添加文本专用别名：`cmp::compare(left, right).to_isize()` 按字典序比较字符串；`find_char(value, character).is_some()` 和 `find_by(value, predicate).is_some()` 检查标量或谓词成员关系。字符串 `starts_with`/`ends_with` 检查固定前缀/后缀。`cut_prefix(value, prefix).unwrap_or(value)` 及类似的 `cut_suffix` 组合移除一个匹配的固定边缘，未匹配输入保持不变；它们不是修剪字符集合。

搜索单字节时，使用 `bytes::find_byte(value.to_bytes().as_slice(), needle)` 或 `rfind_byte`。这种组合复制 UTF-8 字节，可以找到编码内部字节：在 `"é"` 中搜索 `0xa9` 返回字节偏移 1，它不是标量边界。字符串模式操作 `find_bytes` 不能替代这种任意字节搜索。非法 UTF-8 修复同样应在字节边界处理：先 `bytes::replace_invalid_utf8`，再做带检查的 UTF-8 转换。GoML 字符串已是合法 UTF-8；非法的替换字节序列不能保证修复后的字节可以转为字符串。

不同 API 类别对空模式有意采用不同契约：

| 空模式操作 | 结果或策略 |
| --- | --- |
| `find` / `rfind` | 零偏移 / 输入字节长度。 |
| `count` | 标量数加一，空输入也为一。 |
| `cut` / `cut_prefix` / `cut_suffix` | 空前缀加整个输入 / 不变的剩余内容 / 不变的剩余内容。 |
| 保留的 `split` / `split_iter` / `split_after_iter` | 整个输入一次，包括空输入。 |
| 带检查的拆分 | 按标量分解，空输入产生零部分；N 变体应用数量限制。 |
| 保留的 `split_once` | `None`。 |
| 保留的 `replace` | 输入不变。 |
| 带检查的替换 | 在标量边界插入，受替换次数和输出限制约束。 |
| 已编译的 `Replacer` | 按规则顺序匹配，采用其文档规定的标量边界空规则策略。 |

`text::split_iter` 和 `split_after_iter` 产生子串，不预先构建结果向量；后者将每个分隔符保留在前一部分。它们使用一个已编译分隔符，仅在请求时扫描。与现有 GoML 字符串拆分一样，空分隔符恰好产生一次整个输入，即使输入为空；不采用 Go/字节拆分的标量分解方式。非空分隔符保留相邻和边缘空部分。`lines_iter` 与现有 `lines` 一致：省略 LF，每行去掉一个末尾 CR（未终止末行也如此），末尾 LF 后不额外产生空行。`lines_inclusive_iter` 则遵循 Go 行序列约定：保留 LF 和 CR，不额外产生末尾空行。两个行迭代器对空输入均不产生内容。

这些是单次遍历的 `FnIterator[string]` 值：别名共享进度，耗尽后不会恢复，另一次构建启动独立遍历。它们在存活期间保留不可变源字符串。收集结果可能分配无界输出；提前停止不会扫描余下输入。不引入新语法形式。

`text::fields_iter(value)` 惰性产生由 Unicode 15.0.0 空白分隔的非空字段。`fields_by_iter(value, separator)` 则对每个访问到的标量用谓词分类，按从左到右顺序恰好一次，在迭代开始前不调用。开头、结尾和重复分隔符不产生空字段。这些迭代器复制时共享进度，返回 `None` 后始终保持耗尽。

`fields(value, max_fields)` 和 `fields_by(value, separator, max_fields)` 使用显式非负结果数量限制收集相同字段，返回 `Result[Vec[string], bytes::TransformError]`。恰好达到限制可成功；负限制在任何谓词调用前失败。超出限制不返回部分向量，但谓词可能已运行，其副作用不回滚。限制约束结果数量，不约束输入扫描或分配器开销。全部是分隔符的输入在零限制下也成功。

`trim_start_by`、`trim_end_by` 和 `trim_by` 移除满足谓词的边缘标量。左修剪正向扫描，右修剪反向扫描，组合形式依次应用这两种操作，可能对一个保留的标量分类两次。`trim_chars(value, characters)` 移除标量集合成员，不匹配子串模式；`trim_start_chars` 和 `trim_end_chars` 使用相同集合语义，但只处理一端。`trim_space` 使用 Unicode 空白。现有 `trim`、`trim_start` 和 `trim_end` 保持仅 ASCII 的契约。

`equal_fold(left, right)` 逐标量比较 Unicode 简单折叠等价性，不使用区域设置或规范化规则。它接受开尔文符号/K 和希腊 sigma 变体，但不认为 `ß` 与 `ss` 相等；需要多标量折叠时，使用 `std::unicode` 的完整 `case_fold`。

`join_checked(values, separator, max_bytes)`、`repeat_checked(value, count, max_bytes)` 和 `replace_checked(value, old, replacement, max_bytes)` 在非负字节限制内构建合法 UTF-8 文本，返回 `Result[string, bytes::TransformError]`。它们在分配输出存储前检查结果长度。恰好达到限制可成功，算术溢出返回 `LengthOverflow`，负的重复次数返回 `InvalidCount`。空文本可重复任意非负次数，无需循环。现有无界 `join`、`repeat` 和 `replace` 保持不变。

`replace_n_checked` 在 `max_bytes` 前插入最大替换次数：负数表示所有非重叠匹配，零表示不替换。带检查的替换对空 old 字符串采用 Go 的边界插入语义（首个标量前和后续各标量后），区别于保留的无界 `replace`，后者在 old 为空字符串时不改变输入。计算大小和构建两遍匹配共用一个已编译 Finder，不分配所有匹配偏移。

`map_chars(value, transform, max_bytes)` 按输入顺序，对每个访问到的标量调用一次 `transform(char) -> Option[char]`。`None` 删除标量；`Some` 输出一个合法标量。它不进行通过回调预估大小的遍历：每个输出标量的编码字节长度都在追加前立即检查。出错时不返回部分文本，但此前回调副作用保留；负限制在调用回调前失败。当所有输入标量都被移除时，映射可以在零限制下成功。

`split_checked(value, separator, max_parts)` 和 `split_after_checked` 收集有界子串，后者将分隔符保留在前一部分。其 `_n_checked` 形式在 `max_parts` 前插入 `count`：负数表示所有部分，零表示无部分，正数表示最多产生该数量的部分，最后一部分保留未拆分的余下内容。独立的非负结果限制在超出时返回 `TransformError`，不截断结果；即使 count 为零，负限制也失败。不返回部分向量。

带检查的拆分遵循 Go 的空分隔符标量分解方式：不产生空边缘部分，空输入产生零部分，正 count 限制标量部分并保留最终剩余内容。这有意区别于保留的 GoML `split`/`split_iter`，后者对空分隔符产生整个输入。两类 API 对非空分隔符都保留相邻和边缘空部分。限制统计结果部分数，不是字节数或分配器开销。这些 API 不增加新语法形式。

`case_checked(value, unicode::Case, max_bytes)` 执行简单大写/小写/标题大小写映射。`case_with_checked(value, special, kind, max_bytes)` 使用 `unicode::SpecialCase` 覆盖并提供普通回退。这些是单标量映射，不是规范化、完整折叠或单词首字母大写算法。限制统计编码字节，不是字符，也不保证能从分配器耗尽中恢复。

```goml
use std::text;

fn find_second_character() -> Option[isize] {
    text::Finder::new("界").find("é界")
}

fn unicode_text_example() -> bool {
    text::equal_fold("K", "k") && text::trim_space("　x　") == "x"
}

fn bounded_text_example() -> bool {
    text::repeat_checked("界", 3, 9).is_ok()
        && text::repeat_checked("界", 3, 8).is_err()
}

fn scalar_split_example() -> bool {
    text::split_n_checked("é界🙂", "", 2, 2).is_ok()
        && text::find_any("é界", "界") == Some(2)
}
```

<a id="explicit-word-title-policy"></a>

### 显式单词首字母大写策略

简单标题大小写映射会转换每个标量，不是单词首字母大写格式。调用方可通过有状态的 `map_chars` 表达显式单词边界策略。以下有界兼容写法在合法 UTF-8 上匹配 Go 已弃用的 `strings.Title`。它不是推荐的 Unicode 分词方式：ASCII 标点形成边界，但大多数非 ASCII 标点不会。语言学和区域感知的分词属于生态文本处理范围。

```goml
use std::bytes;
use std::text;
use std::unicode;

fn legacy_title_separator(value: char) -> bool {
    if value.to_u32() <= 0x7f {
        !((value >= '0' && value <= '9')
            || (value >= 'a' && value <= 'z')
            || (value >= 'A' && value <= 'Z')
            || value == '_')
    } else {
        !unicode::is_letter(value)
            && !unicode::is_digit(value)
            && unicode::is_whitespace(value)
    }
}

fn legacy_word_title_checked(
    value: string,
    max_bytes: isize,
) -> Result[string, bytes::TransformError] {
    let previous = Ref::new(' ');
    text::map_chars(
        value,
        |character: char| {
            let boundary = legacy_title_separator(previous.get());
            previous.set(character);
            Some(
                if boundary {
                    unicode::to_titlecase(character)
                } else {
                    character
                },
            )
        },
        max_bytes,
    )
}
```

前一个**原始**标量决定是否映射当前标量。因此 `"hELLO world"` 变成 `"HELLO World"`，不将词内字符转为小写；`"foo—bar"` 则变成 `"Foo—bar"`。每次调用使用全新状态。输出字节限制和错误与 `map_chars` 相同，不返回部分构建的文本。此写法使用普通闭包和引用，不是新的标准函数或语法形式。

<a id="typed-field-formatting"></a>

## 带类型的字段格式化

导入 `std::text::format` 可使用显式字段类型组装有界消息。这是纯库 API：`f"{expression}"` 仍使用 ToString，不带格式说明。没有 printf 格式字符串解析器、反射、异构 Any 参数列表、指针/类型名格式化或 scanf 协议。现有数字解析和生态扫描器仍是独立设施。

`Formatter::new(max_bytes)` 拒绝负字节限制。其 `literal`、`text`、`integer`、`unsigned`、`float32`、`float64`、`display` 和 `debug` 方法追加一个字段，返回 `Result[(), Error]`。`len()` 和 `remaining()` 报告字节数。`finish()` 返回独立字符串快照，不关闭或消耗 formatter。Formatter 副本共享私有构建器和预算，不暴露可写构建器或切片。不支持并发使用。

`Padding { width, alignment, fill }` 以 Unicode 标量衡量最终渲染字段，不是 UTF-8 字节、字素簇或终端列。对齐方式为 Left、Right 或 Center；Center 遇到奇数个额外填充标量时，多出的一个放在右侧。Width 是最小宽度，从不截断。Fill 是一个 char，可为多字节标量。`Padding::plain()` 为零宽、右对齐和空格填充。

`TextSpec { padding, max_chars, style }` 可选地在渲染前按 Unicode 标量数截断输入。TextStyle 为 Plain、Quoted 或 QuotedAscii；后两种使用现有文本引用规则。随后填充统计渲染后的标量，包括引号和转义。`TextSpec::plain()` 不截断也不加引号。不使用规范化、字素分割或显示宽度表。

`IntSpec { padding, radix, uppercase, alternate, sign, min_digits }` 接受 2 到 36 的进制。整数字段接受 i64，无符号字段接受 u64；较窄值可以显式扩宽。`alternate` 仅对 2、8、16 进制有效，输出 0b、0o 或 0x（请求时使用大写）。Sign 为 NegativeOnly、Always 或 Space。顺序为符号、前缀、补零后的数字。`min_digits` 非负；即使最小值为零，数值零仍有一位数字。外层填充应用于整个字段：填充字符 `0` 不意味着考虑符号位置的补零。例如，负的十六进制 15，启用 alternate 且 min_digits 为 4 时是 `-0x000f`，而对 `-15` 右对齐、外层补零至宽 6 则得到 `000-15`。`IntSpec::plain()` 使用十进制、NegativeOnly、无填充且无前缀。

`FloatSpec { padding, sign, style }` 复用 num 中分开的 f32 和 f64 格式化器。FloatStyle 支持 Shortest(FloatNotation)、Fixed(decimal_places)、Scientific(decimal_places, uppercase)、General(significant_digits, uppercase)、Binary 和 Hex(fractional_digits, uppercase)。Fixed、Scientific 和 General 精度必须非负；General 的零表示一位有效数字。Hex 接受 -1，表示精确且去掉末尾多余位的输出，或非负小数精度。现有 num 规则决定舍入、指数写法、NaN、无穷和负零。保留已有前导符号：不会给 `+Inf` 添加第二个符号。无符号 NaN 可因 Sign 变为 `+NaN` 或 ` NaN`。`FloatSpec::plain()` 使用 Shortest(General) 和 NegativeOnly。

`display[T: ToString](value, padding)` 和 `debug[T: Debug](value, padding)` 在调用对应转换恰好一次之前验证填充。这些适配器无法限制用户转换内部的分配或时间，也不捕获其 panic。转换可通过别名重入 formatter；即使外层字段失败或 panic，它自身成功的写入及其他副作用仍保留。外层字段在转换后检查当前剩余预算，因此重入无法重置或绕过预算。

所有库字段方法在修改目标前验证格式说明和完整输出大小。返回错误不会部分追加该字段，也不会改变构建器容量；此前成功的字段仍保留。Error 区分 InvalidLimit、InvalidSpec(message)、LengthOverflow、LimitExceeded(total_limit) 和 ConversionFailure(message)。这些错误值支持 ToString、Debug、PartialEq 和 Eq。LimitExceeded 报告 formatter 的总限制，不是当前剩余量。所有输出字节都计入，包括符号、前缀、引用和多字节填充。在分配依赖宽度或精度的缓冲区前检查长度算术。临时字段和目标各自可能使用 O(max_bytes) 存储；这不是 CPU、分配器或进程内存配额。输入扫描和用户转换有独立开销。

`write_to[W: io::Write](writer, value)` 写入已完成的字符串，返回 `Result[usize, io::TransferError]`，统计已确认的 UTF-8 字节。它使用原始 write、验证数量、继续短的成功写入，零进度报告 WriteZero。任何错误（包括 Interrupted）都终止操作，不重试该次调用；未确认的外部副作用不会被重放。空输入不调用 writer。从不刷新或关闭 writer，也不调用被覆盖的 write_all 方法。如果要求格式化失败时无外部输出，应先格式化到私有 Formatter，再调用 write_to；写入本身不是原子的。

```goml
use std::text::format;

fn message(value: i64) -> Result[string, format::Error] {
    let output = format::Formatter::new(128)?;
    output.literal("value=")?;
    output.integer(value, format::IntSpec {
        radix: 16, alternate: true, min_digits: 4,
        ..format::IntSpec::plain()
    })?;
    Ok(output.finish())
}
```

协议模板、上下文相关的 HTML 转义和终端布局仍属于生态范围。在推进 stage0 前，编译器和驱动中的格式化使用方保留已发布的回退实现。

<a id="ascii-and-unicode"></a>

## ASCII 与 Unicode

`std::ascii` 操作 `byte`。分类和大小写转换仅使用 7 位 ASCII 范围，`0x7f` 以上字节保持不变。`escape_default` 对制表符、回车、换行、引号和反斜杠输出短转义，保留可打印 ASCII，对其他字节使用小写 `\\xNN` 转义。

`std::unicode` 将公共数据版本固定为 Unicode 15.0.0。字符谓词操作一个 Unicode 标量值。`lowercase` 和 `uppercase` 对完整字符串应用 Unicode 大小写映射。`case_fold` 使用已检入仓库、由 `tools/generate_unicode_casefold.py` 生成的数据表，支持多标量折叠，且不依赖区域设置。

分类和简单大小写转换使用已检入仓库、由 `tools/generate_unicode_tables.go` 生成的不可变表；生产环境查找和转换使用纯 GoML。谓词包括 `is_control`、`is_letter`、`is_number`、`is_digit`、`is_mark`、`is_punctuation`、`is_symbol`、`is_whitespace`、`is_lowercase`、`is_uppercase`、`is_titlecase`、`is_graphic` 和 `is_printable`。图形字符包含 Unicode 空格分隔符；可打印字符除字母、标记、数字、标点和符号外，只额外包含 ASCII 空格。

`category`、`script`、`property`、`fold_category` 和 `fold_script` 接受区分大小写的 Unicode/Go 表名，返回 `Option[RangeTable]`；对应的 `*_names()` 函数返回新建且已排序的名称向量。折叠表包含类别或文字系统满足简单折叠闭包所需的额外码点。`RangeTable::contains(char)` 检查一个标量；`contains_codepoint(u32)` 还允许检查代理码类别，对高于 U+10FFFF 的值返回 false。`in_any(char, Vec[RangeTable])` 检查它们的并集。

`RangeTable::new(Vec[Range])` 将经检查的区间复制到不可变表。公共 `Range` 字段为 `low`、`high` 和 `stride`，均为 `u32`。上下界包含在内，成员关系为 `(value - low) % stride == 0`，上界本身不必是成员。输入必须已排序、不重叠、位于 U+0000..U+10FFFF 内，且步长非零。失败返回带索引的 `TableError::InvalidRange` 或 `UnsortedRange`。`ranges()` 返回独立副本；后续修改输入或返回的区间都不改变表。

`to_case(char, Case)` 接受 `Upper`、`Lower` 或 `Title`；`to_uppercase`、`to_lowercase` 和 `to_titlecase` 是快捷形式。`simple_fold(char)` 推进标量的简单折叠等价循环，区别于多标量 `case_fold(string)`。字符串 `uppercase` 和 `lowercase` 使用单标量映射，因此 `ß` 转大写仍是 `ß`。它们不执行规范化、上下文大小写转换或字素分割。

`SpecialCase::new(Vec[CaseMapping])` 接受严格递增、不重复的 `from` 字符和显式 `upper`、`lower`、`title` 映射，否则返回 `TableError::UnsortedMapping`。它为输入建立快照，对未列出的字符回退到普通 Unicode 大小写转换。其 `to_case` 方法和字符串 `lowercase_with`/`uppercase_with` 使用这些覆盖；`turkish_case()` 和 `azeri_case()` 提供标准的带点/无点 I 映射。

```goml
use std::unicode;

fn unicode_example() -> bool {
    let letters = unicode::category("L");
    unicode::uppercase_with("iı", unicode::turkish_case()) == "İI"
        && match letters {
            Some(table) => table.contains('界'),
            None => false,
        }
}
```

<a id="mathematics"></a>

## 数学

`std::math` 在 GoML 中实现初等运算，同时保留公共 API 和 IEEE 754 特殊值行为。它不绑定 Go 的 `math` 包。f32 形式通过 f64 计算，再将结果舍入回 f32。超越函数结果不保证与 Go 或跨目标平台逐位相同。

`std::math::complex::Complex` 是由一对 `f64` 值组成的纯 GoML 类型，公开 `real` 和 `imaginary` 字段。`new`、`from_real` 和 `from_polar` 构建值；`add`、`sub`、`neg`、`mul`、`conjugate`、`abs`、`phase` 和 `polar` 提供基本算术和坐标。`try_div` 和 `powi` 为零除数返回 `Result[Complex, DivisionError]`。初等方法包括 `exp`、`log`、`log10`、`sqrt`、`pow`、`sin`、`cos`、`tan`、`sinh`、`cosh`、`tanh` 及其反函数。`tan` 和 `tanh` 在计算出的分母为零时也返回 `Result`。包提供 `atan2(y, x)` 计算相位。有限数除法缩放极端数量级，减少可避免的溢出。对数、平方根和反函数使用主分支，以带符号零处理分支切线。`try_div` 报告零分母而不返回 IEEE 无穷；极端超越函数输入可能通过标量 `exp` 溢出。这些运算复用标量 GoML 数学函数。

```goml
use std::math::complex;

fn quotient() -> Result[complex::Complex, complex::DivisionError] {
    complex::Complex::new(3.0, 4.0).try_div(complex::Complex::new(1.0, -2.0))
}
```

独立的 [`ecosystem::bigmath`](https://github.com/gomlang/bigmath/blob/HEAD/README.md) 模块基于 `ecosystem::bigint` 提供有界的任意精度有理数和二进制浮点算术。`Rational` 精确规范化分数；`Float` 使用显式的 1–4096 位精度和舍入上下文。它报告资源、除零和指数错误，不会静默创建无穷或 NaN。该模块不改变编译器的数值类型。

<a id="fixed-width-bit-operations"></a>

## 定宽位操作

`std::math::bits` 在 GoML 中实现可移植整数算法，不绑定 Go `math/bits`。`ones_count8/16/32/64`、`leading_zeros8/16/32/64`、`trailing_zeros8/16/32/64` 和 `len8/16/32/64` 返回 `isize` 计数。零的置位数/位长度为零，前导/尾随零数等于其宽度。`reverse8/16/32/64` 反转位；`reverse_bytes16/32/64` 反转字节序。`rotate_left8/16/32/64` 接受 `isize` 数量，并按字宽取模；负数向右旋转。

`add32/64(left, right, carry)` 和 `sub32/64(left, right, borrow)` 接受布尔进位/借位，返回 `(word, bool)`。`mul32/64` 返回完整双倍宽度乘积的 `(high, low)`。`div32/64(high, low, divisor)` 返回 `(quotient, remainder)` 或 `DivisionError::DivisionByZero`/`Overflow`；当 `high >= divisor` 时商溢出。`rem32/64` 接受任意高位字，仅在除数为零时失败。这些不是恒定时间密码学原语。不提供原生字宽别名，也不要求新的内建操作或语法。

<a id="portable-simd"></a>

## 可移植 SIMD

`use std::simd;` 导入定宽值类型，具有可移植标量实现和 GoML 自有的 amd64 SSE2/AVX2 汇编后端。实现不导入 Go SIMD 包或第三方 SIMD 库。

| 元素 | 128 位向量 | 256 位向量 | 对应掩码 |
| --- | --- | --- | --- |
| `i8`, `u8` | `i8x16`, `u8x16` | `i8x32`, `u8x32` | `mask8x16`, `mask8x32` |
| `i16`, `u16` | `i16x8`, `u16x8` | `i16x16`, `u16x16` | `mask16x8`, `mask16x16` |
| `i32`, `u32` | `i32x4`, `u32x4` | `i32x8`, `u32x8` | `mask32x4`, `mask32x8` |
| `i64`, `u64` | `i64x2`, `u64x2` | `i64x4`, `u64x4` | `mask64x2`, `mask64x4` |
| `f32` | `f32x4` | `f32x8` | `mask32x4`, `mask32x8` |
| `f64` | `f64x2` | `f64x4` | `mask64x2`, `mask64x4` |

所有向量提供 `from_array`、`to_array`、`splat`、`from_slice` 和 `copy_to_slice`。数组转换使用对应定长标量数组，并复制元素。`from_slice(values, offset)` 返回 `Option[Vector]`；`copy_to_slice(values, offset)` 返回 `bool`。两者都在访问内存前验证完整区间，不要求特殊对齐，并拒绝负偏移或不完整向量。非法存储操作保持目标不变。

所有向量提供 `add`、`sub`、`mul`、`neg`、`abs`、`min`、`max`、`reduce_sum`、`reduce_product`、`reduce_min` 和 `reduce_max`。整数算术按元素宽度回绕，包括归约。对最小有符号整数取负或绝对值仍得到该最小值；无符号 `abs` 是恒等操作。整数向量还提供 `bitand`、`bitor`、`bitxor`、`bitnot`、`shl(count: u32)`、`shr(count: u32)`、`saturating_add`、`saturating_sub`、`reduce_and`、`reduce_or` 和 `reduce_xor`。移位次数按元素宽度取模。有符号右移扩展符号位。饱和算术将结果钳位到标量类型范围。整数向量不提供除法。

浮点向量额外提供 `div`、`sqrt`、`floor`、`ceil`、`trunc`、`round`、`round_ties_even`、`mul_add`、`is_nan`、`is_infinite` 和 `is_finite`。`round` 在恰好居中时远离零舍入；`round_ties_even` 选择最近偶数。普通算术每次操作都舍入到元素类型，因此 `a.mul(b).add(c)` 有两次舍入。`a.mul_add(b, c)` 计算仅一次舍入的融合结果，标量回退也如此。归约采用平衡树，每次操作都舍入。例如，四通道 `reduce_sum` 计算 `(lane0 + lane1) + (lane2 + lane3)`。

浮点 `min` 和 `max` 在恰好一个操作数为 NaN 时返回另一个数值操作数，两者均为 NaN 时返回 NaN。对相等的零，若存在负零则 `min` 选择负零，若存在正零则 `max` 选择正零。算术 NaN 载荷未指定。比较遵循标量规则，包括 NaN 无序，以及正负零相等。

`simd_eq`、`simd_ne`、`simd_lt`、`simd_le`、`simd_gt` 和 `simd_ge` 返回对应掩码；向量 `==` 返回检查所有通道的单个 `bool`。掩码提供 `from_array`、`to_array`、`splat`、`from_bitmask`、`to_bitmask`、`all`、`any`、`bitand`、`bitor`、`bitxor`、`bitnot` 和 `select(if_true, if_false)`。第零位对应第零通道；未使用的高位丢弃。至多八通道的位掩码使用 `u8`，十六通道使用 `u16`，三十二通道使用 `u32`。`select` 通过公共 `SimdSelect[M]` trait 接受任意匹配向量，复制通道位，不做算术。

```goml
use std::simd;

fn select_scaled(a: simd::f64x4, b: simd::f64x4, scale: f64) -> simd::f64x4 {
    a.simd_gt(b).select(a, b).mul(simd::f64x4::splat(scale))
}

fn blend_bytes(a: simd::u8x16, b: simd::u8x16) -> simd::u8x16 {
    a.simd_lt(b).select(a.saturating_add(b), a)
}
```

`swizzle(indices: [u32; N])` 从一个向量选择通道。`shuffle(other, indices: [u32; N])` 从接收者和 `other` 的连接序列中选择。超出各自 `0..N` 或 `0..2*N` 范围的索引产生零通道。`reverse`、`interleave_low`、`interleave_high`、`deinterleave_even` 和 `deinterleave_odd` 提供常见固定重排。交错操作交替排列两个完整向量的元素，返回该序列的下半或上半；解交错从两者连接序列选择偶数或奇数位置。这些操作作用于完整向量，包括跨越 256 位向量的两个 128 位半部。

当具有相同通道数、不同元素类型的目标向量存在时，数值转换使用 `to_i8`、`to_u8`、`to_i16`、`to_u16`、`to_i32`、`to_u32`、`to_i64`、`to_u64`、`to_f32` 或 `to_f64`。整数转换按源的有符号性扩展，缩窄时丢弃多余高位。浮点转整数向零截断，将越界值钳位到目标范围，NaN 映射为零。整数转浮点及 `f64` 转 `f32` 按最近值舍入，中点取偶。浮点 `to_bits` 和 `from_bits` 与对应无符号整数向量互转，不改变任何位。

```goml
use std::simd;

fn widen_and_scale(values: simd::i16x8) -> simd::f32x8 {
    values.to_i32().to_f32().mul(simd::f32x8::splat(0.5))
}

fn fused(a: simd::f32x4, b: simd::f32x4, c: simd::f32x4) -> simd::f32x4 {
    a.mul_add(b, c)
}
```

在 amd64 上，`goml build/run/test` 和 `gomlc run-single` 将受支持的顺序函数提取到生成的 `.s` 文件。内核可组合不同的受支持向量类型、对应掩码、标量广播、转换和归约，也可沿纯辅助函数调用继续提取。它们最多接受八个参数，并将中间值保留在寄存器中。对于包含向量或掩码的受支持签名，内联器保留源函数边界，包括混合标量参数、标量归约和布尔掩码结果，以便将操作一起提取。不支持的表达式、控制流或过高寄存器压力保留标量实现。切片操作和动态 shuffle 索引仍使用标量。某些数值转换和 32/64 位饱和算术也使用标量实现。整数归约可能在原生内核内使用标量指令。原生内核没有分配、回调或栈溢出存放。

基本 128 位内核使用 SSE2。使用 256 位向量或原生浮点舍入的内核要求 AVX2；符合条件的 128 位 FMA 内核独立于 AVX2 要求 FMA。分派通过 CPUID 和 XGETBV 检查 CPU 特性，以及操作系统是否保存 XMM/YMM 状态。不支持的特性保留标量路径。AVX/FMA 内核返回前执行 `VZEROUPPER`。构建时设置 `GOML_SIMD=scalar` 可禁用原生内核，或设置 `GOML_SIMD=sse2` 禁用生成的 AVX2 和 FMA 内核。其他架构及导出的独立 Go 源码保留标量实现；GoML 构建命令使用嵌入生成 Go 源码中的原生元数据。

目前不提供可配置通道数、gather/scatter、向量算术运算符、AVX-512 内核，或普通标量循环的自动向量化。SIMD 使用现有结构体、导入、trait 和方法调用语法，不增加表达式语法。`gomlc/testdata/simd/` 下的汇编快照由 `just update-golden` 和 `just verify-golden` 生成与检查。

<a id="randomness-and-cryptographic-helpers"></a>

## 随机数与密码学辅助函数

`std::crypto::digest_equal(left, right)` 比较等长字节切片，不在遇到不匹配字节时停止；长度不同立即返回 false。`crypto::verify_digest(digest, expected)` 接受 `std::hash::Hasher`，获取不消耗状态的 `sum`，并使用相同比较。HMAC 验证共享此契约。编译器不保证恒定时间执行；长度可观察，调用方必须选择合适算法和密钥处理方式。

`std::crypto::hash` 通过 `std::hash::Hasher` trait 提供纯 GoML 增量 SHA-2 摘要。`Sha256Digest::new224/new256` 和 `Sha512Digest::new384/new512/new512_224/new512_256` 选择标准 SHA-2 变体。`update` 接受字节切片；`sum` 返回独立快照，不消耗摘要；`reset` 恢复初始状态；`copy` 分叉当前状态。普通赋值共享摘要的可变状态，因此需要独立后续计算时必须使用 `copy`，并发访问必须串行化。便捷函数 `sha224`、`sha256`、`sha384`、`sha512`、`sha512_224` 和 `sha512_256` 返回字节缓冲区的小写十六进制摘要。`sha256_file` 在返回前哈希完整文件。这些函数提供无密钥哈希，不是消息认证。

`Sha3Digest::new224/new256/new384/new512` 实现固定输出 SHA-3，具有相同的不消耗状态的 `Hasher` 行为，并提供对应小写十六进制 `sha3_224/sha3_256/sha3_384/sha3_512` 辅助函数。`ShakeDigest::new128/new256` 支持显式的先吸收后挤出生命周期。重复 `squeeze` 继续相同输出流；开始挤出后 `absorb` 返回 `AlreadySqueezing`。一次 squeeze 最多接受 1 MiB，`reset` 开始新流。两类摘要均为纯 GoML Keccak 实现。

`std::crypto::hmac::Digest[H: hash::Hasher]` 实现带密钥的增量 HMAC。`Digest::new(make, key)` 和 `hmac::sum(make, key, input)` 接受无参数工厂，每次调用必须返回新的独立摘要。摘要支持 `update`、`reset`、不消耗状态的 `sum` 和 `verify(expected)`；长度匹配时验证循环比较每个字节。GoML 不保证由编译器强制的恒定时间执行。非法摘要尺寸返回错误。`std::crypto::hkdf::{extract, expand, derive}` 使用该 HMAC 契约完成 RFC 5869 密钥派生。`expand` 拒绝负长度和超过 255 个摘要块的输出；零长度输出合法。两个包都不增加 Go 密码学后端。

`std::crypto::pbkdf2::derive(make, password, salt, iterations, length)` 通过生成新摘要的工厂复用 HMAC。实际执行前检查正迭代次数、RFC 块计数器上限，以及 1 MiB 输出、一百万次迭代和一百万次总 PRF 调用的实用限制。零长度输出合法；非法摘要尺寸及参数/资源限制返回带类型错误。PBKDF2 减慢离线口令猜测，不加密数据，也不能替代应用专用的口令存储策略。

```goml
use std::bytes;
use std::crypto::{hash, hmac, hkdf};

fn derive_key() -> Result[bytes::Bytes, hkdf::Error] {
    hkdf::derive(
        || hash::Sha256Digest::new256(),
        b"salt".as_slice(),
        b"input key".as_slice(),
        b"context".as_slice(),
        32,
    )
}
```

`std::crypto::rand::bytes` 从操作系统密码学随机源读取所需字节数，不暴露随机源句柄。

原有顶层 `std::rand` 函数保留带版本的 `splitmix64-v1` 输出：每次调用接受显式种子，相同输入产生相同结果。有状态的 `Pcg::new(seed_high, seed_low)` 和 `ChaCha8::new(seed: [byte; 32])` 实现 `rand::Source`；两者均提供 `seed` 和 `next_u64`，PCG 还暴露当前 `state`。`Generator[S: Source]` 包装任一来源，提供带检查的 `bounded_u64` 和 `range_i64`，以及 `float64`、`float32`、`bytes`、`shuffled`、`normal_f64` 和 `exponential_f64`。有界采样通过拒绝采样避免模偏差；零上界、空区间和负字节长度返回 `SampleError`。生成器与有状态源的副本共享状态，因此并发调用方必须串行化访问。这些算法是确定性的，用于测试、模拟和采样。此包的 PCG 和 ChaCha8 都不是操作系统密码学随机源；秘密字节请使用 `std::crypto::rand`。正态和指数采样使用文档规定的 Box-Muller 和逆 CDF 公式，因此不保证其值与 Go 的分布流一致。

```goml
use std::rand;

fn roll() -> Result[u64, rand::SampleError] {
    rand::Generator::new(rand::Pcg::new(1, 2)).bounded_u64(6).map(|value| value + 1)
}
```

<a id="incremental-checksums"></a>

## 增量校验和

`std::hash::Hasher` 是独立于集合使用的 prelude `Hash` 的 trait。其方法为 `update(Slice[byte])`、`reset()`、`sum() -> bytes::Bytes`、`size()` 和 `block_size()`。导入 `use std::hash; use hash::Hasher;` 以使用方法调用语法。`sum` 不消耗状态，返回独立字节快照；之后仍可继续 update。以下校验和实现以最高有效字节优先编码结果，块大小为一字节。

- `hash::adler32::checksum(input)` 和 `update(previous, input)` 计算或继续 Adler32 校验和。初始校验和为 1；`previous` 是此算法产生的校验和。`Digest::new()` 支持 `Hasher` 和 `sum32()`。
- `hash::crc32::Table::new(polynomial)` 为反射多项式构建不可变共享表。常量为 `IEEE`、`CASTAGNOLI` 和 `KOOPMAN`。`checksum(input, table)` 从零开始；`update(previous, table, input)` 继续校验和。`Digest::new(table)` 支持 `Hasher` 和 `sum32()`。
- `hash::crc64` 使用相同的表/update/digest 设计，具有 `ISO`、`ECMA` 和 `sum64()`。API 遵循 Go 的反射、取补 CRC 约定；ECMA 多项式名称不意味着使用非反射的 CRC-64/ECMA-182 参数集。
- `hash::fnv::Digest::new32/new64/new128` 构建 FNV-1，`new32a/new64a/new128a` 构建 FNV-1a。它们均实现 `Hasher`；128 位状态通过普通整数算术实现。

赋值摘要会复制共享可变句柄。`digest.copy()` 显式创建独立状态；CRC 副本可共享不可变表。update/reset 不得与相同状态上的其他访问竞争。这些校验和不是密码学哈希或认证机制。不提供硬件加速或兼容 Go 的状态序列化。

`std::hash::stream` 将任意 Hasher 与现有 I/O 组合：`Reader::new(reader, hasher)` 实现 `io::Read`，`Writer::new(writer, hasher)` 实现 `io::Write`。构造函数保留给定哈希状态，允许已有哈希前缀。`inner()` 和 `hasher()` 按普通别名/复制语义返回保存的值；直接访问内层流会绕过哈希，共享标准摘要会共享其状态。使用 `Writer::new(io::Discard::new(), digest)` 可构建只做哈希的写入目标。

每个操作调用底层流一次，验证其返回数量，只哈希成功报告的非空前缀。短读/短写不会自动补齐；继承的 read_exact/write_all 方法提供此行为，并在 Interrupted 时返回错误，不重试。空请求会转发，但零数量和错误从不更新 hasher。负数/过大数量返回 InvalidData。底层错误原样传播，不保持为持续错误；允许后续调用。在返回 Err 前修改缓冲区或消耗字节的流，无法通过当前 Read/Write 契约报告该部分进度，因此这些未报告字节不会被哈希。不回滚底层 I/O。Writer.flush 转发底层 flush，不哈希或重置任何内容。适配器不自动关闭流或重置/复制摘要，不分配传输缓冲区，也不增加同步。流与哈希访问应串行化，调用期间不要修改写入输入。基础 hash 包仍只依赖 bytes；仅 stream 子包增加 io 依赖。

```goml
use std::bytes;
use std::hash;
use hash::Hasher;
use std::hash::fnv;
use std::hash::stream;
use std::io;
use io::Write;

fn hash_output[W: Write](writer: W, data: Slice[byte]) -> Result[bytes::Bytes, io::Error] {
    let digest = fnv::Digest::new64a();
    let output = stream::Writer::new(writer, digest);
    output.write_all(data)?;
    output.flush()?;
    Ok(digest.sum())
}
```

```goml
use std::bytes;
use std::hash;
use hash::Hasher;
use std::hash::crc32;

fn checksum_parts() -> u32 {
    let digest = crc32::Digest::new(crc32::Table::new(crc32::IEEE));
    digest.update(bytes::Bytes::from_string("1234").as_slice());
    digest.update(bytes::Bytes::from_string("56789").as_slice());
    digest.sum32()
}
```

<a id="slice-algorithms-and-checked-vector-edits"></a>

## 切片算法与带检查的向量编辑

纯 `std::collections` 切片辅助函数操作显式 `Slice[T]` 和 `MutSlice[T]` 视图。它们不引入第二种集合类型，也不改变 builtin 存储契约。所有复制都是浅拷贝：条目存储可独立，但引用对象仍共享。

| 能力 | API 与契约 |
| --- | --- |
| 相等与比较 | `slice_equal`、`slice_equal_by`、`slice_compare`、`slice_compare_by`；字典序比较返回 `cmp::Ordering`，`_by` 支持左右不同元素类型 |
| 成员查询 | `slice_index`、`slice_index_by` 返回首个 `Option[isize]`；`slice_contains` 和 `slice_contains_by` 返回 bool |
| 复制/压缩 | 现有 `Slice::to_vec`、`Vec::copy`、`Vec::dedup` 和 `dedup_by`；相邻压缩保留每段相等元素的首个元素 |
| 连接/重复 | `slice_concat(Vec[Slice[T]])`、`slice_repeat(Slice[T], count)` 返回经检查的独立 `Vec` 存储 |
| 区间编辑 | `vec_replace(Vec[T], start, end, replacement)`、`vec_insert_slice(Vec[T], index, inserted)`、`vec_delete_range(Vec[T], start, end)` 返回 `Result[(), SequenceError]` |
| 谓词删除 | `vec_delete_if(Vec[T], predicate)` 移除匹配元素，保留剩余顺序并返回移除数量 |
| 容量预留 | `vec_grow(Vec[T], additional)` 预留前检查负数量和长度溢出；不改变长度或值 |
| 视图修改 | `slice_reverse(MutSlice[T])`、`slice_sort(MutSlice[T])`、`slice_sort_by(MutSlice[T], compare)` 仅影响视图区间；排序稳定，使用 O(n) 临时存储 |
| 有序性/搜索 | `slice_is_sorted`、`slice_is_sorted_by`、`slice_binary_search`、`slice_binary_search_by`；搜索为首个相等值返回 `(insertion_index, found)` |
| 极值 | `slice_min`、`slice_max`、`slice_min_by`、`slice_max_by` 对空视图返回 `None`，平局时返回首个胜出值 |
| 迭代 | `slice_all` 和 `slice_backward` 产生 `(relative_index, element)`；只需值时使用现有 `Slice::iter` |
| 分块 | `slice_chunks(view, positive_size)` 返回经检查的单次遍历迭代器，产生有界只读子视图；最后一块可较短 |
| 收集 | `vec_append_iter(iterator, destination)`、`slice_collect(iterator)`、`slice_sorted(iterator)` 和稳定的 `slice_sorted_by(iterator, compare)` 接受具有对应元素类型的任意 `Iterator`；迭代器优先的参数顺序在检查目标前确定关联元素类型 |

`SequenceError` 区分 `InvalidRange(start, end, original_length)`、`InvalidCount(count)` 和 `LengthOverflow`。负的重复/预留数量、非正块大小和非法编辑区间都是可恢复错误。带检查的长度使用受支持的 Linux amd64 `isize` 范围。算术验证不保证从分配失败中恢复。空切片重复任意非负次数仍为空，不会执行相应次数的循环。

区间编辑在写入前完整构建浅拷贝结果，允许替换/插入输入与目标任意部分重叠，甚至整个向量。索引或长度错误保持目标不变。编辑修改共享 `Vec` 句柄，因此其他别名可观察新长度和条目；已有 `Slice` 头保留原有存储/长度语义。编辑使用 O(结果长度) 临时存储。谓词删除原地压缩；谓词不得结构性修改向量。

切片迭代器和分块保留原有有界视图，在消费时读取元素，不是不可变值快照。因此可观察到底层存储的写入，后续向量增长则不会扩展视图。迭代器只支持单次遍历，不同步并发访问。即使块大小为机器最大值，块边界也避免溢出。将向量自身的普通迭代器追加回向量，会将原先捕获长度的元素追加一次。调用方必须确保给定迭代器可终止，并避免在比较回调期间修改内容。

将带索引的迭代器赋给另一变量会共享游标；需要独立遍历时，再次调用 `slice_all` 或 `slice_backward`。每个返回块都是实时视图，不是副本。相等性检查在长度不等时拒绝匹配，不调用比较器，并在首个不等元素对处停止。谓词删除按顺序访问原始元素，保留剩余元素顺序，报告移除数量；回调 panic 不保证回滚此前压缩写入。

与 Go 切片不同，GoML 公共视图不能重新切片到其长度之外，也不能追加，且没有可公开裁剪的容量。使用 `view.to_vec()` 获得独立、长度匹配的副本，而不是 Go `Clip` 头操作；这也使该视图不再保留视图外的元素。泛型有序辅助函数要求 `cmp::Ord`；浮点调用方必须选择显式比较器/NaN 策略。`slice_equal` 保留普通 `PartialEq` 行为，包括 NaN 不相等。有序搜索要求输入已按给定比较器排序，使用 O(log n) 次比较，不做 O(n) 验证遍历。

```goml
use std::collections;

fn edit(values: Vec[isize]) -> Result[(), collections::SequenceError] {
    collections::vec_insert_slice(values, 0, values.as_slice())?;
    collections::slice_sort(values.as_mut_slice());
    let chunks = collections::slice_chunks(values.as_slice(), 2)?;
    for chunk in chunks {
        println(chunk.len());
    }
    Ok(())
}
```

<a id="map-algorithms"></a>

## 映射算法

导入 `std::collections` 可使用纯 GoML 映射算法。`map_clone` 创建独立条目存储，键和值为浅拷贝：嵌套 `Ref`、`Vec` 和其他共享对象仍共享。`map_copy(destination, source)` 覆盖源中的键，同时保留目标独有条目；将映射复制到自身是安全的。`HashMap` 没有 nil 映射状态。

`map_equal` 比较键是否存在以及值的 `PartialEq`，区分缺失键和存在的零值。`map_equal_by` 接受不同值类型的映射和 `(A, B) -> bool` 比较器。比较可提前停止；回调不得修改任一被比较映射。键使用现有 `Hash + Eq` 契约，包括一致哈希和自反相等性，而非 Go 原始映射键规则。浮点 NaN 等非自反值在普通 `PartialEq` 下仍不相等。

`map_all`、`map_keys` 和 `map_values` 返回调用时捕获的单次遍历 `FnIterator` 快照，顺序未指定。后续条目修改不改变快照，但引用对象仍共享。迭代器保留 O(n) 快照存储；它们不是实时 Go range 迭代器，也不对并发访问同步。`map_insert(destination, iterator)` 和 `map_collect(iterator)` 消费元素为 `(K, V)` 的任意 `Iterator`，按输入顺序覆盖重复键。收集总是分配新映射。

`map_delete_if(map, predicate)` 对原始快照每个条目调用一次 `(K, V) -> bool` 谓词，从当前映射移除选中键，返回实际移除数量。回调期间插入的键不访问。若回调替换了选中键，替换值会被移除；若已删除该键，不重复计数。谓词顺序未指定。并发修改仍需调用方同步。

长度相等但键不相交的映射不调用值比较器；比较器返回 false 时立即停止比较。快照删除仍访问原始条目，即使其键已被较早回调移除。插入消费到迭代器首次 None，并在迭代器回调返回后应用每个产生的键值对，因此该键值对会覆盖回调对同一键的写入。现有快照迭代器可安全地向其来源映射提供条目。这些辅助函数不限制迭代器元素数量，不捕获回调 panic，不回滚此前写入，也不提供同步；调用方必须限制不可信流，并在键存储期间保持哈希/相等行为稳定。

```goml
use std::collections;

fn main() -> () {
    let counts = HashMap::from_array([("a", 1), ("b", 2)]);
    let copy = collections::map_clone(counts);
    let keys = collections::map_keys(copy);
    copy.set("c", 3);
    for key in keys {
        println(key);
    }
    let _ = collections::map_delete_if(copy, |_: string, count: isize| count < 2);
    println(collections::map_equal(counts, copy));
}
```

<a id="immutable-collections"></a>

## 不可变集合

导入 `std::collections::imm` 可使用持久化 `HashMap[K, V]`、`HashSet[T]`、`Vector[T]`、`List[T]`、`BTreeMap[K, V]` 和 `BTreeSet[T]`。持久化意味着更新后旧版本仍可用，并不表示存储到磁盘。插入、移除和替换返回新集合，与原集合共享未变节点。赋值集合以 O(1) 空间和时间保留其版本。要使用新版本，必须保留更新的返回值。

```goml
use std::collections::imm;

fn main() -> () {
    let original: imm::HashMap[string, isize] = imm::HashMap::new();
    let first = original.insert("count", 1);
    let second = first.insert("count", 2);
    println(original.is_empty());
    println(first.get("count").unwrap_or(0));
    println(second.get("count").unwrap_or(0));

    let values = imm::Vector::from_vec(Vec::from_array([10, 20]));
    let appended = values.push(30);
    println(values.len());
    println(appended.get(2).unwrap_or(0));

    let ordered = imm::BTreeSet::from_vec(Vec::from_array([3, 1, 2]));
    for value in ordered.range(1, 3) {
        println(value);
    }
}
```

这六种类型均提供 `new`、`from_vec`、`len`、`is_empty`、`iter`、`Default` 和 `IntoIterator`。映射接受 `Vec[(K, V)]`；序列和集合接受 `Vec[T]`。映射构建按输入顺序处理条目，相等键的后值替换前值。集合丢弃重复项。List 和 Vector 构建保留输入顺序。`from_vec` 使外层集合与后续输入修改隔离。序列/集合的 `to_vec` 和映射的 `entries`、`keys`、`values` 返回独立可变向量。

| 类型 | 操作 | 表示与开销 |
| --- | --- | --- |
| `HashMap[K, V]` | `get`、`contains_key`、`insert`、`remove`、`entries`、`keys`、`values` | 基于 64 位哈希的位图索引 32 路哈希树；更新复制路径和受影响的冲突桶 |
| `HashSet[T]` | `contains`、`insert`、`remove`、`to_vec`、集合代数 | 复用以 unit 为值的 HashMap 实现 |
| `Vector[T]` | `get`、`set`、`push`、`pop`、`first`、`last`、`to_vec` | 带尾缓冲区的 32 路索引树；索引访问和路径复制更新需要 O(log₃₂ n) 的树操作，节点复制有界 |
| `List[T]` | `first`、`push_front`、`pop_front`、`tail`、`reverse`、`to_vec` | 共享尾部的单链表；头部访问、头部插入和取尾表为 O(1)；反转和遍历为 O(n) |
| `BTreeMap[K, V]` | 映射操作，加 `first`、`last`、`lower_bound`、`upper_bound`、`range` | 复制更新路径及受影响的分裂/借用/合并节点的 B 树；查找和单键更新为 O(log n) |
| `BTreeSet[T]` | 集合操作，加 `first`、`last`、`lower_bound`、`upper_bound`、`range` | 复用以 unit 为值的 BTreeMap 实现 |

哈希集合要求 `Hash + Eq`。哈希/相等比较的工作量在树遍历之外另计；完整哈希相同的条目使用逐个检查相等性的桶，因此冲突查找和更新可能与桶大小成线性。哈希集合迭代顺序未指定。树的键/元素要求 `cmp::Ord`，树迭代按升序。比较相等的键标识同一树条目。两类映射在使用等价键替换值时都保留原先存储的键。移除不存在的键/元素返回不变版本。

映射 `get` 返回 `Option[V]`。Vector `get(index)` 返回 `Option[T]`，`set(index, value)` 返回 `Option[Vector[T]]`；负索引及大于等于 `len()` 的索引返回 `None`。`push` 追加一个元素。Vector `pop` 和 List `pop_front` 返回 `Option[(T, Self)]`，将移除值与剩余集合配对。List `tail` 返回 `Option[List[T]]`。空集合的这些移除/访问操作返回 `None`，包括可用时的 `first` 和 `last`。

树的 `lower_bound(key)` 返回首个大于等于键的条目/元素；`upper_bound(key)` 返回首个严格更大的条目/元素。映射返回 `Option[(K, V)]`，集合返回 `Option[T]`。`range(start, end)` 返回半开区间 `[start, end)` 上的惰性迭代器。相等或逆序边界产生空迭代器。边界定位为 O(log n)；访问 m 个元素的区间需要 O(log n + m) 遍历工作和 O(log n) 游标空间。

两种集合均提供 `union`、`intersection`、`difference`、`symmetric_difference`、`is_subset` 和 `is_disjoint`。这些操作比较成员关系并构建持久化结果，不改变任一输入；不保证专门的子树合并复杂度。

`iter()` 返回单次遍历 `FnIterator`，保留集合版本，不先复制所有元素。派生新版本不影响旧版本的迭代。分别调用 `iter()` 具有独立游标；同一迭代器的副本共享可变游标，并发使用需要同步。条件 `PartialEq` 和 `Eq` 实现提供内容相等性：序列顺序有意义，映射/集合插入历史无关。相等比较使用元素/值的 trait，不使用私有节点身份。不提供集合 `Hash` 和 serde 实现。

不可变性是浅层的。元素和值保留普通复制或共享句柄语义，因此 `imm::Vector[Ref[T]]` 仍允许修改被引用的值。哈希键存储期间必须保持稳定的哈希/相等性，树键必须保持稳定顺序。没有 API 暴露可变节点存储。独立不可变版本可在不修改结构的情况下读取；可变载荷的并发访问仍需要载荷自身的同步。

`FrozenVec` 仍是紧凑、连续、具有常量时间索引的不可变快照；`imm::Vector` 支持通过结构共享派生更新版本。持久化集合目前不提供可变临时构建器、深度冻结、任意 Vector 插入/连接，或 List 惰性求值。保留旧版本也会保留其可达的共享节点和载荷。所有算法使用现有 GoML 设施，不增加语法或运行时内建操作。编译器和驱动使用方保留现有集合，直到发布并推进 stage0 后这些 API 可用。

<a id="iterator-adapters"></a>

## 迭代器适配器

新代码应导入 `std::iter`，用于迭代器构建、适配和消费：

```goml
use std::iter;

let values = iter::map(Vec::from_array([1, 2, 3]).iter(), |value: isize| value * 2);
let total = iter::fold(values, 0, |sum: isize, value: isize| sum + value);
```

- 生产者：`empty`、`once` 和 `from_fn`
- 适配器：`map`、`filter`、`filter_map`、`take`、`take_while`、`skip`、`skip_while`、`enumerate`、`zip`、`chain`、`inspect` 和 `map_while`
- 消费操作：`fold`、`collect`、`find`、`find_map`、`any`、`all`、`count`、`position`、`nth`、`last`、`for_each` 和 `reduce`

<a id="sorted-collection-boundaries"></a>

## 有序集合边界

`std::collections::sort` 和 `stable_sort` 保留现有稳定 GoML 归并排序契约：相等元素保留相对顺序，排序修改共享向量。现有 builtin 源码实现仍可用于 stage0 兼容；不引入 Go 排序后端。泛型默认排序使用 `cmp::Ord`；浮点数要求显式比较器和 NaN 排序策略，不采用隐式全序。

`is_sorted`、`is_sorted_by` 和 `is_sorted_by_ordering` 检查相邻元素对，接受空和单元素向量。整数比较器按参数顺序返回负数/零/正数；`Ordering` 比较器返回 `Less`、`Equal` 或 `Greater`。它们必须定义一致顺序，不得修改输入。

`lower_bound(values, expected)` 查找首个不小于目标的元素；`upper_bound` 查找首个大于目标的元素。两者均返回 `0..=len` 内索引，包括目标不存在时的插入点。`equal_range` 返回半开重复区间 `(lower, upper)`。其 `_by` 形式将一个元素与捕获的目标比较；三者也都有 `_by_ordering` 变体，包括 `equal_range_by_ordering`。现有 `binary_search` 仍返回首个匹配索引或 `None`，不改为插入点结果。提供一致的反向比较器即可支持降序。输入必须已按该比较器排序；边界搜索使用 O(log n) 次比较和 O(1) 额外存储，不重新扫描验证有序性。

`search(length, predicate)` 搜索抽象索引区间，不分配向量。谓词随索引增加必须先 false 后 true。结果为 `Some(first_true_index)`，若全为 false 则为 `Some(length)`。负长度不调用回调，返回 `None`；零长度不调用回调，返回 `Some(0)`。中点算术避免溢出，包括机器最大长度。这些辅助函数从不在 `length` 处调用回调。

`search_by(length, compare)` 返回 `Option[(isize, bool)]`：首个比较结果非负的索引（或 `length`），以及该索引是否比较相等。负长度返回 `None`；空区间返回 `Some((0, false))`，均不调用比较器。`search_by_ordering` 接受 `cmp::Ordering` 而非整数比较。回调将索引处元素与捕获目标比较，与 `lower_bound_by` 一致，与 Go `sort.Find` 的目标对元素约定相反。其符号必须形成负、零、正三个连续区域，任一区域均可不存在。候选索引处可能再次调用以检查相等，绝不越出 `0..length`。这些辅助函数使用 O(log n) 次比较和常量空间，不实体化序列；例如，`collections::search_by(100, |index: isize| index * 2 - 37)` 返回 `Some((19, false))`。比较器结果必须保持一致，回调不得修改被搜索序列。

所有排序变体执行 O(n log n) 次比较，使用 O(n) 临时存储，即使名称不含 `stable_` 也保留相等键顺序。比较器必须定义严格弱序，排序期间不得修改向量、其别名或排序键。比较器 panic 可能保留此前已提交的归并轮次；排序不保证事务回滚或同步。共享向量别名和现有元素视图会观察到最终顺序。对于接近有符号边界的整数键，应使用关系比较或 `cmp::compare`，不要使用可能溢出的减法。

```goml
use std::collections;

fn main() -> () {
    let values = Vec::from_array([1, 2, 2, 4]);
    let (start, end) = collections::equal_range(values, 2);
    println(end - start);
    println(collections::lower_bound(values, 3));
    println(collections::is_sorted(values));
}
```

<a id="logical-slash-paths"></a>

## 逻辑斜杠路径

`std::path::slash` 独立于现有宿主文件系统 `std::path` API。它仅将 `/` 视为分隔符；反斜杠和盘符前缀是普通文本。所有操作都不查询工作目录、不解析符号链接、不执行 I/O，也不提供安全沙箱。所有输入都是合法 UTF-8 GoML 字符串。

`clean` 移除空组件和 `.`，按词法解析 `..`。它保留相对路径开头的 `..`，阻止越过词法 `/` 根，空相对结果返回 `.`。`join(Vec[string])` 跳过空元素，以 `/` 连接，并清理结果；全空输入返回空字符串。后续看似绝对路径的元素不会丢弃此前元素。`split` 原样保留输入，返回 `(directory_including_last_slash, final_component)`。`base` 去掉尾随斜杠，空输入返回 `.`，全斜杠返回 `/`。`dir` 清理 `split` 返回的目录。`extension` 返回最后组件中最后一点号起始的后缀，包含点号；不去掉尾随斜杠，也不特殊处理点文件。

`matches(pattern, name)` 匹配整个名称。`*` 匹配零个或多个非斜杠字符，`?` 匹配一个非斜杠 Unicode 标量，`[...]` 指定标量区间，`[^...]` 取反，`\\` 转义下一个标量。字符类不能为空；类内的字面量 `-` 或 `]` 需要转义。降序区间不匹配任何内容。与 Go 的 `path.Match` 一样，字符类可显式或通过否定匹配 `/`，不同于 `*` 和 `?`。没有递归 `**` 扩展。畸形模式返回带零起始 Unicode 字符位置的 `MatchError::InvalidPattern`，即使名称原本也不匹配。

`Pattern::compile` 解析一次，用于重复、独立匹配。`compile_with_limit` 设置模式标量最大数量；默认 `DEFAULT_PATTERN_LIMIT`（4096）。`matches_with_limit` 设置工作预算；`matches` 使用 `DEFAULT_WORK_LIMIT`（1000000）。负限制是错误。匹配对每个初始 token、输入标量、token 转移和字符区间比较计一个单位。它使用与模式大小成线性的状态和多项式动态规划，不使用递归回溯；限制耗尽返回 `PatternLimit` 或 `WorkLimit`，绝不返回 false 匹配结果。调用方选择更大限制，以更多资源使用换取更大输入。Pattern 字段私有，匹配不修改共享编译状态。

匹配总在 Unicode 标量边界前进。这有意区别于 Go 1.26 path.Match 在星号后的逐字节重试：Go 通过将不完整 UTF-8 后缀解码为替换字符，使 `*[�]` 和 `*??` 都能匹配 `世`，而 GoML 对两者均返回 false。真正的替换标量，如 `世�` 中的标量，可匹配 `*[�]`，两个标量如 `世界` 可匹配 `*??`。这些差异经过显式测试，不将内部字节位置解释为字符。连续星号折叠为一个匹配 token，但对源模式标量限制仍逐个计数。每次区间比较都计费，即使当前状态不能匹配；预算失败的调用不影响经同一 Pattern 或别名进行的后续调用。

动态规划匹配器探索所有合法 token 转移。特别地，`*[a/]*` 匹配 `a/`：首个星号消耗 `a`，字符类消耗 `/`，最后星号为空。Go 1.26 path.Match 则在较早的字符类匹配消耗 `a`、余下星号无法消耗 `/` 后拒绝此情况。GoML 在此遵循文档定义的全名称模式语义，不复现该搜索顺序行为。

```goml
use std::path::slash;

fn matches_docs(name: string) -> Result[bool, slash::MatchError] {
    let pattern = slash::Pattern::compile("docs/*.md")?;
    pattern.matches(name)
}

fn main() -> () {
    println(slash::clean("docs/../images//icon.png"));
    println(slash::join(Vec::from_array(["docs", "guide", "../index.md"])));
    match matches_docs("docs/index.md") {
        Ok(matched) => println(matched),
        Err(error) => println(error.to_string()),
    }
}
```

<a id="base32"></a>

## Base32

`std::encoding::base32::stream` 提供 `Encoder[W: io::Write]` 和 `Decoder[R: io::Read]`。两个构造函数都接受底层句柄、Encoding 和累计输出字节限制，返回 `Result[..., io::Error]`。此子包依赖 I/O；核心 Base32 包保持独立于 I/O。

writer 实现 Write 和 Close。`finish()` 或 `close()` 恰好输出一次尾部和填充，不刷新或关闭底层 writer。`flush()` 只转发刷新，不结束待处理位。finish 后写入以 BrokenPipe 失败。输出限制失败保持提交输入未消耗，允许以更小输入重试。底层 write/flush 失败会持续保存，因为输出可能已外泄：后续操作不重放它。`failure()` 保留原始错误；终止性的 Interrupted 对外呈现为 Other，为兼容保留历史终止错误映射。默认 `write_all` 不再重试 Interrupted。`input_len()` 统计状态已接受输入，不代表目标已确认交付。`is_finished()` 要求成功完成且没有终止失败。

reader 增量消费编码字节，最多保留一个解码量子，源 Interrupted 不重试直接返回，解码前验证原始 ASCII。源、解码和限制错误会持续保存，可通过 `failure()` 获取，并保留源错误原始详情。此前返回的字节无法回滚。需一直读到 EOF 才能验证整个输入，包括填充后的尾随字节；只获取预期解码长度还不够。空读取不消耗输入，但仍报告此前记录的失败。句柄副本共享状态，需要外部同步。限制约束累计输出，不约束进程总内存；writer 更新分配与提交输入块成比例的输出。

`std::encoding::base32::{encode, decode}` 使用 RFC 4648 大写字母表和 `=` 填充。`encode_with`/`decode_with` 接受 `Variant::Standard`、`StandardNoPadding`、`Hex` 或 `HexNoPadding`；Hex 表示扩展 Base32 字母表 `0123456789ABCDEFGHIJKLMNOPQRSTUV`，不是十六进制字节编码。输入字节可为任意二进制数据。

解码严格且要么全成要么报错：标准变体拒绝小写、空白/换行、错位或多余填充、非法长度、非法数字和非零的未使用尾部位。`DecodeError::index()` 报告字节偏移，`kind()` 区分 `InvalidLength`、`InvalidByte`、`InvalidPadding` 和 `NonCanonical`。与 Go Base32 解码器不同，不静默移除 CR/LF，也不返回部分输出。这些一次性 API 实体化完整结果；增量状态 API 输出分块，独立的 `base32::stream` 包提供 I/O 适配器。

`Encoding::new(alphabet: string, padding: Option[byte]) -> Result[Encoding, ConfigError]` 构建可复用的不可变配置与解码表。字母表必须恰好包含 32 个不同 ASCII 字节，不含 CR/LF；可选填充字节也必须是 ASCII，不含 CR/LF，且不出现在字母表中。此显式 ASCII 限制保证每个编码结果都是合法 GoML 字符串，比 Go 的任意字节字母表更严格。可以有意配置 NUL、空格和制表符，此时它们是字面符号，绝不会作为空白忽略。`ConfigError` 区分非法长度、非法/重复符号字节偏移和非法填充。

`Encoding::from_variant(variant)` 提供既有字母表和填充设置。`alphabet()` 和 `padding()` 检查配置；`with_padding(option)` 返回新验证的配置，不改变原配置。`.encode(bytes)` 和 `.decode(text)` 共享标准算法和规范尾部位检查。无填充时，`=` 可为字母表数字；否则孤立 `=` 仍被拒绝。例如，`Encoding::new("abcdefghijklmnopqrstuvwxyz234567", Some(b'!'))` 创建小写字母表和感叹号填充。自由函数与 Variant 保留此前结果和错误。这些 API 使用普通结构体、枚举、方法和泛型容器，不改变语法。

`Encoding::encoded_len(input_length)` 计算精确的带填充或原始输出字节长度，不分配输出，返回 `Result[isize, OutputError]`；负长度和算术溢出是可恢复错误。`.encode_checked(input, max_bytes)` 在编码前拒绝负限制和过大/溢出输出。`.decode_checked(input, max_bytes)` 在分配输出缓冲区前检查填充/长度形态和解码结果长度，再执行通常的数字和规范位验证。恰好达到限制可成功；失败不返回部分结果。自由函数 `encode_checked` 和 `decode_checked` 使用标准填充配置。自定义字母表使用其 Encoding 值上的方法。

`OutputError` 区分 `InvalidLength(length)`、`InvalidLimit(limit)`、`LengthOverflow`、`LimitExceeded(limit)` 和 `Decode(original_error)`。解码错误优先级为负限制、非法形态、输出超限、非法数字/尾部位；因此过大的畸形主体可能先报告限制错误，再有数字错误。限制约束输出存储，不约束输入长度或进程总内存。

`Encoder::new(encoding, max_bytes) -> Result[Encoder, OutputError]` 创建具有累计输出限制的增量编码状态。`update(input: Slice[byte]) -> Result[string, StreamError]` 输出当前输入和此前待处理位中所有完整的五位符号；应按调用顺序连接返回块。它仅保留标量位累加器和计数器，不保留输入切片或已输出块。调用可以在输入任意位置拆分。`finish() -> Result[string, StreamError]` 输出最终补零符号和配置的填充，然后关闭编码器；重复 finish 成功返回空字符串。finish 后 update 即使为空输入也返回 `StreamError::Closed`。

每次 update 在分配输出或改变状态前，检查全部已接受输入的最终编码长度，包括 finish 所需填充。`StreamError::Output(error)` 保持整个提交块未消耗、状态不变，因此调用方可以重试更小块。每次成功 update 都会预留足够完成的预算。`input_len()` 报告已接受输入字节；`output_len()` 报告已输出字节，不是预留最终大小。`is_finished()` 报告关闭状态。句柄副本共享状态；`reset()` 重置所有共享计数器和待处理位，重新打开编码器，保留配置和限制。调用需要外部同步。此纯增量状态 API 不依赖 I/O。

`Decoder::new(encoding, max_bytes) -> Result[Decoder, OutputError]` 限制累计解码字节。`update(input: string) -> Result[bytes::Bytes, StreamError]` 输出完整八符号组的解码结果，最多保留七个待处理 ASCII 字节。符号到达时立即检查；完整组使用与一次性解码相同的严格填充和尾部位验证。带填充的最后一组之后不允许任何非空输入。原始编码可以以合法的较短组结束。`finish()` 验证/输出该尾部并关闭解码器；重复成功 finish 返回空字节。失败的 finish 保持解码器打开，以便补充缺失后缀。成功 finish 后 update 返回 `Closed`，包括空输入。

每次 update 对提交块具有事务性：解码或限制失败不返回该块的任何字节，计数器、待处理输入和终止状态均不变。此前成功块不回滚。`StreamError::Output(OutputError::Decode(error))` 报告跨此前已接受块的绝对输入字节偏移；finish 时非法长度可能指向已接受输入末尾后一字节。限制包含待处理数据隐含的字节，因此合法待处理尾部能在预算内完成。错误优先级遵循增量发现，不一定与一次性形态优先验证一致。`input_len()`、`output_len()`、`is_finished()` 和共享句柄 `reset()` 遵循编码器约定。返回字节缓冲区独立，不保留输入字符串和已输出块。解码器调用需要同步。Reader/Writer 适配器位于独立 `base32::stream` 包；这些 API 不增加语法。

```goml
use std::bytes;
use std::encoding::base32;

fn token_text() -> string {
    base32::encode_with(bytes::Bytes::from_string("foobar"), base32::Variant::StandardNoPadding)
}
```

<a id="hexadecimal-and-base64"></a>

## 十六进制与 base64

`std::encoding::hex` 默认将 `bytes::Bytes` 编码为小写十六进制，`encode_upper` 则输出大写数字。解码接受两种大小写，返回 `DecodeError` 时包含奇数长度或非法数字的字节偏移。

`std::encoding::base64` 默认使用带填充的 RFC 4648 标准字母表。`Variant` 选择标准或 URL 安全字母表，并指定要求或省略填充。解码严格拒绝无效长度、混用字母表、位置错误的填充，以及末尾未使用位非零的情况。

<a id="pem-framing"></a>

## PEM 封装

`std::encoding::pem` 在标准 Base64 编解码器之上以纯 GoML 实现文本 PEM 封装。它不加密/解密私钥、不解析 ASN.1、不验证证书，也不实现信任策略。`Block::new(label, headers: Vec[(string, string)], data: bytes::Bytes) -> Result[Block, Error]` 验证输入并保存快照。`label()`、`headers()` 和 `data()` 返回值，其中可变数据是独立快照。标签必须是非空可打印 ASCII，且无首尾空格。头字段键必须是非空可打印 ASCII，不能含空格或冒号；值允许可打印 ASCII 和制表符，首尾 ASCII 空白会被规范化移除。换行和区分大小写的重复键会被拒绝。

`encode(block, max_bytes) -> Result[string, Error]` 输出以 LF 分隔的 BEGIN/END 行，以及每 64 列换行的带填充 Base64。头字段按字典序排序，`Proc-Type` 置于最前，之后跟一个空行。空载荷没有正文行。构造输出前会检查包含头字段和行尾在内的完整字节长度；恰好达到限制时成功，负限制和过大输出失败。不返回部分字符串。

`decode(input: string, max_input, max_data, max_headers) -> Result[Option[(Block, isize)], Error]` 在可选前导内容后寻找首个从行首开始的 BEGIN 候选项。成功时返回块及紧跟其 END 行之后的绝对字节偏移，若有终止换行则包含它。按该偏移切分原字符串可继续解析后续块。`None` 表示没有 BEGIN 候选项。接受 LF 和 CRLF、行末空格/制表符，以及 Base64 正文内的空格/制表符。首先检查整个输入字节限制；解码字节和头字段数量限制彼此独立。临时解析存储受输入限制约束。头字段数量在验证前统计，包含重复字段。

头字段规范化只移除首尾 ASCII 空格和制表符字节。键或值周围的 Unicode 空白会作为无效头字段文本被拒绝，不会在验证前静默移除。这使解码得到的块与显式构造的块遵循相同的可打印 ASCII 契约。

此实现有意比 Go 的搜索式解码器更严格：格式错误的候选项、不匹配或缺失的 END 标签、重复头字段和非规范 Base64 尾位都会返回错误，不会静默跳过。输入是有效 UTF-8 字符串，不是任意二进制前导内容。`Error::Malformed(offset)` 使用原始输入字节偏移；嵌套 Base64 错误的偏移相对于过滤空白后的正文。无效标签/头字段、重复名称、负限制和超出边界各有不同错误变体。解析器实体化一个块，不提供密码学真实性保证，目前也不是流式 I/O 适配器。

例如，`pem::decode(source, 1048576, 65536, 16)` 最多接受一 MiB 输入，其中的块最多含 64 KiB 解码内容和 16 个头字段。这些 API 使用现有泛型容器、结构体、枚举和调用表达，不新增语法或编译器内建函数。

独立的 [`ecosystem::asn1`](https://github.com/gomlang/asn1/blob/HEAD/README.md) 模块通过显式 `Value` 和 `Schema` 树提供有界的规范 DER/OID 编解码。它接受文档列明的 ASN.1 标签子集，要求调用方提供模式进行类型验证，不使用反射。其 `Limits` 限制字节数、序列深度、元素数和 OID 弧数。独立的 [`ecosystem::x509::pkix`](https://github.com/gomlang/x509/blob/HEAD/README.md) 包基于这些有界 DER 原语处理可分辨名称、算法标识符和扩展。它只表示证书数据；证书链验证、信任存储和 TLS 不属于其 API。独立的 [`ecosystem::xml`](https://github.com/gomlang/xml/blob/HEAD/README.md) 模块提供有界增量 XML token 读写器，支持命名空间解析和显式扁平结构体 serde 模式。它接受 UTF-8 XML 1.0，拒绝 DTD 和外部实体，不进行模式验证。对于目标文件和调试元数据，[`ecosystem::object::elf`](https://github.com/gomlang/object/blob/HEAD/README.md) 读取有界 ELF32/64 文件，[`ecosystem::dwarf`](https://github.com/gomlang/dwarf/blob/HEAD/README.md) 通过有界随机访问读取选定的 DWARF32 v2–v5 单元和行号表。各模块文档说明支持的子集和错误限制。

<a id="time"></a>

## 时间

`Duration` 存储非负纳秒数，提供纳秒、微秒、毫秒和秒的构造器及整单位访问器。减法在零处饱和。`Instant` 是单调时钟，适合测量经过时间。`SystemTime` 提供 Unix 纳秒、毫秒和秒时间戳。`time::sleep` 将当前 goroutine 阻塞一个 `Duration`。

`Duration` 提供检查式和饱和式的比例构造、加法、减法及乘法。检查式操作在溢出、下溢或输入为负时返回 `None`。饱和式操作钳制到零或最大有符号 64 位纳秒值。`Duration`、`Instant` 和 `SystemTime` 提供 `compare`；当接收者早于给定时刻时，`Instant::checked_duration_since` 返回 `None`。

对于日历时间和命名时区，独立的 [`ecosystem::datetime::tzdata`](https://github.com/gomlang/datetime/blob/HEAD/tzdata/README.md) 包嵌入带版本的 2026c TZif 快照。`names`、`contains`、`checksum`、`zone_bytes`、`load` 和 `load_versioned` 通过现有 `datetime::TimeZone` 解析器提供经过验证的时区数据。嵌入快照在重新生成前保持固定；加载不依赖宿主的 zoneinfo 目录。

<a id="files-and-standard-streams"></a>

## 文件和标准流

`fs::read_file_structured`、`read_bytes_structured`、`write_file_structured` 和 `write_bytes_structured` 执行整文件 I/O。它们返回 `fs::Error`，包含稳定的 `io::ErrorKind`、操作、路径、可选原始操作系统代码和显示消息。目录创建与删除、规范化、目录列举和文件哈希使用同一错误类型。

`fs::create_dir`、`rename`、`copy`、`hard_link` 和 `symbolic_link` 是立即执行的操作，返回 `fs::Error`。`copy` 读取并写入完整文件，目前以可移植的 `0644` 权限创建目标，不保留源元数据。

`fs::metadata` 和 `symlink_metadata` 返回纯值元数据，包含文件类型、长度、可移植权限位和修改时间。`Metadata::from_parts(file_type, length, mode, modified_unix_nanoseconds)` 无需访问文件系统即可构造相同的值；权限位被掩码为 `0o777`。`read_dir_structured` 立即对目录条目生成快照。`read_dir_names_structured` 返回排序后的名称，不获取每个子项的元数据，并保留结构化列举错误；消失的子项不会使其他名称无效。`atomic_write` 在返回前写入、同步、关闭同目录临时文件并原子重命名；临时文件清理在运行时调用内部完成。`replace` 以替换语义暴露宿主的原子重命名操作。

`io::read_stdin_structured`、`read_stdin_exact_structured`、`write_stdout_structured` 和 `write_stderr_structured` 为标准流提供结构化错误。`read_stdin_to_string` 验证完整输入的 UTF-8 有效性，失败时报告 `InvalidData`。负的精确读取长度在访问 stdin 前报告 `InvalidInput`。

<a id="portable-metadata-values-and-formatting"></a>

## 可移植元数据值和格式化

`fs::FileMode` 增加可移植值语义，不改变现有四分支 FileType 枚举。`NodeKind` 区分 File、Directory、Symlink、BlockDevice、CharacterDevice、NamedPipe、Socket、Irregular 和 Unknown。兼容方法 `file_type()` 直接映射前三者，其他种类都映射为 FileType::Other。这些是 GoML 值，不是 Go 或宿主 ABI 的模式位数值。

`ModeFlags` 有六个公共 bool 字段：set_uid、set_gid、sticky、append_only、exclusive 和 temporary。`FileMode::from_parts(kind, permissions, flags)` 接受 NodeKind、u32 权限位和 Option[ModeFlags]。权限掩码为 `0o777`；flags=None 表示这些属性未知，不同于包含六个 false 值的 Some。种类和标志是否已知彼此独立，因此可以表示已知 NamedPipe 但标志未知的情况。`FileMode::new(kind, permissions, flags)` 是标志已知时的便捷形式。`FileMode::from_file_type(file_type, permissions)` 映射旧分类并将 flags 设为 None，其中 Other 映射为 Unknown。访问器提供 `kind()`、`permissions() -> Permissions`、`flags()` 和兼容的 `file_type()`。NodeKind、ModeFlags 和 FileMode 实现 PartialEq 和 Eq。

`Metadata::from_mode(mode, length, modified_unix_nanoseconds)` 保留完整 FileMode，可通过 `.mode()` 取回。现有 file_type、len、permissions 和时间访问器保持原有含义。旧的 from_parts 构造器仍对权限应用掩码，现在显式记录未知标志。现有宿主查询只提供原有分类和九个权限位，无法恢复特殊属性或区分隐藏在 Other 后的高级种类。因此宿主元数据也报告未知标志，绝不虚构全 false 值。未扩展任何原生接口。

`SnapshotEntry::new(name, metadata)` 存储名称和元数据值，无需文件系统 I/O 即实现 DirectoryEntry。其 metadata 方法始终返回该快照；`.mode()` 直接提供完整模式。构造时不验证给定名称；正常目录收集仍拒绝无效组件名称。这是供提供者使用的公共值适配器，不是新的宿主目录后端，也不声称元数据始终与磁盘保持一致。

`format_metadata(name, metadata, max_bytes)` 和 `format_directory_entry(entry, max_bytes)` 返回 `Result[string, fs::Error]`。它们使用确定性、与区域设置无关且不带换行的显示格式：

```text
kind=file permissions=rw-r----- flags=unknown size=12 modified_ns=-7 name=example
kind=directory name=config/
```

元数据格式包含 snake_case 的高级种类、九个 rwx/短横线权限位置、标志、有符号字节长度和有符号 Unix 纳秒。已知标志显示为 `set_uid:0,set_gid:0,sticky:0,append_only:0,exclusive:0,temporary:0`，设置的标志对应值改为 1。未知标志使用单词 `unknown`。条目格式化在预算非负时只调用 name 和 file_type，各一次，绝不获取 metadata；其兼容种类标签是 file、directory、symlink 和 other。两种格式都为目录追加 `/`，即使给定名称已经以斜杠结尾。名称原样复制，不转义；这些显示字符串不是安全的日志或序列化格式。

负预算在调用任何条目方法前返回 InvalidInput。输出预算不足返回 InvalidData，不返回部分字符串。复制前逐字段检查剩余 UTF-8 字节额度，不在字符内部截断。错误操作为 `"format metadata"` 和 `"format directory entry"`；路径包含给定或获取的名称，但负的条目预算尚无名称。任意条目方法的 panic 遵循正常 panic 传播，不转换为错误。字节额度限制输出长度，不限制分配器容量或进程内存，也不会使分配失败可恢复。

```goml
use std::fs;

fn describe_pipe() -> Result[string, fs::Error] {
    let mode = fs::FileMode::from_parts(NamedPipe, 0o600, None);
    let metadata = fs::Metadata::from_mode(mode, 0, 0);
    fs::format_metadata("events", metadata, 256)
}
```

<a id="portable-read-only-filesystems"></a>

## 可移植只读文件系统

`fs::File` 在 `io::Read + io::Close` 基础上增加 `stat(self: Self) -> Result[fs::Metadata, fs::Error]`。`fs::FileSystem` 有关联类型 `Handle: File` 和 `open(self: Self, name: string) -> Result[Self::Handle, fs::Error]`。实现提供自己的存储，每次成功打开均返回新的逻辑打开句柄。读取遵循普通 `io::Read` 的数量/EOF/错误契约；元数据检查不得移动读取位置。句柄副本可以共享位置和关闭状态；这些 trait 不隐含并发安全性。直接调用 `open` 时，提供者也必须拒绝无效逻辑路径；泛型辅助函数在分派给提供者前也会验证。这些可移植协议与上述宿主路径函数共存，不引入原生文件后端，也不将 OS 文件重新解释为内存快照。

`fs::valid_path(name)` 接受表示根的 `"."`，或由斜杠分隔的非空组件，组件不能为 `"."` 或 `".."`。它拒绝绝对路径、尾随斜杠和空组件，不规范化输入。反斜杠、冒号和 NUL 是普通逻辑名称字符；具体提供者可以拒绝其存储无法表示的名称。字符串使用 GoML 的 UTF-8 表示。验证是词法层面的，不是 OS 路径安全或沙箱保证。

`fs::read_file_from(filesystem, name, max_bytes)` 返回 `Result[bytes::Bytes, fs::Error]`；`fs::stat_from(filesystem, name)` 返回 `Result[fs::Metadata, fs::Error]`。无效名称和负字节限制在打开前失败。打开成功后，每个辅助函数在正常完成（包括失败）时恰好关闭句柄一次；延迟清理也会在 panic 展开时尝试关闭。主操作错误优先于返回的关闭错误；操作成功后的关闭错误会被返回。提供者的 open/stat 错误保持不变。read/close 错误保留种类、原始 OS 代码和消息，使用请求的逻辑名称及操作 `"read file"`/`"close file"`。提供者 panic 遵循普通 panic/defer 规则，不转换为返回错误。

读取使用 `read` 原语，不使用提供者可能重写的 `read_to_end` 辅助函数。它检查数量，遇到任何错误（包括 `Interrupted`）立即停止，并最多多读限制之外的一个字节，以区分恰好 EOF 与超量输入。超量输入或无效数量返回 `InvalidData`；失败时不返回部分字节。不信任元数据大小作为读取边界。限制约束返回字节数，不约束全部分配或提供者副作用；这些辅助函数不会使分配失败可恢复。

`fs::SubFs::new(filesystem, root)` 验证逻辑根，不打开它，也不检查它是否存在或是目录。其 `FileSystem` 实现验证每个请求名称，再添加根前缀，不清理或改写组件。根 `"."` 是恒等适配；打开 `"."` 选择配置的根。嵌套 SubFs 值可以组合。底层句柄和提供者错误原样返回，因此错误可能提及添加前缀后的路径。SubFs 是命名空间适配器，不提供隔离约束：符号链接或提供者策略可以解析到该命名空间之外。

```goml
use std::fs;
use fs::FileSystem;
use std::bytes;

fn load_settings[F: FileSystem](filesystem: F) -> Result[bytes::Bytes, fs::Error] {
    let config = fs::SubFs::new(filesystem, "config")?;
    fs::read_file_from(config, "settings.toml", 65536)
}
```

`fs::DirectoryEntry` 描述稳定的 `name()` 和 `file_type()`，以及可能惰性的 `metadata() -> Result[Metadata, Error]`。读取后续页和关闭目录后，条目必须仍可使用；不能借用枚举器会复用的存储。条目消失或权限变化时元数据可能失败。现有宿主 `DirEntry` 实现此 trait：名称/类型保留列举快照，metadata 在存储的路径上调用 `symlink_metadata`。现有立即执行的 `read_dir_structured` API 不变。

`fs::ReadDirFile: File` 提供 `type Entry: DirectoryEntry` 和 `read_dir(amount: isize) -> Result[Vec[Self::Entry], ReadDirError[Self::Entry]]`。提供者必须在移动游标前以 `InvalidInput` 拒绝非正数量。成功页最多包含 amount 个条目；空页表示该次调用遇到 EOF，短的非空页则不表示 EOF。错误可以包含已确认消费的条目。该 trait 不隐含重试或持续错误行为。名称必须是单个有效组件，不含 `"."` 和 `".."`；NUL、反斜杠和冒号仍是普通逻辑字符。

`ReadDirError::new(entries, error)` 复制向量结构，不复制条目对象。`.entries()` 返回只读切片，`.error()` 返回原始原因。其 display/debug 表示使用原因，不对条目类型施加格式化约束。失败的提供者调用具有未知副作用时，不会虚构条目来表示这些副作用。

对于满足 `F: FileSystem` 且 `F::Handle: ReadDirFile` 的 F，`fs::read_dir_from(filesystem, name, max_entries)` 返回 `Result[Vec[F::Handle::Entry], ReadDirError[F::Handle::Entry]]`。它在打开前检查路径和非负预算，然后请求每页最多 128 个条目。接近限制时多请求一个条目以区分 EOF 和超量数据；零预算仍会打开并探测一个条目。探测项可能被消费但不返回。预算限制结果数量，不限制总分配、名称长度或提供者副作用，也不承诺一致目录快照或使分配失败可恢复。

每页都在接受其中任何条目前完成验证。条目过多或任何名称无效（包括探测项）都会以 `InvalidData` 拒绝整页，并保留之前的页。过大的页在调用条目方法前被拒绝。否则每个名称获取并缓存一次；此辅助函数从不调用 metadata 和 file_type。提供者错误（包括 `Interrupted`）立即停止，保留原始原因以及预算内的有效条目。对于有效错误页，提供者原因优先于超出限制；无效页错误优先于二者。成功但数据超量时则产生 `InvalidData`。保留的前缀按提供者顺序选择，再按缓存名称稳定排序；重复名称保留。成功和所有错误都返回已排序的确认结果，绝不超过预算。

每个成功打开的句柄均关闭一次，包括 panic 清理期间。操作错误优先于普通关闭错误。只有关闭错误时保留完整列表，并使用与 `read_file_from` 相同的错误映射。关闭 panic 遵循正常 panic 传播，可能替代之前的 panic，不转换为普通关闭错误。SubFs 转发其句柄类型，因此嵌套目录组合无需单独适配器或动态接口检测。例如：

```goml
use std::fs;
use fs::{FileSystem, ReadDirFile};

fn list_config[F: FileSystem](filesystem: F) -> Result[
    Vec[F::Handle::Entry], fs::ReadDirError[F::Handle::Entry],
] where F::Handle: ReadDirFile {
    fs::read_dir_from(filesystem, "config", 1024)
}
```

`fs::walk_from(filesystem, root, limits, visit)` 使用非递归方式执行有界、稳定、按字典序排列的深度优先遍历。它返回 `Result[(), fs::Error]`，要求与 `read_dir_from` 相同的 FileSystem/ReadDirFile 约束。回调接受逻辑路径、`Option[fs::WalkEntry[F::Handle::Entry]]` 和 `Option[fs::Error]`，返回 `Result[fs::WalkControl, fs::Error]`。WalkEntry 实现 DirectoryEntry：名称/类型被缓存，根元数据来自初始 stat，子项元数据保持惰性。遍历器不获取子项元数据。缓存名称与目录验证和排序共用，因此构造子路径时不会再次调用提供者的 name 方法。仅对实际访问的子项获取文件类型，各一次。

`WalkControl::Continue` 访问目录的子项；`SkipDir` 跳过它们。对于非目录，SkipDir 跳过其父目录中剩余的兄弟项，不跳过祖先层级的剩余条目。`SkipAll` 成功停止整个遍历。回调 Err 立即停止并原样返回。回调在打开目录前运行，因此可跳过该 I/O。重复名称按提供者稳定顺序保留。会访问子符号链接但不跟随；根使用提供者的 open/stat 行为，可能跟随链接。遍历不承诺隔离约束、循环检测或原子快照。

根通过 `stat_from` 检查，包括独立的打开/关闭。失败时回调收到无条目和原始错误；任何成功的控制响应都会成功结束遍历。目录读取/打开或关闭失败则触发对该目录的第二次回调，携带条目和错误。Continue 遍历其已确认并排序的前缀；SkipDir 丢弃前缀，SkipAll 停止。提供者错误和普通关闭优先级与 `read_dir_from` 一致。所有成功打开的句柄在任何回调前均关闭一次，因此回调不持有句柄。提供者和回调 panic 遵循正常 panic/defer 传播，不转换为 Result。

`WalkLimits` 有四个显式公共 isize 字段：

- `max_depth`：根深度为零。达到最大深度的目录可以被访问并跳过，但 Continue 将超过展开限制并失败。
- `max_directory_entries`：每个目录的结果限制，与 read_dir_from 一样使用限制加一的探测，还受剩余工作预算限制。
- `max_directories`：目录展开尝试次数，包括打开失败；初始根 stat 和展开前跳过的目录不计入。
- `max_work`：初始根 stat 操作计一个单位，每次目录打开尝试计一个，每次目录分页调用计一个，每个数量有效的返回条目计一个。条目单位包括探测项及因名称无效而被拒绝的页。过大数量在调用条目方法或计入条目单位前被拒绝。分页调用至少需要剩余两个单位，用于调用和可能的条目；请求数量上限为剩余工作预算减一。未用的条目容量不计费，但最终 EOF 探测仍需要这份预留。普通回调和必要清理不消耗单位。

无效路径或负限制在 I/O 或回调前返回 InvalidInput。零工作预算不进行 I/O，向回调报告限制错误。算法限制失败使用 InvalidData，操作为 `"walk"`，路径为当前逻辑路径。错误回调可以通过 SkipAll 显式停止，或用 Err 替换错误；Continue 和 SkipDir 不能消除限制失败或恢复遍历。因此即使提供者存在循环或反复失败，工作和目录预算仍能限制它。这些是算法预算，不限制任意提供者或回调的执行时间、名称长度、总分配或进程内存。

```goml
use std::fs;
use fs::{FileSystem, ReadDirFile};

fn inspect_tree[F: FileSystem](filesystem: F) -> Result[(), fs::Error]
where F::Handle: ReadDirFile {
    let limits = fs::WalkLimits {
        max_depth: 32,
        max_directory_entries: 4096,
        max_directories: 1024,
        max_work: 100000,
    };
    fs::walk_from(filesystem, ".", limits, |path, _, error| {
        if let Some(error) = error {
            Err(error)
        } else {
            println(path);
            Ok(Continue)
        }
    })
}
```

`fs::glob_from(filesystem, pattern, limits)` 返回 `Result[Vec[string], fs::GlobError]`，具有相同的 FileSystem/ReadDirFile 约束。它为每个组件使用 `std::path::slash::Pattern`，并使用显式栈，不递归遍历。斜杠始终分隔组件：转义斜杠和字符类中的斜杠在编译相应组件时被拒绝。所有组件在任何文件系统 I/O 前完成编译。空路径、绝对路径、尾随斜杠和空组件路径无效；字面量点/双点组件（包括转义写法）无效，唯有表示根的精确模式 `"."` 除外。路径不清理。重复星号具有普通单组件 Pattern 含义；`**` 不是递归 glob 语法。

完全字面量的模式（包括转义后的字面量元字符）在解码路径上执行一次 `stat_from`，不枚举父目录。包含通配符的模式以最长起始字面量前缀选择首个待枚举目录。后续组件使用编译后的匹配器。最终匹配只需要名称，不需要类型或元数据。中间匹配的 File 和 Other 条目被跳过；Directory 和 Symlink 条目通过提供者打开，以处理下一组件。因此与 walk_from 不同，glob 可能跟随中间链接。两个辅助函数都不提供隔离约束。

成功的 glob 结果按完整逻辑路径进行全局稳定排序，保留重复项；没有匹配是成功的空向量。提供者错误不会被静默忽略，包括字面量路径上的 NotFound 和 Interrupted。中间目录失败会停止，不打开其已确认子项。对于最后一个组件，仍可在剩余预算内匹配已确认目录条目，不再进行文件系统 I/O，然后返回原始提供者或关闭错误。这种已知错误优先于后续匹配或结果限制错误。算法性的目录/工作限制失败则立即停止，不匹配该批次。之前已确认的匹配在任何失败下都保留。

`GlobError::matches()` 提供已确认、排序后的前缀，`.cause()` 返回 `GlobCause::FileSystem(fs::Error)` 或 `GlobCause::Pattern(component_index, slash::MatchError)`。组件索引从零开始；MatchError 内的位置相对于该组件。提供者错误保留字段。glob 生成的 InvalidInput/InvalidData 错误使用操作 `"glob"`；预检查失败的路径为输入模式，执行期间则为当前逻辑路径。GlobError 的 display 和 debug 格式显示原因。`GlobError::new(matches, cause)` 按给定顺序复制向量结构；负责排序的是辅助函数，不是此构造器。

`GlobLimits` 要求七个显式、非负的公共 isize 字段：

- `max_pattern_chars` 限制整个输入的 Unicode 字符数，包括分隔符和转义；`max_components` 限制组件数。根模式 `"."` 计一个字符、一个组件。
- `max_directory_entries` 和 `max_directories` 限制每个目录收集的条目数和总展开尝试次数。完全字面量的 stat 调用不计为目录展开。
- `max_match_calls` 限制跨所有目录的匹配器调用次数。
- `max_matches` 限制返回匹配数。零仍会搜索首个匹配，以区分无结果和超量结果；字面量模式仍对其路径执行 stat。保留结果按遍历顺序选择，再全局排序，并非选择字典序最小的可能子集。
- `max_work` 由所有目录 I/O 和匹配共享。目录工作采用与 walk_from 相同的打开/分页/条目计费和 EOF 预留；完全字面量 stat 消耗一个单位。匹配候选项前，glob 预先扣除 `P + N * (1 + 2 * P)` 单位，其中 P 是组件原始字符数，N 是候选名称的 UTF-8 字节长度。此保守界限覆盖 Pattern 的 token、字符和字符类范围工作；未用的匹配额度不退回。计算成本前通过除法检查避免溢出。模式编译独立受前两个限制约束，不消耗 max_work。

负限制在 I/O 前返回 InvalidInput；超出限制返回 InvalidData。排序、任意提供者执行和总分配不受时间或内存配额约束。成功打开的句柄恰好关闭一次，包括 panic 清理；panic 和普通关闭优先级不变。

```goml
use std::fs;
use fs::{FileSystem, ReadDirFile};

fn text_files[F: FileSystem](filesystem: F) -> Result[Vec[string], fs::GlobError]
where F::Handle: ReadDirFile {
    let limits = fs::GlobLimits {
        max_pattern_chars: 1024,
        max_components: 32,
        max_directory_entries: 4096,
        max_directories: 1024,
        max_match_calls: 100000,
        max_matches: 4096,
        max_work: 10000000,
    };
    fs::glob_from(filesystem, "docs/*.txt", limits)
}
```

`fs::LinkFileSystem: FileSystem` 提供 `read_link(name) -> Result[string, fs::Error]` 和 `symlink_metadata(name) -> Result[fs::Metadata, fs::Error]`。后者检查最终条目但不跟随该链接；前面组件的解析由提供者定义。提供者验证直接调用的逻辑名称。`read_link_from` 和 `symlink_metadata_from` 也在分派前验证，再原样返回提供者结果。辅助函数静态要求 LinkFileSystem：没有动态能力测试，也没有用打开/跟随链接代替检查的回退。

链接目标是不透明的提供者字符串，不作为逻辑输入路径验证。它们可以是相对或绝对路径、包含双点或采用提供者专属语法。当 F 实现 LinkFileSystem 时，`SubFs[F]` 也实现它，像 open 一样给请求名称添加前缀，但保持目标和错误不变。目标不清理，也不重新基于子文件系统定位。这是命名空间组合，不是符号链接安全沙箱。这些协议不增加原生后端。链式关联类型、约束和辅助函数均使用现有语法，无需语法改动。

<a id="testing-assertions"></a>

## 测试断言

测试属性、文件布局和命令见[测试](tooling.zh-CN.md#tests)。

`std::testing` 提供以下断言：

- `testing::fail(message)`：立即使当前测试失败；
- `testing::assert(condition)`：要求条件为 `true`；
- `testing::assert_eq(actual, expected)`：要求两个类型相同且实现 `PartialEq + Debug` 的值相等；
- `testing::assert_ne(actual, expected)`：要求两个类型相同且实现 `PartialEq + Debug` 的值不相等；
- `testing::assert_some`、`assert_none`、`assert_ok` 和 `assert_err`：检查 `Option` 或 `Result` 的变体；
- `testing::expect_some`、`expect_ok` 和 `expect_err`：返回所选载荷，或以给定消息失败。

<a id="in-memory-test-filesystems"></a>

## 内存测试文件系统

导入 `std::testing::fs as memory` 即可使用纯 GoML、只读的 `MemoryFS`。它实现 `fs::FileSystem` 和 `fs::LinkFileSystem`；其 `MemoryHandle` 实现 `fs::File`、`fs::ReadDirFile`、`io::Read`、`io::ReadAt`、`io::Seek` 和 `io::Close`。此包是可移植测试工具，不是 OS 挂载点，也不是宿主文件系统替代品。

`MemoryFS::new(nodes: Vec[MemoryNode], limits: MemoryLimits)` 返回 `Result[MemoryFS, fs::Error]`。每个节点有公共 `path`、`mode`、`modified_unix_nanoseconds` 和 `content` 字段。`MemoryContent` 为 `File(Bytes)`、`Directory`、`Link(string)` 或 `Special`；前三者要求匹配的 NodeKind，Special 要求其他种类。便捷构造器 `MemoryNode::file`、`directory` 和 `link` 分别使用 0444、0555 和 0777 权限、零修改时间，以及已知全 false 标志。构造复制文件字节和容器；后续修改输入不会改变快照。

名称使用 `fs::valid_path`。重复路径以 AlreadyExists 失败；非目录祖先、种类不一致和非目录根以 InvalidInput 失败。缺失的父目录和根推断为 0555 目录，显式提供的目录元数据保留。`node_count()` 包含推断节点。元数据长度为文件数据或原始链接目标的字节长度，目录和特殊节点为零。目录条目名称唯一且已排序。

MemoryLimits 有六个非负字段。`max_nodes` 包含推断节点；`max_path_bytes` 汇总所有唯一存储路径的字节长度，包括根 `.`；`max_data_bytes` 汇总文件和链接目标字节；`max_depth` 统计路径组件，根为零。超出构造限制返回 InvalidData；负限制返回 InvalidInput。这些是逻辑数据边界，不是分配器或 CPU 时间配额。

每次解析分别使用 `max_link_hops` 和 `max_resolve_work`。工作预算按输入字节、一次根查找、每个处理的组件、每次子项查找，以及每个被跟随链接目标的字节数计费。因此解析 `.` 消耗三个单位。逻辑输入验证先于此计数器。这些单位不统计路径拼接或哈希处理的每个字节，不是 CPU 时间配额。链接目标相对于父目录解析，不预先做词法清理：双点应用于实际到达的目录。目标支持重复分隔符和点组件，但穿过文件的遍历（包括尾随分隔符）会失败。绝对目标和试图离开虚拟根以 InvalidInput 失败；空目标和缺失条目以 NotFound 失败；工作/跳数耗尽以 InvalidData 失败。反斜杠、冒号和 NUL 仍是普通名称字符。`read_link` 返回未修改的最终链接目标，`symlink_metadata` 不跟随最终链接；中间链接会被解析。错误保留原始请求名称。

每次打开创建独立游标；同一句柄的副本共享游标和关闭状态。Stat 不移动游标。读取将字节复制到调用方存储，EOF 时返回零，定位读取从不移动游标。短定位读取报告 UnexpectedEof 及已确认字节数；负范围或溢出范围在复制前失败。Seek 允许超过 EOF 的位置，但对负结果/溢出结果拒绝且不移动游标。目录分页要求正数量，返回新向量和缓存的 SnapshotEntry 值，EOF 时反复返回空页。关闭后条目仍可使用。目录拒绝字节读取和 seek；特殊节点的数据和目录操作返回 Unsupported。Close 幂等。已关闭句柄上的原始 stat/read/read_at/seek/read_dir 操作在其他检查前以 InvalidInput 失败。继承的 I/O 辅助函数保持自身语义，包括空 read_exact 无需调用 read 即成功。不支持对同一共享句柄并发操作。

```goml
use std::bytes;
use std::fs;
use std::testing::fs as memory;

fn sample_tree() -> Result[bytes::Bytes, fs::Error] {
    let limits = memory::MemoryLimits {
        max_nodes: 16, max_data_bytes: 1024, max_path_bytes: 1024,
        max_depth: 8, max_link_hops: 16, max_resolve_work: 4096,
    };
    let source = memory::MemoryFS::new(Vec::from_array([
        memory::MemoryNode::file("config/name", bytes::Bytes::from_vec(Vec::from_array([b'o', b'k']))),
        memory::MemoryNode::link("latest", "config/name"),
    ]), limits)?;
    fs::read_file_from(source, "latest", 16)
}
```

<a id="filesystem-contract-reports"></a>

## 文件系统契约报告

`memory::check_fs(source, expected, limits)` 检查稳定只读提供者，返回 `Result[CheckReport, fs::Error]`。它要求 `FileSystem` 及 `ReadDirFile` 句柄。无效检查器输入在任何提供者调用前返回 InvalidInput。关于提供者的发现通过报告返回，不触发测试断言。调用方必须在检查期间保持文件系统稳定；内容变化报告为快照不一致，不能据此证明活动或可变文件系统违反 trait。

`ExpectedTree::Contains(paths)` 要求存在这些实际枚举到的路径，允许额外条目。路径在调用提供者前被验证和复制，重复项忽略。空 Contains 列表不施加覆盖要求。`ExpectedTree::Empty` 要求根 `.` 下没有条目。遍历不跟随子符号链接，因此经此类链接的预期路径不会隐式展开。只有遍历完成且此前无发现或预算耗尽时，才报告缺失路径。

基本检查器探测无效逻辑名称，验证顺序读取数量和空读取，比较使用不同读取宽度的交错独立打开，并检查 stat 和空读取不扰动位置。在稳定快照假设下比较已确认内容长度和元数据。目录检查比较不同页宽，在枚举前和期间探测非正数量，保留错误前缀，并在后续页和关闭后重新检查条目名称、类型和元数据。原始页可以无序且包含重复名称；比较保留重数。NUL、反斜杠和冒号仍是有效组件字符。子符号链接元数据不与跟随链接的 open/stat 比较。特殊节点被检查，但不假设支持普通文件读取。泛型关闭行为不按 MemoryHandle 特有的幂等性或关闭状态策略测试。

CheckLimits 包含非负的 `max_nodes`、`max_depth`、`max_directory_entries`、`max_file_bytes`、`max_total_read_bytes`、`max_provider_calls`、`max_issues` 和 `max_path_bytes`。发现时计入节点和完整路径字节，包括根和重复条目；深度从零开始。每个显式提供者操作（包括条目访问器）消耗一次调用。所有重读中的已确认成功读取数量消耗总字节预算。文件和目录限制允许额外一次探测，以区分 EOF 与超量；读取探测还需剩余总字节额度。没有足够额度证明 EOF 时，检查以未完成状态停止。清理是强制的，不受这些预算约束。边界不会中断阻塞的提供者、限制其内部分配或施加 CPU 时间配额；目录多重集比较可能进行平方级本地比较。输入验证和复制预期列表也先于提供者调用计费。

报告访问器为 `issues()`、`complete()`、`passed()`、`stop_reason()`、`discovered_nodes()`、`path_bytes()`、`read_bytes()`、`provider_calls()` 和 `cleanup_calls()`。`issues()` 是只读快照。每个问题具有路径、操作和 CheckIssueKind：`ContractViolation(message)`、`ProviderError(fs::Error)`、`InconsistentSnapshot(message)` 或 `ProviderPanic(panic::Panic)`。普通提供者错误（包括 Interrupted）会被记录，不自动重试，也不标记为契约违反。任何发现都令 complete/passed 为 false。全局停止原因区分 `Budget(path, CheckBudget)`、`IssueLimit` 和 `ProviderPanic`；CheckBudget 指出耗尽的是 Nodes、Depth、DirectoryEntries、FileBytes、TotalReadBytes、ProviderCalls 或 PathBytes 限制。保留首个停止原因。

检查器对每次成功打开恰好关闭一次，包括无效路径意外打开成功的情况。即使预算停止或发生可捕获 panic，也会尝试关闭。每次关闭有独立的 panic 边界，因此某个清理 panic 不会阻止关闭其他已打开句柄。在报告仍有空间时，清理错误与此前问题一同保留；报告空间耗尽绝不阻止清理。致命运行时终止及语言可捕获 panic 机制之外的失败不会变得可恢复。

`check_seek`、`check_read_at` 和 `check_links` 接受相同的预期树和限制，返回相同报告。每个函数使用一份共享预算运行基本树检查及额外协议检查。能力通过静态约束要求：check_seek 要求句柄实现 ReadDirFile 和 Seek；check_read_at 要求 ReadDirFile 和 ReadAt；check_links 要求 LinkFileSystem 和 ReadDirFile 句柄。不使用动态接口发现。

Seek 检查将有界顺序快照与 Start、Current 和 End 位置的读取比较，验证返回位置，并探测负数或溢出算术。它不要求能够定位到 EOF 之外，也不要求错误后保留全部流状态。无效 seek 后使用绝对重新定位。不支持的 End 操作属于普通提供者错误，使检查未完成，而非契约违反。

ReadAt 检查将已确认前缀与顺序快照比较，要求成功时完整数量、出错时有界进度，拒绝接受无效范围或为无效范围报告进度的行为，并测试 EOF、空缓冲区和交错顺序位置。在真实快照边界处，部分 UnexpectedEof 属于预期；其他错误仍是提供者错误，即使报告完整传输。出错时的已确认前缀数量消耗共享读取预算。使用独立的新句柄探测成功读取和部分 EOF 读取是否保留游标，包括单字节文件。非空文件上的空定位读取也使用新句柄，因此把游标移到已观察到的 EOF 之外不能隐藏缺陷。对于真正的空文件，仅凭 Read 和 ReadAt 无法区分位于 EOF 或其后的数值位置；不会默默假设具有 Seek 能力。无效范围探测期望 InvalidInput 且零进度；其他错误原因仍作为提供者错误，不会悄然通过。检查器不对所有提供者施加 Cursor 特有的目标尾部不变规则。

链接检查重复调用原始 read-link 和不跟随链接的元数据操作，并与保留的链接条目元数据比较。它还为两个方法探测无效输入名称。它既不清理也不验证目标字符串，不打开或递归遍历子链接。绝对、空、逃逸、循环和悬空目标本身不构成协议违反；跟随解析策略不属于此检查。

`check_helpers` 具有与 check_fs 相同的签名和静态要求。它运行基本检查，将完整节点快照与 `stat_from`、`read_file_from` 和 `read_dir_from` 比较，既直接调用，也通过 `SubFs` 调用。文件以父目录作为 SubFs 根；目录以自身为根并查询 `.`。子符号链接保持不跟随。辅助函数的目录结果必须排序，但重复条目保留重数。失败或未完成的文件读取或目录枚举基线不用于推断辅助函数不匹配。

受观察的提供者将每个底层辅助操作计入同一报告预算，包括条目访问器和所有已确认读取字节。缓存条目值不导致额外底层调用。辅助函数没有独立额度，清理仍豁免。普通错误保留提供者分类；主错误和独立清理错误都被保留，只有关闭错误时只报告一次。预算停止和可捕获 panic 会丢弃辅助函数结果，不把内部停止哨兵报告为提供者错误。这些比较测试稳定提供者的组合，不测试宿主隔离或原子快照。

公共类型和函数使用现有语法，不增加新的语法产生式。

```goml
use std::fs;
use std::testing::fs as memory;

fn inspect_tree(source: memory::MemoryFS) -> Result[memory::CheckReport, fs::Error] {
    memory::check_helpers(source, memory::ExpectedTree::Contains(Vec::new()), memory::CheckLimits {
        max_nodes: 256, max_depth: 16, max_directory_entries: 128,
        max_file_bytes: 65536, max_total_read_bytes: 1048576,
        max_provider_calls: 100000, max_issues: 64, max_path_bytes: 65536,
    })
}
```

Ok 结果表示检查器输入有效，不表示提供者通过；请检查返回报告的 `passed()`、问题和停止原因。

<a id="io-test-fixtures"></a>

## I/O 测试用例工具

导入 `std::testing::io as iotest` 可使用确定性的短 I/O 和延迟错误。`FragmentReader::new(data: bytes::Bytes, max_chunk)` 对输入生成快照；`FragmentWriter::new(max_chunk)` 累积输出。非空读取或写入最多传输正的块大小限制数量。`with_error` 接受非负的 `fail_after` 字节偏移和 `io::Error`：恰好传输这么多字节后，下次非空调用返回同一错误。读取器失败偏移不能超过输入长度。空调用返回 `Ok(0)`，不触发错误。`position()` 和 `calls()` 检查读取器进度；`bytes()` 返回写入器输出快照，`calls()` 返回调用次数。测试工具副本共享游标、输出和调用次数，因此调用方需串行访问。这些工具使用标准 `io::Read` 和 `io::Write` trait，用于测试传输行为。

```goml
use std::bytes;
use std::io;
use std::testing::io as iotest;

fn short_reader() -> Result[iotest::FragmentReader, io::Error] {
    iotest::FragmentReader::new(bytes::Bytes::from_vec(Vec::from_array([b'a', b'b'])), 1)
}
```

<a id="resource-scopes"></a>

## 资源作用域

导入 `std::resource` 可显式清理由库管理的资源。`with_cleanup(action, cleanup)` 执行 action，正常返回后执行 cleanup，包括 action 返回 `Err` 的情况。`scope(action)` 将 `Scope[C]` 传入 action，然后关闭它。`Scope::register` 接受 `() -> Result[(), C]` 清理回调。关闭时按注册顺序的逆序运行全部回调，清理出错后仍继续，每个回调运行一次。

`Scope::new`、`register`、`is_closed` 和 `close` 支持并发调用方。开始关闭后，注册返回 `false`；调用方仍负责被拒绝回调对应的资源。并发 close 调用等待首个 close，取得其错误向量副本；后续调用返回相同错误，不重复清理。回调在作用域状态门闩之外运行，可以检查状态或尝试注册，但不得同步关闭自身作用域。

`ScopeError[E, C]` 保留 `action: Option[E]` 和 `cleanup: Vec[C]`；清理失败不会掩盖 action 错误。`finish(action_result, cleanup_errors)` 按同一规则组合已完成操作。成功 action 仅在清理也成功时产生 `Ok`。两种错误类型都实现 `ToString` 时，`ScopeError` 也实现它。`io::with_resource(value, action)` 为任意 `io::Close` 值使用相同机制。

```goml
use std::resource;

fn run_job() -> Result[isize, resource::ScopeError[string, string]] {
    resource::with_cleanup(|| Ok(42), || Ok(()))
}
```

这些辅助函数不实现析构器或隐式 panic 恢复。清理在正常控制流、`Result` 错误和 panic 展开时运行。某个清理 panic 不会阻止更早注册的清理执行，作用域仍会向等待关闭者发布已关闭状态。最新的 panic 继续传播，不转成 `ScopeError`；普通 action 和清理错误保留现有组合结果行为。进程突然终止和致命运行时失败可能绕过清理。逃逸出回调的作用域在之后仍保持关闭。

<a id="composable-io"></a>

## 可组合 I/O

导入 `std::io::{Read, Write, BufRead, Seek, Close}`，通过泛型约束组合流。`Read::read(MutSlice[byte])` 和 `Write::write(Slice[byte])` 返回已传输字节数；短传输是正常情况，非空读取返回零表示 EOF。超出给定缓冲区的数量无效。`read_exact` 和 `write_all` 完成部分传输，进度停止时报告 `UnexpectedEof` 或 `WriteZero`。`read_exact` 和 `write_all` 立即返回任何提供者错误，包括 `Interrupted`。失败调用不报告进度：重试读取可能丢失已消费输入，重放写入可能重复副作用，包括使用 MultiWriter 或 OffsetWriter 时。这有意改变了此前泛型中断重试行为，但不改签名；明确知道中断未产生进度的具体提供者可以内部重试。错误可能发生在部分进度之后，不回滚数据。

`Seek::seek(from: SeekFrom) -> Result[isize, Error]` 提供独立于 `Read` 或 `Write` 的检查式字节定位。`Start(offset)` 相对于逻辑流起点，`Current(delta)` 相对于当前位置，`End(delta)` 在逻辑长度可用时相对于该长度。成功返回新的非负位置。位置算术以 `InvalidInput` 拒绝负结果或溢出；偏移和位置是机器大小的有符号整数。`seek(SeekFrom::Current(0))` 报告当前位置，但仍应用具体流的 seek 副作用。该 trait 不承诺任意 EOF 之外的定位，也不承诺错误后恢复全部状态。

| Seek 实现 | 有效位置及副作用 |
| --- | --- |
| `StringReader` | 允许直到最大 `isize` 的位置，包括 EOF 之外和 UTF-8 标量内部；每次尝试均清除字符回退状态，即使失败也如此。 |
| `Cursor` | 限制在当前字节长度内，而非配置的写入限制内；从不增长存储或创建空洞。End 使用写入后的当前长度。 |
| `SectionReader[R: ReadAt]` | 限制在声明的区段长度内，所有位置均相对于区段。成功清除待处理读取错误；失败保留它。 |
| `OffsetWriter[W: WriteAt]` | Start/Current 相对于其基点，且限制为基点加位置仍可表示。End 返回 `Unsupported`，因为 `WriteAt` 没有长度契约。 |

这些实现通过现有句柄别名共享位置，seek 时不读写，也不 flush 或关闭提供者。无效定位保持位置不变。OffsetWriter seek 成功不保证提供者接受后续在该偏移的写入。现有 `set_position` 和 StringReader 的固有 `seek` 仍可用。缓冲包装器不实现 Seek：其预读/排空状态需要单独定位策略。新 trait 通过现有语法组合，例如 `fn rewind[S: io::Seek](stream: S) -> Result[isize, io::Error] { stream.seek(io::SeekFrom::Start(0)) }`；不新增语法或运行时原语。

`Read::read_to_end(limit)` 返回 `bytes::Bytes`；`read_to_string(limit)` 额外验证 UTF-8。限制是非负字节数。两者立即返回任何提供者错误，包括 `Interrupted`，不重试，也不返回部分输出。两者都拒绝过大输入，最多额外消费一个字节以探测溢出；若限制之外的字节必须保持未读，使用 `Take::new(reader, limit)`。`io::copy(reader, writer)` 复制直到 EOF，返回 `u64` 数量；不 flush 或关闭任一流。

`io::read_at_least(reader, buffer: MutSlice[byte], minimum: isize)` 返回 `Result[isize, TransferError]`，失败时保留已确认进度。负 minimum 或超过视图长度的 minimum 在 I/O 前以 `InvalidInput` 拒绝；零无需调用读取器即可成功。每次 read 接收整个剩余视图，因此成功数量可能超过 minimum，但绝不超过视图长度。将 `buffer.len()` 作为 minimum，即可执行带进度报告的精确读取。达到 minimum 前遇到 EOF 返回 `UnexpectedEof`，即使已确认字节为零；不可能的提供者数量返回 `InvalidData`。此函数不初始化未触及的视图后缀和视图外存储。

任何提供者错误（包括 `Interrupted`）都立即返回，携带此前有效成功数量之和。失败调用可能修改给定缓冲区或消费输入，却不报告多少；这些影响不计入也不回滚，因此该数量不是普遍安全的恢复偏移。此函数和 `Read::read_exact` 都立即返回中断。它不分配暂存缓冲区、不 flush/关闭任何对象，使用普通泛型/函数语法。例如，单字节视图和 minimum 为一时，从 `io::StringReader::new("abc")` 读取一个字节，不消费其余内容。

`io::NoopCloser::new(reader)` 为任意 `Read` 值包装一个始终成功且不触碰底层流的 `Close` 实现。读取原样转发调用、数量和错误，包括空缓冲区；包装器不在此边界验证提供者数量。重复 close 不会禁用后续读取，也不会关闭同时实现 `Close` 的提供者。当作用域借用读取职责但不能接管关闭职责时，可将其与 `io::with_resource` 配合使用。提供者别名和并发规则不变。包装器提供 `Read` 和 `Close`，不自动转发定位、缓冲或优化传输接口。

`io::copy_buffer(reader, writer, buffer: MutSlice[byte]) -> Result[u64, io::Error]` 使用给定非空暂存视图，不自行分配缓冲区。空缓冲区在访问任一流前返回 `InvalidInput`。读取验证数量并返回任何错误，包括 `Interrupted`；写入使用目标的 `write_all` 实现，包括适配器专属的不重试策略。暂存视图会被覆盖，不得与活动源/目标存储形成别名，也不得并发使用。没有优化传输协议会绕过此缓冲区。`copy` 现在使用相同引擎和 8 KiB 缓冲区，保留原有行为。

`io::copy_n(reader, writer, count: isize) -> Result[u64, io::Error]` 通过有界读取器精确复制请求字节数。负 count 在访问任一流前失败；零无需读写或分配复制缓冲区即成功。它绝不读取超过请求数量。提前 EOF 返回 `UnexpectedEof`；已复制字节保留在目标中。这些传统 copy 变体遇到其他错误会停止，不回滚副作用，失败时不返回进度数量。输入可能已被消费后才发生目标失败。它们都不 flush 或关闭流。

`io::copy_progress(reader, writer)` 和 `copy_buffer_progress(reader, writer, scratch: MutSlice[byte])` 返回 `Result[u64, CopyError]`。成功报告已复制字节数。失败暴露 `CopyError::read_count()`、`written_count()` 和 `error()`：分别给出已确认输入/输出数量及原始提供者错误。预读意味着输入数量可能超过输出数量。失败调用可能消费输入、修改 scratch 或产生输出却不报告进度；这些影响不计入也不回滚。两个数量都不是普遍安全的恢复偏移。

这些函数直接调用 `read` 和 `write`，完成成功的短传输，从不重试错误，包括 `Interrupted`。它们不调用自定义 `write_all` 重写，不 flush 或关闭任一流。无效的成功数量产生 `InvalidData`；零写入产生 `WriteZero`。空 scratch 在 I/O 前以 `InvalidInput` 和零数量拒绝。Scratch 不得与活动源/目标存储形成别名，也不得并发访问。`copy_progress` 分配 8 KiB 暂存缓冲区。读取裁剪到剩余 `u64` 计数范围；达到最大数量后，下次迭代在 I/O 前报告 `InvalidData`，即使再读一次本可以发现 EOF。

`io::Discard::new()` 创建无状态 `Write` 实现，接受整个给定缓冲区（包括空缓冲区）而不保留数据；`flush` 无需操作即成功。例如，`io::copy_n(reader, io::Discard::new(), 32)` 精确跳过 32 字节，或报告提前 EOF。

`io::ReadAt::read_at(output: MutSlice[byte], offset: isize)` 和 `io::WriteAt::write_at(input: Slice[byte], offset: isize)` 执行定位传输，不改变顺序位置。结果为 `Result[isize, io::TransferError]`：成功报告完整缓冲区长度；失败可同时报告已传输前缀和原因。`TransferError::new(transferred: usize, error: io::Error)`、`.transferred()` 和 `.error()` 保留两者，不改变现有 `Read`/`Write` 错误签名。实现必须拒绝负偏移和偏移加长度溢出，并将短传输报告为错误，不得报告超过给定缓冲区的进度。

`read_exact_at` 和 `write_all_at` 在调用实现一次前验证区间。它们验证返回数量和错误进度，保留合法底层错误（包括 `Interrupted`），不重试。短的成功读取规范化为 `UnexpectedEof`；短的成功写入或不可能的数量变为 `InvalidData`。不可能的进度报告为零，因为提供者的计数不可信；这不意味着没有发生修改。错误从不回滚已传输字节或其他提供者副作用。实现返回后不得保留给定缓冲区。不同于 Go ReaderAt/WriterAt 的并发契约，这些 trait 本身不保证并发安全：调用方必须遵循具体提供者的同步要求。偏移类型是 GoML 有符号机器字，不是独立的 64 位文件偏移类型。

`Cursor` 也实现 `ReadAt` 和 `WriteAt`。越过末尾的非空读取复制可用前缀，返回带该数量的 `UnexpectedEof`，不改变目标余下部分。任何非负偏移的空读取都成功。定位写入可以覆盖或在当前长度处追加，但不能创建空洞；验证失败保持存储不变。顺序和定位写入都在修改前为输入建立快照，包括从 `fill_buf` 获得的重叠视图。定位操作保持 `position()` 不变，错误时也如此。例如，导入 `use std::io::ReadAt` 后，`cursor.read_exact_at(output.as_mut_slice(), 4)` 从第四字节起填充缓冲区，不移动 cursor。

`BufRead::fill_buf` 暴露当前缓冲字节，`consume(amount)` 检查可用范围。`read_until(delimiter, limit)` 包含找到的分隔符；`read_line(limit)` 包含换行、验证 UTF-8，仅在 EOF 且无字节时返回 `None`。最后未终止行可以恰好填满限制。缓冲视图只应视为在该 reader 下一次操作之前有效。

`BufRead::discard(amount: isize) -> Result[isize, TransferError]` 跨缓冲填充恰好跳过请求的字节数。负数量以 `InvalidInput` 失败；零不调用 `fill_buf` 或 `consume` 即成功。先消费已有缓冲字节；输入不足返回 `UnexpectedEof`。错误中的已传输数量仅包括成功的 `consume` 调用。每个 fill 或 consume 错误都立即返回，不重试，保留原始错误；失败调用未报告的副作用不计入。方法不分配临时存储，不刷新或关闭任何内容。具体缓冲提供者可能预读，因此丢弃字节数不一定等于源位置前进量。导入 `io::BufRead` 以使用此默认方法。

`Cursor::new(bytes)` 复制初始数据，句柄副本共享位置。它实现 `Read`、`Write` 和 `BufRead`；写入覆盖或追加，`set_position` 接受不超过当前长度的偏移，`bytes()` 返回副本。`Cursor::with_limit(bytes, max_bytes)` 在复制前检查初始长度，并按缓冲区总长度而非累计写入字节限制所有后续写入。负限制返回 `InvalidInput`；过多初始数据或写入返回 `InvalidData`。拒绝的写入保留数据和位置，报告零定位进度，即使输入与 `fill_buf` 形成别名。空写入和限制内覆盖仍有效。`limit()` 返回配置限制；`new` 使用最大 `isize`。限制不包括快照、输入副本和分配开销，也不恢复分配失败。

`BufReader::new(reader)` 和 `BufWriter::new(writer)` 使用 8 KiB 缓冲区；`with_capacity(stream, size)` 拒绝非正大小。缓冲 reader 立即返回每个提供者错误，包括 `Interrupted`，不重试失败读取。缓冲 writer 需要显式 `flush`；失败时保留未由成功写入数量确认的后缀，它不一定尚未发送。排空不重试错误，包括 `Interrupted`。重试失败 flush 可能重复失败调用的未知副作用；恢复需要了解具体 writer。底层流实现 `Close` 时，缓冲包装器也实现；writer close 先刷新。这些内存适配器要求调用方串行化共享访问，没有终结器。

`BufReader::peek(amount: isize) -> Result[Slice[byte], PeekError]` 返回恰好请求数量的缓冲字节，不消费它们。小于零或大于 `capacity()` 的请求在 I/O 或缓冲压缩前以 `InvalidInput` 失败；零返回空视图，不做 I/O。合法请求可以压缩未读字节并执行多次短读，使用剩余容量，可能预读。从不增长缓冲区。在请求大小前 EOF 返回 `UnexpectedEof`；提供者错误（包括 `Interrupted`）原样保留，不重试，非法数量产生 `InvalidData`。

`PeekError::available()` 暴露当前缓冲中已确认前缀，`error()` 暴露原因。失败调用在该前缀之外的修改不暴露，但失败调用消耗的输入无法恢复。EOF 和错误都不永久缓存：若需要更多字节，后续显式 peek 可再次读取。成功和错误视图借用共享可变缓冲存储；在其他操作或别名修改 reader 前应复制它们。peek 不是快照、消费操作，也不保证重试失败是安全的。

`BufReader::read_byte()` 返回 `Result[Option[byte], Error]`，`read_char()` 返回 `Result[Option[(char, isize)], Error]`，其中元组包含 Unicode 标量及其原始字节宽度。`None` 表示 EOF 且无剩余字节。畸形或被 EOF 截断的 UTF-8 序列消耗一字节，返回宽度一的 U+FFFD；合法 U+FFFD 消耗其三字节编码。字符读取即使配置容量为一也可工作。物理缓冲至少保留四字节用于组装标量，但 `capacity()` 保持配置值，每次提供者读取受该值限制。因此字符操作后，`buffered_len()` 和 `fill_buf()` 可能暴露最多 `max(capacity(), 4)` 字节。`peek` 仍拒绝超过 `capacity()` 的请求。此固定标量存储不会随输入增长，也不改变 writer 容量。

字符读取立即返回提供者错误，包括 `Interrupted`，不消费当前字符此前已确认的任何前缀。该前缀保留在缓冲中，供后续字节、批量、peek 或 discard 操作使用。失败提供者调用修改或消耗的字节未知，不计入缓冲前缀；重试无法重建丢失输入。EOF 和错误不会持续锁定。这些方法不刷新或关闭源。

`unread_byte()` 和 `unread_char()` 返回 `Result[(), Error]`，允许在符合条件的成功读取后回退一次。字节回退可在 `read_byte`、`read_char` 或非空成功批量 `Read::read` 后使用，仅恢复最后一个字节。字符回退仅可在 `read_char` 后使用，恢复原始编码字节，包括单个畸形字节，不重新编码 U+FFFD。回退仅改变共享逻辑缓冲偏移，不定位提供者，不需 I/O。没有新符合条件读取就重复回退，或回退类别错误，均返回 `InvalidInput`。

每次回退尝试都会清除两种回退许可，包括失败尝试。`peek`、`fill_buf`、`consume` 和 `reset` 尝试也清除它们，即使参数为零或非法；空批量读取、EOF 和读取错误同样清除。`capacity()` 和 `buffered_len()` 仅观察，保留许可。默认辅助方法遵循实际执行的操作：`discard(0)` 和对空视图的 `read_exact` 不调用 reader，保持许可不变。句柄别名共享许可和偏移。复制的字节快照可在后续操作后保留，借用缓冲视图则不能。例如，`reader.read_char()` 返回 `Some(('界', 3))` 后，`reader.unread_char()` 使相同的三个原始字节再次可用。

`BufReader::read_fragment(delimiter: byte)` 返回 `Result[ReadFragment, FragmentError]`，消费逻辑输入的一个有界片段。`ReadFragment::bytes()` 是借用视图，`consumed()` 报告本次调用消耗的已确认原始字节，`end()` 标识边界：

| `FragmentEnd` | 含义 |
| --- | --- |
| `Delimiter` | 找到请求的分隔符，并将其包含在原始片段中。容量边界处的分隔符优先于 BufferFull。 |
| `BufferFull` | 未找到分隔符便达到配置容量。不额外读取探测 EOF；请继续调用。 |
| `Eof` | 缓冲填满前，提供者读取返回零。片段可以非空或为空。 |

每个原始片段最多消耗 `capacity()` 字节，即使此前字符操作留下更多缓冲字节。未消费缓存字节仍可用。读取仅在配置容量内追加，新字节只扫描一次；不分配整个输入，也不增长片段缓冲区。EOF 不永久缓存，因此后续显式调用可观察恢复的源。

`FragmentError::bytes()` 暴露已消费的已确认部分片段，区别于 `PeekError::available()`。`consumed()` 报告其长度，`error()` 保留提供者原因；非法成功计数变为 `InvalidData`。所有错误，包括 `Interrupted`，均立即停止。失败调用的效果不计入，重试无法恢复该调用丢失的输入。成功和失败视图都共享 reader 存储，只在下一次 reader 操作前有效；需要保留时应复制。别名共享消费状态。片段读取总清除字节/字符回退许可，包括失败时，返回后不启用回退。

`read_line_fragment()` 具有相同结果类型，使用 LF 为分隔符，并从 `Delimiter` 载荷中移除 LF 和紧邻其前的 CR。`consumed()` 仍计入被移除的传输字节。不验证 UTF-8。`BufferFull` 表示行仍在继续，不是完整行。满边界处尾随 CR 保留到下一片段，以识别跨片段 CRLF，同时不丢弃单独 CR。容量为一时，开头 CR 使用固定标量存储进行一字节前瞻：CRLF 产生空 `Delimiter` 载荷、消费两个字节；CR 后跟其他字节则产生一字节 `BufferFull` 载荷，并保留下一字节。EOF 处的 CR 作为数据返回；前瞻错误返回已确认 CR 及其原因。每次底层读取仍受容量限制。`BufferFull` 结果从不为空。因此恰好容量长度的行后可能跟随空 `Delimiter` 或 `Eof` 结果。未终止尾部保留单独 CR 字节。现有完整行 `read_line(limit)` 契约不变。

两个包装器都暴露 `capacity()` 和 `reset(replacement)`；替代流具有相同具体类型。Capacity 是配置的逻辑缓冲大小，不是分配器容量。Reset 复用存储、清除缓冲状态，并将所有已有别名切换到替代流。它不对任一流执行 read、write、flush 或 close。Reader reset 丢弃未读预取字节；writer reset 有意丢弃待发送输出，包括失败排空后的未确认后缀。若必须保留这些输出，需先显式 flush，同时注意上述失败限制。reset 后现有 reader 视图不可继续使用；需保留时先复制。

`BufWriter::available()` 报告 `capacity() - buffered_len()`，即逻辑剩余空间。部分排空失败后，这不保证下一次写入无需底层 I/O：包装器会先排空保留的后缀。Reset 同时清除排空偏移和数据。后续 close 针对替代流，不是原流。这些操作要求串行访问，不恢复分配失败，不增加语法或运行时原语。

`BufWriter::write_byte(value)` 接受一个字节后返回 `Result[(), Error]`。`write_char(value)` 和 `write_string(value)` 返回 `Result[isize, TransferError]`，补齐成功短写，并报告缓冲 writer 已接受的新输入字节数量。字符串写入使用至多 4096 字节的临时存储，不复制整个输入；字符写入编码至多四字节。空字符串不做 I/O，也不排空已有输出。错误立即停止，不重试；已传输计数不包括失败调用和本次调用前已缓冲的输出。字符或字符串被错误中断时可能产生部分 UTF-8 输出；这些操作不是原子的，从不刷新或关闭提供者。

`BufWriter::read_from[R: Read](reader: R) -> Result[u64, CopyError]` 组合 `copy_progress(reader, self)`，保留分开的已确认输入计数和缓冲接受计数。EOF 不刷新 writer。接受不代表底层目标已确认交付：操作可能排空旧待处理输出、保留新输出，或在未知目标副作用之后失败。现有部分排空、显式刷新、别名和 reset 契约仍适用。此方法不增加优化传输旁路、自动关闭、重试或恢复保证。

`io::pipe() -> (PipeReader, PipeWriter)` 使用现有通道创建同步内存字节流。`PipeReader` 实现 `Read` 和 `Close`；`PipeWriter` 实现 `Write` 和 `Close`。与可变内存适配器不同，复制的管道句柄支持并发调用。writer 每次发布一个数据包，不交错；reader 可分多次短读消费。非空写入仅在全部字节复制到 reader 缓冲区后成功。写入返回前输入必须不变。不保证 FIFO 或公平性，没有后台任务或额外运行时原语。并发操作不得使用重叠可变缓冲区；同步管道状态不意味着同步操作之外的调用方存储。

`PipeWriter::write_progress(input)` 返回 `Result[isize, TransferError]`，关闭中断写入时保留已确认前缀。普通 `write` 返回相同错误但不带前缀。每个端点提供 `close_error(error)`；重复关闭保留该端点的首个原因。关闭从不等待对端进度，并唤醒阻塞操作。本地已关闭的 reader 或 writer 返回 `BrokenPipe`，优先于另一端点的原因。否则，reader 在 writer 正常关闭后看到 EOF，或看到 writer 自定义错误；writer 看到 reader 自定义错误，或在 reader 正常关闭后看到 `BrokenPipe`。关闭 writer 本身以 `BrokenPipe` 中止其待处理写入，即使它向 reader 提供自定义原因。已完成的数据包回复不能被后续关闭替换。

开放状态下直接空读取和写入立即返回零，不发布数据包，也不向等待的 reader 发送 EOF。关闭状态检查优先于这些空操作。继承的 `flush` 无操作，泛型 `write_all(empty)` 不调用 `write`，即使已关闭。管道操作不接受取消令牌参数：取消清理必须关闭端点以唤醒阻塞任务。非空输入的读写必须并发运行：

```gom
use std::io;
use io::{Read, Write, Close};
use std::task;

task::scope(|scope| {
    let (reader, writer) = io::pipe();
    defer { let _ = reader.close(); let _ = writer.close(); };
    let producer = scope.spawn(|_| {
        defer { let _ = writer.close(); };
        writer.write_all(b"hello".as_slice())
    });
    let received = reader.read_to_end(16);
    let sent = producer.join();
});
```

`Scanner[R: Read]::new(reader, max_buffer)` 创建有界行扫描器，拒绝非正限制。`.next() -> Result[Option[bytes::Bytes], Error]` 返回独立拥有的 token、成功终止后的 None，或由 `.failure()` 保留的持续错误。终止后调用不再读取；句柄副本共享状态，要求串行访问。扫描器从不关闭源。默认 `scan_lines` 去掉 LF 和紧邻其前的 CR，输出空行，并对非空最终未终止行去掉末尾 CR 后输出。最后换行不额外创建 token。保留任意字节内容；调用方将 `std::utf8::BytesUtf8` 引入作用域后，通过 `token.to_string_utf8()` 显式验证。`scan_bytes` 输出单个字节，包括非法 UTF-8。

`Scanner::with_split(reader, max_buffer, split)` 接受 `(Slice[byte], bool) -> Result[ScanStep, Error]` 回调。布尔值表示已确认 EOF。`ScanStep::More(advance)` 丢弃指定前缀并重试，advance 为零时请求更多输入。`Token(advance, start, end)` 返回指定字节区间并消费 advance 字节：`0 <= start <= end <= advance <= input.len()`，且 advance 必须为正。`Final(Some((start, end)))` 输出一个可能为空的最终 token 并停止；Final(None) 不输出 token 直接停止。Final 可以在 EOF 前停止，并丢弃缓冲余下内容。非法回调区间和不前进的 token 返回 InvalidData，避免循环。EOF 时 More(0) 对非空输入返回 UnexpectedEof，对空输入正常结束。回调错误保持终止状态，不重试。

限制约束待处理输入字节，包括分隔符，不只是返回 token 大小或进程总内存。读取使用至多 4096 字节临时存储。待处理输入恰好达到限制、回调仍需数据时，可能额外消耗一个字节以区分 EOF 和溢出：EOF 允许交付最终 token，多出字节则产生终止性的 InvalidData。扫描可能预读，因此调用方不得假设源位置就是最后返回 token 的结尾。所有源错误，包括 Interrupted，都原样保留、不重试；失败时丢弃缓冲中的不完整数据，此前 token 仍有效。回调和返回 token 区间使用字节偏移。这些 API 不增加新语法或原生后端。

`scan_runes` 跨读取等待完整 UTF-8 标量，并将每个标量作为合法 UTF-8 token 返回。每个畸形字节，包括 EOF 截断序列的每个字节，都产生 U+FFFD；字面 U+FFFD 与替换 token 无法区分。`scan_words` 按固定版本的 Unicode White_Space 属性拆分，跳过开头/结尾空白，从不输出空词。它保留每个词的原始字节，包括畸形 UTF-8，在决定不完整标量是否分隔词之前会等待补全。这是按空白分词，不是语言学的词边界分割。

`ScanStep::Emit(advance, token)` 支持替换 rune 等转换后的 token。Advance 必须为正且不超过待处理输入长度。Scanner 复制给定 token，因此回调后续修改不改变已返回 token。待处理输入限制不约束回调创建的输出；替换 rune 可能消耗一个输入字节却占用三个字节。自定义回调自行负责输出/资源限制。

对任意 `R: ReadAt`，`SectionReader::new(reader, start, length)` 返回 `Result[SectionReader[R], io::Error]`。它不访问 reader 就拒绝负边界和结束偏移溢出。它暴露实时逻辑子区间，不是快照或安全沙箱：声明长度目前不必已存在于底层输入。`size()` 返回声明长度；`position()` 和带检查的 `set_position()` 使用从零到长度的区段相对偏移。句柄副本共享位置和待处理错误；独立构建的区段不共享。任何操作都不移动底层 reader 的顺序位置，不关闭它，也不读取声明区段之外。支持嵌套区段。

区段使用相对偏移实现 `ReadAt`，将提供者目标裁剪到区段边界，并验证提供者数量。超出区段的请求报告带可用前缀计数的 `UnexpectedEof`。非负偏移的空定位读取成功，不调用提供者。顺序 `Read` 则在逻辑区段末尾返回零。由于 `Read` 无法同时返回数量和错误，带有效正进度的底层失败会推进区段位置并返回该数量；下一次非空顺序读取不调用提供者，返回保存错误一次。空读取保留此待处理错误。成功 `set_position` 清除它；非法定位和定位读取保持其不变。错误交付后的后续读取可在推进后的位置重试提供者。所有共享区段访问需要调用方同步。例如，`io::SectionReader::new(cursor, 16, 32)` 使用现有泛型/类型/方法语法，最多暴露第 16 到 47 字节。

对 `W: WriteAt`，`OffsetWriter::new(writer, base)` 返回 `Result[OffsetWriter[W], io::Error]`，拒绝负 base。它将相对偏移转换为 `base + offset` 来实现 `WriteAt`，验证两次加法和提供者数量。定位写入保持顺序 `position()` 不变。带检查的 `set_position(relative)` 接受绝对偏移可表示的非负位置；不检查底层目标长度，也不创建空洞。提供者决定该位置的后续写入是否受支持。句柄副本共享位置；独立构建和嵌套的 offset writer 各有独立相对位置。

其顺序 `Write::write` 在成功和失败时都按有效报告进度推进位置，然后立即返回底层错误。需要直接返回错误进度计数时，请使用定位 `write_at`。此适配器的 `write_all` 执行一次经过验证的定位传输，不重试 `Interrupted`：失败传输可能已改变目标，不得重放。伴随全长度传输的错误仍是错误。非法提供者进度不可信，不推进位置；实际提供者副作用无法撤销。空写入转发给提供者，可能失败。`flush` 无操作，因为 `WriteAt` 没有刷新契约；需要时显式刷新具体目标。适配器不关闭提供者，也不改变其顺序位置，调用方必须同步共享访问。例如，`io::OffsetWriter::new(cursor, 16)` 使用现有泛型/类型/方法语法从第 16 字节开始写入。

对 `R: Read`，`MultiReader::new(readers: Slice[R])` 按顺序连接 reader。它对句柄列表建立快照，不对数据或位置建立快照。空读取返回零，不访问任何 reader；非空的零读取使当前 reader 永久耗尽，并释放保存的句柄。它跳过耗尽 reader，直到某个返回正进度、错误，或列表结束。数量会验证；错误不推进到下一 reader，因此后续调用重试同一 reader。副本共享遍历状态，耗尽后永久保持。空列表初始即耗尽。适配器不关闭 reader，调用方必须同步共享使用。

对 `W: Write`，`MultiWriter::new(writers: Slice[W])` 为目标句柄列表建立快照，并按顺序向每个目标写入完整输入，包括空输入。遇到首个错误或短写即停止；非空输入零进度为 `WriteZero`，其他短写或非法数量为 `InvalidData`。此前目标可能已包含完整输入，失败目标可能包含前缀；不回滚，也无聚合进度计数。`write_all` 调用此广播一次，从不重试 `Interrupted`，因为重放会重复此前写入。`flush` 按顺序访问目标，在首个失败处停止。空目标列表接受整个输入。重复句柄有意被重复访问，适配器不关闭目标。

两个构造函数的每个列表都使用一种具体元素类型，可嵌套；异构源可使用应用定义的通用 `Read`/`Write` 实现。嵌套不动态展平。例如，导入 `use std::io::Write` 后，`io::MultiWriter::new(Vec::from_array([left, right]).as_slice()).write_all(data.as_slice())` 向两个相同类型的 writer 广播。不引入新语法或运行时分派机制。

`TeeReader::new(reader, writer)` 对 `R: Read, W: Write` 实现 `Read`。每次正数源读取在成功返回前同步传给镜像 writer 的 `write_all`。短写和中断由该 writer 的实现处理，包括其不重试覆盖。tee 不分配暂存缓冲区、不保留目标视图、不刷新或关闭任一流。空读取和源 EOF 不执行镜像操作；源错误和非法数量直接返回，不调用镜像，也不使 tee 进入终止状态。

镜像失败对该 tee 和所有副本都是终止性的。原始错误可通过 `failure() -> Option[io::Error]` 获取；`failed_read_count() -> isize` 记录失败调用放入调用方缓冲区的源字节数，不是到达镜像的字节数。不包含此前完成读取，失败前为零。失败调用立即返回错误，但源消费、缓冲修改和部分镜像写入保留。所有后续读取，包括空读取，都不触碰任一流，直接返回失败。没有 reset 或重放操作。这有意区别于 Go 返回数量加错误的 TeeReader API：GoML 的 `Read` 无法同时返回两者，因此当部分输入重要时，调用方可以检查显式失败状态。

若镜像最终错误为 `Interrupted`，读取报告带终止镜像消息的 `Other`，`failure()` 则保留完全原始错误。为兼容保留此历史终止错误映射；泛型 `read_exact` 和 `copy` 现在不重试中断，直接返回。其他镜像错误保留原始类别和详情。Tee 句柄需要调用方同步，源、镜像和调用方缓冲区不得形成会修改未读源数据的别名。例如，`io::TeeReader::new(source, io::Discard::new())` 使用现有 trait 和泛型语法转发成功读取，不保留副本。

`stdin()` 实现 `Read`，`stdout()` 和 `stderr()` 实现 `Write`，`TcpStream`、`net::tls::TlsStream` 和 Linux `fd::Fd` 实现 `Read`、`Write` 和 `Close`。现有 socket 固有方法仍可用。需要取消时，直接调用 socket 的感知上下文方法；泛型 trait 不携带上下文参数。

```goml
use std::bytes;
use std::io;
use std::io::{BufRead};

fn first_line(data: bytes::Bytes) -> Result[Option[string], io::Error] {
    let reader = io::BufReader::new(io::Cursor::new(data));
    reader.read_line(4096)
}

fn bounded_memory_output() -> Result[io::Cursor, io::Error] {
    io::Cursor::with_limit(bytes::Bytes::new(), 65536)
}
```

<a id="read-only-string-streams"></a>

## 只读字符串流

`io::StringReader::new(value)` 持有不可变字符串，不将其复制到可写字节缓冲区。`size()` 为原始字节长度，`position()` 为当前字节偏移，`len()`/`is_empty()` 描述未读字节。句柄副本共享位置和 `reset(value)`；reset 替换字符串，回到字节零，并清除字符回退状态。共享访问需要串行化。

reader 实现 `Read` 和 `ReadAt`。顺序读取在 EOF 返回零；空顺序读取成功。定位读取同时保留位置和字符回退状态，遵循带检查的 `ReadAt` 区间契约，短读时返回带可用前缀计数的 `TransferError`。任何非负偏移的空定位读取都成功。两个接口都不授予写访问。

`seek(SeekFrom::Start(offset))`、`Current(delta)` 和 `End(delta)` 返回新的绝对字节位置。允许 EOF 之后的非负位置。负数或溢出结果返回 `InvalidInput`，不移动位置；每次 seek 尝试都使字符回退失效。固有方法和 `io::Seek` 实现共享此行为；现有调用方无需导入 trait 即可继续使用固有方法。

`read_byte()` 返回 `Option[byte]`；`read_char()` 返回 `Option[(char, isize)]`，包含标量和消耗的字节宽度。EOF 为 `None`。在 UTF-8 标量内部开始（字节读取或定位后）会产生替换字符并消耗一字节，与 Go strings.Reader 一样。`unread_byte()` 在任意正位置后退一字节，包括 EOF 之外，零位置失败。`unread_char()` 回退最后一次成功字符读取一次。字节读取、字节回退、顺序读取（包括空/EOF）、seek 尝试、reset 和 EOF 字符读取都会使字符回退失效。`write_to` 也使其失效，包括空传输。元数据查询和定位读取不会。空/EOF 顺序读取的失效规则是有意设计，不承诺 Go 状态机的所有特殊行为。这些操作使用现有字符串原语，不增加语法。

例如，`io::StringReader::new("é界").read_char()` 产生 `Some(('é', 2))`；定位到字节 1 后读取字符，产生 `Some(('�', 1))`。

`reader.write_to(writer)` 返回 `Result[isize, io::TransferError]`，使用 4096 字节临时缓冲区传输未读后缀。与泛型先读后写复制不同，reader 仅按成功且有效的 writer 数量推进。它补齐短写，立即返回所有 writer 错误（包括 `Interrupted`），零进度以 `WriteZero` 拒绝，负数/过大数量以 `InvalidData` 拒绝。失败报告本次调用已确认数量，保留未确认后缀供后续调用；位置可能位于 UTF-8 标量内部。空/EOF 传输不调用 writer，也不改变位置。不隐式刷新或关闭。writer 回调不得并发或重入修改 reader 及其别名。失败或返回非法数量的 writer 的未知副作用无法计数或回滚。即使是合法的 MultiWriter 或 OffsetWriter 适配器，失败调用也可能产生此顺序 Write 结果未表示的副作用。因此恢复需要针对提供者了解失败调用，不能仅凭零已确认计数或 Interrupted 错误。不承诺 Go Reader.WriteTo 的单次调用短写行为：此方法有意遵循 GoML 流约定补齐短写。

<a id="compiled-multi-rule-text-replacement"></a>

## 编译式多规则文本替换

`text::Replacer::new(rules: Slice[(string, string)], max_rules, max_rule_bytes)` 返回已编译替换器或 `ReplaceError`。每对都是字面 old/new 规则，不是正则表达式。规则会建立快照；后续修改输入集合不影响替换器。限制在构建前缀树前约束规则数量和 old/new UTF-8 字节长度总和；空规则仍计数。编译状态私有，替换调用不修改它。

`replacer.replace(input, max_bytes, max_work)` 从左到右扫描。每个位置上，最早给定的匹配规则胜出，不是最长规则。重复 old 模式保留首个规则。替换文本输出一次，从不重新扫描。空规则可在一个位置匹配一次，然后在同一位置尝试非空规则；否则扫描前进一个标量。因此空模式保留 UTF-8 边界，不同于 Go strings.Replacer 的逐字节空模式插入。例如，规则 `("", "-")` 将 `界🙂` 映射为 `-界-🙂-`，不会在任一标量编码内部插入字节。

输出限制统计 UTF-8 字节。匹配对每次根查找和尝试的前缀树字节转移计一个单位，包括失败转移；到达可能最早的规则即停止该次查找。耗尽返回 `WorkLimit`，不是错误的部分匹配。输出复制由 max_bytes 单独限制，不计为前缀树工作。每次调用使用独立输出/工作状态，错误不返回部分输出；同一编译值在失败后仍可复用。不声称对恶意重叠模式具有线性时间复杂度。

`replacer.chunks(input, max_bytes, max_work)` 返回 `Result[FnIterator[Result[string, ReplaceError]], ReplaceError]`。负限制在构建时失败；匹配和输出预算检查按需发生。成功元素是按输出顺序排列的非空 UTF-8 片段。连接它们产生与 `replace` 相同的结果，后者内部使用同一迭代器。删除规则不产生片段，但仍消耗匹配工作。片段边界不是稳定 API 契约；一次替换可能是一个大片段。迭代器不分配完整输出缓冲区。

迭代器报告匹配/输出错误一次，之后保持耗尽。已产生片段不回滚。会超出总输出预算的片段不会产生，即使部分也不会。停止迭代可避免扫描余下输入。迭代器别名共享进度，需要串行使用；独立 `chunks` 调用拥有独立状态。迭代器保留输入和编译规则；惰性不代表保留内存配额。例如，无规则替换器在 `a界b` 上使用输出预算 3，会先产生 `a`，再产生输出限制错误，之后无更多元素。新调用可用更大预算重试。

`ReplaceError` 区分负限制、规则数量/规则字节限制、输出和工作限制，以及非法输入标量边界。限制不包括调用方输入和一般分配开销，不是保留内存配额。正常输入是合法 GoML 字符串；任意字节改写属于 bytes API。构造函数先验证规则数量限制，再验证规则字节限制，然后检查实际规则大小；替换在扫描前先检查输出限制，再检查工作限制。例如，以限制 `(2, 5)` 编译 `[("ab", "x"),
("a", "y")]`，在输出和工作预算充足时替换 `aba` 得到 `xy`。这些 API 不增加语法形式或原生后端。

<a id="streaming-replacement-output"></a>

## 流式替换输出

按需导入的 `std::text::stream` 包将替换与 `std::io::Write` 组合；仅导入 `std::text` 不会导入 I/O。其 `write_replaced(replacer, input, writer, max_bytes, max_work)` 返回已写字节数或 `ReplaceWriteError`。它惰性消费 `Replacer::chunks`，通过固定 4096 字节临时缓冲区复制片段，从不构建完整输出缓冲区。大型替换片段分成有界小块写入。短写会补齐，所有 writer 错误包括 `Interrupted` 都立即返回，零进度变为 `WriteZero`，负数或过大 writer 数量变为 `InvalidData`。不隐式刷新或关闭；这些操作由调用方负责。

`ReplaceWriteError::Replacement(count, error)` 保留 `text::ReplaceError`；`Write(count, error)` 保留 I/O 错误。`transferred()` 返回失败前成功且有效的写调用确认的数量，不包括返回错误或违反数量契约的 writer 的未知副作用。已写字节不回滚，结尾可能位于 UTF-8 标量内部。替换预算保留文本语义：超过输出预算的片段会在任何字节到达 writer 前被拒绝。I/O 调用不消耗匹配工作；此 API 不是超时或取消策略。在同一目标上重试整个函数可能重复此前输出；它不是可恢复写游标。

例如，导入 `use std::text::stream` 和 `use std::io` 后，`stream::write_replaced(replacer, "aba", io::Discard::new(), 64, 1024)` 验证并统计替换输出，不保留内容。这些是普通泛型函数和错误变体，没有新语法或原生算法后端。

<a id="bounded-text-building"></a>

## 有界文本构建

`StringBuilder::capacity()` 报告当前字节容量。`reserve_checked(additional, max_bytes)` 确保至少有 `len() + additional` 字节空间，不改变内容或长度。它先检查负限制（`InvalidLimit`），再检查负增量（`InvalidCount`）、算术溢出（`LengthOverflow`）和逻辑总字节上限（`LimitExceeded`）。验证失败也保持容量不变。保留已有较大容量；分配器增长可能超过请求量。此上限不严格限制实际分配，`clear()` 保留已有容量。例如，空构建器可在限制 16 下预留额外 16 字节，长度仍为零。别名观察到相同预留结果。

`write_bytes(bytes)` 在追加前将完整字节值验证为 UTF-8；非法输入保持内容、长度和容量不变。它复制接受的字节，因此后续源修改不改变构建器。`finish()` 返回不可变快照，不受后续写入或 `clear()` 影响。不同于 Go 面向字节的字符串构建器，此文本 API 不跨调用累积非法 UTF-8 片段。任意字节或拆开的编码标量请使用 `bytes::Builder`，完成后再执行带检查的 UTF-8 转换。直接修改公共 `values` 会绕过文本检查，仍由调用方负责。

`text::clone_checked(value, max_bytes)` 使用新的字节向量和现有不可变字符串转换边界，返回内容相等、与源底层存储分离的字符串。负或不足的输出字节限制在复制前失败。小子串不应保留大源时使用它。空/短字符串不必具有全局唯一地址。上限不包括临时分配和分配器开销，也不恢复分配失败。例如，`text::clone_checked("界", 3)` 成功，限制为 2 则失败。

`text::StringBuilder::write_string_checked(value, max_bytes)`、`write_char_checked(character, max_bytes)` 和 `write_line_checked(value, max_bytes)` 在包含已有内容的总输出字节限制内追加。均返回 `Result[(), bytes::TransformError]`：负限制产生 `InvalidLimit`，长度算术溢出产生 `LengthOverflow`，总长度过大产生 `LimitExceeded`。验证先于缓冲增长和修改；行追加将给定文本和末尾 LF 作为一个操作预检查。UTF-8 字符按编码字节而非标量数计数。即使空追加也拒绝已超限的构建器。

限制按每次调用给定，不保存在构建器中。现有不检查方法、公共 `values` 访问和共享可变别名保持不变；调用方必须保持 UTF-8 有效并串行化访问。失败调用保持缓冲不变，但不回滚此前成功调用。这不是分配失败恢复机制，也不限制单独格式化的输入。

数字和引用追加流程通过组合现有转换函数与这些带检查写入完成，不引入第二套转换算法或 Go append ABI。例如，`builder.write_string_checked(number.to_string(), 4096)` 保留已有前缀，并在追加前检查最终总大小。若临时输出也需受限，应使用有界浮点/引用格式化器并给定适当剩余预算。这些是普通方法，无语法变更。

<a id="bounded-quoted-literals"></a>

## 有界引用字面量

`text::quote(value, max_bytes)` 返回带双引号的 Go 风格字符串字面量；`quote_char(value, max_bytes)` 返回带单引号的标量字面量。`quote_ascii` 和 `quote_char_ascii` 变体转义所有非 ASCII 标量。`quote_graphic` 和 `quote_char_graphic` 变体还保留 Unicode 空格分隔符（Zs），如不换行空格和表意空格。六个函数均返回 `Result[string, bytes::TransformError]`。输出限制包含两侧分隔符；即使空字符串也需两字节。负限制返回 `InvalidLimit`，输出过大返回 `LimitExceeded`，最终输出缓冲区分配前检查长度算术。

与分隔符相同的引号及反斜杠会转义。响铃、退格、制表符、LF、垂直制表符、换页和 CR 分别使用 `\a`、`\b`、`\t`、`\n`、`\v`、`\f` 和 `\r`；其余 ASCII 控制字符使用两位 `\x` 转义。不可打印 Unicode 标量使用四位 `\u` 或八位 `\U` 转义，十六进制数字为小写。普通引用保留 Unicode 15.0.0 可打印标量；Graphic 引用使用相同固定 Unicode 版本，保留图形标量，但仍转义引号、反斜杠、控制字符和非图形值。ASCII 引用转义所有非 ASCII 标量，包括可打印字符。输入是合法 GoML 字符串/字符，因此没有非法字节或代理码输入模式。

`text::can_backquote(value)` 应用 Go 的保守单行判定策略，判断能否在不改变内容的情况下用反引号包围字符串。它拒绝反引号、U+FEFF、DEL 和除制表符外的 ASCII 控制字符（包括 CR 和 LF）。允许其他合法非 ASCII 标量，即使不是图形字符；这不是终端安全谓词。它只扫描，不生成输出字符串。例如，`can_backquote("a\tb")` 为 true，`can_backquote("a\nb")` 为 false。结果针对交换字面量，不针对 GoML 源码语法。

这些 API 编码 Go 风格交换字面量，不是 JSON、shell 命令、HTML 或 GoML 源码语法。特别是完整 Go 转义集合并不是 GoML 源码新接受的转义集合。不隐含解析、规范化、区域规则或安全净化。输出限制不包括已有输入和有界的逐标量临时存储；不恢复分配失败。

`text::unquote_bytes(literal, max_bytes)` 将恰好一个完整 Go 风格双引号、单引号或反引号字面量解码为独立 `bytes::Bytes`。它拒绝尾随文本，不只解析前缀。双/单引号接受上述短转义、`\x` 后恰好两位十六进制、`\u` 后四位、`\U` 后八位，或恰好三位八进制。八进制值必须可放入一个字节；Unicode 转义必须表示合法标量（不能是代理码或超过 U+10FFFF 的值）。只允许转义当前分隔符对应的引号。双/单引号内拒绝未转义 LF。单引号最多接受一个字面标量或转义，包括字节转义；接受空 `''` 形式以匹配 Go Unquote 行为，而非 Go 源码字符字面量语法。反引号内容按字面处理，但丢弃 CR 字节；保留 LF。

`text::quoted_prefix(input, max_bytes)` 验证并返回字节零处的首个完整引用字面量，包括分隔符和原始转义。不检查尾随文本；按返回字符串字节长度前进即可取得剩余内容。不跳过前导空白。限制统计返回前缀的源字节，不是解码字节：即使 `""` 也需两字节，原始 CR 字节仍保留并计入限制。不同于去引用，此操作不构建解码主体，也不要求字节转义组成 UTF-8。它使用有界的逐转义临时存储。

错误使用 `UnquoteError`，先检查负限制。验证起始分隔符后，扫描在预算处停止；`OutputLimit(offset)` 指向首个无法容纳的元素或闭合分隔符。转义验证可能在报告错误前检查超出限制的常量大小转义。语法错误使用源字节偏移，不返回部分前缀。这是 Go 风格交换字面量解析器，不是新的 GoML token 或语法产生式。例如，`quoted_prefix("\"a\"tail", 3)` 返回 `Ok("\"a\"")`。

`text::unquote(literal, max_bytes)` 额外要求完整解码结果是合法 UTF-8，并返回字符串。因此字节转义可以重建多字节 UTF-8，但单独 `\xff` 只能通过 `unquote_bytes` 获取。两个完整输入解码器都返回 `UnquoteError`：负限制为 `InvalidLimit(value)`，畸形语法为 `Syntax(offset)`，下一解码元素会超限为 `OutputLimit(offset)`，或（对 `unquote`）`InvalidUtf8(offset)`。语法和输出限制偏移是零起始输入字节位置；畸形转义指向反斜杠，缺失闭合引号指向输入末尾。UTF-8 错误偏移则指向解码字节序列。诊断不回显输入内容。负限制优先；否则按遍历顺序检查语法和输出限制，成功完整解析后验证 UTF-8。错误不暴露部分结果。

`text::unquote_element(input, context)` 仅解码首个字面标量或转义，返回 `Result[(QuotedElement, string), UnquoteError]`；第二个值是未经检查、保持不变的尾部。`QuoteContext::{Unquoted, Single, Double}` 选择引用规则。Single/Double 拒绝自身未转义分隔符，仅允许该分隔符的引号转义。Unquoted 允许任一种字面引号，但不允许任何引号转义。这里反引号是普通字节，不是原始字面量模式。

`QuotedElement::Byte(byte)` 表示字面 ASCII、短转义和十六进制或八进制字节转义。`Scalar(char)` 表示字面非 ASCII 字符和 Unicode 转义，包括 `\u0041` 等值为 ASCII 的转义。因此 `\xff` 是一个字节，不是要求将 U+00FF 编码为 UTF-8。不同于完整字面量解析器，此原语接受字面 LF 和 CR：外围字面量规则由调用方实施。空或畸形输入产生 `Syntax(0)`，不返回部分结果。它最多检查一个标量/转义（十个输入字节），使用有界临时存储，因此没有可变大小解码输出限制。不引入源码语法形式，也不处理非法 UTF-8 输入字符串。例如，`unquote_element("\\xffrest", QuoteContext::Unquoted)` 返回 `Ok((QuotedElement::Byte(255), "rest"))`。

每次追加前执行解码输出限制，包括每个完整标量编码。它不限制输入扫描（例如丢弃的原始 CR）或分配开销；处理不可信输入的调用方还必须限制输入大小。这些是内存解码器，不是增量流 reader，也不是 GoML 源码语法扩展。

```goml
use std::text;
use std::bytes;

fn quoted_name(value: string) -> Result[string, bytes::TransformError] {
    text::quote_ascii(value, 4096)
}

fn decoded_name(literal: string) -> Result[string, text::UnquoteError] {
    text::unquote(literal, 4096)
}
```

<a id="numeric-parsing-and-checked-arithmetic"></a>

## 数值解析与带检查算术

`num::format_float32_fixed(value, decimal_places, max_bytes)` 和 `format_float64_fixed` 返回 `Result[string, num::FormatFloatError]`，小数点后恰好具有请求的非负位数。零位时省略小数点。转换从 IEEE 位开始，通过纯整数算术构建精确十进制值，再按最近值、中点取偶舍入一次。格式化的是所表示的二进制值，不是猜测的原始十进制输入：例如 2.675 保留两位时遵循实际 f64 值。负零及舍入到零的负值保留负号。

非有限值产生 `NaN`、`+Inf` 或 `-Inf`；合法精度不会填充这些 token。这种写法仅适用于这些 API，不改变现有最短 `ToString` 输出。限制包含符号、小数点和所有数字。负限制在验证精度前返回 `InvalidLimit`；负小数位数返回 `InvalidPrecision`，即使是 NaN/无穷也如此。输出过大返回 `OutputLimit`，不返回部分字符串。分配缓冲区前检查完整输出长度；有界精确十进制临时存储依赖 IEEE 宽度，不依赖请求精度。不恢复分配失败。不引入原生格式化后端、新语法或浮点环境/舍入模式依赖。

`num::format_float32_scientific(value, decimal_places, uppercase, max_bytes)` 和 `format_float64_scientific` 使用相同精确转换和舍入，小数点前一位，后恰好 `decimal_places` 位。`uppercase` 选择 `E` 而非 `e`；指数总带符号，至少两位十进制数字。舍入可能增加指数（9.5 保留零位变为 `1e+01`）。零使用指数零，若有负号则保留。

`num::format_float32_general(value, significant_digits, uppercase, max_bytes)` 和 `format_float64_general` 舍入到请求的非负有效数字位数，零按一处理。它们移除无意义的尾随零，然后在舍入后的十进制指数小于 -4 或大于等于有效精度时使用科学计数法，否则使用定点记法。`uppercase` 只改变科学计数法指数标记。大精度不会强制填充或按精度成比例分配：精确数字可容纳后，就在输出限制内保留它们。负精度是错误，不是请求最短格式化；现有 ToString 仍提供既定最短表示。这两类函数与定点格式化共享 `FormatFloatError`、验证顺序、非有限值写法和完整输出预检查。

`num::format_float32_binary(value, max_bytes)` 和 `format_float64_binary` 输出精确十进制整数有效数，后接 `p` 和带符号、不补零的二进制指数。有效数是 IEEE 整数尾数，不是二进制数字字符串。非正规数保留未规范化尾数。零同样使用非正规数尺度（f32 为 `0p-149`，f64 为 `0p-1074`）；负零保留符号。此表示无精度参数。

`num::format_float32_hex(value, fractional_digits, uppercase, max_bytes)` 和 `format_float64_hex` 输出带二进制指数的规范化十六进制有效数：非零有限值为 `0x1...p±dd`，零为 `0x0...p+00`。大写模式选择 `0X`、十六进制 A–F 和 `P`。指数至少两位十进制数字。非负精度恰好给出相应数量的小数十六进制位，中点取偶舍入，进位时调整指数。精度 `-1` 则输出不含无意义尾随零的精确值；更低精度非法。非正规数会规范化，不会舍入丢弃。大精度仅在检查完整输出字节限制后追加零。

两类函数都使用位/整数操作，与十进制格式化器共享非有限值写法、符号保留、可恢复错误和分配限制。负输出限制先于精度验证被拒绝，非法十六进制精度即使对非有限值也被拒绝。有界尾数和指数临时字符串独立于请求输出大小。不引入新的原生格式化原语或源码语法。

`num::format_float32_shortest(value, notation, max_bytes)` 和 `format_float64_shortest` 返回能解析回原始 f32/f64 位模式的最短有效十进制数字，按 `FloatNotation` 渲染：`Fixed`、`Scientific`、`ScientificUpper`、`General` 或 `GeneralUpper`。这不保证跨不同记法的总字符最少。定点输出可能需要数百个前导/尾随零；科学计数法保留带符号、至少两位的指数。General 最短输出在指数小于 -4 或大于等于 6 时选择科学计数法，与有效数字数量无关；这区别于显式精度的 general 格式化。

转换使用精确十进制数字，按递增有效位数，通过现有纯 GoML 标量解析器检查最近/较低/较高候选。f32 最多检查九种位数，f64 最多十七种，保留负零。最近候选采用中点取偶舍入；相邻候选覆盖非对称舍入区间。若意外找不到匹配候选，返回 `FormatFloatError::RoundtripFailure`，不回退为非最短表示，也不 panic。非有限值写法、负限制拒绝和输出预检查与其他格式化器一致。输出限制不包含有界候选临时存储和转换工作；现有编译器拥有的 ToString 回退和源码语法不变。

```goml
use std::num;

fn fixed_measurement(value: f64) -> Result[string, num::FormatFloatError] {
    num::format_float64_fixed(value, 3, 64)
}

fn compact_measurement(value: f64) -> Result[string, num::FormatFloatError] {
    num::format_float64_general(value, 6, false, 64)
}

fn exact_hex_measurement(value: f64) -> Result[string, num::FormatFloatError] {
    num::format_float64_hex(value, -1, false, 64)
}

fn shortest_measurement(value: f64) -> Result[string, num::FormatFloatError] {
    num::format_float64_shortest(value, num::FloatNotation::General, 64)
}
```

`num::parse_bool_structured(value)` 返回 `Result[bool, num::ParseBoolError]`。它仅接受 `1`、`t`、`T`、`true`、`TRUE`、`True` 为 true，`0`、`f`、`F`、`false`、`FALSE`、`False` 为 false。不执行修剪、通用大小写折叠、数值强制转换、本地化写法或 `yes`/`no` 约定。畸形输入返回错误，绝不默认 false。`error.input()` 保留原字符串供显式检查；其 ToString/Debug 消息不回显输入。错误相等性比较原始输入。若保留错误字符串代价较高，解析不可信数据前应限制输入大小。现有布尔 `ToString` 实现仍是标准格式化器，产生小写 `true`/`false`，无需单独格式化 API。

```goml
use std::num;

fn parse_switch(value: string) -> Result[bool, num::ParseBoolError] {
    num::parse_bool_structured(value)
}
```

`num::format_int(value: i64, radix)` 和 `format_uint(value: u64, radix)` 使用纯 GoML 整数算术，返回 `Result[string, FormatIntError]`。格式化接受 2 到 36 进制，输出小写数字，不含前缀、分隔符或正号，零格式化为 `"0"`。负有符号数使用 `-` 后接绝对大小，包括最小 i64；这不是二进制补码显示。格式化不接受进制零，尽管解析支持它。

`write_int(output: MutSlice[byte], value: i64, radix)` 和 `write_uint` 返回写入的 ASCII 字节数。它们在修改目标前验证：非法进制或容量不足保持所有字节不变，成功时未使用后缀不变。`FormatIntError` 区分 `InvalidRadix(radix)` 和 `BufferTooSmall(needed, available)`；进制验证优先。这些辅助函数分配有界临时数字向量（最多 65 输出字节），不保证无分配转换。现有解析、标量十进制 to_string 和浮点行为不变。不增加语法或编译器/运行时内建操作。

数值解析返回 `Result[_, num::ParseIntError]` 或 `Result[_, num::ParseFloatError]`。整数进制解析接受进制 `0` 或 `2..36`；进制 `0` 识别 `0b`、`0o` 和 `0x` 前缀，并允许 Go 风格数字分隔符。非法进制、畸形输入和溢出返回 `Err`。浮点解析支持十进制和十六进制 IEEE 754 输入、带符号指数、数字分隔符、`inf`、`infinity` 和 `NaN`，直接舍入到请求的 `f32` 或 `f64` 宽度。

对于 `parse_float32_structured` 和 `parse_float64_structured`，`ParseFloatError::is_range()` 区分数值溢出与畸形语法。溢出返回带消息 `value out of range` 的 `Err`，不成功返回无穷或饱和数值。显式无穷写法是成功值。下溢舍入为非正规数或带符号零，并成功。畸形语法（包括巨大指数后的非法尾随字符）具有 `is_range() == false` 和消息 `invalid syntax`。两类错误都保留现有 `kind() == io::ErrorKind::InvalidInput`、输入/上下文、操作和原始操作系统代码接口。该标志参与错误相等比较。兼容构造函数 `ParseFloatError::new(input, message)` 保持标志为 false，不从调用方提供的诊断文本推断分类。不增加语法或运行时原语；解析继续使用现有纯 GoML 标量转换实现。这些 API 无输入长度限制，因此调用方必须在解析或保留错误前限制不可信数值字符串。有理数转换阶段在分配移位整数临时值前检查保守的二进制指数边界。因此确定的溢出和下溢不会分配与巨大二进制指数成比例的存储；边界情况仍使用精确舍入，包括最小非正规数一半处的中点。这不限制扫描和累积长尾数所需工作或存储。

长尾数按数值解释，不截断到固定十进制数字缓冲区。例如，`1` 后跟 2048 个零和 `e-2048`，在两种宽度下都解析为恰好一。这有意区别于 Go 1.26 对某些长输入的固定缓冲区回退；这些情况使用精确有理数测试参考，不将每个 Go 结果视为权威。指数累积只在超过由输入长度推导、且大于尾数可能补偿尺度的边界后饱和。因此长小数零串可以被大指数抵消，不会因任意固定指数截止值而改变数值。饱和后仍检查剩余指数数字和分隔符的语法；非法后缀不是溢出。

```goml
use std::num;

fn parse_float_input(value: string) -> Result[f64, string] {
    num::parse_float64_structured(value).map_err(|error: num::ParseFloatError| {
        if error.is_range() {
            "outside f64 range"
        } else {
            "invalid numeric syntax"
        }
    })
}
```

`num::parse_int_structured`、进制和无符号变体，以及结构化浮点解析器返回领域解析错误。`checked_*_int64` 操作溢出时返回 `None`；对应 `saturating_*_int64` 操作钳位到有符号 64 位边界。

`num::parse_int_bits(input, radix, bits)` 和 `parse_uint_bits` 分别返回 i64 和 u64，受请求的有符号或无符号范围约束。接受宽度 1..64；零选择 GoML 的 64 位原生整数宽度。因此有符号宽度 1 只接受 -1 和 0，无符号宽度 1 接受 0 和 1。现有进制、前缀、符号和下划线规则适用；无符号解析不接受符号。这些函数失败时从不截断或返回饱和值。

`IntWidthError::InvalidWidth(bits)` 优先于输入/进制验证。`IntWidthError::Parse(error)` 为语法、非法进制或全宽溢出保留现有 ParseIntError。64 位成功解析但超出更窄请求范围的值产生 `OutOfRange(effective_bits)`。此有意采用的验证顺序可能区别于 Go 错误类别优先级；调用方获得 Result，而不是 Go 的值加错误对。不改变现有解析器签名、强制转换规则、字面量推断或编译器/运行时边界。

有符号一位范围按数学定义实施：小于 -1 的值是错误。Go 1.26 的 ParseInt 对某些溢出的一位负输入可能成功报告 -1；GoML 有意不复现该饱和特例。

`num::parse_complex_f32(input)` 和 `parse_complex_f64(input)` 返回 `Result[(f32, f32), ParseComplexError]` 和 `Result[(f64, f64), ParseComplexError]`。元组为 `(real, imaginary)`；后缀命名每个分量宽度，不是 Go 的复数总宽度。这些是数值转换函数，不是第二种复数算术类型或源码字面量语法。每个分量复用对应纯浮点解析器，包括直接 f32 舍入和带符号零保留。

接受形式为 `N`、`Ni`、`N+Ni`、`N-Ni`，可选地包围一对圆括号，不修剪空白。数值分量使用现有十进制/十六进制、分隔符、无穷和 NaN 规则。纯虚数形式仍需数值系数：`i`、`+i` 和 `-i` 都是错误。连接用加号可位于负分量（`1+-2i`）或 NaN（`1+NaNi`）之前；双加号和负 NaN 非法。省略分量为正零。`ParseComplexError` 为 `Syntax` 或 `OutOfRange`，通过 `is_range()` 分类；任一分量的畸形语法优先于另一分量的溢出。错误不返回饱和/部分分量，诊断不保留或回显输入。无输入长度限制；解析前应限制不可信字符串。例如，`parse_complex_f64("(1-2i)")` 返回 `Ok((1.0, -2.0))`。

`num::format_complex_f32(real, imaginary, format, max_bytes)` 和 `format_complex_f64` 返回 `Result[string, FormatFloatError]`，始终将两个分量写成 `(real+imaginaryi)` 或 `(real-imaginaryi)`。输入分量的位宽由函数名指定；不涉及复数运行时类型。`ComplexFormat` 可选择 `Shortest(FloatNotation)`、`Fixed(decimal_places)`、`Scientific(decimal_places, uppercase)`、`General(significant_digits, uppercase)`、`Binary` 或 `Hex(fractional_digits, uppercase)`。分量精度、舍入、非有限值标记和验证规则与相应的标量格式化函数完全相同。最短形式需显式选择；只有 Hex 使用 -1 表示精确且去除多余零的输出。

长度限制包括括号、分隔正负号和末尾的 `i`。负零和负虚部保留负号，正无穷本身已带加号，虚部 NaN 前添加加号。返回首个分量错误，不产生部分输出；负输出上限的检查先于精度验证。组装字符串之前，会以不溢出的方式检查完整格式的长度。每个临时分量字符串分别受 max_bytes 限制；这是最终输出限制，不限制总分配量或转换工作量。例如，`format_complex_f64(1.0, -2.0, ComplexFormat::Fixed(1), 16)` 返回 `Ok("(1.0-2.0i)")`。二进制有效数表示是一种显示格式，复数十进制/十六进制解析器不接受它。

`std::num::ToFloat` 为每种标量整数位宽提供 `.to_f32()` 和 `.to_f64()`。转换直接舍入到目标 IEEE 754 位宽，使用最近值、恰好居中时取偶数规则；`u64` 到 `f32` 不经过 `f64` 舍入。大整数可能丢失精度。

`std::num::TryToInt` 为 `f32` 和 `f64` 提供 `.try_to_isize()`、`.try_to_i8/i16/i32/i64()` 及对应的无符号方法。它们先向零截断，再检查目标范围；NaN/无穷返回 `FloatConversionError::NonFinite`，溢出返回 `OutOfRange`。因此 `(-0.9).try_to_u64()` 成功并得到零。`float_to_i64(value, rounding)` 和 `float_to_u64(value, rounding)` 接受 `Rounding::{TowardZero, Floor, Ceil, NearestAway, NearestEven}`，先显式舍入，再执行相同范围检查。这些方法需要导入相应 trait，不增加强制转换语法。

```goml
use std::num::{ToFloat, TryToInt};

let count: u64 = 9007199254740993;
let approximate: f64 = count.to_f64();
let truncated = (12.75).try_to_i32();
```

<a id="c-compatible-strings"></a>

## C 兼容字符串

`std::c::CString` 存储不含内部 NUL 的字节。`CString::new(string)`、`from_bytes(bytes::Bytes)` 和 `from_raw(ffi::String)` 返回 `Result[CString, c::Error]`，通过 `Error::InteriorNul` 拒绝内部 NUL。`bytes()` 返回字节副本，`text()` 验证 UTF-8 并返回 `Result[string, utf8::Utf8Error]`，`as_raw()` 则公开底层 `ffi::String`，供显式外部声明使用。原生内存分配、终止符和生命周期由消费此值的适配器负责。

`c::Error` 有 `InteriorNul` 和 `Native(string)` 变体，实现了 `Debug`、`PartialEq`、`Eq` 和 `ToString`。当 Go 错误为 nil 时，`c::check_error(ffi::Error)` 返回 `Ok(())`，否则返回 `Err(Error::Native(message))`。错误消息不是有效 UTF-8 时，使用文本 `C boundary error contains invalid UTF-8`。

`c::literal(string) -> string` 是一个 `#[comptime]` 函数，通过 `compile_error` 拒绝内部 NUL。常量初始化器或 `comptime` 块会请求编译期验证：

```goml
use std::c;

const DATABASE: string = c::literal(":memory:");

fn database_name() -> Result[c::CString, c::Error] {
    c::CString::new(DATABASE)
}
```

`CString` 构造函数也在运行时强制保证其不变式。普通的运行时 `c::literal` 调用不会被强制进行编译期求值，输入无效时会触发陷阱。这些操作不检查 C 头文件，也不确定原生布局。

<a id="environment-paths-and-processes"></a>

## 环境、路径和进程

`env::current_dir_structured`、`current_exe_structured`、`var_structured`、`path::absolute_structured` 以及 `process` 的结构化执行方法公开整个操作的错误。进程超时方法接受 `time::Duration`，通过命令运行时终止并等待进程，并通过 `process::Error` 返回 `TimedOut`。

路径仍是 UTF-8 `string` 值。`path::separator` 返回宿主分隔符，`components` 识别两种斜杠形式，`relative` 使用宿主路径规则，`windows_prefix` 则独立于宿主操作系统识别驱动器和 UNC 前缀。非 UTF-8 操作系统名称无法表示，因此不能出现在这些 API 中。

<a id="linux-amd64-system-calls"></a>

## Linux amd64 系统调用

`std::os::linux::syscall` 是一个显式的底层访问接口，仅支持 Linux amd64。`syscall6(number: usize, args: [usize; 6]) -> Result[SyscallResult, Error]` 接受任意系统调用号和六个机器字。未使用的参数位置应填零。有符号参数使用其机器字位模式，例如 Linux `AT_FDCWD` 使用 `(-100).to_usize()`。

`SyscallResult` 的公开字段 `r1`、`r2` 和 `errno` 均为 `usize`，保留 Go `syscall.Syscall6` 返回的值。`is_ok()` 检查 `errno == 0`。内核失败仍返回 `Ok(SyscallResult)`，但 errno 非零；调用方必须检查它。外层 `Err(Error::UnsupportedTarget)` 在发出任何系统调用前报告执行目标不是 Linux amd64。不保证其他 Go 目标能够编译。系统调用号、布局和常量特定于 Linux amd64 ABI；本包不会为其他架构选择调用号。

`SyscallResult::into_result()` 将成功时的第一个返回机器字转换为 `Ok(usize)`，将非零 errno 转换为 `Err(Errno)`。`Errno(pub usize)` 中的公开元组字段保留数字代码，包括未来未知的代码。`Errno::name()` 和 `errno_name(usize)` 返回包含规范符号名称的 `Option[string]`；别名解析到原始名称，例如 `EWOULDBLOCK` 解析为 `EAGAIN`。`to_string()` 返回该名称，未知代码则返回 `"errno N"`。零没有 errno 名称。

```goml
use std::os::linux::syscall;

fn process_id() -> Result[usize, string] {
    let result = syscall::syscall6(syscall::SYS_GETPID, [0, 0, 0, 0, 0, 0]).map_err(
        |error| error.to_string(),
    )?;
    if result.is_ok() {
        Ok(result.r1)
    } else {
        Err("getpid errno " + result.errno.to_string())
    }
}
```

`syscall6_with_buffers(number: usize, args: [Arg; 6])` 返回相同的结果类型。`Arg::Word(usize)` 原样传递整数。`Arg::Buffer(MutSlice[byte])` 传递视图首字节地址，空视图则传递空指针。运行时保留并固定每块非空底层分配，直到调用返回；内核的写入可通过该视图观察到。支持多个缓冲区、子视图以及同一分配的别名。缓冲区参数不会隐式添加长度参数或末尾 NUL 字节。

```goml
use std::os::linux::syscall;
use std::os::linux::syscall::{Arg};

fn write_once(fd: usize, data: Vec[byte]) -> Result[syscall::SyscallResult, syscall::Error] {
    syscall::syscall6_with_buffers(syscall::SYS_WRITE, [
        Arg::Word(fd),
        Arg::Buffer(data.as_mut_slice()),
        Arg::Word(data.len().to_usize()),
        Arg::Word(0),
        Arg::Word(0),
        Arg::Word(0),
    ])
}
```

`syscall6_with_pointers(number: usize, args: [Arg; 6], pointers: Slice[Pointer])` 将缓冲区 API 扩展到包含指针的同步原生结构，包括 `iovec` 数组和 `msghdr` 记录。`Pointer` 包含公开字段 `buffer: MutSlice[byte]`、`offset: isize` 和 `target: MutSlice[byte]`。每个描述符临时将目标地址写入缓冲区视图中指定偏移处的八字节小端指针字段；目标为空时写入零。根缓冲区使用 `Arg::Buffer`，描述符可在任意深度连接中间缓冲区和其他缓冲区，包括别名和循环。所有参与的分配都在设置指针字段之前固定，并保持固定直到调用和指针恢复完成。嵌套缓冲区数量不限于六个。

任何修改或系统调用之前都会检查每个指针槽的边界。槽无效时返回 `Error::InvalidPointer { index, offset, length }`，指出描述符及其缓冲区视图。描述符按列表顺序设置。返回时恢复所有原始指针字段字节，包括内核出错的情况；重叠槽也恢复为原始字节。内核在这些槽外的写入仍然可见。应将指针槽视为临时输入字段：内核写入槽中的输出在恢复时会被丢弃。调用期间，其他线程不得并发读取、修改这些缓冲区和指针字段，也不得将其用于另一个系统调用。

例如，Linux amd64 的 `iovec` 记录包含八字节指针，后跟八字节长度。以下代码使用 [readv/writev](https://man7.org/linux/man-pages/man2/readv.2.html) 所记载的原生布局，通过一次内核操作写入两个视图：

```goml
use std::bytes::endian;
use std::os::linux::syscall;
use std::os::linux::syscall::{Arg, Pointer};

fn write_pair(fd: usize, first: MutSlice[byte], second: MutSlice[byte]) -> Result[syscall::SyscallResult, syscall::Error] {
    let vectors: Vec[byte] = Vec::new();
    for _ in 0..32 { vectors.push(0); }
    let _ = endian::write_u64(vectors.as_mut_slice(), 8, first.len().to_u64(), endian::Endian::Little);
    let _ = endian::write_u64(vectors.as_mut_slice(), 24, second.len().to_u64(), endian::Endian::Little);
    let pointers = Vec::from_array([
        Pointer { buffer: vectors.as_mut_slice(), offset: 0, target: first },
        Pointer { buffer: vectors.as_mut_slice(), offset: 16, target: second },
    ]);
    syscall::syscall6_with_pointers(syscall::SYS_WRITEV, [
        Arg::Word(fd), Arg::Buffer(vectors.as_mut_slice()), Arg::Word(2),
        Arg::Word(0), Arg::Word(0), Arg::Word(0),
    ], pointers.as_slice())
}
```

调用方负责系统调用 ABI、有效地址、缓冲区长度、对齐、原生结构编码、资源清理和同步。字节缓冲区可以承载显式编码的原生记录；普通 GoML 结构体不保证内核布局。指针槽边界检查不会验证特定系统调用的结构布局或记录中存储的长度。`Word` 不保留或固定任何 Go 分配，因此不得通过整数夹带 Go 管理的地址。`mmap` 等操作返回的数字地址仍是原始机器字，需要使用 `munmap` 等适当操作显式清理。

缓冲区生命周期仅覆盖同步调用，也包括通过 `Pointer` 提供的嵌套指针图。返回后仍保留指针的操作和任意 Go 对象内存不在此缓冲区 API 的范围内。运行时使用感知调度的 `Syscall6`，而非 `RawSyscall6`。它只执行一次调用，不自动重试 `EINTR`，也不补全部分读取/写入。线程专属操作需要线程亲和性处理，而此 API 不提供这种能力。直接改变线程、创建进程、修改信号处理器或 Go 运行时地址空间可能违反运行时不变式；接受调用号不代表每个内核操作都能从 GoML 安全使用。调用仍受内核可用性、进程权限和沙箱策略约束。

本包导出固定版本 [Linux 7.2 系统调用表](https://github.com/torvalds/linux/blob/v7.2/arch/x86/entry/syscalls/syscall_64.tbl) 中全部 385 个原生 amd64 `SYS_*` 调用号，不包括独立的 x32 ABI。这包括旧接口和 `SYS_OPENAT2`、`SYS_CLONE3`、`SYS_PIDFD_OPEN`、`SYS_IO_URING_SETUP`、`SYS_LANDLOCK_CREATE_RULESET` 等新接口。常量存在不保证运行中的内核实现或允许该调用。还导出了同一版本 `asm-generic/errno-base.h` 和 `errno.h` 中全部 136 个 errno 常量及别名。仍可直接传入其他调用号。数字清单和上游 SHA-256 校验和位于 `tools/syscall/linux-amd64.json`；`python3 tools/syscall/generate.py` 重新生成常量，`--check` 离线验证，`--check --verify-upstream` 还检查原始内核文件。本包使用普通导入、函数、枚举、数组和可变切片，不引入新文法或 `unsafe` 语法，不能在 `comptime` 中使用。

<a id="linux-amd64-abi-records"></a>

## Linux amd64 ABI 记录

`std::os::linux::abi` 提供带公开字段的记录，具有 `Type::size() -> isize`、`encode() -> Vec[byte]` 和 `Type::decode(Slice[byte]) -> Result[Type, endian::BoundsError]`。编码使用 Linux amd64 小端布局，填充字节置零。解码要求包含完整记录，允许尾随字节。这些编解码器验证字节边界；时间范围和标志等语义值由内核验证。普通 GoML 结构体仍不保证原生布局。

| 记录 | 编码大小（字节） |
| --- | --- |
| `Timespec { seconds, nanoseconds }`、`Timeval { seconds, microseconds }` | 各 16 |
| `Itimerspec { interval, value }` | 32 |
| `Rlimit { current, maximum }` | 16 |
| `PollFd { fd, events, revents }`、`EpollEvent { events, data }` | 8、12（紧凑布局） |
| `Iovec { base, length }`、`Msghdr`、`Cmsghdr { length, level, kind }` | 16、56、16 |
| `Flock { kind, whence, start, length, pid }`、`OpenHow { flags, mode, resolve }` | 32、24 |
| `Stat`、`Statx`、`StatxTimestamp { seconds, nanoseconds }` | 144、256、16 |
| `Rusage` | 144 |

`Stat` 公开设备/inode/链接标识、模式和所有权、大小/块信息，以及作为 `Timespec` 的 `atime`、`mtime`、`ctime`。`Statx` 还公开 `mask`、属性、创建时间、设备主/次标识、挂载 ID 和直接 I/O 对齐。使用可选字段前，应检查其返回的掩码。直接 I/O 对齐前缀之后的字段目前由此编解码器保留。`Rusage` 包含用户/系统 `Timeval` 值和 Linux 资源计数器。记录中的指针值字段是机器字；引用 Go 管理的存储时应使用嵌套缓冲区构建器。

`Iovecs::new(buffers: Slice[MutSlice[byte]]) -> Result[Iovecs, string]` 最多接受 1024 个视图，创建编码后的向量数组和指针描述符。`buffer()`、`pointers()` 和 `len()` 为 `readv`、`writev` 及相关调用提供参数。`Message::new(name: MutSlice[byte], vectors: Iovecs, control: MutSlice[byte]) -> Result[Message, string]` 为同步 `sendmsg`/`recvmsg` 构建 `msghdr` 图；空的 name/control 视图表示空指针。其 `buffer()` 和 `pointers()` 与 `syscall6_with_pointers` 配合使用，`header()` 在调用后解码长度和标志。该调用会恢复指针槽。构建器保留视图，不复制内容。调用正在使用构建器时，不得并发访问它、它的视图及其图；修改嵌入的长度需要遵循原始系统调用 ABI。

布局清单位于 `tools/syscall/linux-amd64-layouts.json`。运行 `python3 tools/syscall/layouts.py`，再在 `lib/std` 中运行 `goml fmt`，即可重新生成编解码器。`python3 tools/syscall/layouts.py --check --verify-headers` 检查规范输出，并针对 Linux amd64 系统头文件编译每个大小、偏移和字段位宽的 C 断言。

<a id="linux-amd64-file-descriptors"></a>

## Linux amd64 文件描述符

`std::os::linux::fd` 提供显式描述符所有权，返回 `Result[..., io::Error]`，并重新导出 `Error` 和 `ErrorKind`。内核提供 errno 时，错误保留 `raw_os_code()`。可被中断的调用会重试 `EINTR`，但 `close` 和行为取决于请求的原始 `ioctl` 除外。非阻塞操作在 `EAGAIN` 时返回 `WouldBlock`。这些接口不增加就绪等待或取消机制。

`open(path: string, flags: usize, mode: u32)`、`open_at(directory: isize, path: string, flags: usize, mode: u32)` 和 `open_at2(directory: isize, path: string, how: abi::OpenHow)` 返回 `Fd`。`open_at_bytes` 和 `open_at2_bytes` 接受原始 `Slice[byte]` 路径。路径拒绝嵌入的 NUL，其他字节保持不变；字符串路径编码为 UTF-8。`Fd` 也有 `open_at`、`open_at_bytes` 和 `open_at2` 方法，在操作期间保持目录描述符有效。`memfd(name: string, flags: usize)` 创建内存支持的匿名文件。打开、memfd 和复制操作原子设置 close-on-exec。标志包括 `O_*`、`AT_*`、`RESOLVE_*`、`MFD_*` 和 `F_SEAL_*`；标志或系统调用常量存在不保证内核/文件系统支持。

赋值或传递 `Fd` 会共享其关闭状态。`close()` 是显式且幂等的；没有终结器。它拒绝新操作，等待活动操作/借用完成，然后关闭原始描述符一次，关闭失败时不重试。关闭开始时，`is_closed()` 变为 true。阻塞 I/O 必须完成后 close 才能返回；关闭不会取消它。同一个打开句柄上的操作可以并发运行。调用方必须同步共享缓冲区，以及涉及共享文件偏移的任何复合操作序列。

`duplicate()` 创建独立的描述符所有者，同时保留内核共享的打开文件描述和偏移。`Fd::from_raw(raw: isize)` 验证并接管现有描述符，不改变其标志；成功时，调用方转移唯一的关闭责任。`Fd::from_owned_raw(raw: isize)` 只检查数字范围，接管已拥有的描述符，不进行 `fcntl` 探测；应将其用于成功的系统调用结果，并确保数字有效且独占。`into_raw()` 等待活动操作结束，将该责任转移回来，使所有别名失效。`with_raw_fd[T](action: (usize) -> Result[T, Error])` 在回调期间保持描述符打开。`fd::with_raw_fds[T](descriptors: Slice[Fd], action: (Slice[usize]) -> Result[T, Error])` 以迭代方式借用整个列表，包括重复别名；若后续获取或回调失败，会释放所有已成功的借用。回调不得关闭、替换、保留或重新接管原始编号，也不得对同一共享句柄调用 `close`/`into_raw`。接受原始目录编号的自由函数要求调用方保持其有效；对拥有的目录使用 `with_raw_fd`，当前目录使用 `AT_FDCWD`。

| `Fd` 上的操作 | 行为 |
| --- | --- |
| `read(MutSlice[byte])`、`write(Slice[byte])` | 返回传输的字节数；读取在 EOF 时返回零 |
| `read_at(output, offset: i64)`、`write_at(input, offset: i64)` | 不移动文件偏移的定位 I/O |
| `read_vectored(Slice[MutSlice[byte]])`、`write_vectored(Slice[Slice[byte]])` | 分散/聚集 I/O，最多 1024 个缓冲区 |
| `read_vectored_at(outputs, offset: i64)`、`write_vectored_at(inputs, offset: i64)` | 完整 64 位定位分散/聚集 I/O |
| `read_exact(output)`、`write_all(input)` | 循环处理短传输；停止推进时报告 `UnexpectedEof` 或 `WriteZero` |
| `seek(offset: i64, whence: usize)` | 返回新的 `u64` 位置；支持 `SEEK_SET/CUR/END/DATA/HOLE` |
| `stat()`、`truncate(length: i64)`、`chmod(mode: u32)`、`chown(uid: u32, gid: u32)` | 元数据和文件修改；`0xffffffff` 保持对应所有权 ID 不变 |
| `sync()`、`sync_data()`、`allocate(mode: usize, offset: i64, length: i64)` | `fsync`、`fdatasync` 和带 `FALLOC_FL_*` 标志的 `fallocate` |
| `descriptor_flags()`、`set_descriptor_flags(flags)`、`status_flags()`、`set_status_flags(flags)` | 描述符标志和打开文件状态标志 |
| `flock(operation: usize)`、`record_lock(command: usize, lock: abi::Flock)` | 整文件锁和 POSIX/OFD 字节范围锁；后者返回操作后的 `Flock` |
| `fcntl(command: usize, argument: usize)`、`ioctl(request: usize, buffer: MutSlice[byte])` | 原始机器字参数 fcntl 和同步缓冲区 ioctl；调用方提供请求 ABI |
| `read_dir()` | 读取一批 `Vec[DirEntry]`，EOF 时为空，并推进目录偏移 |

写入会将不可变输入视图复制到临时缓冲区。在输入复制和内核 I/O 期间均不允许并发修改这些视图。`read_exact` 和 `write_all` 返回错误时，可能已经完成部分操作。定位写入保留 Linux `O_APPEND` 行为，即可能无视指定偏移而追加。原始 `fcntl` 复制命令返回调用方拥有的原始描述符；`duplicate()` 是受管理的替代方案。`ioctl` 要求缓冲区足够容纳请求，不能注册返回后继续使用的内存。

相对于目录的自由函数包括 `stat_at`、`statx_at`、`mkdir_at`、`unlink_at`、`rename_at`、`link_at`、`symlink_at`、`readlink_at`、`chmod_at`、`chown_at` 和 `utimens_at`，每个都有 `_bytes` 版本。重命名使用 `renameat2`，支持 `RENAME_NOREPLACE`、`RENAME_EXCHANGE` 和 `RENAME_WHITEOUT`；删除目录使用 `AT_REMOVEDIR`。`utimens_at` 接受访问和修改时间的 `Timespec` 值，纳秒字段可使用 `UTIME_NOW`/`UTIME_OMIT`。`readlink_at` 返回原始目标字节。`DirEntry` 保留原始 `name: Vec[byte]`、`inode: u64`、`next_offset: i64` 和 Linux `kind: byte`；包含 `.` 和 `..`，并保留未知类型。目录顺序和偏移 cookie 由内核定义。

```gom
use std::os::linux::fd;

fn append_bytes(path: string, bytes: Slice[byte]) -> Result[(), fd::Error] {
    let file = fd::open(path, fd::O_WRONLY | fd::O_CREAT | fd::O_APPEND, 0x180)?;
    let written = file.write_all(bytes);
    let closed = file.close();
    written?;
    closed
}
```

这些接口仅支持 Linux amd64。它们使用普通导入、结构体、函数和方法，不增加原生指针或 `unsafe` 语法。

<a id="linux-amd64-processes"></a>

## Linux amd64 进程

`std::os::linux::process` 提供 Linux 进程操作，并重新导出 `io::Error`/`ErrorKind`。PID 参数必须能表示为有符号 32 位整数；信号编号范围为零到 64。信号零用于检查存在性/权限，负 PID 或零 PID 对 `kill` 和 `wait` 保留内核的进程组选择语义。导出的 `SIG*` 常量使用原生 Linux 编号，包括 32/33，而非 glibc 调整后的实时信号范围。

| 函数 | 结果或行为 |
| --- | --- |
| `pid()`、`parent_pid()`、`thread_id()` | `Result[isize, Error]` |
| `uid()`、`effective_uid()`、`gid()`、`effective_gid()` | `Result[u32, Error]` |
| `user_ids()`、`group_ids()`、`groups()` | 实际/有效/保存的 `ResIds`，或附加组 `Vec[u32]` |
| `process_group(pid)`、`session(pid)` | 组/会话 ID；零选择调用方 |
| `set_process_group(pid, group)`、`create_session()` | 应用 Linux 组/会话规则 |
| `kill(pid, signal)`、`signal_thread(group, thread, signal)` | 进程/组信号发送和 `tgkill` |
| `pidfd_open(pid, flags)`、`pidfd_signal(handle, signal)` | 拥有的 close-on-exec `fd::Fd`，以及通过稳定进程句柄发送信号 |
| `wait(pid, options)` | `Result[Option[WaitOutcome], Error]`；`WNOHANG` 找不到就绪子进程时为 `None` |
| `resource_usage(who)` | `abi::Rusage`；支持 `RUSAGE_SELF`、`RUSAGE_CHILDREN`、`RUSAGE_THREAD` |
| `resource_limit(pid, resource)`、`set_resource_limit(pid, resource, limit)` | 读取或原子替换 `abi::Rlimit`，替换时返回先前限制 |
| `priority(which, who)`、`set_priority(which, who, value)` | -20 到 19 的普通 nice 值，受内核权限检查约束 |

`WaitOutcome` 公开 `pid`、`status: WaitStatus` 和 `usage: abi::Rusage`。`WaitStatus { raw: u32 }` 提供 `exited`、`signaled`、`stopped`、`continued`、`core_dumped` 和 `success` 谓词，以及可选的 `exit_code`、`signal` 和 `stop_signal`。等待选项包括 `WNOHANG`、`WUNTRACED` 和 `WCONTINUED`。资源限制公开全部 16 个 `RLIMIT_*` 选择器和 `RLIM_INFINITY`；Linux 最大 RSS 的单位为 KiB。`thread_id`、`RUSAGE_THREAD` 以及 `who == 0` 时的 `PRIO_PROCESS` 指执行中的操作系统线程，该线程会随 GoML 任务调度而变化。它们不能标识或配置一个稳定的 GoML 任务。资源限制设置器直接更新内核。当当前软限制等于初始硬限制减一时，Go 创建子进程可能恢复其原始 `RLIMIT_NOFILE` 软限制，即使该值是显式原始更新后的结果；其他当前值正常继承。

`Command::new(program)` 构建异步子进程命令。链式方法包括 `arg(string)`、`args(Slice[string])`、`current_dir(string)`、`env(key, value)`、`env_clear()`、`stdin(fd)`、`stdout(fd)`、`stderr(fd)`、`extra_fd(fd)`、`new_session(bool)` 和 `new_process_group(bool)`。`spawn()` 返回 `Child`。内建运行时使用 Go 的[运行时协调 ForkExec](https://go.dev/src/syscall/exec_linux.go)；GoML 代码不发起原始 fork/clone。不含斜杠的程序名通过 Go `exec.LookPath` 使用父进程 PATH；含斜杠的路径在应用子进程工作目录后执行。命令环境改动应用于子进程。除非清空，环境在 spawn 时继承，后续覆盖优先。键必须非空且不含 NUL 或 `=`；值拒绝 NUL。空工作目录表示继承。

标准描述符 0/1/2 默认继承。显式描述符在整个 spawn 期间保持借用。额外描述符占用从 3 开始的连续子进程槽位。子进程获得独立的描述符表条目，指向相同的打开文件描述；父进程的所有者仍保持打开。新会话也创建新进程组；同时启用两个选项的效果与 `new_session(true)` 相同。Command 值通过普通 `Vec` 语义共享参数、环境和描述符列表；不要在启动命令时并发修改它。

Spawn 使用 Linux `CLONE_PIDFD` 请求原子 pidfd，需要相应内核能力。`Child.pid()` 返回数字 PID，而 `signal(signal)` 和 `kill()` 使用 pidfd 避免 PID 复用。`wait()` 只回收一次，并在别名之间缓存结果；并发等待方得到缓存结果。子进程仍在运行或其他等待方持有等待操作时，`try_wait()` 返回 `None`。`wait_event(options)` 还可通过支持的等待标志观察停止/继续事件。其他任务等待时仍可发送信号。终止后的回收会关闭内部 pidfd；后续信号操作报告描述符已关闭。成功 spawn 后应调用 `wait()`，包括调用 `kill()` 后；丢弃 Child 不会终止或回收子进程。终止后的回收由其 `Child` 负责，因此其他 wait API 和 SIGCHLD 自动回收不得接管此责任。

```gom
use std::os::linux::process;

fn child_status() -> Result[isize, process::Error] {
    let child = process::Command::new("/bin/sh").arg("-c").arg("exit 7").spawn()?;
    let result = child.wait()?;
    Ok(result.status.exit_code().unwrap_or(-1))
}
```

不提供凭据修改、信号处理器替换和信号掩码替换的便捷封装，因为这些操作需要与多线程运行时协调。这些进程 API 仅支持 Linux amd64，使用普通语言结构，不增加语法。

<a id="linux-amd64-memory-mappings"></a>

## Linux amd64 内存映射

`std::os::linux::memory::anonymous(length: isize, protection: usize, flags: usize)` 和 `file(descriptor: fd::Fd, length: isize, offset: i64, protection: usize, flags: usize)` 返回 `Result[Mapping, io::Error]`。本包重新导出 `Error` 和 `ErrorKind`。匿名映射添加 `MAP_ANONYMOUS`；调用方选择 `MAP_PRIVATE` 或 `MAP_SHARED`。文件映射在 `mmap` 期间借用描述符，描述符关闭后映射仍有效。`MAP_PRIVATE` 提供写时复制行为，而 `MAP_SHARED` 使底层文件修改共享可见。

长度必须为正，文件偏移必须为 `PAGE_SIZE` 的非负整数倍（Linux amd64 上为 4096）。复制访问限制在请求的长度内，即使最后一个映射页只使用了一部分也是如此。`Mapping` 拒绝固定地址、增长栈、显式 hugetlb 和未初始化映射标志。这些专门操作仍可通过原始系统调用访问。受管理的映射不公开原始地址、Go 切片视图、可执行函数指针、重映射或部分解除映射。

| `Mapping` 方法 | 行为 |
| --- | --- |
| `len()`、`is_closed()` | 请求的长度和共享生命周期状态 |
| `read_at(output: MutSlice[byte], offset: isize)`、`write_at(input: Slice[byte], offset: isize)` | 经检查的复制访问，返回实际字节数 |
| `read_exact_at(output, offset)`、`write_all_at(input, offset)` | 重复短传输，直到完成或出错 |
| `protect(protection)` | 使用 `PROT_NONE/READ/WRITE/EXEC/SEM` 改变整个映射的保护 |
| `sync(flags)`、`advise(advice)` | 整个映射的 `msync` 和 `madvise` |
| `lock(flags)`、`unlock()` | 整个映射的 `mlock`/`mlock2` 和 `munlock`；支持 `MLOCK_ONFAULT` |
| `residency()` | `mincore` 为每个基本页返回一个布尔值，描述某一时刻的快照 |
| `close()` | 在所有别名间显式、幂等地解除映射 |

复制通过 [process_vm_readv/process_vm_writev](https://man7.org/linux/man-pages/man2/process_vm_readv.2.html) 操作当前进程，本地 Go 缓冲区只在调用期间固定。它有系统调用/分配开销，在不受支持的内核或系统调用限制下可能返回 `ENOSYS`/`EPERM`。不可访问或文件截断后的页面返回短复制或 `EFAULT`，不会从 Go 直接解引用映射存储。错误保留内核 errno。精确操作出错前可能已修改前缀；写入先复制不可变输入。调用期间不得并发修改调用方缓冲区。

别名共享一个通道，用于串行化复制、保护变更、其他操作和解除映射。`close()` 等待活动访问完成，只解除映射一次；解除失败时保留映射以便重试。关闭后的其他操作返回 `EBADF`；`len()` 保留请求长度。没有终结器。读取要求显式 `PROT_READ`，写入要求显式 `PROT_WRITE`，即使 amd64 硬件隐含额外权限也是如此。缺少声明的访问权限返回 `EACCES`。独立映射和其他进程需要自行同步；映射字节不是原子 IPC 协议。

```gom
use std::os::linux::memory;

fn mapped_bytes() -> Result[Vec[byte], memory::Error] {
    let region = memory::anonymous(4096, memory::PROT_READ | memory::PROT_WRITE, memory::MAP_PRIVATE)?;
    defer { let _ = region.close(); };
    region.write_all_at(b"hello".as_slice(), 0)?;
    let output = b"?????";
    region.read_exact_at(output.as_mut_slice(), 0)?;
    Ok(output)
}
```

本包公开 Linux 的保护、映射、同步、建议和锁定常量；常量存在不保证内核支持。操作遵循内核权限和资源限制。本包仅支持 Linux amd64，不引入指针或布局语法。

<a id="linux-amd64-ipc-and-readiness"></a>

## Linux amd64 IPC 与就绪通知

`std::os::linux::ipc` 使用拥有的 `fd::Fd` 值和结构化的 `io::Error` 结果。`pipe(flags)` 返回 `Pipe { reader, writer }`；`socket_pair(kind, flags)` 使用 Unix 域返回 `SocketPair { first, second }`。套接字类型包括 `SOCK_STREAM`、`SOCK_DGRAM` 和 `SOCK_SEQPACKET`。标志支持非阻塞操作；新建描述符始终具有 close-on-exec。`shutdown(socket, direction)` 支持 `SHUT_RD`、`SHUT_WR` 和 `SHUT_RDWR`。必须显式关闭每个返回的描述符。

`send_message(socket, data: Slice[byte], descriptors: Slice[fd::Fd], flags)` 发送一条消息，可附带 `SCM_RIGHTS` 描述符，返回传输字节数，并添加 `MSG_NOSIGNAL`。它复制载荷输入，并在调用期间借用所有描述符。最多可发送 253 个描述符，传输描述符至少需要一个载荷字节。部分完成的流发送可能已经传输了权限；重新发送整个带描述符的消息可能在接收方产生重复描述符。

`receive_message(socket, output: MutSlice[byte], max_descriptors: isize, flags)` 返回 `ReceivedMessage { byte_count: usize, copied: isize, flags: usize, descriptors: Vec[fd::Fd] }`。它通过 `MSG_CMSG_CLOEXEC` 为接收的描述符原子设置 close-on-exec。数量上限为 0–253。`MSG_TRUNC` 可能使 `byte_count` 大于 `copied`；零字节数据报有效。控制数据截断或描述符数量超限时，接收到的描述符会被关闭，并在接收后返回错误；除非请求了 `MSG_PEEK`，载荷已被消费。成功时，调用方拥有每个返回的描述符。调用期间不得并发修改不可变发送输入或可变接收缓冲区。

`eventfd(initial: u32, flags)` 返回计数器描述符；`EFD_NONBLOCK` 和 `EFD_SEMAPHORE` 控制其行为。`read_counter(fd)` 和 `write_counter(fd, value: u64)` 执行精确的八字节计数器操作。`timerfd(clock, flags)` 创建定时器描述符；`set_timer(timer, flags, value: abi::Itimerspec)` 返回先前设置，`get_timer(timer)` 读取设置。零值解除定时器，间隔使其周期性触发，`TFD_TIMER_ABSTIME` 选择绝对到期时间。通过 `read_counter` 读取到期次数；非阻塞空计数器报告 `WouldBlock`。

`poll(interests: Slice[PollInterest], timeout_ms: i32)` 借用所有描述符并返回 `Vec[PollEvent]`。`PollInterest { descriptor, events: i16 }` 指定 `POLL*` 掩码；`PollEvent { index, events }` 标识就绪输入项。超时 -1 表示无限等待，零表示立即轮询，正值以毫秒计。Poll 返回 `EINTR`，不会重新开始有限超时。

`Epoll::new()` 创建必须显式关闭的共享事件反应器句柄。`add(descriptor, events: u32, token: u64)` 复制并保留被观察的打开文件描述，因此调用方可以关闭原始所有者而不使观察失效。Token 必须在当前观察项中唯一，可使用任意 `u64` 值。`modify(token, events)` 更新或重新启用观察项；`remove(token)` 注销观察项并关闭其保留的描述符。`wait(max_events: isize, timeout_ms: i32)` 返回 `Vec[abi::EpollEvent]`，每个 `data` 字段保存指定 token。它支持 `EPOLL*` 掩码，包括边缘触发和单次触发模式；调用方必须遵循内核的排空/重新启用规则。等待被中断时返回 `EINTR`。

每个 epoll 注册使用永不复用的内部身份；已移除注册的事件会被丢弃，因此并发移除可能产生空结果。等待期间可以添加、修改或移除观察项。`close()` 在释放描述符之前通过私有 eventfd 唤醒所有阻塞等待方；这些等待方返回 `EBADF`。关闭在别名之间是幂等的，并关闭所有保留的观察项。因此，可从另一个任务关闭正在无限等待的空 epoll 实例。

共享内存 IPC 组合使用 `fd::memfd`、`memory::file(..., MAP_SHARED)` 和描述符传输。每个进程分别拥有自己的描述符和映射，并为共享内容提供同步。管道、套接字、计数器、定时器、poll 和 epoll 仅支持 Linux amd64；这些封装不增加文法。

<a id="url-component-and-query-codecs"></a>

## URL 组件与查询编解码器

`std::net::url` 提供纯粹保留字节的组件/查询编解码器、无损引用拆解/解析、显式 authority 验证和 ASCII 序列化。它依赖 `std::bytes`，不依赖套接字、DNS、TLS 或 HTTP 包。原始引用类型不是经过验证的网络端点。生态包 `request::Url` 复用查询编解码器和绝对引用拆解，同时保留自身的 HTTP 端点、UTF-8、片段和相对路径合并策略。

`escape(input: Slice[byte], mode: EscapeMode, limit: isize) -> Result[string, CodecError]` 接受任意字节，输出带大写百分号转义的 ASCII。`EscapeMode::PathSegment` 转义斜杠和其他路径段分隔符，同时保留路径段安全的保留字节 `$&+:=@`；它不是整个路径的编码器。`QueryComponent` 转义所有保留字节，并将空格编码为 `+`。两者都保留 ASCII 字母、数字和 `-._~`。`path_escape(string, limit)` 和 `query_escape(string, limit)` 是针对 UTF-8 字符串的便捷函数。limit 限制编码后的输出字节数，在构造输出缓冲区之前检查。

`unescape(input: string, mode, limit) -> Result[bytes::Bytes, CodecError]`、`path_unescape` 和 `query_unescape` 接受大小写十六进制转义，每个 `%HH` 恰好解码一次。查询模式将 `+` 转为空格；路径模式保留 `+`。解码字节不必是有效 UTF-8：NUL 和 `%FF` 会被保留，而非拒绝或替换。需要文本的调用方应使用经检查的 `utf8::BytesUtf8` 转换。每次成功解码都拥有独立可变存储。limit 限制解码后字节数，而非编码输入长度；处理时间与输入长度成线性关系，当前实现会复制其 UTF-8 字节。解码不验证 UTF-8、主机语法或保留字符策略。

`parse_query(input, max_fields, max_decoded_bytes)` 返回 `Result[Vec[(bytes::Bytes, bytes::Bytes)], CodecError]`。应传入查询本身，本函数不会删除或解释其中的边界分隔符：开头的 `?` 或 `#` 是字面查询数据，不是 URL 边界。解析按 `&` 划分非空字段，再按每个字段的第一个 `=` 分割，缺少 `=` 时值为空。空字段被忽略；空名称/值和重复名称按输入顺序保留。原始分号会被拒绝，`%3B` 则是有效数据。名称和值使用查询反转义，分别放入独立拥有的缓冲区。两个上限分别限制非空字段数和名称/值解码后的总字节数。任何失败都返回错误，不像 Go ParseQuery 那样返回部分映射加错误；已解析字段不会公开。

`parse_query_bounded(input, max_input_bytes, max_fields, max_decoded_bytes)` 还会在复制或扫描前检查原始输入字节长度。它先拒绝任何负上限，原始预算超限时在首个被排除的输入字节处报告 `InputLimit`。这可防止被忽略的分隔符绕过资源限制。旧的三参数 `parse_query` 保留解码/字段限制，但要求调用方单独限制原始输入；两种 API 都不能替代完整的 HTTP 请求大小策略。

`encode_query(fields: Slice[(bytes::Bytes, bytes::Bytes)], limit)` 返回 ASCII `Result[string, CodecError]`。它复制键值对序列，按解码后的名称字节进行字典序排序，相同名称的值保留相对顺序，然后将每对输出为 `name=value`，以 `&` 连接。输入顺序和缓冲区保持不变。排序或构造输出之前进行大小预检，通过检查剩余容量避免长度运算溢出。limit 限制编码后的总字节数；输出之外的排序工作空间为 O(fields)。两遍编码期间都不得并发修改输入视图或共享名称/值缓冲区；这些函数不提供同步。空键值对序列编码为空字符串。

所有上限必须非负，包括空输入的情况。`CodecError::kind()` 返回 `CodecErrorKind::{InvalidLimit, InputLimit, OutputLimit, InvalidEscape, FieldLimit, InvalidSeparator}`。对于组件编解码器和查询解析，`index()` 标识输入字节；查询编码则报告原始键值对索引。无效上限错误使用索引零。格式错误的百分号转义指向 `%`；查询解码错误使用完整查询中的偏移。错误类型支持相等比较和调试格式化，错误支持调试/显示格式化。这些都是普通枚举、结构体和函数，不改变文法或隐式转换。

```goml
use std::net::url;

fn canonical_query(input: string) -> Result[string, url::CodecError] {
    let fields = url::parse_query_bounded(input, 16384, 100, 4096)?;
    url::encode_query(fields.as_slice(), 16384)
}
```

例如，`canonical_query("b=a+b&a=%FF&b=%2B")` 返回 `Ok("a=%FF&b=a+b&b=%2B")`。百分号解码不等于安全净化；应用代码必须根据目标上下文验证解码后的组件。

<a id="raw-url-references-and-relative-resolution"></a>

## 原始 URL 引用与相对解析

`url::Reference::parse(input, max_input_bytes) -> Result[Reference, ParseError]` 将引用拆解为不可变原始组件。在复制字节之前检查输入字节上限。解析器拒绝字面的 ASCII 控制字符（含 DEL）、第一个路径斜杠之前无效的 scheme 语法，以及路径或片段中格式错误的百分号转义。它保留 scheme 大小写、百分号拼写、Unicode、空组件和重复斜杠，不在分割前解码分隔符。`scheme()`、`raw_authority()`、`raw_query()` 和 `raw_fragment()` 返回 `Option[string]`；`raw_path()` 返回字符串，`is_absolute()` 检查是否存在 scheme。`None` 不同于 `Some("")`：`?`、`#` 和 `//` 保留其显式空组件标记。查询转义在查询解码前不作解释。authority 在通过 `authority()` 显式验证之前只是原始子字符串；因此无效主机和端口也可能通过原始引用解析。原始空格、Unicode 和反斜杠不会自动转换为规范的传输 URI，因此解析成功不意味着引用可以用作 HTTP 端点或请求目标。

`decoded_path(limit)` 返回独立的 `bytes::Bytes`；`decoded_fragment(limit)` 在 Result 中返回 `Option[bytes::Bytes]`，保留不存在的状态。两者都使用路径式百分号解码，保留字面加号和非 UTF-8 解码字节。即使片段不存在，负上限也会失败。`ToString` 在解析后无损重组原始组件；它不规范化 scheme/主机、不隐去凭据，也不转义原始字符。记录引用日志前应谨慎。`PartialEq`/`Eq` 比较精确原始组件值，不比较 URI 等价性、DNS 身份或解码字节。

`base.resolve(reference, max_output_bytes)` 返回 `Result[Reference, ParseError]`。它应用 [RFC 3986 第 5 节](https://www.rfc-editor.org/rfc/rfc3986.html#section-5.2) 的组件继承、路径合并和字面点路径段处理规则：具有自身 scheme 的引用替换基准；否则继承基准 scheme，显式 authority 则替换基准 authority。空路径继承基准路径，且仅在没有提供查询标记时继承其查询。非空相对路径与基准目录合并。结果片段始终来自引用，而非基准。相对引用要求绝对基准；独立的绝对引用可以替换相对基准。无根 scheme 路径使用通用 URI 规则，而不是 Go 的 scheme 专属不透明路径行为。

只移除完整的字面 `.`/`..` 路径段。转义的点和斜杠保持转义，重复分隔符被保留，查询/片段文本不进行路径规范化。若结果没有 authority 且规范化路径以 `//` 开头，存储的路径会加上 `/.` 前缀，防止重新解析时凭空产生 authority。这保留了预期的路径解释，而不是直接输出有歧义的 `scheme://...` 文本。解析后的输入本身从不被规范化。空查询和片段标记在解析中保留，不同于将不存在和空值合并的 API。与 Go 处理空引用/主机的方式不同，缺失的引用片段从不继承基准片段，显式空 authority 会替换基准 authority 而非继承它。

在根处产生的重复斜杠也被保留：相对 `https://a` 解析 `..//` 得到 `https://a//`，相对 `file:///a/b` 则得到 `file:////`。Go 1.26 在这两种情况下都少生成一个斜杠；测试记录这一有意差异，而不将 Go 输出视为唯一规范。

解析检查最终重组的字节长度，但不构造该输出字符串；负输出上限优先失败。临时路径工作量与基准/引用路径大小成线性关系，不仅受最终输出上限约束。应限制两份解析输入，以限制中间工作量。返回值只保留不可变字符串，不修改任何输入。`ParseError::kind()` 返回 `ParseErrorKind::{InvalidLimit, InputLimit, OutputLimit, RelativeBase, ControlCharacter, InvalidScheme, InvalidEscape}`；类型支持相等比较/调试格式化，错误支持调试/显示。语法错误的 `index()` 为源字节偏移，输入/输出超限时为被排除的字节边界，无效上限或相对基准错误时为零。这些 API 使用普通结构体、枚举、方法和现有字符串/字节原语，不引入语法、反射或运行时后端。

```goml
use std::net::url;

fn resolve_link(base: string, link: string) -> Result[string, url::ParseError] {
    let base = url::Reference::parse(base, 4096)?;
    let reference = url::Reference::parse(link, 4096)?;
    Ok(base.resolve(reference, 8192)?.to_string())
}
```

例如，相对 `https://example.com/a/b?old#base` 解析 `../d?` 得到 `https://example.com/d?`。这种通用引用解析独立于生态包 request 现有的 HTTP 相对路径合并兼容规则；使用该值层不要求替换应用策略。

<a id="ascii-url-serialization"></a>

## ASCII URL 序列化

`Reference::to_ascii(max_output_bytes) -> Result[string, ParseError]` 显式验证 authority 并输出 ASCII URI 引用。它保留组件边界、scheme/主机拼写、不存在与空标记的区别、点路径段、端口拼写和凭据。它将现有百分号转义中的十六进制数字转为大写而不解码；转义的斜杠、点和无效 UTF-8 字节保留原解释。其他不允许的原始字节，包括空格、反斜杠和 Unicode UTF-8 的每个字节，转换为大写 `%HH`。scheme 语法和分隔符保持不变。路径保留非保留字符、子分隔符、`:`、`@` 和 `/`；查询和片段还保留 `?`。authority 验证后，其方括号和分隔符仍保有结构意义。

此操作在构造输出缓冲区之前检查完整编码长度，在被排除的输出字节边界返回 OutputLimit，并优先拒绝负上限。临时扫描存储受已解析输入大小约束，不仅受输出大小约束。原始查询不会进行表单解码、重排或名称/值对验证，但无效百分号序列会在完整原始引用中的对应字节偏移处以 InvalidEscape 拒绝。authority 验证错误保留下文所述的 authority 相对偏移。`ToString` 仍是无损选项，包括原始查询文本不是有效传输转义序列的情况。

例如，`/café?q=a b+#é` 序列化为 `/caf%C3%A9?q=a%20b+#%C3%A9`；字面加号不会转为空格。重新解析后的序列化是幂等的。这提供规范的转义拼写，不是通用的 URL 等价性规范化：它不移除默认端口、不将主机转为小写、不解析点路径段、不解码非保留字符转义，也不执行 IDNA。百分号编码的 Unicode 注册名称是通用 URI 语法，不是可直接用于 DNS 的主机名。结果不会自动成为 HTTP 请求目标或安全日志文本，其中的凭据也未隐去。应另外应用策略并显式脱敏。此方法不增加文法或隐式转换。

<a id="url-authorities-and-credential-redaction"></a>

## URL authority 与凭据脱敏

`url::Authority::parse(input, max_input_bytes)` 验证不含开头 `//` 的独立 authority，返回 `Result[Authority, ParseError]`。在复制或扫描之前检查输入预算。`Reference::authority()` 显式验证其原始 authority，返回 `Result[Option[Authority], ParseError]`；引用解析本身仍是无损拆解，不是 authority 验证。authority 错误偏移相对于包含用户信息的 authority 子字符串，即使通过 Reference 访问也是如此。

`raw_host()` 保留 IP 字面值外的方括号；`host_kind()` 返回 `HostKind::{Name, Ipv4, Ipv6, IpvFuture}`。带方括号的 IPv6 支持压缩和末尾 IPv4 形式；IPv4 分量禁止冗余前导零。非 IP 注册名称不会作为 DNS 名称验证。IPvFuture 使用 `v` 加十六进制版本、点以及非空的允许地址拼写。`%25` 引入非空 IPv6 zone；其原始字符为非保留 ASCII 或 Unicode，百分号转义不得解码为 ASCII 控制字符。这是文本 zone 值，不进行接口查找。`raw_zone()` 保留不含分隔符的 zone 拼写。`hostname(limit)` 移除方括号，并百分号解码到新的 Bytes，包括解码后的 `%zone` 后缀；`zone(limit)` 返回可选的新 Bytes。

`raw_port()` 区分没有分隔符和显式空端口。非空端口只能包含 ASCII 十进制数字。`port_number()` 单独返回 `Result[Option[u16], ParseError]`，拒绝大于 65535 的值而不回绕；不存在和空端口都返回 None。允许端口零和前导零。仅语法解析不会施加传输端口范围限制。

`userinfo()` 返回可选 UserInfo。`raw_username()`、`raw_password()` 和 `has_password()` 保留密码不存在与为空的区别。`username(limit)` 和 `password(limit)` 解码到独立 Bytes，保留字面加号和任意百分号编码字节。只允许一个原始 `@`；嵌入的 `@` 必须百分号编码，不同于 Go 宽松的最后一个 `@` 分割规则。原始用户信息只能是 ASCII 非保留字符、子分隔符和冒号文本；Unicode 必须转义。注册名称允许 Unicode 和通用百分号编码字节，不同于 Go 对百分号编码 ASCII 主机的限制。IPvFuture 的接受独立于 Go 的主机策略。这些操作均不执行 DNS、IDNA、HTTP 策略或端点验证。

Authority 和 UserInfo 实现原始值相等比较和无损 ToString，但有意不实现 Debug。`Authority::redacted(limit)` 和 `Reference::redacted(limit)` 将存在的密码（包括空密码）替换为 `xxxxx`，保留其他拼写。例如，`https://u:p@host/a?q=secret` 变为 `https://u:xxxxx@host/a?q=secret`。查询、路径、片段和用户名中的秘密不会移除；这不是通用的安全日志净化器。Reference 脱敏先验证 authority。负上限会失败，拼接结果前检查输出字节上限。解码上限限制解码字节数；还应单独限制原始解析输入。

额外的 ParseErrorKind 变体包括 `InvalidHost`、`InvalidPort`、`PortOutOfRange`、`InvalidUserInfo`、`InvalidZone` 和 `InvalidIpLiteral`；格式错误的百分号序列使用 `InvalidEscape`。诊断报告错误类型和字节偏移，不回显凭据。

<a id="networking"></a>

## 网络

`std::net` 使用 Linux amd64 套接字系统调用，在 GoML 中实现 IPv4/IPv6 TCP 和 UDP，不将网络功能委托给 Go 的 `net` 包。所有套接字均为非阻塞且具有 close-on-exec；惰性共享 epoll 工作协程通过通道唤醒等待中的 GoML 任务。一个 epoll 描述符和一个工作协程在进程整个生命周期中保持存在，与套接字数量无关。实现使用单次触发、水平触发就绪模式，以及防止排队事件指向复用描述符的注册身份。原生布局和就绪处理遵循 [socket](https://man7.org/linux/man-pages/man2/socket.2.html)、[connect](https://man7.org/linux/man-pages/man2/connect.2.html)、[accept](https://man7.org/linux/man-pages/man2/accept.2.html) 和 [epoll_ctl](https://man7.org/linux/man-pages/man2/epoll_ctl.2.html)。

`IpAddr` 为 `V4([byte; 4])` 或 `V6([u16; 8])`。`IpAddr::parse(string)` 接受数字地址，包括压缩 IPv6 和 IPv6 中的 IPv4 尾部。IPv4 八位组使用无前导零的十进制。`SocketAddr` 公开 `ip: IpAddr`、`port: u16` 和 `scope_id: u32`；`new(ip, port)` 将 scope 设为零。`SocketAddr::parse` 接受 `127.0.0.1:8080`、`[::1]:8080` 或 `[fe80::1%3]:8080`。IPv6 scope ID 必须是数字；不支持接口名和主机名。IPv4 的 scope 必须为零。两种地址类型均支持相等比较、调试输出和 `to_string()`；IPv6 输出使用小写十六进制，压缩第一个最长的连续零段。绑定时端口零请求内核分配端口；通过 `local_addr()` 读取它。

`IpAddr::bit_len()` 返回 32 或 128。`is_ipv4_mapped()` 只识别 `::ffff:0:0/96`；`unmap()` 将这些地址转换为 V4，其余保持不变。两种操作都不改变原值。分类方法有 `is_unspecified`、`is_loopback`、`is_private`、`is_multicast`、`is_link_local_unicast`、`is_link_local_multicast`、`is_interface_local_multicast` 和 `is_global_unicast`，均返回 bool。私有范围为 IPv4 RFC 1918 和 IPv6 `fc00::/7`。映射地址按嵌入的 IPv4 地址分类，但 `is_unspecified` 只检查字面的 `0.0.0.0` 或 `::`，`is_interface_local_multicast` 仅适用于 IPv6。因此 `::ffff:0.0.0.0` 本身不是未指定地址，但其 unmap 后的值是。全局单播分类排除未指定、环回、多播、链路本地单播和 IPv4 全一广播地址；它包含私有和文档范围，不能作为公网可达性或安全策略检查。这些值操作不使用 DNS、套接字或平台调用，不增加语法。

IpAddr 实现 `std::cmp::{Ord, PartialOrd}`：IPv4 排在 IPv6 前，之后按无符号网络序地址字节排序。映射 IPv6 在排序和相等比较中仍是 IPv6。`next()` 和 `prev()` 返回 Option[IpAddr]，保留地址族，在其数字最大/最小值处返回 None，不回绕也不产生无效地址。它们不修改原值。

`as4()` 对 IPv4 和 IPv4 映射 IPv6 返回 Option[[byte; 4]]，对其他 IPv6 返回 None。`as16()` 返回网络序 [byte; 16]，将 IPv4 映射到 `::ffff:0:0/96`。`IpAddr::from16([byte; 16])` 始终构造 IPv6，包括映射 IPv6；需要 IPv4 时应显式使用 `.unmap()`。`IpAddr::from_slice` 接受恰好 4 或 16 字节的 Slice[byte]，返回 Result[IpAddr, Error]，并复制输入。`to_bytes()` 返回地址族自身宽度（4 或 16 字节）的新 Vec[byte]。返回的数组/向量不与地址存储互为别名。固定宽度 IPv4 可直接用 `IpAddr::V4([byte; 4])` 构造。

IpAddr、ScopedIpAddr、IpPrefix 和 SocketAddr 实现 prelude Hash trait，可用作 HashMap 键。哈希包含相等比较使用的每个字段：地址族、精确 zone 文本、存储的前缀地址/长度以及套接字端口/scope。它不会对 IPv6 做 unmap、屏蔽前缀主机位或规范化 zone。解析后相等的拼写具有相等哈希；不同值即使哈希碰撞也仍是不同键。哈希输出是集合内部实现细节，不是加密摘要或稳定的传输/存储格式。若应用需要这种键等价性，应在插入前显式使用 `unmap()` 或 `masked()` 规范化。

`ScopedIpAddr` 将现有 IpAddr 与不透明、区分大小写的 zone 字符串组合，不改变 V4/V6 枚举或套接字布局。`new(address, zone)` 和 `IpAddr::with_zone(zone)` 返回 Result：非空 zone 要求 IPv6，包括 IPv4 映射 IPv6。不同于 Go WithZone，为 IPv4 提供非空 zone 会报错，而非静默丢弃。空 zone 表示不存在。`ScopedIpAddr::parse(string)` 接受普通 IP 或 `IPv6%zone`；显式空后缀无效。首个 `%` 分隔 zone，其余内容按字面保留，包括额外的百分号和 Unicode。不进行 URL 解码、接口名验证或接口查找。

`address()` 和 `without_zone()` 返回底层 IpAddr；`zone()` 返回字符串。`with_zone(string)` 返回经检查的替换值，不修改原值。`unmap()` 在转换为 IPv4 时移除映射 IPv6 前缀及其 zone，其余地址和 zone 保持不变。相等比较检查地址和精确 zone；Debug/to_string 仅在 zone 存在时附加 `%zone`。例如，`ScopedIpAddr::parse("fe80::1%eth0")?.zone()` 为 `"eth0"`。ScopedIpAddr 还实现 Ord/PartialOrd，先比较 IpAddr，再按字节顺序比较 zone 文本（不存在排在非空前）。其 `next()`/`prev()` 返回 Option[ScopedIpAddr] 并保留 zone，即使跨越映射 IPv6 范围边界也是如此。没有操作会隐式改变地址族。

`to_socket_addr(port: u16)` 接受不存在的 zone 或十进制 u32 zone，返回 Result[SocketAddr, Error]；具名、带符号和溢出的 zone 会失败，不执行网络操作。调用方必须显式解析接口名。`ScopedIpAddr::from_socket_addr(socket)` 验证套接字地址/scope 组合，并将非零 scope 转为十进制。scope 零变为不存在的 zone，因此数字转换有意将 `"0"` 规范化为不存在、将 `"003"` 规范化为 `"3"`。普通带 zone 值的解析/格式化保留这些拼写。前缀仍只接受不带 zone 的 IpAddr 值；因此，为前缀匹配丢弃 zone 时必须显式调用 `address()` 或 `without_zone()`。

`IpPrefix::new(address: IpAddr, bits: isize)` 验证前缀长度：IPv4 为 0..32，IPv6 为 0..128，返回使用现有网络 Error 类型的 Result。`IpPrefix::parse(string)` 接受数字 `address/length` 表示；长度必须是无前导零的无符号十进制（`0` 除外）。拒绝 zone、主机名、缺失长度和超范围长度。两种构造方式都保留主机位：`192.0.2.129/24` 保持原值，直到 `.masked()` 返回 `192.0.2.0/24`。`.address()` 和 `.bits()` 查看原值；`.is_single_ip()` 识别全位宽前缀。相等比较检查地址加前缀长度，而非仅检查网络覆盖范围。Debug/to_string 复用现有 IpAddr 表示。

`.contains(address)` 比较网络位，忽略前缀的主机位。`.overlaps(other)` 检查两个前缀是否有共同地址。两者都要求相同地址族：IPv4 映射 IPv6 仍是 IPv6，不会隐式匹配 IPv4 前缀。零长度前缀包含其地址族中的所有地址。前缀字段为私有，因此公开构造无法产生无效长度。所有操作都是纯值计算，不访问 DNS/套接字/系统调用，也不增加语法。前缀不携带 zone；数字套接字 scope 仍保留在 SocketAddr 上。

可失败的套接字操作返回 `Result[T, net::Error]`；`Error` 和 `ErrorKind` 重新导出 `std::io` 类型。错误保留操作名和可用的数字 errno。取消为 `Interrupted`，超时为 `TimedOut`，关闭后使用为 `InvalidInput`，`read_exact` 中过早的 TCP EOF 为 `UnexpectedEof`。`io::ErrorKind` 未表示的连接错误（如连接被拒绝或重置）使用带 `raw_os_code()` 的 `Other`。地址解析失败为 `InvalidInput`。

| 类型与操作 | 返回值与行为 |
| --- | --- |
| `TcpListener::bind(address)` | `TcpListener`；启用地址复用，backlog 为 128 |
| `TcpListener::bind_with_backlog(address, backlog: isize)` | `TcpListener`；backlog 必须为可表示为有符号 32 位的正值 |
| `TcpListener.accept()` | `(TcpStream, SocketAddr)`；等待一个传入连接并返回对端地址 |
| `TcpStream::connect(address)` | `TcpStream`；等待连接建立，就绪后检查套接字错误 |
| `TcpStream.read(buffer: MutSlice[byte])` | `isize` 字节数；可能少于请求量；非空缓冲区返回零表示 EOF |
| `TcpStream.write(buffer: Slice[byte])` | `isize` 字节数；可能少于请求量 |
| `TcpStream.read_exact(buffer)` / `write_all(buffer)` | `()`；循环直到完成、EOF 或出错 |
| `TcpStream.shutdown(Shutdown::Read / Write / Both)` | `()`；关闭指定方向而不释放描述符 |
| `TcpStream.set_nodelay(bool)` | `()`；启用或禁用 TCP_NODELAY |
| `TcpStream.peer_addr()` | `SocketAddr` |
| `UdpSocket::bind(address)` | `UdpSocket` |
| `UdpSocket.send_to(data: Slice[byte], destination: SocketAddr)` | `isize`；发送一个数据报，包括空数据报 |
| `UdpSocket.recv_from(buffer: MutSlice[byte])` | `Datagram`，公开字段为 `source`、复制的 `len`、完整 `datagram_len` 和 `truncated` |
| 三种套接字类型都有：`local_addr()`、`close()`、`is_closed()` | 分别为 `SocketAddr`、`()` 和普通 `bool` |

UDP 接收恰好消费一个数据报，即使缓冲区为空也是如此。多余字节被丢弃；`datagram_len` 和 `truncated` 按 Linux [MSG_TRUNC 语义](https://man7.org/linux/man-pages/man2/recv.2.html) 报告情况。零长度 UDP 接收是有效数据报，不是 EOF。TCP 零长度读写在打开的套接字上完成，不传输数据。发送使用 `MSG_NOSIGNAL`，使断开的连接表现为错误。发送方法将不可变输入复制到系统调用缓冲区；调用方不得并发修改输入或接收缓冲区。

`WaitOptions::new()` 无限等待。`.with_timeout(time::Duration)` 和 `.with_cancel(task::CancelToken)` 返回修改后的选项。`accept_with(options)`、`connect_with(address, options)`、`read_with(buffer, options)`、`read_exact_with(buffer, options)`、`write_with(buffer, options)`、`write_all_with(buffer, options)`、`send_to_with(data, destination, options)` 和 `recv_from_with(buffer, options)` 应用这些设置。超时覆盖整个操作，包括排在其他读写方后等待以及所有部分传输；取得进展后不会重新计时。零超时不尝试 I/O，直接返回 `TimedOut`。已完成的 I/O 可能在与取消或超时的竞争中获胜。`read_exact` 或 `write_all` 失败后可能已经传输部分字节；缓冲区不会回滚，错误也不报告部分计数。需要跟踪进度时，应使用单次 `read`/`write` 调用。

套接字值是共享句柄。同一套接字上的多个读或多个写会串行执行；一个读方和一个写方可以并发推进。并发 `close()` 唤醒活动和排队操作，只释放套接字一次，且是幂等的。绑定和连接失败时释放描述符。始终显式或通过 `defer` 关闭套接字；垃圾回收不会关闭它们。IPv6 套接字仅支持 IPv6，因此同时需要两种协议时应分别绑定 IPv4 和 IPv6 监听器。Unix 域监听器、HTTP 和其他平台不在此套接字 API 范围内。DNS、TLS 客户端和 TLS 监听器由以下 API 提供。它使用现有导入、枚举、方法、通道和任务，不增加文法或编译期网络功能。

```goml
use std::net;
use std::time;

fn echo_once() -> Result[(), net::Error] {
    let address = net::SocketAddr::parse("127.0.0.1:8080")?;
    let listener = net::TcpListener::bind(address)?;
    defer { let _ = listener.close(); };
    let wait = net::WaitOptions::new().with_timeout(time::Duration::from_seconds(10));
    let (stream, _) = listener.accept_with(wait)?;
    defer { let _ = stream.close(); };
    let buffer = Vec::from_array([0, 0, 0, 0]);
    let count = stream.read_with(buffer.as_mut_slice(), wait)?;
    stream.write_all_with(buffer.slice(0, count), wait)
}
```

<a id="dns-and-tls"></a>

## DNS 与 TLS

`net::resolve(host, port)` 和 `resolve_with(host, port, context)` 使用宿主解析器返回去重的数字 `SocketAddr` 值；数字输入绕过 DNS。空名称或含 NUL 的名称以 `InvalidInput` 失败。解析遵循取消和截止时间；上下文无截止时间时最多 30 秒。`TcpStream::connect_host(host, port, context)` 解析一次，在同一上下文下按解析器顺序尝试地址。应提供截止时间以限制整个连接尝试。

`WaitOptions::with_context(context)` 为 TCP/UDP 操作添加父级取消和绝对截止时间。现有 `with_cancel` token 也继续生效；有效超时取上下文截止时间和 `with_timeout` 中较早者。上下文取消报告 `Interrupted`，到期报告 `TimedOut`，TCP 等待超时后套接字仍可用。

`std::net::tls` 提供 `connect(host, port, config)` 和 `connect_with(host, port, config, context)`。`ClientConfig::new(server_name)` 使用系统根证书验证证书链和服务器名称，要求 TLS 1.2 或更新版本，并将 DNS、拨号和握手的总时间限制为 30 秒。主机为不带端口或 IPv6 方括号的名称或数字地址。配置构建器包括 `with_ca_pem`（非空 PEM 包替换系统根证书；空包选择系统根证书）、`with_client_certificate(certificate_pem, private_key_pem)`、`with_alpn`、`with_minimum_version(Version::Tls12 / Tls13)` 和 `with_connect_timeout`。没有不安全的跳过验证开关。客户端证书和私钥必须同时提供；ALPN 标识为 1–255 字节。

`ServerConfig::new(certificate_pem, private_key_pem)` 使用证书和匹配私钥配置 TLS 服务器，支持 TLS 1.2–1.3，握手超时为 30 秒。`with_client_ca_pem(bundle)` 要求提供可由该非空证书包中的根证书验证的客户端证书。`with_alpn`、`with_versions(minimum, maximum)` 和 `with_handshake_timeout` 配置协商和每个连接的正值握手时限。空或无效的证书/私钥/客户端根证书、反向版本范围、空或过大的 ALPN 标识会产生可恢复错误。

`TlsListener::bind(address: net::SocketAddr, config)` 创建监听器；端口零选择可用端口，通过 `local_addr()` 返回。`accept()` 等待连接并完成 TLS 握手后返回 `TlsStream`。`accept_with(context)` 通过上下文限制等待和握手；配置的握手超时从接受 TCP 连接后开始。Accept 操作在每个监听器上串行执行。握手取消或失败会关闭该连接，监听器仍可用。`close()` 是共享且幂等的，唤醒排队/活动 accept 并取消进行中的握手；已接受的流仍独立拥有，必须分别关闭。本地监听器关闭报告 `BrokenPipe`；上下文取消和截止时间到期分别报告 `Interrupted` 和 `TimedOut`。TLS 服务器与 TLS 客户端使用相同的编译器运行时边界，不需要用户 Go FFI 绑定或使用方的 `go.mod`。

`accept_pending()` 和 `accept_pending_with(context)` 接受 TCP，但不执行 TLS 握手。它们返回所有权独立于监听器的流，因此关闭监听器不会关闭已经返回的待握手流。为每个流在独立任务中调用 `handshake()` 或 `handshake_with(context)`，即可并发协商；配置的握手超时从协商开始时计时。读写也会在传输应用字节前完成待进行的握手。并发握手调用共享一次协商；排队调用遵循取消和流关闭。协商开始前取消会保留流的可用性；活动协商失败或取消则关闭流。应限制已接受待握手流的数量，并在服务器关闭时关闭它们。

```goml
use std::io::Write;
use std::net;
use std::net::tls;
use std::io;

fn serve_one(certificate: Slice[byte], private_key: Slice[byte]) -> Result[(), io::Error] {
    let address = net::SocketAddr::new(net::IpAddr::V4([127, 0, 0, 1]), 8443);
    let listener = tls::TlsListener::bind(address, tls::ServerConfig::new(certificate, private_key))?;
    defer { let _ = listener.close(); };
    let stream = listener.accept()?;
    defer { let _ = stream.close(); };
    stream.write_all("hello".to_bytes().as_slice())
}
```

`TlsStream::local_addr()` 和 `peer_addr()` 在握手前后都可返回套接字地址，关闭后返回 `BrokenPipe`。`connection_info()` 返回协商的版本和 ALPN 协议，握手未完成时返回 `InvalidInput`。`read_with`、`write_with`、`read_exact_with` 和 `write_all_with` 接受 `context::Context`；exact/all 操作持有对应方向的门闩，并在所有部分传输间保持同一截止时间。读方和写方可以并发运行，同一方向的操作串行执行。取消活动 TLS I/O 或使其超时，会关闭连接以中断原生操作；之后需创建新连接。在原生 I/O 开始前检测到取消，包括等待门闩时，连接仍保持打开。显式本地关闭向被中断操作报告 `BrokenPipe`；远程 EOF 仍是成功的零字节读取。`close()` 是显式、共享、幂等的，并唤醒阻塞操作。`abort()` 共享同一关闭状态，但在 TLS 清理前关闭底层套接字，允许立即中断而不发送 TLS 关闭通知，包括普通 `close()` 已经在等待发送通知的情况。普通 `close()` 保留 TLS 关闭通知行为。对端在填满缓冲区前关闭时，精确读取报告 `UnexpectedEof`。失败时可能已经传输了部分数据。

```goml
use std::context;
use std::io;
use std::net::tls;
use std::time;

fn connect_example() -> Result[tls::TlsStream, io::Error] {
    context::with_timeout(context::Context::background(), time::Duration::from_seconds(5), |ctx, _| {
        tls::connect_with("example.com", 443, tls::ClientConfig::new("example.com"), ctx)
    })
}
```

<a id="http-ecosystem-packages"></a>

## HTTP 生态包

独立的 [`ecosystem::request`](https://github.com/gomlang/request/blob/HEAD/README.md) 客户端提供带作用域的 `send_stream` 下载、来自 `io::Read` 的 `send_reader` 上传、支持取消的变体、有界分块和复制，以及增量 SSE 解码。回调返回时响应句柄失效。流式操作通过 TCP 或验证过的 TLS 使用 HTTP/1.1，包括 CONNECT 代理；现有缓冲客户端保留 HTTP/2 路径。对于保留请求体的重定向，流式上传不能自动重放。压缩和 SSE 重连行为在客户端 README 中明确说明。

[`ecosystem::web`](https://github.com/gomlang/web/blob/HEAD/README.md) 服务器提供 TLS 监听器、WebSocket 响应升级、增量 multipart 提取、有界静态文件服务、CORS 和缓冲 gzip 协商。TLS 使用公开的 `std::net::tls` 待握手接受、握手、地址和中止 API；库及其使用方不需要用户 Go FFI 或 `go.mod`。监听器支持 HTTP/1.0 和 HTTP/1.1。Multipart 字段和升级连接遵循请求生命周期和取消规则。静态文件拒绝符号链接和路径穿越，支持 ETag、HEAD 和单个字节范围。这些库 API 使用普通导入，不增加语言文法。

<a id="filesystem-ecosystem-packages"></a>

## 文件系统生态包

递归目录遍历和文件系统通知是独立的第三方模块：[`ecosystem::walkdir`](https://github.com/gomlang/walkdir/blob/HEAD/README.md) 和 [`ecosystem::notify`](https://github.com/gomlang/notify/blob/HEAD/README.md)。它们使用公开的标准文件系统、系统调用、任务和时间 API。其源码不再以 `std::fs::walkdir` 或 `std::fs::notify` 的形式随工具链捆绑。

在模块根清单中添加依赖，并导入新路径：

```toml
[dependencies]
"ecosystem::walkdir" = true
"ecosystem::notify" = true
```

```goml
use ecosystem::walkdir;
use ecosystem::notify;
use std::fs;
use std::time;
```

公开遍历和通知 API 保持原有名称及行为。遍历器提供惰性迭代遍历、深度限制、剪枝、可选的链接跟随和带上下文的错误。观察器提供递归 inotify 维护、事件过滤、忽略子树剪枝、多路径观察集、带超时/可取消的读取和有界订阅。两者目前都面向 Linux amd64，提前停止时要求显式关闭资源。完整 API 示例、生命周期规则和平台限制见各模块 README。使用普通包/依赖语法；此次迁移不增加语言文法。运行 `goml update` 从[中央索引](https://github.com/gomlang/registry)获取模块，再使用普通的 `goml test` 和 `goml run --example <name>` 命令。

<a id="serde-and-json"></a>

## Serde 与 JSON

派生导入和处理器规则见[可编程派生](language/metaprogramming.zh-CN.md#programmable-derive)。

`std::serde` 定义与格式无关的 `Serialize`、`Deserialize`、`Serializer`、`Deserializer` trait 及两种派生。类型化格式使用 `Serialize::serialize` 和 `Deserialize::deserialize`：格式句柄是方法级泛型参数，因此单态化产生普通静态调用，无需运行时序列化器 trait 对象、Go 接口分派或运行时反射。派生结构体直接输出字段，派生枚举直接输出声明索引、传输名称、形状和载荷。派生的具名字段解码允许任意字段顺序，跳过未知字段，拒绝重复字段和缺失的必需字段。结构体字段、枚举变体和具名变体字段支持 `#[serde(rename = "wire_name")]`。字段默认值、省略、展开和枚举表示通过下述 Serde 属性选择。

`serde::Value` 仍是显式动态数据模型。它保留精确的有符号和无符号整数位宽、浮点位宽、枚举声明索引、字段顺序和变体形状。`Binary(Vec[byte])`、`Map(Vec[(Value, Value)])` 和 `Extension { tag: i8, data: Vec[byte] }` 分别保留二进制载荷、有序任意键映射项（含重复项）和不透明扩展数据。`Sequence`、`Tuple` 和 `Optional` 相互区分，`Value::Number(string)` 表示目标类型尚不确定的文本数字。`serde::to_value` 和 `serde::from_value` 通过 `ValueSerializer` 和 `ValueDeserializer` 在类型化值与此模型间转换；两者都返回 `Result`，使用它们意味着有意构造或消费完整值树。单元、布尔值、字符串、字符、数值原语、`Value`、`Vec[T]`、`Option[T]`、`HashMap[K, V]` 以及二元或三元元组都有标准直接实现。

`Serializer::serialize_extension(tag, data)` 和 `Deserializer::deserialize_extension()` 是可选事件，默认返回可恢复的不支持错误。格式实现需显式选择支持。值序列化器复制二进制和扩展载荷，并将映射保留为 `Value::Map`；值反序列化器也接受旧的键值元组序列映射表示。JSON 和 TOML 拒绝扩展事件，Bincode 没有扩展表示。TOML 保留先前值桥接的编码：二进制表示为整数数组，映射表示为键值对数组。`json::try_from_serde_value` 递归拒绝扩展。旧的不可失败 `json::from_serde_value` 仍是有损投影：二进制变为整数数组，映射变为键值对数组，扩展变为 `{ "tag": ..., "data": [...] }`。格式验证应优先使用经检查的转换。

`Serialize` 和 `Deserialize` 没有 `Value` 回退方法，`Deserialize` 不要求运行时 schema。手写实现必须实现泛型直接协议。`Schema` 只供 `bincode::decode_value` 等显式动态操作使用；它不会被派生，类型化解码也不查询它。

该设计消除了强制的中间树，但不保证绝对零开销或 Rust 式零复制。GoML 使用 GC，而非所有权和 `Deserialize<'de>` 生命周期；格式缓冲区和解码字符串仍在需要时分配，Go 编译器决定内联哪些静态调用。直接路径意味着本质上不要求数据模型树或动态分派；仍需针对每种格式和值形状测量分配。

```goml
use std::serde;
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
struct User {
    #[serde(rename = "display_name")]
    name: string,
    active: bool,
}
```

Serde 字段策略适用于结构体字段和具名枚举变体字段：

| 属性 | 编码 | 解码 |
| --- | --- | --- |
| `#[serde(default)]` | 普通字段 | 仅在缺失时调用 `Default::default()` |
| `#[serde(default = "function_path")]` | 普通字段 | 仅在缺失时调用无参数函数 |
| `#[serde(skip)]` | 省略字段 | 忽略输入并使用字段默认值 |
| `#[serde(skip_serializing)]` | 省略字段 | 普通字段；若允许缺失，应添加 `default` |
| `#[serde(skip_deserializing)]` | 普通字段 | 忽略输入并使用字段默认值 |
| `#[serde(skip_serializing_if = "predicate_path")]` | 谓词返回 `true` 时省略字段 | 普通字段；解码省略值需添加 `default` |
| `#[serde(flatten)]` | 将结构体或字符串/字符键映射合并进外层对象 | 将剩余具名字段解码进展开值 |

函数和谓词路径可以是带引号字符串或不带引号路径，在派生所在文件的作用域中解析，并按普通调用检查。谓词接受字段值并返回 `bool`；每次序列化只求值一次。属性可共用列表，例如 `#[serde(default, skip_serializing_if = "Option::is_none")]`。跳过字段隐式使用 `Default`，除非提供自定义 `default` 函数。派生只推断相关字段约束：双向跳过的字段不需要 `Serialize` 或 `Deserialize` 实现，使用普通默认值的泛型字段需要 `Default`。默认值不会修复已存在但格式错误的值或重复字段。除非指定 `default` 或解码跳过策略，`Option` 字段仍是必需的。

普通字段解码完成后，展开字段按声明顺序消费剩余键。嵌套展开结构体保留未消费的键，因此后续兜底映射可以接收它们。`Deserializer::preserve_unknown_fields` 是可选适配器钩子，默认不做任何操作；值适配器通过它在嵌套展开间保留这些键。应将兜底映射放在展开结构体之后。当键和值实现相应 Serde trait 且键实现 `Eq + Hash` 时，`HashMap[K, V]` 实现共享映射协议；重复输入键会被拒绝。映射项顺序遵循映射迭代顺序。展开拒绝重复输出键、非字符串/字符映射键以及非对象值。同一字段上的 `flatten` 不能与 `rename`、`default` 或任一跳过方向组合。未知属性、重复属性、重复传输名称和冲突策略会产生编译期诊断。

枚举默认保留外部标记表示。容器属性可选择以下替代形式：

| 枚举属性 | JSON 示例 |
| --- | --- |
| `#[serde(tag = "kind")]` | `{"kind":"Ready"}`、`{"kind":"Record","id":3}` |
| `#[serde(tag = "kind", content = "data")]` | `{"kind":"Ready"}`、`{"kind":"Number","data":3}`、`{"kind":"Pair","data":[3,true]}` |
| `#[serde(untagged)]` | `3`、`[3,true]`、`{"id":3}`、`null` |

内部标记枚举支持单元变体、具名变体和其值编码为对象的单字段变体。其他元组变体在编译期被拒绝。标记与具名字段冲突会被拒绝；来自展开值的冲突是运行时错误。相邻标记和无标记的单字段变体直接使用字段值；多字段元组变体使用数组，具名变体使用对象，无标记单元变体使用 `null`。无标记解码按声明顺序尝试变体，若没有变体接受完整载荷则返回错误。内部和相邻标记可以以任意字段顺序出现，必须为字符串，且不能重复。`content` 要求 `tag`，两个键必须不同，且都不能与 `untagged` 组合。

普通派生、默认值和固定字段跳过保留直接协议。固定跳过在 Bincode 中也使用连续位置索引。条件省略、展开和自定义枚举表示需要自描述的人类可读格式；Bincode 对这些策略返回可恢复错误。展开将展开字段缓冲为 `serde::Value`；自定义枚举表示缓冲枚举载荷，以处理标记顺序和无标记重试。因此，这些需主动选择的策略不具有普通流式路径的分配行为。它们支持 JSON、数据模型允许该值时的 TOML，以及显式 Serde 值适配器。

```gom
use std::json;
use json::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
struct ApiOptions {
    #[serde(default)]
    retries: isize,
    #[serde(default, skip_serializing_if = "Option::is_none")]
    label: Option[string],
    #[serde(skip)]
    cached: string,
}

#[derive(Serialize, Deserialize)]
#[serde(tag = "kind", content = "data")]
enum ApiReply {
    Ready,
    Options(ApiOptions),
}
```

`Serializer::is_human_readable()` 和 `Deserializer::is_human_readable()` 默认为 `true`。JSON、TOML 和通用值适配器使用此默认值；Bincode 将其覆盖为 `false`。自定义格式应在两端一致覆盖。自定义 Serde 实现可以为文本格式选择可读名称，为二进制格式选择整数表示，无需依赖特定编码器。此查询不改变普通派生序列化。

`std::json` 支持两种有意区分的模式。值模式使用 `json::Value`、`json::parse` 和 `json::encode` 进行无 schema 的检查和编辑。JSON 数字在 `Value::Number` 中保留精确源文本。类型化模式中，`json::try_to_string` 和 `from_string` 通过流式 serde trait 直接写入和消费 JSON；两者都不先构造 `json::Value` 或 `serde::Value` 树。`json::to_value` 和 `from_value` 返回 `Result`，是通向动态 JSON 模型的显式桥接。数值范围和目标位宽检查在反序列化到所请求类型时进行。

对于非整数文本或超出 isize 范围的值，`json::as_int` 返回 None；溢出时不回绕。手工构造的 Number 文本可在此访问器中保留前导零，而 parse 仍要求 JSON 数字语法。包含多字节字符的无效字面量前缀产生可恢复解析错误，而不会在 UTF-8 标量内部切片。

`json` 公开重新导出共享的 `Serialize`、`Deserialize` trait 及派生处理器，因此 `use serde::Serialize` 和 `use json::Serialize` 选择相同实现身份。JSON 序列化器按源码顺序输出结构体字段。直接映射使用 JSON 对象，因此要求键的直接表示为字符串或字符；其他键类型使 `try_to_string` 返回可恢复错误。反序列化器接受任意字段顺序，递归跳过未知值，拒绝重复和缺失字段，并拒绝尾随输入。类型化错误保留字节偏移和嵌套结构体、序列、映射或枚举路径。JSON 默认使用外部标记枚举：单元变体为字符串，元组变体为值是数组的对象，类似结构体的变体为值是另一个对象的对象。Serde 容器属性可按上文选择内部、相邻或无标记表示。

<a id="bounded-json-streaming-and-typed-serde"></a>

### 有界 JSON 流式处理与类型化 serde

`json::Decoder[R]::new(reader, limits)` 接受 `io::Read` 源和显式 DecodeLimits。`next_token()` 返回 `Result[Option[Token], Error]`；token 包括 StartArray、EndArray、StartObject、EndObject、Key(string)、String(string)、Number(string)、Bool(bool) 和 Null。解码器增量验证文法，只保留固定传输缓冲区、容器栈和有界标量载荷。数字保留精确拼写。`next_value()` 使用显式栈组装一棵有界 Value 树，保留对象顺序和重复键。Token 载荷和返回树不借用补充读取缓冲区。

流中的值必须由 JSON 空白分隔：接受 `{} []`，拒绝 `{}[]`。根之间的正常 EOF 返回 None；token 或容器内部的 EOF 是错误。`next_value()` 要求处于根边界，不能从 next_token 已开始的容器中间继续。支持在完成一个根后切换 token、Value 和类型化模式。`finish()` 消费剩余空白并要求 EOF，不解析第二个值。在根中间结束会产生 State 错误。

`decode[T: Deserialize]() -> Result[Option[T], Error]` 通过同一增量 token 文法直接驱动 serde，不构造中间 Value。结构体字段、元组、可选值、字符串/字符映射键和外部标记枚举遵循现有 serde 表示。未知结构体字段以迭代方式跳过，但仍检查语法和字节/结构限制。派生反序列化器拒绝重复或缺失的必需字段。整数位宽和浮点范围错误属于 Schema 失败。显式 deserialize_any 使用迭代栈构造 serde::Value，保留数字拼写。不完整、乱序或多根的自定义 serde 事件序列会失败，包括忽略适配器错误的情况。用户代码保留的适配器在其操作返回时失效。

`from_reader[T, R](reader, DecodeLimits)` 和 `from_string_with_limits[T](input, DecodeLimits)` 要求恰好一个类型化值，后跟 EOF。`value_from_reader(reader, limits)` 和 `parse_with_limits(input, limits)` 提供对应的 Value 操作。对这些单值便捷函数，空输入是 UnexpectedEof。

`offset()` 报告逻辑上已消费的字节数。`buffered()` 返回已确认但未消费输入的独立 Bytes 快照；底层读取可能超前于逻辑游标。严格的 JSON 转义、代理对以及字符串中的 UTF-8 可跨补充读取边界识别。EOF 处不完整的 UTF-8 为 UnexpectedEof；字符串中无效 UTF-8 字节为 Utf8 错误。拒绝注释、BOM、尾随逗号和非 JSON 数字拼写。

`json::Encoder[W]::new(writer, limits)` 接受 `io::Write` 和 EncodeLimits。`encode[T: Serialize](value)` 直接向写入器发送类型化 serde 事件，不先创建 Value 树或完整输出字符串。`encode_value` 显式遍历 Value，而非通过 serde::Value 递归转换。两者在每个成功根后追加 LF，将 LF 计入字节预算，并以 usize 返回本次操作确认的字节数。`written()` 是累计确认字节数。编码保留字段顺序、字符串/字符映射键和外部标记枚举。数字文本必须具有有效 JSON 语法；拒绝 NaN 和无穷。旧的不可失败 `encode(Value)` 保留现有行为和签名。`to_writer(value, writer, EncodeLimits)` 和 `encode_to_writer(Value, writer, EncodeLimits)` 写入单个紧凑根，不附加 LF。`to_string_with_limits(value, limits)` 和 `encode_with_limits(Value, limits)` 返回其有界字符串表示。旧字符串 API 保留原契约，不会隐式获得新限制。

DecodeLimits 包含非负的 max_input_bytes、max_value_bytes、max_depth、max_values、max_string_bytes、max_number_bytes、max_container_entries、max_path_bytes、max_serde_frames 和 max_events。EncodeLimits 字段相同，但将 input/value 字节限制替换为 max_output_bytes。构造函数在任何提供方调用之前拒绝负上限。不提供隐式无限默认值。输入和输出字节配额是整个生命周期的总量；其他结构配额按根重置。Events 统计每次公开操作中可失败的协议调用和内部分派/组装工作，不统计经过的时间。

Depth 统计实际 JSON 数组和对象，根标量深度为零。Values 包括根和嵌套值，不包括对象键。Container entries 分别限制每个数组或对象。String bytes 统计解码后的 UTF-8，包括键；number bytes 统计原始 ASCII 表示。Value bytes 统计原始根的跨度，包括内部空白，但不包括根之间的空白。协议帧有独立限制，使嵌套可选 serde 事件无法绕过深度限制。Path bytes 在添加路径段之前计入字段/变体名字节数和十进制索引位数。新路径段将超出预算时，错误只保留允许的前缀。

每个读取请求均不超过剩余输入额度。即使全部值字节已读取，证明 EOF 或数字结束边界仍可能需要未用额度。不发起额外超预算探测。只确认合法成功的提供者数量；失败调用可能具有未知副作用。Interrupted 不重试，零写入以 WriteZero 失败，无效数量保留 InvalidData I/O 原因。使用原始 read/write 方法，不使用可重写的整流便捷方法。两端均不 flush 或关闭提供者。任何流式失败（包括最终 LF 写入失败）都可能使输出包含部分 JSON 前缀；这不是原子事务。

错误访问器为 kind()、byte_offset()、path()、message()、cause()、confirmed_read() 和 confirmed_written()。ErrorKind 区分 InvalidLimit、Limit(LimitKind)、Syntax、UnexpectedEof、TrailingData、Utf8、Schema、State 和 Io。LimitKind 指明具体配额。路径是只读的结构化 Field(string)、Index(isize) 和 Variant(string) 段，不是有歧义的点分字符串。偏移表示逻辑输入位置或已确认输出字节数；I/O 进度访问器返回整个生命周期的 usize 总量。Error 支持 ToString 和 Debug，无需调用方解析诊断文本。

副本共享游标、预算和终止状态。失败后，后续操作返回首个失败，不再执行 I/O。通过别名重入的操作会被拒绝，并使会话终止。重入失败后提供者返回的合法数量仍计入已确认进度；成功读取的字节仍可通过 buffered 访问，不再解析。用户 serde/提供者 panic 会传播，展开过程使该会话不可再用。用户代码可以独立于这些配额分配或循环；字节限制无法中断阻塞提供者。不支持并发访问。编码时调用方必须保持输入 Value 容器稳定；检测到长度变化会可恢复地失败，但同长度修改不构成快照协议。循环通过结构或事件限制终止。库的 token/Value 遍历及未知字段跳过采用迭代方式；用户编写或派生的递归类型化 serde 代码仍使用自身调用栈。限制是逻辑工作/存储边界，不保证分配失败可恢复。不同于 Go 的宽松字符串替换，格式错误的 UTF-8 和未配对的代理项转义会被拒绝。字段匹配仍严格遵循 GoML serde 规则；没有反射、不区分大小写的字段回退或自动 base64 字节数组映射。Value/token 解码保留对象顺序和重复键；自定义映射/结构体反序列化器自行决定重复字段策略。

```goml
use std::json;
use json::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
struct User {
    name: string,
    active: bool,
}

fn encode_user(value: User) -> Result[string, string] {
    json::try_to_string(value)
}

fn decode_user(input: string) -> Result[User, string] {
    json::from_string(input)
}
```

<a id="bincode-typed-binary-data"></a>

## Bincode 类型化二进制数据

`std::bincode` 主要是类型化 serde 格式。`encode_to_vec` 将请求类型直接写入输出缓冲区，`decode_from_slice` 让请求类型直接从输入拉取字段，并返回值及已消费字节数。类型化路径不构造 `serde::Value`，也不请求运行时模式。底层 `encode_value` 和 `decode_value` 仍可用于显式动态处理；`decode_value` 需要显式提供模式，因为 bincode 在线路数据中不携带字段类型、元组长度或结构体布局。

类型化编码器和解码器验证每个复合值的开始、元素或字段以及结束状态转换。格式错误的手写 serde 实现会返回错误，不产生部分数据或 panic。解码错误包含字节偏移，以及经过位置字段、集合元素、映射条目和枚举变体的路径。

`bincode::standard()` 使用小端变长整数编码。`bincode::legacy()` 使用小端定宽整数。`with_little_endian`、`with_big_endian`、`with_variable_int_encoding` 和 `with_fixed_int_encoding` 返回调整后的配置。线路表示遵循 bincode 2 关于布尔值、ZigZag 有符号变长整数、整数标记、IEEE 浮点位、UTF-8 字符串、集合长度、单字节 option 标签、源码顺序结构体字段和枚举声明索引的约定。

```goml
use std::bincode;
use bincode::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
struct Message {
    id: u32,
    text: string,
}

fn round_trip(value: Message) -> Result[Message, string] {
    let encoded = bincode::encode_to_vec(value, bincode::standard())?;
    let decoded: (Message, isize) = bincode::decode_from_slice(
        encoded.slice(0, encoded.len()),
        bincode::standard(),
    )?;
    Ok(decoded.0)
}
```

类型化 bincode 支持共享 serde 原语、通过自定义直接实现处理的字节切片、`Vec`、`Option`、二元和三元元组，以及派生结构体和枚举。定长数组仍需显式 serde 实现；标准库尚未迁移到 const 泛型数组实现，因为其自举编译器必须保持兼容。动态 `serde::Value`、`json::Value` 和 `toml::Value` 不描述类型化 bincode 解码所需的具体二进制布局。

<a id="toml-values-and-typed-documents"></a>

## TOML 值和类型化文档

`std::toml` 采用与 JSON 相同的公共 API 划分。`toml::Value` 具有字符串、有符号 64 位整数、64 位浮点数、布尔值、日期时间文本、数组和表变体。`parse` 和 `encode` 操作这种无模式表示。`to_value`、`from_value`、`to_string` 和 `from_string` 使用共享 serde trait，并强制目标类型。由于头部和点分路径需要全局文档组织，TOML 仍在输出前规划完整表树；它兼容流式 trait，但目前不是完全直接的格式。超过 TOML 有符号整数范围的无符号值会被拒绝；`()` 和 `None` 无法编码，因为 TOML 没有 null 值。

TOML 解析器接受基本字符串和字面量字符串、Unicode 转义、布尔值、十进制和带进制前缀的整数、浮点数、日期时间文本、数组、内联表、点分键和普通表头。它保留表和字段顺序，以产生确定性输出。尚不支持多行字符串、表数组头和内联表中的点分键。

```goml
use std::toml;
use toml::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
struct Server {
    host: string,
    port: u16,
}

fn load(input: string) -> Result[Server, string] {
    toml::from_string(input)
}
```

<a id="comparison-and-sorting"></a>

## 比较和排序

排序原地修改 `Vec[T]`。`sort` 和 `stable_sort` 使用 `cmp::Ord`；`sort_by_ordering` 和 `stable_sort_by_ordering` 使用 `cmp::Ordering`；`sort_by` 和 `stable_sort_by` 接受返回负数/零/正数的整数比较器。全部排序变体都是稳定排序。`binary_search` 及基于比较器的变体要求向量已排序，返回首个匹配索引。

`std::cmp` 提供 `PartialOrd` 和 `Ord`。`()`、`bool`、`string`、`char` 以及全部有符号和无符号整数类型均实现二者。浮点值只实现 `PartialOrd`，因为 NaN 不构成全序；与 NaN 比较产生 `None`。元组、定长数组、`Vec`、`Slice`、`Option` 和 `Result` 条件性实现比较 trait，使用字典序。`cmp::compare` 返回 `Ordering::Less`、`Equal` 或 `Greater`；`Ordering` 支持谓词、反转、字典序串联，以及转换为负数/零/正数整数约定。`cmp::Reverse[T]` 反转已有顺序。最小值超过最大值时，`cmp::clamp` 返回错误。

<a id="dynamic-channel-selection"></a>

## 动态通道选择

`std::channel` 提供运行时确定数量的同类型选择。`Operation[T]` 为 `Receive(Receiver[T])` 或 `Send(Sender[T], T)`。`select` 阻塞，`try_select` 在没有操作就绪时返回 `Ok(None)`，`try_select_priority` 从最低输入索引开始探测。成功结果为 `Selection::Received { index, value }` 或 `Selection::Sent { index }`；已关闭的接收携带 `None`。空操作切片返回 `SelectError::Empty`。

<a id="structured-concurrency"></a>

## 结构化并发

`std::task` 在 goroutine 之上创建词法任务作用域：

```goml
use std::task;

fn load_pair() -> Result[(isize, isize), string] {
    task::try_scope(
        |scope: task::Scope| {
            let left = scope.spawn_try(|cancel| load_left(cancel));
            let right = scope.spawn_try(|cancel| load_right(cancel));
            Ok((left.join()?, right.join()?))
        },
    )
}
```

上下文 `scope` 和 `spawn` 形式通过隐藏的作用域能力提供相同的结构化生命周期：

```goml
use std::task;

fn load_pair() -> Result[(isize, isize), string] {
    scope {
        let left = spawn |cancel| load_left(cancel);
        let right = spawn |cancel| load_right(cancel);
        Ok((left.join()?, right.join()?))
    }
}
```

`scope { body }` 在 HIR 检查前降级为 `task::scope(|hidden_scope| body)`，每个直接嵌套的 `spawn |cancel| body` 降级为 `hidden_scope.spawn(|cancel| body)`。文件必须导入 `std::task`。直接嵌套的词法作用域使用 `scope_with`，继承父作用域的取消。隐藏作用域值不能被命名或返回。词法 `spawn` 不能跨越用户闭包边界，从而防止返回的闭包捕获该能力；请改为在闭包内部创建嵌套 `scope`。在这些上下文形式之外，`scope` 和 `spawn` 仍是普通标识符。

`scope` 在主体返回后停止接受新工作，等待所有直接子任务，再返回主体的值。即使丢弃 `Task[T]` 句柄，它仍会等待。`Task::join` 可以重复或并发调用；每次调用观察到相同的存储结果。`Scope::try_spawn` 在关闭开始前成功注册任务时返回 `Some(task)`，在作用域开始关闭后返回 `None`。普通 `Scope::spawn` 对已关闭作用域保留更严格的运行时错误行为。

`try_scope` 在主体返回 `Err` 时取消作用域，等待全部直接子任务退出，再返回原始错误。`Scope::spawn_try` 也会在其子任务返回 `Err` 时立即取消兄弟任务。使用 `scope_with(parent_token, body)` 或 `try_scope_with(parent_token, body)` 创建的嵌套作用域继承取消。

取消是协作式的。`Scope::cancel` 改变 `CancelToken::is_cancelled` 观察到的状态，不强制终止 goroutine。阻塞工作可以使用：

- `task::recv_with(token, channel) -> WaitResult[Option[T]]`
- `task::send_with(token, channel, value) -> WaitResult[()]`
- `time::sleep_with(token, duration) -> WaitResult[()]`
- `command.output_cancel_structured(token) -> Result[process::Output, process::Error]`
- `command.status_cancel_structured(token) -> Result[process::ExitStatus, process::Error]`

`CancelToken::done() -> Receiver[()]` 暴露作用域上下文共享的完成通道，`Task::done() -> Receiver[()]` 暴露任务共享的就绪通道。两个方法都不启动桥接 goroutine。任务在关闭就绪通道前存储结果，因此完成事件发生后可以立即通过 `join()` 观察结果。

`std::time::Timer::new(duration)` 创建可停止的单次计时器。`done()` 返回其 `Receiver[()]`，`stop()` 报告是否阻止了待发生的触发，`time::after(duration)` 是单次计时的便捷形式。成功停止会使完成通道保持未就绪。计时器触发和停止经过同步，确保通道最多关闭一次。

取消唤醒通道和 sleep 操作时，它们返回 `WaitResult::Cancelled`。进程操作则返回种类为 `io::ErrorKind::Interrupted` 的 `process::Error`。进程取消使用宿主命令上下文，因此作用域在退出前等待进程操作返回。任务作用域从不自动关闭用户通道。`active_scope_count()` 暴露活动运行时作用域数量，用于测试和泄漏诊断。

GoML 没有生命周期或线性类型系统，因此 `Scope` 值目前可以逃逸出主体。作用域开始关闭后调用 `spawn` 是运行时错误。Panic 不隐式转换为 `Result`。作用域主体或子任务中的 panic 会取消兄弟任务、等待它们、移除运行时作用域，然后在作用域拥有者中重新抛出。在作用域外显式使用 `std::panic::catch`，可以在作用域完成后恢复这个重新抛出的 panic。子任务内部的 catch 会在任务作用域观察到 panic 前处理它。取消仍是协作式的，因此不配合的兄弟任务可能推迟作用域完成和恢复。

`join_all` 按输入顺序返回值。`join_all_results` 等待每个任务，按输入顺序返回错误，不受 goroutine 调度影响。

`join_all_indexed_results` 保留每个错误的输入索引。`for_each_concurrent(limit, values, body)` 同时最多运行 `limit` 个调用，limit 非正时使用一个工作任务，等待每个值，并按输入顺序返回带索引的错误。`ConcurrencyLimit::run` 可将相同的协作限制应用于自定义任务布局。

`race(bodies)` 返回首个完成的值，取消剩余主体，仍等待所有落败主体退出。空输入返回 `None`。因此不配合取消的主体可能推迟 `race` 返回。

<a id="panic-boundaries"></a>

## Panic 边界

`std::panic` 提供对语言和宿主运行时 panic 的显式恢复：

```goml
use std::panic;

fn validate_internal_state(valid: bool) -> () {
    if !valid {
        panic::raise("invalid internal state")
    }
}

fn guarded() -> Result[(), panic::Panic] {
    panic::catch(|| validate_internal_state(false))
}
```

`raise(message: string) -> never` 开始展开。`catch[T](body: () -> T) -> Result[T, Panic]` 在回调正常返回时返回 `Ok`，包括回调返回普通 `Err` 的情况。只有发生 panic，且所有退出块已运行注册的 defer 后，它才返回 `Err`。边界错误和其他可恢复 Go 运行时 panic（包括跨越 FFI 回调的 panic）在同一 goroutine 上捕获。执行从 `catch` 之后恢复，不是从失败表达式处恢复。这不是事务：共享状态修改和 I/O 不回滚。

`Panic.message()` 和 `to_string()` 返回诊断文本；宿主 panic 诊断中的无效 UTF-8 替换为 U+FFFD。`stack_trace()` 返回在首个 `std::panic::catch` 边界捕获的 Go 栈，可能出现生成的函数名。子任务 panic 在转移给作用域拥有者前捕获，而作用域主体 panic 可能在作用域重新抛出后才首次捕获。`resume(info: Panic) -> never` 重新抛出，并在后续 GoML catch 中保留该诊断文本和栈，包括跨生成 Go 包的情况。原始载荷保持私有并由垃圾回收管理，没有全局 panic 注册表。若 defer 在清理时 panic，其余 defer 仍运行，最新 panic 传递到 catch 边界。

未捕获的 panic 终止进程。Catch 不拦截其他 goroutine 的 panic、`process::exit`、运行时致命错误或强制终止。恢复并不证明共享应用状态仍一致；应在隔离请求或任务周围使用边界，对预期失败仍使用 `Result`。运行时钩子属于 `lib/builtin`；公共 API 和消费者均是普通 GoML，无需用户 Go FFI 适配器。

<a id="cancellation-contexts"></a>

## 取消上下文

`std::context` 组合协作式取消和单调截止时间。`Context::background()` 从不取消且无截止时间；`Context::from_token(token)` 适配已有任务 token。`with_cancel(parent, callback)`、`with_timeout(parent, duration, callback)` 和 `with_deadline(parent, Deadline::after(duration), callback)` 在词法任务作用域内调用 `callback(context, cancel_handle)`。子上下文继承较早的截止时间、传播父取消，并在回调返回时取消。内部监视任务和计时器在作用域退出前被等待或停止；逃逸的子上下文已处于取消状态。

`CancelHandle::cancel()` 幂等。首个取消原因保留为 `Error::Cancelled` 或 `Error::DeadlineExceeded`。`Context::check()` 返回 `Result[(), Error]`；`error()`、`deadline()` 和 `remaining()` 提供可选状态。`done()` 是可用于 `select` 的 `Receiver[()]`；`token()` 暴露供现有 API 使用的可选任务 token。`sleep(duration)` 协作式等待并报告取消原因。`Deadline` 使用 `time::Instant`，因此墙上时钟调整不会改变它。取消不会强制停止任意用户代码或不支持上下文的阻塞 I/O。

<a id="collectionsindexmapk-v"></a>

## `collections::IndexMap[K, V]`

`IndexMap` 是按插入顺序排列的哈希映射。键类型必须实现 `Eq` 和 `Hash`，实际键比较使用 `PartialEq`：

```goml
use std::collections;

let headers: collections::IndexMap[string, string] = collections::IndexMap::new();
headers.insert("content-type", "text/plain");
headers.insert("content-length", "12");
headers.insert("content-type", "application/json");
for (name, value) in headers {
    println(name + ": " + value)
}
```

插入新键将其追加到迭代顺序末尾。替换已有值不移动键。移除键后再次插入会将它追加到末尾。

常用方法为 `new`、`with_capacity`、`len`、`is_empty`、`contains`、`get`、`insert`、`remove`、`reserve`、`clear`、`entries`、`keys`、`values` 和 `iter`。`insert` 和 `remove` 以 `Option[V]` 返回旧值。`entries`、`keys` 和 `values` 返回有序快照，`iter` 和 `for` 按插入顺序遍历 `(K, V)` 对。

实现使用稀疏开放寻址索引表和按插入顺序排列的条目数组。已删除条目变为墓碑，在后续增长或删除密度较高时压缩。查找、插入和移除的期望复杂度为 O(1)；迭代和压缩为 O(n)。迭代器活动期间不支持结构性修改。

`IndexMap` 目前没有字面量或索引语法。请使用 `insert` 和 `get`。

<a id="common-tasks"></a>

## 常见任务

| 任务或易错点 | GoML 写法 |
| --- | --- |
| 为领域枚举错误提供可读消息 | `#[derive(ToString, Debug)]` 配合各变体的 `#[to_string("message {field}")]`；无需额外实现 `Error` 方法 |
| `Simd[T, N]` 或向量 `a + b` | 导入 `std::simd`，选择固定 128/256 位向量类型，使用 `a.add(b)` |
| `float_value.to_i32()` | 导入 `std::num::TryToInt`，使用 `float_value.try_to_i32()`；处理非有限值和范围错误 |
| 遍历目录树 | 使用 `ecosystem::walkdir` 依赖，在 Linux amd64 上进行基于系统调用的深度优先迭代，支持深度限制、剪枝、可选链接跟随和逐路径错误 |
| 验证 C 兼容字符串 | `std::c::CString` 拒绝内部 NUL；`c::literal` 支持编译期验证 |
| 操作逻辑斜杠路径或 shell 模式 | `std::path::slash` 提供纯词法操作和有界 Unicode 匹配；宿主路径使用 `std::path` |
| 监视目录树变化 | 在 Linux amd64 使用 `ecosystem::notify` 依赖及其 `watch_recursive` 或 `WatchSet`，通过 `Options` 剪除忽略路径，使用定时读取或作用域订阅，处理 `Event.rescan` 并关闭句柄 |
| TCP 和 UDP 网络 | `std::net` 提供套接字、DNS 解析、共享 epoll 就绪机制、显式关闭和感知上下文的等待；`std::net::tls` 提供经过验证的 TLS 客户端和监听器、待处理连接接收、显式握手及对端地址 |
| 流式 HTTP 请求或消费 SSE | `ecosystem::request` 提供作用域流式回调、`io::Read` 上传、有界响应块，以及 HTTP/1.1 上的增量 SSE 解码 |
| 提供 HTTPS、WebSocket 或上传文件服务 | `ecosystem::web` 提供 TLS 监听器、显式 WebSocket 升级、multipart 字段、静态文件、CORS 和有界缓冲 gzip 中间件 |
| 转义 URL 路径段或查询参数 | `std::net::url` 的有界组件/查询编解码保留解码字节；完整 URL 解析是独立操作 |
| 解析原始 URL 引用 | `url::Reference::parse` 和 `base.resolve` 保留组件标记并移除字面量点段；authority 验证和 ASCII 序列化是显式独立操作 |
| 检查 URL 主机或隐藏密码 | `url::Authority::parse` 验证用户信息、主机字面量和十进制端口语法；`redacted` 只隐藏密码，不隐藏查询或其他秘密 |
| 保持字节快照独立 | `Bytes::copy` 创建可变副本；`Bytes::freeze` 和 `FrozenBytes` 创建不可变快照 |
| 派生集合版本并保留原版本 | 使用 `std::collections::imm` 的映射、集合、Vector 和 List；保存每次更新返回的新集合 |
| 查询有序不可变键或范围 | `imm::BTreeMap` / `BTreeSet` 提供 `lower_bound`、`upper_bound` 和半开 `range(start, end)` |
| 选择不可变索引序列 | 连续快照用 `FrozenVec`；反复进行路径复制更新用 `imm::Vector`；两者都不深度冻结可变元素 |
| 反复搜索二进制数据 | `bytes::Finder` 对模式生成快照并复用线性时间搜索状态；`cut` 返回共享视图，不复制缓冲区 |
| 对二进制输入应用 Unicode 规则 | `bytes::fields`、`trim_space` 和标量搜索保留视图字节；`map_chars` 输出有效 UTF-8；`replace_invalid_utf8` 使用调用方选择的字节和显式输出限制，对每段格式错误序列替换一次 |
| 搜索或遍历文本而不收集各部分 | `text::Finder` 使用字节偏移和线性搜索；`split_iter` 保留 GoML 空分隔符语义，`lines_inclusive_iter` 保留换行字节 |
| Unicode 空白和不区分大小写的文本 | `text::fields`/`fields_iter` 和 `trim_space` 使用固定版本 Unicode 数据；`equal_fold` 是简单标量折叠，不是多标量完整折叠；旧 `trim` 仍仅处理 ASCII |
| 限制生成文本 | `text::*_checked` 构造按 UTF-8 字节计数；`map_chars` 对每个访问的标量调用有状态回调一次，绝不返回部分结果 |
| 编解码带引号字面量 | `text::quote`/`quote_char` 输出有界 Go 风格字面量；`quoted_prefix` 识别一个前导字面量；`unquote_element` 区分字节/标量值；`unquote_bytes` 保留字节，`unquote` 验证解码后的 UTF-8；这些不是 JSON 或 GoML 语法扩展 |
| 解析布尔值 | `num::parse_bool_structured` 接受文档列明的精确写法，对其他输入返回错误；布尔 `ToString` 输出规范小写文本 |
| 区分浮点解析失败 | 结构化浮点解析返回可恢复错误；`ParseFloatError::is_range` 标识溢出，而显式无穷和舍入下溢仍是成功值 |
| 格式化定点浮点值 | `num::format_float32_fixed`/`format_float64_fixed` 使用精确十进制转换、平局舍入到偶数和显式输出字节限制 |
| 解析复数数值文本 | `num::parse_complex_f32`/`parse_complex_f64` 返回类型化实部/虚部元组及语法/范围错误；不引入复数源码字面量或算术类型 |
| 格式化复数分量 | `num::format_complex_f32`/`format_complex_f64` 使用类型化 `ComplexFormat` 模式，限制覆盖两个分量、符号和括号 |
| 复数计算 | `std::math::complex::Complex` 提供 f64 算术和主值初等函数；除零返回 `DivisionError` |
| 复用多个字面量替换规则 | `text::Replacer` 按规则/输出/工作预算编译有序规则；`chunks` 惰性产生可失败片段；`text::stream::write_replaced` 处理短写入并报告部分进度 |
| 从不可变文本读取 | `io::StringReader` 提供顺序/定位字节读取、标量读取、回退、重置和检查式字节 seek，不需要可写缓冲区 |
| 通过共享契约定位流 | `io::Seek` 使用 `SeekFrom` 和检查式字节位置；Cursor/SectionReader 保持有界，StringReader 允许超过 EOF，OffsetWriter 不支持 End |
| 读取并回退缓冲字符 | `io::BufReader` 读取字节或 UTF-8 标量，在符合条件的读取后支持一次回退，即使配置容量为一，也在 I/O 错误时保留已确认标量前缀 |
| 读取有界流片段 | `BufReader::read_fragment` 和 `read_line_fragment` 区分分隔符/满缓冲区/EOF 边界，错误时暴露已消费部分数据；行载荷省略 CRLF，同时报告原始消费量 |
| 组合缓冲输出 | `BufWriter::write_byte/write_char/write_string/read_from` 保留显式 flush，报告输入接收量，不保证目标已收到 |
| 连接并发字节生产者和消费者 | `io::pipe` 使用现有通道；写入等待读取，关闭唤醒等待者，`write_progress` 保留已确认前缀 |
| 注入只读文件系统 | `fs::FileSystem`/`File` 使用关联句柄；有界 read/stat 辅助函数关闭句柄；`SubFs` 组合逻辑名称，不提供沙箱 |
| 枚举可移植目录 | `ReadDirFile` 提供带显式部分错误的分页；`read_dir_from` 对已确认条目施加边界、验证并稳定排序，然后关闭句柄 |
| 遍历可移植文件系统 | `walk_from` 使用显式栈、类型化跳过控制、错误回调，以及显式深度/目录/工作预算 |
| 对可移植路径进行 glob 匹配 | `glob_from` 匹配斜杠分隔组件，使用全局预算、排序后的部分结果和显式提供者错误 |
| 检查可移植符号链接 | `LinkFileSystem` 提供静态 read-link 和不跟随链接的元数据能力；SubFs 保留原始目标 |
| 描述可移植元数据 | `FileMode` 保留高级种类及已知/未知标志；SnapshotEntry 缓存元数据，有界显示辅助函数避免隐式 I/O |
| 构建内存文件系统 | `std::testing::fs::MemoryFS` 对输入字节生成快照、推断目录、在显式预算下解析链接，并打开独立 read/seek/目录句柄 |
| 检查可移植文件系统 | `std::testing::fs::check_fs` 返回有界、可恢复报告；`check_helpers` 增加标准辅助函数/SubFs 比较，静态 Seek/ReadAt/链接检查共享相同错误和预算模型 |
| 测试短 I/O 和延迟失败 | `std::testing::io::{FragmentReader, FragmentWriter}` 限制每次传输，并可在精确字节数后失败 |
| 组装类型化格式字段 | `std::text::format::Formatter` 组合有界文本、整数、浮点和通过 trait 转换的字段；f-string 语法不变 |
| 格式化科学记数法/有效数字值 | `num::format_float*_scientific` 和 `format_float*_general` 系列复用精确舍入；精度分别控制小数位或有效数字 |
| 格式化精确的二的幂表示 | `num::format_float*_binary` 输出整数有效数及二进制指数；`format_float*_hex` 支持精确或舍入的十六进制小数 |
| 格式化最短十进制表示 | `num::format_float*_shortest` 寻找可往返转换的有效数字，并按显式 `FloatNotation` 渲染；不改变默认 ToString |
| 限制文本切分 | `text::split_n_checked` 将数量与结果部分限制分离，并把空分隔符分解为标量；旧 `split` 保留整个字符串 |
| 切分或扩展不可信字节数据 | `bytes::split_n` 将数量语义与部分数量限制分离；`replace_n`、`join` 和 `repeat` 检查显式字节限制并返回 `TransformError` |
| Go `maps` 辅助函数 | `HashMap` 上的 `std::collections::map_*`；浅复制和快照迭代器，没有 nil 映射或隐式 Go map 转换 |
| Go `slices` 辅助函数 | `std::collections::slice_*`、检查式 `vec_*` 编辑及现有 Vec/Slice 方法；只读视图没有重新切片容量或 nil 区别 |
| 增量解码任意 UTF-8 | `std::utf8::Decoder` 严格解码并报告绝对偏移；`decode_rune`、`decode_last_rune` 和 `decode_lossy` 显式替换格式错误字节 |
| Unicode 分类和无大小写匹配 | `std::unicode` 提供固定版本类别/脚本/属性表；区分单标量 `simple_fold` 循环、多标量字符串 `case_fold` 和区域专属大小写转换 |
| Go `sort.Search` 或重复值边界 | `collections::search` 检查负长度；`lower_bound`、`upper_bound` 和 `equal_range` 返回插入点/范围，`binary_search` 保留首个匹配的 `Option` |
| 编码变宽整数 | `std::bytes::endian::{append_uvarint, append_varint}`；检查式读写返回 `Result` |
| 编码固定布局结构体或数组 | 在 endian 的 `read_with`/`write_with` 回调中显式处理字段/循环；内存写入先暂存再原子提交，流操作要求字节限制 |
| 计算流式校验和 | 通过包别名导入 `std::hash::Hasher`，使用 Adler32/CRC/FNV 摘要；校验和不是 MAC |
| 计算密码学摘要 | `std::crypto::hash` 以纯 GoML 实现增量 SHA-2/SHA-3 和显式 SHAKE 吸收/挤出 |
| 比较预期摘要字节 | `std::crypto::{digest_equal, verify_digest}` 比较等长值的全部字节，不声称编译器强制恒定时间 |
| 认证或派生密钥材料 | `std::crypto::hmac` 使用新摘要工厂；`std::crypto::hkdf` 和 `std::crypto::pbkdf2` 强制输出和工作限制 |
| 对成功 I/O 进行哈希 | 用 `std::hash::stream` 包装读取器或写入器；仅成功报告的字节进入给定 Hasher |
| 检测双倍位宽除法溢出 | 使用 `std::math::bits::div32/64`，返回 `Result` 而非 panic |
| 获取可复现随机样本 | `std::rand::Generator` 包装显式 PCG 或 ChaCha8 状态并检查有界范围；旧 splitmix64-v1 函数保留原输出 |
| 保留操作和清理错误 | `std::resource::{with_cleanup, scope, ScopeError}` 或 `io::with_resource` |
| 添加类型化错误上下文 | 导入 `std::error::{ResultExt, ReportContext}`；在报告边界调用一次 `.into_report().with_context(...)`，之后直接追加上下文 |
| 适配库错误报告 | `Report::map` 转换叶节点，同时保留上下文、聚合边界和遍历顺序 |
| 隔离运行时 panic | 使用 `std::panic::catch` 和词法 `defer` 清理；`raise` 和 `resume` 返回 `never` |
| 创建并监督 Linux 子进程 | `std::os::linux::process::Command` 提供运行时协调的 spawn、稳定 pidfd 信号和共享等待结果 |
| 访问映射内存 | `std::os::linux::memory::Mapping` 提供检查式复制访问，具有显式共享解除映射状态 |
| 传递描述符或等待 Linux 就绪 | `std::os::linux::ipc` 提供 Unix 消息、eventfd/timerfd、poll 和 epoll |
| 使用 Linux 描述符操作 | `std::os::linux::fd` 提供显式共享所有权、标量/向量 I/O、字节路径、元数据、锁和相对目录操作 |
| 编码 Linux amd64 原生记录 | `std::os::linux::abi` 提供检查式字节编解码及 `Iovecs`/`Message` 构建器 |
| 将 GoML 整数或结构体视为内核指针/布局 | 在 Linux amd64 上使用 `syscall::Arg::Buffer` 和 `syscall::Pointer` 构建同步字节缓冲区图；显式编码原生 ABI，检查 `SyscallResult.errno` 或 `into_result()` |
