package native

import (
	"crypto/hkdf"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func Hmac256(key, input []byte) string {
	h := hmac.New(sha256.New, key)
	h.Write(input)
	return hex.EncodeToString(h.Sum(nil))
}

func Hmac512(key, input []byte) string {
	h := hmac.New(sha512.New, key)
	h.Write(input)
	return hex.EncodeToString(h.Sum(nil))
}

func Extract256(salt, inputKey []byte) string {
	value, err := hkdf.Extract(sha256.New, inputKey, salt)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(value)
}

func Expand256(prk, info []byte, length int) string {
	value, err := hkdf.Expand(sha256.New, prk, string(info), length)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(value)
}

func Derive256(salt, inputKey, info []byte, length int) string {
	value, err := hkdf.Key(sha256.New, inputKey, salt, string(info), length)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(value)
}

func Derive512(salt, inputKey, info []byte, length int) string {
	value, err := hkdf.Key(sha512.New, inputKey, salt, string(info), length)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(value)
}
