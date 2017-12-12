package encode_test

import (
	"bytes"
	"testing"

	"github.com/composit/jumpc/pkg/encode"
)

func TestParseInput(t *testing.T) {
	input := []byte("password=P@ssw0rd")

	expected := []byte("P@ssw0rd")
	actual, err := encode.GetPwd(input)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(actual, expected) {
		t.Errorf("passwords do not match: want %s, got %s\n", expected, actual)
	}
}

func TestParseBadInput(t *testing.T) {
	input := []byte("pwd=P@ssw0rd")

	_, err := encode.GetPwd(input)
	if err == nil {
		t.Fatal(err)
	}

	msg := []byte(err.Error())
	expected := []byte("improperly formatted")
	if !bytes.Contains(msg, expected) {
		t.Errorf("unknown error message: `%s` does not contain `%s`.", msg, expected)
	}
}
