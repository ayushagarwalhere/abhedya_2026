package main

import (
	"log"

	"abhedya_2026/internal/server"
)

func main() {
	router := server.NewServer()
	if err := router.Run(":3000"); err != nil {
		log.Fatal("Unable to start the server", err)
	}
}
