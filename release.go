package pgmb

import (
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/satori/go.uuid"
)

// Release represents an entry in the release table in the
// MusicBrainz database.
type Release struct {
	ID             int64
	GID            uuid.UUID
	Name           string
	ArtistCreditID int64             `db:"artist_credit"`
	ArtistCredit   *ArtistCredit     `db:"-"`
	StatusID       *int64            `db:"status"`
	Status         *ReleaseStatus    `db:"-"`
	PackagingID    *int64            `db:"packaging"`
	Packaging      *ReleasePackaging `db:"-"`
	ReleaseGroupID int64             `db:"release_group"`
	ReleaseEvents  []*ReleaseEvent   `db:"-"`
	Barcode        *string
	Comment        string
	Quality        int64
}

// EarliestReleaseDate finds the Time of the earliest attached
// ReleaseEvent
func (r *Release) EarliestReleaseDate() time.Time {
	var t time.Time
	for _, event := range r.ReleaseEvents {
		if t.IsZero() {
			t = event.Date()
		} else {
			d := event.Date()
			if d.Before(t) {
				t = d
			}
		}
	}
	return t
}

// FindReleases retrieves a slice of Release based on a dynamically built query
//
func FindReleases(db DB, clauses ...QueryFunc) (releases []*Release, err error) {
	releases = make([]*Release, 0)
	err = Select(db, &releases, ReleaseQuery(), clauses...)
	if err != nil {
		return
	}

	err = loadReleaseArtistCredits(db, releases)
	if err != nil {
		return
	}

	err = loadReleaseEvents(db, releases)
	if err != nil {
		return
	}

	err = loadReleaseStatuses(db, releases)
	if err != nil {
		return
	}

	err = loadReleasePackagings(db, releases)
	return
}

// ReleaseMap returns a mapping of Release IDs to Release structs
func ReleaseMap(db DB, ids []int64) (releases map[int64]*Release, err error) {
	releases = make(map[int64]*Release)
	results, err := FindReleases(db, IDIn(ids))
	for _, release := range results {
		releases[release.ID] = release
	}
	return
}

// ReleaseQuery is the base query for working with release data.
//
func ReleaseQuery() sq.SelectBuilder {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(`
			release.id, release.gid, release.name, release.artist_credit, release.release_group,
			release.status, release.packaging,
			release.comment, release.barcode, release.quality
		`).
		From("release")
}

// WhereReleaseIncludesRecording filters FindReleases to those which
// include the supplied Recording ID on one of their media.
//
func WhereReleaseIncludesRecording(rid uuid.UUID) QueryFunc {
	b := func(b sq.SelectBuilder) sq.SelectBuilder {
		b = b.Where(`
			EXISTS (
				SELECT track.id
				FROM track
				JOIN medium on medium.id = track.medium
				JOIN recording ON recording.id = track.recording
				WHERE recording.gid = ?
				AND medium.release = release.id
			)
		`, rid)
		return b
	}
	return b
}

func loadReleaseArtistCredits(db DB, releases []*Release) error {
	ids := make([]int64, len(releases))
	for i, rel := range releases {
		ids[i] = rel.ArtistCreditID
	}
	credits, err := ArtistCreditMap(db, ids)
	for _, release := range releases {
		release.ArtistCredit, _ = credits[release.ArtistCreditID]
	}
	return err
}

func loadReleaseEvents(db DB, releases []*Release) error {
	ids := make([]int64, len(releases))
	for i, rel := range releases {
		ids[i] = rel.ID
	}
	events, err := ReleaseEventMap(db, ids)
	for _, release := range releases {
		release.ReleaseEvents = events[release.ID]
	}
	return err
}

func loadReleaseStatuses(db DB, releases []*Release) error {
	statuses, err := ReleaseStatusMap(db)
	if err != nil {
		return err
	}
	for _, release := range releases {
		if release.StatusID != nil {
			release.Status, _ = statuses[*release.StatusID]
		}
	}
	return nil
}

func loadReleasePackagings(db DB, releases []*Release) error {
	packagings, err := ReleasePackagingMap(db)
	if err != nil {
		return err
	}
	for _, release := range releases {
		if release.PackagingID != nil {
			release.Packaging, _ = packagings[*release.PackagingID]
		}
	}
	return nil
}
