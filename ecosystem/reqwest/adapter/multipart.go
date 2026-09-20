package adapter

import (
	"bytes"
	"errors"
	"mime"
	"mime/multipart"
	"net/textproto"
	"strings"
)

type Part interface{ part() }
type partState struct {
	name, filename, contentType string
	body                        []byte
	file                        bool
}

func (*partState) part() {}
func NewPart(name, filename, contentType string, body []byte, file bool) Part {
	return &partState{name, filename, contentType, bytes.Clone(body), file}
}

var errMultipartLimit = errors.New("multipart body exceeds request limit")

type limitedBuffer struct {
	bytes.Buffer
	limit int64
}

func (b *limitedBuffer) Write(value []byte) (int, error) {
	if int64(len(value)) > b.limit-int64(b.Len()) {
		return 0, errMultipartLimit
	}
	return b.Buffer.Write(value)
}
func multipartFailure(err error) Failure {
	if errors.Is(err, errMultipartLimit) {
		return fail(5, err.Error())
	}
	return failure(err)
}
func Multipart(parts []Part, limit int64) ([]byte, string, Failure) {
	buffer := limitedBuffer{limit: limit}
	writer := multipart.NewWriter(&buffer)
	for _, handle := range parts {
		p := handle.(*partState)
		if p.name == "" || strings.ContainsAny(p.name+p.filename+p.contentType, "\r\n\x00") {
			return nil, "", fail(1, "invalid multipart field")
		}
		header := textproto.MIMEHeader{}
		parameters := map[string]string{"name": p.name}
		if p.file {
			parameters["filename"] = p.filename
		}
		disposition := mime.FormatMediaType("form-data", parameters)
		if disposition == "" {
			return nil, "", fail(1, "invalid multipart disposition")
		}
		header.Set("Content-Disposition", disposition)
		if p.contentType != "" {
			if _, _, err := mime.ParseMediaType(p.contentType); err != nil {
				return nil, "", fail(1, "invalid multipart content type")
			}
			header.Set("Content-Type", p.contentType)
		}
		field, err := writer.CreatePart(header)
		if err != nil {
			return nil, "", multipartFailure(err)
		}
		if _, err = field.Write(p.body); err != nil {
			return nil, "", multipartFailure(err)
		}
	}
	if err := writer.Close(); err != nil {
		return nil, "", multipartFailure(err)
	}
	return buffer.Bytes(), writer.FormDataContentType(), nil
}
