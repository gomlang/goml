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

func _goml_intrinsic_missing(s string) struct{} {
    println("missing: " + s)
    panic("")
    return struct{}{}
}

func array_get__Array_4_3int(arr [4]int, index int) int {
    return arr[index]
}

func array_get__Array_2_3int(arr [2]int, index int) int {
    return arr[index]
}

type _goml_vec_int32 struct {
    items []int32
}

func vec_new__Vec_5int32() *_goml_vec_int32 {
    return &_goml_vec_int32{
        items: nil,
    }
}

func vec_push__Vec_5int32(vec *_goml_vec_int32, elem int32) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_get__Vec_5int32(vec *_goml_vec_int32, index int) int32 {
    return vec.items[index]
}

func vec_len__Vec_5int32(vec *_goml_vec_int32) int {
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

type ref_Maybe_x struct {
    value Maybe
}

func ref__Ref_5Maybe(value Maybe) *ref_Maybe_x {
    return &ref_Maybe_x{
        value: value,
    }
}

func ref_get__Ref_5Maybe(reference *ref_Maybe_x) Maybe {
    return reference.value
}

func ref_set__Ref_5Maybe(reference *ref_Maybe_x, value Maybe) struct{} {
    reference.value = value
    return struct{}{}
}

func missing__int32(s string) int32 {
    _goml_intrinsic_missing(s)
    var ret int32
    return ret
}

func missing__string(s string) string {
    _goml_intrinsic_missing(s)
    var ret string
    return ret
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
    left int32
    right int32
}

type Ordering uint8

type Maybe uint64

type Either struct {
    _p0 int32
    _tag uint8
}

func unwrap_either(value__0 Either) int32 {
    switch value__0._tag {
    case 0:
        var shared__0 int32 = value__0._p0
        var jp0 int32
        switch value__0._tag {
        case 0:
            jp0 = 0
        case 1:
            jp0 = 1
        default:
            panic("non-exhaustive match")
        }
        var t0 int32 = shared__0 + jp0
        return t0
    default:
        switch value__0._tag {
        case 1:
            var shared__0 int32 = value__0._p0
            var jp1 int32
            switch value__0._tag {
            case 0:
                jp1 = 0
            case 1:
                jp1 = 1
            default:
                panic("non-exhaustive match")
            }
            var t1 int32 = shared__0 + jp1
            return t1
        default:
            var t2 int32 = missing__int32("")
            return t2
        }
    }
}

func describe(value__0 Maybe, numbers__0 *_goml_vec_int32, view__0 []int32) string {
    var jp0 string
    switch value__0 != Maybe(0) {
    case true:
        var x0 int32 = int32(int64(uint64(value__0) - 1) - 2147483648)
        var t24 bool = x0 == 0
        if t24 {
            jp0 = "small"
        } else {
            var t25 bool = x0 == 1
            if t25 {
                jp0 = "small"
            } else {
                var t26 bool = x0 >= 2
                if t26 {
                    var t27 bool = x0 <= 4
                    if t27 {
                        jp0 = "middle"
                    } else {
                        var t28 bool = x0 > 10
                        if t28 {
                            jp0 = "large"
                        } else {
                            jp0 = "other"
                        }
                    }
                } else {
                    var t29 bool = x0 > 10
                    if t29 {
                        jp0 = "large"
                    } else {
                        jp0 = "other"
                    }
                }
            }
        }
    default:
        jp0 = "none"
    }
    var t0 int = vec_len__Vec_5int32(numbers__0)
    var t1 bool = t0 == 0
    var jp1 string
    if t1 {
        jp1 = "empty"
    } else {
        var t12 int = vec_len__Vec_5int32(numbers__0)
        var t13 bool = t12 >= 1
        if t13 {
            var first__1 int32 = vec_get__Vec_5int32(numbers__0, 0)
            var t14 int = vec_len__Vec_5int32(numbers__0)
            var tail__0 []int32 = numbers__0.items[1:t14]
            var t15 int
            var inline0 int = len(tail__0)
            t15 = inline0
            var t16 int32 = int32(int(t15))
            var t17 bool = first__1 == t16
            if t17 {
                jp1 = "balanced"
            } else {
                var t18 int = vec_len__Vec_5int32(numbers__0)
                var t19 bool = t18 >= 1
                if t19 {
                    jp1 = "nonempty"
                } else {
                    var t20 string = missing__string("")
                    jp1 = t20
                }
            }
        } else {
            var t21 int = vec_len__Vec_5int32(numbers__0)
            var t22 bool = t21 >= 1
            if t22 {
                jp1 = "nonempty"
            } else {
                var t23 string = missing__string("")
                jp1 = t23
            }
        }
    }
    var t2 int = len(view__0)
    var t3 bool = t2 >= 2
    var jp2 string
    if t3 {
        var first__0 int32 = view__0[0]
        var t8 int = len(view__0)
        var t9 int = t8 - 1
        var t10 int = t9 + 0
        var last__0 int32 = view__0[t10]
        var t11 bool = first__0 == last__0
        if t11 {
            jp2 = "same ends"
        } else {
            jp2 = "different ends"
        }
    } else {
        jp2 = "different ends"
    }
    var t4 string = jp0 + "/"
    var t5 string = t4 + jp1
    var t6 string = t5 + "/"
    var t7 string = t6 + jp2
    return t7
}

func main0() struct{} {
    var x0 int32 = 3
    var values__0 [4]int = [4]int{1, 2, 3, 1}
    var first__0 int = array_get__Array_4_3int(values__0, 0)
    var last__0 int = array_get__Array_4_3int(values__0, 3)
    var t0 int = array_get__Array_4_3int(values__0, 1)
    var t1 int = array_get__Array_4_3int(values__0, 2)
    var middle__0 [2]int = [2]int{t0, t1}
    var inline23 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(x0)
    _goml_runtime_core_string_println(inline23)
    var t2 int = array_get__Array_2_3int(middle__0, 0)
    var t3 int = first__0 + t2
    var t4 int = t3 + last__0
    var inline21 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t4)
    _goml_runtime_core_string_println(inline21)
    var numbers__0 *_goml_vec_int32 = vec_new__Vec_5int32()
    vec_push__Vec_5int32(numbers__0, 1)
    vec_push__Vec_5int32(numbers__0, 8)
    var t5 int
    var inline20 int = vec_len__Vec_5int32(numbers__0)
    t5 = inline20
    var view__0 []int32 = numbers__0.items[0:t5]
    var t6 Maybe = Maybe(uint64(int64(3) + 2147483648) + 1)
    var t7 string = describe(t6, numbers__0, view__0)
    var inline18 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t7)
    _goml_runtime_core_string_println(inline18)
    var empty__0 *_goml_vec_int32 = vec_new__Vec_5int32()
    var empty_view__0 []int32 = empty__0.items[0:0]
    var t8 string = describe(Maybe(0), empty__0, empty_view__0)
    var inline16 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t8)
    _goml_runtime_core_string_println(inline16)
    var t9 Maybe = Maybe(uint64(int64(7) + 2147483648) + 1)
    var state__0 *ref_Maybe_x
    var inline15 *ref_Maybe_x = ref__Ref_5Maybe(t9)
    state__0 = inline15
    Loop_loop0:
    for {
        var mtmp0 Maybe
        var inline14 Maybe = ref_get__Ref_5Maybe(state__0)
        mtmp0 = inline14
        switch mtmp0 != Maybe(0) {
        case true:
            var x3 int32 = int32(int64(uint64(mtmp0) - 1) - 2147483648)
            var inline12 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(x3)
            _goml_runtime_core_string_println(inline12)
            ref_set__Ref_5Maybe(state__0, Maybe(0))
            continue
        default:
            break Loop_loop0
        }
    }
    var x2 int32 = 6
    var inline9 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(x2)
    _goml_runtime_core_string_println(inline9)
    var jp0 int32
    var value__0 int32 = 5
    var jp1 int32
    var x1 int32 = 5
    jp1 = x1
    var t13 int32 = value__0 + jp1
    jp0 = t13
    var inline7 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(jp0)
    _goml_runtime_core_string_println(inline7)
    var t10 Either = Either{
        _p0: 11,
        _tag: 1,
    }
    var t11 int32 = unwrap_either(t10)
    var inline5 string = _goml_m_trait__impl_i_ToString_i_i32_i_to__string(t11)
    _goml_runtime_core_string_println(inline5)
    var t12 string
    var inline2 rune = 98
    var inline3 bool = inline2 >= 97
    if inline3 {
        var inline4 bool = inline2 <= 99
        if inline4 {
            t12 = "abc"
        } else {
            t12 = "other"
        }
    } else {
        t12 = "other"
    }
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t12)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_i32_i_to__string(self__0 int32) string {
    var inline0 int64 = int64(int32(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func _goml_m_trait__impl_i_ToString_i_isize_i_to__string(self__0 int) string {
    var inline0 int64 = int64(int(self__0))
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
