package handlers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/composit/jumpc/pkg/encode"
)

type handler struct {
	C chan struct{}
}

// Server creates the handler to accept incoming http requests.
func Server(port string, stop chan struct{}) *http.Server {
	h := handler{
		C: stop,
	}
	http.HandleFunc("/", h.PwdHash)

	addr := fmt.Sprintf(":%s", port)
	srv := &http.Server{Addr: addr}

	go func(srv *http.Server) {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("%s\n", err)
		}
	}(srv)

	return srv
}

// PwdHash handles incoming http requests for hashed passwords.
// The password is parsed out of the body, SHA512 hashed,
// Base64 encoded, and returned.
func (h *handler) PwdHash(w http.ResponseWriter, req *http.Request) {
	timez := time.NewTimer(5 * time.Second)

	input, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	if bytes.Equal(input, []byte("graceful shutdown")) {
		close(h.C)
		w.Write([]byte("ok"))
		return
	}

	pwd, err := encode.GetPwd(input)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	hsh := encode.Hash(pwd)
	b := encode.Base64(hsh)

	<-timez.C

	if _, err = w.Write(b); err != nil {
		w.WriteHeader(500)
		return
	}
}
