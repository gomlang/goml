package main

import (
    _goml_ffi_import_math_h0cea0c730c0e331b7d5cc103b112df29 "math"
)

import (
    _goml_os "os"
)

func _goml_runtime_core_string_len(s string) int {
    return int(len(s))
}

func _goml_runtime_core_string_byte_get(s string, i int) uint8 {
    return s[i]
}

func _goml_runtime_core_string_byte_slice(s string, start int, end int) string {
    return s[start:end]
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

func _goml_ffi_math_x00_Float32bits_x0__q__m__z_u32_hbbf8d280343f673a9e2fd959393f1495(arg0 float32) uint32 {
    return _goml_ffi_import_math_h0cea0c730c0e331b7d5cc103b112df29.Float32bits(arg0)
}

type _goml_vec_uint8 struct {
    items []uint8
}

func vec_new__Vec_5uint8() *_goml_vec_uint8 {
    return &_goml_vec_uint8{
        items: nil,
    }
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

func vec_set__Vec_5uint8(vec *_goml_vec_uint8, index int, value uint8) struct{} {
    vec.items[index] = value
    return struct{}{}
}

func vec_len__Vec_5uint8(vec *_goml_vec_uint8) int {
    return int(len(vec.items))
}

type _goml_vec_uint32 struct {
    items []uint32
}

func vec_new__Vec_6uint32() *_goml_vec_uint32 {
    return &_goml_vec_uint32{
        items: nil,
    }
}

func vec_push__Vec_6uint32(vec *_goml_vec_uint32, elem uint32) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_get__Vec_6uint32(vec *_goml_vec_uint32, index int) uint32 {
    return vec.items[index]
}

func vec_set__Vec_6uint32(vec *_goml_vec_uint32, index int, value uint32) struct{} {
    vec.items[index] = value
    return struct{}{}
}

func vec_len__Vec_6uint32(vec *_goml_vec_uint32) int {
    return int(len(vec.items))
}

func vec_truncate__Vec_6uint32(vec *_goml_vec_uint32, new_len int) struct{} {
    if new_len < 0 {
        panic("negative vector length")
    }
    if new_len < int(len(vec.items)) {
        clear(vec.items[new_len:int(len(vec.items))])
        vec.items = vec.items[0:new_len]
    }
    return struct{}{}
}

type Tuple2_4bool_6string struct {
    _0 bool
    _1 string
}

type Tuple2_6string_4bool struct {
    _0 string
    _1 bool
}

type Tuple2_4bool_6uint64 struct {
    _0 bool
    _1 uint64
}

type Tuple2_6uint64_4bool struct {
    _0 uint64
    _1 bool
}

type Tuple2_4bool_3int struct {
    _0 bool
    _1 int
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
    var a__0 uint8
    var inline11 uint8 = 42
    a__0 = inline11
    var t0 string
    var inline10 string = __goml_builtin_uint8_to_string(a__0)
    t0 = inline10
    var inline8 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t0)
    _goml_runtime_core_string_println(inline8)
    var b__0 float32
    var inline7 float32 = 3.14
    b__0 = inline7
    var t1 string
    var inline6 string = __goml_builtin_float32_to_string(b__0)
    t1 = inline6
    var inline4 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t1)
    _goml_runtime_core_string_println(inline4)
    var c__0 int64
    var inline3 int64 = 100
    c__0 = inline3
    var t2 string
    var inline2 string = __goml_builtin_int64_to_string(c__0)
    t2 = inline2
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t2)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func __goml_builtin_uint8_to_string(value__0 uint8) string {
    var t0 uint64 = uint64(uint8(value__0))
    var t1 string = decimal_string(t0)
    return t1
}

func __goml_builtin_float32_to_string(value__0 float32) string {
    var t0 uint32 = _goml_ffi_math_x00_Float32bits_x0__q__m__z_u32_hbbf8d280343f673a9e2fd959393f1495(value__0)
    var t1 uint64 = uint64(uint32(t0))
    var t2 string = format_float_bits(t1, 23, 8, 127)
    return t2
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

func format_float_bits(bits__0 uint64, mantissa_bits__0 int, exponent_bits__0 int, exponent_bias__0 int) string {
    var t0 int = mantissa_bits__0 + exponent_bits__0
    var sign_mask__0 uint64 = 1 << t0
    var t1 uint64 = bits__0 & sign_mask__0
    var negative__0 bool = t1 != 0
    var t2 uint64 = 1 << exponent_bits__0
    var exponent_mask__0 uint64 = t2 - 1
    var t3 uint64 = bits__0 >> mantissa_bits__0
    var exponent__0 uint64 = t3 & exponent_mask__0
    var t4 uint64 = 1 << mantissa_bits__0
    var t5 uint64 = t4 - 1
    var fraction__0 uint64 = bits__0 & t5
    var t6 bool = exponent__0 == exponent_mask__0
    if t6 {
        var t40 bool = fraction__0 == 0
        if t40 {
            if negative__0 {
                return "-inf"
            } else {
                return "inf"
            }
        } else {
            return "NaN"
        }
    } else {
        var t41 bool = exponent__0 == 0
        var jp9 bool
        if t41 {
            var t42 bool = fraction__0 == 0
            jp9 = t42
        } else {
            jp9 = false
        }
        if jp9 {
            if negative__0 {
                return "-0"
            } else {
                return "0"
            }
        } else {
            var t7 bool = exponent__0 == 0
            var jp0 uint64
            if t7 {
                jp0 = fraction__0
            } else {
                var t38 uint64 = 1 << mantissa_bits__0
                var t39 uint64 = fraction__0 | t38
                jp0 = t39
            }
            var t8 bool = exponent__0 == 0
            var jp1 int
            if t8 {
                var t33 int = 1 - exponent_bias__0
                var t34 int = t33 - mantissa_bits__0
                jp1 = t34
            } else {
                var t35 int = int(uint64(exponent__0))
                var t36 int = t35 - exponent_bias__0
                var t37 int = t36 - mantissa_bits__0
                jp1 = t37
            }
            var exact_value__0 FloatNatural = float_natural_from_u64(jp0)
            var t9 bool = jp1 >= 0
            var jp2 int
            if t9 {
                var shifted__0 FloatNatural = float_natural_shift_left(exact_value__0, jp1)
                var digits__0 string = float_natural_decimal(shifted__0)
                var t12 bool = mantissa_bits__0 == 23
                var jp3 int
                if t12 {
                    jp3 = 9
                } else {
                    jp3 = 17
                }
                var t13 int
                var inline3 int = _goml_runtime_core_string_len(digits__0)
                t13 = inline3
                var t14 bool = t13 < jp3
                var jp4 int
                if t14 {
                    var inline2 int = _goml_runtime_core_string_len(digits__0)
                    jp4 = inline2
                } else {
                    jp4 = jp3
                }
                var count__0 int = 1
                Loop_loop0:
                for {
                    var t15 bool = count__0 <= jp4
                    if t15 {
                        var mtmp0 Tuple2_6string_4bool = rounded_float_digits(digits__0, count__0)
                        var x0 string = mtmp0._0
                        var x1 bool = mtmp0._1
                        var rounded__0 string = trim_float_digits(x0)
                        var t16 int
                        var inline1 int = _goml_runtime_core_string_len(digits__0)
                        t16 = inline1
                        var jp5 int
                        if x1 {
                            jp5 = 1
                        } else {
                            jp5 = 0
                        }
                        var point__0 int = t16 + jp5
                        var candidate__0 string = fixed_float_text(rounded__0, point__0, negative__0)
                        var mtmp1 Tuple2_4bool_6uint64 = parsed_float_bits(candidate__0, mantissa_bits__0, exponent_bias__0)
                        var x2 uint64 = mtmp1._1
                        var t17 bool = x2 == bits__0
                        if t17 {
                            return candidate__0
                        } else {
                            var compound_old0 int = count__0
                            var compound_value0 int = 1
                            var t18 int = compound_old0 + compound_value0
                            count__0 = t18
                            continue
                        }
                    } else {
                        break Loop_loop0
                    }
                }
                var inline0 int = _goml_runtime_core_string_len(digits__0)
                jp2 = inline0
                var t10 string = float_natural_decimal(exact_value__0)
                var t11 string = fixed_float_text(t10, jp2, negative__0)
                return t11
            } else {
                var count__1 int = 0
                var t29 int = 0 - jp1
                Loop_loop1:
                for {
                    var t30 bool = count__1 < t29
                    if t30 {
                        float_natural_multiply_small(exact_value__0, 5)
                        var compound_old2 int = count__1
                        var compound_value2 int = 1
                        var t31 int = compound_old2 + compound_value2
                        count__1 = t31
                        continue
                    } else {
                        break Loop_loop1
                    }
                }
                var digits__1 string = float_natural_decimal(exact_value__0)
                var t20 int
                var inline6 int = _goml_runtime_core_string_len(digits__1)
                t20 = inline6
                var point__1 int = t20 + jp1
                var t21 bool = mantissa_bits__0 == 23
                var jp6 int
                if t21 {
                    jp6 = 9
                } else {
                    jp6 = 17
                }
                var t22 int
                var inline5 int = _goml_runtime_core_string_len(digits__1)
                t22 = inline5
                var t23 bool = t22 < jp6
                var jp7 int
                if t23 {
                    var inline4 int = _goml_runtime_core_string_len(digits__1)
                    jp7 = inline4
                } else {
                    jp7 = jp6
                }
                count__1 = 1
                Loop_loop2:
                for {
                    var t24 bool = count__1 <= jp7
                    if t24 {
                        var mtmp2 Tuple2_6string_4bool = rounded_float_digits(digits__1, count__1)
                        var x3 string = mtmp2._0
                        var x4 bool = mtmp2._1
                        var rounded__1 string = trim_float_digits(x3)
                        var jp8 int
                        if x4 {
                            jp8 = 1
                        } else {
                            jp8 = 0
                        }
                        var t25 int = point__1 + jp8
                        var candidate__1 string = fixed_float_text(rounded__1, t25, negative__0)
                        var mtmp3 Tuple2_4bool_6uint64 = parsed_float_bits(candidate__1, mantissa_bits__0, exponent_bias__0)
                        var x5 uint64 = mtmp3._1
                        var t26 bool = x5 == bits__0
                        if t26 {
                            return candidate__1
                        } else {
                            var compound_old1 int = count__1
                            var compound_value1 int = 1
                            var t27 int = compound_old1 + compound_value1
                            count__1 = t27
                            continue
                        }
                    } else {
                        break Loop_loop2
                    }
                }
                jp2 = point__1
                var t10 string = float_natural_decimal(exact_value__0)
                var t11 string = fixed_float_text(t10, jp2, negative__0)
                return t11
            }
        }
    }
}

func float_natural_from_u64(value__0 uint64) FloatNatural {
    var result__0 FloatNatural
    var inline2 *_goml_vec_uint32 = vec_new__Vec_6uint32()
    var inline3 FloatNatural = FloatNatural{
        words: inline2,
    }
    result__0 = inline3
    var t0 bool = value__0 != 0
    if t0 {
        var t1 *_goml_vec_uint32 = result__0.words
        var t2 uint32 = uint32(uint64(value__0))
        vec_push__Vec_6uint32(t1, t2)
        var t3 uint64 = value__0 >> 32
        var high__0 uint32 = uint32(uint64(t3))
        var t4 bool = high__0 != 0
        if t4 {
            var t5 *_goml_vec_uint32 = result__0.words
            vec_push__Vec_6uint32(t5, high__0)
        } else {}
    } else {}
    return result__0
}

func float_natural_shift_left(value__0 FloatNatural, bits__0 int) FloatNatural {
    var t0 bool
    var inline9 *_goml_vec_uint32 = value__0.words
    var inline10 bool = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_is__empty____T__u32(inline9)
    t0 = inline10
    if t0 {
        var inline7 *_goml_vec_uint32 = vec_new__Vec_6uint32()
        var inline8 FloatNatural = FloatNatural{
            words: inline7,
        }
        return inline8
    } else {
        var t19 bool = bits__0 == 0
        if t19 {
            var t20 FloatNatural = float_natural_copy(value__0)
            return t20
        } else {
            var result__0 FloatNatural
            var inline5 *_goml_vec_uint32 = vec_new__Vec_6uint32()
            var inline6 FloatNatural = FloatNatural{
                words: inline5,
            }
            result__0 = inline6
            var word_shift__0 int = bits__0 / 32
            var bit_shift__0 int = bits__0 % 32
            var index__0 int = 0
            Loop_loop0:
            for {
                var t15 bool = index__0 < word_shift__0
                if t15 {
                    var t16 *_goml_vec_uint32 = result__0.words
                    var inline3 uint32 = 0
                    vec_push__Vec_6uint32(t16, inline3)
                    var compound_old1 int = index__0
                    var compound_value1 int = 1
                    var t17 int = compound_old1 + compound_value1
                    index__0 = t17
                    continue
                } else {
                    break Loop_loop0
                }
            }
            var carry__0 uint64 = 0
            index__0 = 0
            Loop_loop1:
            for {
                var t4 *_goml_vec_uint32 = value__0.words
                var t5 int
                var inline2 int = vec_len__Vec_6uint32(t4)
                t5 = inline2
                var t6 bool = index__0 < t5
                if t6 {
                    var t7 *_goml_vec_uint32 = value__0.words
                    var word__0 uint32 = vec_get__Vec_6uint32(t7, index__0)
                    var t8 uint64 = uint64(uint32(word__0))
                    var t9 uint64 = t8 << bit_shift__0
                    var shifted__0 uint64 = t9 | carry__0
                    var t10 *_goml_vec_uint32 = result__0.words
                    var t11 uint32 = uint32(uint64(shifted__0))
                    vec_push__Vec_6uint32(t10, t11)
                    var t12 uint64 = shifted__0 >> 32
                    carry__0 = t12
                    var compound_old0 int = index__0
                    var compound_value0 int = 1
                    var t13 int = compound_old0 + compound_value0
                    index__0 = t13
                    continue
                } else {
                    break Loop_loop1
                }
            }
            var t1 bool = carry__0 != 0
            if t1 {
                var t2 *_goml_vec_uint32 = result__0.words
                var t3 uint32 = uint32(uint64(carry__0))
                vec_push__Vec_6uint32(t2, t3)
            } else {}
            return result__0
        }
    }
}

func float_natural_decimal(value__0 FloatNatural) string {
    var t0 bool
    var inline7 *_goml_vec_uint32 = value__0.words
    var inline8 bool = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_is__empty____T__u32(inline7)
    t0 = inline8
    if t0 {
        return "0"
    } else {
        var current__0 FloatNatural = float_natural_copy(value__0)
        var reversed__0 *_goml_vec_uint8 = vec_new__Vec_5uint8()
        Loop_loop0:
        for {
            var t10 bool
            var inline5 *_goml_vec_uint32 = current__0.words
            var inline6 bool = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_is__empty____T__u32(inline5)
            t10 = inline6
            var t11 bool = !t10
            if t11 {
                var t12 uint32 = float_natural_divide_small(current__0, 10)
                var t13 uint8 = uint8(uint32(t12))
                var t14 uint8 = t13 + 48
                vec_push__Vec_5uint8(reversed__0, t14)
                continue
            } else {
                break Loop_loop0
            }
        }
        var t1 int
        var inline3 int = vec_len__Vec_5uint8(reversed__0)
        t1 = inline3
        var output__0 *_goml_vec_uint8 = vec_with_capacity__Vec_5uint8(t1)
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
                vec_push__Vec_5uint8(output__0, t7)
                var compound_old0 int = offset__0
                var compound_value0 int = 1
                var t8 int = compound_old0 + compound_value0
                offset__0 = t8
                continue
            } else {
                break Loop_loop1
            }
        }
        var mtmp0 Tuple2_4bool_6string = _goml_runtime_core_string_from_utf8(output__0)
        var x0 string = mtmp0._1
        return x0
    }
}

