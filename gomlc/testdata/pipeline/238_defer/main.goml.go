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

type ref_string_x struct {
    value string
}

func ref__Ref_6string(value string) *ref_string_x {
    return &ref_string_x{
        value: value,
    }
}

func ref_get__Ref_6string(reference *ref_string_x) string {
    return reference.value
}

func ref_set__Ref_6string(reference *ref_string_x, value string) struct{} {
    reference.value = value
    return struct{}{}
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

type closure_env_early_return_0 struct {}

type closure_env_early_return_1 struct {}

type closure_env_maybe_2 struct {}

type closure_env_loop_cleanup_3 struct {
    current_0 int
}

type closure_env_observed_at_exit_4 struct {
    value_0 *ref_string_x
}

type closure_env_pattern_cleanup_5 struct {}

type closure_env_closure_cleanup_6 struct {}

type closure_env_run_7 struct {}

type closure_env_run_8 struct {}

type closure_env_main_9 struct {}

type closure_env_main_10 struct {}

type closure_env_main_11 struct {}

type Ordering uint8

type Option__isize struct {
    _p0 int
    _tag uint8
}

func early_return() int {
    var _goml_defer_stack _goml_defer_state
    defer _goml_defer_drain(&_goml_defer_stack)
    var t0 closure_env_early_return_0 = closure_env_early_return_0{}
    var t1 func() struct{} = func() struct{} {
        return _goml_m_inherent_i_closure__en_hfb1ffcf3aaa950d8d3ed3ad9de792fc7_turn__0_i_apply(t0)
    }
    _goml_runtime_core_defer_push(&_goml_defer_stack, t1)
    var t2 closure_env_early_return_1 = closure_env_early_return_1{}
    var t3 func() struct{} = func() struct{} {
        return _goml_m_inherent_i_closure__en_h0b18f602d910aaca7e28529af76b9600_turn__1_i_apply(t2)
    }
    _goml_runtime_core_defer_push(&_goml_defer_stack, t3)
    var defer_return0 int = 7
    _goml_runtime_core_defer_pop(&_goml_defer_stack)
    _goml_runtime_core_defer_pop(&_goml_defer_stack)
    return defer_return0
}

func loop_cleanup() struct{} {
    var _goml_defer_stack _goml_defer_state
    defer _goml_defer_drain(&_goml_defer_stack)
    var index__0 *ref_int_x
    var inline3 int = 0
    var inline4 *ref_int_x = ref__Ref_3int(inline3)
    index__0 = inline4
    Loop_loop0:
    for {
        var t0 int
        var inline2 int = ref_get__Ref_3int(index__0)
        t0 = inline2
        var t1 bool = t0 < 3
        if t1 {
            var current__0 int
            var inline1 int = ref_get__Ref_3int(index__0)
            current__0 = inline1
            var t2 closure_env_loop_cleanup_3 = closure_env_loop_cleanup_3{
                current_0: current__0,
            }
            var t3 func() struct{} = func() struct{} {
                return _goml_m_inherent_i_closure__en_h69ae75ee1025d2960156fd964aebbc5b_anup__3_i_apply(t2)
            }
            _goml_runtime_core_defer_push(&_goml_defer_stack, t3)
            var t4 int = current__0 + 1
            ref_set__Ref_3int(index__0, t4)
            var t5 bool = current__0 == 0
            if t5 {
                _goml_runtime_core_defer_pop(&_goml_defer_stack)
                continue
            } else {
                var t6 bool = current__0 == 1
                if t6 {
                    _goml_runtime_core_defer_pop(&_goml_defer_stack)
                    break Loop_loop0
                } else {
                    _goml_runtime_core_defer_pop(&_goml_defer_stack)
                    continue
                }
            }
        } else {
            break Loop_loop0
        }
    }
    return struct{}{}
}

func pattern_cleanup(value__0 Option__isize) int {
    var _goml_defer_stack _goml_defer_state
    defer _goml_defer_drain(&_goml_defer_stack)
    var t0 closure_env_pattern_cleanup_5 = closure_env_pattern_cleanup_5{}
    var t1 func() struct{} = func() struct{} {
        return _goml_m_inherent_i_closure__en_h146808f0ec2411191d8d80c65887fd74_anup__5_i_apply(t0)
    }
    _goml_runtime_core_defer_push(&_goml_defer_stack, t1)
    switch value__0._tag {
    case 1:
        var x0 int = value__0._p0
        var x1 int = 2
        var defer_tast_result0 int = x0 + x1
        _goml_runtime_core_defer_pop(&_goml_defer_stack)
        return defer_tast_result0
    default:
        var defer_return0 int = 0
        _goml_runtime_core_defer_pop(&_goml_defer_stack)
        return defer_return0
    }
}

func main0() struct{} {
    var _goml_defer_stack _goml_defer_state
    defer _goml_defer_drain(&_goml_defer_stack)
    var t0 closure_env_main_9 = closure_env_main_9{}
    var t1 func() struct{} = func() struct{} {
        return _goml_m_inherent_i_closure__env__main__9_i_closure__env__main__9_i_apply(t0)
    }
    _goml_runtime_core_defer_push(&_goml_defer_stack, t1)
    var t2 closure_env_main_10 = closure_env_main_10{}
    var t3 func() struct{} = func() struct{} {
        return _goml_m_inherent_i_closure__env__main__10_i_closure__env__main__10_i_apply(t2)
    }
    _goml_runtime_core_defer_push(&_goml_defer_stack, t3)
    var t4 closure_env_main_11 = closure_env_main_11{}
    var t5 func() struct{} = func() struct{} {
        return _goml_m_inherent_i_closure__env__main__11_i_closure__env__main__11_i_apply(t4)
    }
    _goml_runtime_core_defer_push(&_goml_defer_stack, t5)
    println__T_string("body")
    _goml_runtime_core_defer_pop(&_goml_defer_stack)
    var t6 int = early_return()
    var t7 string
    var inline27 string = __goml_builtin_int_to_string(t6)
    t7 = inline27
    var inline25 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t7)
    _goml_runtime_core_string_println(inline25)
    var inline20 closure_env_maybe_2 = closure_env_maybe_2{}
    var inline21 func() struct{} = func() struct{} {
        return _goml_m_inherent_i_closure__env__maybe__2_i_closure__env__maybe__2_i_apply(inline20)
    }
    _goml_runtime_core_defer_push(&_goml_defer_stack, inline21)
    _goml_runtime_core_defer_pop(&_goml_defer_stack)
    loop_cleanup()
    var inline14 *ref_string_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__string("before")
    var inline15 closure_env_observed_at_exit_4 = closure_env_observed_at_exit_4{
        value_0: inline14,
    }
    var inline16 func() struct{} = func() struct{} {
        return _goml_m_inherent_i_closure__en_h3cbaa726092214ab80f9b4b4e6e91f7d_exit__4_i_apply(inline15)
    }
    _goml_runtime_core_defer_push(&_goml_defer_stack, inline16)
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__string(inline14, "after")
    _goml_runtime_core_defer_pop(&_goml_defer_stack)
    var t8 Option__isize = Option__isize{
        _p0: 3,
        _tag: 1,
    }
    var t9 int = pattern_cleanup(t8)
    var t10 string
    var inline13 string = __goml_builtin_int_to_string(t9)
    t10 = inline13
    var inline11 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t10)
    _goml_runtime_core_string_println(inline11)
    var t11 int = pattern_cleanup(Option__isize{
        _tag: 0,
    })
    var t12 string
    var inline10 string = __goml_builtin_int_to_string(t11)
    t12 = inline10
    var inline8 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t12)
    _goml_runtime_core_string_println(inline8)
    var inline0 closure_env_closure_cleanup_6 = closure_env_closure_cleanup_6{}
    var inline1 func() struct{} = func() struct{} {
        return _goml_m_inherent_i_closure__en_ha4ac8164b2e8d8c6c7666ec986832880_anup__6_i_apply(inline0)
    }
    _goml_runtime_core_defer_push(&_goml_defer_stack, inline1)
    var inline3 closure_env_run_8 = closure_env_run_8{}
    var inline4 func() struct{} = func() struct{} {
        return _goml_m_inherent_i_closure__env__run__8_i_closure__env__run__8_i_apply(inline3)
    }
    inline4()
    println__T_string("closure:after")
    _goml_runtime_core_defer_pop(&_goml_defer_stack)
    _goml_runtime_core_defer_pop(&_goml_defer_stack)
    _goml_runtime_core_defer_pop(&_goml_defer_stack)
    return struct{}{}
}

