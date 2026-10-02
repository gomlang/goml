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

func _goml_runtime_core_int_hash(x int) uint64 {
    return uint64(x)
}

func _goml_runtime_core_string_println(s string) struct{} {
    _goml_os.Stdout.WriteString(s + "\n")
    return struct{}{}
}

type _goml_vec_uint32 struct {
    items []uint32
}

type hashmap_DerivedKey_string_x_entry struct {
    active bool
    key DerivedKey
    value string
}

type hashmap_DerivedKey_string_x struct {
    buckets map[uint64][]hashmap_DerivedKey_string_x_entry
    hashes []uint64
    len int
}

func hashmap_new__HashMap_10DerivedKey_6string() *hashmap_DerivedKey_string_x {
    return &hashmap_DerivedKey_string_x{
        buckets: make(map[uint64][]hashmap_DerivedKey_string_x_entry),
        len: 0,
        hashes: nil,
    }
}

func hashmap_lookup__HashMap_10DerivedKey_6string(m *hashmap_DerivedKey_string_x, key DerivedKey) (string, bool, int, uint64) {
    if m == nil {
        var zero string
        return zero, false, -1, 0
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_DerivedKey_i_hash(key)
    var bucket []hashmap_DerivedKey_string_x_entry = m.buckets[h]
    var i int = 0
    var reuse_index int = -1
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_DerivedKey_string_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_DerivedKey_i_eq(entry.key, key) {
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

func hashmap_get__HashMap_10DerivedKey_6string(m *hashmap_DerivedKey_string_x, key DerivedKey) Option__string {
    var value string
    var ok bool
    value, ok, _, _ = hashmap_lookup__HashMap_10DerivedKey_6string(m, key)
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

func hashmap_set__HashMap_10DerivedKey_6string(m *hashmap_DerivedKey_string_x, key DerivedKey, value string) struct{} {
    var reuse_index int = -1
    if m == nil {
        return struct{}{}
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_DerivedKey_i_hash(key)
    var bucket []hashmap_DerivedKey_string_x_entry = m.buckets[h]
    if len(bucket) == 0 {
        m.hashes = append(m.hashes, h)
    }
    var i int = 0
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_DerivedKey_string_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_DerivedKey_i_eq(entry.key, key) {
            bucket[i].value = value
            return struct{}{}
        }
        if !entry.active && reuse_index < 0 {
            reuse_index = i
        }
        i = i + 1
    }
    if reuse_index >= 0 {
        bucket[reuse_index] = hashmap_DerivedKey_string_x_entry{
            active: true,
            key: key,
            value: value,
        }
        m.len = m.len + 1
        return struct{}{}
    }
    bucket = append(bucket, hashmap_DerivedKey_string_x_entry{
        active: true,
        key: key,
        value: value,
    })
    m.buckets[h] = bucket
    m.len = m.len + 1
    return struct{}{}
}

type hashmap_CustomKey_string_x_entry struct {
    active bool
    key CustomKey
    value string
}

type hashmap_CustomKey_string_x struct {
    buckets map[uint64][]hashmap_CustomKey_string_x_entry
    hashes []uint64
    len int
}

func hashmap_new__HashMap_9CustomKey_6string() *hashmap_CustomKey_string_x {
    return &hashmap_CustomKey_string_x{
        buckets: make(map[uint64][]hashmap_CustomKey_string_x_entry),
        len: 0,
        hashes: nil,
    }
}

func hashmap_lookup__HashMap_9CustomKey_6string(m *hashmap_CustomKey_string_x, key CustomKey) (string, bool, int, uint64) {
    if m == nil {
        var zero string
        return zero, false, -1, 0
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_CustomKey_i_hash(key)
    var bucket []hashmap_CustomKey_string_x_entry = m.buckets[h]
    var i int = 0
    var reuse_index int = -1
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_CustomKey_string_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_CustomKey_i_eq(entry.key, key) {
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

func hashmap_get__HashMap_9CustomKey_6string(m *hashmap_CustomKey_string_x, key CustomKey) Option__string {
    var value string
    var ok bool
    value, ok, _, _ = hashmap_lookup__HashMap_9CustomKey_6string(m, key)
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

func hashmap_set__HashMap_9CustomKey_6string(m *hashmap_CustomKey_string_x, key CustomKey, value string) struct{} {
    var reuse_index int = -1
    if m == nil {
        return struct{}{}
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_CustomKey_i_hash(key)
    var bucket []hashmap_CustomKey_string_x_entry = m.buckets[h]
    if len(bucket) == 0 {
        m.hashes = append(m.hashes, h)
    }
    var i int = 0
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_CustomKey_string_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_CustomKey_i_eq(entry.key, key) {
            bucket[i].value = value
            return struct{}{}
        }
        if !entry.active && reuse_index < 0 {
            reuse_index = i
        }
        i = i + 1
    }
    if reuse_index >= 0 {
        bucket[reuse_index] = hashmap_CustomKey_string_x_entry{
            active: true,
            key: key,
            value: value,
        }
        m.len = m.len + 1
        return struct{}{}
    }
    bucket = append(bucket, hashmap_CustomKey_string_x_entry{
        active: true,
        key: key,
        value: value,
    })
    m.buckets[h] = bucket
    m.len = m.len + 1
    return struct{}{}
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

type DerivedKey struct {
    value int
}

type CustomKey struct {
    value int
}

type Rank struct {
    value int
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

func _goml_m_trait__impl_i_PartialEq_i_DerivedKey_i_eq(self__0 DerivedKey, other__0 DerivedKey) bool {
    var t0 int = self__0.value
    var t1 int = other__0.value
    var inline0 bool = t0 == t1
    return inline0
}

func _goml_m_trait__impl_i_Hash_i_DerivedKey_i_hash(self__0 DerivedKey) uint64 {
    var t0_source int = 0
    var t0 uint64 = uint64(int(t0_source))
    var h__0 uint64 = t0 + 14695981039346656037
    var t1_source int = 0
    var t1 uint64 = uint64(int(t1_source))
    var t2 uint64 = t1 + 1099511628211
    var t3 uint64 = h__0 * t2
    var t4 int = self__0.value
    var t5 uint64
    var inline0 uint64 = _goml_runtime_core_int_hash(t4)
    t5 = inline0
    var h__1 uint64 = t3 + t5
    return h__1
}

func _goml_m_trait__impl_i_PartialEq_i_CustomKey_i_eq(self__0 CustomKey, other__0 CustomKey) bool {
    var t0 int = self__0.value
    var t1 int = t0 % 10
    var t2 int = other__0.value
    var t3 int = t2 % 10
    var t4 bool = t1 == t3
    return t4
}

func _goml_m_trait__impl_i_Hash_i_CustomKey_i_hash(self__0 CustomKey) uint64 {
    var t0 int = self__0.value
    var t1 int = t0 % 10
    var t2 uint64 = uint64(int(t1))
    return t2
}

func main0() struct{} {
    var derived__0 DerivedKey = DerivedKey{
        value: 7,
    }
    var custom__0 CustomKey = CustomKey{
        value: 9,
    }
    var t0 DerivedKey = DerivedKey{
        value: 7,
    }
    var t1 bool = equal__T_DerivedKey(derived__0, t0)
    println__T_bool(t1)
    var t2 DerivedKey = DerivedKey{
        value: 8,
    }
    var t3 bool = _goml_m_trait__impl_i_PartialEq_i_DerivedKey_i_eq(derived__0, t2)
    println__T_bool(t3)
    var t4 CustomKey = CustomKey{
        value: 19,
    }
    var t5 bool = equal__T_CustomKey(custom__0, t4)
    println__T_bool(t5)
    var t6 CustomKey = CustomKey{
        value: 20,
    }
    var t7 bool = _goml_m_trait__impl_i_PartialEq_i_CustomKey_i_eq(custom__0, t6)
    var t8 bool = !t7
    var t9 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t8)
    println__T_string(t9)
    var derived_map__0 *hashmap_DerivedKey_string_x = _goml_m_inherent_i_HashMap_i_H_h11367fd98aaafe315368863eb9d8e24f_ey____V__string()
    _goml_m_inherent_i_HashMap_i_H_h1cee44072715470f45f5985ef34cf1f9_ey____V__string(derived_map__0, derived__0, "derived")
    var t10 DerivedKey = DerivedKey{
        value: 7,
    }
    var t11 Option__string
    var inline26 Option__string = hashmap_get__HashMap_10DerivedKey_6string(derived_map__0, t10)
    t11 = inline26
    var t12 string = _goml_m_inherent_i_Option_i_Option_l_T_r__i_unwrap__or____T__string(t11, "missing")
    var inline24 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t12)
    _goml_runtime_core_string_println(inline24)
    var custom_map__0 *hashmap_CustomKey_string_x
    var inline23 *hashmap_CustomKey_string_x = hashmap_new__HashMap_9CustomKey_6string()
    custom_map__0 = inline23
    var inline21 string = "custom"
    hashmap_set__HashMap_9CustomKey_6string(custom_map__0, custom__0, inline21)
    var t13 CustomKey = CustomKey{
        value: 19,
    }
    var t14 Option__string
    var inline20 Option__string = hashmap_get__HashMap_9CustomKey_6string(custom_map__0, t13)
    t14 = inline20
    var t15 string
    var inline18 string = "missing"
    switch t14._tag {
    case 0:
        t15 = inline18
    case 1:
        var inline19 string = t14._p0
        t15 = inline19
    default:
        panic("non-exhaustive match")
    }
    var inline16 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t15)
    _goml_runtime_core_string_println(inline16)
    var high__0 Rank = Rank{
        value: 2,
    }
    var low__0 Rank = Rank{
        value: 1,
    }
    var t16 bool
    var inline15 bool = _goml_m_std_p_cmp_p_trait__default_i_PartialOrd_i_lt____Self__Rank(high__0, low__0)
    t16 = inline15
    var t17 string
    var inline14 string = _goml_runtime_core_bool_to_string(t16)
    t17 = inline14
    var inline12 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t17)
    _goml_runtime_core_string_println(inline12)
    var t18 bool
    var inline11 bool = _goml_m_std_p_cmp_p_trait__default_i_PartialOrd_i_le____Self__Rank(high__0, high__0)
    t18 = inline11
    var t19 string
    var inline10 string = _goml_runtime_core_bool_to_string(t18)
    t19 = inline10
    var inline8 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t19)
    _goml_runtime_core_string_println(inline8)
    var t20 bool
    var inline7 bool = _goml_m_std_p_cmp_p_trait__default_i_PartialOrd_i_gt____Self__Rank(low__0, high__0)
    t20 = inline7
    var t21 string
    var inline6 string = _goml_runtime_core_bool_to_string(t20)
    t21 = inline6
    var inline4 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t21)
    _goml_runtime_core_string_println(inline4)
    var t22 bool
    var inline3 bool = _goml_m_std_p_cmp_p_trait__default_i_PartialOrd_i_ge____Self__Rank(low__0, low__0)
    t22 = inline3
    var t23 string
    var inline2 string = _goml_runtime_core_bool_to_string(t22)
    t23 = inline2
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t23)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_trait__impl_i_PartialEq_i_isize_i_eq(self__0 int, other__0 int) bool {
    var t0 bool = self__0 == other__0
    return t0
}

func _goml_m_std_p_cmp_p_trait__default_i_PartialOrd_i_lt____Self__Rank(self__0 Rank, other__0 Rank) bool {
    var commute_field0 Ordering
    var inline0 int = self__0.value
    var inline1 int = other__0.value
    var inline2 bool = inline0 > inline1
    var inline3 Ordering
    if inline2 {
        inline3 = Less
    } else {
        var inline4 int = self__0.value
        var inline5 int = other__0.value
        var inline6 bool = inline4 < inline5
        if inline6 {
            inline3 = Greater
        } else {
            inline3 = Equal
        }
    }
    commute_field0 = inline3
    switch commute_field0 {
    case Less:
        return true
    case Equal:
        return false
    case Greater:
        return false
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_std_p_cmp_p_trait__default_i_PartialOrd_i_le____Self__Rank(self__0 Rank, other__0 Rank) bool {
    var commute_field0 Ordering
    var inline0 int = self__0.value
    var inline1 int = other__0.value
    var inline2 bool = inline0 > inline1
    var inline3 Ordering
    if inline2 {
        inline3 = Less
    } else {
        var inline4 int = self__0.value
        var inline5 int = other__0.value
        var inline6 bool = inline4 < inline5
        if inline6 {
            inline3 = Greater
        } else {
            inline3 = Equal
        }
    }
    commute_field0 = inline3
    var t0 bool
    switch commute_field0 {
    case Less:
        t0 = false
    case Equal:
        t0 = false
    case Greater:
        t0 = true
    default:
        panic("non-exhaustive match")
    }
    var t1 bool = !t0
    return t1
}

func _goml_m_std_p_cmp_p_trait__default_i_PartialOrd_i_gt____Self__Rank(self__0 Rank, other__0 Rank) bool {
    var commute_field0 Ordering
    var inline0 int = self__0.value
    var inline1 int = other__0.value
    var inline2 bool = inline0 > inline1
    var inline3 Ordering
    if inline2 {
        inline3 = Less
    } else {
        var inline4 int = self__0.value
        var inline5 int = other__0.value
        var inline6 bool = inline4 < inline5
        if inline6 {
            inline3 = Greater
        } else {
            inline3 = Equal
        }
    }
    commute_field0 = inline3
    switch commute_field0 {
    case Less:
        return false
    case Equal:
        return false
    case Greater:
        return true
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_std_p_cmp_p_trait__default_i_PartialOrd_i_ge____Self__Rank(self__0 Rank, other__0 Rank) bool {
    var commute_field0 Ordering
    var inline0 int = self__0.value
    var inline1 int = other__0.value
    var inline2 bool = inline0 > inline1
    var inline3 Ordering
    if inline2 {
        inline3 = Less
    } else {
        var inline4 int = self__0.value
        var inline5 int = other__0.value
        var inline6 bool = inline4 < inline5
        if inline6 {
            inline3 = Greater
        } else {
            inline3 = Equal
        }
    }
    commute_field0 = inline3
    var t0 bool
    switch commute_field0 {
    case Less:
        t0 = true
    case Equal:
        t0 = false
    case Greater:
        t0 = false
    default:
        panic("non-exhaustive match")
    }
    var t1 bool = !t0
    return t1
}

func println__T_bool(value__0 bool) struct{} {
    var t0 string
    var inline0 string = _goml_runtime_core_bool_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func equal__T_DerivedKey(left__0 DerivedKey, right__0 DerivedKey) bool {
    var inline0 int = left__0.value
    var inline1 int = right__0.value
    var inline2 bool = _goml_m_trait__impl_i_PartialEq_i_isize_i_eq(inline0, inline1)
    return inline2
}

func equal__T_CustomKey(left__0 CustomKey, right__0 CustomKey) bool {
    var inline0 int = left__0.value
    var inline1 int = inline0 % 10
    var inline2 int = right__0.value
    var inline3 int = inline2 % 10
    var inline4 bool = inline1 == inline3
    return inline4
}

func println__T_string(value__0 string) struct{} {
    var t0 string
    t0 = value__0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_bool_i_to__string(self__0 bool) string {
    var t0 string = _goml_runtime_core_bool_to_string(self__0)
    return t0
}

func _goml_m_inherent_i_HashMap_i_H_h11367fd98aaafe315368863eb9d8e24f_ey____V__string() *hashmap_DerivedKey_string_x {
    var t0 *hashmap_DerivedKey_string_x = hashmap_new__HashMap_10DerivedKey_6string()
    return t0
}

func _goml_m_inherent_i_HashMap_i_H_h1cee44072715470f45f5985ef34cf1f9_ey____V__string(self__0 *hashmap_DerivedKey_string_x, key__0 DerivedKey, value__0 string) struct{} {
    hashmap_set__HashMap_10DerivedKey_6string(self__0, key__0, value__0)
    return struct{}{}
}

func _goml_m_inherent_i_Option_i_Option_l_T_r__i_unwrap__or____T__string(self__0 Option__string, fallback__0 string) string {
    switch self__0._tag {
    case 0:
        return fallback__0
    case 1:
        var x0 string = self__0._p0
        return x0
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func main() {
    main0()
}
