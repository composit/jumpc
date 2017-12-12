package encode

import (
	"crypto/sha512"
	"encoding/base64"
)

// Hash takes a byte slice and hashes it using SHA512.
// It returns a byte slice.
func Hash(pwd []byte) []byte {
	hasher := sha512.New()
	hasher.Write(pwd)
	return hasher.Sum(nil)
}

// Base64 takes a byte slice and Base64 encodes it.
// It returns a byte slice.
func Base64(data []byte) []byte {
	enc := base64.StdEncoding
	buf := make([]byte, enc.EncodedLen(len(data)))
	enc.Encode(buf, data)
	return buf
}
