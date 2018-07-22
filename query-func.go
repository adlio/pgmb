package pgmb

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	uuid "github.com/satori/go.uuid"
)

type QueryFunc func(sq.SelectBuilder) sq.SelectBuilder

func WithGID(gid uuid.UUID) QueryFunc {
	return func(b sq.SelectBuilder) sq.SelectBuilder {
		return b.Where("gid = ?", fmt.Sprintf("%s", gid.String()))
	}
}

func Named(name string) QueryFunc {
	return func(b sq.SelectBuilder) sq.SelectBuilder {
		return b.Where("name = ?", name)
	}
}

func FuzzyNamed(name string) QueryFunc {
	return func(b sq.SelectBuilder) sq.SelectBuilder {
		return b.Where("lower(name) % lower(?)", name)
	}
}

func Query() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}
