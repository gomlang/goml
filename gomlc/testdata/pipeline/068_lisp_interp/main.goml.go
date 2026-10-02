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

func _goml_runtime_core_string_get(s string, i int) rune {
    return rune(s[i])
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

func _goml_runtime_core_char_to_string(x rune) string {
    return string(x)
}

func _goml_runtime_core_char_from_uint32(value uint32) Tuple2_4bool_4char {
    return Tuple2_4bool_4char{
        _0: true,
        _1: rune(value),
    }
}

func _goml_runtime_core_string_println(s string) struct{} {
    _goml_os.Stdout.WriteString(s + "\n")
    return struct{}{}
}

type _goml_vec_Token struct {
    items []Token
}

func vec_new__Vec_5Token() *_goml_vec_Token {
    return &_goml_vec_Token{
        items: nil,
    }
}

func vec_with_capacity__Vec_5Token(capacity int) *_goml_vec_Token {
    return &_goml_vec_Token{
        items: make([]Token, 0, capacity),
    }
}

func vec_push__Vec_5Token(vec *_goml_vec_Token, elem Token) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_get__Vec_5Token(vec *_goml_vec_Token, index int) Token {
    return vec.items[index]
}

func vec_len__Vec_5Token(vec *_goml_vec_Token) int {
    return int(len(vec.items))
}

type _goml_vec_Binding struct {
    items []Binding
}

func vec_new__Vec_7Binding() *_goml_vec_Binding {
    return &_goml_vec_Binding{
        items: nil,
    }
}

func vec_with_capacity__Vec_7Binding(capacity int) *_goml_vec_Binding {
    return &_goml_vec_Binding{
        items: make([]Binding, 0, capacity),
    }
}

func vec_push__Vec_7Binding(vec *_goml_vec_Binding, elem Binding) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_get__Vec_7Binding(vec *_goml_vec_Binding, index int) Binding {
    return vec.items[index]
}

func vec_len__Vec_7Binding(vec *_goml_vec_Binding) int {
    return int(len(vec.items))
}

type _goml_vec_SExpr struct {
    items []SExpr
}

func vec_new__Vec_5SExpr() *_goml_vec_SExpr {
    return &_goml_vec_SExpr{
        items: nil,
    }
}

func vec_with_capacity__Vec_5SExpr(capacity int) *_goml_vec_SExpr {
    return &_goml_vec_SExpr{
        items: make([]SExpr, 0, capacity),
    }
}

func vec_push__Vec_5SExpr(vec *_goml_vec_SExpr, elem SExpr) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_get__Vec_5SExpr(vec *_goml_vec_SExpr, index int) SExpr {
    return vec.items[index]
}

func vec_len__Vec_5SExpr(vec *_goml_vec_SExpr) int {
    return int(len(vec.items))
}

type _goml_vec_Value struct {
    items []Value
}

func vec_new__Vec_5Value() *_goml_vec_Value {
    return &_goml_vec_Value{
        items: nil,
    }
}

func vec_with_capacity__Vec_5Value(capacity int) *_goml_vec_Value {
    return &_goml_vec_Value{
        items: make([]Value, 0, capacity),
    }
}

func vec_push__Vec_5Value(vec *_goml_vec_Value, elem Value) struct{} {
    vec.items = append(vec.items, elem)
    return struct{}{}
}

func vec_get__Vec_5Value(vec *_goml_vec_Value, index int) Value {
    return vec.items[index]
}

func vec_len__Vec_5Value(vec *_goml_vec_Value) int {
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

func vec_len__Vec_6string(vec *_goml_vec_string) int {
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

type ref_string_x struct {
    value string
}

func ref__Ref_6string(value string) *ref_string_x {
    return &ref_string_x{
        value: value,
    }
}

func ref_get__Ref_6string(reference *ref_string_x) string {
    return reference.value
}

func ref_set__Ref_6string(reference *ref_string_x, value string) struct{} {
    reference.value = value
    return struct{}{}
}

type ref_Vec_5Token_x struct {
    value *_goml_vec_Token
}

func ref__Ref_10Vec_5Token(value *_goml_vec_Token) *ref_Vec_5Token_x {
    return &ref_Vec_5Token_x{
        value: value,
    }
}

func ref_get__Ref_10Vec_5Token(reference *ref_Vec_5Token_x) *_goml_vec_Token {
    return reference.value
}

func ref_set__Ref_10Vec_5Token(reference *ref_Vec_5Token_x, value *_goml_vec_Token) struct{} {
    reference.value = value
    return struct{}{}
}

type ref_Value_x struct {
    value Value
}

func ref__Ref_5Value(value Value) *ref_Value_x {
    return &ref_Value_x{
        value: value,
    }
}

func ref_get__Ref_5Value(reference *ref_Value_x) Value {
    return reference.value
}

func ref_set__Ref_5Value(reference *ref_Value_x, value Value) struct{} {
    reference.value = value
    return struct{}{}
}

type ref_Vec_5SExpr_x struct {
    value *_goml_vec_SExpr
}

func ref__Ref_10Vec_5SExpr(value *_goml_vec_SExpr) *ref_Vec_5SExpr_x {
    return &ref_Vec_5SExpr_x{
        value: value,
    }
}

func ref_get__Ref_10Vec_5SExpr(reference *ref_Vec_5SExpr_x) *_goml_vec_SExpr {
    return reference.value
}

func ref_set__Ref_10Vec_5SExpr(reference *ref_Vec_5SExpr_x, value *_goml_vec_SExpr) struct{} {
    reference.value = value
    return struct{}{}
}

type ref_Vec_7Binding_x struct {
    value *_goml_vec_Binding
}

func ref__Ref_12Vec_7Binding(value *_goml_vec_Binding) *ref_Vec_7Binding_x {
    return &ref_Vec_7Binding_x{
        value: value,
    }
}

func ref_get__Ref_12Vec_7Binding(reference *ref_Vec_7Binding_x) *_goml_vec_Binding {
    return reference.value
}

func ref_set__Ref_12Vec_7Binding(reference *ref_Vec_7Binding_x, value *_goml_vec_Binding) struct{} {
    reference.value = value
    return struct{}{}
}

type ref_Vec_6string_x struct {
    value *_goml_vec_string
}

func ref__Ref_11Vec_6string(value *_goml_vec_string) *ref_Vec_6string_x {
    return &ref_Vec_6string_x{
        value: value,
    }
}

func ref_get__Ref_11Vec_6string(reference *ref_Vec_6string_x) *_goml_vec_string {
    return reference.value
}

func ref_set__Ref_11Vec_6string(reference *ref_Vec_6string_x, value *_goml_vec_string) struct{} {
    reference.value = value
    return struct{}{}
}

type ref_Vec_5Value_x struct {
    value *_goml_vec_Value
}

func ref__Ref_10Vec_5Value(value *_goml_vec_Value) *ref_Vec_5Value_x {
    return &ref_Vec_5Value_x{
        value: value,
    }
}

func ref_get__Ref_10Vec_5Value(reference *ref_Vec_5Value_x) *_goml_vec_Value {
    return reference.value
}

func ref_set__Ref_10Vec_5Value(reference *ref_Vec_5Value_x, value *_goml_vec_Value) struct{} {
    reference.value = value
    return struct{}{}
}

type Tuple2_5Token_3int struct {
    _0 Token
    _1 int
}

type Tuple2_10Vec_5SExpr_3int struct {
    _0 *_goml_vec_SExpr
    _1 int
}

type Tuple2_5SExpr_3int struct {
    _0 SExpr
    _1 int
}

type Tuple3_4bool_4char_3int struct {
    _0 bool
    _1 rune
    _2 int
}

type Tuple2_4bool_4char struct {
    _0 bool
    _1 rune
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

type Binding struct {
    name string
    value Value
}

type Lambda struct {
    params *_goml_vec_string
    body SExpr
    env *_goml_vec_Binding
    global *ref_Vec_7Binding_x
}

type Ordering uint8

type Token struct {
    _p0 string
    _p1 int32
    _p2 bool
    _tag uint8
}

type Value interface {
    isValue()
}

type Value_Int struct {
    _0 int32
}

func (_ Value_Int) isValue() {}

type Value_Bool struct {
    _0 bool
}

func (_ Value_Bool) isValue() {}

type Func struct {
    _0 Lambda
}

func (_ Func) isValue() {}

type Nil struct {}

func (_ Nil) isValue() {}

type SExpr interface {
    isSExpr()
}

type SExpr_Int struct {
    _0 int32
}

func (_ SExpr_Int) isSExpr() {}

type SExpr_Bool struct {
    _0 bool
}

func (_ SExpr_Bool) isSExpr() {}

type SExpr_Sym struct {
    _0 string
}

func (_ SExpr_Sym) isSExpr() {}

type List struct {
    _0 *_goml_vec_SExpr
}

func (_ List) isSExpr() {}

type Option__char uint64

func is_int_text(text__0 string) bool {
    var len__0 int
    var inline29 int = _goml_runtime_core_string_len(text__0)
    len__0 = inline29
    var t0 bool = len__0 == 0
    if t0 {
        return false
    } else {
        var i__0 *ref_int_x
        var inline27 int = 0
        var inline28 *ref_int_x = ref__Ref_3int(inline27)
        i__0 = inline28
        var saw_digit__0 *ref_bool_x
        var inline25 bool = false
        var inline26 *ref_bool_x = ref__Ref_4bool(inline25)
        saw_digit__0 = inline26
        var ok__0 *ref_bool_x
        var inline23 bool = true
        var inline24 *ref_bool_x = ref__Ref_4bool(inline23)
        ok__0 = inline24
        var started__0 *ref_bool_x
        var inline21 bool = false
        var inline22 *ref_bool_x = ref__Ref_4bool(inline21)
        started__0 = inline22
        Loop_loop0:
        for {
            var t2 bool
            var inline20 bool = ref_get__Ref_4bool(ok__0)
            t2 = inline20
            var jp0 bool
            if t2 {
                var t12 int
                var inline19 int = ref_get__Ref_3int(i__0)
                t12 = inline19
                var t13 bool = t12 < len__0
                jp0 = t13
            } else {
                jp0 = false
            }
            if jp0 {
                var t3 int
                var inline18 int = ref_get__Ref_3int(i__0)
                t3 = inline18
                var ch__0 rune
                var inline17 rune = string_get(text__0, t3)
                ch__0 = inline17
                var t4 bool
                var inline16 bool = ref_get__Ref_4bool(started__0)
                t4 = inline16
                var t5 bool = !t4
                var jp1 bool
                if t5 {
                    var t11 bool = ch__0 == 45
                    jp1 = t11
                } else {
                    jp1 = false
                }
                if jp1 {
                    var inline4 bool = true
                    ref_set__Ref_4bool(started__0, inline4)
                    var t6 int
                    var inline3 int = ref_get__Ref_3int(i__0)
                    t6 = inline3
                    var t7 int = t6 + 1
                    ref_set__Ref_3int(i__0, t7)
                    continue
                } else {
                    var t8 bool
                    var inline14 bool = ch__0 >= 48
                    if inline14 {
                        var inline15 bool = ch__0 <= 57
                        t8 = inline15
                    } else {
                        t8 = false
                    }
                    if t8 {
                        var inline10 bool = true
                        ref_set__Ref_4bool(started__0, inline10)
                        var inline8 bool = true
                        ref_set__Ref_4bool(saw_digit__0, inline8)
                        var t9 int
                        var inline7 int = ref_get__Ref_3int(i__0)
                        t9 = inline7
                        var t10 int = t9 + 1
                        ref_set__Ref_3int(i__0, t10)
                        continue
                    } else {
                        var inline12 bool = false
                        ref_set__Ref_4bool(ok__0, inline12)
                        continue
                    }
                }
            } else {
                break Loop_loop0
            }
        }
        var t1 bool
        var inline1 bool = ref_get__Ref_4bool(ok__0)
        t1 = inline1
        if t1 {
            var inline0 bool = ref_get__Ref_4bool(saw_digit__0)
            return inline0
        } else {
            return false
        }
    }
}

func parse_int32(text__0 string) int32 {
    var len__0 int = _goml_m_inherent_i_string_i_string_i_len(text__0)
    var i__0 *ref_int_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__isize(0)
    var negative__0 *ref_bool_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__bool(false)
    var started__0 *ref_bool_x
    var inline21 bool = false
    var inline22 *ref_bool_x = ref__Ref_4bool(inline21)
    started__0 = inline22
    var acc__0 *ref_int32_x
    var inline19 int32 = 0
    var inline20 *ref_int32_x = ref__Ref_5int32(inline19)
    acc__0 = inline20
    Loop_loop0:
    for {
        var t3 int
        var inline18 int = ref_get__Ref_3int(i__0)
        t3 = inline18
        var t4 bool = t3 < len__0
        if t4 {
            var t5 int
            var inline17 int = ref_get__Ref_3int(i__0)
            t5 = inline17
            var ch__0 rune
            var inline16 rune = string_get(text__0, t5)
            ch__0 = inline16
            var t6 bool
            var inline15 bool = ref_get__Ref_4bool(started__0)
            t6 = inline15
            var t7 bool = !t6
            var jp0 bool
            if t7 {
                var t15 bool = ch__0 == 45
                jp0 = t15
            } else {
                jp0 = false
            }
            if jp0 {
                var inline7 bool = true
                ref_set__Ref_4bool(started__0, inline7)
                var inline5 bool = true
                ref_set__Ref_4bool(negative__0, inline5)
                var t8 int
                var inline4 int = ref_get__Ref_3int(i__0)
                t8 = inline4
                var t9 int = t8 + 1
                ref_set__Ref_3int(i__0, t9)
                continue
            } else {
                var inline13 bool = true
                ref_set__Ref_4bool(started__0, inline13)
                var d__0 int32
                switch ch__0 {
                case 48:
                    d__0 = 0
                case 49:
                    d__0 = 1
                case 50:
                    d__0 = 2
                case 51:
                    d__0 = 3
                case 52:
                    d__0 = 4
                case 53:
                    d__0 = 5
                case 54:
                    d__0 = 6
                case 55:
                    d__0 = 7
                case 56:
                    d__0 = 8
                case 57:
                    d__0 = 9
                default:
                    d__0 = 0
                }
                var t10 int32
                var inline12 int32 = ref_get__Ref_5int32(acc__0)
                t10 = inline12
                var t11 int32 = t10 * 10
                var t12 int32 = t11 + d__0
                ref_set__Ref_5int32(acc__0, t12)
                var t13 int
                var inline10 int = ref_get__Ref_3int(i__0)
                t13 = inline10
                var t14 int = t13 + 1
                ref_set__Ref_3int(i__0, t14)
                continue
            }
        } else {
            break Loop_loop0
        }
    }
    var t0 bool
    var inline2 bool = ref_get__Ref_4bool(negative__0)
    t0 = inline2
    if t0 {
        var t1 int32
        var inline0 int32 = ref_get__Ref_5int32(acc__0)
        t1 = inline0
        var t2 int32 = 0 - t1
        return t2
    } else {
        var inline1 int32 = ref_get__Ref_5int32(acc__0)
        return inline1
    }
}

func lex_atom(source__0 string, start__0 int) Tuple2_5Token_3int {
    var len__0 int
    var inline22 int = _goml_runtime_core_string_len(source__0)
    len__0 = inline22
    var text__0 *ref_string_x
    var inline20 string = ""
    var inline21 *ref_string_x = ref__Ref_6string(inline20)
    text__0 = inline21
    var i__0 *ref_int_x
    var inline19 *ref_int_x = ref__Ref_3int(start__0)
    i__0 = inline19
    var done__0 *ref_bool_x
    var inline17 bool = false
    var inline18 *ref_bool_x = ref__Ref_4bool(inline17)
    done__0 = inline18
    Loop_loop0:
    for {
        var t8 bool
        var inline16 bool = ref_get__Ref_4bool(done__0)
        t8 = inline16
        var t9 bool = !t8
        var jp1 bool
        if t9 {
            var t17 int
            var inline15 int = ref_get__Ref_3int(i__0)
            t17 = inline15
            var t18 bool = t17 < len__0
            jp1 = t18
        } else {
            jp1 = false
        }
        if jp1 {
            var t10 int
            var inline14 int = ref_get__Ref_3int(i__0)
            t10 = inline14
            var ch__0 rune
            var inline13 rune = string_get(source__0, t10)
            ch__0 = inline13
            var t11 bool
            var inline9 bool = ch__0 == 40
            var inline10 bool
            if inline9 {
                inline10 = true
            } else {
                var inline12 bool = ch__0 == 41
                inline10 = inline12
            }
            if inline10 {
                t11 = true
                if t11 {
                    var inline2 bool = true
                    ref_set__Ref_4bool(done__0, inline2)
                    continue
                } else {
                    var t12 string
                    var inline8 string = ref_get__Ref_6string(text__0)
                    t12 = inline8
                    var t13 string
                    var inline7 string = char_to_string(ch__0)
                    t13 = inline7
                    var t14 string = t12 + t13
                    ref_set__Ref_6string(text__0, t14)
                    var t15 int
                    var inline5 int = ref_get__Ref_3int(i__0)
                    t15 = inline5
                    var t16 int = t15 + 1
                    ref_set__Ref_3int(i__0, t16)
                    continue
                }
            } else {
                var inline11 bool = ch__0 == 32
                t11 = inline11
                if t11 {
                    var inline2 bool = true
                    ref_set__Ref_4bool(done__0, inline2)
                    continue
                } else {
                    var t12 string
                    var inline8 string = ref_get__Ref_6string(text__0)
                    t12 = inline8
                    var t13 string
                    var inline7 string = char_to_string(ch__0)
                    t13 = inline7
                    var t14 string = t12 + t13
                    ref_set__Ref_6string(text__0, t14)
                    var t15 int
                    var inline5 int = ref_get__Ref_3int(i__0)
                    t15 = inline5
                    var t16 int = t15 + 1
                    ref_set__Ref_3int(i__0, t16)
                    continue
                }
            }
        } else {
            break Loop_loop0
        }
    }
    var atom__0 string
    var inline1 string = ref_get__Ref_6string(text__0)
    atom__0 = inline1
    var jp0 Token
    switch atom__0 {
    case "true":
        var t2 Token = Token{
            _p2: true,
            _tag: 4,
        }
        jp0 = t2
    case "false":
        var t3 Token = Token{
            _p2: false,
            _tag: 4,
        }
        jp0 = t3
    default:
        var t4 bool = is_int_text(atom__0)
        if t4 {
            var t5 int32 = parse_int32(atom__0)
            var t6 Token = Token{
                _p1: t5,
                _tag: 3,
            }
            jp0 = t6
        } else {
            var t7 Token = Token{
                _p0: atom__0,
                _tag: 2,
            }
            jp0 = t7
        }
    }
    var t0 int
    var inline0 int = ref_get__Ref_3int(i__0)
    t0 = inline0
    var t1 Tuple2_5Token_3int = Tuple2_5Token_3int{
        _0: jp0,
        _1: t0,
    }
    return t1
}

func lex(source__0 string) *_goml_vec_Token {
    var len__0 int
    var inline22 int = _goml_runtime_core_string_len(source__0)
    len__0 = inline22
    var toks0__0 *_goml_vec_Token
    var inline21 *_goml_vec_Token = vec_new__Vec_5Token()
    toks0__0 = inline21
    var toks__0 *ref_Vec_5Token_x
    var inline20 *ref_Vec_5Token_x = ref__Ref_10Vec_5Token(toks0__0)
    toks__0 = inline20
    var i__0 *ref_int_x
    var inline18 int = 0
    var inline19 *ref_int_x = ref__Ref_3int(inline18)
    i__0 = inline19
    Loop_loop0:
    for {
        var t0 int
        var inline17 int = ref_get__Ref_3int(i__0)
        t0 = inline17
        var t1 bool = t0 < len__0
        if t1 {
            var t2 int
            var inline16 int = ref_get__Ref_3int(i__0)
            t2 = inline16
            var ch__0 rune
            var inline15 rune = string_get(source__0, t2)
            ch__0 = inline15
            var t3 bool = ch__0 == 40
            if t3 {
                var t4 *_goml_vec_Token
                var inline4 *_goml_vec_Token = ref_get__Ref_10Vec_5Token(toks__0)
                t4 = inline4
                var t5 *_goml_vec_Token = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pushed____T__Token(t4, Token{
                    _tag: 0,
                })
                ref_set__Ref_10Vec_5Token(toks__0, t5)
                var t6 int
                var inline2 int = ref_get__Ref_3int(i__0)
                t6 = inline2
                var t7 int = t6 + 1
                ref_set__Ref_3int(i__0, t7)
                continue
            } else {
                var t8 bool = ch__0 == 41
                if t8 {
                    var t9 *_goml_vec_Token
                    var inline8 *_goml_vec_Token = ref_get__Ref_10Vec_5Token(toks__0)
                    t9 = inline8
                    var t10 *_goml_vec_Token = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pushed____T__Token(t9, Token{
                        _tag: 1,
                    })
                    ref_set__Ref_10Vec_5Token(toks__0, t10)
                    var t11 int
                    var inline6 int = ref_get__Ref_3int(i__0)
                    t11 = inline6
                    var t12 int = t11 + 1
                    ref_set__Ref_3int(i__0, t12)
                    continue
                } else {
                    var t13 bool = ch__0 == 32
                    if t13 {
                        var t14 int
                        var inline10 int = ref_get__Ref_3int(i__0)
                        t14 = inline10
                        var t15 int = t14 + 1
                        ref_set__Ref_3int(i__0, t15)
                        continue
                    } else {
                        var t16 int
                        var inline14 int = ref_get__Ref_3int(i__0)
                        t16 = inline14
                        var mtmp2 Tuple2_5Token_3int = lex_atom(source__0, t16)
                        var x0 Token = mtmp2._0
                        var x1 int = mtmp2._1
                        var t17 *_goml_vec_Token
                        var inline13 *_goml_vec_Token = ref_get__Ref_10Vec_5Token(toks__0)
                        t17 = inline13
                        var t18 *_goml_vec_Token = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pushed____T__Token(t17, x0)
                        ref_set__Ref_10Vec_5Token(toks__0, t18)
                        ref_set__Ref_3int(i__0, x1)
                        continue
                    }
                }
            }
        } else {
            break Loop_loop0
        }
    }
    var inline0 *_goml_vec_Token = ref_get__Ref_10Vec_5Token(toks__0)
    return inline0
}

func env_lookup(env__0 *_goml_vec_Binding, name__0 string) Value {
    var t0 int
    var inline13 int = vec_len__Vec_7Binding(env__0)
    t0 = inline13
    var t1 int = t0 - 1
    var i__0 *ref_int_x
    var inline12 *ref_int_x = ref__Ref_3int(t1)
    i__0 = inline12
    var result__0 *ref_Value_x
    var inline11 *ref_Value_x = ref__Ref_5Value(Nil{})
    result__0 = inline11
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
            var t10 int
            var inline7 int = ref_get__Ref_3int(i__0)
            t10 = inline7
            var t11 bool = t10 >= 0
            jp0 = t11
        } else {
            jp0 = false
        }
        if jp0 {
            var t4 int
            var inline6 int = ref_get__Ref_3int(i__0)
            t4 = inline6
            var binding__0 Binding = vec_get__Vec_7Binding(env__0, t4)
            var t5 string = binding__0.name
            var t6 bool = t5 == name__0
            if t6 {
                var t7 Value = binding__0.value
                ref_set__Ref_5Value(result__0, t7)
                var inline1 bool = true
                ref_set__Ref_4bool(done__0, inline1)
                continue
            } else {
                var t8 int
                var inline5 int = ref_get__Ref_3int(i__0)
                t8 = inline5
                var t9 int = t8 - 1
                ref_set__Ref_3int(i__0, t9)
                continue
            }
        } else {
            break Loop_loop0
        }
    }
    var inline0 Value = ref_get__Ref_5Value(result__0)
    return inline0
}

func parse_list(tokens__0 *_goml_vec_Token, start__0 int) Tuple2_10Vec_5SExpr_3int {
    var acc__0 *_goml_vec_SExpr
    var inline18 *_goml_vec_SExpr = vec_new__Vec_5SExpr()
    acc__0 = inline18
    var exprs__0 *ref_Vec_5SExpr_x
    var inline17 *ref_Vec_5SExpr_x = ref__Ref_10Vec_5SExpr(acc__0)
    exprs__0 = inline17
    var i__0 *ref_int_x
    var inline16 *ref_int_x = ref__Ref_3int(start__0)
    i__0 = inline16
    var done__0 *ref_bool_x
    var inline14 bool = false
    var inline15 *ref_bool_x = ref__Ref_4bool(inline14)
    done__0 = inline15
    Loop_loop0:
    for {
        var t3 bool
        var inline13 bool = ref_get__Ref_4bool(done__0)
        t3 = inline13
        var t4 bool = !t3
        var jp0 bool
        if t4 {
            var t11 int
            var inline12 int = ref_get__Ref_3int(i__0)
            t11 = inline12
            var t12 int
            var inline11 int = vec_len__Vec_5Token(tokens__0)
            t12 = inline11
            var t13 bool = t11 < t12
            jp0 = t13
        } else {
            jp0 = false
        }
        if jp0 {
            var t5 int
            var inline10 int = ref_get__Ref_3int(i__0)
            t5 = inline10
            var mtmp0 Token = vec_get__Vec_5Token(tokens__0, t5)
            switch mtmp0._tag {
            case 1:
                var inline4 bool = true
                ref_set__Ref_4bool(done__0, inline4)
                var t6 int
                var inline3 int = ref_get__Ref_3int(i__0)
                t6 = inline3
                var t7 int = t6 + 1
                ref_set__Ref_3int(i__0, t7)
                continue
            default:
                var t8 int
                var inline9 int = ref_get__Ref_3int(i__0)
                t8 = inline9
                var mtmp2 Tuple2_5SExpr_3int = parse_expr(tokens__0, t8)
                var x0 SExpr = mtmp2._0
                var x1 int = mtmp2._1
                var t9 *_goml_vec_SExpr
                var inline8 *_goml_vec_SExpr = ref_get__Ref_10Vec_5SExpr(exprs__0)
                t9 = inline8
                var t10 *_goml_vec_SExpr = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pushed____T__SExpr(t9, x0)
                ref_set__Ref_10Vec_5SExpr(exprs__0, t10)
                ref_set__Ref_3int(i__0, x1)
                continue
            }
        } else {
            break Loop_loop0
        }
    }
    var t0 *_goml_vec_SExpr
    var inline1 *_goml_vec_SExpr = ref_get__Ref_10Vec_5SExpr(exprs__0)
    t0 = inline1
    var t1 int
    var inline0 int = ref_get__Ref_3int(i__0)
    t1 = inline0
    var t2 Tuple2_10Vec_5SExpr_3int = Tuple2_10Vec_5SExpr_3int{
        _0: t0,
        _1: t1,
    }
    return t2
}

func parse_expr(tokens__0 *_goml_vec_Token, start__0 int) Tuple2_5SExpr_3int {
    var mtmp0 Token = vec_get__Vec_5Token(tokens__0, start__0)
    switch mtmp0._tag {
    case 0:
        var t0 int = start__0 + 1
        var mtmp1 Tuple2_10Vec_5SExpr_3int = parse_list(tokens__0, t0)
        var x0 *_goml_vec_SExpr = mtmp1._0
        var x1 int = mtmp1._1
        var t1 SExpr = List{
            _0: x0,
        }
        var t2 Tuple2_5SExpr_3int = Tuple2_5SExpr_3int{
            _0: t1,
            _1: x1,
        }
        return t2
    case 1:
        var t3 SExpr = SExpr_Sym{
            _0: ")",
        }
        var t4 int = start__0 + 1
        var t5 Tuple2_5SExpr_3int = Tuple2_5SExpr_3int{
            _0: t3,
            _1: t4,
        }
        return t5
    case 2:
        var x2 string = mtmp0._p0
        var t6 SExpr = SExpr_Sym{
            _0: x2,
        }
        var t7 int = start__0 + 1
        var t8 Tuple2_5SExpr_3int = Tuple2_5SExpr_3int{
            _0: t6,
            _1: t7,
        }
        return t8
    case 3:
        var x3 int32 = mtmp0._p1
        var t9 SExpr = SExpr_Int{
            _0: x3,
        }
        var t10 int = start__0 + 1
        var t11 Tuple2_5SExpr_3int = Tuple2_5SExpr_3int{
            _0: t9,
            _1: t10,
        }
        return t11
    case 4:
        var x4 bool = mtmp0._p2
        var t12 SExpr = SExpr_Bool{
            _0: x4,
        }
        var t13 int = start__0 + 1
        var t14 Tuple2_5SExpr_3int = Tuple2_5SExpr_3int{
            _0: t12,
            _1: t13,
        }
        return t14
    default:
        panic("non-exhaustive match")
    }
}

func parse_program(tokens__0 *_goml_vec_Token) *_goml_vec_SExpr {
    var i__0 *ref_int_x
    var inline9 int = 0
    var inline10 *ref_int_x = ref__Ref_3int(inline9)
    i__0 = inline10
    var acc__0 *_goml_vec_SExpr
    var inline8 *_goml_vec_SExpr = vec_new__Vec_5SExpr()
    acc__0 = inline8
    var exprs__0 *ref_Vec_5SExpr_x
    var inline7 *ref_Vec_5SExpr_x = ref__Ref_10Vec_5SExpr(acc__0)
    exprs__0 = inline7
    Loop_loop0:
    for {
        var t0 int
        var inline6 int = ref_get__Ref_3int(i__0)
        t0 = inline6
        var t1 int
        var inline5 int = vec_len__Vec_5Token(tokens__0)
        t1 = inline5
        var t2 bool = t0 < t1
        if t2 {
            var t3 int
            var inline4 int = ref_get__Ref_3int(i__0)
            t3 = inline4
            var mtmp0 Tuple2_5SExpr_3int = parse_expr(tokens__0, t3)
            var x0 SExpr = mtmp0._0
            var x1 int = mtmp0._1
            var t4 *_goml_vec_SExpr
            var inline3 *_goml_vec_SExpr = ref_get__Ref_10Vec_5SExpr(exprs__0)
            t4 = inline3
            var t5 *_goml_vec_SExpr = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pushed____T__SExpr(t4, x0)
            ref_set__Ref_10Vec_5SExpr(exprs__0, t5)
            ref_set__Ref_3int(i__0, x1)
            continue
        } else {
            break Loop_loop0
        }
    }
    var inline0 *_goml_vec_SExpr = ref_get__Ref_10Vec_5SExpr(exprs__0)
    return inline0
}

func eval(expr__0 SExpr, local__0 *_goml_vec_Binding, global__0 *ref_Vec_7Binding_x) Value {
    switch expr__0.(type) {
    case SExpr_Int:
        var x0 int32 = expr__0.(SExpr_Int)._0
        var t0 Value = Value_Int{
            _0: x0,
        }
        return t0
    case SExpr_Bool:
        var x1 bool = expr__0.(SExpr_Bool)._0
        var t1 Value = Value_Bool{
            _0: x1,
        }
        return t1
    case SExpr_Sym:
        var x2 string = expr__0.(SExpr_Sym)._0
        var t2 *_goml_vec_Binding
        var inline2 *_goml_vec_Binding = ref_get__Ref_12Vec_7Binding(global__0)
        t2 = inline2
        var inline0 Value = env_lookup(local__0, x2)
        switch inline0.(type) {
        case Nil:
            var inline1 Value = env_lookup(t2, x2)
            return inline1
        default:
            return inline0
        }
    case List:
        var x3 *_goml_vec_SExpr = expr__0.(List)._0
        var inline3 int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_len____T__SExpr(x3)
        var inline4 bool = inline3 == 0
        if inline4 {
            return Nil{}
        } else {
            var inline5 SExpr = vec_get__Vec_5SExpr(x3, 0)
            switch inline5.(type) {
            case SExpr_Sym:
                var inline6 string = inline5.(SExpr_Sym)._0
                var inline7 Value = eval_list_sym(inline6, x3, local__0, global__0)
                return inline7
            default:
                var inline8 Value = eval(inline5, local__0, global__0)
                var inline9 *_goml_vec_Value = eval_args(x3, 1, local__0, global__0)
                var inline10 Value = apply(inline8, inline9, global__0)
                return inline10
            }
        }
    default:
        panic("non-exhaustive match")
    }
}

func eval_list_sym(name__0 string, items__0 *_goml_vec_SExpr, local__0 *_goml_vec_Binding, global__0 *ref_Vec_7Binding_x) Value {
    switch name__0 {
    case "begin":
        var t0 Value = eval_begin(items__0, 1, local__0, global__0)
        return t0
    case "define":
        var t1 int
        var inline2 int = vec_len__Vec_5SExpr(items__0)
        t1 = inline2
        var t2 bool = t1 == 3
        if t2 {
            var mtmp0 SExpr = vec_get__Vec_5SExpr(items__0, 1)
            switch mtmp0.(type) {
            case SExpr_Sym:
                var x0 string = mtmp0.(SExpr_Sym)._0
                var t3 SExpr = vec_get__Vec_5SExpr(items__0, 2)
                var value__0 Value = eval(t3, local__0, global__0)
                var env__0 *_goml_vec_Binding
                var inline1 *_goml_vec_Binding = ref_get__Ref_12Vec_7Binding(global__0)
                env__0 = inline1
                var t4 Binding = Binding{
                    name: x0,
                    value: value__0,
                }
                var updated__0 *_goml_vec_Binding = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pushed____T__Binding(env__0, t4)
                ref_set__Ref_12Vec_7Binding(global__0, updated__0)
                return value__0
            default:
                return Nil{}
            }
        } else {
            return Nil{}
        }
    case "if":
        var t5 int
        var inline6 int = vec_len__Vec_5SExpr(items__0)
        t5 = inline6
        var t6 bool = t5 == 4
        if t6 {
            var t7 SExpr = vec_get__Vec_5SExpr(items__0, 1)
            var cond__0 Value = eval(t7, local__0, global__0)
            var t8 bool
            switch cond__0.(type) {
            case Value_Int:
                var inline3 int32 = cond__0.(Value_Int)._0
                var inline4 bool = inline3 != 0
                t8 = inline4
            case Value_Bool:
                var inline5 bool = cond__0.(Value_Bool)._0
                t8 = inline5
            case Func:
                t8 = true
            case Nil:
                t8 = false
            default:
                panic("non-exhaustive match")
            }
            if t8 {
                var t9 SExpr = vec_get__Vec_5SExpr(items__0, 2)
                var t10 Value = eval(t9, local__0, global__0)
                return t10
            } else {
                var t11 SExpr = vec_get__Vec_5SExpr(items__0, 3)
                var t12 Value = eval(t11, local__0, global__0)
                return t12
            }
        } else {
            return Nil{}
        }
    case "lambda":
        var t13 int
        var inline7 int = vec_len__Vec_5SExpr(items__0)
        t13 = inline7
        var t14 bool = t13 == 3
        if t14 {
            var mtmp1 SExpr = vec_get__Vec_5SExpr(items__0, 1)
            switch mtmp1.(type) {
            case List:
                var x1 *_goml_vec_SExpr = mtmp1.(List)._0
                var params__0 *_goml_vec_string = params_from_sexprs(x1)
                var body__0 SExpr = vec_get__Vec_5SExpr(items__0, 2)
                var t15 Lambda = Lambda{
                    params: params__0,
                    body: body__0,
                    env: local__0,
                    global: global__0,
                }
                var t16 Value = Func{
                    _0: t15,
                }
                return t16
            default:
                return Nil{}
            }
        } else {
            return Nil{}
        }
    case "+":
        var t17 *_goml_vec_Value = eval_args(items__0, 1, local__0, global__0)
        var t18 Value = apply_builtin("+", t17)
        return t18
    case "-":
        var t19 *_goml_vec_Value = eval_args(items__0, 1, local__0, global__0)
        var t20 Value = apply_builtin("-", t19)
        return t20
    case "*":
        var t21 *_goml_vec_Value = eval_args(items__0, 1, local__0, global__0)
        var t22 Value = apply_builtin("*", t21)
        return t22
    case "/":
        var t23 *_goml_vec_Value = eval_args(items__0, 1, local__0, global__0)
        var t24 Value = apply_builtin("/", t23)
        return t24
    case "=":
        var t25 *_goml_vec_Value = eval_args(items__0, 1, local__0, global__0)
        var t26 Value = apply_builtin("=", t25)
        return t26
    default:
        var t27 SExpr = SExpr_Sym{
            _0: name__0,
        }
        var f__0 Value = eval(t27, local__0, global__0)
        var args__0 *_goml_vec_Value = eval_args(items__0, 1, local__0, global__0)
        switch f__0.(type) {
        case Func:
            var inline8 Lambda = f__0.(Func)._0
            var inline9 Value = apply_lambda(inline8, args__0)
            return inline9
        default:
            return Nil{}
        }
    }
}

func eval_begin(items__0 *_goml_vec_SExpr, start__0 int, local__0 *_goml_vec_Binding, global__0 *ref_Vec_7Binding_x) Value {
    var i__0 *ref_int_x
    var inline8 *ref_int_x = ref__Ref_3int(start__0)
    i__0 = inline8
    var last__0 *ref_Value_x
    var inline7 *ref_Value_x = ref__Ref_5Value(Nil{})
    last__0 = inline7
    Loop_loop0:
    for {
        var t0 int
        var inline6 int = ref_get__Ref_3int(i__0)
        t0 = inline6
        var t1 int
        var inline5 int = vec_len__Vec_5SExpr(items__0)
        t1 = inline5
        var t2 bool = t0 < t1
        if t2 {
            var t3 int
            var inline4 int = ref_get__Ref_3int(i__0)
            t3 = inline4
            var t4 SExpr = vec_get__Vec_5SExpr(items__0, t3)
            var v__0 Value = eval(t4, local__0, global__0)
            ref_set__Ref_5Value(last__0, v__0)
            var t5 int
            var inline2 int = ref_get__Ref_3int(i__0)
            t5 = inline2
            var t6 int = t5 + 1
            ref_set__Ref_3int(i__0, t6)
            continue
        } else {
            break Loop_loop0
        }
    }
    var inline0 Value = ref_get__Ref_5Value(last__0)
    return inline0
}

func params_from_sexprs(items__0 *_goml_vec_SExpr) *_goml_vec_string {
    var i__0 *ref_int_x
    var inline12 int = 0
    var inline13 *ref_int_x = ref__Ref_3int(inline12)
    i__0 = inline13
    var acc__0 *_goml_vec_string
    var inline11 *_goml_vec_string = vec_new__Vec_6string()
    acc__0 = inline11
    var params__0 *ref_Vec_6string_x
    var inline10 *ref_Vec_6string_x = ref__Ref_11Vec_6string(acc__0)
    params__0 = inline10
    Loop_loop0:
    for {
        var t0 int
        var inline9 int = ref_get__Ref_3int(i__0)
        t0 = inline9
        var t1 int
        var inline8 int = vec_len__Vec_5SExpr(items__0)
        t1 = inline8
        var t2 bool = t0 < t1
        if t2 {
            var t3 int
            var inline7 int = ref_get__Ref_3int(i__0)
            t3 = inline7
            var mtmp0 SExpr = vec_get__Vec_5SExpr(items__0, t3)
            switch mtmp0.(type) {
            case SExpr_Sym:
                var x0 string = mtmp0.(SExpr_Sym)._0
                var t4 *_goml_vec_string
                var inline4 *_goml_vec_string = ref_get__Ref_11Vec_6string(params__0)
                t4 = inline4
                var t5 *_goml_vec_string = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pushed____T__string(t4, x0)
                ref_set__Ref_11Vec_6string(params__0, t5)
                var t6 int
                var inline2 int = ref_get__Ref_3int(i__0)
                t6 = inline2
                var t7 int = t6 + 1
                ref_set__Ref_3int(i__0, t7)
                continue
            default:
                var t8 int
                var inline6 int = ref_get__Ref_3int(i__0)
                t8 = inline6
                var t9 int = t8 + 1
                ref_set__Ref_3int(i__0, t9)
                continue
            }
        } else {
            break Loop_loop0
        }
    }
    var inline0 *_goml_vec_string = ref_get__Ref_11Vec_6string(params__0)
    return inline0
}

func eval_args(items__0 *_goml_vec_SExpr, start__0 int, local__0 *_goml_vec_Binding, global__0 *ref_Vec_7Binding_x) *_goml_vec_Value {
    var i__0 *ref_int_x
    var inline10 *ref_int_x = ref__Ref_3int(start__0)
    i__0 = inline10
    var acc__0 *_goml_vec_Value
    var inline9 *_goml_vec_Value = vec_new__Vec_5Value()
    acc__0 = inline9
    var args__0 *ref_Vec_5Value_x
    var inline8 *ref_Vec_5Value_x = ref__Ref_10Vec_5Value(acc__0)
    args__0 = inline8
    Loop_loop0:
    for {
        var t0 int
        var inline7 int = ref_get__Ref_3int(i__0)
        t0 = inline7
        var t1 int
        var inline6 int = vec_len__Vec_5SExpr(items__0)
        t1 = inline6
        var t2 bool = t0 < t1
        if t2 {
            var t3 int
            var inline5 int = ref_get__Ref_3int(i__0)
            t3 = inline5
            var t4 SExpr = vec_get__Vec_5SExpr(items__0, t3)
            var v__0 Value = eval(t4, local__0, global__0)
            var t5 *_goml_vec_Value
            var inline4 *_goml_vec_Value = ref_get__Ref_10Vec_5Value(args__0)
            t5 = inline4
            var t6 *_goml_vec_Value = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pushed____T__Value(t5, v__0)
            ref_set__Ref_10Vec_5Value(args__0, t6)
            var t7 int
            var inline2 int = ref_get__Ref_3int(i__0)
            t7 = inline2
            var t8 int = t7 + 1
            ref_set__Ref_3int(i__0, t8)
            continue
        } else {
            break Loop_loop0
        }
    }
    var inline0 *_goml_vec_Value = ref_get__Ref_10Vec_5Value(args__0)
    return inline0
}

func apply_builtin(name__0 string, args__0 *_goml_vec_Value) Value {
    switch name__0 {
    case "=":
        var t0 int
        var inline0 int = vec_len__Vec_5Value(args__0)
        t0 = inline0
        var t1 bool = t0 == 2
        if t1 {
            var t2 Value = vec_get__Vec_5Value(args__0, 0)
            var t3 Value = vec_get__Vec_5Value(args__0, 1)
            switch t3.(type) {
            case Value_Int:
                var x0 int32 = t3.(Value_Int)._0
                switch t2.(type) {
                case Value_Int:
                    var x1 int32 = t2.(Value_Int)._0
                    var t4 bool = x1 == x0
                    var t5 Value = Value_Bool{
                        _0: t4,
                    }
                    return t5
                default:
                    var t6 Value = Value_Bool{
                        _0: false,
                    }
                    return t6
                }
            case Value_Bool:
                var x2 bool = t3.(Value_Bool)._0
                switch t2.(type) {
                case Value_Bool:
                    var x3 bool = t2.(Value_Bool)._0
                    var t7 bool = x3 == x2
                    var t8 Value = Value_Bool{
                        _0: t7,
                    }
                    return t8
                default:
                    var t9 Value = Value_Bool{
                        _0: false,
                    }
                    return t9
                }
            default:
                var t10 Value = Value_Bool{
                    _0: false,
                }
                return t10
            }
        } else {
            var t11 Value = Value_Bool{
                _0: false,
            }
            return t11
        }
    case "+":
        var i__0 *ref_int_x
        var inline13 int = 0
        var inline14 *ref_int_x = ref__Ref_3int(inline13)
        i__0 = inline14
        var acc__0 *ref_int32_x
        var inline11 int32 = 0
        var inline12 *ref_int32_x = ref__Ref_5int32(inline11)
        acc__0 = inline12
        Loop_loop0:
        for {
            var t14 int
            var inline10 int = ref_get__Ref_3int(i__0)
            t14 = inline10
            var t15 int
            var inline9 int = vec_len__Vec_5Value(args__0)
            t15 = inline9
            var t16 bool = t14 < t15
            if t16 {
                var t17 int
                var inline8 int = ref_get__Ref_3int(i__0)
                t17 = inline8
                var mtmp0 Value = vec_get__Vec_5Value(args__0, t17)
                switch mtmp0.(type) {
                case Value_Int:
                    var x4 int32 = mtmp0.(Value_Int)._0
                    var t18 int32
                    var inline5 int32 = ref_get__Ref_5int32(acc__0)
                    t18 = inline5
                    var t19 int32 = t18 + x4
                    ref_set__Ref_5int32(acc__0, t19)
                    var t20 int
                    var inline3 int = ref_get__Ref_3int(i__0)
                    t20 = inline3
                    var t21 int = t20 + 1
                    ref_set__Ref_3int(i__0, t21)
                    continue
                default:
                    var t22 int
                    var inline7 int = ref_get__Ref_3int(i__0)
                    t22 = inline7
                    var t23 int = t22 + 1
                    ref_set__Ref_3int(i__0, t23)
                    continue
                }
            } else {
                break Loop_loop0
            }
        }
        var t12 int32
        var inline1 int32 = ref_get__Ref_5int32(acc__0)
        t12 = inline1
        var t13 Value = Value_Int{
            _0: t12,
        }
        return t13
    case "*":
        var i__1 *ref_int_x = _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__isize(0)
        var acc__1 *ref_int32_x
        var inline25 int32 = 1
        var inline26 *ref_int32_x = ref__Ref_5int32(inline25)
        acc__1 = inline26
        Loop_loop1:
        for {
            var t26 int
            var inline24 int = ref_get__Ref_3int(i__1)
            t26 = inline24
            var t27 int
            var inline23 int = vec_len__Vec_5Value(args__0)
            t27 = inline23
            var t28 bool = t26 < t27
            if t28 {
                var t29 int
                var inline22 int = ref_get__Ref_3int(i__1)
                t29 = inline22
                var mtmp2 Value = vec_get__Vec_5Value(args__0, t29)
                switch mtmp2.(type) {
                case Value_Int:
                    var x5 int32 = mtmp2.(Value_Int)._0
                    var t30 int32
                    var inline19 int32 = ref_get__Ref_5int32(acc__1)
                    t30 = inline19
                    var t31 int32 = t30 * x5
                    ref_set__Ref_5int32(acc__1, t31)
                    var t32 int
                    var inline17 int = ref_get__Ref_3int(i__1)
                    t32 = inline17
                    var t33 int = t32 + 1
                    ref_set__Ref_3int(i__1, t33)
                    continue
                default:
                    var t34 int
                    var inline21 int = ref_get__Ref_3int(i__1)
                    t34 = inline21
                    var t35 int = t34 + 1
                    ref_set__Ref_3int(i__1, t35)
                    continue
                }
            } else {
                break Loop_loop1
            }
        }
        var t24 int32
        var inline15 int32 = ref_get__Ref_5int32(acc__1)
        t24 = inline15
        var t25 Value = Value_Int{
            _0: t24,
        }
        return t25
    case "-":
        var mtmp4 int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_len____T__Value(args__0)
        switch mtmp4 {
        case 1:
            var mtmp5 Value = vec_get__Vec_5Value(args__0, 0)
            switch mtmp5.(type) {
            case Value_Int:
                var x6 int32 = mtmp5.(Value_Int)._0
                var t36 int32 = 0 - x6
                var t37 Value = Value_Int{
                    _0: t36,
                }
                return t37
            default:
                return Nil{}
            }
        case 2:
            var t38 Value = vec_get__Vec_5Value(args__0, 0)
            var t39 Value = vec_get__Vec_5Value(args__0, 1)
            switch t39.(type) {
            case Value_Int:
                var x7 int32 = t39.(Value_Int)._0
                switch t38.(type) {
                case Value_Int:
                    var x8 int32 = t38.(Value_Int)._0
                    var t40 int32 = x8 - x7
                    var t41 Value = Value_Int{
                        _0: t40,
                    }
                    return t41
                default:
                    return Nil{}
                }
            default:
                return Nil{}
            }
        default:
            return Nil{}
        }
    case "/":
        var t42 int = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_len____T__Value(args__0)
        var t43 bool = t42 == 2
        if t43 {
            var t44 Value = vec_get__Vec_5Value(args__0, 0)
            var t45 Value = vec_get__Vec_5Value(args__0, 1)
            switch t45.(type) {
            case Value_Int:
                var x9 int32 = t45.(Value_Int)._0
                switch t44.(type) {
                case Value_Int:
                    var x10 int32 = t44.(Value_Int)._0
                    var t46 int32 = x10 / x9
                    var t47 Value = Value_Int{
                        _0: t46,
                    }
                    return t47
                default:
                    return Nil{}
                }
            default:
                return Nil{}
            }
        } else {
            return Nil{}
        }
    default:
        return Nil{}
    }
}

func apply(func__0 Value, args__0 *_goml_vec_Value, global__0 *ref_Vec_7Binding_x) Value {
    switch func__0.(type) {
    case Func:
        var x0 Lambda = func__0.(Func)._0
        var t0 Value = apply_lambda(x0, args__0)
        return t0
    default:
        return Nil{}
    }
}

func apply_lambda(lambda__0 Lambda, args__0 *_goml_vec_Value) Value {
    var t0 *_goml_vec_Binding = lambda__0.env
    var env__0 *ref_Vec_7Binding_x
    var inline13 *ref_Vec_7Binding_x = ref__Ref_12Vec_7Binding(t0)
    env__0 = inline13
    var i__0 *ref_int_x
    var inline11 int = 0
    var inline12 *ref_int_x = ref__Ref_3int(inline11)
    i__0 = inline12
    Loop_loop0:
    for {
        var t5 int
        var inline10 int = ref_get__Ref_3int(i__0)
        t5 = inline10
        var t6 *_goml_vec_string = lambda__0.params
        var t7 int
        var inline9 int = vec_len__Vec_6string(t6)
        t7 = inline9
        var t8 bool = t5 < t7
        var jp0 bool
        if t8 {
            var t16 int
            var inline8 int = ref_get__Ref_3int(i__0)
            t16 = inline8
            var t17 int
            var inline7 int = vec_len__Vec_5Value(args__0)
            t17 = inline7
            var t18 bool = t16 < t17
            jp0 = t18
        } else {
            jp0 = false
        }
        if jp0 {
            var t9 *_goml_vec_string = lambda__0.params
            var t10 int
            var inline6 int = ref_get__Ref_3int(i__0)
            t10 = inline6
            var name__0 string = vec_get__Vec_6string(t9, t10)
            var t11 int
            var inline5 int = ref_get__Ref_3int(i__0)
            t11 = inline5
            var value__0 Value = vec_get__Vec_5Value(args__0, t11)
            var t12 *_goml_vec_Binding
            var inline4 *_goml_vec_Binding = ref_get__Ref_12Vec_7Binding(env__0)
            t12 = inline4
            var t13 Binding = Binding{
                name: name__0,
                value: value__0,
            }
            var updated__0 *_goml_vec_Binding = _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pushed____T__Binding(t12, t13)
            ref_set__Ref_12Vec_7Binding(env__0, updated__0)
            var t14 int
            var inline2 int = ref_get__Ref_3int(i__0)
            t14 = inline2
            var t15 int = t14 + 1
            ref_set__Ref_3int(i__0, t15)
            continue
        } else {
            break Loop_loop0
        }
    }
    var t1 SExpr = lambda__0.body
    var t2 *_goml_vec_Binding
    var inline0 *_goml_vec_Binding = ref_get__Ref_12Vec_7Binding(env__0)
    t2 = inline0
    var t3 *ref_Vec_7Binding_x = lambda__0.global
    var t4 Value = eval(t1, t2, t3)
    return t4
}

func main0() struct{} {
    var t0 *_goml_vec_Binding
    var inline15 *_goml_vec_Binding = vec_new__Vec_7Binding()
    t0 = inline15
    var global__0 *ref_Vec_7Binding_x
    var inline14 *ref_Vec_7Binding_x = ref__Ref_12Vec_7Binding(t0)
    global__0 = inline14
    var program__0 string = "(begin (define fact (lambda (n) (if (= n 0) 1 (* n (fact (- n 1)))))) (define add3 (lambda (a b c) (+ a (+ b c)))) (fact 6))"
    var t1 *_goml_vec_Token = lex(program__0)
    var exprs__0 *_goml_vec_SExpr = parse_program(t1)
    var t2 SExpr = vec_get__Vec_5SExpr(exprs__0, 0)
    var t3 *_goml_vec_Binding
    var inline13 *_goml_vec_Binding = vec_new__Vec_7Binding()
    t3 = inline13
    var result__0 Value = eval(t2, t3, global__0)
    var t4 string
    switch result__0.(type) {
    case Value_Int:
        var inline9 int32 = result__0.(Value_Int)._0
        var inline10 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline9)
        t4 = inline10
    case Value_Bool:
        var inline11 bool = result__0.(Value_Bool)._0
        var inline12 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(inline11)
        t4 = inline12
    case Func:
        t4 = "<lambda>"
    case Nil:
        t4 = "nil"
    default:
        panic("non-exhaustive match")
    }
    var inline7 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t4)
    _goml_runtime_core_string_println(inline7)
    var t5 *_goml_vec_Token = lex("(add3 10 20 30)")
    var exprs2__0 *_goml_vec_SExpr = parse_program(t5)
    var t6 SExpr = vec_get__Vec_5SExpr(exprs2__0, 0)
    var t7 *_goml_vec_Binding
    var inline6 *_goml_vec_Binding = vec_new__Vec_7Binding()
    t7 = inline6
    var result2__0 Value = eval(t6, t7, global__0)
    var t8 string
    switch result2__0.(type) {
    case Value_Int:
        var inline2 int32 = result2__0.(Value_Int)._0
        var inline3 string = _goml_m_inherent_i_i32_i_i32_i_to__string(inline2)
        t8 = inline3
    case Value_Bool:
        var inline4 bool = result2__0.(Value_Bool)._0
        var inline5 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(inline4)
        t8 = inline5
    case Func:
        t8 = "<lambda>"
    case Nil:
        t8 = "nil"
    default:
        panic("non-exhaustive match")
    }
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t8)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_inherent_i_string_i_string_i_len(self__0 string) int {
    var t0 int = _goml_runtime_core_string_len(self__0)
    return t0
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__isize(value__0 int) *ref_int_x {
    var t0 *ref_int_x = ref__Ref_3int(value__0)
    return t0
}

func _goml_m_inherent_i_Ref_i_Ref_l_T_r__i_new____T__bool(value__0 bool) *ref_bool_x {
    var t0 *ref_bool_x = ref__Ref_4bool(value__0)
    return t0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pushed____T__Token(self__0 *_goml_vec_Token, elem__0 Token) *_goml_vec_Token {
    var t0 int
    var inline4 int = vec_len__Vec_5Token(self__0)
    t0 = inline4
    var t1 int = t0 + 1
    var result__0 *_goml_vec_Token
    var inline3 *_goml_vec_Token = vec_with_capacity__Vec_5Token(t1)
    result__0 = inline3
    var index__0 int = 0
    Loop_loop0:
    for {
        var t2 int
        var inline2 int = vec_len__Vec_5Token(self__0)
        t2 = inline2
        var t3 bool = index__0 < t2
        if t3 {
            var t4 Token = vec_get__Vec_5Token(self__0, index__0)
            vec_push__Vec_5Token(result__0, t4)
            var compound_old0 int = index__0
            var compound_value0 int = 1
            var t5 int = compound_old0 + compound_value0
            index__0 = t5
            continue
        } else {
            break Loop_loop0
        }
    }
    vec_push__Vec_5Token(result__0, elem__0)
    return result__0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pushed____T__SExpr(self__0 *_goml_vec_SExpr, elem__0 SExpr) *_goml_vec_SExpr {
    var t0 int
    var inline4 int = vec_len__Vec_5SExpr(self__0)
    t0 = inline4
    var t1 int = t0 + 1
    var result__0 *_goml_vec_SExpr
    var inline3 *_goml_vec_SExpr = vec_with_capacity__Vec_5SExpr(t1)
    result__0 = inline3
    var index__0 int = 0
    Loop_loop0:
    for {
        var t2 int
        var inline2 int = vec_len__Vec_5SExpr(self__0)
        t2 = inline2
        var t3 bool = index__0 < t2
        if t3 {
            var t4 SExpr = vec_get__Vec_5SExpr(self__0, index__0)
            vec_push__Vec_5SExpr(result__0, t4)
            var compound_old0 int = index__0
            var compound_value0 int = 1
            var t5 int = compound_old0 + compound_value0
            index__0 = t5
            continue
        } else {
            break Loop_loop0
        }
    }
    vec_push__Vec_5SExpr(result__0, elem__0)
    return result__0
}

func _goml_m_inherent_i_i32_i_i32_i_to__string(self__0 int32) string {
    var inline0 int64 = int64(int32(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func _goml_m_trait__impl_i_ToString_i_bool_i_to__string(self__0 bool) string {
    var t0 string = _goml_runtime_core_bool_to_string(self__0)
    return t0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_len____T__SExpr(self__0 *_goml_vec_SExpr) int {
    var t0 int = vec_len__Vec_5SExpr(self__0)
    return t0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pushed____T__Binding(self__0 *_goml_vec_Binding, elem__0 Binding) *_goml_vec_Binding {
    var t0 int
    var inline4 int = vec_len__Vec_7Binding(self__0)
    t0 = inline4
    var t1 int = t0 + 1
    var result__0 *_goml_vec_Binding
    var inline3 *_goml_vec_Binding = vec_with_capacity__Vec_7Binding(t1)
    result__0 = inline3
    var index__0 int = 0
    Loop_loop0:
    for {
        var t2 int
        var inline2 int = vec_len__Vec_7Binding(self__0)
        t2 = inline2
        var t3 bool = index__0 < t2
        if t3 {
            var t4 Binding = vec_get__Vec_7Binding(self__0, index__0)
            vec_push__Vec_7Binding(result__0, t4)
            var compound_old0 int = index__0
            var compound_value0 int = 1
            var t5 int = compound_old0 + compound_value0
            index__0 = t5
            continue
        } else {
            break Loop_loop0
        }
    }
    vec_push__Vec_7Binding(result__0, elem__0)
    return result__0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pushed____T__string(self__0 *_goml_vec_string, elem__0 string) *_goml_vec_string {
    var t0 int
    var inline4 int = vec_len__Vec_6string(self__0)
    t0 = inline4
    var t1 int = t0 + 1
    var result__0 *_goml_vec_string
    var inline3 *_goml_vec_string = vec_with_capacity__Vec_6string(t1)
    result__0 = inline3
    var index__0 int = 0
    Loop_loop0:
    for {
        var t2 int
        var inline2 int = vec_len__Vec_6string(self__0)
        t2 = inline2
        var t3 bool = index__0 < t2
        if t3 {
            var t4 string = vec_get__Vec_6string(self__0, index__0)
            vec_push__Vec_6string(result__0, t4)
            var compound_old0 int = index__0
            var compound_value0 int = 1
            var t5 int = compound_old0 + compound_value0
            index__0 = t5
            continue
        } else {
            break Loop_loop0
        }
    }
    vec_push__Vec_6string(result__0, elem__0)
    return result__0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_pushed____T__Value(self__0 *_goml_vec_Value, elem__0 Value) *_goml_vec_Value {
    var t0 int
    var inline4 int = vec_len__Vec_5Value(self__0)
    t0 = inline4
    var t1 int = t0 + 1
    var result__0 *_goml_vec_Value
    var inline3 *_goml_vec_Value = vec_with_capacity__Vec_5Value(t1)
    result__0 = inline3
    var index__0 int = 0
    Loop_loop0:
    for {
        var t2 int
        var inline2 int = vec_len__Vec_5Value(self__0)
        t2 = inline2
        var t3 bool = index__0 < t2
        if t3 {
            var t4 Value = vec_get__Vec_5Value(self__0, index__0)
            vec_push__Vec_5Value(result__0, t4)
            var compound_old0 int = index__0
            var compound_value0 int = 1
            var t5 int = compound_old0 + compound_value0
            index__0 = t5
            continue
        } else {
            break Loop_loop0
        }
    }
    vec_push__Vec_5Value(result__0, elem__0)
    return result__0
}

func _goml_m_inherent_i_Vec_i_Vec_l_T_r__i_len____T__Value(self__0 *_goml_vec_Value) int {
    var t0 int = vec_len__Vec_5Value(self__0)
    return t0
}

func string_get(value__0 string, index__0 int) rune {
    var mtmp0 Tuple3_4bool_4char_3int = string_decode_utf8_at(value__0, index__0)
    var x0 bool = mtmp0._0
    var x1 rune = mtmp0._1
    if x0 {
        return x1
    } else {
        var t0 rune = _goml_runtime_core_string_get("", -1)
        return t0
    }
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

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func string_decode_utf8_at(value__0 string, index__0 int) Tuple3_4bool_4char_3int {
    var length__0 int = _goml_m_inherent_i_string_i_string_i_byte__len(value__0)
    var t0 bool = index__0 < 0
    var jp0 bool
    if t0 {
        jp0 = true
    } else {
        var t63 bool = index__0 >= length__0
        jp0 = t63
    }
    if jp0 {
        var inline25 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
            _0: false,
            _1: 0,
            _2: 0,
        }
        return inline25
    } else {
        var t1 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, index__0)
        var first__0 uint32 = uint32(uint8(t1))
        var t2 bool = first__0 < 128
        if t2 {
            var inline0 int = 1
            var inline1 Option__char = __goml_builtin_char_from_uint32(first__0)
            switch inline1 != Option__char(0) {
            case false:
                var inline2 Tuple3_4bool_4char_3int = utf8_invalid_decode()
                return inline2
            case true:
                var inline3 rune = rune(uint64(inline1) - 1)
                var inline4 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                    _0: true,
                    _1: inline3,
                    _2: inline0,
                }
                return inline4
            default:
                panic("non-exhaustive match")
            }
        } else {
            var t3 bool = first__0 < 194
            if t3 {
                var inline5 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                    _0: false,
                    _1: 0,
                    _2: 0,
                }
                return inline5
            } else {
                var t4 bool = first__0 < 224
                if t4 {
                    var t5 int = length__0 - index__0
                    var t6 bool = t5 < 2
                    if t6 {
                        var inline15 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                            _0: false,
                            _1: 0,
                            _2: 0,
                        }
                        return inline15
                    } else {
                        var t7 int = index__0 + 1
                        var t8 uint8
                        var inline14 uint8 = _goml_runtime_core_string_byte_get(value__0, t7)
                        t8 = inline14
                        var second__0 uint32 = uint32(uint8(t8))
                        var t9 bool
                        var inline12 bool = second__0 < 128
                        if inline12 {
                            t9 = true
                        } else {
                            var inline13 bool = second__0 > 191
                            t9 = inline13
                        }
                        if t9 {
                            var inline6 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                                _0: false,
                                _1: 0,
                                _2: 0,
                            }
                            return inline6
                        } else {
                            var t10 uint32 = first__0 & 31
                            var t11 uint32 = t10 << 6
                            var t12 uint32 = second__0 & 63
                            var t13 uint32 = t11 | t12
                            var inline7 int = 2
                            var inline8 Option__char = __goml_builtin_char_from_uint32(t13)
                            switch inline8 != Option__char(0) {
                            case false:
                                var inline9 Tuple3_4bool_4char_3int = utf8_invalid_decode()
                                return inline9
                            case true:
                                var inline10 rune = rune(uint64(inline8) - 1)
                                var inline11 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                                    _0: true,
                                    _1: inline10,
                                    _2: inline7,
                                }
                                return inline11
                            default:
                                panic("non-exhaustive match")
                            }
                        }
                    }
                } else {
                    var t14 bool = first__0 < 240
                    if t14 {
                        var t15 int = length__0 - index__0
                        var t16 bool = t15 < 3
                        if t16 {
                            var inline24 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                                _0: false,
                                _1: 0,
                                _2: 0,
                            }
                            return inline24
                        } else {
                            var t17 int = index__0 + 1
                            var t18 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, t17)
                            var second__1 uint32 = uint32(uint8(t18))
                            var t19 int = index__0 + 2
                            var t20 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, t19)
                            var third__0 uint32 = uint32(uint8(t20))
                            var t21 bool = utf8_invalid_continuation(second__1)
                            var jp1 bool
                            if t21 {
                                jp1 = true
                            } else {
                                var inline22 bool = third__0 < 128
                                if inline22 {
                                    jp1 = true
                                } else {
                                    var inline23 bool = third__0 > 191
                                    jp1 = inline23
                                }
                            }
                            var jp2 bool
                            if jp1 {
                                jp2 = true
                            } else {
                                var t31 bool = first__0 == 224
                                if t31 {
                                    var t32 bool = second__1 < 160
                                    jp2 = t32
                                } else {
                                    jp2 = false
                                }
                            }
                            var jp3 bool
                            if jp2 {
                                jp3 = true
                            } else {
                                var t29 bool = first__0 == 237
                                if t29 {
                                    var t30 bool = second__1 >= 160
                                    jp3 = t30
                                } else {
                                    jp3 = false
                                }
                            }
                            if jp3 {
                                var inline16 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                                    _0: false,
                                    _1: 0,
                                    _2: 0,
                                }
                                return inline16
                            } else {
                                var t22 uint32 = first__0 & 15
                                var t23 uint32 = t22 << 12
                                var t24 uint32 = second__1 & 63
                                var t25 uint32 = t24 << 6
                                var t26 uint32 = t23 | t25
                                var t27 uint32 = third__0 & 63
                                var t28 uint32 = t26 | t27
                                var inline17 int = 3
                                var inline18 Option__char = __goml_builtin_char_from_uint32(t28)
                                switch inline18 != Option__char(0) {
                                case false:
                                    var inline19 Tuple3_4bool_4char_3int = utf8_invalid_decode()
                                    return inline19
                                case true:
                                    var inline20 rune = rune(uint64(inline18) - 1)
                                    var inline21 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
                                        _0: true,
                                        _1: inline20,
                                        _2: inline17,
                                    }
                                    return inline21
                                default:
                                    panic("non-exhaustive match")
                                }
                            }
                        }
                    } else {
                        var t33 bool = first__0 < 245
                        if t33 {
                            var t34 int = length__0 - index__0
                            var t35 bool = t34 < 4
                            if t35 {
                                var t61 Tuple3_4bool_4char_3int = utf8_invalid_decode()
                                return t61
                            } else {
                                var t36 int = index__0 + 1
                                var t37 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, t36)
                                var second__2 uint32 = uint32(uint8(t37))
                                var t38 int = index__0 + 2
                                var t39 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, t38)
                                var third__1 uint32 = uint32(uint8(t39))
                                var t40 int = index__0 + 3
                                var t41 uint8 = _goml_m_inherent_i_string_i_string_i_byte__get(value__0, t40)
                                var fourth__0 uint32 = uint32(uint8(t41))
                                var t42 bool = utf8_invalid_continuation(second__2)
                                var jp4 bool
                                if t42 {
                                    jp4 = true
                                } else {
                                    var t60 bool = utf8_invalid_continuation(third__1)
                                    jp4 = t60
                                }
                                var jp5 bool
                                if jp4 {
                                    jp5 = true
                                } else {
                                    var t59 bool = utf8_invalid_continuation(fourth__0)
                                    jp5 = t59
                                }
                                var jp6 bool
                                if jp5 {
                                    jp6 = true
                                } else {
                                    var t57 bool = first__0 == 240
                                    if t57 {
                                        var t58 bool = second__2 < 144
                                        jp6 = t58
                                    } else {
                                        jp6 = false
                                    }
                                }
                                var jp7 bool
                                if jp6 {
                                    jp7 = true
                                } else {
                                    var t55 bool = first__0 == 244
                                    if t55 {
                                        var t56 bool = second__2 > 143
                                        jp7 = t56
                                    } else {
                                        jp7 = false
                                    }
                                }
                                if jp7 {
                                    var t43 Tuple3_4bool_4char_3int = utf8_invalid_decode()
                                    return t43
                                } else {
                                    var t44 uint32 = first__0 & 7
                                    var t45 uint32 = t44 << 18
                                    var t46 uint32 = second__2 & 63
                                    var t47 uint32 = t46 << 12
                                    var t48 uint32 = t45 | t47
                                    var t49 uint32 = third__1 & 63
                                    var t50 uint32 = t49 << 6
                                    var t51 uint32 = t48 | t50
                                    var t52 uint32 = fourth__0 & 63
                                    var t53 uint32 = t51 | t52
                                    var t54 Tuple3_4bool_4char_3int = utf8_valid_decode(t53, 4)
                                    return t54
                                }
                            }
                        } else {
                            var t62 Tuple3_4bool_4char_3int = utf8_invalid_decode()
                            return t62
                        }
                    }
                }
            }
        }
    }
}

