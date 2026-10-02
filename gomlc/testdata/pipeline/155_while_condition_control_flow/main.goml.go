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

func main0() struct{} {
    var i__0 *ref_int_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__isize(0)
    Loop_loop0:
    for {
        var t20 int
        var inline25 int = ref_get__Ref_3int(i__0)
        t20 = inline25
        var t21 bool = t20 < 3
        var jp2 bool
        if t21 {
            jp2 = true
        } else {
            jp2 = false
        }
        if jp2 {
            var t22 int
            var inline24 int = ref_get__Ref_3int(i__0)
            t22 = inline24
            var inline22 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t22)
            _goml_runtime_core_string_println(inline22)
            var t23 int
            var inline21 int = ref_get__Ref_3int(i__0)
            t23 = inline21
            var t24 int = t23 + 1
            ref_set__Ref_3int(i__0, t24)
            continue
        } else {
            break Loop_loop0
        }
    }
    var j__0 *ref_int_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__isize(0)
    var total__0 *ref_int_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__isize(0)
    Loop_loop1:
    for {
        var t9 int
        var inline19 int = ref_get__Ref_3int(j__0)
        t9 = inline19
        var t10 bool = t9 < 4
        var jp1 bool
        if t10 {
            var t16 int
            var inline18 int = ref_get__Ref_3int(j__0)
            t16 = inline18
            var t17 bool = t16 == 1
            if t17 {
                jp1 = true
            } else {
                var t18 int
                var inline17 int = ref_get__Ref_3int(j__0)
                t18 = inline17
                var t19 bool = t18 != 3
                jp1 = t19
            }
        } else {
            jp1 = false
        }
        if jp1 {
            var t11 int
            var inline16 int = ref_get__Ref_3int(total__0)
            t11 = inline16
            var t12 int
            var inline15 int = ref_get__Ref_3int(j__0)
            t12 = inline15
            var t13 int = t11 + t12
            ref_set__Ref_3int(total__0, t13)
            var t14 int
            var inline13 int = ref_get__Ref_3int(j__0)
            t14 = inline13
            var t15 int = t14 + 1
            ref_set__Ref_3int(j__0, t15)
            continue
        } else {
            break Loop_loop1
        }
    }
    var t0 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(total__0)
    println__T_isize(t0)
    var k__0 *ref_int_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__isize(0)
    var sum__0 *ref_int_x
    var inline10 int = 0
    var inline11 *ref_int_x = ref__Ref_3int(inline10)
    sum__0 = inline11
    Loop_loop2:
    for {
        var mtmp0 int
        var inline9 int = ref_get__Ref_3int(k__0)
        mtmp0 = inline9
        var jp0 bool
        switch mtmp0 {
        case 0:
            jp0 = true
        case 1:
            var t7 int
            var inline8 int = ref_get__Ref_3int(sum__0)
            t7 = inline8
            var t8 bool = t7 == 0
            if t8 {
                jp0 = true
            } else {
                jp0 = false
            }
        case 2:
            jp0 = true
        default:
            jp0 = false
        }
        if jp0 {
            var t2 int
            var inline7 int = ref_get__Ref_3int(sum__0)
            t2 = inline7
            var t3 int
            var inline6 int = ref_get__Ref_3int(k__0)
            t3 = inline6
            var t4 int = t2 + t3
            ref_set__Ref_3int(sum__0, t4)
            var t5 int
            var inline4 int = ref_get__Ref_3int(k__0)
            t5 = inline4
            var t6 int = t5 + 1
            ref_set__Ref_3int(k__0, t6)
            continue
        } else {
            break Loop_loop2
        }
    }
    var t1 int
    var inline2 int = ref_get__Ref_3int(sum__0)
    t1 = inline2
    var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t1)
    _goml_runtime_core_string_println(inline0)
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

func _goml_m_trait__impl_i_ToString_i_isize_i_to__string(self__0 int) string {
    var inline0 int64 = int64(int(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
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
