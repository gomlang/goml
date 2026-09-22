package testoracle

import (
	"bytes"
	"compress/flate"
	"encoding/hex"
	"io"
)

func Encode(input string, dictionary string, level int) string {
	data, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	dict, err := hex.DecodeString(dictionary)
	if err != nil {
		panic(err)
	}
	var output bytes.Buffer
	writer, err := flate.NewWriterDict(&output, level, dict)
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

func Decode(input string, dictionary string) string {
	data, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	dict, err := hex.DecodeString(dictionary)
	if err != nil {
		panic(err)
	}
	reader := flate.NewReaderDict(bytes.NewReader(data), dict)
	decoded, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	if err := reader.Close(); err != nil {
		panic(err)
	}
	return hex.EncodeToString(decoded)
}
