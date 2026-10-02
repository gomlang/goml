package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Print(len(slices.Grow([]int{}, 1)))
}
