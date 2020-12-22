package json3

import (
	"encoding/json"
	"testing"

	"github.com/zseartcc/go-vatsim/fixtures/json3"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Rating(t *testing.T) {
	rating := Rating{}
	err := json.Unmarshal(json3.Fixture.Rating, &rating)
	require.NoError(t, err)

	assert.NotEmpty(t, rating.ID)
	assert.NotEmpty(t, rating.Short)
	assert.NotEmpty(t, rating.Long)
}
