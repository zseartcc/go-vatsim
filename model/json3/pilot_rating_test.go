package json3

import (
	"encoding/json"
	"testing"

	"github.com/zseartcc/go-vatsim/fixtures/json3"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_PilotRating(t *testing.T) {
	pilotRating := PilotRating{}
	err := json.Unmarshal(json3.Fixture.PilotRating, &pilotRating)
	require.NoError(t, err)

	assert.NotEmpty(t, pilotRating.ID)
	assert.NotEmpty(t, pilotRating.Short)
	assert.NotEmpty(t, pilotRating.Long)
}
