package pgmb

import (
	"testing"

	_ "github.com/lib/pq"
)

func TestFindArtistAliases(t *testing.T) {
	aliases, err := ArtistAliases(TESTDB).Where("lower(name) = lower(?)", "Crosby Stills & Nash").All()
	if err != nil {
		t.Error(err)
	}
	if len(aliases) < 1 {
		t.Fatalf("No results")
	}
	if aliases[0].Name != "Crosby Stills & Nash" {
		t.Errorf("Got wrong first artist alias match '%s'", aliases[0].Name)
	}
	if aliases[0].ArtistID != 12103 {
		t.Errorf("Got wrong ArtistID: %d", aliases[0].ArtistID)
	}
}
