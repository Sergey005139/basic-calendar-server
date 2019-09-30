package utils

import "encoding/base64"

func Encode(b []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(b))
}

func Decode(b []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(b))
}