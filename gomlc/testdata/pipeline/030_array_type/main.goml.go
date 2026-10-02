package main

import (
    _goml_os "os"
)

func _goml_runtime_core_string_print(s string) struct{} {
    _goml_os.Stdout.WriteString(s)
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

type Buffer struct {
    values [3]int32
}

type Ordering uint8

func main0() struct{} {
    var inline0 string = "array"
    var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
    _goml_runtime_core_string_print(inline1)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func main() {
    main0()
}
