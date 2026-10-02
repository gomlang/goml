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

func _goml_runtime_core_int32_hash(x int32) uint64 {
    return uint64(x)
}

func _goml_runtime_core_string_println(s string) struct{} {
    _goml_os.Stdout.WriteString(s + "\n")
    return struct{}{}
}

type _goml_vec_int struct {
    items []int
}

func vec_new__Vec_3int() *_goml_vec_int {
    return &_goml_vec_int{
        items: nil,
    }
}

func vec_push__Vec_3int(vec *_goml_vec_int, elem int) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_get__Vec_3int(vec *_goml_vec_int, index int) int {
    return vec.items[index]
}

func vec_len__Vec_3int(vec *_goml_vec_int) int {
    return int(len(vec.items))
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

type hashmap_Key_int32_x_entry struct {
    active bool
    key Key
    value int32
}

type hashmap_Key_int32_x struct {
    buckets map[uint64][]hashmap_Key_int32_x_entry
    hashes []uint64
    len int
}

func hashmap_new__HashMap_3Key_5int32() *hashmap_Key_int32_x {
    return &hashmap_Key_int32_x{
        buckets: make(map[uint64][]hashmap_Key_int32_x_entry),
        len: 0,
        hashes: nil,
    }
}

func hashmap_len__HashMap_3Key_5int32(m *hashmap_Key_int32_x) int {
    if m == nil {
        return 0
    }
    return m.len
}

func hashmap_lookup__HashMap_3Key_5int32(m *hashmap_Key_int32_x, key Key) (int32, bool, int, uint64) {
    if m == nil {
        var zero int32
        return zero, false, -1, 0
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_Key_i_hash(key)
    var bucket []hashmap_Key_int32_x_entry = m.buckets[h]
    var i int = 0
    var reuse_index int = -1
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_Key_int32_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_Key_i_eq(entry.key, key) {
            return entry.value, true, i, h
        }
        if !entry.active && reuse_index < 0 {
            reuse_index = i
        }
        i = i + 1
    }
    var zero int32
    return zero, false, reuse_index, h
}

func hashmap_get__HashMap_3Key_5int32(m *hashmap_Key_int32_x, key Key) Option__i32 {
    var value int32
    var ok bool
    value, ok, _, _ = hashmap_lookup__HashMap_3Key_5int32(m, key)
    if ok {
        return Option__i32(uint64(int64(value) + 2147483648) + 1)
    }
    return Option__i32(0)
}

func hashmap_set__HashMap_3Key_5int32(m *hashmap_Key_int32_x, key Key, value int32) struct{} {
    var reuse_index int = -1
    if m == nil {
        return struct{}{}
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_Key_i_hash(key)
    var bucket []hashmap_Key_int32_x_entry = m.buckets[h]
    if len(bucket) == 0 {
        m.hashes = append(m.hashes, h)
    }
    var i int = 0
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_Key_int32_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_Key_i_eq(entry.key, key) {
            bucket[i].value = value
            return struct{}{}
        }
        if !entry.active && reuse_index < 0 {
            reuse_index = i
        }
        i = i + 1
    }
    if reuse_index >= 0 {
        bucket[reuse_index] = hashmap_Key_int32_x_entry{
            active: true,
            key: key,
            value: value,
        }
        m.len = m.len + 1
        return struct{}{}
    }
    bucket = append(bucket, hashmap_Key_int32_x_entry{
        active: true,
        key: key,
        value: value,
    })
    m.buckets[h] = bucket
    m.len = m.len + 1
    return struct{}{}
}

func hashmap_remove__HashMap_3Key_5int32(m *hashmap_Key_int32_x, key Key) struct{} {
    if m == nil {
        return struct{}{}
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_Key_i_hash(key)
    var bucket []hashmap_Key_int32_x_entry = m.buckets[h]
    var i int = 0
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_Key_int32_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_Key_i_eq(entry.key, key) {
            var zero hashmap_Key_int32_x_entry
            bucket[i] = zero
            m.len = m.len - 1
            return struct{}{}
        }
        i = i + 1
    }
    return struct{}{}
}

func hashmap_contains__HashMap_3Key_5int32(m *hashmap_Key_int32_x, key Key) bool {
    if m == nil {
        return false
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_Key_i_hash(key)
    var bucket []hashmap_Key_int32_x_entry = m.buckets[h]
    var i int = 0
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_Key_int32_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_Key_i_eq(entry.key, key) {
            return true
        }
        i = i + 1
    }
    return false
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

type Key uint64

type Option__i32 uint64

func _goml_m_trait__impl_i_Hash_i_Key_i_hash(self__0 Key) uint64 {
    switch self__0 != Key(0) {
    case false:
        var t0_source int = 0
        var t0 uint64 = uint64(int(t0_source))
        var t1 uint64 = t0 + 14695981039346656037
        var h__0 uint64 = t1 + 1
        return h__0
    case true:
        var x0 int32 = int32(int64(uint64(self__0) - 1) - 2147483648)
        var t2_source int = 0
        var t2 uint64 = uint64(int(t2_source))
        var t3 uint64 = t2 + 14695981039346656037
        var h__1 uint64 = t3 + 2
        var t4_source int = 0
        var t4 uint64 = uint64(int(t4_source))
        var t5 uint64 = t4 + 1099511628211
        var t6 uint64 = h__1 * t5
        var t7 uint64
        var inline0 uint64 = _goml_runtime_core_int32_hash(x0)
        t7 = inline0
        var h__2 uint64 = t6 + t7
        return h__2
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_trait__impl_i_PartialEq_i_Key_i_eq(self__0 Key, other__0 Key) bool {
    switch other__0 != Key(0) {
    case false:
        switch self__0 != Key(0) {
        case false:
            return true
        default:
            return false
        }
    case true:
        var x0 int32 = int32(int64(uint64(other__0) - 1) - 2147483648)
        switch self__0 != Key(0) {
        case true:
            var x1 int32 = int32(int64(uint64(self__0) - 1) - 2147483648)
            var inline0 bool = x1 == x0
            return inline0
        default:
            return false
        }
    default:
        panic("non-exhaustive match")
    }
}

func main0() struct{} {
    var v__0 *_goml_vec_int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_new____T__isize()
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_push____T__isize(v__0, 10)
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_push____T__isize(v__0, 20)
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_push____T__isize(v__0, 30)
    var t0 int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_get____T__isize(v__0, 0)
    println__T_isize(t0)
    var t1 int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_get____T__isize(v__0, 1)
    println__T_isize(t1)
    var t2 int
    var inline27 int = 2
    var inline28 int = vec_get__Vec_3int(v__0, inline27)
    t2 = inline28
    var inline25 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t2)
    _goml_runtime_core_string_println(inline25)
    var t3 int
    var inline24 int = vec_len__Vec_3int(v__0)
    t3 = inline24
    var inline22 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t3)
    _goml_runtime_core_string_println(inline22)
    var m__0 *hashmap_Key_int32_x
    var inline21 *hashmap_Key_int32_x = hashmap_new__HashMap_3Key_5int32()
    m__0 = inline21
    var inline19 int32 = 10
    hashmap_set__HashMap_3Key_5int32(m__0, Key(0), inline19)
    var t4 Key = Key(uint64(int64(1) + 2147483648) + 1)
    var inline17 int32 = 20
    hashmap_set__HashMap_3Key_5int32(m__0, t4, inline17)
    var t5 int
    var inline16 int = hashmap_len__HashMap_3Key_5int32(m__0)
    t5 = inline16
    var inline14 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t5)
    _goml_runtime_core_string_println(inline14)
    var t6 Option__i32
    var inline13 Option__i32 = hashmap_get__HashMap_3Key_5int32(m__0, Key(0))
    t6 = inline13
    switch t6 != Option__i32(0) {
    case false:
        println__T_string("none")
    case true:
        var inline11 int32 = int32(int64(uint64(t6) - 1) - 2147483648)
        println__T_i32(inline11)
    default:
        panic("non-exhaustive match")
    }
    var t7 Key = Key(uint64(int64(1) + 2147483648) + 1)
    var t8 bool
    var inline9 bool = hashmap_contains__HashMap_3Key_5int32(m__0, t7)
    t8 = inline9
    var inline7 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t8)
    _goml_runtime_core_string_println(inline7)
    var t9 Key = Key(uint64(int64(1) + 2147483648) + 1)
    hashmap_remove__HashMap_3Key_5int32(m__0, t9)
    var t10 Key = Key(uint64(int64(1) + 2147483648) + 1)
    var t11 bool
    var inline5 bool = hashmap_contains__HashMap_3Key_5int32(m__0, t10)
    t11 = inline5
    var inline3 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t11)
    _goml_runtime_core_string_println(inline3)
    var t12 int
    var inline2 int = hashmap_len__HashMap_3Key_5int32(m__0)
    t12 = inline2
    var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t12)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func println__T_string(value__0 string) struct{} {
    var t0 string
    t0 = value__0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func println__T_i32(value__0 int32) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_int32_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_new____T__isize() *_goml_vec_int {
    var t0 *_goml_vec_int = vec_new__Vec_3int()
    return t0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_push____T__isize(self__0 *_goml_vec_int, elem__0 int) struct{} {
    vec_push__Vec_3int(self__0, elem__0)
    return struct{}{}
}

func println__T_isize(value__0 int) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_int_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_get____T__isize(self__0 *_goml_vec_int, index__0 int) int {
    var t0 int = vec_get__Vec_3int(self__0, index__0)
    return t0
}

func _goml_m_trait__impl_i_ToString_i_isize_i_to__string(self__0 int) string {
    var inline0 int64 = int64(int(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func _goml_m_trait__impl_i_ToString_i_bool_i_to__string(self__0 bool) string {
    var t0 string = _goml_runtime_core_bool_to_string(self__0)
    return t0
}

func __goml_builtin_int32_to_string(value__0 int32) string {
    var t0 int64 = int64(int32(value__0))
    var inline0 bool = t0 < 0
    if inline0 {
        var inline1 uint64 = uint64(int64(t0))
        var inline2 uint64 = 0 - inline1
        var inline3 string = decimal_string(inline2)
        var inline4 string = "-" + inline3
        return inline4
    } else {
        var inline5 uint64 = uint64(int64(t0))
        var inline6 string = decimal_string(inline5)
        return inline6
    }
}

func __goml_builtin_int_to_string(value__0 int) string {
    var t0 int64 = int64(int(value__0))
    var inline0 bool = t0 < 0
    if inline0 {
        var inline1 uint64 = uint64(int64(t0))
        var inline2 uint64 = 0 - inline1
        var inline3 string = decimal_string(inline2)
        var inline4 string = "-" + inline3
        return inline4
    } else {
        var inline5 uint64 = uint64(int64(t0))
        var inline6 string = decimal_string(inline5)
        return inline6
    }
}

func signed_decimal_string(value__0 int64) string {
    var t0 bool = value__0 < 0
    if t0 {
        var t1 uint64 = uint64(int64(value__0))
        var t2 uint64 = 0 - t1
        var t3 string = decimal_string(t2)
        var t4 string = "-" + t3
        return t4
    } else {
        var t5 uint64 = uint64(int64(value__0))
        var t6 string = decimal_string(t5)
        return t6
    }
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
