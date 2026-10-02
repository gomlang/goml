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

type ref_Wide_x struct {
    value Wide
}

func ref__Ref_4Wide(value Wide) *ref_Wide_x {
    return &ref_Wide_x{
        value: value,
    }
}

func ref_get__Ref_4Wide(reference *ref_Wide_x) Wide {
    return reference.value
}

func ref_set__Ref_4Wide(reference *ref_Wide_x, value Wide) struct{} {
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

type Wide struct {
    a int
    b int
    c int
    d int
    e int
    f int
    g int
    h int
    i int
    j int
}

type closure_env_make_reader_0 struct {
    value_0 Wide
}

type closure_env_action_1 struct {
    value_0 *ref_Wide_x
}

type Ordering uint8

func read(value__0 *Wide, action__0 func() struct{}, depth__0 int) int {
    var t0 bool = depth__0 == 0
    if t0 {
        action__0()
        var t1 int = (*value__0).a
        return t1
    } else {
        var t2 int = depth__0 - 1
        var t3 int = read(value__0, action__0, t2)
        return t3
    }
}

func modify(value__0_pointer *Wide, depth__0 int) Wide {
    var value__0 Wide = *value__0_pointer
    var value0 int = 7
    value__0.a = value0
    var t11 bool = depth__0 == 0
    if t11 {
        return value__0
    } else {
        var t12 int = depth__0 - 1
        var t13 Wide = modify(&value__0, t12)
        return t13
    }
}

func make_reader(value__0 *Wide, depth__0 int) func() int {
    var t0 bool = depth__0 == 0
    if t0 {
        var t1 closure_env_make_reader_0 = closure_env_make_reader_0{
            value_0: *value__0,
        }
        var t2 func() int = func() int {
            return _goml_m_inherent_i_closure__en_h058ceb7a4121cc1cd37232df6ad1d333_ader__0_i_apply(t1)
        }
        return t2
    } else {
        var t3 int = depth__0 - 1
        var t4 func() int = make_reader(value__0, t3)
        return t4
    }
}

func main0() struct{} {
    var t0 Wide = Wide{
        a: 1,
        b: 2,
        c: 3,
        d: 4,
        e: 5,
        f: 6,
        g: 7,
        h: 8,
        i: 9,
        j: 10,
    }
    var value__0 *ref_Wide_x = ref__Ref_4Wide(t0)
    var t1 closure_env_action_1 = closure_env_action_1{
        value_0: value__0,
    }
    var action__0 func() struct{} = func() struct{} {
        return _goml_m_inherent_i_closure__env__action__1_i_closure__env__action__1_i_apply(t1)
    }
    var t2 Wide = ref_get__Ref_4Wide(value__0)
    var t3 int = read(&t2, action__0, 2)
    var inline12 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t3)
    _goml_runtime_core_string_println(inline12)
    var t4 Wide = ref_get__Ref_4Wide(value__0)
    var t5 int = t4.a
    var inline10 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t5)
    _goml_runtime_core_string_println(inline10)
    var t6 Wide = ref_get__Ref_4Wide(value__0)
    var result__0 Wide = modify(&t6, 2)
    var t7 int = result__0.a
    var inline8 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t7)
    _goml_runtime_core_string_println(inline8)
    var t8 Wide = ref_get__Ref_4Wide(value__0)
    var t9 int = t8.a
    var inline6 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t9)
    _goml_runtime_core_string_println(inline6)
    var t10 int = result__0.j
    var inline4 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t10)
    _goml_runtime_core_string_println(inline4)
    var t11 Wide = ref_get__Ref_4Wide(value__0)
    var reader__0 func() int = make_reader(&t11, 2)
    var place_root0 Wide = ref_get__Ref_4Wide(value__0)
    var value0 int = 12
    var t21 Wide = place_root0
    t21.a = value0
    ref_set__Ref_4Wide(value__0, t21)
    var t23 int = reader__0()
    var inline2 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t23)
    _goml_runtime_core_string_println(inline2)
    var t24 Wide = ref_get__Ref_4Wide(value__0)
    var t25 int = t24.a
    var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(t25)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
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

func _goml_m_inherent_i_closure__en_h058ceb7a4121cc1cd37232df6ad1d333_ader__0_i_apply(env0 closure_env_make_reader_0) int {
    var value__0 Wide = env0.value_0
    var t0 int = value__0.a
    return t0
}

func _goml_m_inherent_i_closure__env__action__1_i_closure__env__action__1_i_apply(env0 closure_env_action_1) struct{} {
    var value__0 *ref_Wide_x = env0.value_0
    var place_root0 Wide = ref_get__Ref_4Wide(value__0)
    var value0 int = 9
    var t9 Wide = place_root0
    t9.a = value0
    ref_set__Ref_4Wide(value__0, t9)
    return struct{}{}
}

func main() {
    main0()
}
