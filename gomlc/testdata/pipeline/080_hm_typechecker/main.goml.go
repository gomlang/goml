package main

import (
    _goml_os "os"
)

func _goml_runtime_core_string_get(s string, i int) rune {
    return rune(s[i])
}

func _goml_runtime_core_string_from_utf8(bytes *_goml_vec_uint8) Tuple2_4bool_6string {
    return Tuple2_4bool_6string{
        _0: true,
        _1: string(bytes.items),
    }
}

func _goml_runtime_core_char_to_string(x rune) string {
    return string(x)
}

func _goml_runtime_core_string_println(s string) struct{} {
    _goml_os.Stdout.WriteString(s + "\n")
    return struct{}{}
}

type _goml_vec_EnvEntry struct {
    items []EnvEntry
}

func vec_new__Vec_8EnvEntry() *_goml_vec_EnvEntry {
    return &_goml_vec_EnvEntry{
        items: nil,
    }
}

func vec_with_capacity__Vec_8EnvEntry(capacity int) *_goml_vec_EnvEntry {
    return &_goml_vec_EnvEntry{
        items: make([]EnvEntry, 0, capacity),
    }
}

func vec_push__Vec_8EnvEntry(vec *_goml_vec_EnvEntry, elem EnvEntry) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_get__Vec_8EnvEntry(vec *_goml_vec_EnvEntry, index int) EnvEntry {
    return vec.items[index]
}

func vec_len__Vec_8EnvEntry(vec *_goml_vec_EnvEntry) int {
    return int(len(vec.items))
}

type _goml_vec_SubstEntry struct {
    items []SubstEntry
}

func vec_new__Vec_10SubstEntry() *_goml_vec_SubstEntry {
    return &_goml_vec_SubstEntry{
        items: nil,
    }
}

func vec_with_capacity__Vec_10SubstEntry(capacity int) *_goml_vec_SubstEntry {
    return &_goml_vec_SubstEntry{
        items: make([]SubstEntry, 0, capacity),
    }
}

func vec_push__Vec_10SubstEntry(vec *_goml_vec_SubstEntry, elem SubstEntry) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_get__Vec_10SubstEntry(vec *_goml_vec_SubstEntry, index int) SubstEntry {
    return vec.items[index]
}

