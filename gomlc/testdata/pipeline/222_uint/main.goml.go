package main

import (
    _goml_os "os"
)

func _goml_runtime_core_bool_to_string(x bool) string {
    if x {
        return "true"
    } else {
        return "false"
    }
}

func _goml_runtime_core_string_from_utf8(bytes *_goml_vec_uint8) Tuple2_4bool_6string {
    return Tuple2_4bool_6string{
        _0: true,
        _1: string(bytes.items),
    }
}

func _goml_runtime_core_string_println(s string) struct{} {
    _goml_os.Stdout.WriteString(s + "\n")
    return struct{}{}
}

type _goml_vec_uint8 struct {
    items []uint8
}

func vec_with_capacity__Vec_5uint8(capacity int) *_goml_vec_uint8 {
    return &_goml_vec_uint8{
        items: make([]uint8, 0, capacity),
    }
}

func vec_push__Vec_5uint8(vec *_goml_vec_uint8, elem uint8) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_get__Vec_5uint8(vec *_goml_vec_uint8, index int) uint8 {
    return vec.items[index]
}

func vec_len__Vec_5uint8(vec *_goml_vec_uint8) int {
    return int(len(vec.items))
}

type _goml_vec_uint32 struct {
    items []uint32
}

type hashmap_uint_string_x_entry struct {
    active bool
    key uint
    value string
}

type hashmap_uint_string_x struct {
    indices map[uint]int
    entries []hashmap_uint_string_x_entry
    len int
}

func hashmap_new__HashMap_4uint_6string() *hashmap_uint_string_x {
    return &hashmap_uint_string_x{
        indices: make(map[uint]int),
        entries: nil,
        len: 0,
    }
}

func hashmap_lookup__HashMap_4uint_6string(m *hashmap_uint_string_x, key uint) (string, bool, int, uint64) {
    if m == nil {
        var zero string
        return zero, false, -1, 0
    }
    var index int
    var found bool
    index, found = m.indices[key]
    if !found {
        var zero string
        return zero, false, -1, 0
    }
    var entry hashmap_uint_string_x_entry = m.entries[index]
    if entry.active {
        return entry.value, true, index, 0
    }
    var zero string
    return zero, false, index, 0
}

func hashmap_get__HashMap_4uint_6string(m *hashmap_uint_string_x, key uint) Option__string {
    var value string
    var ok bool
    value, ok, _, _ = hashmap_lookup__HashMap_4uint_6string(m, key)
    if ok {
        return Option__string{
            _tag: 1,
            _p0: value,
        }
    }
    return Option__string{
        _tag: 0,
    }
}

func hashmap_set__HashMap_4uint_6string(m *hashmap_uint_string_x, key uint, value string) struct{} {
    if m == nil {
        return struct{}{}
    }
    var index int
    var found bool
    index, found = m.indices[key]
    if found {
        var entry hashmap_uint_string_x_entry = m.entries[index]
        if entry.active {
            m.entries[index].value = value
            return struct{}{}
        }
        m.entries[index] = hashmap_uint_string_x_entry{
            active: true,
            key: key,
            value: value,
        }
        m.len = m.len + 1
        return struct{}{}
    }
    index = len(m.entries)
    m.indices[key] = index
    m.entries = append(m.entries, hashmap_uint_string_x_entry{
        active: true,
        key: key,
        value: value,
    })
    m.len = m.len + 1
    return struct{}{}
}

type Tuple2_4bool_6string struct {
    _0 bool
    _1 string
}

type FloatNatural struct {
    words *_goml_vec_uint32
}

type ParsedFloat struct {
    valid bool
    negative bool
    special int
    numerator FloatNatural
    decimal_exponent int
    binary_exponent int
    hexadecimal bool
    significant_digits int
}

type Ordering uint8

type Option__string struct {
    _p0 string
    _tag uint8
}

