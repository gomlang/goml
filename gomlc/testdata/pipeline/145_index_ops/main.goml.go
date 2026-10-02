package main

import (
    _goml_os "os"
)

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

func array_get__Array_2_3int(arr [2]int, index int) int {
    return arr[index]
}

func array_set__Array_2_3int(arr [2]int, index int, value int) [2]int {
    arr[index] = value
    return arr
}

func array_get__Array_3_3int(arr [3]int, index int) int {
    return arr[index]
}

func array_set__Array_3_3int(arr [3]int, index int, value int) [3]int {
    arr[index] = value
    return arr
}

func array_get__Array_2_5int32(arr [2]int32, index int) int32 {
    return arr[index]
}

func array_set__Array_2_5int32(arr [2]int32, index int, value int32) [2]int32 {
    arr[index] = value
    return arr
}

func array_get__Array_2_14Array_2_5int32(arr [2][2]int32, index int) [2]int32 {
    return arr[index]
}

func array_set__Array_2_14Array_2_5int32(arr [2][2]int32, index int, value [2]int32) [2][2]int32 {
    arr[index] = value
    return arr
}

type _goml_vec_int struct {
    items []int
}

func vec_new__Vec_3int() *_goml_vec_int {
    return &_goml_vec_int{
        items: nil,
    }
}

