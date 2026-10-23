<a id="go-interoperability"></a>

# Go 互操作

[English](bind-go.md) | [简体中文](bind-go.zh-CN.md)

[GoML 语言指南](../goml.zh-CN.md)

GoML 通过显式外部声明调用 Go。编译器检查声明的 GoML 类型，并信任绑定作者提供的签名。最终 Go 构建检查原生符号是否存在、可赋值性、方法集和泛型约束。`goml check` 和编辑器分析不加载 Go 包元数据。

```goml
use std::ffi;

#[go_type("os", "File")]
extern type File;

#[go_ffi("os", "Open")]
extern fn open_file(name: ffi::String) -> (ffi::Ptr[File], ffi::Error);

#[go_method("Close")]
extern fn close_file(file: ffi::Ptr[File]) -> ffi::Error;
```

外部类型是不透明的。其导入路径、Go 名称和声明的类型实参共同标识该类型；GoML 不检查其字段、不规范化 Go 别名、不发现方法，也不推导接口结构。请显式绑定函数和方法，或在原生 API 需要转换时提供一个小型 Go 适配层。`#[go_interface]` 仍可使用，但必须显式提供原始类型、trait 签名和方法映射；最终 Go 编译会对照真实接口检查生成的适配器。

本地 Go 适配层和 Go 依赖使用模块根目录的 `go.mod`。普通项目构建禁用 Go 工作区模式，根据是否存在 `go.mod` 选择模块模式，其他依赖和工具链行为沿用用户的 Go 环境。Go 库导出和原生 SIMD 构建仍采用只读、离线的模块处理方式，要求提前准备依赖。Unicode 有效性、所有权、可变性和并发保证仍属于显式适配器契约；Go 编译成功并不能证明这些性质。

未使用的外部声明可能不会出现在生成的 Go 中，因此最终构建不一定检查它们。构建生成的程序或库时才会检查原生 Go 源码的变化；`goml check` 始终使用声明的 GoML 签名。

