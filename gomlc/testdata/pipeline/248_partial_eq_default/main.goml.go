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

type ref_NoDefault_x struct {
    value NoDefault
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

type Settings struct {
    enabled bool
    retries int
    label string
}

type NoDefault struct {
    value int
}

type Empty struct {}

type PairDefaults__isize struct {
    first int
    second int
    nested Option__isize
}

type Ordering uint8

type State struct {
    _p0 int
    _tag uint8
}

type Message interface {
    isMessage()
}

type Message_Data struct {
    _0 string
    _1 int
}

func (_ Message_Data) isMessage() {}

type Message_Empty struct {}

func (_ Message_Empty) isMessage() {}

type Event interface {
    isEvent()
}

type Event_Data struct {
    _0 string
    _1 int
}

func (_ Event_Data) isEvent() {}

type Event_Empty struct {}

func (_ Event_Empty) isEvent() {}

type Option__isize struct {
    _p0 int
    _tag uint8
}

type Lazy__NoDefault struct {
    _p0 NoDefault
    _tag uint8
}

type Selected__isize interface {
    isSelected__isize()
}

type Values struct {
    _0 int
    _1 int
}

func (_ Values) isSelected__isize() {}

type Ignored struct {
    _0 *ref_NoDefault_x
}

func (_ Ignored) isSelected__isize() {}

func _goml_m_trait__impl_i_PartialEq_i_Settings_i_eq(self__0 Settings, other__0 Settings) bool {
    var jp0 bool
    var t4 bool = self__0.enabled
    var t5 bool = other__0.enabled
    var inline2 bool = t4 == t5
    jp0 = inline2
    var jp1 bool
    if jp0 {
        var t2 int = self__0.retries
        var t3 int = other__0.retries
        var inline1 bool = t2 == t3
        jp1 = inline1
    } else {
        jp1 = false
    }
    if jp1 {
        var t0 string = self__0.label
        var t1 string = other__0.label
        var inline0 bool = t0 == t1
        return inline0
    } else {
        return false
    }
}

func _goml_m_trait__impl_i_Default_i_Settings_i_default() Settings {
    var t0 bool
    t0 = false
    var t1 int
    t1 = 0
    var t2 string
    t2 = ""
    var t3 Settings = Settings{
        enabled: t0,
        retries: t1,
        label: t2,
    }
    return t3
}

func _goml_m_trait__impl_i_Debug_i_Settings_i_debug(self__0 Settings) string {
    var x0 bool = self__0.enabled
    var x1 int = self__0.retries
    var x2 string = self__0.label
    var t0 string = "Settings { " + "enabled: "
    var t1 string
    var inline2 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(x0)
    t1 = inline2
    var t2 string = t0 + t1
    var t3 string = t2 + ", "
    var t4 string = t3 + "retries: "
    var t5 string
    var inline1 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(x1)
    t5 = inline1
    var t6 string = t4 + t5
    var t7 string = t6 + ", "
    var t8 string = t7 + "label: "
    var t9 string
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(x2)
    t9 = inline0
    var t10 string = t8 + t9
    var t11 string = t10 + " }"
    return t11
}

func _goml_m_trait__impl_i_Default_i_Empty_i_default() Empty {
    var t0 Empty = Empty{}
    return t0
}

func _goml_m_trait__impl_i_Debug_i_Empty_i_debug(self__0 Empty) string {
    return "Empty {}"
}

func _goml_m_trait__impl_i_Debug_i_State_i_debug(self__0 State) string {
    switch self__0._tag {
    case 0:
        return "State::Idle"
    case 1:
        var x0 int = self__0._p0
        var t0 string
        var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(x0)
        t0 = inline0
        var t1 string = "State::Running(" + t0
        var t2 string = t1 + ")"
        return t2
    default:
        panic("non-exhaustive match")
    }
}

func _goml_m_trait__impl_i_Debug_i_Message_i_debug(self__0 Message) string {
    switch self__0.(type) {
    case Message_Data:
        var x0 string = self__0.(Message_Data)._0
        var x1 int = self__0.(Message_Data)._1
        var t0 string
        var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(x0)
        t0 = inline1
        var t1 string = "Message::Data(" + t0
        var t2 string = t1 + ", "
        var t3 string
        var inline0 string = _goml_m_trait__impl_i_ToString_i_isize_i_to__string(x1)
        t3 = inline0
        var t4 string = t2 + t3
        var t5 string = t4 + ")"
        return t5
    case Message_Empty:
        return "Message::Empty"
    default:
        panic("non-exhaustive match")
    }
}

func main0() struct{} {
    var settings__0 Settings = _goml_m_trait__impl_i_Default_i_Settings_i_default()
    var t0 string = _goml_m_trait__impl_i_Debug_i_Settings_i_debug(settings__0)
    println__T_string(t0)
    var t1 Settings = Settings{
        enabled: false,
        retries: 0,
        label: "",
    }
    var t2 bool = _goml_m_trait__impl_i_PartialEq_i_Settings_i_eq(settings__0, t1)
    var t3 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t2)
    println__T_string(t3)
    var empty__0 Empty = _goml_m_trait__impl_i_Default_i_Empty_i_default()
    var t4 string = _goml_m_trait__impl_i_Debug_i_Empty_i_debug(empty__0)
    println__T_string(t4)
    var pair__0 PairDefaults__isize = _goml_m_trait__impl_i_Default_i_PairDefaults____isize_i_default()
    var t5 int = pair__0.first
    var t6 int = pair__0.second
    var t7 int = t5 + t6
    var t8 string = _goml_m_inherent_i_isize_i_isize_i_to__string(t7)
    println__T_string(t8)
    var t9 Option__isize = pair__0.nested
    var t10 bool = _goml_m_inherent_i_Option_i_Option_l_T_r__i_is__none____T__isize(t9)
    var t11 string
    var inline30 string = _goml_runtime_core_bool_to_string(t10)
    t11 = inline30
    println__T_string(t11)
    var state__0 State
    state__0 = State{
        _tag: 0,
    }
    var t12 string = _goml_m_trait__impl_i_Debug_i_State_i_debug(state__0)
    var inline28 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t12)
    _goml_runtime_core_string_println(inline28)
    var message__0 Message
    var inline25 string = _goml_m_trait__impl_i_Default_i_string_i_default()
    var inline26 int = _goml_m_trait__impl_i_Default_i_isize_i_default()
    var inline27 Message = Message_Data{
        _0: inline25,
        _1: inline26,
    }
    message__0 = inline27
    var t13 string = _goml_m_trait__impl_i_Debug_i_Message_i_debug(message__0)
    var inline23 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t13)
    _goml_runtime_core_string_println(inline23)
    var event__0 Event
    var inline20 string = _goml_m_trait__impl_i_Default_i_string_i_default()
    var inline21 int = _goml_m_trait__impl_i_Default_i_isize_i_default()
    var inline22 Event = Event_Data{
        _0: inline20,
        _1: inline21,
    }
    event__0 = inline22
    var t14 string
    switch event__0.(type) {
    case Event_Data:
        var inline10 string = event__0.(Event_Data)._0
        var inline11 int = event__0.(Event_Data)._1
        var inline12 string = "Event::Data { " + "name: "
        var inline13 string = _goml_m_trait__impl_i_Debug_i_string_i_debug(inline10)
        var inline14 string = inline12 + inline13
        var inline15 string = inline14 + ", "
        var inline16 string = inline15 + "count: "
        var inline17 string = _goml_m_trait__impl_i_Debug_i_isize_i_debug(inline11)
        var inline18 string = inline16 + inline17
        var inline19 string = inline18 + " }"
        t14 = inline19
    case Event_Empty:
        t14 = "Event::Empty"
    default:
        panic("non-exhaustive match")
    }
    var inline8 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t14)
    _goml_runtime_core_string_println(inline8)
    var lazy__0 Lazy__NoDefault
    lazy__0 = Lazy__NoDefault{
        _tag: 0,
    }
    var jp0 string
    switch lazy__0._tag {
    case 0:
        jp0 = "empty"
    case 1:
        jp0 = "value"
    default:
        panic("non-exhaustive match")
    }
    var inline6 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(jp0)
    _goml_runtime_core_string_println(inline6)
    var selected__0 Selected__isize
    var inline3 int = _goml_m_trait__impl_i_Default_i_isize_i_default()
    var inline4 int = _goml_m_trait__impl_i_Default_i_isize_i_default()
    var inline5 Selected__isize = Values{
        _0: inline3,
        _1: inline4,
    }
    selected__0 = inline5
    var jp1 string
    switch selected__0.(type) {
    case Values:
        var x0 int = selected__0.(Values)._0
        var x1 int = selected__0.(Values)._1
        var t15 int = x0 + x1
        var inline2 string = __goml_builtin_int_to_string(t15)
        jp1 = inline2
    case Ignored:
        jp1 = "ignored"
    default:
        panic("non-exhaustive match")
    }
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(jp1)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_trait__impl_i_Default_i_isize_i_default() int {
    return 0
}

