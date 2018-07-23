package pgmb

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/satori/go.uuid"
)

type Release struct {
	ID               int64
	GID              uuid.UUID
	Name             string
	ArtistCreditID   int64 `db:"artist_credit"`
	ArtistCreditName string
	ReleaseGroupID   int64 `db:"release_group"`
	Barcode          *string
	Comment          string
	Quality          int64
}

func FindReleases(db DB, clauses ...QueryFunc) (rs []*Release, err error) {
	rs = make([]*Release, 0)
	err = Select(db, &rs, ReleaseQuery(), clauses...)
	return
}

func ReleaseQuery() sq.SelectBuilder {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(`
			release.id, release.gid, release.name, release.artist_credit, release.release_group,
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
