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

func _goml_runtime_core_string_print(s string) struct{} {
    _goml_os.Stdout.WriteString(s)
    return struct{}{}
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
    var sum__0 *ref_int_x
    var inline30 int = 0
    var inline31 *ref_int_x = ref__Ref_3int(inline30)
    sum__0 = inline31
    var i__0 *ref_int_x
    var inline28 int = 1
    var inline29 *ref_int_x = ref__Ref_3int(inline28)
    i__0 = inline29
    Loop_loop0:
    for {
        var t10 int
        var inline27 int = ref_get__Ref_3int(i__0)
        t10 = inline27
        var t11 bool = t10 <= 100
        if t11 {
            var t12 int
            var inline26 int = ref_get__Ref_3int(i__0)
            t12 = inline26
            var t13 bool = t12 == 50
            if t13 {
                break Loop_loop0
            } else {
                var t14 int
                var inline25 int = ref_get__Ref_3int(sum__0)
                t14 = inline25
                var t15 int
                var inline24 int = ref_get__Ref_3int(i__0)
                t15 = inline24
                var t16 int = t14 + t15
                ref_set__Ref_3int(sum__0, t16)
                var t17 int
                var inline22 int = ref_get__Ref_3int(i__0)
                t17 = inline22
                var t18 int = t17 + 1
                ref_set__Ref_3int(i__0, t18)
                continue
            }
        } else {
            break Loop_loop0
        }
    }
    var inline18 string = "sum up to break: "
    var inline19 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline18)
    _goml_runtime_core_string_print(inline19)
    var t0 int
    var inline17 int = ref_get__Ref_3int(sum__0)
    t0 = inline17
    var inline15 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t0)
    _goml_runtime_core_string_println(inline15)
    var even_sum__0 *ref_int_x
    var inline13 int = 0
    var inline14 *ref_int_x = ref__Ref_3int(inline13)
    even_sum__0 = inline14
    var j__0 *ref_int_x
    var inline11 int = 1
    var inline12 *ref_int_x = ref__Ref_3int(inline11)
    j__0 = inline12
    Loop_loop1:
    for {
        var t2 int
        var inline10 int = ref_get__Ref_3int(j__0)
        t2 = inline10
        var t3 bool = t2 <= 10
        if t3 {
            var cur__0 int
            var inline9 int = ref_get__Ref_3int(j__0)
            cur__0 = inline9
            var t4 int = cur__0 + 1
            ref_set__Ref_3int(j__0, t4)
            var t5 int = cur__0 / 2
            var t6 int = t5 * 2
            var t7 bool = cur__0 == t6
            if t7 {
                var t8 int
                var inline7 int = ref_get__Ref_3int(even_sum__0)
                t8 = inline7
                var t9 int = t8 + cur__0
                ref_set__Ref_3int(even_sum__0, t9)
                continue
            } else {
                continue
            }
        } else {
            break Loop_loop1
        }
    }
    var inline3 string = "even sum: "
    var inline4 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline3)
    _goml_runtime_core_string_print(inline4)
    var t1 int
    var inline2 int = ref_get__Ref_3int(even_sum__0)
    t1 = inline2
    var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t1)
    _goml_runtime_core_string_println(inline0)
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

func main() {
    main0()
}
