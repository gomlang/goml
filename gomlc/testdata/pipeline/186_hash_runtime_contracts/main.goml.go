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

type ref_int32_x struct {
    value int32
}

func ref__Ref_5int32(value int32) *ref_int32_x {
    return &ref_int32_x{
        value: value,
    }
}

func ref_get__Ref_5int32(reference *ref_int32_x) int32 {
    return reference.value
}

func ref_set__Ref_5int32(reference *ref_int32_x, value int32) struct{} {
    reference.value = value
    return struct{}{}
}

type ref_CollisionKey_x struct {
    value CollisionKey
}

func ref__Ref_12CollisionKey(value CollisionKey) *ref_CollisionKey_x {
    return &ref_CollisionKey_x{
        value: value,
    }
}

func ref_set__Ref_12CollisionKey(reference *ref_CollisionKey_x, value CollisionKey) struct{} {
    reference.value = value
    return struct{}{}
}

func ptr_eq__Ref_12CollisionKey(a *ref_CollisionKey_x, b *ref_CollisionKey_x) bool {
    return a == b
}

func ptr_hash__Ref_12CollisionKey(reference *ref_CollisionKey_x) uint64 {
    return uint64(_goml_reflect.ValueOf(reference).Pointer())
}

type hashmap_CollisionKey_int32_x_entry struct {
    active bool
    key CollisionKey
    value int32
}

type hashmap_CollisionKey_int32_x struct {
    buckets map[uint64][]hashmap_CollisionKey_int32_x_entry
    hashes []uint64
    len int
}

func hashmap_new__HashMap_12CollisionKey_5int32() *hashmap_CollisionKey_int32_x {
    return &hashmap_CollisionKey_int32_x{
        buckets: make(map[uint64][]hashmap_CollisionKey_int32_x_entry),
        len: 0,
        hashes: nil,
    }
}

func hashmap_len__HashMap_12CollisionKey_5int32(m *hashmap_CollisionKey_int32_x) int {
    if m == nil {
        return 0
    }
    return m.len
}

