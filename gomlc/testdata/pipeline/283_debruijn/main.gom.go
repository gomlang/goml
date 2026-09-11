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

type closure_env_lift_substitution_0 struct {
    sigma_0 func(int) Term
}

type closure_env_beta_substitution_1 struct {
    argument_0 Term
}

type closure_env_main_2 struct {}

type closure_env_main_3 struct {}

type Ordering uint8

type Named struct {
    _node *Named_node
}

type Named_node struct {
    _p0 string
    _p1 Named
    _p2 Named
    _tag uint8
}

func _goml_enum_tag_Named(value Named) uint8 {
    if value._node == nil {
        return 0
    }
    return value._node._tag
}

type Term struct {
    _node *Term_node
}

type Term_node struct {
    _p0 int
    _p1 Term
    _p2 Term
    _tag uint8
}

func _goml_enum_tag_Term(value Term) uint8 {
    if value._node == nil {
        return 0
    }
    return value._node._tag
}

type Context struct {
    _node *Context_node
}

type Context_node struct {
    _p0 string
    _p1 Context
    _tag uint8
}

func _goml_enum_tag_Context(value Context) uint8 {
    if value._node == nil {
        return 0
    }
    return value._node._tag
}

type Result__isize__string struct {
    _p1 string
    _p0 int
    _tag uint8
}

type Result__Term__string struct {
    _p1 string
    _p0 Term
    _tag uint8
}

func _goml_m_trait__impl_i_PartialEq_i_Term_i_eq(self__0 Term, other__0 Term) bool {
    switch _goml_enum_tag_Term(other__0) {
    case 0:
        var x0 int = other__0._node._p0
        switch _goml_enum_tag_Term(self__0) {
        case 0:
            var x1 int = self__0._node._p0
            var t0 bool = x1 == x0
            return t0
        default:
            return false
        }
    case 1:
        var x2 Term = other__0._node._p1
        switch _goml_enum_tag_Term(self__0) {
        case 1:
            var x3 Term = self__0._node._p1
            var t1 bool = _goml_m_trait__impl_i_PartialEq_i_Term_i_eq(x3, x2)
            return t1
        default:
            return false
        }
    case 2:
        var x4 Term = other__0._node._p1
        var x5 Term = other__0._node._p2
        switch _goml_enum_tag_Term(self__0) {
        case 2:
            var x6 Term = self__0._node._p1
            var x7 Term = self__0._node._p2
            var t2 bool = _goml_m_trait__impl_i_PartialEq_i_Term_i_eq(x6, x4)
            if t2 {
                var t3 bool = _goml_m_trait__impl_i_PartialEq_i_Term_i_eq(x7, x5)
                return t3
            } else {
                return false
            }
        default:
            return false
        }
    default:
        panic("non-exhaustive match")
    }
}

func lookup(name__0 string, context__0 Context) Result__isize__string {
    switch _goml_enum_tag_Context(context__0) {
    case 0:
        var t0 string = "unbound variable: " + name__0
        var t1 Result__isize__string = Result__isize__string{
            _p1: t0,
            _tag: 1,
        }
        return t1
    case 1:
        var x0 string = context__0._node._p0
        var x1 Context = context__0._node._p1
        var t2 bool = name__0 == x0
        if t2 {
            var t3 Result__isize__string = Result__isize__string{
                _p0: 0,
                _tag: 0,
            }
            return t3
        } else {
            var mtmp0 Result__isize__string = lookup(name__0, x1)
            var jp0 int
            switch mtmp0._tag {
            case 0:
                var x2 int = mtmp0._p0
                jp0 = x2
                var t4 int = jp0 + 1
                var t5 Result__isize__string = Result__isize__string{
                    _p0: t4,
                    _tag: 0,
                }
                return t5
            case 1:
                var x3 string = mtmp0._p1
                var t6 Result__isize__string = Result__isize__string{
                    _p1: x3,
                    _tag: 1,
                }
                return t6
            default:
                panic("non-exhaustive match")
            }
        }
    default:
        panic("non-exhaustive match")
    }
}

