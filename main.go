package main

import (
	"context"
	"log"
	"os"

	"github.com/composit/jumpc/pkg/server"
)

func main() {
	stop := make(chan struct{})

	if len(os.Args) != 2 {
		log.Fatalf("please enter one (and only one) port")
	}

	srv, err := server.NewServer(os.Args[1], stop)
	if err != nil {
		log.Fatalf("failed to initialize the server: %s", err)
	}

	<-stop

	log.Println("waiting for requests to finish")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("failed to shutdown properly: %s", err)
	}
	log.Println("exiting")
}
