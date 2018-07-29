package pgmb

import (
	"fmt"
	"testing"
)

func TestFindDeepTracksByNameAndArtist(t *testing.T) {
	madonnas, err := FindArtistCredits(TESTDB, Where("lower(artist_credit.name) = lower(?)", "Madonna"))
	if err != nil {
		t.Error(err)
	}

	ids := make([]int64, len(madonnas)+1)
	for i, ac := range madonnas {
		ids[i] = ac.ID
	}

	tracks, err := FindDeepTracks(TESTDB, Where("lower(track.name) = lower(?) AND track.artist_credit IN (?)", "Like a Virgin", ids))
	if err != nil {
		t.Error(err)
	}
	if len(tracks) < 1 {
		t.Fatalf("No results")
	}
	for _, track := range tracks {
		fmt.Printf("%s %d - %s by %s on Release Group '%s' by '%s' [%s]\n", track.GID, track.Position, track.Name, track.ArtistCredit.Name, track.ReleaseGroup.Name, track.ReleaseGroup.ArtistCredit.Name, track.Release.EarliestReleaseDate().Format("2006-01-02"))
	}
}
