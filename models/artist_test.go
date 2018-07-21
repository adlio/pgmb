package models

import (
	"log"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func init() {
	var err error
	DB, err = sqlx.Open("postgres", "postgres://mgr_development:e7c87b7607029da0059f4fc7292b81a5@127.0.0.1/mgr_development?sslmode=disable&search_path=musicbrainz,public")
	DB.MapperFunc(ToSnakeCase)
	if err != nil {
		log.Fatal(err)
	}
}
func TestGetArtistByID(t *testing.T) {
	artists, err := FindArtistsNamed(DB, "Crosby Stills Nahs")
	if err != nil {
		t.Error(err)
	}
	if len(artists) < 1 {
		t.Fatalf("No results")
	}
	artist := artists[0]
	if artist.Name != "Crosby, Stills & Nash" {
		t.Errorf("Expected name 'Massive Attack' got '%s'", artist.Name)
	}
	if len(artist.Aliases) != 6 {
		t.Errorf("Expected 6 aliases, got %d", len(artist.Aliases))
	}
}
