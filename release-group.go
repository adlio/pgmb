package pgmb

import (
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

func FindReleaseGroups(db DB, criteria ...Queryer) (rgs []*ReleaseGroup, err error) {
	rgs = make([]*ReleaseGroup, 0)
	q := Query().
		Select("id, gid, name").
		From("release_group").
		Limit(200)
	err = Find(db, &rgs, q, criteria...)
	return
}
