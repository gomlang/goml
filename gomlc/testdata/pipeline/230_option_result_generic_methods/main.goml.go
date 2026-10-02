package main

import (
    _goml_os "os"
)

func _goml_runtime_core_string_len(s string) int {
    return int(len(s))
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

type closure_env_main_0 struct {}

type closure_env_main_1 struct {}

type closure_env_main_2 struct {}

type closure_env_main_3 struct {}

type closure_env_main_4 struct {}

type closure_env_main_5 struct {}

type closure_env_main_6 struct {}

type closure_env_main_7 struct {}

type Ordering uint8

type Option__isize struct {
    _p0 int
    _tag uint8
}

type Option__string struct {
    _p0 string
    _tag uint8
}

type Result__isize__string struct {
    _p1 string
    _p0 int
    _tag uint8
}

type Result__isize__isize struct {
    _p0 int
    _tag uint8
}

type Result__string__string struct {
    _p0 string
    _tag uint8
}

func main0() struct{} {
    var some__0 Option__isize = Option__isize{
        _p0: 3,
        _tag: 1,
    }
    var t0 closure_env_main_0 = closure_env_main_0{}
    var t1 func(int) string = func(p0 int) string {
        return _goml_m_inherent_i_closure__env__main__0_i_closure__env__main__0_i_apply(t0, p0)
    }
    var mapped__0 Option__string = _goml_m_inherent_i_Option_i_Option_l_T_r__i_map____T__isize____U__string(some__0, t1)
    var t2 string = _goml_m_inherent_i_Option_i_Option_l_T_r__i_unwrap__or____T__string(mapped__0, "missing")
    println__T_string(t2)
    var t3 closure_env_main_1 = closure_env_main_1{}
    var t4 func(int) string = func(p0 int) string {
        return _goml_m_inherent_i_closure__env__main__1_i_closure__env__main__1_i_apply(t3, p0)
    }
    var static_mapped__0 Option__string = _goml_m_inherent_i_Option_i_Option_l_T_r__i_map____T__isize____U__string(some__0, t4)
    var t5 string = _goml_m_inherent_i_Option_i_Option_l_T_r__i_unwrap__or____T__string(static_mapped__0, "missing")
    println__T_string(t5)
    var t6 closure_env_main_2 = closure_env_main_2{}
    var t7 func(int) Option__string = func(p0 int) Option__string {
        return _goml_m_inherent_i_closure__env__main__2_i_closure__env__main__2_i_apply(t6, p0)
    }
    var chained__0 Option__string = _goml_m_inherent_i_Option_i_Option_l_T_r__i_and__then____T__isize____U__string(some__0, t7)
    var t8 string = _goml_m_inherent_i_Option_i_Option_l_T_r__i_unwrap__or____T__string(chained__0, "missing")
    println__T_string(t8)
    var none__0 Option__isize = Option__isize{
        _tag: 0,
    }
    var converted__0 Result__isize__string = _goml_m_inherent_i_Option_i_Option_l_T_r__i_ok__or____E__string____T__isize(none__0, "none")
    var t9 closure_env_main_3 = closure_env_main_3{}
    var t10 func(string) int = func(p0 string) int {
        return _goml_m_inherent_i_closure__env__main__3_i_closure__env__main__3_i_apply(t9, p0)
    }
    var t11 int = _goml_m_inherent_i_Result_i_Re_had11e393bde0ae88c9d8324ffd70f925_ing____T__isize(converted__0, t10)
    println__T_isize(t11)
    var ok__0 Result__isize__string = Result__isize__string{
        _p0: 5,
        _tag: 0,
    }
    var t12 closure_env_main_4 = closure_env_main_4{}
    var t13 func(int) int = func(p0 int) int {
        return _goml_m_inherent_i_closure__env__main__4_i_closure__env__main__4_i_apply(t12, p0)
    }
    var t14 Result__isize__string = _goml_m_inherent_i_Result_i_Re_hf15fd215f39b8121388b37682eabc3c0_ize____U__isize(ok__0, t13)
    var t15 int
    var inline16 int = 0
    switch t14._tag {
    case 0:
        var inline17 int = t14._p0
        t15 = inline17
    case 1:
        t15 = inline16
    default:
        panic("non-exhaustive match")
    }
    var inline14 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t15)
    _goml_runtime_core_string_println(inline14)
    var t16 closure_env_main_5 = closure_env_main_5{}
    var t17 func(string) int = func(p0 string) int {
        return _goml_m_inherent_i_closure__env__main__5_i_closure__env__main__5_i_apply(t16, p0)
    }
    var mapped_error__0 Result__isize__isize
    var inline11 string = "bad"
    var inline12 int = t17(inline11)
    var inline13 Result__isize__isize = Result__isize__isize{
        _p0: inline12,
        _tag: 1,
    }
    mapped_error__0 = inline13
    var t18 closure_env_main_6 = closure_env_main_6{}
    var t19 func(int) int = func(p0 int) int {
        return _goml_m_inherent_i_closure__env__main__6_i_closure__env__main__6_i_apply(t18, p0)
    }
    var t20 int
    switch mapped_error__0._tag {
    case 0:
        var inline8 int = mapped_error__0._p0
        t20 = inline8
    case 1:
        var inline9 int = mapped_error__0._p0
        var inline10 int = t19(inline9)
        t20 = inline10
    default:
        panic("non-exhaustive match")
    }
    var inline6 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t20)
    _goml_runtime_core_string_println(inline6)
    var t21 closure_env_main_7 = closure_env_main_7{}
    var t22 func(int) Result__string__string = func(p0 int) Result__string__string {
        return _goml_m_inherent_i_closure__env__main__7_i_closure__env__main__7_i_apply(t21, p0)
    }
    var next__0 Result__string__string
    var inline4 int = 5
    var inline5 Result__string__string = t22(inline4)
    next__0 = inline5
    var t23 string
    var inline2 string = "missing"
    switch next__0._tag {
    case 0:
        var inline3 string = next__0._p0
        t23 = inline3
    case 1:
        t23 = inline2
    default:
        panic("non-exhaustive match")
    }
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t23)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_inherent_i_Option_i_Option_l_T_r__i_map____T__isize____U__string(self__0 Option__isize, map_fn__0 func(int) string) Option__string {
    switch self__0._tag {
    case 0:
        return Option__string{
            _tag: 0,
        }
    case 1:
        var x0 int = self__0._p0
        var t0 string = map_fn__0(x0)
        var t1 Option__string = Option__string{
            _p0: t0,
            _tag: 1,
        }
        return t1
    default:
        panic("non-exhaustive match")
    }
}

