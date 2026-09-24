package native

import (
	"bytes"
	"slices"
	"strings"
	"unicode"
)

func ByteLines(input []byte) [][]byte {
	return slices.Collect(bytes.Lines(input))
}

func TextLines(input string) []string {
	return slices.Collect(strings.Lines(input))
}

func textSeparator(value rune) bool {
	return value == ',' || value == '界' || value == '🙂'
}

func TextFieldsBy(input string) []string {
	return strings.FieldsFunc(input, textSeparator)
}

func TextTrimBy(input string) string {
	return strings.TrimFunc(input, textSeparator)
}

func TextTrimStartBy(input string) string {
	return strings.TrimLeftFunc(input, textSeparator)
}

func TextTrimEndBy(input string) string {
	return strings.TrimRightFunc(input, textSeparator)
}

func TextIndexBy(input string) int {
	return strings.IndexFunc(input, textSeparator)
}

func TextLastIndexBy(input string) int {
	return strings.LastIndexFunc(input, textSeparator)
}

func byteSeparator(value rune) bool {
	return value == ',' || value == '界' || value == unicode.ReplacementChar
}

func ByteFieldsBy(input []byte) [][]byte  { return bytes.FieldsFunc(input, byteSeparator) }
func ByteTrimBy(input []byte) []byte      { return bytes.TrimFunc(input, byteSeparator) }
func ByteTrimStartBy(input []byte) []byte { return bytes.TrimLeftFunc(input, byteSeparator) }
func ByteTrimEndBy(input []byte) []byte   { return bytes.TrimRightFunc(input, byteSeparator) }
func ByteIndexBy(input []byte) int        { return bytes.IndexFunc(input, byteSeparator) }
func ByteLastIndexBy(input []byte) int    { return bytes.LastIndexFunc(input, byteSeparator) }
func ByteLastRune(input []byte, expected rune) int {
	return bytes.LastIndexFunc(input, func(value rune) bool { return value == expected })
}

func ByteMap(input []byte) []byte {
	return bytes.Map(func(value rune) rune {
		if value == ',' {
			return -1
		}
		if value == 'a' {
			return '界'
		}
		return unicode.ToUpper(value)
	}, input)
}

func ByteCase(input []byte, kind int, special bool) []byte {
	if special {
		switch kind {
		case 0:
			return bytes.ToUpperSpecial(unicode.TurkishCase, input)
		case 1:
			return bytes.ToLowerSpecial(unicode.TurkishCase, input)
		default:
			return bytes.ToTitleSpecial(unicode.TurkishCase, input)
		}
	}
	switch kind {
	case 0:
		return bytes.ToUpper(input)
	case 1:
		return bytes.ToLower(input)
	default:
		return bytes.ToTitle(input)
	}
}

func TextMap(input string) string {
	return strings.Map(func(value rune) rune {
		if value == ',' {
			return -1
		}
		if value == 'a' {
			return '界'
		}
		return unicode.ToUpper(value)
	}, input)
}

func TextCase(input string, kind int, special bool) string {
	if special {
		switch kind {
		case 0:
			return strings.ToUpperSpecial(unicode.TurkishCase, input)
		case 1:
			return strings.ToLowerSpecial(unicode.TurkishCase, input)
		default:
			return strings.ToTitleSpecial(unicode.TurkishCase, input)
		}
	}
	switch kind {
	case 0:
		return strings.ToUpper(input)
	case 1:
		return strings.ToLower(input)
	default:
		return strings.ToTitle(input)
	}
}
