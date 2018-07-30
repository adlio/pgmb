package pgmb

import (
	"time"

	sq "github.com/Masterminds/squirrel"
	uuid "github.com/satori/go.uuid"
)

// Recording represents an entry in the recording table in
// the MusicBrainz database.
type Recording struct {
	ID             int64
	GID            uuid.UUID
	Name           string
	ArtistCreditID int64         `db:"artist_credit"`
	ArtistCredit   *ArtistCredit `db:"-"`
	Length         *int64
	Comment        string
	LastUpdated    time.Time
}

// RecordingCollection is an alias for a slice of Recording
type RecordingCollection []*Recording

// FindRecordings returns recordings matching the supplied criteria
func FindRecordings(db DB, clauses ...QueryFunc) (recordings RecordingCollection, err error) {
	recordings = make(RecordingCollection, 0)
	err = Select(db, &recordings, RecordingQuery(), clauses...)
	if err != nil {
		return
	}

	err = loadRecordingArtistCredits(db, recordings)
	return
}

// RecordingMap returns a mapping of Recording IDs to Recording structs
func RecordingMap(db DB, ids []int64) (recordings map[int64]*Recording, err error) {
	recordings = make(map[int64]*Recording)
	results, err := FindRecordings(db, Where("id IN (?)", ids))
	for _, recording := range results {
		recordings[recording.ID] = recording
	}
	return
}

// RecordingQuery builds the default query for working with the recording table
func RecordingQuery() sq.SelectBuilder {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("id, gid, name, artist_credit, length, comment").
		From("recording")
}

func loadRecordingArtistCredits(db DB, recordings RecordingCollection) error {
	ids := make([]int64, len(recordings))
	for i, rec := range recordings {
		ids[i] = rec.ArtistCreditID
	}
	credits, err := ArtistCreditMap(db, ids)
	for _, recording := range recordings {
		recording.ArtistCredit, _ = credits[recording.ArtistCreditID]
	}
	return err
}
