package native

import "math"

func IntegerFloatBits(value uint64, signed bool) (uint32, uint64) {
	if signed {
		return math.Float32bits(float32(int64(value))), math.Float64bits(float64(int64(value)))
	}
	return math.Float32bits(float32(value)), math.Float64bits(float64(value))
}

func RoundedInteger(value float64, mode, width int, signed bool) (uint64, int) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return 0, 1
	}
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
	if signed {
		limit := math.Ldexp(1, width-1)
		if value < -limit || value >= limit {
			return 0, 2
		}
		return uint64(int64(value)), 0
	}
	if value < 0 || value >= math.Ldexp(1, width) {
		return 0, 2
	}
	return uint64(value), 0
}
