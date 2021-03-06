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

// ArtistCreditSelect builds the default select for ArtistCredit
func ArtistCreditSelect() sq.SelectBuilder {
	return sq.StatementBuilder.
		Select(`
			id,
			name,
			artist_count,
			ref_count,
			ARRAY(
				SELECT n.artist
				FROM artist_credit_name n
				WHERE n.artist_credit = artist_credit.id
			) as artist_ids`).
		From("artist_credit")
}

// WithAssociations attaches all relevant associations to each entity in
// the results.
//
func (q ArtistCreditQuery) WithAssociations() ArtistCreditQuery {
	q.processors = append(q.processors, loadArtistCreditArtists)
	return q
}

// WhereCreditOrArtistNameMatches adds a where clause to *exactly*
// match the supplied string on either the artist_credit.name or
// the artist.name.
//
func (q ArtistCreditQuery) WhereCreditOrArtistNameMatches(name string) ArtistCreditQuery {
	q.builder = q.builder.Where(`
		lower(name) = lower(?)
		OR id IN (
			SELECT acn.artist_credit
			FROM artist_credit_name acn
			JOIN artist ON artist.id = acn.artist
			WHERE lower(artist.name) = lower(?)
		)
	`, name, name)
	return q
}

// ArtistCreditMap returns a mapping of ArtistCredit IDs to ArtistCredit structs
func ArtistCreditMap(db DB, ids []int64) (credits map[int64]*ArtistCredit, err error) {
	credits = make(map[int64]*ArtistCredit)
	results, err := ArtistCredits(db).Where("id IN (?)", ids).WithAssociations().All()
	for _, credit := range results {
		credits[credit.ID] = credit
	}
	return
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
