package pgmb

import (
	"bytes"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// QueryFunc allows chaining of squirrel statements
type QueryFunc func(sq.SelectBuilder) sq.SelectBuilder

// Where is a wrapper to a Squirrel Where()
func Where(cmd string, args ...interface{}) QueryFunc {
	return func(b sq.SelectBuilder) sq.SelectBuilder {
		sql, args, err := sqlx.In(cmd, args...)
		if err != nil {
			log.Fatal(err)
		}
		return b.Where(sql, args...)
	}
}

// SelectBuilder blah
type SelectBuilder struct {
	selectClause string
	fromClause   string
	whereClauses []string
	orderClauses []string
	limitClause  string
	args         []interface{}
}

// Select blah
func (s SelectBuilder) Select(clause string) SelectBuilder {
	s.selectClause = clause
	s.whereClauses = make([]string, 0)
	s.args = make([]interface{}, 0)
	return s
}

// From blah
func (s SelectBuilder) From(clause string) SelectBuilder {
	s.fromClause = clause
	return s
}

// Where blah
func (s SelectBuilder) Where(clause string, args ...interface{}) SelectBuilder {
	s.whereClauses = append(s.whereClauses, clause)
	s.args = append(s.args, args...)
	return s
}

// OrderBy blah
func (s SelectBuilder) OrderBy(clause string) SelectBuilder {
	s.orderClauses = append(s.orderClauses, clause)
	return s
}

// Limit blah
func (s SelectBuilder) Limit(n int) SelectBuilder {
	s.limitClause = fmt.Sprintf("LIMIT %d", n)
	return s
}

func (s SelectBuilder) ToSql() (sql string, args []interface{}, err error) {
	sql, args = s.ToSQL()
	return sql, args, nil
}

// ToSQL blah
func (s SelectBuilder) ToSQL() (sql string, args []interface{}) {
	b := bytes.Buffer{}
	b.WriteString("SELECT ")
	b.WriteString(s.selectClause)
	b.WriteString(" FROM ")
	b.WriteString(s.fromClause)
	if len(s.whereClauses) > 0 {
		b.WriteString(" WHERE ")
	}
	for i, c := range s.whereClauses {
		b.WriteString(c)
		if i > 0 && i < len(s.whereClauses)-1 {
			b.WriteString(" AND ")
		}
	}
	if len(s.orderClauses) > 0 {
		b.WriteString(" ORDER BY")
	}
	for i, c := range s.orderClauses {
		b.WriteString(" ")
		b.WriteString(c)
		if i < len(s.orderClauses)-1 {
			b.WriteString(",")
		}
	}
	if s.limitClause != "" {
		b.WriteString(" ")
		b.WriteString(s.limitClause)
	}

	var err error
	sql = b.String()

	sql, args, err = sqlx.In(sql, s.args...)
	if err != nil {
		err = errors.Wrapf(err, "Couldn't run sqlx.In with '%s' and these args: \n%s", sql, s.args)
		panic(err)
	}

	return sql, args
}
