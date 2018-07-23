package pgmb

import (
	"time"

	sq "github.com/Masterminds/squirrel"
	uuid "github.com/satori/go.uuid"
)

// Recording represents an entry in the recording table in
// the MusicBrainz database.
type Recording struct {
	ID          int64
	GID         uuid.UUID
	Name        string
	Length      *int64
	Comment     string
	LastUpdated time.Time
}

// RecordingQuery builds the default query for working with the recording table
func RecordingQuery() sq.SelectBuilder {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("id, gid, name, length, comment").
		From("recording")
}

// FindRecordings returns recordings matching the supplied criteria
func FindRecordings(db DB, clauses ...QueryFunc) (recordings []*Recording, err error) {
	recordings = make([]*Recording, 0)
	err = Select(db, &recordings, RecordingQuery(), clauses...)
	return
}
