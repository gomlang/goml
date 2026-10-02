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

func array_get__Array_3_5uint8(arr [3]uint8, index int) uint8 {
    return arr[index]
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

func answer() int {
    var base__0 int = 40
    var t0 int = base__0 + 2
    return t0
}

func loop_answer() int {
    var jp0 int
    var base__0 int = 6
    var t0 int = base__0 * 7
    jp0 = t0
    return jp0
}

func main0() struct{} {
    var plain__0 *_goml_vec_uint8 = &_goml_vec_uint8{
        items: []uint8{65, 10, 66},
    }
    var empty__0 *_goml_vec_uint8 = &_goml_vec_uint8{
        items: []uint8{},
    }
    var raw__0 *_goml_vec_uint8 = &_goml_vec_uint8{
        items: []uint8{114, 97, 119, 32, 92, 110, 32, 98, 121, 116, 101, 115},
    }
    var quoted__0 *_goml_vec_uint8 = &_goml_vec_uint8{
        items: []uint8{113, 117, 111, 116, 101, 100, 32, 34, 116, 101, 120, 116, 34, 32, 97, 110, 100, 32, 35},
    }
    var fixed__0 [3]uint8 = [3]uint8{65, 66, 67}
    var value__0 int = answer()
    var t0 int = loop_answer()
    var t1 int = value__0 + t0
    var t2 int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_len____T__u8(plain__0)
    var t3 string = _goml_m_inherent_i_isize_i_isize_i_to__string(t2)
    println__T_string(t3)
    var t4 int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_len____T__u8(empty__0)
    var t5 string = _goml_m_inherent_i_isize_i_isize_i_to__string(t4)
    println__T_string(t5)
    var t6 uint8
    var inline30 int = 0
    var inline31 uint8 = vec_get__Vec_5uint8(plain__0, inline30)
    t6 = inline31
    var t7 int = int(uint8(t6))
    var t8 string
    var inline29 string = __goml_builtin_int_to_string(t7)
    t8 = inline29
    var inline27 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t8)
    _goml_runtime_core_string_println(inline27)
    var t9 uint8
    var inline25 int = 1
    var inline26 uint8 = vec_get__Vec_5uint8(plain__0, inline25)
    t9 = inline26
    var t10 int = int(uint8(t9))
    var t11 string
    var inline24 string = __goml_builtin_int_to_string(t10)
    t11 = inline24
    var inline22 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t11)
    _goml_runtime_core_string_println(inline22)
    var t12 uint8
    var inline20 int = 2
    var inline21 uint8 = vec_get__Vec_5uint8(plain__0, inline20)
    t12 = inline21
    var t13 int = int(uint8(t12))
    var t14 string
    var inline19 string = __goml_builtin_int_to_string(t13)
    t14 = inline19
    var inline17 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t14)
    _goml_runtime_core_string_println(inline17)
    var t15 int
    var inline16 int = vec_len__Vec_5uint8(raw__0)
    t15 = inline16
    var t16 string
    var inline15 string = __goml_builtin_int_to_string(t15)
    t16 = inline15
    var inline13 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t16)
    _goml_runtime_core_string_println(inline13)
    var t17 int
    var inline12 int = vec_len__Vec_5uint8(quoted__0)
    t17 = inline12
    var t18 string
    var inline11 string = __goml_builtin_int_to_string(t17)
    t18 = inline11
    var inline9 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t18)
    _goml_runtime_core_string_println(inline9)
    var t19 uint8 = array_get__Array_3_5uint8(fixed__0, 2)
    var t20 int = int(uint8(t19))
    var t21 string
    var inline8 string = __goml_builtin_int_to_string(t20)
    t21 = inline8
    var inline6 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t21)
    _goml_runtime_core_string_println(inline6)
    var t22 string
    var inline5 string = __goml_builtin_int_to_string(t1)
    t22 = inline5
    var inline3 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t22)
    _goml_runtime_core_string_println(inline3)
    var inline0 string = "block condition"
    var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
    _goml_runtime_core_string_println(inline1)
    return struct{}{}
}

func println__T_string(value__0 string) struct{} {
    var t0 string
    t0 = value__0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_len____T__u8(self__0 *_goml_vec_uint8) int {
    var t0 int = vec_len__Vec_5uint8(self__0)
    return t0
}

func _goml_m_inherent_i_isize_i_isize_i_to__string(self__0 int) string {
    var inline0 int64 = int64(int(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
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
