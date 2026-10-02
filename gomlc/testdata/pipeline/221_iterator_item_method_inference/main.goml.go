package main

type _goml_vec_uint8 struct {
    items []uint8
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

type Values struct {}

type Ordering uint8

type Option__i32 uint64

func main0() struct{} {
    return struct{}{}
}

func main() {
    main0()
}
