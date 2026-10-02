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

type _goml_vec_int struct {
    items []int
}

func vec_new__Vec_3int() *_goml_vec_int {
    return &_goml_vec_int{
        items: nil,
    }
}

func vec_get__Vec_3int(vec *_goml_vec_int, index int) int {
    return vec.items[index]
}

func vec_set__Vec_3int(vec *_goml_vec_int, index int, value int) struct{} {
    vec.items[index] = value
    return struct{}{}
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

type ref_string_x struct {
    value string
}

func ref__Ref_6string(value string) *ref_string_x {
    return &ref_string_x{
        value: value,
    }
}

func ref_get__Ref_6string(reference *ref_string_x) string {
    return reference.value
}

func ref_set__Ref_6string(reference *ref_string_x, value string) struct{} {
    reference.value = value
    return struct{}{}
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

func hashmap_len__HashMap_6string_3int(m *hashmap_string_int_x) int {
    if m == nil {
        return 0
    }
    return m.len
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

func hashmap_set__HashMap_6string_3int(m *hashmap_string_int_x, key string, value int) struct{} {
    if m == nil {
        return struct{}{}
    }
    var index int
    var found bool
    index, found = m.indices[key]
    if found {
        var entry hashmap_string_int_x_entry = m.entries[index]
        if entry.active {
            m.entries[index].value = value
            return struct{}{}
        }
        m.entries[index] = hashmap_string_int_x_entry{
            active: true,
            key: key,
            value: value,
        }
        m.len = m.len + 1
        return struct{}{}
    }
    index = len(m.entries)
    m.indices[key] = index
    m.entries = append(m.entries, hashmap_string_int_x_entry{
        active: true,
        key: key,
        value: value,
    })
    m.len = m.len + 1
    return struct{}{}
}

type Tuple2_3int_3int struct {
    _0 int
    _1 int
}

type Tuple2_6string_3int struct {
    _0 string
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

type Point struct {
    x int
    y int
}

type Ordering uint8

type Option__isize struct {
    _p0 int
    _tag uint8
}

func record(log__0 *ref_string_x, label__0 string, value__0 int) int {
    var t0 string
    var inline1 string = ref_get__Ref_6string(log__0)
    t0 = inline1
    var t1 string = t0 + label__0
    ref_set__Ref_6string(log__0, t1)
    return value__0
}

func record_point(log__0 *ref_string_x, label__0 string, value__0 Point) Point {
    var t0 string
    var inline1 string = ref_get__Ref_6string(log__0)
    t0 = inline1
    var t1 string = t0 + label__0
    ref_set__Ref_6string(log__0, t1)
    return value__0
}

func record_vec(log__0 *ref_string_x, label__0 string, value__0 *_goml_vec_int) *_goml_vec_int {
    var t0 string
    var inline1 string = ref_get__Ref_6string(log__0)
    t0 = inline1
    var t1 string = t0 + label__0
    ref_set__Ref_6string(log__0, t1)
    return value__0
}

func main0() struct{} {
    var number__0 int = 5
    var compound_old0 int = number__0
    var compound_value0 int = 3
    var t0 int = compound_old0 + compound_value0
    number__0 = t0
    var compound_old1 int = number__0
    var compound_value1 int = 2
    var t2 int = compound_old1 * compound_value1
    number__0 = t2
    var compound_old2 int = number__0
    var compound_value2 int = 1
    var t4 int = compound_old2 >> compound_value2
    number__0 = t4
    var t6 string = _goml_m_inherent_i_isize_i_isize_i_to__string(number__0)
    println__T_string(t6)
    var direct__0 Point = Point{
        x: 3,
        y: 4,
    }
    var place_root0 Point = direct__0
    var place0 int = place_root0.x
    var value0 int = 5
    var t7 int = place0 + value0
    var t8 int = place_root0.y
    var t9 Point = Point{
        x: t7,
        y: t8,
    }
    direct__0 = t9
    var t11 int = direct__0.x
    var t12 string = _goml_m_inherent_i_isize_i_isize_i_to__string(t11)
    var t13 string = "" + t12
    var t14 string = t13 + ","
    var t15 int = direct__0.y
    var t16 string = _goml_m_inherent_i_isize_i_isize_i_to__string(t15)
    var t17 string = t14 + t16
    println__T_string(t17)
    var pair__0 Tuple2_3int_3int = Tuple2_3int_3int{
        _0: 2,
        _1: 9,
    }
    var place_root1 Tuple2_3int_3int = pair__0
    var place1 int = place_root1._0
    var value1 int = 3
    var t18 int = place1 * value1
    var t19 int = place_root1._1
    var t20 Tuple2_3int_3int = Tuple2_3int_3int{
        _0: t18,
        _1: t19,
    }
    pair__0 = t20
    var t22 int = pair__0._0
    var t23 string = _goml_m_inherent_i_isize_i_isize_i_to__string(t22)
    var t24 string = "" + t23
    var t25 string = t24 + ","
    var t26 int = pair__0._1
    var t27 string = _goml_m_inherent_i_isize_i_isize_i_to__string(t26)
    var t28 string = t25 + t27
    println__T_string(t28)
    var log__0 *ref_string_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__string("")
    var base__0 Point = Point{
        x: 1,
        y: 2,
    }
    var struct_update_field__789__0 int = record(log__0, "F", 7)
    var struct_update_base__0 Point = record_point(log__0, "B", base__0)
    var t29 int = struct_update_base__0.y
    var t30 string = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__string(log__0)
    println__T_string(t30)
    var t31 string = _goml_m_inherent_i_isize_i_isize_i_to__string(struct_update_field__789__0)
    var t32 string = "" + t31
    var t33 string = t32 + ","
    var t34 string = _goml_m_inherent_i_isize_i_isize_i_to__string(t29)
    var t35 string = t33 + t34
    println__T_string(t35)
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__string(log__0, "")
    var t36 int = record(log__0, "A", 10)
    var t37 int = record(log__0, "B", 20)
    var t38 [2]int = [2]int{t36, t37}
    var values__0 *_goml_vec_int = func(values [2]int) *_goml_vec_int {
        var storage struct {
            vector _goml_vec_int
            values [2]int
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t38)
    var place_root2 *_goml_vec_int = record_vec(log__0, "R", values__0)
    var index0 int = record(log__0, "I", 1)
    var place2 int = vec_get__Vec_3int(place_root2, index0)
    var value2 int = record(log__0, "V", 5)
    var t39 int = place2 + value2
    vec_set__Vec_3int(place_root2, index0, t39)
    var t41 string = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__string(log__0)
    println__T_string(t41)
    var t42 int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_get____T__isize(values__0, 0)
    var t43 string = _goml_m_inherent_i_isize_i_isize_i_to__string(t42)
    var t44 string = "" + t43
    var t45 string = t44 + ","
    var t46 int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_get____T__isize(values__0, 1)
    var t47 string = _goml_m_inherent_i_isize_i_isize_i_to__string(t46)
    var t48 string = t45 + t47
    var inline38 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t48)
    _goml_runtime_core_string_println(inline38)
    var inline36 string = ""
    ref_set__Ref_6string(log__0, inline36)
    var t49 string = "" + "k"
    var t50 int
    var inline31 string = "K"
    var inline32 int = 1
    var inline33 string = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__string(log__0)
    var inline34 string = inline33 + inline31
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__string(log__0, inline34)
    t50 = inline32
    var t51 string
    var inline30 string = __goml_builtin_int_to_string(t50)
    t51 = inline30
    var t52 string = t49 + t51
    var t53 int
    var inline25 string = "V"
    var inline26 int = 11
    var inline27 string = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__string(log__0)
    var inline28 string = inline27 + inline25
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__string(log__0, inline28)
    t53 = inline26
    var t54 Tuple2_6string_3int = Tuple2_6string_3int{
        _0: t52,
        _1: t53,
    }
    var t55 int
    var inline20 string = "A"
    var inline21 int = 1
    var inline22 string = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__string(log__0)
    var inline23 string = inline22 + inline20
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__string(log__0, inline23)
    t55 = inline21
    var t56 Tuple2_6string_3int = Tuple2_6string_3int{
        _0: "same",
        _1: t55,
    }
    var t57 int
    var inline15 string = "B"
    var inline16 int = 2
    var inline17 string = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__string(log__0)
    var inline18 string = inline17 + inline15
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__string(log__0, inline18)
    t57 = inline16
    var t58 Tuple2_6string_3int = Tuple2_6string_3int{
        _0: "same",
        _1: t57,
    }
    var table__0 *hashmap_string_int_x = &hashmap_string_int_x{
        indices: make(map[string]int, 3),
        entries: make([]hashmap_string_int_x_entry, 0, 3),
        len: 0,
    }
    hashmap_set__HashMap_6string_3int(table__0, t54._0, t54._1)
    hashmap_set__HashMap_6string_3int(table__0, t56._0, t56._1)
    hashmap_set__HashMap_6string_3int(table__0, t58._0, t58._1)
    var t60 string
    var inline14 string = ref_get__Ref_6string(log__0)
    t60 = inline14
    var inline12 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t60)
    _goml_runtime_core_string_println(inline12)
    var mtmp0 Option__isize
    var inline10 string = "same"
    var inline11 Option__isize = hashmap_get__HashMap_6string_3int(table__0, inline10)
    mtmp0 = inline11
    var jp0 string
    switch mtmp0._tag {
    case 0:
        jp0 = "missing"
    case 1:
        var x0 int = mtmp0._p0
        var inline9 string = __goml_builtin_int_to_string(x0)
        jp0 = inline9
    default:
        panic("non-exhaustive match")
    }
    var inline7 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(jp0)
    _goml_runtime_core_string_println(inline7)
    var empty_values__0 *_goml_vec_int
    var inline6 *_goml_vec_int = vec_new__Vec_3int()
    empty_values__0 = inline6
    var empty_table__0 *hashmap_string_int_x
    var inline5 *hashmap_string_int_x = hashmap_new__HashMap_6string_3int()
    empty_table__0 = inline5
    var t61 string = "" + "empty="
    var t62 int
    var inline4 int = vec_len__Vec_3int(empty_values__0)
    t62 = inline4
    var t63 int
    var inline3 int = hashmap_len__HashMap_6string_3int(empty_table__0)
    t63 = inline3
    var t64 int = t62 + t63
    var t65 string
    var inline2 string = __goml_builtin_int_to_string(t64)
    t65 = inline2
    var t66 string = t61 + t65
    var t67 string = t66 + " {ok}"
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t67)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__string(self__0 *ref_string_x) string {
    var t0 string = ref_get__Ref_6string(self__0)
    return t0
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__string(self__0 *ref_string_x, value__0 string) struct{} {
    ref_set__Ref_6string(self__0, value__0)
    return struct{}{}
}

func println__T_string(value__0 string) struct{} {
    var t0 string
    t0 = value__0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_inherent_i_isize_i_isize_i_to__string(self__0 int) string {
    var inline0 int64 = int64(int(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__string(value__0 string) *ref_string_x {
    var t0 *ref_string_x = ref__Ref_6string(value__0)
    return t0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_get____T__isize(self__0 *_goml_vec_int, index__0 int) int {
    var t0 int = vec_get__Vec_3int(self__0, index__0)
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
