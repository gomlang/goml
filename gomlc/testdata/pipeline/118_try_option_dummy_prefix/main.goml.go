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

type Option__string struct {
    _p0 string
    _tag uint8
}

func cut_prefix(case_id__0 int32) Option__string {
    var t0 bool = case_id__0 == 0
    if t0 {
        var t1 Option__string = Option__string{
            _p0: "ml",
            _tag: 1,
        }
        return t1
    } else {
        return Option__string{
            _tag: 0,
        }
    }
}

func main0() struct{} {
    var t0 Option__string
    var inline14 int32 = 0
    var inline15 Option__string = cut_prefix(inline14)
    var inline16 string
    switch inline15._tag {
    case 0:
        t0 = Option__string{
            _tag: 0,
        }
        var t1 string
        switch t0._tag {
        case 0:
            t1 = "none"
        case 1:
            var inline12 string = t0._p0
            var inline13 string = "some " + inline12
            t1 = inline13
        default:
            panic("non-exhaustive match")
        }
        var inline10 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t1)
        _goml_runtime_core_string_println(inline10)
        var t2 Option__string
        var inline4 int32 = 1
        var inline5 Option__string = cut_prefix(inline4)
        var inline6 string
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
            var inline9 string = inline5._p0
            inline6 = inline9
            var inline7 string = inline6 + "!"
            var inline8 Option__string = Option__string{
                _p0: inline7,
                _tag: 1,
            }
            t2 = inline8
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
        var inline19 string = inline15._p0
        inline16 = inline19
        var inline17 string = inline16 + "!"
        var inline18 Option__string = Option__string{
            _p0: inline17,
            _tag: 1,
        }
        t0 = inline18
        var t1 string
        switch t0._tag {
        case 0:
            t1 = "none"
        case 1:
            var inline12 string = t0._p0
            var inline13 string = "some " + inline12
            t1 = inline13
        default:
            panic("non-exhaustive match")
        }
        var inline10 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t1)
        _goml_runtime_core_string_println(inline10)
        var t2 Option__string
        var inline4 int32 = 1
        var inline5 Option__string = cut_prefix(inline4)
        var inline6 string
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
            var inline9 string = inline5._p0
            inline6 = inline9
            var inline7 string = inline6 + "!"
            var inline8 Option__string = Option__string{
                _p0: inline7,
                _tag: 1,
            }
            t2 = inline8
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
