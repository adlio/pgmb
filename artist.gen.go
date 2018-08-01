// Code generated by go generate; DO NOT EDIT.
package pgmb

import (
	"strings"
	"github.com/pkg/errors"
)

// ArtistQuery is a queryer for Artist data
type ArtistQuery struct {
	db      DB
	builder SelectBuilder
	processors []ArtistCollectionProcessor
}

// Artists is the constructor for ArtistQuery
func Artists(db DB, columns ...string) ArtistQuery {

	var selectClause string
	if len(columns) > 0 {
		selectClause = strings.Join(columns, ", ")
	} else {
		selectClause = "id, gid, name, sort_name, begin_date_year, end_date_year"
	}

	q := ArtistQuery{
		db:      db,
		builder: SelectBuilder{}.Select(selectClause).From("artist"),
	}
	return q
}

// ArtistCollection is a slice of Artist
type ArtistCollection []*Artist

// ArtistCollectionProcessor is a function which modifies each element in a ArtistCollection
// (typically by populting additional data on it)
type ArtistCollectionProcessor func(DB, ArtistCollection) error

// Select adjusts the columns returned from the query
func (q ArtistQuery) Select(columns string) ArtistQuery {
	q.builder = q.builder.Select(columns)
	return q
}

// Where adds an additional where clause to the query
func (q ArtistQuery) Where(cmd string, args ...interface{}) ArtistQuery {
	q.builder = q.builder.Where(cmd, args...)
	return q
}

// OrderBy adjusts the ordering criteria for the query
func (q ArtistQuery) OrderBy(cmd string) ArtistQuery {
	q.builder = q.builder.OrderBy(cmd)
	return q
}

// All returns all results from the query
func (q ArtistQuery) All() (results ArtistCollection, err error) {
	results = make(ArtistCollection, 0)
	err = Select(q.db, &results, q.builder)
	if err != nil {
		err = errors.Wrap(err, "ArtistQuery.All() failed to populate ArtistCollection.")
		return
	}
	for _, f := range q.processors {
		err = f(q.db, results)
		if err != nil {
			err = errors.Wrap(err, "Failed to run processor over ArtistCollection.")
		}
	}
	return
}

// One returns a single result from the query
func (q ArtistQuery) One() (result *Artist, err error) {
	results := make(ArtistCollection, 0, 1)
	err = Select(q.db, &results, q.builder)
	if err != nil {
		err = errors.Wrap(err, "ArtistQuery.One() failed to populate initial result set.")
		return
	}

	for _, f := range q.processors {
		err = f(q.db, results)
		if err != nil {
			err = errors.Wrap(err, "Failed to run processor over ArtistCollection.")
			return
		}
	}

	if len(results) > 0 {
		result = results[0]
	}

	return result, nil
}