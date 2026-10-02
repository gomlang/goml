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

func _goml_runtime_core_string_len(s string) int {
    return int(len(s))
}

func _goml_runtime_core_string_get(s string, i int) rune {
    return rune(s[i])
}

func _goml_runtime_core_string_byte_get(s string, i int) uint8 {
    return s[i]
}

func _goml_runtime_core_string_byte_slice(s string, start int, end int) string {
    return s[start:end]
}

func _goml_runtime_core_string_to_bytes(s string) *_goml_vec_uint8 {
    return &_goml_vec_uint8{
        items: []byte(s),
    }
}

func _goml_runtime_core_string_from_utf8(bytes *_goml_vec_uint8) Tuple2_4bool_6string {
    return Tuple2_4bool_6string{
        _0: true,
        _1: string(bytes.items),
    }
}

func _goml_runtime_core_char_to_string(x rune) string {
    return string(x)
}

func _goml_runtime_core_char_from_uint32(value uint32) Tuple2_4bool_4char {
    return Tuple2_4bool_4char{
        _0: true,
        _1: rune(value),
    }
}

func _goml_runtime_core_string_println(s string) struct{} {
    _goml_os.Stdout.WriteString(s + "\n")
    return struct{}{}
}

type _goml_vec_uint8 struct {
    items []uint8
}

