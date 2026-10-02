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

func array_get__Array_6_5int32(arr [6]int32, index int) int32 {
    return arr[index]
}

func array_get__Array_3_5int32(arr [3]int32, index int) int32 {
    return arr[index]
}

func array_set__Array_3_5int32(arr [3]int32, index int, value int32) [3]int32 {
    arr[index] = value
    return arr
}

func array_get__Array_3_5uint8(arr [3]uint8, index int) uint8 {
    return arr[index]
}

func array_get__Array_2_5int32(arr [2]int32, index int) int32 {
    return arr[index]
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

type Token struct {}

type _goml_m_Buffer_____d_const__value_k_3 struct {
    data [3]int32
}

type Ordering uint8

type _goml_m_Storage____i32_____d_const__value_k_2 struct {
    _p0 [2]int32
    _tag uint8
}

func main0() struct{} {
    var values__0 [6]int32 = [6]int32{1, 2, 3, 4, 5, 6}
    var t0 int32 = _goml_m_first_____d_const__param_k_N___d_const__value_k_6(values__0)
    println__T_i32(t0)
    var t1 uint
    t1 = 6
    println__T_usize(t1)
    var t2 [3]int32 = [3]int32{7, 8, 9}
    var t3 int32
    var inline27 int32 = array_get__Array_3_5int32(t2, 0)
    t3 = inline27
    var inline25 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(t3)
    _goml_runtime_core_string_println(inline25)
    var t4 uint
    t4 = 3
    var inline23 string = _goml_m_trait__impl_i_ToString_i_usize_i_to__string(t4)
    _goml_runtime_core_string_println(inline23)
    var t5 uint
    t5 = 3
    var inline21 string = _goml_m_trait__impl_i_ToString_i_usize_i_to__string(t5)
    _goml_runtime_core_string_println(inline21)
    var t6 uint
    t6 = 6
    var inline19 string = _goml_m_trait__impl_i_ToString_i_usize_i_to__string(t6)
    _goml_runtime_core_string_println(inline19)
    var t7 [3]int32 = [3]int32{9, 8, 7}
    var t8 int32
    var inline18 int32 = array_get__Array_3_5int32(t7, 0)
    t8 = inline18
    var inline16 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(t8)
    _goml_runtime_core_string_println(inline16)
    var t9 [3]int32
    var inline8 [3]int32 = t2
    var inline9 [3]int32 = inline8
    var inline10 int = 0
    var inline11 int32 = array_get__Array_3_5int32(inline9, inline10)
    var inline12 int32 = 1
    var inline13 int32 = inline11 + inline12
    var inline14 [3]int32 = array_set__Array_3_5int32(inline9, inline10, inline13)
    inline8 = inline14
    t9 = inline8
    var t10 int32 = _goml_m_total_____d_const__param_k_N___d_const__value_k_3(t9)
    var inline6 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(t10)
    _goml_runtime_core_string_println(inline6)
    var t11 [3]uint8 = [3]uint8{97, 98, 99}
    var bytes__0 [3]uint8
    bytes__0 = t11
    var t12 uint8 = array_get__Array_3_5uint8(bytes__0, 1)
    var inline4 string = _goml_m_trait__impl_i_ToString_i_u8_i_to__string(t12)
    _goml_runtime_core_string_println(inline4)
    var t13 [2]int32 = [2]int32{1, 2}
    var t15 int32 = array_get__Array_2_5int32(t13, 1)
    var inline2 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(t15)
    _goml_runtime_core_string_println(inline2)
    var t14 uint
    t14 = 3
    var inline0 string = _goml_m_trait__impl_i_ToString_i_usize_i_to__string(t14)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func println__T_i32(value__0 int32) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_int32_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_first_____d_const__param_k_N___d_const__value_k_6(values__0 [6]int32) int32 {
    var t0 int32 = array_get__Array_6_5int32(values__0, 0)
    return t0
}

func println__T_usize(value__0 uint) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_uint_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_total_____d_const__param_k_N___d_const__value_k_3(values__0 [3]int32) int32 {
    var result__0 int32 = 0
    var for_limit0_source uint = 3
    var for_limit0 int = int(uint(for_limit0_source))
    var for_index0 int = 0
    Loop_loop0:
    for {
        var t0 bool = for_index0 < for_limit0
        if t0 {
            var for_item0 int32 = array_get__Array_3_5int32(values__0, for_index0)
            var t1 int = for_index0 + 1
            for_index0 = t1
            var compound_old0 int32 = result__0
            var t2 int32 = compound_old0 + for_item0
            result__0 = t2
            continue
        } else {
            break Loop_loop0
        }
    }
    return result__0
}

func _goml_m_trait__impl_i_ToString_i_i32_i_to__string(self__0 int32) string {
    var inline0 int64 = int64(int32(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func _goml_m_trait__impl_i_ToString_i_usize_i_to__string(self__0 uint) string {
    var inline0 uint64 = uint64(uint(self__0))
    var inline1 string = decimal_string(inline0)
    return inline1
}

func _goml_m_trait__impl_i_ToString_i_u8_i_to__string(self__0 uint8) string {
    var inline0 uint64 = uint64(uint8(self__0))
    var inline1 string = decimal_string(inline0)
    return inline1
}

func __goml_builtin_int32_to_string(value__0 int32) string {
    var t0 int64 = int64(int32(value__0))
    var inline0 bool = t0 < 0
    if inline0 {
        var inline1 uint64 = uint64(int64(t0))
        var inline2 uint64 = 0 - inline1
        var inline3 string = decimal_string(inline2)
        var inline4 string = "-" + inline3
        return inline4
    } else {
        var inline5 uint64 = uint64(int64(t0))
        var inline6 string = decimal_string(inline5)
        return inline6
    }
}

func __goml_builtin_uint_to_string(value__0 uint) string {
    var t0 uint64 = uint64(uint(value__0))
    var t1 string = decimal_string(t0)
    return t1
}

func signed_decimal_string(value__0 int64) string {
    var t0 bool = value__0 < 0
    if t0 {
        var t1 uint64 = uint64(int64(value__0))
        var t2 uint64 = 0 - t1
        var t3 string = decimal_string(t2)
        var t4 string = "-" + t3
        return t4
    } else {
        var t5 uint64 = uint64(int64(value__0))
        var t6 string = decimal_string(t5)
        return t6
    }
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
