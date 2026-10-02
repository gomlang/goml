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

type ref_int_x struct {
    value int
}

func ref__Ref_3int(value int) *ref_int_x {
    return &ref_int_x{
        value: value,
    }
}

func ptr_eq__Ref_3int(a *ref_int_x, b *ref_int_x) bool {
    return a == b
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
    var a__0 *ref_int_x
    var inline6 int = 1
    var inline7 *ref_int_x = ref__Ref_3int(inline6)
    a__0 = inline7
    var c__0 *ref_int_x
    var inline4 int = 1
    var inline5 *ref_int_x = ref__Ref_3int(inline4)
    c__0 = inline5
    var t0 bool = ptr_eq__Ref_3int(a__0, a__0)
    var inline2 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t0)
    _goml_runtime_core_string_println(inline2)
    var t1 bool = ptr_eq__Ref_3int(a__0, c__0)
    var inline0 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t1)
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
