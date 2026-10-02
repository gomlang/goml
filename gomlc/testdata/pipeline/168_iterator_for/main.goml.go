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

func vec_get__Vec_5int32(vec *_goml_vec_int32, index int) int32 {
    return vec.items[index]
}

func vec_len__Vec_5int32(vec *_goml_vec_int32) int {
    return int(len(vec.items))
}

type _goml_vec_Tuple2_5int32_6string struct {
    items []Tuple2_5int32_6string
}

func vec_new__Vec_21Tuple2_5int32_6string() *_goml_vec_Tuple2_5int32_6string {
    return &_goml_vec_Tuple2_5int32_6string{
        items: nil,
    }
}

func vec_push__Vec_21Tuple2_5int32_6string(vec *_goml_vec_Tuple2_5int32_6string, elem Tuple2_5int32_6string) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_get__Vec_21Tuple2_5int32_6string(vec *_goml_vec_Tuple2_5int32_6string, index int) Tuple2_5int32_6string {
    return vec.items[index]
}

func vec_len__Vec_21Tuple2_5int32_6string(vec *_goml_vec_Tuple2_5int32_6string) int {
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

type ref_int32_x struct {
    value int32
}

func ref__Ref_5int32(value int32) *ref_int32_x {
    return &ref_int32_x{
        value: value,
    }
}

func ref_get__Ref_5int32(reference *ref_int32_x) int32 {
    return reference.value
}

func ref_set__Ref_5int32(reference *ref_int32_x, value int32) struct{} {
    reference.value = value
    return struct{}{}
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

type Tuple2_5int32_6string struct {
    _0 int32
    _1 string
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

type FnIterator__i32 struct {
    next_fn func() Option__i32
}

type FnIterator__isize struct {
    next_fn func() Option__isize
}

type closure_env_countdown_0 struct {
    current_0 *ref_int32_x
}

type closure_env_goml_builtin_range_1 struct {
    current_0 *ref_int_x
    end_1 int
}

type Ordering uint8

type Option__i32 uint64

type Option__isize struct {
    _p0 int
    _tag uint8
}

func counted_range(calls__0 *ref_int32_x) FnIterator__isize {
    var t0 int32
    var inline4 int32 = ref_get__Ref_5int32(calls__0)
    t0 = inline4
    var t1 int32 = t0 + 1
    ref_set__Ref_5int32(calls__0, t1)
    var inline0 int = 1
    var inline1 int = 5
    var inline2 FnIterator__isize = __goml_builtin_range(inline0, inline1)
    return inline2
}

func first_even(values__0 FnIterator__isize) int {
    var for_iter0 FnIterator__isize
    for_iter0 = values__0
    Loop_loop0:
    for {
        var for_next0 Option__isize
        var inline0 func() Option__isize = for_iter0.next_fn
        var inline1 Option__isize = inline0()
        for_next0 = inline1
        switch for_next0._tag {
        case 0:
            break Loop_loop0
        case 1:
            var x0 int = for_next0._p0
            var t0 int = x0 / 2
            var t1 int = t0 * 2
            var t2 bool = t1 == x0
            if t2 {
                return x0
            } else {
                continue
            }
        default:
            panic("non-exhaustive match")
        }
    }
    return -1
}

func main0() struct{} {
    var values__0 *_goml_vec_int32 = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_new____T__i32()
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_push____T__i32(values__0, 10)
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_push____T__i32(values__0, 20)
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_push____T__i32(values__0, 30)
    var sum__0 *ref_int32_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__i32(0)
    var for_limit0 int = vec_len__Vec_5int32(values__0)
    var for_index0 int = 0
    Loop_loop0:
    for {
        var t21 bool = for_index0 < for_limit0
        if t21 {
            var for_item2 int32 = vec_get__Vec_5int32(values__0, for_index0)
            var t22 int = for_index0 + 1
            for_index0 = t22
            var t23 bool = for_item2 == 20
            if t23 {
                continue
            } else {
                var t24 int32
                var inline33 int32 = ref_get__Ref_5int32(sum__0)
                t24 = inline33
                var t25 int32 = t24 + for_item2
                ref_set__Ref_5int32(sum__0, t25)
                continue
            }
        } else {
            break Loop_loop0
        }
    }
    var t0 int32 = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__i32(sum__0)
    println__T_i32(t0)
    var pairs__0 *_goml_vec_Tuple2_5int32_6string = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_new____T___o_i32_c_string_q_()
    var t1 Tuple2_5int32_6string = Tuple2_5int32_6string{
        _0: 1,
        _1: "a",
    }
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_push____T___o_i32_c_string_q_(pairs__0, t1)
    var t2 Tuple2_5int32_6string = Tuple2_5int32_6string{
        _0: 2,
        _1: "b",
    }
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_push____T___o_i32_c_string_q_(pairs__0, t2)
    var for_limit1 int = vec_len__Vec_21Tuple2_5int32_6string(pairs__0)
    var for_index1 int = 0
    Loop_loop1:
    for {
        var t17 bool = for_index1 < for_limit1
        if t17 {
            var for_item1 Tuple2_5int32_6string = vec_get__Vec_21Tuple2_5int32_6string(pairs__0, for_index1)
            var t18 int = for_index1 + 1
            for_index1 = t18
            var x2 int32 = for_item1._0
            var x3 string = for_item1._1
            var t19 string
            var inline31 string = __goml_builtin_int32_to_string(x2)
            t19 = inline31
            var t20 string = t19 + x3
            var inline29 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t20)
            _goml_runtime_core_string_println(inline29)
            continue
        } else {
            break Loop_loop1
        }
    }
    var calls__0 *ref_int32_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__i32(0)
    var range_sum__0 *ref_int_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__isize(0)
    var t3 FnIterator__isize = counted_range(calls__0)
    var for_iter0 FnIterator__isize = _goml_m_trait__impl_i_IntoIterator_i_FnIterator____isize_i_into__iter(t3)
    Loop_loop2:
    for {
        var for_next2 Option__isize
        var inline27 func() Option__isize = for_iter0.next_fn
        var inline28 Option__isize = inline27()
        for_next2 = inline28
        switch for_next2._tag {
        case 0:
            break Loop_loop2
        case 1:
            var x1 int = for_next2._p0
            var t15 int
            var inline26 int = ref_get__Ref_3int(range_sum__0)
            t15 = inline26
            var t16 int = t15 + x1
            ref_set__Ref_3int(range_sum__0, t16)
            continue
        default:
            panic("non-exhaustive match")
        }
    }
    var t4 int32 = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__i32(calls__0)
    println__T_i32(t4)
    var t5 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(range_sum__0)
    println__T_isize(t5)
    var slice_sum__0 *ref_int32_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__i32(0)
    var for_source0 []int32 = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_slice____T__i32(values__0, 1, 3)
    var for_limit2 int = len(for_source0)
    var for_index2 int = 0
    Loop_loop3:
    for {
        var t11 bool = for_index2 < for_limit2
        if t11 {
            var for_item0 int32 = for_source0[for_index2]
            var t12 int = for_index2 + 1
            for_index2 = t12
            var t13 int32
            var inline24 int32 = ref_get__Ref_5int32(slice_sum__0)
            t13 = inline24
            var t14 int32 = t13 + for_item0
            ref_set__Ref_5int32(slice_sum__0, t14)
            continue
        } else {
            break Loop_loop3
        }
    }
    var t6 int32
    var inline22 int32 = ref_get__Ref_5int32(slice_sum__0)
    t6 = inline22
    println__T_i32(t6)
    var t7 FnIterator__i32
    var inline17 int32 = 4
    var inline18 *ref_int32_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__i32(inline17)
    var inline19 closure_env_countdown_0 = closure_env_countdown_0{
        current_0: inline18,
    }
    var inline20 func() Option__i32 = func() Option__i32 {
        return _goml_m_inherent_i_closure__en_hc257e8a36c560f54cbc91ed82b2d188c_down__0_i_apply(inline19)
    }
    var inline21 FnIterator__i32 = _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__i32(inline20)
    t7 = inline21
    var for_iter1 FnIterator__i32
    for_iter1 = t7
    Loop_loop4:
    for {
        var for_next1 Option__i32
        var inline15 func() Option__i32 = for_iter1.next_fn
        var inline16 Option__i32 = inline15()
        for_next1 = inline16
        switch for_next1 != Option__i32(0) {
        case false:
            break Loop_loop4
        case true:
            var x0 int32 = int32(int64(uint64(for_next1) - 1) - 2147483648)
            var t10 bool = x0 == 2
            if t10 {
                break Loop_loop4
            } else {
                var inline13 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(x0)
                _goml_runtime_core_string_println(inline13)
                continue
            }
        default:
            panic("non-exhaustive match")
        }
    }
    var empty__0 FnIterator__isize
    var inline10 int = 0
    var inline11 int = 0
    var inline12 FnIterator__isize = __goml_builtin_range(inline10, inline11)
    empty__0 = inline12
    var for_iter2 FnIterator__isize
    for_iter2 = empty__0
    Loop_loop5:
    for {
        var for_next0 Option__isize
        var inline8 func() Option__isize = for_iter2.next_fn
        var inline9 Option__isize = inline8()
        for_next0 = inline9
        switch for_next0._tag {
        case 0:
            break Loop_loop5
        case 1:
            continue
        default:
            panic("non-exhaustive match")
        }
    }
    var t8 FnIterator__isize
    var inline5 int = 3
    var inline6 int = 8
    var inline7 FnIterator__isize = __goml_builtin_range(inline5, inline6)
    t8 = inline7
    var t9 int = first_even(t8)
    var inline3 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t9)
    _goml_runtime_core_string_println(inline3)
    var inline0 string = "done"
    var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
    _goml_runtime_core_string_println(inline1)
    return struct{}{}
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__i32(value__0 int32) *ref_int32_x {
    var t0 *ref_int32_x = ref__Ref_5int32(value__0)
    return t0
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__i32(self__0 *ref_int32_x) int32 {
    var t0 int32 = ref_get__Ref_5int32(self__0)
    return t0
}

func _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__i32(next_fn__0 func() Option__i32) FnIterator__i32 {
    var t0 FnIterator__i32 = FnIterator__i32{
        next_fn: next_fn__0,
    }
    return t0
}

func _goml_m_trait__impl_i_IntoIterator_i_FnIterator____isize_i_into__iter(self__0 FnIterator__isize) FnIterator__isize {
    return self__0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_new____T__i32() *_goml_vec_int32 {
    var t0 *_goml_vec_int32 = vec_new__Vec_5int32()
    return t0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_push____T__i32(self__0 *_goml_vec_int32, elem__0 int32) struct{} {
    vec_push__Vec_5int32(self__0, elem__0)
    return struct{}{}
}

func println__T_i32(value__0 int32) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_int32_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_new____T___o_i32_c_string_q_() *_goml_vec_Tuple2_5int32_6string {
    var t0 *_goml_vec_Tuple2_5int32_6string = vec_new__Vec_21Tuple2_5int32_6string()
    return t0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_push____T___o_i32_c_string_q_(self__0 *_goml_vec_Tuple2_5int32_6string, elem__0 Tuple2_5int32_6string) struct{} {
    vec_push__Vec_21Tuple2_5int32_6string(self__0, elem__0)
    return struct{}{}
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__isize(value__0 int) *ref_int_x {
    var t0 *ref_int_x = ref__Ref_3int(value__0)
    return t0
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(self__0 *ref_int_x) int {
    var t0 int = ref_get__Ref_3int(self__0)
    return t0
}

func println__T_isize(value__0 int) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_int_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_slice____T__i32(self__0 *_goml_vec_int32, start__0 int, end__0 int) []int32 {
    var t0 []int32 = self__0.items[start__0:end__0]
    return t0
}

func __goml_builtin_range(start__0 int, end__0 int) FnIterator__isize {
    var current__0 *ref_int_x = ref__Ref_3int(start__0)
    var t0 closure_env_goml_builtin_range_1 = closure_env_goml_builtin_range_1{
        current_0: current__0,
        end_1: end__0,
    }
    var t1 func() Option__isize = func() Option__isize {
        return _goml_m_inherent_i_closure__en_h07c29ff1f344b08e028033881af7c2d9_ange__1_i_apply(t0)
    }
    var inline0 FnIterator__isize = FnIterator__isize{
        next_fn: t1,
    }
    return inline0
}

func _goml_m_trait__impl_i_ToString_i_i32_i_to__string(self__0 int32) string {
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

func _goml_m_inherent_i_closure__en_hc257e8a36c560f54cbc91ed82b2d188c_down__0_i_apply(env0 closure_env_countdown_0) Option__i32 {
    var current__0 *ref_int32_x = env0.current_0
    var value__0 int32
    var inline1 int32 = ref_get__Ref_5int32(current__0)
    value__0 = inline1
    var t0 bool = value__0 > 0
    if t0 {
        var t1 int32 = value__0 - 1
        ref_set__Ref_5int32(current__0, t1)
        var t2 Option__i32 = Option__i32(uint64(int64(value__0) + 2147483648) + 1)
        return t2
    } else {
        return Option__i32(0)
    }
}

func _goml_m_inherent_i_closure__en_h07c29ff1f344b08e028033881af7c2d9_ange__1_i_apply(env0 closure_env_goml_builtin_range_1) Option__isize {
    var current__0 *ref_int_x = env0.current_0
    var end__0 int = env0.end_1
    var value__0 int = ref_get__Ref_3int(current__0)
    var t0 bool = value__0 < end__0
    if t0 {
        var t1 int = value__0 + 1
        ref_set__Ref_3int(current__0, t1)
        var t2 Option__isize = Option__isize{
            _p0: value__0,
            _tag: 1,
        }
        return t2
    } else {
        return Option__isize{
            _tag: 0,
        }
    }
}

func main() {
    main0()
}
