package main

import (
    _goml_os "os"
)

func _goml_runtime_core_bool_to_string(x bool) string {
    if x {
        return "true"
    } else {
        return "false"
    }
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

type ref_Ref_5int32_x struct {
    value *ref_int32_x
}

func ref__Ref_10Ref_5int32(value *ref_int32_x) *ref_Ref_5int32_x {
    return &ref_Ref_5int32_x{
        value: value,
    }
}

func ref_get__Ref_10Ref_5int32(reference *ref_Ref_5int32_x) *ref_int32_x {
    return reference.value
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

func bump(cell__0 *ref_int32_x) int32 {
    var t0 int32
    var inline2 int32 = ref_get__Ref_5int32(cell__0)
    t0 = inline2
    var t1 int32 = t0 + 1
    ref_set__Ref_5int32(cell__0, t1)
    var inline0 int32 = ref_get__Ref_5int32(cell__0)
    return inline0
}

func flip(flag__0 *ref_bool_x) bool {
    var current__0 bool
    var inline2 bool = ref_get__Ref_4bool(flag__0)
    current__0 = inline2
    var t0 bool = !current__0
    ref_set__Ref_4bool(flag__0, t0)
    var inline0 bool = ref_get__Ref_4bool(flag__0)
    return inline0
}

func nested_total(cell__0 *ref_Ref_5int32_x) int32 {
    var inner__0 *ref_int32_x
    var inline3 *ref_int32_x = ref_get__Ref_10Ref_5int32(cell__0)
    inner__0 = inline3
    var before__0 int32
    var inline2 int32 = ref_get__Ref_5int32(inner__0)
    before__0 = inline2
    var t0 int32 = before__0 + 2
    ref_set__Ref_5int32(inner__0, t0)
    var t1 int32
    var inline0 int32 = ref_get__Ref_5int32(inner__0)
    t1 = inline0
    var t2 int32 = before__0 + t1
    return t2
}

func main0() struct{} {
    var counter__0 *ref_int32_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__i32(39)
    var toggler__0 *ref_bool_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__bool(false)
    var t0 *ref_int32_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__i32(3)
    var nested__0 *ref_Ref_5int32_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__Ref_l_i32_r_(t0)
    var bumped__0 int32 = bump(counter__0)
    var flipped__0 bool = flip(toggler__0)
    var flipped_again__0 bool = flip(toggler__0)
    var inner__0 *ref_int32_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__Ref_l_i32_r_(nested__0)
    var t1 int32
    var inline32 int32 = ref_get__Ref_5int32(inner__0)
    t1 = inline32
    var t2 int32 = t1 + bumped__0
    ref_set__Ref_5int32(inner__0, t2)
    var nested_total_val__0 int32 = nested_total(nested__0)
    var alias_total__0 int32
    var inline27 int32 = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__i32(counter__0)
    var inline28 int32 = inline27 + 5
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__i32(counter__0, inline28)
    var inline30 int32 = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__i32(counter__0)
    alias_total__0 = inline30
    var pair_total__0 int32
    var inline18 *ref_int32_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__i32(4)
    var inline19 *ref_int32_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__i32(6)
    var inline20 int32 = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__i32(inline18)
    var inline21 int32 = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__i32(inline19)
    var inline22 int32 = inline20 + inline21
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__i32(inline18, inline22)
    var inline24 int32 = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__i32(inline18)
    var inline25 int32 = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__i32(inline19)
    var inline26 int32 = inline24 + inline25
    pair_total__0 = inline26
    var reassigned__0 int32
    var inline13 *ref_int32_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__Ref_l_i32_r_(nested__0)
    var inline14 int32 = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__i32(inline13)
    var inline15 int32 = inline14 + 7
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__i32(inline13, inline15)
    var inline17 int32 = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__i32(inline13)
    reassigned__0 = inline17
    var bool_check__0 bool = !false
    var t3 int32
    var inline12 int32 = ref_get__Ref_5int32(counter__0)
    t3 = inline12
    var t4 int32 = bumped__0 + t3
    var t5 string
    var inline11 string = __goml_builtin_int32_to_string(t4)
    t5 = inline11
    var inline9 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t5)
    _goml_runtime_core_string_println(inline9)
    var t6 int32 = nested_total_val__0 + alias_total__0
    var t7 int32 = t6 + reassigned__0
    var t8 string
    var inline8 string = __goml_builtin_int32_to_string(t7)
    t8 = inline8
    var inline6 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t8)
    _goml_runtime_core_string_println(inline6)
    var t9 string
    var inline5 string = __goml_builtin_int32_to_string(pair_total__0)
    t9 = inline5
    var inline3 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t9)
    _goml_runtime_core_string_println(inline3)
    var jp0 bool
    if flipped__0 {
        jp0 = flipped_again__0
    } else {
        jp0 = false
    }
    var jp1 bool
    if jp0 {
        jp1 = bool_check__0
    } else {
        jp1 = false
    }
    var t10 string
    var inline2 string = _goml_runtime_core_bool_to_string(jp1)
    t10 = inline2
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t10)
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

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__Ref_l_i32_r_(self__0 *ref_Ref_5int32_x) *ref_int32_x {
    var t0 *ref_int32_x = ref_get__Ref_10Ref_5int32(self__0)
    return t0
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__i32(value__0 int32) *ref_int32_x {
    var t0 *ref_int32_x = ref__Ref_5int32(value__0)
    return t0
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__bool(value__0 bool) *ref_bool_x {
    var t0 *ref_bool_x = ref__Ref_4bool(value__0)
    return t0
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__Ref_l_i32_r_(value__0 *ref_int32_x) *ref_Ref_5int32_x {
    var t0 *ref_Ref_5int32_x = ref__Ref_10Ref_5int32(value__0)
    return t0
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

func main() {
    main0()
}
