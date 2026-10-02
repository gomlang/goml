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

type Boxed struct {
    value int32
}

type Ordering uint8

func _goml_m_trait__impl_i_Display_i_i32_i_show(self__0 int32) string {
    var inline0 string = __goml_builtin_int32_to_string(self__0)
    return inline0
}

func _goml_m_trait__impl_i_Debug_i_i32_i_show(self__0 int32) string {
    var t0 string
    var inline0 string = __goml_builtin_int32_to_string(self__0)
    t0 = inline0
    var t1 string = "i32(" + t0
    var t2 string = t1 + ")"
    return t2
}

func _goml_m_trait__impl_i_MyHash_i_i32_i_hash(self__0 int32) int32 {
    var t0 int32 = self__0 * 16777619
    var t1 int32 = t0 + 216613626
    return t1
}

func _goml_m_trait__impl_i_Add_i_i32_i_add(self__0 int32, other__0 int32) int32 {
    var t0 int32 = self__0 + other__0
    return t0
}

func _goml_m_trait__impl_i_Inspect_i_i32_i_inspect(self__0 int32) string {
    var t0 string
    var inline0 string = __goml_builtin_int32_to_string(self__0)
    t0 = inline0
    var t1 string = "<" + t0
    var t2 string = t1 + ">"
    return t2
}

func _goml_m_trait__impl_i_Display_i_Boxed_i_show(self__0 Boxed) string {
    var t0 int32 = self__0.value
    var t1 string
    var inline0 string = __goml_builtin_int32_to_string(t0)
    t1 = inline0
    var t2 string = "Boxed(" + t1
    var t3 string = t2 + ")"
    return t3
}

func _goml_m_trait__impl_i_Debug_i_Boxed_i_show(self__0 Boxed) string {
    var t0 int32 = self__0.value
    var t1 string
    var inline0 string = __goml_builtin_int32_to_string(t0)
    t1 = inline0
    var t2 string = "Boxed{value=" + t1
    var t3 string = t2 + "}"
    return t3
}

func _goml_m_trait__impl_i_MyHash_i_Boxed_i_hash(self__0 Boxed) int32 {
    var t0 int32 = self__0.value
    var t1 int32 = t0 * 31
    var t2 int32 = t1 + 7
    var t3 int32 = t2 * 1315423911
    return t3
}

func _goml_m_trait__impl_i_Add_i_Boxed_i_add(self__0 Boxed, other__0 Boxed) Boxed {
    var t0 int32 = self__0.value
    var t1 int32 = other__0.value
    var t2 int32 = t0 + t1
    var t3 Boxed = Boxed{
        value: t2,
    }
    return t3
}

func _goml_m_trait__impl_i_Inspect_i_Boxed_i_inspect(self__0 Boxed) string {
    var t0 int32 = self__0.value
    var t1 string
    var inline0 string = __goml_builtin_int32_to_string(t0)
    t1 = inline0
    var t2 string = "[" + t1
    var t3 string = t2 + "]"
    return t3
}

func main0() struct{} {
    var tag__0 int32 = 7
    var left__0 int32 = 10
    var right__0 int32 = 32
    var sum_tag__0 int32 = 0
    var first__0 int32 = 1
    var second__0 int32 = 2
    var third__0 int32 = 3
    var t0 string
    var inline30 int32 = combine_scaled__T_i32(left__0, right__0, 2)
    var inline31 string = report_pair__Q_i32__T_i32(tag__0, left__0, right__0, inline30)
    t0 = inline31
    var inline28 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t0)
    _goml_runtime_core_string_println(inline28)
    var t1 Boxed = Boxed{
        value: 99,
    }
    var t2 Boxed = Boxed{
        value: 3,
    }
    var t3 Boxed = Boxed{
        value: 4,
    }
    var t4 string
    var inline26 Boxed = combine_scaled__T_Boxed(t2, t3, 2)
    var inline27 string = report_pair__Q_Boxed__T_Boxed(t1, t2, t3, inline26)
    t4 = inline27
    var inline24 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t4)
    _goml_runtime_core_string_println(inline24)
    var t5 string
    var inline14 int32 = _goml_m_trait__impl_i_Add_i_i32_i_add(first__0, second__0)
    var inline15 int32 = _goml_m_trait__impl_i_Add_i_i32_i_add(inline14, third__0)
    var inline16 string = tag_text__Q_i32(sum_tag__0)
    var inline17 int32 = _goml_m_trait__impl_i_MyHash_i_i32_i_hash(inline15)
    var inline18 string = inline16 + " "
    var inline19 string = _goml_m_trait__impl_i_Inspect_i_i32_i_inspect(inline15)
    var inline20 string = inline18 + inline19
    var inline21 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline17)
    var inline22 string = " @" + inline21
    var inline23 string = inline20 + inline22
    t5 = inline23
    var inline12 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t5)
    _goml_runtime_core_string_println(inline12)
    var t6 Boxed = Boxed{
        value: 1,
    }
    var t7 Boxed = Boxed{
        value: 5,
    }
    var t8 Boxed = Boxed{
        value: 6,
    }
    var t9 Boxed = Boxed{
        value: 7,
    }
    var t10 string
    var inline2 Boxed = _goml_m_trait__impl_i_Add_i_Boxed_i_add(t7, t8)
    var inline3 Boxed = _goml_m_trait__impl_i_Add_i_Boxed_i_add(inline2, t9)
    var inline4 string = tag_text__Q_Boxed(t6)
    var inline5 int32 = _goml_m_trait__impl_i_MyHash_i_Boxed_i_hash(inline3)
    var inline6 string = inline4 + " "
    var inline7 string = _goml_m_trait__impl_i_Inspect_i_Boxed_i_inspect(inline3)
    var inline8 string = inline6 + inline7
    var inline9 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline5)
    var inline10 string = " @" + inline9
    var inline11 string = inline8 + inline10
    t10 = inline11
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t10)
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

