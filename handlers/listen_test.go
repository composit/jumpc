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
	handler := http.HandlerFunc(handlers.PwdHash)

	handler.ServeHTTP(rec, req)

	if c := rec.Code; c != http.StatusOK {
		t.Fatalf("server returned a bad response code: %d\n", c)
	}

	expected := "MD-yGkEFDmjAFhaO0iyTQiVPs7JgThHTgB-MwX5AyYrpoA-6oHcfQqS2nCGAmyEoqblHGjp_i7VRhyRJIXGyHw=="
	if actual := rec.Body.String(); expected != actual {
		t.Errorf("server returned a bad hash: want %s, got %s\n", expected, actual)
	}
}