func println__T_string(value__0 string) struct{} {
    var t0 string
    t0 = value__0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__string(value__0 string) *ref_string_x {
    var t0 *ref_string_x = ref__Ref_6string(value__0)
    return t0
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__string(self__0 *ref_string_x, value__0 string) struct{} {
    ref_set__Ref_6string(self__0, value__0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
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

func _goml_m_inherent_i_closure__en_hfb1ffcf3aaa950d8d3ed3ad9de792fc7_turn__0_i_apply(env0 closure_env_early_return_0) struct{} {
    var inline0 string = "return:outer"
    var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
    _goml_runtime_core_string_println(inline1)
    return struct{}{}
}

func _goml_m_inherent_i_closure__en_h0b18f602d910aaca7e28529af76b9600_turn__1_i_apply(env0 closure_env_early_return_1) struct{} {
    var inline0 string = "return:inner"
    var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
    _goml_runtime_core_string_println(inline1)
    return struct{}{}
}

func _goml_m_inherent_i_closure__env__maybe__2_i_closure__env__maybe__2_i_apply(env0 closure_env_maybe_2) struct{} {
    var inline0 string = "try:cleanup"
    var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
    _goml_runtime_core_string_println(inline1)
    return struct{}{}
}

func _goml_m_inherent_i_closure__en_h69ae75ee1025d2960156fd964aebbc5b_anup__3_i_apply(env0 closure_env_loop_cleanup_3) struct{} {
    var current__0 int = env0.current_0
    var t0 string
    var inline2 string = __goml_builtin_int_to_string(current__0)
    t0 = inline2
    var t1 string = "loop:" + t0
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t1)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_inherent_i_closure__en_h3cbaa726092214ab80f9b4b4e6e91f7d_exit__4_i_apply(env0 closure_env_observed_at_exit_4) struct{} {
    var value__0 *ref_string_x = env0.value_0
    var t0 string
    var inline2 string = ref_get__Ref_6string(value__0)
    t0 = inline2
    var t1 string = "observed:" + t0
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t1)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_inherent_i_closure__en_h146808f0ec2411191d8d80c65887fd74_anup__5_i_apply(env0 closure_env_pattern_cleanup_5) struct{} {
    var inline0 string = "pattern:cleanup"
    var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
    _goml_runtime_core_string_println(inline1)
    return struct{}{}
}

func _goml_m_inherent_i_closure__en_ha4ac8164b2e8d8c6c7666ec986832880_anup__6_i_apply(env0 closure_env_closure_cleanup_6) struct{} {
    var inline0 string = "closure:outer"
    var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
    _goml_runtime_core_string_println(inline1)
    return struct{}{}
}

func _goml_m_inherent_i_closure__env__run__7_i_closure__env__run__7_i_apply(env0 closure_env_run_7) struct{} {
    var inline0 string = "closure:inner"
    var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
    _goml_runtime_core_string_println(inline1)
    return struct{}{}
}

func _goml_m_inherent_i_closure__env__run__8_i_closure__env__run__8_i_apply(env0 closure_env_run_8) struct{} {
    var _goml_defer_stack _goml_defer_state
    defer _goml_defer_drain(&_goml_defer_stack)
    var t0 closure_env_run_7 = closure_env_run_7{}
    var t1 func() struct{} = func() struct{} {
        return _goml_m_inherent_i_closure__env__run__7_i_closure__env__run__7_i_apply(t0)
    }
    _goml_runtime_core_defer_push(&_goml_defer_stack, t1)
    var inline0 string = "closure:body"
    var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
    _goml_runtime_core_string_println(inline1)
    _goml_runtime_core_defer_pop(&_goml_defer_stack)
    return struct{}{}
}

func _goml_m_inherent_i_closure__env__main__9_i_closure__env__main__9_i_apply(env0 closure_env_main_9) struct{} {
    var inline0 string = "main:first"
    var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
    _goml_runtime_core_string_println(inline1)
    return struct{}{}
}

func _goml_m_inherent_i_closure__env__main__10_i_closure__env__main__10_i_apply(env0 closure_env_main_10) struct{} {
    var inline0 string = "main:second"
    var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
    _goml_runtime_core_string_println(inline1)
    return struct{}{}
}

func _goml_m_inherent_i_closure__env__main__11_i_closure__env__main__11_i_apply(env0 closure_env_main_11) struct{} {
    var inline0 string = "block"
    var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
    _goml_runtime_core_string_println(inline1)
    return struct{}{}
}

func main() {
    main0()
}
