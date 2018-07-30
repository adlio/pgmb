package pgmb

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/satori/go.uuid"
)

// Track represents an entry in the track table in the MusicBrainz
// database.
type Track struct {
	ID             int64
	GID            uuid.UUID
	Name           string
	Position       int64
	Number         string
	ArtistCreditID int64         `db:"artist_credit"`
	ArtistCredit   *ArtistCredit `db:"-"`
	RecordingID    int64         `db:"recording"`
	Recording      *Recording    `db:"-"`
	Length         *int64
	IsDataTrack    bool `db:"is_data_track"`
}

// TrackCollection is an alias for a slice of Track
type TrackCollection []*Track

// FindTracks rertrieves a slice of Track based on a dynamically built query
//
func FindTracks(db DB, clauses ...QueryFunc) (tracks TrackCollection, err error) {
	tracks = make(TrackCollection, 0)
	err = Select(db, &tracks, TrackQuery(), clauses...)
	if err != nil {
		return
	}

	err = loadTrackArtistCredits(db, tracks)
	if err != nil {
		return
	}

	err = loadTrackRecordings(db, tracks)
	return
}

// TrackQuery is the base query for working with track data.
//
func TrackQuery() sq.SelectBuilder {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(`
			track.id, track.gid, track.name, track.position, track.number,
			track.artist_credit, track.recording,
			track.length
		`).
		From("track")
}

func loadTrackArtistCredits(db DB, tracks TrackCollection) error {
	ids := make([]int64, len(tracks))
	for i, rel := range tracks {
		ids[i] = rel.ArtistCreditID
	}
	credits, err := ArtistCreditMap(db, ids)
	for _, track := range tracks {
		track.ArtistCredit, _ = credits[track.ArtistCreditID]
	}
	return err
}

func loadTrackRecordings(db DB, tracks TrackCollection) error {
	ids := make([]int64, len(tracks))
	for i, rel := range tracks {
		ids[i] = rel.RecordingID
	}
	credits, err := RecordingMap(db, ids)
	for _, track := range tracks {
		track.Recording, _ = credits[track.RecordingID]
	}
	return err
}
