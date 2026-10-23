<a id="control-flow-and-copy-optimization-validation"></a>

# 控制流与复制优化验证

[English](performance-validation.md) | [简体中文](performance-validation.zh-CN.md)

本次改动处理 ESP32 模拟器调查中的前三项编译器工作：共享控制流及代码大小回归、字段/数组复制消除，以及可选的性能解释和 PGO 构建记录。浏览器资源加载和完整的跨平台性能回归服务属于其他工作。

<a id="reproduce-the-control-flow-growth"></a>

## 复现控制流膨胀

核心模式是一连串条件式提前退出，随后接一个共享续延：

```goml
fn check(fail: bool) -> Result[(), string] {
    if fail { Err("failure") } else { Ok(()) }
}

fn work(active: bool, fail: bool) -> Result[isize, string] {
    if active { check(fail)?; }
    if active { check(fail)?; }
    println("shared-tail-marker");
    Ok(7)
}

fn main() -> () {
    let _ = work(true, false);
}
```

重复条件行即可暴露膨胀，无需 `defer`、泛型或大型被调用函数。显式 `if fail { return Err("failure") }` 也会产生同样模式。Core、Mono、Lift 和 ANF 都只保留一份尾部。旧的 ANF 到 Go 降级会将续延复制到两个分支，因为终止分支阻碍了共同目标识别。这发生在 Go 编译器运行之前；禁用下游 Go 内联不会消除这种生成源码膨胀。

后端现在区分终止路径与不同的跳转目标。共同续延在分支后合并。当内联返回跳向另一个续延时，嵌套的带标签区域共享剩余连接点主体，无需 `goto` 或新增函数调用。连接点参数仍处于所属作用域中。DCE 在这些区域中传递 break 和 continue 目标的活跃性，确保跳转后使用的赋值不会被删除。

[`control_size_test.goml`](../gomlc/go_backend/control_size_test.goml) 检查 2、4、8 和 16 个条件，要求静态尾部仅出现一次，限制生成字节数上限，并验证输入翻倍时增长不到三倍。它还覆盖 `match` 和 `select`。这些代码大小断言是确定性的，与 CPU 计时无关。[`shared_join_arguments_test.goml`](../gomlc/go_backend/shared_join_arguments_test.goml) 通过编译并运行生成的 Go，验证多个元组/结构体跳转参数。

[`project112_shared_continuations`](../gomlc/testdata/module/project112_shared_continuations/main.goml) 检查 96 种错误/清理组合、独立的提前返回、复合分支值、当前/外层循环跳转及 panic 展开期间的清理。相同断言在固定的 stage0 和修改后的编译器上运行。

所提供复现路径中保留的 ESP32 生成 Go 输入，在此次验证中为 46,264,513 字节，使用 Go 1.26.0 并禁用 PGO 可复现 `prepwrite: bad off=1073741824`。这是完整的下游复现输入，不是最小化的 Go 编译器缺陷。使用共享续延重新构建保留的、未拆分的 GoML 源码可以成功。这避开了病态输入，并不声称修复了 Go 编译器的诊断。

<a id="reproduce-the-field-read-measurement"></a>

## 复现字段读取测量

确保 `PATH` 中有 Go 1.26，运行 `just make` 后执行：

```sh
python3 tools/performance/projection.py --samples 6
```

基线默认使用 `stage0/bin/gomlc`，候选版本默认使用 `stage2/bin/gomlc`。后续比较可显式传入 `--baseline` 和 `--candidate` 路径。公开源码通过嵌套结构体读取 32 个整数，每次基准调用的总和必须为 528。工具先构建两个版本，采集汇编、Go 逃逸诊断和映射到源码的性能解释，再开始计时。随后执行一次不计入结果的预热，并交替安排基线/候选顺序，将所有原始样本和分配计数保存到 `_artifact/performance-projection`。

通过 `performance.json` 找到源码声明及其生成名称，然后检查 `baseline/generated.go`、`candidate/generated.go` 和各自的 `assembly.txt`。旧生成代码会复制嵌套结构体和数组字段；数组索引还会按值将数组传给辅助函数。不可变字段投影现在可以指向现有局部存储，动态 `array.get` 调用降级为 Go 索引。常量索引在必要时仍保留辅助函数，以免 Go 提前拒绝常量越界索引，使程序无法产生语义要求的运行时 panic。

