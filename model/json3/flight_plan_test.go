package json3

import (
	"encoding/json"
	"testing"

	"github.com/zseartcc/go-vatsim/fixtures/json3"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_FlightPlan(t *testing.T) {
	flightPlan := FlightPlan{}
	err := json.Unmarshal(json3.Fixture.FlightPlan, &flightPlan)
	require.NoError(t, err)

	assert.NotEmpty(t, flightPlan.FlightRules)
	assert.NotEmpty(t, flightPlan.Aircraft)
	assert.NotEmpty(t, flightPlan.Departure)
	assert.NotEmpty(t, flightPlan.Arrival)
	assert.NotEmpty(t, flightPlan.Alternate)
	assert.NotEmpty(t, flightPlan.CruiseTAS)
	assert.NotEmpty(t, flightPlan.Altitude)
	assert.NotEmpty(t, flightPlan.DepTime)
	assert.NotEmpty(t, flightPlan.EnrouteTime)
	assert.NotEmpty(t, flightPlan.FuelTime)
	assert.NotEmpty(t, flightPlan.Remarks)
	assert.NotEmpty(t, flightPlan.Route)
}
