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

func _goml_runtime_core_string_len(s string) int {
    return int(len(s))
}

func _goml_runtime_core_string_byte_get(s string, i int) uint8 {
    return s[i]
}

func _goml_runtime_core_string_from_utf8(bytes *_goml_vec_uint8) Tuple2_4bool_6string {
    return Tuple2_4bool_6string{
        _0: true,
        _1: string(bytes.items),
    }
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

func vec_new__Vec_5uint8() *_goml_vec_uint8 {
    return &_goml_vec_uint8{
        items: nil,
    }
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

type Tuple2_4char_3int struct {
    _0 rune
    _1 int
}

type Tuple3_4bool_4char_3int struct {
    _0 bool
    _1 rune
    _2 int
}

type Tuple2_4bool_4char struct {
    _0 bool
    _1 rune
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

type _goml_m_Option_____o_char_c_isize_q_ struct {
    _p0 Tuple2_4char_3int
    _tag uint8
}

type Option__char uint64

func check_utf8(bytes__0 *_goml_vec_uint8, expected__0 bool) struct{} {
    var expected_length__0 int
    var inline7 int = vec_len__Vec_5uint8(bytes__0)
    expected_length__0 = inline7
    var mtmp0 Tuple2_4bool_6string = string_from_utf8(bytes__0)
    var x0 bool = mtmp0._0
    var x1 string = mtmp0._1
    var t0 bool = x0 == expected__0
    var inline5 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t0)
    _goml_runtime_core_string_println(inline5)
    if x0 {
        var t1 int
        var inline2 int = _goml_runtime_core_string_len(x1)
        t1 = inline2
        var t2 bool = t1 == expected_length__0
        var inline0 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t2)
        _goml_runtime_core_string_println(inline0)
        return struct{}{}
    } else {
        var t3 bool = x1 == ""
        var inline3 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t3)
        _goml_runtime_core_string_println(inline3)
        return struct{}{}
    }
}

func check_scalar(bytes__0 *_goml_vec_uint8, expected__0 uint32, expected_width__0 int) struct{} {
    var mtmp0 Tuple2_4bool_6string = string_from_utf8(bytes__0)
    var x0 bool = mtmp0._0
    var x1 string = mtmp0._1
    var inline16 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(x0)
    _goml_runtime_core_string_println(inline16)
    var commute_field0 Tuple2_4char_3int
    var inline10 int = 0
    var inline11 Tuple3_4bool_4char_3int = string_decode_utf8_at(x1, inline10)
    var inline12 bool = inline11._0
    var inline13 rune = inline11._1
    var inline14 int = inline11._2
    if inline12 {
        var inline15 Tuple2_4char_3int = Tuple2_4char_3int{
            _0: inline13,
            _1: inline14,
        }
        commute_field0 = inline15
        var x2 rune = commute_field0._0
        var x3 int = commute_field0._1
        var t0 uint32 = uint32(rune(x2))
        var t1 bool = t0 == expected__0
        var inline8 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t1)
        _goml_runtime_core_string_println(inline8)
        var t2 bool = x3 == expected_width__0
        var inline6 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t2)
        _goml_runtime_core_string_println(inline6)
        return struct{}{}
    } else {
        var inline3 bool = false
        var inline4 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(inline3)
        _goml_runtime_core_string_println(inline4)
        var inline0 bool = false
        var inline1 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(inline0)
        _goml_runtime_core_string_println(inline1)
        return struct{}{}
    }
}

