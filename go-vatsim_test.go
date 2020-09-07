package govatsim

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_VATSIM(t *testing.T) {
	vatsim := NewVATSIM()
	clients, err := vatsim.Clients()
	assert.NoError(t, err)
	assert.NotEmpty(t, clients)
}
