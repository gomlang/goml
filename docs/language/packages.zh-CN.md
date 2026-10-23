<a id="modules-packages-and-visibility"></a>

# 模块、包与可见性

[English](packages.md) | [简体中文](packages.zh-CN.md)

[GoML 语言指南](../goml.zh-CN.md)

源码中的包身份、导入和可见性。模块清单、工作区、依赖解析和构建布局见[项目工具](../tooling.zh-CN.md)。

本页内容：[模块、包和导入](#modules-packages-and-imports)。

<a id="modules-packages-and-imports"></a>

## 模块、包和导入

<a id="directory-packages"></a>

### 目录包

每个包含 `.goml` 文件的目录都是一个包。同一目录的所有源文件必须声明相同包名：

```goml
package utils;
```

包的规范身份由模块路径和相对目录路径组成。包声明名称决定无别名导入时的本地名称，而不决定包的全局身份。根包的规范身份就是模块路径。

典型项目可以是：

```text
goml.toml
main.goml
utils/utils.goml
```

`utils/utils.goml`:

```goml
package utils;

pub fn message() -> string {
    "hello"
}
```

`main.goml`:

```goml
package main;

use alice::myapp::utils;

fn main() -> () {
    println(utils::message())
}
```

<a id="use"></a>

### `use`

`use` 是文件级声明，必须放在 `package` 之后、任何顶层项之前。不同文件不共享导入。

导入不传递：包 A 导入包 B，并不会让使用 A 的文件自动看到 B。每个文件必须直接导入它引用的每个包。

```goml
use alice::http::client;
use alice::http::client as http_client;
```

上下文路径首段 `module` 表示当前项目 `[module].path` 声明的规范路径。模块中的任何包都可以使用，包括规范路径包含多段的模块：

```goml
use module::http::client;
use module::rendering::api::{Canvas, Render};
pub use module::model::Request;
```

对于声明为 `alice::myapp` 的模块，这些路径解析为 `alice::myapp::http::client`、`alice::myapp::rendering::api` 和 `alice::myapp::model::Request`。此标记仅在 `use` 路径开头且后跟 `::` 时生效；其他位置名为 `module` 的依赖路径和别名保留普通含义。外部依赖继续使用 `alice::http::client` 这样的规范路径。

导入包后，通过本地别名访问其公开项：

```goml
let request = http_client::new_request();
```

花括号导入将选定的公开顶层项直接引入当前文件。可以导入函数、常量、结构体、枚举、类型别名和 trait，并可分别重命名：

```goml
use alice::rendering::api::{Canvas, Color as Paint, Render, render_to_string};

fn describe(canvas: Canvas) -> string {
    render_to_string(canvas)
}
```

花括号前的路径始终是包路径。不支持嵌套花括号导入、`self`、`*` 通配导入、枚举变体导入或子包聚合。导入名仍限于当前文件，并遵循普通顶层可见性规则。

以下特殊形式还可以将已加载包中的 trait 加入方法调用作用域：

```goml
use alice::rendering::api;
use api::Render;
```

在花括号列表中导入 `Render` 也会将其加入方法作用域。之后可对具体值写 `value.render()`。即使未将 trait 加入方法作用域，仍可写限定调用 `api::Render::render(value)`。

`pub use` 从当前包重新导出具名包或公开项，且不改变其身份：

```goml
pub use alice::http::client;
pub use client::{Client, Request, Response};
pub use client::Request as HttpRequest;
```

下游代码可以使用 `facade::HttpRequest`、直接导入它，或将重新导出的 trait 引入方法作用域。编译器将这些路径都解析回原始声明，因此构造器、trait 实现、方法、常量和函数的行为与原始路径完全一致。重新导出的包命名空间也仍可使用，例如 `facade::client::Request`。同包中的重新导出名必须唯一。仍不支持通配重新导出。

不支持以下路径模型：

```text
mod child;
crate::x
self::x
super::x
::rooted::x
```

<a id="visibility"></a>

### 可见性

顶层项默认只在包内可见。顶层函数、结构体、枚举和 trait 用 `pub` 导出；结构体字段和固有 impl 方法也默认私有，跨包访问时须单独标注 `pub`：

```goml
pub struct Point {
    pub x: i32,
    pub y: i32,
}

pub fn origin() -> Point {
    Point { x: 0, y: 0 }
}

impl Point {
    pub fn sum(self) -> i32 {
        self.x + self.y
    }
}
```

同包中的所有文件都可以使用私有顶层项。trait impl 方法继承 trait 方法的可见性，不能写 `pub`。
