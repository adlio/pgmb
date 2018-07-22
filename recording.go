package pgmb

import (
	"time"

	"github.com/Masterminds/squirrel"

	uuid "github.com/satori/go.uuid"
)

type Recording struct {
	ID          int64
	GID         uuid.UUID
	Name        string
	Length      *int64
	Comment     string
	LastUpdated time.Time
}

type RecordingName string

func (rn RecordingName) Query(b squirrel.SelectBuilder) squirrel.SelectBuilder {
	return b.Where("lower(recording.name) % lower(?)", rn).
		OrderBy("similarity(lower(recording.name), lower($1))")
}

// FindRecordings returns recordings matching the supplied criteria
func FindRecordings(db DB, clauses ...QueryFunc) (recordings []*Recording, err error) {
	recordings = make([]*Recording, 0)
	q := Query().
		Select("id, gid, name, length, comment").
		From("recording")
	err = Find(db, &recordings, q, clauses...)
	return
}
