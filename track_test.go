package pgmb

import (
	"testing"
)

func TestFindTracksByNameAndArtist(t *testing.T) {
	/*
		madonnas, err := FindArtistCredits(TESTDB, Where("lower(name) = lower(?)", "Madonna"))
		if err != nil {
			t.Error(err)
		}

		tracks, err := FindTracks(TESTDB, Where("lower(name) = lower(?) AND artist_credit IN (?)", "Like a Virgin", madonnas.MBIDs()))
		if err != nil {
			t.Error(err)
		}
		if len(tracks) < 1 {
			t.Fatalf("No results")
		}
		/*
			for _, track := range tracks {
				fmt.Printf("%s %d - %s by %s\n", track.GID, track.Position, track.Name, track.ArtistCredit.Name)
			}
	*/
}
