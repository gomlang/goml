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

func _goml_runtime_core_string_println(s string) struct{} {
    _goml_os.Stdout.WriteString(s + "\n")
    return struct{}{}
}

type _goml_vec_int struct {
    items []int
}

func vec_get__Vec_3int(vec *_goml_vec_int, index int) int {
    return vec.items[index]
}

func vec_len__Vec_3int(vec *_goml_vec_int) int {
    return int(len(vec.items))
}

type _goml_vec_string struct {
    items []string
}

func vec_new__Vec_6string() *_goml_vec_string {
    return &_goml_vec_string{
        items: nil,
    }
}

func vec_get__Vec_6string(vec *_goml_vec_string, index int) string {
    return vec.items[index]
}

func vec_len__Vec_6string(vec *_goml_vec_string) int {
    return int(len(vec.items))
}

type _goml_vec_Pair struct {
    items []Pair
}

func vec_get__Vec_4Pair(vec *_goml_vec_Pair, index int) Pair {
    return vec.items[index]
}

func vec_len__Vec_4Pair(vec *_goml_vec_Pair) int {
    return int(len(vec.items))
}

type _goml_vec_uint32 struct {
    items []uint32
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

type Pair struct {
    left int
    right int
}

type Ordering uint8

func main0() struct{} {
    var t0 [3]int = [3]int{1, 2, 3}
    var values__0 *_goml_vec_int = func(values [3]int) *_goml_vec_int {
        var storage struct {
            vector _goml_vec_int
            values [3]int
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t0)
    var t1 bool = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_contains____T__isize(values__0, 2)
    var t2 string
    var inline20 string = _goml_runtime_core_bool_to_string(t1)
    t2 = inline20
    var inline18 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t2)
    _goml_runtime_core_string_println(inline18)
    var t3 bool = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_contains____T__isize(values__0, 9)
    var t4 string
    var inline17 string = _goml_runtime_core_bool_to_string(t3)
    t4 = inline17
    var inline15 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t4)
    _goml_runtime_core_string_println(inline15)
    var t5 [2]string = [2]string{"alpha", "beta"}
    var names__0 *_goml_vec_string = func(values [2]string) *_goml_vec_string {
        var storage struct {
            vector _goml_vec_string
            values [2]string
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t5)
    var t6 bool = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_contains____T__string(names__0, "beta")
    var t7 string
    var inline14 string = _goml_runtime_core_bool_to_string(t6)
    t7 = inline14
    var inline12 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t7)
    _goml_runtime_core_string_println(inline12)
    var t8 Pair = Pair{
        left: 1,
        right: 2,
    }
    var t9 [1]Pair = [1]Pair{t8}
    var pairs__0 *_goml_vec_Pair = func(values [1]Pair) *_goml_vec_Pair {
        var storage struct {
            vector _goml_vec_Pair
            values [1]Pair
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t9)
    var t10 Pair = Pair{
        left: 1,
        right: 2,
    }
    var t11 bool = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_contains____T__Pair(pairs__0, t10)
    var t12 string
    var inline11 string = _goml_runtime_core_bool_to_string(t11)
    t12 = inline11
    var inline9 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t12)
    _goml_runtime_core_string_println(inline9)
    var t13 bool
    var inline7 int = 3
    var inline8 bool = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_contains____T__isize(values__0, inline7)
    t13 = inline8
    var t14 string
    var inline6 string = _goml_runtime_core_bool_to_string(t13)
    t14 = inline6
    var inline4 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t14)
    _goml_runtime_core_string_println(inline4)
    var empty__0 *_goml_vec_string
    var inline3 *_goml_vec_string = vec_new__Vec_6string()
    empty__0 = inline3
    var t15 bool = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_contains____T__string(empty__0, "x")
    var t16 string
    var inline2 string = _goml_runtime_core_bool_to_string(t15)
    t16 = inline2
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t16)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_trait__impl_i_PartialEq_i_isize_i_eq(self__0 int, other__0 int) bool {
    var t0 bool = self__0 == other__0
    return t0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_contains____T__isize(self__0 *_goml_vec_int, expected__0 int) bool {
    var index__0 int = 0
    Loop_loop0:
    for {
        var t0 int
        var inline1 int = vec_len__Vec_3int(self__0)
        t0 = inline1
        var t1 bool = index__0 < t0
        if t1 {
            var t2 int = vec_get__Vec_3int(self__0, index__0)
            var t3 bool
            var inline0 bool = t2 == expected__0
            t3 = inline0
            if t3 {
                return true
            } else {
                var compound_old0 int = index__0
                var compound_value0 int = 1
                var t4 int = compound_old0 + compound_value0
                index__0 = t4
                continue
            }
        } else {
            break Loop_loop0
        }
    }
    return false
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_contains____T__string(self__0 *_goml_vec_string, expected__0 string) bool {
    var index__0 int = 0
    Loop_loop0:
    for {
        var t0 int
        var inline1 int = vec_len__Vec_6string(self__0)
        t0 = inline1
        var t1 bool = index__0 < t0
        if t1 {
            var t2 string = vec_get__Vec_6string(self__0, index__0)
            var t3 bool
            var inline0 bool = t2 == expected__0
            t3 = inline0
            if t3 {
                return true
            } else {
                var compound_old0 int = index__0
                var compound_value0 int = 1
                var t4 int = compound_old0 + compound_value0
                index__0 = t4
                continue
            }
        } else {
            break Loop_loop0
        }
    }
    return false
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_contains____T__Pair(self__0 *_goml_vec_Pair, expected__0 Pair) bool {
    var index__0 int = 0
    Loop_loop0:
    for {
        var t0 int
        var inline7 int = vec_len__Vec_4Pair(self__0)
        t0 = inline7
        var t1 bool = index__0 < t0
        if t1 {
            var t2 Pair = vec_get__Vec_4Pair(self__0, index__0)
            var t3 bool
            var inline0 bool
            var inline4 int = t2.left
            var inline5 int = expected__0.left
            var inline6 bool = _goml_m_trait__impl_i_PartialEq_i_isize_i_eq(inline4, inline5)
            inline0 = inline6
            if inline0 {
                var inline1 int = t2.right
                var inline2 int = expected__0.right
                var inline3 bool = _goml_m_trait__impl_i_PartialEq_i_isize_i_eq(inline1, inline2)
                t3 = inline3
                if t3 {
                    return true
                } else {
                    var compound_old0 int = index__0
                    var compound_value0 int = 1
                    var t4 int = compound_old0 + compound_value0
                    index__0 = t4
                    continue
                }
            } else {
                t3 = false
                if t3 {
                    return true
                } else {
                    var compound_old0 int = index__0
                    var compound_value0 int = 1
                    var t4 int = compound_old0 + compound_value0
                    index__0 = t4
                    continue
                }
            }
        } else {
            break Loop_loop0
        }
    }
    return false
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func main() {
    main0()
}
