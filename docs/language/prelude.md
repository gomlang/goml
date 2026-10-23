<a id="prelude-reference"></a>

# Prelude reference

[English](prelude.md) | [简体中文](prelude.zh-CN.md)

[GoML language guide](../goml.md)

Automatically scoped types, traits, and operations. Explicitly imported packages are covered in the [standard library reference](../stdlib.md).

On this page: [Built-in prelude](#built-in-prelude).

<a id="built-in-prelude"></a>

## Built-in prelude

The toolchain library has three separate layers. `builtin` is the hidden compiler and runtime contract: it owns primitive and built-in type identities, runtime hooks, language items, built-in implementations, and derive handlers. `prelude` is an independently compiled package that depends on `builtin` and defines the public names automatically placed in ordinary source scope. `std` contains normal standard-library packages and is available only through explicit imports.

User and third-party packages cannot explicitly import `builtin` or `prelude`; attempts produce a diagnostic instead of resolving a registry package. This keeps raw runtime helpers and derive internals hidden while preserving the stable identities of re-exported types and traits. Names such as `Option`, `Result`, `Vec`, `ToString`, `print`, `println`, and `range` remain available without `use` through the prelude.

Compiler-recognized protocols are declared in `builtin` with reserved `#[lang(...)]` metadata. The metadata identifies the owning type or trait together with protocol members such as enum variants, associated types, and methods, so compiler behavior does not depend on their public spelling. `#[lang(...)]` is rejected in user, third-party, prelude, and standard-library sources.

Runtime-backed collection types retain their builtin identity when re-exported. A qualified type from another package remains distinct even when it has the same short name: after `use std::collections::imm;`, `imm::HashMap[K, V]` is the persistent standard-library map, while the prelude `HashMap[K, V]` is the mutable runtime map. This uses ordinary package-qualified type paths and adds no syntax.

The following names can be used without `use`.

<a id="option-and-result"></a>

### `Option` and `Result`

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

Prefer `Some`, `None`, `Ok`, and `Err` in construction and patterns when the expected type identifies the enum. Qualified names such as `Option::Some` and `Result::Ok` remain available when type context is insufficient or another declaration shadows the variant.

`Option[T]` provides `is_some`, `is_none`, `unwrap_or`, and `unwrap_or_else`, plus type-changing generic methods:

- `value.map(map_fn) -> Option[U]`
- `value.and_then(next) -> Option[U]`
- `value.ok_or(error) -> Result[T, E]`

`Result[T, E]` provides `is_ok`, `is_err`, `unwrap_or`, and `unwrap_or_else`. Its type-changing generic methods are:

- `value.map(map_fn) -> Result[U, E]`
- `value.map_err(map_fn) -> Result[T, F]`
- `value.and_then(next) -> Result[U, E]`

<a id="default"></a>

### `Default`

The prelude `Default` implementations use `()` for the empty tuple, `false` for `bool`, zero for numeric types, and the empty string for `string`. `Vec` and `HashMap` default to empty containers, and the standard collection types `HashSet`, `IndexMap`, `Stack`, `Deque`, `BitSet`, `Arena`, `IndexVec`, and `Interner` likewise default through their empty constructors. `Option[T]` defaults to `None` without requiring `T: Default`, while tuples and fixed arrays default each element. `Result[T, E]` and `Ref[T]` intentionally have no global default implementation.

<a id="output-and-string-conversion"></a>

### Output and string conversion

- `print[T: ToString](value: T) -> ()`
- `println[T: ToString](value: T) -> ()`
- `value.to_string() -> string`, suitable for values that implement `ToString`
- `value.debug() -> string`, suitable for values that implement `Debug`
- `string.len() -> isize` and `string.byte_len() -> isize`, both returning the UTF-8 byte length
- `string.get(index: isize) -> char`, decoding a character at a UTF-8 byte boundary
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
- `string.starts_with_at(start: isize, prefix: string) -> bool`, checking at a byte offset
- `string.find(expected: string) -> Option[isize]`, `string.rfind(expected: string) -> Option[isize]`, and `string.find_bytes(expected: string) -> Option[isize]`, returning byte offsets
- `string.char_count() -> isize`, counting Unicode scalar values
- `string.slice_chars(start: isize, end: isize) -> Option[string]`, using scalar-value indexes
- `string.trim() -> string`, `string.trim_start() -> string`, and `string.trim_end() -> string`, trimming ASCII whitespace
- `string.split(separator: string) -> Vec[string]`, `string.split_once(separator: string) -> Option[(string, string)]`, and `string.lines() -> Vec[string]`
- `string.replace(expected: string, replacement: string) -> string`
- `string.repeat(count: isize) -> string`
- `string.is_ascii() -> bool`
- `string.eq_ignore_ascii_case(other: string) -> bool`
- `string.to_ascii_lowercase() -> string` and `string.to_ascii_uppercase() -> string`

The basic scalar types all implement `ToString` and `Debug`. String concatenation uses `+`.

The key-related traits are the method-bearing `PartialEq::eq(self, other: Self) -> bool`, the marker `Eq: PartialEq`, and `Hash::hash(self) -> u64`. User types can use handwritten impls or derive the corresponding traits.

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
- `ptr_eq(a, b) -> bool`, compare reference identities

The built-in `PartialEq`, `Eq`, and `Hash` implementations for `Ref[T]` use reference identity and do not require `T` to implement those traits. Mutating the referenced value therefore does not change equality or hashing. Separate allocations have distinct identities even when `T` is `()`, an empty struct, or another zero-sized type. Copies of one reference keep its identity. `Ref[T]` does not implement `Default`, because implicit allocation and recursive default construction would be surprising.

<a id="fixed-arrays"></a>

### Fixed arrays

```goml
let mut values: [i32; 3] = [1, 2, 3];
let first = values[0];
values[1] = 20;
```

The index type is `isize`. The underlying `array_get` and `array_set` can also be called; index syntax is preferred for daily code.

<a id="vect"></a>

### `Vec[T]`

```goml
let values = Vec::from_array([1, 2, 3]);
let first = values[0];
```

`Vec::from_array([values...])` creates a vector from a fixed array. Array items are evaluated once from left to right. The vector receives new outer backing storage, so replacing an array element later does not change the vector, while referenced elements remain shared because the copy is shallow. Use `Vec::new()` for an empty vector. Direct nonempty array calls coallocate the vector header and backing array. Passing an array variable retains the same shallow-copy semantics. The associated function is an ordinary first-class value, for example `let make: ([i32; 3]) -> Vec[i32] = Vec::from_array;`.

Commonly used methods:

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

`vec[index]` is equivalent to reading the element, `vec[index] = value;` modifies the element.

`copy()` shallow-copies the outer backing storage. `freeze()` also copies the backing storage and returns a `FrozenVec[T]`, so later structural or element changes through aliases to the original `Vec` cannot affect the frozen value.

`contains(value)` requires the element type to implement `PartialEq` and returns whether any element equals `value`.

Sorting and selection methods take an integer comparator returning negative/zero/positive, or an `Ordering` comparator for the `_ordering` variants. `dedup()` requires `PartialEq`; `position` takes a predicate; `binary_search_*` methods expect the vector to already be ordered and return the first matching index. `join(separator)` requires `ToString` and concatenates the element string forms.

<a id="frozenvect"></a>

### `FrozenVec[T]`

`FrozenVec[T]` is a dynamically sized read-only vector. It exposes `get`, `len`, `is_empty`, `contains`, `iter`, and `to_vec`, but no mutation methods. `to_vec()` returns a new shallow-copied mutable backing store. For nested containers, freezing is shallow: use `FrozenVec[FrozenVec[T]]` when both levels must be read-only.

<a id="slicet-and-mutslicet"></a>

### `Slice[T]` and `MutSlice[T]`

`Slice[T]` is a bounded read-only view on the `Vec[T]` storage:

```goml
let view: Slice[i32] = values.slice(1, 3);
let item = view.get(0);
```

Commonly used read-only methods are `get`, `get_checked`, `len`, `contains`, `sub`, `sub_checked`, `to_vec`, and `iter`. `contains(value)` requires the element type to implement `PartialEq`. `to_vec()` creates a shallow copy in a new `Vec[T]`; `view[index]` can be read, but `view[index] = value;` is rejected.

`MutSlice[T]` is a bounded mutable zero-copy view. `Vec::slice_mut(start, end)` and `Vec::as_mut_slice()` create it; `as_slice()` converts it to a read-only view without copying:

```goml
let values = Vec::from_array([1, 2, 3, 4]);
let output = values.slice_mut(1, 4);
output[0] = 20;
let written = output.copy_from(1, Vec::from_array([7, 8]).as_slice());
```

`get_checked`, `set_checked`, `sub_checked`, `Vec::slice_checked`, and `Vec::slice_mut_checked` report invalid ranges without indexing. `copy_from(offset, source)` validates the complete destination range before writing and uses overlap-safe copy semantics; `copy_within` is the corresponding same-view operation. `fill`, `to_vec`, and iteration are also available. `Slice` and `MutSlice` keep their current backing array alive. A later `Vec` growth may move the vector to a new backing array, so an older view continues to refer to the storage captured when the view was created.

<a id="hashmapk-v"></a>

### `HashMap[K, V]`

The key type must implement both `Hash` and `Eq`; bucket equality dispatches through the `PartialEq` supertrait:

```goml
let counts = HashMap::from_array([("a", 1), ("b", 2)]);
let value: Option[i32] = counts.get("a");
```

`HashMap::from_array([(key, value), ...])` creates a map from an entries array. It first evaluates the complete array from left to right, then inserts its pairs in array order. Later duplicate keys overwrite earlier entries. Use `HashMap::new()` for an empty map. Direct array-literal calls preallocate their build storage from the entry count. The associated function is an ordinary first-class value when its array length is fixed by a function type.

Commonly used methods: `new`, `from_array`, `get`, `set`, `insert`, `get_or_insert_with`, `update`, `remove`, `remove_value`, `entry`, `len`, `contains`, and `entries`. `insert` returns the previous value, `update` applies a `(V) -> V` closure only when the key exists, and `remove_value` returns the removed value. `entries()` returns a snapshot `Vec[(K, V)]`.

`entry(key)` returns a `HashMapEntry[K, V]` with `get`, `or_insert`, `or_insert_with`, `insert`, `and_modify`, `update`, `remove`, and `remove_entry`. The fused entry operations hash and search the key once; the creation closure passed to `or_insert_with` runs only for a missing key:

```goml
let counts: HashMap[string, isize] = HashMap::new();
let count = counts.entry("jobs").or_insert_with(|| 0);
let _ = counts.entry("jobs").update(|value| value + 1);
let removed = counts.entry("jobs").remove_entry();
```

Index reading returns `Option[V]`, and index assignment writes `V`:

```goml
let value = counts["a"];
counts["b"] = 2;
```

See [map algorithms](../stdlib.md#map-algorithms) for cloning, copying, comparing, and iterating maps with `std::collections`.

<a id="channels"></a>

### Channels

`Channel[T]` is a Go channel backend that supports buffered and unbuffered communication:

```goml
let channel = Channel::[string]::new(0);
go || channel.send("ready");
let value: Option[string] = channel.recv();
channel.close();
```

Commonly used methods:

- `Channel::[T]::new(capacity: isize) -> Channel[T]`
- `send(value: T) -> ()`
- `recv() -> Option[T]`, returns `Option::None` when closed and drained
- `close() -> ()`
- `sender() -> Sender[T]`
- `receiver() -> Receiver[T]`
- `split() -> (Sender[T], Receiver[T])`

`Sender[T]` provides `send` and `close`; `Receiver[T]` provides `recv`. Direction is enforced statically and emitted as Go's `chan<- T` and `<-chan T`. Converting a bidirectional channel requires an explicit endpoint method. `Channel[T]` keeps its original send, receive, and close API.

Capacity `0` creates an unbuffered channel. Sending to an unbuffered channel should generally be concurrently received by another `go` closure, and vice versa.

Use `select` when a goroutine must wait on several channels, choose between sending and receiving, observe closure through `Option[T]`, or perform a non-blocking operation with `default`. The channel operands and send values are prepared once before selection.

See [dynamic channel selection](../stdlib.md#dynamic-channel-selection) for runtime-sized selection with `std::channel`.

<a id="iterator"></a>

### Iterator

The built-in protocols are:

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

The protocol entry points are `FnIterator::from_fn(next_fn)`, `iterator.next()` or `Iterator::next(iterator)`. `range(start, end)` and `start..end` produce incrementing half-open `FnIterator[isize]` values; `start >= end` is empty.

See [iterator adapters](../stdlib.md#iterator-adapters) for construction, adapters, and consumers in `std::iter`.

Iterators are single pass. Fixed arrays use native indexed `for` lowering. `Vec[T]`, `Slice[T]` and `MutSlice[T]` can be directly used in `for`, and a value implementing `Iterator` is also directly iterable through the identity `IntoIterator`.
