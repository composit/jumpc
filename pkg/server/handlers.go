package server

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/composit/jumpc/pkg/encode"
)

// PwdHash handles incoming http requests for hashed passwords.
// The password is parsed out of the body, SHA512 hashed,
// Base64 encoded, and returned.
func (h *HandlerChan) PwdHash(w http.ResponseWriter, req *http.Request) {
	timez := time.NewTimer(5 * time.Second)

	input, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	if bytes.Equal(input, []byte("graceful shutdown")) {
		close(h.C)
		if _, err := w.Write([]byte("ok")); err != nil {
			log.Println("failed to respond to graceful shutdown request")
		}
		return
	}

	b, err := encode.Do(input)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	<-timez.C

	if _, err = w.Write(b); err != nil {
		w.WriteHeader(500)
		return
	}
}
