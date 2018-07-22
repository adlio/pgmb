package pgmb

import (
	"bytes"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/satori/go.uuid"

	sq "github.com/Masterminds/squirrel"
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
	convertedDB.Exec("SELECT set_limit(0.5);")
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

type Queryer interface {
	Query(squirrel.SelectBuilder) squirrel.SelectBuilder
}

type WithGID uuid.UUID

func (gid WithGID) Query(b squirrel.SelectBuilder) squirrel.SelectBuilder {
	return b.Where("gid = ?", fmt.Sprintf("%s", uuid.UUID(gid).String()))
}

func Query() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}

func Get(db DB, dest interface{}, q sq.SelectBuilder, criteria ...Queryer) error {
	for _, criteria := range criteria {
		q = criteria.Query(q)
	}
	sql, args, err := q.ToSql()
	if err != nil {
		return err
	}
	err = db.Get(dest, db.Rebind(sql), args...)
	return nil
}

func Find(db DB, dest interface{}, q sq.SelectBuilder, criteria ...Queryer) error {
	for _, criteria := range criteria {
		q = criteria.Query(q)
	}
	sql, args, err := q.ToSql()
	if err != nil {
		return err
	}
	err = db.Select(dest, db.Rebind(sql), args...)
	return nil
}
