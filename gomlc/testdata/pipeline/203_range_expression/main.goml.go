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

type ref_bool_x struct {
    value bool
}

func ref__Ref_4bool(value bool) *ref_bool_x {
    return &ref_bool_x{
        value: value,
    }
}

func ref_get__Ref_4bool(reference *ref_bool_x) bool {
    return reference.value
}

func ref_set__Ref_4bool(reference *ref_bool_x, value bool) struct{} {
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

type FnIterator__isize struct {
    next_fn func() Option__isize
}

type closure_env_goml_builtin_range_inclusive_0 struct {
    finished_0 *ref_bool_x
    current_1 *ref_int_x
    end_2 int
}

type Ordering uint8

type Option__isize struct {
    _p0 int
    _tag uint8
}

func main0() struct{} {
    var for_index0 int = 1
    var for_limit0 int = 4
    Loop_loop0:
    for {
        var t17 bool = for_index0 < for_limit0
        if t17 {
            var for_item3 int = for_index0
            var t18 int = for_index0 + 1
            for_index0 = t18
            var inline38 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(for_item3)
            _goml_runtime_core_string_println(inline38)
            continue
        } else {
            break Loop_loop0
        }
    }
    var calls__0 *ref_int_x
    var inline36 int = 0
    var inline37 *ref_int_x = ref__Ref_3int(inline36)
    calls__0 = inline37
    var for_index1 int
    var inline32 int = 4
    var inline33 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(calls__0)
    var inline34 int = inline33 + 1
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__isize(calls__0, inline34)
    for_index1 = inline32
    var for_limit1 int
    var inline28 int = 6
    var inline29 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(calls__0)
    var inline30 int = inline29 + 1
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__isize(calls__0, inline30)
    for_limit1 = inline28
    var for_done0 bool = for_index1 > for_limit1
    Loop_loop1:
    for {
        var t12 bool = !for_done0
        if t12 {
            var for_item2 int = for_index1
            var t13 bool = for_index1 == for_limit1
            if t13 {
                for_done0 = true
            } else {
                var t15 int = for_index1 + 1
                for_index1 = t15
            }
            var inline26 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(for_item2)
            _goml_runtime_core_string_println(inline26)
            continue
        } else {
            break Loop_loop1
        }
    }
    var for_index2 int = 3
    var for_limit2 int = 1
    var for_done1 bool = for_index2 > for_limit2
    Loop_loop2:
    for {
        var t7 bool = !for_done1
        if t7 {
            var for_item1 int = for_index2
            var t8 bool = for_index2 == for_limit2
            if t8 {
                for_done1 = true
            } else {
                var t10 int = for_index2 + 1
                for_index2 = t10
            }
            var inline24 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(for_item1)
            _goml_runtime_core_string_println(inline24)
            continue
        } else {
            break Loop_loop2
        }
    }
    var maximum__0 int = 9223372036854775807
    var for_index3 int = maximum__0
    var for_done2 bool = for_index3 > maximum__0
    Loop_loop3:
    for {
        var t2 bool = !for_done2
        if t2 {
            var for_item0 int = for_index3
            var t3 bool = for_index3 == maximum__0
            if t3 {
                for_done2 = true
            } else {
                var t5 int = for_index3 + 1
                for_index3 = t5
            }
            var inline22 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(for_item0)
            _goml_runtime_core_string_println(inline22)
            continue
        } else {
            break Loop_loop3
        }
    }
    var iterator__0 FnIterator__isize
    var inline15 int = 8
    var inline16 int = 8
    var inline17 *ref_int_x = ref__Ref_3int(inline15)
    var inline18 *ref_bool_x = ref__Ref_4bool(false)
    var inline19 closure_env_goml_builtin_range_inclusive_0 = closure_env_goml_builtin_range_inclusive_0{
        finished_0: inline18,
        current_1: inline17,
        end_2: inline16,
    }
    var inline20 func() Option__isize = func() Option__isize {
        return _goml_m_inherent_i_closure__en_hb902f75cf29154a7d4df1174edbd9988_sive__0_i_apply(inline19)
    }
    var inline21 FnIterator__isize = _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__isize(inline20)
    iterator__0 = inline21
    var mtmp0 Option__isize
    var inline13 func() Option__isize = iterator__0.next_fn
    var inline14 Option__isize = inline13()
    mtmp0 = inline14
    switch mtmp0._tag {
    case 0:
        var inline8 string = "missing"
        var inline9 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline8)
        _goml_runtime_core_string_println(inline9)
    case 1:
        var x0 int = mtmp0._p0
        var inline11 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(x0)
        _goml_runtime_core_string_println(inline11)
    default:
        panic("non-exhaustive match")
    }
    var t0 int
    var inline7 int = ref_get__Ref_3int(calls__0)
    t0 = inline7
    var inline5 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t0)
    _goml_runtime_core_string_println(inline5)
    var t1 int32
    var inline2 int32 = 10
    var inline3 int32 = 20
    var inline4 int32 = inline2 + inline3
    t1 = inline4
    var inline0 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(t1)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(self__0 *ref_int_x) int {
    var t0 int = ref_get__Ref_3int(self__0)
    return t0
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__isize(self__0 *ref_int_x, value__0 int) struct{} {
    ref_set__Ref_3int(self__0, value__0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_isize_i_to__string(self__0 int) string {
    var inline0 int64 = int64(int(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__isize(next_fn__0 func() Option__isize) FnIterator__isize {
    var t0 FnIterator__isize = FnIterator__isize{
        next_fn: next_fn__0,
    }
    return t0
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
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

func _goml_m_inherent_i_closure__en_hb902f75cf29154a7d4df1174edbd9988_sive__0_i_apply(env0 closure_env_goml_builtin_range_inclusive_0) Option__isize {
    var finished__0 *ref_bool_x = env0.finished_0
    var current__0 *ref_int_x = env0.current_1
    var end__0 int = env0.end_2
    var t0 bool = ref_get__Ref_4bool(finished__0)
    var jp0 bool
    if t0 {
        jp0 = true
    } else {
        var t4 int = ref_get__Ref_3int(current__0)
        var t5 bool = t4 > end__0
        jp0 = t5
    }
    if jp0 {
        return Option__isize{
            _tag: 0,
        }
    } else {
        var value__0 int = ref_get__Ref_3int(current__0)
        var t1 bool = value__0 == end__0
        if t1 {
            ref_set__Ref_4bool(finished__0, true)
        } else {
            var t3 int = value__0 + 1
            ref_set__Ref_3int(current__0, t3)
        }
        var t2 Option__isize = Option__isize{
            _p0: value__0,
            _tag: 1,
        }
        return t2
    }
}

func main() {
    main0()
}