func vec_len__Vec_10SubstEntry(vec *_goml_vec_SubstEntry) int {
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

type ref_Tv_x struct {
    value Tv
}

func ref__Ref_2Tv(value Tv) *ref_Tv_x {
    return &ref_Tv_x{
        value: value,
    }
}

func ref_get__Ref_2Tv(reference *ref_Tv_x) Tv {
    return reference.value
}

func ref_set__Ref_2Tv(reference *ref_Tv_x, value Tv) struct{} {
    reference.value = value
    return struct{}{}
}

func ptr_eq__Ref_2Tv(a *ref_Tv_x, b *ref_Tv_x) bool {
    return a == b
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

type ref_Option__Typ_x struct {
    value Option__Typ
}

func ref__Ref_11Option__Typ(value Option__Typ) *ref_Option__Typ_x {
    return &ref_Option__Typ_x{
        value: value,
    }
}

func ref_get__Ref_11Option__Typ(reference *ref_Option__Typ_x) Option__Typ {
    return reference.value
}

func ref_set__Ref_11Option__Typ(reference *ref_Option__Typ_x, value Option__Typ) struct{} {
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

type Tuple2_3Typ_16Vec_10SubstEntry struct {
    _0 Typ
    _1 *_goml_vec_SubstEntry
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

type CheckerState struct {
    gensym_counter *ref_int32_x
    current_level *ref_int32_x
}

type EnvEntry struct {
    name string
    ty Typ
}

type SubstEntry struct {
    name string
    ty Typ
}

type Ordering uint8

type Exp struct {
    _node *Exp_node
}

type Exp_node struct {
    _p0 string
    _p1 Exp
    _p2 Exp
    _tag uint8
}

func _goml_enum_tag_Exp(value Exp) uint8 {
    if value._node == nil {
        return 0
    }
    return value._node._tag
}

type Typ struct {
    _node *Typ_node
}

type Typ_node struct {
    _p1 string
    _p0 *ref_Tv_x
    _p2 Typ
    _p3 Typ
    _tag uint8
}

func _goml_enum_tag_Typ(value Typ) uint8 {
    if value._node == nil {
        return 0
    }
    return value._node._tag
}

type Tv interface {
    isTv()
}

type Unbound struct {
    _0 string
    _1 int32
}

func (_ Unbound) isTv() {}

type Link struct {
    _0 Typ
}

func (_ Link) isTv() {}

type Option__Typ struct {
    _p0 Typ
    _tag uint8
}

type _goml_m_Result_____o__q_____string struct {
    _p0 string
    _tag uint8
}

type Result__Typ__string struct {
    _p1 string
    _p0 Typ
    _tag uint8
}

func state_new() CheckerState {
    var t0 *ref_int32_x
    var inline2 int32 = 0
    var inline3 *ref_int32_x = ref__Ref_5int32(inline2)
    t0 = inline3
    var t1 *ref_int32_x
    var inline0 int32 = 1
    var inline1 *ref_int32_x = ref__Ref_5int32(inline0)
    t1 = inline1
    var t2 CheckerState = CheckerState{
        gensym_counter: t0,
        current_level: t1,
    }
    return t2
}

func reset_type_variables(st__0 CheckerState) struct{} {
    var inline2 *ref_int32_x = st__0.gensym_counter
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__i32(inline2, 0)
    var inline0 *ref_int32_x = st__0.current_level
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__i32(inline0, 1)
    return struct{}{}
}

func nth_letter(n__0 int32) rune {
    switch n__0 {
    case 0:
        return 97
    case 1:
        return 98
    case 2:
        return 99
    case 3:
        return 100
    case 4:
        return 101
    case 5:
        return 102
    case 6:
        return 103
    case 7:
        return 104
    case 8:
        return 105
    case 9:
        return 106
    case 10:
        return 107
    case 11:
        return 108
    case 12:
        return 109
    case 13:
        return 110
    case 14:
        return 111
    case 15:
        return 112
    case 16:
        return 113
    case 17:
        return 114
    case 18:
        return 115
    case 19:
        return 116
    case 20:
        return 117
    case 21:
        return 118
    case 22:
        return 119
    case 23:
        return 120
    case 24:
        return 121
    case 25:
        return 122
    default:
        return 97
    }
}

func gensym(st__0 CheckerState) string {
    var t0 *ref_int32_x = st__0.gensym_counter
    var n__0 int32
    var inline3 int32 = ref_get__Ref_5int32(t0)
    n__0 = inline3
    var t1 *ref_int32_x = st__0.gensym_counter
    var t2 int32 = n__0 + 1
    ref_set__Ref_5int32(t1, t2)
    var t3 bool = n__0 < 26
    if t3 {
        var t4 rune = nth_letter(n__0)
        var inline0 string = char_to_string(t4)
        return inline0
    } else {
        var t5 string
        var inline1 string = __goml_builtin_int32_to_string(n__0)
        t5 = inline1
        var t6 string = "t" + t5
        return t6
    }
}

func typ_is_arrow(ty__0 Typ) bool {
    switch _goml_enum_tag_Typ(ty__0) {
    case 0:
        var x0 *ref_Tv_x = ty__0._node._p0
        var mtmp0 Tv
        var inline0 Tv = ref_get__Ref_2Tv(x0)
        mtmp0 = inline0
        switch mtmp0.(type) {
        case Link:
            var x1 Typ = mtmp0.(Link)._0
            var t0 bool = typ_is_arrow(x1)
            return t0
        default:
            return false
        }
    case 2:
        return true
    default:
        return false
    }
}

func typ_to_string(ty__0 Typ) string {
    switch _goml_enum_tag_Typ(ty__0) {
    case 0:
        var x0 *ref_Tv_x = ty__0._node._p0
        var mtmp0 Tv
        var inline0 Tv = ref_get__Ref_2Tv(x0)
        mtmp0 = inline0
        switch mtmp0.(type) {
        case Unbound:
            var x1 string = mtmp0.(Unbound)._0
            var t0 string = "'" + x1
            return t0
        case Link:
            var x2 Typ = mtmp0.(Link)._0
            var t1 string = typ_to_string(x2)
            return t1
        default:
            panic("non-exhaustive match")
        }
    case 1:
        var x3 string = ty__0._node._p1
        var t2 string = "'" + x3
        return t2
    case 2:
        var x4 Typ = ty__0._node._p2
        var x5 Typ = ty__0._node._p3
        var t3 bool = typ_is_arrow(x4)
        var jp0 string
        if t3 {
            var t6 string = typ_to_string(x4)
            var t7 string = "(" + t6
            var t8 string = t7 + ")"
            jp0 = t8
        } else {
            var t9 string = typ_to_string(x4)
            jp0 = t9
        }
        var s2__0 string = typ_to_string(x5)
        var t4 string = jp0 + " -> "
        var t5 string = t4 + s2__0
        return t5
    default:
        panic("non-exhaustive match")
    }
}

func env_empty() *_goml_vec_EnvEntry {
    var inline0 *_goml_vec_EnvEntry = vec_new__Vec_8EnvEntry()
    return inline0
}

func env_lookup(env__0 *_goml_vec_EnvEntry, name__0 string) Option__Typ {
    var t0 int
    var inline13 int = vec_len__Vec_8EnvEntry(env__0)
    t0 = inline13
    var t1 int = t0 - 1
    var i__0 *ref_int_x
    var inline12 *ref_int_x = ref__Ref_3int(t1)
    i__0 = inline12
    var found__0 *ref_Option__Typ_x
    var inline11 *ref_Option__Typ_x = ref__Ref_11Option__Typ(Option__Typ{
        _tag: 0,
    })
    found__0 = inline11
    var done__0 *ref_bool_x
    var inline9 bool = false
    var inline10 *ref_bool_x = ref__Ref_4bool(inline9)
    done__0 = inline10
    Loop_loop0:
    for {
        var t2 bool
        var inline8 bool = ref_get__Ref_4bool(done__0)
        t2 = inline8
        var t3 bool = !t2
        var jp0 bool
        if t3 {
            var t11 int
            var inline7 int = ref_get__Ref_3int(i__0)
            t11 = inline7
            var t12 bool = t11 >= 0
            jp0 = t12
        } else {
            jp0 = false
        }
        if jp0 {
            var t4 int
            var inline6 int = ref_get__Ref_3int(i__0)
            t4 = inline6
            var entry__0 EnvEntry = vec_get__Vec_8EnvEntry(env__0, t4)
            var t5 string = entry__0.name
            var t6 bool = t5 == name__0
            if t6 {
                var t7 Typ = entry__0.ty
                var t8 Option__Typ = Option__Typ{
                    _p0: t7,
                    _tag: 1,
                }
                ref_set__Ref_11Option__Typ(found__0, t8)
                var inline1 bool = true
                ref_set__Ref_4bool(done__0, inline1)
                continue
            } else {
                var t9 int
                var inline5 int = ref_get__Ref_3int(i__0)
                t9 = inline5
                var t10 int = t9 - 1
                ref_set__Ref_3int(i__0, t10)
                continue
            }
        } else {
            break Loop_loop0
        }
    }
    var inline0 Option__Typ = ref_get__Ref_11Option__Typ(found__0)
    return inline0
}

func subst_lookup(subst__0 *_goml_vec_SubstEntry, name__0 string) Option__Typ {
    var t0 int
    var inline13 int = vec_len__Vec_10SubstEntry(subst__0)
    t0 = inline13
    var t1 int = t0 - 1
    var i__0 *ref_int_x
    var inline12 *ref_int_x = ref__Ref_3int(t1)
    i__0 = inline12
    var found__0 *ref_Option__Typ_x
    var inline11 *ref_Option__Typ_x = ref__Ref_11Option__Typ(Option__Typ{
        _tag: 0,
    })
    found__0 = inline11
    var done__0 *ref_bool_x
    var inline9 bool = false
    var inline10 *ref_bool_x = ref__Ref_4bool(inline9)
    done__0 = inline10
    Loop_loop0:
    for {
        var t2 bool
        var inline8 bool = ref_get__Ref_4bool(done__0)
        t2 = inline8
        var t3 bool = !t2
        var jp0 bool
        if t3 {
            var t11 int
            var inline7 int = ref_get__Ref_3int(i__0)
            t11 = inline7
            var t12 bool = t11 >= 0
            jp0 = t12
        } else {
            jp0 = false
        }
        if jp0 {
            var t4 int
            var inline6 int = ref_get__Ref_3int(i__0)
            t4 = inline6
            var entry__0 SubstEntry = vec_get__Vec_10SubstEntry(subst__0, t4)
            var t5 string = entry__0.name
            var t6 bool = t5 == name__0
            if t6 {
                var t7 Typ = entry__0.ty
                var t8 Option__Typ = Option__Typ{
                    _p0: t7,
                    _tag: 1,
                }
                ref_set__Ref_11Option__Typ(found__0, t8)
                var inline1 bool = true
                ref_set__Ref_4bool(done__0, inline1)
                continue
            } else {
                var t9 int
                var inline5 int = ref_get__Ref_3int(i__0)
                t9 = inline5
                var t10 int = t9 - 1
                ref_set__Ref_3int(i__0, t10)
                continue
            }
        } else {
            break Loop_loop0
        }
    }
    var inline0 Option__Typ = ref_get__Ref_11Option__Typ(found__0)
    return inline0
}

func occurs(st__0 CheckerState, tvr__0 *ref_Tv_x, ty__0 Typ) _goml_m_Result_____o__q_____string {
    switch _goml_enum_tag_Typ(ty__0) {
    case 0:
        var x0 *ref_Tv_x = ty__0._node._p0
        var t0 bool = ptr_eq__Ref_2Tv(tvr__0, x0)
        if t0 {
            var t1 _goml_m_Result_____o__q_____string = _goml_m_Result_____o__q_____string{
                _p0: "occurs check",
                _tag: 1,
            }
            return t1
        } else {
            var mtmp0 Tv
            var inline3 Tv = ref_get__Ref_2Tv(x0)
            mtmp0 = inline3
            switch mtmp0.(type) {
            case Unbound:
                var x1 string = mtmp0.(Unbound)._0
                var x2 int32 = mtmp0.(Unbound)._1
                var mtmp1 Tv
                var inline2 Tv = ref_get__Ref_2Tv(tvr__0)
                mtmp1 = inline2
                var jp0 int32
                switch mtmp1.(type) {
                case Unbound:
                    var x3 int32 = mtmp1.(Unbound)._1
                    var inline1 bool = x3 < x2
                    if inline1 {
                        jp0 = x3
                    } else {
                        jp0 = x2
                    }
                default:
                    jp0 = x2
                }
                var t2 Tv = Unbound{
                    _0: x1,
                    _1: jp0,
                }
                ref_set__Ref_2Tv(x0, t2)
                var t3 _goml_m_Result_____o__q_____string = _goml_m_Result_____o__q_____string{
                    _tag: 0,
                }
                return t3
            case Link:
                var x4 Typ = mtmp0.(Link)._0
                var t4 _goml_m_Result_____o__q_____string = occurs(st__0, tvr__0, x4)
                return t4
            default:
                panic("non-exhaustive match")
            }
        }
    case 2:
        var x5 Typ = ty__0._node._p2
        var x6 Typ = ty__0._node._p3
        var mtmp3 _goml_m_Result_____o__q_____string = occurs(st__0, tvr__0, x5)
        switch mtmp3._tag {
        case 0:
            var t5 _goml_m_Result_____o__q_____string = occurs(st__0, tvr__0, x6)
            return t5
        case 1:
            var x7 string = mtmp3._p0
            var t6 _goml_m_Result_____o__q_____string = _goml_m_Result_____o__q_____string{
                _p0: x7,
                _tag: 1,
            }
            return t6
        default:
            panic("non-exhaustive match")
        }
    default:
        var t7 _goml_m_Result_____o__q_____string = _goml_m_Result_____o__q_____string{
            _tag: 0,
        }
        return t7
    }
}

func unify(st__0 CheckerState, t1__0 Typ, t2__0 Typ) _goml_m_Result_____o__q_____string {
    switch _goml_enum_tag_Typ(t2__0) {
    case 0:
        var x0 *ref_Tv_x = t2__0._node._p0
        switch _goml_enum_tag_Typ(t1__0) {
        case 0:
            var x1 *ref_Tv_x = t1__0._node._p0
            var t0 bool = ptr_eq__Ref_2Tv(x1, x0)
            if t0 {
                var t1 _goml_m_Result_____o__q_____string = _goml_m_Result_____o__q_____string{
                    _tag: 0,
                }
                return t1
            } else {
                var mtmp0 Tv
                var inline2 Tv = ref_get__Ref_2Tv(x1)
                mtmp0 = inline2
                switch mtmp0.(type) {
                case Unbound:
                    var mtmp1 Tv
                    var inline1 Tv = ref_get__Ref_2Tv(x0)
                    mtmp1 = inline1
                    switch mtmp1.(type) {
                    case Unbound:
                        var t2 Typ = Typ{
                            _node: &Typ_node{
                                _p0: x0,
                                _tag: 0,
                            },
                        }
                        var mtmp2 _goml_m_Result_____o__q_____string = occurs(st__0, x1, t2)
                        switch mtmp2._tag {
                        case 0:
                            var t3 Typ = Typ{
                                _node: &Typ_node{
                                    _p0: x0,
                                    _tag: 0,
                                },
                            }
                            var t4 Tv = Link{
                                _0: t3,
                            }
                            ref_set__Ref_2Tv(x1, t4)
                            var t5 _goml_m_Result_____o__q_____string = _goml_m_Result_____o__q_____string{
                                _tag: 0,
                            }
                            return t5
                        case 1:
                            var x2 string = mtmp2._p0
                            var t6 _goml_m_Result_____o__q_____string = _goml_m_Result_____o__q_____string{
                                _p0: x2,
                                _tag: 1,
                            }
                            return t6
                        default:
                            panic("non-exhaustive match")
                        }
                    case Link:
                        var x3 Typ = mtmp1.(Link)._0
                        var t7 Typ = Typ{
                            _node: &Typ_node{
                                _p0: x1,
                                _tag: 0,
                            },
                        }
                        var t8 _goml_m_Result_____o__q_____string = unify(st__0, t7, x3)
                        return t8
                    default:
                        panic("non-exhaustive match")
                    }
                case Link:
                    var x4 Typ = mtmp0.(Link)._0
                    var t9 Typ = Typ{
                        _node: &Typ_node{
                            _p0: x0,
                            _tag: 0,
                        },
                    }
                    var t10 _goml_m_Result_____o__q_____string = unify(st__0, x4, t9)
                    return t10
                default:
                    panic("non-exhaustive match")
                }
            }
        default:
            var mtmp4 Tv
            var inline4 Tv = ref_get__Ref_2Tv(x0)
            mtmp4 = inline4
            switch mtmp4.(type) {
            case Unbound:
                var mtmp5 _goml_m_Result_____o__q_____string = occurs(st__0, x0, t1__0)
                switch mtmp5._tag {
                case 0:
                    var t11 Tv = Link{
                        _0: t1__0,
                    }
                    ref_set__Ref_2Tv(x0, t11)
                    var t12 _goml_m_Result_____o__q_____string = _goml_m_Result_____o__q_____string{
                        _tag: 0,
                    }
                    return t12
                case 1:
                    var x5 string = mtmp5._p0
                    var t13 _goml_m_Result_____o__q_____string = _goml_m_Result_____o__q_____string{
                        _p0: x5,
                        _tag: 1,
                    }
                    return t13
                default:
                    panic("non-exhaustive match")
                }
            case Link:
                var x6 Typ = mtmp4.(Link)._0
                var t14 _goml_m_Result_____o__q_____string = unify(st__0, t1__0, x6)
                return t14
            default:
                panic("non-exhaustive match")
            }
        }
    case 2:
        var x7 Typ = t2__0._node._p2
        var x8 Typ = t2__0._node._p3
        switch _goml_enum_tag_Typ(t1__0) {
        case 0:
            var x9 *ref_Tv_x = t1__0._node._p0
            var mtmp7 Tv
            var inline6 Tv = ref_get__Ref_2Tv(x9)
            mtmp7 = inline6
            switch mtmp7.(type) {
            case Unbound:
                var mtmp8 _goml_m_Result_____o__q_____string = occurs(st__0, x9, t2__0)
                switch mtmp8._tag {
                case 0:
                    var t15 Tv = Link{
                        _0: t2__0,
                    }
                    ref_set__Ref_2Tv(x9, t15)
                    var t16 _goml_m_Result_____o__q_____string = _goml_m_Result_____o__q_____string{
                        _tag: 0,
                    }
                    return t16
                case 1:
                    var x10 string = mtmp8._p0
                    var t17 _goml_m_Result_____o__q_____string = _goml_m_Result_____o__q_____string{
                        _p0: x10,
                        _tag: 1,
                    }
                    return t17
                default:
                    panic("non-exhaustive match")
                }
            case Link:
                var x11 Typ = mtmp7.(Link)._0
                var t18 _goml_m_Result_____o__q_____string = unify(st__0, x11, t2__0)
                return t18
            default:
                panic("non-exhaustive match")
            }
        case 2:
            var x12 Typ = t1__0._node._p2
            var x13 Typ = t1__0._node._p3
            var mtmp10 _goml_m_Result_____o__q_____string = unify(st__0, x12, x7)
            switch mtmp10._tag {
            case 0:
                var t19 _goml_m_Result_____o__q_____string = unify(st__0, x13, x8)
                return t19
            case 1:
                var x14 string = mtmp10._p0
                var t20 _goml_m_Result_____o__q_____string = _goml_m_Result_____o__q_____string{
                    _p0: x14,
                    _tag: 1,
                }
                return t20
            default:
                panic("non-exhaustive match")
            }
        default:
            var t21 _goml_m_Result_____o__q_____string = _goml_m_Result_____o__q_____string{
                _p0: "unify error",
                _tag: 1,
            }
            return t21
        }
    default:
        switch _goml_enum_tag_Typ(t1__0) {
        case 0:
            var x15 *ref_Tv_x = t1__0._node._p0
            var mtmp11 Tv
            var inline8 Tv = ref_get__Ref_2Tv(x15)
            mtmp11 = inline8
            switch mtmp11.(type) {
            case Unbound:
                var mtmp12 _goml_m_Result_____o__q_____string = occurs(st__0, x15, t2__0)
                switch mtmp12._tag {
                case 0:
                    var t22 Tv = Link{
                        _0: t2__0,
                    }
                    ref_set__Ref_2Tv(x15, t22)
                    var t23 _goml_m_Result_____o__q_____string = _goml_m_Result_____o__q_____string{
                        _tag: 0,
                    }
                    return t23
                case 1:
                    var x16 string = mtmp12._p0
                    var t24 _goml_m_Result_____o__q_____string = _goml_m_Result_____o__q_____string{
                        _p0: x16,
                        _tag: 1,
                    }
                    return t24
                default:
                    panic("non-exhaustive match")
                }
            case Link:
                var x17 Typ = mtmp11.(Link)._0
                var t25 _goml_m_Result_____o__q_____string = unify(st__0, x17, t2__0)
                return t25
            default:
                panic("non-exhaustive match")
            }
        default:
            var t26 _goml_m_Result_____o__q_____string = _goml_m_Result_____o__q_____string{
                _p0: "unify error",
                _tag: 1,
            }
            return t26
        }
    }
}

func gen(st__0 CheckerState, ty__0 Typ) Typ {
    switch _goml_enum_tag_Typ(ty__0) {
    case 0:
        var x0 *ref_Tv_x = ty__0._node._p0
        var mtmp0 Tv
        var inline1 Tv = ref_get__Ref_2Tv(x0)
        mtmp0 = inline1
        switch mtmp0.(type) {
        case Unbound:
            var x1 string = mtmp0.(Unbound)._0
            var x2 int32 = mtmp0.(Unbound)._1
            var t0 *ref_int32_x = st__0.current_level
            var cur__0 int32
            var inline0 int32 = ref_get__Ref_5int32(t0)
            cur__0 = inline0
            var t1 bool = x2 > cur__0
            if t1 {
                var t2 Typ = Typ{
                    _node: &Typ_node{
                        _p1: x1,
                        _tag: 1,
                    },
                }
                return t2
            } else {
                var t3 Typ = Typ{
                    _node: &Typ_node{
                        _p0: x0,
                        _tag: 0,
                    },
                }
                return t3
            }
        case Link:
            var x3 Typ = mtmp0.(Link)._0
            var t4 Typ = gen(st__0, x3)
            return t4
        default:
            panic("non-exhaustive match")
        }
    case 2:
        var x4 Typ = ty__0._node._p2
        var x5 Typ = ty__0._node._p3
        var t5 Typ = gen(st__0, x4)
        var t6 Typ = gen(st__0, x5)
        var t7 Typ = Typ{
            _node: &Typ_node{
                _p2: t5,
                _p3: t6,
                _tag: 2,
            },
        }
        return t7
    default:
        return ty__0
    }
}

func inst_loop(st__0 CheckerState, subst__0 *_goml_vec_SubstEntry, ty__0 Typ) Tuple2_3Typ_16Vec_10SubstEntry {
    switch _goml_enum_tag_Typ(ty__0) {
    case 0:
        var x0 *ref_Tv_x = ty__0._node._p0
        var mtmp0 Tv
        var inline0 Tv = ref_get__Ref_2Tv(x0)
        mtmp0 = inline0
        switch mtmp0.(type) {
        case Link:
            var x1 Typ = mtmp0.(Link)._0
            var t0 Tuple2_3Typ_16Vec_10SubstEntry = inst_loop(st__0, subst__0, x1)
            return t0
        default:
            var t1 Typ = Typ{
                _node: &Typ_node{
                    _p0: x0,
                    _tag: 0,
                },
            }
            var t2 Tuple2_3Typ_16Vec_10SubstEntry = Tuple2_3Typ_16Vec_10SubstEntry{
                _0: t1,
                _1: subst__0,
            }
            return t2
        }
    case 1:
        var x2 string = ty__0._node._p1
        var mtmp1 Option__Typ = subst_lookup(subst__0, x2)
        switch mtmp1._tag {
        case 0:
            var tv__0 Typ
            var inline1 string = gensym(st__0)
            var inline2 *ref_int32_x = st__0.current_level
            var inline3 int32 = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__i32(inline2)
            var inline4 Tv = Unbound{
                _0: inline1,
                _1: inline3,
            }
            var inline5 *ref_Tv_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__Tv(inline4)
            var inline6 Typ = Typ{
                _node: &Typ_node{
                    _p0: inline5,
                    _tag: 0,
                },
            }
            tv__0 = inline6
            var t3 SubstEntry = SubstEntry{
                name: x2,
                ty: tv__0,
            }
            var new_subst__0 *_goml_vec_SubstEntry = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pushed____T__SubstEntry(subst__0, t3)
            var t4 Tuple2_3Typ_16Vec_10SubstEntry = Tuple2_3Typ_16Vec_10SubstEntry{
                _0: tv__0,
                _1: new_subst__0,
            }
            return t4
        case 1:
            var x3 Typ = mtmp1._p0
            var t5 Tuple2_3Typ_16Vec_10SubstEntry = Tuple2_3Typ_16Vec_10SubstEntry{
                _0: x3,
                _1: subst__0,
            }
            return t5
        default:
            panic("non-exhaustive match")
        }
    case 2:
        var x4 Typ = ty__0._node._p2
        var x5 Typ = ty__0._node._p3
        var mtmp2 Tuple2_3Typ_16Vec_10SubstEntry = inst_loop(st__0, subst__0, x4)
        var x6 Typ = mtmp2._0
        var x7 *_goml_vec_SubstEntry = mtmp2._1
        var mtmp3 Tuple2_3Typ_16Vec_10SubstEntry = inst_loop(st__0, x7, x5)
        var x8 Typ = mtmp3._0
        var x9 *_goml_vec_SubstEntry = mtmp3._1
        var t6 Typ = Typ{
            _node: &Typ_node{
                _p2: x6,
                _p3: x8,
                _tag: 2,
            },
        }
        var t7 Tuple2_3Typ_16Vec_10SubstEntry = Tuple2_3Typ_16Vec_10SubstEntry{
            _0: t6,
            _1: x9,
        }
        return t7
    default:
        panic("non-exhaustive match")
    }
}

func typeof(st__0 CheckerState, env__0 *_goml_vec_EnvEntry, e__0 Exp) Result__Typ__string {
    switch _goml_enum_tag_Exp(e__0) {
    case 0:
        var x0 string = e__0._node._p0
        var mtmp0 Option__Typ = env_lookup(env__0, x0)
        switch mtmp0._tag {
        case 0:
            var t0 Result__Typ__string = Result__Typ__string{
                _p1: "unbound var",
                _tag: 1,
            }
            return t0
        case 1:
            var x1 Typ = mtmp0._p0
            var t1 Typ
            var inline0 *_goml_vec_SubstEntry = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_new____T__SubstEntry()
            var inline1 Tuple2_3Typ_16Vec_10SubstEntry = inst_loop(st__0, inline0, x1)
            var inline2 Typ = inline1._0
            t1 = inline2
            var t2 Result__Typ__string = Result__Typ__string{
                _p0: t1,
                _tag: 0,
            }
            return t2
        default:
            panic("non-exhaustive match")
        }
    case 1:
        var x2 Exp = e__0._node._p1
        var x3 Exp = e__0._node._p2
        var mtmp1 Result__Typ__string = typeof(st__0, env__0, x2)
        switch mtmp1._tag {
        case 0:
            var x4 Typ = mtmp1._p0
            var mtmp2 Result__Typ__string = typeof(st__0, env__0, x3)
            switch mtmp2._tag {
            case 0:
                var x5 Typ = mtmp2._p0
                var ty_res__0 Typ
                var inline3 string = gensym(st__0)
                var inline4 *ref_int32_x = st__0.current_level
                var inline5 int32 = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__i32(inline4)
                var inline6 Tv = Unbound{
                    _0: inline3,
                    _1: inline5,
                }
                var inline7 *ref_Tv_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__Tv(inline6)
                var inline8 Typ = Typ{
                    _node: &Typ_node{
                        _p0: inline7,
                        _tag: 0,
                    },
                }
                ty_res__0 = inline8
                var arrow__0 Typ = Typ{
                    _node: &Typ_node{
                        _p2: x5,
                        _p3: ty_res__0,
                        _tag: 2,
                    },
                }
                var mtmp3 _goml_m_Result_____o__q_____string = unify(st__0, x4, arrow__0)
                switch mtmp3._tag {
                case 0:
                    var t3 Result__Typ__string = Result__Typ__string{
                        _p0: ty_res__0,
                        _tag: 0,
                    }
                    return t3
                case 1:
                    var x6 string = mtmp3._p0
                    var t4 Result__Typ__string = Result__Typ__string{
                        _p1: x6,
                        _tag: 1,
                    }
                    return t4
                default:
                    panic("non-exhaustive match")
                }
            case 1:
                var x7 string = mtmp2._p1
                var t5 Result__Typ__string = Result__Typ__string{
                    _p1: x7,
                    _tag: 1,
                }
                return t5
            default:
                panic("non-exhaustive match")
            }
        case 1:
            var x8 string = mtmp1._p1
            var t6 Result__Typ__string = Result__Typ__string{
                _p1: x8,
                _tag: 1,
            }
            return t6
        default:
            panic("non-exhaustive match")
        }
    case 2:
        var x9 string = e__0._node._p0
        var x10 Exp = e__0._node._p1
        var ty_x__0 Typ
        var inline9 string = gensym(st__0)
        var inline10 *ref_int32_x = st__0.current_level
        var inline11 int32 = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__i32(inline10)
        var inline12 Tv = Unbound{
            _0: inline9,
            _1: inline11,
        }
        var inline13 *ref_Tv_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__Tv(inline12)
        var inline14 Typ = Typ{
            _node: &Typ_node{
                _p0: inline13,
                _tag: 0,
            },
        }
        ty_x__0 = inline14
        var t7 EnvEntry = EnvEntry{
            name: x9,
            ty: ty_x__0,
        }
        var env2__0 *_goml_vec_EnvEntry = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pushed____T__EnvEntry(env__0, t7)
        var mtmp4 Result__Typ__string = typeof(st__0, env2__0, x10)
        switch mtmp4._tag {
        case 0:
            var x11 Typ = mtmp4._p0
            var t8 Typ = Typ{
                _node: &Typ_node{
                    _p2: ty_x__0,
                    _p3: x11,
                    _tag: 2,
                },
            }
            var t9 Result__Typ__string = Result__Typ__string{
                _p0: t8,
                _tag: 0,
            }
            return t9
        case 1:
            var x12 string = mtmp4._p1
            var t10 Result__Typ__string = Result__Typ__string{
                _p1: x12,
                _tag: 1,
            }
            return t10
        default:
            panic("non-exhaustive match")
        }
    case 3:
        var x13 string = e__0._node._p0
        var x14 Exp = e__0._node._p1
        var x15 Exp = e__0._node._p2
        var inline20 *ref_int32_x = st__0.current_level
        var inline21 int32 = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__i32(inline20)
        var inline22 *ref_int32_x = st__0.current_level
        var inline23 int32 = inline21 + 1
        _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__i32(inline22, inline23)
        var ty_e__0 Result__Typ__string = typeof(st__0, env__0, x14)
        var inline15 *ref_int32_x = st__0.current_level
        var inline16 int32 = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__i32(inline15)
        var inline17 *ref_int32_x = st__0.current_level
        var inline18 int32 = inline16 - 1
        _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__i32(inline17, inline18)
        switch ty_e__0._tag {
        case 0:
            var x16 Typ = ty_e__0._p0
            var t11 Typ = gen(st__0, x16)
            var t12 EnvEntry = EnvEntry{
                name: x13,
                ty: t11,
            }
            var env2__1 *_goml_vec_EnvEntry = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pushed____T__EnvEntry(env__0, t12)
            var t13 Result__Typ__string = typeof(st__0, env2__1, x15)
            return t13
        case 1:
            var x17 string = ty_e__0._p1
            var t14 Result__Typ__string = Result__Typ__string{
                _p1: x17,
                _tag: 1,
            }
            return t14
        default:
            panic("non-exhaustive match")
        }
    default:
        panic("non-exhaustive match")
    }
}

func exp_var(name__0 string) Exp {
    var t0 Exp = Exp{
        _node: &Exp_node{
            _p0: name__0,
            _tag: 0,
        },
    }
    return t0
}

func exp_lam(name__0 string, body__0 Exp) Exp {
    var t0 Exp = Exp{
        _node: &Exp_node{
            _p0: name__0,
            _p1: body__0,
            _tag: 2,
        },
    }
    return t0
}

func exp_app(a__0 Exp, b__0 Exp) Exp {
    var t0 Exp = Exp{
        _node: &Exp_node{
            _p1: a__0,
            _p2: b__0,
            _tag: 1,
        },
    }
    return t0
}

func exp_let(name__0 string, a__0 Exp, b__0 Exp) Exp {
    var t0 Exp = Exp{
        _node: &Exp_node{
            _p0: name__0,
            _p1: a__0,
            _p2: b__0,
            _tag: 3,
        },
    }
    return t0
}

func show_result(label__0 string, res__0 Result__Typ__string) struct{} {
    switch res__0._tag {
    case 0:
        var x0 Typ = res__0._p0
        var t0 string = label__0 + ": "
        var t1 string = typ_to_string(x0)
        var t2 string = t0 + t1
        var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t2)
        _goml_runtime_core_string_println(inline0)
        return struct{}{}
    case 1:
        var x1 string = res__0._p1
        var t3 string = label__0 + ": "
        var t4 string = t3 + x1
        var inline2 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t4)
        _goml_runtime_core_string_println(inline2)
        return struct{}{}
    default:
        panic("non-exhaustive match")
    }
}

func main0() struct{} {
    var st__0 CheckerState = state_new()
    var t0 Exp = exp_var("x")
    var id__0 Exp = exp_lam("x", t0)
    var t1 Exp = exp_var("x")
    var t2 Exp = exp_var("y")
    var t3 Exp = exp_app(t1, t2)
    var t4 Exp = exp_lam("y", t3)
    var c1__0 Exp = exp_lam("x", t4)
    reset_type_variables(st__0)
    var t5 *_goml_vec_EnvEntry = env_empty()
    var t6 Result__Typ__string = typeof(st__0, t5, id__0)
    show_result("id", t6)
    reset_type_variables(st__0)
    var t7 *_goml_vec_EnvEntry = env_empty()
    var t8 Result__Typ__string = typeof(st__0, t7, c1__0)
    show_result("c1", t8)
    reset_type_variables(st__0)
    var t9 *_goml_vec_EnvEntry = env_empty()
    var t10 Exp = exp_var("x")
    var t11 Exp = exp_let("x", c1__0, t10)
    var t12 Result__Typ__string = typeof(st__0, t9, t11)
    show_result("let_x_c1_x", t12)
    reset_type_variables(st__0)
    var t13 *_goml_vec_EnvEntry = env_empty()
    var t14 Exp = exp_var("z")
    var t15 Exp = exp_lam("z", t14)
    var t16 Exp = exp_var("y")
    var t17 Exp = exp_let("y", t15, t16)
    var t18 Result__Typ__string = typeof(st__0, t13, t17)
    show_result("let_y_id_y", t18)
    reset_type_variables(st__0)
    var t19 *_goml_vec_EnvEntry = env_empty()
    var t20 Exp = exp_var("z")
    var t21 Exp = exp_lam("z", t20)
    var t22 Exp = exp_var("y")
    var t23 Exp = exp_let("y", t21, t22)
    var t24 Exp = exp_lam("x", t23)
    var t25 Result__Typ__string = typeof(st__0, t19, t24)
    show_result("lam_x_let_y_id_y", t25)
    reset_type_variables(st__0)
    var t26 *_goml_vec_EnvEntry = env_empty()
    var t27 Exp = exp_var("z")
    var t28 Exp = exp_lam("z", t27)
    var t29 Exp = exp_var("y")
    var t30 Exp = exp_var("x")
    var t31 Exp = exp_app(t29, t30)
    var t32 Exp = exp_let("y", t28, t31)
    var t33 Exp = exp_lam("x", t32)
    var t34 Result__Typ__string = typeof(st__0, t26, t33)
    show_result("lam_x_let_y_id_yx", t34)
    reset_type_variables(st__0)
    var t35 *_goml_vec_EnvEntry = env_empty()
    var t36 Exp = exp_var("x")
    var t37 Exp = exp_var("x")
    var t38 Exp = exp_app(t36, t37)
    var t39 Exp = exp_lam("x", t38)
    var t40 Result__Typ__string = typeof(st__0, t35, t39)
    show_result("self_apply", t40)
    reset_type_variables(st__0)
    var t41 *_goml_vec_EnvEntry = env_empty()
    var t42 Exp = exp_var("x")
    var t43 Exp = exp_var("x")
    var t44 Exp = exp_let("x", t42, t43)
    var t45 Result__Typ__string = typeof(st__0, t41, t44)
    show_result("unbound_var", t45)
    reset_type_variables(st__0)
    var t46 *_goml_vec_EnvEntry = env_empty()
    var t47 Exp = exp_var("y")
    var t48 Exp = exp_var("y")
    var t49 Exp = exp_var("z")
    var t50 Exp = exp_app(t48, t49)
    var t51 Exp = exp_lam("z", t50)
    var t52 Exp = exp_app(t47, t51)
    var t53 Exp = exp_lam("y", t52)
    var t54 Result__Typ__string = typeof(st__0, t46, t53)
    show_result("max_heiber", t54)
    reset_type_variables(st__0)
    var t55 *_goml_vec_EnvEntry = env_empty()
    var t56 Exp = exp_var("k")
    var t57 Exp = exp_var("k")
    var t58 Exp = exp_var("x")
    var t59 Exp = exp_app(t57, t58)
    var t60 Exp = exp_var("y")
    var t61 Exp = exp_app(t59, t60)
    var t62 Exp = exp_app(t56, t61)
    var t63 Exp = exp_var("k")
    var t64 Exp = exp_var("y")
    var t65 Exp = exp_app(t63, t64)
    var t66 Exp = exp_var("x")
    var t67 Exp = exp_app(t65, t66)
    var t68 Exp = exp_app(t62, t67)
    var t69 Exp = exp_lam("k", t68)
    var t70 Exp = exp_lam("y", t69)
    var t71 Exp = exp_lam("x", t70)
    var t72 Result__Typ__string = typeof(st__0, t55, t71)
    show_result("kirang", t72)
    reset_type_variables(st__0)
    var t73 *_goml_vec_EnvEntry = env_empty()
    var t74 Exp = exp_var("id")
    var t75 Exp = exp_var("id")
    var t76 Exp = exp_app(t74, t75)
    var t77 Exp = exp_let("id", id__0, t76)
    var t78 Result__Typ__string = typeof(st__0, t73, t77)
    show_result("let_id_idid", t78)
    reset_type_variables(st__0)
    var t79 *_goml_vec_EnvEntry = env_empty()
    var t80 Exp = exp_var("x")
    var t81 Exp = exp_app(t80, id__0)
    var t82 Exp = exp_var("z")
    var t83 Exp = exp_let("z", t81, t82)
    var t84 Exp = exp_var("y")
    var t85 Exp = exp_let("y", t83, t84)
    var t86 Exp = exp_let("x", c1__0, t85)
    var t87 Result__Typ__string = typeof(st__0, t79, t86)
    show_result("nested_lets", t87)
    reset_type_variables(st__0)
    var t88 *_goml_vec_EnvEntry = env_empty()
    var t89 Exp = exp_var("x")
    var t90 Exp = exp_var("y")
    var t91 Exp = exp_app(t89, t90)
    var t92 Exp = exp_var("y")
    var t93 Exp = exp_var("x")
    var t94 Exp = exp_app(t92, t93)
    var t95 Exp = exp_lam("x", t94)
    var t96 Exp = exp_let("x", t91, t95)
    var t97 Exp = exp_lam("y", t96)
    var t98 Exp = exp_lam("x", t97)
    var t99 Result__Typ__string = typeof(st__0, t88, t98)
    show_result("fun_x_fun_y_let_x_xy_fun_x_yx", t99)
    reset_type_variables(st__0)
    var t100 *_goml_vec_EnvEntry = env_empty()
    var t101 Exp = exp_var("x")
    var t102 Exp = exp_var("y")
    var t103 Exp = exp_let("y", t101, t102)
    var t104 Exp = exp_lam("x", t103)
    var t105 Result__Typ__string = typeof(st__0, t100, t104)
    show_result("sound_gen_1", t105)
    reset_type_variables(st__0)
    var t106 *_goml_vec_EnvEntry = env_empty()
    var t107 Exp = exp_var("x")
    var t108 Exp = exp_lam("z", t107)
    var t109 Exp = exp_var("y")
    var t110 Exp = exp_let("y", t108, t109)
    var t111 Exp = exp_lam("x", t110)
    var t112 Result__Typ__string = typeof(st__0, t106, t111)
    show_result("sound_gen_2", t112)
    reset_type_variables(st__0)
    var t113 *_goml_vec_EnvEntry = env_empty()
    var t114 Exp = exp_var("x")
    var t115 Exp = exp_var("z")
    var t116 Exp = exp_app(t114, t115)
    var t117 Exp = exp_lam("z", t116)
    var t118 Exp = exp_var("y")
    var t119 Exp = exp_let("y", t117, t118)
    var t120 Exp = exp_lam("x", t119)
    var t121 Result__Typ__string = typeof(st__0, t113, t120)
    show_result("sound_gen_3", t121)
    reset_type_variables(st__0)
    var t122 *_goml_vec_EnvEntry = env_empty()
    var t123 Exp = exp_var("x")
    var t124 Exp = exp_var("y")
    var t125 Exp = exp_app(t123, t124)
    var t126 Exp = exp_var("x")
    var t127 Exp = exp_var("y")
    var t128 Exp = exp_app(t126, t127)
    var t129 Exp = exp_let("x", t125, t128)
    var t130 Exp = exp_lam("y", t129)
    var t131 Exp = exp_lam("x", t130)
    var t132 Result__Typ__string = typeof(st__0, t122, t131)
    show_result("double_apply", t132)
    reset_type_variables(st__0)
    var t133 *_goml_vec_EnvEntry = env_empty()
    var t134 Exp = exp_var("x")
    var t135 Exp = exp_var("y")
    var t136 Exp = exp_var("y")
    var t137 Exp
    var inline39 Exp = Exp{
        _node: &Exp_node{
            _p1: t135,
            _p2: t136,
            _tag: 1,
        },
    }
    t137 = inline39
    var t138 Exp
    var inline37 string = "y"
    var inline38 Exp = Exp{
        _node: &Exp_node{
            _p0: inline37,
            _p1: t134,
            _p2: t137,
            _tag: 3,
        },
    }
    t138 = inline38
    var t139 Exp
    var inline35 string = "x"
    var inline36 Exp = Exp{
        _node: &Exp_node{
            _p0: inline35,
            _p1: t138,
            _tag: 2,
        },
    }
    t139 = inline36
    var t140 Result__Typ__string = typeof(st__0, t133, t139)
    show_result("sound_gen_occurs", t140)
    var inline33 *ref_int32_x = st__0.gensym_counter
    _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__i32(inline33, 0)
    var t141 *_goml_vec_EnvEntry
    var inline32 *_goml_vec_EnvEntry = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_new____T__EnvEntry()
    t141 = inline32
    var t142 Exp
    var inline30 string = "x"
    var inline31 Exp = Exp{
        _node: &Exp_node{
            _p0: inline30,
            _tag: 0,
        },
    }
    t142 = inline31
    var t143 Exp
    var inline29 Exp = Exp{
        _node: &Exp_node{
            _p1: t142,
            _p2: id__0,
            _tag: 1,
        },
    }
    t143 = inline29
    var t144 Exp
    var inline27 string = "z"
    var inline28 Exp = Exp{
        _node: &Exp_node{
            _p0: inline27,
            _tag: 0,
        },
    }
    t144 = inline28
    var t145 Exp
    var inline25 string = "z"
    var inline26 Exp = Exp{
        _node: &Exp_node{
            _p0: inline25,
            _p1: t143,
            _p2: t144,
            _tag: 3,
        },
    }
    t145 = inline26
    var t146 Exp
    var inline23 string = "y"
    var inline24 Exp = Exp{
        _node: &Exp_node{
            _p0: inline23,
            _tag: 0,
        },
    }
    t146 = inline24
    var t147 Exp
    var inline21 string = "y"
    var inline22 Exp = Exp{
        _node: &Exp_node{
            _p0: inline21,
            _p1: t145,
            _p2: t146,
            _tag: 3,
        },
    }
    t147 = inline22
    var t148 Exp
    var inline19 string = "x"
    var inline20 Exp = Exp{
        _node: &Exp_node{
            _p0: inline19,
            _p1: t147,
            _tag: 2,
        },
    }
    t148 = inline20
    var t149 Result__Typ__string = typeof(st__0, t141, t148)
    var inline9 string = "fun_x_let_y_let_z_x_id_z_y"
    switch t149._tag {
    case 0:
        var inline10 Typ = t149._p0
        var inline11 string = inline9 + ": "
        var inline12 string = typ_to_string(inline10)
        var inline13 string = inline11 + inline12
        println__T_string(inline13)
    case 1:
        var inline15 string = t149._p1
        var inline16 string = inline9 + ": "
        var inline17 string = inline16 + inline15
        println__T_string(inline17)
    default:
        panic("non-exhaustive match")
    }
    var inline6 string = ""
    var inline7 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline6)
    _goml_runtime_core_string_println(inline7)
    var inline3 string = "All Done"
    var inline4 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline3)
    _goml_runtime_core_string_println(inline4)
    var inline0 string = ""
    var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
    _goml_runtime_core_string_println(inline1)
    return struct{}{}
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_set____T__i32(self__0 *ref_int32_x, value__0 int32) struct{} {
    ref_set__Ref_5int32(self__0, value__0)
    return struct{}{}
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_get____T__i32(self__0 *ref_int32_x) int32 {
    var t0 int32 = ref_get__Ref_5int32(self__0)
    return t0
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__Tv(value__0 Tv) *ref_Tv_x {
    var t0 *ref_Tv_x = ref__Ref_2Tv(value__0)
    return t0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_new____T__EnvEntry() *_goml_vec_EnvEntry {
    var t0 *_goml_vec_EnvEntry = vec_new__Vec_8EnvEntry()
    return t0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pushed____T__SubstEntry(self__0 *_goml_vec_SubstEntry, elem__0 SubstEntry) *_goml_vec_SubstEntry {
    var t0 int
    var inline4 int = vec_len__Vec_10SubstEntry(self__0)
    t0 = inline4
    var t1 int = t0 + 1
    var result__0 *_goml_vec_SubstEntry
    var inline3 *_goml_vec_SubstEntry = vec_with_capacity__Vec_10SubstEntry(t1)
    result__0 = inline3
    var index__0 int = 0
    Loop_loop0:
    for {
        var t2 int
        var inline2 int = vec_len__Vec_10SubstEntry(self__0)
        t2 = inline2
        var t3 bool = index__0 < t2
        if t3 {
            var t4 SubstEntry = vec_get__Vec_10SubstEntry(self__0, index__0)
            vec_push__Vec_10SubstEntry(result__0, t4)
            var compound_old0 int = index__0
            var compound_value0 int = 1
            var t5 int = compound_old0 + compound_value0
            index__0 = t5
            continue
        } else {
            break Loop_loop0
        }
    }
    vec_push__Vec_10SubstEntry(result__0, elem__0)
    return result__0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_new____T__SubstEntry() *_goml_vec_SubstEntry {
    var t0 *_goml_vec_SubstEntry = vec_new__Vec_10SubstEntry()
    return t0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pushed____T__EnvEntry(self__0 *_goml_vec_EnvEntry, elem__0 EnvEntry) *_goml_vec_EnvEntry {
    var t0 int
    var inline4 int = vec_len__Vec_8EnvEntry(self__0)
    t0 = inline4
    var t1 int = t0 + 1
    var result__0 *_goml_vec_EnvEntry
    var inline3 *_goml_vec_EnvEntry = vec_with_capacity__Vec_8EnvEntry(t1)
    result__0 = inline3
    var index__0 int = 0
    Loop_loop0:
    for {
        var t2 int
        var inline2 int = vec_len__Vec_8EnvEntry(self__0)
        t2 = inline2
        var t3 bool = index__0 < t2
        if t3 {
            var t4 EnvEntry = vec_get__Vec_8EnvEntry(self__0, index__0)
            vec_push__Vec_8EnvEntry(result__0, t4)
            var compound_old0 int = index__0
            var compound_value0 int = 1
            var t5 int = compound_old0 + compound_value0
            index__0 = t5
            continue
        } else {
            break Loop_loop0
        }
    }
    vec_push__Vec_8EnvEntry(result__0, elem__0)
    return result__0
}

func println__T_string(value__0 string) struct{} {
    var t0 string
    t0 = value__0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func char_to_string(value__0 rune) string {
    var t0 uint32 = uint32(rune(value__0))
    var t1 bool
    var inline0 bool = t0 <= 1114111
    if inline0 {
        var inline1 bool = t0 >= 55296
        var inline2 bool
        if inline1 {
            var inline4 bool = t0 <= 57343
            inline2 = inline4
        } else {
            inline2 = false
        }
        var inline3 bool = !inline2
        t1 = inline3
    } else {
        t1 = false
    }
    if t1 {
        var t2 string = _goml_runtime_core_char_to_string(value__0)
        return t2
    } else {
        _goml_runtime_core_string_get("", -1)
        return ""
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

func main() {
    main0()
}
