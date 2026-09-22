package testoracle

import (
	"bytes"
	"compress/lzw"
	"encoding/hex"
	"io"
)

func LzwEncode(input string, order int, literalWidth int) string {
	data, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	var output bytes.Buffer
	writer := lzw.NewWriter(&output, lzw.Order(order), literalWidth)
	if _, err := writer.Write(data); err != nil {
		panic(err)
	}
	if err := writer.Close(); err != nil {
		panic(err)
	}
	return hex.EncodeToString(output.Bytes())
}

func LzwDecode(input string, order int, literalWidth int) string {
	data, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	reader := lzw.NewReader(bytes.NewReader(data), lzw.Order(order), literalWidth)
	decoded, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	if err := reader.Close(); err != nil {
		panic(err)
	}
	return hex.EncodeToString(decoded)
}
