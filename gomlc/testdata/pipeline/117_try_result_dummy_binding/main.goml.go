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

type _goml_m_Result_____o__q_____string struct {
    _p0 string
    _tag uint8
}

type Result__string__string struct {
    _p0 string
    _tag uint8
}

func configure_and_format(config_ok__0 bool, read_ok__0 bool) Result__string__string {
    var mtmp0 _goml_m_Result_____o__q_____string
    if config_ok__0 {
        var inline3 _goml_m_Result_____o__q_____string = _goml_m_Result_____o__q_____string{
            _tag: 0,
        }
        mtmp0 = inline3
    } else {
        var inline4 _goml_m_Result_____o__q_____string = _goml_m_Result_____o__q_____string{
            _p0: "config failed",
            _tag: 1,
        }
        mtmp0 = inline4
    }
    switch mtmp0._tag {
    case 0:
        var mtmp1 Result__string__string
        if read_ok__0 {
            var inline1 Result__string__string = Result__string__string{
                _p0: "2s",
                _tag: 0,
            }
            mtmp1 = inline1
        } else {
            var inline2 Result__string__string = Result__string__string{
                _p0: "duration failed",
                _tag: 1,
            }
            mtmp1 = inline2
        }
        var jp0 string
        switch mtmp1._tag {
        case 0:
            var x0 string = mtmp1._p0
            jp0 = x0
            var t0 string
            var inline0 string = "duration=" + jp0
            t0 = inline0
            var t1 Result__string__string = Result__string__string{
                _p0: t0,
                _tag: 0,
            }
            return t1
        case 1:
            var x1 string = mtmp1._p0
            var t2 Result__string__string = Result__string__string{
                _p0: x1,
                _tag: 1,
            }
            return t2
        default:
            panic("non-exhaustive match")
        }
    case 1:
        var x2 string = mtmp0._p0
        var t3 Result__string__string = Result__string__string{
            _p0: x2,
            _tag: 1,
        }
        return t3
    default:
        panic("non-exhaustive match")
    }
}

func main0() struct{} {
    var t0 Result__string__string = configure_and_format(true, true)
    var t1 string
    switch t0._tag {
    case 0:
        var inline14 string = t0._p0
        var inline15 string = "ok " + inline14
        t1 = inline15
    case 1:
        var inline16 string = t0._p0
        var inline17 string = "err " + inline16
        t1 = inline17
    default:
        panic("non-exhaustive match")
    }
    var inline12 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t1)
    _goml_runtime_core_string_println(inline12)
    var t2 Result__string__string = configure_and_format(true, false)
    var t3 string
    switch t2._tag {
    case 0:
        var inline8 string = t2._p0
        var inline9 string = "ok " + inline8
        t3 = inline9
    case 1:
        var inline10 string = t2._p0
        var inline11 string = "err " + inline10
        t3 = inline11
    default:
        panic("non-exhaustive match")
    }
    var inline6 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t3)
    _goml_runtime_core_string_println(inline6)
    var t4 Result__string__string = configure_and_format(false, true)
    var t5 string
    switch t4._tag {
    case 0:
        var inline2 string = t4._p0
        var inline3 string = "ok " + inline2
        t5 = inline3
    case 1:
        var inline4 string = t4._p0
        var inline5 string = "err " + inline4
        t5 = inline5
    default:
        panic("non-exhaustive match")
    }
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t5)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func main() {
    main0()
}
