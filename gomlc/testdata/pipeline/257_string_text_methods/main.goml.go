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

func _goml_runtime_core_string_len(s string) int {
    return int(len(s))
}

func _goml_runtime_core_string_byte_get(s string, i int) uint8 {
    return s[i]
}

func _goml_runtime_core_string_byte_slice(s string, start int, end int) string {
    return s[start:end]
}

func _goml_runtime_core_string_to_bytes(s string) *_goml_vec_uint8 {
    return &_goml_vec_uint8{
        items: []byte(s),
    }
}

func _goml_runtime_core_string_from_utf8(bytes *_goml_vec_uint8) Tuple2_4bool_6string {
    return Tuple2_4bool_6string{
        _0: true,
        _1: string(bytes.items),
    }
}

func _goml_runtime_core_char_from_uint32(value uint32) Tuple2_4bool_4char {
    return Tuple2_4bool_4char{
        _0: true,
        _1: rune(value),
    }
}

func _goml_runtime_core_string_println(s string) struct{} {
    _goml_os.Stdout.WriteString(s + "\n")
    return struct{}{}
}

type _goml_vec_string struct {
    items []string
}

func vec_new__Vec_6string() *_goml_vec_string {
    return &_goml_vec_string{
        items: nil,
    }
}

func vec_with_capacity__Vec_6string(capacity int) *_goml_vec_string {
    return &_goml_vec_string{
        items: make([]string, 0, capacity),
    }
}

func vec_push__Vec_6string(vec *_goml_vec_string, elem string) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_get__Vec_6string(vec *_goml_vec_string, index int) string {
    return vec.items[index]
}

func vec_set__Vec_6string(vec *_goml_vec_string, index int, value string) struct{} {
    vec.items[index] = value
    return struct{}{}
}

func vec_len__Vec_6string(vec *_goml_vec_string) int {
    return int(len(vec.items))
}

