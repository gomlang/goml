<a id="compile-time-evaluation-architecture"></a>

# 编译期求值架构

[English](comptime.md) | [简体中文](comptime.zh-CN.md)

GoML 编译期求值是一个带类型的前端阶段：

```text
lexer → parser → CST → AST → HIR → TAST
                                      │
                                      ├─ capability validation
                                      ├─ CTIR lowering
dependency interfaces and CTIR ───────┤
                                      ├─ CTIR verification
                                      ├─ deterministic evaluation
                                      └─ value reification
                                                │
                                                ▼
                                      finalized TAST
                                                │
                           Core → Mono → Lift → ANF → Go
```

阶段边界有三项不变量：

- `#[comptime]` 函数体会在定义它的包中验证，即使当前没有调用点使用它。
- 每个导入的 CTIR 单元都被视为不可信数据，在进入依赖环境前必须验证。
- 最终 TAST 不含 `Comptime` 表达式。Core 降级再次检查此不变量，因此 Core 及后续所有 IR 均独立于 CTFE。

CTIR 是带类型、按树遍历的表示。它包含确定性值、局部状态、结构化控制流、模式匹配、定长数组和内建范围迭代、直接及间接调用、闭包、求值器拥有的引用和向量、经过检查的动态分派、确定性字符串和整数格式化内建函数，以及 `compile_error`。它不含运行时钩子、宿主 I/O、通道或并发。闭包共享捕获的局部单元；每次调用获得新的参数和局部单元。可变值和函数值归属于一次求值，不能逃逸到运行时常量中。动态构造要求存在具体的本地实现，且其方法标记为 `#[comptime]`；调用时检查选中方法的签名。动态方法表覆盖对象规范化后的每个约束，以声明 trait 和方法名称共同标识方法；限定调用在求值过程中保留这一标识。trait 方法值降低为闭包，闭包体复用普通调用的降低流程，保留能力验证和限定动态分派。编译期方法入口和静态目标选择包含完整 trait 应用，因此不同类型或常量实参不会共用同一个实现。求值前会验证局部槽位类型、调用签名、控制流目标、源码来源 ID、结构化值类型和目标兼容性。

求值使用语义执行预算（fuel）、调用深度、临时节点数、临时内存和最终结果大小限制，从不使用实际时间截止限制。仅在纯求值内对直接调用结果进行记忆化；创建闭包、引用或向量会禁用该次求值的记忆化。整数值存储符号性、位宽和原始位；CTIR 目标规格决定 `isize` 和 `usize` 的位宽，并参与语义哈希。

整数 `to_string` 降级为带类型的 `comptime.to_string.iN` 或 `comptime.to_string.uN` 内建函数，位宽为 8、16、32 或 64。验证器检查参数位宽和符号性；求值按目标整数语义格式化原始位，包括有符号最小值和无符号最大值。这不会启用 CTFE 中任意 trait 调用。

接口产物导出公共 comptime 入口，以及私有 comptime 辅助函数和常量的可达闭包。公共常量携带规范值。源码来源是调试元数据，使用包相对路径，不参与 CTIR 语义哈希。函数体、局部槽位类型、目标语义、公共值和引用的依赖接口哈希会参与。产物解码器检查格式和语义哈希，随后依赖加载使用完整的导入模块集合运行 CTIR 验证器。

索引降级根据检查后的接收者类型选择操作：定长数组使用 CTIR 索引，向量使用求值器缓冲区读取。接收者和索引的求值顺序保持不变，各操作保留其边界检查。

CTFE 失败是可恢复的编译器诊断，指出失败的源码来源、发起请求的 comptime 位置和编译期调用栈。求值失败或资源耗尽绝不会终止编译器进程。

<a id="programmable-derive-phase"></a>

## 可编程 derive 阶段

可编程 derive 复用已验证的 CTIR，但在消费者侧更早的阶段运行：

