package pgmb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhereCreditOrArtistNameMatches(t *testing.T) {
	t.Run("with a match on artist_credit.name", func(t *testing.T) {
		acs, err := ArtistCredits(TESTDB).WhereCreditOrArtistNameMatches("Crosby, Stills & Nash").All()
		assert.Nil(t, err)
		assert.True(t, len(acs) > 0, "Expected at least 1 match for 'Crosby, Stills & Nash")
		assert.EqualValues(t, 1048685, acs[0].ID)
	})
}
