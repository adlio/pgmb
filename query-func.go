package pgmb

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	uuid "github.com/satori/go.uuid"
)

type QueryFunc func(sq.SelectBuilder) sq.SelectBuilder

// WithGID builds a QueryFunc to exactly match the GID field on any table
func WithGID(gid uuid.UUID) QueryFunc {
	return func(b sq.SelectBuilder) sq.SelectBuilder {
		return b.Where("gid = ?", fmt.Sprintf("%s", gid.String()))
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
