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

type ref_int_x struct {
    value int
}

func ref__Ref_3int(value int) *ref_int_x {
    return &ref_int_x{
        value: value,
    }
}

func ref_get__Ref_3int(reference *ref_int_x) int {
    return reference.value
}

func ref_set__Ref_3int(reference *ref_int_x, value int) struct{} {
    reference.value = value
    return struct{}{}
}

type hashmap_Key_int_x_entry struct {
    active bool
    key Key
    value int
}

type hashmap_Key_int_x struct {
    buckets map[uint64][]hashmap_Key_int_x_entry
    hashes []uint64
    len int
}

func hashmap_new__HashMap_3Key_3int() *hashmap_Key_int_x {
    return &hashmap_Key_int_x{
        buckets: make(map[uint64][]hashmap_Key_int_x_entry),
        len: 0,
        hashes: nil,
    }
}

func hashmap_len__HashMap_3Key_3int(m *hashmap_Key_int_x) int {
    if m == nil {
        return 0
    }
    return m.len
}

func hashmap_lookup__HashMap_3Key_3int(m *hashmap_Key_int_x, key Key) (int, bool, int, uint64) {
    if m == nil {
        var zero int
        return zero, false, -1, 0
    }
    var h uint64 = _goml_m_trait__impl_i_Hash_i_Key_i_hash(key)
    var bucket []hashmap_Key_int_x_entry = m.buckets[h]
    var i int = 0
    var reuse_index int = -1
    for {
        if i >= int(len(bucket)) {
            break
        }
        var entry hashmap_Key_int_x_entry = bucket[i]
        if entry.active && _goml_m_trait__impl_i_PartialEq_i_Key_i_eq(entry.key, key) {
            return entry.value, true, i, h
        }
        if !entry.active && reuse_index < 0 {
            reuse_index = i
        }
        i = i + 1
    }
    var zero int
    return zero, false, reuse_index, h
}

func hashmap_get_or_insert_with__HashMap_3Key_3int(m *hashmap_Key_int_x, key Key, create func() int) int {
    var previous int
    var found bool
    var index int
    var hash uint64
    previous, found, index, hash = hashmap_lookup__HashMap_3Key_3int(m, key)
    _ = previous
    _ = index
    _ = hash
    if found {
        return previous
    }
    var value int = create()
    if m == nil {
        return value
    }
    var bucket []hashmap_Key_int_x_entry = m.buckets[hash]
    if index >= 0 {
        bucket[index] = hashmap_Key_int_x_entry{
            active: true,
            key: key,
            value: value,
        }
    } else {
        if len(bucket) == 0 {
            m.hashes = append(m.hashes, hash)
        }
        bucket = append(bucket, hashmap_Key_int_x_entry{
            active: true,
            key: key,
            value: value,
        })
    }
    m.buckets[hash] = bucket
    m.len = m.len + 1
    return value
}

func hashmap_update__HashMap_3Key_3int(m *hashmap_Key_int_x, key Key, update func(int) int) Option__isize {
    var previous int
    var found bool
    var index int
    var hash uint64
    previous, found, index, hash = hashmap_lookup__HashMap_3Key_3int(m, key)
    _ = previous
    _ = index
    _ = hash
    if !found {
        return Option__isize{
            _tag: 0,
        }
    }
    var value int = update(previous)
    m.buckets[hash][index].value = value
    return Option__isize{
        _tag: 1,
        _p0: value,
    }
}

func hashmap_remove_value__HashMap_3Key_3int(m *hashmap_Key_int_x, key Key) Option__isize {
    var previous int
    var found bool
    var index int
    var hash uint64
    previous, found, index, hash = hashmap_lookup__HashMap_3Key_3int(m, key)
    _ = previous
    _ = index
    _ = hash
    if !found {
        return Option__isize{
            _tag: 0,
        }
    }
    var zero hashmap_Key_int_x_entry
    m.buckets[hash][index] = zero
    m.len = m.len - 1
    return Option__isize{
        _tag: 1,
        _p0: previous,
    }
}

type hashmap_string_int_x_entry struct {
    active bool
    key string
    value int
}

type hashmap_string_int_x struct {
    indices map[string]int
    entries []hashmap_string_int_x_entry
    len int
}

func hashmap_new__HashMap_6string_3int() *hashmap_string_int_x {
    return &hashmap_string_int_x{
        indices: make(map[string]int),
        entries: nil,
        len: 0,
    }
}

