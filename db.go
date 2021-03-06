package pgmb

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

//go:generate go run gen.go
//go:generate go fmt .

// Sqler blah
type Sqler interface {
	ToSql() (string, []interface{}, error)
}

// DB represents a connection to the MusicBrainz database. Its interface is
// satisfied by *sqlx.DB, but to avoid adding a sqlx.DB depedency in clients,
// the NewDB function converts a supplied *sql.DB into *sqlx.DB.
type DB interface {
	Get(interface{}, string, ...interface{}) error
	Select(interface{}, string, ...interface{}) error
	Rebind(string) string
}

// NewDB creates a pgmb.DB from a supplied sql.DB. It is not necessary to use
// this function if your application is using the sqlx package. The sqlx.NewDB
// satisfies our DB interface directly.
func NewDB(db *sql.DB) DB {
	convertedDB := sqlx.NewDb(db, "postgres")
	return convertedDB
}

// Get is inspired by sqlx.Get. It maps a query to a single struct destination.
//
func Get(db DB, dest interface{}, q Sqler) error {
	sql, args, _ := q.ToSql()
	err := db.Get(dest, db.Rebind(sql), args...)
	return err
}

// Select is inspired by sqlx.Select. It maps a query with multiple results to a slice
// of structs.
//
func Select(db DB, dest interface{}, q Sqler) error {
	sql, args, err := q.ToSql()
	if err != nil {
		return errors.Wrap(err, "Squirrel ToSql() failed.")
	}

	sql, args, err = sqlx.In(sql, args...)
	if err != nil {
		return errors.Wrap(err, "sqlx.In failed.")
	}
	err = db.Select(dest, db.Rebind(sql), args...)
	return err
}
