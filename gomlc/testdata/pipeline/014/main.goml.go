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

type Tuple2_4bool_4bool struct {
    _0 bool
    _1 bool
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

func test_nested_match(x__0 Tuple2_4bool_4bool, y__0 Tuple2_4bool_4bool) struct{} {
    var x0 bool = x__0._0
    var x1 bool = x__0._1
    switch x1 {
    case true:
        var x2 bool = y__0._0
        var x3 bool = y__0._1
        switch x3 {
        case true:
            switch x2 {
            case true:
                var inline0 string = "case4"
                var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
                _goml_runtime_core_string_println(inline1)
                return struct{}{}
            case false:
                var inline3 string = "case3"
                var inline4 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline3)
                _goml_runtime_core_string_println(inline4)
                return struct{}{}
            default:
                panic("non-exhaustive match")
            }
        case false:
            var inline6 string = "case4"
            var inline7 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline6)
            _goml_runtime_core_string_println(inline7)
            return struct{}{}
        default:
            panic("non-exhaustive match")
        }
    case false:
        switch x0 {
        case true:
            var x4 bool = y__0._0
            var x5 bool = y__0._1
            switch x5 {
            case true:
                switch x4 {
                case true:
                    var inline9 string = "case2"
                    var inline10 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline9)
                    _goml_runtime_core_string_println(inline10)
                    return struct{}{}
                case false:
                    var inline12 string = "case1"
                    var inline13 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline12)
                    _goml_runtime_core_string_println(inline13)
                    return struct{}{}
                default:
                    panic("non-exhaustive match")
                }
            case false:
                var inline15 string = "case2"
                var inline16 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline15)
                _goml_runtime_core_string_println(inline16)
                return struct{}{}
            default:
                panic("non-exhaustive match")
            }
        case false:
            var x6 bool = y__0._0
            var x7 bool = y__0._1
            switch x7 {
            case true:
                switch x6 {
                case true:
                    var inline18 string = "case4"
                    var inline19 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline18)
                    _goml_runtime_core_string_println(inline19)
                    return struct{}{}
                case false:
                    var inline21 string = "case3"
                    var inline22 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline21)
                    _goml_runtime_core_string_println(inline22)
                    return struct{}{}
                default:
                    panic("non-exhaustive match")
                }
            case false:
                var inline24 string = "case4"
                var inline25 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline24)
                _goml_runtime_core_string_println(inline25)
                return struct{}{}
            default:
                panic("non-exhaustive match")
            }
        default:
            panic("non-exhaustive match")
        }
    default:
        panic("non-exhaustive match")
    }
}

func main0() struct{} {
    var t0 Tuple2_4bool_4bool = Tuple2_4bool_4bool{
        _0: true,
        _1: false,
    }
    var t1 Tuple2_4bool_4bool = Tuple2_4bool_4bool{
        _0: false,
        _1: true,
    }
    test_nested_match(t0, t1)
    var t2 Tuple2_4bool_4bool = Tuple2_4bool_4bool{
        _0: true,
        _1: false,
    }
    var t3 Tuple2_4bool_4bool = Tuple2_4bool_4bool{
        _0: true,
        _1: false,
    }
    test_nested_match(t2, t3)
    var t4 Tuple2_4bool_4bool = Tuple2_4bool_4bool{
        _0: false,
        _1: true,
    }
    var t5 Tuple2_4bool_4bool = Tuple2_4bool_4bool{
        _0: false,
        _1: true,
    }
    test_nested_match(t4, t5)
    var t6 Tuple2_4bool_4bool = Tuple2_4bool_4bool{
        _0: false,
        _1: true,
    }
    var t7 Tuple2_4bool_4bool = Tuple2_4bool_4bool{
        _0: true,
        _1: false,
    }
    test_nested_match(t6, t7)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func main() {
    main0()
}
