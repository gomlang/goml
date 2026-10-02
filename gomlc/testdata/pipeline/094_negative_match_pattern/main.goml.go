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

type Ordering uint8

func main0() struct{} {
    var t0 string
    var inline11 int32 = -1
    switch inline11 {
    case -1:
        t0 = "minus one"
    case 0:
        t0 = "zero"
    case 1:
        t0 = "one"
    default:
        t0 = "other"
    }
    var inline9 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t0)
    _goml_runtime_core_string_println(inline9)
    var t1 string
    var inline8 int32 = 0
    switch inline8 {
    case -1:
        t1 = "minus one"
    case 0:
        t1 = "zero"
    case 1:
        t1 = "one"
    default:
        t1 = "other"
    }
    var inline6 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t1)
    _goml_runtime_core_string_println(inline6)
    var t2 string
    var inline5 int32 = 1
    switch inline5 {
    case -1:
        t2 = "minus one"
    case 0:
        t2 = "zero"
    case 1:
        t2 = "one"
    default:
        t2 = "other"
    }
    var inline3 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t2)
    _goml_runtime_core_string_println(inline3)
    var t3 string
    var inline2 int32 = 42
    switch inline2 {
    case -1:
        t3 = "minus one"
    case 0:
        t3 = "zero"
    case 1:
        t3 = "one"
    default:
        t3 = "other"
    }
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t3)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func main() {
    main0()
}
