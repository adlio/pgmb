package pgmb

import (
	"testing"
)

func TestFindReleaseEvents(t *testing.T) {
	events, err := ReleaseEvents(TESTDB).Where("release IN (?)", []int64{439665, 1991688, 1991696}).All()
	if err != nil {
		t.Error(err)
	}

	if len(events) < 1 {
		t.Error("No results from FindReleaseEvents")
	}
}
