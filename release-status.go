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

// ReleaseStatusSelect builds the default query for release_status data
func ReleaseStatusSelect() sq.SelectBuilder {
	return sq.StatementBuilder.
		Select("id, gid, name, child_order, description").
		From("release_status")
}

// ReleaseStatusMap returns a map of every release_status in the
// database, keyed by its ID for easy linking to associations.
//
func ReleaseStatusMap(db DB) (statuses map[int64]*ReleaseStatus, err error) {
	statuses = make(map[int64]*ReleaseStatus)
	results, err := ReleaseStatuses(db).All()
	for _, status := range results {
		statuses[status.ID] = status
	}
	return
}
