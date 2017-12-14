package server_test

import (
	"strings"
	"testing"

	"github.com/composit/jumpc/pkg/server"
)

func TestNewServerBadPort(t *testing.T) {
	c := make(chan struct{})
	badPort := "abc"

	_, err := server.NewServer(badPort, c)
	if err == nil {
		t.Fatalf("server returned with a bad port: %s\n", badPort)
	}

	expected := "bad port"
	if actual := err.Error(); !strings.Contains(actual, expected) {
		t.Errorf("bad error message. `%s` does not contain `%s`", actual, expected)
	}
}