func main0() struct{} {
    var t0 *_goml_vec_uint8 = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_new____T__u8()
    check_utf8(t0, true)
    var t1 [1]uint8 = [1]uint8{0}
    var t2 *_goml_vec_uint8 = func(values [1]uint8) *_goml_vec_uint8 {
        var storage struct {
            vector _goml_vec_uint8
            values [1]uint8
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t1)
    check_utf8(t2, true)
    var t3 [1]uint8 = [1]uint8{127}
    var t4 *_goml_vec_uint8 = func(values [1]uint8) *_goml_vec_uint8 {
        var storage struct {
            vector _goml_vec_uint8
            values [1]uint8
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t3)
    check_utf8(t4, true)
    var t5 [2]uint8 = [2]uint8{194, 128}
    var t6 *_goml_vec_uint8 = func(values [2]uint8) *_goml_vec_uint8 {
        var storage struct {
            vector _goml_vec_uint8
            values [2]uint8
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t5)
    check_scalar(t6, 128, 2)
    var t7 [2]uint8 = [2]uint8{223, 191}
    var t8 *_goml_vec_uint8 = func(values [2]uint8) *_goml_vec_uint8 {
        var storage struct {
            vector _goml_vec_uint8
            values [2]uint8
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t7)
    check_scalar(t8, 2047, 2)
    var t9 [3]uint8 = [3]uint8{224, 160, 128}
    var t10 *_goml_vec_uint8 = func(values [3]uint8) *_goml_vec_uint8 {
        var storage struct {
            vector _goml_vec_uint8
            values [3]uint8
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t9)
    check_scalar(t10, 2048, 3)
    var t11 [3]uint8 = [3]uint8{237, 159, 191}
    var t12 *_goml_vec_uint8 = func(values [3]uint8) *_goml_vec_uint8 {
        var storage struct {
            vector _goml_vec_uint8
            values [3]uint8
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t11)
    check_scalar(t12, 55295, 3)
    var t13 [3]uint8 = [3]uint8{238, 128, 128}
    var t14 *_goml_vec_uint8 = func(values [3]uint8) *_goml_vec_uint8 {
        var storage struct {
            vector _goml_vec_uint8
            values [3]uint8
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t13)
    check_scalar(t14, 57344, 3)
    var t15 [3]uint8 = [3]uint8{239, 191, 189}
    var t16 *_goml_vec_uint8 = func(values [3]uint8) *_goml_vec_uint8 {
        var storage struct {
            vector _goml_vec_uint8
            values [3]uint8
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t15)
    check_scalar(t16, 65533, 3)
    var t17 [3]uint8 = [3]uint8{239, 191, 191}
    var t18 *_goml_vec_uint8 = func(values [3]uint8) *_goml_vec_uint8 {
        var storage struct {
            vector _goml_vec_uint8
            values [3]uint8
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t17)
    check_scalar(t18, 65535, 3)
    var t19 [4]uint8 = [4]uint8{240, 144, 128, 128}
    var t20 *_goml_vec_uint8 = func(values [4]uint8) *_goml_vec_uint8 {
        var storage struct {
            vector _goml_vec_uint8
            values [4]uint8
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t19)
    check_scalar(t20, 65536, 4)
    var t21 [4]uint8 = [4]uint8{244, 143, 191, 191}
    var t22 *_goml_vec_uint8 = func(values [4]uint8) *_goml_vec_uint8 {
        var storage struct {
            vector _goml_vec_uint8
            values [4]uint8
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t21)
    check_scalar(t22, 1114111, 4)
    var t23 [1]uint8 = [1]uint8{128}
    var t24 *_goml_vec_uint8 = func(values [1]uint8) *_goml_vec_uint8 {
        var storage struct {
            vector _goml_vec_uint8
            values [1]uint8
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t23)
    check_utf8(t24, false)
    var t25 [1]uint8 = [1]uint8{191}
    var t26 *_goml_vec_uint8 = func(values [1]uint8) *_goml_vec_uint8 {
        var storage struct {
            vector _goml_vec_uint8
            values [1]uint8
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t25)
    check_utf8(t26, false)
    var t27 [2]uint8 = [2]uint8{192, 128}
    var t28 *_goml_vec_uint8 = func(values [2]uint8) *_goml_vec_uint8 {
        var storage struct {
            vector _goml_vec_uint8
            values [2]uint8
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t27)
    check_utf8(t28, false)
    var t29 [2]uint8 = [2]uint8{193, 191}
    var t30 *_goml_vec_uint8 = func(values [2]uint8) *_goml_vec_uint8 {
        var storage struct {
            vector _goml_vec_uint8
            values [2]uint8
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t29)
    check_utf8(t30, false)
    var t31 [1]uint8 = [1]uint8{194}
    var t32 *_goml_vec_uint8 = func(values [1]uint8) *_goml_vec_uint8 {
        var storage struct {
            vector _goml_vec_uint8
            values [1]uint8
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t31)
    check_utf8(t32, false)
    var t33 [2]uint8 = [2]uint8{194, 127}
    var t34 *_goml_vec_uint8 = func(values [2]uint8) *_goml_vec_uint8 {
        var storage struct {
            vector _goml_vec_uint8
            values [2]uint8
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t33)
    check_utf8(t34, false)
    var t35 [3]uint8 = [3]uint8{224, 159, 191}
    var t36 *_goml_vec_uint8 = func(values [3]uint8) *_goml_vec_uint8 {
        var storage struct {
            vector _goml_vec_uint8
            values [3]uint8
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t35)
    check_utf8(t36, false)
    var t37 [2]uint8 = [2]uint8{225, 128}
    var t38 *_goml_vec_uint8 = func(values [2]uint8) *_goml_vec_uint8 {
        var storage struct {
            vector _goml_vec_uint8
            values [2]uint8
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t37)
    check_utf8(t38, false)
    var t39 [3]uint8 = [3]uint8{225, 128, 127}
    var t40 *_goml_vec_uint8 = func(values [3]uint8) *_goml_vec_uint8 {
        var storage struct {
            vector _goml_vec_uint8
            values [3]uint8
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t39)
    check_utf8(t40, false)
    var t41 [3]uint8 = [3]uint8{237, 160, 128}
    var t42 *_goml_vec_uint8 = func(values [3]uint8) *_goml_vec_uint8 {
        var storage struct {
            vector _goml_vec_uint8
            values [3]uint8
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t41)
    check_utf8(t42, false)
    var t43 [4]uint8 = [4]uint8{240, 143, 191, 191}
    var t44 *_goml_vec_uint8 = func(values [4]uint8) *_goml_vec_uint8 {
        var storage struct {
            vector _goml_vec_uint8
            values [4]uint8
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t43)
    check_utf8(t44, false)
    var t45 [3]uint8 = [3]uint8{240, 144, 128}
    var t46 *_goml_vec_uint8 = func(values [3]uint8) *_goml_vec_uint8 {
        var storage struct {
            vector _goml_vec_uint8
            values [3]uint8
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t45)
    var inline24 bool = false
    var inline25 int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_len____T__u8(t46)
    var inline26 Tuple2_4bool_6string = string_from_utf8(t46)
    var inline27 bool = inline26._0
    var inline28 string = inline26._1
    var inline29 bool = inline27 == inline24
    println__T_bool(inline29)
    if inline27 {
        var inline31 int = _goml_m_inherent_i_string_i_string_i_byte__len(inline28)
        var inline32 bool = inline31 == inline25
        println__T_bool(inline32)
    } else {
        var inline34 bool = inline28 == ""
        println__T_bool(inline34)
    }
    var t47 [4]uint8 = [4]uint8{244, 144, 128, 128}
    var t48 *_goml_vec_uint8 = func(values [4]uint8) *_goml_vec_uint8 {
        var storage struct {
            vector _goml_vec_uint8
            values [4]uint8
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t47)
    var inline12 bool = false
    var inline13 int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_len____T__u8(t48)
    var inline14 Tuple2_4bool_6string = string_from_utf8(t48)
    var inline15 bool = inline14._0
    var inline16 string = inline14._1
    var inline17 bool = inline15 == inline12
    println__T_bool(inline17)
    if inline15 {
        var inline19 int = _goml_m_inherent_i_string_i_string_i_byte__len(inline16)
        var inline20 bool = inline19 == inline13
        println__T_bool(inline20)
    } else {
        var inline22 bool = inline16 == ""
        println__T_bool(inline22)
    }
    var t49 [4]uint8 = [4]uint8{245, 128, 128, 128}
    var t50 *_goml_vec_uint8 = func(values [4]uint8) *_goml_vec_uint8 {
        var storage struct {
            vector _goml_vec_uint8
            values [4]uint8
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t49)
    var inline0 bool = false
    var inline1 int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_len____T__u8(t50)
    var inline2 Tuple2_4bool_6string = string_from_utf8(t50)
    var inline3 bool = inline2._0
    var inline4 string = inline2._1
    var inline5 bool = inline3 == inline0
    println__T_bool(inline5)
    if inline3 {
        var inline7 int = _goml_m_inherent_i_string_i_string_i_byte__len(inline4)
        var inline8 bool = inline7 == inline1
        println__T_bool(inline8)
        return struct{}{}
    } else {
        var inline10 bool = inline4 == ""
        println__T_bool(inline10)
        return struct{}{}
    }
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_len____T__u8(self__0 *_goml_vec_uint8) int {
    var t0 int = vec_len__Vec_5uint8(self__0)
    return t0
}

func string_from_utf8(bytes__0 *_goml_vec_uint8) Tuple2_4bool_6string {
    var mtmp0 Tuple2_4bool_6string = _goml_runtime_core_string_from_utf8(bytes__0)
    var x0 string = mtmp0._1
    var index__0 int = 0
    Loop_loop0:
    for {
        var t1 int
        var inline0 int = _goml_runtime_core_string_len(x0)
        t1 = inline0
        var t2 bool = index__0 < t1
        if t2 {
            var mtmp1 Tuple3_4bool_4char_3int = string_decode_utf8_at(x0, index__0)
            var x1 bool = mtmp1._0
            var x2 int = mtmp1._2
            if x1 {
                var compound_old0 int = index__0
                var t3 int = compound_old0 + x2
                index__0 = t3
                continue
            } else {
                var t5 Tuple2_4bool_6string = Tuple2_4bool_6string{
                    _0: false,
                    _1: "",
                }
                return t5
            }
        } else {
            break Loop_loop0
        }
    }
    var t0 Tuple2_4bool_6string = Tuple2_4bool_6string{
        _0: true,
        _1: x0,
    }
    return t0
}

func println__T_bool(value__0 bool) struct{} {
    var t0 string
    var inline0 string = _goml_runtime_core_bool_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_inherent_i_string_i_string_i_byte__len(self__0 string) int {
    var t0 int = _goml_runtime_core_string_len(self__0)
    return t0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_new____T__u8() *_goml_vec_uint8 {
    var t0 *_goml_vec_uint8 = vec_new__Vec_5uint8()
    return t0
}

func string_decode_utf8_at(value__0 string, index__0 int) Tuple3_4bool_4char_3int {
    var length__0 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
    var t0 bool = index__0 < 0
    var jp0 bool
    if t0 {
        jp0 = true
    } else {
        var t63 bool = index__0 >= length__0
        jp0 = t63
    }
    if jp0 {
        var inline25 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
            _0: false,
            _1: 0,
            _2: 0,
        }
        return inline25
    } else {
        var t1 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, index__0)
        var first__0 uint32 = uint32(uint8(t1))
        var t2 bool = first__0 < 128
        if t2 {
            var inline0 int = 1
            var inline1 Option__char = __goml_builtin_char_from_uint32(first__0)
            switch inline1 != Option__char(0) {
            case false:
                var inline2 Tuple3_4bool_4char_3int = utf8_invalid_decode()
                return inline2
            case true:
                var inline3 rune = rune(uint64(inline1) - 1)
                var inline4 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                    _0: true,
                    _1: inline3,
                    _2: inline0,
                }
                return inline4
            default:
                panic("non-exhaustive match")
            }
        } else {
            var t3 bool = first__0 < 194
            if t3 {
                var inline5 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                    _0: false,
                    _1: 0,
                    _2: 0,
                }
                return inline5
            } else {
                var t4 bool = first__0 < 224
                if t4 {
                    var t5 int = length__0 - index__0
                    var t6 bool = t5 < 2
                    if t6 {
                        var inline15 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                            _0: false,
                            _1: 0,
                            _2: 0,
                        }
                        return inline15
                    } else {
                        var t7 int = index__0 + 1
                        var t8 uint8
                        var inline14 uint8 = _goml_runtime_core_string_byte_get(value__0, t7)
                        t8 = inline14
                        var second__0 uint32 = uint32(uint8(t8))
                        var t9 bool
                        var inline12 bool = second__0 < 128
                        if inline12 {
                            t9 = true
                        } else {
                            var inline13 bool = second__0 > 191
                            t9 = inline13
                        }
                        if t9 {
                            var inline6 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                                _0: false,
                                _1: 0,
                                _2: 0,
                            }
                            return inline6
                        } else {
                            var t10 uint32 = first__0 & 31
                            var t11 uint32 = t10 << 6
                            var t12 uint32 = second__0 & 63
                            var t13 uint32 = t11 | t12
                            var inline7 int = 2
                            var inline8 Option__char = __goml_builtin_char_from_uint32(t13)
                            switch inline8 != Option__char(0) {
                            case false:
                                var inline9 Tuple3_4bool_4char_3int = utf8_invalid_decode()
                                return inline9
                            case true:
                                var inline10 rune = rune(uint64(inline8) - 1)
                                var inline11 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                                    _0: true,
                                    _1: inline10,
                                    _2: inline7,
                                }
                                return inline11
                            default:
                                panic("non-exhaustive match")
                            }
                        }
                    }
                } else {
                    var t14 bool = first__0 < 240
                    if t14 {
                        var t15 int = length__0 - index__0
                        var t16 bool = t15 < 3
                        if t16 {
                            var inline24 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                                _0: false,
                                _1: 0,
                                _2: 0,
                            }
                            return inline24
                        } else {
                            var t17 int = index__0 + 1
                            var t18 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, t17)
                            var second__1 uint32 = uint32(uint8(t18))
                            var t19 int = index__0 + 2
                            var t20 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, t19)
                            var third__0 uint32 = uint32(uint8(t20))
                            var t21 bool = utf8_invalid_continuation(second__1)
                            var jp1 bool
                            if t21 {
                                jp1 = true
                            } else {
                                var inline22 bool = third__0 < 128
                                if inline22 {
                                    jp1 = true
                                } else {
                                    var inline23 bool = third__0 > 191
                                    jp1 = inline23
                                }
                            }
                            var jp2 bool
                            if jp1 {
                                jp2 = true
                            } else {
                                var t31 bool = first__0 == 224
                                if t31 {
                                    var t32 bool = second__1 < 160
                                    jp2 = t32
                                } else {
                                    jp2 = false
                                }
                            }
                            var jp3 bool
                            if jp2 {
                                jp3 = true
                            } else {
                                var t29 bool = first__0 == 237
                                if t29 {
                                    var t30 bool = second__1 >= 160
                                    jp3 = t30
                                } else {
                                    jp3 = false
                                }
                            }
                            if jp3 {
                                var inline16 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                                    _0: false,
                                    _1: 0,
                                    _2: 0,
                                }
                                return inline16
                            } else {
                                var t22 uint32 = first__0 & 15
                                var t23 uint32 = t22 << 12
                                var t24 uint32 = second__1 & 63
                                var t25 uint32 = t24 << 6
                                var t26 uint32 = t23 | t25
                                var t27 uint32 = third__0 & 63
                                var t28 uint32 = t26 | t27
                                var inline17 int = 3
                                var inline18 Option__char = __goml_builtin_char_from_uint32(t28)
                                switch inline18 != Option__char(0) {
                                case false:
                                    var inline19 Tuple3_4bool_4char_3int = utf8_invalid_decode()
                                    return inline19
                                case true:
                                    var inline20 rune = rune(uint64(inline18) - 1)
                                    var inline21 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                                        _0: true,
                                        _1: inline20,
                                        _2: inline17,
                                    }
                                    return inline21
                                default:
                                    panic("non-exhaustive match")
                                }
                            }
                        }
                    } else {
                        var t33 bool = first__0 < 245
                        if t33 {
                            var t34 int = length__0 - index__0
                            var t35 bool = t34 < 4
                            if t35 {
                                var t61 Tuple3_4bool_4char_3int = utf8_invalid_decode()
                                return t61
                            } else {
                                var t36 int = index__0 + 1
                                var t37 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, t36)
                                var second__2 uint32 = uint32(uint8(t37))
                                var t38 int = index__0 + 2
                                var t39 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, t38)
                                var third__1 uint32 = uint32(uint8(t39))
                                var t40 int = index__0 + 3
                                var t41 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, t40)
                                var fourth__0 uint32 = uint32(uint8(t41))
                                var t42 bool = utf8_invalid_continuation(second__2)
                                var jp4 bool
                                if t42 {
                                    jp4 = true
                                } else {
                                    var t60 bool = utf8_invalid_continuation(third__1)
                                    jp4 = t60
                                }
                                var jp5 bool
                                if jp4 {
                                    jp5 = true
                                } else {
                                    var t59 bool = utf8_invalid_continuation(fourth__0)
                                    jp5 = t59
                                }
                                var jp6 bool
                                if jp5 {
                                    jp6 = true
                                } else {
                                    var t57 bool = first__0 == 240
                                    if t57 {
                                        var t58 bool = second__2 < 144
                                        jp6 = t58
                                    } else {
                                        jp6 = false
                                    }
                                }
                                var jp7 bool
                                if jp6 {
                                    jp7 = true
                                } else {
                                    var t55 bool = first__0 == 244
                                    if t55 {
                                        var t56 bool = second__2 > 143
                                        jp7 = t56
                                    } else {
                                        jp7 = false
                                    }
                                }
                                if jp7 {
                                    var t43 Tuple3_4bool_4char_3int = utf8_invalid_decode()
                                    return t43
                                } else {
                                    var t44 uint32 = first__0 & 7
                                    var t45 uint32 = t44 << 18
                                    var t46 uint32 = second__2 & 63
                                    var t47 uint32 = t46 << 12
                                    var t48 uint32 = t45 | t47
                                    var t49 uint32 = third__1 & 63
                                    var t50 uint32 = t49 << 6
                                    var t51 uint32 = t48 | t50
                                    var t52 uint32 = fourth__0 & 63
                                    var t53 uint32 = t51 | t52
                                    var t54 Tuple3_4bool_4char_3int = utf8_valid_decode(t53, 4)
                                    return t54
                                }
                            }
                        } else {
                            var t62 Tuple3_4bool_4char_3int = utf8_invalid_decode()
                            return t62
                        }
                    }
                }
            }
        }
    }
}

func _goml_m_trait__impl_i_ToString_i_bool_i_to__string(self__0 bool) string {
    var t0 string = _goml_runtime_core_bool_to_string(self__0)
    return t0
}

func utf8_invalid_decode() Tuple3_4bool_4char_3int {
    var t0 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
        _0: false,
        _1: 0,
        _2: 0,
    }
    return t0
}

func _goml_m_inherent_i_string_i_string_i_byte__get(self__0 string, index__0 int) uint8 {
    var t0 uint8 = _goml_runtime_core_string_byte_get(self__0, index__0)
    return t0
}

func utf8_valid_decode(value__0 uint32, width__0 int) Tuple3_4bool_4char_3int {
    var commute_field0 rune
    var inline1 bool = utf8_valid_scalar(value__0)
    if inline1 {
        var inline2 Tuple2_4bool_4char = _goml_runtime_core_char_from_uint32(value__0)
        var inline3 rune = inline2._1
        commute_field0 = inline3
        var t0 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
            _0: true,
            _1: commute_field0,
            _2: width__0,
        }
        return t0
    } else {
        var inline0 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
            _0: false,
            _1: 0,
            _2: 0,
        }
        return inline0
    }
}

func utf8_invalid_continuation(value__0 uint32) bool {
    var t0 bool = value__0 < 128
    if t0 {
        return true
    } else {
        var t1 bool = value__0 > 191
        return t1
    }
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

func utf8_valid_scalar(value__0 uint32) bool {
    var t0 bool = value__0 <= 1114111
    if t0 {
        var t1 bool = value__0 >= 55296
        var jp0 bool
        if t1 {
            var t3 bool = value__0 <= 57343
            jp0 = t3
        } else {
            jp0 = false
        }
        var t2 bool = !jp0
        return t2
    } else {
        return false
    }
}

func main() {
    main0()
}
