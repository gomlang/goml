package native

import (
	"bufio"
	"bytes"
)

func ScanTokens(input []byte, kind string) [][]byte {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	switch kind {
	case "runes":
		scanner.Split(bufio.ScanRunes)
	case "words":
		scanner.Split(bufio.ScanWords)
	}
	var result [][]byte
	for scanner.Scan() {
		result = append(result, append([]byte(nil), scanner.Bytes()...))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return result
}
