package models

import (
	"time"

	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type ArtistAlias struct {
	ID       int64
	ArtistID int64 `db:"artist"`
	Name     string
	SortName string `db:"sort_name"`
	Type     ArtistAliasType
}

type ArtistAliasType struct {
	ID   int64
	Name string
}

type Artist struct {
	ID            int64
	GID           uuid.UUID
	Name          string
	SortName      string `db:"sort_name"`
	Aliases       []*ArtistAlias
	BeginDateYear *int64
	EndDateYear   *int64
	LastUpdated   time.Time
}

func (a *Artist) AddAlias(alias *ArtistAlias) {
	a.Aliases = append(a.Aliases, alias)
}

func FindArtistsNamed(db *sqlx.DB, name string) (artists []*Artist, err error) {
	artists = make([]*Artist, 0)

	sql := `
		SELECT
		id, name, sort_name, begin_date_year, end_date_year
		FROM artist
		WHERE artist.id IN (
			SELECT id
			FROM artist
			WHERE lower(name) % lower(?)
			UNION
			SELECT artist
			FROM artist_alias
			WHERE lower(name) % lower(?)
		)
		ORDER BY similarity(lower(name), lower(?)) DESC`

	err = db.Select(&artists, db.Rebind(sql), name, name, name)
	if err != nil {
		return
	}

	err = LoadArtistAliases(db, artists)
	return
}

// Given a slice of Artist, this function loads and attaches
// all ArtistAliases from the database via a single SQL query.
// This function is designed to operate on < 100 records of input.
//
func LoadArtistAliases(db *sqlx.DB, artists []*Artist) error {
	var err error
	var aliases []*ArtistAlias

	ids := make([]int64, len(artists))
	lu := make(map[int64]*Artist)

	// Collect artist IDs
	for i, artist := range artists {
		ids[i] = artist.ID
		lu[artist.ID] = artist
		// Ensure Aliases starts empty
		artist.Aliases = make([]*ArtistAlias, 0)
	}

	// Fetch all aliases based on those IDs
	sql := `
		SELECT id, artist, name, sort_name
		FROM artist_alias
		WHERE artist IN (?)
	`
	sql, args, err := sqlx.In(sql, ids)
	if err != nil {
		return err
	}
	err = db.Select(&aliases, db.Rebind(sql), args...)
	if err != nil {
		return err
	}

	// Attach Alias objects to the original Artists
	for _, alias := range aliases {
		if artist, ok := lu[alias.ArtistID]; ok {
			artist.AddAlias(alias)
		}
	}

	return err
}
