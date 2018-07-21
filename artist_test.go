package pgmb

import (
	"testing"

	_ "github.com/lib/pq"
)

func TestGetArtistByID(t *testing.T) {
	artists, err := FindArtistsNamed(DB, "Crosby Stills Nash")
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
