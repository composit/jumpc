package server

import (
	"fmt"
	"log"
	"net/http"
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