func vec_new__Vec_5uint8() *_goml_vec_uint8 {
    return &_goml_vec_uint8{
        items: nil,
    }
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

type _goml_vec_string struct {
    items []string
}

func vec_with_capacity__Vec_6string(capacity int) *_goml_vec_string {
    return &_goml_vec_string{
        items: make([]string, 0, capacity),
    }
}

func vec_push__Vec_6string(vec *_goml_vec_string, elem string) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_get__Vec_6string(vec *_goml_vec_string, index int) string {
    return vec.items[index]
}

func vec_len__Vec_6string(vec *_goml_vec_string) int {
    return int(len(vec.items))
}

func vec_capacity__Vec_6string(vec *_goml_vec_string) int {
    return int(cap(vec.items))
}

type _goml_vec_int32 struct {
    items []int32
}

func vec_with_capacity__Vec_5int32(capacity int) *_goml_vec_int32 {
    return &_goml_vec_int32{
        items: make([]int32, 0, capacity),
    }
}

func vec_push__Vec_5int32(vec *_goml_vec_int32, elem int32) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_get__Vec_5int32(vec *_goml_vec_int32, index int) int32 {
    return vec.items[index]
}

func vec_set__Vec_5int32(vec *_goml_vec_int32, index int, value int32) struct{} {
    vec.items[index] = value
    return struct{}{}
}

func vec_len__Vec_5int32(vec *_goml_vec_int32) int {
    return int(len(vec.items))
}

func vec_capacity__Vec_5int32(vec *_goml_vec_int32) int {
    return int(cap(vec.items))
}

func vec_reserve__Vec_5int32(vec *_goml_vec_int32, additional int) struct{} {
    if additional < 0 {
        panic("negative vector capacity")
    }
    var length int = len(vec.items)
    var required int = length + additional
    if required < length {
        panic("vector capacity overflow")
    }
    if required > cap(vec.items) {
        var next_capacity int = cap(vec.items) * 2
        if next_capacity < required {
            next_capacity = required
        }
        var next_items []int32 = make([]int32, length, next_capacity)
        copy(next_items, vec.items)
        vec.items = next_items
    }
    return struct{}{}
}

func vec_truncate__Vec_5int32(vec *_goml_vec_int32, new_len int) struct{} {
    if new_len < 0 {
        panic("negative vector length")
    }
    if new_len < int(len(vec.items)) {
        clear(vec.items[new_len:int(len(vec.items))])
        vec.items = vec.items[0:new_len]
    }
    return struct{}{}
}

type _goml_vec_Tuple2_6string_5int32 struct {
    items []Tuple2_6string_5int32
}

func vec_get__Vec_21Tuple2_6string_5int32(vec *_goml_vec_Tuple2_6string_5int32, index int) Tuple2_6string_5int32 {
    return vec.items[index]
}

func vec_len__Vec_21Tuple2_6string_5int32(vec *_goml_vec_Tuple2_6string_5int32) int {
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

type ref_bool_x struct {
    value bool
}

func ref__Ref_4bool(value bool) *ref_bool_x {
    return &ref_bool_x{
        value: value,
    }
}

func ref_get__Ref_4bool(reference *ref_bool_x) bool {
    return reference.value
}

func ref_set__Ref_4bool(reference *ref_bool_x, value bool) struct{} {
    reference.value = value
    return struct{}{}
}

type hashmap_string_int32_x_entry struct {
    active bool
    key string
    value int32
}

type hashmap_string_int32_x struct {
    indices map[string]int
    entries []hashmap_string_int32_x_entry
    len int
}

func hashmap_new__HashMap_6string_5int32() *hashmap_string_int32_x {
    return &hashmap_string_int32_x{
        indices: make(map[string]int),
        entries: nil,
        len: 0,
    }
}

func hashmap_set__HashMap_6string_5int32(m *hashmap_string_int32_x, key string, value int32) struct{} {
    if m == nil {
        return struct{}{}
    }
    var index int
    var found bool
    index, found = m.indices[key]
    if found {
        var entry hashmap_string_int32_x_entry = m.entries[index]
        if entry.active {
            m.entries[index].value = value
            return struct{}{}
        }
        m.entries[index] = hashmap_string_int32_x_entry{
            active: true,
            key: key,
            value: value,
        }
        m.len = m.len + 1
        return struct{}{}
    }
    index = len(m.entries)
    m.indices[key] = index
    m.entries = append(m.entries, hashmap_string_int32_x_entry{
        active: true,
        key: key,
        value: value,
    })
    m.len = m.len + 1
    return struct{}{}
}

func hashmap_remove__HashMap_6string_5int32(m *hashmap_string_int32_x, key string) struct{} {
    if m == nil {
        return struct{}{}
    }
    var index int
    var found bool
    index, found = m.indices[key]
    if !found {
        return struct{}{}
    }
    var entry hashmap_string_int32_x_entry = m.entries[index]
    if !entry.active {
        return struct{}{}
    }
    var zero hashmap_string_int32_x_entry
    m.entries[index] = zero
    m.len = m.len - 1
    return struct{}{}
}

func hashmap_entries__HashMap_6string_5int32(m *hashmap_string_int32_x) *_goml_vec_Tuple2_6string_5int32 {
    var result []Tuple2_6string_5int32
    if m == nil {
        return &_goml_vec_Tuple2_6string_5int32{
            items: result,
        }
    }
    for _, entry := range m.entries {
        if entry.active {
            result = append(result, Tuple2_6string_5int32{
                _0: entry.key,
                _1: entry.value,
            })
        }
    }
    return &_goml_vec_Tuple2_6string_5int32{
        items: result,
    }
}

type Tuple2_3int_4char struct {
    _0 int
    _1 rune
}

type Tuple2_4bool_6string struct {
    _0 bool
    _1 string
}

type Tuple2_6string_5int32 struct {
    _0 string
    _1 int32
}

type Tuple2_4char_3int struct {
    _0 rune
    _1 int
}

type Tuple3_4bool_4char_3int struct {
    _0 bool
    _1 rune
    _2 int
}

type Tuple2_4bool_4char struct {
    _0 bool
    _1 rune
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

type FnIterator__char struct {
    next_fn func() Option__char
}

type _goml_m_FnIterator_____o_isize_c_char_q_ struct {
    next_fn func() _goml_m_Option_____o_isize_c_char_q_
}

type closure_env_inherent_string_string_chars_0 struct {
    self_0 string
    index_1 *ref_int_x
}

type closure_env_inherent_string_string_char_indices_1 struct {
    index_0 *ref_int_x
    self_1 string
}

type Ordering uint8

type Option__char uint64

type _goml_m_Option_____o_isize_c_char_q_ struct {
    _p0 Tuple2_3int_4char
    _tag uint8
}

type _goml_m_Option_____o_char_c_isize_q_ struct {
    _p0 Tuple2_4char_3int
    _tag uint8
}

type Option__i32 uint64

func print_chars(value__0 string) struct{} {
    var t0 FnIterator__char
    var inline4 *ref_int_x = ref__Ref_3int(0)
    var inline5 closure_env_inherent_string_string_chars_0 = closure_env_inherent_string_string_chars_0{
        self_0: value__0,
        index_1: inline4,
    }
    var inline6 func() Option__char = func() Option__char {
        return _goml_m_inherent_i_closure__en_h3f9733c4625dbd2f543c79fa467f2508_hars__0_i_apply(inline5)
    }
    var inline7 FnIterator__char = _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__char(inline6)
    t0 = inline7
    var for_iter0 FnIterator__char
    for_iter0 = t0
    Loop_loop0:
    for {
        var for_next0 Option__char
        var inline2 func() Option__char = for_iter0.next_fn
        var inline3 Option__char = inline2()
        for_next0 = inline3
        switch for_next0 != Option__char(0) {
        case false:
            break Loop_loop0
        case true:
            var x0 rune = rune(uint64(for_next0) - 1)
            var inline0 string = _goml_m_trait__impl_i_ToString_i_char_i_to__string(x0)
            _goml_runtime_core_string_println(inline0)
            continue
        default:
            panic("non-exhaustive match")
        }
    }
    return struct{}{}
}

func print_char_indices(value__0 string) struct{} {
    var t0 _goml_m_FnIterator_____o_isize_c_char_q_
    var inline6 *ref_int_x = ref__Ref_3int(0)
    var inline7 closure_env_inherent_string_string_char_indices_1 = closure_env_inherent_string_string_char_indices_1{
        index_0: inline6,
        self_1: value__0,
    }
    var inline8 func() _goml_m_Option_____o_isize_c_char_q_ = func() _goml_m_Option_____o_isize_c_char_q_ {
        return _goml_m_inherent_i_closure__en_hf9055ebad38fcb339d5a880925418115_ices__1_i_apply(inline7)
    }
    var inline9 _goml_m_FnIterator_____o_isize_c_char_q_ = _goml_m_inherent_i_FnIterator__hae8fa7fd40ba2a3681525fbecfca8da7_isize_c_char_q_(inline8)
    t0 = inline9
    var for_iter0 _goml_m_FnIterator_____o_isize_c_char_q_
    for_iter0 = t0
    Loop_loop0:
    for {
        var for_next0 _goml_m_Option_____o_isize_c_char_q_
        var inline4 func() _goml_m_Option_____o_isize_c_char_q_ = for_iter0.next_fn
        var inline5 _goml_m_Option_____o_isize_c_char_q_ = inline4()
        for_next0 = inline5
        switch for_next0._tag {
        case 0:
            break Loop_loop0
        case 1:
            var x0 Tuple2_3int_4char = for_next0._p0
            var x1 int = x0._0
            var x2 rune = x0._1
            var t1 string
            var inline3 string = __goml_builtin_int_to_string(x1)
            t1 = inline3
            var t2 string = t1 + ":"
            var t3 string
            var inline2 string = char_to_string(x2)
            t3 = inline2
            var t4 string = t2 + t3
            var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t4)
            _goml_runtime_core_string_println(inline0)
            continue
        default:
            panic("non-exhaustive match")
        }
    }
    return struct{}{}
}

func main0() struct{} {
    var value__0 string = "a你好😀z"
    var t0 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
    println__T_isize(t0)
    var t1 int = _goml_m_inherent_i_string_i_string_i_len(value__0)
    println__T_isize(t1)
    var t2 rune = _goml_m_inherent_i_string_i_string_i_get(value__0, 0)
    println__T_char(t2)
    var t3 rune = _goml_m_inherent_i_string_i_string_i_get(value__0, 1)
    println__T_char(t3)
    var t4 rune = _goml_m_inherent_i_string_i_string_i_get(value__0, 4)
    println__T_char(t4)
    var t5 rune = _goml_m_inherent_i_string_i_string_i_get(value__0, 7)
    println__T_char(t5)
    var t6 rune = _goml_m_inherent_i_string_i_string_i_get(value__0, 11)
    println__T_char(t6)
    var t7 bool = _goml_m_inherent_i_string_i_string_i_is__char__boundary(value__0, 0)
    println__T_bool(t7)
    var t8 bool = _goml_m_inherent_i_string_i_string_i_is__char__boundary(value__0, 2)
    println__T_bool(t8)
    var t9 bool = _goml_m_inherent_i_string_i_string_i_is__char__boundary(value__0, 7)
    println__T_bool(t9)
    var t10 bool = _goml_m_inherent_i_string_i_string_i_is__char__boundary(value__0, 12)
    println__T_bool(t10)
    var t11 string = _goml_m_inherent_i_string_i_string_i_byte__slice(value__0, 1, 7)
    println__T_string(t11)
    var mtmp0 _goml_m_Option_____o_char_c_isize_q_ = _goml_m_inherent_i_string_i_string_i_decode__at(value__0, 7)
    switch mtmp0._tag {
    case 0:
        var inline32 string = "missing"
        var inline33 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline32)
        _goml_runtime_core_string_println(inline33)
    case 1:
        var x8 Tuple2_4char_3int = mtmp0._p0
        var x9 rune = x8._0
        var x10 int = x8._1
        var inline37 string = _goml_m_trait__impl_i_ToString_i_char_i_to__string(x9)
        _goml_runtime_core_string_println(inline37)
        var inline35 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(x10)
        _goml_runtime_core_string_println(inline35)
    default:
        panic("non-exhaustive match")
    }
    print_chars(value__0)
    print_char_indices(value__0)
    var bytes__0 *_goml_vec_uint8 = _goml_m_inherent_i_string_i_string_i_to__bytes(value__0)
    var t12 int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_len____T__u8(bytes__0)
    println__T_isize(t12)
    var t13 uint8 = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_get____T__u8(bytes__0, 0)
    println__T_u8(t13)
    var t14 uint8 = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_get____T__u8(bytes__0, 1)
    println__T_u8(t14)
    var mtmp1 Tuple2_4bool_6string = string_from_utf8(bytes__0)
    var x0 bool = mtmp1._0
    var x1 string = mtmp1._1
    println__T_bool(x0)
    println__T_string(x1)
    var invalid__0 *_goml_vec_uint8 = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_new____T__u8()
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_push____T__u8(invalid__0, 255)
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_push____T__u8(invalid__0, 254)
    var mtmp2 Tuple2_4bool_6string = string_from_utf8(invalid__0)
    var x2 bool = mtmp2._0
    var x3 string = mtmp2._1
    println__T_bool(x2)
    var t15 bool = x3 == ""
    println__T_bool(t15)
    var parts__0 *_goml_vec_string = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_with__capacity____T__string(3)
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_push____T__string(parts__0, "hello")
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_push____T__string(parts__0, "，")
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_push____T__string(parts__0, "世界")
    var t16 int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_capacity____T__string(parts__0)
    var t17 bool = t16 >= 3
    println__T_bool(t17)
    var t18 string = string_concat(parts__0)
    println__T_string(t18)
    var values__0 *_goml_vec_int32 = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_with__capacity____T__i32(1)
    var t19 int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_len____T__i32(values__0)
    println__T_isize(t19)
    var t20 int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_capacity____T__i32(values__0)
    var t21 bool = t20 >= 1
    println__T_bool(t21)
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_reserve____T__i32(values__0, 100)
    var t22 int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_capacity____T__i32(values__0)
    var t23 bool = t22 >= 100
    println__T_bool(t23)
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_push____T__i32(values__0, 1)
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_push____T__i32(values__0, 2)
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_push____T__i32(values__0, 3)
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_insert____T__i32(values__0, 1, 9)
    var t24 int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_len____T__i32(values__0)
    println__T_isize(t24)
    var t25 int32 = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_remove____T__i32(values__0, 2)
    println__T_i32(t25)
    var t26 int32 = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_swap__remove____T__i32(values__0, 0)
    println__T_i32(t26)
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_reverse____T__i32(values__0)
    var t27 int32 = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_get____T__i32(values__0, 0)
    println__T_i32(t27)
    var t28 int32 = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_get____T__i32(values__0, 1)
    println__T_i32(t28)
    var mtmp3 Option__i32 = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_last____T__i32(values__0)
    switch mtmp3 != Option__i32(0) {
    case false:
        var inline27 int = -1
        var inline28 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(inline27)
        _goml_runtime_core_string_println(inline28)
    case true:
        var x7 int32 = int32(int64(uint64(mtmp3) - 1) - 2147483648)
        var inline30 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(x7)
        _goml_runtime_core_string_println(inline30)
    default:
        panic("non-exhaustive match")
    }
    var mtmp4 Option__i32 = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pop____T__i32(values__0)
    switch mtmp4 != Option__i32(0) {
    case false:
        var inline22 int = -1
        var inline23 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(inline22)
        _goml_runtime_core_string_println(inline23)
    case true:
        var x6 int32 = int32(int64(uint64(mtmp4) - 1) - 2147483648)
        var inline25 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(x6)
        _goml_runtime_core_string_println(inline25)
    default:
        panic("non-exhaustive match")
    }
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_truncate____T__i32(values__0, 0)
    var t29 bool = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_is__empty____T__i32(values__0)
    println__T_bool(t29)
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_push____T__i32(values__0, 42)
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_clear____T__i32(values__0)
    var t30 int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_len____T__i32(values__0)
    println__T_isize(t30)
    var map__0 *hashmap_string_int32_x = _goml_m_inherent_i_HashMap_i_HashMap_l_K_c_V_r__i_new____K__string____V__i32()
    _goml_m_inherent_i_HashMap_i_HashMap_l_K_c_V_r__i_set____K__string____V__i32(map__0, "a", 1)
    _goml_m_inherent_i_HashMap_i_HashMap_l_K_c_V_r__i_set____K__string____V__i32(map__0, "b", 2)
    var entries__0 *_goml_vec_Tuple2_6string_5int32 = _goml_m_inherent_i_HashMap_i_HashMap_l_K_c_V_r__i_entries____K__string____V__i32(map__0)
    var t31 int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_len____T___o_string_c_i32_q_(entries__0)
    println__T_isize(t31)
    var inline19 string = "c"
    var inline20 int32 = 3
    hashmap_set__HashMap_6string_5int32(map__0, inline19, inline20)
    var inline17 string = "a"
    hashmap_remove__HashMap_6string_5int32(map__0, inline17)
    var t32 int
    var inline16 int = vec_len__Vec_21Tuple2_6string_5int32(entries__0)
    t32 = inline16
    var inline14 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t32)
    _goml_runtime_core_string_println(inline14)
    var seen_a__0 *ref_bool_x
    var inline12 bool = false
    var inline13 *ref_bool_x = ref__Ref_4bool(inline12)
    seen_a__0 = inline13
    var seen_b__0 *ref_bool_x
    var inline10 bool = false
    var inline11 *ref_bool_x = ref__Ref_4bool(inline10)
    seen_b__0 = inline11
    var for_limit0 int = vec_len__Vec_21Tuple2_6string_5int32(entries__0)
    var for_index0 int = 0
    Loop_loop0:
    for {
        var t35 bool = for_index0 < for_limit0
        if t35 {
            var for_item0 Tuple2_6string_5int32 = vec_get__Vec_21Tuple2_6string_5int32(entries__0, for_index0)
            var t36 int = for_index0 + 1
            for_index0 = t36
            var x4 string = for_item0._0
            var x5 int32 = for_item0._1
            var t37 bool = x4 == "a"
            var jp0 bool
            if t37 {
                var t40 bool = x5 == 1
                jp0 = t40
            } else {
                jp0 = false
            }
            if jp0 {
                var inline6 bool = true
                ref_set__Ref_4bool(seen_a__0, inline6)
                continue
            } else {
                var t38 bool = x4 == "b"
                var jp1 bool
                if t38 {
                    var t39 bool = x5 == 2
                    jp1 = t39
                } else {
                    jp1 = false
                }
                if jp1 {
                    var inline8 bool = true
                    ref_set__Ref_4bool(seen_b__0, inline8)
                    continue
                } else {
                    continue
                }
            }
        } else {
            break Loop_loop0
        }
    }
    var t33 bool
    var inline5 bool = ref_get__Ref_4bool(seen_a__0)
    t33 = inline5
    var inline3 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t33)
    _goml_runtime_core_string_println(inline3)
    var t34 bool
    var inline2 bool = ref_get__Ref_4bool(seen_b__0)
    t34 = inline2
    var inline0 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t34)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func println__T_char(value__0 rune) struct{} {
    var t0 string
    var inline0 string = char_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func println__T_string(value__0 string) struct{} {
    var t0 string
    t0 = value__0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func println__T_isize(value__0 int) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_int_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_inherent_i_string_i_string_i_byte__len(self__0 string) int {
    var t0 int = _goml_runtime_core_string_len(self__0)
    return t0
}

func _goml_m_inherent_i_string_i_string_i_len(self__0 string) int {
    var t0 int = _goml_runtime_core_string_len(self__0)
    return t0
}

func _goml_m_inherent_i_string_i_string_i_get(self__0 string, index__0 int) rune {
    var inline0 Tuple3_4bool_4char_3int = string_decode_utf8_at(self__0, index__0)
    var inline1 bool = inline0._0
    var inline2 rune = inline0._1
    if inline1 {
        return inline2
    } else {
        var inline3 rune = _goml_runtime_core_string_get("", -1)
        return inline3
    }
}

func println__T_bool(value__0 bool) struct{} {
    var t0 string
    var inline0 string = _goml_runtime_core_bool_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_inherent_i_string_i_string_i_is__char__boundary(self__0 string, index__0 int) bool {
    var t0 bool = string_is_char_boundary(self__0, index__0)
    return t0
}

func _goml_m_inherent_i_string_i_string_i_byte__slice(self__0 string, start__0 int, end__0 int) string {
    var inline0 bool = string_is_char_boundary(self__0, start__0)
    var inline1 bool
    if inline0 {
        var inline4 bool = string_is_char_boundary(self__0, end__0)
        inline1 = inline4
    } else {
        inline1 = false
    }
    if inline1 {
        var inline2 string = _goml_runtime_core_string_byte_slice(self__0, start__0, end__0)
        return inline2
    } else {
        var inline3 string = _goml_runtime_core_string_byte_slice(self__0, -1, -1)
        return inline3
    }
}

func _goml_m_inherent_i_string_i_string_i_decode__at(self__0 string, index__0 int) _goml_m_Option_____o_char_c_isize_q_ {
    var mtmp0 Tuple3_4bool_4char_3int = string_decode_utf8_at(self__0, index__0)
    var x0 bool = mtmp0._0
    var x1 rune = mtmp0._1
    var x2 int = mtmp0._2
    if x0 {
        var t0 Tuple2_4char_3int = Tuple2_4char_3int{
            _0: x1,
            _1: x2,
        }
        var t1 _goml_m_Option_____o_char_c_isize_q_ = _goml_m_Option_____o_char_c_isize_q_{
            _p0: t0,
            _tag: 1,
        }
        return t1
    } else {
        return _goml_m_Option_____o_char_c_isize_q_{
            _tag: 0,
        }
    }
}

func _goml_m_inherent_i_string_i_string_i_to__bytes(self__0 string) *_goml_vec_uint8 {
    var t0 *_goml_vec_uint8 = _goml_runtime_core_string_to_bytes(self__0)
    return t0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_len____T__u8(self__0 *_goml_vec_uint8) int {
    var t0 int = vec_len__Vec_5uint8(self__0)
    return t0
}

func println__T_u8(value__0 uint8) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_uint8_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_get____T__u8(self__0 *_goml_vec_uint8, index__0 int) uint8 {
    var t0 uint8 = vec_get__Vec_5uint8(self__0, index__0)
    return t0
}

func string_from_utf8(bytes__0 *_goml_vec_uint8) Tuple2_4bool_6string {
    var mtmp0 Tuple2_4bool_6string = _goml_runtime_core_string_from_utf8(bytes__0)
    var x0 string = mtmp0._1
    var index__0 int = 0
    Loop_loop0:
    for {
        var t1 int
        var inline0 int = _goml_runtime_core_string_len(x0)
        t1 = inline0
        var t2 bool = index__0 < t1
        if t2 {
            var mtmp1 Tuple3_4bool_4char_3int = string_decode_utf8_at(x0, index__0)
            var x1 bool = mtmp1._0
            var x2 int = mtmp1._2
            if x1 {
                var compound_old0 int = index__0
                var t3 int = compound_old0 + x2
                index__0 = t3
                continue
            } else {
                var t5 Tuple2_4bool_6string = Tuple2_4bool_6string{
                    _0: false,
                    _1: "",
                }
                return t5
            }
        } else {
            break Loop_loop0
        }
    }
    var t0 Tuple2_4bool_6string = Tuple2_4bool_6string{
        _0: true,
        _1: x0,
    }
    return t0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_new____T__u8() *_goml_vec_uint8 {
    var t0 *_goml_vec_uint8 = vec_new__Vec_5uint8()
    return t0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_push____T__u8(self__0 *_goml_vec_uint8, elem__0 uint8) struct{} {
    vec_push__Vec_5uint8(self__0, elem__0)
    return struct{}{}
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_with__capacity____T__string(capacity__0 int) *_goml_vec_string {
    var t0 *_goml_vec_string = vec_with_capacity__Vec_6string(capacity__0)
    return t0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_push____T__string(self__0 *_goml_vec_string, elem__0 string) struct{} {
    vec_push__Vec_6string(self__0, elem__0)
    return struct{}{}
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_capacity____T__string(self__0 *_goml_vec_string) int {
    var t0 int = vec_capacity__Vec_6string(self__0)
    return t0
}

func string_concat(values__0 *_goml_vec_string) string {
    var t0 string = __goml_builtin_string_concat(values__0)
    return t0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_with__capacity____T__i32(capacity__0 int) *_goml_vec_int32 {
    var t0 *_goml_vec_int32 = vec_with_capacity__Vec_5int32(capacity__0)
    return t0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_len____T__i32(self__0 *_goml_vec_int32) int {
    var t0 int = vec_len__Vec_5int32(self__0)
    return t0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_capacity____T__i32(self__0 *_goml_vec_int32) int {
    var t0 int = vec_capacity__Vec_5int32(self__0)
    return t0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_reserve____T__i32(self__0 *_goml_vec_int32, additional__0 int) struct{} {
    vec_reserve__Vec_5int32(self__0, additional__0)
    return struct{}{}
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_push____T__i32(self__0 *_goml_vec_int32, elem__0 int32) struct{} {
    vec_push__Vec_5int32(self__0, elem__0)
    return struct{}{}
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_insert____T__i32(self__0 *_goml_vec_int32, index__0 int, value__0 int32) struct{} {
    var len__0 int
    var inline2 int = vec_len__Vec_5int32(self__0)
    len__0 = inline2
    var t0 bool = index__0 == len__0
    if t0 {
        vec_push__Vec_5int32(self__0, value__0)
        return struct{}{}
    } else {
        vec_get__Vec_5int32(self__0, index__0)
        var t1 int = len__0 - 1
        var t2 int32 = vec_get__Vec_5int32(self__0, t1)
        vec_push__Vec_5int32(self__0, t2)
        var current__0 int = len__0 - 1
        Loop_loop0:
        for {
            var t4 bool = current__0 > index__0
            if t4 {
                var index0 int = current__0
                vec_get__Vec_5int32(self__0, index0)
                var t5 int = current__0 - 1
                var value0 int32 = vec_get__Vec_5int32(self__0, t5)
                vec_set__Vec_5int32(self__0, index0, value0)
                var compound_old0 int = current__0
                var compound_value0 int = 1
                var t7 int = compound_old0 - compound_value0
                current__0 = t7
                continue
            } else {
                break Loop_loop0
            }
        }
        vec_get__Vec_5int32(self__0, index__0)
        vec_set__Vec_5int32(self__0, index__0, value__0)
        return struct{}{}
    }
}

func println__T_i32(value__0 int32) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_int32_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_remove____T__i32(self__0 *_goml_vec_int32, index__0 int) int32 {
    var len__0 int
    var inline1 int = vec_len__Vec_5int32(self__0)
    len__0 = inline1
    var value__0 int32 = vec_get__Vec_5int32(self__0, index__0)
    var current__0 int = index__0
    Loop_loop0:
    for {
        var t1 int = current__0 + 1
        var t2 bool = t1 < len__0
        if t2 {
            var index0 int = current__0
            vec_get__Vec_5int32(self__0, index0)
            var t3 int = current__0 + 1
            var value0 int32 = vec_get__Vec_5int32(self__0, t3)
            vec_set__Vec_5int32(self__0, index0, value0)
            var compound_old0 int = current__0
            var compound_value0 int = 1
            var t5 int = compound_old0 + compound_value0
            current__0 = t5
            continue
        } else {
            break Loop_loop0
        }
    }
    var t0 int = len__0 - 1
    vec_truncate__Vec_5int32(self__0, t0)
    return value__0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_swap__remove____T__i32(self__0 *_goml_vec_int32, index__0 int) int32 {
    var len__0 int
    var inline1 int = vec_len__Vec_5int32(self__0)
    len__0 = inline1
    var value__0 int32 = vec_get__Vec_5int32(self__0, index__0)
    var t0 int = index__0 + 1
    var t1 bool = t0 < len__0
    if t1 {
        vec_get__Vec_5int32(self__0, index__0)
        var t3 int = len__0 - 1
        var value0 int32 = vec_get__Vec_5int32(self__0, t3)
        vec_set__Vec_5int32(self__0, index__0, value0)
    } else {}
    var t2 int = len__0 - 1
    vec_truncate__Vec_5int32(self__0, t2)
    return value__0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_reverse____T__i32(self__0 *_goml_vec_int32) struct{} {
    var left__0 int = 0
    var t0 int
    var inline6 int = vec_len__Vec_5int32(self__0)
    t0 = inline6
    var right__0 int = t0 - 1
    Loop_loop0:
    for {
        var t1 bool = left__0 < right__0
        if t1 {
            var inline0 int32 = vec_get__Vec_5int32(self__0, left__0)
            vec_get__Vec_5int32(self__0, left__0)
            var inline2 int32 = vec_get__Vec_5int32(self__0, right__0)
            vec_set__Vec_5int32(self__0, left__0, inline2)
            vec_get__Vec_5int32(self__0, right__0)
            vec_set__Vec_5int32(self__0, right__0, inline0)
            var compound_old0 int = left__0
            var compound_value0 int = 1
            var t2 int = compound_old0 + compound_value0
            left__0 = t2
            var compound_old1 int = right__0
            var compound_value1 int = 1
            var t4 int = compound_old1 - compound_value1
            right__0 = t4
            continue
        } else {
            break Loop_loop0
        }
    }
    return struct{}{}
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_get____T__i32(self__0 *_goml_vec_int32, index__0 int) int32 {
    var t0 int32 = vec_get__Vec_5int32(self__0, index__0)
    return t0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_last____T__i32(self__0 *_goml_vec_int32) Option__i32 {
    var len__0 int
    var inline0 int = vec_len__Vec_5int32(self__0)
    len__0 = inline0
    var t0 bool = len__0 == 0
    if t0 {
        return Option__i32(0)
    } else {
        var t1 int = len__0 - 1
        var t2 int32 = vec_get__Vec_5int32(self__0, t1)
        var t3 Option__i32 = Option__i32(uint64(int64(t2) + 2147483648) + 1)
        return t3
    }
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pop____T__i32(self__0 *_goml_vec_int32) Option__i32 {
    var len__0 int
    var inline1 int = vec_len__Vec_5int32(self__0)
    len__0 = inline1
    var t0 bool = len__0 == 0
    if t0 {
        return Option__i32(0)
    } else {
        var t1 int = len__0 - 1
        var value__0 int32 = vec_get__Vec_5int32(self__0, t1)
        var t2 int = len__0 - 1
        vec_truncate__Vec_5int32(self__0, t2)
        var t3 Option__i32 = Option__i32(uint64(int64(value__0) + 2147483648) + 1)
        return t3
    }
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_truncate____T__i32(self__0 *_goml_vec_int32, len__0 int) struct{} {
    vec_truncate__Vec_5int32(self__0, len__0)
    return struct{}{}
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_is__empty____T__i32(self__0 *_goml_vec_int32) bool {
    var t0 int = vec_len__Vec_5int32(self__0)
    var t1 bool = t0 == 0
    return t1
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_clear____T__i32(self__0 *_goml_vec_int32) struct{} {
    var inline0 int = 0
    vec_truncate__Vec_5int32(self__0, inline0)
    return struct{}{}
}

func _goml_m_inherent_i_HashMap_i_HashMap_l_K_c_V_r__i_new____K__string____V__i32() *hashmap_string_int32_x {
    var t0 *hashmap_string_int32_x = hashmap_new__HashMap_6string_5int32()
    return t0
}

func _goml_m_inherent_i_HashMap_i_HashMap_l_K_c_V_r__i_set____K__string____V__i32(self__0 *hashmap_string_int32_x, key__0 string, value__0 int32) struct{} {
    hashmap_set__HashMap_6string_5int32(self__0, key__0, value__0)
    return struct{}{}
}

func _goml_m_inherent_i_HashMap_i_HashMap_l_K_c_V_r__i_entries____K__string____V__i32(self__0 *hashmap_string_int32_x) *_goml_vec_Tuple2_6string_5int32 {
    var t0 *_goml_vec_Tuple2_6string_5int32 = hashmap_entries__HashMap_6string_5int32(self__0)
    return t0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_len____T___o_string_c_i32_q_(self__0 *_goml_vec_Tuple2_6string_5int32) int {
    var t0 int = vec_len__Vec_21Tuple2_6string_5int32(self__0)
    return t0
}

func _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__char(next_fn__0 func() Option__char) FnIterator__char {
    var t0 FnIterator__char = FnIterator__char{
        next_fn: next_fn__0,
    }
    return t0
}

func _goml_m_trait__impl_i_ToString_i_char_i_to__string(self__0 rune) string {
    var inline0 uint32 = uint32(rune(self__0))
    var inline1 bool = utf8_valid_scalar(inline0)
    if inline1 {
        var inline2 string = _goml_runtime_core_char_to_string(self__0)
        return inline2
    } else {
        _goml_runtime_core_string_get("", -1)
        return ""
    }
}

func _goml_m_inherent_i_FnIterator__hae8fa7fd40ba2a3681525fbecfca8da7_isize_c_char_q_(next_fn__0 func() _goml_m_Option_____o_isize_c_char_q_) _goml_m_FnIterator_____o_isize_c_char_q_ {
    var t0 _goml_m_FnIterator_____o_isize_c_char_q_ = _goml_m_FnIterator_____o_isize_c_char_q_{
        next_fn: next_fn__0,
    }
    return t0
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
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

func char_to_string(value__0 rune) string {
    var t0 uint32 = uint32(rune(value__0))
    var t1 bool
    var inline0 bool = t0 <= 1114111
    if inline0 {
        var inline1 bool = t0 >= 55296
        var inline2 bool
        if inline1 {
            var inline4 bool = t0 <= 57343
            inline2 = inline4
        } else {
            inline2 = false
        }
        var inline3 bool = !inline2
        t1 = inline3
    } else {
        t1 = false
    }
    if t1 {
        var t2 string = _goml_runtime_core_char_to_string(value__0)
        return t2
    } else {
        _goml_runtime_core_string_get("", -1)
        return ""
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

func string_is_char_boundary(value__0 string, index__0 int) bool {
    var t0 bool = index__0 < 0
    var jp0 bool
    if t0 {
        jp0 = true
    } else {
        var t6 int
        var inline2 int = _goml_runtime_core_string_len(value__0)
        t6 = inline2
        var t7 bool = index__0 > t6
        jp0 = t7
    }
    if jp0 {
        return false
    } else {
        var t1 int
        var inline1 int = _goml_runtime_core_string_len(value__0)
        t1 = inline1
        var t2 bool = index__0 == t1
        if t2 {
            return true
        } else {
            var t3 uint8
            var inline0 uint8 = _goml_runtime_core_string_byte_get(value__0, index__0)
            t3 = inline0
            var t4 uint8 = t3 & 192
            var t5 bool = t4 != 128
            return t5
        }
    }
}

func string_decode_utf8_at(value__0 string, index__0 int) Tuple3_4bool_4char_3int {
    var length__0 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
    var t0 bool = index__0 < 0
    var jp0 bool
    if t0 {
        jp0 = true
    } else {
        var t63 bool = index__0 >= length__0
        jp0 = t63
    }
    if jp0 {
        var inline25 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
            _0: false,
            _1: 0,
            _2: 0,
        }
        return inline25
    } else {
        var t1 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, index__0)
        var first__0 uint32 = uint32(uint8(t1))
        var t2 bool = first__0 < 128
        if t2 {
            var inline0 int = 1
            var inline1 Option__char = __goml_builtin_char_from_uint32(first__0)
            switch inline1 != Option__char(0) {
            case false:
                var inline2 Tuple3_4bool_4char_3int = utf8_invalid_decode()
                return inline2
            case true:
                var inline3 rune = rune(uint64(inline1) - 1)
                var inline4 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                    _0: true,
                    _1: inline3,
                    _2: inline0,
                }
                return inline4
            default:
                panic("non-exhaustive match")
            }
        } else {
            var t3 bool = first__0 < 194
            if t3 {
                var inline5 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                    _0: false,
                    _1: 0,
                    _2: 0,
                }
                return inline5
            } else {
                var t4 bool = first__0 < 224
                if t4 {
                    var t5 int = length__0 - index__0
                    var t6 bool = t5 < 2
                    if t6 {
                        var inline15 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                            _0: false,
                            _1: 0,
                            _2: 0,
                        }
                        return inline15
                    } else {
                        var t7 int = index__0 + 1
                        var t8 uint8
                        var inline14 uint8 = _goml_runtime_core_string_byte_get(value__0, t7)
                        t8 = inline14
                        var second__0 uint32 = uint32(uint8(t8))
                        var t9 bool
                        var inline12 bool = second__0 < 128
                        if inline12 {
                            t9 = true
                        } else {
                            var inline13 bool = second__0 > 191
                            t9 = inline13
                        }
                        if t9 {
                            var inline6 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                                _0: false,
                                _1: 0,
                                _2: 0,
                            }
                            return inline6
                        } else {
                            var t10 uint32 = first__0 & 31
                            var t11 uint32 = t10 << 6
                            var t12 uint32 = second__0 & 63
                            var t13 uint32 = t11 | t12
                            var inline7 int = 2
                            var inline8 Option__char = __goml_builtin_char_from_uint32(t13)
                            switch inline8 != Option__char(0) {
                            case false:
                                var inline9 Tuple3_4bool_4char_3int = utf8_invalid_decode()
                                return inline9
                            case true:
                                var inline10 rune = rune(uint64(inline8) - 1)
                                var inline11 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                                    _0: true,
                                    _1: inline10,
                                    _2: inline7,
                                }
                                return inline11
                            default:
                                panic("non-exhaustive match")
                            }
                        }
                    }
                } else {
                    var t14 bool = first__0 < 240
                    if t14 {
                        var t15 int = length__0 - index__0
                        var t16 bool = t15 < 3
                        if t16 {
                            var inline24 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                                _0: false,
                                _1: 0,
                                _2: 0,
                            }
                            return inline24
                        } else {
                            var t17 int = index__0 + 1
                            var t18 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, t17)
                            var second__1 uint32 = uint32(uint8(t18))
                            var t19 int = index__0 + 2
                            var t20 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, t19)
                            var third__0 uint32 = uint32(uint8(t20))
                            var t21 bool = utf8_invalid_continuation(second__1)
                            var jp1 bool
                            if t21 {
                                jp1 = true
                            } else {
                                var inline22 bool = third__0 < 128
                                if inline22 {
                                    jp1 = true
                                } else {
                                    var inline23 bool = third__0 > 191
                                    jp1 = inline23
                                }
                            }
                            var jp2 bool
                            if jp1 {
                                jp2 = true
                            } else {
                                var t31 bool = first__0 == 224
                                if t31 {
                                    var t32 bool = second__1 < 160
                                    jp2 = t32
                                } else {
                                    jp2 = false
                                }
                            }
                            var jp3 bool
                            if jp2 {
                                jp3 = true
                            } else {
                                var t29 bool = first__0 == 237
                                if t29 {
                                    var t30 bool = second__1 >= 160
                                    jp3 = t30
                                } else {
                                    jp3 = false
                                }
                            }
                            if jp3 {
                                var inline16 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                                    _0: false,
                                    _1: 0,
                                    _2: 0,
                                }
                                return inline16
                            } else {
                                var t22 uint32 = first__0 & 15
                                var t23 uint32 = t22 << 12
                                var t24 uint32 = second__1 & 63
                                var t25 uint32 = t24 << 6
                                var t26 uint32 = t23 | t25
                                var t27 uint32 = third__0 & 63
                                var t28 uint32 = t26 | t27
                                var inline17 int = 3
                                var inline18 Option__char = __goml_builtin_char_from_uint32(t28)
                                switch inline18 != Option__char(0) {
                                case false:
                                    var inline19 Tuple3_4bool_4char_3int = utf8_invalid_decode()
                                    return inline19
                                case true:
                                    var inline20 rune = rune(uint64(inline18) - 1)
                                    var inline21 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                                        _0: true,
                                        _1: inline20,
                                        _2: inline17,
                                    }
                                    return inline21
                                default:
                                    panic("non-exhaustive match")
                                }
                            }
                        }
                    } else {
                        var t33 bool = first__0 < 245
                        if t33 {
                            var t34 int = length__0 - index__0
                            var t35 bool = t34 < 4
                            if t35 {
                                var t61 Tuple3_4bool_4char_3int = utf8_invalid_decode()
                                return t61
                            } else {
                                var t36 int = index__0 + 1
                                var t37 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, t36)
                                var second__2 uint32 = uint32(uint8(t37))
                                var t38 int = index__0 + 2
                                var t39 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, t38)
                                var third__1 uint32 = uint32(uint8(t39))
                                var t40 int = index__0 + 3
                                var t41 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, t40)
                                var fourth__0 uint32 = uint32(uint8(t41))
                                var t42 bool = utf8_invalid_continuation(second__2)
                                var jp4 bool
                                if t42 {
                                    jp4 = true
                                } else {
                                    var t60 bool = utf8_invalid_continuation(third__1)
                                    jp4 = t60
                                }
                                var jp5 bool
                                if jp4 {
                                    jp5 = true
                                } else {
                                    var t59 bool = utf8_invalid_continuation(fourth__0)
                                    jp5 = t59
                                }
                                var jp6 bool
                                if jp5 {
                                    jp6 = true
                                } else {
                                    var t57 bool = first__0 == 240
                                    if t57 {
                                        var t58 bool = second__2 < 144
                                        jp6 = t58
                                    } else {
                                        jp6 = false
                                    }
                                }
                                var jp7 bool
                                if jp6 {
                                    jp7 = true
                                } else {
                                    var t55 bool = first__0 == 244
                                    if t55 {
                                        var t56 bool = second__2 > 143
                                        jp7 = t56
                                    } else {
                                        jp7 = false
                                    }
                                }
                                if jp7 {
                                    var t43 Tuple3_4bool_4char_3int = utf8_invalid_decode()
                                    return t43
                                } else {
                                    var t44 uint32 = first__0 & 7
                                    var t45 uint32 = t44 << 18
                                    var t46 uint32 = second__2 & 63
                                    var t47 uint32 = t46 << 12
                                    var t48 uint32 = t45 | t47
                                    var t49 uint32 = third__1 & 63
                                    var t50 uint32 = t49 << 6
                                    var t51 uint32 = t48 | t50
                                    var t52 uint32 = fourth__0 & 63
                                    var t53 uint32 = t51 | t52
                                    var t54 Tuple3_4bool_4char_3int = utf8_valid_decode(t53, 4)
                                    return t54
                                }
                            }
                        } else {
                            var t62 Tuple3_4bool_4char_3int = utf8_invalid_decode()
                            return t62
                        }
                    }
                }
            }
        }
    }
}

func __goml_builtin_string_concat(values__0 *_goml_vec_string) string {
    var length__0 int = 0
    var value_index__0 int = 0
    Loop_loop0:
    for {
        var t9 int
        var inline5 int = vec_len__Vec_6string(values__0)
        t9 = inline5
        var t10 bool = value_index__0 < t9
        if t10 {
            var compound_old2 int = length__0
            var t11 string = vec_get__Vec_6string(values__0, value_index__0)
            var compound_value2 int
            var inline4 int = _goml_runtime_core_string_len(t11)
            compound_value2 = inline4
            var t12 int = compound_old2 + compound_value2
            length__0 = t12
            var compound_old3 int = value_index__0
            var compound_value3 int = 1
            var t14 int = compound_old3 + compound_value3
            value_index__0 = t14
            continue
        } else {
            break Loop_loop0
        }
    }
    var bytes__0 *_goml_vec_uint8 = vec_with_capacity__Vec_5uint8(length__0)
    value_index__0 = 0
    Loop_loop1:
    for {
        var t0 int
        var inline3 int = vec_len__Vec_6string(values__0)
        t0 = inline3
        var t1 bool = value_index__0 < t0
        if t1 {
            var value__0 string = vec_get__Vec_6string(values__0, value_index__0)
            var byte_index__0 int = 0
            Loop_loop2:
            for {
                var t4 int
                var inline2 int = _goml_runtime_core_string_len(value__0)
                t4 = inline2
                var t5 bool = byte_index__0 < t4
                if t5 {
                    var t6 uint8
                    var inline1 uint8 = _goml_runtime_core_string_byte_get(value__0, byte_index__0)
                    t6 = inline1
                    vec_push__Vec_5uint8(bytes__0, t6)
                    var compound_old1 int = byte_index__0
                    var compound_value1 int = 1
                    var t7 int = compound_old1 + compound_value1
                    byte_index__0 = t7
                    continue
                } else {
                    break Loop_loop2
                }
            }
            var compound_old0 int = value_index__0
            var compound_value0 int = 1
            var t2 int = compound_old0 + compound_value0
            value_index__0 = t2
            continue
        } else {
            break Loop_loop1
        }
    }
    var mtmp0 Tuple2_4bool_6string = _goml_runtime_core_string_from_utf8(bytes__0)
    var x0 string = mtmp0._1
    return x0
}

func _goml_m_trait__impl_i_ToString_i_i32_i_to__string(self__0 int32) string {
    var inline0 int64 = int64(int32(self__0))
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

func utf8_valid_scalar(value__0 uint32) bool {
    var t0 bool = value__0 <= 1114111
    if t0 {
        var t1 bool = value__0 >= 55296
        var jp0 bool
        if t1 {
            var t3 bool = value__0 <= 57343
            jp0 = t3
        } else {
            jp0 = false
        }
        var t2 bool = !jp0
        return t2
    } else {
        return false
    }
}

func _goml_m_inherent_i_string_i_string_i_byte__get(self__0 string, index__0 int) uint8 {
    var t0 uint8 = _goml_runtime_core_string_byte_get(self__0, index__0)
    return t0
}

func utf8_invalid_decode() Tuple3_4bool_4char_3int {
    var t0 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
        _0: false,
        _1: 0,
        _2: 0,
    }
    return t0
}

func utf8_valid_decode(value__0 uint32, width__0 int) Tuple3_4bool_4char_3int {
    var commute_field0 rune
    var inline1 bool = utf8_valid_scalar(value__0)
    if inline1 {
        var inline2 Tuple2_4bool_4char = _goml_runtime_core_char_from_uint32(value__0)
        var inline3 rune = inline2._1
        commute_field0 = inline3
        var t0 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
            _0: true,
            _1: commute_field0,
            _2: width__0,
        }
        return t0
    } else {
        var inline0 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
            _0: false,
            _1: 0,
            _2: 0,
        }
        return inline0
    }
}

func utf8_invalid_continuation(value__0 uint32) bool {
    var t0 bool = value__0 < 128
    if t0 {
        return true
    } else {
        var t1 bool = value__0 > 191
        return t1
    }
}

func __goml_builtin_uint8_to_string(value__0 uint8) string {
    var t0 uint64 = uint64(uint8(value__0))
    var t1 string = decimal_string(t0)
    return t1
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

func __goml_builtin_char_from_uint32(value__0 uint32) Option__char {
    var t0 bool
    var inline0 bool = value__0 <= 1114111
    if inline0 {
        var inline1 bool = value__0 >= 55296
        var inline2 bool
        if inline1 {
            var inline4 bool = value__0 <= 57343
            inline2 = inline4
        } else {
            inline2 = false
        }
        var inline3 bool = !inline2
        t0 = inline3
    } else {
        t0 = false
    }
    if t0 {
        var mtmp0 Tuple2_4bool_4char = _goml_runtime_core_char_from_uint32(value__0)
        var x0 rune = mtmp0._1
        var t1 Option__char = Option__char(uint64(uint32(x0)) + 1)
        return t1
    } else {
        return Option__char(0)
    }
}

func _goml_m_inherent_i_closure__en_h3f9733c4625dbd2f543c79fa467f2508_hars__0_i_apply(env0 closure_env_inherent_string_string_chars_0) Option__char {
    var self__0 string = env0.self_0
    var index__0 *ref_int_x = env0.index_1
    var t0 int = ref_get__Ref_3int(index__0)
    var commute_field0 Tuple2_4char_3int
    var inline0 Tuple3_4bool_4char_3int = string_decode_utf8_at(self__0, t0)
    var inline1 bool = inline0._0
    var inline2 rune = inline0._1
    var inline3 int = inline0._2
    if inline1 {
        var inline4 Tuple2_4char_3int = Tuple2_4char_3int{
            _0: inline2,
            _1: inline3,
        }
        commute_field0 = inline4
        var x0 rune = commute_field0._0
        var x1 int = commute_field0._1
        var compound_old0 int = ref_get__Ref_3int(index__0)
        var t1 int = compound_old0 + x1
        ref_set__Ref_3int(index__0, t1)
        var t3 Option__char = Option__char(uint64(uint32(x0)) + 1)
        return t3
    } else {
        return Option__char(0)
    }
}

func _goml_m_inherent_i_closure__en_hf9055ebad38fcb339d5a880925418115_ices__1_i_apply(env0 closure_env_inherent_string_string_char_indices_1) _goml_m_Option_____o_isize_c_char_q_ {
    var index__0 *ref_int_x = env0.index_0
    var self__0 string = env0.self_1
    var current__0 int = ref_get__Ref_3int(index__0)
    var commute_field0 Tuple2_4char_3int
    var inline0 Tuple3_4bool_4char_3int = string_decode_utf8_at(self__0, current__0)
    var inline1 bool = inline0._0
    var inline2 rune = inline0._1
    var inline3 int = inline0._2
    if inline1 {
        var inline4 Tuple2_4char_3int = Tuple2_4char_3int{
            _0: inline2,
            _1: inline3,
        }
        commute_field0 = inline4
        var x0 rune = commute_field0._0
        var x1 int = commute_field0._1
        var t0 int = current__0 + x1
        ref_set__Ref_3int(index__0, t0)
        var t1 Tuple2_3int_4char = Tuple2_3int_4char{
            _0: current__0,
            _1: x0,
        }
        var t2 _goml_m_Option_____o_isize_c_char_q_ = _goml_m_Option_____o_isize_c_char_q_{
            _p0: t1,
            _tag: 1,
        }
        return t2
    } else {
        return _goml_m_Option_____o_isize_c_char_q_{
            _tag: 0,
        }
    }
}

func main() {
    main0()
}
