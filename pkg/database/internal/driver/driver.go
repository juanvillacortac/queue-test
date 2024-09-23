package driver

import (
	"database/sql"

	"github.com/juanvillacortac/bank-queue/pkg/database/internal/sqlhooks"
	"github.com/juanvillacortac/bank-queue/pkg/database/internal/sqlhooks/hooks"
	"github.com/mattn/go-sqlite3"
)

var SQLiteDriver = "sqlite-driver"

func init() {
	sql.Register(SQLiteDriver, sqlhooks.Wrap(&sqlite3.SQLiteDriver{}, getHooks()))
}

func getHooks() sqlhooks.Hooks {
	return sqlhooks.Compose(
		hooks.NewLogHook(),
	)
}
