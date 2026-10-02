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

type Shape struct {}

type Ordering uint8

func _goml_m_inherent_i_Shape_i_Shape_i_name(self__0 Shape) string {
    return "Shape"
}

func _goml_m_inherent_i_Shape_i_Shape_i_rename(self__0 Shape, suffix__0 string) string {
    var t0 string
    t0 = "Shape"
    var t1 string = t0 + suffix__0
    return t1
}

func _goml_m_inherent_i_Shape_i_Shape_i_join(self__0 Shape, left__0 string, right__0 string) string {
    var t0 string
    t0 = "Shape"
    var t1 string = left__0 + t0
    var t2 string = t1 + right__0
    return t2
}

func main0() struct{} {
    var shape__0 Shape = Shape{}
    var inline0 string = _goml_m_inherent_i_Shape_i_Shape_i_name(shape__0)
    var inline1 string = _goml_m_inherent_i_Shape_i_Shape_i_rename(shape__0, "!")
    var inline2 string = _goml_m_inherent_i_Shape_i_Shape_i_join(shape__0, inline0, inline1)
    println__T_string(inline2)
    return struct{}{}
}

func println__T_string(value__0 string) struct{} {
    var t0 string
    t0 = value__0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func main() {
    main0()
}
