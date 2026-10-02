package main

import (
    _goml_os "os"
)

func _goml_runtime_core_string_from_utf8(bytes *_goml_vec_uint8) Tuple2_4bool_6string {
    return Tuple2_4bool_6string{
        _0: true,
        _1: string(bytes.items),
    }
}

func _goml_runtime_core_string_println(s string) struct{} {
    _goml_os.Stdout.WriteString(s + "\n")
    return struct{}{}
}

type _goml_vec_uint8 struct {
    items []uint8
}

func vec_with_capacity__Vec_5uint8(capacity int) *_goml_vec_uint8 {
    return &_goml_vec_uint8{
        items: make([]uint8, 0, capacity),
    }
}

func vec_push__Vec_5uint8(vec *_goml_vec_uint8, elem uint8) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_get__Vec_5uint8(vec *_goml_vec_uint8, index int) uint8 {
    return vec.items[index]
}

func vec_len__Vec_5uint8(vec *_goml_vec_uint8) int {
    return int(len(vec.items))
}

type _goml_vec_uint32 struct {
    items []uint32
}

type Tuple2_4bool_6string struct {
    _0 bool
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

func main0() struct{} {
    var start8__0 uint8 = 200
    var add8__0 uint8 = 55
    var sum8__0 uint8 = start8__0 + add8__0
    var neg8__0 uint8 = -start8__0
    var start16__0 uint16 = 50000
    var add16__0 uint16 = 12000
    var sum16__0 uint16 = start16__0 + add16__0
    var diff16__0 uint16 = sum16__0 - start16__0
    var add32__0 uint32 = 123456789
    var neg32__0 uint32 = -add32__0
    var start64__0 uint64 = 6000000000
    var add64__0 uint64 = 4000000000
    var sum64__0 uint64 = start64__0 + add64__0
    var diff64__0 uint64 = sum64__0 - add64__0
    var t0 string
    var inline6 string = __goml_builtin_uint8_to_string(sum8__0)
    t0 = inline6
    var t1 string = t0 + ", "
    var t2 string
    var inline5 string = __goml_builtin_uint8_to_string(neg8__0)
    t2 = inline5
    var t3 string = t1 + t2
    var t4 string = t3 + "; "
    var t5 string
    var inline4 string = __goml_builtin_uint16_to_string(diff16__0)
    t5 = inline4
    var t6 string = t4 + t5
    var t7 string = t6 + "; "
    var t8 string
    var inline3 string = __goml_builtin_uint32_to_string(neg32__0)
    t8 = inline3
    var t9 string = t7 + t8
    var t10 string = t9 + "; "
    var t11 string
    var inline2 string = __goml_builtin_uint64_to_string(diff64__0)
    t11 = inline2
    var message__0 string = t10 + t11
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(message__0)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func __goml_builtin_uint8_to_string(value__0 uint8) string {
    var t0 uint64 = uint64(uint8(value__0))
    var t1 string = decimal_string(t0)
    return t1
}

func __goml_builtin_uint16_to_string(value__0 uint16) string {
    var t0 uint64 = uint64(uint16(value__0))
    var t1 string = decimal_string(t0)
    return t1
}

func __goml_builtin_uint32_to_string(value__0 uint32) string {
    var t0 uint64 = uint64(uint32(value__0))
    var t1 string = decimal_string(t0)
    return t1
}

func __goml_builtin_uint64_to_string(value__0 uint64) string {
    var t0 string = decimal_string(value__0)
    return t0
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func decimal_string(value__0 uint64) string {
    var t0 bool = value__0 == 0
    if t0 {
        return "0"
    } else {
        var reversed__0 *_goml_vec_uint8 = vec_with_capacity__Vec_5uint8(20)
        var remaining__0 uint64 = value__0
        Loop_loop0:
        for {
            var t10 bool = remaining__0 > 0
            if t10 {
                var t11 uint64 = remaining__0 % 10
                var t12 uint8 = uint8(uint64(t11))
                var t13 uint8 = t12 + 48
                vec_push__Vec_5uint8(reversed__0, t13)
                var compound_old1 uint64 = remaining__0
                var compound_value1 uint64 = 10
                var t14 uint64 = compound_old1 / compound_value1
                remaining__0 = t14
                continue
            } else {
                break Loop_loop0
            }
        }
        var t1 int
        var inline3 int = vec_len__Vec_5uint8(reversed__0)
        t1 = inline3
        var bytes__0 *_goml_vec_uint8 = vec_with_capacity__Vec_5uint8(t1)
        var offset__0 int = 0
        Loop_loop1:
        for {
            var t2 int
            var inline2 int = vec_len__Vec_5uint8(reversed__0)
            t2 = inline2
            var t3 bool = offset__0 < t2
            if t3 {
                var t4 int
                var inline1 int = vec_len__Vec_5uint8(reversed__0)
                t4 = inline1
                var t5 int = t4 - offset__0
                var t6 int = t5 - 1
                var t7 uint8 = vec_get__Vec_5uint8(reversed__0, t6)
                vec_push__Vec_5uint8(bytes__0, t7)
                var compound_old0 int = offset__0
                var compound_value0 int = 1
                var t8 int = compound_old0 + compound_value0
                offset__0 = t8
                continue
            } else {
                break Loop_loop1
            }
        }
        var mtmp0 Tuple2_4bool_6string = _goml_runtime_core_string_from_utf8(bytes__0)
        var x0 string = mtmp0._1
        return x0
    }
}

func main() {
    main0()
}
