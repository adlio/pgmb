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

func (q ArtistQuery) WithAssociations() ArtistQuery {
	q.processors = append(q.processors, loadArtistAliases)
	return q
}

// ArtistMap returns a mapping of Artist IDs to Artist structs, including
// only the ArtistID which was supplied.
func ArtistMap(db DB, ids []int64) (artists map[int64]*Artist, err error) {
	artists = make(map[int64]*Artist)
	results, err := Artists(db).WithAssociations().Where("id IN (?)", ids).All()
	for _, artist := range results {
		artists[artist.ID] = artist
	}
	return
}

// loadArtistAliases lodas and attaches all ArtistAliases for the supplied
// slice of Artist via a single SQL query. This function is designed to operate
// on < 100 records of input.
//
func loadArtistAliases(db DB, artists ArtistCollection) error {
	ids := make([]int64, len(artists))
	lu := make(map[int64]*Artist)

	// Collect artist IDs
	for i, artist := range artists {
		ids[i] = artist.ID
		lu[artist.ID] = artist
		// Ensure Aliases starts empty
		artist.Aliases = make([]*ArtistAlias, 0)
	}

	aliases, err := ArtistAliases(db).Where("artist IN (?)", ids).All()
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
