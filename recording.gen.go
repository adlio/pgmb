// Code generated by go generate; DO NOT EDIT.
package pgmb

import (
	"strings"
	"github.com/pkg/errors"
)

// RecordingQuery is a queryer for Recording data
type RecordingQuery struct {
	db      DB
	builder SelectBuilder
	processors []RecordingCollectionProcessor
}

// Recordings is the constructor for RecordingQuery
func Recordings(db DB, columns ...string) RecordingQuery {

	var selectClause string
	if len(columns) > 0 {
		selectClause = strings.Join(columns, ", ")
	} else {
		selectClause = "id, gid, name, artist_credit, length, comment"
	}

	q := RecordingQuery{
		db:      db,
		builder: SelectBuilder{}.Select(selectClause).From("recording"),
	}
	return q
}

// RecordingCollection is a slice of Recording
type RecordingCollection []*Recording

// RecordingCollectionProcessor is a function which modifies each element in a RecordingCollection
// (typically by populting additional data on it)
type RecordingCollectionProcessor func(DB, RecordingCollection) error

// Select adjusts the columns returned from the query
func (q RecordingQuery) Select(columns string) RecordingQuery {
	q.builder = q.builder.Select(columns)
	return q
}

// Where adds an additional where clause to the query
func (q RecordingQuery) Where(cmd string, args ...interface{}) RecordingQuery {
	q.builder = q.builder.Where(cmd, args...)
	return q
}

// OrderBy adjusts the ordering criteria for the query
func (q RecordingQuery) OrderBy(cmd string) RecordingQuery {
	q.builder = q.builder.OrderBy(cmd)
	return q
}

// All returns all results from the query
func (q RecordingQuery) All() (results RecordingCollection, err error) {
	results = make(RecordingCollection, 0)
	err = Select(q.db, &results, q.builder)
	if err != nil {
		err = errors.Wrap(err, "RecordingQuery.All() failed to populate RecordingCollection.")
		return
	}
	for _, f := range q.processors {
		err = f(q.db, results)
		if err != nil {
			err = errors.Wrap(err, "Failed to run processor over RecordingCollection.")
		}
	}
	return
}

// One returns a single result from the query
func (q RecordingQuery) One() (result *Recording, err error) {
	results := make(RecordingCollection, 0, 1)
	err = Select(q.db, &results, q.builder)
	if err != nil {
		err = errors.Wrap(err, "RecordingQuery.One() failed to populate initial result set.")
		return
	}

	for _, f := range q.processors {
		err = f(q.db, results)
		if err != nil {
			err = errors.Wrap(err, "Failed to run processor over RecordingCollection.")
			return
		}
	}

	if len(results) > 0 {
		result = results[0]
	}

	return result, nil
}