func _goml_m_inherent_i_string_i_string_i_byte__len(self__0 string) int {
    var t0 int = _goml_runtime_core_string_len(self__0)
    return t0
}

func rounded_float_digits(exact__0 string, count__0 int) Tuple2_6string_4bool {
    var t0 int = count__0 + 1
    var output__0 *_goml_vec_uint8 = vec_with_capacity__Vec_5uint8(t0)
    var index__0 int = 0
    Loop_loop0:
    for {
        var t37 bool = index__0 < count__0
        if t37 {
            var t38 uint8
            var inline12 uint8 = _goml_runtime_core_string_byte_get(exact__0, index__0)
            t38 = inline12
            vec_push__Vec_5uint8(output__0, t38)
            var compound_old3 int = index__0
            var compound_value3 int = 1
            var t39 int = compound_old3 + compound_value3
            index__0 = t39
            continue
        } else {
            break Loop_loop0
        }
    }
    var t1 int
    var inline10 int = _goml_runtime_core_string_len(exact__0)
    t1 = inline10
    var t2 bool = count__0 == t1
    if t2 {
        var mtmp3 Tuple2_4bool_6string = _goml_runtime_core_string_from_utf8(output__0)
        var x3 string = mtmp3._1
        var t36 Tuple2_6string_4bool = Tuple2_6string_4bool{
            _0: x3,
            _1: false,
        }
        return t36
    } else {
        var next__0 uint8
        var inline9 uint8 = _goml_runtime_core_string_byte_get(exact__0, count__0)
        next__0 = inline9
        var trailing__0 bool = false
        var t3 int = count__0 + 1
        index__0 = t3
        Loop_loop1:
        for {
            var t30 int
            var inline8 int = _goml_runtime_core_string_len(exact__0)
            t30 = inline8
            var t31 bool = index__0 < t30
            if t31 {
                var t32 uint8
                var inline7 uint8 = _goml_runtime_core_string_byte_get(exact__0, index__0)
                t32 = inline7
                var t33 bool = t32 != 48
                if t33 {
                    trailing__0 = true
                } else {}
                var compound_old2 int = index__0
                var compound_value2 int = 1
                var t34 int = compound_old2 + compound_value2
                index__0 = t34
                continue
            } else {
                break Loop_loop1
            }
        }
        var t4 bool = next__0 > 53
        var jp0 bool
        if t4 {
            jp0 = true
        } else {
            var t23 bool = next__0 == 53
            if t23 {
                if trailing__0 {
                    jp0 = true
                } else {
                    var t24 int
                    var inline6 int = vec_len__Vec_5uint8(output__0)
                    t24 = inline6
                    var t25 int = t24 - 1
                    var t26 uint8 = vec_get__Vec_5uint8(output__0, t25)
                    var t27 uint8 = t26 - 48
                    var t28 uint8 = t27 % 2
                    var t29 bool = t28 == 1
                    jp0 = t29
                }
            } else {
                jp0 = false
            }
        }
        if jp0 {
            var index__1 int
            var inline5 int = vec_len__Vec_5uint8(output__0)
            index__1 = inline5
            Loop_loop2:
            for {
                var t13 bool = index__1 > 0
                if t13 {
                    var compound_old1 int = index__1
                    var compound_value1 int = 1
                    var t14 int = compound_old1 - compound_value1
                    index__1 = t14
                    var t16 uint8 = vec_get__Vec_5uint8(output__0, index__1)
                    var t17 bool = t16 < 57
                    if t17 {
                        var index0 int = index__1
                        var place0 uint8 = vec_get__Vec_5uint8(output__0, index0)
                        var value0 uint8 = 1
                        var t18 uint8 = place0 + value0
                        vec_set__Vec_5uint8(output__0, index0, t18)
                        var mtmp1 Tuple2_4bool_6string = _goml_runtime_core_string_from_utf8(output__0)
                        var x1 string = mtmp1._1
                        var t20 Tuple2_6string_4bool = Tuple2_6string_4bool{
                            _0: x1,
                            _1: false,
                        }
                        return t20
                    } else {
                        var index1 int = index__1
                        vec_get__Vec_5uint8(output__0, index1)
                        var value1 uint8 = 48
                        vec_set__Vec_5uint8(output__0, index1, value1)
                        continue
                    }
                } else {
                    break Loop_loop2
                }
            }
            var t5 int
            var inline4 int = vec_len__Vec_5uint8(output__0)
            t5 = inline4
            var t6 int = t5 + 1
            var carried__0 *_goml_vec_uint8 = vec_with_capacity__Vec_5uint8(t6)
            var inline2 uint8 = 49
            vec_push__Vec_5uint8(carried__0, inline2)
            index__1 = 0
            Loop_loop3:
            for {
                var t8 int
                var inline1 int = vec_len__Vec_5uint8(output__0)
                t8 = inline1
                var t9 bool = index__1 < t8
                if t9 {
                    var t10 uint8 = vec_get__Vec_5uint8(output__0, index__1)
                    vec_push__Vec_5uint8(carried__0, t10)
                    var compound_old0 int = index__1
                    var compound_value0 int = 1
                    var t11 int = compound_old0 + compound_value0
                    index__1 = t11
                    continue
                } else {
                    break Loop_loop3
                }
            }
            var mtmp0 Tuple2_4bool_6string = _goml_runtime_core_string_from_utf8(carried__0)
            var x0 string = mtmp0._1
            var t7 Tuple2_6string_4bool = Tuple2_6string_4bool{
                _0: x0,
                _1: true,
            }
            return t7
        } else {
            var mtmp2 Tuple2_4bool_6string = _goml_runtime_core_string_from_utf8(output__0)
            var x2 string = mtmp2._1
            var t22 Tuple2_6string_4bool = Tuple2_6string_4bool{
                _0: x2,
                _1: false,
            }
            return t22
        }
    }
}

func trim_float_digits(value__0 string) string {
    var length__0 int
    var inline3 int = _goml_runtime_core_string_len(value__0)
    length__0 = inline3
    Loop_loop0:
    for {
        var t0 bool = length__0 > 1
        var jp0 bool
        if t0 {
            var t3 int = length__0 - 1
            var t4 uint8
            var inline2 uint8 = _goml_runtime_core_string_byte_get(value__0, t3)
            t4 = inline2
            var t5 bool = t4 == 48
            jp0 = t5
        } else {
            jp0 = false
        }
        if jp0 {
            var compound_old0 int = length__0
            var compound_value0 int = 1
            var t1 int = compound_old0 - compound_value0
            length__0 = t1
            continue
        } else {
            break Loop_loop0
        }
    }
    var inline0 int = 0
    var inline1 string = string_byte_slice(value__0, inline0, length__0)
    return inline1
}

func fixed_float_text(digits__0 string, decimal_point__0 int, negative__0 bool) string {
    var bytes__0 *_goml_vec_uint8 = vec_new__Vec_5uint8()
    if negative__0 {
        var inline22 uint8 = 45
        vec_push__Vec_5uint8(bytes__0, inline22)
    } else {}
    var t0 bool = decimal_point__0 <= 0
    if t0 {
        var inline7 uint8 = 48
        vec_push__Vec_5uint8(bytes__0, inline7)
        var inline5 uint8 = 46
        vec_push__Vec_5uint8(bytes__0, inline5)
        var index__0 int = 0
        var t6 int = 0 - decimal_point__0
        Loop_loop0:
        for {
            var t7 bool = index__0 < t6
            if t7 {
                var inline3 uint8 = 48
                vec_push__Vec_5uint8(bytes__0, inline3)
                var compound_old1 int = index__0
                var compound_value1 int = 1
                var t8 int = compound_old1 + compound_value1
                index__0 = t8
                continue
            } else {
                break Loop_loop0
            }
        }
        index__0 = 0
        Loop_loop1:
        for {
            var t1 int
            var inline2 int = _goml_runtime_core_string_len(digits__0)
            t1 = inline2
            var t2 bool = index__0 < t1
            if t2 {
                var t3 uint8
                var inline1 uint8 = _goml_runtime_core_string_byte_get(digits__0, index__0)
                t3 = inline1
                vec_push__Vec_5uint8(bytes__0, t3)
                var compound_old0 int = index__0
                var compound_value0 int = 1
                var t4 int = compound_old0 + compound_value0
                index__0 = t4
                continue
            } else {
                break Loop_loop1
            }
        }
        var mtmp0 Tuple2_4bool_6string = _goml_runtime_core_string_from_utf8(bytes__0)
        var x0 string = mtmp0._1
        return x0
    } else {
        var t10 int
        var inline21 int = _goml_runtime_core_string_len(digits__0)
        t10 = inline21
        var t11 bool = decimal_point__0 >= t10
        if t11 {
            var index__1 int = 0
            Loop_loop2:
            for {
                var t15 int
                var inline13 int = _goml_runtime_core_string_len(digits__0)
                t15 = inline13
                var t16 bool = index__1 < t15
                if t16 {
                    var t17 uint8
                    var inline12 uint8 = _goml_runtime_core_string_byte_get(digits__0, index__1)
                    t17 = inline12
                    vec_push__Vec_5uint8(bytes__0, t17)
                    var compound_old3 int = index__1
                    var compound_value3 int = 1
                    var t18 int = compound_old3 + compound_value3
                    index__1 = t18
                    continue
                } else {
                    break Loop_loop2
                }
            }
            Loop_loop3:
            for {
                var t12 bool = index__1 < decimal_point__0
                if t12 {
                    var inline9 uint8 = 48
                    vec_push__Vec_5uint8(bytes__0, inline9)
                    var compound_old2 int = index__1
                    var compound_value2 int = 1
                    var t13 int = compound_old2 + compound_value2
                    index__1 = t13
                    continue
                } else {
                    break Loop_loop3
                }
            }
            var mtmp0 Tuple2_4bool_6string = _goml_runtime_core_string_from_utf8(bytes__0)
            var x0 string = mtmp0._1
            return x0
        } else {
            var index__2 int = 0
            Loop_loop4:
            for {
                var t25 bool = index__2 < decimal_point__0
                if t25 {
                    var t26 uint8
                    var inline20 uint8 = _goml_runtime_core_string_byte_get(digits__0, index__2)
                    t26 = inline20
                    vec_push__Vec_5uint8(bytes__0, t26)
                    var compound_old5 int = index__2
                    var compound_value5 int = 1
                    var t27 int = compound_old5 + compound_value5
                    index__2 = t27
                    continue
                } else {
                    break Loop_loop4
                }
            }
            var inline17 uint8 = 46
            vec_push__Vec_5uint8(bytes__0, inline17)
            Loop_loop5:
            for {
                var t20 int
                var inline16 int = _goml_runtime_core_string_len(digits__0)
                t20 = inline16
                var t21 bool = index__2 < t20
                if t21 {
                    var t22 uint8
                    var inline15 uint8 = _goml_runtime_core_string_byte_get(digits__0, index__2)
                    t22 = inline15
                    vec_push__Vec_5uint8(bytes__0, t22)
                    var compound_old4 int = index__2
                    var compound_value4 int = 1
                    var t23 int = compound_old4 + compound_value4
                    index__2 = t23
                    continue
                } else {
                    break Loop_loop5
                }
            }
            var mtmp0 Tuple2_4bool_6string = _goml_runtime_core_string_from_utf8(bytes__0)
            var x0 string = mtmp0._1
            return x0
        }
    }
}

