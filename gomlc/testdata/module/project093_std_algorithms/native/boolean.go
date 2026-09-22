package native

import "strconv"

func ParseBool(value string) (bool, bool) {
	parsed, err := strconv.ParseBool(value)
	return parsed, err == nil
}
