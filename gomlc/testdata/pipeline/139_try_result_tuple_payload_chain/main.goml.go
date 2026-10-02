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

type Tuple2_6string_6string struct {
    _0 string
    _1 string
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

type _goml_m_Result_____o_string_c_string_q_____string interface {
    is_goml_m_Result_____o_string_c_string_q_____string()
}

type _goml_m_Result_____o_string_c_string_q_____string_Ok struct {
    _0 Tuple2_6string_6string
}

func (_ _goml_m_Result_____o_string_c_string_q_____string_Ok) is_goml_m_Result_____o_string_c_string_q_____string() {}

type _goml_m_Result_____o_string_c_string_q_____string_Err struct {
    _0 string
}

func (_ _goml_m_Result_____o_string_c_string_q_____string_Err) is_goml_m_Result_____o_string_c_string_q_____string() {}

type Result__string__string struct {
    _p0 string
    _tag uint8
}

func split_host_port(ok__0 bool) _goml_m_Result_____o_string_c_string_q_____string {
    if ok__0 {
        var t0 Tuple2_6string_6string = Tuple2_6string_6string{
            _0: "localhost",
            _1: "8080",
        }
        var t1 _goml_m_Result_____o_string_c_string_q_____string = _goml_m_Result_____o_string_c_string_q_____string_Ok{
            _0: t0,
        }
        return t1
    } else {
        var t2 _goml_m_Result_____o_string_c_string_q_____string = _goml_m_Result_____o_string_c_string_q_____string_Err{
            _0: "missing port",
        }
        return t2
    }
}

func render(ok__0 bool) Result__string__string {
    var mtmp0 _goml_m_Result_____o_string_c_string_q_____string
    var inline0 _goml_m_Result_____o_string_c_string_q_____string = split_host_port(ok__0)
    mtmp0 = inline0
    var jp0 Tuple2_6string_6string
    switch mtmp0.(type) {
    case _goml_m_Result_____o_string_c_string_q_____string_Ok:
        var x2 Tuple2_6string_6string = mtmp0.(_goml_m_Result_____o_string_c_string_q_____string_Ok)._0
        jp0 = x2
        var x0 string = jp0._0
        var x1 string = jp0._1
        var t0 string = x0 + "="
        var t1 string = t0 + x1
        var t2 Result__string__string = Result__string__string{
            _p0: t1,
            _tag: 0,
        }
        return t2
    case _goml_m_Result_____o_string_c_string_q_____string_Err:
        var x3 string = mtmp0.(_goml_m_Result_____o_string_c_string_q_____string_Err)._0
        var t3 Result__string__string = Result__string__string{
            _p0: x3,
            _tag: 1,
        }
        return t3
    default:
        panic("non-exhaustive match")
    }
}

func main0() struct{} {
    var t0 Result__string__string = render(true)
    var t1 string
    switch t0._tag {
    case 0:
        var inline8 string = t0._p0
        var inline9 string = "ok " + inline8
        t1 = inline9
    case 1:
        var inline10 string = t0._p0
        var inline11 string = "err " + inline10
        t1 = inline11
    default:
        panic("non-exhaustive match")
    }
    var inline6 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t1)
    _goml_runtime_core_string_println(inline6)
    var t2 Result__string__string = render(false)
    var t3 string
    switch t2._tag {
    case 0:
        var inline2 string = t2._p0
        var inline3 string = "ok " + inline2
        t3 = inline3
    case 1:
        var inline4 string = t2._p0
        var inline5 string = "err " + inline4
        t3 = inline5
    default:
        panic("non-exhaustive match")
    }
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t3)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func main() {
    main0()
}
