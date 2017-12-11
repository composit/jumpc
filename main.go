package main

import (
	"os"

	"github.com/composit/jumpc/handlers"
)

func main() {
	handlers.Listen(os.Args[1])
}
