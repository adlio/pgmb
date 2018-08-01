package pgmb

import (
	"testing"
)

func TestFindRecordingsByFuzzyNameAndArtist(t *testing.T) {
	t.Skipf("TODO: Speed this up. Too slow.")
	madonnas, err := ArtistCredits(TESTDB).Where("lower(name) % lower(?)", "Madonna").All()
	if err != nil {
		t.Error(err)
	}

	recordings, err := Recordings(TESTDB).Where("lower(name) % lower(?) AND artist_credit IN (?)", "Like a Virgin", madonnas.MBIDs()).All()
	if err != nil {
		t.Error(err)
	}
	if len(recordings) < 1 {
		t.Fatalf("No results")
	}
}

func TestFindRecordingsByNameAndArtist(t *testing.T) {
	madonnas, err := ArtistCredits(TESTDB).Where("lower(name) = lower(?)", "Madonna").All()
	if err != nil {
		t.Error(err)
	}

	recordings, err := Recordings(TESTDB).Where("lower(name) = lower(?) AND artist_credit IN (?)", "Like a Virgin", madonnas.MBIDs()).All()
	if err != nil {
		t.Error(err)
	}
	if len(recordings) < 1 {
		t.Fatalf("No results")
	}

	/*
		for _, recording := range recordings {
			fmt.Printf("%s %s by %s\n", recording.GID, recording.Name, recording.ArtistCredit.Name)
		}
	*/
}
