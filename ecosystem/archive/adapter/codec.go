package adapter

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"fmt"
	"io"
	"math"
	"os"
)

func Compress(data []byte, gzipFormat bool) ([]byte, string) {
	var out bytes.Buffer
	var writer io.WriteCloser
	if gzipFormat {
		writer = gzip.NewWriter(&out)
	} else {
		var err error
		writer, err = flate.NewWriter(&out, flate.DefaultCompression)
		if err != nil {
			return nil, err.Error()
		}
	}
	if _, err := writer.Write(data); err != nil {
		return nil, err.Error()
	}
	if err := writer.Close(); err != nil {
		return nil, err.Error()
	}
	return out.Bytes(), ""
}

func Decompress(data []byte, gzipFormat bool, limit int) ([]byte, string) {
	if limit < 0 {
		return nil, "negative decompression limit"
	}
	source := bytes.NewReader(data)
	var reader io.ReadCloser
	if gzipFormat {
		var err error
		reader, err = gzip.NewReader(source)
		if err != nil {
			return nil, err.Error()
		}
	} else {
		reader = flate.NewReader(source)
	}
	defer reader.Close()
	budget := int64(limit)
	if budget < math.MaxInt64 {
		budget++
	}
	out, err := io.ReadAll(io.LimitReader(reader, budget))
	if err != nil {
		return nil, err.Error()
	}
	if len(out) > limit {
		return nil, "decompressed data exceeds limit"
	}
	if source.Len() != 0 {
		return nil, "trailing compressed data"
	}
	return out, ""
}

type Root interface{ ArchiveRoot() }
type root struct{ value *os.Root }

func (*root) ArchiveRoot() {}

func OpenRoot(path string) (Root, string) {
	value, err := os.OpenRoot(path)
	if err != nil {
		return nil, err.Error()
	}
	return &root{value}, ""
}

func CloseRoot(value Root) { _ = value.(*root).value.Close() }

func Mkdir(value Root, path string) string {
	base := value.(*root).value
	if err := base.Mkdir(path, 0755); err == nil {
		return ""
	} else if !os.IsExist(err) {
		return err.Error()
	}
	metadata, err := base.Lstat(path)
	if err != nil {
		return err.Error()
	}
	if !metadata.IsDir() {
		return "existing parent is not a directory"
	}
	return ""
}

func WriteFile(value Root, path string, data []byte, mode uint32) string {
	file, err := value.(*root).value.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, os.FileMode(mode&0777))
	if err != nil {
		return err.Error()
	}
	count, writeErr := file.Write(data)
	closeErr := file.Close()
	if writeErr != nil {
		return writeErr.Error()
	}
	if closeErr != nil {
		return closeErr.Error()
	}
	if count != len(data) {
		return fmt.Sprint(io.ErrShortWrite)
	}
	return ""
}
