// Code generated by go generate; DO NOT EDIT.
package pgmb

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

// DeepTrackQueryFunc can be chained together to modify a DeepTrackQuery
type DeepTrackQueryFunc func(DeepTrackQuery) DeepTrackQuery

// DeepTrackQuery is a queryer for DeepTrack data
type DeepTrackQuery struct {
	db         DB
	builder    sq.SelectBuilder
	processors []DeepTrackCollectionProcessor
}

// DeepTracks is the constructor for DeepTrackQuery
func DeepTracks(db DB) DeepTrackQuery {
	q := DeepTrackQuery{
		db:      db,
		builder: DeepTrackSelect(),
	}
	return q
}

// DeepTrackCollection is a slice of DeepTrack
type DeepTrackCollection []*DeepTrack

// DeepTrackCollectionProcessor is a function which modifies each element in a DeepTrackCollection
// (typically by populting additional data on it)
type DeepTrackCollectionProcessor func(DB, DeepTrackCollection) error

// From sets the table being queried
func (q DeepTrackQuery) From(name string) DeepTrackQuery {
	q.builder = q.builder.From(name)
	return q
}

// Where adds an additional where clause to the query
func (q DeepTrackQuery) Where(cmd string, args ...interface{}) DeepTrackQuery {
	q.builder = q.builder.Where(cmd, args...)
	return q
}

// OrderBy adjusts the ordering criteria for the query
func (q DeepTrackQuery) OrderBy(cmd string) DeepTrackQuery {
	q.builder = q.builder.OrderBy(cmd)
	return q
}

// All returns all results from the query
func (q DeepTrackQuery) All() (results DeepTrackCollection, err error) {
	results = make(DeepTrackCollection, 0)
	err = Select(q.db, &results, q.builder)
	if err != nil {
		err = errors.Wrap(err, "DeepTrackQuery.All() failed to populate DeepTrackCollection.")
		return
	}
	for _, f := range q.processors {
		err = f(q.db, results)
		if err != nil {
			err = errors.Wrap(err, "Failed to run processor over DeepTrackCollection.")
		}
	}
	return
}

// One returns a single result from the query
func (q DeepTrackQuery) One() (result *DeepTrack, err error) {
	results := make(DeepTrackCollection, 0, 1)
	err = Select(q.db, &results, q.builder)
	if err != nil {
		err = errors.Wrap(err, "DeepTrackQuery.One() failed to populate initial result set.")
		return
	}

	for _, f := range q.processors {
		err = f(q.db, results)
		if err != nil {
			err = errors.Wrap(err, "Failed to run processor over DeepTrackCollection.")
			return
		}
	}

	if len(results) > 0 {
		result = results[0]
	}

	return result, nil
}
