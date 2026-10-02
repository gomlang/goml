package main

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

type Point struct {
    x int32
    y int32
}

type Ordering uint8

func main0() int32 {
    var p__0 Point
    var inline1 int32 = 5
    var inline2 int32 = inline1 + 1
    var inline3 Point = Point{
        x: inline1,
        y: inline2,
    }
    p__0 = inline3
    var t0 int32 = p__0.x
    var t1 int32 = t0 + 1
    var t2 int32 = p__0.y
    var t3 int32 = t2 - 2
    var t4 int32
    var inline0 int32 = t1 + t3
    t4 = inline0
    var t5 int32 = t1 + t4
    return t5
}

func main() {
    main0()
}
