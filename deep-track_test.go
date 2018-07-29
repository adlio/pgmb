package pgmb

import (
	"fmt"
	"testing"
)

func TestFindDeepTracksByRecordingGID(t *testing.T) {

	tracks, err := FindDeepTracks(TESTDB, Where("recording.gid = ?", "c0d4a5c9-6992-4b99-bcf2-388dd6be19c0"))
	if err != nil {
		t.Error(err)
	}
	for _, track := range tracks {
		fmt.Printf("%s from %s by %s (%d artists) has %d artists: %s\n", track.Name, track.ReleaseGroup.Name, track.ReleaseGroup.ArtistCredit.Name, track.ReleaseGroup.ArtistCredit.ArtistCount, len(track.ArtistCredit.Artists), track.ArtistCredit.Name)
	}
}
func TestFindDeepTracksByNameAndArtist(t *testing.T) {
	madonnas, err := FindArtistCredits(TESTDB, Where("lower(artist_credit.name) = lower(?)", "Madonna"))
	if err != nil {
		t.Error(err)
	}

	tracks, err := FindDeepTracks(TESTDB, Where("lower(track.name) = lower(?) AND track.artist_credit IN (?)", "Like a Virgin", madonnas.MBIDs()))
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