func _goml_m_trait__impl_i_Default_i_string_i_default() string {
    return ""
}

func _goml_m_trait__impl_i_Debug_i_isize_i_debug(self__0 int) string {
    var inline0 string = __goml_builtin_int_to_string(self__0)
    return inline0
}

func _goml_m_trait__impl_i_Debug_i_string_i_debug(self__0 string) string {
    return self__0
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

func _goml_m_trait__impl_i_Default_i_PairDefaults____isize_i_default() PairDefaults__isize {
    var t0 int
    t0 = 0
    var t1 int
    t1 = 0
    var t2 Option__isize
    t2 = Option__isize{
        _tag: 0,
    }
    var t3 PairDefaults__isize = PairDefaults__isize{
        first: t0,
        second: t1,
        nested: t2,
    }
    return t3
}

func _goml_m_inherent_i_isize_i_isize_i_to__string(self__0 int) string {
    var inline0 int64 = int64(int(self__0))
    var inline1 string = signed_decimal_string(inline0)
    return inline1
}

func _goml_m_inherent_i_Option_i_Option_l_T_r__i_is__none____T__isize(self__0 Option__isize) bool {
    var t0 bool
    switch self__0._tag {
    case 0:
        t0 = false
    case 1:
        t0 = true
    default:
        panic("non-exhaustive match")
    }
    var t1 bool = !t0
    return t1
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

func main() {
    main0()
}
