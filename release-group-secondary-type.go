package pgmb

import uuid "github.com/satori/go.uuid"

type ReleaseGroupSecondaryType struct {
	ID   int64
	GID  uuid.UUID
	Name string
}
