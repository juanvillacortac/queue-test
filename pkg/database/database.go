package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	database "github.com/juanvillacortac/bank-queue/pkg/database/internal"
)

var Instance *sqlx.DB

const DefaultConnectionString = "file::memory:?cache=shared"

func InitDatabase(addr string) error {
	log.Println("Connecting SQLite3 database", addr)

	db, err := database.ConnectToNewSQLInstance(addr)
	if err != nil {
		return err
	}

	Instance = database.GetSQLXInstance(db)

	return nil
}
