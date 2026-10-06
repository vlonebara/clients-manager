package main

import (
	"clients-manager/internal/server"
	"log"
)

type Config struct {
	PORT string
}

func main() {
	log.Println("Starting server")

	config := Config{
		PORT: ":8080",
	}
	server.Start(config.PORT)
}
