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

func array_get__Array_2_15AlwaysDifferent(arr [2]AlwaysDifferent, index int) AlwaysDifferent {
    return arr[index]
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

type AlwaysDifferent struct {
    value int32
}

type Ordering uint8

func main0() struct{} {
    var t0 bool
    t0 = false
    var inline7 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t0)
    _goml_runtime_core_string_println(inline7)
    var t1 bool
    t1 = false
    var t2 bool = !t1
    var inline5 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t2)
    _goml_runtime_core_string_println(inline5)
    var t3 bool
    t3 = false
    var jp0 bool
    if t3 {
        var t13 int = 2
        var t14 int = 2
        var inline4 bool = t13 == t14
        jp0 = inline4
    } else {
        jp0 = false
    }
    var inline2 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(jp0)
    _goml_runtime_core_string_println(inline2)
    var t4 AlwaysDifferent = AlwaysDifferent{
        value: 1,
    }
    var t5 AlwaysDifferent = AlwaysDifferent{
        value: 2,
    }
    var left_array__0 [2]AlwaysDifferent = [2]AlwaysDifferent{t4, t5}
    var t6 AlwaysDifferent = AlwaysDifferent{
        value: 1,
    }
    var t7 AlwaysDifferent = AlwaysDifferent{
        value: 2,
    }
    var right_array__0 [2]AlwaysDifferent = [2]AlwaysDifferent{t6, t7}
    array_get__Array_2_15AlwaysDifferent(left_array__0, 0)
    array_get__Array_2_15AlwaysDifferent(right_array__0, 0)
    var t10 bool
    t10 = false
    var jp1 bool
    if t10 {
        array_get__Array_2_15AlwaysDifferent(left_array__0, 1)
        array_get__Array_2_15AlwaysDifferent(right_array__0, 1)
        jp1 = false
    } else {
        jp1 = false
    }
    var inline0 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(jp1)
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
