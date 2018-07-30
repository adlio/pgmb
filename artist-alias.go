package pgmb

import sq "github.com/Masterminds/squirrel"

// ArtistAlias holds an alternative name for an Artist
type ArtistAlias struct {
	ID       int64
	ArtistID int64 `db:"artist"`
	Name     string
	SortName string `db:"sort_name"`
}

// ArtistAliasCollection is an alias for a slice of ArtistAlias
type ArtistAliasCollection []*ArtistAlias

func (c ArtistAliasCollection) UniqueArtistIDs() []int64 {
	idMap := make(map[int64]bool)
	for _, a := range c {
		idMap[a.ArtistID] = true
	}
	ids := make([]int64, 0, len(idMap))
	for id := range idMap {
		ids = append(ids, id)
	}
	return ids
}

// FindArtistAliases retrieves a slice of ArtistAlias based on a dynamic query
//
func FindArtistAliases(db DB, clauses ...QueryFunc) (aliases ArtistAliasCollection, err error) {
	aliases = make(ArtistAliasCollection, 0)
	err = Select(db, &aliases, ArtistAliasQuery(), clauses...)
	return
}

// ArtistAliasQuery is the base query for working with artist_alias data
func ArtistAliasQuery() sq.SelectBuilder {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("id, artist, name, sort_name").
		From("artist_alias")
}
