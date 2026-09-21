package main

import (
    _goml_os "os"
)

type _goml_defer_state struct {
    actions []func() struct{}
}

func _goml_defer_take(stack *_goml_defer_state) func() struct{} {
    var index int = len(stack.actions) - 1
    var action func() struct{} = stack.actions[index]
    stack.actions[index] = nil
    stack.actions = stack.actions[0:index]
    return action
}

func _goml_defer_drain(stack *_goml_defer_state) {
    if len(stack.actions) != 0 {
        var action func() struct{} = _goml_defer_take(stack)
        defer _goml_defer_drain(stack)
        action()
    }
}

func _goml_runtime_core_defer_push(stack *_goml_defer_state, action func() struct{}) struct{} {
    stack.actions = append(stack.actions, action)
    return struct{}{}
}

func _goml_runtime_core_defer_pop(stack *_goml_defer_state) struct{} {
    var action func() struct{} = _goml_defer_take(stack)
    return action()
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

type _goml_vec_int struct {
    items []int
}

func vec_new__Vec_3int() *_goml_vec_int {
    return &_goml_vec_int{
        items: nil,
    }
}

func vec_push__Vec_3int(vec *_goml_vec_int, elem int) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_get__Vec_3int(vec *_goml_vec_int, index int) int {
    return vec.items[index]
}

func vec_len__Vec_3int(vec *_goml_vec_int) int {
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

type ref_int_x struct {
    value int
}

func ref__Ref_3int(value int) *ref_int_x {
    return &ref_int_x{
        value: value,
    }
}

func ref_get__Ref_3int(reference *ref_int_x) int {
    return reference.value
}

func ref_set__Ref_3int(reference *ref_int_x, value int) struct{} {
    reference.value = value
    return struct{}{}
}

type Tuple2_3int_3int struct {
    _0 int
    _1 int
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

type NumberSource struct {
    value int
}

type closure_env_labeled_cleanup_0 struct {}

type closure_env_labeled_cleanup_1 struct {}

type closure_env_increment_2 struct {
    captured_0 *ref_int_x
}

type Ordering uint8

type Option__isize struct {
    _p0 int
    _tag uint8
}

type dyn__Source_vtable struct {
    get func(any) int
}

type dyn__Source struct {
    data any
    vtable *dyn__Source_vtable
}

func dyn__Source__wrap__NumberSource__get(self any) int {
    return _goml_m_trait__impl_i_Source_i_NumberSource_i_get(self.(NumberSource))
}

func dyn__Source__vtable__NumberSource() *dyn__Source_vtable {
    return &dyn__Source_vtable{
        get: dyn__Source__wrap__NumberSource__get,
    }
}

func _goml_m_trait__impl_i_Source_i_NumberSource_i_get(self__0 NumberSource) int {
    var t0 int = self__0.value
    return t0
}

func labeled_cleanup() struct{} {
    var _goml_defer_stack _goml_defer_state
    defer _goml_defer_drain(&_goml_defer_stack)
    var t0 closure_env_labeled_cleanup_0 = closure_env_labeled_cleanup_0{}
    var t1 func() struct{} = func() struct{} {
        return _goml_m_inherent_i_closure__en_h0223a9befc2ebe8658869638c279e6f5_anup__0_i_apply(t0)
    }
    _goml_runtime_core_defer_push(&_goml_defer_stack, t1)
    var t2 closure_env_labeled_cleanup_1 = closure_env_labeled_cleanup_1{}
    var t3 func() struct{} = func() struct{} {
        return _goml_m_inherent_i_closure__en_hba5cfe2d49e8e7932e64c822e4593e98_anup__1_i_apply(t2)
    }
    _goml_runtime_core_defer_push(&_goml_defer_stack, t3)
    _goml_runtime_core_defer_pop(&_goml_defer_stack)
    _goml_runtime_core_defer_pop(&_goml_defer_stack)
    return struct{}{}
}

func main0() struct{} {
    var t0 NumberSource = NumberSource{
        value: 11,
    }
    var t1 dyn__Source = dyn__Source{
        data: t0,
        vtable: dyn__Source__vtable__NumberSource(),
    }
    var t2 int
    var inline26 int = t1.vtable.get(t1.data)
    t2 = inline26
    var inline24 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t2)
    _goml_runtime_core_string_println(inline24)
    var x0 int = 1
    var x1 int = 2
    var index__0 int = x0
    var compound_old0 int = index__0
    var t3 int = compound_old0 + x1
    index__0 = t3
    var inline22 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(index__0)
    _goml_runtime_core_string_println(inline22)
    var x2 int = 3
    var captured__0 *ref_int_x = ref__Ref_3int(x2)
    var t5 closure_env_increment_2 = closure_env_increment_2{
        captured_0: captured__0,
    }
    var increment__0 func() struct{} = func() struct{} {
        return _goml_m_inherent_i_closure__en_h79b4a32d6db2ef8c8920b1d54dab278c_ment__2_i_apply(t5)
    }
    increment__0()
    var t6 int = ref_get__Ref_3int(captured__0)
    var inline20 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t6)
    _goml_runtime_core_string_println(inline20)
    var x3 int = 4
    var count__0 int = x3
    var compound_old3 int = count__0
    var compound_value2 int = 1
    var t35 int = compound_old3 + compound_value2
    count__0 = t35
    var inline18 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(count__0)
    _goml_runtime_core_string_println(inline18)
    var values__0 *_goml_vec_int
    var inline17 *_goml_vec_int = vec_new__Vec_3int()
    values__0 = inline17
    var inline15 int = 6
    vec_push__Vec_3int(values__0, inline15)
    var for_limit0 int = vec_len__Vec_3int(values__0)
    var for_index0 int = 0
    Loop_loop0:
    for {
        var t31 bool = for_index0 < for_limit0
        if t31 {
            var for_item2 int = vec_get__Vec_3int(values__0, for_index0)
            var t32 int = for_index0 + 1
            for_index0 = t32
            var item__0 int = for_item2
            var compound_old2 int = item__0
            var compound_value1 int = 1
            var t33 int = compound_old2 + compound_value1
            item__0 = t33
            var inline13 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(item__0)
            _goml_runtime_core_string_println(inline13)
            continue
        } else {
            break Loop_loop0
        }
    }
    var tuple__0 Tuple2_3int_3int = Tuple2_3int_3int{
        _0: 8,
        _1: 9,
    }
    var place_root0 Tuple2_3int_3int = tuple__0
    var place0 int = place_root0._0
    var value0 int = 1
    var t7 int = place0 + value0
    var t8 int = place_root0._1
    var t9 Tuple2_3int_3int = Tuple2_3int_3int{
        _0: t7,
        _1: t8,
    }
    tuple__0 = t9
    var place_root1 Tuple2_3int_3int = tuple__0
    var place1 int = place_root1._1
    var value1 int = 1
    var t11 int = place_root1._0
    var t12 int = place1 + value1
    var t13 Tuple2_3int_3int = Tuple2_3int_3int{
        _0: t11,
        _1: t12,
    }
    tuple__0 = t13
    var t15 int = tuple__0._0
    var t16 int = tuple__0._1
    var t17 int = t15 + t16
    var inline11 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t17)
    _goml_runtime_core_string_println(inline11)
    var steps__0 int = 0
    Loop_loop1:
    for {
        var t28 bool = steps__0 < 3
        if t28 {
            var compound_old1 int = steps__0
            var compound_value0 int = 1
            var t29 int = compound_old1 + compound_value0
            steps__0 = t29
            continue
        } else {
            break Loop_loop1
        }
    }
    var inline9 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(steps__0)
    _goml_runtime_core_string_println(inline9)
    var seen__0 *ref_int_x
    var inline7 int = 0
    var inline8 *ref_int_x = ref__Ref_3int(inline7)
    seen__0 = inline8
    var for_index1 int = 0
    var for_limit1 int = 3
    Loop_loop2:
    for {
        var t20 bool = for_index1 < for_limit1
        if t20 {
            var for_item0 int = for_index1
            var t21 int = for_index1 + 1
            for_index1 = t21
            var for_index2 int = 0
            var for_limit2 int = 3
            var t22 bool = for_item0 == 1
            Loop_loop3:
            for {
                var t23 bool = for_index2 < for_limit2
                if t23 {
                    var for_item1 int = for_index2
                    var t24 int = for_index2 + 1
                    for_index2 = t24
                    var t25 int
                    var inline6 int = ref_get__Ref_3int(seen__0)
                    t25 = inline6
                    var t26 int = t25 + 1
                    ref_set__Ref_3int(seen__0, t26)
                    var jp1 bool
                    if t22 {
                        var t27 bool = for_item1 == 1
                        jp1 = t27
                    } else {
                        jp1 = false
                    }
                    if jp1 {
                        var t18 int
                        var inline4 int = ref_get__Ref_3int(seen__0)
                        t18 = inline4
                        var inline2 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t18)
                        _goml_runtime_core_string_println(inline2)
                        var jp0 int
                        jp0 = 42
                        var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(jp0)
                        _goml_runtime_core_string_println(inline0)
                        labeled_cleanup()
                        return struct{}{}
                    } else {
                        continue
                    }
                } else {
                    break Loop_loop3
                }
            }
            continue
        } else {
            break Loop_loop2
        }
    }
    var t18 int
    var inline4 int = ref_get__Ref_3int(seen__0)
    t18 = inline4
    var inline2 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t18)
    _goml_runtime_core_string_println(inline2)
    var jp0 int
    jp0 = 42
    var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(jp0)
    _goml_runtime_core_string_println(inline0)
    labeled_cleanup()
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func _goml_m_trait__impl_i_ToString_i_isize_i_to__string(self__0 int) string {
    var inline0 int64 = int64(int(self__0))
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

func _goml_m_inherent_i_closure__en_h0223a9befc2ebe8658869638c279e6f5_anup__0_i_apply(env0 closure_env_labeled_cleanup_0) struct{} {
    var inline0 string = "outer cleanup"
    var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
    _goml_runtime_core_string_println(inline1)
    return struct{}{}
}

func _goml_m_inherent_i_closure__en_hba5cfe2d49e8e7932e64c822e4593e98_anup__1_i_apply(env0 closure_env_labeled_cleanup_1) struct{} {
    var inline0 string = "inner cleanup"
    var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
    _goml_runtime_core_string_println(inline1)
    return struct{}{}
}

func _goml_m_inherent_i_closure__en_h79b4a32d6db2ef8c8920b1d54dab278c_ment__2_i_apply(env0 closure_env_increment_2) struct{} {
    var captured__0 *ref_int_x = env0.captured_0
    var compound_old0 int = ref_get__Ref_3int(captured__0)
    var compound_value0 int = 1
    var t0 int = compound_old0 + compound_value0
    ref_set__Ref_3int(captured__0, t0)
    return struct{}{}
}

func main() {
    main0()
}
