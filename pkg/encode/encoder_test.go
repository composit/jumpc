package encode_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/composit/jumpc/pkg/encode"
)

func TestDo(t *testing.T) {
	orig := []byte("password=test")
	out, err := encode.Do(orig)
	if err != nil {
		t.Fatalf("failed to encode the password: %s\n", err)
	}

	// the base64 string representation of the hash
	expected := "7iaw3Ur350mqGo7jwQrpkj9hiYB3Lkc/iBml1JQODbJ6wYX4oOHV+E+IvIh/1nsUNzLDBMxfqa2Ob1f1ACio/w=="
	if actual := fmt.Sprintf("%s", out); expected != actual {
		t.Errorf("SHA512 hashes do not match: want %s, got %s\n", expected, actual)
	}
}

func TestDoBadInput(t *testing.T) {
	input := []byte("pwd=P@ssw0rd")

	_, err := encode.Do(input)
	if err == nil {
		t.Fatal(err)
	}

	msg := []byte(err.Error())
	expected := []byte("improperly formatted")
	if !bytes.Contains(msg, expected) {
		t.Errorf("unknown error message: `%s` does not contain `%s`.", msg, expected)
	}
}
