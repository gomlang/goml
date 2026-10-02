package main

import (
    _goml_os "os"
    _goml_reflect "reflect"
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

type _goml_vec_int32 struct {
    items []int32
}

func vec_new__Vec_5int32() *_goml_vec_int32 {
    return &_goml_vec_int32{
        items: nil,
    }
}

func vec_push__Vec_5int32(vec *_goml_vec_int32, elem int32) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_set__Vec_5int32(vec *_goml_vec_int32, index int, value int32) struct{} {
    vec.items[index] = value
    return struct{}{}
}

func vec_len__Vec_5int32(vec *_goml_vec_int32) int {
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

func ptr_eq__Ref_5int32(a *ref_int32_x, b *ref_int32_x) bool {
    return a == b
}

func ptr_hash__Ref_5int32(reference *ref_int32_x) uint64 {
    return uint64(_goml_reflect.ValueOf(reference).Pointer())
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

type Tuple2_5int32_4bool struct {
    _0 int32
    _1 bool
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

type closure_env_goml_builtin_range_1 struct {
    current_0 *ref_int_x
    end_1 int
}

type Ordering uint8

type Option__isize struct {
    _p0 int
    _tag uint8
}

func main0() struct{} {
    var native__0 int = 7
    var small__0 int8 = 8
    var unsigned__0 uint8 = 9
    var t0 string
    var inline19 string = __goml_builtin_int_to_string(native__0)
    t0 = inline19
    var t1 string
    var inline18 string = __goml_builtin_int8_to_string(small__0)
    t1 = inline18
    var t2 string = t0 + t1
    var t3 string
    var inline17 string = __goml_builtin_uint8_to_string(unsigned__0)
    t3 = inline17
    var t4 string = t2 + t3
    var t5 string
    var inline9 string = "abcd"
    var inline10 int = 1
    var inline11 int = 3
    var inline12 bool = string_is_char_boundary(inline9, inline10)
    var inline13 bool
    if inline12 {
        var inline16 bool = string_is_char_boundary(inline9, inline11)
        inline13 = inline16
    } else {
        inline13 = false
    }
    if inline13 {
        var inline14 string = _goml_runtime_core_string_byte_slice(inline9, inline10, inline11)
        t5 = inline14
        var text__0 string = t4 + t5
        var value__0 *ref_int32_x = ref__Ref_5int32(1)
        ref_set__Ref_5int32(value__0, 2)
        var same__0 bool = ptr_eq__Ref_5int32(value__0, value__0)
        ptr_hash__Ref_5int32(value__0)
        var values__0 *_goml_vec_int32 = vec_new__Vec_5int32()
        var t6 int32 = ref_get__Ref_5int32(value__0)
        vec_push__Vec_5int32(values__0, t6)
        vec_push__Vec_5int32(values__0, 3)
        vec_set__Vec_5int32(values__0, 1, 4)
        var t7 int = vec_len__Vec_5int32(values__0)
        var values_slice__0 []int32 = values__0.items[0:t7]
        var t8 int = len(values_slice__0)
        var nested__0 []int32 = values_slice__0[0:t8]
        var channel__0 chan int32 = func(p0 int) chan int32 {
            return make(chan int32, p0)
        }(1)
        var t9 int32 = nested__0[1]
        func(p0 chan int32, p1 int32) struct{} {
            p0 <- p1
            return struct{}{}
        }(channel__0, t9)
        var received__0 Tuple2_5int32_4bool = func(p0 chan int32) Tuple2_5int32_4bool {
            var value int32
            var ok bool
            value, ok = <-p0
            return Tuple2_5int32_4bool{
                _0: value,
                _1: ok,
            }
        }(channel__0)
        func(p0 chan int32) struct{} {
            close(p0)
            return struct{}{}
        }(channel__0)
        var t10 FnIterator__isize
        var inline6 int = 0
        var inline7 int = 3
        var inline8 FnIterator__isize = __goml_builtin_range(inline6, inline7)
        t10 = inline8
        var t11 closure_env_main_0 = closure_env_main_0{}
        var t12 func(int, int) int = func(p0 int, p1 int) int {
            return _goml_m_inherent_i_closure__env__main__0_i_closure__env__main__0_i_apply(t11, p0, p1)
        }
        var total__0 int = _goml_m_std_p_iter_p_fold____A__isize____I__FnIterator_l_isize_r_____T__isize(t10, 0, t12)
        var t13 string
        var inline5 string = _goml_runtime_core_bool_to_string(same__0)
        t13 = inline5
        var t14 string = text__0 + t13
        var t15 int32 = received__0._0
        var t16 string
        var inline4 string = __goml_builtin_int32_to_string(t15)
        t16 = inline4
        var t17 string = t14 + t16
        var t18 bool = received__0._1
        var t19 string
        var inline3 string = _goml_runtime_core_bool_to_string(t18)
        t19 = inline3
        var t20 string = t17 + t19
        var t21 string
        var inline2 string = __goml_builtin_int_to_string(total__0)
        t21 = inline2
        var t22 string = t20 + t21
        var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t22)
        _goml_runtime_core_string_println(inline0)
        return struct{}{}
    } else {
        var inline15 string = _goml_runtime_core_string_byte_slice(inline9, -1, -1)
        t5 = inline15
        var text__0 string = t4 + t5
        var value__0 *ref_int32_x = ref__Ref_5int32(1)
        ref_set__Ref_5int32(value__0, 2)
        var same__0 bool = ptr_eq__Ref_5int32(value__0, value__0)
        ptr_hash__Ref_5int32(value__0)
        var values__0 *_goml_vec_int32 = vec_new__Vec_5int32()
        var t6 int32 = ref_get__Ref_5int32(value__0)
        vec_push__Vec_5int32(values__0, t6)
        vec_push__Vec_5int32(values__0, 3)
        vec_set__Vec_5int32(values__0, 1, 4)
        var t7 int = vec_len__Vec_5int32(values__0)
        var values_slice__0 []int32 = values__0.items[0:t7]
        var t8 int = len(values_slice__0)
        var nested__0 []int32 = values_slice__0[0:t8]
        var channel__0 chan int32 = func(p0 int) chan int32 {
            return make(chan int32, p0)
        }(1)
        var t9 int32 = nested__0[1]
        func(p0 chan int32, p1 int32) struct{} {
            p0 <- p1
            return struct{}{}
        }(channel__0, t9)
        var received__0 Tuple2_5int32_4bool = func(p0 chan int32) Tuple2_5int32_4bool {
            var value int32
            var ok bool
            value, ok = <-p0
            return Tuple2_5int32_4bool{
                _0: value,
                _1: ok,
            }
        }(channel__0)
        func(p0 chan int32) struct{} {
            close(p0)
            return struct{}{}
        }(channel__0)
        var t10 FnIterator__isize
        var inline6 int = 0
        var inline7 int = 3
        var inline8 FnIterator__isize = __goml_builtin_range(inline6, inline7)
        t10 = inline8
        var t11 closure_env_main_0 = closure_env_main_0{}
        var t12 func(int, int) int = func(p0 int, p1 int) int {
            return _goml_m_inherent_i_closure__env__main__0_i_closure__env__main__0_i_apply(t11, p0, p1)
        }
        var total__0 int = _goml_m_std_p_iter_p_fold____A__isize____I__FnIterator_l_isize_r_____T__isize(t10, 0, t12)
        var t13 string
        var inline5 string = _goml_runtime_core_bool_to_string(same__0)
        t13 = inline5
        var t14 string = text__0 + t13
        var t15 int32 = received__0._0
        var t16 string
        var inline4 string = __goml_builtin_int32_to_string(t15)
        t16 = inline4
        var t17 string = t14 + t16
        var t18 bool = received__0._1
        var t19 string
        var inline3 string = _goml_runtime_core_bool_to_string(t18)
        t19 = inline3
        var t20 string = t17 + t19
        var t21 string
        var inline2 string = __goml_builtin_int_to_string(total__0)
        t21 = inline2
        var t22 string = t20 + t21
        var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t22)
        _goml_runtime_core_string_println(inline0)
        return struct{}{}
    }
}

func _goml_m_std_p_iter_p_fold____A__isize____I__FnIterator_l_isize_r_____T__isize(iterator__0 FnIterator__isize, initial__0 int, combine__0 func(int, int) int) int {
    var accumulator__0 int = initial__0
    Loop_loop0:
    for {
        var mtmp0 Option__isize
        var inline0 func() Option__isize = iterator__0.next_fn
        var inline1 Option__isize = inline0()
        mtmp0 = inline1
        switch mtmp0._tag {
        case 1:
            var x0 int = mtmp0._p0
            var t0 int = combine__0(accumulator__0, x0)
            accumulator__0 = t0
            continue
        default:
            break Loop_loop0
        }
    }
    return accumulator__0
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

func __goml_builtin_int8_to_string(value__0 int8) string {
    var t0 int64 = int64(int8(value__0))
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

func __goml_builtin_uint8_to_string(value__0 uint8) string {
    var t0 uint64 = uint64(uint8(value__0))
    var t1 string = decimal_string(t0)
    return t1
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

func __goml_builtin_range(start__0 int, end__0 int) FnIterator__isize {
    var current__0 *ref_int_x = ref__Ref_3int(start__0)
    var t0 closure_env_goml_builtin_range_1 = closure_env_goml_builtin_range_1{
        current_0: current__0,
        end_1: end__0,
    }
    var t1 func() Option__isize = func() Option__isize {
        return _goml_m_inherent_i_closure__en_h07c29ff1f344b08e028033881af7c2d9_ange__1_i_apply(t0)
    }
    var inline0 FnIterator__isize = FnIterator__isize{
        next_fn: t1,
    }
    return inline0
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
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

func _goml_m_inherent_i_closure__env__main__0_i_closure__env__main__0_i_apply(env0 closure_env_main_0, sum__0 int, item__0 int) int {
    var t0 int = sum__0 + item__0
    return t0
}

func _goml_m_inherent_i_closure__en_h07c29ff1f344b08e028033881af7c2d9_ange__1_i_apply(env0 closure_env_goml_builtin_range_1) Option__isize {
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
