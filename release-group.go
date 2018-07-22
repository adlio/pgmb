package pgmb

import (
	"github.com/satori/go.uuid"
)

// ReleaseGroup represents an entry in the MusicBrainz database
// release_group table.
type ReleaseGroup struct {
	ID      int64
	gid     uuid.UUID
	Name    string
	Comment string
}
