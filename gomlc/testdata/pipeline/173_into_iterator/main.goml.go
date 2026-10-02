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

type Numbers struct {
    values *_goml_vec_int32
    conversions *ref_int32_x
}

type FnIterator__i32 struct {
    next_fn func() Option__i32
}

type closure_env_inherent_Vec_Vec_T_iter_T_i32_0 struct {
    index_0 *ref_int_x
    len_1 int
    self_2 *_goml_vec_int32
}

type closure_env_inherent_Slice_Slice_T_iter_T_i32_1 struct {
    index_0 *ref_int_x
    len_1 int
    self_2 []int32
}

type Ordering uint8

type Option__i32 uint64

func main0() struct{} {
    var builds__0 *ref_int32_x
    var inline32 int32 = 0
    var inline33 *ref_int32_x = ref__Ref_5int32(inline32)
    builds__0 = inline33
    var conversions__0 *ref_int32_x
    var inline30 int32 = 0
    var inline31 *ref_int32_x = ref__Ref_5int32(inline30)
    conversions__0 = inline31
    var t0 Numbers
    var inline22 int32 = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__i32(builds__0)
    var inline23 int32 = inline22 + 1
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__i32(builds__0, inline23)
    var inline25 *_goml_vec_int32 = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_new____T__i32()
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_push____T__i32(inline25, 1)
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_push____T__i32(inline25, 2)
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_push____T__i32(inline25, 3)
    var inline29 Numbers = Numbers{
        values: inline25,
        conversions: conversions__0,
    }
    t0 = inline29
    var t1 int32 = sum__S_Numbers(t0)
    var inline20 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(t1)
    _goml_runtime_core_string_println(inline20)
    var t2 int32
    var inline19 int32 = ref_get__Ref_5int32(builds__0)
    t2 = inline19
    var inline17 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(t2)
    _goml_runtime_core_string_println(inline17)
    var t3 int32
    var inline16 int32 = ref_get__Ref_5int32(conversions__0)
    t3 = inline16
    var inline14 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(t3)
    _goml_runtime_core_string_println(inline14)
    var values__0 *_goml_vec_int32
    var inline13 *_goml_vec_int32 = vec_new__Vec_5int32()
    values__0 = inline13
    var inline11 int32 = 10
    vec_push__Vec_5int32(values__0, inline11)
    var inline9 int32 = 20
    vec_push__Vec_5int32(values__0, inline9)
    var inline7 int32 = 30
    vec_push__Vec_5int32(values__0, inline7)
    var t4 int32 = _goml_m_sum____S__Vec_l_i32_r_(values__0)
    var inline5 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(t4)
    _goml_runtime_core_string_println(inline5)
    var t5 []int32
    var inline2 int = 1
    var inline3 int = 3
    var inline4 []int32 = values__0.items[inline2:inline3]
    t5 = inline4
    var t6 int32 = _goml_m_sum____S__Slice_l_i32_r_(t5)
    var inline0 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(t6)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__i32(self__0 *ref_int32_x) int32 {
    var t0 int32 = ref_get__Ref_5int32(self__0)
    return t0
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__i32(self__0 *ref_int32_x, value__0 int32) struct{} {
    ref_set__Ref_5int32(self__0, value__0)
    return struct{}{}
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_iter____T__i32(self__0 *_goml_vec_int32) FnIterator__i32 {
    var index__0 *ref_int_x = ref__Ref_3int(0)
    var len__0 int
    var inline1 int = vec_len__Vec_5int32(self__0)
    len__0 = inline1
    var t0 closure_env_inherent_Vec_Vec_T_iter_T_i32_0 = closure_env_inherent_Vec_Vec_T_iter_T_i32_0{
        index_0: index__0,
        len_1: len__0,
        self_2: self__0,
    }
    var t1 func() Option__i32 = func() Option__i32 {
        return _goml_m_inherent_i_closure__en_h1275f72f5de770912182f2a5cc7ddfae__i32__0_i_apply(t0)
    }
    var inline0 FnIterator__i32 = FnIterator__i32{
        next_fn: t1,
    }
    return inline0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_new____T__i32() *_goml_vec_int32 {
    var t0 *_goml_vec_int32 = vec_new__Vec_5int32()
    return t0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_push____T__i32(self__0 *_goml_vec_int32, elem__0 int32) struct{} {
    vec_push__Vec_5int32(self__0, elem__0)
    return struct{}{}
}

func sum__S_Numbers(source__0 Numbers) int32 {
    var total__0 *ref_int32_x
    var inline12 int32 = 0
    var inline13 *ref_int32_x = ref__Ref_5int32(inline12)
    total__0 = inline13
    var for_iter0 FnIterator__i32
    var inline5 *ref_int32_x = source__0.conversions
    var inline6 *ref_int32_x = source__0.conversions
    var inline7 int32 = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__i32(inline6)
    var inline8 int32 = inline7 + 1
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__i32(inline5, inline8)
    var inline10 *_goml_vec_int32 = source__0.values
    var inline11 FnIterator__i32 = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_iter____T__i32(inline10)
    for_iter0 = inline11
    Loop_loop0:
    for {
        var for_next0 Option__i32
        var inline3 func() Option__i32 = for_iter0.next_fn
        var inline4 Option__i32 = inline3()
        for_next0 = inline4
        switch for_next0 != Option__i32(0) {
        case false:
            break Loop_loop0
        case true:
            var x0 int32 = int32(int64(uint64(for_next0) - 1) - 2147483648)
            var t0 int32
            var inline2 int32 = ref_get__Ref_5int32(total__0)
            t0 = inline2
            var t1 int32 = t0 + x0
            ref_set__Ref_5int32(total__0, t1)
            continue
        default:
            panic("non-exhaustive match")
        }
    }
    var inline0 int32 = ref_get__Ref_5int32(total__0)
    return inline0
}

func _goml_m_sum____S__Vec_l_i32_r_(source__0 *_goml_vec_int32) int32 {
    var total__0 *ref_int32_x
    var inline6 int32 = 0
    var inline7 *ref_int32_x = ref__Ref_5int32(inline6)
    total__0 = inline7
    var for_iter0 FnIterator__i32
    var inline5 FnIterator__i32 = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_iter____T__i32(source__0)
    for_iter0 = inline5
    Loop_loop0:
    for {
        var for_next0 Option__i32
        var inline3 func() Option__i32 = for_iter0.next_fn
        var inline4 Option__i32 = inline3()
        for_next0 = inline4
        switch for_next0 != Option__i32(0) {
        case false:
            break Loop_loop0
        case true:
            var x0 int32 = int32(int64(uint64(for_next0) - 1) - 2147483648)
            var t0 int32
            var inline2 int32 = ref_get__Ref_5int32(total__0)
            t0 = inline2
            var t1 int32 = t0 + x0
            ref_set__Ref_5int32(total__0, t1)
            continue
        default:
            panic("non-exhaustive match")
        }
    }
    var inline0 int32 = ref_get__Ref_5int32(total__0)
    return inline0
}

func _goml_m_sum____S__Slice_l_i32_r_(source__0 []int32) int32 {
    var total__0 *ref_int32_x
    var inline6 int32 = 0
    var inline7 *ref_int32_x = ref__Ref_5int32(inline6)
    total__0 = inline7
    var for_iter0 FnIterator__i32
    var inline5 FnIterator__i32 = _goml_m_inherent_i_Slice_i_Slice_l_T_r__i_iter____T__i32(source__0)
    for_iter0 = inline5
    Loop_loop0:
    for {
        var for_next0 Option__i32
        var inline3 func() Option__i32 = for_iter0.next_fn
        var inline4 Option__i32 = inline3()
        for_next0 = inline4
        switch for_next0 != Option__i32(0) {
        case false:
            break Loop_loop0
        case true:
            var x0 int32 = int32(int64(uint64(for_next0) - 1) - 2147483648)
            var t0 int32
            var inline2 int32 = ref_get__Ref_5int32(total__0)
            t0 = inline2
            var t1 int32 = t0 + x0
            ref_set__Ref_5int32(total__0, t1)
            continue
        default:
            panic("non-exhaustive match")
        }
    }
    var inline0 int32 = ref_get__Ref_5int32(total__0)
    return inline0
}

func _goml_m_trait__impl_i_ToString_i_i32_i_to__string(self__0 int32) string {
    var inline0 int64 = int64(int32(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func _goml_m_inherent_i_Slice_i_Slice_l_T_r__i_iter____T__i32(self__0 []int32) FnIterator__i32 {
    var index__0 *ref_int_x = ref__Ref_3int(0)
    var len__0 int
    var inline1 int = len(self__0)
    len__0 = inline1
    var t0 closure_env_inherent_Slice_Slice_T_iter_T_i32_1 = closure_env_inherent_Slice_Slice_T_iter_T_i32_1{
        index_0: index__0,
        len_1: len__0,
        self_2: self__0,
    }
    var t1 func() Option__i32 = func() Option__i32 {
        return _goml_m_inherent_i_closure__en_h05f447217e0cf6cdcd746af967c79493__i32__1_i_apply(t0)
    }
    var inline0 FnIterator__i32 = FnIterator__i32{
        next_fn: t1,
    }
    return inline0
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

func _goml_m_inherent_i_closure__en_h1275f72f5de770912182f2a5cc7ddfae__i32__0_i_apply(env0 closure_env_inherent_Vec_Vec_T_iter_T_i32_0) Option__i32 {
    var index__0 *ref_int_x = env0.index_0
    var len__0 int = env0.len_1
    var self__0 *_goml_vec_int32 = env0.self_2
    var current__0 int = ref_get__Ref_3int(index__0)
    var t0 bool = current__0 < len__0
    if t0 {
        var value__0 int32 = vec_get__Vec_5int32(self__0, current__0)
        var t1 int = current__0 + 1
        ref_set__Ref_3int(index__0, t1)
        var t2 Option__i32 = Option__i32(uint64(int64(value__0) + 2147483648) + 1)
        return t2
    } else {
        return Option__i32(0)
    }
}

func _goml_m_inherent_i_closure__en_h05f447217e0cf6cdcd746af967c79493__i32__1_i_apply(env0 closure_env_inherent_Slice_Slice_T_iter_T_i32_1) Option__i32 {
    var index__0 *ref_int_x = env0.index_0
    var len__0 int = env0.len_1
    var self__0 []int32 = env0.self_2
    var current__0 int = ref_get__Ref_3int(index__0)
    var t0 bool = current__0 < len__0
    if t0 {
        var value__0 int32
        var inline0 int32 = self__0[current__0]
        value__0 = inline0
        var t1 int = current__0 + 1
        ref_set__Ref_3int(index__0, t1)
        var t2 Option__i32 = Option__i32(uint64(int64(value__0) + 2147483648) + 1)
        return t2
    } else {
        return Option__i32(0)
    }
}

func main() {
    main0()
}
