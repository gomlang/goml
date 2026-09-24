package native

import "strconv"

func Unquote(value string) ([]byte, bool) {
	decoded, err := strconv.Unquote(value)
	return []byte(decoded), err == nil
}
