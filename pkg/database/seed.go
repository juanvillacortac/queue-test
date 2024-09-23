package database

import (
	_ "embed"
	"strings"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

//go:embed scripts/seed.sql
var seedScript string
var password = []byte("admin")

func SeedDatabase(db *sqlx.DB) error {
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	r := strings.NewReplacer("<HASHED_PASSWORD>", string(hashedPassword))
	_, err = db.Exec(r.Replace(seedScript))
	return err
}
