package json3

import (
	"encoding/json"
	"testing"

	"github.com/zseartcc/go-vatsim/fixtures/json3"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_General(t *testing.T) {
	general := General{}
	err := json.Unmarshal(json3.Fixture.General, &general)
	require.NoError(t, err)

	assert.NotEmpty(t, general.Version)
	assert.NotEmpty(t, general.Reload)
	assert.NotEmpty(t, general.Update)
	assert.NotEmpty(t, general.UpdateTimestamp)
	assert.NotEmpty(t, general.ConnectedClients)
	assert.NotEmpty(t, general.UniqueUsers)
}