本次包含旧 `4b2498c84` 分支的不可变具名结构体投影优化，并增加定长数组读取和直接动态索引。这些内部指针不引入源码语言的所有权规则。修改或连接点参数重新赋值会阻止字段共享。`Ref`/`Vec` 读取、返回值和闭包捕获保留值快照。[`project111_projection_value_semantics`](../gomlc/testdata/module/project111_projection_value_semantics/main.goml) 覆盖这些边界、嵌套结构、CPU/总线/外设交互、索引中的副作用，以及常量/动态/结果被丢弃的越界读取。相关修复还让指针参数分析检查每个 `select` 分支中的赋值。

<a id="explanations-and-pgo-evidence"></a>

## 性能解释与 PGO 证据

命令、带版本 JSON 格式、剖析权重规则和局限见[性能解释与可复现的 PGO 构建](performance.zh-CN.md)。报告将成本映射到所属 GoML 声明并记录静态事实，不声称提供表达式级 pprof 映射、最终 Go 逃逸决策或实测剖析覆盖率。

公开 PGO 示例使用确定性解释器步数和校验和，采集原生 CPU 性能剖析数据，构建原生产物及可选 WASM 产物，并保留顺序平衡的原始计时样本。原生和 WASM 构建记录使用构建时哈希与工具链查询，不假设 `go version -m` 能检查所有产物。剖析文件被成功接受，与热点函数匹配及性能提升是不同的事实。

计时工具应在编译器构建和测试结束后运行。构建时长对缓存敏感；原生微基准提升不能证明浏览器或完整模拟器获益。脚本同时报告退化和提升，不根据单次运行声称统计显著性。

<a id="recorded-validation"></a>

## 已记录的验证

2026-10-23 的运行使用 Linux amd64、Intel Core i7-13700K、Go 1.26.0 和 Node 22.14.0。原生计时样本固定到 CPU 0。测量前编译器/测试任务均已结束，编辑器后台服务仍在运行。每项运行时比较都排除一次预热，并保留六轮交替顺序测量。

`just update-golden` 和 `just ci` 均通过，包括 1,306 项编译器测试、407 项驱动器/标准库测试、stage3 固定点验证及打包检查。SIMD 测试同时验证结果与标量实现等价，并保留预期的原生内核。

对于条件错误复现程序，保留的基线和修改后的编译器生成结果如下：

| 条件数 | 基线 Go 字节数 | 修改后 Go 字节数 | 基线静态尾部数 | 修改后静态尾部数 |
| ---: | ---: | ---: | ---: | ---: |
| 1 | 6,360 | 5,763 | 2 | 1 |
| 2 | 9,356 | 6,641 | 4 | 1 |
| 4 | 31,906 | 8,407 | 16 | 1 |
| 8 | 620,774 | 11,969 | 256 | 1 |
| 12 | 8,694,441 | 14,219 | 4,096 | 1 |

对于 12 个条件的生成文件，单次 `go build -pgo=off` 观测在修改前为 1.00 秒 / 299,364 KiB 峰值 RSS，修改后为 0.08 秒 / 57,200 KiB。可执行文件大小分别为 2,221,511 和 2,154,645 字节，两者输出相同。这些测量使用现有 Go 依赖缓存，不是统计意义上的构建时间比较。原始结果和生成源码位于 `_artifact/esp32emu-implementation/scale-final`。

保留的未拆分 ESP32 模块无需提取辅助函数即可成功重建。生成的 Go 从保留文件的 46,264,513 字节降至 9,185,536 字节；`run` 从 32,885,107 降至 360,840 字节，直接静态 `Machine.step` 调用点从 1,008 降至 1。较早项目报告中的约 54.35 MB 是历史测量，不是此处测得的保留文件大小。保留文件指纹为 `e867ea3c9b826ba71d48bdb8f474002074559e669e82aaf153389448f50a8eaf`。

修改后编译器的完整构建及 `--help` 单次观测耗时 13.90 秒，峰值 RSS 为 813,800 KiB。原始生成文件仍因 `prepwrite` 失败，无法提供运行时基线。与较早一次成功的中间构建相比，`hello` 和 `pwm` 示例在 10,000 虚拟周期下产生相同寄存器转储及预期的步数预算错误：分别退休 324 和 10,000 条指令。这是功能冒烟检查，不是模拟器吞吐量测量。`_artifact/esp32emu-implementation/esp32emu-final.json` 记录输入、哈希、命令和局限。

公开字段读取基准在修改前为 61.77 ns/次调用，修改后为 8.8585 ns/次调用，两者均为 0 B/次调用、0 次分配/次调用。生成 Go 分别为 5,841 和 5,531 字节。报告定位 `main.goml:11:4` 的 `read` 和 `main.goml:19:4` 的 `project`，各包含一次借用字段投影。Go 1.26.0 逃逸诊断表明它们的指针参数不逃逸。保留的汇编显示，优化后 `project` 中的聚合值复制被直接索引加载替代，校验和仍为 528。约 6.97 倍的吞吐量变化仅适用于此小型读取基准。原始样本、编译器哈希、生成 Go、下游诊断、汇编和报告位于 `_artifact/performance-projection`。

