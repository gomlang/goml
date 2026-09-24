package native

import "strconv"

func UnquoteElement(value string, delimiter byte) (uint32, bool, string, bool) {
	code, scalar, tail, err := strconv.UnquoteChar(value, delimiter)
	return uint32(code), scalar, tail, err == nil
}
