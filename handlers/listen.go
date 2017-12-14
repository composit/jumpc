package handlers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/composit/jumpc/pkg/encode"
)

// HandlerChan holds the channel that the handlers can use to
// communicate when to stop the server.
type HandlerChan struct {
	C chan struct{}
}

// NewServer creates the server and handlers to accept incoming
// http requests.
func NewServer(port string, stop chan struct{}) (*http.Server, error) {
	if err := validatePort(port); err != nil {
		return nil, err
	}

	h := HandlerChan{
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

	return srv, nil
}

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

// validate port performs simple port validation
// it simply checks if the port is numerical
func validatePort(port string) error {
	if _, err := strconv.Atoi(port); err != nil {
		return fmt.Errorf("bad port specified: %s", port)
	}

	return nil
}
