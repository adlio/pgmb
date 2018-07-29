package pgmb

import (
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

// QueryFunc allows chaining of squirrel statements
type QueryFunc func(sq.SelectBuilder) sq.SelectBuilder

// EchoSQL can be inserted in a find command to output the SQL and arguments accumulated to
// that point.
//
func EchoSQL() QueryFunc {
	return func(b sq.SelectBuilder) sq.SelectBuilder {
		sql, args, err := b.ToSql()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("------------------------------- EchoSQL() ----------------------------------")
		fmt.Println(sql)
		for i, arg := range args {
			fmt.Printf("--- arg %d ----\n", i)
			fmt.Println(arg)
		}
		fmt.Println("------------------------------ End EchoSQL() -------------------------------")
		return b
	}
}

// Where is a wrapper to a Squirrel Where()
func Where(cmd string, args ...interface{}) QueryFunc {
	return func(b sq.SelectBuilder) sq.SelectBuilder {
		sql, args, err := sqlx.In(cmd, args...)
		if err != nil {
			panic(err)
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
