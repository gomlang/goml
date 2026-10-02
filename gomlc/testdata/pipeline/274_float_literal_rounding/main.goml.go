package main

import (
    _goml_ffi_import_math_h0cea0c730c0e331b7d5cc103b112df29 "math"
)

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

func _goml_ffi_math_x00_Float32bits_x0__q__m__z_u32_hbbf8d280343f673a9e2fd959393f1495(arg0 float32) uint32 {
    return _goml_ffi_import_math_h0cea0c730c0e331b7d5cc103b112df29.Float32bits(arg0)
}

func _goml_ffi_math_x00_Float32frombit__q__m__z_f32_h1605b723e1fe36562e475cf246e4961c(arg0 uint32) float32 {
    return _goml_ffi_import_math_h0cea0c730c0e331b7d5cc103b112df29.Float32frombits(arg0)
}

func _goml_ffi_math_x00_Float64bits_x0__q__m__z_u64_h771d143dab10df93b09bfe0f407ff63e(arg0 float64) uint64 {
    return _goml_ffi_import_math_h0cea0c730c0e331b7d5cc103b112df29.Float64bits(arg0)
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
    var below__0 float32 = 1
    var midpoint__0 float32 = 1
    var above__0 float32 = 1.0000001
    var t0 uint32
    var inline27 uint32 = __goml_builtin_float32_to_bits(below__0)
    t0 = inline27
    println__T_u32(t0)
    var t1 uint32
    var inline26 uint32 = __goml_builtin_float32_to_bits(midpoint__0)
    t1 = inline26
    var inline24 string = _goml_m_trait__impl_i_ToString_i_u32_i_to__string(t1)
    _goml_runtime_core_string_println(inline24)
    var t2 uint32
    var inline23 uint32 = __goml_builtin_float32_to_bits(above__0)
    t2 = inline23
    var inline21 string = _goml_m_trait__impl_i_ToString_i_u32_i_to__string(t2)
    _goml_runtime_core_string_println(inline21)
    var generic__0 float32
    var inline20 float32 = 1.0000001
    generic__0 = inline20
    var t3 uint32
    var inline19 uint32 = __goml_builtin_float32_to_bits(generic__0)
    t3 = inline19
    var inline17 string = _goml_m_trait__impl_i_ToString_i_u32_i_to__string(t3)
    _goml_runtime_core_string_println(inline17)
    var t4 float32
    var inline15 uint32 = 1065353217
    var inline16 float32 = __goml_builtin_float32_from_bits(inline15)
    t4 = inline16
    var t5 int
    switch t4 {
    case 1.0000001:
        t5 = 1
    default:
        t5 = 0
    }
    var t6 string
    var inline14 string = __goml_builtin_int_to_string(t5)
    t6 = inline14
    var inline12 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t6)
    _goml_runtime_core_string_println(inline12)
    var tiny32__0 float32 = 0
    var tiny64__0 float64 = 0
    var t7 uint32
    var inline11 uint32 = __goml_builtin_float32_to_bits(tiny32__0)
    t7 = inline11
    var inline9 string = _goml_m_trait__impl_i_ToString_i_u32_i_to__string(t7)
    _goml_runtime_core_string_println(inline9)
    var t8 uint64
    var inline8 uint64 = __goml_builtin_float64_to_bits(tiny64__0)
    t8 = inline8
    var inline6 string = _goml_m_trait__impl_i_ToString_i_u64_i_to__string(t8)
    _goml_runtime_core_string_println(inline6)
    var maximum32__0 float32 = 340282350000000000000000000000000000000
    var maximum64__0 float64 = 17976931348623157e292
    var t9 uint32
    var inline5 uint32 = __goml_builtin_float32_to_bits(maximum32__0)
    t9 = inline5
    var inline3 string = _goml_m_trait__impl_i_ToString_i_u32_i_to__string(t9)
    _goml_runtime_core_string_println(inline3)
    var t10 uint64
    var inline2 uint64 = __goml_builtin_float64_to_bits(maximum64__0)
    t10 = inline2
    var inline0 string = _goml_m_trait__impl_i_ToString_i_u64_i_to__string(t10)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func println__T_u32(value__0 uint32) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_uint32_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_u32_i_to__string(self__0 uint32) string {
    var inline0 uint64 = uint64(uint32(self__0))
    var inline1 string = decimal_string(inline0)
    return inline1
}

func __goml_builtin_float32_to_bits(value__0 float32) uint32 {
    var t0 uint32 = _goml_ffi_math_x00_Float32bits_x0__q__m__z_u32_hbbf8d280343f673a9e2fd959393f1495(value__0)
    return t0
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func __goml_builtin_float32_from_bits(value__0 uint32) float32 {
    var t0 float32 = _goml_ffi_math_x00_Float32frombit__q__m__z_f32_h1605b723e1fe36562e475cf246e4961c(value__0)
    return t0
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

func _goml_m_trait__impl_i_ToString_i_u64_i_to__string(self__0 uint64) string {
    var inline0 string = decimal_string(self__0)
    return inline0
}

func __goml_builtin_float64_to_bits(value__0 float64) uint64 {
    var t0 uint64 = _goml_ffi_math_x00_Float64bits_x0__q__m__z_u64_h771d143dab10df93b09bfe0f407ff63e(value__0)
    return t0
}

func __goml_builtin_uint32_to_string(value__0 uint32) string {
    var t0 uint64 = uint64(uint32(value__0))
    var t1 string = decimal_string(t0)
    return t1
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
