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

// ReleasePackagingSelect builds the default select query for the
// release_packaging table
func ReleasePackagingSelect() sq.SelectBuilder {
	return sq.StatementBuilder.
		Select("id, gid, name, child_order, description").
		From("release_packaging")
}

// ReleasePackagingMap returns a map of every release_packaging in the
// database, keyed by its ID for easy linking to associations.
//
func ReleasePackagingMap(db DB) (packagings map[int64]*ReleasePackaging, err error) {
	packagings = make(map[int64]*ReleasePackaging)
	results, err := ReleasePackagings(db).All()
	for _, status := range results {
		packagings[status.ID] = status
	}
	return
}