func parsed_float_bits(value__0 string, mantissa_bits__0 int, exponent_bias__0 int) Tuple2_4bool_6uint64 {
    var parsed__0 ParsedFloat = parse_float_text(value__0)
    var t0 bool = parsed__0.valid
    var t1 bool = !t0
    if t1 {
        var t58 Tuple2_4bool_6uint64 = Tuple2_4bool_6uint64{
            _0: false,
            _1: 0,
        }
        return t58
    } else {
        var t2 bool = parsed__0.negative
        var jp0 uint64
        if t2 {
            var t55 bool = mantissa_bits__0 == 23
            var jp10 int
            if t55 {
                jp10 = 8
            } else {
                jp10 = 11
            }
            var t56 int = mantissa_bits__0 + jp10
            var t57 uint64 = 1 << t56
            jp0 = t57
        } else {
            jp0 = 0
        }
        var t3 bool = mantissa_bits__0 == 23
        var jp1 int
        if t3 {
            jp1 = 8
        } else {
            jp1 = 11
        }
        var t4 uint64 = 1 << jp1
        var t5 uint64 = t4 - 1
        var exponent_mask__0 uint64 = t5 << mantissa_bits__0
        var t6 int = parsed__0.special
        var t7 bool = t6 == 1
        if t7 {
            var t42 uint64 = jp0 | exponent_mask__0
            var t43 Tuple2_4bool_6uint64 = Tuple2_4bool_6uint64{
                _0: true,
                _1: t42,
            }
            return t43
        } else {
            var t44 int = parsed__0.special
            var t45 bool = t44 == 2
            if t45 {
                var t46 int = mantissa_bits__0 - 1
                var t47 uint64 = 1 << t46
                var t48 uint64 = exponent_mask__0 | t47
                var t49 bool = mantissa_bits__0 == 52
                var jp9 uint64
                if t49 {
                    jp9 = 1
                } else {
                    jp9 = 0
                }
                var t50 uint64 = t48 | jp9
                var t51 Tuple2_4bool_6uint64 = Tuple2_4bool_6uint64{
                    _0: true,
                    _1: t50,
                }
                return t51
            } else {
                var t52 FloatNatural = parsed__0.numerator
                var t53 bool
                var inline3 *_goml_vec_uint32 = t52.words
                var inline4 bool = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_is__empty____T__u32(inline3)
                t53 = inline4
                if t53 {
                    var t54 Tuple2_4bool_6uint64 = Tuple2_4bool_6uint64{
                        _0: true,
                        _1: jp0,
                    }
                    return t54
                } else {
                    var t8 bool = parsed__0.hexadecimal
                    var t9 bool = !t8
                    if t9 {
                        var t33 int = parsed__0.significant_digits
                        var t34 int = parsed__0.decimal_exponent
                        var decimal_position__0 int = t33 + t34
                        var t35 bool = mantissa_bits__0 == 23
                        var jp7 int
                        if t35 {
                            jp7 = 40
                        } else {
                            jp7 = 310
                        }
                        var t36 bool = mantissa_bits__0 == 23
                        var jp8 int
                        if t36 {
                            jp8 = -46
                        } else {
                            jp8 = -325
                        }
                        var t37 bool = decimal_position__0 > jp7
                        if t37 {
                            var t38 uint64 = jp0 | exponent_mask__0
                            var t39 Tuple2_4bool_6uint64 = Tuple2_4bool_6uint64{
                                _0: false,
                                _1: t38,
                            }
                            return t39
                        } else {
                            var t40 bool = decimal_position__0 < jp8
                            if t40 {
                                var t41 Tuple2_4bool_6uint64 = Tuple2_4bool_6uint64{
                                    _0: true,
                                    _1: jp0,
                                }
                                return t41
                            } else {
                                var t10 bool = parsed__0.hexadecimal
                                var t11 bool = !t10
                                var jp2 bool
                                if t11 {
                                    var t31 int = parsed__0.decimal_exponent
                                    var t32 bool = t31 < 0
                                    jp2 = t32
                                } else {
                                    jp2 = false
                                }
                                var jp3 FloatNatural
                                if jp2 {
                                    var t28 int = parsed__0.decimal_exponent
                                    var t29 int = 0 - t28
                                    var t30 FloatNatural = float_natural_power5(t29)
                                    jp3 = t30
                                } else {
                                    var inline0 *_goml_vec_uint32 = vec_new__Vec_6uint32()
                                    vec_push__Vec_6uint32(inline0, 1)
                                    var inline2 FloatNatural = FloatNatural{
                                        words: inline0,
                                    }
                                    jp3 = inline2
                                }
                                var t12 bool = parsed__0.hexadecimal
                                var t13 bool = !t12
                                var jp4 bool
                                if t13 {
                                    var t26 int = parsed__0.decimal_exponent
                                    var t27 bool = t26 > 0
                                    jp4 = t27
                                } else {
                                    jp4 = false
                                }
                                var jp5 FloatNatural
                                if jp4 {
                                    var t20 FloatNatural = parsed__0.numerator
                                    var result__0 FloatNatural = float_natural_copy(t20)
                                    var count__0 int = 0
                                    Loop_loop0:
                                    for {
                                        var t21 int = parsed__0.decimal_exponent
                                        var t22 bool = count__0 < t21
                                        if t22 {
                                            float_natural_multiply_small(result__0, 5)
                                            var compound_old0 int = count__0
                                            var compound_value0 int = 1
                                            var t23 int = compound_old0 + compound_value0
                                            count__0 = t23
                                            continue
                                        } else {
                                            break Loop_loop0
                                        }
                                    }
                                    jp5 = result__0
                                    var t14 bool = parsed__0.hexadecimal
                                    var jp6 int
                                    if t14 {
                                        var t18 int = parsed__0.binary_exponent
                                        jp6 = t18
                                    } else {
                                        var t19 int = parsed__0.decimal_exponent
                                        jp6 = t19
                                    }
                                    var mtmp0 Tuple2_6uint64_4bool = float_rational_bits(jp5, jp3, jp6, mantissa_bits__0, exponent_bias__0)
                                    var x0 uint64 = mtmp0._0
                                    var x1 bool = mtmp0._1
                                    var t15 bool = !x1
                                    var t16 uint64 = jp0 | x0
                                    var t17 Tuple2_4bool_6uint64 = Tuple2_4bool_6uint64{
                                        _0: t15,
                                        _1: t16,
                                    }
                                    return t17
                                } else {
                                    var t25 FloatNatural = parsed__0.numerator
                                    jp5 = t25
                                    var t14 bool = parsed__0.hexadecimal
                                    var jp6 int
                                    if t14 {
                                        var t18 int = parsed__0.binary_exponent
                                        jp6 = t18
                                    } else {
                                        var t19 int = parsed__0.decimal_exponent
                                        jp6 = t19
                                    }
                                    var mtmp0 Tuple2_6uint64_4bool = float_rational_bits(jp5, jp3, jp6, mantissa_bits__0, exponent_bias__0)
                                    var x0 uint64 = mtmp0._0
                                    var x1 bool = mtmp0._1
                                    var t15 bool = !x1
                                    var t16 uint64 = jp0 | x0
                                    var t17 Tuple2_4bool_6uint64 = Tuple2_4bool_6uint64{
                                        _0: t15,
                                        _1: t16,
                                    }
                                    return t17
                                }
                            }
                        }
                    } else {
                        var t10 bool = parsed__0.hexadecimal
                        var t11 bool = !t10
                        var jp2 bool
                        if t11 {
                            var t31 int = parsed__0.decimal_exponent
                            var t32 bool = t31 < 0
                            jp2 = t32
                        } else {
                            jp2 = false
                        }
                        var jp3 FloatNatural
                        if jp2 {
                            var t28 int = parsed__0.decimal_exponent
                            var t29 int = 0 - t28
                            var t30 FloatNatural = float_natural_power5(t29)
                            jp3 = t30
                        } else {
                            var inline0 *_goml_vec_uint32 = vec_new__Vec_6uint32()
                            vec_push__Vec_6uint32(inline0, 1)
                            var inline2 FloatNatural = FloatNatural{
                                words: inline0,
                            }
                            jp3 = inline2
                        }
                        var t12 bool = parsed__0.hexadecimal
                        var t13 bool = !t12
                        var jp4 bool
                        if t13 {
                            var t26 int = parsed__0.decimal_exponent
                            var t27 bool = t26 > 0
                            jp4 = t27
                        } else {
                            jp4 = false
                        }
                        var jp5 FloatNatural
                        if jp4 {
                            var t20 FloatNatural = parsed__0.numerator
                            var result__0 FloatNatural = float_natural_copy(t20)
                            var count__0 int = 0
                            Loop_loop0__2:
                            for {
                                var t21 int = parsed__0.decimal_exponent
                                var t22 bool = count__0 < t21
                                if t22 {
                                    float_natural_multiply_small(result__0, 5)
                                    var compound_old0 int = count__0
                                    var compound_value0 int = 1
                                    var t23 int = compound_old0 + compound_value0
                                    count__0 = t23
                                    continue
                                } else {
                                    break Loop_loop0__2
                                }
                            }
                            jp5 = result__0
                            var t14 bool = parsed__0.hexadecimal
                            var jp6 int
                            if t14 {
                                var t18 int = parsed__0.binary_exponent
                                jp6 = t18
                            } else {
                                var t19 int = parsed__0.decimal_exponent
                                jp6 = t19
                            }
                            var mtmp0 Tuple2_6uint64_4bool = float_rational_bits(jp5, jp3, jp6, mantissa_bits__0, exponent_bias__0)
                            var x0 uint64 = mtmp0._0
                            var x1 bool = mtmp0._1
                            var t15 bool = !x1
                            var t16 uint64 = jp0 | x0
                            var t17 Tuple2_4bool_6uint64 = Tuple2_4bool_6uint64{
                                _0: t15,
                                _1: t16,
                            }
                            return t17
                        } else {
                            var t25 FloatNatural = parsed__0.numerator
                            jp5 = t25
                            var t14 bool = parsed__0.hexadecimal
                            var jp6 int
                            if t14 {
                                var t18 int = parsed__0.binary_exponent
                                jp6 = t18
                            } else {
                                var t19 int = parsed__0.decimal_exponent
                                jp6 = t19
                            }
                            var mtmp0 Tuple2_6uint64_4bool = float_rational_bits(jp5, jp3, jp6, mantissa_bits__0, exponent_bias__0)
                            var x0 uint64 = mtmp0._0
                            var x1 bool = mtmp0._1
                            var t15 bool = !x1
                            var t16 uint64 = jp0 | x0
                            var t17 Tuple2_4bool_6uint64 = Tuple2_4bool_6uint64{
                                _0: t15,
                                _1: t16,
                            }
                            return t17
                        }
                    }
                }
            }
        }
    }
}

func float_natural_multiply_small(value__0 FloatNatural, factor__0 uint32) struct{} {
    var t0 bool = factor__0 == 0
    if t0 {
        var t16 *_goml_vec_uint32 = value__0.words
        _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_truncate____T__u32(t16, 0)
        return struct{}{}
    } else {
        var carry__0 uint64 = 0
        var index__0 int = 0
        var t4 uint64 = uint64(uint32(factor__0))
        Loop_loop0:
        for {
            var t5 *_goml_vec_uint32 = value__0.words
            var t6 int
            var inline1 int = vec_len__Vec_6uint32(t5)
            t6 = inline1
            var t7 bool = index__0 < t6
            if t7 {
                var t8 *_goml_vec_uint32 = value__0.words
                var t9 uint32 = vec_get__Vec_6uint32(t8, index__0)
                var t10 uint64 = uint64(uint32(t9))
                var t11 uint64 = t10 * t4
                var product__0 uint64 = t11 + carry__0
                var place0 *_goml_vec_uint32 = value__0.words
                var index0 int = index__0
                vec_get__Vec_6uint32(place0, index0)
                var value0 uint32 = uint32(uint64(product__0))
                vec_set__Vec_6uint32(place0, index0, value0)
                var t13 uint64 = product__0 >> 32
                carry__0 = t13
                var compound_old0 int = index__0
                var compound_value0 int = 1
                var t14 int = compound_old0 + compound_value0
                index__0 = t14
                continue
            } else {
                break Loop_loop0
            }
        }
        var t1 bool = carry__0 != 0
        if t1 {
            var t2 *_goml_vec_uint32 = value__0.words
            var t3 uint32 = uint32(uint64(carry__0))
            vec_push__Vec_6uint32(t2, t3)
            return struct{}{}
        } else {
            return struct{}{}
        }
    }
}

func float_natural_zero() FloatNatural {
    var t0 *_goml_vec_uint32 = vec_new__Vec_6uint32()
    var t1 FloatNatural = FloatNatural{
        words: t0,
    }
    return t1
}

func float_natural_copy(value__0 FloatNatural) FloatNatural {
    var result__0 FloatNatural
    var inline2 *_goml_vec_uint32 = vec_new__Vec_6uint32()
    var inline3 FloatNatural = FloatNatural{
        words: inline2,
    }
    result__0 = inline3
    var index__0 int = 0
    Loop_loop0:
    for {
        var t0 *_goml_vec_uint32 = value__0.words
        var t1 int
        var inline1 int = vec_len__Vec_6uint32(t0)
        t1 = inline1
        var t2 bool = index__0 < t1
        if t2 {
            var t3 *_goml_vec_uint32 = result__0.words
            var t4 *_goml_vec_uint32 = value__0.words
            var t5 uint32 = vec_get__Vec_6uint32(t4, index__0)
            vec_push__Vec_6uint32(t3, t5)
            var compound_old0 int = index__0
            var compound_value0 int = 1
            var t6 int = compound_old0 + compound_value0
            index__0 = t6
            continue
        } else {
            break Loop_loop0
        }
    }
    return result__0
}

func float_natural_divide_small(value__0 FloatNatural, divisor__0 uint32) uint32 {
    var remainder__0 uint64 = 0
    var t0 *_goml_vec_uint32 = value__0.words
    var index__0 int
    var inline0 int = vec_len__Vec_6uint32(t0)
    index__0 = inline0
    var t2 uint64 = uint64(uint32(divisor__0))
    var t3 uint64 = uint64(uint32(divisor__0))
    Loop_loop0:
    for {
        var t4 bool = index__0 > 0
        if t4 {
            var compound_old0 int = index__0
            var compound_value0 int = 1
            var t5 int = compound_old0 - compound_value0
            index__0 = t5
            var t7 uint64 = remainder__0 << 32
            var t8 *_goml_vec_uint32 = value__0.words
            var t9 uint32 = vec_get__Vec_6uint32(t8, index__0)
            var t10 uint64 = uint64(uint32(t9))
            var current__0 uint64 = t7 | t10
            var place0 *_goml_vec_uint32 = value__0.words
            var index0 int = index__0
            vec_get__Vec_6uint32(place0, index0)
            var t11 uint64 = current__0 / t2
            var value0 uint32 = uint32(uint64(t11))
            vec_set__Vec_6uint32(place0, index0, value0)
            var t13 uint64 = current__0 % t3
            remainder__0 = t13
            continue
        } else {
            break Loop_loop0
        }
    }
    float_natural_trim(value__0)
    var t1 uint32 = uint32(uint64(remainder__0))
    return t1
}

func _goml_m_inherent_i_string_i_string_i_byte__get(self__0 string, index__0 int) uint8 {
    var t0 uint8 = _goml_runtime_core_string_byte_get(self__0, index__0)
    return t0
}

func _goml_m_inherent_i_string_i_string_i_byte__slice(self__0 string, start__0 int, end__0 int) string {
    var inline0 bool = string_is_char_boundary(self__0, start__0)
    var inline1 bool
    if inline0 {
        var inline4 bool = string_is_char_boundary(self__0, end__0)
        inline1 = inline4
    } else {
        inline1 = false
    }
    if inline1 {
        var inline2 string = _goml_runtime_core_string_byte_slice(self__0, start__0, end__0)
        return inline2
    } else {
        var inline3 string = _goml_runtime_core_string_byte_slice(self__0, -1, -1)
        return inline3
    }
}

