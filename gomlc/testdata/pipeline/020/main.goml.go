package main

import (
    _goml_os "os"
)

func _goml_runtime_core_unit_to_string(x struct{}) string {
    return "()"
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

type Point struct {
    x int32
    y int32
}

type Wrapper__i32 struct {
    value int32
}

type _goml_m_Wrapper_____o__q_ struct {
    value struct{}
}

type Ordering uint8

type Shape__i32 struct {
    _p0 Point
    _p1 Wrapper__i32
    _tag uint8
}

type _goml_m_Shape_____o__q_ struct {
    _p0 Point
    _tag uint8
}

func bounce_int(shape__0 Shape__i32) Shape__i32 {
    switch shape__0._tag {
    case 0:
        var x0 Point = shape__0._p0
        var t0 Shape__i32 = Shape__i32{
            _p0: x0,
            _tag: 0,
        }
        return t0
    case 1:
        var x1 Wrapper__i32 = shape__0._p1
        var t1 Shape__i32 = Shape__i32{
            _p1: x1,
            _tag: 1,
        }
        return t1
    case 2:
        return Shape__i32{
            _tag: 2,
        }
    default:
        panic("non-exhaustive match")
    }
}

func point32_to_string(point__0 Point) string {
    var x0 int32 = point__0.x
    var x1 int32 = point__0.y
    var t0 string
    var inline1 string = __goml_builtin_int32_to_string(x0)
    t0 = inline1
    var with_x__0 string = "Point { x: " + t0
    var with_y_label__0 string = with_x__0 + ", y: "
    var t1 string
    var inline0 string = __goml_builtin_int32_to_string(x1)
    t1 = inline0
    var with_y__0 string = with_y_label__0 + t1
    var t2 string = with_y__0 + " }"
    return t2
}

func wrapper_int32_to_string(wrapper__0 Wrapper__i32) string {
    var x0 int32 = wrapper__0.value
    var t0 string
    var inline0 string = __goml_builtin_int32_to_string(x0)
    t0 = inline0
    var prefix__0 string = "Wrapper[i32] { value: " + t0
    var t1 string = prefix__0 + " }"
    return t1
}

func wrapper_unit_to_string(wrapper__0 _goml_m_Wrapper_____o__q_) string {
    var x0 struct{} = wrapper__0.value
    var t0 string
    var inline0 string = _goml_runtime_core_unit_to_string(x0)
    t0 = inline0
    var prefix__0 string = "Wrapper[()] { value: " + t0
    var t1 string = prefix__0 + " }"
    return t1
}

func shape_int32_to_string(shape__0 Shape__i32) string {
    switch shape__0._tag {
    case 0:
        var x0 Point = shape__0._p0
        var t0 string
        var inline0 int32 = x0.x
        var inline1 int32 = x0.y
        var inline2 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline0)
        var inline3 string = "Point { x: " + inline2
        var inline4 string = inline3 + ", y: "
        var inline5 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline1)
        var inline6 string = inline4 + inline5
        var inline7 string = inline6 + " }"
        t0 = inline7
        var prefix__0 string = "Shape::Dot(" + t0
        var t1 string = prefix__0 + ")"
        return t1
    case 1:
        var x1 Wrapper__i32 = shape__0._p1
        var t2 string
        var inline8 int32 = x1.value
        var inline9 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline8)
        var inline10 string = "Wrapper[i32] { value: " + inline9
        var inline11 string = inline10 + " }"
        t2 = inline11
        var prefix__1 string = "Shape::Wrapped(" + t2
        var t3 string = prefix__1 + ")"
        return t3
    case 2:
        return "Shape::Origin"
    default:
        panic("non-exhaustive match")
    }
}

