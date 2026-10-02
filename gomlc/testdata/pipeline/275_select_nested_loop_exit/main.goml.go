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

func default_break_outer_continue() int {
    var rounds__0 *ref_int_x
    var inline5 int = 0
    var inline6 *ref_int_x = ref__Ref_3int(inline5)
    rounds__0 = inline6
    var channel__0 chan int
    var inline3 int = 0
    var inline4 chan int = func(p0 int) chan int {
        return make(chan int, p0)
    }(inline3)
    channel__0 = inline4
    var jp0 int
    Loop_loop_expr0:
    for {
        var t0 int
        var inline2 int = ref_get__Ref_3int(rounds__0)
        t0 = inline2
        var t1 bool = t0 == 0
        if t1 {
            var inline0 int = 1
            ref_set__Ref_3int(rounds__0, inline0)
            Loop_loop_expr1:
            for {
                var _goml_m_______0_i_select__open bool
                select {
                case _, _goml_m_______0_i_select__open = <-channel__0:
                    if _goml_m_______0_i_select__open {}
                    continue
                default:
                    break Loop_loop_expr1
                }
            }
            continue
        } else {
            jp0 = 11
            break Loop_loop_expr0
        }
    }
    return jp0
}

func default_break_outer_break() int {
    var channel__0 chan int
    var inline0 int = 0
    var inline1 chan int = func(p0 int) chan int {
        return make(chan int, p0)
    }(inline0)
    channel__0 = inline1
    var jp0 int
    Loop_loop_expr1:
    for {
        var _goml_m_______0_i_select__open bool
        select {
        case _, _goml_m_______0_i_select__open = <-channel__0:
            if _goml_m_______0_i_select__open {}
            continue
        default:
            break Loop_loop_expr1
        }
    }
    jp0 = 12
    return jp0
}

func arm_break_outer_continue() int {
    var rounds__0 *ref_int_x
    var inline7 int = 0
    var inline8 *ref_int_x = ref__Ref_3int(inline7)
    rounds__0 = inline8
    var channel__0 chan int
    var inline5 int = 1
    var inline6 chan int = func(p0 int) chan int {
        return make(chan int, p0)
    }(inline5)
    channel__0 = inline6
    var inline3 int = 1
    func(p0 chan int, p1 int) struct{} {
        p0 <- p1
        return struct{}{}
    }(channel__0, inline3)
    var jp0 int
    Loop_loop_expr0:
    for {
        var t0 int
        var inline2 int = ref_get__Ref_3int(rounds__0)
        t0 = inline2
        var t1 bool = t0 == 0
        if t1 {
            var inline0 int = 1
            ref_set__Ref_3int(rounds__0, inline0)
            Loop_loop_expr1:
            for {
                var _goml_m_______0_i_select__open bool
                select {
                case _, _goml_m_______0_i_select__open = <-channel__0:
                    if _goml_m_______0_i_select__open {}
                    break Loop_loop_expr1
                default:
                    continue
                }
            }
            continue
        } else {
            jp0 = 13
            break Loop_loop_expr0
        }
    }
    return jp0
}

func arm_break_outer_break() int {
    var channel__0 chan int
    var inline2 int = 1
    var inline3 chan int = func(p0 int) chan int {
        return make(chan int, p0)
    }(inline2)
    channel__0 = inline3
    var inline0 int = 1
    func(p0 chan int, p1 int) struct{} {
        p0 <- p1
        return struct{}{}
    }(channel__0, inline0)
    var jp0 int
    Loop_loop_expr1:
    for {
        var _goml_m_______0_i_select__open bool
        select {
        case _, _goml_m_______0_i_select__open = <-channel__0:
            if _goml_m_______0_i_select__open {}
            break Loop_loop_expr1
        default:
            continue
        }
    }
    jp0 = 14
    return jp0
}

func main0() struct{} {
    var t0 int = default_break_outer_continue()
    var inline6 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t0)
    _goml_runtime_core_string_println(inline6)
    var t1 int = default_break_outer_break()
    var inline4 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t1)
    _goml_runtime_core_string_println(inline4)
    var t2 int = arm_break_outer_continue()
    var inline2 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t2)
    _goml_runtime_core_string_println(inline2)
    var t3 int = arm_break_outer_break()
    var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t3)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
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
