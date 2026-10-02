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

func _goml_intrinsic_missing(s string) struct{} {
    println("missing: " + s)
    panic("")
    return struct{}{}
}

func array_get__Array_2_3int(arr [2]int, index int) int {
    return arr[index]
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

func missing__string(s string) string {
    _goml_intrinsic_missing(s)
    var ret string
    return ret
}

type Tuple2_3int_5uint8 struct {
    _0 int
    _1 uint8
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

const (
    RATIO float64 = 1.5
    COMPUTED_ENABLED bool = true
    CLASSIFIED_AT_COMPILE_TIME int = 1
)

func classify(value__0 int) string {
    switch value__0 {
    case 42:
        return "known"
    case 7:
        return "known"
    default:
        return "other"
    }
}

func classify_bool(value__0 bool) string {
    switch value__0 {
    case true:
        return "enabled"
    case false:
        return "disabled"
    default:
        panic("non-exhaustive match")
    }
}

func classify_computed_bool(value__0 bool) string {
    var t0 bool = value__0 == COMPUTED_ENABLED
    if t0 {
        return "enabled"
    } else {
        var t1 bool = value__0 == false
        if t1 {
            return "disabled"
        } else {
            var t2 string = missing__string("")
            return t2
        }
    }
}

func classify_comptime_guard(value__0 bool) string {
    var t0 bool = value__0 == true
    var jp0 string
    if t0 {
        jp0 = "enabled"
    } else {
        var t1 bool = value__0 == false
        if t1 {
            jp0 = "disabled"
        } else {
            var t2 string = missing__string("")
            jp0 = t2
        }
    }
    return jp0
}

func classify_pair(value__0 Tuple2_3int_5uint8) bool {
    var x0 int = value__0._0
    var x1 uint8 = value__0._1
    switch x1 {
    case 65:
        switch x0 {
        case 42:
            return true
        default:
            return false
        }
    default:
        return false
    }
}

func classify_string(value__0 string) bool {
    switch value__0 {
    case "hello":
        return true
    default:
        return false
    }
}

func classify_float(value__0 float64) bool {
    var t0 bool = value__0 == RATIO
    if t0 {
        return true
    } else {
        return false
    }
}

func for_binding() int {
    var total__0 *ref_int_x
    var inline3 int = 0
    var inline4 *ref_int_x = ref__Ref_3int(inline3)
    total__0 = inline4
    var for_source0 [2]int = [2]int{1, 2}
    var for_limit0 int = 2
    var for_index0 int = 0
    Loop_loop0:
    for {
        var t0 bool = for_index0 < for_limit0
        if t0 {
            var for_item0 int = array_get__Array_2_3int(for_source0, for_index0)
            var t1 int = for_index0 + 1
            for_index0 = t1
            var t2 int
            var inline2 int = ref_get__Ref_3int(total__0)
            t2 = inline2
            var t3 int = t2 + for_item0
            ref_set__Ref_3int(total__0, t3)
            continue
        } else {
            break Loop_loop0
        }
    }
    var inline0 int = ref_get__Ref_3int(total__0)
    return inline0
}

func main0() struct{} {
    var t0 string = classify(42)
    println__T_string(t0)
    var t1 string = classify(7)
    println__T_string(t1)
    var t2 string = classify(0)
    println__T_string(t2)
    var t3 string = classify_bool(true)
    println__T_string(t3)
    var t4 string = classify_bool(false)
    println__T_string(t4)
    var t5 string = classify_computed_bool(true)
    println__T_string(t5)
    var t6 string = classify_computed_bool(false)
    println__T_string(t6)
    var t7 string = classify_comptime_guard(true)
    println__T_string(t7)
    var t8 string = classify_comptime_guard(false)
    println__T_string(t8)
    var t9 Tuple2_3int_5uint8 = Tuple2_3int_5uint8{
        _0: 42,
        _1: 65,
    }
    var t10 bool = classify_pair(t9)
    println__T_bool(t10)
    var t11 Tuple2_3int_5uint8 = Tuple2_3int_5uint8{
        _0: 42,
        _1: 66,
    }
    var t12 bool = classify_pair(t11)
    println__T_bool(t12)
    var t13 bool = classify_string("hello")
    println__T_bool(t13)
    var t14 bool = classify_float(1.5)
    var inline22 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t14)
    _goml_runtime_core_string_println(inline22)
    var t15 int
    var inline21 int = 9
    t15 = inline21
    var inline19 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t15)
    _goml_runtime_core_string_println(inline19)
    var t16 int
    var inline18 int = 11
    t16 = inline18
    var inline16 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t16)
    _goml_runtime_core_string_println(inline16)
    var t17 bool
    var inline15 int = 42
    switch inline15 {
    case 42:
        t17 = true
    default:
        t17 = false
    }
    var inline13 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t17)
    _goml_runtime_core_string_println(inline13)
    var t18 bool
    var inline12 int = 41
    switch inline12 {
    case 42:
        t18 = true
    default:
        t18 = false
    }
    var inline10 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t18)
    _goml_runtime_core_string_println(inline10)
    var t19 bool
    var inline9 int = 42
    switch inline9 {
    case 42:
        t19 = true
    default:
        t19 = false
    }
    var inline7 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t19)
    _goml_runtime_core_string_println(inline7)
    var t20 bool
    var inline6 int = 41
    switch inline6 {
    case 42:
        t20 = true
    default:
        t20 = false
    }
    var inline4 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t20)
    _goml_runtime_core_string_println(inline4)
    var t21 int = for_binding()
    var inline2 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t21)
    _goml_runtime_core_string_println(inline2)
    var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(CLASSIFIED_AT_COMPILE_TIME)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func println__T_string(value__0 string) struct{} {
    var t0 string
    t0 = value__0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func println__T_bool(value__0 bool) struct{} {
    var t0 string
    var inline0 string = _goml_runtime_core_bool_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_bool_i_to__string(self__0 bool) string {
    var t0 string = _goml_runtime_core_bool_to_string(self__0)
    return t0
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
