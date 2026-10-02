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

type Counter struct {
    start uint32
    end uint64
}

type Ordering uint8

func is_flag8(value__0 uint8) bool {
    switch value__0 {
    case 0:
        return true
    case 200:
        return true
    default:
        return false
    }
}

func is_flag16(value__0 uint16) bool {
    switch value__0 {
    case 1024:
        return true
    case 65000:
        return true
    default:
        return false
    }
}

func is_flag32(value__0 uint32) bool {
    switch value__0 {
    case 4000000000:
        return true
    case 1234567890:
        return true
    default:
        return false
    }
}

func report(label__0 string, value__0 bool) string {
    var t0 string
    var inline0 string = _goml_runtime_core_bool_to_string(value__0)
    t0 = inline0
    var t1 string = label__0 + t0
    return t1
}

func main0() struct{} {
    var t0 bool = is_flag8(200)
    var t1 string = report("u8_hit=", t0)
    var t2 bool = is_flag8(15)
    var t3 string = report(",u8_miss=", t2)
    var t4 string = t1 + t3
    var t5 bool = is_flag16(65000)
    var t6 string = report(",u16_hit=", t5)
    var t7 string = t4 + t6
    var t8 bool = is_flag16(42)
    var t9 string = report(",u16_miss=", t8)
    var t10 string = t7 + t9
    var t11 bool = is_flag32(1234567890)
    var t12 string = report(",u32_hit=", t11)
    var t13 string = t10 + t12
    var t14 bool
    var inline23 uint32 = 99
    switch inline23 {
    case 4000000000:
        t14 = true
    case 1234567890:
        t14 = true
    default:
        t14 = false
    }
    var t15 string
    var inline20 string = ",u32_miss="
    var inline21 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t14)
    var inline22 string = inline20 + inline21
    t15 = inline22
    var t16 string = t13 + t15
    var t17 bool
    var inline19 uint64 = 900000000
    switch inline19 {
    case 900000000:
        t17 = true
    case 600000000:
        t17 = true
    default:
        t17 = false
    }
    var t18 string
    var inline16 string = ",u64_hit="
    var inline17 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t17)
    var inline18 string = inline16 + inline17
    t18 = inline18
    var t19 string = t16 + t18
    var t20 bool
    var inline15 uint64 = 700000000
    switch inline15 {
    case 900000000:
        t20 = true
    case 600000000:
        t20 = true
    default:
        t20 = false
    }
    var t21 string
    var inline12 string = ",u64_miss="
    var inline13 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t20)
    var inline14 string = inline12 + inline13
    t21 = inline14
    var t22 string = t19 + t21
    var t23 bool
    var inline10 uint32 = 4000000000
    var inline11 uint64 = 900000000
    switch inline11 {
    case 900000000:
        switch inline10 {
        case 4000000000:
            t23 = true
        default:
            t23 = false
        }
    case 600000000:
        t23 = true
    default:
        t23 = false
    }
    var t24 string
    var inline7 string = ",struct_first="
    var inline8 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t23)
    var inline9 string = inline7 + inline8
    t24 = inline9
    var t25 string = t22 + t24
    var t26 bool
    var inline5 uint32 = 12
    var inline6 uint64 = 600000000
    switch inline6 {
    case 900000000:
        switch inline5 {
        case 4000000000:
            t26 = true
        default:
            t26 = false
        }
    case 600000000:
        t26 = true
    default:
        t26 = false
    }
    var t27 string
    var inline2 string = ",struct_second="
    var inline3 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t26)
    var inline4 string = inline2 + inline3
    t27 = inline4
    var message__0 string = t25 + t27
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(message__0)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_bool_i_to__string(self__0 bool) string {
    var t0 string = _goml_runtime_core_bool_to_string(self__0)
    return t0
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func main() {
    main0()
}
