package main

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

func main0() int32 {
    var value__0 *ref_int32_x
    var inline3 int32 = 1
    var inline4 *ref_int32_x = ref__Ref_5int32(inline3)
    value__0 = inline4
    var inline1 int32 = 2
    ref_set__Ref_5int32(value__0, inline1)
    var inline0 int32 = ref_get__Ref_5int32(value__0)
    return inline0
}

func main() {
    main0()
}