func hashmap_lookup__HashMap_12CollisionKey_5int32(m *hashmap_CollisionKey_int32_x, key CollisionKey) (int32, bool, int, uint64) {
    if m == nil {
        var zero int32
        return zero, false, -1, 0
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_CollisionKey_i_hash(key)
    var bucket []hashmap_CollisionKey_int32_x_entry = m.buckets[h]
    var i int = 0
    var reuse_index int = -1
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_CollisionKey_int32_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_CollisionKey_i_eq(entry.key, key) {
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

func hashmap_get__HashMap_12CollisionKey_5int32(m *hashmap_CollisionKey_int32_x, key CollisionKey) Option__i32 {
    var value int32
    var ok bool
    value, ok, _, _ = hashmap_lookup__HashMap_12CollisionKey_5int32(m, key)
    if ok {
        return Option__i32(uint64(int64(value) + 2147483648) + 1)
    }
    return Option__i32(0)
}

func hashmap_set__HashMap_12CollisionKey_5int32(m *hashmap_CollisionKey_int32_x, key CollisionKey, value int32) struct{} {
    var reuse_index int = -1
    if m == nil {
        return struct{}{}
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_CollisionKey_i_hash(key)
    var bucket []hashmap_CollisionKey_int32_x_entry = m.buckets[h]
    if len(bucket) == 0 {
        m.hashes = append(m.hashes, h)
    }
    var i int = 0
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_CollisionKey_int32_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_CollisionKey_i_eq(entry.key, key) {
            bucket[i].value = value
            return struct{}{}
        }
        if !entry.active && reuse_index < 0 {
            reuse_index = i
        }
        i = i + 1
    }
    if reuse_index >= 0 {
        bucket[reuse_index] = hashmap_CollisionKey_int32_x_entry{
            active: true,
            key: key,
            value: value,
        }
        m.len = m.len + 1
        return struct{}{}
    }
    bucket = append(bucket, hashmap_CollisionKey_int32_x_entry{
        active: true,
        key: key,
        value: value,
    })
    m.buckets[h] = bucket
    m.len = m.len + 1
    return struct{}{}
}

func hashmap_remove__HashMap_12CollisionKey_5int32(m *hashmap_CollisionKey_int32_x, key CollisionKey) struct{} {
    if m == nil {
        return struct{}{}
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_CollisionKey_i_hash(key)
    var bucket []hashmap_CollisionKey_int32_x_entry = m.buckets[h]
    var i int = 0
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_CollisionKey_int32_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_CollisionKey_i_eq(entry.key, key) {
            var zero hashmap_CollisionKey_int32_x_entry
            bucket[i] = zero
            m.len = m.len - 1
            return struct{}{}
        }
        i = i + 1
    }
    return struct{}{}
}

type hashmap_Ref_12CollisionKey_string_x_entry struct {
    active bool
    key *ref_CollisionKey_x
    value string
}

type hashmap_Ref_12CollisionKey_string_x struct {
    buckets map[uint64][]hashmap_Ref_12CollisionKey_string_x_entry
    hashes []uint64
    len int
}

func hashmap_new__HashMap_18Ref_12CollisionKey_6string() *hashmap_Ref_12CollisionKey_string_x {
    return &hashmap_Ref_12CollisionKey_string_x{
        buckets: make(map[uint64][]hashmap_Ref_12CollisionKey_string_x_entry),
        len: 0,
        hashes: nil,
    }
}

func hashmap_lookup__HashMap_18Ref_12CollisionKey_6string(m *hashmap_Ref_12CollisionKey_string_x, key *ref_CollisionKey_x) (string, bool, int, uint64) {
    if m == nil {
        var zero string
        return zero, false, -1, 0
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_Ref_l_CollisionKey_r__i_hash(key)
    var bucket []hashmap_Ref_12CollisionKey_string_x_entry = m.buckets[h]
    var i int = 0
    var reuse_index int = -1
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_Ref_12CollisionKey_string_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_Ref_l_CollisionKey_r__i_eq(entry.key, key) {
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

func hashmap_get__HashMap_18Ref_12CollisionKey_6string(m *hashmap_Ref_12CollisionKey_string_x, key *ref_CollisionKey_x) Option__string {
    var value string
    var ok bool
    value, ok, _, _ = hashmap_lookup__HashMap_18Ref_12CollisionKey_6string(m, key)
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

func hashmap_set__HashMap_18Ref_12CollisionKey_6string(m *hashmap_Ref_12CollisionKey_string_x, key *ref_CollisionKey_x, value string) struct{} {
    var reuse_index int = -1
    if m == nil {
        return struct{}{}
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_Ref_l_CollisionKey_r__i_hash(key)
    var bucket []hashmap_Ref_12CollisionKey_string_x_entry = m.buckets[h]
    if len(bucket) == 0 {
        m.hashes = append(m.hashes, h)
    }
    var i int = 0
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_Ref_12CollisionKey_string_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_Ref_l_CollisionKey_r__i_eq(entry.key, key) {
            bucket[i].value = value
            return struct{}{}
        }
        if !entry.active && reuse_index < 0 {
            reuse_index = i
        }
        i = i + 1
    }
    if reuse_index >= 0 {
        bucket[reuse_index] = hashmap_Ref_12CollisionKey_string_x_entry{
            active: true,
            key: key,
            value: value,
        }
        m.len = m.len + 1
        return struct{}{}
    }
    bucket = append(bucket, hashmap_Ref_12CollisionKey_string_x_entry{
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

type FloatKey struct {
    value float64
}

type CollisionKey struct {
    value int32
}

type Ordering uint8

const (
    Less Ordering = 0
    Equal Ordering = 1
    Greater Ordering = 2
)

type Option__Ordering struct {
    _p0 Ordering
    _tag uint8
}

type Option__string struct {
    _p0 string
    _tag uint8
}

type Option__i32 uint64

func _goml_m_trait__impl_i_PartialEq_i_CollisionKey_i_eq(self__0 CollisionKey, other__0 CollisionKey) bool {
    var t0 int32 = self__0.value
    var t1 int32 = other__0.value
    var t2 bool = t0 == t1
    return t2
}

func _goml_m_trait__impl_i_Hash_i_CollisionKey_i_hash(self__0 CollisionKey) uint64 {
    return 1
}

func print_opt_int(value__0 Option__i32) struct{} {
    switch value__0 != Option__i32(0) {
    case false:
        var inline0 string = "none"
        var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
        _goml_runtime_core_string_println(inline1)
        return struct{}{}
    case true:
        var x0 int32 = int32(int64(uint64(value__0) - 1) - 2147483648)
        var inline3 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(x0)
        _goml_runtime_core_string_println(inline3)
        return struct{}{}
    default:
        panic("non-exhaustive match")
    }
}

func float_comparison_contracts() struct{} {
    var zero32__0 float32 = 0
    var negative_zero32__0 float32 = -zero32__0
    var t0 bool = zero32__0 == negative_zero32__0
    var t1 string
    var inline23 string = _goml_runtime_core_bool_to_string(t0)
    t1 = inline23
    var inline21 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t1)
    _goml_runtime_core_string_println(inline21)
    var zero64__0 float64 = 0
    var negative_zero64__0 float64 = -zero64__0
    var t2 bool = zero64__0 == negative_zero64__0
    var t3 string
    var inline20 string = _goml_runtime_core_bool_to_string(t2)
    t3 = inline20
    var inline18 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t3)
    _goml_runtime_core_string_println(inline18)
    var t4 bool
    var inline17 bool = _goml_m_trait__impl_i_PartialEq_i_f64_i_eq(zero64__0, negative_zero64__0)
    t4 = inline17
    var t5 string
    var inline16 string = _goml_runtime_core_bool_to_string(t4)
    t5 = inline16
    var inline14 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t5)
    _goml_runtime_core_string_println(inline14)
    var nan__0 float64 = zero64__0 / zero64__0
    var t6 bool = nan__0 == nan__0
    var t7 string
    var inline13 string = _goml_runtime_core_bool_to_string(t6)
    t7 = inline13
    var inline11 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t7)
    _goml_runtime_core_string_println(inline11)
    var t8 Option__Ordering
    var inline5 bool = nan__0 < nan__0
    if inline5 {
        var inline6 Option__Ordering = Option__Ordering{
            _p0: Less,
            _tag: 1,
        }
        t8 = inline6
    } else {
        var inline7 bool = nan__0 > nan__0
        if inline7 {
            var inline8 Option__Ordering = Option__Ordering{
                _p0: Greater,
                _tag: 1,
            }
            t8 = inline8
        } else {
            var inline9 bool = nan__0 == nan__0
            if inline9 {
                var inline10 Option__Ordering = Option__Ordering{
                    _p0: Equal,
                    _tag: 1,
                }
                t8 = inline10
            } else {
                t8 = Option__Ordering{
                    _tag: 0,
                }
            }
        }
    }
    var t9 bool
    var inline3 bool = _goml_m_inherent_i_Option_i_Option_l_T_r__i_is__some____T__Ordering(t8)
    var inline4 bool = !inline3
    t9 = inline4
    var t10 string
    var inline2 string = _goml_runtime_core_bool_to_string(t9)
    t10 = inline2
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t10)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func collision_contracts() struct{} {
    var values__0 *hashmap_CollisionKey_int32_x = _goml_m_inherent_i_HashMap_i_H_hc55bb71e9219d0c59c91622ae099ea85_onKey____V__i32()
    var t0 CollisionKey = CollisionKey{
        value: 1,
    }
    _goml_m_inherent_i_HashMap_i_H_h13877db3209e2e8534a556971ced5ab2_onKey____V__i32(values__0, t0, 10)
    var t1 CollisionKey = CollisionKey{
        value: 2,
    }
    _goml_m_inherent_i_HashMap_i_H_h13877db3209e2e8534a556971ced5ab2_onKey____V__i32(values__0, t1, 20)
    var t2 CollisionKey = CollisionKey{
        value: 3,
    }
    _goml_m_inherent_i_HashMap_i_H_h13877db3209e2e8534a556971ced5ab2_onKey____V__i32(values__0, t2, 30)
    var t3 CollisionKey = CollisionKey{
        value: 2,
    }
    _goml_m_inherent_i_HashMap_i_H_h1189b7c51290244a02a1a6d496e4da69_onKey____V__i32(values__0, t3)
    var t4 CollisionKey = CollisionKey{
        value: 1,
    }
    var t5 Option__i32 = _goml_m_inherent_i_HashMap_i_H_h7d750efaf8a1803c39420b89990faf20_onKey____V__i32(values__0, t4)
    print_opt_int(t5)
    var t6 CollisionKey = CollisionKey{
        value: 2,
    }
    var t7 Option__i32 = _goml_m_inherent_i_HashMap_i_H_h7d750efaf8a1803c39420b89990faf20_onKey____V__i32(values__0, t6)
    print_opt_int(t7)
    var t8 CollisionKey = CollisionKey{
        value: 3,
    }
    var t9 Option__i32 = _goml_m_inherent_i_HashMap_i_H_h7d750efaf8a1803c39420b89990faf20_onKey____V__i32(values__0, t8)
    print_opt_int(t9)
    var t10 CollisionKey = CollisionKey{
        value: 4,
    }
    _goml_m_inherent_i_HashMap_i_H_h13877db3209e2e8534a556971ced5ab2_onKey____V__i32(values__0, t10, 40)
    var t11 int = _goml_m_inherent_i_HashMap_i_H_h282dac09c2296c58cbcd9cfca496474b_onKey____V__i32(values__0)
    println__T_isize(t11)
    var t12 CollisionKey = CollisionKey{
        value: 4,
    }
    var t13 Option__i32
    var inline26 Option__i32 = hashmap_get__HashMap_12CollisionKey_5int32(values__0, t12)
    t13 = inline26
    print_opt_int(t13)
    var t14 CollisionKey = CollisionKey{
        value: 4,
    }
    var inline24 int32 = 41
    hashmap_set__HashMap_12CollisionKey_5int32(values__0, t14, inline24)
    var t15 int
    var inline23 int = hashmap_len__HashMap_12CollisionKey_5int32(values__0)
    t15 = inline23
    var inline21 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t15)
    _goml_runtime_core_string_println(inline21)
    var t16 CollisionKey = CollisionKey{
        value: 4,
    }
    var t17 Option__i32
    var inline20 Option__i32 = hashmap_get__HashMap_12CollisionKey_5int32(values__0, t16)
    t17 = inline20
    switch t17 != Option__i32(0) {
    case false:
        println__T_string("none")
    case true:
        var inline18 int32 = int32(int64(uint64(t17) - 1) - 2147483648)
        println__T_i32(inline18)
    default:
        panic("non-exhaustive match")
    }
    var t18 CollisionKey = CollisionKey{
        value: 99,
    }
    hashmap_remove__HashMap_12CollisionKey_5int32(values__0, t18)
    var t19 CollisionKey = CollisionKey{
        value: 99,
    }
    hashmap_remove__HashMap_12CollisionKey_5int32(values__0, t19)
    var t20 int
    var inline14 int = hashmap_len__HashMap_12CollisionKey_5int32(values__0)
    t20 = inline14
    var inline12 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t20)
    _goml_runtime_core_string_println(inline12)
    var index__0 *ref_int32_x
    var inline10 int32 = 0
    var inline11 *ref_int32_x = ref__Ref_5int32(inline10)
    index__0 = inline11
    Loop_loop0:
    for {
        var t22 int32
        var inline9 int32 = ref_get__Ref_5int32(index__0)
        t22 = inline9
        var t23 bool = t22 < 2000
        if t23 {
            var t24 int32
            var inline8 int32 = ref_get__Ref_5int32(index__0)
            t24 = inline8
            var t25 int32 = 1000 + t24
            var key__0 CollisionKey = CollisionKey{
                value: t25,
            }
            var t26 int32
            var inline7 int32 = ref_get__Ref_5int32(index__0)
            t26 = inline7
            hashmap_set__HashMap_12CollisionKey_5int32(values__0, key__0, t26)
            hashmap_remove__HashMap_12CollisionKey_5int32(values__0, key__0)
            var t27 int32
            var inline4 int32 = ref_get__Ref_5int32(index__0)
            t27 = inline4
            var t28 int32 = t27 + 1
            ref_set__Ref_5int32(index__0, t28)
            continue
        } else {
            break Loop_loop0
        }
    }
    var t21 int
    var inline2 int = hashmap_len__HashMap_12CollisionKey_5int32(values__0)
    t21 = inline2
    var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t21)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func reference_contracts() struct{} {
    var values__0 *hashmap_Ref_12CollisionKey_string_x = _goml_m_inherent_i_HashMap_i_H_h88509d3b0dfc2392ba920998ecb79184_r_____V__string()
    var t0 CollisionKey = CollisionKey{
        value: 1,
    }
    var key__0 *ref_CollisionKey_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__CollisionKey(t0)
    var t1 CollisionKey = CollisionKey{
        value: 1,
    }
    var equal_value__0 *ref_CollisionKey_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__CollisionKey(t1)
    var inline23 string = "identity"
    hashmap_set__HashMap_18Ref_12CollisionKey_6string(values__0, key__0, inline23)
    var t2 bool
    var inline22 bool = ptr_eq__Ref_12CollisionKey(key__0, key__0)
    t2 = inline22
    var inline20 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t2)
    _goml_runtime_core_string_println(inline20)
    var t3 bool
    var inline19 bool = ptr_eq__Ref_12CollisionKey(key__0, equal_value__0)
    t3 = inline19
    var inline17 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t3)
    _goml_runtime_core_string_println(inline17)
    var t4 uint64
    var inline16 uint64 = ptr_hash__Ref_12CollisionKey(key__0)
    t4 = inline16
    var t5 uint64
    var inline15 uint64 = ptr_hash__Ref_12CollisionKey(key__0)
    t5 = inline15
    var t6 bool = t4 == t5
    var inline13 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t6)
    _goml_runtime_core_string_println(inline13)
    var t7 Option__string
    var inline12 Option__string = hashmap_get__HashMap_18Ref_12CollisionKey_6string(values__0, key__0)
    t7 = inline12
    switch t7._tag {
    case 0:
        println__T_string("none")
    case 1:
        var inline10 string = t7._p0
        println__T_string(inline10)
    default:
        panic("non-exhaustive match")
    }
    var t8 Option__string
    var inline8 Option__string = hashmap_get__HashMap_18Ref_12CollisionKey_6string(values__0, equal_value__0)
    t8 = inline8
    switch t8._tag {
    case 0:
        println__T_string("none")
    case 1:
        var inline6 string = t8._p0
        println__T_string(inline6)
    default:
        panic("non-exhaustive match")
    }
    var t9 CollisionKey = CollisionKey{
        value: 99,
    }
    ref_set__Ref_12CollisionKey(key__0, t9)
    var t10 Option__string
    var inline3 Option__string = hashmap_get__HashMap_18Ref_12CollisionKey_6string(values__0, key__0)
    t10 = inline3
    switch t10._tag {
    case 0:
        println__T_string("none")
        return struct{}{}
    case 1:
        var inline1 string = t10._p0
        println__T_string(inline1)
        return struct{}{}
    default:
        panic("non-exhaustive match")
    }
}

func main0() struct{} {
    float_comparison_contracts()
    collision_contracts()
    reference_contracts()
    return struct{}{}
}

func _goml_m_trait__impl_i_PartialEq_i_f64_i_eq(self__0 float64, other__0 float64) bool {
    var t0 bool = self__0 == other__0
    return t0
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

func _goml_m_trait__impl_i_ToString_i_bool_i_to__string(self__0 bool) string {
    var t0 string = _goml_runtime_core_bool_to_string(self__0)
    return t0
}

func _goml_m_inherent_i_HashMap_i_H_hc55bb71e9219d0c59c91622ae099ea85_onKey____V__i32() *hashmap_CollisionKey_int32_x {
    var t0 *hashmap_CollisionKey_int32_x = hashmap_new__HashMap_12CollisionKey_5int32()
    return t0
}

func _goml_m_inherent_i_HashMap_i_H_h13877db3209e2e8534a556971ced5ab2_onKey____V__i32(self__0 *hashmap_CollisionKey_int32_x, key__0 CollisionKey, value__0 int32) struct{} {
    hashmap_set__HashMap_12CollisionKey_5int32(self__0, key__0, value__0)
    return struct{}{}
}

func _goml_m_inherent_i_HashMap_i_H_h1189b7c51290244a02a1a6d496e4da69_onKey____V__i32(self__0 *hashmap_CollisionKey_int32_x, key__0 CollisionKey) struct{} {
    hashmap_remove__HashMap_12CollisionKey_5int32(self__0, key__0)
    return struct{}{}
}

func _goml_m_inherent_i_HashMap_i_H_h7d750efaf8a1803c39420b89990faf20_onKey____V__i32(self__0 *hashmap_CollisionKey_int32_x, key__0 CollisionKey) Option__i32 {
    var t0 Option__i32 = hashmap_get__HashMap_12CollisionKey_5int32(self__0, key__0)
    return t0
}

func println__T_isize(value__0 int) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_int_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_inherent_i_HashMap_i_H_h282dac09c2296c58cbcd9cfca496474b_onKey____V__i32(self__0 *hashmap_CollisionKey_int32_x) int {
    var t0 int = hashmap_len__HashMap_12CollisionKey_5int32(self__0)
    return t0
}

func _goml_m_inherent_i_HashMap_i_H_h88509d3b0dfc2392ba920998ecb79184_r_____V__string() *hashmap_Ref_12CollisionKey_string_x {
    var t0 *hashmap_Ref_12CollisionKey_string_x = hashmap_new__HashMap_18Ref_12CollisionKey_6string()
    return t0
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__CollisionKey(value__0 CollisionKey) *ref_CollisionKey_x {
    var t0 *ref_CollisionKey_x = ref__Ref_12CollisionKey(value__0)
    return t0
}

func _goml_m_trait__impl_i_PartialEq_i_Ref_l_CollisionKey_r__i_eq(self__0 *ref_CollisionKey_x, other__0 *ref_CollisionKey_x) bool {
    var t0 bool = ptr_eq__Ref_12CollisionKey(self__0, other__0)
    return t0
}

func _goml_m_trait__impl_i_Hash_i_Ref_l_CollisionKey_r__i_hash(self__0 *ref_CollisionKey_x) uint64 {
    var t0 uint64 = ptr_hash__Ref_12CollisionKey(self__0)
    return t0
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func _goml_m_trait__impl_i_ToString_i_i32_i_to__string(self__0 int32) string {
    var inline0 int64 = int64(int32(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func _goml_m_inherent_i_Option_i_Option_l_T_r__i_is__some____T__Ordering(self__0 Option__Ordering) bool {
    switch self__0._tag {
    case 0:
        return false
    case 1:
        return true
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_trait__impl_i_ToString_i_isize_i_to__string(self__0 int) string {
    var inline0 int64 = int64(int(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
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
