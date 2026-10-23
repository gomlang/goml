<a id="prelude-reference"></a>

# Prelude 参考

[English](prelude.md) | [简体中文](prelude.zh-CN.md)

[GoML 语言指南](../goml.zh-CN.md)

自动进入作用域的类型、trait 和操作。需要显式导入的包见[标准库参考](../stdlib.zh-CN.md)。

本页内容：[内置 prelude](#built-in-prelude)。

<a id="built-in-prelude"></a>

## 内置 prelude

工具链库分为三层。`builtin` 是隐藏的编译器和运行时契约：拥有基本及内置类型身份、运行时钩子、语言项、内置实现和 derive 处理函数。`prelude` 是依赖 `builtin` 的独立编译包，定义自动加入普通源码作用域的公开名称。`std` 包含普通标准库包，仅通过显式导入可用。

用户和第三方包不能显式导入 `builtin` 或 `prelude`；尝试会产生诊断，不会解析注册表包。这隐藏原始运行时辅助函数和 derive 内部机制，同时保留重新导出类型/trait 的稳定身份。`Option`、`Result`、`Vec`、`ToString`、`print`、`println` 和 `range` 等名称通过 prelude 无需 `use` 即可使用。

编译器识别的协议在 `builtin` 中用保留的 `#[lang(...)]` 元数据声明。元数据标识所属类型或 trait 及枚举变体、关联类型和方法等协议成员，使编译器行为不依赖公开拼写。用户、第三方、prelude 和标准库源码中的 `#[lang(...)]` 会被拒绝。

由运行时支持的集合类型在重新导出时保留 builtin 身份。其他包的限定类型即使短名相同也仍不同：`use std::collections::imm;` 后，`imm::HashMap[K, V]` 是持久化标准库映射，而 prelude 的 `HashMap[K, V]` 是可变运行时映射。这使用普通包限定类型路径，不增加语法。

以下名称无需 `use` 即可使用。

<a id="option-and-result"></a>

### `Option` 与 `Result`

```goml
enum Option[T] {
    None,
    Some(T),
}

enum Result[T, E] {
    Ok(T),
    Err(E),
}
```

预期类型能确定枚举时，构造和模式中优先使用 `Some`、`None`、`Ok` 和 `Err`。类型上下文不足或其他声明遮蔽变体时，仍可用 `Option::Some`、`Result::Ok` 等限定名。

`Option[T]` 提供 `is_some`、`is_none`、`unwrap_or` 和 `unwrap_or_else`，以及改变类型的泛型方法：

- `value.map(map_fn) -> Option[U]`
- `value.and_then(next) -> Option[U]`
- `value.ok_or(error) -> Result[T, E]`

`Result[T, E]` 提供 `is_ok`、`is_err`、`unwrap_or` 和 `unwrap_or_else`。改变类型的泛型方法为：

- `value.map(map_fn) -> Result[U, E]`
- `value.map_err(map_fn) -> Result[T, F]`
- `value.and_then(next) -> Result[U, E]`

<a id="default"></a>

### `Default`

prelude 的 `Default` 实现为空元组使用 `()`，为 `bool` 使用 `false`，为数值类型使用零，为 `string` 使用空字符串。`Vec` 和 `HashMap` 默认为空容器，标准集合 `HashSet`、`IndexMap`、`Stack`、`Deque`、`BitSet`、`Arena`、`IndexVec` 和 `Interner` 也通过空构造器默认初始化。`Option[T]` 默认为 `None`，不要求 `T: Default`；元组和定长数组逐元素默认初始化。`Result[T, E]` 和 `Ref[T]` 有意不提供全局默认实现。

<a id="output-and-string-conversion"></a>

### 输出与字符串转换

- `print[T: ToString](value: T) -> ()`
- `println[T: ToString](value: T) -> ()`
- `value.to_string() -> string`，适用于实现 `ToString` 的值
- `value.debug() -> string`，适用于实现 `Debug` 的值
- `string.len() -> isize` 和 `string.byte_len() -> isize`，都返回 UTF-8 字节长度
- `string.get(index: isize) -> char`，在 UTF-8 字节边界解码字符
- `string.byte_get(index: isize) -> u8`
- `string.byte_slice(start: isize, end: isize) -> string`
- `string.is_char_boundary(index: isize) -> bool`
- `string.decode_at(index: isize) -> Option[(char, isize)]`
- `string.to_bytes() -> Vec[u8]`
- `string.chars() -> FnIterator[char]`
- `string.char_indices() -> FnIterator[(isize, char)]`
- `string.starts_with(prefix: string) -> bool`
- `string.ends_with(suffix: string) -> bool`
- `string.contains(expected: string) -> bool`
- `string.starts_with_at(start: isize, prefix: string) -> bool`，在字节偏移处检查
- `string.find(expected: string) -> Option[isize]`、`string.rfind(expected: string) -> Option[isize]` 和 `string.find_bytes(expected: string) -> Option[isize]`，返回字节偏移
- `string.char_count() -> isize`，计数 Unicode 标量值
- `string.slice_chars(start: isize, end: isize) -> Option[string]`，使用标量值索引
- `string.trim() -> string`、`string.trim_start() -> string` 和 `string.trim_end() -> string`，修剪 ASCII 空白
- `string.split(separator: string) -> Vec[string]`、`string.split_once(separator: string) -> Option[(string, string)]` 和 `string.lines() -> Vec[string]`
- `string.replace(expected: string, replacement: string) -> string`
- `string.repeat(count: isize) -> string`
- `string.is_ascii() -> bool`
- `string.eq_ignore_ascii_case(other: string) -> bool`
- `string.to_ascii_lowercase() -> string` 和 `string.to_ascii_uppercase() -> string`

基本标量类型都实现 `ToString` 和 `Debug`。字符串拼接使用 `+`。

与键相关的 trait 包括带方法的 `PartialEq::eq(self, other: Self) -> bool`、标记 `Eq: PartialEq` 和 `Hash::hash(self) -> u64`。用户类型可手写实现或派生相应 trait。

<a id="reft"></a>

### `Ref[T]`

```goml
let cell = Ref::new(1);
let before = cell.get();
cell.set(before + 1);
```

API:

- `Ref::new(value) -> Ref[T]`
- `reference.get() -> T`
- `reference.set(value) -> ()`
- `ptr_eq(a, b) -> bool`，比较引用身份

`Ref[T]` 的内置 `PartialEq`、`Eq` 和 `Hash` 实现使用引用身份，不要求 `T` 实现这些 trait。因此修改引用值不会改变相等性或哈希。即使 `T` 为 `()`、空结构体或其他零大小类型，独立分配也有不同身份。同一引用的副本保留身份。`Ref[T]` 不实现 `Default`，因为隐式分配和递归默认构造容易造成意外。

<a id="fixed-arrays"></a>

### 定长数组

```goml
let mut values: [i32; 3] = [1, 2, 3];
let first = values[0];
values[1] = 20;
```

索引类型为 `isize`。也可调用底层 `array_get` 和 `array_set`；日常代码优先使用索引语法。

<a id="vect"></a>

### `Vec[T]`

```goml
let values = Vec::from_array([1, 2, 3]);
let first = values[0];
```

`Vec::from_array([values...])` 从定长数组创建向量。数组项从左到右求值一次。向量获得新的外层底层存储，因此之后替换数组元素不影响向量；由于浅复制，引用元素仍共享。空向量使用 `Vec::new()`。直接非空数组调用联合分配向量头和底层数组。传入数组变量保留相同浅复制语义。该关联函数是普通一等值，例如 `let make: ([i32; 3]) -> Vec[i32] = Vec::from_array;`。

常用方法：

- `Vec::new() -> Vec[T]`
- `Vec::from_array(values: [T; N]) -> Vec[T]`
- `Vec::with_capacity(capacity) -> Vec[T]`
- `push(value) -> ()`
- `pushed(value) -> Vec[T]`
- `copy() -> Vec[T]`
- `freeze() -> FrozenVec[T]`
- `get(index) -> T`
- `set(index, value) -> ()`
- `len() -> isize`
- `capacity() -> isize`
- `is_empty() -> bool`
- `contains(value) -> bool`
- `position(predicate) -> Option[isize]`
- `sort_by(compare) -> ()`
- `stable_sort_by(compare) -> ()`
- `sort_by_ordering(compare) -> ()`
- `stable_sort_by_ordering(compare) -> ()`
- `binary_search_by(compare) -> Option[isize]`
- `binary_search_by_ordering(compare) -> Option[isize]`
- `min_by(compare) -> Option[T]`
- `max_by(compare) -> Option[T]`
- `min_by_ordering(compare) -> Option[T]`
- `max_by_ordering(compare) -> Option[T]`
- `dedup_by(equal) -> ()`
- `dedup() -> ()`
- `join(separator) -> string`
- `reserve(additional) -> ()`
- `truncate(len) -> ()`
- `clear() -> ()`
- `last() -> Option[T]`
- `pop() -> Option[T]`
- `swap(left, right) -> ()`
- `swap_remove(index) -> T`
- `insert(index, value) -> ()`
- `remove(index) -> T`
- `reverse() -> ()`
- `extend(other) -> ()`
- `slice(start, end) -> Slice[T]`
- `iter() -> FnIterator[T]`

`vec[index]` 读取元素，`vec[index] = value;` 修改元素。

`copy()` 浅复制外层底层存储。`freeze()` 也复制底层存储并返回 `FrozenVec[T]`，因此之后通过原 `Vec` 别名进行结构或元素修改不会影响冻结值。

`contains(value)` 要求元素类型实现 `PartialEq`，返回是否存在等于 `value` 的元素。

排序和选择方法接受返回负数/零/正数的整数比较器，`_ordering` 变体则接受 `Ordering` 比较器。`dedup()` 要求 `PartialEq`；`position` 接受谓词；`binary_search_*` 要求向量已排序，并返回首个匹配索引。`join(separator)` 要求 `ToString`，拼接元素的字符串形式。

<a id="frozenvect"></a>

### `FrozenVec[T]`

`FrozenVec[T]` 是动态大小的只读向量，暴露 `get`、`len`、`is_empty`、`contains`、`iter` 和 `to_vec`，但没有修改方法。`to_vec()` 返回新的浅复制可变底层存储。嵌套容器的冻结是浅层的：两层都须只读时使用 `FrozenVec[FrozenVec[T]]`。

<a id="slicet-and-mutslicet"></a>

### `Slice[T]` 与 `MutSlice[T]`

`Slice[T]` 是 `Vec[T]` 存储上的有界只读视图：

```goml
let view: Slice[i32] = values.slice(1, 3);
let item = view.get(0);
```

常用只读方法有 `get`、`get_checked`、`len`、`contains`、`sub`、`sub_checked`、`to_vec` 和 `iter`。`contains(value)` 要求元素类型实现 `PartialEq`。`to_vec()` 在新 `Vec[T]` 中创建浅副本；可以读取 `view[index]`，但 `view[index] = value;` 会被拒绝。

`MutSlice[T]` 是有界可变零拷贝视图。`Vec::slice_mut(start, end)` 和 `Vec::as_mut_slice()` 创建它；`as_slice()` 无需复制即可转为只读视图：

```goml
let values = Vec::from_array([1, 2, 3, 4]);
let output = values.slice_mut(1, 4);
output[0] = 20;
let written = output.copy_from(1, Vec::from_array([7, 8]).as_slice());
```

`get_checked`、`set_checked`、`sub_checked`、`Vec::slice_checked` 和 `Vec::slice_mut_checked` 在索引前报告无效范围。`copy_from(offset, source)` 写入前验证完整目标范围，并使用重叠安全复制语义；`copy_within` 是对应的同视图操作。还支持 `fill`、`to_vec` 和迭代。`Slice` 和 `MutSlice` 保持当前底层数组存活。之后 `Vec` 增长可能将向量移到新底层数组，因此旧视图继续引用创建时捕获的存储。

<a id="hashmapk-v"></a>

### `HashMap[K, V]`

键类型必须同时实现 `Hash` 和 `Eq`；桶相等比较通过父 trait `PartialEq` 分派：

```goml
let counts = HashMap::from_array([("a", 1), ("b", 2)]);
let value: Option[i32] = counts.get("a");
```

`HashMap::from_array([(key, value), ...])` 从条目数组创建映射。先从左到右求值整个数组，再按数组顺序插入键值对。后面的重复键覆盖先前条目。空映射使用 `HashMap::new()`。直接数组字面量调用根据条目数预分配构建存储。数组长度由函数类型固定时，该关联函数是普通一等值。

常用方法：`new`、`from_array`、`get`、`set`、`insert`、`get_or_insert_with`、`update`、`remove`、`remove_value`、`entry`、`len`、`contains` 和 `entries`。`insert` 返回旧值，`update` 仅键存在时应用 `(V) -> V` 闭包，`remove_value` 返回移除值。`entries()` 返回快照 `Vec[(K, V)]`。

`entry(key)` 返回 `HashMapEntry[K, V]`，提供 `get`、`or_insert`、`or_insert_with`、`insert`、`and_modify`、`update`、`remove` 和 `remove_entry`。融合的 entry 操作只对键哈希和搜索一次；传给 `or_insert_with` 的创建闭包仅在键缺失时运行：

```goml
let counts: HashMap[string, isize] = HashMap::new();
let count = counts.entry("jobs").or_insert_with(|| 0);
let _ = counts.entry("jobs").update(|value| value + 1);
let removed = counts.entry("jobs").remove_entry();
```

索引读取返回 `Option[V]`，索引赋值写入 `V`：

```goml
let value = counts["a"];
counts["b"] = 2;
```

使用 `std::collections` 克隆、复制、比较和迭代映射见[映射算法](../stdlib.zh-CN.md#map-algorithms)。

<a id="channels"></a>

### 通道

`Channel[T]` 使用 Go 通道后端，支持有缓冲和无缓冲通信：

```goml
let channel = Channel::[string]::new(0);
go || channel.send("ready");
let value: Option[string] = channel.recv();
channel.close();
```

常用方法：

- `Channel::[T]::new(capacity: isize) -> Channel[T]`
- `send(value: T) -> ()`
- `recv() -> Option[T]`，关闭且排空时返回 `Option::None`
- `close() -> ()`
- `sender() -> Sender[T]`
- `receiver() -> Receiver[T]`
- `split() -> (Sender[T], Receiver[T])`

`Sender[T]` 提供 `send` 和 `close`；`Receiver[T]` 提供 `recv`。方向在静态检查中强制执行，并生成为 Go 的 `chan<- T` 和 `<-chan T`。转换双向通道需要显式端点方法。`Channel[T]` 保留原有发送、接收和关闭 API。

容量 `0` 创建无缓冲通道。发送到无缓冲通道通常应由另一个 `go` 闭包并发接收，反之亦然。

goroutine 需要等待多个通道、在发送和接收间选择、通过 `Option[T]` 观察关闭，或通过 `default` 做非阻塞操作时，使用 `select`。通道操作数和发送值在选择前准备一次。

使用 `std::channel` 进行运行时数量的选择见[动态通道选择](../stdlib.zh-CN.md#dynamic-channel-selection)。

<a id="iterator"></a>

### 迭代器

内置协议如下：

```goml
trait Iterator {
    type Item;
    fn next(self) -> Option[Self::Item];
}

trait IntoIterator {
    type Item;
    type IntoIter: Iterator;
    fn into_iter(self) -> Self::IntoIter;
}
```

协议入口为 `FnIterator::from_fn(next_fn)`、`iterator.next()` 或 `Iterator::next(iterator)`。`range(start, end)` 和 `start..end` 产生递增半开 `FnIterator[isize]`；`start >= end` 时为空。

`std::iter` 中的构造、适配器和消费操作见[迭代器适配器](../stdlib.zh-CN.md#iterator-adapters)。

迭代器只能单次遍历。定长数组使用原生索引 `for` 降级。`Vec[T]`、`Slice[T]` 和 `MutSlice[T]` 可直接用于 `for`，实现 `Iterator` 的值也通过恒等 `IntoIterator` 直接迭代。
