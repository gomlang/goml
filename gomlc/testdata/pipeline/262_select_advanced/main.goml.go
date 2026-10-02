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

func main0() struct{} {
    var log__0 *ref_string_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__string("")
    var disabled__0 chan int = _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_new____T__isize(1)
    _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_close____T__isize(disabled__0)
    var t0 chan int
    var inline35 string = "c"
    var inline36 string = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__string(log__0)
    var inline37 string = inline36 + inline35
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__string(log__0, inline37)
    t0 = disabled__0
    var t1 int
    var inline30 string = "v"
    var inline31 int = 1
    var inline32 string = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__string(log__0)
    var inline33 string = inline32 + inline30
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__string(log__0, inline33)
    t1 = inline31
    var t2 bool
    var inline25 string = "g"
    var inline26 bool = false
    var inline27 string = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__string(log__0)
    var inline28 string = inline27 + inline25
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__string(log__0, inline28)
    t2 = inline26
    var jp0 string
    var select_channel_0_0 chan int
    if t2 {
        select_channel_0_0 = t0
    }
    select {
    case select_channel_0_0 <- t1:
        jp0 = "sent"
    default:
        jp0 = "default"
    }
    var t3 string
    var inline24 string = ref_get__Ref_6string(log__0)
    t3 = inline24
    var inline22 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t3)
    _goml_runtime_core_string_println(inline22)
    var inline20 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(jp0)
    _goml_runtime_core_string_println(inline20)
    var first__0 chan int
    var inline18 int = 1
    var inline19 chan int = func(p0 int) chan int {
        return make(chan int, p0)
    }(inline18)
    first__0 = inline19
    var second__0 chan int
    var inline16 int = 1
    var inline17 chan int = func(p0 int) chan int {
        return make(chan int, p0)
    }(inline16)
    second__0 = inline17
    var inline14 int = 10
    func(p0 chan int, p1 int) struct{} {
        p0 <- p1
        return struct{}{}
    }(first__0, inline14)
    var inline12 int = 20
    func(p0 chan int, p1 int) struct{} {
        p0 <- p1
        return struct{}{}
    }(second__0, inline12)
    var jp1 int
    var _goml_m_value____0_i_select__value int
    var _goml_m_value____0_i_select__open bool
    var value__0 Option__isize = Option__isize{
        _tag: 0,
    }
    var select_channel_1_0 chan int
    if true {
        select_channel_1_0 = first__0
    }
    var _goml_m_value____1_i_select__value int
    var _goml_m_value____1_i_select__open bool
    var value__1 Option__isize = Option__isize{
        _tag: 0,
    }
    var select_channel_1_1 chan int
    if true {
        select_channel_1_1 = second__0
    }
    select {
    case _goml_m_value____0_i_select__value, _goml_m_value____0_i_select__open = <-select_channel_1_0:
        if _goml_m_value____0_i_select__open {
            value__0 = Option__isize{
                _p0: _goml_m_value____0_i_select__value,
                _tag: 1,
            }
        }
        var inline8 int = -1
        switch value__0._tag {
        case 0:
            jp1 = inline8
        case 1:
            var inline9 int = value__0._p0
            jp1 = inline9
        default:
            panic("non-exhaustive match")
        }
    default:
        select {
        case _goml_m_value____1_i_select__value, _goml_m_value____1_i_select__open = <-select_channel_1_1:
            if _goml_m_value____1_i_select__open {
                value__1 = Option__isize{
                    _p0: _goml_m_value____1_i_select__value,
                    _tag: 1,
                }
            }
            var inline10 int = -1
            switch value__1._tag {
            case 0:
                jp1 = inline10
            case 1:
                var inline11 int = value__1._p0
                jp1 = inline11
            default:
                panic("non-exhaustive match")
            }
        default:
            jp1 = 0
        }
    }
    var inline6 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(jp1)
    _goml_runtime_core_string_println(inline6)
    var events__0 chan int
    var inline4 int = 1
    var inline5 chan int = func(p0 int) chan int {
        return make(chan int, p0)
    }(inline4)
    events__0 = inline5
    var inline2 int = 7
    func(p0 chan int, p1 int) struct{} {
        p0 <- p1
        return struct{}{}
    }(events__0, inline2)
    var t4_lhs int = 1
    var t4_rhs int = 1
    var t4 int = t4_lhs + t4_rhs
    var t5 bool = t4 == 2
    var jp2 int
    var _goml_m__d_select__recv__1184____0_i_select__value int
    var _goml_m__d_select__recv__1184____0_i_select__open bool
    var _goml_m__d_select__recv__1184____0 Option__isize = Option__isize{
        _tag: 0,
    }
    var select_channel_2_0 chan int
    if t5 {
        select_channel_2_0 = events__0
    }
    select {
    case _goml_m__d_select__recv__1184____0_i_select__value, _goml_m__d_select__recv__1184____0_i_select__open = <-select_channel_2_0:
        if _goml_m__d_select__recv__1184____0_i_select__open {
            _goml_m__d_select__recv__1184____0 = Option__isize{
                _p0: _goml_m__d_select__recv__1184____0_i_select__value,
                _tag: 1,
            }
        }
        switch _goml_m__d_select__recv__1184____0._tag {
        case 0:
            jp2 = 0
        case 1:
            var x0 int = _goml_m__d_select__recv__1184____0._p0
            var t6 int = x0 + 1
            jp2 = t6
        default:
            panic("non-exhaustive match")
        }
    default:
        jp2 = -1
    }
    var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(jp2)
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

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__string(value__0 string) *ref_string_x {
    var t0 *ref_string_x = ref__Ref_6string(value__0)
    return t0
}

func _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_new____T__isize(capacity__0 int) chan int {
    var t0 chan int = func(p0 int) chan int {
        return make(chan int, p0)
    }(capacity__0)
    return t0
}

func _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_close____T__isize(self__0 chan int) struct{} {
    func(p0 chan int) struct{} {
        close(p0)
        return struct{}{}
    }(self__0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
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
