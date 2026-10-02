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

func array_get__Array_3_3int(arr [3]int, index int) int {
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

type Pair struct {
    left int
    right int
}

type Ordering uint8

type Choice struct {
    _p0 int
    _tag uint8
}

const (
    ANSWER int = 120
)

func factorial(value__0 int) int {
    var t0 bool = value__0 < 2
    if t0 {
        return 1
    } else {
        var t1 int = value__0 - 1
        var t2 int = factorial(t1)
        var t3 int = value__0 * t2
        return t3
    }
}

func main0() struct{} {
    var values__0 [3]int = [3]int{6, 10, 4}
    println__T_isize(ANSWER)
    var t0 int = array_get__Array_3_3int(values__0, 0)
    var inline30 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t0)
    _goml_runtime_core_string_println(inline30)
    var t1 int = array_get__Array_3_3int(values__0, 1)
    var inline28 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t1)
    _goml_runtime_core_string_println(inline28)
    var t2 int = array_get__Array_3_3int(values__0, 2)
    var inline26 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t2)
    _goml_runtime_core_string_println(inline26)
    var t3 int = 7
    var t4 int = 8
    var t5 int = t3 + t4
    var inline24 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t5)
    _goml_runtime_core_string_println(inline24)
    var x0 int = 9
    var inline22 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(x0)
    _goml_runtime_core_string_println(inline22)
    var t6 int = factorial(5)
    var t7 bool = t6 == ANSWER
    var inline20 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t7)
    _goml_runtime_core_string_println(inline20)
    var shadowed__0 int = 12
    var inline18 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(shadowed__0)
    _goml_runtime_core_string_println(inline18)
    var t8 int
    var inline14 int = 3
    var inline15 int = inline14 + 1
    var inline16 int = inline15 * 2
    var inline17 int = inline15 + inline16
    t8 = inline17
    var t9 bool = t8 == shadowed__0
    var inline12 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t9)
    _goml_runtime_core_string_println(inline12)
    var widened__0 uint64 = 18446744073709551615
    var t10 uint64
    var inline10 int8 = -1
    var inline11 uint64 = uint64(int8(inline10))
    t10 = inline11
    var t11 bool = widened__0 == t10
    var inline8 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t11)
    _goml_runtime_core_string_println(inline8)
    var folded_guard__0 int = 1
    var t12 int
    var inline6 int
    inline6 = 1
    t12 = inline6
    var inline4 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t12)
    _goml_runtime_core_string_println(inline4)
    var inline2 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(folded_guard__0)
    _goml_runtime_core_string_println(inline2)
    var t13 int = factorial(4)
    var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t13)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func println__T_isize(value__0 int) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_int_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_isize_i_to__string(self__0 int) string {
    var inline0 int64 = int64(int(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func _goml_m_trait__impl_i_ToString_i_bool_i_to__string(self__0 bool) string {
    var t0 string = _goml_runtime_core_bool_to_string(self__0)
    return t0
}

func __goml_builtin_int_to_string(value__0 int) string {
    var t0 int64 = int64(int(value__0))
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
