package main

import (
	"example.com/host/gen/calclib"
	"testing"
)

func TestLibrary(t *testing.T) {
	var add func(int64, int64) int64 = calclib.Add
	var pair func(int64) (int64, int64) = calclib.Pair
	if add(20, 22) != 42 || calclib.AbsPlusOne(-3) != 4 {
		t.Fatal("GoML calls and imported math.Abs returned incorrect results")
	}
	a, b := pair(7)
	if a != 7 || b != 8 {
		t.Fatalf("Pair(7) = (%d, %d)", a, b)
	}
}
