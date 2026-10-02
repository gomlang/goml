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

type _goml_vec_int struct {
    items []int
}

func vec_get__Vec_3int(vec *_goml_vec_int, index int) int {
    return vec.items[index]
}

func vec_len__Vec_3int(vec *_goml_vec_int) int {
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

type Ordering uint8

type _goml_m_Option____MutSlice_l_isize_r_ struct {
    _p0 []int
    _tag uint8
}

type _goml_m_Option____Slice_l_isize_r_ struct {
    _p0 []int
    _tag uint8
}

type Option__isize struct {
    _p0 int
    _tag uint8
}

func main0() struct{} {
    var t0 [5]int = [5]int{1, 2, 3, 4, 5}
    var values__0 *_goml_vec_int = func(values [5]int) *_goml_vec_int {
        var storage struct {
            vector _goml_vec_int
            values [5]int
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t0)
    var view__0 []int
    var inline30 int = 1
    var inline31 int = 5
    var inline32 []int = values__0.items[inline30:inline31]
    view__0 = inline32
    var index0 int = 0
    _ = view__0[index0]
    var value0 int = 20
    func(p0 []int, p1 int, p2 int) struct{} {
        p0[p1] = p2
        return struct{}{}
    }(view__0, index0, value0)
    var t2 bool = _goml_m_inherent_i_MutSlice_i_MutSlice_l_T_r__i_set__checked____T__isize(view__0, 9, 30)
    var inline28 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t2)
    _goml_runtime_core_string_println(inline28)
    var t3 bool = _goml_m_inherent_i_MutSlice_i_MutSlice_l_T_r__i_set__checked____T__isize(view__0, 1, 30)
    var inline26 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t3)
    _goml_runtime_core_string_println(inline26)
    var t4 bool = _goml_m_inherent_i_MutSlice_i_MutSlice_l_T_r__i_copy__within____T__isize(view__0, 0, 3, 1)
    var inline24 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t4)
    _goml_runtime_core_string_println(inline24)
    var t5 [2]int = [2]int{7, 8}
    var t6 *_goml_vec_int = func(values [2]int) *_goml_vec_int {
        var storage struct {
            vector _goml_vec_int
            values [2]int
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t5)
    var t7 []int
    var inline22 int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_len____T__isize(t6)
    var inline23 []int = t6.items[0:inline22]
    t7 = inline23
    var t8 bool
    var inline20 int = 0
    var inline21 bool = func(p0 []int, p1 int, p2 []int) bool {
        if p1 < 0 || len(p2) > len(p0) - p1 {
            return false
        }
        copy(p0[p1:p1 + len(p2)], p2)
        return true
    }(view__0, inline20, t7)
    t8 = inline21
    var inline18 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t8)
    _goml_runtime_core_string_println(inline18)
    var mtmp0 _goml_m_Option____MutSlice_l_isize_r_ = _goml_m_inherent_i_MutSlice_i_MutSlice_l_T_r__i_sub__checked____T__isize(view__0, 1, 3)
    switch mtmp0._tag {
    case 0:
    case 1:
        var x0 []int = mtmp0._p0
        _goml_m_inherent_i_MutSlice_i_MutSlice_l_T_r__i_fill____T__isize(x0, 6)
    default:
        panic("non-exhaustive match")
    }
    var for_limit0 int = vec_len__Vec_3int(values__0)
    var for_index0 int = 0
    Loop_loop0:
    for {
        var t14 bool = for_index0 < for_limit0
        if t14 {
            var for_item0 int = vec_get__Vec_3int(values__0, for_index0)
            var t15 int = for_index0 + 1
            for_index0 = t15
            var inline16 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(for_item0)
            _goml_runtime_core_string_println(inline16)
            continue
        } else {
            break Loop_loop0
        }
    }
    var t9 _goml_m_Option____Slice_l_isize_r_ = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_slice__checked____T__isize(values__0, -1, 2)
    var t10 bool
    var inline14 bool = _goml_m_inherent_i_Option_i_Option_l_T_r__i_is__some____T__Slice_l_isize_r_(t9)
    var inline15 bool = !inline14
    t10 = inline15
    var inline12 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t10)
    _goml_runtime_core_string_println(inline12)
    var t11 []int
    var inline11 []int = view__0
    t11 = inline11
    var t12 Option__isize
    var inline4 int = 10
    var inline5 bool = inline4 < 0
    var inline6 bool
    if inline5 {
        inline6 = true
    } else {
        var inline9 int = _goml_m_inherent_i_Slice_i_Slice_l_T_r__i_len____T__isize(t11)
        var inline10 bool = inline4 >= inline9
        inline6 = inline10
    }
    if inline6 {
        t12 = Option__isize{
            _tag: 0,
        }
        var t13 bool
        var inline2 bool = _goml_m_inherent_i_Option_i_Option_l_T_r__i_is__some____T__isize(t12)
        var inline3 bool = !inline2
        t13 = inline3
        var inline0 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t13)
        _goml_runtime_core_string_println(inline0)
        return struct{}{}
    } else {
        var inline7 int = t11[inline4]
        var inline8 Option__isize = Option__isize{
            _p0: inline7,
            _tag: 1,
        }
        t12 = inline8
        var t13 bool
        var inline2 bool = _goml_m_inherent_i_Option_i_Option_l_T_r__i_is__some____T__isize(t12)
        var inline3 bool = !inline2
        t13 = inline3
        var inline0 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t13)
        _goml_runtime_core_string_println(inline0)
        return struct{}{}
    }
}

