package encode_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/composit/jumpc/pkg/encode"
)

func TestParseInput(t *testing.T) {
	input := "password=P@ssw0rd"

	expected := "P@ssw0rd"
	actual, err := encode.GetPwd(input)
	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Errorf("passwords do not match: want %s, got %s\n", expected, actual)
	}
}

func TestParseBadInput(t *testing.T) {
	input := "pwd=P@ssw0rd"

	_, err := encode.GetPwd(input)
	if err == nil {
		t.Fatal(err)
	}

	msg := fmt.Sprintf("%s", err)
	expected := "improperly formatted"
	if !strings.Contains(msg, expected) {
		t.Errorf("unknown error message: `%s` does not contain `%s`.", msg, expected)
	}
}
