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

type DefaultValue struct {
    name string
}

type OverrideValue struct {}

type Ordering uint8

type dyn__Describe_vtable struct {
    name func(any) string
    describe func(any) string
}

type dyn__Describe struct {
    data any
    vtable *dyn__Describe_vtable
}

func dyn__Describe__wrap__DefaultValue__name(self any) string {
    return _goml_m_trait__impl_i_Describe_i_DefaultValue_i_name(self.(DefaultValue))
}

func dyn__Describe__wrap__DefaultValue__describe(self any) string {
    return _goml_m_trait__impl_i_Describe_i_DefaultValue_i_describe(self.(DefaultValue))
}

func dyn__Describe__vtable__DefaultValue() *dyn__Describe_vtable {
    return &dyn__Describe_vtable{
        name: dyn__Describe__wrap__DefaultValue__name,
        describe: dyn__Describe__wrap__DefaultValue__describe,
    }
}

func _goml_m_trait__impl_i_Describe_i_DefaultValue_i_name(self__0 DefaultValue) string {
    var t0 string = self__0.name
    return t0
}

func _goml_m_trait__impl_i_Describe_i_DefaultValue_i_describe(default_arg0 DefaultValue) string {
    var inline0 string = _goml_m_trait__impl_i_Describe_i_DefaultValue_i_name(default_arg0)
    var inline1 string = "default:" + inline0
    return inline1
}

func main0() struct{} {
    var default_value__0 DefaultValue = DefaultValue{
        name: "value",
    }
    var t0 string
    var inline10 string = _goml_m_trait__impl_i_Describe_i_DefaultValue_i_name(default_value__0)
    var inline11 string = "default:" + inline10
    t0 = inline11
    var inline8 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t0)
    _goml_runtime_core_string_println(inline8)
    var t1 string
    var inline7 string = "generic"
    t1 = inline7
    var inline5 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t1)
    _goml_runtime_core_string_println(inline5)
    var t2 string
    t2 = "override"
    var inline3 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t2)
    _goml_runtime_core_string_println(inline3)
    var t3 dyn__Describe = dyn__Describe{
        data: default_value__0,
        vtable: dyn__Describe__vtable__DefaultValue(),
    }
    var t4 string
    var inline2 string = t3.vtable.describe(t3.data)
    t4 = inline2
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t4)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func main() {
    main0()
}
