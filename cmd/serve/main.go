package main

import (
	"log"

	"github.com/juanvillacortac/bank-queue/app"
	"github.com/juanvillacortac/bank-queue/pkg/database"
	"github.com/juanvillacortac/bank-queue/pkg/server"
)

func main() {
	if err := database.InitDatabase(dbAddress); err != nil {
		log.Fatalln(err)
	}

	if err := database.SeedDatabase(database.Instance); err != nil {
		log.Fatalln(err)
	}

	s := server.NewServer(server.ServerOptions{
		ServeFS:        app.AppDistFS,
		FSFallbackFile: "_fallback.html",
	})

	log.Println("Starting server on address", serverAddress)
	if err := s.Listen(serverAddress); err != nil {
		log.Fatalln(err)
	}
}
