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

// FindArtistCredits retrieves a slice of ArtistCredit which
// match the supplied criteria.
//
func FindArtistCredits(db DB, criteria ...Queryer) (credits []*ArtistCredit, err error) {
	credits = make([]*ArtistCredit, 0)
	q := Query().
		Select("id, name, artist_count, ref_count").
		From("artist_credit").
		Limit(1000)
	err = Find(db, &credits, q, criteria...)
	return
}
