package pgmb

import (
	sq "github.com/Masterminds/squirrel"
	uuid "github.com/satori/go.uuid"
)

type ReleaseStatus struct {
	ID          int64
	GID         uuid.UUID
	Name        string
	ChildOrder  int64 `db:"child_order"`
	Description string
}

func ReleaseStatusQuery() sq.SelectBuilder {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("id, gid, name, child_order, description").
		From("release_status")
}

func ReleaseStatusMap(db DB) (statuses map[int64]*ReleaseStatus, err error) {
	statuses = make(map[int64]*ReleaseStatus)
	rs := make([]*ReleaseStatus, 0)
	err = Select(db, &rs, ReleaseStatusQuery())
	for _, status := range rs {
		statuses[status.ID] = status
	}
	return
}
