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

type closure_env_f_0 struct {
    y_0 int32
    z_1 int32
}

type closure_env_add_base_1 struct {
    base_0 int32
}

type closure_env_printer_2 struct {}

type closure_env_no_capture_3 struct {}

type closure_env_play_list_and_point_4 struct {
    list123_0 IntList
    point_1 Point
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

func main0() struct{} {
    var base__0 int32 = 5
    var t0 closure_env_add_base_1 = closure_env_add_base_1{
        base_0: base__0,
    }
    var add_base__0 func(int32) int32 = func(p0 int32) int32 {
        return _goml_m_inherent_i_closure__en_he0443db01d9642f998bbe31aa65b7a79_base__1_i_apply(t0, p0)
    }
    var result__0 int32 = add_base__0(7)
    var t1 closure_env_printer_2 = closure_env_printer_2{}
    var printer__0 func(string, int32) struct{} = func(p0 string, p1 int32) struct{} {
        return _goml_m_inherent_i_closure__env__printer__2_i_closure__env__printer__2_i_apply(t1, p0, p1)
    }
    printer__0("result: ", result__0)
    var t2 closure_env_no_capture_3 = closure_env_no_capture_3{}
    var no_capture__0 func(int32) int32 = func(p0 int32) int32 {
        return _goml_m_inherent_i_closure__en_ha1a1d2e736bef56b17edec01979b6eae_ture__3_i_apply(t2, p0)
    }
    var doubled__0 int32 = no_capture__0(3)
    var t3 string
    var inline12 string = __goml_builtin_int32_to_string(doubled__0)
    t3 = inline12
    var inline10 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t3)
    _goml_runtime_core_string_println(inline10)
    var inline0 int32 = 3
    var inline1 int32 = 5
    var inline2 closure_env_f_0 = closure_env_f_0{
        y_0: inline0,
        z_1: inline1,
    }
    var inline3 func(int32) int32 = func(p0 int32) int32 {
        return _goml_m_inherent_i_closure__env__f__0_i_closure__env__f__0_i_apply(inline2, p0)
    }
    var inline4 int32 = inline3(2)
    var inline5 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline4)
    println__T_string(inline5)
    var inline7 int32 = inline3(3)
    var inline8 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline7)
    println__T_string(inline8)
    var t4 IntList = IntList{
        _node: &IntList_node{
            _p0: 3,
            _p1: IntList{},
            _tag: 1,
        },
    }
    var t5 IntList = IntList{
        _node: &IntList_node{
            _p0: 2,
            _p1: t4,
            _tag: 1,
        },
    }
    var list123__0 IntList = IntList{
        _node: &IntList_node{
            _p0: 1,
            _p1: t5,
            _tag: 1,
        },
    }
    var point__0 Point = Point{
        x: 10,
        y: 20,
    }
    var t6 closure_env_play_list_and_point_4 = closure_env_play_list_and_point_4{
        list123_0: list123__0,
        point_1: point__0,
    }
    var play_list_and_point__0 func() struct{} = func() struct{} {
        return _goml_m_inherent_i_closure__en_h905154c8b1f2c223fea35335436a9056_oint__4_i_apply(t6)
    }
    play_list_and_point__0()
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

func _goml_m_inherent_i_closure__env__f__0_i_closure__env__f__0_i_apply(env0 closure_env_f_0, x__0 int32) int32 {
    var y__0 int32 = env0.y_0
    var z__0 int32 = env0.z_1
    var t0 int32 = x__0 * y__0
    var t1 int32 = t0 * z__0
    return t1
}

func _goml_m_inherent_i_closure__en_he0443db01d9642f998bbe31aa65b7a79_base__1_i_apply(env0 closure_env_add_base_1, x__0 int32) int32 {
    var base__0 int32 = env0.base_0
    var t0 int32 = x__0 + base__0
    return t0
}

func _goml_m_inherent_i_closure__env__printer__2_i_closure__env__printer__2_i_apply(env0 closure_env_printer_2, prefix__0 string, value__0 int32) struct{} {
    var t0 string
    var inline2 string = __goml_builtin_int32_to_string(value__0)
    t0 = inline2
    var message__0 string = prefix__0 + t0
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(message__0)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_inherent_i_closure__en_ha1a1d2e736bef56b17edec01979b6eae_ture__3_i_apply(env0 closure_env_no_capture_3, z__0 int32) int32 {
    var t0 int32 = z__0 * 2
    return t0
}

func _goml_m_inherent_i_closure__en_h905154c8b1f2c223fea35335436a9056_oint__4_i_apply(env0 closure_env_play_list_and_point_4) struct{} {
    var list123__0 IntList = env0.list123_0
    var point__0 Point = env0.point_1
    switch _goml_enum_tag_IntList(list123__0) {
    case 0:
        var inline0 string = "Empty list"
        var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
        _goml_runtime_core_string_println(inline1)
        return struct{}{}
    case 1:
        var x0 int32 = list123__0._node._p0
        var t0 string
        var inline9 string = __goml_builtin_int32_to_string(x0)
        t0 = inline9
        var inline7 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t0)
        _goml_runtime_core_string_println(inline7)
        var x1 int32 = point__0.x
        var x2 int32 = point__0.y
        var t1 string
        var inline6 string = __goml_builtin_int32_to_string(x1)
        t1 = inline6
        var t2 string = "Point: (" + t1
        var t3 string = t2 + ", "
        var t4 string
        var inline5 string = __goml_builtin_int32_to_string(x2)
        t4 = inline5
        var t5 string = t3 + t4
        var t6 string = t5 + ")"
        var inline3 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t6)
        _goml_runtime_core_string_println(inline3)
        return struct{}{}
    default:
        panic("non-exhaustive match")
    }
}

func main() {
    main0()
}
