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

type _goml_m_Option_____o_string_c_string_q_ struct {
    _p0 Tuple2_6string_6string
    _tag uint8
}

type Option__string struct {
    _p0 string
    _tag uint8
}

func cut_pair(ok__0 bool) _goml_m_Option_____o_string_c_string_q_ {
    if ok__0 {
        var t0 Tuple2_6string_6string = Tuple2_6string_6string{
            _0: "left",
            _1: "right",
        }
        var t1 _goml_m_Option_____o_string_c_string_q_ = _goml_m_Option_____o_string_c_string_q_{
            _p0: t0,
            _tag: 1,
        }
        return t1
    } else {
        return _goml_m_Option_____o_string_c_string_q_{
            _tag: 0,
        }
    }
}

func main0() struct{} {
    var t0 Option__string
    var inline13 bool = true
    var inline14 _goml_m_Option_____o_string_c_string_q_ = cut_pair(inline13)
    switch inline14._tag {
    case 0:
        t0 = Option__string{
            _tag: 0,
        }
        var t1 string
        switch t0._tag {
        case 0:
            t1 = "none"
        case 1:
            var inline11 string = t0._p0
            var inline12 string = "some " + inline11
            t1 = inline12
        default:
            panic("non-exhaustive match")
        }
        var inline9 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t1)
        _goml_runtime_core_string_println(inline9)
        var t2 Option__string
        var inline4 bool = false
        var inline5 _goml_m_Option_____o_string_c_string_q_ = cut_pair(inline4)
        switch inline5._tag {
        case 0:
            t2 = Option__string{
                _tag: 0,
            }
            var t3 string
            switch t2._tag {
            case 0:
                t3 = "none"
            case 1:
                var inline2 string = t2._p0
                var inline3 string = "some " + inline2
                t3 = inline3
            default:
                panic("non-exhaustive match")
            }
            var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t3)
            _goml_runtime_core_string_println(inline0)
            return struct{}{}
        case 1:
            var inline7 Option__string = Option__string{
                _p0: "ok",
                _tag: 1,
            }
            t2 = inline7
            var t3 string
            switch t2._tag {
            case 0:
                t3 = "none"
            case 1:
                var inline2 string = t2._p0
                var inline3 string = "some " + inline2
                t3 = inline3
            default:
                panic("non-exhaustive match")
            }
            var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t3)
            _goml_runtime_core_string_println(inline0)
            return struct{}{}
        default:
            panic("non-exhaustive match")
        }
    case 1:
        var inline16 Option__string = Option__string{
            _p0: "ok",
            _tag: 1,
        }
        t0 = inline16
        var t1 string
        switch t0._tag {
        case 0:
            t1 = "none"
        case 1:
            var inline11 string = t0._p0
            var inline12 string = "some " + inline11
            t1 = inline12
        default:
            panic("non-exhaustive match")
        }
        var inline9 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t1)
        _goml_runtime_core_string_println(inline9)
        var t2 Option__string
        var inline4 bool = false
        var inline5 _goml_m_Option_____o_string_c_string_q_ = cut_pair(inline4)
        switch inline5._tag {
        case 0:
            t2 = Option__string{
                _tag: 0,
            }
            var t3 string
            switch t2._tag {
            case 0:
                t3 = "none"
            case 1:
                var inline2 string = t2._p0
                var inline3 string = "some " + inline2
                t3 = inline3
            default:
                panic("non-exhaustive match")
            }
            var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t3)
            _goml_runtime_core_string_println(inline0)
            return struct{}{}
        case 1:
            var inline7 Option__string = Option__string{
                _p0: "ok",
                _tag: 1,
            }
            t2 = inline7
            var t3 string
            switch t2._tag {
            case 0:
                t3 = "none"
            case 1:
                var inline2 string = t2._p0
                var inline3 string = "some " + inline2
                t3 = inline3
            default:
                panic("non-exhaustive match")
            }
            var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t3)
            _goml_runtime_core_string_println(inline0)
            return struct{}{}
        default:
            panic("non-exhaustive match")
        }
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func main() {
    main0()
}
