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

type Tuple2_11Sender_3int_13Receiver_3int struct {
    _0 chan<- int
    _1 <-chan int
}

type Tuple2_3int_4bool struct {
    _0 int
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

type Ordering uint8

type Option__isize struct {
    _p0 int
    _tag uint8
}

func main0() struct{} {
    var channel__0 chan int
    var inline25 int = 2
    var inline26 chan int = func(p0 int) chan int {
        return make(chan int, p0)
    }(inline25)
    channel__0 = inline26
    var mtmp0 Tuple2_11Sender_3int_13Receiver_3int
    var inline24 Tuple2_11Sender_3int_13Receiver_3int = func(p0 chan int) Tuple2_11Sender_3int_13Receiver_3int {
        return Tuple2_11Sender_3int_13Receiver_3int{
            _0: p0,
            _1: p0,
        }
    }(channel__0)
    mtmp0 = inline24
    var x0 chan<- int = mtmp0._0
    var x1 <-chan int = mtmp0._1
    var inline22 int = 7
    func(p0 chan<- int, p1 int) struct{} {
        p0 <- p1
        return struct{}{}
    }(x0, inline22)
    var _goml_m_value____0_i_select__value int
    var _goml_m_value____0_i_select__open bool
    var value__0 Option__isize = Option__isize{
        _tag: 0,
    }
    select {
    case _goml_m_value____0_i_select__value, _goml_m_value____0_i_select__open = <-x1:
        if _goml_m_value____0_i_select__open {
            value__0 = Option__isize{
                _p0: _goml_m_value____0_i_select__value,
                _tag: 1,
            }
        }
        var t2 int
        var inline17 int = 0
        switch value__0._tag {
        case 0:
            t2 = inline17
        case 1:
            var inline18 int = value__0._p0
            t2 = inline18
        default:
            panic("non-exhaustive match")
        }
        var inline15 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t2)
        _goml_runtime_core_string_println(inline15)
    default:
        var inline19 int = 1
        var inline20 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(inline19)
        _goml_runtime_core_string_println(inline20)
    }
    var channel__1 chan int
    var inline13 int = 1
    var inline14 chan int = func(p0 int) chan int {
        return make(chan int, p0)
    }(inline13)
    channel__1 = inline14
    var sender__0 chan<- int
    var inline12 chan<- int = func(p0 chan int) chan<- int {
        return p0
    }(channel__1)
    sender__0 = inline12
    var receiver__0 <-chan int
    var inline11 <-chan int = func(p0 chan int) <-chan int {
        return p0
    }(channel__1)
    receiver__0 = inline11
    select {
    case sender__0 <- 9:
        var t0 Option__isize
        var inline4 Tuple2_3int_4bool = func(p0 <-chan int) Tuple2_3int_4bool {
            var value int
            var ok bool
            value, ok = <-p0
            return Tuple2_3int_4bool{
                _0: value,
                _1: ok,
            }
        }(receiver__0)
        var inline5 int = inline4._0
        var inline6 bool = inline4._1
        if inline6 {
            var inline7 Option__isize = Option__isize{
                _p0: inline5,
                _tag: 1,
            }
            t0 = inline7
        } else {
            t0 = Option__isize{
                _tag: 0,
            }
        }
        var t1 int
        var inline2 int = 0
        switch t0._tag {
        case 0:
            t1 = inline2
        case 1:
            var inline3 int = t0._p0
            t1 = inline3
        default:
            panic("non-exhaustive match")
        }
        var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t1)
        _goml_runtime_core_string_println(inline0)
        return struct{}{}
    default:
        var inline8 int = 2
        var inline9 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(inline8)
        _goml_runtime_core_string_println(inline9)
        return struct{}{}
    }
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
