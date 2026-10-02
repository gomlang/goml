package main

import (
    _goml_os "os"
)

func _goml_runtime_core_string_len(s string) int {
    return int(len(s))
}

func _goml_runtime_core_string_byte_get(s string, i int) uint8 {
    return s[i]
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

type _goml_vec_string struct {
    items []string
}

func vec_get__Vec_6string(vec *_goml_vec_string, index int) string {
    return vec.items[index]
}

func vec_len__Vec_6string(vec *_goml_vec_string) int {
    return int(len(vec.items))
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
    var minimum__0_lhs int64 = -9223372036854775807
    var minimum__0_rhs int64 = 1
    var minimum__0 int64 = minimum__0_lhs - minimum__0_rhs
    var maximum__0 int64 = 9223372036854775807
    var unsigned__0 uint64 = 18446744073709551615
    var inline7 string = _goml_m_trait__impl_i_ToString_i_i64_i_to__string(minimum__0)
    _goml_runtime_core_string_println(inline7)
    var inline5 string = _goml_m_trait__impl_i_ToString_i_i64_i_to__string(maximum__0)
    _goml_runtime_core_string_println(inline5)
    var inline3 string = _goml_m_trait__impl_i_ToString_i_u64_i_to__string(unsigned__0)
    _goml_runtime_core_string_println(inline3)
    var t0 [3]string = [3]string{"go", "ml", "!"}
    var t1 *_goml_vec_string = func(values [3]string) *_goml_vec_string {
        var storage struct {
            vector _goml_vec_string
            values [3]string
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t0)
    var t2 string
    var inline2 string = __goml_builtin_string_concat(t1)
    t2 = inline2
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t2)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_i64_i_to__string(self__0 int64) string {
    var inline0 string = signed_decimal_string(self__0)
    return inline0
}

func _goml_m_trait__impl_i_ToString_i_u64_i_to__string(self__0 uint64) string {
    var inline0 string = decimal_string(self__0)
    return inline0
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func __goml_builtin_string_concat(values__0 *_goml_vec_string) string {
    var length__0 int = 0
    var value_index__0 int = 0
    Loop_loop0:
    for {
        var t9 int
        var inline5 int = vec_len__Vec_6string(values__0)
        t9 = inline5
        var t10 bool = value_index__0 < t9
        if t10 {
            var compound_old2 int = length__0
            var t11 string = vec_get__Vec_6string(values__0, value_index__0)
            var compound_value2 int
            var inline4 int = _goml_runtime_core_string_len(t11)
            compound_value2 = inline4
            var t12 int = compound_old2 + compound_value2
            length__0 = t12
            var compound_old3 int = value_index__0
            var compound_value3 int = 1
            var t14 int = compound_old3 + compound_value3
            value_index__0 = t14
            continue
        } else {
            break Loop_loop0
        }
    }
    var bytes__0 *_goml_vec_uint8 = vec_with_capacity__Vec_5uint8(length__0)
    value_index__0 = 0
    Loop_loop1:
    for {
        var t0 int
        var inline3 int = vec_len__Vec_6string(values__0)
        t0 = inline3
        var t1 bool = value_index__0 < t0
        if t1 {
            var value__0 string = vec_get__Vec_6string(values__0, value_index__0)
            var byte_index__0 int = 0
            Loop_loop2:
            for {
                var t4 int
                var inline2 int = _goml_runtime_core_string_len(value__0)
                t4 = inline2
                var t5 bool = byte_index__0 < t4
                if t5 {
                    var t6 uint8
                    var inline1 uint8 = _goml_runtime_core_string_byte_get(value__0, byte_index__0)
                    t6 = inline1
                    vec_push__Vec_5uint8(bytes__0, t6)
                    var compound_old1 int = byte_index__0
                    var compound_value1 int = 1
                    var t7 int = compound_old1 + compound_value1
                    byte_index__0 = t7
                    continue
                } else {
                    break Loop_loop2
                }
            }
            var compound_old0 int = value_index__0
            var compound_value0 int = 1
            var t2 int = compound_old0 + compound_value0
            value_index__0 = t2
            continue
        } else {
            break Loop_loop1
        }
    }
    var mtmp0 Tuple2_4bool_6string = _goml_runtime_core_string_from_utf8(bytes__0)
    var x0 string = mtmp0._1
    return x0
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