func vec_truncate__Vec_6string(vec *_goml_vec_string, new_len int) struct{} {
    if new_len < 0 {
        panic("negative vector length")
    }
    if new_len < int(len(vec.items)) {
        clear(vec.items[new_len:int(len(vec.items))])
        vec.items = vec.items[0:new_len]
    }
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

func vec_set__Vec_5uint8(vec *_goml_vec_uint8, index int, value uint8) struct{} {
    vec.items[index] = value
    return struct{}{}
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

type Tuple2_6string_6string struct {
    _0 string
    _1 string
}

type Tuple2_13Option__isize_13Option__isize struct {
    _0 Option__isize
    _1 Option__isize
}

type Tuple2_3int_4char struct {
    _0 int
    _1 rune
}

type Tuple2_4bool_6string struct {
    _0 bool
    _1 string
}

type Tuple3_4bool_4char_3int struct {
    _0 bool
    _1 rune
    _2 int
}

type Tuple2_4char_3int struct {
    _0 rune
    _1 int
}

type Tuple2_4bool_4char struct {
    _0 bool
    _1 rune
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

type FnIterator__isize struct {
    next_fn func() Option__isize
}

type FnIterator__char struct {
    next_fn func() Option__char
}

type _goml_m_FnIterator_____o_isize_c_char_q_ struct {
    next_fn func() _goml_m_Option_____o_isize_c_char_q_
}

type closure_env_goml_builtin_range_0 struct {
    current_0 *ref_int_x
    end_1 int
}

type closure_env_inherent_string_string_chars_1 struct {
    self_0 string
    index_1 *ref_int_x
}

type closure_env_inherent_string_string_char_indices_2 struct {
    index_0 *ref_int_x
    self_1 string
}

type Ordering uint8

type _goml_m_Option_____o_string_c_string_q_ struct {
    _p0 Tuple2_6string_6string
    _tag uint8
}

type Option__isize struct {
    _p0 int
    _tag uint8
}

type Option__string struct {
    _p0 string
    _tag uint8
}

type Option__char uint64

type _goml_m_Option_____o_isize_c_char_q_ struct {
    _p0 Tuple2_3int_4char
    _tag uint8
}

type _goml_m_Option_____o_char_c_isize_q_ struct {
    _p0 Tuple2_4char_3int
    _tag uint8
}

func main0() struct{} {
    var text__0 string = "  Héllo, World!  "
    var t0 string = _goml_m_inherent_i_string_i_string_i_trim(text__0)
    var t1 string = "[" + t0
    var t2 string = t1 + "]"
    println__T_string(t2)
    var t3 string = _goml_m_inherent_i_string_i_string_i_trim__start(text__0)
    var t4 string = "[" + t3
    var t5 string = t4 + "]"
    println__T_string(t5)
    var t6 string = _goml_m_inherent_i_string_i_string_i_trim__end(text__0)
    var t7 string = "[" + t6
    var t8 string = t7 + "]"
    println__T_string(t8)
    var t9 *_goml_vec_string = _goml_m_inherent_i_string_i_string_i_split(text__0, ",")
    var t10 string = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_join____T__string(t9, "|")
    println__T_string(t10)
    var t11 _goml_m_Option_____o_string_c_string_q_ = _goml_m_inherent_i_string_i_string_i_split__once(text__0, ",")
    var t12 Tuple2_6string_6string = Tuple2_6string_6string{
        _0: "",
        _1: "",
    }
    var t13 Tuple2_6string_6string = _goml_m_inherent_i_Option_i_Op_h43ad6114bf7944dd0f967bbedf6419b6_ing_c_string_q_(t11, t12)
    var t14 string = t13._1
    println__T_string(t14)
    var t15 *_goml_vec_string = _goml_m_inherent_i_string_i_string_i_lines(text__0)
    var t16 int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_len____T__string(t15)
    var t17 string = _goml_m_inherent_i_isize_i_isize_i_to__string(t16)
    println__T_string(t17)
    var t18 Option__isize = _goml_m_inherent_i_string_i_string_i_find(text__0, "World")
    var t19 int = _goml_m_inherent_i_Option_i_Option_l_T_r__i_unwrap__or____T__isize(t18, -1)
    var t20 string = _goml_m_inherent_i_isize_i_isize_i_to__string(t19)
    println__T_string(t20)
    var t21 Option__isize = _goml_m_inherent_i_string_i_string_i_rfind(text__0, "l")
    var t22 int = _goml_m_inherent_i_Option_i_Option_l_T_r__i_unwrap__or____T__isize(t21, -1)
    var t23 string = _goml_m_inherent_i_isize_i_isize_i_to__string(t22)
    println__T_string(t23)
    var t24 Option__isize
    var inline27 string = "lo"
    var inline28 Option__isize = _goml_m_inherent_i_string_i_string_i_find(text__0, inline27)
    t24 = inline28
    var t25 int = _goml_m_inherent_i_Option_i_Option_l_T_r__i_unwrap__or____T__isize(t24, -1)
    var t26 string
    var inline26 string = __goml_builtin_int_to_string(t25)
    t26 = inline26
    var inline24 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t26)
    _goml_runtime_core_string_println(inline24)
    var t27 bool = _goml_m_inherent_i_string_i_string_i_starts__with__at(text__0, 2, "Hé")
    var t28 string
    var inline23 string = _goml_runtime_core_bool_to_string(t27)
    t28 = inline23
    var inline21 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t28)
    _goml_runtime_core_string_println(inline21)
    var t29 int = _goml_m_inherent_i_string_i_string_i_char__count(text__0)
    var t30 string
    var inline20 string = __goml_builtin_int_to_string(t29)
    t30 = inline20
    var inline18 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t30)
    _goml_runtime_core_string_println(inline18)
    var t31 Option__string = _goml_m_inherent_i_string_i_string_i_slice__chars(text__0, 2, 7)
    var t32 string
    var inline16 string = "none"
    switch t31._tag {
    case 0:
        t32 = inline16
    case 1:
        var inline17 string = t31._p0
        t32 = inline17
    default:
        panic("non-exhaustive match")
    }
    var inline14 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t32)
    _goml_runtime_core_string_println(inline14)
    var t33 string = _goml_m_inherent_i_string_i_string_i_replace(text__0, "l", "L")
    var inline12 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t33)
    _goml_runtime_core_string_println(inline12)
    var t34 string = _goml_m_inherent_i_string_i_string_i_repeat("ab", 3)
    var inline10 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t34)
    _goml_runtime_core_string_println(inline10)
    var t35 bool = _goml_m_inherent_i_string_i_string_i_is__ascii(text__0)
    var t36 string
    var inline9 string = _goml_runtime_core_bool_to_string(t35)
    t36 = inline9
    var inline7 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t36)
    _goml_runtime_core_string_println(inline7)
    var t37 bool = _goml_m_inherent_i_string_i_string_i_eq__ignore__ascii__case("ABC", "abc")
    var t38 string
    var inline6 string = _goml_runtime_core_bool_to_string(t37)
    t38 = inline6
    var inline4 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t38)
    _goml_runtime_core_string_println(inline4)
    var t39 string = _goml_m_inherent_i_string_i_string_i_to__ascii__lowercase("AbC")
    var inline2 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t39)
    _goml_runtime_core_string_println(inline2)
    var t40 string = _goml_m_inherent_i_string_i_string_i_to__ascii__uppercase("aBc")
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t40)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func println__T_string(value__0 string) struct{} {
    var t0 string
    t0 = value__0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_inherent_i_string_i_string_i_trim(self__0 string) string {
    var t0 string = _goml_m_inherent_i_string_i_string_i_trim__start(self__0)
    var t1 string = _goml_m_inherent_i_string_i_string_i_trim__end(t0)
    return t1
}

func _goml_m_inherent_i_string_i_string_i_trim__start(self__0 string) string {
    var start__0 int = 0
    Loop_loop0:
    for {
        var t1 int
        var inline3 int = _goml_runtime_core_string_len(self__0)
        t1 = inline3
        var t2 bool = start__0 < t1
        var jp0 bool
        if t2 {
            var t5 uint8
            var inline2 uint8 = _goml_runtime_core_string_byte_get(self__0, start__0)
            t5 = inline2
            var t6 bool = ascii_is_whitespace(t5)
            jp0 = t6
        } else {
            jp0 = false
        }
        if jp0 {
            var compound_old0 int = start__0
            var compound_value0 int = 1
            var t3 int = compound_old0 + compound_value0
            start__0 = t3
            continue
        } else {
            break Loop_loop0
        }
    }
    var t0 int
    var inline1 int = _goml_runtime_core_string_len(self__0)
    t0 = inline1
    var inline0 string = string_byte_slice(self__0, start__0, t0)
    return inline0
}

func _goml_m_inherent_i_string_i_string_i_trim__end(self__0 string) string {
    var end__0 int
    var inline3 int = _goml_runtime_core_string_len(self__0)
    end__0 = inline3
    Loop_loop0:
    for {
        var t0 bool = end__0 > 0
        var jp0 bool
        if t0 {
            var t3 int = end__0 - 1
            var t4 uint8
            var inline2 uint8 = _goml_runtime_core_string_byte_get(self__0, t3)
            t4 = inline2
            var t5 bool = ascii_is_whitespace(t4)
            jp0 = t5
        } else {
            jp0 = false
        }
        if jp0 {
            var compound_old0 int = end__0
            var compound_value0 int = 1
            var t1 int = compound_old0 - compound_value0
            end__0 = t1
            continue
        } else {
            break Loop_loop0
        }
    }
    var inline0 int = 0
    var inline1 string = string_byte_slice(self__0, inline0, end__0)
    return inline1
}

func _goml_m_inherent_i_string_i_string_i_split(self__0 string, separator__0 string) *_goml_vec_string {
    var result__0 *_goml_vec_string
    var inline7 *_goml_vec_string = vec_new__Vec_6string()
    result__0 = inline7
    var separator_len__0 int
    var inline6 int = _goml_runtime_core_string_len(separator__0)
    separator_len__0 = inline6
    var value_len__0 int
    var inline5 int = _goml_runtime_core_string_len(self__0)
    value_len__0 = inline5
    var t0 bool = separator_len__0 == 0
    if t0 {
        vec_push__Vec_6string(result__0, self__0)
        return result__0
    } else {
        var start__0 int = 0
        Loop_loop_expr0:
        for {
            var mtmp0 Option__isize = string_find_from(self__0, separator__0, start__0)
            switch mtmp0._tag {
            case 0:
                var t1 string
                var inline1 string = string_byte_slice(self__0, start__0, value_len__0)
                t1 = inline1
                vec_push__Vec_6string(result__0, t1)
                break Loop_loop_expr0
            case 1:
                var x0 int = mtmp0._p0
                var t2 string
                var inline3 string = string_byte_slice(self__0, start__0, x0)
                t2 = inline3
                vec_push__Vec_6string(result__0, t2)
                var t3 int = x0 + separator_len__0
                start__0 = t3
                continue
            default:
                panic("non-exhaustive match")
            }
        }
        return result__0
    }
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_join____T__string(self__0 *_goml_vec_string, separator__0 string) string {
    var t0 int
    var inline22 int = vec_len__Vec_6string(self__0)
    t0 = inline22
    var parts__0 *_goml_vec_string
    var inline21 *_goml_vec_string = vec_with_capacity__Vec_6string(t0)
    parts__0 = inline21
    var t1 int
    var inline20 int = vec_len__Vec_6string(self__0)
    t1 = inline20
    var t2 FnIterator__isize
    var inline15 int = 0
    var inline16 *ref_int_x = ref__Ref_3int(inline15)
    var inline17 closure_env_goml_builtin_range_0 = closure_env_goml_builtin_range_0{
        current_0: inline16,
        end_1: t1,
    }
    var inline18 func() Option__isize = func() Option__isize {
        return _goml_m_inherent_i_closure__en_h5b1fcaf2e23588c4625108f446fe7c51_ange__0_i_apply(inline17)
    }
    var inline19 FnIterator__isize = _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__isize(inline18)
    t2 = inline19
    var for_iter0 FnIterator__isize
    for_iter0 = t2
    Loop_loop0:
    for {
        var for_next1 Option__isize
        var inline13 func() Option__isize = for_iter0.next_fn
        var inline14 Option__isize = inline13()
        for_next1 = inline14
        switch for_next1._tag {
        case 0:
            break Loop_loop0
        case 1:
            var x1 int = for_next1._p0
            var t10 string = vec_get__Vec_6string(self__0, x1)
            var t11 string
            t11 = t10
            vec_push__Vec_6string(parts__0, t11)
            continue
        default:
            panic("non-exhaustive match")
        }
    }
    var t3 int
    var inline11 int = vec_len__Vec_6string(parts__0)
    t3 = inline11
    var t4 int = t3 * 2
    var result__0 *_goml_vec_string
    var inline10 *_goml_vec_string = vec_with_capacity__Vec_6string(t4)
    result__0 = inline10
    var t5 int
    var inline9 int = vec_len__Vec_6string(parts__0)
    t5 = inline9
    var t6 FnIterator__isize
    var inline4 int = 0
    var inline5 *ref_int_x = ref__Ref_3int(inline4)
    var inline6 closure_env_goml_builtin_range_0 = closure_env_goml_builtin_range_0{
        current_0: inline5,
        end_1: t5,
    }
    var inline7 func() Option__isize = func() Option__isize {
        return _goml_m_inherent_i_closure__en_h5b1fcaf2e23588c4625108f446fe7c51_ange__0_i_apply(inline6)
    }
    var inline8 FnIterator__isize = _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__isize(inline7)
    t6 = inline8
    var for_iter1 FnIterator__isize
    for_iter1 = t6
    Loop_loop1:
    for {
        var for_next0 Option__isize
        var inline2 func() Option__isize = for_iter1.next_fn
        var inline3 Option__isize = inline2()
        for_next0 = inline3
        switch for_next0._tag {
        case 0:
            break Loop_loop1
        case 1:
            var x0 int = for_next0._p0
            var t8 bool = x0 > 0
            if t8 {
                vec_push__Vec_6string(result__0, separator__0)
            } else {}
            var t9 string = vec_get__Vec_6string(parts__0, x0)
            vec_push__Vec_6string(result__0, t9)
            continue
        default:
            panic("non-exhaustive match")
        }
    }
    var t7 string = __goml_builtin_string_concat(result__0)
    return t7
}

func _goml_m_inherent_i_string_i_string_i_split__once(self__0 string, separator__0 string) _goml_m_Option_____o_string_c_string_q_ {
    var separator_len__0 int
    var inline5 int = _goml_runtime_core_string_len(separator__0)
    separator_len__0 = inline5
    var value_len__0 int
    var inline4 int = _goml_runtime_core_string_len(self__0)
    value_len__0 = inline4
    var t0 bool = separator_len__0 == 0
    if t0 {
        return _goml_m_Option_____o_string_c_string_q_{
            _tag: 0,
        }
    } else {
        var mtmp0 Option__isize
        var inline3 Option__isize = string_find_from(self__0, separator__0, 0)
        mtmp0 = inline3
        switch mtmp0._tag {
        case 0:
            return _goml_m_Option_____o_string_c_string_q_{
                _tag: 0,
            }
        case 1:
            var x0 int = mtmp0._p0
            var t1 string
            var inline1 int = 0
            var inline2 string = string_byte_slice(self__0, inline1, x0)
            t1 = inline2
            var t2 int = x0 + separator_len__0
            var t3 string
            var inline0 string = string_byte_slice(self__0, t2, value_len__0)
            t3 = inline0
            var t4 Tuple2_6string_6string = Tuple2_6string_6string{
                _0: t1,
                _1: t3,
            }
            var t5 _goml_m_Option_____o_string_c_string_q_ = _goml_m_Option_____o_string_c_string_q_{
                _p0: t4,
                _tag: 1,
            }
            return t5
        default:
            panic("non-exhaustive match")
        }
    }
}

func _goml_m_inherent_i_Option_i_Op_h43ad6114bf7944dd0f967bbedf6419b6_ing_c_string_q_(self__0 _goml_m_Option_____o_string_c_string_q_, fallback__0 Tuple2_6string_6string) Tuple2_6string_6string {
    switch self__0._tag {
    case 0:
        return fallback__0
    case 1:
        var x0 Tuple2_6string_6string = self__0._p0
        return x0
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_inherent_i_string_i_string_i_lines(self__0 string) *_goml_vec_string {
    var result__0 *_goml_vec_string = _goml_m_inherent_i_string_i_string_i_split(self__0, "\n")
    var t0 bool
    var inline21 int = vec_len__Vec_6string(result__0)
    var inline22 bool = inline21 == 0
    t0 = inline22
    var t1 bool = !t0
    var jp0 bool
    if t1 {
        var t13 int
        var inline20 int = vec_len__Vec_6string(result__0)
        t13 = inline20
        var t14 int = t13 - 1
        var t15 string = vec_get__Vec_6string(result__0, t14)
        var t16 bool = t15 == ""
        jp0 = t16
    } else {
        jp0 = false
    }
    if jp0 {
        var inline14 int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_len____T__string(result__0)
        var inline15 bool = inline14 == 0
        if inline15 {} else {
            var inline16 int = inline14 - 1
            vec_get__Vec_6string(result__0, inline16)
            var inline18 int = inline14 - 1
            _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_truncate____T__string(result__0, inline18)
        }
    } else {}
    var t2 int
    var inline13 int = vec_len__Vec_6string(result__0)
    t2 = inline13
    var t3 FnIterator__isize
    var inline8 int = 0
    var inline9 *ref_int_x = ref__Ref_3int(inline8)
    var inline10 closure_env_goml_builtin_range_0 = closure_env_goml_builtin_range_0{
        current_0: inline9,
        end_1: t2,
    }
    var inline11 func() Option__isize = func() Option__isize {
        return _goml_m_inherent_i_closure__en_h5b1fcaf2e23588c4625108f446fe7c51_ange__0_i_apply(inline10)
    }
    var inline12 FnIterator__isize = _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__isize(inline11)
    t3 = inline12
    var for_iter0 FnIterator__isize
    for_iter0 = t3
    Loop_loop0:
    for {
        var for_next0 Option__isize
        var inline6 func() Option__isize = for_iter0.next_fn
        var inline7 Option__isize = inline6()
        for_next0 = inline7
        switch for_next0._tag {
        case 0:
            break Loop_loop0
        case 1:
            var x0 int = for_next0._p0
            var line__0 string = vec_get__Vec_6string(result__0, x0)
            var t4 int
            var inline5 int = _goml_runtime_core_string_len(line__0)
            t4 = inline5
            var t5 bool = t4 > 0
            var jp1 bool
            if t5 {
                var t9 int
                var inline4 int = _goml_runtime_core_string_len(line__0)
                t9 = inline4
                var t10 int = t9 - 1
                var t11 uint8
                var inline3 uint8 = _goml_runtime_core_string_byte_get(line__0, t10)
                t11 = inline3
                var t12 bool = t11 == 13
                jp1 = t12
            } else {
                jp1 = false
            }
            if jp1 {
                vec_get__Vec_6string(result__0, x0)
                var t6 int
                var inline2 int = _goml_runtime_core_string_len(line__0)
                t6 = inline2
                var t7 int = t6 - 1
                var value0 string
                var inline0 int = 0
                var inline1 string = string_byte_slice(line__0, inline0, t7)
                value0 = inline1
                vec_set__Vec_6string(result__0, x0, value0)
                continue
            } else {
                continue
            }
        default:
            panic("non-exhaustive match")
        }
    }
    return result__0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_len____T__string(self__0 *_goml_vec_string) int {
    var t0 int = vec_len__Vec_6string(self__0)
    return t0
}

func _goml_m_inherent_i_isize_i_isize_i_to__string(self__0 int) string {
    var inline0 int64 = int64(int(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func _goml_m_inherent_i_string_i_string_i_find(self__0 string, expected__0 string) Option__isize {
    var t0 Option__isize = string_find_from(self__0, expected__0, 0)
    return t0
}

func _goml_m_inherent_i_Option_i_Option_l_T_r__i_unwrap__or____T__isize(self__0 Option__isize, fallback__0 int) int {
    switch self__0._tag {
    case 0:
        return fallback__0
    case 1:
        var x0 int = self__0._p0
        return x0
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_inherent_i_string_i_string_i_rfind(self__0 string, expected__0 string) Option__isize {
    var value_len__0 int
    var inline1 int = _goml_runtime_core_string_len(self__0)
    value_len__0 = inline1
    var expected_len__0 int
    var inline0 int = _goml_runtime_core_string_len(expected__0)
    expected_len__0 = inline0
    var t0 bool = expected_len__0 > value_len__0
    if t0 {
        return Option__isize{
            _tag: 0,
        }
    } else {
        var start__0 int = value_len__0 - expected_len__0
        Loop_loop0:
        for {
            var t1 bool = start__0 >= 0
            if t1 {
                var t2 bool = _goml_m_inherent_i_string_i_string_i_starts__with__at(self__0, start__0, expected__0)
                if t2 {
                    var t3 Option__isize = Option__isize{
                        _p0: start__0,
                        _tag: 1,
                    }
                    return t3
                } else {
                    var compound_old0 int = start__0
                    var compound_value0 int = 1
                    var t4 int = compound_old0 - compound_value0
                    start__0 = t4
                    continue
                }
            } else {
                break Loop_loop0
            }
        }
        return Option__isize{
            _tag: 0,
        }
    }
}

func _goml_m_inherent_i_string_i_string_i_starts__with__at(self__0 string, start__0 int, prefix__0 string) bool {
    var value_len__0 int
    var inline4 int = _goml_runtime_core_string_len(self__0)
    value_len__0 = inline4
    var prefix_len__0 int
    var inline3 int = _goml_runtime_core_string_len(prefix__0)
    prefix_len__0 = inline3
    var t0 bool = start__0 < 0
    var jp0 bool
    if t0 {
        jp0 = true
    } else {
        var t6 bool = start__0 > value_len__0
        jp0 = t6
    }
    var jp1 bool
    if jp0 {
        jp1 = true
    } else {
        var t4 int = value_len__0 - start__0
        var t5 bool = prefix_len__0 > t4
        jp1 = t5
    }
    if jp1 {
        return false
    } else {
        var end__0 int = start__0 + prefix_len__0
        var t1 bool
        var inline2 bool = string_is_char_boundary(self__0, start__0)
        t1 = inline2
        var jp2 bool
        if t1 {
            var inline1 bool = string_is_char_boundary(self__0, end__0)
            jp2 = inline1
        } else {
            jp2 = false
        }
        if jp2 {
            var t2 string
            var inline0 string = string_byte_slice(self__0, start__0, end__0)
            t2 = inline0
            var t3 bool = t2 == prefix__0
            return t3
        } else {
            return false
        }
    }
}

func _goml_m_inherent_i_string_i_string_i_char__count(self__0 string) int {
    var count__0 int = 0
    var t0 FnIterator__char
    var inline2 *ref_int_x = ref__Ref_3int(0)
    var inline3 closure_env_inherent_string_string_chars_1 = closure_env_inherent_string_string_chars_1{
        self_0: self__0,
        index_1: inline2,
    }
    var inline4 func() Option__char = func() Option__char {
        return _goml_m_inherent_i_closure__en_h6cc414bec6b61ab808929b69f0900e8f_hars__1_i_apply(inline3)
    }
    var inline5 FnIterator__char = _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__char(inline4)
    t0 = inline5
    var for_iter0 FnIterator__char
    for_iter0 = t0
    Loop_loop0:
    for {
        var for_next0 Option__char
        var inline0 func() Option__char = for_iter0.next_fn
        var inline1 Option__char = inline0()
        for_next0 = inline1
        switch for_next0 != Option__char(0) {
        case false:
            break Loop_loop0
        case true:
            var compound_old0 int = count__0
            var compound_value0 int = 1
            var t1 int = compound_old0 + compound_value0
            count__0 = t1
            continue
        default:
            panic("non-exhaustive match")
        }
    }
    return count__0
}

func _goml_m_inherent_i_string_i_string_i_slice__chars(self__0 string, start__0 int, end__0 int) Option__string {
    var t0 bool = start__0 < 0
    var jp0 bool
    if t0 {
        jp0 = true
    } else {
        var t20 bool = end__0 < start__0
        jp0 = t20
    }
    if jp0 {
        return Option__string{
            _tag: 0,
        }
    } else {
        var char_index__0 int = 0
        var t1 bool = start__0 == 0
        var jp1 Option__isize
        if t1 {
            var t19 Option__isize = Option__isize{
                _p0: 0,
                _tag: 1,
            }
            jp1 = t19
        } else {
            jp1 = Option__isize{
                _tag: 0,
            }
        }
        var start_byte__0 Option__isize = jp1
        var t2 bool = end__0 == 0
        var jp2 Option__isize
        if t2 {
            var t18 Option__isize = Option__isize{
                _p0: 0,
                _tag: 1,
            }
            jp2 = t18
        } else {
            jp2 = Option__isize{
                _tag: 0,
            }
        }
        var end_byte__0 Option__isize = jp2
        var t3 _goml_m_FnIterator_____o_isize_c_char_q_
        var inline5 *ref_int_x = ref__Ref_3int(0)
        var inline6 closure_env_inherent_string_string_char_indices_2 = closure_env_inherent_string_string_char_indices_2{
            index_0: inline5,
            self_1: self__0,
        }
        var inline7 func() _goml_m_Option_____o_isize_c_char_q_ = func() _goml_m_Option_____o_isize_c_char_q_ {
            return _goml_m_inherent_i_closure__en_hce113b295882d0c03e72873edb11c646_ices__2_i_apply(inline6)
        }
        var inline8 _goml_m_FnIterator_____o_isize_c_char_q_ = _goml_m_inherent_i_FnIterator__hae8fa7fd40ba2a3681525fbecfca8da7_isize_c_char_q_(inline7)
        t3 = inline8
        var for_iter0 _goml_m_FnIterator_____o_isize_c_char_q_
        for_iter0 = t3
        Loop_loop0:
        for {
            var for_next0 _goml_m_Option_____o_isize_c_char_q_
            var inline3 func() _goml_m_Option_____o_isize_c_char_q_ = for_iter0.next_fn
            var inline4 _goml_m_Option_____o_isize_c_char_q_ = inline3()
            for_next0 = inline4
            switch for_next0._tag {
            case 0:
                break Loop_loop0
            case 1:
                var x4 Tuple2_3int_4char = for_next0._p0
                var x5 int = x4._0
                var t12 bool = char_index__0 == start__0
                if t12 {
                    var t17 Option__isize = Option__isize{
                        _p0: x5,
                        _tag: 1,
                    }
                    start_byte__0 = t17
                } else {}
                var t13 bool = char_index__0 == end__0
                if t13 {
                    var t16 Option__isize = Option__isize{
                        _p0: x5,
                        _tag: 1,
                    }
                    end_byte__0 = t16
                } else {}
                var compound_old0 int = char_index__0
                var compound_value0 int = 1
                var t14 int = compound_old0 + compound_value0
                char_index__0 = t14
                continue
            default:
                panic("non-exhaustive match")
            }
        }
        var t4 bool = char_index__0 == start__0
        if t4 {
            var t10 int
            var inline2 int = _goml_runtime_core_string_len(self__0)
            t10 = inline2
            var t11 Option__isize = Option__isize{
                _p0: t10,
                _tag: 1,
            }
            start_byte__0 = t11
        } else {}
        var t5 bool = char_index__0 == end__0
        if t5 {
            var t8 int
            var inline1 int = _goml_runtime_core_string_len(self__0)
            t8 = inline1
            var t9 Option__isize = Option__isize{
                _p0: t8,
                _tag: 1,
            }
            end_byte__0 = t9
        } else {}
        var mtmp0 Tuple2_13Option__isize_13Option__isize = Tuple2_13Option__isize_13Option__isize{
            _0: start_byte__0,
            _1: end_byte__0,
        }
        var x0 Option__isize = mtmp0._0
        var x1 Option__isize = mtmp0._1
        switch x1._tag {
        case 1:
            var x2 int = x1._p0
            switch x0._tag {
            case 1:
                var x3 int = x0._p0
                var t6 string
                var inline0 string = string_byte_slice(self__0, x3, x2)
                t6 = inline0
                var t7 Option__string = Option__string{
                    _p0: t6,
                    _tag: 1,
                }
                return t7
            default:
                return Option__string{
                    _tag: 0,
                }
            }
        default:
            return Option__string{
                _tag: 0,
            }
        }
    }
}

func _goml_m_inherent_i_string_i_string_i_replace(self__0 string, expected__0 string, replacement__0 string) string {
    var t0 int
    var inline8 int = _goml_runtime_core_string_len(expected__0)
    t0 = inline8
    var t1 bool = t0 == 0
    if t1 {
        return self__0
    } else {
        var parts__0 *_goml_vec_string
        var inline7 *_goml_vec_string = vec_new__Vec_6string()
        parts__0 = inline7
        var start__0 int = 0
        Loop_loop_expr0:
        for {
            var mtmp0 Option__isize = string_find_from(self__0, expected__0, start__0)
            switch mtmp0._tag {
            case 0:
                var t3 int
                var inline2 int = _goml_runtime_core_string_len(self__0)
                t3 = inline2
                var t4 string
                var inline1 string = string_byte_slice(self__0, start__0, t3)
                t4 = inline1
                vec_push__Vec_6string(parts__0, t4)
                break Loop_loop_expr0
            case 1:
                var x0 int = mtmp0._p0
                var t5 string
                var inline6 string = string_byte_slice(self__0, start__0, x0)
                t5 = inline6
                vec_push__Vec_6string(parts__0, t5)
                vec_push__Vec_6string(parts__0, replacement__0)
                var t6 int
                var inline3 int = _goml_runtime_core_string_len(expected__0)
                t6 = inline3
                var t7 int = x0 + t6
                start__0 = t7
                continue
            default:
                panic("non-exhaustive match")
            }
        }
        var t2 string = __goml_builtin_string_concat(parts__0)
        return t2
    }
}

func _goml_m_inherent_i_string_i_string_i_repeat(self__0 string, count__0 int) string {
    var t0 bool = count__0 <= 0
    var jp0 bool
    if t0 {
        jp0 = true
    } else {
        var t3 int
        var inline9 int = _goml_runtime_core_string_len(self__0)
        t3 = inline9
        var t4 bool = t3 == 0
        jp0 = t4
    }
    if jp0 {
        return ""
    } else {
        var parts__0 *_goml_vec_string
        var inline8 *_goml_vec_string = vec_with_capacity__Vec_6string(count__0)
        parts__0 = inline8
        var t1 FnIterator__isize
        var inline3 int = 0
        var inline4 *ref_int_x = ref__Ref_3int(inline3)
        var inline5 closure_env_goml_builtin_range_0 = closure_env_goml_builtin_range_0{
            current_0: inline4,
            end_1: count__0,
        }
        var inline6 func() Option__isize = func() Option__isize {
            return _goml_m_inherent_i_closure__en_h5b1fcaf2e23588c4625108f446fe7c51_ange__0_i_apply(inline5)
        }
        var inline7 FnIterator__isize = _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__isize(inline6)
        t1 = inline7
        var for_iter0 FnIterator__isize
        for_iter0 = t1
        Loop_loop0:
        for {
            var for_next0 Option__isize
            var inline1 func() Option__isize = for_iter0.next_fn
            var inline2 Option__isize = inline1()
            for_next0 = inline2
            switch for_next0._tag {
            case 0:
                break Loop_loop0
            case 1:
                vec_push__Vec_6string(parts__0, self__0)
                continue
            default:
                panic("non-exhaustive match")
            }
        }
        var t2 string = __goml_builtin_string_concat(parts__0)
        return t2
    }
}

func _goml_m_inherent_i_string_i_string_i_is__ascii(self__0 string) bool {
    var t0 int
    var inline8 int = _goml_runtime_core_string_len(self__0)
    t0 = inline8
    var t1 FnIterator__isize
    var inline3 int = 0
    var inline4 *ref_int_x = ref__Ref_3int(inline3)
    var inline5 closure_env_goml_builtin_range_0 = closure_env_goml_builtin_range_0{
        current_0: inline4,
        end_1: t0,
    }
    var inline6 func() Option__isize = func() Option__isize {
        return _goml_m_inherent_i_closure__en_h5b1fcaf2e23588c4625108f446fe7c51_ange__0_i_apply(inline5)
    }
    var inline7 FnIterator__isize = _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__isize(inline6)
    t1 = inline7
    var for_iter0 FnIterator__isize
    for_iter0 = t1
    Loop_loop0:
    for {
        var for_next0 Option__isize
        var inline1 func() Option__isize = for_iter0.next_fn
        var inline2 Option__isize = inline1()
        for_next0 = inline2
        switch for_next0._tag {
        case 0:
            break Loop_loop0
        case 1:
            var x0 int = for_next0._p0
            var t2 uint8
            var inline0 uint8 = _goml_runtime_core_string_byte_get(self__0, x0)
            t2 = inline0
            var t3 bool = t2 > 127
            if t3 {
                return false
            } else {
                continue
            }
        default:
            panic("non-exhaustive match")
        }
    }
    return true
}

func _goml_m_inherent_i_string_i_string_i_eq__ignore__ascii__case(self__0 string, other__0 string) bool {
    var t0 int
    var inline14 int = _goml_runtime_core_string_len(self__0)
    t0 = inline14
    var t1 int
    var inline13 int = _goml_runtime_core_string_len(other__0)
    t1 = inline13
    var t2 bool = t0 != t1
    if t2 {
        return false
    } else {
        var t3 int
        var inline12 int = _goml_runtime_core_string_len(self__0)
        t3 = inline12
        var t4 FnIterator__isize
        var inline7 int = 0
        var inline8 *ref_int_x = ref__Ref_3int(inline7)
        var inline9 closure_env_goml_builtin_range_0 = closure_env_goml_builtin_range_0{
            current_0: inline8,
            end_1: t3,
        }
        var inline10 func() Option__isize = func() Option__isize {
            return _goml_m_inherent_i_closure__en_h5b1fcaf2e23588c4625108f446fe7c51_ange__0_i_apply(inline9)
        }
        var inline11 FnIterator__isize = _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__isize(inline10)
        t4 = inline11
        var for_iter0 FnIterator__isize
        for_iter0 = t4
        Loop_loop0:
        for {
            var for_next0 Option__isize
            var inline5 func() Option__isize = for_iter0.next_fn
            var inline6 Option__isize = inline5()
            for_next0 = inline6
            switch for_next0._tag {
            case 0:
                break Loop_loop0
            case 1:
                var x0 int = for_next0._p0
                var t5 uint8
                var inline4 uint8 = _goml_runtime_core_string_byte_get(self__0, x0)
                t5 = inline4
                var t6 uint8
                var inline3 uint8 = _goml_runtime_core_string_byte_get(other__0, x0)
                t6 = inline3
                var t7 bool
                var inline0 uint8 = ascii_to_lowercase(t5)
                var inline1 uint8 = ascii_to_lowercase(t6)
                var inline2 bool = inline0 == inline1
                t7 = inline2
                var t8 bool = !t7
                if t8 {
                    return false
                } else {
                    continue
                }
            default:
                panic("non-exhaustive match")
            }
        }
        return true
    }
}

func _goml_m_inherent_i_string_i_string_i_to__ascii__lowercase(self__0 string) string {
    var values__0 *_goml_vec_uint8
    var inline11 *_goml_vec_uint8 = _goml_runtime_core_string_to_bytes(self__0)
    values__0 = inline11
    var t0 int
    var inline10 int = vec_len__Vec_5uint8(values__0)
    t0 = inline10
    var t1 FnIterator__isize
    var inline5 int = 0
    var inline6 *ref_int_x = ref__Ref_3int(inline5)
    var inline7 closure_env_goml_builtin_range_0 = closure_env_goml_builtin_range_0{
        current_0: inline6,
        end_1: t0,
    }
    var inline8 func() Option__isize = func() Option__isize {
        return _goml_m_inherent_i_closure__en_h5b1fcaf2e23588c4625108f446fe7c51_ange__0_i_apply(inline7)
    }
    var inline9 FnIterator__isize = _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__isize(inline8)
    t1 = inline9
    var for_iter0 FnIterator__isize
    for_iter0 = t1
    var inline0_lhs uint8 = 97
    var inline0_rhs uint8 = 65
    var inline0 uint8 = inline0_lhs - inline0_rhs
    Loop_loop0:
    for {
        var for_next0 Option__isize
        var inline3 func() Option__isize = for_iter0.next_fn
        var inline4 Option__isize = inline3()
        for_next0 = inline4
        switch for_next0._tag {
        case 0:
            break Loop_loop0
        case 1:
            var x1 int = for_next0._p0
            vec_get__Vec_5uint8(values__0, x1)
            var t2 uint8 = vec_get__Vec_5uint8(values__0, x1)
            var value0 uint8
            var inline1 bool = ascii_is_uppercase(t2)
            if inline1 {
                var inline2 uint8 = t2 + inline0
                value0 = inline2
            } else {
                value0 = t2
            }
            vec_set__Vec_5uint8(values__0, x1, value0)
            continue
        default:
            panic("non-exhaustive match")
        }
    }
    var mtmp0 Tuple2_4bool_6string = string_from_utf8(values__0)
    var x0 string = mtmp0._1
    return x0
}

func _goml_m_inherent_i_string_i_string_i_to__ascii__uppercase(self__0 string) string {
    var values__0 *_goml_vec_uint8
    var inline11 *_goml_vec_uint8 = _goml_runtime_core_string_to_bytes(self__0)
    values__0 = inline11
    var t0 int
    var inline10 int = vec_len__Vec_5uint8(values__0)
    t0 = inline10
    var t1 FnIterator__isize
    var inline5 int = 0
    var inline6 *ref_int_x = ref__Ref_3int(inline5)
    var inline7 closure_env_goml_builtin_range_0 = closure_env_goml_builtin_range_0{
        current_0: inline6,
        end_1: t0,
    }
    var inline8 func() Option__isize = func() Option__isize {
        return _goml_m_inherent_i_closure__en_h5b1fcaf2e23588c4625108f446fe7c51_ange__0_i_apply(inline7)
    }
    var inline9 FnIterator__isize = _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__isize(inline8)
    t1 = inline9
    var for_iter0 FnIterator__isize
    for_iter0 = t1
    var inline0_lhs uint8 = 97
    var inline0_rhs uint8 = 65
    var inline0 uint8 = inline0_lhs - inline0_rhs
    Loop_loop0:
    for {
        var for_next0 Option__isize
        var inline3 func() Option__isize = for_iter0.next_fn
        var inline4 Option__isize = inline3()
        for_next0 = inline4
        switch for_next0._tag {
        case 0:
            break Loop_loop0
        case 1:
            var x1 int = for_next0._p0
            vec_get__Vec_5uint8(values__0, x1)
            var t2 uint8 = vec_get__Vec_5uint8(values__0, x1)
            var value0 uint8
            var inline1 bool = ascii_is_lowercase(t2)
            if inline1 {
                var inline2 uint8 = t2 - inline0
                value0 = inline2
            } else {
                value0 = t2
            }
            vec_set__Vec_5uint8(values__0, x1, value0)
            continue
        default:
            panic("non-exhaustive match")
        }
    }
    var mtmp0 Tuple2_4bool_6string = string_from_utf8(values__0)
    var x0 string = mtmp0._1
    return x0
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func _goml_m_inherent_i_string_i_string_i_byte__len(self__0 string) int {
    var t0 int = _goml_runtime_core_string_len(self__0)
    return t0
}

func ascii_is_whitespace(value__0 uint8) bool {
    var t0 bool = value__0 == 9
    var jp0 bool
    if t0 {
        jp0 = true
    } else {
        var t5 bool = value__0 == 10
        jp0 = t5
    }
    var jp1 bool
    if jp0 {
        jp1 = true
    } else {
        var t4 bool = value__0 == 11
        jp1 = t4
    }
    var jp2 bool
    if jp1 {
        jp2 = true
    } else {
        var t3 bool = value__0 == 12
        jp2 = t3
    }
    var jp3 bool
    if jp2 {
        jp3 = true
    } else {
        var t2 bool = value__0 == 13
        jp3 = t2
    }
    if jp3 {
        return true
    } else {
        var t1 bool = value__0 == 32
        return t1
    }
}

func _goml_m_inherent_i_string_i_string_i_byte__get(self__0 string, index__0 int) uint8 {
    var t0 uint8 = _goml_runtime_core_string_byte_get(self__0, index__0)
    return t0
}

func string_find_from(value__0 string, expected__0 string, start__0 int) Option__isize {
    var value_len__0 int
    var inline1 int = _goml_runtime_core_string_len(value__0)
    value_len__0 = inline1
    var expected_len__0 int
    var inline0 int = _goml_runtime_core_string_len(expected__0)
    expected_len__0 = inline0
    var t0 bool = start__0 < 0
    var jp0 bool
    if t0 {
        jp0 = true
    } else {
        var t10 bool = start__0 > value_len__0
        jp0 = t10
    }
    if jp0 {
        return Option__isize{
            _tag: 0,
        }
    } else {
        var t1 bool = expected_len__0 == 0
        if t1 {
            var t2 Option__isize = Option__isize{
                _p0: start__0,
                _tag: 1,
            }
            return t2
        } else {
            var t3 int = value_len__0 - start__0
            var t4 bool = expected_len__0 > t3
            if t4 {
                return Option__isize{
                    _tag: 0,
                }
            } else {
                var limit__0 int = value_len__0 - expected_len__0
                var index__0 int = start__0
                Loop_loop0:
                for {
                    var t5 bool = index__0 <= limit__0
                    if t5 {
                        var t6 bool = _goml_m_inherent_i_string_i_string_i_starts__with__at(value__0, index__0, expected__0)
                        if t6 {
                            var t7 Option__isize = Option__isize{
                                _p0: index__0,
                                _tag: 1,
                            }
                            return t7
                        } else {
                            var compound_old0 int = index__0
                            var compound_value0 int = 1
                            var t8 int = compound_old0 + compound_value0
                            index__0 = t8
                            continue
                        }
                    } else {
                        break Loop_loop0
                    }
                }
                return Option__isize{
                    _tag: 0,
                }
            }
        }
    }
}

func __goml_builtin_string_concat(values__0 *_goml_vec_string) string {
    var length__0 int = 0
    var value_index__0 int = 0
    Loop_loop0:
    for {
        var t9 int
        var inline5 int = vec_len__Vec_6string(values__0)
        t9 = inline5
        var t10 bool = value_index__0 < t9
        if t10 {
            var compound_old2 int = length__0
            var t11 string = vec_get__Vec_6string(values__0, value_index__0)
            var compound_value2 int
            var inline4 int = _goml_runtime_core_string_len(t11)
            compound_value2 = inline4
            var t12 int = compound_old2 + compound_value2
            length__0 = t12
            var compound_old3 int = value_index__0
            var compound_value3 int = 1
            var t14 int = compound_old3 + compound_value3
            value_index__0 = t14
            continue
        } else {
            break Loop_loop0
        }
    }
    var bytes__0 *_goml_vec_uint8 = vec_with_capacity__Vec_5uint8(length__0)
    value_index__0 = 0
    Loop_loop1:
    for {
        var t0 int
        var inline3 int = vec_len__Vec_6string(values__0)
        t0 = inline3
        var t1 bool = value_index__0 < t0
        if t1 {
            var value__0 string = vec_get__Vec_6string(values__0, value_index__0)
            var byte_index__0 int = 0
            Loop_loop2:
            for {
                var t4 int
                var inline2 int = _goml_runtime_core_string_len(value__0)
                t4 = inline2
                var t5 bool = byte_index__0 < t4
                if t5 {
                    var t6 uint8
                    var inline1 uint8 = _goml_runtime_core_string_byte_get(value__0, byte_index__0)
                    t6 = inline1
                    vec_push__Vec_5uint8(bytes__0, t6)
                    var compound_old1 int = byte_index__0
                    var compound_value1 int = 1
                    var t7 int = compound_old1 + compound_value1
                    byte_index__0 = t7
                    continue
                } else {
                    break Loop_loop2
                }
            }
            var compound_old0 int = value_index__0
            var compound_value0 int = 1
            var t2 int = compound_old0 + compound_value0
            value_index__0 = t2
            continue
        } else {
            break Loop_loop1
        }
    }
    var mtmp0 Tuple2_4bool_6string = _goml_runtime_core_string_from_utf8(bytes__0)
    var x0 string = mtmp0._1
    return x0
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

func ascii_to_lowercase(value__0 uint8) uint8 {
    var t0 bool
    var inline0 bool = value__0 >= 65
    if inline0 {
        var inline1 bool = value__0 <= 90
        t0 = inline1
    } else {
        t0 = false
    }
    if t0 {
        var t1_lhs uint8 = 97
        var t1_rhs uint8 = 65
        var t1 uint8 = t1_lhs - t1_rhs
        var t2 uint8 = value__0 + t1
        return t2
    } else {
        return value__0
    }
}

func string_from_utf8(bytes__0 *_goml_vec_uint8) Tuple2_4bool_6string {
    var mtmp0 Tuple2_4bool_6string = _goml_runtime_core_string_from_utf8(bytes__0)
    var x0 string = mtmp0._1
    var index__0 int = 0
    Loop_loop0:
    for {
        var t1 int
        var inline0 int = _goml_runtime_core_string_len(x0)
        t1 = inline0
        var t2 bool = index__0 < t1
        if t2 {
            var mtmp1 Tuple3_4bool_4char_3int = string_decode_utf8_at(x0, index__0)
            var x1 bool = mtmp1._0
            var x2 int = mtmp1._2
            if x1 {
                var compound_old0 int = index__0
                var t3 int = compound_old0 + x2
                index__0 = t3
                continue
            } else {
                var t5 Tuple2_4bool_6string = Tuple2_4bool_6string{
                    _0: false,
                    _1: "",
                }
                return t5
            }
        } else {
            break Loop_loop0
        }
    }
    var t0 Tuple2_4bool_6string = Tuple2_4bool_6string{
        _0: true,
        _1: x0,
    }
    return t0
}

func string_byte_slice(value__0 string, start__0 int, end__0 int) string {
    var t0 bool = string_is_char_boundary(value__0, start__0)
    var jp0 bool
    if t0 {
        var t3 bool = string_is_char_boundary(value__0, end__0)
        jp0 = t3
    } else {
        jp0 = false
    }
    if jp0 {
        var t1 string = _goml_runtime_core_string_byte_slice(value__0, start__0, end__0)
        return t1
    } else {
        var t2 string = _goml_runtime_core_string_byte_slice(value__0, -1, -1)
        return t2
    }
}

func _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__isize(next_fn__0 func() Option__isize) FnIterator__isize {
    var t0 FnIterator__isize = FnIterator__isize{
        next_fn: next_fn__0,
    }
    return t0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_truncate____T__string(self__0 *_goml_vec_string, len__0 int) struct{} {
    vec_truncate__Vec_6string(self__0, len__0)
    return struct{}{}
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

func string_is_char_boundary(value__0 string, index__0 int) bool {
    var t0 bool = index__0 < 0
    var jp0 bool
    if t0 {
        jp0 = true
    } else {
        var t6 int
        var inline2 int = _goml_runtime_core_string_len(value__0)
        t6 = inline2
        var t7 bool = index__0 > t6
        jp0 = t7
    }
    if jp0 {
        return false
    } else {
        var t1 int
        var inline1 int = _goml_runtime_core_string_len(value__0)
        t1 = inline1
        var t2 bool = index__0 == t1
        if t2 {
            return true
        } else {
            var t3 uint8
            var inline0 uint8 = _goml_runtime_core_string_byte_get(value__0, index__0)
            t3 = inline0
            var t4 uint8 = t3 & 192
            var t5 bool = t4 != 128
            return t5
        }
    }
}

func _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__char(next_fn__0 func() Option__char) FnIterator__char {
    var t0 FnIterator__char = FnIterator__char{
        next_fn: next_fn__0,
    }
    return t0
}

func _goml_m_inherent_i_FnIterator__hae8fa7fd40ba2a3681525fbecfca8da7_isize_c_char_q_(next_fn__0 func() _goml_m_Option_____o_isize_c_char_q_) _goml_m_FnIterator_____o_isize_c_char_q_ {
    var t0 _goml_m_FnIterator_____o_isize_c_char_q_ = _goml_m_FnIterator_____o_isize_c_char_q_{
        next_fn: next_fn__0,
    }
    return t0
}

func ascii_is_uppercase(value__0 uint8) bool {
    var t0 bool = value__0 >= 65
    if t0 {
        var t1 bool = value__0 <= 90
        return t1
    } else {
        return false
    }
}

func string_decode_utf8_at(value__0 string, index__0 int) Tuple3_4bool_4char_3int {
    var length__0 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
    var t0 bool = index__0 < 0
    var jp0 bool
    if t0 {
        jp0 = true
    } else {
        var t63 bool = index__0 >= length__0
        jp0 = t63
    }
    if jp0 {
        var inline25 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
            _0: false,
            _1: 0,
            _2: 0,
        }
        return inline25
    } else {
        var t1 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, index__0)
        var first__0 uint32 = uint32(uint8(t1))
        var t2 bool = first__0 < 128
        if t2 {
            var inline0 int = 1
            var inline1 Option__char = __goml_builtin_char_from_uint32(first__0)
            switch inline1 != Option__char(0) {
            case false:
                var inline2 Tuple3_4bool_4char_3int = utf8_invalid_decode()
                return inline2
            case true:
                var inline3 rune = rune(uint64(inline1) - 1)
                var inline4 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                    _0: true,
                    _1: inline3,
                    _2: inline0,
                }
                return inline4
            default:
                panic("non-exhaustive match")
            }
        } else {
            var t3 bool = first__0 < 194
            if t3 {
                var inline5 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                    _0: false,
                    _1: 0,
                    _2: 0,
                }
                return inline5
            } else {
                var t4 bool = first__0 < 224
                if t4 {
                    var t5 int = length__0 - index__0
                    var t6 bool = t5 < 2
                    if t6 {
                        var inline15 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                            _0: false,
                            _1: 0,
                            _2: 0,
                        }
                        return inline15
                    } else {
                        var t7 int = index__0 + 1
                        var t8 uint8
                        var inline14 uint8 = _goml_runtime_core_string_byte_get(value__0, t7)
                        t8 = inline14
                        var second__0 uint32 = uint32(uint8(t8))
                        var t9 bool
                        var inline12 bool = second__0 < 128
                        if inline12 {
                            t9 = true
                        } else {
                            var inline13 bool = second__0 > 191
                            t9 = inline13
                        }
                        if t9 {
                            var inline6 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                                _0: false,
                                _1: 0,
                                _2: 0,
                            }
                            return inline6
                        } else {
                            var t10 uint32 = first__0 & 31
                            var t11 uint32 = t10 << 6
                            var t12 uint32 = second__0 & 63
                            var t13 uint32 = t11 | t12
                            var inline7 int = 2
                            var inline8 Option__char = __goml_builtin_char_from_uint32(t13)
                            switch inline8 != Option__char(0) {
                            case false:
                                var inline9 Tuple3_4bool_4char_3int = utf8_invalid_decode()
                                return inline9
                            case true:
                                var inline10 rune = rune(uint64(inline8) - 1)
                                var inline11 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                                    _0: true,
                                    _1: inline10,
                                    _2: inline7,
                                }
                                return inline11
                            default:
                                panic("non-exhaustive match")
                            }
                        }
                    }
                } else {
                    var t14 bool = first__0 < 240
                    if t14 {
                        var t15 int = length__0 - index__0
                        var t16 bool = t15 < 3
                        if t16 {
                            var inline24 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                                _0: false,
                                _1: 0,
                                _2: 0,
                            }
                            return inline24
                        } else {
                            var t17 int = index__0 + 1
                            var t18 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, t17)
                            var second__1 uint32 = uint32(uint8(t18))
                            var t19 int = index__0 + 2
                            var t20 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, t19)
                            var third__0 uint32 = uint32(uint8(t20))
                            var t21 bool = utf8_invalid_continuation(second__1)
                            var jp1 bool
                            if t21 {
                                jp1 = true
                            } else {
                                var inline22 bool = third__0 < 128
                                if inline22 {
                                    jp1 = true
                                } else {
                                    var inline23 bool = third__0 > 191
                                    jp1 = inline23
                                }
                            }
                            var jp2 bool
                            if jp1 {
                                jp2 = true
                            } else {
                                var t31 bool = first__0 == 224
                                if t31 {
                                    var t32 bool = second__1 < 160
                                    jp2 = t32
                                } else {
                                    jp2 = false
                                }
                            }
                            var jp3 bool
                            if jp2 {
                                jp3 = true
                            } else {
                                var t29 bool = first__0 == 237
                                if t29 {
                                    var t30 bool = second__1 >= 160
                                    jp3 = t30
                                } else {
                                    jp3 = false
                                }
                            }
                            if jp3 {
                                var inline16 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                                    _0: false,
                                    _1: 0,
                                    _2: 0,
                                }
                                return inline16
                            } else {
                                var t22 uint32 = first__0 & 15
                                var t23 uint32 = t22 << 12
                                var t24 uint32 = second__1 & 63
                                var t25 uint32 = t24 << 6
                                var t26 uint32 = t23 | t25
                                var t27 uint32 = third__0 & 63
                                var t28 uint32 = t26 | t27
                                var inline17 int = 3
                                var inline18 Option__char = __goml_builtin_char_from_uint32(t28)
                                switch inline18 != Option__char(0) {
                                case false:
                                    var inline19 Tuple3_4bool_4char_3int = utf8_invalid_decode()
                                    return inline19
                                case true:
                                    var inline20 rune = rune(uint64(inline18) - 1)
                                    var inline21 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                                        _0: true,
                                        _1: inline20,
                                        _2: inline17,
                                    }
                                    return inline21
                                default:
                                    panic("non-exhaustive match")
                                }
                            }
                        }
                    } else {
                        var t33 bool = first__0 < 245
                        if t33 {
                            var t34 int = length__0 - index__0
                            var t35 bool = t34 < 4
                            if t35 {
                                var t61 Tuple3_4bool_4char_3int = utf8_invalid_decode()
                                return t61
                            } else {
                                var t36 int = index__0 + 1
                                var t37 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, t36)
                                var second__2 uint32 = uint32(uint8(t37))
                                var t38 int = index__0 + 2
                                var t39 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, t38)
                                var third__1 uint32 = uint32(uint8(t39))
                                var t40 int = index__0 + 3
                                var t41 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, t40)
                                var fourth__0 uint32 = uint32(uint8(t41))
                                var t42 bool = utf8_invalid_continuation(second__2)
                                var jp4 bool
                                if t42 {
                                    jp4 = true
                                } else {
                                    var t60 bool = utf8_invalid_continuation(third__1)
                                    jp4 = t60
                                }
                                var jp5 bool
                                if jp4 {
                                    jp5 = true
                                } else {
                                    var t59 bool = utf8_invalid_continuation(fourth__0)
                                    jp5 = t59
                                }
                                var jp6 bool
                                if jp5 {
                                    jp6 = true
                                } else {
                                    var t57 bool = first__0 == 240
                                    if t57 {
                                        var t58 bool = second__2 < 144
                                        jp6 = t58
                                    } else {
                                        jp6 = false
                                    }
                                }
                                var jp7 bool
                                if jp6 {
                                    jp7 = true
                                } else {
                                    var t55 bool = first__0 == 244
                                    if t55 {
                                        var t56 bool = second__2 > 143
                                        jp7 = t56
                                    } else {
                                        jp7 = false
                                    }
                                }
                                if jp7 {
                                    var t43 Tuple3_4bool_4char_3int = utf8_invalid_decode()
                                    return t43
                                } else {
                                    var t44 uint32 = first__0 & 7
                                    var t45 uint32 = t44 << 18
                                    var t46 uint32 = second__2 & 63
                                    var t47 uint32 = t46 << 12
                                    var t48 uint32 = t45 | t47
                                    var t49 uint32 = third__1 & 63
                                    var t50 uint32 = t49 << 6
                                    var t51 uint32 = t48 | t50
                                    var t52 uint32 = fourth__0 & 63
                                    var t53 uint32 = t51 | t52
                                    var t54 Tuple3_4bool_4char_3int = utf8_valid_decode(t53, 4)
                                    return t54
                                }
                            }
                        } else {
                            var t62 Tuple3_4bool_4char_3int = utf8_invalid_decode()
                            return t62
                        }
                    }
                }
            }
        }
    }
}

