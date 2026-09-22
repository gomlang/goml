package native

import (
	"io"
	"strings"
)

func StringReaderRune(value string, offset int) (rune, int) {
	reader := strings.NewReader(value)
	if _, err := reader.Seek(int64(offset), io.SeekStart); err != nil {
		panic(err)
	}
	character, width, _ := reader.ReadRune()
	return character, width
}
