package database

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/juanvillacortac/bank-queue/pkg/database/internal/driver"
)

var databaseInstances = map[string]*sql.DB{}
var databaseIntancesMutex = sync.Mutex{}

// Create a handled database instance, returned from a connection pool.
// If connectionString is an empty string, the configmaps default is used instead.
func ConnectToNewSQLInstance(connectionString string) (*sql.DB, error) {
	maxIdleConnections := ValuesPoolConnection.MaxIdleConnections
	maxOpenConnections := ValuesPoolConnection.MaxOpenConnections
	connMaxIdleTime := ValuesPoolConnection.MaxIdleConnections
	connMaxLifetime := ValuesPoolConnection.ConnMaxLifetime

	databaseIntancesMutex.Lock()
	if db, ok := databaseInstances[connectionString]; ok && db != nil {
		ctx, cancelfunc := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancelfunc()
		if err := db.PingContext(ctx); err == nil {
			databaseIntancesMutex.Unlock()
			return db, nil
		}
		db.Close()
		delete(databaseInstances, connectionString)
	}
	databaseIntancesMutex.Unlock()

	driver := driver.SQLiteDriver

	var err error

	// Create connection pool
	db, err := sql.Open(driver, connectionString)
	if err != nil {
		return db, err
	}

	// Maximum Idle Connections
	db.SetMaxIdleConns(maxIdleConnections)
	// Maximum Open Connections
	db.SetMaxOpenConns(maxOpenConnections)
	// Idle Connection Timeout
	db.SetConnMaxIdleTime(time.Duration(connMaxIdleTime) * time.Second)
	// Connection Lifetime
	db.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelfunc()
	if err := db.PingContext(ctx); err != nil {
		return db, err
	}

	databaseIntancesMutex.Lock()
	databaseInstances[connectionString] = db
	databaseIntancesMutex.Unlock()

	return db, nil
}

func GetSQLXInstance(db *sql.DB) *sqlx.DB {
	return sqlx.NewDb(db, "sqlite3")
}
