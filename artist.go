package pgmb

import (
	"time"

	sq "github.com/Masterminds/squirrel"

	uuid "github.com/satori/go.uuid"
)

// Artist represents an entry in the artist table in
// the MusicBrainz database.
type Artist struct {
	ID            int64
	GID           uuid.UUID
	Name          string
	SortName      string `db:"sort_name"`
	Aliases       []*ArtistAlias
	BeginDateYear *int64 `db:"begin_date_year"`
	EndDateYear   *int64 `db:"end_date_year"`
	LastUpdated   time.Time
}

// IsVariousArtists indicates whether the Artist represents a
// "Various Artists" record.
//
func (a *Artist) IsVariousArtists() bool {
	return a.ID == 1
}

// ArtistFuzzyNameOrAlias returns a QueryFunc which matches artists
// whose name or alias names fuzzy-match the supplied string.
func ArtistFuzzyNameOrAlias(name string) QueryFunc {
	return func(b sq.SelectBuilder) sq.SelectBuilder {
		return b.Where(`
		artist.id IN (
			SELECT id
			FROM artist
			WHERE lower(name) % lower(?)
			UNION
			SELECT artist
			FROM artist_alias
			WHERE lower(name) % lower(?)
		)
	`, name, name)
	}
}

// GetArtist fetches a single artist matching the supplied criteria
//
func GetArtist(db DB, clauses ...QueryFunc) (*Artist, error) {
	var err error
	var artist *Artist
	artist = &Artist{}

	err = Get(db, artist, ArtistQuery().Limit(1), clauses...)
	if err != nil {
		return artist, err
	}
	err = loadArtistAliases(db, []*Artist{artist})
	return artist, err
}

// FindArtists retrieves a slice of Artist based on a dynamic query
//
func FindArtists(db DB, clauses ...QueryFunc) (artists []*Artist, err error) {
	artists = make([]*Artist, 0)
	err = Select(db, &artists, ArtistQuery(), clauses...)
	if err != nil {
		return
	}

	if len(artists) > 0 {
		err = loadArtistAliases(db, artists)
	}
	return
}

// ArtistMap returns a mapping of Artist IDs to Artist structs, including
// only the ArtistID which was supplied.
func ArtistMap(db DB, ids []int64) (artists map[int64]*Artist, err error) {
	artists = make(map[int64]*Artist)
	results, err := FindArtists(db, Where("id IN (?)", ids))
	for _, artist := range results {
		artists[artist.ID] = artist
	}
	return
}

// ArtistQuery builds the default query for working with the artist table
func ArtistQuery() sq.SelectBuilder {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("id, gid, name, sort_name, begin_date_year, end_date_year").
		From("artist")
}

// loadArtistAliases lodas and attaches all ArtistAliases for the supplied
// slice of Artist via a single SQL query. This function is designed to operate
// on < 100 records of input.
//
func loadArtistAliases(db DB, artists []*Artist) error {
	ids := make([]int64, len(artists))
	lu := make(map[int64]*Artist)

	// Collect artist IDs
	for i, artist := range artists {
		ids[i] = artist.ID
		lu[artist.ID] = artist
		// Ensure Aliases starts empty
		artist.Aliases = make([]*ArtistAlias, 0)
	}

	aliases, err := FindArtistAliases(db, Where("artist IN (?)", ids))
	if err != nil {
		return err
	}

	// Attach Alias objects to the original Artists
	for _, alias := range aliases {
		if artist, ok := lu[alias.ArtistID]; ok {
			artist.Aliases = append(artist.Aliases, alias)
		}
	}

	return err
}
