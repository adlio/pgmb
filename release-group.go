package pgmb

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"

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

func ReleaseGroupSelect() sq.SelectBuilder {
	return sq.StatementBuilder.
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

// WithAssociations attaches all associations to the ReleaseGroup
func (q ReleaseGroupQuery) WithAssociations() ReleaseGroupQuery {
	q.processors = append(q.processors, loadReleaseGroupArtistCredits)
	q.processors = append(q.processors, loadReleaseGroupPrimaryTypes)
	q.processors = append(q.processors, loadReleaseGroupSecondaryTypes)
	return q
}

// ReleaseGroupMap returns a mapping of ReleaseGroup IDs to ReleaseGroup structs
func ReleaseGroupMap(db DB, ids []int64) (groups map[int64]*ReleaseGroup, err error) {
	groups = make(map[int64]*ReleaseGroup)
	results, err := ReleaseGroups(db).Where("id IN (?)", ids).WithAssociations().All()
	for _, group := range results {
		groups[group.ID] = group
	}
	return
}

func loadReleaseGroupArtistCredits(db DB, groups ReleaseGroupCollection) error {
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

func loadReleaseGroupPrimaryTypes(db DB, groups ReleaseGroupCollection) error {
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

func loadReleaseGroupSecondaryTypes(db DB, groups ReleaseGroupCollection) error {
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
