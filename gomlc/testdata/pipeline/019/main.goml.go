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

type Point struct {
    x int32
    y int32
}

type Wrapper__i32 struct {
    value int32
}

type Ordering uint8

func make_point() Point {
    var t0 Point = Point{
        x: 0,
        y: 0,
    }
    return t0
}

func flip(point__0 Point) Point {
    var x0 int32 = point__0.x
    var x1 int32 = point__0.y
    var t0 Point = Point{
        x: x1,
        y: x0,
    }
    return t0
}

func x_add_1(p__0 Point) Point {
    var x0 int32 = p__0.x
    var x1 int32 = p__0.y
    var t0 int32 = x0 + 1
    var t1 Point = Point{
        x: t0,
        y: x1,
    }
    return t1
}

func point32_to_string(p__0 Point) string {
    var x0 int32 = p__0.x
    var x1 int32 = p__0.y
    var t0 string
    var inline1 string = __goml_builtin_int32_to_string(x0)
    t0 = inline1
    var t1 string = "Point { x: " + t0
    var t2 string = t1 + ", y: "
    var t3 string
    var inline0 string = __goml_builtin_int32_to_string(x1)
    t3 = inline0
    var t4 string = t2 + t3
    var t5 string = t4 + "}"
    return t5
}

func point32_to_string2(p__0 Point) string {
    var x0 int32 = p__0.x
    var x1 int32 = p__0.y
    var t0 string
    var inline1 string = __goml_builtin_int32_to_string(x0)
    t0 = inline1
    var t1 string = "Point { x: " + t0
    var t2 string = t1 + ", y: "
    var t3 string
    var inline0 string = __goml_builtin_int32_to_string(x1)
    t3 = inline0
    var t4 string = t2 + t3
    var t5 string = t4 + "}"
    return t5
}

func main0() struct{} {
    var start__0 Point = make_point()
    var t0 string = point32_to_string(start__0)
    println__T_string(t0)
    var t1 Point = Point{
        x: 1,
        y: 2,
    }
    var swapped__0 Point = flip(t1)
    var t2 string = point32_to_string2(swapped__0)
    println__T_string(t2)
    var a__0 Point = x_add_1(start__0)
    var t3 string
    var inline19 int32 = a__0.x
    var inline20 int32 = a__0.y
    var inline21 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline19)
    var inline22 string = "Point { x: " + inline21
    var inline23 string = inline22 + ", y: "
    var inline24 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline20)
    var inline25 string = inline23 + inline24
    var inline26 string = inline25 + "}"
    t3 = inline26
    var inline17 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t3)
    _goml_runtime_core_string_println(inline17)
    var t4 Point
    var inline13 int32 = start__0.x
    var inline14 int32 = start__0.y
    var inline15 int32 = inline13 + 1
    var inline16 Point = Point{
        x: inline15,
        y: inline14,
    }
    t4 = inline16
    var a__1 Point
    var inline10 int32 = t4.x
    var inline11 int32 = t4.y
    var inline12 Point = Point{
        x: inline11,
        y: inline10,
    }
    a__1 = inline12
    var t5 string
    var inline2 int32 = a__1.x
    var inline3 int32 = a__1.y
    var inline4 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline2)
    var inline5 string = "Point { x: " + inline4
    var inline6 string = inline5 + ", y: "
    var inline7 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline3)
    var inline8 string = inline6 + inline7
    var inline9 string = inline8 + "}"
    t5 = inline9
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t5)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_inherent_i_i32_i_i32_i_to__string(self__0 int32) string {
    var inline0 int64 = int64(int32(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func println__T_string(value__0 string) struct{} {
    var t0 string
    t0 = value__0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
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

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
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
