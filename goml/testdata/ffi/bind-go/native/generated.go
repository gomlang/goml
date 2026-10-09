package native

import (
    native_0 "strconv"
    native_1 "strings"
)

func GomlBind_count(p0 string, p1 string) int {
    return native_1.Count(p0, p1)
}

func GomlBind_parse(p0 string) (int, error) {
    return native_0.Atoi(p0)
}
