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

func _goml_runtime_core_string_get(s string, i int) rune {
    return rune(s[i])
}

func _goml_runtime_core_string_from_utf8(bytes *_goml_vec_uint8) Tuple2_4bool_6string {
    return Tuple2_4bool_6string{
        _0: true,
        _1: string(bytes.items),
    }
}

func _goml_runtime_core_char_to_string(x rune) string {
    return string(x)
}

func _goml_runtime_core_char_from_uint32(value uint32) Tuple2_4bool_4char {
    return Tuple2_4bool_4char{
        _0: true,
        _1: rune(value),
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

type Tuple2_4bool_4char struct {
    _0 bool
    _1 rune
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

type Option__char uint64

func show_u8(value__0 uint8) struct{} {
    var inline0 string = _goml_m_trait__impl_i_ToString_i_u8_i_to__string(value__0)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func unsigned_ops() struct{} {
    var t0_lhs uint8 = 13
    var t0_rhs uint8 = 5
    var t0 uint8 = t0_lhs % t0_rhs
    show_u8(t0)
    var t1_lhs uint8 = 12
    var t1_rhs uint8 = 10
    var t1 uint8 = t1_lhs & t1_rhs
    show_u8(t1)
    var t2_lhs uint8 = 12
    var t2_rhs uint8 = 3
    var t2 uint8 = t2_lhs | t2_rhs
    show_u8(t2)
    var t3_lhs uint8 = 12
    var t3_rhs uint8 = 10
    var t3 uint8 = t3_lhs ^ t3_rhs
    show_u8(t3)
    var t4_lhs uint8 = 1
    var t4_rhs int = 7
    var t4 uint8 = t4_lhs << t4_rhs
    show_u8(t4)
    var t5_lhs uint8 = 128
    var t5_rhs int = 7
    var t5 uint8 = t5_lhs >> t5_rhs
    println__T_u8(t5)
    var t6_operand uint8 = 0
    var t6 uint8 = ^t6_operand
    println__T_u8(t6)
    var t7_lhs uint8 = 1
    var t7_rhs int = 8
    var t7 uint8 = t7_lhs << t7_rhs
    println__T_u8(t7)
    var t8_lhs uint16 = 513
    var t8_rhs uint16 = 256
    var t8 uint16 = t8_lhs % t8_rhs
    println__T_u16(t8)
    var t9_lhs uint16 = 3855
    var t9_rhs uint16 = 255
    var t9 uint16 = t9_lhs & t9_rhs
    println__T_u16(t9)
    var t10_lhs uint16 = 3840
    var t10_rhs uint16 = 15
    var t10 uint16 = t10_lhs | t10_rhs
    println__T_u16(t10)
    var t11_lhs uint16 = 43690
    var t11_rhs uint16 = 3855
    var t11 uint16 = t11_lhs ^ t11_rhs
    println__T_u16(t11)
    var t12_lhs uint16 = 1
    var t12_rhs int = 15
    var t12 uint16 = t12_lhs << t12_rhs
    println__T_u16(t12)
    var t13_lhs uint16 = 32768
    var t13_rhs int = 15
    var t13 uint16 = t13_lhs >> t13_rhs
    println__T_u16(t13)
    var t14_operand uint16 = 0
    var t14 uint16 = ^t14_operand
    println__T_u16(t14)
    var t15_lhs uint32 = 1000000001
    var t15_rhs uint32 = 1000
    var t15 uint32 = t15_lhs % t15_rhs
    println__T_u32(t15)
    var t16_lhs uint32 = 4042322160
    var t16_rhs uint32 = 252645135
    var t16 uint32 = t16_lhs & t16_rhs
    println__T_u32(t16)
    var t17_lhs uint32 = 4042322160
    var t17_rhs uint32 = 252645135
    var t17 uint32 = t17_lhs | t17_rhs
    println__T_u32(t17)
    var t18_lhs uint32 = 4042322160
    var t18_rhs uint32 = 252645135
    var t18 uint32 = t18_lhs ^ t18_rhs
    println__T_u32(t18)
    var t19_lhs uint32 = 1
    var t19_rhs int = 31
    var t19 uint32 = t19_lhs << t19_rhs
    println__T_u32(t19)
    var t20_lhs uint32 = 2147483648
    var t20_rhs int = 31
    var t20 uint32 = t20_lhs >> t20_rhs
    println__T_u32(t20)
    var t21_operand uint32 = 0
    var t21 uint32 = ^t21_operand
    println__T_u32(t21)
    var t22_lhs uint64 = 1000000000001
    var t22_rhs uint64 = 1000
    var t22 uint64 = t22_lhs % t22_rhs
    println__T_u64(t22)
    var t23_lhs uint64 = 17361641481138401520
    var t23_rhs uint64 = 1085102592571150095
    var t23 uint64 = t23_lhs & t23_rhs
    println__T_u64(t23)
    var t24_lhs uint64 = 17361641481138401520
    var t24_rhs uint64 = 1085102592571150095
    var t24 uint64 = t24_lhs | t24_rhs
    println__T_u64(t24)
    var t25_lhs uint64 = 17361641481138401520
    var t25_rhs uint64 = 1085102592571150095
    var t25 uint64 = t25_lhs ^ t25_rhs
    println__T_u64(t25)
    var t26_lhs uint64 = 1
    var t26_rhs int = 63
    var t26 uint64 = t26_lhs << t26_rhs
    println__T_u64(t26)
    var t27_lhs uint64 = 9223372036854775808
    var t27_rhs int = 63
    var t27 uint64 = t27_lhs >> t27_rhs
    println__T_u64(t27)
    var t28_operand uint64 = 0
    var t28 uint64 = ^t28_operand
    println__T_u64(t28)
    return struct{}{}
}

func signed_ops() struct{} {
    var t0_lhs int8 = -13
    var t0_rhs int8 = 5
    var t0 int8 = t0_lhs % t0_rhs
    println__T_i8(t0)
    var t1_lhs int8 = -8
    var t1_rhs int = 2
    var t1 int8 = t1_lhs >> t1_rhs
    println__T_i8(t1)
    var t2_lhs int8 = 1
    var t2_rhs int = 6
    var t2 int8 = t2_lhs << t2_rhs
    println__T_i8(t2)
    var t3_operand int8 = 0
    var t3 int8 = ^t3_operand
    println__T_i8(t3)
    var t4_lhs int8 = -1
    var t4_rhs int = 7
    var t4 int8 = t4_lhs >> t4_rhs
    println__T_i8(t4)
    var t5_lhs int16 = -513
    var t5_rhs int16 = 256
    var t5 int16 = t5_lhs % t5_rhs
    println__T_i16(t5)
    var t6_lhs int16 = -32767
    var t6_rhs int16 = 1
    var t6 int16 = t6_lhs - t6_rhs
    var t7 int16 = t6 >> 15
    println__T_i16(t7)
    var t8_lhs int16 = 1
    var t8_rhs int = 14
    var t8 int16 = t8_lhs << t8_rhs
    println__T_i16(t8)
    var t9_operand int16 = 255
    var t9 int16 = ^t9_operand
    println__T_i16(t9)
    var t10_lhs int32 = -1000000001
    var t10_rhs int32 = 1000
    var t10 int32 = t10_lhs % t10_rhs
    println__T_i32(t10)
    var t11_lhs int32 = -2147483647
    var t11_rhs int32 = 1
    var t11 int32 = t11_lhs - t11_rhs
    var t12 int32 = t11 >> 31
    println__T_i32(t12)
    var t13_lhs int32 = 1
    var t13_rhs int = 30
    var t13 int32 = t13_lhs << t13_rhs
    println__T_i32(t13)
    var t14_operand int32 = 65535
    var t14 int32 = ^t14_operand
    println__T_i32(t14)
    var t15_lhs int64 = -1000000000001
    var t15_rhs int64 = 1000
    var t15 int64 = t15_lhs % t15_rhs
    println__T_i64(t15)
    var t16_lhs int64 = -9223372036854775807
    var t16_rhs int = 62
    var t16 int64 = t16_lhs >> t16_rhs
    println__T_i64(t16)
    var t17_lhs int64 = 1
    var t17_rhs int = 62
    var t17 int64 = t17_lhs << t17_rhs
    println__T_i64(t17)
    var t18_operand int64 = 4294967295
    var t18 int64 = ^t18_operand
    println__T_i64(t18)
    return struct{}{}
}

func precedence() struct{} {
    var t0_lhs uint8 = 3
    var t0_rhs uint8 = 1
    var t0 uint8 = t0_lhs & t0_rhs
    var t1 uint8 = 2 ^ t0
    var t2 uint8 = 1 | t1
    println__T_u8(t2)
    var t3_lhs int = 2
    var t3_rhs int = 1
    var t3 int = t3_lhs + t3_rhs
    var t4 uint8 = 1 << t3
    println__T_u8(t4)
    var t5_lhs int = 1
    var t5_rhs int = 2
    var t5 int = t5_lhs | t5_rhs
    var t6 bool = t5 == 3
    var t7 string
    var inline6 string = _goml_runtime_core_bool_to_string(t6)
    t7 = inline6
    var inline4 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t7)
    _goml_runtime_core_string_println(inline4)
    var t8_lhs int = 8
    var t8_rhs int = 1
    var t8 int = t8_lhs >> t8_rhs
    var t9 bool = t8 < 5
    var t10 string
    var inline3 string = _goml_runtime_core_bool_to_string(t9)
    t10 = inline3
    var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t10)
    _goml_runtime_core_string_println(inline1)
    var t11_operand uint8 = 1
    var t11 uint8 = ^t11_operand
    var t12 uint8 = t11 & 15
    println__T_u8(t12)
    return struct{}{}
}

func casts() struct{} {
    var five_eleven__0 uint16 = 511
    var two_fifty_six__0 uint16 = 256
    var negative_one_i16__0 int16 = -1
    var two_fifty_five__0 uint8 = 255
    var one_twenty_eight__0 uint8 = 128
    var negative_one_twenty_nine__0 int16 = -129
    var max_u16__0 uint16 = 65535
    var negative_one_i32__0 int32 = -1
    var negative_one_i8__0 int8 = -1
    var max_u64__0 uint64 = 18446744073709551615
    var sixty_five__0 uint8 = 65
    var max_u32__0 uint32 = 4294967295
    var three_hundred__0 uint16 = 300
    var t0 uint8 = uint8(uint16(five_eleven__0))
    println__T_u8(t0)
    var t1 uint8 = uint8(uint16(two_fifty_six__0))
    println__T_u8(t1)
    var t2 uint8 = uint8(int16(negative_one_i16__0))
    println__T_u8(t2)
    var t3 int8 = int8(uint8(two_fifty_five__0))
    println__T_i8(t3)
    var t4 int8 = int8(uint8(one_twenty_eight__0))
    println__T_i8(t4)
    var t5 int8 = int8(int16(negative_one_twenty_nine__0))
    println__T_i8(t5)
    var t6 int16 = int16(uint16(max_u16__0))
    println__T_i16(t6)
    var t7 uint16 = uint16(int32(negative_one_i32__0))
    println__T_u16(t7)
    var t8 uint64 = uint64(int8(negative_one_i8__0))
    println__T_u64(t8)
    var t9 int32 = int32(uint64(max_u64__0))
    println__T_i32(t9)
    var t10 uint32 = uint32(uint8(sixty_five__0))
    println__T_u32(t10)
    var t11 int64 = int64(uint32(max_u32__0))
    println__T_i64(t11)
    var t12_source rune = 65
    var t12 uint32 = uint32(rune(t12_source))
    println__T_u32(t12)
    var mtmp0 Option__char
    var inline7 uint32 = 128512
    var inline8 Option__char = __goml_builtin_char_from_uint32(inline7)
    mtmp0 = inline8
    switch mtmp0 != Option__char(0) {
    case false:
        var inline1 string = "invalid"
        var inline2 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline1)
        _goml_runtime_core_string_println(inline2)
    case true:
        var x0 rune = rune(uint64(mtmp0) - 1)
        var t15 string
        var inline6 string = char_to_string(x0)
        t15 = inline6
        var inline4 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t15)
        _goml_runtime_core_string_println(inline4)
    default:
        panic("non-exhaustive match")
    }
    var t13 uint8 = uint8(uint16(three_hundred__0))
    var t14 uint32 = uint32(uint8(t13))
    println__T_u32(t14)
    return struct{}{}
}

