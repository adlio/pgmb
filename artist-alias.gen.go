// Code generated by go generate; DO NOT EDIT.
package pgmb

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

// ArtistAliasQueryFunc can be chained together to modify a ArtistAliasQuery
type ArtistAliasQueryFunc func(ArtistAliasQuery) ArtistAliasQuery

// ArtistAliasQuery is a queryer for ArtistAlias data
type ArtistAliasQuery struct {
	db         DB
	builder    sq.SelectBuilder
	processors []ArtistAliasCollectionProcessor
}

// ArtistAliases is the constructor for ArtistAliasQuery
func ArtistAliases(db DB) ArtistAliasQuery {
	q := ArtistAliasQuery{
		db:      db,
		builder: ArtistAliasSelect(),
	}
	return q
}

// ArtistAliasCollection is a slice of ArtistAlias
type ArtistAliasCollection []*ArtistAlias

// ArtistAliasCollectionProcessor is a function which modifies each element in a ArtistAliasCollection
// (typically by populting additional data on it)
type ArtistAliasCollectionProcessor func(DB, ArtistAliasCollection) error

// Select can be used to replace ArtistAliasSelect() with a different squirrel.SelectBuilder
// to pull different fields or join data differently.
func (q ArtistAliasQuery) Select(b sq.SelectBuilder) ArtistAliasQuery {
	q.builder = b
	return q
}

// Where adds an additional where clause to the query
func (q ArtistAliasQuery) Where(cmd string, args ...interface{}) ArtistAliasQuery {
	q.builder = q.builder.Where(cmd, args...)
	return q
}

// OrderBy adjusts the ordering criteria for the query
func (q ArtistAliasQuery) OrderBy(cmd string) ArtistAliasQuery {
	q.builder = q.builder.OrderBy(cmd)
	return q
}

// All returns all results from the query
func (q ArtistAliasQuery) All() (results ArtistAliasCollection, err error) {
	results = make(ArtistAliasCollection, 0)
	err = Select(q.db, &results, q.builder)
	if err != nil {
		err = errors.Wrap(err, "ArtistAliasQuery.All() failed to populate ArtistAliasCollection.")
		return
	}
	for _, f := range q.processors {
		err = f(q.db, results)
		if err != nil {
			err = errors.Wrap(err, "Failed to run processor over ArtistAliasCollection.")
		}
	}
	return
}

// One returns a single result from the query
func (q ArtistAliasQuery) One() (result *ArtistAlias, err error) {
	results := make(ArtistAliasCollection, 0, 1)
	err = Select(q.db, &results, q.builder)
	if err != nil {
		err = errors.Wrap(err, "ArtistAliasQuery.One() failed to populate initial result set.")
		return
	}

	for _, f := range q.processors {
		err = f(q.db, results)
		if err != nil {
			err = errors.Wrap(err, "Failed to run processor over ArtistAliasCollection.")
			return
		}
	}

	if len(results) > 0 {
		result = results[0]
	}

	return result, nil
}
