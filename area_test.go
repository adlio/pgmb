package pgmb

import (
	"testing"
)

func TestFindAreas(t *testing.T) {
	areas, err := Areas(TESTDB).Where("id IN (?)", []int64{222}).All()
	if err != nil {
		t.Error(err)
	}

	if len(areas) < 1 {
		t.Error("No results from Areas()")
	}
	if areas[0].Name != "United States" {
		t.Error("Should have found United States.")
	}
}
