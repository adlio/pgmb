package pgmb

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/satori/go.uuid"
)

// ReleaseGroup represents an entry in the MusicBrainz database
// release_group table.
type ReleaseGroup struct {
	ID      int64
	GID     uuid.UUID
	Name    string
	Type    *ReleaseGroupPrimaryType
	Comment string
}

type ReleaseGroupPrimaryType struct {
	ID   int64
	GID  uuid.UUID
	Name string
}

type ReleaseGroupSecondaryType struct {
	ID   int64
	GID  uuid.UUID
	Name string
}

func ReleaseGroupQuery() sq.SelectBuilder {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("id, gid, name").
		From("release_group")
}

func FindReleaseGroups(db DB, clauses ...QueryFunc) (rgs []*ReleaseGroup, err error) {
	rgs = make([]*ReleaseGroup, 0)
	err = Find(db, &rgs, ReleaseGroupQuery(), clauses...)
	return
}
