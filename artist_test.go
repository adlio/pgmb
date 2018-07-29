package pgmb

import (
	"testing"

	"github.com/satori/go.uuid"

	_ "github.com/lib/pq"
)

func TestGetArtistsByFuzzyNameOrAlias(t *testing.T) {
	artists, err := FindArtists(TESTDB, ArtistFuzzyNameOrAlias("Crosby Stills Nash"))
	if err != nil {
		t.Error(err)
	}
	if len(artists) < 1 {
		t.Fatalf("No results")
	}
}

func TestGetArtistsByID(t *testing.T) {
	uuid, err := uuid.FromString("79239441-bfd5-4981-a70c-55c3f15c1287")
	if err != nil {
		t.Error(err)
	}
	artist, err := GetArtist(TESTDB, Where("gid = ?", uuid))
	if err != nil {
		t.Error(err)
	}
	if artist.Name != "Madonna" {
		t.Errorf("Expected 'Madonna', got '%s'.", artist.Name)
	}
	if *artist.BeginDateYear != 1958 {
		t.Errorf("Expected begin_date_year = 1958, got %d", artist.BeginDateYear)
	}
}
