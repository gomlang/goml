package main

import (
    _goml_os "os"
)

func _goml_runtime_core_string_println(s string) struct{} {
    _goml_os.Stdout.WriteString(s + "\n")
    return struct{}{}
}

type _goml_vec_uint32 struct {
    items []uint32
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

type Event uint8

const (
    Open Event = 0
    Close Event = 1
    Advance Event = 2
    Error Event = 3
)

func main0() struct{} {
    var running__0 *ref_bool_x
    var inline12 bool = true
    var inline13 *ref_bool_x = ref__Ref_4bool(inline12)
    running__0 = inline13
    Loop_loop0:
    for {
        var t0 bool
        var inline11 bool = ref_get__Ref_4bool(running__0)
        t0 = inline11
        if t0 {
            var event__0 Event = Open
            switch event__0 {
            case Open:
                var scanning__1 *ref_bool_x
                var inline9 bool = false
                var inline10 *ref_bool_x = ref__Ref_4bool(inline9)
                scanning__1 = inline10
                Loop_loop2:
                for {
                    var t2 bool
                    var inline8 bool = ref_get__Ref_4bool(scanning__1)
                    t2 = inline8
                    if t2 {
                        continue
                    } else {
                        break Loop_loop2
                    }
                }
                var scanning__0 *ref_bool_x
                var inline6 bool = false
                var inline7 *ref_bool_x = ref__Ref_4bool(inline6)
                scanning__0 = inline7
                Loop_loop1:
                for {
                    var t1 bool
                    var inline5 bool = ref_get__Ref_4bool(scanning__0)
                    t1 = inline5
                    if t1 {
                        continue
                    } else {
                        break Loop_loop1
                    }
                }
                var inline3 bool = false
                ref_set__Ref_4bool(running__0, inline3)
                continue
            case Close:
                var scanning__0 *ref_bool_x
                var inline6 bool = false
                var inline7 *ref_bool_x = ref__Ref_4bool(inline6)
                scanning__0 = inline7
                Loop_loop1__2:
                for {
                    var t1 bool
                    var inline5 bool = ref_get__Ref_4bool(scanning__0)
                    t1 = inline5
                    if t1 {
                        continue
                    } else {
                        break Loop_loop1__2
                    }
                }
                var inline3 bool = false
                ref_set__Ref_4bool(running__0, inline3)
                continue
            case Advance:
                var scanning__0 *ref_bool_x
                var inline6 bool = false
                var inline7 *ref_bool_x = ref__Ref_4bool(inline6)
                scanning__0 = inline7
                Loop_loop1__3:
                for {
                    var t1 bool
                    var inline5 bool = ref_get__Ref_4bool(scanning__0)
                    t1 = inline5
                    if t1 {
                        continue
                    } else {
                        break Loop_loop1__3
                    }
                }
                var inline3 bool = false
                ref_set__Ref_4bool(running__0, inline3)
                continue
            case Error:
                var scanning__0 *ref_bool_x
                var inline6 bool = false
                var inline7 *ref_bool_x = ref__Ref_4bool(inline6)
                scanning__0 = inline7
                Loop_loop1__4:
                for {
                    var t1 bool
                    var inline5 bool = ref_get__Ref_4bool(scanning__0)
                    t1 = inline5
                    if t1 {
                        continue
                    } else {
                        break Loop_loop1__4
                    }
                }
                var inline3 bool = false
                ref_set__Ref_4bool(running__0, inline3)
                continue
            default:
                panic("non-exhaustive match")
            }
        } else {
            break Loop_loop0
        }
    }
    var inline0 string = "ok"
    var inline1 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(inline0)
    _goml_runtime_core_string_println(inline1)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func main() {
    main0()
}
