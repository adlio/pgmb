package pgmb

import (
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
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

// ReleaseSelect builds the default query for fetching release data
//
func ReleaseSelect() sq.SelectBuilder {
	return sq.StatementBuilder.
		Select("id, gid, name, artist_credit, release_group, status, packaging, comment, barcode, quality").
		From("release")
}

// WithAssociations attaches all associations to the returned Release structs from
// other database tables.
//
func (q ReleaseQuery) WithAssociations() ReleaseQuery {
	q.processors = append(q.processors, loadReleaseArtistCredits)
	q.processors = append(q.processors, loadReleaseEvents)
	q.processors = append(q.processors, loadReleaseStatuses)
	q.processors = append(q.processors, loadReleasePackagings)
	return q
}

// ReleaseMap returns a mapping of Release IDs to Release structs
func ReleaseMap(db DB, ids []int64) (releases map[int64]*Release, err error) {
	releases = make(map[int64]*Release)
	results, err := Releases(db).Where("id IN (?)", ids).WithAssociations().All()
	for _, release := range results {
		releases[release.ID] = release
	}
	return
}

func loadReleaseArtistCredits(db DB, releases ReleaseCollection) error {
	ids := make([]int64, len(releases))
	for i, rel := range releases {
		ids[i] = rel.ArtistCreditID
	}
	credits, err := ArtistCreditMap(db, ids)
	if err != nil {
		return errors.Wrap(err, "failed to build ArtistCreditMap in loadReleaseArtistCredits")
	}
	for _, release := range releases {
		release.ArtistCredit, _ = credits[release.ArtistCreditID]
	}
	return nil
}

func loadReleaseEvents(db DB, releases ReleaseCollection) error {
	ids := make([]int64, len(releases))
	for i, rel := range releases {
		ids[i] = rel.ID
	}
	events, err := ReleaseEventMap(db, ids)
	if err != nil {
		return errors.Wrap(err, "Failed to build ReleaseEventMap in loadReleaseEvents")
	}
	for _, release := range releases {
		release.ReleaseEvents = events[release.ID]
	}
	return nil
}

func loadReleaseStatuses(db DB, releases ReleaseCollection) error {
	statuses, err := ReleaseStatusMap(db)
	if err != nil {
		return errors.Wrap(err, "Failed to build ReleaseStatusMap in loadReleaseStatuses")
	}
	for _, release := range releases {
		if release.StatusID != nil {
			release.Status, _ = statuses[*release.StatusID]
		}
	}
	return nil
}

func loadReleasePackagings(db DB, releases ReleaseCollection) error {
	packagings, err := ReleasePackagingMap(db)
	if err != nil {
		return errors.Wrap(err, "Failed to build ReleasePackagingMap in loadReleasePackagings")
	}
	for _, release := range releases {
		if release.PackagingID != nil {
			release.Packaging, _ = packagings[*release.PackagingID]
		}
	}
	return nil
}
