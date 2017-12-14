package encode

import (
	"crypto/sha512"
	"encoding/base64"
)

// Do takes a string, parses out the password, hashes it and returns
// a base64 string of the hash
func Do(input []byte) ([]byte, error) {
	pwd, err := getPwd(input)
	if err != nil {
		return []byte{}, err
	}

	hsh := hash(pwd)
	return b64(hsh), nil
}

func hash(pwd []byte) []byte {
	hasher := sha512.New()
	hasher.Write(pwd)
	return hasher.Sum(nil)
}

func b64(data []byte) []byte {
	enc := base64.StdEncoding
	buf := make([]byte, enc.EncodedLen(len(data)))
	enc.Encode(buf, data)
	return buf
}