func vec_with_capacity__Vec_3int(capacity int) *_goml_vec_int {
    return &_goml_vec_int{
        items: make([]int, 0, capacity),
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

type _goml_vec_int32 struct {
    items []int32
}

func vec_new__Vec_5int32() *_goml_vec_int32 {
    return &_goml_vec_int32{
        items: nil,
    }
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

type _goml_vec_Array_2_5int32 struct {
    items [][2]int32
}

func vec_new__Vec_14Array_2_5int32() *_goml_vec_Array_2_5int32 {
    return &_goml_vec_Array_2_5int32{
        items: nil,
    }
}

func vec_with_capacity__Vec_14Array_2_5int32(capacity int) *_goml_vec_Array_2_5int32 {
    return &_goml_vec_Array_2_5int32{
        items: make([][2]int32, 0, capacity),
    }
}

func vec_push__Vec_14Array_2_5int32(vec *_goml_vec_Array_2_5int32, elem [2]int32) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_get__Vec_14Array_2_5int32(vec *_goml_vec_Array_2_5int32, index int) [2]int32 {
    return vec.items[index]
}

func vec_set__Vec_14Array_2_5int32(vec *_goml_vec_Array_2_5int32, index int, value [2]int32) struct{} {
    vec.items[index] = value
    return struct{}{}
}

func vec_len__Vec_14Array_2_5int32(vec *_goml_vec_Array_2_5int32) int {
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

type ref_Array_2_5int32_x struct {
    value [2]int32
}

func ref__Ref_14Array_2_5int32(value [2]int32) *ref_Array_2_5int32_x {
    return &ref_Array_2_5int32_x{
        value: value,
    }
}

func ref_get__Ref_14Array_2_5int32(reference *ref_Array_2_5int32_x) [2]int32 {
    return reference.value
}

func ref_set__Ref_14Array_2_5int32(reference *ref_Array_2_5int32_x, value [2]int32) struct{} {
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

func hashmap_lookup__HashMap_6string_5int32(m *hashmap_string_int32_x, key string) (int32, bool, int, uint64) {
    if m == nil {
        var zero int32
        return zero, false, -1, 0
    }
    var index int
    var found bool
    index, found = m.indices[key]
    if !found {
        var zero int32
        return zero, false, -1, 0
    }
    var entry hashmap_string_int32_x_entry = m.entries[index]
    if entry.active {
        return entry.value, true, index, 0
    }
    var zero int32
    return zero, false, index, 0
}

func hashmap_get__HashMap_6string_5int32(m *hashmap_string_int32_x, key string) Option__i32 {
    var value int32
    var ok bool
    value, ok, _, _ = hashmap_lookup__HashMap_6string_5int32(m, key)
    if ok {
        return Option__i32(uint64(int64(value) + 2147483648) + 1)
    }
    return Option__i32(0)
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

type Tuple2_11Array2_3int_3int struct {
    _0 [2]int
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

type Holder struct {
    data [2]int32
    vecs *_goml_vec_Array_2_5int32
}

type Ordering uint8

type Option__i32 uint64

func main0() struct{} {
    var t0 [2]int = [2]int{31, 32}
    var t1 int = array_get__Array_2_3int(t0, 1)
    println__T_isize(t1)
    var arr__0 [3]int = [3]int{1, 2, 3}
    var t2 int = array_get__Array_3_3int(arr__0, 0)
    println__T_isize(t2)
    var arr2__0 [3]int = [3]int{4, 5, 6}
    var place_root0 [3]int = arr2__0
    var index0 int = 1
    array_get__Array_3_3int(place_root0, index0)
    var value0 int = 50
    var t3 [3]int = array_set__Array_3_3int(place_root0, index0, value0)
    arr2__0 = t3
    var t5 int = array_get__Array_3_3int(arr2__0, 1)
    println__T_isize(t5)
    var t6 *_goml_vec_int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_new____T__isize()
    var t7 *_goml_vec_int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pushed____T__isize(t6, 7)
    var t8 *_goml_vec_int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pushed____T__isize(t7, 8)
    var vec__0 *_goml_vec_int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pushed____T__isize(t8, 9)
    var t9 int = vec_get__Vec_3int(vec__0, 2)
    println__T_isize(t9)
    var t10 *_goml_vec_int32
    var inline27 *_goml_vec_int32 = vec_new__Vec_5int32()
    t10 = inline27
    var t11 *_goml_vec_int32 = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pushed____T__i32(t10, 10)
    var t12 *_goml_vec_int32 = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pushed____T__i32(t11, 11)
    var vec2__0 *_goml_vec_int32 = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pushed____T__i32(t12, 12)
    var index1 int = 0
    vec_get__Vec_5int32(vec2__0, index1)
    var value1 int32 = 100
    vec_set__Vec_5int32(vec2__0, index1, value1)
    var t14 int32 = vec_get__Vec_5int32(vec2__0, 0)
    var inline25 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(t14)
    _goml_runtime_core_string_println(inline25)
    var s__0 []int32
    var inline22 int = 0
    var inline23 int = 2
    var inline24 []int32 = vec2__0.items[inline22:inline23]
    s__0 = inline24
    var t15 int32 = s__0[1]
    var inline20 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(t15)
    _goml_runtime_core_string_println(inline20)
    var map__0 *hashmap_string_int32_x
    var inline19 *hashmap_string_int32_x = hashmap_new__HashMap_6string_5int32()
    map__0 = inline19
    var index2 string = "a"
    hashmap_get__HashMap_6string_5int32(map__0, index2)
    var value2 int32 = 13
    hashmap_set__HashMap_6string_5int32(map__0, index2, value2)
    var t17 Option__i32 = hashmap_get__HashMap_6string_5int32(map__0, "a")
    switch t17 != Option__i32(0) {
    case false:
        println__T_string("none")
    case true:
        var inline17 int32 = int32(int64(uint64(t17) - 1) - 2147483648)
        println__T_i32(inline17)
    default:
        panic("non-exhaustive match")
    }
    var t18 Option__i32 = hashmap_get__HashMap_6string_5int32(map__0, "missing")
    switch t18 != Option__i32(0) {
    case false:
        println__T_string("none")
    case true:
        var inline14 int32 = int32(int64(uint64(t18) - 1) - 2147483648)
        println__T_i32(inline14)
    default:
        panic("non-exhaustive match")
    }
    var t19 [2]int32 = [2]int32{1, 2}
    var t20 [2]int32 = [2]int32{3, 4}
    var matrix__0 [2][2]int32 = [2][2]int32{t19, t20}
    var place_root1 [2][2]int32 = matrix__0
    var index3 int = 1
    var place3 [2]int32 = array_get__Array_2_14Array_2_5int32(place_root1, index3)
    var index4 int = 0
    array_get__Array_2_5int32(place3, index4)
    var value3 int32 = 30
    var t21 [2]int32 = array_set__Array_2_5int32(place3, index4, value3)
    var t22 [2][2]int32 = array_set__Array_2_14Array_2_5int32(place_root1, index3, t21)
    matrix__0 = t22
    var t24 [2]int32 = array_get__Array_2_14Array_2_5int32(matrix__0, 1)
    var t25 int32 = array_get__Array_2_5int32(t24, 0)
    var inline11 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(t25)
    _goml_runtime_core_string_println(inline11)
    var t26 [2]int = [2]int{14, 15}
    var pair__0 Tuple2_11Array2_3int_3int = Tuple2_11Array2_3int_3int{
        _0: t26,
        _1: 0,
    }
    var place_root2 Tuple2_11Array2_3int_3int = pair__0
    var place5 [2]int = place_root2._0
    var index5 int = 1
    array_get__Array_2_3int(place5, index5)
    var value4 int = 150
    var t27 [2]int = array_set__Array_2_3int(place5, index5, value4)
    var t28 int = place_root2._1
    var t29 Tuple2_11Array2_3int_3int = Tuple2_11Array2_3int_3int{
        _0: t27,
        _1: t28,
    }
    pair__0 = t29
    var t31 [2]int = pair__0._0
    var t32 int = array_get__Array_2_3int(t31, 1)
    var inline9 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t32)
    _goml_runtime_core_string_println(inline9)
    var t33 [2]int32 = [2]int32{16, 17}
    var t34 *_goml_vec_Array_2_5int32
    var inline8 *_goml_vec_Array_2_5int32 = vec_new__Vec_14Array_2_5int32()
    t34 = inline8
    var t35 [2]int32 = [2]int32{18, 19}
    var t36 *_goml_vec_Array_2_5int32 = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pushed____T___l_i32_x3b_2_r_(t34, t35)
    var t37 [2]int32 = [2]int32{20, 21}
    var t38 *_goml_vec_Array_2_5int32 = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pushed____T___l_i32_x3b_2_r_(t36, t37)
    var holder__0 Holder = Holder{
        data: t33,
        vecs: t38,
    }
    var place_root3 Holder = holder__0
    var place7 [2]int32 = place_root3.data
    var index6 int = 0
    array_get__Array_2_5int32(place7, index6)
    var value5 int32 = 160
    var t39 [2]int32 = array_set__Array_2_5int32(place7, index6, value5)
    var t40 *_goml_vec_Array_2_5int32 = place_root3.vecs
    var t41 Holder = Holder{
        data: t39,
        vecs: t40,
    }
    holder__0 = t41
    var t43 [2]int32 = holder__0.data
    var t44 int32 = array_get__Array_2_5int32(t43, 0)
    var inline6 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(t44)
    _goml_runtime_core_string_println(inline6)
    var place_root4 Holder = holder__0
    var place9 *_goml_vec_Array_2_5int32 = place_root4.vecs
    var index7 int = 1
    var place10 [2]int32 = vec_get__Vec_14Array_2_5int32(place9, index7)
    var index8 int = 0
    array_get__Array_2_5int32(place10, index8)
    var value6 int32 = 200
    var t45 [2]int32 = array_set__Array_2_5int32(place10, index8, value6)
    vec_set__Vec_14Array_2_5int32(place9, index7, t45)
    var t47 *_goml_vec_Array_2_5int32 = holder__0.vecs
    var t48 [2]int32 = vec_get__Vec_14Array_2_5int32(t47, 1)
    var t49 int32 = array_get__Array_2_5int32(t48, 0)
    var inline4 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(t49)
    _goml_runtime_core_string_println(inline4)
    var t50 [2]int32 = [2]int32{22, 23}
    var r__0 *ref_Array_2_5int32_x
    var inline3 *ref_Array_2_5int32_x = ref__Ref_14Array_2_5int32(t50)
    r__0 = inline3
    var place_root5 [2]int32 = ref_get__Ref_14Array_2_5int32(r__0)
    var index9 int = 1
    array_get__Array_2_5int32(place_root5, index9)
    var value7 int32 = 230
    var t51 [2]int32 = array_set__Array_2_5int32(place_root5, index9, value7)
    ref_set__Ref_14Array_2_5int32(r__0, t51)
    var t53 [2]int32
    var inline2 [2]int32 = ref_get__Ref_14Array_2_5int32(r__0)
    t53 = inline2
    var t54 int32 = array_get__Array_2_5int32(t53, 1)
    var inline0 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(t54)
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

func println__T_isize(value__0 int) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_int_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_new____T__isize() *_goml_vec_int {
    var t0 *_goml_vec_int = vec_new__Vec_3int()
    return t0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pushed____T__isize(self__0 *_goml_vec_int, elem__0 int) *_goml_vec_int {
    var t0 int
    var inline4 int = vec_len__Vec_3int(self__0)
    t0 = inline4
    var t1 int = t0 + 1
    var result__0 *_goml_vec_int
    var inline3 *_goml_vec_int = vec_with_capacity__Vec_3int(t1)
    result__0 = inline3
    var index__0 int = 0
    Loop_loop0:
    for {
        var t2 int
        var inline2 int = vec_len__Vec_3int(self__0)
        t2 = inline2
        var t3 bool = index__0 < t2
        if t3 {
            var t4 int = vec_get__Vec_3int(self__0, index__0)
            vec_push__Vec_3int(result__0, t4)
            var compound_old0 int = index__0
            var compound_value0 int = 1
            var t5 int = compound_old0 + compound_value0
            index__0 = t5
            continue
        } else {
            break Loop_loop0
        }
    }
    vec_push__Vec_3int(result__0, elem__0)
    return result__0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pushed____T__i32(self__0 *_goml_vec_int32, elem__0 int32) *_goml_vec_int32 {
    var t0 int
    var inline4 int = vec_len__Vec_5int32(self__0)
    t0 = inline4
    var t1 int = t0 + 1
    var result__0 *_goml_vec_int32
    var inline3 *_goml_vec_int32 = vec_with_capacity__Vec_5int32(t1)
    result__0 = inline3
    var index__0 int = 0
    Loop_loop0:
    for {
        var t2 int
        var inline2 int = vec_len__Vec_5int32(self__0)
        t2 = inline2
        var t3 bool = index__0 < t2
        if t3 {
            var t4 int32 = vec_get__Vec_5int32(self__0, index__0)
            vec_push__Vec_5int32(result__0, t4)
            var compound_old0 int = index__0
            var compound_value0 int = 1
            var t5 int = compound_old0 + compound_value0
            index__0 = t5
            continue
        } else {
            break Loop_loop0
        }
    }
    vec_push__Vec_5int32(result__0, elem__0)
    return result__0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pushed____T___l_i32_x3b_2_r_(self__0 *_goml_vec_Array_2_5int32, elem__0 [2]int32) *_goml_vec_Array_2_5int32 {
    var t0 int
    var inline4 int = vec_len__Vec_14Array_2_5int32(self__0)
    t0 = inline4
    var t1 int = t0 + 1
    var result__0 *_goml_vec_Array_2_5int32
    var inline3 *_goml_vec_Array_2_5int32 = vec_with_capacity__Vec_14Array_2_5int32(t1)
    result__0 = inline3
    var index__0 int = 0
    Loop_loop0:
    for {
        var t2 int
        var inline2 int = vec_len__Vec_14Array_2_5int32(self__0)
        t2 = inline2
        var t3 bool = index__0 < t2
        if t3 {
            var t4 [2]int32 = vec_get__Vec_14Array_2_5int32(self__0, index__0)
            vec_push__Vec_14Array_2_5int32(result__0, t4)
            var compound_old0 int = index__0
            var compound_value0 int = 1
            var t5 int = compound_old0 + compound_value0
            index__0 = t5
            continue
        } else {
            break Loop_loop0
        }
    }
    vec_push__Vec_14Array_2_5int32(result__0, elem__0)
    return result__0
}

func _goml_m_trait__impl_i_ToString_i_i32_i_to__string(self__0 int32) string {
    var inline0 int64 = int64(int32(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
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
