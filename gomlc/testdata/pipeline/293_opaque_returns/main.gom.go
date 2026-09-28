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

type Factory struct {}

type Box__isize struct {
    value int
}

type FnIterator__isize struct {
    next_fn func() Option__isize
}

type Box__string struct {
    value string
}

type closure_env_inherent_Vec_Vec_T_iter_T_isize_0 struct {
    index_0 *ref_int_x
    len_1 int
    self_2 *_goml_vec_int
}

type Ordering uint8

type Option__isize struct {
    _p0 int
    _tag uint8
}

type _goml_m_dyn_____d_dyn_k_dynLabel_u_Source_l_Item_x3d_isize_r___vtable struct {
    _goml_m_Source_i_get func(any) int
    _goml_m_Label_i_label func(any) string
}

type _goml_m_dyn_____d_dyn_k_dynLabel_u_Source_l_Item_x3d_isize_r_ struct {
    data any
    vtable *_goml_m_dyn_____d_dyn_k_dynLabel_u_Source_l_Item_x3d_isize_r___vtable
}

func _goml_m_dyn_____d_dyn_k_dynLab_h943b5aa1e7323357508e1ba03a7630d0____Source_i_get(self any) int {
    return _goml_m_trait__impl_i_Source_i_Box____isize_i_get(self.(Box__isize))
}

func _goml_m_dyn_____d_dyn_k_dynLab_hbfbe86fab294ec44a39edc25a581062e___Label_i_label(self any) string {
    return _goml_m_trait__impl_i_Label_i_Box____isize_i_label(self.(Box__isize))
}

func _goml_m_dyn_____d_dyn_k_dynLab_h5eb07c2866f20cfff747e102e79e4f29____Box____isize() *_goml_m_dyn_____d_dyn_k_dynLabel_u_Source_l_Item_x3d_isize_r___vtable {
    return &_goml_m_dyn_____d_dyn_k_dynLabel_u_Source_l_Item_x3d_isize_r___vtable{
        _goml_m_Source_i_get: _goml_m_dyn_____d_dyn_k_dynLab_h943b5aa1e7323357508e1ba03a7630d0____Source_i_get,
        _goml_m_Label_i_label: _goml_m_dyn_____d_dyn_k_dynLab_hbfbe86fab294ec44a39edc25a581062e___Label_i_label,
    }
}

func forward() Box__isize {
    var inline0 int = 12
    var inline1 Box__isize = make__T_isize(inline0)
    return inline1
}

func _goml_m_inherent_i_Factory_i_Factory_i_wrapped(self__0 Factory) Box__isize {
    var t0 Box__isize = Box__isize{
        value: 21,
    }
    return t0
}

func choose(flag__0 bool) Box__isize {
    if flag__0 {
        var t1 Box__isize = Box__isize{
            value: 31,
        }
        return t1
    } else {
        var t0 Box__isize = Box__isize{
            value: 32,
        }
        return t0
    }
}

