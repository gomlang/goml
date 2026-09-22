package native

import (
	"errors"
	"math"
	"strconv"
)

func ParseComplex(value string, width int) (uint64, uint64, bool, bool) {
	parsed, err := strconv.ParseComplex(value, width*2)
	return math.Float64bits(real(parsed)), math.Float64bits(imag(parsed)), err == nil, errors.Is(err, strconv.ErrRange)
}
