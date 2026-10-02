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

type _goml_vec_int32 struct {
    items []int32
}

func vec_new__Vec_5int32() *_goml_vec_int32 {
    return &_goml_vec_int32{
        items: nil,
    }
}

func vec_push__Vec_5int32(vec *_goml_vec_int32, elem int32) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_len__Vec_5int32(vec *_goml_vec_int32) int {
    return int(len(vec.items))
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

type Marker struct {}

type Ordering uint8

type Shared struct {
    _p1 uint16
    _p0 uint8
    _tag uint8
}

type ZeroSized uint8

type MaybeByte uint16

type MaybeSigned uint64

type MaybeVec struct {
    _value *_goml_vec_int32
}

type Outer struct {
    _p0 Inner
    _tag uint8
}

type Inner struct {
    _p0 uint8
    _tag uint8
}

type List struct {
    _node *List_node
}

type List_node struct {
    _p1 List
    _p0 int32
    _tag uint8
}

func _goml_enum_tag_List(value List) uint8 {
    if value._node == nil {
        return 0
    }
    return value._node._tag
}

func shared_value(value__0 Shared) int {
    switch value__0._tag {
    case 0:
        var x0 uint8 = value__0._p0
        var t0 int = int(uint8(x0))
        return t0
    case 1:
        var x1 uint8 = value__0._p0
        var t1 int = int(uint8(x1))
        var t2 int = t1 + 10
        return t2
    case 2:
        var x2 uint16 = value__0._p1
        var t3 int = int(uint16(x2))
        var t4 int = t3 + 20
        return t4
    case 3:
        return 0
    default:
        panic("non-exhaustive match")
    }
}

func maybe_byte(value__0 MaybeByte) int {
    switch value__0 != MaybeByte(0) {
    case false:
        return 0
    case true:
        var x0 uint8 = uint8(uint16(value__0) - 1)
        var t0 int = int(uint8(x0))
        return t0
    default:
        panic("non-exhaustive match")
    }
}

func maybe_signed(value__0 MaybeSigned) int32 {
    switch value__0 != MaybeSigned(0) {
    case false:
        return 0
    case true:
        var x0 int32 = int32(int64(uint64(value__0) - 1) - 2147483648)
        return x0
    default:
        panic("non-exhaustive match")
    }
}

func list_sum(value__0 List) int32 {
    switch _goml_enum_tag_List(value__0) {
    case 0:
        return 0
    case 1:
        var x0 int32 = value__0._node._p0
        var x1 List = value__0._node._p1
        var t0 int32 = list_sum(x1)
        var t1 int32 = x0 + t0
        return t1
    default:
        panic("non-exhaustive match")
    }
}

func main0() struct{} {
    var t0 Shared = Shared{
        _p0: 3,
        _tag: 0,
    }
    var t1 int = shared_value(t0)
    println__T_isize(t1)
    var t2 Shared = Shared{
        _p0: 4,
        _tag: 1,
    }
    var t3 int = shared_value(t2)
    println__T_isize(t3)
    var t4 Shared = Shared{
        _p1: 5,
        _tag: 2,
    }
    var t5 int = shared_value(t4)
    println__T_isize(t5)
    var t6 MaybeByte = MaybeByte(uint16(uint8(255)) + 1)
    var t7 int = maybe_byte(t6)
    println__T_isize(t7)
    var t8 MaybeSigned = MaybeSigned(uint64(int64(-7) + 2147483648) + 1)
    var t9 int32 = maybe_signed(t8)
    var inline19 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(t9)
    _goml_runtime_core_string_println(inline19)
    var t10 int32
    var inline18 int32 = -2147483648
    t10 = inline18
    var inline16 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(t10)
    _goml_runtime_core_string_println(inline16)
    var values__0 *_goml_vec_int32
    var inline15 *_goml_vec_int32 = vec_new__Vec_5int32()
    values__0 = inline15
    var inline13 int32 = 7
    vec_push__Vec_5int32(values__0, inline13)
    var t11 int
    var inline12 int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_len____T__i32(values__0)
    t11 = inline12
    var inline10 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t11)
    _goml_runtime_core_string_println(inline10)
    var t12 int
    var inline7 uint8 = 8
    var inline8 int = int(uint8(inline7))
    var inline9 int = inline8 + 1
    t12 = inline9
    var inline5 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t12)
    _goml_runtime_core_string_println(inline5)
    var t13 List = List{
        _node: &List_node{
            _p0: 2,
            _p1: List{},
            _tag: 1,
        },
    }
    var t14 List = List{
        _node: &List_node{
            _p0: 1,
            _p1: t13,
            _tag: 1,
        },
    }
    var t15 int32 = list_sum(t14)
    var inline3 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(t15)
    _goml_runtime_core_string_println(inline3)
    var inline0 int = 1
    var inline1 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(inline0)
    _goml_runtime_core_string_println(inline1)
    return struct{}{}
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_len____T__i32(self__0 *_goml_vec_int32) int {
    var t0 int = vec_len__Vec_5int32(self__0)
    return t0
}

func println__T_isize(value__0 int) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_int_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_isize_i_to__string(self__0 int) string {
    var inline0 int64 = int64(int(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func _goml_m_trait__impl_i_ToString_i_i32_i_to__string(self__0 int32) string {
    var inline0 int64 = int64(int32(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
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