func utf8_valid_scalar(value__0 uint32) bool {
    var t0 bool = value__0 <= 1114111
    if t0 {
        var t1 bool = value__0 >= 55296
        var jp0 bool
        if t1 {
            var t3 bool = value__0 <= 57343
            jp0 = t3
        } else {
            jp0 = false
        }
        var t2 bool = !jp0
        return t2
    } else {
        return false
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

func _goml_m_inherent_i_string_i_string_i_byte__len(self__0 string) int {
    var t0 int = _goml_runtime_core_string_len(self__0)
    return t0
}

func utf8_invalid_decode() Tuple3_4bool_4char_3int {
    var t0 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
        _0: false,
        _1: 0,
        _2: 0,
    }
    return t0
}

func _goml_m_inherent_i_string_i_string_i_byte__get(self__0 string, index__0 int) uint8 {
    var t0 uint8 = _goml_runtime_core_string_byte_get(self__0, index__0)
    return t0
}

func utf8_valid_decode(value__0 uint32, width__0 int) Tuple3_4bool_4char_3int {
    var commute_field0 rune
    var inline1 bool = utf8_valid_scalar(value__0)
    if inline1 {
        var inline2 Tuple2_4bool_4char = _goml_runtime_core_char_from_uint32(value__0)
        var inline3 rune = inline2._1
        commute_field0 = inline3
        var t0 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
            _0: true,
            _1: commute_field0,
            _2: width__0,
        }
        return t0
    } else {
        var inline0 Tuple3_4bool_4char_3int = Tuple3_4bool_4char_3int{
            _0: false,
            _1: 0,
            _2: 0,
        }
        return inline0
    }
}

func utf8_invalid_continuation(value__0 uint32) bool {
    var t0 bool = value__0 < 128
    if t0 {
        return true
    } else {
        var t1 bool = value__0 > 191
        return t1
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

func __goml_builtin_char_from_uint32(value__0 uint32) Option__char {
    var t0 bool
    var inline0 bool = value__0 <= 1114111
    if inline0 {
        var inline1 bool = value__0 >= 55296
        var inline2 bool
        if inline1 {
            var inline4 bool = value__0 <= 57343
            inline2 = inline4
        } else {
            inline2 = false
        }
        var inline3 bool = !inline2
        t0 = inline3
    } else {
        t0 = false
    }
    if t0 {
        var mtmp0 Tuple2_4bool_4char = _goml_runtime_core_char_from_uint32(value__0)
        var x0 rune = mtmp0._1
        var t1 Option__char = Option__char(uint64(uint32(x0)) + 1)
        return t1
    } else {
        return Option__char(0)
    }
}

func main() {
    main0()
}
