package json3

import (
	"encoding/json"
	"testing"

	"github.com/zseartcc/go-vatsim/fixtures/json3"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_JSON3(t *testing.T) {
	JSON3 := JSON3{}
	err := json.Unmarshal(json3.Fixture.JSON3, &JSON3)
	require.NoError(t, err)

	assert.NotEmpty(t, JSON3.General)
	assert.NotEmpty(t, JSON3.Pilots)
	assert.NotEmpty(t, JSON3.Controllers)
	assert.NotEmpty(t, JSON3.ATIS)
	assert.NotEmpty(t, JSON3.Servers)
	assert.NotEmpty(t, JSON3.Prefiles)
	assert.NotEmpty(t, JSON3.Facilities)
	assert.NotEmpty(t, JSON3.Ratings)
	assert.NotEmpty(t, JSON3.PilotRatings)
}
