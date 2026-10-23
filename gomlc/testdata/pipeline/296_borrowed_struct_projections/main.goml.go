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

func array_get__Array_16_3int(arr [16]int, index int) int {
    return arr[index]
}

type _goml_vec_Outer struct {
    items []Outer
}

func vec_get__Vec_5Outer(vec *_goml_vec_Outer, index int) Outer {
    return vec.items[index]
}

func vec_set__Vec_5Outer(vec *_goml_vec_Outer, index int, value Outer) struct{} {
    vec.items[index] = value
    return struct{}{}
}

type _goml_vec_Fn0__to_3int struct {
    items []func() int
}

func vec_new__Vec_12Fn0__to_3int() *_goml_vec_Fn0__to_3int {
    return &_goml_vec_Fn0__to_3int{
        items: nil,
    }
}

func vec_push__Vec_12Fn0__to_3int(vec *_goml_vec_Fn0__to_3int, elem func() int) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_get__Vec_12Fn0__to_3int(vec *_goml_vec_Fn0__to_3int, index int) func() int {
    return vec.items[index]
}

func vec_len__Vec_12Fn0__to_3int(vec *_goml_vec_Fn0__to_3int) int {
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

type ref_Outer_x struct {
    value Outer
}

func ref__Ref_5Outer(value Outer) *ref_Outer_x {
    return &ref_Outer_x{
        value: value,
    }
}

func ref_get__Ref_5Outer(reference *ref_Outer_x) Outer {
    return reference.value
}

func ref_set__Ref_5Outer(reference *ref_Outer_x, value Outer) struct{} {
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

type Large struct {
    values [16]int
    shared *ref_int_x
}

type Outer struct {
    inner Large
}

type closure_env_captured_0 struct {
    field_0 Large
}

type closure_env_main_1 struct {
    value_0 Large
}

type Ordering uint8

func _goml_m_make(value__0 int) Outer {
    var t0 [16]int = func(value int) [16]int {
        var result [16]int
        for index := range result {
            result[index] = value
        }
        return result
    }(value__0)
    var t1 *ref_int_x
    var inline0 *ref_int_x = ref__Ref_3int(value__0)
    t1 = inline0
    var t2 Large = Large{
        values: t0,
        shared: t1,
    }
    var t3 Outer = Outer{
        inner: t2,
    }
    return t3
}

func read(value__0 *Large) int {
    var t0 *[16]int = &(*value__0).values
    var t1 int = array_get__Array_16_3int(*t0, 0)
    var t2 *[16]int = &(*value__0).values
    var t3 int = array_get__Array_16_3int(*t2, 15)
    var t4 int = t1 + t3
    var t5 *ref_int_x = (*value__0).shared
    var t6 int
    var inline0 int = ref_get__Ref_3int(t5)
    t6 = inline0
    var t7 int = t4 + t6
    return t7
}

func replace(value__0 *Large, next__0 *Large) int {
    var local__0 Large
    local__0 = *next__0
    var inline0 [16]int = local__0.values
    var inline1 int = array_get__Array_16_3int(inline0, 0)
    var inline2 [16]int = local__0.values
    var inline3 int = array_get__Array_16_3int(inline2, 15)
    var inline4 int = inline1 + inline3
    var inline5 *ref_int_x = local__0.shared
    var inline6 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(inline5)
    var inline7 int = inline4 + inline6
    return inline7
}

func returned(value__0 *Outer) Large {
    var t0 *Large = &(*value__0).inner
    return *t0
}

func captured(value__0 *Outer) func() int {
    var field__0 *Large = &(*value__0).inner
    var t0 closure_env_captured_0 = closure_env_captured_0{
        field_0: *field__0,
    }
    var t1 func() int = func() int {
        return _goml_m_inherent_i_closure__env__captured__0_i_closure__env__captured__0_i_apply(t0)
    }
    return t1
}

func main0() struct{} {
    var first__0 Outer = _goml_m_make(1)
    var field__0 *Large = &first__0.inner
    var t0 *ref_int_x = (*field__0).shared
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__isize(t0, 3)
    var t1 *Large = &first__0.inner
    var t2 int = read(t1)
    println__T_isize(t2)
    var t3 Outer = _goml_m_make(2)
    var t4 *Large = &t3.inner
    var t5 int = replace(field__0, t4)
    println__T_isize(t5)
    var t6 int = read(field__0)
    println__T_isize(t6)
    var t7 Large = returned(&first__0)
    var t8 int = read(&t7)
    println__T_isize(t8)
    var t9 func() int = captured(&first__0)
    var t10 int = t9()
    println__T_isize(t10)
    var current__0 Outer = _goml_m_make(4)
    var snapshot__0 Large = current__0.inner
    var t11 Outer = _goml_m_make(5)
    current__0 = t11
    var t12 int = read(&snapshot__0)
    println__T_isize(t12)
    var t13 Large = current__0.inner
    var t14 int = read(&t13)
    println__T_isize(t14)
    var t15 Outer = _goml_m_make(6)
    var cell__0 *ref_Outer_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__Outer(&t15)
    var t16 Outer = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__Outer(cell__0)
    var from_ref__0 *Large = &t16.inner
    var t17 Outer = _goml_m_make(7)
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__Outer(cell__0, &t17)
    var t18 int = read(from_ref__0)
    println__T_isize(t18)
    var t19 Outer = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__Outer(cell__0)
    var t20 *Large = &t19.inner
    var t21 int = read(t20)
    println__T_isize(t21)
    var t22 Outer = _goml_m_make(8)
    var t23 Outer = _goml_m_make(9)
    var t24 [2]Outer = [2]Outer{t22, t23}
    var values__0 *_goml_vec_Outer = func(values [2]Outer) *_goml_vec_Outer {
        var storage struct {
            vector _goml_vec_Outer
            values [2]Outer
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t24)
    var t25 Outer = vec_get__Vec_5Outer(values__0, 0)
    var from_vector__0 *Large = &t25.inner
    var index0 int = 0
    vec_get__Vec_5Outer(values__0, index0)
    var value0 Outer = _goml_m_make(10)
    vec_set__Vec_5Outer(values__0, index0, value0)
    var t27 int = read(from_vector__0)
    println__T_isize(t27)
    var t28 Outer = vec_get__Vec_5Outer(values__0, 0)
    var t29 *Large = &t28.inner
    var t30 int = read(t29)
    println__T_isize(t30)
    var closures__0 *_goml_vec_Fn0__to_3int
    var inline37 *_goml_vec_Fn0__to_3int = vec_new__Vec_12Fn0__to_3int()
    closures__0 = inline37
    var for_index0 int = 0
    var for_limit0 int = 3
    Loop_loop0:
    for {
        var t44 bool = for_index0 < for_limit0
        if t44 {
            var for_item2 int = for_index0
            var t45 int = for_index0 + 1
            for_index0 = t45
            var current__1 Outer
            var inline33 [16]int = func(value int) [16]int {
                var result [16]int
                for index := range result {
                    result[index] = value
                }
                return result
            }(for_item2)
            var inline34 *ref_int_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__isize(for_item2)
            var inline35 Large = Large{
                values: inline33,
                shared: inline34,
            }
            var inline36 Outer = Outer{
                inner: inline35,
            }
            current__1 = inline36
            var value__0 Large = current__1.inner
            var t46 closure_env_main_1 = closure_env_main_1{
                value_0: value__0,
            }
            var t47 func() int = func() int {
                return _goml_m_inherent_i_closure__env__main__1_i_closure__env__main__1_i_apply(t46)
            }
            vec_push__Vec_12Fn0__to_3int(closures__0, t47)
            continue
        } else {
            break Loop_loop0
        }
    }
    var for_limit1 int = vec_len__Vec_12Fn0__to_3int(closures__0)
    var for_index1 int = 0
    Loop_loop1:
    for {
        var t41 bool = for_index1 < for_limit1
        if t41 {
            var for_item1 func() int = vec_get__Vec_12Fn0__to_3int(closures__0, for_index1)
            var t42 int = for_index1 + 1
            for_index1 = t42
            var t43 int = for_item1()
            var inline30 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t43)
            _goml_runtime_core_string_println(inline30)
            continue
        } else {
            break Loop_loop1
        }
    }
    var loop_value__0 Outer = _goml_m_make(11)
    var for_index2 int = 0
    var for_limit2 int = 3
    Loop_loop2:
    for {
        var t37 bool = for_index2 < for_limit2
        if t37 {
            var for_item0 int = for_index2
            var t38 int = for_index2 + 1
            for_index2 = t38
            var previous__0 Large = loop_value__0.inner
            var t39 Outer
            var inline26 [16]int = func(value int) [16]int {
                var result [16]int
                for index := range result {
                    result[index] = value
                }
                return result
            }(for_item0)
            var inline27 *ref_int_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__isize(for_item0)
            var inline28 Large = Large{
                values: inline26,
                shared: inline27,
            }
            var inline29 Outer = Outer{
                inner: inline28,
            }
            t39 = inline29
            loop_value__0 = t39
            var t40 int
            var inline18 *[16]int = &previous__0.values
            var inline19 int = array_get__Array_16_3int(*inline18, 0)
            var inline20 *[16]int = &previous__0.values
            var inline21 int = array_get__Array_16_3int(*inline20, 15)
            var inline22 int = inline19 + inline21
            var inline23 *ref_int_x = previous__0.shared
            var inline24 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(inline23)
            var inline25 int = inline22 + inline24
            t40 = inline25
            var inline16 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t40)
            _goml_runtime_core_string_println(inline16)
            continue
        } else {
            break Loop_loop2
        }
    }
    var t31 *Large = &first__0.inner
    var t32 *ref_int_x = (*t31).shared
    var t33 int
    var inline15 int = ref_get__Ref_3int(t32)
    t33 = inline15
    var t34 bool = t33 == 3
    var jp0 Outer
    if t34 {
        jp0 = first__0
    } else {
        var inline10 int = 12
        var inline11 [16]int = func(value int) [16]int {
            var result [16]int
            for index := range result {
                result[index] = value
            }
            return result
        }(inline10)
        var inline12 *ref_int_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__isize(inline10)
        var inline13 Large = Large{
            values: inline11,
            shared: inline12,
        }
        var inline14 Outer = Outer{
            inner: inline13,
        }
        jp0 = inline14
    }
    var t35 Large = jp0.inner
    var t36 int
    var inline2 *[16]int = &t35.values
    var inline3 int = array_get__Array_16_3int(*inline2, 0)
    var inline4 *[16]int = &t35.values
    var inline5 int = array_get__Array_16_3int(*inline4, 15)
    var inline6 int = inline3 + inline5
    var inline7 *ref_int_x = t35.shared
    var inline8 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(inline7)
    var inline9 int = inline6 + inline8
    t36 = inline9
    var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t36)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__isize(value__0 int) *ref_int_x {
    var t0 *ref_int_x = ref__Ref_3int(value__0)
    return t0
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(self__0 *ref_int_x) int {
    var t0 int = ref_get__Ref_3int(self__0)
    return t0
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__isize(self__0 *ref_int_x, value__0 int) struct{} {
    ref_set__Ref_3int(self__0, value__0)
    return struct{}{}
}

func println__T_isize(value__0 int) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_int_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__Outer(value__0 *Outer) *ref_Outer_x {
    var t0 *ref_Outer_x = ref__Ref_5Outer(*value__0)
    return t0
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__Outer(self__0 *ref_Outer_x) Outer {
    var t0 Outer = ref_get__Ref_5Outer(self__0)
    return t0
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__Outer(self__0 *ref_Outer_x, value__0 *Outer) struct{} {
    ref_set__Ref_5Outer(self__0, *value__0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_isize_i_to__string(self__0 int) string {
    var inline0 int64 = int64(int(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
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

func _goml_m_inherent_i_closure__env__captured__0_i_closure__env__captured__0_i_apply(env0 closure_env_captured_0) int {
    var field__0 *Large = &env0.field_0
    var inline0 *[16]int = &(*field__0).values
    var inline1 int = array_get__Array_16_3int(*inline0, 0)
    var inline2 *[16]int = &(*field__0).values
    var inline3 int = array_get__Array_16_3int(*inline2, 15)
    var inline4 int = inline1 + inline3
    var inline5 *ref_int_x = (*field__0).shared
    var inline6 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(inline5)
    var inline7 int = inline4 + inline6
    return inline7
}

func _goml_m_inherent_i_closure__env__main__1_i_closure__env__main__1_i_apply(env0 closure_env_main_1) int {
    var value__0 *Large = &env0.value_0
    var inline0 *[16]int = &(*value__0).values
    var inline1 int = array_get__Array_16_3int(*inline0, 0)
    var inline2 *[16]int = &(*value__0).values
    var inline3 int = array_get__Array_16_3int(*inline2, 15)
    var inline4 int = inline1 + inline3
    var inline5 *ref_int_x = (*value__0).shared
    var inline6 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(inline5)
    var inline7 int = inline4 + inline6
    return inline7
}

func main() {
    main0()
}
