package driver

import (
	"database/sql"

	"github.com/juanvillacortac/bank-queue/pkg/database/internal/sqlhooks"
	"github.com/juanvillacortac/bank-queue/pkg/database/internal/sqlhooks/hooks"
	"modernc.org/sqlite"
)

var SQLiteDriver = "sqlite-driver"

func init() {
	sql.Register(SQLiteDriver, sqlhooks.Wrap(&sqlite.Driver{}, getHooks()))
}

func getHooks() sqlhooks.Hooks {
	return sqlhooks.Compose(
		hooks.NewLogHook(),
	)
}
