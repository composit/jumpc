package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/composit/jumpc/handlers"
)

func main() {
	stop := make(chan struct{})
	srv := handlers.Server(os.Args[1], stop)

	<-stop

	log.Println("waiting for requests to finish")
	if err := srv.Shutdown(context.Background()); err != nil {
		fmt.Printf("failed to shutdown properly: %s\n", err)
		os.Exit(1)
	}
	log.Println("exiting")
}
