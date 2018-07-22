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

// ArtistCreditQuery builds the default query for the artist_credit table
//
func ArtistCreditQuery() sq.SelectBuilder {
	return Query().
		Select("id, name, artist_count, ref_count").
		From("artist_credit").
		Limit(1000)
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

// FindArtistCredits retrieves a slice of ArtistCredit which
// match the supplied criteria.
//
func FindArtistCredits(db DB, clauses ...QueryFunc) (credits []*ArtistCredit, err error) {
	credits = make([]*ArtistCredit, 0)
	err = Select(db, &credits, ArtistCreditQuery(), clauses...)
	return
}
