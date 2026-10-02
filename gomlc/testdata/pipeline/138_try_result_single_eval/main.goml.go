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

type Result__i32__string struct {
    _p1 string
    _p0 int32
    _tag uint8
}

func use_try(counter__0 *ref_int32_x, ok__0 bool) Result__i32__string {
    var mtmp0 Result__i32__string
    var inline1 int32 = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__i32(counter__0)
    var inline2 int32 = inline1 + 1
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__i32(counter__0, inline2)
    if ok__0 {
        var inline4 int32 = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__i32(counter__0)
        var inline5 Result__i32__string = Result__i32__string{
            _p0: inline4,
            _tag: 0,
        }
        mtmp0 = inline5
    } else {
        var inline6 Result__i32__string = Result__i32__string{
            _p1: "bump failed",
            _tag: 1,
        }
        mtmp0 = inline6
    }
    var jp0 int32
    switch mtmp0._tag {
    case 0:
        var x0 int32 = mtmp0._p0
        jp0 = x0
        var t0 int32
        var inline0 int32 = ref_get__Ref_5int32(counter__0)
        t0 = inline0
        var t1 int32 = jp0 + t0
        var t2 Result__i32__string = Result__i32__string{
            _p0: t1,
            _tag: 0,
        }
        return t2
    case 1:
        var x1 string = mtmp0._p1
        var t3 Result__i32__string = Result__i32__string{
            _p1: x1,
            _tag: 1,
        }
        return t3
    default:
        panic("non-exhaustive match")
    }
}

func show(res__0 Result__i32__string) string {
    switch res__0._tag {
    case 0:
        var x0 int32 = res__0._p0
        var t0 string
        var inline0 string = __goml_builtin_int32_to_string(x0)
        t0 = inline0
        var t1 string = "ok " + t0
        return t1
    case 1:
        var x1 string = res__0._p1
        var t2 string = "err " + x1
        return t2
    default:
        panic("non-exhaustive match")
    }
}

func main0() struct{} {
    var t0 string
    var inline12 bool = true
    var inline13 *ref_int32_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__i32(0)
    var inline14 Result__i32__string = use_try(inline13, inline12)
    var inline15 string = show(inline14)
    var inline16 string = inline15 + " count="
    var inline17 int32 = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__i32(inline13)
    var inline18 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline17)
    var inline19 string = inline16 + inline18
    t0 = inline19
    var inline10 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t0)
    _goml_runtime_core_string_println(inline10)
    var t1 string
    var inline2 bool = false
    var inline3 *ref_int32_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__i32(0)
    var inline4 Result__i32__string = use_try(inline3, inline2)
    var inline5 string = show(inline4)
    var inline6 string = inline5 + " count="
    var inline7 int32 = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__i32(inline3)
    var inline8 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline7)
    var inline9 string = inline6 + inline8
    t1 = inline9
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t1)
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

func _goml_m_inherent_i_i32_i_i32_i_to__string(self__0 int32) string {
    var inline0 int64 = int64(int32(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__i32(value__0 int32) *ref_int32_x {
    var t0 *ref_int32_x = ref__Ref_5int32(value__0)
    return t0
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

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
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
