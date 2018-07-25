package pgmb

import (
	"fmt"
	"testing"
)

func TestFindRecordingsByFuzzyNameAndArtist(t *testing.T) {
	t.Skipf("TODO: Speed this up. Too slow.")
	madonnas, err := FindArtistCredits(TESTDB, FuzzyNamed("Madonna"))
	if err != nil {
		t.Error(err)
	}

	recordings, err := FindRecordings(TESTDB, FuzzyNamed("Like a Virgin"), ArtistCreditIn(madonnas))
	if err != nil {
		t.Error(err)
	}
	if len(recordings) < 1 {
		t.Fatalf("No results")
	}
}

func TestFindRecordingsByNameAndArtist(t *testing.T) {
	madonnas, err := FindArtistCredits(TESTDB, Named("Madonna"))
	if err != nil {
		t.Error(err)
	}

	recordings, err := FindRecordings(TESTDB, Named("Like a Virgin"), ArtistCreditIn(madonnas))
	if err != nil {
		t.Error(err)
	}
	if len(recordings) < 1 {
		t.Fatalf("No results")
	}

	for _, recording := range recordings {
		fmt.Printf("%s %s by %s\n", recording.GID, recording.Name, recording.ArtistCredit.Name)
	}
}
