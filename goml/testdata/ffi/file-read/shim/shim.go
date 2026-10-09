package shim

import (
	"io"
	"os"
)

func EchoFile(file *os.File) *os.File {
	return file
}

func EchoRead(count int, failure error) (int, error) {
	return count, failure
}

func IsPartialRead(count int, failure error) bool {
	return count == 3 && failure == io.ErrUnexpectedEOF
}
