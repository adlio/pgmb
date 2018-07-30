package pgmb

import (
	"fmt"
	"io"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

// QueryFunc allows chaining of squirrel statements
type QueryFunc func(sq.SelectBuilder) sq.SelectBuilder

// Where is a wrapper to a Squirrel Where()
func Where(cmd string, args ...interface{}) QueryFunc {
	return func(b sq.SelectBuilder) sq.SelectBuilder {
		sql, args, err := sqlx.In(cmd, args...)
		if err != nil {
			log.Fatal(err)
		}
		return b.Where(sql, args...)
	}
}

// Limit builds a QueryFunc to limit the results to the supplied number
func Limit(n uint64) QueryFunc {
	return func(b sq.SelectBuilder) sq.SelectBuilder {
		return b.Limit(n)
	}
}

// EchoSQL can be inserted in a find command to output the SQL and arguments accumulated to
// that point.
//
func EchoSQL(w io.Writer) QueryFunc {
	return func(b sq.SelectBuilder) sq.SelectBuilder {
		sql, args, err := b.ToSql()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(w, "------------------------------- EchoSQL() ----------------------------------")
		fmt.Fprintln(w, sql)
		for i, arg := range args {
			fmt.Fprintf(w, "--- arg %d ----\n", i)
			fmt.Fprintln(w, arg)
		}
		fmt.Fprintln(w, "------------------------------ End EchoSQL() -------------------------------")
		return b
	}
}