func parse_float_text(value__0 string) ParsedFloat {
    var t0 bool = string_equals_ascii_case(value__0, "nan")
    if t0 {
        var t118 FloatNatural
        var inline24 *_goml_vec_uint32 = vec_new__Vec_6uint32()
        var inline25 FloatNatural = FloatNatural{
            words: inline24,
        }
        t118 = inline25
        var t119 ParsedFloat = ParsedFloat{
            valid: true,
            negative: false,
            special: 2,
            numerator: t118,
            decimal_exponent: 0,
            binary_exponent: 0,
            hexadecimal: false,
            significant_digits: 0,
        }
        return t119
    } else {
        var index__0 int = 0
        var negative__0 bool = false
        var t1 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
        var t2 bool = index__0 < t1
        var jp0 bool
        if t2 {
            var t114 uint8
            var inline23 uint8 = _goml_runtime_core_string_byte_get(value__0, index__0)
            t114 = inline23
            var t115 bool = t114 == 43
            if t115 {
                jp0 = true
            } else {
                var t116 uint8
                var inline22 uint8 = _goml_runtime_core_string_byte_get(value__0, index__0)
                t116 = inline22
                var t117 bool = t116 == 45
                jp0 = t117
            }
        } else {
            jp0 = false
        }
        if jp0 {
            var t110 uint8
            var inline21 uint8 = _goml_runtime_core_string_byte_get(value__0, index__0)
            t110 = inline21
            var t111 bool = t110 == 45
            negative__0 = t111
            var compound_old10 int = index__0
            var compound_value10 int = 1
            var t112 int = compound_old10 + compound_value10
            index__0 = t112
        } else {}
        var t3 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
        var special_text__0 string = _goml_m_inherent_i_string_i_string_i_byte__slice(value__0, index__0, t3)
        var t4 bool = string_equals_ascii_case(special_text__0, "inf")
        var jp1 bool
        if t4 {
            jp1 = true
        } else {
            var t109 bool = string_equals_ascii_case(special_text__0, "infinity")
            jp1 = t109
        }
        if jp1 {
            var t107 FloatNatural
            var inline19 *_goml_vec_uint32 = vec_new__Vec_6uint32()
            var inline20 FloatNatural = FloatNatural{
                words: inline19,
            }
            t107 = inline20
            var t108 ParsedFloat = ParsedFloat{
                valid: true,
                negative: negative__0,
                special: 1,
                numerator: t107,
                decimal_exponent: 0,
                binary_exponent: 0,
                hexadecimal: false,
                significant_digits: 0,
            }
            return t108
        } else {
            var t5 int = index__0 + 2
            var t6 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
            var t7 bool = t5 <= t6
            var jp2 bool
            if t7 {
                var t105 uint8
                var inline18 uint8 = _goml_runtime_core_string_byte_get(value__0, index__0)
                t105 = inline18
                var t106 bool = t105 == 48
                jp2 = t106
            } else {
                jp2 = false
            }
            var jp3 bool
            if jp2 {
                var t101 int = index__0 + 1
                var t102 uint8
                var inline17 uint8 = _goml_runtime_core_string_byte_get(value__0, t101)
                t102 = inline17
                var t103 uint8
                var inline12 bool = t102 >= 65
                var inline13 bool
                if inline12 {
                    var inline16 bool = t102 <= 90
                    inline13 = inline16
                } else {
                    inline13 = false
                }
                if inline13 {
                    var inline14_lhs uint8 = 97
                    var inline14_rhs uint8 = 65
                    var inline14 uint8 = inline14_lhs - inline14_rhs
                    var inline15 uint8 = t102 + inline14
                    t103 = inline15
                    var t104 bool = t103 == 120
                    jp3 = t104
                    if jp3 {
                        var compound_old9 int = index__0
                        var compound_value9 int = 2
                        var t99 int = compound_old9 + compound_value9
                        index__0 = t99
                    } else {}
                    var mantissa_start__0 int = index__0
                    var jp4 int
                    if jp3 {
                        jp4 = 16
                    } else {
                        jp4 = 10
                    }
                    var numerator__0 FloatNatural = float_natural_zero()
                    var saw_digit__0 bool = false
                    var saw_dot__0 bool = false
                    var fraction_digits__0 int = 0
                    var significant_digits__0 int = 0
                    var previous_digit__0 bool = false
                    var t70 uint32 = uint32(int(jp4))
                    Loop_loop0:
                    for {
                        var t71 int
                        var inline11 int = _goml_runtime_core_string_len(value__0)
                        t71 = inline11
                        var t72 bool = index__0 < t71
                        if t72 {
                            var current__1 uint8
                            var inline10 uint8 = _goml_runtime_core_string_byte_get(value__0, index__0)
                            current__1 = inline10
                            var mtmp0 Tuple2_4bool_3int = float_digit(current__1, jp4)
                            var x0 bool = mtmp0._0
                            var x1 int = mtmp0._1
                            if x0 {
                                float_natural_multiply_small(numerator__0, t70)
                                var t73 uint32 = uint32(int(x1))
                                float_natural_add_small(numerator__0, t73)
                                saw_digit__0 = true
                                previous_digit__0 = true
                                if saw_dot__0 {
                                    var compound_old6 int = fraction_digits__0
                                    var compound_value6 int = 1
                                    var t80 int = compound_old6 + compound_value6
                                    fraction_digits__0 = t80
                                } else {}
                                var t74 bool = significant_digits__0 > 0
                                var jp16 bool
                                if t74 {
                                    jp16 = true
                                } else {
                                    var t79 bool = x1 != 0
                                    jp16 = t79
                                }
                                if jp16 {
                                    var compound_old5 int = significant_digits__0
                                    var compound_value5 int = 1
                                    var t77 int = compound_old5 + compound_value5
                                    significant_digits__0 = t77
                                } else {}
                                var compound_old4 int = index__0
                                var compound_value4 int = 1
                                var t75 int = compound_old4 + compound_value4
                                index__0 = t75
                                continue
                            } else {
                                var t82 bool = current__1 == 95
                                if t82 {
                                    var t83 int = index__0 + 1
                                    var t84 int
                                    var inline9 int = _goml_runtime_core_string_len(value__0)
                                    t84 = inline9
                                    var t85 bool = t83 >= t84
                                    if t85 {
                                        var inline7 FloatNatural = float_natural_zero()
                                        var inline8 ParsedFloat = ParsedFloat{
                                            valid: false,
                                            negative: false,
                                            special: 0,
                                            numerator: inline7,
                                            decimal_exponent: 0,
                                            binary_exponent: 0,
                                            hexadecimal: false,
                                            significant_digits: 0,
                                        }
                                        return inline8
                                    } else {
                                        var t86 int = index__0 + 1
                                        var t87 uint8
                                        var inline6 uint8 = _goml_runtime_core_string_byte_get(value__0, t86)
                                        t87 = inline6
                                        var mtmp1 Tuple2_4bool_3int = float_digit(t87, jp4)
                                        var x2 bool = mtmp1._0
                                        var jp17 bool
                                        if jp3 {
                                            var t94 bool = !saw_digit__0
                                            jp17 = t94
                                        } else {
                                            jp17 = false
                                        }
                                        var jp18 bool
                                        if jp17 {
                                            var t93 bool = index__0 == mantissa_start__0
                                            jp18 = t93
                                        } else {
                                            jp18 = false
                                        }
                                        var t88 bool = !previous_digit__0
                                        var jp19 bool
                                        if t88 {
                                            var t92 bool = !jp18
                                            jp19 = t92
                                        } else {
                                            jp19 = false
                                        }
                                        var jp20 bool
                                        if jp19 {
                                            jp20 = true
                                        } else {
                                            var t91 bool = !x2
                                            jp20 = t91
                                        }
                                        if jp20 {
                                            var inline4 FloatNatural = float_natural_zero()
                                            var inline5 ParsedFloat = ParsedFloat{
                                                valid: false,
                                                negative: false,
                                                special: 0,
                                                numerator: inline4,
                                                decimal_exponent: 0,
                                                binary_exponent: 0,
                                                hexadecimal: false,
                                                significant_digits: 0,
                                            }
                                            return inline5
                                        } else {
                                            previous_digit__0 = false
                                            var compound_old7 int = index__0
                                            var compound_value7 int = 1
                                            var t89 int = compound_old7 + compound_value7
                                            index__0 = t89
                                            continue
                                        }
                                    }
                                } else {
                                    var t95 bool = current__1 == 46
                                    var jp21 bool
                                    if t95 {
                                        var t98 bool = !saw_dot__0
                                        jp21 = t98
                                    } else {
                                        jp21 = false
                                    }
                                    if jp21 {
                                        saw_dot__0 = true
                                        previous_digit__0 = false
                                        var compound_old8 int = index__0
                                        var compound_value8 int = 1
                                        var t96 int = compound_old8 + compound_value8
                                        index__0 = t96
                                        continue
                                    } else {
                                        break Loop_loop0
                                    }
                                }
                            }
                        } else {
                            break Loop_loop0
                        }
                    }
                    var t8 bool = !saw_digit__0
                    if t8 {
                        var inline2 FloatNatural = float_natural_zero()
                        var inline3 ParsedFloat = ParsedFloat{
                            valid: false,
                            negative: false,
                            special: 0,
                            numerator: inline2,
                            decimal_exponent: 0,
                            binary_exponent: 0,
                            hexadecimal: false,
                            significant_digits: 0,
                        }
                        return inline3
                    } else {
                        var jp5 uint8
                        if jp3 {
                            jp5 = 112
                        } else {
                            jp5 = 101
                        }
                        var t9 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
                        var t10_lhs int = 9223372036854775807
                        var t10_rhs int = 2048
                        var t10 int = t10_lhs - t10_rhs
                        var t11 int = t10 / 4
                        var t12 bool = t9 <= t11
                        var jp6 int
                        if t12 {
                            var t67 int
                            var inline1 int = _goml_runtime_core_string_len(value__0)
                            t67 = inline1
                            var t68 int = t67 * 4
                            var t69 int = t68 + 2048
                            jp6 = t69
                        } else {
                            jp6 = 9223372036854775807
                        }
                        var exponent__0 int = 0
                        var exponent_negative__0 bool = false
                        var t13 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
                        var t14 bool = index__0 < t13
                        var jp7 bool
                        if t14 {
                            var t64 uint8
                            var inline0 uint8 = _goml_runtime_core_string_byte_get(value__0, index__0)
                            t64 = inline0
                            var t65 uint8 = ascii_lower(t64)
                            var t66 bool = t65 == jp5
                            jp7 = t66
                        } else {
                            jp7 = false
                        }
                        if jp7 {
                            var compound_old0 int = index__0
                            var compound_value0 int = 1
                            var t23 int = compound_old0 + compound_value0
                            index__0 = t23
                            var t25 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
                            var t26 bool = index__0 < t25
                            var jp10 bool
                            if t26 {
                                var t59 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, index__0)
                                var t60 bool = t59 == 43
                                if t60 {
                                    jp10 = true
                                } else {
                                    var t61 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, index__0)
                                    var t62 bool = t61 == 45
                                    jp10 = t62
                                }
                            } else {
                                jp10 = false
                            }
                            if jp10 {
                                var t55 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, index__0)
                                var t56 bool = t55 == 45
                                exponent_negative__0 = t56
                                var compound_old3 int = index__0
                                var compound_value3 int = 1
                                var t57 int = compound_old3 + compound_value3
                                index__0 = t57
                            } else {}
                            var exponent_digits__0 bool = false
                            previous_digit__0 = false
                            Loop_loop1:
                            for {
                                var t29 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
                                var t30 bool = index__0 < t29
                                if t30 {
                                    var current__0 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, index__0)
                                    var t31 bool = current__0 >= 48
                                    var jp11 bool
                                    if t31 {
                                        var t54 bool = current__0 <= 57
                                        jp11 = t54
                                    } else {
                                        jp11 = false
                                    }
                                    if jp11 {
                                        exponent_digits__0 = true
                                        previous_digit__0 = true
                                        var t32 uint8 = current__0 - 48
                                        var digit__0 int = int(uint8(t32))
                                        var t33 int = jp6 - digit__0
                                        var t34 int = t33 / 10
                                        var t35 bool = exponent__0 > t34
                                        var jp12 int
                                        if t35 {
                                            jp12 = jp6
                                        } else {
                                            var t38 int = exponent__0 * 10
                                            var t39 int = t38 + digit__0
                                            jp12 = t39
                                        }
                                        exponent__0 = jp12
                                        var compound_old1 int = index__0
                                        var compound_value1 int = 1
                                        var t36 int = compound_old1 + compound_value1
                                        index__0 = t36
                                        continue
                                    } else {
                                        var t40 bool = current__0 == 95
                                        if t40 {
                                            var t41 bool = !previous_digit__0
                                            var jp13 bool
                                            if t41 {
                                                jp13 = true
                                            } else {
                                                var t51 int = index__0 + 1
                                                var t52 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
                                                var t53 bool = t51 >= t52
                                                jp13 = t53
                                            }
                                            var jp14 bool
                                            if jp13 {
                                                jp14 = true
                                            } else {
                                                var t48 int = index__0 + 1
                                                var t49 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, t48)
                                                var t50 bool = t49 < 48
                                                jp14 = t50
                                            }
                                            var jp15 bool
                                            if jp14 {
                                                jp15 = true
                                            } else {
                                                var t45 int = index__0 + 1
                                                var t46 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, t45)
                                                var t47 bool = t46 > 57
                                                jp15 = t47
                                            }
                                            if jp15 {
                                                var t44 ParsedFloat = invalid_parsed_float()
                                                return t44
                                            } else {
                                                previous_digit__0 = false
                                                var compound_old2 int = index__0
                                                var compound_value2 int = 1
                                                var t42 int = compound_old2 + compound_value2
                                                index__0 = t42
                                                continue
                                            }
                                        } else {
                                            break Loop_loop1
                                        }
                                    }
                                } else {
                                    break Loop_loop1
                                }
                            }
                            var t27 bool = !exponent_digits__0
                            if t27 {
                                var t28 ParsedFloat = invalid_parsed_float()
                                return t28
                            } else {
                                var t15 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
                                var t16 bool = index__0 != t15
                                if t16 {
                                    var t22 ParsedFloat = invalid_parsed_float()
                                    return t22
                                } else {
                                    if exponent_negative__0 {
                                        var t21 int = 0 - exponent__0
                                        exponent__0 = t21
                                    } else {}
                                    var jp8 int
                                    if jp3 {
                                        jp8 = 0
                                    } else {
                                        var t20 int = exponent__0 - fraction_digits__0
                                        jp8 = t20
                                    }
                                    var jp9 int
                                    if jp3 {
                                        var t18 int = fraction_digits__0 * 4
                                        var t19 int = exponent__0 - t18
                                        jp9 = t19
                                    } else {
                                        jp9 = 0
                                    }
                                    var t17 ParsedFloat = ParsedFloat{
                                        valid: true,
                                        negative: negative__0,
                                        special: 0,
                                        numerator: numerator__0,
                                        decimal_exponent: jp8,
                                        binary_exponent: jp9,
                                        hexadecimal: jp3,
                                        significant_digits: significant_digits__0,
                                    }
                                    return t17
                                }
                            }
                        } else {
                            if jp3 {
                                var t63 ParsedFloat = invalid_parsed_float()
                                return t63
                            } else {
                                var t15 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
                                var t16 bool = index__0 != t15
                                if t16 {
                                    var t22 ParsedFloat = invalid_parsed_float()
                                    return t22
                                } else {
                                    if exponent_negative__0 {
                                        var t21 int = 0 - exponent__0
                                        exponent__0 = t21
                                    } else {}
                                    var jp8 int
                                    if jp3 {
                                        jp8 = 0
                                    } else {
                                        var t20 int = exponent__0 - fraction_digits__0
                                        jp8 = t20
                                    }
                                    var jp9 int
                                    if jp3 {
                                        var t18 int = fraction_digits__0 * 4
                                        var t19 int = exponent__0 - t18
                                        jp9 = t19
                                    } else {
                                        jp9 = 0
                                    }
                                    var t17 ParsedFloat = ParsedFloat{
                                        valid: true,
                                        negative: negative__0,
                                        special: 0,
                                        numerator: numerator__0,
                                        decimal_exponent: jp8,
                                        binary_exponent: jp9,
                                        hexadecimal: jp3,
                                        significant_digits: significant_digits__0,
                                    }
                                    return t17
                                }
                            }
                        }
                    }
                } else {
                    t103 = t102
                    var t104 bool = t103 == 120
                    jp3 = t104
                    if jp3 {
                        var compound_old9 int = index__0
                        var compound_value9 int = 2
                        var t99 int = compound_old9 + compound_value9
                        index__0 = t99
                    } else {}
                    var mantissa_start__0 int = index__0
                    var jp4 int
                    if jp3 {
                        jp4 = 16
                    } else {
                        jp4 = 10
                    }
                    var numerator__0 FloatNatural = float_natural_zero()
                    var saw_digit__0 bool = false
                    var saw_dot__0 bool = false
                    var fraction_digits__0 int = 0
                    var significant_digits__0 int = 0
                    var previous_digit__0 bool = false
                    var t70 uint32 = uint32(int(jp4))
                    Loop_loop0__2:
                    for {
                        var t71 int
                        var inline11 int = _goml_runtime_core_string_len(value__0)
                        t71 = inline11
                        var t72 bool = index__0 < t71
                        if t72 {
                            var current__1 uint8
                            var inline10 uint8 = _goml_runtime_core_string_byte_get(value__0, index__0)
                            current__1 = inline10
                            var mtmp0 Tuple2_4bool_3int = float_digit(current__1, jp4)
                            var x0 bool = mtmp0._0
                            var x1 int = mtmp0._1
                            if x0 {
                                float_natural_multiply_small(numerator__0, t70)
                                var t73 uint32 = uint32(int(x1))
                                float_natural_add_small(numerator__0, t73)
                                saw_digit__0 = true
                                previous_digit__0 = true
                                if saw_dot__0 {
                                    var compound_old6 int = fraction_digits__0
                                    var compound_value6 int = 1
                                    var t80 int = compound_old6 + compound_value6
                                    fraction_digits__0 = t80
                                } else {}
                                var t74 bool = significant_digits__0 > 0
                                var jp16 bool
                                if t74 {
                                    jp16 = true
                                } else {
                                    var t79 bool = x1 != 0
                                    jp16 = t79
                                }
                                if jp16 {
                                    var compound_old5 int = significant_digits__0
                                    var compound_value5 int = 1
                                    var t77 int = compound_old5 + compound_value5
                                    significant_digits__0 = t77
                                } else {}
                                var compound_old4 int = index__0
                                var compound_value4 int = 1
                                var t75 int = compound_old4 + compound_value4
                                index__0 = t75
                                continue
                            } else {
                                var t82 bool = current__1 == 95
                                if t82 {
                                    var t83 int = index__0 + 1
                                    var t84 int
                                    var inline9 int = _goml_runtime_core_string_len(value__0)
                                    t84 = inline9
                                    var t85 bool = t83 >= t84
                                    if t85 {
                                        var inline7 FloatNatural = float_natural_zero()
                                        var inline8 ParsedFloat = ParsedFloat{
                                            valid: false,
                                            negative: false,
                                            special: 0,
                                            numerator: inline7,
                                            decimal_exponent: 0,
                                            binary_exponent: 0,
                                            hexadecimal: false,
                                            significant_digits: 0,
                                        }
                                        return inline8
                                    } else {
                                        var t86 int = index__0 + 1
                                        var t87 uint8
                                        var inline6 uint8 = _goml_runtime_core_string_byte_get(value__0, t86)
                                        t87 = inline6
                                        var mtmp1 Tuple2_4bool_3int = float_digit(t87, jp4)
                                        var x2 bool = mtmp1._0
                                        var jp17 bool
                                        if jp3 {
                                            var t94 bool = !saw_digit__0
                                            jp17 = t94
                                        } else {
                                            jp17 = false
                                        }
                                        var jp18 bool
                                        if jp17 {
                                            var t93 bool = index__0 == mantissa_start__0
                                            jp18 = t93
                                        } else {
                                            jp18 = false
                                        }
                                        var t88 bool = !previous_digit__0
                                        var jp19 bool
                                        if t88 {
                                            var t92 bool = !jp18
                                            jp19 = t92
                                        } else {
                                            jp19 = false
                                        }
                                        var jp20 bool
                                        if jp19 {
                                            jp20 = true
                                        } else {
                                            var t91 bool = !x2
                                            jp20 = t91
                                        }
                                        if jp20 {
                                            var inline4 FloatNatural = float_natural_zero()
                                            var inline5 ParsedFloat = ParsedFloat{
                                                valid: false,
                                                negative: false,
                                                special: 0,
                                                numerator: inline4,
                                                decimal_exponent: 0,
                                                binary_exponent: 0,
                                                hexadecimal: false,
                                                significant_digits: 0,
                                            }
                                            return inline5
                                        } else {
                                            previous_digit__0 = false
                                            var compound_old7 int = index__0
                                            var compound_value7 int = 1
                                            var t89 int = compound_old7 + compound_value7
                                            index__0 = t89
                                            continue
                                        }
                                    }
                                } else {
                                    var t95 bool = current__1 == 46
                                    var jp21 bool
                                    if t95 {
                                        var t98 bool = !saw_dot__0
                                        jp21 = t98
                                    } else {
                                        jp21 = false
                                    }
                                    if jp21 {
                                        saw_dot__0 = true
                                        previous_digit__0 = false
                                        var compound_old8 int = index__0
                                        var compound_value8 int = 1
                                        var t96 int = compound_old8 + compound_value8
                                        index__0 = t96
                                        continue
                                    } else {
                                        break Loop_loop0__2
                                    }
                                }
                            }
                        } else {
                            break Loop_loop0__2
                        }
                    }
                    var t8 bool = !saw_digit__0
                    if t8 {
                        var inline2 FloatNatural = float_natural_zero()
                        var inline3 ParsedFloat = ParsedFloat{
                            valid: false,
                            negative: false,
                            special: 0,
                            numerator: inline2,
                            decimal_exponent: 0,
                            binary_exponent: 0,
                            hexadecimal: false,
                            significant_digits: 0,
                        }
                        return inline3
                    } else {
                        var jp5 uint8
                        if jp3 {
                            jp5 = 112
                        } else {
                            jp5 = 101
                        }
                        var t9 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
                        var t10_lhs int = 9223372036854775807
                        var t10_rhs int = 2048
                        var t10 int = t10_lhs - t10_rhs
                        var t11 int = t10 / 4
                        var t12 bool = t9 <= t11
                        var jp6 int
                        if t12 {
                            var t67 int
                            var inline1 int = _goml_runtime_core_string_len(value__0)
                            t67 = inline1
                            var t68 int = t67 * 4
                            var t69 int = t68 + 2048
                            jp6 = t69
                        } else {
                            jp6 = 9223372036854775807
                        }
                        var exponent__0 int = 0
                        var exponent_negative__0 bool = false
                        var t13 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
                        var t14 bool = index__0 < t13
                        var jp7 bool
                        if t14 {
                            var t64 uint8
                            var inline0 uint8 = _goml_runtime_core_string_byte_get(value__0, index__0)
                            t64 = inline0
                            var t65 uint8 = ascii_lower(t64)
                            var t66 bool = t65 == jp5
                            jp7 = t66
                        } else {
                            jp7 = false
                        }
                        if jp7 {
                            var compound_old0 int = index__0
                            var compound_value0 int = 1
                            var t23 int = compound_old0 + compound_value0
                            index__0 = t23
                            var t25 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
                            var t26 bool = index__0 < t25
                            var jp10 bool
                            if t26 {
                                var t59 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, index__0)
                                var t60 bool = t59 == 43
                                if t60 {
                                    jp10 = true
                                } else {
                                    var t61 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, index__0)
                                    var t62 bool = t61 == 45
                                    jp10 = t62
                                }
                            } else {
                                jp10 = false
                            }
                            if jp10 {
                                var t55 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, index__0)
                                var t56 bool = t55 == 45
                                exponent_negative__0 = t56
                                var compound_old3 int = index__0
                                var compound_value3 int = 1
                                var t57 int = compound_old3 + compound_value3
                                index__0 = t57
                            } else {}
                            var exponent_digits__0 bool = false
                            previous_digit__0 = false
                            Loop_loop1__2:
                            for {
                                var t29 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
                                var t30 bool = index__0 < t29
                                if t30 {
                                    var current__0 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, index__0)
                                    var t31 bool = current__0 >= 48
                                    var jp11 bool
                                    if t31 {
                                        var t54 bool = current__0 <= 57
                                        jp11 = t54
                                    } else {
                                        jp11 = false
                                    }
                                    if jp11 {
                                        exponent_digits__0 = true
                                        previous_digit__0 = true
                                        var t32 uint8 = current__0 - 48
                                        var digit__0 int = int(uint8(t32))
                                        var t33 int = jp6 - digit__0
                                        var t34 int = t33 / 10
                                        var t35 bool = exponent__0 > t34
                                        var jp12 int
                                        if t35 {
                                            jp12 = jp6
                                        } else {
                                            var t38 int = exponent__0 * 10
                                            var t39 int = t38 + digit__0
                                            jp12 = t39
                                        }
                                        exponent__0 = jp12
                                        var compound_old1 int = index__0
                                        var compound_value1 int = 1
                                        var t36 int = compound_old1 + compound_value1
                                        index__0 = t36
                                        continue
                                    } else {
                                        var t40 bool = current__0 == 95
                                        if t40 {
                                            var t41 bool = !previous_digit__0
                                            var jp13 bool
                                            if t41 {
                                                jp13 = true
                                            } else {
                                                var t51 int = index__0 + 1
                                                var t52 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
                                                var t53 bool = t51 >= t52
                                                jp13 = t53
                                            }
                                            var jp14 bool
                                            if jp13 {
                                                jp14 = true
                                            } else {
                                                var t48 int = index__0 + 1
                                                var t49 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, t48)
                                                var t50 bool = t49 < 48
                                                jp14 = t50
                                            }
                                            var jp15 bool
                                            if jp14 {
                                                jp15 = true
                                            } else {
                                                var t45 int = index__0 + 1
                                                var t46 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, t45)
                                                var t47 bool = t46 > 57
                                                jp15 = t47
                                            }
                                            if jp15 {
                                                var t44 ParsedFloat = invalid_parsed_float()
                                                return t44
                                            } else {
                                                previous_digit__0 = false
                                                var compound_old2 int = index__0
                                                var compound_value2 int = 1
                                                var t42 int = compound_old2 + compound_value2
                                                index__0 = t42
                                                continue
                                            }
                                        } else {
                                            break Loop_loop1__2
                                        }
                                    }
                                } else {
                                    break Loop_loop1__2
                                }
                            }
                            var t27 bool = !exponent_digits__0
                            if t27 {
                                var t28 ParsedFloat = invalid_parsed_float()
                                return t28
                            } else {
                                var t15 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
                                var t16 bool = index__0 != t15
                                if t16 {
                                    var t22 ParsedFloat = invalid_parsed_float()
                                    return t22
                                } else {
                                    if exponent_negative__0 {
                                        var t21 int = 0 - exponent__0
                                        exponent__0 = t21
                                    } else {}
                                    var jp8 int
                                    if jp3 {
                                        jp8 = 0
                                    } else {
                                        var t20 int = exponent__0 - fraction_digits__0
                                        jp8 = t20
                                    }
                                    var jp9 int
                                    if jp3 {
                                        var t18 int = fraction_digits__0 * 4
                                        var t19 int = exponent__0 - t18
                                        jp9 = t19
                                    } else {
                                        jp9 = 0
                                    }
                                    var t17 ParsedFloat = ParsedFloat{
                                        valid: true,
                                        negative: negative__0,
                                        special: 0,
                                        numerator: numerator__0,
                                        decimal_exponent: jp8,
                                        binary_exponent: jp9,
                                        hexadecimal: jp3,
                                        significant_digits: significant_digits__0,
                                    }
                                    return t17
                                }
                            }
                        } else {
                            if jp3 {
                                var t63 ParsedFloat = invalid_parsed_float()
                                return t63
                            } else {
                                var t15 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
                                var t16 bool = index__0 != t15
                                if t16 {
                                    var t22 ParsedFloat = invalid_parsed_float()
                                    return t22
                                } else {
                                    if exponent_negative__0 {
                                        var t21 int = 0 - exponent__0
                                        exponent__0 = t21
                                    } else {}
                                    var jp8 int
                                    if jp3 {
                                        jp8 = 0
                                    } else {
                                        var t20 int = exponent__0 - fraction_digits__0
                                        jp8 = t20
                                    }
                                    var jp9 int
                                    if jp3 {
                                        var t18 int = fraction_digits__0 * 4
                                        var t19 int = exponent__0 - t18
                                        jp9 = t19
                                    } else {
                                        jp9 = 0
                                    }
                                    var t17 ParsedFloat = ParsedFloat{
                                        valid: true,
                                        negative: negative__0,
                                        special: 0,
                                        numerator: numerator__0,
                                        decimal_exponent: jp8,
                                        binary_exponent: jp9,
                                        hexadecimal: jp3,
                                        significant_digits: significant_digits__0,
                                    }
                                    return t17
                                }
                            }
                        }
                    }
                }
            } else {
                jp3 = false
                if jp3 {
                    var compound_old9 int = index__0
                    var compound_value9 int = 2
                    var t99 int = compound_old9 + compound_value9
                    index__0 = t99
                } else {}
                var mantissa_start__0 int = index__0
                var jp4 int
                if jp3 {
                    jp4 = 16
                } else {
                    jp4 = 10
                }
                var numerator__0 FloatNatural = float_natural_zero()
                var saw_digit__0 bool = false
                var saw_dot__0 bool = false
                var fraction_digits__0 int = 0
                var significant_digits__0 int = 0
                var previous_digit__0 bool = false
                var t70 uint32 = uint32(int(jp4))
                Loop_loop0__3:
                for {
                    var t71 int
                    var inline11 int = _goml_runtime_core_string_len(value__0)
                    t71 = inline11
                    var t72 bool = index__0 < t71
                    if t72 {
                        var current__1 uint8
                        var inline10 uint8 = _goml_runtime_core_string_byte_get(value__0, index__0)
                        current__1 = inline10
                        var mtmp0 Tuple2_4bool_3int = float_digit(current__1, jp4)
                        var x0 bool = mtmp0._0
                        var x1 int = mtmp0._1
                        if x0 {
                            float_natural_multiply_small(numerator__0, t70)
                            var t73 uint32 = uint32(int(x1))
                            float_natural_add_small(numerator__0, t73)
                            saw_digit__0 = true
                            previous_digit__0 = true
                            if saw_dot__0 {
                                var compound_old6 int = fraction_digits__0
                                var compound_value6 int = 1
                                var t80 int = compound_old6 + compound_value6
                                fraction_digits__0 = t80
                            } else {}
                            var t74 bool = significant_digits__0 > 0
                            var jp16 bool
                            if t74 {
                                jp16 = true
                            } else {
                                var t79 bool = x1 != 0
                                jp16 = t79
                            }
                            if jp16 {
                                var compound_old5 int = significant_digits__0
                                var compound_value5 int = 1
                                var t77 int = compound_old5 + compound_value5
                                significant_digits__0 = t77
                            } else {}
                            var compound_old4 int = index__0
                            var compound_value4 int = 1
                            var t75 int = compound_old4 + compound_value4
                            index__0 = t75
                            continue
                        } else {
                            var t82 bool = current__1 == 95
                            if t82 {
                                var t83 int = index__0 + 1
                                var t84 int
                                var inline9 int = _goml_runtime_core_string_len(value__0)
                                t84 = inline9
                                var t85 bool = t83 >= t84
                                if t85 {
                                    var inline7 FloatNatural = float_natural_zero()
                                    var inline8 ParsedFloat = ParsedFloat{
                                        valid: false,
                                        negative: false,
                                        special: 0,
                                        numerator: inline7,
                                        decimal_exponent: 0,
                                        binary_exponent: 0,
                                        hexadecimal: false,
                                        significant_digits: 0,
                                    }
                                    return inline8
                                } else {
                                    var t86 int = index__0 + 1
                                    var t87 uint8
                                    var inline6 uint8 = _goml_runtime_core_string_byte_get(value__0, t86)
                                    t87 = inline6
                                    var mtmp1 Tuple2_4bool_3int = float_digit(t87, jp4)
                                    var x2 bool = mtmp1._0
                                    var jp17 bool
                                    if jp3 {
                                        var t94 bool = !saw_digit__0
                                        jp17 = t94
                                    } else {
                                        jp17 = false
                                    }
                                    var jp18 bool
                                    if jp17 {
                                        var t93 bool = index__0 == mantissa_start__0
                                        jp18 = t93
                                    } else {
                                        jp18 = false
                                    }
                                    var t88 bool = !previous_digit__0
                                    var jp19 bool
                                    if t88 {
                                        var t92 bool = !jp18
                                        jp19 = t92
                                    } else {
                                        jp19 = false
                                    }
                                    var jp20 bool
                                    if jp19 {
                                        jp20 = true
                                    } else {
                                        var t91 bool = !x2
                                        jp20 = t91
                                    }
                                    if jp20 {
                                        var inline4 FloatNatural = float_natural_zero()
                                        var inline5 ParsedFloat = ParsedFloat{
                                            valid: false,
                                            negative: false,
                                            special: 0,
                                            numerator: inline4,
                                            decimal_exponent: 0,
                                            binary_exponent: 0,
                                            hexadecimal: false,
                                            significant_digits: 0,
                                        }
                                        return inline5
                                    } else {
                                        previous_digit__0 = false
                                        var compound_old7 int = index__0
                                        var compound_value7 int = 1
                                        var t89 int = compound_old7 + compound_value7
                                        index__0 = t89
                                        continue
                                    }
                                }
                            } else {
                                var t95 bool = current__1 == 46
                                var jp21 bool
                                if t95 {
                                    var t98 bool = !saw_dot__0
                                    jp21 = t98
                                } else {
                                    jp21 = false
                                }
                                if jp21 {
                                    saw_dot__0 = true
                                    previous_digit__0 = false
                                    var compound_old8 int = index__0
                                    var compound_value8 int = 1
                                    var t96 int = compound_old8 + compound_value8
                                    index__0 = t96
                                    continue
                                } else {
                                    break Loop_loop0__3
                                }
                            }
                        }
                    } else {
                        break Loop_loop0__3
                    }
                }
                var t8 bool = !saw_digit__0
                if t8 {
                    var inline2 FloatNatural = float_natural_zero()
                    var inline3 ParsedFloat = ParsedFloat{
                        valid: false,
                        negative: false,
                        special: 0,
                        numerator: inline2,
                        decimal_exponent: 0,
                        binary_exponent: 0,
                        hexadecimal: false,
                        significant_digits: 0,
                    }
                    return inline3
                } else {
                    var jp5 uint8
                    if jp3 {
                        jp5 = 112
                    } else {
                        jp5 = 101
                    }
                    var t9 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
                    var t10_lhs int = 9223372036854775807
                    var t10_rhs int = 2048
                    var t10 int = t10_lhs - t10_rhs
                    var t11 int = t10 / 4
                    var t12 bool = t9 <= t11
                    var jp6 int
                    if t12 {
                        var t67 int
                        var inline1 int = _goml_runtime_core_string_len(value__0)
                        t67 = inline1
                        var t68 int = t67 * 4
                        var t69 int = t68 + 2048
                        jp6 = t69
                    } else {
                        jp6 = 9223372036854775807
                    }
                    var exponent__0 int = 0
                    var exponent_negative__0 bool = false
                    var t13 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
                    var t14 bool = index__0 < t13
                    var jp7 bool
                    if t14 {
                        var t64 uint8
                        var inline0 uint8 = _goml_runtime_core_string_byte_get(value__0, index__0)
                        t64 = inline0
                        var t65 uint8 = ascii_lower(t64)
                        var t66 bool = t65 == jp5
                        jp7 = t66
                    } else {
                        jp7 = false
                    }
                    if jp7 {
                        var compound_old0 int = index__0
                        var compound_value0 int = 1
                        var t23 int = compound_old0 + compound_value0
                        index__0 = t23
                        var t25 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
                        var t26 bool = index__0 < t25
                        var jp10 bool
                        if t26 {
                            var t59 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, index__0)
                            var t60 bool = t59 == 43
                            if t60 {
                                jp10 = true
                            } else {
                                var t61 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, index__0)
                                var t62 bool = t61 == 45
                                jp10 = t62
                            }
                        } else {
                            jp10 = false
                        }
                        if jp10 {
                            var t55 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, index__0)
                            var t56 bool = t55 == 45
                            exponent_negative__0 = t56
                            var compound_old3 int = index__0
                            var compound_value3 int = 1
                            var t57 int = compound_old3 + compound_value3
                            index__0 = t57
                        } else {}
                        var exponent_digits__0 bool = false
                        previous_digit__0 = false
                        Loop_loop1__3:
                        for {
                            var t29 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
                            var t30 bool = index__0 < t29
                            if t30 {
                                var current__0 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, index__0)
                                var t31 bool = current__0 >= 48
                                var jp11 bool
                                if t31 {
                                    var t54 bool = current__0 <= 57
                                    jp11 = t54
                                } else {
                                    jp11 = false
                                }
                                if jp11 {
                                    exponent_digits__0 = true
                                    previous_digit__0 = true
                                    var t32 uint8 = current__0 - 48
                                    var digit__0 int = int(uint8(t32))
                                    var t33 int = jp6 - digit__0
                                    var t34 int = t33 / 10
                                    var t35 bool = exponent__0 > t34
                                    var jp12 int
                                    if t35 {
                                        jp12 = jp6
                                    } else {
                                        var t38 int = exponent__0 * 10
                                        var t39 int = t38 + digit__0
                                        jp12 = t39
                                    }
                                    exponent__0 = jp12
                                    var compound_old1 int = index__0
                                    var compound_value1 int = 1
                                    var t36 int = compound_old1 + compound_value1
                                    index__0 = t36
                                    continue
                                } else {
                                    var t40 bool = current__0 == 95
                                    if t40 {
                                        var t41 bool = !previous_digit__0
                                        var jp13 bool
                                        if t41 {
                                            jp13 = true
                                        } else {
                                            var t51 int = index__0 + 1
                                            var t52 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
                                            var t53 bool = t51 >= t52
                                            jp13 = t53
                                        }
                                        var jp14 bool
                                        if jp13 {
                                            jp14 = true
                                        } else {
                                            var t48 int = index__0 + 1
                                            var t49 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, t48)
                                            var t50 bool = t49 < 48
                                            jp14 = t50
                                        }
                                        var jp15 bool
                                        if jp14 {
                                            jp15 = true
                                        } else {
                                            var t45 int = index__0 + 1
                                            var t46 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, t45)
                                            var t47 bool = t46 > 57
                                            jp15 = t47
                                        }
                                        if jp15 {
                                            var t44 ParsedFloat = invalid_parsed_float()
                                            return t44
                                        } else {
                                            previous_digit__0 = false
                                            var compound_old2 int = index__0
                                            var compound_value2 int = 1
                                            var t42 int = compound_old2 + compound_value2
                                            index__0 = t42
                                            continue
                                        }
                                    } else {
                                        break Loop_loop1__3
                                    }
                                }
                            } else {
                                break Loop_loop1__3
                            }
                        }
                        var t27 bool = !exponent_digits__0
                        if t27 {
                            var t28 ParsedFloat = invalid_parsed_float()
                            return t28
                        } else {
                            var t15 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
                            var t16 bool = index__0 != t15
                            if t16 {
                                var t22 ParsedFloat = invalid_parsed_float()
                                return t22
                            } else {
                                if exponent_negative__0 {
                                    var t21 int = 0 - exponent__0
                                    exponent__0 = t21
                                } else {}
                                var jp8 int
                                if jp3 {
                                    jp8 = 0
                                } else {
                                    var t20 int = exponent__0 - fraction_digits__0
                                    jp8 = t20
                                }
                                var jp9 int
                                if jp3 {
                                    var t18 int = fraction_digits__0 * 4
                                    var t19 int = exponent__0 - t18
                                    jp9 = t19
                                } else {
                                    jp9 = 0
                                }
                                var t17 ParsedFloat = ParsedFloat{
                                    valid: true,
                                    negative: negative__0,
                                    special: 0,
                                    numerator: numerator__0,
                                    decimal_exponent: jp8,
                                    binary_exponent: jp9,
                                    hexadecimal: jp3,
                                    significant_digits: significant_digits__0,
                                }
                                return t17
                            }
                        }
                    } else {
                        if jp3 {
                            var t63 ParsedFloat = invalid_parsed_float()
                            return t63
                        } else {
                            var t15 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
                            var t16 bool = index__0 != t15
                            if t16 {
                                var t22 ParsedFloat = invalid_parsed_float()
                                return t22
                            } else {
                                if exponent_negative__0 {
                                    var t21 int = 0 - exponent__0
                                    exponent__0 = t21
                                } else {}
                                var jp8 int
                                if jp3 {
                                    jp8 = 0
                                } else {
                                    var t20 int = exponent__0 - fraction_digits__0
                                    jp8 = t20
                                }
                                var jp9 int
                                if jp3 {
                                    var t18 int = fraction_digits__0 * 4
                                    var t19 int = exponent__0 - t18
                                    jp9 = t19
                                } else {
                                    jp9 = 0
                                }
                                var t17 ParsedFloat = ParsedFloat{
                                    valid: true,
                                    negative: negative__0,
                                    special: 0,
                                    numerator: numerator__0,
                                    decimal_exponent: jp8,
                                    binary_exponent: jp9,
                                    hexadecimal: jp3,
                                    significant_digits: significant_digits__0,
                                }
                                return t17
                            }
                        }
                    }
                }
            }
        }
    }
}

