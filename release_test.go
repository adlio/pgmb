package pgmb

import (
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestWhereReleaseIncludesRecording(t *testing.T) {
	rid, _ := uuid.FromString("bb883fd9-ab17-434f-b336-9469a2b4f363")
	releases, err := Releases(TESTDB).Where(`
		EXISTS (
			SELECT track.id
			FROM track
			JOIN medium ON medium.id = track.medium
			JOIN recording ON recording.id = track.recording
			WHERE recording.gid = ?
			AND medium.release = release.id
		)
	`, rid).WithAssociations().All()
	if err != nil {
		t.Fatal(err)
	}

	if len(releases) < 10 {
		t.Errorf("Expected at least 10 releases for 'bb883fd9-ab17-434f-b336-9469a2b4f363'. Got %d", len(releases))
	}
}
