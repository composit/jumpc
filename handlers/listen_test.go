package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/composit/jumpc/handlers"
)

func TestListen(t *testing.T) {
	req, err := http.NewRequest("POST", "/", strings.NewReader("password=hunter2"))
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	h := handlers.HandlerChan{}
	handler := http.HandlerFunc(h.PwdHash)

	handler.ServeHTTP(rec, req)

	if c := rec.Code; c != http.StatusOK {
		t.Fatalf("server returned a bad response code: %d\n", c)
	}

	expected := "a5ftaNFOs/GqlZzl1Jx9xhLh6x2v1zsecFhHSD/WpsgJ8s606N9v+ZhMYpj/AoXKzmYUv42qnwBwEBtsiYmeIg=="
	if actual := rec.Body.String(); expected != actual {
		t.Errorf("server returned a bad hash: want %s, got %s\n", expected, actual)
	}
}

func TestNewServerBadPort(t *testing.T) {
	c := make(chan struct{})
	badPort := "abc"

	_, err := handlers.NewServer(badPort, c)
	if err == nil {
		t.Fatalf("server returned with a bad port: %s\n", badPort)
	}

	expected := "bad port"
	if actual := err.Error(); !strings.Contains(actual, expected) {
		t.Errorf("bad error message. `%s` does not contain `%s`", actual, expected)
	}
}
