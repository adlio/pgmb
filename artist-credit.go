package pgmb

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
)

// ArtistCredit represents an entry in the MusicBrainz
// artist_credit table.
type ArtistCredit struct {
	ID          int64
	Name        string
	ArtistIDs   pq.Int64Array `db:"artist_ids"`
	Artists     []*Artist     `db:"-"`
	ArtistCount int64         `db:"artist_count"`
	RefCount    int64         `db:"ref_count"`
}

// ArtistCreditCollection is a slice of ArtistCredits
type ArtistCreditCollection []*ArtistCredit

// MBIDs returns the IDs of the collection
func (acc ArtistCreditCollection) MBIDs() []int64 {
	ids := make([]int64, len(acc))
	for i := range acc {
		ids[i] = acc[i].ID
	}
	return ids
}

// UniqueArtistIDs extracts a list of unique Artist ids from a
// collection of ArtistCredit entities.
//
func (acc ArtistCreditCollection) UniqueArtistIDs() []int64 {
	idMap := make(map[int64]bool)
	for _, c := range acc {
		for _, a := range c.ArtistIDs {
			idMap[a] = true
		}
	}
	ids := make([]int64, 0, len(idMap))
	for id := range idMap {
		ids = append(ids, id)
	}
	return ids
}

// IsVariousArtists indicates whether the ArtistCredit represents
// a "Various Artists" record.
//
func (ac *ArtistCredit) IsVariousArtists() bool {
	return ac.ID == 1
}

// FindArtistCredits retrieves a slice of ArtistCredit which
// match the supplied criteria.
//
func FindArtistCredits(db DB, clauses ...QueryFunc) (credits ArtistCreditCollection, err error) {
	credits = make(ArtistCreditCollection, 0)
	err = Select(db, &credits, ArtistCreditQuery(), clauses...)
	if err != nil {
		return
	}

	return
}

// ArtistCreditMap returns a mapping of ArtistCredit IDs to ArtistCredit structs
func ArtistCreditMap(db DB, ids []int64) (credits map[int64]*ArtistCredit, err error) {
	credits = make(map[int64]*ArtistCredit)
	results, err := FindArtistCredits(db, Where("id IN (?)", ids))
	for _, credit := range results {
		credits[credit.ID] = credit
	}
	return
}

// ArtistCreditQuery builds the default query for the artist_credit table
//
func ArtistCreditQuery() sq.SelectBuilder {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(`
			id, name, artist_count, ref_count,
			ARRAY(
				SELECT n.artist
				FROM artist_credit_name n
				WHERE n.artist_credit = artist_credit.id
			) as artist_ids
		`).
		From("artist_credit")
}

func loadArtistCreditArtists(db DB, credits ArtistCreditCollection) error {
	artistMap, err := ArtistMap(db, credits.UniqueArtistIDs())
	if err != nil {
		return err
	}
	for _, credit := range credits {
		if len(credit.ArtistIDs) > 0 {
			credit.Artists = make([]*Artist, len(credit.ArtistIDs))
			for i, artistID := range credit.ArtistIDs {
				credit.Artists[i], _ = artistMap[artistID]
			}
		}
	}
	return nil
}