func hashmap_lookup__HashMap_6string_3int(m *hashmap_string_int_x, key string) (int, bool, int, uint64) {
    if m == nil {
        var zero int
        return zero, false, -1, 0
    }
    var index int
    var found bool
    index, found = m.indices[key]
    if !found {
        var zero int
        return zero, false, -1, 0
    }
    var entry hashmap_string_int_x_entry = m.entries[index]
    if entry.active {
        return entry.value, true, index, 0
    }
    var zero int
    return zero, false, index, 0
}

func hashmap_get__HashMap_6string_3int(m *hashmap_string_int_x, key string) Option__isize {
    var value int
    var ok bool
    value, ok, _, _ = hashmap_lookup__HashMap_6string_3int(m, key)
    if ok {
        return Option__isize{
            _tag: 1,
            _p0: value,
        }
    }
    return Option__isize{
        _tag: 0,
    }
}

func hashmap_insert__HashMap_6string_3int(m *hashmap_string_int_x, key string, value int) Option__isize {
    var previous int
    var found bool
    var index int
    var hash uint64
    previous, found, index, hash = hashmap_lookup__HashMap_6string_3int(m, key)
    _ = previous
    _ = index
    _ = hash
    if found {
        m.entries[index].value = value
        return Option__isize{
            _tag: 1,
            _p0: previous,
        }
    }
    if m == nil {
        return Option__isize{
            _tag: 0,
        }
    }
    if index >= 0 {
        m.entries[index] = hashmap_string_int_x_entry{
            active: true,
            key: key,
            value: value,
        }
    } else {
        index = len(m.entries)
        m.indices[key] = index
        m.entries = append(m.entries, hashmap_string_int_x_entry{
            active: true,
            key: key,
            value: value,
        })
    }
    m.len = m.len + 1
    return Option__isize{
        _tag: 0,
    }
}

func hashmap_update__HashMap_6string_3int(m *hashmap_string_int_x, key string, update func(int) int) Option__isize {
    var previous int
    var found bool
    var index int
    var hash uint64
    previous, found, index, hash = hashmap_lookup__HashMap_6string_3int(m, key)
    _ = previous
    _ = index
    _ = hash
    if !found {
        return Option__isize{
            _tag: 0,
        }
    }
    var value int = update(previous)
    m.entries[index].value = value
    return Option__isize{
        _tag: 1,
        _p0: value,
    }
}

func hashmap_remove_value__HashMap_6string_3int(m *hashmap_string_int_x, key string) Option__isize {
    var previous int
    var found bool
    var index int
    var hash uint64
    previous, found, index, hash = hashmap_lookup__HashMap_6string_3int(m, key)
    _ = previous
    _ = index
    _ = hash
    if !found {
        return Option__isize{
            _tag: 0,
        }
    }
    var zero hashmap_string_int_x_entry
    m.entries[index] = zero
    m.len = m.len - 1
    return Option__isize{
        _tag: 1,
        _p0: previous,
    }
}

