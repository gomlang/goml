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

func array_get__Array_4_3int(arr [4]int, index int) int {
    return arr[index]
}

func array_get__Array_3_3int(arr [3]int, index int) int {
    return arr[index]
}

func array_get__Array_2_8Ref_3int(arr [2]*ref_int_x, index int) *ref_int_x {
    return arr[index]
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

type closure_env_constructor_0 struct {}

type Ordering uint8

type Option__isize struct {
    _p0 int
    _tag uint8
}

type Result__isize__string struct {
    _p1 string
    _p0 int
    _tag uint8
}

type Option__Result__isize__string struct {
    _p0 Result__isize__string
    _tag uint8
}

type Message__isize struct {
    _p0 int
    _tag uint8
}

type Message__string struct {
    _p0 string
    _tag uint8
}

var REPEATED [4]int = [4]int{7, 7, 7, 7}

func count(counter__0 *ref_int_x) int {
    var t0 int
    var inline2 int = ref_get__Ref_3int(counter__0)
    t0 = inline2
    var t1 int = t0 + 1
    ref_set__Ref_3int(counter__0, t1)
    var inline0 int = ref_get__Ref_3int(counter__0)
    return inline0
}

func choose(value__0 Option__isize) Result__isize__string {
    var x__0 int = 90
    switch value__0._tag {
    case 1:
        var x0 int = value__0._p0
        var t0 bool = x0 > 2
        if t0 {
            var t1 int = x0 + 1
            var t2 int = x0 + t1
            var t3 Result__isize__string = Result__isize__string{
                _p0: t2,
                _tag: 0,
            }
            return t3
        } else {
            var t4 Result__isize__string = Result__isize__string{
                _p0: x__0,
                _tag: 0,
            }
            return t4
        }
    default:
        var t5 Result__isize__string = Result__isize__string{
            _p0: x__0,
            _tag: 0,
        }
        return t5
    }
}

func main0() struct{} {
    var _eq_rhs0 [4]int = [4]int{7, 7, 7, 7}
    var t0 int = array_get__Array_4_3int(REPEATED, 0)
    var t1 int = array_get__Array_4_3int(_eq_rhs0, 0)
    var t2 bool = _goml_m_trait__impl_i_PartialEq_i_isize_i_eq(t0, t1)
    var jp0 bool
    if t2 {
        var t60 int = array_get__Array_4_3int(REPEATED, 1)
        var t61 int = array_get__Array_4_3int(_eq_rhs0, 1)
        var t62 bool
        var inline27 bool = t60 == t61
        t62 = inline27
        if t62 {
            var t63 int = array_get__Array_4_3int(REPEATED, 2)
            var t64 int = array_get__Array_4_3int(_eq_rhs0, 2)
            var t65 bool
            var inline26 bool = t63 == t64
            t65 = inline26
            if t65 {
                var t66 int = array_get__Array_4_3int(REPEATED, 3)
                var t67 int = array_get__Array_4_3int(_eq_rhs0, 3)
                var inline25 bool = t66 == t67
                jp0 = inline25
            } else {
                jp0 = false
            }
        } else {
            jp0 = false
        }
    } else {
        jp0 = false
    }
    println__T_bool(jp0)
    var t3 [3]int = [3]int{1, 1, 1}
    var _eq_rhs1 [3]int = [3]int{1, 1, 1}
    var t4 int = array_get__Array_3_3int(t3, 0)
    var t5 int = array_get__Array_3_3int(_eq_rhs1, 0)
    var t6 bool = _goml_m_trait__impl_i_PartialEq_i_isize_i_eq(t4, t5)
    var jp1 bool
    if t6 {
        var t55 int = array_get__Array_3_3int(t3, 1)
        var t56 int = array_get__Array_3_3int(_eq_rhs1, 1)
        var t57 bool
        var inline24 bool = t55 == t56
        t57 = inline24
        if t57 {
            var t58 int = array_get__Array_3_3int(t3, 2)
            var t59 int = array_get__Array_3_3int(_eq_rhs1, 2)
            var inline23 bool = t58 == t59
            jp1 = inline23
        } else {
            jp1 = false
        }
    } else {
        jp1 = false
    }
    println__T_bool(jp1)
    var t7 int = 2
    println__T_isize(t7)
    println__T_bool(true)
    var counter__0 *ref_int_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__isize(0)
    var t8 int = count(counter__0)
    var values__0 [3]int = func(value int) [3]int {
        var result [3]int
        for index := range result {
            result[index] = value
        }
        return result
    }(t8)
    var _eq_rhs2 [3]int = [3]int{1, 1, 1}
    var t9 int = array_get__Array_3_3int(values__0, 0)
    var t10 int = array_get__Array_3_3int(_eq_rhs2, 0)
    var t11 bool = _goml_m_trait__impl_i_PartialEq_i_isize_i_eq(t9, t10)
    var jp2 bool
    if t11 {
        var t50 int = array_get__Array_3_3int(values__0, 1)
        var t51 int = array_get__Array_3_3int(_eq_rhs2, 1)
        var t52 bool
        var inline22 bool = t50 == t51
        t52 = inline22
        if t52 {
            var t53 int = array_get__Array_3_3int(values__0, 2)
            var t54 int = array_get__Array_3_3int(_eq_rhs2, 2)
            var inline21 bool = t53 == t54
            jp2 = inline21
        } else {
            jp2 = false
        }
    } else {
        jp2 = false
    }
    println__T_bool(jp2)
    var t12 int = count(counter__0)
    func(value int) [0]int {
        var result [0]int
        for index := range result {
            result[index] = value
        }
        return result
    }(t12)
    println__T_bool(true)
    var t13 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(counter__0)
    println__T_isize(t13)
    var shared__0 [2]*ref_int_x = _goml_m_repeated_____d_const___h0a63f16598356e334d8cf47a759652c3__Ref_l_isize_r_(counter__0)
    var t14 *ref_int_x = array_get__Array_2_8Ref_3int(shared__0, 0)
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__isize(t14, 7)
    var t15 *ref_int_x = array_get__Array_2_8Ref_3int(shared__0, 1)
    var t16 int = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(t15)
    println__T_isize(t16)
    var _eq_lhs0 [4]int = _goml_m_repeated_____d_const__param_k_N___d_const__value_k_4____T__isize(5)
    var _eq_rhs3 [4]int = [4]int{5, 5, 5, 5}
    var t17 int = array_get__Array_4_3int(_eq_lhs0, 0)
    var t18 int = array_get__Array_4_3int(_eq_rhs3, 0)
    var t19 bool
    var inline20 bool = t17 == t18
    t19 = inline20
    var jp3 bool
    if t19 {
        var t42 int = array_get__Array_4_3int(_eq_lhs0, 1)
        var t43 int = array_get__Array_4_3int(_eq_rhs3, 1)
        var t44 bool
        var inline19 bool = t42 == t43
        t44 = inline19
        if t44 {
            var t45 int = array_get__Array_4_3int(_eq_lhs0, 2)
            var t46 int = array_get__Array_4_3int(_eq_rhs3, 2)
            var t47 bool
            var inline18 bool = t45 == t46
            t47 = inline18
            if t47 {
                var t48 int = array_get__Array_4_3int(_eq_lhs0, 3)
                var t49 int = array_get__Array_4_3int(_eq_rhs3, 3)
                var inline17 bool = t48 == t49
                jp3 = inline17
            } else {
                jp3 = false
            }
        } else {
            jp3 = false
        }
    } else {
        jp3 = false
    }
    println__T_bool(jp3)
    var t20 Option__isize = Option__isize{
        _p0: 3,
        _tag: 1,
    }
    var t21 Result__isize__string = choose(t20)
    var t22 int = _goml_m_inherent_i_Result_i_Re_h42a4ea439e2cbe7d88a97c5539f34da5_ing____T__isize(t21, -1)
    println__T_isize(t22)
    var t23 Option__isize = Option__isize{
        _p0: 1,
        _tag: 1,
    }
    var t24 Result__isize__string = choose(t23)
    var t25 int = _goml_m_inherent_i_Result_i_Re_h42a4ea439e2cbe7d88a97c5539f34da5_ing____T__isize(t24, -1)
    println__T_isize(t25)
    var t26 Result__isize__string = choose(Option__isize{
        _tag: 0,
    })
    var t27 int = _goml_m_inherent_i_Result_i_Re_h42a4ea439e2cbe7d88a97c5539f34da5_ing____T__isize(t26, -1)
    var inline15 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t27)
    _goml_runtime_core_string_println(inline15)
    var t28 Result__isize__string = Result__isize__string{
        _p0: 8,
        _tag: 0,
    }
    var nested__0 Option__Result__isize__string = Option__Result__isize__string{
        _p0: t28,
        _tag: 1,
    }
    var t29 Result__isize__string = Result__isize__string{
        _p0: 8,
        _tag: 0,
    }
    var t30 Option__Result__isize__string = Option__Result__isize__string{
        _p0: t29,
        _tag: 1,
    }
    var t31 bool = _goml_m_trait__impl_i_PartialEq_i_Option____Result____isize____string_i_eq(nested__0, t30)
    var inline13 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t31)
    _goml_runtime_core_string_println(inline13)
    var t32 closure_env_constructor_0 = closure_env_constructor_0{}
    var constructor__0 func(int) Message__isize = func(p0 int) Message__isize {
        return _goml_m_inherent_i_closure__en_hae69c7cbe9a71685b0e1ce48d4d7b4a7_ctor__0_i_apply(t32, p0)
    }
    var t33 Message__isize = constructor__0(9)
    var t34 string = _goml_m_trait__impl_i_Debug_i_Message____isize_i_debug(t33)
    var inline11 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t34)
    _goml_runtime_core_string_println(inline11)
    var named__0 Message__string = Message__string{
        _p0: "named",
        _tag: 2,
    }
    var t35 string = _goml_m_trait__impl_i_Debug_i_Message____string_i_debug(named__0)
    var inline9 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t35)
    _goml_runtime_core_string_println(inline9)
    var empty_message__0 Message__isize = Message__isize{
        _tag: 0,
    }
    var t36 string = _goml_m_trait__impl_i_Debug_i_Message____isize_i_debug(empty_message__0)
    var inline7 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t36)
    _goml_runtime_core_string_println(inline7)
    var index__0 int = 0
    Loop_loop0:
    for {
        var t38 bool = index__0 < 3
        if t38 {
            var mtmp0 Option__isize = Option__isize{
                _p0: index__0,
                _tag: 1,
            }
            switch mtmp0._tag {
            case 1:
                var x0 int = mtmp0._p0
                var t39 bool = x0 < 2
                if t39 {
                    var inline5 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(x0)
                    _goml_runtime_core_string_println(inline5)
                    var compound_old0 int = index__0
                    var compound_value0 int = 1
                    var t40 int = compound_old0 + compound_value0
                    index__0 = t40
                    continue
                } else {
                    break Loop_loop0
                }
            default:
                break Loop_loop0
            }
        } else {
            break Loop_loop0
        }
    }
    var inline3 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(index__0)
    _goml_runtime_core_string_println(inline3)
    var t37 int
    var inline2 int = ref_get__Ref_3int(counter__0)
    t37 = inline2
    var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t37)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__isize(self__0 *ref_int_x) int {
    var t0 int = ref_get__Ref_3int(self__0)
    return t0
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__isize(self__0 *ref_int_x, value__0 int) struct{} {
    ref_set__Ref_3int(self__0, value__0)
    return struct{}{}
}

func println__T_bool(value__0 bool) struct{} {
    var t0 string
    var inline0 string = _goml_runtime_core_bool_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_trait__impl_i_PartialEq_i_isize_i_eq(self__0 int, other__0 int) bool {
    var t0 bool = self__0 == other__0
    return t0
}

func println__T_isize(value__0 int) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_int_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__isize(value__0 int) *ref_int_x {
    var t0 *ref_int_x = ref__Ref_3int(value__0)
    return t0
}

func _goml_m_repeated_____d_const___h0a63f16598356e334d8cf47a759652c3__Ref_l_isize_r_(value__0 *ref_int_x) [2]*ref_int_x {
    var t0 [2]*ref_int_x = func(value *ref_int_x) [2]*ref_int_x {
        var result [2]*ref_int_x
        for index := range result {
            result[index] = value
        }
        return result
    }(value__0)
    return t0
}

func _goml_m_repeated_____d_const__param_k_N___d_const__value_k_4____T__isize(value__0 int) [4]int {
    var t0 [4]int = func(value int) [4]int {
        var result [4]int
        for index := range result {
            result[index] = value
        }
        return result
    }(value__0)
    return t0
}

func _goml_m_inherent_i_Result_i_Re_h42a4ea439e2cbe7d88a97c5539f34da5_ing____T__isize(self__0 Result__isize__string, fallback__0 int) int {
    switch self__0._tag {
    case 0:
        var x0 int = self__0._p0
        return x0
    case 1:
        return fallback__0
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_trait__impl_i_PartialEq_i_Option____Result____isize____string_i_eq(self__0 Option__Result__isize__string, other__0 Option__Result__isize__string) bool {
    switch other__0._tag {
    case 0:
        switch self__0._tag {
        case 0:
            return true
        default:
            return false
        }
    case 1:
        var x0 Result__isize__string = other__0._p0
        switch self__0._tag {
        case 1:
            var x1 Result__isize__string = self__0._p0
            var t0 bool = _goml_m_trait__impl_i_PartialEq_i_Result____isize____string_i_eq(x1, x0)
            return t0
        default:
            return false
        }
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_trait__impl_i_Debug_i_Message____isize_i_debug(self__0 Message__isize) string {
    switch self__0._tag {
    case 0:
        return "Message::Empty"
    case 1:
        var x0 int = self__0._p0
        var t0 string
        var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(x0)
        t0 = inline0
        var t1 string = "Message::Value(" + t0
        var t2 string = t1 + ")"
        return t2
    case 2:
        var x1 int = self__0._p0
        var t3 string = "Message::Named { " + "value: "
        var t4 string
        var inline1 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(x1)
        t4 = inline1
        var t5 string = t3 + t4
        var t6 string = t5 + " }"
        return t6
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_trait__impl_i_Debug_i_Message____string_i_debug(self__0 Message__string) string {
    switch self__0._tag {
    case 0:
        return "Message::Empty"
    case 1:
        var x0 string = self__0._p0
        var t0 string
        var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(x0)
        t0 = inline0
        var t1 string = "Message::Value(" + t0
        var t2 string = t1 + ")"
        return t2
    case 2:
        var x1 string = self__0._p0
        var t3 string = "Message::Named { " + "value: "
        var t4 string
        var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(x1)
        t4 = inline1
        var t5 string = t3 + t4
        var t6 string = t5 + " }"
        return t6
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_trait__impl_i_ToString_i_bool_i_to__string(self__0 bool) string {
    var t0 string = _goml_runtime_core_bool_to_string(self__0)
    return t0
}

func _goml_m_trait__impl_i_ToString_i_isize_i_to__string(self__0 int) string {
    var inline0 int64 = int64(int(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func _goml_m_trait__impl_i_PartialEq_i_Result____isize____string_i_eq(self__0 Result__isize__string, other__0 Result__isize__string) bool {
    switch other__0._tag {
    case 0:
        var x0 int = other__0._p0
        switch self__0._tag {
        case 0:
            var x1 int = self__0._p0
            var inline0 bool = x1 == x0
            return inline0
        default:
            return false
        }
    case 1:
        var x2 string = other__0._p1
        switch self__0._tag {
        case 1:
            var x3 string = self__0._p1
            var inline1 bool = x3 == x2
            return inline1
        default:
            return false
        }
    default:
        panic("non-exhaustive match")
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

func _goml_m_inherent_i_closure__en_hae69c7cbe9a71685b0e1ce48d4d7b4a7_ctor__0_i_apply(env0 closure_env_constructor_0, ctor_arg_0 int) Message__isize {
    var t0 Message__isize = Message__isize{
        _p0: ctor_arg_0,
        _tag: 1,
    }
    return t0
}

func main() {
    main0()
}
