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

func _goml_runtime_core_string_println(s string) struct{} {
    _goml_os.Stdout.WriteString(s + "\n")
    return struct{}{}
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

type closure_env_goml_builtin_range_0 struct {
    current_0 *ref_int_x
    end_1 int
}

type Ordering uint8

type Option__isize struct {
    _p0 int
    _tag uint8
}

func main0() struct{} {
    var value__0 string = "a你好z"
    var t0 bool = _goml_m_inherent_i_string_i_string_i_starts__with(value__0, "a你")
    var inline31 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t0)
    _goml_runtime_core_string_println(inline31)
    var t1 bool = _goml_m_inherent_i_string_i_string_i_starts__with(value__0, "你")
    var inline29 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t1)
    _goml_runtime_core_string_println(inline29)
    var t2 bool = _goml_m_inherent_i_string_i_string_i_starts__with(value__0, "")
    var inline27 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t2)
    _goml_runtime_core_string_println(inline27)
    var t3 bool = _goml_m_inherent_i_string_i_string_i_ends__with(value__0, "好z")
    var inline25 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t3)
    _goml_runtime_core_string_println(inline25)
    var t4 bool = _goml_m_inherent_i_string_i_string_i_ends__with(value__0, "你好")
    var inline23 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t4)
    _goml_runtime_core_string_println(inline23)
    var t5 bool
    var inline12 string = ""
    var inline13 int = _goml_m_inherent_i_string_i_string_i_byte__len(inline12)
    var inline14 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
    var inline15 bool = inline13 > inline14
    if inline15 {
        t5 = false
    } else {
        var inline16 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
        var inline17 int = _goml_m_inherent_i_string_i_string_i_byte__len(inline12)
        var inline18 int = inline16 - inline17
        var inline19 bool = _goml_m_inherent_i_string_i_string_i_is__char__boundary(value__0, inline18)
        if inline19 {
            var inline20 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
            var inline21 string = _goml_m_inherent_i_string_i_string_i_byte__slice(value__0, inline18, inline20)
            var inline22 bool = inline21 == inline12
            t5 = inline22
        } else {
            t5 = false
        }
    }
    var inline10 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t5)
    _goml_runtime_core_string_println(inline10)
    var t6 bool = _goml_m_inherent_i_string_i_string_i_contains(value__0, "你好")
    var inline8 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t6)
    _goml_runtime_core_string_println(inline8)
    var t7 bool = _goml_m_inherent_i_string_i_string_i_contains(value__0, "好z")
    var inline6 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t7)
    _goml_runtime_core_string_println(inline6)
    var t8 bool = _goml_m_inherent_i_string_i_string_i_contains(value__0, "你z")
    var inline4 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t8)
    _goml_runtime_core_string_println(inline4)
    var t9 bool = _goml_m_inherent_i_string_i_string_i_contains(value__0, "")
    var inline2 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t9)
    _goml_runtime_core_string_println(inline2)
    var t10 bool = _goml_m_inherent_i_string_i_string_i_contains(value__0, "a你好z!")
    var inline0 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t10)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_inherent_i_string_i_string_i_starts__with(self__0 string, prefix__0 string) bool {
    var t0 int
    var inline6 int = _goml_runtime_core_string_len(prefix__0)
    t0 = inline6
    var t1 int
    var inline5 int = _goml_runtime_core_string_len(self__0)
    t1 = inline5
    var t2 bool = t0 <= t1
    var jp0 bool
    if t2 {
        var t6 int
        var inline4 int = _goml_runtime_core_string_len(prefix__0)
        t6 = inline4
        var inline3 bool = string_is_char_boundary(self__0, t6)
        jp0 = inline3
    } else {
        jp0 = false
    }
    if jp0 {
        var t3 int
        var inline2 int = _goml_runtime_core_string_len(prefix__0)
        t3 = inline2
        var t4 string
        var inline0 int = 0
        var inline1 string = string_byte_slice(self__0, inline0, t3)
        t4 = inline1
        var t5 bool = t4 == prefix__0
        return t5
    } else {
        return false
    }
}

