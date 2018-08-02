package pgmb

import (
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

// ReleaseEvent represents an entry in the MusicBrainz
// release_event view (which is a UNION of the
// release_country and release_unknown_country)
//
type ReleaseEvent struct {
	ReleaseID int64  `db:"release"`
	CountryID *int64 `db:"country"`
	Country   *Area  `db:"-"`
	DateYear  *int64 `db:"date_year"`
	DateMonth *int64 `db:"date_month"`
	DateDay   *int64 `db:"date_day"`
}

// ReleaseEventSelect builds the default select for release data
func ReleaseEventSelect() sq.SelectBuilder {
	return sq.StatementBuilder.
		Select("release, country, date_year, date_month, date_day").
		From("release_event")
}

// Date converts the MusicBrainz-supplied YMD values into a
// time.Time
func (re *ReleaseEvent) Date() time.Time {
	var y int
	var m time.Month
	var d int
	if re.DateYear != nil {
		y = int(*re.DateYear)
	} else {
		return time.Time{}
	}
	if re.DateMonth != nil {
		m = time.Month(*re.DateMonth)
	} else {
		m = 1
	}
	if re.DateDay != nil {
		d = int(*re.DateDay)
	} else {
		d = 1
	}
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}

// ReleaseEventMap returns a mapping of Release IDs to slices of related ReleasEvents
//
func ReleaseEventMap(db DB, releaseIDs []int64) (events map[int64][]*ReleaseEvent, err error) {
	events = make(map[int64][]*ReleaseEvent)
	countries := make(map[int64]bool)

	// Collect each ReleaseEvent into a map, taking note of each country we see
	results, err := ReleaseEvents(db).Where("release IN (?)", releaseIDs).All()
	if err != nil {
		return events, errors.Wrap(err, "ReleaseEventMap failed when calling FindReleaseEvents")
	}
	for _, event := range results {
		if _, existed := events[event.ReleaseID]; !existed {
			events[event.ReleaseID] = make([]*ReleaseEvent, 0)
		}
		events[event.ReleaseID] = append(events[event.ReleaseID], event)
		if event.CountryID != nil {
			countries[*event.CountryID] = true
		}
	}

	// Fetch country Areas for each touched country, adding them to
	// the original map
	countryIDs := make([]int64, 0, len(countries))
	for id := range countries {
		countryIDs = append(countryIDs, id)
	}
	areaMap, err := AreaMap(db, countryIDs)
	if err != nil {
		return events, errors.Wrap(err, "ReleaseEventMap failed when building AreaMap")
	}
	for _, event := range results {
		if event.CountryID != nil {
			event.Country, _ = areaMap[*event.CountryID]
		}
	}

	return
}