func ascii_is_lowercase(value__0 uint8) bool {
    var t0 bool = value__0 >= 97
    if t0 {
        var t1 bool = value__0 <= 122
        return t1
    } else {
        return false
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

func utf8_invalid_decode() Tuple3_4bool_4char_3int {
    var t0 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
        _0: false,
        _1: 0,
        _2: 0,
    }
    return t0
}

func utf8_valid_decode(value__0 uint32, width__0 int) Tuple3_4bool_4char_3int {
    var commute_field0 rune
    var inline1 bool = utf8_valid_scalar(value__0)
    if inline1 {
        var inline2 Tuple2_4bool_4char = _goml_runtime_core_char_from_uint32(value__0)
        var inline3 rune = inline2._1
        commute_field0 = inline3
        var t0 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
            _0: true,
            _1: commute_field0,
            _2: width__0,
        }
        return t0
    } else {
        var inline0 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
            _0: false,
            _1: 0,
            _2: 0,
        }
        return inline0
    }
}

func utf8_invalid_continuation(value__0 uint32) bool {
    var t0 bool = value__0 < 128
    if t0 {
        return true
    } else {
        var t1 bool = value__0 > 191
        return t1
    }
}

func __goml_builtin_char_from_uint32(value__0 uint32) Option__char {
    var t0 bool
    var inline0 bool = value__0 <= 1114111
    if inline0 {
        var inline1 bool = value__0 >= 55296
        var inline2 bool
        if inline1 {
            var inline4 bool = value__0 <= 57343
            inline2 = inline4
        } else {
            inline2 = false
        }
        var inline3 bool = !inline2
        t0 = inline3
    } else {
        t0 = false
    }
    if t0 {
        var mtmp0 Tuple2_4bool_4char = _goml_runtime_core_char_from_uint32(value__0)
        var x0 rune = mtmp0._1
        var t1 Option__char = Option__char(uint64(uint32(x0)) + 1)
        return t1
    } else {
        return Option__char(0)
    }
}

