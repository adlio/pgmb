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

// RecordingSelect creates the default select query for recording data
func RecordingSelect() sq.SelectBuilder {
	return sq.StatementBuilder.
		Select("id, gid, name, artist_credit, length, comment").
		From("recording")
}

// WithAssociations adds a processor to load all associated objects
// on the returned Recording entities.
//
func (q RecordingQuery) WithAssociations() RecordingQuery {
	q.processors = append(q.processors, loadRecordingArtistCredits)
	return q
}

// RecordingMap returns a mapping of Recording IDs to Recording structs
func RecordingMap(db DB, ids []int64) (recordings map[int64]*Recording, err error) {
	recordings = make(map[int64]*Recording)
	results, err := Recordings(db).Where("id IN (?)", ids).WithAssociations().All()
	for _, recording := range results {
		recordings[recording.ID] = recording
	}
	return
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
