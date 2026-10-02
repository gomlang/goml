package main

import (
    _goml_ffi_import_strings_hff6d73a8adbc12e9f5edf8b318fbd31b "strings"
    _goml_ffi_import_runtime_hbfd85f862b823c92c4bdefd661241cfa "runtime"
)

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

func _goml_ffi_strings_x00_ToUpper_x00__m__z_string_h76369dffc1fe27f425185fa360e330cd(arg0 string) string {
    return _goml_ffi_import_strings_hff6d73a8adbc12e9f5edf8b318fbd31b.ToUpper(arg0)
}

func _goml_ffi_strings_x00_Cut_x00__o__ng_c_bool_q__hbfb59822899b5650c9939dcbef222236(arg0 string, arg1 string) Tuple3_6string_6string_4bool {
    var result0 string
    var result1 string
    var result2 bool
    result0, result1, result2 = _goml_ffi_import_strings_hff6d73a8adbc12e9f5edf8b318fbd31b.Cut(arg0, arg1)
    return Tuple3_6string_6string_4bool{
        _0: result0,
        _1: result1,
        _2: result2,
    }
}

func _goml_ffi_runtime_x00_Gosched_x00__m__z__o__q__h57c0f10f30d747e7ca95481c1b6b72c3() struct{} {
    _goml_ffi_import_runtime_hbfd85f862b823c92c4bdefd661241cfa.Gosched()
    return struct{}{}
}

type _goml_vec_uint32 struct {
    items []uint32
}

type Tuple3_6string_6string_4bool struct {
    _0 string
    _1 string
    _2 bool
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
    var t0 string = _goml_ffi_strings_x00_ToUpper_x00__m__z_string_h76369dffc1fe27f425185fa360e330cd("goml")
    var inline6 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t0)
    _goml_runtime_core_string_println(inline6)
    var mtmp0 Tuple3_6string_6string_4bool = _goml_ffi_strings_x00_Cut_x00__o__ng_c_bool_q__hbfb59822899b5650c9939dcbef222236("left:right", ":")
    var x0 string = mtmp0._0
    var x1 string = mtmp0._1
    var x2 bool = mtmp0._2
    var inline4 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(x0)
    _goml_runtime_core_string_println(inline4)
    var inline2 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(x1)
    _goml_runtime_core_string_println(inline2)
    var inline0 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(x2)
    _goml_runtime_core_string_println(inline0)
    _goml_ffi_runtime_x00_Gosched_x00__m__z__o__q__h57c0f10f30d747e7ca95481c1b6b72c3()
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func _goml_m_trait__impl_i_ToString_i_bool_i_to__string(self__0 bool) string {
    var t0 string = _goml_runtime_core_bool_to_string(self__0)
    return t0
}

func main() {
    main0()
}
