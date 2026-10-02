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

type _goml_vec_string struct {
    items []string
}

func vec_new__Vec_6string() *_goml_vec_string {
    return &_goml_vec_string{
        items: nil,
    }
}

func vec_push__Vec_6string(vec *_goml_vec_string, elem string) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_get__Vec_6string(vec *_goml_vec_string, index int) string {
    return vec.items[index]
}

func vec_len__Vec_6string(vec *_goml_vec_string) int {
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

type Token struct {}

type Any struct {}

type Counter struct {
    current *ref_int32_x
    end int32
}

type FnIterator__i32 struct {
    next_fn func() Option__i32
}

type FnIterator__isize struct {
    next_fn func() Option__isize
}

type FnIterator__string struct {
    next_fn func() Option__string
}

type closure_env_main_0 struct {}

type closure_env_main_1 struct {}

type closure_env_main_2 struct {}

type closure_env_main_3 struct {}

type closure_env_std_iter_map_A_i32_B_i32_I_Counter_4 struct {
    map_fn_0 func(int32) int32
}

type closure_env_std_iter_map_A_i32_B_i32_I_Counter_5 struct {
    iterator_0 Counter
    map_fn_1 func(int32) int32
}

type closure_env_std_iter_filter_I_FnIterator_i32_T_i32_6 struct {
    iterator_0 FnIterator__i32
    predicate_1 func(int32) bool
}

type closure_env_std_iter_take_I_FnIterator_i32_7 struct {
    remaining_0 *ref_int_x
    iterator_1 FnIterator__i32
}

type closure_env_std_iter_map_A_isize_B_string_I_FnIterator_isize_8 struct {
    map_fn_0 func(int) string
}

type closure_env_std_iter_map_A_isize_B_string_I_FnIterator_isize_9 struct {
    iterator_0 FnIterator__isize
    map_fn_1 func(int) string
}

type closure_env_goml_builtin_range_10 struct {
    current_0 *ref_int_x
    end_1 int
}

type Ordering uint8

type Option__i32 uint64

type Option__isize struct {
    _p0 int
    _tag uint8
}

type Option__string struct {
    _p0 string
    _tag uint8
}

func _goml_m_trait__impl_i_Convert_i__l_i32_r__x40_Token_i_convert(self__0 Token) int32 {
    return 7
}

func _goml_m_trait__impl_i_Convert_i__l_string_r__x40_Token_i_convert(self__0 Token) string {
    return "seven"
}

func main0() struct{} {
    var t0 Token = Token{}
    var t1 int32 = _goml_m_trait__impl_i_Convert_i__l_i32_r__x40_Token_i_convert(t0)
    println__T_i32(t1)
    var t2 Token = Token{}
    var t3 string = _goml_m_trait__impl_i_Convert_i__l_string_r__x40_Token_i_convert(t2)
    println__T_string(t3)
    var t4 Token = Token{}
    var converted__0 int32 = convert_to__T_i32__V_Token(t4)
    println__T_i32(converted__0)
    var t5 string
    t5 = "marked"
    println__T_string(t5)
    var t6 string
    t6 = "marked"
    println__T_string(t6)
    var t7 string
    t7 = "marked"
    var inline34 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t7)
    _goml_runtime_core_string_println(inline34)
    var t8 Counter
    var inline30 int32 = 0
    var inline31 int32 = 8
    var inline32 *ref_int32_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__i32(inline30)
    var inline33 Counter = Counter{
        current: inline32,
        end: inline31,
    }
    t8 = inline33
    var t9 closure_env_main_0 = closure_env_main_0{}
    var t10 func(int32) int32 = func(p0 int32) int32 {
        return _goml_m_inherent_i_closure__env__main__0_i_closure__env__main__0_i_apply(t9, p0)
    }
    var mapped__0 FnIterator__i32
    var inline27 closure_env_std_iter_map_A_i32_B_i32_I_Counter_5 = closure_env_std_iter_map_A_i32_B_i32_I_Counter_5{
        iterator_0: t8,
        map_fn_1: t10,
    }
    var inline28 func() Option__i32 = func() Option__i32 {
        return _goml_m_inherent_i_closure__en_h39f5b01177a579a305108bc4e47ab576_nter__5_i_apply(inline27)
    }
    var inline29 FnIterator__i32 = _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__i32(inline28)
    mapped__0 = inline29
    var t11 closure_env_main_1 = closure_env_main_1{}
    var t12 func(int32) bool = func(p0 int32) bool {
        return _goml_m_inherent_i_closure__env__main__1_i_closure__env__main__1_i_apply(t11, p0)
    }
    var filtered__0 FnIterator__i32
    var inline24 closure_env_std_iter_filter_I_FnIterator_i32_T_i32_6 = closure_env_std_iter_filter_I_FnIterator_i32_T_i32_6{
        iterator_0: mapped__0,
        predicate_1: t12,
    }
    var inline25 func() Option__i32 = func() Option__i32 {
        return _goml_m_inherent_i_closure__en_h920672151d9ffb42cb9c90456440c3f4__i32__6_i_apply(inline24)
    }
    var inline26 FnIterator__i32 = _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__i32(inline25)
    filtered__0 = inline26
    var limited__0 FnIterator__i32
    var inline17 int = 3
    var inline18 bool = inline17 > 0
    var inline19 int
    if inline18 {
        inline19 = inline17
    } else {
        inline19 = 0
    }
    var inline20 *ref_int_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__isize(inline19)
    var inline21 closure_env_std_iter_take_I_FnIterator_i32_7 = closure_env_std_iter_take_I_FnIterator_i32_7{
        remaining_0: inline20,
        iterator_1: filtered__0,
    }
    var inline22 func() Option__i32 = func() Option__i32 {
        return _goml_m_inherent_i_closure__en_h62d5c49070b32f607a729bd540373bb1__i32__7_i_apply(inline21)
    }
    var inline23 FnIterator__i32 = _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__i32(inline22)
    limited__0 = inline23
    var for_iter0 FnIterator__i32
    for_iter0 = limited__0
    Loop_loop0:
    for {
        var for_next0 Option__i32
        var inline15 func() Option__i32 = for_iter0.next_fn
        var inline16 Option__i32 = inline15()
        for_next0 = inline16
        switch for_next0 != Option__i32(0) {
        case false:
            break Loop_loop0
        case true:
            var x0 int32 = int32(int64(uint64(for_next0) - 1) - 2147483648)
            var inline13 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(x0)
            _goml_runtime_core_string_println(inline13)
            continue
        default:
            panic("non-exhaustive match")
        }
    }
    var t13 FnIterator__isize
    var inline10 int = 1
    var inline11 int = 5
    var inline12 FnIterator__isize = __goml_builtin_range(inline10, inline11)
    t13 = inline12
    var t14 closure_env_main_2 = closure_env_main_2{}
    var t15 func(int, int) int = func(p0 int, p1 int) int {
        return _goml_m_inherent_i_closure__env__main__2_i_closure__env__main__2_i_apply(t14, p0, p1)
    }
    var sum__0 int = _goml_m_std_p_iter_p_fold____A__isize____I__FnIterator_l_isize_r_____T__isize(t13, 0, t15)
    var inline8 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(sum__0)
    _goml_runtime_core_string_println(inline8)
    var t16 FnIterator__isize
    var inline5 int = 1
    var inline6 int = 4
    var inline7 FnIterator__isize = __goml_builtin_range(inline5, inline6)
    t16 = inline7
    var t17 closure_env_main_3 = closure_env_main_3{}
    var t18 func(int) string = func(p0 int) string {
        return _goml_m_inherent_i_closure__env__main__3_i_closure__env__main__3_i_apply(t17, p0)
    }
    var t19 FnIterator__string
    var inline2 closure_env_std_iter_map_A_isize_B_string_I_FnIterator_isize_9 = closure_env_std_iter_map_A_isize_B_string_I_FnIterator_isize_9{
        iterator_0: t16,
        map_fn_1: t18,
    }
    var inline3 func() Option__string = func() Option__string {
        return _goml_m_inherent_i_closure__en_h38480fcf7085e17f84e4b3380fa38e2f_size__9_i_apply(inline2)
    }
    var inline4 FnIterator__string = _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__string(inline3)
    t19 = inline4
    var texts__0 *_goml_vec_string = _goml_m_std_p_iter_p_collect____I__FnIterator_l_string_r_____T__string(t19)
    var for_limit0 int = vec_len__Vec_6string(texts__0)
    var for_index0 int = 0
    Loop_loop1:
    for {
        var t20 bool = for_index0 < for_limit0
        if t20 {
            var for_item0 string = vec_get__Vec_6string(texts__0, for_index0)
            var t21 int = for_index0 + 1
            for_index0 = t21
            var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(for_item0)
            _goml_runtime_core_string_println(inline0)
            continue
        } else {
            break Loop_loop1
        }
    }
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

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__i32(self__0 *ref_int32_x, value__0 int32) struct{} {
    ref_set__Ref_5int32(self__0, value__0)
    return struct{}{}
}

func println__T_i32(value__0 int32) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_int32_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func println__T_string(value__0 string) struct{} {
    var t0 string
    t0 = value__0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func convert_to__T_i32__V_Token(value__0 Token) int32 {
    return 7
}

func _goml_m_std_p_iter_p_fold____A__isize____I__FnIterator_l_isize_r_____T__isize(iterator__0 FnIterator__isize, initial__0 int, combine__0 func(int, int) int) int {
    var accumulator__0 int = initial__0
    Loop_loop0:
    for {
        var mtmp0 Option__isize
        var inline0 func() Option__isize = iterator__0.next_fn
        var inline1 Option__isize = inline0()
        mtmp0 = inline1
        switch mtmp0._tag {
        case 1:
            var x0 int = mtmp0._p0
            var t0 int = combine__0(accumulator__0, x0)
            accumulator__0 = t0
            continue
        default:
            break Loop_loop0
        }
    }
    return accumulator__0
}

func _goml_m_std_p_iter_p_collect____I__FnIterator_l_string_r_____T__string(iterator__0 FnIterator__string) *_goml_vec_string {
    var values__0 *_goml_vec_string
    var inline3 *_goml_vec_string = vec_new__Vec_6string()
    values__0 = inline3
    Loop_loop0:
    for {
        var mtmp0 Option__string
        var inline1 func() Option__string = iterator__0.next_fn
        var inline2 Option__string = inline1()
        mtmp0 = inline2
        switch mtmp0._tag {
        case 1:
            var x0 string = mtmp0._p0
            vec_push__Vec_6string(values__0, x0)
            continue
        default:
            break Loop_loop0
        }
    }
    return values__0
}

func _goml_m_trait__impl_i_ToString_i_i32_i_to__string(self__0 int32) string {
    var inline0 int64 = int64(int32(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__i32(next_fn__0 func() Option__i32) FnIterator__i32 {
    var t0 FnIterator__i32 = FnIterator__i32{
        next_fn: next_fn__0,
    }
    return t0
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__isize(value__0 int) *ref_int_x {
    var t0 *ref_int_x = ref__Ref_3int(value__0)
    return t0
}

func __goml_builtin_range(start__0 int, end__0 int) FnIterator__isize {
    var current__0 *ref_int_x = ref__Ref_3int(start__0)
    var t0 closure_env_goml_builtin_range_10 = closure_env_goml_builtin_range_10{
        current_0: current__0,
        end_1: end__0,
    }
    var t1 func() Option__isize = func() Option__isize {
        return _goml_m_inherent_i_closure__en_hfe17758c34b24eefa68773778f3d9cb8_nge__10_i_apply(t0)
    }
    var inline0 FnIterator__isize = FnIterator__isize{
        next_fn: t1,
    }
    return inline0
}

func _goml_m_trait__impl_i_ToString_i_isize_i_to__string(self__0 int) string {
    var inline0 int64 = int64(int(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__string(next_fn__0 func() Option__string) FnIterator__string {
    var t0 FnIterator__string = FnIterator__string{
        next_fn: next_fn__0,
    }
    return t0
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

func _goml_m_inherent_i_closure__env__main__0_i_closure__env__main__0_i_apply(env0 closure_env_main_0, value__0 int32) int32 {
    var t0 int32 = value__0 * 2
    return t0
}

func _goml_m_inherent_i_closure__env__main__1_i_closure__env__main__1_i_apply(env0 closure_env_main_1, value__0 int32) bool {
    var t0 bool = value__0 > 4
    return t0
}

func _goml_m_inherent_i_closure__env__main__2_i_closure__env__main__2_i_apply(env0 closure_env_main_2, total__0 int, value__0 int) int {
    var t0 int = total__0 + value__0
    return t0
}

func _goml_m_inherent_i_closure__env__main__3_i_closure__env__main__3_i_apply(env0 closure_env_main_3, value__0 int) string {
    var t0 string
    var inline0 string = __goml_builtin_int_to_string(value__0)
    t0 = inline0
    var t1 string = "v" + t0
    return t1
}

func _goml_m_inherent_i_closure__en_h645d7a9bc4d79b01cd03faf046af5461_nter__4_i_apply(env0 closure_env_std_iter_map_A_i32_B_i32_I_Counter_4, value__0 int32) int32 {
    var map_fn__0 func(int32) int32 = env0.map_fn_0
    var t0 int32 = map_fn__0(value__0)
    return t0
}

func _goml_m_inherent_i_closure__en_h39f5b01177a579a305108bc4e47ab576_nter__5_i_apply(env0 closure_env_std_iter_map_A_i32_B_i32_I_Counter_5) Option__i32 {
    var iterator__0 Counter = env0.iterator_0
    var map_fn__0 func(int32) int32 = env0.map_fn_1
    var t0 Option__i32
    var inline3 *ref_int32_x = iterator__0.current
    var inline4 int32 = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__i32(inline3)
    var inline5 int32 = iterator__0.end
    var inline6 bool = inline4 < inline5
    if inline6 {
        var inline7 *ref_int32_x = iterator__0.current
        var inline8 int32 = inline4 + 1
        _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__i32(inline7, inline8)
        var inline10 Option__i32 = Option__i32(uint64(int64(inline4) + 2147483648) + 1)
        t0 = inline10
    } else {
        t0 = Option__i32(0)
    }
    var t1 closure_env_std_iter_map_A_i32_B_i32_I_Counter_4 = closure_env_std_iter_map_A_i32_B_i32_I_Counter_4{
        map_fn_0: map_fn__0,
    }
    var t2 func(int32) int32 = func(p0 int32) int32 {
        return _goml_m_inherent_i_closure__en_h645d7a9bc4d79b01cd03faf046af5461_nter__4_i_apply(t1, p0)
    }
    switch t0 != Option__i32(0) {
    case false:
        return Option__i32(0)
    case true:
        var inline0 int32 = int32(int64(uint64(t0) - 1) - 2147483648)
        var inline1 int32 = t2(inline0)
        var inline2 Option__i32 = Option__i32(uint64(int64(inline1) + 2147483648) + 1)
        return inline2
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_inherent_i_closure__en_h920672151d9ffb42cb9c90456440c3f4__i32__6_i_apply(env0 closure_env_std_iter_filter_I_FnIterator_i32_T_i32_6) Option__i32 {
    var iterator__0 FnIterator__i32 = env0.iterator_0
    var predicate__0 func(int32) bool = env0.predicate_1
    Loop_loop0:
    for {
        var mtmp0 Option__i32
        var inline0 func() Option__i32 = iterator__0.next_fn
        var inline1 Option__i32 = inline0()
        mtmp0 = inline1
        switch mtmp0 != Option__i32(0) {
        case true:
            var x0 int32 = int32(int64(uint64(mtmp0) - 1) - 2147483648)
            var t0 bool = predicate__0(x0)
            if t0 {
                var t1 Option__i32 = Option__i32(uint64(int64(x0) + 2147483648) + 1)
                return t1
            } else {
                continue
            }
        default:
            break Loop_loop0
        }
    }
    return Option__i32(0)
}

func _goml_m_inherent_i_closure__en_h62d5c49070b32f607a729bd540373bb1__i32__7_i_apply(env0 closure_env_std_iter_take_I_FnIterator_i32_7) Option__i32 {
    var remaining__0 *ref_int_x = env0.remaining_0
    var iterator__0 FnIterator__i32 = env0.iterator_1
    var t0 int
    var inline4 int = ref_get__Ref_3int(remaining__0)
    t0 = inline4
    var t1 bool = t0 == 0
    if t1 {
        return Option__i32(0)
    } else {
        var t2 int
        var inline3 int = ref_get__Ref_3int(remaining__0)
        t2 = inline3
        var t3 int = t2 - 1
        ref_set__Ref_3int(remaining__0, t3)
        var inline0 func() Option__i32 = iterator__0.next_fn
        var inline1 Option__i32 = inline0()
        return inline1
    }
}

func _goml_m_inherent_i_closure__en_he6638ea00a62480db992a15b1db4c712_size__8_i_apply(env0 closure_env_std_iter_map_A_isize_B_string_I_FnIterator_isize_8, value__0 int) string {
    var map_fn__0 func(int) string = env0.map_fn_0
    var t0 string = map_fn__0(value__0)
    return t0
}

func _goml_m_inherent_i_closure__en_h38480fcf7085e17f84e4b3380fa38e2f_size__9_i_apply(env0 closure_env_std_iter_map_A_isize_B_string_I_FnIterator_isize_9) Option__string {
    var iterator__0 FnIterator__isize = env0.iterator_0
    var map_fn__0 func(int) string = env0.map_fn_1
    var t0 Option__isize
    var inline3 func() Option__isize = iterator__0.next_fn
    var inline4 Option__isize = inline3()
    t0 = inline4
    var t1 closure_env_std_iter_map_A_isize_B_string_I_FnIterator_isize_8 = closure_env_std_iter_map_A_isize_B_string_I_FnIterator_isize_8{
        map_fn_0: map_fn__0,
    }
    var t2 func(int) string = func(p0 int) string {
        return _goml_m_inherent_i_closure__en_he6638ea00a62480db992a15b1db4c712_size__8_i_apply(t1, p0)
    }
    switch t0._tag {
    case 0:
        return Option__string{
            _tag: 0,
        }
    case 1:
        var inline0 int = t0._p0
        var inline1 string = t2(inline0)
        var inline2 Option__string = Option__string{
            _p0: inline1,
            _tag: 1,
        }
        return inline2
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_inherent_i_closure__en_hfe17758c34b24eefa68773778f3d9cb8_nge__10_i_apply(env0 closure_env_goml_builtin_range_10) Option__isize {
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
