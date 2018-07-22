package pgmb

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
func (ac *ArtistCredit) IsVariousArtists() bool {
	return ac.ID == 1
}
