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

type _goml_vec_Dyn_Display struct {
    items []dyn__Display
}

func vec_new__Vec_11Dyn_Display() *_goml_vec_Dyn_Display {
    return &_goml_vec_Dyn_Display{
        items: nil,
    }
}

func vec_push__Vec_11Dyn_Display(vec *_goml_vec_Dyn_Display, elem dyn__Display) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_len__Vec_11Dyn_Display(vec *_goml_vec_Dyn_Display) int {
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

type Point struct {
    x int32
    y int32
}

type Flag struct {
    value bool
}

type Counter struct {
    cell *ref_int32_x
}

type closure_env_f_0 struct {}

type closure_env_make_renderer_1 struct {
    tag_0 string
}

type Ordering uint8

type dyn__Display_vtable struct {
    show func(any) string
    show_with func(any, string, string) string
    tick func(any) struct{}
    bump func(any, int32) int32
}

type dyn__Display struct {
    data any
    vtable *dyn__Display_vtable
}

func dyn__Display__wrap__Counter__show(self any) string {
    return _goml_m_trait__impl_i_Display_i_Counter_i_show(self.(Counter))
}

func dyn__Display__wrap__Counter__show_with(self any, p0 string, p1 string) string {
    return _goml_m_trait__impl_i_Display_i_Counter_i_show__with(self.(Counter), p0, p1)
}

func dyn__Display__wrap__Counter__tick(self any) struct{} {
    return _goml_m_trait__impl_i_Display_i_Counter_i_tick(self.(Counter))
}

func dyn__Display__wrap__Counter__bump(self any, p0 int32) int32 {
    return _goml_m_trait__impl_i_Display_i_Counter_i_bump(self.(Counter), p0)
}

func dyn__Display__vtable__Counter() *dyn__Display_vtable {
    return &dyn__Display_vtable{
        show: dyn__Display__wrap__Counter__show,
        show_with: dyn__Display__wrap__Counter__show_with,
        tick: dyn__Display__wrap__Counter__tick,
        bump: dyn__Display__wrap__Counter__bump,
    }
}

func dyn__Display__wrap__Flag__show(self any) string {
    return _goml_m_trait__impl_i_Display_i_Flag_i_show(self.(Flag))
}

func dyn__Display__wrap__Flag__show_with(self any, p0 string, p1 string) string {
    return _goml_m_trait__impl_i_Display_i_Flag_i_show__with(self.(Flag), p0, p1)
}

func dyn__Display__wrap__Flag__tick(self any) struct{} {
    return _goml_m_trait__impl_i_Display_i_Flag_i_tick(self.(Flag))
}

func dyn__Display__wrap__Flag__bump(self any, p0 int32) int32 {
    return _goml_m_trait__impl_i_Display_i_Flag_i_bump(self.(Flag), p0)
}

func dyn__Display__vtable__Flag() *dyn__Display_vtable {
    return &dyn__Display_vtable{
        show: dyn__Display__wrap__Flag__show,
        show_with: dyn__Display__wrap__Flag__show_with,
        tick: dyn__Display__wrap__Flag__tick,
        bump: dyn__Display__wrap__Flag__bump,
    }
}

func dyn__Display__wrap__Point__show(self any) string {
    return _goml_m_trait__impl_i_Display_i_Point_i_show(self.(Point))
}

func dyn__Display__wrap__Point__show_with(self any, p0 string, p1 string) string {
    return _goml_m_trait__impl_i_Display_i_Point_i_show__with(self.(Point), p0, p1)
}

func dyn__Display__wrap__Point__tick(self any) struct{} {
    return _goml_m_trait__impl_i_Display_i_Point_i_tick(self.(Point))
}

func dyn__Display__wrap__Point__bump(self any, p0 int32) int32 {
    return _goml_m_trait__impl_i_Display_i_Point_i_bump(self.(Point), p0)
}

func dyn__Display__vtable__Point() *dyn__Display_vtable {
    return &dyn__Display_vtable{
        show: dyn__Display__wrap__Point__show,
        show_with: dyn__Display__wrap__Point__show_with,
        tick: dyn__Display__wrap__Point__tick,
        bump: dyn__Display__wrap__Point__bump,
    }
}

func _goml_m_trait__impl_i_Display_i_Point_i_show(self__0 Point) string {
    var t0 int32 = self__0.x
    var t1 string
    var inline1 string = __goml_builtin_int32_to_string(t0)
    t1 = inline1
    var t2 string = "Point(" + t1
    var t3 string = t2 + ","
    var t4 int32 = self__0.y
    var t5 string
    var inline0 string = __goml_builtin_int32_to_string(t4)
    t5 = inline0
    var t6 string = t3 + t5
    var t7 string = t6 + ")"
    return t7
}

func _goml_m_trait__impl_i_Display_i_Point_i_show__with(self__0 Point, prefix__0 string, suffix__0 string) string {
    var t0 string = prefix__0 + "Point("
    var t1 int32 = self__0.x
    var t2 string
    var inline1 string = __goml_builtin_int32_to_string(t1)
    t2 = inline1
    var t3 string = t0 + t2
    var t4 string = t3 + ","
    var t5 int32 = self__0.y
    var t6 string
    var inline0 string = __goml_builtin_int32_to_string(t5)
    t6 = inline0
    var t7 string = t4 + t6
    var t8 string = t7 + ")"
    var t9 string = t8 + suffix__0
    return t9
}

func _goml_m_trait__impl_i_Display_i_Point_i_tick(self__0 Point) struct{} {
    return struct{}{}
}

func _goml_m_trait__impl_i_Display_i_Point_i_bump(self__0 Point, delta__0 int32) int32 {
    var t0 int32 = self__0.x
    var t1 int32 = self__0.y
    var t2 int32 = t0 + t1
    var t3 int32 = t2 + delta__0
    return t3
}

func _goml_m_trait__impl_i_Display_i_Flag_i_show(self__0 Flag) string {
    var t0 bool = self__0.value
    if t0 {
        return "Flag(true)"
    } else {
        return "Flag(false)"
    }
}

func _goml_m_trait__impl_i_Display_i_Flag_i_show__with(self__0 Flag, prefix__0 string, suffix__0 string) string {
    var t0 bool = self__0.value
    if t0 {
        var t1 string = prefix__0 + "Flag(true)"
        var t2 string = t1 + suffix__0
        return t2
    } else {
        var t3 string = prefix__0 + "Flag(false)"
        var t4 string = t3 + suffix__0
        return t4
    }
}

func _goml_m_trait__impl_i_Display_i_Flag_i_tick(self__0 Flag) struct{} {
    return struct{}{}
}

func _goml_m_trait__impl_i_Display_i_Flag_i_bump(self__0 Flag, delta__0 int32) int32 {
    var t0 bool = self__0.value
    if t0 {
        return delta__0
    } else {
        var t1 int32 = -delta__0
        return t1
    }
}

func _goml_m_trait__impl_i_Display_i_Counter_i_show(self__0 Counter) string {
    var t0 *ref_int32_x = self__0.cell
    var t1 int32
    var inline1 int32 = ref_get__Ref_5int32(t0)
    t1 = inline1
    var t2 string
    var inline0 string = __goml_builtin_int32_to_string(t1)
    t2 = inline0
    var t3 string = "Counter(" + t2
    var t4 string = t3 + ")"
    return t4
}

func _goml_m_trait__impl_i_Display_i_Counter_i_show__with(self__0 Counter, prefix__0 string, suffix__0 string) string {
    var t0 string = prefix__0 + "Counter("
    var t1 *ref_int32_x = self__0.cell
    var t2 int32
    var inline1 int32 = ref_get__Ref_5int32(t1)
    t2 = inline1
    var t3 string
    var inline0 string = __goml_builtin_int32_to_string(t2)
    t3 = inline0
    var t4 string = t0 + t3
    var t5 string = t4 + ")"
    var t6 string = t5 + suffix__0
    return t6
}

func _goml_m_trait__impl_i_Display_i_Counter_i_tick(self__0 Counter) struct{} {
    var t0 *ref_int32_x = self__0.cell
    var t1 int32
    var inline1 int32 = ref_get__Ref_5int32(t0)
    t1 = inline1
    var next__0 int32 = t1 + 1
    var t2 *ref_int32_x = self__0.cell
    ref_set__Ref_5int32(t2, next__0)
    return struct{}{}
}

func _goml_m_trait__impl_i_Display_i_Counter_i_bump(self__0 Counter, delta__0 int32) int32 {
    var t0 *ref_int32_x = self__0.cell
    var t1 int32
    var inline1 int32 = ref_get__Ref_5int32(t0)
    t1 = inline1
    var next__0 int32 = t1 + delta__0
    var t2 *ref_int32_x = self__0.cell
    ref_set__Ref_5int32(t2, next__0)
    return next__0
}

func show_dyn(x__0 dyn__Display) string {
    var t0 string = x__0.vtable.show_with(x__0.data, "<", ">")
    return t0
}

func call_via_closure(x__0 dyn__Display, tag__0 string) string {
    var t0 closure_env_f_0 = closure_env_f_0{}
    var f__0 func(dyn__Display, string) string = func(p0 dyn__Display, p1 string) string {
        return _goml_m_inherent_i_closure__env__f__0_i_closure__env__f__0_i_apply(t0, p0, p1)
    }
    var t1 string = f__0(x__0, tag__0)
    return t1
}

func make_renderer(tag__0 string) func(dyn__Display) string {
    var t0 closure_env_make_renderer_1 = closure_env_make_renderer_1{
        tag_0: tag__0,
    }
    var t1 func(dyn__Display) string = func(p0 dyn__Display) string {
        return _goml_m_inherent_i_closure__en_h5c3741356b48d7360bc79df27842b70e_erer__1_i_apply(t0, p0)
    }
    return t1
}

func main0() struct{} {
    var p1__0 Point = Point{
        x: 1,
        y: 2,
    }
    var p2__0 Point = Point{
        x: 3,
        y: 4,
    }
    var f1__0 Flag = Flag{
        value: true,
    }
    var f2__0 Flag = Flag{
        value: false,
    }
    var t0 *ref_int32_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__i32(10)
    var c__0 Counter = Counter{
        cell: t0,
    }
    var dp1__0 dyn__Display = dyn__Display{
        data: p1__0,
        vtable: dyn__Display__vtable__Point(),
    }
    var dp2__0 dyn__Display = dyn__Display{
        data: p2__0,
        vtable: dyn__Display__vtable__Point(),
    }
    var df1__0 dyn__Display = dyn__Display{
        data: f1__0,
        vtable: dyn__Display__vtable__Flag(),
    }
    var df2__0 dyn__Display = dyn__Display{
        data: f2__0,
        vtable: dyn__Display__vtable__Flag(),
    }
    var dc__0 dyn__Display = dyn__Display{
        data: c__0,
        vtable: dyn__Display__vtable__Counter(),
    }
    var render_star__0 func(dyn__Display) string = make_renderer("*")
    var render_angle__0 func(dyn__Display) string = make_renderer("<")
    var s0__0 string = show_dyn(dp2__0)
    var s1__0 string = call_via_closure(df2__0, "*")
    var t1 string = render_star__0(dp1__0)
    var t2 string = t1 + "|"
    var t3 string = render_angle__0(df1__0)
    var s2__0 string = t2 + t3
    var v__0 *_goml_vec_Dyn_Display
    var inline29 *_goml_vec_Dyn_Display = vec_new__Vec_11Dyn_Display()
    v__0 = inline29
    vec_push__Vec_11Dyn_Display(v__0, dp1__0)
    vec_push__Vec_11Dyn_Display(v__0, df1__0)
    vec_push__Vec_11Dyn_Display(v__0, dc__0)
    var vlen__0 int
    var inline25 int = vec_len__Vec_11Dyn_Display(v__0)
    vlen__0 = inline25
    var jp0 int32
    switch vlen__0 {
    case 2:
        jp0 = 3
    default:
        jp0 = 5
    }
    var inline23 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(s0__0)
    _goml_runtime_core_string_println(inline23)
    var inline21 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(s1__0)
    _goml_runtime_core_string_println(inline21)
    var inline19 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(s2__0)
    _goml_runtime_core_string_println(inline19)
    var i__0 *ref_int32_x
    var inline17 int32 = 0
    var inline18 *ref_int32_x = ref__Ref_5int32(inline17)
    i__0 = inline18
    Loop_loop0:
    for {
        var t8 int32
        var inline16 int32 = ref_get__Ref_5int32(i__0)
        t8 = inline16
        var t9 bool = t8 < 3
        if t9 {
            var line__0 string
            dc__0.vtable.tick(dc__0.data)
            var inline11 string = dc__0.vtable.show_with(dc__0.data, "[", "]")
            var inline12 string = inline11 + ":"
            var inline13 int32 = dc__0.vtable.bump(dc__0.data, jp0)
            var inline14 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline13)
            var inline15 string = inline12 + inline14
            line__0 = inline15
            var inline8 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(line__0)
            _goml_runtime_core_string_println(inline8)
            var t10 int32
            var inline7 int32 = ref_get__Ref_5int32(i__0)
            t10 = inline7
            var t11 int32 = t10 + 1
            ref_set__Ref_5int32(i__0, t11)
            continue
        } else {
            break Loop_loop0
        }
    }
    var t4 string
    var inline5 string = __goml_builtin_int_to_string(vlen__0)
    t4 = inline5
    var t5 string = "len:" + t4
    var inline3 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t5)
    _goml_runtime_core_string_println(inline3)
    var t6 string
    var inline2 string = __goml_builtin_int32_to_string(jp0)
    t6 = inline2
    var t7 string = "delta:" + t6
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t7)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_inherent_i_i32_i_i32_i_to__string(self__0 int32) string {
    var inline0 int64 = int64(int32(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__i32(value__0 int32) *ref_int32_x {
    var t0 *ref_int32_x = ref__Ref_5int32(value__0)
    return t0
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

func _goml_m_inherent_i_closure__env__f__0_i_closure__env__f__0_i_apply(env0 closure_env_f_0, v__0 dyn__Display, t__0 string) string {
    var t0 string = v__0.vtable.show_with(v__0.data, t__0, t__0)
    return t0
}

func _goml_m_inherent_i_closure__en_h5c3741356b48d7360bc79df27842b70e_erer__1_i_apply(env0 closure_env_make_renderer_1, x__0 dyn__Display) string {
    var tag__0 string = env0.tag_0
    var t0 string = x__0.vtable.show_with(x__0.data, tag__0, tag__0)
    return t0
}

func main() {
    main0()
}
