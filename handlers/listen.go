package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/composit/jumpc/pkg/hasher"
)

func Listen(port string) {
	http.HandleFunc("/", PwdHash)

	addr := fmt.Sprintf(":%s", port)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func PwdHash(w http.ResponseWriter, req *http.Request) {
	pwd, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	h := hasher.Hash(pwd)
	if _, err = w.Write([]byte(h)); err != nil {
		w.WriteHeader(500)
		return
	}
}
