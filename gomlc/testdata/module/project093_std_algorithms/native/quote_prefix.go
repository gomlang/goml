package native

import "strconv"

func QuotedPrefix(value string) (string, bool) {
	prefix, err := strconv.QuotedPrefix(value)
	return prefix, err == nil
}
