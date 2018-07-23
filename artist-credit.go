package pgmb

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

// ArtistCredit represents an entry in the MusicBrainz
// artist_credit table.
type ArtistCredit struct {
	ID          int64
	Name        string
	ArtistCount int64 `db:"artist_count"`
	RefCount    int64 `db:"ref_count"`
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
func FindArtistCredits(db DB, clauses ...QueryFunc) (credits []*ArtistCredit, err error) {
	credits = make([]*ArtistCredit, 0)
	err = Select(db, &credits, ArtistCreditQuery(), clauses...)
	return
}

// ArtistCreditMap returns a mapping of ArtistCredit IDs to ArtistCredit structs
func ArtistCreditMap(db DB, ids []int64) (credits map[int64]*ArtistCredit, err error) {
	credits = make(map[int64]*ArtistCredit)
	results, err := FindArtistCredits(db, IDIn(ids))
	for _, credit := range results {
		credits[credit.ID] = credit
	}
	return
}

// ArtistCreditQuery builds the default query for the artist_credit table
//
func ArtistCreditQuery() sq.SelectBuilder {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("id, name, artist_count, ref_count").
		From("artist_credit")
}

// ArtistCreditIn builds a QueryFunc for filtering a resultset to a specific
// list of supplied ArtistCredits.
//
func ArtistCreditIn(acs []*ArtistCredit) QueryFunc {
	return func(b sq.SelectBuilder) sq.SelectBuilder {
		ids := make([]interface{}, len(acs))
		for i, ac := range acs {
			ids[i] = ac.ID
		}
		sql, args, _ := sqlx.In("artist_credit IN (?)", ids)
		return b.Where(sql, args...)
	}
}