func numbers() FnIterator__isize {
    var t0 [3]int = [3]int{1, 2, 3}
    var t1 *_goml_vec_int = func(values [3]int) *_goml_vec_int {
        var storage struct {
            vector _goml_vec_int
            values [3]int
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t0)
    var inline0 FnIterator__isize = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_iter____T__isize(t1)
    return inline0
}

func main0() struct{} {
    var number__0 Box__isize = make__T_isize(42)
    var text__0 Box__string = make__T_string("ok")
    var t0 int = _goml_m_trait__impl_i_Source_i_Box____isize_i_get(number__0)
    println__T_isize(t0)
    var t1 string = _goml_m_trait__impl_i_Source_i_Box____string_i_get(text__0)
    println__T_string(t1)
    var t2 string = _goml_m_trait__impl_i_Label_i_Box____isize_i_label(number__0)
    println__T_string(t2)
    var t3 Box__isize = forward()
    var t4 int = _goml_m_trait__impl_i_Source_i_Box____isize_i_get(t3)
    println__T_isize(t4)
    var dynamic__0 _goml_m_dyn_____d_dyn_k_dynLabel_u_Source_l_Item_x3d_isize_r_ = _goml_m_dyn_____d_dyn_k_dynLabel_u_Source_l_Item_x3d_isize_r_{
        data: number__0,
        vtable: _goml_m_dyn_____d_dyn_k_dynLab_h5eb07c2866f20cfff747e102e79e4f29____Box____isize(),
    }
    var t5 int = dynamic__0.vtable._goml_m_Source_i_get(dynamic__0.data)
    println__T_isize(t5)
    var t6 string = dynamic__0.vtable._goml_m_Label_i_label(dynamic__0.data)
    println__T_string(t6)
    var t7 Box__string = Box__string{
        value: "method",
    }
    var t8 Box__string = _goml_m_inherent_i_Box_i_Box_l_T_r__i_wrapped____T__string(t7)
    var t9 string = _goml_m_trait__impl_i_Source_i_Box____string_i_get(t8)
    println__T_string(t9)
    var t10 Factory = Factory{}
    var t11 Box__isize = _goml_m_inherent_i_Factory_i_Factory_i_wrapped(t10)
    var t12 int = _goml_m_trait__impl_i_Source_i_Box____isize_i_get(t11)
    println__T_isize(t12)
    var t13 Box__isize = choose(true)
    var t14 int = _goml_m_trait__impl_i_Source_i_Box____isize_i_get(t13)
    println__T_isize(t14)
    var t15 Box__isize = choose(false)
    var t16 int
    var inline27 int = t15.value
    t16 = inline27
    println__T_isize(t16)
    var iterator__0 FnIterator__isize = numbers()
    var t17 Option__isize = _goml_m_trait__impl_i_Iterator_i_FnIterator____isize_i_next(iterator__0)
    var t18 int = _goml_m_inherent_i_Option_i_Option_l_T_r__i_unwrap__or____T__isize(t17, 0)
    var inline25 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t18)
    _goml_runtime_core_string_println(inline25)
    var t19 Option__isize
    var inline23 func() Option__isize = iterator__0.next_fn
    var inline24 Option__isize = inline23()
    t19 = inline24
    var t20 int
    var inline21 int = 0
    switch t19._tag {
    case 0:
        t20 = inline21
    case 1:
        var inline22 int = t19._p0
        t20 = inline22
    default:
        panic("non-exhaustive match")
    }
    var inline19 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t20)
    _goml_runtime_core_string_println(inline19)
    var t21 Option__isize
    var inline17 func() Option__isize = iterator__0.next_fn
    var inline18 Option__isize = inline17()
    t21 = inline18
    var t22 int
    var inline15 int = 0
    switch t21._tag {
    case 0:
        t22 = inline15
    case 1:
        var inline16 int = t21._p0
        t22 = inline16
    default:
        panic("non-exhaustive match")
    }
    var inline13 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t22)
    _goml_runtime_core_string_println(inline13)
    var t23 Option__isize
    var inline11 func() Option__isize = iterator__0.next_fn
    var inline12 Option__isize = inline11()
    t23 = inline12
    var t24 int
    var inline9 int = 0
    switch t23._tag {
    case 0:
        t24 = inline9
    case 1:
        var inline10 int = t23._p0
        t24 = inline10
    default:
        panic("non-exhaustive match")
    }
    var inline7 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t24)
    _goml_runtime_core_string_println(inline7)
    var t25 int
    var inline5 int = 17
    t25 = inline5
    var inline3 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t25)
    _goml_runtime_core_string_println(inline3)
    var t26 string
    var inline2 string = "generic"
    t26 = inline2
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t26)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func make__T_isize(value__0 int) Box__isize {
    var t0 Box__isize = Box__isize{
        value: value__0,
    }
    return t0
}

func make__T_string(value__0 string) Box__string {
    var t0 Box__string = Box__string{
        value: value__0,
    }
    return t0
}

func println__T_isize(value__0 int) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_int_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_trait__impl_i_Source_i_Box____isize_i_get(self__0 Box__isize) int {
    var t0 int = self__0.value
    return t0
}

func println__T_string(value__0 string) struct{} {
    var t0 string
    t0 = value__0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_trait__impl_i_Source_i_Box____string_i_get(self__0 Box__string) string {
    var t0 string = self__0.value
    return t0
}

func _goml_m_trait__impl_i_Label_i_Box____isize_i_label(self__0 Box__isize) string {
    return "box"
}

func _goml_m_inherent_i_Box_i_Box_l_T_r__i_wrapped____T__string(self__0 Box__string) Box__string {
    return self__0
}

func _goml_m_trait__impl_i_Iterator_i_FnIterator____isize_i_next(self__0 FnIterator__isize) Option__isize {
    var t0 func() Option__isize = self__0.next_fn
    var t1 Option__isize = t0()
    return t1
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

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_iter____T__isize(self__0 *_goml_vec_int) FnIterator__isize {
    var index__0 *ref_int_x = ref__Ref_3int(0)
    var len__0 int
    var inline1 int = vec_len__Vec_3int(self__0)
    len__0 = inline1
    var t0 closure_env_inherent_Vec_Vec_T_iter_T_isize_0 = closure_env_inherent_Vec_Vec_T_iter_T_isize_0{
        index_0: index__0,
        len_1: len__0,
        self_2: self__0,
    }
    var t1 func() Option__isize = func() Option__isize {
        return _goml_m_inherent_i_closure__en_h7a37819447e3ee6c437844690ef4ce6e_size__0_i_apply(t0)
    }
    var inline0 FnIterator__isize = FnIterator__isize{
        next_fn: t1,
    }
    return inline0
}

func _goml_m_trait__impl_i_ToString_i_isize_i_to__string(self__0 int) string {
    var inline0 int64 = int64(int(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
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

func _goml_m_inherent_i_closure__en_h7a37819447e3ee6c437844690ef4ce6e_size__0_i_apply(env0 closure_env_inherent_Vec_Vec_T_iter_T_isize_0) Option__isize {
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
