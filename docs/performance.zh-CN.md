<a id="performance-explanations-and-reproducible-pgo-builds"></a>

# 性能解释与可复现的 PGO 构建

[English](performance.md) | [简体中文](performance.zh-CN.md)

GoML 经 ANF 编译为 Go。编译器事实与下游决策是不同的证据：生成的复制可能被消除，生成的闭包可能留在栈上，语法中的索引操作也可能不再需要边界检查。静态调用数量不是执行次数。

<a id="explain-a-build"></a>

## 解释构建结果

在 GoML 模块中运行：

```sh
goml build --explain-performance
goml build --explain-performance --format json > performance.json
```

即使包产物已缓存，此选项也会强制链接，且不会改变普通构建生成的 Go。默认构建不收集或输出报告。每次链接写入 `goml_generated.go.performance.json`；对于含多个可执行文件的模块，JSON 标准输出将报告包装为 `{"version":1,"kind":"goml-performance-build","reports":[...]}`。构建进度和诊断仍写入标准错误。

每份报告包含 `version: 1`、`kind: "goml-performance"`、编译器版本、生成源码的 SHA-256，以及 `functions`、`sources`、`events` 和 `limitations`。各类别标明产生它们的 `stage`、解释其 `reason`，并标记 `certainty`。`source_precision` 为 `enclosing-declaration`：源码条目包含规范包名、相对于包的文件名、所属声明的行列位置及字节范围。这能定位源码函数，但不一定能定位产生开销的具体表达式。生成函数名及其来源便于检查生成的 Go，但它们并不是跨编译器版本或不相关声明变更的稳定 ABI。

报告包含：

- 每个函数及源码声明的生成字节数。统计格式化后的函数声明，不含方法包装、全局变量、类型和分隔符。共享适配器可能使源码汇总范围重叠。
- 最终裁剪前的具体特化数量。每个普通单态函数计为一个；归属于同一声明的入口适配器或其他具体函数会增加总数。这既不是不同类型实参列表的数量，也不是最终仍被使用的目标函数数量。
- ANF 内联决策，包括是否接受、原因、候选成本、成本上限和调用方剩余预算。目前上限为 16 或 32 个成本单位，调用方初始预算为 48；没有模块级内联预算。
- 部分后端静态计数：聚合值绑定、借用字段投影、索引表达式、已识别的函数值或虚表调用、函数字面量及共享续延区域。共享续延统计生成的结构化退出区域，不统计源码循环或运行时迭代。这些计数并非全部操作的完整清单。

Go 逃逸分析、Go 内联、保留的边界检查和运行时分配均明确标为未测量。`borrowed-projection` 描述内部代码生成选择，不是源码语言的所有权规则。下游决策应以记录中确切 Go 工具链的诊断或实测分配为依据。报告不会将生成的对象直接认定为分配，也不会根据未经验证的成本推导优化建议。

<a id="field-and-array-copy-elimination"></a>

## 消除字段与数组复制

GoML 采用值语义和共享句柄，没有 Rust 式独占借用或生命周期语法。当局部根值和绑定未被赋值、也未复用为连接点参数时，大型具名结构体和定长数组的字段读取可以使用生成的 Go 指针。只读索引随后可直接访问原始不可变存储，无需复制包含该元素的数组。

赋值、返回值和闭包环境仍保留普通值快照。`Ref.get()` 和 `Vec` 查找会先生成快照，再执行字段投影。内嵌的 `Ref` 和 `Vec` 句柄仍共享存储。常量数组索引保留越界 panic 语义。仅凭生成的指针语法不能断定零分配；应检查报告并测量实际工作负载。

<a id="pgo-input-and-build-records"></a>

## PGO 输入与构建记录

```sh
goml build --pgo off
goml build --pgo profiles/representative.pgo
goml run --pgo profiles/representative.pgo -- application-arguments
goml build --pgo auto
GOOS=js GOARCH=wasm CGO_ENABLED=0 goml build --pgo profiles/representative.pgo
```

`--pgo off` 显式禁用 Go PGO。路径相对于 GoML 模块根目录解析。`--pgo auto` 选择模块根目录的 `default.pgo`，不存在时显式禁用 PGO，避免依赖生成文件所在的产物目录。此选项会传给实际的 Go 构建，包括原生 SIMD 编译。不提供 `--pgo` 时，保留现有 Go 环境行为，包括 `GOFLAGS` 和 Go 自身的自动选择。

每次显式指定 `--pgo` 都会写入 `<executable>.build.json`。`--build-record` 也会启用记录，同时保留继承的 PGO 配置。要获得输入完全明确的普通基线，请使用 `--pgo off`；继承的 `GOFLAGS` 和隐式 Go 性能剖析文件选择会被记录，但不会据此解析剖析文件来源。

带版本的 JSON 记录包含：

- 驱动器与编译器版本、编译器路径及 SHA-256。
- 目标 Go 环境与 Go 版本、命令、参数、模块目录及实际 PGO 选择。
- 生成 Go 的路径及 SHA-256、显式性能剖析文件的路径及 SHA-256，以及输出二进制的 SHA-256。
- 构建是否成功、错误，以及成功的 Go 构建是否接受了显式性能剖析文件。
- 局限，包括未测量剖析匹配情况，以及生成源码哈希在可选 SIMD 提取之前计算。

