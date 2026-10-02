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

type _goml_vec_uint32 struct {
    items []uint32
}

type Tuple2_4int8_5int16 struct {
    _0 int8
    _1 int16
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

type PairData struct {
    head int32
    tail int64
}

type Ordering uint8

func is_special8(value__0 int8) bool {
    switch value__0 {
    case 5:
        return true
    case 7:
        return true
    default:
        return false
    }
}

func match_tuple(values__0 Tuple2_4int8_5int16) bool {
    var x0 int8 = values__0._0
    var x1 int16 = values__0._1
    switch x1 {
    case 2:
        switch x0 {
        case 1:
            return true
        default:
            return false
        }
    default:
        return false
    }
}

func match_struct(pair__0 PairData) bool {
    var x0 int32 = pair__0.head
    var x1 int64 = pair__0.tail
    switch x1 {
    case 200:
        switch x0 {
        case 100:
            return true
        default:
            return false
        }
    case 300:
        return true
    default:
        return false
    }
}

func main0() struct{} {
    var tuple_first__0 int8 = 1
    var tuple_second__0 int16 = 2
    var t0 Tuple2_4int8_5int16 = Tuple2_4int8_5int16{
        _0: tuple_first__0,
        _1: tuple_second__0,
    }
    var tuple_result_hit__0 bool = match_tuple(t0)
    var t1 Tuple2_4int8_5int16 = Tuple2_4int8_5int16{
        _0: 3,
        _1: 4,
    }
    var tuple_result_miss__0 bool = match_tuple(t1)
    var t2 PairData = PairData{
        head: 100,
        tail: 200,
    }
    var pair_first__0 bool = match_struct(t2)
    var t3 PairData = PairData{
        head: 10,
        tail: 300,
    }
    var pair_second__0 bool = match_struct(t3)
    var t4 bool = is_special8(5)
    var part1__0 string
    var inline30 string = "i8="
    var inline31 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t4)
    var inline32 string = inline30 + inline31
    part1__0 = inline32
    var t5 bool
    var inline29 int16 = 1024
    switch inline29 {
    case 1024:
        t5 = true
    case 2048:
        t5 = true
    default:
        t5 = false
    }
    var part2__0 string
    var inline26 string = ",i16="
    var inline27 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t5)
    var inline28 string = inline26 + inline27
    part2__0 = inline28
    var t6 bool
    var inline25 int32 = 8192
    switch inline25 {
    case 4096:
        t6 = true
    case 8192:
        t6 = true
    default:
        t6 = false
    }
    var part3__0 string
    var inline22 string = ",i32="
    var inline23 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t6)
    var inline24 string = inline22 + inline23
    part3__0 = inline24
    var t7 bool
    var inline21 int64 = 16384
    switch inline21 {
    case 16384:
        t7 = true
    case 32768:
        t7 = true
    default:
        t7 = false
    }
    var part4__0 string
    var inline18 string = ",int64_a="
    var inline19 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t7)
    var inline20 string = inline18 + inline19
    part4__0 = inline20
    var t8 bool
    var inline17 int64 = 32768
    switch inline17 {
    case 16384:
        t8 = true
    case 32768:
        t8 = true
    default:
        t8 = false
    }
    var part5__0 string
    var inline14 string = ",int64_b="
    var inline15 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t8)
    var inline16 string = inline14 + inline15
    part5__0 = inline16
    var part6__0 string
    var inline11 string = ",tuple_hit="
    var inline12 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(tuple_result_hit__0)
    var inline13 string = inline11 + inline12
    part6__0 = inline13
    var part7__0 string
    var inline8 string = ",tuple_miss="
    var inline9 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(tuple_result_miss__0)
    var inline10 string = inline8 + inline9
    part7__0 = inline10
    var part8__0 string
    var inline5 string = ",struct_first="
    var inline6 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(pair_first__0)
    var inline7 string = inline5 + inline6
    part8__0 = inline7
    var part9__0 string
    var inline2 string = ",struct_second="
    var inline3 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(pair_second__0)
    var inline4 string = inline2 + inline3
    part9__0 = inline4
    var t9 string = part1__0 + part2__0
    var t10 string = t9 + part3__0
    var t11 string = t10 + part4__0
    var t12 string = t11 + part5__0
    var t13 string = t12 + part6__0
    var t14 string = t13 + part7__0
    var t15 string = t14 + part8__0
    var message__0 string = t15 + part9__0
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(message__0)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_bool_i_to__string(self__0 bool) string {
    var t0 string = _goml_runtime_core_bool_to_string(self__0)
    return t0
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func main() {
    main0()
}
