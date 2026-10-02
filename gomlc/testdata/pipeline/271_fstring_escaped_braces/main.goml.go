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

func main0() struct{} {
    var t0 string = "" + "{{"
    var t1 bool = t0 == "{{"
    var inline6 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t1)
    _goml_runtime_core_string_println(inline6)
    var t2 string = "" + "}}"
    var t3 bool = t2 == "}}"
    var inline4 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t3)
    _goml_runtime_core_string_println(inline4)
    var t4 string = "" + "{{"
    var t5 bool = t4 == "{{"
    var inline2 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t5)
    _goml_runtime_core_string_println(inline2)
    var t6 string = "" + "}}"
    var t7 bool = t6 == "}}"
    var inline0 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t7)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_bool_i_to__string(self__0 bool) string {
    var t0 string = _goml_runtime_core_bool_to_string(self__0)
    return t0
}

func main() {
    main0()
}
