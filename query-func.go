package pgmb

import (
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type QueryFunc func(sq.SelectBuilder) sq.SelectBuilder

func EchoSQL() QueryFunc {
	return func(b sq.SelectBuilder) sq.SelectBuilder {
		sql, args, err := b.ToSql()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("------------------------------- EchoSQL() -------------------------------")
		fmt.Println(sql)
		for i, arg := range args {
			fmt.Printf("--- arg %d ----\n", i)
			fmt.Println(arg)
		}
		fmt.Println("------------------------------ End EchoSQL() -------------------------------")
		return b
	}
}

// WithGID builds a QueryFunc to exactly match the GID field on any table
func WithGID(gid uuid.UUID) QueryFunc {
	return func(b sq.SelectBuilder) sq.SelectBuilder {
		return b.Where("gid = ?", fmt.Sprintf("%s", gid.String()))
	}
}

func IDIn(ids []int64) QueryFunc {
	return func(b sq.SelectBuilder) sq.SelectBuilder {
		sql, args, _ := sqlx.In("id IN (?)", ids)
		return b.Where(sql, args...)
	}
}

// Named builds a QueryFunc to exactly match the name fields on any table
func Named(name string) QueryFunc {
	return func(b sq.SelectBuilder) sq.SelectBuilder {
		return b.Where("name = ?", name)
	}
}

// FuzzyNamed builds a QueryFunc to fuzzy-match the name field on any table
func FuzzyNamed(name string) QueryFunc {
	return func(b sq.SelectBuilder) sq.SelectBuilder {
		return b.Where("lower(name) % lower(?)", name)
	}
}
