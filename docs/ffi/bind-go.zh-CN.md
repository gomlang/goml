<a id="go-interoperability"></a>

# Go 互操作

[English](bind-go.md) | [简体中文](bind-go.zh-CN.md)

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

完整边界规则见[语言指南](../goml.zh-CN.md#go-ffi)，显式指针和方法绑定见[文件读取示例](../../goml/testdata/ffi/file-read/README.zh-CN.md)，反向调用见[导出 Go 库](../goml.zh-CN.md#exporting-a-go-library)。
