package testoracle

import (
	"bytes"
	"compress/gzip"
	"encoding/hex"
	"io"
)

func GzipEncode(input string, level int) string {
	data, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	var output bytes.Buffer
	writer, err := gzip.NewWriterLevel(&output, level)
	if err != nil {
		panic(err)
	}
	if _, err := writer.Write(data); err != nil {
		panic(err)
	}
	if err := writer.Close(); err != nil {
		panic(err)
	}
	return hex.EncodeToString(output.Bytes())
}

func GzipDecode(input string) string {
	data, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	decoded, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	if err := reader.Close(); err != nil {
		panic(err)
	}
	return hex.EncodeToString(decoded)
}
