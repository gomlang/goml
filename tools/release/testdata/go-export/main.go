package main

import (
	"example.com/host/gen/calclib"
	"fmt"
)

func main() {
	a, b := calclib.Pair(7)
	fmt.Println(calclib.Add(20, 22), calclib.AbsPlusOne(-3), a, b)
}
