package main

import (
	"log"

	"github.com/altamsh04/go-users-api/internal/database"
	"github.com/altamsh04/go-users-api/internal/server"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	server.Start()
}