公开 PGO 示例特意如实报告了一次观测到的退化：

| 指标 | PGO 关闭 | 提供原生剖析文件 |
| --- | ---: | ---: |
| 原生工作时长中位数，2,000,000 个解释器步骤 | 131.33 ms | 214.00 ms |
| 原生可执行文件字节数 | 3,291,200 | 3,291,574 |
| WASM 字节数 | 3,772,687 | 3,772,274 |

原生吞吐量下降 38.63%；六组配对速度比均在 0.6126 到 0.6180 之间。两个版本在工作区间内均测得零分配。训练使用 40,000,000 个解释器步骤。四份构建记录包含相同的生成源码哈希，两项提供剖析文件的构建均报告成功接受该文件。剖析覆盖率仍未测量；这次下游 PGO 退化的原因尚未确定。

原生、普通 WASM 和 PGO WASM 在 2,000,000 个解释器步骤下均产生校验和 `5556bcccd558a435`。WASM 在 Node 中执行仅用于正确性验证，其原始耗时不纳入性能比较。证据保留在 `_artifact/performance-pgo/results.json`、`samples.json`、`wasm-validation.json` 和四份 `.build.json` 文件中。此示例展示采集、应用、目标验证及退化报告流程，不承诺加速。

<a id="ref-snapshot-narrowing-2026-10-24"></a>

## 缩小 Ref 快照（2026-10-24）

下一项优化在最终 Go 函数裁剪前，缩小相邻且仅使用一次的聚合值快照。它识别只读取字段的 getter 和经过验证的转发包装函数，然后合并紧接着执行的字段和数组投影。例如，`reference.get().nested.flag` 可以直接读取 `reference.value.nested.flag`。中间借用字段指针如果仅用于下一次值读取，也可以被消除。保留的快照、返回的聚合值、闭包捕获、逃逸地址及中间副作用仍遵循值语义。如果索引表达式修改引用，读取仍必须使用修改前捕获的聚合值。这是局部转换，不引入通用别名分析或纯度标注契约。

公开基准位于 [`tools/performance/ref-projection`](../tools/performance/ref-projection)。重建前保留一个完整的旧工具链前缀目录，然后将其中的编译器与候选版本比较：

```sh
python3 tools/performance/projection.py --case ref \
  --baseline /path/to/previous-prefix/bin/gomlc \
  --candidate stage2/bin/gomlc \
  --output _artifact/performance-ref-projection --samples 6
```

`value` 用例仍为默认选择。两个用例遵循相同测量流程：采样前完成所有构建，每个可执行文件执行一次不计入结果的预热，在同一 CPU 上固定运行六轮交替顺序测量，每次测量调用的校验和均为 528，并记录生成 Go、汇编、逃逸诊断、编译器哈希及原始分配/计时样本。`ref` 用例通过 `Ref[Large]` 而非值参数读取并求和 32 个整数。

在 Linux amd64、Go 1.26.0、禁用 PGO 的条件下，`b36644100` 的编译器测得 113.7 ns/次调用，候选版本测得 8.985 ns/次调用，吞吐量约为原来的 12.65 倍。两者均测得 0 B/次调用、0 次分配/次调用。基线汇编在索引循环内复制聚合值；候选版本直接执行索引加载，该循环内没有聚合值复制。结果保留在 `_artifact/performance-ref-projection`。这些是单机微基准结果，不是完整模拟器或浏览器吞吐量。

[`project113_ref_projection`](../gomlc/testdata/module/project113_ref_projection/main.goml) 覆盖嵌套字段、数组读写、保留及返回的快照、捕获的共享句柄、索引中的副作用、循环/match/select 路径，以及结果被丢弃或被使用的越界读取。其断言在 stage0 和候选版本上均通过。候选版本生成的 Go 也能构建为 `js/wasm`，并在 Node 中通过相同断言；这是功能检查，不是 WASM 性能测量。

`just update-golden` 和 `just ci` 在包含相同实现、测试源码及生成快照的隔离工作树中通过。验证包含 1,314 项编译器测试、407 项驱动器/标准库测试、stage3 固定点验证、扩展打包及发布冒烟测试。使用隔离工作树是为了避免替换另一个模拟器测试任务正在使用的 stage2 可执行文件。
