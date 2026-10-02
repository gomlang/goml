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

type Light__i32 uint8

const (
    Light__i32_Red Light__i32 = 0
    Light__i32_Green Light__i32 = 1
)

type Light__string uint8

const (
    Light__string_Red Light__string = 0
    Light__string_Green Light__string = 1
)

func main0() struct{} {
    var t0 Light__i32
    t0 = Light__i32_Green
    var t1 string
    switch t0 {
    case Light__i32_Red:
        t1 = "ri"
    case Light__i32_Green:
        t1 = "gi"
    default:
        panic("non-exhaustive match")
    }
    var inline2 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t1)
    _goml_runtime_core_string_println(inline2)
    var t2 Light__string
    t2 = Light__string_Red
    var t3 string
    switch t2 {
    case Light__string_Red:
        t3 = "rs"
    case Light__string_Green:
        t3 = "gs"
    default:
        panic("non-exhaustive match")
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
