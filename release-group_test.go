package pgmb

import (
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestGetReleaseGroup(t *testing.T) {
	rid, _ := uuid.FromString("8e874f9b-7630-3a4c-baf4-73e0e553b2b4")
	releaseGroup, err := GetReleaseGroup(TESTDB, WithGID(rid))
	if err != nil {
		t.Fatal(err)
	}
	if releaseGroup.Name != "40 oz. to Freedom" {
		t.Errorf("Expected '40 oz. to Freedom'', got '%s'.", releaseGroup.Name)
	}

	if releaseGroup.ArtistCredit == nil {
		t.Fatal("ReleaseGroup.ArtistCredit shouldn't be nil")
	}
	if releaseGroup.ArtistCredit.Name != "Sublime" {
		t.Errorf("Expected 'Sublime', got '%s'", releaseGroup.ArtistCredit.Name)
	}
	if releaseGroup.Type == nil {
		t.Fatal("ReleaseGroup.Type shouldn't be nil")
	}
	if releaseGroup.Type.ID != 1 || releaseGroup.Type.Name != "Album" {
		t.Errorf("Expected '1-Album', got '%d-%s'.", releaseGroup.Type.ID, releaseGroup.Type.Name)
	}
}
