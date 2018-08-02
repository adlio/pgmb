package pgmb

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
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

// DeepTrackSelect is the base query for working with track data.
//
func DeepTrackSelect() sq.SelectBuilder {
	q := sq.StatementBuilder.
		Select(`
			track.id, track.gid, track.name, track.position, track.number,
			track.artist_credit, track.recording,
			track.length,
			release.id as release,
			release.release_group as release_group
		`).
		From("track").
		Join("medium ON medium.id = track.medium").
		Join("release ON release.id = medium.release").
		Join("recording ON recording.id = track.recording")
	return q
}

// WithAssociations attaches all associations to the structs returned
// by a query.
//
func (q DeepTrackQuery) WithAssociations() DeepTrackQuery {
	q.processors = append(q.processors, loadDeepTrackArtistCredits)
	q.processors = append(q.processors, loadDeepTrackRecordings)
	q.processors = append(q.processors, loadDeepTrackReleases)
	q.processors = append(q.processors, loadDeepTrackReleaseGroups)
	return q
}

func loadDeepTrackArtistCredits(db DB, tracks DeepTrackCollection) error {
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

func loadDeepTrackRecordings(db DB, tracks DeepTrackCollection) error {
	ids := make([]int64, len(tracks))
	for i, rel := range tracks {
		ids[i] = rel.RecordingID
	}
	credits, err := RecordingMap(db, ids)
	if err != nil {
		err = errors.Wrap(err, "Error building RecordingMap in loadDeepTrackRecordings")
		return err
	}
	for _, track := range tracks {
		track.Recording, _ = credits[track.RecordingID]
	}
	return nil
}

func loadDeepTrackReleases(db DB, tracks DeepTrackCollection) error {
	ids := make([]int64, len(tracks))
	for i, rel := range tracks {
		ids[i] = rel.ReleaseID
	}
	releases, err := ReleaseMap(db, ids)
	if err != nil {
		return errors.Wrap(err, "Error building ReleaseMap in loadDeepTrackReleases")
	}
	for _, track := range tracks {
		track.Release, _ = releases[track.ReleaseID]
	}
	return nil
}

func loadDeepTrackReleaseGroups(db DB, tracks DeepTrackCollection) error {
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