func utf8_valid_scalar(value__0 uint32) bool {
    var t0 bool = value__0 <= 1114111
    if t0 {
        var t1 bool = value__0 >= 55296
        var jp0 bool
        if t1 {
            var t3 bool = value__0 <= 57343
            jp0 = t3
        } else {
            jp0 = false
        }
        var t2 bool = !jp0
        return t2
    } else {
        return false
    }
}

func _goml_m_inherent_i_closure__en_h5b1fcaf2e23588c4625108f446fe7c51_ange__0_i_apply(env0 closure_env_goml_builtin_range_0) Option__isize {
    var current__0 *ref_int_x = env0.current_0
    var end__0 int = env0.end_1
    var value__0 int = ref_get__Ref_3int(current__0)
    var t0 bool = value__0 < end__0
    if t0 {
        var t1 int = value__0 + 1
        ref_set__Ref_3int(current__0, t1)
        var t2 Option__isize = Option__isize{
            _p0: value__0,
            _tag: 1,
        }
        return t2
    } else {
        return Option__isize{
            _tag: 0,
        }
    }
}

func _goml_m_inherent_i_closure__en_h6cc414bec6b61ab808929b69f0900e8f_hars__1_i_apply(env0 closure_env_inherent_string_string_chars_1) Option__char {
    var self__0 string = env0.self_0
    var index__0 *ref_int_x = env0.index_1
    var t0 int = ref_get__Ref_3int(index__0)
    var commute_field0 Tuple2_4char_3int
    var inline0 Tuple3_4bool_4char_3int = string_decode_utf8_at(self__0, t0)
    var inline1 bool = inline0._0
    var inline2 rune = inline0._1
    var inline3 int = inline0._2
    if inline1 {
        var inline4 Tuple2_4char_3int = Tuple2_4char_3int{
            _0: inline2,
            _1: inline3,
        }
        commute_field0 = inline4
        var x0 rune = commute_field0._0
        var x1 int = commute_field0._1
        var compound_old0 int = ref_get__Ref_3int(index__0)
        var t1 int = compound_old0 + x1
        ref_set__Ref_3int(index__0, t1)
        var t3 Option__char = Option__char(uint64(uint32(x0)) + 1)
        return t3
    } else {
        return Option__char(0)
    }
}