func _goml_m_inherent_i_MutSlice_i_MutSlice_l_T_r__i_set__checked____T__isize(self__0 []int, index__0 int, value__0 int) bool {
    var t0 bool = index__0 < 0
    var jp0 bool
    if t0 {
        jp0 = true
    } else {
        var t2 int
        var inline0 int = len(self__0)
        t2 = inline0
        var t3 bool = index__0 >= t2
        jp0 = t3
    }
    if jp0 {
        return false
    } else {
        _ = self__0[index__0]
        func(p0 []int, p1 int, p2 int) struct{} {
            p0[p1] = p2
            return struct{}{}
        }(self__0, index__0, value__0)
        return true
    }
}

func _goml_m_inherent_i_MutSlice_i_MutSlice_l_T_r__i_copy__within____T__isize(self__0 []int, source_start__0 int, source_end__0 int, dest_start__0 int) bool {
    var t0 bool = source_start__0 < 0
    var jp0 bool
    if t0 {
        jp0 = true
    } else {
        var t5 bool = source_end__0 < source_start__0
        jp0 = t5
    }
    var jp1 bool
    if jp0 {
        jp1 = true
    } else {
        var t3 int
        var inline3 int = len(self__0)
        t3 = inline3
        var t4 bool = source_end__0 > t3
        jp1 = t4
    }
    if jp1 {
        return false
    } else {
        var t1 []int
        var inline2 []int = self__0
        t1 = inline2
        var t2 []int
        var inline1 []int = t1[source_start__0:source_end__0]
        t2 = inline1
        var inline0 bool = func(p0 []int, p1 int, p2 []int) bool {
            if p1 < 0 || len(p2) > len(p0) - p1 {
                return false
            }
            copy(p0[p1:p1 + len(p2)], p2)
            return true
        }(self__0, dest_start__0, t2)
        return inline0
    }
}

func _goml_m_inherent_i_MutSlice_i_MutSlice_l_T_r__i_sub__checked____T__isize(self__0 []int, start__0 int, end__0 int) _goml_m_Option____MutSlice_l_isize_r_ {
    var t0 bool = start__0 < 0
    var jp0 bool
    if t0 {
        jp0 = true
    } else {
        var t5 bool = end__0 < start__0
        jp0 = t5
    }
    var jp1 bool
    if jp0 {
        jp1 = true
    } else {
        var t3 int
        var inline1 int = len(self__0)
        t3 = inline1
        var t4 bool = end__0 > t3
        jp1 = t4
    }
    if jp1 {
        return _goml_m_Option____MutSlice_l_isize_r_{
            _tag: 0,
        }
    } else {
        var t1 []int
        var inline0 []int = self__0[start__0:end__0]
        t1 = inline0
        var t2 _goml_m_Option____MutSlice_l_isize_r_ = _goml_m_Option____MutSlice_l_isize_r_{
            _p0: t1,
            _tag: 1,
        }
        return t2
    }
}

func _goml_m_inherent_i_MutSlice_i_MutSlice_l_T_r__i_fill____T__isize(self__0 []int, value__0 int) struct{} {
    var index__0 int = 0
    Loop_loop0:
    for {
        var t0 int
        var inline0 int = len(self__0)
        t0 = inline0
        var t1 bool = index__0 < t0
        if t1 {
            var index0 int = index__0
            _ = self__0[index0]
            func(p0 []int, p1 int, p2 int) struct{} {
                p0[p1] = p2
                return struct{}{}
            }(self__0, index0, value__0)
            var compound_old0 int = index__0
            var compound_value0 int = 1
            var t3 int = compound_old0 + compound_value0
            index__0 = t3
            continue
        } else {
            break Loop_loop0
        }
    }
    return struct{}{}
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_slice__checked____T__isize(self__0 *_goml_vec_int, start__0 int, end__0 int) _goml_m_Option____Slice_l_isize_r_ {
    var t0 bool = start__0 < 0
    var jp0 bool
    if t0 {
        jp0 = true
    } else {
        var t5 bool = end__0 < start__0
        jp0 = t5
    }
    var jp1 bool
    if jp0 {
        jp1 = true
    } else {
        var t3 int
        var inline0 int = vec_len__Vec_3int(self__0)
        t3 = inline0
        var t4 bool = end__0 > t3
        jp1 = t4
    }
    if jp1 {
        return _goml_m_Option____Slice_l_isize_r_{
            _tag: 0,
        }
    } else {
        var t1 []int = self__0.items[start__0:end__0]
        var t2 _goml_m_Option____Slice_l_isize_r_ = _goml_m_Option____Slice_l_isize_r_{
            _p0: t1,
            _tag: 1,
        }
        return t2
    }
}

func _goml_m_trait__impl_i_ToString_i_bool_i_to__string(self__0 bool) string {
    var t0 string = _goml_runtime_core_bool_to_string(self__0)
    return t0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_len____T__isize(self__0 *_goml_vec_int) int {
    var t0 int = vec_len__Vec_3int(self__0)
    return t0
}

func _goml_m_trait__impl_i_ToString_i_isize_i_to__string(self__0 int) string {
    var inline0 int64 = int64(int(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func _goml_m_inherent_i_Option_i_Option_l_T_r__i_is__some____T__Slice_l_isize_r_(self__0 _goml_m_Option____Slice_l_isize_r_) bool {
    switch self__0._tag {
    case 0:
        return false
    case 1:
        return true
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_inherent_i_Slice_i_Slice_l_T_r__i_len____T__isize(self__0 []int) int {
    var t0 int = len(self__0)
    return t0
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