func translate(term__0 Named, context__0 Context) Result__Term__string {
    switch _goml_enum_tag_Named(term__0) {
    case 0:
        var x0 string = term__0._node._p0
        var mtmp0 Result__isize__string = lookup(x0, context__0)
        var jp0 int
        switch mtmp0._tag {
        case 0:
            var x1 int = mtmp0._p0
            jp0 = x1
            var t0 Term = Term{
                _node: &Term_node{
                    _p0: jp0,
                    _tag: 0,
                },
            }
            var t1 Result__Term__string = Result__Term__string{
                _p0: t0,
                _tag: 0,
            }
            return t1
        case 1:
            var x2 string = mtmp0._p1
            var t2 Result__Term__string = Result__Term__string{
                _p1: x2,
                _tag: 1,
            }
            return t2
        default:
            panic("non-exhaustive match")
        }
    case 1:
        var x3 string = term__0._node._p0
        var x4 Named = term__0._node._p1
        var t3 Context = Context{
            _node: &Context_node{
                _p0: x3,
                _p1: context__0,
                _tag: 1,
            },
        }
        var mtmp1 Result__Term__string = translate(x4, t3)
        var jp1 Term
        switch mtmp1._tag {
        case 0:
            var x5 Term = mtmp1._p0
            jp1 = x5
            var t4 Term = Term{
                _node: &Term_node{
                    _p1: jp1,
                    _tag: 1,
                },
            }
            var t5 Result__Term__string = Result__Term__string{
                _p0: t4,
                _tag: 0,
            }
            return t5
        case 1:
            var x6 string = mtmp1._p1
            var t6 Result__Term__string = Result__Term__string{
                _p1: x6,
                _tag: 1,
            }
            return t6
        default:
            panic("non-exhaustive match")
        }
    case 2:
        var x7 Named = term__0._node._p1
        var x8 Named = term__0._node._p2
        var mtmp2 Result__Term__string = translate(x7, context__0)
        var jp2 Term
        switch mtmp2._tag {
        case 0:
            var x11 Term = mtmp2._p0
            jp2 = x11
            var mtmp3 Result__Term__string = translate(x8, context__0)
            var jp3 Term
            switch mtmp3._tag {
            case 0:
                var x9 Term = mtmp3._p0
                jp3 = x9
                var t7 Term = Term{
                    _node: &Term_node{
                        _p1: jp2,
                        _p2: jp3,
                        _tag: 2,
                    },
                }
                var t8 Result__Term__string = Result__Term__string{
                    _p0: t7,
                    _tag: 0,
                }
                return t8
            case 1:
                var x10 string = mtmp3._p1
                var t9 Result__Term__string = Result__Term__string{
                    _p1: x10,
                    _tag: 1,
                }
                return t9
            default:
                panic("non-exhaustive match")
            }
        case 1:
            var x12 string = mtmp2._p1
            var t10 Result__Term__string = Result__Term__string{
                _p1: x12,
                _tag: 1,
            }
            return t10
        default:
            panic("non-exhaustive match")
        }
    default:
        panic("non-exhaustive match")
    }
}

func render(term__0 Term) string {
    switch _goml_enum_tag_Term(term__0) {
    case 0:
        var x0 int = term__0._node._p0
        var inline0 string = __goml_builtin_int_to_string(x0)
        return inline0
    case 1:
        var x1 Term = term__0._node._p1
        var t0 string = render(x1)
        var t1 string = "(lambda " + t0
        var t2 string = t1 + ")"
        return t2
    case 2:
        var x2 Term = term__0._node._p1
        var x3 Term = term__0._node._p2
        var t3 string = render(x2)
        var t4 string = "(" + t3
        var t5 string = t4 + " "
        var t6 string = render(x3)
        var t7 string = t5 + t6
        var t8 string = t7 + ")"
        return t8
    default:
        panic("non-exhaustive match")
    }
}

func well_scoped(term__0 Term, size__0 int) bool {
    switch _goml_enum_tag_Term(term__0) {
    case 0:
        var x0 int = term__0._node._p0
        var t0 bool = x0 >= 0
        if t0 {
            var t1 bool = x0 < size__0
            return t1
        } else {
            return false
        }
    case 1:
        var x1 Term = term__0._node._p1
        var t2 int = size__0 + 1
        var t3 bool = well_scoped(x1, t2)
        return t3
    case 2:
        var x2 Term = term__0._node._p1
        var x3 Term = term__0._node._p2
        var t4 bool = well_scoped(x2, size__0)
        if t4 {
            var t5 bool = well_scoped(x3, size__0)
            return t5
        } else {
            return false
        }
    default:
        panic("non-exhaustive match")
    }
}