记录标识生成源码和所选性能剖析文件，但不是原生桥接代码、传递 Go 依赖、环境变量或外部资源的完整锁文件。复现构建时应另行保存这些输入。原生和 WASM 均使用构建时证据，不依赖 `go version -m` 读取可执行文件构建信息。性能剖析文件缺失会在调用构建前失败；文件损坏则由 Go 的剖析解析器报告失败。Go 构建失败时，以不含产物哈希的失败记录替换先前成功记录。不记录信息的普通重建会删除旧记录。哈希检查也会拒绝构建期间的源码或剖析文件变化。失败后可能仍有旧二进制文件；请查看 `success` 和产物哈希，不要仅判断二进制文件是否存在。

`profile_read: true` 表示成功的 Go 构建接受了显式性能剖析文件，包括复用 Go 按内容寻址的构建缓存。这不能证明应用热点已匹配，也不能证明进行了优化。`profile_match: "not-measured"` 是有意的：无关但有效的性能剖析文件也能成功构建。此格式尚不测量函数/边覆盖率，也不提供兼容性阈值。路径标识输入，不标识训练工作负载；应在剖析文件旁保留工作负载描述和采集命令。

Go CPU 性能剖析文件可以跨操作系统/架构配置使用，但目标特定代码和工作负载差异会影响适用性。支持将原生训练结果作为 WASM 的剖析输入，但不保证匹配覆盖率或浏览器加速。比较时应测量目标工作负载，并固定工具链版本。参见 [Go PGO 文档](https://go.dev/doc/pgo)。

<a id="collect-merge-build-and-measure"></a>

## 采集、合并、构建与测量

GoML 不引入新的性能剖析文件格式。可通过显式原生桥接调用 Go 的 `runtime/pprof` 采集 CPU 性能剖析数据，或使用应用现有的剖析支持。应保存生成代码、构建记录、工作负载参数、采集时长、工具链和机器信息。

使用 Go 工具合并已有 CPU 性能剖析文件：

```sh
go tool pprof -proto cpu.pgo pwm.pgo rtc.pgo > combined.pgo
goml build --pgo combined.pgo
```

合并会累加样本，不会让每个工作负载获得相同权重。较长的 SDK 工作负载可能压过较短的合成剖析数据。应记录输入哈希、原始样本权重，以及在此命令之外执行的显式归一化或加权。不要将未经归一化的合并称为“等权工作负载”。参见 [Go 的合并指南](https://go.dev/doc/pgo#merging-profiles)。

公开的无依赖解释器示例和测量入口位于 [`tools/performance`](../tools/performance)。确保 `PATH` 中有 Go 1.26+，从仓库根目录运行：

```sh
python3 tools/performance/pgo.py --samples 6 --wasm
```

脚本将示例复制到 `_artifact/performance-pgo/source`，构建基线，采集原生训练剖析数据，再构建 PGO 可执行文件。可选 WASM 构建也在测量前完成。脚本为每个可执行文件执行一次不计入结果的原生预热。指定 `--wasm` 时，随后使用 Node 和所选 Go 工具链的 `lib/wasm/wasm_exec_node.js`，以相同迭代次数分别运行一次普通和 PGO WASM 产物。它们的解释器步数和校验和必须与两次原生预热一致。所有功能检查完成后，脚本才交替执行原生普通/PGO 和 PGO/普通测量轮次。每个进程报告解释器步数、确定性校验和、工作时长、分配次数和分配字节数，功能结果必须一致。`samples.json` 按顺序保存原始测量，包括预热和进程墙钟时间；`results.json` 还记录工具链/机器信息、剖析来源、构建时长、产物大小、记录路径/哈希、中位数及配对比率。`wasm-validation.json` 和报告中的 `wasm_validation` 字段保留 Node 与 Go 路径/版本、执行器/运行时脚本路径及哈希、原生参考结果、命令和 WASM 程序输出。Node 工作时长不纳入性能比较。

Node 执行器接收一个仅含 `TMPDIR` 的小型环境，并将其记录下来，以免向 Go WASM 有限的启动参数/环境区域传入仅用于构建的变量。此示例是合成工作负载，不是浏览器基准或完整编译器回归套件。原生预热不能代替浏览器 JIT 预热。可选 WASM 检查验证构建、记录和 Node 中的功能输出，不测量浏览器性能。`--wasm` 要求 `PATH` 中有 Node，并使用所选 Go 工具链附带的执行器。工作时长与进程墙钟时长须区分，解释器步数也不是硬件 MHz。应在其余工作基本空闲的机器上测量，重复运行未修改的构建来评估噪声，并检查分布和配对比率。工具不推断显著性检验或硬性回归阈值；吞吐量下降也会正常报告。

具体控制流复现、生成大小对比和本次改动的测量见[验证结果](performance-validation.zh-CN.md)。
