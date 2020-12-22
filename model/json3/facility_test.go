package json3

import (
	"encoding/json"
	"testing"

	"github.com/zseartcc/go-vatsim/fixtures/json3"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Facility(t *testing.T) {
	facility := Facility{}
	err := json.Unmarshal(json3.Fixture.Facility, &facility)
	require.NoError(t, err)

	assert.NotEmpty(t, facility.ID)
	assert.NotEmpty(t, facility.Short)
	assert.NotEmpty(t, facility.Long)
}
