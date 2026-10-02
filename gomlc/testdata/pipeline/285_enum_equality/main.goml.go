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

func _goml_runtime_core_string_println(s string) struct{} {
    _goml_os.Stdout.WriteString(s + "\n")
    return struct{}{}
}

type _goml_vec_E struct {
    items []E
}

func vec_get__Vec_1E(vec *_goml_vec_E, index int) E {
    return vec.items[index]
}

func vec_len__Vec_1E(vec *_goml_vec_E) int {
    return int(len(vec.items))
}

type _goml_vec_Fn2_1E_1E_to_4bool struct {
    items []func(E, E) bool
}

func vec_get__Vec_18Fn2_1E_1E_to_4bool(vec *_goml_vec_Fn2_1E_1E_to_4bool, index int) func(E, E) bool {
    return vec.items[index]
}

func vec_len__Vec_18Fn2_1E_1E_to_4bool(vec *_goml_vec_Fn2_1E_1E_to_4bool) int {
    return int(len(vec.items))
}

type _goml_vec_uint32 struct {
    items []uint32
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

type E uint8

const (
    A E = 0
    B E = 1
    C E = 2
)

func same(a__0 E, b__0 E) bool {
    if a__0 <= C {
        return b__0 == a__0
    } else {
        panic("non-exhaustive match")
    }
}

func same_reversed(a__0 E, b__0 E) bool {
    if b__0 <= C {
        return a__0 == b__0
    } else {
        panic("non-exhaustive match")
    }
}

func main0() struct{} {
    var t0 [3]E = [3]E{A, B, C}
    var values__0 *_goml_vec_E = func(values [3]E) *_goml_vec_E {
        var storage struct {
            vector _goml_vec_E
            values [3]E
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t0)
    var t1 [2]func(E, E) bool = [2]func(E, E) bool{same, same_reversed}
    var functions__0 *_goml_vec_Fn2_1E_1E_to_4bool = func(values [2]func(E, E) bool) *_goml_vec_Fn2_1E_1E_to_4bool {
        var storage struct {
            vector _goml_vec_Fn2_1E_1E_to_4bool
            values [2]func(E, E) bool
        }
        storage.values = values
        storage.vector.items = storage.values[0:len(storage.values)]
        return &storage.vector
    }(t1)
    var for_limit0 int = vec_len__Vec_1E(values__0)
    var for_index0 int = 0
    Loop_loop0:
    for {
        var t2 bool = for_index0 < for_limit0
        if t2 {
            var for_item0 E = vec_get__Vec_1E(values__0, for_index0)
            var t3 int = for_index0 + 1
            for_index0 = t3
            var for_limit1 int = vec_len__Vec_1E(values__0)
            var for_index1 int = 0
            Loop_loop1:
            for {
                var t4 bool = for_index1 < for_limit1
                if t4 {
                    var for_item1 E = vec_get__Vec_1E(values__0, for_index1)
                    var t5 int = for_index1 + 1
                    for_index1 = t5
                    var for_limit2 int = vec_len__Vec_18Fn2_1E_1E_to_4bool(functions__0)
                    var for_index2 int = 0
                    Loop_loop2:
                    for {
                        var t6 bool = for_index2 < for_limit2
                        if t6 {
                            var for_item2 func(E, E) bool = vec_get__Vec_18Fn2_1E_1E_to_4bool(functions__0, for_index2)
                            var t7 int = for_index2 + 1
                            for_index2 = t7
                            var t8 bool = for_item2(for_item0, for_item1)
                            var inline0 string = _goml_m_trait__impl_i_ToString_i_bool_i_to__string(t8)
                            _goml_runtime_core_string_println(inline0)
                            continue
                        } else {
                            break Loop_loop2
                        }
                    }
                    continue
                } else {
                    break Loop_loop1
                }
            }
            continue
        } else {
            break Loop_loop0
        }
    }
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_bool_i_to__string(self__0 bool) string {
    var t0 string = _goml_runtime_core_bool_to_string(self__0)
    return t0
}

func main() {
    main0()
}
