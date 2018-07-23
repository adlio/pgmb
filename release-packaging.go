package pgmb

import (
	sq "github.com/Masterminds/squirrel"
	uuid "github.com/satori/go.uuid"
)

type ReleasePackaging struct {
	ID          int64
	GID         uuid.UUID
	Name        string
	ChildOrder  int64 `db:"child_order"`
	Description *string
}

func ReleasePackagingQuery() sq.SelectBuilder {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("id, gid, name, child_order, description").
		From("release_packaging")
}

func ReleasePackagingMap(db DB) (packagings map[int64]*ReleasePackaging, err error) {
	packagings = make(map[int64]*ReleasePackaging)
	rs := make([]*ReleasePackaging, 0)
	err = Select(db, &rs, ReleasePackagingQuery())
	for _, status := range rs {
		packagings[status.ID] = status
	}
	return
}
