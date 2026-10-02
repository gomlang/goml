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

func classify(x__0 int32) string {
    var t0 bool = x__0 < 0
    if t0 {
        return "negative"
    } else {
        var t1 bool = 0 < x__0
        if t1 {
            return "positive"
        } else {
            return "zero"
        }
    }
}

func main0() struct{} {
    var first__0 string = classify(-42)
    var second__0 string = classify(0)
    var third__0 string = classify(17)
    var shape1__0 string
    var inline24 int32 = 1
    var inline25 int32 = 2
    var inline26 int32 = 3
    var inline27 bool = inline24 < inline25
    if inline27 {
        var inline28 bool = inline25 < inline26
        if inline28 {
            shape1__0 = "ascending"
        } else {
            shape1__0 = "peak"
        }
    } else {
        var inline29 bool = inline24 < inline26
        if inline29 {
            shape1__0 = "valley"
        } else {
            shape1__0 = "flat"
        }
    }
    var shape2__0 string
    var inline18 int32 = 3
    var inline19 int32 = 2
    var inline20 int32 = 1
    var inline21 bool = inline18 < inline19
    if inline21 {
        var inline22 bool = inline19 < inline20
        if inline22 {
            shape2__0 = "ascending"
        } else {
            shape2__0 = "peak"
        }
    } else {
        var inline23 bool = inline18 < inline20
        if inline23 {
            shape2__0 = "valley"
        } else {
            shape2__0 = "flat"
        }
    }
    var shape3__0 string
    var inline12 int32 = 2
    var inline13 int32 = 3
    var inline14 int32 = 2
    var inline15 bool = inline12 < inline13
    if inline15 {
        var inline16 bool = inline13 < inline14
        if inline16 {
            shape3__0 = "ascending"
        } else {
            shape3__0 = "peak"
        }
    } else {
        var inline17 bool = inline12 < inline14
        if inline17 {
            shape3__0 = "valley"
        } else {
            shape3__0 = "flat"
        }
    }
    var inline10 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(first__0)
    _goml_runtime_core_string_println(inline10)
    var inline8 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(second__0)
    _goml_runtime_core_string_println(inline8)
    var inline6 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(third__0)
    _goml_runtime_core_string_println(inline6)
    var inline4 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(shape1__0)
    _goml_runtime_core_string_println(inline4)
    var inline2 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(shape2__0)
    _goml_runtime_core_string_println(inline2)
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(shape3__0)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func main() {
    main0()
}