func main0() struct{} {
    unsigned_ops()
    signed_ops()
    precedence()
    casts()
    var t0 uint8
    var inline1 uint8 = 10
    var inline2 uint8 = ^inline1
    var inline3 uint8 = inline2 & 15
    var inline4_lhs uint8 = 1
    var inline4_rhs int = 4
    var inline4 uint8 = inline4_lhs << inline4_rhs
    var inline5 uint8 = inline4 % 31
    var inline6 uint8 = inline3 | inline5
    t0 = inline6
    println__T_u8(t0)
    return struct{}{}
}

func println__T_u8(value__0 uint8) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_uint8_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func println__T_u16(value__0 uint16) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_uint16_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func println__T_u32(value__0 uint32) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_uint32_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func println__T_u64(value__0 uint64) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_uint64_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func println__T_i8(value__0 int8) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_int8_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func println__T_i16(value__0 int16) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_int16_to_string(value__0)
    t0 = inline0
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

func println__T_i64(value__0 int64) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_int64_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_u8_i_to__string(self__0 uint8) string {
    var inline0 uint64 = uint64(uint8(self__0))
    var inline1 string = decimal_string(inline0)
    return inline1
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func __goml_builtin_char_from_uint32(value__0 uint32) Option__char {
    var t0 bool
    var inline0 bool = value__0 <= 1114111
    if inline0 {
        var inline1 bool = value__0 >= 55296
        var inline2 bool
        if inline1 {
            var inline4 bool = value__0 <= 57343
            inline2 = inline4
        } else {
            inline2 = false
        }
        var inline3 bool = !inline2
        t0 = inline3
    } else {
        t0 = false
    }
    if t0 {
        var mtmp0 Tuple2_4bool_4char = _goml_runtime_core_char_from_uint32(value__0)
        var x0 rune = mtmp0._1
        var t1 Option__char = Option__char(uint64(uint32(x0)) + 1)
        return t1
    } else {
        return Option__char(0)
    }
}

func char_to_string(value__0 rune) string {
    var t0 uint32 = uint32(rune(value__0))
    var t1 bool
    var inline0 bool = t0 <= 1114111
    if inline0 {
        var inline1 bool = t0 >= 55296
        var inline2 bool
        if inline1 {
            var inline4 bool = t0 <= 57343
            inline2 = inline4
        } else {
            inline2 = false
        }
        var inline3 bool = !inline2
        t1 = inline3
    } else {
        t1 = false
    }
    if t1 {
        var t2 string = _goml_runtime_core_char_to_string(value__0)
        return t2
    } else {
        _goml_runtime_core_string_get("", -1)
        return ""
    }
}

func __goml_builtin_uint8_to_string(value__0 uint8) string {
    var t0 uint64 = uint64(uint8(value__0))
    var t1 string = decimal_string(t0)
    return t1
}

func __goml_builtin_uint16_to_string(value__0 uint16) string {
    var t0 uint64 = uint64(uint16(value__0))
    var t1 string = decimal_string(t0)
    return t1
}

func __goml_builtin_uint32_to_string(value__0 uint32) string {
    var t0 uint64 = uint64(uint32(value__0))
    var t1 string = decimal_string(t0)
    return t1
}

func __goml_builtin_uint64_to_string(value__0 uint64) string {
    var t0 string = decimal_string(value__0)
    return t0
}

func __goml_builtin_int8_to_string(value__0 int8) string {
    var t0 int64 = int64(int8(value__0))
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

func __goml_builtin_int16_to_string(value__0 int16) string {
    var t0 int64 = int64(int16(value__0))
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

func __goml_builtin_int64_to_string(value__0 int64) string {
    var inline0 bool = value__0 < 0
    if inline0 {
        var inline1 uint64 = uint64(int64(value__0))
        var inline2 uint64 = 0 - inline1
        var inline3 string = decimal_string(inline2)
        var inline4 string = "-" + inline3
        return inline4
    } else {
        var inline5 uint64 = uint64(int64(value__0))
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