func _goml_m_inherent_i_string_i_string_i_ends__with(self__0 string, suffix__0 string) bool {
    var t0 int
    var inline6 int = _goml_runtime_core_string_len(suffix__0)
    t0 = inline6
    var t1 int
    var inline5 int = _goml_runtime_core_string_len(self__0)
    t1 = inline5
    var t2 bool = t0 > t1
    if t2 {
        return false
    } else {
        var t3 int
        var inline4 int = _goml_runtime_core_string_len(self__0)
        t3 = inline4
        var t4 int
        var inline3 int = _goml_runtime_core_string_len(suffix__0)
        t4 = inline3
        var start__0 int = t3 - t4
        var t5 bool
        var inline2 bool = string_is_char_boundary(self__0, start__0)
        t5 = inline2
        if t5 {
            var t6 int
            var inline1 int = _goml_runtime_core_string_len(self__0)
            t6 = inline1
            var t7 string
            var inline0 string = string_byte_slice(self__0, start__0, t6)
            t7 = inline0
            var t8 bool = t7 == suffix__0
            return t8
        } else {
            return false
        }
    }
}

func _goml_m_inherent_i_string_i_string_i_contains(self__0 string, expected__0 string) bool {
    var t0 int
    var inline15 int = _goml_runtime_core_string_len(expected__0)
    t0 = inline15
    var t1 bool = t0 == 0
    if t1 {
        return true
    } else {
        var t2 int
        var inline14 int = _goml_runtime_core_string_len(expected__0)
        t2 = inline14
        var t3 int
        var inline13 int = _goml_runtime_core_string_len(self__0)
        t3 = inline13
        var t4 bool = t2 > t3
        if t4 {
            return false
        } else {
            var t5 int
            var inline12 int = _goml_runtime_core_string_len(self__0)
            t5 = inline12
            var t6 int
            var inline11 int = _goml_runtime_core_string_len(expected__0)
            t6 = inline11
            var t7 int = t5 - t6
            var t8 int = t7 + 1
            var t9 FnIterator__isize
            var inline6 int = 0
            var inline7 *ref_int_x = ref__Ref_3int(inline6)
            var inline8 closure_env_goml_builtin_range_0 = closure_env_goml_builtin_range_0{
                current_0: inline7,
                end_1: t8,
            }
            var inline9 func() Option__isize = func() Option__isize {
                return _goml_m_inherent_i_closure__en_h5b1fcaf2e23588c4625108f446fe7c51_ange__0_i_apply(inline8)
            }
            var inline10 FnIterator__isize = _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__isize(inline9)
            t9 = inline10
            var for_iter0 FnIterator__isize
            for_iter0 = t9
            Loop_loop0:
            for {
                var for_next0 Option__isize
                var inline4 func() Option__isize = for_iter0.next_fn
                var inline5 Option__isize = inline4()
                for_next0 = inline5
                switch for_next0._tag {
                case 0:
                    break Loop_loop0
                case 1:
                    var x0 int = for_next0._p0
                    var t10 int
                    var inline3 int = _goml_runtime_core_string_len(expected__0)
                    t10 = inline3
                    var end__0 int = x0 + t10
                    var t11 bool
                    var inline2 bool = string_is_char_boundary(self__0, x0)
                    t11 = inline2
                    var jp0 bool
                    if t11 {
                        var inline1 bool = string_is_char_boundary(self__0, end__0)
                        jp0 = inline1
                    } else {
                        jp0 = false
                    }
                    var jp1 bool
                    if jp0 {
                        var t12 string
                        var inline0 string = string_byte_slice(self__0, x0, end__0)
                        t12 = inline0
                        var t13 bool = t12 == expected__0
                        jp1 = t13
                    } else {
                        jp1 = false
                    }
                    if jp1 {
                        return true
                    } else {
                        continue
                    }
                default:
                    panic("non-exhaustive match")
                }
            }
            return false
        }
    }
}

func _goml_m_trait__impl_i_ToString_i_bool_i_to__string(self__0 bool) string {
    var t0 string = _goml_runtime_core_bool_to_string(self__0)
    return t0
}

func _goml_m_inherent_i_string_i_string_i_byte__len(self__0 string) int {
    var t0 int = _goml_runtime_core_string_len(self__0)
    return t0
}

func _goml_m_inherent_i_string_i_string_i_is__char__boundary(self__0 string, index__0 int) bool {
    var t0 bool = string_is_char_boundary(self__0, index__0)
    return t0
}

func _goml_m_inherent_i_string_i_string_i_byte__slice(self__0 string, start__0 int, end__0 int) string {
    var inline0 bool = string_is_char_boundary(self__0, start__0)
    var inline1 bool
    if inline0 {
        var inline4 bool = string_is_char_boundary(self__0, end__0)
        inline1 = inline4
    } else {
        inline1 = false
    }
    if inline1 {
        var inline2 string = _goml_runtime_core_string_byte_slice(self__0, start__0, end__0)
        return inline2
    } else {
        var inline3 string = _goml_runtime_core_string_byte_slice(self__0, -1, -1)
        return inline3
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

func main() {
    main0()
}
