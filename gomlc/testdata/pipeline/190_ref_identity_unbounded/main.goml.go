package main

import (
    _goml_os "os"
    _goml_reflect "reflect"
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

type ref_Node_x struct {
    value Node
}

func ref__Ref_4Node(value Node) *ref_Node_x {
    return &ref_Node_x{
        value: value,
    }
}

func ref_set__Ref_4Node(reference *ref_Node_x, value Node) struct{} {
    reference.value = value
    return struct{}{}
}

func ptr_eq__Ref_4Node(a *ref_Node_x, b *ref_Node_x) bool {
    return a == b
}

func ptr_hash__Ref_4Node(reference *ref_Node_x) uint64 {
    return uint64(_goml_reflect.ValueOf(reference).Pointer())
}

type hashmap_Ref_4Node_string_x_entry struct {
    active bool
    key *ref_Node_x
    value string
}

type hashmap_Ref_4Node_string_x struct {
    buckets map[uint64][]hashmap_Ref_4Node_string_x_entry
    hashes []uint64
    len int
}

func hashmap_new__HashMap_9Ref_4Node_6string() *hashmap_Ref_4Node_string_x {
    return &hashmap_Ref_4Node_string_x{
        buckets: make(map[uint64][]hashmap_Ref_4Node_string_x_entry),
        len: 0,
        hashes: nil,
    }
}

func hashmap_len__HashMap_9Ref_4Node_6string(m *hashmap_Ref_4Node_string_x) int {
    if m == nil {
        return 0
    }
    return m.len
}

func hashmap_lookup__HashMap_9Ref_4Node_6string(m *hashmap_Ref_4Node_string_x, key *ref_Node_x) (string, bool, int, uint64) {
    if m == nil {
        var zero string
        return zero, false, -1, 0
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_Ref_l_Node_r__i_hash(key)
    var bucket []hashmap_Ref_4Node_string_x_entry = m.buckets[h]
    var i int = 0
    var reuse_index int = -1
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_Ref_4Node_string_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_Ref_l_Node_r__i_eq(entry.key, key) {
            return entry.value, true, i, h
        }
        if !entry.active && reuse_index < 0 {
            reuse_index = i
        }
        i = i + 1
    }
    var zero string
    return zero, false, reuse_index, h
}

func hashmap_get__HashMap_9Ref_4Node_6string(m *hashmap_Ref_4Node_string_x, key *ref_Node_x) Option__string {
    var value string
    var ok bool
    value, ok, _, _ = hashmap_lookup__HashMap_9Ref_4Node_6string(m, key)
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

func hashmap_set__HashMap_9Ref_4Node_6string(m *hashmap_Ref_4Node_string_x, key *ref_Node_x, value string) struct{} {
    var reuse_index int = -1
    if m == nil {
        return struct{}{}
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_Ref_l_Node_r__i_hash(key)
    var bucket []hashmap_Ref_4Node_string_x_entry = m.buckets[h]
    if len(bucket) == 0 {
        m.hashes = append(m.hashes, h)
    }
    var i int = 0
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_Ref_4Node_string_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_Ref_l_Node_r__i_eq(entry.key, key) {
            bucket[i].value = value
            return struct{}{}
        }
        if !entry.active && reuse_index < 0 {
            reuse_index = i
        }
        i = i + 1
    }
    if reuse_index >= 0 {
        bucket[reuse_index] = hashmap_Ref_4Node_string_x_entry{
            active: true,
            key: key,
            value: value,
        }
        m.len = m.len + 1
        return struct{}{}
    }
    bucket = append(bucket, hashmap_Ref_4Node_string_x_entry{
        active: true,
        key: key,
        value: value,
    })
    m.buckets[h] = bucket
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

type Node struct {
    value int32
}

type Ordering uint8

type Option__string struct {
    _p0 string
    _tag uint8
}

func print_lookup(map__0 *hashmap_Ref_4Node_string_x, key__0 *ref_Node_x) struct{} {
    var mtmp0 Option__string
    var inline5 Option__string = hashmap_get__HashMap_9Ref_4Node_6string(map__0, key__0)
    mtmp0 = inline5
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

func main0() struct{} {
    var t0 Node = Node{
        value: 1,
    }
    var first__0 *ref_Node_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__Node(t0)
    var t1 Node = Node{
        value: 1,
    }
    var distinct__0 *ref_Node_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__Node(t1)
    var t2 bool = _goml_m_trait__impl_i_PartialEq_i_Ref_l_Node_r__i_eq(first__0, first__0)
    println__T_bool(t2)
    var t3 bool = _goml_m_trait__impl_i_PartialEq_i_Ref_l_Node_r__i_eq(first__0, distinct__0)
    println__T_bool(t3)
    var t4 uint64 = _goml_m_trait__impl_i_Hash_i_Ref_l_Node_r__i_hash(first__0)
    var t5 uint64 = _goml_m_trait__impl_i_Hash_i_Ref_l_Node_r__i_hash(first__0)
    var t6 bool = t4 == t5
    var inline25 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t6)
    _goml_runtime_core_string_println(inline25)
    var map__0 *hashmap_Ref_4Node_string_x
    var inline24 *hashmap_Ref_4Node_string_x = hashmap_new__HashMap_9Ref_4Node_6string()
    map__0 = inline24
    var inline22 string = "first"
    hashmap_set__HashMap_9Ref_4Node_6string(map__0, first__0, inline22)
    print_lookup(map__0, first__0)
    var inline18 Option__string = _goml_m_inherent_i_HashMap_i_H_h9ede9e5e4b19497ae51b9af572b3f34a_r_____V__string(map__0, distinct__0)
    switch inline18._tag {
    case 0:
        println__T_string("missing")
    case 1:
        var inline20 string = inline18._p0
        println__T_string(inline20)
    default:
        panic("non-exhaustive match")
    }
    var t7 Node = Node{
        value: 2,
    }
    ref_set__Ref_4Node(first__0, t7)
    var inline13 Option__string = _goml_m_inherent_i_HashMap_i_H_h9ede9e5e4b19497ae51b9af572b3f34a_r_____V__string(map__0, first__0)
    switch inline13._tag {
    case 0:
        println__T_string("missing")
    case 1:
        var inline15 string = inline13._p0
        println__T_string(inline15)
    default:
        panic("non-exhaustive match")
    }
    var inline11 string = "distinct"
    hashmap_set__HashMap_9Ref_4Node_6string(map__0, distinct__0, inline11)
    var t8 int
    var inline10 int = hashmap_len__HashMap_9Ref_4Node_6string(map__0)
    t8 = inline10
    var inline8 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t8)
    _goml_runtime_core_string_println(inline8)
    var inline4 Option__string = _goml_m_inherent_i_HashMap_i_H_h9ede9e5e4b19497ae51b9af572b3f34a_r_____V__string(map__0, first__0)
    switch inline4._tag {
    case 0:
        println__T_string("missing")
    case 1:
        var inline6 string = inline4._p0
        println__T_string(inline6)
    default:
        panic("non-exhaustive match")
    }
    var inline0 Option__string = _goml_m_inherent_i_HashMap_i_H_h9ede9e5e4b19497ae51b9af572b3f34a_r_____V__string(map__0, distinct__0)
    switch inline0._tag {
    case 0:
        println__T_string("missing")
        return struct{}{}
    case 1:
        var inline2 string = inline0._p0
        println__T_string(inline2)
        return struct{}{}
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_inherent_i_HashMap_i_H_h9ede9e5e4b19497ae51b9af572b3f34a_r_____V__string(self__0 *hashmap_Ref_4Node_string_x, key__0 *ref_Node_x) Option__string {
    var t0 Option__string = hashmap_get__HashMap_9Ref_4Node_6string(self__0, key__0)
    return t0
}

func println__T_string(value__0 string) struct{} {
    var t0 string
    t0 = value__0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__Node(value__0 Node) *ref_Node_x {
    var t0 *ref_Node_x = ref__Ref_4Node(value__0)
    return t0
}

func println__T_bool(value__0 bool) struct{} {
    var t0 string
    var inline0 string = _goml_runtime_core_bool_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_trait__impl_i_PartialEq_i_Ref_l_Node_r__i_eq(self__0 *ref_Node_x, other__0 *ref_Node_x) bool {
    var t0 bool = ptr_eq__Ref_4Node(self__0, other__0)
    return t0
}

func _goml_m_trait__impl_i_Hash_i_Ref_l_Node_r__i_hash(self__0 *ref_Node_x) uint64 {
    var t0 uint64 = ptr_hash__Ref_4Node(self__0)
    return t0
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func _goml_m_trait__impl_i_ToString_i_bool_i_to__string(self__0 bool) string {
    var t0 string = _goml_runtime_core_bool_to_string(self__0)
    return t0
}

func _goml_m_trait__impl_i_ToString_i_isize_i_to__string(self__0 int) string {
    var inline0 int64 = int64(int(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
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
