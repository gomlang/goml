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

type _goml_vec_int struct {
    items []int
}

func vec_new__Vec_3int() *_goml_vec_int {
    return &_goml_vec_int{
        items: nil,
    }
}

func vec_with_capacity__Vec_3int(capacity int) *_goml_vec_int {
    return &_goml_vec_int{
        items: make([]int, 0, capacity),
    }
}

func vec_push__Vec_3int(vec *_goml_vec_int, elem int) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_get__Vec_3int(vec *_goml_vec_int, index int) int {
    return vec.items[index]
}

func vec_set__Vec_3int(vec *_goml_vec_int, index int, value int) struct{} {
    vec.items[index] = value
    return struct{}{}
}

func vec_len__Vec_3int(vec *_goml_vec_int) int {
    return int(len(vec.items))
}

func vec_truncate__Vec_3int(vec *_goml_vec_int, new_len int) struct{} {
    if new_len < 0 {
        panic("negative vector length")
    }
    if new_len < int(len(vec.items)) {
        clear(vec.items[new_len:int(len(vec.items))])
        vec.items = vec.items[0:new_len]
    }
    return struct{}{}
}

type _goml_vec_string struct {
    items []string
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

type _goml_vec_Tuple2_3int_3int struct {
    items []Tuple2_3int_3int
}

func vec_with_capacity__Vec_16Tuple2_3int_3int(capacity int) *_goml_vec_Tuple2_3int_3int {
    return &_goml_vec_Tuple2_3int_3int{
        items: make([]Tuple2_3int_3int, 0, capacity),
    }
}

func vec_push__Vec_16Tuple2_3int_3int(vec *_goml_vec_Tuple2_3int_3int, elem Tuple2_3int_3int) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_get__Vec_16Tuple2_3int_3int(vec *_goml_vec_Tuple2_3int_3int, index int) Tuple2_3int_3int {
    return vec.items[index]
}

func vec_set__Vec_16Tuple2_3int_3int(vec *_goml_vec_Tuple2_3int_3int, index int, value Tuple2_3int_3int) struct{} {
    vec.items[index] = value
    return struct{}{}
}

func vec_len__Vec_16Tuple2_3int_3int(vec *_goml_vec_Tuple2_3int_3int) int {
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

type Tuple2_3int_3int struct {
    _0 int
    _1 int
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

type FnIterator__isize struct {
    next_fn func() Option__isize
}

type closure_env_main_0 struct {}

type closure_env_main_1 struct {}

type closure_env_main_2 struct {}

type closure_env_main_3 struct {}

type closure_env_main_4 struct {}

type closure_env_main_5 struct {}

type closure_env_main_6 struct {}

type closure_env_main_7 struct {}

type closure_env_main_8 struct {}

type closure_env_main_9 struct {}

type closure_env_inherent_Vec_Vec_T_dedup_T_isize_10 struct {}

type closure_env_inherent_Vec_Vec_T_sort_by_ordering_T_string_11 struct {
    compare_0 func(string, string) Ordering
}

type closure_env_inherent_Vec_Vec_T_sort_by_ordering_T_isize_isize_12 struct {
    compare_0 func(Tuple2_3int_3int, Tuple2_3int_3int) Ordering
}

type closure_env_goml_builtin_range_13 struct {
    current_0 *ref_int_x
    end_1 int
}

type closure_env_inherent_Vec_Vec_T_iter_T_isize_14 struct {
    index_0 *ref_int_x
    len_1 int
    self_2 *_goml_vec_int
}

type FrozenVec__isize struct {
    values *_goml_vec_int
}

type Ordering uint8

const (
    Less Ordering = 0
    Equal Ordering = 1
    Greater Ordering = 2
)

type Option__isize struct {
    _p0 int
    _tag uint8
}

func main0() struct{} {
    var t0 [8]int = [8]int{3, 1, 4, 1, 5, 9, 2, 6}
    var values__0 *_goml_vec_int = func(values [8]int) *_goml_vec_int {
        var storage struct {
            vector _goml_vec_int
            values [8]int
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t0)
    var t1 closure_env_main_0 = closure_env_main_0{}
    var t2 func(int) bool = func(p0 int) bool {
        return _goml_m_inherent_i_closure__env__main__0_i_closure__env__main__0_i_apply(t1, p0)
    }
    var t3 Option__isize = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_position____T__isize(values__0, t2)
    var t4 int = _goml_m_inherent_i_Option_i_Option_l_T_r__i_unwrap__or____T__isize(t3, -1)
    var t5 string = _goml_m_inherent_i_isize_i_isize_i_to__string(t4)
    println__T_string(t5)
    var t6 closure_env_main_1 = closure_env_main_1{}
    var t7 func(int) bool = func(p0 int) bool {
        return _goml_m_inherent_i_closure__env__main__1_i_closure__env__main__1_i_apply(t6, p0)
    }
    var t8 Option__isize = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_position____T__isize(values__0, t7)
    var t9 int = _goml_m_inherent_i_Option_i_Option_l_T_r__i_unwrap__or____T__isize(t8, -1)
    var t10 string = _goml_m_inherent_i_isize_i_isize_i_to__string(t9)
    println__T_string(t10)
    var t11 closure_env_main_2 = closure_env_main_2{}
    var t12 func(int, int) int = func(p0 int, p1 int) int {
        return _goml_m_inherent_i_closure__env__main__2_i_closure__env__main__2_i_apply(t11, p0, p1)
    }
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_sort__by____T__isize(values__0, t12)
    var t13 string = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_join____T__isize(values__0, ",")
    println__T_string(t13)
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_dedup____T__isize(values__0)
    var t14 string = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_join____T__isize(values__0, ",")
    println__T_string(t14)
    var t15 [5]int = [5]int{1, 2, 4, 4, 5}
    var ordered__0 *_goml_vec_int = func(values [5]int) *_goml_vec_int {
        var storage struct {
            vector _goml_vec_int
            values [5]int
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t15)
    var t16 closure_env_main_3 = closure_env_main_3{}
    var t17 func(int) int = func(p0 int) int {
        return _goml_m_inherent_i_closure__env__main__3_i_closure__env__main__3_i_apply(t16, p0)
    }
    var t18 Option__isize = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_binary__search__by____T__isize(ordered__0, t17)
    var t19 int = _goml_m_inherent_i_Option_i_Option_l_T_r__i_unwrap__or____T__isize(t18, -1)
    var t20 string = _goml_m_inherent_i_isize_i_isize_i_to__string(t19)
    println__T_string(t20)
    var t21 closure_env_main_4 = closure_env_main_4{}
    var t22 func(int) int = func(p0 int) int {
        return _goml_m_inherent_i_closure__env__main__4_i_closure__env__main__4_i_apply(t21, p0)
    }
    var t23 Option__isize = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_binary__search__by____T__isize(ordered__0, t22)
    var t24 int = _goml_m_inherent_i_Option_i_Option_l_T_r__i_unwrap__or____T__isize(t23, -1)
    var t25 string = _goml_m_inherent_i_isize_i_isize_i_to__string(t24)
    println__T_string(t25)
    var t26 closure_env_main_5 = closure_env_main_5{}
    var t27 func(int, int) int = func(p0 int, p1 int) int {
        return _goml_m_inherent_i_closure__env__main__5_i_closure__env__main__5_i_apply(t26, p0, p1)
    }
    var t28 Option__isize = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_min__by____T__isize(ordered__0, t27)
    var t29 int = _goml_m_inherent_i_Option_i_Option_l_T_r__i_unwrap__or____T__isize(t28, 0)
    var t30 string = _goml_m_inherent_i_isize_i_isize_i_to__string(t29)
    println__T_string(t30)
    var t31 closure_env_main_6 = closure_env_main_6{}
    var t32 func(int, int) int = func(p0 int, p1 int) int {
        return _goml_m_inherent_i_closure__env__main__6_i_closure__env__main__6_i_apply(t31, p0, p1)
    }
    var t33 Option__isize = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_max__by____T__isize(ordered__0, t32)
    var t34 int = _goml_m_inherent_i_Option_i_Option_l_T_r__i_unwrap__or____T__isize(t33, 0)
    var t35 string = _goml_m_inherent_i_isize_i_isize_i_to__string(t34)
    println__T_string(t35)
    var t36 [3]string = [3]string{"beta", "alpha", "gamma"}
    var names__0 *_goml_vec_string = func(values [3]string) *_goml_vec_string {
        var storage struct {
            vector _goml_vec_string
            values [3]string
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t36)
    var t37 closure_env_main_7 = closure_env_main_7{}
    var t38 func(string, string) Ordering = func(p0 string, p1 string) Ordering {
        return _goml_m_inherent_i_closure__env__main__7_i_closure__env__main__7_i_apply(t37, p0, p1)
    }
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_sort__by__ordering____T__string(names__0, t38)
    var t39 string = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_join____T__string(names__0, "|")
    var inline29 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t39)
    _goml_runtime_core_string_println(inline29)
    var t40 Tuple2_3int_3int = Tuple2_3int_3int{
        _0: 1,
        _1: 2,
    }
    var t41 Tuple2_3int_3int = Tuple2_3int_3int{
        _0: 1,
        _1: 1,
    }
    var t42 Tuple2_3int_3int = Tuple2_3int_3int{
        _0: 0,
        _1: 9,
    }
    var t43 [3]Tuple2_3int_3int = [3]Tuple2_3int_3int{t40, t41, t42}
    var pairs__0 *_goml_vec_Tuple2_3int_3int = func(values [3]Tuple2_3int_3int) *_goml_vec_Tuple2_3int_3int {
        var storage struct {
            vector _goml_vec_Tuple2_3int_3int
            values [3]Tuple2_3int_3int
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t43)
    var t44 closure_env_main_8 = closure_env_main_8{}
    var t45 func(Tuple2_3int_3int, Tuple2_3int_3int) Ordering = func(p0 Tuple2_3int_3int, p1 Tuple2_3int_3int) Ordering {
        return _goml_m_inherent_i_closure__env__main__8_i_closure__env__main__8_i_apply(t44, p0, p1)
    }
    _goml_m_inherent_i_Vec_i_Vec_l_hfde7b787e35869a5469e590f5ee0e2ea_size_c_isize_q_(pairs__0, t45)
    var t46 int
    var inline28 int = vec_len__Vec_16Tuple2_3int_3int(pairs__0)
    t46 = inline28
    var t47 string
    var inline27 string = __goml_builtin_int_to_string(t46)
    t47 = inline27
    var inline25 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t47)
    _goml_runtime_core_string_println(inline25)
    var t48 [3]int = [3]int{7, 8, 9}
    var t49 *_goml_vec_int = func(values [3]int) *_goml_vec_int {
        var storage struct {
            vector _goml_vec_int
            values [3]int
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t48)
    var view__0 []int
    var inline22 int = 0
    var inline23 int = 2
    var inline24 []int = t49.items[inline22:inline23]
    view__0 = inline24
    var t50 bool = _goml_m_inherent_i_Slice_i_Slice_l_T_r__i_contains____T__isize(view__0, 8)
    var t51 string
    var inline21 string = _goml_runtime_core_bool_to_string(t50)
    t51 = inline21
    var inline19 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t51)
    _goml_runtime_core_string_println(inline19)
    var t52 bool = _goml_m_inherent_i_Slice_i_Slice_l_T_r__i_contains____T__isize(view__0, 9)
    var t53 string
    var inline18 string = _goml_runtime_core_bool_to_string(t52)
    t53 = inline18
    var inline16 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t53)
    _goml_runtime_core_string_println(inline16)
    var t54 [2]int = [2]int{1, 2}
    var t55 *_goml_vec_int = func(values [2]int) *_goml_vec_int {
        var storage struct {
            vector _goml_vec_int
            values [2]int
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t54)
    var frozen__0 FrozenVec__isize
    var inline14 *_goml_vec_int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_copy____T__isize(t55)
    var inline15 FrozenVec__isize = FrozenVec__isize{
        values: inline14,
    }
    frozen__0 = inline15
    var t56 bool = _goml_m_inherent_i_FrozenVec_i_FrozenVec_l_T_r__i_contains____T__isize(frozen__0, 1)
    var t57 string
    var inline13 string = _goml_runtime_core_bool_to_string(t56)
    t57 = inline13
    var inline11 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t57)
    _goml_runtime_core_string_println(inline11)
    var t58 bool = _goml_m_inherent_i_FrozenVec_i_FrozenVec_l_T_r__i_contains____T__isize(frozen__0, 3)
    var t59 string
    var inline10 string = _goml_runtime_core_bool_to_string(t58)
    t59 = inline10
    var inline8 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t59)
    _goml_runtime_core_string_println(inline8)
    var empty__0 *_goml_vec_int
    var inline7 *_goml_vec_int = vec_new__Vec_3int()
    empty__0 = inline7
    var t60 string = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_join____T__isize(empty__0, ",")
    var inline5 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t60)
    _goml_runtime_core_string_println(inline5)
    var t61 closure_env_main_9 = closure_env_main_9{}
    var t62 func(int, int) int = func(p0 int, p1 int) int {
        return _goml_m_inherent_i_closure__env__main__9_i_closure__env__main__9_i_apply(t61, p0, p1)
    }
    var t63 Option__isize = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_min__by____T__isize(empty__0, t62)
    var t64 bool
    var inline3 bool = _goml_m_inherent_i_Option_i_Option_l_T_r__i_is__some____T__isize(t63)
    var inline4 bool = !inline3
    t64 = inline4
    var t65 string
    var inline2 string = _goml_runtime_core_bool_to_string(t64)
    t65 = inline2
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t65)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func println__T_string(value__0 string) struct{} {
    var t0 string
    t0 = value__0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_position____T__isize(self__0 *_goml_vec_int, predicate__0 func(int) bool) Option__isize {
    var t0 int
    var inline7 int = vec_len__Vec_3int(self__0)
    t0 = inline7
    var t1 FnIterator__isize
    var inline2 int = 0
    var inline3 *ref_int_x = ref__Ref_3int(inline2)
    var inline4 closure_env_goml_builtin_range_13 = closure_env_goml_builtin_range_13{
        current_0: inline3,
        end_1: t0,
    }
    var inline5 func() Option__isize = func() Option__isize {
        return _goml_m_inherent_i_closure__en_h705ec68e290747cc19a5685005dd16c1_nge__13_i_apply(inline4)
    }
    var inline6 FnIterator__isize = _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__isize(inline5)
    t1 = inline6
    var for_iter0 FnIterator__isize
    for_iter0 = t1
    Loop_loop0:
    for {
        var for_next0 Option__isize
        var inline0 func() Option__isize = for_iter0.next_fn
        var inline1 Option__isize = inline0()
        for_next0 = inline1
        switch for_next0._tag {
        case 0:
            break Loop_loop0
        case 1:
            var x0 int = for_next0._p0
            var t2 int = vec_get__Vec_3int(self__0, x0)
            var t3 bool = predicate__0(t2)
            if t3 {
                var t4 Option__isize = Option__isize{
                    _p0: x0,
                    _tag: 1,
                }
                return t4
            } else {
                continue
            }
        default:
            panic("non-exhaustive match")
        }
    }
    return Option__isize{
        _tag: 0,
    }
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

func _goml_m_inherent_i_isize_i_isize_i_to__string(self__0 int) string {
    var inline0 int64 = int64(int(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_sort__by____T__isize(self__0 *_goml_vec_int, compare__0 func(int, int) int) struct{} {
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_stable__sort__by____T__isize(self__0, compare__0)
    return struct{}{}
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_join____T__isize(self__0 *_goml_vec_int, separator__0 string) string {
    var t0 int
    var inline23 int = vec_len__Vec_3int(self__0)
    t0 = inline23
    var parts__0 *_goml_vec_string
    var inline22 *_goml_vec_string = vec_with_capacity__Vec_6string(t0)
    parts__0 = inline22
    var t1 int
    var inline21 int = vec_len__Vec_3int(self__0)
    t1 = inline21
    var t2 FnIterator__isize
    var inline16 int = 0
    var inline17 *ref_int_x = ref__Ref_3int(inline16)
    var inline18 closure_env_goml_builtin_range_13 = closure_env_goml_builtin_range_13{
        current_0: inline17,
        end_1: t1,
    }
    var inline19 func() Option__isize = func() Option__isize {
        return _goml_m_inherent_i_closure__en_h705ec68e290747cc19a5685005dd16c1_nge__13_i_apply(inline18)
    }
    var inline20 FnIterator__isize = _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__isize(inline19)
    t2 = inline20
    var for_iter0 FnIterator__isize
    for_iter0 = t2
    Loop_loop0:
    for {
        var for_next1 Option__isize
        var inline14 func() Option__isize = for_iter0.next_fn
        var inline15 Option__isize = inline14()
        for_next1 = inline15
        switch for_next1._tag {
        case 0:
            break Loop_loop0
        case 1:
            var x1 int = for_next1._p0
            var t10 int = vec_get__Vec_3int(self__0, x1)
            var t11 string
            var inline13 string = __goml_builtin_int_to_string(t10)
            t11 = inline13
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
    var inline6 closure_env_goml_builtin_range_13 = closure_env_goml_builtin_range_13{
        current_0: inline5,
        end_1: t5,
    }
    var inline7 func() Option__isize = func() Option__isize {
        return _goml_m_inherent_i_closure__en_h705ec68e290747cc19a5685005dd16c1_nge__13_i_apply(inline6)
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

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_dedup____T__isize(self__0 *_goml_vec_int) struct{} {
    var t0 closure_env_inherent_Vec_Vec_T_dedup_T_isize_10 = closure_env_inherent_Vec_Vec_T_dedup_T_isize_10{}
    var t1 func(int, int) bool = func(p0 int, p1 int) bool {
        return _goml_m_inherent_i_closure__en_h5b5a2c4c0a397a14d2b6a2dc409e76d2_ize__10_i_apply(t0, p0, p1)
    }
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_dedup__by____T__isize(self__0, t1)
    return struct{}{}
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_binary__search__by____T__isize(self__0 *_goml_vec_int, compare__0 func(int) int) Option__isize {
    var low__0 int = 0
    var high__0 int
    var inline1 int = vec_len__Vec_3int(self__0)
    high__0 = inline1
    Loop_loop0:
    for {
        var t6 bool = low__0 < high__0
        if t6 {
            var t7 int = high__0 - low__0
            var t8 int = t7 / 2
            var middle__0 int = low__0 + t8
            var t9 int = vec_get__Vec_3int(self__0, middle__0)
            var t10 int = compare__0(t9)
            var t11 bool = t10 < 0
            if t11 {
                var t12 int = middle__0 + 1
                low__0 = t12
                continue
            } else {
                high__0 = middle__0
                continue
            }
        } else {
            break Loop_loop0
        }
    }
    var t0 int
    var inline0 int = vec_len__Vec_3int(self__0)
    t0 = inline0
    var t1 bool = low__0 < t0
    var jp0 bool
    if t1 {
        var t3 int = vec_get__Vec_3int(self__0, low__0)
        var t4 int = compare__0(t3)
        var t5 bool = t4 == 0
        jp0 = t5
    } else {
        jp0 = false
    }
    if jp0 {
        var t2 Option__isize = Option__isize{
            _p0: low__0,
            _tag: 1,
        }
        return t2
    } else {
        return Option__isize{
            _tag: 0,
        }
    }
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_min__by____T__isize(self__0 *_goml_vec_int, compare__0 func(int, int) int) Option__isize {
    var t0 bool
    var inline8 int = vec_len__Vec_3int(self__0)
    var inline9 bool = inline8 == 0
    t0 = inline9
    if t0 {
        return Option__isize{
            _tag: 0,
        }
    } else {
        var best__0 int = vec_get__Vec_3int(self__0, 0)
        var t1 int
        var inline7 int = vec_len__Vec_3int(self__0)
        t1 = inline7
        var t2 FnIterator__isize
        var inline2 int = 1
        var inline3 *ref_int_x = ref__Ref_3int(inline2)
        var inline4 closure_env_goml_builtin_range_13 = closure_env_goml_builtin_range_13{
            current_0: inline3,
            end_1: t1,
        }
        var inline5 func() Option__isize = func() Option__isize {
            return _goml_m_inherent_i_closure__en_h705ec68e290747cc19a5685005dd16c1_nge__13_i_apply(inline4)
        }
        var inline6 FnIterator__isize = _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__isize(inline5)
        t2 = inline6
        var for_iter0 FnIterator__isize
        for_iter0 = t2
        Loop_loop0:
        for {
            var for_next0 Option__isize
            var inline0 func() Option__isize = for_iter0.next_fn
            var inline1 Option__isize = inline0()
            for_next0 = inline1
            switch for_next0._tag {
            case 0:
                break Loop_loop0
            case 1:
                var x0 int = for_next0._p0
                var value__0 int = vec_get__Vec_3int(self__0, x0)
                var t4 int = compare__0(value__0, best__0)
                var t5 bool = t4 < 0
                if t5 {
                    best__0 = value__0
                    continue
                } else {
                    continue
                }
            default:
                panic("non-exhaustive match")
            }
        }
        var t3 Option__isize = Option__isize{
            _p0: best__0,
            _tag: 1,
        }
        return t3
    }
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_max__by____T__isize(self__0 *_goml_vec_int, compare__0 func(int, int) int) Option__isize {
    var t0 bool
    var inline8 int = vec_len__Vec_3int(self__0)
    var inline9 bool = inline8 == 0
    t0 = inline9
    if t0 {
        return Option__isize{
            _tag: 0,
        }
    } else {
        var best__0 int = vec_get__Vec_3int(self__0, 0)
        var t1 int
        var inline7 int = vec_len__Vec_3int(self__0)
        t1 = inline7
        var t2 FnIterator__isize
        var inline2 int = 1
        var inline3 *ref_int_x = ref__Ref_3int(inline2)
        var inline4 closure_env_goml_builtin_range_13 = closure_env_goml_builtin_range_13{
            current_0: inline3,
            end_1: t1,
        }
        var inline5 func() Option__isize = func() Option__isize {
            return _goml_m_inherent_i_closure__en_h705ec68e290747cc19a5685005dd16c1_nge__13_i_apply(inline4)
        }
        var inline6 FnIterator__isize = _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__isize(inline5)
        t2 = inline6
        var for_iter0 FnIterator__isize
        for_iter0 = t2
        Loop_loop0:
        for {
            var for_next0 Option__isize
            var inline0 func() Option__isize = for_iter0.next_fn
            var inline1 Option__isize = inline0()
            for_next0 = inline1
            switch for_next0._tag {
            case 0:
                break Loop_loop0
            case 1:
                var x0 int = for_next0._p0
                var value__0 int = vec_get__Vec_3int(self__0, x0)
                var t4 int = compare__0(value__0, best__0)
                var t5 bool = t4 > 0
                if t5 {
                    best__0 = value__0
                    continue
                } else {
                    continue
                }
            default:
                panic("non-exhaustive match")
            }
        }
        var t3 Option__isize = Option__isize{
            _p0: best__0,
            _tag: 1,
        }
        return t3
    }
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_sort__by__ordering____T__string(self__0 *_goml_vec_string, compare__0 func(string, string) Ordering) struct{} {
    var t0 closure_env_inherent_Vec_Vec_T_sort_by_ordering_T_string_11 = closure_env_inherent_Vec_Vec_T_sort_by_ordering_T_string_11{
        compare_0: compare__0,
    }
    var t1 func(string, string) int = func(p0 string, p1 string) int {
        return _goml_m_inherent_i_closure__en_h8a0fbf92fe59c5b247b9c312b83aa9e5_ing__11_i_apply(t0, p0, p1)
    }
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_stable__sort__by____T__string(self__0, t1)
    return struct{}{}
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
    var inline17 closure_env_goml_builtin_range_13 = closure_env_goml_builtin_range_13{
        current_0: inline16,
        end_1: t1,
    }
    var inline18 func() Option__isize = func() Option__isize {
        return _goml_m_inherent_i_closure__en_h705ec68e290747cc19a5685005dd16c1_nge__13_i_apply(inline17)
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
    var inline6 closure_env_goml_builtin_range_13 = closure_env_goml_builtin_range_13{
        current_0: inline5,
        end_1: t5,
    }
    var inline7 func() Option__isize = func() Option__isize {
        return _goml_m_inherent_i_closure__en_h705ec68e290747cc19a5685005dd16c1_nge__13_i_apply(inline6)
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

func _goml_m_inherent_i_Vec_i_Vec_l_hfde7b787e35869a5469e590f5ee0e2ea_size_c_isize_q_(self__0 *_goml_vec_Tuple2_3int_3int, compare__0 func(Tuple2_3int_3int, Tuple2_3int_3int) Ordering) struct{} {
    var t0 closure_env_inherent_Vec_Vec_T_sort_by_ordering_T_isize_isize_12 = closure_env_inherent_Vec_Vec_T_sort_by_ordering_T_isize_isize_12{
        compare_0: compare__0,
    }
    var t1 func(Tuple2_3int_3int, Tuple2_3int_3int) int = func(p0 Tuple2_3int_3int, p1 Tuple2_3int_3int) int {
        return _goml_m_inherent_i_closure__en_hc4b58ff73aa2ff3a45662a7349486da0_ize__12_i_apply(t0, p0, p1)
    }
    _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_stable__sort__by____T___o_isize_c_isize_q_(self__0, t1)
    return struct{}{}
}

func _goml_m_inherent_i_Slice_i_Slice_l_T_r__i_contains____T__isize(self__0 []int, expected__0 int) bool {
    var index__0 int = 0
    Loop_loop0:
    for {
        var t0 int
        var inline1 int = len(self__0)
        t0 = inline1
        var t1 bool = index__0 < t0
        if t1 {
            var t2 int = self__0[index__0]
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

func _goml_m_inherent_i_FrozenVec_i_FrozenVec_l_T_r__i_contains____T__isize(self__0 FrozenVec__isize, expected__0 int) bool {
    var t0 *_goml_vec_int = self__0.values
    var for_iter0 FnIterator__isize
    var inline3 FnIterator__isize = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_iter____T__isize(t0)
    for_iter0 = inline3
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
            var t1 bool
            var inline0 bool = x0 == expected__0
            t1 = inline0
            if t1 {
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

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_stable__sort__by____T__isize(self__0 *_goml_vec_int, compare__0 func(int, int) int) struct{} {
    var length__0 int
    var inline18 int = vec_len__Vec_3int(self__0)
    length__0 = inline18
    var t0 bool = length__0 < 2
    if t0 {
        return struct{}{}
    } else {
        var buffer__0 *_goml_vec_int
        var inline17 *_goml_vec_int = vec_with_capacity__Vec_3int(length__0)
        buffer__0 = inline17
        var t1 FnIterator__isize
        var inline12 int = 0
        var inline13 *ref_int_x = ref__Ref_3int(inline12)
        var inline14 closure_env_goml_builtin_range_13 = closure_env_goml_builtin_range_13{
            current_0: inline13,
            end_1: length__0,
        }
        var inline15 func() Option__isize = func() Option__isize {
            return _goml_m_inherent_i_closure__en_h705ec68e290747cc19a5685005dd16c1_nge__13_i_apply(inline14)
        }
        var inline16 FnIterator__isize = _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__isize(inline15)
        t1 = inline16
        var for_iter0 FnIterator__isize
        for_iter0 = t1
        Loop_loop0:
        for {
            var for_next1 Option__isize
            var inline10 func() Option__isize = for_iter0.next_fn
            var inline11 Option__isize = inline10()
            for_next1 = inline11
            switch for_next1._tag {
            case 0:
                break Loop_loop0
            case 1:
                var x1 int = for_next1._p0
                var t37 int = vec_get__Vec_3int(self__0, x1)
                vec_push__Vec_3int(buffer__0, t37)
                continue
            default:
                panic("non-exhaustive match")
            }
        }
        var width__0 int = 1
        Loop_loop1:
        for {
            var t2 bool = width__0 < length__0
            if t2 {
                var left__0 int = 0
                Loop_loop2:
                for {
                    var t8 bool = left__0 < length__0
                    if t8 {
                        var t9 int = left__0 + width__0
                        var middle__0 int
                        var inline8 bool = t9 < length__0
                        if inline8 {
                            middle__0 = t9
                        } else {
                            middle__0 = length__0
                        }
                        var t10 int = middle__0 + width__0
                        var right__0 int
                        var inline7 bool = t10 < length__0
                        if inline7 {
                            right__0 = t10
                        } else {
                            right__0 = length__0
                        }
                        var first__0 int = left__0
                        var second__0 int = middle__0
                        var output__0 int = left__0
                        Loop_loop4:
                        for {
                            var t23 bool = first__0 < middle__0
                            var jp1 bool
                            if t23 {
                                var t36 bool = second__0 < right__0
                                jp1 = t36
                            } else {
                                jp1 = false
                            }
                            if jp1 {
                                var t24 int = vec_get__Vec_3int(self__0, first__0)
                                var t25 int = vec_get__Vec_3int(self__0, second__0)
                                var t26 int = compare__0(t24, t25)
                                var t27 bool = t26 <= 0
                                if t27 {
                                    var index2 int = output__0
                                    vec_get__Vec_3int(buffer__0, index2)
                                    var value3 int = vec_get__Vec_3int(self__0, first__0)
                                    vec_set__Vec_3int(buffer__0, index2, value3)
                                    var compound_old5 int = first__0
                                    var compound_value5 int = 1
                                    var t31 int = compound_old5 + compound_value5
                                    first__0 = t31
                                } else {
                                    var index3 int = output__0
                                    vec_get__Vec_3int(buffer__0, index3)
                                    var value4 int = vec_get__Vec_3int(self__0, second__0)
                                    vec_set__Vec_3int(buffer__0, index3, value4)
                                    var compound_old6 int = second__0
                                    var compound_value6 int = 1
                                    var t34 int = compound_old6 + compound_value6
                                    second__0 = t34
                                }
                                var compound_old4 int = output__0
                                var compound_value4 int = 1
                                var t28 int = compound_old4 + compound_value4
                                output__0 = t28
                                continue
                            } else {
                                break Loop_loop4
                            }
                        }
                        Loop_loop5:
                        for {
                            var t17 bool = first__0 < middle__0
                            if t17 {
                                var index1 int = output__0
                                vec_get__Vec_3int(buffer__0, index1)
                                var value2 int = vec_get__Vec_3int(self__0, first__0)
                                vec_set__Vec_3int(buffer__0, index1, value2)
                                var compound_old2 int = first__0
                                var compound_value2 int = 1
                                var t19 int = compound_old2 + compound_value2
                                first__0 = t19
                                var compound_old3 int = output__0
                                var compound_value3 int = 1
                                var t21 int = compound_old3 + compound_value3
                                output__0 = t21
                                continue
                            } else {
                                break Loop_loop5
                            }
                        }
                        Loop_loop6:
                        for {
                            var t11 bool = second__0 < right__0
                            if t11 {
                                var index0 int = output__0
                                vec_get__Vec_3int(buffer__0, index0)
                                var value1 int = vec_get__Vec_3int(self__0, second__0)
                                vec_set__Vec_3int(buffer__0, index0, value1)
                                var compound_old0 int = second__0
                                var compound_value0 int = 1
                                var t13 int = compound_old0 + compound_value0
                                second__0 = t13
                                var compound_old1 int = output__0
                                var compound_value1 int = 1
                                var t15 int = compound_old1 + compound_value1
                                output__0 = t15
                                continue
                            } else {
                                break Loop_loop6
                            }
                        }
                        left__0 = right__0
                        continue
                    } else {
                        break Loop_loop2
                    }
                }
                var t3 FnIterator__isize
                var inline2 int = 0
                var inline3 *ref_int_x = ref__Ref_3int(inline2)
                var inline4 closure_env_goml_builtin_range_13 = closure_env_goml_builtin_range_13{
                    current_0: inline3,
                    end_1: length__0,
                }
                var inline5 func() Option__isize = func() Option__isize {
                    return _goml_m_inherent_i_closure__en_h705ec68e290747cc19a5685005dd16c1_nge__13_i_apply(inline4)
                }
                var inline6 FnIterator__isize = _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__isize(inline5)
                t3 = inline6
                var for_iter1 FnIterator__isize
                for_iter1 = t3
                Loop_loop3:
                for {
                    var for_next0 Option__isize
                    var inline0 func() Option__isize = for_iter1.next_fn
                    var inline1 Option__isize = inline0()
                    for_next0 = inline1
                    switch for_next0._tag {
                    case 0:
                        break Loop_loop3
                    case 1:
                        var x0 int = for_next0._p0
                        vec_get__Vec_3int(self__0, x0)
                        var value0 int = vec_get__Vec_3int(buffer__0, x0)
                        vec_set__Vec_3int(self__0, x0, value0)
                        continue
                    default:
                        panic("non-exhaustive match")
                    }
                }
                var t4 int = length__0 / 2
                var t5 bool = width__0 > t4
                var jp0 int
                if t5 {
                    jp0 = length__0
                } else {
                    var t6 int = width__0 * 2
                    jp0 = t6
                }
                width__0 = jp0
                continue
            } else {
                break Loop_loop1
            }
        }
        return struct{}{}
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

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_dedup__by____T__isize(self__0 *_goml_vec_int, equal__0 func(int, int) bool) struct{} {
    var t0 int
    var inline9 int = vec_len__Vec_3int(self__0)
    t0 = inline9
    var t1 bool = t0 < 2
    if t1 {
        return struct{}{}
    } else {
        var output__0 int = 1
        var t2 int
        var inline8 int = vec_len__Vec_3int(self__0)
        t2 = inline8
        var t3 FnIterator__isize
        var inline3 int = 1
        var inline4 *ref_int_x = ref__Ref_3int(inline3)
        var inline5 closure_env_goml_builtin_range_13 = closure_env_goml_builtin_range_13{
            current_0: inline4,
            end_1: t2,
        }
        var inline6 func() Option__isize = func() Option__isize {
            return _goml_m_inherent_i_closure__en_h705ec68e290747cc19a5685005dd16c1_nge__13_i_apply(inline5)
        }
        var inline7 FnIterator__isize = _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__isize(inline6)
        t3 = inline7
        var for_iter0 FnIterator__isize
        for_iter0 = t3
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
                var value__0 int = vec_get__Vec_3int(self__0, x0)
                var t4 int = output__0 - 1
                var t5 int = vec_get__Vec_3int(self__0, t4)
                var t6 bool = equal__0(t5, value__0)
                var t7 bool = !t6
                if t7 {
                    var index0 int = output__0
                    vec_get__Vec_3int(self__0, index0)
                    vec_set__Vec_3int(self__0, index0, value__0)
                    var compound_old0 int = output__0
                    var compound_value0 int = 1
                    var t9 int = compound_old0 + compound_value0
                    output__0 = t9
                    continue
                } else {
                    continue
                }
            default:
                panic("non-exhaustive match")
            }
        }
        vec_truncate__Vec_3int(self__0, output__0)
        return struct{}{}
    }
}

func _goml_m_trait__impl_i_std_p_cmp_p_Ord_i_string_i_cmp(self__0 string, other__0 string) Ordering {
    var t0 bool = self__0 < other__0
    if t0 {
        return Less
    } else {
        var t1 bool = self__0 > other__0
        if t1 {
            return Greater
        } else {
            return Equal
        }
    }
}

func _goml_m_trait__impl_i_std_p_cmp_p_Ord_i_isize_i_cmp(self__0 int, other__0 int) Ordering {
    var t0 bool = self__0 < other__0
    if t0 {
        return Less
    } else {
        var t1 bool = self__0 > other__0
        if t1 {
            return Greater
        } else {
            return Equal
        }
    }
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_copy____T__isize(self__0 *_goml_vec_int) *_goml_vec_int {
    var t0 int
    var inline3 int = vec_len__Vec_3int(self__0)
    t0 = inline3
    var result__0 *_goml_vec_int
    var inline2 *_goml_vec_int = vec_with_capacity__Vec_3int(t0)
    result__0 = inline2
    var index__0 int = 0
    Loop_loop0:
    for {
        var t1 int
        var inline1 int = vec_len__Vec_3int(self__0)
        t1 = inline1
        var t2 bool = index__0 < t1
        if t2 {
            var t3 int = vec_get__Vec_3int(self__0, index__0)
            vec_push__Vec_3int(result__0, t3)
            var compound_old0 int = index__0
            var compound_value0 int = 1
            var t4 int = compound_old0 + compound_value0
            index__0 = t4
            continue
        } else {
            break Loop_loop0
        }
    }
    return result__0
}

func _goml_m_inherent_i_Option_i_Option_l_T_r__i_is__some____T__isize(self__0 Option__isize) bool {
    switch self__0._tag {
    case 0:
        return false
    case 1:
        return true
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__isize(next_fn__0 func() Option__isize) FnIterator__isize {
    var t0 FnIterator__isize = FnIterator__isize{
        next_fn: next_fn__0,
    }
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

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_stable__sort__by____T__string(self__0 *_goml_vec_string, compare__0 func(string, string) int) struct{} {
    var length__0 int
    var inline18 int = vec_len__Vec_6string(self__0)
    length__0 = inline18
    var t0 bool = length__0 < 2
    if t0 {
        return struct{}{}
    } else {
        var buffer__0 *_goml_vec_string
        var inline17 *_goml_vec_string = vec_with_capacity__Vec_6string(length__0)
        buffer__0 = inline17
        var t1 FnIterator__isize
        var inline12 int = 0
        var inline13 *ref_int_x = ref__Ref_3int(inline12)
        var inline14 closure_env_goml_builtin_range_13 = closure_env_goml_builtin_range_13{
            current_0: inline13,
            end_1: length__0,
        }
        var inline15 func() Option__isize = func() Option__isize {
            return _goml_m_inherent_i_closure__en_h705ec68e290747cc19a5685005dd16c1_nge__13_i_apply(inline14)
        }
        var inline16 FnIterator__isize = _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__isize(inline15)
        t1 = inline16
        var for_iter0 FnIterator__isize
        for_iter0 = t1
        Loop_loop0:
        for {
            var for_next1 Option__isize
            var inline10 func() Option__isize = for_iter0.next_fn
            var inline11 Option__isize = inline10()
            for_next1 = inline11
            switch for_next1._tag {
            case 0:
                break Loop_loop0
            case 1:
                var x1 int = for_next1._p0
                var t37 string = vec_get__Vec_6string(self__0, x1)
                vec_push__Vec_6string(buffer__0, t37)
                continue
            default:
                panic("non-exhaustive match")
            }
        }
        var width__0 int = 1
        Loop_loop1:
        for {
            var t2 bool = width__0 < length__0
            if t2 {
                var left__0 int = 0
                Loop_loop2:
                for {
                    var t8 bool = left__0 < length__0
                    if t8 {
                        var t9 int = left__0 + width__0
                        var middle__0 int
                        var inline8 bool = t9 < length__0
                        if inline8 {
                            middle__0 = t9
                        } else {
                            middle__0 = length__0
                        }
                        var t10 int = middle__0 + width__0
                        var right__0 int
                        var inline7 bool = t10 < length__0
                        if inline7 {
                            right__0 = t10
                        } else {
                            right__0 = length__0
                        }
                        var first__0 int = left__0
                        var second__0 int = middle__0
                        var output__0 int = left__0
                        Loop_loop4:
                        for {
                            var t23 bool = first__0 < middle__0
                            var jp1 bool
                            if t23 {
                                var t36 bool = second__0 < right__0
                                jp1 = t36
                            } else {
                                jp1 = false
                            }
                            if jp1 {
                                var t24 string = vec_get__Vec_6string(self__0, first__0)
                                var t25 string = vec_get__Vec_6string(self__0, second__0)
                                var t26 int = compare__0(t24, t25)
                                var t27 bool = t26 <= 0
                                if t27 {
                                    var index2 int = output__0
                                    vec_get__Vec_6string(buffer__0, index2)
                                    var value3 string = vec_get__Vec_6string(self__0, first__0)
                                    vec_set__Vec_6string(buffer__0, index2, value3)
                                    var compound_old5 int = first__0
                                    var compound_value5 int = 1
                                    var t31 int = compound_old5 + compound_value5
                                    first__0 = t31
                                } else {
                                    var index3 int = output__0
                                    vec_get__Vec_6string(buffer__0, index3)
                                    var value4 string = vec_get__Vec_6string(self__0, second__0)
                                    vec_set__Vec_6string(buffer__0, index3, value4)
                                    var compound_old6 int = second__0
                                    var compound_value6 int = 1
                                    var t34 int = compound_old6 + compound_value6
                                    second__0 = t34
                                }
                                var compound_old4 int = output__0
                                var compound_value4 int = 1
                                var t28 int = compound_old4 + compound_value4
                                output__0 = t28
                                continue
                            } else {
                                break Loop_loop4
                            }
                        }
                        Loop_loop5:
                        for {
                            var t17 bool = first__0 < middle__0
                            if t17 {
                                var index1 int = output__0
                                vec_get__Vec_6string(buffer__0, index1)
                                var value2 string = vec_get__Vec_6string(self__0, first__0)
                                vec_set__Vec_6string(buffer__0, index1, value2)
                                var compound_old2 int = first__0
                                var compound_value2 int = 1
                                var t19 int = compound_old2 + compound_value2
                                first__0 = t19
                                var compound_old3 int = output__0
                                var compound_value3 int = 1
                                var t21 int = compound_old3 + compound_value3
                                output__0 = t21
                                continue
                            } else {
                                break Loop_loop5
                            }
                        }
                        Loop_loop6:
                        for {
                            var t11 bool = second__0 < right__0
                            if t11 {
                                var index0 int = output__0
                                vec_get__Vec_6string(buffer__0, index0)
                                var value1 string = vec_get__Vec_6string(self__0, second__0)
                                vec_set__Vec_6string(buffer__0, index0, value1)
                                var compound_old0 int = second__0
                                var compound_value0 int = 1
                                var t13 int = compound_old0 + compound_value0
                                second__0 = t13
                                var compound_old1 int = output__0
                                var compound_value1 int = 1
                                var t15 int = compound_old1 + compound_value1
                                output__0 = t15
                                continue
                            } else {
                                break Loop_loop6
                            }
                        }
                        left__0 = right__0
                        continue
                    } else {
                        break Loop_loop2
                    }
                }
                var t3 FnIterator__isize
                var inline2 int = 0
                var inline3 *ref_int_x = ref__Ref_3int(inline2)
                var inline4 closure_env_goml_builtin_range_13 = closure_env_goml_builtin_range_13{
                    current_0: inline3,
                    end_1: length__0,
                }
                var inline5 func() Option__isize = func() Option__isize {
                    return _goml_m_inherent_i_closure__en_h705ec68e290747cc19a5685005dd16c1_nge__13_i_apply(inline4)
                }
                var inline6 FnIterator__isize = _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__isize(inline5)
                t3 = inline6
                var for_iter1 FnIterator__isize
                for_iter1 = t3
                Loop_loop3:
                for {
                    var for_next0 Option__isize
                    var inline0 func() Option__isize = for_iter1.next_fn
                    var inline1 Option__isize = inline0()
                    for_next0 = inline1
                    switch for_next0._tag {
                    case 0:
                        break Loop_loop3
                    case 1:
                        var x0 int = for_next0._p0
                        vec_get__Vec_6string(self__0, x0)
                        var value0 string = vec_get__Vec_6string(buffer__0, x0)
                        vec_set__Vec_6string(self__0, x0, value0)
                        continue
                    default:
                        panic("non-exhaustive match")
                    }
                }
                var t4 int = length__0 / 2
                var t5 bool = width__0 > t4
                var jp0 int
                if t5 {
                    jp0 = length__0
                } else {
                    var t6 int = width__0 * 2
                    jp0 = t6
                }
                width__0 = jp0
                continue
            } else {
                break Loop_loop1
            }
        }
        return struct{}{}
    }
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_stable__sort__by____T___o_isize_c_isize_q_(self__0 *_goml_vec_Tuple2_3int_3int, compare__0 func(Tuple2_3int_3int, Tuple2_3int_3int) int) struct{} {
    var length__0 int
    var inline18 int = vec_len__Vec_16Tuple2_3int_3int(self__0)
    length__0 = inline18
    var t0 bool = length__0 < 2
    if t0 {
        return struct{}{}
    } else {
        var buffer__0 *_goml_vec_Tuple2_3int_3int
        var inline17 *_goml_vec_Tuple2_3int_3int = vec_with_capacity__Vec_16Tuple2_3int_3int(length__0)
        buffer__0 = inline17
        var t1 FnIterator__isize
        var inline12 int = 0
        var inline13 *ref_int_x = ref__Ref_3int(inline12)
        var inline14 closure_env_goml_builtin_range_13 = closure_env_goml_builtin_range_13{
            current_0: inline13,
            end_1: length__0,
        }
        var inline15 func() Option__isize = func() Option__isize {
            return _goml_m_inherent_i_closure__en_h705ec68e290747cc19a5685005dd16c1_nge__13_i_apply(inline14)
        }
        var inline16 FnIterator__isize = _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__isize(inline15)
        t1 = inline16
        var for_iter0 FnIterator__isize
        for_iter0 = t1
        Loop_loop0:
        for {
            var for_next1 Option__isize
            var inline10 func() Option__isize = for_iter0.next_fn
            var inline11 Option__isize = inline10()
            for_next1 = inline11
            switch for_next1._tag {
            case 0:
                break Loop_loop0
            case 1:
                var x1 int = for_next1._p0
                var t37 Tuple2_3int_3int = vec_get__Vec_16Tuple2_3int_3int(self__0, x1)
                vec_push__Vec_16Tuple2_3int_3int(buffer__0, t37)
                continue
            default:
                panic("non-exhaustive match")
            }
        }
        var width__0 int = 1
        Loop_loop1:
        for {
            var t2 bool = width__0 < length__0
            if t2 {
                var left__0 int = 0
                Loop_loop2:
                for {
                    var t8 bool = left__0 < length__0
                    if t8 {
                        var t9 int = left__0 + width__0
                        var middle__0 int
                        var inline8 bool = t9 < length__0
                        if inline8 {
                            middle__0 = t9
                        } else {
                            middle__0 = length__0
                        }
                        var t10 int = middle__0 + width__0
                        var right__0 int
                        var inline7 bool = t10 < length__0
                        if inline7 {
                            right__0 = t10
                        } else {
                            right__0 = length__0
                        }
                        var first__0 int = left__0
                        var second__0 int = middle__0
                        var output__0 int = left__0
                        Loop_loop4:
                        for {
                            var t23 bool = first__0 < middle__0
                            var jp1 bool
                            if t23 {
                                var t36 bool = second__0 < right__0
                                jp1 = t36
                            } else {
                                jp1 = false
                            }
                            if jp1 {
                                var t24 Tuple2_3int_3int = vec_get__Vec_16Tuple2_3int_3int(self__0, first__0)
                                var t25 Tuple2_3int_3int = vec_get__Vec_16Tuple2_3int_3int(self__0, second__0)
                                var t26 int = compare__0(t24, t25)
                                var t27 bool = t26 <= 0
                                if t27 {
                                    var index2 int = output__0
                                    vec_get__Vec_16Tuple2_3int_3int(buffer__0, index2)
                                    var value3 Tuple2_3int_3int = vec_get__Vec_16Tuple2_3int_3int(self__0, first__0)
                                    vec_set__Vec_16Tuple2_3int_3int(buffer__0, index2, value3)
                                    var compound_old5 int = first__0
                                    var compound_value5 int = 1
                                    var t31 int = compound_old5 + compound_value5
                                    first__0 = t31
                                } else {
                                    var index3 int = output__0
                                    vec_get__Vec_16Tuple2_3int_3int(buffer__0, index3)
                                    var value4 Tuple2_3int_3int = vec_get__Vec_16Tuple2_3int_3int(self__0, second__0)
                                    vec_set__Vec_16Tuple2_3int_3int(buffer__0, index3, value4)
                                    var compound_old6 int = second__0
                                    var compound_value6 int = 1
                                    var t34 int = compound_old6 + compound_value6
                                    second__0 = t34
                                }
                                var compound_old4 int = output__0
                                var compound_value4 int = 1
                                var t28 int = compound_old4 + compound_value4
                                output__0 = t28
                                continue
                            } else {
                                break Loop_loop4
                            }
                        }
                        Loop_loop5:
                        for {
                            var t17 bool = first__0 < middle__0
                            if t17 {
                                var index1 int = output__0
                                vec_get__Vec_16Tuple2_3int_3int(buffer__0, index1)
                                var value2 Tuple2_3int_3int = vec_get__Vec_16Tuple2_3int_3int(self__0, first__0)
                                vec_set__Vec_16Tuple2_3int_3int(buffer__0, index1, value2)
                                var compound_old2 int = first__0
                                var compound_value2 int = 1
                                var t19 int = compound_old2 + compound_value2
                                first__0 = t19
                                var compound_old3 int = output__0
                                var compound_value3 int = 1
                                var t21 int = compound_old3 + compound_value3
                                output__0 = t21
                                continue
                            } else {
                                break Loop_loop5
                            }
                        }
                        Loop_loop6:
                        for {
                            var t11 bool = second__0 < right__0
                            if t11 {
                                var index0 int = output__0
                                vec_get__Vec_16Tuple2_3int_3int(buffer__0, index0)
                                var value1 Tuple2_3int_3int = vec_get__Vec_16Tuple2_3int_3int(self__0, second__0)
                                vec_set__Vec_16Tuple2_3int_3int(buffer__0, index0, value1)
                                var compound_old0 int = second__0
                                var compound_value0 int = 1
                                var t13 int = compound_old0 + compound_value0
                                second__0 = t13
                                var compound_old1 int = output__0
                                var compound_value1 int = 1
                                var t15 int = compound_old1 + compound_value1
                                output__0 = t15
                                continue
                            } else {
                                break Loop_loop6
                            }
                        }
                        left__0 = right__0
                        continue
                    } else {
                        break Loop_loop2
                    }
                }
                var t3 FnIterator__isize
                var inline2 int = 0
                var inline3 *ref_int_x = ref__Ref_3int(inline2)
                var inline4 closure_env_goml_builtin_range_13 = closure_env_goml_builtin_range_13{
                    current_0: inline3,
                    end_1: length__0,
                }
                var inline5 func() Option__isize = func() Option__isize {
                    return _goml_m_inherent_i_closure__en_h705ec68e290747cc19a5685005dd16c1_nge__13_i_apply(inline4)
                }
                var inline6 FnIterator__isize = _goml_m_inherent_i_FnIterator_i_FnIterator_l_T_r__i_from__fn____T__isize(inline5)
                t3 = inline6
                var for_iter1 FnIterator__isize
                for_iter1 = t3
                Loop_loop3:
                for {
                    var for_next0 Option__isize
                    var inline0 func() Option__isize = for_iter1.next_fn
                    var inline1 Option__isize = inline0()
                    for_next0 = inline1
                    switch for_next0._tag {
                    case 0:
                        break Loop_loop3
                    case 1:
                        var x0 int = for_next0._p0
                        vec_get__Vec_16Tuple2_3int_3int(self__0, x0)
                        var value0 Tuple2_3int_3int = vec_get__Vec_16Tuple2_3int_3int(buffer__0, x0)
                        vec_set__Vec_16Tuple2_3int_3int(self__0, x0, value0)
                        continue
                    default:
                        panic("non-exhaustive match")
                    }
                }
                var t4 int = length__0 / 2
                var t5 bool = width__0 > t4
                var jp0 int
                if t5 {
                    jp0 = length__0
                } else {
                    var t6 int = width__0 * 2
                    jp0 = t6
                }
                width__0 = jp0
                continue
            } else {
                break Loop_loop1
            }
        }
        return struct{}{}
    }
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_iter____T__isize(self__0 *_goml_vec_int) FnIterator__isize {
    var index__0 *ref_int_x = ref__Ref_3int(0)
    var len__0 int
    var inline1 int = vec_len__Vec_3int(self__0)
    len__0 = inline1
    var t0 closure_env_inherent_Vec_Vec_T_iter_T_isize_14 = closure_env_inherent_Vec_Vec_T_iter_T_isize_14{
        index_0: index__0,
        len_1: len__0,
        self_2: self__0,
    }
    var t1 func() Option__isize = func() Option__isize {
        return _goml_m_inherent_i_closure__en_h9c0087c1b8c373cdd04a2f248cb43de6_ize__14_i_apply(t0)
    }
    var inline0 FnIterator__isize = FnIterator__isize{
        next_fn: t1,
    }
    return inline0
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

func _goml_m_inherent_i_closure__env__main__0_i_closure__env__main__0_i_apply(env0 closure_env_main_0, value__0 int) bool {
    var t0 bool = value__0 == 5
    return t0
}

func _goml_m_inherent_i_closure__env__main__1_i_closure__env__main__1_i_apply(env0 closure_env_main_1, value__0 int) bool {
    var t0 bool = value__0 == 99
    return t0
}

func _goml_m_inherent_i_closure__env__main__2_i_closure__env__main__2_i_apply(env0 closure_env_main_2, left__0 int, right__0 int) int {
    var t0 int = left__0 - right__0
    return t0
}

func _goml_m_inherent_i_closure__env__main__3_i_closure__env__main__3_i_apply(env0 closure_env_main_3, value__0 int) int {
    var t0 int = value__0 - 4
    return t0
}

func _goml_m_inherent_i_closure__env__main__4_i_closure__env__main__4_i_apply(env0 closure_env_main_4, value__0 int) int {
    var t0 int = value__0 - 3
    return t0
}

func _goml_m_inherent_i_closure__env__main__5_i_closure__env__main__5_i_apply(env0 closure_env_main_5, left__0 int, right__0 int) int {
    var t0 int = left__0 - right__0
    return t0
}

func _goml_m_inherent_i_closure__env__main__6_i_closure__env__main__6_i_apply(env0 closure_env_main_6, left__0 int, right__0 int) int {
    var t0 int = left__0 - right__0
    return t0
}

func _goml_m_inherent_i_closure__env__main__7_i_closure__env__main__7_i_apply(env0 closure_env_main_7, left__0 string, right__0 string) Ordering {
    var inline0 Ordering = _goml_m_trait__impl_i_std_p_cmp_p_Ord_i_string_i_cmp(left__0, right__0)
    return inline0
}

func _goml_m_inherent_i_closure__env__main__8_i_closure__env__main__8_i_apply(env0 closure_env_main_8, left__0 Tuple2_3int_3int, right__0 Tuple2_3int_3int) Ordering {
    var t0 int = left__0._0
    var t1 int = right__0._0
    var inline0 Ordering = _goml_m_trait__impl_i_std_p_cmp_p_Ord_i_isize_i_cmp(t0, t1)
    return inline0
}

func _goml_m_inherent_i_closure__env__main__9_i_closure__env__main__9_i_apply(env0 closure_env_main_9, left__0 int, right__0 int) int {
    var t0 int = left__0 - right__0
    return t0
}

func _goml_m_inherent_i_closure__en_h5b5a2c4c0a397a14d2b6a2dc409e76d2_ize__10_i_apply(env0 closure_env_inherent_Vec_Vec_T_dedup_T_isize_10, left__0 int, right__0 int) bool {
    var inline0 bool = left__0 == right__0
    return inline0
}

func _goml_m_inherent_i_closure__en_h8a0fbf92fe59c5b247b9c312b83aa9e5_ing__11_i_apply(env0 closure_env_inherent_Vec_Vec_T_sort_by_ordering_T_string_11, left__0 string, right__0 string) int {
    var compare__0 func(string, string) Ordering = env0.compare_0
    var t0 Ordering = compare__0(left__0, right__0)
    switch t0 {
    case Less:
        return -1
    case Equal:
        return 0
    case Greater:
        return 1
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_inherent_i_closure__en_hc4b58ff73aa2ff3a45662a7349486da0_ize__12_i_apply(env0 closure_env_inherent_Vec_Vec_T_sort_by_ordering_T_isize_isize_12, left__0 Tuple2_3int_3int, right__0 Tuple2_3int_3int) int {
    var compare__0 func(Tuple2_3int_3int, Tuple2_3int_3int) Ordering = env0.compare_0
    var t0 Ordering = compare__0(left__0, right__0)
    switch t0 {
    case Less:
        return -1
    case Equal:
        return 0
    case Greater:
        return 1
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_inherent_i_closure__en_h705ec68e290747cc19a5685005dd16c1_nge__13_i_apply(env0 closure_env_goml_builtin_range_13) Option__isize {
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

func _goml_m_inherent_i_closure__en_h9c0087c1b8c373cdd04a2f248cb43de6_ize__14_i_apply(env0 closure_env_inherent_Vec_Vec_T_iter_T_isize_14) Option__isize {
    var index__0 *ref_int_x = env0.index_0
    var len__0 int = env0.len_1
    var self__0 *_goml_vec_int = env0.self_2
    var current__0 int = ref_get__Ref_3int(index__0)
    var t0 bool = current__0 < len__0
    if t0 {
        var value__0 int = vec_get__Vec_3int(self__0, current__0)
        var t1 int = current__0 + 1
        ref_set__Ref_3int(index__0, t1)
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