func float_natural_power5(exponent__0 int) FloatNatural {
    var result__0 FloatNatural
    var inline0 *_goml_vec_uint32 = vec_new__Vec_6uint32()
    vec_push__Vec_6uint32(inline0, 1)
    var inline2 FloatNatural = FloatNatural{
        words: inline0,
    }
    result__0 = inline2
    var count__0 int = 0
    Loop_loop0:
    for {
        var t0 bool = count__0 < exponent__0
        if t0 {
            float_natural_multiply_small(result__0, 5)
            var compound_old0 int = count__0
            var compound_value0 int = 1
            var t1 int = compound_old0 + compound_value0
            count__0 = t1
            continue
        } else {
            break Loop_loop0
        }
    }
    return result__0
}

func float_rational_bits(numerator__0 FloatNatural, denominator__0 FloatNatural, binary_shift__0 int, mantissa_bits__0 int, exponent_bias__0 int) Tuple2_6uint64_4bool {
    var t0 bool
    var inline0 *_goml_vec_uint32 = numerator__0.words
    var inline1 bool = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_is__empty____T__u32(inline0)
    t0 = inline1
    if t0 {
        var t75 Tuple2_6uint64_4bool = Tuple2_6uint64_4bool{
            _0: 0,
            _1: false,
        }
        return t75
    } else {
        var minimum_exponent__0 int = 1 - exponent_bias__0
        var t1 int = float_natural_bit_length(numerator__0)
        var t2 int = float_natural_bit_length(denominator__0)
        var t3 int = t1 - t2
        var estimated_exponent__0 int = t3 + binary_shift__0
        var t4 int = exponent_bias__0 + 1
        var t5 bool = estimated_exponent__0 > t4
        if t5 {
            var t70 int = exponent_bias__0 + exponent_bias__0
            var t71 int = t70 + 1
            var t72 uint64 = uint64(int(t71))
            var t73 uint64 = t72 << mantissa_bits__0
            var t74 Tuple2_6uint64_4bool = Tuple2_6uint64_4bool{
                _0: t73,
                _1: true,
            }
            return t74
        } else {
            var t6 int = minimum_exponent__0 - mantissa_bits__0
            var t7 int = t6 - 1
            var t8 bool = estimated_exponent__0 < t7
            if t8 {
                var t69 Tuple2_6uint64_4bool = Tuple2_6uint64_4bool{
                    _0: 0,
                    _1: false,
                }
                return t69
            } else {
                var t9 bool = binary_shift__0 >= 0
                var jp0 FloatNatural
                if t9 {
                    var t67 FloatNatural = float_natural_shift_left(numerator__0, binary_shift__0)
                    jp0 = t67
                } else {
                    var t68 FloatNatural = float_natural_copy(numerator__0)
                    jp0 = t68
                }
                var t10 bool = binary_shift__0 >= 0
                var jp1 FloatNatural
                if t10 {
                    var t64 FloatNatural = float_natural_copy(denominator__0)
                    jp1 = t64
                } else {
                    var t65 int = 0 - binary_shift__0
                    var t66 FloatNatural = float_natural_shift_left(denominator__0, t65)
                    jp1 = t66
                }
                var t11 int = float_natural_bit_length(jp0)
                var t12 int = float_natural_bit_length(jp1)
                var exponent__0 int = t11 - t12
                var t13 bool = exponent__0 >= 0
                var jp2 int
                if t13 {
                    var t59 FloatNatural = float_natural_shift_left(jp1, exponent__0)
                    var t60 int = float_natural_compare(jp0, t59)
                    jp2 = t60
                } else {
                    var t61 int = 0 - exponent__0
                    var t62 FloatNatural = float_natural_shift_left(jp0, t61)
                    var t63 int = float_natural_compare(t62, jp1)
                    jp2 = t63
                }
                var t14 bool = jp2 < 0
                if t14 {
                    var compound_old2 int = exponent__0
                    var compound_value2 int = 1
                    var t57 int = compound_old2 - compound_value2
                    exponent__0 = t57
                } else {}
                var t15 bool = exponent__0 > exponent_bias__0
                if t15 {
                    var t52 int = exponent_bias__0 + exponent_bias__0
                    var t53 int = t52 + 1
                    var t54 uint64 = uint64(int(t53))
                    var t55 uint64 = t54 << mantissa_bits__0
                    var t56 Tuple2_6uint64_4bool = Tuple2_6uint64_4bool{
                        _0: t55,
                        _1: true,
                    }
                    return t56
                } else {
                    var t16 bool = exponent__0 < minimum_exponent__0
                    var jp3 uint64
                    if t16 {
                        var t48 int = mantissa_bits__0 - minimum_exponent__0
                        var t49 uint64 = float_rational_quotient(jp0, jp1, t48)
                        jp3 = t49
                    } else {
                        var t50 int = mantissa_bits__0 - exponent__0
                        var t51 uint64 = float_rational_quotient(jp0, jp1, t50)
                        jp3 = t51
                    }
                    var mantissa__0 uint64 = jp3
                    var t17 bool = exponent__0 < minimum_exponent__0
                    if t17 {
                        var t18 bool = mantissa__0 == 0
                        if t18 {
                            var t19 Tuple2_6uint64_4bool = Tuple2_6uint64_4bool{
                                _0: 0,
                                _1: false,
                            }
                            return t19
                        } else {
                            var t20 uint64 = 1 << mantissa_bits__0
                            var t21 bool = mantissa__0 >= t20
                            if t21 {
                                var t22 uint64 = 1 << mantissa_bits__0
                                var t23 uint64 = 1 << mantissa_bits__0
                                var t24 uint64 = mantissa__0 - t23
                                var t25 uint64 = t22 | t24
                                var t26 Tuple2_6uint64_4bool = Tuple2_6uint64_4bool{
                                    _0: t25,
                                    _1: false,
                                }
                                return t26
                            } else {
                                var t27 Tuple2_6uint64_4bool = Tuple2_6uint64_4bool{
                                    _0: mantissa__0,
                                    _1: false,
                                }
                                return t27
                            }
                        }
                    } else {
                        var t28 int = mantissa_bits__0 + 1
                        var t29 uint64 = 1 << t28
                        var t30 bool = mantissa__0 >= t29
                        if t30 {
                            var compound_old0 uint64 = mantissa__0
                            var compound_value0 int = 1
                            var t44 uint64 = compound_old0 >> compound_value0
                            mantissa__0 = t44
                            var compound_old1 int = exponent__0
                            var compound_value1 int = 1
                            var t46 int = compound_old1 + compound_value1
                            exponent__0 = t46
                        } else {}
                        var t31 bool = exponent__0 > exponent_bias__0
                        if t31 {
                            var t32 int = exponent_bias__0 + exponent_bias__0
                            var t33 int = t32 + 1
                            var t34 uint64 = uint64(int(t33))
                            var t35 uint64 = t34 << mantissa_bits__0
                            var t36 Tuple2_6uint64_4bool = Tuple2_6uint64_4bool{
                                _0: t35,
                                _1: true,
                            }
                            return t36
                        } else {
                            var t37 int = exponent__0 + exponent_bias__0
                            var t38 uint64 = uint64(int(t37))
                            var t39 uint64 = t38 << mantissa_bits__0
                            var t40 uint64 = 1 << mantissa_bits__0
                            var t41 uint64 = mantissa__0 - t40
                            var t42 uint64 = t39 | t41
                            var t43 Tuple2_6uint64_4bool = Tuple2_6uint64_4bool{
                                _0: t42,
                                _1: false,
                            }
                            return t43
                        }
                    }
                }
            }
        }
    }
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_is__empty____T__u32(self__0 *_goml_vec_uint32) bool {
    var t0 int = vec_len__Vec_6uint32(self__0)
    var t1 bool = t0 == 0
    return t1
}

func float_natural_trim(value__0 FloatNatural) struct{} {
    Loop_loop0:
    for {
        var t0 *_goml_vec_uint32 = value__0.words
        var t1 bool
        var inline3 int = vec_len__Vec_6uint32(t0)
        var inline4 bool = inline3 == 0
        t1 = inline4
        var t2 bool = !t1
        var jp0 bool
        if t2 {
            var t7 *_goml_vec_uint32 = value__0.words
            var t8 *_goml_vec_uint32 = value__0.words
            var t9 int
            var inline2 int = vec_len__Vec_6uint32(t8)
            t9 = inline2
            var t10 int = t9 - 1
            var t11 uint32 = vec_get__Vec_6uint32(t7, t10)
            var t12 bool = t11 == 0
            jp0 = t12
        } else {
            jp0 = false
        }
        if jp0 {
            var t3 *_goml_vec_uint32 = value__0.words
            var t4 *_goml_vec_uint32 = value__0.words
            var t5 int
            var inline1 int = vec_len__Vec_6uint32(t4)
            t5 = inline1
            var t6 int = t5 - 1
            vec_truncate__Vec_6uint32(t3, t6)
            continue
        } else {
            break Loop_loop0
        }
    }
    return struct{}{}
}

func string_byte_slice(value__0 string, start__0 int, end__0 int) string {
    var t0 bool = string_is_char_boundary(value__0, start__0)
    var jp0 bool
    if t0 {
        var t3 bool = string_is_char_boundary(value__0, end__0)
        jp0 = t3
    } else {
        jp0 = false
    }
    if jp0 {
        var t1 string = _goml_runtime_core_string_byte_slice(value__0, start__0, end__0)
        return t1
    } else {
        var t2 string = _goml_runtime_core_string_byte_slice(value__0, -1, -1)
        return t2
    }
}

func string_equals_ascii_case(value__0 string, expected__0 string) bool {
    var t0 int
    var inline9 int = _goml_runtime_core_string_len(value__0)
    t0 = inline9
    var t1 int
    var inline8 int = _goml_runtime_core_string_len(expected__0)
    t1 = inline8
    var t2 bool = t0 != t1
    if t2 {
        return false
    } else {
        var index__0 int = 0
        var inline0_lhs uint8 = 97
        var inline0_rhs uint8 = 65
        var inline0 uint8 = inline0_lhs - inline0_rhs
        Loop_loop0:
        for {
            var t3 int
            var inline7 int = _goml_runtime_core_string_len(value__0)
            t3 = inline7
            var t4 bool = index__0 < t3
            if t4 {
                var t5 uint8
                var inline6 uint8 = _goml_runtime_core_string_byte_get(value__0, index__0)
                t5 = inline6
                var t6 uint8
                var inline2 bool = t5 >= 65
                var inline3 bool
                if inline2 {
                    var inline5 bool = t5 <= 90
                    inline3 = inline5
                } else {
                    inline3 = false
                }
                if inline3 {
                    var inline4 uint8 = t5 + inline0
                    t6 = inline4
                    var t7 uint8
                    var inline1 uint8 = _goml_runtime_core_string_byte_get(expected__0, index__0)
                    t7 = inline1
                    var t8 bool = t6 != t7
                    if t8 {
                        return false
                    } else {
                        var compound_old0 int = index__0
                        var compound_value0 int = 1
                        var t9 int = compound_old0 + compound_value0
                        index__0 = t9
                        continue
                    }
                } else {
                    t6 = t5
                    var t7 uint8
                    var inline1 uint8 = _goml_runtime_core_string_byte_get(expected__0, index__0)
                    t7 = inline1
                    var t8 bool = t6 != t7
                    if t8 {
                        return false
                    } else {
                        var compound_old0 int = index__0
                        var compound_value0 int = 1
                        var t9 int = compound_old0 + compound_value0
                        index__0 = t9
                        continue
                    }
                }
            } else {
                break Loop_loop0
            }
        }
        return true
    }
}

func ascii_lower(value__0 uint8) uint8 {
    var t0 bool = value__0 >= 65
    var jp0 bool
    if t0 {
        var t3 bool = value__0 <= 90
        jp0 = t3
    } else {
        jp0 = false
    }
    if jp0 {
        var t1_lhs uint8 = 97
        var t1_rhs uint8 = 65
        var t1 uint8 = t1_lhs - t1_rhs
        var t2 uint8 = value__0 + t1
        return t2
    } else {
        return value__0
    }
}

func float_digit(value__0 uint8, base__0 int) Tuple2_4bool_3int {
    var t0 bool = value__0 >= 48
    var jp0 bool
    if t0 {
        var t15 bool = value__0 <= 57
        jp0 = t15
    } else {
        jp0 = false
    }
    var jp1 int
    if jp0 {
        var t4 uint8 = value__0 - 48
        var t5 int = int(uint8(t4))
        jp1 = t5
        var t1 bool = jp1 < base__0
        if t1 {
            var t2 Tuple2_4bool_3int = Tuple2_4bool_3int{
                _0: true,
                _1: jp1,
            }
            return t2
        } else {
            var t3 Tuple2_4bool_3int = Tuple2_4bool_3int{
                _0: false,
                _1: 0,
            }
            return t3
        }
    } else {
        var t6 uint8
        var inline10 bool = value__0 >= 65
        var inline11 bool
        if inline10 {
            var inline14 bool = value__0 <= 90
            inline11 = inline14
        } else {
            inline11 = false
        }
        if inline11 {
            var inline12_lhs uint8 = 97
            var inline12_rhs uint8 = 65
            var inline12 uint8 = inline12_lhs - inline12_rhs
            var inline13 uint8 = value__0 + inline12
            t6 = inline13
            var t7 bool = t6 >= 97
            var jp2 bool
            if t7 {
                var t13 uint8
                var inline5 bool = value__0 >= 65
                var inline6 bool
                if inline5 {
                    var inline9 bool = value__0 <= 90
                    inline6 = inline9
                } else {
                    inline6 = false
                }
                if inline6 {
                    var inline7_lhs uint8 = 97
                    var inline7_rhs uint8 = 65
                    var inline7 uint8 = inline7_lhs - inline7_rhs
                    var inline8 uint8 = value__0 + inline7
                    t13 = inline8
                    var t14 bool = t13 <= 102
                    jp2 = t14
                    if jp2 {
                        var t8 uint8
                        var inline0 bool = value__0 >= 65
                        var inline1 bool
                        if inline0 {
                            var inline4 bool = value__0 <= 90
                            inline1 = inline4
                        } else {
                            inline1 = false
                        }
                        if inline1 {
                            var inline2_lhs uint8 = 97
                            var inline2_rhs uint8 = 65
                            var inline2 uint8 = inline2_lhs - inline2_rhs
                            var inline3 uint8 = value__0 + inline2
                            t8 = inline3
                            var t9 uint8 = t8 - 97
                            var t10 uint8 = t9 + 10
                            var t11 int = int(uint8(t10))
                            jp1 = t11
                            var t1 bool = jp1 < base__0
                            if t1 {
                                var t2 Tuple2_4bool_3int = Tuple2_4bool_3int{
                                    _0: true,
                                    _1: jp1,
                                }
                                return t2
                            } else {
                                var t3 Tuple2_4bool_3int = Tuple2_4bool_3int{
                                    _0: false,
                                    _1: 0,
                                }
                                return t3
                            }
                        } else {
                            t8 = value__0
                            var t9 uint8 = t8 - 97
                            var t10 uint8 = t9 + 10
                            var t11 int = int(uint8(t10))
                            jp1 = t11
                            var t1 bool = jp1 < base__0
                            if t1 {
                                var t2 Tuple2_4bool_3int = Tuple2_4bool_3int{
                                    _0: true,
                                    _1: jp1,
                                }
                                return t2
                            } else {
                                var t3 Tuple2_4bool_3int = Tuple2_4bool_3int{
                                    _0: false,
                                    _1: 0,
                                }
                                return t3
                            }
                        }
                    } else {
                        var t12 Tuple2_4bool_3int = Tuple2_4bool_3int{
                            _0: false,
                            _1: 0,
                        }
                        return t12
                    }
                } else {
                    t13 = value__0
                    var t14 bool = t13 <= 102
                    jp2 = t14
                    if jp2 {
                        var t8 uint8
                        var inline0 bool = value__0 >= 65
                        var inline1 bool
                        if inline0 {
                            var inline4 bool = value__0 <= 90
                            inline1 = inline4
                        } else {
                            inline1 = false
                        }
                        if inline1 {
                            var inline2_lhs uint8 = 97
                            var inline2_rhs uint8 = 65
                            var inline2 uint8 = inline2_lhs - inline2_rhs
                            var inline3 uint8 = value__0 + inline2
                            t8 = inline3
                            var t9 uint8 = t8 - 97
                            var t10 uint8 = t9 + 10
                            var t11 int = int(uint8(t10))
                            jp1 = t11
                            var t1 bool = jp1 < base__0
                            if t1 {
                                var t2 Tuple2_4bool_3int = Tuple2_4bool_3int{
                                    _0: true,
                                    _1: jp1,
                                }
                                return t2
                            } else {
                                var t3 Tuple2_4bool_3int = Tuple2_4bool_3int{
                                    _0: false,
                                    _1: 0,
                                }
                                return t3
                            }
                        } else {
                            t8 = value__0
                            var t9 uint8 = t8 - 97
                            var t10 uint8 = t9 + 10
                            var t11 int = int(uint8(t10))
                            jp1 = t11
                            var t1 bool = jp1 < base__0
                            if t1 {
                                var t2 Tuple2_4bool_3int = Tuple2_4bool_3int{
                                    _0: true,
                                    _1: jp1,
                                }
                                return t2
                            } else {
                                var t3 Tuple2_4bool_3int = Tuple2_4bool_3int{
                                    _0: false,
                                    _1: 0,
                                }
                                return t3
                            }
                        }
                    } else {
                        var t12 Tuple2_4bool_3int = Tuple2_4bool_3int{
                            _0: false,
                            _1: 0,
                        }
                        return t12
                    }
                }
            } else {
                jp2 = false
                if jp2 {
                    var t8 uint8
                    var inline0 bool = value__0 >= 65
                    var inline1 bool
                    if inline0 {
                        var inline4 bool = value__0 <= 90
                        inline1 = inline4
                    } else {
                        inline1 = false
                    }
                    if inline1 {
                        var inline2_lhs uint8 = 97
                        var inline2_rhs uint8 = 65
                        var inline2 uint8 = inline2_lhs - inline2_rhs
                        var inline3 uint8 = value__0 + inline2
                        t8 = inline3
                        var t9 uint8 = t8 - 97
                        var t10 uint8 = t9 + 10
                        var t11 int = int(uint8(t10))
                        jp1 = t11
                        var t1 bool = jp1 < base__0
                        if t1 {
                            var t2 Tuple2_4bool_3int = Tuple2_4bool_3int{
                                _0: true,
                                _1: jp1,
                            }
                            return t2
                        } else {
                            var t3 Tuple2_4bool_3int = Tuple2_4bool_3int{
                                _0: false,
                                _1: 0,
                            }
                            return t3
                        }
                    } else {
                        t8 = value__0
                        var t9 uint8 = t8 - 97
                        var t10 uint8 = t9 + 10
                        var t11 int = int(uint8(t10))
                        jp1 = t11
                        var t1 bool = jp1 < base__0
                        if t1 {
                            var t2 Tuple2_4bool_3int = Tuple2_4bool_3int{
                                _0: true,
                                _1: jp1,
                            }
                            return t2
                        } else {
                            var t3 Tuple2_4bool_3int = Tuple2_4bool_3int{
                                _0: false,
                                _1: 0,
                            }
                            return t3
                        }
                    }
                } else {
                    var t12 Tuple2_4bool_3int = Tuple2_4bool_3int{
                        _0: false,
                        _1: 0,
                    }
                    return t12
                }
            }
        } else {
            t6 = value__0
            var t7 bool = t6 >= 97
            var jp2 bool
            if t7 {
                var t13 uint8
                var inline5 bool = value__0 >= 65
                var inline6 bool
                if inline5 {
                    var inline9 bool = value__0 <= 90
                    inline6 = inline9
                } else {
                    inline6 = false
                }
                if inline6 {
                    var inline7_lhs uint8 = 97
                    var inline7_rhs uint8 = 65
                    var inline7 uint8 = inline7_lhs - inline7_rhs
                    var inline8 uint8 = value__0 + inline7
                    t13 = inline8
                    var t14 bool = t13 <= 102
                    jp2 = t14
                    if jp2 {
                        var t8 uint8
                        var inline0 bool = value__0 >= 65
                        var inline1 bool
                        if inline0 {
                            var inline4 bool = value__0 <= 90
                            inline1 = inline4
                        } else {
                            inline1 = false
                        }
                        if inline1 {
                            var inline2_lhs uint8 = 97
                            var inline2_rhs uint8 = 65
                            var inline2 uint8 = inline2_lhs - inline2_rhs
                            var inline3 uint8 = value__0 + inline2
                            t8 = inline3
                            var t9 uint8 = t8 - 97
                            var t10 uint8 = t9 + 10
                            var t11 int = int(uint8(t10))
                            jp1 = t11
                            var t1 bool = jp1 < base__0
                            if t1 {
                                var t2 Tuple2_4bool_3int = Tuple2_4bool_3int{
                                    _0: true,
                                    _1: jp1,
                                }
                                return t2
                            } else {
                                var t3 Tuple2_4bool_3int = Tuple2_4bool_3int{
                                    _0: false,
                                    _1: 0,
                                }
                                return t3
                            }
                        } else {
                            t8 = value__0
                            var t9 uint8 = t8 - 97
                            var t10 uint8 = t9 + 10
                            var t11 int = int(uint8(t10))
                            jp1 = t11
                            var t1 bool = jp1 < base__0
                            if t1 {
                                var t2 Tuple2_4bool_3int = Tuple2_4bool_3int{
                                    _0: true,
                                    _1: jp1,
                                }
                                return t2
                            } else {
                                var t3 Tuple2_4bool_3int = Tuple2_4bool_3int{
                                    _0: false,
                                    _1: 0,
                                }
                                return t3
                            }
                        }
                    } else {
                        var t12 Tuple2_4bool_3int = Tuple2_4bool_3int{
                            _0: false,
                            _1: 0,
                        }
                        return t12
                    }
                } else {
                    t13 = value__0
                    var t14 bool = t13 <= 102
                    jp2 = t14
                    if jp2 {
                        var t8 uint8
                        var inline0 bool = value__0 >= 65
                        var inline1 bool
                        if inline0 {
                            var inline4 bool = value__0 <= 90
                            inline1 = inline4
                        } else {
                            inline1 = false
                        }
                        if inline1 {
                            var inline2_lhs uint8 = 97
                            var inline2_rhs uint8 = 65
                            var inline2 uint8 = inline2_lhs - inline2_rhs
                            var inline3 uint8 = value__0 + inline2
                            t8 = inline3
                            var t9 uint8 = t8 - 97
                            var t10 uint8 = t9 + 10
                            var t11 int = int(uint8(t10))
                            jp1 = t11
                            var t1 bool = jp1 < base__0
                            if t1 {
                                var t2 Tuple2_4bool_3int = Tuple2_4bool_3int{
                                    _0: true,
                                    _1: jp1,
                                }
                                return t2
                            } else {
                                var t3 Tuple2_4bool_3int = Tuple2_4bool_3int{
                                    _0: false,
                                    _1: 0,
                                }
                                return t3
                            }
                        } else {
                            t8 = value__0
                            var t9 uint8 = t8 - 97
                            var t10 uint8 = t9 + 10
                            var t11 int = int(uint8(t10))
                            jp1 = t11
                            var t1 bool = jp1 < base__0
                            if t1 {
                                var t2 Tuple2_4bool_3int = Tuple2_4bool_3int{
                                    _0: true,
                                    _1: jp1,
                                }
                                return t2
                            } else {
                                var t3 Tuple2_4bool_3int = Tuple2_4bool_3int{
                                    _0: false,
                                    _1: 0,
                                }
                                return t3
                            }
                        }
                    } else {
                        var t12 Tuple2_4bool_3int = Tuple2_4bool_3int{
                            _0: false,
                            _1: 0,
                        }
                        return t12
                    }
                }
            } else {
                jp2 = false
                if jp2 {
                    var t8 uint8
                    var inline0 bool = value__0 >= 65
                    var inline1 bool
                    if inline0 {
                        var inline4 bool = value__0 <= 90
                        inline1 = inline4
                    } else {
                        inline1 = false
                    }
                    if inline1 {
                        var inline2_lhs uint8 = 97
                        var inline2_rhs uint8 = 65
                        var inline2 uint8 = inline2_lhs - inline2_rhs
                        var inline3 uint8 = value__0 + inline2
                        t8 = inline3
                        var t9 uint8 = t8 - 97
                        var t10 uint8 = t9 + 10
                        var t11 int = int(uint8(t10))
                        jp1 = t11
                        var t1 bool = jp1 < base__0
                        if t1 {
                            var t2 Tuple2_4bool_3int = Tuple2_4bool_3int{
                                _0: true,
                                _1: jp1,
                            }
                            return t2
                        } else {
                            var t3 Tuple2_4bool_3int = Tuple2_4bool_3int{
                                _0: false,
                                _1: 0,
                            }
                            return t3
                        }
                    } else {
                        t8 = value__0
                        var t9 uint8 = t8 - 97
                        var t10 uint8 = t9 + 10
                        var t11 int = int(uint8(t10))
                        jp1 = t11
                        var t1 bool = jp1 < base__0
                        if t1 {
                            var t2 Tuple2_4bool_3int = Tuple2_4bool_3int{
                                _0: true,
                                _1: jp1,
                            }
                            return t2
                        } else {
                            var t3 Tuple2_4bool_3int = Tuple2_4bool_3int{
                                _0: false,
                                _1: 0,
                            }
                            return t3
                        }
                    }
                } else {
                    var t12 Tuple2_4bool_3int = Tuple2_4bool_3int{
                        _0: false,
                        _1: 0,
                    }
                    return t12
                }
            }
        }
    }
}

func float_natural_add_small(value__0 FloatNatural, addition__0 uint32) struct{} {
    var carry__0 uint64 = uint64(uint32(addition__0))
    var index__0 int = 0
    Loop_loop0:
    for {
        var t0 bool = carry__0 != 0
        if t0 {
            var t1 *_goml_vec_uint32 = value__0.words
            var t2 int
            var inline2 int = vec_len__Vec_6uint32(t1)
            t2 = inline2
            var t3 bool = index__0 == t2
            if t3 {
                var t11 *_goml_vec_uint32 = value__0.words
                var inline0 uint32 = 0
                vec_push__Vec_6uint32(t11, inline0)
            } else {}
            var t4 *_goml_vec_uint32 = value__0.words
            var t5 uint32 = vec_get__Vec_6uint32(t4, index__0)
            var t6 uint64 = uint64(uint32(t5))
            var sum__0 uint64 = t6 + carry__0
            var place0 *_goml_vec_uint32 = value__0.words
            var index0 int = index__0
            vec_get__Vec_6uint32(place0, index0)
            var value0 uint32 = uint32(uint64(sum__0))
            vec_set__Vec_6uint32(place0, index0, value0)
            var t8 uint64 = sum__0 >> 32
            carry__0 = t8
            var compound_old0 int = index__0
            var compound_value0 int = 1
            var t9 int = compound_old0 + compound_value0
            index__0 = t9
            continue
        } else {
            break Loop_loop0
        }
    }
    return struct{}{}
}

func invalid_parsed_float() ParsedFloat {
    var t0 FloatNatural
    var inline0 *_goml_vec_uint32 = vec_new__Vec_6uint32()
    var inline1 FloatNatural = FloatNatural{
        words: inline0,
    }
    t0 = inline1
    var t1 ParsedFloat = ParsedFloat{
        valid: false,
        negative: false,
        special: 0,
        numerator: t0,
        decimal_exponent: 0,
        binary_exponent: 0,
        hexadecimal: false,
        significant_digits: 0,
    }
    return t1
}

func float_natural_bit_length(value__0 FloatNatural) int {
    var t0 *_goml_vec_uint32 = value__0.words
    var t1 bool
    var inline2 int = vec_len__Vec_6uint32(t0)
    var inline3 bool = inline2 == 0
    t1 = inline3
    if t1 {
        return 0
    } else {
        var t2 *_goml_vec_uint32 = value__0.words
        var t3 *_goml_vec_uint32 = value__0.words
        var t4 int
        var inline1 int = vec_len__Vec_6uint32(t3)
        t4 = inline1
        var t5 int = t4 - 1
        var high__0 uint32 = vec_get__Vec_6uint32(t2, t5)
        var bits__0 int = 0
        Loop_loop0:
        for {
            var t11 bool = high__0 != 0
            if t11 {
                var compound_old0 uint32 = high__0
                var compound_value0 int = 1
                var t12 uint32 = compound_old0 >> compound_value0
                high__0 = t12
                var compound_old1 int = bits__0
                var compound_value1 int = 1
                var t14 int = compound_old1 + compound_value1
                bits__0 = t14
                continue
            } else {
                break Loop_loop0
            }
        }
        var t6 *_goml_vec_uint32 = value__0.words
        var t7 int
        var inline0 int = vec_len__Vec_6uint32(t6)
        t7 = inline0
        var t8 int = t7 - 1
        var t9 int = t8 * 32
        var t10 int = t9 + bits__0
        return t10
    }
}

func float_natural_compare(left__0 FloatNatural, right__0 FloatNatural) int {
    var t0 *_goml_vec_uint32 = left__0.words
    var t1 int
    var inline4 int = vec_len__Vec_6uint32(t0)
    t1 = inline4
    var t2 *_goml_vec_uint32 = right__0.words
    var t3 int
    var inline3 int = vec_len__Vec_6uint32(t2)
    t3 = inline3
    var t4 bool = t1 < t3
    if t4 {
        return -1
    } else {
        var t19 *_goml_vec_uint32 = left__0.words
        var t20 int
        var inline2 int = vec_len__Vec_6uint32(t19)
        t20 = inline2
        var t21 *_goml_vec_uint32 = right__0.words
        var t22 int
        var inline1 int = vec_len__Vec_6uint32(t21)
        t22 = inline1
        var t23 bool = t20 > t22
        if t23 {
            return 1
        } else {
            var t5 *_goml_vec_uint32 = left__0.words
            var index__0 int
            var inline0 int = vec_len__Vec_6uint32(t5)
            index__0 = inline0
            Loop_loop0:
            for {
                var t6 bool = index__0 > 0
                if t6 {
                    var compound_old0 int = index__0
                    var compound_value0 int = 1
                    var t7 int = compound_old0 - compound_value0
                    index__0 = t7
                    var t9 *_goml_vec_uint32 = left__0.words
                    var t10 uint32 = vec_get__Vec_6uint32(t9, index__0)
                    var t11 *_goml_vec_uint32 = right__0.words
                    var t12 uint32 = vec_get__Vec_6uint32(t11, index__0)
                    var t13 bool = t10 < t12
                    if t13 {
                        return -1
                    } else {
                        var t14 *_goml_vec_uint32 = left__0.words
                        var t15 uint32 = vec_get__Vec_6uint32(t14, index__0)
                        var t16 *_goml_vec_uint32 = right__0.words
                        var t17 uint32 = vec_get__Vec_6uint32(t16, index__0)
                        var t18 bool = t15 > t17
                        if t18 {
                            return 1
                        } else {
                            continue
                        }
                    }
                } else {
                    break Loop_loop0
                }
            }
            return 0
        }
    }
}

func float_rational_quotient(numerator__0 FloatNatural, denominator__0 FloatNatural, shift__0 int) uint64 {
    var t0 bool = shift__0 >= 0
    var jp0 FloatNatural
    if t0 {
        var t22 FloatNatural = float_natural_shift_left(numerator__0, shift__0)
        jp0 = t22
    } else {
        var t23 FloatNatural = float_natural_copy(numerator__0)
        jp0 = t23
    }
    var t1 bool = shift__0 >= 0
    var jp1 FloatNatural
    if t1 {
        var t19 FloatNatural = float_natural_copy(denominator__0)
        jp1 = t19
    } else {
        var t20 int = 0 - shift__0
        var t21 FloatNatural = float_natural_shift_left(denominator__0, t20)
        jp1 = t21
    }
    var quotient__0 uint64 = 0
    Loop_loop0:
    for {
        var t8 int = float_natural_compare(jp0, jp1)
        var t9 bool = t8 >= 0
        if t9 {
            var t10 int = float_natural_bit_length(jp0)
            var t11 int = float_natural_bit_length(jp1)
            var offset__0 int = t10 - t11
            var part__0 FloatNatural = float_natural_shift_left(jp1, offset__0)
            var t12 int = float_natural_compare(jp0, part__0)
            var t13 bool = t12 < 0
            if t13 {
                var compound_old2 int = offset__0
                var compound_value2 int = 1
                var t16 int = compound_old2 - compound_value2
                offset__0 = t16
                var t18 FloatNatural = float_natural_shift_left(jp1, offset__0)
                part__0 = t18
            } else {}
            float_natural_subtract(jp0, part__0)
            var compound_old1 uint64 = quotient__0
            var compound_value1 uint64 = 1 << offset__0
            var t14 uint64 = compound_old1 | compound_value1
            quotient__0 = t14
            continue
        } else {
            break Loop_loop0
        }
    }
    var doubled__0 FloatNatural = float_natural_shift_left(jp0, 1)
    var rounding__0 int = float_natural_compare(doubled__0, jp1)
    var t2 bool = rounding__0 > 0
    var jp2 bool
    if t2 {
        jp2 = true
    } else {
        var t5 bool = rounding__0 == 0
        if t5 {
            var t6 uint64 = quotient__0 & 1
            var t7 bool = t6 == 1
            jp2 = t7
        } else {
            jp2 = false
        }
    }
    if jp2 {
        var compound_old0 uint64 = quotient__0
        var compound_value0 uint64 = 1
        var t3 uint64 = compound_old0 + compound_value0
        quotient__0 = t3
    } else {}
    return quotient__0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_truncate____T__u32(self__0 *_goml_vec_uint32, len__0 int) struct{} {
    vec_truncate__Vec_6uint32(self__0, len__0)
    return struct{}{}
}

func string_is_char_boundary(value__0 string, index__0 int) bool {
    var t0 bool = index__0 < 0
    var jp0 bool
    if t0 {
        jp0 = true
    } else {
        var t6 int
        var inline2 int = _goml_runtime_core_string_len(value__0)
        t6 = inline2
        var t7 bool = index__0 > t6
        jp0 = t7
    }
    if jp0 {
        return false
    } else {
        var t1 int
        var inline1 int = _goml_runtime_core_string_len(value__0)
        t1 = inline1
        var t2 bool = index__0 == t1
        if t2 {
            return true
        } else {
            var t3 uint8
            var inline0 uint8 = _goml_runtime_core_string_byte_get(value__0, index__0)
            t3 = inline0
            var t4 uint8 = t3 & 192
            var t5 bool = t4 != 128
            return t5
        }
    }
}

func float_natural_subtract(value__0 FloatNatural, other__0 FloatNatural) struct{} {
    var base__0 uint64 = 4294967296
    var borrow__0 uint64 = 0
    var index__0 int = 0
    Loop_loop0:
    for {
        var t1 *_goml_vec_uint32 = value__0.words
        var t2 int
        var inline1 int = vec_len__Vec_6uint32(t1)
        t2 = inline1
        var t3 bool = index__0 < t2
        if t3 {
            var t4 *_goml_vec_uint32 = other__0.words
            var t5 int
            var inline0 int = vec_len__Vec_6uint32(t4)
            t5 = inline0
            var t6 bool = index__0 < t5
            var jp0 uint64
            if t6 {
                var t17 *_goml_vec_uint32 = other__0.words
                var t18 uint32 = vec_get__Vec_6uint32(t17, index__0)
                var t19 uint64 = uint64(uint32(t18))
                jp0 = t19
            } else {
                jp0 = 0
            }
            var right__0 uint64 = jp0 + borrow__0
            var t7 *_goml_vec_uint32 = value__0.words
            var t8 uint32 = vec_get__Vec_6uint32(t7, index__0)
            var left__0 uint64 = uint64(uint32(t8))
            var t9 bool = left__0 >= right__0
            if t9 {
                var place0 *_goml_vec_uint32 = value__0.words
                var index0 int = index__0
                vec_get__Vec_6uint32(place0, index0)
                var t12 uint64 = left__0 - right__0
                var value0 uint32 = uint32(uint64(t12))
                vec_set__Vec_6uint32(place0, index0, value0)
                borrow__0 = 0
            } else {
                var place2 *_goml_vec_uint32 = value__0.words
                var index1 int = index__0
                vec_get__Vec_6uint32(place2, index1)
                var t14 uint64 = base__0 + left__0
                var t15 uint64 = t14 - right__0
                var value1 uint32 = uint32(uint64(t15))
                vec_set__Vec_6uint32(place2, index1, value1)
                borrow__0 = 1
            }
            var compound_old0 int = index__0
            var compound_value0 int = 1
            var t10 int = compound_old0 + compound_value0
            index__0 = t10
            continue
        } else {
            break Loop_loop0
        }
    }
    float_natural_trim(value__0)
    return struct{}{}
}

func main() {
    main0()
}