type Tuple2_3Key_3int struct {
    _0 Key
    _1 int
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

type Key struct {
    value int
    hashes *ref_int_x
}

type HashMapEntry__Key__isize struct {
    values *hashmap_Key_int_x
    key Key
}

type HashMapEntry__string__isize struct {
    values *hashmap_string_int_x
    key string
}

type closure_env_main_0 struct {
    created_0 *ref_int_x
}

type closure_env_main_1 struct {
    created_0 *ref_int_x
}

type closure_env_main_2 struct {}

type closure_env_main_3 struct {}

type closure_env_inherent_HashMapEntry_HashMapEntry_K_V_remove_entry_K_Key_V_isize_4 struct {
    self_0 HashMapEntry__Key__isize
}

type Ordering uint8

type Option__isize struct {
    _p0 int
    _tag uint8
}

type _goml_m_Option_____o_Key_c_isize_q_ struct {
    _p0 Tuple2_3Key_3int
    _tag uint8
}

func _goml_m_trait__impl_i_PartialEq_i_Key_i_eq(self__0 Key, other__0 Key) bool {
    var t0 int = self__0.value
    var t1 int = other__0.value
    var t2 bool = t0 == t1
    return t2
}

func _goml_m_trait__impl_i_Hash_i_Key_i_hash(self__0 Key) uint64 {
    var t0 *ref_int_x = self__0.hashes
    var t1 *ref_int_x = self__0.hashes
    var t2 int
    var inline1 int = ref_get__Ref_3int(t1)
    t2 = inline1
    var t3 int = t2 + 1
    ref_set__Ref_3int(t0, t3)
    return 7
}

func key(value__0 int, hashes__0 *ref_int_x) Key {
    var t0 Key = Key{
        value: value__0,
        hashes: hashes__0,
    }
    return t0
}

func main0() struct{} {
    var hashes__0 *ref_int_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__isize(0)
    var created__0 *ref_int_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__isize(0)
    var map__0 *hashmap_Key_int_x = _goml_m_inherent_i_HashMap_i_HashMap_l_K_c_V_r__i_new____K__Key____V__isize()
    var t0 Key = key(1, hashes__0)
    var t1 HashMapEntry__Key__isize = _goml_m_inherent_i_HashMap_i_HashMap_l_K_c_V_r__i_entry____K__Key____V__isize(map__0, t0)
    var t2 closure_env_main_0 = closure_env_main_0{
        created_0: created__0,
    }
    var t3 func() int = func() int {
        return _goml_m_inherent_i_closure__env__main__0_i_closure__env__main__0_i_apply(t2)
    }
    var t4 int = _goml_m_inherent_i_HashMapEntr_h08451cacc3f0db11c02874f35c8558b6_Key____V__isize(t1, t3)
    println__T_isize(t4)
    var t5 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(hashes__0)
    println__T_isize(t5)
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__isize(hashes__0, 0)
    var t6 Key = key(1, hashes__0)
    var t7 HashMapEntry__Key__isize = _goml_m_inherent_i_HashMap_i_HashMap_l_K_c_V_r__i_entry____K__Key____V__isize(map__0, t6)
    var t8 closure_env_main_1 = closure_env_main_1{
        created_0: created__0,
    }
    var t9 func() int = func() int {
        return _goml_m_inherent_i_closure__env__main__1_i_closure__env__main__1_i_apply(t8)
    }
    var t10 int = _goml_m_inherent_i_HashMapEntr_h08451cacc3f0db11c02874f35c8558b6_Key____V__isize(t7, t9)
    println__T_isize(t10)
    var t11 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(hashes__0)
    println__T_isize(t11)
    var t12 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(created__0)
    println__T_isize(t12)
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__isize(hashes__0, 0)
    var t13 Key = key(1, hashes__0)
    var t14 HashMapEntry__Key__isize = _goml_m_inherent_i_HashMap_i_HashMap_l_K_c_V_r__i_entry____K__Key____V__isize(map__0, t13)
    var t15 closure_env_main_2 = closure_env_main_2{}
    var t16 func(int) int = func(p0 int) int {
        return _goml_m_inherent_i_closure__env__main__2_i_closure__env__main__2_i_apply(t15, p0)
    }
    var t17 Option__isize = _goml_m_inherent_i_HashMapEntr_h9abc9eb3d8ded3aa79581ab7c647c1de_Key____V__isize(t14, t16)
    var t18 int = _goml_m_inherent_i_Option_i_Option_l_T_r__i_unwrap__or____T__isize(t17, -1)
    println__T_isize(t18)
    var t19 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(hashes__0)
    println__T_isize(t19)
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__isize(hashes__0, 0)
    var t20 Key = key(1, hashes__0)
    var t21 HashMapEntry__Key__isize = _goml_m_inherent_i_HashMap_i_HashMap_l_K_c_V_r__i_entry____K__Key____V__isize(map__0, t20)
    var t22 _goml_m_Option_____o_Key_c_isize_q_ = _goml_m_inherent_i_HashMapEntr_he7df3efa35857dbe82e86ebc4eb03056_Key____V__isize(t21)
    var t23 Key = key(0, hashes__0)
    var t24 Tuple2_3Key_3int = Tuple2_3Key_3int{
        _0: t23,
        _1: -1,
    }
    var t25 Tuple2_3Key_3int = _goml_m_inherent_i_Option_i_Option_l_T_r__i_unwrap__or____T___o_Key_c_isize_q_(t22, t24)
    var t26 int = t25._1
    println__T_isize(t26)
    var t27 int
    var inline31 int = ref_get__Ref_3int(hashes__0)
    t27 = inline31
    var inline29 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t27)
    _goml_runtime_core_string_println(inline29)
    var t28 int
    var inline28 int = hashmap_len__HashMap_3Key_3int(map__0)
    t28 = inline28
    var inline26 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t28)
    _goml_runtime_core_string_println(inline26)
    var native__0 *hashmap_string_int_x
    var inline25 *hashmap_string_int_x = hashmap_new__HashMap_6string_3int()
    native__0 = inline25
    var t29 Option__isize
    var inline22 string = "answer"
    var inline23 int = 41
    var inline24 Option__isize = hashmap_insert__HashMap_6string_3int(native__0, inline22, inline23)
    t29 = inline24
    var t30 bool
    var inline20 bool = _goml_m_inherent_i_Option_i_Option_l_T_r__i_is__some____T__isize(t29)
    var inline21 bool = !inline20
    t30 = inline21
    var inline18 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t30)
    _goml_runtime_core_string_println(inline18)
    var t31 HashMapEntry__string__isize
    var inline16 string = "answer"
    var inline17 HashMapEntry__string__isize = HashMapEntry__string__isize{
        values: native__0,
        key: inline16,
    }
    t31 = inline17
    var t32 closure_env_main_3 = closure_env_main_3{}
    var t33 func(int) int = func(p0 int) int {
        return _goml_m_inherent_i_closure__env__main__3_i_closure__env__main__3_i_apply(t32, p0)
    }
    var t34 HashMapEntry__string__isize
    var inline13 *hashmap_string_int_x = t31.values
    var inline14 string = t31.key
    _goml_m_inherent_i_HashMap_i_H_hfcee0d24f3f45608afc229b1d5b5eca7_ing____V__isize(inline13, inline14, t33)
    t34 = t31
    var t35 Option__isize
    var inline10 *hashmap_string_int_x = t34.values
    var inline11 string = t34.key
    var inline12 Option__isize = _goml_m_inherent_i_HashMap_i_HashMap_l_K_c_V_r__i_get____K__string____V__isize(inline10, inline11)
    t35 = inline12
    var t36 int
    var inline8 int = 0
    switch t35._tag {
    case 0:
        t36 = inline8
    case 1:
        var inline9 int = t35._p0
        t36 = inline9
    default:
        panic("non-exhaustive match")
    }
    var inline6 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t36)
    _goml_runtime_core_string_println(inline6)
    var t37 Option__isize
    var inline4 string = "answer"
    var inline5 Option__isize = hashmap_remove_value__HashMap_6string_3int(native__0, inline4)
    t37 = inline5
    var t38 int
    var inline2 int = 0
    switch t37._tag {
    case 0:
        t38 = inline2
    case 1:
        var inline3 int = t37._p0
        t38 = inline3
    default:
        panic("non-exhaustive match")
    }
    var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t38)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(self__0 *ref_int_x) int {
    var t0 int = ref_get__Ref_3int(self__0)
    return t0
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__isize(self__0 *ref_int_x, value__0 int) struct{} {
    ref_set__Ref_3int(self__0, value__0)
    return struct{}{}
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__isize(value__0 int) *ref_int_x {
    var t0 *ref_int_x = ref__Ref_3int(value__0)
    return t0
}

func _goml_m_inherent_i_HashMap_i_HashMap_l_K_c_V_r__i_new____K__Key____V__isize() *hashmap_Key_int_x {
    var t0 *hashmap_Key_int_x = hashmap_new__HashMap_3Key_3int()
    return t0
}

func println__T_isize(value__0 int) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_int_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_inherent_i_HashMap_i_HashMap_l_K_c_V_r__i_entry____K__Key____V__isize(self__0 *hashmap_Key_int_x, key__0 Key) HashMapEntry__Key__isize {
    var t0 HashMapEntry__Key__isize = HashMapEntry__Key__isize{
        values: self__0,
        key: key__0,
    }
    return t0
}

func _goml_m_inherent_i_HashMapEntr_h08451cacc3f0db11c02874f35c8558b6_Key____V__isize(self__0 HashMapEntry__Key__isize, create__0 func() int) int {
    var t0 *hashmap_Key_int_x = self__0.values
    var t1 Key = self__0.key
    var inline0 int = hashmap_get_or_insert_with__HashMap_3Key_3int(t0, t1, create__0)
    return inline0
}

func _goml_m_inherent_i_HashMapEntr_h9abc9eb3d8ded3aa79581ab7c647c1de_Key____V__isize(self__0 HashMapEntry__Key__isize, update__0 func(int) int) Option__isize {
    var t0 *hashmap_Key_int_x = self__0.values
    var t1 Key = self__0.key
    var inline0 Option__isize = hashmap_update__HashMap_3Key_3int(t0, t1, update__0)
    return inline0
}

func _goml_m_inherent_i_Option_i_Option_l_T_r__i_unwrap__or____T__isize(self__0 Option__isize, fallback__0 int) int {
    switch self__0._tag {
    case 0:
        return fallback__0
    case 1:
        var x0 int = self__0._p0
        return x0
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_inherent_i_HashMapEntr_he7df3efa35857dbe82e86ebc4eb03056_Key____V__isize(self__0 HashMapEntry__Key__isize) _goml_m_Option_____o_Key_c_isize_q_ {
    var t0 Option__isize
    var inline3 *hashmap_Key_int_x = self__0.values
    var inline4 Key = self__0.key
    var inline5 Option__isize = _goml_m_inherent_i_HashMap_i_H_h0e46f47d3ada08b881ff6860cdc9491c_Key____V__isize(inline3, inline4)
    t0 = inline5
    var t1 closure_env_inherent_HashMapEntry_HashMapEntry_K_V_remove_entry_K_Key_V_isize_4 = closure_env_inherent_HashMapEntry_HashMapEntry_K_V_remove_entry_K_Key_V_isize_4{
        self_0: self__0,
    }
    var t2 func(int) Tuple2_3Key_3int = func(p0 int) Tuple2_3Key_3int {
        return _goml_m_inherent_i_closure__en_h51dad037a23a0205e1d674d00df14e80_size__4_i_apply(t1, p0)
    }
    switch t0._tag {
    case 0:
        return _goml_m_Option_____o_Key_c_isize_q_{
            _tag: 0,
        }
    case 1:
        var inline0 int = t0._p0
        var inline1 Tuple2_3Key_3int = t2(inline0)
        var inline2 _goml_m_Option_____o_Key_c_isize_q_ = _goml_m_Option_____o_Key_c_isize_q_{
            _p0: inline1,
            _tag: 1,
        }
        return inline2
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_inherent_i_Option_i_Option_l_T_r__i_unwrap__or____T___o_Key_c_isize_q_(self__0 _goml_m_Option_____o_Key_c_isize_q_, fallback__0 Tuple2_3Key_3int) Tuple2_3Key_3int {
    switch self__0._tag {
    case 0:
        return fallback__0
    case 1:
        var x0 Tuple2_3Key_3int = self__0._p0
        return x0
    default:
        panic("non-exhaustive match")
    }
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

func _goml_m_inherent_i_Option_i_Option_l_T_r__i_is__some____T__isize(self__0 Option__isize) bool {
    switch self__0._tag {
    case 0:
        return false
    case 1:
        return true
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_inherent_i_HashMap_i_H_hfcee0d24f3f45608afc229b1d5b5eca7_ing____V__isize(self__0 *hashmap_string_int_x, key__0 string, update__0 func(int) int) Option__isize {
    var t0 Option__isize = hashmap_update__HashMap_6string_3int(self__0, key__0, update__0)
    return t0
}

func _goml_m_inherent_i_HashMap_i_HashMap_l_K_c_V_r__i_get____K__string____V__isize(self__0 *hashmap_string_int_x, key__0 string) Option__isize {
    var t0 Option__isize = hashmap_get__HashMap_6string_3int(self__0, key__0)
    return t0
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

func _goml_m_inherent_i_HashMap_i_H_h0e46f47d3ada08b881ff6860cdc9491c_Key____V__isize(self__0 *hashmap_Key_int_x, key__0 Key) Option__isize {
    var t0 Option__isize = hashmap_remove_value__HashMap_3Key_3int(self__0, key__0)
    return t0
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

func _goml_m_inherent_i_closure__env__main__0_i_closure__env__main__0_i_apply(env0 closure_env_main_0) int {
    var created__0 *ref_int_x = env0.created_0
    var t0 int
    var inline1 int = ref_get__Ref_3int(created__0)
    t0 = inline1
    var t1 int = t0 + 1
    ref_set__Ref_3int(created__0, t1)
    return 10
}

func _goml_m_inherent_i_closure__env__main__1_i_closure__env__main__1_i_apply(env0 closure_env_main_1) int {
    var created__0 *ref_int_x = env0.created_0
    var t0 int
    var inline1 int = ref_get__Ref_3int(created__0)
    t0 = inline1
    var t1 int = t0 + 1
    ref_set__Ref_3int(created__0, t1)
    return 99
}

func _goml_m_inherent_i_closure__env__main__2_i_closure__env__main__2_i_apply(env0 closure_env_main_2, value__0 int) int {
    var t0 int = value__0 + 5
    return t0
}

func _goml_m_inherent_i_closure__env__main__3_i_closure__env__main__3_i_apply(env0 closure_env_main_3, value__0 int) int {
    var t0 int = value__0 + 1
    return t0
}

func _goml_m_inherent_i_closure__en_h51dad037a23a0205e1d674d00df14e80_size__4_i_apply(env0 closure_env_inherent_HashMapEntry_HashMapEntry_K_V_remove_entry_K_Key_V_isize_4, value__0 int) Tuple2_3Key_3int {
    var self__0 HashMapEntry__Key__isize = env0.self_0
    var t0 Key = self__0.key
    var t1 Tuple2_3Key_3int = Tuple2_3Key_3int{
        _0: t0,
        _1: value__0,
    }
    return t1
}

func main() {
    main0()
}