func main0() struct{} {
    var left__0 uint = 19
    var right__0_source int = 2
    var right__0 uint = uint(int(right__0_source))
    var result__0 uint
    var inline21 uint = left__0 + right__0
    var inline22 uint = inline21 * 2
    result__0 = inline22
    var inline19 string = _goml_m_trait__impl_i_ToString_i_usize_i_to__string(result__0)
    _goml_runtime_core_string_println(inline19)
    var t0 string
    var inline18 string = __goml_builtin_uint_to_string(result__0)
    t0 = inline18
    var inline16 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t0)
    _goml_runtime_core_string_println(inline16)
    var t1 string
    switch result__0 {
    case 0:
        t1 = "zero"
    case 42:
        t1 = "answer"
    default:
        t1 = "other"
    }
    var inline14 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t1)
    _goml_runtime_core_string_println(inline14)
    var t2 bool = result__0 > left__0
    var inline12 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t2)
    _goml_runtime_core_string_println(inline12)
    var t3 uint = result__0 & 15
    var t4 uint64 = uint64(uint(t3))
    var inline10 string = _goml_m_trait__impl_i_ToString_i_u64_i_to__string(t4)
    _goml_runtime_core_string_println(inline10)
    var values__0 *hashmap_uint_string_x
    var inline9 *hashmap_uint_string_x = hashmap_new__HashMap_4uint_6string()
    values__0 = inline9
    var inline7 string = "stored"
    hashmap_set__HashMap_4uint_6string(values__0, result__0, inline7)
    var mtmp0 Option__string
    var inline5 uint = 42
    var inline6 Option__string = hashmap_get__HashMap_4uint_6string(values__0, inline5)
    mtmp0 = inline6
    switch mtmp0._tag {
    case 0:
        var inline0 string = "missing"
        var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
        _goml_runtime_core_string_println(inline1)
        return struct{}{}
    case 1:
        var x0 string = mtmp0._p0
        var inline3 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(x0)
        _goml_runtime_core_string_println(inline3)
        return struct{}{}
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_trait__impl_i_ToString_i_usize_i_to__string(self__0 uint) string {
    var inline0 uint64 = uint64(uint(self__0))
    var inline1 string = decimal_string(inline0)
    return inline1
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func __goml_builtin_uint_to_string(value__0 uint) string {
    var t0 uint64 = uint64(uint(value__0))
    var t1 string = decimal_string(t0)
    return t1
}

func _goml_m_trait__impl_i_ToString_i_bool_i_to__string(self__0 bool) string {
    var t0 string = _goml_runtime_core_bool_to_string(self__0)
    return t0
}

func _goml_m_trait__impl_i_ToString_i_u64_i_to__string(self__0 uint64) string {
    var inline0 string = decimal_string(self__0)
    return inline0
}

func decimal_string(value__0 uint64) string {
    var t0 bool = value__0 == 0
    if t0 {
        return "0"
    } else {
        var reversed__0 *_goml_vec_uint8 = vec_with_capacity__Vec_5uint8(20)
        var remaining__0 uint64 = value__0
        Loop_loop0:
        for {
            var t10 bool = remaining__0 > 0
            if t10 {
                var t11 uint64 = remaining__0 % 10
                var t12 uint8 = uint8(uint64(t11))
                var t13 uint8 = t12 + 48
                vec_push__Vec_5uint8(reversed__0, t13)
                var compound_old1 uint64 = remaining__0
                var compound_value1 uint64 = 10
                var t14 uint64 = compound_old1 / compound_value1
                remaining__0 = t14
                continue
            } else {
                break Loop_loop0
            }
        }
        var t1 int
        var inline3 int = vec_len__Vec_5uint8(reversed__0)
        t1 = inline3
        var bytes__0 *_goml_vec_uint8 = vec_with_capacity__Vec_5uint8(t1)
        var offset__0 int = 0
        Loop_loop1:
        for {
            var t2 int
            var inline2 int = vec_len__Vec_5uint8(reversed__0)
            t2 = inline2
            var t3 bool = offset__0 < t2
            if t3 {
                var t4 int
                var inline1 int = vec_len__Vec_5uint8(reversed__0)
                t4 = inline1
                var t5 int = t4 - offset__0
                var t6 int = t5 - 1
                var t7 uint8 = vec_get__Vec_5uint8(reversed__0, t6)
                vec_push__Vec_5uint8(bytes__0, t7)
                var compound_old0 int = offset__0
                var compound_value0 int = 1
                var t8 int = compound_old0 + compound_value0
                offset__0 = t8
                continue
            } else {
                break Loop_loop1
            }
        }
        var mtmp0 Tuple2_4bool_6string = _goml_runtime_core_string_from_utf8(bytes__0)
        var x0 string = mtmp0._1
        return x0
    }
}

func main() {
    main0()
}
