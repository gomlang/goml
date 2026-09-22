package native

import "encoding/pem"

func EncodePEM(label string, keys, values []string, data []byte) string {
	headers := make(map[string]string)
	for i, key := range keys {
		headers[key] = values[i]
	}
	return string(pem.EncodeToMemory(&pem.Block{Type: label, Headers: headers, Bytes: data}))
}

func DecodePEM(input string) (string, []byte, int, bool) {
	block, rest := pem.Decode([]byte(input))
	if block == nil {
		return "", nil, 0, false
	}
	return block.Type, block.Bytes, len(input) - len(rest), true
}
