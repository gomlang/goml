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

type ref_int32_x struct {
    value int32
}

func ref__Ref_5int32(value int32) *ref_int32_x {
    return &ref_int32_x{
        value: value,
    }
}

func ref_get__Ref_5int32(reference *ref_int32_x) int32 {
    return reference.value
}

func ref_set__Ref_5int32(reference *ref_int32_x, value int32) struct{} {
    reference.value = value
    return struct{}{}
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

type Option__i32 uint64

func accumulate(limit__0 int32) Option__i32 {
    var sum__0 *ref_int32_x
    var inline11 int32 = 0
    var inline12 *ref_int32_x = ref__Ref_5int32(inline11)
    sum__0 = inline12
    var i__0 *ref_int32_x
    var inline9 int32 = 0
    var inline10 *ref_int32_x = ref__Ref_5int32(inline9)
    i__0 = inline10
    Loop_loop0:
    for {
        var t2 int32
        var inline8 int32 = ref_get__Ref_5int32(i__0)
        t2 = inline8
        var t3 bool = t2 < limit__0
        if t3 {
            var cur__0 int32
            var inline7 int32 = ref_get__Ref_5int32(i__0)
            cur__0 = inline7
            var t4 int32 = cur__0 + 1
            ref_set__Ref_5int32(i__0, t4)
            var t5 bool = cur__0 == 1
            if t5 {
                continue
            } else {
                var mtmp0 Option__i32
                var inline3 bool = cur__0 == 2
                if inline3 {
                    mtmp0 = Option__i32(0)
                } else {
                    var inline4 int32 = cur__0 + 10
                    var inline5 Option__i32 = Option__i32(uint64(int64(inline4) + 2147483648) + 1)
                    mtmp0 = inline5
                }
                var jp0 int32
                switch mtmp0 != Option__i32(0) {
                case false:
                    return Option__i32(0)
                case true:
                    var x0 int32 = int32(int64(uint64(mtmp0) - 1) - 2147483648)
                    jp0 = x0
                    var t6 int32
                    var inline2 int32 = ref_get__Ref_5int32(sum__0)
                    t6 = inline2
                    var t7 int32 = t6 + jp0
                    ref_set__Ref_5int32(sum__0, t7)
                    continue
                default:
                    panic("non-exhaustive match")
                }
            }
        } else {
            break Loop_loop0
        }
    }
    var t0 int32
    var inline0 int32 = ref_get__Ref_5int32(sum__0)
    t0 = inline0
    var t1 Option__i32 = Option__i32(uint64(int64(t0) + 2147483648) + 1)
    return t1
}

func main0() struct{} {
    var t0 Option__i32 = accumulate(2)
    var t1 string
    switch t0 != Option__i32(0) {
    case false:
        t1 = "none"
    case true:
        var inline7 int32 = int32(int64(uint64(t0) - 1) - 2147483648)
        var inline8 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline7)
        var inline9 string = "some=" + inline8
        t1 = inline9
    default:
        panic("non-exhaustive match")
    }
    var inline5 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t1)
    _goml_runtime_core_string_println(inline5)
    var t2 Option__i32 = accumulate(4)
    var t3 string
    switch t2 != Option__i32(0) {
    case false:
        t3 = "none"
    case true:
        var inline2 int32 = int32(int64(uint64(t2) - 1) - 2147483648)
        var inline3 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline2)
        var inline4 string = "some=" + inline3
        t3 = inline4
    default:
        panic("non-exhaustive match")
    }
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t3)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_inherent_i_i32_i_i32_i_to__string(self__0 int32) string {
    var inline0 int64 = int64(int32(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
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
