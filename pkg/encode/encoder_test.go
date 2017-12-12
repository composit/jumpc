package encode_test

import (
	"fmt"
	"testing"

	"github.com/composit/jumpc/pkg/encode"
)

func TestHash(t *testing.T) {
	orig := []byte("test")
	hsh := encode.Hash(orig)

	// the base16 string representation of the hash
	expected := "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff"
	if actual := fmt.Sprintf("%x", hsh); expected != actual {
		t.Errorf("SHA512 hashes do not match: want %s, got %s\n", expected, actual)
	}
}

func TestBase64(t *testing.T) {
	orig := []byte("test")
	b64 := encode.Base64(orig)

	expected := "dGVzdA=="
	if actual := string(b64); expected != actual {
		t.Errorf("base64 strings do not match: want %s, got %s\n", expected, actual)
	}
}
