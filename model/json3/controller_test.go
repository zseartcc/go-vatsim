package json3

import (
	"encoding/json"
	"testing"

	"github.com/zseartcc/go-vatsim/fixtures/json3"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Controller(t *testing.T) {
	controller := Controller{}
	err := json.Unmarshal(json3.Fixture.Controller, &controller)
	require.NoError(t, err)

	assert.NotEmpty(t, controller.CID)
	assert.NotEmpty(t, controller.Name)
	assert.NotEmpty(t, controller.Callsign)
	assert.NotEmpty(t, controller.Frequency)
	assert.NotEmpty(t, controller.Facility)
	assert.NotEmpty(t, controller.Rating)
	assert.NotEmpty(t, controller.Server)
	assert.NotEmpty(t, controller.VisualRange)
	assert.NotEmpty(t, controller.TextATIS)
	assert.NotEmpty(t, controller.LastUpdated)
	assert.NotEmpty(t, controller.LogonTime)
}
