package main

import (
    _goml_os "os"
)

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

func show_scalar(value__0 uint32) struct{} {
    var mtmp0 Option__char
    var inline5 Option__char = __goml_builtin_char_from_uint32(value__0)
    mtmp0 = inline5
    switch mtmp0 != Option__char(0) {
    case false:
        var inline0 string = "none"
        var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
        _goml_runtime_core_string_println(inline1)
        return struct{}{}
    case true:
        var x0 rune = rune(uint64(mtmp0) - 1)
        var t0 uint32 = uint32(rune(x0))
        var inline3 string = _goml_m_trait__impl_i_ToString_i_u32_i_to__string(t0)
        _goml_runtime_core_string_println(inline3)
        return struct{}{}
    default:
        panic("non-exhaustive match")
    }
}

func main0() struct{} {
    show_scalar(0)
    show_scalar(65)
    show_scalar(55295)
    show_scalar(55296)
    var inline26 uint32 = 57343
    var inline27 Option__char = char_from_u32(inline26)
    switch inline27 != Option__char(0) {
    case false:
        println__T_string("none")
    case true:
        var inline29 rune = rune(uint64(inline27) - 1)
        var inline30 uint32 = uint32(rune(inline29))
        println__T_u32(inline30)
    default:
        panic("non-exhaustive match")
    }
    var inline20 uint32 = 57344
    var inline21 Option__char = char_from_u32(inline20)
    switch inline21 != Option__char(0) {
    case false:
        println__T_string("none")
    case true:
        var inline23 rune = rune(uint64(inline21) - 1)
        var inline24 uint32 = uint32(rune(inline23))
        println__T_u32(inline24)
    default:
        panic("non-exhaustive match")
    }
    var inline14 uint32 = 1114111
    var inline15 Option__char = char_from_u32(inline14)
    switch inline15 != Option__char(0) {
    case false:
        println__T_string("none")
    case true:
        var inline17 rune = rune(uint64(inline15) - 1)
        var inline18 uint32 = uint32(rune(inline17))
        println__T_u32(inline18)
    default:
        panic("non-exhaustive match")
    }
    var inline8 uint32 = 1114112
    var inline9 Option__char = char_from_u32(inline8)
    switch inline9 != Option__char(0) {
    case false:
        println__T_string("none")
    case true:
        var inline11 rune = rune(uint64(inline9) - 1)
        var inline12 uint32 = uint32(rune(inline11))
        println__T_u32(inline12)
    default:
        panic("non-exhaustive match")
    }
    var mtmp0 Option__char
    var inline6 uint32 = 128512
    var inline7 Option__char = __goml_builtin_char_from_uint32(inline6)
    mtmp0 = inline7
    switch mtmp0 != Option__char(0) {
    case false:
        var inline0 string = "none"
        var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
        _goml_runtime_core_string_println(inline1)
        return struct{}{}
    case true:
        var x0 rune = rune(uint64(mtmp0) - 1)
        var t0 string
        var inline5 string = char_to_string(x0)
        t0 = inline5
        var inline3 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t0)
        _goml_runtime_core_string_println(inline3)
        return struct{}{}
    default:
        panic("non-exhaustive match")
    }
}

func char_from_u32(value__0 uint32) Option__char {
    var inline0 bool = utf8_valid_scalar(value__0)
    if inline0 {
        var inline1 Tuple2_4bool_4char = _goml_runtime_core_char_from_uint32(value__0)
        var inline2 rune = inline1._1
        var inline3 Option__char = Option__char(uint64(uint32(inline2)) + 1)
        return inline3
    } else {
        return Option__char(0)
    }
}

func println__T_string(value__0 string) struct{} {
    var t0 string
    t0 = value__0
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

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func _goml_m_trait__impl_i_ToString_i_u32_i_to__string(self__0 uint32) string {
    var inline0 uint64 = uint64(uint32(self__0))
    var inline1 string = decimal_string(inline0)
    return inline1
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
