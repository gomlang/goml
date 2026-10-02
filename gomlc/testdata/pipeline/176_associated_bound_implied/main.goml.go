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

type TextSource struct {
    value string
}

type Ordering uint8

func _goml_m_trait__impl_i_Label_i_string_i_label(self__0 string) string {
    var t0 string = "label:" + self__0
    return t0
}

func _goml_m_trait__impl_i_Source_i_TextSource_i_get(self__0 TextSource) string {
    var t0 string = self__0.value
    return t0
}

func main0() struct{} {
    var t0 TextSource = TextSource{
        value: "goml",
    }
    var t1 string
    var inline2 string = _goml_m_trait__impl_i_Source_i_TextSource_i_get(t0)
    var inline3 string = _goml_m_trait__impl_i_Label_i_string_i_label(inline2)
    t1 = inline3
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
