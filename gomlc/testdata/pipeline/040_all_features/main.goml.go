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

func array_get__Array_2_20Fn1_5int32_to_5int32(arr [2]func(int32) int32, index int) func(int32) int32 {
    return arr[index]
}

func array_get__Array_2_5int32(arr [2]int32, index int) int32 {
    return arr[index]
}

func array_set__Array_2_5int32(arr [2]int32, index int, value int32) [2]int32 {
    arr[index] = value
    return arr
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

type ref_bool_x struct {
    value bool
}

func ref__Ref_4bool(value bool) *ref_bool_x {
    return &ref_bool_x{
        value: value,
    }
}

func ref_get__Ref_4bool(reference *ref_bool_x) bool {
    return reference.value
}

func ref_set__Ref_4bool(reference *ref_bool_x, value bool) struct{} {
    reference.value = value
    return struct{}{}
}

type Tuple2_6string_6string struct {
    _0 string
    _1 string
}

type Tuple4_7Tracker_24TFunc0_ret_1_hd2fccc77024d42a64d1c9425814244a6_4Record__string struct {
    _0 Tracker
    _1 func() Record__i32
    _2 func(int32) Record__i32
    _3 func() Record__string
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

type Tracker struct {
    label string
    count *ref_int32_x
    toggled *ref_bool_x
}

type closure_env_snapshot_0 struct {
    count_0 *ref_int32_x
}

type closure_env_bump_1 struct {
    count_0 *ref_int32_x
}

type closure_env_flip_2 struct {
    toggled_0 *ref_bool_x
}

type Ordering uint8

type Record__i32 struct {
    _p0 int32
    _p1 int32
    _tag uint8
}

type Record__string interface {
    isRecord__string()
}

type Record__string_Value struct {
    _0 string
}

func (_ Record__string_Value) isRecord__string() {}

type Record__string_Pair struct {
    _0 string
    _1 string
}

func (_ Record__string_Pair) isRecord__string() {}

type Record__string_Empty struct {}

func (_ Record__string_Empty) isRecord__string() {}

type Maybe__i32 uint64

type Maybe__string struct {
    _p0 string
    _tag uint8
}

func _goml_m_trait__impl_i_Describe_i_Tracker_i_describe(self__0 Tracker) string {
    var x0 string = self__0.label
    var x1 *ref_int32_x = self__0.count
    var x2 *ref_bool_x = self__0.toggled
    var current__0 int32
    var inline3 int32 = ref_get__Ref_5int32(x1)
    current__0 = inline3
    var flag__0 bool
    var inline2 bool = ref_get__Ref_4bool(x2)
    flag__0 = inline2
    var with_label__0 string = "Tracker(" + x0
    var with_count_label__0 string = with_label__0 + ", count: "
    var t0 string
    var inline1 string = __goml_builtin_int32_to_string(current__0)
    t0 = inline1
    var with_count__0 string = with_count_label__0 + t0
    var with_flag_label__0 string = with_count__0 + ", toggled: "
    var t1 string
    var inline0 string = _goml_runtime_core_bool_to_string(flag__0)
    t1 = inline0
    var t2 string = with_flag_label__0 + t1
    var t3 string = t2 + ")"
    return t3
}

func _goml_m_trait__impl_i_Describe_i_Record____i32_i_describe(self__0 Record__i32) string {
    switch self__0._tag {
    case 0:
        var x0 int32 = self__0._p0
        var t0 string
        var inline0 string = __goml_builtin_int32_to_string(x0)
        t0 = inline0
        var t1 string = "Value(" + t0
        var t2 string = t1 + ")"
        return t2
    case 1:
        var x1 int32 = self__0._p0
        var x2 int32 = self__0._p1
        var t3 string
        var inline2 string = __goml_builtin_int32_to_string(x1)
        t3 = inline2
        var prefix__0 string = "Pair(" + t3
        var t4 string = prefix__0 + ", "
        var t5 string
        var inline1 string = __goml_builtin_int32_to_string(x2)
        t5 = inline1
        var t6 string = t4 + t5
        var t7 string = t6 + ")"
        return t7
    case 2:
        return "Empty"
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_trait__impl_i_Describe_i_Record____string_i_describe(self__0 Record__string) string {
    switch self__0.(type) {
    case Record__string_Value:
        var x0 string = self__0.(Record__string_Value)._0
        var t0 string = "Value(" + x0
        var t1 string = t0 + ")"
        return t1
    case Record__string_Pair:
        var x1 string = self__0.(Record__string_Pair)._0
        var x2 string = self__0.(Record__string_Pair)._1
        var prefix__0 string = "Pair(" + x1
        var t2 string = prefix__0 + ", "
        var t3 string = t2 + x2
        var t4 string = t3 + ")"
        return t4
    case Record__string_Empty:
        return "Empty"
    default:
        panic("non-exhaustive match")
    }
}

func format_total(total__0 int32) string {
    var t0 string
    var inline0 string = __goml_builtin_int32_to_string(total__0)
    t0 = inline0
    var t1 string = "total: " + t0
    return t1
}

func increment(value__0 int32) int32 {
    var t0 int32 = value__0 + 1
    return t0
}

func triple(value__0 int32) int32 {
    var t0 int32 = value__0 * 3
    return t0
}

func gather(record__0 Record__i32) Maybe__i32 {
    switch record__0._tag {
    case 0:
        var x0 int32 = record__0._p0
        var t0 Maybe__i32 = Maybe__i32(uint64(int64(x0) + 2147483648) + 1)
        return t0
    case 1:
        var x1 int32 = record__0._p1
        var t1 Maybe__i32 = Maybe__i32(uint64(int64(x1) + 2147483648) + 1)
        return t1
    case 2:
        return Maybe__i32(0)
    default:
        panic("non-exhaustive match")
    }
}

func build_counter(label__0 string, start__0 int32) Tuple4_7Tracker_24TFunc0_ret_1_hd2fccc77024d42a64d1c9425814244a6_4Record__string {
    var count__0 *ref_int32_x
    var inline2 *ref_int32_x = ref__Ref_5int32(start__0)
    count__0 = inline2
    var toggled__0 *ref_bool_x
    var inline0 bool = false
    var inline1 *ref_bool_x = ref__Ref_4bool(inline0)
    toggled__0 = inline1
    var tracker__0 Tracker = Tracker{
        label: label__0,
        count: count__0,
        toggled: toggled__0,
    }
    var t0 closure_env_snapshot_0 = closure_env_snapshot_0{
        count_0: count__0,
    }
    var snapshot__0 func() Record__i32 = func() Record__i32 {
        return _goml_m_inherent_i_closure__env__snapshot__0_i_closure__env__snapshot__0_i_apply(t0)
    }
    var t1 closure_env_bump_1 = closure_env_bump_1{
        count_0: count__0,
    }
    var bump__0 func(int32) Record__i32 = func(p0 int32) Record__i32 {
        return _goml_m_inherent_i_closure__env__bump__1_i_closure__env__bump__1_i_apply(t1, p0)
    }
    var t2 closure_env_flip_2 = closure_env_flip_2{
        toggled_0: toggled__0,
    }
    var flip__0 func() Record__string = func() Record__string {
        return _goml_m_inherent_i_closure__env__flip__2_i_closure__env__flip__2_i_apply(t2)
    }
    var t3 Tuple4_7Tracker_24TFunc0_ret_1_hd2fccc77024d42a64d1c9425814244a6_4Record__string = Tuple4_7Tracker_24TFunc0_ret_1_hd2fccc77024d42a64d1c9425814244a6_4Record__string{
        _0: tracker__0,
        _1: snapshot__0,
        _2: bump__0,
        _3: flip__0,
    }
    return t3
}

func main0() struct{} {
    var mtmp0 Tuple4_7Tracker_24TFunc0_ret_1_hd2fccc77024d42a64d1c9425814244a6_4Record__string = build_counter("goml", 2)
    var x0 Tracker = mtmp0._0
    var x1 func() Record__i32 = mtmp0._1
    var x2 func(int32) Record__i32 = mtmp0._2
    var x3 func() Record__string = mtmp0._3
    var tracker_info__0 string = _goml_m_trait__impl_i_Describe_i_Tracker_i_describe(x0)
    var first_record__0 Record__i32 = x1()
    var bumped_record__0 Record__i32 = x2(5)
    var flipped_record__0 Record__string = x3()
    var maybe_first__0 Maybe__i32 = gather(first_record__0)
    var maybe_second__0 Maybe__i32 = gather(bumped_record__0)
    var chosen__0 Maybe__i32 = _goml_m_choose____T__Maybe_l_i32_r_(true, maybe_second__0, maybe_first__0)
    var stringified__0 Maybe__string = map_maybe__T_i32__U_string(chosen__0, format_total)
    var transforms__0 [2]func(int32) int32 = [2]func(int32) int32{increment, triple}
    var results__0 [2]int32
    var inline19 int32 = 4
    var inline20 func(int32) int32 = array_get__Array_2_20Fn1_5int32_to_5int32(transforms__0, 0)
    var inline21 func(int32) int32 = array_get__Array_2_20Fn1_5int32_to_5int32(transforms__0, 1)
    var inline22 int32 = inline20(inline19)
    var inline23 int32 = inline21(inline22)
    var inline24 [2]int32 = [2]int32{inline22, inline19}
    var inline25 [2]int32 = inline24
    var inline26 int = 1
    array_get__Array_2_5int32(inline25, inline26)
    var inline28 [2]int32 = array_set__Array_2_5int32(inline25, inline26, inline23)
    inline24 = inline28
    results__0 = inline24
    var first_result__0 int32 = array_get__Array_2_5int32(results__0, 0)
    var second_result__0 int32 = array_get__Array_2_5int32(results__0, 1)
    var t0 bool = first_result__0 < second_result__0
    var jp0 bool
    if t0 {
        jp0 = true
    } else {
        jp0 = false
    }
    var first_text__0 string = _goml_m_trait__impl_i_Describe_i_Record____i32_i_describe(first_record__0)
    var bumped_text__0 string = _goml_m_trait__impl_i_Describe_i_Record____i32_i_describe(bumped_record__0)
    var flipped_text__0 string = _goml_m_trait__impl_i_Describe_i_Record____string_i_describe(flipped_record__0)
    var jp1 string
    switch stringified__0._tag {
    case 0:
        var x4 string = stringified__0._p0
        var t3 string = "Snapshot: " + x4
        jp1 = t3
    case 1:
        jp1 = "Snapshot: none"
    default:
        panic("non-exhaustive match")
    }
    var t1 string
    var inline18 string = __goml_builtin_int32_to_string(first_result__0)
    t1 = inline18
    var t2 string
    var inline17 string = __goml_builtin_int32_to_string(second_result__0)
    t2 = inline17
    var pair_text__0 string
    var inline15 string = t1 + " -> "
    var inline16 string = inline15 + t2
    pair_text__0 = inline16
    var bool_text__0 string
    var inline14 string = _goml_runtime_core_bool_to_string(jp0)
    bool_text__0 = inline14
    var inline12 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(tracker_info__0)
    _goml_runtime_core_string_println(inline12)
    var inline10 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(first_text__0)
    _goml_runtime_core_string_println(inline10)
    var inline8 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(bumped_text__0)
    _goml_runtime_core_string_println(inline8)
    var inline6 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(flipped_text__0)
    _goml_runtime_core_string_println(inline6)
    var inline4 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(jp1)
    _goml_runtime_core_string_println(inline4)
    var inline2 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(pair_text__0)
    _goml_runtime_core_string_println(inline2)
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(bool_text__0)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_choose____T__Maybe_l_i32_r_(flag__0 bool, when_true__0 Maybe__i32, when_false__0 Maybe__i32) Maybe__i32 {
    if flag__0 {
        return when_true__0
    } else {
        return when_false__0
    }
}

func map_maybe__T_i32__U_string(value__0 Maybe__i32, f__0 func(int32) string) Maybe__string {
    switch value__0 != Maybe__i32(0) {
    case true:
        var x0 int32 = int32(int64(uint64(value__0) - 1) - 2147483648)
        var t0 string = f__0(x0)
        var t1 Maybe__string = Maybe__string{
            _p0: t0,
            _tag: 0,
        }
        return t1
    case false:
        return Maybe__string{
            _tag: 1,
        }
    default:
        panic("non-exhaustive match")
    }
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

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
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

func _goml_m_inherent_i_closure__env__snapshot__0_i_closure__env__snapshot__0_i_apply(env0 closure_env_snapshot_0) Record__i32 {
    var count__0 *ref_int32_x = env0.count_0
    var t0 int32
    var inline0 int32 = ref_get__Ref_5int32(count__0)
    t0 = inline0
    var t1 Record__i32 = Record__i32{
        _p0: t0,
        _tag: 0,
    }
    return t1
}

func _goml_m_inherent_i_closure__env__bump__1_i_closure__env__bump__1_i_apply(env0 closure_env_bump_1, delta__0 int32) Record__i32 {
    var count__0 *ref_int32_x = env0.count_0
    var before__0 int32
    var inline2 int32 = ref_get__Ref_5int32(count__0)
    before__0 = inline2
    var t0 int32 = before__0 + delta__0
    ref_set__Ref_5int32(count__0, t0)
    var t1 int32
    var inline0 int32 = ref_get__Ref_5int32(count__0)
    t1 = inline0
    var t2 Record__i32 = Record__i32{
        _p0: before__0,
        _p1: t1,
        _tag: 1,
    }
    return t2
}

func _goml_m_inherent_i_closure__env__flip__2_i_closure__env__flip__2_i_apply(env0 closure_env_flip_2) Record__string {
    var toggled__0 *ref_bool_x = env0.toggled_0
    var before__0 bool
    var inline4 bool = ref_get__Ref_4bool(toggled__0)
    before__0 = inline4
    var t0 bool = !before__0
    ref_set__Ref_4bool(toggled__0, t0)
    var after__0 bool
    var inline2 bool = ref_get__Ref_4bool(toggled__0)
    after__0 = inline2
    var t1 string
    var inline1 string = _goml_runtime_core_bool_to_string(before__0)
    t1 = inline1
    var t2 string
    var inline0 string = _goml_runtime_core_bool_to_string(after__0)
    t2 = inline0
    var t3 Record__string = Record__string_Pair{
        _0: t1,
        _1: t2,
    }
    return t3
}

func main() {
    main0()
}
