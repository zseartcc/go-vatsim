package json3

import (
	"encoding/json"
	"testing"

	"github.com/zseartcc/go-vatsim/fixtures/json3"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Prefile(t *testing.T) {
	prefile := Prefile{}
	err := json.Unmarshal(json3.Fixture.Prefile, &prefile)
	require.NoError(t, err)

	assert.NotEmpty(t, prefile.CID)
	assert.NotEmpty(t, prefile.Name)
	assert.NotEmpty(t, prefile.Callsign)
	assert.NotEmpty(t, prefile.FlightPlan)
	assert.NotEmpty(t, prefile.LastUpdated)
}
