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

type Tuple2_3int_4bool struct {
    _0 int
    _1 bool
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

type closure_env_roundtrip_T_string_0 struct {
    channel_0 chan string
    value_1 string
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
    var buffered__0 chan int = _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_new____T__isize(2)
    _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_send____T__isize(buffered__0, 10)
    _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_send____T__isize(buffered__0, 20)
    _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_close____T__isize(buffered__0)
    var mtmp0 Option__isize
    var inline23 Tuple2_3int_4bool = func(p0 chan int) Tuple2_3int_4bool {
        var value int
        var ok bool
        value, ok = <-p0
        return Tuple2_3int_4bool{
            _0: value,
            _1: ok,
        }
    }(buffered__0)
    var inline24 int = inline23._0
    var inline25 bool = inline23._1
    if inline25 {
        var inline26 Option__isize = Option__isize{
            _p0: inline24,
            _tag: 1,
        }
        mtmp0 = inline26
    } else {
        mtmp0 = Option__isize{
            _tag: 0,
        }
    }
    var jp0 int
    switch mtmp0._tag {
    case 0:
        jp0 = -1
    case 1:
        var x2 int = mtmp0._p0
        jp0 = x2
    default:
        panic("non-exhaustive match")
    }
    var inline21 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(jp0)
    _goml_runtime_core_string_println(inline21)
    var mtmp1 Option__isize
    var inline17 Tuple2_3int_4bool = func(p0 chan int) Tuple2_3int_4bool {
        var value int
        var ok bool
        value, ok = <-p0
        return Tuple2_3int_4bool{
            _0: value,
            _1: ok,
        }
    }(buffered__0)
    var inline18 int = inline17._0
    var inline19 bool = inline17._1
    if inline19 {
        var inline20 Option__isize = Option__isize{
            _p0: inline18,
            _tag: 1,
        }
        mtmp1 = inline20
    } else {
        mtmp1 = Option__isize{
            _tag: 0,
        }
    }
    var jp1 int
    switch mtmp1._tag {
    case 0:
        jp1 = -1
    case 1:
        var x1 int = mtmp1._p0
        jp1 = x1
    default:
        panic("non-exhaustive match")
    }
    var inline15 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(jp1)
    _goml_runtime_core_string_println(inline15)
    var mtmp2 Option__isize
    var inline11 Tuple2_3int_4bool = func(p0 chan int) Tuple2_3int_4bool {
        var value int
        var ok bool
        value, ok = <-p0
        return Tuple2_3int_4bool{
            _0: value,
            _1: ok,
        }
    }(buffered__0)
    var inline12 int = inline11._0
    var inline13 bool = inline11._1
    if inline13 {
        var inline14 Option__isize = Option__isize{
            _p0: inline12,
            _tag: 1,
        }
        mtmp2 = inline14
    } else {
        mtmp2 = Option__isize{
            _tag: 0,
        }
    }
    var jp2 string
    switch mtmp2._tag {
    case 0:
        jp2 = "closed"
    case 1:
        jp2 = "open"
    default:
        panic("non-exhaustive match")
    }
    var inline9 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(jp2)
    _goml_runtime_core_string_println(inline9)
    var unbuffered__0 chan string
    var inline8 chan string = _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_new____T__string(0)
    unbuffered__0 = inline8
    var mtmp3 Option__string
    var inline3 string = "ready"
    var inline4 closure_env_roundtrip_T_string_0 = closure_env_roundtrip_T_string_0{
        channel_0: unbuffered__0,
        value_1: inline3,
    }
    var inline5 func() struct{} = func() struct{} {
        return _goml_m_inherent_i_closure__en_h6de6ca36633f2d445fcbfaf4a81e6dce_ring__0_i_apply(inline4)
    }
    go inline5()
    var inline7 Option__string = _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_recv____T__string(unbuffered__0)
    mtmp3 = inline7
    var jp3 string
    switch mtmp3._tag {
    case 0:
        jp3 = "closed"
    case 1:
        var x0 string = mtmp3._p0
        jp3 = x0
    default:
        panic("non-exhaustive match")
    }
    var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(jp3)
    _goml_runtime_core_string_println(inline1)
    func(p0 chan string) struct{} {
        close(p0)
        return struct{}{}
    }(unbuffered__0)
    return struct{}{}
}

func _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_new____T__string(capacity__0 int) chan string {
    var t0 chan string = func(p0 int) chan string {
        return make(chan string, p0)
    }(capacity__0)
    return t0
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

func _goml_m_inherent_i_Channel_i_Channel_l_T_r__i_close____T__isize(self__0 chan int) struct{} {
    func(p0 chan int) struct{} {
        close(p0)
        return struct{}{}
    }(self__0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_isize_i_to__string(self__0 int) string {
    var inline0 int64 = int64(int(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
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

func _goml_m_inherent_i_closure__en_h6de6ca36633f2d445fcbfaf4a81e6dce_ring__0_i_apply(env0 closure_env_roundtrip_T_string_0) struct{} {
    var channel__0 chan string = env0.channel_0
    var value__0 string = env0.value_1
    func(p0 chan string, p1 string) struct{} {
        p0 <- p1
        return struct{}{}
    }(channel__0, value__0)
    return struct{}{}
}

func main() {
    main0()
}
