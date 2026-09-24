package native

import (
	"crypto/pbkdf2"
	"crypto/sha256"
	"crypto/sha3"
	"crypto/sha512"
	"encoding/hex"
)

func Sum224(input []byte) string {
	value := sha3.Sum224(input)
	return hex.EncodeToString(value[:])
}

func Sum256(input []byte) string {
	value := sha3.Sum256(input)
	return hex.EncodeToString(value[:])
}

func Sum384(input []byte) string {
	value := sha3.Sum384(input)
	return hex.EncodeToString(value[:])
}

func Sum512(input []byte) string {
	value := sha3.Sum512(input)
	return hex.EncodeToString(value[:])
}

func Shake128(input []byte, length int) string {
	return hex.EncodeToString(sha3.SumSHAKE128(input, length))
}

func Shake256(input []byte, length int) string {
	return hex.EncodeToString(sha3.SumSHAKE256(input, length))
}

func PBKDF2SHA256(password, salt []byte, iterations, length int) string {
	value, err := pbkdf2.Key(sha256.New, string(password), salt, iterations, length)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(value)
}

func PBKDF2SHA512(password, salt []byte, iterations, length int) string {
	value, err := pbkdf2.Key(sha512.New, string(password), salt, iterations, length)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(value)
}

func PBKDF2SHA3_256(password, salt []byte, iterations, length int) string {
	value, err := pbkdf2.Key(sha3.New256, string(password), salt, iterations, length)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(value)
}
