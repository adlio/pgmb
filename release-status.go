package pgmb

import (
	sq "github.com/Masterminds/squirrel"
	uuid "github.com/satori/go.uuid"
)

// ReleaseStatus represents an entry in the release_status table
// in the MusicBrainz database.
type ReleaseStatus struct {
	ID          int64
	GID         uuid.UUID
	Name        string
	ChildOrder  int64 `db:"child_order"`
	Description string
}

// ReleaseStatusMap returns a map of every release_status in the
// database, keyed by its ID for easy linking to associations.
//
func ReleaseStatusMap(db DB) (statuses map[int64]*ReleaseStatus, err error) {
	statuses = make(map[int64]*ReleaseStatus)
	rs := make([]*ReleaseStatus, 0)
	err = Select(db, &rs, ReleaseStatusQuery())
	for _, status := range rs {
		statuses[status.ID] = status
	}
	return
}

// ReleaseStatusQuery is the base query for working with release_status data
//
func ReleaseStatusQuery() sq.SelectBuilder {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("id, gid, name, child_order, description").
		From("release_status")
}
