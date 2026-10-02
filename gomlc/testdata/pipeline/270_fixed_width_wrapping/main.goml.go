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

const (
    WRAPPED_I8 int8 = -128
    WRAPPED_U8 uint8 = 255
    WRAPPED_PRODUCT uint8 = 0
    WRAPPED_DIVISION int8 = -128
    WRAPPED_NEGATION int8 = -128
    FORWARD_WRAPPED int8 = -128
    MAX_I8 int8 = 127
    NEGATIVE_ONE_I16 int16 = -1
    PREFIX string = "fixed"
)

func main0() struct{} {
    var t0_lhs int8 = 127
    var t0_rhs int8 = 1
    var t0 int8 = t0_lhs + t0_rhs
    println__T_i8(t0)
    var t1_lhs uint8 = 0
    var t1_rhs uint8 = 1
    var t1 uint8 = t1_lhs - t1_rhs
    println__T_u8(t1)
    var t2_lhs uint8 = 128
    var t2_rhs uint8 = 2
    var t2 uint8 = t2_lhs * t2_rhs
    println__T_u8(t2)
    var t3_lhs int8 = -128
    var t3_rhs int8 = -1
    var t3 int8 = t3_lhs / t3_rhs
    println__T_i8(t3)
    var t4_operand int8 = -128
    var t4 int8 = -t4_operand
    println__T_i8(t4)
    var t5_lhs int64 = 9223372036854775807
    var t5_rhs int64 = 1
    var t5 int64 = t5_lhs + t5_rhs
    println__T_i64(t5)
    var t6_lhs uint64 = 18446744073709551615
    var t6_rhs uint64 = 1
    var t6 uint64 = t6_lhs + t6_rhs
    println__T_u64(t6)
    println__T_i8(WRAPPED_I8)
    println__T_u8(WRAPPED_U8)
    println__T_u8(WRAPPED_PRODUCT)
    println__T_i8(WRAPPED_DIVISION)
    println__T_i8(WRAPPED_NEGATION)
    println__T_i8(FORWARD_WRAPPED)
    var t7_lhs int8 = MAX_I8
    var t7_rhs int8 = 1
    var t7 int8 = t7_lhs + t7_rhs
    println__T_i8(t7)
    var t8_lhs int8 = WRAPPED_I8
    var t8_rhs int8 = -1
    var t8 int8 = t8_lhs / t8_rhs
    println__T_i8(t8)
    var t9_operand int8 = WRAPPED_I8
    var t9 int8 = -t9_operand
    println__T_i8(t9)
    var converted__0_source int16 = NEGATIVE_ONE_I16
    var converted__0 uint8 = uint8(int16(converted__0_source))
    println__T_u8(converted__0)
    var runtime_value__0 int8 = 127
    var t10 int8 = runtime_value__0 + 1
    println__T_i8(t10)
    var t11 string = PREFIX + " width"
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t11)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func println__T_i8(value__0 int8) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_int8_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func println__T_u8(value__0 uint8) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_uint8_to_string(value__0)
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

func println__T_u64(value__0 uint64) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_uint64_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
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

func __goml_builtin_uint8_to_string(value__0 uint8) string {
    var t0 uint64 = uint64(uint8(value__0))
    var t1 string = decimal_string(t0)
    return t1
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

func __goml_builtin_uint64_to_string(value__0 uint64) string {
    var t0 string = decimal_string(value__0)
    return t0
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
