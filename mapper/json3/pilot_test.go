package json3

import (
	"testing"

	"github.com/zseartcc/go-vatsim/model/json3"

	"github.com/stretchr/testify/assert"
)

func Test_PilotsToMap(t *testing.T) {
	pilotMap := PilotsToMap([]json3.Pilot{
		{
			CID:  123,
			Name: "Test Pilot",
		},
		{
			CID:  234,
			Name: "Pilot",
		},
	})

	assert.NotNil(t, pilotMap[123])
	assert.Equal(t, "Test Pilot", pilotMap[123].Name)
	assert.NotNil(t, pilotMap[234])
	assert.Equal(t, "Pilot", pilotMap[234].Name)
	assert.Nil(t, pilotMap[456])
}
