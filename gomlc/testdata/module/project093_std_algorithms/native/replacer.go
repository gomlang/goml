package native

import (
	"strings"
	"unicode/utf8"
)

func ReplaceMany(value string, old, replacement []string) string {
	pairs := make([]string, 0, len(old)*2)
	for index := range old {
		pairs = append(pairs, old[index], replacement[index])
	}
	return strings.NewReplacer(pairs...).Replace(value)
}

func ReplaceScalarMany(value string, old, replacement []string) string {
	var output strings.Builder
	skipEmpty := false
	for offset := 0; ; {
		matched := -1
		for index, pattern := range old {
			if !(skipEmpty && pattern == "") && strings.HasPrefix(value[offset:], pattern) {
				matched = index
				break
			}
		}
		if matched >= 0 {
			output.WriteString(replacement[matched])
			skipEmpty = old[matched] == ""
			offset += len(old[matched])
			continue
		}
		if offset == len(value) {
			break
		}
		_, width := utf8.DecodeRuneInString(value[offset:])
		output.WriteString(value[offset : offset+width])
		offset += width
		skipEmpty = false
	}
	return output.String()
}
