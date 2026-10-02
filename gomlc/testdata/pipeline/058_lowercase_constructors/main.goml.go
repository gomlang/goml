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

type Maybe__i32 uint64

func main0() int32 {
    var commute_field0 int32
    var inline5 int32 = 5
    commute_field0 = inline5
    var pt__0 Point
    var inline3 int32 = 7
    var inline4 Point = Point{
        x: commute_field0,
        y: inline3,
    }
    pt__0 = inline4
    var t0 int32 = pt__0.x
    var t1 int32
    var inline0 int32 = pt__0.x
    var inline1 int32 = pt__0.y
    var inline2 int32 = inline0 + inline1
    t1 = inline2
    var t2 int32 = t0 + t1
    return t2
}

func main() {
    main0()
}
