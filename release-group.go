package pgmb

import (
	"github.com/lib/pq"

	sq "github.com/Masterminds/squirrel"
	"github.com/satori/go.uuid"
)

// ReleaseGroup represents an entry in the MusicBrainz database
// release_group table.
type ReleaseGroup struct {
	ID               int64
	GID              uuid.UUID
	Name             string
	ArtistCreditID   int64                        `db:"artist_credit"`
	ArtistCredit     *ArtistCredit                `db:"-"`
	TypeID           *int64                       `db:"type"`
	Type             *ReleaseGroupPrimaryType     `db:"-"`
	SecondaryTypeIDs pq.Int64Array                `db:"secondary_type_ids"`
	SecondaryTypes   []*ReleaseGroupSecondaryType `db:"-"`
	Comment          string
}

// GetReleaseGroup returns the first ReleaseGroup result from the supplied dynamic query
// parameters.
//
func GetReleaseGroup(db DB, clauses ...QueryFunc) (releaseGroup *ReleaseGroup, err error) {
	clauses = append(clauses, Limit(1))
	groups, err := FindReleaseGroups(db, clauses...)
	if err != nil {
		return
	}
	if len(groups) > 0 {
		releaseGroup = groups[0]
	}
	return
}

// FindReleaseGroups retrieves a slice of ReleaseGroup based on a dynamically built
// query.
//
func FindReleaseGroups(db DB, clauses ...QueryFunc) (groups []*ReleaseGroup, err error) {
	groups = make([]*ReleaseGroup, 0)
	err = Select(db, &groups, ReleaseGroupQuery(), clauses...)
	if err != nil {
		return
	}

	err = loadReleaseGroupArtistCredits(db, groups)
	if err != nil {
		return
	}

	err = loadReleaseGroupPrimaryTypes(db, groups)
	if err != nil {
		return
	}

	err = loadReleaseGroupSecondaryTypes(db, groups)
	return
}

// ReleaseGroupMap returns a mapping of ReleaseGroup IDs to ReleaseGroup structs
func ReleaseGroupMap(db DB, ids []int64) (groups map[int64]*ReleaseGroup, err error) {
	groups = make(map[int64]*ReleaseGroup)
	results, err := FindReleaseGroups(db, Where("id IN (?)", ids))
	for _, group := range results {
		groups[group.ID] = group
	}
	return
}

// ReleaseGroupQuery is the base query for working with release_group data
//
func ReleaseGroupQuery() sq.SelectBuilder {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(`
			id, gid, name, type, artist_credit, comment,
			ARRAY(
				SELECT j.secondary_type
				FROM release_group_secondary_type_join j
				WHERE j.release_group = release_group.id
			) as secondary_type_ids
		`).
		From("release_group")
}

func loadReleaseGroupArtistCredits(db DB, groups []*ReleaseGroup) error {
	ids := make([]int64, len(groups))
	for i, rel := range groups {
		ids[i] = rel.ArtistCreditID
	}
	credits, err := ArtistCreditMap(db, ids)
	for _, group := range groups {
		group.ArtistCredit, _ = credits[group.ArtistCreditID]
	}
	return err
}

func loadReleaseGroupPrimaryTypes(db DB, groups []*ReleaseGroup) error {
	types, err := ReleaseGroupPrimaryTypeMap(db)
	if err != nil {
		return err
	}
	for _, group := range groups {
		if group.TypeID != nil {
			group.Type, _ = types[*group.TypeID]
		}
	}
	return nil
}

func loadReleaseGroupSecondaryTypes(db DB, groups []*ReleaseGroup) error {
	var err error
	typeMap, err := ReleaseGroupSecondaryTypeMap(db)
	for _, group := range groups {
		if len(group.SecondaryTypeIDs) > 0 {
			group.SecondaryTypes = make([]*ReleaseGroupSecondaryType, len(group.SecondaryTypeIDs))
			for i, typeID := range group.SecondaryTypeIDs {
				group.SecondaryTypes[i], _ = typeMap[typeID]
			}
		}
	}
	return err
}
