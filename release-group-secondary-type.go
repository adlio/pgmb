package pgmb

import (
	sq "github.com/Masterminds/squirrel"
	uuid "github.com/satori/go.uuid"
)

// ReleaseGroupSecondaryType represents an entry in the release_group_secondary_type table
// in the MusicBrainz database.
type ReleaseGroupSecondaryType struct {
	ID   int64
	GID  uuid.UUID
	Name string
}

// ReleaseGroupSecondaryTypeMap returns a map of every release_group_secondary_type in the database
// keyed by its ID for easy linking to associations.
//
func ReleaseGroupSecondaryTypeMap(db DB) (types map[int64]*ReleaseGroupSecondaryType, err error) {
	types = make(map[int64]*ReleaseGroupSecondaryType)
	results := make([]*ReleaseGroupSecondaryType, 0)
	err = Select(db, &results, ReleaseGroupSecondaryTypeQuery())
	for _, rgt := range results {
		types[rgt.ID] = rgt
	}
	return
}

// ReleaseGroupSecondaryTypeQuery is the base query for working with release_group_primary_type data
//
func ReleaseGroupSecondaryTypeQuery() sq.SelectBuilder {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("id, gid, name").
		From("release_group_secondary_type")
}
