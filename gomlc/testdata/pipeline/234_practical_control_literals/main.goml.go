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

type Option__isize struct {
    _p0 int
    _tag uint8
}

func unwrap_or_negative(value__0 Option__isize) int {
    switch value__0._tag {
    case 1:
        var x0 int = value__0._p0
        return x0
    default:
        return -1
    }
}

func count_to(limit__0 int) int {
    var counter__0 *ref_int_x
    var inline2 int = 0
    var inline3 *ref_int_x = ref__Ref_3int(inline2)
    counter__0 = inline3
    var jp0 int
    Loop_loop_expr0:
    for {
        var current__0 int
        var inline1 int = ref_get__Ref_3int(counter__0)
        current__0 = inline1
        var t0 bool = current__0 >= limit__0
        if t0 {
            jp0 = current__0
            break Loop_loop_expr0
        } else {
            var t1 int = current__0 + 1
            ref_set__Ref_3int(counter__0, t1)
            continue
        }
    }
    return jp0
}

func loop_option(value__0 Option__isize) int {
    var jp0 int
    switch value__0._tag {
    case 1:
        var x0 int = value__0._p0
        jp0 = x0
        return jp0
    default:
        jp0 = -2
        return jp0
    }
}

func nested_loop_value() int {
    var jp0 int
    jp0 = 7
    return jp0
}

func main0() struct{} {
    println__T_string("C:\\tmp\\\"quoted\\\"")
    var t0 string = "" + "}"
    var inline22 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t0)
    _goml_runtime_core_string_println(inline22)
    var t1 Option__isize = Option__isize{
        _p0: 11,
        _tag: 1,
    }
    var t2 int = unwrap_or_negative(t1)
    var t3 string
    var inline21 string = __goml_builtin_int_to_string(t2)
    t3 = inline21
    var inline19 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t3)
    _goml_runtime_core_string_println(inline19)
    var t4 int
    t4 = -1
    var t5 string
    var inline18 string = __goml_builtin_int_to_string(t4)
    t5 = inline18
    var inline16 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t5)
    _goml_runtime_core_string_println(inline16)
    var t6 int = count_to(4)
    var t7 string
    var inline15 string = __goml_builtin_int_to_string(t6)
    t7 = inline15
    var inline13 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t7)
    _goml_runtime_core_string_println(inline13)
    var t8 Option__isize = Option__isize{
        _p0: 9,
        _tag: 1,
    }
    var t9 int = loop_option(t8)
    var t10 string
    var inline12 string = __goml_builtin_int_to_string(t9)
    t10 = inline12
    var inline10 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t10)
    _goml_runtime_core_string_println(inline10)
    var t11 int = loop_option(Option__isize{
        _tag: 0,
    })
    var t12 string
    var inline9 string = __goml_builtin_int_to_string(t11)
    t12 = inline9
    var inline7 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t12)
    _goml_runtime_core_string_println(inline7)
    var t13 int = nested_loop_value()
    var t14 string
    var inline6 string = __goml_builtin_int_to_string(t13)
    t14 = inline6
    var inline4 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t14)
    _goml_runtime_core_string_println(inline4)
    var t15 bool
    var inline3 string = "C:\\tmp"
    switch inline3 {
    case "C:\\tmp":
        t15 = true
    default:
        t15 = false
    }
    var t16 string
    var inline2 string = _goml_runtime_core_bool_to_string(t15)
    t16 = inline2
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t16)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func println__T_string(value__0 string) struct{} {
    var t0 string
    t0 = value__0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
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
