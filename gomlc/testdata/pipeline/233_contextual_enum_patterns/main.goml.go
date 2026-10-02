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

type _goml_vec_Boxed__isize struct {
    items []Boxed__isize
}

func vec_get__Vec_12Boxed__isize(vec *_goml_vec_Boxed__isize, index int) Boxed__isize {
    return vec.items[index]
}

func vec_len__Vec_12Boxed__isize(vec *_goml_vec_Boxed__isize) int {
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

type ref_Option__isize_x struct {
    value Option__isize
}

func ref__Ref_13Option__isize(value Option__isize) *ref_Option__isize_x {
    return &ref_Option__isize_x{
        value: value,
    }
}

func ref_get__Ref_13Option__isize(reference *ref_Option__isize_x) Option__isize {
    return reference.value
}

func ref_set__Ref_13Option__isize(reference *ref_Option__isize_x, value Option__isize) struct{} {
    reference.value = value
    return struct{}{}
}

type ref_int_x struct {
    value int
}

func ref__Ref_3int(value int) *ref_int_x {
    return &ref_int_x{
        value: value,
    }
}

func ref_get__Ref_3int(reference *ref_int_x) int {
    return reference.value
}

func ref_set__Ref_3int(reference *ref_int_x, value int) struct{} {
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

type Second struct {
    _p0 int
    _tag uint8
}

type First__isize interface {
    isFirst__isize()
}

type First__isize_Shared struct {
    _0 int
}

func (_ First__isize_Shared) isFirst__isize() {}

type Idle struct {}

func (_ Idle) isFirst__isize() {}

type Data struct {
    _0 int
    _1 string
}

func (_ Data) isFirst__isize() {}

type Result__isize__string struct {
    _p1 string
    _p0 int
    _tag uint8
}

type Option__Result__isize__string struct {
    _p0 Result__isize__string
    _tag uint8
}

type Option__isize struct {
    _p0 int
    _tag uint8
}

type Boxed__isize struct {
    _p0 int
    _tag uint8
}

func classify(value__0 First__isize) string {
    switch value__0.(type) {
    case First__isize_Shared:
        var x0 int = value__0.(First__isize_Shared)._0
        var t0 string
        var inline0 string = __goml_builtin_int_to_string(x0)
        t0 = inline0
        var t1 string = "shared:" + t0
        return t1
    case Idle:
        return "idle"
    case Data:
        var x1 int = value__0.(Data)._0
        var x2 string = value__0.(Data)._1
        var t2 string = x2 + ":"
        var t3 string
        var inline1 string = __goml_builtin_int_to_string(x1)
        t3 = inline1
        var t4 string = t2 + t3
        return t4
    default:
        panic("non-exhaustive match")
    }
}

func nested(value__0 Option__Result__isize__string) string {
    switch value__0._tag {
    case 0:
        return "none"
    case 1:
        var x0 Result__isize__string = value__0._p0
        switch x0._tag {
        case 0:
            var x1 int = x0._p0
            var t0 string
            var inline0 string = __goml_builtin_int_to_string(x1)
            t0 = inline0
            var t1 string = "ok:" + t0
            return t1
        case 1:
            var x2 string = x0._p1
            var t2 string = "err:" + x2
            return t2
        default:
            panic("non-exhaustive match")
        }
    default:
        panic("non-exhaustive match")
    }
}

func take_once(value__0 Option__isize) int {
    var current__0 *ref_Option__isize_x
    var inline6 *ref_Option__isize_x = ref__Ref_13Option__isize(value__0)
    current__0 = inline6
    var result__0 *ref_int_x
    var inline4 int = 0
    var inline5 *ref_int_x = ref__Ref_3int(inline4)
    result__0 = inline5
    Loop_loop0:
    for {
        var mtmp0 Option__isize
        var inline3 Option__isize = ref_get__Ref_13Option__isize(current__0)
        mtmp0 = inline3
        switch mtmp0._tag {
        case 1:
            var x0 int = mtmp0._p0
            ref_set__Ref_3int(result__0, x0)
            ref_set__Ref_13Option__isize(current__0, Option__isize{
                _tag: 0,
            })
            continue
        default:
            break Loop_loop0
        }
    }
    var inline0 int = ref_get__Ref_3int(result__0)
    return inline0
}

func sum_boxed(values__0 *_goml_vec_Boxed__isize) int {
    var result__0 *ref_int_x
    var inline3 int = 0
    var inline4 *ref_int_x = ref__Ref_3int(inline3)
    result__0 = inline4
    var for_limit0 int = vec_len__Vec_12Boxed__isize(values__0)
    var for_index0 int = 0
    Loop_loop0:
    for {
        var t0 bool = for_index0 < for_limit0
        if t0 {
            var for_item0 Boxed__isize = vec_get__Vec_12Boxed__isize(values__0, for_index0)
            var t1 int = for_index0 + 1
            for_index0 = t1
            switch for_item0._tag {
            case 0:
                var x0 int = for_item0._p0
                var t2 int
                var inline2 int = ref_get__Ref_3int(result__0)
                t2 = inline2
                var t3 int = t2 + x0
                ref_set__Ref_3int(result__0, t3)
                continue
            default:
                panic("non-exhaustive match")
            }
        } else {
            break Loop_loop0
        }
    }
    var inline0 int = ref_get__Ref_3int(result__0)
    return inline0
}

func main0() struct{} {
    var t0 Boxed__isize = Boxed__isize{
        _p0: 19,
        _tag: 0,
    }
    var t1 Boxed__isize = Boxed__isize{
        _p0: 23,
        _tag: 0,
    }
    var t2 [2]Boxed__isize = [2]Boxed__isize{t0, t1}
    var boxed__0 *_goml_vec_Boxed__isize = func(values [2]Boxed__isize) *_goml_vec_Boxed__isize {
        var storage struct {
            vector _goml_vec_Boxed__isize
            values [2]Boxed__isize
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t2)
    var t3 First__isize = First__isize_Shared{
        _0: 7,
    }
    var t4 string = classify(t3)
    println__T_string(t4)
    var t5 string = classify(Idle{})
    println__T_string(t5)
    var t6 First__isize = Data{
        _0: 9,
        _1: "data",
    }
    var t7 string = classify(t6)
    println__T_string(t7)
    var t8 Result__isize__string = Result__isize__string{
        _p0: 11,
        _tag: 0,
    }
    var t9 Option__Result__isize__string = Option__Result__isize__string{
        _p0: t8,
        _tag: 1,
    }
    var t10 string = nested(t9)
    println__T_string(t10)
    var t11 Result__isize__string = Result__isize__string{
        _p1: "bad",
        _tag: 1,
    }
    var t12 Option__Result__isize__string = Option__Result__isize__string{
        _p0: t11,
        _tag: 1,
    }
    var t13 string = nested(t12)
    println__T_string(t13)
    var t14 string = nested(Option__Result__isize__string{
        _tag: 0,
    })
    var inline16 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t14)
    _goml_runtime_core_string_println(inline16)
    var t15 int
    var inline15 int = 13
    t15 = inline15
    var inline13 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t15)
    _goml_runtime_core_string_println(inline13)
    var t16 int
    t16 = 0
    var inline11 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t16)
    _goml_runtime_core_string_println(inline11)
    var t17 bool
    t17 = true
    var inline9 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t17)
    _goml_runtime_core_string_println(inline9)
    var t18 bool
    t18 = false
    var inline7 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t18)
    _goml_runtime_core_string_println(inline7)
    var t19 Option__isize = Option__isize{
        _p0: 15,
        _tag: 1,
    }
    var t20 int = take_once(t19)
    var inline5 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t20)
    _goml_runtime_core_string_println(inline5)
    var t21 int
    var inline4 int = 17
    t21 = inline4
    var inline2 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t21)
    _goml_runtime_core_string_println(inline2)
    var t22 int = sum_boxed(boxed__0)
    var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t22)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func println__T_string(value__0 string) struct{} {
    var t0 string
    t0 = value__0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
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

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
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
