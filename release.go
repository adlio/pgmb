package pgmb

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/satori/go.uuid"
)

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
	Barcode        *string
	Comment        string
	Quality        int64
}

func FindReleases(db DB, clauses ...QueryFunc) (releases []*Release, err error) {
	releases = make([]*Release, 0)
	err = Select(db, &releases, ReleaseQuery(), clauses...)
	if err != nil {
		return
	}

	err = loadArtistCredits(db, releases)
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

func ReleaseQuery() sq.SelectBuilder {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(`
			release.id, release.gid, release.name, release.artist_credit, release.release_group,
			release.status, release.packaging,
			release.comment, release.barcode, release.quality
		`).
		From("release")
}

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

func loadArtistCredits(db DB, releases []*Release) error {
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
