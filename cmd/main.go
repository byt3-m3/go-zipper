package main

import (
	server2 "go-zip/internal/server"
	"log"
)

func main() {
	server := server2.NewServer(server2.WithListenAddress("localhost:8080"))

	if err := server.Run(); err != nil {
		log.Fatalln(err)
	}
}
