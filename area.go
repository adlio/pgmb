package pgmb

import (
	sq "github.com/Masterminds/squirrel"
	uuid "github.com/satori/go.uuid"
)

// Area represents an entry in the MusicBrainz area table
//
type Area struct {
	ID   int64
	GID  uuid.UUID
	Name string
}

// FindAreas retrieves a slice of Area which match the supplied
// criteria.
func FindAreas(db DB, clauses ...QueryFunc) (areas []*Area, err error) {
	areas = make([]*Area, 0)
	err = Select(db, &areas, AreaQuery(), clauses...)
	return
}

// AreaMap returns a mapping of Area IDs to Area structs
//
func AreaMap(db DB, ids []int64) (areas map[int64]*Area, err error) {
	areas = make(map[int64]*Area)
	results, err := FindAreas(db, IDIn(ids))
	for _, area := range results {
		areas[area.ID] = area
	}
	return
}

// AreaQuery builds the default query for the area table in the MusicBrainz
// database.
func AreaQuery() sq.SelectBuilder {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("id, gid, name").
		From("area")
}
