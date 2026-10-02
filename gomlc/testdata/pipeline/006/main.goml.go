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

type Ordering uint8

func main0() struct{} {
    var x0 bool = true
    var x1 bool = false
    switch x0 {
    case true:
        var inline2 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(x1)
        _goml_runtime_core_string_print(inline2)
    case false:
    default:
        panic("non-exhaustive match")
    }
    var x2 bool = true
    var x3 bool = true
    switch x2 {
    case true:
        var inline0 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(x3)
        _goml_runtime_core_string_print(inline0)
        return struct{}{}
    case false:
        return struct{}{}
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_trait__impl_i_ToString_i_bool_i_to__string(self__0 bool) string {
    var t0 string = _goml_runtime_core_bool_to_string(self__0)
    return t0
}

func main() {
    main0()
}
