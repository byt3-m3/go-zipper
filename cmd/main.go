package main

import (
	server2 "go-zip/internal/server"
	"go-zip/internal/vars"
	"log"
)

func main() {
	server := server2.NewServer(server2.WithListenAddress(vars.WebserverListenAddress))

	if err := server.Run(); err != nil {
		log.Fatalln(err)
	}
}
