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

type Tuple2_6string_4bool struct {
    _0 string
    _1 bool
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

type closure_env_read_0 struct {
    value_0 Option__isize
}

type Ordering uint8

type Option__isize struct {
    _p0 int
    _tag uint8
}

type Option__string struct {
    _p0 string
    _tag uint8
}

func main0() struct{} {
    var received__0 chan int = _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_new____T__isize(1)
    _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_send____T__isize(received__0, 7)
    var jp0 int
    var _goml_m_value____1_i_select__value int
    var _goml_m_value____1_i_select__open bool
    var value__1 Option__isize = Option__isize{
        _tag: 0,
    }
    select {
    case _goml_m_value____1_i_select__value, _goml_m_value____1_i_select__open = <-received__0:
        if _goml_m_value____1_i_select__open {
            value__1 = Option__isize{
                _p0: _goml_m_value____1_i_select__value,
                _tag: 1,
            }
        }
        var t7 closure_env_read_0 = closure_env_read_0{
            value_0: value__1,
        }
        var read__0 func() int = func() int {
            return _goml_m_inherent_i_closure__env__read__0_i_closure__env__read__0_i_apply(t7)
        }
        var t8 int = read__0()
        jp0 = t8
    default:
        jp0 = 0
    }
    println__T_isize(jp0)
    var sent__0 chan string = _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_new____T__string(1)
    var jp1 string
    select {
    case sent__0 <- "ready":
        jp1 = "sent"
    default:
        jp1 = "blocked"
    }
    var inline38 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(jp1)
    _goml_runtime_core_string_println(inline38)
    var t0 Option__string = _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_recv____T__string(sent__0)
    var t1 string
    var inline36 string = "missing"
    switch t0._tag {
    case 0:
        t1 = inline36
    case 1:
        var inline37 string = t0._p0
        t1 = inline37
    default:
        panic("non-exhaustive match")
    }
    var inline34 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t1)
    _goml_runtime_core_string_println(inline34)
    var empty__0 chan int
    var inline32 int = 0
    var inline33 chan int = func(p0 int) chan int {
        return make(chan int, p0)
    }(inline32)
    empty__0 = inline33
    var jp2 int
    var _goml_m_______0_i_select__open bool
    select {
    case _, _goml_m_______0_i_select__open = <-empty__0:
        if _goml_m_______0_i_select__open {}
        jp2 = -1
    default:
        jp2 = 42
    }
    var inline30 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(jp2)
    _goml_runtime_core_string_println(inline30)
    var log__0 *ref_string_x
    var inline28 string = ""
    var inline29 *ref_string_x = ref__Ref_6string(inline28)
    log__0 = inline29
    var left__0 chan int
    var inline26 int = 0
    var inline27 chan int = func(p0 int) chan int {
        return make(chan int, p0)
    }(inline26)
    left__0 = inline27
    var right__0 chan int
    var inline24 int = 0
    var inline25 chan int = func(p0 int) chan int {
        return make(chan int, p0)
    }(inline24)
    right__0 = inline25
    var t2 chan int
    var inline20 string = "a"
    var inline21 string = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__string(log__0)
    var inline22 string = inline21 + inline20
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__string(log__0, inline22)
    t2 = left__0
    var t3 int
    var inline15 string = "1"
    var inline16 int = 1
    var inline17 string = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__string(log__0)
    var inline18 string = inline17 + inline15
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__string(log__0, inline18)
    t3 = inline16
    var t4 chan int
    var inline11 string = "b"
    var inline12 string = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__string(log__0)
    var inline13 string = inline12 + inline11
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__string(log__0, inline13)
    t4 = right__0
    var t5 int
    var inline6 string = "2"
    var inline7 int = 2
    var inline8 string = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__string(log__0)
    var inline9 string = inline8 + inline6
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__string(log__0, inline9)
    t5 = inline7
    select {
    case t2 <- t3:
    case t4 <- t5:
    default:
    }
    var t6 string
    var inline5 string = ref_get__Ref_6string(log__0)
    t6 = inline5
    var inline3 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t6)
    _goml_runtime_core_string_println(inline3)
    func(p0 chan int) struct{} {
        close(p0)
        return struct{}{}
    }(received__0)
    var jp3 string
    var _goml_m_value____0_i_select__value int
    var _goml_m_value____0_i_select__open bool
    var value__0 Option__isize = Option__isize{
        _tag: 0,
    }
    select {
    case _goml_m_value____0_i_select__value, _goml_m_value____0_i_select__open = <-received__0:
        if _goml_m_value____0_i_select__open {
            value__0 = Option__isize{
                _p0: _goml_m_value____0_i_select__value,
                _tag: 1,
            }
        }
        switch value__0._tag {
        case 0:
            jp3 = "closed"
        case 1:
            jp3 = "open"
        default:
            panic("non-exhaustive match")
        }
    }
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(jp3)
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

func _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_new____T__isize(capacity__0 int) chan int {
    var t0 chan int = func(p0 int) chan int {
        return make(chan int, p0)
    }(capacity__0)
    return t0
}

func _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_send____T__isize(self__0 chan int, value__0 int) struct{} {
    func(p0 chan int, p1 int) struct{} {
        p0 <- p1
        return struct{}{}
    }(self__0, value__0)
    return struct{}{}
}

func println__T_isize(value__0 int) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_int_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_new____T__string(capacity__0 int) chan string {
    var t0 chan string = func(p0 int) chan string {
        return make(chan string, p0)
    }(capacity__0)
    return t0
}

func _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_recv____T__string(self__0 chan string) Option__string {
    var mtmp0 Tuple2_6string_4bool = func(p0 chan string) Tuple2_6string_4bool {
        var value string
        var ok bool
        value, ok = <-p0
        return Tuple2_6string_4bool{
            _0: value,
            _1: ok,
        }
    }(self__0)
    var x0 string = mtmp0._0
    var x1 bool = mtmp0._1
    if x1 {
        var t0 Option__string = Option__string{
            _p0: x0,
            _tag: 1,
        }
        return t0
    } else {
        return Option__string{
            _tag: 0,
        }
    }
}

func _goml_m_trait__impl_i_ToString_i_isize_i_to__string(self__0 int) string {
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

func _goml_m_inherent_i_closure__env__read__0_i_closure__env__read__0_i_apply(env0 closure_env_read_0) int {
    var value__0 Option__isize = env0.value_0
    var inline0 int = -1
    switch value__0._tag {
    case 0:
        return inline0
    case 1:
        var inline1 int = value__0._p0
        return inline1
    default:
        panic("non-exhaustive match")
    }
}

func main() {
    main0()
}
