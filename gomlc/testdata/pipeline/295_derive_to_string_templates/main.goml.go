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

type Opaque struct {}

type Counted struct {
    calls *ref_int_x
    text string
}

type Ordering uint8

type Status interface {
    isStatus()
}

type Ready struct {}

func (_ Ready) isStatus() {}

type Retry struct {
    _0 int
}

func (_ Retry) isStatus() {}

type Failed struct {
    _0 int
    _1 string
}

func (_ Failed) isStatus() {}

type Unknown struct {}

func (_ Unknown) isStatus() {}

type DefaultTuple struct {
    _0 int
}

func (_ DefaultTuple) isStatus() {}

type DefaultNamed struct {
    _0 string
}

func (_ DefaultNamed) isStatus() {}

type Empty struct {}

func (_ Empty) isStatus() {}

type Raw struct {}

func (_ Raw) isStatus() {}

type Repeated interface {
    isRepeated()
}

type Pair struct {
    _0 Counted
    _1 Counted
}

func (_ Pair) isRepeated() {}

type Ignored struct {
    _0 Counted
}

func (_ Ignored) isRepeated() {}

type Hidden__Opaque struct {
    _p0 string
    _tag uint8
}

type Mixed__isize struct {
    _p0 int
    _tag uint8
}

func _goml_m_trait__impl_i_ToString_i_Status_i_to__string(self__0 Status) string {
    switch self__0.(type) {
    case Ready:
        var t0 string = "" + "ready"
        return t0
    case Retry:
        var x0 int = self__0.(Retry)._0
        var _goml_m__i_derive1__formatted____0 string
        var inline0 string = __goml_builtin_int_to_string(x0)
        _goml_m__i_derive1__formatted____0 = inline0
        var t1 string = "" + "retry after "
        var t2 string = t1 + _goml_m__i_derive1__formatted____0
        var t3 string = t2 + "s"
        return t3
    case Failed:
        var x1 string = self__0.(Failed)._1
        var t4 string = "" + "错误 "
        var t5 string = t4 + "{"
        var t6 string = t5 + x1
        var t7 string = t6 + "}"
        return t7
    case Unknown:
        return "Status::Unknown"
    case DefaultTuple:
        var x2 int = self__0.(DefaultTuple)._0
        var t8 string
        var inline1 string = __goml_builtin_int_to_string(x2)
        t8 = inline1
        var t9 string = "Status::DefaultTuple(" + t8
        var t10 string = t9 + ")"
        return t10
    case DefaultNamed:
        var x3 string = self__0.(DefaultNamed)._0
        var t11 string = "Status::DefaultNamed { " + "value: "
        var t12 string = t11 + x3
        var t13 string = t12 + " }"
        return t13
    case Empty:
        return ""
    case Raw:
        var t14 string = "" + "raw\\path "
        var t15 string = t14 + "{"
        var t16 string = t15 + "ok"
        var t17 string = t16 + "}"
        return t17
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_trait__impl_i_Debug_i_Status_i_debug(self__0 Status) string {
    switch self__0.(type) {
    case Ready:
        return "Status::Ready"
    case Retry:
        var x0 int = self__0.(Retry)._0
        var t0 string
        var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(x0)
        t0 = inline0
        var t1 string = "Status::Retry(" + t0
        var t2 string = t1 + ")"
        return t2
    case Failed:
        var x1 int = self__0.(Failed)._0
        var x2 string = self__0.(Failed)._1
        var t3 string = "Status::Failed { " + "code: "
        var t4 string
        var inline2 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(x1)
        t4 = inline2
        var t5 string = t3 + t4
        var t6 string = t5 + ", "
        var t7 string = t6 + "message: "
        var t8 string
        var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(x2)
        t8 = inline1
        var t9 string = t7 + t8
        var t10 string = t9 + " }"
        return t10
    case Unknown:
        return "Status::Unknown"
    case DefaultTuple:
        var x3 int = self__0.(DefaultTuple)._0
        var t11 string
        var inline3 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(x3)
        t11 = inline3
        var t12 string = "Status::DefaultTuple(" + t11
        var t13 string = t12 + ")"
        return t13
    case DefaultNamed:
        var x4 string = self__0.(DefaultNamed)._0
        var t14 string = "Status::DefaultNamed { " + "value: "
        var t15 string
        var inline4 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(x4)
        t15 = inline4
        var t16 string = t14 + t15
        var t17 string = t16 + " }"
        return t17
    case Empty:
        return "Status::Empty"
    case Raw:
        return "Status::Raw"
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_trait__impl_i_ToString_i_Repeated_i_to__string(self__0 Repeated) string {
    switch self__0.(type) {
    case Pair:
        var x0 Counted = self__0.(Pair)._0
        var x1 Counted = self__0.(Pair)._1
        var _goml_m__i_derive1__formatted____0 string
        var inline10 *ref_int_x = x0.calls
        var inline11 *ref_int_x = x0.calls
        var inline12 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(inline11)
        var inline13 int = inline12 + 1
        _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__isize(inline10, inline13)
        var inline15 string = x0.text
        var inline16 *ref_int_x = x0.calls
        var inline17 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(inline16)
        var inline18 string = _goml_m_inherent_i_isize_i_isize_i_to__string(inline17)
        var inline19 string = inline15 + inline18
        _goml_m__i_derive1__formatted____0 = inline19
        var _goml_m__i_derive3__formatted____0 string
        var inline0 *ref_int_x = x1.calls
        var inline1 *ref_int_x = x1.calls
        var inline2 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(inline1)
        var inline3 int = inline2 + 1
        _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__isize(inline0, inline3)
        var inline5 string = x1.text
        var inline6 *ref_int_x = x1.calls
        var inline7 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(inline6)
        var inline8 string = _goml_m_inherent_i_isize_i_isize_i_to__string(inline7)
        var inline9 string = inline5 + inline8
        _goml_m__i_derive3__formatted____0 = inline9
        var t0 string = "" + _goml_m__i_derive3__formatted____0
        var t1 string = t0 + ":"
        var t2 string = t1 + _goml_m__i_derive1__formatted____0
        var t3 string = t2 + ":"
        var t4 string = t3 + _goml_m__i_derive3__formatted____0
        return t4
    case Ignored:
        var t5 string = "" + "ignored"
        return t5
    default:
        panic("non-exhaustive match")
    }
}

func main0() struct{} {
    var ready__0 Status = Ready{}
    var retry__0 Status = Retry{
        _0: 3,
    }
    var failed__0 Status = Failed{
        _0: 500,
        _1: "超时",
    }
    println__T_Status(ready__0)
    println__T_Status(retry__0)
    var inline30 string = _goml_m_trait__impl_i_ToString_i_Status_i_to__string(failed__0)
    _goml_runtime_core_string_println(inline30)
    var t0 string = _goml_m_trait__impl_i_Debug_i_Status_i_debug(ready__0)
    var inline28 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t0)
    _goml_runtime_core_string_println(inline28)
    var unknown__0 Status = Unknown{}
    var tuple__0 Status = DefaultTuple{
        _0: 7,
    }
    var named__0 Status = DefaultNamed{
        _0: "default",
    }
    var empty__0 Status = Empty{}
    var raw__0 Status = Raw{}
    var inline26 string = _goml_m_trait__impl_i_ToString_i_Status_i_to__string(unknown__0)
    _goml_runtime_core_string_println(inline26)
    var inline24 string = _goml_m_trait__impl_i_ToString_i_Status_i_to__string(tuple__0)
    _goml_runtime_core_string_println(inline24)
    var inline22 string = _goml_m_trait__impl_i_ToString_i_Status_i_to__string(named__0)
    _goml_runtime_core_string_println(inline22)
    var t1 string = _goml_m_trait__impl_i_ToString_i_Status_i_to__string(empty__0)
    var t2 bool = t1 == ""
    var inline20 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t2)
    _goml_runtime_core_string_println(inline20)
    var inline18 string = _goml_m_trait__impl_i_ToString_i_Status_i_to__string(raw__0)
    _goml_runtime_core_string_println(inline18)
    var hidden__0 Hidden__Opaque = Hidden__Opaque{
        _tag: 0,
    }
    var selected__0 Hidden__Opaque = Hidden__Opaque{
        _p0: "visible",
        _tag: 1,
    }
    var inline16 string = _goml_m_trait__impl_i_ToString_i_Hidden____Opaque_i_to__string(hidden__0)
    _goml_runtime_core_string_println(inline16)
    var inline14 string = _goml_m_trait__impl_i_ToString_i_Hidden____Opaque_i_to__string(selected__0)
    _goml_runtime_core_string_println(inline14)
    var mixed__0 Mixed__isize = Mixed__isize{
        _p0: 9,
        _tag: 1,
    }
    var inline12 string = _goml_m_trait__impl_i_ToString_i_Mixed____isize_i_to__string(mixed__0)
    _goml_runtime_core_string_println(inline12)
    var calls__0 *ref_int_x
    var inline10 int = 0
    var inline11 *ref_int_x = ref__Ref_3int(inline10)
    calls__0 = inline11
    var t5 Counted = Counted{
        calls: calls__0,
        text: "a",
    }
    var t6 Counted = Counted{
        calls: calls__0,
        text: "b",
    }
    var repeated__0 Repeated = Pair{
        _0: t5,
        _1: t6,
    }
    var inline8 string = _goml_m_trait__impl_i_ToString_i_Repeated_i_to__string(repeated__0)
    _goml_runtime_core_string_println(inline8)
    var t7 int
    var inline7 int = ref_get__Ref_3int(calls__0)
    t7 = inline7
    var inline5 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t7)
    _goml_runtime_core_string_println(inline5)
    var t8 Counted = Counted{
        calls: calls__0,
        text: "unused",
    }
    var ignored__0 Repeated = Ignored{
        _0: t8,
    }
    var inline3 string = _goml_m_trait__impl_i_ToString_i_Repeated_i_to__string(ignored__0)
    _goml_runtime_core_string_println(inline3)
    var t9 int
    var inline2 int = ref_get__Ref_3int(calls__0)
    t9 = inline2
    var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t9)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_isize_i_to__string(self__0 int) string {
    var inline0 int64 = int64(int(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(self__0 *ref_int_x) int {
    var t0 int = ref_get__Ref_3int(self__0)
    return t0
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__isize(self__0 *ref_int_x, value__0 int) struct{} {
    ref_set__Ref_3int(self__0, value__0)
    return struct{}{}
}

func _goml_m_inherent_i_isize_i_isize_i_to__string(self__0 int) string {
    var inline0 int64 = int64(int(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func println__T_Status(value__0 Status) struct{} {
    var t0 string = _goml_m_trait__impl_i_ToString_i_Status_i_to__string(value__0)
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

func _goml_m_trait__impl_i_ToString_i_bool_i_to__string(self__0 bool) string {
    var t0 string = _goml_runtime_core_bool_to_string(self__0)
    return t0
}

func _goml_m_trait__impl_i_ToString_i_Hidden____Opaque_i_to__string(self__0 Hidden__Opaque) string {
    switch self__0._tag {
    case 0:
        var t0 string = "" + "hidden"
        return t0
    case 1:
        var x0 string = self__0._p0
        var t1 string = "" + x0
        return t1
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_trait__impl_i_ToString_i_Mixed____isize_i_to__string(self__0 Mixed__isize) string {
    switch self__0._tag {
    case 0:
        var t0 string = "" + "hidden"
        return t0
    case 1:
        var x0 int = self__0._p0
        var t1 string
        var inline0 string = __goml_builtin_int_to_string(x0)
        t1 = inline0
        var t2 string = "Mixed::Shown(" + t1
        var t3 string = t2 + ")"
        return t3
    default:
        panic("non-exhaustive match")
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
    } else {}
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

func main() {
    main0()
}
