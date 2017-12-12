package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/composit/jumpc/pkg/encode"
)

// Listen creates the handler to accept incoming http requests.
func Listen(port string) {
	http.HandleFunc("/", PwdHash)

	addr := fmt.Sprintf(":%s", port)
	log.Fatal(http.ListenAndServe(addr, nil))
}

// PwdHash handles incoming http requests for hashed passwords.
// The password is parsed out of the body, SHA512 hashed,
// Base64 encoded, and returned.
func PwdHash(w http.ResponseWriter, req *http.Request) {
	input, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	pwd, err := encode.GetPwd(input)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	h := encode.Hash(pwd)
	b := encode.Base64(h)
	if _, err = w.Write([]byte(b)); err != nil {
		w.WriteHeader(500)
		return
	}
}
