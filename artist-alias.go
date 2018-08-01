package pgmb

// ArtistAlias holds an alternative name for an Artist
type ArtistAlias struct {
	ID       int64
	ArtistID int64 `db:"artist"`
	Name     string
	SortName string `db:"sort_name"`
}

// UniqueArtistIDs pulls the unique artist IDs from the slice of ArtistAlias
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
