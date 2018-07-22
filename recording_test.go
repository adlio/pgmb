package pgmb

import "testing"

func TestFindRecordingsByFuzzyName(t *testing.T) {
	recordings, err := FindRecordings(TESTDB, RecordingName("I'm a Cool Girl"))
	if err != nil {
		t.Error(err)
	}
	if len(recordings) < 1 {
		t.Fatalf("No results")
	}
}