func shift(term__0 Term, amount__0 int, cutoff__0 int) Term {
    switch _goml_enum_tag_Term(term__0) {
    case 0:
        var x0 int = term__0._node._p0
        var t0 bool = x0 < cutoff__0
        var jp0 int
        if t0 {
            jp0 = x0
        } else {
            var t2 int = x0 + amount__0
            jp0 = t2
        }
        var t1 Term = Term{
            _node: &Term_node{
                _p0: jp0,
                _tag: 0,
            },
        }
        return t1
    case 1:
        var x1 Term = term__0._node._p1
        var t3 int = cutoff__0 + 1
        var t4 Term = shift(x1, amount__0, t3)
        var t5 Term = Term{
            _node: &Term_node{
                _p1: t4,
                _tag: 1,
            },
        }
        return t5
    case 2:
        var x2 Term = term__0._node._p1
        var x3 Term = term__0._node._p2
        var t6 Term = shift(x2, amount__0, cutoff__0)
        var t7 Term = shift(x3, amount__0, cutoff__0)
        var t8 Term = Term{
            _node: &Term_node{
                _p1: t6,
                _p2: t7,
                _tag: 2,
            },
        }
        return t8
    default:
        panic("non-exhaustive match")
    }
}

func instantiate(body__0 Term, argument__0 Term, depth__0 int) Term {
    switch _goml_enum_tag_Term(body__0) {
    case 0:
        var x0 int = body__0._node._p0
        var t0 bool = x0 < depth__0
        if t0 {
            var t1 Term = Term{
                _node: &Term_node{
                    _p0: x0,
                    _tag: 0,
                },
            }
            return t1
        } else {
            var t2 bool = x0 == depth__0
            if t2 {
                var t3 Term = shift(argument__0, depth__0, 0)
                return t3
            } else {
                var t4 int = x0 - 1
                var t5 Term = Term{
                    _node: &Term_node{
                        _p0: t4,
                        _tag: 0,
                    },
                }
                return t5
            }
        }
    case 1:
        var x1 Term = body__0._node._p1
        var t6 int = depth__0 + 1
        var t7 Term = instantiate(x1, argument__0, t6)
        var t8 Term = Term{
            _node: &Term_node{
                _p1: t7,
                _tag: 1,
            },
        }
        return t8
    case 2:
        var x2 Term = body__0._node._p1
        var x3 Term = body__0._node._p2
        var t9 Term = instantiate(x2, argument__0, depth__0)
        var t10 Term = instantiate(x3, argument__0, depth__0)
        var t11 Term = Term{
            _node: &Term_node{
                _p1: t9,
                _p2: t10,
                _tag: 2,
            },
        }
        return t11
    default:
        panic("non-exhaustive match")
    }
}

func substitute(term__0 Term, sigma__0 func(int) Term) Term {
    switch _goml_enum_tag_Term(term__0) {
    case 0:
        var x0 int = term__0._node._p0
        var t0 Term = sigma__0(x0)
        return t0
    case 1:
        var x1 Term = term__0._node._p1
        var t1 func(int) Term
        var inline0 closure_env_lift_substitution_0 = closure_env_lift_substitution_0{
            sigma_0: sigma__0,
        }
        var inline1 func(int) Term = func(p0 int) Term {
            return _goml_m_inherent_i_closure__en_h4945d894b03216ac23e7f2c330830aeb_tion__0_i_apply(inline0, p0)
        }
        t1 = inline1
        var t2 Term = substitute(x1, t1)
        var t3 Term = Term{
            _node: &Term_node{
                _p1: t2,
                _tag: 1,
            },
        }
        return t3
    case 2:
        var x2 Term = term__0._node._p1
        var x3 Term = term__0._node._p2
        var t4 Term = substitute(x2, sigma__0)
        var t5 Term = substitute(x3, sigma__0)
        var t6 Term = Term{
            _node: &Term_node{
                _p1: t4,
                _p2: t5,
                _tag: 2,
            },
        }
        return t6
    default:
        panic("non-exhaustive match")
    }
}

