package pgmb

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/satori/go.uuid"
)

// DeepTrack represents an entry in the track table in the MusicBrainz
// database, with deep associations bubbled up to the top level for
// convenience (ArtistCredit, Release and ReleaseGroup in particular).
// This might also have been called AlbumTrack.
//
type DeepTrack struct {
	ID             int64
	GID            uuid.UUID
	Name           string
	Position       int64
	Number         string
	ArtistCreditID int64         `db:"artist_credit"`
	ArtistCredit   *ArtistCredit `db:"-"`
	RecordingID    int64         `db:"recording"`
	Recording      *Recording    `db:"-"`
	ReleaseID      int64         `db:"release"`
	Release        *Release      `db:"-"`
	ReleaseGroupID int64         `db:"release_group"`
	ReleaseGroup   *ReleaseGroup `db:"-"`
	Length         *int64
	IsDataTrack    bool `db:"is_data_track"`
}

// FindDeepTracks rertrieves a slice of Track based on a dynamically built query
//
func FindDeepTracks(db DB, clauses ...QueryFunc) (tracks []*DeepTrack, err error) {
	tracks = make([]*DeepTrack, 0)
	err = Select(db, &tracks, DeepTrackQuery(), clauses...)
	if err != nil {
		return
	}

	err = loadDeepTrackArtistCredits(db, tracks)
	if err != nil {
		return
	}

	err = loadDeepTrackRecordings(db, tracks)
	if err != nil {
		return
	}

	err = loadDeepTrackReleases(db, tracks)
	if err != nil {
		return
	}

	err = loadDeepTrackReleaseGroups(db, tracks)
	return
}

// DeepTrackQuery is the base query for working with track data.
//
func DeepTrackQuery() sq.SelectBuilder {
	q := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(`
			track.id, track.gid, track.name, track.position, track.number,
			track.artist_credit, track.recording,
			track.length,
			release.id as release,
			release.release_group as release_group
		`).
		From("track").
		Join("medium ON medium.id = track.medium").
		Join("release ON release.id = medium.release")
	return q
}

func loadDeepTrackArtistCredits(db DB, tracks []*DeepTrack) error {
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

func loadDeepTrackRecordings(db DB, tracks []*DeepTrack) error {
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

func loadDeepTrackReleases(db DB, tracks []*DeepTrack) error {
	ids := make([]int64, len(tracks))
	for i, rel := range tracks {
		ids[i] = rel.ReleaseID
	}
	releases, err := ReleaseMap(db, ids)
	for _, track := range tracks {
		track.Release, _ = releases[track.ReleaseID]
	}
	return err
}

func loadDeepTrackReleaseGroups(db DB, tracks []*DeepTrack) error {
	ids := make([]int64, len(tracks))
	for i, rel := range tracks {
		ids[i] = rel.ReleaseGroupID
	}
	groups, err := ReleaseGroupMap(db, ids)
	for _, track := range tracks {
		track.ReleaseGroup, _ = groups[track.ReleaseGroupID]
	}
	return err
}