完整边界规则见 [Go FFI](#go-ffi)，显式指针和方法绑定见[文件读取示例](../../goml/testdata/ffi/file-read/README.zh-CN.md)，反向调用见[导出 Go 库](#exporting-a-go-library)。

<a id="go-ffi"></a>

## Go FFI

可运行的[文件读取夹具](../../goml/testdata/ffi/file-read/README.zh-CN.md) 综合展示原始指针、方法、共享字节缓冲区、部分读取、错误身份、显式 Close 和 Go↔GoML 结果传输。

类型化 Go FFI 将顶层 GoML 声明绑定到导出的包级 Go 函数：

```goml
#[go_ffi("strings", "ToUpper")]
extern fn to_upper(value: string) -> string;

#[go_ffi("strings", "Cut")]
extern fn cut(value: string, separator: string) -> (string, string, bool);

fn example() -> string {
    let (before, after, found) = cut("left:right", ":");
    if found {
        to_upper(before + after)
    } else {
        ""
    }
}
```

属性第一个实参是 Go 导入路径，第二个是导出的 ASCII Go 标识符。GoML 函数名是本地名，可不同于 Go 符号。在 `extern fn` 前加 `pub` 可通过其他 GoML 包暴露绑定；接口和 Core 产物保留 Go 导入路径及符号。

使用 `#[go_type("time", "Duration")] pub extern type Duration;` 绑定不透明 Go 类型。GoML 名可不同于 Go 对象名。声明不能有主体或 `= Type` 目标。导入路径、Go 名和类型实参共同标识类型：`Duration` 不同于 `i64`，而两个 GoML 包绑定的 `time.Duration` 共享类型。值可通过普通 GoML 函数和显式类型的外部调用传递；生成的 Go 使用原始 Go 类型。编译器不检查外部字段、不规范化 Go 别名，也不发现方法和接口。不可访问外部字段或构造外部结构体；请使用显式函数/方法绑定或 Go shim。

```goml
#[go_type("time", "Duration")]
extern type Duration;

#[go_ffi("time", "Sleep")]
extern fn sleep(duration: Duration) -> ();

fn wait(duration: Duration) -> () {
    sleep(duration)
}
```

对于 Go 的 `type Point struct { X, Y int64 }`，从其 Go 包绑定 `NewPoint`、`PointX` 等显式辅助函数：

```goml
#[go_type("example.com/geometry", "Point")]
extern type Point;

#[go_ffi("example.com/geometry", "NewPoint")]
extern fn point(x: i64, y: i64) -> Point;

#[go_ffi("example.com/geometry", "PointX")]
extern fn point_x(value: Point) -> i64;
```

`use std::ffi;` 暴露 `ffi::Ptr[T]`，它是指向 T 的原始 Go 指针的透明写法。`ffi::null()` 构造类型从上下文推断的 nil 指针，`ffi::is_nil(value)` 检查它。非空指针通过显式 Go 绑定进入。例如，`let empty: ffi::Ptr[Cell] = ffi::null();` 声明 nil 指针。这些函数也可作为普通函数值。指针在函数、闭包或容器间传递不复制所指对象。包含 nil 的 `Option::Some` 仍不同于 `Option::None`。GoML 不能从整数或 `Ref` 构造指针。独立声明的 Go 别名如 `extern type Handle` 仍不透明，不会规范化为 `ffi::Ptr[Cell]`。它不能传给 `ffi::is_nil()` 或接收 `ffi::null()` 的值；请在绑定签名中使用显式指针表示，或为不透明类型绑定 Go 辅助函数。没有 `*T` 源类型语法。

通过 `#[go_method("Method")] extern fn` 绑定 Go 方法；第一个参数为接收者。声明必须单态，且恰好有一个 ABI 属性。方法名必须为导出的 ASCII Go 标识符。最终 Go 构建检查接收者实际方法集：仅指针方法要求适当的指针接收者，值方法也可通过指针调用。不会对复制的值隐式取地址。接口和提升方法遵循相同 Go 规则。绑定可公开、跨包导入或作为普通函数值传递，且始终仅限运行时，即使从产物加载也一样。

```goml
use std::ffi;

#[go_type("example.com/shim", "Cell")]
extern type Cell;

#[go_method("Read")]
extern fn read(cell: ffi::Ptr[Cell]) -> i64;

#[go_method("Write")]
extern fn write(cell: ffi::Ptr[Cell], value: i64) -> ();
```

这些绑定生成 `(*shim.Cell).Read(cell)` 等 Go 方法表达式。nil 接收者原样传给 Go 方法，包括该方法定义的 panic 或特殊 nil 行为。方法不会自动分配对象、关闭资源或将错误转换为 `Result`。取得非空值仍需显式 Go 绑定。

`ffi::Error` 是原始的可空 Go `error` 接口。`ffi::nil_error()` 构造 nil 接口；`ffi::error_is_nil(value)` 检查接口是否为 nil。持有带类型 nil 指针的接口不是 nil。错误在函数、闭包和产物间保留原对象及错误链。原始 `(T, error)` 返回仍是元组；包含 nil 错误的 `Option::Some` 仍不同于 `None`。

`ffi::NonNilError::from_error(error)` 对 nil 接口返回 `None`，对非 nil 接口（含带类型 nil 错误）返回 `Some`。其私有存储不能直接构造；`as_error()` 返回原错误。因此 `Result[T, ffi::NonNilError]` 不能接受 `Err(ffi::nil_error())`。此包装器不调用 `Error()`，也不检查具体错误载荷。

`ffi::Outcome::new(value, error)` 在公开 `value` 和 `error` 字段中保留两个结果，包括失败时同时返回的部分值。显式命名的 `into_result_discarding_value_on_error()` 适配器返回 `Result[T, ffi::NonNilError]`：仅成功时保留值，失败时丢弃值。此转换不会自动发生。例如：

```goml
let (count, error) = read(buffer);
let outcome = ffi::Outcome::new(count, error);
let partial_count = outcome.value;
let result = outcome.into_result_discarding_value_on_error();
```

反向转换也需显式进行。`ffi::Outcome::from_raw_result(result, failure_value)` 接受 `Result[T, ffi::Error]`，返回 `Result[ffi::Outcome[T], ffi::BoundaryError]`。`Ok(value)` 产生该值和 nil Go 错误。`Err(error)` 使用调用者提供的失败值并保留错误。若错误是真正的 nil 接口，转换以 `BoundaryError::NilError` 失败，消息为 `Go FFI boundary rejected Err(nil)`。带类型 nil 错误仍是失败。即使输入为 `Ok`，失败值实参也正常求值。

`ffi::Outcome::from_result(result, failure_value)` 接受 `Result[T, ffi::NonNilError]` 并应用相同边界检查。`outcome.into_tuple()` 返回用于显式 Go 调用的精确 `(value, error)` 对；不会规范化部分结果或 nil 值。这些适配器都不臆造默认失败值，也不会隐式将契约错误变成 Go 成功。

`ffi::error_matches(value, target)` 使用 Go `errors.Is`，包括包装链、自定义 `Is` 方法和 Go 的 nil 匹配语义。`NonNilError::matches(target)` 为已检查包装器提供相同操作。匹配比较原始错误对象和错误链，从不比较消息字符串。

`NonNilError::message_bytes()` 调用原始 Go `Error()` 方法，返回其字节的新 `std::bytes::Bytes` 副本。`message()` 返回 `Result[string, std::utf8::Utf8Error]`，接受有效 UTF-8，拒绝无效字节而不替换。访问消息不替换存储的错误。带类型 nil 接口可进入已检查包装器，因此底层 Go 方法的 nil 处理、副作用和 panic 仍由其自身决定。这些 API 不捕获 Go panic，也不隐式将错误转为字符串。

泛型外部声明如 `#[go_type("example.com/shim", "Box")] extern type Box[T];` 接受 `Box[i64]`，以及泛型 GoML 代码中的符号形式 `Box[T]`。GoML 检查声明参数数量，并跨包和产物保留外部身份及实参。不查询 Go 声明、不检查约束、不展开别名。最终 Go 编译检查特化后可达的具体实例，包括实际类型和实参能否作为合法运行时值。未使用的泛型函数可能不产生 Go 实例。查询和 LSP 分析使用同样的声明不透明类型。

外部调用 ABI 支持生成的 Go 表示可直接赋值的值。原始边界类型和适配器详见后续各节：

| GoML 类型 | Go 表示 |
| --- | --- |
| `bool`、数值基本类型、`string` | 对应 Go 基本类型 |
| 带具体实参的外部具名类型 | 原始 Go 包类型，不创建新的定义类型 |
| `ffi::Ptr[T]` | 可空 Go 指针 `*T` |
| `ffi::Error` | 可空 Go `error` 接口 |
| `ffi::String` | 保留任意字节的 Go `string` |
| `ffi::Rune` | 不做标量验证的 Go `rune` / `int32` |
| `ffi::RawSlice[T]` | 可空 Go 切片 `[]T` |
| `ffi::RawMap[K, V]` | 键可比较的可空 Go 映射 `map[K]V` |
| `char` | `rune` / `int32` |
| `byte` | `byte` / `uint8` |
| `[T; N]` | `[N]T` |
| `Slice[T]` | `[]T` |
| `MutSlice[T]` | `[]T` |
| `ffi::Func[F]` | 签名为 F 的可空 Go 函数；空元组/扁平元组结果对应零个/多个 Go 结果 |
| `Channel[T]` | `chan T` |
| `Sender[T]` | `chan<- T` |
| `Receiver[T]` | `<-chan T` |
| 直接 `dyn Marker` 参数或单个返回值 | 通过 trait 对象 `data` 字段传递 Go `any` |
| `()` 返回 | 无结果的 Go 函数 |
| `(A, B, ...)` 返回 | 同顺序的多个 Go 结果 |

元组类型仅支持作为完整返回类型。元组元素及数组、切片、通道元素本身也必须满足 FFI 安全。参数不能是 `()`。

`dyn Marker` 仅支持直接参数或完整的单返回类型。`Marker` 必须严格为空：不能声明泛型参数、谓词、父 trait、关联类型或方法，`dyn` 类型不能有实参、关联类型绑定或附加约束。产物验证在声明函数的包内解析本地未限定标记名，包括私有标记。包装器将对象的 `data` 字段作为 Go `any` 传递，再将返回 Go 值包装回空标记对象。空标记对象暂不支持嵌套在元组、数组、切片或通道内。

声明必须单态。GoML 检查声明的边界形态，并信任显式签名。`goml check` 和编辑器分析不检查原生包。最终 Go 构建根据实际符号、接收者方法集、可赋值性和泛型约束验证生成调用。错误声明在该构建步骤失败；未使用绑定可能不生成，因此可能不检查。即使底层表示匹配，`time.Duration` 这样的 Go 具名类型仍不同于 GoML `i64`。

最终可执行文件构建使用选定的 Go 可执行程序、目标、标志、模块策略和生成调用者目录。Go 检查导入合法性，包括 `internal` 限制。普通构建禁用 Go 工作区模式，并根据模块根是否存在 `go.mod` 选择模块模式；其他 Go 依赖和工具链行为遵循用户 Go 环境。不会额外实施 FFI 专用只读或离线策略。Go 库导出和原生 SIMD 构建保留只读、离线模块处理，请单独准备依赖。Go 构建成功不能证明 Unicode、可变性、生命周期或并发契约。返回空元组的旧绑定可能丢弃 Go 结果；普通调用保留 Go 泛型推断和可变参数调用规则。

此 ABI 不支持 `Vec`、`Ref`、`HashMap`、`Option`、`Result`、用户结构体和枚举、非空 trait 对象、普通 GoML 函数值、嵌套元组、泛型函数声明、隐式回调转换、自动 Go 对象生命周期管理和自动 `error` 转换。适配此类 API 时，编写带导出函数和 FFI 安全参数的小型 Go shim：

```go
package goshim

import "os"

func ReadText(name string) (string, bool) {
    data, err := os.ReadFile(name)
    return string(data), err == nil
}
```

```goml
#[go_ffi("example.com/myapp/goshim", "ReadText")]
extern fn read_text(name: string) -> (string, bool);
```

GoML 模块根包含 `go.mod` 时，`goml build`、`goml run` 和测试链接以模块模式调用 Go，从而可导入本地 shim 包及声明的 Go 模块依赖。Go 工作区模式仍禁用。无 `go.mod` 时保留现有关闭模块的行为。模块模式 Go 构建始终执行，并将依赖和源码新鲜度交给 Go 自身构建缓存，因此会观察到 `go.mod`、`go.sum` 和本地 `.go` shim 的变化。

包函数形式从 GoML 调用 Go；方法调用使用上述显式 `go_method` 形式。动态符号查找和 C ABI 互操作需要独立机制。

<a id="go-interface-wrappers"></a>

### Go 接口包装器

模块编译中，顶层 trait 上的 `go_interface` 生成持有显式绑定的原生 Go 接口并转发方法的包装器：

```goml
use std::ffi;

#[go_type("io", "Reader")]
pub extern type GoReader;

#[go_interface(GoReader, ReaderAdapter, read = "Read")]
pub trait Reader {
    fn read(self: Self, buffer: ffi::RawSlice[u8]) -> (isize, ffi::Error);
}

fn wrap(raw: GoReader) -> ReaderAdapter {
    ReaderAdapter::from_go(raw)
}
```

第一个实参是原始类型路径，可指向导入的外部类型。第二个是未限定包装器名，其可见性随 trait。具名字符串实参将每个声明的 trait 方法映射到恰好一个导出 Go 方法。GoML 检查 trait 映射完整性、重复/未导出名称和声明的原始边界类型。生成类型/辅助名称及保留包装器成员不能与用户声明冲突。最终 Go 构建检查真实原生接口和方法签名；GoML 不发现其方法集。

`ReaderAdapter::from_go(raw)` 将原接口值存入私有字段；`wrapper.into_go()` 返回它。二者保留 nil 接口、带类型 nil 接收者及所引用 Go 对象身份。转发调用保留 Go 接收者行为，包括 nil 处理和 panic。不替换接收者，也不自动关闭。原始 Go 接口仍不同于 `dyn Reader`；用 `wrapper as dyn Reader` 显式转换为 GoML trait。

方法签名使用显式原始边界：文本使用 `ffi::String`/`ffi::Rune`，切片使用 `ffi::RawSlice`。在调用方显式应用复制/共享适配器。所有 Go 结果均保留：`()` 表示无结果，`(isize, ffi::Error)` 同时保留部分计数和原错误。普通 GoML 文本、只读切片、隐式闭包和运行时容器不会自动转换成接口方法边界。

`ReaderAdapter::from_trait(value: dyn Reader)` 创建由生成的 Go 桥接支撑的包装器。用 `ReaderAdapter::from_trait(implementation as dyn Reader).into_go()` 将 GoML 实现传给 Go。类型化转发闭包在创建函数返回后仍保留所提供 trait 对象。Go 调用保留共享捕获状态和每个返回分量；panic 在同一 goroutine 传播。生成的桥接包含编译期 Go 接口满足性断言。结果接口非 nil，即使提供的 trait 对象包装 nil 原始接口也如此。`from_go`、`into_go` 和 `from_trait` 是包装器保留方法名。取消订阅和关闭资源仍由调用者显式负责。

适配器要求原始外部类型和带实例方法的非泛型 GoML trait，不能有父 trait、关联类型或额外方法谓词。显式签名和映射提供生成桥接所需信息。最终 Go 构建检查原始类型是否为可用接口，以及生成桥接是否满足它，包括嵌入方法和签名兼容性。仅约束用途的接口不能用作 Go 运行时值。

<a id="raw-go-function-values"></a>

### 原始 Go 函数值

`ffi::Func[F]` 是可空原始 Go 函数值，不同于签名 F 的 GoML 闭包。例如，`ffi::Func[(i64) -> (i64, ffi::Error)]` 表示 Go `func(int64) (int64, error)`，`ffi::Func[() -> ()]` 表示 `func()`。F 必须解析为参数和结果均 FFI 安全的函数签名；函数不能作原始映射键。

`ffi::func_nil()` 构造签名从上下文推断的 nil，`ffi::func_is_nil(value)` 检查它。Go 返回的值可通过包、泛型函数、容器和捕获它的 GoML 闭包传递，不丢失 Go 闭包环境。包含 nil 函数的 `Option::Some` 仍不同于 `Option::None`。Go 绑定可接收这些原始值并在 Go 中调用。不能用 GoML 调用语法直接调用原始函数，不支持隐式闭包转换或结构体构造。

```goml
use std::ffi;

fn empty_callback() -> ffi::Func[(i64) -> i64] {
    ffi::func_nil()
}
```

`ffi::func_from_closure(value)` 显式将 GoML 函数或闭包转为非 nil 原始 Go 函数。`ffi::func_to_closure(value)` 对 nil 返回 `None`，否则返回可调用 GoML 闭包 `Some(closure)`。两个方向都保留捕获状态；转换同签名的两个闭包时仍保留各自环境。捕获的 `Ref` 单元保持普通共享行为。

```goml
use std::ffi;

fn round_trip() -> bool {
    let callback: ffi::Func[(i64) -> i64] = ffi::func_from_closure(|x| x + 1);
    let Some(call) = ffi::func_to_closure(callback) else {
        return false
    };
    call(41) == 42
}
```

转换生成保留原函数及环境的类型化 Go 方法值。GoML 闭包仍经过闭包提升。空元组和扁平元组结果在 GoML 存储与 Go 的零个/多个结果间显式转换。原始错误和其他受支持原始值保持不变。回调文本值必须使用 `ffi::String` 或 `ffi::Rune`；普通 GoML `string`/`char`（包括数组、切片、通道或原始映射中的）会在适配器验证时被拒绝。请在回调内部做带检查的 Unicode 转换，并显式选择失败结果形态。

适配器不添加同步、全局注册表、反射或隐式 panic 恢复。panic 沿同一 goroutine 传播，包括穿过两个转换方向，Go 和 GoML 的 defer 在栈展开时运行。panic 不会自动成为错误结果。恢复需要显式 `std::panic::catch` 边界，或调用方拥有的、在该 goroutine 中用 defer/recover 的 Go 适配器；二者都不能恢复其他 goroutine 的 panic 或拦截进程退出。`panic::resume` 用内部运行时包装器保留捕获 panic 的诊断信息；Go 调用方在 resume 后 recover 看到的是包装器，而非原载荷动态类型。Go 可在注册回调的 GoML 调用返回后继续保留转换后的回调，在另一 goroutine 调用，或递归重入 GoML。捕获状态通过 Go 函数值保持可达。并发修改需要普通同步；转换不会让 `Ref` 线程安全。

可运行的[回调订阅夹具](../../goml/testdata/ffi/callbacks/README.zh-CN.md) 覆盖显式注销、受控异步分派、并发原子更新、递归重入和注册压力。注销会移除其回调引用；已选中的调用仍可能执行，因此调用方在释放捕获资源前等待相应批次。GC 可达性和注册生命周期是不同概念。夹具及其生成的 Go 程序使用竞态检测器测试。仍不支持 GoLibrary 函数值导出。

可执行文件、测试和库模块链接在生成的 Go 字符串常量 `_goml_source_origins` 中嵌入带版本 JSON 索引，将保留的函数及回调适配器关联到外层 GoML 声明。共享适配器可有多个源码条目。外部调用包装器还标识实际使用的 extern 声明，包括 `go_method`；同一原生目标的未使用别名被排除。索引记录规范符号、源包、包相对源文件和声明名称位置，包括显式定义的 impl 方法；它不是表达式级栈追踪行号表。

索引具有 `version: 1`、`scope: "enclosing-declaration"` 和 `entries` 数组。条目包含 `generated_symbol`、`source_symbol`、`source_package`、`source_file`、`start`、`end`、`line` 和 `column`；回调方法用 `ReceiverType.Apply` 作为生成符号。解析 JSON 前先解码 Go 字符串字面量。

<a id="raw-go-strings"></a>

### 原始 Go 字符串

`std::ffi::String` 表示任意 Go 字符串字节，不同于 GoML Unicode `string`。跨外部调用、包接口、泛型函数和捕获闭包传递时字节不变。不支持隐式转换或直接构造。

| 显式适配器 | 结果与策略 |
| --- | --- |
| `ffi::string_from_bytes(bytes::Bytes)` | `ffi::String`；复制字节，包括无效 UTF-8 和 NUL |
| `ffi::string_from_text(string)` | `ffi::String`；保留已验证文本 |
| `ffi::string_bytes(ffi::String)` | `bytes::Bytes`；返回独立可变字节副本 |
| `ffi::string_text(ffi::String)` | `Result[string, utf8::Utf8Error]`；验证 UTF-8，不替换无效字节 |

```goml
use std::ffi;
use std::bytes;

fn checked_text() -> bool {
    let raw = ffi::string_from_bytes(bytes::Bytes::from_vec(Vec::from_array([255, 0])));
    match ffi::string_text(raw) {
        Ok(_) => false,
        Err(error) => error.valid_up_to() == 0,
    }
}
```

用 GoML `string` 声明的旧外部绑定保留现有受信文本契约。Go 生产者可能返回任意字节时，使用 `ffi::String`。GoLibrary 导出可用 `ffi::String` 选择原始字节策略。普通 GoML `string` 和 `char` 导出仍不受支持；请在边界使用显式带检查适配器。

<a id="raw-go-runes"></a>

### 原始 Go rune

`ffi::Rune` 是 `i32` 的透明别名，与 Go `rune` 一致。原始值可以为负、代理码点或超出 Unicode 标量范围。`ffi::rune_from_char(char) -> ffi::Rune` 保留有效字符。`ffi::rune_char(ffi::Rune) -> Result[char, ffi::BoundaryError]` 检查标量范围，无效输入返回 `InvalidRune(original_value)`，从不替换成替换字符。这些适配器使用普通类型别名和函数语法。

<a id="raw-go-slices"></a>

### 原始 Go 切片

`ffi::RawSlice[T]` 是 `MutSlice[T]` 的别名，直接表示 Go `[]T` 头。头赋值复制长度、容量和底层存储引用，不复制元素。现有带检查索引和子切片方法仍可用。

| 适配器或操作 | 策略 |
| --- | --- |
| `ffi::slice_nil::[T]()` / `ffi::slice_is_nil(value)` | 保留并区分 nil 与非 nil 空切片 |
| `ffi::slice_capacity(value)` | 返回原始头的容量 |
| `ffi::slice_append(value, item)` | 返回 Go append 的新头；保留它才能观察新长度 |
| `ffi::slice_shared(vector)` | 共享可变元素存储，使用 Vec 头的快照 |
| `ffi::slice_from_vec_copy(vector)` / `ffi::slice_copy(value)` | 将元素复制到独立存储 |
| `ffi::slice_to_vec_copy(value)` | 将元素复制到新 Vec |
| `ffi::slice_trusted_readonly(value)` | 在绑定作者承诺下，通过只读 `Slice[T]` 视图共享存储 |

复制为浅复制：复制指针切片会隔离元素槽，但所指对象仍共享。空输入的复制适配器生成非 nil 空分配，不保留 nil 身份或剩余容量。追加可能复用底层存储，也可能分配新底层数组；Go append 和 Vec append 都不更新先前复制的头。Go 写入通过共享底层存储的别名可见。所有适配器都不添加同步。

`trusted_readonly` 是作者对存储适合只读使用的显式承诺，包括协调其他写入者；Go 签名无法证明此承诺。转换不冻结原切片，也不递归冻结引用对象。需要独立元素存储时先复制。Vec 和 HashMap 仍无隐式外部转换。

```goml
use std::ffi;

fn append_header() -> bool {
    let original: ffi::RawSlice[i64] = ffi::slice_nil();
    let updated = ffi::slice_append(original, 7);
    ffi::slice_is_nil(original) && updated.len() == 1
}
```

仍不支持 GoLibrary 切片导出。

<a id="raw-go-maps"></a>

### 原始 Go 映射

`ffi::RawMap[K, V]` 直接表示 Go `map[K]V`，不同于 GoML `HashMap[K, V]`。在赋值、外部调用、包接口、泛型和捕获闭包间保留 nil 和共享存储。对别名的 Go 操作观察到相同条目。

`ffi::map_nil::[K, V]()` 产生 nil `ffi::RawMap[K, V]`。`ffi::map_is_nil(value) -> bool` 区分 nil 和已分配空映射。包含 nil 的 `Option::Some` 仍不同于 `None`。原始类型没有公开结构体构造器，也不能与 `HashMap` 隐式互转；二者的存储、哈希和相等契约不同。原始映射语法使用普通限定泛型别名和函数。

| 操作 | 语义 |
| --- | --- |
| `ffi::map_new::[K, V]()` | 分配非 nil 空原始映射 |
| `ffi::map_len(value)` | 条目数；nil 长度为零 |
| `ffi::map_get(value, key)` | `Option[V]`；区分缺失与存在的零值或 nil 值 |
| `ffi::map_set(value, key, item)` | 使用 Go 相等规则插入或替换；写入 nil 映射保留 Go 的 panic |
| `ffi::map_delete(value, key)` | 存在则删除；缺失键和 nil 映射不做操作 |
| `ffi::map_copy(value)` | 独立条目存储，值为浅复制；保留 nil |

赋值和传参共享条目存储。需要独立条目时用 `map_copy`。副本中的指针值仍引用原对象。复制直接遍历键值对，因此即使 NaN 不能通过相等查找，以 NaN 为键的条目仍保留值。这些操作均不冻结存储、不引入同步，也不转换 HashMap 哈希语义。

```goml
use std::ffi;

#[go_ffi("example.com/myapp/goshim", "Entries")]
extern fn entries() -> ffi::RawMap[i64, i64];

fn available() -> bool {
    !ffi::map_is_nil(entries())
}
```

最终 Go 编译器实施 Go 映射键可比较规则。GoML 检查拒绝已知不可比较键，包括切片、原始映射、包含它们的数组或元组，以及含此类字段的本地结构体；推断出的泛型应用在推断后检查。重复的本地泛型包装器在每个具体实例检查：若 `Key[T]` 直接存储 `T`，则 `Key[Key[RawSlice[i64]]]` 被拒绝；仅存储指向 `T` 指针的包装器仍可比较。递归结构体布局及超过 64 层嵌套结构体实例的结构检查产生可恢复诊断。指向不可比较类型的指针仍是合法键。公开 GoML 结构体接口保留私有字段的可比较性摘要，下游键检查可遵循限制而不暴露字段名或类型。摘要可依赖泛型实参；未知隐藏结构不能证明可比较。产物解码器也拒绝直接不可比较的键形态。外部具名键是不透明的；最终 Go 构建检查实际字段和具体类型实参的可比较性。GoML 分析不导入原生字段结构或约束。模块检查从泛型函数体推断内部可比较性要求，包括经函数值和嵌套闭包到达的要求。包检查跨文件传播，直到签名稳定，设有可恢复的 256 轮限制。导出函数接口保留推断要求，因此下游模块检查可在 Go 构建前拒绝无效键。私有结构键包装器被归约为所需类型参数或关联类型投影。这些谓词没有用户编写的约束语法，也不引入 let 泛化。仅源码依赖分析会在检查调用者前检查依赖体并保留其推断要求，包括未保存源码覆盖。结构体和枚举接口也保留字段或变体载荷中映射的键要求，包括私有存储和别名。这些要求在名义声明间传播，设有可恢复的 256 轮限制，并在具体类型应用时替换。递归类型应用通过声明要求检查，不反复展开布局。默认 trait 方法单独保留推断映射键要求，与方法声明约束分开。使用默认体的实现必须在替换接收者和关联类型后满足这些要求。显式方法覆盖不继承未使用默认体的要求。仍不支持 GoLibrary 映射导出。原始映射不添加同步或自动深复制。

<a id="declaring-go-bindings"></a>

### 声明 Go 绑定

编写显式 `go_ffi`、`go_type` 和 `go_method` 声明，API 需要适配时使用 Go 转发函数。声明模型见 [Go 互操作](bind-go.zh-CN.md)，显式指针和方法绑定见[文件读取示例](../../goml/testdata/ffi/file-read/README.zh-CN.md)。

<a id="c-compatible-strings"></a>

### C 兼容字符串

`std::c` 通过普通 GoML 类型和调用提供 `CString`、带检查的 C 字符串字面量及边界错误。API 和编译期验证见 [C 兼容字符串](../stdlib.zh-CN.md#c-compatible-strings)。这些 API 不添加指针算术、原生布局或 `extern "C"` 语法。

<a id="exporting-a-go-library"></a>

### 导出 Go 库

前端识别 `#[go_export("Add")] pub fn add(x: i64, y: i64) -> i64 { x + y }`。检查声明必须为普通顶层公开函数，不带泛型、comptime/derive 能力或测试属性。Go 名以 ASCII 大写字母开头，仅含 ASCII 字母、数字或下划线；所选 GoML 包内名称必须唯一。重复属性及方法、extern、类型、字段和其他非函数项上的属性会被拒绝。

导出参数目前允许 bool、整数和浮点基本类型、`ffi::String` 及这些值的定长数组。返回额外允许空元组或由允许的非 `()` 值组成的扁平元组。透明别名规范化为底层类型。普通 GoML 字符串和字符、切片、通道、用户类型、函数值、元组参数及嵌套元组返回不在此导出边界内。

文本边界策略在 GoML 类型和适配器体中显式体现：

| 导出的 GoML 类型 | Go 签名 | 边界策略 |
| --- | --- | --- |
| `ffi::String` | `string` | 保留任意字节，包括无效 UTF-8 和 NUL；无隐式 Unicode 转换 |
| `ffi::Rune`（`i32` 的别名） | `int32` / `rune` | 保留原始整数；无隐式标量验证 |
| 普通 `string` / `char` | 拒绝 | 选择原始类型并编写显式带检查适配器 |

在导出函数内用 `ffi::string_text` 或 `ffi::rune_char` 验证值后再视为 Unicode。失败须通过适配器支持的返回签名或有意选择的边界失败策略表达。以下适配器对无效 UTF-8 返回原字节和 `false`，否则返回已验证文本和 `true`，从不静默替换为替换字符：

```goml
use std::ffi;

#[go_export("CheckedText")]
pub fn checked_text(value: ffi::String) -> (ffi::String, bool) {
    match ffi::string_text(value) {
        Ok(text) => (ffi::string_from_text(text), true),
        Err(_) => (value, false),
    }
}
```

返回 `ffi::String` 的原始回显可不经解码保留字节。rune 适配器同样可返回 `(ffi::Rune, bool)`。定长数组和扁平元组返回保留这些元素策略。`Result` 本身不是 Go 导出 ABI 类型；选择并记录失败形态由适配器作者负责。

接口和 Core 产物保留所选 Go 导出名、稳定包/函数身份、解析签名及相对源码来源。导出名和签名参与接口语义哈希，源码位置不参与。内部库链接器在死代码消除和单态化前，仅选取目标包声明的导出作为函数根；可达私有及跨包辅助函数被保留，无需 `main`。内部 Go 库后端为标量和定长数组参数/结果生成可调用包装器，将空元组映射为无 Go 结果，将扁平元组映射为多个 Go 结果。库包装器使用包级 OnceCell 初始化并保留 panic。生成的内部包声明使用私有 Go 标识符；只有显式选定的导出名成为公开 API。单独 `pub` 不请求 Go 导出。

使用 `goml export-go` 从当前 GoML 模块中的一个包生成库：

```text
goml export-go [PACKAGE_DIR] --import-path example.com/host/gen/calclib --out ./gen/calclib
```

默认源为 GoML 模块根。显式源路径和输出路径相对于当前工作目录。源必须是当前模块中的包，且不能声明 `package main`。仅其 `go_export` 声明成为 Go API；依赖导出仅在可达时保留。输出目录基本名是 Go 包名，必须是非 main 的 ASCII Go 标识符。

GoML 模块根必须已有 `go.mod`。输出必须位于该 Go 模块内、嵌套 Go 模块外，提供的导入路径必须等于 Go 模块路径加规范相对输出路径。拒绝父目录穿越，检查边界前解析符号链接。命令从不创建或修复 `go.mod` / `go.sum`，也不下载缺失依赖。

输出包含 `goml_generated.go` 和 `goml_exports.json`。清单记录公开签名、生成版本、Go 构建先决条件和所有权摘要。先决条件包括 Go 工具链、目标、cgo 设置、标志和构建标签。保留手写 Go 文件，并与候选生成源码一起检查。发布前通过 Go overlay 检查新输出；无效 Go、重复符号和导入循环会失败，不发布候选。已修改或不归生成器所有的文件不会被覆盖。同一输出的并发生成会被拒绝。发布会回滚已报告的文件替换错误；断电或强制终止进程时，两文件发布不具有崩溃原子性。

发布前，Go 编译检查候选生成包和每个已生成外部调用。GoML 分析仍信任外部声明；未使用声明不必出现在候选中。支持 `--compiler`、`--target-dir`、`--jobs` 和 `--dry-run`。生成的纯 Go 库仅需 Go 工具链即可构建。完整双向示例（含 Go 可执行程序和 Go 测试）见[互操作夹具](../../tools/release/testdata/go-export)。
