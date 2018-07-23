package pgmb

import (
	sq "github.com/Masterminds/squirrel"
	uuid "github.com/satori/go.uuid"
)

// ReleasePackaging represents an entry in the release_packaging table
// in the MusicBrainz database.
type ReleasePackaging struct {
	ID          int64
	GID         uuid.UUID
	Name        string
	ChildOrder  int64 `db:"child_order"`
	Description *string
}

// ReleasePackagingMap returns a map of every release_packaging in the
// database, keyed by its ID for easy linking to associations.
//
func ReleasePackagingMap(db DB) (packagings map[int64]*ReleasePackaging, err error) {
	packagings = make(map[int64]*ReleasePackaging)
	rs := make([]*ReleasePackaging, 0)
	err = Select(db, &rs, ReleasePackagingQuery())
	for _, status := range rs {
		packagings[status.ID] = status
	}
	return
}

// ReleasePackagingQuery is the base query for working with release_packaging data
//
func ReleasePackagingQuery() sq.SelectBuilder {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("id, gid, name, child_order, description").
		From("release_packaging")
}
