package native

import (
	"errors"
	"math"
	"math/big"
	"strconv"
	"strings"
)

func ParseFloat(value string, bits int) (uint64, bool, bool) {
	parsed, err := strconv.ParseFloat(value, bits)
	return math.Float64bits(parsed), err == nil, errors.Is(err, strconv.ErrRange)
}

func ParseExactFloat(value string, bits int) (uint64, bool, bool) {
	rational, valid := new(big.Rat).SetString(value)
	if !valid {
		return 0, false, false
	}
	var parsed float64
	if bits == 32 {
		narrow, _ := rational.Float32()
		parsed = float64(narrow)
	} else {
		parsed, _ = rational.Float64()
	}
	if parsed == 0 && strings.HasPrefix(value, "-") {
		parsed = math.Copysign(0, -1)
	}
	overflow := math.IsInf(parsed, 0)
	return math.Float64bits(parsed), !overflow, overflow
}
