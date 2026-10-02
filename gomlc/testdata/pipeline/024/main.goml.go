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

type Line struct {
    from Point
    to Point
    color Color
}

type Ordering uint8

type Color uint8

const (
    Red Color = 0
    Green Color = 1
    Blue Color = 2
)

func line_to_string(l__0 Line) string {
    var x0 Point = l__0.from
    var x1 Point = l__0.to
    var x2 Color = l__0.color
    var t0 string
    var inline8 int32 = x0.x
    var inline9 int32 = x0.y
    var inline10 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline8)
    var inline11 string = "Point { x: " + inline10
    var inline12 string = inline11 + ", y: "
    var inline13 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline9)
    var inline14 string = inline12 + inline13
    var inline15 string = inline14 + " }"
    t0 = inline15
    var t1 string = "Line { from: " + t0
    var t2 string = t1 + ", to: "
    var t3 string
    var inline0 int32 = x1.x
    var inline1 int32 = x1.y
    var inline2 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline0)
    var inline3 string = "Point { x: " + inline2
    var inline4 string = inline3 + ", y: "
    var inline5 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline1)
    var inline6 string = inline4 + inline5
    var inline7 string = inline6 + " }"
    t3 = inline7
    var t4 string = t2 + t3
    var t5 string = t4 + ", color: "
    var t6 string
    switch x2 {
    case Red:
        t6 = "Red"
    case Green:
        t6 = "Green"
    case Blue:
        t6 = "Blue"
    default:
        panic("non-exhaustive match")
    }
    var t7 string = t5 + t6
    var t8 string = t7 + " }"
    return t8
}

func main0() struct{} {
    var p0__0 Point = Point{
        x: 0,
        y: 0,
    }
    var t0 string
    var inline4 int32 = 0
    var inline5 int32 = 0
    switch inline4 {
    case 0:
        switch inline5 {
        case 0:
            t0 = "origin"
        case 1:
            t0 = "up"
        default:
            var inline6 bool = 0 < inline5
            switch inline6 {
            case true:
                t0 = "above"
            case false:
                t0 = "below"
            default:
                panic("non-exhaustive match")
            }
        }
    case 1:
        switch inline5 {
        case 0:
            t0 = "right"
        default:
            t0 = "unknown"
        }
    default:
        t0 = "unknown"
    }
    var inline2 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t0)
    _goml_runtime_core_string_println(inline2)
    var p1__0 Point = Point{
        x: 10,
        y: 10,
    }
    var line__0 Line = Line{
        from: p0__0,
        to: p1__0,
        color: Red,
    }
    var t1 string = line_to_string(line__0)
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t1)
    _goml_runtime_core_string_println(inline0)
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
