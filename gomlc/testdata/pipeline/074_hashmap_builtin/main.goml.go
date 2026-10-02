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

func _goml_runtime_core_int32_hash(x int32) uint64 {
    return uint64(x)
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

type ref_Point_x struct {
    value Point
}

func ref__Ref_5Point(value Point) *ref_Point_x {
    return &ref_Point_x{
        value: value,
    }
}

func ref_set__Ref_5Point(reference *ref_Point_x, value Point) struct{} {
    reference.value = value
    return struct{}{}
}

func ptr_eq__Ref_5Point(a *ref_Point_x, b *ref_Point_x) bool {
    return a == b
}

func ptr_hash__Ref_5Point(reference *ref_Point_x) uint64 {
    return uint64(_goml_reflect.ValueOf(reference).Pointer())
}

type ref_Key_x struct {
    value Key
}

func ref__Ref_3Key(value Key) *ref_Key_x {
    return &ref_Key_x{
        value: value,
    }
}

func ptr_eq__Ref_3Key(a *ref_Key_x, b *ref_Key_x) bool {
    return a == b
}

func ptr_hash__Ref_3Key(reference *ref_Key_x) uint64 {
    return uint64(_goml_reflect.ValueOf(reference).Pointer())
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

type hashmap_Ref_5Point_int32_x_entry struct {
    active bool
    key *ref_Point_x
    value int32
}

type hashmap_Ref_5Point_int32_x struct {
    buckets map[uint64][]hashmap_Ref_5Point_int32_x_entry
    hashes []uint64
    len int
}

func hashmap_new__HashMap_10Ref_5Point_5int32() *hashmap_Ref_5Point_int32_x {
    return &hashmap_Ref_5Point_int32_x{
        buckets: make(map[uint64][]hashmap_Ref_5Point_int32_x_entry),
        len: 0,
        hashes: nil,
    }
}

func hashmap_lookup__HashMap_10Ref_5Point_5int32(m *hashmap_Ref_5Point_int32_x, key *ref_Point_x) (int32, bool, int, uint64) {
    if m == nil {
        var zero int32
        return zero, false, -1, 0
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_Ref_l_Point_r__i_hash(key)
    var bucket []hashmap_Ref_5Point_int32_x_entry = m.buckets[h]
    var i int = 0
    var reuse_index int = -1
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_Ref_5Point_int32_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_Ref_l_Point_r__i_eq(entry.key, key) {
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

func hashmap_get__HashMap_10Ref_5Point_5int32(m *hashmap_Ref_5Point_int32_x, key *ref_Point_x) Option__i32 {
    var value int32
    var ok bool
    value, ok, _, _ = hashmap_lookup__HashMap_10Ref_5Point_5int32(m, key)
    if ok {
        return Option__i32(uint64(int64(value) + 2147483648) + 1)
    }
    return Option__i32(0)
}

func hashmap_set__HashMap_10Ref_5Point_5int32(m *hashmap_Ref_5Point_int32_x, key *ref_Point_x, value int32) struct{} {
    var reuse_index int = -1
    if m == nil {
        return struct{}{}
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_Ref_l_Point_r__i_hash(key)
    var bucket []hashmap_Ref_5Point_int32_x_entry = m.buckets[h]
    if len(bucket) == 0 {
        m.hashes = append(m.hashes, h)
    }
    var i int = 0
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_Ref_5Point_int32_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_Ref_l_Point_r__i_eq(entry.key, key) {
            bucket[i].value = value
            return struct{}{}
        }
        if !entry.active && reuse_index < 0 {
            reuse_index = i
        }
        i = i + 1
    }
    if reuse_index >= 0 {
        bucket[reuse_index] = hashmap_Ref_5Point_int32_x_entry{
            active: true,
            key: key,
            value: value,
        }
        m.len = m.len + 1
        return struct{}{}
    }
    bucket = append(bucket, hashmap_Ref_5Point_int32_x_entry{
        active: true,
        key: key,
        value: value,
    })
    m.buckets[h] = bucket
    m.len = m.len + 1
    return struct{}{}
}

type hashmap_Ref_3Key_int32_x_entry struct {
    active bool
    key *ref_Key_x
    value int32
}

type hashmap_Ref_3Key_int32_x struct {
    buckets map[uint64][]hashmap_Ref_3Key_int32_x_entry
    hashes []uint64
    len int
}

func hashmap_new__HashMap_8Ref_3Key_5int32() *hashmap_Ref_3Key_int32_x {
    return &hashmap_Ref_3Key_int32_x{
        buckets: make(map[uint64][]hashmap_Ref_3Key_int32_x_entry),
        len: 0,
        hashes: nil,
    }
}

func hashmap_lookup__HashMap_8Ref_3Key_5int32(m *hashmap_Ref_3Key_int32_x, key *ref_Key_x) (int32, bool, int, uint64) {
    if m == nil {
        var zero int32
        return zero, false, -1, 0
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_Ref_l_Key_r__i_hash(key)
    var bucket []hashmap_Ref_3Key_int32_x_entry = m.buckets[h]
    var i int = 0
    var reuse_index int = -1
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_Ref_3Key_int32_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_Ref_l_Key_r__i_eq(entry.key, key) {
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

func hashmap_get__HashMap_8Ref_3Key_5int32(m *hashmap_Ref_3Key_int32_x, key *ref_Key_x) Option__i32 {
    var value int32
    var ok bool
    value, ok, _, _ = hashmap_lookup__HashMap_8Ref_3Key_5int32(m, key)
    if ok {
        return Option__i32(uint64(int64(value) + 2147483648) + 1)
    }
    return Option__i32(0)
}

func hashmap_set__HashMap_8Ref_3Key_5int32(m *hashmap_Ref_3Key_int32_x, key *ref_Key_x, value int32) struct{} {
    var reuse_index int = -1
    if m == nil {
        return struct{}{}
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_Ref_l_Key_r__i_hash(key)
    var bucket []hashmap_Ref_3Key_int32_x_entry = m.buckets[h]
    if len(bucket) == 0 {
        m.hashes = append(m.hashes, h)
    }
    var i int = 0
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_Ref_3Key_int32_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_Ref_l_Key_r__i_eq(entry.key, key) {
            bucket[i].value = value
            return struct{}{}
        }
        if !entry.active && reuse_index < 0 {
            reuse_index = i
        }
        i = i + 1
    }
    if reuse_index >= 0 {
        bucket[reuse_index] = hashmap_Ref_3Key_int32_x_entry{
            active: true,
            key: key,
            value: value,
        }
        m.len = m.len + 1
        return struct{}{}
    }
    bucket = append(bucket, hashmap_Ref_3Key_int32_x_entry{
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

type Point struct {
    x int32
    y int32
}

type Ordering uint8

type Key struct {
    _p1 Point
    _p0 int32
    _tag uint8
}

type Option__i32 uint64

func _goml_m_trait__impl_i_Hash_i_Key_i_hash(self__0 Key) uint64 {
    switch self__0._tag {
    case 0:
        var t0_source int = 0
        var t0 uint64 = uint64(int(t0_source))
        var t1 uint64 = t0 + 14695981039346656037
        var h__0 uint64 = t1 + 1
        return h__0
    case 1:
        var x0 int32 = self__0._p0
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
    case 2:
        var x1 Point = self__0._p1
        var t8_source int = 0
        var t8 uint64 = uint64(int(t8_source))
        var t9 uint64 = t8 + 14695981039346656037
        var h__3 uint64 = t9 + 3
        var t10_source int = 0
        var t10 uint64 = uint64(int(t10_source))
        var t11 uint64 = t10 + 1099511628211
        var t12 uint64 = h__3 * t11
        var t13 uint64
        var inline1_source int = 0
        var inline1 uint64 = uint64(int(inline1_source))
        var inline2 uint64 = inline1 + 14695981039346656037
        var inline3_source int = 0
        var inline3 uint64 = uint64(int(inline3_source))
        var inline4 uint64 = inline3 + 1099511628211
        var inline5 uint64 = inline2 * inline4
        var inline6 int32 = x1.x
        var inline7 uint64 = _goml_m_trait__impl_i_Hash_i_i32_i_hash(inline6)
        var inline8 uint64 = inline5 + inline7
        var inline9_source int = 0
        var inline9 uint64 = uint64(int(inline9_source))
        var inline10 uint64 = inline9 + 1099511628211
        var inline11 uint64 = inline8 * inline10
        var inline12 int32 = x1.y
        var inline13 uint64 = _goml_m_trait__impl_i_Hash_i_i32_i_hash(inline12)
        var inline14 uint64 = inline11 + inline13
        t13 = inline14
        var h__4 uint64 = t12 + t13
        return h__4
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_trait__impl_i_PartialEq_i_Key_i_eq(self__0 Key, other__0 Key) bool {
    switch other__0._tag {
    case 0:
        switch self__0._tag {
        case 0:
            return true
        default:
            return false
        }
    case 1:
        var x0 int32 = other__0._p0
        switch self__0._tag {
        case 1:
            var x1 int32 = self__0._p0
            var inline0 bool = x1 == x0
            return inline0
        default:
            return false
        }
    case 2:
        var x2 Point = other__0._p1
        switch self__0._tag {
        case 2:
            var x3 Point = self__0._p1
            var inline1 bool
            var inline5 int32 = x3.x
            var inline6 int32 = x2.x
            var inline7 bool = _goml_m_trait__impl_i_PartialEq_i_i32_i_eq(inline5, inline6)
            inline1 = inline7
            if inline1 {
                var inline2 int32 = x3.y
                var inline3 int32 = x2.y
                var inline4 bool = _goml_m_trait__impl_i_PartialEq_i_i32_i_eq(inline2, inline3)
                return inline4
            } else {
                return false
            }
        default:
            return false
        }
    default:
        panic("non-exhaustive match")
    }
}

func print_opt_int(x__0 Option__i32) struct{} {
    switch x__0 != Option__i32(0) {
    case false:
        var inline0 string = "none"
        var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
        _goml_runtime_core_string_println(inline1)
        return struct{}{}
    case true:
        var x0 int32 = int32(int64(uint64(x__0) - 1) - 2147483648)
        var inline3 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(x0)
        _goml_runtime_core_string_println(inline3)
        return struct{}{}
    default:
        panic("non-exhaustive match")
    }
}

func main0() struct{} {
    var m1__0 *hashmap_Key_int32_x = _goml_m_inherent_i_HashMap_i_HashMap_l_K_c_V_r__i_new____K__Key____V__i32()
    _goml_m_inherent_i_HashMap_i_HashMap_l_K_c_V_r__i_set____K__Key____V__i32(m1__0, Key{
        _tag: 0,
    }, 10)
    var t0 Key = Key{
        _p0: 1,
        _tag: 1,
    }
    _goml_m_inherent_i_HashMap_i_HashMap_l_K_c_V_r__i_set____K__Key____V__i32(m1__0, t0, 20)
    var t1 int = _goml_m_inherent_i_HashMap_i_HashMap_l_K_c_V_r__i_len____K__Key____V__i32(m1__0)
    println__T_isize(t1)
    var t2 Option__i32 = _goml_m_inherent_i_HashMap_i_HashMap_l_K_c_V_r__i_get____K__Key____V__i32(m1__0, Key{
        _tag: 0,
    })
    print_opt_int(t2)
    var t3 Key = Key{
        _p0: 1,
        _tag: 1,
    }
    var t4 bool = _goml_m_inherent_i_HashMap_i_HashMap_l_K_c_V_r__i_contains____K__Key____V__i32(m1__0, t3)
    println__T_bool(t4)
    var t5 Key = Key{
        _p0: 1,
        _tag: 1,
    }
    _goml_m_inherent_i_HashMap_i_HashMap_l_K_c_V_r__i_remove____K__Key____V__i32(m1__0, t5)
    var t6 Key = Key{
        _p0: 1,
        _tag: 1,
    }
    var t7 bool = _goml_m_inherent_i_HashMap_i_HashMap_l_K_c_V_r__i_contains____K__Key____V__i32(m1__0, t6)
    println__T_bool(t7)
    var t8 int = _goml_m_inherent_i_HashMap_i_HashMap_l_K_c_V_r__i_len____K__Key____V__i32(m1__0)
    println__T_isize(t8)
    var m2__0 *hashmap_Ref_5Point_int32_x = _goml_m_inherent_i_HashMap_i_H_h32f703bdb7b587c4ce71a39164c77ac4_nt_r_____V__i32()
    var t9 Point = Point{
        x: 1,
        y: 2,
    }
    var p1__0 *ref_Point_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__Point(t9)
    var t10 Point = Point{
        x: 1,
        y: 2,
    }
    var p2__0 *ref_Point_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__Point(t10)
    _goml_m_inherent_i_HashMap_i_H_h92e4b8ba44d2c4c111acc209f0277885_nt_r_____V__i32(m2__0, p1__0, 99)
    var t11 Option__i32 = _goml_m_inherent_i_HashMap_i_H_hf74f21a0b6b06890f12f275bddb17d70_nt_r_____V__i32(m2__0, p1__0)
    switch t11 != Option__i32(0) {
    case false:
        println__T_string("none")
    case true:
        var inline21 int32 = int32(int64(uint64(t11) - 1) - 2147483648)
        println__T_i32(inline21)
    default:
        panic("non-exhaustive match")
    }
    var t12 Option__i32
    var inline19 Option__i32 = hashmap_get__HashMap_10Ref_5Point_5int32(m2__0, p2__0)
    t12 = inline19
    switch t12 != Option__i32(0) {
    case false:
        println__T_string("none")
    case true:
        var inline17 int32 = int32(int64(uint64(t12) - 1) - 2147483648)
        println__T_i32(inline17)
    default:
        panic("non-exhaustive match")
    }
    var t13 Point = Point{
        x: 9,
        y: 8,
    }
    ref_set__Ref_5Point(p1__0, t13)
    var t14 Option__i32
    var inline14 Option__i32 = hashmap_get__HashMap_10Ref_5Point_5int32(m2__0, p1__0)
    t14 = inline14
    switch t14 != Option__i32(0) {
    case false:
        println__T_string("none")
    case true:
        var inline12 int32 = int32(int64(uint64(t14) - 1) - 2147483648)
        println__T_i32(inline12)
    default:
        panic("non-exhaustive match")
    }
    var t15 bool
    var inline10 bool = ptr_eq__Ref_5Point(p1__0, p2__0)
    t15 = inline10
    var inline8 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t15)
    _goml_runtime_core_string_println(inline8)
    var m3__0 *hashmap_Ref_3Key_int32_x
    var inline7 *hashmap_Ref_3Key_int32_x = hashmap_new__HashMap_8Ref_3Key_5int32()
    m3__0 = inline7
    var t16 Key = Key{
        _p0: 7,
        _tag: 1,
    }
    var k1__0 *ref_Key_x
    var inline6 *ref_Key_x = ref__Ref_3Key(t16)
    k1__0 = inline6
    var inline4 int32 = 123
    hashmap_set__HashMap_8Ref_3Key_5int32(m3__0, k1__0, inline4)
    var t17 Option__i32
    var inline3 Option__i32 = hashmap_get__HashMap_8Ref_3Key_5int32(m3__0, k1__0)
    t17 = inline3
    switch t17 != Option__i32(0) {
    case false:
        println__T_string("none")
        return struct{}{}
    case true:
        var inline1 int32 = int32(int64(uint64(t17) - 1) - 2147483648)
        println__T_i32(inline1)
        return struct{}{}
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_trait__impl_i_Hash_i_i32_i_hash(self__0 int32) uint64 {
    var t0 uint64 = _goml_runtime_core_int32_hash(self__0)
    return t0
}

func _goml_m_trait__impl_i_PartialEq_i_i32_i_eq(self__0 int32, other__0 int32) bool {
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

func _goml_m_inherent_i_HashMap_i_HashMap_l_K_c_V_r__i_new____K__Key____V__i32() *hashmap_Key_int32_x {
    var t0 *hashmap_Key_int32_x = hashmap_new__HashMap_3Key_5int32()
    return t0
}

func _goml_m_inherent_i_HashMap_i_HashMap_l_K_c_V_r__i_set____K__Key____V__i32(self__0 *hashmap_Key_int32_x, key__0 Key, value__0 int32) struct{} {
    hashmap_set__HashMap_3Key_5int32(self__0, key__0, value__0)
    return struct{}{}
}

func println__T_isize(value__0 int) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_int_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_inherent_i_HashMap_i_HashMap_l_K_c_V_r__i_len____K__Key____V__i32(self__0 *hashmap_Key_int32_x) int {
    var t0 int = hashmap_len__HashMap_3Key_5int32(self__0)
    return t0
}

func _goml_m_inherent_i_HashMap_i_HashMap_l_K_c_V_r__i_get____K__Key____V__i32(self__0 *hashmap_Key_int32_x, key__0 Key) Option__i32 {
    var t0 Option__i32 = hashmap_get__HashMap_3Key_5int32(self__0, key__0)
    return t0
}

func println__T_bool(value__0 bool) struct{} {
    var t0 string
    var inline0 string = _goml_runtime_core_bool_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_inherent_i_HashMap_i_HashMap_l_K_c_V_r__i_contains____K__Key____V__i32(self__0 *hashmap_Key_int32_x, key__0 Key) bool {
    var t0 bool = hashmap_contains__HashMap_3Key_5int32(self__0, key__0)
    return t0
}

func _goml_m_inherent_i_HashMap_i_HashMap_l_K_c_V_r__i_remove____K__Key____V__i32(self__0 *hashmap_Key_int32_x, key__0 Key) struct{} {
    hashmap_remove__HashMap_3Key_5int32(self__0, key__0)
    return struct{}{}
}

func _goml_m_inherent_i_HashMap_i_H_h32f703bdb7b587c4ce71a39164c77ac4_nt_r_____V__i32() *hashmap_Ref_5Point_int32_x {
    var t0 *hashmap_Ref_5Point_int32_x = hashmap_new__HashMap_10Ref_5Point_5int32()
    return t0
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__Point(value__0 Point) *ref_Point_x {
    var t0 *ref_Point_x = ref__Ref_5Point(value__0)
    return t0
}

func _goml_m_inherent_i_HashMap_i_H_h92e4b8ba44d2c4c111acc209f0277885_nt_r_____V__i32(self__0 *hashmap_Ref_5Point_int32_x, key__0 *ref_Point_x, value__0 int32) struct{} {
    hashmap_set__HashMap_10Ref_5Point_5int32(self__0, key__0, value__0)
    return struct{}{}
}

func _goml_m_inherent_i_HashMap_i_H_hf74f21a0b6b06890f12f275bddb17d70_nt_r_____V__i32(self__0 *hashmap_Ref_5Point_int32_x, key__0 *ref_Point_x) Option__i32 {
    var t0 Option__i32 = hashmap_get__HashMap_10Ref_5Point_5int32(self__0, key__0)
    return t0
}

func _goml_m_trait__impl_i_PartialEq_i_Ref_l_Point_r__i_eq(self__0 *ref_Point_x, other__0 *ref_Point_x) bool {
    var t0 bool = ptr_eq__Ref_5Point(self__0, other__0)
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

func _goml_m_trait__impl_i_Hash_i_Ref_l_Point_r__i_hash(self__0 *ref_Point_x) uint64 {
    var t0 uint64 = ptr_hash__Ref_5Point(self__0)
    return t0
}

func _goml_m_trait__impl_i_PartialEq_i_Ref_l_Key_r__i_eq(self__0 *ref_Key_x, other__0 *ref_Key_x) bool {
    var t0 bool = ptr_eq__Ref_3Key(self__0, other__0)
    return t0
}

func _goml_m_trait__impl_i_Hash_i_Ref_l_Key_r__i_hash(self__0 *ref_Key_x) uint64 {
    var t0 uint64 = ptr_hash__Ref_3Key(self__0)
    return t0
}

func main() {
    main0()
}
