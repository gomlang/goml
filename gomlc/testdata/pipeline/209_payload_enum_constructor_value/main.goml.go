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

type closure_env_cons_0 struct {}

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

func sum(values__0 List__i32) int32 {
    switch _goml_enum_tag_List__i32(values__0) {
    case 0:
        return 0
    case 1:
        var x0 int32 = values__0._node._p0
        var x1 List__i32 = values__0._node._p1
        var t0 int32 = sum(x1)
        var t1 int32 = x0 + t0
        return t1
    default:
        panic("non-exhaustive match")
    }
}

func main0() struct{} {
    var t0 closure_env_cons_0 = closure_env_cons_0{}
    var cons__0 func(int32, List__i32) List__i32 = func(p0 int32, p1 List__i32) List__i32 {
        return _goml_m_inherent_i_closure__env__cons__0_i_closure__env__cons__0_i_apply(t0, p0, p1)
    }
    var t1 List__i32
    var inline5 int32 = 2
    var inline6 List__i32 = cons__0(inline5, List__i32{})
    t1 = inline6
    var values__0 List__i32
    var inline3 int32 = 1
    var inline4 List__i32 = cons__0(inline3, t1)
    values__0 = inline4
    var t2 int32 = sum(values__0)
    var t3 string
    var inline2 string = __goml_builtin_int32_to_string(t2)
    t3 = inline2
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t3)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
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

func _goml_m_inherent_i_closure__env__cons__0_i_closure__env__cons__0_i_apply(env0 closure_env_cons_0, ctor_arg_0 int32, ctor_arg_1 List__i32) List__i32 {
    var t0 List__i32 = List__i32{
        _node: &List__i32_node{
            _p0: ctor_arg_0,
            _p1: ctor_arg_1,
            _tag: 1,
        },
    }
    return t0
}

func main() {
    main0()
}
