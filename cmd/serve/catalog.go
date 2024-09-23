package main

import (
	"os"

	database "github.com/juanvillacortac/bank-queue/pkg/database"
)

var (
	dbAddress     = os.Getenv("DATABASE")
	serverAddress = os.Getenv("ADDRESS")
)

func init() {
	if serverAddress == "" {
		serverAddress = ":8080"
	}
	if dbAddress == "" {
		dbAddress = database.DefaultConnectionString
	}
}
