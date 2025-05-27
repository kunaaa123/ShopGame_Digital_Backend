package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kunaaa123/SHOPGAME/internal/server"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	if err := server.RunServer(); err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}
