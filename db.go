package pgmb

import (
	"bytes"
	"database/sql"
	"strings"

	"github.com/jmoiron/sqlx"
)

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

// ToSnakeCase converts a string to snake case, words separated with underscores.
// It's intended to be used with NameMapper to map struct field names to snake case database fields.
func ToSnakeCase(src string) string {
	thisUpper := false
	prevUpper := false

	buf := bytes.NewBufferString("")
	for i, v := range src {
		if v >= 'A' && v <= 'Z' {
			thisUpper = true
		} else {
			thisUpper = false
		}
		if i > 0 && thisUpper && !prevUpper {
			buf.WriteRune('_')
		}
		prevUpper = thisUpper
		buf.WriteRune(v)
	}
	return strings.ToLower(buf.String())
}
