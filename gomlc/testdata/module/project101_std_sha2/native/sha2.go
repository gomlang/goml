package native

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func Sum224(value []byte) string {
	sum := sha256.Sum224(value)
	return hex.EncodeToString(sum[:])
}

func Sum256(value []byte) string {
	sum := sha256.Sum256(value)
	return hex.EncodeToString(sum[:])
}

func Sum384(value []byte) string {
	sum := sha512.Sum384(value)
	return hex.EncodeToString(sum[:])
}

func Sum512(value []byte) string {
	sum := sha512.Sum512(value)
	return hex.EncodeToString(sum[:])
}

func Sum512_224(value []byte) string {
	sum := sha512.Sum512_224(value)
	return hex.EncodeToString(sum[:])
}

func Sum512_256(value []byte) string {
	sum := sha512.Sum512_256(value)
	return hex.EncodeToString(sum[:])
}
