package pgmb

import (
	sq "github.com/Masterminds/squirrel"
	uuid "github.com/satori/go.uuid"
)

// ReleaseGroupPrimaryType represents an entry in the release_group_primary_type
// table in the MusicBrainz database.
type ReleaseGroupPrimaryType struct {
	ID         int64
	GID        uuid.UUID
	Name       string
	ChildOrder int64 `db:"child_order"`
}

// ReleaseGroupPrimaryTypeSelect builds the default select query for release_group_primary_type
// data.
func ReleaseGroupPrimaryTypeSelect() sq.SelectBuilder {
	return sq.StatementBuilder.
		Select("id, gid, name, child_order").
		From("release_group_primary_type")
}

// ReleaseGroupPrimaryTypeMap returns a map of every release_group_primary_type in the database
// keyed by its ID for easy linking to associations.
//
func ReleaseGroupPrimaryTypeMap(db DB) (types map[int64]*ReleaseGroupPrimaryType, err error) {
	types = make(map[int64]*ReleaseGroupPrimaryType)
	results, err := ReleaseGroupPrimaryTypes(db).All()
	for _, rgt := range results {
		types[rgt.ID] = rgt
	}
	return
}
