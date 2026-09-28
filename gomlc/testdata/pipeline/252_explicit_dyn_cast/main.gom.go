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

type Number struct {
    value int
}

type Ordering uint8

type dyn__Display_vtable struct {
    display func(any) string
}

type dyn__Display struct {
    data any
    vtable *dyn__Display_vtable
}

type _goml_m_dyn_____d_dyn_k_dynSource_l_Item_x3d_isize_r___vtable struct {
    _goml_m_Source_i_get func(any) int
}

type _goml_m_dyn_____d_dyn_k_dynSource_l_Item_x3d_isize_r_ struct {
    data any
    vtable *_goml_m_dyn_____d_dyn_k_dynSource_l_Item_x3d_isize_r___vtable
}

func _goml_m_dyn_____d_dyn_k_dynSou_h6aa6e687a03d41d06284de45c4274b7f____Source_i_get(self any) int {
    return _goml_m_trait__impl_i_Source_i_Number_i_get(self.(Number))
}

func _goml_m_dyn_____d_dyn_k_dynSource_l_Item_x3d_isize_r_____vtable____Number() *_goml_m_dyn_____d_dyn_k_dynSource_l_Item_x3d_isize_r___vtable {
    return &_goml_m_dyn_____d_dyn_k_dynSource_l_Item_x3d_isize_r___vtable{
        _goml_m_Source_i_get: _goml_m_dyn_____d_dyn_k_dynSou_h6aa6e687a03d41d06284de45c4274b7f____Source_i_get,
    }
}

func dyn__Display__wrap__Number__display(self any) string {
    return _goml_m_trait__impl_i_Display_i_Number_i_display(self.(Number))
}

func dyn__Display__vtable__Number() *dyn__Display_vtable {
    return &dyn__Display_vtable{
        display: dyn__Display__wrap__Number__display,
    }
}

func _goml_m_trait__impl_i_Display_i_Number_i_display(self__0 Number) string {
    var t0 int = self__0.value
    var inline0 string = __goml_builtin_int_to_string(t0)
    return inline0
}

func _goml_m_trait__impl_i_Source_i_Number_i_get(self__0 Number) int {
    var t0 int = self__0.value
    return t0
}

func main0() struct{} {
    var t0 Number = Number{
        value: 42,
    }
    var display__0 dyn__Display = dyn__Display{
        data: t0,
        vtable: dyn__Display__vtable__Number(),
    }
    var t1 string = display__0.vtable.display(display__0.data)
    var inline7 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t1)
    _goml_runtime_core_string_println(inline7)
    var t2 Number = Number{
        value: 7,
    }
    var erased__0 dyn__Display
    var inline6 dyn__Display = dyn__Display{
        data: t2,
        vtable: dyn__Display__vtable__Number(),
    }
    erased__0 = inline6
    var t3 string = erased__0.vtable.display(erased__0.data)
    var inline4 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t3)
    _goml_runtime_core_string_println(inline4)
    var t4 Number = Number{
        value: 11,
    }
    var source__0 _goml_m_dyn_____d_dyn_k_dynSource_l_Item_x3d_isize_r_ = _goml_m_dyn_____d_dyn_k_dynSource_l_Item_x3d_isize_r_{
        data: t4,
        vtable: _goml_m_dyn_____d_dyn_k_dynSource_l_Item_x3d_isize_r_____vtable____Number(),
    }
    var t5 int = source__0.vtable._goml_m_Source_i_get(source__0.data)
    var inline2 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t5)
    _goml_runtime_core_string_println(inline2)
    var t6 Number = Number{
        value: 13,
    }
    var same__0 dyn__Display = dyn__Display{
        data: t6,
        vtable: dyn__Display__vtable__Number(),
    }
    var t7 string = same__0.vtable.display(same__0.data)
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t7)
    _goml_runtime_core_string_println(inline0)
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
