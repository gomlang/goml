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

type NoTraits struct {
    value int
}

type Wrapper__NoTraits struct {
    value NoTraits
}

type Generic__NoTraits__NoTraits struct {
    first Wrapper__NoTraits
    second Wrapper__NoTraits
}

type Ordering uint8

type GenericChoice__NoTraits__NoTraits struct {
    _p0 Wrapper__NoTraits
    _tag uint8
}

func main0() struct{} {
    var t0 NoTraits = NoTraits{
        value: 1,
    }
    var wrapped__0 Wrapper__NoTraits = Wrapper__NoTraits{
        value: t0,
    }
    var left__0 Generic__NoTraits__NoTraits = Generic__NoTraits__NoTraits{
        first: wrapped__0,
        second: wrapped__0,
    }
    var right__0 Generic__NoTraits__NoTraits = Generic__NoTraits__NoTraits{
        first: wrapped__0,
        second: wrapped__0,
    }
    var t1 string = _goml_m_trait__impl_i_ToString_i_Generic____NoTraits____NoTraits_i_to__string(left__0)
    println__T_string(t1)
    var t2 bool = _goml_m_trait__impl_i_PartialEq_i_Generic____NoTraits____NoTraits_i_eq(left__0, right__0)
    println__T_bool(t2)
    var t3 uint64 = _goml_m_trait__impl_i_Hash_i_Generic____NoTraits____NoTraits_i_hash(left__0)
    var t4 uint64 = _goml_m_trait__impl_i_Hash_i_Generic____NoTraits____NoTraits_i_hash(right__0)
    var t5 bool = t3 == t4
    var inline24 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t5)
    _goml_runtime_core_string_println(inline24)
    var empty__0 GenericChoice__NoTraits__NoTraits = GenericChoice__NoTraits__NoTraits{
        _tag: 0,
    }
    var value__0 GenericChoice__NoTraits__NoTraits = GenericChoice__NoTraits__NoTraits{
        _p0: wrapped__0,
        _tag: 1,
    }
    var t6 string = _goml_m_trait__impl_i_ToString_hfd40b94e3e10293076a83269859fcdb0_ts_i_to__string(empty__0)
    var inline22 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t6)
    _goml_runtime_core_string_println(inline22)
    var t7 string = _goml_m_trait__impl_i_ToString_hfd40b94e3e10293076a83269859fcdb0_ts_i_to__string(value__0)
    var inline20 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t7)
    _goml_runtime_core_string_println(inline20)
    var t8 bool = _goml_m_trait__impl_i_PartialEq_i_GenericChoice____NoTraits____NoTraits_i_eq(empty__0, value__0)
    var t9 bool = !t8
    var inline18 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t9)
    _goml_runtime_core_string_println(inline18)
    var t10 uint64
    var inline10_source int = 0
    var inline10 uint64 = uint64(int(inline10_source))
    var inline11 uint64 = inline10 + 14695981039346656037
    var inline12 uint64 = inline11 + 2
    var inline13_source int = 0
    var inline13 uint64 = uint64(int(inline13_source))
    var inline14 uint64 = inline13 + 1099511628211
    var inline15 uint64 = inline12 * inline14
    var inline16 uint64 = _goml_m_trait__impl_i_Hash_i_Wrapper____NoTraits_i_hash(wrapped__0)
    var inline17 uint64 = inline15 + inline16
    t10 = inline17
    var t11 uint64
    var inline2_source int = 0
    var inline2 uint64 = uint64(int(inline2_source))
    var inline3 uint64 = inline2 + 14695981039346656037
    var inline4 uint64 = inline3 + 2
    var inline5_source int = 0
    var inline5 uint64 = uint64(int(inline5_source))
    var inline6 uint64 = inline5 + 1099511628211
    var inline7 uint64 = inline4 * inline6
    var inline8 uint64 = _goml_m_trait__impl_i_Hash_i_Wrapper____NoTraits_i_hash(wrapped__0)
    var inline9 uint64 = inline7 + inline8
    t11 = inline9
    var t12 bool = t10 == t11
    var inline0 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t12)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func println__T_string(value__0 string) struct{} {
    var t0 string
    t0 = value__0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_Generic____NoTraits____NoTraits_i_to__string(self__0 Generic__NoTraits__NoTraits) string {
    var t0 string = "Generic { " + "first: "
    var t1 string
    t1 = "wrapped"
    var t2 string = t0 + t1
    var t3 string = t2 + ", "
    var t4 string = t3 + "second: "
    var t5 string
    t5 = "wrapped"
    var t6 string = t4 + t5
    var t7 string = t6 + " }"
    return t7
}

func println__T_bool(value__0 bool) struct{} {
    var t0 string
    var inline0 string = _goml_runtime_core_bool_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_trait__impl_i_PartialEq_i_Generic____NoTraits____NoTraits_i_eq(self__0 Generic__NoTraits__NoTraits, other__0 Generic__NoTraits__NoTraits) bool {
    var jp0 bool
    jp0 = true
    if jp0 {
        return true
    } else {
        return false
    }
}

func _goml_m_trait__impl_i_Hash_i_Generic____NoTraits____NoTraits_i_hash(self__0 Generic__NoTraits__NoTraits) uint64 {
    var t0_source int = 0
    var t0 uint64 = uint64(int(t0_source))
    var h__0 uint64 = t0 + 14695981039346656037
    var t1_source int = 0
    var t1 uint64 = uint64(int(t1_source))
    var t2 uint64 = t1 + 1099511628211
    var t3 uint64 = h__0 * t2
    var t4 uint64
    t4 = 7
    var h__1 uint64 = t3 + t4
    var t5_source int = 0
    var t5 uint64 = uint64(int(t5_source))
    var t6 uint64 = t5 + 1099511628211
    var t7 uint64 = h__1 * t6
    var t8 uint64
    t8 = 7
    var h__2 uint64 = t7 + t8
    return h__2
}

func _goml_m_trait__impl_i_ToString_hfd40b94e3e10293076a83269859fcdb0_ts_i_to__string(self__0 GenericChoice__NoTraits__NoTraits) string {
    switch self__0._tag {
    case 0:
        return "GenericChoice::Empty"
    case 1:
        var t0 string
        t0 = "wrapped"
        var t1 string = "GenericChoice::Value(" + t0
        var t2 string = t1 + ")"
        return t2
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_trait__impl_i_PartialEq_i_GenericChoice____NoTraits____NoTraits_i_eq(self__0 GenericChoice__NoTraits__NoTraits, other__0 GenericChoice__NoTraits__NoTraits) bool {
    switch other__0._tag {
    case 0:
        switch self__0._tag {
        case 0:
            return true
        default:
            return false
        }
    case 1:
        switch self__0._tag {
        case 1:
            return true
        default:
            return false
        }
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func _goml_m_trait__impl_i_ToString_i_bool_i_to__string(self__0 bool) string {
    var t0 string = _goml_runtime_core_bool_to_string(self__0)
    return t0
}

func _goml_m_trait__impl_i_Hash_i_Wrapper____NoTraits_i_hash(self__0 Wrapper__NoTraits) uint64 {
    return 7
}

func main() {
    main0()
}
