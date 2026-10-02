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

func ref_get__Ref_3int(reference *ref_int_x) int {
    return reference.value
}

func ref_set__Ref_3int(reference *ref_int_x, value int) struct{} {
    reference.value = value
    return struct{}{}
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

type Option__isize struct {
    _p0 int
    _tag uint8
}

func find(wanted__0 int) Option__isize {
    var current__0 *ref_int_x
    var inline2 int = 0
    var inline3 *ref_int_x = ref__Ref_3int(inline2)
    current__0 = inline3
    for {
        var value__0 int
        var inline1 int = ref_get__Ref_3int(current__0)
        value__0 = inline1
        var t0 bool = value__0 >= 5
        if t0 {
            return Option__isize{
                _tag: 0,
            }
        } else {
            var t1 bool = value__0 == wanted__0
            if t1 {
                var t2 Option__isize = Option__isize{
                    _p0: value__0,
                    _tag: 1,
                }
                return t2
            } else {
                var t3 int = value__0 + 1
                ref_set__Ref_3int(current__0, t3)
                continue
            }
        }
    }
}

func main0() struct{} {
    var t0 bool
    var inline6 int = 3
    var inline7 Option__isize = find(inline6)
    switch inline7._tag {
    case 0:
        t0 = false
    case 1:
        t0 = true
    default:
        panic("non-exhaustive match")
    }
    var inline4 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t0)
    _goml_runtime_core_string_println(inline4)
    var t1 bool
    var inline2 int = 8
    var inline3 Option__isize = find(inline2)
    switch inline3._tag {
    case 0:
        t1 = false
    case 1:
        t1 = true
    default:
        panic("non-exhaustive match")
    }
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
