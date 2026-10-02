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

type IntList struct {
    _node *IntList_node
}

type IntList_node struct {
    _p1 IntList
    _p0 int32
    _tag uint8
}

func _goml_enum_tag_IntList(value IntList) uint8 {
    if value._node == nil {
        return 0
    }
    return value._node._tag
}

func print_int_list(xs__0 IntList) struct{} {
    switch _goml_enum_tag_IntList(xs__0) {
    case 0:
        var inline0 string = "Nil"
        var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
        _goml_runtime_core_string_println(inline1)
        return struct{}{}
    case 1:
        var x0 int32 = xs__0._node._p0
        var x1 IntList = xs__0._node._p1
        var inline15 string = "Cons"
        var inline16 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline15)
        _goml_runtime_core_string_println(inline16)
        var inline12 string = "("
        var inline13 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline12)
        _goml_runtime_core_string_println(inline13)
        var t0 string
        var inline11 string = __goml_builtin_int32_to_string(x0)
        t0 = inline11
        var inline9 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t0)
        _goml_runtime_core_string_println(inline9)
        var inline6 string = ", "
        var inline7 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline6)
        _goml_runtime_core_string_println(inline7)
        print_int_list(x1)
        var inline3 string = ")"
        var inline4 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline3)
        _goml_runtime_core_string_println(inline4)
        return struct{}{}
    default:
        panic("non-exhaustive match")
    }
}

func int_list_rev_aux(xs__0 IntList, acc__0 IntList) IntList {
    switch _goml_enum_tag_IntList(xs__0) {
    case 0:
        return acc__0
    case 1:
        var x0 int32 = xs__0._node._p0
        var x1 IntList = xs__0._node._p1
        var t0 IntList = IntList{
            _node: &IntList_node{
                _p0: x0,
                _p1: acc__0,
                _tag: 1,
            },
        }
        var t1 IntList = int_list_rev_aux(x1, t0)
        return t1
    default:
        panic("non-exhaustive match")
    }
}

func int_list_length(xs__0 IntList) int32 {
    switch _goml_enum_tag_IntList(xs__0) {
    case 0:
        return 0
    case 1:
        var x0 IntList = xs__0._node._p1
        var t0 int32 = int_list_length(x0)
        var t1 int32 = 1 + t0
        return t1
    default:
        panic("non-exhaustive match")
    }
}

func main0() struct{} {
    var x__0 IntList = IntList{}
    print_int_list(x__0)
    var inline22 string = ""
    var inline23 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline22)
    _goml_runtime_core_string_println(inline23)
    println__T_string("Length: ")
    var inline19 int32 = int_list_length(x__0)
    var inline20 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline19)
    println__T_string(inline20)
    var x__1 IntList = IntList{
        _node: &IntList_node{
            _p0: 1,
            _p1: IntList{},
            _tag: 1,
        },
    }
    print_int_list(x__1)
    var inline15 string = ""
    var inline16 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline15)
    _goml_runtime_core_string_println(inline16)
    println__T_string("Length: ")
    var inline12 int32 = int_list_length(x__1)
    var inline13 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline12)
    println__T_string(inline13)
    var t0 IntList = IntList{
        _node: &IntList_node{
            _p0: 3,
            _p1: IntList{},
            _tag: 1,
        },
    }
    var t1 IntList = IntList{
        _node: &IntList_node{
            _p0: 2,
            _p1: t0,
            _tag: 1,
        },
    }
    var x__2 IntList = IntList{
        _node: &IntList_node{
            _p0: 1,
            _p1: t1,
            _tag: 1,
        },
    }
    print_int_list(x__2)
    var inline8 string = ""
    var inline9 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline8)
    _goml_runtime_core_string_println(inline9)
    println__T_string("Length: ")
    var inline5 int32 = int_list_length(x__2)
    var inline6 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline5)
    println__T_string(inline6)
    var y__0 IntList
    var inline3 IntList = int_list_rev_aux(x__2, IntList{})
    y__0 = inline3
    print_int_list(y__0)
    var inline0 string = ""
    var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
    _goml_runtime_core_string_println(inline1)
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
