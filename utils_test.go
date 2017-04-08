package gorandom

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestUtils_pnInt16(t *testing.T) {
	assert.Equal(t, pnInt16(-1), int16(1))
	assert.Equal(t, pnInt16(0), int16(1))
	assert.Equal(t, pnInt16(1), int16(1))
	assert.Equal(t, pnInt16(2), int16(2))
}

func TestUtils_pnInt32(t *testing.T) {
	assert.Equal(t, pnInt32(-1), int32(1))
	assert.Equal(t, pnInt32(0), int32(1))
	assert.Equal(t, pnInt32(1), int32(1))
	assert.Equal(t, pnInt32(2), int32(2))
}

func TestUtils_pnInt(t *testing.T) {
	assert.Equal(t, pnInt(-1), int(1))
	assert.Equal(t, pnInt(0), int(1))
	assert.Equal(t, pnInt(1), int(1))
	assert.Equal(t, pnInt(2), int(2))
}

func TestUtils_pnInt64(t *testing.T) {
	assert.Equal(t, pnInt64(-1), int64(1))
	assert.Equal(t, pnInt64(0), int64(1))
	assert.Equal(t, pnInt64(1), int64(1))
	assert.Equal(t, pnInt64(2), int64(2))
}
