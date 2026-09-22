package native

import "strconv"

func ParseSignedWidth(input string, radix, bits int) (int64, bool) {
	value, err := strconv.ParseInt(input, radix, bits)
	return value, err == nil
}

func ParseUnsignedWidth(input string, radix, bits int) (uint64, bool) {
	value, err := strconv.ParseUint(input, radix, bits)
	return value, err == nil
}
