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

type ref_int32_x struct {
    value int32
}

func ref__Ref_5int32(value int32) *ref_int32_x {
    return &ref_int32_x{
        value: value,
    }
}

func ref_get__Ref_5int32(reference *ref_int32_x) int32 {
    return reference.value
}

func ref_set__Ref_5int32(reference *ref_int32_x, value int32) struct{} {
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

type closure_env_main_0 struct {
    signal_0 *ref_int32_x
}

type Ordering uint8

func main0() struct{} {
    var signal__0 *ref_int32_x
    var inline4 int32 = 0
    var inline5 *ref_int32_x = ref__Ref_5int32(inline4)
    signal__0 = inline5
    var t0 closure_env_main_0 = closure_env_main_0{
        signal_0: signal__0,
    }
    var t1 func() struct{} = func() struct{} {
        return _goml_m_inherent_i_closure__env__main__0_i_closure__env__main__0_i_apply(t0)
    }
    go t1()
    Loop_loop0:
    for {
        var t2 int32
        var inline3 int32 = ref_get__Ref_5int32(signal__0)
        t2 = inline3
        var t3 bool = t2 < 1
        if t3 {
            continue
        } else {
            break Loop_loop0
        }
    }
    var inline0 string = "main"
    var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
    _goml_runtime_core_string_println(inline1)
    return struct{}{}
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__i32(self__0 *ref_int32_x, value__0 int32) struct{} {
    ref_set__Ref_5int32(self__0, value__0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func _goml_m_inherent_i_closure__env__main__0_i_closure__env__main__0_i_apply(env0 closure_env_main_0) struct{} {
    var signal__0 *ref_int32_x = env0.signal_0
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__i32(signal__0, 1)
    return struct{}{}
}

func main() {
    main0()
}
