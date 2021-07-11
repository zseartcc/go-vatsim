package json3

import (
	"testing"

	"github.com/zseartcc/go-vatsim/model/json3"

	"github.com/stretchr/testify/assert"
)

func Test_ControllersToMap(t *testing.T) {
	controllerMap := ControllersToMap([]json3.Controller{
		{
			CID:  123,
			Name: "Test Controller",
		},
		{
			CID:  234,
			Name: "Controller",
		},
	})

	assert.NotNil(t, controllerMap[123])
	assert.Equal(t, "Test Controller", controllerMap[123].Name)
	assert.NotNil(t, controllerMap[234])
	assert.Equal(t, "Controller", controllerMap[234].Name)
	assert.Nil(t, controllerMap[456])
}