func _goml_m_inherent_i_closure__en_hce113b295882d0c03e72873edb11c646_ices__2_i_apply(env0 closure_env_inherent_string_string_char_indices_2) _goml_m_Option_____o_isize_c_char_q_ {
    var index__0 *ref_int_x = env0.index_0
    var self__0 string = env0.self_1
    var current__0 int = ref_get__Ref_3int(index__0)
    var commute_field0 Tuple2_4char_3int
    var inline0 Tuple3_4bool_4char_3int = string_decode_utf8_at(self__0, current__0)
    var inline1 bool = inline0._0
    var inline2 rune = inline0._1
    var inline3 int = inline0._2
    if inline1 {
        var inline4 Tuple2_4char_3int = Tuple2_4char_3int{
            _0: inline2,
            _1: inline3,
        }
        commute_field0 = inline4
        var x0 rune = commute_field0._0
        var x1 int = commute_field0._1
        var t0 int = current__0 + x1
        ref_set__Ref_3int(index__0, t0)
        var t1 Tuple2_3int_4char = Tuple2_3int_4char{
            _0: current__0,
            _1: x0,
        }
        var t2 _goml_m_Option_____o_isize_c_char_q_ = _goml_m_Option_____o_isize_c_char_q_{
            _p0: t1,
            _tag: 1,
        }
        return t2
    } else {
        return _goml_m_Option_____o_isize_c_char_q_{
            _tag: 0,
        }
    }
}

func main() {
    main0()
}
