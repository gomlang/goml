<a id="goml-api-documentation"></a>

# GoML API 文档

[English](documentation.md) | [简体中文](documentation.zh-CN.md)

`gomldoc` 从已检查的 GoML 源码生成文档。它与编译器查询引擎共用包分析、依赖导航、derive 展开和类型检查。文档生成不构建原生可执行文件。编译期求值和 Go FFI 验证仍遵循普通分析规则。

<a id="commands"></a>

## 命令

在模块内任意目录运行：

```sh
goml doc
goml doc --format json
goml doc --document-private-items
goml doc --target-dir build --dry-run
```

驱动器准备依赖并为所在模块调用 `gomldoc`。默认输出为 `<target-dir>/doc`，目标目录来自 `goml.toml` 或 `--target-dir`。包含生产包；排除测试、`testdata`、隐藏目录、构建目标目录和嵌套模块。

文档生成器依次通过 `--documenter`、`GOMLDOC`、`goml` 同目录程序、`GOML_HOME/bin` 和 `PATH` 查找。`--compiler` 选择驱动器项目发现使用的编译器。

独立命令使用本地可用的依赖源码：

```text
gomldoc [OPTIONS] [MODULE_DIR]

--format html|json        default: html
--output DIR              default: <target-dir>/doc
--target-dir DIR          override the module build target directory
--document-private-items  include private declarations and members
-h, --help
--version
```

`MODULE_DIR` 默认为当前目录；所在的清单确定模块。显式相对输出路径相对于工作目录。相对的 `--target-dir` 相对于模块根目录，并将该构建目录排除在发现范围之外。如果缺少依赖，先用 `goml update` 准备，或使用 `goml doc`。

独立命令成功时退出码为 0，分析或输出失败时为 1，参数无效时为 2。链接警告不会导致生成失败。只有分析和渲染成功后才写入输出。运行失败会保留原有文档。已有输出目录必须带有 `.gomldoc` 归属标记；任意已有目录不会被替换。生成文件归文档生成器管理，并在下次成功运行时替换。

<a id="comments"></a>

## 注释

```goml
//! Package overview.
package example;

/// A named value.
///
/// The public field can be read by callers.
pub struct Value {
    /// Display name.
    pub name: string,
}
```

`///` 和 `//!` 的规则见[注释语法与附着关系](language/basics.zh-CN.md#comments-and-api-documentation)。文档可以出现在函数、类型、常量、静态项、字段、变体、trait 成员、实现、方法和公共重导出上。[格式化器会保留附着关系](formatting.zh-CN.md#comments-and-literals)。

<a id="rendering-and-links"></a>

## 渲染和链接

支持的 Markdown 子集包含段落、ATX 标题、扁平的 `-` 或 `*` 项目符号列表、单反引号行内代码、三反引号围栏代码块和链接。原始 HTML 会被转义。尚未实现表格、强调、嵌套列表、图片和完整 CommonMark 语义。

`[Value]`、``[`Value`]`` 和 `[a value](goml:models::Value)` 是 API 引用。名称在声明文件的包及导入上下文中解释，包括显式重命名的导入和包声明的默认别名。支持规范包路径和 `module::` 路径。解析不会根据全局名称后缀猜测目标。无法解析或有歧义的引用产生诊断。已知但未生成页面的目标（如外部依赖）会渲染为文本。

普通链接目标接受 `https:`、`http:`、`mailto:` 和本地片段。其他协议会渲染为文本。所有 HTML 和 CSS 均在本地，打开 `index.html` 即可离线使用。

公共 API 输出排除私有声明、私有字段、私有固有方法，以及属于私有类型或 trait 的实现。trait 成员继承 trait 的可见性。泛型参数、约束、`where` 子句、别名和关联类型保留在签名中。方法体和常量/静态项初始化器被省略。重导出保留原始目标，生成的 API 会被标记为生成项。

<a id="output"></a>

## 输出

HTML 输出包含模块索引、每个包的页面、专门的结构体/枚举/trait 页面和样式表。符号标识符和页面名使用规范标识，trait 实现通过实现签名区分。页面显示相对源文件名和从 1 开始的源码行号；生成的声明使用第 0 行。

JSON 输出为 `module.json`：

```json
{"schema_version":1,"module":"example","packages":[]}
```

每个包包含 `name`、合并后的 `documentation`、`sections` 和 `items`。每节保留相对 `source`、`documentation` 和已解析的 `links`，使来自不同文件的包描述保留各自的导入上下文。每个条目包含 `id`、`name`、`qualified_name`、`kind`、`public`、`signature`、`documentation`、`source`、`line`、`generated`、`target`、`links`，以及递归嵌套的 `members`。链接包含引用 `text` 和规范 `target`。消费者必须检查 `schema_version`；此模式独立于编译器产物格式。

输出顺序是确定性的。生成内容既不含时间戳，也不含绝对源码路径。HTML 和 JSON 是同一文档模型的不同投影。

<a id="development"></a>

## 开发

`just make` 构建 `stage2/bin/gomldoc`；`just install` 和发布归档均包含它。它使用编译器版本模块和相对于可执行文件的库资源。

实现位于 `gomlc/doc_comment`、`gomlc/query/documentation*`、`gomlc/doc` 和 `gomlc/cmd/gomldoc`。项目集成位于 `goml/commands/doc.goml` 和 `goml/cmd/goml/doc.goml`。测试覆盖附着关系、可见性、导入、输出、错误和 CLI 集成。工具链及打包改动需要运行 `just ci`，包括不动点和归档迁移位置后的检查。

依赖文档站点、搜索、HTTP 服务器、自动启动浏览器和文档测试属于后续工作。
