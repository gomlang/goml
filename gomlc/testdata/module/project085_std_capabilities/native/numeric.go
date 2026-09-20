package native

import "math"

func IntegerFloats(value uint64) (uint64, uint32, uint64, uint32) {
	return math.Float64bits(float64(value)), math.Float32bits(float32(value)), math.Float64bits(float64(int64(value))), math.Float32bits(float32(int64(value)))
}

func FloatIntegers(bits uint64, mode int) (bool, int64, bool, uint64) {
	value := math.Float64frombits(bits)
	switch mode {
	case 0:
		value = math.Trunc(value)
	case 1:
		value = math.Floor(value)
	case 2:
		value = math.Ceil(value)
	case 3:
		value = math.Round(value)
	case 4:
		value = math.RoundToEven(value)
	}
	signed := !math.IsNaN(value) && value >= -0x1p63 && value < 0x1p63
	unsigned := !math.IsNaN(value) && value >= 0 && value < 0x1p64
	var i int64
	var u uint64
	if signed {
		i = int64(value)
	}
	if unsigned {
		u = uint64(value)
	}
	return signed, i, unsigned, u
}
