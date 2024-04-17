package main

import (
	"log"

	"github.com/denismathan/goAuth/src/httpHandler"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("hello World")
	server := httpHandler.NewHttpServer()
	server.Start()
}
