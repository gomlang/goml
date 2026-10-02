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

func _goml_runtime_core_string_println(s string) struct{} {
    _goml_os.Stdout.WriteString(s + "\n")
    return struct{}{}
}

type _goml_vec_uint32 struct {
    items []uint32
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

func test_int_comparisons() struct{} {
    var a__0 int32 = 10
    var b__0 int32 = 20
    var c__0 int32 = 10
    var less__0 bool = a__0 < b__0
    var t0 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(less__0)
    var t1 string = "10 < 20: " + t0
    var inline27 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t1)
    _goml_runtime_core_string_println(inline27)
    var greater__0 bool = b__0 > a__0
    var t2 string
    var inline26 string = _goml_runtime_core_bool_to_string(greater__0)
    t2 = inline26
    var t3 string = "20 > 10: " + t2
    var inline24 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t3)
    _goml_runtime_core_string_println(inline24)
    var less_eq1__0 bool = a__0 <= b__0
    var t4 string
    var inline23 string = _goml_runtime_core_bool_to_string(less_eq1__0)
    t4 = inline23
    var t5 string = "10 <= 20: " + t4
    var inline21 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t5)
    _goml_runtime_core_string_println(inline21)
    var less_eq2__0 bool = a__0 <= c__0
    var t6 string
    var inline20 string = _goml_runtime_core_bool_to_string(less_eq2__0)
    t6 = inline20
    var t7 string = "10 <= 10: " + t6
    var inline18 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t7)
    _goml_runtime_core_string_println(inline18)
    var greater_eq1__0 bool = b__0 >= a__0
    var t8 string
    var inline17 string = _goml_runtime_core_bool_to_string(greater_eq1__0)
    t8 = inline17
    var t9 string = "20 >= 10: " + t8
    var inline15 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t9)
    _goml_runtime_core_string_println(inline15)
    var greater_eq2__0 bool = c__0 >= a__0
    var t10 string
    var inline14 string = _goml_runtime_core_bool_to_string(greater_eq2__0)
    t10 = inline14
    var t11 string = "10 >= 10: " + t10
    var inline12 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t11)
    _goml_runtime_core_string_println(inline12)
    var eq1__0 bool = a__0 == c__0
    var t12 string
    var inline11 string = _goml_runtime_core_bool_to_string(eq1__0)
    t12 = inline11
    var t13 string = "10 == 10: " + t12
    var inline9 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t13)
    _goml_runtime_core_string_println(inline9)
    var eq2__0 bool = a__0 == b__0
    var t14 string
    var inline8 string = _goml_runtime_core_bool_to_string(eq2__0)
    t14 = inline8
    var t15 string = "10 == 20: " + t14
    var inline6 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t15)
    _goml_runtime_core_string_println(inline6)
    var neq1__0 bool = a__0 != b__0
    var t16 string
    var inline5 string = _goml_runtime_core_bool_to_string(neq1__0)
    t16 = inline5
    var t17 string = "10 != 20: " + t16
    var inline3 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t17)
    _goml_runtime_core_string_println(inline3)
    var neq2__0 bool = a__0 != c__0
    var t18 string
    var inline2 string = _goml_runtime_core_bool_to_string(neq2__0)
    t18 = inline2
    var t19 string = "10 != 10: " + t18
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t19)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func test_float_comparisons() struct{} {
    var x__0 float64 = 3.14
    var y__0 float64 = 2.71
    var z__0 float64 = 3.14
    var less__0 bool = y__0 < x__0
    var t0 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(less__0)
    var t1 string = "2.71 < 3.14: " + t0
    var inline27 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t1)
    _goml_runtime_core_string_println(inline27)
    var greater__0 bool = x__0 > y__0
    var t2 string
    var inline26 string = _goml_runtime_core_bool_to_string(greater__0)
    t2 = inline26
    var t3 string = "3.14 > 2.71: " + t2
    var inline24 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t3)
    _goml_runtime_core_string_println(inline24)
    var less_eq1__0 bool = y__0 <= x__0
    var t4 string
    var inline23 string = _goml_runtime_core_bool_to_string(less_eq1__0)
    t4 = inline23
    var t5 string = "2.71 <= 3.14: " + t4
    var inline21 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t5)
    _goml_runtime_core_string_println(inline21)
    var less_eq2__0 bool = x__0 <= z__0
    var t6 string
    var inline20 string = _goml_runtime_core_bool_to_string(less_eq2__0)
    t6 = inline20
    var t7 string = "3.14 <= 3.14: " + t6
    var inline18 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t7)
    _goml_runtime_core_string_println(inline18)
    var greater_eq1__0 bool = x__0 >= y__0
    var t8 string
    var inline17 string = _goml_runtime_core_bool_to_string(greater_eq1__0)
    t8 = inline17
    var t9 string = "3.14 >= 2.71: " + t8
    var inline15 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t9)
    _goml_runtime_core_string_println(inline15)
    var greater_eq2__0 bool = z__0 >= x__0
    var t10 string
    var inline14 string = _goml_runtime_core_bool_to_string(greater_eq2__0)
    t10 = inline14
    var t11 string = "3.14 >= 3.14: " + t10
    var inline12 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t11)
    _goml_runtime_core_string_println(inline12)
    var eq1__0 bool = x__0 == z__0
    var t12 string
    var inline11 string = _goml_runtime_core_bool_to_string(eq1__0)
    t12 = inline11
    var t13 string = "3.14 == 3.14: " + t12
    var inline9 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t13)
    _goml_runtime_core_string_println(inline9)
    var eq2__0 bool = x__0 == y__0
    var t14 string
    var inline8 string = _goml_runtime_core_bool_to_string(eq2__0)
    t14 = inline8
    var t15 string = "3.14 == 2.71: " + t14
    var inline6 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t15)
    _goml_runtime_core_string_println(inline6)
    var neq1__0 bool = x__0 != y__0
    var t16 string
    var inline5 string = _goml_runtime_core_bool_to_string(neq1__0)
    t16 = inline5
    var t17 string = "3.14 != 2.71: " + t16
    var inline3 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t17)
    _goml_runtime_core_string_println(inline3)
    var neq2__0 bool = x__0 != z__0
    var t18 string
    var inline2 string = _goml_runtime_core_bool_to_string(neq2__0)
    t18 = inline2
    var t19 string = "3.14 != 3.14: " + t18
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t19)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func main0() struct{} {
    var inline6 string = "=== Integer Comparisons ==="
    var inline7 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline6)
    _goml_runtime_core_string_println(inline7)
    test_int_comparisons()
    var inline3 string = ""
    var inline4 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline3)
    _goml_runtime_core_string_println(inline4)
    var inline0 string = "=== Float Comparisons ==="
    var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
    _goml_runtime_core_string_println(inline1)
    test_float_comparisons()
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_bool_i_to__string(self__0 bool) string {
    var t0 string = _goml_runtime_core_bool_to_string(self__0)
    return t0
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func main() {
    main0()
}
