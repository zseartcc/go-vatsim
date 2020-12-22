package json3

import (
	"encoding/json"
	"testing"

	"github.com/zseartcc/go-vatsim/fixtures/json3"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Pilot(t *testing.T) {
	pilot := Pilot{}
	err := json.Unmarshal(json3.Fixture.Pilot, &pilot)
	require.NoError(t, err)

	assert.NotEmpty(t, pilot.CID)
	assert.NotEmpty(t, pilot.Name)
	assert.NotEmpty(t, pilot.Callsign)
	assert.NotEmpty(t, pilot.Server)
	assert.NotEmpty(t, pilot.PilotRating)
	assert.NotEmpty(t, pilot.Latitude)
	assert.NotEmpty(t, pilot.Longitude)
	assert.NotEmpty(t, pilot.Altitude)
	assert.NotEmpty(t, pilot.Groundspeed)
	assert.NotEmpty(t, pilot.Transponder)
	assert.NotEmpty(t, pilot.Heading)
	assert.NotEmpty(t, pilot.QNHInchesOfMercury)
	assert.NotEmpty(t, pilot.QNHMillibars)
	assert.NotEmpty(t, pilot.FlightPlan)
	assert.NotEmpty(t, pilot.LogonTime)
	assert.NotEmpty(t, pilot.LastUpdated)
}