```text
dependency interface → verify derive CTIR ──────────────┐
                                                        │
source → parser → CST → AST → resolve derive entry → evaluate handler
                                                        │
                                                        ▼
                                                generated AST items
                                                        │
                                                        ▼
                                          HIR → TAST → ordinary CTFE
```

处理函数在其定义包编译时已完成类型检查并降级到 CTIR。消费包从不执行无类型源码或宿主代码。求值接收不透明的 `DeriveInput` 句柄，必须返回不透明的 `DeriveOutput` 句柄。编译器内建函数通过不透明的 `MetaAttribute`、`MetaSpan` 和 `MetaType` 句柄暴露结构化属性、源码位置和类型形状。表达式、模式、分支、块、方法和声明构建器构造结构化输出项。结果不能包含任意 token 或声明，随后进入正常的 HIR 降级、名称解析、一致性检查、类型检查、单态化和代码生成流程。

公共 `#[comptime_derive]` 入口及私有 derive 辅助函数的可达闭包属于接口 CTIR 语义部分。因此 derive 函数体变化会改变接口哈希。derive 入口不是运行时导出项。接口解码器在求值前验证元类型、内建函数签名、直接调用目标、ID 和语义哈希。

定义处构建器将未限定的 trait、类型和函数名限定到处理函数所在包。显式调用处构建器将名称保留在目标包作用域中。生成的局部绑定使用处理函数选择的名称；`derive_fresh_name` 提供无冲突的编译器名称。生成节点最初使用发起请求的 derive 属性作为诊断来源。`MetaSpan` 句柄选择字段类型和属性参数的范围；`derive_set_span` 改变后续构建器的来源，`derive_error_span`、`derive_error_at` 和 `derive_error_at_argument` 选择精确诊断位置。

derive 求值使用普通的执行预算、深度、值节点数和临时内存限制，此外最多允许 100,000 次元数据或语法构建器操作。derive 调用不进行记忆化，因为其不透明 arena 句柄只在本次求值内有效。

查询层按源文件保留格式化后的展开 AST。`gomlc run-single --dump-expanded-ast` 和 LSP 的 `goml/expandedDerive` 请求暴露相同的展开边界，不绕过正常的 derive 求值或诊断。

derive 输出起始于一个 trait 实现或一个固有实现。`derive_output_inherent(input)` 保留目标的泛型参数和标识。`derive_output_add_method` 将固有方法保持为私有，`derive_output_add_public_method` 将固有方法标记为公共，并对在 trait 实现上的使用给出诊断。两者均进入普通的可见性和类型检查流水线。输出可以添加关联类型和常量、设置 trait 实参和等式谓词、合并更多 impl，并生成私有辅助函数、结构体和类型别名。这些声明经过正常的名称解析和类型检查。

同包处理函数在展开前准备。受限 AST 遍历保留声明和编译期函数，移除 derive 请求和运行时函数体，并构建经过验证的本地 CTIR。最终遍历使用这些本地处理函数展开原始 AST，并检查完整包。本地处理函数不能依赖其自身展开尚未生成的声明。私有命名处理函数保持私有；导入的 derive 仍要求公共入口。项目编译、单文件编译和查询分析共用相同的准备过程。

结构化属性访问对括号平衡的参数列表进行 token 化，保留字面量种类、名称、嵌套类型和表达式语法。原始字符串使用普通解析器解码。类型和表达式访问器在派生文件作用域中恰好解析一个值。原始属性文本仍可访问。生成的闭包使用带类型的参数列表；`meta_expr_type_member` 携带所属类型和成员的类型实参，不解析调用方的同名类型。

内建 `ToString` derive 在 GoML derive 处理函数中实现[枚举变体模板](language/metaprogramming.zh-CN.md#attributes-and-derive)。它通过结构化属性访问读取 `#[to_string("...")]`，在占位符的源码范围验证它们，并为引用的字段使用新局部变量构造普通 match 分支。只有模板使用的字段会贡献 `ToString` 约束；无模板变体保留默认格式及其字段约束。编译器提供元数据和语法构建器，而不自行解释模板语法。
