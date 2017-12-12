package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/composit/jumpc/pkg/encoder"
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

	h := encoder.Hash(pwd)
	b := encoder.Base64(h)
	if _, err = w.Write([]byte(b)); err != nil {
		w.WriteHeader(500)
		return
	}
}
