package pgmb

import (
	"bytes"
	"regexp"
	"testing"
)

func TestEchoSQL(t *testing.T) {
	buf := &bytes.Buffer{}
	EchoSQL(buf)(ArtistQuery().Where("id = ?", 999))
	re := regexp.MustCompile("SELECT id, gid.* FROM artist")
	if !re.MatchString(buf.String()) {
		t.Errorf("Echo SQL output was different than expected '%s'", buf)
	}
}
