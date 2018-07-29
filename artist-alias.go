package pgmb

import sq "github.com/Masterminds/squirrel"

// ArtistAlias holds an alternative name for an Artist
type ArtistAlias struct {
	ID       int64
	ArtistID int64 `db:"artist"`
	Name     string
	SortName string `db:"sort_name"`
}

func FindArtistAliases(db DB, clauses ...QueryFunc) (aliases []*ArtistAlias, err error) {
	aliases = make([]*ArtistAlias, 0)
	err = Select(db, &aliases, ArtistAliasQuery(), clauses...)
	if err != nil {
		return
	}
	return
}

// ArtistAliasQuery is the base query for working with artist_alias data
func ArtistAliasQuery() sq.SelectBuilder {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("id, artist, name, sort_name").
		From("artist_alias")
}