func println__T_string(value__0 string) struct{} {
    var t0 string
    t0 = value__0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_inherent_i_Option_i_Option_l_T_r__i_unwrap__or____T__string(self__0 Option__string, fallback__0 string) string {
    switch self__0._tag {
    case 0:
        return fallback__0
    case 1:
        var x0 string = self__0._p0
        return x0
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_inherent_i_Option_i_Option_l_T_r__i_and__then____T__isize____U__string(self__0 Option__isize, next__0 func(int) Option__string) Option__string {
    switch self__0._tag {
    case 0:
        return Option__string{
            _tag: 0,
        }
    case 1:
        var x0 int = self__0._p0
        var t0 Option__string = next__0(x0)
        return t0
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_inherent_i_Option_i_Option_l_T_r__i_ok__or____E__string____T__isize(self__0 Option__isize, error__0 string) Result__isize__string {
    switch self__0._tag {
    case 0:
        var t0 Result__isize__string = Result__isize__string{
            _p1: error__0,
            _tag: 1,
        }
        return t0
    case 1:
        var x0 int = self__0._p0
        var t1 Result__isize__string = Result__isize__string{
            _p0: x0,
            _tag: 0,
        }
        return t1
    default:
        panic("non-exhaustive match")
    }
}

func println__T_isize(value__0 int) struct{} {
    var t0 string
    var inline0 string = __goml_builtin_int_to_string(value__0)
    t0 = inline0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_inherent_i_Result_i_Re_had11e393bde0ae88c9d8324ffd70f925_ing____T__isize(self__0 Result__isize__string, fallback__0 func(string) int) int {
    switch self__0._tag {
    case 0:
        var x0 int = self__0._p0
        return x0
    case 1:
        var x1 string = self__0._p1
        var t0 int = fallback__0(x1)
        return t0
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_inherent_i_Result_i_Re_hf15fd215f39b8121388b37682eabc3c0_ize____U__isize(self__0 Result__isize__string, map_fn__0 func(int) int) Result__isize__string {
    switch self__0._tag {
    case 0:
        var x0 int = self__0._p0
        var t0 int = map_fn__0(x0)
        var t1 Result__isize__string = Result__isize__string{
            _p0: t0,
            _tag: 0,
        }
        return t1
    case 1:
        var x1 string = self__0._p1
        var t2 Result__isize__string = Result__isize__string{
            _p1: x1,
            _tag: 1,
        }
        return t2
    default:
        panic("non-exhaustive match")
    }
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

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func _goml_m_trait__impl_i_ToString_i_isize_i_to__string(self__0 int) string {
    var inline0 int64 = int64(int(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
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

func _goml_m_inherent_i_closure__env__main__0_i_closure__env__main__0_i_apply(env0 closure_env_main_0, value__0 int) string {
    var inline0 string = __goml_builtin_int_to_string(value__0)
    return inline0
}

func _goml_m_inherent_i_closure__env__main__1_i_closure__env__main__1_i_apply(env0 closure_env_main_1, value__0 int) string {
    var t0 string
    var inline0 string = __goml_builtin_int_to_string(value__0)
    t0 = inline0
    var t1 string = "static:" + t0
    return t1
}

func _goml_m_inherent_i_closure__env__main__2_i_closure__env__main__2_i_apply(env0 closure_env_main_2, value__0 int) Option__string {
    var t0 string
    var inline0 string = __goml_builtin_int_to_string(value__0)
    t0 = inline0
    var t1 string = "value:" + t0
    var t2 Option__string = Option__string{
        _p0: t1,
        _tag: 1,
    }
    return t2
}

func _goml_m_inherent_i_closure__env__main__3_i_closure__env__main__3_i_apply(env0 closure_env_main_3, error__0 string) int {
    var inline0 int = _goml_runtime_core_string_len(error__0)
    return inline0
}

func _goml_m_inherent_i_closure__env__main__4_i_closure__env__main__4_i_apply(env0 closure_env_main_4, value__0 int) int {
    var t0 int = value__0 + 2
    return t0
}

func _goml_m_inherent_i_closure__env__main__5_i_closure__env__main__5_i_apply(env0 closure_env_main_5, value__0 string) int {
    var inline0 int = _goml_runtime_core_string_len(value__0)
    return inline0
}

func _goml_m_inherent_i_closure__env__main__6_i_closure__env__main__6_i_apply(env0 closure_env_main_6, value__0 int) int {
    return value__0
}

func _goml_m_inherent_i_closure__env__main__7_i_closure__env__main__7_i_apply(env0 closure_env_main_7, value__0 int) Result__string__string {
    var t0 string
    var inline0 string = __goml_builtin_int_to_string(value__0)
    t0 = inline0
    var t1 string = "next:" + t0
    var t2 Result__string__string = Result__string__string{
        _p0: t1,
        _tag: 0,
    }
    return t2
}

func main() {
    main0()
}
