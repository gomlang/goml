<a id="goml-formatting"></a>

# GoML 格式化

[English](formatting.md) | [简体中文](formatting.zh-CN.md)

`gomlfmt` 是 GoML 源码的规范格式化器。它将源码解析为无损具体语法树，再将该树转换为描述可选换行位置的文档，最后以固定宽度渲染。

格式化从不使用降级后的 AST。注释、标识符、数字写法、字符和字节转义、字符串及插值字符串文本均从原始词法 token 输出。多行字符串保留其标记和内容文本，仅规范化每个标记之前的布局缩进。

<a id="fixed-settings"></a>

## 固定设置

- 最大行宽：100 列
- 缩进：四个空格
- 行尾：LF
- 非空文件结尾：恰好一个 LF

格式化器没有配置文件或命令行样式设置。测试可以使用更窄的宽度来覆盖换行行为。

列按 Unicode 码点计数，不特殊衡量东亚显示宽度和字素簇。尾随的 `//` 注释可能超过最大行宽。

<a id="files-and-items"></a>

## 文件和项

文件头注释始终位于最前。`package` 声明独占一行，后跟空行。连续的 `use` 声明占据连续行，整个 use 块后跟空行。导入不排序、不合并，也不调整顺序。

顶层项之间以一个空行分隔。属性保持附着于后续项，每行一个属性。

结构体、枚举、trait、impl、函数、match 和其他非空主体采用多行布局：

```goml
fn choose(value: bool) -> string {
    if value {
        "yes"
    } else {
        "no"
    }
}
```

空块写为 `{}`。`else` 与前一个右花括号保持在同一行。不含注释的空 `else {}` 分支会被省略，因为没有 `else` 的 `if` 具有相同的隐式 `()` 回退值。

块末尾恰好为 `()` 的表达式及独立的 `();` 在不含注释时会被省略。多余分号构成的空语句也会被移除。这些规范化保留块的 `()` 结果，不删除注释。

无标签 `while` 和 `for` 表达式后的分号会被移除。若后面还有语句，`if`、`match`、`select` 和无标签 `loop` 后的分号会被移除。以隐式 `()` 为回退值的末尾 `if` 也省略分号。带标签循环语句保留分号，因为解析器要求在后续语句前保留它。末尾控制流分号在可能丢弃值时保留；若下一个 token 可能延续前一个表达式，如 `(`、`[`、`.`、`?`、`as`、范围运算符或二元运算符，也会保留控制流分号。

只有一个字段的结构体式枚举变体，在所有此类变体都能放入单行时保持单行。如果任一结构体式变体有多个字段、含有注释或超过行宽，则枚举内所有非空结构体式变体采用相同的展开布局：

```goml
enum Message[T] {
    Empty,
    Value { value: T },
}

enum Pair[T] {
    Left {
        value: T,
    },
    Both {
        left: T,
        right: T,
    },
}
```

<a id="spaces"></a>

## 空格

- 逗号或冒号后若同行还有内容，则跟一个空格。
- 二元运算符、`=`、`=>`、`->` 和 `as` 两侧各一个空格。
- `::`、`.`、后缀 `?`、调用和索引周围不添加空格。
- 前缀运算符与操作数之间不加空格。
- 圆括号和方括号的内侧边缘不加空格。

数组重复语法保留初始化器和长度之间的分号，如 `[value; N]`。let 链分隔符和 trait 约束分隔符采用普通运算符间距：`if let Some(x) = value && x > 0`、`impl Source[Item = T] + Label` 和 `dyn Consumer[string] + Close`。

格式化器保留 CST 的括号和结合性，不插入或移除分组括号。

<a id="lists-and-expressions"></a>

## 列表和表达式

一组内容能放下时保持单行：

```goml
call(first, second, third)
```

放不下时，每个列表元素放在各自的缩进行上：

```goml
call(
    first,
    second,
    third,
)
```

语法允许时，换行列表会添加尾随逗号。除了上述块规范化外，条件性的尾随逗号是格式化唯一可以添加或移除的非空白注释类 token。嵌套组独立决定宽度布局。

较长的二元运算链在运算符之前换行：

```goml
first
    + second
    + third
```

<a id="comments-and-literals"></a>

## 注释和字面量

`//` 注释的文本和顺序保持不变。同行注释仍附着于前面的语法，输出在该行末尾。独立注释仍独占一行。独立注释之间最多保留一个空行。

附着的 `///` 文档块与后续声明保持在一起，即使中间存在属性。格式化保留它们对字段、变体、trait 成员和方法的附着关系，不在文档块与声明之间插入空行。`//!` 包文档保持在文件头。见[文档说明](documentation.zh-CN.md)。

列表内的注释会强制换行。尾随注释不折行，也不改写。

字面量 token 文本逐字节保留，包括进制前缀、数字分隔符、转义和插值字符串文本。多行字符串另起一行，其所有 `\\` 标记对齐，结束语法恢复到外层缩进。只改写 `\\` 标记前的空格和制表符；标记后的内容逐字节保留。

```goml
let poem =
    \\roses are red
    \\violets are blue
;
```

<a id="errors"></a>

## 错误

无效源码不会被部分格式化。库返回解析器诊断，stdout 模式不输出格式化源码，`-w` 在写入任何文件前验证全部输入。

<a id="command-line"></a>

## 命令行

```text
gomlfmt [OPTIONS] [FILE...]

-w          write formatted source back to files
-l          list files whose contents would change
--check     exit 1 when any file needs formatting
-h, --help
--version
```

文件参数必须使用 `.goml` 后缀。没有文件参数或只有一个 `-` 时，格式化器读取 stdin 并写入 stdout。`-w` 不能与 stdin 一起使用。未指定模式的单文件会格式化到 stdout。多个文件要求指定 `-w`、`-l` 或 `--check`。

`--check` 在全部文件格式正确时退出码为 0，需要改动时为 1，参数、I/O 或解析错误时为 2。

在模块内任意位置，使用项目驱动器格式化全部 `.goml` 文件：

```text
goml fmt
goml fmt --check
```

`goml fmt` 向上查找 `goml.toml`，并使用与模块构建和测试相同的包图。它格式化当前模块的生产文件、内部 `*_test.goml` 文件、黑盒 `tests` 包，以及命名 `examples/` 及其测试，同时排除外部依赖。测试和示例导入会解析开发依赖。包发现排除 `testdata`、配置的构建目标目录、隐藏目录和嵌套模块。`--check` 不写入文件。

格式化器可执行文件依次从 `--formatter`、`GOMLFMT`、`goml` 所在目录、`GOML_HOME/bin` 和 `PATH` 解析。

<a id="editor-formatting"></a>

## 编辑器格式化

`gomllsp` 实现 `textDocument/formatting`。它格式化最新的内存文档，源码发生变化时返回一次整篇文档编辑。已格式化或语法无效的文档不返回编辑。

编辑器提供的缩进选项会被忽略，因为 GoML 格式化使用上述固定设置。整篇文档的编辑范围按语言服务器协议要求使用 UTF-16 位置。
