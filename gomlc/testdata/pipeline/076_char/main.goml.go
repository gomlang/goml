package main

import (
    _goml_os "os"
)

func _goml_runtime_core_string_get(s string, i int) rune {
    return rune(s[i])
}

func _goml_runtime_core_char_to_string(x rune) string {
    return string(x)
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
    var c__0 rune = 65
    var t0 string
    var inline4 string = char_to_string(c__0)
    t0 = inline4
    var inline2 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t0)
    _goml_runtime_core_string_println(inline2)
    var d__0 rune = 98
    var jp0 string
    switch d__0 {
    case 97:
        jp0 = "A"
    case 98:
        jp0 = "B"
    default:
        jp0 = "?"
    }
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(jp0)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func char_to_string(value__0 rune) string {
    var t0 uint32 = uint32(rune(value__0))
    var t1 bool
    var inline0 bool = t0 <= 1114111
    if inline0 {
        var inline1 bool = t0 >= 55296
        var inline2 bool
        if inline1 {
            var inline4 bool = t0 <= 57343
            inline2 = inline4
        } else {
            inline2 = false
        }
        var inline3 bool = !inline2
        t1 = inline3
    } else {
        t1 = false
    }
    if t1 {
        var t2 string = _goml_runtime_core_char_to_string(value__0)
        return t2
    } else {
        _goml_runtime_core_string_get("", -1)
        return ""
    }
}

func main() {
    main0()
}
