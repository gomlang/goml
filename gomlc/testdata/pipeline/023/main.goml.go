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

type Tuple2_5int32_6string struct {
    _0 int32
    _1 string
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

type Mixed struct {
    _p1 string
    _p0 int32
    _tag uint8
}

func match_mixed_pair(pair__0 Tuple2_5int32_6string) int32 {
    var x0 int32 = pair__0._0
    var x1 string = pair__0._1
    switch x1 {
    case "zero":
        switch x0 {
        case 0:
            return 1
        default:
            return 4
        }
    case "one":
        switch x0 {
        case 0:
            return 2
        case 1:
            return 3
        default:
            return 5
        }
    default:
        switch x0 {
        case 0:
            return 2
        default:
            return 5
        }
    }
}

func match_mixed_enum(value__0 Mixed) int32 {
    switch value__0._tag {
    case 0:
        var x0 int32 = value__0._p0
        switch x0 {
        case 0:
            return 6
        default:
            return 7
        }
    case 1:
        var x1 string = value__0._p1
        switch x1 {
        case "zero":
            return 8
        default:
            return 9
        }
    case 2:
        var x2 int32 = value__0._p0
        var x3 string = value__0._p1
        switch x3 {
        case "zero":
            switch x2 {
            case 0:
                return 10
            default:
                return 12
            }
        default:
            switch x2 {
            case 0:
                return 11
            default:
                return 13
            }
        }
    default:
        panic("non-exhaustive match")
    }
}

func main0() struct{} {
    var t0 Tuple2_5int32_6string = Tuple2_5int32_6string{
        _0: 0,
        _1: "zero",
    }
    var t1 int32 = match_mixed_pair(t0)
    var t2 string = _goml_m_inherent_i_i32_i_i32_i_to__string(t1)
    println__T_string(t2)
    var t3 Tuple2_5int32_6string = Tuple2_5int32_6string{
        _0: 0,
        _1: "other",
    }
    var t4 int32 = match_mixed_pair(t3)
    var t5 string = _goml_m_inherent_i_i32_i_i32_i_to__string(t4)
    println__T_string(t5)
    var t6 Tuple2_5int32_6string = Tuple2_5int32_6string{
        _0: 1,
        _1: "one",
    }
    var t7 int32 = match_mixed_pair(t6)
    var t8 string = _goml_m_inherent_i_i32_i_i32_i_to__string(t7)
    println__T_string(t8)
    var t9 Tuple2_5int32_6string = Tuple2_5int32_6string{
        _0: 2,
        _1: "zero",
    }
    var t10 int32 = match_mixed_pair(t9)
    var t11 string = _goml_m_inherent_i_i32_i_i32_i_to__string(t10)
    var inline27 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t11)
    _goml_runtime_core_string_println(inline27)
    var t12 Tuple2_5int32_6string = Tuple2_5int32_6string{
        _0: 2,
        _1: "two",
    }
    var t13 int32 = match_mixed_pair(t12)
    var t14 string
    var inline26 string = __goml_builtin_int32_to_string(t13)
    t14 = inline26
    var inline24 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t14)
    _goml_runtime_core_string_println(inline24)
    var t15 Mixed = Mixed{
        _p0: 0,
        _tag: 0,
    }
    var t16 int32 = match_mixed_enum(t15)
    var t17 string
    var inline23 string = __goml_builtin_int32_to_string(t16)
    t17 = inline23
    var inline21 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t17)
    _goml_runtime_core_string_println(inline21)
    var t18 Mixed = Mixed{
        _p0: 5,
        _tag: 0,
    }
    var t19 int32 = match_mixed_enum(t18)
    var t20 string
    var inline20 string = __goml_builtin_int32_to_string(t19)
    t20 = inline20
    var inline18 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t20)
    _goml_runtime_core_string_println(inline18)
    var t21 Mixed = Mixed{
        _p1: "zero",
        _tag: 1,
    }
    var t22 int32 = match_mixed_enum(t21)
    var t23 string
    var inline17 string = __goml_builtin_int32_to_string(t22)
    t23 = inline17
    var inline15 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t23)
    _goml_runtime_core_string_println(inline15)
    var t24 Mixed = Mixed{
        _p1: "hello",
        _tag: 1,
    }
    var t25 int32 = match_mixed_enum(t24)
    var t26 string
    var inline14 string = __goml_builtin_int32_to_string(t25)
    t26 = inline14
    var inline12 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t26)
    _goml_runtime_core_string_println(inline12)
    var t27 Mixed = Mixed{
        _p0: 0,
        _p1: "zero",
        _tag: 2,
    }
    var t28 int32 = match_mixed_enum(t27)
    var t29 string
    var inline11 string = __goml_builtin_int32_to_string(t28)
    t29 = inline11
    var inline9 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t29)
    _goml_runtime_core_string_println(inline9)
    var t30 Mixed = Mixed{
        _p0: 0,
        _p1: "hello",
        _tag: 2,
    }
    var t31 int32 = match_mixed_enum(t30)
    var t32 string
    var inline8 string = __goml_builtin_int32_to_string(t31)
    t32 = inline8
    var inline6 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t32)
    _goml_runtime_core_string_println(inline6)
    var t33 Mixed = Mixed{
        _p0: 2,
        _p1: "zero",
        _tag: 2,
    }
    var t34 int32 = match_mixed_enum(t33)
    var t35 string
    var inline5 string = __goml_builtin_int32_to_string(t34)
    t35 = inline5
    var inline3 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t35)
    _goml_runtime_core_string_println(inline3)
    var t36 Mixed = Mixed{
        _p0: 3,
        _p1: "three",
        _tag: 2,
    }
    var t37 int32 = match_mixed_enum(t36)
    var t38 string
    var inline2 string = __goml_builtin_int32_to_string(t37)
    t38 = inline2
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t38)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func println__T_string(value__0 string) struct{} {
    var t0 string
    t0 = value__0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_inherent_i_i32_i_i32_i_to__string(self__0 int32) string {
    var inline0 int64 = int64(int32(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
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
