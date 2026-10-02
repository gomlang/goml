package main

import (
    _goml_os "os"
)

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

type Value struct {}

type Ordering uint8

func _goml_m_trait__impl_i_Base_i_Value_i_base(self__0 Value) string {
    return "B"
}

func _goml_m_trait__impl_i_Left_i_Value_i_left(self__0 Value) string {
    return "L"
}

func _goml_m_trait__impl_i_Right_i_Value_i_right(self__0 Value) string {
    return "R"
}

func _goml_m_trait__impl_i_Diamond_i_Value_i_diamond(self__0 Value) string {
    return "D"
}

func main0() struct{} {
    var t0 Value = Value{}
    var t1 string
    var inline2 string = _goml_m_trait__impl_i_Base_i_Value_i_base(t0)
    var inline3 string = _goml_m_trait__impl_i_Left_i_Value_i_left(t0)
    var inline4 string = inline2 + inline3
    var inline5 string = _goml_m_trait__impl_i_Right_i_Value_i_right(t0)
    var inline6 string = inline4 + inline5
    var inline7 string = _goml_m_trait__impl_i_Diamond_i_Value_i_diamond(t0)
    var inline8 string = inline6 + inline7
    t1 = inline8
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t1)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func main() {
    main0()
}
