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

func main0() struct{} {
    var a__0 int8
    a__0 = 90
    var b__0 int8
    b__0 = -20
    var c__0 int8
    c__0 = 3
    var sum__0 int8 = a__0 + b__0
    var diff__0 int8 = a__0 - c__0
    var prod__0 int8 = b__0 * c__0
    var quot__0 int8 = a__0 / c__0
    var neg__0 int8 = -b__0
    var less__0 bool = b__0 < a__0
    var inline32 string = "a="
    var inline33 string = _goml_m_trait__impl_i_ToString_i_i8_i_to__string(a__0)
    var inline34 string = inline32 + inline33
    println__T_string(inline34)
    var inline28 string = "b="
    var inline29 string = _goml_m_trait__impl_i_ToString_i_i8_i_to__string(b__0)
    var inline30 string = inline28 + inline29
    println__T_string(inline30)
    var inline24 string = "c="
    var inline25 string = _goml_m_trait__impl_i_ToString_i_i8_i_to__string(c__0)
    var inline26 string = inline24 + inline25
    println__T_string(inline26)
    var inline20 string = "sum="
    var inline21 string = _goml_m_trait__impl_i_ToString_i_i8_i_to__string(sum__0)
    var inline22 string = inline20 + inline21
    println__T_string(inline22)
    var inline16 string = "diff="
    var inline17 string = _goml_m_trait__impl_i_ToString_i_i8_i_to__string(diff__0)
    var inline18 string = inline16 + inline17
    println__T_string(inline18)
    var inline12 string = "prod="
    var inline13 string = _goml_m_trait__impl_i_ToString_i_i8_i_to__string(prod__0)
    var inline14 string = inline12 + inline13
    println__T_string(inline14)
    var inline8 string = "quot="
    var inline9 string = _goml_m_trait__impl_i_ToString_i_i8_i_to__string(quot__0)
    var inline10 string = inline8 + inline9
    println__T_string(inline10)
    var inline4 string = "neg="
    var inline5 string = _goml_m_trait__impl_i_ToString_i_i8_i_to__string(neg__0)
    var inline6 string = inline4 + inline5
    println__T_string(inline6)
    var inline0 string = "b<a="
    var inline1 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(less__0)
    var inline2 string = inline0 + inline1
    println__T_string(inline2)
    return struct{}{}
}

func println__T_string(value__0 string) struct{} {
    var t0 string
    t0 = value__0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_i8_i_to__string(self__0 int8) string {
    var inline0 int64 = int64(int8(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func _goml_m_trait__impl_i_ToString_i_bool_i_to__string(self__0 bool) string {
    var t0 string = _goml_runtime_core_bool_to_string(self__0)
    return t0
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