func shape_unit_to_string(shape__0 _goml_m_Shape_____o__q_) string {
    switch shape__0._tag {
    case 0:
        var x0 Point = shape__0._p0
        var t0 string
        var inline0 int32 = x0.x
        var inline1 int32 = x0.y
        var inline2 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline0)
        var inline3 string = "Point { x: " + inline2
        var inline4 string = inline3 + ", y: "
        var inline5 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline1)
        var inline6 string = inline4 + inline5
        var inline7 string = inline6 + " }"
        t0 = inline7
        var prefix__0 string = "Shape::Dot(" + t0
        var t1 string = prefix__0 + ")"
        return t1
    case 1:
        var x1 _goml_m_Wrapper_____o__q_ = _goml_m_Wrapper_____o__q_{}
        var t2 string
        var inline8 struct{} = x1.value
        var inline9 string = _goml_m_trait__impl_i_ToString_i__o__q__i_to__string(inline8)
        var inline10 string = "Wrapper[()] { value: " + inline9
        var inline11 string = inline10 + " }"
        t2 = inline11
        var prefix__1 string = "Shape::Wrapped(" + t2
        var t3 string = prefix__1 + ")"
        return t3
    case 2:
        return "Shape::Origin"
    default:
        panic("non-exhaustive match")
    }
}

func main0() struct{} {
    var t0 Point = Point{
        x: 3,
        y: 4,
    }
    var t1 string = point32_to_string(t0)
    println__T_string(t1)
    var t2 Wrapper__i32 = Wrapper__i32{
        value: 7,
    }
    var t3 string = wrapper_int32_to_string(t2)
    println__T_string(t3)
    var t4 _goml_m_Wrapper_____o__q_ = _goml_m_Wrapper_____o__q_{
        value: struct{}{},
    }
    var t5 string = wrapper_unit_to_string(t4)
    println__T_string(t5)
    var bounced_origin__0 Shape__i32 = bounce_int(Shape__i32{
        _tag: 2,
    })
    var t6 Point = Point{
        x: 3,
        y: 4,
    }
    var t7 Shape__i32 = Shape__i32{
        _p0: t6,
        _tag: 0,
    }
    var t8 string = shape_int32_to_string(t7)
    println__T_string(t8)
    var t9 Wrapper__i32 = Wrapper__i32{
        value: 7,
    }
    var t10 Shape__i32 = Shape__i32{
        _p1: t9,
        _tag: 1,
    }
    var t11 string = shape_int32_to_string(t10)
    var inline11 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t11)
    _goml_runtime_core_string_println(inline11)
    var t12 string = shape_int32_to_string(bounced_origin__0)
    var inline9 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t12)
    _goml_runtime_core_string_println(inline9)
    var t13 Point = Point{
        x: 3,
        y: 4,
    }
    var t14 _goml_m_Shape_____o__q_ = _goml_m_Shape_____o__q_{
        _p0: t13,
        _tag: 0,
    }
    var t15 string = shape_unit_to_string(t14)
    var inline7 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t15)
    _goml_runtime_core_string_println(inline7)
    var t17 _goml_m_Shape_____o__q_ = _goml_m_Shape_____o__q_{
        _tag: 1,
    }
    var t18 string = shape_unit_to_string(t17)
    var inline5 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t18)
    _goml_runtime_core_string_println(inline5)
    var t19 string
    t19 = "Shape::Origin"
    var inline3 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t19)
    _goml_runtime_core_string_println(inline3)
    var t20 Shape__i32
    t20 = Shape__i32{
        _tag: 2,
    }
    switch t20._tag {
    case 0:
    case 1:
    case 2:
    default:
        panic("non-exhaustive match")
    }
    var inline0 string = "struct enums!"
    var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
    _goml_runtime_core_string_println(inline1)
    return struct{}{}
}

func _goml_m_inherent_i_i32_i_i32_i_to__string(self__0 int32) string {
    var inline0 int64 = int64(int32(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func _goml_m_trait__impl_i_ToString_i__o__q__i_to__string(self__0 struct{}) string {
    var t0 string = _goml_runtime_core_unit_to_string(self__0)
    return t0
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
