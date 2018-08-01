package pgmb

import (
	uuid "github.com/satori/go.uuid"
)

// Area represents an entry in the MusicBrainz area table
//
type Area struct {
	ID   int64
	GID  uuid.UUID
	Name string
}

// AreaMap returns a mapping of Area IDs to Area structs
//
func AreaMap(db DB, ids []int64) (areas map[int64]*Area, err error) {
	areas = make(map[int64]*Area)
	results, err := Areas(db).Where("id IN (?)", ids).All()
	for _, area := range results {
		areas[area.ID] = area
	}
	return
}
