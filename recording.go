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

// FindRecordings returns recordings matching the supplied criteria
func FindRecordings(db DB, clauses ...QueryFunc) (recordings []*Recording, err error) {
	recordings = make([]*Recording, 0)
	err = Select(db, &recordings, RecordingQuery(), clauses...)
	return
}

// RecordingMap returns a mapping of Recording IDs to Recording structs
func RecordingMap(db DB, ids []int64) (recordings map[int64]*Recording, err error) {
	recordings = make(map[int64]*Recording)
	results, err := FindRecordings(db, IDIn(ids))
	for _, recording := range results {
		recordings[recording.ID] = recording
	}
	return
}

// RecordingQuery builds the default query for working with the recording table
func RecordingQuery() sq.SelectBuilder {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("id, gid, name, length, comment").
		From("recording")
}
