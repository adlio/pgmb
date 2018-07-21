package pgmb

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

func WrapDB(db *sql.DB) *sqlx.DB {
	return sqlx.NewDb(db, "postgres")
}
