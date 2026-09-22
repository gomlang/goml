package native

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"
)

func JSONTokens(input string) (bool, string) {
	if !json.Valid([]byte(input)) {
		return false, ""
	}
	decoder := json.NewDecoder(strings.NewReader(input))
	decoder.UseNumber()
	var output strings.Builder
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			return true, output.String()
		}
		if err != nil {
			return false, ""
		}
		switch value := token.(type) {
		case json.Delim:
			output.WriteRune(rune(value))
		case string:
			output.WriteString("s" + strconv.Itoa(len(value)) + ":" + value + ";")
		case json.Number:
			output.WriteString("n" + string(value) + ";")
		case bool:
			output.WriteString(strconv.FormatBool(value) + ";")
		case nil:
			output.WriteString("null;")
		}
	}
}
