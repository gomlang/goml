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

type Choice struct {
    _p1 int32
    _p0 bool
    _tag uint8
}

type Result__i32__string struct {
    _p1 string
    _p0 int32
    _tag uint8
}

func choose(choice__0 Choice) Result__i32__string {
    var jp0 int32
    switch choice__0._tag {
    case 0:
        var x0 bool = choice__0._p0
        var commute_field0 int32
        var commute_field1 string
        if x0 {
            commute_field0 = 10
            jp0 = commute_field0
            var t0 Result__i32__string = Result__i32__string{
                _p0: jp0,
                _tag: 0,
            }
            return t0
        } else {
            commute_field1 = "left failed"
            var t1 Result__i32__string = Result__i32__string{
                _p1: commute_field1,
                _tag: 1,
            }
            return t1
        }
    case 1:
        var x1 bool = choice__0._p0
        var mtmp0 Result__i32__string
        if x1 {
            var inline0 Result__i32__string = Result__i32__string{
                _p0: 20,
                _tag: 0,
            }
            mtmp0 = inline0
        } else {
            var inline1 Result__i32__string = Result__i32__string{
                _p1: "right failed",
                _tag: 1,
            }
            mtmp0 = inline1
        }
        var jp1 int32
        switch mtmp0._tag {
        case 0:
            var x2 int32 = mtmp0._p0
            jp1 = x2
            var t2 int32 = jp1 + 1
            jp0 = t2
            var t0 Result__i32__string = Result__i32__string{
                _p0: jp0,
                _tag: 0,
            }
            return t0
        case 1:
            var x3 string = mtmp0._p1
            var t3 Result__i32__string = Result__i32__string{
                _p1: x3,
                _tag: 1,
            }
            return t3
        default:
            panic("non-exhaustive match")
        }
    case 2:
        var x4 int32 = choice__0._p1
        jp0 = x4
        var t0 Result__i32__string = Result__i32__string{
            _p0: jp0,
            _tag: 0,
        }
        return t0
    default:
        panic("non-exhaustive match")
    }
}

func show(res__0 Result__i32__string) string {
    switch res__0._tag {
    case 0:
        var x0 int32 = res__0._p0
        var t0 string
        var inline0 string = __goml_builtin_int32_to_string(x0)
        t0 = inline0
        var t1 string = "ok " + t0
        return t1
    case 1:
        var x1 string = res__0._p1
        var t2 string = "err " + x1
        return t2
    default:
        panic("non-exhaustive match")
    }
}

func main0() struct{} {
    var t0 Choice = Choice{
        _p0: true,
        _tag: 0,
    }
    var t1 Result__i32__string = choose(t0)
    var t2 string = show(t1)
    var inline23 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t2)
    _goml_runtime_core_string_println(inline23)
    var t3 Choice = Choice{
        _p0: true,
        _tag: 1,
    }
    var t4 Result__i32__string = choose(t3)
    var t5 string = show(t4)
    var inline21 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t5)
    _goml_runtime_core_string_println(inline21)
    var t6 Choice = Choice{
        _p1: 5,
        _tag: 2,
    }
    var t7 Result__i32__string = choose(t6)
    var t8 string
    switch t7._tag {
    case 0:
        var inline16 int32 = t7._p0
        var inline17 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline16)
        var inline18 string = "ok " + inline17
        t8 = inline18
    case 1:
        var inline19 string = t7._p1
        var inline20 string = "err " + inline19
        t8 = inline20
    default:
        panic("non-exhaustive match")
    }
    var inline14 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t8)
    _goml_runtime_core_string_println(inline14)
    var t9 Choice = Choice{
        _p0: false,
        _tag: 0,
    }
    var t10 Result__i32__string = choose(t9)
    var t11 string
    switch t10._tag {
    case 0:
        var inline9 int32 = t10._p0
        var inline10 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline9)
        var inline11 string = "ok " + inline10
        t11 = inline11
    case 1:
        var inline12 string = t10._p1
        var inline13 string = "err " + inline12
        t11 = inline13
    default:
        panic("non-exhaustive match")
    }
    var inline7 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t11)
    _goml_runtime_core_string_println(inline7)
    var t12 Choice = Choice{
        _p0: false,
        _tag: 1,
    }
    var t13 Result__i32__string = choose(t12)
    var t14 string
    switch t13._tag {
    case 0:
        var inline2 int32 = t13._p0
        var inline3 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline2)
        var inline4 string = "ok " + inline3
        t14 = inline4
    case 1:
        var inline5 string = t13._p1
        var inline6 string = "err " + inline5
        t14 = inline6
    default:
        panic("non-exhaustive match")
    }
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t14)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_inherent_i_i32_i_i32_i_to__string(self__0 int32) string {
    var inline0 int64 = int64(int32(self__0))
    var inline1 string = signed_decimal_string(inline0)
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