func combine_scaled__T_i32(a__0 int32, b__0 int32, factor__0 int32) int32 {
    var t0 int32
    var inline1 int32 = a__0 + b__0
    t0 = inline1
    var inline0 int32 = t0 * factor__0
    return inline0
}

func report_pair__Q_i32__T_i32(tag__0 int32, a__0 int32, b__0 int32, combined__0 int32) string {
    var same__0 bool
    var inline12 bool = a__0 == b__0
    same__0 = inline12
    var header__0 string
    var inline7 string = _goml_m_trait__impl_i_Debug_i_i32_i_show(tag__0)
    var inline8 string = inline7 + "#"
    var inline9 int32 = _goml_m_trait__impl_i_MyHash_i_i32_i_hash(tag__0)
    var inline10 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline9)
    var inline11 string = inline8 + inline10
    header__0 = inline11
    var repr__0 string
    var inline3 string = _goml_m_trait__impl_i_Debug_i_i32_i_show(combined__0)
    var inline4 string = inline3 + " / "
    var inline5 string = _goml_m_trait__impl_i_Display_i_i32_i_show(combined__0)
    var inline6 string = inline4 + inline5
    repr__0 = inline6
    var h__0 int32
    var inline1 int32 = combined__0 * 16777619
    var inline2 int32 = inline1 + 216613626
    h__0 = inline2
    var t0 string = header__0 + " "
    var t1 string = t0 + repr__0
    var t2 string
    if same__0 {
        t2 = "true"
    } else {
        t2 = "false"
    }
    var t3 string = " | eq=" + t2
    var t4 string
    var inline0 string = __goml_builtin_int32_to_string(h__0)
    t4 = inline0
    var t5 string = " | hash=" + t4
    var t6 string = t3 + t5
    var t7 string = t1 + t6
    return t7
}

func combine_scaled__T_Boxed(a__0 Boxed, b__0 Boxed, factor__0 int32) Boxed {
    var t0 Boxed
    var inline3 int32 = a__0.value
    var inline4 int32 = b__0.value
    var inline5 int32 = inline3 + inline4
    var inline6 Boxed = Boxed{
        value: inline5,
    }
    t0 = inline6
    var inline0 int32 = t0.value
    var inline1 int32 = inline0 * factor__0
    var inline2 Boxed = Boxed{
        value: inline1,
    }
    return inline2
}

func report_pair__Q_Boxed__T_Boxed(tag__0 Boxed, a__0 Boxed, b__0 Boxed, combined__0 Boxed) string {
    var same__0 bool
    var inline14 int32 = a__0.value
    var inline15 int32 = b__0.value
    var inline16 bool = inline14 == inline15
    same__0 = inline16
    var header__0 string
    var inline9 string = _goml_m_trait__impl_i_Debug_i_Boxed_i_show(tag__0)
    var inline10 string = inline9 + "#"
    var inline11 int32 = _goml_m_trait__impl_i_MyHash_i_Boxed_i_hash(tag__0)
    var inline12 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline11)
    var inline13 string = inline10 + inline12
    header__0 = inline13
    var repr__0 string
    var inline5 string = _goml_m_trait__impl_i_Debug_i_Boxed_i_show(combined__0)
    var inline6 string = inline5 + " / "
    var inline7 string = _goml_m_trait__impl_i_Display_i_Boxed_i_show(combined__0)
    var inline8 string = inline6 + inline7
    repr__0 = inline8
    var h__0 int32
    var inline1 int32 = combined__0.value
    var inline2 int32 = inline1 * 31
    var inline3 int32 = inline2 + 7
    var inline4 int32 = inline3 * 1315423911
    h__0 = inline4
    var t0 string = header__0 + " "
    var t1 string = t0 + repr__0
    var t2 string
    if same__0 {
        t2 = "true"
    } else {
        t2 = "false"
    }
    var t3 string = " | eq=" + t2
    var t4 string
    var inline0 string = __goml_builtin_int32_to_string(h__0)
    t4 = inline0
    var t5 string = " | hash=" + t4
    var t6 string = t3 + t5
    var t7 string = t1 + t6
    return t7
}

func tag_text__Q_i32(tag__0 int32) string {
    var t0 string
    var inline3 string = _goml_m_inherent_i_i32_i_i32_i_to__string(tag__0)
    var inline4 string = "i32(" + inline3
    var inline5 string = inline4 + ")"
    t0 = inline5
    var t1 string = t0 + "#"
    var t2 int32
    var inline1 int32 = tag__0 * 16777619
    var inline2 int32 = inline1 + 216613626
    t2 = inline2
    var t3 string
    var inline0 string = __goml_builtin_int32_to_string(t2)
    t3 = inline0
    var t4 string = t1 + t3
    return t4
}

func tag_text__Q_Boxed(tag__0 Boxed) string {
    var t0 string
    var inline5 int32 = tag__0.value
    var inline6 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline5)
    var inline7 string = "Boxed{value=" + inline6
    var inline8 string = inline7 + "}"
    t0 = inline8
    var t1 string = t0 + "#"
    var t2 int32
    var inline1 int32 = tag__0.value
    var inline2 int32 = inline1 * 31
    var inline3 int32 = inline2 + 7
    var inline4 int32 = inline3 * 1315423911
    t2 = inline4
    var t3 string
    var inline0 string = __goml_builtin_int32_to_string(t2)
    t3 = inline0
    var t4 string = t1 + t3
    return t4
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
