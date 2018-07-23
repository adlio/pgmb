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

// ReleaseGroupPrimaryTypeMap returns a map of every release_group_primary_type in the database
// keyed by its ID for easy linking to associations.
//
func ReleaseGroupPrimaryTypeMap(db DB) (types map[int64]*ReleaseGroupPrimaryType, err error) {
	types = make(map[int64]*ReleaseGroupPrimaryType)
	results := make([]*ReleaseGroupPrimaryType, 0)
	err = Select(db, &results, ReleaseGroupPrimaryTypeQuery())
	for _, rgt := range results {
		types[rgt.ID] = rgt
	}
	return
}

// ReleaseGroupPrimaryTypeQuery is the base query for working with release_group_primary_type data
//
func ReleaseGroupPrimaryTypeQuery() sq.SelectBuilder {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("id, gid, name, child_order").
		From("release_group_primary_type")
}
