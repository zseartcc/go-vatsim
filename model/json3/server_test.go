package json3

import (
	"encoding/json"
	"testing"

	"github.com/zseartcc/go-vatsim/fixtures/json3"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Server(t *testing.T) {
	server := Server{}
	err := json.Unmarshal(json3.Fixture.Server, &server)
	require.NoError(t, err)

	assert.NotEmpty(t, server.Ident)
	assert.NotEmpty(t, server.HostnameOrIP)
	assert.NotEmpty(t, server.Location)
	assert.NotEmpty(t, server.Name)
	assert.NotEmpty(t, server.ClientsConnectionAllowed)
}
