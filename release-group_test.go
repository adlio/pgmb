package pgmb

import (
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestReleaseGroup40OzToFreedom(t *testing.T) {
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
	if len(releaseGroup.SecondaryTypes) != 0 {
		t.Errorf("Expected no secondary types. Got %d.", len(releaseGroup.SecondaryTypes))
	}
}

func TestReleaseGroupLiveWideOpen(t *testing.T) {
	rid, _ := uuid.FromString("12990b55-d95f-3b5c-9ae6-d95b0c8c44d4")
	releaseGroup, err := GetReleaseGroup(TESTDB, WithGID(rid))
	if err != nil {
		t.Fatal(err)
	}
	if releaseGroup.Name != "Live Wide Open" {
		t.Errorf("Expected 'Live Wide Open', got '%s'.", releaseGroup.Name)
	}

	if releaseGroup.ArtistCredit == nil {
		t.Fatal("ReleaseGroup.ArtistCredit shouldn't be nil")
	}
	if releaseGroup.ArtistCredit.Name != "Martin Sexton" {
		t.Errorf("Expected 'Martin Sexton', got '%s'", releaseGroup.ArtistCredit.Name)
	}
	if releaseGroup.Type == nil {
		t.Fatal("ReleaseGroup.Type shouldn't be nil")
	}
	if releaseGroup.Type.ID != 1 || releaseGroup.Type.Name != "Album" {
		t.Errorf("Expected '1-Album', got '%d-%s'.", releaseGroup.Type.ID, releaseGroup.Type.Name)
	}
	if len(releaseGroup.SecondaryTypes) != 1 {
		t.Errorf("Expected 1 secondary type Got %d.", len(releaseGroup.SecondaryTypes))
	}
	if releaseGroup.SecondaryTypes[0].ID != 6 || releaseGroup.SecondaryTypes[0].Name != "Live" {
		t.Errorf("Expected 6-Live, got %d-%s.", releaseGroup.SecondaryTypes[0].ID, releaseGroup.SecondaryTypes[0].Name)
	}
}
