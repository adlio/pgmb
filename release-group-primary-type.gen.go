// Code generated by go generate; DO NOT EDIT.
package pgmb

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

// ReleaseGroupPrimaryTypeQueryFunc can be chained together to modify a ReleaseGroupPrimaryTypeQuery
type ReleaseGroupPrimaryTypeQueryFunc func(ReleaseGroupPrimaryTypeQuery) ReleaseGroupPrimaryTypeQuery

// ReleaseGroupPrimaryTypeQuery is a queryer for ReleaseGroupPrimaryType data
type ReleaseGroupPrimaryTypeQuery struct {
	db         DB
	builder    sq.SelectBuilder
	processors []ReleaseGroupPrimaryTypeCollectionProcessor
}

// ReleaseGroupPrimaryTypes is the constructor for ReleaseGroupPrimaryTypeQuery
func ReleaseGroupPrimaryTypes(db DB) ReleaseGroupPrimaryTypeQuery {
	q := ReleaseGroupPrimaryTypeQuery{
		db:      db,
		builder: ReleaseGroupPrimaryTypeSelect(),
	}
	return q
}

// ReleaseGroupPrimaryTypeCollection is a slice of ReleaseGroupPrimaryType
type ReleaseGroupPrimaryTypeCollection []*ReleaseGroupPrimaryType

// ReleaseGroupPrimaryTypeCollectionProcessor is a function which modifies each element in a ReleaseGroupPrimaryTypeCollection
// (typically by populting additional data on it)
type ReleaseGroupPrimaryTypeCollectionProcessor func(DB, ReleaseGroupPrimaryTypeCollection) error

// From sets the table being queried
func (q ReleaseGroupPrimaryTypeQuery) From(name string) ReleaseGroupPrimaryTypeQuery {
	q.builder = q.builder.From(name)
	return q
}

// Where adds an additional where clause to the query
func (q ReleaseGroupPrimaryTypeQuery) Where(cmd string, args ...interface{}) ReleaseGroupPrimaryTypeQuery {
	q.builder = q.builder.Where(cmd, args...)
	return q
}

// OrderBy adjusts the ordering criteria for the query
func (q ReleaseGroupPrimaryTypeQuery) OrderBy(cmd string) ReleaseGroupPrimaryTypeQuery {
	q.builder = q.builder.OrderBy(cmd)
	return q
}

// All returns all results from the query
func (q ReleaseGroupPrimaryTypeQuery) All() (results ReleaseGroupPrimaryTypeCollection, err error) {
	results = make(ReleaseGroupPrimaryTypeCollection, 0)
	err = Select(q.db, &results, q.builder)
	if err != nil {
		err = errors.Wrap(err, "ReleaseGroupPrimaryTypeQuery.All() failed to populate ReleaseGroupPrimaryTypeCollection.")
		return
	}
	for _, f := range q.processors {
		err = f(q.db, results)
		if err != nil {
			err = errors.Wrap(err, "Failed to run processor over ReleaseGroupPrimaryTypeCollection.")
		}
	}
	return
}

// One returns a single result from the query
func (q ReleaseGroupPrimaryTypeQuery) One() (result *ReleaseGroupPrimaryType, err error) {
	results := make(ReleaseGroupPrimaryTypeCollection, 0, 1)
	err = Select(q.db, &results, q.builder)
	if err != nil {
		err = errors.Wrap(err, "ReleaseGroupPrimaryTypeQuery.One() failed to populate initial result set.")
		return
	}

	for _, f := range q.processors {
		err = f(q.db, results)
		if err != nil {
			err = errors.Wrap(err, "Failed to run processor over ReleaseGroupPrimaryTypeCollection.")
			return
		}
	}

	if len(results) > 0 {
		result = results[0]
	}

	return result, nil
}
