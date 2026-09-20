package adapter

import (
	"fmt"
	"strings"
)

func tokens(text string) ([]string, error) {
	result := make([]string, 0)
	for i := 0; i < len(text); {
		b := text[i]
		if b == ' ' || b == '\t' || b == '\r' || b == '\n' || b == '\f' {
			i++
			continue
		}
		if i+1 < len(text) && text[i:i+2] == "--" {
			for i < len(text) && text[i] != '\n' {
				i++
			}
			continue
		}
		if i+1 < len(text) && text[i:i+2] == "/*" {
			end := strings.Index(text[i+2:], "*/")
			if end < 0 {
				return nil, fmt.Errorf("%w: unterminated SQL comment", ErrArgument)
			}
			i += end + 4
			continue
		}
		if b == '\'' || b == '"' || b == '`' || b == '[' {
			end := b
			if b == '[' {
				end = ']'
			}
			i++
			closed := false
			for i < len(text) {
				if text[i] == end {
					i++
					if end != ']' && i < len(text) && text[i] == end {
						i++
						continue
					}
					closed = true
					break
				}
				i++
			}
			if !closed {
				return nil, fmt.Errorf("%w: unterminated SQL quote", ErrArgument)
			}
			result = append(result, "?")
			continue
		}
		if b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' || b == '_' || b >= 128 {
			start := i
			i++
			for i < len(text) {
				b = text[i]
				if !(b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' || b >= '0' && b <= '9' || b == '_' || b >= 128) {
					break
				}
				i++
			}
			result = append(result, strings.ToUpper(text[start:i]))
			continue
		}
		result = append(result, string(b))
		i++
	}
	return result, nil
}
func singleStatement(text string) error {
	values, err := tokens(text)
	if err != nil {
		return err
	}
	for len(values) > 0 && values[0] == ";" {
		values = values[1:]
	}
	if len(values) == 0 {
		return fmt.Errorf("%w: empty SQL statement", ErrArgument)
	}
	switch values[0] {
	case "BEGIN", "COMMIT", "END", "ROLLBACK", "SAVEPOINT", "RELEASE":
		return fmt.Errorf("%w: use the transaction API for transaction control", ErrArgument)
	}
	trigger := false
	if values[0] == "CREATE" {
		position := 1
		if position < len(values) && (values[position] == "TEMP" || values[position] == "TEMPORARY") {
			position++
		}
		trigger = position < len(values) && values[position] == "TRIGGER"
	}
	body, finished, cases := false, false, 0
	for _, value := range values {
		if finished {
			if value != ";" {
				return fmt.Errorf("%w: expected one SQL statement", ErrArgument)
			}
			continue
		}
		if trigger && value == "BEGIN" && !body {
			body = true
			continue
		}
		if body {
			if value == "CASE" {
				cases++
			}
			if value == "END" {
				if cases > 0 {
					cases--
				} else {
					body = false
					trigger = false
				}
			}
		}
		if value == ";" && !body {
			finished = true
		}
	}
	return nil
}