func check_translation(label__0 string, term__0 Named, context__0 Context, expected__0 Term) struct{} {
    var mtmp0 Result__Term__string = translate(term__0, context__0)
    switch mtmp0._tag {
    case 0:
        var x0 Term = mtmp0._p0
        var t0 string = label__0 + ": "
        var t1 string = render(x0)
        var t2 string = t0 + t1
        var t3 string = t2 + " equal="
        var t4 bool = _goml_m_trait__impl_i_PartialEq_i_Term_i_eq(x0, expected__0)
        var t5 string
        var inline2 string = _goml_runtime_core_bool_to_string(t4)
        t5 = inline2
        var t6 string = t3 + t5
        var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t6)
        _goml_runtime_core_string_println(inline0)
        return struct{}{}
    case 1:
        var x1 string = mtmp0._p1
        var t7 string = label__0 + ": "
        var t8 string = t7 + x1
        var inline3 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t8)
        _goml_runtime_core_string_println(inline3)
        return struct{}{}
    default:
        panic("non-exhaustive match")
    }
}

func check_beta(label__0 string, body__0 Term, argument__0 Term, size__0 int, expected__0 Term) struct{} {
    var actual__0 Term = instantiate(body__0, argument__0, 0)
    var t0 func(int) Term
    var inline7 closure_env_beta_substitution_1 = closure_env_beta_substitution_1{
        argument_0: argument__0,
    }
    var inline8 func(int) Term = func(p0 int) Term {
        return _goml_m_inherent_i_closure__en_h04b89aaa6ab63af6c11902bf658d4e55_tion__1_i_apply(inline7, p0)
    }
    t0 = inline8
    var simultaneous__0 Term = substitute(body__0, t0)
    var t1 string = label__0 + ": "
    var t2 string = render(actual__0)
    var t3 string = t1 + t2
    var inline5 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t3)
    _goml_runtime_core_string_println(inline5)
    var t4 bool = _goml_m_trait__impl_i_PartialEq_i_Term_i_eq(actual__0, expected__0)
    var t5 string
    var inline4 string = _goml_runtime_core_bool_to_string(t4)
    t5 = inline4
    var t6 string = "expected=" + t5
    var t7 string = t6 + " simultaneous="
    var t8 bool = _goml_m_trait__impl_i_PartialEq_i_Term_i_eq(actual__0, simultaneous__0)
    var t9 string
    var inline3 string = _goml_runtime_core_bool_to_string(t8)
    t9 = inline3
    var t10 string = t7 + t9
    var t11 string = t10 + " scoped="
    var t12 int = size__0 + 1
    var t13 bool = well_scoped(body__0, t12)
    var jp0 bool
    if t13 {
        var t17 bool = well_scoped(argument__0, size__0)
        jp0 = t17
    } else {
        jp0 = false
    }
    var jp1 bool
    if jp0 {
        var t16 bool = well_scoped(actual__0, size__0)
        jp1 = t16
    } else {
        jp1 = false
    }
    var t14 string
    var inline2 string = _goml_runtime_core_bool_to_string(jp1)
    t14 = inline2
    var t15 string = t11 + t14
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t15)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func main0() struct{} {
    var empty__0 Context = Context{}
    var x__0 Named = Named{
        _node: &Named_node{
            _p0: "x",
            _tag: 0,
        },
    }
    var y__0 Named = Named{
        _node: &Named_node{
            _p0: "y",
            _tag: 0,
        },
    }
    var z__0 Named = Named{
        _node: &Named_node{
            _p0: "z",
            _tag: 0,
        },
    }
    var zero__0 Term = Term{
        _node: &Term_node{
            _p0: 0,
            _tag: 0,
        },
    }
    var one__0 Term = Term{
        _node: &Term_node{
            _p0: 1,
            _tag: 0,
        },
    }
    var two__0 Term = Term{
        _node: &Term_node{
            _p0: 2,
            _tag: 0,
        },
    }
    var identity__0 Term = Term{
        _node: &Term_node{
            _p1: zero__0,
            _tag: 1,
        },
    }
    var t0 Named = Named{
        _node: &Named_node{
            _p0: "x",
            _p1: x__0,
            _tag: 1,
        },
    }
    check_translation("identity x", t0, empty__0, identity__0)
    var t1 Named = Named{
        _node: &Named_node{
            _p0: "y",
            _p1: y__0,
            _tag: 1,
        },
    }
    check_translation("identity y", t1, empty__0, identity__0)
    var t2 Named = Named{
        _node: &Named_node{
            _p0: "y",
            _p1: x__0,
            _tag: 1,
        },
    }
    var t3 Named = Named{
        _node: &Named_node{
            _p0: "x",
            _p1: t2,
            _tag: 1,
        },
    }
    var t4 Term = Term{
        _node: &Term_node{
            _p1: one__0,
            _tag: 1,
        },
    }
    var t5 Term = Term{
        _node: &Term_node{
            _p1: t4,
            _tag: 1,
        },
    }
    check_translation("outer", t3, empty__0, t5)
    var t6 Named = Named{
        _node: &Named_node{
            _p0: "y",
            _p1: y__0,
            _tag: 1,
        },
    }
    var t7 Named = Named{
        _node: &Named_node{
            _p0: "x",
            _p1: t6,
            _tag: 1,
        },
    }
    var t8 Term = Term{
        _node: &Term_node{
            _p1: identity__0,
            _tag: 1,
        },
    }
    check_translation("inner", t7, empty__0, t8)
    var t9 Named = Named{
        _node: &Named_node{
            _p1: x__0,
            _p2: y__0,
            _tag: 2,
        },
    }
    var t10 Named = Named{
        _node: &Named_node{
            _p0: "y",
            _p1: t9,
            _tag: 1,
        },
    }
    var t11 Named = Named{
        _node: &Named_node{
            _p0: "x",
            _p1: t10,
            _tag: 1,
        },
    }
    var t12 Term = Term{
        _node: &Term_node{
            _p1: one__0,
            _p2: zero__0,
            _tag: 2,
        },
    }
    var t13 Term = Term{
        _node: &Term_node{
            _p1: t12,
            _tag: 1,
        },
    }
    var t14 Term = Term{
        _node: &Term_node{
            _p1: t13,
            _tag: 1,
        },
    }
    check_translation("application", t11, empty__0, t14)
    var t15 Named = Named{
        _node: &Named_node{
            _p0: "x",
            _p1: x__0,
            _tag: 1,
        },
    }
    var t16 Named = Named{
        _node: &Named_node{
            _p1: t15,
            _p2: x__0,
            _tag: 2,
        },
    }
    var t17 Named = Named{
        _node: &Named_node{
            _p0: "x",
            _p1: t16,
            _tag: 1,
        },
    }
    var t18 Term = Term{
        _node: &Term_node{
            _p1: identity__0,
            _p2: zero__0,
            _tag: 2,
        },
    }
    var t19 Term = Term{
        _node: &Term_node{
            _p1: t18,
            _tag: 1,
        },
    }
    check_translation("shadowing", t17, empty__0, t19)
    var t20 Named = Named{
        _node: &Named_node{
            _p1: x__0,
            _p2: y__0,
            _tag: 2,
        },
    }
    var t21 Named = Named{
        _node: &Named_node{
            _p0: "y",
            _p1: t20,
            _tag: 1,
        },
    }
    var t22 Named = Named{
        _node: &Named_node{
            _p1: x__0,
            _p2: t21,
            _tag: 2,
        },
    }
    var t23 Named = Named{
        _node: &Named_node{
            _p0: "x",
            _p1: t22,
            _tag: 1,
        },
    }
    var t24 Term = Term{
        _node: &Term_node{
            _p1: one__0,
            _p2: zero__0,
            _tag: 2,
        },
    }
    var t25 Term = Term{
        _node: &Term_node{
            _p1: t24,
            _tag: 1,
        },
    }
    var t26 Term = Term{
        _node: &Term_node{
            _p1: zero__0,
            _p2: t25,
            _tag: 2,
        },
    }
    var t27 Term = Term{
        _node: &Term_node{
            _p1: t26,
            _tag: 1,
        },
    }
    check_translation("relative address", t23, empty__0, t27)
    var t28 Named = Named{
        _node: &Named_node{
            _p1: x__0,
            _p2: z__0,
            _tag: 2,
        },
    }
    var t29 Named = Named{
        _node: &Named_node{
            _p1: y__0,
            _p2: z__0,
            _tag: 2,
        },
    }
    var t30 Named = Named{
        _node: &Named_node{
            _p1: t28,
            _p2: t29,
            _tag: 2,
        },
    }
    var t31 Named = Named{
        _node: &Named_node{
            _p0: "z",
            _p1: t30,
            _tag: 1,
        },
    }
    var t32 Named = Named{
        _node: &Named_node{
            _p0: "y",
            _p1: t31,
            _tag: 1,
        },
    }
    var t33 Named = Named{
        _node: &Named_node{
            _p0: "x",
            _p1: t32,
            _tag: 1,
        },
    }
    var t34 Term = Term{
        _node: &Term_node{
            _p1: two__0,
            _p2: zero__0,
            _tag: 2,
        },
    }
    var t35 Term = Term{
        _node: &Term_node{
            _p1: one__0,
            _p2: zero__0,
            _tag: 2,
        },
    }
    var t36 Term = Term{
        _node: &Term_node{
            _p1: t34,
            _p2: t35,
            _tag: 2,
        },
    }
    var t37 Term = Term{
        _node: &Term_node{
            _p1: t36,
            _tag: 1,
        },
    }
    var t38 Term = Term{
        _node: &Term_node{
            _p1: t37,
            _tag: 1,
        },
    }
    var t39 Term = Term{
        _node: &Term_node{
            _p1: t38,
            _tag: 1,
        },
    }
    check_translation("three binders", t33, empty__0, t39)
    var t40 Context = Context{
        _node: &Context_node{
            _p0: "z",
            _p1: empty__0,
            _tag: 1,
        },
    }
    var external__0 Context = Context{
        _node: &Context_node{
            _p0: "y",
            _p1: t40,
            _tag: 1,
        },
    }
    check_translation("free y", y__0, external__0, zero__0)
    check_translation("free z", z__0, external__0, one__0)
    var t41 Named = Named{
        _node: &Named_node{
            _p0: "x",
            _p1: y__0,
            _tag: 1,
        },
    }
    var t42 Term = Term{
        _node: &Term_node{
            _p1: one__0,
            _tag: 1,
        },
    }
    check_translation("open y", t41, external__0, t42)
    var t43 Named = Named{
        _node: &Named_node{
            _p0: "x",
            _p1: z__0,
            _tag: 1,
        },
    }
    var t44 Term = Term{
        _node: &Term_node{
            _p1: two__0,
            _tag: 1,
        },
    }
    check_translation("open z", t43, external__0, t44)
    var t45 Named = Named{
        _node: &Named_node{
            _p0: "y",
            _p1: y__0,
            _tag: 1,
        },
    }
    check_translation("external shadowing", t45, external__0, identity__0)
    var t46 Named = Named{
        _node: &Named_node{
            _p0: "x",
            _p1: x__0,
            _tag: 1,
        },
    }
    var t47 Named = Named{
        _node: &Named_node{
            _p1: t46,
            _p2: x__0,
            _tag: 2,
        },
    }
    check_translation("missing", t47, empty__0, zero__0)
    var t48 Term = Term{
        _node: &Term_node{
            _p1: one__0,
            _tag: 1,
        },
    }
    var t49 Term = Term{
        _node: &Term_node{
            _p1: t48,
            _tag: 1,
        },
    }
    var t50 Term = Term{
        _node: &Term_node{
            _p1: identity__0,
            _tag: 1,
        },
    }
    var t51 bool = _goml_m_trait__impl_i_PartialEq_i_Term_i_eq(t49, t50)
    var t52 bool = !t51
    var t53 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t52)
    var t54 string = "different binders=" + t53
    println__T_string(t54)
    var t55 Term = Term{
        _node: &Term_node{
            _p1: one__0,
            _p2: zero__0,
            _tag: 2,
        },
    }
    var open_function__0 Term = Term{
        _node: &Term_node{
            _p1: t55,
            _tag: 1,
        },
    }
    var shifted__0 Term = shift(open_function__0, 1, 0)
    var t56 string = render(shifted__0)
    var t57 string = "shift cutoff: " + t56
    println__T_string(t57)
    var t58 Term = Term{
        _node: &Term_node{
            _p1: two__0,
            _p2: zero__0,
            _tag: 2,
        },
    }
    var t59 Term = Term{
        _node: &Term_node{
            _p1: t58,
            _tag: 1,
        },
    }
    var t60 bool = _goml_m_trait__impl_i_PartialEq_i_Term_i_eq(shifted__0, t59)
    var t61 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t60)
    var t62 string = "shift expected=" + t61
    var inline27 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t62)
    _goml_runtime_core_string_println(inline27)
    var t63 Term = shift(open_function__0, 0, 0)
    var t64 bool = _goml_m_trait__impl_i_PartialEq_i_Term_i_eq(t63, open_function__0)
    var t65 string
    var inline26 string = _goml_runtime_core_bool_to_string(t64)
    t65 = inline26
    var t66 string = "shift zero=" + t65
    var inline24 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t66)
    _goml_runtime_core_string_println(inline24)
    var t67 Term = shift(identity__0, 3, 0)
    var t68 bool = _goml_m_trait__impl_i_PartialEq_i_Term_i_eq(t67, identity__0)
    var t69 string
    var inline23 string = _goml_runtime_core_bool_to_string(t68)
    t69 = inline23
    var t70 string = "shift closed=" + t69
    var inline21 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t70)
    _goml_runtime_core_string_println(inline21)
    var t71 closure_env_main_2 = closure_env_main_2{}
    var t72 func(int) Term = func(p0 int) Term {
        return _goml_m_inherent_i_closure__env__main__2_i_closure__env__main__2_i_apply(t71, p0)
    }
    var t73 Term = substitute(open_function__0, t72)
    var t74 bool = _goml_m_trait__impl_i_PartialEq_i_Term_i_eq(t73, open_function__0)
    var t75 string
    var inline20 string = _goml_runtime_core_bool_to_string(t74)
    t75 = inline20
    var t76 string = "substitution identity=" + t75
    var inline18 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t76)
    _goml_runtime_core_string_println(inline18)
    var t77 closure_env_main_3 = closure_env_main_3{}
    var t78 func(int) Term = func(p0 int) Term {
        return _goml_m_inherent_i_closure__env__main__3_i_closure__env__main__3_i_apply(t77, p0)
    }
    var t79 Term = substitute(open_function__0, t78)
    var t80 bool = _goml_m_trait__impl_i_PartialEq_i_Term_i_eq(t79, shifted__0)
    var t81 string
    var inline17 string = _goml_runtime_core_bool_to_string(t80)
    t81 = inline17
    var t82 string = "substitution shift=" + t81
    var inline15 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t82)
    _goml_runtime_core_string_println(inline15)
    check_beta("identity beta", zero__0, identity__0, 0, identity__0)
    var t83 Term = Term{
        _node: &Term_node{
            _p1: one__0,
            _tag: 1,
        },
    }
    var t84 Term = Term{
        _node: &Term_node{
            _p1: one__0,
            _tag: 1,
        },
    }
    check_beta("avoid capture", t83, zero__0, 1, t84)
    var t85 Term = Term{
        _node: &Term_node{
            _p1: two__0,
            _p2: one__0,
            _tag: 2,
        },
    }
    var t86 Term = Term{
        _node: &Term_node{
            _p1: t85,
            _p2: zero__0,
            _tag: 2,
        },
    }
    var t87 Term = Term{
        _node: &Term_node{
            _p1: t86,
            _tag: 1,
        },
    }
    var t88 Term = Term{
        _node: &Term_node{
            _p1: one__0,
            _p2: one__0,
            _tag: 2,
        },
    }
    var t89 Term = Term{
        _node: &Term_node{
            _p1: t88,
            _p2: zero__0,
            _tag: 2,
        },
    }
    var t90 Term = Term{
        _node: &Term_node{
            _p1: t89,
            _tag: 1,
        },
    }
    check_beta("all three cases", t87, zero__0, 1, t90)
    var t91 Term = Term{
        _node: &Term_node{
            _p1: one__0,
            _tag: 1,
        },
    }
    var t92 Term = Term{
        _node: &Term_node{
            _p1: shifted__0,
            _tag: 1,
        },
    }
    check_beta("argument local binder", t91, open_function__0, 1, t92)
    var t93 Term = Term{
        _node: &Term_node{
            _p1: two__0,
            _p2: zero__0,
            _tag: 2,
        },
    }
    var t94 Term = Term{
        _node: &Term_node{
            _p1: t93,
            _tag: 1,
        },
    }
    var t95 Term = Term{
        _node: &Term_node{
            _p1: t94,
            _tag: 1,
        },
    }
    var t96 Term = Term{
        _node: &Term_node{
            _p1: two__0,
            _p2: zero__0,
            _tag: 2,
        },
    }
    var t97 Term = Term{
        _node: &Term_node{
            _p1: t96,
            _tag: 1,
        },
    }
    var t98 Term = Term{
        _node: &Term_node{
            _p1: t97,
            _tag: 1,
        },
    }
    check_beta("two nested binders", t95, zero__0, 1, t98)
    var t99 Term = Term{
        _node: &Term_node{
            _p1: zero__0,
            _tag: 1,
        },
    }
    check_beta("unused argument", t99, open_function__0, 1, identity__0)
    check_beta("remove external slot", one__0, identity__0, 1, zero__0)
    var t100 Term = Term{
        _node: &Term_node{
            _p1: zero__0,
            _p2: zero__0,
            _tag: 2,
        },
    }
    var t101 Term = Term{
        _node: &Term_node{
            _p1: identity__0,
            _p2: identity__0,
            _tag: 2,
        },
    }
    check_beta("duplicate argument", t100, identity__0, 0, t101)
    var t102 bool = well_scoped(identity__0, 0)
    var t103 string
    var inline14 string = _goml_runtime_core_bool_to_string(t102)
    t103 = inline14
    var t104 string = "closed scope=" + t103
    var inline12 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t104)
    _goml_runtime_core_string_println(inline12)
    var t105 bool = well_scoped(open_function__0, 1)
    var t106 string
    var inline11 string = _goml_runtime_core_bool_to_string(t105)
    t106 = inline11
    var t107 string = "open scope=" + t106
    var inline9 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t107)
    _goml_runtime_core_string_println(inline9)
    var t108 bool = well_scoped(open_function__0, 0)
    var t109 string
    var inline8 string = _goml_runtime_core_bool_to_string(t108)
    t109 = inline8
    var t110 string = "missing external scope=" + t109
    var inline6 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t110)
    _goml_runtime_core_string_println(inline6)
    var t111 Term = Term{
        _node: &Term_node{
            _p0: -1,
            _tag: 0,
        },
    }
    var t112 bool = well_scoped(t111, 1)
    var t113 string
    var inline5 string = _goml_runtime_core_bool_to_string(t112)
    t113 = inline5
    var t114 string = "negative index=" + t113
    var inline3 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t114)
    _goml_runtime_core_string_println(inline3)
    var t115 bool = well_scoped(one__0, 1)
    var t116 string
    var inline2 string = _goml_runtime_core_bool_to_string(t115)
    t116 = inline2
    var t117 string = "index at boundary=" + t116
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t117)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func println__T_string(value__0 string) struct{} {
    var t0 string
    t0 = value__0
    _goml_runtime_core_string_println(t0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_bool_i_to__string(self__0 bool) string {
    var t0 string = _goml_runtime_core_bool_to_string(self__0)
    return t0
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

func _goml_m_inherent_i_closure__en_h4945d894b03216ac23e7f2c330830aeb_tion__0_i_apply(env0 closure_env_lift_substitution_0, index__0 int) Term {
    var sigma__0 func(int) Term = env0.sigma_0
    var t0 bool = index__0 == 0
    if t0 {
        var t1 Term = Term{
            _node: &Term_node{
                _p0: 0,
                _tag: 0,
            },
        }
        return t1
    } else {
        var t2 int = index__0 - 1
        var t3 Term = sigma__0(t2)
        var t4 Term = shift(t3, 1, 0)
        return t4
    }
}

func _goml_m_inherent_i_closure__en_h04b89aaa6ab63af6c11902bf658d4e55_tion__1_i_apply(env0 closure_env_beta_substitution_1, index__0 int) Term {
    var argument__0 Term = env0.argument_0
    var t0 bool = index__0 == 0
    if t0 {
        return argument__0
    } else {
        var t1 int = index__0 - 1
        var t2 Term = Term{
            _node: &Term_node{
                _p0: t1,
                _tag: 0,
            },
        }
        return t2
    }
}

func _goml_m_inherent_i_closure__env__main__2_i_closure__env__main__2_i_apply(env0 closure_env_main_2, index__0 int) Term {
    var t0 Term = Term{
        _node: &Term_node{
            _p0: index__0,
            _tag: 0,
        },
    }
    return t0
}

func _goml_m_inherent_i_closure__env__main__3_i_closure__env__main__3_i_apply(env0 closure_env_main_3, index__0 int) Term {
    var t0 int = index__0 + 1
    var t1 Term = Term{
        _node: &Term_node{
            _p0: t0,
            _tag: 0,
        },
    }
    return t1
}

func main() {
    main0()
}
