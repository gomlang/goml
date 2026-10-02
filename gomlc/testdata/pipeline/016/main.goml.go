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

type List__i32 struct {
    _node *List__i32_node
}

type List__i32_node struct {
    _p1 List__i32
    _p0 int32
    _tag uint8
}

func _goml_enum_tag_List__i32(value List__i32) uint8 {
    if value._node == nil {
        return 0
    }
    return value._node._tag
}

type List__isize struct {
    _node *List__isize_node
}

type List__isize_node struct {
    _p0 int
    _p1 List__isize
    _tag uint8
}

func _goml_enum_tag_List__isize(value List__isize) uint8 {
    if value._node == nil {
        return 0
    }
    return value._node._tag
}

type _goml_m_List_____o__q_ struct {
    _node *_goml_m_List_____o__q__node
}

type _goml_m_List_____o__q__node struct {
    _p0 _goml_m_List_____o__q_
    _tag uint8
}

func _goml_enum_tag__goml_m_List_____o__q_(value _goml_m_List_____o__q_) uint8 {
    if value._node == nil {
        return 0
    }
    return value._node._tag
}

type List__bool struct {
    _node *List__bool_node
}

type List__bool_node struct {
    _p1 List__bool
    _p0 bool
    _tag uint8
}

func _goml_enum_tag_List__bool(value List__bool) uint8 {
    if value._node == nil {
        return 0
    }
    return value._node._tag
}

func int_list_length(xs__0 List__i32) int32 {
    switch _goml_enum_tag_List__i32(xs__0) {
    case 0:
        return 0
    case 1:
        var x0 List__i32 = xs__0._node._p1
        var t0 int32 = int_list_length(x0)
        var t1 int32 = 1 + t0
        return t1
    default:
        panic("non-exhaustive match")
    }
}

func main0() struct{} {
    var x__0 List__isize = List__isize{
        _node: &List__isize_node{
            _p0: 1,
            _p1: List__isize{},
            _tag: 1,
        },
    }
    var length__0 int32 = list_length__T_isize(x__0)
    var inline10 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(length__0)
    _goml_runtime_core_string_println(inline10)
    var t0 List__isize = List__isize{
        _node: &List__isize_node{
            _p0: 2,
            _p1: List__isize{},
            _tag: 1,
        },
    }
    var x__1 List__isize = List__isize{
        _node: &List__isize_node{
            _p0: 1,
            _p1: t0,
            _tag: 1,
        },
    }
    var length__1 int32 = list_length__T_isize(x__1)
    var inline8 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(length__1)
    _goml_runtime_core_string_println(inline8)
    var t1 List__i32 = List__i32{
        _node: &List__i32_node{
            _p0: 2,
            _p1: List__i32{},
            _tag: 1,
        },
    }
    var t2 List__i32 = List__i32{
        _node: &List__i32_node{
            _p0: 1,
            _p1: t1,
            _tag: 1,
        },
    }
    var x__2 List__i32 = List__i32{
        _node: &List__i32_node{
            _p0: 0,
            _p1: t2,
            _tag: 1,
        },
    }
    var length__2 int32 = int_list_length(x__2)
    var inline6 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(length__2)
    _goml_runtime_core_string_println(inline6)
    var x__3 _goml_m_List_____o__q_ = _goml_m_List_____o__q_{
        _node: &_goml_m_List_____o__q__node{
            _p0: _goml_m_List_____o__q_{},
            _tag: 1,
        },
    }
    var length__3 int32 = _goml_m_list__length____T___o__q_(x__3)
    var inline4 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(length__3)
    _goml_runtime_core_string_println(inline4)
    var t3 _goml_m_List_____o__q_ = _goml_m_List_____o__q_{
        _node: &_goml_m_List_____o__q__node{
            _p0: _goml_m_List_____o__q_{},
            _tag: 1,
        },
    }
    var x__4 _goml_m_List_____o__q_ = _goml_m_List_____o__q_{
        _node: &_goml_m_List_____o__q__node{
            _p0: t3,
            _tag: 1,
        },
    }
    var length__4 int32 = _goml_m_list__length____T___o__q_(x__4)
    var inline2 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(length__4)
    _goml_runtime_core_string_println(inline2)
    var t4 List__bool = List__bool{
        _node: &List__bool_node{
            _p0: false,
            _p1: List__bool{},
            _tag: 1,
        },
    }
    var x__5 List__bool = List__bool{
        _node: &List__bool_node{
            _p0: true,
            _p1: t4,
            _tag: 1,
        },
    }
    var length__5 int32 = list_length__T_bool(x__5)
    var inline0 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(length__5)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func list_length__T_isize(xs__0 List__isize) int32 {
    switch _goml_enum_tag_List__isize(xs__0) {
    case 0:
        return 0
    case 1:
        var x0 List__isize = xs__0._node._p1
        var t0 int32 = list_length__T_isize(x0)
        var t1 int32 = 1 + t0
        return t1
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_list__length____T___o__q_(xs__0 _goml_m_List_____o__q_) int32 {
    switch _goml_enum_tag__goml_m_List_____o__q_(xs__0) {
    case 0:
        return 0
    case 1:
        var x0 _goml_m_List_____o__q_ = xs__0._node._p0
        var t0 int32 = _goml_m_list__length____T___o__q_(x0)
        var t1 int32 = 1 + t0
        return t1
    default:
        panic("non-exhaustive match")
    }
}

func list_length__T_bool(xs__0 List__bool) int32 {
    switch _goml_enum_tag_List__bool(xs__0) {
    case 0:
        return 0
    case 1:
        var x0 List__bool = xs__0._node._p1
        var t0 int32 = list_length__T_bool(x0)
        var t1 int32 = 1 + t0
        return t1
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_trait__impl_i_ToString_i_i32_i_to__string(self__0 int32) string {
    var inline0 int64 = int64(int32(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
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
