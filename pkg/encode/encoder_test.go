package encode_test

import (
	"fmt"
	"testing"

	"github.com/composit/jumpc/pkg/encoder"
)

func TestHash(t *testing.T) {
	orig := []byte("test")
	hsh := encoder.Hash(orig)

	// the base16 string representation of the hash
	expected := "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff"
	if actual := fmt.Sprintf("%x", hsh); expected != actual {
		t.Errorf("encoder returned a bad hash: want %s, got %s\n", expected, actual)
	}
}

func TestBase64(t *testing.T) {
	orig := []byte("test")
	b64 := encoder.Base64(orig)

	expected := "dGVzdA=="
	if actual := string(b64); expected != actual {
		t.Errorf("encoder returned a bad base64 string: want %s, got %s\n", expected, actual)
	}
}
