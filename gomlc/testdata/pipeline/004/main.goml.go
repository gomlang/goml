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

type Color uint8

const (
    Red Color = 0
    Green Color = 1
    Blue Color = 2
)

func main0() bool {
    var x0 Color = Blue
    var x1 Color = Blue
    switch x1 {
    case Red:
        switch x0 {
        case Red:
            return true
        default:
            return false
        }
    case Green:
        switch x0 {
        case Red:
            return true
        default:
            return false
        }
    case Blue:
        switch x0 {
        case Blue:
            var inline0 bool = true
            var inline1 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(inline0)
            _goml_runtime_core_string_print(inline1)
            return false
        default:
            return false
        }
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
