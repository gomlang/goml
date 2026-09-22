package testoracle

import (
	"bytes"
	"compress/zlib"
	"encoding/hex"
	"io"
)

func ZlibEncode(input string, dictionary string, level int) string {
	data, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	dict, err := hex.DecodeString(dictionary)
	if err != nil {
		panic(err)
	}
	var output bytes.Buffer
	var writer *zlib.Writer
	if dictionary == "" {
		writer, err = zlib.NewWriterLevel(&output, level)
	} else {
		writer, err = zlib.NewWriterLevelDict(&output, level, dict)
	}
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

func ZlibDecode(input string, dictionary string) string {
	data, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	dict, err := hex.DecodeString(dictionary)
	if err != nil {
		panic(err)
	}
	var reader io.ReadCloser
	if dictionary == "" {
		reader, err = zlib.NewReader(bytes.NewReader(data))
	} else {
		reader, err = zlib.NewReaderDict(bytes.NewReader(data), dict)
	}
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
