package govatsim

import (
	"testing"

	"github.com/zseartcc/go-vatsim/fixtures"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ParseClient(t *testing.T) {
	client, err := parseClient(fixtures.Clients[0])
	require.NoError(t, err)

	assert.NotNil(t, client)
	assert.Equal(t, "KAL6831", client.Callsign)
}

func Test_ParseTime(t *testing.T) {
	time, err := parseTime("20200907060246")
	assert.NoError(t, err)
	assert.Equal(t, int64(1599458566), time.Unix())
}
