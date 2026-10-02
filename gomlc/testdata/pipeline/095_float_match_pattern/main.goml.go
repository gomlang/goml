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
    var inline13 float64 = 0
    switch inline13 {
    case 0:
        t0 = "zero"
    case 1:
        t0 = "one"
    case -1:
        t0 = "minus one"
    case 3.14:
        t0 = "pi"
    default:
        t0 = "other"
    }
    var inline11 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t0)
    _goml_runtime_core_string_println(inline11)
    var t1 string
    var inline10 float64 = 1
    switch inline10 {
    case 0:
        t1 = "zero"
    case 1:
        t1 = "one"
    case -1:
        t1 = "minus one"
    case 3.14:
        t1 = "pi"
    default:
        t1 = "other"
    }
    var inline8 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t1)
    _goml_runtime_core_string_println(inline8)
    var t2_operand float64 = 1
    var t2 float64 = -t2_operand
    var t3 string
    switch t2 {
    case 0:
        t3 = "zero"
    case 1:
        t3 = "one"
    case -1:
        t3 = "minus one"
    case 3.14:
        t3 = "pi"
    default:
        t3 = "other"
    }
    var inline6 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t3)
    _goml_runtime_core_string_println(inline6)
    var t4 string
    var inline5 float64 = 3.14
    switch inline5 {
    case 0:
        t4 = "zero"
    case 1:
        t4 = "one"
    case -1:
        t4 = "minus one"
    case 3.14:
        t4 = "pi"
    default:
        t4 = "other"
    }
    var inline3 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t4)
    _goml_runtime_core_string_println(inline3)
    var t5 string
    var inline2 float64 = 42
    switch inline2 {
    case 0:
        t5 = "zero"
    case 1:
        t5 = "one"
    case -1:
        t5 = "minus one"
    case 3.14:
        t5 = "pi"
    default:
        t5 = "other"
    }
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t5)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func main() {
    main0()
}
