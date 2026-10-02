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

type Worker struct {
    name string
}

type Station struct {
    worker Worker
}

type Ordering uint8

func _goml_m_trait__impl_i_Ready_i_Worker_i_state(self__0 Worker) string {
    var t0 string = self__0.name
    var t1 string = t0 + ":ready"
    return t1
}

func _goml_m_trait__impl_i_Service_i__l_Worker_r__x40_Station_i_get(self__0 Station) Worker {
    var t0 Worker = self__0.worker
    return t0
}

func main0() struct{} {
    var t0 Worker = Worker{
        name: "build",
    }
    var t1 Station = Station{
        worker: t0,
    }
    var t2 string
    var inline2 Worker = _goml_m_trait__impl_i_Service_i__l_Worker_r__x40_Station_i_get(t1)
    var inline3 string = _goml_m_trait__impl_i_Ready_i_Worker_i_state(inline2)
    t2 = inline3
    var inline0 string = _goml_m_trait__impl_i_ToString_i_string_i_to__string(t2)
    _goml_runtime_core_string_println(inline0)
    return struct{}{}
}

func _goml_m_trait__impl_i_ToString_i_string_i_to__string(self__0 string) string {
    return self__0
}

func main() {
    main0()
}
