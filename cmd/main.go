package main

import (
	"github.com/byt3-m3/go-zipper/internal/server"
	"github.com/byt3-m3/go-zipper/internal/vars"
	"log"
)

func main() {
	server := server.NewServer(server.WithListenAddress(vars.WebserverListenAddress))

	if err := server.Run(); err != nil {
		log.Fatalln(err)
	}
}
